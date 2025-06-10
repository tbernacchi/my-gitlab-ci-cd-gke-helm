# GitLab CI/CD using GKE and Helm 👋

> A comprehensive pipeline setup with GitLab CI/CD using Google Kubernetes Engine (GKE) and Helm.

![GitLab](/.github/assets/img/gitlab-pipe.png)

<div align="center">
	<img align="center" width="300px" src="/.github/assets/img/google-cloud-logo.png">
</div>

## Requirements

Before you begin, ensure you have the following installed:
* GitLab account
* Docker installed locally
* Google Cloud Platform (GCP) account
* `kubectl` - Kubernetes command-line tool
* `helm` - Kubernetes package manager
* `gcloud` - Google Cloud SDK

## Step by Step Guide

### 1. Create a Kubernetes Cluster on GKE

```bash
# Create a new cluster
gcloud container clusters create my-k8s-dev --zone southamerica-east1-a --project <your-project-id>

# Verify your current context
kubectl config current-context 

# Configure kubectl to use the new cluster
gcloud container clusters get-credentials my-k8s-dev --zone southamerica-east1-a
```

### 2. Set up GCP Service Account

1. Create a service account in GCP to allow GitLab to interact with GKE
   * Navigate to GCP Console > IAM & Admin > Service Accounts
   * Create a new service account with appropriate permissions
   * Generate and download a JSON key file
   * For detailed instructions, visit: https://cloud.google.com/compute/docs/access/service-accounts

### 3. Configure GitLab Registry Access

Create a Kubernetes secret to allow GKE to pull images from GitLab's registry:

```bash
kubectl create secret docker-registry gitlab-registry \
  --docker-server=registry.gitlab.com \
  --docker-username=<your-gitlab-username> \
  --docker-password=<your-registry-token> \
  --docker-email=<your-email>
```

Then create a Deploy Token in GitLab:
1. Go to your GitLab project
2. Navigate to Settings > Repository > Deploy tokens
3. Create a new token with appropriate permissions

### 4. Create Kubernetes Service Account for GitLab

Create the service account configuration file:

```bash
# Create the configuration file
cat << EOF > gitlab-admin-service-account.yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: gitlab
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: gitlab-admin
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: gitlab
    namespace: kube-system
EOF

# Apply the configuration
kubectl create -f gitlab-admin-service-account.yaml
```

### 5. Integrate GKE with GitLab

Collect the following information to configure in GitLab:

```bash
# Get the cluster API URL
kubectl cluster-info | grep -E 'Kubernetes master|Kubernetes control plane' | awk '/http/ {print $NF}'

# Get the CA Certificate
kubectl get secret default-token-l76rk -o jsonpath="{['data']['ca\.crt']}" | base64 --decode

# Get the Service Account Token
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep gitlab | awk '{print $1}')
```

### 6. Configure GitLab CI/CD

1. In your GitLab project, go to Infrastructure > Kubernetes
2. Add a new Kubernetes cluster
3. Fill in the following information:
   * API URL (from step 5)
   * CA Certificate (from step 5)
   * Service Account Token (from step 5)
4. Save the configuration

### 7. Configure Helm

Helm will be used to manage your Kubernetes applications. Make sure your `.gitlab-ci.yml` includes Helm commands for deployment.

Example `.gitlab-ci.yml` structure:
```yaml
stages:
  - build
  - deploy

build:
  stage: build
  script:
    - docker build -t $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA .
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA

deploy:
  stage: deploy
  script:
    - helm upgrade --install my-app ./helm-chart \
      --set image.tag=$CI_COMMIT_SHA \
      --namespace my-namespace
```

## Troubleshooting

Common issues and their solutions:
* If you can't pull images from GitLab registry, verify your registry secret is correctly configured
* If GitLab can't connect to GKE, check your service account permissions
* For Helm issues, ensure your chart is properly configured and validated

## References
* [GitLab Kubernetes Integration](https://docs.gitlab.com/ee/user/project/clusters/add_remove_clusters.html)
* [GitLab Cluster Configuration Guide](https://about.gitlab.com/handbook/customer-success/demo-systems/tutorials/getting-started/configuring-group-cluster)
* [Automated Kubernetes Deployments](https://medium.com/@yanick.witschi/automated-kubernetes-deployments-with-gitlab-helm-and-traefik-4e54bec47dcf)

## Author

👤 **Tadeu Bernacchi**

* Website: http://www.tadeubernacchi.com.br/
* Twitter: [@tadeuuuuu](https://twitter.com/tadeuuuuu)
* Github: [@tbernacchi](https://github.com/tbernacchi)
* LinkedIn: [@tadeubernacchi](https://linkedin.com/in/tadeubernacchi)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_

<h1 align="">Gitlab CI/CD using GKE and Helm 👋</h1>
<p>
</p>

> My pipeline with Gitlab CI/CD using GKE and Helm.

![GitLab](/.github/assets/img/gitlab-pipe.png)

<div align=>
	<img align="center" width="300px" src=/.github/assets/img/google-cloud-logo.png>
</div>


## Requirements
* Gitlab;
* Docker;
* GCP;
* kubectl;
* helm;

## Usage

```
gcloud container clusters create my-k8s-dev --zone southamerica-east1-a --project <my-project>
kubectl config current-context 
gcloud container clusters get-credentials my-k8s-dev --zone southamerica-east1-a
````

You're going to need to create an account service on GCP to allow Gitlab to read/write on GKE:

> https://cloud.google.com/compute/docs/access/service-accounts

With the account service created you need to generate a key to allow GKE access GitLab Docker registry to pull the docker images, with the key:

```
kubectl create secret docker-registry gitlab-registry --docker-server=registry.gitlab.com --docker-username=tbernacchi --docker-password=my-key-generate-at-iam-console-at-gcp --docker-email=tbernacchi@gmail.com
```

At GitLab we create a deploy token with the same key (Settings, Repository, Deploy tokens);

We also must create an account into the k8s cluster for gitlab, it must have cluster-admin privileges:

```
$ touch gitlab-admin-service-account.yaml 
```

```
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
```

```
kubectl create -f gitlab-admin-service-account.yaml
```

To integrate GKE with Gitlab we need:

```
kubectl cluster-info | grep -E 'Kubernetes master|Kubernetes control plane' | awk '/http/ {print $NF}'
kubectl get secret default-token-l76rk -o jsonpath="{['data']['ca\.crt']}" | base64 --decode
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep gitlab | awk '{print $1}')
```

# References
https://docs.gitlab.com/ee/user/project/clusters/add_remove_clusters.html
https://about.gitlab.com/handbook/customer-success/demo-systems/tutorials/getting-started/configuring-group-cluster  
https://medium.com/@yanick.witschi/automated-kubernetes-deployments-with-gitlab-helm-and-traefik-4e54bec47dcf

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

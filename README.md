# My Gitlab CI/CD using GKE and Helm.

```bash
comandos
comandos
```
<h1 align="">Gitlab CI/CD using GKE and Helm 👋</h1>
<p>
</p>

> My pipeline using Gitlab CI/CD using GKE and Helm.

![Datadog](/.github/assets/img/gitlab-pipe.png)

<div align=>
	<img align="center" width="400px" src=/.github/assets/img/gitlab-pipe.png>
</div>

## Table of Contents

* **GitLab**  
  * [Official Website](https://gitlab.com/)
  * [Official Docs](https://docs.gitlab.com/)
  * [Official Github](https://github.com/gitlabhq)

* **GKE**  
  * [Official Website](https://cloud.google.com/kubernetes-engine)
  * [Official Docs](https://cloud.google.com/kubernetes-engine/docs/quickstart)
  * [Official Github](https://github.com/GoogleCloudPlatform/kubernetes-engine-samples)

* **Helm**  
  * [Official Website](https://helm.sh/)
  * [Official Docs](https://helm.sh/docs/)
  * [Official Github](https://github.com/helm/helm)

## Requirements
* Google Cloud SDK;
* gcloud CLI;  
* kubectl;  
* helm;
* Datadog account + Apikey;

## Usage

```
gcloud auth login
gcloud components update
gcloud config set project <my-project>
gcloud container clusters create my-cluster-datadog --zone southamerica-east1-a --project <my-project>
```

```
helm repo add datadog https://helm.datadoghq.com
helm repo add stable https://charts.helm.sh/stable
helm repo update
kubectl create namespace datadog
```

```
helm install tadeu-teste-datadog -f values.yml -n datadog --set datadog.site='datadoghq.com' --set datadog.apiKey=myapikeyec747bdd2b92a3fc42345678 datadog/datadog
```

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

#!/bin/bash
function init_helm() {
  mkdir -p /etc/deploy
  echo ${GKE_SERVICE_ACCOUNT} | base64 -d > /etc/deploy/sa.json;
  gcloud auth activate-service-account --key-file /etc/deploy/sa.json --project=${GKE_PROJECT};
  gcloud container clusters get-credentials ${GKE_CLUSTER_NAME} --zone ${GKE_ZONE} --project ${GKE_PROJECT};
  helm init --service-account tiller --wait --upgrade;
}


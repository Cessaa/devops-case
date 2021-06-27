#!/usr/bin/env bash
set -Eeuo pipefail

gcloud_cluster_name=$1
gcloud_cluster_region_name=$2
gcloud_project_id=$(gcloud config get-value project -q)

docker build -t go-project:latest .

docker tag \
       go-project \
       eu.gcr.io/$gcloud_project_id/go-project:latest

docker push \
       eu.gcr.io/$gcloud_project_id/go-project:latest

gcloud container clusters \
       get-credentials $gcloud_cluster_name \
       --region $gcloud_cluster_region_name

template=`cat "deployment.yml" | sed "s/{{PROJECT_ID}}/$gcloud_project_id/g"`
echo "$template" | kubectl apply -f -

kubectl apply -f service.yml
kubectl apply -f ingress.yml

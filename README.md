# Semantic search demo

### Docker compose 

```bash 

docker-compose --env-file ./.env up -d --build

```

Read logs:

```bash 

docker-compose logs -f -t

```

### Kubernetes (Google Cloud)

Create cluster with name: ss-demo

```bash 
gcloud container clusters create ss-demo \
    --num-nodes=3 --zone us-central1-a --machine-type g1-small
```

Connect the kubectl client to the cluster:

```bash 
gcloud container clusters get-credentials ss-demo --zone us-central1-a
```

Create disk with name: pg-data-disk

```bash 
gcloud compute disks create pg-data-disk --size 50GB --zone us-central1-a
```

Change google project id `PROJECT_ID` in yaml files to yours.

Build docker images:

```bash 
export PROJECT_ID= .... # your google project id

docker build -t gcr.io/${PROJECT_ID}/text2embeddings:v0.0.1 -f text2embeddings/Dockerfile .
docker push gcr.io/${PROJECT_ID}/text2embeddings:v0.0.1

docker build -t gcr.io/${PROJECT_ID}/app:v0.0.1 -f app/Dockerfile . && \
docker push gcr.io/${PROJECT_ID}/app:v0.0.1
```

Deploy DB:

```bash 
kubectl apply -f k8s/db-volume.yaml

kubectl apply -f k8s/db-volume-claim.yaml

kubectl apply -f k8s/db-secret.yaml

kubectl apply -f k8s/db-config-map.yaml

kubectl apply -f k8s/db-deployment.yaml

kubectl apply -f k8s/db-service.yaml

```

Deploy NATS service:

```bash 


kubectl apply -f k8s/nats-secret.yaml

kubectl apply -f k8s/nats-deployment.yaml

kubectl apply -f k8s/nats-service.yaml

```

Deploy text2embeddings service (todo: CUDA support):

```bash 
kubectl apply -f k8s/text2embeddings-deployment.yaml

kubectl apply -f k8s/text2embeddings-service.yaml
```

Deploy main app service:

```bash 
kubectl apply -f k8s/app-deployment.yaml

kubectl apply -f k8s/app-service.yaml
```

Check pods:

```bash 
kubectl get pods
```

## Try

Create a post with some text:

```bash 

curl --request POST  \
  --data "Man in the car ..."  \
  $host/post

```

Find a post with similar text (synonymous):

```bash 

curl --request GET \
    $host/post/search?q=dude%20in%20the%20pickup
  
```
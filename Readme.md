# Semantic search demo

# Build 

```bash 

docker-compose --env-file ./.env up -d 

```

Read logs:

```bash 

docker-compose logs -f -t

```

# Try

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
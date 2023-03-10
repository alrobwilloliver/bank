// gen migration
```
migrate create -ext sql -dir ./db/migration -seq init
```

// Catch to mockgen error 
error: "mockgen missing go.sum entry for module providing package github .com/golang/mock/mockgen/model"
```
mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/alrobwilloliver/bank/db/sqlc Store
```

// build dockerfile 
```
docker build -t bank:latest .
```
// run dockerfile 
```
docker run -p 8080:8080 bank
```

## add postgres to bank-network
```
docker network create bank-network
docker network connect bank-network postgres12
```

## command to run application
### the ip address of postgres12 can be replaced with the container name
```
docker run --name bank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres12:5432/bank?sslmode=disable" bank
```

### once deployed you can pull the image from aws ecr
## login to aws ecr
```
aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin $LINK_OF_DEPLOYED_IMAGE
```
## pull image from aws ecr
```
docker pull $LINK_OF_DEPLOYED_IMAGE
```
## run image from aws ecr
```
docker run --name bank --network bank-network -p 8080:8080 -e GIN_MODE=release $LINK_OF_DEPLOYED_IMAGE
```
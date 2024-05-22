# Go Hexa example

## prerequisites

```
go > 1.21.4
```

## configuration


### util commands
```shell
# generate prisma models
## 1 set prisma env
cp dist.env .env


## 2 fetch the golang prisma tool
go get github.com/steebchen/prisma-client-go

## 3 run the prisma generator binary
go run github.com/steebchen/prisma-client-go generate

## run the db container
docker compose up -d postgres

## 4 push the database
go run github.com/steebchen/prisma-client-go db push

## 5 fetch golang packages
go mod tidy
```

### run project locally
```shell
docker compose watch
```

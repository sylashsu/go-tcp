``# Simple golang tcp server/client

## server

```
go run sever.go -a $(ip) -p $(port)
```

## client

```
go run client.go -url $(ip:port) -file $(file)
```
---

## build

### linux
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

### mac
```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```

### windows
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

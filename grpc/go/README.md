# Rebuild

## Generation
```bash
go version
go help gopath
export PATH=$PATH:$GOPATH/bin
go get -u google.golang.org/grpc
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
sudo apt  install protobuf-compiler
protoc -I proto/ proto/kubedge.proto --go_out=plugins=grpc:go
```
- Do not forget the plugins=grpc otherwise you will be missing code


## On server:

```bash
go run kubedge_server/main.go
```

## On client:
```bash
go run kubedge_client/main.go
```

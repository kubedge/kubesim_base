# Rebuild

```bash
go version
go help gopath
export PATH=$PATH:$GOPATH/bin
go get -u google.golang.org/grpc
go get -u github.com/kubedge/kubesim_base/grpc/go
```

On server:
```bash
go run kubedge_server/main.go
```

On client:
```bash
go run kubedge_client/main.go
```

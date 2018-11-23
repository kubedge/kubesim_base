# Use Swagger to implement readiness, liveness and config api command for any KUBEDGE simulator

```bash
swagger generate server --exclude-main -A cfgapisrv -t . -f ./swagger/swagger.yml
```

Running the Config API Server

```bash
$ go run ./main.go
```

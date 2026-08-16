# Tasks — go-version-bump

- [x] Decide the target go line (align with the operator side where practical). → **go 1.26** (matches the kubedge-operator-base pilot + local toolchain).
- [x] Raise the `go` directive in each module go.mod (config, connected, grpc/go, health, arpscan, root). → all six `go 1.20` → `go 1.26.0`.
- [x] `go mod tidy` per module; fix any breakage from the bump.
  - `grpc/go`: the go 1.24+ vet check `non-constant format string` (gated on the go directive) fired on 26 `log.Printf(message)` calls in `kubedge_server/server.go` → changed to `log.Print(message)` (no format args). Also fixed a pre-existing `log.Fatalf("%s...", simuname, err)` 1-verb/2-arg call (dropped the stray `err`).
  - `go mod tidy` backfilled two missing `/go.mod` hash lines in `grpc/go/go.sum` + `health/go.sum` (benign).
- [x] `go build ./... && go vet ./... && go test ./... -race` green in each module. → all 6 green.
- [ ] Sequence with `multimodule-tag-realign` (bump then re-tag together). → handled by the `multimodule-tag-realign` change; the go bump lands first, the re-tag/release contract follows there.

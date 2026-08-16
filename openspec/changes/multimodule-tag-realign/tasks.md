# Tasks — multimodule-tag-realign

- [ ] Pick the next `<version>` (plain `v0.1.NN`, > current max 0.1.24).
- [ ] Bump `grpc/go/go.mod`'s `github.com/kubedge/kubesim_base/config` require to `<version>`; `go mod tidy` in `grpc/go`.
- [ ] Build/vet/test each consumed module (`config`, `connected`, `grpc/go`) green.
- [ ] Tag every consumed sub-module path together: `config/<version>`, `connected/<version>`, `grpc/go/<version>` (and `health/<version>` only if it is ever consumed).
- [ ] Push tags. Verify `go list -m github.com/kubedge/kubesim_base/config@<version>` resolves.
- [ ] Hand off: consumers realign all three requires to `<version>` (their own realign change).

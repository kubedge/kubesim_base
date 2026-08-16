# Tasks — multimodule-ci

- [x] `detect` job emits a `gomodules` JSON output listing every dir with a `go.mod`. → `[".","arpscan","config","connected","grpc/go","health"]`.
- [x] `go-build` / `go-vet` / `go-test` run as a matrix over `fromJSON(detect.gomodules)`, each `cd`-ing into the module and using its `go.mod` for `setup-go` (`fail-fast: false`).
- [x] Keep `go-build` as `go build -o /dev/null ./...` (cmdrun collision).
- [x] `golangci-lint` left root-only; multi-module-lint follow-up noted (blocked by lint-cleanup).
- [x] Fix arpscan CI breakage: it imports `gopacket/pcap` (cgo), so the ubuntu runner needs `libpcap-dev` (present on macOS, hence green locally). Added a `matrix.module == 'arpscan'` step installing `libpcap-dev` to each go job.
- [ ] Trial push; confirm build/vet/test are green for all 6 modules (matrix cells). → pending PR CI.

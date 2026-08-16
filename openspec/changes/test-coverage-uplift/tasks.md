# Tasks — test-coverage-uplift

- [ ] `go test ./... -cover` per module; note 0% packages.
- [ ] Test `config.Configdata` construction/round-trip.
- [ ] Test `connected.Connecteddata` state handling.
- [ ] Test the `grpc/go` client/server path against the protobuf messages (encode/decode).
- [ ] Keep `go test ./... -race` green; record the coverage delta.

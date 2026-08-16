## ADDED Requirements

### Requirement: The shared config, connection, and RPC modules are unit-tested

The sim base SHALL have unit tests for `config.Configdata`, `connected.Connecteddata`,
and the `grpc/go` client/server against the protobuf message set, run under
`go test ./... -race`, so the contract the sim fleet imports is protected.

#### Scenario: shared modules have passing tests
- **WHEN** `go test ./... -race` runs across the consumed modules
- **THEN** config, connected, and grpc/go have passing tests (no longer 0 coverage)

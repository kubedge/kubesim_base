# Uplift test coverage

## Why

The sim base carries the shared config/connection types and the RPC protocol every
simulator depends on, but has little/no test coverage. Since a change here ripples to the
whole sim fleet, the shared contract deserves tests.

## What Changes

- Add unit tests for the highest-value shared pieces: `config.Configdata` and
  `connected.Connecteddata` round-trips, and the `grpc/go` client/server against the
  protobuf message set.
- Wire `go test ./... -race` into the local loop and CI per module.

## Capabilities

### New Capabilities
- test-coverage: the sim base's shared contract is protected by tests.

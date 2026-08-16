# go-toolchain Specification

## Purpose

The Go language/toolchain version the sim base targets. The base publishes multiple
modules from one repo, and its `go` directive is the floor every `kubedge-sim-*` /
`kubesim_*` consumer must meet to build against it.

## Requirements

### Requirement: Sim-base modules target a current Go toolchain

Every sim-base module (`config`, `connected`, `grpc/go`, and root/`health`/`arpscan`)
SHALL declare a current, supported `go` directive (raised from `go 1.20`; currently
`go 1.26.0`), and each module SHALL build/vet/test green after `go mod tidy`.

#### Scenario: modules build on the bumped toolchain
- **WHEN** `go build ./... && go test ./... -race` runs in each module after the bump
- **THEN** they pass, and no module still declares the stale `go 1.20`/`1.15` line

### Requirement: CI lint tooling matches the toolchain floor

The `golangci-lint` used in CI SHALL be built with a Go version at least the module's
`go` directive, because golangci-lint refuses to lint a module whose language version
exceeds the Go it was built with.

#### Scenario: golangci-lint can lint the bumped modules
- **WHEN** the CI `golangci-lint` job runs against a `go 1.26` module
- **THEN** it uses a golangci-lint (v2.12.2, `golangci-lint-action@v8`) built with go1.26 and does not fail with a "language version lower than targeted" error

# Multi-module tag + realign strategy

## Why

This repo publishes several Go modules from one repo, and consumers pin them
independently (`config@vX`, `connected@vX`, `grpc/go@vX`). Today the fleet is split across
`v0.1.24` (sim-ecds/mme, linkio) and `v0.1.20` (5gc/elte/epc/lte/nr), and the internal
`grpc/go → config` edge must stay consistent. There is no defined, repeatable release
procedure — so realigning the sims safely needs one.

## What Changes

- Define the release procedure: cut ONE new tag `<version>` and apply it to every
  **consumed** sub-module path together (`config/<version>`, `connected/<version>`,
  `grpc/go/<version>`), after bumping the internal `grpc/go/go.mod` `config` require to
  the same `<version>`.
- Publish; then consumers realign all their kubesim_base require lines to `<version>` in
  one `go get ... && go mod tidy`.
- Tag scheme stays plain `v0.1.NN` (no k8s encoding).

## Capabilities

### New Capabilities
- module-versioning: how the multi-module sim base is tagged and consumed.

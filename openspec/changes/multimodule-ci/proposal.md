# Multi-module CI coverage

## Why

The sim base publishes six Go modules (root, `config`, `connected`, `grpc/go`, `health`,
`arpscan`), each with its own `go.mod`. The shipped class-M CI runs the Go jobs only at
the repo root, and `go build/vet/test ./...` stops at nested module boundaries — so the
five sub-modules, where the real code (and the forthcoming `test-coverage-uplift` tests)
live, are never gated. `adopt-go-ci` shipped root-only and deferred this.

## What Changes

- `detect` emits the list of module directories (every dir containing a `go.mod`) as a
  JSON array output.
- `go-build`, `go-vet`, `go-test` run as a matrix over that list, so each module is
  built / vetted / tested independently against its own `go.mod`.
- `golangci-lint` stays root-only for now — multi-module lint is blocked by ~33 findings
  across the five sub-modules (the lint-cleanup initiative). Matrixing it is a follow-up
  once those are clean.

## Non-goals

- Lint-cleanup of the sub-modules (separate initiative).
- Changing the Makefile targets (CI runs raw `go`).

## Impact

Every module's build/vet/test green is enforced on push, which is the prerequisite for
gating `test-coverage-uplift`. Fleet-wide: the sim consumers are also multi-module, so
this CI shape is a template for them.

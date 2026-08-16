# continuous-integration Specification

## Purpose

The push-time gate that enforces build/vet/test/lint green, independently of the
Makefile. Carried by the class-M `.github/workflows/ci.yml` shipped from claude-meta.

## Requirements

### Requirement: The repo runs Go CI on every push

The repository SHALL carry the class-M `.github/workflows/ci.yml` that auto-detects
`go.mod` and runs `go-build`, `go-vet`, `go-test -race`, and `golangci-lint` as required
checks, so build/vet/test/lint green is enforced on every push independently of the
Makefile.

The `go-build` job SHALL be a pure compile check (`go build -o /dev/null ./...`), because
the root module's sole package is `main` at `cmdrun/` and a plain `go build ./...` would
emit a `cmdrun` binary that collides with the `cmdrun/` directory.

#### Scenario: CI gates a push
- **WHEN** a commit is pushed to a branch with `go.mod` present
- **THEN** the four Go jobs run and a failure blocks the change

#### Scenario: go-build does not emit an artifact
- **WHEN** the `go-build` job runs at the repo root
- **THEN** it compile-checks with `-o /dev/null` and does not write a binary that would collide with the `cmdrun/` directory

## MODIFIED Requirements

### Requirement: The repo runs Go CI on every push

The repository SHALL carry the class-M `.github/workflows/ci.yml` that auto-detects
`go.mod` and runs `go-build`, `go-vet`, `go-test -race`, and `golangci-lint` as required
checks, so build/vet/test/lint green is enforced on every push independently of the
Makefile.

Because the base is multi-module, the `go-build`, `go-vet`, and `go-test` jobs SHALL run
as a **matrix over every module** (every directory containing a `go.mod`), each built /
vetted / tested against its own `go.mod` — not only the root module, whose `./...` stops
at nested module boundaries. The `detect` job SHALL emit the module list as a JSON output
that these jobs consume.

The `go-build` job SHALL be a pure compile check (`go build -o /dev/null ./...`), because
the root module's sole package is `main` at `cmdrun/` and a plain `go build ./...` would
emit a `cmdrun` binary that collides with the `cmdrun/` directory.

The `golangci-lint` job MAY run at the root module only until sub-module lint findings are
cleaned up; multi-module lint is a follow-up.

#### Scenario: CI gates a push
- **WHEN** a commit is pushed to a branch with `go.mod` present
- **THEN** the Go jobs run and a failure blocks the change

#### Scenario: every module is built, vetted, and tested
- **WHEN** the Go jobs run
- **THEN** `go-build`, `go-vet`, and `go-test -race` each execute once per module (root, `config`, `connected`, `grpc/go`, `health`, `arpscan`) against that module's `go.mod`

#### Scenario: go-build does not emit an artifact
- **WHEN** the `go-build` job runs for a module
- **THEN** it compile-checks with `-o /dev/null` and does not write a binary that would collide with the `cmdrun/` directory

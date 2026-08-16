# Tasks — adopt-go-ci

- [x] From the meta session, run `/alemax:update-skills` so the class-M set (incl. `ci.yml`) is staged for this repo.
- [x] In this repo's session, run `/alemax:complete-update` to apply the update branch onto the working branch.
- [x] Confirm `.github/workflows/ci.yml` present and its jobs gate on `go.mod`.
- [x] Trial push the branch; confirm `go-build`/`go-vet`/`go-test`/`golangci-lint` are green.
  - Root `go-build` initially failed: the root module's sole package is `main` at
    `cmdrun/`, so `go build ./...` emits a `cmdrun` binary into the repo root and collides
    with the `cmdrun/` directory. Fixed by switching the job to `go build -o /dev/null ./...`
    (pure compile check, no artifact). `go-vet`/`go-test`/`golangci-lint` were already green at root.
- [x] Confirm the rest of class-M landed: `.editorconfig`, `.gitattributes`, `.github/*`, `dependabot.yml`, `.pre-commit-config.yaml`, `bin/set-secret.sh`.

## Deferred — multi-module coverage (follow-up, not this change)

The shipped go jobs run at the **root module only**; `go build/vet/test ./...` stops at
nested `go.mod` boundaries, so `config/ connected/ grpc/go/ health/ arpscan/` are NOT
covered. Making the four checks multi-module + green is blocked by other initiatives and
is intentionally out of scope here:

- `golangci-lint` is red in all 5 submodules → **lint-cleanup** initiative.
- `grpc/go` has a `go vet`/`go test` failure (`log.Fatalf` 2-arg call at
  `kubedge_server/server.go:246`) → lint/vet cleanup.

Track multi-module CI as its own change once those land.

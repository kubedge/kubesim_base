# sim-config Specification

## Purpose

The shared configuration data type simulators use. Published as its own module
`github.com/kubedge/kubesim_base/config` (`config/config.go`, package `config`).

## Requirements

### Requirement: A shared Configdata type carries simulator configuration

The `config` module SHALL export a `Configdata` type that consumer simulators
(`kubedge-sim-*`, `kubesim_*`) import to hold their runtime configuration, so all sims
share one configuration shape rather than each defining its own.

#### Scenario: a simulator loads its configuration
- **WHEN** a simulator constructs its config
- **THEN** it uses `config.Configdata` from `github.com/kubedge/kubesim_base/config` at the pinned tag

### Requirement: config is an independently versioned module

The `config` module SHALL be a standalone Go module (its own `config/go.mod`) tagged
independently (`config/vX.Y.Z`), because the sim base publishes multiple modules from one
repo and consumers pin `config` on its own line.

#### Scenario: config is pinned on its own require line
- **WHEN** a consumer's go.mod is inspected
- **THEN** it requires `github.com/kubedge/kubesim_base/config` at a specific tag, separate from other kubesim_base sub-modules

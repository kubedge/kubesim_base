# sim-connected Specification

## Purpose

The shared connection/topology state type simulators use. Published as its own module
`github.com/kubedge/kubesim_base/connected` (`connected/connected.go`, package `connected`).

## Requirements

### Requirement: A shared Connecteddata type represents simulator connection state

The `connected` module SHALL export a `Connecteddata` type that simulators import to
represent their connection/topology state, keeping that model consistent across the sim
fleet.

#### Scenario: a simulator tracks connection state
- **WHEN** a simulator manages its peers/topology
- **THEN** it uses `connected.Connecteddata` from `github.com/kubedge/kubesim_base/connected` at the pinned tag

### Requirement: connected is an independently versioned module

The `connected` module SHALL be a standalone Go module (`connected/go.mod`) tagged
independently (`connected/vX.Y.Z`) and pinned on its own require line by consumers.

#### Scenario: connected is pinned on its own require line
- **WHEN** a consumer's go.mod is inspected
- **THEN** it requires `github.com/kubedge/kubesim_base/connected` at a specific tag

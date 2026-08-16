# sim-grpc-protocol Specification

## Purpose

The gRPC/protobuf contract simulators communicate over, plus the generated client and
server. Published as `github.com/kubedge/kubesim_base/grpc/go` (`grpc/go/kubedge.pb.go`,
`grpc/go/kubedge_client`, `grpc/go/kubedge_server`).

## Requirements

### Requirement: A shared protobuf message set defines the sim RPC protocol

The `grpc/go` module SHALL define the generated protobuf message types the simulators
exchange — including the telecom/EPC/MME simulation messages (`EnodeRequest`, `EPCReply`,
`UERequest`, `SIMReply`, `MME_PDN_CONNECT_MSG`, …) — as the single source of truth for the
wire format across the sim fleet.

#### Scenario: two simulators speak the same protocol
- **WHEN** a sim client and a sim server exchange messages
- **THEN** both use the message types from `github.com/kubedge/kubesim_base/grpc/go` at a compatible tag

### Requirement: The module ships generated client and server helpers

The module SHALL provide `kubedge_client` (a `Client(...)` entrypoint) and
`kubedge_server` helpers so simulators connect/serve without regenerating boilerplate.

#### Scenario: a simulator uses the shared client
- **WHEN** a simulator initiates an RPC
- **THEN** it calls `kubedge_client.Client(...)` from this module

### Requirement: grpc/go depends on config within the sim base

The `grpc/go` module SHALL import `kubesim_base/config` (`grpc/go` requires
`config`) — an internal edge inside the multi-module base — so the client is configured
from the shared `Configdata`. A base re-tag MUST bump this internal `grpc/go → config`
require in lockstep.

#### Scenario: internal edge stays consistent on re-tag
- **WHEN** the sim base is re-tagged
- **THEN** `grpc/go/go.mod`'s `config` require is bumped to the same new tag as the published `config` module

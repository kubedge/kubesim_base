## ADDED Requirements

### Requirement: Consumed sub-modules are tagged together with a lockstep internal edge

A release of the sim base SHALL tag all externally-consumed sub-module paths
(`config`, `connected`, `grpc/go`) with the same `<version>` in one release, and the
internal `grpc/go → config` require SHALL be bumped to that same `<version>` before
tagging, so a consumer can pin all three at one consistent version.

#### Scenario: a release produces consistent sub-module tags
- **WHEN** the sim base is released as `<version>`
- **THEN** `config/<version>`, `connected/<version>`, and `grpc/go/<version>` all exist and `grpc/go`'s `config` require equals `<version>`

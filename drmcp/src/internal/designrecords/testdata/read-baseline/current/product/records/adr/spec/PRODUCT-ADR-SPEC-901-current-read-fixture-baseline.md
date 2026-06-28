# PRODUCT-ADR-SPEC-901: Use shared current read fixtures

- **status**: accepted
- **date**: 2026-06-28
- **depends_on**: []
- **supersedes**: []
- **migrated_to_spec**: null

## Context

Current read implementation needs stable app-aware fixture records across separate app roots.

## Decision

Use package-local current fixtures with canonical PRODUCT and DRMCP identities.

## Rationale

Stable files separate fixture structure from operation-specific runtime assertions.

## Rejected alternatives

- Bare IDs: They are not canonical current references.
- Repository bootstrap data: It mixes test input with the working repository.

## Consequences

Implementation tests may consume this record through the shared fixture manifest.

## Evidence

`DRMCP-TASK-MCP-008-02` owns this fixture file.

# DRMCP-ADR-MCP-012: Unified Current Records state and application lifecycle

- **status**: accepted
- **date**: 2026-07-04
- **depends_on**:
  - DRMCP-ADR-MCP-001
- **supersedes**:
  - DRMCP-ADR-MCP-009
- **migrated_to_spec**: null

## Context

ADR-009 kept Guidance outside Current Records state because Guidance used a separate source capability.
That source model conflicts with the portable standards package authority.

PRODUCT-REQ-SPEC-003 produces an ordinary current Spec tree rooted at `<exe-dir>/design-records/`.
The package uses `spec:design_records.*` identities and must work outside the host repository.

The application needs one state model that reuses Current Records semantics while preserving Legacy Archive separation and request freshness.

## Decision

Composition / Lifecycle owns runtime configuration, dependency construction, wiring, server startup, and ordered shutdown.
Runtime configuration remains immutable for one server lifetime.

Configuration selects one portable standards spec-tree source and associates it with `app_namespace: design_records`.
The source may use the bundled default or one explicit override.
An explicit override does not silently fall back to the bundled source.
Exact configuration serialization remains downstream.

A configured current source may expose:

- a normal records root whose Spec tree is below `spec/`; or
- an explicit spec-tree root whose own `index.md` is the app root Spec.

Both source forms use the same current Spec discovery, path-derived identity, parsing, logical-tree, active-index, retrieval, resolution, and validation semantics.
DRMCP performs no runtime namespace rewrite.

Each Read, Validation, or Guidance request builds one fresh immutable Current Records snapshot from every configured mandatory current source.
The request uses the snapshot from start to finish and discards it afterward.

The Current Records snapshot includes:

- host application records;
- portable package Specs under the `design_records` app namespace;
- one active logical tree and index across current identities;
- current validation subjects and relation inputs.

The active index retains each source's explicit app association.
The portable package does not receive a separate package index, snapshot, cache, or lifecycle.

Current Records and Legacy Archive remain separate source capabilities and separate request state.
Legacy Archive provides exact issued-ID compatibility lookup only.
Each use case decides whether Legacy state is required.

DRMCP does not introduce:

- a process-wide mutable Current Records index;
- a shared Current Records cache across requests;
- a package-specific index or snapshot;
- filesystem watchers;
- background refresh;
- incremental index patching;
- stale-snapshot reuse after source failure.

A mandatory current source failure prevents construction of a complete trustworthy Current Records snapshot.
The operation fails instead of returning partial normal data.

This ADR supersedes ADR-009.
It preserves request freshness, immutable request state, operation-specific Legacy loading, Current-versus-Legacy separation, and composition-owned lifecycle.

## Rationale

The portable package already follows current Spec semantics.
A shared snapshot avoids duplicate source-state and lookup models.

Explicit spec-tree root mapping supports the packaged physical layout without changing canonical identity rules.
The configured app namespace supplies `design_records`; path-relative Spec derivation supplies the suffix.

Request-scoped construction keeps package and host edits visible on the next request.
It also avoids cache invalidation and package-specific lifecycle contracts.

Legacy Archive stays separate because compatibility lookup has different identity and listing semantics.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Keep Guidance outside Current Records | Indexed package Specs would require duplicate parsing and identity behavior. |
| Build a separate package index | The package uses ordinary current Spec identities and lookup semantics. |
| Rewrite `product` refs at runtime | PRODUCT package generation already owns deterministic ref rewriting. |
| Require the package to mimic `<app>/records/spec/` physically | PRODUCT-REQ-SPEC-003 fixes the portable root at `<exe-dir>/design-records/`. |
| Merge Legacy Archive into the active index | Legacy issued-ID compatibility would leak into current listing and validation. |
| Reuse stale package state after failure | A normal response would conceal an incomplete current snapshot. |

## Consequences

- ADR-009 becomes historical and superseded.
- Guidance requests build and consume the normal Current Records snapshot.
- Current source contracts must support an explicit spec-tree root with an assigned app namespace.
- Package Specs become available to normal exact retrieval, resolution, and validation.
- Guidance list scope remains an Application-level restriction, not an index boundary.
- Performance refinements must preserve request freshness and one logical Current Records model.
- A package-specific lifecycle or request-spanning mutable index requires new architecture decision work.

Affected design areas:

- `spec:drmcp.application_architecture.runtime_and_state`;
- `spec:drmcp.application_architecture.failure_and_evolution`;
- `spec:drmcp.design_records_mcp.namespace_scanning`;
- `spec:drmcp.design_records_mcp.schema.discovery`;
- `spec:drmcp.design_records_mcp.schema.id_normalization`;
- `spec:drmcp.design_records_mcp.schema.record_source`.

## Evidence

- Source Requirements: `DRMCP-REQ-MCP-003` and `DRMCP-REQ-MCP-005`.
- Source Work Item: `DRMCP-WORK-MCP-016`.
- Review finding: F-BLK-01 in `DRMCP-TASK-MCP-016-09`.
- Revised decisions: D-019 and D-023 in `DRMCP-TASK-MCP-016-12`.
- ADR routing authority: B-06 in `DRMCP-TASK-MCP-016-13`.
- Preserved upstream authority: `DRMCP-ADR-MCP-001`.

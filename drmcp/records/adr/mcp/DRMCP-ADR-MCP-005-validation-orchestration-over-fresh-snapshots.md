# DRMCP-ADR-MCP-005: Validation orchestration over fresh snapshots

- **status**: accepted
- **date**: 2026-06-30
- **depends_on**:
  - DRMCP-ADR-MCP-002
  - DRMCP-ADR-MCP-003
  - DRMCP-ADR-MCP-004
- **supersedes**: []
- **migrated_to_spec**: 2026-06-30

## Context

`DRMCP-REQ-MCP-001` requires validation over current sources, active-index state, cross-namespace relations, and configured current-to-legacy relations.

Standalone `validate_records` must report source and graph findings without allowing detectors to observe different filesystem states.

Any future persisted-write integration must validate persisted bytes instead of candidate or pre-write state.
W011 does not define when the current authoring transaction invokes this architecture.
That integration requires the current-format realignment owned by `DRMCP-REQ-MCP-002`.

Validators that scan files independently would duplicate infrastructure behavior and could produce findings from inconsistent source generations.
Validators that call public MCP tools would couple core validation to transport and operation response contracts.

## Decision

Standalone `validate_records` builds one fresh request-scoped snapshot from the configured filesystem before validation passes begin.

Validation orchestration applies these stages:

1. enumerate and read configured sources;
2. parse sources and build current and configured legacy lookup state;
3. run per-source syntax, metadata, identity, and document-shape validation;
4. run relation and Topics graph validation against the complete index state;
5. aggregate, deduplicate, order, and project findings into the operation output.

Individual validators:

- perform no filesystem I/O;
- do not rescan roots;
- do not call public MCP tools;
- do not call one another through transport boundaries;
- consume retained source, parsed, index, or finding data;
- return findings as data;
- do not format MCP responses.

Legacy archive records remain lookup targets where accepted contracts require them.
They do not become normal current-repository validation subjects.

Diagnostic aggregation, stable ordering, semantic duplicate suppression, and MCP response formatting occur outside individual validators.

Any future caller that validates after persisted filesystem writes must discard candidate and pre-write state.
The caller must rebuild validation input from persisted files before using this validation architecture.

Current authoring-transaction timing, source loading, and current-format validation integration are deferred to `DRMCP-REQ-MCP-002`.
This ADR does not integrate YAML or V01-SPEC authoring semantics with current-format validation.

## Rationale

One snapshot gives every validation pass a consistent view of sources and relations.
Per-source validation must precede graph validation because graph checks require complete identity and retained invalid-source state.

Pure validators are easier to test and reuse.
Keeping filesystem I/O and MCP formatting outside detectors preserves the architecture boundaries in `DRMCP-ADR-MCP-003` and `DRMCP-ADR-MCP-004`.

A persisted-state rebuild validates actual written state.
Reusing candidate or pre-write state would validate intended content instead of repository state.
This rule remains a constraint for future authoring integration.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Let every validator scan the filesystem | Passes could observe different states and duplicate source-loading policy. |
| Invoke validators through public tools | Core validation would depend on MCP transport and operation wrappers. |
| Run graph validation before per-source state is complete | Relation and Topics checks would operate on incomplete or ambiguous identity state. |
| Format diagnostics inside each validator | Ordering, deduplication, and operation projection would fragment across detectors. |
| Reuse the pre-write snapshot for post-write validation | Validation would not prove the bytes actually persisted. |
| Validate legacy archive records as normal subjects | Archive records would leak into current repository validation. |

## Consequences

- Standalone `validate_records` depends on the W011 fresh-snapshot lifecycle.
- Detectors require explicit input and finding types.
- Relation and Topics graph validators run only after the complete index state exists.
- A future persisted-write integration incurs a new source enumeration and index build.
- The current authoring transaction is not required by this ADR to perform that rebuild until `DRMCP-REQ-MCP-002` defines the integration.
- Stable ordering and duplicate suppression remain application or shared projection responsibilities.
- A snapshot construction failure prevents a trustworthy normal validation result.

Affected Specification targets:

- `spec:drmcp.design_records_mcp.tools.validate_records`;
- `spec:drmcp.design_records_mcp.schema.diagnostics`;
- `spec:drmcp.design_records_mcp.schema.record_model`.

Deferred integration target:

- `DRMCP-REQ-MCP-002` for current-format authoring-transaction realignment.

## Evidence

- Source Requirement: `DRMCP-REQ-MCP-001`.
- Source Work Item: `DRMCP-WORK-MCP-011`.
- Accepted decision: D-006 in `DRMCP-TASK-MCP-011-01`.
- Runtime snapshot authority: `DRMCP-ADR-MCP-002`.
- Component and contract authorities: `DRMCP-ADR-MCP-003` and `DRMCP-ADR-MCP-004`.
- Retained detector ownership: `DRMCP-WORK-MCP-007`, `DRMCP-WORK-SPEC-001`, and `DRMCP-WORK-SPEC-002`.
- Deferred current authoring-transaction integration: `DRMCP-REQ-MCP-002`.

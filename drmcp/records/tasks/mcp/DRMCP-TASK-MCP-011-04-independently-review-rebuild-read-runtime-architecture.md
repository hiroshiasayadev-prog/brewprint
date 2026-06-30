# DRMCP-TASK-MCP-011-04: Independently review rebuild read-runtime architecture

- **id**: DRMCP-TASK-MCP-011-04
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-011
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-011-03
- **outputs**:
  - DRMCP-TASK-MCP-011-04
  - DRMCP-WORK-MCP-011

## Goal

Persist the completed independent design-review result for the rebuild read-runtime architecture.

Make the review verdict, findings, implementation-planning readiness, and exact next gate canonical without performing correction work.

## Work

- Record the accepted independent reviewer result without rerunning the review.
- Confirm that the accepted result does not fundamentally contradict the named decision, ADR, and Specification artifacts.
- Record the verdict and finding inventory.
- Preserve the exact observed problem and required correction outcome for F-MAJ-01 and F-MAJ-02.
- Record whether either finding requires a new user decision.
- Synchronize Phase D and review Evidence into `DRMCP-WORK-MCP-011`.
- Stop before finding correction, finding-closure re-review, design closure, implementation planning, or production implementation.

## Done condition

- The verdict is recorded as `NEEDS REVISION`.
- Blocking findings are recorded as none.
- Major findings are recorded as F-MAJ-01 and F-MAJ-02.
- Minor findings and advisories are recorded as none.
- Both findings preserve `New user decision required: no`.
- Implementation-planning readiness is recorded as blocked pending finding correction.
- The exact correction, re-review, and closure sequence is recorded.
- W011 lists this Task and records Phase D as `done: NEEDS REVISION`.
- W011 remains `in_progress`.
- T05 is not created or added to W011.
- No ADR, Specification, Requirement, production source, or implementation Task changes.

## Verification

- Compare the accepted reviewer result with T01 decision summaries, ADR-002 through ADR-006, and the named Specifications.
- Confirm that F-MAJ-01 identifies the unresolved operation-set boundary without inventing a new design decision.
- Confirm that F-MAJ-02 identifies the authoring-contract and current-format validation contradiction without changing D-006.
- Confirm that the Work Item Task list contains T01 through T04 only.
- Confirm that Phase D records `done: NEEDS REVISION` and Phase E remains conditional.
- Inspect Git state, textual diff, and whitespace only for this Task and W011.
- Do not inspect or infer repository-wide cleanliness.

## Evidence

### Review scope and authority check

The completed independent review result was accepted as the input to this synchronization Task.
The review was not rerun.

The result was checked for fundamental contradiction against:

- `DRMCP-TASK-MCP-011-01` through `DRMCP-TASK-MCP-011-03`;
- `DRMCP-ADR-MCP-001` through `DRMCP-ADR-MCP-006`;
- `spec:drmcp.implementation`;
- the named Design Records MCP responsibility, tools, authoring, and schema contracts.

No fundamental contradiction prevents persistence of the accepted review result.

### Verdict

| item | result |
|---|---|
| Verdict | `NEEDS REVISION` |
| Blocking findings | None |
| Major findings | F-MAJ-01, F-MAJ-02 |
| Minor findings | None |
| Advisories | None |
| Implementation-planning readiness | `NOT READY — FINDING CORRECTION REQUIRED` |

Exact next gate:

```text
F-MAJ-01 and F-MAJ-02 correction
  -> independent finding-closure re-review
  -> design closure synchronization
```

### F-MAJ-01

| field | value |
|---|---|
| Finding ID | F-MAJ-01 |
| Severity | major |
| Affected decisions | D-001, D-002, D-007, D-009 |
| New user decision required | no |

Observed problem:

- The decision ledger and ADRs use `each MCP tool invocation` or `each public MCP tool dispatch`.
- `spec:drmcp.implementation` defines application packages for only four operations:
  - `list_records`;
  - `get_records`;
  - `resolve_reference`;
  - `validate_records`.
- The tools catalog also contains authoring guidance and authoring transaction tools.
- W011 excludes the authoring transaction runtime from scope.
- The architecture does not uniquely identify whether it covers four read-runtime operations or all MCP tools.
- Authoring-guidance snapshot lifecycle and package ownership remain undecided.
- The package tree does not state whether it is a complete inventory or a partial example.
- The ambiguity leaves an architecture decision to implementation Task authoring.

Required correction outcome:

- Make the W011 architecture operation set canonical and unambiguous.
- Limit fresh snapshots, dedicated use cases, and package layout to these four read-runtime operations under W011 and `DRMCP-REQ-MCP-001`:
  - `list_records`;
  - `get_records`;
  - `resolve_reference`;
  - `validate_records`.
- Align these artifacts to the same boundary:
  - T01 decision summaries;
  - ADR-002;
  - ADR-003;
  - `spec:drmcp.implementation`;
  - responsibility boundary;
  - tools overview;
  - package responsibility table.
- Do not decide authoring-guidance or authoring-transaction architecture in this Work Item.

### F-MAJ-02

| field | value |
|---|---|
| Finding ID | F-MAJ-02 |
| Severity | major |
| Affected decisions | D-006 |
| New user decision required | no |

Observed problem:

- Phase C added a fresh post-write validation rule to authoring transaction contracts.
- The existing authoring transaction contract assumes YAML front matter, V01-SPEC identity, and `design_record` fields.
- Current read and validation authority rejects YAML front matter as a current metadata source and uses path-derived `spec:` identity.
- Persisted output produced by the authoring contract can therefore be rejected by post-write validation in the same response.
- W011 excludes authoring transaction runtime and assigns it to `DRMCP-REQ-MCP-002`.
- Phase C normatively connected unresolved legacy authoring semantics to the current read architecture.

Required correction outcome:

- Preserve D-006's decision to rebuild a fresh snapshot from persisted filesystem state.
- Define only the read-runtime side of post-write validation architecture in W011.
- Defer application to authoring transaction contracts until `DRMCP-REQ-MCP-002` completes current-format realignment.
- Remove or explicitly defer the normative authoring transaction Specification synchronization added by W011 Phase C.
- Do not treat YAML/V01-SPEC authoring and current-format validation as one simultaneously valid integrated contract.

### Review conclusion

Implementation planning is not ready.
A dedicated correction Task must own only F-MAJ-01 and F-MAJ-02 and their direct consistency effects.
The correction author must not close the findings.

# Reference: Artifact boundary

- **id**: `spec:product.design_records.authoring_standards.artifact_boundary`
- **status**: draft
- **date**: 2026-07-01
- **parent**: `spec:product.design_records.authoring_standards`

## What this is

Lightweight authoring-time projection of the artifact responsibility boundary for DRMCP-managed design records.

Use this spec to select the correct artifact kind before writing. Canonical artifact ownership is defined by `spec:product.design_records.artifact_model.artifact_responsibility_matrix`. When wording conflicts between this spec and the canonical source, the canonical source wins.

## Covered artifact kinds

| artifact | authoring-time purpose |
|---|---|
| ADR | Record an adopted design decision and its rationale. |
| spec | Record the currently valid specification or contract. |
| investigation | Record research, evidence, uncertainty, and options before a decision. |
| requirement | Record a stable need, gap, or requested outcome. |
| work item | Record a bounded resolution flow and Task graph for direct material sources. |
| task | Record concrete short-term work and its verification. |

Metadata, file shape, status lifecycle, and authoring interface rules for each kind are defined in the corresponding per-artifact authoring guide under `spec:product.design_records.authoring_standards`.

## Not covered

These artifact families are outside the scope of this spec:

- BPDSL YAML
- Render artifacts
- Internal design
- Target implementation

Ownership for the complete artifact system, including non-DRMCP artifacts, is defined by the canonical responsibility matrix.

## Distinguishing adjacent artifacts

Use the following decision rules when the correct artifact kind is unclear.

| question | artifact |
|---|---|
| Does an adopted decision require durable rationale or supersession history? | ADR |
| Is a decision still under exploration? | investigation |
| Is a decision answer being checkpointed for resumable workflow state? | task |
| Is the currently valid contract or rule being stated? | spec |
| Has a stable need or gap been identified? | requirement |
| Is a bounded resolution flow and Task graph being tracked for direct material sources? | work item |
| Is short-term concrete work being tracked with completion conditions? | task |

Rules:

- Exploration logs and research belong in an investigation, not an ADR.
- Currently valid contracts belong in a spec, not an ADR.
- Progress and implementation steps belong in Work Items or Tasks, not an ADR.
- Durable decision history belongs in ADRs, not a Specification.
- A Task records concrete work closeable in the short term.
- A Work Item records a bounded resolution flow, completion boundary, and Task graph for its direct material sources.
- A Work Item created or decomposed from a Task records that exact Task as direct provenance.
- Task-originated decomposition does not create implicit parent or child Work Item metadata.
- A decision occurrence does not automatically require an ADR.

## Decision workflow projection

The artifact boundary separates workflow checkpoints from canonical design state.

| stage | artifact responsibility |
|---|---|
| Decision Task checkpoint | Persist selected option, concise reason, canonical target, routing state, and workflow cursor. |
| ADR routing | Decide whether durable rationale, alternatives, consequences, or supersession history require an ADR. |
| ADR authoring when required | Record the durable choice and rationale in a separate authoring Task. |
| Direct Specification synchronization | Write decisions without an ADR requirement into the current normative Specification. |
| Specification synchronization after ADR | Write the accepted ADR outcome into the current normative Specification. |

Required flow:

```text
decision Task checkpoint
  -> ADR routing
  -> ADR authoring when required
  -> direct Specification synchronization when ADR is not required
```

A Task checkpoint is temporary and historical workflow state.
An ADR is durable decision rationale.
A Specification is the current normative contract.
A Task ledger must not replace either canonical artifact.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.authoring_standards` | Parent index. |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | Canonical artifact ownership. |
| PRODUCT-REQ-SPEC-005 | Typed single-responsibility Task contract. |
| PRODUCT-REQ-SPEC-006 | Generic workflow source-relation requirement. |
| PRODUCT-ADR-SPEC-006 | Decision checkpoints and canonical design-state boundary. |
| PRODUCT-ADR-SPEC-007 | Work Item provenance and Task-originated decomposition. |
| PRODUCT-ADR-SPEC-008 | Legacy workflow relation migration boundary. |

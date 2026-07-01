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

The artifact boundary separates workflow checkpoints, ADR routing, canonical design state, and closure propagation.

| stage | artifact responsibility |
|---|---|
| Decision Task checkpoint | Own selected option, concise reason, responsibility boundary, scope, canonical target, routing state, and workflow cursor. |
| ADR routing | Classify each decision as `required`, `covered`, `not_required`, or `blocked`. |
| ADR boundary partitioning | Group coherent decisions and prevent both omnibus and mechanically fragmented ADR boundaries. |
| ADR disposition | Select `create`, `amend`, `reuse`, or `supersede` and map exact decision IDs to canonical targets. |
| ADR authoring when required | Record the durable choice and rationale in a separate authoring Task. |
| Specification synchronization | Write the accepted current normative rule into the canonical Specification. |
| Integrated review | Own the independent verdict and finding set for the final combined state. |
| Closure synchronization | Propagate only exact mechanically derivable lifecycle, Evidence, and relation state. |

ADR routing and ADR authoring are separate responsibilities.
Routing does not write ADR body content.
Authoring consumes the completed route.

A decision Task owns the workflow checkpoint but does not own canonical design state.
A completed decision Task retains its checkpoint unchanged.
Downstream authoring, review, and synchronization Tasks own their own references and Evidence.
Downstream progress must not be written back into the completed decision Task.

An ADR owns durable decision rationale.
A Specification owns the current normative contract.
A Task ledger must not replace either canonical artifact.

Closure synchronization must not author or correct canonical content.
Closure synchronization must not create or amend the Task graph.
Closure stops when missing work, graph change, or unresolved judgment is discovered.

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
| PRODUCT-ADR-SPEC-014 | Completed-record preservation and closure-synchronization boundary. |

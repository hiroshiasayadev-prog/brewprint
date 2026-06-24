# Reference: Artifact boundary

- **id**: `spec:product.design_records.authoring_standards.artifact_boundary`
- **status**: draft
- **date**: 2026-06-23
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
| work item | Record the complete resolution flow for a requirement. |
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
| Has a decision been made and needs to be recorded? | ADR |
| Is a decision still under exploration? | investigation |
| Is the currently valid contract or rule being stated? | spec |
| Has a stable need or gap been identified? | requirement |
| Is the complete resolution flow for a requirement being tracked? | work item |
| Is short-term concrete work being tracked with completion conditions? | task |

Rules:

- Exploration logs and research belong in an investigation, not an ADR.
- Currently valid contracts belong in a spec, not an ADR.
- Progress and implementation steps belong in work items or tasks, not an ADR.
- Decision history and research belong in ADRs and investigations, not a spec.
- A task records concrete work closeable in the short term; a work item records the complete resolution flow and task graph for a requirement.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.authoring_standards` | Parent index. |
| `spec:product.design_records.artifact_model.artifact_responsibility_matrix` | Canonical artifact ownership. |

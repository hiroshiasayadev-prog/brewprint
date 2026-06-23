# PRODUCT-TASK-SPEC-011-02: Create artifact boundary and common guide shape

- **id**: PRODUCT-TASK-SPEC-011-02
- **status**: done
- **date**: 2026-06-23
- **work_item**: PRODUCT-WORK-SPEC-011
- **source_requirement**: PRODUCT-REQ-SPEC-002
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-011-01
- **outputs**:
  - `product/records/spec/concepts/authoring-standards/artifact-boundary.md`
  - `product/records/spec/concepts/authoring-standards/index.md`

## Goal

Create the lightweight authoring-time artifact boundary. The common per-artifact guide structure is defined normatively in PRODUCT-WORK-SPEC-011 and enforced through individual per-artifact guides; it is not published as a separate spec artifact in this task.

## Work

- Create `spec:product.concepts.authoring_standards.artifact_boundary` as a Reference spec.
- Cover ADR, spec, investigation, requirement, work item, and task.
- Define the minimum artifact-selection guidance needed before authoring.
- Cite the project artifact responsibility matrix as the canonical ownership source.
- State that the canonical matrix wins when wording conflicts.
- Exclude BPDSL YAML, render artifacts, internal design, and target implementation.
- Keep metadata, file-shape, and lifecycle rules in per-artifact guides.
- Update the authoring-standards Index.
- Correct the ADR authoring Index summary to match its current scope.

## Done condition

| item | done when |
|---|---|
| Boundary published | `artifact-boundary.md` exists with the correct semantic ref and parent. |
| Six artifact kinds covered | ADR, spec, investigation, requirement, work item, and task are distinguished. |
| Canonical authority preserved | The spec cites the responsibility matrix and defines conflict precedence. |
| Scope remains lightweight | The spec does not duplicate per-artifact metadata, file-shape, or lifecycle rules. |
| Index updated | The boundary is listed and the ADR summary is current. |

## Verification

- Confirm the boundary is a Reference spec.
- Confirm the responsibility matrix is cited as canonical.
- Confirm all six DRMCP-managed design record kinds are covered.
- Confirm excluded artifact families are explicit.
- Confirm the Index uses canonical refs and English summaries.

## Evidence

- `product/records/spec/concepts/authoring-standards/artifact-boundary.md` created as Reference spec with id `spec:product.concepts.authoring_standards.artifact_boundary`.
- All six DRMCP-managed artifact kinds covered: ADR, spec, investigation, requirement, work item, task.
- Canonical responsibility matrix cited; conflict precedence stated (canonical source wins).
- Non-DRMCP artifact families (BPDSL YAML, render artifacts, internal design, target implementation) explicitly excluded.
- Per-artifact metadata, file shape, and lifecycle rules are deferred to per-artifact guides; not duplicated.
- `product/records/spec/concepts/authoring-standards/index.md` updated: artifact boundary added to Topics table; ADR authoring summary corrected to match current guide scope (removed "status enum", "responsibility boundary", "DRMCP authoring transaction"; replaced with "status lifecycle", "kind-specific writing rules", "authoring interface requirements").
- Common per-artifact guide structure: defined normatively in PRODUCT-WORK-SPEC-011, not published as a separate Reference spec. Task Goal updated to reflect this scope decision.

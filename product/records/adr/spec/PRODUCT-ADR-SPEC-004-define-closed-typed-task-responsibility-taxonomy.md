# PRODUCT-ADR-SPEC-004: Define the closed typed Task responsibility taxonomy

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**: []
- **supersedes**: []
- **migrated_to_spec**: null

## Context

`PRODUCT-REQ-SPEC-005` requires every Task to declare one primary Task type and own one matching responsibility.

The current Task authoring contract has no persisted type field or closed semantic taxonomy.
A prose-only Task description cannot reliably prevent mixed responsibilities.
A free-form label would also prevent consistent authoring and downstream validation.

This decision defines the durable Task-type taxonomy.
It does not define DRMCP parser, validation, diagnostic, index, or projection behavior.
It does not migrate existing Task records.

## Decision

Use the required scalar metadata field `task_type`.

Permit exactly one value from this closed set:

- `investigation`
- `decision`
- `authoring`
- `implementation`
- `review`
- `correction`
- `verification`
- `coordination`
- `synchronization`

Each value owns one primary outcome and one completion judgment.

| task type | owned outcome | completion judgment | prohibited primary overlaps |
|---|---|---|---|
| `investigation` | One Investigation record for one bounded research question. | The Investigation record satisfies its completion requirements. | Decision adoption, canonical authoring, implementation, independent review, finding correction, or lifecycle synchronization. |
| `decision` | One bounded decision ledger. | Every owned item is `decided`, `deferred`, or validly `blocked`. | Investigation authoring, canonical ADR or Specification authoring, implementation, independent review, correction, or final synchronization. |
| `authoring` | One bounded canonical artifact set from decided inputs. | The artifact set satisfies its authoring requirements and the Task Done condition. | Unresolved decision work, implementation, independent review, finding correction, or lifecycle synchronization. |
| `implementation` | One bounded implementation outcome. | The implementation contract and declared verification pass. | Unresolved design decisions, canonical design authoring, independent review, correction ownership, or workflow coordination. |
| `review` | One bounded independent verdict and finding set. | The result is `PASS` or `NEEDS REVISION` with complete finding evidence. | Authoring, implementation, finding correction, finding closure by the correction author, or lifecycle synchronization. |
| `correction` | One bounded named finding set repaired. | The named repairs and direct verification pass. | Independent finding closure, unrelated improvement, new decision adoption, or lifecycle synchronization. |
| `verification` | One bounded objective acceptance gate. | Every predefined check is executed and the overall result is `PASS`, `FAIL`, or validly `BLOCKED`. | Artifact modification, undefined semantic judgment, repair, independent review verdict, or lifecycle synchronization. |
| `coordination` | One parent Work Item overview of child Work Items and responsibility boundaries. | Required child Work Items exist with distinguishable, non-overlapping responsibilities. | Child-owned investigation, decision, authoring, implementation, review, correction, or verification deliverables. |
| `synchronization` | One bounded accepted-state propagation. | All specified lifecycle, Evidence, completion-result, and relation state expresses the same accepted result. | New design judgment, decomposition, substantive deliverable creation, implementation, review, or correction. |

Task authoring must persist `task_type` on create.
Task authoring may change `task_type` through an update.

PRODUCT owns the field meaning and closed value set.
DRMCP owns the later parser, validation, diagnostic, indexing, and tool-projection contracts.

## Rationale

A closed taxonomy makes Task responsibility explicit before execution begins.

Each value has a distinct outcome and completion authority.
The taxonomy therefore supports writer ownership, review independence, and reliable dependency graphs.

A required scalar field avoids ambiguous multi-label Tasks.
The field also gives future validators a stable input without assigning validation behavior to PRODUCT.

The nine values cover the complete accepted design workflow and implementation workflow without introducing a generic catch-all type.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Infer the Task type from Goal or Work prose. | Inference is ambiguous and cannot support deterministic validation. |
| Allow free-form type values. | Free-form values do not establish stable responsibility semantics. |
| Permit several primary types on one Task. | Multiple primary types allow multiple outcomes and completion judgments. |
| Keep one generic Task class. | A generic class does not expose writer, verifier, or completion ownership. |
| Define labels without type-specific contracts. | Labels alone do not prevent mixed responsibility. |
| Include model routing or validator types. | Model routing and validator design are outside REQ-005. |

## Consequences

- `spec:product.design_records.authoring_standards.task_authoring` must add `task_type` to its metadata and authoring interface.
- The Task authoring Specification must define the nine type contracts.
- `spec:product.design_records.authoring_standards.work_item_authoring` must preserve the coordination boundary where needed.
- Downstream DRMCP work must consume the accepted field and closed value set.
- Existing Task migration remains separate future work.
- The current W016 and W017 bootstrap Tasks remain under the pre-REQ-005 metadata contract.

## Evidence

- `PRODUCT-REQ-SPEC-005`: accepted typed single-responsibility Task requirement.
- `PRODUCT-TASK-SPEC-016-02`: D-001 through D-011.
- `PRODUCT-TASK-SPEC-016-04`: C-001 and C-002 conflict dispositions.
- `PRODUCT-TASK-SPEC-016-05`: ADR routing and this ADR boundary.
- `V01-ADR-091`: compatible historical Work Item and Task separation context.

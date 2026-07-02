# PRODUCT-ADR-SPEC-015: Define Task-local semantic responsibility validation semantics

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**:
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
- **supersedes**: []
- **migrated_to_spec**: null

## Context

`PRODUCT-ADR-SPEC-004` defines the closed Task-type taxonomy.
`PRODUCT-ADR-SPEC-005` requires one primary outcome and one completion judgment per Task.

`PRODUCT-REQ-SPEC-007` requires semantic validation of that responsibility boundary.
The validator must report both overall compliance and criterion-level judgments.

A free-form model verdict would not provide deterministic aggregation.
External-context inference would also exceed the Task-local compliance claim.
Structural failure and execution failure must not be reported as semantic non-compliance.

## Decision

Evaluate exactly one Task record per validation invocation.
Use the Task record as the only semantic Evidence source.
Do not infer supporting context from other records or repository state.

Select the validation checklist automatically from the declared `task_type`.
Compose the applied checklist from:

- one common responsibility-boundary criterion set;
- one criterion set for the declared `task_type`.

Evaluate every combined criterion independently.
Return one binary result and one concise Task-local reason for every criterion.
Require reasons for both compliant and non-compliant results.

Treat missing required Task-local content as a non-compliant criterion.
Do not replace missing content with external inference.

Derive overall compliance through logical AND across all criterion results.
Overall compliance is true only when every evaluated criterion is true.
The semantic evaluator does not produce a separate free-form overall judgment.

Do not start semantic evaluation when:

- the Task cannot be read or parsed;
- `task_type` is missing;
- `task_type` is invalid.

Report those conditions as structural precondition failures.
Do not emit semantic criterion results or overall compliance for a precondition failure.

Keep model, runtime, timeout, response-decode, and incomplete-response failures separate from semantic judgment.
Do not emit semantic criterion results or overall compliance for an execution failure.
Do not synthesize omitted criterion judgments.

## Rationale

Task-local Evidence keeps the compliance claim bounded to the evaluated Task.
Automatic selection prevents callers from changing the applicable responsibility contract.

Common and type-specific criteria preserve shared rules without duplicating every checklist.
Independent criterion judgments make non-compliance reviewable.

Logical-AND aggregation makes the overall result deterministic.
The rule also preserves the meaning that every criterion is required.

Separating precondition, execution, and semantic outcomes prevents false claims about Task quality.
A semantic false result means evaluation completed and found a responsibility-boundary violation.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Infer missing Evidence from external records or repository context. | The validator would exceed the Task-local compliance claim. |
| Let the caller select arbitrary criteria. | Callers could bypass the contract associated with the declared `task_type`. |
| Ask the model for one free-form overall verdict. | The overall result would not follow deterministic criterion aggregation. |
| Treat structural precondition failure as semantic false. | The validator could not establish that semantic evaluation occurred. |
| Treat execution failure as semantic false. | Runtime failure does not establish Task non-compliance. |

## Consequences

- `spec:product.responsibility_boundary_validator` must define the current normative result and failure semantics.
- PRODUCT-WORK-SPEC-020 must author the exact common and type-specific checklist content.
- PRODUCT-WORK-SPEC-021 must preserve deterministic checklist selection and aggregation.
- Checklist wording, criterion identifiers, and physical storage remain outside this ADR.
- External response field names remain outside this ADR.
- Model, provider, runtime, timeout, retry, and implementation choices remain outside this ADR.

## Evidence

- `PRODUCT-TASK-SPEC-019-01`: D-001 through D-005, D-007, D-009, and D-010.
- `PRODUCT-TASK-SPEC-019-03`: preserved Task-local semantic decisions under the corrected standalone boundary.
- `PRODUCT-TASK-SPEC-019-07`: R-001 and R-002 canonical ownership inputs.
- `PRODUCT-TASK-SPEC-019-12`: B-001 routing boundary and `create` disposition.

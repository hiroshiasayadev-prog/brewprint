# PRODUCT-REQ-SPEC-007: Semantic Task Responsibility Boundary Validation

- **id**: PRODUCT-REQ-SPEC-007
- **status**: accepted
- **date**: 2026-07-01
- **source_refs**:
  - PRODUCT-REQ-SPEC-005
  - spec:product.design_records.authoring_standards.task_authoring
- **work_items**: []

## Requirement

Task authoring workflows must provide semantic validation of Task responsibility boundaries.

The validator must determine whether one Task owns one responsibility and one completion judgment consistent with its declared Task type.

The validator must base the judgment only on information contained in the Task record.

## Evidence

- PRODUCT-REQ-SPEC-005 requires every Task to declare one primary Task type and own one matching responsibility.
- Structural validation cannot determine whether Task sections combine semantically independent responsibilities.
- Mixed responsibilities may remain hidden when each required Task section is structurally present.
- A repeatable semantic check is needed before a Task is accepted for downstream execution.
- An overall binary result does not identify which responsibility-boundary criteria passed or failed.
- Each criterion judgment needs Task-local supporting evidence or a concise reason for review and correction.

## Required Outcome

- Validate the semantic responsibility boundary of one Task record.
- Evaluate whether the Task responsibility matches its declared Task type.
- Detect multiple independent responsibilities within one Task.
- Detect multiple independent completion judgments within one Task.
- Treat missing Task-local evidence as non-compliance instead of inferring external context.
- Keep validation read-only.
- Return a machine-readable overall compliance result.
- Return one binary compliance judgment for each evaluated responsibility-boundary criterion.
- Return Task-local supporting evidence or a concise reason for each criterion judgment.
- Support use during Task authoring or before Task release.
- Limit the compliance claim to Task responsibility boundaries.
- Do not claim complete executor readiness.

## Explicitly Excluded Scope

- Concrete language model, provider, or runtime selection.
- Concrete MCP tool request, response, or error schemas.
- Checklist item wording or checklist storage format.
- Retry policy and error taxonomy.
- Complete executor-readiness validation.
- Multi-Task or Work Item graph validation.
- Authority or Specification consistency validation.
- Automatic Task correction or rewriting.
- Production implementation scheduling.

## Boundary

PRODUCT owns the need for semantic Task responsibility-boundary validation and the scope of its compliance claim.

Subsequent ADRs and Specifications own validator criteria, checklist contracts, tool contracts, result schemas, failure handling, and implementation choices.

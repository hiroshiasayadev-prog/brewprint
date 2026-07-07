# PRODUCT-ADR-SPEC-017: Validate Task responsibility after authoring and final Evidence with human-owned exceptions

- **status**: accepted
- **date**: 2026-07-03
- **depends_on**:
  - PRODUCT-ADR-SPEC-009
  - PRODUCT-ADR-SPEC-015
  - PRODUCT-ADR-SPEC-016
- **supersedes**: []
- **migrated_to_spec**: null

## Context

A Task responsibility boundary can be defective before execution begins.
The completed Task can also expose a boundary violation through its final Evidence and outcome.

One validation point cannot cover both states.
The validator must remain a semantic evaluator rather than a workflow gate owner.

Semantic violations also require an explicit exception owner.
Automatic acceptance would hide responsibility-boundary deviations.
Automatic rejection would prevent justified exceptions.

## Decision

Run the same semantic Task responsibility validator at two points:

- immediately after Task authoring;
- after final Evidence is written and before the Task becomes `done`.

Do not run the second invocation as a prerequisite for `cancelled`.
A cancelled Task has no completed outcome to validate and follows its separate terminal section-readiness and Evidence contract.

Use the same validator and criterion-level result contract at both points.
The validator reports the result defined by PRODUCT-ADR-SPEC-015.
The validator does not own workflow enforcement.

The authoring workflow owns the post-authoring call.
The successful Task completion or release workflow owns the post-Evidence call.

A fully compliant semantic result permits the caller to continue its normal route.
Any semantic violation must route to explicit human acceptance or rejection.

When a human accepts a violation, preserve these items in Task Evidence:

- each violated criterion;
- the validator rationale for each violation;
- the acceptance decision;
- the acceptance reason.

Human acceptance records an exception for that Task.
Human acceptance does not change the validator result or the underlying criterion contract.

When a human rejects a violation, route the work to:

- Task correction; or
- responsibility-boundary reconsideration.

The caller selects the applicable return route from the accepted workflow state.
The validator does not correct, rewrite, split, release, or reject the Task automatically.

## Rationale

Post-authoring validation detects mixed responsibility before execution consumes the Task.
Post-Evidence validation checks the completed record against its actual outcome and Evidence.

Using one validator contract prevents invocation-specific meanings from diverging.
Keeping enforcement outside the validator preserves its read-only semantic boundary.

Human judgment permits explicit exceptions without weakening the default responsibility contract.
Recorded exception Evidence keeps accepted deviations reviewable.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Validate only after Task authoring. | Final Evidence and actual outcome would not be evaluated. |
| Validate only after final Evidence. | Invalid responsibility boundaries could reach execution unchecked. |
| Require the post-Evidence invocation before cancellation. | Cancellation may occur before substantive Work or Verification exists and may propagate across several Tasks without a completed outcome to evaluate. |
| Use separate result contracts at each point. | The same Task responsibility rule would have different meanings by invocation context. |
| Let the validator enforce continuation or release. | Semantic evaluation and workflow enforcement have different owners. |
| Accept semantic violations automatically. | Exceptions would lack explicit human ownership and rationale. |
| Reject every semantic violation automatically. | The workflow could not preserve justified Task-specific exceptions. |
| Record acceptance without violated criteria and reasons. | Reviewers could not determine what exception was accepted. |

## Consequences

- PRODUCT-REQ-SPEC-007 remains unchanged and continues to own the stable semantic-validation need rather than these downstream workflow-policy decisions.
- `spec:product.responsibility_boundary_validator` must define the shared result contract and exception boundary.
- `spec:product.design_records.authoring_standards.task_authoring` must define the narrow post-authoring usage rule.
- Successful Task completion or release workflows must invoke the validator after final Evidence and before `done`.
- Cancellation does not invoke the post-Evidence validator.
- Accepted violations require explicit Task Evidence.
- Rejected violations return to correction or responsibility-boundary reconsideration.
- UI, CLI, transport, identity, notification, retry, and implementation details remain outside this ADR.
- Cancelled-state structural and Evidence readiness remains owned by Task authoring authority rather than the semantic responsibility validator.

## Evidence

- `PRODUCT-TASK-SPEC-019-01`: D-012 shared validator and caller-owned enforcement boundary.
- `PRODUCT-TASK-SPEC-019-07`: R-005 two-point invocation and human exception route.
- `PRODUCT-TASK-SPEC-019-07`: validation and human-acceptance portions of R-006.
- `PRODUCT-TASK-SPEC-019-12`: B-003 routing boundary and `create` disposition.
- `PRODUCT-ADR-SPEC-009`: design convergence and production implementation completion boundary.
- `PRODUCT-TASK-SPEC-023-04`: J-004 successful-completion-only post-Evidence invocation.
- `PRODUCT-TASK-SPEC-023-06`: B-004 non-material amendment route.

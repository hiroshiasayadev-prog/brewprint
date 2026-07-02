# Overview: Responsibility-boundary validator

- **id**: `spec:product.responsibility_boundary_validator`
- **status**: draft
- **date**: 2026-07-02
- **parent**: `spec:product`

## What this is

PRODUCT-owned semantic contract for Task responsibility-boundary validation.

TRV consumes this contract for app-local design and later implementation.
Current DRMCP does not own or implement this semantic contract.

## Current contract

| concern | contract |
|---|---|
| Semantic contract owner | PRODUCT. |
| Application owner | TRV. |
| Invocation scope | Evaluate exactly one Task record per invocation. |
| Semantic Evidence | Use only content in the evaluated Task record. Do not infer missing support from other records or repository state. |
| Checklist selection | Select the applicable checklist automatically from the declared `task_type`. |
| Checklist composition | Apply the common responsibility-boundary criteria and the criteria for the declared `task_type`. |
| Criterion result | Return one binary result and one concise Task-local reason for every evaluated criterion. |
| Evidence locator | Include a Task section reference for each criterion reason. An excerpt is optional. Line numbers are not required. |
| Overall compliance | Derive the overall result through logical AND across all criterion results. |
| Result identity | Do not require checklist revision identifiers or stable criterion identifiers in the external result contract. |
| Outcome classes | Keep structural precondition failure, completed semantic evaluation, and execution failure distinct. |
| Workflow use | Use the same validator contract after Task authoring and after final Evidence is written. |
| Violation disposition | Route semantic violations to explicit human acceptance or rejection. |

## Non-goals

- Exact common or Task-type-specific checklist wording.
- Checklist storage format or placement.
- Concrete response field names.
- TRV language model, provider, runtime, timeout, retry, or decode policy.
- TRV UI, CLI, transport, identity, packaging, deployment, or notification mechanisms.
- Current DRMCP integration or current DRMCP contract changes.
- Complete executor-readiness validation.
- Multi-Task or Work Item graph validation.
- Automatic Task correction, rewriting, splitting, release, or rejection.

## Rules

### Evaluation boundary

- The input is one readable Task record.
- The Task record is the only semantic Evidence source.
- Missing required Task-local content produces a non-compliant criterion result.
- External context must not substitute for missing Task-local content.

### Checklist selection and composition

- The validator selects the checklist from the declared `task_type`.
- The caller must not select or omit individual responsibility criteria.
- The applied checklist combines:
  - the common responsibility-boundary criterion set;
  - the criterion set for the declared `task_type`.
- PRODUCT-WORK-SPEC-020 owns exact checklist wording, storage format, placement, and checklist review.

### Criterion results and aggregation

- Evaluate every applied criterion independently.
- Return one binary result for every criterion.
- Return one concise reason for both compliant and non-compliant results.
- Base every reason on the evaluated Task.
- Include a Task section reference with each reason.
- An excerpt may supplement a section reference.
- Do not require line-number references.
- Derive overall compliance through logical AND.
- Overall compliance is true only when every criterion result is true.
- The semantic evaluator must not produce a separate free-form overall verdict.
- The external result contract does not require checklist revision identifiers or stable criterion identifiers.

### Outcome separation

| outcome | meaning | semantic results allowed |
|---|---|---|
| Structural precondition failure | The Task cannot be read or parsed, or `task_type` is missing or invalid. | No criterion results and no overall compliance result. |
| Semantic evaluation | Every applied criterion was evaluated. | Criterion results and one derived overall compliance result. |
| Execution failure | Model, runtime, timeout, response-decode, or incomplete-response failure prevented complete evaluation. | No criterion results and no overall compliance result. Do not synthesize missing judgments. |

A semantic false result means evaluation completed and found at least one responsibility-boundary violation.
A structural or execution failure does not establish semantic non-compliance.

### Workflow invocation and exceptions

- Invoke the validator immediately after Task authoring.
- Invoke the same validator after final Task Evidence is written.
- Use the same checklist-selection and result semantics at both points.
- The validator reports semantic results but does not enforce workflow continuation or release.
- A fully compliant result permits the caller to continue its normal route.
- Any semantic violation requires explicit human acceptance or rejection.
- Human acceptance preserves:
  - each violated criterion;
  - the validator reason for each violation;
  - the acceptance decision;
  - the acceptance reason.
- Human acceptance does not change the validator result or criterion contract.
- Human rejection returns work to Task correction or responsibility-boundary reconsideration.
- The caller selects the applicable return route.

## Boundary

| concern | owner |
|---|---|
| Semantic evaluation, aggregation, outcome classes, shared invocation semantics, and exception boundary | This Specification. |
| Canonical Task types, responsibility outcomes, completion judgments, and prohibited overlaps | `spec:product.design_records.authoring_standards.task_authoring`. |
| Narrow Task-authoring workflow usage relation | `spec:product.design_records.authoring_standards.task_authoring`. |
| Exact checklist artifacts and checklist review | PRODUCT-WORK-SPEC-020. |
| TRV app-local Requirement, ADR, Specification, interface, runtime, tests, and concrete human interaction | TRV under `spec:trv`. |
| PRODUCT conceptual-design authoring and TRV namespace bootstrap | PRODUCT-WORK-SPEC-021. |
| Current DRMCP parsing, structural validation, diagnostics, indexing, and tool projections | Current DRMCP authority. This Specification does not modify that authority. |
| Future DRMCP integration | A separate future Requirement or Work Item. |

TRV consumes the PRODUCT Task contract and this semantic validator contract.
TRV has no current normative dependency on DRMCP.

Concrete TRV design choices remain outside this Specification.
Historical technology candidates do not become canonical without TRV-local design decisions.

## Related specs

| ref | relation |
|---|---|
| PRODUCT-REQ-SPEC-007 | Requires semantic Task responsibility-boundary validation. |
| PRODUCT-REQ-SPEC-005 | Defines the typed single-responsibility Task requirement being evaluated. |
| PRODUCT-ADR-SPEC-015 | Defines Task-local semantic result and failure semantics. |
| PRODUCT-ADR-SPEC-016 | Defines the standalone ownership boundary and excludes current DRMCP integration. |
| PRODUCT-ADR-SPEC-017 | Defines the two invocation points and human-owned violation exceptions. |
| `spec:product.design_records.authoring_standards.task_authoring` | Defines the canonical Task responsibility contract and narrow usage relation. |
| PRODUCT-WORK-SPEC-020 | Owns exact checklist authoring and review. |
| PRODUCT-WORK-SPEC-021 | Owns reviewed PRODUCT conceptual design and TRV namespace bootstrap. |
| `spec:trv` | Owns the TRV namespace overview and routes app-local design. |

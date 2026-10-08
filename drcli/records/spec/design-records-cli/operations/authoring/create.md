# Contract: DRCLI create operation

- **id**: `spec:drcli.design_records_cli.operations.authoring.create`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.operations.authoring`
- **contract_class**: interface

## What this is

Defines the semantic inputs and outcomes for creating one PRODUCT-governed Design Record through DRCLI.
Normal success writes the created record in the same invocation; deferred proposal preparation is an explicit alternative.

## Request

A create request supplies the following semantic inputs independent of concrete CLI flag spelling.

| input | requirement |
|---|---|
| Selected repository | The repository chosen by the DRCLI repository-selection contract. |
| Record kind | One supported kind from the authoring overview. |
| Create identity | The exact identity or PRODUCT-permitted allocation placeholder and any required namespace or parent context. |
| Title | The author-supplied title required by the applicable PRODUCT authoring contract. |
| Author fields | The create-required fields owned by the applicable PRODUCT authoring standard. |
| Body source | Inline body, standard-input body, or existing body cache ID when body content is required. |
| Completion disposition | Direct write by default, or deferred proposal only when explicitly requested. |

Create identity is logical record identity, not a caller-selected physical file path.

Sequential artifacts use the exact-ID or allocation-placeholder behavior allowed by PRODUCT authority.
Task creation uses an explicit parent Work Item relation and must not infer that parent only from ID shape.

Specification creation must identify the logical spec target and distinguish leaf creation from topic-index creation without inferring the intended form from current repository state.
The concrete CLI spelling of that distinction belongs to the CLI mapping contract.

The caller supplies the author-facing fields and body defined by PRODUCT.
DRCLI generates H1 projection, metadata serialization, resolved identity projection, filename, and placement according to PRODUCT authority.

When body content is required:

- exactly one body source is accepted;
- inline and standard-input bodies are treated identically after receipt;
- a body cache ID supplies body content only;
- the body contains author-controlled body sections and excludes generated H1 and generated metadata content.

DRCLI constructs one candidate affected set containing the new record and every additional PRODUCT-required coupled change.
All create-required semantics are validated against the complete candidate before write eligibility is granted.

## Response

Direct completion has these semantic outcomes:

- the request resolves to one final canonical target identity;
- the complete candidate affected set passes required validation;
- the shared write contract confirms current identity availability and stale-state guards immediately before modification;
- repository modification occurs in the same invocation;
- the result distinguishes a completed write from a non-writing failure and exposes the resolved target identity and applicable validation diagnostics.

Deferred completion has these semantic outcomes:

- no repository files are modified during proposal preparation;
- a retained proposal is created only after the candidate can be prepared and candidate-state validation has been evaluated;
- the result returns an opaque proposal identity usable from a later independent invocation;
- later acceptance follows the deferred-proposal and shared write contracts.

If a retryable failure occurs after inline or standard-input body content has been received, the response includes an opaque retry body cache ID according to the retry-cache contract.
The caller can retry without resending that body.

## Errors

| condition | machine-readable classification | required effect |
|---|---|---|
| Unsupported record kind | `authoring_invalid_request` (`reason=unsupported_kind`), exit 1 | Reject without modification. |
| Missing or invalid create identity or required parent context | `authoring_invalid_request` (`reason=create_identity_missing_or_invalid` or `parent_context_missing_or_invalid`), exit 1; missing required CLI operand is usage 2 | Reject without modification. |
| Existing canonical identity or generated placement collision | `authoring_identity_collision` (`collision_kind=canonical_identity` or `generated_placement`), exit 1 | Reject without modification. |
| Caller supplies a generated-only value as authoritative input | `authoring_invalid_request` (`reason=generated_only_input`), exit 1 | Reject without modification. |
| Required author field or body content is missing or invalid under PRODUCT authority | `authoring_invalid_request` (`reason=required_field_missing_or_invalid` or `body_content_missing_or_invalid`), exit 1; CLI-required operand/body-source absence is usage 2 | Reject without modification. If full candidate validation identifies blocking canonical findings, use `authoring_candidate_validation_failed` instead. |
| More than one body source is supplied | CLI `usage_error`, exit 2; `authoring_invalid_request` (`reason=body_source_multiple`), exit 1 only for a type-valid logical request that reaches this semantic boundary | Reject without modification. |
| Body cache ID is unknown or no longer retained | `authoring_body_cache_unavailable` (`reason=unknown`, `expired`, or `invalid`), exit 3 | Reject without modification. |
| Candidate-state validation fails | `authoring_candidate_validation_failed`, exit 1 | Reject without modification; retain canonical finding identities. |
| Stale allocation, target-state change, or collision is detected at write time | `authoring_stale_state` (`guard=allocation` or `target_state`) or `authoring_identity_collision`, exit 1 | Reject without modification. |
| Retryable preparation failure occurs after body receipt | Underlying specific authoring code; `authoring_operational_failure` (`reason=candidate_preparation`), exit 3 where no trustworthy result exists | Do not claim success; preserve the received body and return `error.context.body_cache_id`; if preservation fails include `error.context.cache_error` instead. |

If candidate preparation, repository inspection, or write-completion cannot produce a trustworthy result for an otherwise valid request, return `authoring_operational_failure` with `reason=candidate_preparation`, `repository_state`, or `write_completion` as established; exit 3. Retry-cache context follows the diagnostic catalog whenever the received body is retryable. The shared write contract determines actual write-state reporting.

No error path may require an interactive confirmation to continue.

## Boundary

PRODUCT owns all record-kind semantics, required fields, lifecycle gates, canonical body sections, identity grammar, and generated placement.
This contract owns only DRCLI create orchestration, body transport choice, direct-versus-deferred completion, candidate preparation, and the handoff to shared write guards.

## Related specs

| ref | relation |
|---|---|
| spec:drcli.design_records_cli.operations.authoring | Parent authoring overview. |
| spec:drcli.design_records_cli.operations.authoring.write | Shared write guards and completion eligibility. |
| spec:drcli.design_records_cli.operations.authoring.retry_cache | Retry body preservation. |
| spec:drcli.design_records_cli.operations.authoring.deferred_proposal | Optional deferred completion. |
| spec:product.design_records.authoring_standards | Record-kind authoring authority. |

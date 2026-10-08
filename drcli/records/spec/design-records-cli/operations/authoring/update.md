# Contract: DRCLI update operation

- **id**: `spec:drcli.design_records_cli.operations.authoring.update`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.operations.authoring`
- **contract_class**: interface

## What this is

Defines DRCLI partial update semantics for one existing PRODUCT-governed Design Record.
One invocation may compose the requested partial changes into one candidate state and normally writes that candidate immediately.

## Request

An update request supplies the following semantic inputs independent of concrete CLI flag spelling.

| input | requirement |
|---|---|
| Selected repository | The repository chosen by the DRCLI repository-selection contract. |
| Existing target | One exact existing canonical Design Record identity or spec ref. Allocation placeholders are invalid for update. |
| Partial changes | One or more author-controlled metadata-field changes and, when needed, one or more named body-section replacements supported by PRODUCT authority. |
| Body source | Inline body, standard-input body, or existing body cache ID for each body-bearing replacement. |
| Completion disposition | Direct write by default, or deferred proposal only when explicitly requested. |

Partial update preserves every omitted author-controlled field and section.
Generated identity, generated H1 projection, generated placement, and other PRODUCT-owned generated values are not ordinary update targets.

All requested changes in one invocation form one candidate.
No intermediate partial state is persisted or validated as though it were the final requested result.

If multiple requested changes conflict over the same persisted subject, DRCLI rejects the request before modification rather than choosing an order that changes meaning.
Where several non-conflicting changes are supplied, validation is performed against their final combined candidate state.

Named section replacement must resolve one intended existing section under the applicable PRODUCT section rules.
Zero-match or ambiguous selection is a non-writing failure.
Diagnostics may expose candidate headings when that helps the caller repair the selector.

A body-bearing replacement accepts exactly one body source.
Metadata-only changes do not consume a body source.
Using a body cache ID supplies only replacement body content and never changes repository selection, target identity, or the remaining update request.

The affected set contains the requested target and every additional PRODUCT-required coupled change.
PRODUCT-required lifecycle or relation effects are not optional DRCLI conveniences.

## Response

Direct completion has these semantic outcomes:

- the target resolves to the same existing record intended by the caller;
- the complete final candidate affected set passes required validation;
- shared write guards re-check current target state and required validation immediately before modification;
- all requested changes that are part of the candidate complete in the same invocation;
- the result distinguishes a completed write, a valid no-op, and a non-writing failure.

A valid no-op occurs when the fully interpreted candidate would not change persisted content.
A no-op does not modify repository files and does not create a deferred proposal.

Deferred completion has these semantic outcomes:

- candidate preparation and validation occur without repository modification;
- a retained proposal is created only when the requested update is not a no-op and preparation succeeds;
- later acceptance uses the shared write contract against current repository state.

If a retryable failure occurs after an inline or standard-input replacement body has been received, the response includes an opaque retry body cache ID according to the retry-cache contract.

## Errors

| condition | machine-readable classification | required effect |
|---|---|---|
| Target does not resolve exactly to the requested existing record | `authoring_target_unresolved` (`reason=malformed`, `unresolved`, or `conflicted`), exit 1 | Reject without modification. |
| Allocation placeholder is used as an update target | `authoring_invalid_request` (`reason=update_placeholder`), exit 1 | Reject without modification. |
| Caller attempts to update a generated-only identity or placement value | `authoring_invalid_request` (`reason=generated_only_input`), exit 1 | Reject without modification. |
| Partial changes conflict with each other | `authoring_invalid_request` (`reason=conflicting_partial_changes`), exit 1 | Reject without modification. |
| Named section selection has zero or multiple matches | `authoring_section_selection_failed` (exact `h2_title`, `match_count`, `available_h2_headings`), exit 1 | Reject without modification. |
| Body source is missing where required or supplied where not consumed | CLI `usage_error`, exit 2 when recognized by CLI input mapping; otherwise `authoring_invalid_request` (`reason=body_source_missing` or `body_source_unexpected`), exit 1 for an admitted type-valid logical request | Reject without modification. |
| More than one body source is supplied for one body-bearing replacement | CLI `usage_error`, exit 2; `authoring_invalid_request` (`reason=body_source_multiple`), exit 1 only at the type-valid semantic boundary | Reject without modification. |
| Body cache ID is unknown or no longer retained | `authoring_body_cache_unavailable` (`reason=unknown`, `expired`, or `invalid`), exit 3 | Reject without modification. |
| Final candidate-state validation fails | `authoring_candidate_validation_failed`, exit 1 | Reject without modification; retain canonical finding identities. |
| Target becomes stale or changes identity before write | `authoring_stale_state` (`guard=target_state` or `target_identity`), exit 1 | Reject without modification. |
| Retryable preparation failure occurs after body receipt | Underlying specific authoring code; `authoring_operational_failure` (`reason=candidate_preparation`), exit 3 where no trustworthy result exists | Do not claim success; preserve received body as `error.context.body_cache_id`, or report a distinct `error.context.cache_error` when preservation fails. |

If candidate preparation, repository inspection, or write-completion cannot produce a trustworthy result for an otherwise valid request, return `authoring_operational_failure` with `reason=candidate_preparation`, `repository_state`, or `write_completion` as established; exit 3. Retry-cache context follows the catalog when an inline/stdin replacement body has been received. The shared write contract determines actual write-state reporting.

No failure may be repaired by silently weakening PRODUCT validation or changing the requested target.

## Boundary

PRODUCT owns which metadata fields, lifecycle changes, section headings, relation updates, and coupled effects are valid.
This contract owns DRCLI partial-update composition, body transport choice, no-op handling, direct-versus-deferred completion, and the handoff to shared write guards.

## Related specs

| ref | relation |
|---|---|
| spec:drcli.design_records_cli.operations.authoring | Parent authoring overview. |
| spec:drcli.design_records_cli.operations.authoring.write | Shared write guards and completion eligibility. |
| spec:drcli.design_records_cli.operations.authoring.retry_cache | Retry body preservation. |
| spec:drcli.design_records_cli.operations.authoring.deferred_proposal | Optional deferred completion. |
| spec:product.design_records.authoring_standards | Record-kind authoring authority. |

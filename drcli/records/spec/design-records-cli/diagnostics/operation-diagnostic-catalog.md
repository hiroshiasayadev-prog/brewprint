# Reference: Operation diagnostic catalog

- **id**: `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.diagnostics`

## What this is

Defines canonical semantic operation diagnostic codes and code-specific context contracts used by DRCLI operations.

## Non-goals

- Defining operation trigger conditions or normal result fields.
- Defining validation finding codes, validation classification, or validation result projection.
- Defining CLI parser, command-selection, help, usage, stream, or exit-status diagnostics.
- Exposing repair guidance, internal exceptions, debug information, or implementation paths unless an operation result explicitly owns a path field.

## Catalog rules

| rule | contract |
|---|---|
| Code spelling | Lowercase snake_case. |
| Code identity | The semantic operation classification name and diagnostic `code` are identical. |
| Class | Each code is fixed as either `error` or `warning`. |
| Placement | Each code may appear only at cataloged placements. |
| Context | Each matching variant permits exactly its cataloged fields. |
| Unknown codes | Prohibited. |
| Unknown context fields | Prohibited. |
| Message | Required by the shared envelope but exact wording is not cataloged. |
| Expansion | A new semantic operation outcome requires an operation Specification update and a catalog entry before external use. |
| CLI usage separation | CLI usage diagnostics are not entries in this catalog unless a later command/runtime Specification explicitly establishes a distinct mapping. |

## Error codes

| code | placement | context | trigger owner |
|---|---|---|---|
| `invalid_selector` | Top-level error. | Required variant. | Discovery, listing, search, and scope-validation operation Specifications. |
| `invalid_projection` | Top-level error. | Required. | Sequential record listing and tree child listing. |
| `invalid_limit` | Top-level error. | Required. | Sequential record listing, tree child listing, and validation output limits. |
| `invalid_offset` | Top-level error. | Required. | Sequential record listing and tree child listing. |
| `invalid_depth` | Top-level error. | Required. | Tree overview. |
| `invalid_node_limit` | Top-level error. | Required. | Tree overview. |
| `invalid_selector_count` | Top-level error. | Required. | Batch retrieval results and exact record validation. |
| `invalid_source_content_limit` | Top-level error. | Required. | Retrieval output limits. |
| `invalid_match_limit` | Top-level error. | Required. | Search results and limits. |
| `invalid_snippet_limit` | Top-level error. | Required. | Search results and limits. |
| `empty_query` | Top-level error. | Prohibited. | Lexical record search. |
| `invalid_pattern` | Top-level error. | Prohibited. | Lexical match model. |
| `configuration_failure` | Top-level error. | Prohibited. | Every completed semantic operation. |
| `execution_failure` | Top-level error. | Prohibited. | Every completed semantic operation. |
| `malformed_selector` | Selector-level error. | Prohibited. | Exact record retrieval, H2 section retrieval, and exact record validation. |
| `unresolved_record` | Selector-level error. | Prohibited. | Exact record retrieval, H2 section retrieval, and exact record validation. |
| `conflicted_record` | Selector-level error. | Prohibited. | Exact record retrieval, H2 section retrieval, and exact record validation. |
| `node_not_found` | Top-level error. | Required variant. | Tree child listing, tree overview, lexical record search, and scope validation. |
| `section_not_found` | Selector-level error. | Required. | H2 section retrieval. |
| `section_data_unavailable` | Selector-level error. | Prohibited. | H2 section retrieval. |
| `authoring_invalid_request` | Authoring failure `error`. | Required. | Create and update. |
| `authoring_target_unresolved` | Authoring failure `error`. | Required. | Update and write. |
| `authoring_section_selection_failed` | Authoring failure `error`. | Required. | Update. |
| `authoring_identity_collision` | Authoring failure `error`. | Required. | Create and write. |
| `authoring_stale_state` | Authoring failure `error`. | Required. | Write and deferred acceptance. |
| `authoring_repository_mismatch` | Authoring failure `error`. | Required. | Write and deferred proposal. |
| `authoring_candidate_validation_failed` | Authoring failure `error`. | Required. | Create, update, write, and deferred proposal. |
| `authoring_body_cache_unavailable` | Authoring failure `error`. | Required. | Retry cache, create, and update. |
| `authoring_proposal_unavailable` | Authoring failure `error`. | Required. | Deferred proposal. |
| `authoring_proposal_ineligible` | Authoring failure `error`. | Required. | Deferred proposal. |
| `authoring_operational_failure` | Authoring failure `error`. | Required. | All authoring operations. |
| `authoring_cache_preservation_failed` | Nested authoring `cache_error` only. | Prohibited. | Retry cache. |

### `invalid_selector`

`invalid_selector` uses exactly one matching context variant.

| condition | required context fields | field order and value rules |
|---|---|---|
| Required scope selectors are absent | `missing_selector_fields` | List of missing field names in operation request-field order. |
| A supplied selector combination is not accepted | `accepted_selector_combinations` | List of field-name lists in operation selector-progression order. Each inner list uses request-field order. An empty inner list represents no selectors. |
| App namespace is unavailable | `app_namespace`, `available_app_namespaces` | Supplied string and available strings in ascending simple string order. |
| Artifact kind is unavailable without a structure restriction | `artifact_kind`, `available_artifact_kinds` | Supplied string and available strings in ascending simple string order. |
| Artifact kind is unavailable or ineligible for a sequential operation | `artifact_kind`, `available_sequential_artifact_kinds` | Supplied string and available strings in ascending simple string order. |
| Sequential domain is unavailable | `domain_namespace`, `available_domain_namespaces` | Supplied string and available strings in ascending simple string order. |
| `node_ref` is malformed or does not identify a tree record kind | `node_ref`, `accepted_tree_record_kinds` | Supplied string and accepted tree record-kind strings in ascending simple string order. |
| `subtree_root_ref` is malformed or does not identify a tree record kind | `subtree_root_ref`, `accepted_tree_record_kinds` | Supplied string and accepted tree record-kind strings in ascending simple string order. |

No `invalid_selector` variant permits another context field.
The trigger-owning operation determines which variant applies.

### `invalid_projection`

| field | requirement | value |
|---|---|---|
| `invalid_projection_values` | required | Unknown supplied values in request occurrence order with duplicate occurrences preserved. |
| `available_projection_values` | required | Accepted values in ascending simple string order. |

No other context field is permitted.

### Numeric request errors

| code | required context fields | prohibited common fields |
|---|---|---|
| `invalid_limit` | `minimum`, `maximum`, `default`, `actual` | Byte-suffixed fields. |
| `invalid_offset` | `minimum`, `default`, `actual` | `maximum` and byte-suffixed fields. |
| `invalid_depth` | `minimum`, `maximum`, `default`, `actual` | Byte-suffixed fields. |
| `invalid_node_limit` | `minimum`, `maximum`, `default`, `actual` | Byte-suffixed fields. |
| `invalid_selector_count` | `minimum`, `maximum`, `actual` | `default` and byte-suffixed fields. |
| `invalid_source_content_limit` | `minimum_bytes`, `maximum_bytes`, `actual_bytes` | Unsuffixed numeric-bound fields and `default`. |
| `invalid_match_limit` | `minimum`, `maximum`, `default`, `actual` | Byte-suffixed fields. |
| `invalid_snippet_limit` | `minimum`, `maximum`, `default`, `actual` | Byte-suffixed fields. |

Each context value is an integer.
`actual` or `actual_bytes` is the exact supplied value in the type-valid logical operation request.
A value that cannot be decoded into the required request type is a CLI usage concern, not a semantic operation diagnostic.
No numeric request error permits another context field.

### `node_not_found`

`node_not_found` uses exactly one matching context variant.

| consuming selector | required context field | value |
|---|---|---|
| Tree navigation `node_ref` | `node_ref` | Exact supplied node ref. |
| Search or scope-validation `subtree_root_ref` | `subtree_root_ref` | Exact supplied subtree root ref. |

No variant permits another context field.
The diagnostic does not include available nodes, suggestions, or physical paths.

### `section_not_found`

| field | requirement | value |
|---|---|---|
| `available_h2_headings` | required | Real H2 title strings in source order with duplicate occurrences preserved. |

The list may be empty.
The list excludes H1 and H3-or-deeper headings.
The list preserves exact title text without the leading `## ` marker.
The parent result already contains the requested `h2_title`.
No other context field is permitted.

### Context-prohibited errors

The following non-authoring codes omit `context`:

- `configuration_failure`;
- `execution_failure`;
- `malformed_selector`;
- `unresolved_record`;
- `conflicted_record`;
- `section_data_unavailable`;
- `empty_query`;
- `invalid_pattern`.

A selector-level parent result retains every field established before the error.
Conflict members, source defects, and physical source locations belong to validation and current-record contracts rather than operation diagnostics.

## Authoring failure codes and context

The following authoring-owned classifications apply after a type-valid logical request exists. Each `authoring_*` code above has fixed class `error`. An authoring failure is one top-level `error` diagnostic document; its code and required context discriminator (when present) identify the exact failure class. Codes and discriminator values must not be replaced by prose or implementation-specific alternatives. A trustworthy negative semantic result uses stdout/status `1`; an unavailable retained state, I/O, permission, repository/package resource, or otherwise untrustworthy completion uses stderr/status `3`. CLI usage errors and internal failures retain statuses `2` and `4` under the CLI contract.

| code | required context | normative exit category |
|---|---|---|
| `authoring_invalid_request` | `reason` | `contract_result_failure` (1) |
| `authoring_target_unresolved` | `record_ref`, `reason` | `contract_result_failure` (1) |
| `authoring_section_selection_failed` | `h2_title`, `match_count`, `available_h2_headings` | `contract_result_failure` (1) |
| `authoring_identity_collision` | `collision_kind` | `contract_result_failure` (1) |
| `authoring_stale_state` | `guard` | `contract_result_failure` (1) |
| `authoring_repository_mismatch` | `phase`; `proposal_id` iff `phase=accept` | `contract_result_failure` (1) |
| `authoring_candidate_validation_failed` | `affected_findings` | `contract_result_failure` (1) |
| `authoring_body_cache_unavailable` | `body_cache_id`, `reason` | `execution_failure` (3) |
| `authoring_proposal_unavailable` | `proposal_id`, `reason` | `execution_failure` (3) |
| `authoring_proposal_ineligible` | `proposal_id`, `state` | `contract_result_failure` (1) |
| `authoring_operational_failure` | `reason` | `execution_failure` (3) |
| `authoring_cache_preservation_failed` | No context | Nested `cache_error`; does not independently override the enclosing failure category |

`authoring_invalid_request.reason` is exactly one of: `unsupported_kind`, `create_identity_missing_or_invalid`, `parent_context_missing_or_invalid`, `generated_only_input`, `required_field_missing_or_invalid`, `body_content_missing_or_invalid`, `body_source_missing`, `body_source_unexpected`, `body_source_multiple`, `update_placeholder`, `conflicting_partial_changes`. These describe DRCLI request rejection, not new PRODUCT validation rule codes. The `unsupported_kind` variant additionally requires `record_kind` containing the exact supplied kind string. The `required_field_missing_or_invalid` variant additionally requires nonempty `field_names` in the applicable PRODUCT-defined author-field order, identifying the affected fields without redefining why PRODUCT rejects them. No other variant permits `record_kind` or `field_names`. Where the corresponding CLI surface detects missing required command input, mutually exclusive body-source flags, invalid option values, or another invalid CLI shape, that failure instead follows the existing actionable CLI `usage_error` contract (2) **before** a logical authoring operation begins; a CLI usage error must not be recast as this operation code.

`authoring_target_unresolved.record_ref` is the exact supplied target string; `reason` is `malformed`, `unresolved`, or `conflicted`. `authoring_section_selection_failed.h2_title` is the exact supplied H2 title, `match_count` is the actual nonnegative count of matching H2 headings, and `available_h2_headings` is the complete real H2 title list in source order, preserving duplicates; the failure applies only when `match_count` is not one. `authoring_identity_collision.collision_kind` is `canonical_identity` or `generated_placement`.

`authoring_stale_state.guard` is `allocation`, `target_state`, `target_identity`, `affected_set`, or `target_resolution`, identifying the first failed write guard in the shared write-contract order. `authoring_repository_mismatch.phase` is `prepared_write` or `accept`; `proposal_id` is the exact supplied opaque ID only for acceptance. These are established, trustworthy non-writing guards; failure to construct or inspect the underlying state reliably instead uses `authoring_operational_failure`.

`authoring_candidate_validation_failed.affected_findings` is a nonempty array of `{record_ref, findings}` entries, ordered by each entry's own canonical `record_ref` in ascending simple string order, independently of candidate affected-set enumeration, traversal, or construction order. Each entry's `record_ref` identifies that same entry's affected candidate record; `findings` is a nonempty array of blocking PRODUCT-owned findings in that record's owning canonical validation order. Each finding retains exactly the owning validation contract's `classification`, `code`, `message`, `contract_ref`, and `subject`; DRCLI neither assigns replacement codes nor changes the finding's subject semantics. No nonaffected repository findings are included. If candidate validation itself cannot produce trustworthy findings, use `authoring_operational_failure`, not a fabricated PRODUCT finding.

`authoring_body_cache_unavailable.reason` and `authoring_proposal_unavailable.reason` are `unknown`, `expired`, or `invalid`. The ID fields are the exact supplied opaque strings. The selected installation's unavailable retained state cannot supply a trustworthy authoring result and therefore uses status 3, without assuming a retention duration. `authoring_proposal_ineligible.state` is `discarded` or `accepted`; both are reliably known lifecycle states, and acceptance must never repeat a write.

`authoring_operational_failure.reason` is `candidate_preparation`, `repository_state`, `retained_state_io`, `proposal_persistence`, `body_cache_persistence`, or `write_completion`. All primary authoring error contexts, regardless of code, additionally require `resolved_target_ref` whenever the canonical target is safely resolved (omitted otherwise), and require `write_state`: `not_modified`, `modified`, or `unknown` as supported by actual write-state evidence. When `not_modified`, `modified_refs` is required as an empty array; when `modified` or `unknown`, `modified_refs` is forbidden. `known_modified_refs` is permitted only for `write_state=modified`, only for identities independently established as modified, in actual modification order; it is omitted otherwise. All semantic rejections before modification, including failed pre-write guards and proposal lifecycle rejections, have `write_state=not_modified` and `modified_refs=[]`. A failure of reliable completion after some modification must explicitly use `write_state=modified` and must never be relabeled as a verified non-write. An indeterminate post-write I/O state uses `unknown`, not an unsupported success/no-write claim. These fields communicate state without specifying rollback or atomicity machinery.

Every authoring `error.context` additionally permits the two retry-cache fields `body_cache_id` and `cache_error` under these exclusive rules: `body_cache_id` is required if a retryable failed attempt received inline/stdin body and preservation succeeded; it is the exact retained opaque ID. `cache_error` is required instead if preservation was attempted and failed, with the nested shape `{code: "authoring_cache_preservation_failed", message: <nonempty English text>}`. Neither appears if preservation was not applicable. They never appear together. This preserves a received body only when an actual retained entry exists and reports primary authoring failure and preservation failure separately in one document. `authoring_cache_preservation_failed` is forbidden as a standalone primary `error`. Except for the stated code-specific fields, the required write-state fields, `known_modified_refs` when allowed, and these conditional retry fields, no additional authoring context fields are permitted.

### Authoring error selection

CLI parsing and actionable usage validation precede semantic authoring diagnostics. Once a type-valid request exists, create/update request checks follow the listed Error-row order within each operation, then retained body-source resolution, candidate preparation and complete candidate validation, then the ordered shared write guards. A conclusive specific collision or stale-guard failure is reported instead of a generic request or operational failure. For deferred proposal actions, establish the retained ID's availability, then its lifecycle eligibility for the requested action, repository binding, and the shared write guards in their defined order. Within a shared write-guard step, evaluate target identity/resolution, target staleness, canonical-identity availability, and generated-placement availability in that order where applicable; the first established failure is selected. If an operational condition prevents reliable evaluation of a semantic check, do not fabricate the negative semantic result; report the applicable operational failure instead. A single candidate-validation failure contains all blocking findings in canonical order; it is not split into multiple primary errors. This ordering selects the first primary `error` when several conditions coexist and does not bypass any required eligibility check before a write.

The presence of a returned `body_cache_id` asserts that the exact received body has actually been retained and is reusable at response time, subject to subsequent retention expiry. No cache ID is emitted for unsuccessful preservation.

## Warning codes

| code | placement | required context | trigger owner |
|---|---|---|---|
| `source_content_limit_ignored` | Top-level warning. | `source_content_limit_bytes` | Exact record retrieval. |
| `source_content_omitted_by_limit` | Selector-level warning. | `content_size_bytes` | Exact record retrieval and retrieval output limits. |
| `duplicate_h2` | Selector-level warning. | None. | H2 section retrieval. |
| `section_content_omitted_by_limit` | Selector-level warning. | `content_size_bytes` | H2 section retrieval and retrieval output limits. |
| `search_target_unavailable` | Top-level warning. | `canonical_ref`, `target` | Lexical match model. |
| `duplicate_requested_ref` | Top-level warning. | `ref`, `first_index`, `duplicate_indexes` | Exact record validation. |

### `duplicate_requested_ref`

| field | requirement | value |
|---|---|---|
| `ref` | required | Exact supplied selector string. |
| `first_index` | required | Zero-based index of the effective first occurrence. |
| `duplicate_indexes` | required | Later zero-based occurrence indexes in ascending order. |

The warning appears once for each exact duplicate selector string.
`duplicate_indexes` contains at least one integer.
No other context field is permitted.
The trigger-owning operation defines warning collection order.

### `search_target_unavailable`

| field | requirement | value |
|---|---|---|
| `canonical_ref` | required | Canonical ref of the uniquely addressable current record. |
| `target` | required | `title`, `metadata_source`, `h2_heading`, or `h2_section_content`. |

No other context field is permitted.
The diagnostic does not include defect details, source positions, or physical paths.
The lexical match model defines warning aggregation and order.

### `source_content_limit_ignored`

| field | requirement | value |
|---|---|---|
| `source_content_limit_bytes` | required | Exact supplied integer that was ignored. |

This warning applies when the exact-record retrieval operation receives the field while source-content projection is disabled.
The ignored value is not range-validated.
No other context field is permitted.

### Content omission warnings

| code | field | requirement | value |
|---|---|---|---|
| `source_content_omitted_by_limit` | `content_size_bytes` | required | Complete record source-content size as UTF-8 bytes. |
| `section_content_omitted_by_limit` | `content_size_bytes` | required | Complete H2 section-content size as UTF-8 bytes. |

No other context field is permitted.
The warnings preserve successful selector results.

### `duplicate_h2`

`duplicate_h2` prohibits `context`.
The parent result contains the requested `h2_title`.
Duplicate counts and source positions are not operation diagnostic fields.

## List ordering

| context field | normative order |
|---|---|
| `available_app_namespaces` | Ascending simple string order. |
| `available_artifact_kinds` | Ascending simple string order. |
| `available_sequential_artifact_kinds` | Ascending simple string order. |
| `available_domain_namespaces` | Ascending simple string order. |
| `available_projection_values` | Ascending simple string order. |
| `accepted_tree_record_kinds` | Ascending simple string order. |
| `missing_selector_fields` | Operation request-field order. |
| `accepted_selector_combinations` | Operation selector-progression order; each combination uses request-field order. |
| `invalid_projection_values` | Request occurrence order with duplicate occurrences preserved. |
| `available_h2_headings` | Source order with duplicate occurrences preserved. |
| `duplicate_indexes` | Ascending numeric order. |

## Boundary

| concern | owner |
|---|---|
| Canonical semantic code spelling, class, placement eligibility, and context fields | This Specification. |
| Shared diagnostic object and response-state rules | `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_envelope`. |
| Trigger conditions, result fields established before failure, and warning order | The consuming operation Specification. |
| Current record, identity, node, and scope semantics | `spec:drcli.design_records_cli.current_record_model`. |
| Validation finding semantics and artifact conformance rules | The canonical Specification that owns the validated rule. |
| CLI syntax and usage diagnostic behavior | `spec:drcli.design_records_cli.cli`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_envelope` | Defines the object and placement contract used by every catalog entry. |
| `spec:drcli.design_records_cli.operations.validation.scope_validation` | Owns broad validation selector and scope diagnostic triggers. |
| `spec:drcli.design_records_cli.operations.validation.exact_record_validation` | Owns exact validation selector outcomes and duplicate-selector warnings. |
| `spec:drcli.design_records_cli.operations.validation.validation_output_limits` | Owns validation output-limit diagnostic triggers. |
| `spec:drcli.design_records_cli.cli` | Owns CLI usage diagnostics and concrete CLI presentation. |

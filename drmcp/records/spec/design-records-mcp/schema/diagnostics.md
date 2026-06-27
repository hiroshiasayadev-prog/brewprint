# Reference: Diagnostics

- **id**: `spec:drmcp.design_records_mcp.schema.diagnostics`
- **status**: draft
- **date**: 2026-06-28
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines the shared machine-readable envelope, associations, category vocabulary, severity, ordering, and duplicate-suppression rules for DRMCP operation warnings, repository diagnostics, and authoring diagnostics.

PRODUCT specifications own whether record content, document shape, lifecycle state, or declared relations are semantically valid. DRMCP maps those authority-owned conditions into the representation defined here.

## Current contract

### Placement by operation

| surface | collection | contract |
|---|---|---|
| `list_records` | top-level `warnings` | Contains only W004-defined operation warning triggers. |
| `get_records` | top-level `warnings` | Contains only W004-defined per-ref and duplicate warning triggers. |
| `resolve_reference` | top-level `diagnostics` | Explains the cause of `unresolved` or `unsupported` without changing the W005 public status vocabulary. |
| `validate_records` | top-level `diagnostics` | Contains the unified source, record, conflict, and relation findings selected by the T02 execution contract. |
| Authoring proposal validation | `validation.diagnostics` | Uses the same repository-diagnostic representation as `validate_records`, limited to the proposal-local affected set. |
| Authoring operation outcomes | operation-level `diagnostics` | Uses the shared envelope while preserving authoring-owned categories and trigger semantics. |

Request-shape errors and execution failures are not normal warning or validation-diagnostic entries when the owning operation contract says that no normal response wrapper is returned.

### Shared envelope

Every warning or diagnostic entry contains:

```json
{
  "category": "invalid_field_value",
  "severity": "error",
  "message": "status value is not accepted for this record kind"
}
```

| field | required | meaning |
|---|---:|---|
| `category` | yes | Stable machine-readable cause identifier. |
| `severity` | yes | `error`, `warning`, or `info`. |
| `message` | yes | Human-readable explanation. It is not an identity or ordering field. |
| `subject` | category-dependent | The record, source, configuration area, or request item to which the entry applies. |
| `field` | category-dependent | Affected metadata field, narrative section, or ordered field item. |
| `value` | category-dependent | Invalid present value and an optional explicit expectation. |
| `target` | category-dependent | Relation or lookup target association. |
| `occurrence` | category-dependent | Exact-request first occurrence and duplicate occurrence association. |
| `conflict` | category-dependent | Current identity or legacy lookup conflict association. |
| `location` | category-dependent | T04-owned machine-readable source location for a repairable source-backed finding. |

Each category profile defines which optional associations are required, optional, or prohibited.

The following are not shared diagnostic-envelope fields:

- scalar `record_id`, `source_ref`, or `requested_id` duplicates beside canonical `ref` association;
- scalar `field` or scalar `value` shortcuts;
- a standalone raw top-level `path` field;
- an untyped generic `details` object.

### Subject association

`subject` is a discriminated object.

| `subject.type` | required fields | optional fields | use |
|---|---|---|---|
| `record` | `ref` | `record_kind` | A current canonical record or spec identity, including an identity whose duplicate conflict prevents unique addressability. |
| `source` | `app_namespace`, `record_kind` | `ref` when determinable | A current discovered source, including parse-failed and identity-less validation-only sources. |
| `configuration` | `component` | `app_namespace` | Repository, current-root, legacy-root, active-index, or legacy-lookup configuration context. |
| `request_item` | `operation` | none | One item supplied to a read or resolver operation. |

Canonical current identity uses `ref`.
Physical paths are never stored inside `subject`.
An identity-less source omits `ref`; its exact repair target is supplied by `location` when known.

### Field, value, and target associations

`field` contains:

| field | required | meaning |
|---|---:|---|
| `name` | yes | Metadata field name or canonical required section heading. |
| `item_index` | no | Zero-based index of one ordered list item. |

`value` contains:

| field | required | meaning |
|---|---:|---|
| `actual` | yes | Present invalid value. A missing field or section omits `value`. |
| `expected` | no | Explicit machine-readable accepted value, values, or shape when the authority can be represented without prose inference. |

`target` contains:

| field | required | meaning |
|---|---:|---|
| `ref` | category-dependent | Exact current canonical or accepted legacy issued ref being checked. |
| `actual_kind` | no | Kind found for a uniquely identified target when it differs from the required kind. |
| `expected_kinds` | no | Ordered accepted target kinds. |
| `reciprocal_field` | no | Counterpart field required by the applicable reciprocity rule. |
| `lookup_state` | no | Machine-readable non-resolved lookup cause. |

Accepted `target.lookup_state` values are:

- `current_unresolved`;
- `legacy_disabled`;
- `legacy_unresolved`;
- `legacy_unreadable`.

A declared relation value that matches no accepted current or legacy lookup grammar has no lookup state because no lookup runs.

A current duplicate target is distinguished through `conflict.type: current_identity` rather than another lookup-state value.
Physical paths are not placed in `field`, `value`, or `target`.

### Request occurrence association

`get_records` per-ref warnings use `occurrence`:

```json
{
  "occurrence": {
    "ref": "DRMCP-WORK-MCP-006",
    "first_index": 0,
    "duplicate_indexes": [2, 4]
  }
}
```

| field | required | meaning |
|---|---:|---|
| `ref` | yes | Exact string supplied by the caller. |
| `first_index` | yes | Zero-based index of the first occurrence. |
| `duplicate_indexes` | only for `duplicate_requested_ref` | Ordered zero-based indexes of every ignored later exact-equal occurrence. |

Exact string equality controls duplicate detection.
The first occurrence remains effective.
One `duplicate_requested_ref` entry is emitted per duplicated exact string, not per ignored occurrence.

### Conflict association

`conflict.type` is one of:

| type | required fields | meaning |
|---|---|---|
| `current_identity` | `source_count`, `members` | Two or more current sources claim the same canonical current identity. |
| `legacy_lookup` | `candidate_count`, `candidates` | Two or more configured legacy sources produce the same issued legacy ID. |

For `current_identity`:

- `source_count` is the number of conflicting current sources;
- `members` contains one entry per conflicting source;
- each member contains the T04-owned `location` association;
- no source winner is selected.

For `legacy_lookup`:

- `candidate_count` is the number of configured legacy candidates;
- `candidates` contains one entry per conflicting legacy source;
- each candidate contains the T04-owned `location` association;
- no candidate winner is selected;
- legacy archive sources remain lookup targets and do not become repository-validation subjects.

Disabled, missing, unreadable, resolved, and unsupported legacy lookup outcomes are not conflict objects.

### Source-location slot and T04 boundary

`location` is the shared source-location association reserved by this contract.

It is required when a repository or proposal-local validation diagnostic identifies a known repairable source file, including:

- identity-less, parse-failed, or unreadable current sources when their source is known;
- source-backed metadata, section, identity, and relation diagnostics;
- every member of a current identity conflict;
- every candidate of a legacy lookup conflict;
- persisted-file authoring diagnostics when the authoring operation exposes a repair target.

It is omitted when no repository file exists, including malformed request items, unresolved request refs without a source, unknown proposal IDs, and expired body-cache IDs.

This T03 contract does not define the fields inside `location`.
T04 owns:

- repository-relative versus absolute representation;
- current-root and legacy-root identification;
- separator normalization and Windows-path handling;
- the stable source-location sort key;
- operation-specific exceptional path-exposure policy.

An opaque internal source token alone is not an acceptable replacement for a repairable location.
Normal successful list, retrieval, and resolver target objects remain path-free under W004 and W005.

### Severity and validation blocking

| severity | meaning |
|---|---|
| `error` | Applicable repository or record contract is violated. Blocks `validate_records.ok`. |
| `warning` | Operation completed normally but one item failed, a projection degraded, or repair attention is warranted without invalidating the entire response wrapper. |
| `info` | Non-blocking notice, ignored duplicate occurrence, or repair guidance. |

`validate_records.ok` is `false` when at least one returned diagnostic has severity `error`.
It is `true` when no returned diagnostic has severity `error`, including responses containing only `warning` or `info` entries.

Malformed validation requests and untrustworthy mandatory configuration or index state remain request or execution failures outside the normal validation wrapper.

Severity is context-sensitive but deterministic.
The same condition under the same operation and applicable authority always maps to the same severity.

### Repository-validation categories

| category | normal severity | required associations | meaning |
|---|---|---|---|
| `source_unreadable` | `error` | `subject: source`, `location` when known | A selected current source cannot be read. |
| `source_syntax_invalid` | `error` | `subject`, `location`; optional `field`, `value` | Source or metadata structure cannot be consumed under the current format contract. |
| `identity_invalid` | `error` | `subject`, `location`; optional `field`, `value` | Canonical identity grammar, H1 identity, filename consistency, metadata identity consistency, or path-derived spec identity is invalid. |
| `missing_required_field` | `error` | `subject`, `field`, `location` | An applicable authority requires a field that is absent. |
| `empty_required_field` | `error` | `subject`, `field`, `value`, `location` | An applicable authority requires a present field to be non-empty and the parsed value is empty. |
| `invalid_field_value` | `error` | `subject`, `field`, `value`, `location`; optional `target` | A present value violates its applicable grammar, vocabulary, kind-specific rule, or relation-value form. |
| `missing_required_section` | `error` | `subject`, `field`, `location` | An applicable authority requires a narrative section that is absent. |
| `empty_required_section` | `error` | `subject`, `field`, `value`, `location` | An applicable authority requires a present narrative section to be non-empty and its content is empty. |
| `noncanonical_section_heading` | `info` | `subject`, `field`, `value`, `location` | A case-only heading variant provides repair guidance and does not suppress a required-section error. |
| `relation_target_unavailable` | authority-dependent; normally `error` | `subject`, `field`, `target`, `location`; exactly one of a non-conflict `target.lookup_state` or `conflict: current_identity` | A relation value accepted by current or legacy lookup grammar has no selectable target. The lookup state or current identity conflict preserves the cause. |
| `relation_target_kind_mismatch` | `error` | `subject`, `field`, `target`, `location` | A uniquely found target has a kind outside `target.expected_kinds`. |
| `relation_reciprocity_mismatch` | `error` | `subject`, `field`, `target`, `location` | A PRODUCT-owned reciprocal relation does not contain the required counterpart. |
| `current_identity_conflict` | `error` | `subject: record`, `conflict: current_identity` | Current sources claim one canonical identity and no winner exists. |
| `legacy_lookup_conflict` | `error` in repository validation | `subject: record`, `field`, `target`, `conflict: legacy_lookup`, `location` for the referring source | A selected current record relation cannot resolve because the accepted legacy issued ID has multiple configured candidates. |

`empty_required_field` and `empty_required_section` do not mean that every required field or section must be non-empty.
They fire only when the applicable PRODUCT authority requires non-empty content.
A field or section whose authority permits an empty value produces no empty-required diagnostic.

A noncanonical or unsupported declared relation value that matches no accepted current or legacy lookup grammar must use `invalid_field_value`.
No lookup runs for that condition, no `target.lookup_state` is attached, and `relation_target_unavailable` must not be emitted for the same condition.
`relation_target_unavailable` applies only after accepted current or legacy lookup grammar and no usable target is selected.

### Operation warning and resolver-diagnostic categories

| category | normal severity | surfaces | required associations | trigger mapping |
|---|---|---|---|---|
| `missing_compact_field` | `warning` | `list_records` | `subject: record`, `field`; `location` only when T04 allows it | A returned compact record lacks `title`, `status`, or `date` and W004 projects `null`. |
| `malformed_requested_ref` | `warning` | `get_records`, `resolve_reference` | `subject: request_item`; `occurrence` for `get_records`, otherwise `value.actual` | The owning operation classifies the supplied string as malformed. |
| `unsupported_requested_ref` | `warning` | `get_records`, `resolve_reference` | `subject: request_item`; `occurrence` for `get_records`, otherwise `value.actual` | The supplied string belongs to no supported operation input family. |
| `unresolved_requested_ref` | `warning` | `get_records`, `resolve_reference` | `subject: request_item`, `target`; `occurrence` for `get_records`; optional `conflict` | An accepted exact current, spec, or legacy input has no selectable target. |
| `legacy_lookup_unavailable` | `warning` | `get_records`, `resolve_reference` | `subject: request_item`, `target`; `occurrence` for `get_records`; `location` when T04 allows and a source exists | Accepted legacy lookup is disabled or a unique indexed source is unreadable. |
| `legacy_lookup_conflict` | `warning` in read operations | `get_records`, `resolve_reference` | `subject: request_item`, `target`, `conflict: legacy_lookup`; `occurrence` for `get_records` | Accepted legacy lookup has multiple configured candidates. |
| `duplicate_requested_ref` | `info` | `get_records` | `subject: request_item`, `occurrence` | Later exact-equal occurrences are ignored after the first occurrence. |

The operation contracts own trigger conditions, response wrappers, and placement.
This contract owns the entry representation and category names.

For an accepted current input whose active-index identity is conflicted, `unresolved_requested_ref` may include `conflict.type: current_identity` so the cause remains machine-readable without adding another public resolver status.

### Resolver status alignment

`resolve_reference` keeps the W005 public statuses `resolved`, `unresolved`, and `unsupported`.
It returns a top-level `diagnostics` array in every normal resolver response.

| resolver state | diagnostic result |
|---|---|
| Current or legacy target resolves. | `diagnostics: []`. |
| String is malformed under operation classification. | One `malformed_requested_ref` warning. |
| Neither current nor accepted legacy grammar supports the string. | One `unsupported_requested_ref` warning. |
| Accepted current input has no active-index target. | One `unresolved_requested_ref` warning with `target.lookup_state: current_unresolved`. |
| Accepted current identity is conflicted. | One `unresolved_requested_ref` warning with `conflict.type: current_identity`. |
| Accepted legacy input has fallback disabled. | One `legacy_lookup_unavailable` warning with `target.lookup_state: legacy_disabled`. |
| Accepted legacy input has no candidate. | One `unresolved_requested_ref` warning with `target.lookup_state: legacy_unresolved`. |
| Accepted legacy input has duplicate candidates. | One `legacy_lookup_conflict` warning. |
| Accepted legacy input has one unreadable indexed source. | One `legacy_lookup_unavailable` warning with `target.lookup_state: legacy_unreadable`. |

The diagnostic cause never changes the resolver status or successful target projection.

### PRODUCT semantic authority mapping

DRMCP categories do not define which fields, statuses, sections, or relations are valid.
They consume the following authorities:

| semantic concern | authority |
|---|---|
| Current canonical lookup, duplicate identity, and declared relation invalidity | `spec:product.design_records.traceability.resolve_and_validation` |
| Investigation and workflow relation fields | `spec:product.design_records.traceability.metadata_schema` |
| Current sequential artifact identity | `spec:product.design_records.namespace_model.artifact_id_grammar` |
| Current spec document shape | `spec:product.design_records.spec_format.document_shape` |
| Current spec path-derived identity | `spec:product.design_records.spec_format.spec_id_as_ref` |
| Migration-sensitive spec validation policy | `spec:product.design_records.spec_format.validation_policy` |
| ADR metadata, lifecycle, and narrative rules | `spec:product.design_records.authoring_standards.adr_authoring` |
| Spec author-facing rules | `spec:product.design_records.authoring_standards.spec_authoring` |
| Investigation metadata, lifecycle, and narrative rules | `spec:product.design_records.authoring_standards.investigation_authoring` |
| Requirement metadata, lifecycle, and narrative rules | `spec:product.design_records.authoring_standards.requirement_authoring` |
| Work-item metadata, lifecycle, and narrative rules | `spec:product.design_records.authoring_standards.work_item_authoring` |
| Task metadata, lifecycle, and narrative rules | `spec:product.design_records.authoring_standards.task_authoring` |

The category remains stable when two kinds share the same invalidity cause.
Kind-specific requiredness, allowed emptiness, gated-section activation, and accepted values are represented by the cited authority plus `field`, `value`, and `target`, not by creating one category per PRODUCT rule.

### Authoring shared-envelope boundary

Proposal-local `validation.diagnostics` uses the same repository category, severity, association, ordering, and duplicate-suppression rules as `validate_records` for the same candidate or materialized state.

Authoring operation diagnostics also use required `category`, `severity`, and `message`, plus applicable shared associations.
Existing authoring category names and trigger semantics remain owned by the authoring operation contracts.

Authoring-only category fields remain allowed when they express operation concepts that do not fit the shared associations, including:

- `proposal_id`;
- `body_cache_id`;
- `candidate_headings`;
- `stripped_heading`;
- `stripped_level`.

This shared-envelope alignment does not change:

- proposal lifecycle or retention;
- body-cache behavior;
- proposal creation, accept, discard, or retrieval behavior;
- stale, target-change, or ID-collision checks;
- write eligibility or post-write repair behavior;
- authoring response wrappers.

Authoring diagnostic entries do not use scalar `record_id`, scalar `field`, scalar `value`, or standalone raw diagnostic `path` after this contract is applied.

### Deterministic ordering

Repository and proposal-local validation diagnostics use this primary subject-type order:

1. `configuration`;
2. `source`;
3. `record`;
4. `request_item`.

Subjects then sort as follows:

| subject type | ascending key |
|---|---|
| `configuration` | `app_namespace`, then `component` |
| `source` | `app_namespace`, then `record_kind`, then the T04 stable location sort key |
| `record` | canonical `ref` |
| `request_item` | `occurrence.first_index` when present |

Within one subject, diagnostics sort by:

1. `category`;
2. `field.name`;
3. numeric `field.item_index`;
4. `target.ref`;
5. `target.lookup_state`;
6. `target.reciprocal_field`;
7. `conflict.type`.

Missing optional keys sort before present values for the same preceding key.
`message` is never a sort key.

W004 operation-specific ordering remains authoritative:

- `get_records` warnings follow first-occurrence request order;
- duplicate-request entries use ascending `occurrence.first_index`;
- `list_records` missing compact-field warnings follow returned-record order and then `title`, `status`, `date`.

### Duplicate suppression

Duplicate suppression occurs before final deterministic sorting.

Semantic diagnostic identity consists of:

- `category`;
- canonical `subject`;
- `field`, including `item_index`;
- `target`, including `ref`, `lookup_state`, and `reciprocal_field`;
- requested `ref` plus `occurrence.first_index` when present;
- `conflict.type` when present;
- source or finding location identity when the diagnostic is location-specific.

The following do not define identity:

- `message`;
- `severity`;
- `value.actual` or `value.expected`;
- counts;
- aggregated duplicate indexes;
- aggregated conflict-member or candidate-location lists;
- the internal check path that detected the condition.

The same semantic finding discovered by multiple checks is emitted once.
Different subjects, fields, item indexes, targets, lookup states, reciprocal fields, conflict types, or finding locations remain distinct.

Aggregation rules:

- duplicate requested refs aggregate every later index into one ordered `duplicate_indexes` list;
- one current identity conflict aggregates every conflicting current source location;
- one legacy lookup conflict aggregates every configured candidate location for the same lookup.

If the same semantic identity is produced with different severities, the implementation violates the deterministic severity contract and must not silently select one.

### Removed stale behavior

This contract does not define or retain:

- resolver statuses named `unresolved_reference`, `ambiguous_reference`, or `unsupported_reference`;
- `get_records` request field `ids`, item-level failure wrappers, or `record_not_found`;
- `requested_id` or `record_id` diagnostic shortcuts;
- one generic duplicate category that merges current identity conflict and legacy lookup conflict;
- `spec_status_mismatch` based on obsolete YAML or nested metadata;
- front-matter `semantic_refs`, `sections`, section-target, or section-alias diagnostics;
- DRMCP-owned copies of PRODUCT required-section or lifecycle rules;
- raw path as a mandatory standard diagnostic field;
- V01 ADR or requirement citations as current diagnostic authority.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.tools.validate_records` | Validation request, subject selection, relation lookup, and normal wrapper. |
| `spec:drmcp.design_records_mcp.tools.resolve_reference` | Three-status resolver response plus W006 cause diagnostics. |
| `spec:drmcp.design_records_mcp.tools.list_records` | Compact-list warning triggers and placement. |
| `spec:drmcp.design_records_mcp.tools.get_records` | Exact-retrieval warning triggers, ordered occurrence behavior, and placement. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Current source retention, addressability, and conflict groups. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Current and legacy root state and separate legacy lookup map. |
| `spec:product.design_records.traceability.resolve_and_validation` | PRODUCT-owned relation and identity invalidity. |
| `spec:product.design_records.authoring_standards` | PRODUCT-owned kind-specific authoring rules. |

## Sources

- `DRMCP-TASK-MCP-006-01`: Authority, contradiction, and changed-file baseline.
- `DRMCP-TASK-MCP-006-02`: Validation execution, subject, relation-lookup, and response-wrapper contract.
- `DRMCP-TASK-MCP-006-03`: Accepted D01 through D11 diagnostic representation decisions.

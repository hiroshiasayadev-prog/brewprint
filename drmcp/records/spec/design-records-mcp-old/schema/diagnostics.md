# Reference: Diagnostics

- **id**: `spec:drmcp.design_records_mcp.schema.diagnostics`
- **status**: draft
- **date**: 2026-06-30
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

### Internal producer boundary

Core validators return transport-neutral findings as data. Validators do not format MCP responses and perform no filesystem I/O.

Application use cases aggregate findings, suppress semantic duplicates, apply deterministic ordering, and project operation diagnostics or warnings.

The MCP adapter encodes the application result. The adapter does not redefine category, severity, association, ordering, or execution-error meaning.

A failure that prevents a trustworthy diagnostic collection is an application execution error. It is not converted into a normal diagnostic entry.

The complete internal architecture is defined by `spec:drmcp.implementation`.

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
| `location` | category-dependent | Portable machine-readable source location for a repairable source-backed finding. Its concrete shape and exposure rules are defined below. |

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
An identity-less source omits `ref`; its exact repair target is supplied by the required `location` association. If that location cannot be constructed, the operation fails under the missing-required-location boundary below.

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
- each member contains the concrete `location` association defined below;
- no source winner is selected.

For `legacy_lookup`:

- `candidate_count` is the number of configured legacy candidates;
- `candidates` contains one entry per conflicting legacy source;
- each candidate contains the concrete `location` association defined below;
- no candidate winner is selected;
- legacy archive sources remain lookup targets and do not become repository-validation subjects.

Disabled, missing, unreadable, resolved, and unsupported legacy lookup outcomes are not conflict objects.

### Source-location association

`location` identifies one repairable current or legacy source through one shared object shape.

Current example:

```json
{
  "source_scope": "current",
  "app_namespace": "drmcp",
  "records_root": "drmcp/records",
  "path": "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-04-example.md"
}
```

Legacy example:

```json
{
  "source_scope": "legacy",
  "records_root": "v01/records",
  "path": "v01/records/tasks/mcp/V01-TASK-MCP-003-01-example.md"
}
```

| field | required | meaning |
|---|---:|---|
| `source_scope` | yes | `current` or `legacy`. |
| `records_root` | yes | Configured repository-relative records root containing the source. |
| `path` | yes | Repository-relative source file path. |
| `app_namespace` | current only | Explicit app namespace bound to the configured current root. Prohibited for legacy locations. |

Portable path rules:

- `records_root` and `path` are relative to the configured `repository_root`;
- `/` is the only exposed separator;
- leading slash, trailing slash, duplicate separator, empty segment, `.`, and `..` are prohibited;
- Windows drive-qualified, UNC, device, URI, `file://`, and other absolute forms are prohibited;
- path spelling is preserved without case folding, Unicode normalization, or locale-dependent rewriting;
- `records_root` must remain canonically within `repository_root`;
- `path` must remain canonically within both `repository_root` and its declared `records_root`;
- symlink, junction, reparse-point, or other alias escape cannot produce a valid normal location.

Location semantic identity is:

```text
(source_scope, records_root, path)
```

`app_namespace` is required current-root context but is not another identity component.

The stable location sort key is:

```text
(source_scope_rank, records_root, path)
```

where `current = 0` and `legacy = 1`.
Normalized strings compare in locale-independent UTF-8 bytewise ascending order without case folding.
Configuration order, scan order, discovery order, host filesystem collation, and absolute paths are never identity or sort inputs.

### Location exposure by surface

A direct `location` identifies the one source that contains or causes a finding.
Conflict-source locations appear only in their typed conflict collection and are not duplicated at diagnostic top level.

| surface or outcome | exposure |
|---|---|
| Source-backed `validate_records` diagnostic | Direct `location` required when one current source is the repair target. |
| Current identity conflict | Every `conflict.members[].location` required; no direct diagnostic location. |
| Repository-validation legacy relation conflict | Direct location of the referring current source plus every `conflict.candidates[].location`. |
| Proposal-local validation | Direct current location required when a deterministic candidate repository target exists, including an unmaterialized create target. |
| Persisted-file authoring diagnostic | Direct current location allowed and required when the file is the cause or repair target. |
| `list_records` `missing_compact_field` | Direct current source location required. |
| Current conflict in `get_records` or `resolve_reference` | Every conflict member location required; no direct diagnostic location. |
| Unique unreadable legacy source | Direct legacy source location required. |
| Legacy lookup conflict in a read operation | Every conflict candidate location required; no direct diagnostic location. |
| Malformed, unsupported, duplicate, source-less unresolved, disabled-lookup, proposal-lifecycle, or body-cache-lifecycle condition | Location prohibited. |
| Successful list, retrieval, resolver, or normal target projection | Location prohibited. |

Paths do not move into `subject`, `field`, `value`, `target`, successful records, or successful resolver targets.
Normal successful list, retrieval, and resolver projections remain path-free.

### Missing required location

A source-backed entry is valid only when every required direct, member, or candidate location can be constructed.

When a required location cannot be constructed:

- do not emit the entry without it;
- do not drop an affected item, member, or candidate;
- do not return a partial location or conflict collection;
- do not substitute an opaque source token or absolute path;
- do not weaken category or severity to continue;
- fail the operation before emitting a normal response or beginning a write.

For `validate_records`, this is an execution failure and no normal validation wrapper is returned.
For `list_records`, `get_records`, and `resolve_reference`, this is operation execution failure rather than per-item partial success or an unresolved status.
Proposal creation retains no proposal.
`accept_proposed_write` begins no write and returns `written: false` with an empty `files_written` list when the invariant fails before writing.
An implementation failure discovered after files were actually modified must not misreport the write as absent.

No shared “location unavailable” diagnostic is introduced because it would itself lack the required repair target.

### Absolute physical-path boundary

No current operation or shared diagnostic field exposes an absolute physical path.

A future separately contracted and host-enabled privileged debug or emergency operation may expose a distinct `physical_path` field.
It must not replace portable `location`, add a hidden debug flag to an existing operation, or use the absolute value for identity, ordering, or duplicate suppression.
That future operation must define its request, response, privilege, scope, and failure or redaction behavior.

Authoring `target.path`, `diff.files[].path`, unified-diff path operands, and `files_written[].path` are separate explicit transaction outputs governed by the authoring transaction contract; they are not diagnostic `location` objects.

### Severity and validation blocking

| severity | meaning |
|---|---|
| `error` | Applicable repository or record contract is violated. Blocks `validate_records.ok`. |
| `warning` | Operation completed normally but one item failed, a projection degraded, or repair attention is warranted without invalidating the entire response wrapper. |
| `info` | Non-blocking notice, ignored duplicate occurrence, or repair guidance. |

`validate_records.ok` is `false` when at least one returned diagnostic has severity `error`.
It is `true` when no returned diagnostic has severity `error`, including responses containing only `warning` or `info` entries.

Malformed validation requests, untrustworthy mandatory configuration or index state, and inability to construct a required source-backed location remain request or execution failures outside the normal validation wrapper.

Severity is context-sensitive but deterministic.
The same condition under the same operation and applicable authority always maps to the same severity.

### Repository-validation categories

| category | normal severity | required associations | meaning |
|---|---|---|---|
| `source_unreadable` | `error` | `subject: source`, `location` | A selected current source cannot be read. |
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
| `missing_compact_field` | `warning` | `list_records` | `subject: record`, `field`, `location` | A returned compact record lacks `title`, `status`, or `date` and W004 projects `null`. |
| `malformed_requested_ref` | `warning` | `get_records`, `resolve_reference` | `subject: request_item`; `occurrence` for `get_records`, otherwise `value.actual` | The owning operation classifies the supplied string as malformed. |
| `unsupported_requested_ref` | `warning` | `get_records`, `resolve_reference` | `subject: request_item`; `occurrence` for `get_records`, otherwise `value.actual` | The supplied string belongs to no supported operation input family. |
| `unresolved_requested_ref` | `warning` | `get_records`, `resolve_reference` | `subject: request_item`, `target`; `occurrence` for `get_records`; optional `conflict: current_identity` whose members each require `location` | An accepted exact current, spec, or legacy input has no selectable target. |
| `legacy_lookup_unavailable` | `warning` | `get_records`, `resolve_reference` | `subject: request_item`, `target`; `occurrence` for `get_records`; direct `location` required only for `legacy_unreadable` and prohibited for `legacy_disabled` | Accepted legacy lookup is disabled or a unique indexed source is unreadable. |
| `legacy_lookup_conflict` | `warning` in read operations | `get_records`, `resolve_reference` | `subject: request_item`, `target`, `conflict: legacy_lookup` whose candidates each require `location`; `occurrence` for `get_records` | Accepted legacy lookup has multiple configured candidates. |
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
| `source` | `app_namespace`, then `record_kind`, then `(source_scope_rank, records_root, path)` using the location comparison rules above |
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
- one current identity conflict aggregates every conflicting current source location, sorted by the stable location sort key;
- one legacy lookup conflict aggregates every configured candidate location for the same lookup, sorted by the stable location sort key.

When location participates in semantic diagnostic identity, equality uses the complete `(source_scope, records_root, path)` tuple.

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
| `spec:drmcp.implementation` | Finding production, aggregation, projection, and execution-error ownership. |
| `spec:product.design_records.traceability.resolve_and_validation` | PRODUCT-owned relation and identity invalidity. |
| `spec:product.design_records.authoring_standards` | PRODUCT-owned kind-specific authoring rules. |

## Sources

- `DRMCP-TASK-MCP-006-01`: Authority, contradiction, and changed-file baseline.
- `DRMCP-TASK-MCP-006-02`: Validation execution, subject, relation-lookup, and response-wrapper contract.
- `DRMCP-TASK-MCP-006-03`: Accepted D01 through D11 diagnostic representation decisions.
- `DRMCP-ADR-MCP-004`: Operation result and diagnostic ownership.
- `DRMCP-ADR-MCP-005`: Validation finding aggregation and projection boundary.

# Reference: Diagnostics

- **id**: `spec:drmcp.design_records_mcp.schema.diagnostics`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines all validation diagnostic categories, severity levels, required additional fields, and diagnostic policy for Design Records MCP.

## Current contract

### `resolve_reference` resolution responses

`resolve_reference` returns these response types for direct queries:

| response | condition |
|---|---|
| `unresolved_reference` | Supported ref form but target does not exist |
| `ambiguous_reference` | Ref matches multiple targets; single resolution not possible |
| `unsupported_reference` | Input is defined as out of MVP resolver scope |

Reserved prefix `yaml:` public resolver input / direct query response behavior and investigation metadata validation behavior are not defined in MVP. These resolution responses are distinct from validation diagnostics produced by `validate_records` for reference fields and index defects.

### `get_records` retrieval diagnostics

`get_records` returns these request- and retrieval-level diagnostics:

| category | severity | placement | required additional fields | meaning |
|---|---|---|---|---|
| `record_not_found` | error | item-level | `requested_id` | Requested exact record ID lookup key is not in the index |
| `duplicate_requested_id_ignored` | info | top-level | `requested_id`, `first_index`, `duplicate_indexes` | Second and subsequent occurrences of the same requested ID are ignored; only the first-occurrence item is returned |

`first_index` and `duplicate_indexes` are zero-based indexes into the request `ids` array. One `duplicate_requested_id_ignored` diagnostic is returned per duplicated requested ID. These are request/retrieval diagnostics; they do not use the `record_id` field that indicates a record metadata defect.

### Authoring transaction diagnostics

These diagnostic categories are used in proposal, accept, discard, get-proposal, and body-cache retry responses.

| category | severity | meaning |
|---|---|---|
| `proposal_not_found` | error | Requested proposal ID does not exist |
| `proposal_expired` | error | Requested proposal ID is past expiry |
| `proposal_discarded` | error | Proposal is already discarded; cannot be accepted |
| `proposal_already_accepted` | error | Proposal is already accepted; cannot be applied again |
| `proposal_stale` | error | Proposal base state and current target state do not match |
| `target_changed` | error | Target record kind / path / identity differs from proposal creation time |
| `id_collision` | error | Create proposal resolved ID was claimed before acceptance |
| `required_follow_up_not_satisfied` | error | Required reciprocal metadata update or other follow-up is not satisfied |
| `invalid_body_source` | error | Body source rule violation: both `body` and `body_cache_id` supplied, or required body source is missing |
| `body_cache_not_found` | error | Requested body cache ID does not exist |
| `body_cache_expired` | error | Requested body cache ID is past expiry |
| `proposal_preparation_failed` | error | Proposal preparation failed before proposal persistence |
| `section_selector_no_match` | error | Named section selector matched no sections in the target record |
| `section_selector_ambiguous` | error | Named section selector matched multiple sections; single target cannot be resolved |
| `section_replacement_body_heading_stripped` | warning | The first non-empty line of the `named_section_replace` replacement body was a Markdown ATX heading matching `section_selector`; it was stripped before proposal creation to prevent a duplicate heading |

`section_selector_no_match` / `section_selector_ambiguous` diagnostics should include `candidate_headings` when possible. Candidate heading entries carry at least `heading`, `level`, and `ordinal`.

`section_replacement_body_heading_stripped` diagnostics MUST include `stripped_heading` (the stripped heading text) and `stripped_level` (ATX level as an integer). This diagnostic does not block retained proposal creation.

### `validate_records` diagnostics

Diagnostics fire independently per check axis. Multiple diagnostics may attach to a single record.

| category | severity | meaning |
|---|---|---|
| `duplicate_id` | error | Multiple records share the same normalized record ID |
| `filename_id_mismatch` | error | Canonical ID or metadata ID of a decision / investigation / workflow artifact record does not match the filename ID segment |
| `invalid_h1_title` | error | H1 is absent or does not match the expected format |
| `invalid_workflow_id` | error | Requirement / work item / task metadata ID or H1 ID does not follow workflow ID grammar |
| `missing_required_metadata` | error | Required metadata field is absent for a requirement / work item / task |
| `empty_required_metadata` | error | Required scalar metadata field is empty, or a required list metadata field contains an empty item |
| `missing_required_section` | error | A workflow artifact in a gated status is missing a required narrative section heading |
| `empty_required_section` | error | A workflow artifact in a gated status has the required narrative section heading but the section body is empty or whitespace-only |
| `section_heading_case_mismatch` | info | A heading exists that differs from a canonical required heading only by case |
| `invalid_metadata_value` | error | Required metadata field is non-empty but does not satisfy the value contract (e.g. `date` is not strict `YYYY-MM-DD`) |
| `invalid_status_for_kind` | error | `status` value is not allowed for the record `kind` |
| `spec_status_mismatch` | error | Spec top-level `status` and `design_record.status` do not agree |
| `missing_depends_on_target` | error | A `depends_on` referenced ID does not exist |
| `missing_supersedes_target` | error | A `supersedes` referenced ID does not exist |
| `invalid_migrated_to_spec` | error | `migrated_to_spec` value is invalid |
| `missing_record_path` | error | Discovery found a candidate path but read/stat failed |
| `invalid_semantic_ref_declaration` | error | A spec front matter `semantic_refs` entry or `sections` key does not follow active `spec:` grammar |
| `missing_section_target` | error | A spec front matter `sections` value has no matching Markdown heading |
| `ambiguous_section_target` | error | A spec front matter `sections` value matches multiple headings in the same document; single resolution not possible |
| `duplicate_semantic_ref` | error | An active `spec:` document-level or section-level ref is declared to multiple targets and cannot be resolved to a single one |
| `unresolved_source_ref` | error | An investigation `source_refs` entry is a supported canonical ref but cannot be resolved |
| `unresolved_follow_up_result` | error | An investigation `follow_up_results` entry is a supported canonical ref but cannot be resolved |
| `unresolved_follow_up_candidate` | info | An investigation `follow_up_candidates` entry is a supported canonical ref but is unresolved (planned artifact not yet created) |
| `noncanonical_source_ref` | error | An investigation `source_refs` entry is a physical path rather than a canonical ref |
| `noncanonical_follow_up_result` | error | An investigation `follow_up_results` entry is a physical path rather than a canonical ref |
| `noncanonical_follow_up_candidate` | info | An investigation `follow_up_candidates` entry is a physical path rather than a canonical ref |
| `unsupported_reference` | error / info | MVP-defined unsupported reference in investigation metadata. `source_refs` / `follow_up_results`: error. `follow_up_candidates`: info. Includes `TASK-*` in investigation fields. Reserved `yaml:` prefix is not in scope for this category; its behavior is not defined in MVP. |
| `unresolved_workflow_relation` | error | A workflow relation field entry is a supported `REQ-*` / `WORK-*` / `TASK-*` ref but cannot be resolved |
| `invalid_workflow_relation_target` | error | A workflow relation field contains a target that is not the expected kind or ID form |
| `workflow_relation_mismatch` | error | Declared bidirectional relation between `REQ.work_items` ↔ `WORK.source_requirement`, or `WORK.tasks` ↔ `TASK.work_item` is inconsistent |
| `workflow_source_requirement_mismatch` | error | Task `source_requirement` does not match the parent work item's `source_requirement` |

### Required additional fields by diagnostic type

**Workflow metadata diagnostics** (`missing_required_metadata` / `empty_required_metadata` / `invalid_metadata_value`): in addition to the standard `category` / `severity` / `record_id` / `path` / `message`, must return `field`. When the input value is present, must also return `value`.

**Workflow required section diagnostics** (`missing_required_section` / `empty_required_section`): in addition to standard fields, must return `section` (the required narrative section heading text) and `status` (the workflow artifact status that activated the rule).

**`section_heading_case_mismatch`:** must include `section` (canonical required heading), `actual_heading` (matched non-canonical heading), and `status` (workflow artifact status that activated the required-section rule). Should include `candidate_headings` (heading text, level, ordinal) when available.

This diagnostic is repair guidance only. It does not relax canonical required-section validation and does not suppress any `missing_required_section` / `empty_required_section` error. It is returned when: a canonical required section is missing for the target record kind and current gated status, and exactly one heading exists whose text differs from the canonical heading only by case.

**Workflow relation diagnostics** (`unresolved_workflow_relation` / `invalid_workflow_relation_target` / `workflow_relation_mismatch` / `workflow_source_requirement_mismatch`): in addition to standard fields, must return `field` (one of `work_items` / `source_requirement` / `tasks` / `work_item` / `depends_on`), `value` (the input ID-as-ref), and `ref_status` (one of `unresolved` / `invalid_target` / `mismatch`). When the target ID can be identified, must also return `target_id`.

**Investigation reference diagnostics** (`unresolved_*` / `noncanonical_*` / metadata-field-originated `unsupported_reference`): in addition to standard fields, must return `field` (one of `source_refs` / `follow_up_results` / `follow_up_candidates`), `value` (the input reference string), and `ref_status` (one of `unresolved` / `unsupported` / `noncanonical`). When investigation metadata points to a duplicate semantic ref or duplicate record ID that cannot be uniquely resolved, no field-specific diagnostic is added — only `duplicate_semantic_ref` or `duplicate_id` is returned. These duplicate diagnostics and spec declaration/section lookup diagnostics do not require the additional fields defined for investigation metadata diagnostics.

### Required narrative section policy

| artifact kind | gated status | required non-empty narrative sections |
|---|---|---|
| `work_item` | `done` | `Goal`, `Boundary`, `Evidence` |
| `task` | `done` | `Goal`, `Work`, `Done condition`, `Verification`, `Evidence` |
| `requirement` | `accepted` | `Requirement`, `Required Outcome` |

Only headings listed for the target record kind in this table are canonical workflow required headings for the case-only repair behavior defined by `propose_record_update` named section replacement. The target record does not need to currently be in the gated status for the authoring selector fallback to apply — the record kind and requested heading determine fallback eligibility. Authoring guide format headings not listed for the target kind, and user-defined optional headings, are not canonicalized by this rule.

`requirement` `accepted` is treated as an adoption-readiness gate, not a close/completion state. Therefore `Evidence` / `Boundary` / `Explicitly Excluded Scope` are not required non-empty sections for `REQ accepted`.

Required narrative section body is non-empty when: the section body (excluding the heading line) trimmed of leading and trailing whitespace contains at least one non-whitespace character. Whitespace-only body is empty. Body quality, sufficiency, and semantic content are not evaluated — placeholder text such as `Pending` or `None` is non-empty.

### Diagnostic policy and MVP exclusions

`follow_up_candidates` pointing to a planned-but-uncreated artifact is not an error. `validate_records.ok` is determined solely by the presence of error-severity diagnostics; info diagnostics do not cause failure.

The following are **not** included in MVP diagnostic scope:

- Coverage mapping, semantic realization relation, `internal-design:` / `coverage:` / `COV-*` resolution and diagnostics.
- Workflow relation validation beyond declared relation existence and consistency: no orphan artifact diagnostics, no task dependency cycle detection, no execution order projection, no task status–derived progress projection.
- `accepted_but_not_migrated`
- `missing_design_record`
- Status combination validity
- Per-spec-section origin insufficiency
- Semantic mismatch between natural language body and metadata

`missing_record_path` is issued when the filesystem scan or path normalization detects a record candidate path but the actual read/stat fails. Examples include a file deleted after scan, permission denied, symlink target missing, or a path normalization result that does not exist.

> Source: V01-ADR-077 §validate_records の責務, V01-ADR-090 §Partial result / Ordering と duplicate requested ID, V01-ADR-092 §4–§7, V01-REQ-MCP-017 / V01-TASK-MCP-016-01 required narrative section policy

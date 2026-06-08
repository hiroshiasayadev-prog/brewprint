# V01-INV-DATA-002: UC-002 notes retreat inventory for M15 boundary assessment

- **status**: concluded
- **date**: 2026-05-27
- **trigger**: M15 close boundary 判断前に、現行 YAML の `note` 退避実態を inventory 化するため
- **scope**: UC-002 / self-hosting 由来 YAML と関連 YAML に残る schema / contract / relation / type constraint の `note` 退避箇所の抽出と分類
- **non_scope**: ADR 採否判断、M15 release boundary 確定、新規設計提案、REQ / WORK / TASK / ADR 起票、spec / implementation / fixture / YAML 変更、git commit
- **source_refs**:
  - V01-INV-DATA-001
  - V01-ADR-067
  - V01-ADR-069
  - V01-ADR-070
  - V01-ADR-073
  - V01-ADR-078
  - V01-ADR-079
  - V01-ADR-080
- **follow_up_candidates**:
  - なし

# UC-002 notes retreat inventory for M15 boundary assessment

## 1. Scope checked

| path / area | checked? | relevant YAML found? | notes |
|---|---:|---:|---|
| `docs/uc/002-brewprint-self-hosting/**/*.yaml` | yes | yes | Primary target. 37 YAML files checked, including MCP task / model / store / view / render index files. |
| `docs/uc/002-brewprint-self-hosting/docs/coverage.md` | yes | n/a | Used only as a clue for known v1 expressiveness gaps. |
| `docs/uc/002-brewprint-self-hosting/docs/phase-a-work-split.md` | yes | n/a | Used only as a clue for the intended `any + note` / `str + note` fallback policy. |
| `docs/uc/002-brewprint-self-hosting/TASKS-UC-002.md` | yes | n/a | Used only as a clue for known self-hosting schema gaps. |
| `docs/uc/001-ec-checkout-flow/**/*.yaml` | yes | yes | Secondary related YAML under `docs/uc`; only representative enum-like schema debt was inventoried. Most notes were human-facing labels or already represented by `pk` / `fk`. |
| `docs/tasks/m14-self-hosting.md` | yes | n/a | Used as background for UC-002 scope and self-hosting representation choices. |
| `docs/tasks/m15-data-layer-expressiveness.md` | yes | n/a | Used as background for M15 / Phase C terminology only. |
| `yaml/**` | yes | no | Directory does not exist in this repository. |
| `examples/**` | yes | no | Directory does not exist in this repository. |
| `fixtures/**` | yes | no | Directory does not exist in this repository. |
| `testdata/**` | yes | no | Directory does not exist in this repository. |
| `docs/investigations/data/INV-DATA-001-m15-data-layer-expressiveness-derivation-and-completion-boundary.md` | yes | n/a | Used to confirm the M15 / Phase C / V01-ADR-078..080 background; not re-evaluated. |

## 2. Extraction method

Searches and checks used:

- `rg --files` over `docs/uc`, `yaml`, `examples`, `fixtures`, `testdata`, and the M14 / M15 task documents.
- Text searches for `note:`, `type: any`, `type: str`, `element: any`, `value: any`, `list<any>`, `dict<any>`, `kind`, `severity`, `direction`, `discriminator`, `enum`, `値集合`, `FileID`, `QualifiedID`, `synthetic`, `transition`, and `state_file`.
- UTF-8 reads of the relevant Markdown and YAML files.

Judgment rules:

- Included when `note` carried information that could plausibly be represented as a machine-readable schema / contract / relation / type constraint.
- Included when `any` hid object shape, list entry shape, dict / map shape, union, recursive shape, or tool payload shape.
- Included when `str` hid a closed or semi-closed value set such as `out / in / both`, `file / error`, object kinds, diagnostic severity, or error codes.
- Included when `note` carried identity semantics such as `FileID`, `QualifiedID`, `synthetic ID`, file-local IDs, or semantic object registries.
- Excluded when `note` was only a label, prose explanation, renderer background, UI copy, operational description, or already represented by structured YAML fields such as `pk` / `fk`.
- For repeated task-level notes that restated the model-level contract, the inventory records the model field or model note as the owner and does not duplicate every task note.

## 3. Inventory

| ID | YAML file | object / field path | current representation | note-retreated meaning | primary category | primary candidate capability | secondary candidates | confidence | evidence snippet summary |
|---|---|---|---|---|---|---|---|---|---|
| N-001 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_request.yaml` | `analyze_impact_request.change` | `any` | `change.kind` discriminates payload rules for `rename`, `remove`, `change_type`, `change_contract`, `change_transition_target`, and `add`. | `tagged_union_candidate` | `V01-ADR-073 tagged union model` | `not covered by existing ADRs` for cross-field payload requirements | high | Field note says discriminator object; model note lists kind-specific required payloads. |
| N-002 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_request.yaml` | `analyze_impact_request.scope_modules` | `any` | String array for module scope filter. | `opaque_container_shape` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | `V01-ADR-070 file-private helper model` | high | Note says module list / string array represented as `any`. |
| N-003 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml` | `analyze_impact_response.change` | `any` | Echoes the same discriminated change object from input. | `tagged_union_candidate` | `V01-ADR-073 tagged union model` | `not covered by existing ADRs` | high | Note says discriminated object cannot be represented in v1 model. |
| N-004 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml` | `analyze_impact_response.summary` | `any` | Count dictionaries keyed by severity, fixability, and kind. | `dict_key_semantics` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | `V01-ADR-070 file-private helper model`, `V01-ADR-067 enum model` | high | Note names `by_severity / by_fixability / by_kind` dict shape. |
| N-005 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml` | `analyze_impact_response.impacts` | `any` | Impact entry array with fields such as `id`, `kind`, `severity`, `fixability`, `object`, `reason`, `via`, `source`, `recommended_action`, and `suggested_fixes`. | `opaque_container_shape` | `V01-ADR-070 file-private helper model` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption`, `V01-ADR-067 enum model`, `V01-ADR-073 tagged union model` | high | Note lists entry fields and says no dedicated list model. |
| N-006 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml` | `analyze_impact_response.coverage` | `any` | Coverage object has `analyzed`, `not_analyzed`, and `note`; nested vocabularies are kept in note. | `named_or_helper_shape_candidate` | `V01-ADR-070 file-private helper model` | `V01-ADR-067 enum model` | high | Note names fields and model note says coverage lists have standard vocabulary. |
| N-007 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml` | `analyze_impact_response.assumptions` | `any` | String array for tool assumptions / limits. | `opaque_container_shape` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | `V01-ADR-070 file-private helper model` | high | Note says string array represented as `any`. |
| N-008 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml` | `analyze_impact_response.truncated_reasons` | `any` | String array of truncation reasons. | `opaque_container_shape` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | `V01-ADR-070 file-private helper model` | high | Note says reason string array represented as `any`. |
| N-009 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/diagnostic.yaml` | `diagnostic.related` | `any` | Union list of `SourceLocation` or `ObjectRef`. | `constraint_not_covered` | `not covered by existing ADRs` | `V01-ADR-073 tagged union model` | high | Note says union list cannot be represented in v1. |
| N-010 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_request.yaml` | `get_reference_tree_request.direction` | `str` | Required enum: `out / in / both`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `none` | high | Note explicitly says enum constraint is kept in note. |
| N-011 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_request.yaml` | `get_reference_tree_request.depth` | `int` | Numeric range `0..4`; invalid values produce `invalid_depth`. | `constraint_not_covered` | `not covered by existing ADRs` | `none` | high | Note and model note carry range constraint. |
| N-012 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_request.yaml` | `get_reference_tree_request.kinds` | `any` | String array of `Reference.kind` filters. | `opaque_container_shape` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | `V01-ADR-067 enum model`, `V01-ADR-070 file-private helper model` | high | Note says string array and references external Reference.kind vocabulary. |
| N-013 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml` | `get_reference_tree_response.direction` | `str` | Actual direction enum: `out / in / both`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `none` | high | Note lists the value set. |
| N-014 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml` | `get_reference_tree_response.nodes` | `any` | Node entry array; each entry has `object:ObjectRef`, `depth:int`, and `via:string[]`. | `opaque_container_shape` | `V01-ADR-070 file-private helper model` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | high | Note lists entry fields and says no dedicated list model. |
| N-015 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml` | `get_reference_tree_response.edges` | `any` | Reference entry array with added `depth:int`. | `opaque_container_shape` | `V01-ADR-070 file-private helper model` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | high | Note says `Reference` plus depth cannot be represented with existing `reference_list`. |
| N-016 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml` | `get_reference_tree_response.truncated_reasons` | `any` | String array of max-node / max-edge truncation reasons. | `opaque_container_shape` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | `V01-ADR-070 file-private helper model` | high | Note says string array represented as `any`. |
| N-017 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_request.yaml` | `get_references_request.direction` | `str` | Optional enum: `out / in / both`, default `out`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `constraint_not_covered` for default value | high | Note lists value set and default. |
| N-018 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_request.yaml` | `get_references_request.kinds` | `any` | String array of `Reference.kind` filters. | `opaque_container_shape` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | `V01-ADR-067 enum model`, `V01-ADR-070 file-private helper model` | high | Note says string-list filter and model note points to Reference.kind vocabulary. |
| N-019 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_response.yaml` | `get_references_response.direction` | `str` | Actual direction enum: `out / in / both`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `none` | high | Note lists value set. |
| N-020 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_signature_request.yaml` | `get_signature_request.selector` | `object_selector` + model note | Selector valid `object / kind` combinations are delegated to external support matrix. | `constraint_not_covered` | `not covered by existing ADRs` | `V01-ADR-067 enum model` | medium | Model note says valid selector combinations follow MCP schema support matrix. |
| N-021 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_signature_response.yaml` | `get_signature_response.signature` | `any` | Kind-specific signature union for task / model / store / event / state / transition / field. | `tagged_union_candidate` | `V01-ADR-073 tagged union model` | `V01-ADR-070 file-private helper model` | high | Note says kind-specific union cannot be represented in v1. |
| N-022 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_request.yaml` | `get_source_request.fallback` | `str` | Optional enum `file / error`, with default-equivalent behavior for omitted value. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `constraint_not_covered` for default / behavior coupling | high | Field and model notes list values and branch behavior. |
| N-023 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_response.yaml` | `get_source_response.snippet` | `any` | Small object with `language: yaml` and `text`. | `named_or_helper_shape_candidate` | `V01-ADR-070 file-private helper model` | `V01-ADR-067 enum model` | high | Note names the object fields and literal language value. |
| N-024 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_response.yaml` | `get_source_response.fallback` | `str` | Fallback marker enum; current note only names `file`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `unknown` | medium | Related request field defines `file / error`; response field note says fallback case is `file`. |
| N-025 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/inspect_request.yaml` | `inspect_request.detail` | `str` | Optional enum `brief / normal / full`, default `normal`, unknown values are errors. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `constraint_not_covered` for default / unknown-value rule | high | Field and model notes list enum and default. |
| N-026 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/inspect_response.yaml` | `inspect_response.signature` | `any` | `get_signature`-like kind-specific response union. | `tagged_union_candidate` | `V01-ADR-073 tagged union model` | `V01-ADR-070 file-private helper model` | high | Note says shape differs by task / store / model / state / event / view / file. |
| N-027 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/inspect_response.yaml` | `inspect_response.members` | `any` | Kind-specific member payloads such as flow entries, view aggregation, sequence hints, and nested elements. | `tagged_union_candidate` | `V01-ADR-073 tagged union model` | `V01-ADR-070 file-private helper model`, `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | high | Model note says `members` differs by inspected object kind. |
| N-028 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_endpoints_request.yaml` | `list_endpoints_request.api_table_id` | `str` + model note | If omitted and multiple API Tables exist, response groups results into `tables[]`. | `constraint_not_covered` | `not covered by existing ADRs` | `V01-ADR-070 file-private helper model` | medium | Model note carries cross-field / cross-response behavior. |
| N-029 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_endpoints_response.yaml` | `list_endpoints_response.tables` | `any` | Nested list object: table entries have `id`, `http_root_path`, `sections`; endpoint entries have `method`, `path`, `leaf_path`, `task`, optional `params`, `returns`, and `source`. | `opaque_container_shape` | `V01-ADR-070 file-private helper model` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | high | Field and model notes describe nested table / section / endpoint shapes. |
| N-030 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_request.yaml` | `list_objects_request.object` | `str` | Object type filter enum-like values such as `node / view / transition / field`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `V01-ADR-078+ MCP semantic identity series` | medium | Field note and model note treat `object` as enum-like. |
| N-031 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_request.yaml` | `list_objects_request.kind` | `str` | Object-kind filter values such as `task / model / api_table / transition / field`; value set varies by object type. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `constraint_not_covered` for object-dependent vocabulary | medium | Note says object-specific kind filter. |
| N-032 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_request.yaml` | `list_objects_request.file` | `str` | `FileID` based on slash-normalized relative path under `yaml/`. | `identity_or_reference_semantics` | `V01-ADR-078+ MCP semantic identity series` | `not covered by existing ADRs` | high | Note defines FileID path semantics. |
| N-033 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_response.yaml` | `list_objects_response.objects` | `any` | ObjectRef-like array with list-specific summary fields such as `module`, `file`, and `source`. | `opaque_container_shape` | `V01-ADR-070 file-private helper model` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption`, `V01-ADR-078+ MCP semantic identity series` | high | Model note lists `objects[]` fields. |
| N-034 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/mcp_error.yaml` | `mcp_error.code` | `str` | Closed error code vocabulary. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `none` | high | Note lists concrete error codes. |
| N-035 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/common.yaml` | `object_selector.id` | `str` | Selector ID can be `QualifiedID`, actor global ID, or synthetic ID. | `identity_or_reference_semantics` | `V01-ADR-078+ MCP semantic identity series` | `not covered by existing ADRs` | high | Note names multiple identity schemes. |
| N-036 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/common.yaml` | `object_selector.object` | `str` | Optional object enum: `node / view / transition / asset / field / file / primitive`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `V01-ADR-078+ MCP semantic identity series` | high | Note lists enum values. |
| N-037 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/common.yaml` | `object_selector.kind` | `str` | Expected kind filter; value set depends on resolved object type. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `constraint_not_covered` for object-dependent vocabulary | medium | Note says expected kind is validated against resolution result. |
| N-038 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/common.yaml` | `object_selector.file` | `str` | `FileID` used for file-local object selection. | `identity_or_reference_semantics` | `V01-ADR-078+ MCP semantic identity series` | `not covered by existing ADRs` | high | Note names FileID and file-local usage. |
| N-039 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/common.yaml` | `object_selector.local_id` | `str` | File-local object ID for sub task / field / asset. | `identity_or_reference_semantics` | `V01-ADR-078+ MCP semantic identity series` | `not covered by existing ADRs` | high | Note defines file-local object identity. |
| N-040 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/common.yaml` | `object_selector` | all fields optional + model note | Valid selector field combinations vary by tool. | `constraint_not_covered` | `not covered by existing ADRs` | `V01-ADR-067 enum model`, `V01-ADR-078+ MCP semantic identity series` | high | Model note says JSON schema fields are optional but valid combinations are tool-specific. |
| N-041 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/object_ref.yaml` | `object_ref.object` | `str` | Object category enum: `node / view / transition / asset / field / file / primitive`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `V01-ADR-078+ MCP semantic identity series` | high | Note lists enum values. |
| N-042 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/object_ref.yaml` | `object_ref.kind` | `str` | Object-specific kind vocabulary, e.g. node kind values. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `constraint_not_covered` for dependent vocabulary | medium | Note says value set varies per object. |
| N-043 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/object_ref.yaml` | `object_ref.id / qualified_id / file / local_id` | `str` fields | ObjectRef identity can be QualifiedID, synthetic ID, file-local FileID + local ID, or optional qualified ID. | `identity_or_reference_semantics` | `V01-ADR-078+ MCP semantic identity series` | `not covered by existing ADRs` | high | Field notes define multiple identity variants. |
| N-044 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/object_ref.yaml` | `object_ref.parent` | `any` | Recursive parent `ObjectRef` for fields and nested objects. | `constraint_not_covered` | `not covered by existing ADRs` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | high | Field and model notes say recursion cannot be expressed in v1. |
| N-045 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/reference.yaml` | `reference.kind` | `str` | Closed semantic reference kind vocabulary, including model, store, state, event, transition, and scenario refs. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `V01-ADR-078+ MCP semantic identity series` | high | Note lists the full kind vocabulary. |
| N-046 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/reference.yaml` | `reference.direction` | `str` | Reference direction enum: `out / in`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `none` | high | Note lists enum values. |
| N-047 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/resolved_project.yaml` | `resolved_project.semantic_objects` | `any` | Registry of semantic objects such as node, view, transition, asset, field, and file. | `identity_or_reference_semantics` | `V01-ADR-078+ MCP semantic identity series` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | medium | Note defines semantic object registry but says internal shape is not public contract. |
| N-048 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/resolved_project.yaml` | `resolved_project.reference_indexes` | `any` | Maps such as `referencesBySource` / `referencesByTarget`; key and value semantics are internal semantic reference indexes. | `dict_key_semantics` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | `V01-ADR-078+ MCP semantic identity series` | high | Note names map structures and says arbitrary map / union value cannot be represented. |
| N-049 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/resolved_project.yaml` | `resolved_project.render_context` | `any` | Render index group / output mapping used as material for `analyze_impact`. | `opaque_container_shape` | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | `V01-ADR-070 file-private helper model` | medium | Note says mapping shape is not fixed as public schema. |
| N-050 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/source_location.yaml` | `source_location.file` | `str` | `FileID` for source location. | `identity_or_reference_semantics` | `V01-ADR-078+ MCP semantic identity series` | `not covered by existing ADRs` | high | Note says `FileID` is required. |
| N-051 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/string_list.yaml` | `string_list.element` | `list` of `str` + model note | Element enum set depends on usage site, e.g. reference kind filter or coverage vocabulary. | `constraint_not_covered` | `not covered by existing ADRs` | `V01-ADR-067 enum model` | medium | Model note says array element enum values are supplemented per usage note. |
| N-052 | `docs/uc/001-ec-checkout-flow/yaml/cart/model/cart.yaml` | `cart.status` | `str` | Cart status enum-like values: `active / locked / checked_out`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `none` | high | Secondary UC-001 check; note lists state values. |
| N-053 | `docs/uc/001-ec-checkout-flow/yaml/order/model/order.yaml` | `order.status` | `str` | Order status enum-like values: `pending / processing / confirmed / failed`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `none` | high | Secondary UC-001 check; note lists state values. |
| N-054 | `docs/uc/001-ec-checkout-flow/yaml/payment/model/payment_event.yaml` | `payment_event.result` | `str` | Payment result enum-like values: `succeeded / failed`. | `enum_like_closed_vocabulary` | `V01-ADR-067 enum model` | `none` | high | Secondary UC-001 check; note lists result values. |
| N-055 | `docs/uc/002-brewprint-self-hosting/yaml/actors.yaml` | `mcp_client`, `mcp_server` notes | `note only` | Human-facing actor descriptions, not schema debt. | `human_explanation_only` | `none` | `none` | high | Notes describe external MCP client and server boundary. |
| N-056 | `docs/uc/002-brewprint-self-hosting/yaml/views/er.yaml` | `mcp_contract_er.note` | `note only` | View / renderer explanation; no hidden field shape or type constraint. | `human_explanation_only` | `none` | `none` | high | Note explains ER view scope and renderer limitation. |

## 4. Grouped patterns

### 4.1 Enum-like closed vocabulary

| recurring pattern | occurrences | representative inventory IDs | candidate capability |
|---|---:|---|---|
| Direction-like fields encoded as `str + note` | 5 | N-010, N-013, N-017, N-019, N-046 | `V01-ADR-067 enum model` |
| Object / kind / reference vocabularies encoded as `str + note` | 8 | N-030, N-031, N-036, N-037, N-041, N-042, N-045 | `V01-ADR-067 enum model` |
| Tool option / detail / fallback / error code vocabularies encoded as `str + note` | 4 | N-022, N-024, N-025, N-034 | `V01-ADR-067 enum model` |
| Domain status / result values in secondary UC-001 YAML | 3 | N-052, N-053, N-054 | `V01-ADR-067 enum model` |

### 4.2 Named/helper shape candidates

| recurring pattern | occurrences | representative inventory IDs | candidate capability |
|---|---:|---|---|
| Small tool-specific object shape hidden behind `any` | 2 | N-006, N-023 | `V01-ADR-070 file-private helper model` |
| Entry / response object shapes that could become local helper models | 7 | N-005, N-014, N-015, N-029, N-033 | `V01-ADR-070 file-private helper model` |

### 4.3 Tagged union candidates

| recurring pattern | occurrences | representative inventory IDs | candidate capability |
|---|---:|---|---|
| Discriminator-based request / response payload | 2 | N-001, N-003 | `V01-ADR-073 tagged union model` |
| Kind-specific signature / inspect payload | 3 | N-021, N-026, N-027 | `V01-ADR-073 tagged union model` |
| Untagged union list needing a union-like representation | 1 | N-009 | `not covered by existing ADRs`; `V01-ADR-073 tagged union model` may be adjacent only |

### 4.4 Opaque container / dict key semantics

| recurring pattern | occurrences | representative inventory IDs | candidate capability |
|---|---:|---|---|
| String arrays represented as `any` | 6 | N-002, N-007, N-008, N-012, N-016, N-018 | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` |
| Nested list entry shapes represented as `any` | 5 | N-005, N-014, N-015, N-029, N-033 | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption`; `V01-ADR-070 file-private helper model` |
| Dict / map key semantics in `any` | 2 | N-004, N-048 | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` |
| Internal mapping / render context retained as opaque `any` | 1 | N-049 | `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` |

### 4.5 Identity / reference semantics and non-data-layer debt

| recurring pattern | occurrences | representative inventory IDs | candidate capability |
|---|---:|---|---|
| FileID / path-normalized identity semantics | 3 | N-032, N-038, N-050 | `V01-ADR-078+ MCP semantic identity series` |
| QualifiedID / synthetic ID / file-local ID semantics | 4 | N-035, N-039, N-043, N-047 | `V01-ADR-078+ MCP semantic identity series` |
| Semantic reference kind / source-target index semantics | 2 | N-045, N-048 | `V01-ADR-078+ MCP semantic identity series`; partly `V01-ADR-067 enum model` / `V01-ADR-069` |

### 4.6 Not covered by existing ADR candidates

| recurring pattern | occurrences | representative inventory IDs | why existing ADR candidates appear insufficient |
|---|---:|---|---|
| Numeric range / default / unknown-value behavior | 3 | N-011, N-017, N-025 | V01-ADR-067 can model value sets, but not numeric range, defaults, or error behavior by itself. |
| Tool-specific selector combination matrix | 2 | N-020, N-040 | Valid combinations depend on tool / object-kind support matrix, not just enum or helper shape. |
| Cross-response behavior based on omitted input | 1 | N-028 | The `api_table_id` omission changes response grouping behavior. |
| Recursive object references | 1 | N-044 | Recursive type expression is not clearly covered by V01-ADR-067 / 069 / 070 / 073. |
| Usage-site-dependent list element vocabularies | 1 | N-051 | Shared `string_list` cannot carry a single enum; element vocabulary varies by usage site. |
| Untagged union list | 1 | N-009 | V01-ADR-073 is discriminator-oriented; `SourceLocation or ObjectRef[]` has no explicit discriminator in the YAML note. |

## 5. Count summary

| primary category | count | files affected | representative IDs |
|---|---:|---:|---|
| `enum_like_closed_vocabulary` | 19 | 14 | N-010, N-034, N-045, N-052 |
| `named_or_helper_shape_candidate` | 2 | 2 | N-006, N-023 |
| `tagged_union_candidate` | 5 | 4 | N-001, N-021, N-027 |
| `opaque_container_shape` | 12 | 6 | N-005, N-014, N-029, N-033 |
| `dict_key_semantics` | 2 | 2 | N-004, N-048 |
| `identity_or_reference_semantics` | 7 | 5 | N-032, N-035, N-043, N-050 |
| `constraint_not_covered` | 7 | 7 | N-009, N-011, N-040, N-044 |
| `human_explanation_only` | 2 | 2 | N-055, N-056 |
| `unknown` | 0 | 0 | none |

- Total inventory item count: 56, including 2 representative `human_explanation_only` rows.
- Schema-debt / contract-debt item count excluding `human_explanation_only`: 54.
- `any`-family retreat count: 24.
- `str + closed vocabulary` retreat count: 19.
- Items potentially relatable to data-layer ADR candidates V01-ADR-067 / V01-ADR-069 / V01-ADR-070 / V01-ADR-073: 40.
- Items that look like V01-ADR-078..080 semantic identity series or another non-data-layer domain: 7.
- Items that look difficult to handle with the existing ADR candidates alone: 7.

## 6. Capability coverage map — descriptive only

| capability candidate | inventory items potentially related | categories covered | categories not covered / unclear | notes |
|---|---:|---|---|---|
| `V01-ADR-067 enum model` | 19 primary, plus 5 secondary | `enum_like_closed_vocabulary` | dependent vocabularies, default values, unknown-value behavior, usage-site-dependent list element enum | Direct fit for explicit closed vocabularies such as direction, severity, object, reference kind, and error code. |
| `V01-ADR-069 opaque TypeRef warning / inline struct non-adoption` | 15 primary/secondary | `opaque_container_shape`, `dict_key_semantics`, some recursive / map cases | Does not by itself introduce richer machine-readable shapes. | Describes or warns about opacity in list / dict / map shapes; may not remove the retreat. |
| `V01-ADR-070 file-private helper model` | 14 primary/secondary | `named_or_helper_shape_candidate`, some nested entry shapes | Does not directly cover enum, discriminator semantics, identity semantics, or recursive references. | Could potentially name local response entry / snippet / coverage / endpoint shapes. |
| `V01-ADR-073 tagged union model` | 5 primary, plus 1 secondary | `tagged_union_candidate` | Untagged union list, numeric range, selector matrix, identity semantics | Directly adjacent to `change.kind` and kind-specific `signature` / `members` payloads. |
| `V01-ADR-078+ MCP semantic identity series` | 7 primary, plus 7 secondary | `identity_or_reference_semantics` and some semantic reference vocabulary | Does not solve container shape, helper shape, enum, or tagged payload representation. | Covers FileID / QualifiedID / synthetic ID / semantic object identity style questions more than data-layer expressiveness. |
| `not covered by existing ADRs` | 7 primary | `constraint_not_covered` | n/a | Covers numeric ranges, defaults, support matrices, recursive types, cross-response behavior, usage-site-dependent list enum, and untagged unions. |

This section is descriptive only. It does not select M15 scope, release blockers, implementation sequence, or ADR adoption status.

## 7. Inputs for later boundary judgment

- The strongest concentration is `enum_like_closed_vocabulary` (19 rows), followed by `opaque_container_shape` (12 rows).
- The most reusable candidate capabilities appear to be enum modeling for repeated `str + note` vocabularies, and helper / container shape handling for repeated response-entry arrays.
- Patterns not cleanly covered by the existing ADR candidates: numeric range / default / behavior rules, tool-specific selector support matrices, recursive `ObjectRef`, usage-site-dependent list element vocabularies, and untagged union lists.
- Identity / reference semantics are visibly present in YAML notes, but they appear to sit closer to MCP semantic identity / state-machine identity style work than to pure data-layer expressiveness.
- Classification uncertainty remains around object-dependent `kind` vocabularies: they are enum-like, but not globally closed without the object category or support matrix.
- Additional confirmation targets for later judgment: `docs/spec/mcp/schema.md` selector and reference vocabularies, the per-tool specs under `docs/spec/mcp/tools/`, and implementation-side ResolvedProject / ObjectRef identity structures. Those were not used to decide M15 boundary here.
- No ADR adoption, M15 close boundary, release blocker, implementation order, or new artifact proposal is made by this investigation.
- No spec / implementation / fixture / YAML file was changed by this investigation.
- No commit was made.

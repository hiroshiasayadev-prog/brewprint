# Reference: Diagnostics

- **id**: `spec:bpdsl.dsl.diagnostics`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.overview`

## What this is

External-facing specification for brewprint validation diagnostics. Defines the diagnostic object structure, severity levels, text and JSON output formats, sort order, and the complete diagnostic code catalog.

Diagnostics are used in:

- `brewprint validate --yaml-root <path>` text output
- `brewprint validate --yaml-root <path> --format json` JSON output
- `brewprint mcp --yaml-root <path>` — semantic error detection before startup

## Diagnostic object

In JSON output, each diagnostic has the following shape.

```json
{
  "severity": "error",
  "code": "unresolved_model",
  "file": "order/state.yaml",
  "message": "unresolved event payload model: payment_event"
}
```

| field | required | description |
|---|---|---|
| `severity` | yes | `error` or `warning` |
| `code` | no | Machine-readable diagnostic code. May be omitted for unclassified diagnostics. |
| `file` | no | FileID of the YAML file the diagnostic applies to. |
| `message` | yes | Human-readable description. |

## File representation

The `file` field uses FileID.

For files under `yaml/`, the `yaml/` prefix is stripped from the project-root-relative path.

```text
rawyaml.File.Path: yaml/order/state.yaml
rawyaml.File.ID:   order/state.yaml
diagnostic.file:   order/state.yaml
```

For files outside `yaml/` (e.g., `render_index.yaml`), FileID and Path are the same.

```text
rawyaml.File.Path: render_index.yaml
rawyaml.File.ID:   render_index.yaml
diagnostic.file:   render_index.yaml
```

This convention keeps FileID consistent across diagnostic output, MCP responses, and renderer internals.

> Source: V01-ADR-052

## Severity

| severity | meaning |
|---|---|
| `error` | Validation failed. `brewprint validate` returns a non-zero exit code. |
| `warning` | Validation succeeded. Diagnostic is still displayed. |

## Text format

`brewprint validate --format text` outputs one line per diagnostic.

```text
<severity> <code> <file>: <message>
```

Example:

```text
error unresolved_model order/state.yaml: unresolved event payload model: payment_event
```

When there are no diagnostics:

```text
ok
```

## JSON format

`brewprint validate --format json` outputs:

```json
{
  "diagnostics": [],
  "error_count": 0,
  "warning_count": 0
}
```

When error diagnostics exist, JSON is still written to stdout. The command returns a non-zero exit code after the JSON is written.

## Diagnostic ordering

Diagnostics are stable-sorted before being returned from `resolve.Build`.

Sort keys, in order:

1. Severity rank: `error` → `warning` → unknown
2. File ID
3. Code
4. Message

## Semantic reference / model validation

| code | severity | description |
|---|---|---|
| `invalid_model_id` | error | Model ID uses a primitive reserved word. |
| `unresolved_model` | error | Model reference cannot be resolved. |
| `unresolved_field_type` | error | Model field type cannot be resolved as a primitive or defined model. |
| `unresolved_fk` | error | FK reference target field cannot be resolved. |
| `unresolved_store` | error | Store reference cannot be resolved. |
| `invalid_endpoint` | error | Endpoint task method/path definition is invalid. |
| `invalid_store_kind` | error | Store kind is not an allowed value. |
| `invalid_model_kind` | error | Model kind is not an allowed value. |
| `duplicate_model_field` | error | Model field name is duplicated. |
| `duplicate_model_id` | error | Model ID is duplicated within a file, or a public model and a file-private helper model in the same module share the same name. |
| `invalid_private_model_reference` | error | `params[].model` of a task / branch / fork / join references a file-private helper model in the same task file. |
| `duplicate_primary_key` | error | A model has more than one primary key. |
| `missing_required_field` | error | A required field is missing. |
| `invalid_type_ref` | error | TypeRef syntax is invalid, an unsupported container kind is used, or the parser safety limit is exceeded. |
| `invalid_enum_model` | error | `kind: enum` definition is invalid: `values` missing, empty, non-string values, empty string values, or prohibited `fields` / `element` / `value` fields. |
| `duplicate_enum_value` | error | `values` contains duplicates in the same enum model. |
| `invalid_tagged_union_model` | error | `kind: tagged_union` definition is invalid: `discriminator` missing, `variants` missing or empty, or prohibited field presence. |
| `duplicate_variant_tag` | error | `variants[].tag` is duplicated in the same tagged union model. |
| `invalid_variant_field` | error | A variant payload field is invalid: discriminator redefinition, prohibited attributes (`pk`, `fk`, `unique`, etc.), or missing `name` / `type`. |

`invalid_enum_model` is emitted when a `kind: enum` model has a missing or empty `values`, non-string values, empty strings in `values`, or prohibited fields (`fields`, `element`, `value`).

`duplicate_enum_value` is emitted when `values` contains duplicates within the same enum model.

`invalid_enum_value` is not added in v1.1 minimum — brewprint YAML holds schema / model definitions and does not store runtime literal values for enum-typed fields.

`invalid_tagged_union_model` is emitted when: `discriminator` is missing, empty, or a dot path; `variants` is missing, empty, or not a list; or the model contains prohibited fields (`fields`, `element`, `value`, `values`, etc.).

If `variants[].tag` is missing, empty, or not a string, it is treated as `invalid_tagged_union_model`.

`duplicate_variant_tag` is emitted when valid `variants[].tag` values are duplicated. It is not used for missing / empty / non-string tags.

`invalid_variant_field` is emitted for invalid `variants[].fields[]` payload fields. Representative cases: redefinition of the discriminator field name, prohibited attributes (`pk`, `fk`, `unique`, etc.), missing `name` or `type`, or invalid field object shape. It is not used for `variants[].tag` invalidity.

Variant field name duplication may reuse `duplicate_model_field`. Variant field `type` syntax errors use `invalid_type_ref`; unresolved named model uses `unresolved_field_type`. Unsupported tagged union `kind` uses `invalid_model_kind`.

Tagged union definition diagnostics do not validate runtime MCP request / response payloads (discriminator presence, tag value in allowed set, payload field existence, or unknown additional fields) — those are outside V01-WORK-DATA-010 minimum scope.

When a tagged union model is already invalid via `invalid_tagged_union_model`, lower-level checks (variant field TypeRef resolution, `duplicate_variant_tag`, etc.) may be suppressed by cascade. Implementations may still emit additional local diagnostics where feasible, consistent with the policy of not stacking flow-wiring diagnostics on unresolved TypeRefs.

`duplicate_model_id` covers task-file helper model identity validation. Name collision rules: [`spec:bpdsl.dsl.naming`](naming.md) §Helper model name collision rules. Helper model semantics: [`spec:bpdsl.dsl.nodes.data`](nodes/data.md).

When a task-file helper model is not in the TypeRef resolution scope of the reference site, the existing `unresolved_model` or `unresolved_field_type` may be used.

`invalid_private_model_reference` is emitted when a TypeRef resolved to a file-private helper model in the same task file appears in `params[].model` (a public input contract). Since the reference resolved, `unresolved_model` is not stacked on the same reference. Applies to task / branch / fork / join params. `returns.model` referencing a file-private helper model is valid; no diagnostic in minimum scope. When the helper model identity is already invalid (e.g., `duplicate_model_id`), `invalid_private_model_reference` may be suppressed by cascade.

> Source: V01-ADR-067, V01-ADR-070

## TypeRef lint / warning

| code | severity | description |
|---|---|---|
| `opaque_type_ref` | warning | A container TypeRef contains `any` and its shape meaning is not captured by a named model. |

`opaque_type_ref` is treated as validation success. The message may suggest extracting a named model.

`opaque_type_ref` targets `any` inside container TypeRefs as a debt-visibility baseline — it does not warn about or require resolution of bare `any` fields or `any + note` used as main response shapes.

`unclear_dict_key` / `deep_type_ref` are future lint candidates and are not added to v1.1 minimum diagnostics.

> Source: V01-ADR-069

## Duplicate / symbol validation

| code | severity | description |
|---|---|---|
| `duplicate_node` | error | A public node QualifiedID is duplicated project-wide. |
| `duplicate_main_node` | error | More than one main node in a single file. |
| `duplicate_sub_node` | error | A file-private sub-node local ID is duplicated within the same file. |
| `duplicate_actor` | error | Actor ID is duplicated. |
| `duplicate_initialized_store` | error | An initialized store is duplicated within the same file. |

`duplicate_node` is limited to main node collisions on public QualifiedIDs. File-private sub-nodes have no public QualifiedID, so the same local ID in different files is not `duplicate_node`. For same-file sub-node local ID duplication, use `duplicate_sub_node`.

## Flow validation

| code | severity | description |
|---|---|---|
| `unsupported_flow_entry` | warning | Empty or unsupported flow entry. |
| `unresolved_flow_task` | error | Flow step / foreach task cannot be resolved. |
| `unresolved_flow_node` | error | Branch / fork / join node cannot be resolved. |
| `invalid_flow_branch` | error | Step definition inside a fork branch is invalid. |
| `unmatched_join_param` | error | A join param has no matching branch terminal step `returns`. |
| `incompatible_wiring_type` | error | Source TypeRef and target param TypeRef are incompatible. |
| `invalid_wiring_source` | error | Source resolves but cannot be used as a wiring source in this context. |
| `unresolved_wiring_source` | error | Source token resolves to nothing. |
| `invalid_foreach_over_type` | error | `foreach.over` result cannot be treated as a list. |
| `invalid_foreach_returns` | error | `foreach.returns` specified but apply-target has no `returns`, or the foreach's own `params` reference its own `returns` name. |
| `unresolved_return_source` | error | `returns.source` token resolves to nothing. |
| `invalid_return_source` | error | Source resolves but cannot be used as a task return source. |
| `incompatible_return_type` | error | `returns.source` TypeRef and `returns.model` TypeRef are incompatible. |
| `duplicate_flow_source` | error | Same name used by more than one of: node id, `foreach.returns`, `initializes[].name` in the same file. |

`invalid_type_ref` messages include the invalid TypeRef string and its location. Parser safety limit violations also include the TypeRef and the fact that the limit was exceeded.

When TypeRef syntax is valid but the named model is unresolved: use `unresolved_field_type` for `fields[].type`, and `unresolved_model` for `params[].model` / `returns.model` / `model.element` / `model.value`. `invalid_type_ref` is used only when the syntax is broken, an unsupported container kind is used, or the safety limit is exceeded.

`incompatible_wiring_type` messages include the source TypeRef, target TypeRef, and wiring location.

`duplicate_flow_source` is emitted when a node id, a `foreach.returns`-declared collected asset source name, or an `initializes[].name`-declared initialized source name duplicates another in the same flow file. Task `returns.name` is not a bare wiring source, so a `returns.name` matching another bare source name is not `duplicate_flow_source`.

`invalid_foreach_returns` is emitted when:
- The apply-target task has no `returns` but `foreach.returns` is specified.
- The foreach's own `params` reference its own `foreach.returns` name.

`unresolved_wiring_source` vs. `invalid_wiring_source`:
- `unresolved_wiring_source`: reference target does not exist (e.g., typo). Source token resolves to nothing as any of: node id / `$params.<name>` / `$item` / collected asset source / initialized source.
- `invalid_wiring_source`: reference target exists but cannot be used as a wiring source in this context (e.g., a node with no `returns`, or `$item` outside a foreach).

Initialized source is a valid wiring source kind and is not subject to `invalid_wiring_source`. When `initializes[].model` is unresolvable, suppress `incompatible_wiring_type` for wiring that references that initialized source.

Do not stack `incompatible_wiring_type` on top of TypeRef resolution failure, unresolved references, `invalid_foreach_over_type` affecting `$item` wiring, or unresolvable collected asset / initialized source TypeRefs.

`unresolved_return_source` vs. `invalid_return_source`:
- `unresolved_return_source`: `returns.source` token resolves to nothing as any of: node id / `$params.<name>` / collected asset source / initialized source.
- `invalid_return_source`: source resolves but cannot be used as a task return source (e.g., a node with no `returns`, or `$item`).

Initialized source is a valid return source kind and is not subject to `invalid_return_source`.

`incompatible_return_type` messages include the source TypeRef, target TypeRef, and `returns.source` location. Suppress when source or `returns.model` TypeRef is unresolvable.

> Source: V01-ADR-060, V01-ADR-061, V01-ADR-062, V01-ADR-063

## Transition validation

| code | severity | description |
|---|---|---|
| `unresolved_transition_state` | error | Transition `from` / `to` state cannot be resolved. |
| `unresolved_transition_event` | error | Transition `on` event cannot be resolved. |
| `duplicate_transition` | error | Transition `from` / `on` / `guard` combination is duplicated. |
| `missing_transition_guard` | error | A branched transition is missing its guard. |

## View / scenario validation

| code | severity | description |
|---|---|---|
| `duplicate_view` | error | API View / ER View / Sequence Scenario ID is duplicated. |
| `invalid_view_definition` | error | A required view / scenario definition is missing or invalid. |
| `duplicate_view_module` | error | A module definition is duplicated within an API View or ER View. |
| `unresolved_sequence_step` | error | A sequence scenario step cannot be resolved to a transition. |
| `non_continuous_sequence` | error | A sequence scenario step is not continuous from the previous step's `to` state. |

## File classification validation

| code | severity | description |
|---|---|---|
| `unsupported_file` | warning | YAML file could not be classified (`as:` and `nodes:` both absent, and not `render_index.yaml`). Detects forgotten `as:` keys. See [`spec:bpdsl.dsl.file_types`](file-types.md) §Unsupported file handling. |

> Source: V01-ADR-051

## Generic fallback

| code | severity | description |
|---|---|---|
| `semantic_validation` | error / warning | Semantic validation diagnostic not yet classified into a specific code. |

## Compatibility note

Diagnostic `message` is human-facing and may change in future versions. External tools should prefer `severity` / `code` / `file` for programmatic handling.

# Reference: MCP schema

- **id**: `spec:bpdsl.mcp.schema`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:bpdsl.mcp.overview`

## What this is

Common schema shared across MCP tools: selector, ObjectRef, Reference, Diagnostic, and the common ID representations (QualifiedID, FileID, synthetic ID).

> Source: V01-ADR-027, V01-ADR-031, V01-ADR-035, V01-ADR-043, V01-ADR-048, V01-ADR-054, V01-ADR-056, V01-ADR-062

## Object selector

Each tool receives a `selector` that designates the target object.

```json
{
  "selector": {
    "id": "auth.task.login"
  }
}
```

### Selector fields

| field | required | type | content |
|---|---:|---|---|
| `id` | optional | string | Target object ID. Usually a QualifiedID. For `actor`, the global actor ID. For view objects (scenario etc.), a view-specific ID. For transition / asset / private sub-node, use the synthetic ID forms below. Omittable for objects addressable via `file` + `local_id`. |
| `object` | optional | enum | `node` / `view` / `transition` / `asset` / `field` / `file` / `primitive`. Defaults to attempting resolution as `node` when omitted. |
| `kind` | optional | string | Expected kind. Value set depends on `object`. If specified and the resolved kind doesn't match, returns `kind_mismatch` tool error. |
| `file` | optional | FileID | Used to designate a file-local object, e.g. a private sub-node. |
| `local_id` | optional | string | Local object ID within `file`. Used to reference private objects like sub tasks. |

A selector that omits `object` is resolved as `object: node`. Specifying only `kind` does not infer the object class — when querying anything other than `node`, `object` should be specified explicitly.

`invalid_selector` tool error when the selector shape is malformed. `not_found` when it cannot be resolved, `ambiguous` when it resolves to multiple candidates, `kind_mismatch` when the resolved result doesn't match `kind`. When resolution succeeds but the tool doesn't handle that object/kind, follow the selector support matrix in §Selector support matrix below.

A normal externally-referenceable node is designated by `id` alone. To query a private sub-node directly, use a synthetic ID or the `file` + `local_id` form.

```json
{
  "selector": {
    "object": "node",
    "id": "order/task/checkout.yaml#build_order"
  }
}
```

```json
{
  "selector": {
    "object": "node",
    "file": "order/task/checkout.yaml",
    "local_id": "build_order"
  }
}
```

To query an asset directly, use the synthetic ID composed of producer and asset name, or the `id` + `local_id` form.

```json
{
  "selector": {
    "object": "asset",
    "id": "order.task.build_order#draft_order"
  }
}
```

Direct private sub-node queries are handled via `get_signature` / `get_references` / `inspect`. `inspect(main task)`'s `members.sub_tasks` uses the same ObjectRef representation.

### Object-dependent kind vocabulary

The `kind` value set differs per `object`. This table is the shared vocabulary for selector validation and ObjectRef responses.

| object | allowed kind values | notes |
|---|---|---|
| `node` | `task` / `model` / `store` / `state` / `event` / `actor` | brewprint node kind. A private sub-node is also represented as `object: node`, `kind: task`. |
| `view` | `sequence_diagram` / `api_table` / `er_diagram` | View object. The computed route list for API endpoints is `list_endpoints`'s responsibility. |
| `transition` | `transition` | Synthetic transition object within a state file. |
| `asset` | `asset` | Asset object implicitly generated from a task/join's `returns`. |
| `field` | `field` / `model_field` | Model field object. MCP v1 responses may return `field` for backward compatibility; new specs may use `model_field` as the descriptive term. |
| `file` | `node` / `state_file` / `sequence_diagram` / `api_table` / `er_diagram` / `render_index` | File kind under `yaml/`. Represents file kind, not node kind. |
| `primitive` | `primitive` | Primitive type-reference target. Not a direct query target. |

This vocabulary is not a DATA DSL dependent-enum feature — it is treated as the runtime selector contract and response contract on the MCP schema / tool contract.

## Selector support matrix

MCP v1's selector support range is as follows.

| object / kind | [`get_signature`](tools/get-signature.md) | [`get_references`](tools/get-references.md) | [`get_reference_tree`](tools/get-reference-tree.md) | [`analyze_impact`](tools/analyze-impact.md) | [`inspect`](tools/inspect.md) | status |
|---|---:|---:|---:|---:|---:|---|
| `node: task` | yes | yes | yes | yes | yes | supported |
| `node: model` | yes | yes | yes | yes | yes | supported |
| `node: store` | yes | yes | yes | yes | yes | supported |
| `node: state` | yes | yes | yes | yes | yes | supported |
| `node: event` | yes | yes | yes | yes | yes | supported |
| `node: actor` | yes | yes | yes | yes | limited | supported / limited inspect |
| `view: sequence_diagram` | yes | yes | yes | no | yes | supported |
| `transition` | yes | yes | yes | yes | yes | supported |
| `field` / `model_field` | yes | yes | yes | yes | yes | supported |
| `file: node` | no | limited | limited | no | yes | supported / limited references |
| `file: state_file` | no | yes | yes | no | yes | supported |
| `file: sequence_diagram` | no | no | no | no | yes | supported |
| `file: api_table` | no | no | no | no | yes | supported |
| `file: er_diagram` | no | no | no | no | yes | supported |
| `file: render_index` | no | no | no | no | yes | supported |
| `asset` | yes | yes | yes | no | yes | supported |
| `view: api_table` | no | no | no | no | yes | supported; `list_endpoints` is dedicated to the computed route list |
| `view: er_diagram` | no | no | no | no | yes | supported |
| private sub node | yes | yes | yes | no | yes | supported |
| `primitive` | no | no | no | no | no | reference target only |

Cell value meaning:

| value | meaning |
|---|---|
| `yes` | The tool treats this selector as a normal query target. |
| `no` | The tool does not treat this selector as a normal query target. |
| `limited` | The tool handles this selector, but with limited return scope or information volume. |

Status meaning:

| status | meaning |
|---|---|
| `supported` | Treated as a query target in MCP v1. |
| `supported / limited inspect` | signature / references are handled, but the dedicated inspect's information volume is limited. |
| `supported / limited references` | Reference meaning is limited, e.g. at file granularity. |
| `partial` | Only some tools support it. |
| `future` | A candidate for design-conversation coverage, not yet implemented. |
| `v1 optional` | Permitted by the spec, but not implementation-mandatory. |
| `reference target only` | Returned as a reference target, but not a direct query target. |

`no` handling differs per tool. `get_signature` / `get_references` / `get_reference_tree` / `inspect` return `unsupported_object` tool error in principle when given a `no` selector. `analyze_impact`'s `no` instead means returning a normal response with empty `impacts` and an `unsupported_selector` diagnostic, not a tool error.

`get_reference_tree`'s `limited` for `file: node` means it follows `get_references(file: node)`'s support range — only expanding references to in-file nodes. `primitive` is reachable as a reference target but is not treated as a traversal root.

> Source: V01-ADR-054 §Decision

## QualifiedID

A node with module scope follows the QualifiedID format defined in V01-ADR-027.

```text
<module-path>.<node-kind>.<id>
```

Examples:

```text
auth.task.login
order.store.order_db
catalog.model.item
payment.stripe.task.receive_webhook
```

`<module-path>` is a dot-separated module path of one or more segments. `<node-kind>` is a brewprint node-kind sentinel.

### Actor exception

`actor` is project-global per V01-ADR-031 and does not belong to a module.

Actor references are always a direct ID reference.

```text
stripe
scheduler
end_user
```

MCP responses still return a `qualified_id` field for actors; its value may simply equal the global actor ID.

```json
{
  "object": "node",
  "kind": "actor",
  "id": "stripe",
  "qualified_id": "stripe"
}
```

## FileID

FileID is the slash-normalized string of the path relative to the brewprint project's `yaml/` directory, as defined in V01-ADR-043.

```text
auth/task/login.yaml
order/state.yaml
views/scenarios/checkout_flow.yaml
```

Windows `\` is normalized to `/` in MCP responses.

## Synthetic ID

A file-local object without a QualifiedID uses a synthetic ID as a stable reference in MCP responses.

### Private sub-node ID

```text
<file-id>#<local-id>
```

Examples:

```text
order/task/checkout.yaml#build_order
order/task/checkout.yaml#reserve_inventory
```

A sub-node is not externally referenceable from other modules, but an MCP response still needs to identify it as a reference target or as a member in `inspect(main task)`, hence the synthetic ID.

### Asset ID

```text
<producer-qualified-id>#<asset-name>
```

Examples:

```text
order.task.build_order#draft_order
auth.task.login#auth_token
```

An asset is implicitly generated from a task/join's `returns`. To query it directly, use this synthetic ID, or specify the producer as `id` and the asset name as `local_id`.

### Transition ID

The transition ID corresponds to the `(stateFileID, fromStateID, eventID, guard?)` tuple from V01-ADR-035 / V01-ADR-048, represented as:

```text
<state-file-id>#<from-state>:<event>
<state-file-id>#<from-state>:<event>[<guard>]
```

Examples:

```text
auth/state.yaml#idle:login_submitted
order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']
```

The guard string is used verbatim as decoded from YAML. No trim, whitespace normalization, Unicode normalization, or expression-AST comparison is performed. This matches the guard exact-match policy of V01-ADR-035 / V01-ADR-048.

## SourceLocation

```json
{
  "file": "auth/task/login.yaml",
  "line": 12,
  "column": 5,
  "end_line": 42,
  "end_column": 1
}
```

| field | required | content |
|---|---:|---|
| `file` | ✓ | FileID. |
| `line` | optional | 1-origin line number. |
| `column` | optional | 1-origin column number. |
| `end_line` | optional | Range end line. |
| `end_column` | optional | Range end column. |

Implementations that cannot obtain line/column may return only `file`.

## ObjectRef

The common form used to point at an object within an MCP response.

```json
{
  "object": "node",
  "kind": "task",
  "id": "auth.task.login",
  "qualified_id": "auth.task.login",
  "label": "login",
  "source": {
    "file": "auth/task/login.yaml",
    "line": 3
  }
}
```

| field | required | content |
|---|---:|---|
| `object` | ✓ | `node` / `view` / `transition` / `asset` / `field` / `file` / `primitive`. |
| `kind` | ✓ | Object kind subtype, e.g. `task` / `model` for a node. |
| `id` | ✓ | Object ID — QualifiedID, actor global ID, or file-local synthetic ID. |
| `qualified_id` | optional | Resolved QualifiedID; only for objects that have one. May equal `id`. |
| `file` | optional | Owning FileID, for file-local objects. |
| `local_id` | optional | Local ID within `file`, for file-local objects. |
| `label` | optional | Human-readable short display name. |
| `source` | optional | SourceLocation. |
| `parent` | optional | Parent ObjectRef, e.g. for a field. |

A private sub-node is represented as:

```json
{
  "object": "node",
  "kind": "task",
  "id": "order/task/checkout.yaml#build_order",
  "file": "order/task/checkout.yaml",
  "local_id": "build_order",
  "label": "build_order"
}
```

A primitive is represented as:

```json
{
  "object": "primitive",
  "kind": "primitive",
  "id": "str",
  "label": "str"
}
```

A model field is represented as:

```json
{
  "object": "field",
  "kind": "model_field",
  "id": "auth.model.user.email",
  "label": "email",
  "parent": {
    "object": "node",
    "kind": "model",
    "id": "auth.model.user"
  }
}
```

## TransitionRef

TransitionRef is an ObjectRef extension representing a transition.

```json
{
  "object": "transition",
  "kind": "transition",
  "id": "order/state.yaml#processing:payment_webhook_received[payload.status == 'succeeded']",
  "state_file": "order/state.yaml",
  "from": "processing",
  "on": "payment_webhook_received",
  "to": "confirmed",
  "guard": "payload.status == 'succeeded'",
  "action": "order.task.confirm_order",
  "source": {
    "file": "order/state.yaml",
    "line": 42
  }
}
```

| field | required | content |
|---|---:|---|
| `object` | ✓ | Fixed: `transition`. |
| `kind` | ✓ | Fixed: `transition`. |
| `id` | ✓ | TransitionID. |
| `state_file` | ✓ | State FileID where the transition is defined. |
| `from` | ✓ | Source state ID. |
| `on` | ✓ | Event ID. |
| `to` | ✓ | Target state ID. |
| `guard` | optional | Guard string; exact-match target. |
| `action` | optional | Action task QualifiedID. |
| `source` | optional | SourceLocation. |

When ObjectRefs to the state / event / task are needed, they are returned in `references` as `transition_from` / `transition_event` / `transition_to` / `transition_action`. The `from` / `on` / `to` fields on TransitionRef itself are a shorthand for readability in state diagrams / scenario steps.

## AssetRef

`asset` has no standalone YAML file — it is implicitly generated from a task's `returns`. Since an implicit asset has no QualifiedID, it can be queried directly via the synthetic ID composed of producer and asset name. `AssetRef` is returned with producer context.

```json
{
  "object": "asset",
  "id": "auth.task.login#auth_token",
  "name": "auth_token",
  "producer": "auth.task.login",
  "model": "auth.model.token",
  "scope_file": "auth/task/login.yaml"
}
```

| field | required | content |
|---|---:|---|
| `object` | ✓ | Fixed: `asset`. |
| `id` | ✓ | Asset synthetic ID, `<producer>#<name>`. |
| `name` | ✓ | `task.returns.name`. |
| `producer` | ✓ | QualifiedID or file-local synthetic ID of the task that produces the asset. |
| `model` | ✓ | Asset's model QualifiedID or primitive. |
| `scope_file` | optional | FileID where the asset arises. |

`task.returns.name` is treated as the asset label on the DAG, but since a sub-task or a different file may have a same-named `returns`, MCP responses always return it together with producer context. The synthetic ID `<producer>#<name>` is used for direct queries.

## Diagnostic

```json
{
  "severity": "warning",
  "code": "uncovered_module",
  "message": "module catalog is not covered by render_index.yaml; implicit group will be used",
  "source": {
    "file": "render_index.yaml"
  }
}
```

| field | required | content |
|---|---:|---|
| `severity` | ✓ | `error` / `warning` / `info` / `hint`. |
| `code` | ✓ | Machine-readable code. |
| `message` | ✓ | Human-readable message. |
| `source` | optional | SourceLocation. |
| `related` | optional | Array of related SourceLocation or ObjectRef. |

## Reference schema

### Reference

`Reference` represents a direct reference between brewprint objects.

```json
{
  "kind": "reads",
  "direction": "out",
  "from": {
    "object": "node",
    "kind": "task",
    "id": "auth.task.login"
  },
  "to": {
    "object": "node",
    "kind": "store",
    "id": "auth.store.user_db"
  },
  "source": {
    "file": "auth/task/login.yaml",
    "line": 10
  }
}
```

| field | required | content |
|---|---:|---|
| `kind` | ✓ | Reference kind. |
| `direction` | ✓ | Direction relative to the query target: `out` / `in`. |
| `from` | ✓ | Source ObjectRef. |
| `to` | ✓ | Target ObjectRef. |
| `source` | optional | SourceLocation where this reference is defined. |
| `doc` | optional | Natural-language supplement about the reference, e.g. a branch case label or transition note. |

### Reference kind

MCP v1 returns the following reference kinds.

| kind | from | to | meaning |
|---|---|---|---|
| `param_model` | task / branch / join | model / primitive | A param type-references a model or primitive. |
| `return_model` | task / join | model / primitive | `returns` type-references a model or primitive. |
| `produces_asset` | task / join | asset | `returns` implicitly generates an asset. |
| `consumes_asset` | asset | task / join | Flow wiring passes an implicit asset to a consumer node. |
| `reads` | task | store | Task reads a store. |
| `writes` | task | store | Task writes a store. |
| `store_of` | store | model | Store holds a model. |
| `field_type` | model field | model / primitive | Model field type-references something. |
| `field_fk` | model field | model field | Model field FK-references another field. |
| `transition_event` | transition | event | Transition references an event as its trigger. |
| `transition_from` | transition | state | Transition's from state. |
| `transition_to` | transition | state | Transition's to state. |
| `transition_action` | transition | task | Transition calls an action task. |
| `event_payload` | event | model | Event payload references a model. |
| `event_actor` | event | actor | An external event references an actor. |
| `event_watches` | event | store | An ER event watches a store. |
| `scenario_state_file` | scenario | state file | Sequence scenario references a state file. |
| `scenario_step_transition` | scenario step | transition | Sequence scenario step references a transition. |

Flow wiring (information corresponding to `flow_step` / `flow_param`) is not returned by `get_references` in MCP v1 — it is internal local structure within a DAG file and stays scoped to `inspect(task).members.flow.entries` in MCP v1.

Task return wiring via `returns.source` is likewise not returned by `get_references` in MCP v1. `returns.source` is returned as raw/resolved info in `inspect(task).members.return_source` and is not a mandatory target of the global reverse index (`referencesBySource` / `referencesByTarget`).

This policy is maintained for the M11 scope: `flow_step` / `flow_param` / `flow_branch_case` / `flow_foreach_over` / `task_return_source` are treated as flow/task-inspect vocabulary, not `Reference.kind`. These may become traversal material for a future `get_reference_tree` / `analyze_impact`, but are not included in direct-references-v1's return scope.

### Direction

`direction` represents the direction relative to the query target.

| direction | meaning |
|---|---|
| `out` | The party the query target references. |
| `in` | The party that references the query target. |

Example: querying `inspect(auth.store.user_db)` where `auth.task.login` reads `user_db` yields:

```json
{
  "kind": "reads",
  "direction": "in",
  "from": { "id": "auth.task.login", "kind": "task" },
  "to": { "id": "auth.store.user_db", "kind": "store" }
}
```

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.mcp.overview` | Parent overview; tool catalog. |
| `spec:bpdsl.mcp.errors` | Error model that consumes the selector / Diagnostic shapes defined here. |
| `spec:bpdsl.dsl.naming` | QualifiedID format and bare ID resolution that this schema's QualifiedID section is derived from. |

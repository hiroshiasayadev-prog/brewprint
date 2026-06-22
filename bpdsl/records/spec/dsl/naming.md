# Reference: Name resolution

- **id**: `spec:bpdsl.dsl.naming`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.overview`

## What this is

Namespace and ID resolution rules for brewprint YAML. Covers the module / folder hierarchy, QualifiedID format, sentinel-based parsing, bare ID resolution within a module, actor global namespace, FK field reference resolution, and `reads` / `writes` store reference resolution.

> Source: V01-ADR-002, V01-ADR-003, V01-ADR-020, V01-ADR-027, V01-ADR-031, V01-ADR-033, V01-ADR-058, V01-ADR-070, V01-ADR-078

## Module and folder hierarchy

The directory hierarchy under `yaml/` maps directly to the module hierarchy.

```
yaml/
  auth/
    dag.yaml        ← module: auth
  commerce/
    cart/
      dag.yaml      ← module: commerce.cart
    order/
      dag.yaml      ← module: commerce.order
```

Nesting depth is unbounded. Module names are expressed as dot-separated paths.

> Source: V01-ADR-002, V01-ADR-027

## QualifiedID

Full-path node references use the following form.

```
<module-path>.<node-kind>.<id>
```

- `<module-path>`: dot-separated path of one or more segments (e.g., `auth`, `commerce.cart`)
- `<node-kind>`: a sentinel keyword (see §Node kind sentinel keywords)
- `<id>`: public node ID — only main nodes have a public QualifiedID

Sub-nodes have file-private local IDs and no public QualifiedID. They cannot be referenced from other files as `<module>.<kind>.<sub-node-id>`.

Task-file helper model semantics are defined in [`spec:bpdsl.dsl.nodes.data`](nodes/data.md). This spec defines only that `<module>.model.<id>` QualifiedID targets public models only.

MCP query layer may use `<semantic-anchor-id>#<local-id>` synthetic IDs (V01-ADR-078) for file-private / generated objects, but these are not public QualifiedIDs and are not used in YAML authoring. MCP schema / ObjectRef migration details are out of scope for this section and are not changed by this resolver spec.

Examples:

```yaml
auth.task.login
commerce.cart.store.cart_db
analysis.state.session
```

> Source: V01-ADR-027

## Node kind sentinel keywords

QualifiedID parsing determines the boundary between module path and node ID by detecting the first occurrence of a sentinel keyword.

| keyword | node kind |
|---|---|
| `task` | Task |
| `model` | Model |
| `store` | Store |
| `actor` | Actor |
| `event` | Event |
| `state` | State |
| `branch` | Branch |
| `fork` | Fork |
| `join` | Join |

Parse example:

```
auth.oauth.task.login
           ↑
           sentinel found → left side is module path, right side is id
→ module: auth.oauth
→ type:   task
→ id:     login
```

> Source: V01-ADR-027

## Bare ID resolution

Within-module main node references do not require the full QualifiedID; bare IDs are resolved.

Bare IDs in `flow.step` / `reads` / `writes` and similar fields are first resolved against file-private sub-nodes / sources within the same file. If no match is found, the resolver falls back to same-module main nodes.

Cross-module references require a QualifiedID.

```yaml
# Within auth/dag.yaml
flow:
  - step: login        # resolves to auth.task.login

# Cross-module (full path required)
transitions:
  - action: auth.task.login
```

For TypeRef bare model name resolution see [`spec:bpdsl.dsl.type_ref`](type-ref.md) §Named model TypeRef. This section defines only the name collision rules for task-file helper models vs. public models.

> Source: V01-ADR-003, V01-ADR-027, V01-ADR-058

### Helper model name collision rules

Task-file helper model visibility, identity, and reference scope are defined in [`spec:bpdsl.dsl.nodes.data`](nodes/data.md). This section defines the name collision rules that protect TypeRef readability.

| case | result |
|---|---|
| Two helper models in the same file share the same id | invalid |
| A helper model, main node, or private sub-node in the same file shares a local id | invalid |
| A helper model and a public model in the same module share the same id | invalid |
| Helper models in different files within the same module share the same id | valid |
| A helper model and a public model in different modules share the same id | valid |
| Helper models in different modules share the same id | valid |

Helper shapes that need to be reused externally must be promoted to public models.

> Source: V01-ADR-070

## Actor global namespace

`actor` is a project-global node kind and does not belong to any module. Actor QualifiedIDs do not use `<module-path>.actor.<id>` form — the actor ID itself is the reference.

```yaml
nodes:
  - id: stripe
    type: actor
    note: "External payment service"
```

From any module:

```yaml
actor: stripe
```

Actor IDs must be unique across the project. Duplicate actor IDs in multiple files are a validation error.

> Source: V01-ADR-031, V01-ADR-025

## FK field reference resolution

`model.fields[].fk` references a field in another model.

### Same-module bare FK

Within the same module, a short form is allowed.

```text
<model-id>.<field-name>
```

Example:

```yaml
# yaml/order/model/order_item.yaml
fields:
  - name: order_id
    type: str
    fk: order.id
```

When the current module is `order`, this resolves to:

```text
order.model.order.id
```

Resolution rule:

```text
<model-id>.<field-name>  →  <current-module>.model.<model-id>.<field-name>
```

### Cross-module FK

Cross-module field references require a qualified field reference.

```text
<module-path>.model.<model-id>.<field-name>
```

Example:

```yaml
# yaml/payment/model/payment_event.yaml
fields:
  - name: order_id
    type: str
    fk: order.model.order.id

# Nested module
    fk: commerce.order.model.order.id
```

A bare FK used cross-module is always interpreted as a reference within the current module — it cannot be used for cross-module references.

### Post-resolve normalization

Regardless of input form, the semantic model / reference index / MCP response always uses the resolved, normalized field ID.

```text
input:     fk: order.id
resolved:  order.model.order.id
reference: field_fk from order.model.order_item.order_id to order.model.order.id
```

This normalization allows `get_references` and similar reverse-lookup operations to treat bare FK and qualified field references uniformly.

> Source: V01-ADR-033, V01-ADR-027

## Reads / writes store reference resolution

`task.reads` / `task.writes` reference store nodes. Same-module references allow bare store IDs; cross-module references require QualifiedIDs.

### Same-module store reference

```yaml
# yaml/order/task/checkout.yaml
nodes:
  - id: checkout
    type: task
    reads: [order_db]
    writes: [order_db]
```

When the current module is `order`, this resolves to `order.store.order_db`.

### Cross-module store reference

```yaml
# yaml/order/task/checkout.yaml
reads:
  - cart.store.cart_session
  - auth.store.user_db
writes:
  - order_db     # resolves to order.store.order_db
```

A bare store ID used cross-module is interpreted as a reference within the current module.

> Source: V01-ADR-020, V01-ADR-027

## Post-resolve ID representation

After name resolution, the semantic model / reference index / MCP response returns normalized IDs, regardless of input form.

| YAML input | context | resolved |
|---|---|---|
| `order_db` | `reads` / `writes` in `order` module | `order.store.order_db` |
| `order.id` | `fk:` in `order` module | `order.model.order.id` |
| `cart.store.cart_session` | cross-module store reference | `cart.store.cart_session` |
| `commerce.order.model.order.id` | cross-module FK reference | `commerce.order.model.order.id` |
| `stripe` | actor reference | `stripe` |
| `validate_request` | private sub-task reference inside `mcp/task/get_signature.yaml` | file-local sub-node identity; not normalized to public QualifiedID |
| `mcp.task.get_signature#validate_request` | MCP query layer synthetic ID for a private sub-task | semantic-anchor-based synthetic ID; not a public QualifiedID |

This approach lets YAML authors write concise same-module references while the implementation, MCP layer, and diagnostics all use unambiguous IDs.

Not all resolved IDs are public QualifiedIDs. File-private sub-nodes have no public QualifiedID and are treated as file-local identities. When a stable MCP reference is needed, `<semantic-anchor-id>#<local-id>` synthetic IDs (V01-ADR-078) are used instead. MCP schema / ObjectRef migration details are out of scope for this section and are not changed by this resolver spec.

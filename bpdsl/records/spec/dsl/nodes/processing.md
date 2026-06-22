# Reference: Processing nodes

- **id**: `spec:bpdsl.dsl.nodes.processing`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.nodes.overview`

## What this is

Field definitions for the five Processing-layer node kinds: `task`, `asset`, `branch`, `fork`, and `join`. These nodes appear in the DAG view. `task` is the core unit; `asset`, `branch`, `fork`, `join` are derived or control-flow nodes.

## task

> Source: V01-ADR-009, V01-ADR-010, V01-ADR-011, V01-ADR-014, V01-ADR-017, V01-ADR-020, V01-ADR-070, V01-ADR-071

Processing layer. The unit of processing. A `returns` declaration implicitly creates an `asset` node in the DAG.

```yaml
- id: login
  type: task
  main: true
  endpoint: true
  method: POST
  path: login
  params:
    - name: credentials
      model: credential
  returns:
    name: auth_token
    model: token
  reads: [session_store]
  writes: [session_store]
  initializes:
    - name: login_log
      model: login_log
      note: "Initialize empty login_log"
  note: "Validate credentials and issue a token"
```

### task fields

| field | required | type | description | source |
|---|---|---|---|---|
| `main` | optional | bool | `true` declares this as the main node. One per file. Applies to `task` only (V01-ADR-011). | V01-ADR-011 |
| `params` | optional | list\<param\> | Input parameter list. TypeRef references. | V01-ADR-009, V01-ADR-060 |
| `returns` | optional | returns | Output. TypeRef reference, asset name declaration, optional return source wiring. | V01-ADR-009, V01-ADR-060, V01-ADR-062 |
| `reads` | optional | list\<store-id\> | Store IDs this task reads from. | V01-ADR-020 |
| `writes` | optional | list\<store-id\> | Store IDs this task writes to. | V01-ADR-020 |
| `endpoint` | optional | bool | `true` includes this task in `list_endpoints` MCP output and API Table aggregation. | V01-ADR-005, V01-ADR-017, V01-ADR-028 |
| `method` | required if endpoint | enum | HTTP method: GET / POST / PUT / DELETE / PATCH. | V01-ADR-005 |
| `path` | optional | string | Endpoint leaf path (e.g. `login`). Omitted defaults to `task.id`. Single segment only — no `/`. Full path is composed from `http_root_path` + module hierarchy (V01-ADR-028). | V01-ADR-005, V01-ADR-028 |
| `initializes` | optional | list\<init\> | Store declarations for use within this file. Main node only. | V01-ADR-014 |
| task-file helper model | optional | `type: model` node | File-private schema definition within the same task file. See [data.md](./data.md) §Task-file private helper model. | V01-ADR-070, V01-ADR-071 |

### param object

| field | required | type | description |
|---|---|---|---|
| `name` | ✓ | string | Parameter name within this task. |
| `model` | ✓ | TypeRef | Type reference: primitive / named model / inline `list<T>` / inline `dict<T>`. |

### returns object

| field | required | type | description |
|---|---|---|---|
| `name` | ✓ | string | Name of the generated asset. |
| `model` | ✓ | TypeRef | Type reference: primitive / named model / inline `list<T>` / inline `dict<T>`. |
| `source` | optional | string (wiring source syntax) | Where the return value comes from: node id / collected asset source / initialized source / `$params.<name>`. `$item` is not allowed. |

`returns` is single-valued. Wrap multiple outputs in a struct model (V01-ADR-009).

`returns.source` omitted: no implicit name-match connection is made, even if `returns.name` matches a flow source name.

### init object (inside initializes)

| field | required | type | description |
|---|---|---|---|
| `name` | ✓ | string | Store reference name within this file. |
| `model` | ✓ | model-id | Store type (model ID). |
| `note` | optional | string | Initial value or initialization method description. |

`initializes` stores are file-private. External reference is not allowed (V01-ADR-014).

`initializes[].name` participates in the bare wiring source namespace of the same file and can be referenced as an initialized source from `returns.source` and flow wiring tokens.

### Task-file private helper model signature exposure

- `params[].model` references a task-file private helper model → `invalid_private_model_reference` error. Params are a public input contract.
- `returns.model` may reference a same-file private helper model. No diagnostic in minimum scope.

> Source: V01-ADR-062 §1–7, V01-ADR-063 §1

---

## asset

> Source: V01-ADR-010, V01-ADR-065

Processing layer. A flow-level entity. **Has no standalone file.** Created implicitly from a `task` node's `returns` declaration.

No direct YAML definition syntax exists for `asset`. In DAG diagrams and MCP tools, `task.returns.name` is treated as the asset ID.

### asset immutability

`asset` is an immutable output snapshot from a single task execution.

| rule | description |
|---|---|
| Not a cross-edge target | `reads:` / `writes:` accept only store IDs; asset is excluded. |
| No write syntax | No YAML construct exists to mutate an asset after creation. |
| Not an accumulator | `asset` cannot model "a box that appends over time." Use `store` for that. |

For mutable runtime instances, use `store`. Runtime-language reference or mutation behavior is out of scope for brewprint (V01-ADR-065 §4).

> Source: V01-ADR-065 §1–5

---

## branch

> Source: V01-ADR-012, V01-ADR-023

Processing layer. Exclusive branch. Selects **exactly one** downstream path based on a condition.

```yaml
- id: route_by_role
  type: branch
  params:
    - name: user
      model: user
  note: "Routes to admin_flow or user_flow depending on user.role"
```

### branch fields

| field | required | type | description | source |
|---|---|---|---|---|
| `params` | optional | list\<param\> | Input used for the branching decision. | V01-ADR-012 |

Merge point is implicit — read from edge structure. No explicit merge node (V01-ADR-012).

**Scope rule**: Assets generated inside a branch block cannot be referenced from outside that block. Pass data out via `initializes` + `writes` / `reads`. If each path ends independently with no convergence needed, the floating downstream tasks render as going directly to END in the DAG (V01-ADR-023).

For `flow:` syntax see [edges/data-flow.md](../edges/data-flow.md) §Exclusive branch entry.

---

## fork

> Source: V01-ADR-012

Processing layer. Parallel split. Executes **all** downstream paths concurrently. Must always be paired with a `join`.

```yaml
- id: fan_out
  type: fork
  note: "Run static analysis, dynamic analysis, and dep check in parallel"
```

`fork` without a paired `join` is invalid (V01-ADR-012). Inputs to each branch task are declared in `flow:` via `branches[].steps[].params` (V01-ADR-040).

No additional fields beyond the common fields (`id`, `type`, `note`).

---

## join

> Source: V01-ADR-012

Processing layer. Parallel merge. Waits for all branches of the paired `fork` to complete, then aggregates the results. Must always be paired with a `fork`.

```yaml
- id: aggregate
  type: join
  params:
    - name: static_result
      model: static_result
    - name: dynamic_result
      model: dynamic_result
    - name: dep_result
      model: dep_result
  returns:
    name: full_report
    model: full_report
  note: "Combine results from 3 branches into full_report"
```

### join fields

| field | required | type | description | source |
|---|---|---|---|---|
| `params` | optional | list\<param\> | Inputs from each branch. List all branch outputs. | V01-ADR-012 |
| `returns` | optional | returns | Aggregated result asset declaration. | V01-ADR-009, V01-ADR-012 |

`join` without a paired `fork` is invalid (V01-ADR-012).

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.dsl.nodes.overview` | Parent overview; node kind boundary matrix. |
| `spec:bpdsl.dsl.edges.data_flow` | `flow:` wiring syntax for branch, fork, foreach. |
| `spec:bpdsl.dsl.type_ref` | TypeRef syntax used in params and returns. |

# Reference: Data flow edges

- **id**: `spec:bpdsl.dsl.edges.data_flow`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.edges.overview`

## What this is

Syntax and rules for the `flow:` section: all Processing-layer node wiring, control flow constructs (step, fork, branch, foreach), the sigil system, flow wiring type compatibility, and task return wiring.

> Source: V01-ADR-015, V01-ADR-016, V01-ADR-060, V01-ADR-061

## Design principle

Node definitions in `nodes:` hold signature only (types and names). All inter-node wiring — which output feeds which input — is consolidated in the `flow:` section (V01-ADR-015).

Analogues: Airflow (`@task` + `>>`), Prefect (`@task` + `@flow`), Temporal (activity + workflow).

## Control flow scope

> Source: V01-ADR-023

**Assets created inside a control flow construct (`branch` / `fork` / `foreach`) cannot be referenced from outside that construct's scope.**

Exception: a `foreach.returns` declared collected asset source is exported and can be referenced as a bare source from subsequent flow entries in the same file. This is an explicit escape hatch — it does not make individual iteration assets externally accessible.

To pass data outside a scope when `foreach.returns` collect is insufficient, use `initializes` to pre-declare a store, `writes` to store the value, and `reads` for the downstream task to consume it.

```yaml
# Invalid: referencing an asset from inside a branch scope
- step: finalize
  params:
    result: admin_flow    # admin_flow is inside a branch scope

# Valid: pass data out via a store
- id: admin_flow
  type: task
  writes: [role_result_store]

- id: finalize
  type: task
  reads: [role_result_store]
```

When paths end independently with no convergence needed, no store is required. Floating downstream tasks render as going directly to END in the DAG (V01-ADR-023).

## Step entry

```yaml
flow:
  - step: fetch_data
    params:
      config: $params.config   # reference field "config" from main node params

  - step: transform
    params:
      raw: fetch_data          # reference the entire returns of fetch_data
```

### Step entry fields

| field | required | description | source |
|---|---|---|---|
| `step` | ✓ | Node ID to execute. | V01-ADR-015 |
| `params` | optional | Input wiring. key = param name, value = wiring source. | V01-ADR-015 |

### Wiring source notation

| notation | meaning |
|---|---|
| `source_node` | Entire `returns` of source_node. |
| `collected_asset` | Collected asset source from a preceding `foreach.returns`. |
| `$params.field` | A specific field from the main node's params (file boundary input). |
| `$item` | Current iteration element (valid inside `foreach.params` only). |

**Wiring unit is always the entire task `returns`** (V01-ADR-015). `source_node.field` notation does not exist. To access a sub-field, insert an explicit extract task.

`$params.field` requires an explicit field name. No implicit name-match via bare `$params` (V01-ADR-015).

```yaml
# Invalid: direct sub-field access
- step: static_analysis
  params:
    raw: fetch_data.raw_data

# Valid: explicit extract task
- step: extract_raw
  params:
    raw: fetch_data
- step: static_analysis
  params:
    raw: extract_raw
```

## Fork entry

> Source: V01-ADR-012, V01-ADR-015, V01-ADR-040

```yaml
flow:
  - fork: fan_out
    branches:
      - steps:
          - step: static_analysis
            params:
              raw: fetch_data
      - steps:
          - step: dynamic_analysis
            params:
              raw: fetch_data
      - steps:
          - step: dep_check
            params:
              raw: fetch_data
    join: aggregate
```

### Fork entry fields

| field | required | description | source |
|---|---|---|---|
| `fork` | ✓ | Fork node ID. | V01-ADR-012, V01-ADR-015 |
| `branches` | ✓ | List of branch objects. | V01-ADR-040 |
| `join` | ✓ | Corresponding join node ID. | V01-ADR-015 |

`fork` and `join` must always be paired. Either alone is invalid (V01-ADR-012).

### Branch object (inside fork)

| field | required | description | source |
|---|---|---|---|
| `steps` | ✓ | List of step objects for this branch. | V01-ADR-040 |

Each element of `steps[]` uses the same format as a regular flow step entry.

| field | required | description | source |
|---|---|---|---|
| `step` | ✓ | Node ID to execute. | V01-ADR-015, V01-ADR-040 |
| `params` | optional | Input wiring. | V01-ADR-015, V01-ADR-040 |

The old-form `branches: - [step_a, step_b]` is not used. Always use `steps:` with explicit `step:` entries, even when no params are needed.

Implicit propagation via `fork.params` to branch steps is not adopted. Each branch step's inputs are written explicitly in `steps[].params` (V01-ADR-040).

### join.params resolution

The `params` of the `join` node are resolved by name-matching against each branch's terminal step `returns.name`.

```yaml
nodes:
  - id: static_analysis
    type: task
    returns:
      name: static_result
      model: static_result

  - id: aggregate
    type: join
    params:
      - name: static_result
        model: static_result
```

`static_analysis.returns.name == static_result` maps to `aggregate.params.static_result`. A parser error is raised if no matching branch terminal step returns is found.

## Exclusive branch entry

> Source: V01-ADR-012, V01-ADR-023, V01-ADR-040

```yaml
flow:
  - branch: route_by_role
    params:
      user: fetch_user
    cases:
      - label: admin
        step: admin_flow
        params:
          user: fetch_user
      - label: user
        step: user_flow
        params:
          user: fetch_user
```

### Branch entry fields

| field | required | description | source |
|---|---|---|---|
| `branch` | ✓ | Branch node ID. | V01-ADR-012 |
| `params` | optional | Input wiring for the branch node's own decision logic (same rules as step entry). | V01-ADR-023, V01-ADR-040 |
| `cases` | ✓ | List of case entry objects. | V01-ADR-023 |

`branch.params` supplies only the branch node's own decision input. Case entry task inputs are written explicitly in `cases[].params` (V01-ADR-040).

### Cases entry fields

| field | required | description | source |
|---|---|---|---|
| `label` | ✓ | Condition label. Human/LLM description. Evaluation is outside brewprint scope. | V01-ADR-023 |
| `step` | ✓ | Entry point task ID for this case (single node). | V01-ADR-023 |
| `params` | optional | Input wiring for the case entry task. | V01-ADR-040 |

`step` is a single node ID. Subsequent steps after the entry point are determined by the DAG wiring structure (params references), so no step list inside cases is needed. Only one case executes per branch invocation (V01-ADR-023).

Assets created inside a case block cannot be referenced from outside the branch. `cases[].params` declares the case entry task input; it does not allow external access to case-internal assets.

## Foreach entry

> Source: V01-ADR-013 (superseded), V01-ADR-016, V01-ADR-060, V01-ADR-061

`foreach` is a control flow construct in the `flow:` section, not a node type (V01-ADR-016).

```yaml
flow:
  - foreach: process_item
    mode: sequential
    over: fetch_items
    params:
      item: $item
      config: $params.config
    returns: results
```

### Foreach entry fields

| field | required | description | source |
|---|---|---|---|
| `foreach` | ✓ | Apply-target task ID (same-file sub-node or external main node). | V01-ADR-016 |
| `over` | ✓ | List source to iterate: a node ID (preceding task / join `returns` asset) or `$params.field` (main task params). TypeRef resolution and `$item` type derivation: see §Type compatibility. | V01-ADR-016, V01-ADR-060 |
| `mode` | optional | `sequential` (default) or `map` (parallel). | V01-ADR-016 |
| `params` | optional | Apply-target task params wiring (same rules as step entry). Required if the apply target has params. `$item` references the current iteration element. | V01-ADR-016 |
| `returns` | optional | Name for the collected asset source — the per-iteration `returns` from the apply target task, collected into a list. Specify when downstream flow needs the collected result. Side-effect-only foreach may omit. | V01-ADR-016, V01-ADR-061 |

### Mode semantics

| mode | meaning | implementation pattern |
|---|---|---|
| `sequential` | Process elements one at a time in order. | Standard for loop. |
| `map` | Process all elements in parallel. | ProcessPoolExecutor / asyncio.gather. |

### foreach.returns collected asset source

`foreach.returns` declares a file-local source name for the per-iteration collected results. The collected TypeRef is `list<T>` where `T` is the apply-target task's `returns.model`.

```yaml
flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items

  - step: summarize_cart
    params:
      items: validated_items    # references the collected list
```

`validated_items` is a file-local source name declared at the foreach invocation level — it is not the apply-target task's `returns.name`.

Rules:
- `foreach.returns` is optional. Omit when downstream flow does not need the collected result (side-effect only).
- If the apply-target task has no `returns` and `foreach.returns` is specified: `invalid_foreach_returns`.
- Referencing own `returns` name from within the same foreach's `params`: `invalid_foreach_returns`.
- If apply-target `returns.model` is `any`: collected TypeRef is `list<any>`.
- If apply-target `returns.model` is unresolvable: collected TypeRef is also unresolvable; suppress `incompatible_wiring_type` for downstream wiring.

`foreach.id` is not introduced. When the same apply-target task is used in multiple foreach entries, distinguish the collected results by different `returns` names.

DAG rendering: foreach decorates the apply-target task's box with a ↻ icon. Foreach is not rendered as a standalone box (V01-ADR-016).

> Source: V01-ADR-061 §1–8

## Sigil reference

> Source: V01-ADR-015, V01-ADR-016

| sigil | meaning | valid locations |
|---|---|---|
| `$params.field` | A specific field from the main node's params (file boundary input). Field name is required. | `flow:` wiring |
| `$item` | Current iteration element from `foreach`. | `foreach.params` only |

Sigils are syntactically distinct from node IDs and explicitly mark "externally injected" values (V01-ADR-015).

## Type compatibility

> Source: V01-ADR-060, V01-ADR-061

Flow wiring validates assignment compatibility from source TypeRef to target param TypeRef.

### Wiring validation scope

| wiring location | source | target |
|---|---|---|
| `step.params` | wiring source | step task params |
| `branch.params` | wiring source | branch node params |
| `branch.cases[].params` | wiring source | case entry task params |
| `fork.branches[].steps[].params` | wiring source | branch step task params |
| `foreach.params` | wiring source | foreach apply task params |
| `join.params` | implicit source via terminal step `returns.name` match | join node params |

### TypeRef compatibility rules

Named list/dict models are first normalized to their container TypeRef form. Assignment from source `S` to target `T` is valid only when:

1. `S` or `T` is `any`
2. Both are the same primitive
3. Both are non-list/dict named models with the same QualifiedID
4. Both are list and their element TypeRefs are compatible
5. Both are dict and their value TypeRefs are compatible

Otherwise: `incompatible_wiring_type`.

```text
str -> str            OK
str -> int            NG
user -> user          OK
user -> order         NG
any -> user           OK
user -> any           OK
list<user> -> list<user>     OK
list<user> -> list<order>    NG
list<any> -> list<user>      OK
user_list -> list<user>      OK  (named list normalized)
config_map -> dict<config>   OK  (named dict normalized)
str -> user           NG
```

For named list/dict normalization and TypeRef syntax see [`spec:bpdsl.dsl.type_ref`](../type-ref.md).

### Wiring source TypeRef resolution

| source notation | resolved TypeRef |
|---|---|
| node ID / QualifiedID | `returns.model` of the referenced task or join |
| `$params.<name>` | `model` of the matching entry in the same file's main task `params` |
| `$item` | Element TypeRef derived from `foreach.over`: `T` if over is `list<T>` or a named list model; `any` if over is `any` |
| `foreach.returns` collected asset source name | `list<T>` where `T` is the apply-target task's `returns.model` |
| `initializes[].name` initialized source name | Named model TypeRef derived from `initializes[].model` |

Bare node/source resolution order: same-file file-private sub-node / source first, then same-module main node fallback. Cross-module references require QualifiedID.

### Resolution failure handling

Type compatibility is only checked when both source and target TypeRef are successfully resolved. `incompatible_wiring_type` is suppressed when:

- Source TypeRef is unresolvable
- Target param TypeRef is unresolvable
- `foreach.over` cannot be treated as a list (suppresses `$item` incompatibility)
- Collected asset source TypeRef is unresolvable

This prevents cascading diagnostics from unresolved references.

### Error codes

| code | condition |
|---|---|
| `incompatible_wiring_type` | Source and target TypeRefs are incompatible. |
| `invalid_wiring_source` | Reference exists but cannot be used as a wiring source in this context. |
| `unresolved_wiring_source` | Source token resolves to nothing. |
| `duplicate_flow_source` | Same name used by more than one of: node id, `foreach.returns`, `initializes[].name` in the same file. |
| `invalid_foreach_over_type` | `foreach.over` result cannot be treated as a list. |
| `invalid_foreach_returns` | `foreach.returns` specified but apply-target has no `returns`, or self-reference from params. |

> Source: V01-ADR-060 §5, V01-ADR-061 §3–6, §9, V01-ADR-063 §2, §6, §7

## Task return wiring

> Source: V01-ADR-062

`task.returns.source` is the task return wiring field — it declares where the task gets the value it exposes as its `returns`. Unlike `flow:` wiring (internal node-to-node), return wiring connects a task's internal flow output to its external signature.

`returns.name` / `returns.model` define the external signature. `returns.source` specifies where that value comes from inside the task.

`returns.source` is optional. Omit for leaf tasks, note-only tasks, and external boundary tasks. Specify when a main task with `flow:` returns an internal flow result, or when explicitly declaring a pass-through return.

```yaml
nodes:
  - id: validate_cart
    type: task
    main: true
    params:
      - name: cart_items
        model: cart_item_list
    returns:
      name: validated_items
      model: cart_item_list
      source: validated_items

flow:
  - foreach: validate_item
    over: $params.cart_items
    params:
      cart_item: $item
    returns: validated_items
```

### Valid source forms for returns.source

| source | meaning |
|---|---|
| node ID / QualifiedID | Output of a task or join that has `returns`. |
| collected asset source | A `foreach.returns` collected asset source in the same file. |
| initialized source | A file-private source from `initializes[].name`. |
| `$params.<name>` | Pass-through of a main task param field as-is. |

`$item` is not valid in `returns.source`. To return a foreach aggregate, use `foreach.returns` and reference the collected source.

Evaluation timing: `returns.source` is evaluated at task / flow completion (END point), not at source position. This means `returns.source` does not apply flow-entry ordering (no forward reference restriction).

Bare node/source resolution in `returns.source` follows the same order as flow wiring: same-file file-private first, then same-module main node fallback. Cross-module requires QualifiedID.

### Error codes

| code | condition |
|---|---|
| `unresolved_return_source` | Source token resolves to nothing. |
| `invalid_return_source` | Source resolves but cannot be used as a return source (e.g., a node with no `returns`, or `$item`). |
| `incompatible_return_type` | Source TypeRef and `returns.model` TypeRef are incompatible (same rules as §Type compatibility). |

`incompatible_return_type` is suppressed when source or `returns.model` TypeRef is unresolvable.

`returns.name` matching a flow source name does not constitute an implicit connection — `returns.source` must be specified explicitly. (The `join.params` implicit name-match is an existing exception maintained as-is.)

> Source: V01-ADR-062 §1–8, V01-ADR-063 §1, §3, §5

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.dsl.edges.overview` | Parent overview; edge kind summary. |
| `spec:bpdsl.dsl.type_ref` | TypeRef syntax and named list/dict normalization rules. |
| `spec:bpdsl.dsl.naming` | Bare ID and QualifiedID resolution rules. |
| `spec:bpdsl.dsl.nodes.processing` | Processing node field definitions including `initializes`. |

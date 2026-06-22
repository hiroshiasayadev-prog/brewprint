# Reference: Data nodes

- **id**: `spec:bpdsl.dsl.nodes.data`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.nodes.overview`

## What this is

Field definitions for the two Data-layer node kinds: `model` and `store`. `model` defines type schemas. `store` holds runtime data instances. Both appear in the ER diagram; `store` also appears in the DAG.

## model

> Source: V01-ADR-007 (superseded; content carried forward in V01-ADR-010), V01-ADR-008, V01-ADR-010, V01-ADR-021, V01-ADR-060, V01-ADR-067, V01-ADR-073

Data layer. Pure type definition. Does not appear in the DAG.

Public model files live in the `model/` subdirectory as module-level schema definitions (one file per definition). Task files may also contain file-private helper models; see §Task-file private helper model below.

Model file rendering and the model file helper view are owned by `bpdsl/records/spec/views/model-file.md`, not this spec. Model catalog and the UC-002 model response helper-shape migration are not defined in this task-file helper minimum.

```yaml
# struct
- id: user
  type: model
  kind: struct
  fields:
    - name: id
      type: str
      pk: true
    - name: email
      type: str
    - name: role_id
      type: str
      fk: role.id
    - name: profile
      type: user_profile
    - name: created_at
      type: datetime

# list
- id: item_list
  type: model
  kind: list
  element: item

# dict
- id: config_map
  type: model
  kind: dict
  value: config

# enum
- id: mcp_diagnostic_severity
  type: model
  kind: enum
  values:
    - error
    - warning
    - info
    - hint

# tagged_union
- id: analyze_impact_change
  type: model
  kind: tagged_union
  discriminator: kind
  variants:
    - tag: rename
      fields:
        - name: new_id
          type: str
    - tag: remove
      fields: []
    - tag: change_type
      fields:
        - name: new_type
          type: str
```

### model fields

| field | required | type | description | source |
|---|---|---|---|---|
| `kind` | ✓ | enum | `struct` / `list` / `dict` / `enum` / `tagged_union` | V01-ADR-007→010, V01-ADR-021, V01-ADR-067, V01-ADR-073 |
| `fields` | required for struct | list\<field\> | Field definitions (struct only). | V01-ADR-008 |
| `element` | required for list | TypeRef | Element type. | V01-ADR-021, V01-ADR-060 |
| `value` | required for dict | TypeRef | Value type. Key is always `str`. | V01-ADR-021, V01-ADR-060 |
| `values` | required for enum | list\<string\> | Allowed value set. Non-empty; no empty strings; no duplicates within the same enum. | V01-ADR-067 |
| `discriminator` | required for tagged_union | string | Top-level field name used to identify the variant. Non-empty; no dot paths; no external discriminator. | V01-ADR-073 |
| `variants` | required for tagged_union | list\<variant\> | Variant definitions. Non-empty; `tag` values unique within the model. | V01-ADR-073 |

`kind: scalar` is retired — use primitive literals directly (V01-ADR-021).
`kind: dict` key is always `str`; no `key` field exists (V01-ADR-021).

### Primitive reserved words

The following are primitive reserved words and cannot be used as model IDs (V01-ADR-021). Full TypeRef syntax is defined in [`spec:bpdsl.dsl.type_ref`](../type-ref.md).

| primitive | meaning |
|---|---|
| `str` | string |
| `int` | integer |
| `float` | floating point |
| `bool` | boolean |
| `bytes` | byte sequence |
| `datetime` | date-time |
| `any` | untyped (use sparingly) |

### enum semantics

`kind: enum` represents a string-valued finite vocabulary. `values` is required; `fields` / `element` / `value` are absent. Usage is via named model TypeRef — no inline enum TypeRef syntax exists.

`values` order is preserved for display and schema generation but carries no type-compatibility meaning. Per-value metadata (`note`, `label`, `deprecated`) is not introduced in v1.1.

> Source: V01-ADR-067 §1–5

### tagged_union semantics

`kind: tagged_union` is an object model where the variant is identified by a `discriminator` field within the same object. Usage is via named model TypeRef — no inline `union<...>` / `tagged_union<...>` TypeRef syntax.

`discriminator` is a top-level field name only. Dot paths, external discriminators, and adjacent discriminators are not supported in the V01-WORK-DATA-010 minimum. The discriminator field is implicitly present in the tagged union object and must not appear in `variants[].fields`.

`variants` is a non-empty list. Each `variants[].tag` set is the allowed discriminator value set; duplicates within the same model are not allowed. Variants with no payload fields use `fields: []`.

Variant and field ordering follows YAML order and is used for rendering and schema generation but not for type-compatibility.

> Source: V01-ADR-073 §1–4, §7, §10

### variant object (inside tagged_union)

| field | required | type | description | source |
|---|---|---|---|---|
| `tag` | ✓ | string | Discriminator field value. Non-empty; unique within this model. | V01-ADR-073 |
| `fields` | ✓ | list\<variant-field\> | Payload field definitions. `[]` for no payload. | V01-ADR-073 |

### variant-field object (inside tagged_union variant)

| field | required | type | description | source |
|---|---|---|---|---|
| `name` | ✓ | string | Payload field name. Unique within the variant. Must not match `discriminator`. | V01-ADR-073 |
| `type` | ✓ | TypeRef | Type: primitive / named model / inline `list<T>` / inline `dict<T>`. | V01-ADR-060, V01-ADR-073 |
| `note` | optional | string | Human docstring and LLM semantic contract. | V01-ADR-008, V01-ADR-073 |

`pk` / `fk` / `unique` are not used in variant-fields (ER/struct-only semantics).

### struct field object

| field | required | type | description | source |
|---|---|---|---|---|
| `name` | ✓ | string | Field name. Unique within the struct. | V01-ADR-008 |
| `type` | ✓ | TypeRef | Type: primitive / named model / inline `list<T>` / inline `dict<T>`. | V01-ADR-008, V01-ADR-060 |
| `pk` | optional | bool | `true` = primary key column. One per struct. | V01-ADR-021 |
| `fk` | optional | `<model-id>.<field-name>` | FK reference target. Omitted = JSON embed. | V01-ADR-021 |
| `unique` | optional | bool | `true` = 1:1 relation. Used with `fk:`. Omitted = many-to-one. | V01-ADR-026 |
| `note` | optional | string | Human docstring and LLM semantic contract. | V01-ADR-008, V01-ADR-021 |

**`fk` field semantics:**

| notation | DB treatment |
|---|---|
| `type: str, fk: role.id` | FK column referencing `role.id`. |
| `type: user_profile` (no fk) | JSON-embedded column. |
| `type: tag_list` (list kind model) | Variant / JSON column (implementation-dependent). |

`type` validation: TypeRef syntax check + existence check against primitive reserved words or defined model IDs. `invalid_type_ref` if syntax is invalid; `unresolved_field_type` if named model is not found.

### Task-file private helper model semantics

A `type: model` node without `main: true` inside a task file is a file-private helper model.

| property | rule |
|---|---|
| visibility | file-private. No `visibility:` field introduced. |
| public identity | No public QualifiedID. |
| local identity | Local model ID within the defining task file. |
| reference scope | Bare TypeRef from within the same YAML file only. |
| external reference | No cross-file reference; no QualifiedID syntax for task-file helper models. |
| render role | Not rendered in DAG body. May appear in task-file render `## Private models` section. |

A helper model ID must be unique among local node IDs in the same file and must not collide with public model IDs in the same module (prevents ambiguous bare TypeRef resolution).

For helper model visibility in task signature exposure, see [`spec:bpdsl.dsl.nodes.processing`](./processing.md) §Task-file private helper model signature exposure.

> Source: V01-ADR-070, V01-ADR-071

---

## store

> Source: V01-ADR-007 (kind definitions; content carried forward in V01-ADR-010), V01-ADR-019 (distinction from state), V01-ADR-065 (two-form summary)

Processing / Data layer. Runtime data holder. Distinct from FSM `state` (V01-ADR-019).

`store` is the umbrella term for all runtime data instances in brewprint and comes in two forms depending on how it is declared:

| form | declaration | scope | source |
|---|---|---|---|
| store node | `type: store` in `store/*.yaml` | module-level (public QualifiedID) | V01-ADR-007 → V01-ADR-010 |
| initialized store | `initializes[]` in a task file | file-private (same file only) | V01-ADR-014, V01-ADR-063 |

Both forms are valid targets for cross-edge `reads:` / `writes:` and have a mutation vocabulary. For contrast with `asset`, see [processing.md](./processing.md) §asset.

The fields below apply to both forms. `initialized store` does not use `kind`; it is declared with `model` only (see [processing.md](./processing.md) §init object).

```yaml
- id: user_db
  type: store
  kind: db
  of: user
  note: "User table"

- id: session_store
  type: store
  kind: session
  of: session
  note: "HTTP session"

- id: user_collection
  type: store
  kind: collection
  of: user
  note: |
    - find_by_email: match by email
    - active_users: is_active = true
```

### store fields

| field | required | type | description | source |
|---|---|---|---|---|
| `kind` | ✓ | enum | `db` / `session` / `collection` / `context` | V01-ADR-007 → V01-ADR-010 |
| `of` | optional | model-id | ID of the model this store holds. | V01-ADR-007 |

`store.kind: db` has no own field definitions. The ER diagram follows `store.of` → model → fields to draw column definitions (V01-ADR-021).

`store.of` is treated as a model-id reference (not a TypeRef). For the fields that do accept a TypeRef see [`spec:bpdsl.dsl.type_ref`](../type-ref.md) §Fields that accept TypeRef.

### collection note convention

| write in note | do not write in note |
|---|---|
| Filter / search on a single collection's fields | Joins across multiple stores |
| Conditions expressible in words (equality, range, bool flag) | Complex queries involving aggregation, transformation, or sorting |

Collection notes serve as LLM semantic contracts (same role as V01-ADR-007, V01-ADR-008 `note` fields).

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.dsl.nodes.overview` | Parent overview; node kind boundary matrix. |
| `spec:bpdsl.dsl.type_ref` | TypeRef syntax used in model field and element types. |
| `spec:bpdsl.dsl.edges.cross_edges` | Cross-edge reads/writes targeting stores. |

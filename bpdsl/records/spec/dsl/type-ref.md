# Reference: TypeRef

- **id**: `spec:bpdsl.dsl.type_ref`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.overview`

## What this is

Type reference syntax and compatibility rules for brewprint v1.1. Defines the TypeRef form used in `params` / `returns` / model fields, primitive reserved words, named model references, inline list/dict containers, named list/dict model normalization, TypeRef error codes, enum/tagged-union compatibility, and the opaque container warning.

TypeRef is not a runtime type system — it is a minimal representation for statically verifying the design contract shared between humans and LLMs. Brewprint uses TypeRef to check flow wiring type compatibility, derive `$item` type in foreach, and handle shallow container I/O at external boundaries.

> Source: V01-ADR-060

## Fields that accept TypeRef

TypeRef is used in the following fields.

| location | field | meaning |
|---|---|---|
| task / branch / fork / join params | `params[].model` | input param type |
| task / join returns | `returns.model` | output asset type |
| struct model field | `fields[].type` | field type |
| list model | `element` | list element type |
| dict model | `value` | dict value type |

`store.of`, `initializes[].model`, and `event.payload.model` are not TypeRef targets — they use existing model-id references.

TypeRef is also used to derive `$item` type during foreach. `$item` is not a field that receives a TypeRef — it is derived from the TypeRef of `foreach.over` at flow wiring resolution time. See [`spec:bpdsl.dsl.edges.data_flow`](edges/data-flow.md) §Type compatibility.

> Source: V01-ADR-060

## TypeRef syntax

A TypeRef is one of the following.

| kind | example | meaning |
|---|---|---|
| primitive | `str`, `int`, `float`, `bool`, `bytes`, `datetime`, `any` | primitive reserved word |
| named model | `user`, `order`, `catalog.product` | reference to a model ID or model QID |
| inline list | `list<user>` | list with element type `user` |
| inline dict | `dict<config>` | dict with value type `config`; key is always `str` |

Inline union syntax (`union<...>`, `oneOf<...>`, `anyOf<...>`, `str | int`) does not exist in TypeRef. For untagged unions / general oneOf, use a tagged union envelope model referenced as a named model.

```yaml
params:
  - name: files
    model: list<source_file>

returns:
  name: diagnostics
  model: list<diagnostic>

nodes:
  - id: module_config
    type: model
    kind: struct
    fields:
      - name: settings
        type: dict<any>
```

`list<T>` and `dict<T>` are built-in container TypeRefs, not user-defined generics. User-declared type variables, `model<T>`, `task<T>`, `T extends X`, and generic function inference do not exist.

> Source: V01-ADR-060

## Primitive reserved words

The following words are treated as primitive TypeRefs. They must not be used as model IDs.

This list covers all primitive reserved words from V01-ADR-021 §3. The primitive table in V01-ADR-060 §1 is representative, not exhaustive — this section is authoritative.

| primitive | meaning |
|---|---|
| `str` | string |
| `int` | integer |
| `float` | floating-point number |
| `bool` | boolean |
| `bytes` | byte sequence |
| `datetime` | date and time |
| `any` | untyped; treated as a wildcard in both directions in flow wiring compatibility |

`any` is an escape hatch. Use should be minimal. In v1.1 there is no narrowing syntax, so `any` is compatible with all types in both source and target positions.

> Source: V01-ADR-021, V01-ADR-060

## Named model TypeRef

Any TypeRef that is neither primitive nor an inline container is treated as a named model reference. `enum` and `tagged_union` models are also referenced as named model TypeRefs.

Task-file helper model semantics are defined in [`spec:bpdsl.dsl.nodes.data`](nodes/data.md). This section defines only the resolution order used by the TypeRef resolver.

When a bare named model TypeRef appears in a task file, the resolver uses the following order:

1. If it matches a primitive reserved word → treat as primitive TypeRef.
2. If it matches `list<T>` / `dict<T>` → treat as inline container TypeRef; resolve `T` recursively.
3. Bare name → look for a file-private helper model in the same YAML file first.
4. If not found → look for a public model in the same module.
5. If still not found → emit an unresolved diagnostic at the use site.

QualifiedID named model TypeRefs target public models only; they do not resolve to task-file helper models.

Resolution to a file-private helper model in the same file does not mean the reference is valid at all TypeRef use sites. If `params[].model` references a file-private helper model from the same task file, the post-resolution diagnostic is `invalid_private_model_reference`. Referencing a file-private helper model from `returns.model` in the same file is valid; no diagnostic in minimum scope.

```yaml
nodes:
  - id: get_preview
    type: task
    main: true
    returns:
      name: preview
      model: preview_response

  - id: preview_response
    type: model
    kind: struct
    fields:
      - name: items
        type: list<preview_item>

  - id: preview_item
    type: model
    kind: struct
    fields:
      - name: title
        type: str
```

In this example `returns.model: preview_response` and `fields[].type: list<preview_item>` both resolve to helper models in the same file.

Named model TypeRef allows recursive references. Use an existing named model TypeRef for the field; do not introduce an inline recursive shape. Resolvers and renderers treat recursive references as named references, not unbounded expansions.

```yaml
nodes:
  - id: object_ref
    type: model
    kind: struct
    fields:
      - name: parent
        type: object_ref
        optional: true
```

Named models are compared nominally. `user` and `customer` with identical fields are not type-compatible. `enum` models are compared nominally — they are not implicitly compatible with `str` even though their underlying JSON representation is a string. `tagged_union` models are compared nominally — models with the same variant set or field structure are not compatible if they have different IDs.

For shape conversion between types, introduce an explicit adapter / normalize task.

Public model vs. task-file helper model name collision rules: [`spec:bpdsl.dsl.naming`](naming.md) §Helper model name collision rules. In valid YAML, a public model in the same module cannot be shadowed by a helper model in the same file.

> Source: V01-ADR-060, V01-ADR-070

## Inline list / dict TypeRef

`list<T>` represents a list with element type `T`.

```yaml
params:
  - name: users
    model: list<user>
```

`dict<T>` represents a dict with value type `T`. The key is always `str`; there is no syntax for specifying a key type.

```yaml
params:
  - name: config_by_name
    model: dict<config>
```

TypeRef is recursive, so the following are syntactically valid.

```text
list<any>
list<dict<user>>
dict<list<diagnostic>>
```

Nested `list<T>` / `dict<T>` is syntactically valid and is not a validation error.

For implementation safety, a TypeRef with container nesting depth exceeding 16 is an `invalid_type_ref` error. Container nesting depth is the number of nested `list<T>` / `dict<T>` layers; primitives and named models have depth 0.

| TypeRef | depth |
|---|---:|
| `diagnostic` | 0 |
| `list<diagnostic>` | 1 |
| `dict<list<diagnostic>>` | 2 |
| `list<dict<list<any>>>` | 3 |

Anonymous inline struct TypeRef is not supported in v1.1.

```text
list<{ id: str, severity: str }>   # not valid TypeRef
```

For `dict<T>`, since the value type alone cannot express key semantics, the key meaning should be made explicit via a field name, model name, or `note`.

> Source: V01-ADR-060, V01-ADR-069

## Named list / dict model normalization

Named list and dict models from V01-ADR-021 are valid in v1.1.

```yaml
- id: user_list
  type: model
  kind: list
  element: user
```

For type compatibility checking, this is normalized to the same container shape as:

```text
list<user>
```

Similarly:

```yaml
- id: config_map
  type: model
  kind: dict
  value: config
```

is treated as compatible with:

```text
dict<config>
```

The `id` / `note` of a named list/dict model are retained for human/LLM documentation. Flow wiring type compatibility is checked using the normalized container shape.

```text
user_list → list<user>      OK
list<user> → user_list      OK
user_list → list<order>     NG
```

Struct models are not normalized; they are compared as named models nominally.

> Source: V01-ADR-060

## TypeRef syntax errors

A string that cannot be parsed as a TypeRef is an `invalid_type_ref` diagnostic.

Examples:

```text
list<
dict<>
list<user
map<user>
```

Using a container kind other than `list` or `dict` (e.g., `map<user>`) is also `invalid_type_ref` since it cannot be parsed as a TypeRef.

When TypeRef syntax is valid but the named model inside cannot be resolved, use the existing unresolved diagnostic for the use site: `unresolved_field_type` for `fields[].type`; `unresolved_model` for `params[].model` / `returns.model` / `model.element` / `model.value`. Use `invalid_type_ref` only when the syntax itself is broken, an unsupported container kind is used, or the parser safety limit is exceeded.

If a task-file helper model is not in the resolution scope of the reference site, the existing unresolved diagnostics above may be used. If the TypeRef resolved to a helper model in the same file but the use site is `params[].model`, emit `invalid_private_model_reference` rather than an unresolved diagnostic.

When TypeRef resolution fails, do not additionally emit `incompatible_wiring_type` for wiring that uses that TypeRef. Type compatibility is only checked when both source and target TypeRef are successfully resolved.

Parser safety limit violations are also `invalid_type_ref`. This limit is a hard error for parser / implementation safety, not a readability lint.

> Source: V01-ADR-060, V01-ADR-069

## Enum model TypeRef compatibility

Enum models are referenced as named model TypeRefs. There is no inline enum variant syntax in TypeRef.

```yaml
nodes:
  - id: diagnostic
    type: model
    kind: struct
    fields:
      - name: severity
        type: mcp_diagnostic_severity
```

Enum model type compatibility is nominal.

```text
mcp_diagnostic_severity → mcp_diagnostic_severity  OK
mcp_diagnostic_severity → str                      NG
str → mcp_diagnostic_severity                      NG
mcp_diagnostic_severity → impact_severity          NG
any → mcp_diagnostic_severity                      OK
mcp_diagnostic_severity → any                      OK
```

Implicit `enum` / `str` compatibility is not introduced. `any` follows V01-ADR-060: compatible in both source and target positions.

> Source: V01-ADR-067

## Tagged union model TypeRef compatibility

Tagged union models are referenced as named model TypeRefs. Inline union syntax (`union<...>`, `tagged_union<...>`, `oneOf<...>`, `anyOf<...>`, scalar union) does not exist in TypeRef.

```yaml
nodes:
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
```

Tagged union model type compatibility is nominal.

```text
analyze_impact_change → analyze_impact_change  OK
analyze_impact_change → other_change           NG
analyze_impact_change → any                    OK
any → analyze_impact_change                    OK
analyze_impact_change → str                    NG
```

`any` follows V01-ADR-060: compatible in both source and target positions.

Variant field `type` is treated as a regular TypeRef. `invalid_type_ref` for broken syntax; `unresolved_field_type` for unresolved named model.

Untagged union / general oneOf are not introduced either as a model kind or as a TypeRef syntax. For machine-readable schema where union inference is needed, use a tagged union envelope model. For intentionally opaque locations, use `any + note` / prose.

> Source: V01-ADR-073

## Opaque container TypeRef warning

When a container TypeRef contains `any` and the shape meaning is not captured by a named model, it is subject to an `opaque_type_ref` warning diagnostic.

Targets:

```text
list<any>
dict<any>
list<dict<any>>
dict<list<any>>
list<dict<list<any>>>
```

`opaque_type_ref` is a warning — validation succeeds. The message may suggest extracting a named model.

`opaque_type_ref` targets `any` inside container TypeRefs as a debt visibility baseline. It does not warn about or require resolution of bare `any` fields or `any + note` used as the main response shape.

`unclear_dict_key` / `deep_type_ref` are future lint candidates and are not added to v1.1 minimum diagnostics.

> Source: V01-ADR-069

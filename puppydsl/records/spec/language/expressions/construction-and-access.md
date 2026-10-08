# Reference: PuppyDSL construction and access

- **id**: `spec:puppydsl.language.expressions.construction_and_access`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.expressions`

## What this is

Defines transparent-record construction, immutable dictionary construction, field projection, and required-key dictionary lookup.

## Transparent-record construction

A transparent-record constructor names one concrete declared record type.

```yaml
location: >-
  SourceLocation{
    path: path,
    line: line,
    column: column
  }
```

A zero-field transparent record uses an empty body.

```yaml
not_formed: >-
  CandidateNotFormed{}
```

| rule | contract |
|---|---|
| Authoring form | One YAML folded scalar using `>-`. |
| Target | One concrete transparent-record type. |
| Field set | Every declared field exactly once and no unknown field. |
| Field value | One literal or visible binding reference. |
| Field order | No semantic effect. |
| Trailing comma | Optional. |
| Result | Exact named record type. |

Record-constructor syntax is invalid for primitives, named scalars, enums, unions, opaque types, and generic composite types.

## Dictionary construction

A dictionary constructor declares exact key and value types.

```yaml
entries: >-
  dict<string, SourceLocation>{
    [first_name]: first_location,
    [second_name]: second_location
  }
```

The form is:

```text
dict<K, V>{[key]: value, ...}
```

| rule | contract |
|---|---|
| Authoring form | One YAML folded scalar using `>-`. |
| Key type | Must satisfy the dictionary key restriction. |
| Key and value | One literal or visible binding reference. |
| Duplicate key | Never overwrites an earlier entry. |
| Entry order | No semantic meaning. |
| Empty constructor | Valid because `K` and `V` are explicit. |
| Result | Exact immutable type `dict<K, V>`. |

A statically evident duplicate key is a program error. Equal runtime binding values discovered during construction cause execution failure.

## Access chains

An access chain begins with one visible binding and applies zero or more typed postfix operations.

```text
<binding>(.<field> | [<key>])*
```

Examples:

```yaml
path: location.path
app: inventory.apps[app_name]
source_root: inventory.apps[app_name].source_root
```

### Field projection

`.<field>` requires a transparent-record receiver. The field must exist in the declared record shape.

The result type is the declared field type.

### Required-key dictionary lookup

`[<key>]` requires a `dict<K, V>` receiver. The key is one primitive literal, enum literal, or visible binding whose type is exactly `K`.

The result type is `V`.

A missing runtime key terminates the current invocation as an execution failure. Required-key lookup does not produce `optional<V>`, insert a value, or select a fallback.

## Access restrictions

| receiver or operation | result |
|---|---|
| Opaque value field projection | Invalid. |
| Union value projection before switch | Invalid. |
| `optional<T>` projection before `present` switch | Invalid. |
| Optional target in `absent` case | No value exists. |
| List bracket access | Invalid. |
| Access chain beginning with a call or constructor | Invalid. |
| Lookup key containing a call, projection, or constructor | Invalid. |

Inside a union switch case, the target may expose fields defined by the selected transparent-record variant. Inside an optional `present` case, the target may expose operations defined by `T`.

## Nested construction restriction

Record fields, dictionary keys, and dictionary values accept only literals or visible bindings.

Inline calls, projections, access chains, nested constructors, arithmetic, comparison, and boolean expressions are invalid in those positions.

## Validation rules

- Reject a constructor target that is unknown, generic but uninstantiated, or not the required constructible type.
- Reject missing, unknown, or duplicate record fields.
- Reject constructor values that are not assignable to their declared targets.
- Reject invalid dictionary key types.
- Reject duplicate dictionary keys.
- Reject access to an undeclared field.
- Reject bracket access on a non-dictionary receiver.
- Reject a lookup key whose type differs from exact `K`.
- Reject direct access to opaque, union, or uneliminated optional representation.
- Reject any prohibited nested expression.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.expressions` | Parent expression overview. |
| `spec:puppydsl.language.type_system.declared_types` | Defines transparent and opaque types. |
| `spec:puppydsl.language.type_system.assignability` | Defines constructor compatibility. |
| `spec:puppydsl.language.control_flow.switch` | Defines union and optional narrowing. |

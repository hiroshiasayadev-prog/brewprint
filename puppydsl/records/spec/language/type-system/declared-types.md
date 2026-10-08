# Reference: PuppyDSL declared types

- **id**: `spec:puppydsl.language.type_system.declared_types`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.type_system`

## What this is

Defines the declaration shapes and value semantics for PuppyDSL named scalars, enums, transparent records, opaque types, and unions.

## Declaration body selection

Each type declaration contains `detail` and exactly one type body.

| body field | declared form |
|---|---|
| `scalar` | Named scalar. |
| `enum` | Closed enumeration. |
| `record` | Transparent record. |
| `opaque: true` | Opaque nominal type. |
| `union` | Closed union. |

Combining multiple type bodies is invalid.

## Named scalar

A named scalar uses one supported primitive representation while preserving distinct type identity.

```yaml
types:
  ArtifactKind:
    detail: Identifies one application-defined artifact kind.
    scalar: string
```

Named scalars:

- are nongeneric;
- expose no fields;
- do not declare a closed value set;
- are distinct from their backing primitive;
- are distinct from every other named scalar.

## Enum

An enum declares a non-empty duplicate-free member list.

```yaml
types:
  RecordStructure:
    detail: Identifies one supported record structure.
    enum:
      - tree
      - sequential
```

Every member is a value of the enum type. A member literal uses `RecordStructure.tree` form.

Enums are nongeneric, expose no fields, and remain distinct from strings and other enums.

## Transparent record

A transparent record declares the complete field set visible to PuppyDSL.

```yaml
types:
  SourceLocation:
    detail: Identifies one source position.
    record:
      path: string
      line: integer
      column: integer
```

The declaration owns field names and field types. Native implementations cannot expose additional fields through the language.

A zero-field transparent record is valid.

```yaml
types:
  CandidateNotFormed:
    detail: Represents one explicit information-bearing outcome.
    record: {}
```

A zero-field record is a nominal singleton value. It is distinct from `void`, optional absence, and every other zero-field record.

### Generic records

Generic parameters appear in the declaration key.

```yaml
types:
  SetComparison<T>:
    detail: Preserves asymmetric differences and equality.
    generic:
      T: equatable
    record:
      left_only: list<T>
      right_only: list<T>
      equal: boolean
```

The `generic` field is omitted when no parameter requires a constraint.

### Recursive records

A record-declaration cycle is valid only when every cycle crosses at least one `list<T>` or dictionary value boundary.

```yaml
types:
  TreeNode:
    detail: Preserves one recursively nested tree node.
    record:
      children: list<TreeNode>
```

These cycles are invalid:

- direct self-reference through a record field;
- mutual recursion through direct record fields;
- recursion whose only boundary is `optional<T>`.

PuppyDSL has no pointer, nullable-reference, or address-identity type.

## Opaque type

An opaque declaration exposes type identity without exposing representation.

```yaml
types:
  ParsedDocument:
    detail: Represents parsed source whose internal form remains hidden.
    opaque: true
```

Opaque types:

- are nongeneric;
- expose no fields;
- cannot be constructed by PuppyDSL expressions;
- admit no implicit conversion;
- are produced and consumed only through declared typed boundaries.

Native representation and code-generation rules are outside the language specification.

## Union

A union declares a non-empty duplicate-free list of separately declared variant types.

```yaml
types:
  ParseOutcome:
    detail: Carries either one parsed value or one expected invalid-input value.
    union:
      - ParsedDocument
      - InvalidInput
```

A union value contains one active variant and one value of that variant. The active variant is interpreter-owned type information.

PuppyDSL cannot:

- construct a union literal;
- project a tag or payload representation;
- inspect the active variant as data;
- perform arbitrary runtime type assertions.

A union is consumed through exhaustive switch. Inside one case, the switch target is narrowed to that case's declared variant.

Unions are nongeneric in the initial language profile.

## Validation rules

- Reject an unknown or unsupported named-scalar primitive.
- Reject empty or duplicate enum members.
- Reject duplicate record fields.
- Reject unknown field type expressions or generic parameters.
- Reject invalid recursive record cycles.
- Reject `opaque: true` combined with another type body or generic parameters.
- Reject empty, duplicate, unknown, or generic union variants.
- Reject any declaration with more than one type body.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.type_system` | Parent type-system overview. |
| `spec:puppydsl.language.type_system.type_expressions` | Defines field and variant type references. |
| `spec:puppydsl.language.type_system.assignability` | Defines compatibility among declared types. |
| `spec:puppydsl.language.control_flow.switch` | Eliminates enum, union, and optional values. |

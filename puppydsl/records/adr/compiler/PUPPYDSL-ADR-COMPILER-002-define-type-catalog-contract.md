# PUPPYDSL-ADR-COMPILER-002: Define the TypeCatalog contract

- **status**: accepted
- **date**: 2026-08-05
- **depends_on**:
  - `PUPPYDSL-ADR-COMPILER-001`
- **supersedes**: []
- **migrated_to_spec**: null

## Context

`PUPPYDSL-ADR-COMPILER-001` established the type-information consumers and queries that the compiler frontend must support.

The next design step is to define the logical contract of the global type-definition catalog without prematurely fixing Go structs, maps, pointer identity, parsing phases, or backend representation.

The catalog must distinguish globally declared definitions from declaration-local type parameters and from concrete type uses such as `list<CurrentRecord>` or `SetComparison<ArtifactKind>`.

## Decision

One PuppyDSL source set produces one global `TypeCatalog`.

The catalog resolves one globally visible type identity to one type definition. It covers both built-in definitions and user-declared definitions through one lookup boundary.

### Registered definitions

The catalog registers these built-in atomic and control definitions:

- `string`;
- `boolean`;
- `integer`;
- `float`;
- `void`;
- `never`.

The catalog registers these built-in composite type-constructor definitions:

- `list` with one type parameter;
- `optional` with one type parameter;
- `dict` with ordered key and value type parameters.

The catalog registers every user-declared type definition of these kinds:

- named scalar;
- enum;
- transparent record;
- opaque type;
- union.

### Definition contract

Every catalog definition exposes:

- its globally unique type identity;
- its definition category;
- ordered generic parameters and constraints when declared;
- its category-specific shape.

Category-specific shape consists of:

| category | retained shape |
|---|---|
| Primitive | Primitive identity. |
| Control result | Exact `void` or `never` identity. |
| Composite constructor | Ordered parameter identities and constructor arity. |
| Named scalar | Resolved backing primitive definition. |
| Enum | Ordered declared members. |
| Transparent record | Ordered declared fields and each field's resolved type use. Zero fields remain valid. |
| Opaque type | Nominal identity only; no exposed representation. |
| Union | Ordered resolved direct-variant definitions. |

Record fields and enum members remain owner-local child definitions rather than independent global type-catalog entries. A record field retains its owning record definition, authored field name, declaration position, and declared `ResolvedTypeUse`. An enum member retains its owning enum definition, authored member name, and declaration position. Their semantic identities are owner-local, so equal authored names under different owners remain distinct.

Union variants do not introduce a separate owner-local variant definition. Each direct union variant is the already-resolved declared `TypeDefinition` referenced by the union definition, and exhaustive union-switch case identity uses that declared type identity directly.

### Generic boundaries

Active generic parameters are declaration-local definitions. They are not global `TypeCatalog` entries.

Concrete applications such as these are resolved type uses rather than new catalog definitions:

- `list<CurrentRecord>`;
- `optional<string>`;
- `dict<ArtifactKind, CurrentRecord>`;
- `SetComparison<ArtifactKind>`.

A generic declared type contributes one catalog definition containing its ordered parameter declarations and its shape expressed using declaration-local parameter references.

Each declaration-local generic parameter is represented by one owner-local definition retaining its owning declaration, declaration position, authored parameter name, and optional declared constraint. Parameter identity is owner-local rather than global; parameters with the same authored name in different declarations are distinct.

For a generic declared type application, the target definition's ordered parameters are paired with the `ResolvedTypeUse` ordered arguments when substitution is required. The resolved type use does not retain a duplicate substitution map.

For a validated generic callable invocation, the inferred substitution is retained explicitly as ordered associations from resolved callable-parameter definitions to resolved type arguments. This avoids repeating inference when consumers need the concrete invocation signature. Constraint declarations remain on parameter definitions and are not duplicated into each validated substitution.

### Resolved type use

Every concrete use of a type is represented logically as one `ResolvedTypeUse` containing:

- one resolved target, which is either a global `TypeCatalog` definition or an active declaration-local type-parameter definition;
- zero or more ordered resolved type arguments.

The same recursive form covers primitive types, control results, nongeneric declared types, built-in composites, active type parameters, and generic declared applications.

Exact resolved-type identity is determined by the same resolved target together with recursively identical ordered type arguments. This contract does not require pointer identity, interning, or a global concrete-instance cache.

A `ResolvedTypeUse` does not duplicate a substituted declaration shape. Consumers obtain an instantiated field or other declaration-local type by applying the use's ordered arguments to the target definition's declared parameters when required.

A validated `ResolvedTypeUse` has already established target existence, constructor or declaration arity, type-argument resolution, applicable generic constraints, dictionary-key eligibility, and context-sensitive `void` or `never` placement. Source spelling, formatting, Go representation, and precomputed substituted shapes are not part of resolved type identity.

### Resolution boundary

The final validated `TypeCatalog` contains resolved definition bodies:

- named-scalar backing primitives resolve to built-in primitive definitions;
- record field types are `ResolvedTypeUse` values;
- union variants resolve to declared type definitions;
- declared generic parameters resolve within their owning declaration.

Construction may use header registration and later body resolution to support forward and recursive references. This decision does not require a specific construction algorithm or intermediate representation.

## Rationale

A single lookup boundary prevents built-in type behavior from being duplicated across declaration validation, invocation validation, control-flow typing, and lowering.

Keeping global definitions separate from resolved uses avoids treating every concrete composite or generic application as a new declaration.

Retaining declaration order supports deterministic enum, record, union, and backend processing without making source order part of unrelated lookup semantics.

Keeping opaque representation outside the catalog preserves the language boundary between nominal type identity and native implementation representation.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Maintain separate unrelated catalogs for built-in and user-declared types. | Consumers would need duplicated dispatch and lookup behavior. |
| Register every concrete generic or composite application globally. | Concrete applications are uses of existing definitions rather than declarations and would introduce unnecessary canonicalization requirements. |
| Register active generic parameters globally. | Parameter identity and lifetime are declaration-local. |
| Register enum members, record fields, or union memberships as global type identities. | These identities are meaningful only within their owning declaration. |
| Store opaque native representation in the type catalog. | Native representation is a backend and binding concern, not a language type-definition property. |
| Fix concrete Go structs, maps, or pointer identity now. | The logical contract can be established independently of implementation representation. |
| Intern every resolved composite or generic application. | Exact type identity can be determined from the resolved target and ordered arguments without requiring canonical instance identity. |
| Copy a fully substituted declared shape into each resolved use. | Instantiated shape can be derived from the target definition and ordered arguments, so copying it would duplicate definition information. |

## Consequences

- Type-name resolution uses one global `TypeCatalog` boundary.
- Built-in atomic types, control results, and composite constructors are represented as definitions visible through the same boundary as declared types.
- User-declared definitions retain ordered parameters and category-specific shape.
- Record fields and enum members are owner-local child definitions; union variants reuse the referenced declared `TypeDefinition` identity rather than introducing another semantic identity.
- Final record-field type references and other definition-body references are resolved rather than retained as raw source strings.
- Active generic parameters are owner-local definitions carrying owner, position, authored name, and optional constraint; they are not global catalog entries.
- Generic declared type applications derive substitution by pairing target parameters with `ResolvedTypeUse` ordered arguments rather than storing a duplicate substitution map.
- Validated generic callable invocations retain their inferred ordered parameter-to-type substitution because later invocation consumers require the concrete signature without repeating inference.
- Concrete applications use `ResolvedTypeUse` rather than global catalog entries.
- `ResolvedTypeUse` retains one resolved target plus ordered resolved type arguments and does not retain a duplicated substituted declaration shape.
- Exact resolved-type identity does not require interning or a global concrete-instance catalog.
- Dictionary-key eligibility, assignability, recursive-record analysis, and Go representation remain consumers of the retained definition information rather than fields duplicated into every definition.
- No concrete catalog data structure, parser staging model, generic-instance cache, or backend type representation is established.

## Evidence

- `spec:puppydsl.language.type_system.declared_types` defines named scalar, enum, transparent record, opaque, union, generic-record, and recursive-record semantics.
- `spec:puppydsl.language.type_system.type_expressions` defines primitive, control-result, declared, parameter, composite, and generic-declared type uses.
- `spec:puppydsl.language.type_system.assignability` requires exact nominal identity and direct variant-to-union compatibility.
- `PUPPYDSL-ADR-COMPILER-001` requires one global type-catalog boundary and distinguishes global definitions from active parameters and concrete applications.

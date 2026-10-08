# PUPPYDSL-ADR-COMPILER-001: Define type-information consumers and catalog requirements

- **status**: accepted
- **date**: 2026-07-29
- **depends_on**: []
- **supersedes**: []
- **migrated_to_spec**: null

## Context

The PuppyDSL language Specifications define type declarations, type expressions, callable signatures, expression typing, control-flow typing, contexts, events, and whole-program validity.

Those rules establish many uses of type information, but they do not yet provide one compiler-facing inventory of the consumers and queries that internal artifacts must support.

Choosing `TypeExpr`, type-declaration, or catalog structures before identifying those consumers would make the representation depend on guessed implementation patterns rather than language requirements.

PuppyDSL type information primarily supports static validation and later Go lowering. The language does not require general runtime reflection or runtime manipulation of compiler type objects.

## Decision

Define compiler-facing type artifacts from the type-information queries required by the current language Specifications.

### Type-information consumers

The compiler must make type information available to these consumers:

| consumer | required use of type information |
|---|---|
| Type-declaration validation | Resolve type annotations; distinguish built-in types, declared types, and active type parameters; validate declaration kind, generic arity and constraints, dictionary-key eligibility, `void` and `never` placement, union variants, and recursive record references. |
| Function-signature validation | Resolve every input and output type, preserve ordered input names, validate generic parameters, and identify the exact callable output type. |
| Invocation validation | Resolve the callable, determine literal and binding argument types, validate argument assignment, infer one consistent generic substitution when applicable, and determine the call result type. |
| Step and return validation | Assign a static type to each information-bearing binding; create no binding for `void` or `never`; validate completion-label use; retain that `void` continues the current lexical path while `never` terminates it; and compare return bindings with the declared function output. |
| Record and dictionary construction | Determine whether the target type is constructible; obtain record fields or dictionary key and value types; and validate supplied values. Validated transparent-record constructors retain the resolved record target and supplied-field mappings. Dictionary-constructor lowering retention remains deferred until a concrete PuppyDSL source exercises the form. |
| Field and dictionary access | Determine the receiver category; resolve each field or lookup step in order; determine the resulting type; and retain the root binding identity plus each resolved access segment target and segment result type for later lowering. |
| Conditional validation | Require an exact `boolean` condition; compare the result types of value-producing branches; and represent an omitted branch of an information-bearing conditional as `IterationResultContributionAbsent` rather than `void` or an implicit `continue`. |
| Exhaustive switch validation | Obtain union variants, enum members, or the contained optional type; validate exhaustive cases; and establish case-local narrowed types. |
| Bounded iteration validation | Obtain list element or dictionary key and value types, type iteration bindings, determine returned element type, form the complete `list<U>` result when present, and omit an element when the per-iteration result is `IterationResultContributionAbsent`. |
| Provider invocation validation | Resolve the provider target as an ordinary callable; bind every callable input exactly once to a resolved context-provided value; validate supplied types, generic substitution, and output type; and retain the complete resolved provider invocation for later context and route analysis. |
| Context provided-entry validation | Derive each provided-value type from its validated provider invocation, reject `void` or `never` provider outputs, and make the resolved provided entry available to `requires`, dependency analysis, and route validation. |
| Context destructor validation | Type `self` as the provided entry's derived type, require one ordinary callable returning exact `void`, and prohibit all other argument sources. |
| Event validation | Resolve event payload types and derive generated event-callable signatures in the shared callable namespace. |
| Flow analysis | Carry binding types through lexical scopes, narrowing, path completion, absent iteration-result contribution, explicit `continue`, and `never` termination. |
| Go lowering | Consume resolved declared-type kinds and shapes plus resolved callable signatures. Exact Go representations remain a later decision. |

### Required type queries

Compiler artifacts must support at least these logical queries:

| query group | required answer |
|---|---|
| Identity | Whether two resolved types have the same exact type identity. |
| Category | Whether a type is primitive, control result, named scalar, enum, transparent record, opaque, union, list, optional, dictionary, or active type parameter. |
| Declared shape | Named-scalar backing primitive, ordered enum members, record fields and field types, direct union variants, and ordered generic parameters and constraints. |
| Composite shape | List element, optional contained type, and dictionary key and value types. |
| Compatibility | Whether one resolved source type is assignable to one resolved target type under the language assignability rules. |
| Narrowing | The case-local type established by union and optional switch elimination. |
| Special properties | Whether a type is information-bearing, `void`, or `never`; dictionary-key eligible; constraint eligible; constructible; or field-projectable; and whether a structured control-flow result carries its typed value or represents `IterationResultContributionAbsent`. |

### Logical artifacts

The compiler frontend must provide these distinct logical artifacts.

#### Type catalog

One source set produces one global type catalog covering both:

- built-in atomic and control types;
- built-in composite type constructors;
- user-defined named-scalar, enum, record, opaque, and union declarations.

The type catalog owns global type-name lookup and the definition information required by the type queries above.

Active generic type parameters are declaration-local and are not global type-catalog entries. Concrete generic applications such as `list<string>` or `SetComparison<ArtifactKind>` are resolved type uses, not independently declared catalog entries.

#### Callable catalog

One source set produces one global callable catalog covering both:

- ordinary declared functions;
- generated event callables.

Ordinary functions and generated event callables share one callable identity namespace.

For static typing, each callable entry must expose at least:

- callable identity;
- ordered input names and resolved input types;
- exactly one resolved output type;
- ordered generic parameters and constraints when present;
- whether the callable originates from an ordinary function or an event declaration.

Concrete implementation form, runtime binding, generated Go symbol, and execution mechanism are outside this catalog requirement.

#### Lexical binding static type environment

Function-body analysis derives a `LexicalBindingStaticTypeEnvironment` from:

- function inputs;
- required context-provided values;
- earlier step results;
- iteration bindings;
- active ancestor bindings;
- case-local narrowed switch targets.

The `LexicalBindingStaticTypeEnvironment` maps lexically visible binding identities to resolved static types. It is function-body-local analysis state, not a global declaration catalog.

Function-body analysis separately derives a `LexicalPathCompletionState`. A `void` result continues the current lexical path without creating a binding. A `never` result terminates the current lexical path without creating a binding. `IterationResultContributionAbsent` also completes normally without creating a binding, but it is neither a `void` value nor an implicit `continue`; when consumed as a bounded iteration's per-iteration result, it causes that iteration to contribute no list element. Completion labels never enter the `LexicalBindingStaticTypeEnvironment`.

### Representation boundary

This decision fixes required consumers, queries, and logical artifact responsibilities only.

It does not decide:

- concrete Go interfaces, structs, maps, or pointer identity;
- unresolved or resolved type-expression node shapes;
- parsing algorithms;
- name-resolution order or implementation algorithm;
- generic substitution representation or caching;
- canonicalization of concrete generic applications;
- diagnostic schemas;
- exact Go type generation;
- runtime type descriptors or reflection.

Those choices must be derived later from the requirements fixed here and from the specific artifact being mapped.

## Rationale

The language Specifications already define the semantic questions that validation must answer. Making those consumers explicit provides a stable basis for deciding what later artifacts must retain without prematurely choosing implementation structures.

A type catalog and callable catalog answer different questions and use different identities. Keeping them logically separate avoids turning one catch-all registry into an implicit program model. An implementation may aggregate access behind a larger program artifact, but the two catalog responsibilities remain distinct.

Including built-in and user-defined types behind one type-catalog boundary prevents built-in knowledge from being duplicated across invocation, construction, access, control-flow, and lowering logic.

A separate `LexicalBindingStaticTypeEnvironment` distinguishes globally declared definitions from binding types established while analyzing one function body.

Deferring exact Go lowering preserves the current frontend boundary while ensuring that the frontend retains the declared kind, shape, and callable information the backend will require.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Define `TypeExpr` and catalog structs before identifying consumers. | Required fields and boundaries would be selected from implementation habit rather than language use. |
| Keep only user-defined types in the type catalog and hard-code built-ins in each consumer. | Built-in categories, composite arity, dictionary rules, and control-result behavior would be duplicated across the compiler. |
| Combine type, callable, and lexical binding information into one undifferentiated catalog. | Global declaration identity and local analysis state have different ownership, lifetime, and query semantics. |
| Register active type parameters or concrete generic applications as global declarations. | Type parameters are declaration-local, and concrete applications are resolved uses of existing definitions rather than new declarations. |
| Delegate all type validation to the generated Go compiler. | Go compilation cannot own PuppyDSL-specific rules for nominal assignment, union narrowing, exhaustive switch, `void`, `never`, contexts, events, and restricted expressions. |
| Define exact Go representations in this decision. | Backend representation is not required to determine the frontend information contract and remains premature. |

## Consequences

- Later source-to-artifact ADRs can derive `TypeExpr`, type-declaration, resolved-type, and catalog representations from an explicit query inventory.
- The global type catalog must cover built-in definitions and user-defined declarations through one lookup boundary.
- The global callable catalog must include ordinary functions and generated event callables in one callable namespace.
- Function-body validation must maintain a separate `LexicalBindingStaticTypeEnvironment` for lexically visible binding identities and their resolved static types.
- Function-body validation must maintain a separate `LexicalPathCompletionState`; `void` continues the lexical path, `never` terminates it, and `IterationResultContributionAbsent` continues without contributing an iteration result; none creates a lexical binding.
- An information-bearing conditional with an omitted branch has the logical result `T | IterationResultContributionAbsent`; the latter is a control-flow result state, not a declared union type, `void`, or an implicit `continue` operation.
- A bounded iteration collects only contributing `T` results into `list<T>` and omits iterations whose per-iteration result is `IterationResultContributionAbsent`.
- Type declaration representations must retain all shape information needed by validation and later lowering, including named-scalar backing primitives, ordered enum members, record fields, direct union variants, and generic parameters.
- Validated transparent-record constructors must retain the resolved constructor target identity and each supplied field's resolved declared-field identity for later lowering.
- Validated field projections must retain the root binding identity and each segment's resolved declaring-record identity, resolved declared-field identity, and resolved segment result type for later lowering.
- Validated required-key dictionary lookups must retain the root binding identity, resolved dictionary receiver type, key operand identity and resolved static type, resolved dictionary key type, and resolved dictionary value and result types for later lowering.
- Validated exhaustive union switches must retain the target binding identity, original union type, resolved direct-variant set, each case's resolved variant identity and case-local narrowed target type, each case completion and result type, and the complete switch result type.
- Validated exhaustive enum switches must retain the target binding identity, original enum type, resolved declared-member set, each case's resolved enum-member identity, each case completion and result type, and the complete switch result type. Enum cases do not create case-local narrowed target types.
- Validated bounded dictionary iterations must retain the input collection binding identity and resolved dictionary type, resolved key and value types, iteration key and value binding identities and static types, per-iteration return binding identity, contributing result type, accepted non-contributing completion states, and complete `list<U>` result type.
- Validated bounded list iterations must retain the input collection binding identity and resolved list type, resolved element type, iteration item binding identity and static type, per-iteration return binding identity, contributing result type, accepted non-contributing completion states, and complete `list<U>` result type.
- Validated provider invocations must retain the resolved ordinary callable target, ordered callable-input-to-argument associations, each supplied resolved provided-value entry and static type, any inferred generic substitution, and the resolved information-bearing output type.
- Validated context provided entries must retain their full provided-value identity, owning context identity, provided key, complete validated provider invocation, and static type derived from the provider output. Provider dependency and route requirements are derived from the retained provider arguments rather than duplicated as separate typed-edge artifacts.
- Context destructor validation must resolve one ordinary callable invocation, type `self` as the provided entry's derived static type, prohibit every non-`self` argument source and hidden function requirement, and require exact `void` output. Concrete validated-destructor retention remains deferred until a PuppyDSL source exercises the form.
- Validated function declarations must retain each direct `requires` entry as a resolved reference to its context provided entry. Body binding types and owning-context requirements are obtained from that entry, while transitive requirements are derived through the call graph rather than duplicated on each function.
- Validated function declarations must retain ordered `opens` entries as resolved references to context declarations. Active context chains, provider dependency availability, and transitive requirements are derived per selected entry route rather than retained as one precomputed route map.
- The function call graph is a whole-program analysis artifact derived from resolved ordinary invocations, provider invocations, and later destructor or event-dispatch invocations. Individual function declarations do not duplicate a precomputed callee list.
- Go lowering consumes resolved declared-type kinds and shapes, recursively resolved type uses, ordered callable signatures, and validated expression and control-flow artifacts. It must not reparse source type expressions or reconstruct semantic resolution.
- Go package organization, generated-module layout, `void` and `never` representation, opaque-type binding, union and optional representation, generated symbol naming, native binder form, and whether steps are generated or interpreted remain backend decisions for later ADRs.
- Callable representations must retain ordered typed signatures independently of implementation form.
- No generic-instance cache, concrete Go struct layout, resolver algorithm, or runtime type-object contract is established.

## Evidence

- `spec:puppydsl.language.type_system.type_expressions` defines the type-expression forms and placement restrictions consumed by resolution.
- `spec:puppydsl.language.type_system.declared_types` defines named-scalar, enum, record, opaque, and union shapes.
- `spec:puppydsl.language.type_system.assignability` defines exact nominal compatibility and direct variant-to-union assignment.
- `spec:puppydsl.language.functions` defines ordered typed callable signatures and generic parameters.
- `spec:puppydsl.language.functions.steps_and_bindings` requires static step-result, discard-label, and return typing.
- `spec:puppydsl.language.expressions.invocation` requires callable resolution, argument typing, generic substitution, and result typing.
- `spec:puppydsl.language.expressions.construction_and_access` requires record, dictionary, field, and lookup type queries.
- `spec:puppydsl.language.control_flow` and its child Specifications require boolean conditions, exact result agreement, exhaustive narrowing, and collection-element typing.
- `spec:puppydsl.language.contexts` derives provided-value and destructor types from callable signatures.
- `spec:puppydsl.language.events` derives generated callable contracts from payload types.
- `spec:puppydsl.language.program_validity.name_and_type_resolution` defines separate logical type and callable registries.
- `spec:puppydsl.language.program_validity.flow_analysis` requires lexical binding types, narrowing, and completion analysis.
- The user accepted deriving internal artifacts from these concrete consumers and explicitly requested that the callable catalog be included in the requirement inventory.

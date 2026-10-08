# Partial executable-design validation

## Status

This document records the current validation model for incomplete top-down YAML design.
Undefined declarations are not permitted. Typed stubs are the only supported incomplete function state.

## Need

Top-down executable design must permit an operation entry function to call semantic functions whose signatures and responsibilities are known before their bodies are implemented.

The runtime must preserve full declaration, type, context, call-graph, and traceability validation during that process.
Allowing unknown functions or types would block those checks and would make typos indistinguishable from intentional design gaps.

The supported incomplete state is therefore:

```text
declared and type-checkable
but not executable yet
```

rather than:

```text
undeclared and ignored
```

## Stub declaration

A stub is a complete function declaration with no executable body.

```yaml
# responsibility: Defines current-record selection functions used by executable Design Records operations.
# excludes: Reference parsing, repository discovery, operation response projection, and unrelated current-record behaviors.

functions:
  current-records.select-exact:
    detail: >-
      Selects one uniquely addressable current record for a canonical
      artifact reference.
    implements:
      - spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability
    signature: >-
      (reference: ArtifactReference) -> CurrentRecordSelection
    requires:
      - operation.current_identity_index
    implementation:
      stub: native
```

`implementation.stub` must be exactly one of `native`, `steps`, or `undecided`.
`native` and `steps` record a decided final implementation form. `undecided` records an unresolved native-versus-steps ownership decision and emits an authoring warning.
A stub must not contain `steps`, `return`, or `native`, and every stub remains non-executable.

A stub declaration must still satisfy every ordinary declaration rule that does not require an executable body, including:

- a valid declaration-file responsibility and exclusion header;
- stable function identity;
- nonempty semantic `detail`;
- complete typed signature;
- declared input and output types;
- valid `opens` context IDs and order;
- valid `requires` dependencies;
- valid traceability metadata;
- duplicate-ID and callable-namespace rules.

## Unknown declarations

Every referenced function, type, context, event, and module must be declared in every validation profile.

The validator rejects:

- a call to an undeclared function;
- a signature, field, union, event, or module slot that names an undeclared type;
- an unknown context ID in `opens`;
- an unknown provided-value ID in `requires` or provider arguments;
- any other unresolved declaration reference.

There is no `allow undefined` mode.
A planned function must first be declared as a typed stub before another function may call it.

## Validation profiles

### Authoring profile

Purpose: permit top-down refinement while preserving complete static design checks.

Behavior:

- unknown declarations are errors;
- typed stubs are allowed;
- `stub: native` and `stub: steps` are authoring gaps without an implementation-direction warning;
- `stub: undecided` emits one unresolved-implementation-direction warning;
- calls to stubs are type-checked normally;
- downstream field projections and union handling are validated from stub signatures;
- `opens`, `requires`, active-context routes, provider dependencies, and provider cycles are validated;
- call-graph cycles are validated across steps-defined and stub declarations where edges exist;
- traceability and generated projections include stub declarations;
- execution is prohibited when the selected reachable graph contains a stub.

A stub is reported as an authoring gap, not as a successful implementation.

### Design-closure profile

Purpose: prove that the selected executable design is structurally complete independent of native implementation activation.

Behavior:

- all authoring checks pass;
- no reachable stub is allowed;
- every function uses exactly one executable implementation form: `steps` or `native`;
- declarative bodies type-check;
- all field projections, union switches, call edges, context routes, and provider dependency graphs validate;
- every native leaf has one declared stable native binding ID;
- generated documentation and coverage are reproducible from the declaration set.

Whether an unavailable but declared native binding may remain at this profile is still pending.

### Activation profile

Purpose: prove that the selected declaration set can execute.

Behavior:

- all design-closure checks pass;
- every native binding resolves to an activated implementation;
- native signatures conform exactly;
- required host-owned contexts and entry routes are constructible;
- no stub exists in the activated reachable graph.

## Reachability boundary

Validation distinguishes the repository declaration inventory from the graph reachable from selected entry functions.

At minimum:

- authoring validation checks every loaded declaration for local correctness;
- design closure rejects stubs reachable from the selected entries;
- activation rejects stubs or unresolved bindings reachable from the selected entries.

Whether repository-wide design closure also rejects unreachable stubs remains pending.

## Context validation during authoring

Stub declarations participate fully in context analysis.

For each selected entry route, validation must derive the ordered active context chain from:

```text
host-opened contexts
  + each called function's ordered opens list
```

It must reject:

- unknown or duplicate context IDs in `opens`;
- reopening a context ID already active on the route;
- an `opens` order in which an earlier context depends on a later context;
- a direct or transitive `requires` dependency unavailable from the active chain;
- provider arguments that reference unavailable contexts;
- direct or indirect cycles in the provided-value dependency graph.

Stub status does not suppress these checks because the function's context contract is already declared.

## Diagnostic projection

Partial-design validation should report machine-readable gap kinds such as:

- `explicit_function_stub`;
- `stub_implementation_direction_undecided`;
- `stub_reachable_from_entry`;
- `stub_prohibits_execution`;
- `unknown_function`;
- `unknown_type`;
- `unknown_context`;
- `context_reopen`;
- `invalid_context_open_order`;
- `context_provider_cycle`;
- `unresolved_native_binding`.

Unknown declarations remain validation errors in every profile.
Stub diagnostics are authoring gaps in the authoring profile and errors in design-closure or activation profiles.

These are design-tool diagnostics. They are not DRMCP operation diagnostics and never appear in normal operation responses.

## Generated projection

Generated function pages for stubs show the complete declared contract together with:

```text
implementation status: stub
implementation intent: native | steps | undecided
execution: prohibited
closure: incomplete
```

The generated DAG may include call edges to stubs because their signatures are known.
Coverage reports may show their direct `implements` relations, but must distinguish declared coverage from executable evidence.

## Closure principle

```text
authoring
  -> every declaration exists
  -> typed stubs are allowed
  -> full static validation continues

refinement
  -> each stub becomes steps or native

design closure
  -> no reachable stubs

activation
  -> all native bindings and runtime dependencies resolve
```

The validator must never treat a stub as executable evidence or convert an unknown declaration into an allowed design gap.

# V01-TASK-DATA-015-01: Decide recursive reference and untagged-union representation boundary

- **id**: V01-TASK-DATA-015-01
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-DATA-015
- **source_requirement**: V01-REQ-DATA-008
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - Contract boundary decision for recursive named model references
  - Explicit decision to reject, defer, or constrain untagged union / general oneOf support
  - V01-ADR-073 relationship decision and successor ADR need assessment
  - Follow-up task split input for spec / implementation / fixture work if selected

## Goal

Decide the contract boundary for recursive references and untagged-union representation captured by `V01-REQ-DATA-008` / `V01-WORK-DATA-015`.

This task covers the `recursive / union structure` bucket from the UC-002 notes-retreat successor classification: N-009 and N-044.

## Work

- Review the source bucket evidence for N-009 and N-044.
- Confirm whether recursive references are allowed in brewprint data models.
- If recursive references are allowed, decide the minimum safe representation boundary, especially whether recursion is limited to named model references.
- Confirm whether untagged union / general `oneOf` / `anyOf` / scalar union support is introduced, explicitly deferred, or rejected.
- Check whether untagged-like surfaces can be represented by tagged union envelope models instead of adding general shape-inference union support.
- Decide whether V01-ADR-073 remains limited to tagged / discriminated unions or requires a successor / broadening ADR.
- Identify the follow-up spec, diagnostic, YAML, render, fixture, and implementation tasks needed after the contract decision.

## Included Scope

- Contract decision for recursive named model references.
- Contract decision for untagged union / general `oneOf` support.
- Boundary check against V01-ADR-073 tagged union support.
- Follow-up task split proposal if implementation or spec work is selected.

## Excluded Scope

- Parser, resolver, renderer, validator, MCP, or fixture implementation changes.
- Direct UC-002 YAML migration.
- Golden regeneration.
- Reopening M15 or completed DATA work items.
- Expanding V01-WORK-DATA-010 without an explicit boundary decision.

## Done condition

- Recursive reference handling is explicitly decided.
- Untagged union / general `oneOf` handling is explicitly decided.
- V01-ADR-073 relationship is classified as unchanged, broadened, or requiring a successor ADR.
- Follow-up work is identified, or an explicit no-action / defer outcome is recorded.
- `V01-WORK-DATA-015` can be moved forward from blocked planning state.

## Verification

- Confirm the decision is traceable to `V01-REQ-DATA-008`, `V01-WORK-DATA-015`, `V01-TASK-DATA-009-03`, and `V01-TASK-DATA-009-04`.
- Confirm the task does not perform implementation, YAML migration, or fixture regeneration.

## Evidence

Completed on 2026-06-07.

### Sources reviewed

- `V01-REQ-DATA-008`
- `V01-WORK-DATA-015`
- `V01-TASK-DATA-009-01`
- `V01-TASK-DATA-009-02`
- `V01-TASK-DATA-009-03`
- `V01-ADR-073`
- `docs/spec/type-ref.md`

### Decision summary

Recursive references are accepted only through named model references.

Untagged union / general `oneOf` / `anyOf` / scalar union support is not introduced.

V01-ADR-073 remains limited to tagged / discriminated unions. No V01-ADR-073 broadening is selected by this task.

### Recursive reference boundary

The selected boundary is:

- Recursive structure may be represented when a field uses an existing named model TypeRef.
- Inline recursive shapes are not introduced.
- Renderer / resolver follow-up must avoid recursive expansion and should display recursive fields as references instead of unfolding the full target shape.
- This decision is intended for cases such as recursive `object_ref.parent` from N-044.

Example boundary:

```yaml
id: object_ref
kind: struct
fields:
  parent:
    type: object_ref
    optional: true
```

The example is a contract illustration only. This task does not perform YAML migration.

### Untagged union boundary

The selected boundary is:

- Do not add untagged union / general `oneOf` / `anyOf` to TypeRef or model kind vocabulary.
- Do not add shape-inference variant selection to validator, resolver, renderer, or MCP schema generation.
- Do not add scalar union such as `str | int`.
- Do not add `SourceLocation | ObjectRef` style inferred union support.

Rationale:

- The confirmed customer surface is narrow: N-009, represented by `diagnostic.related`-like `SourceLocation` or `ObjectRef` lists.
- V01-ADR-073 already records that this is outside tagged-union minimum and that shape-inference union support is independently complex.
- Untagged union would require readers, validators, renderers, and LLM consumers to infer variant identity from shape, which conflicts with brewprint's traceability and human-review goals.
- The likely customer can be modeled with an explicit tagged union envelope when machine-readable schema is required.

### Replacement pattern for untagged-like surfaces

When a surface needs machine-readable schema and would otherwise require untagged union, model it as a tagged union envelope.

Example boundary:

```yaml
id: diagnostic_related
kind: tagged_union
discriminator: kind
variants:
  - tag: source_location
    fields:
      location:
        type: source_location
  - tag: object_ref
    fields:
      object:
        type: object_ref
```

If the surface is intentionally opaque or explanatory, it may remain `any + note` / prose instead of forcing a schema model.

### V01-ADR-073 relationship

V01-ADR-073 remains unchanged in scope:

- tagged / discriminated union support remains the accepted union mechanism;
- untagged union / general `oneOf` remains outside the model system;
- future work should not silently broaden V01-WORK-DATA-010 into untagged union support.

No successor ADR is required for untagged union rejection. A future ADR may be created only if a later concrete use case proves tagged envelopes insufficient.

### Follow-up work input

Recommended follow-up split after this task:

1. Spec update task: document recursive named model reference support and explicitly document that untagged union / general `oneOf` is unsupported.
2. Implementation task: update resolver / validator / renderer behavior for recursive named model references if current implementation rejects or recursively expands them incorrectly.
3. UC-002 cleanup task if selected: handle N-044 `object_ref.parent` after spec / implementation support exists.
4. Optional UC-002 cleanup task: decide whether N-009 should remain `any + note` or become a tagged union envelope model.

### Verification note

This task performed contract decision only.

No parser, resolver, renderer, validator, MCP implementation, UC-002 YAML migration, or fixture / golden regeneration was performed.

Traceability preserved:

- N-009 / N-044 remain traceable through `V01-REQ-DATA-008` and `V01-WORK-DATA-015`.
- Tagged union scope remains owned by V01-ADR-073 / `V01-REQ-DATA-004` / `V01-WORK-DATA-010` without silent expansion.

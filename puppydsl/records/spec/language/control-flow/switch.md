# Reference: PuppyDSL exhaustive switch

- **id**: `spec:puppydsl.language.control_flow.switch`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.control_flow`

## What this is

Defines exhaustive switching over one previously bound union, enum, or `optional<T>` value.

## Authoring form

```yaml
<result-step>:
  switch: <target-binding>
  cases:
    <case-key>:
      steps:
        <ordered steps>
      return: <binding>
```

The target is one visible binding. Inline calls, projections, constructors, and other expressions are invalid.

## Union switch

Each case key names one variant type directly declared by the target union.

Inside one case, the target binding keeps its name and narrows to the named variant type.

Every union variant appears exactly once. Unknown, missing, duplicate, and `default` cases are invalid.

## Enum switch

Each case key is one qualified enum member literal.

```yaml
structure_result:
  switch: structure
  cases:
    RecordStructure.tree:
      steps:
        result: example.for_tree()
      return: result
    RecordStructure.sequential:
      steps:
        result: example.for_sequential()
      return: result
```

Every declared enum member appears exactly once.

The target retains its enum type inside each case. Member selection does not create a member-specific type.

## Optional switch

An `optional<T>` switch contains exactly these cases:

```text
present
absent
```

Inside `present`, the target keeps its name and narrows to `T`.

Inside `absent`, the target is unavailable in value positions because no `T` exists.

Outside the switch, the original target remains `optional<T>`.

## Case completion

| case shape | result rule |
|---|---|
| No continuing case returns | Complete switch is `void`. |
| Every case terminates through `never` | Complete switch is `never`. |
| Continuing cases return values | Every continuing case returns the same exact type. |
| Case terminates through `never` | Case does not participate in result agreement. |
| Case terminates through `continue` inside `for` | Case contributes no iteration output. |

PuppyDSL does not infer a common union or optional from different case return types.

## Scope

Every case is an independent child scope.

Union and optional narrowing is case-local. The narrowed type does not persist after the switch.

## Validation rules

- Reject a target that is not a previously bound union, enum, or optional value.
- Reject inline expression targets.
- Reject missing, unknown, duplicate, or `default` cases.
- Reject an unqualified or wrong-enum member case.
- Reject optional cases other than exactly `present` and `absent`.
- Reject any value use of the optional target in `absent`.
- Reject incompatible continuing case results.
- Reject use of a case-local binding outside its case.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.control_flow` | Parent control-flow overview. |
| `spec:puppydsl.language.type_system.declared_types` | Defines enum and union case sets. |
| `spec:puppydsl.language.type_system.assignability` | Defines case result compatibility. |
| `spec:puppydsl.language.expressions.construction_and_access` | Applies case-local narrowing to access. |

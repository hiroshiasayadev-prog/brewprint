# Reference: PuppyDSL conditional execution

- **id**: `spec:puppydsl.language.control_flow.conditional`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.control_flow`

## What this is

Defines `if` branching over one previously bound boolean value.

## Authoring form

```yaml
<result-step>:
  if:
    condition: <boolean-binding>
    then:
      steps:
        <ordered steps>
      return: <binding>
    else:
      steps:
        <ordered steps>
      return: <binding>
```

`condition` names one visible binding of exact type `boolean`. Inline calls and expressions are invalid.

## Branch completion

| branch shape | result rule |
|---|---|
| No branch returns and one or more paths continue | Complete `if` is `void`. |
| Every possible branch terminates through `never` | Complete `if` is `never`. |
| Continuing branches return values | Every continuing branch returns the same exact type. |
| One branch terminates through `never` | Other continuing branches determine the result type. |
| One branch uses explicit `continue` inside `for` | Other value-producing branches determine the iteration contribution type. |
| An information-bearing `if` used as a bounded iteration result omits one branch | The value-producing branch determines `T`; the omitted branch produces `IterationResultContributionAbsent`. |

A non-void `if` outside a bounded iteration requires an explicit `else`.

Inside a bounded `for`, an information-bearing `if` used as the per-iteration result may omit `else`. Its logical result is `T | IterationResultContributionAbsent`. `IterationResultContributionAbsent` is a control-flow result state, not a declared union type, `void`, or an implicit `continue` operation.

The omitted branch completes normally without producing the per-iteration result binding. The enclosing `for` interprets that state by contributing no element for the current iteration.

A `void` `if` may omit `else`. The false path is a no-op.

## Scope

`then` and `else` are independent child scopes.

Branch locals do not escape. The complete information-bearing `if` result enters the enclosing scope through its step key.

## Restrictions

`if` does not perform:

- union or optional narrowing;
- enum member dispatch;
- inline comparison or arithmetic;
- boolean composition;
- runtime type testing.

Use exhaustive `switch` for union, enum, and optional values.

## Validation rules

- Reject a condition that is unknown, not previously bound, or not `boolean`.
- Reject a non-void `if` outside `for` with omitted `else`.
- Reject an omitted branch as an absent result contribution unless the complete `if` supplies a bounded iteration's per-iteration result.
- Reject incompatible value-producing branch result types.
- Reject a continuing branch that fails to establish the required result.
- Reject `continue` outside the permitted iteration position.
- Reject a later step after terminal `never` or `continue` in the same branch.

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.control_flow` | Parent control-flow overview. |
| `spec:puppydsl.language.control_flow.iteration` | Defines collection of contributing values, omission of absent iteration-result contributions, and explicit `continue`. |
| `spec:puppydsl.language.functions.steps_and_bindings` | Defines branch lexical scope. |

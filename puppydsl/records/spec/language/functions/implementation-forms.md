# Reference: PuppyDSL implementation forms

- **id**: `spec:puppydsl.language.functions.implementation_forms`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language.functions`

## What this is

Defines the mutually exclusive native, steps-defined, and stub implementation forms behind one declared function contract.

## Implementation forms

| form | declaration | meaning |
|---|---|---|
| Native | `implementation.native` | An activated native binding supplies the implementation. |
| Steps | `implementation.steps` | Ordered PuppyDSL steps compose declared functions and expressions. |
| Stub | `implementation.stub` | Authoring-only typed placeholder with no executable implementation. |

Exactly one implementation form is required.

## Native form

```yaml
implementation:
  native: text.trim
```

The native value is one non-empty binding ID. The binding ID identifies an activated implementation registry entry.

The binding ID is not a source path, package name, or native-language symbol.

The activated binding must expose the exact declared function signature.

## Steps form

```yaml
implementation:
  steps:
    normalized: common.text.trim(body)
    empty: common.text.is_empty(normalized)
  return: empty
```

A steps implementation contains a non-empty ordered `steps` mapping. `return` is governed by the declared output and step-path rules.

## Stub form

A stub records the accepted implementation direction.

| stub value | meaning | authoring validation |
|---|---|---|
| `native` | Final implementation is native but not authored. | Valid without warning. |
| `steps` | Final implementation is steps-defined but not authored. | Valid without warning. |
| `undecided` | Implementation ownership remains unresolved. | Valid with one warning. |

```yaml
implementation:
  stub: native
```

A stub cannot contain `native`, `steps`, or `return`. Every stub is non-executable.

## Validation profiles

| profile | stub treatment |
|---|---|
| Authoring | Accept `native` and `steps`; accept `undecided` with warning. |
| Design closure | Reject every reachable stub. |
| Activation | Reject every activated stub. |

Validation profile names define validity conditions. Diagnostic transport and command behavior belong to validator specifications.

## `void` output

A `void` function completes normally without an information-bearing result.

A direct `void` call uses a discard step label and does not create a local binding. A steps-defined `void` function omits `return`.

## `never` output

A `never` function has no normally continuing result path.

A direct `never` call:

- uses a terminal discard step label;
- creates no local binding;
- prohibits later reachable steps on that lexical path.

A steps-defined `never` function omits `return`. Every reachable path must terminate through another `never` call.

PuppyDSL exposes no catch, recovery, or value-level inspection of execution failure.

## Invalid combinations

| condition | result |
|---|---|
| More than one implementation form | Invalid. |
| No implementation form | Invalid. |
| Native or stub form contains `return` | Invalid. |
| Stub uses an unsupported scalar or boolean | Invalid. |
| Steps mapping is empty | Invalid. |
| Native binding ID is empty or unresolved | Invalid for activation. |
| Native binding signature differs | Invalid for activation. |

## Boundary

| owned here | owned elsewhere |
|---|---|
| Language meaning of native binding ID. | Native registry implementation. |
| Stub intent and profile validity. | Diagnostic code and output format. |
| `void` and `never` implementation behavior. | Host execution-failure envelope. |
| Mutual exclusion of implementation forms. | Code generation and native representation mapping. |

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language.functions` | Parent function overview. |
| `spec:puppydsl.language.functions.steps_and_bindings` | Defines steps implementation semantics. |
| `spec:puppydsl.language.type_system` | Defines `void` and `never`. |
| `spec:puppydsl.language.program_validity` | Applies profile and activation validity. |

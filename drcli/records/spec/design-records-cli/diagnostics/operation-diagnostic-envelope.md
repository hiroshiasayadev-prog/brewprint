# Contract: Operation diagnostic envelope

- **id**: `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_envelope`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.diagnostics`
- **contract_class**: `format`

## What this is

Defines the transport-independent logical shape and placement rules for semantic DRCLI operation errors and warnings.

## Current contract

An operation diagnostic uses this logical shape:

```text
code
message
context?
```

| field | presence | value |
|---|---|---|
| `code` | always | Stable lowercase snake_case identifier registered by the operation diagnostic catalog. |
| `message` | always | Non-empty human-readable English explanation. |
| `context` | only when the code requires or permits context | Object whose fields follow the selected code-specific context contract. |

The same diagnostic object shape is used beneath `error` and within `warnings`.
Placement determines whether the diagnostic is an error or warning.
The object does not contain `severity` or `operation`.
The logical shape does not prescribe CLI serialization, stream placement, or rendering.

## Non-goals

- Defining CLI tokenization, flag parsing, command selection, or argument decoding failures.
- Defining help text, usage text, terminal formatting, stdout or stderr placement, or exit-status policy.
- Defining validation findings, validation classification, or artifact-specific conformance rules.
- Defining internal exceptions, debug output, logs, or telemetry.

## Rules

### Code

Each semantic operation diagnostic code must exist in `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog`.
Each code has one fixed class: `error` or `warning`.
An implementation must not expose implementation-specific or unregistered semantic operation codes.
An unexpected operation failure maps to `execution_failure` when no more specific accepted outcome applies.

### Message

`message` explains the specific diagnostic for a human reader.
Exact wording is not a stable machine-readable contract.
A consumer must not parse `message` for branching, identity, ordering, or equality.
A consumer uses `code` and `context` for machine processing.
The same code may use different message wording for different accepted context variants.

### Context

The catalog defines required, optional, and prohibited context fields for each code and context variant.
A context object contains only fields allowed by the matching catalog entry.
An implementation must not add unknown context fields.
Object field order is not normative.
A required context field must not use `null` as a substitute for an unavailable value.

### Placement

| response state | top-level fields | selector-result diagnostic fields |
|---|---|---|
| Complete response without request-wide warnings | Normal operation result fields. | Operation-specific success or selector-error state. |
| Complete response with request-wide warnings | Normal operation result fields and top-level `warnings`. | Operation-specific success or selector-error state. |
| Selector success with warnings | Normal response fields. | Success fields and `warnings`. |
| Selector-level error | Normal response fields and the ordered `results` collection. | Fields established before failure and exactly one `error`. |
| Request-wide error | Exactly one top-level `error`. | No selector results. |

The shared logical placements are:

```text
error:
  code
  message
  context?
```

```text
<normal-response-fields>
warnings:
  - code
    message
    context?
```

```text
results:
  - <supplied-selector-fields>
    <fields-established-before-error>
    error:
      code
      message
      context?
```

```text
results:
  - <supplied-selector-fields>
    <success-fields>
    warnings:
      - code
        message
        context?
```

A top-level warning may coexist with selector-level errors in `results`.
A top-level `error` must not coexist with normal result fields, top-level `warnings`, or `results`.
A selector-level `error` must not coexist with selector success fields or selector-level `warnings`.
An operation Specification defines every allowed warning placement.
The envelope does not create a general warning surface for operations that define none.

### Authoring negative outcome placement

All authoring errors use exactly one top-level `error` object from this envelope, with code and context fixed by `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog`. A trustworthy non-writing semantic rejection is a complete **negative operation result**, serialized as the single `error` document on stdout in JSON mode (exit 1). It has no `results`, `warnings`, or unrelated normal-success fields. An unavailable retained-state, preparation, persistence, or write-completion condition that prevents a trustworthy operation result is the single `error` document on stderr (exit 3). A cache preservation failure accompanying the original authoring error is represented by the nested cataloged `cache_error` in its context, not a second document or independent process status. A successfully preserved retry body is represented only by the `body_cache_id` in that same context. This is also true for errors after body receipt in deferred preparation or write. The presence of a retry ID never implies a successful repository write.

The distinct `authoring_cache_preservation_failed` code is valid only as the nested `cache_error` for a primary authoring error. The `cache_error` object contains exactly `code` and `message`; it has no `context`. Authoring context follows the catalog's exact code-specific fields and conditional retry fields. The enclosing error determines stream and exit category. An authoring failure is never represented as a selector-level error, validation finding, optional warning, or implementation-defined JSON shape.

### Omission

An absent diagnostic surface is omitted.
The response must not use these empty representations:

- `error: null`;
- `warnings: []`;
- `context: null`;
- `context: {}`.

A diagnostic omits `context` when the catalog prohibits context for its code.
A complete response or selector result omits `warnings` where no warning applies.

### Warning order

Warning collection order is normative.
The consuming operation Specification defines deterministic warning order.
Diagnostics do not apply a global code sort.
An operation that introduces multiple warnings defines their relative order before exposing them.

### Request-wide error selection

A response contains at most one request-wide error.
Each operation Specification defines a deterministic request-validation and execution order.
The operation returns the first request-wide error established by that order.
Diagnostics do not define one global precedence across operations.

Current-state-independent semantic request validation precedes current-state construction.
Current-state availability validation follows successful current-state construction.
Selector-level processing begins only after request-wide semantic validation succeeds.

### CLI usage boundary

This envelope begins after DRCLI has a type-valid logical operation request.
CLI syntax, tokenization, command lookup, flag recognition, missing CLI operands, and argument decoding belong to the command and runtime contract.
Those failures are CLI usage diagnostics, not semantic operation diagnostics defined here.
This Specification does not map CLI usage diagnostics to operation codes.
Concrete help and usage rendering belongs to `spec:drcli.design_records_cli.cli`.

An operation may define selector-presence or selector-combination rules as semantic validation when the logical operation request admits that shape.
The consuming operation Specification owns that decision.

## Validation rules

| condition | result |
|---|---|
| `code` is absent or is not cataloged | Invalid diagnostic object. |
| `message` is absent, empty, or not a string | Invalid diagnostic object. |
| `severity` or `operation` is present | Invalid diagnostic object. |
| `context` is present for a context-prohibited code | Invalid diagnostic object. |
| A required context field is absent | Invalid diagnostic object. |
| An unknown context field is present | Invalid diagnostic object. |
| A code appears under a placement not allowed by its cataloged class and placement | Invalid diagnostic placement. |
| A top-level `error` coexists with normal result fields, `results`, or top-level `warnings` | Invalid response state. |
| A selector-level `error` coexists with selector success fields or selector-level `warnings` | Invalid selector-result state. |
| An empty or null diagnostic surface is present | Invalid response state. |

## Boundary

| concern | owner |
|---|---|
| Diagnostic object fields, placement, omission, exclusivity, and common validation rules | This Specification. |
| Canonical semantic operation codes and code-specific context fields | `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog`. |
| Trigger condition, semantic classification, request-wide or selector-level scope, and warning order | The consuming operation Specification. |
| Normal operation result fields and indicators | The consuming operation Specification. |
| Validation finding semantics and artifact conformance rules | The canonical Specification that owns the validated rule. |
| CLI syntax, parser diagnostics, help or usage rendering, stream placement, and exit status | `spec:drcli.design_records_cli.cli`. |
| Logging, telemetry, stack traces, and internal exceptions | Implementation contracts. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog` | Defines every semantic operation code and context contract accepted by this envelope. |
| `spec:drcli.design_records_cli.operations.validation` | Defines validation triggers, result placement, and validation-specific warning order. |
| `spec:drcli.design_records_cli.current_record_model` | Defines addressability, scope availability, and trustworthy current-state boundaries referenced by operation outcomes. |
| `spec:drcli.design_records_cli.cli` | Owns CLI usage diagnostics and concrete CLI rendering outside this semantic envelope. |

# Contract: Operation diagnostic envelope

- **id**: `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_envelope`
- **status**: draft
- **date**: 2026-07-14
- **parent**: `spec:drmcp.design_records_mcp.diagnostics`
- **contract_class**: `format`

## What this is

Defines the shared external object shape and placement rules for DRMCP operation errors and warnings.

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
The placement determines whether the diagnostic is an error or warning.
The object does not contain `severity` or `operation`.

## Rules

### Code

Each external diagnostic code must exist in `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_catalog`.
Each code has one fixed class: `error` or `warning`.
An implementation must not expose implementation-specific or unregistered codes.
An unexpected internal failure maps to `execution_failure` when no more specific accepted operation outcome applies.

### Message

`message` explains the specific diagnostic for a human reader.
Exact wording is not a stable API contract.
A caller must not parse `message` for branching, identity, ordering, or equality.
A caller must use `code` and `context` for machine processing.
The same code may use different message wording for different accepted context variants.

### Context

The catalog defines required, optional, and prohibited context fields for each code and context variant.
A context object must contain only fields allowed by the matching catalog entry.
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
An operation Specification must define every allowed warning placement.
The envelope does not create a general top-level warning surface for operations that define none.

### Omission

An absent diagnostic surface is omitted.
The response must not use these empty representations:

- `error: null`;
- `warnings: []`;
- `context: null`;
- `context: {}`.

A diagnostic omits `context` when the catalog prohibits context for its code.
A complete response or selector result omits `warnings` at each placement where no warning applies.

### Warning order

Warning collection order is normative.
The consuming operation Specification defines deterministic warning order.
Diagnostics do not apply a global code sort.
An operation that introduces multiple warnings must define their relative order before exposing them.

### Request-wide error selection

A response contains at most one request-wide error.
Each operation Specification defines a deterministic request-validation and execution order.
The operation returns the first request-wide error established by that order.
Diagnostics do not define one global precedence across operations.

Current-state-independent request validation precedes current-state construction.
Current-state availability validation follows successful current-state construction.
Selector-level processing begins only after request-wide validation succeeds.

### Input-schema boundary

MCP input-schema violations occur before this diagnostic envelope applies.
Schema violations include omissions rejected by the input schema, unknown fields, wrong value types, and malformed schema-owned object shapes.
An operation may define selector-presence rules as semantic validation when its input schema accepts the incomplete selector set.
Operation diagnostics represent only semantic outcomes accepted by the operation Specifications.

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
| Canonical codes and code-specific context fields | `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_catalog`. |
| Trigger condition, semantic classification, request-wide or selector-level scope, and warning order | The consuming operation Specification. |
| Normal operation result fields and indicators | The consuming operation Specification. |
| Artifact conformance findings, physical paths, repair guidance, validation severity, and validation aggregation | Validation Specifications. |
| Transport, JSON-RPC, MCP SDK, logging, telemetry, stack traces, and internal exceptions | Implementation and transport contracts. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_catalog` | Defines every code and context contract accepted by this envelope. |
| `spec:drmcp.design_records_mcp.operations.discovery_and_listing` | Defines completed discovery and listing triggers and result placement. |
| `spec:drmcp.design_records_mcp.operations.retrieval` | Defines completed retrieval triggers, batch results, and warning order. |
| `spec:drmcp.design_records_mcp.current_record_model` | Defines addressability, scope availability, and trustworthy current-state boundaries referenced by operation outcomes. |

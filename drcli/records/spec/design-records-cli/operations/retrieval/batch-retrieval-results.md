# Contract: Batch retrieval results

- **id**: `spec:drcli.design_records_cli.operations.retrieval.batch_retrieval_results`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations.retrieval`
- **contract_class**: `interface`
## What this is

Defines shared batch-result behavior for `get_records` and `get_record_sections`.

## Non-goals

- Operation-specific selector syntax.
- Operation-specific success fields, warning classifications, or error classifications.
- Source-content budget measurement and allocation.
- Shared diagnostic object, code, context, or message rules.

## Request

Each retrieval request supplies one required ordered selector collection.
The operation-specific Specification defines the collection field and selector syntax.

| constraint | value |
|---|---|
| minimum selector count | `1` |
| maximum selector count | `100` |

A selector count outside this range produces request-level `invalid_selector_count`.
Selector processing does not begin after this error.
The response contains no `results` or top-level `warnings` after this error.

The error context includes:

| field | value |
|---|---|
| `minimum` | `1` |
| `maximum` | `100` |
| `actual` | Supplied selector count. |

## Response

A normal complete response contains one ordered `results` collection.

The following invariants apply:

- Each supplied selector occurrence produces exactly one result item.
- Result count equals supplied selector count.
- Result order equals supplied selector order.
- Each result preserves the exact supplied selector string.
- Duplicate selector occurrences remain separate.
- DRCLI does not deduplicate or reorder selectors.
- One selector-level error does not stop later selector processing.

A result item uses the selector field defined by its operation:

| operation | selector field |
|---|---|
| `get_records` | `record_ref` |
| `get_record_sections` | `section_selector` |

## Result states

### Success without warning

A successful result contains its operation-specific success fields.
The result omits `error` and `warnings` when no warning applies.
The result does not contain a separate success outcome field.

### Success with warnings

A successful selector result contains `warnings` only when one or more selector-level warnings apply.
The collection contains only applicable warnings.
The operation-specific Specification defines warning classifications and context.

A normal response may contain top-level `warnings` only when the operation Specification defines a request-wide warning.
A top-level warning may coexist with selector-level errors in `results`.

When multiple warnings apply at one placement, the operation-specific Specification defines their deterministic order.

### Selector-level error

An unsuccessful selector result contains exactly one `error`.
The result omits fields that were not established before the error condition.
The operation-specific Specification defines the error classification and available context.

A selector-level error does not remove successful results for other selectors.

## Errors

A request-wide configuration or execution failure produces a top-level `error`.
The response contains no partial `results` collection or top-level `warnings`.

A request-level semantic error defined by a retrieval Specification follows the same no-results boundary.
Logical input-shape violations are rejected before this result contract applies.

Shared diagnostic objects, canonical codes, context fields, placement, and messages belong to Diagnostics Specifications.
Retrieval Specifications define only triggering conditions, classifications, ordering, and operation-specific context.

## Boundary

| concern | owner |
|---|---|
| Selector count and selector-to-result correspondence | This Specification. |
| Partial selector success and continuation | This Specification. |
| Request-wide no-partial-result boundary | This Specification. |
| Exact record selectors and result fields | `spec:drcli.design_records_cli.operations.retrieval.exact_record_retrieval`. |
| H2 section selectors and result fields | `spec:drcli.design_records_cli.operations.retrieval.h2_section_retrieval`. |
| Source-content budget behavior | `spec:drcli.design_records_cli.operations.retrieval.retrieval_output_limits`. |
| Shared diagnostic representation | `spec:drcli.design_records_cli.diagnostics`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.operations.retrieval.exact_record_retrieval` | Applies this contract to exact record selectors. |
| `spec:drcli.design_records_cli.operations.retrieval.h2_section_retrieval` | Applies this contract to H2 section selectors. |
| `spec:drcli.design_records_cli.current_record_model` | Defines unresolved and conflicted exact refs. |
| `spec:drcli.design_records_cli.diagnostics` | Defines shared response-state and diagnostic object rules. |
| `spec:drcli.design_records_cli.diagnostics` | Defines canonical batch and selector diagnostic codes. |
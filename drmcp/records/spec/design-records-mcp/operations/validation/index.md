# Overview: Validation operations

- **id**: `spec:drmcp.design_records_mcp.operations.validation`
- **status**: draft
- **date**: 2026-07-14
- **parent**: `spec:drmcp.design_records_mcp.operations`

## What this is

Defines broad-scope validation and detailed validation of exact current records.

## Current contract

| operation | purpose |
|---|---|
| `validate_scope` | Validate one repository, app, artifact-kind, sequential-domain, or tree-subtree scope and return aggregates with abnormal subject summaries. |
| `validate_records` | Validate one or more exact current canonical refs and return detailed findings for every successfully selected current record. |

Each operation name is also the MCP tool name.
Broad-scope validation and exact-record validation are independent operations.

## Non-goals

- Defining artifact-specific validation rules.
- Defining shared validation rule components or finding-code catalogs.
- Defining metadata-field, H2-section, reference, or structure locator schemas.
- Repairing, completing, normalizing, or inferring nonconforming source content.
- Selecting unadmitted sources or tree nodes by physical source path.
- Parser, cache, storage, watcher, concurrency, or filesystem architecture.

## Boundary

| concern | owner |
|---|---|
| Current source, admission, addressability, scope, and tree-node semantics | `spec:drmcp.design_records_mcp.current_record_model`. |
| Base and artifact-specific conformance rules | `spec:drmcp.design_records_mcp.artifacts`. |
| Broad-scope request, abnormal summaries, aggregates, and operation outcomes | `spec:drmcp.design_records_mcp.operations.validation.scope_validation`. |
| Exact-record request, selector outcomes, and detailed finding placement | `spec:drmcp.design_records_mcp.operations.validation.exact_record_validation`. |
| Broad-result and exact-finding output limits | `spec:drmcp.design_records_mcp.operations.validation.validation_output_limits`. |
| Shared operation errors and warnings | `spec:drmcp.design_records_mcp.diagnostics`. |
| Finding codes, affected-subject locators, rule composition, and artifact validation mappings | Other Validation Specifications. |

Validation operations select subjects and project outcomes.
Applicable validation contracts define evaluation dependencies, rules, findings, and lookup behavior.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Scope validation | Contract | `spec:drmcp.design_records_mcp.operations.validation.scope_validation` | Defines `validate_scope`, broad selectors, abnormal result summaries, aggregates, ordering, and request-wide outcomes. |
| Exact record validation | Contract | `spec:drmcp.design_records_mcp.operations.validation.exact_record_validation` | Defines `validate_records`, exact selector handling, detailed finding placement, duplicate handling, and selector-level errors. |
| Validation output limits | Contract | `spec:drmcp.design_records_mcp.operations.validation.validation_output_limits` | Defines request-wide broad-result and exact-finding limits without partial result objects. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.current_record_model` | Supplies current sources, scopes, tree nodes, and uniquely addressable current records. |
| `spec:drmcp.design_records_mcp.artifacts` | Supplies base and artifact-specific contracts consumed by validation. |
| `spec:drmcp.design_records_mcp.diagnostics` | Supplies shared operation diagnostic objects and canonical codes. |

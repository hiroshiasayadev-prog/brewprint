# DRMCP-WORK-MCP-006: Validation, diagnostics, and path-exposure contract realignment

- **id**: DRMCP-WORK-MCP-006
- **status**: not_started
- **date**: 2026-06-27
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-INV-MCP-002
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-004
  - DRMCP-WORK-MCP-005
  - DRMCP-TASK-MCP-001-06
  - spec:product.design_records.traceability
  - spec:product.design_records.spec_format
  - spec:drmcp.design_records_mcp.responsibility_boundary
  - spec:drmcp.design_records_mcp.schema.diagnostics
  - spec:drmcp.design_records_mcp.schema.record_model
  - spec:drmcp.design_records_mcp.tools.list_records
  - spec:drmcp.design_records_mcp.tools.get_records
  - spec:drmcp.design_records_mcp.tools.validate_records
- **tasks**: []

## Goal

Establish the corrected current-repository validation, machine-readable diagnostic, and exceptional physical-path exposure contracts.

Separate PRODUCT-owned semantic invalidity from DRMCP-owned validation execution and response representation.

## Boundary

This Work Item owns:

- current active-index repository validation scope;
- current cross-namespace relation validation;
- configured current-to-legacy relation validation;
- behavior when a referenced legacy target is unavailable because fallback is disabled;
- exclusion of legacy archive records from normal current repository validation;
- mapping semantic invalidity into DRMCP diagnostic representation;
- shared machine-readable diagnostic structure and category placement;
- representation of unsupported, unresolved, duplicate, disabled-fallback, and source-format failures;
- narrow path exposure for source-location diagnostics, explicit patch output, and debug or emergency inspection;
- synchronization of validation, diagnostic, exceptional-path, and responsibility contracts;
- scoped validation and independent review for this contract boundary.

This Work Item does not own:

- canonical identity, relation, traceability, or spec-format semantics;
- discovery conditions or duplicate detection rules owned by `DRMCP-WORK-MCP-003`;
- tool-specific query and retrieval behavior owned by `DRMCP-WORK-MCP-004`;
- normal listing and exact-retrieval response projection, including physical-path hiding, owned by `DRMCP-WORK-MCP-004`;
- resolution order and fallback eligibility owned by `DRMCP-WORK-MCP-005`;
- individual warning-trigger behavior already owned by an operation contract;
- PRODUCT validation-policy owner-pointer updates;
- disposition of `DRMCP-WORK-SPEC-001` or `DRMCP-WORK-SPEC-002`;
- fixture authoring;
- validator, diagnostic, or response implementation;
- authoring transaction behavior.

Diagnostic representation may cite PRODUCT semantic authorities.
Diagnostic representation must not redefine those authorities.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Source Requirement for validation execution, diagnostics, and path exposure. |
| `DRMCP-ADR-MCP-001` | Governs current validation scope, archive exclusion, and path hiding. |
| `DRMCP-INV-MCP-002` | Supplies validation, diagnostic, and response-boundary findings. |
| PRODUCT traceability and spec-format authorities | Supply semantic validity without defining DRMCP response representation. |
| `DRMCP-WORK-MCP-003` | Supplies active-index, duplicate, and source-format outcomes. |
| `DRMCP-WORK-MCP-004` | Supplies normal listing, retrieval, and operation warning boundaries. |
| `DRMCP-WORK-MCP-005` | Supplies resolver outcomes and fallback states. |
| `spec:drmcp.design_records_mcp.schema.diagnostics` | Correct common diagnostic representation and category ownership. |
| `spec:drmcp.design_records_mcp.tools.validate_records` | Correct current repository validation behavior. |
| Normal list and retrieval contracts | Supply W004-owned normal physical-path hiding; W006 must not redefine their response projection. |
| DRMCP responsibility boundary | Synchronize semantic-owner and execution-owner statements. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Authority and affected-file confirmation | `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-005` | Confirm semantic authorities, operation inputs, and T07 exclusions. |
| B. Validation execution contract | Phase A | Define current repository, cross-namespace, and current-to-legacy validation scope. |
| C. Diagnostic representation contract | Phases A-B | Define common machine-readable representation and semantic-invalidity mapping. |
| D. Exceptional path-exposure contract | Phases A-C | Define source-location diagnostic, patch, debug, and emergency exposure without redefining W004-owned normal response hiding. |
| E. Cross-spec synchronization and review | Phases B-D | Synchronize affected contracts, validate, review, correct, and close. |

Validation Work Item disposition proceeds through a separate cross-owner gate.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Confirm affected specs, semantic authorities, W003-W005 inputs, and the T07 disposition boundary. | W003-W005 boundaries available. |
| T02 | Correct current repository and relation-validation execution contracts. | T01. |
| T03 | Correct machine-readable diagnostic representation and semantic-invalidity mapping. | T01-T02. |
| T04 | Correct source-location diagnostic, patch, debug, and emergency path-exposure exceptions while preserving W004-owned normal response hiding. | T02-T03. |
| T05 | Synchronize contracts, confirm no T07 overlap, validate, review, correct, and close. | T04. |

## Completion Condition

This Work Item is complete when all of the following are true:

- current repository validation operates on the active index only;
- current cross-namespace relations resolve across configured current roots;
- configured current-to-legacy relations have explicit validation behavior;
- disabled legacy fallback produces an explicit diagnostic outcome;
- legacy archive records are excluded from normal current repository validation;
- semantic invalidity and DRMCP diagnostic representation remain separate;
- required failures have machine-readable diagnostics;
- W004 remains the sole owner of physical-path hiding in normal listing and exact-retrieval responses;
- W006 path exposure is limited to explicit source-location diagnostics, patch output, and debug or emergency surfaces;
- T07 owns validation Work Item disposition and PRODUCT owner-pointer synchronization;
- fixtures and implementation remain delegated;
- all changed specs pass scoped validation;
- independent review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records changed files, validation results, review verdict, and residual limitations.

## Evidence

- `DRMCP-ADR-MCP-001`: Accepted validation-scope, diagnostic-ownership, and path-hiding direction.
- `DRMCP-REQ-MCP-001`: Source Requirement.
- `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-005`: Upstream contract owners.
- `DRMCP-TASK-MCP-001-06`: Hub lifecycle gate for this Work Item.
- 2026-06-27 planning-record correction: removed overlapping ownership of normal list and exact-retrieval path hiding; W004 remains the operation-response authority and W006 owns only exceptional path exposure.
- Contract correction, validation, and independent-review evidence: pending Task execution.

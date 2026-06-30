# DRMCP-WORK-MCP-010: Configured legacy archive fallback implementation

- **id**: DRMCP-WORK-MCP-010
- **status**: blocked
- **date**: 2026-06-30
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-INV-MCP-002
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-005
  - DRMCP-WORK-MCP-006
  - DRMCP-WORK-MCP-008
  - DRMCP-WORK-MCP-009
  - DRMCP-WORK-MCP-011
  - DRMCP-WORK-MCP-012
  - DRMCP-TASK-MCP-001-10
  - PRODUCT-WORK-SPEC-014
  - spec:drmcp.implementation
  - spec:product.brewprint.compatibility.legacy_id_compatibility
  - spec:drmcp.design_records_mcp.resolver
  - spec:drmcp.design_records_mcp.schema.diagnostics
  - spec:drmcp.design_records_mcp.schema.discovery
  - spec:drmcp.design_records_mcp.tools.get_records
  - spec:drmcp.design_records_mcp.tools.resolve_reference
  - spec:drmcp.design_records_mcp.tools.validate_records
- **tasks**: []

## Goal

Add an optional read-only legacy archive fallback layer to the completed W012 current runtime.

Keep Brewprint history configuration-gated, separately indexed, and absent from normal current operations.

## Boundary

This Work Item owns:

- explicit `legacy_roots` configuration loading;
- disabled fallback when no legacy roots are configured;
- invalid, missing, duplicate, and overlapping root handling;
- legacy archive parsing for approved sequential artifact kinds;
- separate legacy archive index construction and duplicate detection;
- exact retrieval for approved V01 sequential IDs;
- fallback resolution only after current resolution returns unresolved;
- preservation of issued legacy IDs;
- configured current-to-legacy relation validation;
- disabled-fallback, unsupported-family, and unresolved-legacy diagnostics;
- read-only enforcement for legacy archive records;
- exclusion of legacy records from normal listing, current validation, authoring targets, and the active index;
- accepted, rejected, disabled-fallback, regression, and leakage tests using W008 fixtures;
- implementation review and closure evidence.

This Work Item does not own:

- accepted legacy-family semantics;
- resolver, fallback, configuration, diagnostic, or path-exposure contract design;
- legacy fixture design or fixture authoring;
- current parser, active-index, query, exact-current-retrieval, or current validation implementation;
- `V01-SPEC-*` compatibility;
- app-prefixless IDs, path inference, or fuzzy normalization;
- legacy-to-current ID translation or migration;
- archive content modification or archive relocation;
- automatic repository `v01/` discovery;
- authoring transaction behavior.

Without configured `legacy_roots`, runtime behavior must remain equivalent to the completed W012 current runtime.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Source Requirement for optional legacy archive fallback. |
| `DRMCP-ADR-MCP-001` | Governs configuration gating, accepted families, separate indexes, and read-only behavior. |
| `DRMCP-INV-MCP-002` | Supplies legacy resolver, fixture, implementation, and leakage findings. |
| `PRODUCT-WORK-SPEC-014` | Supplies the accepted `V01-SPEC-*` removal prerequisite. |
| Brewprint legacy compatibility spec | Supplies approved sequential families and archive facts. |
| `DRMCP-WORK-MCP-005` | Supplies resolver and configured-fallback contracts. |
| `DRMCP-WORK-MCP-006` | Supplies diagnostic, validation, and path-exposure contracts. |
| `DRMCP-WORK-MCP-008` | Supplies accepted legacy, rejected, disabled, and leakage fixtures. |
| `DRMCP-WORK-MCP-009` | Historical retired implementation planning. Its extension seam is not authority. |
| `DRMCP-WORK-MCP-011` | Supplies the accepted replacement runtime architecture. |
| `DRMCP-WORK-MCP-012` | Supplies the completed replacement current runtime consumed by this Work Item. |
| `spec:drmcp.implementation` | Supplies the accepted replacement implementation architecture. |
| DRMCP configuration, parser, index, resolver, retrieval, and validation code | Receive the optional legacy layer. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Legacy code and test inventory | W005, W006, W008, completed W012, and `PRODUCT-WORK-SPEC-014` | Identify the exact rebaselined implementation boundary without treating the W009 seam as authority. |
| B. Configuration, parsing, and archive index | Phase A | Implement `legacy_roots`, root validation, legacy parser, and separate archive index. |
| C. Retrieval and fallback resolution | Phases A-B | Implement exact legacy retrieval and current-first fallback resolution. |
| D. Relation validation and isolation | Phases B-C | Implement current-to-legacy validation, diagnostics, read-only behavior, and leakage protection. |
| E. Verification | Phases B-D | Run accepted, rejected, disabled-fallback, current-regression, and leakage tests. |
| F. Review and closure | Phase E | Run implementation review, apply corrections, record evidence, and close. |

Integrated hub verification begins after this Work Item closes.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Inventory legacy-related code and tests against accepted contracts, fixtures, and the completed W012 runtime. | W005, W006, W008, and W012 accepted. |
| T02 | Implement `legacy_roots`, root validation, legacy parsing, and separate archive-index construction. | T01. |
| T03 | Implement exact legacy retrieval and current-first fallback resolution. | T01-T02. |
| T04 | Implement current-to-legacy relation validation, diagnostics, read-only enforcement, and non-leakage behavior. | T02-T03. |
| T05 | Run accepted, rejected, disabled-fallback, current-regression, and leakage tests. | T02-T04. |
| T06 | Run implementation review, apply required corrections, record evidence, and close. | T05. |

## Completion Condition

This Work Item is complete when all of the following are true:

- legacy fallback activates only for explicitly configured `legacy_roots`;
- missing, invalid, duplicate, and overlapping roots have deterministic behavior;
- approved legacy sequential records build a separate read-only archive index;
- exact retrieval and fallback accept only approved V01 sequential families;
- current resolution runs before any legacy grammar or archive lookup;
- successful legacy resolution preserves the issued legacy ID;
- configured current-to-legacy relations validate correctly;
- disabled, unsupported, unresolved, and duplicate legacy cases produce accepted diagnostics;
- legacy records do not appear in normal listing, current validation, authoring targets, or the active index;
- no `V01-SPEC-*`, bare-ID, path-inference, or fuzzy behavior is implemented;
- no automatic `v01/` discovery occurs;
- completed W012 current-runtime behavior remains unchanged without configured legacy roots;
- accepted legacy and leakage fixtures have automated coverage;
- all relevant automated tests pass;
- implementation review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records changed code, tests, commands, results, review verdict, and residual limitations.

## Evidence

- `DRMCP-ADR-MCP-001`: Accepted configuration-gated archive fallback direction.
- `DRMCP-INV-MCP-002`: Legacy implementation and contract-drift evidence.
- `PRODUCT-WORK-SPEC-014`: Compatibility prerequisite.
- `DRMCP-WORK-MCP-005`, `DRMCP-WORK-MCP-006`, and `DRMCP-WORK-MCP-008`: Accepted contract and fixture inputs.
- `DRMCP-WORK-MCP-012`: Required completed replacement current-runtime input.
- `DRMCP-WORK-MCP-009`: Historical retired planning only; its production structure and extension seam are not authority.
- `DRMCP-TASK-MCP-001-10`: Hub lifecycle gate for this Work Item.
- Implementation, automated-test, and independent-review evidence: pending Task execution.

### Rebuild-line disposition

Disposition: retained scope; implementation graph rebaseline required.

- Optional configured legacy fallback scope remains valid.
- Existing W010 Task Candidates must not execute yet.
- The W009 current-only implementation baseline and extension seam referenced by W010 are retired.
- W010 depends on the completed output of `DRMCP-WORK-MCP-012`.
- The retired W009 extension seam is not authority.
- W010 requires a new execution-graph authoring, independent review, and release sequence after W012 completion.
- Existing Task Candidates are planning only and are not released implementation contracts.
- No W010 production Task may start before rebaseline and release.
- W010 remains `blocked` until rebaseline completes.
- Legacy public behavior and W008 fixture contracts remain unchanged.
- No implementation Task, including T01, is created by T13.

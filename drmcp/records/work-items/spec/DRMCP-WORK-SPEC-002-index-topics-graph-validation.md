# DRMCP-WORK-SPEC-002: Index Topics graph validation

- **id**: DRMCP-WORK-SPEC-002
- **status**: not_started
- **date**: 2026-06-28
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-006
  - DRMCP-WORK-MCP-007
  - DRMCP-WORK-MCP-008
  - DRMCP-WORK-MCP-009
  - spec:product.design_records.spec_format.document_shape
  - spec:product.design_records.spec_format.topics_table
  - spec:product.design_records.spec_format.spec_id_as_ref
  - spec:product.design_records.spec_format.validation_policy
  - spec:product.design_records.spec_format.follow_up_boundary
- **tasks**: []

## Goal

Implement cross-file Topics graph validation for current PRODUCT spec-format records.

Consume accepted per-file detector results and current active-index state.

## Boundary

This Work Item owns:

- extraction of graph edges from locally valid `## Topics` rows;
- the current Topics row shape `title`, `kind`, `ref`, and `summary` as an implementation input;
- canonical child lookup by exact `ref` against current active-index state;
- authoritative-parent derivation from the declaring `Index` or `Overview` spec;
- consistency checks between the declaring parent and the child H1-adjacent `parent` marker;
- duplicate authoritative-parent detection for one child `ref`;
- graph cycle detection;
- graph detector tests;
- runtime integration tests for the accepted `validate_records` execution path.

This Work Item consumes:

- the accepted callable detector result boundary from `DRMCP-WORK-SPEC-001`;
- current parsed sources, canonical refs, conflict state, and active-index lookup state from `DRMCP-WORK-MCP-003`;
- validation selection, diagnostic categories, severity, ordering, deduplication, and representation from `DRMCP-WORK-MCP-006`;
- PRODUCT Topics, identity, parent, and validation semantics from the referenced authorities;
- shared graph fixtures, manifests, and fixture-local checks from `DRMCP-WORK-MCP-008`.

This Work Item does not own:

- H1 count or shape detection;
- H1-adjacent metadata detection;
- accepted spec-kind detection;
- `contract_class` detection;
- kind-specific required-section detection;
- local `## Topics` table presence, header, or row-shape detection;
- current-root discovery, Markdown parsing contracts, canonical identity derivation, or active-index construction;
- public resolver orchestration or fallback behavior;
- validation request forms, diagnostic taxonomy, severity, source locations, or response representation;
- PRODUCT rule text or severity policy;
- shared fixture authoring or fixture-local structural checks;
- general current-read implementation or tests owned by `DRMCP-WORK-MCP-009`;
- PRODUCT owner-pointer synchronization.

The canonical Topics row does not contain `file` or row-level `parent`.
Alias, redirect, and stale-ref lookup are not accepted graph inputs.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Source Requirement for the retained DRMCP graph-validation implementation. |
| `DRMCP-WORK-SPEC-001` | Supplies the accepted local detector result and structural-eligibility boundary. |
| `DRMCP-WORK-MCP-003` | Supplies parsed current sources, canonical refs, conflicts, and active-index lookup state. |
| `DRMCP-WORK-MCP-006` | Supplies validation execution and diagnostic representation contracts. |
| `DRMCP-WORK-MCP-007` | Supplies the accepted `retain` disposition and rebaseline boundary. |
| `DRMCP-WORK-MCP-008` | Supplies shared Topics and graph fixtures. |
| `DRMCP-WORK-MCP-009` | Retains general current-read implementation and test ownership outside this validator. |
| PRODUCT Topics-table authority | Supplies the `title/kind/ref/summary` row and authoritative-parent semantics. |
| PRODUCT document-shape authority | Supplies child H1-adjacent `parent` marker semantics. |
| PRODUCT ID-as-ref authority | Supplies canonical `ref` grammar and exact path-derived identity semantics. |
| PRODUCT validation policy | Supplies rule ownership and migration-phase severity semantics. |
| DRMCP graph validator implementation and tests | Receive edge extraction, lookup, consistency, duplicate-parent, cycle, and integration changes. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Graph input boundary | Accepted W-SPEC-001 detector result; W003, W006, and PRODUCT authorities | Define eligible sources, edge records, exact lookup inputs, and graph outputs. |
| B. Edge extraction and lookup | Phase A | Extract accepted `ref` edges and resolve canonical children through current active-index state. |
| C. Parent and graph algorithms | Phase B | Implement child-marker consistency, duplicate authoritative-parent detection, and cycle detection. |
| D. Graph tests | Phases B-C; W008 fixtures | Add focused algorithm and invalid-graph tests. |
| E. Runtime integration | Phases C-D | Integrate graph diagnostics into the W006-owned validation path. |
| F. Validation and review | Phases D-E | Run scoped automated tests, independent review, correction, and closure. |

Implementation Tasks begin only after the W-SPEC-001 callable detector result contract is accepted.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Confirm PRODUCT graph semantics, W003 lookup inputs, W006 output boundary, and W008 fixture coverage. | Accepted upstream contracts and fixtures. |
| T02 | Define graph-eligible detector results, edge records, and exact canonical lookup behavior. | W-SPEC-001 accepted detector boundary; T01. |
| T03 | Implement `## Topics` edge extraction and canonical child lookup. | T02. |
| T04 | Implement child parent-marker consistency and duplicate authoritative-parent detection. | T03. |
| T05 | Implement cycle detection and deterministic graph results. | T03-T04. |
| T06 | Add graph algorithm tests and runtime integration tests using W008 fixtures. | T04-T05. |
| T07 | Run scoped verification, independent review, correction, and closure. | T06. |

## Completion Condition

This Work Item is complete when all of the following are true:

- graph validation consumes only sources accepted by the W-SPEC-001 structural boundary;
- no W-SPEC-001 local detector is duplicated;
- accepted rows use exactly `title`, `kind`, `ref`, and `summary` as the current implementation input;
- edge extraction uses canonical child `ref`, not a physical `file` column;
- canonical child lookup uses W003 current active-index state without invoking public resolver fallback;
- the declaring `Index` or `Overview` is the authoritative parent for each row;
- each resolved child H1-adjacent `parent` marker is checked against the declaring parent ref;
- duplicate authoritative-parent declarations are detected without choosing one as authoritative;
- graph cycles are detected deterministically;
- alias, redirect, stale-ref, `file`, and row-level `parent` assumptions are absent;
- graph results use W006 diagnostic contracts without inventing categories, severity, or representation;
- shared fixtures are consumed from W008 without duplicating fixture ownership;
- graph algorithm tests and runtime integration tests pass;
- general current-read tests remain in W009;
- PRODUCT semantic text is cited rather than copied as independent DRMCP authority;
- independent review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records changed code, tests, commands, results, review verdict, and residual limitations.

## Evidence

- `DRMCP-TASK-MCP-007-02` accepted disposition `retain` for this Work Item.
- `DRMCP-TASK-MCP-007-03` rebaselined this record on 2026-06-28.
- Historical `file`, row-level `parent`, alias-waiting, temporary-tooling, and generic redesign assumptions were removed.
- `DRMCP-WORK-SPEC-001` remains the per-file detector owner and supplies the graph-eligibility boundary.
- `DRMCP-WORK-MCP-003` remains the parsed-state, canonical-ref, conflict, and active-index authority.
- `DRMCP-WORK-MCP-006` remains the validation-execution and diagnostic-representation authority.
- PRODUCT spec-format records remain the semantic and severity authorities.
- `DRMCP-WORK-MCP-008` remains the shared fixture owner.
- `DRMCP-WORK-MCP-009` remains the general current-read implementation and test owner.
- Implementation, automated-test, command, and independent-review evidence remain pending future child Tasks.

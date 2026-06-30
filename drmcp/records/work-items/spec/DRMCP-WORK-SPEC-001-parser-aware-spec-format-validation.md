# DRMCP-WORK-SPEC-001: Parser-aware spec format validation

- **id**: DRMCP-WORK-SPEC-001
- **status**: not_started
- **date**: 2026-06-28
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-006
  - DRMCP-WORK-MCP-007
  - DRMCP-WORK-MCP-008
  - DRMCP-WORK-MCP-012
  - DRMCP-WORK-SPEC-002
  - spec:product.design_records.spec_format.document_shape
  - spec:product.design_records.spec_format.spec_id_as_ref
  - spec:product.design_records.spec_format.validation_policy
  - spec:product.design_records.spec_format.follow_up_boundary
- **tasks**: []

## Goal

Implement parser-aware per-file detectors for current PRODUCT spec-format rules.

Integrate those detectors with the accepted DRMCP validation path without redefining parsing, identity, diagnostics, or PRODUCT semantics.

## Boundary

This Work Item owns:

- real ATX H1 count detection outside YAML front matter and fenced code blocks;
- accepted H1 shape and spec-kind detection;
- H1-adjacent metadata presence and shape detection;
- `contract_class` required, prohibited, and allowed-value detection;
- kind-specific required and prohibited section detection;
- local `## Topics` table presence and column-shape detection;
- path-derived canonical ID mismatch detection;
- front-matter migration-policy detection;
- a callable per-file detector result boundary for downstream graph validation;
- detector unit tests;
- runtime integration tests for the accepted `validate_records` execution path.

This Work Item consumes:

- parsed current-source state and canonical identity from `DRMCP-WORK-MCP-003`;
- validation selection, diagnostic categories, severity, ordering, deduplication, and representation from `DRMCP-WORK-MCP-006`;
- PRODUCT semantic rules from the referenced spec-format authorities;
- shared fixture files, manifests, and fixture-local checks from `DRMCP-WORK-MCP-008`.

This Work Item does not own:

- current-root discovery, Markdown parsing contracts, parsed-state shape, or canonical identity derivation;
- validation request forms, subject selection, diagnostic taxonomy, severity, source locations, or response representation;
- PRODUCT rule text, migration severity policy, accepted spec kinds, or required-section matrices;
- cross-file Topics lookup, parent consistency, duplicate-parent detection, or cycle detection;
- shared fixture authoring or fixture-local structural checks;
- general current-read implementation or tests owned by `DRMCP-WORK-MCP-012`;
- PRODUCT owner-pointer synchronization.

`DRMCP-WORK-SPEC-002` consumes the accepted detector boundary from this Work Item.
It must not reimplement local per-file checks.

## Impact Scope

| ref or area | impact |
|---|---|
| `DRMCP-REQ-MCP-001` | Source Requirement for the retained DRMCP validation implementation. |
| `DRMCP-WORK-MCP-003` | Supplies parsed current-source state, headings, metadata, canonical identity, and active-index state. |
| `DRMCP-WORK-MCP-006` | Supplies validation execution and diagnostic representation contracts. |
| `DRMCP-WORK-MCP-007` | Supplies the accepted `retain` disposition and rebaseline boundary. |
| `DRMCP-WORK-MCP-008` | Supplies shared current-format and invalid-case fixtures. |
| `DRMCP-WORK-MCP-012` | Retains general current-read implementation and test ownership outside this validator. |
| `DRMCP-WORK-SPEC-002` | Consumes the accepted local detector result boundary before graph validation. |
| PRODUCT document-shape authority | Supplies H1, metadata, contract-class, kind, and section semantics. |
| PRODUCT ID-as-ref authority | Supplies path-derived canonical ID semantics and mismatch policy. |
| PRODUCT validation policy | Supplies rule ownership and migration-phase severity semantics. |
| DRMCP validator implementation and tests | Receive the detector, integration, and automated-test changes. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Detector boundary | W003, W006, and PRODUCT authorities | Define detector inputs, results, and downstream eligibility without redefining upstream contracts. |
| B. Per-file detector implementation | Phase A | Implement H1, metadata, kind, contract-class, section, local Topics shape, ID mismatch, and front-matter detectors. |
| C. Detector tests | Phases A-B; W008 fixtures | Add focused unit tests for each detector and migration-state branch. |
| D. Runtime integration | Phases B-C | Integrate detector results into the W006-owned validation path. |
| E. Validation and review | Phases C-D | Run scoped automated tests, independent review, correction, and closure. |

W-SPEC-002 implementation starts after Phase A accepts the callable detector result boundary.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Confirm PRODUCT rule mapping, W003 parsed inputs, W006 output boundary, and W008 fixture coverage. | Accepted upstream contracts and fixtures. |
| T02 | Define the callable per-file detector result boundary and graph-eligibility handoff. | T01. |
| T03 | Implement H1, metadata, spec-kind, contract-class, section, local Topics shape, ID mismatch, and front-matter detectors. | T02. |
| T04 | Add detector unit tests using W008 fixtures. | T03. |
| T05 | Integrate detectors with the W006 validation path and add runtime integration tests. | T03-T04. |
| T06 | Run scoped verification, independent review, correction, and closure. | T05. |

## Completion Condition

This Work Item is complete when all of the following are true:

- every retained per-file rule has one implementation detector;
- real H1 counting ignores YAML front matter and fenced code blocks through the accepted parser boundary;
- accepted H1 shape and spec-kind checks use PRODUCT authority;
- required H1-adjacent metadata checks use PRODUCT authority;
- `contract_class` checks cover required, prohibited, and invalid values;
- kind-specific required and prohibited section checks are implemented;
- local `## Topics` presence and column shape are detected without cross-file lookup;
- path-derived ID mismatch detection consumes W003 canonical identity;
- front-matter handling follows PRODUCT migration policy without treating front matter as current metadata;
- detector results use W006 diagnostic contracts without inventing categories, severity, or representation;
- W-SPEC-002 can consume one accepted callable detector result boundary;
- shared fixtures are consumed from W008 without duplicating fixture ownership;
- detector unit tests and runtime integration tests pass;
- general current-read tests remain in W012;
- PRODUCT semantic text is cited rather than copied as independent DRMCP authority;
- independent review reports no blocking or major findings;
- `DRMCP-REQ-MCP-001` lists this Work Item in `work_items`;
- final evidence records changed code, tests, commands, results, review verdict, and residual limitations.

## Evidence

- `DRMCP-TASK-MCP-007-02` accepted disposition `retain` for this Work Item.
- `DRMCP-TASK-MCP-007-03` rebaselined this record on 2026-06-28.
- Historical migration-era metadata, completed PRODUCT Work Item dependencies, generic redesign gates, and diagnostic ownership claims were removed.
- `DRMCP-WORK-MCP-003` remains the parsed-state and canonical-identity authority.
- `DRMCP-WORK-MCP-006` remains the validation-execution and diagnostic-representation authority.
- PRODUCT spec-format records remain the semantic and severity authorities.
- `DRMCP-WORK-MCP-008` remains the shared fixture owner.
- `DRMCP-WORK-MCP-012` is the replacement general current-read implementation and test owner.
- Implementation, automated-test, command, and independent-review evidence remain pending future child Tasks.

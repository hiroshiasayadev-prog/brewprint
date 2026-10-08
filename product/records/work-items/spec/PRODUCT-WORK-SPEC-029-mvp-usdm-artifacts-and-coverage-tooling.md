# PRODUCT-WORK-SPEC-029: MVP USDM artifacts and coverage tooling

- **id**: PRODUCT-WORK-SPEC-029
- **status**: done
- **date**: 2026-07-13
- **source_refs**:
  - PRODUCT-REQ-SPEC-015
- **impact_refs**:
  - `spec:product.design_records`
  - `spec:product.design_records.usdm`
  - `spec:product.design_records.usdm.artifact_format`
  - `spec:product.design_records.usdm.coverage_format`
  - `spec:product.design_records.usdm.coverage_tools`
  - `spec:product.design_records.authoring_standards`
  - `spec:product.design_records.authoring_standards.usdm_authoring`
- **tasks**:
  - PRODUCT-TASK-SPEC-029-01
  - PRODUCT-TASK-SPEC-029-02
  - PRODUCT-TASK-SPEC-029-03
  - PRODUCT-TASK-SPEC-029-04
  - PRODUCT-TASK-SPEC-029-05

## Goal

Introduce the MVP USDM artifact specification, authoring standard, standalone coverage tools, and repository-local MCP wrapper.

The Work Item keeps USDM independent from the canonical artifact model and repository layout until a later migration Work Item.

## Boundary

This Work Item owns:

- the MVP USDM artifact specification under `product/records/spec/design-records/usdm/`;
- the MVP USDM authoring standard under `product/records/spec/design-records/authoring-standards/usdm-authoring.md`;
- standalone USDM validation and coverage tooling under `tools/usdm/`;
- an MCP wrapper around standalone USDM tooling under `tools/usdm-mcp/`;
- one minimal smoke path that proves a USDM requirement can be covered by an implementation Specification.

This Work Item does not own:

- full artifact-model integration;
- repository-layout migration;
- migration of existing records into USDM;
- DRMCP integrated USDM read support;
- final DRMCP Domain component input/output decisions;
- section-level coverage unless file-level `usdm_covers` is insufficient for the MVP;
- production packaging beyond standalone repository tooling.

## Impact Scope

| target | impact |
|---|---|
| `spec:product.design_records` | Add a Topics row for the USDM spec area. |
| `product/records/spec/design-records/usdm/` | Add MVP USDM overview, artifact format, coverage format, and standalone tool contract Specifications. |
| `spec:product.design_records.authoring_standards` | Add a Topics row for the USDM authoring standard. |
| `product/records/spec/design-records/authoring-standards/usdm-authoring.md` | Add MVP USDM authoring rules. |
| `tools/usdm/` | Add standalone validation and coverage tools. |
| `tools/usdm-mcp/` | Add MCP wrapper for the standalone USDM tools. |
| implementation Specification metadata | Allow optional `usdm_covers` metadata for MVP coverage checks. |
| `<app>/records/usdm/` | Treat as MVP USDM record placement. |

## Task flow

```text
PRODUCT-TASK-SPEC-029-01 author MVP USDM artifact Specification and parent topic row
  -> PRODUCT-TASK-SPEC-029-02 author USDM authoring standard
  -> PRODUCT-TASK-SPEC-029-03 implement standalone USDM tools
  -> PRODUCT-TASK-SPEC-029-04 smoke MVP USDM coverage path
  -> PRODUCT-TASK-SPEC-029-05 implement USDM MCP wrapper
```

PRODUCT-TASK-SPEC-029-01 through PRODUCT-TASK-SPEC-029-05 completed the owned delivery boundary.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| PRODUCT-TASK-SPEC-029-01 | authoring | Author the MVP USDM artifact Specification under `product/records/spec/design-records/usdm/` and add the parent topic row. | none |
| PRODUCT-TASK-SPEC-029-02 | authoring | Author `usdm-authoring.md` for MVP record shape, ID rules, requirement tables, and coverage metadata. | T01 |
| PRODUCT-TASK-SPEC-029-03 | implementation | Implement standalone `tools/usdm/` commands for `validate_usdm`, `check_usdm_coverage`, and `usdm_covered_by`. | T01, T02 |
| PRODUCT-TASK-SPEC-029-04 | verification | Add or run a minimal smoke path proving validation and coverage checks over one small USDM sample and one covering Specification. | T03 |
| PRODUCT-TASK-SPEC-029-05 | implementation | Implement an MCP wrapper around the standalone USDM tools. | T03, T04 |

## Completion Condition

- MVP USDM overview, artifact-format, coverage-format, and standalone-tool contract specs are specified.
- MVP USDM authoring rules are specified.
- Standalone USDM tool contracts are specified before implementation.
- Standalone USDM tools exist under `tools/usdm/` after PRODUCT-TASK-SPEC-029-03 completes.
- A repository-local MCP wrapper exists under `tools/usdm-mcp/` after PRODUCT-TASK-SPEC-029-05 completes.
- The tools can validate USDM records and detect uncovered USDM requirement IDs.
- The tools can list implementation Specifications covering one USDM requirement ID.
- The MVP explicitly documents that artifact-model, repository-layout, and migration integration are deferred.

## Evidence

- PRODUCT-REQ-SPEC-015 requires MVP USDM artifacts and coverage checks.
- PRODUCT-TASK-SPEC-029-01 materializes the first authoring Task.
- PRODUCT-TASK-SPEC-029-02 materializes the USDM authoring-standard Task.
- PRODUCT-TASK-SPEC-029-03 materializes the standalone USDM tool implementation Task.
- PRODUCT-TASK-SPEC-029-04 records the inline PowerShell/Python smoke verification result `USDM smoke PASS`.
- PRODUCT-TASK-SPEC-029-05 materializes the repository-local USDM MCP wrapper Task and records successful `@usdm-mcp` runtime calls.
- PRODUCT-TASK-SPEC-029-01 through PRODUCT-TASK-SPEC-029-05 are `done` and collectively satisfy every retained Completion Condition.
- The former DRMCP Domain inventory handoff was removed from this Work Item because the referenced DRMCP workflow was retired and the handoff was not part of the delivered USDM MVP boundary.
- This legacy Work Item had no separate integrated-review Task. Closure uses the recorded Specification authoring results, standalone-tool verification, `USDM smoke PASS`, and successful MCP runtime calls as its accepted completion Evidence.
- No artifact-model integration, repository-layout migration, existing-record migration, or DRMCP integrated USDM read support is claimed by this closure.

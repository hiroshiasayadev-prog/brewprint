# DRMCP-TASK-MCP-004-02: Reflect compact active list-records contract

- **id**: DRMCP-TASK-MCP-004-02
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-004
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-004-01
- **outputs**:
  - spec:drmcp.design_records_mcp.tools.list_records

## Goal

Reflect the accepted compact active-index listing baseline into the authoritative `list_records` operation contract.

Remove broad, range-based, mixed-kind, legacy, spec, path-bearing, and full-metadata listing behavior without entering resolver, diagnostic-taxonomy, fixture, or implementation scope.

## Work

- Rewrite `spec:drmcp.design_records_mcp.tools.list_records` around one required active app, supported sequential kind, and domain scope.
- Define exact request fields, request validation, optional status filtering, canonical ordering, and bounded limits.
- Define compact `ref`, `title`, `status`, and `date` results plus `has_more`.
- Keep invalid-but-addressable records listable with nullable missing compact fields.
- Place missing-compact-field warning triggers at the operation response level.
- Preserve parsed invalid compact values without repair.
- Define valid zero-match behavior.
- Exclude duplicate-conflict identities without choosing a filesystem-order winner.
- Hide physical paths and internal index state from normal responses.
- Remove exact ID, range, broad, mixed-kind, spec, legacy, arbitrary-ordering, and unbounded listing claims.
- Synchronize `DRMCP-WORK-MCP-004` with the Task link and execution evidence.

Changed-file scope:

- create this Task;
- update `DRMCP-WORK-MCP-004` task linkage and Evidence;
- rewrite `spec:drmcp.design_records_mcp.tools.list_records`;
- change no tool catalog, MVP scope, resolver, diagnostics, validation, fixture, implementation, or automated-test file.

Ownership exclusions:

- W003 retains current-root discovery, parsing, canonical identity, active-index construction, and duplicate-conflict construction.
- W005 retains current-first resolution and configured legacy fallback orchestration.
- W006 retains diagnostic category names, severity, shared fields, source-location representation, validation semantics, and exceptional path exposure.
- T04 retains tool-catalog, MVP-scope, responsibility-boundary, and broader pointer synchronization.

## Done condition

- The request has exact required and optional fields.
- One configured active app, one supported sequential kind, and one domain are required.
- Supported kinds are decision, investigation, requirement, work item, and task only.
- Optional status filtering uses the selected kind's PRODUCT vocabulary.
- Ordering uses canonical `ref`, defaults to descending, and allows ascending only.
- Limit defaults to 20 and accepts integers from 1 through 100.
- Each result contains only `ref`, `title`, `status`, and `date`.
- Missing `title`, `status`, or `date` appears as `null` and triggers an operation-level warning.
- Parsed invalid compact values remain unchanged.
- The response contains `has_more` and operation-level warnings.
- Valid zero-match requests return an empty result without warnings or errors.
- Duplicate-conflict identities produce no selected result or path-based tie-break.
- Exact ID, range, broad, mixed-kind, spec, legacy, arbitrary-ordering, unbounded, full-metadata, and physical-path listing behavior is removed.
- W003, W005, and W006 ownership boundaries remain intact.
- Scoped strict validation passes.
- Independent review reports no blocking or major findings.
- Required review corrections are applied and recorded.

## Verification

- Compare the rewritten operation against `DRMCP-TASK-MCP-004-01` accepted baseline.
- Compare addressability and duplicate-conflict handling against `spec:drmcp.design_records_mcp.schema.record_model`.
- Compare common field availability against `spec:drmcp.design_records_mcp.schema.fields`.
- Confirm no exact ID, `id_range`, one-sided range, broad omitted scope, mixed-kind, spec, legacy, path tie-break, arbitrary ordering, unbounded result, full metadata, or physical-path response remains.
- Confirm warning trigger placement does not define W006 category, severity, shared-field, or source-location schemas.
- Run the scoped strict spec validator and `git diff --check`.
- Obtain independent review before changing the Task status to `done`.

## Evidence

### Accepted contract input

| concern | accepted input |
|---|---|
| Scope | Active index only; exactly one configured `app_namespace`, one supported sequential `kind`, and one `domain`. |
| Kinds | `decision`, `investigation`, `requirement`, `work_item`, and `task`. |
| Exclusions | `spec` and legacy archive records are not normal-listing targets. |
| Status | Optional exact filter using the selected kind's PRODUCT vocabulary. |
| Ordering | Canonical `ref`; default `desc`; alternative `asc` only. |
| Limit | Default 20; accepted range 1 through 100; no unbounded listing. |
| Result | Fixed compact fields `ref`, `title`, `status`, and `date`. |
| Truncation | Top-level `has_more`. |
| Invalid addressable source | Missing compact values become `null`; parsed invalid values remain unchanged. |
| Warning placement | Missing returned compact fields trigger top-level operation warnings; W006 defines warning representation. |
| Zero match | Empty results, `has_more: false`, empty warnings, and no error. |
| Duplicate conflict | No winner, no result for the conflicted identity, and no path-order tie-break. |
| Path boundary | Normal response contains no physical path, source provenance, or internal index state. |

T01 fixed these inputs. T02 does not reopen them.

### Changed files

- Created `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-004-02-reflect-compact-active-list-records-contract.md`.
- Updated `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md`.
- Rewritten `drmcp/records/spec/design-records-mcp/tools/list-records.md`.
- No pointer-only schema change was required. `schema.fields` and `schema.record_model` already delegate public listing representation to W004.

### Removed stale claims

- Optional or omitted app, kind, and domain scope.
- Exact `id` filtering.
- `id_range` and one-sided range behavior.
- Mixed-kind listing and kind derivation from range endpoints.
- Spec and legacy archive listing.
- `order_by` and arbitrary ordering keys.
- Path-based duplicate tie-breaking.
- Unbounded result behavior.
- Full common and kind-specific metadata response.
- Physical path response.
- V01-specific listing examples and historical source notes.

### Retained claims

- `list_records` narrows candidates before exact retrieval.
- Status filtering remains optional after scope is fixed.
- Canonical identity controls ordering.
- Request limit bounds result count.
- Invalid request shape fails the operation rather than broadening scope.

### W003 inputs consumed

- The active index contains current records only.
- Every addressable record has canonical identity and kind.
- Invalid-but-addressable records preserve only parsed values and do not receive fabricated fields.
- Duplicate canonical identity creates no addressable winner.
- Source paths remain internal provenance and validation inputs.

### W005 and W006 exclusions

- T02 does not invoke or redefine current-first resolution or legacy fallback.
- T02 does not define warning category names, severity, shared diagnostic fields, or source-location fields.
- T02 does not define validation execution or duplicate-conflict diagnostic representation.
- T02 does not define diagnostic, patch, debug, or emergency path exposure.

### Unresolved decisions

None.

The request field names follow existing contracts: `app_namespace` from current-root configuration, `domain` from domain-scoped tool requests, and existing `kind`, `status`, `order`, and `limit` names.
The response retains `records` and adds top-level `has_more` and `warnings`.

### Validation and review state

Scoped strict validator command:

```powershell
python -X utf8 product/src/tools/validate_spec.py `
  drmcp/records/spec/design-records-mcp/tools/list-records.md `
  --strict `
  --no-color
```

Whitespace and patch-integrity command:

```powershell
git diff --check -- `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-004-02-reflect-compact-active-list-records-contract.md `
  drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md `
  drmcp/records/spec/design-records-mcp/tools/list-records.md
```

- Scoped strict validator result: `[strict]  All 1 file(s) OK.`
- `git diff --check` result: PASS; no whitespace errors reported.
- Git emitted an LF-to-CRLF working-copy warning for `tools/list-records.md`; the warning does not indicate a diff-check failure.
- Independent review verdict: `PASS`.
- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- Advisories: none.
- Review confirmed the request scope, supported kinds, ordering, limit, compact response, invalid-but-addressable behavior, duplicate-conflict behavior, and W003/W005/W006/T04 ownership boundaries.
- Closure readiness: `ready`.
- Task closed as `done` on 2026-06-27.

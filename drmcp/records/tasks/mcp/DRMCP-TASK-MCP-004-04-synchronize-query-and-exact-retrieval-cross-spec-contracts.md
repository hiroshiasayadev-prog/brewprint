# DRMCP-TASK-MCP-004-04: Synchronize query and exact-retrieval cross-spec contracts

- **id**: DRMCP-TASK-MCP-004-04
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-004
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-004-03
- **outputs**:
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.tools.overview
  - spec:drmcp.design_records_mcp.mvp_scope
  - spec:drmcp.design_records_mcp.responsibility_boundary

## Goal

Synchronize the accepted T02 and T03 query and exact-retrieval contracts into broader DRMCP catalog, scope, and responsibility specifications.

Keep shared schema contracts as pointers to W004 behavior without duplicating operation response definitions.

## Work

- Synchronize the public tool catalog with compact `list_records` and sole exact-retrieval `get_records`.
- Remove public `get_record`, stale range and `ids` claims, legacy/spec normal listing, resolver-mediated exact retrieval, and normal physical-path projection.
- Synchronize the current-format read baseline without asserting implementation or release status.
- Separate W003 discovery and model ownership, W004 query and retrieval plus normal path-hiding ownership, W005 resolver ownership, and W006 diagnostic plus exceptional path-exposure ownership.
- Recheck `schema.fields`, `schema.record_model`, and `schema.record_source` for sufficient W004 delegation.
- Recheck the T02 and T03 operation specs without reopening their accepted contracts.
- Update the parent Work Item task relation and execution evidence.

Changed-file scope:

- create this Task;
- update `DRMCP-WORK-MCP-004` task linkage and Evidence;
- correct the `DRMCP-WORK-MCP-006` planning boundary when independent review identifies overlap with W004 normal response ownership, without starting W006 Task execution;
- update `spec:drmcp.design_records_mcp.overview` pointer summaries when required by the synchronized child specs;
- update `spec:drmcp.design_records_mcp.tools.overview`;
- update `spec:drmcp.design_records_mcp.mvp_scope`;
- update `spec:drmcp.design_records_mcp.responsibility_boundary`;
- change shared schema files only when an existing pointer is insufficient;
- change no resolver, diagnostic, validation, fixture, implementation, automated-test, or authoring-transaction normative spec or operation file.

Ownership exclusions:

- W003 retains discovery, parsing, canonical identity, active-index construction, normalized record model, invalid-source retention, and duplicate-conflict state.
- W005 retains resolver invocation, current-first resolution, configured legacy fallback, and fallback ordering.
- W006 retains warning and diagnostic schemas, category names, severity, shared fields, source-location representation, validation semantics, and exceptional path exposure.
- T05 retains final Work Item validation, independent review, required corrections, and Work Item closure.

## Done condition

- The tools overview matches the accepted T02 and T03 operation split.
- The public catalog contains no invokable `get_record` entry.
- `get_records` is the sole exact-retrieval operation, including single retrieval through one `refs` element.
- The MVP scope matches compact active listing and exact retrieval without implementation-status claims.
- The responsibility boundary separates W003, W004, W005, and W006 without redefining their contracts.
- Shared field, record-model, and record-source specs delegate public listing and retrieval representation to W004.
- No old `ids`, range listing, broad listing, dual retrieval tool, resolver-mediated exact retrieval, spec or legacy normal listing, or normal physical-path response claim remains in the changed broader specs.
- No unnecessary shared schema change is made.
- No implementation, fixture, validation, or authoring-transaction ownership is absorbed.
- The changed-file scope and no-change decisions are recorded in Evidence.
- Scoped strict validation and `git diff --check` commands are recorded.
- The Task remains `in_progress` until independent review reports closure readiness.

## Verification

- Compare the changed broader specs against `DRMCP-TASK-MCP-004-02` and `DRMCP-TASK-MCP-004-03`.
- Compare responsibility ownership against W003, W004, W005, and W006.
- Confirm the public catalog omits `get_record` and `suggest_next_record` as available tools.
- Confirm no `ids`, `id_range`, range listing, broad listing, spec or legacy normal listing, resolver invocation during exact retrieval, or normal path-bearing response claim remains in the changed broader specs.
- Recheck `schema.fields`, `schema.record_model`, and `schema.record_source` for operation-pointer sufficiency.
- Run the scoped strict spec validator and `git diff --check`.
- Obtain independent review before changing the Task status to `done`.

## Evidence

### Affected-file manifest

| file | disposition | reason |
|---|---|---|
| `DRMCP-WORK-MCP-006` | changed after independent review | Planning-record correction removes overlapping ownership of normal list and exact-retrieval path hiding while preserving W006 exceptional path exposure. No W006 Task or normative spec was started. |
| `overview.md` | changed | Topics summaries for the MVP and tools child specs became stale after synchronization and required pointer-only correction. |
| `tools/overview.md` | changed | Public catalog, operation summary, response ownership, and stale tool entries required synchronization. |
| `mvp-scope.md` | changed | P0 surface, retired tools, stale YAML source claim, and W004/W005/W006 separation required synchronization. |
| `responsibility-boundary.md` | changed | Read-operation and filesystem boundaries required explicit W003-W006 ownership separation. |
| `schema/fields.md` | rechecked; unchanged | Already delegates public list and retrieval response shapes plus headings and body inclusion to W004. |
| `schema/record-model.md` | rechecked; unchanged | Already delegates public omission, null, warning, heading, body, and path representation to W004 and W006. |
| `schema/record-source.md` | rechecked; unchanged | Already limits headings and body to source availability and delegates public inclusion to W004. |
| `tools/list-records.md` | rechecked; unchanged | T02 contract already matches the accepted compact active listing baseline. |
| `tools/get-record.md` | rechecked; unchanged | T03 retirement marker already defines no public request, response, or compatibility alias. |
| `tools/get-records.md` | rechecked; unchanged | T03 contract already defines sole exact retrieval, ordered partial success, request-wide body inclusion, and normal path hiding. |

### Changed broader contracts

- Root `overview.md` now describes the MVP scope as the current-format read baseline and the tools overview as a navigation-first catalog.
- `tools/overview.md` is now navigation-first and delegates request, response, error, warning, and schema details to their owning contracts.
- The public catalog contains `list_records` and `get_records` without an invokable `get_record` entry.
- The overview removes stale path-bearing common response examples, `ids`, range errors, failure-item behavior, and duplicated diagnostic authority.
- `mvp-scope.md` classifies the current-format read contract without claiming implementation or release completion.
- `mvp-scope.md` removes `get_record`, `suggest_next_record`, YAML-front-matter source claims, broad/range listing, and W004 ownership over resolver or diagnostics.
- `responsibility-boundary.md` records W003 discovery/model, W004 query/retrieval plus normal path hiding, W005 resolver, and W006 diagnostics plus exceptional path exposure.
- `DRMCP-WORK-MCP-006` planning language now consumes W004-owned normal path hiding instead of claiming it, and limits its own path authority to source-location diagnostic, patch, debug, and emergency surfaces.
- `responsibility-boundary.md` removes next-record suggestion and normal path projection from the Design Records MCP responsibility summary.

### Removed stale claims

- Public or invokable `get_record`.
- Coexisting single and batch exact-retrieval tools.
- `ids` as the batch request field.
- ID-range and broad listing behavior.
- Spec or legacy archive inclusion in normal listing.
- Resolver invocation during exact retrieval.
- Physical path in normal list or exact-retrieval records.
- Mixed success and failure response examples owned by the overview.
- `suggest_next_record` as a current auxiliary tool.
- YAML front matter as a current spec source.

### Ownership confirmation

- W003 remains authoritative for discovery, parsing, identity, active-index state, normalized fields, addressability, invalid-source retention, and duplicate conflicts.
- W004 remains authoritative for list and exact-retrieval request and result behavior, batch ordering, partial-success triggers, body inclusion, and normal path hiding.
- W005 remains authoritative for resolver invocation, current-first behavior, configured legacy fallback, and fallback order.
- W006 remains authoritative for warning and diagnostic representation, severity, source locations, validation semantics, and exceptional path exposure; it no longer claims normal list or exact-retrieval path hiding.

### Unresolved decisions

None.

### Validation state

Scoped strict validator command:

```powershell
python -X utf8 product/src/tools/validate_spec.py `
  drmcp/records/spec/design-records-mcp/overview.md `
  drmcp/records/spec/design-records-mcp/tools/overview.md `
  drmcp/records/spec/design-records-mcp/mvp-scope.md `
  drmcp/records/spec/design-records-mcp/responsibility-boundary.md `
  --strict `
  --no-color
```

Whitespace and patch-integrity command:

```powershell
git diff --check -- `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-004-04-synchronize-query-and-exact-retrieval-cross-spec-contracts.md `
  drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md `
  drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md `
  drmcp/records/spec/design-records-mcp/overview.md `
  drmcp/records/spec/design-records-mcp/tools/overview.md `
  drmcp/records/spec/design-records-mcp/mvp-scope.md `
  drmcp/records/spec/design-records-mcp/responsibility-boundary.md
```

- Scoped strict validator result: `[strict]  All 4 file(s) OK.` The validated normative spec set has not changed since that run.
- Initial `git diff --check` result before Evidence and review-correction updates: PASS; no whitespace errors reported.
- Git emitted LF-to-CRLF working-copy warnings for `overview.md`, `tools/overview.md`, `mvp-scope.md`, and `responsibility-boundary.md`; these warnings are non-blocking.
- Independent review verdict: `NEEDS REVISION`.
- Review finding `F-MAJ-01`: W004 and W006 both claimed normal path hiding. Corrected by narrowing the W006 planning record to exceptional path exposure and making W004 the sole normal list/exact-retrieval response authority.
- Review finding `F-MAJ-02`: final `git diff --check` evidence did not cover the Evidence-updated state.
- Post-correction seven-file `git diff --check` result: PASS; no whitespace errors reported.
- That run emitted LF-to-CRLF working-copy warnings only for `overview.md`, `tools/overview.md`, `mvp-scope.md`, and `responsibility-boundary.md`; the Task, W004, and W006 planning records emitted no warning.
- The line-ending warnings remain non-blocking because the command reported no whitespace error or failure.
- Final seven-file `git diff --check` result after Evidence synchronization: PASS; no whitespace errors or command failure reported.
- The final run emitted LF-to-CRLF working-copy warnings only for `overview.md`, `tools/overview.md`, `mvp-scope.md`, and `responsibility-boundary.md`; the Task, W004, and W006 planning records emitted no warning.
- Independent re-review verdict: `PASS`.
- Previous findings `F-MAJ-01` and `F-MAJ-02`: `CLOSED`.
- Remaining blocking, major, and minor findings: none.
- Advisories: none affecting closure readiness.
- Validation evidence assessment: scoped strict validator and final seven-file diff check are sufficient; LF-to-CRLF warnings are non-blocking.
- T04 closure readiness: `ready`.
- `DRMCP-TASK-MCP-004-04` closed as `done` on 2026-06-27.
- Post-closure seven-file `git diff --check` result before final Evidence synchronization: PASS; no whitespace errors or command failure reported.
- LF-to-CRLF working-copy warnings were limited to `overview.md`, `tools/overview.md`, `mvp-scope.md`, and `responsibility-boundary.md`; the Task, W004, and W006 planning records emitted no warning.
- This final Evidence synchronization updates the Task and W004 after that run, so one final seven-file `git diff --check` is required.

# DRMCP-TASK-MCP-004-05: Validate and close query and exact-retrieval contract realignment

- **id**: DRMCP-TASK-MCP-004-05
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-004
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-004-04
- **outputs**:
  - spec:drmcp.design_records_mcp.tools.list_records
  - spec:drmcp.design_records_mcp.tools.get_record
  - spec:drmcp.design_records_mcp.tools.get_records
  - spec:drmcp.design_records_mcp.tools.suggest_next_record
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.tools.overview
  - spec:drmcp.design_records_mcp.mvp_scope
  - spec:drmcp.design_records_mcp.responsibility_boundary

## Goal

Validate the complete W004 query and exact-retrieval contract set.

Confirm stale-claim removal, ownership separation, scoped validation, patch integrity, and independent-review readiness before T05 and W004 closure.

## Work

- Recheck the eight normative specs changed by T02 through T05 review correction.
- Recheck `schema.fields`, `schema.record_model`, and `schema.record_source` without changing them unless a W004-owned contradiction remains.
- Audit the public tool catalog, listing scope, exact-retrieval split, partial-success behavior, path hiding, and current spec source claims.
- Audit W003 through W006 ownership and preserve the W006 planning boundary corrected during T04.
- Apply only necessary W004-scope corrections. Do not redesign accepted T02 or T03 behavior.
- Run scoped strict validation for the eight normative specs.
- Confirm the actual changed-file set and run `git diff --check` after final Evidence synchronization.
- Prepare an independent final review request.
- After independent review reports PASS, apply required corrections and close T05 and W004 in a separate closure step.

Changed-file scope before review:

- create this Task;
- update `DRMCP-WORK-MCP-004` task linkage and Evidence;
- change a normative spec only when the final audit identifies a W004-owned contradiction;
- do not start W005, W006 Task execution, fixtures, implementation, tests, or authoring-transaction work.

## Done condition

- The final affected-file manifest matches the actual worktree diff for W004.
- The eight normative changed specs are recorded.
- Rechecked but unchanged shared schema files are recorded.
- No stale public `get_record`, active `suggest_next_record`, `ids`, broad or range listing, mixed result wrapper, resolver-mediated exact retrieval, normal physical-path projection, or YAML-front-matter current-source claim remains.
- W003, W004, W005, and W006 ownership remains non-overlapping.
- The scoped strict validator passes for all eight normative specs.
- A final `git diff --check` passes after Evidence synchronization.
- LF-to-CRLF warnings are recorded separately from command failure and whitespace errors.
- Independent review reports no blocking, major, or minor findings.
- Required review corrections are applied and revalidated.
- Final Evidence records the review verdict, residual limitations, and closure state.
- T05 and W004 are changed to `done` only after all preceding conditions are satisfied.

## Verification

- Compare the final contracts against T01 through T04 accepted evidence.
- Confirm the public tool catalog omits invokable `get_record` and `suggest_next_record` entries.
- Confirm `get_records` is the sole exact retrieval operation and single retrieval uses one `refs` element.
- Confirm `list_records` requires app namespace, supported sequential kind, and domain scope.
- Confirm specs and legacy archive records remain outside normal listing.
- Confirm exact retrieval does not invoke the resolver or repair inputs.
- Confirm responses contain successful records only and no `retrieval_status` or `record: null` wrapper.
- Confirm normal listing and retrieval responses expose no physical path.
- Confirm overview specs summarize rather than redefine operation response or diagnostic taxonomy.
- Run the scoped strict validator and final changed-file `git diff --check`.
- Obtain independent review before closure.

## Evidence

### Final normative changed-spec set

- `drmcp/records/spec/design-records-mcp/tools/list-records.md`
- `drmcp/records/spec/design-records-mcp/tools/get-record.md`
- `drmcp/records/spec/design-records-mcp/tools/get-records.md`
- `drmcp/records/spec/design-records-mcp/tools/suggest-next-record.md`
- `drmcp/records/spec/design-records-mcp/overview.md`
- `drmcp/records/spec/design-records-mcp/tools/overview.md`
- `drmcp/records/spec/design-records-mcp/mvp-scope.md`
- `drmcp/records/spec/design-records-mcp/responsibility-boundary.md`

### Rechecked but unchanged files

- `drmcp/records/spec/design-records-mcp/schema/fields.md`
- `drmcp/records/spec/design-records-mcp/schema/record-model.md`
- `drmcp/records/spec/design-records-mcp/schema/record-source.md`

The shared schema files already delegate public list, exact-retrieval, heading, body, warning, and path representation to the owning operation or Work Item contracts. No W004-owned contradiction required a shared schema change.

### Final audit result

The first independent review identified one major finding, `F-MAJ-01`: `tools/suggest-next-record.md` still defined an active P1 operation despite the accepted removal from the public surface.

The correction rewrote that file as an explicit retirement marker. It now defines no current request, response, compatibility alias, or tool-execution error contract, and removes the former ADR-number and V01 path-suggestion behavior.

Confirmed results after correction:

- The public catalog contains no invokable `get_record` entry.
- `get_records` is the sole exact-retrieval operation. Single retrieval uses one `refs` element.
- `list_records` has required `app_namespace`, supported sequential `kind`, and `domain` scope.
- Normal listing excludes specs and legacy archive records.
- No current `ids`, broad listing, range listing, mixed success/failure item wrapper, `retrieval_status`, or `record: null` contract remains.
- Exact retrieval classifies each exact input once and does not invoke `resolve_reference`.
- Normal list and exact-retrieval projections exclude physical paths, source provenance, resolver traces, and internal index state.
- `suggest_next_record` is absent from the public catalog and its own operation spec is an explicit retirement marker with no compatibility surface.
- Current specs use H1-adjacent metadata and path-derived identity. YAML front matter is not a current source.
- Root and tools overviews remain navigation-first and do not define W006 diagnostic taxonomy.
- MVP scope classifies the contract without asserting implementation or release completion.

### Accepted contract summary

- Listing: compact active-index results for one app namespace, sequential kind, and domain; optional status; canonical ref ordering; bounded limit; fixed nullable compact fields; `has_more`; valid empty results; no path projection.
- Exact retrieval: required 1 through 20 `refs`; request-wide `include_body`; direct active or configured legacy index lookup by accepted exact family; exact-string ordered deduplication; partial success; successful records only; normalized metadata, headings, and conditional verbatim body; no path projection.

### Ownership audit

- W003 owns discovery, parsing, canonical identity, active-index construction, normalized record state, invalid-source retention, and duplicate-conflict state.
- W004 owns listing and exact-retrieval requests, results, ordering, warning triggers, partial success, body inclusion, and normal physical-path hiding.
- W005 owns resolver invocation, current-first resolution, configured legacy fallback, and fallback order.
- W006 owns warning and diagnostic representation, category names, severity, shared fields, source-location representation, validation semantics, and exceptional path exposure.
- W006 remains `not_started` with no child Tasks. T05 does not start W006 execution.

### Affected-file manifest state

Known W004 final-review files:

- this T05 Task;
- `DRMCP-WORK-MCP-004`;
- the T04-corrected `DRMCP-WORK-MCP-006` planning record;
- the eight normative specs listed above.

User execution of `git status --short` confirmed that the current W004 worktree set also contains untracked T01 through T04 Task files. The final affected-file manifest therefore contains:

- `DRMCP-TASK-MCP-004-01` through `DRMCP-TASK-MCP-004-05`;
- `DRMCP-WORK-MCP-004`;
- the T04-corrected `DRMCP-WORK-MCP-006` planning record;
- the eight normative specs listed above.

Other modified or untracked repository files shown by `git status --short` are outside W004 and are excluded from this Task evidence.

### Scoped strict validation

```powershell
python -X utf8 product/src/tools/validate_spec.py `
  drmcp/records/spec/design-records-mcp/tools/list-records.md `
  drmcp/records/spec/design-records-mcp/tools/get-record.md `
  drmcp/records/spec/design-records-mcp/tools/get-records.md `
  drmcp/records/spec/design-records-mcp/tools/suggest-next-record.md `
  drmcp/records/spec/design-records-mcp/overview.md `
  drmcp/records/spec/design-records-mcp/tools/overview.md `
  drmcp/records/spec/design-records-mcp/mvp-scope.md `
  drmcp/records/spec/design-records-mcp/responsibility-boundary.md `
  --strict `
  --no-color
```

Previous seven-file result: `[strict]  All 7 file(s) OK.`

Final post-correction result: `[strict]  All 8 file(s) OK.`

### Final `git diff --check`

```powershell
git diff --check -- `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-004-05-validate-and-close-query-and-exact-retrieval-contract-realignment.md `
  drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md `
  drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md `
  drmcp/records/spec/design-records-mcp/tools/list-records.md `
  drmcp/records/spec/design-records-mcp/tools/get-record.md `
  drmcp/records/spec/design-records-mcp/tools/get-records.md `
  drmcp/records/spec/design-records-mcp/tools/suggest-next-record.md `
  drmcp/records/spec/design-records-mcp/overview.md `
  drmcp/records/spec/design-records-mcp/tools/overview.md `
  drmcp/records/spec/design-records-mcp/mvp-scope.md `
  drmcp/records/spec/design-records-mcp/responsibility-boundary.md
```

Final post-correction tracked normative-spec result:

- command failure: none reported;
- whitespace errors: none reported;
- LF-to-CRLF warnings:
  - `mvp-scope.md`;
  - `overview.md`;
  - `responsibility-boundary.md`;
  - `tools/get-record.md`;
  - `tools/get-records.md`;
  - `tools/list-records.md`;
  - `tools/overview.md`;
  - `tools/suggest-next-record.md`;
- warning classification: non-blocking because the command completed without whitespace errors.

Refreshed post-correction no-index checks covered:

- `DRMCP-TASK-MCP-004-01` through `DRMCP-TASK-MCP-004-05`;
- `DRMCP-WORK-MCP-004`;
- `DRMCP-WORK-MCP-006`.

Each no-index check returned exit code `1`, which indicates a detected difference from the temporary empty file. No check returned exit code `2` or higher. No whitespace error was reported.

Git emitted an LF-to-CRLF warning for each of the seven untracked records. These warnings are non-blocking because every command completed with the expected no-index difference code and no whitespace error.

A post-correction no-index recheck covered T05 and W004 before this final Evidence synchronization.

- T05 result: exit code `1`; no whitespace error; LF-to-CRLF warning only.
- W004 result: exit code `1`; no whitespace error; LF-to-CRLF warning only.
- No command returned exit code `2` or higher.
- The warnings are non-blocking.

This final Evidence synchronization changes T05 and W004 after that run. One final post-synchronization no-index check is required. Its output must be supplied directly to the independent re-review without another file edit.

### Unresolved decisions

None.

### Independent review state

First independent review verdict: `NEEDS REVISION`.

- Blocking findings: none.
- Major finding: `F-MAJ-01`, active `suggest_next_record` operation contract remained in `tools/suggest-next-record.md`.
- Minor findings: none.
- Advisories: none.

Independent re-review verdict: `PASS`.

- `F-MAJ-01`: `CLOSED`.
- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- Advisories: none.
- Validation evidence assessment: valid.
- Changed-file integrity evidence assessment: valid.
- Ownership assessment: valid and non-overlapping.
- T05 closure readiness: `ready`.
- W004 closure readiness: `ready`.

The re-review confirmed that `tools/suggest-next-record.md` is an explicit retirement marker with no public, invokable, request, response, error, or compatibility surface. It also confirmed alignment with the public catalog, MVP scope, ADR decision, T05 outputs, W004 scope, and the final eight-spec manifest.

### Closure state

- T05 status changed to `done` on 2026-06-27.
- W004 closure synchronization is performed in the parent Work Item.
- The final normative set contains eight specs.
- Final scoped validation result: `[strict]  All 8 file(s) OK.`
- No unresolved decision remains.
- No residual W004-scope limitation remains.
- W005 and W006 execution remain outside this closure.
- A final post-closure no-index check of T05 and W004 must be executed after this synchronization and supplied directly as external closure evidence without another file edit.

### Residual limitations

None within W004 contract scope. The only remaining action is the external post-closure integrity check for the two closure-updated records.

# DRMCP-TASK-MCP-001-04: Track query and exact-retrieval contract correction

- **id**: DRMCP-TASK-MCP-001-04
- **status**: done
- **date**: 2026-06-26
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d coordination
- **depends_on**:
  - DRMCP-TASK-MCP-001-03
- **outputs**:
  - DRMCP-WORK-MCP-004

## Goal

Accept the corrected query and exact-retrieval contract gate.

## Work

- Track `DRMCP-WORK-MCP-004` as the exact child Work Item selected by T01.
- Delegate scoped listing, ordering, range removal, exact retrieval, batch behavior, and unsupported-input contracts to that child Work Item.
- Confirm that the child Work Item keeps exact retrieval separate from reference resolution.
- Track the child Work Item through contract review and `done`.
- Record the child Work Item ID and accepted evidence here.

This Task does not modify query or retrieval operation specs.
All detailed contract work belongs to the selected child Work Item.

## Done condition

- The selected child Work Item is `done`.
- Normal listing queries only the active index.
- Query filters, defaults, ordering, range behavior, and partial-result behavior are explicit.
- Exact retrieval accepts only supported exact current identities and configured exact legacy IDs.
- Batch retrieval does not infer app prefixes or perform fuzzy normalization.
- The child review has no blocking or major findings.
- The exact child Work Item ID and evidence pointer are recorded here.

## Verification

- Review the final query and retrieval operation contracts.
- Confirm stable ordering and duplicate-request behavior.
- Confirm that this Task contains no direct contract implementation evidence.

## Evidence

- T01 selected `DRMCP-WORK-MCP-004` as the query and exact-retrieval contract owner.
- The child Work Item records the accepted compact `list_records` and `get_records`-only baseline.
- `DRMCP-WORK-MCP-004` and `DRMCP-TASK-MCP-004-01` moved to `in_progress` on 2026-06-27.
- T01 records the authority baseline, affected-file manifest, and W005/W006 exclusions without normative spec edits.
- `DRMCP-TASK-MCP-004-01` independent review verdict: `PASS`; findings and advisories: none; closure readiness: `ready`.
- `DRMCP-TASK-MCP-004-01` closed as `done` on 2026-06-27.
- `DRMCP-WORK-MCP-004` completed its T02 through T05 contract correction flow and changed to `done` on 2026-06-27.
- Accepted listing contract: active-index-only compact listing scoped by `app_namespace`, supported sequential `kind`, and `domain`, with bounded ordering, nullable compact fields, `has_more`, and no physical-path projection.
- Accepted exact-retrieval contract: `get_records` is the sole exact-retrieval operation; one-record retrieval uses one `refs` element; exact lookup performs no prefix inference, fuzzy repair, path interpretation, or resolver invocation.
- `get_record` and `suggest_next_record` are explicit retirement markers with no public, invokable, or compatibility surface.
- Batch results preserve first-occurrence request order, use exact-string ordered deduplication, retain retrievable records under partial failure, and return successful records only.
- Final normative changed-spec set contains eight specs, including `spec:drmcp.design_records_mcp.tools.suggest_next_record`.
- Final scoped validator result: `[strict]  All 8 file(s) OK.`
- Final independent re-review verdict: `PASS`; previous finding `F-MAJ-01`: `CLOSED`; blocking, major, minor findings, and advisories: none.
- Review accepted validation evidence, changed-file integrity evidence, and W003 through W006 ownership separation.
- Final post-closure integrity check for T05 and W004 returned expected no-index exit code `1` for both files, with no whitespace error and LF-to-CRLF warnings only.
- The exact child evidence pointer is `DRMCP-WORK-MCP-004` and its child Tasks `DRMCP-TASK-MCP-004-01` through `DRMCP-TASK-MCP-004-05`.
- This lifecycle-tracking Task contains no direct query or retrieval contract implementation.

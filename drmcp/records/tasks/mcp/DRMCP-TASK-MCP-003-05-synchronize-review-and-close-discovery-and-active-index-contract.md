# DRMCP-TASK-MCP-003-05: Synchronize, review, and close discovery and active-index contract

- **id**: DRMCP-TASK-MCP-003-05
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-003
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-003-04
- **outputs**:
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.responsibility_boundary
  - spec:drmcp.design_records_mcp.schema.overview
  - spec:drmcp.design_records_mcp.schema.record_source

## Goal

Synchronize the final W003 contract across overview and ownership summaries.

Validate and independently review the complete discovery and active-index contract before closing the Work Item and its hub gate.

## Work

- Compare T01 through T04 outputs with the overview, responsibility, namespace-scanning, and schema navigation specifications.
- Remove stale automatic root derivation, V01-only current identity, YAML `design_record`, and physical-path response claims from the T05-owned summaries.
- Preserve explicit configured current roots, separate current and legacy indexes, current path-derived spec identity, and shared invalid-source behavior.
- Keep normal list and exact-retrieval response representation delegated to `DRMCP-WORK-MCP-004`.
- Keep diagnostic identifiers, severity, and source-location response representation delegated to `DRMCP-WORK-MCP-006`.
- Run scoped strict validation after normative synchronization.
- Obtain an independent review of W003, T01 through T05, the hub task, changed DRMCP specifications, and PRODUCT authorities.
- Apply required corrections for blocking or major findings.
- Synchronize final evidence and lifecycle state in this Task, `DRMCP-WORK-MCP-003`, and `DRMCP-TASK-MCP-001-03`.

This Task does not define implementation, fixtures, tests, query behavior, exact retrieval behavior, resolver behavior, or diagnostic taxonomy.

## Done condition

- T01 through T04 outcomes are synchronized into the overview and ownership summaries.
- No active current-source claim uses YAML `design_record` metadata.
- No active current-root claim derives `namespace_prefix` or app identity from the app directory.
- No V01-only ID model appears as the current identity model.
- `DRMCP-WORK-MCP-004` retains normal list and exact-retrieval response ownership.
- `DRMCP-WORK-MCP-006` retains diagnostic taxonomy, severity, and source-location response ownership.
- The W003 changed contract set passes strict validation.
- Independent review reports no blocking or major findings.
- W003 Evidence records the final changed files, validation result, review verdict, and delegated residual scope.
- `DRMCP-WORK-MCP-003` is `done`.
- `DRMCP-TASK-MCP-001-03` accepts the child completion evidence and is `done`.

## Verification

- Search the W003-owned specification set for stale `design_record`, YAML current-source, automatic namespace derivation, V01 current-ID, and normal physical-path response claims.
- Compare the synchronized contract with `DRMCP-TASK-MCP-003-02`, `DRMCP-TASK-MCP-003-03`, and `DRMCP-TASK-MCP-003-04`.
- Compare identity and source semantics with the referenced PRODUCT authorities.
- Run `validate_spec.py` against `drmcp/records/spec` with strict mode.
- Run an independent final review without changing files during review.

## Evidence

- Upstream Tasks `DRMCP-TASK-MCP-003-01` through `DRMCP-TASK-MCP-003-04` are `done`.
- T04 final scoped validation passed for five changed specifications: `[strict] All 5 file(s) OK.`
- T04 final independent review verdict: `PASS`; no blocking, major, or minor findings remained.
- Initial T05 synchronization found stale current-contract summaries in:
  - `spec:drmcp.design_records_mcp.overview`;
  - `spec:drmcp.design_records_mcp.responsibility_boundary`;
  - `spec:drmcp.design_records_mcp.schema.overview`.
- `spec:drmcp.design_records_mcp.namespace_scanning` was rechecked against T02 and required no T05 change.
- Normative synchronization applied on 2026-06-27:
  - overview now uses explicit configured current roots and current canonical identity;
  - responsibility boundary now uses H1-adjacent current metadata and path-derived spec identity;
  - schema overview now describes complete current IDs, path-derived spec refs, source provenance, and active-index outcomes;
  - W004 and W006 ownership boundaries remain explicit.
- Strict validation passed on 2026-06-27: `[strict] All 30 file(s) OK.`
- Initial W003 independent review verdict: `NEEDS REVISION`.
  - Blocking findings: none.
  - Major finding: `schema.record_source` defined concrete `get_record` / `get_records`, `include_body`, headings response-field, and found-record response behavior owned by W004.
  - Minor findings: none.
  - Advisories: none.
- Major finding correction applied on 2026-06-27:
  - `schema.record_source` now defines Markdown headings and body only as readable source material;
  - concrete tool names, request flags, and public response fields were removed;
  - public heading inclusion, optional body inclusion, tool selection, and response representation are explicitly delegated to `DRMCP-WORK-MCP-004`.
- Post-correction validation completed successfully for all 30 DRMCP spec files on 2026-06-27.
- Independent re-review verdict: `PASS`.
  - Previous Finding 1: `CLOSED`.
  - Blocking findings: none.
  - Major findings: none.
  - Minor findings: none.
  - Advisories: none.
- `DRMCP-WORK-MCP-003` closure synchronized to `done`.
- Hub task `DRMCP-TASK-MCP-001-03` accepted the child completion evidence and closed as `done`.

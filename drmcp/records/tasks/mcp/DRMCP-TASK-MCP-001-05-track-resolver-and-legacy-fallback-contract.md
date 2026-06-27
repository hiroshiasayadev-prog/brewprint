# DRMCP-TASK-MCP-001-05: Track resolver and legacy-fallback contract correction

- **id**: DRMCP-TASK-MCP-001-05
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d coordination
- **depends_on**:
  - DRMCP-TASK-MCP-001-03
  - DRMCP-TASK-MCP-001-04
- **outputs**:
  - DRMCP-WORK-MCP-005

## Goal

Accept the corrected current-first resolver and configured legacy-fallback contract gate.

## Work

- Track `DRMCP-WORK-MCP-005` as the exact child Work Item selected by T01.
- Delegate current-first resolution order, accepted legacy grammar, legacy-root configuration, and fallback lookup contracts to that child Work Item.
- Confirm that the child Work Item excludes fuzzy repair, path inference, and legacy spec aliases.
- Track the child Work Item through contract review and `done`.
- Record the child Work Item ID and accepted evidence here.

This Task does not modify resolver or configuration specs.
All detailed contract work belongs to the selected child Work Item.

## Done condition

- The selected child Work Item is `done`.
- Current canonical resolution runs before legacy grammar checks.
- Legacy fallback is disabled without configured `legacy_roots`.
- Only approved exact V01 sequential families are accepted.
- `V01-SPEC-*`, app-prefixless IDs, paths, and fuzzy normalization are rejected.
- Legacy resolution preserves the issued legacy ID.
- The child review has no blocking or major findings.
- The exact child Work Item ID and evidence pointer are recorded here.

## Verification

- Review the final resolver and legacy-root configuration contracts.
- Confirm that current and legacy indexes remain separate.
- Confirm that this Task contains no direct contract implementation evidence.

## Evidence

- Selected child Work Item: `DRMCP-WORK-MCP-005`.
- `DRMCP-WORK-MCP-003` is `done` and supplies configured current roots, current identity, and separate current/legacy index scopes.
- `DRMCP-WORK-MCP-004` is `done` and supplies `get_records` as sole exact retrieval without resolver invocation.
- `PRODUCT-WORK-SPEC-014` is `done` and removed `V01-SPEC-*` compatibility authority.
- Accepted legacy fallback families remain `V01-ADR-*`, `V01-INV-*`, `V01-REQ-*`, `V01-WORK-*`, and `V01-TASK-*`.
- `DRMCP-WORK-MCP-005` moved to `in_progress` on 2026-06-27.
- `DRMCP-TASK-MCP-005-01` completed on 2026-06-27 after independent final re-review returned `PASS` with all findings closed.
- T02 start readiness and T03 design-decision readiness were both `READY`.
- `DRMCP-TASK-MCP-005-02` through `DRMCP-TASK-MCP-005-04` completed with all reported findings closed.
- `DRMCP-TASK-MCP-005-05` completed on 2026-06-28 after final validation, correction, and limited independent re-review.
- Final normative manifest: `namespace-scanning.md`, `resolver.md`, `tools/resolve-reference.md`, and `tools/get-records.md`.
- Post-correction scoped validator result: `[strict]  All 4 file(s) OK.`
- `git diff --check` reported no whitespace error; LF-to-CRLF warnings were non-blocking.
- Initial final review finding `F-MIN-FINAL-01` was corrected and is `CLOSED`.
- Limited independent re-review verdict: `PASS` with no blocking, major, or minor findings.
- Advisories A-01 and A-02 remain non-blocking.
- `DRMCP-WORK-MCP-005` changed to `done` on 2026-06-28.
- The selected child Work Item satisfies the current-first resolver, configured legacy fallback, accepted-family, rejection, issued-ID preservation, and separate-index gates.
- Detailed resolver, configuration, and fallback contract work remained delegated to `DRMCP-WORK-MCP-005`.
- This lifecycle-tracking Task contains no direct resolver or configuration contract implementation.
- `DRMCP-TASK-MCP-001-05` changed to `done` on 2026-06-28 after accepting the child Work Item closure evidence.

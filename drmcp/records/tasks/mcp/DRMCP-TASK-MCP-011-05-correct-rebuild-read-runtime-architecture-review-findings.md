# DRMCP-TASK-MCP-011-05: Correct rebuild read-runtime architecture review findings

- **id**: DRMCP-TASK-MCP-011-05
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-011
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-011-04
- **outputs**:
  - DRMCP-TASK-MCP-011-05
  - DRMCP-WORK-MCP-011
  - DRMCP-TASK-MCP-011-01
  - DRMCP-ADR-MCP-002
  - DRMCP-ADR-MCP-003
  - DRMCP-ADR-MCP-005
  - DRMCP-ADR-MCP-006
  - spec:drmcp.implementation
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.responsibility_boundary
  - spec:drmcp.design_records_mcp.tools.overview
  - spec:drmcp.design_records_mcp.tools.authoring_transaction_model
  - spec:drmcp.design_records_mcp.tools.accept_proposed_write

## Goal

Correct F-MAJ-01 and F-MAJ-02 without changing the accepted architecture decisions.

Make the W011 operation scope and authoring-integration authority unambiguous. Leave finding closure to an independent T06 re-review.

## Work

- Limit W011 architecture to `list_records`, `get_records`, `resolve_reference`, and `validate_records`.
- Clarify that the package tree is complete only for the W011 read-runtime slice.
- Keep authoring-guidance and authoring-transaction use cases, snapshots, and package placement outside W011.
- Preserve D-006's persisted-state fresh-rebuild rule as a future integration constraint.
- Keep standalone `validate_records` fresh-snapshot validation as the current W011 normative runtime.
- Defer current authoring-transaction integration to `DRMCP-REQ-MCP-002`.
- Remove or defer the Phase C normative connection from authoring transaction Specifications.
- Preserve public request, response, status, warning, diagnostic, and error behavior.
- Synchronize this correction into W011.
- Stop before independent re-review, finding closure, design closure, implementation planning, or production implementation.

## Done condition

- W011 architecture is limited to the four read-runtime operations.
- The package tree is identified as the complete inventory for the W011 read-runtime scope, not the complete server package inventory.
- Authoring-guidance and authoring-transaction architecture are explicitly outside W011.
- D-006's persisted-state fresh-rebuild rule remains accepted.
- Current authoring-transaction integration is deferred to `DRMCP-REQ-MCP-002`.
- Phase C's normative authoring-transaction connection is removed or explicitly deferred.
- Public read-operation behavior is unchanged.
- This Task does not declare either finding closed.
- T06 is not created.
- Scoped Git diff and whitespace inspection pass.

## Verification

- Check T01, ADR-002, ADR-003, ADR-006, and the affected Specifications for one exact four-operation scope.
- Check ADR-005 and validation sections for standalone `validate_records` authority and future persisted-write freshness.
- Check authoring transaction Specifications for deferred current-format integration.
- Confirm T04 is unchanged.
- Confirm W011 remains `in_progress`.
- Confirm no T06 record exists.
- Inspect only the 14-file writable boundary for textual completeness, whitespace, and staged changes.
- Run a Specification validator only when an exact command is present in the read startup authority.

## Evidence

### F-MAJ-01 correction

- W011 architecture now applies only to `list_records`, `get_records`, `resolve_reference`, and `validate_records`.
- T01 D-001, D-002, D-007, and D-009 use the same four-operation boundary.
- ADR-002 and ADR-003 limit snapshot lifecycle and dedicated use cases to those operations.
- ADR-006 and `spec:drmcp.implementation` define a complete package inventory only for the W011 read-runtime slice.
- The package tree is not a complete inventory for the Design Records MCP server.
- Authoring-guidance and authoring-transaction use cases, snapshots, and package placement remain outside W011.
- Public authoring-guidance and authoring-transaction tools remain in the catalog under their owning contracts.

### F-MAJ-02 correction

- Standalone `validate_records` remains the current W011 fresh-snapshot validation runtime.
- Validators continue to perform no filesystem I/O.
- Per-source validation continues to precede relation and Topics graph validation.
- Finding aggregation and MCP formatting remain outside individual validators.
- D-006 continues to require any future persisted-write caller to discard candidate and pre-write state.
- A future persisted-write caller must rebuild validation input from persisted files.
- Current authoring-transaction integration is deferred to `DRMCP-REQ-MCP-002`.
- Phase C's normative connection from current authoring contracts to W011 snapshot architecture was removed.
- YAML and V01-SPEC authoring semantics are not treated as integrated with current-format validation.

### Changed files

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-011-01-run-rebuild-read-runtime-architecture-decision-loop.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-011-05-correct-rebuild-read-runtime-architecture-review-findings.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-011-rebuild-read-runtime-application-architecture.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-002-request-scoped-read-runtime-snapshot-and-lifecycle.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-003-layered-read-runtime-application-and-adapter-boundaries.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-005-validation-orchestration-over-fresh-snapshots.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-006-read-runtime-go-package-boundaries.md`
- `drmcp/records/spec/implementation/index.md`
- `drmcp/records/spec/design-records-mcp/overview.md`
- `drmcp/records/spec/design-records-mcp/namespace-scanning.md`
- `drmcp/records/spec/design-records-mcp/responsibility-boundary.md`
- `drmcp/records/spec/design-records-mcp/tools/overview.md`
- `drmcp/records/spec/design-records-mcp/tools/authoring-transaction-model.md`
- `drmcp/records/spec/design-records-mcp/tools/accept-proposed-write.md`

### Preserved decisions and behavior

- D-003, D-004, D-005, and D-008 retain their accepted meaning.
- D-006's persisted-state freshness rule is preserved.
- ADR status, date, `depends_on`, `supersedes`, and `migrated_to_spec` metadata remain unchanged.
- Public read request, response, status, warning, diagnostic, and error behavior remains unchanged.
- Proposal lifecycle, write eligibility, pre-write validation, `written`, `files_written`, repair guidance, actual-write reporting, and no-automatic-rollback semantics remain unchanged.
- T04 was read but not modified.
- W011 remains `in_progress`.

### Deferred scope

- Current-format authoring-transaction source loading and post-write validation integration belong to `DRMCP-REQ-MCP-002`.
- Authoring-guidance application architecture remains with its owning contracts.
- T06 was not created or added to W011.
- Production implementation and implementation Task authoring did not begin.

### Verification result

- The scoped semantic search found no remaining all-tool forms of `each MCP invocation`, `each MCP tool`, `each public MCP tool`, or `each public read tool` in the correction boundary.
- The exact four-operation set appears across the corrected decision, ADR, Work Item, and Specification authority.
- Scoped `git.inspect_diff` returned `pass` and returned the complete 14-file patch.
- Scoped `git.inspect_worktree` returned `pass`.
- Whitespace status was `pass` with no findings.
- No scoped file was staged.
- LF-to-CRLF conversion messages were advisory working-copy warnings.
- Repository-wide cleanliness was not checked or inferred.
- Specification validator: NOT RUN. No exact validator command was present in the read startup authority or authoring standards.
- The task-directory inventory contains T01 through T05 and no T06 file.
- This lifecycle and Evidence update changes T05 bytes. Final post-update scoped Git inspection is reported by the correction execution output.

### Finding disposition

Correction complete; closure pending independent re-review.

The correction author does not independently close F-MAJ-01 or F-MAJ-02.
The next gate is T06 independent finding-closure re-review.

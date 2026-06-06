# TASK-MCP-021-01: review current named_section_replace body flow

- **id**: TASK-MCP-021-01
- **status**: done
- **date**: 2026-06-05
- **work_item**: WORK-MCP-021
- **source_requirement**: REQ-MCP-022
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Current-flow notes for propose_record_update named_section_replace body and body_cache_id handling
  - Implementation touchpoint list for heading-safe normalization

## Goal

Clarify the current authoring update flow for `propose_record_update` with `update.type = named_section_replace`, including direct `body` input, `body_cache_id` resolution, retained proposal creation, diff generation, and diagnostics.

## Work

- Inspect the current code paths for `propose_record_update` update handling.
- Identify where replacement body content is resolved and validated.
- Identify where retained proposal diffs are generated.
- Identify the safest insertion point for first-line heading normalization.
- Check interactions with REQ-MCP-021 heading canonicalization work and avoid making this task depend on tolerant selector matching.

## Done condition

- The update flow is documented enough for implementation work to proceed.
- Candidate implementation touchpoints are identified.
- Any dependency or collision with REQ-MCP-021 / WORK-MCP-019 is explicitly noted.

## Verification

- Review relevant implementation files and tests.
- Confirm that both `body` and `body_cache_id` sources are covered by the investigation notes.

## Evidence
Verdict: PASS — current-flow review completed.

## Findings

- MCP `propose_record_update` dispatch enters `designrecordsmcp/tools_call.go`, then calls `designrecords.ProposeRecordUpdate`.
- `ProposeRecordUpdate` validates update kind, resolves the body source through `resolveBodySource`, and passes the resolved body into `prepareUpdate`.
- Direct `body` and `body_cache_id` are already unified before `prepareUpdate`; therefore normalization inside the named-section replacement path can cover both sources.
- `prepareUpdate` loads the target record file and calls `replaceNamedSection` for `named_section_replace`.
- `replaceNamedSection` currently trims trailing newlines, prepends the selected heading, and appends replacement body content. If the body itself starts with the same selected heading, this creates duplicated headings.
- The best insertion point is the replacement-body handling inside `replaceNamedSection`, before constructing the final replacement lines.
- New diagnostic category needed: `section_replacement_body_heading_stripped`.
- Important implementation risk: `prepareUpdate` currently treats any diagnostics returned from `replaceNamedSection` as proposal-blocking. REQ-MCP-022 needs warning diagnostics to be retained and surfaced without blocking proposal creation. Error-severity diagnostics should still block proposal creation.
- Existing tests to extend include named-section selector tests, body source/cache tests, and spacing preservation tests. A direct `replaceNamedSection` unit test is also recommended.
- REQ-MCP-022 can be implemented independently from REQ-MCP-021 / WORK-MCP-019 because it operates after section selection has already succeeded.

## Files inspected by reviewer

- `internal/designrecordsmcp/tools_call.go`
- `internal/designrecords/authoring.go`
- `internal/designrecords/authoring_test.go`
- `internal/designrecords/types.go`
- `internal/designrecordsmcp/tools_call_test.go`

## Commands / inspection

Claude Code performed focused repository inspection using Glob/Grep/Read. No repository files were changed during the review.

## Follow-up

Proceed to `TASK-MCP-021-02` to update the MCP authoring contract and diagnostic documentation before implementation.

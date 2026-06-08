# V01-REQ-MCP-027: propose_record_update git-style accurate diff and no-op detection

- **id**: V01-REQ-MCP-027
- **status**: accepted
- **date**: 2026-06-06
- **source_refs**:
  - V01-REQ-MCP-020
- **work_items**:
  - V01-WORK-MCP-023

## Requirement

`propose_record_update` MUST return reviewable and accurate diff output for retained update proposals.

For modify proposals, the returned `diff.text` MUST use a git-style unified diff format suitable for proposal review.

The diff MUST compare the current persisted file content against the proposed content. It MUST NOT render the entire target record as newly added content unless the target file is actually new.

`propose_record_update` MUST also detect no-op updates. If the proposed content is byte-equivalent to the current persisted content after the operation is applied, the operation MUST surface a no-op result and MUST NOT create a retained write proposal.

## Evidence

A runtime `propose_record_update` call using `metadata_fields_replace` for `V01-TASK-MCP-020-04` returned a modify proposal whose `diff.text` contained the full target record as `+` lines, even though the target file already existed and the requested metadata value was already `done`.

Expected behavior for an unchanged `status: done` update is no retained proposal, or an explicit no-op response. Expected behavior for a real status change is a small modify diff around the changed metadata line, not a whole-file addition.

This makes proposal review unreliable because reviewers cannot distinguish a one-line metadata update from an apparent full-file creation.

## Required Outcome

`propose_record_update` returns accurate review diffs for update proposals.

Acceptance criteria:

- For existing-file update proposals, `diff.text` compares the current persisted file content with the proposed content.
- Modify diffs are emitted in git-style unified diff format, including `diff --git a/<path> b/<path>`, `--- a/<path>`, `+++ b/<path>`, and `@@` hunk headers.
- When available, `diff.text` SHOULD include an `index <oldhash>..<newhash> 100644` line or an equivalent stable old/new content hash representation.
- Metadata-only changes show only changed lines plus bounded context, not the full record as newly added content.
- Whole-file `+` output is only valid for actual create/add proposals.
- No-op updates are detected after applying operation semantics.
- No-op updates do not create retained write proposals.
- No-op responses include clear diagnostics or response fields that identify the request as no-op.
- Existing proposal accept guards remain unchanged.

## Explicitly Excluded Scope

- Changing the semantics of `metadata_block_replace`, `metadata_fields_replace`, or `named_section_replace`.
- Automatically accepting update proposals.
- Replacing repository validation or accept-time guards.
- Requiring exact byte preservation for unrelated renderer/normalizer behavior outside the authoring update path.

## Boundary

This requirement belongs to the Design Records MCP authoring update contract. It concerns proposal review output and no-op handling for update proposals, not DATA-domain record semantics.

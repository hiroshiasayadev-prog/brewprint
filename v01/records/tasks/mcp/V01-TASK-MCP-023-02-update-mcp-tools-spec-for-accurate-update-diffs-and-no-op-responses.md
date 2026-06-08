# V01-TASK-MCP-023-02: Update MCP tools spec for accurate update diffs and no-op responses

- **id**: V01-TASK-MCP-023-02
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-023
- **source_requirement**: V01-REQ-MCP-027
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-023-01
- **outputs**:
  - SPEC-design-records-mcp-tools update
  - diff and no-op response contract

## Goal

Update the Design Records MCP tools spec so `propose_record_update` has an explicit contract for accurate existing-file update diffs and no-op update responses.

## Work

- Document that existing-file update proposals compare persisted current content with proposed content.
- Document git-style unified diff expectations for modify proposals.
- Clarify that whole-file addition output is valid only for actual create/add proposals.
- Define the response shape or diagnostics for no-op update requests that do not retain proposals.
- Confirm whether V01-ADR-093 needs a short compatibility note; leave ADR unchanged if the behavior is a contract clarification under the existing authoring transaction model.

## Done condition

- `SPEC-design-records-mcp-tools` describes accurate update diff behavior.
- `SPEC-design-records-mcp-tools` describes no-op update behavior.
- Any ADR impact is explicitly classified as changed or no-change.

## Verification

- Review the updated spec against V01-REQ-MCP-027 acceptance criteria.
- Run Design Records validation for affected records when available.

## Evidence

Verdict: PASS.

Spec contract updated in `SPEC-design-records-mcp-tools` for V01-REQ-MCP-027.

Defined behavior:

- `propose_record_update` no-op definition: an update request whose proposed persisted content is byte-equivalent to current persisted file content after requested update operation semantics are applied.
- No-op updates must not create retained proposals.
- No-op responses use `proposal_created:false`, `operation:"update"`, target information, `validation.ok:true` when there are no error diagnostics, and an info diagnostic with category `no_op_update`.
- No-op responses omit `proposal_id` because no retained proposal exists.
- No-op responses omit or null `diff` because there is no retained proposal diff.
- `no_op_update` is an authoring response diagnostic, not a tool execution error and not a validation error.

Diff contract updated:

- Existing-file update proposals must compare current persisted content with proposed persisted content.
- Modify proposal `diff.text` must be git-style unified diff.
- Modify diffs must include `diff --git a/<path> b/<path>`, `--- a/<path>`, `+++ b/<path>`, and `@@` hunk headers when content differs.
- `index <oldhash>..<newhash> 100644` or equivalent stable old/new content hash representation is required when available.
- Metadata-only real changes must show changed lines plus bounded context, not whole-file addition output.
- Whole-file `+` output remains valid only for actual create/add proposals.

ADR impact:

- V01-ADR-093 required no change. This is a contract clarification and correction under the existing proposal-first authoring transaction model, not a new transaction model decision.

Implementation handoff:

- Proceed to `V01-TASK-MCP-023-03`.
- Implementation should make diff generation change-aware and add no-op detection after operation semantics and normalization, before proposal retention.

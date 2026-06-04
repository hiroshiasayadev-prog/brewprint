# TASK-MCP-019-02: Update tools spec for required heading case canonicalization

- **id**: TASK-MCP-019-02
- **status**: done
- **date**: 2026-06-05
- **work_item**: WORK-MCP-019
- **source_requirement**: REQ-MCP-021
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-MCP-019-01
- **outputs**:
  - `SPEC-design-records-mcp-tools` contract update for case-only required-heading canonicalization
  - Diagnostic contract for validation repair hints
  - Implementation and test inputs for later tasks

## Goal

Update the Design Records MCP tools spec so REQ-MCP-021 has a precise public contract before implementation.

## Work

- Update `SPEC-design-records-mcp-tools` around `propose_record_update` named section replacement.
- Define that exact heading matching remains the default.
- Define a narrow case-only fallback for validation-required workflow headings only.
- Define that a successful fallback must canonicalize the matched heading in the retained proposal diff.
- Define ambiguous case-insensitive match behavior and candidate heading reporting.
- Define validation repair diagnostic behavior for case-only required-heading mismatches while keeping validation strict.
- Confirm ADR update is not required unless the transaction model boundary changes.

## Done condition

- The tools spec defines the narrow case-only fallback scope.
- The tools spec excludes user-defined optional headings from canonicalization.
- The tools spec defines proposal diff canonicalization behavior.
- The tools spec defines ambiguity and non-case mismatch behavior.
- The tools spec defines validation repair diagnostics without relaxing required-section validation.
- Follow-up implementation and test tasks have clear inputs.

## Verification

- Review the updated spec section for consistency with REQ-MCP-021 and TASK-MCP-019-01 evidence.
- Validate related design records with Design Records MCP where possible.

## Evidence
Spec update completed and reviewed.

Initial spec update:

- `SPEC-design-records-mcp-tools` added the `section_heading_case_mismatch` diagnostic category as info-level repair guidance.
- `propose_record_update` named section replacement now defines narrow case-only fallback for validation-required workflow headings.
- The fallback excludes user-defined optional headings and non-required authoring guide headings.
- Successful fallback must canonicalize the matched heading in the retained proposal diff.
- Ambiguous case-insensitive matches fail with candidate headings.
- Non-case differences remain governed by existing exact selector rules.

Codex review result:

- Verdict: NEEDS_FIX before implementation.
- The review found that `schema.md` also needed the new diagnostic category and field contract.
- The review found that fallback scope needed to be target-record-kind specific.
- The review found that validation repair diagnostics should be MUST, not MAY, for the narrow case where the canonical required heading is missing and exactly one case-only non-canonical heading exists.
- The review confirmed that ADR-093 does not need an update.

Fixes applied:

- Updated `SPEC-design-records-mcp-schema` Diagnostic category section to include `section_heading_case_mismatch`.
- Defined required diagnostic fields: `section`, `actual_heading`, and `status`; `candidate_headings` should be included when available.
- Updated tools/spec wording so the heading must be required for the target record kind.
- Explicitly stated that authoring selector fallback does not require the target record to currently be in the gated status.
- Changed validation repair diagnostic wording from MAY to MUST for the narrow case.
- Updated spec `last_updated` fields to `2026-06-05` for the touched spec files.

Final review result:

- Verdict: PASS_WITH_MINOR_NOTES.
- Previous NEEDS_FIX findings are resolved.
- Minor note 1: fallback success wording should not use optional `MAY` language.
- Minor note 2: fallback should explicitly inherit `section_selector.level` constraint.

Final minor fixes applied:

- Reworded fallback success so an eligible single case-insensitive match resolves through fallback and proposal creation proceeds unless an independent proposal-preparation error applies.
- Added explicit wording that fallback candidate matching uses the same `level` constraint before determining zero, one, or ambiguous case-insensitive matches.

Accepted MCP proposals:

- `pw_000086`: initial tools diagnostic update.
- `pw_000088`: initial tools named section replacement update.
- `pw_000099`: schema diagnostic category update.
- `pw_000104`: tools diagnostic review fix.
- `pw_000109`: tools named section replacement review fix.
- `pw_000159`: final minor tools named section replacement fix.

Note:

- Front matter date updates were applied by minimal filesystem line edit because MCP `metadata_block_replace` would have reformatted the YAML front matter unexpectedly.
- Discarded proposal `pw_000102` was not written.

Next step: proceed to `TASK-MCP-019-03` for regression test creation.

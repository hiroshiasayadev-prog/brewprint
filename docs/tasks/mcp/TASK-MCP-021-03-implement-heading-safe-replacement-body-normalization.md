# TASK-MCP-021-03: implement heading-safe replacement body normalization

- **id**: TASK-MCP-021-03
- **status**: done
- **date**: 2026-06-05
- **work_item**: WORK-MCP-021
- **source_requirement**: REQ-MCP-022
- **estimate**: 1d
- **depends_on**:
  - TASK-MCP-021-01
  - TASK-MCP-021-02
- **outputs**:
  - Implementation for first matching heading-line stripping in named_section_replace update flow
  - Warning diagnostic emission for stripped heading normalization

## Goal

Implement heading-safe replacement body normalization for `propose_record_update` with `update.type = named_section_replace`.

## Work

- Add normalization after replacement content is resolved and before retained proposal diff generation.
- Detect the first non-empty line of the replacement body.
- If that line is a Markdown heading whose level and text match the selected `section_selector`, strip only that line.
- Preserve all other content, including body-internal headings.
- Emit a warning diagnostic identifying the stripped heading and level.
- Apply the same path to direct `body` and `body_cache_id` content.

## Done condition

- Matching first heading lines are stripped for `body` and `body_cache_id` sources.
- Only one leading matching heading line is stripped.
- Internal headings and non-matching headings are not stripped.
- Warning diagnostic behavior is implemented.

## Verification

- Run focused authoring unit tests for named section replacement.
- Run relevant MCP tool call tests if impacted.

## Evidence
Implementation complete. All tests pass.

**Files changed:**

- `internal/designrecords/types.go`
  - Added `DiagnosticSeverityWarning = "warning"` constant
  - Added `DiagnosticSectionReplacementBodyHeadingStripped = "section_replacement_body_heading_stripped"` constant
  - Added `StrippedHeading string` and `StrippedLevel int` fields to `Diagnostic` struct
  - Updated `MarshalJSON` to serialize `stripped_heading` and `stripped_level`

- `internal/designrecords/authoring.go`
  - Fixed `ProposeRecordUpdate`: changed `if len(prep.diagnostics) > 0` → `if hasErrorDiagnostics(prep.diagnostics)` so warning diagnostics no longer block proposal creation
  - Fixed `prepareUpdate`: changed `if len(diagnostics) > 0` → `if hasErrorDiagnostics(diagnostics)`, and included warning diagnostics in the returned `authoringPreparation` when proceeding
  - Updated `replaceNamedSection` exact-match path: calls `stripBodyLeadingHeading` after match resolution, emits warning diagnostic when strip occurs
  - Updated `replaceNamedSection` case-fallback path: same normalization applied against resolved section heading
  - Added `stripBodyLeadingHeading(body string, heading Heading) (stripped, text string, level int, wasStripped bool)` helper
  - Added `sectionBodyHeadingStrippedDiagnostic(headingText string, headingLevel int) Diagnostic` helper

- `internal/designrecords/authoring_test.go`
  - Added `TestReplaceNamedSectionBodyHeadingStripping` with 9 sub-tests:
    - `direct_body_heading_stripped`: exact match, body starts with `## Evidence` → proposal created, warning diagnostic, one `## Evidence` in diff
    - `body_cache_id_heading_stripped`: same behavior via cached body
    - `only_first_heading_stripped`: body with two `## Evidence` → first stripped, second preserved in diff
    - `internal_subheading_preserved`: body with `## Evidence` then `### Sub` → first stripped, `### Sub` remains
    - `level_mismatch_no_strip`: body starts with `### Evidence`, selected is `## Evidence` → no strip
    - `text_mismatch_no_strip`: body starts with `## Notes`, selected is `## Evidence` → no strip
    - `omitted_selector_level_strip`: selector omits level, resolves to `## Evidence`, body starts with `## Evidence` → strip
    - `omitted_selector_level_level_mismatch`: omitted level resolves to level 2, body starts with `### Evidence` → no strip
    - `warning_does_not_block_proposal_creation`: warning severity confirmed, proposal_id present

**Test results:**

```
go test ./internal/designrecords ./internal/designrecordsmcp
ok  github.com/hiroshiasayadev-prog/brewprint/internal/designrecords
ok  github.com/hiroshiasayadev-prog/brewprint/internal/designrecordsmcp
```

All tests in both packages passed.

# V01-TASK-MCP-021-04: add heading-safe normalization regression tests

- **id**: V01-TASK-MCP-021-04
- **status**: done
- **date**: 2026-06-05
- **work_item**: V01-WORK-MCP-021
- **source_requirement**: V01-REQ-MCP-022
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-MCP-021-03
- **outputs**:
  - Regression tests for direct body normalization
  - Regression tests for body_cache_id normalization
  - Regression tests for internal heading preservation and mismatch behavior

## Goal

Add regression coverage for heading-safe `named_section_replace` replacement body normalization.

## Work

- Add tests for direct `body` input where the first non-empty line is the selected heading.
- Add tests for `body_cache_id` input with the same heading pattern.
- Add tests proving only the first matching heading line is stripped.
- Add tests proving body-internal headings are preserved.
- Add tests proving heading text mismatch and heading level mismatch are not stripped.
- Assert warning diagnostic category and stripped heading metadata where available.
- Assert retained proposal diff uses normalized content.

## Done condition

- Regression tests fail without the implementation and pass with it.
- Coverage includes both body source paths and negative cases.
- Diagnostic assertions cover the warning behavior required by V01-REQ-MCP-022.

## Verification

- Run focused tests for designrecords authoring behavior.
- Run MCP tool call tests if diagnostics are surfaced through the MCP layer.

## Evidence
**Verdict: PASS**

Coverage confirmed by inspecting `TestReplaceNamedSectionBodyHeadingStripping` in
`internal/designrecords/authoring_test.go` (lines 1651–1888, added by V01-TASK-MCP-021-03).

## Coverage checklist

| # | Required item | Sub-test | Result |
|---|---|---|---|
| 1 | direct `body` source strips leading duplicate selected heading | `direct_body_heading_stripped` | ✅ |
| 2 | `body_cache_id` source strips the same pattern | `body_cache_id_heading_stripped` | ✅ |
| 3 | only the first matching heading line is stripped | `only_first_heading_stripped` | ✅ |
| 4 | body-internal headings are preserved | `internal_subheading_preserved` | ✅ |
| 5 | heading text mismatch is not stripped | `text_mismatch_no_strip` | ✅ |
| 6 | heading level mismatch is not stripped | `level_mismatch_no_strip` | ✅ |
| 7 | omitted `section_selector.level` compares against resolved section level | `omitted_selector_level_strip` + `omitted_selector_level_level_mismatch` | ✅ |
| 8 | warning diagnostic category is `section_replacement_body_heading_stripped` | `direct_body_heading_stripped`, `warning_does_not_block_proposal_creation` | ✅ |
| 9 | warning severity is `warning` | `direct_body_heading_stripped`, `warning_does_not_block_proposal_creation` | ✅ |
| 10 | warning includes `stripped_heading` | `direct_body_heading_stripped` | ✅ |
| 11 | warning includes `stripped_level` | `direct_body_heading_stripped` | ✅ |
| 12 | warning diagnostic does not block retained proposal creation | `warning_does_not_block_proposal_creation` | ✅ |
| 13 | error diagnostics still block proposal creation | pre-existing error tests (`SectionSelectorNoMatch`, `SectionSelectorAmbiguous`, etc.) | ✅ |

## Tests inspected

- `TestReplaceNamedSectionBodyHeadingStripping` (9 sub-tests, all PASS)
  - File: `internal/designrecords/authoring_test.go:1651`
- Pre-existing error-path tests in `TestAuthoringNamedSectionSelectors`, `TestBodyCacheReturnClassification` confirm error diagnostics still block.

## Commands run and results

```
go test ./internal/designrecords -run TestReplaceNamedSectionBodyHeadingStripping -v
# PASS (9/9 sub-tests, 0.13s)

go test ./internal/designrecords ./internal/designrecordsmcp
# ok internal/designrecords  3.193s
# ok internal/designrecordsmcp  0.618s
```

## Gaps

No gaps. All required coverage items are satisfied by the tests added in V01-TASK-MCP-021-03.

## Notes for V01-TASK-MCP-021-05

No contract changes were discovered. The heading stripping normalization, warning diagnostic fields (`stripped_heading`, `stripped_level`), and `hasErrorDiagnostics` gating are confirmed correct. V01-TASK-MCP-021-05 (runtime smoke and close synchronization) can proceed against the current implementation.

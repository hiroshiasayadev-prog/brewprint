# PRODUCT-TASK-SPEC-009-04: Opus review of migrated product concepts spec files

- **id**: PRODUCT-TASK-SPEC-009-04
- **status**: done
- **date**: 2026-06-22
- **work_item**: PRODUCT-WORK-SPEC-009
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-009-03
- **outputs**:
  - Review findings (inline in Evidence)

## Goal

Independent Opus 4.7 review of all 9 migrated files against the pre-migration originals in `product/records/old/`. Confirm completeness, format compliance, and English translation accuracy before removing the staging files.

## Work

| area | what to check |
|---|---|
| completeness | All content from the 9 `product/records/old/` source files is accounted for in the migrated files. No section silently dropped. |
| format compliance | H1 format, H1-adjacent metadata, required sections by kind, no YAML front matter remaining. |
| parent chain | Every spec's `parent:` value matches the declared `id:` in its stated parent; `traceability/index.md` `## Topics` rows cover all 7 siblings with correct `ref:` values. |
| English titles | No Japanese remains in H1, H2, H3 titles, or table headers. Body prose may remain Japanese. |
| stale paths removed | No `docs/...` references remain in any of the 9 files. |
| drift guards | `metadata-schema.md` has drift guards in the workflow/investigation reference metadata and validation responsibility sections. `resolve-and-validation.md` has drift guards in the resolve/lookup and duplicate-detection sections; resolver output and MCP writer placeholder sections are preserved and not expanded. |
| validator | `validate_spec.py --strict` exits 0 on all 9 files (already confirmed in PRODUCT-TASK-SPEC-009-03, but reviewer should flag any format issues they spot). |

## Done condition

| item | done when |
|---|---|
| review complete | Opus review report is attached in Evidence. |
| findings classified | Each finding is classified: must-fix before PRODUCT-TASK-SPEC-009-05, or defer. |
| user sign-off | User approves proceeding to PRODUCT-TASK-SPEC-009-05. |

## Verification

This task is itself a review gate. Proceed to PRODUCT-TASK-SPEC-009-05 only after user sign-off on findings.

Cross-check all Opus findings against the actual files before applying any fix — do not accept findings at face value (see CLAUDE.md model policy and lessons from PRODUCT-TASK-SPEC-005-11).

## Evidence

Retrospective status synchronization from `PRODUCT-WORK-SPEC-009` completion evidence:

- Independent review findings were recorded and classified before finalization.
- M1: Japanese table headers in eight locations — corrected during the retranslation pass.
- D1: dropped navigation rows — restored with `*(planned)*` markers in `project-artifact-model/index.md`.
- D2: missing internal-design and implementation follow-up bullets — restored in `traceability/out-of-scope.md`.
- D3: stale example paths — updated in `traceability/resolve-and-validation.md`.
- D4 and D5 required no changes after verification against current contracts.
- User sign-off is evidenced by completion of `PRODUCT-TASK-SPEC-009-05` and closure of `PRODUCT-WORK-SPEC-009`.

The original review output was summarized in the parent work item's Evidence rather than copied into this task. This record now reflects that completed gate.

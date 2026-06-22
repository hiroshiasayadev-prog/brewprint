# PRODUCT-TASK-SPEC-005-11: BPDSL — Opus review of views/ spec files

- **id**: PRODUCT-TASK-SPEC-005-11
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-10
- **outputs**:
  - Review findings (inline in Evidence)

## Goal

Independent Opus review of the `bpdsl/records/spec/views/` migration output against the content staged in `bpdsl/old_views/`. Confirm completeness, format compliance, and `contract_class: format` correctness before removing the staging files.

## Work

| area | what to check |
|---|---|
| completeness | All content from `bpdsl/old_views/` files is accounted for in `views/`. No section silently dropped. |
| format compliance | H1 format, H1-adjacent metadata, required sections per kind, `## Topics` columns. |
| contract shape | Each view spec has `## Current contract` / `## Rules` / `## Validation rules` and `contract_class: format`. |
| parent chain | Every child spec's `parent:` matches an existing `## Topics` row in `views/overview.md`. |
| translation / encoding | No mojibake or unresolved encoding artifacts (6 of 7 source files had corrupted H1 bytes). No Japanese remains in H1 titles, H2 section titles, or table headers. |
| internal refs | Cross-spec `[]()` links point to files that exist under `views/` (not to `bpdsl/old_views/`). |
| validator | `validate_spec.py bpdsl/records/spec/views/ --strict` exits 0. |

## Done condition

| item | done when |
|---|---|
| review complete | Opus review report is attached in Evidence. |
| findings classified | Each finding is classified: must-fix before 005-12, or defer. |
| user sign-off | User approves proceeding to 005-12. |

## Verification

This task is itself a review gate. Proceed to PRODUCT-TASK-SPEC-005-12 only after user sign-off on findings.

## Evidence

- Opus review spawned via Agent tool (model: opus). All 8 files read in full (including `dag.md` at 992-line source).
- Verdict: **pass-with-findings**. One must-fix, two should-fix, two defer findings identified.

### Finding disposition

| finding | severity | disposition | action taken |
|---|---|---|---|
| `views/overview.md` Topics headers use `title\|kind\|ref\|summary` instead of `slug\|kind\|id\|description` | must-fix (by Opus) | **false alarm** — constraint in review prompt was wrong. `mcp/overview.md` uses identical headers and passed --strict. No fix. | None |
| `er.md` `## Rules` missing generation-source statement (`render_er`, no hand-written Mermaid) | should-fix | **real omission** — `state-diagram.md` and `sequence-diagram.md` both have equivalent statements. Fixed. | Added one sentence to `er.md` `## Rules` end. |
| `state-diagram.md` choice-pseudostate Validation rules adds "treated as a renderer bug" | should-fix (by Opus) | **false alarm** — Opus cited text that does not exist in the file. Actual line 127: "Mermaid draws it as a normal node rather than a diamond — it must always be collected into the leading block." Faithful translation of source. No fix. | None |
| `dag.md` hex color case inconsistency (lowercase `#e0e0e0` vs uppercase `#E0E0E0`) | defer | Inherited from source; not a migration defect. | None |
| All other areas | — | No findings (api-table, wireframe, sequence-diagram: clean; dag: clean; er, model-file: clean after fix). | — |

- `validate_spec.py bpdsl/records/spec/views/ --strict` → `All 8 file(s) OK.` (post-fix)

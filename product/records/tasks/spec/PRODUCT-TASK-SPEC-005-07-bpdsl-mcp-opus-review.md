# PRODUCT-TASK-SPEC-005-07: BPDSL — Opus review of mcp/ spec files

- **id**: PRODUCT-TASK-SPEC-005-07
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-06
- **outputs**:
  - Review findings (inline in Evidence)

## Goal

Independent Opus review of the `bpdsl/records/spec/mcp/` migration output against the content staged in `bpdsl/old_mcp/`. Confirm completeness, format compliance, and `contract_class: interface` correctness before removing the staging files.

## Work

| area | what to check |
|---|---|
| completeness | All content from `bpdsl/old_mcp/` files is accounted for in `mcp/`. No section silently dropped. |
| format compliance | H1 format, H1-adjacent metadata, required sections per kind, `## Topics` columns. |
| contract shape | Each tool spec has `## Request` / `## Response` / `## Errors` and `contract_class: interface`. |
| parent chain | Every child spec's `parent:` matches an existing `## Topics` row in the declared parent. |
| translation | No Japanese remains in H1 titles, H2 section titles, or table headers. (Note: an earlier mojibake concern for `overview.md`'s source H1 was a false alarm — Bash/grep was printing through the Windows console codepage; the Read tool decoded the source cleanly. No encoding corruption was found in any source file.) |
| internal refs | Cross-spec `[]()` links point to files that exist under `mcp/` (not to `bpdsl/old_mcp/`). |
| validator | `validate_spec.py bpdsl/records/spec/mcp/ --strict` exits 0. |

## Done condition

| item | done when |
|---|---|
| review complete | Opus review report is attached in Evidence. |
| findings classified | Each finding is classified: must-fix before 005-08, or defer. |
| user sign-off | User approves proceeding to 005-08. |

## Verification

This task is itself a review gate. Proceed to PRODUCT-TASK-SPEC-005-08 only after user sign-off on findings.

## Evidence

Opus review verdict (spawned via Agent tool, model opus): **pass with minor findings**. `validate_spec.py bpdsl/records/spec/mcp/ --strict` → `[strict] All 12 file(s) OK.` All 12 old files confirmed fully migrated with no content silently dropped; parent chain confirmed clean from all 11 children to `spec:bpdsl.mcp.overview` to `spec:bpdsl.overview`; all 8 tool specs confirmed `contract_class: interface` with matching `## Request`/`## Response`/`## Errors` content; zero Japanese found; all internal links resolve (including forward links into not-yet-migrated `dsl/`/`views/` files, correctly not flagged as mcp/-migration defects); all 12 ids match path-derived form.

Findings cross-checked against source and disposed of:

| finding | category | disposition |
|---|---|---|
| `inspect.md` task-inspect example dropped 2nd `sub_tasks` entry (`reserve_inventory`) and 2nd flow entry | completeness | **Restored** — confirmed against `bpdsl/old_mcp/tools/inspect.md`; illustrative duplicate, but cheap to keep faithful. |
| `inspect.md` field-inspect example dropped 2nd `field_fk` (`payment_event.order_id`) | completeness | **Restored** — same fidelity reasoning. |
| `get-reference-tree.md` calls `analyze_impact` "the future ... " and links to itself correctly but with stale wording | completeness | **Fixed** — confirmed old source wrote this before `analyze_impact` existed (pointed at `versioning.md` as placeholder); new spec correctly links to the real `analyze-impact.md` tool spec but kept the now-inaccurate "future" qualifier. Removed "future." |
| `versioning.md` reworded "track in `docs/TASKS.md`" to generic guidance | completeness | **Left as-is** — intentional; `docs/TASKS.md` doesn't exist in the new repo layout, so a literal carryover would be a dead reference. |
| ADRs cited only in old front-matter `depends_on` (not old body) are absent from new body `> Source:` lines | adr-attribution | **Left as-is** — migration rule is front-matter removed, body citations preserved; front-matter-only ADRs were never body-cited in the source either. |

Re-ran `validate_spec.py bpdsl/records/spec/mcp/ --strict` after edits → `[strict] All 12 file(s) OK.`

User sign-off received 2026-06-17. Proceeding to PRODUCT-TASK-SPEC-005-08.

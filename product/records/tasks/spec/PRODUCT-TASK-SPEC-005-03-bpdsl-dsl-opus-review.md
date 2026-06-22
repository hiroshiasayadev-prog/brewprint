# PRODUCT-TASK-SPEC-005-03: BPDSL — Opus review of dsl/ spec files

- **id**: PRODUCT-TASK-SPEC-005-03
- **status**: done
- **date**: 2026-06-16
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-02
- **outputs**:
  - Review findings (inline in Evidence)

## Goal

Independent Opus review of the `bpdsl/records/spec/dsl/` migration output against the content in `bpdsl/old/`. Confirm completeness, format compliance, and translation accuracy before removing the staging files.

## Work

| area | what to check |
|---|---|
| completeness | All content from `bpdsl/old/` files is accounted for in `dsl/`. No section silently dropped. |
| format compliance | H1 format, H1-adjacent metadata, required sections by kind, `## Topics` columns. |
| parent chain | Every child spec's `parent:` matches an existing `## Topics` row in the declared parent. |
| English translation | No Japanese remains in H1 titles, H2 section titles, or table headers. Body prose may retain Japanese until a separate pass if needed. |
| internal refs | Cross-spec `[]()` links point to files that exist under `dsl/` (not to `bpdsl/old/`). |
| validator | `validate_spec.py bpdsl/records/spec/dsl/ --strict` exits 0. |

## Done condition

| item | done when |
|---|---|
| review complete | Opus review report is attached in Evidence. |
| findings classified | Each finding is classified: must-fix before 005-04, or defer. |
| user sign-off | User approves proceeding to 005-04. |

## Verification

This task is itself a review gate. Proceed to PRODUCT-TASK-SPEC-005-04 only after user sign-off on findings.

## Evidence

Opus review verdict: **pass with minor findings**. `validate_spec.py bpdsl/records/spec/dsl/ --strict` → `[strict] All 16 file(s) OK.` Coverage check confirmed all 7 `bpdsl/old/` files fully migrated; parent chain confirmed clean from all 16 leaves to `spec:bpdsl.overview`; no Japanese remains; no links into `bpdsl/old/`.

Findings cross-checked against source and disposed of:

| finding | category | disposition |
|---|---|---|
| `nodes/data.md` → `type-ref.md §Applicability` (heading doesn't exist) | internal-ref | **Fixed** — corrected to `§Fields that accept TypeRef`. |
| `nodes/processing.md` → `data-flow.md §Branch entry` (heading doesn't exist) | internal-ref | **Fixed** — corrected to `§Exclusive branch entry`. |
| `naming.md` dropped 2x "MCP schema / ObjectRef migration is out of scope" disclaimer | completeness | **Restored** — real scope-boundary statement, not filler. |
| `nodes/data.md` dropped "model file render / catalog ownership → views/model-file.md" pointer | completeness | **Restored** — points to existing `bpdsl/records/spec/views/model-file.md` (pre-T06, still old format but file exists at that path). |
| `nodes/overview.md` / `nodes.md` old §74 minimum-scope disclaimer (helper model/tagged_union/etc out of scope) | completeness | **Left dropped** — verified stale: `tagged_union` is already in scope in `nodes/data.md`, so restoring would contradict current content. |
| `file-types.md` dropped §6 implementation reference (`classify.go`/`loader.go`) | completeness | **Left dropped** — implementation file pointer, not contract; consistent with no other `dsl/` Reference spec citing Go source paths. |
| `project-layout.md` dropped §8 `go-mN-summary.md` reference | completeness | **Left dropped** — internal/transient work-tracking doc, not a stable spec target. |
| ADR section-anchors stripped (e.g. `§1, §9` → bare ADR id); V01-REQ/V01-TASK refs dropped | adr-attribution | **Left as-is** — ADR id set itself preserved in full; REQ/TASK are work-tracking refs outside the `spec` concept. |

Re-ran `validate_spec.py bpdsl/records/spec/dsl/ --strict` after all edits → `[strict] All 16 file(s) OK.`

User sign-off received 2026-06-17. Proceeding to PRODUCT-TASK-SPEC-005-04.

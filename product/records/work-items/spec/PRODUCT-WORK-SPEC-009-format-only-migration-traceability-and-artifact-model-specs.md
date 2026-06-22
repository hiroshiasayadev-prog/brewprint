# PRODUCT-WORK-SPEC-009: Format-only migration of traceability and artifact model specs

- **id**: PRODUCT-WORK-SPEC-009
- **status**: done
- **date**: 2026-06-15
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **impact_refs**:
- **tasks**:
  - PRODUCT-TASK-SPEC-009-01
  - PRODUCT-TASK-SPEC-009-02
  - PRODUCT-TASK-SPEC-009-03
  - PRODUCT-TASK-SPEC-009-04
  - PRODUCT-TASK-SPEC-009-05

## Goal

Apply the accepted spec format (PRODUCT-WORK-SPEC-001 contract) to `project-artifact-model/index.md` and `traceability/**` without executing ownership relocation. This is the format-migration phase split from PRODUCT-WORK-SPEC-005, enabling format compliance before PRODUCT-WORK-SPEC-004 ownership boundary decisions.

## Boundary

Format migration only. No ownership relocation, no hybrid-section ownership splits. File-level splitting for readability is permitted — if a target file is split into topic-level sibling files, the 9-file list expands to however many result. Stale `docs/...` path cleanup identified by PRODUCT-INV-SPEC-002 is included as migration cleanup, not ownership change.

`namespace-model/index.md` added as 9th target per PRODUCT-TASK-SPEC-004-03 (accepted in PRODUCT-WORK-SPEC-004): it is pre-migration format and will receive relocated content from DRMCP in Phase 2; format migration must complete first.

## Impact Scope

9 spec files under `product/records/spec/concepts/`:

- `project-artifact-model/index.md`
- `traceability/index.md`
- `traceability/artifact-refs.md`
- `traceability/coverage-mapping.md`
- `traceability/metadata-schema.md`
- `traceability/out-of-scope.md`
- `traceability/resolve-and-validation.md`
- `traceability/semantic-ref.md`
- `namespace-model/index.md`

## Task flow

| task | description | gate |
|---|---|---|
| T01 | Inventory — run `validate_spec.py` on all 9 target files and record current diagnostics; stage all 9 to `product/records/old/`. | — |
| T02 | Migrate all 9 files in place to accepted spec format; add drift guards to hybrid sections in `metadata-schema.md` and `resolve-and-validation.md`. | — |
| T03 | Run `validate_spec.py --strict` on all 9 migrated files; confirm exit 0. | — |
| T04 | Independent Opus 4.7 review — compare `product/records/old/` (before) against migrated files (after). | **GATE** |
| T05 | Apply review findings; remove `product/records/old/`; run final `validate_spec.py --strict`; close work item. | — |

## Task Candidates

- PRODUCT-TASK-SPEC-009-01
- PRODUCT-TASK-SPEC-009-02
- PRODUCT-TASK-SPEC-009-03
- PRODUCT-TASK-SPEC-009-04
- PRODUCT-TASK-SPEC-009-05

## Completion Condition

- All 9 target files pass PRODUCT-WORK-SPEC-006 format validation.
- Stale `docs/...` path references removed or updated in target files.
- No ownership relocation or hybrid-section ownership split performed. File-level splits for readability are acceptable.
- Drift guards documented in `metadata-schema.md` hybrid sections (workflow/investigation reference metadata; validation responsibility) (see PRODUCT-TASK-SPEC-004-01 decision).
- Drift guards documented in `resolve-and-validation.md` hybrid sections (resolve/resolver input/lookup; duplicate detection/unresolved refs/declared relation integrity) (see PRODUCT-TASK-SPEC-004-01 decision).

## Evidence

### Migration complete (2026-06-22)

**Source files:** 9 spec files under `product/records/spec/concepts/` (3 Overviews: traceability, project-artifact-model, namespace-model)

**Output files:** 18 files — 7 traceability (1:1), 5 project-artifact-model (1→5 split), 6 namespace-model (1→6 split)

**T02 corrections applied:**
- All YAML front matter removed; H1 `# <Kind>: <Title>` added with H1-adjacent metadata
- All body prose translated to English (original T02 instruction "body prose may remain Japanese" was incorrect per user correction)
- Drift guards added: `metadata-schema.md` (3 sections), `resolve-and-validation.md` (6 sections)
- Stale `docs/...` path hyperlinks removed; table cell paths updated to `<namespace>/records/` pattern
- `REFERENCE_NO_TABLE`: `design-flow.md` and `change-and-investigation-flow.md` received summary tables

**T04 review findings disposition:**
- M1 (Japanese table headers, 8 locations) — applied in T02 retranslation pass
- D1 (dropped navigation rows) — placeholder rows added with `*(planned)*` markers in `project-artifact-model/index.md`
- D2 (internal-design/impl bullets in out-of-scope) — placeholder bullets added with `*(planned)*` markers in `traceability/out-of-scope.md`
- D3 (stale example paths) — `resolve-and-validation.md` example block updated to current namespace-qualified spec refs and full V01 filenames
- D4/D5 — no action required; consistent with current state

**Staging directory:** `product/records/old/` removed after final validation

**Final validator:** `[strict]  All 18 file(s) OK.`

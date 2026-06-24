# PRODUCT-TASK-SPEC-012-11: Validate and clean migration diagnostics

- **id**: PRODUCT-TASK-SPEC-012-11
- **status**: done
- **date**: 2026-06-25
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-10
- **outputs**:
  - Validation results recorded in this task
  - Migration-caused stale refs and structural diagnostics corrected

## Goal

Validate the restructured tree and correct only diagnostics caused by this migration.

## Work

- Run available PRODUCT spec validation against the final target tree.
- Check canonical refs, parent markers, topic tables, H1 shape, and required sections.
- Run scoped stale-ref searches for removed paths and refs.
- Compare diagnostics with the T01 baseline.
- Correct migration-caused failures.
- Record pre-existing or unrelated diagnostics without expanding scope.
- Confirm `product/records/spec/concepts/` contains no active specifications.

## Done condition

- No unresolved reference or structural diagnostic is caused by this migration.
- Old target paths and canonical refs are absent from active files.
- Pre-existing diagnostics remain separately attributed.
- `v01/` is unchanged.
- Validation commands and complete results are recorded.

## Verification

- Re-run all scoped validation after corrections.
- Confirm validation results against the T01 baseline.
- Confirm stale-ref searches return no migration-caused matches.
- Confirm the final PRODUCT spec tree matches `PRODUCT-ADR-SPEC-001`.

## Evidence

### 1. Execution summary

**Status**: done
**Shell execution available**: yes
**Files changed by T11**:
- `product/records/spec/design-records/artifact-model/traceability-boundary.md` — Candidate 1 path corrected
- `product/src/tools/validate_spec.py` — Candidate 2 comment ref corrected
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-11-validate-and-clean-migration-diagnostics.md` — this file (status + Evidence)

**Corrections applied**: 2 (both known candidates, mechanical)
**Migration-caused diagnostics found and corrected**: 2
**Pre-existing / deferred / false-positive diagnostics**: 4 (see §10)
**Historical-evidence occurrences preserved**: multiple (see §5–7)
**No T12 semantic review edits or T13 Work Item closure edits performed.**

### 2. Baseline comparison

| check | T01 baseline | T11 final |
|---|---|---|
| PRODUCT spec count | 40 | 47 |
| PRODUCT strict diagnostics | 0 | 0 (exit 0) |
| DRMCP strict diagnostics | 0 | 0 (exit 0) |
| BPDSL strict diagnostics | 0 | 0 (exit 0) |
| active `spec:product.concepts.*` refs | 266 occurrences before migration | 0 active (historical Evidence only) |
| active `product/records/spec/concepts/` paths | — | 0 active after Candidate 1 correction |
| active concepts specs | 40-tree included concepts | 0 under `concepts/` |
| combined ID graph | not available at T01 | 114 IDs, 0 duplicates, all parents resolve |

Notes:
- T01 had no strict validator diagnostics; the same holds at T11.
- File-count increase from 40 to 47 is expected: routing overviews, split areas, compatibility sub-areas, and BPDSL staging added by T02–T10.
- Historical old refs (217 at T10) are not unresolved active references; they are preserved Evidence rows in completed task and work-item records.
- Combined graph validation is a T11 addition — the temporary validator does not resolve refs.

### 3. Strict validator results

Commands run from repo root (`C:\Users\imved\projects\brewprint`):

```powershell
python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color
python -X utf8 product/src/tools/validate_spec.py drmcp/records/spec/design-records-mcp --strict --no-color
python -X utf8 product/src/tools/validate_spec.py bpdsl/records/spec --strict --no-color
```

Pre-correction run:

| root | output | exit |
|---|---|---|
| `product/records/spec` | `[strict]  All 47 file(s) OK.` | 0 |
| `drmcp/records/spec/design-records-mcp` | `[strict]  All 30 file(s) OK.` | 0 |
| `bpdsl/records/spec` | `[strict]  All 37 file(s) OK.` | 0 |

Post-correction re-run:

| root | output | exit |
|---|---|---|
| `product/records/spec` | `[strict]  All 47 file(s) OK.` | 0 |
| `drmcp/records/spec/design-records-mcp` | `[strict]  All 30 file(s) OK.` | 0 |
| `bpdsl/records/spec` | `[strict]  All 37 file(s) OK.` | 0 |

Total: 114 spec files, zero errors in strict mode.

### 4. Combined spec graph results

Inline Python script run against all three spec roots. Results recorded below.

| check | count | result |
|---|---|---|
| indexed specs | 114 (47+30+37) | pass |
| duplicate IDs | 0 | pass |
| unresolved parents | 0 | pass |
| parent/Topics mismatches | 0 | pass |
| unresolved Topics refs | 0 | pass |
| unresolved Related specs refs | 1 active expected-0 | see §10 |
| unresolved active body refs | 4 items | see §10 |
| accepted historical unresolved refs | 1 | `spec:product.concepts.repository_layout` in `bpdsl/design-flow.md` previous-context row |

The 1 unresolved Related specs ref and 4 unresolved body refs are classified in §10 as pre-existing or false-positive. None are migration-caused.

### 5. Stale canonical-ref audit

Search pattern: `spec:product\.concepts[\.A-Za-z0-9_]*`

**PRODUCT current specs (`product/records/spec/**`)**

| file | line | text | class |
|---|---|---|---|
| `product/records/spec/bpdsl/design-flow.md` | 97 | `spec:product.concepts.repository_layout` in previous-context source-map row | `historical_evidence` |

Active stale refs: **0**.

**PRODUCT workflow metadata fields**

Zero matches in H1-adjacent `- **field**:` metadata of any ADR, investigation, requirement, work-item, or task record.

**PRODUCT workflow Evidence body**

Many matches in completed task and work-item Evidence sections (body text, not metadata fields). All record pre-migration paths and IDs as they existed at execution time. All are `historical_evidence`. Correcting them is prohibited by the correction rules.

Representative files containing historical Evidence matches:
- `product/records/tasks/namespace/PRODUCT-TASK-NAMESPACE-001-*.md` — task Evidence rows
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-011-*.md` — work item Evidence rows
- `product/records/work-items/namespace/PRODUCT-WORK-NAMESPACE-001-*.md` — work item Evidence rows
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-003-*.md`, `PRODUCT-WORK-SPEC-005-*.md`, `PRODUCT-WORK-SPEC-006-*.md`, `PRODUCT-WORK-SPEC-007-*.md` — Evidence rows
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-001-03-*.md`, `PRODUCT-TASK-SPEC-005-*.md`, `PRODUCT-TASK-SPEC-006-*.md` — Evidence rows

**DRMCP records (`drmcp/records/**`)**

Zero matches.

**BPDSL records (`bpdsl/records/**`)**

Zero matches.

**Active supporting instructions (CLAUDE.md)**

Zero matches.

**Source/tooling files**

| file | line | text | class |
|---|---|---|---|
| `product/src/tools/validate_spec.py` | 36 (before correction) | `spec:product.concepts.spec_format.document_shape` | `migration_caused_corrected` |

After correction: `spec:product.design_records.spec_format.document_shape`. Active stale refs in source/tooling: **0**.

### 6. Stale physical-path audit

Search string: `product/records/spec/concepts/`

**Current PRODUCT specs (`product/records/spec/**`)**

| file | line | text | class |
|---|---|---|---|
| `product/records/spec/bpdsl/repository-implementation-flow.md` | 71 | source-map row: old path → preserved here | `historical_evidence` |
| `product/records/spec/bpdsl/repository-implementation-flow.md` | 72 | source-map row: old path → preserved here | `historical_evidence` |
| `product/records/spec/bpdsl/artifact-responsibilities.md` | 46 | source-map row: old path → preserved here | `historical_evidence` |
| `product/records/spec/bpdsl/artifact-responsibilities.md` | 47 | source-map row: old path → preserved here | `historical_evidence` |
| `product/records/spec/bpdsl/design-flow.md` | 87 | source-map row: old path → preserved here | `historical_evidence` |
| `product/records/spec/bpdsl/design-flow.md` | 88 | source-map row: old path → preserved here | `historical_evidence` |
| `product/records/spec/design-records/artifact-model/traceability-boundary.md` | 49 (before correction) | active present-tense statement | `migration_caused_corrected` |

After Candidate 1 correction: line 49 now reads `product/records/spec/design-records/traceability/`. Active stale physical paths in current specs: **0**.

**Source/tooling and supporting docs**: zero matches in `product/src/` or `CLAUDE.md`.

Every residual physical-path occurrence is in a BPDSL staging source-map table. These are historical_evidence and remain unchanged.

### 7. Removed-target audit

Removed targets:
- `spec:product.concepts.traceability.coverage_mapping`
- `spec:product.concepts.traceability.out_of_scope`

| file | line | text | class |
|---|---|---|---|
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-009-02-format-migration.md` | 48 | historical output-map row in format-migration task inventory | `historical_evidence` |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-009-02-format-migration.md` | 50 | historical output-map row in format-migration task inventory | `historical_evidence` |

Active live pointers to removed targets: **0**. Both known occurrences are historical output-map rows. Preserved unchanged.

### 8. Concepts and final-tree audit

**Current product/records/spec/ top-level:**

```text
product/records/spec/
  index.md              ← spec:product (Overview)
  design-records/       ← spec:product.design_records
  brewprint/            ← spec:product.brewprint
  bpdsl/                ← spec:product.bpdsl (non-canonical staging)
  concepts/             ← empty directory shell, no .md files
```

Confirmed:
- No Markdown spec files under `product/records/spec/concepts/` (directory is an empty shell). ✓
- No top-level `compatibility/` area. ✓
- No top-level `deferred-integration/` area. ✓
- `brewprint/compatibility/` exists (contains compatibility Reference specs). ✓
- `brewprint/layout/` exists. ✓
- `brewprint/namespaces/` exists. ✓
- `bpdsl/` declares non-canonical status and exit obligations in its index. ✓

Matches accepted ADR `PRODUCT-ADR-SPEC-001`.

### 9. Corrections

| file | diagnostic | old value | new value | why migration-caused |
|---|---|---|---|---|
| `product/records/spec/design-records/artifact-model/traceability-boundary.md` line 49 | Active present-tense path references the removed `concepts/traceability/` location | `` `product/records/spec/concepts/traceability/` defines semantic ref grammar... `` | `` `product/records/spec/design-records/traceability/` defines semantic ref grammar... `` | T07 relocated `concepts/traceability/` to `design-records/traceability/`. This active sentence was carried as a later advisory in T07/T10 and is now corrected in T11. |
| `product/src/tools/validate_spec.py` line 36 | Active source comment references the removed `spec:product.concepts.spec_format.document_shape` ID | `# Section requirement matrix derived from spec:product.concepts.spec_format.document_shape` | `# Section requirement matrix derived from spec:product.design_records.spec_format.document_shape` | The spec was relocated from `concepts/spec-format/` to `design-records/spec-format/`, changing its path-derived ID. The source comment was not updated during migration. |

No surrounding contract or validator behavior was changed.

### 10. Pre-existing and deferred diagnostics

| finding | file | context | class | action |
|---|---|---|---|---|
| `spec:drmcp.design_records_mcp.tools` (Related specs) | `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md` | Non-goals note and Related specs table reference a planned DRMCP tool spec not yet created. File `drmcp/records/spec/design-records-mcp/tools.md` does not exist. | `pre_existing_unrelated` | Leave for DRMCP tooling work. Not migration-caused. |
| `spec:drmcp.design_records_mcp.tools` (example table) | `product/records/spec/design-records/authoring-standards/spec-authoring.md` line 43 | Example table row illustrating path→ID derivation algorithm. Not an active reference to the spec. | `false_positive` | No action. The example is illustrative. |
| `spec:drmcp.design_records_mcp.tools` (example table) | `product/records/spec/design-records/spec-format/spec-id-as-ref.md` line 53 | Example table row illustrating path→ID derivation algorithm. Not an active reference to the spec. | `false_positive` | No action. The example is illustrative. |
| `spec:trace.semantic` (regex extraction) | `product/records/spec/design-records/spec-format/spec-id-as-ref.md` line 88 | Regex matches `spec:trace.semantic` from inside the string `` `spec:trace.semantic-ref` `` used as a hyphen-form ID example in the Boundary table. Not an actual ref; the example illustrates the legacy hyphen-form pattern. | `false_positive` | No action. Not a real reference. |

Known deferred findings outside T11 scope (already recorded in T08/T09):
- DRMCP visible-metadata compatibility gap (YAML front matter in DRMCP specs): `deferred_semantic_review` — belongs to DRMCP redesign.
- DRMCP P0/write-tool contradiction: `deferred_semantic_review` — belongs to DRMCP spec review.
- BPDSL historical `dsl/` versus current `yaml/` inconsistency: `deferred_semantic_review` — T12 BPDSL review.
- BPDSL self-hosting ownership gap: `deferred_semantic_review` — T12/T13.
- Semantic review of `product/records/spec/bpdsl/` content: `deferred_semantic_review` — T12.
- Work Item closure: `deferred_semantic_review` — T13.

### 11. Scope and Git evidence

**Staged diff**:
```
(no output — nothing staged)
```

**v01 diff**:
```
(no output — v01 unchanged)
```

**T11 edit scope**: exactly 3 files changed:
1. `product/records/spec/design-records/artifact-model/traceability-boundary.md` — one-line physical-path correction
2. `product/src/tools/validate_spec.py` — one-line source comment ref correction
3. `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-11-validate-and-clean-migration-diagnostics.md` — this file (status + Evidence)

No files relocated. No T12 review edits. No T13 Work Item closure edits. No semantic redesign.

Temporary graph check script `tmp_graph_check.py` was created and deleted within the session; it is not present in the working tree.

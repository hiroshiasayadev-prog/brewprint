# PRODUCT-WORK-SPEC-005: Existing spec format migration and restructuring

- **id**: PRODUCT-WORK-SPEC-005
- **status**: done
- **date**: 2026-06-10
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-001
  - PRODUCT-WORK-SPEC-002
  - PRODUCT-WORK-SPEC-003
  - PRODUCT-WORK-SPEC-004
  - PRODUCT-WORK-SPEC-006
- **task_refs**:
  - PRODUCT-TASK-SPEC-005-01
  - PRODUCT-TASK-SPEC-005-02
  - PRODUCT-TASK-SPEC-005-03
  - PRODUCT-TASK-SPEC-005-04
  - PRODUCT-TASK-SPEC-005-05
  - PRODUCT-TASK-SPEC-005-06
  - PRODUCT-TASK-SPEC-005-07
  - PRODUCT-TASK-SPEC-005-08
  - PRODUCT-TASK-SPEC-005-09
  - PRODUCT-TASK-SPEC-005-10
  - PRODUCT-TASK-SPEC-005-11
  - PRODUCT-TASK-SPEC-005-12
  - PRODUCT-TASK-SPEC-005-13
  - PRODUCT-TASK-SPEC-005-14
  - PRODUCT-TASK-SPEC-005-15
  - PRODUCT-TASK-SPEC-005-16
  - PRODUCT-TASK-SPEC-005-17
  - PRODUCT-TASK-SPEC-005-18
  - PRODUCT-TASK-SPEC-005-19
  - PRODUCT-TASK-SPEC-005-20
  - PRODUCT-TASK-SPEC-005-21

## Summary

Migrate existing PRODUCT / DRMCP / BPDSL spec files to the accepted spec format after compatibility, authoring guide, ownership decisions, and temporary validation tooling are complete.

## Scope

| area | in scope |
|---|---|
| format migration | Add accepted H1 format, H1-adjacent metadata, required sections, and visible source/topic structures. |
| restructuring | Split or restructure specs that cannot be migrated directly. |
| ownership relocation | Execute accepted relocation plan from PRODUCT-WORK-SPEC-004. |
| alias / compatibility | Apply compatibility decisions from PRODUCT-WORK-SPEC-002. |
| temporary validation | Use PRODUCT-WORK-SPEC-006 tooling to validate target files during migration review. |
| validation hardening | Coordinate when validation severity can move from migration warning to error. |

## Non-scope

| area | owner |
|---|---|
| format design | PRODUCT-WORK-SPEC-001 |
| ID/ref compatibility design | PRODUCT-WORK-SPEC-002 |
| authoring guide design | PRODUCT-WORK-SPEC-003 |
| ownership decision | PRODUCT-WORK-SPEC-004 |
| temporary validator/tooling build | PRODUCT-WORK-SPEC-006 |
| DRMCP implementation | DRMCP-WORK-SPEC-001 / DRMCP-WORK-SPEC-002 |

## Dependencies

| dependency | reason |
|---|---|
| PRODUCT-WORK-SPEC-001 | Accepted format contract. |
| PRODUCT-WORK-SPEC-002 | Compatibility and alias behavior. |
| PRODUCT-WORK-SPEC-003 | Authoring guidance for migration edits. |
| PRODUCT-WORK-SPEC-004 | Ownership boundary and relocation plan. |
| PRODUCT-WORK-SPEC-006 | Temporary validation support for migration; this replaces any need to patch current DRMCP before migration. |

## Done condition

| item | done when |
|---|---|
| migrated specs | Target spec set follows the accepted format. |
| relocation complete | Accepted ownership relocation is executed. |
| temporary validation used | Migration batches are checked with PRODUCT-WORK-SPEC-006 tooling or an explicitly accepted equivalent. |
| validation state clean | Validation diagnostics match post-migration severity expectations. |
| stale aliases addressed | Alias / redirect mappings are present where needed and no untracked stale references remain. |
| review complete | Migration is reviewed in phases rather than as one opaque rewrite. |

## Source records

| ref | role |
|---|---|
| PRODUCT-REQ-SPEC-001 | Requirement for MCP-readable spec format and topic tree support. |
| PRODUCT-INV-SPEC-001 | Investigation evidence and migration classification basis. |
| PRODUCT-WORK-SPEC-001 | Format contract and follow-up split. |

## Evidence

### BPDSL DSL batch (PRODUCT-TASK-SPEC-005-01 .. -04) — done 2026-06-17

- Restructured `bpdsl/records/spec/` from flat layout into three peer domains: `dsl/`, `mcp/`, `views/` (mcp/views pre-existed; dsl/ is new). `bpdsl/records/spec/overview.md` `## Topics` reduced to 3 rows pointing at the three domain overviews.
- Migrated and split the 7 flat DSL files (staged at `bpdsl/old/` during the work, now removed) into 16 spec-format-compliant files under `dsl/`:
  - `nodes.md` (687 lines, Japanese) → `dsl/nodes/{overview,processing,data,application}.md`
  - `edges.md` (657 lines, Japanese) → `dsl/edges/{overview,data-flow,state-transitions,cross-edges,cross-file-refs}.md`
  - `file-types.md`, `naming.md`, `type-ref.md`, `project-layout.md`, `diagnostics.md` → 1:1 migrated under `dsl/`
  - `design-philosophy.md` moved from spec root into `dsl/`
  - `dsl/overview.md` newly authored as the dsl/ domain entry point
- All H1 titles and section headings translated to English; YAML front matter removed; H1-adjacent metadata (`id`/`status`/`date`/`parent`) added; parent chain traced clean to `spec:bpdsl.overview` for all 16 files.
- Independent Opus review (PRODUCT-TASK-SPEC-005-03) found 2 broken internal section-anchor links and 3 content judgment calls (2 worth restoring — real scope-boundary disclaimers; 1 correctly left dropped as stale/self-contradictory in the source). All resolved; cross-checked against source rather than accepted as-is.
- Final state: `validate_spec.py bpdsl/records/spec/dsl/ --strict` and `bpdsl/records/spec/overview.md --strict` both exit 0. `bpdsl/old/` removed.
- `mcp/` and `views/` remain in pre-migration format — out of scope for this batch, owned by PRODUCT-TASK-SPEC-005-05 .. -12.

### BPDSL mcp/ batch (PRODUCT-TASK-SPEC-005-05 .. -08) — done 2026-06-17

- Staged all 12 existing `mcp/` files (unchanged) to `bpdsl/old_mcp/` for diff reference, then migrated the same 12 files in place (directory layout unchanged) into spec-format-compliant form: `overview.md` (Overview), `errors.md` / `schema.md` / `versioning.md` (Reference), and 8 `tools/*.md` (Contract, `contract_class: interface`, restructured from numbered sections into `## What this is` / `## Request` / `## Response` / `## Errors`).
- Caught and corrected a false mojibake alarm: terminal/grep output had shown corrupted bytes in several source H1s (e.g. `# MCP�d�l Overview`), but re-reading with the Read tool confirmed all source files were clean, properly-encoded Japanese — the corruption was a Windows console codepage display artifact in Bash/grep, not file corruption. No content was guessed from corrupted bytes.
- Independent Opus review (PRODUCT-TASK-SPEC-005-07, spawned via Agent tool) found 5 defer-level findings, no must-fix. Cross-checked all 5 against source: restored 2 trimmed illustrative examples in `inspect.md` for fidelity, fixed 1 stale "future `analyze_impact`" reference now that `analyze_impact` is a real v1 spec, left 2 as correctly intentional (a dead `docs/TASKS.md` reference reworded; an ADR-attribution gap matching the stated front-matter-vs-body migration rule).
- Final state: `validate_spec.py bpdsl/records/spec/mcp/ --strict` exits 0 on all 12 files. `bpdsl/old_mcp/` removed.
- `views/` remains in pre-migration format — owned by PRODUCT-TASK-SPEC-005-09 .. -12.

### BPDSL views/ batch (PRODUCT-TASK-SPEC-005-09 .. -12) — done 2026-06-17

- Staged all 7 existing `views/` files (unchanged) to `bpdsl/old_views/` for diff reference, then migrated the same 7 files in place into `Contract: format` shape (`contract_class: format`, `## What this is` / `## Current contract` / `## Rules` / `## Validation rules`) plus a newly authored `views/overview.md` (Overview, 7-row `## Topics`). Total: 8 files.
- Source files were clean UTF-8 Japanese — no actual mojibake; earlier terminal/grep display artifacts (same Windows console issue as the mcp/ batch) were caught early after re-reading with the Read tool.
- The section split between `## Rules` (substantive render/behavior rules) and `## Validation rules` (edge cases, omission conditions, parser-error conditions) was modeled directly on `spec:product.concepts.spec_format.document_shape`, which is itself a `Contract: format` spec.
- `dag.md` (largest source at 992 lines) migrated as a single spec — all 6 worked render examples, all node-shape classDef hex values, all edge-kind rules preserved in full.
- Independent review (PRODUCT-TASK-SPEC-005-11): review was run via Agent tool (Opus 4.8) and findings were cross-checked in full against actual files before any fix was applied. Two of the three non-defer findings were false alarms (Opus cited non-existent text and an invented format constraint). One real omission found: `er.md` `## Rules` was missing the `render_er` generation-source statement present in the two peer files (`state-diagram.md`, `sequence-diagram.md`). Applied.
- Lesson recorded: do not spawn Agent `model: "opus"` for reviews — maps to 4.8 which hallucinates. Future review gates use the T03 pattern: give the user a ready-to-run Opus 4.7 prompt to run externally, then cross-check findings before applying.
- Final state: `validate_spec.py bpdsl/records/spec/views/ --strict` exits 0 on all 8 files. `bpdsl/old_views/` removed. `validate_spec.py bpdsl/records/spec/ --strict` exits 0 on all 37 files across `dsl/` + `mcp/` + `views/`. **Full BPDSL namespace spec-format migration complete.**
- DRMCP migration (PRODUCT-TASK-SPEC-005-13 .. -16) is tracked separately in this work item.

### Ownership boundary decision (PRODUCT-WORK-SPEC-004) — done 2026-06-17

PRODUCT-WORK-SPEC-004 closed. Accepted decision recorded in PRODUCT-TASK-SPEC-004-01. Key impact on this work item:

- PRODUCT-TASK-SPEC-005-13 done: `drmcp/old/` staged, `schema/` + `tools/` dirs created.
- PRODUCT-TASK-SPEC-005-14..-16 proceed as **format-only** migration. Four sections flagged as deferred relocation candidates (per PRODUCT-INV-SPEC-004); PRODUCT-TASK-SPEC-005-15 review must record these explicitly.
- **Phase 2 relocation batch** (future tasks under this work item, TBD after -16 completes): four DRMCP-side sections will relocate to `namespace-model/index.md` / `repository-layout/index.md` per the accepted plan. Task numbers and exact file targets set at that time.
- PRODUCT-WORK-SPEC-009 scope updated to 9 files (added `namespace-model/index.md`); that work item is independent of DRMCP migration and can run in parallel. PRODUCT-WORK-SPEC-009 completed before Phase 2 began.

### DRMCP Phase 2 relocation batch (PRODUCT-TASK-SPEC-005-17 .. -21) — done 2026-06-22

Relocated 4 deferred DRMCP sections to PRODUCT namespace. All deferred relocation notes removed from DRMCP files.

| step | source (DRMCP) | destination (PRODUCT) | notes |
|---|---|---|---|
| -17 | `namespace-scanning.md` core content | `namespace-model/v1-namespace-algorithm.md` (new) | namespace_prefix derivation formula, kind-level prefix table, multi-root scan |
| -17 | `schema/id-normalization.md` | `namespace-model/v1-id-grammar.md` (new) | public ID and bare ID grammar tables |
| -18 | `schema/discovery.md` path-pattern table | `repository-layout/record-discovery-paths.md` (new) | DRMCP kind-filter stays in `schema/discovery.md` |
| -18 | `repository-layout/index.md` | — | Format-migrated in place (was not in WORK-009 scope; handled as -18 prerequisite) |
| -19 | `resolver.md` duplicated resolver semantics | — | Trimmed to pointer → `spec:product.concepts.traceability.resolve_and_validation`; DRMCP-specific sections retained |
| -19 | drift guards | `traceability/resolve-and-validation.md` | Already present from WORK-009; no edits needed |

ADR discovery pattern corrected during -20 review: `<records_root>/adr/*.md` → `<records_root>/adr/<namespace_prefix>ADR-*.md` in `record-discovery-paths.md` (legacy artifact from pre-migration DRMCP spec, not the intended format).

Independent Opus 4.7 review (PRODUCT-TASK-SPEC-005-20): **PASS** — 0 must-fix items, 2 defer items (stale `docs/spec/project-layout.md` path in `repository-layout/index.md`; task-spec wording on §MVP exclusions — implementation correct).

Final state: `validate_spec.py product/records/spec/concepts/ drmcp/records/spec/design-records-mcp/ --strict` exits 0 on all 62 files. No deferred relocation notes remain in DRMCP. **DRMCP Phase 2 relocation complete.**

### DRMCP design-records-mcp/ batch (PRODUCT-TASK-SPEC-005-13 .. -16) — done 2026-06-17

- Staged all 3 existing DRMCP spec files to `drmcp/old/` (PRODUCT-TASK-SPEC-005-13), created `schema/` and `tools/` subdirectories.
- Authored 30 spec-format-compliant files under `drmcp/records/spec/design-records-mcp/` from the 3 source files (PRODUCT-TASK-SPEC-005-14):
  - `drmcp/old/overview.md` (220 lines, Japanese) → `overview.md` + 4 flat Reference files (`responsibility-boundary`, `resolver`, `namespace-scanning`, `mvp-scope`)
  - `drmcp/old/schema.md` (835 lines, Japanese) → `schema/overview.md` + 9 Reference files
  - `drmcp/old/tools.md` (1944 lines, Japanese) → `tools/overview.md` + `tools/authoring-transaction-model.md` (Reference) + 13 Contract (`contract_class: interface`) files
  - All H1 titles, H2 section headings, and table headers translated to English; YAML front matter removed; H1-adjacent metadata (`id`/`status`/`date`/`parent`) added; parent chain traces clean to `spec:drmcp.design_records_mcp.overview` for all children.
  - 3 `REFERENCE_NO_TABLE` errors fixed during authoring: `resolver.md`, `schema/authoring-guidance-source.md`, `schema/metadata-grammar.md` — resolved by adding a Markdown table to `## Current contract` in each file.
- Independent Opus 4.7 review (PRODUCT-TASK-SPEC-005-15): **PASS** — 0 must-fix items, 2 defer items. Cross-checked all findings against actual files. One defer item confirmed accurate: cross-ref wording in `tools/validate-records.md:107` is imprecise for 5 authoring-only diagnostic categories not defined in `schema/diagnostics.md`. Deferred to a future fixup.
- `drmcp/old/` removed. `validate_spec.py drmcp/records/spec/design-records-mcp/ --strict` exits 0 on all 30 files.
- **Deferred relocation candidates** (PRODUCT-INV-SPEC-004 / PRODUCT-WORK-SPEC-004 Phase 2 — open): the following 4 files contain PRODUCT-owned semantics migrated in place as format-only. Each file carries a `> **Deferred relocation note**`. Relocation target is `spec:product.concepts.namespace_model` / `spec:product.concepts.repository_layout`.

  | file | PRODUCT-owned semantics |
  |---|---|
  | `resolver.md` | namespace_prefix derivation, public/bare ID grammar |
  | `namespace-scanning.md` | namespace_prefix derivation algorithm, multi-root scan, public/bare ID grammar |
  | `schema/id-normalization.md` | public ID and bare ID grammar |
  | `schema/discovery.md` | discovery path pattern conventions (partial; DRMCP kind-filter stays) |

  Phase 2 relocation tasks: PRODUCT-TASK-SPEC-005-17 .. -21 — **done 2026-06-22** (see Phase 2 entry below).

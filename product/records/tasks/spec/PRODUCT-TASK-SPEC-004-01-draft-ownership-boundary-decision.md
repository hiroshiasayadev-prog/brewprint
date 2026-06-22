# PRODUCT-TASK-SPEC-004-01: Draft ownership boundary decision and relocation plan

- **id**: PRODUCT-TASK-SPEC-004-01
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-004
- **estimate**: 1d
- **depends_on**:
- **outputs**:
  - Draft ownership boundary decision (inline in Evidence, or as a new ADR if the user prefers a formal decision artifact)
  - Draft relocation plan: source file/section → destination → move order

## Goal

Synthesize PRODUCT-INV-SPEC-002 and PRODUCT-INV-SPEC-004's classification tables into a single accepted ownership boundary decision, plus a concrete relocation plan with source/destination/order — without executing any relocation.

## Work

| area | required work |
|---|---|
| merge classifications | Combine PRODUCT-INV-SPEC-002 (traceability / project-artifact-model, 8 files) and PRODUCT-INV-SPEC-004 (namespace-model vs. DRMCP, 2 files) section-level classification tables into one decision document. |
| resolve open ambiguities | Decide the ambiguities both investigations left open — notably: where in `namespace-model/index.md` the relocated v1 ID-grammar/namespace-prefix content lands (new section vs. new sibling file); whether `resolve-and-validation.md`'s hybrid sections need drift guards. |
| relocation plan | For each relocation candidate from both investigations, specify: source file/section, destination file/section, and move order (what must move before what). |
| WORK-SPEC-009 scope note | Recommend whether `namespace-model/index.md` should be added as a 9th target file to PRODUCT-WORK-SPEC-009's format-only migration scope. Do not edit PRODUCT-WORK-SPEC-009 directly in this task — that edit belongs to PRODUCT-TASK-SPEC-004-03 once the decision is accepted. |
| dependency update plan | List records/references that must be updated once relocation executes (e.g. `overview.md`'s pointer to `resolve-and-validation.md` instead of restating resolver semantics). |

## Done condition

| item | done when |
|---|---|
| decision drafted | A single ownership boundary decision covering both investigations exists. |
| relocation plan drafted | Source/destination/order table exists for every relocation candidate from both investigations. |
| ambiguities resolved | Every open ambiguity from PRODUCT-INV-SPEC-002 and PRODUCT-INV-SPEC-004 has a stated resolution. |
| WORK-SPEC-009 recommendation made | A clear recommendation (add file / don't add) is recorded, pending PRODUCT-TASK-SPEC-004-03. |

## Verification

- Cross-check the draft against both investigations' "Relocation candidates" tables — no candidate should be silently dropped.
- Confirm the relocation plan does not contradict PRODUCT-INV-SPEC-002's recommendation ("no whole file relocated to DRMCP before migration").

## Evidence

### Decision summary

1. **No whole PRODUCT file relocates to DRMCP.** PRODUCT retains full ownership of: project-artifact-model semantics, traceability/canonical-reference semantics and vocabulary, namespace/ID-grammar model, repository layout conventions.
2. **Four DRMCP spec sections are misplaced and will relocate to PRODUCT**, in a future relocation batch (Phase 2) after Phase 1 format migrations complete. DRMCP's remaining content (tool contracts, response shapes, diagnostic detail, scope decisions, responsibility boundaries) stays DRMCP-owned.
3. **`namespace-model/index.md` is added to PRODUCT-WORK-SPEC-009's format-only migration scope** as a 9th target file. It is currently incomplete (missing the v1 algorithm it will receive from DRMCP relocation) and pre-migration format. Format migration must complete before relocation content can be integrated cleanly.
4. **Two hybrid sections in `traceability/resolve-and-validation.md` get explicit drift guards** — documented inline once format-migrated — stating which clause is a PRODUCT semantic rule vs. a DRMCP API vocabulary item.

---

### Merged classification — all reviewed sections

Sources: PRODUCT-INV-SPEC-002 (8 traceability/artifact-model files) + PRODUCT-INV-SPEC-004 (namespace-model/index.md vs. DRMCP namespace/ID-grammar content).

| file / section | final ownership | action |
|---|---|---|
| `project-artifact-model/index.md` — all sections | PRODUCT | No relocation. Format-migrate under PRODUCT-WORK-SPEC-009. Stale `docs/...` path cleanup during migration. |
| `traceability/index.md` — all sections | PRODUCT | No relocation. Migrate as `Overview` + `## Topics` (not a pure `Index` — see ambiguity resolution below). |
| `traceability/artifact-refs.md` — all sections | PRODUCT | No relocation. DRMCP implements lookup; PRODUCT owns vocabulary. |
| `traceability/coverage-mapping.md` — all sections | PRODUCT | No relocation. Deferred realization mapping is cross-app governance. |
| `traceability/metadata-schema.md` — trace metadata / semantic_refs / sections / front matter | PRODUCT | No relocation. Do not add concrete DRMCP parser/response schema here. |
| `traceability/metadata-schema.md` — workflow and investigation reference metadata | hybrid, PRODUCT-led | No relocation. Add drift guard: PRODUCT owns concrete bidirectional integrity rule statements; DRMCP owns parser/response/diagnostic details. |
| `traceability/metadata-schema.md` — validation responsibility | hybrid, PRODUCT-led | No relocation. Add drift guard: PRODUCT owns invalid conditions; DRMCP owns diagnostic category names/JSON shape. |
| `traceability/out-of-scope.md` — all sections | PRODUCT | No relocation. |
| `traceability/resolve-and-validation.md` — resolve / resolver input / lookup sources / section anchor lookup | hybrid, PRODUCT-led | No relocation. Add drift guard: PRODUCT owns canonical input semantics and tool behavior descriptions expressed in PRODUCT metadata; DRMCP may expose them as tool API. |
| `traceability/resolve-and-validation.md` — duplicate detection / unresolved refs / declared relation integrity validation | hybrid, PRODUCT-led, drift-sensitive | No relocation. Add drift guard: PRODUCT owns invalid conditions and scope boundary; DRMCP owns diagnostic category names, JSON response shape, tool response vocabulary. Do not let DRMCP response vocabulary accumulate in these clauses. |
| `traceability/resolve-and-validation.md` — resolver output | DRMCP-owned pointer | No relocation. Current section correctly delegates to DRMCP tools spec. Preserve. |
| `traceability/resolve-and-validation.md` — MCP writer contract placeholder | DRMCP-owned future pointer | No relocation. Do not expand into PRODUCT-owned API schema. |
| `traceability/semantic-ref.md` — all sections | PRODUCT | No relocation. |
| `namespace-model/index.md` — existing sections (app/domain namespace definitions, domain catalog, subdomain model, v2 grammar, existing artifact ownership) | PRODUCT | No relocation. Add format migration (WORK-SPEC-009). Will receive new content from DRMCP relocation in Phase 2. |
| `drmcp/overview.md` §Record scanning と namespace prefix — namespace_prefix derivation formula, multi-root scan default, kind-別 prefix application table | PRODUCT (misplaced in DRMCP) | **Relocate Phase 2**: move core content to `namespace-model/index.md` new section (see relocation plan). Replace with pointer in DRMCP. |
| `drmcp/schema.md` §ID normalization model — Public ID / Bare ID grammar tables | PRODUCT (misplaced in DRMCP) | **Relocate Phase 2**: move to `namespace-model/index.md` as v1 record ID grammar section. Replace with pointer in DRMCP. |
| `drmcp/schema.md` §Discovery and index inclusion — discovery path pattern per kind | PRODUCT (misplaced in DRMCP) | **Relocate Phase 2** (path pattern only): move path-pattern convention to `repository-layout/index.md` new section. DRMCP kind-inclusion filter (e.g. requiring `design_record` front matter) stays in DRMCP. |
| `drmcp/overview.md` §Resolver responsibility | hybrid, duplicate-content risk | **Trim Phase 2**: reduce to a pointer into `traceability/resolve-and-validation.md`. No content moves to PRODUCT — PRODUCT already owns it via that file; DRMCP is duplicating it. |
| `drmcp/overview.md` §既存brewprint MCPとの責務境界 / §filesystem tool との責務境界 | DRMCP | No relocation. Should cross-ref `project-artifact-model/index.md` boundary table instead of stating the rule independently, but ownership stays DRMCP. This is a format-migration editorial note for PRODUCT-TASK-SPEC-005-14, not a relocation candidate. |
| `drmcp/overview.md` §目的, §対象record, §MVP tool set, §MVP外 | DRMCP | No relocation. Legitimate DRMCP scope decisions. |
| All `drmcp/tools.md` tool contracts (13 tools) | DRMCP | No relocation. These are the genuine DRMCP implementation contracts. |
| All `drmcp/schema.md` non-ID-grammar sections (authoring guidance source model, authoring transaction schema concepts, field definitions, record model, diagnostic categories, etc.) | DRMCP | No relocation. These define concrete implementation-level data shapes that DRMCP owns. |

---

### Ambiguity resolutions

| ambiguity | resolution |
|---|---|
| PRODUCT semantic validation conditions vs. DRMCP diagnostic categories | PRODUCT owns invalid conditions/scope-boundary clauses; DRMCP owns diagnostic category names, JSON shape, tool response vocabulary. Drift guards implement this. |
| `spec:` grammar uses hyphens; WORK-SPEC-001 path-derived IDs use underscores | Out of WORK-SPEC-004's scope. Carry forward to PRODUCT-WORK-SPEC-002 as an open compatibility question. |
| `resolve-and-validation.md` — migrate as `Contract` or `Reference`/`Concept` | Migrate as PRODUCT-owned **`Reference`** kind. The "Contract" label would mislead future agents into treating it as a DRMCP API contract. The file defines semantic resolution/validation boundaries, not a tool request/response shape. |
| How much workflow relation metadata belongs in PRODUCT traceability vs. workflow authoring guides | PRODUCT traceability owns the semantic validity rule (what constitutes a valid/invalid declared relation). Authoring guides own the how-to narrative for authors. No duplication. PRODUCT-WORK-SPEC-003 authoring guides should not restate traceability semantic rules — reference them instead. |
| Whether `traceability/index.md` should migrate as `Overview`+`## Topics` or pure `Index` | **`Overview`+`## Topics`**: splitting a substantive overview from its topic table would add unnecessary churn and the current file's content warrants a real `## What this is` section. |
| Stale `scope:` front-matter paths pointing at `docs/spec/...` | Migration cleanup in PRODUCT-WORK-SPEC-009, not ownership change. |
| Legacy DRMCP pointer destinations not verified | Carried to PRODUCT-TASK-SPEC-004-03: dependency-update-plan step will enumerate PRODUCT files that currently point to legacy `docs/spec/design-records-mcp/...` and list their correct `drmcp/records/spec/design-records-mcp/...` targets. |
| Where v1 namespace_prefix/ID-grammar content lands in `namespace-model/index.md` — new section vs. sibling file | **New sections in existing `namespace-model/index.md`**: two additions — `## v1 namespace resolution algorithm` (namespace_prefix formula + kind-別 prefix table + multi-root scan behavior) and `## v1 record ID grammar` (public ID / bare ID tables). The content volume (~30 lines each) does not justify new sibling files. |
| Whether `namespace-model/index.md` should be added to PRODUCT-WORK-SPEC-009's scope | **Yes — add as 9th target file.** It is pre-migration format and will receive new content from DRMCP relocation. Format migration must precede content integration. |

---

### Relocation plan

No relocation executes until Phase 1 format migrations complete (PRODUCT-WORK-SPEC-009 for PRODUCT files, PRODUCT-TASK-SPEC-005-13..-16 for DRMCP files).

**Phase 1 — format migration only (no content moves):**

| work item | target files |
|---|---|
| PRODUCT-WORK-SPEC-009 (updated scope) | 8 original PRODUCT-INV-SPEC-002 files + `namespace-model/index.md` (9th) |
| PRODUCT-TASK-SPEC-005-13..-16 | `drmcp/records/spec/design-records-mcp/` 30 output files; deferred relocation candidates flagged in -15 evidence |

**Phase 2 — relocation execution (new batch under PRODUCT-WORK-SPEC-005):**

| step | source | destination | notes |
|---|---|---|---|
| 1 | `drmcp/records/spec/design-records-mcp/namespace-scanning.md` §core content | `namespace-model/index.md` new `## v1 namespace resolution algorithm` section | Namespace_prefix formula + kind-別 prefix table + multi-root scan default. Replace DRMCP file with a one-paragraph pointer to `spec:product.concepts.namespace_model`. |
| 2 | `drmcp/records/spec/design-records-mcp/schema/id-normalization.md` | `namespace-model/index.md` new `## v1 record ID grammar` section | Public ID / bare ID grammar tables. Replace DRMCP file with pointer. |
| 3 | `drmcp/records/spec/design-records-mcp/schema/discovery.md` path-pattern content | `repository-layout/index.md` new `## Record discovery paths` section | Kind → `<records_root>/.../*.md` pattern convention. DRMCP kind-inclusion filter stays in remaining `discovery.md` content. |
| 4 | `drmcp/records/spec/design-records-mcp/resolver.md` | Trim to pointer → `traceability/resolve-and-validation.md` | No content moves to PRODUCT — resolver semantics already live there. DRMCP was duplicating. |
| 5 | Drift guards for `traceability/resolve-and-validation.md` hybrid sections | In-file editorial notes (not relocation) | Added during Phase 2 or during PRODUCT-WORK-SPEC-009 migration pass, whichever is more natural. |

Phase 2 should be proposed as a follow-on batch under PRODUCT-WORK-SPEC-005 (task numbers TBD at TASK-004-03 time, after Phase 1 completes and the exact file shapes are known).

---

### WORK-SPEC-009 scope recommendation

Add `namespace-model/index.md` as a 9th target file to PRODUCT-WORK-SPEC-009's `Impact Scope`. Same boundary as the existing 8 files: format migration only, no ownership relocation, no content edits beyond stale-path cleanup. PRODUCT-TASK-SPEC-004-03 executes this update.

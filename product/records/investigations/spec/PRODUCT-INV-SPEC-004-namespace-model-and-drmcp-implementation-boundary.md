# PRODUCT-INV-SPEC-004: Namespace model and DRMCP implementation boundary

- **id**: PRODUCT-INV-SPEC-004
- **status**: concluded
- **date**: 2026-06-17
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-005

## Question

`product/records/spec/concepts/namespace-model/index.md` defines the app-namespace / domain-namespace model, including per-app architecture sketches. DRMCP's own spec (`drmcp/records/spec/design-records-mcp/overview.md`, `schema.md`) independently defines a namespace-prefix derivation algorithm, multi-root scan behavior, and a public/bare record-ID grammar — none of which is DRMCP-specific business logic; any namespace-aware tool in the repo would need the identical rules. Which sections are PRODUCT-owned cross-app namespace/identity semantics that DRMCP merely implements, versus genuinely DRMCP-owned tool contract, and is `namespace-model/index.md` itself missing content as a result?

This question was not covered by PRODUCT-INV-SPEC-002, whose reviewed-file set was limited to `traceability/**` and `project-artifact-model/index.md` and did not include `namespace-model/index.md`.

## Scope

| area | in scope |
|---|---|
| namespace model | Inspect `product/records/spec/concepts/namespace-model/index.md` in full. |
| DRMCP namespace/ID content | Inspect `drmcp/records/spec/design-records-mcp/overview.md` §Record scanning と namespace prefix, §Resolver responsibility, and `schema.md` §ID normalization model, §Discovery and index inclusion. |
| ownership classification | Classify each section as PRODUCT-owned semantics, DRMCP-owned tool contract, or hybrid, using the same method as PRODUCT-INV-SPEC-002. |
| completeness gap | Identify what `namespace-model/index.md` would need to add to become the canonical home for content currently only described inside DRMCP's spec. |
| relocation candidates | Identify files/sections that may need relocation before DRMCP's spec format migration (PRODUCT-TASK-SPEC-005-13..-16) is finalized. |

## Non-scope

| area | owner |
|---|---|
| actual relocation | PRODUCT-WORK-SPEC-004 / PRODUCT-WORK-SPEC-005 |
| DRMCP spec format migration | PRODUCT-TASK-SPEC-005-13..-16 |
| v2 artifact ID grammar redesign | Out of scope — this investigation concerns the *current* (v1) namespace_prefix / ID grammar already implemented and described in DRMCP's spec, not the not-yet-adopted v2 grammar `namespace-model/index.md` already documents. |
| DRMCP implementation correctness | DRMCP work items |

## Expected evidence

| evidence | purpose |
|---|---|
| section-level classification table | Shows PRODUCT / DRMCP / hybrid ownership for the reviewed sections. |
| completeness gap list | What `namespace-model/index.md` is missing. |
| relocation candidate list | Feeds PRODUCT-WORK-SPEC-004. |
| recommendation | Tells PRODUCT-WORK-SPEC-005 / PRODUCT-TASK-SPEC-005-13..-16 whether to proceed with format-only migration now. |

## Done condition

| item | done when |
|---|---|
| classification complete | Target sections are classified. |
| completeness gap identified | Missing content in `namespace-model/index.md` is listed. |
| relocation candidates identified | Candidate sections and reasons are listed. |
| decision handoff ready | PRODUCT-WORK-SPEC-004 can fold this into its ownership boundary decision alongside PRODUCT-INV-SPEC-002. |

## Source records

| ref | role |
|---|---|
| PRODUCT-INV-SPEC-002 | Established the classification methodology and recommendation pattern this investigation extends; its reviewed-file set did not include `namespace-model/index.md`. |
| PRODUCT-WORK-SPEC-004 | Downstream ownership decision work item; currently `not_started`, empty `task_refs`. |
| PRODUCT-WORK-SPEC-005 | Where this question surfaced, while planning PRODUCT-TASK-SPEC-005-13..-16 (DRMCP spec format migration). |
| PRODUCT-WORK-SPEC-009 | Precedent: format-only migration of the 8 PRODUCT-INV-SPEC-002 target files, explicitly deferring relocation until PRODUCT-WORK-SPEC-004 decides. |

## Evidence

### Files reviewed

| file | reviewed scope |
|---|---|
| `product/records/spec/concepts/namespace-model/index.md` | Full document. |
| `drmcp/records/spec/design-records-mcp/overview.md` | Full document (220 lines). |
| `drmcp/records/spec/design-records-mcp/schema.md` | §ID normalization model (L436–477), §Discovery and index inclusion (L414–435). |
| `PRODUCT-INV-SPEC-002` | Methodology and recommendation pattern. |

### Classification summary

| target | ownership | reason | relocation impact |
|---|---|---|---|
| `overview.md` §Record scanning と namespace prefix — `namespace_prefix = ToUpper(appNamespaceDir) + "-"`, multi-root scan default, kind-別 prefix application table | PRODUCT | Generic cross-app namespace resolution rule. Any namespace-aware tool (a future BPDSL tool, DRUI, etc.) needs the identical algorithm — it is not DRMCP business logic. `namespace-model/index.md` defines the namespace *concept* but never states this *algorithm*. | Move to `namespace-model/index.md` (or a new sibling reference file under that concept). DRMCP's spec should reference it, not restate it. |
| `schema.md` §ID normalization model (Public ID / Bare ID grammar) | PRODUCT | Public/bare ID grammar is defined per record *kind* across the whole repo, not per DRMCP tool. Parallels PRODUCT-INV-SPEC-002's classification of `artifact-refs.md` as "shared reference vocabulary; DRMCP implements lookup." | Move to `namespace-model/index.md`'s v2-grammar sibling area, but as the *current* (v1) grammar — `namespace-model/index.md` today only documents the not-yet-adopted v2 grammar, so this is new content, not a straight cut-paste. |
| `schema.md` §Discovery and index inclusion (discovery-path table: kind → `<records_root>/.../*.md` pattern) | hybrid | The path *pattern* itself overlaps with `repository-layout/index.md`'s "records/ internal structure" section (kind-first directory layout). Which record kinds DRMCP actually chooses to index (e.g. requiring `design_record` front matter for specs) is a legitimate DRMCP tool-scope decision. | Split: path pattern → PRODUCT (`repository-layout` or `namespace-model`); DRMCP-specific inclusion filtering stays in DRMCP. |
| `overview.md` §Resolver responsibility | hybrid, PRODUCT-led | Restates resolver input/output semantics already owned by `traceability/resolve-and-validation.md` (classified hybrid, PRODUCT-led in PRODUCT-INV-SPEC-002 — "PRODUCT owns canonical input semantics; DRMCP may expose them through tool API"). This is a duplicate-content risk, not just a placement question: two specs currently describe the same resolver semantics independently. | Trim to a pointer to `traceability/resolve-and-validation.md`; keep only DRMCP's concrete tool-API exposure of it. |
| `overview.md` §既存brewprint MCPとの責務境界 / §filesystem tool との責務境界 | DRMCP | DRMCP stating its own boundary against other concrete tools is legitimate tool-contract content, consistent with PRODUCT-INV-SPEC-002's classification of `resolve-and-validation.md`'s "resolver output" section as a DRMCP-owned tool-contract pointer. | No relocation. Should reference `project-artifact-model/index.md`'s cross-app boundary table rather than restate it independently, but this is a cross-reference cleanup, not an ownership move. |
| `overview.md` §目的, §MVP tool set, §MVP外 | DRMCP | Legitimate DRMCP-specific scope decisions (which tools this specific MCP server ships). | No relocation. |

### Completeness gap in `namespace-model/index.md`

| gap | detail |
|---|---|
| no current (v1) namespace_prefix algorithm | The spec documents only a future v2 ID grammar (`<APP_NAMESPACE>-<KIND>-<DOMAIN>-<SEQUENCE>`, not yet adopted). The algorithm actually implemented today (`namespace_prefix = ToUpper(appNamespaceDir) + "-"`, producing e.g. `V01-`, `DRMCP-`) has no PRODUCT-owned home — it currently exists only as prose inside DRMCP's `overview.md`. |
| no multi-root scan behavior | The "scan all `*/records/` directories, derive prefix per directory, build a unified index" behavior is a cross-app concept (it is what makes `PRODUCT` / `DRMCP` / `BPDSL` namespaces interoperate under one query surface) but is currently described only as a DRMCP implementation detail. |
| no public/bare ID grammar | Same gap — these two ID layers are referenced by `artifact-refs.md` (PRODUCT-owned, per PRODUCT-INV-SPEC-002) but the grammar itself is defined only in DRMCP's `schema.md`. |

### Relocation candidates

| candidate | recommended action | reason | target owner |
|---|---|---|---|
| `overview.md` §Record scanning と namespace prefix | Relocate core algorithm/table to `namespace-model/index.md`; leave a pointer in DRMCP's spec. | Cross-app namespace resolution rule, not DRMCP business logic. | PRODUCT |
| `schema.md` §ID normalization model | Relocate to `namespace-model/index.md` as the current (v1) ID grammar, alongside the existing (future v2) grammar section. | Shared identity grammar across all record kinds and namespaces. | PRODUCT |
| `schema.md` §Discovery and index inclusion | Split: path-pattern convention → PRODUCT (`repository-layout` or `namespace-model`); kind-inclusion filtering → stays DRMCP. | Hybrid: structural convention vs. tool-specific scope choice. | PRODUCT for pattern; DRMCP for filter |
| `overview.md` §Resolver responsibility | Trim to a pointer into `traceability/resolve-and-validation.md`. | Duplicate-content risk — same semantics already PRODUCT-owned elsewhere. | PRODUCT (existing owner) |

### Ambiguities and required decisions

| ambiguity | impact | owner to decide |
|---|---|---|
| Whether the relocated v1 ID-grammar/namespace-prefix content becomes a new section in `namespace-model/index.md` or a new sibling file under that concept. | Affects `namespace-model/index.md`'s eventual `## Topics`/structure once it migrates. | PRODUCT-WORK-SPEC-004 |
| Whether `namespace-model/index.md` should be added as a 9th target file to PRODUCT-WORK-SPEC-009's format-only migration scope, given it currently covers only the 8 PRODUCT-INV-SPEC-002 files. | If not added, `namespace-model/index.md` stays pre-migration format while DRMCP/BPDSL specs around it move to the new format. | PRODUCT-WORK-SPEC-004, with a PRODUCT-WORK-SPEC-009 scope update as follow-up |
| Whether PRODUCT-TASK-SPEC-005-13..-16 (DRMCP spec format migration, already drafted) should proceed now against the current (misplaced) content, or wait for relocation. | Affects whether DRMCP's migrated `namespace-scanning.md` / `schema/id-normalization.md` files become throwaway work. | Already decided by the user for this round: proceed with format migration now, flag relocation candidates in the PRODUCT-TASK-SPEC-005-15 review gate, defer actual relocation to PRODUCT-WORK-SPEC-004's accepted plan. |

### Migration risk notes

| risk | mitigation |
|---|---|
| Migrating DRMCP's `overview.md`/`schema.md` to the new format without flagging relocation candidates would bury this finding again — exactly what happened the first time (PRODUCT-INV-SPEC-002 didn't review this file). | PRODUCT-TASK-SPEC-005-15 (DRMCP Opus review) explicitly classifies the same sections found here as deferred relocation candidates in its own evidence, cross-referencing this investigation. |
| Treating this as a reason to block DRMCP's format migration entirely would repeat the WORK-SPEC-004 stall (the work item already sat `not_started` while WORK-SPEC-005 executed one full batch around it). | Format migration and ownership relocation are independent axes — formatting now does not foreclose relocation later, and produces cleaner source material for the eventual move. |
| `namespace-model/index.md` is outside PRODUCT-WORK-SPEC-009's current 8-file scope; it could be silently skipped if nobody updates that scope. | Recorded explicitly as an ambiguity above and as a recommendation below; PRODUCT-WORK-SPEC-004's task flow should include updating PRODUCT-WORK-SPEC-009's `Impact Scope`. |

### Recommendation

`namespace-model/index.md` is the correct long-term owner of: the namespace_prefix derivation algorithm, multi-root scan behavior, and public/bare record-ID grammar — all currently described only inside DRMCP's spec. This mirrors PRODUCT-INV-SPEC-002's existing pattern (PRODUCT owns identity/semantics, DRMCP owns concrete tool/schema implementation) extended to a file PRODUCT-INV-SPEC-002 never reviewed.

PRODUCT-WORK-SPEC-004 should treat this investigation as a second input alongside PRODUCT-INV-SPEC-002, and its task flow should:

1. Fold this classification into one accepted ownership boundary decision and relocation plan (covering both investigations).
2. Add `namespace-model/index.md` to PRODUCT-WORK-SPEC-009's format-only migration scope (or an equivalent follow-up), since it is materially incomplete in its pre-migration form and now has new content destined for it.
3. Produce the migration handoff PRODUCT-WORK-SPEC-005 needs, so PRODUCT-TASK-SPEC-005-13..-16's flagged sections have a concrete destination once relocation actually executes.

PRODUCT-TASK-SPEC-005-13..-16 should proceed now as format-only migration against current (misplaced) content — see the ambiguities table above for why this was the user's explicit call for this round, not unilaterally decided here.

# PRODUCT-TASK-SPEC-011-07: Create investigation authoring guide

- **id**: PRODUCT-TASK-SPEC-011-07
- **status**: done
- **date**: 2026-06-23
- **work_item**: PRODUCT-WORK-SPEC-011
- **source_requirement**: PRODUCT-REQ-SPEC-002
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-011-02
  - PRODUCT-TASK-SPEC-011-03
- **outputs**:
  - `product/records/spec/concepts/authoring-standards/investigation-authoring.md`
  - `product/records/spec/concepts/authoring-standards/index.md`

## Goal

Create the canonical investigation authoring guide under PRODUCT authoring standards.

## Work

- Use the common per-artifact guide structure.
- Define namespace-aware investigation ID grammar and file layout.
- Define canonical English body headings.
- Define investigation metadata meaning and requiredness for create, partial update, and persisted state.
- Define investigation status lifecycle.
- Define evidence, uncertainty, option comparison, conclusion, and follow-up candidate rules.
- Distinguish investigation findings from adopted ADR decisions and current specs.
- Reference `artifact_boundary` for cross-artifact selection.
- Define author-facing exact-ID and `new` forms.
- Exclude concrete DRMCP request, response, and diagnostic schemas.
- Update the authoring-standards Index.

## Done condition

| item | done when |
|---|---|
| Guide published | `investigation-authoring.md` exists as a Reference spec. |
| Common shape applied | The guide follows PRODUCT-WORK-SPEC-011 section structure. |
| English headings canonical | All prescribed investigation headings use English. |
| Metadata states separated | Create, partial update, and persisted requirements are explicit. |
| Decision boundary clear | Findings, uncertainty, options, decisions, specs, and follow-up artifacts are distinguished. |
| Index updated | The guide appears with its canonical ref and current summary. |

## Verification

- Confirm the guide uses abstract v2 IDs as primary forms.
- Confirm investigation headings and table headers use English.
- Confirm evidence and uncertainty remain separate from adopted decisions.
- Confirm follow-up candidates do not imply accepted requirements or work.
- Confirm no current DRMCP operating status appears.

## Evidence

- `product/records/spec/concepts/authoring-standards/investigation-authoring.md` created as a Reference spec; revised per review findings R1–R8.
- Common guide shape from PRODUCT-WORK-SPEC-011 applied: all sections present in order.
- All prescribed headings use English. Japanese v01 body headings translated: 調査スコープ → Investigation scope, 非スコープ → Out of scope, 背景 → Background, 調査したもの → What was investigated, 調査項目ごとの確認結果 → Findings, 横断的な観測事実 → Cross-cutting observations, 後続判断に渡す候補 → Follow-up judgment candidates, 推奨案 → Recommendation, 後続 artifact 候補 → Follow-up artifact candidates, 未確定点 → Open questions.
- Investigation follows ADR pattern: no explicit `id` field in bullet metadata. Public ID appears in H1 and file name only. Stated explicitly in Metadata schema introduction.
- Metadata schema distinguishes required fields (status, date, trigger, scope, non_scope, source_refs, follow_up_candidates) from optional fields (supersedes, related_*, follow_up_results) in separate tables with create input / partial update / persisted columns.
- `trigger`, `scope`, `non_scope` are investigation-specific required fields with no analog in other guides. `scope` and `non_scope` described as short inline summaries that mirror the corresponding body sections.
- `follow_up_candidates` required at create; empty list allowed; empty-list representation used when no candidate exists (no literal `none`). `follow_up_candidates` does not accept empty list items.
- `follow_up_results` create input: optional. Defined as canonical refs for artifacts actually created or updated as follow-up results (not "adopted"); may be supplied at create for migrated or already-concluded investigations. No restriction to post-conclusion population.
- `supersedes` partial update: optional; correction only. Defined as a list of investigation IDs (not a singular ID). Correction allowed when lineage metadata was omitted or recorded incorrectly.
- Status lifecycle: `investigating` / `concluded` / `superseded`. Explicit rule: `concluded` does not mean follow-up candidates are adopted. `proposed` excluded to avoid conflict with ADR terminology.
- Investigation conclusion-readiness rule (not "done gate"): `## Investigation scope` and `## Findings` must contain substantive content when `status` is `concluded`; `TBD` does not satisfy either condition. No DRMCP enforcement language in authoring guide.
- Supersession lineage: canonical lineage established only via `supersedes` metadata, not via `## Background`. `## Background` may explain reason and history. This is not a bidirectional workflow relation; no DRMCP reciprocal-update statement.
- `## Follow-up artifact candidates` content broadened: "Proposed follow-up artifacts or artifact updates." Supported candidates include ADR, spec, investigation, requirement, work item; `TASK-*` excluded from investigation metadata canonical references.
- Decision boundary updated: acting on a recommendation requires an ADR, spec, investigation, requirement, or work item — not an investigation update.
- `related_internal_design` and `related_coverage` described as parser-recognized optional auxiliary identifiers with no active canonical resolution or validation semantics defined by this guide.
- Guide cites `spec:product.concepts.authoring_standards.artifact_boundary`; does not duplicate the ownership matrix.
- No DRMCP operating status recorded.
- `product/records/spec/concepts/authoring-standards/index.md` updated with new entry.

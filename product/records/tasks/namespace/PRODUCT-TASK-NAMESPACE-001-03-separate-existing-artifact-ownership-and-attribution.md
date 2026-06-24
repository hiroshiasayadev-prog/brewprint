# PRODUCT-TASK-NAMESPACE-001-03: Separate historical ownership, issued identity, effective attribution, and new-artifact ownership

- **id**: PRODUCT-TASK-NAMESPACE-001-03
- **status**: completed
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-NAMESPACE-001
- **source_requirement**: V01-REQ-PRODUCT-001
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-NAMESPACE-001-01
- **outputs**:
  - `product/records/spec/concepts/namespace-model/existing-artifacts.md` (updated)

## Goal

Rewrite `existing-artifacts.md` to cover exactly three concerns — V01-ADR-096 historical ownership decision, effective attribution for navigation and grouping, and new-artifact ownership policy — without restating issued-ID retention. Reference `spec:product.concepts.namespace_model.legacy_id_compatibility` (T02) where issued identity is relevant.

## Work

The three concerns to state clearly:

| concern | description |
|---|---|
| Historical ownership decision | V01-ADR-096 decision: all artifacts from the single-app era are treated as PRODUCT-owned; per-app migration is not performed. This is a historical judgment, not a statement about issued-ID form. |
| Effective attribution | For navigation and grouping purposes, which app namespace each domain-prefix group effectively belongs to (`MCP` → DRMCP, `DATA`/`RESOLVE` → BPDSL, `SELFHOST` → cross-app). This is not a retroactive ownership reassignment. |
| New-artifact ownership | When creating new artifacts, use `<APP>-...` when the owning app namespace is confirmed; use `PRODUCT` when cross-app or attribution is unclear. |

Steps:

| area | required work |
|---|---|
| read current file | Read `existing-artifacts.md` in full. |
| identify conflation | Locate where effective attribution and historical ownership are mixed in the same sentence or table row. |
| rewrite | Restructure into three clearly named subsections or sections for the three concerns above. Retain the existing attribution table under the effective-attribution section. Where issued-ID retention was referenced, replace with a reference to `spec:product.concepts.namespace_model.legacy_id_compatibility`. |
| stable ref / parent | Retain the existing `id` (`spec:product.concepts.namespace_model.existing_artifacts`) and `parent` unchanged. Update `date`. |

## Done condition

| item | done when |
|---|---|
| historical ownership | V01-ADR-096 decision is stated as a historical ownership judgment, distinct from attribution. |
| effective attribution | The domain-prefix → app-namespace attribution table is clearly labeled as a navigation/grouping tool, not a retroactive ownership reassignment. |
| new-artifact ownership | New-artifact guidance is clearly separated from the legacy artifact policy. |
| issued identity not restated | The file does not restate issued-ID retention or renaming policy; it references `spec:product.concepts.namespace_model.legacy_id_compatibility` where the topic is relevant. |
| no contradiction | No sentence implies issued IDs are being renamed or the historical ownership decision is being reversed. |

## Verification

- Confirm the stable ref `spec:product.concepts.namespace_model.existing_artifacts` is retained.
- Confirm no sentence conflates effective attribution with ownership reassignment.
- Confirm issued-ID retention is not stated in this file; a reference to `spec:product.concepts.namespace_model.legacy_id_compatibility` is used instead.

## Evidence

- Rewrote `product/records/spec/concepts/namespace-model/existing-artifacts.md` in place. Stable ref `spec:product.concepts.namespace_model.existing_artifacts` and parent retained unchanged. Date updated to 2026-06-24.
- Three sections now clearly separated: (1) "Historical ownership decision" — V01-ADR-096 judgment, per-app migration not performed; references `spec:product.concepts.namespace_model.legacy_id_compatibility` for issued-ID details. (2) "Effective attribution" — combined domain-prefix → effective-app-namespace table (incorporating mapping content from `v2-grammar.md`), labeled as logical projection not retroactive reassignment, includes display-form prepend note. (3) "New-artifact ownership" — guidance on when to use `<APP>-...` vs `PRODUCT`; new PRODUCT artifact examples; references `spec:product.concepts.namespace_model.artifact_id_grammar`.
- Stale reference to `spec:product.concepts.namespace_model.v2_grammar` removed; replaced with reference to `spec:product.concepts.namespace_model.artifact_id_grammar`.
- Issued-ID retention is not stated in this file; `spec:product.concepts.namespace_model.legacy_id_compatibility` is referenced instead.

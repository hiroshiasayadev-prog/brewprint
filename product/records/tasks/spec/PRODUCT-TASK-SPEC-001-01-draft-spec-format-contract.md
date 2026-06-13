# PRODUCT-TASK-SPEC-001-01: Draft spec format contract

- **id**: PRODUCT-TASK-SPEC-001-01
- **status**: done
- **date**: 2026-06-10
- **work_item**: PRODUCT-WORK-SPEC-001
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **source_investigation**: PRODUCT-INV-SPEC-001
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - `product/records/spec/concepts/spec-format/index.md`

## Goal

Draft the PRODUCT-level spec format contract that defines the MCP-readable spec document format and topic table rules accepted by PRODUCT-WORK-SPEC-001.

## Work

| area | required work |
|---|---|
| deliverable path | Create or confirm `product/records/spec/concepts/spec-format/index.md`. |
| spec kind set | Define accepted kinds: `Overview`, `Index`, `Concept`, `Reference`, `Contract`. |
| deferred kinds | Define deferred kinds and revisit conditions: `Guide`, `Process`, `Architecture`, `Glossary`. |
| H1 format | Define `# <SpecKind>: <Title>` and parser-aware one-H1 rule. |
| required sections | Add required / recommended / prohibited section matrix by kind. |
| Topics table | Define required columns: `title`, `kind`, `parent`, `file`, `summary`. |
| duplicate parent | Define duplicate-parent prohibition for authoritative topic declarations. |
| Overview+Topics | Define when `Overview` may contain `## Topics` and when pure `Index` is preferred. |
| spec ID-as-ref derivation | Define path-derived `spec:` ID suggestion, including `index.md` omission and underscore segment policy. |
| validation warning | Define mismatch warning when visible `id` drifts from path-derived default ID. |
| parent grammar | Define safe interim grammar and explicitly prohibit path / filename / H1 title / derived topic ref as parent values. |
| stable ID-as-ref boundary | Preserve stable `spec:` IDs as canonical and defer alias / derived topic compatibility to PRODUCT-WORK-SPEC-002. |

## Done condition

| item | done when |
|---|---|
| format spec draft | `product/records/spec/concepts/spec-format/index.md` exists or an alternative path is documented. |
| matrices | Spec kind matrix and required section matrix are present. |
| ID-as-ref rule | `index.md` omission, underscore segment policy, no separate `ref`, and path-derived default warning rule are present. |
| parent grammar | Interim grammar and forbidden parent forms are present. |
| migration safety | Existing specs are not bulk-migrated. |

## Verification

- Review the drafted spec against PRODUCT-WORK-SPEC-001 Done Condition.
- Confirm no `v01/records/spec/**` files were modified.
- Confirm no DRMCP implementation files were changed.

## Evidence

- Created `product/records/spec/concepts/spec-format/index.md`.
- Draft includes accepted/deferred spec kind matrix.
- Draft includes required section matrix.
- Draft defines `## Topics` table columns and duplicate-parent rule.
- Draft defines path-derived `spec:` ID-as-ref suggestion with `index.md` omission and underscore segment policy.
- Draft removes separate `ref`; spec `id` is the canonical reference target.
- Draft defines warning-only behavior for path-derived default spec ID mismatch.
- Draft preserves stable `spec:` ID-as-ref behavior and defers alias / derived topic compatibility to PRODUCT-WORK-SPEC-002.
- No existing spec bulk migration performed.
- No DRMCP implementation files changed.
- No `v01/records/spec/**` files changed.

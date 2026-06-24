# PRODUCT-TASK-NAMESPACE-001-04: Restore DRMCP ID normalization ownership

- **id**: PRODUCT-TASK-NAMESPACE-001-04
- **status**: completed
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-NAMESPACE-001
- **source_requirement**: V01-REQ-PRODUCT-001
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-NAMESPACE-001-01
- **outputs**:
  - `drmcp/records/spec/design-records-mcp/schema/id-normalization.md` (restored)
  - `product/records/spec/concepts/namespace-model/v1-id-grammar.md` (retained temporarily; deletion owned by T07)

## Goal

Restore the complete DRMCP ID-normalization contract. The obsolete PRODUCT source file remains temporarily until T07 deletes it.

## Work

| area | required work |
|---|---|
| restore DRMCP spec | Restore `drmcp/records/spec/design-records-mcp/schema/id-normalization.md`: remove the placeholder text, write the full public-ID formula and examples, bare-ID grammar table, and parser-stripping rule. Update `date` for the substantive contract change; preserve `status` unless a separate maturity decision is established. Content source: current `v1-id-grammar.md`. |
| PRODUCT-side disposition | Using the T01 ownership map: retain `v1-id-grammar.md` temporarily in this task. Do not introduce a redirect stub. Deletion is owned by T07 after T06 completes reference migration. If T01 records a concrete compatibility requirement for retention beyond T07, document it explicitly in Evidence. |
| no index.md edit | `index.md` Topics table is handled by T02. Do not edit `index.md` in this task. |

## Done condition

| item | done when |
|---|---|
| DRMCP spec restored | `id-normalization.md` contains the full public-ID formula, examples, and bare-ID grammar table; placeholder text removed. |
| PRODUCT file retained | `v1-id-grammar.md` remains in the repository temporarily; all DRMCP parser behavior has been moved to `id-normalization.md`. Deletion is deferred to T07. |

## Verification

- Confirm `id-normalization.md` no longer contains "relocated to PRODUCT" text.
- Confirm `id-normalization.md` contains the full DRMCP parser contract. The deletion of `v1-id-grammar.md` is deferred to T07.

## Evidence

- Restored `drmcp/records/spec/design-records-mcp/schema/id-normalization.md`. Placeholder text ("relocated to PRODUCT") removed. Full DRMCP ID-normalization contract written: public-ID formula (`namespace_prefix + bare_id`), public-ID examples table (single-root mode), tool-API statement (bare IDs not accepted as inputs), bare-ID grammar table (all six kinds), disambiguation note. Date updated to 2026-06-24; status preserved as draft.
- `product/records/spec/concepts/namespace-model/v1-id-grammar.md` retained in repository; deletion deferred to T07.
- T01 ownership map confirmed: entire content of `v1-id-grammar.md` is `DRMCP-implementation`. No PRODUCT-compatibility content was present in that file. No index.md edit made.

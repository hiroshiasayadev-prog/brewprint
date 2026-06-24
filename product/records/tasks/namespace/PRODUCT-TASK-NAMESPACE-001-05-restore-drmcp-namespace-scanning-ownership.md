# PRODUCT-TASK-NAMESPACE-001-05: Restore DRMCP namespace scanning ownership

- **id**: PRODUCT-TASK-NAMESPACE-001-05
- **status**: completed
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-NAMESPACE-001
- **source_requirement**: V01-REQ-PRODUCT-001
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-NAMESPACE-001-01
- **outputs**:
  - `drmcp/records/spec/design-records-mcp/namespace-scanning.md` (restored)
  - `product/records/spec/concepts/namespace-model/v1-namespace-algorithm.md` (retained temporarily; deletion owned by T07)

## Goal

Restore the complete DRMCP namespace-scanning contract. The obsolete PRODUCT source file remains temporarily until T07 deletes it.

## Work

| area | required work |
|---|---|
| restore DRMCP spec | Restore `drmcp/records/spec/design-records-mcp/namespace-scanning.md`: remove the placeholder text, write the full namespace_prefix derivation formula and examples, kind-level prefix application table, and multi-root scan behavior (default mode and single-root mode). Update `date` for the substantive contract change; preserve `status` unless a separate maturity decision is established. Content source: current `v1-namespace-algorithm.md`. |
| PRODUCT-side disposition | Using the T01 ownership map: `v1-namespace-algorithm.md` is entirely DRMCP-implementation content. Retain it temporarily in this task. Do not introduce a redirect stub. Deletion is owned by T07 after T06 completes reference migration. If T01 records a concrete compatibility requirement for retention beyond T07, document it explicitly in Evidence. |
| no index.md edit | `index.md` Topics table is handled by T02. Do not edit `index.md` in this task. |

## Done condition

| item | done when |
|---|---|
| DRMCP spec restored | `namespace-scanning.md` contains the full namespace_prefix derivation, kind-level prefix table, and multi-root scan behavior; placeholder text removed. |
| PRODUCT file retained | `v1-namespace-algorithm.md` remains in the repository temporarily; all DRMCP scanning and prefix-derivation behavior has been moved to `namespace-scanning.md`. Deletion is deferred to T07. |

## Verification

- Confirm `namespace-scanning.md` no longer contains "relocated to PRODUCT" text.
- Confirm `namespace-scanning.md` contains the full DRMCP scanning contract. The deletion of `v1-namespace-algorithm.md` is deferred to T07.

## Evidence

- Restored `drmcp/records/spec/design-records-mcp/namespace-scanning.md`. Placeholder text ("relocated to PRODUCT") removed. Full DRMCP namespace-scanning contract written: `namespace_prefix` derivation formula with `strings.ToUpper(appNamespaceDir) + "-"`, derivation table (v01/records → V01-, drmcp/records → DRMCP-), kind-level prefix application table (four kinds: ADR/Spec/Investigation/Workflow), multi-root scan (default and single-root modes). Date updated to 2026-06-24; status preserved as draft.
- Internal cross-reference updated: "Bare ID grammar is defined in `spec:drmcp.design_records_mcp.schema.id_normalization`" (previously pointed to the now-obsolete PRODUCT location).
- `product/records/spec/concepts/namespace-model/v1-namespace-algorithm.md` retained in repository; deletion deferred to T07.
- T01 ownership map confirmed: entire content of `v1-namespace-algorithm.md` is `DRMCP-implementation`. No index.md edit made.

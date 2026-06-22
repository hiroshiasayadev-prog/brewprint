# PRODUCT-TASK-SPEC-005-20: DRMCP Phase 2 — Opus review of relocation output

- **id**: PRODUCT-TASK-SPEC-005-20
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-17
  - PRODUCT-TASK-SPEC-005-18
  - PRODUCT-TASK-SPEC-005-19
- **outputs**:
  - Review findings (inline in Evidence)

## Goal

Independent Opus review of the Phase 2 relocation output. Confirm content completeness (nothing dropped between DRMCP source and PRODUCT destination), cross-ref accuracy, drift-guard correctness, and validator cleanliness before finalizing.

## Work

| area | what to check |
|---|---|
| completeness — namespace-model | `namespace-model/index.md §v1 namespace resolution algorithm` contains all content from the former `namespace-scanning.md` core body (namespace_prefix formula, kind-別 prefix table, multi-root scan default). |
| completeness — id grammar | `namespace-model/index.md §v1 record ID grammar` contains all content from the former `schema/id-normalization.md`. |
| completeness — discovery | `repository-layout/index.md §Record discovery paths` contains the path-pattern conventions. `schema/discovery.md` retains the DRMCP kind-inclusion filter. |
| resolver trim | `resolver.md` retains only DRMCP-specific scope sections; duplicate resolver semantics removed; pointer to `spec:product.concepts.traceability.resolve_and_validation` present. |
| drift guards | Both hybrid sections in `resolve-and-validation.md` carry `> **Drift guard**:` notes. Wording accurately reflects the PRODUCT-owned vs. DRMCP-vocabulary split. |
| deferred relocation notes | Removed from all 4 formerly-flagged files (`namespace-scanning.md`, `schema/id-normalization.md`, `schema/discovery.md`, `resolver.md`). |
| cross-refs | DRMCP pointer files use canonical `spec:product.concepts.*` refs, not physical paths. |
| validator | `validate_spec.py product/records/spec/concepts/ --strict` and `drmcp/records/spec/design-records-mcp/ --strict` both exit 0. |

## Done condition

| item | done when |
|---|---|
| review complete | Opus review report attached in Evidence. |
| findings classified | Each finding classified: must-fix before -21, or defer. |
| user sign-off | User approves proceeding to PRODUCT-TASK-SPEC-005-21. |

## Verification

This task is itself a review gate. Proceed to PRODUCT-TASK-SPEC-005-21 only after user sign-off.

Per Opus model policy (feedback memory): give user a ready-to-run Opus 4.7 prompt to run externally. Never spawn Agent model "opus". Cross-check all findings against actual files before applying.

## Evidence

### Verdict: PASS — no must-fix issues

2 defer items from Opus review (stale `docs/spec/project-layout.md` path refs in `repository-layout/index.md`; task-spec wording mismatch on §MVP exclusions — implementation correct). 2 additional corrections applied from user review:

1. **ADR discovery pattern corrected**: `<records_root>/adr/*.md` → `<records_root>/adr/<namespace_prefix>ADR-*.md` in `record-discovery-paths.md`. The old glob was a legacy artifact from the pre-migration DRMCP spec; the intended format follows the same `<namespace_prefix>KIND-*` convention as all other kinds. Note added to `schema/discovery.md`.
2. **Hyphen in namespace_prefix**: confirmed correct — `namespace_prefix` includes the trailing hyphen (e.g. `V01-`), so `<namespace_prefix>INV-*-*.md` = `V01-INV-*-*.md`. No change needed.

### User sign-off

Obtained 2026-06-22.

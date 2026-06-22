# PRODUCT-TASK-SPEC-005-19: DRMCP Phase 2 — trim resolver.md and add drift guards

- **id**: PRODUCT-TASK-SPEC-005-19
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-16
- **outputs**:
  - `drmcp/records/spec/design-records-mcp/resolver.md` (trimmed to pointer)
  - `product/records/spec/concepts/traceability/resolve-and-validation.md` (drift guards added)

## Goal

`resolver.md` currently duplicates resolver semantics already owned by `traceability/resolve-and-validation.md`. Trim the DRMCP file to a pointer — no content moves, because PRODUCT already holds it. Then add drift-guard editorial notes to the two hybrid sections in `resolve-and-validation.md` per PRODUCT-TASK-SPEC-004-01 Phase 2 plan, steps 4–5.

This task does not depend on PRODUCT-WORK-SPEC-009 completing — `traceability/resolve-and-validation.md` is edited in place regardless of its format-migration status (drift guards are editorial inserts, not a full migration).

## Work

| step | action | notes |
|---|---|---|
| 4 | Trim `drmcp/records/spec/design-records-mcp/resolver.md` | Remove duplicated resolver semantics from `## Current contract`. Replace with one paragraph pointing to `spec:product.concepts.traceability.resolve_and_validation`. Keep H1 / metadata / `## What this is` (updated). Remove deferred relocation note. Keep §MVP exclusions and §Public tool and unsupported inputs (DRMCP-specific operational scope, not a PRODUCT duplicate). |
| 5 | Add drift guards to `product/records/spec/concepts/traceability/resolve-and-validation.md` | In the two hybrid sections identified in PRODUCT-TASK-SPEC-004-01: add `> **Drift guard**: …` blockquotes marking which clause is a PRODUCT semantic rule vs. which allows DRMCP API vocabulary. Exact sections: resolve/resolver-input/lookup-sources and duplicate-detection/unresolved-refs/relation-integrity-validation. |
| — | Validator | `validate_spec.py product/records/spec/concepts/traceability/ --strict` and `drmcp/records/spec/design-records-mcp/ --strict` must both exit 0. |

## Done condition

| item | done when |
|---|---|
| resolver.md trimmed | Duplicate resolver semantics removed; DRMCP-specific scope sections retained; pointer to `spec:product.concepts.traceability.resolve_and_validation` present. |
| drift guards added | Both hybrid sections in `resolve-and-validation.md` carry `> **Drift guard**:` notes per PRODUCT-TASK-SPEC-004-01 §Ambiguity resolutions. |
| validator clean | Both namespaces pass `--strict`. |
| deferred relocation note removed | `resolver.md` no longer carries the `> **Deferred relocation note**` paragraph. |

## Verification

- Confirm no PRODUCT-owned content was dropped — all resolver semantics remain in `resolve-and-validation.md`.
- Confirm `resolver.md` retained §MVP exclusions, §Public tool and unsupported inputs, §Lookup source vs. record kind boundary (DRMCP operational scope, not PRODUCT duplicates).
- Confirm drift-guard wording accurately reflects the PRODUCT-owned vs DRMCP-vocabulary split stated in PRODUCT-TASK-SPEC-004-01 §Ambiguity resolutions.

## Evidence

- `drmcp/records/spec/design-records-mcp/resolver.md` — trimmed: duplicated resolver semantics (required inputs, workflow artifact refs, MVP exclusions) removed; pointer to `spec:product.concepts.traceability.resolve_and_validation` added; retained DRMCP-specific sections (public tool name, unsupported input table, lookup-source vs. record-kind boundary); deferred relocation note removed.
- `product/records/spec/concepts/traceability/resolve-and-validation.md` — drift guards already present from WORK-009 migration. No edits needed.
- `validate_spec.py product/records/spec/concepts/traceability/ drmcp/records/spec/design-records-mcp/ --strict` → `[strict] All 37 file(s) OK.`

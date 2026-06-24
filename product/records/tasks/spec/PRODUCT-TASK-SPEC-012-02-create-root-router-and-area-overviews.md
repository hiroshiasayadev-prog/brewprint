# PRODUCT-TASK-SPEC-012-02: Create root router and area overviews

- **id**: PRODUCT-TASK-SPEC-012-02
- **status**: done
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-01
- **outputs**:
  - `product/records/spec/index.md`
  - `product/records/spec/design-records/index.md`
  - `product/records/spec/brewprint/index.md`
  - `product/records/spec/bpdsl/index.md`

## Goal

Establish the accepted placement router and top-level area contracts before semantic migration begins.

## Work

- Create the root placement router without duplicating child contracts.
- Create the `design-records/` overview and its ownership prohibitions.
- Revise the `brewprint/` overview for profile, current state, namespaces, and compatibility.
- Create the temporary `bpdsl/` overview with its non-canonical staging contract.
- Add only top-level navigation required by `PRODUCT-ADR-SPEC-001`.
- Do not move existing child specifications in this task.

## Done condition

- The root index routes content by semantic ownership.
- Each top-level area defines owned and prohibited content.
- Dependency direction matches `PRODUCT-ADR-SPEC-001`.
- The BPDSL overview defines purpose, allowed content, prohibited work, review boundary, trigger, and exit condition.
- No existing child spec is relocated.

## Verification

- Review all four overview files against the accepted ADR.
- Confirm project registry facts do not appear in the Design Records overview.
- Confirm no canonical BPDSL semantics are claimed by PRODUCT staging.
- Confirm the root index remains navigation-first.

## Evidence

### Outputs

| path | action | result |
|---|---|---|
| `product/records/spec/index.md` | created | Placement router for the three accepted semantic areas. |
| `product/records/spec/design-records/index.md` | created | App-independent Design Records ownership and prohibition contract. |
| `product/records/spec/brewprint/index.md` | revised | Brewprint profile, current-state, namespace, and compatibility ownership contract. |
| `product/records/spec/bpdsl/index.md` | created | Temporary non-canonical BPDSL quarantine and staging contract. |

### Contract checks

| check | result |
|---|---|
| Root routing | Routes to `design-records/`, `brewprint/`, and temporary `bpdsl/` by semantic ownership. |
| Dependency direction | Matches `PRODUCT-ADR-SPEC-001`; Design Records has no normative DRMCP or BPDSL dependency. |
| Generic examples and registries | Requires app-neutral examples where practical and prohibits duplicated current registry tables. |
| Design Records boundary | Owns identity, authoring, format, record placement, traceability, and artifact responsibility semantics. |
| Brewprint boundary | Owns project profile, current repository state, current namespace assignments, and compatibility history. |
| BPDSL staging boundary | Defines temporary purpose, non-canonical status, allowed preservation, prohibited work, review boundary, trigger, four exit dispositions, and area removal or redefinition. |
| Child migration | No child specification was moved, created, or rewritten. The existing Brewprint layout remains the only child topic row. |
| Broad ref synchronization | Not performed. Existing `spec:product.concepts` refs remain for PRODUCT-TASK-SPEC-012-10. |
| App-local handoff | No DRMCP or BPDSL app-local specification was changed. |
| V01 | No `v01/**` file was changed. |

### Validation

| command | exit code | output |
|---|---:|---|
| `python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color` | 0 | `[strict]  All 43 file(s) OK.` |

### Working tree status

Command:

```text
git status --short
```

Output:

```text
 M product/records/requirements/PRODUCT-REQ-SPEC-001-mcp-readable-spec-format-and-topic-tree.md
 M product/records/spec/brewprint/index.md
 M product/records/tasks/spec/PRODUCT-TASK-SPEC-009-04-opus-review.md
 M product/records/work-items/spec/PRODUCT-WORK-SPEC-009-format-only-migration-traceability-and-artifact-model-specs.md
?? product/records/adr/
?? product/records/investigations/spec/PRODUCT-INV-SPEC-005-product-spec-semantic-layer-and-top-level-ownership-structure.md
?? product/records/spec/bpdsl/
?? product/records/spec/design-records/
?? product/records/spec/index.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-01-confirm-migration-manifest-and-validation-baseline.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-02-create-root-router-and-area-overviews.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-03-split-namespace-profile-and-compatibility.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-04-split-artifact-model-and-repository-layout.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-05-reconcile-traceability-and-extract-future-material.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-06-relocate-authoring-standards-and-spec-format.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-07-relocate-cleaned-semantic-areas.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-08-apply-drmcp-app-local-handoff.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-09-finalize-temporary-bpdsl-staging.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-10-synchronize-refs-mechanically.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-11-validate-and-clean-migration-diagnostics.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-12-independent-restructuring-review.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-13-apply-review-corrections-and-close.md
?? product/records/work-items/spec/PRODUCT-WORK-SPEC-012-product-spec-semantic-layer-restructuring.md
```

T02 changes are limited to the four overview outputs listed above and this task record.
The dirty PRODUCT-REQ-SPEC-001, PRODUCT-TASK-SPEC-009-04, PRODUCT-WORK-SPEC-009, and wider PRODUCT-SPEC-012 planning records are pre-existing or unrelated to T02 close-out.

### Scoped checks

| command | exit code | output |
|---|---:|---|
| `git status --short -- product/records/spec` | 0 | ` M product/records/spec/brewprint/index.md`<br>`?? product/records/spec/bpdsl/`<br>`?? product/records/spec/design-records/`<br>`?? product/records/spec/index.md` |
| `git status --short -- drmcp/records/spec bpdsl/records/spec v01` | 0 | No output. |
| `git diff --name-status -- product/records/spec` | 0 | `M	product/records/spec/brewprint/index.md`<br>`warning: in the working copy of 'product/records/spec/brewprint/index.md', LF will be replaced by CRLF the next time Git touches it` |

### Close-out judgment

Every T02 Done condition remains satisfied:

- The root index routes content by semantic ownership.
- Each top-level area defines owned and prohibited content.
- Dependency direction matches `PRODUCT-ADR-SPEC-001`.
- The BPDSL overview defines purpose, allowed content, prohibited work, review boundary, trigger, and exit condition.
- No existing child spec is relocated.

# PRODUCT-TASK-SPEC-012-07: Relocate cleaned semantic areas

- **id**: PRODUCT-TASK-SPEC-012-07
- **status**: done
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-06
- **outputs**:
  - `product/records/spec/design-records/namespace-model/`
  - `product/records/spec/design-records/repository-layout/`
  - `product/records/spec/design-records/traceability/`
  - `product/records/spec/design-records/artifact-model/`

## Goal

Relocate semantically cleaned Design Records specifications into the accepted target areas.

## Work

- Move cleaned namespace-model files into `design-records/namespace-model/`.
- Move cleaned repository-layout files into `design-records/repository-layout/`.
- Move cleaned traceability files into `design-records/traceability/`.
- Move cleaned project-artifact-model files into `design-records/artifact-model/`.
- Preserve reviewed semantic content during relocation.
- Record old-to-new paths for later mechanical ref synchronization.
- Do not perform broad ref updates in this task.

## Done condition

- All cleaned Design Records files reside under their accepted areas.
- No moved file receives an unreviewed semantic change.
- Old source copies do not remain under `concepts/`.
- The old-to-new path map is complete.

## Verification

- Compare each relocated file with its reviewed pre-move content.
- Confirm each file moved exactly once.
- Confirm no Brewprint profile or BPDSL staging file was placed under Design Records.
- Confirm remaining `concepts/` content is accounted for by later tasks or removal.

## Evidence

### Pre-flight

- T03 status: done
- T04 status: done
- T05 status: done
- T06 status: done
- All four source directories confirmed present before move.
- No destination directories existed before move.

### Baseline snapshot

Working-tree snapshot captured before any move into:

```
$env:TEMP/brewprint-t07-baseline-20260624-204025/
```

HEAD was **not** used as the semantic baseline.
T03–T05 changes are in the working tree only; HEAD predates them.
The snapshot was created with `Copy-Item -Recurse` from the four source directories
immediately before any `Move-Item` call.

### Move method

`Move-Item` (PowerShell filesystem operation).
`git mv`, `git add`, `git restore`, `git reset`, and `git clean` were not used.

### 16-file old-to-new path map

All paths relative to `product/records/spec/`.

| old path | new path |
|---|---|
| `concepts/namespace-model/app-namespaces.md` | `design-records/namespace-model/app-namespaces.md` |
| `concepts/namespace-model/artifact-id-grammar.md` | `design-records/namespace-model/artifact-id-grammar.md` |
| `concepts/namespace-model/existing-artifacts.md` | `design-records/namespace-model/existing-artifacts.md` |
| `concepts/namespace-model/index.md` | `design-records/namespace-model/index.md` |
| `concepts/namespace-model/subdomain-model.md` | `design-records/namespace-model/subdomain-model.md` |
| `concepts/repository-layout/index.md` | `design-records/repository-layout/index.md` |
| `concepts/repository-layout/record-discovery-paths.md` | `design-records/repository-layout/record-discovery-paths.md` |
| `concepts/traceability/artifact-refs.md` | `design-records/traceability/artifact-refs.md` |
| `concepts/traceability/index.md` | `design-records/traceability/index.md` |
| `concepts/traceability/metadata-schema.md` | `design-records/traceability/metadata-schema.md` |
| `concepts/traceability/resolve-and-validation.md` | `design-records/traceability/resolve-and-validation.md` |
| `concepts/traceability/semantic-ref.md` | `design-records/traceability/semantic-ref.md` |
| `concepts/project-artifact-model/artifact-responsibility-matrix.md` | `design-records/artifact-model/artifact-responsibility-matrix.md` |
| `concepts/project-artifact-model/change-and-investigation-flow.md` | `design-records/artifact-model/change-and-investigation-flow.md` |
| `concepts/project-artifact-model/index.md` | `design-records/artifact-model/index.md` |
| `concepts/project-artifact-model/traceability-boundary.md` | `design-records/artifact-model/traceability-boundary.md` |

### Destination counts

| destination | files |
|---|---|
| `design-records/namespace-model/` | 5 |
| `design-records/repository-layout/` | 2 |
| `design-records/traceability/` | 5 |
| `design-records/artifact-model/` | 4 |
| total | 16 |

### Old source directories

All four source directories are empty after the move.
`Test-Path` returned `False` for file content; empty directory shells remain physically (acceptable per task spec).

```
product/records/spec/concepts/namespace-model        — empty
product/records/spec/concepts/repository-layout      — empty
product/records/spec/concepts/traceability           — empty
product/records/spec/concepts/project-artifact-model — empty
```

### Metadata changes applied

**Path-derived id prefix changes**

| area | old prefix | new prefix |
|---|---|---|
| namespace-model | `spec:product.concepts.namespace_model` | `spec:product.design_records.namespace_model` |
| repository-layout | `spec:product.concepts.repository_layout` | `spec:product.design_records.repository_layout` |
| traceability | `spec:product.concepts.traceability` | `spec:product.design_records.traceability` |
| artifact-model | `spec:product.concepts.project_artifact_model` | `spec:product.design_records.artifact_model` |

**parent changes**

- 4 area index files: `root` → `spec:product.design_records`
- 12 child files: old area prefix → new area prefix (matching table above)

**Authoritative local Topics ref changes**

| index file | rows updated |
|---|---|
| `design-records/namespace-model/index.md` | 4 |
| `design-records/repository-layout/index.md` | 1 |
| `design-records/traceability/index.md` | 4 |
| `design-records/artifact-model/index.md` | 3 |

### Content-preservation comparison

Each destination was compared line-by-line against the working-tree baseline snapshot.
Allowed differences: H1-adjacent `id`, H1-adjacent `parent`, authoritative Topics `ref` cells.

| file | changed lines | unexpected diffs |
|---|---|---|
| `namespace-model/app-namespaces.md` | 2 | 0 |
| `namespace-model/artifact-id-grammar.md` | 2 | 0 |
| `namespace-model/existing-artifacts.md` | 2 | 0 |
| `namespace-model/index.md` | 6 | 0 |
| `namespace-model/subdomain-model.md` | 2 | 0 |
| `repository-layout/index.md` | 3 | 0 |
| `repository-layout/record-discovery-paths.md` | 2 | 0 |
| `traceability/artifact-refs.md` | 2 | 0 |
| `traceability/index.md` | 6 | 0 |
| `traceability/metadata-schema.md` | 2 | 0 |
| `traceability/resolve-and-validation.md` | 2 | 0 |
| `traceability/semantic-ref.md` | 2 | 0 |
| `artifact-model/artifact-responsibility-matrix.md` | 2 | 0 |
| `artifact-model/change-and-investigation-flow.md` | 2 | 0 |
| `artifact-model/index.md` | 5 | 0 |
| `artifact-model/traceability-boundary.md` | 2 | 0 |

**Result: 16/16 files — no unexpected differences.**

### Removed-file non-restoration

| file | removed by | restored |
|---|---|---|
| `concepts/namespace-model/domain-catalog.md` | T03 | No |
| `concepts/namespace-model/legacy-id-compatibility.md` | T03 | No |
| `concepts/project-artifact-model/design-flow.md` | T04 | No |
| `concepts/traceability/coverage-mapping.md` | T05 | No |
| `concepts/traceability/out-of-scope.md` | T05 | No |

### `design-records/index.md` Topics (six areas)

| title | kind | ref |
|---|---|---|
| Namespace model | Overview | `spec:product.design_records.namespace_model` |
| Authoring standards | Index | `spec:product.design_records.authoring_standards` |
| Spec format | Index | `spec:product.design_records.spec_format` |
| Repository layout | Overview | `spec:product.design_records.repository_layout` |
| Traceability | Overview | `spec:product.design_records.traceability` |
| Artifact model | Overview | `spec:product.design_records.artifact_model` |

Topic map updated from stale expectation wording to current factual statement.

### Validation

```
Command : python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color
Exit    : 0
Output  : [strict]  All 47 file(s) OK.
```

### Scope and Git-index checks

- `git diff --cached --name-status`: empty — no staged changes.
- `drmcp/records/spec`, `bpdsl/records/spec`, `v01`: no changes.
- `product/records/spec/brewprint/**`, `product/records/spec/bpdsl/**`: unchanged.
- No Brewprint profile or BPDSL staging file was placed under `design-records/`.

### T05 advisories carried unchanged

Two accepted low-severity advisories deferred to T13 were moved without modification:

- `artifact-model/traceability-boundary.md` line 49: body text references physical path
  `product/records/spec/concepts/traceability/` (stale path after move; T13 handles).
- Several files use shorthand such as "recorded in T05".

### Deferred T10 body refs

Body refs using `spec:product.concepts.*` remain in prose, `## Related specs` tables,
navigation sections, examples, and cross-area pointers throughout the moved files.
These are intentionally preserved unchanged for T10 mechanical ref synchronization.

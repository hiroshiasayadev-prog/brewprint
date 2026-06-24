# PRODUCT-TASK-SPEC-012-06: Relocate authoring standards and spec format

- **id**: PRODUCT-TASK-SPEC-012-06
- **status**: done
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-03
  - PRODUCT-TASK-SPEC-012-04
  - PRODUCT-TASK-SPEC-012-05
- **outputs**:
  - `product/records/spec/design-records/authoring-standards/`
  - `product/records/spec/design-records/spec-format/`

## Goal

Relocate clear single-owner Design Records specifications without semantic rewrites.

## Work

- Move `authoring-standards/` under `design-records/`.
- Move `spec-format/` under `design-records/`.
- Preserve file content except path-derived metadata required by the accepted move.
- Record all moved paths for later mechanical ref synchronization.
- Do not update downstream refs in this task.
- Do not combine unrelated semantic cleanup with the moves.

## Done condition

- Both areas exist under `design-records/`.
- Source copies under `concepts/` no longer exist.
- File content changes are limited to move-required metadata.
- Downstream refs remain intentionally deferred to the ref-sync task.

## Verification

- Compare moved files against their pre-move versions.
- Confirm no substantive contract wording changed.
- Confirm all expected files moved exactly once.
- Confirm `v01/` remains unchanged.

## Evidence

### Move operations

| operation | method | result |
|---|---|---|
| `concepts/authoring-standards/` → `design-records/authoring-standards/` | PowerShell `Move-Item` | 10 files moved |
| `concepts/spec-format/` → `design-records/spec-format/` | PowerShell `Move-Item` | 7 files moved |

### Path map — authoring-standards (10 files)

| source path | destination path |
|---|---|
| `concepts/authoring-standards/index.md` | `design-records/authoring-standards/index.md` |
| `concepts/authoring-standards/writing-standard.md` | `design-records/authoring-standards/writing-standard.md` |
| `concepts/authoring-standards/agent-authoring-policy.md` | `design-records/authoring-standards/agent-authoring-policy.md` |
| `concepts/authoring-standards/artifact-boundary.md` | `design-records/authoring-standards/artifact-boundary.md` |
| `concepts/authoring-standards/adr-authoring.md` | `design-records/authoring-standards/adr-authoring.md` |
| `concepts/authoring-standards/requirement-authoring.md` | `design-records/authoring-standards/requirement-authoring.md` |
| `concepts/authoring-standards/work-item-authoring.md` | `design-records/authoring-standards/work-item-authoring.md` |
| `concepts/authoring-standards/task-authoring.md` | `design-records/authoring-standards/task-authoring.md` |
| `concepts/authoring-standards/investigation-authoring.md` | `design-records/authoring-standards/investigation-authoring.md` |
| `concepts/authoring-standards/spec-authoring.md` | `design-records/authoring-standards/spec-authoring.md` |

### Path map — spec-format (7 files)

| source path | destination path |
|---|---|
| `concepts/spec-format/index.md` | `design-records/spec-format/index.md` |
| `concepts/spec-format/document-shape.md` | `design-records/spec-format/document-shape.md` |
| `concepts/spec-format/follow-up-boundary.md` | `design-records/spec-format/follow-up-boundary.md` |
| `concepts/spec-format/overview.md` | `design-records/spec-format/overview.md` |
| `concepts/spec-format/spec-id-as-ref.md` | `design-records/spec-format/spec-id-as-ref.md` |
| `concepts/spec-format/topics-table.md` | `design-records/spec-format/topics-table.md` |
| `concepts/spec-format/validation-policy.md` | `design-records/spec-format/validation-policy.md` |

### Metadata changes — authoring-standards

| file | field | old value | new value |
|---|---|---|---|
| `authoring-standards/index.md` | `id` | `spec:product.concepts.authoring_standards` | `spec:product.design_records.authoring_standards` |
| `authoring-standards/index.md` | `parent` | `root` | `spec:product.design_records` |
| `authoring-standards/index.md` | `## Topics` rows (9) | `spec:product.concepts.authoring_standards.*` | `spec:product.design_records.authoring_standards.*` |
| `authoring-standards/writing-standard.md` | `id` | `spec:product.concepts.authoring_standards.writing_standard` | `spec:product.design_records.authoring_standards.writing_standard` |
| `authoring-standards/writing-standard.md` | `parent` | `spec:product.concepts.authoring_standards` | `spec:product.design_records.authoring_standards` |
| `authoring-standards/agent-authoring-policy.md` | `id` | `spec:product.concepts.authoring_standards.agent_authoring_policy` | `spec:product.design_records.authoring_standards.agent_authoring_policy` |
| `authoring-standards/agent-authoring-policy.md` | `parent` | `spec:product.concepts.authoring_standards` | `spec:product.design_records.authoring_standards` |
| `authoring-standards/artifact-boundary.md` | `id` | `spec:product.concepts.authoring_standards.artifact_boundary` | `spec:product.design_records.authoring_standards.artifact_boundary` |
| `authoring-standards/artifact-boundary.md` | `parent` | `spec:product.concepts.authoring_standards` | `spec:product.design_records.authoring_standards` |
| `authoring-standards/adr-authoring.md` | `id` | `spec:product.concepts.authoring_standards.adr_authoring` | `spec:product.design_records.authoring_standards.adr_authoring` |
| `authoring-standards/adr-authoring.md` | `parent` | `spec:product.concepts.authoring_standards` | `spec:product.design_records.authoring_standards` |
| `authoring-standards/requirement-authoring.md` | `id` | `spec:product.concepts.authoring_standards.requirement_authoring` | `spec:product.design_records.authoring_standards.requirement_authoring` |
| `authoring-standards/requirement-authoring.md` | `parent` | `spec:product.concepts.authoring_standards` | `spec:product.design_records.authoring_standards` |
| `authoring-standards/work-item-authoring.md` | `id` | `spec:product.concepts.authoring_standards.work_item_authoring` | `spec:product.design_records.authoring_standards.work_item_authoring` |
| `authoring-standards/work-item-authoring.md` | `parent` | `spec:product.concepts.authoring_standards` | `spec:product.design_records.authoring_standards` |
| `authoring-standards/task-authoring.md` | `id` | `spec:product.concepts.authoring_standards.task_authoring` | `spec:product.design_records.authoring_standards.task_authoring` |
| `authoring-standards/task-authoring.md` | `parent` | `spec:product.concepts.authoring_standards` | `spec:product.design_records.authoring_standards` |
| `authoring-standards/investigation-authoring.md` | `id` | `spec:product.concepts.authoring_standards.investigation_authoring` | `spec:product.design_records.authoring_standards.investigation_authoring` |
| `authoring-standards/investigation-authoring.md` | `parent` | `spec:product.concepts.authoring_standards` | `spec:product.design_records.authoring_standards` |
| `authoring-standards/spec-authoring.md` | `id` | `spec:product.concepts.authoring_standards.spec_authoring` | `spec:product.design_records.authoring_standards.spec_authoring` |
| `authoring-standards/spec-authoring.md` | `parent` | `spec:product.concepts.authoring_standards` | `spec:product.design_records.authoring_standards` |

### Metadata changes — spec-format

| file | field | old value | new value |
|---|---|---|---|
| `spec-format/index.md` | `id` | `spec:product.concepts.spec_format` | `spec:product.design_records.spec_format` |
| `spec-format/index.md` | `parent` | `root` | `spec:product.design_records` |
| `spec-format/index.md` | `## Topics` rows (6) | `spec:product.concepts.spec_format.*` | `spec:product.design_records.spec_format.*` |
| `spec-format/document-shape.md` | `id` | `spec:product.concepts.spec_format.document_shape` | `spec:product.design_records.spec_format.document_shape` |
| `spec-format/document-shape.md` | `parent` | `spec:product.concepts.spec_format` | `spec:product.design_records.spec_format` |
| `spec-format/follow-up-boundary.md` | `id` | `spec:product.concepts.spec_format.follow_up_boundary` | `spec:product.design_records.spec_format.follow_up_boundary` |
| `spec-format/follow-up-boundary.md` | `parent` | `spec:product.concepts.spec_format` | `spec:product.design_records.spec_format` |
| `spec-format/overview.md` | `id` | `spec:product.concepts.spec_format.overview` | `spec:product.design_records.spec_format.overview` |
| `spec-format/overview.md` | `parent` | `spec:product.concepts.spec_format` | `spec:product.design_records.spec_format` |
| `spec-format/spec-id-as-ref.md` | `id` | `spec:product.concepts.spec_format.spec_id_as_ref` | `spec:product.design_records.spec_format.spec_id_as_ref` |
| `spec-format/spec-id-as-ref.md` | `parent` | `spec:product.concepts.spec_format` | `spec:product.design_records.spec_format` |
| `spec-format/topics-table.md` | `id` | `spec:product.concepts.spec_format.topics_table` | `spec:product.design_records.spec_format.topics_table` |
| `spec-format/topics-table.md` | `parent` | `spec:product.concepts.spec_format` | `spec:product.design_records.spec_format` |
| `spec-format/validation-policy.md` | `id` | `spec:product.concepts.spec_format.validation_policy` | `spec:product.design_records.spec_format.validation_policy` |
| `spec-format/validation-policy.md` | `parent` | `spec:product.concepts.spec_format` | `spec:product.design_records.spec_format` |

### design-records/index.md change

Added `## Topics` table (2 rows) before the existing `## Topic map` section:

```
| Authoring standards | Index | `spec:product.design_records.authoring_standards` | Design record selection, authoring, lifecycle, and writing standards. |
| Spec format | Index | `spec:product.design_records.spec_format` | Visible spec shape, identity, topic organization, and validation policy. |
```

### Content-preservation comparison

Baseline: `git show HEAD:<old-path>` for each of the 17 source files.
Comparison: current working-tree file at destination path.
Allowed diff lines: H1-adjacent `id` marker, H1-adjacent `parent` marker, authoritative `## Topics` table ref cells.

| file | changed-line pairs | unexpected diffs |
|---|---|---|
| `authoring-standards/index.md` | 11 (id + parent + 9 Topics rows) | none |
| `authoring-standards/writing-standard.md` | 2 (id + parent) | none |
| `authoring-standards/agent-authoring-policy.md` | 2 (id + parent) | none |
| `authoring-standards/artifact-boundary.md` | 2 (id + parent) | none |
| `authoring-standards/adr-authoring.md` | 2 (id + parent) | none |
| `authoring-standards/requirement-authoring.md` | 2 (id + parent) | none |
| `authoring-standards/work-item-authoring.md` | 2 (id + parent) | none |
| `authoring-standards/task-authoring.md` | 2 (id + parent) | none |
| `authoring-standards/investigation-authoring.md` | 2 (id + parent) | none |
| `authoring-standards/spec-authoring.md` | 2 (id + parent) | none |
| `spec-format/index.md` | 8 (id + parent + 6 Topics rows) | none |
| `spec-format/document-shape.md` | 2 (id + parent) | none |
| `spec-format/follow-up-boundary.md` | 2 (id + parent) | none |
| `spec-format/overview.md` | 2 (id + parent) | none |
| `spec-format/spec-id-as-ref.md` | 2 (id + parent) | none |
| `spec-format/topics-table.md` | 2 (id + parent) | none |
| `spec-format/validation-policy.md` | 2 (id + parent) | none |

Result: 17/17 files — no unexpected diffs. All body content, H1 titles, lifecycle status fields, contract wording, and section structure preserved exactly.

### Deferred body refs (T10)

All `## Related specs`, prose, and example-table refs using `spec:product.concepts.authoring_standards.*` and `spec:product.concepts.spec_format.*` in moved files are intentionally left unchanged. Mechanical ref synchronization is owned by PRODUCT-TASK-SPEC-012-10.

### Validation output

```
python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color
[strict]  All 47 file(s) OK.
Exit code: 0
```

Matches T05 baseline of 47 files.

### Scope checks

| check | result |
|---|---|
| `git diff --cached --name-status` | Empty — git index not modified |
| `git status --short -- drmcp/records/spec bpdsl/records/spec v01` | Empty — no changes outside product/records/spec |
| Source concepts/authoring-standards/ files remaining | None |
| Source concepts/spec-format/ files remaining | None |
| New design-records/ files git status | `??` (untracked, not staged) |

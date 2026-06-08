# V01-TASK-MCP-006-02: Workflow metadata strictness contract を判断する

- **id**: V01-TASK-MCP-006-02
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-MCP-006
- **source_requirement**: V01-REQ-MCP-006
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-006-01
- **outputs**:
  - Workflow metadata strictness contract decision
  - Diagnostic category decision for V01-TASK-MCP-006-03 and V01-TASK-MCP-006-04

## Goal

V01-TASK-MCP-006-01 の evidence に基づき、Design Records MCP の workflow artifact metadata validation strictness を public contract として判断する。

## Scope

- 対象 artifact kind は `requirement` / `work_item` / `task` とする。
- 判断対象は required metadata presence、empty scalar、missing list、empty list、date validation、diagnostic category とする。
- `fixture_pending` は valid work item status として扱う。
- V01-REQ-MCP-003 / V01-WORK-MCP-003 の declared relation validation 完了 scope は再オープンしない。
- Orphan diagnostics、progress projection、workflow traversal、task dependency cycle / execution order projection、physical path relation support、`req:` / `work:` / `task:` semantic prefix support は対象外とする。

## Work

1. V01-TASK-MCP-006-01 の gap table を確認する。
2. Required metadata presence validation の採用範囲を決める。
3. Empty scalar / missing list / empty list の扱いを決める。
4. Workflow artifact の `date` validation を導入するか決める。
5. 新規 diagnostic category を追加するか、既存 category に寄せるかを決める。
6. Spec / implementation / tests に反映すべき contract を箇条書きで確定する。

## Expected output

- 採用する workflow metadata strictness contract。
- 採用しない validation とその理由。
- 追加または再利用する diagnostic category の一覧。
- V01-TASK-MCP-006-03 / V01-TASK-MCP-006-04 に渡す implementation input。

## Completion condition

Spec / authoring guidance / implementation / tests に反映可能な粒度で strictness contract が確定している。

## Decision

2026-06-01: Adopt workflow metadata strictness validation for required metadata presence, required scalar emptiness, workflow date format, and metadata-specific diagnostics.

### Adopted contract

- Required metadata presence validation is adopted for `requirement` / `work_item` / `task`.
- Required scalar fields must be present and non-empty.
- Required list fields must be present.
- Empty required list values are allowed unless a later artifact-specific rule explicitly requires non-empty values.
- Empty list items are invalid metadata values.
- Workflow artifact `date` is required and must use strict `YYYY-MM-DD` format.
- Existing status enum validation remains owned by `invalid_status_for_kind`.
- Existing relation target format / existence / reciprocal validation remains owned by the existing workflow relation diagnostics.
- `fixture_pending` remains a valid `work_item` status and must be guarded by regression test.

### Required metadata fields

| kind | required fields |
|---|---|
| `requirement` | `id`, `status`, `date`, `source_refs`, `work_items` |
| `work_item` | `id`, `status`, `date`, `source_requirement`, `impact_refs`, `tasks` |
| `task` | `id`, `status`, `date`, `work_item`, `source_requirement`, `estimate`, `depends_on`, `outputs` |

### List field semantics

| field class | field presence | empty list | empty list item |
|---|---|---|---|
| required list metadata | required | allowed | error |

Rationale:

- Missing list field and intentionally empty list have different authoring meaning.
- Public response may still normalize missing / empty values to `[]` after parsing, but validation must be able to diagnose missing required list metadata.
- Empty list is allowed because early artifact lifecycle states often need no current relation entries.
- `task.depends_on: []` is explicitly valid for first tasks.

### Scalar field semantics

| field class | missing | empty |
|---|---|---|
| required scalar metadata | error | error |

Required scalar relation fields such as `work_item`, `source_requirement`, and task `source_requirement` must not be silently skipped when missing or empty. Relation validation still handles non-empty malformed / unresolved / mismatched relation targets.

### Date semantics

Workflow artifact `date` is treated as required metadata, not merely authoring decoration.

- Missing `date`: error.
- Empty `date`: error.
- Invalid `date` format: error.
- Accepted format: `YYYY-MM-DD`.

Date validation requires parser support for raw date value and field presence tracking, because current parser discards workflow `date`.

### Diagnostic categories

Add metadata-specific diagnostic categories instead of overloading existing relation / status diagnostics.

| category | meaning |
|---|---|
| `missing_required_metadata` | Required metadata field is absent. |
| `empty_required_metadata` | Required scalar metadata field is present but empty, or required list field has an empty item. |
| `invalid_metadata_value` | Required metadata field is present and non-empty but does not satisfy its value contract, such as invalid date format. |

Existing categories remain in use for their current responsibility:

- `invalid_status_for_kind`: non-empty status not allowed for artifact kind, and current status-empty side effect until implementation is refined.
- `invalid_workflow_id` / `filename_id_mismatch`: workflow identity / filename / H1 consistency.
- `invalid_workflow_relation_target`: non-empty relation target with malformed or wrong-kind reference.
- `unresolved_workflow_relation`: supported non-empty relation target does not resolve.
- `workflow_relation_mismatch`: reciprocal relation mismatch.

### Explicitly out of scope for this decision

- Changing `diagnostics:null` versus `diagnostics:[]` response shape. This is a response contract cleanup candidate and should not block metadata strictness implementation.
- Orphan diagnostics, traversal, progress projection, task dependency cycle / execution order projection, physical path relations, and `req:` / `work:` / `task:` semantic prefixes.
- Making required list fields non-empty by default.

## Implementation input for V01-TASK-MCP-006-03 / V01-TASK-MCP-006-04

- Update public spec / guidance to say required list fields must be present but may be empty unless explicitly stated otherwise.
- Add diagnostic contract for `missing_required_metadata`, `empty_required_metadata`, and `invalid_metadata_value`.
- Add parser metadata presence tracking for required fields.
- Store raw workflow `date` or equivalent validation metadata for `requirement` / `work_item` / `task`.
- Preserve response normalization where useful, but do not let normalization hide validation diagnostics.
- Add tests for missing required metadata, empty scalar metadata, missing list field, empty list item, invalid date, missing date, and valid empty list.
- Add explicit regression test that `fixture_pending` is valid for `work_item`.
- Leave `diagnostics:null` versus `[]` as follow-up unless V01-TASK-MCP-006-03 finds the spec wording forces a local clarification.

## Evidence

- V01-TASK-MCP-006-01 evidence review found that required workflow fields are documented but currently not generally validated.
- V01-TASK-MCP-006-01 evidence review found that workflow `date` is currently required by docs but ignored by parser.
- V01-TASK-MCP-006-01 evidence review found that missing list fields and empty list fields currently collapse to empty slices, so parser presence tracking is required for this decision.
- User decision on 2026-06-01 accepted the proposed contract: required fields are validated, scalar empty values are invalid, list fields are required but empty list values are allowed, metadata-specific diagnostic categories are added, and `diagnostics:null` versus `[]` is not included in this task.

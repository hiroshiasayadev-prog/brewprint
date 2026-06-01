# TASK-MCP-006-05: Runtime verification と close evidence を整理する

- **id**: TASK-MCP-006-05
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-MCP-006
- **source_requirement**: REQ-MCP-006
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-006-03
  - TASK-MCP-006-04
- **outputs**:
  - Runtime verification evidence
  - Close evidence for REQ-MCP-006 / WORK-MCP-006

## Goal

WORK-MCP-006 の完了判断に必要な runtime verification と close evidence を整理し、REQ-MCP-006 / WORK-MCP-006 を close 可能な状態にする。

## Scope

- `validate_records` の runtime behavior が spec / tests と一致することを確認する。
- `requirement` / `work_item` / `task` の metadata validation strictness を確認する。
- `fixture_pending` は valid work item status として扱う。
- REQ-MCP-003 / WORK-MCP-003 の完了済み scope は再オープンしない。
- SELFHOST artifact は直接編集対象にしない。

## Work

1. TASK-MCP-006-03 / TASK-MCP-006-04 の変更結果を確認する。
2. `validate_records` の runtime verification を実施する。
3. 採用した diagnostic category / response shape が spec と一致することを確認する。
4. Required metadata / empty value / date validation の採用範囲が tests と runtime evidence の両方で確認されていることを整理する。
5. REQ-MCP-006 / WORK-MCP-006 の status 更新案と close evidence を作成する。
6. 必要なら follow-up candidate を分離し、WORK-MCP-006 の scope に混ぜない。

## Expected output

- Runtime verification command / result。
- Close evidence summary。
- REQ-MCP-006 / WORK-MCP-006 の status 更新案。
- 残す follow-up がある場合は、その理由と scope boundary。

## Completion condition

WORK-MCP-006 の completion condition を満たす evidence が揃い、close 判断に進める。

## Verification

- `go test -count=1 ./internal/designrecords ./internal/designrecordsmcp` passed during TASK-MCP-006-04 implementation review.
- Runtime validation passed:
  - `validate_records(kind="task", id_range=TASK-MCP-006-01..TASK-MCP-006-05)` returned `ok: true`, `diagnostics: null`.
  - `validate_records(kind="work_item")` returned `ok: true`, `diagnostics: null`.
- Build verification was executed after adding the local Windows build helper script; build completed successfully and generated binaries were updated.
- `diagnostics: null` remains an existing response-shape behavior and is explicitly out of scope for WORK-MCP-006.

## Close review

WORK-MCP-006 completion condition is satisfied.

Completed task chain:

- `TASK-MCP-006-01`: current validator behavior and metadata field gap review completed.
- `TASK-MCP-006-02`: workflow metadata strictness contract decision completed.
- `TASK-MCP-006-03`: Design Records MCP schema / tools spec and authoring guidance updated and reviewed.
- `TASK-MCP-006-04`: parser / validator / tests implemented and reviewed.
- `TASK-MCP-006-05`: runtime verification and close evidence completed.

Requirement coverage:

- Required metadata presence validation is implemented for `requirement` / `work_item` / `task`.
- Required scalar missing / empty validation is implemented.
- Required list field presence validation is implemented.
- Empty list is valid.
- Empty list item returns `empty_required_metadata`.
- Workflow artifact `date` is required and strict `YYYY-MM-DD` validation is implemented.
- Metadata-specific diagnostics are defined in spec and implemented:
  - `missing_required_metadata`
  - `empty_required_metadata`
  - `invalid_metadata_value`
- Metadata diagnostics expose `field`; `value` is emitted when input value exists, including explicit empty string.
- `fixture_pending` remains valid work item status and has regression coverage.
- Workflow `date` remains validation-only metadata and is not added to public get/list responses.

Out-of-scope items intentionally not closed here:

- `diagnostics:null` versus `diagnostics:[]` response shape cleanup.
- README / legacy authoring guide cleanup, tracked separately by `REQ-MCP-009`.
- Orphan diagnostics, workflow traversal, progress projection, task dependency cycle / execution order projection.
- Physical path workflow relations and `req:` / `work:` / `task:` semantic prefixes.

## Close decision

- `TASK-MCP-006-05` can be marked `done`.
- `WORK-MCP-006` can be marked `done`.
- `REQ-MCP-006` remains `accepted`; close evidence is recorded on the requirement because requirement status has no `done` state.

# V01-TASK-MCP-006-04: Workflow metadata validator 実装と tests を追加する

- **id**: V01-TASK-MCP-006-04
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-MCP-006
- **source_requirement**: V01-REQ-MCP-006
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-MCP-006-02
  - V01-TASK-MCP-006-03
- **outputs**:
  - Workflow metadata validator implementation
  - Unit tests / MCP transport tests

## Goal

V01-TASK-MCP-006-02 / V01-TASK-MCP-006-03 で確定・反映した public contract に合わせて、Design Records MCP の workflow artifact metadata parser / validator / MCP tests を実装する。

## Scope

- 主対象は `internal/designrecords` と `internal/designrecordsmcp` の parser / validation / tests とする。
- `requirement` / `work_item` / `task` の metadata validation を対象にする。
- `fixture_pending` は valid work item status として regression test で守る。
- SELFHOST artifact は直接編集対象にしない。
- Orphan diagnostics、progress projection、workflow traversal、task dependency cycle / execution order projection、physical path relation support、`req:` / `work:` / `task:` semantic prefix support は対象外とする。

## Work

1. V01-TASK-MCP-006-03 の spec / guidance update を読む。
2. Parser / validator に required metadata / empty value / date / diagnostic category contract を実装する。
3. Unit tests で missing field / empty scalar / missing list / empty list / invalid date / invalid status / invalid relation target のうち採用範囲を確認する。
4. MCP transport tests で `validate_records` response が public spec と一致することを確認する。
5. `fixture_pending` が valid work item status として通ること、許可外 status が invalid になることを regression test で明示する。

## Expected output

- Parser / validator implementation update。
- Unit tests / MCP tests update。
- 実行した command と結果。

## Completion condition

採用した workflow metadata strictness contract が implementation / tests に反映され、関連 tests が pass している。

## Verification

- `go test -count=1 ./internal/designrecords ./internal/designrecordsmcp` passed.
- Runtime validation passed:
  - `validate_records(kind="task", id_range=V01-TASK-MCP-006-01..V01-TASK-MCP-006-05)` returned `ok: true`, `diagnostics: null`.
  - `validate_records(kind="work_item")` returned `ok: true`, `diagnostics: null`.
- Implementation review verdict: OK to proceed.

## Evidence

Implemented changes:

- `internal/designrecords/parser.go`
  - Added workflow artifact required metadata presence tracking.
  - Added empty list item tracking.
  - Added raw workflow `date` tracking for validation.
  - Did not add workflow `date` to public get/list response.
- `internal/designrecords/types.go`
  - Added validation-only metadata model.
  - Added metadata diagnostic categories: `missing_required_metadata`, `empty_required_metadata`, `invalid_metadata_value`.
  - Adjusted diagnostic JSON marshal so `value: ""` can be emitted when explicitly present.
- `internal/designrecords/validation.go`
  - Implemented `missing_required_metadata`.
  - Implemented `empty_required_metadata`.
  - Implemented `invalid_metadata_value`.
  - Preserved existing relation diagnostic responsibilities for non-empty malformed / unresolved / mismatched relation targets.
- `internal/designrecords/validation_test.go`
  - Added requirement / work_item / task strictness regression tests.
  - Added `fixture_pending` valid work item status regression coverage.
- `internal/designrecordsmcp/tools_call_test.go`
  - Added MCP response coverage for `category`, `severity`, `record_id`, `path`, `message`, `field`, and `value`, including empty string value.

Review result:

- No findings.
- Parser distinguishes missing list, present empty list, and empty list item.
- Validator maps missing required field to `missing_required_metadata`.
- Validator maps empty required scalar and empty list item to `empty_required_metadata`.
- Validator maps invalid workflow `date` to `invalid_metadata_value`.
- Public get/list response does not expose workflow `date`.
- `diagnostics:null` versus `[]` remains out of scope.
- Existing relation validation, task `depends_on` empty list, workflow id_range filtering, get/list response model, investigation / ADR / spec validation are preserved by tests.

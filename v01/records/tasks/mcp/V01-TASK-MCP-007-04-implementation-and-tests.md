# V01-TASK-MCP-007-04: Implementation and tests

- **id**: V01-TASK-MCP-007-04
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-MCP-007
- **source_requirement**: V01-REQ-MCP-007
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-MCP-007-03
- **outputs**:
  - Updated `internal/designrecords` range filter implementation
  - Updated `internal/designrecordsmcp` tool schema / call handling
  - Regression tests for workflow artifact range behavior

## Goal

`V01-TASK-MCP-007-03` で更新した public contract に合わせて、Design Records MCP の `list_records` range navigation を実装し、tests を追加する。

## Work

- `internal/designrecords` の list records query / filtering behavior を更新する。
- `internal/designrecordsmcp` の tool schema と call validation を spec に合わせる。
- Requirement / work item / task の valid range cases を test する。
- Mixed-domain / mixed-kind / malformed endpoint / unsupported kind の invalid cases を test する。
- 既存 ADR decision range behavior の regression test を維持または追加する。

## Done condition

- Workflow artifact range navigation が contract 通りに動作する。
- Unsupported range request が silent reinterpretation されない。
- Existing ADR decision range behavior が regression していない。
- Targeted Go tests が pass している、または out-of-scope failure が明確に分離されている。

## Verification

- `go test ./internal/designrecords ./internal/designrecordsmcp` を実行する。
- 必要なら `go test ./cmd/design-records-mcp ./internal/designrecords ./internal/designrecordsmcp` を実行する。

## Evidence

2026-06-01 implementation update: done.

Implemented files:

- Added `internal/designrecords/id_range.go`.
- Updated `internal/designrecords/types.go`.
- Updated `internal/designrecords/tools.go`.
- Updated `internal/designrecords/validation.go`.
- Updated `internal/designrecords/list_records_test.go`.
- Updated `internal/designrecords/validation_test.go`.
- Updated `internal/designrecordsmcp/tools_call_test.go`.
- Updated `internal/designrecords/types_test.go`.

Implementation summary:

- Added family-aware `recordIDRange` parsing for:
  - `ADR-NNN` / `decision`.
  - `REQ-<DOMAIN>-NNN` / `requirement`.
  - `WORK-<DOMAIN>-NNN` / `work_item`.
  - `TASK-<DOMAIN>-NNN-MM` / `task`.
- `list_records` now parses `id_range` with explicit `kind` validation and derives effective kind from endpoint family when `kind` is omitted.
- `validate_records` now uses the same parser and selection semantics as `list_records`.
- Added `invalid_id_range` error code while leaving `id_range_requires_decision_kind` as legacy.
- Range membership now uses `recordIDRange.containsRecord` / `containsID` instead of ADR-only numeric checks.
- Empty `id_range` preserves the old behavior as an effective kind filter instead of accidentally returning no workflow records.

Test updates:

- Added `TestListRecordsWorkflowIDRangeFilter` for requirement, work item, task, omitted-kind workflow range, one-sided workflow range, and empty workflow range.
- Updated `TestListRecordsRequestErrors` to expect `invalid_id_range` for invalid range cases.
- Added invalid mixed-domain / mixed-family / mixed task work sequence cases.
- Added `TestValidateRecordsWorkflowIDRangeFilter` for workflow range scoped validation.
- Updated `TestValidateRecordsRequestErrors` to expect `invalid_id_range`.
- Updated MCP `tools/call` test error expectation and added a workflow work item range success case.
- Updated `types_test` to include `invalid_id_range`.

Verification status:

- User ran `go test ./internal/designrecords ./internal/designrecordsmcp` on Windows PowerShell.
- Result:
  - `ok github.com/hiroshiasayadev-prog/brewprint/internal/designrecords 1.212s`
  - `ok github.com/hiroshiasayadev-prog/brewprint/internal/designrecordsmcp 0.475s`

Follow-up:

- Runtime verification remains in `V01-TASK-MCP-007-05`.

# V01-TASK-MCP-026-03: tests — multi-op update regression and new cases

- **id**: V01-TASK-MCP-026-03
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-026
- **source_requirement**: V01-REQ-MCP-025
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-MCP-026-02
- **outputs**:
  - Updated internal/designrecords/authoring_test.go

## Goal

`operations` 配列を使った multi-op update の回帰テストを追加する。

## Work

`internal/designrecords/authoring_test.go` に以下のテストケースを追加する:

- 正常系: `metadata_fields_replace` + `named_section_replace` を 1 proposal で適用（Evidence + status:done）
- 後方互換: 既存 `update: {...}` 形式が動作すること
- conflict 系: 同一フィールドに 2 つの metadata op → `conflicting_operations` error diagnostic
- conflict 系: `named_section_replace` が 2 op → `multiple_section_replace_not_supported` error diagnostic
- no-op 系: operations 配列内の全 op が no-op → no-op response
- MCP tool call layer の既存テストが全通過すること（`internal/designrecordsmcp/tools_call_test.go`）

## Done condition

`go test ./internal/designrecords/... ./internal/designrecordsmcp/...` が全通過する。

## Verification

テスト実行ログで全ケース PASS を確認する。

## Evidence
2026-06-07: テスト追加完了。

`internal/designrecords/authoring_test.go` の末尾（旧 line 2289 以降）に以下の11テスト関数を追加:

| 関数名 | 内容 |
|---|---|
| `TestMultiOpUpdateNormalCase` | metadata_fields_replace + named_section_replace で diff に両変更が現れること |
| `TestMultiOpUpdateDoneStateEvidenceGatePassesWithCombinedOps` | status:done + Evidence 同時更新で Evidence セクション diagnostic が出ないこと |
| `TestMultiOpUpdateBackwardCompatSingleOp` | 既存 `update:{}` 形式が引き続き動作すること |
| `TestMultiOpUpdateExclusivity` | update + operations 両方指定 → `invalid_request` |
| `TestMultiOpUpdateEmptyOperations` | 空 operations → `invalid_request` |
| `TestMultiOpUpdateNeitherUpdateNorOperations` | update も operations もなし → `invalid_request` |
| `TestMultiOpConflictDuplicateMetadataField` | 同一フィールドへの 2 metadata op → `conflicting_operations` |
| `TestMultiOpConflictMetadataBlockPlusFields` | metadata_block_replace + metadata_fields_replace → `conflicting_operations` |
| `TestMultiOpConflictMultipleSectionReplace` | named_section_replace が 2 op → `multiple_section_replace_not_supported` |
| `TestMultiOpUpdateNoOp` | 全 op が no-op → `no_op_update` diagnostic |
| `TestMultiOpTopLevelBodyForbidden` | top-level body + operations → `invalid_body_source` |

ヘルパー `assertHasDiagnosticCategory` を末尾に追加（既存 `hasDiagnosticCategory` に委譲）。

`go test ./internal/designrecords/... ./internal/designrecordsmcp/...` 全件 PASS 確認済み。

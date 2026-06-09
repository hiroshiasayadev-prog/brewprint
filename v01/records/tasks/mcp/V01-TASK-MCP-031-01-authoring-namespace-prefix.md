# V01-TASK-MCP-031-01: authoring namespace prefix バグを実装修正しテストを通す

- **id**: V01-TASK-MCP-031-01
- **status**: done
- **date**: 2026-06-09
- **work_item**: V01-WORK-MCP-031
- **source_requirement**: V01-REQ-MCP-034
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - drmcp/src/internal/designrecords/authoring.go
  - drmcp/src/internal/designrecords/authoring_guidance.go
  - drmcp/src/internal/designrecords/authoring_guidance_test.go
  - drmcp/src/internal/designrecords/validation.go
  - drmcp/src/internal/designrecordsmcp/tools_call_test.go

## Goal

authoring.go・authoring_guidance.go・validation.go の namespace prefix バグを修正し、`go test ./drmcp/src/...` を全件 PASS させる。

## Work

- `authoring.go`: `bareRecordID(id, nsPrefix)` ヘルパー追加
- `authoring.go` / `prepareCreate`: `resolved = ns + bare` でファイル生成に public ID を使用
- `authoring.go` / `prepareUpdate`・`prepareMultiOpUpdate`: `Domain` フィールドを bare ID から導出
- `authoring.go` / `nextWorkflowID`・`maxWorkflowSeq`・`nextTaskID`・`maxTaskSeqForParent`・`nextDecisionID`・`collectSubdomainValues`: index record の bare ID で比較
- `authoring.go` / `exactIDGapWarning`・`subdomainAdvisoryDiagnostics`: parentID / id を bare ID 化してから処理
- `authoring_guidance.go` / `loadAuthoringGuides`: `RecordsRoot/guides/` を参照するよう修正
- `validation.go` / `validateWorkflowRelationTarget`: index lookup を bare ID フォーマットチェックより先に実行
- テスト fixture: `docs/guides/` → `v01/records/guides/` に更新

## Done condition

- `go test ./drmcp/src/...` 全件 PASS
- `propose_record_create` が public ID でファイルを生成する

## Verification

- `go test ./drmcp/src/...` を実行して全件 PASS を確認する

## Evidence

- 2026-06-09: 上記ファイルを修正、テスト fixture 更新。`go test ./drmcp/src/...` 全件 PASS 確認。commit 145dcc9（authoring/guidance/test）、commit 46d4875（validation）。

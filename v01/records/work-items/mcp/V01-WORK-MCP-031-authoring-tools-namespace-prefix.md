# V01-WORK-MCP-031: authoring tools の namespace prefix バグを修正する

- **id**: V01-WORK-MCP-031
- **status**: done
- **date**: 2026-06-09
- **source_requirement**: V01-REQ-MCP-034
- **impact_refs**:
  - V01-REQ-MCP-034
- **tasks**:
  - V01-TASK-MCP-031-01

## Goal

WORK-DRMCP-001 で実装した namespace prefix 対応の後、authoring tools 内部のシーケンス解決関数群が public ID をそのまま bare ID positional parser に渡していたことによる複数のバグを修正する。

## Background

WORK-DRMCP-001 完了後に発見されたバグ群（commit 145dcc9 / 46d4875）:

- `nextWorkflowID`・`maxWorkflowSeq`・`nextTaskID`・`maxTaskSeqForParent`・`nextDecisionID`・`collectSubdomainValues` が index 内の public ID をそのまま `workflowDomain/Sequence` に渡す → 常に sequence `001` が返る ID 衝突リスク
- `prepareCreate` が bare ID でファイルを生成 → 作成ファイルの H1 / metadata id に namespace prefix が付かない
- `prepareUpdate` / `prepareMultiOpUpdate` の `target.domain` が誤った domain を返す
- `exactIDGapWarning`・`subdomainAdvisoryDiagnostics` が public `parentID` を bare ID として扱う
- `loadAuthoringGuides` が `docs/guides/` をハードコード → `v01/records/guides/` が参照されない
- `validateWorkflowRelationTarget` が bare ID パターンチェックを index lookup より先に行う → public ID を `invalid_target` 扱いにする

## Boundary

このWORKが所有するもの:
- `drmcp/src/internal/designrecords/authoring.go`: `bareRecordID` ヘルパー追加、`prepareCreate` の public ID 生成、`prepareUpdate`/`prepareMultiOpUpdate` の `Domain` フィールド修正、sequence 解決関数群の bare ID 化、`exactIDGapWarning`/`subdomainAdvisoryDiagnostics` の parentID 処理修正
- `drmcp/src/internal/designrecords/authoring_guidance.go`: `loadAuthoringGuides` の RecordsRoot ベース化
- `drmcp/src/internal/designrecords/validation.go`: `validateWorkflowRelationTarget` の lookup-first 化

このWORKが所有しないもの:
- multi-root スキャン対応（V01-WORK-MCP-030）
- v01/records のコンテンツ変更

## Completion Condition

- `go test ./drmcp/src/...` 全件 PASS
- `propose_record_create` が public ID でファイルを生成する
- relation フィールドに public ID を指定しても validation エラーにならない

## Evidence

- 2026-06-09: `authoring.go`・`authoring_guidance.go`・`validation.go` を修正、テスト fixture 更新。`go test ./drmcp/src/...` 全件 PASS。commit 145dcc9 / 46d4875。

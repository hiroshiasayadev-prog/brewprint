# V01-TASK-MCP-030-05: authoring tools namespace routing 対応

- **id**: V01-TASK-MCP-030-05
- **status**: done
- **date**: 2026-06-09
- **work_item**: V01-WORK-MCP-030
- **source_requirement**: V01-REQ-MCP-033
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-MCP-030-02
  - V01-TASK-MCP-030-03
- **outputs**:
  - drmcp/src/internal/designrecords/authoring.go 更新
  - drmcp/src/internal/designrecords/authoring_test.go 更新

## Goal

`propose_record_create` / `propose_record_update` が public ID の namespace prefix から対象 records root を特定し、正しい records tree にファイルを生成・更新できるようにする。

## Work

- authoring 内の records root 解決ロジックを、単一 `cfg.RecordsRoot` の固定参照から `cfg.RecordsRoots` の prefix lookup に変更
- public ID prefix（例: `DRMCP-`）から対応する `RecordsEntry`（records root + namespace prefix）を特定するルーティング関数を実装
- `propose_record_create` / `propose_record_update` の path 解決を namespace routing 経由に変更
- `authoring_test.go` を multi-root 環境でのルーティングシナリオでカバー

## Done condition

- `propose_record_create` で `DRMCP-*` ID を指定すると `drmcp/records/` 配下にファイルが生成される
- `propose_record_create` で `V01-*` ID を指定すると `v01/records/` 配下にファイルが生成される
- `authoring_test.go` がすべて pass する

## Verification

`go test ./drmcp/src/internal/designrecords/... -run TestAuthoring` が pass する。

## Evidence
`detectCreateNamespace` / `namespacePrefixForID` / `recordsRootForNamespace` / `entryForPath` ヘルパー追加。`prepareCreate` で namespace prefix をリクエスト ID から自動検出し `createRecordPath` / `resolveCreateID` にルーティング。`nextDecisionID` / `nextWorkflowID` / `nextTaskID` に namespace filter 追加。`buildHypotheticalIndex` で per-path namespace ルックアップ実装。全テスト pass。

# V01-TASK-DRMCP-001-05: テストを v01/records ベースに更新し全件 PASS を確認する

- **id**: V01-TASK-DRMCP-001-05
- **status**: done
- **date**: 2026-06-09
- **work_item**: V01-WORK-DRMCP-001
- **source_requirement**: (TBD)
- **estimate**: 3d
- **depends_on**:
  - V01-TASK-DRMCP-001-03
  - V01-TASK-DRMCP-001-04
- **outputs**:
  - drmcp/src/internal/designrecords/ 全テストファイル
  - drmcp/src/internal/designrecordsmcp/tools_call_test.go
  - drmcp/src/cmd/design-records-mcp/main_test.go
  - scripts/verify.bat

## Goal

drmcp/src/ 全パッケージのテストを v01/records ベースに更新し、namespace prefix 付き ID の parse / index / resolve が通ることを確認する。scripts/verify.bat を正しいパスに更新する。

## Work

- designrecords パッケージ: `docs/` パスを `records/` に変更、ADR H1 を `# ADR-NNN:` 形式に更新、bootstrap テストを `NewConfig(root, "v01/records")` ベースに更新（V01- prefix ID を使用）
- designrecordsmcp パッケージ: `NewConfig(root)` → `NewConfig(root, "docs")` に修正（4 箇所）、`toolsCallTestIndex` に `RecordsRoot: "docs"` 追加、suggest_next_record expected path 更新、`TestToolsCallValidateWorkflowMetadataDiagnosticShape` の Config.RecordsRoot 修正
- cmd/design-records-mcp: リポジトリルートパスを `../..` → `../../../../` に修正、`ADR-076` → `V01-ADR-076` に更新
- scripts/verify.bat: `go test` を `./drmcp/src/...` に変更、`go build` パスを `./v01/src/cmd/brewprint` / `./drmcp/src/cmd/design-records-mcp` に更新

## Done condition

- `go test ./drmcp/src/...` 全件 PASS
- `scripts/verify.bat` が OK を返す

## Verification

- `go test ./drmcp/src/...` および `scripts/verify.bat` を実行する

## Evidence

- 2026-06-09: 全テストファイル更新、`go test ./drmcp/src/...` 全件 PASS 確認。`scripts/verify.bat` OK 確認。commit fadb914。

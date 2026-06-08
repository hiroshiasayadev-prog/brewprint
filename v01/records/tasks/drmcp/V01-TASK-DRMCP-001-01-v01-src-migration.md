# V01-TASK-DRMCP-001-01: v01/src の designrecords コードを drmcp/src に移管する

- **id**: V01-TASK-DRMCP-001-01
- **status**: done
- **date**: 2026-06-09
- **work_item**: V01-WORK-DRMCP-001
- **source_requirement**: (TBD)
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - drmcp/src/internal/designrecords/
  - drmcp/src/internal/designrecordsmcp/
  - drmcp/src/cmd/design-records-mcp/

## Goal

v01/src/internal/designrecords / v01/src/internal/designrecordsmcp / v01/src/cmd/design-records-mcp を drmcp/src 配下に移管し、Go import パスを `github.com/hiroshiasayadev-prog/brewprint/drmcp/src/...` に更新する。

## Work

- drmcp/src/internal/designrecords/, drmcp/src/internal/designrecordsmcp/, drmcp/src/cmd/design-records-mcp/ を作成
- v01/src/ の対象ファイルをコピーし import パスを更新
- v01/src/ の対象ファイルを削除

## Done condition

- drmcp/src/ 配下のビルドが通る
- v01/src/ から designrecords / designrecordsmcp / design-records-mcp が削除されている

## Verification

- `go build ./drmcp/src/...` を実行して成功することを確認

## Evidence

- 2026-06-09: v01/src/internal/designrecords・designrecordsmcp・cmd/design-records-mcp を drmcp/src/ に移管。import パスを drmcp/src に更新。v01/src/ 対象ファイル削除（未コミット、後続コミットで反映予定）。`go build ./drmcp/src/...` 成功確認。

# V01-WORK-DRMCP-001: Design Records MCP parser に app namespace prefix 導出ルールを実装する

- **id**: V01-WORK-DRMCP-001
- **status**: done
- **date**: 2026-06-08
- **source_requirement**: (TBD — V01-WORK-PRODUCT-004 migration 起因の follow-up。正式 REQ は別途起票する)
- **impact_refs**:
  - V01-WORK-PRODUCT-004
- **tasks**:
  - V01-TASK-DRMCP-001-01
  - V01-TASK-DRMCP-001-02
  - V01-TASK-DRMCP-001-03
  - V01-TASK-DRMCP-001-04
  - V01-TASK-DRMCP-001-05

## Goal

V01-WORK-PRODUCT-004 で実行した v01/ migration により、Design Records MCP の parser は V01- prefix 付きの artifact ID を認識できない状態になった。本 WORK では、records root のディレクトリ名から namespace prefix を動的に導出するルールを spec 化し、parser に実装することで MCP を復旧させる。

## Background

V01-ADR-097 と V01-ADR-099 によって確定した移行後のリポジトリ構造では:
- `v01/records/` がリポジトリの design records を格納する
- すべての既存 artifact ID に `V01-` prefix が付与されている
- 将来的に `drmcp/records/`、`bpdsl/records/`、`product/records/` が加わる

現在の parser は `docs/` 配下のベア ID（`REQ-*`, `WORK-*`, `ADR-NNN` 等）のみ認識する。namespace prefix 導出ルール（ディレクトリ名大文字化 + `-`）は spec 化されておらず、実装もない。

## Boundary

このWORKが所有するもの:
- `v01/src/internal/designrecords/` / `v01/src/internal/designrecordsmcp/` / `v01/src/cmd/design-records-mcp/` を `drmcp/src/` へ移管し、import パスを更新する
- namespace prefix 導出ルールの spec への記述（`drmcp/records/spec/design-records-mcp/` 内）
- `drmcp/src/internal/designrecords/config.go`: `RecordsRoot` フィールドと `NamespacePrefix()` 導出メソッド追加
- `drmcp/src/internal/designrecords/index.go`: scan パスを `RecordsRoot` ベースに変更し、namespace prefix をパーサーに渡す
- `drmcp/src/internal/designrecords/parser.go`: namespace prefix を受け取り、ファイル名・H1・metadata ID の prefix ストリップ＋復元を実装
- `drmcp/src/cmd/design-records-mcp/main.go`: `--records-root` フラグ追加（デフォルト: `v01/records`）
- Go test の更新・追加

このWORKが所有しないもの:
- multi-root 対応（複数の app namespace を同時にスキャンする機能）
- v01/ コンテンツの廃止・削除
- 各 app namespace への新規 artifact 移行
- bpdsl / product 側の app namespace 移管

## Task Candidates

- TASK-00: `v01/src` の designrecords 関連コード（`internal/designrecords/`・`internal/designrecordsmcp/`・`cmd/design-records-mcp/`）を `drmcp/src/` に移管し、import パスを更新する
- TASK-01: `drmcp/records/spec/design-records-mcp/` に namespace prefix 導出ルール（`<dir>/records/` → `strings.ToUpper(dir) + "-"`）を記述する → spec review ゲート
- TASK-02: `drmcp/src/internal/designrecords/config.go` / `index.go` / `parser.go` に namespace prefix 導出と prefix 対応 ID 処理を実装する
- TASK-03: `drmcp/src/cmd/design-records-mcp/main.go` に `--records-root` フラグを追加し、デフォルトを `v01/records` にする
- TASK-04: 既存テストを `v01/records` ベースに更新し、namespace prefix 付き ID の parse・index が通ることを確認する

## Completion Condition

- `design-records-mcp --root <repo>` が `v01/records/` を scan し、`V01-*` prefix 付き全 records を返せる
- namespace prefix 導出ルールが spec に記述されている
- 既存の Go テストが通る

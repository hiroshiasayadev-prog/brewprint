# Go M7-12 diagnostics spec 追加タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: validation diagnosticsの外部向けspec追加

## 1. 目的

M7-1〜M7-11でdiagnostic code、validate CLI、JSON出力、diagnostic orderingを整えた。

M7-12では、これらを外部向け仕様として `docs/spec/diagnostics.md` にまとめる。

## 2. 実装範囲

- `docs/spec/diagnostics.md` を追加
- Front Matterをdoc-policyに従って付与
- Diagnostic objectのJSON schemaを記述
- text / JSON validate出力を記述
- severityの意味を記述
- diagnostic orderingを記述
- 現在のdiagnostic code一覧をカテゴリ別に記述
- messageは人間向けであり、外部ツールはseverity/code/fileを優先することを明記

## 3. 境界

今回はdoc追加のみ。

```text
実装変更なし
```

## 4. 受け入れ条件

- [x] `docs/spec/diagnostics.md` が追加される
- [x] Front Matterを持つ
- [x] `brewprint validate` text / JSON形式が書かれている
- [x] severity / ordering / code一覧が書かれている
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る

## 5. 検証メモ

2026-04-28にユーザー環境で以下を確認済み。

```sh
cd C:\Users\imved\projects\brewprint
go fmt ./...
go test ./...
```

全package通過。

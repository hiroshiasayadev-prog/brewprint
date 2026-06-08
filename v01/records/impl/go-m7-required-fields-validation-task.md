# Go M7-7 required fields validation 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: semantic validationに必須フィールド欠落チェックを追加

---

## 1. 目的

M7-1〜M7-6でdiagnostic code、validate CLI、JSON出力、warning-only挙動、出力順を整えた。

M7-7では、YAML入力の基本的な欠落をsemantic validationで検出する。

---

## 2. 実装範囲

新しいdiagnostic codeを追加する。

```text
missing_required_field
```

対象は、semantic model上で安全に判定できる以下に限定する。

- node `id`
- model `kind`
- model field `name`
- model field `type`
- store `kind`
- task / branch / fork / join param `name`
- task / branch / fork / join param `model`
- task return `name`
- task return `model`
- join return `name`
- join return `model`
- initialized store `name`
- initialized store `model`

---

## 3. 境界

validationは引き続き `internal/resolve` に置く。

```text
resolve -> rawyaml, semantic
query   -> semantic
mcp     -> query
```

M7-7ではsemanticに残っている値だけを使う。`internal/query` / `internal/mcp` / renderer側にはvalidationを追加しない。

---

## 4. 受け入れ条件

- [x] `missing_required_field` diagnostic codeを追加
- [x] node id欠落を検出する
- [x] model kind欠落を検出する
- [x] model field name/type欠落を検出する
- [x] store kind欠落を検出する
- [x] param name/model欠落を検出する
- [x] return name/model欠落を検出する
- [x] initialized store name/model欠落を検出する
- [x] unit testで `missing_required_field` を固定する
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

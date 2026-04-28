# Go M5 ER 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: Go実装 Milestone 5 第3段の ER renderer vertical slice 実装チェックリスト

---

## 1. 目的

M5第3段では、ER Diagram rendererをUC-001の横断view YAMLで縦切り実装する。

対象:

```text
views/er.yaml -> renders/_cross/er.md
```

まずは `as: er_diagram` の decode / semantic view build / ER render / golden test 1本までを通す。
`render/er` は `semantic.Project` のみを読み、Raw YAML structs には依存しない。

---

## 2. 実装範囲

- `as: er_diagram` view file decode
- semantic.ERView / ERViewModule
- view YAML の `modules[].module` による対象module列挙
- `store.kind: db` のみ対象化
- `store.of` から model fields を辿る
- primitive型のER型マッピング
- fkありfieldの `FK` 表示
- view内modelへのFK relation描画
- view外modelへのFKはrelationなし
- UC-001 golden test 1本

---

## 3. 境界

```text
source    -> rawyaml
resolve   -> rawyaml, semantic
render/er -> semantic
```

禁止:

```text
render/er -> rawyaml
render/er内でview YAMLを直接読む
Wireframe / API Table を混ぜる
```

---

## 4. 受け入れ条件

- [x] `views/er.yaml` をdecodeできる
- [x] `semantic.Project` にER viewを保持できる
- [x] `ec_er` をrenderできる
- [x] `store.kind: db` のみER対象にできる
- [x] `store.of` のmodel fieldsをカラム化できる
- [x] cross module FKをrelationとして描画できる
- [x] view外FKをrelationなしで扱える
- [x] `render/er` が `rawyaml` をimportしていない
- [x] golden test 1本が通る
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る

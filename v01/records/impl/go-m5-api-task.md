# Go M5 API Table 実装タスク

- **status**: completed
- **last_updated**: 2026-04-28
- **scope**: Go実装 Milestone 5 第4段の API Table renderer vertical slice 実装チェックリスト

---

## 1. 目的

M5第4段では、API Table rendererをUC-001の横断view YAMLで縦切り実装する。

対象:

```text
views/api_table.yaml -> renders/_cross/api.md
```

まずは `as: api_table` の decode / semantic view build / API Table render / golden test 1本までを通す。
`render/api` は `semantic.Project` のみを読み、Raw YAML structs には依存しない。

---

## 2. 実装範囲

- `as: api_table` view file decode
- semantic.APIView / APIViewModule
- view YAML の `modules[].module` と `include_submodules` による対象module列挙
- endpoint taskのみ対象化
- `http_root_path` + section起点相対module path + task leaf path によるroute合成
- params / returns table render
- section内 task id ASCII昇順sort
- UC-001 golden test 1本

---

## 3. 境界

```text
source     -> rawyaml
resolve    -> rawyaml, semantic
render/api -> semantic
```

禁止:

```text
render/api -> rawyaml
render/api内でview YAMLを直接読む
Wireframe / ER rendererを混ぜる
```

---

## 4. 受け入れ条件

- [x] `views/api_table.yaml` をdecodeできる
- [x] `semantic.Project` にAPI viewを保持できる
- [x] `ec_api` をrenderできる
- [x] `endpoint: true` taskのみAPI Table対象にできる
- [x] `include_submodules` に従って対象taskを収集できる
- [x] full routeを合成できる
- [x] params / returnsをtable表示できる
- [x] `render/api` が `rawyaml` をimportしていない
- [x] golden test 1本が通る
- [x] `go fmt ./...` が通る
- [x] `go test ./...` が通る

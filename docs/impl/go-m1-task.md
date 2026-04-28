# Go M1 実装タスク

- **status**: completed
- **last_updated**: 2026-04-27
- **scope**: Go実装 Milestone 1 の実装チェックリスト（完了済み）

---

## 1. 目的

このファイルは、`docs/impl/go-skeleton.md` に従って Milestone 1 を実装するための作業チェックリストである。

M1の最終ゴールは、UC-001のログインDAG 1本を以下の経路で通し、golden fixtureと一致させること。

```text
docs/uc/001-ec-checkout-flow/yaml/auth/task/login.yaml
  ↓ source load / classify / decode
Raw YAML structs
  ↓ resolve / semantic build
ResolvedProject
  ↓ DAG renderer
docs/uc/001-ec-checkout-flow/renders/auth/dag-login.md と比較
```

### M1完了メモ

M1は実装完了済みとして扱う。

- `docs/TASKS.md` の Milestone 1 は完了済み。
- UC-001 の `auth.task.login` は Raw YAML → ResolvedProject → DAG renderer → golden test まで到達済み。
- `go test ./...` 通過は `docs/TASKS.md` に記録済み。
- このファイル内の詳細タスク一覧は、M1実装前チェックリストとして残す。
- 完了判定は「7. 受け入れ条件」の完了済みサマリを正とする。

---

## 2. 前提ドキュメント

実装前に最低限読むこと。

- `docs/doc-policy.md`
- `docs/impl/go-skeleton.md`
- `docs/TASKS.md` の Milestone 1
- `docs/adr/014-initializes-field.md`
- `docs/adr/027-module-nesting-and-name-resolution.md`
- `docs/adr/030-yaml-file-type-declaration.md`
- `docs/adr/031-actor-global-definition.md`
- `docs/adr/047-go-semantic-model-query-layer-boundary.md`
- `docs/adr/048-resolved-project-index-strategy.md`
- `docs/spec/nodes.md` の task / model / store / actor 範囲
- `docs/spec/views/dag.md`
- `docs/uc/001-ec-checkout-flow/yaml/auth/task/login.yaml`
- `docs/uc/001-ec-checkout-flow/renders/auth/dag-login.md`

---

## 3. 実装原則

- M1は縦切りを優先する。
- Raw YAML structs と ResolvedProject を分ける。
- rendererは `rawyaml` をimportしない。
- QueryService / MCP / CLI command設計はM1では作らない。
- view file / render_index.yaml は分類だけ行い、M1ではdecodeしない。
- `yaml.Node` を使う場合も `internal/source` 内部に閉じる。
- `rawyaml` / `semantic` / `resolve` / `render` は `yaml.Node` に依存しない。
- `go.mod` のmodule pathは `github.com/hiroshiasayadev-prog/brewprint` とする。
- YAML decoderは `gopkg.in/yaml.v3` を使う。

---

## 4. 作るディレクトリ / ファイル

```text
brewprint/
  go.mod

  cmd/
    brewprint/
      main.go

  internal/
    source/
      loader.go
      classify.go
      file.go

    rawyaml/
      node.go
      task.go
      model.go
      store.go
      actor.go

    semantic/
      id.go
      project.go
      node.go
      task.go
      model.go
      store.go
      actor.go
      asset.go
      diagnostic.go

    resolve/
      builder.go
      symbols.go
      names.go

    render/
      dag/
        renderer.go
        viewmodel.go
        renderer_test.go

    testutil/
      golden/
        golden.go
```

M1では以下を作らない。

```text
internal/cli/
internal/query/
internal/golden/
rawyaml/view.go
rawyaml/render_index.go
```

---

## 5. タスク一覧

### 5.1 Go skeleton

- [ ] `go.mod` を作る
  - module path: `github.com/hiroshiasayadev-prog/brewprint`
  - dependency: `gopkg.in/yaml.v3`
- [ ] `cmd/brewprint/main.go` を空殻で作る
  - M1ではCLI引数設計をしない
  - build可能な最小mainでよい
- [ ] `internal/` 配下のpackageディレクトリを作る

完了条件:

- [ ] `go test ./...` が実行できる
- [ ] まだ実装が空でもpackage構成がコンパイルを阻害しない

---

### 5.2 `internal/source`

実装対象:

- [ ] yaml rootを再帰走査するLoaderを作る
- [ ] `.yaml` / `.yml` を対象にする
- [ ] FileIDをyaml rootからの相対slash pathとして正規化する
- [ ] `render_index.yaml` を検出する
- [ ] top-level `as:` があればview fileとして分類する
- [ ] top-level `nodes:` があればnode fileとして分類する
- [ ] node fileを `rawyaml` structへdecodeする
- [ ] view file / render_index.yaml は分類だけしてsemantic build対象からskipする

M1で扱う分類:

| file | 判定 |
|---|---|
| top-level `as:` あり | view file |
| top-level `nodes:` あり | node file |
| file名 `render_index.yaml` | render index file |
| その他 | unsupported diagnostic またはerror |

FileID / module path例:

```text
auth/task/login.yaml
  FileID: auth/task/login.yaml
  module: auth
  qid: auth.task.login

payment/webhooks/task/process_payment.yaml
  FileID: payment/webhooks/task/process_payment.yaml
  module: payment.webhooks
  qid: payment.webhooks.task.process_payment

actors.yaml
  actor: project global
  external reference form: end_user / stripe
  internal qid form: actor.end_user / actor.stripe
```

完了条件:

- [ ] UC-001 yaml rootをloadできる
- [ ] `auth/task/login.yaml` をnode fileとしてdecodeできる
- [ ] `views/api_table.yaml` 等をview fileとしてskipできる
- [ ] `render_index.yaml` をskipできる

---

### 5.3 `internal/rawyaml`

実装対象:

- [ ] node file root struct
- [ ] common node fields
- [ ] task struct
- [ ] model struct
- [ ] store struct
- [ ] actor struct

M1で扱うtask field:

- `id`
- `type`
- `main`
- `note`
- `params`
- `returns`
- `reads`
- `writes`
- `initializes`
- `endpoint`
- `method`
- `path`

M1で扱うmodel field:

- `id`
- `type`
- `kind`
- `fields`
- `note`

M1で扱うstore field:

- `id`
- `type`
- `kind`
- `of`
- `note`

M1で扱うactor field:

- `id`
- `type`
- `note`

禁止:

- [ ] `rawyaml` に `*yaml.Node` を保持しない
- [ ] `rawyaml` でQualifiedID生成をしない
- [ ] `rawyaml` で参照解決をしない
- [ ] `rawyaml` でimplicit asset生成をしない

完了条件:

- [ ] UC-001の task / model / store / actor をdecodeできる
- [ ] 未対応node種別はM1でsemantic build対象外またはunsupportedとして扱える

---

### 5.4 `internal/semantic`

実装対象:

- [ ] `FileID`
- [ ] `QualifiedID`
- [ ] node kind
- [ ] resolved task
- [ ] resolved model
- [ ] resolved store
- [ ] resolved actor
- [ ] implicit asset
- [ ] diagnostic
- [ ] `Project` index

M1の `Project` field:

```go
type Project struct {
    NodesByQID     map[QualifiedID]Node
    NodesByFile    map[FileID][]Node
    MainNodeByFile map[FileID]QualifiedID

    TasksByQID  map[QualifiedID]*Task
    ModelsByQID map[QualifiedID]*Model

    StoresByQID       map[QualifiedID]*Store
    StoresByFileLocal map[FileID]map[string]*Store

    ActorsByQID map[QualifiedID]*Actor

    TasksReadingStore map[QualifiedID][]QualifiedID
    TasksWritingStore map[QualifiedID][]QualifiedID
}
```

M1では作らないfield:

- `ReferencesBySource`
- `ReferencesByTarget`
- `TransitionsByStateEventGuard`
- `ActionsByTask`
- `ScenariosByID`
- `FilesByTopLevelModule`

完了条件:

- [ ] rendererが `semantic.Project` だけを読んでDAGを生成できる
- [ ] `semantic` が `rawyaml` に依存しない

---

### 5.5 `internal/resolve`

実装対象:

- [ ] file pathからmodule pathを導出する
- [ ] node kind sentinelに基づきQualifiedIDを作る
- [ ] actorはproject globalとして扱う
- [ ] actor内部QIDは `actor.<id>` とする
- [ ] symbol tableを作る
- [ ] duplicate QIDを検出する
- [ ] duplicate actor IDを検出する
- [ ] `NodesByQID` を構築する
- [ ] `NodesByFile` を構築する
- [ ] `MainNodeByFile` を構築する
- [ ] `TasksByQID` / `ModelsByQID` / `StoresByQID` / `ActorsByQID` を構築する
- [ ] task paramsをmodel参照へ解決する
- [ ] task returnsをmodel参照へ解決する
- [ ] store.ofをmodel参照へ解決する
- [ ] task returnsからimplicit assetを生成する
- [ ] task initializesからfile-private storeを生成する
- [ ] `StoresByFileLocal` を構築する
- [ ] task reads / writesをstore参照へ解決する
- [ ] `TasksReadingStore` / `TasksWritingStore` を構築する

store参照解決順:

```text
reads / writes store resolution:
1. 同一FileIDの StoresByFileLocal[name]
2. 同一moduleの store.<name>
3. フルパス QualifiedID
```

名前解決のM1範囲:

- [ ] 同一module内ID直書き
- [ ] クロスモジュールフルパス
- [ ] module nesting
- [ ] actor global ID

完了条件:

- [ ] `auth.task.login` がResolvedProject上のtaskとして取得できる
- [ ] `form` param が `auth.model.login_form` に解決される
- [ ] `auth_token` returns が `auth.model.token` に解決される
- [ ] `user_db` / `request_context_store` reads がstoreに解決される
- [ ] `session_store` writes がstoreに解決される
- [ ] `login_log_db` writes がfile-private initialized storeに解決される
- [ ] `login_log_db` が外部storeとして `StoresByQID` に混入しない

---

### 5.6 `internal/render/dag`

実装対象:

- [ ] ResolvedProjectからmain taskを取得する
- [ ] task noteをMarkdown本文に出力する
- [ ] endpoint taskなら `**API**` 行を出力する
- [ ] params boundaryを描画する
- [ ] returns boundaryを描画する
- [ ] read store edgeを描画する
- [ ] write store edgeを描画する
- [ ] implicit asset / returns edgeを描画する
- [ ] Mermaid flowchartを出力する
- [ ] Tasks detail sectionを出力する

M1のAPI行:

```text
**API**: [POST /api/login](../_cross/api.md)
```

M1ではAPI Table viewをdecodeしない。
DAG冒頭のAPI行は、taskの `endpoint` / `method` / `path` だけを使い、UC-001 fixtureに合わせる。

禁止:

- [ ] `internal/render/dag` から `internal/rawyaml` をimportしない
- [ ] renderer内で名前解決をしない
- [ ] renderer内でreverse lookup indexを再構築しない
- [ ] renderer内でRaw YAMLを走査しない

完了条件:

- [ ] `docs/uc/001-ec-checkout-flow/renders/auth/dag-login.md` 相当のMarkdownを生成できる
- [ ] Mermaidのstore edge labelが `read` / `write` になる
- [ ] Task detail sectionがgoldenと一致する

---

### 5.7 `internal/testutil/golden`

実装対象:

- [ ] expected fileを読む
- [ ] actualと比較する
- [ ] CRLFをLFへ正規化する
- [ ] expected / actual の末尾改行を1つに揃える
- [ ] 差分をテスト失敗時に読みやすく出す

M1ではやらないこと:

- [ ] trailing spacesを勝手にtrimしない
- [ ] Markdown内部の空行を勝手に正規化しない

完了条件:

- [ ] golden helperをrenderer testから使える
- [ ] 差分が出た時にexpected / actualの違いを把握できる

---

### 5.8 Golden test

実装対象:

- [ ] `internal/render/dag/renderer_test.go` を作る
- [ ] UC-001 yaml rootをloadする
- [ ] ResolvedProjectをbuildする
- [ ] `auth/task/login.yaml` のmain taskをrender対象にする
- [ ] actual Markdownを生成する
- [ ] `docs/uc/001-ec-checkout-flow/renders/auth/dag-login.md` と比較する

完了条件:

- [ ] `go test ./...` が通る
- [ ] `dag-login.md` golden testが通る

---

## 6. M1で実装しないもの

- CLI command設計
- `brewprint render` コマンド
- QueryService
- MCP server wrapper
- view file decode
- render_index.yaml decode / validation
- API Table renderer
- ER renderer
- State Diagram renderer
- Sequence Diagram renderer
- Wireframe renderer
- flow
- branch
- fork
- join
- foreach
- state
- event
- scenario
- output placement / renders index生成
- import boundary lint導入

---

## 7. 受け入れ条件

M1完了条件:

- [x] `go test ./...` が通る
- [x] UC-001 yaml rootを読み込める
- [x] node file / view file / render_index.yaml を分類できる
- [x] view file / render_index.yaml をM1対象外としてskipできる
- [x] `auth.task.login` をRaw YAMLからResolvedProjectへbuildできる
- [x] `auth.task.login` のparams / returns / reads / writes / initializesを解決できる
- [x] task returnsからimplicit assetを生成できる
- [x] task initializesからfile-private storeを生成できる
- [x] store read/write indexを構築できる
- [x] DAG rendererが `rawyaml` をimportしていない
- [x] `docs/uc/001-ec-checkout-flow/renders/auth/dag-login.md` とgolden一致する

---

## 8. 推奨実装順

1. `go.mod` / `cmd/brewprint/main.go`
2. `internal/source` のwalk / classify
3. `internal/rawyaml` のstruct decode
4. `internal/semantic` のID / Project / node型
5. `internal/resolve` のsymbol table / QID生成
6. model / store / task / actor index構築
7. task params / returns / store.of 解決
8. task returns由来のimplicit asset生成
9. task initializes由来のfile-private store生成
10. reads / writes解決とstore access index構築
11. `internal/render/dag` のMarkdown生成
12. `internal/testutil/golden`
13. `internal/render/dag/renderer_test.go`
14. `go test ./...`

---

## 9. 実装後に確認すること

- [ ] `go fmt ./...`
- [ ] `go test ./...`
- [ ] `git diff` でdocs以外の不要変更がないこと
- [ ] `render/dag` が `rawyaml` をimportしていないこと
- [ ] `semantic` が `rawyaml` をimportしていないこと
- [ ] `rawyaml` に `*yaml.Node` が露出していないこと
- [ ] M1完了後、`docs/TASKS.md` のMilestone 1進捗更新を検討する

---

## 10. 実装者向けメモ

M1では「きれいな全体実装」よりも、`auth.task.login` の1本を壊さず通すことを優先する。

ただし、以下の境界は崩さない。

```text
source      -> rawyaml
resolve     -> rawyaml, semantic
render/dag  -> semantic
```

禁止:

```text
render/dag -> rawyaml
semantic   -> rawyaml
rawyaml    -> yaml.Node
```

この境界を守れていれば、M2以降でDAG対象を広げても手戻りは小さい。

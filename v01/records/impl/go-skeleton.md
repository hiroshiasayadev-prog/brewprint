# Go skeleton 方針

- **status**: draft
- **last_updated**: 2026-04-27
- **scope**: Go実装 Milestone 1 の初期package構成と責務境界

---

## 1. 目的

このメモは、`docs/TASKS.md` の Milestone 1「Go project skeleton を作る」に入る前に、初期Go実装のpackage構成・責務境界・M1で作るもの / 作らないものを固定するための作業メモである。

設計判断の根拠は主に以下に従う。

- V01-ADR-014: initializesフィールド設計
- V01-ADR-027: モジュールネスト許容と名前解決ルール拡張
- V01-ADR-030: YAMLファイル種別の自己宣言
- V01-ADR-031: actorのグローバル定義
- V01-ADR-047: Go semantic model / query layer boundary
- V01-ADR-048: ResolvedProject index strategy
- `docs/TASKS.md` Milestone 1

このメモはADRではない。
M1実装開始時のpackage方針を残すための実装メモとして扱う。

---

## 2. M1のゴール

M1では、UC-001のDAG 1本を縦に通す。

```text
UC-001 yaml/auth/task/login.yaml
  ↓ load / classify / decode
Raw YAML structs
  ↓ semantic build
ResolvedProject
  ↓ DAG renderer
renders/auth/dag-login.md と golden 比較
```

最初の成功条件は、以下である。

- `auth.task.login` をRaw YAMLから読み込める
- 必要なmodel / store / actorをsymbol tableへ登録できる
- task params / returns / reads / writes / initializes を解決できる
- task returnsからimplicit assetを生成できる
- task initializesからfile-private storeを構築できる
- store read/write indexを構築できる
- endpoint taskについてDAG冒頭の `**API**` 行を出力できる
- DAG rendererがRaw YAML structsを直接読まず、ResolvedProjectから `dag-login.md` 相当を生成できる
- golden fixture `docs/uc/001-ec-checkout-flow/renders/auth/dag-login.md` と比較できる

---

## 3. 基本方針

### 3.1 vertical sliceを優先する

M1では、全仕様を横断的に実装しない。

まず以下の1本を通す。

```text
Raw YAML → ResolvedProject → DAG renderer → golden test
```

MCP / QueryService / State / Sequence / ER / API / Wireframe は後続milestoneで扱う。

### 3.2 Raw YAMLとResolvedProjectを分ける

V01-ADR-047に従い、Go実装は以下の境界を守る。

```text
YAML files
  ↓ load / classify / decode
Raw YAML structs
  ↓ validate / name resolution / derived model build
ResolvedProject
  ↓
QueryService / Renderer
```

Raw YAML structsは、YAML decode後の入力形を保持するだけにする。
名前解決、implicit asset生成、index構築はResolvedProject buildで行う。

### 3.3 rendererはrawyamlをimportしない

DAG rendererは `rawyaml` packageを直接importしない。

rendererが読む入力は以下のいずれかとする。

1. `semantic.Project` / ResolvedProjectのindex・resolved field
2. ResolvedProjectから構築したDAG用view model

renderer内でRaw YAML走査、未解決文字列の再解釈、reverse lookupの再構築をしない。

### 3.4 M1ではCLIを広げない

M1では `cmd/brewprint/main.go` は空殻でよい。

DAG vertical sliceはgolden testから直接動かす。
CLIの引数設計や `brewprint render` コマンド設計は、render対象・出力配置が増える後続milestoneで扱う。

### 3.5 Go module / YAML decoder

`go.mod` のmodule pathは、docsも含むこのrepo全体をGo module rootとして扱い、以下を採用する。

```go
module github.com/hiroshiasayadev-prog/brewprint
```

UC fixtureやdocsは同一repo内のtestdata相当の入力として扱う。
Go packageとしてimportしないため、docsが同一repoに存在していてもmodule path上の問題はない。

YAML decoderはM1では以下を採用する。

```go
gopkg.in/yaml.v3
```

M1では通常のstruct decodeを基本とする。
`yaml.Node` はline/column等のsource locationや柔軟な未知フィールド検出が必要になった時点で使う。
M1のDAG 1本では、全ファイルを `yaml.Node` として保持する必要はない。

`yaml.Node` を使う場合も、利用範囲は `internal/source` 内部に閉じる。
`rawyaml` / `semantic` / `resolve` / `render` は `yaml.Node` に依存しない。

禁止例:

```go
type Task struct {
    Node *yaml.Node
}
```

source locationが必要になった場合は、`source` が `yaml.Node.Line` / `yaml.Node.Column` を読み取り、`rawyaml.SourceLoc` または `semantic.Diagnostic` に写像する。

例:

```go
type SourceLoc struct {
    File   string
    Line   int
    Column int
}
```

この方針により、YAML decode実装をstruct decodeから `yaml.Node` 併用へ変更しても、semantic build / renderer / query layer には影響を広げない。

---

## 4. 初期ディレクトリ構成

M1の初期構成は以下とする。

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

M1では以下をまだ作らない。

```text
internal/cli/
internal/query/
internal/golden/
rawyaml/view.go
rawyaml/render_index.go
```

理由:

- `internal/cli/` はCLI引数設計が固まってから切る
- `internal/query/` はM3 QueryService vertical sliceで作る
- `internal/golden/` ではなく、`internal/testutil/golden` をhelperとして使う
- view / render_index のraw structはM1のDAG 1本には不要

---

## 5. package責務

### 5.1 `internal/source`

`source` は、YAMLファイルのload / classify / decodeの入口を担う。

担当:

- yaml rootの走査
- FileIDのslash正規化
- top-level `as:` の有無によるview file判定
- top-level `nodes:` の有無によるnode file判定
- `render_index.yaml` の検出
- node fileを `rawyaml` structへdecode
- view file / render_index.yaml はM1では検出のみ行い、semantic build対象からはskipする

`source` が行うのはdecode-time validationまでである。
名前解決、semantic validation、derived model buildは行わない。

FileID / module path導出例:

```text
yaml root:
  docs/uc/001-ec-checkout-flow/yaml

auth/task/login.yaml
  FileID: auth/task/login.yaml
  module: auth
  node kind: task
  qid: auth.task.login

auth/model/token.yaml
  FileID: auth/model/token.yaml
  module: auth
  node kind: model
  qid: auth.model.token

payment/webhooks/task/process_payment.yaml
  FileID: payment/webhooks/task/process_payment.yaml
  module: payment.webhooks
  node kind: task
  qid: payment.webhooks.task.process_payment

actors.yaml
  FileID: actors.yaml
  actor: project global
  external reference form: end_user / stripe
  internal qid form: actor.end_user / actor.stripe
```

actorはV01-ADR-031に従い、YAML上の参照形式はglobal ID直参照とする。
内部index用の `QualifiedID` は、通常nodeと衝突しないよう `actor.<id>` 形式で保持する。

### 5.2 `internal/rawyaml`

`rawyaml` はYAML decode用structを持つ。

M1で扱うnode種別:

- task
- model
- store
- actor

M1で扱うtask field:

- `id`
- `type`
- `note`
- `params`
- `returns`
- `reads`
- `writes`
- `initializes`
- `endpoint`
- `method`
- `path`

`rawyaml` は以下をしない。

- QualifiedID生成
- 参照解決
- actor global解決
- implicit asset生成
- index構築
- renderer向け整形

### 5.3 `internal/semantic`

`semantic` はResolvedProject側の型を持つ。

担当:

- `QualifiedID`
- `FileID`
- resolved node型
- task / model / store / actor
- implicit asset
- diagnostics
- ResolvedProject index

M1の `Project` は、以下のfieldを持つ。

```go
type Project struct {
    NodesByQID     map[QualifiedID]Node
    NodesByFile    map[FileID][]Node
    MainNodeByFile map[FileID]QualifiedID

    TasksByQID  map[QualifiedID]*Task
    ModelsByQID map[QualifiedID]*Model
    StoresByQID       map[QualifiedID]*Store
    StoresByFileLocal map[FileID]map[string]*Store
    ActorsByQID       map[QualifiedID]*Actor

    TasksReadingStore map[QualifiedID][]QualifiedID
    TasksWritingStore map[QualifiedID][]QualifiedID
}
```

M1では以下のfieldは作らない。

```text
ReferencesBySource
ReferencesByTarget
TransitionsByStateEventGuard
ActionsByTask
ScenariosByID
FilesByTopLevelModule
```

これらは後続milestoneで追加する。

- `ReferencesBySource` / `ReferencesByTarget`: M3 QueryService
- `TransitionsByStateEventGuard`: State / Sequence
- `ActionsByTask`: State / Sequence / QueryService
- `ScenariosByID`: Sequence
- `FilesByTopLevelModule`: M4 render_index / output placement

### 5.4 `internal/resolve`

`resolve` はRaw YAML structsからResolvedProjectを構築する。

担当:

- file pathからmodule pathを導出する
- node kind sentinelに基づきQualifiedIDを作る
- actorはV01-ADR-031に従いglobal IDとして扱う
- symbol tableを作る
- `NodesByQID` / `NodesByFile` / `MainNodeByFile` を構築する
- task params / returns / reads / writes を解決する
- store.of を解決する
- task returnsからimplicit assetを生成する
- task initializesからfile-private storeを構築し、`StoresByFileLocal` に登録する
- initialized storeを同一ファイル内のwrites解決対象に含める
- `TasksReadingStore` / `TasksWritingStore` を構築する
- duplicate actor IDを検出する

M1のstore参照解決順は以下とする。

```text
reads / writes store resolution:
1. 同一FileIDの StoresByFileLocal[name]
2. 同一moduleの store.<name>
3. フルパス QualifiedID
```

`initializes` 由来のfile-private storeは外部ファイルから参照できない。
M1では少なくとも `writes` の解決対象に含める。
`reads` でfile-private storeを読むケースも、同一ファイル内であれば同じ解決順に従って許容する。

`resolve` はrenderer向けのMermaid / Markdown整形をしない。

### 5.5 `internal/render/dag`

`render/dag` はDAG rendererを持つ。

担当:

- ResolvedProjectからmain taskを取得する
- endpoint taskについてDAG冒頭の `**API**` 行を出力する
- task params / returns boundaryを描画材料に変換する
- implicit assetを描画材料に変換する
- reads / writes store edgeを描画材料に変換する
- task detail sectionを出力する
- markdownを生成する
- golden testでfixtureと比較する

M1ではAPI Table viewをdecodeしない。
そのためDAG冒頭の `**API**` 行は、taskの `endpoint` / `method` / `path` だけを使い、UC-001のgolden fixtureに合わせて `/api/<leaf path>` と `../_cross/api.md` を出力する。
V01-ADR-028に基づく正式なroute合成は、M5のAPI Table rendererで扱う。

`render/dag` は `rawyaml` をimportしない。

### 5.6 `internal/testutil/golden`

`testutil/golden` はgolden test用helperを持つ。

担当:

- expected fileの読み込み
- actualとの比較
- 差分表示補助
- 改行コードをCRLFからLFへ正規化する
- expected / actual の末尾改行を1つに揃える

M1ではtrailing spacesやMarkdown内部の空行は意味のある出力差分として扱い、勝手にtrimしない。

rendererごとのテスト本体は、各renderer package側に置く。

例:

```text
internal/render/dag/renderer_test.go
```

---

## 6. classify / decode方針

M1では `source` がload / classify / decodeの入口になる。

```text
source.Loader
  ↓ walk yaml root
  ↓ normalize FileID
  ↓ classify file
  ↓ decode node files into rawyaml structs
  ↓ return loaded raw project files
```

分類ルール:

| file | 判定 |
|---|---|
| top-levelに `as:` がある | view file |
| top-levelに `nodes:` がある | node file |
| file名が `render_index.yaml` | render index file |
| それ以外 | decode-time error またはM1ではunsupported diagnostic |

M1での扱い:

| file種別 | M1での扱い |
|---|---|
| node file | decodeしてsemantic build対象にする |
| view file | classifyだけしてskip |
| render_index.yaml | classifyだけしてskip |

---

## 7. actor方針

M1でも `actors.yaml` はloadする。

理由:

- UC-001のyaml rootを走査すると `actors.yaml` が存在する
- actorはV01-ADR-031によりproject globalである
- V01-ADR-048のbase indexに `actorsByQID` が含まれる

M1で行うこと:

- actor nodeをdecodeする
- `ActorsByQID` に登録する
- actor ID重複をsemantic validation errorにする

M1で行わないこと:

- actorをSequence Diagram participantとして使う
- event.actor参照を解決する

---

## 8. M1で扱うYAML範囲

M1で実装するもの:

- node file decode
- task
- model
- store
- actor
- task params
- task returns
- task reads
- task writes
- task initializes
- endpoint / method / path
- store.of
- implicit asset
- file-private initialized store
- store read/write index
- DAG rendererの `**API**` 行出力

M1で実装しないもの:

- view file decode
- render_index.yaml decode / validation
- flow
- branch
- fork
- join
- foreach
- state
- event
- sequence scenario
- api table
- er view
- wireframe
- QueryService
- MCP server wrapper
- CLI command design

---

## 9. import boundary

M1では外部lintを導入しない。
ただし、以下のimport境界をレビュー基準として守る。

```text
source      -> rawyaml
resolve     -> rawyaml, semantic
render/dag  -> semantic
query       -> semantic   # M3以降
mcp wrapper -> query      # M6以降
```

禁止:

```text
render/dag -> rawyaml
query      -> rawyaml
mcp        -> rawyaml
```

将来、境界違反が増えそうな場合は、go-arch-lint等の導入を検討する。

---

## 10. M1実装順

推奨順:

1. `go.mod` を作る
2. `cmd/brewprint/main.go` を空殻で作る
3. `internal/source` のfile walk / classifyを作る
4. `internal/rawyaml` のtask / model / store / actor structを作る
5. `internal/semantic` のID / Project / node型を作る
6. `internal/resolve` のsymbol table / minimal builderを作る
7. task initializesからfile-private storeを構築する
8. `TasksReadingStore` / `TasksWritingStore` を構築する
9. endpoint taskの `**API**` 行をDAG rendererで出力する
10. `internal/render/dag` のrendererを作る
11. `internal/testutil/golden` を作る
12. `internal/render/dag/renderer_test.go` で `dag-login.md` と比較する

---

## 11. 判断保留事項

以下はM1では決めない。

- CLI command構成
- public Go APIの安定化
- QueryService interface
- MCP wrapper package構成
- view-specific view modelの最終形
- all renderer共通interface
- render_indexによる出力配置
- import boundary lint導入

M1で実装中に必要になった場合は、TASKS更新または別メモ / ADRで扱う。

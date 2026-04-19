# TASKS

brewprintの積みタスク一覧。会話をまたいでコンテキストを引き継ぐために使う。

---

## ステータス凡例

- `[ ]` 未着手
- `[~]` 進行中
- `[x]` 完了

---

- [x] **YAMLの階層構造とディレクトリ構造の1:1対応設計 → ADR-010**
  - CA強制・1ファイル=1ノード・model/asset分離・ビュー自動導出を確定
  - ADR-007（asset定義）・ADR-002（ビュー別ファイル分け）を部分supersede

- [x] **1ファイル=1メインノード + サブノード構造 → ADR-011**
  - `main: true` フラグ・サブノードのprivateスコープを確定
  - ADR-010の「1ファイル=1ノード」を「1ファイル=1メインノード」に改訂

- [x] **foreach / cond / initializes の設計 → ADR-012, 013（→016), 014**
  - `branch`（旧`cond`）/ `fork` / `join` のノード種別を確定（ADR-012）
  - `foreach` はnode typeから廃止し `flow:` の制御構文として設計、`$item` シジル導入（ADR-013 superseded by ADR-016）
  - `initializes` をmain nodeのフィールドとして設計（ADR-014）

- [x] **ファイル内edgeの記述構造（flow:セクション） → ADR-015**
  - tasks=signatureのみ・wiringはflow:に分離
  - fork/join記法・$paramsシジル・reads/writesフィールドを確定

- [ ] **eventノードのスキーマを決める → ADR-017**
  - 制御フローの起点として `event` ノードをDAGに導入する方向は確定
  - `source` 属性（ui / time / external / er）でタグ付け
  - 具体的なスキーマが未定（spec/overview.md `open_issues` より）

- [ ] **Edgeの管理方式を決める → ADR-018**
  - クロスエッジに `kind` 属性が必要なため、Nodeのadjacency listではなく別管理が有力
  - 未決（spec/overview.md `open_issues` より）

---

## ユースケース

- [ ] **UC-001: login flowのユースケースを書く**
  - `docs/uc/001-login-flow.md` を新規作成
  - sequence → DAG → ER が全部繋がるのを1ユースケースで通しで確認する
  - actor / endpoint / store(kind=db) の実際のYAML記述を確認する（store=旧state）
  - ADR-012/013（foreach/cond/edge）確定後に着手推奨

---

- [ ] **spec/nodes.md を新規作成**
  - 各ノード種別（task / asset / store / actor）のフィールド定義を網羅
  - endpointフラグ含む（ADR 006-008に基づく）

- [ ] **spec/views.md を新規作成**
  - 各図（DAG / ER / state / class / sequence）のrenderルールを定義
  - どのノード・エッジをどの図に出力するかのマッピング

- [ ] **spec/edges.md を新規作成**
  - クロスエッジの種類（write / read / trigger / reflect / hydrate）の定義
  - Edgeの管理方式が決まり次第着手（ADR 007に依存）

---

## 実装

- [ ] **GoのAST struct定義**
  - ADR 001に基づき、ノード種別ごとにstructを用意
  - eventノードのスキーマ確定後に着手推奨

- [ ] **YAMLパーサー実装**
  - AST構築 + 名前解決（ADR 003）
  - エラーはASTビルド時に返す

- [ ] **MCPツール実装**
  - `get_signature` / `get_deps` / `inspect`

- [ ] **Mermaidレンダラー実装**
  - 各view（DAG / ER / state / class / sequence）のrenderロジック

# TASKS

brewprintの積みタスク一覧。会話をまたいでコンテキストを引き継ぐために使う。

---

## ステータス凡例

- `[ ]` 未着手
- `[~]` 進行中
- `[x]` 完了

---

- [ ] **YAMLの階層構造とディレクトリ構造の1:1対応設計 → ADR 010**
  - 「設計と実装の乖離を構造的に防ぐ」ための仕組み
  - YAMLの階層がそのままディレクトリ構造に対応するスキーマ設計が必要
  - スキーマ設計に大きく影響するため独立したセッションで議論



- [ ] **UC-001: login flowのユースケースを書く**
  - `docs/uc/001-login-flow.md` を新規作成
  - sequence → DAG → ER が全部繋がるのを1ユースケースで通しで確認する
  - actor / endpoint / state(kind=db) の実際のYAML記述を確認する

- [ ] **eventノードのスキーマを決める → ADR 009**
  - 制御フローの起点として `event` ノードをDAGに導入する方向は確定
  - `source` 属性（ui / time / external / er）でタグ付け
  - 具体的なスキーマが未定（spec/overview.md `open_issues` より）

- [ ] **Edgeの管理方式を決める → ADR 010**
  - クロスエッジに `kind` 属性が必要なため、Nodeのadjacency listではなく別管理が有力
  - 未決（spec/overview.md `open_issues` より）

---

## spec整備

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

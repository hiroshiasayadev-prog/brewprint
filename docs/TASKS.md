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

- [x] **eventノードのスキーマを決める → ADR-018**

- [x] **stateノード（FSM）の設計 → ADR-019**
  - `state` ノード種別・`transitions:` セクションを確定
  - `store`（データ保持）との概念区別を明記
  - 制御フローの起点として `event` ノードをDAGに導入する方向は確定
  - `source` 属性（ui / time / external / er）でタグ付け
  - 具体的なスキーマが未定（spec/overview.md `open_issues` より）

- [x] **Edgeの管理方式を決める → ADR-020**
  - クロスエッジを `write` / `read` の2種に絞り、task nodeの `reads`/`writes` フィールドとして記述
  - `trigger`/`reflect`/`hydrate` は既存のevent/transition構造で表現済みのため廃止
  - ADR-015の `reads`/`writes`（flow:ステップ記述）をsupersede

---

## ユースケース

- [ ] **UC-001: login flowのユースケースを書く**
  - `docs/uc/001-login-flow.md` を新規作成
  - sequence → DAG → ER が全部繋がるのを1ユースケースで通しで確認する
  - actor / endpoint / store(kind=db) の実際のYAML記述を確認する（store=旧state）
  - ADR-012/013（foreach/cond/edge）確定後に着手推奨

---

- [x] **spec/nodes.md を新規作成**
  - 全ノード種別のフィールド定義を網羅（ADR-001〜021に基づく）

- [x] **spec/edges.md を新規作成**
  - flow: / transitions: / reads/writes / $シジル体系を定義（ADR-015, 016, 019, 020に基づく）

- [~] **spec/views/ ディレクトリを作成し図ごとにspecを書く**
  - [x] `spec/views/dag.md` — DAGのrenderルール（ノード・エッジのマッピング）　status: confirmed
  - [x] `spec/views/er.md` — ER図のrenderルール
  - [ ] `spec/views/state.md` — State Diagramのrenderルール
  - [ ] `spec/views/sequence.md` — Sequence Diagramのrenderルール
  - [ ] `spec/views/api-table.md` — API Table（list_endpoints MCPツール出力）のルール
  - 1ファイル = 1図。各ファイルにFront Matter（doc-policy準拠）

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

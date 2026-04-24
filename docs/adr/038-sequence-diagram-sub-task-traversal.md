# 038: sequence diagram における sub task reads/writes の辿り方

- **status**: accepted
- **date**: 2026-04-24

## 背景

ADR-011 で「1ファイル = 1 main task + 複数の sub task」構造を確定した。
ADR-020 で `reads` / `writes` は task node のフィールドとして定義された。
ADR-032 / ADR-036 / ADR-037 で sequence diagram の描画ルールを確定したが、
これらはいずれも「`transition.action` で指定された task の `reads` / `writes` を参照する」という暗黙前提で書かれており、
**main task の `reads` / `writes` が空で sub task のみが DB 操作を持つケース** の扱いが未定義だった。

具体例（UC-001 `order/task/checkout.yaml`）:

```yaml
nodes:
  - id: checkout
    type: task
    main: true
    endpoint: true
    # reads / writes なし（直接DBに触らない）
    ...

  - id: build_order
    type: task              # sub task（ファイル内private）
    reads:  [cart.store.cart_session, auth.store.user_db]
    writes: [order_db]
    ...

  - id: reserve_inventory
    type: task              # sub task
    reads:  [catalog.store.inventory_db]
    writes: [catalog.store.inventory_db]
    ...
```

`transition.action: order.task.checkout` が指すのは main task。
main task の `reads` / `writes` のみを参照すると、このAPIコールがDBに一切触れない図が生成される。
実際には build_order / reserve_inventory が複数DBを操作しており、図と事実が乖離する。

## 決定

### 1. 同一ファイル内の全taskの `reads` / `writes` を集約する

sequence diagramのバックエンドは、`transition.action` が指すtaskと**同一ファイル内の全task**（main task + 全sub task）の `reads` / `writes` を集約してDB participant生成・`API→DB` 矢印生成・DB操作table出力を行う。

- 辿る範囲は**ファイル境界まで**（ADR-011 の「1ファイル = 1 subgraph」と一致）
- sub task がさらに別ファイルの main task を flow 経由で呼ぶケースは辿らない（別APIコールとして独立している扱い）
- クロスファイル参照で別 main task を flow step として呼ぶケースも辿らない（そもそも別endpointの責務）

### 2. Mermaid図上の集約方針は既存のまま

`API→DB` 矢印は既存設計通りstep内で `reads` / `writes` を1本ずつに集約して描画する（sub taskごとに個別矢印を引かない）。

```
API->>DB: reads
DB-->>API:
API->>DB: writes
DB-->>API:
```

### 3. DB操作tableに `sub_task` 列を追加する

内訳を人間およびMCPが辿れるよう、DB操作tableに `sub_task` 列を追加する。

```markdown
| step | task | sub_task | store | 操作 |
|------|------|----------|-------|------|
| 2 | order.task.checkout | build_order | order_db | writes |
| 2 | order.task.checkout | build_order | auth.store.user_db | reads |
| 2 | order.task.checkout | reserve_inventory | catalog.store.inventory_db | reads |
| 2 | order.task.checkout | reserve_inventory | catalog.store.inventory_db | writes |
```

| 列 | 内容 |
|---|---|
| `task` | `transition.action` が指す main task の qualified ID（外部から参照可能な公開ID） |
| `sub_task` | sub task の short ID（ファイル内private、ADR-011）。main task 自身の `reads` / `writes` の行は `-` |

### 4. main task 自身が reads/writes を持つ場合は `-` で統一する

列の出力有無を条件分岐させず、常に `sub_task` 列を出力する。main task 直接のDB操作は `-` で表現する。

```markdown
| step | task | sub_task | store | 操作 |
|------|------|----------|-------|------|
| 1 | auth.task.login | - | user_db | reads |
| 1 | auth.task.login | - | session_store | writes |
```

## 理由

### ファイルスコープでの集約を選んだ理由

候補:

| 案 | 内容 |
|---|---|
| A | main task のみ参照（辿らない） |
| B | main task に aggregate `reads` / `writes` を書かせる |
| C | **同一ファイル内の main + sub tasks を辿る（本決定）** |
| D | flow経由で到達する全taskを再帰的に辿る |

A は DB に触らない図が生成されるため誤解を招く。却下。
B は sub task との情報二重管理が発生し、同期ミスのリスクがある。却下。
D は「1つの sequence diagram = 1つのAPIコール」というスコープを越える。別 main task の呼び出しはそれ自体が別エンドポイントの責務であり、sequence diagram で混ぜると抽象化レベルが崩れる。却下。
C は ADR-011 の「1ファイル = 1 subgraph」境界と自然に一致し、sub task が main task の「inlined helper」である性質と整合する。採用。

### `sub_task` 列を常に出す理由

条件付きで列を省略すると以下の問題が生じる:

- パーサー/レンダラー側で列数が動的に変わる実装負荷
- 同一プロジェクト内のsequence diagram間でtable schema が不統一
- MCP で table を機械的に読むとき、列の有無判定が必要

`-` での統一は上記を回避する。sub task を持たない既存の単純なケースでも `-` が1列増えるだけで情報の欠落はない。

### tree表示を採用しなかった理由

Markdownの字下げやASCII treeはレンダラー依存で壊れやすく、MCP側で機械的にパースしにくい。`sub_task` を独立した列として扱うほうが構造化データとして堅牢。

### sub task名をqualifiedにしない理由

ADR-011 で sub task はファイル内privateと定義されており、外部参照可能なIDを持たない。`task` 列が qualified ID を持つため、`(task, sub_task)` のペアでファイル内一意に特定できる。`sub_task` 列は short ID で十分。

## 影響

- `spec/views/sequence-diagram.md` の以下を更新:
  - `participants` の DB 生成条件に「同一ファイル内 main + sub tasks の reads/writes を参照」と明記
  - `バックエンドによる自動解決` の table で `task の reads / writes` を「同一ファイル内 main + sub tasks の reads / writes」に修正
  - `出力フォーマット` の DB 操作 table 例に `sub_task` 列追加
  - `Mermaid出力イメージ > ログインフロー` の DB 操作 table に `sub_task: -` の行を追加
  - `Mermaid出力イメージ > Webhookフロー` の DB 操作 table に `sub_task: -` の行を追加
  - `depends_on:` に本ADRを追加
- Goパーサー／MCP実装は、sequence diagram 生成時にファイル内全taskのreads/writesを集約する
- UC-001 の `order.task.checkout` を含むシナリオは本ADRのルールでDB参加者が出る

## Evidence
- commit: 6d834d4
- impl commit: tbd
- 参考: 特になし

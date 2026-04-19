# 016: foreachをflow:制御構文に降格

- **status**: accepted
- **date**: 2026-04-19
- **supersedes**: 013

## 背景

ADR-013ではforeachをnode typeとして定義していた。しかしforeachは「繰り返せ」という命令を持つだけであり、実行本体はapplyに指定したtaskが担う。これはfork/joinが実行本体を持たず`flow:`の制御構文として定義されているのと同じ構造であり、foreachもnode typeではなく`flow:`の制御構文として扱うべきという判断に至った。

合わせて、各イテレーションの要素をapply先taskに渡す記法として`$item`シジルを導入する。

## 決定

### foreachはnode typeから廃止し、flow:の制御構文とする

`nodes:`内にforeachを定義しない。`flow:`内で`foreach:`キーワードを使って記述する。

```yaml
flow:
  - foreach: process_item      # apply先taskのID
    mode: sequential           # sequential（デフォルト）or map
    over: fetch_items          # iterateするlistの参照元node ID
    params:
      item: $item              # 現在のイテレーション要素
      config: $params.config   # 他のparamも含めapply先taskのparams wiring
    returns: results           # applyの結果をcollectしたasset名
```

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `foreach` | ✓ | apply先taskのID |
| `over` | ✓ | iterateするlistの参照元node ID |
| `mode` | 任意 | `sequential`（デフォルト）または `map`（並列実行） |
| `params` | 任意 | apply先taskのparams wiring（stepエントリと同じルール）。apply先にparamsがある場合は必須。`$item`で現在のイテレーション要素を参照 |
| `returns` | 任意 | applyの結果をcollectしたasset名 |

### $itemシジルの導入

`$item`はforeachのループ境界から注入される現在のイテレーション要素を指す。`$params`（ファイル境界からの入力）と同じシジル体系の拡張。

| シジル | 意味 |
|--------|------|
| `$params.field` | ファイル境界からの入力（main nodeのparams）の特定フィールドを参照 |
| `$item` | ループ境界からの入力（現在のイテレーション要素） |

### 型の解決は暗黙

`$item`の型は`over`で参照したlistのelement型から暗黙に決まる。apply先taskのparamのmodelとの型一致はGo実装の静的検証で担保する。`$item`に型アノテーションを付ける記法は定義しない。

### DAGのrender

foreachはapply先taskのboxに↻アイコンを装飾する形でレンダリングする。foreachが独立したboxとして描画されることはない。

## 理由

### foreachをnodeから廃止

forkはflow:に存在し`nodes:`に定義がない。foreachも「繰り返せ」という制御命令を持つだけであり、実行本体をapplyに持つ構造はforkのbranchesと同質。`nodes:`にforeachを残すと、「node = 実行単位」という概念の一貫性が崩れる。

### $itemシジル

`$params`が「ファイル境界をまたいで注入される値」を意味するのと対称に、`$item`は「ループ境界をまたいで注入される値」を意味する。node IDとシジルが記法レベルで区別されるため、「fetch_itemsのreturnsからwiring」と「ループ要素の注入」が混同されない。

### 型の暗黙解決

element型をforeach側で明示すると、listのmodel定義とforeachで型情報が二重管理になる（ADR-013と同じ理由）。Go実装でASTを追えば型の整合は静的に検証できるため、YAML上での明示は不要。

## 影響

- ADR-013（foreachノード設計）はsupersededとする
- `spec/nodes.md` のforeachに関する記述を本ADRに基づき更新する
- ADR-015の`$`シジルテーブルに`$item`を追記する

## Evidence
- commit: 86be72f
- impl commit: tbd
- 参考: BPMN 2.0 Multi-Instance Activity（sequential / parallel モード、OMG）

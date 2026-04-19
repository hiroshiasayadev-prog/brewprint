# 013: foreachノード設計

- **status**: accepted
- **date**: 2026-04-19

## 背景

ADR-011にてforeachノードの設計をADR-012に委譲していた。ADR-012で制御フローノード（branch/fork/join）が独立したため、本ADRでforeachを扱う。

foreachは「listを受け取り、各要素に同じ処理を適用する」ループノード。DAGにおけるmap/sequential loopパターンを表現する。

## 決定

### フィールド定義

```yaml
- id: process_items
  type: foreach
  mode: sequential        # sequential（デフォルト）or map
  over:
    name: items
    model: item_list      # list型modelのID参照
  apply: process_item     # 適用するtaskのID（同ファイルのsub taskまたは外部main node）
  params:
    - name: config
      model: app_config   # 各イテレーションに共通で渡す追加パラメータ
  returns:
    name: results
    model: result_list    # applyの結果をcollectしたlist型model
```

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `mode` | 任意 | `sequential`（デフォルト）または `map`（並列実行） |
| `over` | ✓ | ループ対象。list型modelへの参照 |
| `apply` | ✓ | 各要素に適用するtaskのID |
| `params` | 任意 | 各イテレーションに共通で渡す追加パラメータ |
| `returns` | 任意 | applyの結果をcollectしたasset |

### over の型

`over.model` はlist型として定義されたmodelを参照する。element型はmodel定義側で明示するため、foreach側では参照のみ。

```yaml
# model定義側でelement型まで持つ
- id: item_list
  type: model
  kind: list
  element: item           # element型はここで定義
```

### apply の参照先

同ファイル内のsub taskおよび外部ファイルのmain node、どちらも参照可能。

```yaml
apply: process_item       # 同ファイルのsub task
apply: external_module    # 外部main nodeも可
```

### mode の使い分け

| mode | 意味 | 対応する実装パターン |
|------|------|-------------------|
| `sequential` | 要素を順番に1つずつ処理 | 通常のforループ |
| `map` | 全要素を並列処理 | ProcessPoolExecutor / asyncio.gather等 |

デフォルトは `sequential`。

### returns の扱い

foreachはapplyの結果をcollectしてlist型assetとして返す。storeの特定フィールドへの蓄積は、foreachの責務ではなくstore側で解決する（ADR-014に委譲）。

## 理由

### over はmodel参照のみ

element型をforeach側に持たせると、model定義とforeachで型情報が二重管理になる。型の単一管理はmodelに集約する。

### apply の外部参照を許可

参照先をsub taskに限定すると「再利用可能な処理をforeachで回す」パターンが表現できない。ADR-011のスコープ制限は「サブノードが外部から参照されることの禁止」であり、foreachが外部を参照することは別の問題。制限する理由がないため許可する。

### returns はcollectのみ

storeの特定フィールドへの蓄積をforeach側で解決しようとすると、foreachがstore操作の責務を持つことになり、責務が混ざる。foreachはcollectしてassetを返すことに専念し、storeへの書き込みはADR-014（initializes設計）で解決する。

却下した代替案：
- `returns` にstoreのフィールドパスを指定して直接蓄積する → foreachの責務過多として却下

## 影響

- `spec/overview.md` のノード種別表に `foreach` を追記する
- storeの特定フィールドへの蓄積パターンはADR-014（initializes設計）で解決する
- `apply` 先のsub taskとのedge記述はADR-015（ファイル内edge記述構造）に委譲

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: Python ProcessPoolExecutor / asyncio.gather参考

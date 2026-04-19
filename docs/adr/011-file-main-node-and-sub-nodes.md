# 011: 1ファイル=1メインノード + サブノード構造

- **status**: accepted
- **date**: 2026-04-19
- **supersedes**: ADR-010（「1ファイル=1ノード」の部分的上書き）

## 背景

ADR-010は「1ファイル=1ノード」を原則とした。しかし以下のケースで実態と合わなくなることが明確になった。

- `foreach` のようなtaskは、applyする処理（sub task）を必然的に持つ
- mermaid DAGは「このファイルが担う処理のまとまり」をsubgraphとして表現する
- sub taskをすべて別ファイルに切り出すと、単純な処理でもファイルが爆発し、subgraphも描けなくなる

GoやRustの「1ファイル = 1 public型 + 複数のprivate helper」と同じ構造が、YAMLにも必要。

## 決定

### 1. 1ファイル = 1メインノード + 複数のサブノード

1ファイルに複数のノードを定義できる。ただしファイルの代表となる**メインノードは1つだけ**。

```yaml
# task/process_items.yaml

nodes:
  - id: process_items
    type: foreach
    main: true              # メインノード（publicノード）
    over:
      name: items
      model: item_list
    params:
      - name: config
        model: app_config
    apply: process_item     # 同ファイル内のsub task IDを参照
    returns:
      name: results
      model: result_list

  - id: process_item        # サブノード（このファイル内にprivate）
    type: task
    params:
      - name: item
        model: item
      - name: config
        model: app_config
    returns:
      name: result
      model: result
    note: "itemにconfigを適用して変換する"
```

### 2. メインノードの宣言

`main: true` フラグをノードに付与する。1ファイルに1つだけ許容。

### 3. サブノードの可視性スコープ

サブノードは**ファイル内にprivate**。外部モジュールからの参照不可。

| ノード種別 | 外部からの参照 |
|-----------|-------------|
| メインノード（`main: true`） | ✅ 可能（通常のID参照） |
| サブノード | ❌ 不可（ファイル内専用） |

外部から見える表面積をメインノードに絞ることで、Claude Codeが「どのIDを参照していいか」を迷わない。

### 4. mermaidのsubgraph単位

mermaidでDAGを描く際、**1ファイル = 1 subgraph**。
subgraphの中にメインノード・サブノード双方のフローを描く。

```
subgraph process_items[process_items.yaml]
    [items] ──→ ◇process_items(foreach)
    [config] ──→ ◇process_items
                      │ apply
                      ▽
                 [process_item] ──→ ◆result
                      │ collect
                      ▽
                 ◆results
end
```

## 理由

**privateスコープ強制**：sub taskを外部参照可能にすると「どこからでも呼べる内部ヘルパー」が生まれ、モジュール境界が崩れる。CAの層構造強制というADR-010の思想と矛盾する。

**main: true フラグ**：ファイル名からmain nodeを推論するアプローチ（`process_items.yaml` → `process_items`が main）も検討したが、明示フラグのほうが「1ファイルにmainが2つある」誤りを静的検出できる。

**1ファイル=1subgraph**：mermaidのsubgraphはファイル単位の責務の境界を視覚化する。ADR-010が目指した「設計と実装の乖離を構造的に防ぐ」仕組みをsubgraphレベルで担保する。

## 影響

- ADR-010の「1ファイル=1ノード」は本ADRにより「1ファイル=1メインノード」に改訂
- ADR-010のその他の決定（CA強制・ビュー自動導出・model/asset分離・ディレクトリ構造）は継続有効
- `spec/nodes.md`（未作成）にてメインノード・サブノードのフィールド定義を詳細化する
- `branch`（旧仮称 `cond`）/ `fork` / `join` の設計はADR-012に委譲
- `foreach` の設計はADR-013にて確定
- `initializes` の設計はADR-014にて確定
- ファイル内edgeの記述構造（control/data section）はADR-015に委譲

## Evidence
- commit: dfe1020
- impl commit: tbd
- 参考: GoおよびRustの「1ファイル=1public型+privateヘルパー」構造慣習参考

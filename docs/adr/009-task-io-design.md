# 009: taskのI/O設計（params/returns・インライン定義禁止）

- **status**: accepted
- **date**: 2026-04-18

## 背景

taskのI/Oフィールド名として `input`/`output` を使っていたが、複数入力の対応・返り値の cardinality・静的検証性の観点から設計を見直した。

## 決定

### 1. `params` / `returns` への改名

| 旧 | 新 | 理由 |
|----|----|------|
| `input` | `params` | 複数のassetを受け取る場合があるため。関数シグネチャの慣習（Python・Go・TS）に倣う |
| `output` | `returns` | 同上。`return` は予約語リスクがあるため名詞複数形 `returns` を採用 |

参照した公知概念：Python関数定義・Go関数シグネチャ・TypeScript関数型

### 2. `params` は複数asset参照を許容

```yaml
- id: merge_task
  type: task
  params:
    - name: users
      type: user_list    # asset ID
    - name: orders
      type: order_list   # asset ID
  returns: merge_result
```

### 3. `returns` は単一asset参照のみ（複数返し禁止）

複数の値を返す必要がある場合は、名前付きstructのassetでwrapすることを強制する。

```yaml
# NG: returnsに複数を並べる
returns:
  - train_dataset
  - test_dataset

# OK: struct assetでwrapして単一にする
- id: split_result
  type: asset
  kind: struct
  fields:
    - type: dataset
      name: train
    - type: dataset
      name: test

- id: split_task
  type: task
  params:
    - name: data
      type: dataset
  returns: split_result
```

Pythonのtuple unpackのような「記法による複数返し」は、動的型付け＋構文糖衣があって成立する。静的検証性を優先するbrewprintでは採用しない。

GoのMultiple return（`value, error`）は事実上のtry/catchであり「返り値」とは別物として扱う。brewprintはハッピーパス前提（`spec/overview.md`）のためerrorは表現しない。

### 4. インライン定義禁止・すべてのassetは名前付き定義を強制

task内にassetの定義をインラインで書くことを禁止する。

```yaml
# NG: task内にインライン定義
- id: split_task
  type: task
  returns:
    kind: struct
    fields:
      - type: dataset
        name: train

# OK: トップレベルで名前付き定義
- id: split_result
  type: asset
  kind: struct
  fields:
    - type: dataset
      name: train
```

「使い捨てstruct」であっても名前付き定義を強制する。

## 理由

**静的検証性の極振り**：brewprintの実装者はほぼAIを想定している。人間が「書きやすい」インライン記法の利点より、IDによる参照の一意性・静的検証可能性を優先する。インライン定義を許容すると参照不可能な孤立定義が生まれうる。

名前付き強制の副次効果として、実装上も `class SplitResult` のように明示的なクラス定義が生まれ、AIによるコード生成の追跡性が向上する。

参照した公知概念：
- Prisma / GraphQL SDL（スキーマはすべて名前付きで定義する思想）
- TypeScriptの `interface` / `type alias`（インライン型より名前付き型を推奨する慣習）

### 5. `returns` 単一強制のMermaid render

`returns` が単一assetである前提で、そのassetのフィールドの参照状況に応じてDAG図の描画を変える：

**全フィールドをそのまま参照するtaskしかない場合**
→ assetをboxとしてそのまま描画

```
[split_task] --> [split_result] --> [train_task]
```

**一部フィールドのみ参照するtaskがある場合**
→ subgraph内に個別フィールドのboxを描画し、そこから矢印を引く

```
[split_task] --> subgraph split_result
                    [train] --> [train_task]
                    [test]  --> [test_task]
                end
```

これがbrewprintにおけるunpackの視覚的表現であり、YAML上で特別な記法を用意しない。

## 影響

- `spec/overview.md` のendpoint例・taskのYAML例を `params`/`returns` に更新する
- `spec/nodes.md`（未作成）にてtaskのフィールド定義を詳細化する
- `spec/views.md`（未作成）にてDAG図のrender詳細を定義する

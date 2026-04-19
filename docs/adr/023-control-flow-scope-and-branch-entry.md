# 023: 制御フロースコープルールとbranchエントリ設計

- **status**: accepted
- **date**: 2026-04-20

## 背景

ADR-012でbranch/fork/joinのノード種別を確定し、ADR-016でforeachをflow:制御構文に統合したが、以下の2点が未定義のままだった。

1. **制御フロー構文のスコープ境界**: branch/fork/foreachの内部で生成されたassetを外部から参照できるか否かのルールが明文化されていなかった
2. **branchのflow:エントリ**: `step:` / `fork:` / `foreach:` と異なり、branchに対応するflow:エントリが未定義だった

branchのflow:エントリを設計する過程で、「分岐後のassetを合流taskのwiring（`params`）でどう参照するか」という問いが生じた。`|` 記法やbranchノードのプロキシ化など複数の案を検討したが、いずれもADR-015のwiring原則との整合が困難だった。発想を転換し、「制御フロー内部のassetはスコープ外から参照不可」とするルールを制定することで問いを原理的に解消する。

## 決定

### 1. 制御フロースコープルール

**制御フロー構文（`branch` / `fork` / `foreach`）の内部で生成されたassetは、その構文のスコープ外から直接参照不可。**

スコープ外にデータを渡す必要がある場合は、`initializes` で事前宣言したstoreに `writes` で格納し、後続taskが `reads` で参照する。

```yaml
# OK: storeを介して分岐外にデータを渡す
nodes:
  - id: role_result_store
    type: store
    ...

  - id: process_order
    type: task
    initializes: [role_result_store]

  - id: admin_flow
    type: task
    writes: [role_result_store]

  - id: user_flow
    type: task
    writes: [role_result_store]

  - id: finalize
    type: task
    reads: [role_result_store]

flow:
  - branch: route_by_role
    params:
      user: fetch_user
    cases:
      - label: admin
        step: admin_flow
      - label: user
        step: user_flow

  - step: finalize         # role_result_storeをreadsで参照
```

```yaml
# NG: 分岐内のassetをwiring内で直接参照
- step: finalize
  params:
    result: admin_flow     # admin_flowはbranchスコープ内のため参照不可
```

収束が不要な場合（各パスが独立して終端する場合）はstoreも不要。分岐後のtaskはどこからも参照されない「floatingノード」となり、DAGではENDに直行する形でrenderする。

### 2. branchエントリの記法

```yaml
flow:
  - branch: route_by_role
    params:
      user: fetch_user
    cases:
      - label: admin
        step: admin_flow
      - label: user
        step: user_flow
```

#### branchエントリのフィールド

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `branch` | ✓ | branch node ID |
| `params` | 任意 | 分岐判断に使う入力のwiring（stepエントリと同じルール） |
| `cases` | ✓ | 各パスのエントリポイントのリスト |

#### casesエントリのフィールド

| フィールド | 必須 | 内容 |
|-----------|------|------|
| `label` | ✓ | 条件ラベル（人間とLLMへの意味記述。評価はbrewprintのスコープ外） |
| `step` | ✓ | このケースのエントリポイントとなるtask ID（単一） |

`step` は単一のnode IDのみ。エントリポイント以降の後続stepはwiring（`params`参照）から導出されるDAG構造によって決まるため、`cases` 内での列挙は不要。

`fork` の `branches` がstep列を持つのは「どのstepが並列のどのブランチに属するか」を実行エンジンが識別する必要があるため。`branch` は1パスしか実行されず後続はDAGから自明に読めるため、エントリポイントのみで十分（ADR-012と同じ理由）。

### 3. floatingノードのDAGレンダリング

branchのcases内の各stepが後続からwiring参照されない場合（収束しないケース）、そのtaskはfloatingノードとして扱い、DAGではENDに向かうエッジを暗黙に追加してrenderする。

## 理由

### 制御フロースコープルール

fork/foreachとの一貫性から導かれる。

- `foreach` 内の `$item` はループ外から参照不可（ADR-016）
- `fork` のbranches内で生成されたassetを外からwiring参照する記法は存在しない（ADR-015）

branchだけ例外にする根拠がない。「制御フロー構文の内部は独立したスコープ」として統一する。

合流後にデータが必要な場合のstore経由パターンは、ADR-020で確立したクロスエッジ（reads/writes）の自然な適用。新たな記法を導入せずに解決できる。

### branchエントリの`cases:`にラベルを持たせる理由

`fork` の `branches:` は匿名のstep列でよい（全ブランチが実行されるため条件の意味は不要）。`branch` は「どのケースがどの条件で実行されるか」がセマンティクスの核心であり、ラベルがないとLLMが条件の意味を読めない。ADR-008のcommentフィールドと同様、`label` はLLMへの意味コントラクトとして機能する。

## 影響

- `spec/edges.md` に本ADRの内容を反映する
  - 「制御フロースコープ」原則セクションを追加
  - `1-4. branchエントリ` セクションを追加
  - `$シジル体系まとめ` の節番号を更新
- `spec/nodes.md` のbranchセクションにスコープルールの参照を追記する

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: 特になし

# 040: 制御フロー内step wiringの明示化

- **status**: accepted
- **date**: 2026-04-25

## 背景

ADR-015で `flow:` セクションを導入し、通常のtask実行は以下の `step` エントリでparams wiringを明示できるようにした。

```yaml
flow:
  - step: transform
    params:
      raw: fetch_data
```

一方で、fork / branch の制御フロー内部では、実行されるtaskをIDだけで指定する構造になっていた。

```yaml
flow:
  - fork: parallel_processing
    branches:
      - [reserve_inventory]
      - [notify_payment_gateway]
    join: finalize_checkout
    params:
      draft_order: build_order
```

```yaml
flow:
  - branch: route_by_inventory
    params:
      order: check_inventory
    cases:
      - label: in_stock
        step: confirm_order
      - label: out_of_stock
        step: cancel_order
```

この構造では、`reserve_inventory` / `notify_payment_gateway` / `confirm_order` / `cancel_order` といった制御フロー内taskに対して、`params:` を個別に書く場所がない。
そのため以下のような暗黙解釈が必要になっていた。

- `fork.params` を branches 内stepへ同名paramとして暗黙伝播する
- `branch.params` を cases の `step` へ同名paramとして暗黙伝播する

しかしこの暗黙伝播は、`flow:` の基本方針である「wiringはparamsで明示する」と相性が悪い。
また、branchごとに異なる入力を渡したいケースや、同じ値を異なるparam名で渡したいケースを表現できない。

## 決定

### 1. fork.branches は branch object のlistに統一する

`fork.branches` は、各branchを `steps:` を持つobjectとして記述する。
`steps[]` の各要素は通常の `flow` の `step` エントリと同じ形式を使う。

```yaml
flow:
  - fork: parallel_processing
    branches:
      - steps:
          - step: reserve_inventory
            params:
              draft_order: build_order

      - steps:
          - step: notify_payment_gateway
            params:
              draft_order: build_order
    join: finalize_checkout
```

### 2. 旧 `branches: - [step_a, step_b]` 形式は廃止する

以下の短縮形は採用しない。

```yaml
# NG
branches:
  - [reserve_inventory]
  - [notify_payment_gateway]
```

paramsが不要な場合も、`steps:` 配下に `step:` を明示する。

```yaml
branches:
  - steps:
      - step: reserve_inventory
  - steps:
      - step: notify_payment_gateway
```

記法を1つに統一し、パーサー・LLM・人間の読み方を単純にする。

### 3. fork.params による暗黙伝播は採用しない

`fork.params` を branches 内stepへ暗黙伝播する仕様は採用しない。
branch内taskへの入力は、各 `steps[].params` に明示的に書く。

```yaml
# NG: fork.params による暗黙伝播
flow:
  - fork: parallel_processing
    params:
      draft_order: build_order
    branches:
      - steps:
          - step: reserve_inventory
      - steps:
          - step: notify_payment_gateway
    join: finalize_checkout

# OK: 各stepにparamsを書く
flow:
  - fork: parallel_processing
    branches:
      - steps:
          - step: reserve_inventory
            params:
              draft_order: build_order
      - steps:
          - step: notify_payment_gateway
            params:
              draft_order: build_order
    join: finalize_checkout
```

### 4. branch.cases は `params` を持てる

`branch.cases[]` は、caseのエントリtaskに渡す `params:` を持てる。

```yaml
flow:
  - branch: route_by_inventory
    params:
      order: check_inventory
    cases:
      - label: in_stock
        step: confirm_order
        params:
          order: check_inventory
      - label: out_of_stock
        step: cancel_order
        params:
          order: check_inventory
```

`branch.params` は branch node 自身の判定入力としてのみ扱う。
`cases[].params` は case entry task へのwiringとして扱う。
両者は同じ値を参照してもよいが、意味は別である。

### 5. branch caseは引き続き単一entry stepのみ

ADR-023の設計は維持する。
`cases[].step` は分岐先のエントリポイントとなる単一task IDであり、case内のstep列は持たない。
エントリポイント以降の後続taskは、通常のwiringから導出される。

複数stepをcase内に列挙する構文は導入しない。

### 6. join.params の解決はbranch終端stepのreturns.nameで行う

forkの `join:` で指定されたjoin nodeの `params` は、各branch終端stepの `returns.name` と同名一致で解決する。

```yaml
nodes:
  - id: reserve_inventory
    type: task
    returns:
      name: reserved
      model: order

  - id: notify_payment_gateway
    type: task
    returns:
      name: notified
      model: order

  - id: finalize_checkout
    type: join
    params:
      - name: reserved
        model: order
      - name: notified
        model: order
```

上記では以下のように解決する。

- `reserve_inventory.returns.name = reserved` → `finalize_checkout.params.reserved`
- `notify_payment_gateway.returns.name = notified` → `finalize_checkout.params.notified`

一致するbranch終端stepのreturnsが存在しない場合はparser errorとする。

## 理由

### params wiringを明示するため

通常の `flow.step` では `params:` によってwiringを明示する。
制御フロー内部だけ暗黙伝播を許すと、同じ `flow:` セクション内で読み方が分裂する。

`fork.branches[].steps[].params` と `branch.cases[].params` を導入することで、制御フロー内部のtaskも通常stepと同じ規則で読める。

### 短縮形を廃止する理由

`branches: - [step_a, step_b]` の短縮形を残すと、以下の2つの記法が併存する。

```yaml
branches:
  - [step_a]
```

```yaml
branches:
  - steps:
      - step: step_a
```

paramsが不要な場合でも後者で十分に短い。
複数記法を持つメリットより、schemaと実装を単純に保つメリットの方が大きい。

### fork.paramsを暗黙伝播に使わない理由

`fork.params` を共通入力の省略記法として残す案も検討したが、通常stepの `params` と意味がずれ、暗黙展開が必要になる。
また、同じparam名が `fork.params` と `steps[].params` の両方に出た場合の優先順位・衝突規則が必要になる。

制御フロー内taskへの入力は常にそのtaskの `params` に書く、と統一する方が単純である。

### branch.casesにparamsを置く理由

branchはforkと異なり、各caseのエントリポイントは単一taskである。
そのため `cases[].step` と同じ階層に `params:` を置けば、case entry taskへのwiringを過不足なく表現できる。

case内に `steps:` listを導入すると、ADR-023の「branchはエントリポイントのみを持ち、以降はDAG構造から導出する」という設計を崩すため採用しない。

## 影響

- ADR-015のfork branches記法を本ADRで更新する
- ADR-023のbranch cases記法に `cases[].params` を追加する
- `docs/spec/edges.md` の以下を更新する必要がある
  - forkエントリの `branches` schema
  - fork.params の扱い削除
  - branch cases schemaへの `params` 追加
  - join.params 解決ルールの追記
- UC-001 の以下YAMLを更新する必要がある
  - `docs/uc/001-ec-checkout-flow/yaml/order/task/checkout.yaml`
  - `docs/uc/001-ec-checkout-flow/yaml/order/task/process_order.yaml`
- UC-001 のDAG render期待値は、data線を各branch/case taskへ明示的に引く形へ更新する

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: Airflow / Prefect の workflow 内 task call における明示的引数渡し慣習

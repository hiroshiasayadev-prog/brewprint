# 015: ファイル内edge記述構造（flow:セクション）

- **status**: accepted
- **date**: 2026-04-19

## 背景

ADR-011〜014にて以下がADR-015に委譲されていた。

- ファイル内ノード間のデータフローedge記述方法（ADR-011）
- fork/joinの対応関係の記述方法（ADR-012）
- foreachのapply先sub taskとのedge記述（ADR-013）
- storeの参照・更新のedge表現（ADR-014）

また、task nodeのYAML定義（`params`/`returns`）がsignatureのみを担うべきか、wiring情報も持つべきかが未確定だった。

## 決定

### 1. tasks = signatureのみ。wiringは`flow:`セクションに分離

`nodes:`内の各taskは `params`/`returns` によるsignature定義のみを担う。ノード間のwiring（どのnodeの出力がどのnodeの入力になるか）はすべて `flow:` セクションに記述する。

```yaml
nodes:
  - id: fetch_data
    type: task
    params:
      - name: config
        model: app_config
    returns:
      name: raw
      model: raw_data

  - id: transform
    type: task
    params:
      - name: raw
        model: raw_data
    returns:
      name: result
      model: result_data

flow:
  - step: fetch_data
    params:
      config: $params

  - step: transform
    params:
      raw: fetch_data
```

この分離はAirflow（`@task` + `>>`）、Prefect（`@task` + `@flow`）、Temporal（activity + workflow）と同じ構造。

### 2. flow内のデータwiring記法

`params` のvalueに**参照元node ID**を書く。

```yaml
params:
  name: source_node            # source_nodeのreturnsから型・name一致で解決
```

同一modelを持つフィールドが複数あり型・nameだけでは一意に解決できない場合のみ、`.field` で明示する。

```yaml
params:
  raw: fetch_data.rawa         # 曖昧性解消が必要な場合のみ
```

**wiringの単位は常にtaskのreturns全体**。`.field`記法は曖昧性解消のためにのみ使い、「taskのreturnsの一部フィールドを取り出してwiring先taskに渡す」用途では使わない。返り値の一部だけが必要な場合は、extract taskをノードとして明示的に定義する。

```yaml
# NG: returnsの一部フィールドをwiringで直接引く
- step: static_analysis
  params:
    raw: fetch_data.rawa       # fetch_dataのreturns(raw_data)のフィールドをwiring

# OK: extract taskを明示的に挟む
- step: extract_rawa
  params:
    raw: fetch_data
- step: static_analysis
  params:
    raw: extract_rawa
```

### 3. main paramsの参照

ファイルへの外部入力はmain nodeの `params` のみ。flow内からは `$params` シジルで参照する。

```yaml
flow:
  - step: fetch_data
    params:
      config: $params          # main nodeのparamsからname一致で解決

  - step: transform
    params:
      config: $params.config   # 曖昧性解消が必要な場合はフィールド指定
```

`# 015: ファイル内edge記述構造（flow:セクション）

- **status**: accepted
- **date**: 2026-04-19

## 背景

ADR-011〜014にて以下がADR-015に委譲されていた。

- ファイル内ノード間のデータフローedge記述方法（ADR-011）
- fork/joinの対応関係の記述方法（ADR-012）
- foreachのapply先sub taskとのedge記述（ADR-013）
- storeの参照・更新のedge表現（ADR-014）

また、task nodeのYAML定義（`params`/`returns`）がsignatureのみを担うべきか、wiring情報も持つべきかが未確定だった。

## 決定

### 1. tasks = signatureのみ。wiringは`flow:`セクションに分離

`nodes:`内の各taskは `params`/`returns` によるsignature定義のみを担う。ノード間のwiring（どのnodeの出力がどのnodeの入力になるか）はすべて `flow:` セクションに記述する。

```yaml
nodes:
  - id: fetch_data
    type: task
    params:
      - name: config
        model: app_config
    returns:
      name: raw
      model: raw_data

  - id: transform
    type: task
    params:
      - name: raw
        model: raw_data
    returns:
      name: result
      model: result_data

flow:
  - step: fetch_data
    params:
      config: $params

  - step: transform
    params:
      raw: fetch_data
```

この分離はAirflow（`@task` + `>>`）、Prefect（`@task` + `@flow`）、Temporal（activity + workflow）と同じ構造。

### 2. flow内のデータwiring記法

`params` のvalueに**参照元node ID**を書く。

```yaml
params:
  name: source_node            # source_nodeのreturnsから型・name一致で解決
```

同一modelを持つフィールドが複数あり型・nameだけでは一意に解決できない場合のみ、`.field` で明示する。

```yaml
params:
  raw: fetch_data.rawa         # 曖昧性解消が必要な場合のみ
```

**wiringの単位は常にtaskのreturns全体**。`.field`記法は曖昧性解消のためにのみ使い、「taskのreturnsの一部フィールドを取り出してwiring先taskに渡す」用途では使わない。返り値の一部だけが必要な場合は、extract taskをノードとして明示的に定義する。

```yaml
# NG: returnsの一部フィールドをwiringで直接引く
- step: static_analysis
  params:
    raw: fetch_data.rawa       # fetch_dataのreturns(raw_data)のフィールドをwiring

# OK: extract taskを明示的に挟む
- step: extract_rawa
  params:
    raw: fetch_data
- step: static_analysis
  params:
    raw: extract_rawa
```

### 3. main paramsの参照

ファイルへの外部入力はmain nodeの `params` のみ。flow内からは `$params` シジルで参照する。

```yaml
flow:
  - step: fetch_data
    params:
      config: $params          # main nodeのparamsからname一致で解決

  - step: transform
    params:
      config: $params.config   # 曖昧性解消が必要な場合はフィールド指定
```

シジルは「境界をまたいで注入される値」を意味する。

| シジル | 意味 |
|--------|------|
| `$params` | ファイル境界からの入力（main nodeのparams） |
| `$item` | ループ境界からの入力（foreachの現在のイテレーション要素）。ADR-016参照 |

### 4. fork/joinの記法

fork側に `branches:` と `join:` をまとめて記述する。

```yaml
flow:
  - fork: fan_out
    branches:
      - [static_analysis]
      - [dynamic_analysis]
      - [dep_check]
    join: aggregate
    params:
      raw: fetch_data          # 各branchへの共通入力
```

`join:` はこのforkに対応するjoin nodeのIDを指定する。branchesは各ブランチのstep列をリストで記述する。

### 5. storeのreads/writes

flow内の各stepに `reads:` / `writes:` フィールドで明示する。`params`/`returns`（asset）とは直交するフィールド。

```yaml
flow:
  - step: enrich_report
    params:
      result: aggregate
    reads: [cache]
    writes: [report]
```

| フィールド | 対象 | 意味 |
|-----------|------|------|
| `params` | asset | 上流taskのreturnsからデータを受け取る |
| `returns` | asset | 下流taskへデータを渡す（signatureで定義済み） |
| `reads` | store | storeの状態を参照する |
| `writes` | store | storeの状態を更新する |

storeに書き込みつつassetをreturnするステップは正常なユースケース。

### 6. pointer/replicaはスコープ外

flow内のwiring記法は「何が何に繋がるか」を表現する。データのコピー渡しか参照渡しかは実装言語依存であり、brewprintの言語仕様に含めない。意味的に重要な場合はextract taskの `note` に記述する。

## 理由

### tasks = signatureのみ

taskのYAML定義にwiring情報を混在させると、「このtaskはどのtaskと繋がっているか」という情報が分散する。flow:セクションに集約することで、ファイル内のデータフローが1箇所で把握できる。

Airflow/Prefect/Temporalの「task定義とorchestrationの分離」と同じ構造であり、業界標準の設計思想に沿う。

### wiringはreturns全体単位

`fetch_data.rawa` のようにtaskのreturnsのフィールドを直接wiringで参照すると、「関数の返り値の中身を呼び出し元が直接展開する」セマンティクスになる。これはtaskのsignature設計を無意味にし、静的検証の複雑化を招く。

ADR-009でも「returns単一強制」と「フィールドレベルの分岐はsubgraphで視覚的に表現」という同じ判断をしている。extract taskを明示することで責務が明確になり、LLMによる追跡性も向上する。

### fork側にまとめる

branchesとjoinの対応関係をfork側に書くことで、「このforkは何を並列実行し、どこで合流するか」が1箇所で読める。join側に`fork:`参照を置く逆方向の設計だと、forkとjoinを別々に読まないと全体像が分からない。

### $paramsシジル

`config: process_report`（main node ID参照）にすると「main nodeのreturnsから引く」という誤読が生じる。`$params`にすることで、node IDとファイル境界入力が記法レベルで区別できる。

## 影響

- `spec/nodes.md`（未作成）にてtask/foreach/branch/fork/joinの`flow:`内での記法を詳細化する
- `spec/views.md`（未作成）にてflow:セクションからDAG図を導出するrender規則を定義する
- ADR-013のforeachの `apply:` はsub taskのID参照であり、`flow:`セクションのwiring記法とは別の機構。spec/nodes.mdにて整合を確認する

## Evidence
- commit: ed75964
- impl commit: tbd
- 参考: Airflow BranchPythonOperator / task decorator、Prefect @flow / @task、Temporal workflow / activity参考

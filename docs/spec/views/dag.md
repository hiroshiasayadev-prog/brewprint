---
scope: docs/spec/views/dag.md
status: confirmed
last_updated: 2026-04-26
summary: >
  DAG（Directed Acyclic Graph）のrenderルール定義。
  Processingレイヤーのノード・エッジをMermaid flowchartとして出力する際の
  ノード形状・エッジ種別・スコープ・特殊ケースを定義する。
depends_on:
  - docs/adr/009-task-io-design.md
  - docs/adr/011-file-main-node-and-sub-nodes.md
  - docs/adr/012-control-flow-nodes.md
  - docs/adr/015-file-internal-edge-structure.md
  - docs/adr/016-foreach-as-flow-construct.md
  - docs/adr/017-diagram-layers-and-scope.md
  - docs/adr/020-cross-edge-management.md
  - docs/adr/022-dag-node-shapes-and-edge-types.md
  - docs/adr/023-control-flow-scope-and-branch-entry.md
  - docs/adr/024-dag-boundary-nodes.md
  - docs/adr/040-control-flow-step-wiring.md
  - docs/adr/044-store-access-edge-labels.md
---

# DAG renderルール

## スコープ

1ファイル = 1DAG。メインノードのファイルを単位としてDAGを描画する（ADR-011）。
複数ファイルをまたぐモジュール全体DAGは定義しない。

他ファイルのメインノードを参照する場合（foreach.applyの外部参照等）は、
参照先ノードを外部ノードとして描画する（→ [外部参照ノード](#外部参照ノード)）。

---

## 出力フォーマット

```markdown
# {メインノードID}

**API**: [{method} {path}](../views/api-table.md)

{メインノードのnote}

​```mermaid
flowchart TD
  ...
​```

## Tasks

### {task_id}
...
```

- H1 = メインノードの `id`
- **API行** = メインノードが `endpoint: true` の場合のみ出力。`note` の前に置く
- 説明文 = メインノードの `note`。`note` がない場合は省略
- Mermaid記法: `flowchart TD`（上から下）
- **Tasks詳細セクション** = Mermaid図の後に続く。各taskのsignature・reads/writes・noteを一覧する

### Tasks詳細セクションのフォーマット

```markdown
## Tasks

### login
**params**: credentials (credential)
**returns**: auth_token (token)
**reads**: session_store
**writes**: session_store
**note**: 認証情報を検証しトークンを発行する

### other_task
**外部参照**: [auth.task.validate](../../auth/task/validate.md)
```

- 同ファイル内のtask: signature（params/returns）・reads/writes・noteを列挙
- 外部参照taskのみ: `**外部参照**:` でリンクを示し詳細は省略
- `reads` / `writes` がない場合は該当行を省略
- `note` がない場合は該当行を省略

---

## ノードのrender

### start / end

```
_start([Start])
_end([End])
```

形状: スタジアム（Mermaid `([label])`）。ISO 5807:1985 Terminal記号に対応（ADR-022）。
Mermaid ID は `_start` / `_end`（キーワード衝突回避のためアンダースコアプレフィックス付与）（ADR-024）。
DAGの先頭に `_start([Start])` を置き、最後のtaskから `_end([End])` へ制御線（`==>`）を引く。floatingノード（ADR-023）も `_end` へ接続する。

**floatingノード**: flow内で後続のwiring参照がないtask。branchの各caseタスクのように、分岐先で処理が完結しそれ以降のstepがない場合に発生する。floatingノードには暗黙的に `==> _end([End])` を追加して描画する。

```
admin_flow ==> _end([End])
user_flow ==> _end
```

（`_end` は1つのノードなので、2本以上の制御線が収束してよい）

### task

```
task_id[task_id]
```

形状: 矩形（Mermaid `[label]`）

### asset

```
asset_name([asset_name])
```

形状: スタジアム（Mermaid `([label])`）。taskの `returns` から暗黙に生成される中間ノード（ADR-009）。

### store

```
store_id[(store_id)]
```

形状: シリンダー（Mermaid `[(label)]`）。`reads` / `writes` エッジで task と接続される（ADR-020）。

### branch

```
branch_id{branch_id}
```

形状: ひし形（Mermaid `{label}`）。排他分岐（ADR-012）。

### fork / join

```
fork_id{{fork_id}}
join_id{{join_id}}
```

形状: 六角形（Mermaid `{{label}}`）。並列分岐・合流（ADR-012）。

> UML 2.x標準のfork/joinはバー記号（━━）だが、Mermaid flowchartはバー記号を再現できないため六角形を代替として採用する（ADR-022）。

### 外部参照ノード

同ファイル外のメインノードを参照する場合、classDefで色を変えて区別する。

```
classDef external fill:#e0e0e0,stroke:#999,color:#555
class other_task external
```

外部ノードの形状はノード種別に従う（task → 矩形、等）。

### subgraph params / returns

メインノードの入出力を図の境界として `subgraph` で囲む（ADR-024）。

```
subgraph params
  config([config])
end
subgraph returns
  result([result])
end
```

- `subgraph params`: メインノードに `params` がある場合、各paramをassetノードとして列挙する
- `subgraph returns`: メインノードに `returns` がある場合のみ描画。`returns` がない場合は省略
- 境界assetには `boundaryNode` classDef を適用する

### foreachの↻装飾

`foreach` はnode typeではなく `flow:` の制御構文のため、独立したノードとして描画しない（ADR-016）。
apply先taskのノードラベルに `↻` を付与して表現する。

```
process_item["↻ process_item"]
```

apply先が外部参照ノードの場合は外部ノードのclassDefと組み合わせる。

### ノードの色付け

種別ごとにclassDefで色分けする。WCAG 2.1 Level AA（コントラスト比4.5:1以上）に準拠（ADR-022）。

```
classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
classDef branchNode   fill:#9B6BBD,stroke:#6B3D8F,color:#fff
classDef forkNode     fill:#8A8A8A,stroke:#5A5A5A,color:#fff
classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
classDef external     fill:#E0E0E0,stroke:#999,color:#555
```

`boundaryNode` は `subgraph params` / `subgraph returns` 内の境界assetに適用する（ADR-024）。

各ノードに対応するclassを付与する。

```
class login taskNode
class auth_token assetNode
class session_store storeNode
class route_by_role branchNode
class fan_out,aggregate forkNode
class _start,_end terminalNode
class config,result boundaryNode
```

---

## エッジのrender

エッジはUML Activity DiagramのControlFlow / ObjectFlowの区別に対応する（OMG UML 2.x / ADR-022）。

| 種別 | UML対応 | Mermaid記法 | 用途 |
|------|---------|------------|------|
| データ線 | ObjectFlow | `-->` | データの受け渡し |
| ラベル付きデータ線 | ObjectFlow | `--"label"-->` | foreach等、意味を付与するデータの受け渡し |
| store access線 | ObjectFlow | `-- "read" -->` / `-- "write" -->` / `<-- "read/write" -->` | storeへの読み書きアクセス |
| 制御線 | ControlFlow | `==>` | 実行順序の制御 |
| ラベル付き制御線 | ControlFlow | `== "label" ==>` | branch・forkの条件付き制御フロー |

### データ線（`-->`）

task → asset、asset → task のwiring。`flow:` の `params` wiringから導出する。

```
fetch_data --> raw([raw])
raw --> transform
```

store の `reads` / `writes` は store access線として描く。通常の asset dataflow と区別するため、アクセス種別を edge label で明示する（ADR-044）。

| YAML上の指定 | Mermaid表現 | 意味 |
|-------------|-------------|------|
| `reads: [store]` | `store -- "read" --> task` | task が store を読む |
| `writes: [store]` | `task -- "write" --> store` | task が store に書く |
| `reads` と `writes` の両方 | `task <-- "read/write" --> store` | task が store を読み書きする |

store edge の向きは従来どおり維持するが、意味の主表現はラベルに寄せる。

```
session_store[(session_store)] -- "read" --> login       %% reads のみ
login -- "write" --> audit_log[(audit_log)]              %% writes のみ
login <-- "read/write" --> session_store[(session_store)] %% reads + writes 両方
```

### 制御線（`==>`）

以下の全ケースで制御線を引く。

- `_start` → 最初のtask（DAG起点）
- 最後のtask → `_end`（DAG終点）
- floatingノード → `_end`（暗黙終端）
- **flow上で連続するtask間**（data線があっても必ず併記する）
- task → branch / fork（分岐への入り口）
- branch / join → 後続task（分岐・合流の出口）

```
process_report ==> transform       %% 連続するtask間
fetch_user ==> route_by_role       %% task → branch
fetch_data ==> fan_out{{fan_out}}  %% task → fork
```

> データ線（`-->`）だけでは実行順序がLLMに伝わりにくい。UML ControlFlowとして明示することで、DAGの制御構造を機械的に解析可能にする（ADR-022）。

**branch（排他分岐）**

`cases[].label` をエッジラベルに使う（ADR-022, ADR-023）。

```
route_by_role{route_by_role} == "admin" ==> admin_flow
route_by_role{route_by_role} == "user" ==> user_flow
```

**branchのparams wiring（データ線）**: branchノードが受け取るassetは2種のデータ線を引く。
1本はbranchノード自身（ルーティング判断用）、もう1本は各branch task（実行時に使うデータ）。

```
user --> route_by_role      %% ルーティング判断用
user --> admin_flow         %% admin ブランチへのデータ
user --> user_flow          %% user ブランチへのデータ
```

**fork / join（並列実行）**

forkからの各ブランチに `"parallel"` ラベルを付ける（BPMN 2.0 Parallel Gateway準拠 / ADR-022）。
joinノードへは、各branchタスクからの **制御線（合流）とデータ線（結果）の両方** を引く。

```
fan_out{{fan_out}} == "parallel" ==> static_analysis
fan_out{{fan_out}} == "parallel" ==> dynamic_analysis
fan_out{{fan_out}} == "parallel" ==> dep_check
static_analysis --> static_result([static_result])   %% データ線（branchタスクの出力）
static_result --> aggregate{{aggregate}}              %% データ線（joinへの入力）
static_analysis ==> aggregate{{aggregate}}            %% 制御線（合流）
dynamic_analysis --> dynamic_result([dynamic_result])
dynamic_result --> aggregate
dynamic_analysis ==> aggregate
dep_check --> dep_result([dep_result])
dep_result --> aggregate
dep_check ==> aggregate
```

**foreach**

「foreach」はデータがどう流れるか（1件ずつ）の意味なので、データ線（ObjectFlow）にラベルとして乗せる。
タスク間の実行順序は制御線（ControlFlow）で示す（BPMN 2.0 Multi-Instance Activity準拠 / ADR-022）。

```
fetch_items ==> process_item["↻ process_item"]    %% 制御線（実行順序）
items --"foreach"--> process_item                 %% データ線（foreachラベル付きObjectFlow）
```

---

## render例

### 基本的なDAG（2 task + asset）

YAML:
```yaml
nodes:
  - id: process_report
    type: task
    main: true
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
  - step: process_report
    params:
      config: $params.config
  - step: transform
    params:
      raw: process_report
```

Mermaid出力:
```mermaid
flowchart TD
  subgraph params
    config([config])
  end
  subgraph returns
    result([result])
  end

  _start([Start]) ==> process_report
  config --> process_report[process_report]
  process_report --> raw([raw])
  process_report ==> transform
  raw --> transform[transform]
  transform --> result
  transform ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class process_report,transform taskNode
  class raw assetNode
  class _start,_end terminalNode
  class config,result boundaryNode
```

### fork / join を含むDAG

YAML:
```yaml
nodes:
  - id: analyze_code
    type: task
    main: true
    params:
      - name: source_code
        model: source_code
    returns:
      name: full_report
      model: full_report

  - id: fetch_data
    type: task
    params:
      - name: source_code
        model: source_code
    returns:
      name: raw
      model: raw_data

  - id: fan_out
    type: fork

  - id: static_analysis
    type: task
    params:
      - name: raw
        model: raw_data
    returns:
      name: static_result
      model: static_result

  - id: dynamic_analysis
    type: task
    params:
      - name: raw
        model: raw_data
    returns:
      name: dynamic_result
      model: dynamic_result

  - id: dep_check
    type: task
    params:
      - name: raw
        model: raw_data
    returns:
      name: dep_result
      model: dep_result

  - id: aggregate
    type: join
    params:
      - name: static_result
        model: static_result
      - name: dynamic_result
        model: dynamic_result
      - name: dep_result
        model: dep_result
    returns:
      name: full_report
      model: full_report

flow:
  - step: fetch_data
    params:
      source_code: $params.source_code
  - fork: fan_out
    branches:
      - steps:
          - step: static_analysis
            params:
              raw: fetch_data
      - steps:
          - step: dynamic_analysis
            params:
              raw: fetch_data
      - steps:
          - step: dep_check
            params:
              raw: fetch_data
    join: aggregate
```

Mermaid出力:
```mermaid
flowchart TD
  subgraph params
    source_code([source_code])
  end
  subgraph returns
    full_report([full_report])
  end

  _start([Start]) ==> fetch_data
  source_code --> fetch_data[fetch_data]
  fetch_data --> raw([raw])
  fetch_data ==> fan_out{{fan_out}}
  raw --> static_analysis[static_analysis]
  raw --> dynamic_analysis[dynamic_analysis]
  raw --> dep_check[dep_check]
  fan_out == "parallel" ==> static_analysis
  fan_out == "parallel" ==> dynamic_analysis
  fan_out == "parallel" ==> dep_check
  static_analysis --> static_result([static_result])
  dynamic_analysis --> dynamic_result([dynamic_result])
  dep_check --> dep_result([dep_result])
  static_result --> aggregate{{aggregate}}
  dynamic_result --> aggregate
  dep_result --> aggregate
  static_analysis ==> aggregate
  dynamic_analysis ==> aggregate
  dep_check ==> aggregate
  aggregate --> full_report
  aggregate ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef forkNode     fill:#8A8A8A,stroke:#5A5A5A,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class fetch_data,static_analysis,dynamic_analysis,dep_check taskNode
  class raw,static_result,dynamic_result,dep_result assetNode
  class fan_out,aggregate forkNode
  class _start,_end terminalNode
  class source_code,full_report boundaryNode
```

### storeを含むDAG

YAML:
```yaml
nodes:
  - id: authenticate
    type: task
    main: true
    params:
      - name: credentials
        model: credential
    returns:
      name: auth_token
      model: token

  - id: login
    type: task
    params:
      - name: credentials
        model: credential
    returns:
      name: auth_token
      model: token
    reads: [session_store]
    writes: [session_store]

  - id: session_store
    type: store
    model: session

flow:
  - step: login
    params:
      credentials: $params.credentials
```

Mermaid出力:
```mermaid
flowchart TD
  subgraph params
    credentials([credentials])
  end
  subgraph returns
    auth_token([auth_token])
  end

  _start([Start]) ==> login
  credentials --> login[login]
  login <-- "read/write" --> session_store[(session_store)]
  login --> auth_token
  login ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef storeNode    fill:#E8A838,stroke:#B07820,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class login taskNode
  class session_store storeNode
  class _start,_end terminalNode
  class credentials,auth_token boundaryNode
```

### foreachを含むDAG

YAML:
```yaml
nodes:
  - id: process_reports
    type: task
    main: true
    params:
      - name: config
        model: app_config
    returns:
      name: results
      model: result_list

  - id: fetch_items
    type: task
    params:
      - name: config
        model: app_config
    returns:
      name: items
      model: item_list

  - id: process_item
    type: task
    params:
      - name: item
        model: item
    returns:
      name: result
      model: result

flow:
  - step: fetch_items
    params:
      config: $params.config
  - foreach: process_item
    over: fetch_items
    params:
      item: $item
    returns: results
```

Mermaid出力:
```mermaid
flowchart TD
  subgraph params
    config([config])
  end
  subgraph returns
    results([results])
  end

  _start([Start]) ==> fetch_items
  config --> fetch_items[fetch_items]
  fetch_items --> items([items])
  fetch_items ==> process_item["↻ process_item"]
  items --"foreach"--> process_item
  process_item --> results
  process_item ==> _end([End])

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class fetch_items,process_item taskNode
  class items assetNode
  class _start,_end terminalNode
  class config,results boundaryNode
```

### branchを含むDAG（floatingノードあり）

returnsなし・floatingノードを含むパターン。branch caseタスクが後続なしで完結するため、
両caseタスクがfloatingノードとなり `_end` へ暗黙接続される。

YAML:
```yaml
nodes:
  - id: handle_request
    type: task
    main: true
    params:
      - name: user_id
        model: user_id

  - id: fetch_user
    type: task
    params:
      - name: user_id
        model: user_id
    returns:
      name: user
      model: user

  - id: route_by_role
    type: branch
    params:
      - name: user
        model: user

  - id: admin_flow
    type: task
    params:
      - name: user
        model: user

  - id: user_flow
    type: task
    params:
      - name: user
        model: user

flow:
  - step: fetch_user
    params:
      user_id: $params.user_id
  - branch: route_by_role
    params:
      user: fetch_user
    cases:
      - label: admin
        step: admin_flow
        params:
          user: fetch_user
      - label: user
        step: user_flow
        params:
          user: fetch_user
```

Mermaid出力:
```mermaid
flowchart TD
  subgraph params
    user_id([user_id])
  end

  _start([Start]) ==> fetch_user
  user_id --> fetch_user[fetch_user]
  fetch_user --> user([user])
  user --> route_by_role{route_by_role}
  fetch_user ==> route_by_role
  user --> admin_flow[admin_flow]
  user --> user_flow[user_flow]
  route_by_role == "admin" ==> admin_flow
  route_by_role == "user" ==> user_flow
  admin_flow ==> _end([End])
  user_flow ==> _end

  classDef taskNode     fill:#4A90D9,stroke:#2C5F8A,color:#fff
  classDef assetNode    fill:#5BA55B,stroke:#3A6B3A,color:#fff
  classDef branchNode   fill:#9B6BBD,stroke:#6B3D8F,color:#fff
  classDef terminalNode fill:#2C2C2C,stroke:#000,color:#fff
  classDef boundaryNode fill:#2D7D9A,stroke:#1A5068,color:#fff
  class fetch_user,admin_flow,user_flow taskNode
  class user assetNode
  class route_by_role branchNode
  class _start,_end terminalNode
  class user_id boundaryNode
```


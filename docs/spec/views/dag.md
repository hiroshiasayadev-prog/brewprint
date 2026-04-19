---
scope: docs/spec/views/dag.md
status: wip
last_updated: 2026-04-20
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

{メインノードのnote}

​```mermaid
flowchart TD
  ...
​```
```

- H1 = メインノードの `id`
- 説明文 = メインノードの `note`。`note` がない場合は説明文を省略する
- Mermaid記法: `flowchart TD`（上から下）

---

## ノードのrender

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

### foreachの↻装飾

`foreach` はnode typeではなく `flow:` の制御構文のため、独立したノードとして描画しない（ADR-016）。
apply先taskのノードラベルに `↻` を付与して表現する。

```
process_item["↻ process_item"]
```

apply先が外部参照ノードの場合は外部ノードのclassDefと組み合わせる。

---

## エッジのrender

エッジはUML Activity DiagramのControlFlow / ObjectFlowの区別に対応する（OMG UML 2.x / ADR-022）。

| 種別 | UML対応 | Mermaid記法 | 用途 |
|------|---------|------------|------|
| データ線 | ObjectFlow | `-->` | データの受け渡し |
| 制御線 | ControlFlow | `==>` | 実行順序の制御 |

### データ線（`-->`）

task → asset、asset → task のwiring。`flow:` の `params` wiringから導出する。

```
fetch_data --> raw([raw])
raw --> transform
```

store の reads / writes も同じデータ線で描く。方向で reads / writes を区別する。

```
session_store[(session_store)] --> login
login --> session_store[(session_store)]
```

### 制御線（`==>`）

branch・fork・join を含む制御フローに使う。

**branch（排他分岐）**

`cases[].label` をエッジラベルに使う（ADR-022, ADR-023）。

```
route_by_role{route_by_role} == "admin" ==> admin_flow
route_by_role{route_by_role} == "user" ==> user_flow
```

**fork / join（並列実行）**

forkからの各ブランチに `"parallel"` ラベルを付ける（BPMN 2.0 Parallel Gateway準拠 / ADR-022）。

```
fan_out{{fan_out}} == "parallel" ==> static_analysis
fan_out{{fan_out}} == "parallel" ==> dynamic_analysis
fan_out{{fan_out}} == "parallel" ==> dep_check
static_analysis ==> aggregate{{aggregate}}
dynamic_analysis ==> aggregate{{aggregate}}
dep_check ==> aggregate{{aggregate}}
```

**foreach**

apply先taskへのエッジに `"foreach"` ラベルを付ける（BPMN 2.0 Multi-Instance Activity準拠 / ADR-022）。

```
items([items]) == "foreach" ==> process_item["↻ process_item"]
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
  process_report[process_report] --> raw([raw])
  raw --> transform[transform]
  transform --> result([result])
```

### fork / join を含むDAG

```mermaid
flowchart TD
  fetch_data[fetch_data] --> raw([raw])
  raw --> fan_out{{fan_out}}
  fan_out == "parallel" ==> static_analysis[static_analysis]
  fan_out == "parallel" ==> dynamic_analysis[dynamic_analysis]
  fan_out == "parallel" ==> dep_check[dep_check]
  static_analysis ==> aggregate{{aggregate}}
  dynamic_analysis ==> aggregate{{aggregate}}
  dep_check ==> aggregate{{aggregate}}
  aggregate --> full_report([full_report])
```

### storeを含むDAG

```mermaid
flowchart TD
  credentials([credentials]) --> login[login]
  session_store[(session_store)] --> login
  login --> session_store
  login --> auth_token([auth_token])
```

### foreachを含むDAG

```mermaid
flowchart TD
  fetch_items[fetch_items] --> items([items])
  items == "foreach" ==> process_item["↻ process_item"]
  process_item --> results([results])
```

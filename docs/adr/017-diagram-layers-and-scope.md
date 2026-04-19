# 017: 図レイヤー定義とbrewprintのスコープ確定

- **status**: accepted
- **date**: 2026-04-19

## 背景

eventノードの設計を進める中で、「eventをDAGに書くべきか」という問いが発生した。
議論の過程で、brewprintが扱う図を抽象レイヤーで整理しないと、ノードの所属・レイヤー間の依存方向・スコープ外の明示ができないことが明らかになった。

合わせて、class diagramの位置づけとrepo/infraのモデリング要否についても判断が必要だった。

## 決定

### 1. 図を3レイヤーで分類する

| レイヤー | 図 | 答える問い |
|---------|---|-----------|
| **Application** | Sequence Diagram, State Diagram, API Table | 誰が・何がいつ起きるか。actorの行動・eventの発生・状態遷移・APIルート |
| **Processing** | DAG | 処理の中身はどうなっているか。taskのデータフロー・制御フロー |
| **Data** | ER Diagram | データはどんな構造か。storeの定義 |

### 2. レイヤー間の依存方向

```
Application → Processing → Data
```

- ApplicationレイヤーはProcessingレイヤーのノード（task等）を参照できる
- ProcessingレイヤーはApplicationレイヤーのノード（event / state等）を参照しない
- 逆方向の参照は禁止

具体的には：
- Sequence DiagramはDAGのtask（endpoint=true）を参照する（矢印ラベルにtask IDを記載）
- DAGにevent / stateノードは登場しない

### 3. class diagramを廃止し、API TableをMCPツールとして提供

`class diagram`（`endpoint: true`なtaskをモジュール別にまとめた図）は廃止する。

理由：
- brewprintはOOPのclass構造（継承・合成・interface）のモデリングを目的としない
- APIルート一覧は図として描画するよりMDテーブルとして出力する方が実用的
- Swagger/OpenAPIと同様の構造で `list_endpoints` MCPツールとして提供する

`list_endpoints` の出力イメージ：
```markdown
| Method | Path | Task ID | Params | Returns |
|--------|------|---------|--------|---------|
| POST | /auth/login | auth.task.login | login_request | token |
| GET  | /users/{id} | user.task.get_user | — | user |
```

### 4. repo/infraのモデリングはスコープ外

Repository / Infrastructure層の詳細（interfaceとimplの対応・DI設定等）はbrewprintの対象外とする。

### 5. brewprintはclean-arch前提であり、宣言は不要

| brewprintのノード | clean-archでの対応 |
|---|---|
| `model` | Domain Entity（型定義） |
| `task` | Use Case / Application Service |
| `store` | Repository / Infrastructure |
| `task(endpoint=true)` | Controller / Interface Adapter |
| `asset` | Use Case間を流れるDTO |

brewprintのノード設計にclean-archの層構造が織り込まれているため、`pattern: clean-arch`のような宣言フィールドは不要。

## 理由

### レイヤー分類の根拠

「eventはDAGに出るべきか」という問いに答えるには、「DAGとは何を表現する図か」を先に定義する必要があった。レイヤーという概念を導入することで、ノードの所属とレイヤー間の依存方向を原則として記述できる。

「問いの整理」ではなくレイヤー分類を採用した理由は、依存方向の制約（DAGはeventを参照しない等）をレイヤー原則として明文化できるため。

### class diagram廃止の根拠

brewprintは「ドメインロジックの設計言語」であり、OOPの構造モデリングツールではない。class diagramが担っていた内容（APIルート一覧）はMCPツールとして提供する方がLLMにとっても人間にとっても扱いやすい。

### repo/infraのスコープ外

brewprintはDAG・State Diagram・Sequence Diagramというドメイン色の強い図に重きを置いている。repo/infraは「どう実装するか」の関心であり、「何を・どんな順序で・どんなデータで処理するか」というbrewprintの主眼から外れる。

Claude Codeによる自立コーディングを支援する上でも、パターンが明示されていれば詳細なrepo/infra specがなくても十分な精度で動作する。むしろ中途半端な情報がClaude Codeを惑わせるリスクがある。

## 影響

- `spec/overview.md` の「書ける図の一覧」をレイヤー別テーブルに更新する
- `spec/overview.md` の class diagram に関する記述（`endpoint: true`のclass diagram viewへの出力）を削除し、MCPツールによる提供に変更する
- `task`の `endpoint: true` フラグは存続するが、出力先は class diagram ではなく `list_endpoints` MCPツール
- eventノードの設計はADR-018へ
- stateノード（FSM）の設計はADR-019へ

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: clean-arch層構造参考

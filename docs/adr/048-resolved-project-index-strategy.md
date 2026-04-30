# 048: ResolvedProject index strategy

- **status**: accepted
- **date**: 2026-04-27
- **supersedes**: なし

※ ADR-049 により `GetDeps` / `get_deps` は `GetReferences` / `get_references` に名称変更。

## 背景

ADR-047で、Go実装では `Raw YAML structs` を application-facing layer から隠し、`ResolvedProject` を `QueryService` / `Renderer` / `MCP wrapper` の入力境界にすることを決めた。

次に決めるべきことは、`ResolvedProject` がどのようなindexを持つかである。

brewprintでは、同じYAMLから以下のような問い合わせが繰り返し発生する。

- あるQualifiedIDのnodeを取得する
- あるfileに定義されたnodeを取得する
- あるstoreを読むtask / 書くtaskを探す
- あるFSM state + event + guardからtransitionを特定する
- あるtaskをactionとして呼ぶtransitionを探す
- あるscenario IDからsequence scenarioを取得する
- rendererがDAG / State / Sequence / ER / API / Wireframeの入力を集める
- MCP toolが `get_signature` / `get_references` / `inspect` の返却材料を集める

これらを都度Raw YAMLまたはResolvedProject内の全slice走査で実装すると、以下の問題が起きる。

- renderer / QueryService / MCP toolごとに探索ロジックが分散する
- `reads` / `writes` / `transition.action` / scenarioなどの逆引きが毎回再実装される
- name resolution済みの参照と未解決文字列が混ざりやすくなる
- later implementationで性能問題が出たときに、どこを最適化するか分かりにくい

そのため、`ResolvedProject` build時に必要なindexを構築し、外部出力に近い層はそのindexを読む方針を定める。

## 決定

### 1. indexはResolvedProject build時に構築する

`ResolvedProject` のindexは、Raw YAML decode後の semantic build pipeline で構築する。

```text
Raw YAML structs
  ↓ collect symbols
  ↓ resolve references
  ↓ validate semantic constraints
  ↓ build derived nodes / relations
  ↓ build indexes
ResolvedProject
```

indexは `ResolvedProject` の一部として扱う。
`QueryService` / `Renderer` / `MCP wrapper` は、都度Raw YAML structsを走査して逆引きを作らない。

### 2. indexは解決済みIDをkeyにする

indexのkeyには、原則として解決済みのIDを使う。

- node参照: `QualifiedID`
- file参照: file ID。file IDはblueprint YAML root（通常 `yaml/`）からの相対パスをslash正規化した文字列とする（例: `auth/task/login.yaml`）。変更が必要になった場合は別ADRで判断する
- store参照: store nodeの `QualifiedID`
- task参照: task nodeの `QualifiedID`
- scenario参照: scenario ID
- FSM transition参照: state file ID + from state + event + guard

Raw YAMLに書かれた未解決文字列を、application-facing layer向けindexのkeyにしてはならない。
ただし、diagnostic用に元のraw文字列やsource locationを保持することは許可する。

### 3. indexはread-onlyな派生情報とする

`ResolvedProject` build完了後、indexは読み取り専用とする。

- rendererがindexを更新しない
- QueryServiceがindexを更新しない
- MCP wrapperがindexを更新しない
- lazyにindexを増やす場合でも、外部から観測される意味がbuild時indexと一致することを必須とする

初期実装では、実装を単純にするためeager buildを採用する。

### 4. common base index

全view / queryで共通して使うbase indexとして、以下を構築する。

| index | key | value | 用途 |
|---|---|---|---|
| `nodesByQID` | `QualifiedID` | resolved node | `get_signature` / 各rendererの参照解決 |
| `nodesByFile` | file ID | resolved node list | 1ファイル=1DAG / 1ファイル=1FSMのrender入力 |
| `mainNodeByFile` | file ID | main task node | DAG render対象の特定 |
| `filesByTopLevelModule` | top-level module | file list | render_index.yaml group解決 / output配置 |
| `modelsByQID` | `QualifiedID` | resolved model node | params / returns / store.model の型参照 |
| `storesByQID` | `QualifiedID` | resolved store node | reads/writes / ER / Sequence DB participant |
| `tasksByQID` | `QualifiedID` | resolved task node | DAG / API / QueryService |
| `actorsByQID` | `QualifiedID` | resolved actor node | Sequence external participant |

`nodesByQID` が主indexであり、type別indexはtype assertionやfilterを散らさないための補助indexとする。

`filesByTopLevelModule` のtop-level moduleは、ADR-045の `render_index.yaml` の `groups[].modules` で指定可能な最上位moduleを指す。nested moduleは親top-level moduleのkeyに集約し、値側のfile listに含める。

### 5. initial reverse lookup index

ADR-048時点で必須とする初期reverse lookup indexは以下とする。

| index | key | value | 主な利用者 |
|---|---|---|---|
| `referencesBySource` | source object ID | reference list | `get_references(direction=out)` / `inspect(model)` / `inspect(task)` |
| `referencesByTarget` | target object ID | reference list | `get_references(direction=in)` / `inspect(model)` / `inspect(store)` / `inspect(task)` |
| `tasksReadingStore` | store `QualifiedID` | task `QualifiedID` list | `inspect(store)` / ER変化影響 / Sequence DB操作 |
| `tasksWritingStore` | store `QualifiedID` | task `QualifiedID` list | `inspect(store)` / ER変化影響 / Sequence DB操作 |
| `transitionsByStateEventGuard` | `(stateFileID, fromStateID, eventID, guard?)` | transition ref | Sequence scenario解決 / State render |
| `actionsByTask` | task `QualifiedID` | transition ref list | `inspect(task)` / `get_references(task)` / Sequence逆引き |
| `scenariosByID` | scenario ID | resolved scenario | MCP query / Sequence render |

`referencesBySource` / `referencesByTarget` は、ADR-049の `references` schemaに対応する汎用semantic reference indexである。
対象には、task params / returns / reads / writes、store.of、model field type / fk、event payload / actor / watches、transition from / to / event / action、scenario step transition等のdirect referenceを含める。
MCP v1ではdirect referenceのみをindexし、transitive closureは持たない。

`tasksReadingStore` / `tasksWritingStore` / `actionsByTask` は、rendererや特定queryで頻繁に使う用途別indexとして残す。
ただし、MCP外部schemaの中心語彙は `references` であり、`get_references` / `inspect` の一般的な参照返却は `referencesBySource` / `referencesByTarget` を主材料にする。

`transition ref` は、transition本体への参照に加えて、定義元state file、source location、resolved action taskがあればその `QualifiedID` を持つ。

### 6. transition indexの詳細

FSM transitionはADR-035により、同一 `(from, on)` に複数transitionを持てる。
そのため、transition indexではguardの有無をkeyに含める。

```text
(stateFileID, fromStateID, eventID, guard?) -> transition ref
```

解決ルールは以下とする。

- guardなしtransitionは `guard = nil` としてindexする
- guardありtransitionはguard文字列のexact matchでindexする
- guard文字列はYAML decode後の文字列をそのままkeyにする。trim、空白正規化、Unicode正規化、式AST比較は行わない
- 同一keyが複数存在した場合は semantic validation error
- scenario step解決では、まず `(stateFileID, fromStateID, eventID)` で候補を確認し、必要に応じてguard exact matchで1件に特定する

実装上は、候補列挙用に `(stateFileID, fromStateID, eventID) -> transition ref list` の補助indexを持ってよい。
ただし外部的な意味は `transitionsByStateEventGuard` の一意特定ルールに従う。

### 7. actionsByTaskの詳細

`actionsByTask` は、どのFSM transitionが特定taskをactionとして呼ぶかを逆引きするindexである。

```text
task QualifiedID -> transition ref list
```

対象は、semantic buildで `transition.action` がtask nodeに解決できたtransitionのみとする。
`action` が省略されたtransitionはindexしない。

このindexにより、`inspect(task)` は「このtaskはどのstate transitionから呼ばれるか」をRaw YAML走査なしで返せる。
またSequence Diagram rendererは、scenario stepから解決したtransitionを通じてaction taskへ到達できる。

### 8. scenariosByIDの詳細

`scenariosByID` は、`as: sequence_diagram` のview fileをscenario IDで取得するindexである。

```text
scenario ID -> resolved scenario
```

制約は以下とする。

- project内でscenario IDは一意（ADR-032）
- 重複した場合は semantic validation error（ADR-032の一意制約をsemantic buildで強制する）
- scenario内の `state_file` はcanonical file IDへ解決済みで保持する
- scenario stepsはtransition refへ解決済みで保持する

Sequence Diagram rendererは、scenario IDからscenarioを取得し、各stepのtransition refを辿って描画材料を集める。

### 9. 依存グラフ全体は初期indexにしない

`get_references` 用の完全なtransitive reference graphはADR-048では事前構築しない。

初期実装では、以下のbase / reverse indexを組み合わせてdirect referenceを返す。

- task params / returns / reads / writes
- transition action
- scenario steps
- model field references
- store.model

transitive closureのcacheや依存グラフ専用indexは、QueryService vertical sliceで必要になった時点で追加ADRまたは実装判断として扱う。

### 10. Rendererがindexから読む範囲

RendererはRaw YAML structsを読まない。

Rendererが使う入力は以下のいずれかとする。

1. `ResolvedProject` のindex / accessor
2. `ResolvedProject` のindex / accessorから構築されたview-specific view model

ただし、rendererが表示対象nodeのresolved fieldを読むことは許可する。
禁止するのは、Raw YAML structsの走査、未解決文字列の再解釈、reverse lookupの再構築である。

### 11. 今回決めないこと

本ADRでは以下を決めない。

- Goの具体的なstruct名 / package名
- mapの具体型、slice順序の実装詳細
- MCP toolのinput / output schema
- view-specific view modelの具体型
- transitive dependency graphの完全な事前構築
- indexの永続化やキャッシュファイル化

## 理由

### build時indexにする理由

indexを利用時に都度作ると、renderer / QueryService / MCP wrapperの各所に探索ロジックが分散する。
特に `reads` / `writes` / `transition.action` / scenario step解決は複数viewで使われるため、各viewで個別実装すると解釈差分が生まれる。

build時に一度だけindexを作ることで、validationとindex構築を同じ意味解決結果に基づかせられる。

### reverse lookupを明示的に持つ理由

YAMLは基本的に「taskがstoreを読む」「transitionがtaskを呼ぶ」のような順方向参照で書かれる。
一方、MCP queryでは「このstoreを読むtaskは？」「このtaskを呼ぶtransitionは？」のような逆方向問い合わせが重要になる。

逆方向問い合わせを毎回全node走査で実装すると、MCP queryの中心価値であるinspectが不安定になる。
そのため、reverse lookup indexを `ResolvedProject` のfirst-classな派生情報として持つ。

### 初期indexを絞る理由

最初から全依存グラフや全view用のindexを作ると、まだ実装していないrendererの都合まで先取りすることになる。
ADR-048では、Milestone 1〜3で必要になる最小限のindexに絞る。

- DAG vertical sliceに必要な `nodesByQID` / `nodesByFile` / `mainNodeByFile` / store access index
- State / Sequence解決に必要なtransition index
- MCP `inspect` / `get_references` に必要な汎用reverse lookup index
- Sequence scenario取得に必要な `scenariosByID`

### ResolvedProjectのfield直読みを禁止しない理由

すべての読み取りをindex経由に限定すると、実装が不自然になる。
例えば `task.params` や `task.returns` は、そのtask nodeを取得した後にresolved fieldとして読むのが自然である。

禁止すべきなのは、fieldを読むことではなく、application-facing layerがRaw YAMLを再走査したり、未解決文字列を再解釈したり、逆引きをその場で作ることである。

## 影響

- `ResolvedProject` build pipelineは、semantic validation後にbase indexとreverse lookup indexを構築する。
- `QueryService` は `ResolvedProject` のindex / accessorを使って `GetSignature` / `Inspect` / `GetReferences` を実装する。
- RendererはRaw YAML structsではなく、index / accessorまたはview-specific view modelを使う。
- Milestone 1のDAG vertical sliceでは、最低限 `nodesByQID` / `nodesByFile` / `mainNodeByFile` / `tasksReadingStore` / `tasksWritingStore` を実装する。
- Milestone 3のQueryService vertical sliceでは、本ADRのreverse lookup indexを利用する。特に `GetReferences` / MCP `get_references` は `referencesBySource` / `referencesByTarget` を主材料にする。
- `docs/spec/mcp/overview.md` および `docs/spec/mcp/tools/*.md` では、本ADRのindex方針を前提に `get_signature` / `get_references` / `inspect` の外部I/Oを定義する。

## Evidence
- commit: da28fa9
- impl commit: tbd
- 参考: compiler symbol table / semantic model、Goのmap-based index慣習

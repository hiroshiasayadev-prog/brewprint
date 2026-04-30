# 054: MCP query coverage は設計対話を基準に拡張する

- **status**: accepted
- **date**: 2026-04-30
- **depends on**: ADR-047, ADR-048, ADR-049, ADR-050

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

brewprint の MCP / QueryService は、当初は実装補助のために `get_signature` / `get_references` / `inspect` / `list_endpoints` を提供する query layer として整備された。
ADR-047 で Raw YAML structs / ResolvedProject / QueryService / Renderer の責務境界を固定し、ADR-048 で ResolvedProject index strategy、ADR-049 で reference 語彙を固定した。

M9 / Post-M9 までに、MCP は以下のように拡張されている。

- state / event / scenario inspect
- transition object の references / signature / inspect
- field object の references / signature / inspect
- file object の references 一部
- same-module bare FK normalization

一方で、brewprint の目的は単に YAML から Mermaid / HTML / Markdown を生成することではない。
brewprint は、人間と LLM が共通の設計表現を見ながら会話するための中間表現である。

具体的には、利用者は以下のような会話をしたい。

```text
このDAGの pending_order はどこで作られて、どこで使われる？
このER図にこのmodelが出ないのはなぜ？
このAPI Tableはどのmoduleを集約している？
このstateから出るtransitionはどれ？
このfieldを変えると何に影響する？
このYAML fileには何が定義されている？
```

このユースケースでは、MCP は単なる実装補助APIでは足りない。
DAG / State Diagram / Sequence Diagram / ER / API Table / Wireframe など、図やviewに現れるsemantic objectを、LLMが直接問い合わせられる必要がある。

## 決定

MCP / QueryService の query coverage は、今後 **設計対話 coverage** を基準に拡張する。

ここでいう設計対話 coverage とは、利用者が render された図・表・view を見ながら、LLMに対して対象要素の意味、参照関係、生成元、影響範囲を質問できる範囲を指す。

### 1. MCPは設計対話用 query layer である

brewprint MCP は、単なる Raw YAML 確認APIでも、renderer の内部補助APIでもない。
LLMが ResolvedProject 上の semantic object を辿り、人間との設計対話を支援するための query layer と位置づける。

MCP response は引き続き Raw YAML AST を直接公開しない。
ADR-047 の境界を維持し、MCP は QueryService を通じて ResolvedProject 上の情報を返す。

```text
YAML files
  ↓ load / classify / decode
Raw YAML structs
  ↓ validate / name resolution / derived model build / index build
ResolvedProject
  ↓
QueryService
  ↓
MCP response
```

### 2. 図やviewに現れるsemantic objectは原則query可能にする

DAG / State Diagram / Sequence Diagram / ER / API Table / Wireframe などの render に現れる主要なsemantic objectは、原則として MCP から問い合わせ可能にする。

対象例:

- task / model / store / state / event / actor
- model field
- transition
- sequence scenario view
- API Table view
- ER Diagram view
- implicit asset
- file-local sub task / branch / fork / join
- flow entry / flow wiring
- source file

すべてを一度に実装する必要はない。
ただし、MCP拡張の判断では「そのobjectが図やview上で利用者に見えており、会話対象になりうるか」を重視する。

### 3. project探索・file理解・view理解を優先する

LLMが project を自律的に理解するには、個別object問い合わせだけでは不足する。
そのため、以下の coverage を段階的に拡張する。

- project内objectを探索する機能
- file単位で定義内容を把握する機能
- view定義が何を集約しているかを把握する機能
- 図上の個別要素を直接queryする機能
- 変更影響を直接referenceより深く辿る機能

具体的な実装順序は `docs/TASKS.md` で管理する。

### 4. direct references v1方針は維持しつつ、impact traversalを将来拡張する

ADR-049 / `docs/spec/mcp.md` で定義した通り、現在の `get_references` は direct references を返す。
この方針は維持する。

ただし、設計変更相談では transitive impact traversal が必要になるため、将来的には以下のどちらかで拡張する。

- `get_reference_tree` のような別toolを追加する
- `get_references` に depth 指定を追加する

どちらを採用するかは、実装時点のQueryService設計と利用感を見て決める。

### 5. Raw YAML source accessは補助機能として扱う

`get_source` 相当の機能は将来候補とする。
ただし、これは Raw YAML AST を公開するという意味ではない。
MCP上で対象semantic objectに対応するsource snippetを確認する補助機能として扱う。

source snippet を返す場合でも、QueryService / source mapping を通じて、semantic objectに紐づく範囲を返す。
Raw YAML全体を無構造に公開するAPIにはしない。

## 理由

### 設計対話 coverage を採用する理由

brewprint の価値は、人間と LLM が同じ設計表現を見ながら会話できる点にある。
render は人間向けの視覚化であり、MCP は LLM向けの構造化queryである。

人間が図上の要素を見て質問したとき、LLMがその要素をqueryできなければ、図とMCPの間に断絶が生まれる。

例:

- DAGに asset node が出るのに、asset を直接queryできない
- ER図が横断viewとして表示されるのに、ER viewが何を含むかinspectできない
- API Tableが route を表示するのに、そのview定義が何を集約しているかqueryできない
- state fileを見ても、file単位で何が定義されているかMCPで確認できない

これらは、brewprintを設計対話の基盤にする上で中核的な欠落となる。

### Raw YAML AST公開を避ける理由

MCPがRaw YAML ASTを直接公開すると、以下の問題が起きる。

- name resolution 済みのIDではなく、未解決文字列をLLMが読んでしまう
- validation / derived model / implicit asset / reference index を経由しないため、semantic contractが崩れる
- renderer / QueryService / source layer の責務境界が曖昧になる

そのため、MCPは引き続き ResolvedProject 上の semantic object を返す。
YAML source は必要に応じて補助情報として返すに留める。

### 具体実装をTASKSで管理する理由

本ADRは、MCP拡張の判断基準を固定するための方針ADRである。
個別toolやselectorのschemaは、実装しながら `docs/spec/mcp.md` に追随する。

最初から全selector / 全toolをADR本文に固定すると、実装しづらくなる。
そのため、ADRでは方向性と原則を定め、具体的なmilestoneは `docs/TASKS.md` で管理する。

## 影響

### MCP / QueryServiceへの影響

今後のMCP拡張では、以下が優先候補になる。

- `list_objects`
- `inspect(file)`
- `inspect(view: api_table)`
- `inspect(view: er_diagram)`
- implicit asset selector
- private sub node selector
- flow entry / flow wiring references
- `get_source`
- `get_reference_tree` または depth指定つきreference traversal

これらは、設計対話 coverage を広げるための候補であり、すべてを即時実装することを要求するものではない。
実装順序は `docs/TASKS.md` で管理する。

### specへの影響

`docs/spec/mcp.md` は、実装済みselector / toolの範囲を明確にする必要がある。
特に、selector support matrix を追加し、以下を区別する。

- supported
- partial
- future
- reference target only
- v1 optional

### renderer / view specへの影響

renderに現れるobjectをMCPでquery可能にするため、view specとMCP specの対応関係を今後整理する必要がある。

例:

- DAGの asset node ↔ asset selector / asset references
- DAGの flow wiring ↔ flow_step / flow_param references
- ER Diagram view ↔ ER view inspect
- API Table view ↔ API table inspect / list_endpoints
- wireframe element ↔ 将来のwireframe element selector候補

### 実装境界への影響

ADR-047の境界は維持する。
MCP拡張によって `internal/mcp` が Raw YAML や renderer を直接読むことは許容しない。
必要な情報は QueryService / semantic model / index に追加する。

## Evidence
- commit: 0ea083b
- impl commit: tbd
- 参考: 特になし

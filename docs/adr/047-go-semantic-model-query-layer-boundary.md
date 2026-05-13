# 047: Go semantic model / query layer boundary

- **status**: accepted
- **date**: 2026-04-27
- **supersedes**:

※ ADR-049 により `GetDeps` / `get_deps` は `GetReferences` / `get_references` に名称変更。

## 背景

ADR-001〜046でbrewprintのYAML仕様、名前解決、各viewのrender規則、render出力配置はおおむね確定した。
次の段階ではGo実装に入るが、実装境界を曖昧にしたまま進めると、以下の問題が起きる。

- renderer が Raw YAML structs を直接走査し、独自に名前解決や派生情報構築を始める
- MCP wrapper / QueryService が renderer とは別に Raw YAML structs を走査し、同じ解決処理が二重実装される
- YAML decode、validation、name resolution、derived model build、render、query の責務が混ざる
- 後続の `get_signature` / `get_deps` / `inspect` が、YAMLの形に強く依存した不安定なAPIになる

brewprintは、人間向けにはMermaid/HTML renderを、LLM向けにはMCP queryを提供する。
どちらも同じ設計情報から派生するため、Go実装では Raw YAML をそのまま各利用者に読ませるのではなく、解決済みの意味モデルを境界として共有する必要がある。

## 決定

### 1. Go実装の主要境界

Go実装は以下の流れを基本構造とする。

```text
YAML files
  ↓ load / classify / decode
Raw YAML structs
  ↓ validate / name resolution / derived model build
ResolvedProject
  ↓
QueryService / Renderer
  ↓
MCP output / render output
```

`ResolvedProject` を、Go実装内の application-facing layer が読む中心モデルとする。
ここでいう application-facing layer とは、brewprint YAML上のApplicationレイヤーではなく、Go実装上の `QueryService` / `Renderer` / `MCP wrapper` など、外部出力に近い層を指す。

### 2. Raw YAML structs の責務

Raw YAML structs は、YAMLをGoで受け取るための入力構造に限定する。

Raw YAML structs が担う責務は以下とする。

- YAML decode後のフィールド保持
- YAMLファイル種別ごとの最小構造表現
- source file path / line number など、診断に必要なsource location保持
- YAML形状に由来する最低限の構文エラー検出

Raw YAML structs は以下を担わない。

- 名前解決済みIDの提供
- module nesting / actor global / sentinel などを考慮した参照解決
- implicit asset の生成
- task / store / model / state / scenario の横断index提供
- rendererやMCP向けの問い合わせAPI提供

### 3. ResolvedProject の責務

`ResolvedProject` は、Raw YAML structs から構築される解決済みの意味モデルである。

`ResolvedProject` build では以下を実行する。

- 必須フィールド、型、参照形式、構造制約のvalidation
- QualifiedID parser / symbol table によるname resolution
- 同一module内ID、クロスモジュールフルパス、module nesting、actor global、sentinel方式の解決
- task `returns` からの implicit asset 構築
- task params / returns / reads / writes / flow step / state transition / scenario などの意味的な接続関係の構築
- renderer / QueryService が安定して読めるnode registryと参照関係の提供
- node registryに加え、reverse lookup index等の派生indexを構築する（詳細はADR-048）
- validation error / warning にsource locationを付与できる情報の保持

`ResolvedProject` は、Raw YAMLの記述形に依存した低レベル構造ではなく、brewprintの意味に沿って問い合わせ可能なモデルとする。

### 4. Renderer の入力境界

Renderer は Raw YAML structs を直接読まない。

Renderer は以下のいずれかを入力とする。

1. `ResolvedProject`
2. `ResolvedProject` から構築された view-specific view model

view-specific view model は、render対象のviewに必要な情報だけを持つ読み取り専用モデルである。
具体的なview model型は、最初のDAG vertical slice実装中に必要性を見て固める。

Renderer 内で許容される処理は、表示形式への変換、並び順の決定、Mermaid/HTML/Markdownの生成である。
Renderer 内で名前解決、Raw YAML走査、implicit asset生成、横断依存解析を再実装してはならない。

### 5. QueryService の入力境界

`QueryService` は `ResolvedProject` を読む問い合わせ層とする。

`QueryService` は、MCP tool から直接呼ばれることを想定するが、MCP transportには依存しない通常のGo APIとして実装する。

`QueryService` が担う責務は以下とする。

- `GetSignature`
- `GetDeps`
- `Inspect`
- その他、LLM向けに必要な問い合わせの集約

`QueryService` は Raw YAML structs を直接読まない。
また、名前解決や依存解析の基本材料は `ResolvedProject` に存在するものを使う。
reverse lookup index の具体戦略はADR-048で別途決める。

### 6. MCP wrapper の責務

MCP wrapper は transport / protocol adapter として扱う。

MCP wrapper が担う責務は以下に限定する。

- MCP request のparse
- input validationのうち、transport境界で必要な形式チェック
- `QueryService` 呼び出し
- `QueryService` の結果をMCP responseへ変換
- protocol-level errorへの変換

MCP wrapper は Raw YAML structs を直接読まない。
MCP wrapper 内でname resolution、dependency traversal、renderer相当の整形を実装してはならない。

### 7. validation / name resolution / derived model build の分担

Go実装では、validationを以下の2段階に分ける。

1. **decode-time validation**
   - YAMLとして読めるか
   - ファイル種別を判定できるか
   - 必須トップレベルキーの形が大きく壊れていないか

2. **semantic validation**
   - 参照先が存在するか
   - QualifiedIDとして解決できるか
   - module nesting / actor global / sentinel方式に従っているか
   - flow / transition / scenario / view が仕様上成立しているか
   - render_index.yaml のgroup解決が成立しているか

name resolution と derived model build は semantic validation と同じbuild pipelineで実行する。
以後の `Renderer` / `QueryService` / `MCP wrapper` は、このbuild pipelineを通過した `ResolvedProject` を前提にする。

### 8. 今回決めないこと

本ADRでは以下を決めない。

- `ResolvedProject` が保持する具体的なreverse lookup index一覧
  - ADR-048で決める
- `get_signature` / `get_deps` / `inspect` の具体的なinput / output schema
  - `docs/spec/mcp/overview.md` および `docs/spec/mcp/tools/*.md` で決める
- DAG / State / Sequence / ER / API / Wireframeごとの具体的なview model型
  - 各rendererのvertical slice実装中に固める
- Go package名、struct名、interface名の最終形
  - 実装時に本ADRの境界を守る範囲で決める

## 理由

### Raw YAML structsをapplication-facing layerから隠す理由

Raw YAML structs は入力形式そのものであり、brewprintの意味モデルではない。
Raw YAML structs を renderer や MCP wrapper に読ませると、各所で以下の処理が必要になる。

- short idからQualifiedIDへの解決
- module nestingの解釈
- actor globalの特別扱い
- task returnsからのimplicit asset生成
- reads / writes / transitions / scenariosの横断走査

これらを複数箇所に分散させると、render結果とMCP query結果が同じYAMLから異なる解釈を返す危険がある。
そのため、Raw YAML structs を読む層を loader / decoder / resolver / builder に限定する。

### ResolvedProjectを中心にする理由

brewprintは、同じYAMLから複数のviewとMCP queryを導出する。
そのため、各出力系が個別にYAMLを解釈するより、解決済みの意味モデルを1回構築し、それを共有する方が一貫性が高い。

`ResolvedProject` は、YAMLの表記ゆれや参照解決の複雑さを吸収し、Renderer / QueryService がbrewprintの意味だけを扱える境界になる。

### QueryServiceをMCP wrapperから分離する理由

MCPは外部公開手段であり、問い合わせロジックそのものではない。
`QueryService` を通常のGo APIとして分離することで、以下が可能になる。

- MCP transportなしでunit testできる
- CLIやgolden testから同じ問い合わせを直接叩ける
- MCP protocol変更時に意味解析ロジックへ影響を出さない

### view-specific view modelを今決めきらない理由

Rendererが `ResolvedProject` を直接読めば、初期実装は単純になる。
一方、DAGやSequenceなどは表示用に整形済みのview modelを挟んだ方が見通しが良くなる可能性がある。

現時点では具体的な型を決める材料が不足しているため、ADR-047では「RendererはRaw YAMLを読まない」という境界だけを固定し、view modelの必要性と型はDAG vertical slice中に判断する。

## 影響

- Go実装は Raw YAML structs / ResolvedProject / QueryService / Renderer / MCP wrapper の責務を分けて進める。
- Renderer と MCP wrapper は Raw YAML structs を直接importしない方針とする。
- 名前解決、implicit asset生成、横断依存解析は `ResolvedProject` build pipeline に集約する。
- DAG vertical sliceでは、まず `auth.task.login` を Raw YAML → ResolvedProject → Renderer の経路で通す。
- QueryService vertical sliceでは、`ResolvedProject` を入力に `GetSignature` / `Inspect` / `GetDeps` を実装する。
- ADR-048で `ResolvedProject` のindex strategyを続けて決める必要がある。
- `docs/spec/mcp/overview.md` および `docs/spec/mcp/tools/*.md` でMCP toolの外部仕様を別途定義する必要がある。

## Evidence
- commit: da28fa9
- impl commit: tbd
- 参考: Goのlayering慣習

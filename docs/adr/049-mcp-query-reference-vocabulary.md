# 049: MCP / QueryService の reference 語彙統一

- **status**: accepted
- **date**: 2026-04-27
- **supersedes**: ADR-047, ADR-048
- **depends on**: ADR-047, ADR-048

## 背景

ADR-047では、QueryService の責務として `GetSignature` / `GetDeps` / `Inspect` を仮置きした。
ADR-048でも、MCP tool が `get_deps` を返却材料として使う前提の記述が残っている。

しかし `docs/spec/mcp/overview.md` / `docs/spec/mcp/schema.md` / `docs/spec/mcp/tools/get-references.md` の設計中に、brewprint MCPが返す関係は単なる dependency より広いことが明確になった。

例:

- task params / returns が model を参照する
- task が store を reads / writes する
- transition が event / state / action task を参照する
- sequence scenario step が transition を参照する
- model field が type / fk を参照する
- event が actor / payload model / watched store を参照する

これらは広義には依存と呼べるが、LLMが読む外部APIとしては `references` の方が自然である。

また、MCP APIはPython `inspect` に近い操作感を採用する方針とした。
その語彙体系では `signature` / `doc` / `source` / `members` / `references` が自然に対応する。

## 決定

### 1. 外部MCP tool名は `get_references` とする

`get_deps` は外部MCP tool名として採用しない。

```text
NG: get_deps
OK: get_references
```

`get_references` は、対象objectが参照するもの、および対象objectを参照するものを返す。

方向は `direction` で指定する。

```json
{
  "selector": { "id": "auth.task.login" },
  "direction": "both"
}
```

### 2. QueryService内部APIも `GetReferences` に統一する

ADR-047で仮置きされた `GetDeps` は廃止し、Go実装上のQueryService method名も `GetReferences` とする。

```go
GetSignature(...)
GetReferences(...)
Inspect(...)
```

外部MCP toolと内部QueryServiceで別名を使わない。
命名差分による実装者・LLMの混乱を避ける。

### 3. MCP responseの中心語彙を `references` とする

MCP responseでは、依存・参照・逆参照を `references` として返す。

```json
{
  "references": [
    {
      "kind": "reads",
      "direction": "out",
      "from": { "id": "auth.task.login", "kind": "task" },
      "to": { "id": "auth.store.user_db", "kind": "store" }
    }
  ]
}
```

`edges` / `deps` / `dependencies` はMCP外部schemaの中心語彙として使わない。

### 4. `references` はsemantic relationであり、graph edgeに限定しない

`references` はDAG edgeだけを意味しない。

以下もreferenceに含める。

- model field type
- model field fk
- store.of
- task params / returns
- task reads / writes
- transition from / to / on / action
- scenario step transition
- event payload / actor / watches

つまり `references` は、ResolvedProject上のsemantic object間の直接関係を表す総称である。

### 5. v1ではdirect referencesのみ返す

ADR-048の方針通り、v1ではtransitive dependency graphを事前構築しない。

`GetReferences` / `get_references` は直接referenceのみを返す。
transitive traversalが必要になった場合は、将来のspec更新または別ADRで扱う。

本ADRは、MCP外部schema / QueryService API の語彙統一を扱う。
Renderer内部 / Raw YAML structs / DAG render output（Mermaid edge label等）の用語は対象外とし、それらでは `edge` / `dependency` 等の語を必要に応じて使ってよい。

## 理由

### `deps` より `references` が適切な理由

`deps` / `dependencies` は、実行順序・ビルド依存・package依存のような意味に寄りやすい。
brewprint MCPが返す関係はそれより広い。

例えば `model field -> primitive type` や `scenario step -> transition` は、依存というより参照である。
これらを `dependency` と呼ぶと、LLMが実行依存やDAG依存として誤読する可能性がある。

`references` は、IDEやPython introspectionの文脈でも自然であり、LLMが追加説明なしに理解しやすい。

### 外部名と内部名を揃える理由

外部MCP toolを `get_references`、内部QueryServiceを `GetDeps` とすると、同じ概念に2つの名前が生じる。

brewprintはAI実装を前提にしているため、用語差分はそのまま実装時の迷いになる。
外部schemaとGo APIの命名を揃え、`GetReferences` に統一する。

### ADR-047 / ADR-048を全面supersedeしない理由

ADR-047のQueryService境界、ADR-048のResolvedProject index strategyは継続有効。
変更対象は `GetDeps` / `get_deps` という命名部分のみである。

そのため、本ADRはADR-047/048全体をsupersedeせず、命名部分だけを上書きする。

## 影響

- `docs/spec/mcp/tools/get-references.md` は `get_references` を正式tool名として定義する。
- `docs/spec/mcp/schema.md` のresponse語彙は `references` を中心にする。
- ADR-047の `GetDeps` 記述は `GetReferences` と読み替える。
- ADR-048の `get_deps` / `GetDeps` 記述は `get_references` / `GetReferences` と読み替える。
- TASKS.md 等に `get_deps` 表記があれば `get_references` に更新する。
- Go実装のQueryService methodは `GetReferences` とする。

## Evidence
- commit: 870feb1
- impl commit: tbd
- 参考: Python inspect / IDE references / compiler semantic model

# UC-002 Phase A: MCP公開contract blueprint方針

このメモは、M14 Phase Aで MCP公開contract を brewprint YAML に落とす際の表現方針を固定する。

参照元:

- `docs/tasks/m14-self-hosting.md` §Phase A
- `docs/spec/mcp/overview.md`
- `docs/spec/mcp/schema.md`
- `docs/spec/mcp/errors.md`
- `docs/spec/mcp/tools/*.md`
- V01-ADR-047 / V01-ADR-048 / V01-ADR-049 / V01-ADR-054 / V01-ADR-055 / V01-ADR-056

---

## 1. 表現対象

Phase Aで扱うMCP公開contractは、MCP v1の外部I/O契約に限定する。

対象:

- MCP tool名
- tool request / response model
- 共通schema model
- MCP client / server actor
- tool呼び出しからQueryService境界までの概略flow

対象外:

- Go package / struct / interfaceの具体名
- MCP transport実装の詳細
- Raw YAML decode struct
- `ResolvedProject` 内部indexの具体実装
- renderer内部のMermaid / HTML / Markdown出力詳細

---

## 2. node種別への対応

| MCP contract要素 | brewprint表現 | 理由 |
|---|---|---|
| MCP client | `actor` | LLM / MCP host / Claude Code など、brewprint外からtoolを呼ぶ外部主体 |
| MCP server | `actor` | MCP requestを受け取る外部境界。Sequence等で外部主体として扱える |
| MCP tool | `task` | requestを受け取りresponseを返す処理単位。QueryServiceを呼ぶ公開contract単位 |
| request / response schema | `model` | JSON schemaの構造をData layerの型として表現する |
| common schema | `model` | ObjectSelector / ObjectRef / Reference / Diagnostic等は複数toolで共有される型 |
| QueryService境界 | `task`のflow / note | v1にはservice/component nodeがないため、fake actor/modelにはしない |
| ResolvedProject | `store kind: context` | 各tool taskが読み取る解決済みproject contextとして表す |
| MCP tool invocation event | Phase Aでは原則未使用 | Phase Aは公開I/O contract中心。FSMやsequence化が必要になった時にevent化する |

---

## 3. MCP toolはHTTP endpointではない

MCP toolを `task` として表すが、`endpoint: true` は付けない。

理由:

- `endpoint: true` は `docs/spec/nodes.md` 上、HTTP API Table集計対象を意味する
- MCP toolはHTTP routeではなく、MCP protocol上のtool callである
- `method` / `path` を無理に割り当てると、API Table viewがHTTP APIとして誤読する

したがって、Phase AのMCP tool taskは以下の形を基本とする。

```yaml
nodes:
  - id: get_signature
    type: task
    main: true
    params:
      - name: request
        model: get_signature_request
    returns:
      name: response
      model: get_signature_response
    note: |
      MCP tool `get_signature` の公開contract。
      HTTP endpointではないため endpoint: true は付けない。
```

---

## 4. tool taskの内部flow

各MCP tool taskは、公開contract上は1つのmain taskとして扱う。

ただしDAGでMCP serverからQueryService境界までの概略を見えるようにするため、各tool file内に以下のfile-local sub taskを置く。

1. `validate_request`
2. `query_service`
3. `build_response`

この3段は実装package名ではなく、外部contractから見た概念的な境界である。

```text
request
  -> validate_request
  -> query_service      # ResolvedProject / QueryService境界
  -> build_response
  -> response
```

QueryServiceはv1のnode種別では表せないため、独立nodeにはしない。
`query_service` sub taskの `reads: [resolved_project_store]` と `note` で境界を表す。

---

## 5. JSON schemaの表現粒度

MCP specのJSON schemaには、optional field / enum / union / discriminated object / arbitrary JSON object が含まれる。

v1のbrewprint `model` は `struct` / `list` / `dict` と primitive / model参照を表せるが、以下を型レベルでは直接表現できない。

- optional / required の細かい区別
- enum値集合
- `oneOf` / union
- discriminated objectのkind別payload制約
- JSON objectの任意shape

Phase Aでは、これらを以下の暫定ルールで表す。

- 型として表せる部分は `model` field に落とす
- optional / enum / union / validation制約は `note` に明記する
- 任意JSON payloadは `any` を使い、`note` で意味を補足する
- v1表現では情報が落ちる箇所は `TASKS-UC-002.md` のspec gapとして追跡する

---

## 6. Phase AのYAML配置

```text
yaml/
  actors.yaml
  mcp/
    model/
      common.yaml
      tools.yaml
    store/
      resolved_project_store.yaml
    task/
      list_objects.yaml
      get_signature.yaml
      get_source.yaml
      get_references.yaml
      get_reference_tree.yaml
      analyze_impact.yaml
      inspect.yaml
      list_endpoints.yaml
  views/
    er.yaml
```

`render_index.yaml` の初期groupは `mcp` moduleを対象にしているため、上記 `yaml/mcp/**` は `mcp` groupに出力される。

---

## 7. ADR起票判断

現時点では新規ADRは起票しない。

理由:

- MCP toolを `task` として扱う判断は既存の `task = 処理単位` 仕様と整合する
- `endpoint: true` を使わない判断は、既存のHTTP endpoint仕様の自然な帰結である
- QueryServiceを独立nodeにしない判断は、v1 node種別の範囲内での表現方針に留まる

ただし、将来 `service` / `component` / `protocol_endpoint` のようなnode種別を追加する場合は、v2向け構造変更としてADR起票対象になる。

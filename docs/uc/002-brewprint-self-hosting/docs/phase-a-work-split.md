# UC-002 Phase A: MCP公開contract 並列作業分割

このメモは、M14 Phase A（MCP公開contractのblueprint化）を複数レーンで並列に進めるための作業分割を定義する。

参照元:

- `docs/tasks/m14-self-hosting.md` §Phase A
- `docs/uc/002-brewprint-self-hosting/docs/phase-a-mcp-contract.md`
- `docs/spec/mcp/overview.md`
- `docs/spec/mcp/schema.md`
- `docs/spec/mcp/errors.md`
- `docs/spec/mcp/tools/*.md`

---

## 1. 並列化の基本方針

Phase Aは以下の理由により並列化できる。

- MCP toolごとのrequest / response schemaは独立している
- MCP toolごとのtask YAMLも独立している
- 共通schema modelを先に固定すれば、各tool YAMLは同じ命名規約で作れる
- render生成・coverage整理は後段に集約できる

ただし、並列作業前に以下を固定する。

1. 共通model名
2. tool別 request / response model名
3. task id
4. file配置
5. v1 modelで表現できないJSON schema制約の扱い

---

## 2. レーン分割

### Lane A: 方針 / gap整理

目的:

- MCP公開contractをbrewprint YAMLへ写像する方針を維持する
- v1表現では落ちる情報をspec gapとして記録する
- 並列作業中の命名・配置のずれを防ぐ

担当ファイル:

- `docs/uc/002-brewprint-self-hosting/docs/phase-a-mcp-contract.md`
- `docs/uc/002-brewprint-self-hosting/docs/phase-a-work-split.md`
- `docs/uc/002-brewprint-self-hosting/TASKS-UC-002.md`
- `docs/uc/002-brewprint-self-hosting/editor-viewer-notes.md`（必要時のみ）

主な作業:

- `task` / `model` / `actor` / `store` への対応表を維持する
- optional / enum / union / discriminated object の扱いを明記する
- v2向け構造変更候補を `TASKS-UC-002.md` に追記する
- 新規ADRが必要かを判断する

完了条件:

- Phase Aの表現方針がdoc化されている
- spec gap候補が未記録のまま放置されていない

---

### Lane B: 共通model

目的:

- MCP tool間で共有されるschema modelを定義する

担当ファイル:

```text
yaml/mcp/model/common.yaml                 # object_selector
yaml/mcp/model/source_location.yaml
yaml/mcp/model/object_ref.yaml
yaml/mcp/model/diagnostic.yaml
yaml/mcp/model/diagnostic_list.yaml
yaml/mcp/model/reference.yaml
yaml/mcp/model/reference_list.yaml
yaml/mcp/model/string_list.yaml
yaml/mcp/model/resolved_project.yaml
yaml/mcp/model/mcp_error.yaml
```

注意:

- 原則として `model/` 配下は 1ファイル = 1 main model とする
- 複数の補助modelを1ファイルに詰めない
- 再帰型・union・任意JSON objectは `any` + `note` で暫定表現する
- `any` を使った箇所は、必要に応じてspec gapへ記録する

完了条件:

- tool別modelが参照する共通modelが存在する
- `object_selector` / `object_ref` / `reference` / `diagnostic` が最低限揃っている

---

### Lane C: tool別 request / response model

目的:

- MCP v1 toolごとの公開I/O schemaをmodel化する

担当ファイル:

```text
yaml/mcp/model/list_objects_request.yaml
yaml/mcp/model/list_objects_response.yaml
yaml/mcp/model/get_signature_request.yaml
yaml/mcp/model/get_signature_response.yaml
yaml/mcp/model/get_source_request.yaml
yaml/mcp/model/get_source_response.yaml
yaml/mcp/model/get_references_request.yaml
yaml/mcp/model/get_references_response.yaml
yaml/mcp/model/get_reference_tree_request.yaml
yaml/mcp/model/get_reference_tree_response.yaml
yaml/mcp/model/analyze_impact_request.yaml
yaml/mcp/model/analyze_impact_response.yaml
yaml/mcp/model/inspect_request.yaml
yaml/mcp/model/inspect_response.yaml
yaml/mcp/model/list_endpoints_request.yaml
yaml/mcp/model/list_endpoints_response.yaml
```

命名規約:

| tool | request model | response model |
|---|---|---|
| `list_objects` | `list_objects_request` | `list_objects_response` |
| `get_signature` | `get_signature_request` | `get_signature_response` |
| `get_source` | `get_source_request` | `get_source_response` |
| `get_references` | `get_references_request` | `get_references_response` |
| `get_reference_tree` | `get_reference_tree_request` | `get_reference_tree_response` |
| `analyze_impact` | `analyze_impact_request` | `analyze_impact_response` |
| `inspect` | `inspect_request` | `inspect_response` |
| `list_endpoints` | `list_endpoints_request` | `list_endpoints_response` |

表現ルール:

- `selector` は `object_selector` を参照する
- `diagnostics` は `diagnostic_list` を参照する
- `references` は `reference_list` を参照する
- JSON arrayは専用list modelを作るか、粒度が細かすぎる場合は `any` + `note` にする
- enum制約は `type: str` + `note` に列挙する
- discriminated objectは `kind: str` + optional payload field + `note` で制約を記述する

完了条件:

- 8 toolすべてに request / response modelが存在する
- 各modelの `note` から元specのどの制約を表しているか追える

---

### Lane D: tool task / flow

目的:

- MCP tool呼び出しを `task` / `flow` としてblueprint化する

担当ファイル:

```text
yaml/mcp/task/list_objects.yaml
yaml/mcp/task/get_signature.yaml
yaml/mcp/task/get_source.yaml
yaml/mcp/task/get_references.yaml
yaml/mcp/task/get_reference_tree.yaml
yaml/mcp/task/analyze_impact.yaml
yaml/mcp/task/inspect.yaml
yaml/mcp/task/list_endpoints.yaml
```

各tool taskの基本形:

```yaml
nodes:
  - id: <tool_id>
    type: task
    main: true
    params:
      - name: request
        model: <tool_id>_request
    returns:
      name: response
      model: <tool_id>_response
    note: |
      MCP tool `<tool_id>` の公開contract。
      MCP toolでありHTTP endpointではないため endpoint: true は付けない。

  - id: validate_request
    type: task
    params:
      - name: request
        model: <tool_id>_request
    returns:
      name: validated_request
      model: <tool_id>_request
    note: "tool input schema / selector / enum等を検証する。"

  - id: query_service
    type: task
    params:
      - name: request
        model: <tool_id>_request
    returns:
      name: query_result
      model: any
    reads:
      - resolved_project_store
    note: |
      QueryService境界。
      Raw YAML ASTではなくResolvedProject上のsemantic objectを読む。

  - id: build_response
    type: task
    params:
      - name: query_result
        model: any
    returns:
      name: response
      model: <tool_id>_response
    note: "QueryService結果をMCP response payloadへ整形する。"

flow:
  - step: validate_request
    params:
      request: $params.request

  - step: query_service
    params:
      request: validate_request

  - step: build_response
    params:
      query_result: query_service
```

注意:

- `endpoint: true` は付けない
- QueryServiceを独立node種別として捏造しない
- `query_service` sub taskで `reads: [resolved_project_store]` を付ける
- tool固有の分岐が必要な場合は、まず `note` で表現し、必要になってから `branch` 化する

完了条件:

- 8 toolすべてにtask YAMLが存在する
- 各taskが request model / response model を参照している
- 各taskのflowがrender可能なDAGとして成立している

---

### Lane E: view / render / coverage

目的:

- Phase AのYAMLをrenderし、表現力とcoverageを確認する

担当ファイル:

```text
yaml/views/er.yaml
docs/coverage.md
renders/**              # renderer出力。手書きしない
```

主な作業:

- Phase A範囲のER viewを定義する
- `brewprint render` を実行する
- render結果を確認する
- `docs/coverage.md` を起票する
- MCP tool / model / render coverageを対応づける

実行コマンド:

```powershell
brewprint render --yaml-root docs/uc/002-brewprint-self-hosting/yaml --out docs/uc/002-brewprint-self-hosting/renders --clean
go test ./...
```

完了条件:

- renderが通る
- `renders/` がrenderer生成物として更新されている
- `docs/coverage.md` にPhase A coverageが記録されている
- render結果で表現力gapがあれば `TASKS-UC-002.md` に記録されている

---

## 3. 依存関係

```text
Lane A 方針固定
  ↓
Lane B 共通model
  ↓
Lane C tool request/response model
  ↓
Lane D tool task/flow
  ↓
Lane E render/coverage
```

並列可能な範囲:

- Lane B内の各共通modelは、名前が固定されていれば並列可能
- Lane Cはtool単位で並列可能
- Lane Dはtool単位で並列可能。ただし対応するLane C modelが先に必要
- Lane EはLane B〜Dが一通り入ってから実行する

---

## 4. Phase A tool一覧

MCP v1のPhase A対象toolは以下の8つ。

| tool | primary spec | task file | request model | response model |
|---|---|---|---|---|
| `list_objects` | `docs/spec/mcp/tools/list-objects.md` | `yaml/mcp/task/list_objects.yaml` | `list_objects_request` | `list_objects_response` |
| `get_signature` | `docs/spec/mcp/tools/get-signature.md` | `yaml/mcp/task/get_signature.yaml` | `get_signature_request` | `get_signature_response` |
| `get_source` | `docs/spec/mcp/tools/get-source.md` | `yaml/mcp/task/get_source.yaml` | `get_source_request` | `get_source_response` |
| `get_references` | `docs/spec/mcp/tools/get-references.md` | `yaml/mcp/task/get_references.yaml` | `get_references_request` | `get_references_response` |
| `get_reference_tree` | `docs/spec/mcp/tools/get-reference-tree.md` | `yaml/mcp/task/get_reference_tree.yaml` | `get_reference_tree_request` | `get_reference_tree_response` |
| `analyze_impact` | `docs/spec/mcp/tools/analyze-impact.md` | `yaml/mcp/task/analyze_impact.yaml` | `analyze_impact_request` | `analyze_impact_response` |
| `inspect` | `docs/spec/mcp/tools/inspect.md` | `yaml/mcp/task/inspect.yaml` | `inspect_request` | `inspect_response` |
| `list_endpoints` | `docs/spec/mcp/tools/list-endpoints.md` | `yaml/mcp/task/list_endpoints.yaml` | `list_endpoints_request` | `list_endpoints_response` |

---

## 5. merge前チェックリスト

各レーンの成果物を統合する前に確認する。

- [ ] `model/` 配下が原則 1ファイル = 1 main model になっている
- [ ] 参照するmodel IDが存在する
- [ ] tool taskに `endpoint: true` が付いていない
- [ ] tool taskの `params.request.model` が `<tool_id>_request` になっている
- [ ] tool taskの `returns.model` が `<tool_id>_response` になっている
- [ ] `query_service` sub taskが `resolved_project_store` を `reads` している
- [ ] optional / enum / union / discriminated object の制約が `note` に残っている
- [ ] v1表現で落ちる情報が `TASKS-UC-002.md` に記録されている
- [ ] render生成前に `yaml/views/er.yaml` がPhase A範囲を含んでいる

---

## 6. 現時点の注意点

### 6.1 v1 modelの表現力gap

MCP JSON schemaはbrewprint v1 modelより表現力が高い。
特に以下はv1 modelでは厳密表現できない。

- optional / requiredの厳密な違い
- enum値集合
- union / oneOf
- discriminated object
- 再帰的ObjectRef
- 任意JSON object

Phase Aでは `note` によるsemantic contractとして保持する。
必要に応じて、v2向け構造変更候補として `TASKS-UC-002.md` に記録する。

### 6.2 MCP toolとHTTP API endpointを混同しない

`list_endpoints` という名前のMCP toolは、HTTP endpoint一覧を返すquery toolである。
しかし `list_endpoints` 自体はHTTP endpointではない。

そのため、`yaml/mcp/task/list_endpoints.yaml` のmain taskには `endpoint: true` を付けない。

### 6.3 QueryServiceはnode化しない

v1 node種別には `service` / `component` / `package` がない。
QueryServiceを無理に `actor` や `store` として表現すると意味がずれる。

Phase Aでは `query_service` sub taskの `note` と `reads: [resolved_project_store]` で境界を表す。
将来、内部レイヤーblueprint化（Phase B）で必要になれば、v2向けnode種別として検討する。

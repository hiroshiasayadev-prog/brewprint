# UC-002 Phase A: MCP公開contract coverage

このdocは、M14 Phase Aで MCP公開contract をbrewprint YAMLへ落とす作業のcoverageを追跡する。

参照:

- `docs/tasks/m14-self-hosting.md` §Phase A
- `docs/uc/002-brewprint-self-hosting/docs/phase-a-mcp-contract.md`
- `docs/uc/002-brewprint-self-hosting/docs/phase-a-work-split.md`
- `docs/uc/002-brewprint-self-hosting/render_index.yaml`

---

## 1. 現時点の結論

Phase AのMCP公開contract blueprintは、YAML配置・validate・render・generated render review まで完了している。

- actor定義は配置済み
- 共通modelは配置済み
- ResolvedProject context storeは配置済み
- 8 MCP toolすべての request / response model は配置済み
- 8 MCP toolすべての task / flow は配置済み
- Phase A ER view YAML は配置済み
- `go test ./...` は `TASK-SELFHOST-001-01` で pass 済み
- UC-002 Phase A validate は `TASK-SELFHOST-001-01` で ok 済み
- canonical renders は11 files生成済み
- generated render review は `TASK-SELFHOST-001-02` で完了し、canonical fixture として採用可能と判定済み

---

## 2. MCP tool coverage

| tool | request model | response model | task / flow | status |
|---|---|---|---|---|
| `list_objects` | `yaml/mcp/model/list_objects_request.yaml` | `yaml/mcp/model/list_objects_response.yaml` | `yaml/mcp/task/list_objects.yaml` | 配置済み |
| `get_signature` | `yaml/mcp/model/get_signature_request.yaml` | `yaml/mcp/model/get_signature_response.yaml` | `yaml/mcp/task/get_signature.yaml` | 配置済み |
| `get_source` | `yaml/mcp/model/get_source_request.yaml` | `yaml/mcp/model/get_source_response.yaml` | `yaml/mcp/task/get_source.yaml` | 配置済み |
| `get_references` | `yaml/mcp/model/get_references_request.yaml` | `yaml/mcp/model/get_references_response.yaml` | `yaml/mcp/task/get_references.yaml` | 配置済み |
| `get_reference_tree` | `yaml/mcp/model/get_reference_tree_request.yaml` | `yaml/mcp/model/get_reference_tree_response.yaml` | `yaml/mcp/task/get_reference_tree.yaml` | 配置済み |
| `analyze_impact` | `yaml/mcp/model/analyze_impact_request.yaml` | `yaml/mcp/model/analyze_impact_response.yaml` | `yaml/mcp/task/analyze_impact.yaml` | 配置済み |
| `inspect` | `yaml/mcp/model/inspect_request.yaml` | `yaml/mcp/model/inspect_response.yaml` | `yaml/mcp/task/inspect.yaml` | 配置済み |
| `list_endpoints` | `yaml/mcp/model/list_endpoints_request.yaml` | `yaml/mcp/model/list_endpoints_response.yaml` | `yaml/mcp/task/list_endpoints.yaml` | 配置済み |

---

## 3. node / file coverage

| 対象 | 使用箇所 | status | notes |
|---|---|---|---|
| MCP client actor | `yaml/actors.yaml` | 配置済み | `mcp_client` として定義済み。 |
| MCP server actor | `yaml/actors.yaml` | 配置済み | `mcp_server` として定義済み。 |
| `object_selector` | `yaml/mcp/model/common.yaml` | 配置済み | MCP tool共通selector。optional / enum制約は `note` に保持。 |
| `source_location` | `yaml/mcp/model/source_location.yaml` | 配置済み | source file / line / column系の共通model。 |
| `object_ref` | `yaml/mcp/model/object_ref.yaml` | 配置済み | `parent` の再帰参照は `any` + `note` で暫定表現。 |
| `diagnostic` | `yaml/mcp/model/diagnostic.yaml` | 配置済み | `related` のunion listは `any` + `note` で暫定表現。 |
| `diagnostic_list` | `yaml/mcp/model/diagnostic_list.yaml` | 配置済み | `element: diagnostic` のlist model。 |
| `reference` | `yaml/mcp/model/reference.yaml` | 配置済み | semantic object間の直接参照。 |
| `reference_list` | `yaml/mcp/model/reference_list.yaml` | 配置済み | `element: reference` のlist model。 |
| `string_list` | `yaml/mcp/model/string_list.yaml` | 配置済み | string array用の共通list model。 |
| `resolved_project` | `yaml/mcp/model/resolved_project.yaml` | 配置済み | 内部index等は `any` + `note` で暫定表現。 |
| `mcp_error` | `yaml/mcp/model/mcp_error.yaml` | 配置済み | MCP-level tool error payload。 |
| ResolvedProject context store | `yaml/mcp/store/resolved_project_store.yaml` | 配置済み | `store kind: context`。各toolの `query_service` sub taskがreadsする。 |
| tool request / response model | `yaml/mcp/model/*_request.yaml`, `yaml/mcp/model/*_response.yaml` | 配置済み | 8 MCP tool分を配置済み。 |
| MCP tool task / flow | `yaml/mcp/task/*.yaml` | 配置済み | 8 MCP tool分を配置済み。HTTP endpointではないため `endpoint: true` は付けない。 |
| ER view | `yaml/views/er.yaml` | 配置済み | Phase A対象として `mcp` moduleを含める。 |

---

## 4. task / flow coverage

各MCP tool taskは以下の構造で統一している。

```text
<tool main task>
  -> validate_request
  -> query_service
  -> build_response
```

確認項目:

| check | status | notes |
|---|---|---|
| 8 toolすべてにmain taskがある | 確認済み | task idはtool idと一致。 |
| 8 toolすべてに `validate_request` がある | 確認済み | input schema / selector / enum等の検証をnoteに記録。 |
| 8 toolすべてに `query_service` がある | 確認済み | `reads: [resolved_project_store]` を持つ。 |
| 8 toolすべてに `build_response` がある | 確認済み | QueryService結果をMCP response payloadへ整形する。 |
| MCP tool taskに `endpoint: true` が付いていない | 確認済み | MCP toolはHTTP endpointではないため。 |
| QueryServiceを独立node種別として捏造していない | 確認済み | file-local sub task + noteで境界を表現。 |

---

## 5. render coverage

| render | 対象ファイル | status | notes |
|---|---|---|---|
| Project index | `renders/index.md` | 生成・review済み | 11 file canonical render set のproject indexとして生成済み。 |
| MCP group index | `renders/mcp/index.md` | 生成・review済み | `render_index.yaml` の `mcp` groupに対応し、8 DAG renders を列挙する。 |
| DAG render | `renders/mcp/dag-*.md` | 生成・review済み | 8 MCP tool taskのDAGが生成済み。各DAGは `request -> validate_request -> query_service -> build_response -> response` のPhase A flowを保持する。 |
| ER render | `renders/_cross/er.md` | 生成・review済み | 空の `erDiagram` として生成済み。Phase AはDB storeを持たないため、現行仕様どおりの非issueとして扱う。 |
| API Table render | `renders/_cross/api.md` | 対象外 | MCP toolはHTTP endpointではないため、tool taskに `endpoint: true` は付けない。 |
| State / Sequence / Wireframe | `renders/mcp/state-*.md`, `renders/mcp/seq-*.md`, `renders/mcp/wireframe-*.html` | 対象外 | Phase Aは公開I/O contract中心。FSM / scenario / wireframeは未使用。 |

---

## 6. ER view coverage

Phase AのER viewは以下を意図している。

```yaml
as: er_diagram
id: mcp_contract_er
modules:
  - module: mcp
```

ただし `docs/spec/views/er.md` の現行仕様では、ER図の描画対象は `store.kind: db` と、そこから辿れる `model.kind: struct` に限定される。
Phase Aでは `ResolvedProject` を `store kind: context` として表すため、ER図には登場しない。
MCP request / response / common schema modelもDB entityではなくcontract schemaなので、ER図ではなくcoverage表で確認する。

これは現行仕様どおりの挙動であり、現時点ではspec gapとしては扱わない。

---

## 7. v1表現力gapの確認

既知のv1表現力gapは `docs/uc/002-brewprint-self-hosting/docs/phase-a-mcp-contract.md` と `docs/uc/002-brewprint-self-hosting/docs/phase-a-work-split.md` に記録済み。

| gap | 現時点の扱い | 発見元 |
|---|---|---|
| optional / requiredの厳密な区別 | `note` に保持 | common model / tool model方針 |
| enum値集合 | `type: str` + `note` に保持 | common model / tool model方針 |
| union / oneOf | `any` + `note` に保持 | `diagnostic.related` など |
| discriminated object | `any` + `note` に保持 | `analyze_impact_request.change` など |
| 再帰的ObjectRef | `any` + `note` に保持 | `object_ref.parent` |
| 任意JSON object | `any` + `note` に保持 | tool別response payload / nested members |
| nested list object | `any` + `note` に保持 | `get_reference_tree.nodes[]`, `analyze_impact.impacts[]`, `list_endpoints.tables[]` など |

上記は `TASKS-UC-002.md` のspec gap発見ログにも記録済み。

---

## 8. render実行チェック

実行済みコマンド:

```powershell
go test ./...
go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml
go run ./cmd/brewprint render --yaml-root docs\uc\002-brewprint-self-hosting\yaml --out docs\uc\002-brewprint-self-hosting\renders --clean
go run ./cmd/brewprint render --yaml-root docs\uc\002-brewprint-self-hosting\yaml --out $env:TEMP\brewprint-uc002-render-review --clean
```

実行結果:

- `go test ./...`: pass (`TASK-SELFHOST-001-01`)
- UC-002 Phase A validate: `ok` (`TASK-SELFHOST-001-01`)
- canonical render generation: `rendered 11 file(s)` (`TASK-SELFHOST-001-01`)
- generated render review: blockingなし、canonical fixture として採用可能 (`TASK-SELFHOST-001-02`)
- temp render comparison: canonical と file list / byte-level content が一致 (`TASK-SELFHOST-001-02`)

生成済みrender files:

- `renders/index.md`
- `renders/mcp/index.md`
- `renders/mcp/dag-analyze_impact.md`
- `renders/mcp/dag-get_reference_tree.md`
- `renders/mcp/dag-get_references.md`
- `renders/mcp/dag-get_signature.md`
- `renders/mcp/dag-get_source.md`
- `renders/mcp/dag-inspect.md`
- `renders/mcp/dag-list_endpoints.md`
- `renders/mcp/dag-list_objects.md`
- `renders/_cross/er.md`

---

## 9. 次の確認ポイント

Phase A validate / render / review に残作業はない。

次の未完項目は、`WORK-SELFHOST-001` の範囲外として扱う。

- M14 Phase B internal layer blueprinting は未着手のまま。
- v1表現力gapは既存ログどおり v2向け構造変更候補として残る。
- MCP coverage実用検証レポートと editor / viewer notes のspec昇格判断は、UC-002全体作業で扱う。

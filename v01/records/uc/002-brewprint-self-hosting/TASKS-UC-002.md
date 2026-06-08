# TASKS-UC-002: brewprint Self-hosting

UC-002 固有の作業・spec gap・editor / viewer 要件の発見ログを追跡するための作業台。
プロジェクト全体の入口は `docs/TASKS.md`、milestone単位の方針は `docs/tasks/m14-self-hosting.md` に置く。

---

## ステータス

- [x] UC-002 ディレクトリ骨格を作成
  - [x] `yaml/` を作成
  - [x] `renders/` を作成
- [x] `README.md` を起票
- [x] `render_index.yaml` を起票
- [x] `TASKS-UC-002.md` を起票
- [x] `editor-viewer-notes.md` を起票
- [x] Phase A: MCP公開contract の blueprint 化を開始
- [x] Phase A の YAML が入った後に `docs/coverage.md` を起票

---

## 作業方針

M14の方針に従い、UC-002は v1.0.0-spec を基準に self-hosting を進める。

作業順序:

1. MCP公開contract を blueprint 化する
2. Phase A render を生成・確認する
3. 内部レイヤーを layer 単位で blueprint 化する
4. UC-002全体ERとMCP coverage実用検証レポートをまとめる
5. self-hosting中に発見したspec gapを本ファイルへ記録する
6. editor / viewer 要件は `editor-viewer-notes.md` に蓄積する

---

## Phase A: MCP公開contract

並列作業時は `docs/phase-a-work-split.md` を入口にし、命名規約・レーン分割・merge前チェックリストに従う。
表現方針の詳細は `docs/phase-a-mcp-contract.md` を参照する。

- [x] MCP tools の blueprint 表現方針を決める
  - `docs/phase-a-mcp-contract.md` に記録
  - MCP toolは `task`、request / responseは `model`、client / serverは `actor`、ResolvedProjectは `store kind: context` で表す
  - MCP toolはHTTP endpointではないため、tool taskに `endpoint: true` は付けない
- [x] Phase A の並列作業分割を記録する
  - `docs/phase-a-work-split.md` にレーン分割・命名規約・merge前チェックリストを記録
- [x] MCP server / client actor を定義する
  - `yaml/actors.yaml` に `mcp_client` / `mcp_server` を定義
- [x] ResolvedProject context store を定義する
  - `yaml/mcp/store/resolved_project_store.yaml` に `store kind: context` として定義
- [x] MCP tool request / response model を定義する
  - 8 MCP tool分の `*_request.yaml` / `*_response.yaml` を配置済み
- [x] MCP tool 呼び出しを task / flow として表現する
  - 8 MCP tool分のtask YAMLを配置済み
  - 各tool taskは `validate_request -> query_service -> build_response` のflowを持つ
  - 各 `query_service` sub taskは `resolved_project_store` を `reads` する
  - MCP toolはHTTP endpointではないため `endpoint: true` は付けていない
- [x] Phase A範囲のrenderを生成・確認する
  - `yaml/views/er.yaml` と `docs/coverage.md` は起票済み
  - Phase A YAMLは配置済み
  - `go test ./...` は `V01-TASK-SELFHOST-001-01` で pass 済み
  - Phase A validate は `V01-TASK-SELFHOST-001-01` で ok 済み
  - canonical renders は11 files生成済み
  - generated render review は `V01-TASK-SELFHOST-001-02` で完了し、blockingなし

---

## Phase B: 内部レイヤー

- [ ] source layer を blueprint 化する
- [ ] rawyaml layer を blueprint 化する
- [ ] semantic layer を blueprint 化する
- [ ] resolve layer を blueprint 化する
- [ ] query layer を blueprint 化する
- [ ] render layer を blueprint 化する
- [ ] CLI（`brewprint render` / `brewprint validate`）を blueprint 化する

---

## spec gap 発見ログ

### 1. MCP schemaの再帰型 / union list / 任意内部indexをv1 modelで厳密表現できない

- 対象: `docs/spec/mcp/schema.md`, `docs/spec/mcp/errors.md`
- 発見元: `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/object_ref.yaml`, `diagnostic.yaml`, `resolved_project.yaml`
- 状況:
  - `ObjectRef.parent` は本来 `ObjectRef` への再帰参照。
  - `Diagnostic.related` は `SourceLocation` または `ObjectRef` の配列。
  - `ResolvedProject` の内部registry / reverse lookup index / render mappingは任意mapや実装内部shapeを含む。
- 論点:
  - brewprint v1 modelは recursive struct、union list、arbitrary JSON object / map value unionを型レベルで表せない。
  - そのため、MCP公開contractを完全なschemaとしてblueprint化するには v2 で model 表現力の拡張候補になる。
- 暫定対応:
  - v1では該当fieldを `any` とし、`note` に元schema上の意味を残す。
- 分類: v2向け構造変更

### 2. tool別request / responseのoptional・enum・discriminated object・nested listをv1 modelで厳密表現できない

- 対象: `docs/spec/mcp/tools/*.md`
- 発見元: `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/*_request.yaml`, `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/*_response.yaml`
- 状況:
  - `direction`, `detail`, `fallback`, `severity`, `fixability` などはenum相当だが、v1 modelではenum値集合を型制約として表せない。
  - `analyze_impact.change` は `kind` ごとにpayloadが変わるdiscriminated objectだが、v1 modelではpayload相関を表せない。
  - `get_reference_tree.nodes[]`, `get_reference_tree.edges[]`, `list_endpoints.tables[]`, `analyze_impact.impacts[]` などはnested list objectで、専用modelを大量に作らない限りshapeを厳密に保持できない。
  - `get_signature.signature` と `inspect.signature` / `inspect.members` は対象kindごとのunion相当。
- 論点:
  - MCP公開contractをmachine-readableなblueprintとして使うには、optional / enum / union / discriminated object / list element schema をmodelで表す拡張が必要か。
  - あるいはPhase Aのように外部contractの出典と制約を `note` に残す運用で十分か。
- 暫定対応:
  - enum相当は `type: str` + `note` に列挙。
  - discriminated object / union / nested list object / 任意payloadは `type: any` + `note` に出典specと制約を記録。
- 分類: v2向け構造変更

発見時は以下の形式で追記する。

```markdown
### N. タイトル

- 対象: `docs/spec/...`
- 発見元: `docs/uc/002-brewprint-self-hosting/yaml/...`
- 状況:
- 論点:
- 暫定対応:
- 分類: v1範囲内修正 / v2向け構造変更 / 棄却 / 判断待ち
```

---

## coverage.md 起票方針

`docs/coverage.md` は、Phase A の YAML が入った後に起票する。
空の骨格段階ではなく、MCP公開contractの実例YAMLに対して、以下の観点を対応づけて記録する。

- MCP tool coverage
- node / field coverage
- render coverage
- spec gap の発見元

---

## MCP coverage / tool不足ログ

現時点では未記録。

MCP toolsだけでは探索・確認が難しかった点があれば、ここに記録する。

---

## 完了条件

- [ ] MCP公開contract が blueprint 化されている
- [ ] 内部レイヤーが layer 単位で blueprint 化されている
- [ ] Phase A / Phase B のrender結果を確認している
- [ ] `docs/coverage.md` が実例YAMLに追随している
- [ ] spec gap 発見ログをレビューし、v1範囲内修正 / v2向け構造変更 / 棄却に分類している
- [ ] MCP coverage 実用検証レポートをまとめている
- [ ] editor / viewer notes のspec昇格要否を判断している

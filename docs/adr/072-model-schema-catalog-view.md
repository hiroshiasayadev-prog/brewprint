# 072: model / schema catalog view

- **status**: accepted
- **date**: 2026-05-11

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-069 では anonymous inline struct TypeRef を導入しないことを決めた。
ADR-070 では、その補完策として public model / file-private helper model の visibility 概念を導入した。
ADR-071 では、task file 内 file-private helper model を DAG Markdown detail section に表示する方針を決めた。

これにより、file-local な helper model は task DAG Markdown 上で確認できるようになる。
一方で、module / contract 単位で model 群を俯瞰する view はまだ存在しない。

現行 render には、ER 図と API Table がある。
しかし ER 図は `store.kind: db` から辿れる `model.kind: struct` を対象とする view であり、以下を十分に俯瞰できない。

- store に紐づかない public contract model
- task params / returns の request / response model
- enum / list / dict model
- JSON 埋め込み用 model
- file-private helper model
- MCP schema / QueryService contract のような schema 群

API Table は endpoint task の一覧であり、model / schema の一覧ではない。

UC-002 self-hosting では、MCP tool の request / response / common schema / diagnostic / ObjectRef / Reference など、多数の public contract model が存在する。
ADR-070 によって helper model を導入すると、public model と file-private helper model の関係も増える。
これらを YAML や MCP query だけで追うのでは、人間向け render として弱い。

本ADRは、public model と file-private helper model を module / contract 単位で俯瞰する curated Markdown view として、model / schema catalog view を導入する。

ただし、本ADRは opt-in の cross-cutting catalog view を定めるものであり、ADR-070 §10 の不可視禁止制約を単体で履行するものではない。
model file 内 helper model の file-local な自動 render exposure は、後続 ADR-075 で扱う。

## 決定

### 1. model / schema catalog view を導入する

brewprint に、model / schema catalog view を導入する。

この view は、対象 module 群に含まれる model node を収集し、public model と、必要に応じて file-private helper model を Markdown table として一覧する。

model catalog は user が `as: model_catalog` view YAML を定義した場合に生成される opt-in の curated view である。model file 内 helper model を常に自動表示する file-local render ではない。

目的は以下である。

- public contract model の一覧を人間が確認できるようにする
- file-private helper model がどの public model / field / task から使われるかを確認できるようにする
- ER 図に出ない schema model を俯瞰できるようにする
- anonymous inline struct 不採用により増える named model 群を見失わないようにする

### 2. view YAML の `as:` は `model_catalog` とする

view YAML の file type は `as: model_catalog` とする。

```yaml
as: model_catalog
id: mcp_contract_models
note: MCP公開contract model一覧
modules:
  - module: mcp
    include_submodules: true
include:
  public_models: true
  file_private_models: true
  enums: true
  list_dict_models: true
  tagged_union_models: true
```

`model_catalog` は、人間向け Markdown render を主目的とする view である。
MCP query の `list_objects` とは異なり、module / contract 単位で model 群を読ませるための curated view として扱う。

ユースケース別の分割は `include` flag では扱わない。
ユースケース単位で catalog を分けたい場合は、UC ディレクトリ、view YAML、または `modules` scope を分ける。

### 3. view schema を定義する

`model_catalog` view は以下のフィールドを持つ。

| field | required | type | meaning |
|---|---:|---|---|
| `as` | yes | string | `model_catalog` 固定 |
| `id` | yes | string | view id。Markdown H1 と output filename に使う |
| `note` | no | string | view 全体の説明 |
| `modules` | yes | list<module-entry> | 収集対象 module |
| `include` | no | include options | 出力対象 model 種別の filter |

`modules[]` の schema は API Table と同じ方針を流用する。

| field | required | type | meaning |
|---|---:|---|---|
| `module` | yes | string | 対象 module path |
| `include_submodules` | no | bool | true の場合、配下 submodule も収集する。省略時 false |

`include` は省略可能であり、省略時は以下とする。

```yaml
include:
  public_models: true
  file_private_models: false
  enums: true
  list_dict_models: true
  tagged_union_models: true
```

`include` は visibility 軸と kind 軸を組み合わせた filter として扱う。
起票時点では、各 flag の意味を以下とする。

| flag | 対象 |
|---|---|
| `public_models` | visibility = public かつ kind = struct の model |
| `enums` | visibility = public かつ kind = enum の model |
| `list_dict_models` | visibility = public かつ kind = list / dict の model |
| `tagged_union_models` | visibility = public かつ kind = tagged_union の model |
| `file_private_models` | visibility = file-private の helper model 全 kind |

visibility = file-private の enum / list / dict / tagged_union は、`file_private_models: true` のとき表示対象に含まれる。
`file_private_models: false` の場合、private helper model は kind に関係なく catalog には出さない。

file-private helper model は default では非表示とする。
本ADRの model catalog は、ADR-070 §10 を単体で履行するものではなく、user が module / contract 単位で schema 群を読むための opt-in view である。
file-private helper model の自動表示は ADR-071（task file）および後続 ADR-075（model file）で担う。
そのため、catalog の default は public surface 中心とし、helper model を catalog 経由で確認したい場合は `file_private_models: true` を明示する。

### 4. 出力 Markdown は public models と private helper models を分ける

model catalog Markdown は、public model と file-private helper model を分けて表示する。

推奨出力構成:

```markdown
# mcp_contract_models

MCP公開contract model一覧

## Public models

| model | kind | module | file | shape | used by | note |
|---|---|---|---|---|---|---|
| get_reference_tree_response | struct | mcp | mcp/model/get_reference_tree_response.yaml | root: object_ref<br/>nodes: list<reference_tree_node> | get_reference_tree.returns | response model |

## Enums

| enum | module | values | used by | note |
|---|---|---|---|---|
| mcp_diagnostic_severity | mcp | error<br/>warning<br/>info<br/>hint | diagnostic.severity | diagnostic severity |

## List / dict models

| model | kind | module | shape | key semantics | used by | note |
|---|---|---|---|---|---|---|
| diagnostic_list | list | mcp | list<diagnostic> | — | response.diagnostics | Diagnostic list |

## Tagged unions

| model | module | discriminator | variants | used by | note |
|---|---|---|---|---|---|
| analyze_impact_change | mcp | kind | rename<br/>remove<br/>change_type | analyze_impact_request.change | change payload |

## Private helper models

| file | model | kind | used by | shape | note |
|---|---|---|---|---|---|
| mcp/model/get_reference_tree_response.yaml | reference_tree_node | struct | get_reference_tree_response.nodes | object: object_ref<br/>depth: int | traversal node entry |
```

実際の section 分割、列、sort order の詳細は spec 反映時に `docs/spec/views/model-catalog.md` で確定する。
ただし、public model と file-private helper model は同じ table に混ぜず、visibility が明確に分かる出力にする。

### 5. model catalog は ER 図の代替ではない

model catalog は schema 一覧であり、ER 図の代替ではない。

ER 図は DB entity / FK relation を描く view である。
model catalog は、DB に紐づかない contract model、request / response model、enum、list / dict、tagged union、helper model を含めて一覧する view である。

したがって、ER 図に出る model も model catalog に出てよいが、表示目的は異なる。

### 6. model catalog は MCP list_objects の代替ではない

MCP `list_objects` は LLM / tool が project 内 object を探索する query API である。
model catalog は人間向け Markdown render であり、module / contract 単位の設計レビューを助ける view である。

両者は対象が重なるが、責務は異なる。

- `list_objects`: query layer。機械が selector を得るための object listing
- `model_catalog`: render layer。人間が schema 群を読むための curated Markdown

### 7. file-private helper model の used-by を表示する

file-private helper model は外部 QualifiedID を持たないため、catalog 上では所属 file と local id を組み合わせて表示する。

used-by には、同一 file 内で当該 helper model を参照している field / param / returns / element / value を表示する。
表記形式は ADR-071 の `## Private models` と揃える。

深い transitive usage の完全展開は本ADRの主対象外とする。
初期導入では depth 1 の直接参照を基本とし、必要に応じて spec 側で拡張する。

### 8. render output placement は view spec 反映時に確定する

model catalog は cross-cutting view であり、複数 module を対象にできる。

render output の配置は、既存の render placement 方針に従い spec 反映時に確定する。
候補としては以下がある。

- `_cross/model-{id}.md`
- `_cross/schema-{id}.md`
- group 配下の `model-{id}.md`

本ADRでは、model catalog が Markdown render を生成する view であることを決めるに留め、最終 output path は `docs/spec/project-layout.md` / `docs/spec/views/model-catalog.md` 反映時に確定する。

## 理由

### なぜ model catalog が必要か

ADR-069 で anonymous inline struct を導入しないと決めたため、意味を持つ shape は named model として定義する方針になる。
ADR-070 で file-private helper model を認めたことで、public model と private helper model の関係も増える。

これらを YAML だけで追うと、人間が schema 群を俯瞰しづらい。
MCP query で問い合わせられるとしても、人間向け render として一覧が必要である。

model catalog は、named model を増やしても設計全体を見失わないための view である。

### なぜ `model_catalog` という名前か

候補として `schema_catalog` もある。
しかし brewprint の YAML 上の型定義単位は `type: model` であり、struct / enum / list / dict / helper model をまとめて扱う中心概念は model である。

そのため、view file type は `model_catalog` とする。
ただし、説明上は MCP schema や public contract schema の catalog として使ってよい。

### なぜ public / private を分けて表示するか

public model と file-private helper model は外部 surface が異なる。
これらを同じ table に混ぜると、外部から参照してよい schema と file-local helper shape の区別が弱くなる。

ADR-070 の visibility 方針に従い、render 上も public / private を明確に分ける。

### なぜ default で private helper model を非表示にするか

model catalog は public contract model の一覧として使われる場面が多い。
そのため、default では public model を中心に出し、file-private helper model は `include.file_private_models: true` で明示した場合に出す。

本ADRの catalog は opt-in の curated view であり、ADR-070 §10 の不可視禁止制約を単体で履行するものではない。
file-private helper model の自動表示は、task file については ADR-071、model file については後続 ADR-075 が担う。
この分担により、catalog は public surface を中心にした一覧として使いやすく保ちつつ、必要な場合には helper model も含められる。

### なぜ render output placement を本ADRで決めきらないか

output placement は `render_index.yaml`、group / `_cross` 配置、nested module の扱いに関わる。
これは view 導入時の spec 反映で既存 render placement 方針と整合させる必要がある。

本ADRでは view の必要性・責務・基本 schema を決め、output path の詳細は spec に委ねる。

## 却下した代替案

### 代替案A: ER 図を拡張して model catalog を兼ねる

- 利点: 既存 view を増やさずに済む
- 欠点: ER 図は DB entity / relation の view であり、request / response model、enum、list / dict、helper model を扱うには責務が違う

→ 却下。model catalog は別 view とする。

### 代替案B: API Table を拡張して params / returns model を深く表示する

- 利点: endpoint contract と model を同時に見られる
- 欠点: API Table は endpoint task の一覧であり、model / schema 群そのものの catalog ではない。MCP tool のような non-endpoint task や共通 model を扱いづらい

→ 却下。API Table とは別 view とする。

### 代替案C: MCP list_objects / inspect だけで対応する

- 利点: render view を増やさずに済む
- 欠点: 人間が Markdown render だけで schema 群を俯瞰できない。brewprint の人間向け render 方針と合わない

→ 却下。MCP query は補助であり、人間向け catalog render は別に持つ。

### 代替案D: file-private helper model を常に catalog に表示する

- 利点: helper model が見落とされない
- 欠点: public contract model の一覧として使いたい場合に noise が増える。public surface と internal helper shape の区別が弱くなる

→ 却下。default では public model 中心とし、helper model は include option で明示する。

### 代替案E: view 名を `schema_catalog` にする

- 利点: MCP schema / contract schema を扱う用途が伝わりやすい
- 欠点: brewprint の型定義単位は `model` であり、既存用語とずれる

→ 却下。file type は `model_catalog` とする。説明上は schema catalog と呼んでもよい。

## 影響

### spec への影響

本ADR受理後、以下の spec 更新が必要になる。

- `docs/spec/views/model-catalog.md`
  - `as: model_catalog` の schema を定義する
  - modules / include の schema を定義する
  - `tagged_union_models` include flag を定義する
  - Markdown output の section / table / sort order / truncation を定義する
  - public tagged union model を `## Tagged unions` section に出すことを定義する
  - public model と file-private helper model の出力ルールを定義する

- `docs/spec/file-types.md`
  - view file type として `model_catalog` を追加する

- `docs/spec/project-layout.md`
  - model catalog render の output placement を追加する

- `docs/spec/nodes.md` / `docs/spec/type-ref.md`
  - model catalog が参照する public / file-private model visibility と TypeRef 解決規則へのリンクを追加する

### render 実装への影響

renderer は `as: model_catalog` view を読み、対象 modules から model node を収集して Markdown を生成する責務を持つ。

renderer は `include.tagged_union_models` が true の場合、public tagged union model を `## Tagged unions` section に出力する。
file-private tagged union model は、他の private helper model と同じく `include.file_private_models: true` のとき `## Private helper models` に出力する。

初期実装の優先範囲、fixture 追加、golden 更新、truncation 閾値の詳細は task file / UC task file で追跡する。

### UC-002 への影響

UC-002 self-hosting では、MCP contract model 群を俯瞰する model catalog view を追加できる。

特に以下を確認しやすくなる。

- MCP request / response model
- ObjectRef / Reference / Diagnostic などの共通 schema
- enum model
- tagged union model
- response model 内部の file-private helper model
- `any + note` から named helper model / tagged union model へ移行した箇所

具体的な view YAML 追加、fixture migration、golden 更新は task file / UC task file で追跡する。

### ADR-070 / ADR-071 / ADR-075 との関係

ADR-070 の「file-private helper model を人間向け render から不可視にしてはならない」制約は、複数の render ADR に分担して履行する。

| 対象 | 履行責任 | render 経路 |
|---|---|---|
| task file 内 helper model | ADR-071 | DAG Markdown `## Private models` 自動表示 |
| model file 内 helper model | ADR-075（後続） | model file 単体 render の自動表示 |
| module / contract 単位の curated 俯瞰 | 本ADR | `as: model_catalog` opt-in view |

本ADRは、public model と file-private helper model を module / contract 単位で俯瞰する curated catalog view を導入する。
model file 内 helper model の file-local な自動 render exposure は、本ADRではなく後続 ADR-075 で扱う。

## Evidence

- commit: 5ae7769
- impl commit: tbd
- close boundary: M15 / `v1.1.0-spec` では follow-up scope として deferred。実装は含めない。
- 参考: ADR-069 anonymous inline struct 不採用、ADR-070 model visibility、ADR-071 task file helper model render exposure、UC-002 MCP contract model 群

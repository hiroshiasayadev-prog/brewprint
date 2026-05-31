# 075: model file render

- **status**: proposed
- **date**: 2026-05-11
- **depends_on**: ADR-070, ADR-071, ADR-072, ADR-073

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-070 では、model に public / file-private の visibility 概念を導入し、file-private helper model を認めた。
同ADRでは、file-private helper model を人間向け render から不可視にしてはならない、という制約も残した。

ADR-071 では、task file 内 file-private helper model を DAG Markdown の `## Private models` section に自動表示する方針を定めた。
これにより、task file 内 helper model は 1 task YAML に対応する DAG Markdown render で確認できる。

一方で、model file 内 file-private helper model については、まだ file-local な自動 render exposure が定義されていない。
ADR-072 では `as: model_catalog` view を導入したが、これは user が view YAML を書いた場合に生成される opt-in の curated catalog view であり、ADR-070 の不可視禁止制約を単体で履行するものではない。

現行 render 体系では、task file には `dag-{task-id}.md` という単体 render がある。
API Table は endpoint task を module / contract 単位で俯瞰する集約 view であり、task file 単体 render の代替ではない。

model についても同じ構造に揃えるのが自然である。

```text
task/*.yaml
  単体 render: dag-{task-id}.md
  集約 view:   api_table

model/*.yaml
  単体 render: model-{model-id}.md
  集約 view:   model_catalog
```

ER 図は `store.kind: db` から辿れる DB schema / relation を描く view であり、model file 単体の schema render ではない。
store に紐づかない request / response model、enum / list / dict model、file-private helper model は ER 図では十分に確認できない。

本ADRは、model file に対する file-local な自動 Markdown render として model file render を導入する。

## 決定

### 1. model file render を導入する

brewprint は `model/*.yaml` に対して、1 YAML file = 1 Markdown render を生成する。

model file render は、model file の public main model と、その file 内の file-private helper model を人間向けに表示する file-local render である。

model file render は opt-in view YAML ではなく、model file から自動生成される render とする。
これにより、model file 内 file-private helper model は、user が `model_catalog` view YAML を書かなくても人間向け render から確認できる。

### 2. model file render は model file 内の public main model と helper model を表示する

model file render の対象は以下である。

| 対象 | 表示 |
|---|---|
| public main model | `## Public model` section に表示 |
| file-private helper model | `## Private models` section に表示 |

public main model は ADR-070 に従い `main: true` を持つ model である。
file-private helper model は同一 file 内の `type: model` かつ `main: true` を持たない model である。

### 3. task file 内 helper model は model file render の対象外とする

task file 内 file-private helper model は、ADR-071 に従い DAG Markdown の `## Private models` section に表示する。

model file render は `model/*.yaml` の render であり、task file 内 helper model は対象外とする。

### 4. model file render は Markdown table を基本とする

model file render は Mermaid 図ではなく Markdown table を基本とする。

model は Data layer の schema 定義であり、processing flow の node / edge を持たない。
そのため、Mermaid DAG や ER 図として描くのではなく、fields / values / element / value などを table で表示する。

### 5. 出力フォーマットを定義する

model file render の基本フォーマットは以下とする。

```markdown
# {main model id}

{main model note}

## Public model

| property | value |
|---|---|
| kind | struct |
| visibility | public |
| file | mcp/model/get_reference_tree_response.yaml |

### Fields

| field | type | note |
|---|---|---|
| root | object_ref | traversal root |
| nodes | list<reference_tree_node> | reached node entries |

## Private models

| model | kind | used by | shape | note |
|---|---|---|---|---|
| reference_tree_node | struct | get_reference_tree_response.nodes | object: object_ref<br/>depth: int | traversal node entry |
```

`## Private models` は、対象 model file に file-private helper model が1個以上存在する場合のみ出力する。
存在しない場合は section 自体を省略する。

### 6. public model の kind 別表示を定義する

`## Public model` section は、main model の `kind` に応じて表示を変える。

| kind | 表示 |
|---|---|
| `struct` | `### Fields` table を出す |
| `enum` | `### Values` table または values list を出す |
| `list` | `### Element` として element TypeRef を出す |
| `dict` | `### Value` として value TypeRef を出す。key は常に `str` とし、key semantics は note / model名 / field名から確認する |
| `tagged_union` | `### Discriminator` と `### Variants` table を出す |

struct field の table は以下を基本とする。

| column | 内容 |
|---|---|
| `field` | field name |
| `type` | TypeRef |
| `note` | field note。ない場合は `—` |

PK / FK / unique などの field metadata をどう表示するかは spec 反映時に `docs/spec/views/model.md` または `docs/spec/views/model-file.md` で確定する。

`tagged_union` の public model は、discriminator field と variant payload を分けて表示する。

`### Discriminator` では、variant 判定に使う discriminator field name を表示する。

```markdown
### Discriminator

| property | value |
|---|---|
| field | kind |
```

`### Variants` では、各 `variants[].tag` と、その variant が持つ payload fields を表示する。
payload field は `name: type` を基本とし、field note がある場合は同じ cell 内に補足として表示する。
payload field がない variant、つまり `fields: []` は `—` と表示する。

variant の表示順は YAML の `variants[]` 順を保持する。
variant payload fields の表示順も YAML の `fields[]` 順を保持する。

```markdown
### Variants

| tag | payload fields |
|---|---|
| rename | `new_id: str` — new object id |
| remove | — |
| change_type | `new_type: str` |
```

複数 field を持つ variant は、`payload fields` cell 内で `<br/>` 区切りで表示する。

```markdown
| tag | payload fields |
|---|---|
| change_contract | `old_contract: str`<br/>`new_contract: str` |
```

### 7. private helper model 表示は ADR-071 の `## Private models` と揃える

model file render の `## Private models` table は、ADR-071 の DAG Markdown `## Private models` table と同じ基本列を使う。

| column | 内容 |
|---|---|
| `model` | file-private helper model の local id |
| `kind` | `struct` / `list` / `dict` / `enum` / `tagged_union` |
| `used by` | 同一 file 内で当該 helper model を直接参照している箇所の一覧 |
| `shape` | kind に応じた概要 |
| `note` | helper model の note。ない場合は `—` |

`used by` と `shape` の表記ルールは ADR-071 と揃える。

- `used by` は `<parent_id>.<location>` 形式
- 複数参照は `<br/>` 区切り
- depth 1 の直接参照だけを列挙
- nested helper model は型名参照のみ
- 深い展開は model catalog / schema catalog view に委ねる

`tagged_union` helper model の `shape` は、`discriminator: <field>` と、各 variant を `<tag>: <payload fields>` 形式で `<br/>` 区切りで表示する。
payload field がない variant は `<tag>: —` と表示する。

例:

```markdown
| model | kind | used by | shape | note |
|---|---|---|---|---|
| analyze_impact_change | tagged_union | analyze_impact_request.change | discriminator: kind<br/>rename: `new_id: str`<br/>remove: —<br/>change_type: `new_type: str` | change payload |
```

### 8. output path は group 配下の `model-{model-id}.md` とする

model file render は group 配下に出力する。

```text
renders/{group-id}/model-{model-id}.md
```

例:

```text
yaml/mcp/model/get_reference_tree_response.yaml
→ renders/mcp/model-get_reference_tree_response.md
```

`model-id` は public main model の local id を使う。
同一 group 内で output path が衝突する場合は、既存 render output collision 方針と同じく validation error とする。silent overwrite は禁止する。

model file render の所属 group は、public main model の module path を `render_index.yaml` の group module 定義に照合して決定する。
この規則は task DAG など既存 group render の module-based grouping と揃える。

初期方針は以下とする。

| ケース | 扱い |
|---|---|
| module path が1つの group に一致する | その group 配下に出力する |
| nested module が group の `modules` entry 配下にある | 既存の `include_submodules` / module matching rule に従う |
| module path がどの group にも一致しない | project-layout の default / uncovered module grouping に従う |
| module path が複数 group に一致する | render output placement が曖昧なため validation error とする |

uncovered module、nested module、複数 group 一致時の詳細な matching rule は `docs/spec/project-layout.md` への spec 反映時に確定する。

### 9. group index / master index に model render を含める

model file render は正式な group render の一種として扱う。

`renders/{group-id}/index.md` の group index には、kind = `Model` の行として model render を含める。

master `renders/index.md` には、既存の DAG / State / Sequence / Wireframe の count と同様に、Model count を表示する列を追加する。

最終的な index table の列順は spec 反映時に `docs/spec/project-layout.md` で確定する。
候補としては以下を基本とする。

```markdown
| group | DAG | Model | State | Sequence | Wireframe | ER | API |
```

### 10. model_catalog との責務を分ける

model file render と model catalog view は責務が異なる。

| render | 性質 | 目的 |
|---|---|---|
| model file render | 自動 / file-local / 1 YAML = 1 Markdown | 1つの model file の public main model と helper model を読む |
| model_catalog | opt-in / module・contract 単位 / view YAML | 複数 module の model 群を curated に俯瞰する |

model file render は ADR-070 §10 の不可視禁止制約を file-local に履行する。
model_catalog は module / contract 単位で schema 群をレビューするための集約 view である。

### 11. ER 図との責務を分ける

model file render は ER 図の代替ではない。

ER 図は `store.kind: db` と `store.of` を起点に DB entity / FK relation を描く view である。
model file render は DB に出るかどうかに関係なく、model file の schema 定義を表示する view である。

同じ model が ER 図と model file render の両方に現れてよい。
表示目的が異なるためである。

## 理由

### なぜ model file render が必要か

ADR-070 により、model file は 1 public main model + 0個以上の file-private helper model を持てるようになる。

helper model が YAML 上に存在するにもかかわらず、人間向け render に出ない場合、人間は YAML 直接確認か MCP inspect を使わなければ helper schema を確認できない。
これは ADR-070 の不可視禁止制約と brewprint の人間向け render 方針に反する。

model file render により、model file 内 helper model は自動 render から確認できる。

### なぜ 1 YAML = 1 render か

brewprint の task file には DAG Markdown render があり、1 task YAML の内容を人間が読むための file-local render を提供している。

model file も同様に、1 model YAML の内容を人間が読むための file-local render を持つべきである。

この方針により、task / model ともに以下の二層構造が揃う。

```text
file-local render:
  1 YAML = 1 Markdown

aggregate view:
  view YAML で scope を指定して横断表示
```

### なぜ Markdown table か

model は schema 定義であり、flow node / edge を持たない。
そのため、Mermaid diagram よりも Markdown table の方が情報を正確に表せる。

struct fields、enum values、list element、dict value、helper model の used-by は table 表示に向いている。

### なぜ group 配下に出すか

model file render は file-local render であり、task DAG / state / sequence / wireframe と同じく group に属する render である。

ER / API Table のような cross-cutting view ではないため、`_cross/` ではなく group 配下に配置する。

### なぜ model_catalog と分けるか

model_catalog は module / contract 単位の curated catalog であり、user が view YAML で scope を指定して生成する。

一方、model file render は model file に対して自動生成される file-local render である。

この二者を分けることで、以下を両立できる。

- どの model file も自動 render から読める
- 必要なときは model_catalog で module / contract 単位に俯瞰できる

## 却下した代替案

### 代替案A: model file 内 helper model は model_catalog だけで表示する

- 利点: render 種別を増やさずに済む
- 欠点: model_catalog は opt-in view であり、view YAML がない場合 helper model が不可視になる。ADR-070 §10 を単体では満たせない

→ 却下。file-local な自動 render として model file render を導入する。

### 代替案B: ER 図を model file render の代替にする

- 利点: 既存 ER view を使える
- 欠点: ER 図は DB schema / relation の view であり、store.kind: db に辿り着かない model、enum、list/dict、helper model を扱うには責務が違う

→ 却下。ER 図とは別に model file render を導入する。

### 代替案C: MCP inspect だけで対応する

- 利点: render 実装を増やさずに済む
- 欠点: 人間が Markdown render だけで model file の schema を確認できない。brewprint の人間向け render 方針と合わない

→ 却下。MCP inspect は query layer であり、人間向け render の代替にはしない。

### 代替案D: model file render を `_cross/` に出す

- 利点: model 関連 render を一箇所に集められる
- 欠点: model file render は file-local render であり、cross-cutting view ではない。task DAG と同じく group に属する方が自然

→ 却下。model file render は group 配下に出す。

### 代替案E: public model だけを render し、helper model は省略する

- 利点: render が短くなる
- 欠点: ADR-070 の不可視禁止制約を満たせない。helper model の存在を確認するには YAML / MCP が必要になる

→ 却下。helper model が存在する場合は `## Private models` section に表示する。

## 影響

### spec への影響

本ADR受理後、以下の spec 更新が必要になる。

- `docs/spec/views/model-file.md` または `docs/spec/views/model.md`
  - model file render の出力フォーマットを定義する
  - public model の kind 別表示を定義する
  - `tagged_union` の `### Discriminator` / `### Variants` 表示を定義する
  - `## Private models` table を定義する
  - output path を定義する

- `docs/spec/project-layout.md`
  - group 配下 render として `model-{model-id}.md` を追加する
  - model file render の所属 group を public main model の module path から解決する規則を追加する
  - uncovered module / nested module / 複数 group 一致時の扱いを定義する
  - group index / master index に Model を追加する
  - output collision rule を model render に適用する

- `docs/spec/nodes.md`
  - model file が render 対象になること
  - model file 内 helper model が `## Private models` に表示されること

### render 実装への影響

renderer は `model/*.yaml` の public main model ごとに model file render Markdown を生成する責務を持つ。

model file renderer は `tagged_union` の `discriminator` / `variants` を読み、public model section と private helper model table の両方で表示できる必要がある。

renderer は public main model の module path から所属 group を解決し、複数 group に一致する場合は validation error として扱う。

具体的な実装順序、fixture migration、golden 更新は task file / UC task file で追跡する。

### UC / fixture への影響

既存 UC / fixture の model YAML に対して、model file render の golden が追加される。

ADR-070 により既存 model YAML には `main: true` migration が必要になるため、model file render 導入時にはこの migration と合わせて fixture 更新を行う。

UC-002 では、MCP contract model file の public model と file-private helper model を model file render で確認できるようになる。

### ADR-070 / ADR-071 / ADR-072 との関係

ADR-070 の不可視禁止制約は、task file 側を ADR-071、model file 側を本ADRで履行する。

ADR-071 は task file 内 helper model を DAG Markdown `## Private models` に出す。
本ADRは model file 内 helper model を model file render `## Private models` に出す。

ADR-072 の model_catalog は、module / contract 単位で model 群を俯瞰する opt-in view であり、file-local render である本ADRとは責務が異なる。

### 後続ADR候補

本ADRとは別に、DAG asset node label に TypeRef を表示する ADR-074 がある。

ADR-074 は DAG readability 改善であり、model file render とは別論点である。

## Evidence

- commit: tbd
- impl commit: tbd
- 参考: ADR-070 model visibility、ADR-071 task file helper model render exposure、ADR-072 model catalog view、docs/spec/project-layout.md の group render 方針

# V01-ADR-070: model visibility と file-private helper model

- **status**: accepted
- **date**: 2026-05-11

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

V01-ADR-069 では、TypeRef に anonymous inline struct を導入しないことを決定した。

brewprint は実装言語ではなく、人間と LLM が設計意図を共有するための設計層である。
そのため、設計上意味を持つ shape には名前を付ける方針を採る。

一方で、anonymous inline struct を使わずにすべての helper shape を public model として定義すると、以下の問題が起きる。

- 小さな one-off helper shape まで public model file として増える
- public contract と内部補助 shape の区別がつきにくくなる
- MCP / render / LLM から見た外部参照可能 surface が広がりすぎる
- model file が増えたとき、どれが外部から参照してよい schema なのか判断しづらくなる

V01-ADR-011 では task file について「1ファイル = 1メインノード + 複数サブノード」を認め、サブノードを file-private として扱う方針を導入した。
V01-ADR-058 では、この file-private 性を実装上も保証するため、private sub node は外部 QualifiedID 参照対象ではなく、file-local object として扱うことを補強した。

しかし、同じ file-private helper の考え方を model に適用するかは未定義である。
現行 `docs/spec/nodes.md` では model は Data レイヤーの型定義として `model/` サブディレクトリに 1ファイル=1定義で置く前提のままであり、model-local helper model や task-local helper model は定義されていない。

UC-002 self-hosting の MCP contract YAML では、response model 内部の entry shape や集約 object を `any + note` で暫定表現している箇所が複数ある。
例として、`get_reference_tree_response.nodes` / `edges`、`list_endpoints_response.tables`、`analyze_impact_response.impacts` / `summary` / `coverage` などは、本来 named helper model に切り出したい shape である。
これらは public contract の代表 model そのものではなく、特定 response model の内部 shape として閉じるのが自然なものを含む。

本ADRは、anonymous inline struct を採用しない方針を実用上成立させるため、model に public / file-private の visibility 概念を導入し、file-private helper model を定義する。

### 判断根拠としての UC-002 実例の扱い

本ADRで参照する UC-002 の YAML は、ADR起票時点の self-hosting 実例である。
これらの YAML は今後の spec 反映・implementation・UC-002 migration によって更新されうるため、本文中の具体的な field 名や `any + note` の配置は恒久仕様ではない。

本ADRが判断根拠として扱うのは、具体的な YAML shape そのものではなく、そこから抽出された以下の設計上の性質である。

- anonymous inline struct を使わずに named model へ切り出したい内部 shape が存在する
- その一部は複数 tool で再利用する public model ではなく、特定 file / parent model に閉じた helper shape である
- すべてを public model にすると、public surface と model file 数が過剰に増える
- `any + note` のままでは、machine-readable schema として LLM / validator / MCP schema generation が扱いづらい

具体的な fixture migration の範囲や順序は task file / UC task file で管理する。

## 決定

### 1. model に public / file-private の visibility 概念を導入する

brewprint の model には、仕様上の visibility として以下の2種類を導入する。

| visibility | 意味 |
|---|---|
| `public` | 外部 file / module から参照可能な model。public contract、store.of、event.payload、cross-file TypeRef から参照できる |
| `file-private` | 同一 YAML file 内だけで参照可能な helper model。外部 file / module から QualifiedID 参照できない |

### 2. model でも `main: true` を認める

V01-ADR-011 では `main: true` を task にのみ適用するとしていたが、本ADRでは model にも `main: true` を拡張する。

model file では、`main: true` が付いた model をその file の代表 model とする。
単一 model file であっても、public model として扱うには `main: true` を明示しなければならない。

```yaml
nodes:
  - id: get_reference_tree_response
    type: model
    main: true
    kind: struct
    fields:
      - name: nodes
        type: list<reference_tree_node>

  - id: reference_tree_node
    type: model
    kind: struct
    fields:
      - name: object
        type: object_ref
      - name: depth
        type: int
```

この例では、`get_reference_tree_response` が public model、`reference_tree_node` が file-private helper model である。

### 3. `main: true` model を public model とする

model node の visibility は以下の規則で決まる。

| model node | visibility |
|---|---|
| `main: true` を持つ model | `public` |
| `main: true` を持たない model | `file-private` |

YAML field として `visibility: public` / `visibility: file-private` は導入しない。
既存 task subnode と同じく、main node と非 main node の区別で外部 surface を決める。

単一 model のみを持つ model file でも、public model として扱うには `main: true` を明示しなければならない。

### 4. task file / model file の両方で file-private helper model を認める

file-private helper model は、以下の両方で定義できる。

| file種別 | 用途 |
|---|---|
| task file | task-local response / intermediate shape、特定 task の内部 flow でのみ使う補助 schema |
| model file | public model の内部 entry / nested object / helper enum / helper dict など、特定 model file に閉じた補助 schema |

ただし、helper model はその file 内でのみ TypeRef 参照できる。
再利用が始まった helper model は public model へ昇格させる。

### 5. model file は 1 public main model + 0個以上の file-private helper model を持てる

`model/*.yaml` では、1つの public main model を `main: true` で明示する。
その他の model は file-private helper model とする。

model file の構成制約は以下の通り。

- 1 file = 1 public main model + 0個以上の file-private helper model
- public main model は `main: true` で必ず明示する
- public model を複数持つ model file は不正とする
- `main: true` model が存在しない model file は不正とする
- `main: true` を持たない model は file-private helper model として扱う

この規則は、task file の main node / private subnode 方針と揃えるためのものである。

### 6. file-private helper model は同一 file 内からのみ TypeRef 参照できる

file-private helper model は、同一 YAML file 内の TypeRef からのみ参照できる。

外部 file / module から helper model を参照することはできない。
QualifiedID で file-private helper model を参照する構文も導入しない。

```yaml
# OK: 同一file内の helper model 参照
fields:
  - name: nodes
    type: list<reference_tree_node>
```

```yaml
# NG: 別fileから helper model を参照
fields:
  - name: nodes
    type: mcp.model.reference_tree_node
```

### 7. TypeRef の bare name 解決順を定義する

TypeRef に bare model name が現れた場合、以下の順で解決する。

1. 同一 file 内の file-private helper model
2. 同一 module 内の public model

QualifiedID を指定した場合は public model のみを対象とする。
file-private helper model は外部 QualifiedID を持たない。

ただし、§8 で同一 module 内の public model と file-private helper model の同名を禁止するため、通常は 1 と 2 の競合は発生しない。
spec 反映時は、本解決順を `docs/spec/type-ref.md` §named model TypeRef、および `docs/spec/naming.md` の file-local object / QualifiedID 周辺に反映する。

### 8. 同一 module 内の public model と file-private helper model の同名を禁止する

file-private helper model は、同一 module 内の public model と同じ id を持ってはならない。

また、同一 file 内の model id は visibility に関係なく一意でなければならない。

| ケース | 判定 |
|---|---|
| 同一 file 内で helper model 同士が同名 | invalid |
| 同一 file 内で public model と helper model が同名 | invalid |
| 同一 module 内の public model と helper model が同名 | invalid |
| 同一 module 内の別 file にある helper model 同士が同名 | valid |
| 別 module の public model と helper model が同名 | valid |
| 別 module の helper model 同士が同名 | valid |

shadowing を許すと、TypeRef の読み手が「この名前は public model なのか file-local helper なのか」を判断しづらくなる。
brewprint は設計意図を明瞭に残す言語であるため、同一 module 内では shadowing を禁止する。

### 9. MCP exposure は既存 private sub node 方針を流用する

file-private helper model は外部 QualifiedID を持たないが、MCP response 内では識別可能でなければならない。
そのため、V01-ADR-058 / MCP schema の private sub node と同じく、file-local synthetic ID を使う。

```text
<file-id>#<local-id>
```

例:

```text
mcp/model/get_reference_tree_response.yaml#reference_tree_node
```

MCP exposure の基本方針は以下とする。

- file-private helper model は public QualifiedID を持たない
- private sub node と同じ synthetic ID 形式で MCP response 内の安定識別子を持つ
- direct query が必要な場合は synthetic ID、または `file` + `local_id` selector を使う
- file-private helper model の listing / signature / inspect / references における詳細な扱いは MCP spec 側で定義する

通常の object listing に helper model を含めるか、private object を含める option を導入するかは、本ADRでは決めない。
詳細な tool response schema は spec 反映時に `docs/spec/mcp/*` で定義する。

### 10. render exposure は後続ADRで扱う。ただし不可視にはしない

file-private helper model は、MCP で問い合わせなければ見えないだけの存在にしてはならない。
人間向け render からも確認できる必要がある。

ただし、本ADRでは具体的な render 表示方式は決めない。
Mermaid DAG 本体への描画、DAG Markdown detail section の `Private models` table、model / schema catalog view などは render 仕様の論点であり、後続ADRで扱う。

本ADRでは、少なくとも以下を制約として残す。

- file-private helper model は人間向け render から不可視にしてはならない
- Processing flow を表す Mermaid DAG 本体に schema detail を混ぜるかどうかは別ADRで判断する
- file単位または catalog単位の Markdown render で helper model を確認できる導線を用意する

## 理由

### なぜ anonymous inline struct ではなく helper model か

V01-ADR-069 で決定した通り、brewprint は anonymous inline struct TypeRef を導入しない。
設計層で匿名 shape を許すと、後から読んだ人間や別 session の LLM が、その shape が何の概念だったのかを復元しにくくなる。

helper shape に名前を付けることで、以下が得られる。

- stable ID
- field / element / value の machine-readable schema
- note による semantic contract
- TypeRef / MCP / render からの参照可能性
- public model へ昇格しやすい migration path

### なぜすべて public model にしないか

すべての helper shape を public model として定義すると、小さな内部 entry shape まで外部参照可能な contract に見えてしまう。

UC-002 の MCP response model には、`nodes[]` entry、`edges[]` entry、endpoint section、impact entry など、特定 response model の内部構造として閉じるのが自然な shape が存在する。
これらをすべて public model にすると、public surface が膨らみ、LLM が「外部から参照してよい schema」と「内部補助 shape」を区別しづらくなる。

file-private helper model により、名前付き schema と public surface 抑制を両立できる。

### なぜ task file / model file の両方で認めるか

helper model を task file に限定すると、model-local な内部 shape を表現できない。

UC-002 で強く現れているのは task 内部の一時データだけではなく、response model の内部 entry shape である。
これは model file 側に置くのが自然であり、task file に置くと public contract schema の定義が task 実装側へ漏れる。

一方で、task-local な intermediate shape や task-local response shape もありうるため、task file 側でも helper model を認める。

### なぜ `visibility:` field を導入しないか

V01-ADR-011 以降、brewprint は file の代表 node を `main: true` で明示し、それ以外を file-private helper として扱う設計を採っている。

model にだけ `visibility:` field を導入すると、task subnode と model helper model で説明体系が分かれる。
人間と LLM の説明コストを下げるため、model でも `main: true` を拡張し、main / non-main の区別で public / file-private を決める。

### なぜ `main: true` を model に拡張するか

ファイル名から representative model を推論する案もある。
しかし V01-ADR-011 では、ファイル名推論ではなく `main: true` による明示を採用した。
その理由は、1ファイル内に代表 node が複数存在する誤りを静的検出しやすくし、設計上の代表を YAML 内で明確にするためである。

model でも同じ理由が当てはまる。
したがって、model file でも `main: true` を使う方が task file と一貫する。

### なぜ同一 module 内 shadowing を禁止するか

file-private helper model が public model を shadow できると、bare TypeRef を読んだときに名前の意味が文脈依存になりすぎる。

プログラミング言語ではローカル shadowing が許される場合が多いが、brewprint は実装言語ではなく設計言語である。
設計上の概念名が同一 module 内で複数の意味を持つと、人間の確認コストと LLM への説明コストが増える。

そのため、同一 module 内では public model と helper model の同名を禁止する。

### なぜ render exposure を本ADRで詳細化しないか

本ADRの主対象は model visibility と名前解決である。
DAG Markdown detail section に private model table を出すか、model/schema catalog view を追加するかは render 仕様の論点であり、責務が異なる。

ただし、file-private helper model が MCP query でしか見えない状態は brewprint の人間向け render 方針に反する。
そのため、本ADRでは「人間向け render から不可視にしてはならない」という制約だけを決め、具体的な表示方式は後続ADRに分離する。

## 却下した代替案

### 代替案A: helper model を導入せず、すべて public model にする

- 利点: 名前解決と MCP exposure は単純
- 欠点: 小さな内部 shape まで public model になり、public surface と model file 数が過剰に増える。外部から参照してよい schema と内部補助 schema の区別が弱くなる

→ 却下。

### 代替案B: anonymous inline struct TypeRef を導入する

```text
list<{ id: str, severity: str }>
```

- 利点: 小さな one-off shape を手軽に書ける
- 欠点: V01-ADR-069 の判断に反する。匿名 shape は設計意図が後から追いにくく、TypeRef variant / compatibility / lint の論点も増える

→ 却下。

### 代替案C: helper model を task file に限定する

- 利点: task-local intermediate shape だけを扱うなら単純
- 欠点: response model 内部の entry shape など、model-local helper model が必要なケースを扱えない。UC-002 で観測された主な需要と合わない

→ 却下。

### 代替案D: `visibility:` field を導入する

```yaml
- id: response
  type: model
  visibility: public

- id: response_entry
  type: model
  visibility: file-private
```

- 利点: visibility が直接書かれて明示的
- 欠点: task subnode の main / non-main 方針と別体系になる。`main: true` と `visibility:` の不整合 validation も必要になる

→ 却下。YAML field は増やさず、`main: true` の有無で visibility を決める。

### 代替案E: ファイル名と同じ id の model を public model と推論する

- 利点: `main: true` を書かずに済む
- 欠点: V01-ADR-011 で避けたファイル名推論に戻る。representative model が YAML 内で明示されず、複数候補の検出も分かりづらくなる

→ 却下。model でも `main: true` を使う。

### 代替案F: file-private helper model による public model shadowing を許可する

- 利点: プログラミング言語のローカル名前空間に近い柔軟性がある
- 欠点: TypeRef の読み手が文脈依存の解釈を強いられる。同一 module 内で同じ設計概念名が複数意味を持ち、人間と LLM の説明コストが増える

→ 却下。同一 module 内の public model と file-private helper model の同名は禁止する。

## 影響

### spec への影響

本ADR受理後、以下の spec 更新が必要になる。

- `docs/spec/nodes.md`
  - model に `main: true` を適用できること
  - model file が 1 public main model + 0個以上の file-private helper model を持てること
  - model file では単一 model のみの場合も public model には `main: true` が必須であること
  - `visibility:` field は導入しないこと
  - `main: true` model は public、mainなし model は file-private とすること

- `docs/spec/type-ref.md`
  - named model TypeRef が同一 file 内 helper model を参照できること
  - TypeRef の bare name 解決順として、同一 file 内 file-private helper model を同一 module 内 public model より優先すること
  - external QualifiedID では helper model を参照できないこと

- `docs/spec/naming.md`
  - file-private helper model が外部 QualifiedID を持たないこと
  - QualifiedID は public model のみを対象にすること
  - 同一 module 内 public model と file-private helper model の同名禁止を定義すること
  - 同一 module 内の別 file にある helper model 同士は衝突しないこと

- `docs/spec/mcp/schema.md` / `docs/spec/mcp/tools/*`
  - file-private helper model の ObjectRef / selector / synthetic ID 方針を追加すること
  - `inspect(file)` / `inspect(model)` / `get_signature` における helper model exposure を定義すること

### diagnostics への影響

本ADRにより、以下の validation diagnostic が必要になる可能性がある。
具体的な code 名と message format は spec / implementation 反映時に定義する。

- model file に `main: true` model がない
- model file に public model が複数ある
- 同一 file 内で model id が重複している
- 同一 module 内の public model と file-private helper model が同名である
- 外部 file / module から file-private helper model を参照している

### MCP への影響

MCP は file-private helper model を public QualifiedID object として扱わない。
一方で、実装判断やレビューでは helper model の中身を取得できる必要があるため、synthetic ID または `file` + `local_id` selector で直接問い合わせ可能にする。

通常の object listing に helper model を含めるか、private object を含める option を導入するかは MCP spec 反映時に決める。
ただし、public listing の既定値では public model と helper model を混ぜない方針を推奨する。

### render への影響

file-private helper model は人間向け render から不可視にしてはならない。
ただし、具体的な表示方式は後続ADRで扱う。

候補としては、以下がある。

- DAG Markdown detail section に `Private models` table を追加する
- model file / schema render に file-private helper model を表示する
- module横断の model catalog / schema catalog view を追加する

本ADRは render exposure の必要性を制約として残すが、具体的な render schema は所有しない。

### UC-002 への影響

本ADRにより、既存の model YAML には public model を明示するための `main: true` 追加 migration が必要になる。

UC-002 の MCP contract YAML では、現在 `any + note` で暫定表現している response 内部 shape の一部を file-private helper model へ移行できる。

代表的な候補には以下がある。

- `get_reference_tree_response.nodes` / `edges` の entry shape
- `list_endpoints_response.tables` / sections / endpoint entry
- `analyze_impact_response.impacts` / summary / coverage

ただし、具体的な migration 対象、順序、完了条件は `docs/tasks/m15-data-layer-expressiveness.md` および UC-002 側の task file で追跡する。

### M15 への影響

本ADRは M15 Phase C の model visibility / helper model 方針を定める。
V01-ADR-069 で anonymous inline struct を導入しないことを決めたため、本ADRはその補完策となる。

M15 では、本ADR acceptance 後に spec / implementation / UC-002 migration へ反映する。
具体的な作業項目は task file で管理する。

### 後続ADR候補

本ADRから、少なくとも以下の後続ADRが必要になる。

- file-private helper model の Markdown render exposure
  - Mermaid DAG 本体ではなく、DAG Markdown detail section / model render でどう表示するか
- model catalog / schema catalog view
  - public model と file-private helper model を module / contract 単位で俯瞰する view

## Evidence

- commit: 49391ff
- impl commit: tbd
- 参考: V01-ADR-011 の main node / sub node 方針、V01-ADR-058 の file-private sub node scope、V01-ADR-069 の anonymous inline struct 不採用、UC-002 MCP contract YAML における `any + note` 暫定表現

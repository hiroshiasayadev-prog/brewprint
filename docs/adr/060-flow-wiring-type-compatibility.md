# 060: v1.1 TypeRef と flow wiring の型互換性

- **status**: accepted
- **date**: 2026-05-03

> こ�EADRは起票時点での決定を記録したスナップショチE��である、E> 現在の仕様�E spec を参照すること、E
## 背景

ADR-058 / ADR-059�E�E14a�E�で B1: subnode scope と B2: return primitive support の実裁E��正を行った結果、E`task.returns.model` および `join.returns.model` で primitive (`any` / `str` / `int` 筁E を許容するようになった、E
一方で、brewprint v1.0 系の型表現は self-hosting (UC-002) に対して不足してぁE��ことが�E確になった、E特に以下�E問題がある、E
- flow wiring における型互換性ルールが存在しなぁE- `any` の代入互換性が未定義
- `foreach` は `$item` の型を list element から導�Eする忁E��があるにもかかわらず、型表現ぁEnamed list model に閉じてぁE��
- cloud/local 墁E��めEMCP tool I/O では `list<T>` / `dict<T>` 相当�E shallow container type が忁E��E- すべての一時的な collection に named model を強制すると、設計ノイズが増えめE
したがって、これ以降�E型整合性 validation は v1.0 補修ではなく、E*v1.1 の TypeRef を前提に設計すめE*、E
こ�EADRは、v1.1 におけめETypeRef の最小形と、それを使っぁEflow wiring の型互換性ルールを定める、Eただし、brewprint を�Eログラミング言語�E型シスチE��にはしなぁE��E本ADRで導�Eするのは built-in container type-ref (`list<T>` / `dict<T>`) までであり、user-defined generics / subtyping / structural typing / class inheritance は導�EしなぁE��E
### 現状の課顁E
`internal/resolve/flow.go` の `buildParamWirings` / `buildJoinParamWirings` は wiring source の参�E解決のみを行い、source側の型と target param の型�E整合性は一刁E��査しなぁE��E
具体侁E

```yaml
# auth/task/login.yaml
nodes:
  - id: fetch_user
    type: task
    returns:
      name: user
      model: user

  - id: charge_payment
    type: task
    params:
      - name: order
        model: order

flow:
  - step: fetch_user
  - step: charge_payment
    params:
      order: fetch_user  # user -> order だが、現状は no error
```

こ�Ewiringは型として明らかに不整合！Eser ↁEorder�E�だが、現行validationは「signature が解決できる」までしか検査しなぁE��めE��ってしまぁE��E
また、v1.0 の named list model 方式だけでは、以下�Eような自然なI/Oを直接表現しづらい、E
```yaml
params:
  - name: files
    model: list<source_file>

returns:
  name: diagnostics
  model: list<diagnostic>
```

`list<T>` / `dict<T>` は full generic ではなく、brewprint の flow / foreach / external boundary を扱ぁE��め�E built-in TypeRef として導�Eする、E
## 決宁E
### 1. TypeRef を導�Eする

v1.1 の params / returns / model field / list element / dict value で使ぁE��参�EめE**TypeRef** と呼ぶ、E
TypeRef は以下�E形を持つ、E
| TypeRef | 侁E| 意味 |
|---|---|---|
| primitive | `str`, `int`, `bool`, `any` | ADR-021 の primitive 予紁E��E|
| named model | `user`, `order`, `catalog.product` | model QID への参�E |
| inline list | `list<user>` | 要素垁E`user` の list を表ぁEinline TypeRef 構文 |
| inline dict | `dict<config>` | value垁E`config` の dict を表ぁEinline TypeRef 構文。key は ADR-021 と同じく常に `str` |

`list<T>` / `dict<T>` は built-in container TypeRef であり、user-defined generic ではなぁE��Eこ�E節でぁE�� `list<T>` / `dict<T>` は inline TypeRef 構文を指す、EDR-021 の named list/dict model は §2 で定める正規化により、型互換性チェチE��時に inline container TypeRef と同じ container shape として扱ぁE��ETypeRef は再帰皁E��定義されるため、構文上�E `list<dict<user>>` のような入れ子も表現できる。ただし、深ぁE�Eれ子を推奨するも�Eではなく、深さ制限また�E lint 方針�E M15 の spec 整琁E��扱ぁE��E型変数 `T` をユーザーが宣言する構文、`model<T>`、`task<T>`、`T extends X`、generic function inference は導�EしなぁE��E
### 2. 既孁Enamed list/dict model の正規化

ADR-021 で導�EされぁEnamed list/dict model は v1.1 でも有効とする、Eただし型互換性チェチE��では、list/dict kind の named model めETypeRef に正規化できる、E
```yaml
- id: user_list
  type: model
  kind: list
  element: user
```

上記�E型互換性チェチE��上、以下と同じ container shape を持つ、E
```txt
list<user>
```

同様に、E
```yaml
- id: config_map
  type: model
  kind: dict
  value: config
```

は型互換性チェチE��上、以下と同じ container shape を持つ、E
```txt
dict<config>
```

named list/dict model の `id` / `note` は、人間�ELLM向けの意味付けとして保持される、Eただし、flow wiring の型互換性では container shape に正規化して比輁E��る、E
侁E

```txt
user_list(kind:list, element:user) -> list<user>       OK
list<user> -> user_list(kind:list, element:user)       OK
user_list(kind:list, element:user) -> list<order>      NG
```

list/dict 以外�E named model は正規化せず、named model として nominal に比輁E��る、Efield構造の一致による互換�E�Etructural typing�E��E行わなぁE��E
### 3. flow wiring type compatibility ルール

任意�E wiring において、source type S から target type T への代入は、named list/dict model めEcontainer TypeRef に正規化したぁE��で、以下�E場合�Eみ valid とする、E
1. **S また�E T ぁE`any`**
2. **primitive 同士で同一**
3. **list/dict 以外�E named model 同士で QID が同一**
4. **list同士で、element TypeRef が互換**
5. **dict同士で、value TypeRef が互換**

それ以外�E不正であり、`incompatible_wiring_type` diagnostic�E�Eeverity: error�E�を出す、E
侁E

```txt
str -> str                         OK
str -> int                         NG
user -> user                       OK
user -> order                      NG
any -> user                        OK
user -> any                        OK
list<user> -> list<user>           OK
list<user> -> list<order>          NG
list<any> -> list<user>            OK
user_list -> list<user>            OK  # user_list.kind=list, element=user
config_map -> dict<config>         OK  # config_map.kind=dict, value=config
str -> user                        NG
```

`any` は container 冁E��めEwildcard として扱ぁE��E
```txt
list<any> -> list<user>    OK
list<user> -> list<any>    OK
dict<any> -> dict<config>  OK
```

### 4. 検証対象 wiring

以下�Eすべての wiring に本ルールを適用する、E
| wiring箁E�� | source の持E��方況E| target |
|---|---|---|
| `step.params` | wiring source 構文 | step task の params |
| `branch.params` | wiring source 構文 | branch node 自身の params |
| `branch.cases[].params` | wiring source 構文 | case entry task の params |
| `fork.branches[].steps[].params` | wiring source 構文 | step task の params |
| `foreach.params` | wiring source 構文 | foreach apply task の params |
| `join.params`�E�暗黙）| 吁Efork branch の terminal step の `returns.name` 一致解決 | join node の params |

`join.params` は wiring 構文を持たず、各 fork branch の terminal step の `returns.name` と join.params の `name` の一致で暗黙に解決される！Epec/edges.md §1-2 join.params の解決�E�。本 ADR ではこ�E暗黁Ewiring も検証対象とする、E
### 5. wiring source の型解決ルール

wiring source 構文ごとに source TypeRef を以下�Eように解決する、E
#### 5-1. node ID 参�E

source = 同ファイル冁E�E node ID また�E QualifiedID、E
- task の場吁E `task.returns.model` めETypeRef として解決する
- join の場吁E `join.returns.model` めETypeRef として解決する
- branch / fork 等、returns を持たなぁEnode の場吁E node 自体�E解決できてめEwiring source としては不正。`invalid_wiring_source` diagnostic�E�Eeverity: error�E�を出ぁE
> 注: foreach の collected asset (`foreach.returns`) を後綁Eflow から参�E可能にする場合�E source id 解決ルールは ADR-061�E�予定）で扱ぁE��本 ADR の node ID 参�Eルールには含めなぁE��E
#### 5-2. `$params.<name>`

source = ファイルの main task の params 配�Eから `name == <name>` の param を引き、その `model` めETypeRef として解決する、E
`$params.<name>` の解釈�E struct 冁E�� field アクセスではなく、main params 名引きである�E�Epec/edges.md §1-1 wiringの記法）、E
#### 5-3. `$item`

`$item` は `foreach.params` 冁E��のみ有効な wiring source とする。foreach 外で `$item` を指定した場合�E `invalid_wiring_source` diagnostic�E�Eeverity: error�E�を出す、E
source = `foreach.over` の解決結果から要素型を引く、E
`foreach.over` は spec/edges.md §1-5 により以下�E2種を取りうめE

| over の持E��E| 要素型解決 |
|---|---|
| node ID�E�Eask/join のID�E�E| source node の returns TypeRef を引き、それが `list<T>` なめE`T` を返す |
| `$params.<name>` | main task params の `<name>` を引き、その TypeRef ぁE`list<T>` なめE`T` を返す |

要素型解決の特殊ケース:

- **over の解決結果ぁE`any` の場吁E*: `$item` の型も `any` として扱ぁE��Eny の伝播�E�。後続�E wiring は any wildcard ルールにより常に互換成竁E- **over の解決結果ぁE`list<T>` の場吁E*: `$item` の型�E `T`
- **over の解決結果ぁEnamed list model の場吁E*: 正規化後�E `list<T>` から `T` を引く
- **over の解決結果ぁElist ではなぁE��吁E*: `invalid_foreach_over_type` diagnostic�E�Eeverity: error�E�を出す。当該 foreach 冁E�E `$item` を含む wiring は型整合性検証めEskip する
- **over の解決自体が失敗する場吁E*�E�Enresolved node 等！E source type が解決できなぁE��め、当該 foreach 冁E�E `$item` を含む wiring は型整合性検証めEskip

> 注: foreach の collected asset (`foreach.returns`) を後綁Eflow から参�E可能にするか、その source id 解決ルールをどぁE��るかは ADR-061�E�予定）で扱ぁE��本 ADR のスコープでは `$item` の入力型解決のみを扱ぁE��E
### 6. 新 diagnostic コーチE
| code | severity | 意味 |
|---|---|---|
| `incompatible_wiring_type` | error | wiring source の型と target param の型が互換しなぁE|
| `invalid_wiring_source` | error | returns 相当�E出力型を持たなぁEnode、また�E有効篁E��外�E `$item` めEwiring source として持E��しぁE|
| `invalid_foreach_over_type` | error | `foreach.over` が指ぁEsource ぁElist として扱えなぁE|
| `invalid_type_ref` | error | TypeRef 構文が不正、また�E TypeRef として扱えなぁEcontainer kind 等を持E��しぁE|

`incompatible_wiring_type` のメチE��ージには source TypeRef / target TypeRef / wiring位置を含めること、E`invalid_type_ref` のメチE��ージには、不正な TypeRef 斁E���Eとそ�E出現位置を含めること、E
### 7. 型解決失敗時の扱ぁE
型互換性チェチE��は、source TypeRef と target TypeRef の両方が正常に解決できた場合�Eみ行う、E
以下�E条件のぁE��れかに該当すめEwiring につぁE��は、型互換性チェチE��を行わず、`incompatible_wiring_type` を発行しなぁE��E
- source TypeRef が解決不�E�E�未解決 node、returns を持たなぁEsource、未定義の collected asset 等！E- target param の TypeRef が解決不�E�E�Earam model が未解決等！E- foreach の場合、`foreach.over` の要素型が解決不�E�E�Einvalid_foreach_over_type` 等！E
こ�E方針により、実裁E�E「既出 diagnostic が存在するか」を検索する忁E���EなぁE��type resolution が失敗した時点で `incompatible_wiring_type` を抑制する、E
### 8. 本 ADR で行わなぁE��と

本 ADR では以下を導�EしなぁE��E
- **subtyping**�E�Eodel 階層・継承による互換�E�E- **structural typing**�E�Eield 構造の一致による互換�E�E- **user-defined generics**�E�Emodel<T>` / `task<T>` / 型変数 `T` の宣言�E�E- **generic type unification**�E�型変数推論！E- **class inheritance / extends / implements**
- **variance rules**�E�Elist<Dog>` めE`list<Animal>` として扱ぁE��！E
`list<T>` / `dict<T>` は built-in TypeRef として扱ぁE��、これ�E user-defined generics ではなぁE��E
### 9. 仕様反映允E
本 ADR 受理後、以下�E spec を更新する、E
- **`docs/spec/nodes.md`**: `param.model` / `returns.model` / `field.type` / `model.element` / `model.value` ぁETypeRef を受け取ることを定義
- **`docs/spec/edges.md` §1�E�Elow:セクション�E�末尾に「§型互換性ルール」節を追加**: 本 ADR §3〜§7 の冁E��を仕様として記述
- **`docs/spec/diagnostics.md`**: `incompatible_wiring_type` / `invalid_wiring_source` / `invalid_foreach_over_type` / `invalid_type_ref` を追加、由来に ADR-060 を記輁E
TypeRef の構文詳細めE`nodes.md` に置くか、新要E`spec/type-ref.md` に刁E��か�E M15 の spec 整琁E��に決める、E
## 琁E��

### なぁEv1.1 TypeRef 前提にするぁE
v1.0 の named list/dict model だけでは、self-hosting (UC-002) めEcloud/local 墁E��のI/Oで忁E��になる一時的な collection 型を表現するには重い、E
すべての `list<diagnostic>` / `list<reference>` / `dict<any>` に named model を強制すると、設計ノイズが増える、E一方で、無制限に深ぁE`list<dict<list<...>>>` を推奨したぁE��けでもなぁE��E
したがって、v1.1 では built-in container TypeRef を許容し、意味を持つ collection は named model に昁E��できる設計にする。深ぁE�Eれ子�E制限や lint は、TypeRef 構文の spec 反映時に扱ぁE��E
### なぁEnominal + built-in container recursion ぁE
brewprint の目皁E�E、�Eログラミング言語�E型シスチE��を�E現することではなく、人間とAIの設計認識を揁E��ることである、E
list/dict 以外�E named model は nominal に扱ぁE��E`user` と `customer` が同ぁEfield 構造を持ってぁE��も、設計上同じ意味とは限らなぁE��E暗黙�E structural typing は設計意図を壊しめE��ぁE��め導�EしなぁE��E
一方、`list<user>` / `dict<config>` は built-in container の中身を比輁E��なければ foreach / collection flow の型整合が取れなぁE��Eそ�Eため、container につぁE��のみ再帰皁E�� TypeRef を比輁E��る、E
### なぁEnamed list/dict model と inline list<T>/dict<T> を互換にするぁE
v1.0 の `kind:list` / `kind:dict` model は既存設計賁E��として残る、EこれめE`list<T>` / `dict<T>` と互換にしなぁE��合、v1.1 への移行時に同じ container shape を二重定義する忁E��が生じる、E
named list/dict model の名前めEnote は意味付けとして保持しつつ、flow wiring の型互換性では container shape に正規化する方が実用皁E��ある、E
### なぁEany を両方向許容�E�Eildcard�E�にするぁE
`any` は ADR-021 §3 で「型不定（最小限�E�」と定義された送E��口、E
- return 側で `any` を使ぁE�Eは「呼び出し�Eに型を強制したくなぁE��signal
- param 側で `any` を使ぁE�Eは「呼び出し�Eの型を問わなぁE��signal

どちら�E方向も blueprint 言語として忁E��であり、片方向�Eみ許容する非対称な設計を選ぶ積極皁E��由がなぁE��E
LLM-friendly な緩衝地帯として、any は両方吁Ewildcard が�E然、E
### なぁE`$item` を含めるぁE
`foreach` は正弁Esupport の構文�E�EDR-016�E�。`$item` への型検証なしでは foreach 冁E�E wiring が「半刁E��か検証されなぁE��状態になる、E
`$item` の要素型解決は user-defined generic ではなく、`foreach.over` の TypeRef から `list<T>` の `T` を読むだけである、E
### なぁE`join.params`�E�暗黁Ewiring�E�を含めるぁE
join.params は wiring 斁E��に直接出なぁE��、各 fork branch の terminal step の returns.name と join.params の name の一致による暗黙�E wiring である、E実質皁E�� wiring であり、検証対象から除外する根拠がなぁE��E
### 却下した代替桁E
#### 代替桁E: 型互換性ルールを導�EしなぁE
- primitive return が解禁された結果、any の代入互換挙動が未定義のまま実裁E��進む
- LLM が生成しぁEblueprint の型不整合が機械検�EされなぁE��まになめE- self-hosting で品質保証としての型検証が機�EしなぁE- ↁE却下。TypeRef を導�Eする以上、flow wiring の互換性ルールは忁E��E
#### 代替桁E: v1.0 の named model QID 一致だけで実裁E��めE
- 実裁E�E簡単だが、`list<T>` / `dict<T>` 導�E時にすぐ壊れめE- foreach の `$item` めEcollected asset の型導�EぁETypeRef と噛み合わなぁE- ↁE却下。これから実裁E��めEvalidation は v1.1 TypeRef 前提にする

#### 代替桁E: full structural typing を導�Eする

- 「struct field の雁E��が一致すれば互換」「kind:list の element が互換なら互換」等�E構造皁E��輁E��導�Eする
- 実裁E��ストが大きく、subtyping / variance の議論を巻き込む
- brewprint の「意味を持つ model は名前で識別する」思想と相性が悪ぁE- ↁE却下。外部shapeとの変換は structural互換ではなぁEadapter / normalize task で明示する

#### 代替桁E: user-defined generics を導�Eする

- `model<T>` / `task<T>` / 型変数推論を導�Eする
- brewprint が小さなプログラミング言語に寁E��すぎめE- v1.1 の目皁E�E AI実裁E��忁E��な設計契紁E�E明確化であり、多相型シスチE��の再現ではなぁE- ↁE却下。built-in container TypeRef までに留めめE
#### 代替桁E: any ↁET を禁止する�E�一方吁Ewildcard�E�E
- TypeScript any ではなぁERust dyn Any 寁E��の設訁E- v1.1 に narrow 構文が存在しなぁE��め、any めEreturn で使った後�E値は事実上どこにめEwiring できなくなめE- ↁE却下。any を実用皁E��送E��口として機�Eさせるには両方向許容が忁E��E
## 影響

### 既存実裁E��の影響

- `internal/semantic` に TypeRef 表現を追加する
- `Param` / `Return` / `ModelField` / `Model.Element` / `Model.Value` から TypeRef を引けるよぁE��する
- 既存�E `ModelName` / `Model` は migration 期間中は保持してよいが、flow wiring validation は TypeRef ベ�Eスに寁E��めE- `internal/resolve/validation.go` に `validateFlowWiringTypes` を追加�E�Ealidate phase で実行！E- `internal/resolve/validation.go` に TypeRef 互換判定�Eルパ�E�E�仮称 `typeRefsCompatible(a, b semantic.TypeRef) bool` 等）を追加
- `internal/resolve/diagnostics.go` 等�E diagnostic コード一覧に `incompatible_wiring_type` / `invalid_wiring_source` / `invalid_foreach_over_type` / `invalid_type_ref` を追加

### 既孁Espec への影響

- `spec/nodes.md` また�E新要E`spec/type-ref.md` に TypeRef 構文を追加
- `spec/edges.md` §1 末尾に「§型互換性ルール」節を新設
- `spec/diagnostics.md` に新 diagnostic 4件を追加

### 既孁EUC への影響

- **UC-001**: 既孁Ewiring はすべて同一型同士で揁E��てぁE��はずであり、回帰チE��トでパス想定。丁E��一不整合があれば bug として修正
- **UC-002**: v1.1 TypeRef 導�E後、MCP tool I/O めEself-hosting blueprint で `list<T>` / `dict<T>` を使った�E構築が可能になめE
### v1.0.0-spec タグへの影響

本 ADR は v1.1 の設計判断であり、v1.0.0-spec の凍結対象には含まなぁE��EM14a ぁEpatch release�E�Ev1.0.1-spec`�E�として v1.0.0-spec から刁E��する�Eに対し、本 ADR は M15 / v1.1 系の基礎となる、E
タグ発行方釁E

- M14b を独立しぁEv1.0.x patch として扱わなぁE- M14b 相当�E flow wiring type validation は M15 / data layer expressiveness の前段また�E一部として扱ぁE- M15 完亁E��に TypeRef / wiring type validation / enum / discriminated object 等を合わせて **`v1.1.0-spec`** タグを発行すめE
ADR-050 §7 / ADR-057 Non-goals が禁じる「v1 篁E��の spec / ADR の遡及修正」には該当しなぁE��E本 ADR は v1.0.0-spec で未定義だった領域に v1.1 ルールを追加する forward 拡張である、E
### M15 への影響

本 ADR の TypeRef と flow wiring type compatibility は M15 / data layer expressiveness の前提となる、EM15 では以下を合わせて扱ぁE��E
- TypeRef 構文の spec 反映
- enum / discriminated object / inline struct の要否判断
- `list<T>` / `dict<T>` の深さ制限また�E lint 方釁E- named list/dict model との互換性
- flow wiring type validation 実裁E
### 関連 ADR への影響

- **ADR-021**: primitive 予紁E��およ�E list/dict model の定義は維持する。ただぁEv1.1 では TypeRef として `list<T>` / `dict<T>` も許容する
- **ADR-016**: foreach 構文は変更なし。`$item` の型解決規則を本 ADR で明文化すめE- **ADR-059**: primitive return の解禁�E本 ADR の前提条件、EDR-060 受理後、ADR-059 §影響に「型互換性ルールは ADR-060 で別途規定」を追記すめE- **ADR-061�E�予定！E*: foreach.returns の collected asset 参�Eルール�E�Eource id 自動生成、`foreach.id` フィールド導�E、E��褁E��ール、ADR-016 との整合）�E ADR-061 で別途扱ぁE
## Evidence

- commit: f507485
- impl commit: 01e7127
- 参老E TypeScript any の両方向互換挙動、nominal typing�E�Eava / C# の class identity�E�、Dagster asset の type 解決

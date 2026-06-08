# V01-ADR-060: v1.1 TypeRef と flow wiring の型互換性

- **status**: accepted
- **date**: 2026-05-03

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

V01-ADR-058 / V01-ADR-059（M14a）で B1: subnode scope と B2: return primitive support の実装修正を行った結果、
`task.returns.model` および `join.returns.model` で primitive (`any` / `str` / `int` 等) を許容するようになった。

一方で、brewprint v1.0 系の型表現は self-hosting (UC-002) に対して不足していることが明確になった。
特に以下の問題がある。

- flow wiring における型互換性ルールが存在しない
- `any` の代入互換性が未定義
- `foreach` は `$item` の型を list element から導出する必要があるにもかかわらず、型表現が named list model に閉じている
- cloud/local 境界や MCP tool I/O では `list<T>` / `dict<T>` 相当の shallow container type が必要
- すべての一時的な collection に named model を強制すると、設計ノイズが増える

したがって、これ以降の型整合性 validation は v1.0 補修ではなく、**v1.1 の TypeRef を前提に設計する**。

このADRは、v1.1 における TypeRef の最小形と、それを使った flow wiring の型互換性ルールを定める。
ただし、brewprint をプログラミング言語の型システムにはしない。
本ADRで導入するのは built-in container type-ref (`list<T>` / `dict<T>`) までであり、user-defined generics / subtyping / structural typing / class inheritance は導入しない。

### 現状の課題

`internal/resolve/flow.go` の `buildParamWirings` / `buildJoinParamWirings` は wiring source の参照解決のみを行い、source側の型と target param の型の整合性は一切検査しない。

具体例:

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

このwiringは型として明らかに不整合（user → order）だが、現行validationは「signature が解決できる」までしか検査しないため通ってしまう。

また、v1.0 の named list model 方式だけでは、以下のような自然なI/Oを直接表現しづらい。

```yaml
params:
  - name: files
    model: list<source_file>

returns:
  name: diagnostics
  model: list<diagnostic>
```

`list<T>` / `dict<T>` は full generic ではなく、brewprint の flow / foreach / external boundary を扱うための built-in TypeRef として導入する。

## 決定

### 1. TypeRef を導入する

v1.1 の params / returns / model field / list element / dict value で使う型参照を **TypeRef** と呼ぶ。

TypeRef は以下の形を持つ。

| TypeRef | 例 | 意味 |
|---|---|---|
| primitive | `str`, `int`, `bool`, `any` | V01-ADR-021 の primitive 予約語 |
| named model | `user`, `order`, `catalog.product` | model QID への参照 |
| inline list | `list<user>` | 要素型 `user` の list を表す inline TypeRef 構文 |
| inline dict | `dict<config>` | value型 `config` の dict を表す inline TypeRef 構文。key は V01-ADR-021 と同じく常に `str` |

`list<T>` / `dict<T>` は built-in container TypeRef であり、user-defined generic ではない。
この節でいう `list<T>` / `dict<T>` は inline TypeRef 構文を指す。V01-ADR-021 の named list/dict model は §2 で定める正規化により、型互換性チェック時に inline container TypeRef と同じ container shape として扱う。
TypeRef は再帰的に定義されるため、構文上は `list<dict<user>>` のような入れ子も表現できる。ただし、深い入れ子を推奨するものではなく、深さ制限または lint 方針は M15 の spec 整理で扱う。
型変数 `T` をユーザーが宣言する構文、`model<T>`、`task<T>`、`T extends X`、generic function inference は導入しない。

### 2. 既存 named list/dict model の正規化

V01-ADR-021 で導入された named list/dict model は v1.1 でも有効とする。
ただし型互換性チェックでは、list/dict kind の named model を TypeRef に正規化できる。

```yaml
- id: user_list
  type: model
  kind: list
  element: user
```

上記は型互換性チェック上、以下と同じ container shape を持つ。

```txt
list<user>
```

同様に、

```yaml
- id: config_map
  type: model
  kind: dict
  value: config
```

は型互換性チェック上、以下と同じ container shape を持つ。

```txt
dict<config>
```

named list/dict model の `id` / `note` は、人間・LLM向けの意味付けとして保持される。
ただし、flow wiring の型互換性では container shape に正規化して比較する。

例:

```txt
user_list(kind:list, element:user) -> list<user>       OK
list<user> -> user_list(kind:list, element:user)       OK
user_list(kind:list, element:user) -> list<order>      NG
```

list/dict 以外の named model は正規化せず、named model として nominal に比較する。
field構造の一致による互換（structural typing）は行わない。

### 3. flow wiring type compatibility ルール

任意の wiring において、source type S から target type T への代入は、named list/dict model を container TypeRef に正規化したうえで、以下の場合のみ valid とする。

1. **S または T が `any`**
2. **primitive 同士で同一**
3. **list/dict 以外の named model 同士で QID が同一**
4. **list同士で、element TypeRef が互換**
5. **dict同士で、value TypeRef が互換**

それ以外は不正であり、`incompatible_wiring_type` diagnostic（severity: error）を出す。

例:

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

`any` は container 内でも wildcard として扱う。

```txt
list<any> -> list<user>    OK
list<user> -> list<any>    OK
dict<any> -> dict<config>  OK
```

### 4. 検証対象 wiring

以下のすべての wiring に本ルールを適用する。

| wiring箇所 | source の指定方法 | target |
|---|---|---|
| `step.params` | wiring source 構文 | step task の params |
| `branch.params` | wiring source 構文 | branch node 自身の params |
| `branch.cases[].params` | wiring source 構文 | case entry task の params |
| `fork.branches[].steps[].params` | wiring source 構文 | step task の params |
| `foreach.params` | wiring source 構文 | foreach apply task の params |
| `join.params`（暗黙）| 各 fork branch の terminal step の `returns.name` 一致解決 | join node の params |

`join.params` は wiring 構文を持たず、各 fork branch の terminal step の `returns.name` と join.params の `name` の一致で暗黙に解決される（spec/edges.md §1-2 join.params の解決）。本 ADR ではこの暗黙 wiring も検証対象とする。

### 5. wiring source の型解決ルール

wiring source 構文ごとに source TypeRef を以下のように解決する。

#### 5-1. node ID 参照

source = 同ファイル内の node ID または QualifiedID。

- task の場合: `task.returns.model` を TypeRef として解決する
- join の場合: `join.returns.model` を TypeRef として解決する
- branch / fork 等、returns を持たない node の場合: node 自体は解決できても wiring source としては不正。`invalid_wiring_source` diagnostic（severity: error）を出す

> 注: foreach の collected asset (`foreach.returns`) を後続 flow から参照可能にする場合の source id 解決ルールは V01-ADR-061（予定）で扱う。本 ADR の node ID 参照ルールには含めない。

#### 5-2. `$params.<name>`

source = ファイルの main task の params 配列から `name == <name>` の param を引き、その `model` を TypeRef として解決する。

`$params.<name>` の解釈は struct 内部 field アクセスではなく、main params 名引きである（spec/edges.md §1-1 wiringの記法）。

#### 5-3. `$item`

`$item` は `foreach.params` 内でのみ有効な wiring source とする。foreach 外で `$item` を指定した場合は `invalid_wiring_source` diagnostic（severity: error）を出す。

source = `foreach.over` の解決結果から要素型を引く。

`foreach.over` は spec/edges.md §1-5 により以下の2種を取りうる:

| over の指定 | 要素型解決 |
|---|---|
| node ID（task/join のID） | source node の returns TypeRef を引き、それが `list<T>` なら `T` を返す |
| `$params.<name>` | main task params の `<name>` を引き、その TypeRef が `list<T>` なら `T` を返す |

要素型解決の特殊ケース:

- **over の解決結果が `any` の場合**: `$item` の型も `any` として扱う（any の伝播）。後続の wiring は any wildcard ルールにより常に互換成立
- **over の解決結果が `list<T>` の場合**: `$item` の型は `T`
- **over の解決結果が named list model の場合**: 正規化後の `list<T>` から `T` を引く
- **over の解決結果が list ではない場合**: `invalid_foreach_over_type` diagnostic（severity: error）を出す。当該 foreach 内の `$item` を含む wiring は型整合性検証を skip する
- **over の解決自体が失敗する場合**（unresolved node 等）: source type が解決できないため、当該 foreach 内の `$item` を含む wiring は型整合性検証を skip

> 注: foreach の collected asset (`foreach.returns`) を後続 flow から参照可能にするか、その source id 解決ルールをどうするかは V01-ADR-061（予定）で扱う。本 ADR のスコープでは `$item` の入力型解決のみを扱う。

### 6. 新 diagnostic コード

| code | severity | 意味 |
|---|---|---|
| `incompatible_wiring_type` | error | wiring source の型と target param の型が互換しない |
| `invalid_wiring_source` | error | returns 相当の出力型を持たない node、または有効範囲外の `$item` を wiring source として指定した |
| `invalid_foreach_over_type` | error | `foreach.over` が指す source が list として扱えない |
| `invalid_type_ref` | error | TypeRef 構文が不正、または TypeRef として扱えない container kind 等を指定した |

`incompatible_wiring_type` のメッセージには source TypeRef / target TypeRef / wiring位置を含めること。
`invalid_type_ref` のメッセージには、不正な TypeRef 文字列とその出現位置を含めること。

### 7. 型解決失敗時の扱い

型互換性チェックは、source TypeRef と target TypeRef の両方が正常に解決できた場合のみ行う。

以下の条件のいずれかに該当する wiring については、型互換性チェックを行わず、`incompatible_wiring_type` を発行しない。

- source TypeRef が解決不能（未解決 node、returns を持たない source、未定義の collected asset 等）
- target param の TypeRef が解決不能（param model が未解決等）
- foreach の場合、`foreach.over` の要素型が解決不能（`invalid_foreach_over_type` 等）

この方針により、実装は「既出 diagnostic が存在するか」を検索する必要はない。type resolution が失敗した時点で `incompatible_wiring_type` を抑制する。

### 8. 本 ADR で行わないこと

本 ADR では以下を導入しない。

- **subtyping**（model 階層・継承による互換）
- **structural typing**（field 構造の一致による互換）
- **user-defined generics**（`model<T>` / `task<T>` / 型変数 `T` の宣言）
- **generic type unification**（型変数推論）
- **class inheritance / extends / implements**
- **variance rules**（`list<Dog>` を `list<Animal>` として扱う等）

`list<T>` / `dict<T>` は built-in TypeRef として扱うが、これは user-defined generics ではない。

### 9. 仕様反映先

本 ADR 受理後、以下の spec を更新する。

- **`docs/spec/nodes.md`**: `param.model` / `returns.model` / `field.type` / `model.element` / `model.value` が TypeRef を受け取ることを定義
- **`docs/spec/edges.md` §1（flow:セクション）末尾に「§型互換性ルール」節を追加**: 本 ADR §3〜§7 の内容を仕様として記述
- **`docs/spec/diagnostics.md`**: `incompatible_wiring_type` / `invalid_wiring_source` / `invalid_foreach_over_type` / `invalid_type_ref` を追加、由来に V01-ADR-060 を記載

TypeRef の構文詳細を `nodes.md` に置くか、新規 `spec/type-ref.md` に切るかは M15 の spec 整理時に決める。

## 理由

### なぜ v1.1 TypeRef 前提にするか

v1.0 の named list/dict model だけでは、self-hosting (UC-002) や cloud/local 境界のI/Oで必要になる一時的な collection 型を表現するには重い。

すべての `list<diagnostic>` / `list<reference>` / `dict<any>` に named model を強制すると、設計ノイズが増える。
一方で、無制限に深い `list<dict<list<...>>>` を推奨したいわけでもない。

したがって、v1.1 では built-in container TypeRef を許容し、意味を持つ collection は named model に昇格できる設計にする。深い入れ子の制限や lint は、TypeRef 構文の spec 反映時に扱う。

### なぜ nominal + built-in container recursion か

brewprint の目的は、プログラミング言語の型システムを再現することではなく、人間とAIの設計認識を揃えることである。

list/dict 以外の named model は nominal に扱う。
`user` と `customer` が同じ field 構造を持っていても、設計上同じ意味とは限らない。
暗黙の structural typing は設計意図を壊しやすいため導入しない。

一方、`list<user>` / `dict<config>` は built-in container の中身を比較しなければ foreach / collection flow の型整合が取れない。
そのため、container についてのみ再帰的に TypeRef を比較する。

### なぜ named list/dict model と inline list<T>/dict<T> を互換にするか

v1.0 の `kind:list` / `kind:dict` model は既存設計資産として残る。
これを `list<T>` / `dict<T>` と互換にしない場合、v1.1 への移行時に同じ container shape を二重定義する必要が生じる。

named list/dict model の名前や note は意味付けとして保持しつつ、flow wiring の型互換性では container shape に正規化する方が実用的である。

### なぜ any を両方向許容（wildcard）にするか

`any` は V01-ADR-021 §3 で「型不定（最小限）」と定義された逃げ口。

- return 側で `any` を使うのは「呼び出し側に型を強制したくない」signal
- param 側で `any` を使うのは「呼び出し側の型を問わない」signal

どちらの方向も blueprint 言語として必要であり、片方向のみ許容する非対称な設計を選ぶ積極的理由がない。

LLM-friendly な緩衝地帯として、any は両方向 wildcard が自然。

### なぜ `$item` を含めるか

`foreach` は正式 support の構文（V01-ADR-016）。`$item` への型検証なしでは foreach 内の wiring が「半分しか検証されない」状態になる。

`$item` の要素型解決は user-defined generic ではなく、`foreach.over` の TypeRef から `list<T>` の `T` を読むだけである。

### なぜ `join.params`（暗黙 wiring）を含めるか

join.params は wiring 文法に直接出ないが、各 fork branch の terminal step の returns.name と join.params の name の一致による暗黙の wiring である。
実質的に wiring であり、検証対象から除外する根拠がない。

### 却下した代替案

#### 代替案A: 型互換性ルールを導入しない

- primitive return が解禁された結果、any の代入互換挙動が未定義のまま実装が進む
- LLM が生成した blueprint の型不整合が機械検出されないままになる
- self-hosting で品質保証としての型検証が機能しない
- → 却下。TypeRef を導入する以上、flow wiring の互換性ルールは必要

#### 代替案B: v1.0 の named model QID 一致だけで実装する

- 実装は簡単だが、`list<T>` / `dict<T>` 導入時にすぐ壊れる
- foreach の `$item` や collected asset の型導出が TypeRef と噛み合わない
- → 却下。これから実装する validation は v1.1 TypeRef 前提にする

#### 代替案C: full structural typing を導入する

- 「struct field の集合が一致すれば互換」「kind:list の element が互換なら互換」等の構造的比較を導入する
- 実装コストが大きく、subtyping / variance の議論を巻き込む
- brewprint の「意味を持つ model は名前で識別する」思想と相性が悪い
- → 却下。外部shapeとの変換は structural互換ではなく adapter / normalize task で明示する

#### 代替案D: user-defined generics を導入する

- `model<T>` / `task<T>` / 型変数推論を導入する
- brewprint が小さなプログラミング言語に寄りすぎる
- v1.1 の目的は AI実装に必要な設計契約の明確化であり、多相型システムの再現ではない
- → 却下。built-in container TypeRef までに留める

#### 代替案E: any → T を禁止する（一方向 wildcard）

- TypeScript any ではなく Rust dyn Any 寄りの設計
- v1.1 に narrow 構文が存在しないため、any を return で使った後の値は事実上どこにも wiring できなくなる
- → 却下。any を実用的な逃げ口として機能させるには両方向許容が必要

## 影響

### 既存実装への影響

- `internal/semantic` に TypeRef 表現を追加する
- `Param` / `Return` / `ModelField` / `Model.Element` / `Model.Value` から TypeRef を引けるようにする
- 既存の `ModelName` / `Model` は migration 期間中は保持してよいが、flow wiring validation は TypeRef ベースに寄せる
- `internal/resolve/validation.go` に `validateFlowWiringTypes` を追加（validate phase で実行）
- `internal/resolve/validation.go` に TypeRef 互換判定ヘルパー（仮称 `typeRefsCompatible(a, b semantic.TypeRef) bool` 等）を追加
- `internal/resolve/diagnostics.go` 等の diagnostic コード一覧に `incompatible_wiring_type` / `invalid_wiring_source` / `invalid_foreach_over_type` / `invalid_type_ref` を追加

### 既存 spec への影響

- `spec/nodes.md` または新規 `spec/type-ref.md` に TypeRef 構文を追加
- `spec/edges.md` §1 末尾に「§型互換性ルール」節を新設
- `spec/diagnostics.md` に新 diagnostic 4件を追加

### 既存 UC への影響

- **UC-001**: 既存 wiring はすべて同一型同士で揃っているはずであり、回帰テストでパス想定。万が一不整合があれば bug として修正
- **UC-002**: v1.1 TypeRef 導入後、MCP tool I/O や self-hosting blueprint で `list<T>` / `dict<T>` を使った再構築が可能になる

### v1.0.0-spec タグへの影響

本 ADR は v1.1 の設計判断であり、v1.0.0-spec の凍結対象には含まない。
M14a が patch release（`v1.0.1-spec`）として v1.0.0-spec から分岐するのに対し、本 ADR は M15 / v1.1 系の基礎となる。

タグ発行方針:

- M14b を独立した v1.0.x patch として扱わない
- M14b 相当の flow wiring type validation は M15 / data layer expressiveness の前段または一部として扱う
- M15 完了時に TypeRef / wiring type validation / enum / discriminated object 等を合わせて **`v1.1.0-spec`** タグを発行する

V01-ADR-050 §7 / V01-ADR-057 Non-goals が禁じる「v1 範囲の spec / ADR の遡及修正」には該当しない。
本 ADR は v1.0.0-spec で未定義だった領域に v1.1 ルールを追加する forward 拡張である。

### M15 への影響

本 ADR の TypeRef と flow wiring type compatibility は M15 / data layer expressiveness の前提となる。
M15 では以下を合わせて扱う。

- TypeRef 構文の spec 反映
- enum / discriminated object / inline struct の要否判断
- `list<T>` / `dict<T>` の深さ制限または lint 方針
- named list/dict model との互換性
- flow wiring type validation 実装

### 関連 ADR への影響

- **V01-ADR-021**: primitive 予約語および list/dict model の定義は維持する。ただし v1.1 では TypeRef として `list<T>` / `dict<T>` も許容する
- **V01-ADR-016**: foreach 構文は変更なし。`$item` の型解決規則を本 ADR で明文化する
- **V01-ADR-059**: primitive return の解禁は本 ADR の前提条件。V01-ADR-060 受理後、V01-ADR-059 §影響に「型互換性ルールは V01-ADR-060 で別途規定」を追記する
- **V01-ADR-061（予定）**: foreach.returns の collected asset 参照ルール（source id 自動生成、`foreach.id` フィールド導入、重複ルール、V01-ADR-016 との整合）は V01-ADR-061 で別途扱う

## Evidence

- commit: f507485
- impl commit: 01e7127
- 参考: TypeScript any の両方向互換挙動、nominal typing（Java / C# の class identity）、Dagster asset の type 解決

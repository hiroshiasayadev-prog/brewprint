---
scope: docs/spec/type-ref.md
status: confirmed
last_updated: 2026-05-31
summary: >
  brewprint v1.1 TypeRef構文と型互換性の基礎定義。
  primitive / named model / inline list / inline dict、named list/dict model の正規化、
  enum model compatibility、および container TypeRef safety / warning 境界を定義する。
depends_on:
  - docs/adr/021-model-field-structure.md
  - docs/adr/060-flow-wiring-type-compatibility.md
  - docs/adr/067-enum-model.md
  - docs/adr/069-type-ref-container-complexity-lint.md
  - docs/adr/070-model-visibility-file-private-helper-model.md
---

# TypeRef仕様

TypeRef は、brewprint v1.1 で params / returns / model field / list element / dict value に使う型参照である。

TypeRef は runtime 型システムではなく、人間とLLMが共有する設計契約を静的に検証するための最小表現である。brewprint は TypeRef によって flow wiring の型互換性、foreach の `$item` 型導出、外部境界の shallow container I/O を扱う。

> 由来: ADR-060 §1

---

## 1. TypeRefを受け取るフィールド

TypeRef は以下のフィールドで使う。

| 場所 | フィールド | 意味 |
|---|---|---|
| task / branch / fork / join params | `params[].model` | 入力paramの型 |
| task / join returns | `returns.model` | 出力assetの型 |
| struct model field | `fields[].type` | fieldの型 |
| list model | `element` | list要素の型 |
| dict model | `value` | dict値の型 |

`store.of`、`initializes[].model`、`event.payload.model` は本仕様では TypeRef 適用対象に含めない。これらは既存の model-id 参照として扱う。

TypeRef は上記のフィールド値として直接指定されるほか、foreach の `$item` の型導出にも使われる。`$item` は TypeRef を受け取るフィールドではなく、`foreach.over` の TypeRef から flow wiring 解決時に導出される。詳細は [edges.md](./edges.md) §1-7 を参照する。

TypeRef を受け取るフィールドの詳細なスキーマは [nodes.md](./nodes.md) を参照する。

> 由来: ADR-060 §1, §9

---

## 2. TypeRef構文

TypeRef は以下のいずれかである。

| 種別 | 例 | 意味 |
|---|---|---|
| primitive | `str`, `int`, `float`, `bool`, `bytes`, `datetime`, `any` | primitive予約語 |
| named model | `user`, `order`, `catalog.product` | model ID または model QID への参照 |
| inline list | `list<user>` | 要素型 `user` の list |
| inline dict | `dict<config>` | value型 `config` の dict。key は常に `str` |

```yaml
params:
  - name: files
    model: list<source_file>

returns:
  name: diagnostics
  model: list<diagnostic>

nodes:
  - id: module_config
    type: model
    kind: struct
    fields:
      - name: settings
        type: dict<any>
```

`list<T>` / `dict<T>` は built-in container TypeRef であり、user-defined generic ではない。ユーザーが型変数 `T` を宣言する構文、`model<T>`、`task<T>`、`T extends X`、generic function inference は存在しない。

> 由来: ADR-060 §1

---

## 3. primitive予約語

以下の語は primitive TypeRef として扱う。model ID として定義してはならない。

この一覧は ADR-021 §3 の primitive 予約語を網羅したものである。ADR-060 §1 の primitive 表は代表例であり、完全な一覧は本節を正とする。

| primitive | 意味 |
|---|---|
| `str` | 文字列 |
| `int` | 整数 |
| `float` | 浮動小数点数 |
| `bool` | 真偽値 |
| `bytes` | バイト列 |
| `datetime` | 日時 |
| `any` | 型不定。flow wiring compatibility では両方向 wildcard として扱う |

`any` は設計上の逃げ口であり、使用は最小限にする。ただし v1.1 には narrow 構文がないため、互換性判定では source / target のどちらに現れても互換とする。

> 由来: ADR-021 §3, ADR-060 §1

---

## 4. named model TypeRef

primitive でも inline container でもない TypeRef は named model 参照として扱う。enum model も named model TypeRef として参照する。

Task-file helper model の基本 semantics は [nodes.md](./nodes.md#task-file-private-helper-model-semantics) が定義する。本節では、TypeRef resolver における解決順序だけを定義する。

Task file 内で bare named model TypeRef が現れた場合、resolver は以下の順で解決する。

1. primitive 予約語に一致する場合は primitive TypeRef として扱う。
2. `list<T>` / `dict<T>` は inline container TypeRef として内側の `T` を再帰的に解決する。
3. named model の bare name は、同一 YAML file 内の file-private helper model を先に探す。
4. 該当 helper model がない場合、同一 module 内の public model を探す。
5. それでも解決できない場合、出現箇所に応じて未解決 diagnostic を出す。

QualifiedID を使った named model TypeRef は public model のみを対象とし、task-file helper model には解決しない。

同一 file 内の task-file helper model に解決できることは、その参照がすべての TypeRef use-site で valid であることを意味しない。task / branch / fork / join の `params[].model` から task-file private helper model を参照した場合、解決後に `invalid_private_model_reference` error とする。task / join の `returns.model` から同一 file 内の private helper model を参照することは valid とし、minimum scope では diagnostic を出さない。

```yaml
nodes:
  - id: get_preview
    type: task
    main: true
    returns:
      name: preview
      model: preview_response

  - id: preview_response
    type: model
    kind: struct
    fields:
      - name: items
        type: list<preview_item>

  - id: preview_item
    type: model
    kind: struct
    fields:
      - name: title
        type: str
```

この例では、`returns.model: preview_response` と `fields[].type: list<preview_item>` は同一 file 内の helper model に解決する。

```yaml
params:
  - name: user
    model: user

returns:
  name: order
  model: commerce.order
```

list/dict 以外の named model は nominal に比較する。`user` と `customer` が同じ fields を持っていても、型互換とはみなさない。enum model も同じく nominal に比較し、underlying JSON 表現が string であっても `str` とは暗黙互換にしない。

外部shapeとの変換が必要な場合は、adapter / normalize task を明示して別の asset を生成する。

Public model と task-file helper model の同名衝突 rule は [naming.md](./naming.md) §4.1 が所有する。正常な YAML では同一 module 内の public model が同一 file 内の helper model に shadow される状態は作れない。

> 由来: ADR-060 §2, §3; ADR-070 §6〜§8

---

## 5. inline list / dict TypeRef

`list<T>` は要素型 `T` を持つ list を表す。

```yaml
params:
  - name: users
    model: list<user>
```

`dict<T>` は value型 `T` を持つ dict を表す。key は常に `str` であり、key型を指定する構文は存在しない。

```yaml
params:
  - name: config_by_name
    model: dict<config>
```

TypeRef は再帰的に定義されるため、構文上は以下も有効である。

```txt
list<any>
list<dict<user>>
dict<list<diagnostic>>
```

nested `list<T>` / `dict<T>` は構文上 valid のまま維持する。通常の nested TypeRef は validation error にしない。

ただし parser / implementation safety のため、container nesting depth が 16 を超える TypeRef は `invalid_type_ref` error とする。container nesting depth は `list<T>` / `dict<T>` の入れ子数で数え、primitive / named model は depth 0 とする。

| TypeRef | depth |
|---|---:|
| `diagnostic` | 0 |
| `list<diagnostic>` | 1 |
| `dict<list<diagnostic>>` | 2 |
| `list<dict<list<any>>>` | 3 |

v1.1 では anonymous inline struct TypeRef は導入しない。以下のような構文は TypeRef として扱わない。

```txt
list<{ id: str, severity: str }>
```

`dict<T>` は value 型だけでは key semantics を表現できないため、field 名 / model 名 / `note` のいずれかで key の意味を明示することを推奨する。

> 由来: ADR-060 §1, ADR-069 §1〜§6, §10

---

## 6. named list/dict model の正規化

ADR-021 の named list/dict model は v1.1 でも有効である。

```yaml
- id: user_list
  type: model
  kind: list
  element: user
```

型互換性チェックでは、上記は以下と同じ container shape として正規化する。

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

は以下と互換な container shape として扱う。

```txt
dict<config>
```

named list/dict model の `id` / `note` は、人間・LLM向けの意味付けとして保持される。ただし flow wiring の型互換性では container shape に正規化して比較する。

```txt
user_list -> list<user>      OK
list<user> -> user_list      OK
user_list -> list<order>     NG
```

struct model は正規化せず、named model として nominal に比較する。

> 由来: ADR-060 §2

---

## 7. TypeRef構文エラー

TypeRef として解釈できない文字列は `invalid_type_ref` diagnostic とする。

例:

```txt
list<
dict<>
list<user
map<user>
```

`map<user>` のように `list` / `dict` 以外の container kind を使った場合も、TypeRef として扱えないため `invalid_type_ref` とする。

TypeRef 構文は valid だが内部の named model が解決できない場合は、出現箇所に応じて既存の未解決 diagnostic を使う。`fields[].type` では `unresolved_field_type`、`params[].model` / `returns.model` / `model.element` / `model.value` では `unresolved_model` を使う。構文自体が壊れている場合、または TypeRef として扱えない container kind を指定した場合のみ `invalid_type_ref` を出す。

Task-file helper model が参照元の名前解決 scope に存在しない場合、上記の既存未解決 diagnostic を出してよい。TypeRef が同一 file 内の helper model に解決できたが、その use-site が `params[].model` である場合は未解決ではなく `invalid_private_model_reference` を出す。

TypeRef の解決に失敗した場合、その TypeRef を使う flow wiring では `incompatible_wiring_type` を追加で発行しない。型互換性チェックは source TypeRef と target TypeRef の両方が正常に解決できた場合のみ行う。

Parser safety limit を超えた TypeRef も `invalid_type_ref` とする。この limit は可読性 lint ではなく、parser / implementation safety のための hard error である。

> 由来: ADR-060 §6, §7, M15 Phase A task, ADR-069 §2, §10

---

## 8. enum model TypeRef compatibility

enum model は、既存 TypeRef の named model 参照として扱う。TypeRef に inline enum variant は存在しない。

```yaml
nodes:
  - id: diagnostic
    type: model
    kind: struct
    fields:
      - name: severity
        type: mcp_diagnostic_severity
```

enum model の型互換性は nominal である。

```txt
mcp_diagnostic_severity -> mcp_diagnostic_severity  OK
mcp_diagnostic_severity -> str                      NG
str -> mcp_diagnostic_severity                      NG
mcp_diagnostic_severity -> impact_severity          NG
any -> mcp_diagnostic_severity                      OK
mcp_diagnostic_severity -> any                      OK
```

`enum` と `str` の暗黙互換は導入しない。`any` は ADR-060 の既存ルール通り、source / target のどちらに現れても互換とする。

> 由来: ADR-067 §3〜§5

---

## 9. opaque container TypeRef warning

container TypeRef の内部に `any` が含まれ、shape の意味が named model に回収されていない場合、`opaque_type_ref` warning diagnostic の対象とする。

対象例:

```txt
list<any>
dict<any>
list<dict<any>>
dict<list<any>>
list<dict<list<any>>>
```

`opaque_type_ref` は validation 成功扱いの warning であり、message では必要に応じて named model への切り出しを促す。

`opaque_type_ref` は container TypeRef 内の `any` を対象とする debt visibility baseline であり、bare `any` field や `any + note` の主要 response shape 全体を warning 化または解消するものではない。

`unclear_dict_key` / `deep_type_ref` は将来 lint 候補であり、v1.1 minimum の diagnostic としては追加しない。

> 由来: ADR-069 §4〜§10

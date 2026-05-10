# 067: enum model 導入

- **status**: proposed
- **date**: 2026-05-09

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

M15 Phase C では、UC-002 self-hosting で見つかった data layer 表現力不足を整理している。

UC-002 Phase A の MCP公開contract YAML では、MCP schema 上は閉じた値集合である項目を、現状すべて `type: str` + `note` で表現している。

代表例:

- `diagnostic.severity`: `error` / `warning` / `info` / `hint`
- `object_ref.object`: `node` / `view` / `transition` / `asset` / `field` / `file` / `primitive`
- `get_reference_tree.direction`: `out` / `in` / `both`
- `analyze_impact` の impact severity: `breaking` / `warning` / `info`
- `analyze_impact` の fixability: `mechanical` / `suggested` / `manual_review` / `unknown`

`str + note` は人間には読めるが、validator / MCP schema / LLM にとって machine-readable な制約ではない。
そのため、以下ができない。

- validator が許容値集合を検査する
- MCP schema 生成時に enum として出力する
- LLM が自由文字列ではなく閉じた語彙として扱う
- self-hosting において brewprint 自身の public contract を型として表現する

ADR-021 では `model.kind: scalar` を廃止し、型エイリアス的な意味付けは `note` で担う方針にした。
しかし enum は単なる型エイリアスではなく、有限の値集合という machine-readable constraint である。
したがって、`str` + `note` ではなく model として明示する価値がある。

一方で、ADR-060 で導入した TypeRef は primitive / named model / inline `list<T>` / inline `dict<T>` の最小構成であり、M15 Phase A の TypeRef variant をむやみに増やしたくない。

本ADRは、TypeRef 構文を拡張せずに enum を v1.1 へ導入する最小形を定める。

### 判断根拠としての UC-002 実例の扱い

本ADRで参照する UC-002 の YAML は、ADR起票時点の self-hosting 実例である。
これらの YAML は今後の M15 Phase C 実装・spec反映・UC-002 migration によって更新されうるため、本文中の具体的な field 名や `any` / `str + note` の配置は恒久仕様ではない。

本ADRが判断根拠として扱うのは、個々の YAML shape そのものではなく、UC-002 において以下の性質を持つ値集合が繰り返し現れたという事実である。

- 値集合が有限である
- 外部 contract 上、自由文字列ではなく閉じた語彙として扱う必要がある
- 現行表現では `str + note` に閉じており、validator / MCP schema / LLM に machine-readable な制約として渡らない
- TypeRef variant を増やさず named model として表現できる

したがって、本文中の UC-002 例は実装固定の仕様例ではなく、enum model 導入判断の evidence として扱う。
具体的な fixture migration の範囲や順序は M15 / UC-002 の task file で管理する。

## 決定

### 1. `model.kind: enum` を追加する

brewprint v1.1 の model kind に `enum` を追加する。

```yaml
nodes:
  - id: mcp_diagnostic_severity
    type: model
    kind: enum
    values:
      - error
      - warning
      - info
      - hint
```

`enum` は string-valued finite vocabulary を表す model kind である。
JSON / MCP schema 等への出力では string enum として扱う。

### 2. enum model は `values: []` を持つ

`kind: enum` の model は `values` を必須とする。

`values` の制約:

- non-empty list である
- 各要素は string である
- 空文字は不可
- 同一 enum model 内で値は重複不可
- 値の順序は表示・schema生成時の順序として保持してよいが、型互換性の意味には使わない

値ごとの `note` / `label` / `deprecated` 等の metadata は v1.1 では導入しない。
必要になった場合は別ADRで扱う。

### 3. 使用側は既存 TypeRef の named model として参照する

enum model を使う側は、既存 TypeRef の named model 参照として指定する。

```yaml
nodes:
  - id: diagnostic
    type: model
    kind: struct
    fields:
      - name: severity
        type: mcp_diagnostic_severity
      - name: code
        type: str
      - name: message
        type: str
```

TypeRef に inline enum variant は追加しない。
`mcp_diagnostic_severity` は named model TypeRef として解決される。

### 4. enum の型互換性は nominal に扱う

ADR-060 の TypeRef compatibility において、enum model は list/dict 以外の named model と同じく nominal に扱う。

互換例:

```text
mcp_diagnostic_severity -> mcp_diagnostic_severity  OK
mcp_diagnostic_severity -> str                      NG
str -> mcp_diagnostic_severity                      NG
mcp_diagnostic_severity -> impact_severity          NG
any -> mcp_diagnostic_severity                      OK
mcp_diagnostic_severity -> any                      OK
```

`any` は ADR-060 の既存ルール通り、両方向 wildcard として扱う。

`enum` と `str` の暗黙互換は導入しない。
暗黙互換を許すと、enum の値集合制約を flow wiring validation 上すり抜けられるためである。

### 5. inline enum TypeRef は導入しない

以下のような inline enum TypeRef 構文は v1.1 では導入しない。

```text
enum<error|warning|info|hint>
```

enum は名前を持つ model として定義する。

これにより、ADR-060 の TypeRef variant 追加に戻らず、既存の named model TypeRef の範囲内で enum を導入できる。

### 6. enum validation diagnostic を追加する

新しい diagnostic code として以下を追加する。

| code | severity | 意味 |
|---|---|---|
| `invalid_enum_model` | error | `kind: enum` の定義が不正。`values` 欠落、空、非string値、空文字など |
| `duplicate_enum_value` | error | 同一 enum model 内で `values` が重複している |

`invalid_enum_value` は v1.1 の初期導入では追加しない。
現時点の brewprint YAML は主に schema / model 定義であり、enum-typed field に対する具体的な runtime literal 値を保持しないためである。

将来、example / default / literal payload / fixture data を YAML 上で検証する場合は、その時点で `invalid_enum_value` を追加する。

### 7. UC-002 の移行対象は public contract の共通 enum から始める

ADR-067 acceptance 後、UC-002 の `str + note` enum 候補を段階的に enum model へ移行する。
初期移行対象は以下とする。

| enum model | values | 主な置換対象 |
|---|---|---|
| `mcp_object_type` | `node` / `view` / `transition` / `asset` / `field` / `file` / `primitive` | `object_selector.object`, `object_ref.object` |
| `mcp_diagnostic_severity` | `error` / `warning` / `info` / `hint` | `diagnostic.severity` |
| `reference_tree_direction` | `out` / `in` / `both` | `get_reference_tree_request.direction`, `get_reference_tree_response.direction` |

`analyze_impact_response.impacts` 内の `impact.severity` / `fixability` は、現状 `impacts: any` の note 内に閉じている。
これらを enum 化するには、先に `impact_entry` model を切り、`impacts: list<impact_entry>` へ移行する必要がある。
そのため、`impact_severity` / `impact_fixability` の導入と `impact_entry` model 化は ADR-067 acceptance 後の UC-002 migration task として扱う。

## 理由

### なぜ enum を v1.1 に入れるか

UC-002 は brewprint 自身の MCP public contract を blueprint 化する self-hosting 実例である。
この contract には閉じた語彙が複数存在する。

それらを `str + note` に留めると、人間向け説明としては成立するが、機械的な制約としては失われる。
MCP schema 生成、validator、LLM の設計理解にとって enum は有効な最小拡張である。

### なぜ `model.kind: enum` か

既存の brewprint data layer では、型として意味を持つものは model に置く。
`struct` / `list` / `dict` と同じく、enum も再利用可能な型定義であり、model として置くのが自然である。

また、named model にすることで、以下が得られる。

- enum に安定した ID を与えられる
- field / param / returns から既存 TypeRef で参照できる
- MCP schema 生成時に同じ enum を再利用できる
- LLM に対して語彙の意味を `note` で補足できる
- inline enum の構文・互換性設計を増やさずに済む

### なぜ `str` と暗黙互換にしないか

enum は underlying JSON 表現としては string だが、設計契約としては string より狭い型である。

`enum -> str` または `str -> enum` を暗黙互換にすると、flow wiring 上は自由文字列と閉じた語彙が混ざり、enum を導入する価値が弱くなる。

brewprint はプログラミング言語の型システムを再現しないが、設計契約としての nominal identity は守る。
必要な変換は adapter / normalize task として明示する。

### なぜ inline enum TypeRef を導入しないか

inline enum は一見便利だが、v1.1 の最小拡張としては重い。

導入すると以下の設計論点が発生する。

- 同じ値集合なら互換とするのか
- 値の順序違いをどう扱うか
- anonymous enum に note を持たせるか
- TypeRef parser に新しい variant を足すか
- named enum と inline enum の互換性をどうするか

ADR-060 では TypeRef を primitive / named model / inline list / inline dict に絞った。
本ADRではその方針を維持し、enum は named model として導入する。

### なぜ値ごとの metadata を入れないか

MCP contract の現時点の要求は、許容値集合を machine-readable にすることである。
値ごとの説明は MCP spec 側の表や model の `note` で当面足りる。

`values` を object list にすると、最小導入の割に schema が重くなる。
値ごとの説明、deprecated、alias 等が必要になった場合に別ADRで拡張する。

### なぜ `invalid_enum_value` は初期導入しないか

brewprint の model field は schema を定義するものであり、通常の runtime 値を保持しない。
そのため、enum-typed field に具体値が入っていて、それが `values` に含まれるかを検査する場面はまだない。

初期導入では enum model 自体の定義検証に集中し、literal value validation は将来の example/default/payload 検証と同時に扱う。

## 却下した代替案

### 代替案A: `str + note` を継続する

- 利点: 仕様・実装変更が不要
- 欠点: 値集合が machine-readable にならない。validator / MCP schema / LLM が閉じた語彙として扱えない。UC-002 self-hosting の public contract 表現として弱い

→ 却下。

### 代替案B: TypeRef に inline enum variant を追加する

```text
enum<error|warning|info|hint>
```

- 利点: 小さな field には手軽
- 欠点: TypeRef parser / compatibility に新 variant が必要。anonymous enum の互換性・命名・note・再利用性が未整理になる。ADR-060 Phase A の TypeRef 最小構成を広げる

→ 却下。

### 代替案C: `model.kind: scalar` を復活させ、`base: str` + `values` を持たせる

```yaml
- id: severity
  type: model
  kind: scalar
  base: str
  values: [error, warning]
```

- 利点: enum を string の特殊化として表現できる
- 欠点: ADR-021 で廃止した scalar kind を復活させることになる。型エイリアスと enum 制約の責務が混ざる

→ 却下。enum は scalar alias ではなく finite vocabulary model として扱う。

### 代替案D: `values` を `fields` の属性として持たせる

```yaml
fields:
  - name: severity
    type: str
    values: [error, warning]
```

- 利点: field local な制約として書ける
- 欠点: 同じ enum を複数fieldで再利用できない。MCP schema 生成時に共通語彙として扱いづらい。TypeRef の named model 方針とずれる

→ 却下。

### 代替案E: enum と `str` を互換にする

- 利点: 既存 `str` model との migration が楽
- 欠点: enum の制約が flow wiring validation 上弱くなる。自由文字列を enum として渡せてしまう

→ 却下。移行のための緩和は `any` または明示 adapter で扱う。

## 影響

### spec への影響

- `docs/spec/nodes.md`
  - `model.kind` に `enum` を追加する
  - `kind: enum` の `values` フィールドを定義する
  - `kind: enum` は `fields` / `element` / `value` を持たないことを明記する

- `docs/spec/type-ref.md`
  - named model TypeRef が enum model を参照できることを明記する
  - enum model の TypeRef compatibility は nominal であることを明記する
  - enum と `str` の暗黙互換を行わないことを明記する

- `docs/spec/diagnostics.md`
  - `invalid_enum_model`
  - `duplicate_enum_value`

### 実装への影響

- raw YAML / semantic model に `kind: enum` と `values` を追加する
- model validation で `values` の存在・非空・string・空文字・重複を検査する
- TypeRef resolution では enum model を named model として解決する
- TypeRef compatibility では enum model を list/dict 以外の named model と同じ nominal 比較にする
- MCP schema 生成を行う場合、enum model を string enum として出力できるようにする

### 既存 UC への影響

- UC-001 には影響しない想定
- UC-002 の MCP public contract model で、`str + note` の enum 候補を段階的に enum model へ移行する
- `diagnostic.severity` / `object_ref.object` / `get_reference_tree.direction` は初期移行対象
- `analyze_impact_response.impacts` は `impact_entry` model 化が必要なため、後続 migration task として扱う

### M15 への影響

M15 Phase C では、本ADRを受けて enum model の spec / implementation / UC-002 migration を扱う。

ただし、具体的な作業項目、fixture migration の順序、`analyze_impact_response.impacts` の `impact_entry` model 化などは ADR 本文では管理しない。
これらは `docs/tasks/m15-data-layer-expressiveness.md` および UC-002 側の task file で追跡する。

### 他設計への影響

- discriminated object は本ADRでは扱わない
- optional / required 制約は本ADRでは扱わない
- recursive struct / union list / arbitrary JSON object は本ADRでは扱わない
- inline struct は本ADRでは扱わない
- TypeRef 深さ制限 / lint 方針は本ADRでは扱わない

## Evidence

- commit: 693e3c0
- impl commit: tbd
- 参考: JSON Schema enum, OpenAPI string enum, nominal typing 方針（ADR-060）

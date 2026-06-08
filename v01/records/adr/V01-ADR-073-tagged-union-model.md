# V01-ADR-073: tagged union model 導入

- **status**: accepted
- **date**: 2026-05-11

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

V01-WORK-DATA-010 では、UC-002 self-hosting で見つかった data layer 表現力不足のうち、tagged union / discriminator payload support を独立した successor work として扱う。

UC-002 Phase A の MCP公開contract YAML では、JSON / MCP contract 上は object の種別によって payload shape が変わる構造を、現状 `any` + `note` で表現している。

代表例は `analyze_impact.change` である。

`docs/spec/mcp/tools/analyze-impact.md` §3 では、`change` は `kind` を discriminator とする object として定義されている。

例:

```json
{ "kind": "rename", "new_id": "auth.model.user.email_address" }
{ "kind": "remove" }
{ "kind": "change_type", "new_type": "str" }
{ "kind": "add", "added_id": "auth.model.user.locale" }
```

`kind` の値によって、必要な payload field が変わる。

一方、UC-002 の現行 YAML では以下のように `any` + `note` に退避している。

```yaml
fields:
  - name: change
    type: any
    note: "必須。kindをdiscriminatorとする変更種別object。rename/remove/change_type/change_contract/change_transition_target/add。v1 modelではdiscriminated objectを表せないためany。"
```

この表現は人間には読めるが、validator / MCP schema / LLM にとって machine-readable な schema ではない。
そのため、以下ができない。

- `kind` の許可値集合を model として扱う
- `kind` ごとの payload field を機械的に列挙する
- MCP public contract から `any` を減らす
- LLM が `rename` では `new_id` が必要、`remove` では payload を持たない、と構造的に理解する
- 将来 MCP schema 生成や validation に反映する

V01-ADR-067 では enum model を proposed として起票し、閉じた値集合を `model.kind: enum` で表現する案を定めた。
しかし enum は「値集合」だけを表す。
`analyze_impact.change` の問題は、値集合に加えて「tag value と payload shape の対応」を表す必要がある点であり、enum だけでは解けない。

したがって、本ADRでは `model.kind: tagged_union` を追加し、discriminator field が明示された object variant を model として表現する。

本ADRは、V01-REQ-DATA-004 / V01-WORK-DATA-010 の successor work として tagged union model の最小 contract を accepted とする。

### 判断根拠としての UC-002 実例の扱い

本ADRで参照する UC-002 の YAML は、ADR起票時点の self-hosting 実例である。
これらの YAML は今後の V01-WORK-DATA-010 実装・spec反映・UC-002 migration によって更新されうるため、本文中の具体的な field 名や `any` の配置は恒久仕様ではない。

本ADRが判断根拠として扱うのは、個々の YAML shape そのものではなく、UC-002 において以下の性質を持つ payload が現れたという事実である。

- object 内に discriminator field が明示されている
- discriminator の値集合が有限である
- discriminator の値ごとに payload field が異なる
- 現行表現では `any + note` に閉じており、validator / MCP schema / LLM に machine-readable な制約として渡らない
- TypeRef variant を増やさず named model として表現できる

したがって、本文中の UC-002 例は実装固定の仕様例ではなく、tagged union model 導入判断の evidence として扱う。
具体的な fixture migration の範囲や順序は V01-WORK-DATA-010 または後続 UC-002 task file で管理する。

## 決定

### 1. `model.kind: tagged_union` を追加する

brewprint の model kind に `tagged_union` を追加する。

```yaml
nodes:
  - id: analyze_impact_change
    type: model
    kind: tagged_union
    discriminator: kind
    variants:
      - tag: rename
        fields:
          - name: new_id
            type: str
      - tag: remove
        fields: []
      - tag: change_type
        fields:
          - name: new_type
            type: str
```

`tagged_union` は、同一 object 内の discriminator field を見て variant を判定する object model である。

上記の例では、`discriminator: kind` により、object は常に `kind` field を持つ。
`kind` の値が `rename` なら `rename` variant、`remove` なら `remove` variant、`change_type` なら `change_type` variant として解釈する。

### 2. `discriminator` は必須とする

`kind: tagged_union` の model は `discriminator` を必須とする。

```yaml
kind: tagged_union
discriminator: kind
```

`discriminator` は、variant 判定に使う field name である。

制約:

- non-empty string である
- object 内の top-level field name を表す
- dot path は V01-WORK-DATA-010 minimum では許可しない
- external discriminator は扱わない

以下のような nested path / external path は V01-WORK-DATA-010 minimum では扱わない。

```yaml
discriminator: object.kind   # non-goal
```

### 3. `variants` は必須とする

`kind: tagged_union` の model は `variants` を必須とする。

`variants` の制約:

- non-empty list である
- 各 variant は `tag` を必須で持つ
- 各 variant の `tag` は string である
- 空文字 tag は不可
- 同一 tagged union model 内で tag は重複不可
- 各 variant は `fields` を持つ。payload field がない variant は `fields: []` と書く

```yaml
variants:
  - tag: remove
    fields: []
```

### 4. discriminator field は variant fields に重複して書かない

`discriminator` で指定された field は、tagged union object の variant 判定 field として暗黙に存在する。
variant payload の `fields` には discriminator field 自体を書かない。

つまり、以下のように書く。

```yaml
- id: analyze_impact_change
  type: model
  kind: tagged_union
  discriminator: kind
  variants:
    - tag: rename
      fields:
        - name: new_id
          type: str
```

これは、JSON shape としては以下を意味する。

```json
{ "kind": "rename", "new_id": "..." }
```

以下のように variant field 内で `kind` を再定義することは invalid とする。

```yaml
variants:
  - tag: rename
    fields:
      - name: kind
        type: str
      - name: new_id
        type: str
```

### 5. variant fields は struct field の最小サブセットを使う

`variants[].fields[]` は、struct model の `fields[]` と同じ field object を基礎にする。

ただし V01-WORK-DATA-010 minimum の tagged union payload では、以下の最小サブセットのみを許可する。

| field | 必須 | 内容 |
|---|---:|---|
| `name` | ✓ | payload field 名 |
| `type` | ✓ | TypeRef |
| `note` | 任意 | field の補足説明 |

`pk` / `fk` / `unique` は V01-WORK-DATA-010 minimum の tagged union payload では扱わない。
これらは ER / struct model 向けの意味が強く、variant payload に持ち込むと責務が混ざるためである。

variant field の `type` は通常の TypeRef として扱う。
primitive / named model / inline `list<T>` / inline `dict<T>` を指定できる。

### 6. 使用側は既存 TypeRef の named model として参照する

tagged union model を使う側は、既存 TypeRef の named model 参照として指定する。

```yaml
nodes:
  - id: analyze_impact_request
    type: model
    kind: struct
    fields:
      - name: selector
        type: object_selector
      - name: change
        type: analyze_impact_change
```

TypeRef に `union<...>` や `tagged_union<...>` の inline 構文は追加しない。
`analyze_impact_change` は named model TypeRef として解決される。

### 7. discriminator tag は string literal 列挙として扱う

`variants[].tag` の集合が discriminator の許可値集合を表す。

tagged union の discriminator tag を表すために、別途 enum model を必須にはしない。

```yaml
discriminator: kind
variants:
  - tag: rename
  - tag: remove
```

この場合、`kind` の許可値は `rename` / `remove` である。

V01-ADR-067 の enum model は、通常 field の閉じた値集合を表すための model kind である。
一方、tagged union の discriminator tag は tagged union model 自体の variant 定義に内包される。

将来、enum model と tagged union discriminator の整合性を明示したくなった場合は別ADRで扱う。

### 8. `tagged_union` の型互換性は named model として nominal に扱う

V01-ADR-060 の TypeRef compatibility において、tagged union model は list/dict 以外の named model と同じく nominal に扱う。

互換例:

```text
analyze_impact_change -> analyze_impact_change  OK
analyze_impact_change -> any                    OK
any -> analyze_impact_change                    OK
analyze_impact_change -> str                    NG
analyze_impact_change -> other_change           NG
```

variant set や field 構造が同じであっても、別 ID の tagged union model は互換とはみなさない。
structural typing は V01-ADR-060 の方針どおり導入しない。

### 9. untagged union / general oneOf は導入しない

本ADRで導入するのは、同一 object 内の discriminator field が明示された tagged union のみである。

以下は V01-WORK-DATA-010 minimum では扱わない。

- untagged union
- general `oneOf` / `anyOf`
- `str | int` のような scalar union
- `SourceLocation | ObjectRef` のような shape 推論 union
- external discriminator
- adjacent discriminator
- discriminator path
- variant narrowing / exhaustiveness checking

例として、MCP schema の `diagnostic.related` は `SourceLocation` または `ObjectRef` の配列であり、これは untagged union / oneOf 領域である。
本ADRの対象には含めない。

### 10. tagged union の contract semantics を定義する

`tagged_union` は runtime payload validator ではなく、brewprint model 上の contract schema である。
runtime payload validation を non-goal とすることは、tagged union model の contract semantics を未定義にすることを意味しない。

`discriminator` は、その object が variant 判定 field を持つことを表す。
`variants[].tag` の集合は、discriminator field の許可値集合を表す。
`variants[].fields` は、その tag のときに object が持つ payload field の schema を表す。

payload field がない variant は `fields: []` と書く。
`fields: []` は「payload field を持たない variant」を意味し、「payload shape 未定義」や「任意 payload 許可」を意味しない。
payload shape をまだ定義しない場合は、`fields: []` ではなく、本ADRの範囲外として `any + note` 継続を選ぶ。

variant の順序は YAML の `variants[]` 順を保持する。
variant field の順序は YAML の `fields[]` 順を保持する。
この順序は render / schema generation / MCP inspect の表示順として使ってよいが、型互換性の意味には使わない。

本ADRは、tagged union model から schema / render / MCP inspect が variant 構造を読み取れることを定義する。
一方で、実際の MCP request / response JSON payload を実行時に検査する責務は持たない。

tagged union model は schema generation の入力となる。
ただし、V01-WORK-DATA-010 minimum では具体的な JSON Schema / MCP schema 出力形式、additionalProperties の扱い、required 配列の生成規則までは定義しない。
これらは schema generation 実装または別ADRで扱う。

### 11. runtime payload validation は V01-WORK-DATA-010 minimum では non-goal とする

本ADRで定義する validation は、tagged union model 定義そのものの検証である。

以下は runtime payload validation または将来の schema generation policy の範囲であり、V01-WORK-DATA-010 minimum では扱わない。

- 実際の MCP request / response JSON payload を tagged union model に照合する runtime validation
- 実際の JSON payload に discriminator field が存在するかの実行時検査
- 実際の JSON payload の discriminator value が `variants[].tag` に含まれるかの実行時検査
- variant payload field が実際の JSON payload に存在するかの実行時検査
- unknown additional payload field の許可 / 不許可
- optional / required の完全表現
- `new_to` / `new_action` の少なくとも一方が必要、のような cross-field constraint
- variant ごとの payload value literal validation

unknown additional payload field の許可 / 不許可は runtime payload validation および将来の schema generation policy の範囲とし、本ADRでは決めない。

`analyze_impact` tool の runtime request validation は、MCP tool 実装側の責務として維持する。
brewprint model はその public contract を machine-readable に近づけるが、MCP server の実行時 validator を置き換えない。

### 12. tagged union validation diagnostic を追加する

新しい diagnostic code として以下を追加する。

| code | severity | 意味 |
|---|---|---|
| `invalid_tagged_union_model` | error | `kind: tagged_union` の定義が不正。`discriminator` 欠落、`variants` 欠落、空 variants など |
| `duplicate_variant_tag` | error | 同一 tagged union model 内で `variants[].tag` が重複している |
| `invalid_variant_field` | error | variant field が不正。discriminator field の再定義、許可されない field 属性など |

variant field の TypeRef 検証には既存の `invalid_type_ref` / `unresolved_field_type` を使う。
variant field 名の重複には既存の `duplicate_model_field` を流用してよい。

## Acceptance review result

V01-TASK-DATA-010-01 reviewed this ADR and classified it as `revise-before-acceptance`.

The core tagged union contract is accepted under V01-REQ-DATA-004 / V01-WORK-DATA-010, with the following boundaries:

- This ADR owns schema-level tagged union model expressiveness.
- Runtime MCP request / response payload validation remains out of scope.
- Untagged union / general oneOf / scalar union / external discriminator remain out of scope.
- DAG asset TypeRef hint, MCP identity, UC-002 duplicate task QID repair, and broad notes-retreat cleanup remain outside this ADR.
- Tagged union render support is not assumed to be already implemented by the V01-WORK-DATA-003 model-file render minimum.
- V01-WORK-DATA-010 owns the follow-up spec / implementation / fixture sequence.

## 理由

### なぜ tagged union を V01-WORK-DATA-010 minimum に入れるか

UC-002 は brewprint 自身の MCP public contract を blueprint 化する self-hosting 実例である。
この contract には、`kind` の値に応じて payload shape が変わる object が存在する。

これを `any + note` に留めると、人間向け説明としては成立するが、machine-readable な schema としては失われる。
MCP schema 生成、validator、LLM の設計理解にとって、tagged union は有効な拡張である。

特に `analyze_impact.change` は、discriminator field が明示され、variant set も有限で、payload 差分も比較的小さい。
初期導入の evidence として適切である。

### なぜ `model.kind: tagged_union` か

既存の brewprint data layer では、型として意味を持つものは model に置く。
`struct` / `list` / `dict` / `enum` と同じく、tagged union も再利用可能な型定義であり、model として置くのが自然である。

また、named model にすることで、以下が得られる。

- tagged union に stable ID を与えられる
- field / param / returns から既存 TypeRef で参照できる
- MCP schema 生成時に同じ union shape を再利用できる
- TypeRef に新 variant を増やさずに済む
- LLM に対して union の意味を `note` で補足できる

### なぜ discriminator を必須にするか

本ADRの目的は、判定 field が明示された union を安全に表現することである。

discriminator がない union は、shape 推論 / oneOf / optional field ambiguity の問題を持つ。
これは tagged union とは別機能であり、V01-WORK-DATA-010 minimum としては重い。

`discriminator` を必須にすることで、validator / schema generator / LLM が variant 判定方法を一意に理解できる。

### なぜ external discriminator を扱わないか

`get_signature_response.signature` や `inspect_response.signature` は、対象 object kind によって shape が変わる。
しかし、その discriminator は `signature` object 内ではなく、外側の `object.kind` に存在する。

```json
{
  "object": { "kind": "task" },
  "signature": { ... }
}
```

これは external discriminator / adjacent discriminator の領域であり、今回の `discriminator: <field-name>` 方式とは異なる。
ここまで扱うと、MCP response envelope 全体の schema composition を設計する必要があり、ADR のスコープが大きくなりすぎる。

したがって、本ADRでは同一 object 内に discriminator field がある tagged union に限定する。

### なぜ enum model と自動連携しないか

tagged union の discriminator tag 集合は、`variants[].tag` から自然に導出できる。
ここに enum model 参照を必須にすると、同じ値集合を enum model と variants の両方に書くことになり、同期問題が生じる。

V01-ADR-067 の enum model は、通常 field の値集合を表すためのものとして扱い、tagged union discriminator の tag 集合は tagged union model の内部定義として持つ方が単純である。

将来、MCP schema 出力時に discriminator property の enum を明示したい場合も、`variants[].tag` から生成できる。

### なぜ runtime payload validation を non-goal にするか

brewprint model は public contract / design schema を表すが、MCP server の runtime validator そのものではない。

実際の JSON request payload が `analyze_impact_change` に一致するかを検査するには、実行時入力、optional/required、unknown field、cross-field constraint などの別領域が必要になる。
これを本ADRに含めると、tagged union model 導入の範囲を超える。

V01-WORK-DATA-010 minimum では、model 定義を machine-readable にすることに集中する。

### なぜ variant fields を struct field の完全コピーにしないか

variant payload の field は struct field と似ているが、ER / DB 的な意味を持つ `pk` / `fk` / `unique` は不要である。

payload object は API / MCP contract の variant shape であり、ER entity ではない。
そのため、V01-WORK-DATA-010 minimum では `name` / `type` / `note` の最小サブセットに絞る。

## 却下した代替案

### 代替案A: `any + note` を継続する

- 利点: 仕様・実装変更が不要
- 欠点: `kind` ごとの payload shape が machine-readable にならない。UC-002 self-hosting の public contract 表現として弱い

→ proposed 時点では保留せず、導入案を起票する。ただし実装コスト見合いの確認後、accepted に進めない選択肢は残す。

### 代替案B: TypeRef に inline union variant を追加する

```text
union<rename_change | remove_change>
```

- 利点: TypeRef 上で union を直接表現できる
- 欠点: TypeRef parser / compatibility に新 variant が必要。V01-ADR-060 の TypeRef 最小構成を広げる。tagged / untagged / oneOf の境界が曖昧になる

→ 却下。tagged union は named model として定義し、使用側は既存 TypeRef の named model 参照に留める。

### 代替案C: enum model + struct model 群で表現する

```yaml
- id: change_kind
  kind: enum
  values: [rename, remove]

- id: rename_change
  kind: struct
  fields: ...
```

- 利点: 既存 struct / enum の組み合わせで表現できる
- 欠点: `kind` と payload shape の対応を1つの model として表せない。使用側 field が結局 `any` または別の union 表現を必要とする

→ 却下。tag value と payload shape の対応を model として表す必要がある。

### 代替案D: untagged union / general oneOf も同時に導入する

- 利点: `diagnostic.related` なども表現できる
- 欠点: discriminator がないため shape 推論が必要になる。optional field と相性が悪く、曖昧性の扱いが重い。これだけで独立した機能になる

→ 却下。V01-WORK-DATA-010 minimum では tagged union のみに限定する。

### 代替案E: external discriminator も扱う

- 利点: `get_signature_response.signature` / `inspect_response.signature` も表現できる
- 欠点: discriminator が同一 object 内にないため、envelope と payload の相関を表す別仕様が必要になる。MCP response schema 全体の composition 問題になり、初期導入として重い

→ 却下。V01-WORK-DATA-010 minimum では同一 object 内 discriminator に限定する。

### 代替案F: discriminator field を variant fields に明示的に書かせる

```yaml
variants:
  - tag: rename
    fields:
      - name: kind
        type: str
      - name: new_id
        type: str
```

- 利点: JSON上の field がすべて fields に見える
- 欠点: `tag` と `kind` field の値が二重定義になる。enum / literal value 制約も別途必要になる

→ 却下。discriminator field は tagged union model が暗黙に持つものとし、variant fields は payload 差分だけを書く。

## 影響

### spec への影響

本ADR受理後、以下を更新する。

- `docs/spec/nodes.md`
  - `model.kind` に `tagged_union` を追加する
  - `kind: tagged_union` の `discriminator` / `variants` / `variants[].tag` / `variants[].fields` を定義する
  - `fields: []` は payload field を持たない variant を意味し、任意 payload 許可や shape 未定義を意味しないことを明記する
  - variant / variant field の順序保持を render / schema generation / MCP inspect の表示順として明記する
  - variant fields で許可する field 属性を `name` / `type` / `note` に限定する

- `docs/spec/type-ref.md`
  - named model TypeRef が tagged union model を参照できることを明記する
  - tagged union model の TypeRef compatibility は nominal であることを明記する
  - TypeRef に inline union variant を追加しないことを明記する

- `docs/spec/diagnostics.md`
  - `invalid_tagged_union_model`
  - `duplicate_variant_tag`
  - `invalid_variant_field`

### 実装への影響

- raw YAML / semantic model に `kind: tagged_union`、`discriminator`、`variants` を追加する
- model validation で tagged union 定義の妥当性を検査する
- variant field の TypeRef を通常 field と同様に解決する
- TypeRef resolution では tagged union model を named model として解決する
- TypeRef compatibility では tagged union model を list/dict 以外の named model と同じ nominal 比較にする
- MCP schema 生成を行う場合、tagged union model を discriminator 付き schema として出力できるようにする
- schema generation の具体出力形式、additionalProperties、required 配列の生成規則は本ADRでは決めず、schema generation 実装または別ADRで扱う

### UC-002 への影響

- V01-WORK-DATA-010 の初期 fixture / migration candidate は `analyze_impact_change` とする
- `analyze_impact_request.change` は `any` から `analyze_impact_change` へ移行候補になる
- `analyze_impact_response.change` は input change を返すため同じ `analyze_impact_change` へ移行候補になる
- `suggested_fixes[]` は fix kind 依存 payload を持つ可能性があるが、現状 `impacts: any` の内部に埋もれているため、初期 migration candidate には含めない
- `get_signature_response.signature` / `inspect_response.signature` / `inspect_response.members` は broader kind-specific payload candidate として残すが、external discriminator / envelope correlation を含むため本ADRの初期 migration candidate には含めない
- `diagnostic.related` は untagged union / oneOf 領域であり、本ADRの対象には含めない

### render / catalog への影響

tagged union は新しい model kind であるため、model を表示する render / catalog 仕様にも反映が必要である。

The existing V01-WORK-DATA-003 model-file render minimum is already implemented for `struct` / `enum` / `list` / `dict` and same-file private helper models.

Tagged union discriminator / variants rendering is not covered by that minimum. It remains a follow-up under V01-WORK-DATA-010 or a later render-specific task.

- model file render follow-up should display `tagged_union` discriminator / variants when tagged union render support is implemented
- model catalog is expected to list / filter / summarize `tagged_union` as a model kind
  - V01-ADR-072 already includes `tagged_union_models` as an include filter
  - public tagged union model can appear in a `## Tagged unions` section
  - file-private tagged union model can appear under private helper model display when private helpers are included
- DAG asset node TypeRef hint remains outside this ADR and is owned by V01-ADR-074 / REQ-DATA-005-style successor work

Model catalog concrete kind display / shape display is handled by V01-ADR-072 and later spec / implementation work.

### Successor work boundary

This ADR is accepted under V01-REQ-DATA-004 / V01-WORK-DATA-010.

It does not reopen M15, V01-WORK-DATA-001, V01-WORK-DATA-002, V01-WORK-DATA-003, or V01-WORK-DATA-004.

V01-WORK-DATA-010 owns the follow-up sequence for spec alignment, implementation, fixtures, verification, and close.

### 他設計への影響

- enum model は V01-ADR-067 で別途扱う
- TypeRef container complexity / anonymous inline struct 不採用は V01-ADR-069 で別途扱う
- optional / required 制約は本ADRでは扱わない
- recursive struct / union list / arbitrary JSON object は本ADRでは扱わない
- file-private helper model は本ADRでは扱わない
- untagged union / oneOf は本ADRでは扱わない

## Evidence

- original proposal commit: 5ae7769
- acceptance review: V01-TASK-DATA-010-01
- successor requirement: V01-REQ-DATA-004
- successor work item: V01-WORK-DATA-010
- implementation commit: tbd
- close boundary: accepted as V01-WORK-DATA-010 successor scope; implementation is not included in this ADR update.
- 参考: UC-002 MCP公開contract YAML における `analyze_impact.change` の `any + note` 暫定表現、JSON Schema / OpenAPI discriminator、TypeScript discriminated union

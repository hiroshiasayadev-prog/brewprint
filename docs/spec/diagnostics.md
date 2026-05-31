---
scope: docs/spec/diagnostics.md
status: draft
last_updated: 2026-05-31
summary: >
  brewprint validation diagnosticsの外部向け仕様。
  severity / code / file / message の意味、diagnostic file表現、現在のdiagnostic code一覧を定義する。
depends_on:
  - docs/adr/047-go-semantic-model-query-layer-boundary.md
  - docs/adr/051-unsupported-yaml-file-warning.md
  - docs/adr/052-source-file-path-normalization.md
  - docs/adr/058-subnode-file-private-scope-enforcement.md
  - docs/adr/060-flow-wiring-type-compatibility.md
  - docs/adr/061-foreach-returns-collected-asset.md
  - docs/adr/062-task-return-source.md
  - docs/adr/063-task-return-source-initialized-store.md
  - docs/adr/067-enum-model.md
  - docs/adr/069-type-ref-container-complexity-lint.md
  - docs/adr/070-model-visibility-file-private-helper-model.md
---

# Diagnostics仕様

## 目的

brewprintは `resolve.Build` 時にsemantic validation diagnosticsを生成する。

Diagnosticsは以下で利用する。

- `brewprint validate --yaml-root <path>` のtext出力
- `brewprint validate --yaml-root <path> --format json` のJSON出力
- `brewprint mcp --yaml-root <path>` 起動前のsemantic error検出

## Diagnostic object

JSON出力では、diagnosticは以下の形で表現される。

```json
{
  "severity": "error",
  "code": "unresolved_model",
  "file": "order/state.yaml",
  "message": "unresolved event payload model: payment_event"
}
```

| field | required | description |
|---|---:|---|
| `severity` | yes | `error` または `warning` |
| `code` | no | machine-readableなdiagnostic code。未分類の場合は省略されうる |
| `file` | no | diagnostic対象のYAML FileID |
| `message` | yes | 人間向け説明 |

## File representation

Diagnostic の `file` は FileID を使う。

`yaml/` 配下のファイルでは、プロジェクトルート相対pathではなく、`yaml/` prefix を除いたIDを出力する。

```text
rawyaml.File.Path: yaml/order/state.yaml
rawyaml.File.ID:   order/state.yaml
diagnostic.file:   order/state.yaml
```

`render_index.yaml` のように `yaml/` 外のファイルでは、FileID と Path は同じ値になる。

```text
rawyaml.File.Path: render_index.yaml
rawyaml.File.ID:   render_index.yaml
diagnostic.file:   render_index.yaml
```

この規約により、diagnostic出力・MCP response・renderer内部のFileID表現を揃える。

> 由来: ADR-052 §決定

## Severity

| severity | meaning |
|---|---|
| `error` | validation失敗。`brewprint validate` はnon-zero errorを返す |
| `warning` | validationは成功扱い。diagnosticは表示する |

## Text format

`brewprint validate --format text` は以下の形式で1行ずつ出力する。

```text
<severity> <code> <file>: <message>
```

例:

```text
error unresolved_model order/state.yaml: unresolved event payload model: payment_event
```

Diagnosticsが空の場合は以下を出力する。

```text
ok
```

## JSON format

`brewprint validate --format json` は以下の形式で出力する。

```json
{
  "diagnostics": [],
  "error_count": 0,
  "warning_count": 0
}
```

Error diagnosticsが存在する場合でも、JSONはstdoutへ出力される。その後、commandはnon-zero errorを返す。

## Diagnostic ordering

Diagnosticsは `resolve.Build` の返却前に安定ソートされる。

Sort key:

1. severity rank
   - `error`
   - `warning`
   - unknown severity
2. file id
3. code
4. message

## Diagnostic codes

### Semantic reference / model validation

| code | severity | description |
|---|---|---|
| `invalid_model_id` | error | model idがprimitive予約語を使っている |
| `unresolved_model` | error | model参照が解決できない |
| `unresolved_field_type` | error | model field typeがprimitiveまたは定義済みmodelとして解決できない |
| `unresolved_fk` | error | FK参照先fieldが解決できない |
| `unresolved_store` | error | store参照が解決できない |
| `invalid_endpoint` | error | endpoint taskのmethod/path定義が不正 |
| `invalid_store_kind` | error | store kindが許可値ではない |
| `invalid_model_kind` | error | model kindが許可値ではない |
| `duplicate_model_field` | error | model field名が重複している |
| `duplicate_model_id` | error | 同一 file 内の model id 重複、または同一 module 内の public model と file-private helper model の同名衝突 |
| `duplicate_primary_key` | error | model内にprimary keyが複数ある |
| `missing_required_field` | error | 必須fieldが欠落している |
| `invalid_type_ref` | error | TypeRef構文が不正、TypeRef として扱えない container kind 等を指定している、または parser safety limit を超過している |
| `invalid_enum_model` | error | `kind: enum` の定義が不正。`values` 欠落、空、非 string 値、空文字、または enum で禁止される `fields` / `element` / `value` 指定など |
| `duplicate_enum_value` | error | 同一 enum model 内で `values` が重複している |

`invalid_enum_model` は `kind: enum` の `values` が欠落している、空 list である、string list ではない、空文字を含む、または enum model では禁止される `fields` / `element` / `value` を持つ場合に出す。

`duplicate_enum_value` は同一 enum model 内の `values` に重複がある場合に出す。

`invalid_enum_value` は v1.1 minimum では追加しない。現時点の brewprint YAML は主に schema / model 定義であり、enum-typed field に対する runtime literal 値を保持しないためである。

`duplicate_model_id` は task-file helper minimum の model identity validation に使う。名前衝突 rule は [naming.md](./naming.md#41-task-file-helper-model-の名前衝突) §4.1、task-file helper model の基本 semantics は [nodes.md](./nodes.md#task-file-private-helper-model-semantics) を正とする。

Task-file helper model が参照元の TypeRef scope に入らない場合、初期実装では、出現箇所に応じて既存の `unresolved_model` または `unresolved_field_type` を出してよい。専用 code が必要な場合は、TASK-DATA-002-03 で `invalid_private_model_reference` 等の追加可否を判断する。

> 由来: ADR-067 §6, ADR-070 §7〜§8

### TypeRef lint / warning

| code | severity | description |
|---|---|---|
| `opaque_type_ref` | warning | container TypeRef の内部に `any` が含まれ、shape の意味が named model に回収されていない |

`opaque_type_ref` は validation 成功扱いとする。message では、必要に応じて named model への切り出しを提案する。

`opaque_type_ref` は container TypeRef 内の `any` を対象とする warning であり、bare `any` field や `any + note` の主要 response shape 全体を警告・解消するものではない。

`unclear_dict_key` / `deep_type_ref` は将来 lint 候補であり、v1.1 minimum の diagnostic code としては追加しない。

> 由来: ADR-069 §4〜§10

### Duplicate / symbol validation

| code | severity | description |
|---|---|---|
| `duplicate_node` | error | public node QualifiedID が project-wide に重複している |
| `duplicate_main_node` | error | 1 file内にmain nodeが複数ある |
| `duplicate_sub_node` | error | file-private sub node の local ID が同一 file 内で重複している |
| `duplicate_actor` | error | actor idが重複している |
| `duplicate_initialized_store` | error | initialized storeが同一file内で重複している |

`duplicate_node` は public QualifiedID を持つ main node の衝突に限定する。file-private sub node は public QualifiedID を持たないため、別 file 間で同じ local ID を持っていても `duplicate_node` にはしない。

同一 file 内で sub node local ID が重複している場合は `duplicate_sub_node` を出す。

### Flow validation

| code | severity | description |
|---|---|---|
| `unsupported_flow_entry` | warning | 空または未対応のflow entry |
| `unresolved_flow_task` | error | flow step / foreach taskが解決できない |
| `unresolved_flow_node` | error | branch / fork / join nodeが解決できない |
| `invalid_flow_branch` | error | fork branch内のstep定義が不正 |
| `unmatched_join_param` | error | join paramに対応するbranch returnがない |
| `incompatible_wiring_type` | error | flow wiring source の TypeRef と target param の TypeRef が互換しない |
| `invalid_wiring_source` | error | source は解決できたが、その文脈では wiring source として使えない |
| `unresolved_wiring_source` | error | wiring source が node id / `$params.<name>` / `$item` / collected asset source / initialized source のいずれとしても解決できない |
| `invalid_foreach_over_type` | error | `foreach.over` が指す source が list として扱えない |
| `invalid_foreach_returns` | error | apply 先 task に `returns` がないのに `foreach.returns` が指定されている、または当該 foreach 自身の `params` 内から自分の `returns` 名を参照している |
| `unresolved_return_source` | error | `returns.source` が node id / `$params.<name>` / collected asset source / initialized source のいずれとしても解決できない |
| `invalid_return_source` | error | source は解決できたが task return source として使えない |
| `incompatible_return_type` | error | `returns.source` の TypeRef と `returns.model` の TypeRef が互換しない |
| `duplicate_flow_source` | error | 同一 flow file 内で node id / collected asset source 名 / initialized source 名のいずれかが他のものと重複している |

`invalid_type_ref` の message には、不正な TypeRef 文字列とその出現位置を含める。parser safety limit 超過の場合も、対象 TypeRef と limit 超過が分かる message にする。

TypeRef 構文は valid だが内部の named model が解決できない場合は、出現箇所に応じて既存の未解決 diagnostic を使う。`fields[].type` では `unresolved_field_type`、`params[].model` / `returns.model` / `model.element` / `model.value` では `unresolved_model` を使う。構文自体が壊れている場合、TypeRef として扱えない container kind を指定した場合、または parser safety limit を超過した場合のみ `invalid_type_ref` を出す。

`incompatible_wiring_type` の message には、source TypeRef / target TypeRef / wiring位置を含める。

`duplicate_flow_source` は、同一 flow file 内で node id / `foreach.returns` で宣言された collected asset source 名 / `initializes[].name` で宣言された initialized source 名のいずれかが他のものと重複している場合に出す。task の `returns.name` は通常 flow の wiring source ではないため、task `returns.name` と他の bare source 名が同名でも `duplicate_flow_source` にはしない。

`invalid_foreach_returns` は、以下の場合に出す。

- apply 先 task に `returns` がないにもかかわらず `foreach.returns` が指定されている
- 当該 foreach 自身の `params` 内から自分自身の `returns` 名を参照している

`unresolved_wiring_source` と `invalid_wiring_source` は区別する。

- `unresolved_wiring_source`: typo などにより参照先が存在しない。wiring source が node id / `$params.<name>` / `$item` / collected asset source / initialized source のいずれとしても解決できない
- `invalid_wiring_source`: 参照先は存在するが、その文脈では source として使えない。例: returns を持たない node、foreach 外の `$item`

initialized source は valid な wiring source 種別であり、`invalid_wiring_source` の対象にはならない。`initializes[].model` が解決不能な場合、initialized source を参照する wiring の `incompatible_wiring_type` は抑制する。

TypeRef の解決失敗、未解決参照、`invalid_foreach_over_type` が発生している `$item` wiring、または TypeRef 解決不能な collected asset source / initialized source を参照する wiring では、重複して `incompatible_wiring_type` を発行しない。

`unresolved_return_source` と `invalid_return_source` は区別する。

- `unresolved_return_source`: typo などにより参照先が存在しない。`returns.source` が node id / `$params.<name>` / collected asset source / initialized source のいずれとしても解決できない
- `invalid_return_source`: 参照先は存在するが、task return source として使えない。例: returns を持たない node、`$item`

initialized source は valid な return source 種別であり、`invalid_return_source` の対象にはならない。

`incompatible_return_type` の message には、source TypeRef / target TypeRef / `returns.source` 位置を含める。

`returns.source` の TypeRef または `returns.model` の TypeRef が解決不能な場合、重複して `incompatible_return_type` を発行しない。

> 由来: ADR-060 §6, §7; ADR-061 §9; ADR-062 §8; ADR-063 §9

### Transition validation

| code | severity | description |
|---|---|---|
| `unresolved_transition_state` | error | transition from/to stateが解決できない |
| `unresolved_transition_event` | error | transition eventが解決できない |
| `duplicate_transition` | error | transition from/on/guardが重複している |
| `missing_transition_guard` | error | branched transitionでguardが欠落している |

### View / scenario validation

| code | severity | description |
|---|---|---|
| `duplicate_view` | error | API View / ER View / Sequence Scenarioのidが重複している |
| `invalid_view_definition` | error | view/scenarioの必須定義が欠落または不正 |
| `duplicate_view_module` | error | API View / ER View内のmodule定義が重複している |
| `unresolved_sequence_step` | error | sequence scenario stepがtransitionへ解決できない |
| `non_continuous_sequence` | error | sequence scenario stepが前stepのto stateから連続していない |

### File classification validation

| code | severity | description |
|---|---|---|
| `unsupported_file` | warning | YAMLファイルが種別判定不能（`as:` も `nodes:` も持たず、`render_index.yaml` でもない）。`as:` の書き忘れ等を検出する。詳細は [file-types.md](./file-types.md) §4 |

> 由来: ADR-051

### Generic fallback

| code | severity | description |
|---|---|---|
| `semantic_validation` | error/warning | まだ具体codeへ分類されていないsemantic validation diagnostic |

## Compatibility note

Diagnostic `message` は人間向けであり、将来文言が変わる可能性がある。
外部ツールは `severity` / `code` / `file` を優先して扱う。

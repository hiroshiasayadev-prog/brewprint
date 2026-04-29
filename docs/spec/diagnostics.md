---
scope: docs/spec/diagnostics.md
status: draft
last_updated: 2026-04-29
summary: >
  brewprint validation diagnosticsの外部向け仕様。
  severity / code / file / message の意味と、現在のdiagnostic code一覧を定義する。
depends_on:
  - docs/adr/047-go-semantic-model-query-layer-boundary.md
  - docs/adr/051-unsupported-yaml-file-warning.md
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
| `file` | no | diagnostic対象のYAML file id |
| `message` | yes | 人間向け説明 |

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
| `duplicate_primary_key` | error | model内にprimary keyが複数ある |
| `missing_required_field` | error | 必須fieldが欠落している |

### Duplicate / symbol validation

| code | severity | description |
|---|---|---|
| `duplicate_node` | error | node qualified idが重複している |
| `duplicate_main_node` | error | 1 file内にmain nodeが複数ある |
| `duplicate_actor` | error | actor idが重複している |
| `duplicate_initialized_store` | error | initialized storeが同一file内で重複している |

### Flow validation

| code | severity | description |
|---|---|---|
| `unsupported_flow_entry` | warning | 空または未対応のflow entry |
| `unresolved_flow_task` | error | flow step / foreach taskが解決できない |
| `unresolved_flow_node` | error | branch / fork / join nodeが解決できない |
| `invalid_flow_branch` | error | fork branch内のstep定義が不正 |
| `unmatched_join_param` | error | join paramに対応するbranch returnがない |

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

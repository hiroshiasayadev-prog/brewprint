---
scope: docs/spec/mcp/tools/analyze-impact.md
status: draft
last_updated: 2026-06-07
summary: >
  analyze_impact toolの仕様を定義する。
  対象objectとchange kindから影響範囲を解釈付きで返す。
  severity / fixability / coverage / suggested_fixes を規定する。
depends_on:
  - docs/adr/049-mcp-query-reference-vocabulary.md
  - docs/adr/054-mcp-query-coverage-for-design-conversation.md
  - docs/adr/055-mcp-reference-tree-traversal.md
  - docs/adr/056-mcp-analyze-impact-tool-design.md
---

# `analyze_impact`

## 1. Purpose

`analyze_impact` は、対象 object と変更種別（`change`）を受け取り、 影響範囲を **意味づけ済みの impact list** として返す。

`get_references` / `get_reference_tree` が返す raw な reference 情報の上に、 change kind ごとの解釈、severity 判定、 修正可能性、推奨アクションを載せる。

返すもの:

- `summary` — severity / kind / fixability ごとの件数集計
- `impacts[]` — 各影響箇所の意味づけ済み entry
- `coverage` — 何を分析したか / 何を分析していないか
- `assumptions` — tool 側の前提・限界
- `truncated` — 返却上限による打ち切り情報
- `diagnostics` — diagnostic list

返さないもの:

- 完全な YAML snippet（`get_source` の責務）
- raw な reference graph（`get_reference_tree` の責務）
- renderer 出力 md 内の presentation 詳細
- semantic contract の意味互換性判断

> 由来: ADR-056 §決定 §1

---

## 2. Input

```json
{
  "selector": {
    "object": "field",
    "id": "auth.model.user.email"
  },
  "change": {
    "kind": "rename",
    "new_id": "auth.model.user.email_address"
  },
  "scope_modules": ["auth"],
  "max_impacts": 200
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `selector` | ✓ | 対象 object selector |
| `change` | ✓ | 変更種別。discriminated object（§3 参照） |
| `scope_modules` | 任意 | 分析範囲を絞る module list |
| `max_impacts` | 任意 | impact 返却上限。省略時 `200` |

`depth` は input に持たない。
影響範囲の探索深さは change kind ごとに tool 側が決める。

`selector` の object / kind 対応範囲は `docs/spec/mcp/schema.md` の selector support matrix を正本とする。
`analyze_impact` で matrix が `no` の selector を受け取った場合は、tool error ではなく、空 `impacts`、`coverage`、および `unsupported_selector` diagnostic を含む通常responseとして返す。

> 由来: ADR-056 §決定 §2

---

## 3. `change` discriminated object

`change` は `kind` を discriminator とする object とする。
v1 で扱う `kind` は以下。

```jsonc
// rename
{ "kind": "rename", "new_id": "..." }

// remove
{ "kind": "remove" }

// change_type （field の type 変更）
{ "kind": "change_type", "new_type": "auth.model.account" }

// change_contract （task の params / returns 等の構造変更）
{ "kind": "change_contract", "note": "params に session_id を追加" }

// change_transition_target （transition の to / action 変更）
{
  "kind": "change_transition_target",
  "new_to": "order.state.shipped",
  "new_action": "order.task.notify_user"
}

// add （新規追加。既存影響というより整合性分析）
{ "kind": "add", "added_id": "auth.model.user.locale" }
```

`add` を投げた場合、 `coverage.analyzed` は他 kind と異なる（§7 参照）。

### validation

`change.kind` ごとの必須payloadは以下とする。

| kind | 必須payload | optional payload | validation |
|---|---|---|---|
| `rename` | `new_id` | - | `new_id` が空なら `invalid_change_payload` |
| `remove` | - | - | 追加payloadは無視せず `invalid_change_payload` |
| `change_type` | `new_type` | - | `new_type` が空なら `invalid_change_payload` |
| `change_contract` | - | `note` | payloadなしでも有効 |
| `change_transition_target` | - | `new_to`, `new_action` | `new_to` / `new_action` の少なくとも一方が必要 |
| `add` | `added_id` | - | `added_id` が空なら `invalid_change_payload` |

kind と payload の不正な組み合わせは tool error として `invalid_change_payload` を返す。
unsupported selector は tool error ではなく、空 `impacts`、`unsupported_selector` diagnostic、coverage を含む通常responseとして返す。

> 由来: ADR-056 §決定 §2

---

## 4. Output

```json
{
  "target": {
    "object": "field",
    "id": "auth.model.user.email"
  },
  "change": {
    "kind": "rename",
    "new_id": "auth.model.user.email_address"
  },
  "summary": {
    "by_severity": {
      "breaking": 3,
      "warning": 5,
      "info": 2
    },
    "by_fixability": {
      "mechanical": 3,
      "suggested": 2,
      "manual_review": 3,
      "unknown": 2
    },
    "by_kind": {
      "field_consumer": 4,
      "transition_action": 2,
      "render_output": 1
    }
  },
  "impacts": [
    {
      "id": "impact-001",
      "kind": "field_consumer",
      "severity": "breaking",
      "fixability": "mechanical",
      "object": {
        "object": "node",
        "kind": "task",
        "id": "auth.task.login"
      },
      "reason": "auth.task.login が field 'email' を reads しているため、rename後に参照解決できなくなる",
      "via": ["reads", "model_field"],
      "source": {
        "file": "auth/task/login.yaml",
        "line": 42,
        "column": 7,
        "end_line": 42,
        "end_column": 18
      },
      "recommended_action": "reads field reference を email_address へ更新する",
      "suggested_fixes": [
        {
          "kind": "replace_reference",
          "confidence": "high",
          "from": "email",
          "to": "email_address",
          "source": {
            "file": "auth/task/login.yaml",
            "line": 42,
            "column": 7
          }
        }
      ]
    }
  ],
  "coverage": {
    "analyzed": [
      "direct_references",
      "reference_tree",
      "model_field_resolution",
      "transition_action_resolution",
      "type_signature_identity",
      "render_output_files"
    ],
    "not_analyzed": [
      "type_structural_compatibility",
      "semantic_contract_compatibility",
      "render_presentation_details",
      "wireframe_element_binding"
    ],
    "note": "v1では reference 経路の到達可能性と型 signature identity のみを判定する。"
  },
  "assumptions": [
    "rename後のID衝突は検証対象外",
    "note内の自然言語参照は解析対象外"
  ],
  "truncated": false,
  "truncated_reasons": [],
  "diagnostics": []
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `target` | ✓ | 分析対象 ObjectRef |
| `change` | ✓ | input で指定された change（そのまま返す） |
| `summary` | ✓ | severity / fixability / kind ごとの件数集計 |
| `impacts` | ✓ | 影響箇所一覧 |
| `coverage` | ✓ | 分析範囲の明示 |
| `assumptions` | ✓ | tool の前提・限界 |
| `truncated` | ✓ | `max_impacts` により打ち切ったか |
| `truncated_reasons` | ✓ | 打ち切り理由 |
| `diagnostics` | ✓ | Diagnostic list |

---

## 5. Impact entry

各 `impacts[]` entry は以下の形を持つ。

```jsonc
{
  "id": "impact-NNN",
  "kind": "field_consumer",
  "severity": "breaking",
  "fixability": "mechanical",
  "object": { /* ObjectRef */ },
  "reason": "...",
  "via": ["reads", "model_field"],
  "source": { /* SourceLocation */ },
  "recommended_action": "...",
  "suggested_fixes": [ /* SuggestedFix[] */ ]
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `id` | ✓ | response 内 unique な impact id |
| `kind` | ✓ | impact の semantic kind |
| `severity` | ✓ | §6 参照 |
| `fixability` | ✓ | §6 参照 |
| `object` | ✓ | 影響を受ける ObjectRef または render output |
| `reason` | ✓ | なぜ影響するかの自然言語説明 |
| `via` | ✓ | rootから到達した reference kind の経路 |
| `source` | ✓ | 影響箇所の SourceLocation |
| `recommended_action` | ✓ | 人間向け推奨アクション |
| `suggested_fixes` | 任意 | 機械的修正候補（§8 参照） |

`via` は最短到達経路のみを表す軽量表現とする。
完全な経路復元が必要な場合は、別途 `get_reference_tree` を呼ぶ。

> 由来: ADR-056 §決定 §5, §9

---

## 6. severity / fixability

severity と fixability は別軸として独立に判定する。

### severity

| 値 | 意味 |
|---|---|
| `breaking` | そのまま変更すると semantic build / validation / render / query のいずれかが失敗する可能性が高い |
| `warning` | 壊れるとは限らないが、意味・到達経路・表示・設計意図が変わる可能性がある |
| `info` | 関連情報として提示するが、変更対応は不要または低優先 |

### fixability

| 値 | 意味 |
|---|---|
| `mechanical` | source location と置換内容が一意に決まり、機械的に直せる |
| `suggested` | 修正方針は提案できるが、人間レビュー前提 |
| `manual_review` | 設計判断が必要で、tool では直し方を決めない方がよい |
| `unknown` | tool が判断できない（情報不足、coverage 外、source range 欠落 等） |

### change kind ごとの典型的組み合わせ（参考）

| change kind | typical severity | typical fixability |
|---|---|---|
| `rename` | breaking | mechanical |
| `remove` | breaking | manual_review |
| `change_type`（primitive） | breaking または warning | suggested |
| `change_contract` | breaking | suggested または manual_review |
| `change_transition_target` | warning | manual_review |
| `add` | info または warning | manual_review |

これは目安であり、 個別状況に応じて実装が判断する。

> 由来: ADR-056 §決定 §3

---

## 7. coverage

`coverage` は output 必須フィールドとする。

```jsonc
{
  "coverage": {
    "analyzed": [...],
    "not_analyzed": [...],
    "note": "..."
  }
}
```

`analyzed` / `not_analyzed` は文字列列挙。
LLM はこの内容を人間に提示することで、 「分析対象だが結果0件」 と 「分析していない」 を区別できる。

### v1 標準セット

`coverage.analyzed` の v1 標準語彙:

- `direct_references`
- `reference_tree`
- `model_field_resolution`
- `transition_action_resolution`
- `flow_step_task_resolution`
- `flow_param_field_resolution`
- `sequence_step_task_resolution`
- `type_signature_identity`
- `render_output_files`

`coverage.not_analyzed` の v1 標準語彙:

- `type_structural_compatibility`
- `semantic_contract_compatibility`
- `render_presentation_details`
- `wireframe_element_binding`

### change.kind ごとの coverage サブセット

`change.kind` や対象 object kind により、 `coverage.analyzed` はサブセットになってよい。

例:

- `change.kind = "rename"` の field 対象 → `direct_references`, `reference_tree`, `model_field_resolution`, `flow_param_field_resolution`, `render_output_files`
- `change.kind = "remove"` の task 対象 → `direct_references`, `reference_tree`, `transition_action_resolution`, `flow_step_task_resolution`, `sequence_step_task_resolution`, `render_output_files`
- `change.kind = "add"` の field 対象 → name collision check / type resolution / writer coverage 等。`direct_references` ではない

`add` を投げた場合の `coverage.analyzed` v1 最小実装:

- `name_collision`

`type_resolution` / `writer_coverage` は `add` 専用の将来 coverage 語彙だが、M13 v1 では実体 collector を持たないため `coverage.not_analyzed` に入れる。

### coverage 必須化のルール

- 分析対象0件で返す場合も `coverage` は必須
- `not_analyzed` が空でも、 array としては必ず返す（省略不可）
- `note` は任意。 LLM 向けに補足説明を入れたいときに使う

> 由来: ADR-056 §決定 §6, §7

---

## 8. recommended_action / suggested_fixes

各 impact は `recommended_action`（人間向け）と `suggested_fixes[]`（機械的候補）を二段で返す。

### recommended_action

人間向けの自然言語推奨アクション。 多少抽象でもよい。
LLM が人間に提示する文言の素材となる。

### suggested_fixes

機械的に直せそうな修正候補。
`fixability` の値により出力可否が変わる。

| fixability | suggested_fixes |
|---|---|
| `mechanical` | source location 付きの具体的 fix を出してよい |
| `suggested` | confidence 付きで概念的 fix を出してよい。 source 必須ではない |
| `manual_review` | 空、または非破壊的 advisory のみ |
| `unknown` | 空 |

### SuggestedFix shape

```jsonc
{
  "kind": "replace_reference",
  "confidence": "high",
  "from": "email",
  "to": "email_address",
  "source": {
    "file": "auth/task/login.yaml",
    "line": 42,
    "column": 7
  },
  "note": "..."
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `kind` | ✓ | fix kind（実装裁量。例: `replace_reference`, `update_param_model`） |
| `confidence` | ✓ | `high` / `medium` / `low` |
| `source` | 条件付き | `mechanical` のときは必須。 それ以外は任意 |
| `note` | 任意 | 補足説明 |

`from` / `to` 等の payload は `kind` 依存。

### `fixability=mechanical` の必要条件

`fixability=mechanical` を返してよいのは、 以下を**すべて**満たす場合のみとする。

1. 置換対象 source location が一意に特定できる（file / line / column range）
2. 置換前 token が source 上で一意（誤一致しない）
3. 置換後 token が明確に1つに定まる（衝突なし）
4. 置換後の reference 解決先が変わらない
5. YAML 構造を変えない単純 token 置換である

ひとつでも欠ければ、 最低でも `suggested` に下げる。
不確実性が高い場合は `manual_review` または `unknown` を返す。

実装はこの5要件を judgement gate として持つこと。
個別 change kind ごとの追加 heuristic は実装裁量とする。

> 由来: ADR-056 §決定 §4, §5

---

## 9. SourceLocation

`source` は file / line / column を inline で持つ。
YAML snippet 全文は含めない。

```jsonc
{
  "file": "auth/task/login.yaml",
  "line": 42,
  "column": 7,
  "end_line": 42,
  "end_column": 18
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `file` | ✓ | project root 相対 path |
| `line` | 条件付き | 1-based line number。取得できる場合は必ず返す |
| `column` | 条件付き | 1-based column number。取得できる場合は必ず返す |
| `end_line` | 任意 | range end |
| `end_column` | 任意 | range end |

実装が line / column を取得できない場合も、impact 自体は落とさず `source.file` だけで返してよい。
その場合、該当impactは `fixability=unknown` または `manual_review` に下げ、`diagnostics[]` に `source_location_unavailable` を含める。
`fixability=mechanical` を返すには、file / line / column range が一意に特定できている必要がある。

完全な YAML snippet が必要な場合、 LLM は `get_source` を別途呼ぶ。
`source_preview` として短い行範囲を optional に持ってもよいが、 `analyze_impact` v1 では必須としない。

> 由来: ADR-056 §決定 §9

---

## 10. coverage scope details

### flow wiring の扱い

`analyze_impact` v1 は flow wiring を coverage に含める。
具体的には以下を分析する。

- flow step の task 参照（`flow_step_task_resolution`）
- flow param の field 参照（`flow_param_field_resolution`）

M13 v1 の `flow_param_field_resolution` は最小範囲とする。
対象は、flow param wiring の target/source/return asset 名や source task/join の return model identity が、対象 field / field model に関係すると判定できるケースに限る。
model 間の structural compatibility や、任意の式解析は行わない。

実装は `inspect(task).members.flow.entries` 相当の経路を内部的に読む。
`get_reference_tree` の reference kind は拡張しない（ADR-055 維持）。

### sequence step の扱い

`analyze_impact` v1 は sequence step を coverage に含める。
具体的には以下を分析する。

- sequence step の task 参照（`sequence_step_task_resolution`）

### render output の扱い

`analyze_impact` v1 は render output を **file 粒度** で扱う。

- 変更対象を含む render group の特定
- 該当 render output file path の返却

md 内の presentation 詳細（DAG node shape の変化、ER の線の変化、等）は v1 除外。
`coverage.not_analyzed` に `render_presentation_details` として明示する。

### 型整合性の扱い

`analyze_impact` v1 は **型 signature の identity 比較** までを対象とする。

含まれるもの（`type_signature_identity`）:

- primitive 型一致（`str` / `int` / `bool` 等の比較）
- model 型一致（model id の identity 比較）

含まれないもの（`type_structural_compatibility` として `not_analyzed` に明示）:

- model 間の subtyping 判定
- model fields の structural 互換性
- nullable / optional / required の互換性判定
- semantic contract 互換性

> 由来: ADR-056 §決定 §7, §8

---

## 11. M13 v1 implementation constraints

M13 は full spec を一度に満たすのではなく、強い public contract を守った上での v1 最小実装として close する。

M13 v1 で実装するもの:

- `change` discriminated object validation
- unsupported selector の normal response + `unsupported_selector` diagnostic
- task rename/remove/change_contract の transition action / flow step / sequence step impact
- field rename/remove/change_type の direct/reference-tree based impact
- field 変更に対する最小 `flow_param_field_resolution`
- transition `change_transition_target` の `new_to` / `new_action` 解決チェック
- render output file 粒度 impact
- add の `name_collision`
- `fixability=mechanical` の共通 judgement gate

M13 v1 の known limitations:

- field rename は source line/column が不足する場合、`mechanical` ではなく `unknown` / `manual_review` に落とす。
- `fixability=mechanical` は gate を満たす場合のみ返す。gate を満たさない rename は `suggested` または `unknown` とする。
- flow param 解析は wiring identity の最小判定に限る。
- add の `type_resolution` / `writer_coverage` は `coverage.not_analyzed` に残す。
- state / event / actor / store に対する専用 analyze collector は M13 v1 では限定的であり、主に reference/render 経路で扱う。

---

## 12. assumptions

`assumptions` は tool 側の前提・限界を文字列で列挙する。

例:

- `"rename後のID衝突は検証対象外"`
- `"note内の自然言語参照は解析対象外"`
- `"transitive な flow wiring 影響は depth 1 までのみ"`

これらは LLM が人間に「**こういう前提で分析しています**」と伝えるための情報。
implementation 都合の制約も、 LLM が誤った安心感を人間に与えないために明示する。

---

## 13. Selector support

`analyze_impact` の対象 selector は、 [`get_references`](./get-references.md) で supported な selector を起点にする。

v1 で supported:

- `node: task`
- `node: model`
- `node: store`
- `node: state`
- `node: event`
- `node: actor`
- `transition`
- `field` / `model_field`

v1 で unsupported（`coverage.not_analyzed` に該当 kind を含めて空 impact を返す）:

- `view: api_table`
- `view: er_diagram`
- `view: sequence_diagram`
- `file: *`
- `asset`
- `primitive`
- `render_index`

この一覧は共有 selector support matrix の `analyze_impact` 列と一致させる。
unsupported selector が来た場合、 `diagnostics[]` に `unsupported_selector` を含めて返す。

---

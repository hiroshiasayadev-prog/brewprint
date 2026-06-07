---
scope: docs/spec/mcp/errors.md
status: draft
last_updated: 2026-06-07
summary: >
  MCP toolのerror modelを定義する。
  tool errorとdiagnosticの使い分け、およびerror codeとpayloadを規定する。
depends_on:
  - docs/adr/054-mcp-query-coverage-for-design-conversation.md
---

# Error model

## 1. MCP-level error vs diagnostic

MCP toolの実行自体が成立しない場合はtool errorを返す。

例:

- projectがsemantic validationを通過していない
- selectorの形式が壊れている
- 対象objectが存在しない
- guard未指定でtransitionが曖昧

一方、tool実行は成立したが注意すべき情報がある場合は `diagnostics` に入れる。

例:

- source lineが取得できない
- uncovered moduleが暗黙groupになった
- noteが存在しない
- optionalな周辺情報が未実装

## 2. Error code

MCP v1で定義するerror code:

| code | 意味 |
|---|---|
| `project_invalid` | semantic buildに失敗しておりqueryできない |
| `invalid_args` | tool input JSON または input schema が不正 |
| `invalid_selector` | selectorの形式が不正 |
| `invalid_change_payload` | `analyze_impact.change` の kind / payload の組み合わせが不正 |
| `not_found` | 対象objectが存在しない |
| `kind_mismatch` | selector.kind と解決結果のkindが一致しない |
| `ambiguous` | 候補が複数あり一意に解決できない |
| `unsupported_object` | v1ではquery対象外のobject |
| `unsupported_detail` | `detail` の値が未対応 |
| `unsupported_direction` | `direction` の値が未対応 |
| `invalid_depth` | traversal depth が未対応範囲外 |
| `source_range_unavailable` | object単位のsource rangeを特定できない |
| `internal_error` | 実装内部エラー |

`analyze_impact` では、unsupported selector は tool error にしない。
空 `impacts`、`coverage`、および `unsupported_selector` diagnostic を含む通常responseとして返す。
一方、`change.kind` に対して必須payloadが欠けている場合や、kind と payload の組み合わせが不正な場合は `invalid_change_payload` を返す。

Request option behavior:

- enum-like request option が仕様上の値集合外の場合は、tool-specific code がある場合はそれを優先し、ない場合は `invalid_args` を返す。
- `direction` の未知値は `unsupported_direction` を返す。
- `detail` の未知値は `unsupported_detail` を返す。
- `get_reference_tree.depth` が `0..4` の範囲外の場合は `invalid_depth` を返す。
- `get_source.fallback` が `file` / `error` 以外の場合は `invalid_args` を返す。
- `source_range_unavailable` は object単位のsource rangeを特定できない同じ根本条件を表す。
- `get_source(fallback=file)` または `fallback` 省略時は、`source_range_unavailable` を warning diagnostic として返す。
- `get_source(fallback=error)` では、`source_range_unavailable` を tool error として返す。
- `source_range_unavailable` の surface / severity は request の `fallback` option により決まる。
- request option の default / omitted behavior は各 tool spec の contract として扱い、DATA DSL の default 構文としては扱わない。

## 3. Error payload

```json
{
  "error": {
    "code": "not_found",
    "message": "object not found: auth.task.missing_login",
    "selector": {
      "id": "auth.task.missing_login"
    },
    "diagnostics": []
  }
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `code` | ✓ | error code |
| `message` | ✓ | human-readable message |
| `selector` | 任意 | 入力selector |
| `diagnostics` | 任意 | 関連diagnostic |

---

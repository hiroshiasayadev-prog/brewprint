---
scope: docs/spec/mcp/errors.md
status: draft
last_updated: 2026-04-30
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
| `invalid_selector` | selectorの形式が不正 |
| `not_found` | 対象objectが存在しない |
| `kind_mismatch` | selector.kind と解決結果のkindが一致しない |
| `ambiguous` | 候補が複数あり一意に解決できない |
| `unsupported_object` | v1ではquery対象外のobject |
| `unsupported_detail` | `detail` の値が未対応 |
| `unsupported_direction` | `direction` の値が未対応 |
| `invalid_depth` | traversal depth が未対応範囲外 |
| `internal_error` | 実装内部エラー |

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

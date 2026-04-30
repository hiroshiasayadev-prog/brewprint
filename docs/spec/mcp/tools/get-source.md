---
scope: docs/spec/mcp/tools/get-source.md
status: draft
last_updated: 2026-04-30
summary: >
  get_source toolの仕様を定義する。
  semantic objectに対応するYAML source snippetを返す。
  fallback時の挙動とselector supportを規定する。
depends_on:
  - docs/adr/054-mcp-query-coverage-for-design-conversation.md
---

# `get_source`

## 1. Purpose

`get_source` は、対象semantic objectに対応するYAML source snippetを返す。

返すもの:

- object identity
- source file / range
- YAML snippet
- fallbackした場合の理由を示すdiagnostic

返さないもの:

- Raw YAML AST
- semantic build前の未解決構造全体
- project外fileの内容
- renderer output

`get_source` はADR-054の方針に従い、Raw YAML AST公開APIではなく、ResolvedProject上のsemantic objectに紐づくsource補助情報として扱う。

## 2. Input

```json
{
  "selector": {
    "id": "auth.task.login"
  },
  "fallback": "file"
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `selector` | ✓ | Object selector |
| `fallback` | 任意 | `file` / `error`。省略時は `file` と同等 |

`fallback=file` または省略時は、object単位のrangeが特定できない場合に、同じFileIDのYAML全体を返し、`diagnostics[]` に `source_range_unavailable` warningを入れる。
`fallback=error` の場合は、object単位のrangeが特定できないとtool errorを返す。

## 3. Output

```json
{
  "object": {
    "object": "node",
    "kind": "task",
    "id": "auth.task.login",
    "qualified_id": "auth.task.login",
    "label": "login",
    "file": "auth/task/login.yaml"
  },
  "source": {
    "file": "auth/task/login.yaml",
    "line": 3,
    "column": 5,
    "end_line": 18,
    "end_column": 1
  },
  "snippet": {
    "language": "yaml",
    "text": "  - id: login\n    type: task\n    ..."
  },
  "diagnostics": []
}
```

| フィールド | 必須 | 内容 |
|---|---:|---|
| `object` | ✓ | 対象ObjectRef |
| `source` | ✓ | SourceLocation。line / columnが取得できない場合は `file` のみでもよい |
| `snippet` | ✓ | `language: yaml` と snippet text |
| `fallback` | 任意 | fallbackした場合は `file` |
| `diagnostics` | ✓ | Diagnostic list |

## 4. Selector support

`get_source` は、MCP v1でquery可能なsemantic objectを対象にする。

初期実装でsnippet rangeを返す対象:

- `node` / private sub node: `nodes[]` 内の該当item
- `field`: parent modelの `fields[]` 内の該当item
- `transition`: `transitions[]` 内の該当item
- `asset`: producer nodeの `returns` block。特定不能時はproducer nodeへfallbackしてよい
- `view`: view file全体
- `file`: file全体

source line/columnが取得できない実装、または対応objectの局所rangeを特定できない実装では、`fallback=file` により同一FileID全体を返してよい。
この場合は `diagnostics[]` に以下を入れる。

```json
{
  "severity": "warning",
  "code": "source_range_unavailable",
  "file": "auth/task/login.yaml",
  "message": "source range is unavailable; returned whole file"
}
```

> 由来: ADR-054 §決定 §5

---

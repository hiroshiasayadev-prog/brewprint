---
scope: docs/spec/design-records-mcp/tools.md
status: draft
last_updated: 2026-05-12
summary: >
  Design Records MCP MVP の read-only tool interface と責務境界を定義する。
depends_on:
  - docs/adr/076-design-records-mcp.md
  - docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md
design_record:
  id: SPEC-design-records-mcp-tools
  kind: spec
  status: draft
  depends_on:
    - ADR-076
    - ADR-077
---

# Design Records MCP tools

## Tool set

Design Records MCP MVP の P0 tool は以下である。

| tool | priority | purpose |
|---|---|---|
| `list_records` | P0 | record index を構造化して返す |
| `get_record` | P0 | record ID から metadata / path / headings / raw body を取得する |
| `validate_records` | P0 | record metadata の基本整合性を検査する |

P1 の任意補助 tool として以下を許容する。

| tool | priority | purpose |
|---|---|---|
| `suggest_next_record` | P1 | 次の ADR ID と推奨 path を提案する |

MVP tool は read-only である。
ファイル作成・更新・Evidence 書き換え・commit 操作は行わない。

> 由来: ADR-077 §P0: MVP必須tool, ADR-077 §P1: MVPに含めてもよい補助tool

## Common response conventions

MVP tool response は、record を返す場合、以下の基本形を使う。

```json
{
  "id": "ADR-076",
  "kind": "decision",
  "title": "Design Records MCP",
  "status": "accepted",
  "path": "docs/adr/076-design-records-mcp.md",
  "depends_on": ["ADR-050", "ADR-068"],
  "supersedes": [],
  "migrated_to_spec": null
}
```

`title` は H1 から抽出する。
`path` は repository root からの相対 path とする。
repository root は、Design Records MCP 起動時の current working directory、または起動引数で明示された root path とする。

MVP では response 内で Markdown 本文を整形・要約・正規化しない。

> 由来: ADR-077 §list_records の責務, ADR-077 §get_record の責務

## `list_records`

### Purpose

`list_records` は、ADR/spec record index を構造化して返す query tool である。

目的は、Markdown 本文を読む前に候補 record を絞り込むことである。
単なる filesystem listing ではなく、ADR の箇条書きmetadataまたは spec の YAML front matter から正規化した record metadata と H1 title を含む一覧を返す。

> 由来: ADR-077 §list_records の責務

### Request

MVP request schema:

```json
{
  "kind": "decision",
  "status": "accepted",
  "id": "ADR-076",
  "id_range": {
    "from": "ADR-067",
    "to": "ADR-077"
  },
  "order_by": "id",
  "order": "asc",
  "limit": 20
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `kind` | no | string | `decision` / `spec` で絞り込む |
| `status` | no | string | status で絞り込む |
| `id` | no | string | exact ID で絞り込む |
| `id_range` | no | object | ID range で絞り込む |
| `order_by` | no | string | MVP では `id` のみ |
| `order` | no | string | `asc` / `desc` |
| `limit` | no | integer | 最大件数 |

`order_by` は MVP では `id` のみをサポートする。
`head` / `tail` は採用しない。

`id_range` は両端を含む範囲とする。
`from` または `to` の片方だけを指定してもよい。

MVP の `id_range` は `ADR-NNN` 形式の `decision` record にのみ適用する。
比較は `NNN` の数値比較とする。
`kind` が省略され、かつ `id_range` が指定された場合は `kind: decision` と同等に扱う。
`id_range` を指定する request では、`kind` は省略されているか `decision` でなければならない。
`kind: spec` と `id_range` の併用、または `SPEC-*` への range 指定は request error とする。

### Response

```json
{
  "records": [
    {
      "id": "ADR-076",
      "kind": "decision",
      "title": "Design Records MCP",
      "status": "accepted",
      "path": "docs/adr/076-design-records-mcp.md",
      "depends_on": ["ADR-050", "ADR-068"],
      "supersedes": [],
      "migrated_to_spec": null
    }
  ]
}
```

`records[]` の並び順は `order_by` / `order` に従う。

## `get_record`

### Purpose

`get_record` は、record ID から metadata、path、headings、必要に応じて Markdown 本文を取得する read tool である。

ADR 番号から path や本文を取得できることで、候補絞り込み後に filesystem tool へ戻る回数を減らす。

> 由来: ADR-077 §get_record の責務

### Request

```json
{
  "id": "ADR-076",
  "include_body": false
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `id` | yes | string | 取得対象 record ID |
| `include_body` | no | bool | Markdown raw body を返すか。default は `false` |

### Response without body

```json
{
  "record": {
    "id": "ADR-076",
    "kind": "decision",
    "title": "Design Records MCP",
    "status": "accepted",
    "path": "docs/adr/076-design-records-mcp.md",
    "depends_on": ["ADR-050", "ADR-068"],
    "supersedes": [],
    "migrated_to_spec": null,
    "headings": [
      { "level": 1, "text": "076: Design Records MCP" },
      { "level": 2, "text": "背景" },
      { "level": 2, "text": "決定" }
    ]
  }
}
```

### Response with body

`include_body: true` の場合、`record.body` に Markdown 本文を追加する。

```json
{
  "record": {
    "id": "ADR-076",
    "kind": "decision",
    "title": "Design Records MCP",
    "status": "accepted",
    "path": "docs/adr/076-design-records-mcp.md",
    "depends_on": ["ADR-050", "ADR-068"],
    "supersedes": [],
    "migrated_to_spec": null,
    "headings": [],
    "body": "# 076: Design Records MCP\n\n- **status**: accepted\n..."
  }
}
```

`body` は元ファイル内容をそのまま返す。
整形・要約・正規化を行ってはならない。
構造化 metadata や headings は body とは別 field として返す。

## `validate_records`

### Purpose

`validate_records` は、Design Records MCP の metadata index が信頼できる状態かを検証する tool である。

MVP では record metadata の基本整合性検査に限定する。
運用 gap 診断や semantic trace は扱わない。

> 由来: ADR-077 §validate_records の責務

### Request

```json
{
  "kind": "decision",
  "id_range": {
    "from": "ADR-067",
    "to": "ADR-077"
  }
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `kind` | no | string | 検証対象 kind を絞る |
| `id_range` | no | object | 検証対象 ID 範囲を絞る |

request が空の場合、MVP index 対象の全 record を検証する。

`id_range` の扱いは `list_records` と同じく、`ADR-NNN` 形式の `decision` record 専用とする。
`kind: spec` と `id_range` の併用、または `SPEC-*` への range 指定は request error とする。

### Response

```json
{
  "ok": false,
  "diagnostics": [
    {
      "category": "missing_depends_on_target",
      "severity": "error",
      "record_id": "ADR-077",
      "path": "docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md",
      "message": "depends_on references missing record ADR-999",
      "target_id": "ADR-999"
    }
  ]
}
```

| field | meaning |
|---|---|
| `ok` | error diagnostic がない場合 true |
| `diagnostics` | diagnostic list |

Diagnostic object は少なくとも以下を持つ。

| field | required | meaning |
|---|---:|---|
| `category` | yes | diagnostic category |
| `severity` | yes | MVP では `error` |
| `record_id` | no | 問題がある record ID |
| `path` | no | 問題がある path |
| `message` | yes | human-readable message |
| `target_id` | no | 参照切れなどの対象 ID |

### Diagnostic categories

MVP diagnostic category は `schema.md` の定義に従う。

- `duplicate_id`
- `filename_id_mismatch`
- `invalid_h1_title`
- `invalid_status_for_kind`
- `spec_status_mismatch`
- `missing_depends_on_target`
- `missing_supersedes_target`
- `invalid_migrated_to_spec`
- `missing_record_path`

`accepted_but_not_migrated` / `missing_design_record` などの運用 gap 診断は MVP 外である。

`missing_record_path` は、filesystem scan または path normalization により record 候補 path を検出したが、実際の read/stat に失敗した場合に出す。
例として、scan 後に file が削除された場合、permission denied、symlink target missing、path normalization 後の path が存在しない場合を含む。

## `suggest_next_record`

### Purpose

`suggest_next_record` は、新規 ADR 起票を補助する P1 read-only tool である。

既存 record index から次の ADR ID と推奨 path を提案する。
ファイル作成は行わない。

`next_number` は既存 `decision` record の最大番号 + 1 とする。
欠番は埋めない。

> 由来: ADR-077 §suggest_next_record の責務

### Request

```json
{
  "kind": "decision",
  "title": "Design Records MCP implementation package layout"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `kind` | yes | string | MVP では `decision` のみ |
| `title` | yes | string | 新規 ADR title |

MVP では spec 新規作成の path 提案は扱わない。

### Response

```json
{
  "kind": "decision",
  "title": "Design Records MCP implementation package layout",
  "next_id": "ADR-078",
  "next_number": 78,
  "suggested_path": "docs/adr/078-design-records-mcp-implementation-package-layout.md",
  "existing_max_id": "ADR-077"
}
```

`suggested_path` は提案であり、ファイル作成は行わない。

filename slug は `title` から生成する。
MVP の slug 生成規則は以下とする。

- ASCII 英数字を lowercase 化する
- ASCII 英数字以外の連続を `-` に変換する
- 連続する `-` は1つにまとめる
- 前後の `-` は除去する
- 非 ASCII 文字は `-` として扱う

slug が空になる場合、`suggested_path` は `docs/adr/{NNN}.md` としてよい。
`suggested_path` は提案であり、人間が起票時に上書きしてよい。

## Error handling

MVP tool error code は以下を最小とする。

| code | meaning |
|---|---|
| `record_not_found` | 指定された record ID が存在しない |
| `invalid_request` | request schema または field value が不正。例: `list_records` に `kind: task` を指定した場合 |
| `unsupported_kind` | tool が対象外の `kind` を指定された。例: `suggest_next_record` に `kind: spec` を指定した場合 |
| `id_range_requires_decision_kind` | `id_range` が `decision` 以外の kind と併用された、または `SPEC-*` range が指定された |

MVP では、存在しない record ID を指定した場合、tool は machine-readable な error を返す。

例:

```json
{
  "error": {
    "code": "record_not_found",
    "message": "record ADR-999 was not found"
  }
}
```

`record_not_found` は tool 実行 error であり、`validate_records` diagnostic category ではない。

## Write tool policy

MVP では write 系 tool を提供しない。

以下は MVP 外である。

- `create_record`
- `update_record`
- `set_evidence`
- `add_record_metadata`
- `migrate_record_to_spec`

write 系 tool を導入する場合は、dry-run diff、ユーザー確認、衝突処理、template 責務、git 運用との境界を別 ADR / spec で定義する。

> 由来: ADR-077 §MVP外, ADR-077 §理由

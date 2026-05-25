---
scope: docs/spec/design-records-mcp/tools.md
status: draft
last_updated: 2026-05-26
summary: >
  Design Records MCP MVP の read-only tool interface と責務境界を定義する。
depends_on:
  - docs/adr/076-design-records-mcp.md
  - docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
  - docs/adr/090-design-records-mcp-batch-retrieval-tool-boundary.md
design_record:
  id: SPEC-design-records-mcp-tools
  kind: spec
  status: draft
  depends_on:
    - ADR-076
    - ADR-077
    - ADR-087
    - ADR-088
    - ADR-090
---

# Design Records MCP tools

## Tool set

Design Records MCP MVP の P0 tool は以下である。

| tool | priority | purpose |
|---|---|---|
| `list_records` | P0 | record index を構造化して返す |
| `get_record` | P0 | 単一 record ID から metadata / path / headings / raw body を取得する |
| `get_records` | P0 | 明示された複数 record ID の detail representation をまとめて取得する |
| `validate_records` | P0 | record metadata の基本整合性と canonical reference validation を検査する |
| `resolve_reference` | P0 | canonical semantic/artifact reference を document / section / record へ解決する |

P1 の任意補助 tool として以下を許容する。

| tool | priority | purpose |
|---|---|---|
| `suggest_next_record` | P1 | 次の ADR ID と推奨 path を提案する |

MVP tool は read-only である。
ファイル作成・更新・Evidence 書き換え・commit 操作は行わない。

> 由来: ADR-077 §P0: MVP必須tool, ADR-077 §P1: MVPに含めてもよい補助tool, ADR-090 §決定

## Common response conventions

record を返す tool response は、共通 field と kind 固有 detail object を分離する。

Decision example:

```json
{
  "id": "ADR-076",
  "kind": "decision",
  "title": "Design Records MCP",
  "status": "accepted",
  "path": "docs/adr/076-design-records-mcp.md",
  "decision": {
    "depends_on": ["ADR-050", "ADR-068"],
    "supersedes": [],
    "migrated_to_spec": null
  }
}
```

Investigation example:

```json
{
  "id": "INV-MCP-001",
  "kind": "investigation",
  "title": "Design Records MCP investigation support",
  "status": "concluded",
  "path": "docs/investigations/mcp/INV-MCP-001-design-records-mcp-investigation-support.md",
  "investigation": {
    "trigger": "ADR-087",
    "scope": "investigation MCP integration",
    "non_scope": "writer tools",
    "source_refs": ["ADR-086", "ADR-087"],
    "follow_up_candidates": ["ADR-088"]
  }
}
```

旧 flat response field と kind 固有 detail object は併存させない。spec を新 contract に更新した後、実装と tests は同一の切替単位で追従する。

`title` は H1 から抽出する。
`path` は repository root からの相対 path とする。
repository root は、Design Records MCP 起動時の current working directory、または起動引数で明示された root path とする。

MVP では response 内で Markdown 本文を整形・要約・正規化しない。

> 由来: ADR-077 §list_records の責務, ADR-077 §get_record の責務

## `list_records`

### Purpose

`list_records` は、decision / spec / investigation を扱う record index を構造化して返す query tool である。

目的は、Markdown 本文を読む前に候補 record を絞り込むことである。
単なる filesystem listing ではなく、ADR / investigation の箇条書きmetadataまたは spec の YAML front matter から正規化した record metadata と H1 title を含む一覧を返す。

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
| `kind` | no | string | `decision` / `spec` / `investigation` で絞り込む |
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
`kind: spec` / `kind: investigation` と `id_range` の併用、または `SPEC-*` / `INV-*` への range 指定は request error とする。
Investigation の domain-scoped ID に対する range / domain filter はこの版では扱わず、後続 spec refinement に委ねる。

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
      "decision": {
        "depends_on": ["ADR-050", "ADR-068"],
        "supersedes": [],
        "migrated_to_spec": null
      }
    }
  ]
}
```

`records[]` の並び順は `order_by` / `order` に従う。

`order_by: id` で mixed kind の record を返す場合は、canonical `id` の ASCII lexical order を用いる。同一 canonical `id` が複数 entry に存在する場合も並び順は path の ASCII lexical order で安定化し、`duplicate_id` diagnostic は別途返す。
`decision` の `id_range` は従来どおり `ADR-NNN` の `NNN` を数値比較する。

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
    "decision": {
      "depends_on": ["ADR-050", "ADR-068"],
      "supersedes": [],
      "migrated_to_spec": null
    },
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
    "decision": {
      "depends_on": ["ADR-050", "ADR-068"],
      "supersedes": [],
      "migrated_to_spec": null
    },
    "headings": [],
    "body": "# 076: Design Records MCP\n\n- **status**: accepted\n..."
  }
}
```

`body` は元ファイル内容をそのまま返す。
整形・要約・正規化を行ってはならない。
構造化 metadata や headings は body とは別 field として返す。

## `get_records`

### Purpose

`get_records` は、呼び出し側が明示した複数 record ID について、`get_record` と同じ record representation をまとめて取得する read-only tool である。

候補探索と filter / range query は `list_records`、canonical reference resolution は `resolve_reference`、index integrity validation は `validate_records` の責務である。`get_records` はこれらを兼務しない。

> 由来: ADR-090 §1〜§3

### Request

```json
{
  "ids": [
    "ADR-077",
    "SPEC-design-records-mcp-tools",
    "INV-DOCS-001",
    "ADR-077",
    "INV-DOCS-999"
  ],
  "include_body": false
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `ids` | yes | non-empty array of string | 取得対象の exact record ID lookup key。入力順を保持する |
| `include_body` | no | bool | 各 found record に Markdown raw body を含めるか。default は `false` |

`ids` の欠落、空配列、array 以外の値、または string 以外の element は `invalid_request` tool error とする。

`ids[]` は record index に対する exact lookup key としてのみ評価する。前後 whitespace の trim、case normalization、canonical reference resolution、input kind classification は行わない。したがって `spec:trace`、`REQ-MCP-002`、physical path、`adr-077`、` ADR-077 ` のような string input が indexed record ID と一致しない場合、tool error や `unsupported` ではなく item-level `not_found` とする。

`get_records` は `kind` / `status` / `id_range` / `limit` を request field として持たない。record ごとに異なる `include_body` を指定する query plan 形式も採用しない。

### Response

```json
{
  "items": [
    {
      "id": "ADR-077",
      "retrieval_status": "found",
      "record": {
        "id": "ADR-077",
        "kind": "decision",
        "title": "Design Records MCP MVP boundary and tool prioritization",
        "status": "accepted",
        "path": "docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md",
        "decision": {
          "depends_on": ["ADR-076"],
          "supersedes": [],
          "migrated_to_spec": null
        },
        "headings": []
      },
      "diagnostics": []
    },
    {
      "id": "SPEC-design-records-mcp-tools",
      "retrieval_status": "found",
      "record": {
        "id": "SPEC-design-records-mcp-tools",
        "kind": "spec",
        "title": "Design Records MCP tools",
        "status": "draft",
        "path": "docs/spec/design-records-mcp/tools.md",
        "spec": {
          "depends_on": ["ADR-076", "ADR-077", "ADR-087", "ADR-088", "ADR-090"]
        },
        "headings": []
      },
      "diagnostics": []
    },
    {
      "id": "INV-DOCS-001",
      "retrieval_status": "found",
      "record": {
        "id": "INV-DOCS-001",
        "kind": "investigation",
        "title": "investigation artifact format and lifecycle",
        "status": "concluded",
        "path": "docs/investigations/docs/INV-DOCS-001-investigation-artifact-format-and-lifecycle.md",
        "investigation": {
          "trigger": "ADR-085",
          "scope": "investigation artifact format and lifecycle",
          "non_scope": "public contract decisions",
          "source_refs": ["ADR-085"],
          "follow_up_candidates": []
        },
        "headings": []
      },
      "diagnostics": []
    },
    {
      "id": "INV-DOCS-999",
      "retrieval_status": "not_found",
      "record": null,
      "diagnostics": [
        {
          "category": "record_not_found",
          "severity": "error",
          "requested_id": "INV-DOCS-999",
          "message": "record INV-DOCS-999 was not found"
        }
      ]
    }
  ],
  "diagnostics": [
    {
      "category": "duplicate_requested_id_ignored",
      "severity": "info",
      "requested_id": "ADR-077",
      "first_index": 0,
      "duplicate_indexes": [3],
      "message": "duplicate requested record ID was ignored after its first occurrence"
    }
  ]
}
```

Top-level response fields:

| field | required | meaning |
|---|---:|---|
| `items` | yes | dedupe 後の first occurrence order で並ぶ retrieval item list |
| `diagnostics` | yes | request-level diagnostic list。正常時は empty list |

Retrieval item fields:

| field | required | meaning |
|---|---:|---|
| `id` | yes | request で指定された lookup key |
| `retrieval_status` | yes | `found` / `not_found` |
| `record` | yes | `found` では `get_record.record` と同一 representation、`not_found` では `null` |
| `diagnostics` | yes | item-level diagnostic list。`found` では empty list |

`items` は `records` と呼ばない。missing item も同じ collection に含むためである。

全 ID が存在しない場合も tool error にはせず、normal response として各 first-occurrence input に `retrieval_status: "not_found"` item を返す。

同一 ID が複数回指定された場合、最初の出現だけを `items` に返す。重複した ID ごとに top-level `duplicate_requested_id_ignored` diagnostic を一件返し、`first_index` と `duplicate_indexes` は request `ids` 配列の zero-based index とする。

`include_body: true` の場合、各 `found` item の `record.body` に元 Markdown 全文を追加する。`get_record` と同様に本文の整形・要約・正規化・truncate は行わない。Response total length / body size の public numeric limit は定義しない。

> 由来: ADR-090 §4〜§7

## `resolve_reference`

### Purpose

`resolve_reference` は、MVP canonical reference を単一の document / section / record target に解決する read-only tool である。Validation はこの resolver と同一の lookup 規則を用い、別の解決規則を持ってはならない。

### Request

```json
{
  "ref": "spec:trace.semantic-ref.definition"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `ref` | yes | string | 解決対象の canonical reference candidate。前後 whitespace は許容せず、入力文字列をそのまま評価する |

Supported input は以下のみとする。

| input form | ref kind | lookup source |
|---|---|---|
| active `spec:` document-level ref | `semantic_ref` | spec front matter `semantic_refs` |
| active `spec:` section-level ref | `semantic_ref` | spec front matter `sections` |
| `ADR-NNN` | `record_id` | `decision` record index |
| `SPEC-<slug>` | `record_id` | `spec` record index |
| `INV-<DOMAIN>-NNN` | `record_id` | `investigation` record index |

`internal-design:` / `coverage:`、`COV-*`、`REQ-*`、`WORK-*`、physical path、および grammar に合わない ID form は supported input ではなく、direct query では tool execution error ではなく `status: "unsupported"` を返す。`yaml:` は reserved prefix だが、MVP は public resolver input または direct query response behavior を定義しない。

### Response

MVP が behavior を定義する direct query response は、常に以下の top-level field を持つ。

| field | required | meaning |
|---|---:|---|
| `ref` | yes | request で受け取った文字列 |
| `ref_kind` | yes | `semantic_ref` / `record_id` / `unsupported` |
| `status` | yes | `resolved` / `unresolved` / `unsupported` |
| `target` | yes | `resolved` の場合は target object、それ以外は `null` |
| `diagnostics` | yes | resolution diagnostic list。正常解決では empty list |

Resolved section-level `spec:` example:

```json
{
  "ref": "spec:trace.semantic-ref.definition",
  "ref_kind": "semantic_ref",
  "status": "resolved",
  "target": {
    "target_type": "section",
    "path": "docs/spec/concepts/traceability/semantic-ref.md",
    "section": "Semantic ref definition"
  },
  "diagnostics": []
}
```

Resolved document-level `spec:` target は `target_type: "document"` と `path` を返し、`section` を持たない。入力 canonical ref は top-level `ref` に保持するため、target に重複して返さない。Resolved section-level `spec:` target は `target_type: "section"`、`path`、`section` を返す。MVP は section-level ref と document-level ref の親子 relation を public response として定義せず、section-level ref の文字列 prefix から親 document ref を推定しない。Resolved record ID-as-ref target は `target_type: "record"`、`path`、`record_id`、`record_kind`、`title`、`status` を返す。

Supported form だが lookup target が存在しない場合は `status: "unresolved"`、`target: null` とし、`diagnostics` に `unresolved_reference` を含める。同一 `spec:` ref または record ID が複数 target へ解決される場合、任意の一件を返してはならない。`status: "unresolved"`、`target: null` とし、`ambiguous_reference` diagnostic を返す。Validation では同一原因を `duplicate_semantic_ref` または `duplicate_id` の `error` として報告する。

Unsupported example:

```json
{
  "ref": "internal-design:resolver.semantic-ref-index",
  "ref_kind": "unsupported",
  "status": "unsupported",
  "target": null,
  "diagnostics": [
    {
      "category": "unsupported_reference",
      "severity": "info",
      "message": "reference form is outside the MVP resolver contract"
    }
  ]
}
```

Direct query の `unsupported_reference` は resolver の failure ではなく input boundary の可視化であるため `info` とする。ただし unsupported input が investigation metadata の validation 対象 field に現れた場合の severity は、下記 `validate_records` の `unsupported_reference` contract に従う。Reserved `yaml:` の public resolver input / direct query response behavior、および investigation metadata validation behavior は MVP では定義しない。

## `validate_records`

### Purpose

`validate_records` は、Design Records MCP の metadata index が信頼できる状態かを検証する tool である。

record metadata の基本整合性検査に加え、active `spec:` semantic ref、record ID-as-ref (`ADR-*` / `SPEC-*` / `INV-*`)、および investigation の `source_refs` / 記載済み `follow_up_results` が canonical reference として解決可能であることを検査する。
`follow_up_candidates` に artifact reference が記載された場合は canonical form を検査する。Canonical form の unresolved candidate は予定された後続 artifact が未作成であることを示す `info` diagnostic とし、physical path による candidate は noncanonical candidate を示す `info` diagnostic とする。
ADR-088 により、`internal-design:` / `coverage:` / `COV-*`、semantic realization relation、coverage mapping query は MVP required scope として扱わない。

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
`kind: spec` / `kind: investigation` と `id_range` の併用、または `SPEC-*` / `INV-*` への range 指定は request error とする。

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
| `severity` | yes | `error` または `info` |
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

Canonical reference / investigation validation の concrete category と severity は以下とする。

| category | severity | field / condition |
|---|---|---|
| `invalid_semantic_ref_declaration` | `error` | spec front matter の `semantic_refs` entry または `sections` key が active `spec:` grammar に従わない |
| `missing_section_target` | `error` | spec front matter の `sections` value と一致する Markdown heading が存在しない |
| `ambiguous_section_target` | `error` | spec front matter の `sections` value が同一 document 内の複数 heading に一致し、section target を単一解決できない |
| `duplicate_semantic_ref` | `error` | 同一 active `spec:` ref が複数 target に宣言された |
| `unresolved_source_ref` | `error` | investigation `source_refs` の supported canonical ref が解決不能 |
| `unresolved_follow_up_result` | `error` | investigation `follow_up_results` の supported canonical ref が解決不能 |
| `unresolved_follow_up_candidate` | `info` | investigation `follow_up_candidates` の supported canonical ref が未解決 |
| `noncanonical_source_ref` | `error` | investigation `source_refs` に physical path が記載された |
| `noncanonical_follow_up_result` | `error` | investigation `follow_up_results` に physical path が記載された |
| `noncanonical_follow_up_candidate` | `info` | investigation `follow_up_candidates` に physical path が記載された |
| `unsupported_reference` | `error` / `info` | MVP が unsupported と定義する metadata reference。`source_refs` / `follow_up_results` では `error`、`follow_up_candidates` では `info`。Reserved `yaml:` はこの category の対象に含めず、MVP では behavior を定義しない |

Investigation reference diagnostic (`unresolved_*` / `noncanonical_*` / metadata field 由来の `unsupported_reference`) は、既存の diagnostic field に加えて `field`（`source_refs` / `follow_up_results` / `follow_up_candidates`）、`value`（入力 ref 文字列）、`ref_status`（`unresolved` / `unsupported` / `noncanonical`）を必須で返す。対象が record ID-as-ref の場合は `target_id` も返してよい。Investigation metadata が duplicate semantic ref または duplicate record ID を指して単一解決できない場合は field-specific diagnostic を追加せず、index defect を示す `duplicate_semantic_ref` または `duplicate_id` のみを返す。これら duplicate diagnostic および spec declaration / section lookup diagnostic は investigation metadata field 由来の追加 field を要求しない。

`depends_on` は `ADR-*` / `SPEC-*` / `INV-*` の canonical record ID-as-ref を参照できる。したがって `ADR-086` の `depends_on: INV-DOCS-001` は contract 上 valid であり、investigation integration 実装前に返る `missing_depends_on_target` は既知 implementation gap である。M19 implementation 後は `INV-DOCS-001` の index integration によりこの diagnostic が解消されなければならない。

`ok` は `error` diagnostic がない場合に `true` とし、`info` diagnostic が存在しても `false` にしない。
Coverage mapping、semantic realization relation、`internal-design:` / `coverage:` / `COV-*` の解決・診断は MVP tool acceptance に含めない。

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
| `record_not_found` | `get_record` で指定された単一 record ID が存在しない。`get_records` では tool error ではなく item-level diagnostic として用いる |
| `invalid_request` | request schema または field value が不正。例: `list_records` に `kind: task` を指定した場合、または `get_records.ids` が欠落・空・非 array・非 string element を含む場合 |
| `unsupported_kind` | tool が対象外の `kind` を指定された。例: `suggest_next_record` に `kind: spec` を指定した場合 |
| `id_range_requires_decision_kind` | `id_range` が `decision` 以外の kind と併用された、または `SPEC-*` / `INV-*` range が指定された |

`get_record` では、存在しない単一 record ID を指定した場合、tool は machine-readable な error を返す。

例:

```json
{
  "error": {
    "code": "record_not_found",
    "message": "record ADR-999 was not found"
  }
}
```

`get_record` の `record_not_found` は tool execution error であり、`validate_records` diagnostic category ではない。`get_records` の missing requested ID は batch response 内の item-level `record_not_found` diagnostic として返し、batch tool execution 自体は成功とする。

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

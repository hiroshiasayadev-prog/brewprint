---
scope: docs/spec/design-records-mcp/tools.md
status: draft
last_updated: 2026-06-05
summary: >
  Design Records MCP の read/navigation/guidance tool interface と
  authoring transaction tool interface の責務境界を定義する。
depends_on:
  - docs/adr/076-design-records-mcp.md
  - docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
  - docs/adr/090-design-records-mcp-batch-retrieval-tool-boundary.md
  - docs/adr/092-design-records-mcp-workflow-artifact-record-and-relation-boundary.md
  - docs/adr/093-design-records-mcp-authoring-transaction-model.md
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
    - ADR-092
    - ADR-093
---

# Design Records MCP tools

## Tool set

Design Records MCP の P0 read / navigation / guidance tool は以下である。

| tool | priority | purpose |
|---|---|---|
| `list_records` | P0 | record index を構造化して返す |
| `get_record` | P0 | 単一 record ID から metadata / path / headings / raw body を取得する |
| `get_records` | P0 | 明示された複数 record ID の detail representation をまとめて取得する |
| `validate_records` | P0 | record metadata の基本整合性と canonical reference validation を検査する |
| `resolve_reference` | P0 | canonical semantic/artifact reference を document / section / record へ解決する |
| `list_authoring_guides` | P0 | project authoring guide catalog を guide ID ベースで返す |
| `get_authoring_guidance` | P0 | guide ID から authoring guidance Markdown を取得する |

P1 の任意補助 tool として以下を許容する。

| tool | priority | purpose |
|---|---|---|
| `suggest_next_record` | P1 | 次の ADR ID と推奨 path を提案する |

Authoring transaction MVP の write tool は proposal-first で別 surface として定義する。
Proposal creation は repository file を変更しない。
Repository file を変更できる authoring operation は `accept_proposed_write` のみである。

Authoring transaction MVP tool は以下である。

| tool | priority | purpose |
|---|---|---|
| `propose_record_create` | P0 | artifact-oriented create proposal を作成し、diff と validation result を返す |
| `propose_record_update` | P0 | metadata block または named section の replacement proposal を作成し、diff と validation result を返す |
| `get_proposed_write` | P0 | proposal ID から retained proposal detail を取得する |
| `accept_proposed_write` | P0 | proposal ID の diff を accept し、repository file への write を試みる |
| `discard_proposed_write` | P0 | proposal ID を破棄し、以後 accept できない状態にする |

> 由来: ADR-077 §P0: MVP必須tool, ADR-077 §P1: MVPに含めてもよい補助tool, ADR-090 §決定, ADR-093 §決定

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

Workflow artifact example:

```json
{
  "id": "WORK-MCP-003",
  "kind": "work_item",
  "title": "Workflow artifact MCP support の最小 public contract を判断・実現する",
  "status": "in_progress",
  "path": "docs/work-items/mcp/WORK-MCP-003-workflow-artifact-mcp-support.md",
  "work_item": {
    "source_requirement": "REQ-MCP-003",
    "impact_refs": ["ADR-092", "SPEC-design-records-mcp-tools"],
    "tasks": ["TASK-MCP-003-01", "TASK-MCP-003-02", "TASK-MCP-003-03"]
  }
}
```

旧 flat response field と kind 固有 detail object は併存させない。spec を新 contract に更新した後、実装と tests は同一の切替単位で追従する。

`title` は H1 から抽出する。
`path` は repository root からの相対 path とする。
repository root は、Design Records MCP 起動時の current working directory、または起動引数で明示された root path とする。

MVP では response 内で Markdown 本文を整形・要約・正規化しない。

Authoring guidance tool response は design record response ではない。Guide source path は public response contract に含めず、guide ID から内部解決する。

> 由来: ADR-077 §list_records の責務, ADR-077 §get_record の責務

## `list_records`

### Purpose

`list_records` は、decision / spec / investigation / requirement / work_item / task を扱う record index を構造化して返す query tool である。

目的は、Markdown 本文を読む前に候補 record を絞り込むことである。
単なる filesystem listing ではなく、ADR / investigation / requirement / work item / task の箇条書きmetadataまたは spec の YAML front matter から正規化した record metadata と H1 title を含む一覧を返す。

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
| `kind` | no | string | `decision` / `spec` / `investigation` / `requirement` / `work_item` / `task` で絞り込む |
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

`id_range` は以下の ID family に適用できる。

| family | effective `kind` | endpoint form | ordering |
|---|---|---|---|
| decision | `decision` | `ADR-NNN` | `NNN` の数値比較 |
| requirement | `requirement` | `REQ-<DOMAIN>-NNN` | 同一 `<DOMAIN>` 内の `NNN` 数値比較 |
| work item | `work_item` | `WORK-<DOMAIN>-NNN` | 同一 `<DOMAIN>` 内の `NNN` 数値比較 |
| task | `task` | `TASK-<DOMAIN>-NNN-MM` | 同一 `<DOMAIN>` かつ同一 work sequence `NNN` 内の task sequence `MM` 数値比較 |

`kind` が省略され、かつ `id_range` が指定された場合は、指定された endpoint family から effective `kind` を決める。
`kind` が指定されている場合、endpoint family は指定 `kind` と一致しなければならない。

Workflow artifact range は同一 family / 同一 domain の endpoint に限定する。
Task range はさらに同一 work sequence に限定し、`TASK-MCP-006-01` .. `TASK-MCP-007-05` のような work item をまたぐ task range は request error とする。

One-sided workflow range は、指定された endpoint の family / domain / work sequence scope 内で評価する。
たとえば `kind: work_item` と `id_range.from: WORK-DATA-004` は `WORK-DATA-*` の sequence `004` 以上を対象にする。

`SPEC-*` / `INV-*` range はこの版では扱わず、request error とする。
Mixed family、mixed domain、mixed task work sequence、malformed endpoint、および指定 `kind` と endpoint family の不一致は request error とし、lexical ordering や広い listing へ silent fallback してはならない。

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
`id_range` 指定時の range membership は request section の ID family rule に従う。Response ordering は通常どおり `order_by` / `order` に従う。

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

`ids[]` は record index に対する exact lookup key としてのみ評価する。前後 whitespace の trim、case normalization、canonical reference resolution、input kind classification は行わない。したがって `spec:trace`、physical path、`adr-077`、` ADR-077 `、または grammar に合わない workflow ID のような string input が indexed record ID と一致しない場合、tool error や `unsupported` ではなく item-level `not_found` とする。`REQ-MCP-003` / `WORK-MCP-003` / `TASK-MCP-003-01` は index 済みの場合に found record を返す exact lookup key である。

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
          "depends_on": ["ADR-076", "ADR-077", "ADR-087", "ADR-088", "ADR-090", "ADR-092", "ADR-093"]
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

## `list_authoring_guides`

### Purpose

`list_authoring_guides` は、project authoring guidance の catalog を返す read-only tool である。

目的は、AI assistant が起票・更新対象に応じて必要な authoring guide を guide ID で発見できるようにすることである。

この tool は Design Records record index を返さない。Guide は record kind ではなく、authoring guidance retrieval surface として扱う。

### Request

```json
{}
```

request field は持たない。未知 field を許容する compatibility contract は定義しない。実装は未知 field を含む request を `invalid_request` tool error としてよい。

### Response

```json
{
  "guides": [
    {
      "id": "adr-authoring",
      "title": "ADR Authoring Guide",
      "abstract": "ADR を起票・レビュー・更新するときの実践ルールを定める。ADR は設計判断の履歴を所有し、現行仕様本文や作業 checklist を所有しない。"
    }
  ]
}
```

`guides[]` fields:

| field | required | meaning |
|---|---:|---|
| `id` | yes | guide ID。`docs/guides/<id>.md` の filename stem から導出する |
| `title` | yes | guide file の first H1 text |
| `abstract` | yes | guide file の `## Abstract` section content |

Response MUST NOT expose guide source file path.

`guides[]` の並び順は `id` の ASCII lexical order とする。

## `get_authoring_guidance`

### Purpose

`get_authoring_guidance` は、guide ID から authoring guidance Markdown を取得する read-only tool である。

この tool は Markdown guidance をそのまま返し、record metadata、record path、record headings、record lifecycle status を返さない。

### Request

```json
{
  "id": "adr-authoring"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `id` | yes | string | 取得対象 guide ID |

`id` は exact guide ID lookup key として扱う。前後 whitespace の trim、case normalization、physical path lookup、record ID resolution は行わない。

### Response

```json
{
  "id": "adr-authoring",
  "title": "ADR Authoring Guide",
  "content": "# ADR Authoring Guide\n\n## Abstract\n\n..."
}
```

| field | required | meaning |
|---|---:|---|
| `id` | yes | request で指定された guide ID |
| `title` | yes | guide file の first H1 text |
| `content` | yes | guide file の Markdown content |

`content` は元 Markdown をそのまま返す。整形・要約・正規化・truncate を行ってはならない。

Response MUST NOT expose guide source file path.

指定された guide ID が存在しない場合、tool は machine-readable な `guide_not_found` error を返す。

例:

```json
{
  "error": {
    "code": "guide_not_found",
    "message": "authoring guide unknown-guide was not found"
  }
}
```

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
| `REQ-<DOMAIN>-NNN` | `record_id` | `requirement` record index |
| `WORK-<DOMAIN>-NNN` | `record_id` | `work_item` record index |
| `TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>` | `record_id` | `task` record index |

`internal-design:` / `coverage:`、`COV-*`、physical path、および grammar に合わない ID form は supported input ではなく、direct query では tool execution error ではなく `status: "unsupported"` を返す。Workflow ID grammar として、requirement / work item sequence および task の work sequence は3桁ゼロ埋め、task sequence は2桁ゼロ埋めとする。`yaml:` は reserved prefix だが、MVP は public resolver input または direct query response behavior を定義しない。

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

Direct query の `unsupported_reference` は resolver の failure ではなく input boundary の可視化であるため `info` とする。ただし unsupported input が investigation metadata の validation 対象 field に現れた場合の severity は、下記 `validate_records` の `unsupported_reference` contract に従う。Reserved `yaml:` については、public resolver input / direct query response behavior、および investigation metadata に現れた場合の validation behavior を MVP では定義しない。

## `validate_records`
### Purpose

`validate_records` は、Design Records MCP の metadata index が信頼できる状態かを検証する tool である。

record metadata の基本整合性検査に加え、active `spec:` semantic ref、record ID-as-ref (`ADR-*` / `SPEC-*` / `INV-*` / `REQ-*` / `WORK-*` / `TASK-*`)、investigation の `source_refs` / 記載済み `follow_up_results` が canonical reference として解決可能であること、および workflow relation field の宣言済み integrity を検査する。
`follow_up_candidates` に artifact reference が記載された場合は canonical form を検査する。Canonical form の unresolved candidate は予定された後続 artifact が未作成であることを示す `info` diagnostic とし、physical path による candidate は noncanonical candidate を示す `info` diagnostic とする。

Workflow artifact については、metadata / relation validation に加えて、status-gated required narrative section validation を行う。
この検査は section heading の存在と section body の non-empty 条件のみを扱い、本文の品質・十分性・意味的妥当性は判定しない。

ADR-088 / ADR-092 により、`internal-design:` / `coverage:` / `COV-*`、semantic realization relation、coverage mapping query、orphan workflow artifact diagnostics、progress projection、workflow traversal query は MVP required scope として扱わない。Investigation metadata の canonical workflow reference 拡張は `REQ-*` / `WORK-*` に限定し、同 field 内の `TASK-*` は unsupported reference として扱う。

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

`id_range` の endpoint family、effective `kind`、one-sided range、unsupported range の扱いは `list_records` と同じとする。
したがって `validate_records` でも `REQ-<DOMAIN>-NNN` / `WORK-<DOMAIN>-NNN` / `TASK-<DOMAIN>-NNN-MM` の safe workflow artifact range を指定できる。
`SPEC-*` / `INV-*` range、mixed family、mixed domain、mixed task work sequence、malformed endpoint、および指定 `kind` と endpoint family の不一致は request error とする。

### Response

```json
{
  "ok": false,
  "diagnostics": [
    {
      "category": "missing_required_section",
      "severity": "error",
      "record_id": "WORK-MCP-014",
      "path": "docs/work-items/mcp/WORK-MCP-014-normalize-propose-record-create-id-fields-body-contract.md",
      "message": "required section \"Boundary\" must be non-empty when work_item status is \"done\"",
      "section": "Boundary",
      "status": "done"
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
| `severity` | yes | `error`, `warning`, または `info` |
| `record_id` | no | 問題がある record ID |
| `path` | no | 問題がある path |
| `message` | yes | human-readable message |
| `target_id` | no | 参照切れなどの対象 ID |

`error` severity は `validation.ok` を `false` にし、retained proposal の作成をブロックする。
`warning` severity は retained proposal 上で返されるが、`validation.ok` を `false` にせず、proposal creation をブロックしない。
`info` severity は `validation.ok` を `false` にしない。

### Diagnostic categories
MVP diagnostic category は `schema.md` の定義に従う。

- `duplicate_id`
- `filename_id_mismatch`
- `invalid_h1_title`
- `invalid_workflow_id`
- `missing_required_metadata`
- `empty_required_metadata`
- `missing_required_section`
- `empty_required_section`
- `section_heading_case_mismatch`
- `invalid_metadata_value`
- `missing_required_metadata_batch`
- `invalid_status_for_kind`
- `spec_status_mismatch`
- `missing_depends_on_target`
- `missing_supersedes_target`
- `invalid_migrated_to_spec`
- `missing_record_path`
- `no_op_update`
- `reciprocal_follow_up_mode_required`
- `conflicting_operations`
- `multiple_section_replace_not_supported`

Canonical reference / investigation / workflow validation の concrete category と severity は以下とする。

| category | severity | field / condition |
|---|---|---|
| `missing_required_metadata` | `error` | workflow artifact の required metadata field が存在しない |
| `missing_required_metadata_batch` | `error` | authoring create 時に複数の required metadata field が欠落している。`required_fields` に欠落フィールド名一覧、`target_kind` に対象 kind を含める |
| `empty_required_metadata` | `error` | workflow artifact の required scalar metadata field が empty、または required list metadata field に empty item が含まれる |
| `missing_required_section` | `error` | workflow artifact が gated status にあるとき、required narrative section heading が存在しない |
| `empty_required_section` | `error` | workflow artifact が gated status にあるとき、required narrative section heading は存在するが section body が empty または whitespace-only である |
| `section_heading_case_mismatch` | `info` | workflow artifact の target-kind-specific validation-required narrative section heading と case-only で一致する non-canonical heading が存在する |
| `invalid_metadata_value` | `error` | workflow artifact の required metadata field が non-empty だが value contract を満たさない。例: `date` が strict `YYYY-MM-DD` format ではない、または status field の値が対象 kind の allowed set に含まれない。後者の場合は `allowed_values` に許容値一覧を含め、`repair_suggestion` に最小修正の metadata patch を含めてよい |
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
| `unsupported_reference` | `error` / `info` | MVP が unsupported と定義する investigation metadata reference。`source_refs` / `follow_up_results` では `error`、`follow_up_candidates` では `info`。Investigation field に現れる `TASK-*` を含む。Reserved `yaml:` はこの category の対象に含めず、MVP では behavior を定義しない |
| `unresolved_workflow_relation` | `error` | workflow relation field に記載された supported `REQ-*` / `WORK-*` / `TASK-*` が解決不能 |
| `invalid_workflow_relation_target` | `error` | workflow relation field に field が要求する kind / ID form ではない target が記載された |
| `workflow_relation_mismatch` | `error` | `REQ.work_items` と `WORK.source_requirement`、または `WORK.tasks` と `TASK.work_item` の宣言済み双方向 relation が一致しない |
| `workflow_source_requirement_mismatch` | `error` | task の `source_requirement` が parent work item の `source_requirement` と一致しない |

Workflow metadata diagnostic (`missing_required_metadata` / `empty_required_metadata` / `invalid_metadata_value`) は、既存の diagnostic field に加えて `field` を必須で返す。入力 value が存在する場合は `value` も返す。

Workflow required section diagnostic (`missing_required_section` / `empty_required_section`) は、通常の `category` / `severity` / `record_id` / `path` / `message` に加えて `section` と `status` を必須で返す。
`section` は required narrative section heading text とする。
`status` は required section rule を発火させた workflow artifact status とする。

When validation detects that a canonical required section is missing for the target record kind and current gated status, and exactly one heading exists whose text differs only by case from that canonical heading, validation MUST also return `section_heading_case_mismatch` with severity `info`.
This diagnostic is repair guidance only. It does not relax canonical required-section validation and does not suppress any `missing_required_section` / `empty_required_section` error.

`section_heading_case_mismatch` MUST include:

- `section`: canonical required heading text for the target record kind.
- `actual_heading`: matched non-canonical heading text.
- `status`: workflow artifact status that activated the required-section rule.

It SHOULD include `candidate_headings` with heading text, level, and ordinal when available.

Workflow relation diagnostic (`unresolved_workflow_relation` / `invalid_workflow_relation_target` / `workflow_relation_mismatch` / `workflow_source_requirement_mismatch`) は、既存の diagnostic field に加えて `field`（`work_items` / `source_requirement` / `tasks` / `work_item` / `depends_on`）、`value`（入力 ID-as-ref）、`ref_status`（`unresolved` / `invalid_target` / `mismatch`）を必須で返す。

Investigation reference diagnostic (`unresolved_*` / `noncanonical_*` / metadata field 由来の `unsupported_reference`) は、既存の diagnostic field に加えて `field`（`source_refs` / `follow_up_results` / `follow_up_candidates`）、`value`（入力 ref 文字列）、`ref_status`（`unresolved` / `unsupported` / `noncanonical`）を必須で返す。対象が record ID-as-ref の場合は `target_id` も返してよい。Investigation metadata が duplicate semantic ref または duplicate record ID を指して単一解決できない場合は field-specific diagnostic を追加せず、index defect を示す `duplicate_semantic_ref` または `duplicate_id` のみを返す。これら duplicate diagnostic および spec declaration / section lookup diagnostic は investigation metadata field 由来の追加 field を要求しない。

Required narrative section policy:

| artifact kind | gated status | required non-empty narrative sections |
|---|---|---|
| `work_item` | `done` | `Goal`, `Boundary`, `Evidence` |
| `task` | `done` | `Goal`, `Work`, `Done condition`, `Verification`, `Evidence` |
| `requirement` | `accepted` | `Requirement`, `Required Outcome` |

Only headings listed for the target record kind in the Required narrative section policy are canonical workflow required headings for the case-only repair behavior defined by `propose_record_update` named section replacement.
The target record does not need to currently be in the gated status for the authoring selector fallback to apply; the target record kind and requested heading determine fallback eligibility.
Authoring guide format headings that are not listed for the target record kind, and user-defined optional headings, are not canonicalized by this rule.

`requirement` の `accepted` は close/completion state ではなく adoption-readiness gate として扱う。
したがって `Evidence` / `Boundary` / `Explicitly Excluded Scope` は `REQ accepted` の required non-empty section には含めない。

Required narrative section body は、heading 行を除いた section body の前後 whitespace を trim した結果、少なくとも 1 つの non-whitespace character を含む場合に non-empty とする。
Whitespace-only body は empty とする。
本文の品質・十分性・意味内容は判定しないため、`Pending` や `None` のような placeholder text も non-empty として扱う。

`depends_on` は `ADR-*` / `SPEC-*` / `INV-*` の canonical record ID-as-ref を参照できる。したがって `ADR-086` の `depends_on: INV-DOCS-001` は contract 上 valid であり、investigation integration 実装前に返る `missing_depends_on_target` は既知 implementation gap である。M19 implementation 後は `INV-DOCS-001` の index integration によりこの diagnostic が解消されなければならない。

`ok` は `error` diagnostic がない場合に `true` とし、`warning` および `info` diagnostic が存在しても `false` にしない。
Coverage mapping、semantic realization relation、`internal-design:` / `coverage:` / `COV-*` の解決・診断は MVP tool acceptance に含めない。Workflow relation validation は宣言済み relation の存在と整合性に限り、未接続 artifact の orphan diagnostics、task dependency cycle detection、execution order projection、task status 由来 progress projection は含めない。

`accepted_but_not_migrated` / `missing_design_record` などの運用 gap 診断、および orphan requirement / orphan work item / orphan task diagnostic は MVP 外である。

`missing_record_path` は、filesystem scan または path normalization により record 候補 path を検出したが、実際の read/stat に失敗した場合に出す。
例として、scan 後に file が削除された場合、permission denied、symlink target missing、path normalization 後の path が存在しない場合を含む。

`no_op_update` は `propose_record_update` 時に提案内容がファイルの現在の内容と byte 単位で一致した場合に出す。Severity は `info`。`proposal_created: false` を伴い、retained proposal は作成されない。`record_id` と `path` を含める。

`reciprocal_follow_up_mode_required` は `propose_record_create` を `reciprocal_update_mode: "report_required_follow_up"` で呼び出し、かつ必須の reciprocal follow-up update が存在するときに出す。Severity は `warning`。`include_required` モードが accept に必要であることを `message` で明示する。Repair guidance として `repair_suggestion` に `{"reciprocal_update_mode": "include_required"}` を含めてよい。

`conflicting_operations` は `propose_record_update` の `operations` 配列内に同一 metadata field を対象とする複数の metadata operation が含まれる場合に出す（`metadata_fields_replace` 同士の `metadata` key 重複、または `metadata_block_replace` と他の metadata operation の共存）。Severity は `error`。`proposal_created: false` を伴い、retained proposal は作成されない。

`multiple_section_replace_not_supported` は `propose_record_update` の `operations` 配列内に 2 つ以上の `named_section_replace` operation が含まれる場合に出す。MVP 制約として、対象セクションが同一か否かに関わらず適用される。Severity は `error`。`proposal_created: false` を伴い、retained proposal は作成されない。

Authoring diagnostic additional fields:

| field | applicable categories | meaning |
|---|---|---|
| `allowed_values` | `invalid_metadata_value` (status) | 対象 kind で許容される status 値の一覧 |
| `required_fields` | `missing_required_metadata_batch` | 欠落している required field 名の一覧 |
| `target_kind` | `missing_required_metadata_batch` | フィールド要件が定義されている record kind |
| `repair_suggestion` | `invalid_metadata_value`, `reciprocal_follow_up_mode_required` | 最小修正を示す advisory metadata patch。実装が決定論的に導出できる場合のみ含める。Caller は自動適用せず、必ず確認してから使用する |

`repair_suggestion` の value は最小修正を示す plain object（例: `{"status": "not_started"}`）とする。非決定論的な修正（ユーザーのビジネス意図が必要なもの）には含めない。`repair_suggestion` は advisory であり、validation や accept guard を迂回しない。

`missing_required_metadata_batch` の diagnostic 例:

```json
{
  "category": "missing_required_metadata_batch",
  "severity": "error",
  "message": "fields is missing required metadata fields for kind task: work_item, estimate, outputs",
  "required_fields": ["work_item", "estimate", "outputs"],
  "target_kind": "task"
}
```

`invalid_metadata_value` (status) with `allowed_values` の diagnostic 例:

```json
{
  "category": "invalid_metadata_value",
  "severity": "error",
  "field": "status",
  "value": "todo",
  "message": "status \"todo\" is not valid for kind task",
  "allowed_values": ["not_started", "in_progress", "blocked", "done"],
  "repair_suggestion": {
    "status": "not_started"
  }
}
```

`reciprocal_follow_up_mode_required` の diagnostic 例:

```json
{
  "category": "reciprocal_follow_up_mode_required",
  "severity": "warning",
  "message": "required reciprocal follow-up updates are present; use reciprocal_update_mode: \"include_required\" for a safe accept",
  "repair_suggestion": {
    "reciprocal_update_mode": "include_required"
  }
}
```

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

## Authoring transaction model

Authoring transaction tool は artifact-oriented write surface である。
Physical filesystem path は public request の primary input として受け取らない。
Tool input は record kind、record ID または `new` placeholder ID、domain、parent context、section selector、structured authoring fields、および必要時の body source に基づく。
`path` は response で透明性のために返してよいが、caller が write target を path で直接指定する contract ではない。

Authoring write は propose -> accept flow とする。

1. `propose_record_create` または `propose_record_update` は、変更案を作成し、proposal ID、resolved target、previewable diff、validation result、expiry、diagnostics、note を返す。
2. Proposal creation は repository file を変更してはならない。
3. `get_proposed_write` は retained proposal の内容と lifecycle state を返す。
4. `discard_proposed_write` は proposal を discard し、以後 accept できない状態にする。
5. `accept_proposed_write` は proposal ID を受け取り、accept-time validation / staleness / collision guard を通過した場合だけ repository file を書き込む。

Proposal lifecycle state は `proposed` / `accepted` / `discarded` とする。
Expired proposal は retained proposal として返さず、`proposal_expired` diagnostic を返す。
Proposal retention は 3 days とし、proposal response は `expires_at` を必ず返す。

Proposal は base-state information を保持しなければならない。
Accept は write の直前に少なくとも以下を再確認する。

- proposal が unknown / expired / discarded / already accepted ではないこと
- target file state が proposal 作成時から変わっていないこと
- target kind が proposal 作成時の resolved target と一致していること
- create proposal の resolved ID がまだ使用可能であること
- update proposal の target ID がまだ同一 record として解決できること
- pre-write validation に error diagnostic がないこと

上記のいずれかに失敗した場合、`accept_proposed_write` は `written: false` を返し、repository file を変更してはならない。

### Common authoring response fields

Proposal response は少なくとも以下を返す。

| field | required | meaning |
|---|---:|---|
| `proposal_id` | yes | proposal lookup key。Format は opaque string とし、caller は構造を解釈しない |
| `state` | yes | `proposed` |
| `operation` | yes | `create` / `update` |
| `target_kind` | yes | resolved target record kind |
| `target` | yes | requested / resolved target identity object |
| `expires_at` | yes | proposal expiry timestamp |
| `retention_days` | yes | `3` |
| `diff` | yes | previewable diff object |
| `validation` | yes | validation result object |
| `diagnostics` | yes | request / proposal diagnostic list |
| `note` | yes | repository file がまだ書かれていないことと、適用には accept が必要であることを明示する note |

`target` は少なくとも以下を返す。

| field | required | meaning |
|---|---:|---|
| `requested_id` | yes | caller が指定した ID。`new` placeholder を含んでよい |
| `resolved_id` | yes | MCP が index から解決した final record ID |
| `kind` | yes | target record kind |
| `domain` | no | domain-scoped workflow record の domain |
| `parent_id` | no | task create など parent-aware resolution に使われた parent ID |
| `path` | yes | resolved repository-relative path。Transparency output only; request primary input ではない |

`diff` は少なくとも以下を返す。

| field | required | meaning |
|---|---:|---|
| `format` | yes | MVP では `unified` |
| `text` | yes | previewable unified diff text |
| `files` | yes | changed file summary list |

`diff.files[]` は `path`、`change` (`create` / `modify`)、および必要に応じて `record_id` / `record_kind` を含む。
Workflow reciprocal metadata update を含む proposal では、`files[]` は複数 entry になりうる。

`validation` は少なくとも `ok` と `diagnostics` を返す。
`ok` は error diagnostic がない場合に `true` とする。
Validation failure と write failure は同じ state として扱わない。

Proposal-time validation and accept-time pre-write validation are proposal-local.
`validation.diagnostics` must describe only diagnostics for the affected record set in the candidate state represented by the proposal.
Unrelated existing repository diagnostics must not be mixed into proposal-local blocking diagnostics.
If an implementation chooses to expose broader repository health while preparing or accepting a proposal, it must use a field separate from `validation`, such as `repository_health`, or an equivalently distinct response category.
Repository health diagnostics must not affect proposal-local `validation.ok` or accept-time write eligibility unless the diagnosed record is also in the affected record set.

Proposal note の標準意味は以下である。

```text
No repository files have been written. Call accept_proposed_write with this proposal_id to apply the diff.
```

Exact wording は implementation detail だが、この意味を弱めてはならない。

### Body source and body cache

Operations that require Markdown body input accept a caller-supplied `body` or, where explicitly supported, a cached `body_cache_id`.
Supplying both `body` and `body_cache_id` is invalid and must not create a proposal or body cache entry.
Supplying neither is valid only for operations that do not require body input.

For `propose_record_create`, structured `fields` and caller-supplied `body` may be combined when `body` is content sections only.
In this mode, `fields` render the H1-following metadata block and `body` supplies the Markdown sections after that generated metadata block.
The MCP owns H1 and metadata rendering for structured creates.
The caller must not include the H1, metadata block, metadata `id`, or any server-resolved ID in the `body`.

`propose_record_create` uses a strict structured create contract: `fields` is required.
`body` and `body_cache_id` are optional section-only content sources and are valid only when combined with `fields`.
Legacy full-record create body input without `fields` is invalid.

MVP body source rules:

| operation case | body requirement |
|---|---|
| `propose_record_create` structured metadata only | `fields` present; `body` / `body_cache_id` omitted |
| `propose_record_create` structured metadata plus content sections | `fields` and section-only `body` present; `body_cache_id` omitted |
| `propose_record_create` retry with cached content sections | `fields` and `body_cache_id` present; `body` omitted |
| `propose_record_create` body-only / cache-only create | invalid; `fields` is required |
| `propose_record_update` `metadata_block_replace` | `body` / `body_cache_id` must be omitted |
| `propose_record_update` `metadata_fields_replace` | `body` / `body_cache_id` must be omitted |
| `propose_record_update` `named_section_replace` | exactly one of `body` / `body_cache_id` |

If a request supplies both `body` and `body_cache_id`, the tool returns `invalid_body_source` and creates neither proposal nor cache.
If `body_cache_id` is unknown or expired, the tool returns `body_cache_not_found` or `body_cache_expired` and creates no proposal.

If a request supplies a large `body` and proposal/write preparation fails before the body can be persisted into a proposal, the MCP should preserve the submitted body when possible and return a retryable body cache object:

```json
{
  "proposal_created": false,
  "body_cache": {
    "body_cache_id": "bc_opaque",
    "expires_at": "2026-06-05T00:00:00Z",
    "retention_days": 3
  },
  "diagnostics": [
    {
      "category": "proposal_preparation_failed",
      "severity": "error",
      "message": "proposal could not be prepared; retry with body_cache_id"
    }
  ]
}
```

Body cache retention is 3 days.
Body cache responses must include `expires_at`.
Body cache entries remain reusable within the 3 day retention period, including after they have been used to create a proposal.
Expired body cache entries must not be used to create proposals.

### Proposal validation affected record set

The proposal-local affected record set is the set of records whose content is created or modified by the proposal.

For `propose_record_create`, the affected record set contains:

- the proposed target record; and
- any related records that the proposal actually modifies for required reciprocal workflow metadata under `reciprocal_update_mode`, such as the parent requirement or work item updated by `include_required`.

For `propose_record_update`, the affected record set contains the target record being updated.
If a future update operation explicitly modifies related records in the same retained proposal, those related records become part of the affected record set.

Proposal-time validation must run against the candidate repository state represented by applying the proposal diff to the current repository state, but returned proposal-local diagnostics must be filtered to the affected record set.
Accept-time pre-write validation uses the same affected-record-set model after staleness, target-change, and ID-collision guards.
Diagnostics returned in `validation.diagnostics` must be reproducible by running the same `validate_records` rules against the same affected record set in the same candidate state, or after accepting/materializing that candidate state.
This contract does not suppress diagnostics for affected related records; it only prevents unrelated repository diagnostics from becoming proposal-local blockers.

> 由来: REQ-MCP-012, TASK-MCP-011-01

## `propose_record_create`

### Purpose

`propose_record_create` creates a retained proposal for a new record or workflow artifact.
It does not write repository files.

MVP create support covers:

- `decision`
- `requirement`
- `work_item`
- `task`

Spec skeleton creation is outside this MVP authoring surface. Existing spec records may still be updated by `propose_record_update` when the target record already exists and the metadata / section selector is unambiguous.
Investigation creation is outside this MVP authoring surface.

### Request

```json
{
  "kind": "task",
  "id": "TASK-MCP-008-new",
  "domain": "MCP",
  "parent_id": "WORK-MCP-008",
  "title": "MCP tools spec reflection",
  "fields": {
    "status": "not_started",
    "date": "2026-06-01",
    "source_requirement": "REQ-MCP-008",
    "estimate": "1d-2d",
    "depends_on": ["TASK-MCP-008-03"],
    "outputs": ["Updated SPEC-design-records-mcp-tools"]
  },
  "body": "## Goal\n\nReflect the accepted authoring transaction contract in the Design Records MCP tools spec.\n\n## Work\n\n- Update the public tool contract.\n- Record verification evidence.\n\n## Done condition\n\nThe spec documents the current request and response contract.\n\n## Verification\n\n- go test ./internal/designrecords ./internal/designrecordsmcp\n\n## Evidence\n",
  "reciprocal_update_mode": "include_required"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `kind` | yes | string | create target kind |
| `id` | yes | string | exact ID or allowed `new` placeholder ID |
| `domain` | conditional | string | domain-scoped workflow create の domain。ID に domain が含まれる場合は一致必須 |
| `parent_id` | conditional | string | task create では required。Parent work item ID |
| `title` | yes | string | H1 title |
| `fields` | yes | object | kind-specific structured authoring fields。MCP が H1 / metadata を生成するため必須 |
| `body` | conditional | string | caller-supplied section-only content body。`fields` と組み合わせる場合のみ valid |
| `body_cache_id` | conditional | string | cached section-only body lookup key for `fields + body` retry form |
| `reciprocal_update_mode` | no | string | workflow reciprocal metadata handling mode |

`propose_record_create` has these content input combinations:

| mode | required content input | forbidden content input |
|---|---|---|
| structured metadata only | `fields` | `body`, `body_cache_id` |
| structured metadata plus content sections | `fields` plus section-only `body` | `body_cache_id` |
| retry with cached content sections | `fields` plus `body_cache_id` | `body` |
| body-only create | none | invalid because `fields` is required |
| cache-only create | none | invalid because `fields` is required |

In all create modes, top-level `kind`, `id`, `domain` when applicable, `parent_id` when applicable, and `title` remain request-level target inputs.
For structured create, `title` is also a rendering input.
The top-level `id` is the canonical create target ID input.
For exact create IDs, it is the canonical target ID after request normalization.
For `new` placeholders, it is the canonical target family and server-side resolution request, and the response `target.resolved_id` is the final canonical target ID.

When `fields` is present, the MCP renders the record H1 and metadata block from top-level `id`, `title`, server-resolved target identity, and `fields`.
If `body` is also present, `body` is appended after the generated metadata block as content sections only.
The `body` must start at the first content section, such as `## Goal`, `## Requirement`, or `## 背景`, and must not include a leading H1, bullet metadata block, YAML metadata, metadata `id`, or guessed resolved ID.
For `id` placeholders such as `TASK-MCP-008-new`, the generated H1 and metadata must use `target.resolved_id`, not the literal `new` placeholder.

Full-record body create without `fields` is invalid.
Callers must not submit H1, metadata block, metadata `id`, or guessed server-resolved ID through `body` / `body_cache_id`.
If a caller submits `body` without `fields`, the request is rejected as `invalid_request`; when the submitted `body` is a string, the MCP should return a new `body_cache` so the caller can retry without regenerating the content.

`fields.id` is not required for create.
Structured create rendering must use `target.resolved_id` as the record metadata ID and must not require callers to duplicate top-level `id` inside `fields`.
If `fields.id` is supplied with an exact top-level ID, it must match the top-level ID after the same canonical ID normalization used by the request.
If `fields.id` is supplied with a `new` placeholder top-level ID, the request is invalid because the final ID is not known to the caller before server-side resolution.
If `fields.id` does not match the exact top-level ID, the tool rejects the request as `invalid_request` and creates no proposal.

For domain-scoped workflow creates, `domain` is compared with the ID domain case-insensitively.
The canonical ID domain remains the uppercase domain segment in IDs such as `REQ-MCP-011`.
The response `target.domain` uses the canonical ID domain.
Repository paths use the lowercase normalized domain directory, such as `docs/requirements/mcp/`.
Therefore `domain: "mcp"` and `id: "REQ-MCP-011"` are consistent, while `domain: "data"` and `id: "REQ-MCP-011"` are rejected as `invalid_request`.

> 由来: REQ-MCP-011, TASK-MCP-011-01

Allowed `id` placeholder forms:

| kind | allowed create ID forms |
|---|---|
| `decision` | exact `ADR-NNN` or `ADR-new` |
| `spec` | not supported for create in MVP |
| `requirement` | exact `REQ-<DOMAIN>-NNN` or `REQ-<DOMAIN>-new` |
| `work_item` | exact `WORK-<DOMAIN>-NNN` or `WORK-<DOMAIN>-new` |
| `task` | exact `TASK-<DOMAIN>-<WORK-SEQUENCE>-NN` or `TASK-<DOMAIN>-<WORK-SEQUENCE>-new` |

The `new` placeholder is the literal token `new` in the sequence position. `ADR-new`, `REQ-<DOMAIN>-new`, `WORK-<DOMAIN>-new`, and `TASK-<DOMAIN>-<WORK-SEQUENCE>-new` are the only accepted placeholder forms. Any other token in the sequence position, such as `ADR-newish` or `REQ-MCP-newer`, is treated as malformed and rejected as `invalid_request`.

`new` is valid only for create operations.
The MCP resolves the final ID using the current record index.
`SPEC-new` and spec skeleton create are rejected in the MVP because spec record placement cannot be derived safely from ID alone. Spec placement discovery / domain tree support is tracked as REQ-MCP-010.
For numeric families, the MCP uses the next available sequence in the relevant family / domain / parent scope and does not fill gaps unless a later spec explicitly changes that rule.

Exact ID create remains allowed for workflow artifacts, but it may return a non-blocking authoring diagnostic when the requested exact ID would create a sequence gap. This warning is caller feedback only; it must not reject proposal creation or acceptance by itself.

The diagnostic category is `exact_id_sequence_gap`, and severity is `info`. The diagnostic should recommend the matching `*-new` placeholder when the caller does not intentionally need the exact ID.
This diagnostic is returned in the proposal-level `diagnostics` field, not `validation.diagnostics`, because it is a proposal advisory rather than a record content validation result.

Exact ID gap-warning scopes are:

| kind | sequence scope | warning condition |
|---|---|---|
| `requirement` | same requirement kind and domain | requested `NNN` is greater than current max `NNN` plus one |
| `work_item` | same work item kind and domain | requested `NNN` is greater than current max `NNN` plus one |
| `task` | same domain and parent work item sequence | requested task `NN` is greater than current max task `NN` plus one |

`decision` / `ADR-*` create is outside this workflow artifact warning scope.
`new` placeholder create must not emit `exact_id_sequence_gap`, because server-side allocation already uses the current max sequence plus one.
Exact ID create that fills an existing gap must not emit `exact_id_sequence_gap`; existing duplicate ID checks still reject an exact ID that already exists.

Task create requires `parent_id`.
The parent work item must resolve to an indexed `work_item` record.
For task placeholder IDs, `<DOMAIN>` and `<WORK-SEQUENCE>` must match the parent work item ID.
Task parent relation must be written from explicit metadata, not inferred from ID shape alone.

`reciprocal_update_mode` values:

| value | meaning |
|---|---|
| `include_required` | Default. Include required reciprocal workflow metadata updates in the same proposal when needed to keep relation validation valid |
| `report_required_follow_up` | Do not include reciprocal file updates; return explicit required follow-up updates and reject acceptance until they are represented by an accepted proposal or current index state |

`include_required` may create a multi-file proposal for workflow relation validity.
This is allowed only for required reciprocal metadata updates such as adding a new work item to `REQ.work_items` or a new task to `WORK.tasks`.
It is not a general-purpose multi-record atomic transaction with rollback semantics, and arbitrary unrelated record bundling remains outside MVP.

### Response

```json
{
  "proposal_id": "pw_opaque",
  "state": "proposed",
  "operation": "create",
  "target_kind": "task",
  "target": {
    "requested_id": "TASK-MCP-008-new",
    "resolved_id": "TASK-MCP-008-04",
    "kind": "task",
    "domain": "MCP",
    "parent_id": "WORK-MCP-008",
    "path": "docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md"
  },
  "expires_at": "2026-06-05T00:00:00Z",
  "retention_days": 3,
  "diff": {
    "format": "unified",
    "files": [
      {
        "path": "docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md",
        "change": "create",
        "record_id": "TASK-MCP-008-04",
        "record_kind": "task"
      }
    ],
    "text": "--- /dev/null\n+++ docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md\n..."
  },
  "validation": {
    "ok": true,
    "diagnostics": []
  },
  "diagnostics": [],
  "note": "No repository files have been written. Call accept_proposed_write with this proposal_id to apply the diff."
}
```

The proposal response may include `required_follow_up_updates` when `reciprocal_update_mode: "report_required_follow_up"` is used or when the implementation cannot include a required reciprocal update in the same proposal.
If required follow-up updates are present, acceptance must be rejected with `written: false` until the follow-up requirement is satisfied.

Exact ID sequence-gap warning example:

```json
{
  "diagnostics": [
    {
      "category": "exact_id_sequence_gap",
      "severity": "info",
      "message": "REQ-MCP-020 skips the next available sequence REQ-MCP-019; prefer REQ-MCP-new unless this ID is intentional"
    }
  ]
}
```

This example shows only the relevant response field. The full proposal response still includes the common proposal response fields.

## `propose_record_update`

### Purpose

`propose_record_update` creates a retained proposal for a single-operation update (whole metadata block replacement, field-level metadata patch, or named Markdown section replacement) or an atomic multi-operation update combining multiple supported operations into one retained proposal for the same record.
It does not write repository files.

If the requested update operation is a no-op after operation semantics are applied, `propose_record_update` MUST NOT create a retained proposal.
A no-op update is an update request whose proposed persisted content is byte-equivalent to the current persisted file content after the requested update operation semantics are applied.

MVP update support covers `decision`, `spec`, `requirement`, `work_item`, and `task`.
Update operations reject any ID containing the literal `new` token in the sequence position.

### Request

Metadata block replacement:

```json
{
  "kind": "task",
  "id": "TASK-MCP-008-04",
  "update": {
    "type": "metadata_block_replace",
    "metadata": {
      "id": "TASK-MCP-008-04",
      "status": "done",
      "date": "2026-06-01",
      "work_item": "WORK-MCP-008",
      "source_requirement": "REQ-MCP-008",
      "estimate": "1d-2d",
      "depends_on": ["TASK-MCP-008-03"],
      "outputs": ["Updated SPEC-design-records-mcp-tools"]
    }
  }
}
```

Named section replacement:

```json
{
  "kind": "task",
  "id": "TASK-MCP-008-04",
  "update": {
    "type": "named_section_replace",
    "section_selector": {
      "heading": "Evidence",
      "match": "exact"
    }
  },
  "body": "2026-06-02: Spec reflection completed.\n"
}
```

Metadata field replacement:

```json
{
  "kind": "task",
  "id": "TASK-MCP-020-02",
  "update": {
    "type": "metadata_fields_replace",
    "metadata": {
      "status": "done"
    }
  }
}
```

Operations array:

```json
{
  "kind": "task",
  "id": "TASK-MCP-008-04",
  "operations": [
    {
      "type": "metadata_fields_replace",
      "metadata": { "status": "done" }
    },
    {
      "type": "named_section_replace",
      "section_selector": { "heading": "Evidence", "match": "exact" },
      "body": "2026-06-07: Implementation verified.\n"
    }
  ]
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `kind` | yes | string | update target kind |
| `id` | yes | string | exact existing record ID. `new` placeholder is invalid |
| `update` | conditional | object | single update operation object; mutually exclusive with `operations` |
| `operations` | conditional | array | atomic multi-operation list; mutually exclusive with `update` |
| `body` | conditional | string | replacement Markdown body for `named_section_replace` when using `update` |
| `body_cache_id` | conditional | string | cached replacement body for `named_section_replace` when using `update` |

Exactly one of `update` or `operations` must be present.

`update.type` values:

| value | meaning |
|---|---|
| `metadata_block_replace` | Replace the kind-specific metadata block as a whole |
| `metadata_fields_replace` | Patch one or more metadata fields; preserve unspecified existing fields |
| `named_section_replace` | Replace exactly one Markdown section as a whole |

#### Metadata block replacement

Metadata block replacement targets the kind-specific metadata block.

| kind | replacement target |
|---|---|
| `spec` | recognized spec metadata fields inside YAML front matter |
| `decision` | H1-following ADR bullet metadata block |
| `requirement` | H1-following requirement bullet metadata block |
| `work_item` | H1-following work item bullet metadata block |
| `task` | H1-following task bullet metadata block |

For `spec`, metadata replacement is scoped to recognized fields only. Unknown or auxiliary YAML front matter fields must be preserved. The recognized spec metadata fields are `scope`, top-level `status`, and `design_record.id` / `design_record.kind` / `design_record.status` / `design_record.depends_on`.
Decision metadata replacement must validate recognized ADR fields required by the current ADR metadata contract: `status`, `date`, `depends_on`, `supersedes`, and `migrated_to_spec`.
Workflow artifact metadata replacement must validate the required fields defined in `schema.md`.

Missing required fields must produce `missing_required_metadata` diagnostics.
Empty required scalar fields or empty list items must produce `empty_required_metadata`.
Invalid recognized field values must produce `invalid_metadata_value` or the existing kind-specific diagnostic.

#### Metadata field replacement

Metadata field replacement reads the existing metadata of the target record, applies only the caller-supplied field changes, and preserves all unspecified existing metadata fields.

`update.metadata` must contain only the fields to be patched. Fields omitted from `update.metadata` are kept from the current record as-is.

After applying the patch, the resulting complete metadata block is validated using the same metadata validation rules as `metadata_block_replace`:

- Missing required fields produce `missing_required_metadata` diagnostics.
- Empty required scalar fields or empty list items produce `empty_required_metadata`.
- Invalid recognized field values produce `invalid_metadata_value` or the existing kind-specific diagnostic.

`body` and `body_cache_id` must be omitted for `metadata_fields_replace`. Supplying either returns `invalid_body_source` and must not create a proposal.

`metadata_block_replace` remains available for intentional whole-block replacement when the caller needs to provide or reset the complete metadata set.

#### Named section replacement
Named section replacement is valid only when `section_selector` resolves to exactly one Markdown ATX section in the target record. Section matching uses the same ATX heading source rules as the `headings` field defined in `schema.md`; YAML front matter and fenced code block content are not section sources, and setext headings are not section sources in the MVP.

`section_selector` fields:

| field | required | meaning |
|---|---:|---|
| `heading` | yes | heading text to match |
| `match` | no | MVP supports `exact` only |
| `level` | no | optional ATX heading level constraint |

Exact matching compares the parsed heading text after removing ATX marker syntax and trimming surrounding whitespace.
Exact matching is case-sensitive.
No Unicode normalization, punctuation folding, or prefix / contains matching is applied in the MVP.
If `level` is supplied, both heading text and level must match.

Narrow case-only fallback:

- The fallback applies only to workflow artifact records (`requirement`, `work_item`, `task`).
- The fallback applies only when `section_selector.heading` is listed as a validation-required canonical heading for the target record kind in the Required narrative section policy.
- The target record does not need to currently be in the gated status for the fallback to apply; target record kind and requested heading determine fallback eligibility.
- Authoring guide format headings that are not validation-required for the target record kind, and user-defined optional headings, are not canonicalized by this fallback.
- The fallback is attempted only after exact matching finds zero matches.
- The fallback compares parsed heading text case-insensitively, without Unicode normalization, punctuation folding, typo correction, prefix matching, or contains matching.
- If `level` is supplied, fallback candidate matching uses the same level constraint before determining zero, one, or ambiguous case-insensitive matches.
- If exactly one eligible case-insensitive match is found, the selector resolves through fallback and proposal creation proceeds unless an independent proposal-preparation error applies.
- When proposal creation proceeds through this fallback, the retained proposal diff MUST rewrite the matched heading line to the canonical `section_selector.heading` text.
- If multiple case-insensitive matches are found, proposal creation MUST fail with `section_selector_ambiguous` and MUST NOT create a proposal.
- Non-case differences remain governed by the existing exact selector rules and MUST NOT use this fallback.

The replacement range includes the matched heading line and all following lines until the next heading whose level is less than or equal to the matched heading level.
Nested headings below the matched heading are part of the replaced section.

Zero matches must return `section_selector_no_match` and must not create a proposal.
Multiple matches must return `section_selector_ambiguous` and must not create a proposal.
When possible, diagnostics should include `candidate_headings` with heading text, level, and ordinal.

**Heading-safe replacement body normalization:**

This normalization is evaluated after `section_selector` resolves to exactly one target section. The comparison uses the resolved selected section heading text and resolved selected section heading level — not the raw `section_selector` input values. The strip condition therefore applies regardless of whether `section_selector.level` was supplied by the caller.

When the first non-empty line of the replacement body is a Markdown ATX heading whose text equals the resolved selected section heading text and whose ATX level equals the resolved selected section heading level, that heading line is stripped before retained proposal creation.

- This normalization applies to both direct `body` and `body_cache_id` replacement content.
- Only the first matching heading line is stripped.
- Body-internal headings after the first content line are preserved as section content.
- If the first non-empty line is a heading whose text does not equal the resolved selected section heading text, or whose ATX level does not equal the resolved selected section heading level, no stripping occurs and existing selector validation behavior applies.
- When stripping occurs, a `section_replacement_body_heading_stripped` warning diagnostic is returned with `stripped_heading` (the stripped heading text) and `stripped_level` (the ATX level of the stripped heading as an integer).
- The warning does not block retained proposal creation.
- Error-severity diagnostics continue to block proposal creation regardless of this normalization.
- Multi-section replacement, arbitrary string replacement, and canonical workflow section name changes are not supported by this operation.

Spec metadata block replacement example:

```json
{
  "kind": "spec",
  "id": "SPEC-design-records-mcp-tools",
  "update": {
    "type": "metadata_block_replace",
    "metadata": {
      "scope": "docs/spec/design-records-mcp/tools.md",
      "status": "draft",
      "design_record": {
        "id": "SPEC-design-records-mcp-tools",
        "kind": "spec",
        "status": "draft",
        "depends_on": ["ADR-076", "ADR-077", "ADR-087", "ADR-088", "ADR-090", "ADR-092", "ADR-093"]
      }
    }
  }
}
```

#### Operations array

`operations` is a list of update operation objects applied atomically to the same record in a single retained proposal.
`update` and `operations` are mutually exclusive. Supplying both MUST return `invalid_request` and MUST NOT create a proposal.
An empty `operations` array MUST return `invalid_request` and MUST NOT create a proposal.
A single-element `operations` array is valid and applies the same semantics as the corresponding single `update` operation.

Each element in `operations` is an operation object with the following fields:

| field | required | type | meaning |
|---|---:|---|---|
| `type` | yes | string | operation type; same values as `update.type` |
| `metadata` | conditional | object | required for `metadata_block_replace` and `metadata_fields_replace` |
| `section_selector` | conditional | object | required for `named_section_replace` |
| `body` | conditional | string | replacement body for `named_section_replace`; mutually exclusive with `body_cache_id` |
| `body_cache_id` | conditional | string | cached replacement body for `named_section_replace`; mutually exclusive with `body` |

Each operation in `operations` applies the full semantics of the corresponding single-operation type.
For `named_section_replace` operations, this includes section selector resolution, case-only fallback matching, and heading-safe replacement body normalization, as specified in `#### Named section replacement`.
For `metadata_fields_replace` operations, this includes reading the existing metadata, applying the patch, and validating the merged result, as specified in `#### Metadata field replacement`.

Per-type body source rules apply to each operation element identically to the single `update` operation:
- `body` and `body_cache_id` MUST be omitted for `metadata_block_replace` and `metadata_fields_replace` operations; supplying either MUST return `invalid_body_source` and MUST NOT create a proposal.
- Exactly one of `body` or `body_cache_id` MUST be present for `named_section_replace` operations.
- Supplying both `body` and `body_cache_id` in the same operation element MUST return `invalid_body_source` and MUST NOT create a proposal.

For each `named_section_replace` operation that supplies an inline `body`, if proposal preparation fails after the body has been received (for example, due to conflict detection, a validation error, or a section selector resolution failure), the response SHOULD include a retryable `body_cache` entry for that body, following the shared Body source and body cache rules.
Per-operation `body_cache_id` resolution uses the same `body_cache_not_found` and `body_cache_expired` behavior as single `update` requests.

Top-level `body` and `body_cache_id` fields are applicable only to single `update` requests. Supplying either alongside `operations` MUST return `invalid_body_source` and MUST NOT create a proposal.

**Operation ordering**

Operations are applied in the following deterministic order regardless of their position in the array:

1. `metadata_block_replace` and `metadata_fields_replace` operations, in the order they appear in the `operations` array.
2. `named_section_replace` operations, in the order they appear in the `operations` array.

**Conflict detection**

Conflict detection runs before any operation is applied. A conflicting `operations` array MUST NOT create a proposal.

- Two or more operations targeting the same metadata field MUST return `conflicting_operations` and MUST NOT create a proposal. For `metadata_fields_replace`, conflict is determined by comparing the `metadata` keys of the patch objects. A `metadata_block_replace` operation conflicts with any other metadata operation in the same array.
- Two or more `named_section_replace` operations in a single `operations` array MUST return `multiple_section_replace_not_supported` and MUST NOT create a proposal, regardless of whether they target the same or different sections. This is an MVP constraint.

**Validation**

Validation is performed against the final record state after all operations are applied, not against any intermediate state.

An `operations` array combining `metadata_fields_replace { status: done }` with a `named_section_replace` for `## Evidence` will pass done-state Evidence validation when the combined final state satisfies the required Evidence gate.

**No-op detection**

An `operations` request is a no-op when the combined result of all applied operations yields byte-equivalent persisted content to the current file. No retained proposal is created. The response has `proposal_created: false` and a `no_op_update` info diagnostic, identical to the single-operation no-op response.

**Response shape**

`operations` proposals use the same retained proposal response shape as single-operation proposals. `diff.text` covers all changes from all operations in a single unified diff for the target file. `diff.files` contains one entry for the target record with `change: modify`.

### Response

When an update request produces changed content, `propose_record_update` returns the common retained proposal response fields.
For retained update proposals, `target.resolved_id` is the existing record ID, and `diff.files[].change` is `modify`.

For existing-file update proposals, `diff.text` MUST compare the current persisted file content with the proposed persisted content after the requested operation semantics are applied.
It MUST NOT render the entire target record as newly added content unless the target file is actually new.

Modify proposal `diff.text` MUST use a git-style unified diff suitable for proposal review.
For each modified file, it MUST include:

- `diff --git a/<path> b/<path>`
- either an `index <oldhash>..<newhash> 100644` line or an equivalent stable old/new content hash representation when available
- `--- a/<path>`
- `+++ b/<path>`
- one or more `@@` hunk headers when file content differs
- changed lines plus bounded context, rather than the whole record as added content for metadata-only updates

Whole-file `+` output is valid for actual create/add proposals. It is not valid for existing-file modify proposals.

A no-op update MUST NOT create a retained write proposal.
A no-op update response MUST have:

| field | required | meaning |
|---|---:|---|
| `proposal_created` | yes | `false` |
| `operation` | yes | `update` |
| `target_kind` | yes | resolved target record kind |
| `target` | yes | requested / resolved target identity object |
| `validation` | yes | proposal-local validation result after applying operation semantics; `ok` is `true` when there are no error diagnostics |
| `diagnostics` | yes | includes an info diagnostic with category `no_op_update` |
| `diff` | no | omitted or `null` because no retained proposal exists |
| `proposal_id` | no | omitted because no retained proposal exists |

The `no_op_update` diagnostic identifies a successful update request that would not change the persisted file content.
It MUST have severity `info` and SHOULD include `record_id`, `path`, and `message`.
It MUST NOT be returned as a tool execution error.
It MUST NOT make `validation.ok` false.

No-op detection is evaluated after operation-specific semantics and normalization, including metadata field preservation, metadata validation, section selector resolution, and heading-safe replacement body normalization.
For example:

- `metadata_fields_replace` with `status: done` is a real update when the current status is not `done`.
- `metadata_fields_replace` with `status: done` is a no-op update when the current status is already `done` and the resulting persisted content is byte-equivalent.
- `named_section_replace` is a no-op update when the replacement body, after heading-safe normalization, yields byte-equivalent persisted content.

Metadata block replacement response example:

```json
{
  "proposal_id": "pw_opaque",
  "state": "proposed",
  "operation": "update",
  "target_kind": "task",
  "target": {
    "requested_id": "TASK-MCP-008-04",
    "resolved_id": "TASK-MCP-008-04",
    "kind": "task",
    "domain": "MCP",
    "path": "docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md"
  },
  "expires_at": "2026-06-05T00:00:00Z",
  "retention_days": 3,
  "diff": {
    "format": "unified",
    "files": [
      {
        "path": "docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md",
        "change": "modify",
        "record_id": "TASK-MCP-008-04",
        "record_kind": "task"
      }
    ],
    "text": "diff --git a/docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md b/docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md\nindex oldhash..newhash 100644\n--- a/docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md\n+++ b/docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md\n@@ -1,7 +1,7 @@\n # TASK-MCP-008-04: MCP tools spec reflection\n \n - **id**: TASK-MCP-008-04\n-- **status**: todo\n+- **status**: done\n - **date**: 2026-06-01\n"
  },
  "validation": {
    "ok": true,
    "diagnostics": []
  },
  "diagnostics": [],
  "note": "No repository files have been written. Call accept_proposed_write with this proposal_id to apply the diff."
}
```

No-op metadata field replacement response example:

```json
{
  "proposal_created": false,
  "operation": "update",
  "target_kind": "task",
  "target": {
    "requested_id": "TASK-MCP-008-04",
    "resolved_id": "TASK-MCP-008-04",
    "kind": "task",
    "domain": "MCP",
    "path": "docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md"
  },
  "validation": {
    "ok": true,
    "diagnostics": []
  },
  "diagnostics": [
    {
      "category": "no_op_update",
      "severity": "info",
      "record_id": "TASK-MCP-008-04",
      "path": "docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md",
      "message": "update produced no persisted content changes"
    }
  ]
}
```

Named section replacement response uses the same proposal shape for real changes. `diff.files[].change` is `modify`, and `target.resolved_id` is the existing record ID.

## `get_proposed_write`

### Purpose

`get_proposed_write` retrieves a retained proposal by proposal ID.
It does not write repository files.

### Request

```json
{
  "proposal_id": "pw_opaque"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `proposal_id` | yes | string | proposal lookup key |

### Response

For a retained proposal, response returns the proposal detail using the common proposal fields and current `state`.

Unknown proposal IDs return `proposal_not_found`.
Expired proposal IDs return `proposal_expired`.

## `accept_proposed_write`

### Purpose

`accept_proposed_write` applies a retained proposal after accept-time checks.
This is the only Design Records MCP authoring tool that may write repository files.

### Request

```json
{
  "proposal_id": "pw_opaque"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `proposal_id` | yes | string | proposal lookup key |

### Response

```json
{
  "proposal_id": "pw_opaque",
  "state": "accepted",
  "written": true,
  "files_written": [
    {
      "path": "docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md",
      "record_id": "TASK-MCP-008-04",
      "record_kind": "task"
    }
  ],
  "validation": {
    "ok": true,
    "diagnostics": []
  },
  "repair_guidance": [],
  "diagnostics": []
}
```

| field | required | meaning |
|---|---:|---|
| `proposal_id` | yes | accepted proposal ID |
| `state` | yes | `accepted` on successful write; otherwise the retained or rejection state |
| `written` | yes | whether repository files were modified by this accept call |
| `files_written` | yes | written file list; empty when `written: false` |
| `validation` | yes | post-accept validation result when applicable, otherwise current validation result |
| `repair_guidance` | yes | actionable repair suggestions; empty when no repair is needed |
| `diagnostics` | yes | accept diagnostics |

`written: false` is required for unknown, expired, discarded, already accepted, stale, changed target, ID collision, unresolved target, invalid proposal, and required-follow-up-not-satisfied outcomes.
These outcomes must not modify repository files.

`written: true` means repository files were modified.
If post-write validation fails after files were written, the response must still return `written: true`, include validation diagnostics, and provide repair guidance.
The MVP does not automatically roll back accepted writes after post-write validation failure.
The caller should create a repair proposal.

Force-accepting invalid proposals is outside MVP.
If pre-write validation has error diagnostics, accept returns `written: false`.

## `discard_proposed_write`

### Purpose

`discard_proposed_write` discards a retained proposal.
It does not write repository files.

### Request

```json
{
  "proposal_id": "pw_opaque"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `proposal_id` | yes | string | proposal lookup key |

### Response

```json
{
  "proposal_id": "pw_opaque",
  "state": "discarded",
  "discarded": true,
  "written": false,
  "diagnostics": []
}
```

Discarding an accepted proposal must not undo the accepted write.
Discard request for unknown or expired proposal IDs returns diagnostics and `discarded: false`.

## Error handling

Tool error code は以下を最小とする。

| code | meaning |
|---|---|
| `record_not_found` | `get_record` で指定された単一 record ID が存在しない。`get_records` では tool error ではなく item-level diagnostic として用いる |
| `guide_not_found` | `get_authoring_guidance` で指定された guide ID が存在しない |
| `invalid_request` | request schema または field value が不正。例: `list_records` に未知の `kind` を指定した場合、`get_records.ids` が欠落・空・非 array・非 string element を含む場合、`get_authoring_guidance.id` が欠落・非 string の場合、`propose_record_create` で required `fields` を省略した場合、top-level `id` と `fields.id` が一致しない場合、`new` placeholder create で `fields.id` を指定した場合、または `domain` が ID domain と case-insensitive に一致しない場合 |
| `unsupported_kind` | tool が対象外の `kind` を指定された。例: `suggest_next_record` に `kind: spec` を指定した場合 |
| `invalid_id_range` | `id_range` endpoint が malformed、unsupported family、mixed family、mixed domain、mixed task work sequence、または指定 `kind` と endpoint family 不一致である |
| `id_range_requires_decision_kind` | legacy error code。REQ-MCP-007 以前の decision-only `id_range` boundary を示す。新規 implementation では `invalid_id_range` を用いる |
| `proposal_not_found` | requested proposal ID が存在しない |
| `proposal_expired` | requested proposal ID は expiry を過ぎており、取得・accept・discard できない |
| `proposal_discarded` | proposal は discard 済みであり accept できない |
| `proposal_already_accepted` | proposal は accepted 済みであり再適用できない |
| `proposal_stale` | proposal base state と current target state が一致せず、accept できない |
| `target_changed` | target record kind / path / identity が proposal 作成時と異なる |
| `id_collision` | create proposal の resolved ID が accept 前に使用済みになった |
| `required_follow_up_not_satisfied` | workflow reciprocal metadata など required follow-up updates が未完了で accept できない |
| `invalid_body_source` | body source rule 違反。例: `body` と `body_cache_id` の両方を指定した、または required body source がない |
| `body_cache_not_found` | requested body cache ID が存在しない |
| `body_cache_expired` | requested body cache ID が expiry を過ぎている |
| `proposal_preparation_failed` | proposal preparation failed before proposal persistence; retry guidance may include `body_cache_id` |
| `section_selector_no_match` | named section selector が target record 内の section に一致しない |
| `section_selector_ambiguous` | named section selector が複数 section に一致し、単一 target に解決できない |

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

Authoring diagnostics may appear in normal authoring responses rather than tool execution errors when the tool can return `written: false`, proposal state, validation result, or retry guidance.
Invalid request shape may still be returned as a tool execution error.

## Authoring write boundary

The authoring transaction MVP intentionally excludes:

- generic filesystem write tools
- path-first authoring APIs
- immediate write create/update tools such as `create_record` or `update_record`
- `set_evidence` convenience operation
- `add_record_metadata` or add/remove relation convenience operations
- `migrate_record_to_spec`
- partial Markdown AST editing
- general-purpose multi-record atomic transactions with rollback semantics
- spec skeleton creation and `SPEC-new` placeholder create
- arbitrary unrelated record bundling in one proposal
- automatic close cascades across requirement / work item / task
- automatic rollback after accepted post-write validation failure
- formatter integration
- indefinite proposal or body cache retention
- force-accepting invalid proposals

Workflow artifact create proposals may include required reciprocal metadata updates in the same proposal when needed to keep relation validation valid.
That allowance is limited to required reciprocal updates and does not create a general-purpose multi-record atomic transaction.

Existing read / navigation / guidance tools keep their read-only behavior and request / response semantics.
`suggest_next_record` remains a read-only suggestion tool and does not create files; authoring creates use `propose_record_create`.

> 由来: ADR-077 §MVP外, ADR-077 §理由, ADR-093 §決定

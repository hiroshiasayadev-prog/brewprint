# PRODUCT-REQ-SPEC-001: MCP-readable spec format and topic tree support

- **id**: PRODUCT-REQ-SPEC-001
- **status**: accepted
- **date**: 2026-06-10
- **source_refs**:
  - V01-REQ-PRODUCT-003
  - PRODUCT-INV-SPEC-001
- **work_items**:
  - PRODUCT-WORK-SPEC-001
  - PRODUCT-WORK-SPEC-012
  - PRODUCT-WORK-SPEC-015

Note: `V01-REQ-PRODUCT-003` is listed as the repository-layout prerequisite for the current `product/records/**` placement. `PRODUCT-INV-SPEC-001` is the investigation evidence used to accept this requirement. This requirement itself is a new PRODUCT app-domain record and is not intended to be a V01 snapshot continuation.

## Requirement

brewprint の spec documents は、Design Records MCP が topic tree を抽出し、spec review 時に必要な topic だけを読めるようにするため、最小限の validation-friendly な書式を持たなければならない。

この書式は、front matter の `sections` や手書き topic metadata を正本にしない。topic structure は、Markdown の parser-aware heading、spec kind、Index table、MCP-managed visible metadata から構築できる形にする。

front matter は引き続き `design_record` などの record metadata を保持してよい。ただし、topic tree structure、parent relation、および新たに導入する topic reference の正本にはしない。

既存の stable `spec:` semantic ref を title-derived な identifier で置き換えるか、別物の derived topic reference を導入するかは、後続 spec で明示的に定義する。このREQ自体は既存の stable `spec:` semantic ref contract を変更しない。

## Evidence

- 現行の spec は、MVP 時代の stale な文章、現行 contract、履歴メモ、実装上の例外規則が同一文書に混在しており、読むべき範囲を topic 単位で絞り込めない。
- spec 文書には Markdown examples や migration remnants が混在しやすいため、parser-aware な heading rule がないと、実 heading と fenced code block 内の例示 heading を区別できない。
- front matter や文書先頭の `sections` に semantic metadata を集約すると、本文編集時に更新されず stale しやすい。
- spec review 時、context に巨大な spec 全体を渡す必要があり、LLM review の認知コストと誤読リスクが増えている。
- work item の task flow のように、本文から離れた複製的な進捗・関係情報は stale しやすい。spec topic relation も同じ失敗を避ける必要がある。

## Proposed Initial Format

このREQで検討する初期フォーマット案は以下とする。後続 spec / work item で既存 spec への適用可否を調査し、必要に応じて調整する。

### H1 format

各 spec markdown file は、YAML front matter と fenced code block の外側に ATX H1 をちょうど1つだけ持つ。

```markdown
# <SpecKind>: <Title>
```

例:

```markdown
# Overview: Design Records MCP
# Index: DRMCP record model
# Concept: Record identity model
# Reference: Record metadata fields
# Contract: validate_records tool
```

### Initial spec kind candidates

| spec kind | intent | expected content |
|---|---|---|
| `Overview` | spec area の入口 | What this is / current contract / non-goals / topic map / related specs |
| `Index` | topic list と navigation | Topics table。原則として仕様本文を持たない |
| `Concept` | 概念モデル・用語・境界 | concept model / rules / boundary |
| `Reference` | field / enum / grammar / fixed rule の辞書 | table-based reference |
| `Contract` | API / tool / request / response / validation / error contract | request / response / errors / validation rules |
| `Guide` | usage / authoring / operation procedure | steps / examples / constraints |
| `Process` | lifecycle / workflow / transition model | states / transitions / responsibilities |
| `Architecture` | component / runtime / storage / dependency / scanning structure | components / data flow / boundaries |
| `Glossary` | term list | term table |

### Common required sections

All spec kinds must have:

```markdown
## What this is
```

### Overview format candidate

```markdown
# Overview: <Title>

## What this is

## Current contract

## Non-goals

## Topic map

## Related specs
```

### Index format candidate

```markdown
# Index: <Title>

## What this is

## Topics

| title | kind | parent | file | summary |
|---|---|---|---|---|
```

### Concept format candidate

```markdown
# Concept: <Title>

- **parent**: `<parent-ref>`

## What this is

## Concept model

## Rules

## Boundary

## Related specs
```

The parent line is visible near the title but must be MCP-managed when it is derived from an Index table.

### Reference format candidate

```markdown
# Reference: <Title>

- **parent**: `<parent-ref>`

## What this is

## <Reference table title>

| name | description |
|---|---|

## Related specs
```

Reference specs may define stricter table columns per reference type, for example `field`, `type`, `required`, `description`, or `value`, `meaning`, `description`.

### Contract format candidate

```markdown
# Contract: <Title>

- **parent**: `<parent-ref>`

## What this is

## Request

## Response

## Errors

## Validation rules

## Related specs
```

### Guide format candidate

```markdown
# Guide: <Title>

- **parent**: `<parent-ref>`

## What this is

## Procedure

## Examples

## Constraints

## Related specs
```

### Process format candidate

```markdown
# Process: <Title>

- **parent**: `<parent-ref>`

## What this is

## States

## Transitions

## Responsibilities

## Related specs
```

### Architecture format candidate

```markdown
# Architecture: <Title>

- **parent**: `<parent-ref>`

## What this is

## Components

## Data flow

## Boundaries

## Related specs
```

The investigation must decide whether `Guide`, `Process`, and `Architecture` remain supported initial kinds or are deferred until existing spec migration proves they are necessary.

### Glossary format candidate

```markdown
# Glossary: <Title>

## What this is

## Terms

| term | definition | related |
|---|---|---|
```

## Required Outcome

- spec H1 format が定義されている。
  - 例: `# Overview: <Title>`, `# Concept: <Title>`, `# Reference: <Title>`, `# Contract: <Title>`
- supported spec kinds が定義されている。
- initial spec kind candidates と kind ごとの intent が評価されている。
- kind ごとの初期 format candidate が定義されている。
- 各 spec markdown file は YAML front matter と fenced code block の外側に ATX H1 をちょうど1つだけ持つ、という parser-aware 制約が定義されている。
- 全 spec が `## What this is` を持つ、という制約が定義されている。
- spec kind ごとの必須 section が定義されている。
- Index spec は `## Topics` table により child topics を列挙する、という制約が定義されている。
- `## Topics` table の必須列が定義されている。
  - child title
  - child spec kind
  - child target file
  - parent reference or root marker
- derived topic reference を導入する場合は、parent reference と normalized title から機械的に導出するルールが定義されている。
- derived topic reference と既存 stable `spec:` semantic ref の関係が定義されている。
  - derived topic reference が stable `spec:` semantic ref を置き換える場合は、rename / move / split / merge / collision / redirect / alias / legacy `semantic_refs` and `sections` compatibility behavior が定義されている。
  - derived topic reference を stable `spec:` semantic ref とは別物として扱う場合は、両者の resolver / validation / display boundary が定義されている。
- parent relation は Index table を正本とし、child spec の H1 直下に置く visible parent metadata は MCP-managed とするルールが定義されている。
- front matter は topic structure の正本にしない、という方針が明示されている。
- Design Records MCP が検証できる violation categories が定義されている。
  - missing required section
  - invalid spec H1 format
  - invalid H1 count outside front matter and fenced code
  - invalid or missing Index Topics table column
  - unresolved child target
  - H1 / Topics table kind mismatch
  - duplicate child parent
  - orphan child topic
  - parent metadata mismatch
  - topic cycle, or an explicit decision to defer cycle validation
- 既存 spec set に対して、この初期フォーマットへ migrate 可能かが調査されている。
  - directly migratable specs
  - split required specs
  - ownership relocation required specs
  - incompatible or ambiguous specs
  - format candidate revision points

## Investigation Expectations

The follow-on investigation must produce evidence and recommendations. It must not finalize the derived topic reference contract by itself.

The investigation should sample at least:

- `drmcp/records/spec/design-records-mcp/schema.md`
- `drmcp/records/spec/design-records-mcp/tools.md`
- `drmcp/records/spec/design-records-mcp/overview.md`
- `product/records/spec/concepts/traceability/**`
- `product/records/spec/concepts/project-artifact-model/index.md`
- `product/records/spec/concepts/namespace-model/index.md`
- `bpdsl/records/spec/overview.md`
- `bpdsl/records/spec/nodes.md`
- `bpdsl/records/spec/edges.md`
- `bpdsl/records/spec/mcp/schema.md`
- `bpdsl/records/spec/mcp/tools/inspect.md`
- `bpdsl/records/spec/views/dag.md`
- `bpdsl/records/spec/views/sequence-diagram.md`
- corresponding `v01/records/spec/**` snapshots where migration remnants need comparison.

Expected evidence:

- migration classification table per sampled file
- kind-fit matrix for proposed spec kinds
- parser scan showing real H1s versus fenced-code example headings
- validation rule matrix: mechanical / semantic / deferred
- compatibility matrix for stable `spec:` refs versus derived topic refs
- recommended later WORK items

## Explicitly Excluded Scope

- 既存 spec set の一括再構造化
- Design Records MCP の topic tree extraction 実装
- Design Records MCP の validation diagnostics 実装
- ADR / spec semantic reference linkage の詳細設計
- spec authoring guide / ADR authoring guide / workflow authoring guide の実更新
- UI における topic tree 表示

## Boundary

このREQは、brewprint records 全体で使う spec document format と topic tree extraction に必要な最小構造を定義する要求を所有する。

既存 spec の移行、MCP 実装、authoring guides 更新、ADR semantic reference linkage は後続 work item / requirement / task に委ねる。

## Follow-up Candidates

- spec format authoring guide を作成または更新し、H1 format、spec kind、required sections、Index Topics table、derived topic reference policy、MCP-managed parent metadata の運用手順を定義する。
- Design Records MCP に spec format validation を追加する。
  - H1 count validation
  - H1 spec kind validation
  - required section validation
  - Index Topics table validation
  - parent ref resolution
  - MCP-managed parent metadata mismatch detection
- ADR が関連する spec topic semantic ref を宣言する仕組みを定義する。spec 側に historical ADR backlink を手書きしない。
- ADR semantic refs の resolve validation と、spec topic から ADR を逆引きする query / projection を検討する。
- 既存の DRMCP / BPDSL / PRODUCT spec set を新 format に migrate 可能か inventory し、direct migration / split required / ownership relocation / incompatible に分類する。
- 既存の DRMCP / BPDSL / PRODUCT spec set を新 format に分割・再構造化する。
- `v01/records/spec/concepts` 配下の spec ownership を見直し、DRMCP-owned concepts と PRODUCT-owned concepts の移管方針を決める。
- stable `spec:` semantic ref と derived topic reference の互換性方針を定義する。
- stale な `depends_on` / legacy `docs/` path references / old MVP wording の cleanup を行う。

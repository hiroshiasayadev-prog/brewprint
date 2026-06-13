# V01-REQ-PRODUCT-004: MCP-readable spec format and topic tree support

- **id**: V01-REQ-PRODUCT-004
- **status**: captured
- **date**: 2026-06-10
- **source_refs**:
  - V01-REQ-PRODUCT-003
- **work_items**:

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

## Required Outcome

- spec H1 format が定義されている。
  - 例: `# Overview: <Title>`, `# Concept: <Title>`, `# Reference: <Title>`, `# Contract: <Title>`
- supported spec kinds が定義されている。
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
- 既存の DRMCP / BPDSL / PRODUCT spec set を新 format に分割・再構造化する。
- `v01/records/spec/concepts` 配下の spec ownership を見直し、DRMCP-owned concepts と PRODUCT-owned concepts の移管方針を決める。
- stable `spec:` semantic ref と derived topic reference の互換性方針を定義する。
- stale な `depends_on` / legacy `docs/` path references / old MVP wording の cleanup を行う。

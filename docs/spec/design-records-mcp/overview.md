---
scope: docs/spec/design-records-mcp/overview.md
status: draft
last_updated: 2026-05-27
summary: >
  Design Records MCP の目的、対象範囲、既存 brewprint MCP との責務境界、
  MVP の非目標を定義する。
depends_on:
  - docs/adr/076-design-records-mcp.md
  - docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
  - docs/adr/090-design-records-mcp-batch-retrieval-tool-boundary.md
  - docs/adr/092-design-records-mcp-workflow-artifact-record-and-relation-boundary.md
design_record:
  id: SPEC-design-records-mcp-overview
  kind: spec
  status: draft
  depends_on:
    - ADR-076
    - ADR-077
    - ADR-087
    - ADR-088
    - ADR-090
    - ADR-092
---

# Design Records MCP overview

## 目的

Design Records MCP は、brewprint の design record / workflow artifact 運用を machine-readable metadata と MCP query / validation で支援する補助 MCP である。ADR / spec / investigation に加え、requirement / work item / task も対象とする。

主目的は以下である。

- ADR / spec / investigation / requirement / work item / task の record index を構築する
- record の ID / kind / status / path と kind 固有 metadata を構造化して取得できるようにする
- 選択済みの複数 record ID について detail representation をまとめて取得できるようにする
- record metadata の基本不整合を検出する
- docs artifact 間の semantic/artifact ref を解決し、参照切れ検査に利用できるようにする
- 別セッションの LLM が、読むべき design record を本文読解前に絞り込めるようにする

Design Records MCP は、spec-first ドキュメント運用を置き換えない。
現行仕様の唯一の正は引き続き `docs/spec/**` であり、ADR は設計判断の根拠記録である。
Design Records MCP は、その関係を機械的に辿りやすくする query / validation layer として扱う。

> 由来: ADR-076 §決定, ADR-076 §理由

## 対象 record

index / query / validation 対象は以下とする。

| kind | 対象 |
|---|---|
| `decision` | `docs/adr/**` に置かれる ADR |
| `spec` | `docs/spec/**` のうち YAML front matter に `design_record.id` と `design_record.kind` を持つ spec |
| `investigation` | `docs/investigations/<domain>/INV-<DOMAIN>-NNN-*.md` に置かれる investigation artifact |
| `requirement` | `docs/requirements/<domain>/REQ-<DOMAIN>-NNN-*.md` に置かれる requirement artifact |
| `work_item` | `docs/work-items/<domain>/WORK-<DOMAIN>-NNN-*.md` に置かれる work item artifact |
| `task` | `docs/tasks/<domain>/TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>-*.md` に置かれる新形式 task artifact |

`design_record` を持たない既存 spec は index 対象外とし、`missing_design_record` diagnostic も出さない。
既存 `docs/tasks/m*.md` と `docs/TASKS.md` の M-series は legacy record / historical label であり、`task` record discovery の対象に含めない。

本specは record kind 全体を閉じた列挙として確定しない。後続判断により他の artifact kind を追加しうる。
MVP では UC docs / impl notes は record kind として index 対象に含めない。

> 由来: ADR-076 §MVP対象, ADR-087 §1, ADR-091 §1〜§6, ADR-092 §1

## 既存 brewprint MCP との責務境界

Design Records MCP は、既存 brewprint MCP とは別の対象データソースを扱う。

| MCP | 対象 | 主な責務 |
|---|---|---|
| brewprint MCP | brewprint YAML から構築された `ResolvedProject` | semantic object の query / inspect / impact analysis |
| Design Records MCP | `docs/adr/**` / `docs/investigations/**` / `docs/requirements/**` / `docs/work-items/**` / 新形式 `docs/tasks/<domain>/TASK-*.md` の箇条書きmetadataと、`design_record` を持つ `docs/spec/**` の YAML front matter | design record / workflow artifact の index / read / validation、および traceability spec に従う semantic/artifact ref resolve |

Design Records MCP は、既存 brewprint MCP とは独立して起動・検証できる構成を第一候補とする。
既存 `QueryService` の責務を docs 管理へそのまま拡張しない。

この spec は既存 brewprint MCP の `docs/spec/mcp/tools/**` には混ぜず、`docs/spec/design-records-mcp/**` に置く。

> 由来: ADR-076 §既存brewprint MCPとの関係

## Resolver responsibility

Design Records MCP は、traceability spec が定める canonical reference model に従い、ref の resolve とその結果を用いた validation を担う。

ADR-088 / ADR-092 により、MVP で必須とする resolver input は active `spec:` semantic ref、Design Records MCP が扱う record ID-as-ref (`ADR-*` / `SPEC-*` / `INV-*` / `REQ-*` / `WORK-*` / `TASK-*`)、investigation canonical reference validation、および workflow relation integrity validation とする。

Investigation metadata が canonical reference として追加利用できる workflow ID-as-ref は `REQ-*` / `WORK-*` に限定し、`TASK-*` は direct resolver input および workflow artifact 間 relation のみで support する。

`internal-design:` / `coverage:` / `COV-*`、coverage mapping、semantic realization relation は MVP required resolver scope に含めない。Orphan workflow artifact diagnostics、task status 由来 progress projection、workflow 専用 traversal tool も MVP に含めない。

resolver が lookup source として読む artifact と、`list_records` / `get_record` が record kind として公開する artifact は同一集合である必要はない。

Resolver の public tool 名は `resolve_reference` とし、active `spec:` semantic ref と record ID-as-ref (`ADR-*` / `SPEC-*` / `INV-*` / `REQ-*` / `WORK-*` / `TASK-*`) を active lookup input として扱う contract を `tools.md` に定義する。`internal-design:` / `coverage:` / `COV-*`、physical path、および未採用 ID form は `unsupported` response とする。Reserved prefix である `yaml:` の public resolver input / direct query response behavior は MVP で定義しない。

> 由来: ADR-087 §4, ADR-088, ADR-092 §3〜§7

## MVP tool set

MVP の P0 tool は以下である。

- `list_records`
- `get_record`
- `get_records`
- `validate_records`
- `resolve_reference`

P1 の任意補助 tool として、以下を許容する。

- `suggest_next_record`

P0 tool は read-only であり、record 本文を読む前の候補絞り込みと metadata 整合性検証を目的とする。

> 由来: ADR-077 §P0: MVP必須tool, ADR-077 §P1: MVPに含めてもよい補助tool, ADR-090 §決定

## MVP 外

MVP では以下を扱わない。

- `trace_record`
- `list_gaps`
- `create_record`
- `update_record`
- `set_evidence`
- その他の write 系 tool
- 自然言語本文から依存関係を推定すること
- spec 本文との厳密な意味照合
- git 履歴解析
- code static analysis
- Web UI
- 複数プロジェクト横断管理
- 汎用 OSS CLI としての公開 contract
- section 単位の完全な traceability
- `topics` / `affects` / `refines` / `conflicts_with` metadata
- legacy M-series task record / UC docs / impl notes の record kind としての index 化
- `internal-design:` / `coverage:` / `COV-*` の resolve と semantic realization relation validation
- coverage mapping query
- orphan requirement / orphan work item / orphan task diagnostics
- task status から work item progress を導出する projection
- workflow 専用 traversal / tree / graph query tool
- task dependency cycle detection / execution order projection
- investigation metadata から `TASK-*` を canonical reference として辿ること

MVP は ADR / investigation / requirement / work item / task の箇条書きmetadata、spec の YAML front matter、H1、path から得られる明示情報だけを扱う。
自然言語本文の推定や運用 gap 診断は、MVP の validator を実データへ当てた後に追加可否を判断する。

> 由来: ADR-076 §MVPスコープ外, ADR-077 §MVP外, ADR-092 §7

## filesystem tool との責務境界

Design Records MCP は汎用 filesystem tool の代替ではない。

Design Records MCP が扱うもの:

- 単一または明示された複数の record ID から metadata / path / headings / raw body を取得する
- record 一覧を構造化して返す
- record metadata の基本不整合を検証する
- 次の ADR 番号と推奨 path を提案する

Design Records MCP が扱わないもの:

- 任意ファイルの読み書き
- Markdown 一般編集
- ADR 本文の自動生成・自動更新
- commit hash の自動書き換え
- git 操作

> 由来: ADR-077 §filesystemとの責務境界, ADR-090 §決定

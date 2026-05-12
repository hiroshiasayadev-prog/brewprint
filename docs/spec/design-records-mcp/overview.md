---
scope: docs/spec/design-records-mcp/overview.md
status: draft
last_updated: 2026-05-12
summary: >
  Design Records MCP の目的、対象範囲、既存 brewprint MCP との責務境界、
  MVP の非目標を定義する。
depends_on:
  - docs/adr/076-design-records-mcp.md
  - docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md
design_record:
  id: SPEC-design-records-mcp-overview
  kind: spec
  status: draft
  depends_on:
    - ADR-076
    - ADR-077
---

# Design Records MCP overview

## 目的

Design Records MCP は、brewprint の ADR / spec 運用を machine-readable metadata と MCP query / validation で支援する補助 MCP である。

主目的は以下である。

- ADR / spec の record index を構築する
- record の ID / kind / status / path / 依存関係を構造化して取得できるようにする
- ADR / spec metadata の基本不整合を検出する
- 別セッションの LLM が、読むべき ADR / spec を本文読解前に絞り込めるようにする

Design Records MCP は、spec-first ドキュメント運用を置き換えない。
現行仕様の唯一の正は引き続き `docs/spec/**` であり、ADR は設計判断の根拠記録である。
Design Records MCP は、その関係を機械的に辿りやすくする query / validation layer として扱う。

> 由来: ADR-076 §決定, ADR-076 §理由

## 対象 record

MVP の index 対象は以下に限定する。

| kind | 対象 |
|---|---|
| `decision` | `docs/adr/**` に置かれる ADR |
| `spec` | `docs/spec/**` のうち YAML front matter に `design_record.id` と `design_record.kind` を持つ spec |

`design_record` を持たない既存 spec は MVP index 対象外とし、`missing_design_record` diagnostic も出さない。

MVP では task file / UC docs / impl notes は index 対象に含めない。
これらは ADR-068 の責務境界上の artifact ではあるが、作業状態・実例・実装引継ぎの性質が強く、ADR/spec とは更新頻度と正とする情報が異なるためである。

> 由来: ADR-076 §MVP対象

## 既存 brewprint MCP との責務境界

Design Records MCP は、既存 brewprint MCP とは別の対象データソースを扱う。

| MCP | 対象 | 主な責務 |
|---|---|---|
| brewprint MCP | brewprint YAML から構築された `ResolvedProject` | semantic object の query / inspect / impact analysis |
| Design Records MCP | `docs/adr/**` の箇条書きmetadataと、`design_record` を持つ `docs/spec/**` の YAML front matter | design record の index / read / validation |

Design Records MCP は、既存 brewprint MCP とは独立して起動・検証できる構成を第一候補とする。
既存 `QueryService` の責務を docs 管理へそのまま拡張しない。

この spec は既存 brewprint MCP の `docs/spec/mcp/tools/**` には混ぜず、`docs/spec/design-records-mcp/**` に置く。

> 由来: ADR-076 §既存brewprint MCPとの関係

## MVP tool set

MVP の P0 tool は以下である。

- `list_records`
- `get_record`
- `validate_records`

P1 の任意補助 tool として、以下を許容する。

- `suggest_next_record`

P0 tool は read-only であり、record 本文を読む前の候補絞り込みと metadata 整合性検証を目的とする。

> 由来: ADR-077 §P0: MVP必須tool, ADR-077 §P1: MVPに含めてもよい補助tool

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
- task file / UC docs / impl notes の index 化

MVP は ADR の箇条書きmetadata、spec の YAML front matter、H1、path から得られる明示情報だけを扱う。
自然言語本文の推定や運用 gap 診断は、MVP の validator を実データへ当てた後に追加可否を判断する。

> 由来: ADR-076 §MVPスコープ外, ADR-077 §MVP外

## filesystem tool との責務境界

Design Records MCP は汎用 filesystem tool の代替ではない。

Design Records MCP が扱うもの:

- record ID から metadata / path / headings / raw body を取得する
- record 一覧を構造化して返す
- record metadata の基本不整合を検証する
- 次の ADR 番号と推奨 path を提案する

Design Records MCP が扱わないもの:

- 任意ファイルの読み書き
- Markdown 一般編集
- ADR 本文の自動生成・自動更新
- commit hash の自動書き換え
- git 操作

> 由来: ADR-077 §filesystemとの責務境界

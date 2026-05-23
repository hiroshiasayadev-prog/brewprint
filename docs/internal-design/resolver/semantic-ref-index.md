---
scope: docs/internal-design/resolver/semantic-ref-index.md
status: draft
last_updated: 2026-05-24
summary: >
  Design Records MCP が canonical semantic/artifact ref を解決するための
  index と validation route の internal design を記録する。
depends_on:
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
---

# Semantic ref index resolver

## Purpose

この document は、Design Records MCP が ADR-087 / ADR-088 に基づく canonical reference resolve / validation を実装するときの lookup route を記録する internal design である。

これは public MCP contract ではない。Resolve tool の名称、request / response schema、diagnostic category は `docs/spec/design-records-mcp/` および M19 で確定・追従する。

ADR-088 により、internal design document 自体は semantic trace MVP の active endpoint または required acceptance example ではない。本 document は implementation-facing design artifact としてのみ存続する。

## MVP lookup sources

Resolver index は、少なくとも以下の explicit metadata / record ID を lookup source として扱う。

| ref kind | source metadata / index | resolves to |
|---|---|---|
| `spec:...` | spec front matter の `semantic_refs` / `sections` | spec document / section |
| `ADR-*` / `SPEC-*` / `INV-*` | Design Records MCP record index | design record artifact |

Investigation metadata の `source_refs` / 記載済み `follow_up_results` は resolver の validation input として扱う。`follow_up_candidates` に artifact reference が記載されている場合は canonical form を検査するが、未作成候補の存在は要求しない。

Resolver が lookup source として読む artifact と、`list_records` / `get_record` の record kind として公開する artifact は同一集合である必要はない。

## Deferred lookup sources

MVP required lookup source に含めないもの:

| ref kind / mechanism | reason |
|---|---|
| `internal-design:...` | internal design endpoint の concrete consumer requirement が未確認 |
| `coverage:...` / `COV-*` | external coverage artifact は MVP 外 |
| coverage mapping relation | semantic realization relation は MVP 外 |
| `yaml:...` | reserve only |
| `REQ-*` / `WORK-*` public resolve | concrete public contract は後続判断 |

## Validation route

M19 実装では、resolve 結果を少なくとも以下の検証に利用する。

1. `spec:` semantic ref または record ID-as-ref を index から検索する。
2. `spec:` ref が一意の document / section に解決できるかを検査する。
3. `ADR-*` / `SPEC-*` / `INV-*` が一意の record artifact に解決できるかを検査する。
4. Investigation の `source_refs` と記載済み `follow_up_results` が canonical reference であり解決可能かを検査する。
5. `follow_up_candidates` は canonical form を検査できるが、未作成 target であること自体は error にしない。
6. Physical path が canonical reference 欄へ現れた場合は noncanonical input として診断候補にする。

MVP では internal-design / coverage relation や `COV-*` mapping を resolve / validate しない。

## Verification targets

Concrete test は M19 が所有する。Required acceptance target は、少なくとも以下を含む方向で整理する。

- `spec:trace.resolve-and-validation` 等の active `spec:` ref が解決できること
- investigation record が取得できること
- investigation の `source_refs` / 記載済み `follow_up_results` の unresolved を error とできること
- unresolved `follow_up_candidates` を orphan error にしないこと
- path-based canonical reference input を noncanonical と診断できること

旧 `COV-TRACE-001` や `spec:` → `internal-design:` mapping は acceptance target にしない。

## Non-goals

- resolve tool の concrete public request / response schema
- diagnostic category 名と severity の確定
- `internal-design:` / `coverage:` / `yaml:` endpoint の active 化
- semantic realization relation validation
- coverage mapping query tool
- MCP writer tools

> 由来: ADR-083, ADR-087, ADR-088; INV-DOCS-002; INV-DOCS-003

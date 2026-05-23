# Internal design

`docs/internal-design/` は、現行 spec semantics と brewprint DSL YAML model を target implementation へ写像する internal wiring / route を記録する artifact layer である。

Internal design は現行仕様の source of truth ではなく、brewprint DSL YAML や target implementation の source of truth でもない。公開 contract、YAML semantics、diagnostic behavior、MCP public response schema は `docs/spec/` が所有する。

## When to create internal design

以下のような、外部仕様ではないが複数 component にまたがって長期的に維持すべき内部構造を記録する場合に作る。

- resolver / index / cache の責務と lookup route
- parser / semantic model / validator 間の境界
- MCP tool implementation の内部 phase 分担
- spec semantics と実装 component の多対多 mapping を検討するための設計記述

局所的な実装メモや一時的な handoff は `docs/impl/` または code 側に置く。

## Layout

Physical layout は semantic identity ではないため、topic に応じて分割してよい。

初期 layout:

```text
docs/internal-design/<domain>/<topic>.md
```

例:

```text
docs/internal-design/resolver/semantic-ref-index.md
```

Front matter は doc-policy に従う。

```yaml
---
scope: docs/internal-design/resolver/semantic-ref-index.md
status: draft
last_updated: YYYY-MM-DD
summary: >
  semantic ref resolver の lookup route を記録する。
depends_on:
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
---
```

## Semantic trace MVP boundary

ADR-088 により、internal design layer の存在と semantic trace endpoint 化は分離された。

Semantic trace MVP では、以下を要求しない。

- `internal-design:` semantic ref の宣言・解決・validation
- internal design metadata による source `spec:` relation declaration
- `spec:` から internal design への reverse graph
- internal design を endpoint とする `maps_to` / `covers` relation

Internal design document は引き続き作成・レビューできるが、MVP の canonical reference resolution / validation acceptance target には含めない。

## Future activation trigger

以下のような concrete requirement が成立した場合、`internal-design:` endpoint、metadata、resolver / validation contract の導入を後続 ADR / requirement / work item で再判断する。

- spec から複数 internal design への機械的 navigation / impact analysis が必要になった場合
- investigation / work item / MCP query が internal design artifact の canonical resolve を必要とする場合
- `yaml:` active 化により、spec / internal design / YAML の realization chain が必要になった場合

## Boundary

- spec が external / authoring contract を所有する。
- internal design は implementation-facing wiring route を所有する。
- semantic trace MVP は `spec:` と record / investigation canonical references の resolve / validation に限定する。
- external coverage artifact と semantic realization relation は MVP 外であり、必要性が確認された場合に再判断する。
- `yaml:` は semantic trace MVP では reserve only であり、active endpoint ではない。

> 由来: ADR-083, ADR-084, ADR-087, ADR-088; INV-DOCS-003

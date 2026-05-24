# REQ-MCP-003: Requirement / work item / task の MCP support が必要

- **id**: REQ-MCP-003
- **status**: captured
- **date**: 2026-05-25
- **source_refs**:
  - ADR-081
  - ADR-083
  - ADR-087
- **work_items**:

## 要求

LLM が requirement / work item / task による運用チェーンを確認するとき、文書を個別に探索するだけに依存せず、関連 artifact を query / resolve / validate できる MCP capability が必要である。

## 発見根拠

M19 は、既に以下の運用チェーンで管理されている。

```text
REQ-MCP-001 -> WORK-MCP-001 -> M19
```

一方、現行 Design Records MCP と M19 の canonical reference foundation は `ADR-*` / `SPEC-*` / `INV-*` および active `spec:` ref を対象とし、`REQ-*` / `WORK-*` / task identity は public record / resolver contract に含めていない。

このため、運用 artifact を用いて work を進め始めているにもかかわらず、LLM は requirement から work item、work item から task、task から完了根拠を MCP 経由で一貫して辿れない。

## Candidate capability

後続判断では、少なくとも以下を検討する。

- requirement / work item / task の record kind または別 query interface としての公開
- `REQ-*` / `WORK-*` / task identity の canonical resolve boundary
- requirement の `work_items`、work item の `source_requirement` / `tasks` の validation
- orphan requirement / orphan work item / orphan task の diagnostic boundary
- M19 のような実運用結果を基にした必要最小限の query surface

## Boundary

- 本 requirement は workflow artifact MCP support の必要性を捕捉するものであり、record kind、resolver input、validation category、tool 名を確定しない。
- M19 の必須完了条件には含めない。
- Design Records MCP の batch retrieval capability は別 requirement で扱う。
- requirements / work-items / tasks の docs artifact 運用自体は、本 requirement の採用判断を待たず継続できる。

## Next decision

既存の `REQ-MCP-001 -> WORK-MCP-001 -> M19` を用いた M19 implementation / close 運用から evidence を収集し、MCP support を public contract として採用するか、その最小 scope を判断する。

# REQ-MCP-003: Requirement / work item / task の MCP support が必要

- **id**: REQ-MCP-003
- **status**: captured
- **date**: 2026-05-25
- **source_refs**:
  - ADR-081
  - ADR-083
  - ADR-087
  - ADR-091
- **work_items**:

## 要求

LLM が requirement / work item / task による運用チェーンを確認するとき、文書を個別に探索するだけに依存せず、関連 artifact を query / resolve / validate できる MCP capability が必要である。

## 発見根拠

REQ-MCP-002 は、ADR-091 の判断根拠となる dogfooding として、以下の運用チェーンで完了した。

```text
REQ-MCP-002 -> WORK-MCP-002 -> TASK-MCP-002-01 ... TASK-MCP-002-06
```

この flow では、evidence 整理、ADR 判断、spec 更新、implementation、runtime verification、close review を短期 task として管理し、work item 上の Mermaid task graph と status 遷移が有効であることを確認した。

一方、現行 Design Records MCP と canonical reference foundation は `ADR-*` / `SPEC-*` / `INV-*` および active `spec:` ref を対象とし、`REQ-*` / `WORK-*` / `TASK-*` ID-as-ref は public record / resolver contract に含めていない。

このため、workflow artifact による work を実運用で進められる一方、LLM は requirement から work item、work item から task、task 間 dependency と完了根拠を MCP 経由で一貫して辿れない。

## Candidate capability

後続判断では、少なくとも以下を検討する。

- requirement / work item / task の record kind または別 query interface としての公開
- `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref の canonical resolve boundary
- requirement の `work_items`、work item の `source_requirement` / `tasks`、task の `work_item` / `source_requirement` / `depends_on` の validation
- orphan requirement / orphan work item / orphan task の diagnostic boundary
- task status から導出する progress projection の要否。手書き checkbox を source of truth として追加しない境界
- `WORK-MCP-002` の運用結果を基にした必要最小限の query surface

## Boundary

- 本 requirement は workflow artifact MCP support の必要性を捕捉するものであり、record kind、resolver input、validation category、tool 名を確定しない。
- ADR-091 に従い、workflow artifact 間の canonical relation は `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref を前提とし、physical path は supported canonical relation に含めない。
- `req:` / `work:` / `task:` の semantic prefix は導入前提としない。
- Design Records MCP の batch retrieval capability は REQ-MCP-002 / ADR-090 で完了済みであり、本 requirement は `get_records` 等が workflow artifact を扱う場合の拡張境界を後続判断する。
- requirements / work-items / tasks の docs artifact 運用自体は、本 requirement の採用判断を待たず継続できる。

## Next decision

`WORK-MCP-002` dogfooding と ADR-091 を evidence として、`REQ-*` / `WORK-*` / `TASK-*` ID-as-ref を public record / resolver / validation contract に含める最小 scope を判断する。

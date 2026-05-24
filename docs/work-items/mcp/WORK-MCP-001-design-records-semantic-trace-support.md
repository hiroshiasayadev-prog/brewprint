# WORK-MCP-001: Design Records MCP semantic trace support を実装追従する

- **id**: WORK-MCP-001
- **status**: implementation_pending
- **date**: 2026-05-23
- **source_requirement**: REQ-MCP-001
- **impact_refs**:
  - ADR-087
  - ADR-088
  - spec:trace.resolve-and-validation
  - spec:trace.metadata-schema
- **tasks**:
  - M19

## Goal

ADR-087 / ADR-088 と traceability spec で確定した canonical reference resolution foundation と investigation record integration を、Design Records MCP の contract・実装・tests に一貫して反映する。

## Impact scope

| layer | status | impact |
|---|---|---|
| decision | done | ADR-087 / ADR-088 accepted |
| design spec | done | active `spec:` / record ID / investigation canonical ref の concrete resolve tool / diagnostic contract を M19 Phase A で確定した |
| internal design | done | M19 Phase B0 で finalized public contract に追従する resolver / record parser / validation route を具体化した。MVP endpoint としては扱わない |
| implementation | pending | M19 Phase B/C で investigation kind、kind-specific response、canonical resolver、validator を実装する |
| verification | pending | M19 Phase D で investigation record / `spec:` ref / record ID-as-ref の参照解決・診断 tests を追加する |

## Boundary

- 本 work item は横断進捗と影響範囲を所有する。
- 具体的な実装順序とチェックリストは M19 task が所有する。
- `internal-design:` / `coverage:` / `COV-*` の resolve、semantic realization relation、coverage mapping query、MCP writer tools は本 work item の必須完了条件に含めない。

## Done condition

M19 の完了条件を満たし、Design Records MCP が ADR-087 / ADR-088 に基づく investigation integration と canonical reference resolution / validation を contract・implementation・tests の同一切替単位で提供できること。

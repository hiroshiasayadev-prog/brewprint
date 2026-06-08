# V01-WORK-MCP-001: Design Records MCP semantic trace support を実装追従する

- **id**: V01-WORK-MCP-001
- **status**: done
- **date**: 2026-05-23
- **source_requirement**: V01-REQ-MCP-001
- **impact_refs**:
  - V01-ADR-087
  - V01-ADR-088
  - spec:trace.resolve-and-validation
  - spec:trace.metadata-schema
- **tasks**:
  - V01-TASK-MCP-001-01

## Goal

V01-ADR-087 / V01-ADR-088 と traceability spec で確定した canonical reference resolution foundation と investigation record integration を、Design Records MCP の contract・実装・tests に一貫して反映する。

## Impact scope

| layer | status | impact |
|---|---|---|
| decision | done | V01-ADR-087 / V01-ADR-088 accepted |
| design spec | done | active `spec:` / record ID / investigation canonical ref の concrete resolve tool / diagnostic contract を M19 Phase A で確定した |
| internal design | done | M19 Phase B0 で finalized public contract に追従する resolver / record parser / validation route を具体化した。MVP endpoint としては扱わない |
| implementation | done | M19 Phase B/C で investigation kind、kind-specific response、canonical resolver、validator を実装した |
| verification | done | M19 Phase D で investigation record / `spec:` ref / record ID-as-ref の参照解決・診断 tests を追加し、`go test ./internal/designrecords ./internal/designrecordsmcp` と `go test ./...` が通過した |

## Boundary

- 本 work item は横断進捗と影響範囲を所有する。
- 過去の具体的な実装順序、チェックリスト、完了 evidence の正本は legacy record `docs/tasks/m19-design-records-semantic-trace-support.md` が保持する。
- `V01-TASK-MCP-001-01` は、legacy `M19` の完了 evidence を current canonical workflow relation に接続する retrospective migration bridge であり、過去工程を再実行または再分解するものではない。
- `internal-design:` / `coverage:` / `COV-*` の resolve、semantic realization relation、coverage mapping query、MCP writer tools は本 work item の必須完了条件に含めない。

## Done condition

M19 の完了条件を満たし、Design Records MCP が V01-ADR-087 / V01-ADR-088 に基づく investigation integration と canonical reference resolution / validation を contract・implementation・tests の同一切替単位で提供できること。

## Migration bridge evidence

- 2026-05-27: Workflow artifact MCP validation の runtime verification により、旧 relation `tasks: M19` が `invalid_workflow_relation_target` として検出された。
- 2026-05-27: Legacy `M19` は historical completion evidence の正本として維持し、current workflow relation のみを `V01-TASK-MCP-001-01` に接続した。
- 2026-05-27: 本修正は closed M-series 全体の archive / migration ではなく、既に新形式で存在する `V01-WORK-MCP-001` の canonical relation clean-up に限定する。

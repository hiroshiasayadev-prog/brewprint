# TASK-MCP-001-01: Legacy M19 completion evidence を canonical task relation に接続する

- **id**: TASK-MCP-001-01
- **status**: done
- **date**: 2026-05-27
- **work_item**: WORK-MCP-001
- **source_requirement**: REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - WORK-MCP-001 の `tasks` relation を `TASK-*` ID-as-ref に整合させる migration bridge
  - Legacy `M19` completion evidence への追跡可能な接続

## Goal

ADR-091 / ADR-092 後の workflow relation contract に従い、既に完了している legacy M-series record `M19` の成果を `WORK-MCP-001` の canonical `TASK-*` relation へ接続する。

## Work

- `WORK-MCP-001.tasks` が legacy label `M19` を直接参照していた migration gap を確認する。
- `WORK-MCP-001.tasks` を本 task の `TASK-*` ID-as-ref へ更新する。
- 過去の具体作業・検証 evidence の正本が `docs/tasks/m19-design-records-semantic-trace-support.md` に残ることを明記する。

## Boundary

- 本 task は retrospective migration bridge であり、M19 の過去工程を新形式 task として再実行または再分解しない。
- Legacy `docs/tasks/m*.md` 全体の archive 化、open legacy record の分解、M-series migration 方針全体の確定は扱わない。
- `WORK-MCP-001` の capability completion 判定を変更せず、canonical relation の整合だけを修復する。

## Done condition

- `WORK-MCP-001.tasks` が `TASK-MCP-001-01` を canonical ID-as-ref として参照している。
- M19 が historical completion evidence の正本であることが失われず記録されている。
- `WORK-MCP-001.tasks = M19` に起因した `invalid_workflow_relation_target` の解消対象が明確になっている。

## Verification

- Design Records MCP の `validate_records` で `WORK-MCP-001.tasks = M19` 由来の `invalid_workflow_relation_target` が解消されることを確認する。
- `WORK-MCP-001` の status を `done` のまま維持し、過去実装の再完了を主張していないことを確認する。

## Evidence

- Legacy completion evidence source: `docs/tasks/m19-design-records-semantic-trace-support.md` (`M19`, `status: closed`, `closed_at: 2026-05-26`)。
- `M19` は `REQ-MCP-001` を実現する concrete implementation milestone であり、横断進捗は `WORK-MCP-001` が所有すると記録している。
- `WORK-MCP-003` の runtime verification により、`WORK-MCP-001.tasks = M19` が current validator で `invalid_workflow_relation_target` となる legacy migration gap として観測された。
- 本 task は、この既知 gap を canonical `TASK-*` relation に接続する最小 bridge として追加する。

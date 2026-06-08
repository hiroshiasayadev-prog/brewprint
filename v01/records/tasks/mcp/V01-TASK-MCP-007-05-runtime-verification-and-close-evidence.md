# V01-TASK-MCP-007-05: Runtime verification and close evidence

- **id**: V01-TASK-MCP-007-05
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-MCP-007
- **source_requirement**: V01-REQ-MCP-007
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-007-04
- **outputs**:
  - Runtime verification evidence
  - Close evidence for V01-WORK-MCP-007 / V01-REQ-MCP-007

## Goal

Implemented workflow artifact range navigation を runtime MCP call で確認し、V01-WORK-MCP-007 を close できる evidence を残す。

## Work

- Representative valid workflow artifact range query を実行する。
- Representative invalid mixed-domain / mixed-kind query を実行する。
- Existing ADR decision range query を実行し、regression がないことを確認する。
- Targeted tests の結果を evidence に残す。
- 必要なら V01-WORK-MCP-007 の status / close evidence と V01-REQ-MCP-007 の status を同期する。

## Done condition

- Valid workflow artifact range query が期待する records を返す。
- Invalid range query が contract 通りに拒否される。
- ADR decision range query が従来どおり動作する。
- Tests と runtime verification の結果が evidence として残っている。
- V01-WORK-MCP-007 / V01-REQ-MCP-007 の close 状態が整合している。

## Verification

- Design Records MCP runtime call 結果を確認する。
- Targeted Go tests の結果を確認する。
- `validate_records` で workflow relation の整合性を確認する。

## Evidence

2026-06-01 runtime verification result: done.

Pre-runtime verification already completed in `V01-TASK-MCP-007-04`:

- User ran `go test ./internal/designrecords ./internal/designrecordsmcp` on Windows PowerShell.
- Result:
  - `ok github.com/hiroshiasayadev-prog/brewprint/internal/designrecords 1.212s`
  - `ok github.com/hiroshiasayadev-prog/brewprint/internal/designrecordsmcp 0.475s`

Runtime MCP verification from active ChatGPT tool surface after rebuild / MCP reload:

Valid cases:

- `list_records(kind: work_item, id_range: V01-WORK-MCP-007..V01-WORK-MCP-007, order_by: id, order: asc, limit: 5)` returned `V01-WORK-MCP-007`.
- `list_records(kind: task, id_range: V01-TASK-MCP-007-01..V01-TASK-MCP-007-05, order_by: id, order: asc, limit: 10)` returned `V01-TASK-MCP-007-01` through `V01-TASK-MCP-007-05`.
- `list_records(id_range: V01-WORK-MCP-007..V01-WORK-MCP-007, order_by: id, order: asc, limit: 5)` returned `V01-WORK-MCP-007`, confirming omitted `kind` derives effective `work_item` kind from endpoint family.
- `validate_records(kind: task, id_range: V01-TASK-MCP-007-01..V01-TASK-MCP-007-05)` returned `ok: true` with no diagnostics.
- `list_records(kind: decision, id_range: V01-ADR-067..V01-ADR-077, order_by: id, order: asc, limit: 3)` returned `V01-ADR-067`, `V01-ADR-068`, `V01-ADR-069`, confirming ADR range compatibility.

Invalid cases:

- `list_records(kind: work_item, id_range: V01-WORK-DATA-001..V01-WORK-MCP-010)` returned `invalid_id_range` with message `workflow id_range endpoints must use the same domain`.
- `list_records(kind: task, id_range: V01-TASK-MCP-006-01..V01-TASK-MCP-007-05)` returned `invalid_id_range` with message `task id_range endpoints must use the same work sequence`.
- `list_records(id_range: SPEC-design-records-mcp-tools..SPEC-design-records-mcp-schema)` returned `invalid_id_range` with message `invalid id_range endpoint "SPEC-design-records-mcp-tools"`.

Close condition:

- Runtime MCP verification passed.
- `V01-WORK-MCP-007` and `V01-REQ-MCP-007` may now be closed if relation validation remains clean.

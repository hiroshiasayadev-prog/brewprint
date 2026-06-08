# V01-TASK-MCP-007-03: MCP tools spec update

- **id**: V01-TASK-MCP-007-03
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-MCP-007
- **source_requirement**: V01-REQ-MCP-007
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-007-02
- **outputs**:
  - Updated `docs/spec/design-records-mcp/tools.md`

## Goal

`V01-TASK-MCP-007-02` で決めた workflow artifact range filter contract を Design Records MCP tools spec に反映する。

## Work

- `list_records` request schema の range / filter field 定義を更新する。
- Supported workflow artifact ID family、domain、sequence comparison rules を明記する。
- Mixed-domain、mixed-kind、unsupported kind、malformed endpoint の扱いを明記する。
- ADR decision range の既存 behavior を維持する場合、その互換性を明記する。
- 必要なら Error handling section を更新する。
- `last_updated` と関連 source refs の必要性を確認する。

## Done condition

- Workflow artifact range navigation の入力、成功時 response、失敗時 response が public MCP contract として読める。
- `V01-REQ-MCP-007` の acceptance expectation と矛盾しない。
- Implementation task が spec を根拠に作業できる状態になっている。

## Verification

- Spec update diff を確認する。
- `id_range` の decision-only 既存記述が stale として残っていないことを確認する。

## Evidence

2026-06-01 spec update result: done.

Updated `docs/spec/design-records-mcp/tools.md`:

- `list_records.id_range` is no longer documented as decision-only.
- `id_range` now supports these endpoint families:
  - `ADR-NNN` for `decision`.
  - `REQ-<DOMAIN>-NNN` for `requirement`.
  - `WORK-<DOMAIN>-NNN` for `work_item`.
  - `TASK-<DOMAIN>-NNN-MM` for `task`.
- Requirement and work item ranges are defined as same-family + same-domain numeric sequence comparisons.
- Task ranges are defined as same-family + same-domain + same work sequence numeric task sequence comparisons.
- Omitted `kind` with `id_range` now derives effective `kind` from the endpoint family.
- Explicit `kind` must match endpoint family.
- One-sided workflow ranges are scoped by the provided endpoint family / domain / task work sequence.
- `SPEC-*` / `INV-*`, mixed family, mixed domain, mixed task work sequence, malformed endpoint, and endpoint family mismatch with explicit `kind` are request errors.
- Response section now refers to the request section's ID family rule for range membership.

Updated `validate_records` request contract:

- `validate_records.id_range` now shares the same endpoint family / effective `kind` / one-sided range / unsupported range behavior as `list_records`.
- This keeps `list_records` and `validate_records` tool behavior consistent.

Updated Error handling:

- Added `invalid_id_range` for malformed, unsupported, mixed, or kind-mismatched `id_range` requests.
- Kept `id_range_requires_decision_kind` as a legacy error code for the pre-REQ-MCP-007 decision-only boundary.
- New implementation should use `invalid_id_range`.

Verification:

- The spec now reflects the selected contract from `V01-TASK-MCP-007-02`.
- Implementation can proceed in `V01-TASK-MCP-007-04` against this public contract.

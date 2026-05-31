# TASK-MCP-007-01: Current list_records range contract gap review

- **id**: TASK-MCP-007-01
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-MCP-007
- **source_requirement**: REQ-MCP-007
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Current behavior evidence for `list_records.id_range`
  - Contract gap summary for workflow artifact range navigation

## Goal

現行の `list_records.id_range` contract、implementation、tests を確認し、workflow artifact range navigation を追加する前提差分を明確にする。

## Work

- `docs/spec/design-records-mcp/tools.md` の `list_records` request / error handling を確認する。
- `internal/designrecords` と `internal/designrecordsmcp` の `list_records` range handling を確認する。
- 既存 tests が ADR range、workflow artifact range rejection、invalid request をどう扱うか確認する。
- `REQ-MCP-007` の要求に対して、現行 contract と implementation の gap を分類する。

## Done condition

- 現行 `id_range` の supported / unsupported behavior が整理されている。
- Workflow artifact range support に必要な contract decision points が明示されている。
- 既存 behavior の regression guard として残すべき test case が分かっている。

## Verification

- Relevant spec / implementation / tests の確認結果を evidence として記録する。
- 実行可能なら targeted Go tests の現状結果を記録する。

## Evidence

2026-06-01 review result: done.

Read / inspected:

- `docs/spec/design-records-mcp/tools.md` `list_records` section.
- `docs/spec/design-records-mcp/tools.md` Error handling section.
- `internal/designrecords/tools.go`.
- `internal/designrecords/types.go`.
- `internal/designrecords/validation.go`.
- `internal/designrecords/list_records_test.go`.
- `internal/designrecords/validation_test.go`.
- `internal/designrecordsmcp/tools.go`.
- `internal/designrecordsmcp/tools_call.go`.
- `internal/designrecordsmcp/tools_call_test.go`.

Current spec contract:

- `list_records.id_range` is inclusive and supports one-sided `from` / `to`.
- `list_records.id_range` is explicitly limited to `ADR-NNN` decision records.
- If `kind` is omitted with `id_range`, the request behaves as `kind: decision`.
- `kind: spec` / `kind: investigation` / `kind: requirement` / `kind: work_item` / `kind: task` with `id_range` is request error.
- `SPEC-*` / `INV-*` / `REQ-*` / `WORK-*` / `TASK-*` range endpoints are request error.
- Current error code for this boundary is `id_range_requires_decision_kind`.
- `validate_records.id_range` mirrors the same decision-only range contract.

Current implementation behavior:

- `ListRecords` builds `listRecordsScope` via `newListRecordsScope`.
- `newListRecordsScope` rejects any `id_range` request with non-empty `kind` other than `decision`.
- `newListRecordsScope` parses endpoints using `parseDecisionIDRange`.
- `parseDecisionIDRange` delegates to `parseDecisionIDEndpoint`, which only accepts normalized `ADR-NNN` via `adrIDPattern`.
- When `id_range` is accepted, `scope.kind` is forced to `decision` and selection uses numeric ADR comparison only.
- MCP tool schema exposes only a generic `id_range` object with `from` / `to` strings; the decision-only rule is enforced by handler logic, not schema enum constraints.

Current test coverage:

- `TestListRecordsIDRangeFilter` covers inclusive endpoints, one-sided from, one-sided to, omitted kind behaving as decision, and explicit decision kind.
- `TestListRecordsRequestErrors` covers non-decision kind with `id_range`, `SPEC-*` endpoint, malformed ADR endpoint, invalid order, and invalid limit.
- `TestListRecordsRepositoryBootstrapQueries` covers repository ADR range and confirms workflow artifact kinds can be listed by `kind` without range.
- `TestValidateRecordsIDRangeFilter` and `TestValidateRecordsRequestErrors` cover the same decision-only boundary for validation scope.
- `TestToolsCallToolErrors` covers MCP transport error for `list_records` with `kind: spec` and `id_range`.
- Existing tests do not cover valid workflow artifact range behavior because current contract rejects it.

Runtime MCP observations:

- `list_records(id_range: ADR-067..ADR-077, order_by: id, order: asc, limit: 5)` returned ADR records starting at `ADR-067`; runtime behavior matches the ADR-only range contract.
- `list_records(kind: work_item, id_range: WORK-DATA-001..WORK-DATA-004)` returned tool error `id_range_requires_decision_kind` with message `id_range requires kind decision`.
- `list_records(id_range: WORK-DATA-001..WORK-DATA-004)` returned tool error `id_range_requires_decision_kind` with message `id_range endpoints must use ADR-NNN decision IDs`.
- `list_records(kind: work_item, order_by: id, order: asc, limit: 10)` returned `WORK-DATA-001` through `WORK-DATA-010`, showing workflow artifact listing works without range, but range narrowing is unavailable.

Gap classification:

- This is not an implementation bug against current spec. Spec, implementation, tests, and runtime behavior are aligned on decision-only `id_range`.
- This is a public contract gap / usability gap captured by `REQ-MCP-007`.
- `WORK-DATA-004` occupancy can be discovered by exact ID lookup or broad `kind: work_item` listing, but cannot be expressed as a safe bounded workflow artifact range query.

Decision points for `TASK-MCP-007-02`:

1. Whether to extend existing `id_range` to workflow artifact IDs or introduce a new workflow-specific filter field.
2. Whether `validate_records.id_range` should be extended together with `list_records`, or remain decision-only.
3. Requirement / work item range ordering can be same-kind + same-domain + 3-digit sequence numeric comparison.
4. Task range ordering needs an explicit unit: full `(domain, work_sequence, task_sequence)` ordering or same-domain + same-work-sequence only.
5. Error behavior needs names/messages for mixed-domain, mixed-kind, malformed workflow endpoint, unsupported `SPEC-*` / `INV-*`, and cross-family endpoints.

Regression guards to preserve:

- ADR `id_range` remains inclusive.
- One-sided ADR `from` and `to` continue to work.
- Omitted `kind` with ADR `id_range` continues to behave as `kind: decision`, unless explicitly superseded by a new contract.
- Non-range workflow listing by `kind` remains unchanged.
- Unsupported mixed or malformed ranges must not silently fall back to lexical ordering.

Go tests were not executed in this assistant session because no repo-local command execution tool is available here. Runtime behavior was verified through the active Design Records MCP tool surface.

# TASK-MCP-007-02: Workflow artifact range filter contract decision

- **id**: TASK-MCP-007-02
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-MCP-007
- **source_requirement**: REQ-MCP-007
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-007-01
- **outputs**:
  - Selected workflow artifact range filter contract
  - Rejected alternatives and rationale

## Goal

Workflow artifact range navigation を `list_records` public contract としてどう表現するかを決める。

## Work

- `list_records.id_range` を workflow artifact ID に拡張する案を評価する。
- Decision-only `id_range` を維持し、workflow artifact 専用 filter を追加する案を評価する。
- Requirement / work item / task それぞれの valid range grammar と ordering unit を定義する。
- Mixed-domain、mixed-kind、unsupported kind、malformed endpoint の error behavior を決める。
- `validate_records.id_range` に同様の拡張を入れるか、`list_records` に限定するかを判断する。

## Done condition

- 採用 contract が一つに決まっている。
- `REQ-MCP-007` が要求する workflow artifact navigation が満たせることを確認している。
- Cross-domain / mixed-kind comparison を silent reinterpretation しない拒否方針が決まっている。
- `TASK-MCP-007-03` が spec update に進める粒度で decision result が残っている。

## Verification

- Decision result が `REQ-MCP-007` の Acceptance expectation と矛盾しないことを確認する。
- `docs/spec/design-records-mcp/tools.md` の既存 `id_range` wording と衝突しない更新方針になっていることを確認する。

## Evidence

2026-06-01 decision result: done.

Selected contract:

- Extend existing `list_records.id_range` to support workflow artifact IDs.
- Do not introduce a separate workflow artifact range filter field in this work item.
- Preserve the existing ADR decision range behavior.
- `id_range` remains an inclusive range with optional one-sided `from` / `to` endpoints.
- When `id_range` is present and `kind` is omitted, the endpoint family determines the effective kind.
- When `kind` is present, it must match the endpoint family.

Supported endpoint families:

| family | kind | valid endpoint form | ordering unit |
|---|---|---|---|
| ADR | `decision` | `ADR-NNN` | numeric ADR sequence |
| requirement | `requirement` | `REQ-<DOMAIN>-NNN` | same domain + numeric requirement sequence |
| work item | `work_item` | `WORK-<DOMAIN>-NNN` | same domain + numeric work item sequence |
| task | `task` | `TASK-<DOMAIN>-NNN-MM` | same domain + same work sequence + numeric task sequence |

Task range rule:

- `TASK-<DOMAIN>-NNN-MM` ranges are valid only when both endpoints have the same `<DOMAIN>` and the same work sequence `NNN`.
- Example valid range: `TASK-MCP-007-01` .. `TASK-MCP-007-05`.
- Example invalid range: `TASK-MCP-006-01` .. `TASK-MCP-007-05`.
- This avoids defining global task ordering across work items.

One-sided workflow ranges:

- One-sided ranges are allowed for workflow families when `kind` is explicit.
- For one-sided workflow ranges with omitted `kind`, the provided endpoint determines the effective kind and domain / work sequence scope.
- Example: `kind: work_item`, `id_range.from: WORK-DATA-004` lists `WORK-DATA-*` records whose sequence is >= 004.
- Example: omitted `kind`, `id_range.from: WORK-DATA-004` behaves as `kind: work_item` scoped to `DATA`.

Rejected alternative:

- A new workflow-specific filter field was rejected for this work item.
- Reason: `id_range` already means ID range and can safely support multiple ID families when each family has explicit grammar and ordering rules.
- Adding a parallel field would create API duplication and require users to know which range field belongs to which artifact family.

Validation scope decision:

- Extend `validate_records.id_range` with the same endpoint parsing and selection semantics as `list_records.id_range`.
- Reason: `validate_records` already exposes `id_range`; leaving it decision-only while `list_records` supports workflow ranges would create inconsistent tool behavior.
- Validation should use the same valid family / same domain / same task work-sequence rules.

Error behavior decision:

- Unsupported `SPEC-*` and `INV-*` range endpoints remain invalid in this work item.
- Mixed families are invalid: e.g. `REQ-MCP-001` .. `WORK-MCP-001`.
- Mixed domains are invalid: e.g. `WORK-DATA-001` .. `WORK-MCP-010`.
- Mixed task work sequences are invalid: e.g. `TASK-MCP-006-01` .. `TASK-MCP-007-05`.
- Malformed endpoints are invalid.
- Invalid range requests must return a tool error and must not silently fall back to lexical ordering or broad listing.
- Existing error code `id_range_requires_decision_kind` should be replaced or generalized in spec / implementation, because the selected contract no longer means decision-only. Suggested direction: introduce `invalid_id_range` for malformed, mixed, or unsupported range endpoints. Exact naming can be finalized in `TASK-MCP-007-03` spec update.

Examples expected after implementation:

Valid:

- `list_records(id_range: ADR-067..ADR-077)` -> decision records in numeric ADR range.
- `list_records(kind: requirement, id_range: REQ-MCP-001..REQ-MCP-007)` -> MCP requirements in numeric sequence range.
- `list_records(kind: work_item, id_range: WORK-DATA-001..WORK-DATA-004)` -> DATA work items in numeric sequence range.
- `list_records(kind: task, id_range: TASK-MCP-007-01..TASK-MCP-007-05)` -> tasks for `WORK-MCP-007` in task sequence range.
- `list_records(id_range: WORK-DATA-004..WORK-DATA-004)` -> exact same-family range for `WORK-DATA-004`.

Invalid:

- `list_records(kind: work_item, id_range: WORK-DATA-001..WORK-MCP-010)`.
- `list_records(kind: requirement, id_range: REQ-DATA-001..REQ-MCP-010)`.
- `list_records(kind: task, id_range: TASK-MCP-006-01..TASK-MCP-007-05)`.
- `list_records(id_range: REQ-MCP-001..TASK-MCP-001-01)`.
- `list_records(id_range: SPEC-design-records-mcp-tools..SPEC-design-records-mcp-schema)`.
- `list_records(id_range: INV-DOCS-001..INV-DOCS-010)`.

Verification:

- This decision satisfies `REQ-MCP-007` because workflow artifacts can be listed by safe family/domain-scoped ID ranges without relying on filesystem listing.
- The decision preserves the existing ADR range behavior and narrows new workflow ordering to explicit same-family scopes.
- `TASK-MCP-007-03` can now update the public tools spec from this contract.

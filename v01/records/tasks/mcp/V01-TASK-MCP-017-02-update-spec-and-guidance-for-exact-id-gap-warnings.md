# V01-TASK-MCP-017-02: Update spec and guidance for exact ID gap warnings

- **id**: V01-TASK-MCP-017-02
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-017
- **source_requirement**: V01-REQ-MCP-018
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-MCP-017-01
- **outputs**:
  - Updated SPEC-design-records-mcp-tools exact ID create warning contract
  - Updated authoring guidance for new placeholder preference and exact ID warning

## Goal

Document the exact ID create gap-warning behavior in the public MCP tool contract and authoring guidance.

## Work

- Update the Design Records MCP tools spec around propose_record_create exact ID and new placeholder behavior.
- Document that the warning is non-blocking and does not prohibit exact ID create.
- Add guidance that workflow artifact creation should prefer server-side new placeholders unless an exact ID is intentionally required.
- Ensure TASK warning scope is described as parent work item scope.

## Done condition

Spec and guidance describe the warning contract clearly enough for implementation and caller behavior.

## Verification

- Review updated spec section for consistency with V01-REQ-MCP-018.
- Validate affected design records.

## Evidence
Updated public contract and authoring guidance for exact ID gap warnings.

Files updated:

- `docs/spec/design-records-mcp/tools.md`
- `docs/guides/requirement-authoring.md`
- `docs/guides/work-item-authoring.md`
- `docs/guides/task-authoring.md`

Summary:

- `propose_record_create` now documents non-blocking `exact_id_sequence_gap` info diagnostics for workflow artifact exact ID create.
- REQ and WORK gap warning scope is same kind plus same domain.
- TASK gap warning scope is same domain plus parent work item sequence.
- ADR / decision create is explicitly outside this warning scope.
- `new` placeholder create must not emit the warning.
- Exact ID create that fills an existing gap must not emit the warning, while existing duplicate ID checks remain unchanged.
- Requirement, work item, and task authoring guides now recommend server-side `new` placeholders unless a specific exact ID is intentional.

Claude Code review result:

- Verdict: OK with minor fixes.
- `docs/spec/diagnostics.md` update is not required because `exact_id_sequence_gap` belongs to the Design Records MCP authoring response surface, not brewprint DSL semantic validation diagnostics.
- V01-TASK-MCP-017-02 can remain done.
- V01-TASK-MCP-017-03 can start after clarifying the proposal-level diagnostic field.

Review fixes applied:

- Clarified that `exact_id_sequence_gap` is returned in proposal-level `diagnostics`, not `validation.diagnostics`.
- Added an exact ID sequence-gap warning response example.
- Normalized work item authoring guide wording from current-tense `skips` to conditional `would skip`.

Verification performed:

- Reviewed updated spec and guidance text against V01-REQ-MCP-018 and V01-TASK-MCP-017-01 evidence.
- Design Records MCP validation was run after metadata updates for V01-TASK-MCP-017-02 and V01-WORK-MCP-017.

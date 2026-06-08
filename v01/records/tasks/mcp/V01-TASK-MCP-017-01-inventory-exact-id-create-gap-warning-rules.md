# V01-TASK-MCP-017-01: Inventory exact ID create gap-warning rules

- **id**: V01-TASK-MCP-017-01
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-017
- **source_requirement**: V01-REQ-MCP-018
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Exact ID gap-warning rule inventory
  - Decision on REQ / WORK / TASK warning scope

## Goal

Define the proposal-time gap-warning rules for workflow artifact exact ID create requests.

## Work

- Inspect current `propose_record_create` exact ID and `*-new` allocation behavior.
- Classify sequence scopes for REQ, WORK, and TASK create.
- Decide whether REQ exact ID create should warn in addition to WORK / TASK.
- Document warning trigger examples and non-trigger examples.

## Done condition

The implementation task has a concrete rule for when to emit a non-blocking warning and what scope is used for sequence comparison.

## Verification

- Review `SPEC-design-records-mcp-tools` create contract.
- Review current authoring implementation and tests before patching.

## Evidence
Claude Code review completed V01-TASK-MCP-017-01 inventory and found no blocking issue.

Summary:

- Current `*-new` allocation uses max existing sequence plus one and does not fill existing gaps.
- Current exact ID create returns the requested exact ID after format/domain/parent/duplicate checks and performs no gap warning.
- Proposed warning rule: emit a non-blocking info diagnostic when `requestedSeq > maxSeq + 1`.
- REQ and WORK warning scope is domain-scoped sequence.
- TASK warning scope is parent work item scoped sequence.
- ADR is outside scope because V01-REQ-MCP-018 targets workflow artifacts.
- `*-new` placeholder create must skip gap-warning checks.
- Existing gaps filled by exact ID create should not warn unless the target already exists, which remains covered by existing duplicate checks.
- Implementation must make `ProposeRecordCreate` diagnostics severity-aware, because current behavior fails proposal creation when any diagnostic is present.

Recommended implementation touch points:

- Add `exact_id_sequence_gap` diagnostic category.
- Change `ProposeRecordCreate` gating from any diagnostic to error diagnostics only.
- Add exact ID gap detection after ID resolution and before proposal persistence.
- Add regression tests for REQ, WORK, TASK warning and non-warning cases, plus placeholder and ADR non-scope cases.

Verdict: V01-TASK-MCP-017-01 done. V01-TASK-MCP-017-02 / V01-TASK-MCP-017-03 can proceed from this rule set.

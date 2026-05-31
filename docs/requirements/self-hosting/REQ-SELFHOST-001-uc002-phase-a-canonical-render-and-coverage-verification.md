# REQ-SELFHOST-001: UC-002 Phase A canonical render and coverage verification

- **id**: REQ-SELFHOST-001
- **status**: captured
- **date**: 2026-05-31
- **source_refs**:
  - ADR-057
  - REQ-DATA-001
  - WORK-DATA-001
  - REQ-RESOLVE-001
  - WORK-RESOLVE-001
- **work_items**:
  - WORK-SELFHOST-001

## Requirement

UC-002 Phase A MCP public contract YAML needs to be validated and rendered as the current canonical baseline after the M15 v1.1 data-layer work and the file-private sub node resolver fix.

The Phase A YAML is already present, but the canonical `renders/` output and render review evidence remain incomplete. The project needs a current verification point before M14 self-hosting can safely continue into internal layer blueprinting.

## Evidence

M14 recorded Phase A as the MCP public contract blueprinting stage and left "Phase A render" open. UC-002 coverage currently states that the Phase A YAML is placed, while `brewprint render`, `go test ./...`, and generated render review are still unconfirmed.

`REQ-DATA-001` / `WORK-DATA-001` closed the M15 minimum-expressiveness boundary, including the initial enum migration used by UC-002. `REQ-RESOLVE-001` / `WORK-RESOLVE-001` resolved the UC-002 duplicate task QID / unresolved flow task issue that previously blocked full validate / render.

## Required Outcome

- UC-002 Phase A YAML validates under the current spec and implementation baseline.
- UC-002 Phase A canonical renders are generated under the UC-002 `renders/` directory and reviewed as generated artifacts.
- UC-002 Phase A coverage and TASKS-UC-002 gap log are synchronized with the current M15 / RESOLVE outcomes.
- Any remaining Phase A render, coverage, or model expressiveness gaps are classified without reopening M14a or M15.

## Explicitly Excluded Scope

- M14 Phase B internal layer blueprinting.
- M14 close.
- M14a reopening.
- M15 / `v1.1.0-spec` reopening.
- New data-layer capability work such as helper model, tagged union, or DAG TypeRef hint.
- MCP semantic identity / state-machine identity migration.

## Boundary

This requirement captures the need for a Phase A verification baseline. It does not own concrete commands, generated output diffs, coverage edits, short task ordering, or close evidence. Those belong to the linked work item and its future task artifacts.


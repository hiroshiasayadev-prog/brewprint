# V01-WORK-DATA-003: Resolve model-file helper render boundary

- **id**: V01-WORK-DATA-003
- **status**: done
- **date**: 2026-05-31
- **source_requirement**: V01-REQ-DATA-002
- **impact_refs**:
  - V01-INV-DATA-002
  - V01-ADR-070
  - V01-ADR-072
  - V01-ADR-073
  - V01-ADR-075
- **tasks**:
  - V01-TASK-DATA-003-01
  - V01-TASK-DATA-003-02
  - V01-TASK-DATA-003-03
  - V01-TASK-DATA-003-04

## Goal

Resolve the model-file helper model and render exposure boundary separately from the task-file helper minimum path.

This work item exists to prevent model response helper migration from being pulled into V01-WORK-DATA-002 before model-file render exposure is decided.

V01-WORK-DATA-002 explicitly sends V01-ADR-072 catalog follow-up, V01-ADR-075 model file render resolution, model-file helper render exposure, and UC-002 model response helper-shape migration review to this work item. Actual UC-002 migration is delegated to V01-WORK-DATA-004 or later follow-up work.

## Boundary

### Included

- Receive deferred scope from V01-WORK-DATA-002 after the Option A task-file helper model minimum boundary.
- Decide whether V01-ADR-075 should be accepted as-is, revised, or split.
- Decide whether V01-ADR-072 model / schema catalog follow-up is needed before or after model-file render exposure.
- Decide the minimum model-file helper render capability needed to satisfy V01-ADR-070's human-visible render constraint for model-file helper models.
- Determine whether the V01-ADR-075 dependency on V01-ADR-073 is intrinsic to model-file render, or only needed for tagged-union rendering details.
- Classify UC-002 model response helper-shape candidates and delegate actual migration to V01-WORK-DATA-004 or later follow-up work.
- Track any required spec, render, YAML, fixture, and verification follow-up for model-file helper render.

### Excluded

- V01-WORK-DATA-002 Option A task-file helper minimum.
- V01-ADR-071 task-file DAG Markdown private model exposure, except as contrast.
- Implementing V01-ADR-073 tagged union model.
- Implementing V01-ADR-074 DAG asset TypeRef hint.
- Implementing V01-ADR-078 MCP semantic identity or helper model MCP exposure schema.
- M15 / v1.1.0-spec reopening.
- V01-REQ-DATA-001 / V01-WORK-DATA-001 edits.

## Impact Scope

| layer | current state | handling in this work item |
|---|---|---|
| source requirement | V01-REQ-DATA-002 captured | This work item owns the model-file helper render sub-boundary |
| helper model decision | V01-ADR-070 accepted | Keep helper models human-visible; do not migrate model-file helpers without render exposure |
| catalog decision | V01-ADR-072 accepted | Treat model catalog as opt-in curated view, not automatic model-file render |
| model file render decision | V01-ADR-075 proposed | Resolve before implementation or migration |
| tagged union | V01-ADR-073 proposed | Do not implement here; classify whether V01-ADR-075 depends on it intrinsically |
| UC-002 migration | V01-INV-DATA-002 model response candidates exist | Classify candidates here; delegate actual migration to V01-WORK-DATA-004 or later follow-up work |

## Task flow

```mermaid
flowchart TD
  T1["V01-TASK-DATA-003-01: V01-ADR-075 dependency and split review"]
  T2["V01-TASK-DATA-003-02: Model-file helper render spec boundary"]
  T3["V01-TASK-DATA-003-03: Model-file render implementation and verification"]
  T4["V01-TASK-DATA-003-04: UC-002 model response helper candidate review"]
  T1 --> T2 --> T3 --> T4
```

## Task Candidates

- `V01-TASK-DATA-003-01`: V01-ADR-075 dependency and split review.
- `V01-TASK-DATA-003-02`: Model-file helper render spec boundary.
- `V01-TASK-DATA-003-03`: Model-file render implementation and verification.
- `V01-TASK-DATA-003-04`: UC-002 model response helper candidate review.

`V01-TASK-DATA-003-01`, `V01-TASK-DATA-003-02`, `V01-TASK-DATA-003-03`, and `V01-TASK-DATA-003-04` are created task artifacts and are listed in metadata `tasks`.

## Completion Condition

This work item can be marked `done` when model-file helper render exposure is decided, reflected in the relevant spec and renderer behavior, verified with fixture / golden evidence, and UC-002 model response helper candidates are classified without pulling in tagged union, DAG TypeRef hint, MCP identity, M15 reopen scope, or actual UC-002 migration.

## Close Evidence

- V01-TASK-DATA-003-01 is done: V01-ADR-075 dependency and split review completed.
- V01-TASK-DATA-003-02 is done: model-file render minimum spec alignment completed.
- V01-TASK-DATA-003-03 is done: model-file render implementation, fixture / golden update, and verification completed.
- V01-TASK-DATA-003-04 is done: UC-002 model response helper candidates classified.
- Model-file helper render exposure has been decided, specified, implemented, and verified by the DATA-003 task chain.
- UC-002 response helper-shape candidates are classified into model-file helper migration candidates, V01-REQ-DATA-003 / V01-WORK-DATA-004 wait candidates, and unchanged candidates.
- Actual UC-002 YAML migration is delegated to V01-WORK-DATA-004 or later follow-up work.
- Tagged union rendering, DAG TypeRef hint, MCP identity, M15 reopen scope, V01-WORK-DATA-002 reopening, and actual UC-002 migration were not pulled into this work item.
- No UC-002 YAML or render output change is required to close V01-WORK-DATA-003.

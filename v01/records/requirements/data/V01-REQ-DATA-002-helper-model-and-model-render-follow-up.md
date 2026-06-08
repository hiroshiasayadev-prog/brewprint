# V01-REQ-DATA-002: Helper model and model render follow-up for UC-002 contract shapes

- **id**: V01-REQ-DATA-002
- **status**: captured
- **date**: 2026-05-31
- **source_refs**:
  - V01-INV-DATA-001
  - V01-INV-DATA-002
  - V01-REQ-DATA-001
  - V01-WORK-DATA-001
  - V01-ADR-070
  - V01-ADR-071
  - V01-ADR-072
  - V01-ADR-075
- **work_items**:
  - V01-WORK-DATA-002
  - V01-WORK-DATA-003
  - V01-WORK-DATA-005
  - V01-WORK-DATA-006
  - V01-WORK-DATA-008
  - V01-WORK-DATA-009
  - V01-WORK-DATA-011
  - V01-WORK-DATA-012

## Requirement

UC-002 still contains MCP contract response and helper shapes that are represented with `any + note` after the M15 minimum-expressiveness release. The project needs a follow-up requirement for the V01-ADR-070 / V01-ADR-071 / V01-ADR-072 / V01-ADR-075 dependency chain: file-private helper model capability and the human-readable render surfaces needed to keep helper models visible.

## Evidence

`V01-REQ-DATA-001` / `V01-WORK-DATA-001` explicitly excluded helper model, private model render, model catalog, and model file render from the M15 / `v1.1.0-spec` close boundary.

`V01-INV-DATA-001` and `V01-INV-DATA-002` identify helper / nested response shapes in UC-002 as a deferred data-layer follow-up. V01-ADR-070 accepted the file-private helper model direction, V01-ADR-071 accepted task-file helper render exposure, V01-ADR-072 accepted model / schema catalog view, and V01-ADR-075 remains proposed for model file render.

## Required Outcome

- The helper model / model render follow-up boundary is split into implementable task-sized work.
- The project decides which part of V01-ADR-070 / 071 / 072 / 075 is the first practical capability to implement.
- UC-002 helper-shape migration candidates are identified without turning all notes retreat debt into one blocker.
- Any spec, implementation, render, YAML, fixture, and verification work is tracked through a work item and future task artifacts.

## Explicitly Excluded Scope

- Tagged union / discriminator payload representation from V01-ADR-073.
- DAG asset TypeRef hint from V01-ADR-074.
- MCP semantic identity and state-machine identity from V01-ADR-078 / V01-ADR-079 / V01-ADR-080.
- M15 / `v1.1.0-spec` reopening.
- Reclassifying the completed `V01-REQ-DATA-001` F1 boundary.

## Boundary

This requirement captures the deferred helper-model and model-render chain. It does not decide the exact implementation sequence, spec wording, or UC-002 migration set. Those are owned by the linked work item and future tasks.

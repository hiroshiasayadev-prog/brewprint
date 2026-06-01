# TASK-DATA-010-02: ADR-073 revision and spec / diagnostics alignment

- **id**: TASK-DATA-010-02
- **status**: todo
- **date**: 2026-06-01
- **work_item**: WORK-DATA-010
- **source_requirement**: REQ-DATA-004
- **estimate**: 1d-2d
- **depends_on**:
  - TASK-DATA-010-01
- **outputs**:
  - Revised and accepted ADR-073 or explicit blocker note
  - Tagged union model spec alignment
  - Tagged union TypeRef compatibility spec alignment
  - Tagged union diagnostic spec alignment
  - Follow-up input for implementation / fixture task

## Goal

Revise ADR-073 so that it can be accepted under the current REQ-DATA-004 / WORK-DATA-010 successor boundary, then align the active specs for the accepted tagged union / discriminator payload minimum.

This task exists because TASK-DATA-010-01 concluded that ADR-073 does not need a conceptual split, but should not be accepted as-is while it still contains stale M15 wording and while active specs do not yet define `kind: tagged_union`.

## Work

- Review TASK-DATA-010-01 evidence before editing.
- Revise ADR-073 to remove or update stale M15 / Phase C / close-scope wording.
- Make REQ-DATA-004 / WORK-DATA-010 the active acceptance boundary for ADR-073.
- Move ADR-073 from `proposed` to `accepted` only if the revised text cleanly preserves the minimum tagged union contract and excluded scope.
- Ensure ADR-073 acceptance does not imply tagged union render support is already implemented or that all remaining ADR-075 proposed text is accepted.
- Update `docs/spec/nodes.md` for the accepted tagged union model shape:
  - `model.kind: tagged_union`
  - required `discriminator`
  - required `variants`
  - required unique `variants[].tag`
  - `variants[].fields`
  - `fields: []` semantics for payload-less variants
  - discriminator field exclusion from variant payload fields
  - variant field subset rules for `name` / `type` / `note`
- Update `docs/spec/type-ref.md` for tagged union TypeRef behavior:
  - tagged union model is referenced as a named model TypeRef
  - no inline union / tagged union TypeRef syntax is introduced
  - tagged union compatibility is nominal like other non-container named models
  - `any` compatibility follows existing ADR-060 behavior
- Update `docs/spec/diagnostics.md` for tagged union definition diagnostics and chaining guidance:
  - `invalid_tagged_union_model`
  - `duplicate_variant_tag`
  - `invalid_variant_field`
  - interaction with existing `invalid_model_kind`, `duplicate_model_field`, `invalid_type_ref`, `unresolved_field_type`, and `unresolved_model`
- Record implementation follow-up input for parser / resolver / validation / render / fixtures.

## Included Scope

- ADR-073 wording revision and status update if acceptance conditions are met.
- Spec alignment for model shape, TypeRef compatibility, and diagnostics.
- Clarifying that model-file render minimum is already implemented for non-tagged-union kinds, while tagged union render remains separate follow-up scope.
- Follow-up task input for implementation / fixtures.

## Excluded Scope

- Parser implementation.
- Resolver implementation.
- Validation implementation.
- Renderer implementation for tagged union discriminator / variants.
- Model catalog implementation.
- UC-002 YAML migration.
- Fixture / golden regeneration.
- Go tests or render command execution unless an implementation agent later takes over.
- DAG asset TypeRef hint from ADR-074.
- MCP semantic identity / state machine identity from ADR-078 / ADR-079 / ADR-080.
- UC-002 duplicate task QID / unresolved flow task repair.
- Broad remaining UC-002 notes-retreat cleanup.
- Reopening M15, WORK-DATA-001, WORK-DATA-002, WORK-DATA-003, or WORK-DATA-004.

## Done Condition

- ADR-073 is either accepted with updated successor-boundary wording or left proposed with an explicit blocker recorded.
- Active specs describe the tagged union model minimum consistently enough for implementation to begin.
- Tagged union TypeRef compatibility is defined without adding inline union syntax.
- Tagged union diagnostics are listed with enough boundary guidance to implement validation without diagnostic category ambiguity.
- The task records what should be handed to the next implementation / fixture task.
- Excluded implementation, render, UC-002 YAML, fixture, golden, and unrelated DATA / MCP work remain untouched.

## Verification

- Confirm ADR-073 no longer implies M15 / WORK-DATA-003 reopening.
- Confirm ADR-073 acceptance, if performed, does not accept ADR-075 tagged-union render details as already implemented.
- Confirm `docs/spec/nodes.md`, `docs/spec/type-ref.md`, and `docs/spec/diagnostics.md` are mutually consistent.
- Confirm spec updates do not introduce inline union TypeRef syntax, external discriminator support, untagged union, general oneOf, scalar union, or runtime payload validation.
- Confirm no parser / resolver / validation / renderer / UC-002 YAML / fixture / golden changes are performed in this task.
- Confirm follow-up implementation task input is concrete enough to create the next task.

## Evidence

Not started.

# V01-TASK-DATA-010-01: V01-ADR-073 acceptance / split review

- **id**: V01-TASK-DATA-010-01
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-DATA-010
- **source_requirement**: V01-REQ-DATA-004
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - V01-ADR-073 acceptance / revision / split recommendation
  - Tagged union minimum contract boundary input
  - Follow-up task input for spec / diagnostics alignment

## Goal

Decide whether V01-ADR-073 can be accepted as the tagged union / discriminator payload contract for V01-WORK-DATA-010, or whether it must be revised, split, or deferred before implementation.

This task exists because V01-WORK-DATA-010 must not proceed into spec or implementation work while V01-ADR-073 is still only proposed and while the minimum tagged-union contract boundary is not confirmed.

## Work

- Review V01-ADR-073 as the proposed tagged union model decision.
- Review V01-REQ-DATA-004 to confirm the requirement boundary and explicitly excluded scope.
- Review V01-INV-DATA-002 tagged-union candidate inventory only enough to confirm the evidence used by V01-ADR-073.
- Review prior task evidence from V01-TASK-DATA-003-04 and V01-TASK-DATA-005-02 only enough to confirm that tagged union work was intentionally deferred to V01-WORK-DATA-010.
- Check whether V01-ADR-073's proposed contract is internally coherent for:
  - `model.kind: tagged_union`
  - required `discriminator`
  - required `variants`
  - variant tag uniqueness
  - discriminator field exclusion from variant payload fields
  - variant field subset rules
  - named model TypeRef usage
  - nominal compatibility
  - non-goals for untagged union / general oneOf / external discriminator
  - runtime payload validation non-goal
  - validation diagnostic boundary
- Decide whether V01-ADR-073 should be:
  - accepted as-is,
  - revised before acceptance,
  - split into a smaller accepted minimum plus follow-up decision,
  - or deferred.
- Identify the minimum follow-up task set needed after this review, without implementing spec, validation, render, fixture, or YAML changes.

## Decision Questions

1. Is V01-ADR-073's tagged union contract small enough to accept for the V01-WORK-DATA-010 minimum?
2. Are discriminator payload validation and runtime payload validation clearly separated?
3. Does V01-ADR-073 avoid pulling in DAG TypeRef hint, MCP identity, UC-002 duplicate task QID repair, or broad notes-retreat cleanup?
4. Are UC-002 tagged-union candidates sufficient as evidence without treating their old YAML shape as a fixed implementation contract?
5. Does any part of V01-ADR-073 need to be split before spec / implementation tasks are created?

## Included Scope

- V01-ADR-073 acceptance / revision / split review.
- Tagged union minimum contract boundary classification.
- Requirement and prior evidence alignment review.
- Follow-up recommendation for V01-WORK-DATA-010 task flow.

## Excluded Scope

- Editing V01-ADR-073.
- Accepting V01-ADR-073 by changing its status.
- Updating specs.
- Implementing parser / resolver / validation / renderer support.
- Updating UC-002 YAML.
- Regenerating fixtures or golden output.
- Implementing DAG TypeRef hint from V01-ADR-074.
- Implementing MCP semantic identity / state machine identity from V01-ADR-078 / V01-ADR-079 / V01-ADR-080.
- Repairing UC-002 duplicate task QID / unresolved flow task issues.
- Reopening M15, V01-WORK-DATA-001, V01-WORK-DATA-002, V01-WORK-DATA-003, or V01-WORK-DATA-004.

## Done Condition

- V01-ADR-073 is classified as accept-as-is, revise-before-acceptance, split-before-acceptance, or defer.
- Any blocking ambiguity in the tagged union minimum contract is listed.
- The boundary between schema-level tagged union validation and runtime payload validation is confirmed or flagged.
- Follow-up task input for spec / diagnostics alignment is concrete enough to create the next task.
- No spec, implementation, renderer, YAML, fixture, golden, or test changes are performed by this task.

## Verification

- Confirm this task only records review and decision input.
- Confirm V01-REQ-DATA-004 excluded scope remains excluded.
- Confirm V01-WORK-DATA-010 remains the owning work item.
- Confirm V01-ADR-073 status is not changed by this task.
- Confirm no UC-002 YAML or generated render artifacts are changed.

## Evidence

Reviewed on 2026-06-01.

### Sources reviewed

- `V01-REQ-DATA-004`: confirms tagged union / discriminator payload support is now owned by V01-WORK-DATA-010 and must stay separate from helper model render, DAG TypeRef hint, MCP identity, UC-002 duplicate task QID repair, broad notes-retreat cleanup, and reopening earlier DATA work items.
- `V01-WORK-DATA-010`: confirms this work item must review V01-ADR-073 before spec / implementation work and must not mix tagged union with unrelated DATA / MCP follow-ups.
- `V01-ADR-073`: reviewed proposed tagged union contract, acceptance checks, spec impact, implementation impact, UC-002 migration boundary, render/catalog impact, and non-goals.
- `V01-INV-DATA-002`: reviewed tagged-union candidate inventory and related non-covered cases.
- `V01-TASK-DATA-003-04`: confirms tagged union / kind-specific response payload candidates were intentionally left unchanged by V01-WORK-DATA-003.
- `V01-TASK-DATA-005-02`: confirms V01-ADR-073 was classified as DATA expressiveness requiring its own requirement and successor work item.
- `docs/spec/nodes.md`: confirms current active model kinds are `struct` / `list` / `dict` / `enum`, with no active `tagged_union` spec yet.
- `docs/spec/type-ref.md`: confirms named model TypeRef compatibility already treats list/dict specially and all other named models nominally; this is compatible with V01-ADR-073's tagged-union TypeRef proposal.
- `docs/spec/diagnostics.md`: confirms enum and TypeRef diagnostics exist, but tagged-union-specific diagnostics are not yet present.
- `docs/spec/views/model-file.md`: confirms model-file render minimum is specified for `struct` / `enum` / `list` / `dict`, while tagged union discriminator / variants rendering remains explicitly deferred from the V01-WORK-DATA-003 scope.
- `V01-TASK-DATA-003-03`: confirms model-file render minimum is implemented and verified, including same-file private helper model visibility, placement, indexes, fixtures, and tests; it also confirms tagged union discriminator / variants rendering was explicitly excluded.
- `V01-ADR-067`: reviewed enum precedent for named model TypeRef, nominal compatibility, and runtime literal validation non-goal.
- `V01-ADR-072`: reviewed accepted model catalog decision, including `tagged_union_models` as a planned include flag.
- `V01-ADR-075`: reviewed proposed model-file render decision; it includes tagged-union rendering detail, but that detail was not part of the implemented V01-WORK-DATA-003 model-file render minimum.

### Review result

Verdict: `revise-before-acceptance`.

V01-ADR-073 does not need a conceptual split before V01-WORK-DATA-010 continues. Its core tagged union contract is coherent enough to serve as the minimum for the next spec / diagnostics task:

- `model.kind: tagged_union` is a model-layer expressiveness feature and fits V01-REQ-DATA-004.
- `discriminator` and `variants` are mandatory and small enough for an initial schema-level contract.
- Variant `tag` uniqueness, empty-field variants, and discriminator-field exclusion are clear validation targets.
- Variant payload fields reuse the existing struct field subset of `name` / `type` / `note`, avoiding ER-specific metadata in payload variants.
- Use-site remains existing named model TypeRef; no inline union TypeRef is introduced.
- Compatibility is nominal like enum / struct named models, with list/dict remaining the only special container normalization.
- Runtime payload validation remains a non-goal and is separated from tagged-union model-definition validation.
- Untagged union / general oneOf / scalar union / external discriminator remain explicit non-goals.

However, V01-ADR-073 should not be accepted as-is because several acceptance / impact statements are stale relative to the current successor-work boundary:

1. V01-ADR-073 still frames the decision around `M15 Phase C` and `v1.1` initial adoption. V01-WORK-DATA-010 is now a dedicated successor work item, so acceptance wording should be updated to avoid implying M15 reopening.
2. V01-ADR-073's acceptance checklist asks whether the impact fits M15 Phase C. That is no longer the controlling boundary; the controlling boundary is V01-REQ-DATA-004 / V01-WORK-DATA-010.
3. Model-file render minimum itself is already implemented and verified by V01-TASK-DATA-003-03 for `struct` / `enum` / `list` / `dict` and same-file private helper models. However, tagged union discriminator / variants rendering was explicitly excluded from that implementation and remains deferred to V01-ADR-073 / V01-WORK-DATA-010 or a later render follow-up. Therefore accepting V01-ADR-073 should not imply that tagged union render behavior is already implemented, nor should it implicitly accept all remaining V01-ADR-075 proposed text.
4. V01-ADR-073 says initial migration target is `analyze_impact_change`. That is still a reasonable first fixture candidate, but acceptance should phrase it as V01-WORK-DATA-010 minimum candidate, not as M15 close scope.
5. V01-ADR-073's diagnostic proposal is acceptable as a starting point, but spec alignment should decide exact placement and chaining behavior with existing `invalid_model_kind`, `duplicate_model_field`, `invalid_type_ref`, `unresolved_field_type`, and `unresolved_model` diagnostics.

### Decision question answers

1. Is V01-ADR-073's tagged union contract small enough to accept for the V01-WORK-DATA-010 minimum?
   - Yes, after wording revision. No split required for the core contract.

2. Are discriminator payload validation and runtime payload validation clearly separated?
   - Yes. V01-ADR-073 defines model-definition validation and explicitly excludes runtime request / response payload validation. This separation should be preserved in spec wording.

3. Does V01-ADR-073 avoid pulling in DAG TypeRef hint, MCP identity, UC-002 duplicate task QID repair, or broad notes-retreat cleanup?
   - Mostly yes. It explicitly excludes untagged union / general oneOf / external discriminator and treats `diagnostic.related` as out of scope. V01-REQ-DATA-004 further excludes DAG TypeRef hint, MCP identity, duplicate task QID repair, and broad cleanup. The revised ADR should point to that successor boundary.

4. Are UC-002 tagged-union candidates sufficient as evidence without treating old YAML as fixed implementation contract?
   - Yes. V01-ADR-073 already states that UC-002 examples are evidence of the pattern rather than fixed implementation shape. V01-INV-DATA-002 supports N-001 / N-003 as direct discriminator-based request / response payload candidates, and N-021 / N-026 / N-027 as broader kind-specific payload candidates.

5. Does any part of V01-ADR-073 need to be split before spec / implementation tasks are created?
   - No conceptual split is required. The next task should update V01-ADR-073 wording/status and then align specs/diagnostics. Render/catalog details should remain spec-task scope or later render task scope, not a blocker for accepting the core tagged-union model contract.

### Recommended follow-up

Create or perform the next task as `V01-TASK-DATA-010-02: V01-ADR-073 revision and spec / diagnostics alignment`.

Recommended work for the next task:

- Revise V01-ADR-073 from `proposed` toward `accepted` after removing stale M15-specific acceptance wording.
- Make V01-REQ-DATA-004 / V01-WORK-DATA-010 the active boundary.
- Ensure acceptance does not imply tagged union render support is already implemented or that all remaining V01-ADR-075 proposed text is accepted.
- Update `docs/spec/nodes.md` to include `kind: tagged_union`, `discriminator`, `variants`, and variant field rules.
- Update `docs/spec/type-ref.md` to state tagged union is a named model TypeRef and nominally compatible like non-container named models.
- Update `docs/spec/diagnostics.md` to add tagged-union diagnostics and chaining guidance.
- Defer actual parser / resolver / validation / render implementation and UC-002 YAML migration to later tasks.

### Verification result

- This task performed review and decision input only.
- V01-REQ-DATA-004 excluded scope remains excluded.
- V01-WORK-DATA-010 remains the owning work item.
- V01-ADR-073 status was not changed by this task.
- No spec, implementation, renderer, UC-002 YAML, fixture, golden, or test changes were performed by this task.
- No Opus review is required before proceeding, because the remaining issue is a bounded wording/status/spec-alignment follow-up rather than an unresolved design split.

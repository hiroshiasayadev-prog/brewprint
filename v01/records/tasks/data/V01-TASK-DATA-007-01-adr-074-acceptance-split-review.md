# V01-TASK-DATA-007-01: V01-ADR-074 acceptance / split review

- **id**: V01-TASK-DATA-007-01
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-DATA-007
- **source_requirement**: V01-REQ-DATA-005
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - V01-ADR-074 acceptance / revision / split recommendation
  - DAG asset TypeRef hint render boundary input
  - Follow-up task input for DAG view spec alignment

## Goal

Decide whether V01-ADR-074 can be accepted as the DAG asset TypeRef hint render contract for V01-WORK-DATA-007, or whether it must be revised, split, or deferred before spec / renderer / fixture work starts.

This task exists because V01-WORK-DATA-007 must not proceed into view spec alignment or renderer changes while V01-ADR-074 is still only proposed and while the exact TypeRef hint render boundary is not confirmed.

## Work

- Review V01-ADR-074 as the proposed DAG asset node TypeRef hint decision.
- Review V01-REQ-DATA-005 to confirm the requirement boundary and explicitly excluded scope.
- Review V01-WORK-DATA-007 to confirm this task remains a focused DATA render / view successor.
- Review V01-TASK-DATA-005-02 only enough to confirm that V01-ADR-074 was intentionally classified as separate DATA render / view support.
- Check whether V01-ADR-074's proposed contract is internally coherent for:
  - asset node label format,
  - top-level TypeRef hint display,
  - primitive / named model / inline list / inline dict hint behavior,
  - named list / dict model behavior,
  - local id versus shortened QID fallback,
  - params boundary asset inclusion,
  - subgraph returns exclusion,
  - full TypeRef retention in Markdown detail sections,
  - unresolved TypeRef omission behavior,
  - Mermaid label readability / escaping boundary.
- Decide whether V01-ADR-074 should be:
  - accepted as-is,
  - revised before acceptance,
  - split into a smaller accepted minimum plus follow-up decision,
  - or deferred.
- Identify the minimum follow-up task set needed after this review, without implementing spec, renderer, YAML, fixture, golden, or test changes.

## Decision Questions

1. Is V01-ADR-074's DAG TypeRef hint contract small enough to accept for the V01-WORK-DATA-007 minimum?
2. Is the boundary between top-level DAG label hint and full TypeRef detail section clear enough?
3. Should params boundary assets be included in the same minimum as ordinary asset nodes?
4. Is shortened QID fallback needed in the initial implementation, or should it be revised / split as a later collision-handling follow-up?
5. Does V01-ADR-074 avoid pulling in tagged union support, MCP semantic identity, UC-002 duplicate task QID repair, or broad notes-retreat cleanup?
6. Does any part of V01-ADR-074 need to be revised before DAG view spec alignment starts?

## Included Scope

- V01-ADR-074 acceptance / revision / split review.
- DAG asset TypeRef hint minimum boundary classification.
- Requirement and prior evidence alignment review.
- Follow-up recommendation for V01-WORK-DATA-007 task flow.

## Excluded Scope

- Editing V01-ADR-074.
- Accepting V01-ADR-074 by changing its status.
- Updating specs.
- Implementing parser / resolver / renderer behavior.
- Updating UC-002 YAML.
- Regenerating fixtures or golden output.
- Implementing V01-ADR-073 tagged union / discriminator payload support.
- Implementing V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP semantic identity / state machine identity.
- Repairing UC-002 duplicate task QID / unresolved flow task issues.
- Performing broad remaining UC-002 notes retreat cleanup.
- Reopening M15, V01-WORK-DATA-001, V01-WORK-DATA-002, V01-WORK-DATA-003, or V01-WORK-DATA-004.

## Done Condition

- V01-ADR-074 is classified as accept-as-is, revise-before-acceptance, split-before-acceptance, or defer.
- Any blocking ambiguity in the DAG TypeRef hint minimum contract is listed.
- The boundary between Mermaid label hint and Markdown detail TypeRef is confirmed or flagged.
- The scope of params boundary asset support is confirmed or flagged.
- Follow-up task input for DAG view spec alignment is concrete enough to create the next task.
- No spec, implementation, renderer, YAML, fixture, golden, or test changes are performed by this task.

## Verification

- Confirm this task only records review and decision input.
- Confirm V01-REQ-DATA-005 excluded scope remains excluded.
- Confirm V01-WORK-DATA-007 remains the owning work item.
- Confirm V01-ADR-074 status is not changed by this task.
- Confirm no UC-002 YAML or generated render artifacts are changed.

## Evidence

Reviewed on 2026-06-02.

### Sources reviewed

- `V01-WORK-DATA-007`: confirms this work item owns DAG asset TypeRef hint render support as a focused DATA render / view successor to V01-ADR-074. It also excludes tagged union support, MCP semantic identity, UC-002 duplicate task QID repair, remaining notes-retreat cleanup, and reopening earlier M15 / DATA work.
- `V01-REQ-DATA-005`: confirms the requirement is to expose useful top-level TypeRef information in DAG asset nodes without changing the underlying TypeRef model, and that render implementation is owned by V01-WORK-DATA-007 and its tasks.
- `V01-ADR-074`: reviewed proposed asset label hint contract, top-level hint rules, named model local-id behavior, shortened QID fallback, params boundary asset inclusion, subgraph returns exclusion, full TypeRef detail retention, unresolved TypeRef omission, and render / fixture impact.
- `V01-TASK-DATA-005-02`: confirms V01-ADR-074 was intentionally classified as DATA render / view support and should not be mixed with tagged union implementation or UC-002 cleanup.
- `docs/spec/views/dag.md`: confirms current active DAG spec still renders asset labels as `asset_name([asset_name])` and params boundary assets as bare asset nodes, with no active TypeRef hint rule yet.
- `docs/spec/type-ref.md`: confirms TypeRef currently supports primitive, named model, inline `list<T>`, inline `dict<T>`, named list / dict model normalization, and existing unresolved / invalid TypeRef diagnostics.
- `docs/spec/diagnostics.md`: confirms current TypeRef-related lint / warning coverage includes `opaque_type_ref`, but DAG render hint behavior does not require a new diagnostic category.

### Review result

Verdict: `revise-before-acceptance`.

V01-ADR-074's core proposal is coherent and should continue as the basis for V01-WORK-DATA-007. The following parts are small enough and useful enough for the V01-WORK-DATA-007 minimum:

- DAG asset labels may show a top-level TypeRef hint as `{asset_name}: {type_hint}`.
- Primitive TypeRef hints use the primitive name directly.
- Named model TypeRef hints use the model local id for Mermaid readability.
- Inline `list<T>` and `dict<T>` hints collapse to `list` and `dict` rather than expanding nested TypeRef details.
- Named list / dict models remain named model hints rather than being displayed as bare `list` / `dict`.
- Main task params boundary assets should be included because they are asset nodes in the DAG render and benefit from the same hint rule.
- `subgraph returns` remains out of scope because V01-ADR-064 already removed returns boundary assets from DAG render.
- Full TypeRef must remain in Markdown detail sections rather than being expanded in Mermaid labels.
- Unresolved / invalid TypeRef should rely on existing diagnostics; Mermaid label generation should omit the hint rather than inventing a second diagnostic surface.

However, V01-ADR-074 should not be accepted as-is.

The main reason is the shortened QID fallback rule. The rule is conceptually reasonable, but it is heavier than the minimum render readability improvement because it requires render-scope collision detection and suffix-qualified model identity calculation. There is no current evidence that V01-WORK-DATA-007's first implementation needs cross-module same-local-id collision handling to validate the core behavior. Keeping it in the accepted minimum risks turning a focused render-label improvement into a broader name-disambiguation implementation.

The second reason is wording / boundary drift. V01-ADR-074 was originally captured as an M15 / v1.1 follow-up. Acceptance should now name V01-REQ-DATA-005 / V01-WORK-DATA-007 as the active boundary and make clear that M15 / V01-WORK-DATA-001 remain closed.

The third reason is implementation staging. Current active DAG spec does not yet define TypeRef hint labels, and current TypeRef spec already owns TypeRef syntax / resolution / diagnostics. The next task should update DAG view spec and accepted ADR wording without changing the underlying TypeRef model or adding new diagnostics.

### Decision question answers

1. Is V01-ADR-074's DAG TypeRef hint contract small enough to accept for the V01-WORK-DATA-007 minimum?
   - Yes for top-level label hints on existing asset nodes. No for requiring shortened QID fallback in the first minimum unless implementation inspection proves it is already trivial.

2. Is the boundary between top-level DAG label hint and full TypeRef detail section clear enough?
   - Mostly yes. The intended boundary is sound: Mermaid gets only the top-level hint, while Markdown detail sections keep full TypeRef. The next ADR/spec revision should state this as the active V01-WORK-DATA-007 rule.

3. Should params boundary assets be included in the same minimum as ordinary asset nodes?
   - Yes. Params boundary entries are rendered as asset nodes and should use the same hint rule. This is not a separate feature.

4. Is shortened QID fallback needed in the initial implementation, or should it be revised / split as a later collision-handling follow-up?
   - It should be revised out of the initial minimum or explicitly staged as a follow-up. It is a valid future behavior, but not required to prove the core DAG TypeRef hint support.

5. Does V01-ADR-074 avoid pulling in tagged union support, MCP semantic identity, UC-002 duplicate task QID repair, or broad notes-retreat cleanup?
   - Yes. V01-ADR-074 is render / view behavior for existing TypeRef information and does not require tagged union, MCP identity, duplicate task QID repair, or broad UC-002 cleanup. V01-REQ-DATA-005 and V01-WORK-DATA-007 should remain the controlling boundary.

6. Does any part of V01-ADR-074 need to be revised before DAG view spec alignment starts?
   - Yes. Revise the acceptance boundary from the old M15 follow-up framing to V01-REQ-DATA-005 / V01-WORK-DATA-007, and either defer shortened QID fallback or mark it as non-minimum collision handling.

### Recommended follow-up

Create or perform the next task as `V01-TASK-DATA-007-02: V01-ADR-074 revision and DAG view spec alignment`.

Recommended work for the next task:

- Revise V01-ADR-074 from `proposed` toward `accepted` after updating the active boundary to V01-REQ-DATA-005 / V01-WORK-DATA-007.
- Preserve the core top-level hint rules for primitive, named model, inline list, inline dict, and named list / dict model.
- Preserve params boundary asset inclusion and `subgraph returns` exclusion.
- Explicitly keep full TypeRef in Markdown detail sections and omit full container expansion from Mermaid labels.
- Revise shortened QID fallback as either:
  - deferred collision-handling follow-up, or
  - non-minimum behavior that may be implemented only if already supported by existing renderer / symbol data.
- Update `docs/spec/views/dag.md` to define the accepted minimum label format and hint calculation rule.
- Do not change TypeRef syntax, TypeRef compatibility, diagnostics, UC-002 YAML, renderer implementation, fixtures, or golden output in the next task unless it is explicitly split beyond spec alignment.

### Verification result

- This task performed review and decision input only.
- V01-REQ-DATA-005 excluded scope remains excluded.
- V01-WORK-DATA-007 remains the owning work item.
- V01-ADR-074 status was not changed by this task.
- No spec, implementation, renderer, UC-002 YAML, fixture, golden, or test changes were performed by this task.
- Opus review is optional but useful before accepting V01-ADR-074 because the only material judgment is whether shortened QID fallback belongs in the accepted minimum or should be split.

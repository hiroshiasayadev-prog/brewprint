# PRODUCT-ADR-SPEC-006: Separate decision workflow checkpoints from canonical design state

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**:
  - PRODUCT-ADR-SPEC-004
- **supersedes**: []
- **migrated_to_spec**: null

## Context

The current Task authoring and artifact-selection guidance implies two absolute rules:

- every recorded design decision belongs immediately in an ADR;
- a Task must not retain decision rationale.

The accepted decision workflow requires a different sequence.
A decision Task must checkpoint each answer immediately so interrupted work can resume safely.
ADR routing occurs only after all relevant decisions are explicit.
Some decisions need durable rationale in an ADR.
Other decisions are direct Specification projections or bounded editorial details.

Without a clear boundary, Task ledgers may be prohibited from storing required workflow state.
The opposite failure is also possible: a Task ledger may be treated as the canonical design source.

## Decision

A Task does not own canonical design state.

A `decision` Task may own temporary and historical decision-workflow state:

- unresolved-decision inventory;
- current cursor;
- questions and options;
- explicit user answers;
- selected option;
- rejected-option identifiers;
- a concise one-to-three-sentence reason;
- canonical target;
- ADR-routing state;
- dependency, blocked, or deferred state.

A `decision` Task ends when the selected outcome, concise reason, and canonical target are fixed in its ledger.
The Task does not create or update canonical ADR or Specification body content.

An `authoring` Task begins when decided inputs are written into canonical artifacts.
Authoring must stop when the inputs permit several materially different interpretations.
The unresolved choice returns to a `decision` Task.

### Conditional ADR routing

Every explicit decision is first checkpointed in the decision ledger.
A later ADR-routing step classifies the decision.

| routing result | action |
|---|---|
| Durable rationale required | Create, amend, or supersede an ADR in a separate authoring Task. |
| Existing accepted ADR covers the decision | Record reuse in routing Evidence and synchronize the Specification. |
| ADR not required | Synchronize the accepted result directly into the relevant Specification or scope record. |
| Routing blocked | Stop authoring until the named authority or decision exists. |

ADR routing also owns:

- coherent ADR-boundary partitioning across several decision items;
- prevention of oversized omnibus ADRs;
- prevention of mechanical one-row-per-ADR fragmentation;
- create, amend, reuse, or supersede disposition;
- exact decision-ID and Specification-target mapping.

ADR routing does not author ADR body content.

### Amendment and supersession materiality

Use an in-place ADR amendment when:

- the selected alternative remains unchanged;
- the core architecture and rationale remain valid;
- the change clarifies or extracts a responsibility inside that architecture;
- the amendment does not conceal a reversal.

Responsibility extraction is not automatically a material ownership change.
Routing must record the materiality judgment.

Use supersession when:

- the selected alternative changes;
- the core ownership architecture changes materially;
- an accepted constraint is removed or reversed;
- the previous rationale no longer justifies the current state.

An ADR is not required merely because a decision occurred.
An ADR is required when alternatives, rationale, consequences, ownership, or supersession history must remain understandable.

### Canonical ownership

| information | canonical owner |
|---|---|
| Decision workflow state and resumable checkpoint history | `decision` Task |
| Durable choice, alternatives, rationale, consequences, and supersession | ADR |
| Current normative behavior, structure, boundary, and constraint | Specification |
| Complete requirement-resolution graph | Work Item |

After ADR authoring, the completed decision Task retains its workflow history unchanged.
ADR authoring, Specification synchronization, review, and closure Tasks own their own references and Evidence.
Downstream progress must not be written back into the completed decision Task.
The ADR becomes the canonical durable rationale.
The Specification remains the canonical current-state contract.

A Task ledger must not be used as a substitute for either artifact.

## Rationale

Workflow persistence and canonical design ownership solve different problems.

Immediate Task checkpoints prevent lost answers, repeated questions, and unsafe session resumption.
Conditional ADR routing prevents a new ADR for every editorial or mechanically derived decision.

Separating decision and authoring phases also preserves writer boundaries.
The decision phase fixes the input.
The authoring phase records that input without reopening it.

Keeping durable rationale in ADRs prevents Specifications from accumulating history.
Keeping current contracts in Specifications prevents ADRs from becoming stale operational manuals.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Create an ADR after every answered question. | Many decisions are direct Specification projections with no durable trade-off. |
| Prohibit all rationale in Tasks. | Decision workflows need concise resumable Evidence before ADR routing. |
| Treat the Task ledger as canonical design authority. | Workflow state is not the current normative contract or durable rationale source. |
| Let decision Tasks author ADR body content. | Decision confirmation and canonical authoring have different completion judgments. |
| Put full decision history in Specifications. | Specifications must describe current truth rather than decision history. |
| Delete Task decision history after ADR authoring. | Historical workflow Evidence remains useful for traceability and review. |

## Consequences

- `spec:product.design_records.authoring_standards.task_authoring` must permit concise decision-workflow rationale.
- Task authoring must replace the unconditional every-decision-to-ADR wording with conditional routing.
- `spec:product.design_records.authoring_standards.artifact_boundary` must distinguish workflow checkpointing from durable ADR recording.
- `spec:product.design_records.artifact_model.artifact_responsibility_matrix` must distinguish Task workflow state from canonical design state.
- Decision Tasks and authoring Tasks remain separate phases even when one session executes both.
- ADR routing and ADR authoring remain separate responsibilities.
- ADR routing records whether a responsibility refinement preserves or materially changes the accepted ownership architecture.
- `spec:product.design_records.authoring_standards.adr_authoring` and the active ADR-routing workflow must use the same amendment boundary.
- Completed decision Tasks do not track downstream ADR, Specification, review, or closure progress.
- Independent review remains a later Task and is not performed by the authoring session.
- Existing Task migration remains outside W016.

## Evidence

- `PRODUCT-REQ-SPEC-005`: accepted typed single-responsibility Task requirement.
- `PRODUCT-ADR-SPEC-004`: accepted Task-type taxonomy.
- `PRODUCT-TASK-SPEC-016-02`: D-012.
- `PRODUCT-TASK-SPEC-016-04`: C-006, C-007, C-013, and C-014.
- `PRODUCT-TASK-SPEC-016-05`: ADR routing and this ADR boundary.
- `skills/design-decision-workflow/adr-routing.md`: conditional ADR-routing authority.
- `PRODUCT-TASK-SPEC-018-01`: D-019 and D-021.
- `PRODUCT-TASK-SPEC-018-02`: B-007 clarification-amendment route.
- `PRODUCT-TASK-SPEC-018-03`: meaning-preserving clarification amendment.
- `PRODUCT-TASK-SPEC-018-07`: F-BLK-01.
- `PRODUCT-TASK-SPEC-018-16`: accepted ADR amendment materiality decision.

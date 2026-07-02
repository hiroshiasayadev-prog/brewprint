# PRODUCT-TASK-SPEC-024-01: Decide Work Item framing workflow

- **id**: PRODUCT-TASK-SPEC-024-01
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-024
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - PRODUCT-TASK-SPEC-024-01
  - PRODUCT-TASK-SPEC-024-02
  - PRODUCT-WORK-SPEC-024

## Goal

Fix one bounded MVP contract for Work Item framing and its exact downstream authoring route.

## Work

- Align the framing purpose and Requirement entry boundary.
- Decide source-disposition and Work Item contract ownership.
- Decide conditional Investigation, review, synchronization, and Task materialization rules.
- Decide the design-convergence handoff.
- Decide the existing Work Item exclusion.
- Record the accepted decisions and their canonical targets.
- Directly materialize the uniquely required T02 authoring Task.
- Update the parent Work Item Task flow and Task list.

This Task does not author the framing skill or canonical Specification content.

The direct T02 materialization is the accepted framing bootstrap exception.
No separate graph judgment remained.

### Decision ledger

| ID | topic | status | decision summary | reason | canonical target | ADR route |
|---|---|---|---|---|---|---|
| D-001 | Framing purpose | `decided` | Framing aligns Requirement intent, decides source handling, and fixes the downstream Work Item contract before design or execution planning. | Design convergence cannot safely start while user and agent understanding of the work differs. | Framing skill | `not_required` |
| D-002 | Workflow separation | `decided` | Work Item framing is a separate upstream skill. Design convergence is one conditional downstream route. | Goal negotiation and design convergence own different completion judgments. | Framing and design-convergence skills | `not_required` |
| D-003 | Existing Requirement entry | `decided` | Treat the existing Problem as accepted input. Confirm Desired Outcome against Required Outcome. | Re-eliciting an accepted Problem adds noise without resolving the actual alignment risk. | Framing skill | `not_required` |
| D-004 | Requirement mismatch | `decided` | Route mismatches through amendment, split, follow-up, replacement, rejection, deferral, Investigation, or blocking as appropriate. | A mismatch must not be hidden inside Work Item planning. | Framing routing | `not_required` |
| D-005 | No-Requirement entry | `decided` | Identify Problem and Desired Outcome sufficiently to capture a Requirement before repository-persistent framing continues. | Work Items require direct material provenance and Requirement owns the stable need. | Framing skill and Requirement authoring | `not_required` |
| D-006 | Framing conclusion | `decided` | Framing succeeds when it reaches an explicit disposition. Proceeding is not required. | Reject, defer, investigate, block, or proceed are all valid conclusions. | Framing routing | `not_required` |
| D-007 | Proceed contract | `decided` | Proceeding fixes Goal, Boundary, Completion Condition, unknown handling, and initial downstream route. | These fields define the Work Item's execution identity without requiring detailed downstream design. | Framing loop and Work Item authoring | `not_required` |
| D-008 | Decision ownership | `decided` | One `decision` Task owns the framing inventory, cursor, user answers, disposition, conditional Work Item contract, unknown handling, and selected route. | Framing outcomes are judgments, not a new Task responsibility type. | Task authoring and framing loop | `not_required` |
| D-009 | Task type set | `decided` | Do not add a `framing` Task type. | Existing Task types already own decision, Investigation, authoring, decomposition, coordination, and synchronization. | Task authoring and framing skill | `not_required` |
| D-010 | Initial graph | `decided` | Start with exactly one framing `decision` Task. | Later responsibilities are unknown until the framing decision chooses a route. | Framing skill | `not_required` |
| D-011 | Conditional Tasks | `decided` | Create T02 and later Tasks only when the active framing decision determines they are required. | Speculative Tasks create unused workflow overhead. | Framing skill | `not_required` |
| D-012 | Formal Investigation | `decided` | Formal Investigation is conditional. Limited downstream research may be recorded without an Investigation when the route remains coherent. | Not every framing uncertainty needs a separate research artifact. | Framing routing and design convergence | `not_required` |
| D-013 | Independent review | `decided` | Framing has no mandatory independent review. | Direct user judgment is the framing acceptance boundary. | Framing skill | `not_required` |
| D-014 | Direct Task materialization | `decided` | The active framing decision may create and register same-Work-Item Tasks when the selected route uniquely fixes their type, outcome, and dependency. | A coordination Task solely to create a uniquely determined Task is self-referential overhead. | Task authoring and framing skill | `not_required` |
| D-015 | Coordination boundary | `decided` | Use coordination only for graph repair, alternative graph design, shared-writer sequencing, cross-Work-Item change, or release-order judgment. | These cases own an independent graph judgment. | Task authoring and graph coordination | `not_required` |
| D-016 | Work Item creation | `decided` | Create a required downstream Work Item through `work_item_decomposition`. | Work Item creation owns an independent artifact and completion judgment. | Framing routing and decomposition | `not_required` |
| D-017 | Simple closure | `decided` | The framing workflow may directly close its Work Item after required outcomes and materialized Tasks complete. | A synchronization Task is unnecessary when no separate propagation judgment exists. | Framing skill | `not_required` |
| D-018 | Design-convergence entry | `decided` | Design convergence starts only when framing selects it. Impact Investigation is conditional there. | Design convergence must consume an accepted Work Item boundary rather than form it implicitly. | Design-convergence skill | `not_required` |
| D-019 | Existing Work Items | `decided` | Do not migrate, repair, or retroactively legitimize existing unframed Work Items. | The MVP establishes the future framing flow only. | Work Item boundary | `not_required` |
| D-020 | MVP skill shape | `decided` | Use `SKILL.md`, `interactive-framing-loop.md`, and `framing-routing.md`. | Three files separate activation, interaction, and disposition routing without premature fragmentation. | `skills/work-item-framing/` | `not_required` |
| D-021 | Downstream route | `decided` | One authoring Task must create and activate the skill and amend directly conflicting authority. | The accepted decisions fix one writer, one artifact set, and one completion judgment. | PRODUCT-TASK-SPEC-024-02 | `not_required` |

### Current cursor

- Decision: none
- Loop state: `decision_complete`
- Every owned decision is terminal.

## Done condition

- D-001 through D-021 are `decided`, `deferred`, or validly `blocked`.
- The framing purpose, entry rules, disposition boundary, and proceed contract are fixed.
- Conditional Investigation, review, synchronization, and Task materialization rules are fixed.
- The design-convergence handoff and existing Work Item exclusion are fixed.
- The exact T02 authoring route is materialized.
- No framing skill or canonical Specification content is authored.

## Verification

- Confirm every user judgment is represented once in the ledger.
- Confirm no unresolved decision remains.
- Confirm the initial graph was not populated speculatively.
- Confirm T02 has one authoring outcome and depends only on T01.
- Confirm no formal Investigation, independent review, coordination, or synchronization Task was created.
- Confirm no existing Work Item migration or repair entered the Work Item boundary.

## Evidence

- The user accepted a separate Work Item framing skill.
- The user accepted Requirement-first repository persistence and Required Outcome alignment.
- The user accepted explicit reject, defer, Investigation, limited research, blocked, and proceed conclusions.
- The user accepted one decision Task as the initial graph.
- The user accepted conditional T02 and later Task materialization.
- The user accepted no mandatory independent review or synchronization Task.
- The user accepted same-Work-Item direct Task materialization by the active decision when the route is uniquely fixed.
- The user excluded existing Work Item migration and repair.
- T02 is the only required downstream Task for the accepted MVP route.
- The standalone Task responsibility validator evaluated all common and `decision` criteria after final Evidence. Every criterion returned `true`.
- DRMCP is non-operational. Filesystem authoring is the required fallback.

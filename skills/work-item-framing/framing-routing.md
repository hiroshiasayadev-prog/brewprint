# Framing routing

## Purpose

Select one explicit source disposition and the minimum required downstream route.

Routing does not perform formal Investigation, canonical authoring, downstream Work Item creation, design convergence, or implementation.

## Source dispositions

| disposition | meaning | required route |
|---|---|---|
| `proceed` | The Requirement is suitable for concrete downstream work. | Fix the downstream Work Item contract and select its initial route. |
| `amend` | The same Requirement identity remains valid but its expression or boundary must change. | Materialize bounded Requirement `authoring`, then resume framing when needed. |
| `split` | An independently acceptable outcome must receive separate Requirement identity. | Materialize Requirement `authoring`; frame each selected Requirement separately. |
| `follow_up` | The original Requirement remains completable and an adjacent need appears. | Create a follow-up Requirement through bounded `authoring`. |
| `replace` | The motivating problem or Required Outcome no longer represents the intended request. | Create replacement Requirement authority and stop using the original as the downstream source. |
| `reject` | The Requirement is not adopted for downstream work. | Persist the reason and close framing without a downstream Work Item. |
| `investigate` | One formal research question must be answered before disposition can complete. | Materialize an `investigation` Task and block the active decision item. |
| `defer` | The Requirement is recognized but not handled now. | Persist the reason and restart condition; close framing when no other action is required. |
| `blocked` | A named external input or authority prevents a valid disposition. | Persist the blocker and keep the framing Work Item blocked. |

A framing disposition describes source handling.
Limited research is not a separate disposition when the accepted route can proceed coherently.

## Requirement identity routing

| observed condition | route |
|---|---|
| Problem and Required Outcome remain the same; wording or same-outcome boundary is inaccurate. | `amend` |
| Problem or Required Outcome changes materially and can be accepted independently. | `split` or `replace` |
| Original Requirement remains completable and an adjacent need appears. | `follow_up` |
| Requirement remains valid but the proposed downstream work departed from it. | Preserve the Requirement and revise the framing decision. |
| Adoption is not justified. | `reject` |
| Evidence is insufficient for identity or adoption judgment. | `investigate`, `defer`, or `blocked` |

Do not hide Requirement identity changes inside Work Item Goal or Boundary text.

## Unknown routing

| unknown type | route |
|---|---|
| Direct accepted authority resolves the fact. | Read it during framing. |
| A small scoped repository read resolves the fact. | Read it during framing without a formal Investigation. |
| Limited research can occur downstream without changing Goal or completion meaning. | Record the research need in the downstream Work Item contract. |
| One bounded research question needs durable Evidence, uncertainty, and options. | `investigate` |
| The topic has insufficient current value. | `defer` |
| A named external input is unavailable. | `blocked` |

The amount of reading alone does not decide whether formal Investigation is required.
Use the ownership and durability of the research outcome.

## Proceed route

`proceed` requires these accepted inputs:

| input | required meaning |
|---|---|
| Goal | One downstream outcome. |
| Boundary | Owned and excluded responsibility. |
| Completion Condition | One observable completion meaning. |
| Direct source | The Requirement that materially motivates the Work Item. |
| Unknown handling | Formal Investigation, limited downstream research, resolved fact, deferment, or blocker. |
| Initial route | The first downstream responsibility. |

Possible initial routes include:

- design convergence;
- formal Investigation;
- Work Item decomposition;
- mechanical authoring from accepted authority;
- implementation planning from an already accepted contract;
- another explicitly owned workflow.

Do not select design convergence automatically.

## Task route selection

| required outcome | owner |
|---|---|
| Framing judgment | active `decision` Task |
| Formal research record | `investigation` Task |
| Requirement or accepted artifact writing | `authoring` Task |
| Downstream Work Item creation or split | `work_item_decomposition` Task |
| Independent Task-graph judgment or repair | `coordination` Task |
| Separate mechanical propagation | `synchronization` Task |

The active framing decision directly materializes a same-Work-Item Task only under the bounded exception in `SKILL.md`.

## Design-convergence route

Select design convergence when the downstream Work Item needs repository-persistent design decisions, ADR routing, normative Specification authoring, and integrated independent review.

The handoff must state whether formal impact Investigation is required.
Design convergence may later add Investigation when a new bounded research need appears.

## Closure routes

| disposition state | framing Work Item outcome |
|---|---|
| `reject`, `defer`, or another complete non-proceed disposition | `done` after the reason and route are persisted. |
| `proceed` with no conditional Task | `done` after the downstream contract and handoff are persisted. |
| `proceed` with conditional Tasks | `done` after every required framing Task completes and the handoff is persisted. |
| unresolved named external dependency | `blocked` |

A rejected Requirement can represent successful framing.
Do not mark the framing Work Item unsuccessful merely because no downstream Work Item is created.

## Stop conditions

Stop routing when:

- Desired Outcome remains unclear;
- Required Outcome has materially different valid interpretations;
- Requirement identity cannot be selected safely;
- the proceed contract permits several independent completion meanings;
- the downstream owner cannot be identified;
- direct Task materialization would require a separate graph judgment.

Return unresolved judgment to the active framing decision Task.

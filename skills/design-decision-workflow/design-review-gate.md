# Design review gate

This document defines the independent review, finding correction, finding-closure re-review, and lifecycle closure sequence for the design decision workflow.

It is a required companion to `SKILL.md` when creating or evaluating the review side of the workflow.

## Purpose

Verify that interactive decisions were converted into coherent canonical design without hidden assumptions, missing ADRs, stale Specifications, or premature lifecycle closure.

The gate reviews the complete trace:

```text
Work Item scope and completion conditions
  -> decision inventory and explicit decisions
  -> ADR routing and ADRs when required
  -> current normative Specifications
  -> deferred scope and downstream obligations
```

The review is not an implementation review and must not expand into production-code assessment unless production artifacts are explicitly part of the design contract under review.

## Preconditions

Begin independent design review only when:

- the decision loop has reached `decision_complete` or an explicitly reviewable blocked state;
- every `decided` item has a resolved ADR route;
- required ADRs exist at the lifecycle stage expected by the Work Item;
- Specification synchronization is complete for the review boundary;
- exact changed or authored artifacts are known;
- the authoring session has stopped making design changes;
- the reviewer is independent from the authoring judgment.

If these conditions are not met, return `NOT READY` rather than reviewing an incomplete design as though it were final.

## Review boundary

The review Task must name exact artifacts:

- parent Work Item;
- decision-loop Task;
- ADR authoring Task or routing Evidence;
- each new, amended, reused, or superseded ADR in scope;
- each changed Specification file or section;
- directly governing Requirements and accepted authority needed for judgment;
- explicit deferred or blocked decisions;
- repository authoring and lifecycle standards required for the artifacts.

Do not perform repository-wide traversal to discover already-known review inputs.

Read adjacent artifacts only when a scoped relation or contradiction requires them.

## Reviewer independence

The review Task is read-only.

The reviewer must not:

- change decision status;
- edit ADRs or Specifications;
- correct findings;
- synchronize lifecycle state;
- start implementation;
- accept the author's summary as proof without reading the named artifacts.

The reviewer may identify a newly revealed design gap when it is directly exposed by the scoped trace.

## Review criteria

### Work Item and workflow integrity

Verify:

- the Work Item topic is bounded;
- the Task graph follows the required sequence;
- conditional ADR, correction, and re-review routes are represented correctly;
- writer ownership is unambiguous;
- completion conditions match the actual design deliverables;
- implementation work is not hidden inside the design workflow.

### Decision inventory integrity

Verify:

- all material decisions revealed by the scoped topic are represented;
- decisions answered by accepted authority were not unnecessarily delegated to the user;
- each explicit user decision is represented accurately;
- dependencies and blockers are coherent;
- no item is incorrectly marked `recorded` before canonical synchronization;
- deferred scope is explicit and does not leave the current design internally incomplete;
- no required decision exists only in chat or narrative Evidence.

### ADR routing integrity

Verify:

- every durable non-obvious decision has an ADR when required;
- `not_required` and `covered` outcomes are justified;
- existing ADR reuse is scope-compatible and not superseded;
- ADR granularity is coherent;
- amendments do not conceal material reversals;
- supersession is explicit and historically honest;
- rationale and consequences are sufficient without duplicating the Specification.

### Specification integrity

Verify:

- every normative accepted decision appears in the correct Specification;
- Specification language states current behavior rather than discussion history;
- ADR and Specification statements agree;
- obsolete or contradictory text is removed or dispositioned;
- Requirements remain satisfied;
- scope exclusions are explicit;
- references and relations follow repository policy;
- no unresolved design choice is delegated implicitly to an implementation executor.

### Cross-artifact traceability

For each decision ID entering review, verify a valid review-ready trace:

```text
decided decision
  -> ADR required / covered / not_required disposition
  -> exact canonical Specification target
  -> review evidence produced by this gate
```

The closure-synchronization Task changes `decided` to `recorded` only after this trace passes independent review.

A `deferred` or `blocked` item must instead have:

```text
status
  -> exact reason
  -> dependency or destination
  -> proof that current design remains coherent without it
```

### Review readiness for implementation

Verify that downstream implementation Task authoring can proceed without discovering:

- product behavior;
- architecture ownership;
- compatibility policy;
- persistence authority;
- failure semantics;
- data shape;
- validation boundary;
- unresolved cross-component responsibility.

The review does not need to define implementation files, symbols, or commands unless the design Work Item explicitly includes executor Task authoring.

## Verdicts

Use one of:

- `PASS`: no blocking, major, or required minor finding prevents design closure;
- `NEEDS REVISION`: at least one material finding requires correction;
- `NOT READY`: authoring prerequisites are incomplete, so independent review cannot be validly performed;
- `BLOCKED`: required authority or evidence is unavailable and cannot be recovered within the scoped boundary.

Do not use `PASS WITH CHANGES` when the reviewer has not performed or cannot independently verify those changes.

## Finding severity

### Blocking

Use blocking when:

- the selected design conflicts with an accepted Requirement or controlling authority;
- a required decision remains unresolved but the workflow claims closure;
- canonical ownership cannot be determined;
- the design would permit materially incompatible implementations;
- review independence or evidence is insufficient to establish a valid verdict;
- a required ADR supersession is missing and current authority is contradictory.

### Major

Use major when:

- a durable decision lacks a required ADR;
- ADR and Specification materially disagree;
- a normative decision is absent from Specification;
- decision status or traceability materially misrepresents repository state;
- stale canonical text would direct implementation incorrectly;
- deferred scope leaves the current design incomplete;
- the Work Item completion contract cannot be satisfied as written.

### Minor

Use minor when:

- traceability, references, consequences, or wording are incomplete but the design remains materially unambiguous;
- a non-critical status, Evidence, or scope detail needs correction;
- an ADR or Specification has a bounded clarity defect that does not change the selected design.

### Advisory

Use advisory for optional improvements that are not required for closure.

Do not inflate editorial preferences into findings.

## Review output

Default output:

1. Verdict
2. Reviewed artifacts
3. Decision-to-ADR-to-Specification trace result
4. Blocking findings
5. Major findings
6. Minor findings
7. Advisories
8. Implementation-planning readiness
9. Exact next gate

Each material finding must include:

- finding ID;
- severity;
- affected decision IDs;
- affected artifact and section;
- observed contradiction or omission;
- required correction outcome;
- whether a new user decision is required.

The reviewer should not prescribe implementation details beyond what is needed to make the design unambiguous.

## Finding correction

A correction Task must:

- quote or reference exact finding IDs;
- name exact writable ADR, Specification, Task, or Work Item files;
- preserve accepted decisions not implicated by the findings;
- identify whether the correction is editorial, traceability-only, canonical synchronization, ADR routing, or renewed decision confirmation;
- stop when a finding requires a new user decision;
- avoid unrelated cleanup;
- avoid declaring findings closed;
- avoid lifecycle closure.

When correction requires a new design decision:

1. return the affected decision ID or append a new one;
2. resume the interactive decision loop for that one decision;
3. persist the answer;
4. repeat ADR routing and Specification synchronization as required;
5. then request closure re-review.

Do not let the correction author invent the missing decision.

## Finding-closure re-review

The re-review is read-only and independent.

It must verify:

- each specified finding is closed;
- direct cross-artifact consistency effects are correct;
- no accepted decision was unintentionally changed;
- no direct regression was introduced;
- required new decision evidence exists when applicable;
- canonical ADR and Specification state now agree.

Do not reopen the full design unless the correction directly reveals a blocking contradiction.

A new finding may be reported only when it is caused by, or directly exposed by, the correction under review.

## Closure synchronization gate

Lifecycle and Evidence synchronization may begin only after:

- initial review `PASS`; or
- correction is complete and closure re-review confirms all required findings closed.

The closure Task must:

- update `decided` items to `recorded` only when exact canonical refs exist;
- record ADR routing outcomes;
- record exact ADR and Specification refs;
- record review verdict and finding dispositions;
- preserve deferred and blocked items with reasons and destinations;
- confirm Work Item completion conditions;
- update only explicitly owned lifecycle records;
- avoid production implementation and executor prompt issuance unless separately authorized.

A Work Item must not be closed when a required item remains `open`, `in_discussion`, or `decided`.

A `blocked` item may remain only when the Work Item completion contract explicitly permits that blocked scope and the current design is still coherent.

## Relationship to implementation Task authoring

After design closure:

- treat accepted ADRs and Specifications as implementation authority;
- create or revise executor-ready implementation Tasks separately;
- evaluate the execution-hub pattern before issuing implementation prompts;
- do not copy interactive discussion into executor prompts;
- do not ask implementation agents to rediscover decisions closed by this workflow.

If implementation Task authoring exposes a missing design choice, return to a new or reopened design-decision workflow rather than deciding it implicitly inside implementation.

## Stop conditions

Stop and report the exact blocker when:

- a named review artifact is missing;
- the decision ledger and canonical artifacts disagree so fundamentally that the intended decision is unclear;
- a finding requires a new user decision;
- accepted authority conflicts and precedence cannot be resolved from repository rules;
- the review boundary omits an artifact necessary for a valid verdict;
- closure synchronization would claim evidence that does not exist.

Do not broaden into repository-wide review to compensate for an incomplete review contract.

## Anti-patterns

Do not:

- let the design author perform the independent review;
- combine correction and closure re-review;
- accept Task summaries instead of canonical artifacts;
- review only ADR prose while ignoring Specification state;
- review only Specification diffs while ignoring decision and ADR routing;
- close findings based on author claims without direct evidence;
- mark decisions `recorded` before canonical synchronization;
- close the Work Item before required findings are dispositioned;
- start implementation during review or closure;
- reopen unrelated design areas during a scoped re-review;
- treat advisory wording preferences as release blockers.

## Readiness checklist

Before returning `PASS`, verify:

- review prerequisites were satisfied;
- the reviewer inspected exact scoped artifacts;
- every required decision has a valid canonical trace;
- ADR routing is complete and correct;
- Specifications express current normative design;
- ADR and Specification content agree;
- deferred or blocked scope is explicit and safe;
- no hidden design judgment remains for implementation;
- no material finding remains open;
- closure synchronization has an exact next owner.

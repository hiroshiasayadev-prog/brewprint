# ADR routing

## Purpose

Classify every terminal decision and partition required durable choices into coherent ADR boundaries.

ADR routing is separate from ADR authoring.
It decides what ADR work is required without writing ADR body content.

## Inputs

Read:

- parent Work Item;
- completed decision Task and exact decision IDs;
- completed impact Investigation when one was selected;
- reconciliation disposition;
- accepted ADRs that may cover or conflict with the decision;
- affected Specification targets;
- ADR authoring and lifecycle standards.

Do not reopen accepted decisions unless direct authority conflict makes their intended meaning unclear.

## Routing outcomes

Classify each terminal decision as:

| outcome | meaning |
|---|---|
| `required` | A durable choice requires ADR creation, amendment, or supersession. |
| `covered` | An accepted, non-superseded ADR already owns the same durable choice. |
| `not_required` | No durable trade-off needs ADR history; project directly to the correct canonical artifact. |
| `blocked` | ADR disposition requires named authority, Investigation, or decision input. |

Inventory states such as `candidate` and `unknown` are provisional only.
Resolve them before canonical authoring or record the exact blocker.

## ADR-required criteria

An ADR is normally required when:

- several reasonable alternatives existed;
- the choice establishes or changes architecture or ownership boundaries;
- the choice governs canonical state, lifecycle, validation, compatibility, persistence, integration, or failure handling;
- a meaningful trade-off or limitation is accepted;
- several Specifications or future implementation waves are constrained;
- future maintainers need the decisive rationale;
- migration or future evolution is constrained;
- an accepted ADR must change or be superseded;
- a durable exception to a repository rule is introduced.

Affected file count is not a routing criterion.

## ADR-not-required criteria

An ADR is normally not required when:

- the result is mechanically derived from accepted authority;
- wording or references are corrected without changing behavior or ownership;
- a bounded inclusion or exclusion has no durable trade-off;
- the choice is an intentionally delegated local implementation detail;
- only a typo, identifier, path, example, or format is corrected;
- an existing accepted ADR already determines the canonical correction;
- no meaningful alternative or rationale needs preservation.

Do not create an ADR merely because a Specification changes.
Do not omit one merely to reduce artifact count.

## Existing ADR coverage

Classify `covered` only when the existing ADR:

- is accepted and usable as current authority;
- includes the current topic;
- states the same durable choice;
- has consequences compatible with the new decision;
- is not superseded;
- does not omit a new independently durable trade-off.

When the core choice remains valid but wording, references, or consequences are stale, choose an in-place clarification amendment only when repository policy permits it and history remains honest.

## Create, amend, reuse, or supersede

Use `create` when no current ADR owns the durable choice.

Use `reuse` when an accepted ADR fully covers it.

Use `amend` only when:

- the selected alternative remains unchanged;
- the core architecture and rationale remain valid;
- the change clarifies wording, references, consequences, or an overloaded responsibility;
- any responsibility extraction stays inside the accepted architecture;
- the routing record explicitly classifies the change as non-material;
- the amendment does not conceal a reversal.

Responsibility extraction alone does not require supersession.

Use `supersede` when:

- the selected alternative changes;
- the core ownership architecture changes materially;
- an accepted constraint is removed or reversed;
- previous rationale no longer justifies current state;
- keeping the old ADR current would mislead implementation or review.

A superseding ADR preserves the old record as history and updates canonical Specifications separately.

## ADR-boundary partitioning

Routing owns the boundary between ADRs.

Keep decisions in one ADR when:

- they select one inseparable alternative;
- their alternatives, rationale, and consequences are shared;
- reviewing or superseding one independently would create an invalid state.

Split them when:

- they can change independently;
- they have distinct alternatives or rationale;
- they govern different ownership or architecture boundaries;
- one may be superseded while another remains valid;
- combining them would create an oversized omnibus ADR.

Do not default to one ADR per decision row.
Do not combine a full workflow merely because all decisions were made in one loop.

## Routing record

Record the result in the ADR-routing Task Evidence.
For each decision include:

- decision ID;
- routing outcome;
- ADR boundary ID when required;
- exact existing, new, amended, or superseding ADR ID when known;
- disposition `create`, `amend`, `reuse`, or `supersede`;
- reason for `covered` or `not_required`;
- affected Specification and other canonical targets;
- exact blocker when blocked.

Also record for each ADR boundary:

- included decision IDs;
- coherent bounded question;
- why the decisions belong together;
- required authoring Task and writer;
- dependency on other ADR amendments or creations.

Do not write routing results into the completed decision Task.

## ADR authoring handoff

Create one or more separate `authoring` Tasks after routing.

Split ADR authoring Tasks only when artifact ownership, writer, dependency, or completion judgment differs.
Several ADR files may share one Task when they consume the same routed input set, share one writer, and share one acceptance boundary.

The authoring Task receives:

- exact ADR IDs and dispositions;
- exact decision IDs per boundary;
- accepted authority and existing ADR dependencies;
- required alternatives, rationale, and consequence boundaries;
- affected Specification targets;
- explicit prohibited scope.

## Review expectations

Integrated review verifies:

- every decision has a routing outcome;
- durable choices were not omitted;
- unnecessary ADRs do not duplicate accepted authority;
- boundaries are coherent;
- amendments preserve history;
- supersession is explicit;
- ADRs and Specifications agree;
- `not_required` decisions still reach the correct canonical target.

## Stop conditions

Stop and report the exact blocker when:

- a decision permits materially different durable choices;
- accepted Requirements or ADRs conflict with the chosen result;
- existing ADR coverage or status is unclear;
- supersession is required but the replacement choice is not explicit;
- coherent ADR boundaries cannot be selected without new judgment;
- authoring would require changing an out-of-scope Requirement or topic.

## Readiness checklist

Before ADR authoring:

- every terminal decision has one resolved route;
- every `required` item belongs to one coherent boundary;
- every `covered` item names an accepted non-superseded ADR;
- every `not_required` item has a reason and canonical target;
- every `blocked` item names its missing input;
- create, amend, reuse, and supersede dispositions are explicit;
- every `amend` that changes responsibility wording records a materiality judgment;
- authoring Tasks have one responsibility and one completion judgment;
- no ADR body was written during routing.

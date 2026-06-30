# ADR routing

This document defines when a confirmed design decision requires an ADR, how ADR scope is chosen, and how ADR authoring relates to current Specification synchronization.

It is a required companion to `SKILL.md` when classifying, creating, updating, or superseding ADRs.

## Purpose

Use ADRs for durable decisions whose alternatives, rationale, consequences, or supersession history must remain understandable after the interactive conversation ends.

Do not use ADRs as a duplicate Specification or as a container for every small clarification.

## Routing outcomes

Classify every `decided` item as one of:

- `required`: an ADR must be created, amended, or superseded before Specification synchronization is complete;
- `not_required`: the decision may be synchronized directly to Specification or scope records;
- `covered`: an existing accepted ADR already owns the decision and remains valid;
- `blocked`: ADR disposition cannot be completed until a named decision, authority, or investigation exists.

The decision inventory may use `candidate` or `unknown` before the user decision is explicit. Resolve those provisional states before design closure.

## ADR required criteria

An ADR is normally required when one or more of the following is true:

- multiple reasonable alternatives existed and the choice is not obvious from accepted authority;
- the decision establishes or changes an architecture boundary;
- the decision establishes ownership of canonical state, lifecycle, validation, compatibility, persistence, integration, or failure handling;
- the decision introduces a meaningful trade-off or accepted limitation;
- the choice affects several Specifications, components, or future implementation waves;
- future maintainers need to understand why the chosen design is preferred;
- the decision constrains future evolution or migration;
- an accepted ADR must be changed or superseded;
- rejecting a plausible alternative is itself important design context;
- the decision creates a durable exception to a broader repository rule.

A small textual change may still require an ADR when the underlying decision is durable and non-obvious.

## ADR not required criteria

An ADR is normally not required when all relevant conditions are true:

- the result is mechanically derived from an accepted Requirement, ADR, or Specification;
- the change only clarifies wording without changing behavior or ownership;
- the decision is a bounded inclusion or exclusion with no durable trade-off;
- the choice is a local implementation detail intentionally delegated by an accepted design;
- only a typo, path, identifier, example, or formatting error is corrected;
- the current Specification was stale and one accepted ADR already determines the correction;
- no meaningful alternative or rationale needs to be preserved.

Do not skip an ADR merely because only one file or one symbol is affected.

Do not create an ADR merely because a Specification section changes.

## Existing ADR coverage

Classify a decision `covered` only when an accepted ADR already states the same durable choice and remains compatible with the new decision.

Before reusing an ADR, verify:

- its status permits use as current authority;
- its scope includes the current design topic;
- its decision is not narrower or materially different;
- its consequences remain accurate enough for the current change;
- no later ADR supersedes it;
- the new decision does not add a separate durable trade-off.

When the existing ADR covers the choice but its consequences or references are stale, update it only if repository policy permits an in-place clarification that does not rewrite history. Otherwise author a follow-up ADR.

## Amend versus supersede

Use an in-place amendment only when:

- the original decision remains the same;
- the change corrects or clarifies non-material wording, references, or consequences;
- repository policy allows the amendment;
- the amendment does not conceal that a different alternative is now selected.

Supersede the ADR when:

- the selected alternative changes;
- the ownership or architecture boundary changes materially;
- an earlier accepted constraint is removed or reversed;
- the previous rationale no longer justifies the current system;
- retaining the old ADR as current would mislead implementation or review.

When superseding:

- create the new ADR according to repository authoring policy;
- mark the old ADR superseded through the supported lifecycle mechanism;
- preserve traceability in both directions when the format supports it;
- update affected Specifications separately;
- do not rewrite the old ADR to make the new decision appear historical.

## ADR granularity

One ADR should own one coherent decision boundary.

Keep decisions in one ADR when:

- they select one inseparable alternative;
- their rationale and consequences are shared;
- reviewing or superseding one without the others would create an invalid state.

Split decisions into separate ADRs when:

- they can change independently;
- they have distinct alternatives or rationales;
- they affect different ownership or architecture boundaries;
- one may be superseded while the other remains valid;
- combining them would produce an oversized omnibus decision.

Do not create one ADR per decision-ledger row mechanically. Several rows may refine one durable choice, and some rows may require no ADR.

## ADR authoring inputs

The ADR authoring Task must read:

- the parent Work Item;
- the decision-loop Task and exact relevant decision IDs;
- directly governing accepted Requirement, ADR, Specification, or investigation records;
- repository ADR authoring standards and lifecycle rules;
- existing ADRs that may cover or conflict with the decision.

Do not reopen every decision in the topic. Treat explicit decision-ledger entries as accepted inputs unless a direct authority conflict is found.

## ADR content contract

The ADR must preserve, according to repository format:

- the bounded question or context;
- the selected decision;
- material alternatives considered;
- the decisive rationale;
- consequences, limitations, and follow-up obligations;
- supersession relationships when applicable;
- target Specifications or affected design areas when repository relations support them.

Do not copy the entire interactive discussion transcript.

Do not put detailed current-state normative contracts into the ADR when they belong in Specifications.

## ADR and Specification separation

Use this rule:

```text
ADR = why this durable choice was made and what consequences follow
Specification = what behavior, structure, boundary, or constraint is currently normative
```

Example:

```text
ADR decision:
Use the database as canonical persistence because transactional history and
recovery consistency outweigh human-editable configuration convenience.

Specification statements:
The database is the canonical persistence authority.
The application shall not reconstruct canonical state from YAML files.
YAML persistence is outside the current scope.
```

The ADR does not eliminate the need to update Specification.

Specification text should not reproduce the complete ADR rationale.

## Routing record

For each decided item, record:

- decision ID;
- routing result;
- exact existing, new, amended, or superseding ADR ID when known;
- reason for `not_required` or `covered`;
- affected Specification targets;
- blocker when routing is `blocked`.

This may be recorded in the decision-loop Task, ADR authoring Task Evidence, or closure synchronization record according to writer ownership.

## Review expectations

Independent design review must verify:

- every durable decision received the correct ADR route;
- no required ADR was omitted;
- no unnecessary ADR duplicates accepted authority;
- the ADR scope is coherent;
- supersession is explicit and historically honest;
- ADR decisions and Specification statements agree;
- `not_required` decisions still appear in the appropriate canonical artifact when normative.

## Stop conditions

Stop ADR authoring and report the exact issue when:

- the decision summary permits materially different ADR decisions;
- accepted Requirements or ADRs conflict with the selected choice;
- an existing ADR may cover the choice but its status or scope is unclear;
- supersession is required but the replacement decision is not explicit;
- the correct ADR ownership boundary cannot be chosen without another user decision;
- authoring would require changing a Requirement or adjacent design topic outside scope.

Do not invent a rationale or alternative to make the ADR appear complete.

## Anti-patterns

Do not:

- create an ADR for every Specification edit;
- omit an ADR only to reduce file count;
- use patch size as the routing criterion;
- treat an implementation choice as a durable design decision without evidence;
- hide a reversal inside an in-place amendment;
- bundle independent choices into an omnibus ADR;
- split one inseparable choice across multiple ADRs;
- leave the current Specification stale after accepting the ADR;
- copy the conversation transcript into the ADR;
- mark an ADR accepted without the repository-required review or lifecycle process.

## Readiness checklist

Before handing off to Specification synchronization, verify:

- every `decided` item has a resolved routing outcome;
- every `required` decision has an exact ADR artifact and valid lifecycle state for the next gate;
- every `covered` decision points to an accepted, non-superseded ADR;
- every `not_required` decision has a recorded reason;
- superseded ADRs are handled according to repository policy;
- ADR granularity is coherent;
- affected Specification targets are explicit;
- no unresolved ADR decision remains hidden in prose.

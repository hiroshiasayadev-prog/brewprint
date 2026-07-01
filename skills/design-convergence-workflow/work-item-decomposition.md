# Work Item decomposition

## Purpose

Create or split child Work Items after a decision fixes their identity and completion boundaries.

This phase authors Work Item boundaries.
It does not decide whether a split is required or change the Task graph.

## Inputs

Read:

- parent Work Item;
- completed reconciliation decision;
- direct material sources for each child Work Item;
- accepted responsibility and completion boundaries;
- Work Item authoring standards;
- required parent-level routing.

Do not begin while Work Item identity or scope remains unresolved.

## Owned outcome

A `work_item_decomposition` Task owns one bounded parent-to-child decomposition outcome.

It may:

- create one or more child Work Items selected by the accepted decision;
- assign distinguishable Goals and Boundaries;
- preserve direct material `source_refs`;
- record coarse parent-to-child routing;
- leave child-internal Task composition undecided.

It must not:

- select whether the parent continues or splits;
- change existing Task dependencies, blockers, owners, or release conditions;
- author child-owned decisions, investigations, Specifications, ADRs, or implementation;
- duplicate child-internal Task graphs in the parent Work Item;
- perform review or lifecycle synchronization.

## Completion condition

Decomposition is complete when:

- every selected child Work Item exists;
- each child has one coherent Goal and completion boundary;
- child responsibilities do not overlap;
- direct material sources are preserved;
- parent-level routing is recorded;
- no child-owned deliverable was authored.

## Verification

Confirm:

- each child Work Item ID and path follow the active namespace rules;
- each child lists the source decomposition Task in `source_refs`;
- the parent records only coarse child routing;
- no Task-graph change was performed;
- no unresolved identity decision was introduced.

## Stop conditions

Stop when:

- the accepted decision does not fix child identity or completion boundaries;
- direct material sources cannot be selected safely;
- child responsibilities overlap materially;
- decomposition requires a new Requirement decision;
- Task-graph changes are required before or after decomposition.

Route unresolved judgment to `decision`.
Route Task-graph changes to `coordination`.

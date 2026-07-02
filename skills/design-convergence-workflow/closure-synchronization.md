# Closure synchronization

## Purpose

Mechanically propagate one accepted reviewed result into lifecycle, Evidence, relation, and successful Work Item closure state.

Closure synchronization is not design authoring, correction, coordination, review, or cancellation execution.

## Entry gate

Begin only after:

- integrated review returns `PASS`; or
- every required finding from `NEEDS REVISION` is independently `CLOSED`;
- exact accepted artifacts and verdict Evidence exist;
- no required decision remains `open` or `in_discussion`;
- blocked or deferred scope is permitted by the Work Item completion contract;
- the synchronization Task names every writable target.

If any input is missing, return `BLOCKED` or stop at the appropriate earlier owner.

## Synchronization Task contract

Use `task_type: synchronization`.

The Task owns one bounded propagation of accepted state.
Its contract must name:

- parent Work Item;
- accepted review Task or finding-closure review Tasks;
- exact ADRs and canonical artifacts in the accepted result;
- exact lifecycle or relation targets;
- exact writable files or records;
- completion conditions that are mechanically derivable;
- stop conditions for missing work or judgment.

## Writable boundary

A closure-synchronization Task may write only:

- its own status and Evidence;
- the parent Work Item status and closure Evidence;
- exact mechanically derivable lifecycle fields;
- exact mechanically derivable relation fields;
- exact completion-result fields named by its contract.

It may update another record only when:

- the target is explicitly named;
- the accepted value is uniquely determined by reviewed Evidence;
- no design or release judgment remains;
- the target's authoring standard permits the update.

## Prohibited writes

Closure synchronization must not:

- rewrite a completed decision Task or decision entry;
- change a `decided` item into a downstream status;
- alter authoring Task Evidence or outputs;
- alter a review verdict or finding set;
- mark a finding closed;
- create, remove, split, or amend the Task graph;
- choose a Requirement, Work Item, ADR, or Specification disposition;
- author or correct canonical design content;
- start production implementation;
- cancel a Work Item or Task;
- hide missing work by weakening Completion Conditions.

When any prohibited action is needed, stop and return to the correct owner.

## Accepted route selection

### Direct PASS route

When integrated review returns `PASS`:

- cite the accepted review Task;
- record exact reviewed artifacts;
- synchronize declared lifecycle and relations;
- evaluate Work Item Completion Conditions;
- close the Work Item only when all conditions are satisfied.

### Finding-closure route

When initial review returns `NEEDS REVISION`:

- preserve the initial verdict and finding set unchanged;
- cite correction Tasks as repair Evidence;
- cite independent closure-review dispositions;
- require every closure-blocking finding to be `CLOSED`;
- synchronize the final accepted combined state;
- do not rewrite the initial review as `PASS`.

## Closure Evidence

Record in the synchronization Task:

- accepted review route;
- exact initial verdict;
- exact finding IDs and final dispositions when applicable;
- exact ADR IDs and lifecycle states;
- exact Specification and originating-artifact refs;
- exact workflow-support files when in scope;
- deferred or blocked scope with reason and destination;
- exact lifecycle and relation changes made;
- Work Item Completion Condition results;
- confirmation that no canonical content, graph, or verdict was changed.

Record downstream refs here rather than in completed decision Tasks.

## Work Item closure

Set the Work Item to `done` only when:

- every required decision is terminal;
- every selected Investigation is complete;
- required ADRs and canonical artifacts are accepted for the closure contract;
- integrated review passed or every required finding is independently closed;
- all Completion Conditions are satisfied;
- no required downstream design work remains hidden;
- lifecycle and relations express the same accepted result.

Do not close the Work Item merely because all currently authored Tasks are done.
Do not use reviewed-success closure synchronization to set `cancelled`.
Cancellation follows the separate atomic lifecycle operation and does not require integrated-review PASS.

A validly blocked item may remain only when:

- the Completion Conditions explicitly allow it;
- the reason and dependency are recorded;
- current design remains coherent without it;
- the blocked item has an exact destination or future owner when required.

## Relation synchronization

Synchronize only relations uniquely determined by accepted artifacts.

Verify:

- Work Item `tasks` matches existing owned Tasks;
- each Task `work_item` points to the same parent;
- exact ADR and Specification refs are recorded by the owning downstream Evidence;
- supersession relations remain historically honest;
- no relation claims an artifact was reviewed when it was not in the review boundary.

Do not infer relations from file proximity or naming alone.

## Lifecycle synchronization

Update lifecycle only when the accepted route uniquely determines the new state.

Examples:

- synchronization Task becomes `done` after its propagation verifies;
- parent Work Item becomes `done` after every Completion Condition passes;
- an exact ADR migration field may change only when the canonical Specification projection and review Evidence exist and the Task contract explicitly owns it.

Do not use closure to advance unrelated records.

## Missing work route

When closure discovers:

- missing canonical content;
- missing Task ownership;
- incorrect dependency or writer order;
- an unresolved design choice;
- an open required finding;
- absent review Evidence;
- a Completion Condition that cannot be mechanically evaluated;

stop synchronization.

Route:

- unresolved judgment to `decision`;
- graph defects to `coordination`;
- named projection defects to `correction` when a valid finding exists;
- absent or incomplete review to `review`;
- missing authoring without new judgment to its authoring owner or coordination when no owner exists.

Do not repair the issue inside closure.

## Verification

After writing:

1. inspect the scoped diff;
2. confirm every changed target was declared writable;
3. confirm values are mechanically supported by accepted Evidence;
4. confirm completed decision, authoring, and review records are unchanged except this synchronization Task's own state;
5. confirm Work Item status matches Completion Condition results;
6. confirm relations are bidirectionally coherent where required;
7. confirm no production, canonical authoring, correction, or graph change occurred.

## Completion condition

Closure synchronization is complete when:

- every declared propagation is written and verified;
- every changed lifecycle and relation value matches the accepted route;
- closure Evidence names exact accepted inputs and changes;
- the parent Work Item status is correct;
- no prohibited artifact was changed;
- no missing work or judgment remains concealed.

## Stop conditions

Stop when:

- the review route is not accepted;
- a required finding remains open;
- an exact writable target is absent from the Task contract;
- a value requires judgment rather than mechanical derivation;
- canonical artifacts and review Evidence disagree;
- closing the Work Item would leave an unsatisfied Completion Condition;
- the requested lifecycle outcome is `cancelled` rather than reviewed-success `done`;
- relation synchronization would require guessing.

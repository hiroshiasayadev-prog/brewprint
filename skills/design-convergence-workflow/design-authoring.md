# Design authoring

## Purpose

Write already decided design into bounded canonical artifact sets without reopening judgment.

This companion governs ADR, Specification, Requirement, Work Item, and other originating-artifact authoring order.
Artifact-specific authoring standards remain authoritative for file shape and metadata.

## Preconditions

Begin authoring only when:

- the decision boundary is terminal;
- mandatory impact Investigation is complete;
- originating-artifact disposition is fixed;
- required graph owners and dependencies exist;
- ADR routing is complete for the authoring scope;
- exact target artifacts or sections are known;
- shared-writer order is fixed.

When these conditions are incomplete, return to decision, investigation, routing, or coordination rather than filling gaps during writing.

## Authoring Task boundary

Use `task_type: authoring` for one bounded artifact set from already decided inputs.

One Task may write several files when:

- every file projects the same accepted input set;
- one writer owns them;
- one completion judgment applies;
- one review boundary applies;
- no file introduces an independent decision.

Split authoring when writer, dependency, completion judgment, artifact authority, or release boundary differs.

Do not combine authoring with:

- unresolved decision work;
- impact Investigation;
- graph coordination;
- implementation;
- independent review;
- finding correction;
- lifecycle synchronization.

## Authoring order

Default order:

1. clarify or supersede existing ADR authority when required;
2. create newly routed ADRs;
3. amend originating Requirement or Work Item when the accepted disposition requires it;
4. write current normative Specification state;
5. author or activate workflow support files such as skills when they are direct targets;
6. produce exact changed-artifact Evidence for integrated review.

Change the order only when dependencies or shared-writer serialization require it.

## ADR authoring

ADR authoring consumes completed routing.

For each routed boundary:

- use the exact create, amend, reuse, or supersede disposition;
- preserve the bounded context and selected decision;
- preserve material alternatives and decisive rationale;
- state consequences and follow-up obligations;
- preserve honest supersession history;
- identify affected design areas without duplicating full Specification text.

Do not copy the interactive transcript.
Do not invent alternatives or rationale absent from accepted inputs.

An ADR explains why a durable choice was made.
It does not replace current normative Specification content.

## Specification authoring

After required ADR work:

- state current normative behavior, structure, ownership, constraints, and exclusions;
- project every applicable accepted decision;
- remove or revise stale contradictory text;
- preserve Requirement intent;
- use active `spec:` refs according to repository policy;
- add ADR relationships when required by the canonical format;
- keep history and rejected alternatives out of normative sections;
- identify exact changed sections for review.

Do not mark the Work Item complete merely because Specification writing finished.

## Originating Requirement authoring

Amend an existing Requirement only when the accepted reconciliation decision preserves the same motivating problem and Required Outcome.

Create a new or follow-up Requirement when the accepted identity boundary requires it.
Do not make the identity decision during authoring.

Preserve:

- the motivating problem;
- required outcome identity;
- explicit scope and exclusions;
- relation to the correct Work Item.

## Work Item authoring

Amend the current Work Item only when its accepted resolution and completion identity remain coherent.

Authoring may update:

- Goal, Boundary, Impact Scope, and Completion Conditions fixed by decision;
- Task candidates or flow fixed by coordination;
- exact artifact and review targets;
- conditional routes already accepted.

Do not let Work Item authoring perform graph coordination implicitly.
Task creation, dependency changes, and writer order belong to a coordination owner.

## Skill and workflow-support authoring

When an accepted design targets a repository skill or instruction source:

- keep the entry document focused on activation, end-to-end flow, and phase routing;
- place detailed judgment authority in bounded companions;
- reference canonical authoring standards rather than duplicating their schemas;
- activate the successor only after all required files are substantive;
- remove replaced authority only after successor readiness is verified;
- do not retain a deprecated stub when the accepted replacement policy forbids one.

Instruction-pointer changes and old-authority removal may share the same Task only when they are one atomic activation outcome.

## Shared writers

Follow the writer order persisted by coordination.

Each later writer must:

- read the final accepted output of earlier writers;
- preserve earlier accepted semantics;
- write only its assigned sections or consistency effects;
- stop when preservation would require reinterpretation or reversal.

A required semantic change returns to investigation and decision.
Do not resolve it as merge cleanup.

## Ambiguity stop rule

Stop authoring when decided inputs permit several materially different outputs.

Report:

- exact ambiguous input;
- affected artifact and section;
- materially different interpretations;
- current owner status;
- whether coordination is needed to create or repair the decision route.

Do not choose the most convenient interpretation.

## Authoring Evidence

The authoring Task records:

- exact inputs and decision IDs;
- exact created or updated artifacts;
- exact sections when relevant;
- applied ADR dispositions;
- writer-order dependencies;
- verification performed;
- unresolved blockers or excluded scope;
- confirmation that no independent review occurred.

Do not write downstream review, finding, or closure progress into completed decision Tasks.

## Completion condition

Authoring is complete when:

- every assigned artifact satisfies its canonical authoring requirements;
- every accepted input is projected exactly once in the correct owner;
- no stale contradiction remains within the Task boundary;
- ADR and Specification roles remain distinct;
- shared-writer obligations are preserved;
- exact review inputs are available;
- no unresolved judgment remains hidden in prose.

## Verification

Inspect scoped files and diffs.
Confirm:

- artifact IDs, metadata, and sections match canonical standards;
- output files match the Task contract;
- no prohibited artifact changed;
- no decision, coordination, correction, review, or synchronization work was absorbed;
- no completed workflow record was rewritten to mirror downstream progress;
- integrated review can read exact final artifacts.

## Stop conditions

Stop when:

- the canonical target is unknown;
- accepted authorities conflict;
- authoring would alter a completed decision outcome;
- an ADR reversal lacks supersession authority;
- shared-writer order is missing;
- a target falls outside the Work Item boundary;
- the Task would need an independent completion judgment for part of its output.

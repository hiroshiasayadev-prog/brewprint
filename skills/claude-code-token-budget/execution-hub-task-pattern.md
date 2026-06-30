# Execution hub Task pattern

This document defines the persistent execution-graph pattern used before issuing Claude Code prompts for multi-executor implementation, correction, or integration work.

It is a required companion to `SKILL.md` and `executor-ready-task-design.md` when the application trigger below is met.

## Purpose

Keep implementation ownership, model routing, writer boundaries, dependency order, verification ownership, independent review, and release state in repository Design Records instead of leaving them only in chat prompts.

The prompt given to Claude Code is an execution view of an accepted Task. It is not the canonical source of the execution contract.

## Application trigger

Evaluate this pattern before issuing an implementation or finding-correction prompt.

Apply the pattern when any of the following is true:

- one existing Task is split across multiple executors;
- more than one model class is routed within the same implementation or correction scope;
- writable files, fixtures, generated artifacts, or lifecycle records need explicit writer allocation;
- focused verification and broader package, integration, or release verification have different owners;
- parallel and sequential execution are mixed;
- one executor consumes another executor's type, API, fixture, generated artifact, or other output;
- an existing production Task must become an aggregate verification or lifecycle gate;
- correction work contains multiple leaves followed by independent closure review;
- review acceptance must explicitly release some leaves while keeping others blocked.

Do not issue a ready-to-run implementation or correction prompt while an applicable execution graph exists only in conversation text.

## Non-trigger cases

Do not create a hub merely because a patch is small or because a lower-cost model is available.

The pattern is normally unnecessary when one executor can complete one responsibility with:

- one accepted contract;
- one exact writable boundary;
- no shared writer;
- no dependent executor output;
- one focused verification owner;
- no separate aggregate gate beyond the existing Task lifecycle.

When uncertain, compare the cost of persistent coordination against the risk of hidden ownership, duplicated verification, or chat-only state.

## Artifact topology

The preferred persistent structure is:

```text
Execution reorganization or correction Work Item
├─ authoring or scope-freeze Task
├─ independent review Task
└─ release synchronization Task

Target implementation Work Item
├─ executor Task A
├─ executor Task B
├─ optional integration or aggregate verification Task
└─ existing lane lifecycle or closure Task
```

The coordinator triplet may remain in the target Work Item when the scope is small and its ownership stays coherent. Create a separate execution-reorganization Work Item when the graph spans multiple existing Tasks, multiple lanes, or a substantial redesign of execution ownership.

Do not model Tasks as owning child Tasks. The Work Item is the canonical owner of the Task graph.

## Canonical ownership

Persist each concern in the following artifact:

| concern | canonical owner |
|---|---|
| complete execution graph and Task relations | Work Item |
| exact implementation contract | executor Task |
| independent judgment and findings | independent review Task |
| accepted start order and released or deferred leaves | release synchronization Task |
| focused implementation evidence | executor Task Evidence after execution |
| broader package or integration evidence | aggregate or integration Task |
| final lifecycle and finding disposition synchronization | explicitly named closure synchronization Task |

A Claude Code prompt may summarize these records, but must not introduce essential writable paths, symbols, tests, dependencies, or ownership that are absent from the persistent graph.

## Coordinator triplet

### Authoring or scope-freeze Task

This Task creates or revises the execution graph without starting production implementation.

It must define:

- existing Task or finding scope being reorganized;
- exact executor-card IDs and filenames when new Tasks are created;
- exact responsibility of each executor Task;
- model routing and why each leaf is suitable for that model class;
- exact writer ownership;
- predecessor and consumer dependencies;
- focused test ownership;
- aggregate, integration, and release gate ownership;
- protected existing Tasks and artifacts;
- the next independent reviewer.

It must not perform production implementation or independent review.

### Independent review Task

This Task independently judges whether the graph is executable without hidden design work or broad repository exploration.

It must verify:

- every leaf is self-contained;
- exact writer uniqueness;
- dependency acyclicity;
- model routing by contract density and ambiguity;
- fixture and generated-artifact producer order;
- focused and aggregate verification separation;
- reviewer and release ownership;
- stop conditions and failure owners;
- preservation of existing public Task IDs and lifecycle roles when applicable.

It records `PASS` or `NEEDS REVISION` with blocking, major, minor, and advisory findings.

It does not correct findings, implement production, or release executors.

### Release synchronization Task

This Task records independent review acceptance and makes start eligibility explicit.

It must:

- require review `PASS` without blocking or major findings;
- list executor Tasks released for execution;
- state allowed parallelism and required sequence;
- list deferred leaves and their exact blockers;
- preserve producer and consumer order;
- keep executor statuses unchanged until their own execution begins;
- synchronize only explicitly owned Work Items or lifecycle records.

It does not create executor cards, correct findings, rerun independent review, or implement production.

## Executor Task profile

An executor Task is the persistent source of the implementation contract.

Its `## Work` section should contain the following structured subsections when applicable:

```text
### Read
### Change
### Implement
### Test
### Stop
### Prohibited operations
### Output
```

### Read

List only the exact Task itself and exact implementation, test, fixture, or generated-input files needed for execution.

A fully frozen executor Task should not require the executor to read:

- `prompt_chappy.md` merely to recover missing implementation details;
- parent Work Items;
- Requirements, ADRs, or broad Specifications;
- the old large Task that was decomposed;
- sibling executor Tasks;
- repository-wide authority.

Include a governing contract only when the executor must directly apply it and the Task cannot contain a safe, complete execution contract without duplicating normative content.

### Change

List every writable production, test, fixture, schema, generated artifact, or script path.

No path may have more than one writer in the same execution wave.

### Implement

Freeze, as applicable:

- exact symbols and seams;
- request and response shapes;
- field types and cardinalities;
- filtering, ordering, deduplication, projection, warning, and error behavior;
- compatibility behavior to preserve;
- obsolete behavior and symbols to remove;
- protected helpers, files, and refactors;
- fixture and generated-artifact contracts.

The executor performs accepted implementation decisions. It does not discover the product or architecture contract while coding.

### Test

Define:

- exact test files;
- exact test cases or behavior assertions;
- exact focused command;
- expected success condition;
- the downstream owner of broader package, integration, or release verification.

Do not assign the same full suite to every leaf.

### Stop

Require `BLOCKED` without broader exploration when:

- a required file outside `Change` must be modified;
- a fixed symbol or contract is ambiguous;
- a predecessor output is absent;
- a failure belongs to another writer;
- an accepted contract must change;
- a dependency, manifest, lock, or generated-artifact ownership change is required;
- the exact verification environment cannot run.

The executor names the missing decision, input, owner, command failure, or exact file.

### Prohibited operations

Prohibit only relevant boundary violations, including as applicable:

- repository-wide traversal;
- broad recursive grep;
- unrelated refactor or cleanup;
- Design Record or lifecycle changes;
- sibling implementation;
- independent review;
- stage or commit.

### Output

Default concise output:

1. Result
2. Changed files
3. Verification command and exit status
4. Blocker
5. Review readiness

Include full logs only when a failure excerpt is needed.

## Writer ownership

The authoring Task or Work Item must record a single-writer map for the wave.

```text
exact path or artifact family -> sole writer Task
```

Apply this to:

- production files;
- test files;
- fixture families;
- schemas and examples;
- sidecars and manifests;
- generated outputs;
- shared helper files;
- Work Item and lifecycle synchronization records.

When two leaves need the same file, choose one of these strategies before release:

1. sequence both responsibilities under one writer;
2. move the shared write into a dedicated integration Task;
3. redesign the boundary so files are disjoint.

Do not use later merge-conflict resolution as the ownership strategy.

## Dependency and parallelism

Leaves may execute in parallel only when all of the following hold:

- writable paths are disjoint;
- neither leaf consumes output produced by the other;
- fixture and generated-artifact ownership is disjoint;
- no shared lifecycle record is written concurrently;
- the integration point and integration owner are explicit.

Otherwise, persist a predecessor edge and execute sequentially.

Parallelism is an optimization after ownership correctness. Do not create extra sessions when the token benefit is smaller than the coordination and verification cost.

## Model routing

Route by contract density, ambiguity, and judgment ownership rather than patch size.

### Lower-cost model leaf

Use a lower-cost model only when:

- the accepted contract is explicit;
- exact files and symbols are fixed;
- exact tests and commands are fixed;
- no architecture, semantic, or ownership decision remains;
- failures can be stopped and routed without investigation outside the boundary.

### Stronger coding model

Use Sonnet, Codex, or another stronger coding model when the Task requires:

- semantic production choices within an accepted boundary;
- shared-writer cleanup;
- cross-file helper ownership;
- ambiguous compile or test diagnosis;
- cross-package integration;
- correction requiring contract interpretation.

### Strong reasoning or reviewer model

Keep the following on a strong reasoning model:

- architecture review;
- independent review of the execution graph;
- finding disposition;
- authority escalation;
- cross-lane integration judgment;
- final acceptance.

A one-file change is not automatically lower-cost-model work.

## Focused and aggregate verification

Each implementation leaf owns only the focused verification required to prove its behavior.

A named aggregate or integration Task owns, as applicable:

- full affected-package tests;
- cross-leaf integration tests;
- generated-artifact drift checks;
- architecture checks;
- final scoped whitespace and changed-path verification;
- release-readiness judgment.

The aggregate Task is read-only unless repair ownership is explicitly assigned elsewhere.

### Failure owner map

An aggregate Task must contain an exact owner map when a failure must be routed without reading sibling executor Tasks.

```text
focused test / production file / fixture or generated artifact -> owning executor Task
```

If a diagnostic names no owned file or mapped behavior, report it without guessing an owner and stop `BLOCKED`.

## Correction and re-review

When findings require more than one correction leaf, persist this flow before issuing correction prompts:

```text
finding set
  -> correction authoring or scope-freeze
  -> correction executor leaves
  -> focused verification
  -> aggregate correction verification when required
  -> independent finding closure review
  -> closure synchronization
```

The persistent graph must name:

- exact finding IDs;
- exact correction leaves;
- writable boundaries;
- model routing;
- predecessor order;
- focused verification owner;
- aggregate verification owner;
- closure reviewer;
- lifecycle synchronization owner.

Correction executors must not claim findings are closed. Only independent re-review may close a finding.

## Closure

A coordinator Work Item or hub scope is not complete merely because executor code exists.

Closure requires:

- all released executor Tasks completed or explicitly superseded;
- required focused verification recorded;
- required aggregate or integration gate passed;
- independent review findings dispositioned;
- release or closure synchronization completed;
- parent Work Item relations and evidence synchronized by the sole named writer.

Do not copy canonical child Task statuses into the parent Work Item as duplicated lifecycle state.

## Anti-patterns

Do not:

- keep A/B/C slice ownership only in chat;
- issue prompts whose essential contract is absent from the Task;
- create a hub for every trivial one-executor patch;
- create one author-review-release triplet per tiny mechanical edit when a coherent wave shares one review and release gate;
- assign overlapping writers and plan to resolve conflicts later;
- make every leaf run the same full suite;
- let an aggregate verifier repair failures;
- let a release Task correct findings or advance executor lifecycle status;
- route final architecture or acceptance work to a lower-cost model;
- reserve large Task ID ranges without a bounded, reviewed execution need.

## Readiness checklist

Before issuing the first implementation or correction prompt for an applicable wave, confirm:

- the Work Item contains the persistent Task graph;
- authoring, review, and release ownership are explicit;
- every executor Task exists at an exact path;
- every executor Task contains an exact execution contract;
- writer ownership is unique;
- dependencies are acyclic;
- released parallel leaves are writer-disjoint;
- model routing is justified;
- focused verification is assigned per leaf;
- broader verification has one named owner;
- a failure owner map exists when aggregate routing needs it;
- independent review recorded acceptance;
- release synchronization identifies executable and deferred leaves;
- the Claude Code prompt introduces no essential contract absent from Design Records.

If any required check fails, return to scope-freeze or Task authoring. Do not issue the implementation prompt.

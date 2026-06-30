# Executor-ready Task design

This document is a required companion to `SKILL.md` for Claude Code implementation prompt and Task-graph authoring.

When `execution-hub-task-pattern.md` applies, all three documents are required.

## Read condition

Read this document together with `SKILL.md` before preparing or revising:

- a Claude Code implementation prompt;
- a Task graph intended for Haiku, Sonnet, Opus, or another Claude Code model;
- implementation Task decomposition or parallelization;
- model routing;
- writer ownership or integration-gate ownership;
- test ownership or verification ownership;
- execution-hub applicability, coordinator ownership, or release synchronization.

Re-read it when the phase, model routing, Task partitioning, writer ownership, verification ownership, or execution-hub applicability changes.

## Executor-ready gate

Do not issue an implementation prompt until the target Task or accepted execution contract defines:

- exact production files allowed to change;
- exact test or fixture files allowed to change;
- exact production symbols or seams to add, modify, or remove;
- exact behavior to implement;
- input, output, ordering, warning, and error behavior;
- accepted compatibility behavior that must remain unchanged;
- predecessor and integration dependencies;
- focused verification commands;
- the owner of broader package, integration, and release verification.

Do not ask Claude Code to discover the implementation contract while coding.

When implementation ownership, symbols, fixtures, tests, or verification are unknown, run a read-only inventory or scope-freeze phase first. When Design Records are being authored or reorganized, persist the resulting executor contract in the Task instead of leaving essential behavior only in a chat prompt.

## Persistent execution-graph gate

Before issuing an implementation or finding-correction prompt, evaluate `execution-hub-task-pattern.md`.

When the trigger applies, require all of the following before prompt issuance:

- a Work Item-owned persistent Task graph;
- exact executor Tasks with accepted execution contracts;
- explicit model routing;
- unique writer ownership;
- focused and broader verification ownership;
- an independent graph review;
- release synchronization that names executable and deferred leaves.

A chat-only decomposition, coordinator summary, or handoff prompt is not an accepted execution graph.

The Work Item owns the Task graph. Executor Tasks own implementation contracts. Review Tasks own judgment. Release Tasks own start eligibility.

Do not create a hub for a single self-contained leaf that has no shared writer, dependent output, separate aggregate gate, or release-coordination need.

## Haiku-ready leaves

When designing an implementation Task graph, actively look for safe low-context leaves that Haiku can execute.

A Task is Haiku-ready only when:

- the governing contract is explicit and accepted;
- no architecture, product, or ownership decision remains;
- writable files are exact;
- parallel writable boundaries do not overlap;
- implementation symbols and required behavior are exact;
- test cases and verification commands are exact;
- all predecessors are complete;
- fixtures, generated artifacts, and shared files have one named writer;
- the agent can stop without guessing when work would cross the declared boundary.

A small patch is not automatically Haiku-ready. Contract density and ambiguity matter more than line count.

Use Sonnet, Codex, or another stronger model when the Task requires:

- contract interpretation;
- semantic production choices;
- architecture decisions;
- ownership resolution;
- ambiguous failure diagnosis;
- cross-package integration;
- unresolved fixture or generated-artifact ownership.

When a larger Task contains mechanical and semantic work, split only along real dependency and writer boundaries. Route the mechanical leaves to Haiku and keep semantic or integration ownership on the stronger model.

## Exact implementation specification

Every implementation Task and implementation prompt identifies, as applicable:

- functions, methods, types, handlers, adapters, schemas, fixtures, or generated artifacts to add or change;
- obsolete symbols or behavior to remove;
- accepted request and response shapes;
- filtering, ordering, deduplication, projection, warning, error, and boundary behavior;
- existing behavior and shared helpers that must remain;
- protected files, protected symbols, and forbidden refactors.

Do not replace this with a generic instruction to inspect the repository and implement the Task appropriately.

The implementation agent executes an already-made design decision. It is not the implicit owner of missing Task design.

## Exact test specification

Every implementation Task defines its verification target before execution.

State:

- test files to add or change;
- behavior each test must prove;
- test function names when already known or safely fixed;
- the exact focused test command;
- package-level verification required by the Task;
- the downstream Task that owns broader integration or release verification.

Leaf Tasks run focused verification first.

Do not make every parallel leaf rerun the same full package or repository suite. Assign each broader gate to one explicit owner. Distinguish focused executor verification from package, integration, and release gates.

## Parallelism and writer ownership

Tasks may run in parallel only when:

- writable file sets do not overlap;
- fixture and generated-artifact writers do not overlap;
- neither Task depends on types, APIs, fixtures, or output produced by the other;
- the integration point and integration owner are explicit.

Every shared file, fixture family, sidecar, manifest, generated output, or lifecycle record has one named writer for the wave.

When ownership overlaps, sequence the Tasks or move the shared write into a dedicated integration Task. Do not use later conflict resolution as the ownership model.

## Separate execution slices

Keep these responsibilities separate:

1. implementation;
2. independent review;
3. finding correction;
4. finding closure review;
5. closure synchronization.

Implementation sessions do not perform independent review or lifecycle closure.

Review sessions do not modify files.

Finding-correction sessions do not claim closure before a separate closure review.

Design Record lifecycle and Evidence updates belong to closure synchronization unless explicitly assigned otherwise.

## Readiness checklist

Before returning an implementation prompt, verify:

- the Task is executor-ready;
- the exact code boundary is stated;
- the exact behavior is stated;
- the exact tests and commands are stated;
- model routing is justified by contract density and ambiguity;
- shared writers and integration ownership are resolved;
- broader verification has one named owner;
- architecture or contract discovery is not delegated to the implementation agent;
- execution-hub applicability was evaluated;
- when applicable, the persistent graph was reviewed and released;
- the prompt adds no essential contract absent from the executor Task.

When any check fails, return to inventory, scope-freeze, or Task authoring instead of issuing the implementation prompt.

# PRODUCT-TASK-SPEC-021-04: Coordinate executor-ready implementation graph

- **id**: PRODUCT-TASK-SPEC-021-04
- **status**: blocked
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-021-03
- **outputs**:
  - PRODUCT-WORK-SPEC-021
  - PRODUCT-TASK-SPEC-021-04

## Goal

Materialize one executor-ready implementation graph from terminal decisions and the concluded Investigation.

## Work

- Consume terminal T02 decisions without reopening them.
- Consume PRODUCT-INV-SPEC-009 Evidence without adopting unsupported choices.
- Determine whether one implementation Task satisfies the accepted boundary.
- Evaluate the execution-hub trigger from accepted decisions and Investigation Evidence.
- Materialize the required executor leaf count and exact dependencies.
- Assign exact model routing by contract density and ambiguity.
- Assign one writer for every production, test, fixture, configuration, and generated artifact path.
- Name exact production and test files.
- Name exact symbols and implementation seams.
- Assign focused verification ownership.
- Assign aggregate or integration verification ownership when required.
- Materialize one objective verification Task when the accepted gate requires separate ownership.
- Materialize one independent implementation review Task.
- Materialize one closure synchronization Task.
- Persist dependency, writer, review, and release order in W021.
- Keep correction and finding-closure review Tasks conditional on named findings.

When the execution-hub trigger applies, materialize this persistent route before any Claude Code implementation prompt:

- scope-freeze or executor-contract authoring;
- independent graph review;
- release synchronization;
- released executor Tasks.

This Task must not:

- introduce an implementation decision absent from T02 or PRODUCT-INV-SPEC-009;
- perform production implementation;
- author or modify checklist content;
- issue an independent review verdict;
- release an unreviewed execution-hub graph;
- create a Claude Code implementation prompt;
- perform finding correction, lifecycle closure, stage, or commit work.

## Done condition

- W021 contains one persistent executor-ready implementation graph.
- Every implementation Task has one bounded implementation outcome.
- Exact model routing, writers, files, symbols, seams, dependencies, and focused verification are persisted.
- Broader verification has one named owner when required.
- Objective verification, independent implementation review, and closure synchronization owners exist.
- The execution-hub route includes exact authoring, independent graph review, and release-synchronization owners when the trigger applies.
- Conditional finding work remains unmaterialized.
- No Claude Code implementation prompt is required to invent missing contract details.

## Verification

- Confirm every materialized Task has one canonical `task_type` and one completion judgment.
- Confirm exact file writers do not overlap in one execution wave.
- Confirm dependencies are acyclic and producer-consumer order is explicit.
- Confirm model routing follows contract density and ambiguity.
- Confirm focused and aggregate verification ownership do not overlap.
- Confirm independent review and release ownership remain separate from implementation.
- Confirm no implementation prompt exists before any required graph review and release synchronization.
- Confirm no production, checklist, ADR, Requirement, or Specification content changed.

## Evidence

- PRODUCT-TASK-SPEC-021-01 created this coordination owner.
- T04 remains downstream of terminal T02 decisions and concluded PRODUCT-INV-SPEC-009.
- PRODUCT-INV-SPEC-009 exposed that no standalone validator app namespace or app-local design authority exists.
- The user clarified that PRODUCT owns conceptual design only.
- App namespace bootstrap and app-local implementation Specifications must precede implementation planning.
- This Task would cross the PRODUCT ownership boundary by materializing executor-ready implementation work directly.
- PRODUCT-TASK-SPEC-021-05 repairs the graph.
- PRODUCT-TASK-SPEC-021-06 owns the missing PRODUCT-to-app boundary and namespace decisions.
- PRODUCT-TASK-SPEC-021-07 owns graph materialization after T06.
- This Task must not execute unless a later accepted decision explicitly restores its responsibility.
- No executor Task ID, implementation graph, or Claude Code prompt has been produced.

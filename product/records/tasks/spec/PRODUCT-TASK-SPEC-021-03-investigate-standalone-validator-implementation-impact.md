# PRODUCT-TASK-SPEC-021-03: Investigate standalone validator implementation impact

- **id**: PRODUCT-TASK-SPEC-021-03
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-021
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-021-02
- **outputs**:
  - PRODUCT-TASK-SPEC-021-03
  - PRODUCT-INV-SPEC-009

## Goal

Produce one bounded implementation impact Investigation for an executor-ready standalone validator graph.

## Work

- Create PRODUCT-INV-SPEC-009 for the bounded research question below.
- Use the terminal T02 decision ledger and accepted W019 and W020 contracts as fixed inputs.
- Inspect active app namespace and source placement.
- Inspect existing standalone tools and command patterns.
- Inspect current language and package boundaries.
- Identify the checklist filesystem loading seam.
- Identify the Task Markdown input seam.
- Identify local LLM runtime availability and invocation seams.
- Identify structured JSON parsing seams.
- Inspect configuration and environment handling.
- Identify test framework and fixture placement.
- Identify one runtime smoke-test command.
- Identify focused and aggregate verification boundaries.
- Identify shared writer candidates.
- Evaluate execution-hub trigger applicability.
- Record uncertainty and blockers without adopting decisions.

This Task must not:

- change any T02 decision;
- author production or test code;
- create source directories;
- materialize executor Tasks;
- issue a Claude Code implementation prompt;
- perform independent review, correction, synchronization, stage, or commit work.

### Bounded research question

```text
Given the accepted standalone-validator contract and terminal implementation-boundary decisions, what repository placement, existing implementation seams, local runtime constraints, test surfaces, writer boundaries, and verification commands constrain an executor-ready implementation graph?
```

## Done condition

- PRODUCT-INV-SPEC-009 exists and is `concluded`.
- The Investigation answers only the bounded research question.
- Every required repository seam and runtime constraint has direct Evidence.
- Exact production and test file candidates are identified.
- Focused and aggregate verification candidates are distinguished.
- Shared writers and execution-hub trigger facts are explicit.
- Every uncertainty or blocker has a named downstream owner.
- No decision or implementation is adopted.

## Verification

- Confirm PRODUCT-INV-SPEC-009 uses the canonical Investigation structure.
- Confirm the Investigation consumes terminal T02 decisions without changing them.
- Confirm every claimed file, symbol, seam, and command has repository Evidence.
- Confirm writer and verification candidates remain Investigation findings.
- Confirm no production, test, checklist, ADR, Requirement, or Specification file changed.
- Confirm no executor Task or Claude Code prompt was created.

## Evidence

- `prompt_chappy.md` was read first. `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational for current authoring, so repository-local filesystem read and write operations were used under the accepted fallback policy.
- PRODUCT-TASK-SPEC-021-01 reserved PRODUCT-INV-SPEC-009 for this Task.
- The Investigation namespace contained PRODUCT-INV-SPEC-001 through PRODUCT-INV-SPEC-008 before writing; PRODUCT-INV-SPEC-009 was unused.
- PRODUCT-TASK-SPEC-021-02 was `done`, and D-001 through D-019 were all terminal `decided` inputs.
- Created `product/records/investigations/spec/PRODUCT-INV-SPEC-009-standalone-validator-implementation-impact.md` with `status: concluded` and all ten canonical Investigation sections.
- The Investigation found complete eleven-type checklist coverage and a deterministic common-plus-type checklist loading seam.
- The Investigation found that current DRMCP provides reference patterns for Go command, stdio, JSON-RPC, strict tool decoding, and tests, but its current Task parser and `TaskDetail` do not retain `task_type` and its `internal` packages are not an accepted standalone dependency.
- The Investigation identified candidate production paths, test and fixture paths, focused commands, aggregate commands, Windows build verification, writer boundaries, and execution-hub trigger facts without adopting an implementation option.
- The active namespace profile does not assign a dedicated standalone validator app namespace. This is recorded as an executor-readiness blocker for T04 rather than resolved by this Investigation.
- Verification confirmed no explicit Investigation `id` metadata, no Task ID in Investigation canonical reference metadata, substantive `Investigation scope` and `Findings`, and no change to T02 decisions.
- No production, test, fixture, checklist, script, ADR, Requirement, Specification, Work Item, or Task-graph change was performed.
- No executor Task, implementation prompt, review, closure, stage, or commit was performed.
- The only write targets for this Task were PRODUCT-INV-SPEC-009 and this Task record.

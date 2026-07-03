# TRV-TASK-SPEC-003-08: Decompose architecture-derived contract Specification Work Item

- **id**: TRV-TASK-SPEC-003-08
- **status**: done
- **date**: 2026-07-03
- **work_item**: TRV-WORK-SPEC-003
- **task_type**: work_item_decomposition
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-001-13
- **outputs**:
  - TRV-TASK-SPEC-003-08
  - TRV-WORK-SPEC-005

## Goal

Create one replacement Work Item for architecture-derived TRV contract Specification design.

## Work

- Create `TRV-WORK-SPEC-005` for architecture-derived contract Specification design.
- Make Specification topic-tree and Markdown-placement judgment the first design boundary.
- Preserve W002 architecture as the controlling design input.
- Preserve W003 as historical Evidence of the incomplete external-interface-oriented route.
- Keep exact implementation-ready detail in W004.
- Leave W005 child Task materialization to a later coordination owner.

This Task must not change the parent Task graph, author child Specifications or ADRs, decide placement, review, synchronize lifecycle, or perform implementation.

## Done condition

- TRV-WORK-SPEC-005 exists with one coherent contract-Specification design goal.
- W005 lists this decomposition Task as a direct source.
- W005 references the reviewed architecture and retired W003 route as direct material sources.
- W005 owns an independent completion and review boundary.
- No W005-owned canonical deliverable is authored.

## Verification

- Confirm the W005 ID and path follow active Work Item rules.
- Confirm all canonical Work Item sections exist.
- Confirm W005 does not overlap W004 implementation-ready detail.
- Confirm no parent execution relation or child canonical artifact was created by this Task.

## Evidence

- The user explicitly selected a separate Work Item after identifying that Specification placement must precede contract authoring.
- TRV-TASK-SPEC-001-13 created this decomposition owner and retired the incomplete W003 execution route.
- Created `trv/records/work-items/spec/TRV-WORK-SPEC-005-define-architecture-derived-trv-contract-specifications.md`.
- W005 was created with W002 architecture, W003, and this Task as direct material sources.
- W005 defines placement-first contract Specification design and leaves implementation-ready detail in W004.
- No ADR, Specification, review, synchronization, implementation, stage, or commit work occurred.
- Result: `PASS`.

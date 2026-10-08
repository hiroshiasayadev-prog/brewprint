# PRODUCT-TASK-SPEC-030-01: Frame hierarchical USDM coverage semantics

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-030
- **task_type**: decision
- **estimate**: 0.25d
- **depends_on**:
- **outputs**:
  - PRODUCT-TASK-SPEC-030-02

## Goal

Fix the framing disposition and downstream contract for PRODUCT-REQ-SPEC-016.

## Work

- Compare the requested hierarchy and coverage behavior with PRODUCT-REQ-SPEC-016.
- Resolve current USDM format and coverage facts from canonical Specifications.
- Select the source disposition, unknown handling, and downstream route.
- Materialize the uniquely determined Work Item decomposition Task.

## Done condition

Every framing decision item is terminal and the downstream Work Item contract is fixed.

## Verification

The decision ledger contains one disposition, one aligned downstream boundary, and one next route with no unresolved framing question.

## Evidence

| item | state | decision | reason |
|---|---|---|---|
| F001 Desired Outcome | decided | Add pure USDM requirement decomposition and derived coverage semantics. | The requested behavior is the purpose of PRODUCT-REQ-SPEC-016. |
| F002 Outcome alignment | decided | Desired Outcome matches the Required Outcome. | No Requirement amendment is needed. |
| F003 Source disposition | decided | `proceed`. | The Requirement is accepted and has an actionable design boundary. |
| F004 Unknown handling | decided | Resolve from scoped canonical reads; no formal Investigation. | Current artifact-format and coverage-format contracts state the exact flat baseline being changed. |
| F005 Goal | decided | Establish durable hierarchy and effective coverage semantics in ADR and USDM Specifications. | One coherent design outcome. |
| F006 Boundary | decided | ADR and normative Specification semantics only; no tool implementation or migration. | Implementation has a separate completion judgment. |
| F007 Completion Condition | decided | Accepted ADR and synchronized Specifications pass integrated independent review. | This is an observable design closure boundary. |
| F008 Direct source | decided | PRODUCT-REQ-SPEC-016. | The Requirement directly motivates the downstream work. |
| F009 Initial route | decided | Design convergence. | The change establishes durable validation and coverage semantics across several Specifications. |

The accepted framing decision directly materialized PRODUCT-TASK-SPEC-030-02 under the framing exception.

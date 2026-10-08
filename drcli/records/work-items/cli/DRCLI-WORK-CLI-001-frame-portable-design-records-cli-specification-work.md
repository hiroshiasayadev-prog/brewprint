# DRCLI-WORK-CLI-001: Frame portable Design Records CLI specification work

- **status**: done
- **date**: 2026-10-07
- **source_refs**:
  - DRCLI-REQ-CLI-001
- **impact_refs**: []
- **tasks**:
  - DRCLI-TASK-CLI-001-01
  - DRCLI-TASK-CLI-001-02

## Goal

Fix how DRCLI-REQ-CLI-001 proceeds into repository-persistent Specification design.

## Boundary

This Work Item owns only Requirement framing and downstream Work Item creation.

It does not own:

- DRCLI Specification content;
- implementation;
- DRMCP modification or replacement;
- production CLI code;
- independent design review of the downstream Specification.

## Impact Scope

No canonical Specification is modified by this framing Work Item.

The downstream design Work Item owns the new DRCLI Specification area.

## Task flow

```text
DRCLI-TASK-CLI-001-01 framing decision
  -> DRCLI-TASK-CLI-001-02 create downstream Specification Work Item
  -> framing closure
```

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `DRCLI-TASK-CLI-001-01` | `decision` | Fix source disposition, downstream contract, unknown handling, and route. | none |
| `DRCLI-TASK-CLI-001-02` | `work_item_decomposition` | Create the downstream Specification convergence Work Item. | T01 |

## Completion Condition

- DRCLI-REQ-CLI-001 has one explicit source disposition.
- Desired Outcome and Required Outcome alignment is recorded.
- The downstream Goal, Boundary, Completion Condition, unknown handling, and route are fixed.
- The downstream Specification Work Item exists.
- No DRCLI Specification content is authored inside this framing Work Item.

## Evidence

- The user directed DRCLI-REQ-CLI-001 to proceed into Specification work.
- The user selected the existing DRMCP Specification as the design baseline while retaining DRMCP independently.
- The user required parallelizable Specification authoring followed by integrated review.
- The user added mandatory CLI help and actionable invalid-usage diagnostics before framing closure.
- T01 fixed the downstream design-convergence contract.
- T02 created DRCLI-WORK-CLI-002.
- No formal Investigation was required for framing.

# PRODUCT-TASK-SPEC-019-03: Decide temporary standalone validator boundary

- **id**: PRODUCT-TASK-SPEC-019-03
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-02
- **outputs**:
  - PRODUCT-TASK-SPEC-019-03

## Goal

Decide the corrected product and integration boundary for the temporary semantic Task responsibility validator.

## Work

- Preserve T01 criterion, evidence, aggregation, and failure-boundary decisions.
- Decide whether the immediate validator is standalone or part of DRMCP.
- Decide ownership of future DRMCP integration.
- Fix the corrected canonical target and downstream Investigation scope.

This Task does not perform Investigation, graph coordination, ADR authoring, Specification authoring, implementation, review, correction, or synchronization.

## Done condition

The temporary-tool boundary and future DRMCP-integration disposition are explicit and ready for impact Investigation.

## Verification

- Confirm the decision evaluates one Task record per invocation.
- Confirm the temporary tool owns checklist selection and local-LLM orchestration.
- Confirm W019 does not modify the current DRMCP tool surface.
- Confirm future DRMCP integration remains separately decidable.
- Confirm T01 remains historical Evidence rather than being rewritten.

## Evidence

### Loop state

- Status: `decision_complete`
- Current decision: `none`
- Terminal decisions: `D-001`
- Open decisions: none
- Blocked decisions: none

### Decision inventory

| ID | Topic | Status | Depends on | Decision summary | Reason | Canonical target | ADR route |
|---|---|---|---|---|---|---|---|
| D-001 | Temporary standalone tool boundary | `decided` | PRODUCT-TASK-SPEC-019-01 | Build the immediate validator as a temporary standalone tool. One invocation evaluates one Task record. The tool reads that Task, selects the predefined common and `task_type`-specific checklist, attaches the checklist to the local-LLM request, receives criterion-level binary judgments and reasons, and derives the final compliance flag. Do not integrate the tool with DRMCP or modify current DRMCP contracts in W019. Route any future DRMCP integration through a separate Requirement or Work Item. | The immediate need is a lightweight interim validator. DRMCP integration introduces a separate tool-surface, diagnostics, architecture, and lifecycle design that is not required for the temporary solution. | PRODUCT-owned temporary semantic Task validator Specification; downstream standalone implementation boundary | `candidate` |

### Preserved T01 decisions

- The validator evaluates exactly one Task record per invocation.
- The Task record is the only semantic evidence source.
- Checklist selection uses the declared `task_type`.
- The applied checklist combines common and type-specific criteria.
- The local LLM returns one binary result and concise reason per criterion.
- The tool derives the final compliance flag through logical AND.
- Structural precondition failure and execution failure remain separate from semantic non-compliance.
- The validator remains read-only and produces no automatic correction or Task-splitting proposal.

### Corrected boundary

- `MCP` in the temporary product shape means an independent interim tool surface, not Design Records MCP integration.
- Current `validate_records` behavior is not extended or reused by this decision.
- Current DRMCP Specifications, diagnostics, source, and tests are not impact targets for W019.
- Future integration may reuse the accepted semantic contract, but that route requires separate design authority.

### Decision closure

- D-001 is `decided`.
- No decision remains `open`, `in_discussion`, or `blocked`.
- Mandatory impact and conflict Investigation is the next responsibility.

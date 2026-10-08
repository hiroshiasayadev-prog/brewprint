# DRCLI-TASK-CLI-001-01: Frame DRCLI Specification convergence

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-001
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - DRCLI-WORK-CLI-001
  - DRCLI-TASK-CLI-001-02

## Goal

Fix the downstream contract for DRCLI Specification convergence.

## Work

### Decision ledger

| ID | topic | status | decision summary | reason | canonical target | ADR route |
|---|---|---|---|---|---|---|
| D-001 | Desired Outcome | `decided` | Produce a complete current DRCLI Specification before implementation. | The user explicitly selected Specification convergence as the next phase. | DRCLI downstream Work Item | `not_required` |
| D-002 | Outcome alignment | `decided` | The Desired Outcome matches DRCLI-REQ-CLI-001. | The Requirement already captures portable CLI, read, validation, guidance, authoring, persistence, Windows/Linux, and CLI help. | DRCLI-REQ-CLI-001 | `not_required` |
| D-003 | Source disposition | `decided` | `proceed`. | The Requirement is accepted and sufficiently bounded for downstream design. | DRCLI-WORK-CLI-002 | `not_required` |
| D-004 | Baseline | `decided` | Use current and old DRMCP Specifications as design evidence and adapt them to DRCLI requirements. | DRMCP already contains mature record, operation, validation, guidance, and authoring contracts. | DRCLI Specification | `not_required` |
| D-005 | DRMCP relation | `decided` | DRCLI does not replace or supersede DRMCP. | The user explicitly retained DRMCP as an independent reference. | DRCLI boundary | `not_required` |
| D-006 | Downstream Goal | `decided` | Author the complete normative DRCLI Specification and pass one integrated independent review. | This is one coherent design completion boundary. | DRCLI-WORK-CLI-002 | `not_required` |
| D-007 | Downstream Boundary | `decided` | Specification only; no production implementation and no DRMCP edits. | Implementation follows design closure. | DRCLI-WORK-CLI-002 | `not_required` |
| D-008 | Completion Condition | `decided` | All required DRCLI spec areas are canonical, cross-consistent, and integrated-review PASS. | Parallel drafts alone are not design closure. | DRCLI-WORK-CLI-002 | `not_required` |
| D-009 | Unknown handling | `decided` | Resolve bounded details from existing DRMCP/Product authority inside downstream design; stop and return to decision only for materially new choices. | No formal framing Investigation is needed. | DRCLI-WORK-CLI-002 | `not_required` |
| D-010 | Initial route | `decided` | Design convergence with parallel writer-disjoint Specification authoring, integration, then independent review. | The requested design naturally partitions by Specification area. | DRCLI-WORK-CLI-002 | `not_required` |
| D-011 | CLI discoverability | `decided` | Help and invalid-usage guidance are mandatory normative behavior, not optional polish. | CLI callers do not receive MCP schemas. | DRCLI command-interface Specification | `not_required` |

### Current cursor

- Decision: none
- Loop state: `decision_complete`

## Done condition

- Every decision item is terminal.
- The proceed contract is complete.
- The downstream design-convergence route is fixed.
- No Specification content or implementation is authored in this Task.

## Verification

- Confirm DRCLI-REQ-CLI-001 contains the added CLI help and usage-diagnostic requirement.
- Confirm no decision requires a formal Investigation before downstream design.
- Confirm DRMCP remains outside the modification boundary.

## Evidence

- The user's current instructions directly decide D-001 through D-011.
- DRCLI-REQ-CLI-001 was amended with mandatory help and actionable invalid-usage diagnostics before this Task completed.
- Filesystem authoring is used because current DRMCP authoring transactions are non-operational.

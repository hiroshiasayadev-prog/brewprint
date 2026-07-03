# TRV-TASK-SPEC-003-02: Route TRV contract ADRs

- **id**: TRV-TASK-SPEC-003-02
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-003
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-003-01
- **outputs**:
  - TRV-TASK-SPEC-003-02

## Goal

Produce the ADR-routing ledger for the accepted TRV application-contract decisions.

## Work

- Route D-002, D-003, D-009, D-010, and D-014 from TRV-TASK-SPEC-001-02.
- Preserve the closed W002 architecture.
- Classify each decision as `required`, `covered`, `not_required`, or validly `blocked`.
- Record create, amend, reuse, or supersede disposition.
- Partition coherent boundaries for external MCP, Task-path input, caller workflow ownership, and future DRMCP compatibility.
- Leave concrete schemas, Go types, path mechanics, and library choices to W004.

This Task must not:

- reopen accepted contract choices;
- change W002 architecture;
- author an ADR or Specification;
- change the Task graph;
- perform detailed design, implementation, review, synchronization, stage, or commit work.

## Done condition

- Every contract decision has one terminal ADR route.
- Every required choice belongs to one coherent boundary.
- Covered and not-required results name exact authority or reason.
- T03 can materialize exact writers without new routing judgment.

## Verification

- Confirm no accepted decision changed.
- Confirm PRODUCT-ADR-SPEC-017 coverage is exact.
- Confirm W004-owned detail remains excluded.
- Confirm no canonical artifact body changed.

## Evidence

- TRV-TASK-SPEC-001-02 supplies terminal contract decisions.
- TRV-WORK-SPEC-002 is done.
- TRV-TASK-SPEC-003-01 created this routing owner.

| decision | route | disposition | ADR boundary | canonical target |
|---|---|---|---|---|
| D-002 | `required` | `create` | TRV-ADR-SPEC-003 | `spec:trv.mcp_interface` |
| D-009 | `required` | `create` | TRV-ADR-SPEC-003 | `spec:trv.mcp_interface` |
| D-003 | `required` | `create` | TRV-ADR-SPEC-004 | `spec:trv.task_input` |
| D-010 | `covered` | `reuse` | PRODUCT-ADR-SPEC-017 | `spec:trv.caller_integration` |
| D-014 | `required` | `create` | TRV-ADR-SPEC-005 | `spec:trv.compatibility` |

ADR boundaries:

- `TRV-ADR-SPEC-003`: standalone stdio MCP server, accepted identity, strict path input, and tagged outcome classes.
- `TRV-ADR-SPEC-004`: one repository-root-relative Task path with rejection of absolute or root-escaping input.
- `TRV-ADR-SPEC-005`: semantic compatibility only across future DRMCP integration; transport, path input, configuration, and implementation remain replaceable.
- PRODUCT-ADR-SPEC-017 fully covers caller-owned human acceptance and rejection.
- Concrete JSON Schema encoding, Go types, path-normalization mechanics, symlink behavior, and MCP library selection remain W004-owned.
- Result: `PASS`.

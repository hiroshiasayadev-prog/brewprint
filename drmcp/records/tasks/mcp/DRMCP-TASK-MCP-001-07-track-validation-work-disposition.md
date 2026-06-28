# DRMCP-TASK-MCP-001-07: Track validation-work disposition

- **id**: DRMCP-TASK-MCP-001-07
- **status**: not_started
- **date**: 2026-06-26
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d coordination
- **depends_on**:
  - DRMCP-TASK-MCP-001-01
- **outputs**:
  - DRMCP-WORK-MCP-007
  - PRODUCT-WORK-SPEC-015

## Goal

Accept the coordinated disposition of existing DRMCP validation Work Items and PRODUCT owner pointers.

## Work

- Track `DRMCP-WORK-MCP-007` and `PRODUCT-WORK-SPEC-015` as the exact owner-local Work Items selected by T01.
- Delegate the disposition of `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002` to the selected owner.
- Delegate matching PRODUCT validation-policy owner-pointer changes to the correctly owned PRODUCT Work Item.
- Require one explicit disposition for each existing Work Item: retain, supersede, absorb, or close.
- Track all required child Work Items through review and `done`.
- Record exact Work Item IDs and accepted evidence here.

This Task does not decide validation semantics or edit owner pointers directly.
All detailed decisions and updates belong to the selected Work Items.

## Done condition

- Every affected existing Work Item has an explicit accepted disposition.
- PRODUCT validation-policy owner pointers match the accepted disposition.
- No replacement Work Item creates duplicated authority.
- Every selected child Work Item is `done`.
- Cross-owner review has no blocking or major findings.
- Exact Work Item IDs and evidence pointers are recorded here.

## Verification

- Compare the final dispositions with `spec:product.design_records.spec_format.validation_policy`.
- Confirm reciprocal references and source Requirements for all created Work Items.
- Confirm that this Task contains no direct validation-policy implementation evidence.

## Evidence

Selected child Work Items:

- `DRMCP-WORK-MCP-007`: DRMCP validation Work Item disposition and rebaseline.
- `PRODUCT-WORK-SPEC-015`: PRODUCT validation owner-pointer synchronization.

T07 remains `not_started` until both child Work Items begin and later reach `done`.

# DRMCP-TASK-MCP-009-31: Review T15 verification ownership and execution graph

- **id**: DRMCP-TASK-MCP-009-31
- **status**: not_started
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-30
- **outputs**:
  - DRMCP-TASK-MCP-009-31

## Goal

Independently judge whether the amended T15 graph restores verification ownership without weakening T05 correction acceptance.

Confirm that every released executor can act without hidden ownership judgment.

## Work

### Read

Read only:

- W009;
- T06, T07, T08, T15, T16, T17, and T23;
- T30 through T36;
- `drmcp/src/internal/designrecords/resolve_reference_test.go` only for top-level test declarations;
- `drmcp/src/internal/designrecords/validation_test.go` only for top-level test declarations;
- `drmcp/src/internal/designrecords/authoring_test.go` only for top-level test declarations;
- `drmcp/src/internal/designrecords/authoring_guidance_test.go` only for top-level test declarations.

Do not read production source.
Do not run tests, formatter, generator, or repository-wide search.

### Review

Verify:

- the first T15 result proves a graph defect but does not prove T05 implementation failure;
- T06 and T07 own focused acceptance before T08 integration;
- T08 retains the first accepted full-package PASS;
- T33 through T36 each have one responsibility;
- T33 and T35 command ownership is separated from T15 routing judgment;
- T36 observation is separated from T16 mapping acceptance;
- T15 performs no command execution or source inspection;
- T15 is not routed to Haiku;
- every exact T05, T06, T07, and T08 test name matches its persistent file or Task boundary;
- unmapped and non-test failures remain blocking;
- mixed failures containing only T06, T07, and T08 names can route downstream without deciding T05 closure;
- a T05-mapped failure remains blocking;
- writer ownership is unique;
- dependencies are acyclic;
- T36 uses a complete-T34-report start gate without requiring T34 lifecycle synchronization before T15;
- T17 cannot claim T08 overall readiness while T06 or T07 remains incomplete.

### Prohibited operations

- Do not modify files.
- Do not correct findings.
- Do not release T33 through T36.
- Do not update lifecycle state.
- Do not start T15, T16, or T17.
- Do not stage or commit.

### Output

1. Verdict: `PASS`, `NEEDS REVISION`, or `BLOCKED`.
2. Reviewed records and declaration files.
3. Graph-defect classification.
4. Exact owner-map result.
5. Writer and dependency result.
6. Blocking findings.
7. Major findings.
8. Minor findings.
9. Advisories.
10. T32 readiness.

## Done condition

- The graph receives an independent verdict.
- Every material finding has an ID and severity.
- The exact owner maps are checked against declaration files.
- Verification and mapping responsibilities are confirmed separate.
- No file is changed.

## Verification

- Compare W009, T06-T08, and T15-T17 ownership directly.
- Compare every T30 owner-map entry with the exact top-level test declarations.
- Confirm T33-T36 depend on T32 and T15 depends on all four leaves.
- Confirm T16 and T17 retain independent review and closure ownership.
- Confirm no production or Go test change belongs to T30 authoring.

## Evidence

Independent review pending.

# DRMCP-TASK-MCP-016-17: Coordinate F-MAJ-02 correction route

- **id**: DRMCP-TASK-MCP-016-17
- **status**: done
- **date**: 2026-07-05
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: coordination
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-016-16
- **outputs**:
  - DRMCP-WORK-MCP-016
  - DRMCP-TASK-MCP-016-10
  - DRMCP-TASK-MCP-016-17
  - DRMCP-TASK-MCP-016-18
  - DRMCP-TASK-MCP-016-19

## Goal

Materialize the exact correction and independent finding-closure route for T16 finding F-MAJ-02.

## Work

- Preserve T16 as the historical `NEEDS REVISION` verdict.
- Preserve F-BLK-01 and F-MAJ-01 as independently `CLOSED`.
- Create one correction Task for F-MAJ-02.
- Limit the correction target to `spec:drmcp.design_records_mcp.namespace_scanning`.
- Create one later independent review Task for F-MAJ-02 closure.
- Block T10 until that review independently closes F-MAJ-02.
- Update W016 Task ownership, flow, dependencies, and release route.

## Done condition

- T18 owns only the F-MAJ-02 projection repair.
- T19 owns only independent F-MAJ-02 closure review.
- The dependency order is T16 -> T17 -> T18 -> T19 -> T10.
- T10 cannot consume T16 as final release proof while F-MAJ-02 is open.
- No completed Task is reopened or substantively changed.

## Verification

- Confirm W016 lists T17 through T19.
- Confirm T18 and T19 point to W016.
- Confirm T18 depends on T17.
- Confirm T19 depends on T18.
- Confirm T10 depends on T19.
- Confirm T16 remains `done` with `NEEDS REVISION` and F-MAJ-02.
- Confirm this Task authors no Specification correction and no review verdict.

## Evidence

- The repository-root `prompt_chappy.md` was read first.
- `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational. Design Record access used filesystem fallback.
- T16 independently closed F-BLK-01 and F-MAJ-01.
- T16 recorded F-MAJ-02 as a Major closure blocker in `spec:drmcp.design_records_mcp.namespace_scanning` at `## Runtime snapshot boundary`.
- F-MAJ-02 requires no new design judgment.
- T18 is the correction owner. T19 is the independent closure-review owner.
- T10 is blocked until T19 records F-MAJ-02 as `CLOSED` and no new closure blocker.
- Scoped graph inspection confirmed W016 lists T17 through T19 and T10 depends on T19.
- Scoped whitespace inspection passed. LF-to-CRLF notices were advisory only.
- No scoped file is staged.
- The standalone semantic responsibility validator was not executed because no operational invocation tool is available. No validator PASS was synthesized.
- No stage or commit was performed.

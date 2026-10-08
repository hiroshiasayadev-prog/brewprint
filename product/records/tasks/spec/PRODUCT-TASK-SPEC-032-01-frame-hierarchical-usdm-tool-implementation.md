# PRODUCT-TASK-SPEC-032-01: Frame hierarchical USDM tool implementation

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-032
- **task_type**: decision
- **estimate**: 0.25d
- **depends_on**:
- **outputs**:
  - PRODUCT-TASK-SPEC-032-02

## Goal

Fix the implementation framing for the reviewed hierarchical USDM contract.

## Work

- Confirm design closure and implementation readiness.
- Identify implementation surfaces affected by hierarchical requirement IDs and effective coverage.
- Select the downstream implementation and review boundaries.

## Done condition

The implementation Work Item can be created without unresolved design judgment.

## Verification

PRODUCT-WORK-SPEC-031 is `done`; PRODUCT-TASK-SPEC-031-09 is `PASS`; accepted Specifications define exact behavior.

## Evidence

| item | state | decision |
|---|---|---|
| Desired Outcome | decided | Make standalone USDM tools and the MCP wrapper conform to hierarchical row identity and effective coverage. |
| Outcome alignment | decided | Directly realizes PRODUCT-REQ-SPEC-016 after reviewed design closure. |
| Source disposition | decided | `proceed`. |
| Unknown handling | decided | Implementation-local code discovery and tests; no design Investigation required. |
| Goal | decided | Implement hierarchical parsing, validation, coverage evaluation, scope reporting, and MCP exposure. |
| Boundary | decided | `tools/usdm`, relevant similarity loader/tests, and `tools/usdm-mcp/server.py`; no Design Record semantic changes. |
| Completion Condition | decided | Scoped tests pass, current corpus validates, cross-app coverage remains correct, and independent implementation review passes. |
| Initial route | decided | Implementation followed by independent implementation review. |

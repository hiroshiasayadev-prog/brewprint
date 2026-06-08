# V01-WORK-MCP-016: Validate required narrative sections for closed workflow artifacts

- **id**: V01-WORK-MCP-016
- **status**: done
- **date**: 2026-06-03
- **source_requirement**: V01-REQ-MCP-017
- **impact_refs**:
  - SPEC-design-records-mcp-tools
  - SPEC-design-records-mcp-schema
- **tasks**:
  - V01-TASK-MCP-016-01
  - V01-TASK-MCP-016-02
  - V01-TASK-MCP-016-03
  - V01-TASK-MCP-016-04

## Goal
Implement status-gated validation so workflow artifacts in `done` or `accepted` states cannot have required narrative sections empty.

This work item completes `V01-REQ-MCP-017` by defining required narrative section policy for workflow artifact status gates, updating the public MCP validation/diagnostic contract, implementing the validator/authoring guard behavior, and recording runtime verification evidence.

`REQ status: accepted` is treated as an adoption-readiness gate, not a close/completion state.

## Boundary

This work item owns required narrative section non-empty validation for workflow artifacts in Design Records MCP validation and authoring guards.

In scope:

- Required narrative section policy for `WORK status: done`, `TASK status: done`, and `REQ status: accepted`
- Diagnostic contract for empty required narrative sections, including record ID, section name, status, and severity
- Design Records MCP validation behavior
- Authoring guard behavior where applicable
- Regression tests and runtime smoke evidence
- Close synchronization for `V01-REQ-MCP-017` and this work item

Out of scope:

- Automated judgment of narrative section quality or sufficiency
- Markdown formatter introduction
- New status transition workflow/state machine
- Non-workflow design record content quality checks

## Impact Scope
Expected impacted areas:

- `docs/spec/design-records-mcp/tools.md`
- `docs/spec/design-records-mcp/schema.md` for Design Records MCP diagnostic category/catalog updates
- `internal/designrecords` validation/parser helpers
- `internal/designrecordsmcp` tool-call/runtime behavior and tests
- Authoring guidance only if the existing guide policy is insufficient or ambiguous

Non-target:

- `docs/spec/diagnostics.md` is the brewprint DSL semantic diagnostics catalog and is not the primary target for this Design Records MCP validation diagnostic.

## Task flow

text flow:

1. Define policy from existing authoring guides and current workflow artifact formats.
2. Update spec/diagnostic contract for the validation behavior.
3. Implement validation and regression tests.
4. Run runtime smoke and close synchronization.

## Task Candidates

- Policy inventory and close-state section matrix
- Spec/diagnostic contract update
- Implementation and regression tests
- Runtime smoke and close synchronization

## Completion Condition

This work item can be closed when:

- Required narrative section policy is explicit for `WORK done`, `TASK done`, and `REQ accepted`.
- Validation diagnostics are emitted for empty required narrative sections in close states.
- Regression tests cover at least `WORK done` with empty `Goal` / `Boundary` / `Evidence`, `TASK done`, and `REQ accepted` policy cases.
- Runtime smoke confirms the MCP validation path reports the new diagnostic.
- `V01-REQ-MCP-017` is updated with accepted close evidence and this work item reference.

## Evidence

V01-WORK-MCP-016 is closed on 2026-06-03 because all owned tasks completed and runtime validation evidence passed.

Completed task evidence:

- `V01-TASK-MCP-016-01` defined the status-gated required narrative section policy matrix for WORK/TASK/REQ.
- `V01-TASK-MCP-016-02` updated `SPEC-design-records-mcp-tools` and `SPEC-design-records-mcp-schema` with the diagnostic contract, policy matrix, and non-empty definition.
- `V01-TASK-MCP-016-03` implemented `missing_required_section` and `empty_required_section` diagnostics, added required `section` / `status` fields, and added regression tests.
- `V01-TASK-MCP-016-04` completed targeted tests, full tests, newline-delimited JSON-RPC runtime smoke, and close synchronization.

Verification evidence:

- `go test ./internal/designrecords ./internal/designrecordsmcp` passed.
- `go test ./...` passed.
- Runtime smoke through `go run ./cmd/design-records-mcp --root .` passed using newline-delimited JSON-RPC requests.
- `validate_records` returned `ok:true` / `diagnostics:null` for `V01-TASK-MCP-016-01..V01-TASK-MCP-016-04`, `V01-WORK-MCP-016`, and `V01-REQ-MCP-017`.
- MCP stderr was empty and process exit code was `0`.
- Final smoke output was `runtime smoke PASS`.

Implemented behavior:

- `WORK status: done` requires non-empty `Goal`, `Boundary`, and `Evidence`.
- `TASK status: done` requires non-empty `Goal`, `Work`, `Done condition`, `Verification`, and `Evidence`.
- `REQ status: accepted` requires non-empty `Requirement` and `Required Outcome` as an adoption-readiness gate.
- `REQ accepted` does not require `Evidence`, `Boundary`, or `Explicitly Excluded Scope`.
- Required section diagnostics use `missing_required_section` and `empty_required_section` with `error` severity and include `section` and `status`.

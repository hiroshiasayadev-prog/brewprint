# V01-TASK-DATA-013-01: Decide contract boundary for request option and response behavior constraints

- **id**: V01-TASK-DATA-013-01
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-DATA-013
- **source_requirement**: V01-REQ-DATA-006
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - Boundary decision for N-011, N-017, N-022, N-024, N-025, and N-028
  - Classification of each candidate into DATA DSL schema, MCP tool behavior contract, diagnostic/error contract, YAML migration, or no-action follow-up
  - Next task split recommendation for V01-WORK-DATA-013

## Goal

Decide the contract boundary for the V01-WORK-DATA-013 numeric / default behavior bucket before any spec, YAML, fixture, or implementation update begins.

This task covers N-011, N-017, N-022, N-024, N-025, and N-028 from the UC-002 notes retreat inventory.

## Work

- Review the six V01-WORK-DATA-013 candidates:
  - N-011: `get_reference_tree_request.depth` numeric range `0..4` and invalid-depth behavior.
  - N-017: `get_references_request.direction` optional enum and default `out`.
  - N-022: `get_source_request.fallback` optional enum and omitted-value fallback behavior.
  - N-024: `get_source_response.fallback` response fallback marker contract.
  - N-025: `inspect_request.detail` optional enum, default `normal`, and unknown-value behavior.
  - N-028: `list_endpoints_request.api_table_id` omitted-input cross-response grouping behavior.
- Decide for each candidate whether the owner should be DATA DSL schema, MCP tool behavior contract, diagnostic/error contract, UC-002 YAML migration, or no-action follow-up.
- Identify which spec files and diagnostic/error surfaces need updates after the boundary decision.
- Decide whether an ADR is required, or whether task evidence plus spec updates are enough.

## Done condition

- Each of N-011, N-017, N-022, N-024, N-025, and N-028 has an explicit ownership decision.
- The decision separates data shape constraints from tool runtime behavior instead of mixing them into one bucket.
- Follow-up work is split into concrete next tasks or explicitly closed as no-action.
- V01-WORK-DATA-013 remains limited to V01-REQ-DATA-006 and does not reopen unrelated DATA or MCP identity work.

## Verification

- Compare the decision against V01-REQ-DATA-006, V01-WORK-DATA-013, V01-INV-DATA-002, V01-TASK-DATA-009-03, and V01-TASK-DATA-009-04.
- Check that selected spec / diagnostic / YAML follow-up surfaces match the ownership decision.
- Confirm that excluded scopes from V01-WORK-DATA-013 remain excluded.

## Evidence

Decision result:

- V01-WORK-DATA-013 will not introduce a DATA DSL extension for request-option defaults, omitted-input behavior, unknown-value behavior, fallback behavior, or cross-response grouping at this stage.
- The main contract owner is the UC-002 MCP tool contract: `docs/spec/mcp/tools/*.md` and related MCP error documentation.
- Numeric range for N-011 (`depth: 0..4`) is also kept as MCP tool behavior/spec text for now instead of adding a generic DATA DSL numeric constraint. This avoids expanding parser / validator / renderer / fixture scope before real operation proves the need.
- The decision is intentionally operational-first: run with explicit MCP tool behavior specs, then revisit DATA DSL expressiveness only if repeated real usage shows the abstraction is worth carrying.

Candidate ownership:

| ID | Decision |
|---|---|
| N-011 | Specify `get_reference_tree.depth` range and invalid-depth behavior in MCP tool spec / errors; no DATA DSL range extension now. |
| N-017 | Specify `get_references.direction` omitted/default behavior in MCP tool spec. |
| N-022 | Specify `get_source.fallback` omitted/default behavior and `file` / `error` branch behavior in MCP tool spec. |
| N-024 | Specify `get_source` response fallback marker in MCP tool response contract. |
| N-025 | Specify `inspect.detail` omitted/default and unknown-value behavior in MCP tool spec / errors. |
| N-028 | Specify `list_endpoints.api_table_id` omitted-input grouped response behavior in MCP tool spec. |

Follow-up recommendation:

- Next task should update the relevant MCP tool specs and error documentation.
- YAML note cleanup and fixture/golden updates should be separate follow-up work after spec wording is accepted.
- Go implementation changes should only be introduced if the current implementation conflicts with the accepted MCP tool contract.

Verification:

- Compared against V01-REQ-DATA-006 / V01-WORK-DATA-013 scope: the decision covers numeric ranges, defaults, omitted values, unknown values, fallback behavior, and cross-response grouping without mixing in excluded DATA/MCP identity scopes.
- No ADR is required for this decision at this time because the chosen direction avoids a new cross-cutting DSL capability and keeps the contract in existing MCP tool spec surfaces.

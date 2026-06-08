# V01-TASK-DATA-013-02: Update MCP tool specs for request option and response behavior contracts

- **id**: V01-TASK-DATA-013-02
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-DATA-013
- **source_requirement**: V01-REQ-DATA-006
- **estimate**: 1d-2d
- **depends_on**:
  - V01-TASK-DATA-013-01
- **outputs**:
  - Updated MCP tool specs for get-reference-tree, get-references, get-source, inspect, and list-endpoints request option / response behavior contracts
  - Updated MCP error documentation if invalid-depth, unknown-detail, fallback, or related behavior errors need explicit documentation
  - Review notes for whether UC-002 YAML note cleanup and implementation verification should be split into later tasks

## Goal

Update the MCP tool contract specs for the request option and response behavior decisions made in `V01-TASK-DATA-013-01`.

This task records the accepted behavior in spec surfaces rather than introducing new DATA DSL constraint syntax.

## Work

- Update the relevant MCP tool specs so the following behavior is explicit:
  - `get_reference_tree`: `depth` range `0..4` and invalid-depth behavior.
  - `get_references`: omitted `direction` defaults to `out`.
  - `get_source`: omitted `fallback` behavior and `file` / `error` branch behavior.
  - `get_source`: response `fallback` marker contract.
  - `inspect`: omitted `detail` defaults to `normal`, and unknown detail values are errors.
  - `list_endpoints`: omitted `api_table_id` response grouping behavior using `tables[]`.
- Update MCP error documentation when the behavior requires named error codes or existing error-code clarification.
- Keep this task limited to specs and error documentation unless a tiny wording-only cross-reference is needed.
- Record any follow-up needed for UC-002 YAML note cleanup, fixture/golden regeneration, or Go implementation verification.

## Done condition

- Relevant MCP tool specs explicitly describe the default, omitted-input, unknown-value, fallback, numeric range, and response grouping behavior selected by `V01-TASK-DATA-013-01`.
- Error documentation is updated or explicitly judged unchanged.
- No DATA DSL numeric range/default/fallback syntax is introduced by this task.
- Follow-up tasks are identified for YAML cleanup, fixtures, or implementation only if needed.

## Verification

- Review updated spec text against `V01-TASK-DATA-013-01` Evidence.
- Check that updates do not reopen excluded V01-WORK-DATA-013 scopes such as MCP identity, selector support matrix, tagged unions, recursive structures, or DAG TypeRef hints.
- Validate design records after proposal acceptance.

## Evidence

Verdict: PASS for spec/error documentation reflection.

Files updated:

- `docs/spec/mcp/tools/get-reference-tree.md`
  - Updated `last_updated` to `2026-06-07`.
  - Clarified that `depth < 0` or `depth > 4` returns `invalid_depth` tool error.
  - Clarified that the `depth` range is an MCP tool runtime constraint, not a DATA DSL numeric range extension.
  - Clarified unknown `direction` values return `unsupported_direction`.
- `docs/spec/mcp/tools/get-references.md`
  - Updated `last_updated` to `2026-06-07`.
  - Clarified omitted `direction` uses `out` and response `direction` returns the actual `out` value.
  - Clarified unknown `direction` values return `unsupported_direction`.
  - Clarified the default is MCP tool behavior, not DATA DSL default syntax.
- `docs/spec/mcp/tools/get-source.md`
  - Updated `last_updated` to `2026-06-07`.
  - Clarified omitted / `file` fallback returns whole-file fallback with `fallback: "file"` and `source_range_unavailable` warning when object range is unavailable.
  - Clarified `fallback=error` returns `source_range_unavailable` tool error instead of a fallback response.
  - Clarified invalid fallback values return `invalid_args`.
- `docs/spec/mcp/tools/inspect.md`
  - Updated `last_updated` to `2026-06-07`.
  - Clarified omitted `detail` uses `normal`.
  - Clarified unknown detail values return `unsupported_detail`.
  - Clarified the default is MCP tool behavior, not DATA DSL default syntax.
- `docs/spec/mcp/tools/list-endpoints.md`
  - Updated `last_updated` to `2026-06-07`.
  - Clarified `api_table_id` specified vs omitted behavior.
  - Clarified omitted `api_table_id` returns all API Tables grouped by `tables[]`.
  - Clarified response grouping is MCP tool behavior, not DATA DSL default / union syntax.
- `docs/spec/mcp/errors.md`
  - Updated `last_updated` to `2026-06-07`.
  - Added `source_range_unavailable` as a tool error for `get_source(fallback=error)` when object-level source range is unavailable.
  - Added request option behavior guidance for `unsupported_direction`, `unsupported_detail`, `invalid_depth`, `invalid_args`, and `source_range_unavailable`.

No DATA DSL numeric range/default/fallback syntax was introduced.

Follow-up needed:

- UC-002 YAML note cleanup should be split into a later task if the YAML notes should point to the newly clarified MCP tool specs.
- Implementation/runtime verification should be split into a later task if current Go behavior must be checked against these contracts.
- Fixture/golden regeneration should only be done if YAML cleanup or implementation changes affect rendered outputs.

Verification performed:

- Checked the updated text against `V01-TASK-DATA-013-01` Evidence.
- Confirmed the changes stay within V01-WORK-DATA-013 scope and do not reopen MCP identity, selector matrix, tagged union, recursive structure, or DAG TypeRef hint work.

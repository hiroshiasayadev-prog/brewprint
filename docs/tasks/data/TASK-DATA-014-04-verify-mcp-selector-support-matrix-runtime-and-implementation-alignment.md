# TASK-DATA-014-04: Verify MCP selector support matrix runtime and implementation alignment

- **id**: TASK-DATA-014-04
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-DATA-014
- **source_requirement**: REQ-DATA-007
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-014-03
- **outputs**:
  - Runtime and implementation alignment report for MCP selector support matrix
  - Follow-up decision for implementation fixes or WORK-DATA-014 close

## Goal

Verify whether the current MCP runtime / Go implementation aligns with the selector support matrix and object-dependent vocabulary accepted by `TASK-DATA-014-02` and reflected into UC-002 YAML notes by `TASK-DATA-014-03`.

This task is an investigation / verification task. It should not change implementation unless a separate follow-up task is created.

## Work

- Review the accepted selector support matrix and object-dependent vocabulary in `docs/spec/mcp/schema.md`.
- Review tool-specific selector behavior in:
  - `docs/spec/mcp/tools/get-signature.md`
  - `docs/spec/mcp/tools/list-objects.md`
  - `docs/spec/mcp/tools/get-references.md`
  - `docs/spec/mcp/tools/get-reference-tree.md`
  - `docs/spec/mcp/tools/analyze-impact.md`
  - `docs/spec/mcp/tools/inspect.md`
  - `docs/spec/mcp/errors.md`
- Inspect the Go implementation for the corresponding MCP query/runtime behavior.
- Run targeted runtime smoke tests or unit tests where practical.
- Record whether implementation matches the accepted contract, or whether follow-up implementation tasks are required.

## Done condition

- Runtime / implementation behavior is checked for the selector-related tools in scope.
- Supported selector behavior, unsupported selector behavior, `analyze_impact` exception behavior, `field` ObjectRef behavior, and `list_objects` object/kind filter behavior are explicitly covered.
- Evidence states either:
  - implementation matches the accepted spec and `WORK-DATA-014` can proceed to close, or
  - specific follow-up implementation/spec tasks are required.
- This task does not silently mix in implementation changes. If code changes are required, create or recommend a separate task.

## Verification

- Run targeted tests / runtime smoke checks relevant to the selector support matrix.
- Run Design Records MCP validation for this task and `WORK-DATA-014` after evidence is updated.
- If no Go tests are run, explain why in Evidence.

## Evidence

Verdict: NEEDS FOLLOW-UP.

Investigation scope:

- Reviewed `REQ-DATA-007`, `WORK-DATA-014`, and `TASK-DATA-014-01` through `TASK-DATA-014-04`.
- Reviewed canonical selector matrix / vocabulary specs:
  - `docs/spec/mcp/schema.md`
  - `docs/spec/mcp/tools/get-signature.md`
  - `docs/spec/mcp/tools/list-objects.md`
  - `docs/spec/mcp/tools/get-references.md`
  - `docs/spec/mcp/tools/get-reference-tree.md`
  - `docs/spec/mcp/tools/analyze-impact.md`
  - `docs/spec/mcp/tools/inspect.md`
  - `docs/spec/mcp/errors.md`
- Inspected implementation areas:
  - `internal/query/*`
  - `internal/mcp/*`
  - `internal/designrecords/*`
  - `internal/designrecordsmcp/*`
  - `internal/resolve/references.go`
  - `internal/semantic/reference.go`

Runtime / test results:

- `go test ./internal/designrecords ./internal/designrecordsmcp`: PASS.
- `go test ./internal/query ./internal/mcp`: PASS.
- Temporary MCP smoke was created and deleted by the investigator. No investigation changes were intentionally left.
- Key smoke outputs:
  - `get_signature(file: state_file)` returned `not_found`.
  - `inspect(primitive)` returned `not_found`.
  - `field kind=model_field` full ID form returned normal response.
  - `field kind=model_field` with `id + local_id` returned `unsupported_object`.
  - `list_objects unknown_kind` returned normal empty response.
  - `analyze_impact(view: sequence_diagram)` returned normal response with `unsupported_selector` diagnostic.

Findings:

| severity | area | finding |
|---|---|---|
| major | unsupported selector handling | `get_signature`, `inspect`, and MCP server handling do not consistently classify unsupported-but-resolvable selectors as `unsupported_object`. Smoke showed `get_signature(file: state_file)` and `inspect(primitive)` returned `not_found`, while spec expects `unsupported_object` for matrix `no` cells in these tools. |
| major | file selector references | `get_references(file: node)` / `file: state_file` limited behavior is not aligned with the accepted contract. File selectors are mapped to `StateFileObjectKey(fileID)`, so node-file references are not aggregated by nodes in that file, and state-file results do not include transition-owned `transition_from` / `transition_event` / `transition_to` / `transition_action` references. |
| major | `model_field` alias | `kind: model_field` field selector behavior is inconsistent. Full ID works because it resolves before kind validation, but `id + local_id` returns `unsupported_object`. The shared schema allows `field` or `model_field` as field kind vocabulary. |
| major | `list_objects` validation | `list_objects` does not validate unknown `object` / `kind`; it silently returns empty results. The accepted spec says unknown object or kind is `invalid_args`. |
| minor | `list_objects(kind: model_field)` | `kind: model_field` returns empty because listed field refs use `kind: field`. This may be acceptable if `model_field` is only explanatory selector/response vocabulary, but it is currently ambiguous against the object-omitted kind filter text. |

Selector matrix coverage:

| topic | result |
|---|---|
| supported selector behavior | partial |
| unsupported selector behavior | partial |
| `analyze_impact` exception behavior | confirmed |
| field ObjectRef object value | partial; output normalizes to `object: field`, but selector behavior is inconsistent |
| `list_objects` object/kind filter behavior | partial |
| limited file selector behavior | not confirmed / likely mismatched |

Conclusion:

- Do not close `WORK-DATA-014` yet.
- Create a follow-up implementation task to align MCP runtime behavior with the accepted selector support matrix.
- The follow-up should cover selector support classification, proper tool error codes, `model_field` alias handling, `list_objects` filter validation, and `file: node` / `file: state_file` reference aggregation.

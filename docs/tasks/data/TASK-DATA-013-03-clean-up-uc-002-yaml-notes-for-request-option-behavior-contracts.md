# TASK-DATA-013-03: Clean up UC-002 YAML notes for request option behavior contracts

- **id**: TASK-DATA-013-03
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-DATA-013
- **source_requirement**: REQ-DATA-006
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-013-02
- **outputs**:
  - UC-002 YAML notes for N-011, N-017, N-022, N-024, N-025, and N-028 aligned with updated MCP tool specs
  - Review result for whether fixture or golden regeneration is required after YAML note cleanup

## Goal

Clean up the UC-002 YAML notes for the WORK-DATA-013 request option and response behavior candidates after the MCP tool specs were clarified in `TASK-DATA-013-02`.

This task should keep the YAML notes as lightweight pointers or concise summaries, not as the primary source of behavior truth.

## Work

- Review the UC-002 YAML model files for the six WORK-DATA-013 candidates:
  - N-011: `get_reference_tree_request.depth`.
  - N-017: `get_references_request.direction`.
  - N-022: `get_source_request.fallback`.
  - N-024: `get_source_response.fallback`.
  - N-025: `inspect_request.detail`.
  - N-028: `list_endpoints_request.api_table_id`.
- Update stale or overly detailed notes so they align with the clarified MCP tool specs.
- Avoid introducing DATA DSL numeric range/default/fallback syntax.
- Keep the authoritative behavior contract in `docs/spec/mcp/tools/*.md` and `docs/spec/mcp/errors.md`.
- Determine whether YAML-only note changes affect rendered fixtures or golden outputs.

## Done condition

- The relevant UC-002 YAML notes no longer conflict with the MCP tool specs updated by `TASK-DATA-013-02`.
- YAML remains descriptive and does not become the authoritative behavior contract.
- No DATA DSL schema capability is introduced by this task.
- Fixture/golden impact is checked and either verified unchanged or recorded as a follow-up.

## Verification

- Compare updated YAML notes against:
  - `docs/spec/mcp/tools/get-reference-tree.md`
  - `docs/spec/mcp/tools/get-references.md`
  - `docs/spec/mcp/tools/get-source.md`
  - `docs/spec/mcp/tools/inspect.md`
  - `docs/spec/mcp/tools/list-endpoints.md`
  - `docs/spec/mcp/errors.md`
- Run relevant validation/render commands if YAML note changes may affect outputs.
- Report whether any fixture/golden files changed.

## Evidence
Verdict: PASS.

Files updated:

- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/inspect_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_endpoints_request.yaml`
- `docs/uc/002-brewprint-self-hosting/renders/` regenerated after YAML note cleanup.

YAML note cleanup result:

- The six candidate notes now point to MCP tool behavior as the authoritative contract for range/default/fallback/unknown-value/grouping behavior.
- The YAML remains descriptive and does not introduce DATA DSL numeric range, default, fallback, or grouping syntax.
- Stale wording that implied behavior constraints were held as DATA model expressiveness gaps was replaced with MCP tool spec pointer wording.

Verification performed:

- Compared the YAML notes against the updated MCP specs from `TASK-DATA-013-02`:
  - `docs/spec/mcp/tools/get-reference-tree.md`
  - `docs/spec/mcp/tools/get-references.md`
  - `docs/spec/mcp/tools/get-source.md`
  - `docs/spec/mcp/tools/inspect.md`
  - `docs/spec/mcp/tools/list-endpoints.md`
  - `docs/spec/mcp/errors.md`
- `go run ./cmd/brewprint validate --yaml-root docs/uc/002-brewprint-self-hosting/yaml` -> PASS (`ok`).
- `go run ./cmd/brewprint render --yaml-root docs/uc/002-brewprint-self-hosting/yaml --out docs/uc/002-brewprint-self-hosting/renders --clean` -> PASS (`rendered 46 file(s)`).
- `go test ./...` -> PASS.

Fixture/golden impact:

- UC-002 render regeneration was needed after the YAML note edits and was performed.
- `docs/uc/002-brewprint-self-hosting/renders/` changed: 23 tracked files modified and 5 generated files newly present after render.
- The target model render files reflect the YAML note cleanup. Broader regenerated render changes also reflect the repository's current UC-002 YAML state; no DATA DSL syntax was introduced by this task.

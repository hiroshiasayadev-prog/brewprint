# TASK-DATA-006-03: UC-002 helper migration verification review

- **id**: TASK-DATA-006-03
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-DATA-006
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.25d
- **depends_on**:
  - TASK-DATA-006-02
- **outputs**:
  - Verification review result for TASK-DATA-006-02
  - Temp render / canonical render file list and SHA-256 comparison evidence
  - Validation and focused test evidence

## Goal

Verify that TASK-DATA-006-02's UC-002 YAML migration and render updates stayed within the 7 candidates selected by TASK-DATA-006-01.

Do not perform additional YAML migration or canonical render regeneration unless required to resolve a verification finding.

## Review Result

### Verdict

OK to proceed.

### Files Reviewed

- `docs/tasks/data/TASK-DATA-006-01-uc-002-helper-shape-migration-set.md`
- `docs/tasks/data/TASK-DATA-006-02-uc-002-model-file-helper-yaml-migration.md`
- `docs/work-items/data/WORK-DATA-006-helper-shape-migration.md`
- `docs/adr/070-model-visibility-file-private-helper-model.md`
- `docs/adr/071-file-private-helper-model-render-exposure.md`
- `docs/adr/075-model-file-render.md`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_endpoints_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/analyze_impact.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_reference_tree.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_references.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_signature.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_source.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/inspect.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/list_endpoints.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/list_objects.yaml`
- `docs/uc/002-brewprint-self-hosting/renders/index.md`
- `docs/uc/002-brewprint-self-hosting/renders/mcp/index.md`
- `docs/uc/002-brewprint-self-hosting/renders/mcp/model-analyze_impact_response.md`
- `docs/uc/002-brewprint-self-hosting/renders/mcp/model-get_reference_tree_response.md`
- `docs/uc/002-brewprint-self-hosting/renders/mcp/model-get_source_response.md`
- `docs/uc/002-brewprint-self-hosting/renders/mcp/model-list_endpoints_response.md`
- `docs/uc/002-brewprint-self-hosting/renders/mcp/model-list_objects_response.md`

### Findings

No blocking, major, minor, or nit findings.

### Non-issues

- The selected 7 candidates are represented as same-file private helper model migrations under the owning response model files.
- The owning response models remain public via `main: true`.
- The selected helper models were not cut out into standalone public model YAML files or standalone model render files.
- N-029 legitimately creates three private helper models (`list_endpoints_table`, `list_endpoints_section`, `list_endpoints_endpoint`) because the selected candidate includes nested table / section / endpoint response shapes.
- The 8 UC-002 MCP task files still preserve the `query_service.returns.model: any` to `build_response.params[].model: any` `query_result` pattern.
- Tagged union / discriminator payloads, identity semantics, optional semantics, literal / enum constraints, and other notes-retreat debt remain note-based and outside this task.
- The current renderer produced 40 UC-002 render files; the count is consistent with 29 model renders plus the existing 8 DAG renders, group index, project index, and cross ER output.
- Temp render and canonical render file lists matched.
- Temp render and canonical render SHA-256 hashes matched for all 40 files.
- Pre-existing unrelated dirty / untracked files were not reverted, staged, or committed.

### Suggested Next Step

Proceed to the WORK-DATA-006 close / follow-up split task when ready.

## Evidence

Completed on 2026-06-01.

Verification commands:

- `go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml` -> `ok`
- `go test ./internal/resolve ./internal/render/model ./internal/render/project` -> pass
- `validate_records(kind="task")` -> pass before creating this task
- `validate_records(kind="work_item")` -> pass before updating this work item
- `go run ./cmd/brewprint render --yaml-root docs\uc\002-brewprint-self-hosting\yaml --out $env:TEMP\brewprint-uc002-data006-03-render --clean` -> `rendered 40 file(s)`
- Temp render file count: 40
- Canonical render file count: 40
- Temp render and canonical render relative file lists: matched
- Temp render and canonical render SHA-256 hashes: matched

Additional verification checks:

- Static text check confirmed the selected response fields now reference same-file helper model types:
  - `analyze_impact_response.impacts -> list<analyze_impact_impact>`
  - `analyze_impact_response.coverage -> analyze_impact_coverage`
  - `get_reference_tree_response.nodes -> list<get_reference_tree_node>`
  - `get_reference_tree_response.edges -> list<get_reference_tree_edge>`
  - `get_source_response.snippet -> get_source_snippet`
  - `list_endpoints_response.tables -> list<list_endpoints_table>`
  - `list_objects_response.objects -> list<list_objects_object>`
- Static text check confirmed exactly one `main: true` public model in each selected response model file.
- Static file check confirmed no standalone YAML or render files exist for the private helper IDs.
- Static task-file check confirmed each of the 8 UC-002 MCP task files still has two `query_result` occurrences and two `model: any` occurrences.
- Render review confirmed `## Private models` sections expose the selected helper shapes in the owning response model renders.

No additional canonical render regeneration was performed during this verification review.

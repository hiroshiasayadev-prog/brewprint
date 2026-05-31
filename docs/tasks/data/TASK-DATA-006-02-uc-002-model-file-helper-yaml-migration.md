# TASK-DATA-006-02: UC-002 model-file helper YAML migration

- **id**: TASK-DATA-006-02
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-DATA-006
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-006-01
- **outputs**:
  - UC-002 selected model response YAML migrated to same-file private helper models
  - Deferred task-file any candidates preserved
  - Verification evidence for YAML parse, render, golden impact, and focused Go tests

## Goal

Migrate the selected UC-002 model-file response-local helper-shape candidates from `any + note` into same-file private helper models, while keeping the owning response models public.

This task does not migrate task-file `query_result:any` patterns, because TASK-DATA-006-01 classified those as deferred under the accepted DATA-004 / REQ-DATA-003 task `params[].model` private-helper policy.

## Work

- Add `main: true` to the selected owning response models so they remain public when same-file private helper models are introduced.
- Convert the selected response-local nested shapes to same-file private helper models.
- Preserve remaining tagged union, enum / literal, optional-field, identity, and other out-of-scope constraints as notes.
- Leave UC-002 MCP task files and task-file `query_result:any` patterns unchanged.
- Regenerate UC-002 render fixtures if the current renderer produces changed output.
- Record verification evidence and update the parent work item metadata.

## Done condition

- The selected 7 model-file response candidates are represented by same-file private helper models.
- No public model files are created for those candidates.
- The owning response models remain public.
- UC-002 MCP task files are not changed for `query_result:any` migration.
- YAML validation, render generation, render fixture comparison, and focused Go tests have passing evidence.

## Verification

- Validate UC-002 YAML with the documented brewprint CLI.
- Render UC-002 to a temp directory and to the canonical `renders/` directory.
- Confirm temp and canonical render outputs have matching file lists and hashes.
- Run focused Go tests for resolver and model/project renderer packages.
- Validate task and work item design records through Design Records MCP.

## Evidence

Completed on 2026-06-01.

Initial repository check:

- `git status --short` showed existing unrelated modified / untracked files before this task, including design-records implementation/spec files and other workflow artifacts. Those unrelated files were preserved.
- `git log --oneline -1` returned `09ba38b docs(data): triage M15 deferred follow-ups`, which differs from the reported pre-work HEAD `593d4b2 feat(data): enforce private helper signature policy`.

Selected candidates migrated:

- N-005: `analyze_impact_response.impacts` -> `list<analyze_impact_impact>`
- N-006: `analyze_impact_response.coverage` -> `analyze_impact_coverage`
- N-014: `get_reference_tree_response.nodes` -> `list<get_reference_tree_node>`
- N-015: `get_reference_tree_response.edges` -> `list<get_reference_tree_edge>`
- N-023: `get_source_response.snippet` -> `get_source_snippet`
- N-029: `list_endpoints_response.tables` -> `list<list_endpoints_table>` with nested `list_endpoints_section` and `list_endpoints_endpoint`
- N-033: `list_objects_response.objects` -> `list<list_objects_object>`

Files changed by this task:

- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_endpoints_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_response.yaml`
- `docs/uc/002-brewprint-self-hosting/renders/index.md`
- `docs/uc/002-brewprint-self-hosting/renders/mcp/index.md`
- `docs/uc/002-brewprint-self-hosting/renders/mcp/model-*.md`
- `docs/tasks/data/TASK-DATA-006-02-uc-002-model-file-helper-yaml-migration.md`
- `docs/work-items/data/WORK-DATA-006-helper-shape-migration.md`

Deferred candidates intentionally unchanged:

- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/analyze_impact.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_reference_tree.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_references.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_signature.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_source.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/inspect.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/list_endpoints.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/list_objects.yaml`

Out-of-scope constraints preserved as notes:

- `analyze_impact_response.change` remains `any` for discriminated object payload.
- `analyze_impact_impact.suggested_fixes` remains `any` for fix kind-dependent payloads.
- severity / fixability / coverage vocabulary constraints remain notes instead of enum migration.
- `get_source_snippet.language` keeps the `yaml` literal constraint as a note.
- `list_endpoints_endpoint` optional fields keep optional semantics as notes.
- `list_objects_object` identity semantics remain ADR-078+ / MCP identity follow-up scope.
- `get_reference_tree_edge` explicitly repeats Reference-like fields plus `depth` instead of introducing inheritance / extension.

Verification commands:

- `go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml` -> `ok`
- `go run ./cmd/brewprint render --yaml-root docs\uc\002-brewprint-self-hosting\yaml --out $env:TEMP\brewprint-uc002-data006-02-render --clean` -> `rendered 40 file(s)`
- `go run ./cmd/brewprint render --yaml-root docs\uc\002-brewprint-self-hosting\yaml --out docs\uc\002-brewprint-self-hosting\renders --clean` -> `rendered 40 file(s)`
- Temp render output and canonical render output file lists matched.
- Temp render output and canonical render output SHA-256 hashes matched.
- `go test ./internal/resolve ./internal/render/model ./internal/render/project` -> pass
- `validate_records(kind="task")` -> pass
- `validate_records(kind="work_item")` -> pass

Render fixture impact:

- The current renderer produced 40 UC-002 render files, including 29 model-file renders and the existing 8 DAG, group index, project index, and cross ER outputs.
- This exposed a stale 11-file canonical render set from before UC-002 model-file render fixtures were generated.
- Canonical UC-002 renders were regenerated from the current renderer output to keep indexes and linked model render files internally consistent.

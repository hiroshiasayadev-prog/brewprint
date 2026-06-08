# V01-TASK-DATA-012-01: Review UC-002 enum, literal, and vocabulary cleanup candidates

- **id**: V01-TASK-DATA-012-01
- **status**: done
- **date**: 2026-06-05
- **work_item**: V01-WORK-DATA-012
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
- **outputs**:
  - Candidate review table for V01-WORK-DATA-012 owned UC-002 enum-like / literal / vocabulary notes
  - Classification of each candidate as named enum, literal constraint, usage-site vocabulary, covered by other work, or no-action
  - Input for the next cleanup path decision task

## Goal

Review the UC-002 enum-like, literal, and usage-site vocabulary cleanup candidates owned by `V01-WORK-DATA-012` and produce a bounded candidate classification before any YAML migration or fixture regeneration.

This task owns only candidate review and boundary classification. It does not modify UC-002 YAML, fixtures, golden outputs, parser, renderer, validator, MCP code, or existing DATA work items.

## Work

- Review the `V01-WORK-DATA-012` candidate set: N-019, N-030, N-034, N-045, N-046, N-051, and residual vocabulary or literal notes from N-006, N-015, N-023, and N-029.
- Use `V01-TASK-DATA-009-03`, `V01-TASK-DATA-009-04`, `V01-INV-DATA-002`, and the relevant UC-002 YAML files as input evidence.
- For each candidate, classify the cleanup path as one of:
  - named enum candidate;
  - literal constraint candidate;
  - usage-site vocabulary / note cleanup candidate;
  - covered by another existing or successor work item;
  - obsolete / no-action.
- Explicitly separate behavior/default constraints, selector support matrix constraints, recursive/untagged-union structure, tagged union scope, and MCP identity scope from this task.
- Record any unclear candidate as a decision input for the next task rather than silently migrating it.

## Done condition

- Every candidate owned by `V01-WORK-DATA-012` has a recorded classification.
- The classification distinguishes YAML-cleanup-ready items from items that need a spec/ADR/requirement decision first.
- Out-of-scope items are explicitly routed to their owner or marked no-action with evidence.
- No UC-002 YAML, fixture, golden, parser, renderer, validator, MCP implementation, or closed DATA work item is changed by this task.
- The task has enough evidence for a follow-up cleanup path decision task.

## Verification

- Confirm the reviewed candidate list matches `V01-WORK-DATA-012` and the `V01-TASK-DATA-009-03` / `V01-TASK-DATA-009-04` successor split.
- Confirm classifications do not reopen M15, `V01-WORK-DATA-001`, `V01-WORK-DATA-006`, `V01-WORK-DATA-009`, or `V01-WORK-DATA-010`.
- Confirm default/behavior candidates remain outside this task and point to `V01-REQ-DATA-006` / `V01-WORK-DATA-013` where applicable.
- Confirm selector/support-matrix candidates remain outside this task and point to `V01-REQ-DATA-007` / `V01-WORK-DATA-014` where applicable.
- Run Design Records MCP validation for the created task and `V01-WORK-DATA-012` after creation.

## Evidence

Completed on 2026-06-05.

### Sources reviewed

- `docs/prompt_chappy.md`
- `AGENTS.md`
- `docs/AGENTS.md`
- `docs/doc-policy.md`
- `docs/tasks/data/TASK-DATA-012-01-review-uc-002-enum-literal-and-vocabulary-cleanup-candidates.md`
- `docs/work-items/data/WORK-DATA-012-uc-002-enum-literal-vocabulary-cleanup.md`
- `docs/tasks/data/TASK-DATA-009-03-decide-remaining-uc-002-notes-retreat-successor-outcomes.md`
- `docs/tasks/data/TASK-DATA-009-04-create-follow-up-split-or-close-remaining-uc-002-notes-retreat-classification.md`
- `docs/investigations/data/INV-DATA-002-uc002-notes-retreat-inventory-and-m15-release-boundary-input.md`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_endpoints_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/mcp_error.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/reference.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/string_list.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/common.yaml`
- `docs/spec/mcp/tools/analyze-impact.md`
- `docs/spec/mcp/tools/get-reference-tree.md`
- `docs/spec/mcp/tools/get-references.md`
- `docs/spec/mcp/tools/get-source.md`
- `docs/spec/mcp/tools/list-endpoints.md`
- `docs/spec/mcp/tools/list-objects.md`
- `docs/spec/mcp/errors.md`
- `docs/spec/mcp/schema.md`
- Design Records MCP records for `V01-ADR-067`, `V01-ADR-070`, `V01-ADR-073`, `V01-REQ-DATA-004`, `V01-WORK-DATA-006`, `V01-WORK-DATA-010`, `V01-REQ-DATA-006`, `V01-WORK-DATA-013`, `V01-REQ-DATA-007`, `V01-WORK-DATA-014`, `V01-REQ-DATA-008`, `V01-WORK-DATA-015`, `V01-REQ-MCP-004`, and `V01-WORK-MCP-004`.

### Candidate review table

| candidate ID | source YAML file | field / model path | current type | note summary | classification | recommended next handling | out-of-scope routing if applicable | confidence | evidence |
|---|---|---|---|---|---|---|---|---|---|
| N-006 residual | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml` | `analyze_impact_response.coverage` -> `analyze_impact_coverage.analyzed` / `not_analyzed` | `analyze_impact_coverage`; child fields are `string_list` | Helper shape is already explicit, but coverage element vocabularies remain in notes. | usage-site vocabulary / note cleanup candidate | Next task should decide whether `coverage.analyzed` and `coverage.not_analyzed` need one or two named vocabulary enums, then replace only those usage sites if selected. | Helper-shape migration is already covered by closed `V01-WORK-DATA-006`. `suggested_fixes` tagged payload remains outside this task and belongs to tagged-union scope if reopened. | high | YAML uses `analyze_impact_coverage` and says coverage vocabulary enum migration was not done; spec `analyze-impact` section 7 lists v1 standard coverage vocabularies. |
| N-015 residual | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml` | `get_reference_tree_response.edges` -> `get_reference_tree_edge.kind` / `direction`; also `get_reference_tree_node.via` | `list<get_reference_tree_edge>`; residual fields are `str` or `string_list` | Edge helper shape is explicit, but `kind`, `direction`, and `via` still depend on Reference vocabularies in notes. | usage-site vocabulary / note cleanup candidate | Do not create a separate edge enum. Reuse the outcome of N-045 for reference kind and N-046 for reference direction if those named enums are selected. | Helper-shape migration is already covered by closed `V01-WORK-DATA-006`. Depth/range/default behavior remains outside this task. | high | YAML defines `get_reference_tree_edge` as a local helper and notes `Reference kind`, `out / in`, and `Reference.kind` path vocabulary; spec says edges are Reference plus `depth`. |
| N-019 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_response.yaml` | `get_references_response.direction` | `str` | Actual direction is `out / in / both`. | named enum candidate | Decide whether to reuse/rename the existing `reference_tree_direction` enum or introduce a more general `reference_query_direction` enum for `out / in / both`. | Request default behavior from the paired request field is outside this task and belongs to `V01-REQ-DATA-006` / `V01-WORK-DATA-013`; do not reopen M15 or `V01-WORK-DATA-001`. | high | YAML still has `type: str`; `common.yaml` already has `reference_tree_direction` with `out / in / both`; spec `get-references` defines the same values for query direction. |
| N-023 residual | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_response.yaml` | `get_source_response.snippet` -> `get_source_snippet.language` | `get_source_snippet`; `language` is `str` | Helper shape is explicit, but `language` has literal value `yaml` only in a note. | literal constraint candidate | Next task should decide whether brewprint model syntax can express a field literal constraint, or whether this remains a note-only contract. Do not treat this as tagged-union discriminator work. | Snippet helper shape is already covered by closed `V01-WORK-DATA-006`. `fallback` default/behavior and response fallback marker are N-024-style behavior scope under `V01-REQ-DATA-006` / `V01-WORK-DATA-013`, not this task. | high | YAML defines `get_source_snippet.language` as `str` with literal `yaml`; spec `get-source` output requires `language: "yaml"`. |
| N-029 residual | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_endpoints_response.yaml` | `list_endpoints_response.tables` -> `list_endpoints_endpoint.method`, `params`, `returns`, `source` | `list<list_endpoints_table>`; endpoint fields are `str` / `source_location` | Nested helper shape is explicit. Current YAML/spec notes do not list a closed HTTP method vocabulary or literal value; remaining notes are mostly optional-field semantics. | obsolete / no-action | Do not create enum/literal cleanup from `HTTP method` wording alone. A future task may reopen only if a spec decision lists a closed method set. | Helper-shape migration is already covered by closed `V01-WORK-DATA-006`. Optional field semantics are not the enum/literal/vocabulary bucket and should not be smuggled into this work item. | medium | YAML defines table/section/endpoint local helpers and only says `method` is `HTTP method`; spec `list-endpoints` section 4 also does not enumerate method values. |
| N-030 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_request.yaml` | `list_objects_request.object` | `str` | Optional object filter values are `node / view / transition / field`; omitted means all object categories in this tool. | named enum candidate | Prefer a list-objects-specific object filter enum unless the next task deliberately decides the broader `mcp_object_type` enum is valid here. | Object-dependent `kind` vocabulary and selector matrix constraints remain outside this task under `V01-REQ-DATA-007` / `V01-WORK-DATA-014`. MCP identity semantics remain under `V01-REQ-MCP-004` / `V01-WORK-MCP-004`. | medium | YAML and spec `list-objects` list the same four object filter values; existing `mcp_object_type` is broader than this field's current spec. |
| N-034 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/mcp_error.yaml` | `mcp_error.code` | `str` | Closed MCP error code vocabulary is listed in the field note. | named enum candidate | Create or select a named `mcp_error_code` enum and use it on `mcp_error.code` if the next cleanup task proceeds. | None identified. | high | YAML and spec `errors` list the same code set, including `project_invalid`, `invalid_args`, `invalid_selector`, `invalid_change_payload`, `not_found`, `kind_mismatch`, `ambiguous`, `unsupported_object`, `unsupported_detail`, `unsupported_direction`, `invalid_depth`, and `internal_error`. |
| N-045 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/reference.yaml` | `reference.kind` | `str` | Closed semantic reference kind vocabulary is listed in the field note. | named enum candidate | Create or select a named `reference_kind` enum. Reuse it at usage sites such as N-015 `edge.kind` and selected `string_list` filters if list element enum cleanup is chosen. | Reference identity semantics, semantic-object IDs, and source-target index identity remain under `V01-REQ-MCP-004` / `V01-WORK-MCP-004`; this row owns only the value vocabulary. | high | YAML lists all `Reference.kind` values; spec `schema` section 2.2 defines the same public Reference kind vocabulary. `V01-TASK-DATA-009-03` explicitly keeps N-045 residue in the enum/literal bucket, not primary MCP identity. |
| N-046 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/reference.yaml` | `reference.direction` | `str` | Reference direction is `out / in`. | named enum candidate | Create or select a named `reference_direction` enum for `out / in`. Reuse it for N-015 `get_reference_tree_edge.direction` if selected. | None identified. Do not conflate this two-value enum with N-019's `out / in / both` query direction without an explicit subset decision. | high | YAML lists `out / in`; spec `schema` section 2.3 defines Reference direction with those two values. |
| N-051 | `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/string_list.yaml` | `string_list.element` | list element is `str` | Shared `string_list` cannot own one enum because element vocabulary changes by usage site. | usage-site vocabulary / note cleanup candidate | Keep `string_list` as the free-text/shared fallback. For usage sites with a known closed value set, replace the field type with `list<selected_enum>` or a usage-specific list model only after a next-task decision. | Object-dependent vocabularies and selector/support matrices remain under `V01-REQ-DATA-007` / `V01-WORK-DATA-014`. Default/behavior constraints remain under `V01-REQ-DATA-006` / `V01-WORK-DATA-013`. | high | YAML says element enum values are supplemented per usage note; inventory classifies N-051 as not covered by existing ADRs because the shared list has usage-site-dependent vocabularies. |

### Out-of-scope routing notes

- Default, omitted-value, fallback, unknown-value, and other behavior constraints remain outside `V01-WORK-DATA-012` and route to `V01-REQ-DATA-006` / `V01-WORK-DATA-013`. Encountered examples include `get_source_response.fallback` adjacency and request-side direction/default behavior.
- Selector support matrices and object-dependent `kind` vocabularies remain outside `V01-WORK-DATA-012` and route to `V01-REQ-DATA-007` / `V01-WORK-DATA-014`. N-030 should not absorb N-031-style `kind` matrix work.
- Recursive and untagged-union representation remains outside `V01-WORK-DATA-012` and routes to `V01-REQ-DATA-008` / `V01-WORK-DATA-015`.
- Tagged union and discriminator payload support remains outside `V01-WORK-DATA-012` and is already covered by `V01-REQ-DATA-004` / closed `V01-WORK-DATA-010`.
- MCP identity / semantic reference identity remains outside `V01-WORK-DATA-012` and routes to `V01-REQ-MCP-004` / `V01-WORK-MCP-004`. N-045 owns only the public `Reference.kind` value set here.
- Closed work items `V01-WORK-DATA-001`, `V01-WORK-DATA-006`, `V01-WORK-DATA-009`, and `V01-WORK-DATA-010` were treated as closed evidence and were not reopened.

### Next-task input summary

- Cleanup-ready named enum candidates: N-019, N-030, N-034, N-045, and N-046, subject to enum naming/reuse decisions.
- Literal constraint candidate: N-023 residual `get_source_snippet.language = yaml`; this likely needs a model-syntax decision before YAML cleanup.
- Usage-site vocabulary cleanup candidates: N-006 residual coverage vocabularies, N-015 residual reference-kind/direction usage sites, and N-051 shared-list usage sites.
- Explicit no-action candidate: N-029 residual, because current source evidence does not enumerate an endpoint HTTP method value set and the remaining notes are helper-shape or optionality notes.
- Open decision for the next task: whether to generalize existing `reference_tree_direction` or create a new query-direction enum for N-019; whether N-030 should use a narrow list-objects filter enum or the broader `mcp_object_type`; whether literal constraints have a supported representation or remain note-only.

### Verification notes

- Pre-edit Design Records MCP validation for `V01-TASK-DATA-012-01`: passed.
- Pre-edit Design Records MCP validation for `V01-WORK-DATA-012`: passed.
- `go run ./cmd/design-records-mcp --help`: printed usage for `design-records-mcp`; command exited 1 with `flag: help requested`, which is the binary's help behavior.
- Post-edit Design Records MCP validation for `V01-TASK-DATA-012-01`: passed with `ok: true`, `diagnostics: null`.
- Post-edit Design Records MCP validation for `V01-WORK-DATA-012`: passed with `ok: true`, `diagnostics: null`.
- `git diff -- docs/tasks/data/TASK-DATA-012-01-review-uc-002-enum-literal-and-vocabulary-cleanup-candidates.md docs/work-items/data/WORK-DATA-012-uc-002-enum-literal-vocabulary-cleanup.md`: normal tracked diff showed only the existing `V01-WORK-DATA-012.tasks` relation addition for `V01-TASK-DATA-012-01`; the task file is currently untracked, so its content is not shown by normal `git diff`.
- `git status --short`: showed `?? docs/tasks/data/TASK-DATA-012-01-review-uc-002-enum-literal-and-vocabulary-cleanup-candidates.md`, a tracked modification to `docs/work-items/data/WORK-DATA-012-uc-002-enum-literal-vocabulary-cleanup.md`, and unrelated existing MCP docs/workflow changes.

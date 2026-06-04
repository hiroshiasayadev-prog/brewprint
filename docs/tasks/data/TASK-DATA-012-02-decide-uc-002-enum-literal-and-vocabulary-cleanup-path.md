# TASK-DATA-012-02: Decide UC-002 enum, literal, and vocabulary cleanup path

- **id**: TASK-DATA-012-02
- **status**: done
- **date**: 2026-06-05
- **work_item**: WORK-DATA-012
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-012-01
- **outputs**:
  - Cleanup path decision table for WORK-DATA-012 candidates
  - Explicit enum reuse / new enum / literal-note / no-action decisions
  - Input for follow-up YAML cleanup, spec-gap, or no-action tasks

## Goal

Decide the cleanup path for the UC-002 enum-like, literal, and usage-site vocabulary candidates classified by the previous candidate review task.

This task must turn the candidate review table into concrete next actions, but it must not perform YAML migration, fixture regeneration, implementation changes, or closed-work reopen.

The task should reduce decision burden for the project owner by recording a recommended default decision for each unclear point with evidence and fallback options.

## Work

- Review the previous candidate classification result for:
  - N-006 residual;
  - N-015 residual;
  - N-019;
  - N-023 residual;
  - N-029 residual;
  - N-030;
  - N-034;
  - N-045;
  - N-046;
  - N-051.
- For N-019, decide whether to reuse or generalize an existing direction enum, or create a new query/reference direction enum.
- For N-030, decide whether to use a narrow list-objects filter enum or a broader MCP object type enum.
- For N-023 residual, decide whether `language = yaml` has a supported literal-constraint representation, should remain note-only, or needs a separate requirement / ADR.
- For N-034, N-045, and N-046, decide the named enum outcome and intended reuse boundary.
- For N-006 residual, N-015 residual, and N-051, decide which usage-site vocabulary notes are YAML-cleanup-ready and which should remain generic or note-only.
- Confirm N-029 residual no-action unless new evidence identifies an explicit closed vocabulary.
- Keep default/behavior, selector/support matrix, recursive/untagged-union, tagged-union, and MCP identity concerns routed to their existing owners.

## Done condition

- Each candidate from the previous classification has exactly one selected cleanup path or explicit no-action outcome.
- Any recommended new enum has a proposed name, value set, and reuse boundary.
- Any literal-related item is classified as supported now, note-only, or requiring a separate follow-up requirement / ADR.
- The resulting decisions are sufficient input for a follow-up YAML cleanup task or a spec-gap follow-up task.
- No UC-002 YAML, fixture, golden, parser, renderer, validator, MCP implementation, ADR, spec, or closed DATA work item is changed by this task.

## Verification

- Confirm the decision table covers all candidates reviewed by the previous candidate review task.
- Confirm selected outcomes do not route default/behavior scope into `WORK-DATA-012`.
- Confirm selected outcomes do not route selector/support-matrix scope into `WORK-DATA-012`.
- Confirm selected outcomes do not reopen M15, `WORK-DATA-001`, `WORK-DATA-006`, `WORK-DATA-009`, or `WORK-DATA-010`.
- Run Design Records MCP validation for this task and `WORK-DATA-012` after creation.

## Evidence
Completed on 2026-06-05.

### Sources reviewed

- `docs/prompt_chappy.md`
- `AGENTS.md`
- `docs/AGENTS.md`
- `docs/doc-policy.md`
- `docs/tasks/data/TASK-DATA-012-02-decide-uc-002-enum-literal-and-vocabulary-cleanup-path.md`
- `docs/tasks/data/TASK-DATA-012-01-review-uc-002-enum-literal-and-vocabulary-cleanup-candidates.md`
- `docs/work-items/data/WORK-DATA-012-uc-002-enum-literal-vocabulary-cleanup.md`
- `docs/tasks/data/TASK-DATA-009-03-decide-remaining-uc-002-notes-retreat-successor-outcomes.md`
- `docs/tasks/data/TASK-DATA-009-04-create-follow-up-split-or-close-remaining-uc-002-notes-retreat-classification.md`
- `docs/investigations/data/INV-DATA-002-uc002-notes-retreat-inventory-and-m15-release-boundary-input.md`
- `docs/adr/067-enum-model.md`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/reference.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/mcp_error.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/common.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/list_objects_request.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_reference_tree_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_references_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/analyze_impact_response.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/string_list.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/get_source_response.yaml`
- `docs/spec/mcp/schema.md`
- `docs/spec/mcp/errors.md`
- `docs/spec/mcp/tools/list-objects.md`
- `docs/spec/mcp/tools/get-references.md`
- `docs/spec/mcp/tools/analyze-impact.md`

### Decision summary

| candidate ID | selected outcome | proposed enum / note | values | reuse boundary | confidence |
|---|---|---|---|---|---|
| N-034 | create new named enum | `mcp_error_code` | `project_invalid`, `invalid_args`, `invalid_selector`, `invalid_change_payload`, `not_found`, `kind_mismatch`, `ambiguous`, `unsupported_object`, `unsupported_detail`, `unsupported_direction`, `invalid_depth`, `internal_error` | `mcp_error.code` only | high |
| N-045 | create new named enum | `reference_kind` | `param_model`, `return_model`, `produces_asset`, `consumes_asset`, `reads`, `writes`, `store_of`, `field_type`, `field_fk`, `transition_event`, `transition_from`, `transition_to`, `transition_action`, `event_payload`, `event_actor`, `event_watches`, `scenario_state_file`, `scenario_step_transition` | `reference.kind`, `get_reference_tree_edge.kind`, and selected `via` usage sites | high |
| N-046 | create new named enum | `reference_direction` | `out`, `in` | `reference.direction`, `get_reference_tree_edge.direction` | high |
| N-030 | create new named enum | `list_objects_object_filter` | `node`, `view`, `transition`, `field` | `list_objects_request.object` only | high |
| N-019 | create new named enum | `reference_query_direction` | `out`, `in`, `both` | `get_references_response.direction`; future flat reference-query direction fields only | medium |
| N-015 residual | reuse named enums | `reference_kind` and `reference_direction` | inherits N-045 / N-046 | `get_reference_tree_edge.kind` and `get_reference_tree_edge.direction`; no independent edge enum | high |
| N-006 residual | create new named enum | `analyze_impact_coverage_vocabulary` | `direct_references`, `reference_tree`, `model_field_resolution`, `transition_action_resolution`, `flow_step_task_resolution`, `flow_param_field_resolution`, `sequence_step_task_resolution`, `type_signature_identity`, `render_output_files`, `name_collision`, `type_structural_compatibility`, `semantic_contract_compatibility`, `render_presentation_details`, `wireframe_element_binding` | `analyze_impact_coverage.analyzed` and `analyze_impact_coverage.not_analyzed` | medium |
| N-051 | usage-site vocabulary cleanup only | keep shared `string_list` generic; replace only closed-vocabulary usage sites with typed lists | n/a | `get_reference_tree_node.via` to `list<reference_kind>`; coverage lists to `list<analyze_impact_coverage_vocabulary>` if N-006 is implemented | high |
| N-023 residual | literal note-only | keep `type: str` plus note that `language` is literal `yaml` | n/a | `get_source_snippet.language`; no enum / literal model change now | high |
| N-029 residual | obsolete / no-action | no enum or literal cleanup | n/a | reopen only if a future spec defines a closed HTTP method value set | medium |

### Key decisions

- N-034 is ready for YAML cleanup as `mcp_error_code`; the YAML note and MCP errors spec define the same closed 12-value set.
- N-045 is ready for YAML cleanup as `reference_kind`; the public `Reference.kind` vocabulary is distinct from MCP identity / semantic reference identity work.
- N-046 is ready for YAML cleanup as `reference_direction` with only `out` and `in`. It must not be reused for query fields that allow `both`.
- N-030 uses a narrow `list_objects_object_filter` enum. Broader `mcp_object_type` reuse is rejected because it would admit invalid filter values such as `asset`, `file`, and `primitive` for the current list-objects request field.
- N-019 uses a new `reference_query_direction` enum. `reference_tree_direction` has the same values, but its name and notes are tree-specific. Reusing it for flat `get_references` query output would create semantic overreach and future confusion.
- N-015 residual reuses N-045 / N-046 outcomes for tree edge `kind` and `direction`; no separate edge enum is needed.
- N-006 creates one shared `analyze_impact_coverage_vocabulary` for both `analyzed` and `not_analyzed`. The enum should contain only the current v1 standard vocabulary; future coverage terms should update the enum through future spec-aligned work.
- N-051 does not change the shared `string_list` helper. Only confirmed closed-vocabulary usage sites should move to typed lists.
- N-023 remains note-only. `language = yaml` is a literal constraint, not a finite selection enum in current ADR-067 terms, and no separate literal-constraint requirement is needed now.
- N-029 remains no-action because no closed HTTP method value set is defined by current evidence.

### YAML cleanup input

The follow-up YAML cleanup task should include these items if selected:

| candidate | YAML cleanup input | prerequisite |
|---|---|---|
| N-034 | change `mcp_error.code` from `str` to `mcp_error_code` | create `mcp_error_code` enum |
| N-045 | change `reference.kind` from `str` to `reference_kind` | create `reference_kind` enum |
| N-046 | change `reference.direction` from `str` to `reference_direction` | create `reference_direction` enum |
| N-015 residual | change `get_reference_tree_edge.kind` to `reference_kind`; change `get_reference_tree_edge.direction` to `reference_direction` | N-045 / N-046 enum creation |
| N-030 | change `list_objects_request.object` from `str` to `list_objects_object_filter` | create `list_objects_object_filter` enum |
| N-019 | change `get_references_response.direction` from `str` to `reference_query_direction` | create `reference_query_direction` enum |
| N-006 residual | change `coverage.analyzed` and `coverage.not_analyzed` from `string_list` to list of `analyze_impact_coverage_vocabulary` | create `analyze_impact_coverage_vocabulary` enum |
| N-051 selected | change `get_reference_tree_node.via` from `string_list` to list of `reference_kind`; keep open/free-text `string_list` usages unchanged | N-045 enum creation |

### No-action / note-only outcomes

- N-023 residual: keep `get_source_snippet.language` as `type: str` with a literal `yaml` note. Do not create a one-value enum now.
- N-029 residual: no enum/literal cleanup for HTTP method wording unless a future spec introduces an explicit closed value set.
- No separate requirement / ADR is needed from this task. Literal constraints and coverage vocabulary expansion can be revisited only when a concrete future requirement appears.

### Out-of-scope routing confirmation

- Behavior / default / fallback rules remain owned by `REQ-DATA-006` / `WORK-DATA-013`.
- Selector / support matrix / object-dependent kind rules remain owned by `REQ-DATA-007` / `WORK-DATA-014`.
- Recursive / untagged-union representation remains owned by `REQ-DATA-008` / `WORK-DATA-015`.
- Tagged union / discriminator payload remains owned by `REQ-DATA-004` / closed `WORK-DATA-010`.
- MCP identity / semantic reference identity remains owned by `REQ-MCP-004` / `WORK-MCP-004`.

### Risks and ambiguity

- `reference_query_direction` duplicates the value set of `reference_tree_direction`, but this is intentional to avoid tree-specific naming leaking into flat query semantics.
- `analyze_impact_coverage_vocabulary` may require future updates as coverage vocabulary evolves. Current cleanup should include only current v1 standard terms.
- Typed list migration for usage sites depends on existing parser / renderer support for list-of-named-model TypeRef and must be verified in the YAML cleanup task.

### Verification note

The review result was accepted as the basis for this task's decisions, with the project owner explicitly choosing the new `reference_query_direction` enum for N-019 rather than reusing `reference_tree_direction`.

This task intentionally performed no UC-002 YAML migration, fixture regeneration, golden update, parser change, renderer change, validator change, MCP implementation change, ADR/spec update, or closed DATA work reopen.

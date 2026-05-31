# TASK-DATA-006-01: UC-002 helper-shape migration set selection

- **id**: TASK-DATA-006-01
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-DATA-006
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d
- **depends_on**:
  - TASK-DATA-003-04
- **outputs**:
  - Selected UC-002 model-file response helper migration set
  - Deferred task-file any / private-helper-blocked candidate classification
  - Boundary for TASK-DATA-006-02 YAML migration

## Goal

Select the exact UC-002 helper-shape migration set for WORK-DATA-006 before any UC-002 YAML, fixture, golden output, or Go implementation changes.

This task converts the TASK-DATA-003-04 candidate inventory and the WORK-DATA-004 / REQ-DATA-003 task-file signature policy into a concrete migration boundary for the next task.

## Work

- Review the TASK-DATA-003-04 model-file helper migration candidates.
- Use the WORK-DATA-004 / REQ-DATA-003 task-file private helper signature policy as a constraint.
- Select the response-local model-file helper candidates that can be migrated as same-file private helper models.
- Classify task-file `query_result:any` candidates as blocked / deferred for this work item.
- Keep UC-002 YAML, fixture / golden outputs, and Go implementation unchanged in this task.

## Selected migration set

These candidates are selected for TASK-DATA-006-02.

They are model-file response-local nested shapes. The owning response model remains public, and the nested shape may be migrated to a same-file private helper model. Public model files are not required for these selected candidates.

| candidate | location | selected handling |
|---|---|---|
| N-005 | `analyze_impact_response.impacts` | migrate as same-file private helper; keep `suggested_fixes` kind-dependent payload as tagged-union / note debt |
| N-006 | `analyze_impact_response.coverage` | migrate as same-file private helper; keep vocabulary constraints as note / enum follow-up |
| N-014 | `get_reference_tree_response.nodes` | migrate as same-file private helper |
| N-015 | `get_reference_tree_response.edges` | migrate as same-file private helper; avoid assuming inheritance / extension of `reference` |
| N-023 | `get_source_response.snippet` | migrate as same-file private helper; keep `language: yaml` literal constraint as note |
| N-029 | `list_endpoints_response.tables` | migrate nested table / section / endpoint response shapes as same-file private helpers; keep optional semantics as notes |
| N-033 | `list_objects_response.objects` | migrate as same-file private helper; keep identity semantics out of scope |

## Deferred / blocked candidates

### Task-file `query_result:any` candidates

The 8 UC-002 MCP task files keep the current `query_service.returns.model: any` to `build_response.params[].model: any` pattern for WORK-DATA-006.

- `analyze_impact.yaml`
- `get_reference_tree.yaml`
- `get_references.yaml`
- `get_signature.yaml`
- `get_source.yaml`
- `inspect.yaml`
- `list_endpoints.yaml`
- `list_objects.yaml`

These candidates are not migrated to same-file task-file private helper models in WORK-DATA-006. `returns.model` alone could use a task-file private helper, but using the same private helper from `build_response.params[].model` would violate the accepted DATA-004 / REQ-DATA-003 policy, where task `params[].model` references to same-file private helper models are invalid.

Later work may choose one of these approaches:

- keep `any` / note-based shape,
- use public model files,
- migrate only returns-side helper shapes,
- or introduce a later design that handles cross-step task-local result typing without exposing private helper params.

### Other excluded debt

The following categories remain outside this task and outside TASK-DATA-006-02:

- tagged union / discriminator payloads, including `analyze_impact_response.change` and `suggested_fixes` kind-dependent payloads,
- recursive / union / arbitrary member shapes, including `get_signature_response.signature`, `inspect_response.signature`, `inspect_response.members`, `diagnostic.related`, and `object_ref.parent`,
- container-only or constraint-only notes such as string arrays, dictionary summaries, enum / range / optional constraints,
- stale UC-002 README render-status cleanup,
- Go implementation changes for task-file private helper params policy or cross-step task result typing.

## Done condition

- Selected model-file response helper migration set is recorded.
- Task-file `query_result:any` candidates are explicitly classified as blocked / deferred for this work item.
- Public model file cut-out is not required for the selected 7 model-file response candidates.
- TASK-DATA-006-02 can start from this selected set without reopening WORK-DATA-002, WORK-DATA-003, WORK-DATA-004, tagged union, MCP identity, or M15 scope.
- No UC-002 YAML, fixture / golden output, or Go implementation change is performed in this task.

## Verification

- Confirm this task only records selection and boundary.
- Confirm selected candidates match the model-file helper migration candidates recorded by TASK-DATA-003-04.
- Confirm task-file `query_result:any` candidates are not selected for migration because of the DATA-004 / REQ-DATA-003 `params[].model` private-helper policy.
- Confirm WORK-DATA-006 is the parent work item.

## Evidence

Reviewed on 2026-06-01.

The selected model-file response helper migration set is:

- N-005: `analyze_impact_response.impacts`
- N-006: `analyze_impact_response.coverage`
- N-014: `get_reference_tree_response.nodes`
- N-015: `get_reference_tree_response.edges`
- N-023: `get_source_response.snippet`
- N-029: `list_endpoints_response.tables`
- N-033: `list_objects_response.objects`

The selected candidates remain same-file private helper migrations because they are response-local nested shapes under public response models.

The 8 task-file `query_result:any` candidates are deliberately deferred. They would require either public model files, returns-only migration, keeping `any`, or later design work because task `params[].model` cannot reference same-file task-file private helper models under the accepted DATA-004 / REQ-DATA-003 policy.

No UC-002 YAML, render output, fixture, golden output, or Go implementation change was performed by this task.

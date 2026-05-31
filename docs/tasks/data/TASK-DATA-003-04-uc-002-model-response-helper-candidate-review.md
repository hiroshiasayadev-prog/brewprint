# TASK-DATA-003-04: UC-002 model response helper candidate review

- **id**: TASK-DATA-003-04
- **status**: done
- **date**: 2026-05-31
- **work_item**: WORK-DATA-003
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-003-03
- **outputs**:
  - UC-002 model response helper candidate classification
  - Boundary note delegating actual migration to WORK-DATA-004 or later work
  - WORK-DATA-003 close-readiness input

## Goal

Review UC-002 model response helper-shape candidates after model-file render support is available, and classify what should be migrated later without performing the migration in this task.

This task is the WORK-DATA-003 close-readiness review for UC-002 migration scope. Actual YAML / fixture / golden migration is delegated to WORK-DATA-004 or later follow-up work so that WORK-DATA-003 does not mix model-file render exposure with later helper policy work.

## Work

- Review the relevant UC-002 model response helper-shape candidates recorded by prior investigations / requirements.
- Classify candidates into:
  - model-file helper migration candidates
  - later helper policy candidates that must wait for REQ-DATA-003 / WORK-DATA-004
  - candidates that should remain unchanged
- Confirm model-file render support from TASK-DATA-003-03 removes the render-blocker for model-file helper candidates.
- Record why actual migration is delegated out of WORK-DATA-003.
- Update WORK-DATA-003 close-readiness evidence if needed.

## Included Scope

- Review and classification only.
- UC-002 migration boundary clarification.
- WORK-DATA-003 close-readiness input.
- Delegation note for actual migration to WORK-DATA-004 or later follow-up work.

## Excluded Scope

- UC-002 YAML migration.
- Fixture / golden regeneration for UC-002.
- Renderer implementation.
- REQ-DATA-003 implementation.
- WORK-DATA-004 implementation or status correction.
- WORK-DATA-002 reopening.
- Tagged union rendering / ADR-073 implementation.

## Done Condition

- UC-002 model response helper-shape candidates are classified.
- Actual migration is explicitly delegated out of WORK-DATA-003.
- Any dependency on REQ-DATA-003 / WORK-DATA-004 is identified without implementing it.
- WORK-DATA-003 has enough evidence to decide whether it can close after model-file render exposure and candidate classification.
- No UC-002 YAML, fixture / golden, renderer, or later helper policy implementation is performed.

## Verification

- Confirm no UC-002 YAML / render output files were changed by this task.
- Confirm TASK-DATA-003-03 model-file render support is already done.
- Confirm REQ-DATA-003 / WORK-DATA-004 scope is not implemented here.
- Confirm WORK-DATA-002 remains closed.

## Evidence

Reviewed on 2026-06-01.

### Sources reviewed

- `REQ-DATA-002`: captures the deferred helper-model and model-render follow-up for UC-002 contract shapes, without deciding the exact migration set.
- `WORK-DATA-002`: closed the task-file helper model minimum and deferred model-file helper render plus UC-002 response helper-shape migration to WORK-DATA-003 or later.
- `WORK-DATA-003`: owns model-file helper render exposure and candidate classification, while delegating actual UC-002 migration to WORK-DATA-004 or later follow-up work.
- `INV-DATA-002`: records the UC-002 note-retreat inventory and the concrete candidate IDs.
- `TASK-DATA-003-01`: classifies N-014, N-015, and N-023 as model-file helper candidates once model-file render is available.
- `TASK-DATA-003-02`: aligns `docs/spec/views/model-file.md` and explicitly excludes UC-002 migration from the model-file render minimum.
- `TASK-DATA-003-03`: records implemented and verified model-file render support for public model files plus same-file private helper models.
- `REQ-DATA-003` / `WORK-DATA-004`: define the separate task-file private helper signature policy; read as dependency context only.
- UC-002 YAML reviewed for the relevant model response files and MCP task files; no UC-002 YAML or render artifact was changed.

### Candidate classification

#### Model-file helper migration candidates

These candidates are response-model fields in `docs/uc/002-brewprint-self-hosting/yaml/mcp/model/*.yaml` where the owning public response model can remain public and the hidden shape can later become same-file private helper models. TASK-DATA-003-03 removes the render visibility blocker for these candidates because model files now render public models and same-file private helpers.

| candidate | location | classification | notes |
|---|---|---|---|
| N-005 | `analyze_impact_response.impacts` | model-file helper migration candidate | Outer impact-entry list / entry shape can be migrated later. Nested `suggested_fixes` kind-dependent payload remains tagged-union / note debt and must not pull ADR-073 into this migration. |
| N-006 | `analyze_impact_response.coverage` | model-file helper migration candidate | Small response-local coverage object. Vocabulary constraints can remain notes or enum follow-up. |
| N-014 | `get_reference_tree_response.nodes` | model-file helper migration candidate | Node-entry list with `object:ObjectRef`, `depth:int`, and `via:string[]`; not intrinsically tagged union. |
| N-015 | `get_reference_tree_response.edges` | model-file helper migration candidate | Edge-entry list with Reference-like fields plus `depth:int`; migration needs care because v1 has no inheritance / extension of `reference`. |
| N-023 | `get_source_response.snippet` | model-file helper migration candidate | Small response-local snippet object; `language: yaml` literal constraint can remain a note or later enum-style constraint. |
| N-029 | `list_endpoints_response.tables` | model-file helper migration candidate | Nested table / section / endpoint response shape can be represented with response-local helper models, while optional-field semantics remain notes. |
| N-033 | `list_objects_response.objects` | model-file helper migration candidate | ObjectRef-like list entry with list-specific summary fields. Identity semantics remain ADR-078+ / note debt and are not solved by helper migration. |

#### REQ-DATA-003 / WORK-DATA-004 wait candidates

The UC-002 MCP task files all use the same task-local pattern: `query_service` returns `model: any`, and `build_response.params[].model` consumes that `query_result` as `any`.

Files checked:

- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/analyze_impact.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_reference_tree.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_references.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_signature.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/get_source.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/inspect.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/list_endpoints.yaml`
- `docs/uc/002-brewprint-self-hosting/yaml/mcp/task/list_objects.yaml`

If these `query_result` shapes are migrated as task-file private helper models, the helper would appear both in `query_service.returns.model` and in `build_response.params[].model`. REQ-DATA-003 accepts an asymmetric policy: task-file private helper references from `returns.model` may be allowed, but `params[].model` references are invalid because params are caller-provided input contracts. Therefore these task-file helper attempts must wait for WORK-DATA-004 or a later design that avoids private-helper params. They are not WORK-DATA-003 model-file migration candidates.

#### Unchanged candidates

The following UC-002 note-retreat candidates should remain unchanged by this task and by the WORK-DATA-003 close boundary:

- Tagged union / kind-specific response payloads: N-003, N-021, N-026, N-027, plus related request-side N-001. These remain ADR-073 or later tagged-union work.
- Union / recursive / currently unsupported structural constraints: N-009 and N-044.
- Pure enum / value-set / default / numeric-range / support-matrix constraints, including N-010, N-011, N-013, N-017, N-019, N-020, N-022, N-024, N-025, N-030, N-031, N-034, N-036, N-037, N-040, N-041, N-042, N-045, N-046, and N-051. Some enum migration was already handled by M15; the remaining constraints are not model-file helper render blockers.
- MCP identity / semantic reference debt: N-032, N-035, N-038, N-039, N-043, N-047, N-048, and N-050. These remain ADR-078+ / MCP identity style work, not helper-shape migration.
- Request-side or generic container candidates whose primary issue is not response-local helper shape, such as N-002, N-004, N-007, N-008, N-012, N-016, N-018, N-028, and N-049. They can be reconsidered by later notes-retreat work, but they are not required to close WORK-DATA-003.
- N-055 and N-056 are human explanation / view-renderer notes and remain non-issues for helper migration.

### Delegation rationale

Actual migration is delegated out of WORK-DATA-003 because WORK-DATA-003's required capability was render exposure for model-file helpers, not UC-002 YAML reshaping. Migrating UC-002 would change YAML and require fixture / golden regeneration, and task-file `query_result` helper attempts intersect with REQ-DATA-003 / WORK-DATA-004 signature policy. Keeping this task to review and classification preserves the boundary stated in WORK-DATA-003 and avoids pulling tagged union, MCP identity, DAG TypeRef hint, WORK-DATA-004, or remaining notes-retreat debt into the model-file render close path.

### Close-readiness input

WORK-DATA-003 has enough evidence to close once this task update is accepted:

- TASK-DATA-003-01 records ADR-075 as usable for the model-file render minimum while deferring tagged union rendering.
- TASK-DATA-003-02 records the spec owner and scope for model-file render.
- TASK-DATA-003-03 records implementation, fixture / golden evidence, and passing verification for model-file render.
- This task classifies the UC-002 model response helper candidates and delegates actual migration out of WORK-DATA-003.
- No UC-002 YAML, UC-002 render output, renderer implementation, REQ-DATA-003 / WORK-DATA-004 implementation, WORK-DATA-002 reopening, or ADR-073 implementation was performed.

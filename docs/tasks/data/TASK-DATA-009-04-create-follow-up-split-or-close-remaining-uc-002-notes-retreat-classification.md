# TASK-DATA-009-04: Create follow-up split or close remaining UC-002 notes retreat classification

- **id**: TASK-DATA-009-04
- **status**: done
- **date**: 2026-06-01
- **work_item**: WORK-DATA-009
- **source_requirement**: REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - TASK-DATA-009-03
- **outputs**:
  - Follow-up requirements and/or work items created where needed
  - Explicit no-action / obsolete close outcomes
  - WORK-DATA-009 close evidence

## Goal

Create the selected successor planning artifacts from `TASK-DATA-009-03`, record explicit no-action outcomes for buckets that need no new artifact, and close `WORK-DATA-009` without performing direct UC-002 cleanup.

## Work

- Review the `TASK-DATA-009-03` successor outcome table.
- Create only the follow-up requirement and work item artifacts selected by that table.
- Reuse existing owners for buckets already covered by existing work.
- Record no-action and obsolete outcomes explicitly.
- Update `WORK-DATA-009` task flow, status, and close evidence.
- Do not perform implementation, UC-002 YAML migration, fixture / golden regeneration, or ADR implementation.

## Done Condition

- Every `TASK-DATA-009-03` outcome has either a created successor artifact, an existing owner, or an explicit no-action / obsolete close rationale.
- Newly created requirements and work items have valid metadata, source references, and scope boundaries.
- `WORK-DATA-009` includes `TASK-DATA-009-04`, assigns `T4` to it, and has close evidence.
- `TASK-DATA-009-04` and `WORK-DATA-009` can be marked `done`.
- No out-of-scope implementation or UC-002 cleanup file is changed.

## Verification

- Validate `TASK-DATA-009-04`, `WORK-DATA-009`, and all newly created successor requirements / work items through Design Records MCP.
- Retrieve the updated records through Design Records MCP where relevant to confirm metadata and relation updates.
- Check the working tree for touched files and confirm no out-of-scope files were edited by this task.

## Evidence

Completed on 2026-06-01.

### Sources reviewed

- `AGENTS.md`
- `docs/AGENTS.md`
- `docs/prompt_chappy.md`
- `docs/doc-policy.md`
- `docs/guides/requirement-authoring.md`
- `docs/guides/work-item-authoring.md`
- `docs/guides/task-authoring.md`
- Design Records MCP authoring guidance for `requirement-authoring`, `work-item-authoring`, `task-authoring`, and `artifact-boundary`
- `docs/requirements/data/REQ-DATA-002-helper-model-and-model-render-follow-up.md`
- `docs/requirements/data/REQ-DATA-004-tagged-union-and-discriminator-payload-support.md`
- `docs/requirements/data/REQ-DATA-005-dag-asset-typeref-hint-render-support.md`
- `docs/investigations/data/INV-DATA-002-uc002-notes-retreat-inventory-and-m15-release-boundary-input.md`
- `docs/tasks/data/TASK-DATA-003-04-uc-002-model-response-helper-candidate-review.md`
- `docs/tasks/data/TASK-DATA-005-01-m15-deferred-item-inventory.md`
- `docs/tasks/data/TASK-DATA-005-02-deferred-ownership-classification.md`
- `docs/tasks/data/TASK-DATA-005-03-create-successor-split.md`
- `docs/tasks/data/TASK-DATA-006-04-close-and-follow-up-split.md`
- `docs/tasks/data/TASK-DATA-009-01-reconcile-remaining-uc-002-notes-retreat-candidates.md`
- `docs/tasks/data/TASK-DATA-009-02-classify-remaining-uc-002-notes-retreat-successor-buckets.md`
- `docs/tasks/data/TASK-DATA-009-03-decide-remaining-uc-002-notes-retreat-successor-outcomes.md`
- `docs/work-items/data/WORK-DATA-006-helper-shape-migration.md`
- `docs/work-items/data/WORK-DATA-007-dag-asset-typeref-hint-render-support.md`
- `docs/work-items/data/WORK-DATA-008-uc-002-duplicate-task-qid-unresolved-flow-task.md`
- `docs/work-items/data/WORK-DATA-009-remaining-uc-002-notes-retreat-classification.md`
- `docs/work-items/data/WORK-DATA-010-tagged-union-and-discriminator-payload-support.md`

### Input summary from TASK-DATA-009-03

`TASK-DATA-009-03` selected these successor outcomes:

- `request-side / generic container`: create a new DATA work item under existing `REQ-DATA-002`.
- `enum-like / literal constraint`: create a new DATA work item under existing `REQ-DATA-002`.
- `numeric / default behavior`: create a new requirement plus work item for request option behavior and response behavior constraints.
- `selector matrix / support matrix`: create a new requirement plus work item for selector support matrices and object-dependent kind vocabulary.
- `recursive / union structure`: create a new DATA expressiveness requirement plus work item for recursive and untagged-union representation.
- `MCP identity / semantic reference`: do not create new DATA work; use existing `REQ-MCP-004` / `WORK-MCP-004`.
- `human explanation / view-renderer note`: no action.
- `obsolete / no-action candidate`: no action.
- `other still-unowned`: no action because the bucket is empty.

### Follow-up split table

| bucket | outcome from TASK-DATA-009-03 | action taken in this task | created artifact ID | no-action / deferred rationale |
|---|---|---|---|---|
| `request-side / generic container` | `new work item needed under existing requirement` | Created a successor DATA cleanup work item under `REQ-DATA-002`. | `WORK-DATA-011` | n/a |
| `enum-like / literal constraint` | `new work item needed under existing requirement` | Created a successor DATA cleanup work item under `REQ-DATA-002`. | `WORK-DATA-012` | n/a |
| `numeric / default behavior` | `new requirement needed` | Created a DATA requirement and matching work item for request option and response behavior constraints. | `REQ-DATA-006`, `WORK-DATA-013` | n/a |
| `selector matrix / support matrix` | `new requirement needed` | Created a DATA requirement and matching work item for selector support matrices and object-dependent vocabulary. | `REQ-DATA-007`, `WORK-DATA-014` | n/a |
| `recursive / union structure` | `new requirement needed` | Created a DATA requirement and matching work item for recursive and untagged-union representation. | `REQ-DATA-008`, `WORK-DATA-015` | n/a |
| `MCP identity / semantic reference` | `covered by existing work` | No new DATA artifact created. Existing MCP owner recorded. | n/a | Covered by `REQ-MCP-004` / `WORK-MCP-004`; duplicating DATA identity work would conflict with the prior split. |
| `human explanation / view-renderer note` | `obsolete / no-action` | No artifact created. | n/a | N-055 and N-056 are explanatory notes rather than hidden machine-readable schema debt. |
| `obsolete / no-action candidate` | `obsolete / no-action` | No artifact created. | n/a | N-052 to N-054 are secondary UC-001 rows; N-049 remains intentionally non-public render-context mapping; residual N-048 does not create DATA work unless future MCP identity work needs a public reference-index map schema. |
| `other still-unowned` | `obsolete / no-action` | No artifact created. | n/a | `TASK-DATA-009-02` found no remaining candidates for this bucket, so no placeholder successor is needed. |

### Created artifact summary

- `WORK-DATA-011`: owns request-side / generic container cleanup planning for N-002, N-004, N-007, N-008, N-012, N-016, N-018, and `TF-QUERY-RESULT` under `REQ-DATA-002`.
- `WORK-DATA-012`: owns remaining enum / literal / usage-site vocabulary cleanup planning for N-019, N-030, N-034, N-045, N-046, N-051, and residual N-006 / N-015 / N-023 / N-029 vocabulary or literal notes under `REQ-DATA-002`.
- `REQ-DATA-006` / `WORK-DATA-013`: own numeric range, default, omitted-value, unknown-value, fallback, and cross-response behavior constraints for N-011, N-017, N-022, N-024, N-025, and N-028.
- `REQ-DATA-007` / `WORK-DATA-014`: own selector support matrices and object-dependent kind vocabulary for N-020, N-031, N-037, N-040, and N-042.
- `REQ-DATA-008` / `WORK-DATA-015`: own recursive and untagged-union representation for N-009 and N-044, separate from the tagged / discriminated union scope in `REQ-DATA-004` / `WORK-DATA-010`.

### Ownership clarification

`TASK-DATA-009-03` used "likely MCP-domain" / "MCP/tool-contract" for the `numeric / default behavior` and `selector matrix / support matrix` buckets as candidate decision input, not as final artifact ownership. This task intentionally resolved those UC-002 note-retreat contract gaps as DATA-domain follow-ups (`REQ-DATA-006` / `WORK-DATA-013` and `REQ-DATA-007` / `WORK-DATA-014`) because they concern data-shape expressiveness and constraint representation in UC-002 examples. Primary MCP identity ownership remains outside this split and stays with existing MCP-domain ownership, especially `REQ-MCP-004` / `WORK-MCP-004`. Do not create duplicate MCP requirements or work items for the same numeric/default or selector/support-matrix buckets unless new evidence changes that boundary.

### Explicit no-action / obsolete close outcomes

- No new DATA artifact was created for primary MCP identity / semantic-reference candidates because `REQ-MCP-004` / `WORK-MCP-004` already own that successor scope.
- No artifact was created for N-055 and N-056 because they are human explanation / view-renderer notes without hidden machine-readable schema debt.
- No artifact was created for N-052, N-053, and N-054 because they are secondary UC-001 rows outside the remaining UC-002 cleanup scope.
- No artifact was created for N-049 because `resolved_project.render_context` remains intentionally non-public mapping material for `analyze_impact`.
- No DATA artifact was created for the residual N-048 map-shape question; if `WORK-MCP-004` later needs a public reference-index map schema, that MCP work owns the split.
- No artifact was created for `other still-unowned` because the bucket is empty.

### Deferred outcomes

No bucket was kept deferred by this task.

The only conditional reopen evidence recorded is for the residual N-048 map-shape question: reopen through the MCP identity owner only if `WORK-MCP-004` identifies a concrete public reference-index map schema need. The deferred note is preserved in `TASK-DATA-009-03` and this task evidence.

### WORK-DATA-009 close readiness statement

`WORK-DATA-009` is ready to close because:

- `TASK-DATA-009-01` reconciled the remaining candidate inventory.
- `TASK-DATA-009-02` classified the remaining candidates into successor buckets.
- `TASK-DATA-009-03` selected covered, obsolete / no-action, and new-artifact outcomes.
- `TASK-DATA-009-04` created the selected follow-up artifacts and recorded explicit no-action outcomes.

No remaining `TASK-DATA-009-03` outcome lacks an owner, created artifact, or close rationale.

### Verification note

This task performed follow-up planning artifact creation and close evidence only.

No UC-002 YAML migration, fixture / golden regeneration, parser / renderer / validator / MCP implementation change, ADR-073 implementation, ADR-074 implementation, ADR-078 / ADR-079 / ADR-080 MCP identity implementation, or UC-002 duplicate task QID / unresolved flow task fix was performed.

### Post-edit verification

Design Records MCP validation passed for:

- `TASK-DATA-009-04`
- `WORK-DATA-009` through `WORK-DATA-015`
- `REQ-DATA-002` through `REQ-DATA-008`

Design Records MCP retrieval confirmed:

- `TASK-DATA-009-04` exists with status `done`.
- `WORK-DATA-009.tasks` includes `TASK-DATA-009-01`, `TASK-DATA-009-02`, `TASK-DATA-009-03`, and `TASK-DATA-009-04`.
- `WORK-DATA-009.status` is `done`.
- `REQ-DATA-006`, `REQ-DATA-007`, and `REQ-DATA-008` point to their matching work items.
- `WORK-DATA-011` through `WORK-DATA-015` point back to their source requirements and source planning artifacts.

Working tree review confirmed this task touched only DATA planning artifacts. Pre-existing unrelated dirty / untracked files remain present and were not edited by this task.

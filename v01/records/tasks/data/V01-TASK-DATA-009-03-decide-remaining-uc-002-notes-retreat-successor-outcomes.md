# V01-TASK-DATA-009-03: Decide remaining UC-002 notes retreat successor outcomes

- **id**: V01-TASK-DATA-009-03
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-DATA-009
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-DATA-009-02
- **outputs**:
  - Successor outcome decision table for remaining UC-002 notes retreat buckets
  - Explicit no-action / obsolete outcomes
  - Follow-up artifact creation input for V01-TASK-DATA-009-04 or later

## Goal

Decide the successor outcome for each remaining UC-002 notes retreat bucket classified by `V01-TASK-DATA-009-02`.

This task is decision / planning only. It does not create the follow-up requirement, work item, or task artifacts and does not perform UC-002 cleanup implementation.

## Work

- Use the bucket classification from `V01-TASK-DATA-009-02` as the input set.
- Select exactly one primary successor outcome for each bucket.
- Preserve already separated ownership for helper-shape migration, DAG TypeRef hint rendering, tagged union support, MCP identity, and the stale duplicate task QID / unresolved flow task issue.
- Record no-action / obsolete outcomes explicitly.
- Provide input for a later `V01-TASK-DATA-009-04` or later follow-up split / close task.

## Included Scope

- Bucket-level successor outcome decisions.
- Target existing owner references where a bucket is already covered.
- Proposed future artifact type where a new artifact is needed.
- Explicit no-action / obsolete rows.

## Excluded Scope

- UC-002 YAML migration.
- Fixture / golden regeneration.
- Parser, renderer, validator, or MCP implementation changes.
- V01-ADR-073 tagged union implementation.
- V01-ADR-074 DAG TypeRef hint implementation.
- V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP identity implementation.
- UC-002 duplicate task QID / unresolved flow task fix.
- Creation of new requirements or new work items.
- Creation of `V01-TASK-DATA-009-04` or later tasks.
- Marking `V01-WORK-DATA-009` as `done`.
- Reopening M15, V01-WORK-DATA-001, V01-WORK-DATA-002, V01-WORK-DATA-003, V01-WORK-DATA-004, V01-WORK-DATA-005, V01-WORK-DATA-006, V01-WORK-DATA-007, V01-WORK-DATA-008, or V01-WORK-DATA-010.

## Done Condition

- Each `V01-TASK-DATA-009-02` bucket has exactly one primary successor outcome.
- Existing ownership, no-action / obsolete outcomes, and proposed future artifact types are explicit.
- Follow-up artifact creation input is available for `V01-TASK-DATA-009-04` or later.
- No implementation, UC-002 cleanup, new requirement, new work item, or later task artifact is created.

## Verification

- Confirm only `V01-TASK-DATA-009-03` and `V01-WORK-DATA-009` documentation files are changed by this task.
- Confirm `V01-TASK-DATA-009-03` status is `done`.
- Confirm `V01-WORK-DATA-009.tasks` includes `V01-TASK-DATA-009-01`, `V01-TASK-DATA-009-02`, and `V01-TASK-DATA-009-03`.
- Confirm `V01-WORK-DATA-009` is not marked `done`.
- Run Design Records MCP validation for `V01-TASK-DATA-009-03` and `V01-WORK-DATA-009`.
- Retrieve both records through Design Records MCP if available.

## Evidence

Completed on 2026-06-01.

### Sources reviewed

- `AGENTS.md`
- `docs/AGENTS.md`
- `docs/prompt_chappy.md`
- `docs/doc-policy.md`
- `docs/guides/task-authoring.md`
- `docs/guides/work-item-authoring.md`
- Design Records MCP authoring guidance for `task-authoring` and `work-item-authoring`
- `docs/requirements/data/REQ-DATA-002-helper-model-and-model-render-follow-up.md`
- `docs/requirements/data/REQ-DATA-003-private-helper-model-signature-exposure-boundary.md`
- `docs/requirements/data/REQ-DATA-004-tagged-union-and-discriminator-payload-support.md`
- `docs/requirements/data/REQ-DATA-005-dag-asset-typeref-hint-render-support.md`
- `docs/requirements/mcp/REQ-MCP-004-mcp-semantic-identity-and-state-machine-identity-follow-up.md`
- `docs/investigations/data/INV-DATA-002-uc002-notes-retreat-inventory-and-m15-release-boundary-input.md`
- `docs/tasks/data/TASK-DATA-003-04-uc-002-model-response-helper-candidate-review.md`
- `docs/tasks/data/TASK-DATA-005-01-m15-deferred-item-inventory.md`
- `docs/tasks/data/TASK-DATA-005-02-deferred-ownership-classification.md`
- `docs/tasks/data/TASK-DATA-009-01-reconcile-remaining-uc-002-notes-retreat-candidates.md`
- `docs/tasks/data/TASK-DATA-009-02-classify-remaining-uc-002-notes-retreat-successor-buckets.md`
- `docs/work-items/data/WORK-DATA-006-helper-shape-migration.md`
- `docs/work-items/data/WORK-DATA-007-dag-asset-typeref-hint-render-support.md`
- `docs/work-items/data/WORK-DATA-008-uc-002-duplicate-task-qid-unresolved-flow-task.md`
- `docs/work-items/data/WORK-DATA-009-remaining-uc-002-notes-retreat-classification.md`
- `docs/work-items/data/WORK-DATA-010-tagged-union-and-discriminator-payload-support.md`
- `docs/work-items/mcp/WORK-MCP-004-mcp-semantic-identity-and-state-machine-identity-follow-up.md`

### Input summary from V01-TASK-DATA-009-02

`V01-TASK-DATA-009-02` classified the remaining `still-unowned` input from `V01-TASK-DATA-009-01` into these buckets:

- `request-side / generic container`: N-002, N-004, N-007, N-008, N-012, N-016, N-018, N-049, `TF-QUERY-RESULT`, and residual public-schema question from N-048.
- `enum-like / literal constraint`: N-019, N-030, N-034, N-045, N-046, N-051, and residual vocabulary / literal notes from N-006, N-015, N-023, and N-029.
- `numeric / default behavior`: N-011, N-017, N-022, N-024, N-025, and N-028.
- `selector matrix / support matrix`: N-020, N-031, N-037, N-040, and N-042.
- `recursive / union structure`: N-009 and N-044.
- `MCP identity / semantic reference`: no primary remaining `still-unowned` bucket; primary identity candidates were already reconciled as covered by `V01-REQ-MCP-004` / `V01-WORK-MCP-004`.
- `human explanation / view-renderer note`: no remaining `still-unowned` bucket; N-055 and N-056 were already reconciled as obsolete / no-action.
- `obsolete / no-action candidate`: likely N-049, residual N-048, and N-024 depending on later evidence; already reconciled obsolete / no-action candidates N-052, N-053, N-054, N-055, and N-056.
- `other still-unowned`: none.

Already separated work remains outside this task:

- Helper-shape migration is covered by `V01-WORK-DATA-006`.
- DAG asset TypeRef hint render support is covered by `V01-REQ-DATA-005` / `V01-WORK-DATA-007`.
- UC-002 duplicate task QID / unresolved flow task repair is covered and closed as stale by `V01-WORK-DATA-008`.
- Tagged union / discriminator payload candidates are covered by `V01-REQ-DATA-004` / `V01-WORK-DATA-010`.
- Primary MCP identity candidates are covered by `V01-REQ-MCP-004` / `V01-WORK-MCP-004`.

### Successor outcome decision table

| bucket name | candidate IDs / source note IDs | selected outcome | rationale | target existing owner if covered | proposed future artifact type if new artifact is needed |
|---|---|---|---|---|---|
| `request-side / generic container` | N-002, N-004, N-007, N-008, N-012, N-016, N-018, `TF-QUERY-RESULT`; excludes N-049 and residual public-schema question from N-048, which are decided in the no-action row | `new work item needed under existing requirement` | The bucket is still within the UC-002 notes-retreat cleanup lineage captured by `V01-REQ-DATA-002`, but it is not owned by the completed helper-shape migration. The work should avoid per-field over-splitting and address the remaining public schema candidates together. | n/a | New DATA work item under `V01-REQ-DATA-002` for request-side / generic container cleanup. |
| `enum-like / literal constraint` | N-019, N-030, N-034, N-045, N-046, N-051; residual vocabulary / literal notes from N-006, N-015, N-023, and N-029 | `new work item needed under existing requirement` | These are leftover value-set, literal, and usage-site vocabulary constraints after the M15 enum minimum. `V01-WORK-DATA-001` must not be reopened, but the remaining UC-002 cleanup can be handled as one follow-up under the existing `V01-REQ-DATA-002` cleanup-planning lineage. | n/a | New DATA work item under `V01-REQ-DATA-002` for remaining UC-002 enum / literal / usage-site vocabulary cleanup. |
| `numeric / default behavior` | N-011, N-017, N-022, N-024, N-025, N-028 | `new requirement needed` | The hidden contract is behavior-oriented: numeric ranges, defaults, omitted-value behavior, unknown-value behavior, fallback branches, and cross-response grouping. Existing DATA helper, enum, tagged-union, and MCP identity work do not own this product/spec capability. | n/a | New requirement plus work item, likely in the MCP/tool-contract domain, for request option behavior and response behavior constraints. This may be combined with the selector/support-matrix requirement if `V01-TASK-DATA-009-04` confirms one coherent contract owner. |
| `selector matrix / support matrix` | N-020, N-031, N-037, N-040, N-042 | `new requirement needed` | These rows need a support-matrix contract before YAML cleanup is meaningful. They are adjacent to MCP selector/ObjectRef behavior, not to helper-shape migration or tagged-union support. | n/a | New requirement plus work item, likely MCP-domain, for selector support matrices and object-dependent kind vocabulary. This may share the same requirement as numeric/default behavior if the future split keeps them together as MCP tool-contract rules. |
| `recursive / union structure` | N-009, N-044 | `new requirement needed` | These are type-system expressiveness gaps: an untagged union list and recursive `ObjectRef.parent`. `V01-WORK-DATA-010` covers tagged/discriminated unions only and should not be expanded to silently include untagged or recursive type semantics. | n/a | New DATA expressiveness requirement plus work item for recursive and untagged-union representation, or an explicit successor to V01-ADR-073 if later review chooses to broaden that decision. |
| `MCP identity / semantic reference` | No primary remaining `still-unowned` bucket. Previously reconciled identity candidates include N-032, N-035, N-038, N-039, N-043, N-047, and N-050; N-045 and N-048 residue stays in non-identity buckets. | `covered by existing work` | Primary identity and semantic-reference questions are already owned outside DATA cleanup by `V01-REQ-MCP-004` / `V01-WORK-MCP-004`. Creating duplicate DATA identity work would conflict with the prior split. | `V01-REQ-MCP-004` / `V01-WORK-MCP-004` | n/a |
| `human explanation / view-renderer note` | N-055, N-056 | `obsolete / no-action` | These rows are human explanation or view-renderer explanatory notes, not hidden machine-readable schema debt. No future artifact should be created unless new evidence identifies a concrete contract gap. | n/a | n/a |
| `obsolete / no-action candidate` | N-052, N-053, N-054, N-055, N-056, N-049; residual N-048 map-shape question unless `V01-WORK-MCP-004` later requires public schema | `obsolete / no-action` | N-052 to N-054 are secondary UC-001 rows, not remaining UC-002 cleanup. N-055 and N-056 are explanatory notes. N-049 is intentionally non-public render-context mapping per the inventory note. Residual N-048 should not create DATA work unless MCP identity work exposes a public map-shape need. | n/a | n/a |
| `other still-unowned` | None | `obsolete / no-action` | `V01-TASK-DATA-009-02` found no residual bucket after classification. No placeholder successor artifact should be created for an empty bucket. | n/a | n/a |

### Explicit no-action / obsolete outcomes

No future artifact should be created for these outcomes unless evidence changes:

- N-052, N-053, and N-054: secondary UC-001 enum-like rows, outside the remaining UC-002 notes-retreat cleanup scope.
- N-055 and N-056: human explanation / view-renderer notes without hidden machine-readable schema debt.
- N-049: `resolved_project.render_context` remains intentionally non-public mapping material for `analyze_impact`; no public schema cleanup is selected by this decision.
- Residual N-048 map-shape question: no DATA follow-up is selected here. If `V01-WORK-MCP-004` later needs public reference-index map schema, that MCP work owns the split.
- Empty `other still-unowned` bucket: no placeholder artifact.

N-024 is not selected as no-action in this decision. It remains in the `numeric / default behavior` bucket because the response fallback marker is coupled to request fallback behavior and should be resolved with that future contract owner rather than silently dropped.

### Explicit follow-up artifact candidates for V01-TASK-DATA-009-04

`V01-TASK-DATA-009-04` or later should create only the selected follow-up artifacts, or record an explicit close decision if the user chooses not to create them:

- New DATA work item under `V01-REQ-DATA-002`: request-side / generic container cleanup for N-002, N-004, N-007, N-008, N-012, N-016, N-018, and `TF-QUERY-RESULT`.
- New DATA work item under `V01-REQ-DATA-002`: remaining enum / literal / usage-site vocabulary cleanup for N-019, N-030, N-034, N-045, N-046, N-051, and residual N-006/N-015/N-023/N-029 vocabulary or literal notes.
- New MCP/tool-contract requirement plus work item: numeric range, default, omitted-value, unknown-value, fallback, and cross-response behavior constraints for N-011, N-017, N-022, N-024, N-025, and N-028.
- New MCP/tool-contract requirement plus work item, or the same requirement as the previous item if one owner is clearer: selector support matrices and object-dependent kind vocabulary for N-020, N-031, N-037, N-040, and N-042.
- New DATA expressiveness requirement plus work item: recursive and untagged-union representation for N-009 and N-044, kept separate from `V01-WORK-DATA-010` unless a later V01-ADR-073 review explicitly broadens that scope.

`V01-TASK-DATA-009-04` should record no-action close notes rather than create artifacts for N-049, residual N-048 map-shape questions, N-052 through N-056, and the empty `other still-unowned` bucket.

No follow-up DATA artifact should be created for primary MCP identity rows already covered by `V01-REQ-MCP-004` / `V01-WORK-MCP-004`, for N-052 through N-056, for N-049, for residual N-048 map-shape questions, or for the empty `other still-unowned` bucket.

### Verification note

This task performed successor outcome decision only.

No UC-002 YAML migration, fixture / golden regeneration, parser / renderer / validator / MCP implementation change, V01-ADR-073 implementation, V01-ADR-074 implementation, V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP identity implementation, UC-002 duplicate task QID / unresolved flow task fix, new requirement creation, new work item creation, or `V01-TASK-DATA-009-04` creation was performed.

`V01-WORK-DATA-009` was not marked `done`; it remains open because `T4` still needs to create the selected follow-up split artifacts or close with explicit no-action outcomes.

### Post-edit verification

Design Records MCP validation passed for:

- `V01-TASK-DATA-009-03`
- `V01-WORK-DATA-009`

Design Records MCP retrieval confirmed:

- `V01-TASK-DATA-009-03` exists with status `done`.
- `V01-WORK-DATA-009.tasks` includes `V01-TASK-DATA-009-01`, `V01-TASK-DATA-009-02`, and `V01-TASK-DATA-009-03`.
- `V01-WORK-DATA-009` status is `decision_pending`, not `done`.

Working tree check showed this task touched only:

- `docs/tasks/data/TASK-DATA-009-03-decide-remaining-uc-002-notes-retreat-successor-outcomes.md`
- `docs/work-items/data/WORK-DATA-009-remaining-uc-002-notes-retreat-classification.md`

Other modified / untracked files were already present in the worktree and were not edited by this task.

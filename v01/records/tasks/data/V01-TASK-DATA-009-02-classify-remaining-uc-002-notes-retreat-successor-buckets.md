# V01-TASK-DATA-009-02: Classify remaining UC-002 notes retreat successor buckets

- **id**: V01-TASK-DATA-009-02
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-DATA-009
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-DATA-009-01
- **outputs**:
  - Successor bucket classification for remaining UC-002 notes retreat candidates
  - Bucket-level rationale
  - Input evidence for later covered / obsolete / new-work decision

## Goal

Classify the `still-unowned` UC-002 notes retreat candidates from `V01-TASK-DATA-009-01` into stable successor buckets.

This task is classification only. It does not decide final successor actions, create follow-up work, or perform UC-002 cleanup implementation.

## Work

- Use `V01-TASK-DATA-009-01` as the input set for remaining `still-unowned` candidates.
- Preserve already separated ownership for `V01-WORK-DATA-006`, `V01-WORK-DATA-007`, `V01-WORK-DATA-008`, and `V01-WORK-DATA-010`.
- Group the remaining candidates into small successor buckets that can feed a later covered / obsolete / new-work decision.
- Identify candidates that are likely no-action / obsolete candidates for the next decision task.

## Included Scope

- Bucket classification for remaining UC-002 notes retreat candidates.
- Bucket-level rationale and non-coverage notes.
- Later decision type hints without making the later decision.

## Excluded Scope

- UC-002 YAML migration.
- Fixture / golden regeneration.
- Parser, renderer, validator, or MCP implementation changes.
- V01-ADR-073 tagged union implementation.
- V01-ADR-074 DAG TypeRef hint implementation.
- V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP identity implementation.
- UC-002 duplicate task QID / unresolved flow task fix.
- Creation of new requirements or new work items.
- Creation of `V01-TASK-DATA-009-03` or later tasks.
- Marking `V01-WORK-DATA-009` as `done`.
- Reopening M15, V01-WORK-DATA-001, V01-WORK-DATA-002, V01-WORK-DATA-003, V01-WORK-DATA-004, V01-WORK-DATA-005, V01-WORK-DATA-006, V01-WORK-DATA-007, V01-WORK-DATA-008, or V01-WORK-DATA-010.

## Done Condition

- Remaining `still-unowned` candidates from `V01-TASK-DATA-009-01` are classified into successor buckets.
- Each bucket records candidate IDs, rationale, non-coverage by existing successor work, and expected later decision type.
- Probable no-action / obsolete candidates are identified for the next decision task.
- No implementation, UC-002 YAML migration, fixture / golden regeneration, new requirement, new work item, or later task creation is performed.

## Verification

- Confirm only `V01-TASK-DATA-009-02` and `V01-WORK-DATA-009` documentation files are changed by this task.
- Confirm `V01-WORK-DATA-009.tasks` includes `V01-TASK-DATA-009-01` and `V01-TASK-DATA-009-02`.
- Confirm `V01-TASK-DATA-009-02` status is `done`.
- Run Design Records MCP validation for `V01-TASK-DATA-009-02` and `V01-WORK-DATA-009`.
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
- Accepted ADR title/status list under `docs/adr/`, especially the accepted/proposed boundary around V01-ADR-067, V01-ADR-069, V01-ADR-070, V01-ADR-073, V01-ADR-074, V01-ADR-078, V01-ADR-079, and V01-ADR-080
- `docs/investigations/data/INV-DATA-002-uc002-notes-retreat-inventory-and-m15-release-boundary-input.md`
- `docs/investigations/data/review-INV-DATA-002-m15-minimum-expressiveness-boundary.md`
- `docs/investigations/data/final-review-INV-DATA-002-m15-release-boundary.md`
- `docs/tasks/data/TASK-DATA-003-04-uc-002-model-response-helper-candidate-review.md`
- `docs/tasks/data/TASK-DATA-005-01-m15-deferred-item-inventory.md`
- `docs/tasks/data/TASK-DATA-005-02-deferred-ownership-classification.md`
- `docs/tasks/data/TASK-DATA-009-01-reconcile-remaining-uc-002-notes-retreat-candidates.md`
- `docs/work-items/data/WORK-DATA-006-helper-shape-migration.md`
- `docs/work-items/data/WORK-DATA-007-dag-asset-typeref-hint-render-support.md`
- `docs/work-items/data/WORK-DATA-008-uc-002-duplicate-task-qid-unresolved-flow-task.md`
- `docs/work-items/data/WORK-DATA-009-remaining-uc-002-notes-retreat-classification.md`
- `docs/work-items/data/WORK-DATA-010-tagged-union-and-discriminator-payload-support.md`

### Input summary from V01-TASK-DATA-009-01

`V01-TASK-DATA-009-01` reconciled the broad `V01-INV-DATA-002` inventory against existing successor ownership.

Already covered candidates are not reopened here:

- Helper-shape migration is covered by `V01-WORK-DATA-006`.
- DAG asset TypeRef hint render support is already separated to `V01-WORK-DATA-007`.
- UC-002 duplicate task QID / unresolved flow task repair is covered by `V01-WORK-DATA-008` and closed as stale / already resolved.
- Tagged union / discriminator payload candidates are covered for ownership by `V01-WORK-DATA-010`.
- Primary MCP identity candidates are covered by `V01-REQ-MCP-004` / `V01-WORK-MCP-004`.

The remaining `still-unowned` classification input from `V01-TASK-DATA-009-01` is:

- String-array and generic container cleanup: N-002, N-007, N-008, N-012, N-016, N-018, N-049, plus `TF-QUERY-RESULT`.
- Dict / map semantics: N-004, and residual public-schema question from N-048.
- Numeric/default/behavior constraints: N-011, N-017, N-022, N-024, N-025, N-028.
- Selector and object-kind matrices: N-020, N-031, N-037, N-040, N-042.
- Remaining enum-like / vocabulary cleanup not already migrated by `V01-TASK-DATA-001-04`: N-019, N-030, N-034, N-045, N-046, N-051.
- Untagged union / recursive constraints: N-009, N-044.
- Residual secondary notes from covered helper-shape rows: coverage vocabulary from N-006, literal `yaml` from N-023, optional/list vocabulary details from N-015 and N-029, and any public-schema map-shape issue left after N-048 identity classification.

### Bucket classification table

| bucket | candidate IDs | rationale | not already covered by V01-WORK-DATA-006 / 007 / 008 / 010 | expected later decision type |
|---|---|---|---|---|
| `request-side / generic container` | N-002, N-004, N-007, N-008, N-012, N-016, N-018, N-049, `TF-QUERY-RESULT`, residual public-schema question from N-048 | These are `any` / map / string-list / generic container surfaces where the open question is whether a public schema should replace note-retreated structure. They are likely handled together by a generic container/schema cleanup pass, with per-surface no-action checks for intentionally internal shapes. | `V01-WORK-DATA-006` handled selected response-local helper shapes, not generic request-side containers, response summary maps, task-file `query_result:any`, or render-context mapping. `V01-WORK-DATA-007` is only DAG TypeRef hint render support. `V01-WORK-DATA-008` is only the stale duplicate QID / unresolved flow issue. `V01-WORK-DATA-010` is discriminator/tagged-union work, not arbitrary string-list or map schema cleanup. | Later covered / obsolete / new-work decision. Likely split between public-schema cleanup candidates and explicit no-action for intentionally opaque internal mappings. |
| `enum-like / literal constraint` | N-019, N-030, N-034, N-045, N-046, N-051, residual vocabulary/literal notes from N-006, N-015, N-023, and N-029 | These are closed or semi-closed value sets, usage-site vocabulary constraints, or literal constraints left after the M15 enum minimum. They likely share a future enum/vocabulary cleanup pass, but some residuals may only need explicit no-action if the value set is local prose rather than public schema. | `V01-WORK-DATA-006` migrated helper object shapes but explicitly left vocabulary/literal/optional semantics outside that migration. `V01-WORK-DATA-007`, `V01-WORK-DATA-008`, and `V01-WORK-DATA-010` do not own ordinary enum or literal cleanup. `V01-WORK-DATA-010` owns discriminator payloads only, not general value-set constraints. | Later decision should classify which values become enums/literals, which remain usage-site notes, and which need no action. |
| `numeric / default behavior` | N-011, N-017, N-022, N-024, N-025, N-028 | These pair schema-like values with range, default, omitted-value, unknown-value, or cross-response behavior. They are more than simple enum creation because the hidden contract includes behavior. | `V01-WORK-DATA-006` is helper-shape migration only. `V01-WORK-DATA-007` is render hint support. `V01-WORK-DATA-008` is diagnostic blocker closure. `V01-WORK-DATA-010` is tagged-union support and does not define numeric ranges, defaults, or omitted-field response grouping behavior. | Later decision should decide whether these are spec/validation contracts, YAML cleanup only, or explicit no-action where behavior is intentionally prose. |
| `selector matrix / support matrix` | N-020, N-031, N-037, N-040, N-042 | These depend on valid combinations or object-dependent vocabularies rather than standalone field types. A future pass probably needs support-matrix ownership before YAML cleanup is meaningful. | `V01-WORK-DATA-006` does not cover tool-specific selector rules. `V01-WORK-DATA-007` and `V01-WORK-DATA-008` are unrelated. `V01-WORK-DATA-010` does not cover selector support matrices. MCP identity work may own identity semantics, but these rows preserve the remaining matrix/vocabulary rules after identity ownership is separated. | Later decision should determine whether the support matrix belongs to DATA validation, MCP public contract/spec, or explicit note-only behavior. |
| `recursive / union structure` | N-009, N-044 | These are structural expressiveness gaps: an untagged union list and a recursive `ObjectRef` parent. They are not the same as V01-ADR-073 discriminator payloads. | `V01-WORK-DATA-010` covers tagged/discriminated union candidates, not untagged unions or recursive self-reference. `V01-WORK-DATA-006` helper migration cannot express these without a deeper type-system decision. `V01-WORK-DATA-007` and `V01-WORK-DATA-008` are unrelated. | Later decision should decide whether to create a separate DATA expressiveness requirement, defer explicitly, or mark as no-action for v1-style documentation-only constraints. |
| `MCP identity / semantic reference` | None as a primary remaining `still-unowned` bucket; N-045 remains in `enum-like / literal constraint`, and residual N-048 remains in `request-side / generic container` only for any public map-shape question. | `V01-TASK-DATA-009-01` already classified primary identity candidates as covered by `V01-REQ-MCP-004` / `V01-WORK-MCP-004`. The remaining pieces are vocabulary or map-shape residue, not identity implementation. | `V01-WORK-DATA-006` / `007` / `008` / `010` do not own MCP identity, but the primary identity work is already outside V01-WORK-DATA-009's DATA successor split. | Later decision should avoid creating duplicate DATA identity work; only non-identity residue should be considered. |
| `human explanation / view-renderer note` | None from the remaining `still-unowned` input; N-055 and N-056 were already reconciled as obsolete / no-action by `V01-TASK-DATA-009-01`. | No remaining machine-readable hidden schema debt is classified here. Human explanation and view-renderer notes should not become DATA implementation work unless a later source introduces a concrete contract gap. | Not covered by the listed work items because no implementation ownership is needed. | Later decision should preserve explicit no-action outcomes for N-055 and N-056 rather than creating work. |
| `obsolete / no-action candidate` | Likely: N-049 if render-context mapping remains intentionally non-public; residual N-048 if identity work does not require a public map schema; N-024 if the response fallback marker is intentionally a single literal note. Already reconciled obsolete/no-action candidates: N-052, N-053, N-054, N-055, N-056. | These need an explicit later decision because no-action should be recorded deliberately, not inferred from absence of implementation. | They are not covered by `V01-WORK-DATA-006` / `007` / `008` / `010` because the likely outcome is no DATA cleanup implementation, not existing successor ownership. | Later covered / obsolete / new-work decision should explicitly mark no-action / obsolete outcomes. |
| `other still-unowned` | None after this classification pass. | The preferred buckets cover every remaining `still-unowned` input from `V01-TASK-DATA-009-01`. | No residual bucket is needed. | No later decision type beyond confirming the table remains complete. |

### Bucket rationale notes

- `request-side / generic container` keeps string-list, map, task-file `query_result:any`, and residual public-map questions together because they require the same future question: is this a public schema surface or intentionally opaque/internal?
- `enum-like / literal constraint` deliberately separates simple vocabulary cleanup from `numeric / default behavior`; default and omitted-value behavior likely needs spec or validation ownership, not only enum creation.
- `selector matrix / support matrix` keeps object-dependent `kind` vocabularies with selector valid-combination rows because both require a matrix-style contract before implementation.
- `recursive / union structure` stays separate from `V01-WORK-DATA-010` because V01-ADR-073 is discriminator-oriented and does not settle untagged union or recursive TypeRef representation.
- `MCP identity / semantic reference` is present as a negative classification bucket to prevent future DATA work from re-creating already separated MCP identity ownership.

### Probable no-action / obsolete candidates for the next task

The next decision task should probably carry explicit no-action / obsolete outcomes for:

- N-052, N-053, and N-054, because `V01-TASK-DATA-009-01` found them to be secondary UC-001 rows, not remaining UC-002 cleanup candidates.
- N-055 and N-056, because they are human explanation / view-renderer notes without hidden machine-readable schema debt.
- N-049, if the render-context mapping remains intentionally non-public and only material for `analyze_impact`.
- Residual N-048, if MCP identity work does not require a public schema for the reference index map shape.
- N-024, if the response fallback marker is intentionally only the `file` fallback case and no broader public response enum is needed.

These are not decided here; they are input for the later covered / obsolete / new-work task.

### Verification note

This task performed bucket classification only.

No UC-002 YAML migration, fixture / golden regeneration, parser / renderer / validator / MCP implementation change, V01-ADR-073 implementation, V01-ADR-074 implementation, V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP identity implementation, UC-002 duplicate task QID / unresolved flow task fix, new requirement creation, new work item creation, or `V01-TASK-DATA-009-03` creation was performed.

`V01-WORK-DATA-009` was not marked `done`; it remains open because later task(s) still need to decide concrete successor actions or explicit no-action / obsolete outcomes.

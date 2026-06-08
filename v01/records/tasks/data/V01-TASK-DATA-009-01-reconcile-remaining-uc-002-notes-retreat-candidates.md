# V01-TASK-DATA-009-01: Reconcile remaining UC-002 notes retreat candidates

- **id**: V01-TASK-DATA-009-01
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-DATA-009
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-DATA-003-04
  - V01-TASK-DATA-005-01
  - V01-TASK-DATA-005-02
- **outputs**:
  - Remaining UC-002 notes retreat candidate reconciliation
  - Covered / obsolete / still-unowned candidate table
  - Input evidence for successor bucket classification

## Goal

Reconcile the remaining UC-002 notes retreat candidates after the known successor split has separated helper-shape migration, DAG TypeRef hint render support, tagged union support, MCP identity-related work, and the UC-002 duplicate task QID / unresolved flow task issue.

This task is inventory reconciliation only. It must not perform direct UC-002 cleanup implementation.

## Work

- Review `V01-INV-DATA-002`, `V01-TASK-DATA-003-04`, `V01-TASK-DATA-005-01`, and `V01-TASK-DATA-005-02` as the primary candidate sources.
- Confirm which UC-002 notes retreat candidates are already covered by existing successor work items.
- Confirm which candidates became obsolete or no-action after the completed helper/model render and signature-policy work.
- Identify candidates that remain unowned and need later bucket classification.
- Record enough evidence for the next V01-WORK-DATA-009 task to classify successor buckets without reopening prior completed work.

## Included Scope

- Candidate reconciliation by N-id or equivalent source note identifier.
- Covered / obsolete / still-unowned classification.
- Cross-reference to existing successor work items where already separated.
- Evidence notes for later bucket classification.

## Excluded Scope

- UC-002 YAML migration.
- Fixture / golden regeneration.
- Parser, renderer, validator, or MCP implementation changes.
- V01-ADR-073 tagged union implementation.
- V01-ADR-074 DAG TypeRef hint implementation.
- V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP identity implementation.
- UC-002 duplicate task QID / unresolved flow task fix.
- Creating additional requirements or work items in this task.
- Reopening M15, V01-WORK-DATA-001, V01-WORK-DATA-002, V01-WORK-DATA-003, V01-WORK-DATA-004, V01-WORK-DATA-005, V01-WORK-DATA-006, V01-WORK-DATA-007, V01-WORK-DATA-008, or V01-WORK-DATA-010.

## Done Condition

- Remaining UC-002 notes retreat candidates are reconciled against existing successor ownership.
- Each reviewed candidate is classified as covered, obsolete / no-action, or still-unowned.
- Still-unowned candidates are preserved as input for later successor bucket classification.
- No implementation, UC-002 YAML migration, fixture / golden regeneration, or new successor artifact creation is performed.

## Verification

- Confirm only documentation files for V01-WORK-DATA-009 / V01-TASK-DATA-009-01 are changed.
- Confirm existing successor work items are referenced rather than reopened.
- Confirm this task does not decide final implementation buckets beyond reconciliation status.
- Confirm V01-WORK-DATA-009 remains a cleanup-planning work item, not direct cleanup implementation.

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
- Accepted ADR title/status list under `docs/adr/`
- `docs/investigations/data/INV-DATA-002-uc002-notes-retreat-inventory-and-m15-release-boundary-input.md`
- `docs/investigations/data/review-INV-DATA-002-m15-minimum-expressiveness-boundary.md`
- `docs/investigations/data/final-review-INV-DATA-002-m15-release-boundary.md`
- `docs/tasks/data/TASK-DATA-001-04-implement-enum-minimum-atomic-migration.md`
- `docs/tasks/data/TASK-DATA-003-04-uc-002-model-response-helper-candidate-review.md`
- `docs/tasks/data/TASK-DATA-005-01-m15-deferred-item-inventory.md`
- `docs/tasks/data/TASK-DATA-005-02-deferred-ownership-classification.md`
- `docs/tasks/data/TASK-DATA-005-03-create-successor-split.md`
- `docs/tasks/data/TASK-DATA-005-04-sync-links-and-close-triage.md`
- `docs/tasks/data/TASK-DATA-006-01-uc-002-helper-shape-migration-set.md`
- `docs/tasks/data/TASK-DATA-006-04-close-and-follow-up-split.md`
- `docs/tasks/data/TASK-DATA-008-01-reproduce-and-localize-uc-002-duplicate-task-qid.md`
- `docs/tasks/data/TASK-DATA-010-01-adr-073-acceptance-split-review.md`
- `docs/work-items/data/WORK-DATA-005-triage-m15-deferred-follow-ups.md`
- `docs/work-items/data/WORK-DATA-006-helper-shape-migration.md`
- `docs/work-items/data/WORK-DATA-007-dag-asset-typeref-hint-render-support.md`
- `docs/work-items/data/WORK-DATA-008-uc-002-duplicate-task-qid-unresolved-flow-task.md`
- `docs/work-items/data/WORK-DATA-009-remaining-uc-002-notes-retreat-classification.md`
- `docs/work-items/data/WORK-DATA-010-tagged-union-and-discriminator-payload-support.md`
- `docs/requirements/mcp/REQ-MCP-004-mcp-semantic-identity-and-state-machine-identity-follow-up.md`
- `docs/work-items/mcp/WORK-MCP-004-mcp-semantic-identity-and-state-machine-identity-follow-up.md`
- Focused current UC-002 YAML reads for helper-shape and enum status only; no UC-002 YAML was edited.

### Reconciliation rules

Classification below is inventory reconciliation, not final successor bucket design.

- `covered` means an existing work item or completed task already owns the primary candidate aspect.
- `obsolete / no-action` means no V01-WORK-DATA-009 successor action is needed because the row is outside UC-002 cleanup scope or is not hidden machine-readable schema debt.
- `still-unowned` means the candidate remains input for later V01-WORK-DATA-009 successor bucket classification.

Where an V01-INV-DATA-002 N-id mixes a primary helper-shape issue with secondary vocabulary, literal, optional, identity, or tagged-union notes, the table classifies the primary N-id and preserves the residual note in the evidence column.

### UC-002 notes retreat candidate reconciliation

| candidate | classification | owning artifact / later input | evidence |
|---|---|---|---|
| N-001 | covered | `V01-WORK-DATA-010` | Discriminated `analyze_impact_request.change` is part of the tagged-union / discriminator payload successor. |
| N-002 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Request-side module string-array container remains `any + note`; no existing successor owns generic request-side string-list cleanup. |
| N-003 | covered | `V01-WORK-DATA-010` | Discriminated `analyze_impact_response.change` is part of the tagged-union successor. |
| N-004 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Dict summary key semantics remain separate from helper-shape migration and tagged-union work. |
| N-005 | covered | `V01-WORK-DATA-006`; residual tagged-union note aligns with `V01-WORK-DATA-010` | The response-local impact entry shape was selected and migrated by V01-WORK-DATA-006. `suggested_fixes` kind-dependent payload remains outside helper migration and is tagged-union-adjacent. |
| N-006 | covered | `V01-WORK-DATA-006`; residual vocabulary note remains later cleanup input | The response-local coverage helper shape was selected and migrated by V01-WORK-DATA-006. Coverage vocabulary constraints were explicitly left as note / enum follow-up. |
| N-007 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Assumptions string-array container remains a generic `any + note` cleanup candidate. |
| N-008 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Truncated-reasons string-array container remains a generic `any + note` cleanup candidate. |
| N-009 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Untagged union list of `SourceLocation` or `ObjectRef` is not covered by the discriminator-oriented tagged-union successor. |
| N-010 | covered | `V01-TASK-DATA-001-04` / `V01-WORK-DATA-001` | `get_reference_tree_request.direction` was migrated to the `reference_tree_direction` enum in the M15 enum minimum. |
| N-011 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Numeric range and `invalid_depth` behavior are not covered by enum, helper-shape, or tagged-union successors. |
| N-012 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Reference-kind filter string-array remains a request-side / generic container candidate. |
| N-013 | covered | `V01-TASK-DATA-001-04` / `V01-WORK-DATA-001` | `get_reference_tree_response.direction` was migrated to the `reference_tree_direction` enum in the M15 enum minimum. |
| N-014 | covered | `V01-WORK-DATA-006` | Response-local reference-tree node entry shape was selected and migrated by V01-WORK-DATA-006. |
| N-015 | covered | `V01-WORK-DATA-006`; residual reference vocabulary notes remain later cleanup input | Response-local reference-tree edge entry shape was selected and migrated by V01-WORK-DATA-006; kind / direction value constraints remain note-based elsewhere. |
| N-016 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Response truncated-reasons string-array remains a generic `any + note` cleanup candidate. |
| N-017 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Optional direction enum plus default behavior was explicitly excluded from the initial enum migration. |
| N-018 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Reference-kind filter string-array remains a request-side / generic container candidate. |
| N-019 | still-unowned | Later V01-WORK-DATA-009 bucket classification | `get_references_response.direction` was explicitly not migrated by `V01-TASK-DATA-001-04`. |
| N-020 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Selector valid-combination matrix is tool-specific behavior, not owned by helper, tagged-union, or MCP identity split. |
| N-021 | covered | `V01-WORK-DATA-010` | Kind-specific `get_signature_response.signature` payload is a tagged-union successor candidate. |
| N-022 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Fallback enum plus default / branch behavior remains note-based. |
| N-023 | covered | `V01-WORK-DATA-006`; residual literal note remains later cleanup input | Response-local snippet helper shape was selected and migrated by V01-WORK-DATA-006. The `language: yaml` literal constraint remains outside that migration. |
| N-024 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Response fallback marker enum remains note-based and only partially specified. |
| N-025 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Detail enum plus default / unknown-value behavior remains note-based. |
| N-026 | covered | `V01-WORK-DATA-010` | Kind-specific `inspect_response.signature` payload is a tagged-union successor candidate. |
| N-027 | covered | `V01-WORK-DATA-010` | Kind-specific `inspect_response.members` payload is a tagged-union successor candidate. |
| N-028 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Cross-response grouping behavior based on omitted `api_table_id` is not covered by existing successors. |
| N-029 | covered | `V01-WORK-DATA-006`; residual optional semantics remain later cleanup input | Nested list-endpoints table / section / endpoint shapes were selected and migrated by V01-WORK-DATA-006. |
| N-030 | still-unowned | Later V01-WORK-DATA-009 bucket classification | `list_objects_request.object` remains an enum-like filter outside the initial enum migration. |
| N-031 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Object-dependent kind vocabulary remains note-based and needs later classification. |
| N-032 | covered | `V01-WORK-MCP-004` / `V01-REQ-MCP-004` | FileID semantics belong to the MCP semantic identity / state-machine identity successor split. |
| N-033 | covered | `V01-WORK-DATA-006`; residual identity semantics align with `V01-WORK-MCP-004` | Response-local list object shape was selected and migrated by V01-WORK-DATA-006; identity semantics remain outside DATA helper migration. |
| N-034 | still-unowned | Later V01-WORK-DATA-009 bucket classification | MCP error-code vocabulary remains note-based and was not part of the initial enum migration. |
| N-035 | covered | `V01-WORK-MCP-004` / `V01-REQ-MCP-004` | QualifiedID / actor global ID / synthetic ID selector semantics belong to MCP identity work. |
| N-036 | covered | `V01-TASK-DATA-001-04` / `V01-WORK-DATA-001` | `object_selector.object` was migrated to the `mcp_object_type` enum in the M15 enum minimum. |
| N-037 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Object-dependent selector kind vocabulary remains note-based. |
| N-038 | covered | `V01-WORK-MCP-004` / `V01-REQ-MCP-004` | FileID selector semantics belong to MCP identity work. |
| N-039 | covered | `V01-WORK-MCP-004` / `V01-REQ-MCP-004` | File-local object ID semantics belong to MCP identity work. |
| N-040 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Tool-specific valid selector field combinations remain unowned after identity split. |
| N-041 | covered | `V01-TASK-DATA-001-04` / `V01-WORK-DATA-001` | `object_ref.object` was migrated to the `mcp_object_type` enum in the M15 enum minimum. |
| N-042 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Object-specific `object_ref.kind` vocabulary remains note-based. |
| N-043 | covered | `V01-WORK-MCP-004` / `V01-REQ-MCP-004` | ObjectRef QualifiedID / synthetic ID / file-local identity variants belong to MCP identity work. |
| N-044 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Recursive `object_ref.parent` is not covered by current helper-shape, enum, or tagged-union successors. |
| N-045 | still-unowned | Later V01-WORK-DATA-009 bucket classification | `reference.kind` is primarily a closed vocabulary row in V01-TASK-DATA-003-04. V01-INV-DATA-002 notes MCP identity adjacency, but no existing successor owns the vocabulary cleanup itself. |
| N-046 | still-unowned | Later V01-WORK-DATA-009 bucket classification | `reference.direction` was explicitly excluded from the initial enum migration. |
| N-047 | covered | `V01-WORK-MCP-004` / `V01-REQ-MCP-004` | Semantic object registry identity semantics belong to MCP identity work. |
| N-048 | covered | `V01-WORK-MCP-004` / `V01-REQ-MCP-004`; residual map-shape question remains later cleanup input if public schema is required | Reference index key/value semantics are identity-related, but arbitrary map representation is not solved by that split. |
| N-049 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Render-context mapping is still an opaque / generic container candidate; later classification may decide it is intentionally non-public. |
| N-050 | covered | `V01-WORK-MCP-004` / `V01-REQ-MCP-004` | `source_location.file` FileID semantics belong to MCP identity work. |
| N-051 | still-unowned | Later V01-WORK-DATA-009 bucket classification | Usage-site-dependent list element vocabularies remain note-based. |
| N-052 | obsolete / no-action | No V01-WORK-DATA-009 successor | Secondary UC-001 `cart.status` row, not a remaining UC-002 cleanup candidate. |
| N-053 | obsolete / no-action | No V01-WORK-DATA-009 successor | Secondary UC-001 `order.status` row, not a remaining UC-002 cleanup candidate. |
| N-054 | obsolete / no-action | No V01-WORK-DATA-009 successor | Secondary UC-001 `payment_event.result` row, not a remaining UC-002 cleanup candidate. |
| N-055 | obsolete / no-action | No V01-WORK-DATA-009 successor | Human-facing actor explanation note; no hidden machine-readable schema debt. |
| N-056 | obsolete / no-action | No V01-WORK-DATA-009 successor | View / renderer explanation note; no hidden machine-readable schema debt. |
| TF-QUERY-RESULT | still-unowned | Later V01-WORK-DATA-009 bucket classification | Equivalent non-N-id source from V01-TASK-DATA-003-04 / V01-TASK-DATA-006-01: eight UC-002 MCP task-file `query_service.returns.model:any` to `build_response.params[].model:any` patterns remain deferred by the DATA-004 / V01-REQ-DATA-003 private-helper params policy. |

### Known successor split check

- Helper-shape migration is not reopened. The selected response-local helper-shape candidates are covered by `V01-WORK-DATA-006`, which is already `done`.
- Tagged union / discriminator payload support is not implemented here. The tagged-union candidates are covered for ownership by `V01-WORK-DATA-010`, which remains `not_started`.
- DAG asset TypeRef hint render support has no UC-002 N-id in `V01-INV-DATA-002`; the deferred V01-ADR-074 bucket is already separated to `V01-WORK-DATA-007`.
- MCP identity-related work is not implemented here. Identity candidates are covered for ownership by `V01-REQ-MCP-004` / `V01-WORK-MCP-004`, which remain outside DATA cleanup implementation.
- The UC-002 duplicate task QID / unresolved flow task issue has no V01-INV-DATA-002 N-id. It is covered by `V01-WORK-DATA-008`, which closed it as a stale follow-up already resolved by `V01-WORK-RESOLVE-001` / ADR-058-aligned resolver behavior.

### Still-unowned input for later V01-WORK-DATA-009 bucket classification

The remaining still-unowned inventory should be carried forward as classification input only:

- String-array and generic container cleanup: N-002, N-007, N-008, N-012, N-016, N-018, N-049, plus TF-QUERY-RESULT.
- Dict / map semantics: N-004, and any residual public-schema question from N-048.
- Numeric/default/behavior constraints: N-011, N-017, N-022, N-024, N-025, N-028.
- Selector and object-kind matrices: N-020, N-031, N-037, N-040, N-042.
- Remaining enum-like / vocabulary cleanup not already migrated by `V01-TASK-DATA-001-04`: N-019, N-030, N-034, N-045, N-046, N-051.
- Untagged union / recursive constraints: N-009, N-044.
- Residual secondary notes from covered helper-shape rows: coverage vocabulary from N-006, literal `yaml` from N-023, optional/list vocabulary details from N-015 and N-029, and any public-schema map-shape issue left after N-048 identity classification.

### Verification note

This task performed inventory reconciliation only.

No UC-002 YAML migration, fixture / golden regeneration, parser / renderer / validator / MCP implementation change, V01-ADR-073 implementation, V01-ADR-074 implementation, V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP identity implementation, UC-002 duplicate task QID / unresolved flow task fix, or new requirement / work item / task creation was performed.

`V01-WORK-DATA-009` was not marked `done`; it remains the parent cleanup-planning work item for later bucket classification.

# PRODUCT-TASK-SPEC-012-08: Apply DRMCP app-local handoff

- **id**: PRODUCT-TASK-SPEC-012-08
- **status**: done
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-07
- **outputs**:
  - `drmcp/records/spec/design-records-mcp/overview.md`
  - `drmcp/records/spec/design-records-mcp/namespace-scanning.md`
  - `drmcp/records/spec/design-records-mcp/resolver.md`
  - `drmcp/records/spec/design-records-mcp/schema/id-normalization.md`
  - `drmcp/records/spec/design-records-mcp/schema/overview.md`

## Goal

Complete the DRMCP handoff for operational content removed from PRODUCT specifications.

## Work

- Use the T01 manifest to identify concrete DRMCP-owned statements requiring app-local preservation.
- Update only app-local files needed to prevent contract loss.
- Reference PRODUCT-owned Design Records semantics instead of restating them.
- Keep parser, scanning, resolver, schema, tool, UI, and projection behavior under DRMCP.
- Do not redesign DRMCP or expand its current specification hierarchy.
- Keep this handoff separate from BPDSL work and broad ref synchronization.

## Done condition

- Every DRMCP-owned statement removed from PRODUCT has a retained app-local owner or deletion rationale.
- DRMCP specs point to PRODUCT-owned generic contracts where appropriate.
- PRODUCT semantics are not duplicated as DRMCP authority.
- No unrelated DRMCP redesign occurs.

## Verification

- Compare handoff items with `PRODUCT-INV-SPEC-005` and T01.
- Confirm each app-local update is required by the accepted ownership split.
- Confirm no BPDSL content is changed.
- Confirm downstream ref cleanup remains deferred.

## Evidence

### Baseline validation

```
python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color
[strict]  All 47 file(s) OK.
exit 0

python -X utf8 product/src/tools/validate_spec.py drmcp/records/spec/design-records-mcp --strict --no-color
[strict]  All 30 file(s) OK.
exit 0
```

### Files changed

| file | change |
|---|---|
| `drmcp/records/spec/design-records-mcp/resolver.md` | narrow_update: 2 occurrences of `spec:product.concepts.traceability.resolve_and_validation` → `spec:product.design_records.traceability.resolve_and_validation` |
| `drmcp/records/spec/design-records-mcp/schema/discovery.md` | narrow_update: 2 occurrences of `spec:product.concepts.repository_layout.record_discovery_paths` → `spec:product.design_records.repository_layout.record_discovery_paths` |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-08-apply-drmcp-app-local-handoff.md` | status: not_started → done; Evidence added |

### PRODUCT pointers corrected

| file | old ref | new ref |
|---|---|---|
| `resolver.md` | `spec:product.concepts.traceability.resolve_and_validation` | `spec:product.design_records.traceability.resolve_and_validation` |
| `schema/discovery.md` | `spec:product.concepts.repository_layout.record_discovery_paths` | `spec:product.design_records.repository_layout.record_discovery_paths` |

### Disposition map

All 34 mandatory concerns covered.

**Namespace and discovery (9)**

| concern | source evidence | current app-local owner | disposition | file update, if any | rationale |
|---|---|---|---|---|---|
| 1. DRMCP architecture and tool list removed from PRODUCT app-namespace text | T03 Evidence; `spec:product.design_records.namespace_model.app_namespaces` boundary | `spec:drmcp.design_records_mcp.overview`; `spec:drmcp.design_records_mcp.tools.overview` | retain_existing_owner | none | overview.md carries the tool boundary table and record scope table; tools/overview.md carries the full tool set with P0/P1 classification, contract class, and Topics |
| 2. Namespace scanning | T04 Evidence; `spec:product.design_records.repository_layout.record_discovery_paths` T08 handoff | `spec:drmcp.design_records_mcp.namespace_scanning` | retain_existing_owner | none | namespace-scanning.md defines multi-root auto-discovery and single-root `--records-root` mode with kind-level prefix application table |
| 3. namespace_prefix derivation from records_root | T04 Evidence; `spec:product.design_records.repository_layout.record_discovery_paths` T08 handoff | `spec:drmcp.design_records_mcp.namespace_scanning` | retain_existing_owner | none | namespace-scanning.md defines `namespace_prefix = strings.ToUpper(appNamespaceDir) + "-"` with multi-root examples |
| 4. Multi-root and single-root scanning behavior | T04 Evidence; `spec:product.design_records.repository_layout.record_discovery_paths` T08 handoff | `spec:drmcp.design_records_mcp.namespace_scanning` | retain_existing_owner | none | namespace-scanning.md defines both modes: multi-root auto-discovers all `*/records/` directories; single-root uses `--records-root <path>` |
| 5. Record discovery inclusion conditions | T04 Evidence; `spec:product.design_records.repository_layout.record_discovery_paths` T08 handoff | `spec:drmcp.design_records_mcp.schema.discovery` | narrow_update | `schema/discovery.md`: stale PRODUCT ref updated | discovery.md already owns DRMCP-specific inclusion conditions per kind; stale ref corrected to `spec:product.design_records.repository_layout.record_discovery_paths` |
| 6. Tool-owned scanning filters | T04 Evidence; `spec:product.design_records.repository_layout.record_discovery_paths` T08 handoff | `spec:drmcp.design_records_mcp.schema.discovery` | retain_existing_owner | none | discovery.md defines DRMCP-specific filters: spec kind requires `design_record.id` + `design_record.kind` in YAML front matter; other kinds use path-pattern only |
| 7. Existing-artifact UI or MCP attribution projection | T03 Evidence; `spec:product.design_records.namespace_model.existing_artifacts` boundary | none | delete_with_rationale | none | Attribution facts are owned by `spec:product.brewprint.compatibility`. No accepted DRMCP app-local contract for attribution projection behavior exists. |
| 8. DRMCP MCP-domain subdomain example and tool concern | T03 Evidence; `spec:product.design_records.namespace_model.subdomain_model` boundary | none | delete_with_rationale | none | Domain assignments (MCP, SPEC) are observable from existing DRMCP record IDs. PRODUCT generic subdomain-model owns the semantics. No normative DRMCP subdomain catalog spec is required. |
| 9. Provenance of path-pattern content formerly held by DRMCP | T04 Evidence; `spec:product.design_records.repository_layout.record_discovery_paths` T08 handoff | `spec:drmcp.design_records_mcp.schema.discovery` | narrow_update | `schema/discovery.md`: stale PRODUCT ref updated | discovery.md already carries the provenance note ("Phase 2 relocation per PRODUCT-WORK-SPEC-004"); stale ref in that note corrected by the same edit as concern 5 |

**Resolve, validation, and indexing (11)**

| concern | source evidence | current app-local owner | disposition | file update, if any | rationale |
|---|---|---|---|---|---|
| 10. Index construction behavior | T04+T05 Evidence; `spec:product.design_records.artifact_model.traceability_boundary` T08 handoff | `spec:drmcp.design_records_mcp.namespace_scanning`; `spec:drmcp.design_records_mcp.schema.discovery`; `spec:drmcp.design_records_mcp.schema.record_source` | retain_existing_owner | none | namespace-scanning.md covers multi-root index building; discovery.md covers inclusion conditions per kind; record-source.md covers metadata source per record kind |
| 11. Resolver implementation behavior | T04+T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` boundary | `spec:drmcp.design_records_mcp.resolver` | narrow_update | `resolver.md`: stale PRODUCT ref updated | resolver.md already owns public tool name, unsupported inputs, and lookup-source vs. record-kind boundary; stale ref corrected to `spec:product.design_records.traceability.resolve_and_validation` |
| 12. Validation implementation behavior | T05 Evidence; `spec:product.design_records.artifact_model.traceability_boundary` T08 handoff | `spec:drmcp.design_records_mcp.tools.validate_records`; `spec:drmcp.design_records_mcp.schema.diagnostics` | retain_existing_owner | none | validate-records.md defines the full `validate_records` tool contract including valid status values per kind; diagnostics.md defines all diagnostic categories, severities, and MVP exclusions |
| 13. Resolver request shape | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.tools.resolve_reference` | retain_existing_owner | none | resolve-reference.md ## Request defines the full supported input forms table and input field contract |
| 14. Resolver response shape and status vocabulary | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.tools.resolve_reference`; `spec:drmcp.design_records_mcp.schema.diagnostics` | retain_existing_owner | none | resolve-reference.md ## Response defines response shape and status vocabulary (`resolved`/`unresolved`/`unsupported`); diagnostics.md defines resolve_reference response categories |
| 15. Diagnostic categories and severities | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.schema.diagnostics` | retain_existing_owner | none | diagnostics.md defines all four diagnostic category groups (resolve_reference, get_records, authoring transaction, validate_records) with full category names and severities |
| 16. Duplicate and unresolved diagnostic presentation | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.schema.diagnostics` | retain_existing_owner | none | diagnostics.md defines `duplicate_id`, `ambiguous_reference`, `unresolved_reference`, and `unresolved_source_ref` in the validate_records category group |
| 17. Parser behavior | T05 Evidence; `spec:product.design_records.artifact_model.traceability_boundary` T08 handoff | `spec:drmcp.design_records_mcp.schema.record_source` | retain_existing_owner | none | record-source.md defines metadata source per record kind; compatibility gap (YAML front matter model vs. new H1-adjacent contract) deferred to DRMCP-WORK-SPEC-001 |
| 18. Persistence and indexing implementation | T04+T05 Evidence; `spec:product.design_records.artifact_model.traceability_boundary` T08 handoff | `spec:drmcp.design_records_mcp.overview` | retain_existing_owner | none | overview.md establishes DRMCP as the index owner; no dedicated persistence spec exists — persistence implementation is not a current normative contract |
| 19. Record/source projection and retrieval response behavior | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.tools.overview` | retain_existing_owner | none | tools/overview.md defines common response conventions with JSON examples covering all tool responses including retrieval |
| 20. UI and tool API behavior | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.tools.overview` and all tool specs | retain_existing_owner | none | tools/overview.md defines error codes and request/response conventions; individual tool specs define per-tool APIs |

**Writer and authoring behavior (8)**

| concern | source evidence | current app-local owner | disposition | file update, if any | rationale |
|---|---|---|---|---|---|
| 21. MCP writer tools | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.tools.overview`; `spec:drmcp.design_records_mcp.tools.authoring_transaction_model`; `spec:drmcp.design_records_mcp.tools.propose_record_create`; `spec:drmcp.design_records_mcp.tools.propose_record_update`; `spec:drmcp.design_records_mcp.tools.get_proposed_write`; `spec:drmcp.design_records_mcp.tools.accept_proposed_write`; `spec:drmcp.design_records_mcp.tools.discard_proposed_write` | retain_existing_owner | none | Current tool contracts retain propose/get/accept/discard behavior; see compatibility gap for P0/MVP classification contradiction between tools/overview.md and mvp-scope.md |
| 22. Proposal/dry-run diff behavior | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.tools.authoring_transaction_model` | retain_existing_owner | none | authoring-transaction-model.md defines diff_mode (`summary`/`patch`/`none`) and the diff response shape including per-file change entries |
| 23. Explicit propose/accept confirmation flow | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.tools.authoring_transaction_model`; `spec:drmcp.design_records_mcp.tools.accept_proposed_write` | retain_existing_owner | none | authoring-transaction-model.md defines the two-step Propose→Accept flow and proposal lifecycle states; accept-proposed-write.md defines the accept tool contract |
| 24. Conflict, stale-target, and ID-collision handling | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.schema.diagnostics`; `spec:drmcp.design_records_mcp.tools.accept_proposed_write` | retain_existing_owner | none | diagnostics.md defines `proposal_stale`, `target_changed`, `id_collision`; accept-proposed-write.md defines `written: false` conditions and accept-time checks |
| 25. Formatting preservation behavior | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.tools.propose_record_update`; `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | retain_existing_owner | none | `metadata_fields_replace` preserves all unspecified existing metadata fields; `metadata_block_replace` for spec preserves unknown or auxiliary YAML front matter fields; `named_section_replace` limits replacement to the selected section only; no-op detection compares final persisted result with current file. Preservation is scoped to each operation's defined semantics — this is not a global byte-preservation contract. |
| 26. Permission and filesystem-write boundary | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.tools.overview`; `spec:drmcp.design_records_mcp.responsibility_boundary` | retain_existing_owner | none | tools/overview.md ## Authoring write boundary defines MVP write scope; responsibility-boundary.md defines boundary against filesystem tools |
| 27. Authoring proposal and transaction behavior | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.tools.authoring_transaction_model`; `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema` | retain_existing_owner | none | authoring-transaction-model.md defines proposal lifecycle and Propose→Accept flow; authoring-transaction-schema.md defines proposal model, body cache model, metadata block replacement target, and section selector model |
| 28. Agent/tool interface behavior | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` excluded impl. | `spec:drmcp.design_records_mcp.tools.overview` and all individual tool specs | retain_existing_owner | none | tools/overview.md defines common request/response conventions; all tool specs define per-tool request/response/error contracts |

**Operational projections (6)**

Concerns 29–33 are retained as app-local MVP exclusions. Retention does not activate or require these features; the app-local exclusion record is the adequate owner. Concern 34 has a deletion rationale and remains unadopted.

| concern | source evidence | current app-local owner | disposition | file update, if any | rationale |
|---|---|---|---|---|---|
| 29. Orphan requirement/work-item/task diagnostics | T05 Evidence; `spec:product.design_records.traceability.resolve_and_validation` ## Resolve and validation boundary | `spec:drmcp.design_records_mcp.mvp_scope`; `spec:drmcp.design_records_mcp.schema.diagnostics` | retain_existing_owner | none | mvp-scope.md explicitly excludes "Orphan requirement / orphan work item / orphan task diagnostics"; diagnostics.md Diagnostic policy confirms the MVP exclusion |
| 30. Work-item progress projection from task status | T05 Evidence; same | `spec:drmcp.design_records_mcp.mvp_scope` | retain_existing_owner | none | mvp-scope.md explicitly excludes "Deriving work item progress from task status" |
| 31. Workflow traversal/tree/graph queries | T05 Evidence; same | `spec:drmcp.design_records_mcp.mvp_scope` | retain_existing_owner | none | mvp-scope.md explicitly excludes "Workflow-specific traversal / tree / graph query tools" |
| 32. Task dependency-cycle detection | T05 Evidence; same | `spec:drmcp.design_records_mcp.mvp_scope` | retain_existing_owner | none | mvp-scope.md explicitly excludes "Task dependency cycle detection / execution order projection" |
| 33. Execution-order projection | T05 Evidence; same | `spec:drmcp.design_records_mcp.mvp_scope` | retain_existing_owner | none | mvp-scope.md explicitly excludes "Task dependency cycle detection / execution order projection" |
| 34. Implementation-update projection from Design Records | `spec:product.design_records.artifact_model.change_and_investigation_flow` ## Deferred implementation tracking disposition | none | delete_with_rationale | none | Unadopted integration mechanism removed from PRODUCT in T04/T05. No accepted DRMCP app-local contract exists. No new contract is activated in T08. |

### Compatibility gaps deferred

| gap | deferred to |
|---|---|
| `schema/discovery.md` requires YAML front matter (`design_record.id` + `design_record.kind`) for spec inclusion; conflicts with new PRODUCT H1-adjacent bullet metadata contract | DRMCP-WORK-SPEC-001 |
| `schema/record-source.md` describes YAML front matter as the spec metadata source | DRMCP-WORK-SPEC-001 |
| `tools/resolve-reference.md` examples use `namespace_prefix = V01-` (MVP single-namespace assumption) | DRMCP implementation |
| Broad `spec:product.concepts.*` → `spec:product.design_records.*` ref sweep across remaining DRMCP files | T10 |
| `tools/overview.md` labels authoring transaction tools as MVP/P0; `mvp-scope.md` lists P0 tools as read-only and excludes other write tools — internal DRMCP spec contradiction | DRMCP follow-up |

### Inspection evidence

**PRODUCT task and evidence inputs**

- `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-08-apply-drmcp-app-local-handoff.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-01-confirm-migration-manifest-and-validation-baseline.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-03-split-namespace-profile-and-compatibility.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-04-split-artifact-model-and-repository-layout.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-05-reconcile-traceability-and-extract-future-material.md`
- `product/records/investigations/spec/PRODUCT-INV-SPEC-005-product-spec-semantic-layer-and-top-level-ownership-structure.md`

**Current PRODUCT handoff pointers**

- `product/records/spec/design-records/namespace-model/app-namespaces.md`
- `product/records/spec/design-records/namespace-model/existing-artifacts.md`
- `product/records/spec/design-records/namespace-model/subdomain-model.md`
- `product/records/spec/design-records/namespace-model/index.md`
- `product/records/spec/design-records/repository-layout/record-discovery-paths.md`
- `product/records/spec/design-records/artifact-model/traceability-boundary.md`
- `product/records/spec/design-records/traceability/resolve-and-validation.md`
- `product/records/spec/design-records/traceability/index.md`
- `product/records/spec/design-records/spec-format/follow-up-boundary.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`

**Named DRMCP outputs**

- `drmcp/records/spec/design-records-mcp/overview.md`
- `drmcp/records/spec/design-records-mcp/namespace-scanning.md`
- `drmcp/records/spec/design-records-mcp/resolver.md`
- `drmcp/records/spec/design-records-mcp/schema/id-normalization.md`
- `drmcp/records/spec/design-records-mcp/schema/overview.md`

**Directly implicated DRMCP owner specs**

- `drmcp/records/spec/design-records-mcp/schema/discovery.md`
- `drmcp/records/spec/design-records-mcp/schema/diagnostics.md`
- `drmcp/records/spec/design-records-mcp/schema/record-source.md`
- `drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md`
- `drmcp/records/spec/design-records-mcp/tools/overview.md`
- `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`
- `drmcp/records/spec/design-records-mcp/tools/validate-records.md`
- `drmcp/records/spec/design-records-mcp/tools/authoring-transaction-model.md`
- `drmcp/records/spec/design-records-mcp/tools/propose-record-create.md`
- `drmcp/records/spec/design-records-mcp/tools/propose-record-update.md`
- `drmcp/records/spec/design-records-mcp/tools/accept-proposed-write.md`
- `drmcp/records/spec/design-records-mcp/mvp-scope.md`
- `drmcp/records/spec/design-records-mcp/responsibility-boundary.md`

### Scope evidence

```
git status --short -- product/records/tasks/spec/PRODUCT-TASK-SPEC-012-08-apply-drmcp-app-local-handoff.md drmcp/records/spec/design-records-mcp
 M drmcp/records/spec/design-records-mcp/resolver.md
 M drmcp/records/spec/design-records-mcp/schema/discovery.md
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-08-apply-drmcp-app-local-handoff.md
```

The task file is untracked (`??`) and is excluded from `git diff` output. `resolver.md` and `schema/discovery.md` are the only tracked DRMCP files modified by T08.

```
git diff --name-status -- product/records/tasks/spec/PRODUCT-TASK-SPEC-012-08-apply-drmcp-app-local-handoff.md drmcp/records/spec/design-records-mcp
M	drmcp/records/spec/design-records-mcp/resolver.md
M	drmcp/records/spec/design-records-mcp/schema/discovery.md
```

```
git status --short -- bpdsl/records/spec v01
no output
```

```
git diff --name-status -- bpdsl/records/spec v01
no output
```

```
git diff --cached --name-status
no output
```

No BPDSL spec changes. No v01 changes. No staged changes.

### Post-edit validation

```
python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color
[strict]  All 47 file(s) OK.
exit 0

python -X utf8 product/src/tools/validate_spec.py drmcp/records/spec/design-records-mcp --strict --no-color
[strict]  All 30 file(s) OK.
exit 0
```

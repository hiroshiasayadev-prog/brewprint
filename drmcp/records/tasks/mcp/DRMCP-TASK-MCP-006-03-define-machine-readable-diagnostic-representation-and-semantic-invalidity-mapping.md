# DRMCP-TASK-MCP-006-03: Define machine-readable diagnostic representation and semantic-invalidity mapping

- **id**: DRMCP-TASK-MCP-006-03
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-006
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 2d
- **depends_on**:
  - DRMCP-TASK-MCP-006-02
- **outputs**:
  - spec:drmcp.design_records_mcp.schema.diagnostics
  - DRMCP-WORK-MCP-006

## Goal

Define the shared machine-readable representation for operation warnings and repository diagnostics.

Map PRODUCT-owned semantic invalidity into DRMCP-owned categories, severities, and associations without redefining semantic rules.

## Work

- Consume the accepted T01 authority baseline and T02 validation-execution contract.
- Confirm the exact `DRMCP-TASK-MCP-006-*` inventory before creating this Task.
- Record D01 through D11 as separate design decisions.
- Present and accept one design decision at a time.
- Record each accepted decision immediately in this Task Evidence.
- Keep normative specifications unchanged until D01 through D11 are accepted.
- Define shared envelope fields, subject associations, request occurrences, field values, targets, and conflicts.
- Define category ownership, severity vocabulary, deterministic ordering, and duplicate suppression.
- Define the boundary between operation warnings, repository diagnostics, and authoring diagnostics.
- Carry T02 review advisory A-01 into category and severity mapping.
- Reserve only the T04-owned source-location association required by the accepted envelope.
- Reflect the accepted contract into `schema/diagnostics.md` only after all eleven decisions are accepted.
- Change conditional files only when the accepted contract requires an authoritative pointer or operation-specific association.
- Run scoped validation and changed-file whitespace checks after normative reflection.
- Run independent review before changing this Task to `done`.

## Done condition

- D01 through D11 are accepted and recorded separately.
- The shared warning and diagnostic envelope is explicit.
- Canonical subject, request occurrence, field, value, target, and conflict associations are explicit.
- Category names preserve distinct operational causes without copying every PRODUCT rule into a category.
- Severity vocabulary and the T02 blocking severity are explicit.
- Operation warnings and repository diagnostics have a clear shared boundary.
- The authoring shared-envelope boundary is explicit, and authoring trigger semantics remain operation-local.
- Ordering and duplicate suppression are deterministic and independent of filesystem or map iteration order.
- Current duplicate identity and legacy issued-ID conflict remain distinct.
- Applicable PRODUCT authoring standards are cited for kind-specific semantic invalidity.
- The T04 source-location slot is reserved only when required; its concrete shape remains undecided here.
- `schema/diagnostics.md` contains no stale V01, YAML-front-matter, section-ref, retired request, or raw-path assumptions.
- Conditional files change only when a concrete pointer or operation field is required.
- Changed normative specifications pass scoped strict validation.
- Tracked and untracked changed files pass the applicable whitespace checks.
- Independent review reports no blocking, major, or minor findings before this Task is marked `done`.

## Verification

- Compare diagnostic ownership with `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`.
- Compare validation execution with the accepted T02 D01 through D07 decisions.
- Compare current-source and conflict associations with final W003 contracts.
- Compare operation warning triggers and occurrence behavior with final W004 contracts.
- Compare resolver and legacy lookup distinctions with final W005 contracts.
- Compare semantic invalidity with PRODUCT traceability and spec-format authorities.
- Compare kind-specific mappings with applicable PRODUCT authoring standards.
- Confirm that no PRODUCT semantic rule is redefined as DRMCP authority.
- Confirm that source-location shape and physical-path policy remain T04-owned.
- Run the strict spec validator against only changed normative specifications.
- Run `git diff --check` for tracked T03 files.
- Run `git diff --no-index --check` for any untracked T03 file.
- Run independent review before changing status to `done`.

## Evidence

### Exact Task inventory

The exact directory `drmcp/records/tasks/mcp/` was listed on 2026-06-28.

Existing `DRMCP-TASK-MCP-006-*` records were:

- `DRMCP-TASK-MCP-006-01-establish-validation-diagnostics-and-path-exposure-correction-baseline.md`;
- `DRMCP-TASK-MCP-006-02-define-current-repository-and-relation-validation-execution-contract.md`.

No `DRMCP-TASK-MCP-006-03` record existed.
The next Task is therefore `DRMCP-TASK-MCP-006-03`.

### Authority boundary

| concern | authority or accepted input | T03 treatment |
|---|---|---|
| Semantic ownership split | `DRMCP-ADR-MCP-001` | PRODUCT owns invalidity. DRMCP owns category, severity, associations, ordering, and response representation. |
| Required diagnostic outcomes | `DRMCP-REQ-MCP-001` | Provide machine-readable distinctions without reopening discovery, retrieval, or resolver behavior. |
| Historical contradictions | `DRMCP-INV-MCP-002` | Use audit findings as correction evidence, not current normative authority. |
| Current source and conflict state | `DRMCP-WORK-MCP-003` | Consume retained sources, addressable records, validation-only sources, and current conflict groups. |
| Operation warning triggers | `DRMCP-WORK-MCP-004` | Preserve list and exact-retrieval triggers, ordering, deduplication, and wrappers. |
| Resolver and legacy states | `DRMCP-WORK-MCP-005` | Preserve public statuses and distinct disabled, missing, duplicate, unreadable, and resolved lookup states. |
| Validation request and execution | `DRMCP-TASK-MCP-006-02` | Preserve request, execution-failure, subject-selection, relation-lookup, wrapper, and `ok` boundaries. |
| Canonical lookup invalidity | `spec:product.design_records.traceability.resolve_and_validation` | Consume canonical, unresolved, unsupported, duplicate, and relation invalidity. |
| Trace relation fields | `spec:product.design_records.traceability.metadata_schema` | Consume investigation and workflow relation semantics. |
| Current spec shape and identity | PRODUCT spec-format contracts | Consume document-shape, identity, and migration-severity inputs. |
| Kind-specific metadata and lifecycle | PRODUCT authoring standards | Consume required metadata, lifecycle, and gated-section rules by explicit authority pointer. |
| Diagnostic representation | T03 and `schema/diagnostics.md` | Own field envelope, categories, severities, associations, ordering, and suppression. |
| Source-location representation | T04 | T03 may reserve an association slot but does not define its concrete shape or path policy. |

### T02 fixed inputs

Accepted request forms are:

```json
{}
```

```json
{
  "app_namespace": "drmcp"
}
```

```json
{
  "ref": "DRMCP-WORK-MCP-006"
}
```

T03 preserves these T02 rules:

- `domain`, `kind`, `status`, `id_range`, ranges, legacy IDs, and physical paths are not validation selectors.
- Repository and app scopes include all W003-retained current validation inputs.
- Exact-ref scope selects one source, record, or complete duplicate-conflict group.
- Relation targets are lookup inputs and are not recursively added as validation subjects.
- Legacy archive records are never repository-validation subjects.
- Malformed validation scopes are request errors.
- Untrustworthy configuration or index state is an execution failure.
- Selected repository invalidity is returned as a normal validation diagnostic.
- Request errors and execution failures return no normal validation wrapper.
- Current relation lookup uses every configured current root without same-app restriction.
- Repository validation does not invoke public `resolve_reference`.
- Accepted legacy relation lookup uses W005 separate configured lookup state.
- Disabled, missing, duplicate, unreadable, resolved, and unsupported legacy states remain distinct inputs.
- Normal validation returns `ok`, `scope`, `summary`, and one unified `diagnostics` array.
- Validation does not add a separate `warnings` array.
- `ok` is false when at least one T03-defined blocking diagnostic exists.

### Contradiction inventory

| stale or incomplete claim | T03 disposition |
|---|---|
| Resolver public outcomes are named `unresolved_reference`, `ambiguous_reference`, and `unsupported_reference`. | Replace stale response claims while preserving W005 public statuses. |
| `get_records` uses `ids`, item-level `record_not_found`, `requested_id`, and one item wrapper per failure. | Replace with W004 top-level warnings associated with exact `refs` and request occurrences. |
| Duplicate request occurrences have fixed `info` severity. | Decide severity in D07 without changing W004 ordered deduplication. |
| Every standard diagnostic requires raw `path`. | Remove raw-path requiredness. Reserve any location association for T04. |
| One `duplicate_id` category covers current and legacy conflicts. | Separate W003 current conflict groups from W005 legacy lookup conflicts. |
| `spec_status_mismatch` compares current spec YAML values. | Remove the obsolete category. |
| Current spec `semantic_refs`, `sections`, and section targets remain validation inputs. | Remove obsolete categories and associations. |
| DRMCP locally owns required narrative sections and placeholder semantics. | Replace copied rules with applicable PRODUCT authoring-standard pointers. |
| Repository and authoring diagnostics share categories without an explicit boundary. | Define shared-envelope compatibility while preserving authoring operation-local triggers. |
| Error severity alone controls `validate_records.ok`. | Replace with the T03 blocking-severity rule accepted under D07. |
| V01 ADR and requirement citations remain category authority. | Replace with current PRODUCT and DRMCP authority pointers. |
| Source-location fields and exceptional path policy are mixed into the general envelope. | Reserve only the accepted T04 handoff slot. |

### Changed-file manifest

Unconditional files:

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-03-define-machine-readable-diagnostic-representation-and-semantic-invalidity-mapping.md`;
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`;
- `drmcp/records/spec/design-records-mcp/schema/diagnostics.md`.

Conditional files:

- `drmcp/records/spec/design-records-mcp/tools/validate-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`;
- `drmcp/records/spec/design-records-mcp/namespace-scanning.md`;
- `drmcp/records/spec/design-records-mcp/schema/discovery.md`;
- `drmcp/records/spec/design-records-mcp/schema/record-model.md`;
- `drmcp/records/spec/design-records-mcp/tools/propose-record-create.md`;
- `drmcp/records/spec/design-records-mcp/tools/propose-record-update.md`;
- `drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md`.

A conditional file changes only when the accepted T03 contract requires an authoritative pointer or operation-specific association field.

The authoring operation specs were added to the conditional manifest after D10 acceptance because their normative diagnostic examples use scalar `record_id`, scalar `field`, scalar `value`, and standalone diagnostic `path`, which directly contradict the accepted shared associations. `schema/authoring-transaction-schema.md` was added after static recheck found direct references to retired repository-validation categories. Target objects, diff summaries, files-written objects, and authoring operation semantics remain outside this diagnostic-only correction.

### Recheck-only manifest

The following files remain recheck-only unless a concrete contradiction is found:

- `drmcp/records/spec/design-records-mcp/tools/list-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/get-records.md`;
- `drmcp/records/spec/design-records-mcp/resolver.md`;
- `drmcp/records/spec/design-records-mcp/schema/record-source.md`;
- `drmcp/records/spec/design-records-mcp/schema/fields.md`;
- `drmcp/records/spec/design-records-mcp/schema/metadata-grammar.md`;
- `drmcp/records/spec/design-records-mcp/schema/id-normalization.md`.

### Decision register

| decision | question | status | accepted result |
|---|---|---|---|
| D01 | Which top-level fields form the shared envelope for operation warnings, repository diagnostics, and authoring diagnostics? | accepted | Required common fields are `category`, `severity`, and `message`. Optional common associations are `subject`, `field`, `value`, `target`, `occurrence`, and `conflict`. Each category defines which optional associations are required or prohibited. No untyped generic `details` object is introduced. Authoring-specific fields and trigger semantics remain operation-local pending D10. Source location remains deferred to D11 and T04. |
| D02 | How does the envelope identify a current record, spec, identity-less source, configuration subject, or request item without a physical path? | accepted | Use one structured `subject` object for all diagnostic subjects. `subject.type` distinguishes `record`, `source`, `configuration`, and `request_item`. Current canonical record and spec identities use canonical `ref`, including an identity whose D05 duplicate conflict prevents unique addressability; identity-less current sources use `app_namespace` and `record_kind` and are disambiguated later by the T04-owned location association; configuration subjects use `component`; request-item occurrence details remain D03-owned. Physical paths and duplicate identity fields such as `record_id`, `source_ref`, and `requested_id` are not placed in `subject`. |
| D03 | How are exact requested refs, first occurrences, duplicate occurrences, and request order associated with W004 warnings? | accepted | Use the structured `occurrence` association for W004 request-item warnings. For a duplicated exact requested ref, `occurrence` contains canonical `ref`, zero-based `first_index`, and ordered zero-based `duplicate_indexes`. Exact string equality controls duplicate detection; the first occurrence remains effective; later occurrences are ignored. Emit one warning per duplicated ref and order warnings by ascending `first_index`. Category and severity remain D06- and D07-owned. |
| D04 | How are metadata fields, invalid values, relation targets, expected kinds, and reciprocal counterparts associated? | accepted | Use separate structured `field`, `value`, and `target` associations. `field.name` identifies the affected metadata field or section and may include zero-based `item_index` for an array element. Missing fields omit `value`; invalid present values use `value.actual`, with `value.expected` only when an explicit machine-readable expectation exists. Relation diagnostics use `target.ref`; kind mismatches may add `actual_kind` and `expected_kinds`; reciprocal-link failures may add `reciprocal_field`. No physical path is placed in these associations, and D04 does not change diagnostic trigger semantics. |
| D05 | How are current duplicate-conflict groups and legacy issued-ID lookup conflicts represented without merging their states? | accepted | Distinguish `current_identity` conflicts from `legacy_lookup` conflicts through `conflict.type`. Current identity conflicts report `source_count`; legacy lookup conflicts report `candidate_count`. Neither conflict selects a winner. Every conflicting source or candidate must expose a machine-readable location so an authoring or repair agent can identify the files that require correction. The concrete location field shape and physical-path exposure policy remain D11- and T04-owned. Non-conflict legacy states such as disabled, missing, unreadable, and unsupported are not represented as conflicts. |
| D06 | Which category taxonomy preserves operational causes while avoiding one category per PRODUCT semantic rule? | accepted | Use stable cause-oriented categories rather than one category per PRODUCT rule. Repository validation categories are `source_unreadable`, `source_syntax_invalid`, `identity_invalid`, `missing_required_field`, `empty_required_field`, `invalid_field_value`, `missing_required_section`, `empty_required_section`, `noncanonical_section_heading`, `relation_target_unavailable`, `relation_target_kind_mismatch`, `relation_reciprocity_mismatch`, `current_identity_conflict`, and `legacy_lookup_conflict`. Operation-warning categories are `missing_compact_field`, `malformed_requested_ref`, `unsupported_requested_ref`, `unresolved_requested_ref`, `legacy_lookup_unavailable`, `legacy_lookup_conflict`, and `duplicate_requested_ref`. PRODUCT authority determines whether a field or section is required or must be non-empty; the category does not create that rule. An unsupported or noncanonical declared relation value that matches no accepted lookup grammar is `invalid_field_value`; no relation lookup runs and no lookup state is attached. `relation_target_unavailable` applies only after accepted current or legacy grammar and uses `target.lookup_state` for non-conflict lookup outcomes. Existing authoring categories remain unchanged pending D10. |
| D07 | Which severity vocabulary applies, and which severity makes T02 `ok` false? | accepted | Use `error`, `warning`, and `info`. `validate_records.ok` is `false` when at least one returned diagnostic has severity `error`; `warning` and `info` alone preserve `ok: true`. Malformed requests and untrustworthy configuration or index state remain request or execution failures outside the normal validation wrapper. Repository semantic invalidity is normally `error`; partial-success operation degradation is normally `warning`; non-blocking notices and repair guidance are `info`. Severity may vary by operation and applicable PRODUCT rule, but an identical condition in an identical context must always map to the same severity. |
| D08 | Which deterministic ordering key spans configuration, source, record, field, relation, and request-occurrence findings? | accepted | Diagnostics must be deterministically ordered without depending on filesystem traversal or execution completion order. Subject-type order is `configuration`, `source`, `record`, then `request_item`. Configuration subjects sort by `app_namespace` and `component`; source subjects by `app_namespace`, `record_kind`, and a stable T04-owned location sort key; record subjects by canonical `ref`; request items by zero-based `occurrence.first_index`. Within one subject, sort by `category`, `field.name`, `field.item_index`, `target.ref`, `target.lookup_state`, `target.reciprocal_field`, and `conflict.type`. `message` is never a sort key. W004 operation-specific request and result ordering remains authoritative. |
| D09 | Which exact diagnostic identity controls duplicate suppression without collapsing distinct causes or occurrences? | accepted | Duplicate suppression uses semantic identity composed from `category`, canonical `subject`, `field`, `target`, request occurrence identity (`ref` plus `first_index`), `conflict.type`, and source or finding location identity. Distinct subjects, fields, item indexes, targets, lookup states, reciprocal fields, conflict types, or locations remain distinct diagnostics. `message`, `severity`, actual or expected values, counts, and aggregated member-location lists do not define identity. Identical findings from multiple check paths collapse to one diagnostic; duplicate-request indexes and conflict member locations aggregate into that diagnostic. |
| D10 | Which envelope parts are shared with authoring diagnostics, and which authoring fields and trigger semantics remain operation-local? | accepted | Proposal-local validation diagnostics use the same category, severity, subject, field, value, target, conflict, ordering, and duplicate-suppression contracts as `validate_records`. Authoring-operation diagnostics also use the shared `category`, `severity`, `message`, and applicable structured associations. Existing authoring categories and trigger semantics remain unchanged. Authoring-only fields such as `proposal_id`, `body_cache_id`, `candidate_headings`, `stripped_heading`, and `stripped_level` remain category-specific and are not forced into generic associations. Proposal lifecycle, retention, write eligibility, body-cache behavior, and response wrappers are not redesigned. |
| D11 | Does T03 reserve a source-location association slot, while leaving its concrete shape and path policy to T04? | accepted | Add `location` as a shared diagnostic association. Source-backed repository and proposal-local validation diagnostics must expose a machine-readable location that identifies the repair target; an opaque source token alone is insufficient. Current identity conflicts and legacy lookup conflicts expose one location per conflicting source or candidate. Request-only and lifecycle-only diagnostics without a repository file omit `location`. T03 fixes the requirement to expose repairable locations; T04 owns the concrete object shape, repository-relative versus absolute path policy, root identification, separator normalization, and Windows-path handling. |

D01 was accepted on 2026-06-28.
D02 was accepted on 2026-06-28.
D03 was accepted on 2026-06-28.
D04 was accepted on 2026-06-28.
D05 was accepted on 2026-06-28.
D06 was accepted on 2026-06-28.
D07 was accepted on 2026-06-28.
D08 was accepted on 2026-06-28.
D09 was accepted on 2026-06-28.
D10 was accepted on 2026-06-28.
D11 was accepted on 2026-06-28.
All D01 through D11 decisions are accepted.
No normative specification may change until all eleven decisions are accepted.

### D01 accepted envelope families

D01 defines the shared field names only. It does not yet define the internal shape of each structured association.

Required common fields:

- `category`;
- `severity`;
- `message`.

Optional common association fields:

- `subject`;
- `field`;
- `value`;
- `target`;
- `occurrence`;
- `conflict`.

D01 rules:

- Each category defines which optional association fields are required, optional, or prohibited.
- The detailed structure of `subject`, `target`, `occurrence`, and `conflict` remains owned by later decisions.
- A generic untyped `details` object is not introduced.
- Authoring-specific fields and authoring trigger semantics remain operation-local pending D10.
- Source-location association remains deferred to D11 and T04.
- D01 does not change operation warning triggers, repository validation subjects, resolver statuses, or authoring transaction behavior.

### D02 accepted canonical subject association

All diagnostic subject association uses one structured `subject` object.

Supported `subject.type` values:

- `record` for a current canonical record or spec identity, including an identity whose duplicate conflict prevents unique addressability;
- `source` for a current source, including parse-failed or identity-less validation-only sources;
- `configuration` for repository configuration, root, index, or lookup-state subjects;
- `request_item` for an individual operation request item.

D02 rules:

- A current canonical record or spec identity uses canonical `ref` in `subject`, including a conflicted identity aggregated under D05.
- A `record` subject may include `record_kind` when the diagnostic contract requires the kind to be explicit.
- An identity-less `source` subject has no `ref` and uses `app_namespace` and `record_kind` as available subject context.
- Exact identity-less source discrimination is completed by the source-location association whose concrete representation and physical-path policy remain owned by D11 and T04.
- A `configuration` subject uses `component` to identify the affected configuration area and may include `app_namespace` when the subject is app-specific.
- A `request_item` subject identifies the operation context; requested values, first occurrence, duplicate occurrence, and request ordering remain owned by D03 and the `occurrence` association.
- Physical paths are not stored in `subject`.
- Duplicate identity fields such as `record_id`, `source_ref`, and `requested_id` are not introduced beside canonical `ref`.

### D03 accepted request occurrence association

W004 request-item warnings use the structured `occurrence` association.

For an exact requested ref that occurs more than once, `occurrence` contains:

- canonical `ref`;
- zero-based `first_index`;
- ordered zero-based `duplicate_indexes`.

D03 rules:

- Exact string equality controls whether requested refs are duplicates.
- The first occurrence remains effective and retains its request order.
- Every later occurrence of the same exact ref is ignored.
- One warning is emitted per duplicated ref, not one warning per ignored occurrence.
- `duplicate_indexes` preserves ascending request order.
- When more than one ref is duplicated, warnings are ordered by ascending `first_index`.
- Category naming remains D06-owned.
- Severity and blocking behavior remain D07-owned.
- D03 does not change W004 ordered deduplication, partial success, successful-record-only projection, or top-level warning placement.

### D04 accepted field, value, and relation-target associations

Field-, value-, and relation-specific context uses separate structured associations.

`field` association:

- `name` identifies the affected metadata field or required section;
- zero-based `item_index` may identify one element of an ordered array field.

`value` association:

- a missing field or section omits `value` rather than inventing a null actual value;
- an invalid present value uses `actual`;
- `expected` is included only when the accepted values or shape can be represented explicitly and machine-readably.

`target` association:

- `ref` identifies a relation target by canonical ref;
- a relation-kind mismatch may include `actual_kind` and ordered `expected_kinds`;
- a missing reciprocal relation may include `reciprocal_field` to identify the required counterpart field.

D04 rules:

- `field` says where the invalidity occurs.
- `value` says which present value was invalid and, when useful, what was expected.
- `target` says which related record or spec participates in the relation invalidity.
- Physical paths are not placed in `field`, `value`, or `target`.
- D04 does not add, remove, or redefine PRODUCT semantic rules, required metadata, lifecycle rules, narrative-section rules, relation lookup behavior, or diagnostic triggers.

### D05 accepted current and legacy conflict separation

Duplicate current identity and ambiguous legacy issued-ID lookup remain separate conflict families.

Current identity conflict:

- `conflict.type` is `current_identity`;
- `source_count` reports the number of current sources claiming the same canonical identity;
- all conflicting current sources remain part of one validation conflict group;
- no winner is selected.

Legacy issued-ID lookup conflict:

- `conflict.type` is `legacy_lookup`;
- `candidate_count` reports the number of configured legacy lookup candidates;
- the diagnostic remains associated with the current record relation that attempted the legacy lookup;
- legacy archive sources do not become repository-validation subjects;
- no candidate winner is selected.

Repair-location requirement:

- Every current conflicting source and every legacy lookup candidate must expose a machine-readable source location.
- This requirement exists so an authoring or repair agent can identify every file that must be corrected.
- D05 requires the locations to be available but does not define their concrete object shape, path normalization, absolute-versus-relative policy, or operation-specific exposure boundary.
- Those concrete decisions remain owned by D11 and T04.

D05 rules:

- Current identity and legacy lookup conflicts must not collapse into one generic duplicate-reference state.
- Legacy states such as disabled, missing, unreadable, resolved, and unsupported are lookup outcomes, not conflict objects.
- The exact category names remain D06-owned.
- Severity and blocking behavior remain D07-owned.

### D06 accepted category taxonomy

Categories are stable cause-oriented identifiers. They do not duplicate every PRODUCT-owned semantic rule.

Repository-validation categories:

- `source_unreadable`;
- `source_syntax_invalid`;
- `identity_invalid`;
- `missing_required_field`;
- `empty_required_field`;
- `invalid_field_value`;
- `missing_required_section`;
- `empty_required_section`;
- `noncanonical_section_heading`;
- `relation_target_unavailable`;
- `relation_target_kind_mismatch`;
- `relation_reciprocity_mismatch`;
- `current_identity_conflict`;
- `legacy_lookup_conflict`.

Operation-warning categories:

- `missing_compact_field`;
- `malformed_requested_ref`;
- `unsupported_requested_ref`;
- `unresolved_requested_ref`;
- `legacy_lookup_unavailable`;
- `legacy_lookup_conflict`;
- `duplicate_requested_ref`.

D06 rules:

- PRODUCT authority determines whether a metadata field or narrative section is required.
- PRODUCT authority also determines whether a present field or section must be non-empty.
- `empty_required_field` and `empty_required_section` fire only when the applicable PRODUCT rule requires non-empty content and the actual content is empty.
- A field or section whose applicable authority permits an empty value produces no empty-required diagnostic.
- PRODUCT rule details are represented through `field`, `value`, `target`, and expected-value associations rather than by creating one category per kind or rule.
- An unsupported or noncanonical declared relation value that matches no accepted current or legacy lookup grammar uses `invalid_field_value`; no lookup runs and no `target.lookup_state` is attached.
- `relation_target_unavailable` applies only after accepted current or legacy grammar and uses `target.lookup_state` to distinguish `current_unresolved`, `legacy_disabled`, `legacy_unresolved`, and `legacy_unreadable` causes.
- A duplicate legacy lookup remains the distinct `legacy_lookup_conflict` category and conflict family.
- Existing authoring-transaction category names and trigger semantics remain unchanged pending D10.
- Severity and blocking behavior remain D07-owned.

### D07 accepted severity and blocking contract

The shared severity vocabulary is:

- `error`;
- `warning`;
- `info`.

Validation-wrapper rule:

- `validate_records.ok` is `false` when at least one returned diagnostic has severity `error`.
- `validate_records.ok` is `true` when no returned diagnostic has severity `error`, including responses containing only `warning` or `info` diagnostics.
- A malformed validation request is a request error and does not return the normal validation wrapper.
- Configuration or active-index state that prevents trustworthy execution is an execution failure and does not return the normal validation wrapper.

Severity intent:

- `error` represents repository or record semantic invalidity under the applicable authority and blocks `validate_records.ok`.
- `warning` represents a normally returned operation result with partial failure, degraded projection, or an item that could not be fulfilled without failing the whole operation.
- `info` represents a non-blocking notice, ignored duplicate occurrence, or repair guidance that does not independently invalidate the repository.

Context rules:

- Repository validation uses `error` for violated applicable metadata, section, relation, identity, source, or conflict contracts.
- `list_records` and `get_records` partial-success warning entries normally use `warning`.
- Duplicate requested refs ignored after the first occurrence use `info`.
- `noncanonical_section_heading` uses `info` as repair guidance and does not suppress the corresponding blocking missing-section diagnostic when that diagnostic is required.
- The same category may have a different severity in a different operation context. For example, `legacy_lookup_conflict` is `error` when it makes a current record relation semantically invalid during repository validation and `warning` when one legacy exact-retrieval item cannot be fulfilled in a normal partial-success response.
- PRODUCT authority may define a relation such as a planned follow-up candidate as non-blocking. Such a condition may use `info` rather than `error`.
- An identical condition under the same operation and applicable authority must always produce the same severity.

### D08 accepted deterministic ordering

Diagnostic output order must not depend on filesystem traversal order, map iteration order, parallel completion order, or nondeterministic implementation details.

Primary subject-type order:

1. `configuration`;
2. `source`;
3. `record`;
4. `request_item`.

Subject ordering within each type:

- `configuration`: ascending `app_namespace`, then ascending `component`;
- `source`: ascending `app_namespace`, then ascending `record_kind`, then a stable source-location sort key supplied by the D11/T04 location contract;
- `record`: ascending canonical `ref`;
- `request_item`: ascending zero-based `occurrence.first_index`.

Diagnostic ordering within one subject:

1. ascending `category`;
2. ascending `field.name` when present;
3. ascending numeric `field.item_index` when present;
4. ascending `target.ref` when present;
5. ascending `target.lookup_state` when present;
6. ascending `target.reciprocal_field` when present;
7. ascending `conflict.type` when present.

D08 rules:

- Missing optional sort components sort before present values for the same preceding key.
- `message` is never used as a sort key, so wording changes do not reorder diagnostics.
- Duplicate suppression defined by D09 occurs before final deterministic sorting.
- `get_records` request-item warnings preserve W004 first-occurrence request order.
- Duplicate-request warnings use ascending `occurrence.first_index`.
- `list_records` missing compact-field warnings preserve returned-record order and then use the fixed compact-field order `title`, `status`, `date`.
- The D11/T04 location contract must supply a stable location sort key for identity-less sources and conflict members without making raw filesystem traversal order observable.

### D09 accepted diagnostic identity and duplicate suppression

Duplicate suppression uses semantic diagnostic identity rather than message text or detector origin.

Identity components:

- `category`;
- canonical `subject`;
- `field`, including `name` and `item_index` when present;
- `target`, including `ref`, `lookup_state`, and `reciprocal_field` when present;
- request occurrence identity, consisting of requested `ref` and zero-based `first_index`, when present;
- `conflict.type` when present;
- source or finding location identity when the diagnostic is location-specific.

The following do not define diagnostic identity:

- `message`;
- `severity`;
- `value.actual`;
- `value.expected`;
- source, candidate, or occurrence counts;
- aggregated duplicate indexes;
- aggregated conflict-member or candidate-location lists;
- the internal validation check or implementation path that discovered the condition.

D09 rules:

- The same semantic finding discovered through more than one check path is emitted once.
- Different subjects, fields, array item indexes, relation targets, lookup states, reciprocal fields, conflict types, or finding locations remain distinct diagnostics.
- If duplicate suppression encounters the same identity with different severities, the implementation has violated the D07 deterministic severity contract; it must not silently select one.
- Duplicate requested refs aggregate all later zero-based indexes into one ordered `duplicate_indexes` list for the first occurrence.
- One current identity conflict diagnostic aggregates every conflicting current source location for that canonical ref.
- One legacy lookup conflict diagnostic aggregates every candidate location for the same current relation lookup.
- Duplicate suppression occurs before the D08 deterministic final sort.

### D10 accepted authoring shared-envelope boundary

Proposal-local validation diagnostics use the same diagnostic contract as `validate_records`.

Shared proposal-local validation behavior:

- the same candidate or materialized record state produces the same category and severity;
- `subject`, `field`, `value`, `target`, `conflict`, and later location associations use the same shapes;
- D08 ordering and D09 duplicate suppression apply;
- authoring code does not translate repository-validation findings into a separate authoring-only taxonomy.

Authoring-operation diagnostics also use the shared envelope:

- `category`;
- `severity`;
- `message`;
- applicable structured `subject`, `field`, `value`, `target`, `occurrence`, `conflict`, and later location associations.

Authoring-only category fields remain operation-local when they express authoring concepts that do not fit the shared associations. These include, as applicable:

- `proposal_id`;
- `body_cache_id`;
- `candidate_headings`;
- `stripped_heading`;
- `stripped_level`.

D10 rules:

- Existing authoring category names remain unchanged.
- Existing authoring diagnostic trigger conditions remain unchanged.
- Existing proposal creation, accept, discard, get-proposal, body-cache, and retry semantics remain unchanged.
- Proposal lifecycle state, retention, stale checks, target-change checks, ID-collision checks, write eligibility, and post-write repair behavior are not redesigned.
- Authoring response wrappers remain operation-owned.
- Existing scalar `record_id`, scalar `field`, scalar `value`, and standalone diagnostic `path` representations are realigned to the accepted shared associations when the normative contracts are reflected.
- Authoring-specific information must not be hidden inside an untyped generic `details` object.

### D11 accepted source-location handoff

`location` is a shared diagnostic association for source-backed findings.

T03 requirements:

- Repository-validation diagnostics that identify a repairable source must expose a machine-readable `location`.
- Proposal-local validation diagnostics use the same location association when the candidate or persisted source has a repairable file target.
- An opaque internal source token without a usable path is insufficient for repair tooling.
- Identity-less, parse-failed, and unreadable current sources require location association when the source location is known.
- A current identity conflict exposes one location for every conflicting current source.
- A legacy lookup conflict exposes one location for every conflicting configured legacy candidate.
- Persisted-file authoring diagnostics may use the shared location association.
- Request-only diagnostics, unresolved request items with no source, unknown proposal IDs, expired cache IDs, and other lifecycle-only conditions omit `location`.

T04-owned concrete decisions:

- the exact `location` object fields;
- repository-relative versus absolute path representation;
- current-root and legacy-root identification;
- separator normalization and Windows-path handling;
- the stable source-location sort key required by D08;
- operation-specific exceptional path-exposure rules.

D11 fixes that repairable source locations are exposed. T04 may choose their concrete representation but must not replace them with a non-actionable opaque identifier.

### T02 review advisory A-01

T03 carries A-01 as a mandatory authority rule.

When a category or severity maps any of these conditions, the mapping must cite the applicable PRODUCT authoring standard:

- required metadata;
- lifecycle or status;
- done-gated or accepted-gated required sections;
- kind-specific narrative requirements.

T03 does not restate those semantic rules as DRMCP authority.

Applicable authorities include:

- `spec:product.design_records.authoring_standards.adr_authoring`;
- `spec:product.design_records.authoring_standards.spec_authoring`;
- `spec:product.design_records.authoring_standards.investigation_authoring`;
- `spec:product.design_records.authoring_standards.requirement_authoring`;
- `spec:product.design_records.authoring_standards.work_item_authoring`;
- `spec:product.design_records.authoring_standards.task_authoring`.

### Normative reflection

Normative reflection began only after D01 through D11 were accepted.

Changed normative specifications:

- `drmcp/records/spec/design-records-mcp/schema/diagnostics.md`;
- `drmcp/records/spec/design-records-mcp/tools/validate-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`;
- `drmcp/records/spec/design-records-mcp/tools/propose-record-create.md`;
- `drmcp/records/spec/design-records-mcp/tools/propose-record-update.md`;
- `drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md`.

Reflection results:

- `schema/diagnostics.md` now defines the accepted shared envelope, structured associations, cause-oriented categories, `error` / `warning` / `info` severity, deterministic ordering, semantic duplicate suppression, resolver cause diagnostics, authoring-envelope compatibility, and the T04-owned `location` handoff.
- Stale resolver response names, `get_records.ids`, item-level failure wrappers, scalar diagnostic identity shortcuts, generic current-and-legacy duplicate merging, obsolete YAML and section-ref categories, copied PRODUCT lifecycle and section rules, mandatory raw path, and V01 diagnostic authority claims were removed.
- `tools/validate-records.md` now states that `ok` is false exactly when at least one returned diagnostic has severity `error` and points to the shared diagnostics contract.
- `tools/resolve-reference.md` now includes top-level `diagnostics` in normal responses without changing W005 public statuses, current-first lookup, fallback eligibility, or successful target projection.
- `tools/propose-record-create.md` and `tools/propose-record-update.md` use structured shared associations in the affected authoring-operation diagnostic examples without changing their triggers or proposal behavior.
- `schema/authoring-transaction-schema.md` and `tools/propose-record-update.md` use the shared repository-validation categories for proposal-local validation.

Conditional-file disposition:

- `namespace-scanning.md`: unchanged; existing root, conflict, and diagnostic-owner pointers remain sufficient.
- `schema/discovery.md`: unchanged; existing invalid-source, validation-input, provenance, and W006 pointers remain sufficient.
- `schema/record-model.md`: unchanged; existing source-path retention and W006 repair-diagnostic delegation remain sufficient.

Recheck-only disposition:

- `tools/list-records.md`: unchanged; W004 warning triggers, compact projection, and normal path hiding remain authoritative.
- `tools/get-records.md`: unchanged; W004 exact classification, partial success, warning placement, and successful-record-only projection remain authoritative.
- `resolver.md`: unchanged; W005 lookup order and public status vocabulary remain authoritative.
- `schema/record-source.md`, `schema/fields.md`, `schema/metadata-grammar.md`, and `schema/id-normalization.md`: unchanged; no concrete T03 contradiction required an edit.

Static inspection found no positive remaining use of the retired repository-validation category set. A stale name appears only inside the explicit removed-behavior inventory in `schema/diagnostics.md`.
Scalar `record_id`, `requested_id`, and `path` values that remain in authoring target, diff, or files-written objects are not diagnostic envelope fields and were not changed.

### Validation state

- Required instruction and authoring standards: read.
- Exact authoring-standards directory inventory: complete.
- Exact W006 Task inventory: complete.
- Authority and upstream records: read.
- Required operation and schema contracts: read.
- PRODUCT semantic authorities: read.
- T03 Task creation: complete.
- W006 T03 linkage and start Evidence: synchronized.
- D01: accepted.
- D02: accepted.
- D03: accepted.
- D04: accepted.
- D05: accepted.
- D06: accepted.
- D07: accepted.
- D08: accepted.
- D09: accepted.
- D10: accepted.
- D11: accepted.
- D01 through D11: all accepted.
- Normative specification reflection: complete for the six-file T03 changed-spec set.
- Conditional file disposition: complete and recorded above.
- Static stale-category and scalar-diagnostic inspection: complete.
- Changed-spec H1, H1-adjacent metadata, and kind-required major-section inspection: complete.
- Pre-review external scoped strict validator executed on 2026-06-28 against the complete six-file T03 normative set.
  - Result: `[strict]  All 6 file(s) OK.`
- Pre-review external tracked-file whitespace check completed.
  - Result: `tracked_exit=0`.
  - No whitespace error was reported for the six changed specs and W006.
- Pre-review external untracked Task whitespace check completed with `git diff --no-index --check -- NUL <T03 Task>`.
  - Result: `untracked_exit=1`.
  - Exit code `1` is expected because the new Task differs from `NUL`; no whitespace error or exit code `2` or greater was reported.
- Targeted status confirmed the complete T03 changed-file set:
  - `M drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md`;
  - `M drmcp/records/spec/design-records-mcp/schema/diagnostics.md`;
  - `M drmcp/records/spec/design-records-mcp/tools/propose-record-create.md`;
  - `M drmcp/records/spec/design-records-mcp/tools/propose-record-update.md`;
  - `M drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`;
  - `M drmcp/records/spec/design-records-mcp/tools/validate-records.md`;
  - `M drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`;
  - `?? drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-03-define-machine-readable-diagnostic-representation-and-semantic-invalidity-mapping.md`.
- Initial independent review verdict: `NEEDS REVISION`.
  - `F-MAJ-01`: unsupported legacy relation syntax could map either to `invalid_field_value` or `relation_target_unavailable` with `legacy_unsupported`, breaking deterministic category and duplicate identity.
  - `F-MIN-01`: `tools/propose-record-update.md` referenced a removed generic Required narrative section policy instead of the applicable PRODUCT authoring standards.
- Corrections applied on 2026-06-28:
  - D06 now maps a relation value accepted by no current or legacy lookup grammar exclusively to `invalid_field_value`, performs no lookup, and attaches no lookup state.
  - `legacy_unsupported` was removed from the accepted repository `target.lookup_state` vocabulary.
  - `relation_target_unavailable` is limited to accepted-grammar lookup attempts that produce a non-conflict unavailable state.
  - The narrow case-only authoring fallback now cites the requirement, work-item, and task PRODUCT authoring standards directly without changing its trigger or heading set.
- Post-correction external scoped strict validation completed against the complete six-file T03 normative set.
  - Result: `[strict]  All 6 file(s) OK.`
- Post-correction external tracked-file whitespace check completed.
  - Result: `tracked_exit=0`.
  - No whitespace error was reported for the six changed specs and W006.
- Post-correction external untracked Task whitespace check completed with `git diff --no-index --check -- NUL <T03 Task>`.
  - Result: `untracked_exit=1`.
  - Exit code `1` is expected because the new Task differs from `NUL`; no whitespace error or exit code `2` or greater was reported.
- Limited independent re-review verdict: `PASS`.
  - `F-MAJ-01`: `CLOSED`.
  - `F-MIN-01`: `CLOSED`.
  - No blocking, major, or minor findings remain.
  - No advisories remain.
  - Post-correction validator and tracked/untracked whitespace evidence were assessed as valid.
  - No regression was found in T02 execution boundaries, W005 lookup-state separation, D07-D09 determinism, authoring operation triggers, or the T04 location handoff.
- T03 closure readiness: `READY`.
- T03 closed as `done` on 2026-06-28.
- Repository-wide clean status: not inferred.

The filesystem MCP did not execute these repository-local commands. The external results above were supplied by the user and recorded without alteration.

### Independent review prompt skeleton

```text
C:\Users\imved\projects\brewprint

Review `DRMCP-TASK-MCP-006-03` independently.
Do not modify files.
Use filesystem MCP because DRMCP is unavailable.
Do not clone the repository into a sandbox.
Do not report repository-local commands as executed unless they were actually run.

Read first:
- prompt_chappy.md
- task, work-item, spec, writing, and agent authoring standards
- applicable ADR, investigation, requirement, and kind-specific authoring standards

Accepted baseline:
- DRMCP-ADR-MCP-001
- DRMCP-REQ-MCP-001
- DRMCP-WORK-MCP-003 through DRMCP-WORK-MCP-006
- DRMCP-TASK-MCP-006-01 through DRMCP-TASK-MCP-006-03

Review targets:
- drmcp/records/spec/design-records-mcp/schema/diagnostics.md
- drmcp/records/spec/design-records-mcp/tools/validate-records.md
- drmcp/records/spec/design-records-mcp/tools/resolve-reference.md
- drmcp/records/spec/design-records-mcp/tools/propose-record-create.md
- drmcp/records/spec/design-records-mcp/tools/propose-record-update.md
- drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md
- drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md
- drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-03-define-machine-readable-diagnostic-representation-and-semantic-invalidity-mapping.md

Confirm:
- D01 through D11 are explicit and reflected exactly;
- PRODUCT semantic invalidity is cited, not redefined;
- W003 current conflict state and W005 legacy conflict state remain distinct;
- W004 warning triggers, ordering, deduplication, and wrappers are unchanged;
- W005 public resolver statuses and lookup order are unchanged;
- T02 request, execution, wrapper, and `ok` boundaries are unchanged;
- authoring operation triggers and proposal behavior are not redesigned;
- scalar IDs and paths that remain in authoring target, diff, or files-written objects are not misclassified as diagnostic-envelope regressions;
- pre-existing broader authoring-format migration concerns outside the accepted D10 diagnostic-envelope boundary are not reopened without a direct T03 contradiction;
- T04 still owns concrete source-location and physical-path exposure;
- deterministic ordering and duplicate suppression are complete;
- `schema/diagnostics.md` no longer positively defines stale V01 authority, YAML or section-ref categories, `ids`, item wrappers, or mandatory raw-path fields;
- scoped validator and complete-manifest whitespace evidence are valid.

Output:
1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition when applicable
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. D01-D11 assessment
8. PRODUCT authority assessment
9. W003-W005 and T02 non-regression assessment
10. Authoring-envelope boundary assessment
11. T04 handoff assessment
12. Changed-file and validation evidence assessment
13. T03 closure readiness
```

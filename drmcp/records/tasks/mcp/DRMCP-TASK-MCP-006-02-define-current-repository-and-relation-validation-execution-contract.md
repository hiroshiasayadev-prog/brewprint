# DRMCP-TASK-MCP-006-02: Define current repository and relation-validation execution contract

- **id**: DRMCP-TASK-MCP-006-02
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-006
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1.5d
- **depends_on**:
  - DRMCP-TASK-MCP-006-01
- **outputs**:
  - spec:drmcp.design_records_mcp.tools.validate_records
  - DRMCP-WORK-MCP-006

## Goal

Define the execution contract for current repository validation and relation validation.

Reflect the accepted contract into `validate_records` without redefining PRODUCT semantic invalidity, W003 input construction, W004 read behavior, or W005 resolver behavior.

## Work

- Consume the accepted T01 authority matrix, upstream inputs, contradiction inventory, manifests, and T02-T05 split.
- Confirm the exact `DRMCP-TASK-MCP-006-*` inventory and reserve `DRMCP-TASK-MCP-006-02`.
- Define the `validate_records` request shape without restoring `id_range` or list-query semantics.
- Define validation selection across configured current roots, app namespaces, and exact current canonical refs.
- Define the W003-retained validation input set for every accepted request scope.
- Define startup, configuration, request, source, record, and relation failure boundaries.
- Define current canonical relation validation across all configured current roots.
- Define current-to-legacy relation validation from W005 underlying lookup states without invoking the public resolver operation.
- Exclude legacy archive records as repository-validation subjects.
- Define the repository-validation response wrapper and execution meaning of `ok`.
- Leave diagnostic envelope, category names, severity vocabulary, ordering, deduplication, and source-location representation to T03 and T04.
- Update `schema/discovery.md` or `schema/record-model.md` only when an authoritative validation-input pointer is required.
- Run scoped strict validation only after normative spec changes are complete.
- Prepare an independent review prompt and keep this Task open until review corrections are complete.

## Done condition

- D01 through D07 are explicitly decided and recorded.
- The request contract supports only accepted validation selectors and defines empty-request behavior.
- Every accepted selector maps to an explicit W003-retained validation input set.
- Active-index entries are not treated as the complete repository-validation input set.
- Startup and configuration failures are separated from source, record, and relation invalidity.
- Current relation targets are checked across all configured current roots without same-app restriction.
- Current-to-legacy relation checks distinguish accepted fallback states without validating legacy archive records as subjects.
- Malformed requests, unsupported selectors, execution failures, and validation diagnostics have separate boundaries.
- The response wrapper defines selection information, result collections, execution counts when retained, and `ok` semantics.
- `tools/validate-records.md` cites PRODUCT semantic authorities instead of copying obsolete V01, YAML front-matter, section-ref, or spec-status rules.
- T03 retains diagnostic envelope, category, severity, ordering, and deduplication ownership.
- T04 retains source-location and exceptional physical-path representation ownership.
- Conditional schema files change only when the accepted execution contract requires a pointer.
- Changed normative specs pass scoped strict validation.
- Independent review reports no blocking, major, or minor findings before this Task is marked `done`.

## Verification

- Compare execution ownership with `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`.
- Compare validation inputs and conflict states with final W003 contracts.
- Compare request and response exclusions with final W004 contracts.
- Compare current and accepted legacy lookup-state handling with final W005 contracts.
- Compare semantic invalidity with PRODUCT traceability, spec-format, and artifact-ID authorities.
- Confirm that repository validation does not call or redefine public `resolve_reference`.
- Confirm that legacy archive records are not validation subjects.
- Confirm that no detailed diagnostic or source-location schema is finalized in T02.
- Run the repository strict spec validator against only changed normative spec files.
- Run normal `git diff --check` against tracked T02 files and `git diff --no-index --check` against any untracked T02 file.
- Run an independent review before changing status to `done`.

## Evidence

### Task discovery

The exact directory `drmcp/records/tasks/mcp/` was listed on 2026-06-28.

The only existing `DRMCP-TASK-MCP-006-*` record was:

- `DRMCP-TASK-MCP-006-01-establish-validation-diagnostics-and-path-exposure-correction-baseline.md`.

No `DRMCP-TASK-MCP-006-02` record existed.
The next Task is therefore `DRMCP-TASK-MCP-006-02`.

### Accepted ownership boundary

| concern | T02 treatment |
|---|---|
| PRODUCT semantic invalidity | Consume by authority pointer. Do not redefine. |
| W003 discovery, parsing, identity, addressability, conflict, provenance, and index construction | Fixed input. Do not reopen. |
| W004 listing, exact retrieval, ordered deduplication, partial success, warning triggers, and normal path hiding | Fixed input and exclusion. Do not reopen. |
| W005 current-first resolver order, accepted legacy grammar, configured lookup state, public statuses, and successful target projection | Fixed input. Reuse underlying state only where relation validation requires it. |
| Repository validation request and execution | T02-owned. |
| Diagnostic envelope, category, severity, ordering, and deduplication | T03-owned. |
| Source location and exceptional path exposure | T04-owned. |
| Fixtures, runtime implementation, and tests | Outside W006 contract scope. |

### Fixed validation inputs

T02 consumes these W003 states:

- configured current roots;
- discovered current sources;
- parsed source state;
- addressable current records;
- identity-less validation-only sources;
- duplicate-conflict groups;
- source provenance;
- active-index state;
- separate current and legacy index scopes.

T02 does not redefine discovery conditions, parsing, canonical identity, addressability, duplicate detection, winner policy, or index construction.

### Decision register

| decision | question | status | accepted result |
|---|---|---|---|
| D01 | Request shape and empty-request meaning. | accepted | Empty request validates all configured current roots. `app_namespace` selects one configured current app. `ref` selects one exact current canonical ref. `app_namespace` and `ref` are mutually exclusive. No `domain`, `kind`, `id_range`, range, legacy-ID, or path selector is supported. Relation lookup may still use all configured current roots. |
| D02 | Validation input set for each accepted scope. | accepted | Validation subjects are limited to the selected repository, app, or exact current ref scope. Relation target lookup may use all configured current roots and configured legacy lookup state. Relation targets are not recursively added as validation subjects. |
| D03 | Startup and configuration failure boundary. | accepted | Any invalid mandatory current root or configured legacy root prevents construction of a trustworthy validation input set and causes startup or repository-validation execution failure. No normal `{ ok, diagnostics }` validation result is produced, and no partial index or app-local fallback validation is allowed. Empty valid roots and disabled legacy fallback remain valid configurations. |
| D04 | Current relation validation across configured current roots. | accepted | Every selected subject checks current canonical relation targets against the complete active index across all configured current roots. Exact lookup determines whether the target exists and is uniquely addressable. Missing targets remain unresolved; duplicate-conflict identities do not count as resolved. The public `resolve_reference` operation is not invoked. |
| D05 | Current-to-legacy relation validation outcomes. | accepted | Accepted legacy-family relation targets use W005 configured lookup state. One readable unique target satisfies the relation. Disabled fallback, missing target, duplicate conflict, and unreadable source remain distinct non-resolved outcomes. Unsupported syntax performs no legacy lookup. Invalid configured roots or unusable required lookup state remain D03 execution failures. |
| D06 | Request failure versus validation diagnostic boundary. | accepted | Malformed requests, unsupported selectors, unknown configured app selectors, and exact current refs with no selectable subject are request errors. Invalid mandatory configuration or unavailable required index state is an execution failure. Source, record, duplicate, current relation, and legacy relation invalidity discovered after successful subject selection is returned as validation diagnostics. |
| D07 | Response wrapper and `ok` execution semantics. | accepted | Normal validation returns one wrapper with `ok`, effective `scope`, `summary`, and unified `diagnostics`. `ok` is false when any T03-blocking diagnostic exists and true otherwise. Request errors and D03 execution failures return no normal wrapper. Separate validation `warnings` are not introduced. |

D01 was accepted on 2026-06-28.
D02 was accepted on 2026-06-28.
D03 was accepted on 2026-06-28.
D04 was accepted on 2026-06-28.
D05 was accepted on 2026-06-28.
D06 was accepted on 2026-06-28.
D07 was accepted on 2026-06-28.
All T02 design decisions are accepted.

### D01 accepted request contract

| request form | selected validation scope |
|---|---|
| `{}` | All configured current roots. |
| `{ "app_namespace": "<configured-app>" }` | All W003-retained validation inputs for exactly one configured current app. |
| `{ "ref": "<exact-current-canonical-ref>" }` | The exact current canonical record or current spec source identified by the ref, plus relation checks required by that subject. |

D01 rules:

- `app_namespace` and `ref` are mutually exclusive.
- `ref` accepts current canonical refs only.
- Accepted legacy IDs are not validation-scope selectors.
- Physical paths are not validation-scope selectors.
- `domain`, `kind`, `status`, `id_range`, one-sided ranges, and list-query filters are not supported.
- Empty request means repository-wide validation across every configured current root.
- Selecting one app does not restrict relation-target lookup to that app. Current relation lookup may use all configured current roots.
- Selecting one exact ref does not call the public `resolve_reference` operation.
- Identity-less validation-only sources cannot be selected by exact ref; they remain reachable through repository-wide or app-scoped validation.

Reason for excluding `domain`: domain-scoped validation would require another selector and source-classification boundary while providing limited value beyond app and exact-ref scopes. The request remains intentionally small.

### D02 accepted validation-input contract

| request scope | validation subjects |
|---|---|
| Empty request | Configuration and active-index state for all configured current roots; every discovered current source; parse-failed source; identity-less validation-only source; addressable current record; and duplicate-conflict group. |
| `app_namespace` | The same W003-retained input classes, limited to the selected configured current app. |
| `ref` | The exact current source or addressable record identified by the ref. A conflicted ref selects the complete duplicate-conflict group for that identity. |

D02 rules:

- Validation subjects are determined directly by the selected scope; they are not discovered by traversing relations.
- Current relation target lookup uses all configured current roots, including other app namespaces.
- Accepted legacy relation target lookup uses W005 configured legacy lookup state.
- A relation target is not recursively added as a validation subject.
- Validation checks the selected subject's relation existence, canonical form, reciprocity, and other PRODUCT-owned relation conditions.
- Validation does not recursively validate every metadata field, body section, or outgoing relation of the target record.
- Identity-less validation-only sources are included by empty or app-scoped validation and cannot be selected through `ref`.
- An exact-ref request remains a focused diagnostic operation rather than expanding to a transitive relation closure.
- Empty request validates every retained current subject across all configured current roots. Current relation targets are therefore also validated when they are part of that repository-wide subject set, but not because relation traversal added them.
- App-scoped request validates every retained current subject in that app. A current relation target in another app is lookup-only for the selected subject's relation check and is not added as a validation subject.

### D03 accepted startup and configuration failure contract

A trustworthy validation input set cannot be built when any mandatory configured root is invalid.

The following conditions cause startup or repository-validation execution failure:

- any configured current root is missing, not a directory, unreadable, outside `repository_root`, inconsistent with its `app_namespace`, duplicated, canonically conflicting, or overlapping a legacy root;
- any configured legacy root is missing, unreadable, outside `repository_root`, duplicated, overlapping another configured legacy root, or overlapping a configured current root;
- any other configuration condition prevents construction of the complete active index or configured legacy lookup state required by the request.

D03 rules:

- DRMCP does not omit an invalid configured root and continue with a partial active index or partial legacy lookup map.
- `validate_records` does not return its normal repository-validation wrapper when the required validation input set cannot be built.
- An `app_namespace` request does not bypass a failure in another mandatory configured current root.
- Configuration failure is not represented as ordinary source, record, or relation invalidity.
- Exact diagnostic or tool-error representation remains T03-owned.
- A valid current root with zero discovered records is valid and contributes an empty subject set.
- Missing `legacy_roots` and an explicit empty list are valid and mean legacy fallback is disabled.

### D04 accepted current-relation validation contract

Every selected validation subject checks its declared current canonical relation targets against the complete active index across all configured current roots.

D04 rules:

- Current relation lookup is repository-wide across configured current roots, regardless of whether the request scope is repository-wide, app-scoped, or exact-ref-scoped.
- Lookup uses exact current canonical identity.
- A uniquely addressable active-index target satisfies the target-existence check.
- No matching target is unresolved.
- A duplicate-conflict identity has no selected target and does not satisfy the target-existence check.
- A value outside accepted current canonical grammar is not repaired, prefix-completed, path-interpreted, or fuzzily normalized.
- The public `resolve_reference` operation is not called.
- Same-app restriction is prohibited; cross-app current relations are supported.
- PRODUCT-owned reciprocity and relation-integrity checks use the same resolved target state when those checks apply.
- D04 does not recursively add the relation target as a validation subject; D02 remains controlling.

### D05 accepted current-to-legacy relation validation contract

A current validation subject may declare a relation target in an accepted legacy sequential-ID family.

| legacy lookup state | relation result |
|---|---|
| Fallback configured; one readable source matches the exact issued ID. | Relation target exists. |
| Fallback disabled because `legacy_roots` is missing or empty. | Relation is not resolved. Preserve a distinct disabled-fallback outcome. |
| Fallback configured; no source matches the exact issued ID. | Relation is not resolved. Preserve an unresolved-target outcome. |
| Two or more sources produce the same issued ID. | Relation is not resolved. Preserve a duplicate-conflict outcome. |
| One indexed source exists but cannot be read. | Relation is not resolved. Preserve an unreadable-source outcome. |
| Input does not match accepted legacy-family grammar. | Unsupported relation value. Perform no legacy lookup. |
| Configured legacy roots or required lookup state cannot be built. | D03 startup or repository-validation execution failure. Do not return a normal validation result. |

D05 rules:

- Legacy archive records are lookup targets only and never become repository-validation subjects.
- Lookup uses the W005 exact issued-ID mapping and separate configured legacy lookup state.
- Validation does not invoke the public `resolve_reference` operation.
- Disabled fallback and unresolved target remain distinguishable because their causes differ.
- Duplicate and unreadable states do not count as an existing usable target.
- `V01-SPEC-*`, app-prefixless IDs, physical paths, fuzzy inputs, and other unsupported forms perform no legacy lookup.
- Exact diagnostic categories, severity, messages, and source-location fields remain T03 and T04 ownership.

### D06 accepted request, execution, and diagnostic boundary

| condition class | handling |
|---|---|
| Malformed request shape or field type. | Tool request error. No normal validation result. |
| Unsupported selector such as `domain`, `kind`, `status`, `id_range`, range, legacy ID, or physical path. | Tool request error. No normal validation result. |
| `app_namespace` and `ref` supplied together. | Tool request error. No normal validation result. |
| Selected `app_namespace` is not configured. | Tool request error. No normal validation result. |
| Selected `ref` does not match current canonical grammar. | Tool request error. No normal validation result. |
| Selected current canonical `ref` has no addressable source, record, or duplicate-conflict group. | Tool request error. No normal validation result. |
| Mandatory configuration or required index state cannot be built. | D03 startup or repository-validation execution failure. No normal validation result. |
| Valid configured app contains zero retained validation subjects. | Successful validation with an empty selected subject set. |
| Selected source, record, duplicate group, current relation, or legacy relation is invalid. | Validation diagnostic inside the normal result. |

D06 rules:

- Request errors mean the caller failed to identify a supported validation scope.
- Execution failures mean DRMCP could not build the trustworthy state required to validate that scope.
- Validation diagnostics mean DRMCP selected a valid scope and found invalid repository content or relation state.
- An unresolved exact `ref` selector is not converted into `ok: true` with an empty diagnostic set.
- Unsupported relation values inside a selected record are validation diagnostics, not request errors.
- Exact request-error codes and diagnostic categories remain T03-owned.

### D07 accepted response-wrapper contract

A successfully started validation execution returns one normal response wrapper:

```json
{
  "ok": true,
  "scope": {
    "app_namespace": "drmcp"
  },
  "summary": {
    "sources": 42,
    "addressable_records": 39,
    "validation_only_sources": 1,
    "conflict_groups": 1
  },
  "diagnostics": []
}
```

D07 rules:

- `scope` reports the effective supported selector: empty object for repository-wide validation, `app_namespace` for app scope, or `ref` for exact current-ref scope.
- `summary` reports selected-subject counts and does not include relation-target lookup counts.
- `summary.sources` counts selected discovered current sources.
- `summary.addressable_records` counts selected uniquely addressable current records.
- `summary.validation_only_sources` counts selected sources retained for validation without an addressable current identity.
- `summary.conflict_groups` counts selected duplicate-identity conflict groups.
- `diagnostics` is the single collection for source, record, duplicate, current-relation, and legacy-relation findings.
- Validation does not add a separate top-level `warnings` collection.
- `ok` is `false` when at least one diagnostic has the T03-defined blocking severity.
- `ok` is `true` when no blocking diagnostic exists, including an empty diagnostic collection.
- Request errors and D03 startup or execution failures do not return this normal wrapper.
- Diagnostic envelope, allowed categories, severity vocabulary, deterministic ordering, and deduplication remain T03-owned.

### Normative reflection

`drmcp/records/spec/design-records-mcp/tools/validate-records.md` was realigned after D01 through D07 were accepted.

The reflected contract now defines:

- empty, app-scoped, and exact-current-ref request forms;
- W003-retained validation subjects for every accepted scope;
- mandatory configuration and index prerequisites;
- repository-wide current relation lookup across configured current roots;
- accepted legacy relation lookup outcomes without validating archive records as subjects;
- request error, execution failure, and validation diagnostic boundaries;
- the normal `ok`, `scope`, `summary`, and unified `diagnostics` wrapper;
- PRODUCT semantic-authority pointers instead of copied obsolete V01 semantic rules.

No T03 diagnostic envelope, category set, severity vocabulary, ordering, deduplication, or T04 source-location representation was finalized in T02.

### Changed-file manifest

| file | disposition |
|---|---|
| `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-02-define-current-repository-and-relation-validation-execution-contract.md` | Created and updated with accepted D01-D07 decisions, normative-reflection evidence, conditional-file disposition, and verification state. |
| `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md` | Updated with T02 progress and normative-reflection evidence. |
| `drmcp/records/spec/design-records-mcp/tools/validate-records.md` | Realigned to the accepted current repository and relation-validation execution contract. |

### Conditional-file disposition

| file | condition for change | current disposition |
|---|---|---|
| `drmcp/records/spec/design-records-mcp/schema/discovery.md` | Change only if the accepted validation execution contract requires an explicit pointer from discovered invalid-source state to repository validation. | Rechecked after D02. No change required; W003 already retains parse-failed and identity-less sources for validation, and `validate-records.md` now consumes that state directly. |
| `drmcp/records/spec/design-records-mcp/schema/record-model.md` | Change only if the accepted validation execution contract requires an explicit pointer from retained source, record, conflict, or index state to repository validation. | Rechecked after D02. No change required; W003 already defines retained source, addressable record, conflict-group, provenance, and active-index state. |

### Exclusions

T02 does not define or change:

- PRODUCT semantic validity rules;
- current or legacy root discovery;
- Markdown parsing or metadata grammar;
- canonical identity or addressability;
- duplicate detection, conflict construction, or index winner policy;
- list or exact-retrieval request, ordering, partial-success, wrapper, or projection behavior;
- resolver grammar order, lookup order, public status vocabulary, or target projection;
- accepted legacy lexical grammar;
- diagnostic field envelope, category names, severity vocabulary, ordering, or deduplication;
- source-location structure or exceptional physical-path representation;
- authoring proposal-local validation;
- W007 disposition or PRODUCT owner-pointer updates;
- fixtures, runtime implementation, or automated tests.

### Validation state

- Normative spec edits: completed for `drmcp/records/spec/design-records-mcp/tools/validate-records.md`.
- Conditional schema recheck: completed; `schema/discovery.md` and `schema/record-model.md` require no T02 change.
- Scoped strict validator executed externally on 2026-06-28:
  - command target: `drmcp/records/spec/design-records-mcp/tools/validate-records.md`;
  - result: `[strict] All 1 file(s) OK.`
- Task and Work Item records are outside the strict spec-validator scope.
- External `git diff --check` passed for the two tracked T02 files:
  - `drmcp/records/spec/design-records-mcp/tools/validate-records.md`;
  - `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`.
- LF-to-CRLF warnings for those tracked files are non-blocking working-copy conversion notices.
- Targeted `git status --short` confirmed the complete T02 changed-file set:
  - `M drmcp/records/spec/design-records-mcp/tools/validate-records.md`;
  - `M drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`;
  - `?? drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-02-define-current-repository-and-relation-validation-execution-contract.md`.
- Independent review verdict: `NEEDS REVISION` with one minor finding, `F-MIN-01`.
  - The normal tracked-file `git diff --check` did not inspect the untracked T02 Task.
  - Evidence therefore must not claim that the complete T02 manifest has passed whitespace validation.
- `F-MIN-01` correction verification executed externally after the Evidence correction:
  - tracked-file check: `git diff --check -- <tracked T02 files>` returned `tracked_exit=0` with no whitespace error;
  - untracked Task check: `git diff --no-index --check -- NUL <T02 Task>` returned `untracked_exit=1` with no whitespace error;
  - exit code `1` is expected because the new Task differs from `NUL`; no exit code `2` or greater occurred;
  - targeted `git status --short` still reports the same two tracked modifications and one untracked T02 Task.
- Limited independent re-review verdict: `PASS`.
  - `F-MIN-01`: closed.
  - No blocking, major, or minor findings remain.
  - Recorded tracked and untracked whitespace evidence is sufficient.
  - No regression was found in the normative contract or changed-file manifest.
- Review advisory `A-01` is handed to T03: kind-specific required metadata, lifecycle, and done-gated section diagnostics should cite the applicable authoring standards when categories and severities are defined.
- T02 closure readiness: `READY`.
- No repository-wide clean working tree is inferred.
- Task status changed to `done` on 2026-06-28 after limited re-review returned `PASS`.

### Independent review prompt

```text
C:\Users\imved\projects\brewprint

`DRMCP-TASK-MCP-006-02`の独立reviewを行う。

ファイルは変更しないこと。

DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
sandboxへrepositoryを複製しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。

## 最初に読む

- `prompt_chappy.md`
- `product/records/spec/design-records/authoring-standards/task-authoring.md`
- `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
- `product/records/spec/design-records/authoring-standards/spec-authoring.md`
- `product/records/spec/design-records/authoring-standards/writing-standard.md`
- `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## Accepted baseline and authority

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-01-establish-validation-diagnostics-and-path-exposure-correction-baseline.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-06-track-validation-diagnostics-and-path-exposure-contract.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-design-records-mcp-contract-baseline-and-realignment.md`
- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/investigations/mcp/DRMCP-INV-MCP-002-design-records-mcp-contract-consistency-and-realignment-audit.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-003-current-discovery-and-active-index-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-005-resolver-and-configured-legacy-fallback-contract-realignment.md`

## Review targets

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-02-define-current-repository-and-relation-validation-execution-contract.md`
- `drmcp/records/spec/design-records-mcp/tools/validate-records.md`
- `drmcp/records/spec/design-records-mcp/schema/discovery.md` when changed
- `drmcp/records/spec/design-records-mcp/schema/record-model.md` when changed
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`

## PRODUCT semantic authorities

- `product/records/spec/design-records/traceability/resolve-and-validation.md`
- `product/records/spec/design-records/traceability/metadata-schema.md`
- `product/records/spec/design-records/spec-format/document-shape.md`
- `product/records/spec/design-records/spec-format/spec-id-as-ref.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`
- `product/records/spec/design-records/namespace-model/artifact-id-grammar.md`

## Review scope

Review only T02 validation execution behavior.
Do not redesign PRODUCT semantic invalidity.
Do not reopen accepted W003, W004, or W005 contracts.
Do not finalize T03 diagnostic-envelope details or T04 source-location representation.
Do not review fixtures, implementation, tests, authoring proposal-local validation, or W007 disposition.

Confirm:

- D01 through D07 are all explicit and internally consistent;
- request selectors do not restore `id_range` or list-query semantics;
- empty-request behavior is explicit;
- each scope validates the required W003-retained sources, records, conflict groups, provenance, and index state;
- active-index entries are not the sole validation inputs;
- startup, configuration, request, source, record, and relation boundaries are distinct;
- current relations resolve across all configured current roots without same-app restriction;
- current-to-legacy checks consume accepted W005 lookup states without invoking the public resolver operation;
- legacy archive records are excluded as validation subjects;
- malformed requests and unsupported selectors are tool errors rather than repository diagnostics;
- the response wrapper and `ok` semantics are execution-level and do not preempt T03 severity design;
- stale V01 current identity, YAML front matter, section-ref, retired range, invented spec-status, raw path, and proposal-local claims are removed;
- conditional schema edits are pointer-only and do not reopen W003;
- changed normative specs pass scoped strict validation;
- Task and Work Item authoring shapes are valid.

Repository-local commands, when available:

`python -X utf8 product/src/tools/validate_spec.py <changed normative spec files> --strict --no-color`

`git diff --check -- <complete T02 changed-file set>`

Do not infer a clean working tree.
Do not use `git add .`.

出力形式:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition when applicable
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. D01-D07 assessment
8. W003 input-boundary assessment
9. W004/W005 non-regression assessment
10. PRODUCT authority assessment
11. Conditional-file assessment
12. Validation evidence assessment
13. T02 closure readiness
```

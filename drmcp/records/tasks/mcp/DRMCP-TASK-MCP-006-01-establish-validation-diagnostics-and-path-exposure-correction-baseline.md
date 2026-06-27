# DRMCP-TASK-MCP-006-01: Establish validation, diagnostics, and path-exposure correction baseline

- **id**: DRMCP-TASK-MCP-006-01
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-006
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-005-05
- **outputs**:
  - DRMCP-TASK-MCP-001-06
  - DRMCP-WORK-MCP-006

## Goal

Establish the validation, diagnostics, warning, source-location, and exceptional path-exposure correction baseline for `DRMCP-WORK-MCP-006`.

Confirm the authority tree, upstream ownership inputs, contradiction inventory, candidate file manifest, and T02 through T05 split before normative spec edits begin.

## Work

- Confirm the actual `DRMCP-TASK-MCP-006-*` inventory and select the next Task without guessing an existing sequence.
- Read the accepted ADR, source Requirement, audit, hub records, and completed W003 through W005 outputs.
- Read the PRODUCT traceability, spec-format, artifact, namespace, and repository-layout authorities referenced by W006 and upstream contracts.
- Inspect the current diagnostic schema, validation operation, response-boundary summaries, and operation-specific warning triggers.
- Classify stale claims as W006 correction, upstream accepted input, conditional pointer synchronization, W007 disposition scope, authoring scope, or no correction required.
- Record the W006-owned normative candidate manifest and the recheck-only manifest.
- Separate PRODUCT-owned semantic invalidity from DRMCP-owned validation execution and diagnostic representation.
- Separate operation warning triggers from warning-entry taxonomy and representation.
- Separate normal path hiding owned by W004 and W005 from exceptional path exposure owned by W006.
- Record T02 through T05 design questions without deciding them prematurely.
- Prepare an independent review prompt for this baseline.

This Task does not edit normative DRMCP specs.
It establishes the reviewed correction boundary used by T02 through T05.

## Done condition

- The actual next Task is established from the exact task-directory inventory.
- The authority matrix and W003 through W005 ownership inputs are recorded.
- Every material stale claim in the current diagnostic and validation contracts has a classification.
- The validation input boundary includes W003-retained current sources, addressable records, conflict groups, and active-index state without reopening W003 behavior.
- Legacy archive records remain excluded as repository-validation subjects.
- Operation warning triggers remain owned by their operation contracts.
- Warning and diagnostic schema, severity, source-location, and exceptional path representation remain W006-owned.
- W004 normal list and exact-retrieval response projection remains unchanged.
- W005 resolver order, lookup scopes, public status vocabulary, and successful target projection remain unchanged.
- The candidate changed-file manifest and recheck-only manifest are explicit.
- Authoring transactions, validation Work Item disposition, fixtures, implementation, and tests are excluded.
- T02 through T05 open design questions are explicit.
- An independent review reports no blocking or major finding.

## Verification

- Compare the baseline against `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`.
- Compare semantic-invalidity claims against PRODUCT traceability, spec-format, and authoring authorities.
- Compare validation inputs and source provenance against final W003 contracts.
- Compare list and exact-retrieval warning triggers and normal response hiding against final W004 contracts.
- Compare resolver outcomes and configured legacy states against final W005 contracts.
- Confirm that no normative DRMCP spec changed during this Task.
- Run an independent review before changing this Task to `done`.

## Evidence

### Actual task discovery

The exact directory `drmcp/records/tasks/mcp/` was listed once.
No file beginning with `DRMCP-TASK-MCP-006-` existed.

The actual next Task is therefore `DRMCP-TASK-MCP-006-01`.
No existing Task was overwritten or inferred from a broader repository search.

### Authority baseline

| concern | authority or accepted input | T01 use |
|---|---|---|
| DRMCP authority split and path boundary | `DRMCP-ADR-MCP-001` | Accept PRODUCT semantic ownership, DRMCP validation and diagnostic ownership, normal path hiding, and narrow exceptional exposure. |
| Required read-validation outcome | `DRMCP-REQ-MCP-001` | Require current and cross-namespace validation, configured current-to-legacy relation handling, machine-readable diagnostics, archive exclusion, and path-exposure limits. |
| Historical contradiction evidence | `DRMCP-INV-MCP-002` | Use findings F06, F09-F13, F18-F19, and F21-F22 as gap evidence rather than current normative authority. |
| Current source, record, and index state | `DRMCP-WORK-MCP-003` | Consume current discovery, parsing, identity, addressability, invalid-source retention, conflict groups, source provenance, and separate index scopes. |
| Operation warning triggers and normal response shape | `DRMCP-WORK-MCP-004` | Preserve list and exact-retrieval request, ordering, partial success, wrapper, successful projection, and normal path hiding. |
| Resolver and configured legacy states | `DRMCP-WORK-MCP-005` | Preserve resolver order, lookup scopes, public statuses, exact legacy classification, and the distinctions W006 may represent diagnostically. |
| Canonical reference invalidity | `spec:product.design_records.traceability.resolve_and_validation` | Consume supported refs, relation invalidity, duplicate identity, and noncanonical-path conditions without copying them as DRMCP semantic authority. |
| Trace metadata and workflow relations | `spec:product.design_records.traceability.metadata_schema` | Consume investigation and workflow relation fields and invalid conditions. |
| Current spec document validity | `spec:product.design_records.spec_format.document_shape` | Consume visible metadata, H1, kind, contract class, and section-shape requirements. |
| Current spec identity | `spec:product.design_records.spec_format.spec_id_as_ref` | Consume path-derived identity, mismatch invalidity, and no automatic rewrite. |
| Spec migration severity and owner pointers | `spec:product.design_records.spec_format.validation_policy` | Consume the migration-phase severity input. Leave Work Item owner-pointer disposition to W007 and PRODUCT-WORK-SPEC-015. |
| Artifact lifecycle and required sections | PRODUCT authoring standards | Consume kind-specific required metadata and done-gated section semantics. Do not restate them as DRMCP authority. |
| Sequential identity grammar | `spec:product.design_records.namespace_model.artifact_id_grammar` | Consume complete app-aware IDs and reject bare grammar fragments as canonical refs. |
| Candidate placement | `spec:product.design_records.repository_layout.record_discovery_paths` | Consume path-pattern semantics without redefining discovery. |
| Diagnostic and validation representation | W006 and its normative outputs | Own category, severity, message, associations, ordering, validation wrapper, and exceptional source-location representation. |

### Upstream ownership inputs

| upstream owner | fixed input to W006 | W006 must not reopen |
|---|---|---|
| W003 | Configured current roots, current discovery, source parsing, path-derived spec identity, sequential identity, addressability, validation-only sources, conflict groups, source provenance, active-index construction, and separate current/legacy scopes. | Candidate discovery, parser behavior, canonical identity, duplicate detection, conflict winner policy, or index construction. |
| W004 | `list_records` and `get_records` request shapes, ordered deduplication, first-occurrence ordering, partial success, successful-record-only responses, top-level warning placement, request-wide `include_body`, and normal list/get path hiding. | Request fields, wrapper, record projection, ordering, partial-success semantics, or warning triggers. |
| W005 | Resolver current-first order, current resolved stop, unresolved-only legacy eligibility, configured legacy lookup, public statuses `resolved` / `unresolved` / `unsupported`, exact-retrieval classification split, and legacy state distinctions. | Resolver lookup order, status vocabulary, exact-retrieval lookup scope, legacy-root semantics, lexical grammar, or successful target projection. |

Repository validation cannot operate only on addressable active-index entries.
W003 retains identity-less invalid sources and duplicate-conflict sources as validation inputs.
W006 therefore validates W003-retained current source, record, conflict, and index state while excluding legacy archive records as validation subjects.

### Contradiction inventory

| file and current claim | classification | required disposition |
|---|---|---|
| `schema/diagnostics.md`: resolver responses use `unresolved_reference`, `ambiguous_reference`, and `unsupported_reference`. | W006 correction | Align diagnostic distinctions with W005 while preserving the fixed public resolver statuses. Do not add resolver statuses. |
| `schema/diagnostics.md`: `get_records` uses item-level `record_not_found`, `requested_id`, and an `ids` request array. | W006 correction | Replace stale item/error representation with W004 top-level warning entries associated with exact `refs` and occurrences. |
| `schema/diagnostics.md`: duplicate request handling is fixed to `info`. | W006 design question | Define severity under the shared warning policy without changing W004 ordered deduplication. |
| `schema/diagnostics.md`: standard diagnostics require raw `path`. | W006 correction | Replace the unqualified path field with a narrow source-location contract and an explicit absolute-path prohibition for normal diagnostics. |
| `schema/diagnostics.md`: `duplicate_id` describes normalized IDs without separating current conflict state and legacy lookup conflict. | W006 correction | Define operation and validation categories over W003 and W005 states without merging current and legacy validation scopes. |
| `schema/diagnostics.md`: `spec_status_mismatch` compares top-level and `design_record.status`. | W006 correction | Remove obsolete YAML/front-matter status behavior. Current spec metadata is H1-adjacent. |
| `schema/diagnostics.md`: semantic-ref declaration and section-target categories depend on front-matter `semantic_refs` and `sections`. | W006 correction | Remove obsolete current validation categories. Current spec refs are path-derived and section refs are inactive. |
| `schema/diagnostics.md`: required narrative-section table and placeholder policy are stated as DRMCP semantics. | duplicated authority | Replace semantic restatement with PRODUCT authority pointers and keep only diagnostic mapping. |
| `schema/diagnostics.md`: authoring transaction categories share the same document without a read-validation boundary. | conditional narrowing | Preserve shared-envelope reuse where valid. Do not redesign authoring transaction triggers or operation behavior in W006. |
| `tools/validate-records.md`: current canonical refs and examples are `V01-*`, including `V01-SPEC-*`. | W006 correction | Use current app-aware IDs and path-derived `spec:` refs. Treat accepted V01 sequential IDs only as configured relation targets, never current validation subjects. |
| `tools/validate-records.md`: request scope uses optional `kind` and `id_range` and imports old `list_records` range rules. | W006 correction | Define an independent validation request contract. Do not restore W004-retired range or broad-list semantics. |
| `tools/validate-records.md`: empty request validates all MVP-indexed records. | W006 correction | Define current repository validation over W003-retained current validation inputs, not legacy archive records and not only addressable index entries. |
| `tools/validate-records.md`: response exposes a raw physical `path`. | W006 correction | Apply the source-location and exceptional path policy. |
| `tools/validate-records.md`: spec status vocabulary is `confirmed`, `draft`, `wip`. | W006 correction | Remove the invented enum. PRODUCT currently defines no complete spec status vocabulary. |
| `tools/validate-records.md`: spec front-matter semantic refs, sections, and `spec_status_mismatch` are active checks. | W006 correction | Replace with current H1-adjacent and path-derived spec validation pointers. |
| `tools/validate-records.md`: repository validation and proposal-local authoring diagnostics are mixed. | boundary correction | Keep repository validation authoritative here. Leave proposal-local affected-set and authoring trigger behavior to authoring contracts. |
| W006 planning text: repository validation operates on the active index only. | planning correction | Include W003 validation-only and conflict sources while preserving active-index and source-state inputs. |
| W006 planning text: current-to-legacy validation could imply validating legacy archive records. | clarification | Validate current records that refer to configured accepted legacy targets. Do not validate legacy archive records as subjects. |
| Normal list/get/resolver records omit physical paths. | accepted upstream input | Recheck only. W006 must not redefine those projections. |

### Candidate changed-file manifest

| file | candidate action | planned Task |
|---|---|---|
| `schema/diagnostics.md` | Rewrite the shared warning and diagnostic envelope, category ownership, severity, association fields, ordering, duplicate suppression, and source-location representation. Remove stale V01, YAML, section-ref, range, and item-wrapper assumptions. | T03-T04 |
| `tools/validate-records.md` | Rewrite validation request, current validation-input scope, relation checks, request-failure boundary, response wrapper, and PRODUCT authority pointers. | T02-T04 |
| `tools/resolve-reference.md` | Conditional operation-pointer synchronization if T03 adds warning or diagnostic fields. Preserve W005 statuses, lookup order, and target projection. | T03-T05 conditional |
| `namespace-scanning.md` | Conditional diagnostic-pointer synchronization for invalid root configuration, duplicate legacy issued IDs, excluded alias candidates, and unreadable legacy sources. Preserve W003/W005 configuration behavior. | T03-T05 conditional |
| `schema/discovery.md` | Conditional source-location and validation-input pointer synchronization. Preserve W003 discovery and addressability behavior. | T02-T05 conditional |
| `schema/record-model.md` | Conditional validation-input and source-location pointer synchronization. Preserve W003 record states and W004 response delegation. | T02-T05 conditional |
| `responsibility-boundary.md` | Synchronize final W006 validation, warning, source-location, and exceptional path ownership without changing W004 normal hiding. | T05 conditional |
| `overview.md` | Pointer-only synchronization when final validation or exceptional path summaries require it. | T05 conditional |
| `tools/overview.md` | Pointer-only synchronization for the final shared warning and diagnostic authority. | T05 conditional |
| `mvp-scope.md` | Pointer-only synchronization for the final validation scope and exceptional path boundary. | T05 conditional |

Only `schema/diagnostics.md` and `tools/validate-records.md` are unconditional normative correction candidates at T01.
Conditional files change only when later decisions require an authoritative pointer or operation field.

### Recheck-only and unchanged candidate manifest

| file | recheck purpose | current disposition |
|---|---|---|
| `tools/list-records.md` | Confirm warning triggers, top-level placement, compact projection, and normal path hiding remain W004-owned. | Recheck only. No W006 change currently required. |
| `tools/get-records.md` | Confirm exact classification, partial success, successful-only wrapper, warning triggers, and normal path hiding remain unchanged. | Recheck only. No W006 change currently required. |
| `resolver.md` | Confirm W005 orchestration and public status vocabulary remain unchanged. | Recheck only. |
| `schema/record-source.md` | Confirm source material and provenance pointers remain sufficient. | Recheck only. |
| `schema/fields.md` | Confirm current parsed fields and the absent spec enum remain correct. | Recheck only. |
| `schema/metadata-grammar.md` | Confirm parser grammar remains W003-owned and semantic invalidity remains authority-linked. | Recheck only. |
| `schema/id-normalization.md` | Confirm identity mapping and no-repair behavior remain W003-owned. | Recheck only. |
| PRODUCT traceability and spec-format specs | Confirm authority inputs only. | No W006 edit. |
| Authoring transaction specs | Confirm shared-envelope compatibility only when necessary. | Outside W006 changed-file scope. |

### T02 through T05 responsibility split

| Task | responsibility | candidate normative scope |
|---|---|---|
| T02 | Define current repository validation inputs, request scope, execution axes, current cross-namespace relations, configured current-to-legacy target checks, legacy-subject exclusion, and request-failure versus diagnostic boundaries. | `tools/validate-records.md`; conditional pointers in `schema/discovery.md` and `schema/record-model.md` |
| T03 | Define shared warning and diagnostic fields, category mapping, severity, canonical ref and occurrence association, field/value association, deterministic ordering, duplicate suppression, and operation-versus-repository sharing. | `schema/diagnostics.md`; conditional operation pointers |
| T04 | Define structured source location and narrow path exposure for diagnostics, patch, debug, and emergency surfaces. Synchronize the affected validation and diagnostic contracts. | `schema/diagnostics.md`, `tools/validate-records.md`, conditional path-owner pointers |
| T05 | Recheck upstream operation contracts, synchronize overview and responsibility pointers, confirm no W007 overlap, validate the final manifest, run independent review, correct findings, and close W006. | Final changed set plus workflow records |

### Open design questions

#### T02 validation execution

- What request shape selects current repository validation without reintroducing `id_range` or broad listing semantics?
- Does one invocation validate all configured current roots, one app namespace, one exact current ref, or a constrained combination?
- How are startup and configured-root failures represented when no trustworthy validation input set can be built?
- Which W003 states are validation inputs: discovered source, parsed source, addressable record, conflict group, and index-level configuration state?
- How does internal relation validation consume W005 current and legacy lookup states without invoking or redefining the public resolver operation?
- How are request-shape failures separated from per-source and per-record diagnostics?
- Does `ok` remain false only when at least one error-severity diagnostic exists?

#### T03 warning and diagnostic representation

- Which fields are common to operation warnings and repository validation diagnostics?
- Is canonical association represented by `ref`, `source_ref`, `record_ref`, or another single field?
- How are request-item occurrences represented for exact retrieval duplicates and failures?
- How are `field`, `value`, target association, and conflict-source association represented?
- Which categories distinguish malformed input, unsupported input, unresolved accepted ref, disabled legacy fallback, duplicate request occurrence, invalid source, duplicate identity, unreadable source, and validation failure?
- Which severities apply to each operation warning and repository diagnostic?
- What deterministic ordering key applies across configuration, source, record, field, relation, and operation-item diagnostics?
- Are identical diagnostics suppressed, and what exact identity defines a duplicate diagnostic?
- Which envelope fields are shared with authoring diagnostics without adopting authoring trigger semantics?

#### T04 source location and exceptional path exposure

- Is source location repository-relative, configured-root-relative, or a structured pair containing app namespace and repository-relative path?
- Are line, column, range, and heading ordinal optional location fields?
- Are absolute paths always prohibited from normal warning and validation responses?
- How are root configuration errors represented when a candidate resolves outside `repository_root` without exposing the escaped absolute path?
- Do `get_records` and `resolve_reference` warnings ever include legacy source locations, or are those locations limited to explicit configuration, diagnostic, debug, or emergency surfaces?
- Which explicit patch, debug, or emergency operations may expose a path, and how is that surface distinguished from normal read responses?
- Does an exceptional path field use the same structured source-location schema or a separate privileged representation?

### Explicit exclusions

T01 and W006 do not own:

- current or legacy root discovery;
- Markdown parsing or metadata grammar;
- canonical identity or addressability;
- active-index or legacy lookup-map construction;
- list or exact-retrieval request, ordering, wrapper, and successful-record projection;
- resolver grammar order, lookup order, status vocabulary, or target projection;
- accepted legacy-family semantics or lexical grammar;
- authoring transaction request, proposal, patch, or write behavior;
- disposition of `DRMCP-WORK-SPEC-001` or `DRMCP-WORK-SPEC-002`;
- PRODUCT validation-policy owner-pointer changes;
- fixture authoring;
- runtime implementation;
- automated implementation tests;
- migration or repair of record files.

### Independent review finding disposition

The 2026-06-28 independent baseline review returned `NEEDS REVISION` with one minor finding.

`F-MIN-01`: CLOSED in the working files.

- `DRMCP-WORK-MCP-006` used `T07` in four places where the ownership boundary is `W007`.
- The Task-flow exclusion, T01 disposition boundary, and T05 overlap check now name `W007`.
- The Completion Condition now states that `W007` owns DRMCP validation Work Item disposition.
- The Completion Condition separately states that `PRODUCT-WORK-SPEC-015` owns PRODUCT validation-policy owner-pointer synchronization.
- No normative DRMCP spec changed while correcting the finding.
- The limited independent re-review returned `PASS` with `F-MIN-01` closed and no blocking, major, minor, or advisory findings.
- External `git diff --check` completed successfully for the three T01 workflow files.
- The LF-to-CRLF message for `DRMCP-WORK-MCP-006` is a non-blocking working-copy conversion warning, not a whitespace error.

### Current verification state

- Required instruction and authoring standards read: complete.
- Exact Task-directory inventory: complete.
- Authority and prerequisite records read: complete.
- Minimum normative DRMCP specs read: complete.
- PRODUCT authority paths resolved and read: complete.
- Normative DRMCP spec changes: none.
- Workflow changes: T01 created; W006 and hub T06 moved to `in_progress`; T01 linked from W006.
- Repository-local validator execution: not run; the current tool boundary cannot execute repository-local commands.
- Repository-local Git status: not run; no clean working tree is inferred.
- Repository-local `git diff --check`: PASS for the T01 Task, parent W006, and hub T06 files; one non-blocking LF-to-CRLF warning was reported for W006.
- T01 initial independent review: `NEEDS REVISION`; sole minor finding `F-MIN-01` corrected.
- T01 limited independent re-review: `PASS`; `F-MIN-01` closed with no remaining findings or advisories.
- T01 closure readiness: complete.

### Closure evidence

- Final status: `done`.
- Normative DRMCP spec changes: none.
- Changed workflow records:
  - `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-01-establish-validation-diagnostics-and-path-exposure-correction-baseline.md`
  - `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`
  - `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-06-track-validation-diagnostics-and-path-exposure-contract.md`
- Verification command:
  - `git diff --check -- drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-01-establish-validation-diagnostics-and-path-exposure-correction-baseline.md drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-06-track-validation-diagnostics-and-path-exposure-contract.md`
- Verification result: PASS; no whitespace errors.
- Residual limitation: repository-wide clean status was not asserted.

### Independent review prompt

```text
C:\Users\imved\projects\brewprint

`DRMCP-TASK-MCP-006-01`の独立baseline reviewを行う。

ファイルは変更しないこと。

DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
sandboxへrepositoryを複製しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。

## 最初に読む

- `prompt_chappy.md`
- `product/records/spec/design-records/authoring-standards/task-authoring.md`
- `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
- `product/records/spec/design-records/authoring-standards/writing-standard.md`
- `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`
- `product/records/spec/design-records/authoring-standards/spec-authoring.md`

## Authority and upstream baseline

- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-design-records-mcp-contract-baseline-and-realignment.md`
- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/investigations/mcp/DRMCP-INV-MCP-002-design-records-mcp-contract-consistency-and-realignment-audit.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-003-current-discovery-and-active-index-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-005-resolver-and-configured-legacy-fallback-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-06-track-validation-diagnostics-and-path-exposure-contract.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-01-establish-validation-diagnostics-and-path-exposure-correction-baseline.md`

## Normative contracts to inspect

- `drmcp/records/spec/design-records-mcp/schema/diagnostics.md`
- `drmcp/records/spec/design-records-mcp/tools/validate-records.md`
- `drmcp/records/spec/design-records-mcp/responsibility-boundary.md`
- `drmcp/records/spec/design-records-mcp/overview.md`
- `drmcp/records/spec/design-records-mcp/tools/overview.md`
- `drmcp/records/spec/design-records-mcp/mvp-scope.md`
- `drmcp/records/spec/design-records-mcp/tools/list-records.md`
- `drmcp/records/spec/design-records-mcp/tools/get-records.md`
- `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`
- `drmcp/records/spec/design-records-mcp/resolver.md`
- `drmcp/records/spec/design-records-mcp/namespace-scanning.md`
- `drmcp/records/spec/design-records-mcp/schema/discovery.md`
- `drmcp/records/spec/design-records-mcp/schema/record-model.md`
- `drmcp/records/spec/design-records-mcp/schema/record-source.md`
- `drmcp/records/spec/design-records-mcp/schema/fields.md`
- `drmcp/records/spec/design-records-mcp/schema/metadata-grammar.md`
- `drmcp/records/spec/design-records-mcp/schema/id-normalization.md`

## PRODUCT authorities

- `product/records/spec/design-records/traceability/index.md`
- `product/records/spec/design-records/traceability/resolve-and-validation.md`
- `product/records/spec/design-records/traceability/metadata-schema.md`
- `product/records/spec/design-records/spec-format/document-shape.md`
- `product/records/spec/design-records/spec-format/spec-id-as-ref.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`
- `product/records/spec/design-records/namespace-model/artifact-id-grammar.md`

## Review scope

Review only the T01 baseline.
Do not redesign the diagnostic schema or decide T02-T04 open questions.
Do not reopen accepted W003, W004, or W005 behavior.
Do not review implementation, fixtures, authoring transactions, or W007 disposition decisions.

Confirm:

- actual Task discovery and `DRMCP-TASK-MCP-006-01` selection are correct;
- authority matrix separates PRODUCT semantic invalidity from DRMCP representation;
- W003 validation-only and conflict sources are included without redefining discovery or index behavior;
- legacy archive records are excluded as validation subjects while accepted configured legacy targets remain relation-resolution inputs;
- W004 request, wrapper, ordering, partial-success, warning-trigger, and normal path-hiding contracts are preserved;
- W005 resolver order, lookup scope, status vocabulary, and target projection are preserved;
- contradiction inventory covers stale V01, YAML front matter, section refs, ranges, spec status, raw path, and authoring/repository-validation mixing;
- unconditional changed-file candidates are appropriately narrow;
- conditional and recheck-only manifests are correctly classified;
- T02-T05 split is coherent and leaves concrete design questions open;
- W007 owner-pointer and validation Work Item disposition scope is excluded;
- no normative DRMCP spec changed during T01;
- Task and Work Item authoring shapes are valid.

Repository-local commands, when available:

`git diff --check -- drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-01-establish-validation-diagnostics-and-path-exposure-correction-baseline.md drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-06-track-validation-diagnostics-and-path-exposure-contract.md`

Do not infer a clean working tree.
Do not use `git add .`.

出力形式:

1. Verdict: PASS / NEEDS REVISION
2. Blocking findings
3. Major findings
4. Minor findings
5. Advisories
6. Authority baseline assessment
7. Upstream ownership assessment
8. Contradiction inventory assessment
9. Candidate changed-file manifest assessment
10. Recheck-only manifest assessment
11. Open-question and Task-split assessment
12. T02 start readiness
13. T01 closure readiness
```

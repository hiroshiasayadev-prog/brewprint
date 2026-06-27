# DRMCP-TASK-MCP-005-01: Establish resolver and configured legacy-fallback correction baseline

- **id**: DRMCP-TASK-MCP-005-01
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-005
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**: []
- **outputs**:
  - DRMCP-TASK-MCP-001-05
  - DRMCP-WORK-MCP-005

## Goal

Establish the resolver and configured legacy-fallback correction baseline for `DRMCP-WORK-MCP-005`.

Confirm the authority baseline, claim classification, affected-file manifest, and W004/W006 ownership boundaries before normative spec edits begin.

## Work

- Confirm the completed W003, W004, and PRODUCT compatibility inputs.
- Extract the accepted resolver baseline from the ADR, Requirement, audit, and PRODUCT authorities.
- Inspect the resolver, namespace-scanning, discovery, identity, tool, validation, and response-boundary contracts.
- Classify candidate claims as W005 correction, W006 correction, already-correct upstream input, authoring transaction scope, historical evidence only, or no correction required.
- Record W005-owned normative candidates and rechecked-but-unchanged files.
- Separate resolver outcomes and trigger conditions from machine-readable diagnostics and path representation.
- Separate exact retrieval from reference resolution.
- Record the T02 through T05 responsibility split.
- Separate T01-level authority and ownership decisions from T03-specific configuration and archive-index design decisions.
- Record any T03 design decisions that remain intentionally open after the baseline is established.
- Prepare an independent review prompt for the baseline.

This Task does not edit normative DRMCP specs.
It establishes the manifest and ownership boundary used by T02 through T05.

## Done condition

- The authority baseline is recorded.
- Current-first resolution order is explicit.
- Fallback eligibility and `legacy_roots` gating are explicit.
- Accepted and rejected input families are explicit.
- Active-index and legacy-index separation is explicit.
- Issued legacy ID preservation is explicit.
- W004 exact-retrieval ownership remains unchanged.
- W006 diagnostic and path-representation ownership remains unchanged.
- Every candidate stale claim has a classification.
- Every candidate spec has a disposition.
- Authoring transaction, fixture, implementation, and test scope are excluded.
- T02 through T05 can proceed without reopening W003 or W004 decisions.
- No unresolved T01-level authority, scope, or ownership decision remains.
- T03-specific configuration and archive-index design questions are explicit rather than incorrectly treated as settled.
- The baseline is ready for independent review.

## Verification

- Compare the baseline against `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`.
- Compare current identity and index assumptions against final W003 contracts.
- Compare exact retrieval and resolver separation against final W004 contracts.
- Compare accepted legacy families against `spec:product.brewprint.compatibility.legacy_id_compatibility`.
- Compare current canonical inputs against PRODUCT namespace, spec identity, and traceability authorities.
- Compare diagnostic and path-representation claims against W006 ownership.
- Confirm that no normative DRMCP spec changed during this Task.
- Run an independent review before changing this Task to `done`.

## Evidence

### Authority and upstream baseline

| concern | authority or accepted input | T01 use |
|---|---|---|
| DRMCP authority and operation split | `DRMCP-ADR-MCP-001` | Accept current-first resolution, configured fallback, separate indexes, exact legacy grammar, and issued-ID preservation. |
| Required resolver outcome | `DRMCP-REQ-MCP-001` | Require current grammar first, configured legacy fallback, exact accepted families, rejection rules, and archive isolation. |
| Historical audit evidence | `DRMCP-INV-MCP-002` | Use resolver and namespace findings as gap evidence, not as current normative authority. |
| Current discovery and identity | `DRMCP-WORK-MCP-003` | Consume configured current roots, app-aware sequential IDs, path-derived `spec:` refs, and separate current/legacy index scopes. |
| Exact retrieval boundary | `DRMCP-WORK-MCP-004` | Keep `get_records` exact-only and prevent resolver invocation from exact retrieval. |
| Brewprint legacy compatibility | `spec:product.brewprint.compatibility.legacy_id_compatibility` | Accept only issued V01 decision, investigation, requirement, work-item, and task families. |
| V01 spec compatibility removal | `PRODUCT-WORK-SPEC-014` | Treat `V01-SPEC-*` as rejected and keep current spec identity path-derived. |
| Current sequential identity | `spec:product.design_records.namespace_model.artifact_id_grammar` | Accept complete app-aware current IDs and reject app-prefixless grammar fragments. |
| Current spec identity | `spec:product.design_records.spec_format.spec_id_as_ref` | Accept path-derived document-level `spec:` refs only. |
| Canonical traceability inputs | `spec:product.design_records.traceability.resolve_and_validation` | Consume current lookup sources and the rejection of physical paths and obsolete section/front-matter refs. |
| Legacy archive source and record model | `DRMCP-WORK-MCP-005` | Own the legacy-source parsing boundary, archive-record construction, duplicate handling, and lookup model required by configured fallback and exact legacy retrieval. T03 decides whether to add separate legacy sections to shared schema specs or introduce a dedicated legacy archive schema. |
| Diagnostics and path representation | `DRMCP-WORK-MCP-006` | Delegate diagnostic object shape, category, severity, source location, operation warning schema, and exceptional path exposure. |

Upstream readiness is complete:

- `DRMCP-WORK-MCP-003` is `done`.
- `DRMCP-WORK-MCP-004` is `done`.
- `DRMCP-TASK-MCP-001-03` is `done`.
- `DRMCP-TASK-MCP-001-04` is `done`.
- `PRODUCT-WORK-SPEC-014` is `done`.

### Accepted resolver baseline

#### Current-first order

1. Evaluate the input against current canonical grammar.
2. Query the active index when current grammar accepts the input.
3. Stop when a current target resolves.
4. Evaluate legacy fallback eligibility only after current resolution remains unresolved.
5. Test exact accepted legacy grammar without repairing or transforming the input.
6. Query the separate legacy archive index only after an exact accepted legacy grammar match.
7. Return an unresolved or unsupported outcome when neither accepted path resolves.

A current input is never rewritten into a legacy input.
A resolved current target prevents any legacy grammar or legacy-index lookup.

A string may satisfy both the current app-aware grammar and an accepted V01 legacy grammar.
Grammar overlap does not change the accepted order: query the active index first, stop on a resolved current target, and only after an unresolved current lookup evaluate the exact legacy grammar and configured legacy index.
T03 does not reopen this fallback condition.

#### Legacy fallback gate

- Legacy fallback is enabled only when configuration declares one or more `legacy_roots`.
- Missing `legacy_roots` disables fallback.
- DRMCP does not auto-discover `v01/` or another archive directory.
- Current roots and legacy roots remain separate configuration scopes.
- Active and legacy records remain in separate indexes.
- Legacy records never enter the active index.

#### Accepted legacy families

| accepted family | example |
|---|---|
| `V01-ADR-*` | `V01-ADR-096` |
| `V01-INV-*` | `V01-INV-MCP-001` |
| `V01-REQ-*` | `V01-REQ-PRODUCT-001` |
| `V01-WORK-*` | `V01-WORK-MCP-004` |
| `V01-TASK-*` | `V01-TASK-MCP-003-01` |

The resolver applies exact grammar matching.
The compatibility authority remains the owner of accepted family semantics.

#### Rejected inputs

- `V01-SPEC-*`.
- Bare `SPEC-*` IDs.
- App-prefixless current IDs.
- Physical file paths.
- Case repair.
- Whitespace repair.
- Prefix completion.
- Domain completion.
- Sequence repair.
- Fuzzy matching.
- Legacy YAML semantic-ref aliases.
- The reserved `yaml:` prefix as a direct resolver input.
- `fixture:` lookup inputs.
- Metadata aliases.
- Source-path aliases.

#### Successful legacy resolution

- The lookup uses the separate legacy archive index.
- The result preserves the issued legacy ID.
- The result is not renamed or rewritten into a current canonical ID.
- The legacy source is not promoted into a current authoring target.

#### W004 exact-retrieval boundary

- `get_records` remains the sole exact-retrieval operation.
- `get_records` classifies each exact input once.
- `get_records` never invokes `resolve_reference`.
- W005 does not redesign the `get_records` request or response.
- W005 does not change normal listing or exact-retrieval projection.

#### W006 diagnostic and path boundary

W005 owns resolver behavior, resolution order, fallback eligibility, outcome conditions, and successful target projection except for physical-path representation.
This includes the non-path target discriminator, canonical resolved identity, record kind, and any title or status fields retained by the corrected public operation contract.
W006 owns the machine-readable representation of warnings and diagnostics.
W006 also owns source-location and exceptional path-exposure representation.

W005 does not define:

- diagnostic object shape;
- diagnostic category names;
- severity;
- message fields;
- source-location fields;
- operation warning schema;
- `validate_records` execution semantics;
- exceptional path representation.

### Stale-claim classification

| file and claim | classification | required disposition |
|---|---|---|
| `tools/resolve-reference.md`: active `spec:` lookup uses front-matter `semantic_refs`. | W005 correction | Replace with path-derived document-level `spec:` identity and active-index lookup. |
| `tools/resolve-reference.md`: section-level `spec:` refs use front-matter `sections`. | W005 correction | Remove section-level refs from the active resolver input surface. |
| `tools/resolve-reference.md`: `namespace_prefix = V01-` represents the current grammar. | W005 correction | Replace with complete app-aware current IDs and separate legacy grammar evaluation. |
| `tools/resolve-reference.md`: `V01-SPEC-*` is supported. | W005 correction | Remove the family and classify it as unsupported. |
| `tools/resolve-reference.md`: current app-aware sequential IDs are absent. | W005 correction | Add current decision, investigation, requirement, work-item, and task input families by authority pointer. |
| `tools/resolve-reference.md`: current-first order is absent. | W005 correction | Define current grammar and active-index lookup before fallback. |
| `tools/resolve-reference.md`: fallback does not require configured `legacy_roots`. | W005 correction | Add the explicit configuration gate. |
| `tools/resolve-reference.md`: active and legacy lookup share one index model. | W005 correction | Define separate active and legacy lookup scopes. |
| `tools/resolve-reference.md`: validation must share the same resolver implementation logic. | W005 correction | Limit the operation contract to resolver behavior and outcomes. Validation execution remains W006-owned. |
| `tools/resolve-reference.md`: resolved targets always include physical paths. | split W005/W006 correction | W005 defines the successful non-path target projection. Remove normal path fields; W006 alone defines any exceptional path representation. |
| `tools/resolve-reference.md`: resolved document, section, and record target fields reflect obsolete section refs and V01-era path-first projection. | W005 correction | Remove section targets, retain only supported current/legacy target classes, and define the corrected non-path successful target projection. |
| `tools/resolve-reference.md`: operation diagnostics define categories, severity, and messages. | W006 correction | Keep only W005-owned outcome and trigger conditions. Defer representation to W006. |
| `tools/resolve-reference.md`: ambiguity and unsupported outcomes prescribe concrete diagnostics. | W006 correction | Retain the outcome condition only. Defer fields and severity to W006. |
| `tools/resolve-reference.md`: input is evaluated as-is and whitespace is not repaired. | already-correct upstream input | Retain exact input handling. |
| `tools/resolve-reference.md`: physical paths are unsupported canonical inputs. | no correction required | Retain the rejection. |
| `tools/resolve-reference.md`: `yaml:` is reserved and direct public behavior is undefined. | W005 correction | Classify `yaml:` as an unsupported current resolver input. Do not preserve an undefined public branch. |
| `tools/resolve-reference.md`: `fixture:` lookup has no active PRODUCT resolver authority. | W005 correction | Classify `fixture:` as an unsupported resolver input. |
| `tools/resolve-reference.md`: `internal-design:`, `coverage:`, `COV-*`, and unrecognized forms are not active resolver inputs. | no correction required | Retain the unsupported boundary by PRODUCT traceability pointer. |
| `resolver.md`: PRODUCT traceability pointer identifies canonical input and lookup-source authority. | already-correct upstream input | Retain the pointer. |
| `resolver.md`: current-first order and configured fallback are absent. | W005 correction | Add resolver orchestration and configuration gating. |
| `resolver.md`: `list_records` and retired `get_record` define the record-kind boundary. | W005 correction | Remove stale tool references and point to the current operation split. |
| `resolver.md`: lookup-source and exposed-record-kind sets may differ. | no correction required | Retain the general distinction without stale tool names. |
| `resolver.md`: V01 ADR citations are the direct contract source. | historical evidence only | Keep only as optional provenance or replace with current authority pointers. |
| `namespace-scanning.md`: current roots are explicit and active/legacy roots remain separate. | already-correct upstream input | Consume without reopening W003 decisions. |
| `namespace-scanning.md`: legacy roots are optional and never auto-discovered. | already-correct upstream input | Retain. |
| `namespace-scanning.md`: concrete `legacy_roots` entry contract is not defined. | W005 correction | Add the configured legacy-root inputs and archive-index construction boundary. |
| `namespace-scanning.md`: invalid legacy-root conditions lack resolver-facing gating detail. | W005 correction | Define when legacy fallback is unavailable. Exact diagnostic representation remains W006-owned. |
| `namespace-scanning.md`: resolver fallback order is explicitly excluded. | no correction required | Retain the ownership boundary. |
| `schema/record-model.md`: only discovered current sources and active-index records are defined. | W005 correction | Add or point to a separate legacy archive-record model, including addressability and duplicate-ID behavior, without weakening the current active-index model. |
| `schema/record-source.md`: only current source types and metadata sources are defined. | W005 correction | Define or delegate the accepted legacy source material and source-provenance boundary used to construct archive records. |
| `schema/fields.md`: only current normalized field vocabulary is defined while `get_records` projects normalized metadata for legacy records. | W005 correction | Define which normalized non-identity fields a legacy archive record may supply and how issued legacy identity remains separate from current canonical identity. |
| `schema/metadata-grammar.md`: visible metadata grammar is explicitly current-only. | W005 conditional correction | Recheck in T03. Add a separate legacy grammar boundary or an explicit pointer to a dedicated legacy archive schema; do not broaden current grammar implicitly. |
| `schema/overview.md`: authoritative schema Topics table does not yet contain a dedicated legacy archive child. | W005 conditional correction | When T03 creates a dedicated legacy archive schema, add the matching child Topics row and summary in the same Task. Otherwise leave unchanged. |
| `schema/discovery.md`: current identity excludes legacy aliases. | already-correct upstream input | Recheck unchanged. |
| `schema/discovery.md`: resolver fallback order is delegated. | no correction required | Recheck unchanged. |
| `schema/id-normalization.md`: current identity is app-aware or path-derived and is not repaired. | already-correct upstream input | Recheck unchanged. |
| `schema/id-normalization.md`: legacy IDs belong to separate compatibility and index contracts. | already-correct upstream input | Recheck unchanged. |
| `tools/get-records.md`: exact retrieval classifies once and does not invoke the resolver. | already-correct upstream input | Recheck unchanged. |
| `tools/overview.md`: W005 owns current-first resolution and configured fallback. | already-correct upstream input | Recheck and change only if T02-T03 introduce a new pointer need. |
| `mvp-scope.md`: `resolve_reference` belongs to W005 and `get_records` does not invoke it. | already-correct upstream input | Recheck unchanged. |
| `responsibility-boundary.md`: W003-W006 ownership is separated. | already-correct upstream input | Recheck unchanged. |
| `schema/diagnostics.md`: resolver response categories, retrieval diagnostics, V01 spec claims, and path fields are current authority. | W006 correction | Leave unchanged in W005 and route to W006. |
| `tools/validate-records.md`: validation uses `V01-SPEC-*`, range coupling, obsolete spec refs, and old current grammar. | W006 correction | Leave unchanged in W005 and route to W006. |
| `tools/propose-record-create.md`: V01 IDs and `V01-SPEC-*` output examples define authoring inputs. | authoring transaction scope | Record as stale authoring evidence only. Do not edit in W005. |
| `tools/propose-record-update.md`: V01 IDs and YAML spec update examples define authoring behavior. | authoring transaction scope | Do not edit in W005. |
| `schema/authoring-transaction-schema.md`: V01 authoring identity and YAML spec metadata replacement remain current. | authoring transaction scope | Do not edit in W005. |

### Affected-file candidate manifest

#### W005-owned normative candidates

| file | candidate action | planned Task |
|---|---|---|
| `resolver.md` | Rewrite the current resolver responsibility, current-first order, fallback eligibility, and operation boundaries. | T02-T04 |
| `tools/resolve-reference.md` | Rewrite accepted input classes, current lookup, fallback trigger, legacy lookup, rejection behavior, and W006 delegation. | T02-T04 |
| `namespace-scanning.md` | Add the explicit `legacy_roots` configuration and legacy archive-index input boundary without reopening current-root rules. | T03 |
| `schema/record-model.md` | Define or point to the separate legacy archive-record and duplicate-ID model while preserving the current active-index contract. | T03 |
| `schema/record-source.md` | Define or point to accepted legacy source material and provenance used for archive-record construction. | T03 |
| `schema/fields.md` | Define the normalized legacy metadata subset required by `get_records` and resolver successful-target projection. | T03 |
| `schema/metadata-grammar.md` | Recheck and, when needed, define a separate legacy parsing boundary or point to a dedicated legacy archive schema. | T03 conditional |
| Dedicated legacy archive schema, if T03 selects separation over shared-schema extension | Define legacy source parsing, archive-record construction, and index-entry contract without redefining current record semantics. | T03 conditional new spec |
| `schema/overview.md` | If T03 creates a dedicated legacy archive schema, add the matching authoritative Topics row and summary in the same Task. Otherwise retain unchanged. | T03 conditional parent synchronization |

#### Rechecked-but-unchanged candidates

| file | recheck purpose | current disposition |
|---|---|---|
| `schema/discovery.md` | Confirm current identity and no-legacy-alias behavior remain current-only. | No W005 correction currently required. |
| `schema/id-normalization.md` | Confirm current canonical identity remains separate from legacy compatibility. | No W005 correction currently required. |
| `tools/get-records.md` | Confirm exact retrieval remains resolver-free and classification-once. | No W005 correction currently required. |
| `tools/overview.md` | Confirm the catalog points `resolve_reference` to W005. | Recheck in T04. Change only if synchronization requires it. |
| `mvp-scope.md` | Confirm the read baseline separates exact retrieval and resolver behavior. | Recheck in T04. Change only if synchronization requires it. |
| `responsibility-boundary.md` | Confirm W003-W006 ownership remains non-overlapping. | Recheck in T04. Change only if synchronization requires it. |

#### Routed to W006

| file or claim area | reason |
|---|---|
| `schema/diagnostics.md` | Owns diagnostic categories, severity, shared fields, source locations, and response representation. |
| `tools/validate-records.md` | Owns validation execution and current-to-legacy relation validation. |
| Diagnostic and path sections removed from `tools/resolve-reference.md` | W005 removes stale authority. W006 defines replacement representation. |

#### Authoring transaction scope

- `tools/propose-record-create.md`.
- `tools/propose-record-update.md`.
- `schema/authoring-transaction-schema.md`.

These files are excluded from the W005 changed-file manifest.

### Ownership exclusions

W005 does not own:

- current-root discovery or current active-index construction;
- current spec parsing or identity derivation;
- normal listing or exact retrieval request and response contracts;
- diagnostic object shape, category, severity, message, or source-location fields;
- `validate_records` request, execution, or response semantics;
- fixture creation;
- resolver, configuration, or legacy-index implementation;
- automated implementation tests;
- authoring transaction contracts;
- migration or rewriting of legacy files;
- accepted legacy-family semantics;
- normal list or exact-retrieval path hiding;
- W006 exceptional path representation.

### T02 through T05 responsibility split

| Task | responsibility | candidate normative scope |
|---|---|---|
| T02 | Define current grammar evaluation, active-index lookup, resolved-current stop condition, the already-decided unresolved-current fallback gate, resolver/exact-retrieval separation, and current successful non-path target projection. | `resolver.md`, `tools/resolve-reference.md` |
| T03 | Resolve the remaining `legacy_roots`, legacy source parsing, archive-record/index model, duplicate-ID, root-failure, lookup, and legacy successful-target projection details; then define exact accepted legacy grammar and issued-ID preservation. If a dedicated legacy archive schema is created, synchronize `schema/overview.md` in the same Task. The current-first fallback condition is not reopened. | `namespace-scanning.md`, `resolver.md`, `tools/resolve-reference.md`, legacy schema candidates, conditional `schema/overview.md` parent synchronization |
| T04 | Remove rejected-input and stale-tool claims. Synchronize catalogs and ownership pointers. Recheck unchanged candidates. | W005-owned files plus conditional pointer-only changes in `tools/overview.md`, `mvp-scope.md`, or `responsibility-boundary.md` |
| T05 | Validate the final changed contract set, run independent review, apply required corrections, and close W005 after findings are resolved. | Task and Work Item synchronization plus final normative manifest |

T02 through T04 must not edit W006-owned diagnostics or validation contracts.
T05 must not close W005 before independent review findings are resolved.

### Remaining design decisions

No unresolved T01-level authority, scope, operation-split, or ownership decision remains.
The accepted ADR, Requirement, PRODUCT authorities, W003 output, and W004 output determine that baseline.

T03 must still confirm and, where authority does not already decide the matter, resolve the concrete configuration and archive-index contract. Candidate questions include:

- the normative fields and identity of each `legacy_roots` entry;
- duplicate issued IDs across multiple configured legacy roots;
- missing, unreadable, overlapping, or otherwise unusable legacy roots and whether they disable one root, all fallback, or startup;
- the minimum legacy source material required to construct an archive-index entry;
- the archive-index contents needed for exact issued-ID lookup while preserving isolation from the active index;
- whether the legacy source and archive-record contract is expressed as separate sections in `schema/record-model.md`, `schema/record-source.md`, `schema/fields.md`, and `schema/metadata-grammar.md`, or as a dedicated legacy archive schema;
- when a dedicated legacy archive schema is selected, its child ref, title, and summary plus the matching `schema/overview.md` Topics-row synchronization in the same T03 change set;
- the legacy successful-target projection, excluding W006-owned diagnostic and path representation and aligned with W004 normalized retrieval fields.

These questions do not block T02 because T02 is limited to current-grammar evaluation, active-index lookup, the already-decided unresolved-only fallback gate, exact-retrieval separation, and current successful non-path target projection.
They must be resolved before T03 claims its normative correction complete.

### Current verification result

- Required instruction and authoring standards read: complete.
- Authority and prerequisite records read: complete.
- Candidate normative specs read: complete.
- W006 routing specs read: complete.
- Authoring transaction examples classified without edits: complete.
- Normative DRMCP spec changes: none.
- Workflow changes: hub T05 and W005 moved to `in_progress`; T01 created and linked.
- Unresolved T01-level authority, scope, or ownership decisions: none.
- T03-specific configuration, legacy-source, archive-record/index, duplicate, root-failure, and legacy target-projection design questions: recorded and intentionally open.
- Current/legacy grammar overlap fallback order: accepted baseline; not open for T03 redesign.
- `yaml:` and `fixture:` direct resolver inputs: classified as unsupported.
- Successful non-path resolver target projection owner: W005; path representation remains W006-owned.
- Scoped spec validation: not required because no normative spec changed.
- Repository-local whitespace check: deferred to staged pre-commit validation; not required for this re-review.
- Initial independent review: `NEEDS REVISION` with F-MAJ-01, F-MAJ-02, and F-MIN-01; all closed by the first correction pass.
- First independent re-review: `NEEDS REVISION` with F-MAJ-03 and F-MIN-02.
- F-MAJ-03 correction: `schema/overview.md` added as a T03 conditional parent-synchronization candidate when a dedicated legacy archive schema is selected.
- F-MIN-02 correction: W005 date updated to `2026-06-27` for the substantive scope correction.
- Independent final re-review: `PASS`.
- Previous findings closed: F-MAJ-01, F-MAJ-02, F-MIN-01, F-MAJ-03, and F-MIN-02.
- Remaining blocking, major, minor, and advisory findings: none.
- T02 start readiness: `READY`.
- T03 design-decision readiness: `READY`.
- T01 closure: complete on 2026-06-27.

### Independent review record

The final independent re-review used the following prompt and returned `PASS`:

```text
C:\Users\imved\projects\brewprint

`DRMCP-TASK-MCP-005-01`の第2修正後独立final re-reviewを行う。

最初に以下を読む。

- `prompt_chappy.md`
- `product/records/spec/design-records/authoring-standards/task-authoring.md`
- `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
- `product/records/spec/design-records/authoring-standards/writing-standard.md`
- `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`
- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-design-records-mcp-contract-baseline-and-realignment.md`
- `drmcp/records/investigations/mcp/DRMCP-INV-MCP-002-design-records-mcp-contract-consistency-and-realignment-audit.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-003-current-discovery-and-active-index-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-005-resolver-and-configured-legacy-fallback-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`
- `drmcp/records/spec/design-records-mcp/schema/record-model.md`
- `drmcp/records/spec/design-records-mcp/schema/record-source.md`
- `drmcp/records/spec/design-records-mcp/schema/fields.md`
- `drmcp/records/spec/design-records-mcp/schema/metadata-grammar.md`
- `drmcp/records/spec/design-records-mcp/schema/overview.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-014-remove-v01-spec-compatibility-authority.md`
- `product/records/spec/brewprint/compatibility/legacy-id-compatibility.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-005-01-establish-resolver-and-configured-legacy-fallback-correction-baseline.md`

必要に応じてT01に記録されたcandidate normative specsを読む。

これまでのreview findingは以下。

初回review:

- F-MAJ-01: current/legacy grammar overlap時のfallback条件をT03未決事項として再オープンしていた。
- F-MAJ-02: legacy source・archive record/index modelを扱うschema candidateがmanifestから欠落していた。
- F-MIN-01: `yaml:`と`fixture:` resolver input dispositionが未分類だった。

第1回re-review:

- F-MAJ-03: 専用legacy archive schema分岐に必要な`schema/overview.md`がmanifestに含まれていなかった。
- F-MIN-02: W005のsubstantive scope correctionに対するdate更新がなかった。

レビュー対象はT01のauthority baseline、claim classification、affected-file manifest、rechecked-but-unchanged判断、ownership exclusions、T02-T05 split、および全findingのclosure。

確認観点:

- current-first resolution orderがADR・REQ・PRODUCT authorityと一致するか
- `legacy_roots` gateとactive/legacy index separationが明確か
- accepted legacy familiesとrejected inputsが完全か
- W004 exact retrieval boundaryを再オープンしていないか
- W006 diagnostic・validation・path representationを侵食していないか
- `tools/resolve-reference.md`、`resolver.md`、`namespace-scanning.md`のclaim classificationが完全か
- `schema/discovery.md`、`schema/id-normalization.md`、`get-records.md`をunchanged候補とする判断が妥当か
- authoring transaction scopeが正しく除外されているか
- T02開始前に解消すべきauthority・scope・ownership判断が残っていないか
- current grammarとlegacy grammarの両方に一致し得る入力でも、active unresolved後にexact legacy fallbackする既決順序が再オープンされていないか
- `yaml:`と`fixture:`がunsupported resolver inputとして明示分類されているか
- successful targetの非path projectionがW005、diagnosticとpath representationがW006として分離されているか
- T03で解消すべき`legacy_roots`・legacy source・archive-record/index・duplicate・root-failure・legacy projection詳細が過不足なく明示されているか
- `schema/record-model.md`、`schema/record-source.md`、`schema/fields.md`、`schema/metadata-grammar.md`または専用legacy schemaがmanifestに含まれているか
- 専用legacy schema分岐を選ぶ場合、`schema/overview.md`のTopics rowとsummary同期が同じT03 conditional scopeに含まれているか
- T03固有の未決事項が、誤って既決事項やT01 blockerとして扱われていないか

ファイルは変更しない。

出力形式:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition
   - F-MAJ-01: CLOSED / OPEN
   - F-MAJ-02: CLOSED / OPEN
   - F-MIN-01: CLOSED / OPEN
   - F-MAJ-03: CLOSED / OPEN
   - F-MIN-02: CLOSED / OPEN
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. Authority baseline assessment
8. Claim-classification assessment
9. Affected-file manifest assessment
10. Ownership-boundary assessment
11. T02 start readiness
12. T03 design-decision readiness
```

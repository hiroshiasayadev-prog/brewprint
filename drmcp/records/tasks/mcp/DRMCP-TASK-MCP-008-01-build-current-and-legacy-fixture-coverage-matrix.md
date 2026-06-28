# DRMCP-TASK-MCP-008-01: Build current and legacy fixture coverage matrix

- **id**: DRMCP-TASK-MCP-008-01
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-008
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-001-07
- **outputs**:
  - DRMCP-WORK-MCP-008
  - DRMCP-TASK-MCP-001-08
  - DRMCP-WORK-MCP-009
  - DRMCP-WORK-MCP-010
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002

## Goal

Establish the requirement-to-fixture coverage baseline for current-format reads and configured legacy archive fallback.

Map accepted W003 through W007 contracts to stable fixture cases before any shared fixture file or production implementation is created.

## Work

- Confirm the bounded pre-creation Task inventory.
- Map every required accepted, rejected, unresolved, disabled, duplicate, and isolation outcome to a stable T01-local case.
- Inventory the existing inline, temporary-directory, repository-bootstrap, and helper-built test data in the bounded implementation test set.
- Separate stale implementation assertions from accepted fixture inputs.
- Propose package-local current, legacy, invalid-configuration, duplicate, overlap, and manifest locations.
- Define a candidate manifest schema and fixture-local structural checker boundary.
- Assign fixture authoring to T02 through T04.
- Assign future runtime assertions to W009, W010, W-SPEC-001, W-SPEC-002, or another accepted owner.
- Record the exact three-file workflow change boundary.
- Record ready-to-run scoped verification and independent review instructions.

This Task does not create fixture files.
It does not change production implementation or existing tests.

## Done condition

- Every required current accepted case has a matrix row.
- Every approved legacy family and configured fallback case has a matrix row.
- Every required rejection, disabled-fallback, unresolved, duplicate, overlap, and leakage case has a matrix row.
- Retained per-file spec and Topics graph validators have explicit shared-fixture ownership coverage.
- Current and legacy fixture roots are proposed as physically and logically separate areas.
- Current-only configuration is explicit.
- Invalid, duplicate, and overlapping root arrangements are explicit.
- Fallback-disabled and unresolved-legacy cases are distinct.
- Existing tests are classified without treating stale behavior as authority.
- Fixture-local checks are separated from runtime assertions.
- T02 through T04 authoring ownership is explicit.
- The changed-file manifest contains only this Task, W008, and hub T08.
- Repository-local commands are not reported as executed unless externally run.
- An independent review reports no blocking or major finding before this Task changes to `done`.

## Verification

- Compare every matrix row with `DRMCP-REQ-MCP-001`, `DRMCP-ADR-MCP-001`, W003 through W007, and PRODUCT compatibility authorities.
- Confirm the approved legacy-family set contains exactly ADR, INV, REQ, WORK, and TASK.
- Confirm no accepted row treats `V01-SPEC-*`, a bare ID, a physical path, inferred prefix, fuzzy repair, or YAML front matter as a canonical current input.
- Confirm current and legacy roots never share one normal fixture root.
- Confirm manifest checks do not execute production parser, index, resolver, validator, or MCP operations.
- Confirm existing production tests and implementation files remain unchanged.
- Run the recorded scoped lifecycle check and targeted whitespace commands externally.
- Run the recorded independent review before closure.

## Evidence

### Pre-creation inventory

The exact directory `drmcp/records/tasks/mcp/` was listed once.
No file beginning with `DRMCP-TASK-MCP-008-` existed before this Task was created.

The bounded starting state was:

- `DRMCP-WORK-MCP-008.tasks` was empty;
- `DRMCP-WORK-MCP-008.status` was `not_started`;
- `DRMCP-TASK-MCP-001-08.status` was `not_started`;
- no equivalent requirement-to-fixture matrix Task existed.

The selected Task ID is therefore `DRMCP-TASK-MCP-008-01`.
No existing Task was overwritten.

### Authority and ownership baseline

| concern | accepted authority or owner | T01 use |
|---|---|---|
| Source Requirement | `DRMCP-REQ-MCP-001` | Supplies required current, legacy, rejection, isolation, and fixture outcomes. |
| Accepted direction | `DRMCP-ADR-MCP-001` | Requires current-format-first behavior, separate legacy fallback, approved families, and path hiding. |
| Current discovery and identity | `DRMCP-WORK-MCP-003` | Supplies configured current roots, app-aware IDs, H1-adjacent spec metadata, path-derived refs, and duplicate behavior. |
| Listing and exact retrieval | `DRMCP-WORK-MCP-004` | Supplies compact active listing, exact retrieval, archive exclusion, and normal path hiding. |
| Resolver and legacy fallback | `DRMCP-WORK-MCP-005` | Supplies current-first order, `legacy_roots`, accepted legacy grammar, fallback states, and no-repair behavior. |
| Validation and diagnostics | `DRMCP-WORK-MCP-006` | Supplies current validation scope, current-to-legacy relation handling, archive exclusion, source locations, and leakage boundaries. |
| Validation-work disposition | `DRMCP-WORK-MCP-007` | Retains W-SPEC-001 and W-SPEC-002 and assigns shared fixtures to W008. |
| Legacy family policy | `PRODUCT-WORK-SPEC-014`; `spec:product.brewprint.compatibility.legacy_id_compatibility` | Accepts only V01 ADR, INV, REQ, WORK, and TASK families. Rejects `V01-SPEC-*`. |
| Current spec format | `spec:product.design_records.spec_format` | Supplies current H1-adjacent metadata and path-derived identity authorities. |
| Current runtime assertions | `DRMCP-WORK-MCP-009` | Owns general current read implementation and tests consuming accepted current fixtures. |
| Legacy runtime assertions | `DRMCP-WORK-MCP-010` | Owns configured legacy fallback, rejection, disabled, isolation, and leakage tests. |
| Per-file spec detectors | `DRMCP-WORK-SPEC-001` | Owns parser-aware per-file detector and integration tests consuming W008 fixtures. |
| Topics graph validation | `DRMCP-WORK-SPEC-002` | Owns graph algorithm and integration tests consuming W008 fixtures. |

W008 owns fixture material and fixture-local structure only.
The owners above remain responsible for runtime behavior.

### Requirement-to-fixture coverage matrix

#### Current accepted cases

| case ID | contract source | fixture class | root arrangement | record or config input | expected classification | future runtime owner | fixture-local verification | planned authoring Task | notes |
|---|---|---|---|---|---|---|---|---|---|
| C01 | `DRMCP-REQ-MCP-001`; `DRMCP-WORK-MCP-003`; `spec:product.design_records.namespace_model.artifact_id_grammar` | current | current-only | app-aware ADR such as `PRODUCT-ADR-001` | accepted | W009 | File is under the declared app root and manifest ref equals the visible ID. | T02 | No prefix inference. |
| C02 | `DRMCP-REQ-MCP-001`; `DRMCP-WORK-MCP-003`; `spec:product.design_records.namespace_model.artifact_id_grammar` | current | current-only | app-aware INV such as `DRMCP-INV-MCP-001` | accepted | W009 | Kind path, H1 ID, and manifest ref agree. | T02 | Uses current sequential grammar. |
| C03 | `DRMCP-REQ-MCP-001`; `DRMCP-WORK-MCP-003`; `spec:product.design_records.namespace_model.artifact_id_grammar` | current | current-only | app-aware REQ such as `DRMCP-REQ-MCP-901` | accepted | W009 | Kind path, metadata ID, and manifest ref agree. | T02 | Relation targets are separate cases. |
| C04 | `DRMCP-REQ-MCP-001`; `DRMCP-WORK-MCP-003`; `spec:product.design_records.namespace_model.artifact_id_grammar` | current | current-only | app-aware WORK such as `DRMCP-WORK-MCP-901` | accepted | W009 | Work Item is placed below the declared current root and references a current REQ. | T02 | Lifecycle semantics remain PRODUCT-owned. |
| C05 | `DRMCP-REQ-MCP-001`; `DRMCP-WORK-MCP-003`; `spec:product.design_records.namespace_model.artifact_id_grammar` | current | current-only | app-aware TASK such as `DRMCP-TASK-MCP-901-01` | accepted | W009 | Task placement, visible ID, parent Work Item ref, and manifest ref agree. | T02 | No bare `TASK-*` alias. |
| C06 | `DRMCP-WORK-MCP-003`; `spec:product.design_records.spec_format.spec_id_as_ref` | current | current-only | path-derived document ref such as `spec:product.fixture_baseline.overview` | accepted | W009 | Manifest derives the expected ref from the fixture-relative spec path. | T02 | No `SPEC-*` current identity. |
| C07 | `DRMCP-WORK-MCP-003`; `spec:product.design_records.spec_format.document_shape`; `DRMCP-WORK-SPEC-001` | current | current-only | H1 plus adjacent `id`, `status`, `date`, and `parent` metadata | accepted | W009 / W-SPEC-001 | Required visible metadata is adjacent to the real H1 and contains no YAML metadata authority. | T02 | Runtime parsing is W009; detector assertions are W-SPEC-001. |
| C08 | `DRMCP-WORK-MCP-003` | config | current-only | two declared active app roots, for example PRODUCT and DRMCP | accepted | W009 | Current root declarations are unique, contained, and use distinct app namespaces. | T02 | Proves multi-root current discovery. |
| C09 | `DRMCP-WORK-MCP-006`; `DRMCP-WORK-MCP-003` | relation | current-only | current record in one app refers to current record in another app | accepted | W009 | Both refs exist in separate declared current app roots and relation text is exact. | T02 | No public resolver invocation is asserted locally. |
| C10 | `DRMCP-WORK-MCP-005`; `DRMCP-WORK-MCP-009` | config | current-only | valid current roots with `legacy_roots` omitted | accepted | W009 | Manifest explicitly marks legacy configuration as omitted. | T02 | Distinct from an empty configured list. |
| C11 | `DRMCP-WORK-MCP-004`; `spec:drmcp.design_records_mcp.tools.get_records` | current | current-only | exact app-aware current ref and exact path-derived `spec:` ref | accepted | W009 | Manifest contains exact refs with no whitespace, case, path, or prefix variants. | T02 | Runtime ordering and projection remain W009 assertions. |
| C12 | `DRMCP-WORK-MCP-004`; `spec:drmcp.design_records_mcp.tools.list_records` | current | current-only | active sequential records from one app, kind, and domain | accepted | W009 | Manifest marks only current sequential records as list candidates. | T02 | Specs are addressable but not normal-list candidates. |
| C13 | `DRMCP-WORK-MCP-005`; `spec:drmcp.design_records_mcp.resolver` | current | current-only | exact current canonical ref with a unique current target | accepted | W009 | Manifest marks one current target and no competing legacy lookup input. | T02 | Current success stops fallback at runtime. |
| C14 | `DRMCP-WORK-MCP-006`; `spec:drmcp.design_records_mcp.tools.validate_records` | current | current-only | current repository containing valid sources and cross-root relations | accepted | W009 | Manifest enumerates current validation subjects only. | T02 | Specialized spec detectors remain retained owners. |
| C15 | `DRMCP-WORK-MCP-006`; `spec:drmcp.design_records_mcp.schema.diagnostics` | current | current-only | intentionally invalid current source with a portable source location | accepted | W009 / W-SPEC-001 | Invalidity is declared and the expected repository-relative source path exists. | T04 | Fixture-local check does not assert diagnostic category or severity. |
| C16 | `DRMCP-WORK-MCP-004`; `DRMCP-WORK-MCP-006` | isolation | current-only | normal listing and retrieval projection expectation | accepted | W009 | Manifest identifies response surfaces that must remain path-free without invoking them. | T04 | Runtime path-hiding assertion belongs to W009. |
| C17 | `DRMCP-WORK-SPEC-002`; `spec:product.design_records.spec_format.topics_table` | current | current-only | valid Index or Overview Topics row and canonical child `ref` | accepted | W-SPEC-002 | Parent and child files exist and the row contains `title`, `kind`, `ref`, and `summary`. | T02 | Supplemental retained-validator coverage. |

#### Configured legacy accepted cases

| case ID | contract source | fixture class | root arrangement | record or config input | expected classification | future runtime owner | fixture-local verification | planned authoring Task | notes |
|---|---|---|---|---|---|---|---|---|---|
| L01 | `DRMCP-WORK-MCP-005`; `spec:drmcp.design_records_mcp.namespace_scanning` | config | current+legacy | one declared legacy archive root separate from all current roots | accepted | W010 | Legacy root exists, is contained, and is neither equal to nor nested within a current root. | T03 | No automatic `v01/` discovery. |
| L02 | `spec:product.brewprint.compatibility.legacy_id_compatibility`; `DRMCP-WORK-MCP-005` | legacy | current+legacy | `V01-ADR-*` file keyed by issued filename ID | accepted | W010 | Filename-derived ID matches the ADR family grammar. | T03 | No current-model rewrite. |
| L03 | `spec:product.brewprint.compatibility.legacy_id_compatibility`; `DRMCP-WORK-MCP-005` | legacy | current+legacy | `V01-INV-*` | accepted | W010 | Filename-derived ID matches the INV family grammar. | T03 | Issued ID is preserved. |
| L04 | `spec:product.brewprint.compatibility.legacy_id_compatibility`; `DRMCP-WORK-MCP-005` | legacy | current+legacy | `V01-REQ-*` | accepted | W010 | Filename-derived ID matches the REQ family grammar. | T03 | Archive record remains read-only. |
| L05 | `spec:product.brewprint.compatibility.legacy_id_compatibility`; `DRMCP-WORK-MCP-005` | legacy | current+legacy | `V01-WORK-*` | accepted | W010 | Filename-derived ID matches the WORK family grammar. | T03 | Archive record remains outside the active index. |
| L06 | `spec:product.brewprint.compatibility.legacy_id_compatibility`; `DRMCP-WORK-MCP-005` | legacy | current+legacy | `V01-TASK-*` | accepted | W010 | Filename-derived ID matches the TASK family grammar. | T03 | Archive record remains outside authoring targets. |
| L07 | `DRMCP-WORK-MCP-004`; `DRMCP-WORK-MCP-005`; `spec:drmcp.design_records_mcp.tools.get_records` | legacy | current+legacy | exact approved V01 ID with fallback enabled | accepted | W010 | Manifest points to exactly one readable legacy source. | T03 | Runtime uses legacy exact classification, not resolver fallback. |
| L08 | `DRMCP-WORK-MCP-006`; `DRMCP-WORK-MCP-005` | relation | current+legacy | current record relation to an approved configured V01 target | accepted | W010 | Current source and exact legacy target both exist in separate roots. | T03 | Legacy target is not a validation subject. |
| L09 | `DRMCP-WORK-MCP-005`; `spec:drmcp.design_records_mcp.resolver` | relation | current+legacy | current-stage unresolved exact input followed by approved legacy lookup success | accepted | W010 | Manifest records one current miss and one unique accepted legacy source for the same exact input. | T03 | Current-first order is a runtime assertion. |
| L10 | `DRMCP-WORK-MCP-004`; `DRMCP-WORK-MCP-010` | isolation | current+legacy | approved legacy records present beside current roots | accepted | W010 | Manifest marks every legacy source `normal_listing_candidate: false`. | T04 | Leakage assertion is runtime-owned. |
| L11 | `DRMCP-WORK-MCP-006`; `DRMCP-WORK-MCP-010` | isolation | current+legacy | approved legacy records present during current validation | accepted | W010 | Manifest marks every legacy source `current_validation_subject: false`. | T04 | Current relations may still target legacy. |
| L12 | `DRMCP-WORK-MCP-010`; `DRMCP-ADR-MCP-001` | isolation | current+legacy | approved legacy records present during authoring discovery | accepted | W010 | Manifest marks every legacy source `authoring_target: false`. | T04 | W008 does not change authoring behavior. |
| L13 | `DRMCP-WORK-MCP-003`; `DRMCP-WORK-MCP-005`; `DRMCP-WORK-MCP-010` | isolation | current+legacy | approved legacy records present beside the active current index | accepted | W010 | Manifest marks every legacy source `active_index_candidate: false`. | T04 | Separate archive lookup input remains allowed. |

#### Rejected, unresolved, disabled, duplicate, and leakage cases

| case ID | contract source | fixture class | root arrangement | record or config input | expected classification | future runtime owner | fixture-local verification | planned authoring Task | notes |
|---|---|---|---|---|---|---|---|---|---|
| R01 | `PRODUCT-WORK-SPEC-014`; `spec:product.brewprint.compatibility.legacy_id_compatibility`; `DRMCP-WORK-MCP-005` | rejection | current+legacy | `V01-SPEC-*` | rejected | W010 | Manifest forbids accepted-family classification and records intentional rejection. | T04 | Must never become a positive legacy fixture. |
| R02 | `DRMCP-WORK-MCP-005`; `spec:product.design_records.namespace_model.artifact_id_grammar` | rejection | current-only | app-prefixless bare sequential ID | rejected | W009 / W010 | Input is recorded only in a rejection case and has no accepted target file. | T04 | No inferred app prefix. |
| R03 | `DRMCP-WORK-MCP-004`; `DRMCP-WORK-MCP-005` | rejection | current-only | repository-relative or absolute physical path used as canonical input | rejected | W009 / W010 | Manifest input type is `path_input` and no canonical-ref alias is declared. | T04 | Diagnostic locations are not canonical inputs. |
| R04 | `DRMCP-WORK-MCP-004`; `DRMCP-WORK-MCP-005` | rejection | current-only | misspelled or partial prefix requiring fuzzy repair | rejected | W009 / W010 | Only the malformed literal is recorded; no repaired value is stored. | T04 | No nearest-match behavior. |
| R05 | `DRMCP-WORK-MCP-004`; `DRMCP-WORK-MCP-005` | rejection | current-only | valid-looking suffix missing the app prefix | rejected | W009 / W010 | Manifest contains no app-inference hint or alternate accepted ref. | T04 | Separate from R02 by fixture description only. |
| R06 | `DRMCP-WORK-MCP-003`; `spec:product.design_records.spec_format.document_shape`; `DRMCP-WORK-SPEC-001` | rejection | current-only | YAML-front-matter current spec metadata | rejected | W-SPEC-001 / W009 | Intentional front matter is present and no H1-adjacent metadata authority is declared. | T04 | Existing stale tests are not authority. |
| R07 | `DRMCP-WORK-MCP-003` | duplicate | current-only | two current files with one canonical identity | duplicate | W009 | Manifest names every conflicting source and declares no winner. | T04 | Unrelated current records remain separate fixture inputs. |
| R08 | `DRMCP-WORK-MCP-003` | config | invalid | missing, unreadable, non-directory, escaped, or malformed current root | rejected | W009 | Manifest declares the exact intentional invalid-root condition. | T04 | No partial active index is expected. |
| R09 | `DRMCP-WORK-MCP-005`; `DRMCP-WORK-MCP-010` | config | invalid | missing, unreadable, non-directory, escaped, or malformed legacy root | rejected | W010 | Manifest declares the exact intentional invalid-root condition. | T04 | Configured legacy roots are mandatory when present. |
| R10 | `DRMCP-WORK-MCP-003`; `DRMCP-WORK-MCP-005` | duplicate | invalid | duplicate configured root declaration | duplicate | W009 / W010 | Manifest repeats one normalized root intentionally and identifies both declaration positions. | T04 | No deduplicated acceptance. |
| R11 | `DRMCP-WORK-MCP-003`; `DRMCP-WORK-MCP-005` | config | overlapping | current and legacy roots equal, nested, or otherwise overlapping | rejected | W010 | Structural checker proves the declared root paths overlap. | T04 | Physical separation is mandatory. |
| R12 | `DRMCP-WORK-MCP-005` | config | current-only | `legacy_roots` key missing | disabled | W010 | Manifest encodes `legacy_roots_state: omitted`. | T04 | Not unresolved legacy. |
| R13 | `DRMCP-WORK-MCP-005` | config | current-only | `legacy_roots` present as an empty list | disabled | W010 | Manifest encodes `legacy_roots_state: empty`. | T04 | Kept distinct from R12. |
| R14 | `DRMCP-WORK-MCP-005` | rejection | current-only | accepted current grammar with no active target | unresolved | W009 | Exact canonical input is valid and no current target exists. | T04 | Does not become unsupported. |
| R15 | `DRMCP-WORK-MCP-005` | rejection | current+legacy | approved V01 grammar with configured roots but no usable target | unresolved | W010 | Exact approved V01 input exists only in the manifest and no matching readable archive source exists. | T04 | Distinct from disabled fallback. |
| R16 | `spec:product.brewprint.compatibility.legacy_id_compatibility`; `DRMCP-WORK-MCP-005` | rejection | current+legacy | unsupported legacy family such as `V01-SPEC-*` or another non-approved family | rejected | W010 | Family is absent from the exact five-family allowlist. | T04 | `V01-SPEC-*` also has dedicated R01 coverage. |
| R17 | `DRMCP-WORK-MCP-004`; `DRMCP-WORK-MCP-010` | isolation | current+legacy | legacy record appears in normal active listing | rejected | W010 | Manifest declares a forbidden leakage observation without creating an accepted listing entry. | T04 | Runtime assertion belongs to W010. |
| R18 | `DRMCP-WORK-MCP-006`; `DRMCP-WORK-MCP-010` | isolation | current+legacy | legacy record becomes a current validation subject | rejected | W010 | Manifest declares a forbidden subject classification. | T04 | Current-to-legacy target lookup remains allowed. |
| R19 | `DRMCP-WORK-MCP-010` | isolation | current+legacy | legacy record is exposed as an authoring target | rejected | W010 | Manifest declares a forbidden authoring-target classification. | T04 | No authoring fixture is created. |
| R20 | `DRMCP-WORK-MCP-004`; `DRMCP-WORK-MCP-006` | isolation | current-only | normal listing or retrieval includes a physical path | rejected | W009 | Manifest marks physical-path fields forbidden on normal surfaces. | T04 | Source-location diagnostics remain exceptional accepted surfaces. |
| R21 | `DRMCP-WORK-SPEC-002`; `spec:product.design_records.spec_format.topics_table` | rejection | current-only | Topics row with unresolved canonical child `ref` | unresolved | W-SPEC-002 | Parent file and row exist; the referenced child fixture is intentionally absent. | T04 | Supplemental graph coverage. |
| R22 | `DRMCP-WORK-SPEC-002`; `spec:product.design_records.spec_format.topics_table` | duplicate | current-only | one child declared by two authoritative parents | duplicate | W-SPEC-002 | Manifest names both parent sources and one child ref without selecting a winner. | T04 | Supplemental graph coverage. |
| R23 | `DRMCP-WORK-SPEC-002`; `spec:product.design_records.spec_format.topics_table` | rejection | current-only | canonical Topics cycle | rejected | W-SPEC-002 | Manifest lists the complete closed edge sequence. | T04 | Runtime cycle detection remains W-SPEC-002-owned. |
| R24 | `DRMCP-WORK-MCP-003`; `DRMCP-WORK-MCP-005`; `DRMCP-WORK-MCP-010` | isolation | current+legacy | legacy record is inserted into the active current index | rejected | W010 | Manifest declares a forbidden active-index classification and retains the source only under the legacy root. | T04 | Detects index-scope leakage. |

### T02 through T04 authoring allocation

| Task | fixture authoring responsibility | excluded responsibility |
|---|---|---|
| T02 | Current app roots, app-aware sequential kinds, current specs, current-only configuration, cross-namespace relations, exact-current inputs, listing candidates, valid Topics parent-child inputs. | Legacy archive files, invalid roots, leakage, runtime assertions. |
| T03 | Separate configured legacy root, exact five approved V01 families, exact legacy retrieval inputs, current-to-legacy relations, successful fallback arrangement. | `V01-SPEC-*`, disabled fallback, leakage, runtime assertions. |
| T04 | Invalid source and identity inputs, duplicate identity, invalid/duplicate/overlapping roots, omitted and empty legacy configuration, unresolved refs, unsupported families, path leakage, archive leakage, invalid Topics graph arrangements. | Production code, existing test rewrites, runtime pass/fail assertions. |

T05 remains the manifest structural-verification, scoped validation, independent review, correction, and closure phase.

### Existing fixture and test inventory

This inventory is static.
No test was executed and no existing test result is inferred.

| file | fixture form | current and legacy separation | stale behavior found | shared fixture candidate | future runtime assertion owner |
|---|---|---|---|---|---|
| `designrecords/config_test.go` | `t.TempDir`, directory creation, direct config calls | Not separated; auto-detects `v01/records` and current roots together. | Automatic root detection, `v01/records` fallback, and inferred `V01-` namespace prefix contradict W003/W005. | Root-layout and invalid-config arrangements only after rewriting expectations. | W009 / W010 |
| `designrecords/index_test.go` | `t.TempDir` empty repository | No legacy content. | Default-config behavior may still depend on stale config implementation. | Empty current root case. | W009 |
| `designrecords/parser_index_test.go` | inline Markdown, `t.TempDir`, repository-bootstrap scan, shared write helper | Mixed current and V01 repository state in bootstrap tests. | YAML current specs, `SPEC-*`, non-app IDs, accepted `V01-SPEC-*`, and one mixed index. | Current kind files, invalid source files, and duplicate sources after authority correction. | W009 / W-SPEC-001 |
| `designrecords/list_records_test.go` | inline records, `t.TempDir`, direct `Index`, repository bootstrap | Legacy records appear in the normal listing baseline. | Broad listing, spec listing, ID ranges, detailed records, physical paths, V01 records, and `V01-SPEC-*` conflict with W004. | Compact current-list candidate corpus only. | W009; leakage cases W010 |
| `designrecords/get_records_test.go` | direct in-memory `Index` records | No physical root separation. | Old `ids`, item wrappers, failure placeholders, record paths, and non-app IDs conflict with W004. | Exact current input ordering corpus after replacing direct model authority with shared files. | W009 / W010 |
| `designrecords/get_record_test.go` | inline Markdown and `t.TempDir`; repository bootstrap | Legacy and current use one index path. | Retired public `get_record`, physical paths, first-winner duplicate behavior, YAML specs, and `V01-SPEC-*`. | Source files may be reusable only after migration into current or legacy roots. | W009 / W010 |
| `designrecords/resolve_reference_test.go` | inline Markdown, `t.TempDir`, direct `Index`, repository bootstrap | Current semantic aliases and V01 records share one index. | YAML `semantic_refs`, section aliases, path-bearing targets, bare IDs, and mixed-index legacy resolution conflict with W003/W005. | Current exact refs, unresolved inputs, duplicate targets, and separate fallback arrangements after correction. | W009 / W010 |
| `designrecords/validation_test.go` | many inline files, `t.TempDir`, direct `Index`, repository bootstrap, helper-built records | Mostly one synthetic current index; bootstrap may mix legacy assumptions. | YAML spec authority, raw path fields, old categories, kind/range request behavior, and helper-defined relation semantics include stale contracts. | Invalid current sources, relation corpora, and required-section documents; graph cases can feed retained validators. | W009 / W-SPEC-001 / W-SPEC-002 / W010 |
| `designrecordsmcp/tools_call_test.go` | direct in-memory index, `t.TempDir`, inline RPC JSON, file helper | No separate archive index in the helper. | Public `get_record` and `suggest_next_record`, broad listing, old `get_records`, path-bearing resolver target, and helper-owned response assumptions conflict with W004-W006. | MCP integration may consume shared fixture roots later; direct helper data is not fixture authority. | W009 / W010 |

No shared on-disk fixture package exists in the bounded inventory.
The dominant pattern is inline text written into a temporary directory.
Repository-bootstrap tests also treat the working repository as test data.

The following helpers currently encode behavior implicitly and must not become fixture authority:

- `writeTestFile` and `buildTestIndex`;
- `buildListRecordsTestIndex`;
- `workflowTestIndex` and direct `Index{Records: ...}` construction;
- `toolsCallTestIndex` and `toolsCallRecord`;
- repository-root discovery helpers used as bootstrap fixtures.

W008 may replace repeated source material with shared files.
W009, W010, W-SPEC-001, and W-SPEC-002 must retain operation and detector assertions in their own tests.

### Fixture root and manifest proposal

The bounded Go tests are package-local under `drmcp/src/internal/designrecords`.
No existing shared `testdata` directory was present in either inspected package directory.

Candidate package-local baseline:

```text
drmcp/src/internal/designrecords/testdata/read-baseline/
├── manifest.json
├── current/
│   ├── product/records/
│   └── drmcp/records/
├── legacy/
│   └── v01/records/
├── arrangements/
│   ├── duplicate-current/
│   ├── invalid-root/
│   ├── duplicate-root/
│   └── overlapping-root/
└── config/
    ├── current-only/
    ├── current-plus-legacy/
    ├── legacy-roots-omitted/
    └── legacy-roots-empty/
```

This path is a candidate until T02 accepts it.
It does not introduce a repository top-level fixture directory.
It follows the Go convention of package-local `testdata` consumed by tests without becoming production package input.

Physical separation rules:

- current files live only below `current/<app_namespace_lower>/records/`;
- archive files live only below `legacy/v01/records/`;
- no normal current root is equal to, beneath, or above a legacy root;
- invalid and overlap arrangements live below `arrangements/` and are never listed as normal roots;
- current-only configurations omit or explicitly empty `legacy_roots` according to the case;
- duplicate identity arrangements list every source and declare no winner.

Candidate `manifest.json` schema:

| field | required content |
|---|---|
| `schema_version` | Integer version, initially `1`. |
| `suite` | Stable suite name `current-and-legacy-read-baseline`. |
| `roots.current[]` | Stable name, `app_namespace`, repository-relative `records_root`, and normal-or-arrangement role. |
| `roots.legacy[]` | Stable name, repository-relative `records_root`, and normal-or-arrangement role. |
| `cases[].id` | Stable case ID from this matrix. |
| `cases[].fixture_class` | `current`, `legacy`, `config`, `duplicate`, `relation`, `rejection`, or `isolation`. |
| `cases[].root_arrangement` | `current-only`, `current+legacy`, `invalid`, or `overlapping`. |
| `cases[].inputs[]` | Fixture-relative path, exact ref or config subject, and optional relation role. |
| `cases[].legacy_roots_state` | `omitted`, `empty`, `configured`, or `not_applicable`. |
| `cases[].expected_classification` | `accepted`, `rejected`, `unresolved`, `disabled`, or `duplicate`. |
| `cases[].intentional_invalidity` | Null for normal cases; otherwise a stable kind and human-readable reason. |
| `cases[].runtime_owner` | W009, W010, W-SPEC-001, W-SPEC-002, or another accepted owner. |
| `cases[].planned_task` | T02, T03, or T04. |
| `cases[].notes` | Residual constraint that must not be inferred from file contents alone. |

Intentional invalidity is positive manifest data.
An invalid source, absent path, duplicated declaration, overlapping root, unsupported family, or leakage expectation must not depend on a comment or filename alone.

The fixture-local structural checker owns only:

- manifest JSON shape and enum validation;
- unique case IDs and required-case presence;
- declared path existence or intentional nonexistence;
- repository-relative `/`-separated path spelling and containment;
- current and legacy normal-root disjointness;
- exact five-family legacy allowlist and explicit `V01-SPEC-*` rejection;
- duplicate-source completeness and no-winner declaration;
- intentional-invalidity reason presence;
- manifest-to-file placement consistency;
- deterministic manifest ordering if T02 adopts an ordering rule.

The checker must not:

- call the production parser, index builder, resolver, validator, or MCP server;
- assert listing, retrieval, resolver, validation, diagnostic, or authoring runtime output;
- redefine accepted identity, format, relation, or compatibility rules;
- repair an invalid fixture into an accepted one.

### Ownership boundary

W008 owns:

- fixture files;
- fixture manifests;
- fixture root separation;
- intentional invalid fixture structure;
- fixture-local structural verification.

W008 does not own:

- production parser, index, resolver, validator, configuration, or MCP implementation;
- runtime behavior assertions;
- accepted contract redesign;
- PRODUCT compatibility policy changes;
- W-SPEC-001 per-file detector implementation;
- W-SPEC-002 Topics graph validation implementation;
- existing production-test rewrites.

Future runtime assertion ownership:

- W009 owns current-format discovery, parsing, indexing, listing, exact retrieval, current resolution, current validation, diagnostics, and normal path-hiding assertions.
- W010 owns configured legacy roots, archive indexing, exact legacy retrieval, fallback, current-to-legacy relations, disabled states, unsupported families, read-only behavior, and leakage assertions.
- W-SPEC-001 owns per-file current spec detector and validation-integration assertions.
- W-SPEC-002 owns Topics edge, exact child lookup, parent consistency, duplicate-parent, cycle, and graph-integration assertions.

### Lifecycle changes

The bounded lifecycle update is:

- this new Task was created as `in_progress`;
- `DRMCP-WORK-MCP-008.tasks` now contains `DRMCP-TASK-MCP-008-01`;
- W008 changed to `in_progress` because fixture-baseline work began;
- hub T08 changed to `in_progress` because selected-child lifecycle tracking began;
- W008 and hub T08 record only bounded start evidence;
- `DRMCP-WORK-MCP-001` was rechecked and remains unchanged because no relation contradiction was found.

W008 and hub T08 remain `in_progress` after T01 review and closure.
This Task remains `in_progress` until independent review acceptance.

### Changed-file manifest

| action | file |
|---|---|
| new | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-01-build-current-and-legacy-fixture-coverage-matrix.md` |
| modify | `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md` |
| modify | `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md` |

No fourth artifact is justified.
The matrix is T01 execution evidence rather than a durable runtime manifest contract.

Files checked but unchanged:

- `DRMCP-WORK-MCP-001`;
- W003 through W007;
- W009 and W010;
- W-SPEC-001 and W-SPEC-002;
- PRODUCT compatibility and spec-format authorities;
- all nine bounded existing test files;
- all production implementation files.

### Scoped lifecycle validation command

Run from the repository root after the final T01 evidence edit:

```powershell
@'
from pathlib import Path

paths = {
    "task": Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-01-build-current-and-legacy-fixture-coverage-matrix.md"),
    "work": Path("drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md"),
    "hub": Path("drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md"),
}
text = {name: path.read_text(encoding="utf-8") for name, path in paths.items()}

assert text["task"].startswith("# DRMCP-TASK-MCP-008-01:")
assert "- **status**: done" in text["task"]
task_lines = text["task"].splitlines()
for heading in ["## Goal", "## Work", "## Done condition", "## Verification", "## Evidence"]:
    assert task_lines.count(heading) == 1, (heading, task_lines.count(heading))
assert "- **status**: in_progress" in text["work"]
assert "  - DRMCP-TASK-MCP-008-01" in text["work"]
assert "- **status**: in_progress" in text["hub"]
print("task_shape=OK")
'@ | python -
```

Expected result:

```text
task_shape=OK
```

This command validates workflow shape and lifecycle synchronization only.
It does not validate future fixture files or production behavior.

### Targeted status command

```powershell
git status --short -- `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-01-build-current-and-legacy-fixture-coverage-matrix.md `
  drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md
```

Expected bounded output:

- the new T01 Task is untracked;
- W008 is modified;
- hub T08 is modified.

This command does not establish repository-wide clean status.

### Targeted whitespace command

```powershell
$trackedPaths = @(
  "drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md",
  "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md"
)

$untrackedPath = "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-01-build-current-and-legacy-fixture-coverage-matrix.md"

git diff --check -- $trackedPaths
$tracked_exit = $LASTEXITCODE

git diff --no-index --check -- NUL $untrackedPath
$untracked_exit = $LASTEXITCODE

"tracked_exit=$tracked_exit"
"untracked_exit=$untracked_exit"
```

Expected result:

- `tracked_exit=0`;
- `untracked_exit=1`;
- no whitespace error;
- no exit code `2` or greater.

An LF-to-CRLF working-copy warning is non-blocking when no whitespace error is reported.

### Command execution state

Repository-local commands were not executed by this assistant because the available filesystem MCP does not provide repository-local command execution.

Existing Go tests and repository-wide validation or status remain `NOT RUN`.
No repository-wide clean status is asserted.

The user externally ran the scoped lifecycle and targeted Git checks after independent review.
The externally reported closure evidence is:

- scoped lifecycle validation: `task_shape=OK`;
- targeted status: exactly the declared T01 Task, W008, and hub T08 paths were reported;
- all three declared paths were confirmed untracked with `git ls-files --error-unmatch`;
- each path was therefore checked with `git diff --no-index --check -- NUL <path>`;
- all three no-index checks returned exit `1`;
- no whitespace error was reported;
- no exit code `2` or greater was reported;
- LF-to-CRLF working-copy warnings were reported and treated as non-blocking.

The earlier mixed tracked/untracked expectation did not match the actual Git state because W008 and hub T08 were also untracked.
The all-untracked no-index checks supersede that expectation.

Independent review returned `PASS` with no blocking, major, or minor finding.
The review advisories only required the external scoped checks recorded above.
No command result beyond the user-provided output is inferred.

### Remaining ambiguities

No accepted contract ambiguity blocks T02.
The following fixture-authoring details remain intentionally deferred:

- T02 must accept or replace the candidate package-local path before creating files.
- T02 must decide whether one `manifest.json` remains sufficient or whether arrangement-specific fragments are required; one authoritative schema must still result.
- Exact fixture record numbers and titles are authoring details as long as canonical grammar and relation coverage remain unchanged.
- Source-location fixtures record portable source identity locally; exact runtime line, column, and diagnostic projection remain implementation-owner assertions.
- Existing test data may be migrated, rewritten, or deleted later by W009/W010; T01 does not select a mechanical migration strategy.

These are authoring choices, not permission to redesign W003 through W007.

### Independent review prompt

```text
C:\Users\imved\projects\brewprint

`DRMCP-TASK-MCP-008-01`の独立fixture coverage baseline reviewを行う。

ファイルは変更しないこと。
DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
sandboxへrepositoryを複製しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。
repository-wide clean statusを推測しないこと。
無制限なrepository traversalを行わないこと。

## 最初に読む

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/task-authoring.md`
3. `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
4. `product/records/spec/design-records/authoring-standards/writing-standard.md`
5. `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## Primary review records

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-008-01-build-current-and-legacy-fixture-coverage-matrix.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-08-track-current-and-legacy-fixture-baseline.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-001-current-read-baseline-and-realignment-coordination.md`
- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-design-records-mcp-contract-baseline-and-realignment.md`

## Accepted upstream and future owners

- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-003-current-discovery-and-active-index-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-005-resolver-and-configured-legacy-fallback-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-009-current-format-read-implementation.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-010-configured-legacy-archive-fallback-implementation.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md`

## PRODUCT authorities

- `product/records/work-items/spec/PRODUCT-WORK-SPEC-014-remove-v01-spec-compatibility-authority.md`
- `product/records/spec/brewprint/compatibility/legacy-id-compatibility.md`
- `product/records/spec/design-records/spec-format/index.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`

## Bounded existing-test inventory

- `drmcp/src/internal/designrecords/config_test.go`
- `drmcp/src/internal/designrecords/index_test.go`
- `drmcp/src/internal/designrecords/parser_index_test.go`
- `drmcp/src/internal/designrecords/list_records_test.go`
- `drmcp/src/internal/designrecords/get_records_test.go`
- `drmcp/src/internal/designrecords/get_record_test.go`
- `drmcp/src/internal/designrecords/resolve_reference_test.go`
- `drmcp/src/internal/designrecords/validation_test.go`
- `drmcp/src/internal/designrecordsmcp/tools_call_test.go`

Review only the T01 matrix, inventory, fixture-root and manifest proposal, ownership split, lifecycle update, and verification readiness.
Do not redesign W003 through W007.
Do not require fixture files or implementation changes from T01.
Do not treat stale existing tests as accepted authority.

Confirm at minimum:

- the matrix covers every fixture requirement in `DRMCP-REQ-MCP-001`;
- W003 through W007 accepted contracts are not contradicted;
- `V01-SPEC-*` is never an accepted fixture;
- accepted legacy families are limited to ADR, INV, REQ, WORK, and TASK;
- current and legacy roots are physically and logically separate;
- current-only configuration is explicit;
- invalid, duplicate, and overlapping root cases exist;
- missing and empty `legacy_roots` are separate disabled cases;
- disabled fallback and unresolved accepted legacy input are distinct;
- current-to-legacy relation coverage exists;
- legacy listing, validation, and authoring leakage coverage exists;
- physical-path leakage and accepted source-location exposure are distinct;
- fixture-local structural checks and runtime assertions are separated;
- W009, W010, W-SPEC-001, and W-SPEC-002 ownership is correct;
- existing tests were statically inventoried without inferring pass or fail;
- production implementation, existing tests, fixtures, and PRODUCT records were not changed;
- the changed-file manifest matches exactly the three workflow files;
- W008 and hub T08 lifecycle updates are correct;
- T01 remains `in_progress` before review acceptance;
- validation evidence is not fabricated;
- repository-wide clean status is not inferred.

When repository-local commands are available, run the scoped lifecycle command and targeted whitespace commands recorded in T01 Evidence.
Do not write the final command output back into the checked files before review.

出力形式:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. Matrix completeness assessment
8. Contract consistency assessment
9. Existing-test inventory assessment
10. Fixture root and manifest assessment
11. Ownership assessment
12. Lifecycle and changed-file assessment
13. Validation-evidence assessment
14. T01 closure readiness
15. T02 start readiness
```

### Current readiness

- Coverage matrix: complete for required and retained-validator cases.
- Existing test and fixture inventory: complete for the bounded nine-file set.
- Fixture root and manifest proposal: complete as a candidate for T02 acceptance.
- T02 through T04 allocation: complete.
- Changed-file manifest: complete.
- Independent review prompt: retained as closure evidence.
- Repository-wide validation and existing Go tests: not run and not required for T01 closure.
- Independent review: `PASS`; no blocking, major, or minor finding.
- External scoped lifecycle and whitespace checks: passed with the corrected all-untracked verification shape.
- T01 closure readiness: complete; this Task is `done`.
- T02 start readiness: ready.

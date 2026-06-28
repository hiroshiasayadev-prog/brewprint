# DRMCP-TASK-MCP-007-01: Establish validation Work Item audit baseline

- **id**: DRMCP-TASK-MCP-007-01
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-007
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-006-05
- **outputs**:
  - DRMCP-WORK-MCP-007

## Goal

Establish the audit baseline for the existing DRMCP spec-validation Work Items.

Record their current scope, stale metadata, obsolete phase assumptions, overlap with W003 through W006, residual implementation scope, and T02 decision inputs without accepting a disposition.

## Work

- Confirm the exact `DRMCP-TASK-MCP-007-*` inventory and select the next Task without guessing.
- Audit `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002` against current Work Item authoring rules.
- Compare both candidates with the accepted W003 through W006 ownership boundaries.
- Separate PRODUCT semantic authority from DRMCP execution and implementation responsibility.
- Record stale metadata, obsolete phase assumptions, implementation assumptions, and fixture assumptions.
- Record the residual scope that may still justify retention, supersession, or absorption.
- Define the evidence required before a `retain`, `supersede`, `absorb`, or `close` decision can be accepted.
- Record the future handoff requirements for `PRODUCT-WORK-SPEC-015`.
- Record the candidate changed-file manifest, recheck-only manifest, scoped verification commands, and independent review prompt.

This Task does not accept a final disposition.
It does not edit either candidate Work Item or any PRODUCT record.

## Done condition

- The exact Task inventory and canonical candidate paths are recorded.
- Both candidate Work Items have complete audit entries.
- Current metadata conformance and stale metadata are explicit.
- Obsolete phase assumptions are explicit.
- W003 through W006 overlap is explicit without reopening accepted boundaries.
- Residual parser-aware validation and Topics graph implementation scope is explicit.
- PRODUCT semantic authority and DRMCP execution responsibility remain separate.
- All four candidate disposition meanings and their evidence requirements are explicit.
- No final disposition is accepted.
- T02 receives a complete unresolved-decision list.
- The candidate changed-file manifest and recheck-only manifest are explicit.
- Scoped whitespace evidence covers the tracked and untracked T01 files.
- Independent review reports no blocking, major, or minor finding.

## Verification

- Compare the audit against `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`.
- Compare metadata and body shape against the current Work Item authoring standard.
- Compare spec-format assumptions against the current PRODUCT document-shape, Topics-table, ID-as-ref, validation-policy, and follow-up-boundary contracts.
- Compare overlap classifications against final W003 through W006 boundaries.
- Confirm that neither candidate Work Item nor any PRODUCT record changed during T01.
- Run the scoped tracked and untracked whitespace commands recorded in Evidence.
- Run an independent review before changing this Task to `done`.

## Evidence

### Exact Task and candidate discovery

The exact directory `drmcp/records/tasks/mcp/` was listed once.
No file beginning with `DRMCP-TASK-MCP-007-` existed.

The next Task is therefore `DRMCP-TASK-MCP-007-01`.
No existing Task was overwritten or inferred from a repository-wide search.

The supplied W-SPEC-001 path was canonical.
The supplied W-SPEC-002 filename was not present.

The exact directory `drmcp/records/work-items/spec/` was listed once.
The canonical W-SPEC-002 path is:

`drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md`

The supplied PRODUCT-WORK-SPEC-015 filename was not present.
The canonical path identified from the exact PRODUCT spec Work Item directory is:

`product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`

### Authority and lifecycle baseline

| concern | authority or accepted input | T01 use |
|---|---|---|
| DRMCP contract authority and reimplementation sequence | `DRMCP-ADR-MCP-001` | Keep PRODUCT semantic authority separate from DRMCP parsing, validation execution, diagnostics, and implementation. |
| Required read and validation outcome | `DRMCP-REQ-MCP-001` | Treat W007 as the disposition owner and preserve W003-W006 accepted boundaries. |
| Current Work Item shape and lifecycle | `spec:product.design_records.authoring_standards.work_item_authoring` | Audit metadata fields, canonical sections, reciprocal Requirement linkage, and task linkage. |
| Current Task shape and lifecycle | `spec:product.design_records.authoring_standards.task_authoring` | Create and track this T01 Task. |
| Current spec document validity | `spec:product.design_records.spec_format.document_shape` | Identify the current H1, metadata, contract-class, and section rules consumed by future validation. |
| Current Topics contract | `spec:product.design_records.spec_format.topics_table` | Identify the current `title/kind/ref/summary` table and authoritative-parent rules. |
| Current spec identity | `spec:product.design_records.spec_format.spec_id_as_ref` | Identify path-derived canonical refs and the absence of current alias or redirect mappings. |
| Validation rule and severity ownership | `spec:product.design_records.spec_format.validation_policy` | Preserve PRODUCT rule and severity authority while auditing implementation-owner pointers. |
| Follow-up ownership | `spec:product.design_records.spec_format.follow_up_boundary` | Confirm that the existing candidate IDs remain current PRODUCT pointers pending W007 disposition. |
| PRODUCT pointer synchronization | `PRODUCT-WORK-SPEC-015` | Defer all PRODUCT pointer edits until W007 accepts a disposition. |
| Current discovery and active-index contract | `DRMCP-WORK-MCP-003` | Preserve current source parsing, canonical identity, conflict handling, and active-index construction. |
| Query and exact-retrieval contract | `DRMCP-WORK-MCP-004` | Preserve listing and retrieval request, response, ordering, and path-hiding behavior. |
| Resolver and legacy-fallback contract | `DRMCP-WORK-MCP-005` | Preserve canonical grammar evaluation, lookup order, and configured legacy behavior. |
| Validation execution and diagnostic representation | `DRMCP-WORK-MCP-006` | Preserve validation selection, diagnostic taxonomy, severity, location, and exceptional path contracts. |

`DRMCP-WORK-MCP-006` and `DRMCP-TASK-MCP-006-05` are `done`.
W003 through W006 are accepted fixed inputs to this audit.

`DRMCP-WORK-MCP-007` moves to `in_progress` when this Task starts.

Hub `DRMCP-TASK-MCP-001-07` remains `not_started` during T01.
Its current Evidence requires both `DRMCP-WORK-MCP-007` and `PRODUCT-WORK-SPEC-015` to begin before the hub starts.
`PRODUCT-WORK-SPEC-015` remains `not_started` until W007 supplies an accepted disposition.
No hub edit is required during T01.

### Candidate audit: DRMCP-WORK-SPEC-001

| audit item | current finding |
|---|---|
| current status | `not_started`. No implementation or test evidence is recorded. |
| source Requirement | Declared through legacy metadata `requirement_refs: PRODUCT-REQ-SPEC-001`. The current required singular `source_requirement` field is absent. |
| dependencies | Body lists `PRODUCT-WORK-SPEC-001`, `PRODUCT-WORK-SPEC-002`, `PRODUCT-WORK-SPEC-006`, and a generic DRMCP redesign/reimplementation plan. |
| outputs | No current `outputs` field exists for Work Items. The body implies parser-aware validation implementation, diagnostics, and automated tests. |
| stated Goal | Implement parser-aware per-file validation for the PRODUCT-owned spec format in a future DRMCP implementation phase. |
| owned boundary | Intended scope covers real H1 counting, accepted H1 kind and title shape, H1-adjacent metadata, kind-specific sections, front-matter handling, canonical ID mismatch, and parser behavior around fenced headings. |
| implementation assumptions | Reuse or extend parser behavior; do not patch the historical DRMCP implementation; start during redesign or reimplementation. |
| fixture assumptions | Tests should distinguish real headings from fenced headings and cover metadata, front matter, and required sections. Temporary PRODUCT tooling may supply fixture and diagnostic expectations. |
| PRODUCT authority assumptions | PRODUCT owns document shape, identity, validation rules, and migration severity. The candidate points mainly to old PRODUCT Work Items rather than current canonical spec refs. |
| current metadata conformance | Nonconforming. Uses `requirement_refs`, `source_work_items`, and `task_refs` instead of `source_requirement`, `impact_refs`, and `tasks`. `task_refs` has no explicit empty list. |
| current body conformance | Nonconforming. Uses `Summary`, `Scope`, `Non-scope`, `Dependencies`, `Done condition`, and `Source records` instead of the canonical Work Item section set. |
| reciprocal linkage | `PRODUCT-REQ-SPEC-001` does not list `DRMCP-WORK-SPEC-001` in `work_items`. A retained or replacement owner needs one accepted source Requirement and reciprocal linkage. |
| stale metadata | Legacy field names; missing `impact_refs`; missing `tasks`; missing singular `source_requirement`; no current dependency representation; no current owner pointers to W003-W006. |
| obsolete phase assumptions | PRODUCT-WORK-SPEC-001, PRODUCT-WORK-SPEC-002, PRODUCT-WORK-SPEC-005, and PRODUCT-WORK-SPEC-006 are complete. W003-W006 contract realignment is also complete. The generic future-redesign gate is no longer a precise dependency. |
| W003 overlap | High for source parsing, H1-adjacent metadata parsing, accepted spec candidate handling, canonical identity derivation, and invalid-source retention. Residual validation must consume W003 output rather than redefine parsing or identity. |
| W004 overlap | None for the core validator. The candidate must not define list or exact-retrieval request, response, warning-trigger, or path-hiding behavior. |
| W005 overlap | Low. Canonical ref grammar and active-index lookup are fixed inputs. The candidate must not invoke or redefine the public resolver. |
| W006 overlap | High for validation execution, source selection, diagnostic categories, severity, response wrapper, source locations, and exceptional path exposure. Residual implementation must implement checks under W006 contracts instead of owning those contracts. |
| possible residual scope | Runtime implementation of per-file semantic checks not delivered by W003-W006; integration with `validate_records`; detector-level tests; parser reuse behind the accepted source model. |
| future PRODUCT-WORK-SPEC-015 handoff | Supply the accepted current owner ID for every validation-policy row now pointing to W-SPEC-001. Preserve PRODUCT rule text and severity semantics. |
| unresolved disposition inputs | Whether the same Work Item can be rebaselined cleanly; whether residual implementation belongs in a later current-read implementation Work Item; whether one dedicated validator Work Item remains reviewable; exact source Requirement and dependency chain. |

#### W-SPEC-001 assumption corrections required by any retained or replacement scope

- YAML front matter is not an active current metadata source.
- Front-matter presence may still produce PRODUCT-defined migration or source-format invalidity, but no retained scope may parse it as current metadata.
- Current canonical spec identity is path-derived and one-to-one.
- PRODUCT-WORK-SPEC-002 accepted no alias, redirect, or stale-ref mapping.
- Contract specs require `contract_class`; this rule is absent from the candidate scope.
- Current PRODUCT specs, not historical Work Item prose, are the semantic authority.
- Diagnostic names, severity, ordering, and location shape remain W006-owned.

### Candidate audit: DRMCP-WORK-SPEC-002

| audit item | current finding |
|---|---|
| current status | `not_started`. No implementation or test evidence is recorded. |
| source Requirement | Declared through legacy metadata `requirement_refs: PRODUCT-REQ-SPEC-001`. The current required singular `source_requirement` field is absent. |
| dependencies | Body lists `PRODUCT-WORK-SPEC-001`, `PRODUCT-WORK-SPEC-002`, `DRMCP-WORK-SPEC-001`, `PRODUCT-WORK-SPEC-006`, and a generic DRMCP redesign/reimplementation plan. |
| outputs | No current `outputs` field exists for Work Items. The body implies a Topics graph validator, actionable diagnostics, and automated graph tests. |
| stated Goal | Implement cross-file graph validation for spec topic trees declared by `## Topics` tables in a future DRMCP implementation phase. |
| owned boundary | Intended scope covers Topics table shape, child resolution, parent consistency, duplicate parent detection, cycle detection, and graph diagnostics. |
| implementation assumptions | Run per-file validation first; resolve graph edges across files; do not patch the historical DRMCP implementation. |
| fixture assumptions | Tests should cover valid graph, missing child, duplicate parent, invalid parent, parent mismatch, and cycle. Temporary PRODUCT tooling may supply cases. |
| PRODUCT authority assumptions | PRODUCT owns Topics table columns, authoritative parent semantics, canonical refs, and graph invalidity. The candidate encodes an obsolete pre-correction Topics shape. |
| current metadata conformance | Nonconforming. Uses `requirement_refs`, `source_work_items`, and `task_refs` instead of `source_requirement`, `impact_refs`, and `tasks`. `task_refs` has no explicit empty list. |
| current body conformance | Nonconforming. Uses the legacy Work Item section shape rather than the current canonical section set. |
| reciprocal linkage | `PRODUCT-REQ-SPEC-001` does not list `DRMCP-WORK-SPEC-002` in `work_items`. A retained or replacement owner needs one accepted source Requirement and reciprocal linkage. |
| stale metadata | Legacy field names; missing `impact_refs`; missing `tasks`; missing singular `source_requirement`; dependency on W-SPEC-001 encoded as `source_work_items`; no W003-W006 owner pointers. |
| obsolete phase assumptions | PRODUCT stabilization, compatibility, temporary validator, and migration gates are complete. W003-W006 contract realignment is complete. The generic redesign gate is no longer precise. |
| W003 overlap | Medium. Current spec discovery, canonical refs, source parsing, active-index construction, conflict state, and source provenance are fixed inputs to graph validation. Graph extraction must not redefine them. |
| W004 overlap | None for graph validation. The candidate must not add graph navigation behavior to normal listing or exact retrieval. |
| W005 overlap | Low. Canonical current lookup behavior is a fixed input. Graph validation should use internal active-index lookup and must not redefine public resolver orchestration. |
| W006 overlap | High for validation subject selection, current relation execution, diagnostic taxonomy, severity, ordering, source locations, and response representation. Residual graph implementation must emit W006-conformant diagnostics. |
| possible residual scope | Runtime extraction and validation of `## Topics` edges; canonical child-ref lookup; duplicate parent detection; child `parent` marker consistency; cycle detection; detector-level and integrated tests. |
| future PRODUCT-WORK-SPEC-015 handoff | Supply the accepted current owner ID for every validation-policy row now pointing to W-SPEC-002. Preserve PRODUCT rule text and severity semantics. |
| unresolved disposition inputs | Whether graph implementation remains a dedicated Work Item; whether it is absorbed into a broader validator implementation Work Item; dependency on the final per-file validator owner; exact source Requirement and dependency chain. |

#### W-SPEC-002 obsolete normative assumptions

| current candidate claim | current PRODUCT contract |
|---|---|
| Required columns are `title/kind/parent/file/summary`. | Required columns are `title/kind/ref/summary`. |
| `file` resolves the child target. | Canonical child identity is `ref`; tooling derives the path from the canonical ref. |
| Each row carries a canonical `parent`. | Row-level `parent` is not canonical and must not be required. |
| Row `parent` must match the declaring spec ID. | The declaring Index or Overview owns the row parent; the child H1-adjacent `parent` marker must match the declaring spec ID. |
| PRODUCT-WORK-SPEC-002 still needs to settle alias or redirect exceptions. | PRODUCT-WORK-SPEC-002 is `done` and accepted no current alias or redirect mapping. |
| Temporary PRODUCT tooling is a future fixture source. | PRODUCT-WORK-SPEC-006 is `done`; any retained fixture use must identify exact reusable cases rather than a future dependency. |

### W003-W006 overlap matrix

| candidate scope | W003 | W004 | W005 | W006 | residual after overlap removal |
|---|---|---|---|---|---|
| Per-file source and parser handling | Owns current candidate discovery, H1-adjacent metadata parsing, canonical identity, source invalidity, and conflicts. | No ownership. | No ownership. | Consumes invalid source states for validation. | Detector implementation for PRODUCT semantic rules over W003-parsed state. |
| Canonical spec identity mismatch | Owns path-derived identity and index addressability. | No ownership. | Owns public resolver orchestration only. | Owns diagnostic mapping and representation. | Comparison detector and integration tests without redefining identity or diagnostics. |
| H1, metadata, kind, contract class, and section validation | Supplies parsed state only. | No ownership. | No ownership. | Owns execution scope, categories, severity, wrapper, ordering, and locations. | PRODUCT-rule detector implementation and tests. |
| Topics child resolution | Supplies active-index and source state. | No ownership. | Supplies canonical lookup rules but not graph validation. | Owns validation execution and relation diagnostic representation. | Graph-edge extraction and internal exact child lookup. |
| Duplicate parent and cycle detection | Supplies sources and identity conflicts only. | No ownership. | No ownership. | Owns diagnostic representation, not graph algorithms. | Cross-file graph algorithm and tests. |
| Public read response behavior | No candidate ownership. | Sole owner for list and exact retrieval. | Sole owner for resolver outcomes and successful targets. | Owns diagnostics and exceptional locations. | None. |
| Runtime implementation | W003 is contract-only. | W004 is contract-only. | W005 is contract-only. | W006 is contract-only. | Both candidates may retain implementation value after their contract overlap is removed. |

W006 did not implement parser-aware spec validation or Topics graph validation.
W006 defined how current repository validation selects inputs and represents results.
The absence of runtime implementation evidence prevents T01 from treating either candidate as already delivered.

### Stale metadata and obsolete-assumption inventory

| item | W-SPEC-001 | W-SPEC-002 | T02 significance |
|---|---|---|---|
| legacy Work Item metadata keys | Yes. | Yes. | Any retained record requires current metadata rebaseline. |
| missing singular `source_requirement` | Yes. | Yes. | T02 must choose the source Requirement before reciprocal correction. |
| missing `impact_refs` | Yes. | Yes. | A retained or replacement record must point to PRODUCT authorities and W003-W006 inputs. |
| missing current `tasks` field | Yes. | Yes. | A retained record needs an explicit empty or populated list. |
| legacy body section shape | Yes. | Yes. | Retain requires a full canonical Work Item rewrite, not a metadata-only patch. |
| PRODUCT Requirement reciprocal link absent | Yes. | Yes. | Source Requirement selection and reciprocal update are mandatory decision inputs. |
| generic redesign/reimplementation dependency | Yes. | Yes. | Replace with exact current implementation dependency or owner relation. |
| PRODUCT stabilization still future | Yes. | Yes. | Obsolete; PRODUCT format, compatibility, temporary tooling, and migration are complete. |
| W003-W006 contracts still future | Implied. | Implied. | Obsolete; consume them as fixed inputs. |
| legacy Topics `parent/file` columns | N/A. | Yes. | Any retained graph scope requires complete rule replacement. |
| unresolved alias or redirect design | Partly implied. | Yes. | Obsolete; PRODUCT-WORK-SPEC-002 accepted no current mapping. |
| diagnostic ownership inside candidate | Yes. | Yes. | Remove from owner boundary; W006 owns representation. |
| fixture ownership | Underspecified. | Underspecified. | W007 does not assign fixture work; later implementation planning must identify the owner. |

### Candidate disposition meanings and current applicability

| option | precise meaning for W007 | W-SPEC-001 T01 assessment | W-SPEC-002 T01 assessment | evidence required before acceptance |
|---|---|---|---|---|
| `retain` | Keep the same Work Item ID and rebaseline its metadata, canonical sections, authority pointers, dependencies, boundary, and completion conditions. | Applicable. Residual per-file validator implementation exists, but overlap removal requires a substantial rewrite. | Applicable. Residual graph implementation exists, but obsolete Topics rules require a substantial rewrite. | Exact source Requirement; non-overlapping boundary; dependency chain; implementation outputs; fixture owner; reciprocal links; PRODUCT handoff target. |
| `supersede` | Replace the candidate with a new DRMCP Work Item because preserving the old record would obscure the accepted baseline or combine incompatible scopes. | Applicable. A new implementation-focused owner may be clearer than rewriting the historical migration-era record. | Applicable. A new current Topics graph owner may be clearer because the existing rule model is materially obsolete. | Replacement scope and ID strategy; explicit replacement relation and closure evidence; no duplicate authority; PRODUCT handoff target. |
| `absorb` | Move the residual scope into an existing accepted implementation Work Item and close the candidate with explicit absorption evidence. | Applicable only if a later implementation Work Item explicitly accepts per-file validator implementation and tests. | Applicable only if an existing implementation Work Item explicitly accepts graph algorithms and tests. | Exact absorbing owner; updated boundary and completion conditions; proof that scope is not lost or duplicated; PRODUCT handoff target. |
| `close` | Close the candidate without a retained or replacement implementation owner because no remaining executable scope exists or the scope is no longer required. | Not supported by current T01 evidence. No runtime implementation evidence exists and PRODUCT still requires durable validation. | Not supported by current T01 evidence. No graph implementation evidence exists and PRODUCT still points to durable graph validation. | Evidence that the implementation is already delivered, intentionally canceled by authority, or rendered unnecessary by a separate accepted contract. |

T01 accepts none of these options.
The table records applicability and evidence needs only.

### PRODUCT-WORK-SPEC-015 future handoff requirements

After W007 accepts a disposition, the handoff must contain:

- one accepted owner target for every `validation_policy` row currently pointing to `DRMCP-WORK-SPEC-001`;
- one accepted owner target for every row currently pointing to `DRMCP-WORK-SPEC-002`;
- the exact disposition of each old Work Item ID;
- confirmation that PRODUCT validation rules and severity text do not change;
- confirmation that `follow_up_boundary` points to current Work Item owners, not lifecycle Tasks;
- confirmation that no retained, replacement, or absorbing owner duplicates W003-W006;
- the earliest point at which PRODUCT-WORK-SPEC-015 may move to `in_progress`.

T01 does not edit `validation_policy`, `follow_up_boundary`, PRODUCT-WORK-SPEC-015, or PRODUCT-REQ-SPEC-001.

### T02 unresolved decision list

T02 must decide one question at a time.

1. What disposition applies to `DRMCP-WORK-SPEC-001`?
2. What disposition applies to `DRMCP-WORK-SPEC-002`?
3. What source Requirement owns each retained, replacement, or absorbing DRMCP implementation scope?
4. Which exact Work Item owns residual per-file validator implementation?
5. Which exact Work Item owns residual Topics graph implementation?
6. Must per-file validation and graph validation remain separate Work Items?
7. What exact dependency connects graph validation to per-file validation?
8. Who owns fixture authoring and automated implementation tests?
9. Which DRMCP records require reciprocal source-Requirement or dependency correction?
10. What exact target set must be handed to PRODUCT-WORK-SPEC-015?

No new implementation Work Item may be created before T02 accepts the relevant disposition.

### Candidate changed-file manifest

#### T01 actual workflow changes

| path | change |
|---|---|
| `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-01-establish-validation-work-item-audit-baseline.md` | New canonical T01 Task containing the audit baseline and review prompt. |
| `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md` | Move to `in_progress`, add T01 to `tasks`, and record the T01 opening state. |

No normative DRMCP spec changes in T01.
No candidate Work Item changes in T01.
No PRODUCT record changes in T01.
No hub T07 change in T01.

#### Later conditional DRMCP candidates

| path or record | condition |
|---|---|
| `DRMCP-WORK-SPEC-001` | Change only after T02 accepts retain, supersede, absorb, or close. |
| `DRMCP-WORK-SPEC-002` | Change only after T02 accepts retain, supersede, absorb, or close. |
| `DRMCP-REQ-MCP-001` | Change only when accepted source-Requirement or reciprocal linkage requires it. |
| `DRMCP-WORK-MCP-007` | Synchronize accepted decisions and later closure evidence. |
| `DRMCP-TASK-MCP-001-07` | Change lifecycle state only when its own two-child gate is satisfied. |
| New DRMCP Work Item | Create only when T02 accepts `supersede`; never create prospectively in T01. |

PRODUCT pointer files remain outside the W007 changed-file manifest.
They belong to PRODUCT-WORK-SPEC-015 after accepted disposition.

### Recheck-only manifest

| record | recheck purpose | T01 disposition |
|---|---|---|
| `DRMCP-WORK-MCP-003` | Current parsing, canonical identity, source state, conflict, and active-index boundary. | Recheck only; no edit. |
| `DRMCP-WORK-MCP-004` | List and exact-retrieval non-overlap boundary. | Recheck only; no edit. |
| `DRMCP-WORK-MCP-005` | Resolver and canonical lookup non-overlap boundary. | Recheck only; no edit. |
| `DRMCP-WORK-MCP-006` | Validation execution, diagnostics, severity, location, and implementation exclusion. | Recheck only; no edit. |
| `DRMCP-TASK-MCP-001-07` | Confirm lifecycle gate and selected child Work Items. | Recheck only; remains `not_started`. |
| `PRODUCT-WORK-SPEC-015` | Confirm downstream pointer-only ownership and start gate. | Authority read only; no edit. |
| `spec:product.design_records.spec_format.validation_policy` | Confirm current owner pointers and preserve rule and severity semantics. | Authority read only; no edit. |
| `spec:product.design_records.spec_format.follow_up_boundary` | Confirm current follow-up pointers. | Authority read only; no edit. |
| PRODUCT spec-format authority specs | Confirm current H1, metadata, Topics, and identity rules. | Authority read only; no edit. |
| PRODUCT-WORK-SPEC-001/002/005/006 | Confirm completed phase gates and obsolete candidate assumptions. | Historical input only; no edit. |
| `PRODUCT-REQ-SPEC-001` | Confirm current reciprocal `work_items` list. | Authority read only; no edit. |

### Scoped verification commands

Run from:

`C:\Users\imved\projects\brewprint`

```powershell
$trackedPath = "drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md"
$untrackedPath = "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-01-establish-validation-work-item-audit-baseline.md"

git diff --check -- $trackedPath
$tracked_exit = $LASTEXITCODE

git diff --no-index --check -- NUL $untrackedPath
$untracked_exit = $LASTEXITCODE

"tracked_exit=$tracked_exit"
"untracked_exit=$untracked_exit"

if ($tracked_exit -ge 2 -or $untracked_exit -ge 2) {
    throw "Whitespace verification command failed."
}

if ($tracked_exit -ne 0 -or $untracked_exit -ne 1) {
    throw "Unexpected whitespace verification exit code."
}
```

Expected results when no whitespace error exists:

- `tracked_exit=0`;
- `untracked_exit=1`;
- no whitespace error;
- no exit code `2` or greater.

LF-to-CRLF working-copy warnings are non-blocking when no whitespace error is reported.

Task and Work Item records are outside the strict spec-format validator scope.
No normative spec changed, so no spec validator command is required for T01.

Repository-local commands were not executed by this assistant.
The user supplied the scoped whitespace results after executing the recorded commands:

- `tracked_exit=0`;
- `untracked_exit=1`;
- no whitespace error reported;
- no exit code `2` or greater reported.

Repository-wide clean status is not inferred.

### Independent review result

Independent baseline review completed with verdict `PASS`.

- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- Advisories:
  - repository-local Git commands were not independently rerun because the review environment exposed filesystem access only;
  - the recorded external whitespace results were consistent with the expected tracked and untracked outcomes;
  - the pre-creation absence of another T007 Task could not be reconstructed from the current filesystem state, but the recorded limited directory-listing evidence was internally consistent.
- T02 decision-input readiness: `READY`.
- T01 closure readiness: `READY`.

The review confirmed that no disposition was accepted, W003-W006 were not reopened, PRODUCT pointers were not changed, and residual parser-aware and Topics graph implementation value remains explicit.

### Current T01 state

- Required instruction and authoring standards read: complete.
- Exact Task inventory: complete.
- Canonical W-SPEC-002 and PRODUCT-WORK-SPEC-015 paths: confirmed.
- W-SPEC-001/002 audit matrix: complete.
- W003-W006 overlap matrix: complete.
- Stale metadata and obsolete assumptions: inventoried.
- Candidate disposition applicability: recorded without acceptance.
- T02 unresolved decision list: complete.
- Candidate changed-file and recheck-only manifests: complete.
- PRODUCT records changed: none.
- Candidate Work Items changed: none.
- Normative DRMCP specs changed: none.
- Scoped whitespace evidence before closure synchronization: complete.
  - tracked Work Item check: `tracked_exit=0`;
  - untracked T01 check: `untracked_exit=1`;
  - no whitespace error reported;
  - no exit code `2` or greater reported;
  - the untracked exit code `1` is the expected difference result against `NUL`.
- Independent review: `PASS`; no blocking, major, or minor finding.
- Final disposition: intentionally not decided.
- T01 status: `done`.
- Final post-closure whitespace verification must run against the changed bytes; its result is not written back into this file to avoid invalidating the check.

### Independent T01 review prompt

```text
C:\Users\imved\projects\brewprint

`DRMCP-TASK-MCP-007-01`の独立baseline reviewを行う。

ファイルは変更しないこと。

DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
sandboxへrepositoryを複製しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。
repository-wide clean statusを推測しないこと。

無制限なrepository traversalや広範なsearchを行わず、以下のexact pathだけを必要な範囲で読むこと。

## 最初に読む

- `prompt_chappy.md`
- `product/records/spec/design-records/authoring-standards/task-authoring.md`
- `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
- `product/records/spec/design-records/authoring-standards/writing-standard.md`
- `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## Authority and planning records

- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-design-records-mcp-contract-baseline-and-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-01-establish-validation-work-item-audit-baseline.md`

## Disposition candidates

- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md`

## Accepted non-overlap boundaries

- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-003-current-discovery-and-active-index-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-005-resolver-and-configured-legacy-fallback-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`

## PRODUCT authority and handoff boundary

- `product/records/requirements/PRODUCT-REQ-SPEC-001-mcp-readable-spec-format-and-topic-tree.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-001-spec-format-contract-and-follow-up-split.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-002-path-derived-canonical-spec-refs-and-ref-first-topic-index.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-005-existing-spec-format-migration-and-restructuring.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-006-temporary-spec-format-validator-tooling.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`
- `product/records/spec/design-records/spec-format/document-shape.md`
- `product/records/spec/design-records/spec-format/topics-table.md`
- `product/records/spec/design-records/spec-format/spec-id-as-ref.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`
- `product/records/spec/design-records/spec-format/follow-up-boundary.md`

## Review scope

Review only the T01 audit baseline.
Do not decide `retain`, `supersede`, `absorb`, or `close`.
Do not create or propose a replacement implementation Work Item.
Do not edit PRODUCT owner pointers.
Do not reopen W003 through W006 accepted contracts.
Do not review runtime implementation or fixtures beyond whether their ownership remains unresolved correctly.

Confirm:

- Task inventory and selected T01 ID are correct;
- canonical candidate and PRODUCT handoff paths are correct;
- W007 is `in_progress` and lists T01;
- hub T07 correctly remains `not_started` while PRODUCT-WORK-SPEC-015 has not begun;
- each candidate audit covers status, source Requirement, dependencies, outputs, goal, boundary, implementation assumptions, fixture assumptions, PRODUCT authority, metadata conformance, stale metadata, obsolete assumptions, W003-W006 overlap, residual scope, future handoff, and unresolved inputs;
- current Work Item metadata and body-shape findings are accurate;
- PRODUCT-REQ-SPEC-001 reciprocal-link findings are accurate;
- W-SPEC-002 obsolete `parent/file` Topics assumptions are identified completely;
- W003 parsing and identity ownership is not reopened;
- W004 list and exact-retrieval ownership is not reopened;
- W005 resolver ownership is not reopened;
- W006 validation execution and diagnostic representation ownership is not reopened;
- the audit preserves residual parser-aware validator and Topics graph implementation value;
- `retain`, `supersede`, `absorb`, and `close` are defined precisely without accepting one;
- the T02 unresolved-decision list is sufficient;
- T01 changed-file and recheck-only manifests are accurate;
- no candidate Work Item, PRODUCT record, or normative DRMCP spec changed;
- Task and Work Item authoring shapes are valid.

Repository-local commands, when available:

```powershell
$trackedPath = "drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md"
$untrackedPath = "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-01-establish-validation-work-item-audit-baseline.md"

git diff --check -- $trackedPath
$tracked_exit = $LASTEXITCODE

git diff --no-index --check -- NUL $untrackedPath
$untracked_exit = $LASTEXITCODE

"tracked_exit=$tracked_exit"
"untracked_exit=$untracked_exit"
```

Expected no-whitespace-error result:

- `tracked_exit=0`
- `untracked_exit=1`
- no exit code `2` or greater

LF-to-CRLF warnings are non-blocking when no whitespace error exists.
Do not infer a clean working tree.
Do not use `git add .`.

出力形式:

1. Verdict: PASS / NEEDS REVISION
2. Blocking findings
3. Major findings
4. Minor findings
5. Advisories
6. Task and lifecycle linkage assessment
7. W-SPEC-001 audit assessment
8. W-SPEC-002 audit assessment
9. W003-W006 overlap assessment
10. Stale metadata and obsolete-assumption assessment
11. Candidate disposition-option assessment
12. PRODUCT-WORK-SPEC-015 handoff assessment
13. Changed-file manifest assessment
14. Recheck-only manifest assessment
15. T02 decision-input readiness
16. T01 closure readiness
```

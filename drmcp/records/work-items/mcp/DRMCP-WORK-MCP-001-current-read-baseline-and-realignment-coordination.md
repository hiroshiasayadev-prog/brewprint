# DRMCP-WORK-MCP-001: Current read baseline and realignment coordination

- **id**: DRMCP-WORK-MCP-001
- **status**: in_progress
- **date**: 2026-06-30
- **source_requirement**: DRMCP-REQ-MCP-001
- **impact_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-REQ-MCP-002
  - DRMCP-REQ-MCP-003
  - PRODUCT-REQ-SPEC-001
  - PRODUCT-REQ-SPEC-003
  - PRODUCT-REQ-SPEC-004
  - PRODUCT-WORK-SPEC-013
  - PRODUCT-WORK-SPEC-014
  - PRODUCT-WORK-SPEC-015
  - DRMCP-WORK-MCP-002
  - DRMCP-WORK-MCP-003
  - DRMCP-WORK-MCP-004
  - DRMCP-WORK-MCP-005
  - DRMCP-WORK-MCP-006
  - DRMCP-WORK-MCP-007
  - DRMCP-WORK-MCP-008
  - DRMCP-WORK-MCP-009
  - DRMCP-WORK-MCP-010
  - DRMCP-WORK-MCP-011
  - DRMCP-WORK-MCP-012
  - DRMCP-WORK-MCP-013
  - DRMCP-WORK-MCP-014
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002
  - spec:product.brewprint.compatibility.legacy_id_compatibility
  - spec:product.design_records.authoring_standards.spec_authoring
  - spec:product.design_records.spec_format.validation_policy
  - spec:drmcp.implementation
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.schema.overview
  - spec:drmcp.design_records_mcp.tools.overview
- **tasks**:
  - DRMCP-TASK-MCP-001-01
  - DRMCP-TASK-MCP-001-02
  - DRMCP-TASK-MCP-001-03
  - DRMCP-TASK-MCP-001-04
  - DRMCP-TASK-MCP-001-05
  - DRMCP-TASK-MCP-001-06
  - DRMCP-TASK-MCP-001-07
  - DRMCP-TASK-MCP-001-08
  - DRMCP-TASK-MCP-001-09
  - DRMCP-TASK-MCP-001-10
  - DRMCP-TASK-MCP-001-11
  - DRMCP-TASK-MCP-001-12
  - DRMCP-TASK-MCP-001-13
  - DRMCP-TASK-MCP-001-14
  - DRMCP-TASK-MCP-001-15
  - DRMCP-TASK-MCP-001-16
  - DRMCP-TASK-MCP-001-17
  - DRMCP-TASK-MCP-001-18
  - DRMCP-TASK-MCP-001-19

## Goal

Coordinate the child Work Items that establish the corrected DRMCP current-format read baseline required by `DRMCP-REQ-MCP-001`.

Keep detailed contract correction, fixtures, implementation, tests, and local review inside independently owned child Work Items.

Accept cross-owner gates, run the integrated review, and emit downstream readiness signals required by `DRMCP-ADR-MCP-001`.

## Boundary

This Work Item owns:

- the child Work Item split for the `DRMCP-REQ-MCP-001` read baseline;
- ownership and source-Requirement checks for each child Work Item;
- lifecycle tracking from child Work Item creation or selection through `done`;
- cross-owner gate acceptance for Brewprint compatibility and validation-policy pointers;
- integrated validation and independent review after child implementation completion;
- readiness signals to `DRMCP-WORK-MCP-002` for PRODUCT package-producer tracking and downstream REQ-002/REQ-003 Work Item creation.

This Work Item does not directly own:

- detailed discovery, parsing, index, query, retrieval, resolver, validation, diagnostic, or path-exposure contract changes;
- current-format or legacy-fallback fixture authoring;
- current read-baseline or legacy-fallback implementation;
- child implementation tests or local child review;
- full portable standards-package implementation;
- full authoring transaction implementation;
- workflow artifact create or update implementation;
- spec or investigation authoring implementation;
- BPDSL design or migration;
- migration of legacy YAML specs into the active tree;
- renaming issued legacy sequential IDs;
- creation or lifecycle tracking of downstream Work Items for `PRODUCT-REQ-SPEC-003`, `DRMCP-REQ-MCP-003`, or `DRMCP-REQ-MCP-002`;
- reopening `PRODUCT-ADR-SPEC-001` or `PRODUCT-WORK-SPEC-012`.

Each independently executable contract, fixture, or implementation workstream belongs to a child Work Item.
Each child Work Item owns its detailed Tasks, evidence, and local review.
The hub Task records only the child Work Item gate and evidence pointer.

`DRMCP-REQ-MCP-002` and `DRMCP-REQ-MCP-003` remain independent source requirements.
Their implementation Tasks must belong to separate Work Items with matching `source_requirement` metadata.

## Impact Scope

| ref | impact |
|---|---|
| `DRMCP-ADR-MCP-001` | Governs authority, compatibility, phase order, path hiding, and retained transaction boundaries. |
| `DRMCP-REQ-MCP-001` | Source requirement resolved by this Work Item. |
| `DRMCP-REQ-MCP-002` | Downstream authoring requirement. Its implementation begins only after required read and package gates. |
| `DRMCP-REQ-MCP-003` | Downstream portable-package consumer requirement. Its P0 Work Item must precede authoring runtime implementation. |
| `PRODUCT-REQ-SPEC-001` | Source Requirement for the PRODUCT-owned validation owner-pointer synchronization tracked through T07. |
| `PRODUCT-REQ-SPEC-003` | Owns projection sources, package production, synchronization, drift detection, and release validation outside DRMCP Work Items. |
| `PRODUCT-WORK-SPEC-013` | Dedicated PRODUCT package-producer Work Item tracked by the milestone; this Work Item emits readiness only. |
| `PRODUCT-REQ-SPEC-004` | PRODUCT-owned Requirement for removing obsolete `V01-SPEC-*` compatibility authority. |
| `PRODUCT-WORK-SPEC-014` | PRODUCT-owned child Work Item tracked through `DRMCP-TASK-MCP-001-02`. |
| `DRMCP-WORK-MCP-003` | DRMCP-owned discovery, current spec parsing, and active-index contract Work Item tracked through T03. |
| `DRMCP-WORK-MCP-004` | DRMCP-owned compact query and exact batch-retrieval contract Work Item tracked through T04. |
| `DRMCP-WORK-MCP-005` | DRMCP-owned current-first resolver and configured legacy-fallback contract Work Item tracked through T05. |
| `DRMCP-WORK-MCP-006` | DRMCP-owned validation, diagnostics, and path-exposure contract Work Item tracked through T06. |
| `DRMCP-WORK-MCP-007` | DRMCP-owned disposition and rebaseline Work Item for existing validation Work Items tracked through T07. |
| `PRODUCT-WORK-SPEC-015` | PRODUCT-owned validation owner-pointer synchronization Work Item tracked through T07. |
| `DRMCP-WORK-MCP-008` | DRMCP-owned current and legacy fixture-baseline Work Item tracked through T08. |
| `DRMCP-WORK-MCP-009` | Historical retired current-format implementation planning. It is not a completion gate for the rebuild line. |
| `DRMCP-WORK-MCP-010` | Retained configured legacy archive fallback implementation Work Item tracked through T10 after W012 completion and rebaseline. |
| `DRMCP-WORK-MCP-011` | Accepted replacement read-runtime architecture input. |
| `DRMCP-WORK-MCP-012` | Replacement current-read implementation Work Item. It remains blocked until W013 and W014 complete reviewed closure. |
| `DRMCP-WORK-MCP-013` | Responsibility-contract design hub tracked through T18. |
| `DRMCP-WORK-MCP-014` | Function-level internal-specification design hub tracked through T19 after W013. |
| `DRMCP-WORK-MCP-002` | Owns downstream child Work Item creation, lifecycle tracking, and integrated milestone closure. |
| `spec:product.brewprint.compatibility.legacy_id_compatibility` | Must remove `V01-SPEC-*` before legacy fallback implementation lands. |
| `spec:product.design_records.authoring_standards.spec_authoring` | Requires later synchronization for logical spec create selectors. Not changed by the read-baseline tasks. |
| `spec:product.design_records.spec_format.validation_policy` | Owner pointers must be coordinated before replacing validation Work Items. |
| `DRMCP-WORK-SPEC-001` | Existing per-file validation work requires disposition against the current spec format. |
| `DRMCP-WORK-SPEC-002` | Existing Topics graph work requires disposition against the current Topics contract. |
| Active DRMCP overview, schema, and tool specs | Require ownership-focused correction batches before implementation. |
| DRMCP implementation and fixtures | Must stop treating legacy behavior and passing tests as semantic authority. |

## Task flow

| phase | Task | dependency | outcome |
|---|---|---|---|
| A. Child Work Item plan | `DRMCP-TASK-MCP-001-01` | `DRMCP-ADR-MCP-001`, `DRMCP-REQ-MCP-001` | Create or select each independently owned child Work Item and record its owner, source Requirement, and gate. |
| B. Cross-owner prerequisites | `DRMCP-TASK-MCP-001-02`, `DRMCP-TASK-MCP-001-07` | Phase A | Complete Brewprint compatibility correction and validation-work disposition through their owning Work Items. |
| C. Read-contract correction | `DRMCP-TASK-MCP-001-03` through `DRMCP-TASK-MCP-001-06` | Phase A; compatibility policy fixed where required | Complete discovery, index, query, retrieval, resolver, validation, diagnostic, and path-exposure child Work Items. |
| D. Fixture baseline | `DRMCP-TASK-MCP-001-08` | Phases B-C | Complete the current-format and legacy-fallback fixture child Work Item. |
| E. Architecture closure | W011 `done` | Accepted D-001 through D-009 | Supply the replacement runtime architecture to the amended graph. |
| F. Historical graph-amendment defect | `DRMCP-TASK-MCP-001-13`; T14 and T15 blocked | Phase E | Preserve evidence that the first W012 release path omitted responsibility-contract and function-level internal-specification design. Release nothing. |
| G. Responsibility contract design | `DRMCP-TASK-MCP-001-18` | Phase E and T13 | Track W013 through work-partition investigation, child decision workflows, overall review, and closure. |
| H. Function-level internal specification | `DRMCP-TASK-MCP-001-19` | T18 and W013 reviewed closure | Track W014 through work-partition investigation, child decision workflows, overall review, and closure. |
| I. Replacement current runtime | `DRMCP-TASK-MCP-001-09` | T19 and corrected contracts and fixtures | Track blocked W012 through execution-graph authoring, implementation, review, and `done`. |
| J. Retained validation | `DRMCP-TASK-MCP-001-16` -> `DRMCP-TASK-MCP-001-17` | T09; T17 after T16 | Track W-SPEC-001, then W-SPEC-002, through reviewed completion. |
| K. Legacy fallback | `DRMCP-TASK-MCP-001-10` | T09 and its future rebaseline release | Track rebaselined W010 after the completed W012 runtime. T10 may proceed in parallel with T16 after their own graphs are reviewed and released. |
| L. Integrated verification | `DRMCP-TASK-MCP-001-11` | T09, T10, T16, and T17 | Run the sole integrated validation and independent review gate. Route every failure to its owning Work Item. |
| M. Downstream handoff | `DRMCP-TASK-MCP-001-12` | T11 | Record readiness signals for `DRMCP-WORK-MCP-002`. |

Detailed work does not execute inside lifecycle-tracking Tasks.
A lifecycle-tracking Task closes only after its named child Work Item is `done` and its completion evidence is accepted.

## Task Candidates

| Task | hub scope | delegated Work Item scope | dependency |
|---|---|---|---|
| `DRMCP-TASK-MCP-001-01` | Define the child Work Item split, owners, source Requirements, and lifecycle gates. | None. | None. |
| `DRMCP-TASK-MCP-001-02` | Track `PRODUCT-WORK-SPEC-014` as the Brewprint compatibility correction gate. | Remove `V01-SPEC-*` authority under `PRODUCT-REQ-SPEC-004`. | T01. |
| `DRMCP-TASK-MCP-001-03` | Track `DRMCP-WORK-MCP-003` as the discovery and active-index contract gate. | Correct active discovery, current spec parsing, and index separation contracts. | T01-T02. |
| `DRMCP-TASK-MCP-001-04` | Track `DRMCP-WORK-MCP-004` as the query and exact-retrieval contract gate. | Reflect the accepted compact active listing and `get_records`-only exact retrieval baseline. | T03. |
| `DRMCP-TASK-MCP-001-05` | Track `DRMCP-WORK-MCP-005` as the resolver and legacy-fallback contract gate. | Correct current-first resolver and configured legacy-fallback contracts. | T03-T04. |
| `DRMCP-TASK-MCP-001-06` | Track `DRMCP-WORK-MCP-006` as the validation and response-boundary contract gate. | Correct validation, diagnostics, and path-exposure contracts. | T03-T05. |
| `DRMCP-TASK-MCP-001-07` | Track `DRMCP-WORK-MCP-007` and `PRODUCT-WORK-SPEC-015` as the validation-work disposition gate. | Decide `DRMCP-WORK-SPEC-001/002` disposition and synchronize PRODUCT owner pointers. | T01. |
| `DRMCP-TASK-MCP-001-08` | Track `DRMCP-WORK-MCP-008` as the fixture-baseline gate. | Create current-format and legacy-fallback fixtures. | T02-T07. |
| `DRMCP-TASK-MCP-001-09` | Track `DRMCP-WORK-MCP-012` as the sole replacement current-read implementation gate. | W012 owns graph authoring, clean implementation, aggregate verification, independent review, and closure. | T03-T08 and T19; W012 T01 begins only after reviewed W013 and W014 closure. |
| `DRMCP-TASK-MCP-001-10` | Track rebaselined `DRMCP-WORK-MCP-010`. | W010 owns configured legacy fallback after completed W012 output and its own reviewed graph release. | T02, T05, T08, T09. |
| `DRMCP-TASK-MCP-001-11` | Run the sole integrated validation and independent review gate. Route every failure to its owning Work Item. | W012, W010, W-SPEC-001, and W-SPEC-002 own corrective implementation. | T07-T10, T16, T17. |
| `DRMCP-TASK-MCP-001-12` | Record downstream readiness and evidence pointers for `DRMCP-WORK-MCP-002`. | Downstream Work Item creation and lifecycle tracking remain outside this Work Item. | T11. |
| `DRMCP-TASK-MCP-001-13` | Author the amended W001 graph and exact future contracts. | No production implementation. | W011 closure. |
| `DRMCP-TASK-MCP-001-14` | Preserve the blocked review disposition for the incomplete T13 graph. | No releaseable verdict. | T13. |
| `DRMCP-TASK-MCP-001-15` | Preserve an empty release set for the stale T13 path. | Release nothing. | T14 remains blocked. |
| `DRMCP-TASK-MCP-001-16` | Track retained `DRMCP-WORK-SPEC-001`. | W-SPEC-001 owns per-file detectors. | T09. |
| `DRMCP-TASK-MCP-001-17` | Track retained `DRMCP-WORK-SPEC-002`. | W-SPEC-002 owns Topics graph validation. | T16. |
| `DRMCP-TASK-MCP-001-18` | Track `DRMCP-WORK-MCP-013`. | W013 owns responsibility-contract partitioning, child decision workflows, overall review, and closure. | T13 and W011 closure. |
| `DRMCP-TASK-MCP-001-19` | Track `DRMCP-WORK-MCP-014`. | W014 owns function-level internal-specification partitioning, child decision workflows, overall review, and closure. | T18. |

Tasks T02-T10 and T16-T19 are lifecycle-tracking Tasks, not implementation containers.
T14 and T15 remain blocked as the stale review and release path created by T13.
T18 and T19 establish the required detailed-design predecessors for W012.
T01 records the historical initial child split; T13 created W012 but omitted its required detailed-design gates.
Each delegated Work Item owns detailed Tasks, execution evidence, and local review.

## Completion Condition

This Work Item is complete when all of the following are true:

- corrected DRMCP read contracts implement the authority split in `DRMCP-ADR-MCP-001`;
- current specs use H1-adjacent metadata and path-derived `spec:` refs as the only active spec source format;
- current and legacy records use separate indexes;
- legacy fallback is disabled without configured `legacy_roots`;
- only accepted V01 sequential families resolve through exact read-only fallback;
- `V01-SPEC-*`, app-prefixless bare IDs, path inference, and fuzzy normalization are rejected;
- legacy records do not leak into normal listing, current repository validation, or authoring targets;
- physical paths are absent from normal listing and retrieval responses except for explicit diagnostic or debug surfaces;
- current-format and legacy-fallback fixtures cover accepted and rejected cases;
- W013 completes reviewed responsibility-contract design;
- W014 completes reviewed function-level internal-specification design after W013;
- W012 completes the clean replacement current runtime after W013 and W014 without treating W009 structure as authority;
- W-SPEC-001 completes retained per-file detector implementation;
- W-SPEC-002 completes retained Topics graph validation after W-SPEC-001;
- W010 completes configured legacy fallback after W012 and its own graph rebaseline;
- implementation tests validate the corrected contracts rather than legacy assumptions;
- `DRMCP-WORK-SPEC-001/002` disposition is coordinated with PRODUCT validation-policy owner pointers;
- T11 remains the sole integrated verification owner;
- an independent integrated review returns no blocking or major findings;
- `DRMCP-WORK-MCP-002` accepts the readiness signals for `PRODUCT-WORK-SPEC-013` lifecycle tracking and REQ-003/REQ-002 child Work Item creation; this Work Item does not own their creation or lifecycle tracking;
- `DRMCP-REQ-MCP-001` links this Work Item in its `work_items` metadata;
- completion evidence and any residual limitations are recorded here before status changes to `done`.

## Evidence

- `DRMCP-INV-MCP-002`: Contract audit and implementation-drift evidence.
- `DRMCP-ADR-MCP-001`: Accepted realignment baseline and implementation sequence.
- `DRMCP-REQ-MCP-001`: Source requirement for the current read baseline.
- `DRMCP-REQ-MCP-002`: Downstream authoring transaction requirement.
- `DRMCP-REQ-MCP-003`: Downstream portable standards-package consumer requirement.
- `PRODUCT-REQ-SPEC-003`: PRODUCT-owned package projection and production requirement.
- `PRODUCT-REQ-SPEC-004`: PRODUCT-owned Brewprint legacy spec compatibility correction.
- `PRODUCT-WORK-SPEC-013`: Dedicated PRODUCT package-producer Work Item handed off to milestone tracking.
- `PRODUCT-WORK-SPEC-014`: PRODUCT-owned `V01-SPEC-*` compatibility correction tracked by T02.
- `DRMCP-WORK-MCP-003`: DRMCP-owned current discovery and active-index contract realignment tracked by T03.
- `DRMCP-WORK-MCP-004`: DRMCP-owned compact query and exact batch-retrieval contract realignment tracked by T04.
  - The child Work Item is `done` and the query/exact-retrieval gate is accepted.
  - The accepted public surface uses compact active-index `list_records` and `get_records` as the sole exact-retrieval operation; `get_record` and `suggest_next_record` are retired without compatibility aliases.
  - Final scoped validator result: `[strict]  All 8 file(s) OK.`
  - Final independent re-review verdict: `PASS`; `F-MAJ-01`: `CLOSED`; no remaining blocking, major, or minor findings.
  - Validation evidence, changed-file integrity evidence, and W003 through W006 ownership separation were accepted.
  - Final post-closure no-index checks for the closure-updated T05 and W004 records returned expected exit code `1`, no whitespace errors, and LF-to-CRLF warnings only.
- `DRMCP-WORK-MCP-005`: DRMCP-owned resolver and configured legacy-fallback contract realignment tracked by T05.
  - The child Work Item is `done` and the resolver/configured-fallback gate is accepted.
  - Final normative manifest: `namespace-scanning.md`, `resolver.md`, `tools/resolve-reference.md`, and `tools/get-records.md`.
  - The accepted resolver contract is current-first; accepted legacy fallback begins only after the current stage remains unresolved.
  - `get_records` retains its operation-specific legacy-first exact classification, one lookup scope, and no resolver invocation.
  - Missing or empty `legacy_roots` disables fallback; configured roots are mandatory; legacy identity remains filename-derived and exact.
  - Rejected inputs are not repaired, inferred, redirected, or resolved through aliases, paths, sections, headings, fixtures, or obsolete prefixes.
  - Post-correction scoped validator result: `[strict]  All 4 file(s) OK.`
  - `git diff --check` reported no whitespace error; LF-to-CRLF warnings were non-blocking.
  - Initial final-review finding `F-MIN-FINAL-01` was corrected and closed.
  - Limited independent re-review verdict: `PASS`; no remaining blocking, major, or minor findings.
  - Advisories A-01 and A-02 remain non-blocking.
  - `DRMCP-TASK-MCP-001-05` is `done` and accepted the child closure evidence.
- `DRMCP-WORK-MCP-006`: DRMCP-owned validation, diagnostics, and path-exposure contract realignment tracked by T06.
- `DRMCP-WORK-MCP-007`: DRMCP-owned validation Work Item disposition and rebaseline tracked by T07.
- `PRODUCT-WORK-SPEC-015`: PRODUCT-owned validation owner-pointer synchronization tracked by T07.
- `DRMCP-WORK-MCP-008`: DRMCP-owned current and legacy fixture baseline tracked by T08.
- `DRMCP-WORK-MCP-009`: Historical retired implementation planning. It is no longer tracked as the T09 completion gate.
- `DRMCP-WORK-MCP-010`: Retained configured legacy archive fallback implementation tracked by T10 after W012 completion and rebaseline.
- `DRMCP-WORK-MCP-011`: Accepted replacement runtime architecture.
- `DRMCP-WORK-MCP-012`: Blocked replacement current-read implementation gate tracked by T09.
- `DRMCP-WORK-MCP-013`: Responsibility-contract design hub tracked by T18.
- `DRMCP-WORK-MCP-014`: Function-level internal-specification design hub tracked by T19 after T18.
- `DRMCP-WORK-SPEC-001`: Retained per-file detector owner tracked by T16.
- `DRMCP-WORK-SPEC-002`: Retained Topics graph owner tracked by T17 after T16.
- `DRMCP-WORK-MCP-002`: Owner of downstream child Work Item creation, lifecycle tracking, and integrated closure.
- `PRODUCT-ADR-SPEC-001`: Accepted semantic ownership boundary.
- 2026-06-26 hub restructuring: the Work Item now delegates independently executable contract, fixture, and implementation work to child Work Items.
- `DRMCP-TASK-MCP-001-01` through `DRMCP-TASK-MCP-001-12`: historical accepted hub graph before rebuild amendment.
- `DRMCP-TASK-MCP-001-13`: completed graph-amendment authoring contract.
- `DRMCP-TASK-MCP-001-14` and `DRMCP-TASK-MCP-001-15`: blocked stale review and release path.
- `DRMCP-TASK-MCP-001-16` and `DRMCP-TASK-MCP-001-17`: retained-validation lifecycle gates.
- `DRMCP-TASK-MCP-001-18` and `DRMCP-TASK-MCP-001-19`: responsibility-contract and internal-specification lifecycle gates.
- Production implementation, validation implementation, and integrated-review evidence: pending downstream Task execution.

### W011 architecture input

- W011 is `done` and supplies the accepted replacement architecture.
- D-001 through D-009 are `recorded`.
- W009 remains replaced and retired.
- W010 remains retained and `blocked` pending post-W012 rebaseline.
- W-SPEC-001 and W-SPEC-002 remain retained, separate, and `not_started`.

### Rebuild detailed-design and implementation state

- T13 is `done`, but its graph omitted required detailed-design gates.
- T14 is `blocked` and cannot produce a releaseable review verdict.
- T15 is `blocked` with an empty release set.
- W013 exists as the responsibility-contract design hub tracked by T18.
- W014 exists as the function-level internal-specification design hub tracked by T19 after T18.
- W012 and W012 T01 are `blocked` until W013 and W014 complete reviewed closure.
- T09 tracks W012 rather than W009 and depends on T19.
- T16 tracks W-SPEC-001.
- T17 tracks W-SPEC-002 after T16.
- T10 remains blocked pending completed W012 and its own rebaseline.
- T11 remains the sole integrated verification owner.
- T12 remains dependent on T11.
- No execution leaf is released.
- Production implementation remains blocked.
- W001 remains `in_progress`.

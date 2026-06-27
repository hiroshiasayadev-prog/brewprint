# DRMCP-WORK-MCP-001: Current read baseline and realignment coordination

- **id**: DRMCP-WORK-MCP-001
- **status**: in_progress
- **date**: 2026-06-26
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
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002
  - spec:product.brewprint.compatibility.legacy_id_compatibility
  - spec:product.design_records.authoring_standards.spec_authoring
  - spec:product.design_records.spec_format.validation_policy
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
| `DRMCP-WORK-MCP-009` | DRMCP-owned current-format read implementation Work Item tracked through T09. |
| `DRMCP-WORK-MCP-010` | DRMCP-owned configured legacy archive fallback implementation Work Item tracked through T10. |
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
| E. Implementation | `DRMCP-TASK-MCP-001-09`, `DRMCP-TASK-MCP-001-10` | Corrected contracts and fixtures | Complete current read and configured legacy-fallback implementation child Work Items. |
| F. Integrated verification | `DRMCP-TASK-MCP-001-11` | Phase E | Run hub-level validation and independent review. Route substantive corrections back to the owning child Work Item. |
| G. Downstream handoff | `DRMCP-TASK-MCP-001-12` | Phase F | Record readiness signals for `DRMCP-WORK-MCP-002`. |

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
| `DRMCP-TASK-MCP-001-09` | Track `DRMCP-WORK-MCP-009` as the current read implementation gate. | Implement and test the corrected active index and current record handling. | T03-T08. |
| `DRMCP-TASK-MCP-001-10` | Track `DRMCP-WORK-MCP-010` as the configured legacy-fallback implementation gate. | Implement and test the separate legacy archive index and fallback behavior. | T05-T09. |
| `DRMCP-TASK-MCP-001-11` | Run integrated validation and independent review. Route substantive corrections to child Work Items. | Child Work Items own all corrective implementation. | T07-T10. |
| `DRMCP-TASK-MCP-001-12` | Record downstream readiness and evidence pointers for `DRMCP-WORK-MCP-002`. | Downstream Work Item creation and lifecycle tracking remain outside this Work Item. | T11. |

Tasks T02-T10 are lifecycle-tracking Tasks, not implementation containers.
T01 selected an exact existing or newly created Work Item ID for every T02-T10 gate.
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
- implementation tests validate the corrected contracts rather than legacy assumptions;
- `DRMCP-WORK-SPEC-001/002` disposition is coordinated with PRODUCT validation-policy owner pointers;
- an independent review returns no blocking or major findings;
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
- `DRMCP-WORK-MCP-009`: DRMCP-owned current-format read implementation tracked by T09.
- `DRMCP-WORK-MCP-010`: DRMCP-owned configured legacy archive fallback implementation tracked by T10.
- `DRMCP-WORK-MCP-002`: Owner of downstream child Work Item creation, lifecycle tracking, and integrated closure.
- `PRODUCT-ADR-SPEC-001`: Accepted semantic ownership boundary.
- 2026-06-26 hub restructuring: the Work Item now delegates independently executable contract, fixture, and implementation work to child Work Items.
- `DRMCP-TASK-MCP-001-01` through `DRMCP-TASK-MCP-001-12`: accepted hub task graph. T01 selected and created the exact T02-T10 child Work Item graph on 2026-06-26.
- Implementation, fixture, validation, and independent-review evidence: pending task execution.

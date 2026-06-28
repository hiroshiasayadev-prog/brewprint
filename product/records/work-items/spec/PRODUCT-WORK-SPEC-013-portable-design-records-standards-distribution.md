# PRODUCT-WORK-SPEC-013: Portable Design Records standards distribution

- **id**: PRODUCT-WORK-SPEC-013
- **status**: done
- **date**: 2026-06-26
- **source_requirement**: PRODUCT-REQ-SPEC-003
- **impact_refs**:
  - PRODUCT-ADR-SPEC-001
  - PRODUCT-ADR-SPEC-002
  - PRODUCT-ADR-SPEC-003
  - DRMCP-ADR-MCP-001
  - DRMCP-REQ-MCP-003
  - DRMCP-WORK-MCP-001
  - DRMCP-WORK-MCP-002
  - spec:product.design_records.authoring_standards
  - spec:product.design_records.namespace_model
  - spec:product.design_records.repository_layout
  - spec:product.design_records.spec_format
  - spec:product.design_records.traceability
  - spec:product.design_records.artifact_model
- **tasks**:
  - PRODUCT-TASK-SPEC-013-01
  - PRODUCT-TASK-SPEC-013-02
  - PRODUCT-TASK-SPEC-013-03
  - PRODUCT-TASK-SPEC-013-04
  - PRODUCT-TASK-SPEC-013-05
  - PRODUCT-TASK-SPEC-013-06
  - PRODUCT-TASK-SPEC-013-07
  - PRODUCT-TASK-SPEC-013-08

## Goal

Deliver a portable, operationally standalone distribution of PRODUCT-owned Design Records semantics for consumers such as DRMCP.

Define the package as a whole-tree copy of `product/records/spec/design-records/` to `<exe-dir>/design-records/`, with only the allowed canonical spec-ref prefix rewrite from `spec:product.design_records` to `spec:design_records`.

Separate source authoring cleanup from package generation. If app-local or wiring content exists in the Design Records source tree, the producer warns during generation or check execution. When a human reviewer decides that tracked correction is required, the relevant warning output may be copied into a manually created source-authoring Requirement. Warning output is not a persistent producer-owned evidence artifact and is not a generation dependency.

## Boundary

This Work Item owns:

- the authoritative PRODUCT source tree boundary;
- emission of source-authoring warnings during generation or check execution;
- human review and optional manual creation of a source-authoring Requirement for warnings that require tracked correction, without blocking generation;
- the fixed package namespace `design_records`;
- the bundled installation location `<exe-dir>/design-records/`;
- the logical package root `design-records/`;
- the package spec root `spec:design_records`;
- the authoring guidance root `spec:design_records.authoring_standards`;
- the canonical spec-ref prefix rewrite contract;
- producer warning boundaries for duplicate, unresolved, external, or unrewritten refs;
- deterministic package generation or synchronization;
- operational generation failure handling;
- production of an initial reviewable package distribution with an operational generation result;
- producer-side portability fixtures and consumer handoff documentation.

This Work Item does not own:

- resolving app-local or wiring semantics in `product/records/spec/design-records/`;
- moving or rewriting PRODUCT semantic source specs;
- separate wiring, domain, or source-placement redesign;
- section-level content selection during package generation;
- package-time prose generalization or app-local semantic filtering;
- a package-specific manifest or wrapper overview;
- DRMCP package-root configuration or loader implementation;
- DRMCP operational loading, localized indexing, or diagnostics;
- DRMCP guidance request and response contracts;
- DRMCP authoring transaction behavior or proposal staleness;
- Brewprint legacy compatibility or archive behavior;
- relative refs, namespace remapping, or package-local aliases for the first package release;
- network package registries or remote distribution;
- BPDSL design or migration;
- bulk migration of existing Design Records.

Implementation placement for generation and warning-emission tooling must be selected without moving semantic authority out of PRODUCT.
A neutral or shared tool location is permitted only when PRODUCT remains the owner of its inputs, release policy, and expected outputs.

## Impact Scope

| ref or area | impact |
|---|---|
| `PRODUCT-REQ-SPEC-003` | Source requirement resolved by this Work Item. |
| `PRODUCT-ADR-SPEC-001` | Governs PRODUCT semantic ownership and cross-owner pointer boundaries. |
| `PRODUCT-ADR-SPEC-002` | Selects the PRODUCT-owned rebuild command, fixed paths, replacement model, build integration, test seam, and generated-artifact authority boundary. |
| `PRODUCT-ADR-SPEC-003` | Defines operational checks, warning classes, exit behavior, cleanup, recovery, concurrency, independence, and the T06 verification minimum. |
| `DRMCP-ADR-MCP-001` | Requires a portable package while leaving runtime consumption to DRMCP. |
| `DRMCP-REQ-MCP-003` | Defines the consumer expectations this package must satisfy without owning package production. |
| `DRMCP-WORK-MCP-001` | Emits read-baseline and cross-owner readiness signals; it does not track this Work Item lifecycle. |
| `DRMCP-WORK-MCP-002` | Milestone owner that tracks this Work Item through T02 and consumes its producer handoff evidence. |
| `spec:product.design_records.authoring_standards` | Primary authoring contract area copied into the package. |
| `spec:product.design_records.namespace_model` | Supplies package-required identity and namespace semantics. |
| `spec:product.design_records.repository_layout` | Supplies logical placement mappings required by authoring consumers. |
| `spec:product.design_records.spec_format` | Supplies current spec shape, kinds, metadata, and selector semantics. |
| `spec:product.design_records.traceability` | Supplies package-required relation and reference forms. |
| `spec:product.design_records.artifact_model` | Supplies artifact responsibility and authoring-selection semantics. |
| Portable package artifact | New whole-tree copied distribution with bundled installation location `<exe-dir>/design-records/`, logical package root `design-records/`, and refs rewritten into `spec:design_records`. |
| Configured package override location | Consumer-owned runtime configuration surface; this Work Item defines only the bundled default and package contract consumed there. |
| Package generation and release tooling | New or revised producer-side mechanism; physical implementation owner must not become semantic owner. |

## Task flow

| phase | dependency | outcome |
|---|---|---|
| A. Source-boundary correction | `PRODUCT-REQ-SPEC-003` | Reclassify T01's prior package-selection inventory as source-authoring audit evidence and record the whole-tree package boundary. |
| B. Package interface contract | Phase A; consumer expectations from `DRMCP-REQ-MCP-003` | Define fixed namespace, bundled root, spec roots, ref-prefix rewrite, producer warnings, localized consumer indexing behavior, and producer/consumer responsibility split. |
| C. Source-authoring warning handling | Phases A-B | Define human review and optional manual source-authoring Requirement creation for producer-reported warnings without introducing persistent warning evidence or requiring warning closure before generation. |
| D. Generation and tooling placement | Phases B-C | Select deterministic copy/rewrite behavior and a tooling location that preserves PRODUCT authority. |
| E. Generation checks and warning emission | Phases B-D | Define destination-root checks, prefix-rewrite checks, semantic warning emission, host-registry independence, working-directory independence, and operational generation failures. |
| F. Initial package production | Phases D-E | Generate the first reviewable whole-tree package with semantic warnings emitted during execution and an operational generation result. |
| G. Portability and release fixtures | Phase F | Verify deterministic output, host independence, prefix rewriting, warning reporting, localized defect handling, and operational generation success. |
| H. Independent review and handoff | Phases F-G | Review warning behavior, operational results, producer/consumer interface, and ownership boundaries; then record accepted producer evidence in `DRMCP-WORK-MCP-002` as input to T03 P0 Work Item creation. |

Contract and fixture work may proceed before the final tooling implementation.
Later generation phases do not depend on source-warning closure. The initial package must have reviewable generation, prefix rewrite, warning behavior, operational result, and portability checks.

## Task Candidates

| candidate | scope | dependency |
|---|---|---|
| T01 | Record the corrected whole-tree source boundary and retain prior detailed findings only as source-authoring audit evidence. | None. |
| T02 | Define the minimal package interface: namespace, bundled root, spec roots, ref-prefix rewrite, producer warnings, localized consumer indexing behavior, and producer/consumer split. | T01. |
| T03 | Define human review and optional manual source-authoring Requirement creation for producer-reported warnings without persistent warning evidence or required warning closure. | T01-T02. |
| T04 | Select the generation or synchronization model and implementation placement while preserving PRODUCT semantic authority. | T02-T03. |
| T05 | Define package generation checks for copy destination, prefix rewrite, warning emission, operational failures, host-registry independence, and working-directory independence. | T02-T04. |
| T06 | Implement or configure deterministic whole-tree copy/ref-rewrite generation and produce the initial package. | T03-T05. |
| T07 | Add release and portability fixtures for reproducibility, host independence, prefix rewriting, warning reporting, localized duplicate/unresolved behavior, and operational generation success. | T05-T06. |
| T08 | Run independent producer/consumer review, apply required corrections, record release evidence, update the reciprocal Requirement relation, and record the accepted producer contract and release evidence in `DRMCP-WORK-MCP-002` for T03 P0 Work Item creation. | T07. |

Generator implementation, operational generation results, and review evidence must remain separately reviewable where practical. A source-authoring Requirement is created separately only when a human reviewer decides that tracked correction is required.

## Completion Condition

This Work Item is complete when all of the following are true:

- the authoritative PRODUCT source tree boundary is explicit as `product/records/spec/design-records/`;
- the package is defined as a deterministic whole-tree copy to `<exe-dir>/design-records/`;
- the package namespace `design_records` is selected and documented;
- package spec root `spec:design_records` and guidance root `spec:design_records.authoring_standards` are documented;
- the canonical ref-prefix rewrite from `spec:product.design_records` to `spec:design_records` is defined without full-text replacement;
- deterministic generation or synchronization is implemented and reproducible from the source tree;
- manual package edits cannot become independent semantic authority;
- semantic findings are emitted as producer warnings during generation or check execution, and a human reviewer may manually create a separately owned source-authoring Requirement only when tracked correction is required;
- operational generation failures are distinguished from semantic warnings;
- generation succeeds when only semantic warnings exist;
- partial semantic defects are not silently repaired, filtered, removed, generalized, or reinterpreted by the producer;
- duplicate, unresolved, external, and unrewritten refs are emitted as warnings with localized consumer behavior;
- the package producer does not perform semantic filtering, section selection, or prose generalization;
- the initial package can be generated and operationally checked without depending on the Brewprint process working directory or host app registry;
- producer-side fixtures cover destination root, namespace, prefix rewrite, warning emission, operational failure cases, reproducibility, host independence, and working-directory independence;
- the producer/consumer interface supplies the roots, namespace, rewrite, warning boundary, partial-loading, localized-indexing, and guidance-projection information expected by `DRMCP-REQ-MCP-003`;
- an independent review reports no blocking, major, or minor findings;
- `PRODUCT-REQ-SPEC-003` lists this Work Item in `work_items`;
- the resulting package contract and release evidence are available in `DRMCP-WORK-MCP-002` for T03 creation of the DRMCP package-consumer P0 Work Item;
- final evidence identifies the exact package contents, generation command or mechanism, operational generation result, review findings, and residual limitations.

## Evidence

- `PRODUCT-REQ-SPEC-003`: Source requirement for the portable standards distribution.
- `PRODUCT-ADR-SPEC-001`: PRODUCT semantic ownership boundary.
- `DRMCP-ADR-MCP-001`: Accepted producer/consumer package direction.
- `DRMCP-REQ-MCP-003`: Consumer-side package loading and runtime integration expectations.
- `DRMCP-WORK-MCP-001`: Read-baseline and cross-owner readiness signaling.
- `DRMCP-WORK-MCP-002`: Milestone lifecycle tracking and producer-handoff consumer.
- T01's previous detailed inventory remains source-authoring audit evidence only; it is not a package-generation instruction set.
- T02 records the accepted minimal package interface: whole-tree copy, exact canonical prefix rewrite, semantic warning boundary, operational generation failures, consumer partial loading, localized duplicate/unresolved/unrewritten-ref behavior, and exact guidance list/get behavior.
- T03 defines human review and optional manual source-authoring Requirement creation only; it does not define persistent warning evidence or require concrete warning results before T05 or T06 runs.
- T04 and `PRODUCT-ADR-SPEC-002` select a PRODUCT-owned Python rebuild command, fixed source and destination paths, temporary generation with replacement and rollback, build integration, internal test seams, and the non-authoritative generated-tree boundary.
- T05 and `PRODUCT-ADR-SPEC-003` define temporary-tree checks, copy and rewrite integrity, warning classes, exit behavior, cleanup, recovery, unsupported concurrency, host independence, and the minimum T06 verification cases. Independent review passed with no blocking, major, minor, or advisory findings, and T05 is closed.
- T06 emits warnings during generation or check execution, but warning output need not be retained in Task Evidence.
- When a human reviewer decides that a warning requires tracked source correction, the relevant output may be copied into a manually created source-authoring Requirement. A new source-authoring Requirement is required only when tracked correction is necessary and no existing Requirement already covers the finding.
- Later generation phases must support warning emission and record operational generation results, but do not depend on source-warning closure.
- T06 implemented the PRODUCT-owned generator at `product/src/tools/generate_design_records_package.py`, added 27 standard-library automated tests, and integrated both tests and generation into `scripts/verify.bat`.
- T06 operational result: all 27 tests passed; the public generator produced the 34-file `bin/design-records/` package with 79 non-blocking semantic warnings and exit `0`; the complete repository-local verification flow also exited `0`.
- T06 generated-artifact result: `bin/design-records/` matched the source file count, remained covered by the existing `.gitignore` `bin/` rule, and was not staged or committed.
- T06 warning classes observed during the initial generation were `external canonical ref` and `source-authoring boundary finding`. Warning output remains non-persistent and does not block later tasks.
- T06 independent review passed with no blocking or major findings. The single minor test-coverage finding was corrected with destination-backup failure and post-publication rollback-failure tests; the 27-test suite and `scripts\verify.bat` passed afterward. The warning-detail assertion advisory is deferred to T07.
- T06 is closed. T07 release and portability fixtures and T08 final producer/consumer review and handoff remain later Work Item phases.
- T07 investigation found four release-level gaps beyond the T06 failure-injection suite: repeated byte and warning reproducibility, repository-external current-working-directory execution, exact warning shape, and successful generation with localized duplicate and unresolved defects.
- T07 uses a dedicated `product/tests/tools/test_generate_design_records_package_release.py` module with temporary focused trees. No persistent generated-package golden or committed `bin/` fixture is required.
- T07 fixture coverage includes repeated replacement generation, explicit prefix-preservation bytes, exact warning class/path ordering with semantic detail projection, unchanged warning-scan content, localized duplicate and unresolved output, isolated public-command execution, public-command operational failure, standard-library-only dependency inspection, and `scripts/verify.bat` propagation checks.
- T07 added directory traversal, direct-content re-ignore, and `*.py` re-inclusion rules for `product/tests/tools/` because the existing repository-wide `tools/` rule hid both T06 and T07 test sources. Non-Python files remain ignored, and the generated `bin/` rule remains unchanged.
- T07 scoped Git inspection passed whitespace checks and showed no `bin/design-records/` status entry.
- T07 automated verification passed: all 35 Python tests exited `0`; the public generator exited `0` with 34 files and 79 non-blocking warnings; `scripts\\verify.bat` exited `0` with final `OK`; and `git check-ignore -v bin/design-records/index.md` exited `0`.
- T07 corrected one over-coupled operational-failure assertion discovered by the first test run. The corrected fixture checks the public error boundary instead of a private wrapper operation label.
- T07 independent review returned `NEEDS REVISION` with two minor findings and no blocking or major findings.
- T07 corrected F-MIN-01 by limiting Git visibility to direct Python files under `product/tests/tools/`.
- T07 corrected F-MIN-02 by preserving warning order and exact class/path assertions while checking canonical refs, declaring path sets, and boundary group sets instead of private detail wrapper text.
- Post-correction verification passed: all 35 tests, the public generator, and `scripts\\verify.bat` exited `0`; `bin/design-records/index.md` and a non-Python test-directory probe remained ignored; and the Python release test remained visible to Git.
- T07 final review result: PASS. F-MIN-01 and F-MIN-02 are closed. T07 is done, while T08 remains open.
- T08 final producer review confirms the released implementation remains aligned with `PRODUCT-REQ-SPEC-003` and `PRODUCT-ADR-SPEC-001` through `PRODUCT-ADR-SPEC-003`. No new producer-contract conflict was found.
- T08 consolidated accepted release evidence: 34 generated files, 79 non-blocking semantic warnings, generator exit `0`, 35 passing tests with exit `0`, `scripts\\verify.bat` exit `0`, generated-artifact ignore exit `0`, scoped whitespace pass, and T07 independent re-review `PASS`.
- T08 confirmed that `PRODUCT-REQ-SPEC-003` already lists `PRODUCT-WORK-SPEC-013`; no Requirement rewrite is required.
- T08 prepared the producer handoff for `DRMCP-REQ-MCP-003` through `DRMCP-WORK-MCP-002`. The handoff fixes the source and package roots, namespace, prefix rewrite, warning boundary, operational failure boundary, generation command, portability evidence, consumer-owned scope, and residual limitations.
- No Work Item sourced from `DRMCP-REQ-MCP-003` exists in the scoped consumer Work Item set. The next consumer action is P0 Work Item creation through the `DRMCP-WORK-MCP-002` T03 sequence; T08 does not create or implement that Work Item.
- T08 residual limitations are explicit: 79 semantic warnings, no manifest, no version negotiation, no remote registry, no incremental generation, unsupported concurrent generation, unimplemented consumer integration, non-authoritative generated output, and separately owned source correction.
- T08 initial independent review returned `NEEDS REVISION` with one minor finding, F-MIN-01, and no blocking or major findings.
- F-MIN-01 was corrected by synchronizing Phase H, Task Candidate T08, and the Completion Condition with the `DRMCP-WORK-MCP-002` T03 handoff sequence and by requiring closure of blocking, major, and minor findings.
- T08 independent re-review returned `PASS`. F-MIN-01 is `CLOSED`; T07 F-MIN-01 and F-MIN-02 remain `CLOSED`; no blocking, major, or minor findings remain.
- Scoped whitespace passed after the wording-only correction. Repository-local tests, generator, and `scripts\\verify.bat` were not rerun; accepted T07 execution evidence remains valid.
- `PRODUCT-TASK-SPEC-013-08` and this Work Item are closed as `done`.
- Producer closure is complete. The next consumer-side action is P0 Work Item creation through `DRMCP-WORK-MCP-002` T03; consumer implementation is not claimed complete.

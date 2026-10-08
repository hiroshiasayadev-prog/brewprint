# DRCLI-TASK-CLI-002-10: Review integrated DRCLI Specification

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: review
- **estimate**: 1d
- **depends_on**:
  - DRCLI-TASK-CLI-002-09
- **outputs**:
  - DRCLI-TASK-CLI-002-10

## Goal

Independently judge whether the complete DRCLI Specification satisfies DRCLI-REQ-CLI-001 and is implementation-ready.

## Work

Read only:

- DRCLI-REQ-CLI-001;
- DRCLI-WORK-CLI-002;
- T02 decision Evidence;
- the complete `drcli/records/spec/design-records-cli/**` tree;
- exact PRODUCT and DRMCP source authorities needed to test traceability.

Review:

- Requirement coverage;
- PRODUCT/DRCLI/DRMCP ownership boundaries;
- semantic consistency across operation areas;
- command-to-operation consistency;
- help and invalid-usage discoverability;
- Windows/Linux and portable deployment contract;
- cwd repository selection and ignore behavior;
- guidance;
- validation/diagnostics;
- direct authoring, body stdin/arg/cache, persistence, and optional proposal behavior;
- absence of hidden implementation-time design decisions.

Do not edit repository files.

## Done condition

- Verdict is PASS, NEEDS REVISION, NOT READY, or BLOCKED.
- Every material finding has ID, severity, artifact/section, observed issue, required outcome, user-judgment need, and required owner type.
- PASS is used only when no blocking, major, or required minor finding prevents design closure.

## Verification

- Trace every DRCLI-REQ-CLI-001 Required Outcome to final Specification.
- Trace every T02 decision to final Specification.
- Check final combined state rather than individual author summaries.

## Evidence

Review completed 2026-10-08 against the final combined Specification state after T13/T14 and resumed T09 integration.

### Verdict

**NEEDS REVISION**

The integrated Specification is coherent across read/navigation, repository discovery, guidance, validation, deployment, retained state, and the T02 CLI decisions, but one major Requirement gap remains in authoring diagnostics. That gap prevents implementation planning from treating authoring failure behavior as closed.

### Reviewed scope

- DRCLI-REQ-CLI-001 and DRCLI-WORK-CLI-002.
- DRCLI-TASK-CLI-002-02 decision Evidence and DRCLI-TASK-CLI-002-09 integration Evidence.
- All 82 Markdown files under `drcli/records/spec/design-records-cli/**`.
- T13/T14 selected-repository current-source decisions as projected into the final current-record-model and CLI repository-selection contracts.
- Exact PRODUCT authorities needed for ownership/traceability, including repository discovery paths, namespace/ID grammar, canonical refs, spec format, authoring standards, and PRODUCT-REQ-SPEC-003 portable standards distribution.
- Exact DRMCP design-reference requirements needed for comparison: DRMCP-REQ-MCP-001, DRMCP-REQ-MCP-002, and DRMCP-REQ-MCP-003; authoring diagnostic/update source contracts were consulted only to test the inherited design-reference boundary.

Independent structural verification found 82 Specification files, 82 unique DRCLI Specification IDs, zero duplicate IDs, and zero unresolved internal `spec:drcli.design_records_cli...` references. No unresolved TBD/TODO contract marker was found in the final Specification tree.

### Requirement trace

| DRCLI-REQ-CLI-001 Required Outcome | review result | canonical final Specification |
|---|---|---|
| Portable deployment / shared installation / Windows-Linux parity | PASS | `cli.runtime_and_distribution`, `cli.output_and_exit_status` |
| cwd / `--repo` repository selection | PASS | `cli.repository_selection_and_traversal` |
| Git-compatible ignore / `--no-ignore` traversal | PASS | `cli.repository_selection_and_traversal`, `current_record_model.current_source_corpus` |
| Automatic records-root discovery and non-path-derived app namespace resolution | PASS | `current_record_model.current_source_corpus`, `current_record_model.current_record_scopes` |
| Read/navigation/search/reference/validation operation coverage | PASS | `operations.discovery_and_listing`, `operations.retrieval`, `operations.search`, `operations.reference_resolution`, `operations.validation` |
| Portable authoring guidance list/read | PASS | `operations.guidance` |
| Direct create/update and optional deferred proposal behavior | PASS except diagnostic projection below | `operations.authoring` |
| Inline/stdin/body-cache handling and process-independent retry state | PASS | `cli.authoring_body_input`, `operations.authoring.retry_cache`, `cli.runtime_and_distribution` |
| Process-independent proposal state and repository binding | PASS | `operations.authoring.deferred_proposal`, `cli.runtime_and_distribution` |
| Help and actionable invalid usage without MCP schema knowledge | PASS | `cli.help_and_usage_diagnostics`, `cli.command_surface` |
| PRODUCT-owned validation semantics and validation result distinctions | PASS | `operations.validation`, `diagnostics`, artifact/current-record-model contracts |
| Machine-readable validation and authoring diagnostics | **FAIL** for authoring | Finding DRCLI-T10-F-MAJ-001 |
| Output limiting does not change semantic validation outcome | PASS | `operations.validation.validation_output_limits` |

### T02 decision trace

All T02 decisions D-001 through D-014 are represented consistently in the final Specification: the seven-group command hierarchy and command-to-operation map; exact cwd/`--repo` behavior; ignore/`--no-ignore`; repository-independent help; actionable usage failures; `--body` / `--body-stdin` / `--body-cache-id`; explicit human/JSON output; stdout/stderr responsibility; exit statuses 0-4; copyable runtime distribution; process-independent body-cache/proposal state; shared installation; and Windows/Linux parity.

The T13/T14 replacement of DRMCP's configured-root assumption is also consistently projected: candidate `records` roots are discovered beneath the selected repository, ignore pruning precedes discovery, namespace evidence comes from syntactically extractable canonical record identities rather than path naming, zero/multiple namespace evidence and duplicate app roots fail deterministically, and no repository manifest or registration is introduced.

### Findings

#### DRCLI-T10-F-MAJ-001

- **severity**: major
- **affected decisions/outcomes**: DRCLI-REQ-CLI-001 Required Outcome “Validation and authoring diagnostics are available in machine-readable form”; T02 D-008 and D-010 require machine-readable JSON output and stable result/failure exit categories.
- **affected artifact/section**: `spec:drcli.design_records_cli.operations.authoring.create` — Response/Errors; `...authoring.update` — Response/Errors; `...authoring.write` — Response/Errors; `...authoring.retry_cache` — Response/Errors; `...authoring.deferred_proposal` — Response/Errors; `spec:drcli.design_records_cli.diagnostics`; `spec:drcli.design_records_cli.cli.output_and_exit_status`.
- **observed problem**: Authoring contracts enumerate human-readable failure conditions and required effects, but do not define stable machine-readable authoring failure/result classifications, diagnostic codes/context, or a normative mapping of those negative outcomes to trustworthy `contract_result_failure` versus result-preventing `execution_failure`. The shared DRCLI diagnostic catalog contains no authoring-specific diagnostic entries, and the authoring contracts do not define an alternative structured machine-readable failure envelope. As a result, two conforming implementations could expose different JSON failure identities and different exit-category treatment for the same authoring condition, including invalid authoring input, candidate validation failure, stale write guards, unknown/expired body-cache IDs, and proposal lifecycle failures.
- **required outcome**: Add a DRCLI-owned machine-readable authoring diagnostic/result contract that gives every authoring failure class a stable machine-readable identity and deterministic placement/context, and classifies each authoring negative outcome sufficiently for `cli.output_and_exit_status` to select the required process-status category. PRODUCT-owned validation findings must remain PRODUCT-owned rather than being redefined. The existing direct-write, retry-cache, proposal, and write-guard semantics must remain unchanged.
- **new user judgment required**: no. The required semantic failure conditions are already fixed by the Requirement and current authoring contracts; this is a DRCLI correction/design-completion task.
- **owner type**: correction.

### Implementation-planning readiness

**Not ready.** The major finding leaves authoring machine-readable failure semantics implementation-defined. Read/navigation and repository-discovery areas are otherwise review-ready, but the Work Item must not enter implementation planning until this finding is corrected and independently re-reviewed.

### Exact next gate

Finding-driven coordination/correction under the design-review workflow: materialize a correction task for DRCLI-T10-F-MAJ-001 and a dependent independent closure-review task. T11 closure synchronization is not eligible until that closure review passes.

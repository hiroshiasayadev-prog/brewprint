# DRCLI-TASK-CLI-002-09: Integrate complete DRCLI Specification

- **status**: done
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: authoring
- **estimate**: 1d
- **depends_on**:
  - DRCLI-TASK-CLI-002-03
  - DRCLI-TASK-CLI-002-04
  - DRCLI-TASK-CLI-002-05
  - DRCLI-TASK-CLI-002-06
  - DRCLI-TASK-CLI-002-07
  - DRCLI-TASK-CLI-002-08
  - DRCLI-TASK-CLI-002-14
- **outputs**:
  - spec:drcli.design_records_cli
  - spec:drcli.design_records_cli.operations

## Goal

Produce one navigable and cross-consistent final DRCLI Specification state from all completed authoring lanes.

## Work

- Create or finalize `drcli/records/spec/design-records-cli/index.md`.
- Create or finalize `drcli/records/spec/design-records-cli/operations/index.md`.
- Verify all topic refs, parents, naming, terminology, and ownership boundaries.
- Normalize only direct consistency defects across the completed DRCLI Specification tree.
- Verify every DRCLI-REQ-CLI-001 Required Outcome has a canonical owner.
- Verify DRMCP is referenced only as design evidence, not treated as DRCLI authority.
- Verify no MCP request schema is required for CLI use.
- Preserve all accepted semantics from T02-T08.

T09 may make bounded consistency edits across `drcli/records/spec/design-records-cli/**` because all predecessor writers are complete.
T09 must stop if a required repair needs new design judgment.

## Done condition

- The DRCLI root and operation indexes are complete.
- Every spec file has correct parent/ref navigation.
- No direct contradiction exists among predecessor outputs.
- Requirement coverage is complete.
- Exact final artifacts are ready for independent review.

## Verification

- Traverse only the DRCLI Specification tree and named source authorities.
- Run available spec-format validation when possible.
- Check for stale DRCLI refs, accidental DRMCP ownership refs, missing parents, and broken Topics entries.
- Record any judgment-requiring issue as BLOCKED rather than resolving it.

## Evidence

- Integration preconditions passed: DRCLI-TASK-CLI-002-03 through DRCLI-TASK-CLI-002-08 are all `done`.
- Traversed all 80 predecessor Specification files under `drcli/records/spec/design-records-cli/**`.
- Structural checks found no path-derived ID mismatch, no existing parent mismatch, and no missing direct-child `Topics` entry.
- Before T09 authoring, the only unresolved DRCLI spec refs are the two T09-owned records that do not yet exist: `spec:drcli.design_records_cli` and `spec:drcli.design_records_cli.operations`.
- CLI command mapping contains 19 command/action rows targeting 15 unique semantic operation refs; every mapped operation ref resolves.
- No `spec:drmcp.design_records_mcp...` ref exists in the DRCLI Specification tree. Remaining DRMCP/MCP mentions are explicit baseline evidence or statements that MCP/session/schema knowledge is not required.
- Help and usage requirements remain explicit: root/group/command help is repository-independent, command help must expose accepted operands/options and a valid example, and every defined invalid-usage class requires corrective content plus an exact scoped help route.
- PRODUCT spec-format authority was checked directly. No dedicated spec-format validator was found under repository `scripts/**` or `tools/**`; equivalent structural checks were run against `spec:product.design_records.spec_format.document_shape` and `spec:product.design_records.spec_format.topics_table`.
- **BLOCKER:** selected-repository current-source construction has no complete canonical contract. `spec:drcli.design_records_cli.current_record_model.current_source_corpus` requires an explicit app-namespace-to-records-root configuration, states that DRCLI does not auto-discover configured current sources, and delegates configuration representation/diagnostics to unspecified other DRCLI Specifications. The CLI/runtime contracts select only one repository root by cwd or `--repo`; they define no source for those app/records-root associations.
- PRODUCT authority `spec:product.design_records.repository_layout.record_discovery_paths` does not close the gap: it states that app namespace is supplied by implementation context and is not derived from the physical records-root path.
- Repair therefore requires a new observable design decision for how one selected repository yields the app namespace and records-root associations consumed by the current-record model. Choosing repository scanning/inference, a repository manifest/configuration source, explicit CLI configuration, or another mechanism would be materially new judgment and is outside T09.
- Affected consistency surface includes `current-record-model/index.md`, `current-record-model/current-source-corpus.md`, downstream configured-source terminology, and any CLI/runtime contract needed to supply the association.
- No DRCLI Specification file, DRMCP file, parent Work Item lifecycle, or Git state was modified. Root and operations indexes were not created because doing so would present an internally incomplete final Specification as integrated.

### Blocker resolution — 2026-10-08

- The historical blocker above is preserved as the reason T09 originally stopped.
- DRCLI-TASK-CLI-002-13 resolved the missing selected-repository current-source decision without reopening T09 predecessor semantics.
- DRCLI-TASK-CLI-002-14 projected T13 into `current_source_corpus`, the current-record-model root, and CLI repository traversal, then completed as the final T09 prerequisite.
- T09 resumed from that accepted state and did not introduce another product-semantic decision.
- The remaining direct consistency repair in `current-record-model/current-record-scopes.md` replaces explicit configured-source association with resolved current-source semantics.
- T09 also normalized the mechanically implied T13 source-root boundary from `<APP_NAMESPACE>/records/<ARTIFACT_DIRECTORY>/` to `<RESOLVED_RECORDS_ROOT>/<ARTIFACT_DIRECTORY>/`; this preserves T13's decision that the candidate `records` root parent directory is not app-namespace authority.
- Sequential identity consistency wording now requires agreement with the resolved app namespace rather than a preconfigured app namespace.
- Validation and ordering contracts now refer to selected-repository current-state construction and records-root discovery order rather than a configured-source/configuration order.
- The canonical diagnostic code remains `configuration_failure`; no new diagnostic classification or context contract was introduced.

### Completion outputs — 2026-10-08

- Created `drcli/records/spec/design-records-cli/index.md` as `spec:drcli.design_records_cli`.
- Created `drcli/records/spec/design-records-cli/operations/index.md` as `spec:drcli.design_records_cli.operations`.
- Bounded consistency edits were made only in:
  - `drcli/records/spec/design-records-cli/artifacts/base/definitions/identity-declaration.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/definitions/record-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/templates/identity-and-structure/sequential.md`
  - `drcli/records/spec/design-records-cli/artifacts/base/templates/identity-and-structure/tree.md`
  - `drcli/records/spec/design-records-cli/artifacts/decision/identity-and-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/investigation/identity-and-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/requirement/identity-and-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/task/identity-and-structure.md`
  - `drcli/records/spec/design-records-cli/artifacts/work-item/identity-and-structure.md`
  - `drcli/records/spec/design-records-cli/current-record-model/artifact-candidate-and-admission.md`
  - `drcli/records/spec/design-records-cli/current-record-model/current-record-scopes.md`
  - `drcli/records/spec/design-records-cli/current-record-model/identity-conflict-and-addressability.md`
  - `drcli/records/spec/design-records-cli/operations/discovery-and-listing/index.md`
  - `drcli/records/spec/design-records-cli/operations/search/search-results-and-limits.md`
  - `drcli/records/spec/design-records-cli/operations/validation/exact-record-validation.md`
  - `drcli/records/spec/design-records-cli/operations/validation/scope-validation.md`
  - `drcli/records/spec/design-records-cli/operations/validation/validation-output-limits.md`
- No implementation source, DRMCP file, or parent Work Item lifecycle was modified.
- No Git operation was performed.

### Final verification — 2026-10-08

- Traversed the completed 82-file DRCLI Specification tree, including the two T09-owned indexes.
- Path-derived ID verification passed for 82 / 82 Specification files.
- Parent verification passed with zero mismatches.
- Direct-child `Topics` coverage passed with zero missing child refs, including all five root children and all seven direct operation children.
- DRCLI-internal Specification-reference validation checked 629 reference occurrences and found zero unresolved refs.
- CLI command mapping remains 19 command/action rows targeting 15 unique semantic operation refs; every target resolves.
- Focused stale-source scan found zero remaining occurrences of the old preconfigured-source assumptions, including `<APP_NAMESPACE>/records`, configured app namespace/current source/records root, configured current-state construction, configuration-order precedence, and the old empty-app-source state.
- No `spec:drmcp.design_records_mcp...` authority ref exists in the completed DRCLI Specification tree. DRMCP remains baseline/design evidence where explicitly named; PRODUCT remains authoritative for Product-owned Design Records semantics.
- DRCLI-REQ-CLI-001 Required Outcome ownership is complete:
  - portable deployment and shared-installation behavior: `spec:drcli.design_records_cli.cli.runtime_and_distribution`;
  - repository selection and ignore-aware traversal: `spec:drcli.design_records_cli.cli.repository_selection_and_traversal`;
  - repository records-root and app-namespace discovery: `spec:drcli.design_records_cli.current_record_model.current_source_corpus`, `spec:drcli.design_records_cli.current_record_model.current_record_scopes`, and `spec:drcli.design_records_cli.operations.discovery_and_listing.record_scope_discovery`;
  - read, navigation, retrieval, search, and reference resolution: the discovery/listing, retrieval, search, and reference-resolution operation Specifications;
  - portable authoring guidance: `spec:drcli.design_records_cli.operations.guidance`;
  - create/update/write orchestration: `spec:drcli.design_records_cli.operations.authoring` and its child contracts;
  - inline/stdin/cache body handling and retry persistence: `spec:drcli.design_records_cli.cli.authoring_body_input`, `spec:drcli.design_records_cli.operations.authoring.retry_cache`, and `spec:drcli.design_records_cli.cli.runtime_and_distribution`;
  - CLI help, usage correction, output, and exit behavior: `spec:drcli.design_records_cli.cli.command_surface`, `spec:drcli.design_records_cli.cli.help_and_usage_diagnostics`, and `spec:drcli.design_records_cli.cli.output_and_exit_status`;
  - broad/exact validation and machine-readable semantic diagnostics: `spec:drcli.design_records_cli.operations.validation` and `spec:drcli.design_records_cli.diagnostics`.
- CLI use still requires no MCP schema knowledge: the help contract states this explicitly, command help exposes every accepted operand and option, required and mutually exclusive inputs, applicable global options, at least one valid invocation example, and an exact scoped help command.
- Invalid CLI usage remains actionable: unknown command/group, unknown option, missing operand/value, mutually exclusive inputs, and malformed selector combinations each require the offending input or accepted shape, correction information, and the exact scoped help route; parser-only messages are insufficient.
- PRODUCT authority checked for this completion includes repository-layout record discovery and Specification authoring/format rules. The completed integration adds no new Product-owned semantic authority to DRCLI.
- Every external PRODUCT ref newly introduced by the T09 root index resolves to an existing canonical Specification: `spec:product.design_records.repository_layout`, `spec:product.design_records.authoring_standards`, and `spec:product.design_records.traceability.artifact_refs`. Predecessor external refs were not changed by T09 consistency edits.
- No remaining issue requires new design judgment. T09 is complete and ready for the independent T10 review.

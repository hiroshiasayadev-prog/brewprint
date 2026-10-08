# DRMCP-TASK-MCP-009-30: Amend T15 verification ownership and execution graph

- **id**: DRMCP-TASK-MCP-009-30
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-009-23
- **outputs**:
  - DRMCP-WORK-MCP-009
  - DRMCP-TASK-MCP-009-15
  - DRMCP-TASK-MCP-009-16
  - DRMCP-TASK-MCP-009-17
  - DRMCP-TASK-MCP-009-30
  - DRMCP-TASK-MCP-009-31
  - DRMCP-TASK-MCP-009-32
  - DRMCP-TASK-MCP-009-33
  - DRMCP-TASK-MCP-009-34
  - DRMCP-TASK-MCP-009-35
  - DRMCP-TASK-MCP-009-36

## Goal

Correct the T15 verification ownership defect exposed by the first read-only T15 execution.

Split command execution from aggregate routing, mapping observation, independent review, and lifecycle synchronization.

## Work

### Defect classification

The first T15 execution established these observations:

- all T05 mandatory focused commands exited `0`;
- the diagnostic full-package command exited `1`;
- the report named resolver and validation failures outside the frozen T08 allowlist;
- T06 and T07 remain `not_started`;
- T08 owns the first accepted full `designrecords` package PASS after T05, T06, and T07 integration;
- the report omitted the exact failing test names required by the old T15 Evidence contract.

The observations do not establish a T05 implementation failure.

The old T15 contract was defective because it converted incomplete T06 and T07 lane failures into a T05 correction blocker despite depending only on T23.

The first T15 report is amendment-trigger evidence only.
It is not accepted as T15 execution Evidence and does not change T15 lifecycle state.

### Amended execution graph

```text
T23 focused repair Evidence complete
  -> T30 graph amendment
  -> T31 independent graph review
  -> T32 reviewed release synchronization
       |-> T33 T05 focused command execution ---------|
       |-> T34 T05 static boundary verification -> T36 final mapping observation
       |-> T35 full-package failure inventory --------|

T33 + T34 + T35 + T36 reports
  -> T15 aggregate Evidence synchronization and routing
  -> T16 independent finding and mapping review
  -> T17 T05 correction closure synchronization
```

T33, T34, and T35 may run in parallel after T32 release.
T36 runs after the T34 report exists.
All four leaves are read-only and have no shared file writer.
Their complete executor reports satisfy the T15 execution gate before lifecycle synchronization.
T15 is the sole writer that later persists leaf Evidence and lifecycle.

T15 runs only after all four reports exist.
T15 does not run Go, Git, grep, formatter, generator, or filesystem inspection commands.

### Responsibility split

| Task | responsibility | model route | writes |
|---|---|---|---|
| T30 | Author the persistent amendment graph. | ChatGPT coordinator | W009, T15-T17, T30-T36 Design Records. |
| T31 | Independently review graph executability, ownership, and exact routing maps. | Opus or equivalent strong reviewer | None during review. |
| T32 | Synchronize accepted review and release T33-T36. | Haiku | Named Design Records only. |
| T33 | Run only the three T05 mandatory focused commands. | Haiku | None. |
| T34 | Run only static absence, presence, scoped diff, protected-boundary, and whitespace checks. | Haiku | None. |
| T35 | Run one diagnostic full-package command and inventory exact failures. | Sonnet or equivalent coding model | None. |
| T36 | Observe only the frozen mapping items. | Haiku | None. |
| T15 | Persist leaf reports, route exact diagnostic names, and determine aggregate result. | Sonnet or ChatGPT coordinator | T33-T36 and T15 Design Records only. |
| T16 | Independently judge finding closure and mapping acceptance. | Opus or equivalent strong reviewer | None during review. |
| T17 | Synchronize accepted T05 closure. | Haiku | Existing closure records only. |

Haiku does not assign diagnostic ownership, accept mappings, or close findings.

### Aggregate acceptance policy

T15 overall `PASS` requires:

- T33 focused result `PASS`;
- T34 boundary result `PASS`;
- T36 mapping observation result `PASS`;
- T35 package inventory result `PASS` or `COMPLETE_FAILURE_INVENTORY`;
- every exact failing test from T35 maps through the frozen owner map below;
- no exact failing test maps to T05;
- no test name is unmapped;
- no compile, package setup, package initialization, panic, unknown, or mechanically unclassifiable failure exists.

Mapped T06, T07, and T08 test failures produce T15 diagnostic result `ROUTED_DOWNSTREAM`.
They do not decide T05 finding closure.

Any T05-mapped, unmapped, non-test, or mechanically unclassifiable failure produces T15 diagnostic result `BLOCKED`.

T08 retains ownership of the first accepted full-package PASS.

### Frozen exact diagnostic test-owner map

The map assigns execution and repair routing by persistent Task boundary.
The map does not infer semantic ownership from a failure message.

#### T05 correction lane

```text
TestListRecordsCurrentCompactFiltersDefaultsAndNoPath
TestListRecordsCurrentStatusOrderLimitAndHasMore
TestListRecordsCurrentRejectsObsoleteAndInvalidInputs
TestGetRecordsCurrentExactBatchSuccessWarningsAndNoPath
TestGetRecordsCurrentRequestErrors
TestNormalReadGetRecordsPathIsNotSerialized
```

#### T06 resolver test-file lane

```text
TestResolveReferenceSemanticAndRecordTargets
TestResolveReferenceWorkflowRecordTargets
TestResolveReferenceUsesSemanticRefsFromNonRecordSpec
TestValidateRecordsSemanticRefDiagnosticsFromNonRecordSpec
TestResolveReferenceRepositoryBootstrapIDs
TestResolveReferenceUnresolvedAmbiguousAndUnsupported
```

`TestValidateRecordsSemanticRefDiagnosticsFromNonRecordSpec` routes to T06 because T06 owns its exact test file.
The route does not assign validation semantic authority to T06.

#### T07 validation test-file lane

```text
TestValidateRecordsDiagnosticCategories
TestValidateRecordsAllowsMultipleDiagnosticsForOneRecord
TestValidateRecordsOKWhenNoErrorDiagnostics
TestValidateRecordsWorkflowStatusAndParseDiagnostics
TestValidateRecordsWorkflowMetadataStrictness
TestValidateRecordsWorkflowRelationHappyPath
TestValidateRecordsWorkflowRelationUnresolvedTargets
TestValidateRecordsWorkflowRelationInvalidTargets
TestValidateRecordsWorkflowRelationBidirectionalMismatch
TestValidateRecordsTaskSourceRequirementMismatch
TestValidateRecordsTaskDependsOnBoundary
TestValidateRecordsWorkflowRelationKindFilter
TestValidateRecordsRepositoryWorkflowRelationBootstrap
TestValidateRecordsInvestigationKnownDependencyTargetExists
TestValidateRecordsSemanticRefDeclarationDiagnostics
TestValidateRecordsAllowsRootSemanticRefDeclarations
TestValidateRecordsInvestigationReferenceDiagnostics
TestValidateRecordsInvestigationWorkflowReferenceBoundary
TestValidateRecordsInvestigationSourceRefsResolveRootSemanticRefs
TestValidateRecordsRejectsInvalidSemanticRefForms
TestValidateRecordsInfoOnlyDiagnosticsKeepOKTrue
TestValidateRecordsIgnoresYAMLRefsInInvestigationMetadata
TestValidateRecordsDuplicateTargetsDoNotAddFieldSpecificAmbiguousDiagnostics
TestValidateRecordsSpecStatusMismatchRequiresBothStatuses
TestValidateRecordsADRFilenameNumberMissingIsMismatch
TestValidateRecordsKindFilter
TestValidateRecordsIDRangeFilter
TestValidateRecordsWorkflowIDRangeFilter
TestValidateRecordsRequestErrors
TestValidateRecordsRequiredNarrativeSections
TestRequiredSectionHeadingCaseMismatchDiagnostics
```

#### T08 authoring-test migration lane

```text
TestAuthoringProposalAcceptAndLifecycle
TestAuthoringDiscardExpiredAndUnknownProposal
TestAuthoringAcceptGuardsStaleTargetAndCreateCollision
TestAuthoringAcceptPartialMultiFileWriteReportsWrittenFiles
TestAuthoringCreateIDResolutionAndRejectedNewForms
TestAuthoringCreateInputContractNormalization
TestAuthoringTaskCreateRequiresExplicitParentMetadataAndReciprocalUpdate
TestAuthoringBodySourceAndCache
TestAuthoringMetadataReplacement
TestAuthoringMetadataReplacementMissingRequiredFields
TestAuthoringMetadataFieldsReplace
TestAuthoringMetadataFieldsReplaceNoOp
TestAuthoringMetadataBlockReplaceNoOp
TestAuthoringMetadataFieldsReplaceBodyForbidden
TestAuthoringMetadataFieldsReplaceRequiredFieldValidation
TestAuthoringNamedSectionSelectors
TestAuthoringNamedSectionReplaceNoOp
TestAuthoringSectionSelectorIgnoresFrontMatterAndFenceHeadings
TestAuthoringProposalValidationIsAffectedRecordSetOnly
TestAuthoringHypotheticalIndexPreservesUnchangedSemanticRefSources
TestProposeRecordUpdateRequiredHeadingCaseFallback
TestProposeRecordCreateStatusDiagnostic
TestProposeRecordCreateReciprocalFollowUpMode
TestProposeRecordCreateReciprocalUpdateIncludedDiagnostic
TestBodyCacheReturnClassification
TestExactIDSequenceGapWarning
TestReplaceNamedSectionSpacingPreservation
TestReplaceNamedSectionSpacingLastSection
TestReplaceNamedSectionBodyHeadingStripping
TestMultiOpUpdateNormalCase
TestMultiOpUpdateDoneStateEvidenceGatePassesWithCombinedOps
TestMultiOpUpdateBackwardCompatSingleOp
TestMultiOpUpdateExclusivity
TestMultiOpUpdateEmptyOperations
TestMultiOpUpdateNeitherUpdateNorOperations
TestMultiOpConflictDuplicateMetadataField
TestMultiOpConflictMetadataBlockPlusFields
TestMultiOpConflictMultipleSectionReplace
TestMultiOpUpdateNoOp
TestMultiOpTopLevelBodyForbidden
TestDiffModeRequestParameter
TestListAuthoringGuides
TestGetAuthoringGuidance
TestGetAuthoringGuidanceErrors
TestBuildIndexIgnoresAuthoringGuides
```

`TestBuildDiffMultiHunkModify` and `TestBuildDiffCreateUsesAddDiff` remain unmapped and therefore block T15 when they fail.

### Writer boundaries

T30 may change only:

- W009;
- T15, T16, and T17;
- T30 through T36.

T31 is read-only.
T32 may synchronize only W009, T15, and T30 through T36.
T33 through T36 are read-only executors.
T15 may write only T33 through T36 Evidence and lifecycle plus T15 Evidence and lifecycle.
T16 is read-only.
T17 retains its existing closure writer boundary.

No Task in this amendment may modify production source, Go tests, fixtures, schemas, manifests, ADRs, Requirements, or Specifications.

## Done condition

- The first T15 result is classified as graph-defect evidence, not T05 implementation failure Evidence.
- T30 through T36 exist with one responsibility each.
- T33, T34, and T35 may run in parallel; T36 has an explicit complete-T34-report start gate.
- T15 depends on T33 through T36 instead of T23 directly.
- Command execution and aggregate routing are separate.
- Mapping observation and mapping acceptance are separate.
- Exact T05, T06, T07, and T08 diagnostic routing maps are persistent.
- T06 and T07 failures do not decide T05 closure.
- T08 retains the first accepted full-package PASS.
- T31 and T32 own independent review and release synchronization.
- T15, T16, and T17 remain unstarted.
- No production source or Go test changes occur.

## Verification

- Confirm every new Task ID and filename is unique.
- Confirm W009 lists T30 through T36.
- Confirm dependencies are acyclic.
- Confirm T33 through T36 depend on T32 and T36 has a complete-T34-report start gate without a lifecycle dependency cycle.
- Confirm T15 depends on T33 through T36.
- Confirm T16 depends on T15 and T17 depends on T16.
- Confirm each leaf has one responsibility and no write boundary.
- Confirm only the allowed Design Records changed.
- Run scoped Git and whitespace inspection on the allowed records.

## Evidence

```yaml
result: PASS
classification:
  first_t15_execution: graph_contract_defect_trigger
  t05_implementation_failure_established: false
  first_report_accepted_as_t15_evidence: false

amendment_graph:
  authoring: DRMCP-TASK-MCP-009-30
  independent_review: DRMCP-TASK-MCP-009-31
  release_synchronization: DRMCP-TASK-MCP-009-32
  focused_commands: DRMCP-TASK-MCP-009-33
  static_boundary: DRMCP-TASK-MCP-009-34
  package_failure_inventory: DRMCP-TASK-MCP-009-35
  mapping_observation: DRMCP-TASK-MCP-009-36
  aggregate: DRMCP-TASK-MCP-009-15
  independent_finding_review: DRMCP-TASK-MCP-009-16
  closure_synchronization: DRMCP-TASK-MCP-009-17

authoring_verification:
  task_ids_unique: true
  filenames_unique: true
  w009_lists_t30_through_t36: true
  dependency_cycle_found: false
  changed_record_scope:
    - DRMCP-WORK-MCP-009
    - DRMCP-TASK-MCP-009-15
    - DRMCP-TASK-MCP-009-16
    - DRMCP-TASK-MCP-009-17
    - DRMCP-TASK-MCP-009-30
    - DRMCP-TASK-MCP-009-31
    - DRMCP-TASK-MCP-009-32
    - DRMCP-TASK-MCP-009-33
    - DRMCP-TASK-MCP-009-34
    - DRMCP-TASK-MCP-009-35
    - DRMCP-TASK-MCP-009-36
  scoped_whitespace: PASS
  repository_wide_clean: not_checked
  line_ending_warning_only: true

implementation_started: false
t15_retried: false
t16_started: false
stage_or_commit: false
```

DRMCP authoring transaction tools were unavailable in the current tool surface.
Filesystem authoring was used under the active `drmcp/records` namespace.

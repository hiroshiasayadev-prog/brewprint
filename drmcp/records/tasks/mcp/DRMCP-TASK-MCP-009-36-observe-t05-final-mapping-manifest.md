# DRMCP-TASK-MCP-009-36: Observe T05 final mapping manifest

- **id**: DRMCP-TASK-MCP-009-36
- **status**: not_started
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-32
- **outputs**:
  - DRMCP-TASK-MCP-009-36

## Goal

Mechanically assemble the frozen complete replacement mapping from T34 boundary observations.

Do not judge contract significance or accept the mapping.

## Work

### Read

Read only:

- this Task;
- T05 provisional implementation mapping;
- the complete T34 executor report.

Do not read production source, Go tests, T15, T16, or unrelated Design Records.
Do not run commands or explore the repository.

### Observation gate

Do not start until the complete T34 executor report is supplied.
This report gate is separate from lifecycle synchronization, which remains owned by T15.

Require T34 to provide complete item-level observations for:

- every retained implementation symbol;
- every retained verification test;
- every removed path, symbol, test, and helper;
- every unchanged protected path;
- every exact T22 boundary item;
- scoped diff and whitespace.

Return `BLOCKED` when T34 is missing, incomplete, or not mechanically comparable with the frozen manifest.

Return `MISMATCH` when any frozen item is absent, present contrary to the manifest, or protected-path Evidence fails.

Return `PASS` only when every frozen item matches T34 Evidence.

### Frozen complete replacement manifest

Record the following manifest without additions, deletions, semantic reordering, reference conversion, or supplementation:

```yaml
final_implementation_mapping:
  status: proposed_for_independent_acceptance

  contract_refs:
    - DRMCP-REQ-MCP-001
    - DRMCP-WORK-MCP-009
    - DRMCP-TASK-MCP-009-05

  fixture_cases:
    - C11
    - C12
    - C16
    - R02
    - R03
    - R04
    - R05
    - R14
    - R20

  implementation:
    - path: drmcp/src/internal/designrecords/tools.go
      symbols:
        - ListCurrentRecords
        - currentListedRecord
        - GetCurrentRecords
        - currentGetRecord

    - path: drmcp/src/internal/designrecordsmcp/tools.go
      symbols:
        - Tools

  verification:
    - path: drmcp/src/internal/designrecords/list_records_test.go
      tests:
        - TestListRecordsCurrentCompactFiltersDefaultsAndNoPath
        - TestListRecordsCurrentStatusOrderLimitAndHasMore
        - TestListRecordsCurrentRejectsObsoleteAndInvalidInputs

    - path: drmcp/src/internal/designrecords/get_records_test.go
      tests:
        - TestGetRecordsCurrentExactBatchSuccessWarningsAndNoPath
        - TestGetRecordsCurrentRequestErrors
        - TestNormalReadGetRecordsPathIsNotSerialized

    - path: drmcp/src/internal/designrecordsmcp/tools_call_test.go
      tests:
        - TestToolsCallSuccess
        - TestToolsListWorkflowKindEnums
        - TestToolsCallToolErrors

  removed:
    paths:
      - drmcp/src/internal/designrecords/get_record_test.go
      - drmcp/src/internal/designrecords/suggest_next_record_test.go

    symbols:
      - path: drmcp/src/internal/designrecords/tools.go
        names:
          - ListRecords
          - GetRecords
          - newListRecordsScope
          - listRecordsScope
          - listRecordsScope.selectRecord

    tests:
      - path: drmcp/src/internal/designrecords/list_records_test.go
        names:
          - TestLegacyListingBasicFiltersAndResponseShape
          - TestLegacyListingStatusAndIDFilters
          - TestLegacyListingIDRangeFilter
          - TestLegacyListingWorkflowIDRangeFilter
          - TestLegacyListingRequestErrors
          - TestLegacyListingSortOrderAndLimit
          - TestLegacyListingRepositoryBootstrapQueries

      - path: drmcp/src/internal/designrecords/get_records_test.go
        names:
          - TestGetRecordsMixedKindsPartialResultsAndDuplicates
          - TestGetRecordsAllMissingIsNormalExactLookupResponse
          - TestGetRecordsRequestErrors

    helpers:
      - path: drmcp/src/internal/designrecords/list_records_test.go
        names:
          - buildListRecordsTestIndex
          - assertListRecordsErrorCode
          - listedRecordIDs
          - findListedRecord
          - latestDecisionRecordID

      - path: drmcp/src/internal/designrecordsmcp/tools.go
        names:
          - idRangeSchema

  unchanged_protected:
    - drmcp/src/internal/designrecords/types.go
    - drmcp/src/internal/designrecords/config.go
    - drmcp/src/internal/designrecords/id_range.go
    - drmcp/src/internal/designrecords/validation.go
    - drmcp/src/internal/designrecords/validation_test.go
    - drmcp/src/internal/designrecords/authoring.go
    - drmcp/src/internal/designrecords/resolver.go

  excluded_as_non_contract_significant:
    - path: drmcp/src/internal/designrecords/list_records_test.go
      reason: Retained helpers intPtr and currentReadTestIndex support the frozen tests but are not verification test identities.
    - path: drmcp/src/internal/designrecords/get_records_test.go
      reason: Retained helpers hasOperationWarning, jsonContainsKey, and jsonValueContainsKey support assertions but are not verification test identities.
    - path: drmcp/src/internal/designrecordsmcp/tools_call.go
      reason: Generic dispatch glue is unchanged by this correction; operation symbols and adapter tests carry the correction mapping.
    - path: drmcp/src/internal/designrecords/resolve_reference_test.go
      reason: T22 removes only stale retired list and get assertions. The surviving resolver assertions remain T06-owned, are not executed by T22 or T33, and are not T05 contract-significant verification.
    - path: drmcp/src/internal/designrecordsmcp/jsonrpc_test.go
      reason: Unchanged protocol coverage is outside the correction-significant retained test set.
    - path: drmcp/src/cmd/design-records-mcp/main_test.go
      reason: Unchanged command bootstrap coverage is outside the correction-significant retained test set.

  future_canonicalization:
    internal_design_ref: pending
    bpdsl_ref: pending
```

### Observation result

Record item-level `observed: true | false` values separately from the frozen manifest.
Do not alter the manifest to reflect an observation failure.

T15 may copy the manifest only when T36 result is `PASS` and T34 result is `PASS`.
T16 remains the only mapping acceptance owner.

### Prohibited operations

- Do not modify files.
- Do not run Go, Git, grep, formatter, generator, or filesystem discovery commands.
- Do not inspect source or test meaning.
- Do not change contract refs, fixture cases, paths, symbols, tests, removals, exclusions, reasons, or future-canonicalization values.
- Do not accept the mapping or close findings.
- Do not update Design Records.
- Do not stage or commit.

### Output

1. Result: `PASS`, `MISMATCH`, or `BLOCKED`.
2. T34 Evidence completeness.
3. Item-level observations.
4. Exact frozen manifest.
5. T15 mapping readiness.

## Done condition

- T34 Evidence is mechanically compared with every manifest item it owns.
- The exact frozen manifest is reproduced without discretionary change.
- Observation result is explicit.
- No contract-significance or acceptance judgment occurs.
- No command or file change occurs.

## Verification

Compare only T34 item-level Evidence with the frozen manifest in this Task.

## Evidence

Execution pending the complete T34 executor report after reviewed release by T32.

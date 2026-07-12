# USDM requirement: Validation result projection

- **id**: `usdm:drmcp.mcp_capabilities.validation.validation_result_projection`
- **status**: draft
- **date**: 2026-07-12
- **kind**: requirement
- **parent**: `usdm:drmcp.mcp_capabilities.validation`

## What this is

Product requirements for projecting broad-scope validation summaries, exact-ref validation findings, admission failures, deterministic ordering, bounded output, and validation failure boundaries.

## Requirements: Validation result modes
> source: literal

| id | requirement | notes |
|---|---|---|
| R001 | DRMCP must allow broad-scope validation and detailed validation of one or more exact current canonical refs to be selected and executed independently. | Broad scopes include repository, app, artifact-kind, and domain scopes. Concrete operation boundaries are defined by downstream Specifications. |
| R002 | DRMCP must make contract-violation findings distinguishable from advisory findings that do not make a validated record nonconforming. | Exact classification and severity vocabulary are defined by downstream Specifications. |

## Requirements: Broad-scope result projection
> source: literal

| id | requirement | notes |
|---|---|---|
| R003 | DRMCP must allow a broad-scope validation result to include either only scope-level aggregates or scope-level aggregates together with summaries for current records in which an abnormal condition was detected. | Conforming record refs are not listed by broad-scope validation. |
| R004 | Each returned record summary must identify the current record by canonical ref and expose its abnormal state, contract-violation finding count, and advisory finding count. | Exact state vocabulary is defined by downstream Specifications. |
| R006 | Scope-level aggregates must include the number of validated records, records with detected abnormalities, contract-violation findings, and advisory findings. | Aggregate values describe the complete selected scope rather than only returned record summaries. |
| R007 | DRMCP must report discovered sources that could not be admitted as current records and current identity conflicts separately from summaries of validated records. | These conditions occur before detailed record validation. |
| R008 | DRMCP must identify each unadmitted source by a repository-root-relative path and each identity conflict by its claimed canonical ref and the repository-root-relative paths of every conflicting source. | Absolute physical paths are not required. |
| R009 | DRMCP must not report a broad scope containing an unadmitted source or identity conflict as having completed without an abnormal condition. | Exact overall-outcome representation is defined by downstream Specifications. |
| R010 | Broad-scope validation must not be required to return detailed findings for each abnormal record, and DRMCP must make those detailed findings available through exact-ref validation. | Unadmitted sources and identity conflicts are not selectable for detailed validation by source path. |

## Requirements: Exact-ref result projection
> source: literal

| id | requirement | notes |
|---|---|---|
| R011 | DRMCP must return detailed validation findings for every exact current canonical ref that is successfully selected as one uniquely selectable current record. |  |
| R012 | Each detailed finding must identify the affected record and, when applicable, the affected metadata field, H2 section, reference target, or structural relationship. | Concrete diagnostic association schemas are defined by downstream Specifications. |
| R013 | When some supplied exact-ref selectors are malformed, unresolved, conflicted, or duplicated, DRMCP must preserve the detailed validation results for every selector that successfully selects a current record. |  |
| R014 | DRMCP must return warnings for unsuccessful and duplicate exact-ref selectors after the successfully selected records' detailed validation results. | Concrete response collection structure is defined by downstream Specifications. |
| R015 | When the same exact ref is supplied more than once, DRMCP must validate the selected record once, keep the first occurrence effective, and aggregate the later occurrences into one duplicate-selector warning. |  |

## Requirements: Deterministic and bounded output
> source: literal

| id | requirement | notes |
|---|---|---|
| R016 | DRMCP must return validation results in a deterministic order for the same request and current-record state. | Exact ordering keys are defined by downstream Specifications. |
| R017 | DRMCP must suppress a semantic finding that is detected through more than one validation path so that it is returned once. |  |
| R018 | DRMCP must not suppress distinct findings that differ by affected record, field, section, reference target, structural relationship, source, or identity conflict. |  |
| R019 | DRMCP must allow the caller to limit the amount of validation result output and must apply a default limit and a server-side maximum. | Measurement units and numeric values are defined by downstream Specifications. |
| R020 | DRMCP must apply broad-scope output limits to complete record summaries and exact-ref output limits to complete findings, without returning a partially truncated summary or finding. |  |
| R021 | When an output limit omits part of a result, DRMCP must make the existence of additional results and the applied limit identifiable. |  |
| R022 | Scope-level aggregates and the overall validation outcome must reflect every detected result in the selected scope rather than only the subset returned within the output limit. |  |

## Requirements: Failure boundary
> source: literal

| id | requirement | notes |
|---|---|---|
| R023 | DRMCP must make malformed requests, invalid broad-scope selectors, configuration failures, and execution failures distinguishable from normal validation findings and exact-ref selector warnings. |  |
| R024 | When DRMCP cannot construct the selected scope, current-record state, or required validation input as a trustworthy complete state, it must not represent a partial result as a complete normal validation result. |  |

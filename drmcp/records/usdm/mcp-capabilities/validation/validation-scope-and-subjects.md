# USDM requirement: Validation scope and subjects

- **id**: `usdm:drmcp.mcp_capabilities.validation.validation_scope_and_subjects`
- **status**: draft
- **date**: 2026-07-12
- **kind**: requirement
- **parent**: `usdm:drmcp.mcp_capabilities.validation`

## What this is

Product requirements for selecting current-record validation scopes, admitting uniquely selectable current records, and reporting discovered sources that cannot be admitted as records.

## Requirements: Validation scope selection
> source: literal

| id | requirement | notes |
|---|---|---|
| R001 | DRMCP must support repository-wide validation across every configured app namespace. |  |
| R002 | DRMCP must support validation within one selected app namespace. |  |
| R003 | DRMCP must support validation within one selected app namespace and artifact kind. |  |
| R004 | DRMCP must support validation within one selected app namespace, sequential artifact kind, and domain namespace. | Domain selection does not apply to tree artifact kinds. |
| R005 | DRMCP must support selecting one or more exact current canonical refs for validation in one request. |  |
| R006 | One validation request must select either one repository, app, kind, or domain scope, or a collection of exact current canonical refs. | Broad scopes and exact-ref selections are not mixed in one request. |
| R007 | DRMCP must validate every uniquely selectable current record included in the selected scope. |  |
| R008 | DRMCP must not broaden a validation scope when a required selector is omitted or invalid. |  |
| R009 | Selectable app, artifact-kind, and domain scopes must remain consistent with the current scopes exposed by DRMCP discovery capabilities. | Validation does not redefine scope discovery. |
| R010 | A valid broad scope containing no uniquely selectable current record and no record-admission failure must produce a successful empty validation result. |  |
| R011 | An exact current canonical ref that identifies no uniquely selectable current record must be identifiable as an unsuccessful selector and must not be treated as a successful empty validation result. | Exact outcome representation is defined by downstream Specifications. |

## Requirements: Current record admission
> source: literal

| id | requirement | notes |
|---|---|---|
| R012 | DRMCP must admit every discovered current source that has one uniquely selectable current canonical ref within the selected broad scope as a current record for validation. |  |
| R013 | DRMCP must separately identify every discovered current source within the selected broad scope that cannot be admitted as a uniquely selectable current record. | Includes unreadable, unparseable, identity-unavailable, and identity-inconsistent sources. Exact admission-failure classification belongs to downstream Specifications. |
| R014 | When multiple current sources claim one canonical identity, DRMCP must report an identity conflict, must not admit any conflicting source as the current record for that identity, and must not select a winner. | An exact selector for the conflicted identity is not uniquely selectable. |
| R015 | DRMCP must not add a referenced current record to the selected validation records solely because a selected record refers to it. | A record outside the selected scope may still be used for relation lookup. |
| R016 | DRMCP must not admit legacy archive sources as current records for validation. | DRMCP read MVP does not provide legacy lookup. |
| R017 | DRMCP must identify each unadmitted source and each identity-conflict member by a path relative to the repository root. | Absolute physical paths are not required. |
| R018 | DRMCP must not support detailed validation selection by repository-relative or absolute source path. | Sources must first be corrected until they can be admitted and selected by canonical ref. |

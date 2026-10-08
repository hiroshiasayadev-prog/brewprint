# USDM requirement: Current source configuration

- **id**: `usdm:drmcp.mcp_capabilities.discovery_and_listing.current_source_configuration`
- **status**: draft
- **date**: 2026-07-13
- **kind**: requirement
- **parent**: `usdm:drmcp.mcp_capabilities.discovery_and_listing`

## What this is

Product requirements for explicitly configured current sources, app-namespace association, records-root discovery boundaries, and configuration uniqueness.

## Requirements: Current source selection
> source: literal

| id | requirement | notes |
|---|---|---|
| R001 | DRMCP must obtain the current sources used for current-record discovery from explicit configuration. | The concrete configuration representation is defined by downstream Specifications. |
| R002 | Each configured current source must explicitly associate one app namespace with one records root. |  |
| R003 | The configured records root must establish the physical discovery boundary corresponding to the app namespace's Design Records `records/` directory. | Discovery begins beneath the records root rather than the repository root or its parent app directory. |
| R004 | DRMCP must not auto-discover a current source or infer its app namespace from repository directories, directory names, source content, filenames, or declared record identities. |  |

## Requirements: Namespace and physical-root boundary
> source: literal

| id | requirement | notes |
|---|---|---|
| R005 | DRMCP must not derive the configured app namespace from the records root or its parent-directory name. | The physical directory name need not equal the app namespace. |
| R006 | DRMCP must determine current record identities using the configured app namespace and the applicable artifact identity rules relative to the configured records root. | Parent-directory names and the physical `records` path segment do not become canonical identity segments. |
| R007 | DRMCP must not admit a source as a current record when the app namespace determined under the applicable artifact identity rules disagrees with the app namespace associated with its configured current source. | Exact validation and admission-failure representation are defined by downstream Specifications. |
| R008 | DRMCP must exclude sources outside every configured records root from the current-record discovery corpus. | Unconfigured repository areas such as `tools/` remain outside the corpus. |

## Requirements: Configuration uniqueness
> source: literal

| id | requirement | notes |
|---|---|---|
| R009 | One app namespace must be associated with at most one configured current records root. | Multiple current records roots for one app namespace are outside the MVP. |
| R010 | One physical records root must be associated with at most one configured app namespace. | The same source tree must not produce identities under multiple app namespaces. |
| R011 | DRMCP must not process one physical source through more than one configured current source. | Canonicalized root overlap handling is defined by downstream Specifications. |

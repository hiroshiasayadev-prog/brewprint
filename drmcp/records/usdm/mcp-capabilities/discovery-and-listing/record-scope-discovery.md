# USDM requirement: Record scope discovery

- **id**: `usdm:drmcp.mcp_capabilities.discovery_and_listing.record_scope_discovery`
- **status**: draft
- **date**: 2026-07-11
- **kind**: requirement
- **parent**: `usdm:drmcp.mcp_capabilities.discovery_and_listing`

## What this is

Product requirements for discovering selectable current-record scopes before listing or navigation.

## Requirements: Record scope discovery
> source: literal

| id | requirement | notes |
|---|---|---|
| R001 | DRMCP must make each app namespace associated with a configured current source discoverable. | An app remains discoverable even when its configured source contains no discovered record source. |
| R002 | DRMCP must make the artifact kinds available within a selected app namespace discoverable. | Availability may depend on the configured current-source form. |
| R003 | DRMCP must expose whether each available artifact kind uses a sequential or tree record structure. | The concrete representation is defined by downstream Specifications. |
| R004 | DRMCP must make the domain namespaces of a selected app namespace and sequential artifact kind discoverable. |  |
| R005 | DRMCP must include a domain namespace when at least one record source has been discovered in that domain, even when no record in the domain is addressable. | Discovery of the scope must not depend on successful record addressability. |
| R006 | DRMCP must make the root canonical ref of a selected app namespace and tree artifact kind discoverable. |  |
| R007 | When a requested app namespace is unavailable, DRMCP must report the invalid selector and expose the available app namespaces. | Exact diagnostic representation belongs to downstream Specifications. |
| R008 | When an artifact kind is unavailable within a selected app namespace, DRMCP must report the invalid selector and expose the artifact kinds available within that app namespace. | Exact diagnostic representation belongs to downstream Specifications. |
| R009 | DRMCP must not expose physical source paths in normal record-scope discovery results. |  |
| R010 | DRMCP must discover current sources for each artifact kind only beneath the artifact directory declared by that artifact kind's applicable identity-and-structure Specification under the app namespace's `records` directory. | The declared artifact directory establishes the source-discovery boundary for that artifact kind. |
| R011 | DRMCP must not classify or admit a source outside an artifact kind's declared artifact directory as that artifact kind solely from its content, filename, or declared identity. | Sources outside all declared artifact directories are outside the current-record discovery corpus. |

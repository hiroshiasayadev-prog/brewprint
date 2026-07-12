# USDM requirement: Tree record navigation

- **id**: `usdm:drmcp.mcp_capabilities.discovery_and_listing.tree_record_navigation`
- **status**: draft
- **date**: 2026-07-11
- **kind**: requirement
- **parent**: `usdm:drmcp.mcp_capabilities.discovery_and_listing`

## What this is

Product requirements for navigating current records whose artifact kind uses a tree record structure.

## Requirements: Tree record navigation
> source: literal

| id | requirement | notes |
|---|---|---|
| R001 | DRMCP must expose the root canonical ref for a selected app namespace and tree artifact kind. |  |
| R002 | DRMCP must list the direct child records of a selected tree record. | Recursive traversal may be performed through repeated navigation. |
| R003 | DRMCP must return the canonical ref of each listed child record. | Additional compact projection fields may be defined by downstream Specifications. |
| R004 | DRMCP must expose whether each listed child can itself be navigated or is a terminal record. | The internal node representation is not fixed here. |
| R005 | When a requested canonical ref does not identify an available current tree record, DRMCP must return a not-found outcome. | Exact outcome representation belongs to downstream Specifications. |
| R006 | DRMCP must not perform child navigation from a terminal tree record. | Exact error or normal-outcome classification belongs to downstream Specifications. |
| R007 | DRMCP must return tree-navigation results in a deterministic order for the same current-record state. | The ordering key is defined by downstream Specifications. |
| R008 | DRMCP must not expose physical source paths in normal tree record navigation results. |  |

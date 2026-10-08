# USDM requirement: Tree node navigation

- **id**: `usdm:drmcp.mcp_capabilities.discovery_and_listing.tree_record_navigation`
- **status**: draft
- **date**: 2026-07-13
- **kind**: requirement
- **parent**: `usdm:drmcp.mcp_capabilities.discovery_and_listing`

## What this is

Product requirements for navigating current tree nodes whose artifact kind uses a tree record structure.

## Requirements: Tree node navigation
> source: literal

| id | requirement | notes |
|---|---|---|
| R001 | DRMCP must expose the root canonical ref for a selected app namespace and tree artifact kind. |  |
| R002 | DRMCP must list the direct child nodes of a selected current tree node. | Recursive traversal may be performed through repeated navigation. |
| R003 | DRMCP must expose each listed child node's canonical identity through either its complete canonical ref or a deterministic representation derived from the selected node ref. | Additional compact projection fields may be defined by downstream Specifications. |
| R004 | DRMCP must expose whether each listed child node can itself be navigated or is terminal. | The internal node representation is not fixed here. |
| R005 | When a requested canonical ref does not identify an available current tree node, DRMCP must return a not-found outcome. | Exact outcome representation belongs to downstream Specifications. |
| R006 | DRMCP must not perform child navigation from a terminal current tree node. | Exact error or normal-outcome classification belongs to downstream Specifications. |
| R007 | DRMCP must return tree-navigation results in a deterministic order for the same current-record state. | The ordering key is defined by downstream Specifications. |
| R008 | DRMCP must not expose physical source paths in normal tree record navigation results. |  |
| R009 | DRMCP must treat every discovered directory at or beneath a tree artifact source root, including the artifact source root and an empty directory, as an available current tree node independently of a corresponding `index.md` record source. | A missing index record does not remove the node or prevent child navigation. |
| R010 | DRMCP must allow callers to inspect a bounded multi-level projection of a selected current tree subtree without navigating each level separately. | Concrete depth, output-limit, ordering, and projection rules are defined by downstream Specifications. |

# Overview: Discovery and listing operations

- **id**: `spec:drmcp.design_records_mcp.operations.discovery_and_listing`
- **status**: draft
- **date**: 2026-07-14
- **parent**: `spec:drmcp.design_records_mcp.operations`

## What this is

Defines the DRMCP operations that discover selectable current-record scopes and inspect sequential or tree-structured current state.

## Current contract

| operation | purpose |
|---|---|
| `discover_record_scopes` | Discover configured apps, available artifact kinds, sequential domains, and tree roots. |
| `list_sequential_records` | List uniquely addressable current records within one selected sequential domain. |
| `list_tree_children` | List the direct children of one selected current tree node with compact record details. |
| `inspect_tree` | Inspect a bounded multi-level structural projection of one selected current tree subtree. |

Each operation name is also the MCP tool name.

The normal caller flow branches by record structure and required projection:

```text
discover_record_scopes
  -> sequential artifact: list_sequential_records
  -> tree artifact detail: list_tree_children
  -> tree artifact overview: inspect_tree
```

A caller chooses the operation that matches the selected artifact's record structure and the required projection.

## Non-goals

- Current-source configuration serialization.
- Retrieval, search, validation, or legacy lookup.
- Shared diagnostic object, code, context, or message rules.
- Parser, cache, storage, snapshot, concurrency, or filesystem architecture.

## Boundary

| concern | owner |
|---|---|
| Scope availability, tree-node existence, and record addressability | `spec:drmcp.design_records_mcp.current_record_model`. |
| Artifact structure, identity segments, and sequence declarations | `spec:drmcp.design_records_mcp.artifacts`. |
| Request fields, result projections, ordering, limits, and operation outcomes | Child Specifications in this area. |
| Shared failure and advisory representation | `spec:drmcp.design_records_mcp.diagnostics`. |
| Physical source paths for validation or repair | Validation Specifications. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Record scope discovery | Contract | `spec:drmcp.design_records_mcp.operations.discovery_and_listing.record_scope_discovery` | Defines hierarchical discovery of selectable current-record scopes. |
| Sequential record listing | Contract | `spec:drmcp.design_records_mcp.operations.discovery_and_listing.sequential_record_listing` | Defines bounded listing within one selected sequential domain. |
| Tree child listing | Contract | `spec:drmcp.design_records_mcp.operations.discovery_and_listing.tree_child_listing` | Defines detailed direct-child navigation for one selected current tree node. |
| Tree overview | Contract | `spec:drmcp.design_records_mcp.operations.discovery_and_listing.tree_overview` | Defines bounded multi-level structural inspection of one selected current tree subtree. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.current_record_model` | Supplies current scopes, nodes, and addressable records consumed by these operations. |
| `spec:drmcp.design_records_mcp.artifacts` | Supplies artifact structure and identity declarations. |
| `spec:drmcp.design_records_mcp.diagnostics` | Supplies shared operation diagnostic objects and canonical codes. |

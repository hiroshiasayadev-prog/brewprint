# Contract: Tree overview

- **id**: `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_overview`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations.discovery_and_listing`
- **contract_class**: `interface`
## What this is

Defines the `inspect_tree` operation for bounded multi-level structural inspection of one current tree subtree.

## Non-goals

- Returning record titles, status values, or record availability.
- Replacing detailed direct-child listing.
- Returning physical source paths or validation findings.
- Distinguishing physical directory nodes from physical file nodes.

## Request

The request accepts these fields:

| field | requirement | value | default |
|---|---|---|---|
| `node_ref` | required | Exact canonical ref for a tree-structured artifact node. | - |
| `depth` | optional | Integer from `0` through `5`. | `2` |
| `node_limit` | optional | Integer from `1` through `500`. | `100` |

`node_ref` identifies the record kind, app namespace, and selected tree position.
The request does not repeat app or artifact-kind selectors.
The operation does not infer, normalize, or complete `node_ref`.

`depth` counts tree edges from the selected node.

| depth | included boundary |
|---:|---|
| `0` | Selected node only. |
| `1` | Selected node and direct children. |
| `n` | Selected node and descendants at most `n` edges away. |

### Request validation order

The operation evaluates request-wide outcomes in this order:

1. `node_ref` syntax and tree-kind validity.
2. `depth` bounds.
3. `node_limit` bounds.
4. Selected-repository current-state construction.
5. Current tree-node availability.
6. Operation execution.

The first established request-wide error is returned.

## Response

A successful response uses this logical shape:

```text
node_ref
tree:
  segment
  has_children
  children:
    - segment
      has_children
      children
truncated
```

`node_ref` is the complete canonical ref of the selected node.
The selected tree node and every nested node use a canonical identity `segment`.
The selected node's segment is the final identity segment in `node_ref`.
For a tree root, the selected segment is the app namespace.
Each descendant canonical ref is derived by joining the ancestor ref and returned segments with `.`.

### Node projection

| field | presence | value |
|---|---|---|
| `segment` | always | Canonical identity segment for the represented node. |
| `has_children` | always | `true` when the current tree node has at least one direct child. |
| `children` | always | Returned child nodes within the requested depth and node limit. |

`has_children` describes current tree state rather than response expansion.
A node may have `has_children: true` and an empty returned `children` list when the requested depth or node limit prevents expansion.
The overview does not include `has_record`, `title`, or `status`.

### Ordering, traversal, and limit

Sibling nodes use ascending simple string order by canonical `segment`.
DRCLI selects nodes for the response through breadth-first traversal:

```text
selected node
  -> all depth-1 nodes
  -> all depth-2 nodes
  -> later depths
```

Within one depth, parent order follows the prior breadth-first level and each sibling set follows canonical-segment order.
The response may render selected nodes as a nested tree after breadth-first selection.

The selected node counts toward `node_limit`.
`truncated` is `true` only when the node limit excludes a node that otherwise falls within the requested depth.
A requested depth boundary does not itself set `truncated` to `true`.
Reaching the node limit is normal success rather than failure.

### Terminal node

A selected leaf node or empty directory node produces normal success:

```text
tree:
  segment
  has_children: false
  children: []
truncated: false
```

Normal success results do not expose physical source paths.

## Errors

| condition | classification | operation-specific context |
|---|---|---|
| `node_ref` is malformed or uses a non-tree record kind | `invalid_selector` | `node_ref`, `accepted_tree_record_kinds`. |
| A valid tree `node_ref` identifies no available current tree node | `node_not_found` | `node_ref`. |
| `depth` is outside `0..5` | `invalid_depth` | `minimum`, `maximum`, `default`, `actual`. |
| `node_limit` is outside `1..500` | `invalid_node_limit` | `minimum`, `maximum`, `default`, `actual`. |
| Selected-repository current state cannot be constructed reliably | `configuration_failure` | None. |
| The operation cannot produce a reliable complete result | `execution_failure` | None. |

The operation returns no partial result for a configuration or execution failure.
Shared diagnostic objects, canonical codes, context fields, and messages belong to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| Current tree-node existence, descendant relations, and terminal status | `spec:drcli.design_records_cli.current_record_model`. |
| Tree canonical identity mapping | `spec:drcli.design_records_cli.artifacts`. |
| Depth, node limit, breadth-first selection, structural projection, and truncation | This Specification. |
| Detailed child record projection | `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_child_listing`. |
| Detailed source defects and physical source paths | Validation Specifications. |
| Shared error and advisory representation | `spec:drcli.design_records_cli.diagnostics`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.current_record_model` | Supplies the selected current tree subtree. |
| `spec:drcli.design_records_cli.artifacts` | Defines tree canonical refs and path-segment mapping. |
| `spec:drcli.design_records_cli.operations.discovery_and_listing.record_scope_discovery` | Discovers tree root refs. |
| `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_child_listing` | Provides detailed direct-child inspection. |
| `spec:drcli.design_records_cli.diagnostics` | Defines shared error placement and object shape. |
| `spec:drcli.design_records_cli.diagnostics` | Defines canonical codes and context fields used by this operation. |
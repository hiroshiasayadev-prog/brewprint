# Contract: Tree child listing

- **id**: `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_child_listing`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations.discovery_and_listing`
- **contract_class**: `interface`
## What this is

Defines the `list_tree_children` operation for detailed listing of one current tree node's direct children.

## Non-goals

- Recursive or multi-level subtree inspection.
- Listing sequential records.
- Returning physical source paths or validation findings.
- Distinguishing physical directory nodes from physical file nodes.

## Request

The request accepts these fields:

| field | requirement | value | default |
|---|---|---|---|
| `node_ref` | required | Exact canonical ref for a tree-structured artifact node. | - |
| `limit` | optional | Integer from `1` through `500`. | `100` |
| `offset` | optional | Non-negative integer. | `0` |
| `projection` | optional | List containing `title`, `status`, or both. | Empty list. |

`node_ref` identifies the record kind, app namespace, and selected tree position.
The request does not repeat app or artifact-kind selectors.
The operation does not infer, normalize, or complete `node_ref`.

`projection` is an unordered selection set.
Duplicate projection values have the same effect as one occurrence.
Any unknown projection value makes the request invalid.

### Request validation order

The operation evaluates request-wide outcomes in this order:

1. `node_ref` syntax and tree-kind validity.
2. Projection validity.
3. `limit` bounds.
4. `offset` bounds.
5. Selected-repository current-state construction.
6. Current tree-node availability.
7. Operation execution.

The first established request-wide error is returned.

## Response

A successful child-listing response uses this logical shape:

```text
node_ref
children:
  - segment
    has_children
    has_record
    title?
    status?
applied_offset
applied_limit
has_more
```

### Child identity

`node_ref` is the complete canonical ref of the selected node.
Each child exposes one canonical identity `segment` relative to `node_ref`.
The complete child canonical ref is derived as:

```text
<node_ref> + "." + <segment>
```

The operation does not repeat the complete shared canonical-ref prefix for every child.

### Child items

| field | presence | value |
|---|---|---|
| `segment` | always | Child canonical identity segment. |
| `has_children` | always | `true` when at least one direct child current tree node exists. |
| `has_record` | always | `true` when one uniquely addressable current record corresponds to the child node. |
| `title` | only when requested | Parsed title string, or `null` when unavailable. |
| `status` | only when requested | Parsed status string, or `null` when unavailable. |

A node remains listable when `has_record` is `false`.
A requested `title` or `status` is `null` when no corresponding record exists or the field is unavailable.
Detailed source defects belong to validation.

### Ordering and range

Sibling nodes use ascending simple string order by canonical `segment`.
DRCLI applies `offset` and then `limit` after ordering.

`applied_offset` and `applied_limit` report the effective request values.
`has_more` is `true` only when another direct child exists after the returned range.
An offset at or beyond the direct-child count returns a successful empty `children` list with `has_more: false`.
Reaching the output limit is normal success rather than failure.

### Terminal node

A selected leaf node or empty directory node produces a distinct normal `terminal node` outcome.
The outcome includes the selected `node_ref` and does not include a child-listing result.
A terminal-node outcome is not an operation failure.

Normal responses do not expose physical source paths.

## Errors

| condition | classification | operation-specific context |
|---|---|---|
| `node_ref` is malformed or uses a non-tree record kind | `invalid_selector` | `node_ref`, `accepted_tree_record_kinds`. |
| A valid tree `node_ref` identifies no available current tree node | `node_not_found` | `node_ref`. |
| `projection` contains an unknown value | `invalid_projection` | `invalid_projection_values`, `available_projection_values`. |
| `limit` is outside `1..500` | `invalid_limit` | `minimum`, `maximum`, `default`, `actual`. |
| `offset` is negative | `invalid_offset` | `minimum`, `default`, `actual`. |
| Selected-repository current state cannot be constructed reliably | `configuration_failure` | None. |
| The operation cannot produce a reliable complete result | `execution_failure` | None. |

The operation returns no partial result for a configuration or execution failure.
Shared diagnostic objects, canonical codes, context fields, and messages belong to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| Current tree-node existence, child relations, and terminal status | `spec:drcli.design_records_cli.current_record_model`. |
| Corresponding record admission and addressability | `spec:drcli.design_records_cli.current_record_model`. |
| Tree canonical identity mapping | `spec:drcli.design_records_cli.artifacts`. |
| Compact child identity, detailed projection, sibling order, range, and outcomes | This Specification. |
| Multi-level structural inspection | `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_overview`. |
| Detailed source defects and physical source paths | Validation Specifications. |
| Shared error and advisory representation | `spec:drcli.design_records_cli.diagnostics`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.current_record_model` | Supplies current tree nodes and direct-child relations. |
| `spec:drcli.design_records_cli.current_record_model` | Defines corresponding uniquely addressable records. |
| `spec:drcli.design_records_cli.artifacts` | Defines tree canonical refs and path-segment mapping. |
| `spec:drcli.design_records_cli.operations.discovery_and_listing.record_scope_discovery` | Discovers tree root refs. |
| `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_overview` | Provides bounded multi-level structural inspection. |
| `spec:drcli.design_records_cli.diagnostics` | Defines shared error placement and object shape. |
| `spec:drcli.design_records_cli.diagnostics` | Defines canonical codes and context fields used by this operation. |
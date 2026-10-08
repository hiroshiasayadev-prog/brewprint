# Contract: Sequential record listing

- **id**: `spec:drmcp.design_records_mcp.operations.discovery_and_listing.sequential_record_listing`
- **status**: draft
- **date**: 2026-07-14
- **parent**: `spec:drmcp.design_records_mcp.operations.discovery_and_listing`
- **contract_class**: `interface`
- **usdm_covers**:
  - usdm:drmcp.mcp_capabilities.discovery_and_listing.sequential_record_listing#R001-R014

## What this is

Defines the `list_sequential_records` operation for bounded listing within one selected sequential domain scope.

## Non-goals

- Discovering available apps, artifact kinds, or domains.
- Listing tree nodes.
- Returning unadmitted sources, identity-conflict members, or physical source paths.
- Defining artifact-specific identity segments or sequence formats.

## Request

The request accepts these fields:

| field | requirement | value | default |
|---|---|---|---|
| `app_namespace` | required | Exact configured app namespace. | - |
| `artifact_kind` | required | Exact available sequential artifact kind. | - |
| `domain_namespace` | required | Exact discovered domain namespace. | - |
| `order` | optional | `ascending` or `descending`. | `descending` |
| `limit` | optional | Integer from `1` through `100`. | `30` |
| `offset` | optional | Non-negative integer. | `0` |
| `projection` | optional | List containing `title`, `status`, or both. | Empty list. |

All three scope selectors are operation-required.
The input schema permits omission so the operation can return `invalid_selector` with `missing_selector_fields`.
An omitted or invalid selector does not broaden the selected scope.
The operation does not infer, normalize, or complete selector values.

`projection` is an unordered selection set.
Duplicate projection values have the same effect as one occurrence.
Any unknown projection value makes the request invalid.

### Request validation order

The operation evaluates request-wide outcomes in this order:

1. Required selector presence.
2. Projection validity.
3. `limit` bounds.
4. `offset` bounds.
5. Configured current-state construction.
6. App-namespace availability.
7. Sequential artifact-kind availability.
8. Domain-namespace availability.
9. Operation execution.

The first established request-wide error is returned.

## Response

A successful response uses this logical shape:

```text
records:
  - canonical_ref
    title?
    status?
applied_offset
applied_limit
has_more
```

### Record items

| field | presence | value |
|---|---|---|
| `canonical_ref` | always | Canonical ref of one uniquely addressable current record. |
| `title` | only when requested | Parsed title string, or `null` when unavailable. |
| `status` | only when requested | Parsed status string, or `null` when unavailable. |

An unavailable requested projection field does not remove the record from the result.
Projection failure for one field does not change repository-wide addressability.
Detailed source defects belong to validation.

### Ordering

DRMCP constructs the ordering tuple from the selected artifact's identity-and-structure Specification.
The tuple contains every artifact-specific segment declared with `sequence: yes`.
Tuple elements follow the row order of the `Artifact-specific segments` table.
Each element uses the comparison defined by its declared segment format.

DRMCP compares tuples from left to right.
`order` applies to the complete tuple.
Records with identical tuples use ascending simple string order by canonical ref as a fixed tie-break.
The canonical-ref tie-break does not reverse for descending order.

### Range application

DRMCP applies listing stages in this order:

```text
selected domain scope
  -> uniquely addressable records
  -> ordering
  -> offset
  -> limit
```

`applied_offset` and `applied_limit` report the effective request values.
`has_more` is `true` only when another matching record exists after the returned range.
An offset at or beyond the matching record count returns a successful empty `records` list with `has_more: false`.
A valid domain with no eligible record also returns a successful empty `records` list.

Normal success results do not expose physical source paths.
Reaching the output limit is normal success rather than failure.

## Errors

| condition | classification | operation-specific context |
|---|---|---|
| A required scope selector is omitted | `invalid_selector` | `missing_selector_fields`. |
| Requested app namespace is unavailable | `invalid_selector` | `app_namespace`, `available_app_namespaces`. |
| Requested artifact kind is unknown, non-sequential, or unavailable within the selected app | `invalid_selector` | `artifact_kind`, `available_sequential_artifact_kinds`. |
| Requested domain is unavailable within the selected app and artifact kind | `invalid_selector` | `domain_namespace`, `available_domain_namespaces`. |
| `projection` contains an unknown value | `invalid_projection` | `invalid_projection_values`, `available_projection_values`. |
| `limit` is outside `1..100` | `invalid_limit` | `minimum`, `maximum`, `default`, `actual`. |
| `offset` is negative | `invalid_offset` | `minimum`, `default`, `actual`. |
| Configured current state cannot be constructed reliably | `configuration_failure` | None. |
| The operation cannot produce a reliable complete result | `execution_failure` | None. |

The operation returns no partial result for a configuration or execution failure.
Shared diagnostic objects, canonical codes, context fields, and messages belong to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| Domain availability and included addressable records | `spec:drmcp.design_records_mcp.current_record_model.current_record_scopes`. |
| Unique addressability and no-winner conflict handling | `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability`. |
| Sequence-bearing segments, row priority, and segment formats | Applicable artifact identity-and-structure Specifications. |
| Request fields, projection, ordering direction, tie-break, range, and result indicators | This Specification. |
| Detailed record defects and physical source paths | Validation Specifications. |
| Shared error and advisory representation | `spec:drmcp.design_records_mcp.diagnostics`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.current_record_model.current_record_scopes` | Supplies the selected sequential domain and its addressable-record set. |
| `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability` | Defines uniquely addressable current records. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` | Defines sequence declarations and ordering-tuple construction. |
| `spec:drmcp.design_records_mcp.operations.discovery_and_listing.record_scope_discovery` | Discovers valid sequential selectors. |
| `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_envelope` | Defines shared error placement and object shape. |
| `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_catalog` | Defines canonical codes and context fields used by this operation. |

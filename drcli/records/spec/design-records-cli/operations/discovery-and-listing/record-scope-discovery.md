# Contract: Record scope discovery

- **id**: `spec:drcli.design_records_cli.operations.discovery_and_listing.record_scope_discovery`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations.discovery_and_listing`
- **contract_class**: `interface`
## What this is

Defines the `discover_record_scopes` operation for hierarchical discovery of selectable current-record scopes.

## Non-goals

- Listing records within a sequential domain.
- Listing tree children or inspecting a subtree.
- Reconstructing scope availability rules owned by the current record model.
- Exposing physical source paths or validation findings.

## Request

The request accepts these fields:

| field | requirement | value |
|---|---|---|
| `app_namespace` | optional | Exact discovered app namespace. |
| `artifact_kind` | optional | Exact artifact-kind value declared by an artifact Specification. |

The accepted selector combinations are:

| `app_namespace` | `artifact_kind` | result variant |
|---|---|---|
| omitted | omitted | `app_scopes` |
| present | omitted | `artifact_kind_scopes` |
| present | present, sequential kind | `domain_scopes` |
| present | present, tree kind | `tree_root` |

An `artifact_kind` without an `app_namespace` is an invalid selector combination.
The operation does not infer, normalize, or complete selector values.

### Request validation order

The operation evaluates request-wide outcomes in this order:

1. Selector-combination validity.
2. Selected-repository current-state construction.
3. App-namespace availability.
4. Artifact-kind availability within the selected app.
5. Operation execution.

The first established request-wide error is returned.

## Response

A successful response contains exactly one result variant.

### App scopes

```text
app_scopes:
  - app_namespace
```

The operation returns every discovered app scope.
Items use ascending simple string order by `app_namespace`.
A selected repository with no discovered app scope returns `app_scopes: []`.

### Artifact-kind scopes

```text
artifact_kind_scopes:
  - artifact_kind
    record_structure
```

`record_structure` is `sequential` or `tree` as declared by the artifact Specification.
Items use ascending simple string order by `artifact_kind`.
A valid app with no available artifact-kind scope returns `artifact_kind_scopes: []`.

### Domain scopes

```text
domain_scopes:
  - domain_namespace
```

The operation returns discovered sequential domain scopes for the selected app and artifact kind.
Items use ascending simple string order by `domain_namespace`.
A valid sequential artifact-kind scope with no discovered domain returns `domain_scopes: []`.

### Tree root

```text
tree_root:
  canonical_ref
```

The operation returns the root canonical ref for the selected tree artifact-kind scope.
A valid tree artifact-kind scope always has one root current tree node.

Normal success results do not expose physical source paths.

## Errors

| condition | classification | operation-specific context |
|---|---|---|
| `artifact_kind` is supplied without `app_namespace` | `invalid_selector` | `accepted_selector_combinations`. |
| Requested app namespace is unavailable | `invalid_selector` | `app_namespace`, `available_app_namespaces`. |
| Requested artifact kind is unavailable within the selected app | `invalid_selector` | `artifact_kind`, `available_artifact_kinds`. |
| Selected-repository current state cannot be constructed reliably | `configuration_failure` | None. |
| The operation cannot produce a reliable complete result | `execution_failure` | None. |

The operation returns no partial result for a configuration or execution failure.
Shared diagnostic objects, canonical codes, context fields, and messages belong to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| App, artifact-kind, domain, and tree-root availability | `spec:drcli.design_records_cli.current_record_model`. |
| Artifact directories and record structures | `spec:drcli.design_records_cli.artifacts`. |
| Selector progression and result variants | This Specification. |
| Invalid-selector trigger and applicable context variant | This Specification. |
| Shared error and advisory representation | `spec:drcli.design_records_cli.diagnostics`. |
| Unadmitted-source paths and validation guidance | Validation Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.current_record_model` | Defines the scopes discovered by this operation. |
| `spec:drcli.design_records_cli.current_record_model` | Defines selected-repository current sources and artifact source roots. |
| `spec:drcli.design_records_cli.artifacts` | Defines sequential and tree record structures. |
| `spec:drcli.design_records_cli.operations.discovery_and_listing.sequential_record_listing` | Consumes a discovered sequential domain. |
| `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_child_listing` | Consumes a discovered tree root or child node. |
| `spec:drcli.design_records_cli.operations.discovery_and_listing.tree_overview` | Consumes a discovered tree root or child node. |
| `spec:drcli.design_records_cli.diagnostics` | Defines shared error placement and object shape. |
| `spec:drcli.design_records_cli.diagnostics` | Defines canonical codes and context fields used by this operation. |
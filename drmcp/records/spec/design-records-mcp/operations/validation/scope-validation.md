# Contract: Scope validation

- **id**: `spec:drmcp.design_records_mcp.operations.validation.scope_validation`
- **status**: draft
- **date**: 2026-07-14
- **parent**: `spec:drmcp.design_records_mcp.operations.validation`
- **contract_class**: `interface`
- **usdm_covers**:
  - usdm:drmcp.mcp_capabilities.validation.validation_scope_and_subjects#R001-R004,R006-R010,R015,R023
  - usdm:drmcp.mcp_capabilities.validation.validation_result_projection#R003-R010,R016,R022-R024
  - usdm:drmcp.mcp_capabilities.validation.reference_and_structure_validation#R010

## What this is

Defines the `validate_scope` operation for broad validation of one selected current-state boundary.

## Non-goals

- Returning detailed findings for abnormal current records.
- Selecting one current record by exact canonical ref.
- Defining artifact-specific validation rules or finding schemas.
- Selecting a source by repository-relative or absolute path.
- Returning conforming record refs or tree nodes without advisories.
- Repairing source content or reference values.

## Request

The request accepts these optional fields:

| field | value | default |
|---|---|---|
| `app_namespace` | Exact configured app namespace. | Omitted. |
| `artifact_kind` | Exact artifact-kind value declared by an artifact Specification. | Omitted. |
| `domain_namespace` | Exact discovered sequential domain namespace. | Omitted. |
| `subtree_root_ref` | Exact canonical ref of one available current tree node. | Omitted. |
| `result_limit` | Maximum returned complete abnormal result items. | `200`. |

The accepted scope-selector combinations are:

| supplied selector fields | selected scope |
|---|---|
| none | Repository current-state boundary. |
| `app_namespace` | One app scope. |
| `app_namespace`, `artifact_kind` | One artifact-kind scope. |
| `app_namespace`, `artifact_kind`, `domain_namespace` | One sequential-domain scope. |
| `subtree_root_ref` | One current tree subtree. |

The prefix selector fields must be supplied without gaps.
`subtree_root_ref` is exclusive with every prefix selector field.
A domain selector applies only to a sequential artifact kind.
DRMCP does not trim, repair, normalize, complete, or infer a supplied selector.
An omitted trailing prefix selector selects the corresponding broader scope.

Validation subjects are exactly the sources, records, and tree nodes included by the selected scope.
A referenced current record outside the selected scope is not added solely because a selected record refers to it.

Input-schema violations are rejected before operation result construction.
Schema violations include unknown fields, wrong value types, and a non-integer `result_limit`.

### Tree subtree scope

A valid `subtree_root_ref` may identify a directory node or a leaf node.
The selected scope contains the root node and every descendant current tree node.
A corresponding current record is not required for the selected root node.
A leaf root produces a valid one-node subtree.

The subtree includes discovered sources, unadmitted sources, candidates, identity-conflict members, addressable records, and tree nodes placed within the selected physical tree-node boundary.
A record outside the subtree is not added to the validation subjects solely because a selected record refers to it.

### Request validation order

The operation evaluates request-wide outcomes in this order:

1. Scope-selector combination validity.
2. `result_limit` bounds.
3. Configured current-state construction.
4. Semantic validity and availability of the selected scope.
5. Validation execution.

Prefix selectors use this semantic order:

1. App availability.
2. Artifact-kind availability.
3. Sequential-kind eligibility when `domain_namespace` is supplied.
4. Domain availability.

A subtree selector uses this semantic order:

1. Canonical tree-ref form and tree-kind eligibility.
2. Current tree-node availability.

The first established request-wide error is returned.
Scope-selector combination and numeric-limit validation precede configured current-state construction.

## Response

A normal response uses this logical shape:

```text
ok
summary:
  validated_records
  abnormal_records
  abnormal_tree_nodes
  contract_violations
  advisories
  unadmitted_sources
  identity_conflicts
results:
  - <abnormal-result>
applied_result_limit
total_results
results_truncated
```

The operation always returns complete scope aggregates and abnormal result summaries.
The operation does not provide an aggregates-only projection mode.
The response does not echo the selected scope.
The normal response does not contain top-level `warnings`.

### Overall outcome

`ok` is `false` when the complete selected scope contains at least one of these conditions:

- a contract-violation finding on a validated current record;
- an unadmitted source;
- an identity conflict.

`ok` is `true` when none of those conditions exists.
Advisory-only record or tree-node results do not make `ok` false.
A request-wide error is not represented as `ok: false`.

### Scope summary

| field | value |
|---|---|
| `validated_records` | Number of uniquely addressable current records validated in the complete selected scope. |
| `abnormal_records` | Number of validated records with at least one contract-violation or advisory finding. |
| `abnormal_tree_nodes` | Number of non-record tree nodes with at least one advisory finding. |
| `contract_violations` | Total contract-violation findings across all validated records. |
| `advisories` | Total advisory findings across validated records and non-record tree nodes. |
| `unadmitted_sources` | Number of unadmitted source results. |
| `identity_conflicts` | Number of conflicting canonical identities. |

Every summary value describes the complete selected scope.
Output limiting does not change the summary or `ok`.

### Abnormal result collection

`results` is one tagged union with these variants:

```text
unadmitted_source
identity_conflict
record
tree_node
```

Conforming records are omitted.
Tree nodes without advisories are omitted.
The collection does not contain detailed record findings.
A caller uses `validate_records` with a returned `record_ref` to obtain detailed findings.

#### Unadmitted source

```text
type: unadmitted_source
path
reason
```

`path` is relative to the repository root.
`reason` is exactly one of:

```text
source_unreadable
source_unparseable
identity_unavailable
identity_inconsistent
```

Each unadmitted source produces one result item.

#### Identity conflict

```text
type: identity_conflict
claimed_ref
member_paths:
  - <repository-root-relative-path>
```

`claimed_ref` is the canonical identity claimed by multiple current artifact candidates.
`member_paths` contains every conflict member path in ascending simple string order.
The collection contains at least two member paths.
DRMCP does not select or identify a winner.

An identity conflict is included when at least one member source belongs to the selected scope.
`member_paths` still contains every repository-wide conflict member, including members outside the selected scope.
An outside member is diagnostic context and does not become a selected validation subject.

#### Abnormal record

```text
type: record
record_ref
contract_violations
advisories
```

`record_ref` is the current canonical ref of the uniquely addressable record.
Both counts describe the record's complete validation outcome.
At least one count is greater than zero.
The counts make the abnormal state identifiable without a separate state field.

#### Abnormal tree node

```text
type: tree_node
node_ref
advisories
```

`node_ref` is the canonical ref of the current tree node.
`advisories` is greater than zero.
This variant represents advisories not attached to a current record.
Every tree-node advisory established by the current scope model produces this variant.
The current advisory set includes a directory node whose corresponding `index.md` source is absent.

### Deterministic ordering

DRMCP orders complete abnormal results by variant in this order:

```text
unadmitted_source
identity_conflict
record
tree_node
```

DRMCP then orders items within each variant as follows:

| variant | ordering key |
|---|---|
| `unadmitted_source` | `path` ascending simple string order. |
| `identity_conflict` | `claimed_ref` ascending simple string order. |
| `record` | `record_ref` ascending simple string order. |
| `tree_node` | `node_ref` ascending simple string order. |

Filesystem order, discovery order, and validation execution order do not affect result order.
`validation_output_limits` applies `result_limit` to this complete ordered collection.

### Empty result

A valid scope with no abnormal condition returns:

```text
ok: true
summary: <all-zero-summary>
results: []
applied_result_limit
total_results: 0
results_truncated: false
```

A valid scope may contain no addressable current record.
The empty state remains normal success when no unadmitted source, identity conflict, record finding, or tree-node advisory exists.

## Errors

| condition | classification | operation-specific context |
|---|---|---|
| Supplied scope-selector fields do not form one accepted combination | `invalid_selector` | `accepted_selector_combinations`. |
| Requested app namespace is unavailable | `invalid_selector` | `app_namespace`, `available_app_namespaces`. |
| Requested artifact kind is unavailable within the selected app | `invalid_selector` | `artifact_kind`, `available_artifact_kinds`. |
| `domain_namespace` is supplied for an available non-sequential artifact kind | `invalid_selector` | `artifact_kind`, `available_sequential_artifact_kinds`. |
| Requested sequential domain is unavailable | `invalid_selector` | `domain_namespace`, `available_domain_namespaces`. |
| `subtree_root_ref` is malformed or does not identify a tree artifact kind | `invalid_selector` | `subtree_root_ref`, `accepted_tree_record_kinds`. |
| A valid tree `subtree_root_ref` identifies no available current tree node | `node_not_found` | `subtree_root_ref`. |
| `result_limit` is outside `1..2000` | `invalid_limit` | `minimum`, `maximum`, `default`, `actual`. |
| Configured current state cannot be constructed reliably | `configuration_failure` | None. |
| The operation cannot produce a reliable complete validation result | `execution_failure` | None. |

A request-wide error returns exactly one top-level `error`.
The response contains no normal result fields or top-level warnings.
Shared diagnostic objects, canonical codes, context fields, and messages belong to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| Repository, app, artifact-kind, domain, and subtree scope contents | `spec:drmcp.design_records_mcp.current_record_model.current_record_scopes`. |
| Source admission and unadmitted-source boundaries | `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission`. |
| Repository-wide identity conflicts and member provenance | `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability`. |
| Broad request, result variants, aggregates, ordering, and request-wide outcomes | This Specification. |
| Broad-result limit bounds and application | `spec:drmcp.design_records_mcp.operations.validation.validation_output_limits`. |
| Detailed findings for one current record | `spec:drmcp.design_records_mcp.operations.validation.exact_record_validation`. |
| Finding rules, finding codes, and affected-subject locators | Other Validation Specifications. |
| Shared errors and warnings | `spec:drmcp.design_records_mcp.diagnostics`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.current_record_model.current_record_scopes` | Defines selectable current-state boundaries and tree subtree membership. |
| `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission` | Defines current records and unadmitted sources. |
| `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability` | Defines identity conflicts and repository-wide member sets. |
| `spec:drmcp.design_records_mcp.operations.validation.exact_record_validation` | Provides detailed findings for returned abnormal record refs. |
| `spec:drmcp.design_records_mcp.operations.validation.validation_output_limits` | Defines complete-item result limiting. |
| `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_envelope` | Defines top-level error placement and response exclusivity. |
| `spec:drmcp.design_records_mcp.diagnostics.operation_diagnostic_catalog` | Defines canonical scope-validation error codes and contexts. |

# Contract: Scope validation

- **id**: `spec:drcli.design_records_cli.operations.validation.scope_validation`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.operations.validation`
- **contract_class**: `interface`

## What this is

Defines the logical `validate_scope` operation for broad validation of one selected current-state boundary.

## Non-goals

- Returning detailed findings for abnormal current records.
- Selecting one current record by exact canonical ref.
- Defining artifact-specific validation rules or finding schemas.
- Selecting a source by repository-relative or absolute path.
- Returning conforming record refs or tree nodes without advisories.
- Repairing source content, selectors, references, or metadata.
- Defining CLI parsing, help, usage, rendering, streams, or exit status.

## Request

The operation consumes a type-valid logical request with these optional fields:

| field | value | default |
|---|---|---|
| `app_namespace` | Exact discovered app namespace. | Omitted. |
| `artifact_kind` | Exact artifact-kind value declared by an applicable artifact Specification. | Omitted. |
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

Prefix selector fields must be supplied without gaps.
`subtree_root_ref` is exclusive with every prefix selector field.
A domain selector applies only to a sequential artifact kind.
DRCLI does not trim, repair, normalize, complete, or infer a supplied semantic selector.
An omitted trailing prefix selector selects the corresponding broader scope.

Validation subjects are exactly the sources, records, and tree nodes included by the selected scope.
A referenced current record outside the selected scope is not added solely because a selected record refers to it.
CLI syntax and conversion into this logical request belong to `spec:drcli.design_records_cli.cli`.
CLI usage failures are not validation findings or semantic operation diagnostics defined by this contract.

### Tree subtree scope

A valid `subtree_root_ref` may identify a directory node or a leaf node.
The selected scope contains the root node and every descendant current tree node.
A corresponding current record is not required for the selected root node.
A leaf root produces a valid one-node subtree.

The subtree includes discovered sources, unadmitted sources, candidates, identity-conflict members, addressable records, and tree nodes placed within the selected physical tree-node boundary.
A record outside the subtree is not added solely because a selected record refers to it.

### Request validation order

The operation evaluates request-wide semantic outcomes in this order:

1. Scope-selector combination validity.
2. `result_limit` bounds.
3. Selected-repository current-state construction.
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

The first established request-wide semantic error is returned.
Selector-combination and numeric-limit validation precede selected-repository current-state construction.

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

The operation always computes complete scope aggregates and abnormal result summaries before output limiting.
The operation does not provide an aggregates-only projection mode.
The response does not echo the selected scope.
The normal response does not contain top-level `warnings`.
The logical response does not prescribe CLI serialization or rendering.

### Overall outcome

`ok` is `false` when the complete selected scope contains at least one of these conditions:

- a contract-violation finding on a validated current record;
- an unadmitted source;
- an identity conflict.

`ok` is `true` when none of those conditions exists.
Advisory-only record or tree-node results do not make `ok` false.
A request-wide operation error is not represented as `ok: false`.

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
The current-record model owns source admission and reason semantics.

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
DRCLI does not select or identify a winner.

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

DRCLI orders the complete abnormal result collection by variant in this order:

```text
unadmitted_source
identity_conflict
record
tree_node
```

DRCLI then orders items within each variant as follows:

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
| Selected-repository current state cannot be constructed reliably | `configuration_failure` | None. |
| The operation cannot produce a reliable complete validation result | `execution_failure` | None. |

A request-wide semantic error returns exactly one top-level `error`.
The response contains no normal result fields or top-level warnings.
Shared diagnostic objects, canonical codes, and context fields belong to Diagnostics Specifications.

## Boundary

| concern | owner |
|---|---|
| Repository, app, artifact-kind, domain, and subtree scope contents | `spec:drcli.design_records_cli.current_record_model.current_record_scopes`. |
| Source admission and unadmitted-source boundaries | `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission`. |
| Repository-wide identity conflicts and member provenance | `spec:drcli.design_records_cli.current_record_model.identity_conflict_and_addressability`. |
| Broad request, result variants, aggregates, deterministic ordering, and request-wide outcomes | This Specification. |
| Broad-result limit bounds and application | `spec:drcli.design_records_cli.operations.validation.validation_output_limits`. |
| Detailed findings for one current record | `spec:drcli.design_records_cli.operations.validation.exact_record_validation`. |
| Finding rules, finding codes, affected-subject locators, and conformance semantics | The canonical Specification that owns each validated rule. |
| Shared semantic operation errors and warnings | `spec:drcli.design_records_cli.diagnostics`. |
| CLI usage and rendering behavior | `spec:drcli.design_records_cli.cli`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.current_record_model.current_record_scopes` | Defines selectable current-state boundaries and tree subtree membership. |
| `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission` | Defines current records and unadmitted sources. |
| `spec:drcli.design_records_cli.current_record_model.identity_conflict_and_addressability` | Defines identity conflicts and repository-wide member sets. |
| `spec:drcli.design_records_cli.operations.validation.exact_record_validation` | Provides detailed findings for returned abnormal record refs. |
| `spec:drcli.design_records_cli.operations.validation.validation_output_limits` | Defines complete-item result limiting. |
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_envelope` | Defines top-level error placement and response exclusivity. |
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog` | Defines canonical scope-validation error codes and contexts. |
| `spec:drcli.design_records_cli.cli` | Owns CLI usage diagnostics and concrete output behavior. |

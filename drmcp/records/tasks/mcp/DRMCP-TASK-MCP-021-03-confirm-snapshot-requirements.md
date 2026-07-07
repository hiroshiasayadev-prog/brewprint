# DRMCP-TASK-MCP-021-03: Confirm snapshot requirements

- **id**: DRMCP-TASK-MCP-021-03
- **status**: done
- **date**: 2026-07-07
- **work_item**: DRMCP-WORK-MCP-021
- **task_type**: decision
- **estimate**: 1.0d
- **depends_on**:
  - DRMCP-TASK-MCP-021-02
- **outputs**:
  - DRMCP-TASK-MCP-021-03

## Goal

Run a decision loop over the T02 investigation tables and confirm W021 snapshot and lookup-map requirements.

## Work

### Decision boundary

This Task consumes T02 investigation evidence.

It decides:

- the accepted public-use-case consumption matrix for Current Records Snapshot and Legacy exact lookup map;
- the accepted Current Records Snapshot minimum requirements;
- the accepted Legacy exact lookup-map minimum requirements;
- the accepted Application-to-Domain construction-output requirements;
- which items are already covered by existing Specifications;
- which concrete gaps route to downstream Domain structural work;
- the downstream closure route into W022.

It does not decide:

- public operation request or response contracts;
- Domain parser internals;
- Domain logical-tree object layout;
- Domain relation-graph object layout;
- resolver or validator internal algorithms;
- concrete Infrastructure adapter fields;
- Go signatures, structs, interfaces, functions, packages, algorithms, fixtures, or tests.

### Decision inventory

| ID | topic | status | depends on | decision summary | reason / evidence | canonical target | ADR route |
|---|---|---|---|---|---|---|---|
| D-001 | Public use-case consumption matrix | `decided` | T02 | Accepted T02 public-use-case consumption matrix as the W021 consumer baseline. | User accepted the clarified matrix. Guidance diagnostic and source-location uncertainty remains for later gap classification. | W021 downstream specification route | not_required |
| D-002 | Current Records Snapshot minimum requirements | `decided` | D-001 | Accept the minimum requirement set for Current identity lookup, compact fields, metadata, headings, validation subjects, current relation inputs, duplicate conflict groups, source provenance, and body-retention exclusion. | User accepted the set as requirements. Application-visible surface shape remains undecided and is reserved for D-004. | W021 downstream specification route | not_required |
| D-003 | Legacy exact lookup-map minimum requirements | `decided` | D-001 | Accept the minimum requirement set for separate request state, request scope, conditional construction, exact lookup, lookup outcomes, source provenance, retrieval, resolver input, validation input, operation-error boundary, and exclusions. | User accepted the set after clarifying that required map construction failure means an Application operation error, not a partial normal response. Direct reader shape remains D-004 material. | W021 downstream specification route | not_required |
| D-004 | Application-to-Domain construction-output requirements | `decided` | D-002, D-003 | Accept D-004a through D-004d as the complete Application-to-Domain requirement baseline for T03. | User concluded that further D-004 subdivision repeats the same question. Remaining precision belongs in Domain detailed-contract work or D-005 gap classification, not more T03 requirement decisions. | W022 Domain structural route | not_required |
| D-005 | Existing-spec coverage and concrete gaps | `decided` | D-002, D-003, D-004 | Existing schema specs cover current identity, field vocabulary, source retention, invalid-source behavior, and duplicate-conflict behavior. W021 closes with the requirement baseline. Domain detailed contracts still have semantic output gaps. | User accepted minimal coverage and gap classification. The classification routes precision work without duplicating existing schema specifications. | W022 Domain structural route | not_required |
| D-006 | Downstream synchronization, authoring, review, and closure route | `decided` | D-005 | Route Domain-facing gaps to W022. Close W021 after materializing the W022 decision Task that references this Task. | User accepted that W022 should carry the next decisions and W021 need not continue into Domain detail or W021 Specification authoring. T03 does not create implementation tasks, Go contracts, ADRs, fixtures, tests, or production plans. | W021 closure and W022 task graph | not_required |

### Accepted Current Records Snapshot minimum requirements

| requirement | accepted minimum |
|---|---|
| Current identity lookup | The snapshot must support current canonical identity lookup and active-index addressability. Duplicate-conflict identities have no addressable winner. |
| Compact fields | The snapshot must support `list_records` projection of `ref`, `title`, `status`, and `date`. |
| Metadata projection | The snapshot must support `get_records` projection of normalized current metadata. |
| Heading projection | The snapshot must support `get_records` and Guidance projection of headings or H1 title. |
| Validation subjects | The snapshot must expose validation subjects. Subjects include retained sources, parse-failed sources, identity-less validation-only sources, addressable records, and duplicate-conflict groups. |
| Current relation inputs | The snapshot must provide Current relation lookup inputs. The snapshot must provide Current relation validation inputs. |
| Duplicate conflict groups | The snapshot must preserve duplicate conflict groups for validation and diagnostics. Listing must keep conflicted identities winnerless. |
| Source provenance | The snapshot must preserve source provenance or source references. Diagnostics and exact source retrieval use those references. |
| Body retention exclusion | The snapshot is not required to retain every record body in memory. Operations may read body content through source references or source contracts. |

Application-visible surface shape is not decided by D-002. Direct fields, query capability, or opaque Domain outcomes remain D-004 material.

### Accepted Legacy exact lookup-map minimum requirements

| requirement | accepted minimum |
|---|---|
| Separate request state | Legacy exact lookup map is separate from Current Records Snapshot. |
| Request scope | The map is fresh, immutable, and discarded after the request. |
| Conditional construction | Application builds the map only when the operation requires Legacy compatibility lookup. |
| Guidance exclusion | Guidance operations do not build or read the map. |
| Exact lookup | The map supports exact lookup by accepted legacy issued ID or accepted legacy key. |
| Lookup outcomes | The map distinguishes unique, absent, duplicate, and unreadable outcomes. |
| Source provenance | The map preserves Legacy source provenance or source references. |
| Legacy body retrieval | The map can support body retrieval through Legacy source access when `get_records` requires body content. |
| Resolver input | The map can supply Legacy fallback input after Current resolution has no target. |
| Validation input | The map can supply Legacy relation-target lookup input for validation. |
| Operation-error boundary | Application returns an operation error when a required map cannot be built trustworthily. |
| No partial normal map | Application does not return a partial normal map when required map state is incomplete or untrustworthy. |
| Exclusion | The map excludes Current Records Snapshot, active Current index, MCP policy, concrete filesystem access, and mutable cache. |

D-003 does not decide who directly reads the map. Direct Application access, Reference Resolution access, or Relation Graph Validation access remains D-004 material.

### D-004 accepted subdecisions

| ID | topic | accepted outcome |
|---|---|---|
| D-004a | Domain and Application responsibility split | Domain provides semantic views over parsed records. The views include identity, record summary data, document structure, relation inputs, validation inputs, parse findings, and conflict state. Application decides operation scope, use-case selection, aggregation, public projection, and operation-error selection. |
| D-004b | Markdown scan and body retention boundary | Current Records Snapshot construction may scan Markdown source to derive fields, headings, parse findings, provenance, and conflict state. The snapshot is not required to retain every verbatim body. Public body projection and section extraction remain operation-specific. |
| D-004c | Record summary view boundary | Domain provides a record summary view for Application selection. The view includes identity, common fields, parsed kind-specific fields, heading summary, source reference, and validity state. Domain does not provide public response-specific compact projections. Application owns list, retrieval, Guidance, and response projection choices. |
| D-004d | Validation material boundary | Domain provides validation materials for Application use-case judgment. The materials include validation subject candidates, parse findings, local check results, and relation check results. Application decides validation scope, selected subjects, pass order, aggregation, public diagnostics, and operation-error selection. D-004d does not decide validation handler layout, rule IDs, finding schema, or Go types. |

D-004 is complete with D-004a through D-004d. Additional identity, addressability, tree, graph, lookup, parser, and validation details are outside T03 requirement decisions. D-005 classifies whether those details are covered by existing specifications or require Domain detailed-contract follow-up.

### Existing-spec coverage and concrete gaps

| classification | items |
|---|---|
| Covered by existing schema specifications | Current identity rules, field vocabulary, source retention, invalid-source behavior, duplicate-conflict behavior, and source provenance baseline. |
| W021 requirement baseline | Current Records Snapshot minimum requirements, Legacy exact lookup-map minimum requirements, operation-error boundary, and request-scope coordination requirements. |
| Domain detailed-contract gaps | Semantic views, record summary view, validation materials, resolver inputs, validation inputs, relation inputs, and exact Domain output boundaries. |
| Outside T03 scope | Go types, package layout, parser mode names, algorithms, fixtures, tests, public response shapes, and production implementation planning. |

### Downstream synchronization route

1. Route Domain-facing gaps to W022 Domain structural work.
2. Materialize a W022 decision Task that references this Task's decisions.
3. Close W021 after W022 can consume the requirement baseline without guessing.
4. Keep public response shape work outside T03.
5. Keep ADR creation, implementation Tasks, Go contracts, fixtures, tests, and production plans outside T03.

## Done condition

- Every owned decision item is `decided`, `deferred`, or validly `blocked`.
- Accepted Current Records Snapshot minimum requirements are recorded.
- Accepted Legacy exact lookup-map minimum requirements are recorded.
- Accepted Application-to-Domain construction-output requirements are recorded without deciding Domain internals.
- Existing-spec coverage and concrete gaps are classified.
- The downstream synchronization and closure route is known.
- No Specification, ADR, implementation Task, production implementation, Go contract-shape probe, fixture, test, or independent review is authored by this Task.

## Verification

- Scoped diff inspected for this Task file.
- Scoped whitespace check passed by inspection of the generated patch.
- No stage, commit, push, Specification authoring, ADR authoring, implementation Task authoring, Go contract-shape probe, fixture, test, or production implementation planning was performed.

## Evidence

- DRMCP-TASK-MCP-021-02 is the required predecessor investigation-type Task.
- User clarified that the investigation tables should be reviewed through a later decision loop.
- User clarified that W022 should carry the next Domain decisions by referencing this Task.
- This Task is created before any W021 detailed Specification authoring Task.

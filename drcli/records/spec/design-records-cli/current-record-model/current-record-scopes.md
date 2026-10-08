# Concept: Current record scopes

- **id**: `spec:drcli.design_records_cli.current_record_model.current_record_scopes`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.current_record_model`
- **usdm_covers**:
  - usdm:drmcp.mcp_capabilities.discovery_and_listing.record_scope_discovery#R001-R006,R009
  - usdm:drmcp.mcp_capabilities.discovery_and_listing.tree_record_navigation#R001,R009
  - usdm:drmcp.mcp_capabilities.validation.reference_and_structure_validation#R010
  - usdm:drmcp.mcp_capabilities.validation.validation_scope_and_subjects#R009

## What this is

Defines current-state boundaries for resolved app scopes, artifact kinds, sequential domains, tree nodes, subtrees, and broad validation.

## Concept model

| concept | definition |
|---|---|
| Repository current-state boundary | Every resolved current source and all source, identity, record, and tree-node states derived from them. |
| App scope | The current-state boundary associated with one resolved app namespace. |
| Artifact-kind scope | The part of one app scope beneath one existing artifact source root. |
| Sequential domain scope | The part of one sequential artifact-kind scope beneath one discovered physical domain directory. |
| Current tree node | A path-derived node represented by a tree directory or non-index source path. |
| Tree root node | The directory node represented by the tree artifact source root and its root canonical ref. |
| Subtree scope | One available current tree node and every descendant current tree node. |
| Scope addressable-record set | The uniquely addressable current records included by one scope. |

A current-record scope is a current-state boundary.
A scope is not only a collection of admitted records.

## Rules

### Scope availability

| scope | availability condition |
|---|---|
| Repository | Selected-repository current-source construction completes as a trustworthy complete state. |
| App | One resolved current source associates the app namespace with one resolved records root. |
| Artifact kind | The artifact kind's derived artifact source root exists beneath the app's resolved records root. |
| Sequential domain | At least one current source is discovered directly within the physical domain directory. |
| Tree root | The tree artifact source root exists. |
| Subtree | The supplied canonical ref identifies an available current tree node. |

An app scope remains available when its resolved records root yields no addressable current record.
Each artifact-kind scope retains the `sequential` or `tree` structure declared by its artifact Specification.
Scope discovery exposes that declared structure.
An artifact-kind scope remains available when its artifact source root is empty.
A sequential domain scope remains available when its discovered sources produce no addressable record.
Scope availability does not depend on successful record addressability unless the scope explicitly selects records.

### Sequential domains

The directory directly beneath a sequential artifact source root determines the domain scope name.
A source inside that directory makes the domain discoverable before candidate formation.

An H1 public ID with another domain makes the source unadmitted.
The source does not move to the H1-declared domain scope.
The physically discovered domain remains available.

### Tree nodes

The tree artifact source root establishes the root directory node.
Its canonical ref is valid independently of a corresponding root `index.md` source.

Tree nodes are derived as follows:

| physical element | tree-node effect |
|---|---|
| Artifact source root | Establishes the root directory node. |
| Directory beneath the artifact source root | Establishes one directory node, including an empty directory. |
| `index.md` inside a directory | Supplies a possible current record for the existing directory node. |
| Non-index source file | Establishes one leaf node and supplies a possible current record for that leaf. |

An `index.md` source does not create its directory node.
Removing or invalidating the corresponding record does not remove the node.

An unreadable or otherwise unadmitted non-index tree source still establishes its path-derived leaf node.
The node remains navigable as a terminal node, but no current record exists for exact retrieval.

### Tree navigation and subtree scopes

Tree navigation operates on available current tree nodes rather than admitted current records.
A directory node exposes its direct directory and leaf children.
A leaf node is terminal.
An empty directory node is terminal.

A subtree scope may begin at a directory or leaf node.
A corresponding current record is not required.
The subtree includes the selected node and every descendant node.

A missing directory `index.md` source does not invalidate its subtree scope.
Broad validation reports the missing source as an advisory.
The advisory applies to the artifact source root, non-empty directories, and empty directories.

### Scope contents by consumer

| consumer | scope contents consumed |
|---|---|
| Sequential listing | Uniquely addressable sequential records in the selected domain scope. |
| Exact retrieval | One uniquely addressable current record selected by canonical ref. |
| Lexical search | Uniquely addressable current records included by the selected scopes. |
| Tree navigation | Available current tree nodes, including nodes without corresponding records. |
| Broad validation | Discovered sources, unadmitted sources, candidates, identity conflicts, addressable records, and tree nodes. |
| Exact-ref validation | Uniquely addressable current records selected by exact canonical ref. |

A nonconforming admitted record remains in the scope addressable-record set.
An unadmitted source remains in broad validation scope but not in the addressable-record set.
An identity-conflict member remains in broad validation scope but produces no addressable record.

Discovery, listing, search, and validation consume this shared scope model.
A consuming operation does not redefine scope availability or inclusion.

### Valid empty states

These conditions produce valid scopes with an empty addressable-record set:

- a repository with no discovered candidate records root;
- a resolved app scope with no addressable current record;
- an empty artifact source root;
- a discovered sequential domain containing only unadmitted or conflicting sources;
- a tree node with no corresponding record;
- a subtree containing no addressable records.

A valid empty addressable-record set does not make the scope unavailable.

### Physical path exposure

Normal scope discovery and tree navigation expose canonical refs and logical scope information.
They do not expose physical source paths.
Repository-root-relative paths remain available only where validation or repair context requires them.

## Boundary

| concern | owner |
|---|---|
| Scope availability and inclusion | This Specification. |
| Sequential domain discovery | This Specification. |
| Tree-node and subtree existence | This Specification. |
| Tree node versus current record distinction | This Specification. |
| Missing tree-directory `index.md` advisory condition | This Specification. |
| Resolved current source and artifact source-root derivation | `spec:drcli.design_records_cli.current_record_model.current_source_corpus`. |
| Candidate, unadmitted-source, and admission states | `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission`. |
| Repository-wide unique addressability | `spec:drcli.design_records_cli.current_record_model.identity_conflict_and_addressability`. |
| Tool selectors, result projection, ordering, limits, and diagnostics | Consuming operation Specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.current_record_model.current_source_corpus` | Establishes app and artifact source boundaries. |
| `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission` | Defines source and record states included by scopes. |
| `spec:drcli.design_records_cli.current_record_model.identity_conflict_and_addressability` | Supplies the repository-wide addressable-record set. |
| `spec:drcli.design_records_cli.artifacts.base.definitions.record_structure` | Defines sequential and tree structure semantics. |

# Validation atomic component inventory

## Status

- Working inventory created on 2026-07-15.
- This document is not a canonical Design Record.
- Each category is added only after explicit user confirmation.
- Confirmed entries define capability boundaries for later artifact/base recipes and artifact-to-component mapping.
- Provisional engine syntax and final evaluation shapes remain outside this inventory.

## Inventory criteria

An atomic capability remains in the inventory only when it:

- has one independently meaningful responsibility;
- has at least one concrete consumer;
- has reusable value or owns clear artifact/base meaning;
- is not merely a trivial implementation helper;
- may remain even when its native implementation is trivial when its name exposes an independently meaningful design operation and hides a lower-level comparison or manipulation;
- does not embed artifact validity or finding classification;
- returns output reusable by later components or recipes.

## Confirmed inventory

### `common.markdown`

`common.markdown` owns generic Markdown structure selection and explicit source-location data exposed by declared occurrence types only.
It does not own Design Record H1 grammar, metadata notation, artifact validity, permitted headings, or finding meaning.

#### `common.markdown.heading-occurrences`

| field | content |
|---|---|
| capability ID | `common.markdown.heading-occurrences` |
| purpose | Enumerate real headings from a parsed Markdown document and optionally filter them by level or exact heading text. |
| input | `document: MarkdownDocument`, `level?: integer`, `text?: string` |
| output | `list<HeadingOccurrence>` containing level, text, heading form, raw marker, and source location. |
| owner | `common.markdown` |
| first consumers | H1 count, H1 first-line placement, ATX H1 requirement, H2 occurrence counts, required/optional/prohibited H2 rules, and unlisted-H2 discovery across artifacts. |
| custom needed | no |
| keep / merge / reject recommendation | **keep**. Separate H1 and H2 selectors are merged into this capability. |

The capability returns heading structure only.
Interpretation of `# <H1_PREFIX>: <Title>` remains under `artifact/base` or the artifact namespace.

#### `common.markdown.section-occurrences`

| field | content |
|---|---|
| capability ID | `common.markdown.section-occurrences` |
| purpose | Enumerate sections rooted at headings of a supplied level and expose each section body range. |
| input | `document: MarkdownDocument`, `level: integer`, `heading?: string` |
| output | `list<SectionOccurrence>` containing the heading occurrence, body fragment, and source range. |
| owner | `common.markdown` |
| first consumers | Substantive H2 body validation, Reference Specification table-presence validation, Topics section parsing, and accepted Requirement Required Outcome body selection. |
| custom needed | no |
| keep / merge / reject recommendation | **keep**. H2 occurrence selection and section-body retrieval are one responsibility. |

A separate `section-body` capability is unnecessary because the selected section occurrence already contains its body fragment.

#### `common.markdown.yaml-front-matter-occurrence`

| field | content |
|---|---|
| capability ID | `common.markdown.yaml-front-matter-occurrence` |
| purpose | Select a YAML front matter block at the beginning of a Markdown source when one exists. |
| input | `source: MarkdownSource` |
| output | `optional<FrontMatterOccurrence>` containing raw body, delimiters, and source range. |
| owner | `common.markdown` |
| first consumers | Current-source YAML front matter prohibition shared by all current artifacts. |
| custom needed | no |
| keep / merge / reject recommendation | **keep**. It exposes the occurrence and its declared source range without deciding whether front matter is allowed. |

The prohibition itself belongs to `artifact/base/validation`.

#### `common.markdown.table-occurrences`

| field | content |
|---|---|
| capability ID | `common.markdown.table-occurrences` |
| purpose | Enumerate Markdown tables within a supplied Markdown document or fragment. |
| input | `content: MarkdownFragment` |
| output | `list<TableOccurrence>` containing header cells, body rows, cells, and source locations. |
| owner | `common.markdown` |
| first consumers | Reference Specification requirement for at least one table in an H2 body and Specification Topics table parsing. |
| custom needed | no |
| keep / merge / reject recommendation | **keep**. Table existence and counts are evaluated later through collection and generic validation checks. |

#### `common.markdown.table-column-occurrences`

| field | content |
|---|---|
| capability ID | `common.markdown.table-column-occurrences` |
| purpose | Select every column occurrence whose header exactly matches a supplied column name and preserve row-aligned cells. |
| input | `table: TableOccurrence`, `column_name: string` |
| output | `list<TableColumnOccurrence>` containing each matching header location and its row-aligned cells. |
| owner | `common.markdown` |
| first consumers | Topics table `title`, `kind`, `ref`, and `summary` column extraction. |
| custom needed | no |
| keep / merge / reject recommendation | **keep**. Returning all matching columns preserves missing-column and duplicate-column evidence for later checks. |

### Rejected or merged `common.markdown` candidates

| candidate | disposition | reason |
|---|---|---|
| `common.markdown.h1-data` | merge and split ownership | H1 occurrence selection belongs to `heading-occurrences`; H1 prefix/title interpretation belongs to `artifact/base`. |
| `common.markdown.h2-occurrences` | merge | Expressed by `section-occurrences` with `level: 2`. |
| `common.markdown.section-body` | merge | Body is part of `SectionOccurrence`; a separate projection is a trivial helper. |
| `common.markdown.first-line` | reject | Source line information is already carried by occurrences and can be checked generically. |
| `common.markdown.table-exists` | reject | Table selection plus collection count and a generic check is sufficient. |
| `common.markdown.metadata-field-occurrences` | reject from `common` | Visible field notation and the H1-to-first-H2 metadata region are Design Record semantics owned by `artifact/base`. |
| `common.markdown.h1-prefix-title` | reject from `common` | The H1 delimiter shape is an artifact/base source contract rather than generic Markdown structure. |
| `common.markdown.list-items` | reject as a public capability | No independently meaningful current consumer requires a public generic list-item selector. |

### `common.text`

`common.text` owns scalar text operations that are meaningful in design-facing YAML.
A capability may remain even when its Go implementation is a one-line operation when the function name hides a lower-level comparison and makes the intended processing step explicit.

#### `common.text.trim`

| field | content |
|---|---|
| capability ID | `common.text.trim` |
| purpose | Remove leading and trailing whitespace before later semantic interpretation or validation. |
| input | `value: string` |
| output | `string` |
| owner | `common.text` |
| first consumers | H1 titles, metadata scalar values, metadata child items, H2 bodies, Topics `title` and `summary`, Task `estimate`, and Investigation `trigger`, `scope`, and `non_scope`. |
| custom needed | no |
| implementation form | Native Go is the initial implementation candidate. |
| keep / merge / reject recommendation | **keep**. The YAML explicitly records that whitespace normalization occurs before later meaning is evaluated. |

#### `common.text.is-empty`

| field | content |
|---|---|
| capability ID | `common.text.is-empty` |
| purpose | Determine whether a string is empty without exposing a low-level equality comparison in design-facing YAML. |
| input | `value: string` |
| output | `boolean` |
| owner | `common.text` |
| first consumers | Emptiness checks after trimming H1 titles, metadata values, list items, and section bodies. |
| custom needed | no |
| implementation form | Native Go is the initial implementation candidate. |
| keep / merge / reject recommendation | **keep**. It is a semantic function rather than a generic `value == ""` expression. |

#### `common.text.equals`

| field | content |
|---|---|
| capability ID | `common.text.equals` |
| purpose | Determine whether two text values are exactly equal. |
| input | `left: string`, `right: string` |
| output | `boolean` |
| owner | `common.text` |
| first consumers | Steps-defined artifact/base placeholder recognition against the exact `TBD` literal. |
| custom needed | no |
| implementation form | Native Go is the initial implementation candidate. |
| keep / merge / reject recommendation | **keep**. A concrete steps-defined semantic consumer now requires exact text equality while keeping comparison operators out of the authoring surface. |

### Deferred or rejected `common.text` candidates

| candidate | disposition | reason |
|---|---|---|
| `common.text.is-blank` | defer | Initially express as `trim` followed by `is-empty`; promote to a graph-defined reusable function only when mapping demonstrates repeated design value. |
| `common.text.is-non-empty` | defer | Retain only if artifact mapping shows that direct non-empty vocabulary is clearer than boolean composition. |
| `common.text.matches` | reject | Supplying arbitrary patterns would expose implementation syntax in design YAML. A named grammar or format function should own each accepted meaning. |
| `common.text.starts-with` | reject | Current consumers are artifact identity or file-projection rules rather than generic text meaning. |
| `common.text.split` | reject | H1, identity, and metadata parsing are implementation details of the owning semantic function. |
| `common.text.length` | reject | No accepted validation rule currently constrains character count. |
| `common.text.normalize-whitespace` | reject | The current contracts define trimming but no broader whitespace normalization. |
| `common.text.is-single-line` | reject | Current heading structure already supplies this boundary and there is no independent consumer. |

The current inventory records scalar functions only.
Applying them to collection elements is evidence for later investigation of minimal iteration support rather than a reason to add `trim-all` or `all-non-empty` convenience functions.

### `common.collection`

`common.collection` owns artifact-independent collection analysis over its declared input values and output records.
It does not attach or propagate hidden source or evidence metadata.
It does not decide whether duplicates, missing values, or unexpected values violate an artifact contract.

#### `common.collection.is-empty`

| field | content |
|---|---|
| capability ID | `common.collection.is-empty` |
| purpose | Determine whether a supplied collection contains no elements without exposing `count(values) == 0` in design-facing YAML. |
| input | `values: list<T>` |
| output | `boolean` |
| owner | `common.collection` |
| first consumers | Duplicate-discovery results, set-comparison differences, metadata child collections, and heading or table occurrence collections. |
| custom needed | no |
| implementation form | Native Go is the initial implementation candidate. |
| keep / merge / reject recommendation | **keep**. It expresses collection emptiness as a named design operation. |

#### `common.collection.find-duplicates`

| field | content |
|---|---|
| capability ID | `common.collection.find-duplicates` |
| purpose | Discover values that occur more than once and preserve every matching input occurrence through declared output fields. |
| input | `values: list<T comparable>` |
| output | `list<DuplicateGroup<T>>`, where each group contains the duplicated value and all occurrences. |
| owner | `common.collection` |
| first consumers | Specification Topics refs, expanded `usdm_covers`, and persisted reference collections across artifact kinds. |
| custom needed | no |
| implementation form | Native Go is the initial implementation candidate. |
| keep / merge / reject recommendation | **keep**. Returning duplicate groups retains evidence needed by later validation findings. |

A separate `has-duplicates` capability is unnecessary because `find-duplicates` can be followed by `common.collection.is-empty` without discarding duplicate evidence.

#### `common.collection.compare-sets`

| field | content |
|---|---|
| capability ID | `common.collection.compare-sets` |
| purpose | Compare two collections as sets and expose values present only on either side. |
| input | `left: list<T comparable>`, `right: list<T comparable>` |
| output | `SetComparison<T>` containing at least `left_only`, `right_only`, and `equal`. |
| owner | `common.collection` |
| first consumers | Exact completeness comparison between Specification Topics refs and physical direct-child Specification refs. |
| custom needed | no |
| implementation form | Native Go is the initial implementation candidate. |
| keep / merge / reject recommendation | **keep**. One comparison result supports equality and precise missing or unexpected-value findings. |

Duplicate discovery remains a separate operation.
Set comparison must not silently turn duplicated input into a valid collection.

### Deferred, rejected, or merged `common.collection` candidates

| candidate | disposition | reason |
|---|---|---|
| `common.collection.count` | defer | Current contracts mostly require semantic cardinality checks such as exactly once or at most once rather than the raw numeric count. Revisit during generic-check inventory. |
| `common.collection.has-duplicates` | merge | `find-duplicates` plus `is-empty` preserves both the decision value and duplicate evidence. |
| `common.collection.set-equals` | merge | Available as `compare-sets.equal`. |
| `common.collection.set-difference` | merge | Available as `compare-sets.left_only` and `right_only`. |
| `common.collection.ordered-equals` | reject | No accepted current rule requires collection-order equality. |
| `common.collection.subset` | reject | No concrete current consumer. |
| `common.collection.contains` | reject | Current candidate consumers such as direct self-reference have stronger semantic names and should not be represented as low-level value lookup. |
| `common.collection.distinct` | reject | It would silently discard duplicate evidence. |
| `common.collection.first` / `single` | reject | Selecting one occurrence could hide a cardinality violation. |
| `common.collection.map` / `filter` | defer as callable capabilities | Determine whether minimal runtime iteration should own these operations after artifact mapping exposes concrete requirements. |
| `common.collection.any` / `all` | defer | Determine together with boolean aggregation and minimal `if` / `for` behavior. |

Current constraints already provide evidence that some bounded iteration mechanism is likely required for metadata items, reference resolution, duplicate findings, Topics child checks, and graph-edge construction.
The exact authoring syntax remains deferred.

### `common.tree`

`common.tree` owns artifact-independent traversal over an already constructed tree model.
It does not own Design Record tree-node kinds, canonical refs, corresponding-record semantics, or missing `index.md` meaning.

#### `common.tree.direct-children`

| field | content |
|---|---|
| capability ID | `common.tree.direct-children` |
| purpose | Enumerate the direct child nodes of one supplied tree node. |
| input | `tree: Tree<T>`, `node: TreeNode<T>` |
| output | `list<TreeNode<T>>` |
| owner | `common.tree` |
| first consumers | Tree navigation and physical direct-child selection for Specification Topics completeness. |
| custom needed | no |
| implementation form | Native Go is the initial implementation candidate. |
| keep / merge / reject recommendation | **keep**. Direct-child traversal is an independently meaningful tree operation. |

The capability does not project children into canonical Specification refs.
That projection belongs to `artifact/base` or the applicable artifact namespace.

#### `common.tree.subtree-nodes`

| field | content |
|---|---|
| capability ID | `common.tree.subtree-nodes` |
| purpose | Enumerate one supplied root node and every descendant node beneath it. |
| input | `tree: Tree<T>`, `root: TreeNode<T>` |
| output | `list<TreeNode<T>>` |
| owner | `common.tree` |
| first consumers | `validate_scope` subtree selection and validation-subject enumeration for tree artifact scopes. |
| custom needed | no |
| implementation form | Native Go is the initial implementation candidate. |
| keep / merge / reject recommendation | **keep**. It directly represents the accepted subtree scope of root plus all descendants. |

### Deferred, rejected, or relocated `common.tree` candidates

| candidate | disposition | reason |
|---|---|---|
| `common.tree.descendants` | merge | No current consumer requires descendants while excluding the supplied root. |
| `common.tree.all-nodes` | merge | Call `subtree-nodes` with the tree root. |
| `common.tree.directory-nodes` | relocate to `artifact/base` | Directory-node classification is part of the Design Record tree structure model. |
| `common.tree.corresponding-record` | relocate to `artifact/base` | Node-to-current-record correspondence belongs to artifact current-state semantics. |
| `common.tree.has-corresponding-record` | relocate to `artifact/base` | It contributes to the artifact/base tree `index.md` rule. |
| `common.tree.node-by-ref` | reject from `common` | Canonical-ref interpretation and resolution belong to `artifact/base/reference`. |
| `common.tree.parent` | reject for now | No direct current validation consumer requires a standalone generic parent selector. |
| `common.tree.leaves` / `depth` / `ancestors` / `path` | reject | No concrete current consumer. |
| `common.tree.map` / `filter` | defer | Evaluate as part of the minimal runtime iteration requirement. |

Tree artifact mapping adds further evidence that bounded iteration is likely required, including child-record parent checks across every physical direct child.
The exact `for` or collection-iteration syntax remains deferred.

### `common.graph`

`common.graph` owns artifact-independent graph algorithms over already normalized relation edges.
It does not select relation fields, normalize artifact-specific edges, combine unrelated relation types, or decide finding meaning.

#### `common.graph.find-directed-cycles`

| field | content |
|---|---|
| capability ID | `common.graph.find-directed-cycles` |
| purpose | Detect cyclic regions in supplied directed edges and return the participating nodes, edges, and at least one witness cycle. Any evidence data is carried explicitly by the declared edge type. |
| input | `edges: list<DirectedEdge<N>>` |
| output | `list<DirectedCycleGroup<N>>`, where each group contains participating nodes and edges plus at least one witness path. |
| owner | `common.graph` |
| first consumers | Task dependency cycles, Decision dependency cycles, Decision supersession cycles, and a future Work Item provenance cycle rule if accepted. |
| custom needed | no |
| implementation form | Native Go is the initial implementation candidate. |
| keep / merge / reject recommendation | **keep**. It provides reusable directed-cycle analysis while preserving evidence for artifact-owned findings. |

The output does not enumerate every elementary cycle because their number may grow exponentially.
An implementation may group cyclic strongly connected regions while preserving at least one concrete witness cycle for each returned group.
A self-loop is a cycle.

Each artifact relation supplies its own normalized edge collection and calls this function separately.
Decision dependency and supersession edges, for example, must not be combined into one graph merely because both are directed relations.

### Rejected or relocated `common.graph` candidates

| candidate | disposition | reason |
|---|---|---|
| `common.graph.is-acyclic` | merge | Apply `common.collection.is-empty` to the cycle-discovery result while preserving cycle evidence. |
| `common.graph.build` | reject | The supplied edge collection is already the graph input; another wrapper has no independent meaning. |
| `common.graph.add-edge` | reject | This is an implementation-level construction operation. |
| Task dependency edge construction | relocate to `artifact/task` | Field selection and normalization carry Task relation meaning. |
| Decision dependency or supersession edge construction | relocate to `artifact/decision` | Each relation owns separate artifact semantics. |
| Work Item provenance edge construction | relocate to `artifact/work-item` | Provenance normalization is artifact-specific and may require a native semantic function. |
| `common.graph.reachable` | reject for now | No concrete accepted validation consumer. |
| `common.graph.connected-components` | reject for now | No independent current rule consumes generic connectivity. |
| `common.graph.topological-sort` | reject for now | No accepted validation contract requires ordering by graph topology. |
| `common.graph.find-path` | reject for now | Cycle witness generation is part of `find-directed-cycles`; no other current consumer exists. |

## Confirmed parse, candidate, and validation boundary

Source parsing, candidate formation, admission, and conformance validation are separate responsibilities.

```text
source
  -> syntax and source-structure parsing
  -> artifact facts and identity determination
  -> candidate formation and repository-wide admission
  -> conformance validation
```

### Source parsing

`artifact.base.source.parse` is the shared Design Record source parser boundary.
It preserves discoverable source structure, including explicit source-location fields in declared occurrence records, even for partial or malformed states where reliable structure can still be retained.
Its output may contain H1 occurrences, an H1 prefix/title candidate, H1-adjacent metadata occurrences, H2 sections, tables, source locations, and explicit unavailable projections.

Source parsing does not decide:

- artifact field requiredness;
- value vocabularies;
- required or prohibited H2 headings;
- title validity;
- reference resolution;
- finding classification;
- complete artifact conformance.

Artifact-specific facts select and interpret relevant values from this source representation without discarding invalid actual values that validation needs as evidence.

### Identity and candidate formation

Only identity availability and expected-context agreement gate candidate formation.

For sequential artifacts, artifact-specific identity determination must obtain one complete H1 public ID that conforms to the applicable identity grammar and agrees with the configured app namespace, corpus-selected artifact kind, and physical domain.
For tree artifacts, canonical node identity is path-derived and tree-node construction does not depend on full source validation.

Nonidentity defects do not remove candidate or admitted-record status.
Examples include invalid status values, missing nonidentity metadata, prohibited metadata `id`, file-name projection mismatch, H2 violations, title violations, and unresolved persisted references.

### Listing and retrieval

Listing, retrieval, navigation, and search consume admitted uniquely addressable records without requiring full conformance validation.
A requested projection may return an invalid but structurally available value, such as an unrecognized status string.
A projection returns unavailable or `null` only when the source structure cannot produce that projection reliably.
Invalid and unavailable are distinct states.

### Validation

Artifact validation consumes parsed source structure, artifact facts, identity results, physical context, repository state, resolver state, and tree state.
It evaluates the complete artifact contract and emits findings without removing nonconforming admitted records from listing, retrieval, search, or exact validation eligibility.

Each artifact therefore owns separate concerns rather than one "validated parse":

- validation-independent fact projection;
- structure-specific identity determination;
- conformance validation recipes and semantic functions.

## Pending categories

The remaining categories are intentionally undecided:

1. `artifact/base/reference`;
2. `artifact/base` facts;
3. `operation.validation.check`;
4. artifact-specific facts and custom component candidates.

# Artifact constraint inventory

## Status and boundary

This is a working investigation record, not a canonical Specification.

The investigation examined the validation-relevant source contracts for the six current DRMCP artifact kinds:

- Specification;
- Requirement;
- Decision;
- Investigation;
- Work Item;
- Task.

The inventory focused on constraints that can produce validation findings. It did not define operation request schemas, output limits, parser implementation, authoring transactions, or repair behavior.

## Primary source set

### Shared artifact contracts

- `drmcp/records/spec/design-records-mcp/artifacts/base/definitions/source.md`
- `drmcp/records/spec/design-records-mcp/artifacts/base/definitions/h1-adjacent-metadata.md`
- `drmcp/records/spec/design-records-mcp/artifacts/base/definitions/identity-declaration.md`
- `drmcp/records/spec/design-records-mcp/artifacts/base/definitions/record-structure.md`

### Artifact-specific DRMCP contracts

For each artifact kind, the investigation used:

- `source.md`;
- `h1-adjacent-metadata.md`;
- `identity-and-structure.md`.

These are under:

```text
drmcp/records/spec/design-records-mcp/artifacts/<artifact-kind>/
```

### Product and operation authority used to close gaps

- `product/records/spec/design-records/authoring-standards/requirement-authoring.md`
- `product/records/spec/design-records/spec-format/topics-table.md`
- the Product authoring and traceability Specifications referenced by each artifact contract;
- `drmcp/records/spec/design-records-mcp/current-record-model/current-record-scopes.md`;
- `drmcp/records/spec/design-records-mcp/operations/validation/exact-record-validation.md`.

## Investigation dimensions

Each artifact was inspected across the same validation surfaces:

| surface | examples |
|---|---|
| source placement | artifact directory, domain directory, direct file placement, tree source root |
| identity | H1 authority, path-derived identity, kind and app agreement, sequential segments |
| file projection | file-name public-ID prefix and optional slug |
| H1 | marker, prefix, title, count, first-line placement |
| metadata | field allow-list, occurrence count, form, type, vocabulary, conditional presence |
| H2 structure | required, optional, recommended, prohibited, unlisted policy |
| content | empty body, `TBD`, table presence, implementation-contract shape |
| references | canonical form, permitted target kinds, resolution, self-reference |
| collections | empty items, duplicates, order significance, set equality |
| relations | parent agreement, reciprocal ownership, segment agreement |
| graphs | dependency and supersession cycles, relation-specific edge construction |
| tree consistency | direct children, Topics rows, directory `index.md` presence |

## Shared artifact constraints

### Source and H1

- H1 is the first source line.
- The source shape is `# <H1_PREFIX>: <Title>`.
- Sequential artifacts use the complete H1 identity prefix as the sole source-internal canonical identity authority.
- Tree artifacts derive canonical identity from source placement and path mapping.
- H1-adjacent metadata appears after H1 and before the first H2.
- Current artifacts do not use YAML front matter.
- The title must be non-empty after trimming and must remain on one line.

### Metadata

- Metadata uses visible bullet fields: `- **<field_name>**: <value>`.
- Each artifact declares an exact field inventory.
- Unlisted metadata fields are contract violations.
- Sequential metadata does not persist `id`.
- A tree artifact may expose `id` only as a projection checked against path-derived identity.
- Field requirement, form, value type, value vocabulary, and permitted target kinds remain artifact-owned.

### Sequential structure

- Source placement is `<APP_NAMESPACE>/records/<ARTIFACT_DIRECTORY>/<DOMAIN_NAMESPACE>/<record-file>.md`.
- Standard sequential records are files directly beneath the domain directory.
- H1 identity must agree with configured app namespace, selected artifact kind, and physical domain directory.
- The file name projects the H1 public ID but does not establish identity.
- The accepted file-name rule is `<PUBLIC-ID>.md` or `<PUBLIC-ID>-<slug>.md`.
- Slug-to-title equality is not validated.

### Tree structure

- A tree artifact source root represents the root node.
- Directories represent directory nodes.
- `index.md` represents its containing directory node.
- Non-index Markdown files represent leaf nodes.
- Nodes exist and remain navigable independently of whether a corresponding record is admitted.
- Latest accepted direction: every tree directory node requires a corresponding `index.md`; absence is a contract violation but does not remove the node or subtree.

## Artifact overview

| artifact | structure | artifact directory | identity form |
|---|---|---|---|
| Specification | tree | `spec` | `spec:<APP_NAMESPACE>(.<PATH_SEGMENT>)*` |
| Requirement | sequential | `requirements` | `<APP_NAMESPACE>-REQ-<DOMAIN_NAMESPACE>-<SEQUENCE>` |
| Decision | sequential | `adr` | `<APP_NAMESPACE>-ADR-<DOMAIN_NAMESPACE>-<SEQUENCE>` |
| Investigation | sequential | `investigations` | `<APP_NAMESPACE>-INV-<DOMAIN_NAMESPACE>-<SEQUENCE>` |
| Work Item | sequential | `work-items` | `<APP_NAMESPACE>-WORK-<DOMAIN_NAMESPACE>-<SEQUENCE>` |
| Task | sequential | `tasks` | `<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>` |

## Artifact-specific constraint summary

### Specification

#### H1 and metadata

- H1 prefix is one of `Overview`, `Index`, `Concept`, `Reference`, or `Contract`.
- Exactly one real ATX H1 appears outside fenced code blocks.
- YAML front matter is prohibited for current sources.
- Metadata fields are `id`, `status`, `date`, `parent`, conditional `contract_class`, and optional `usdm_covers`.
- `id` must equal the path-derived canonical Specification ref.
- `date` uses strict `YYYY-MM-DD`.
- `parent` is a Specification ref or the literal `root` or `-`.
- `contract_class` is required only for `Contract` and is one of `interface` or `format`.
- Expanded `usdm_covers` requirement IDs must not contain duplicates.

#### H2 and content

- H2 applicability depends on H1 kind and, for Contract, `contract_class`.
- Unlisted H2 headings are allowed when Product authority permits them.
- A missing recommended heading is an advisory.
- A `Reference` Specification must contain at least one Markdown table inside a body H2 section; the table schema and rows are not validated by this rule.
- An H2 body is substantive when its trimmed text is neither empty nor exactly `TBD`.
- No generic sentence-count or block-count requirement applies.

#### Topics and tree consistency

For an `Index` or an `Overview` that declares `## Topics`:

- Topics `ref` values contain no duplicates;
- every Topics `ref` resolves to a Specification;
- Topics refs equal the physical direct-child Specification refs as a set;
- every direct child appears exactly once;
- every listed ref is a direct child;
- each child `parent` metadata value equals the declaring Specification ref.

Every Specification tree directory node requires `index.md` under the latest accepted direction.

### Requirement

#### Identity and placement

- Sequential `REQ` identity with a three-digit domain-scoped sequence.
- The H1 identity agrees with app, kind, and physical domain.
- File name starts with the complete public ID and may append a slug.

#### Metadata

Latest Product authority overrides the stale optional/`Any string` DRMCP field table:

- `status` is required and one of `captured`, `decision_needed`, `accepted`, `deferred`, or `rejected`;
- `date` is required and strict `YYYY-MM-DD`;
- `source_refs` is required, with an empty list allowed;
- empty child items are prohibited;
- only declared fields are allowed;
- metadata `id` is prohibited.

#### H2 and content

- `## Requirement` appears exactly once.
- `## Evidence` appears exactly once.
- `## Required Outcome` appears exactly once only when `status` is valid and equals `accepted`; otherwise it is optional.
- When status is missing or invalid, the Required Outcome condition is indeterminate and produces no additional finding beyond the status finding.
- For accepted Requirements, `## Requirement` and `## Required Outcome` must be substantive under the shared empty-or-`TBD` rule.
- Unlisted headings are allowed when Product authority permits them.

### Decision

#### Metadata

- `status` is one of `proposed`, `accepted`, or `superseded`.
- `date` is strict `YYYY-MM-DD`.
- `depends_on` and `supersedes` are Decision-ref lists; an empty list is represented by the field marker with no child items.
- `migrated_to_spec` is strict `YYYY-MM-DD` or `null`.
- Unlisted metadata and metadata `id` are prohibited.

#### H2

These headings are required exactly once:

- `## Context`;
- `## Decision`;
- `## Rationale`;
- `## Consequences`;
- `## Evidence`.

`## Rejected alternatives` is optional. Unlisted headings are prohibited.

#### Relations

- Direct self-reference checking is reusable but must be explicitly applied to the relevant fields.
- Dependency and supersession cycle checks use separate relation graphs.
- Artifact-specific edge selection and normalization are separate from the generic cycle detector.

### Investigation

#### Metadata

Mandatory fields:

- `status`: `investigating`, `concluded`, or `superseded`;
- strict `date`;
- non-empty `trigger`;
- non-empty `scope`;
- non-empty `non_scope`;
- `source_refs`, where an empty list is allowed;
- `follow_up_candidates`, where an empty list is allowed.

Optional relation fields include `supersedes`, `related_requirements`, `related_work_items`, `related_adrs`, `related_specs`, and `follow_up_results` with field-specific target kinds.

`follow_up_candidates` is string-valued candidate text. Canonical-looking text may be treated as a ref candidate by the field-specific rule, but Task ID text is prohibited.

Unlisted metadata and metadata `id` are prohibited.

#### H2

All declared Investigation headings are required exactly once and unlisted headings are prohibited:

- Investigation scope;
- Out of scope;
- Background;
- What was investigated;
- Findings;
- Cross-cutting observations;
- Follow-up judgment candidates;
- Recommendation;
- Follow-up artifact candidates;
- Open questions.

### Work Item

#### Metadata

- `status` is one of `not_started`, `in_progress`, `blocked`, `done`, or `cancelled`.
- `date` is strict `YYYY-MM-DD`.
- `source_refs` contains one or more refs; duplicate items, empty items, and direct self-reference are prohibited.
- `impact_refs` contains zero or more refs; duplicate and empty items are prohibited.
- `tasks` contains zero or more Task refs; duplicate and empty items are prohibited.
- `source_refs` is the only persisted source-provenance field.
- Unlisted metadata and metadata `id` are prohibited.

#### H2

All declared headings are required exactly once and unlisted headings are prohibited:

- Goal;
- Boundary;
- Impact Scope;
- Task flow;
- Task Candidates;
- Completion Condition;
- Evidence.

#### Relations

Work Item provenance may require artifact-specific edge normalization before generic cycle analysis. The normalization remains an artifact component rather than a generic graph rule.

### Task

#### Identity and relations

- Identity contains inherited `<WORK_SEQUENCE>` and allocated `<TASK_SEQUENCE>`.
- `<WORK_SEQUENCE>` must agree with the referenced parent Work Item identity.
- `work_item` is required and targets a Work Item.
- `work_item_ref` follows task-type-specific Product authority.
- `depends_on` contains Task refs.
- Task dependency cycles are detected from Task-specific edges by the generic cycle detector.

#### Metadata

- `status` is one of `not_started`, `in_progress`, `blocked`, `done`, or `cancelled`.
- `date` is strict `YYYY-MM-DD`.
- `task_type` is one of the accepted workflow task types.
- `estimate` is non-empty.
- `outputs` is `ref_or_literal`: a value interpreted as an allowed canonical ref is validated as a ref; otherwise it is a non-empty literal.
- Task metadata must not contain source-provenance fields such as `source_refs` or `source_requirement`.
- Unlisted metadata and metadata `id` are prohibited.

#### H2

Required exactly once:

- Goal;
- Work;
- Done condition;
- Verification;
- Evidence.

`## Implementation contract` is required for `task_type = implementation` and prohibited otherwise. Unlisted headings are prohibited.

## Gap resolutions

| gap | accepted resolution | canonical synchronization note |
|---|---|---|
| GAP-001 Requirement metadata authority | `status`, `date`, and `source_refs` are persisted required fields. Status uses the Product lifecycle vocabulary. Date is strict `YYYY-MM-DD`. `source_refs` may be an empty list. | Current DRMCP Requirement metadata Specification is stale and must be updated. |
| GAP-002 Required Outcome applicability | Required exactly once only when status is valid and equals `accepted`. Missing or invalid status makes applicability indeterminate; report the status violation only. | Requires artifact-owned conditional validation. |
| GAP-003 unlisted metadata | Any metadata field not declared by the artifact contract is a contract violation. | Shared allow-list recipe with artifact field inventory. |
| GAP-004 YAML front matter | Prohibited for current artifacts. Legacy sources are outside this current validation rule. | Shared current-source rule. |
| GAP-005 H1 title | Required, one line, and non-empty after trimming. No special `TBD` title detection. | Shared H1 validation. |
| GAP-006 sequential file name | File name is `<PUBLIC-ID>.md` or `<PUBLIC-ID>-<slug>.md`. Slug is optional. No slug/title equality rule. | Shared sequential projection validation. |
| GAP-007 reference policy | Reference behavior is field-specific. `ref_or_literal` validates as a ref when interpretable as an allowed canonical ref; otherwise the field-specific literal rule applies. | Reference grammar belongs to artifact/base, not `common`. |
| GAP-008 duplicates | Use one generic duplicate-discovery component. Each artifact validation explicitly opts relevant collections into duplicate prohibition. The likely direction is duplicate prohibition for all persisted ref collections, but complete field coverage still needs one final inventory pass. | Component boundary confirmed; exhaustive mapping pending. |
| GAP-009 direct self-reference | Use a generic direct self-reference checker, explicitly selected by artifact validation. Semantic self-reference through ownership or projections remains artifact-specific. | No generalized semantic-projection framework. |
| GAP-010 cycles | Use a generic directed-cycle detector over supplied edges. Each artifact/relation defines edge extraction and normalization. Different relations remain separate graphs. | Work Item provenance may use a custom edge builder. |
| GAP-011 Topics consistency | Topics refs and physical direct-child refs must be equal as sets; no duplicates; all refs resolve to Specifications; every child `parent` equals the declaring Specification. | Stronger than table-shape-only validation. |
| GAP-012 Reference Specification table | Require at least one Markdown table in an H2 body. Do not validate table schema, row count, or semantic content under this rule. | Artifact Specification validation. |
| GAP-013 recommended heading | Absence is an `advisory`, not a contract violation. | Finding classification belongs to artifact validation. |
| GAP-014 substantive content | Trim the body. Empty or exactly `TBD` is non-substantive. Any other text, including `-`, counts as substantive. | Authoring may later standardize placeholders separately. |
| GAP-015 sentence/block count | No generic sentence or block count. Use only the substantive-content rule unless an artifact contract adds a specific structure. | Avoid prose-style validation. |
| GAP-016 tree directory `index.md` | Every tree directory node requires `index.md`. Missing source is a contract violation, but the node remains discovered, navigable, and part of current subtree structure. The rule applies to tree artifacts generally, including future USDM tree structure. | Current `current-record-scopes.md` advisory classification is stale. |

## Investigation conclusion

The constraints are sufficiently regular to support reusable atomic components, but the meaning of each rule remains artifact-owned.

The inventory does not support one universal hard-coded validator. It supports:

- shared source and structure facts;
- reusable collection, tree, graph, and text algorithms;
- artifact/base validation recipes for genuinely shared artifact constraints;
- artifact-specific recipes and custom fact builders for semantic relations;
- generic validation result projection that does not know artifact meaning.

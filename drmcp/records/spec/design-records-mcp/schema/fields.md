# Reference: Fields

- **id**: `spec:drmcp.design_records_mcp.schema.fields`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines the parsed field vocabulary for current design records and the source rules used to obtain those fields.

## Current contract

DRMCP maps source Markdown into common fields plus kind-specific fields. This specification defines field meaning and source mapping. Public list and retrieval response shapes are owned by `DRMCP-WORK-MCP-004`.

### Common fields

| field | valid record | source | meaning |
|---|---:|---|---|
| `id` | required | Kind-specific canonical identity rule | Canonical current record ID. |
| `kind` | required | Discovery path kind | `decision`, `spec`, `investigation`, `requirement`, `work_item`, or `task`. |
| `title` | required | Real ATX H1 | Human-readable title. |
| `status` | required | H1-adjacent metadata | Lifecycle state carried by the source record. |
| `date` | required | H1-adjacent metadata | Source date under the owning PRODUCT authoring or spec-format rule. |

An invalid but addressable record may lack a required field or carry a parsed value that fails validation. DRMCP preserves parsed values and does not fabricate replacements.

The physical source path belongs to source provenance defined by `spec:drmcp.design_records_mcp.schema.record_model`; it is not a normalized semantic record field.

### Kind-specific fields

| kind | fields sourced from the current record |
|---|---|
| `decision` | `depends_on`, `supersedes`, `migrated_to_spec` |
| `spec` | `spec_kind`, `parent`, conditional `contract_class` |
| `investigation` | `trigger`, `scope`, `non_scope`, `source_refs`, `follow_up_candidates`, optional `supersedes`, `related_requirements`, `related_work_items`, `related_adrs`, `related_specs`, `related_internal_design`, `related_coverage`, `follow_up_results` |
| `requirement` | `source_refs`, `work_items`, optional `subdomain` |
| `work_item` | `source_requirement`, `impact_refs`, `tasks`, optional `subdomain` |
| `task` | `work_item`, `source_requirement`, `estimate`, `depends_on`, `outputs`, optional `subdomain` |

Current specs do not define `depends_on`, `supersedes`, or `migrated_to_spec` fields. DRMCP does not create empty or null spec values for those obsolete mappings.

### Canonical `id`

Current sequential artifacts use the complete canonical app-aware artifact ID directly.

```text
ADR / investigation / requirement / work item:
<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<SEQUENCE>

Task:
<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>
```

- `<SEQUENCE>` and `<WORK_SEQUENCE>` are three-digit zero-padded decimal values.
- `<TASK_SEQUENCE>` is a two-digit zero-padded decimal value.
- Subdomains are not ID segments.
- DRMCP does not strip and later reattach an app prefix.
- Bare forms such as `REQ-*` or `WORK-*` are grammar fragments, not canonical current IDs.

Examples, not exhaustive:

| kind | synthetic example |
|---|---|
| `decision` | `EXAMPLEAPP-ADR-ARCH-001` |
| `investigation` | `EXAMPLEAPP-INV-DATA-002` |
| `requirement` | `EXAMPLEAPP-REQ-SEARCH-003` |
| `work_item` | `EXAMPLEAPP-WORK-SEARCH-004` |
| `task` | `EXAMPLEAPP-TASK-SEARCH-004-01` |

These examples illustrate the grammar only. They do not define an app-specific contract.

Current spec IDs are path-derived canonical `spec:` refs. They follow `spec:product.design_records.spec_format.spec_id_as_ref` and `spec:drmcp.design_records_mcp.schema.id_normalization`.

### Identity source by kind

| kind | canonical identity authority | required consistency checks |
|---|---|---|
| `decision` | Complete canonical ID in H1 | Filename ID prefix must match. ADR metadata has no `id` field. |
| `investigation` | Complete canonical ID in H1 | Filename ID prefix must match. Investigation metadata has no `id` field. |
| `requirement` | Complete canonical ID in H1 | Required metadata `id` and filename ID prefix must match H1. |
| `work_item` | Complete canonical ID in H1 | Required metadata `id` and filename ID prefix must match H1. |
| `task` | Complete canonical ID in H1 | Required metadata `id` and filename ID prefix must match H1. |
| `spec` | Canonical ref derived from configured app namespace and spec-relative path | H1-adjacent metadata `id` must match the derived ref. |

For every sequential kind, H1 is the canonical identity authority. Metadata `id` and filename prefixes are consistency values only and never become fallback identity sources.

- A valid H1 with missing, malformed, or mismatched metadata remains addressable under the H1 canonical ID and fails validation.
- A missing or malformed H1 leaves a sequential source without a canonical ID even when metadata or filename text resembles one.
- DRMCP does not infer a missing sequential ID from metadata or filename.

### H1 and title extraction

Sequential artifact H1 format:

```markdown
# <CANONICAL_ARTIFACT_ID>: <Title>
```

DRMCP parses the complete canonical ID before the first ASCII colon and the non-empty title after it. It validates the ID against the PRODUCT artifact grammar without prefix stripping or case repair.

Current spec H1 format:

```markdown
# <SpecKind>: <Title>
```

`spec_kind` is the trimmed value before the first ASCII colon. `title` is the non-empty trimmed value after it. The accepted spec-kind set is consumed from `spec:product.design_records.spec_format.document_shape`; DRMCP does not maintain a separate local set.

Title is not inferred from the filename, metadata, or body when H1 is missing or malformed.

### `status`

| kind | PRODUCT-defined vocabulary |
|---|---|
| `decision` | `proposed`, `accepted`, `superseded` |
| `spec` | Marker is required, but no complete allowed-value set is currently defined. |
| `investigation` | `investigating`, `concluded`, `superseded` |
| `requirement` | `captured`, `decision_needed`, `accepted`, `deferred`, `rejected` |
| `work_item` | `not_started`, `in_progress`, `blocked`, `done` |
| `task` | `not_started`, `in_progress`, `blocked`, `done` |

DRMCP consumes these vocabularies from PRODUCT authority. It does not invent a spec status vocabulary while PRODUCT leaves the complete set undefined.

### Relation and list fields

- ADR `depends_on` and `supersedes` contain complete canonical artifact IDs.
- Task `depends_on` contains complete canonical task IDs.
- Workflow relation fields use complete canonical artifact IDs for their declared target kind.
- Spec references use active canonical `spec:` refs.
- Empty-list normalization follows the owning PRODUCT authoring rule.
- Resolution, reciprocal-link validation, and diagnostic identifiers are defined outside this field-vocabulary contract.

### Retrieval-only content

Markdown headings and verbatim body content are retrieval projections, not persisted semantic metadata fields. Their public inclusion and representation are owned by `DRMCP-WORK-MCP-004`.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Canonical sequential artifact ID grammar. |
| `spec:product.design_records.spec_format.document_shape` | Current spec H1 kind and metadata requiredness. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Current spec canonical ref authority. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Addressable record and invalid-source retention behavior. |
| `spec:drmcp.design_records_mcp.schema.id_normalization` | Canonical identity mapping by kind. |
| `spec:drmcp.design_records_mcp.schema.metadata_grammar` | Source metadata parsing grammar. |

## Sources

- `DRMCP-TASK-MCP-003-04`: Shared field, invalid-source, and current identity decisions.
- PRODUCT authoring standards for ADR, investigation, requirement, work item, task, and spec records.

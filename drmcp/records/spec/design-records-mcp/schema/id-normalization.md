# Reference: ID normalization

- **id**: `spec:drmcp.design_records_mcp.schema.id_normalization`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines canonical current identity mapping for sequential design records and path-derived spec records.

## Current contract

### Sequential artifact canonical ID

Current sequential artifacts use one complete canonical app-aware artifact ID. DRMCP validates that complete value directly; it does not split identity into a runtime prefix and a bare ID and does not reconstruct an ID after parsing.

General grammar:

```text
ADR / investigation / requirement / work item:
<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<SEQUENCE>

Task:
<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>
```

| segment | rule |
|---|---|
| `<APP_NAMESPACE>` | Owning app namespace. |
| `<ARTIFACT_KIND>` | `ADR`, `INV`, `REQ`, or `WORK` for the corresponding current kind. |
| `<DOMAIN_NAMESPACE>` | Owning domain namespace. Subdomains are excluded. |
| `<SEQUENCE>` | Three-digit zero-padded decimal sequence. |
| `<WORK_SEQUENCE>` | Three-digit sequence inherited from the parent Work Item. |
| `<TASK_SEQUENCE>` | Two-digit zero-padded sequence within the parent Work Item. |

Examples, not exhaustive:

| kind | synthetic canonical ID |
|---|---|
| decision | `EXAMPLEAPP-ADR-ARCH-001` |
| investigation | `EXAMPLEAPP-INV-DATA-002` |
| requirement | `EXAMPLEAPP-REQ-SEARCH-003` |
| work item | `EXAMPLEAPP-WORK-SEARCH-004` |
| task | `EXAMPLEAPP-TASK-SEARCH-004-01` |

These values are examples only. They do not define an app-specific namespace or compatibility rule.

### Sequential identity source and agreement

| kind | canonical identity authority | required agreement |
|---|---|---|
| decision | Complete canonical ID in H1 | Filename ID prefix matches H1. |
| investigation | Complete canonical ID in H1 | Filename ID prefix matches H1. |
| requirement | Complete canonical ID in H1 | Required metadata `id` and filename ID prefix match H1. |
| work item | Complete canonical ID in H1 | Required metadata `id` and filename ID prefix match H1. |
| task | Complete canonical ID in H1 | Required metadata `id` and filename ID prefix match H1. |

For every current sequential artifact, H1 is the canonical identity authority.

- ADR and investigation metadata do not carry an `id` field under current PRODUCT authoring authority.
- Requirement, work-item, and task metadata `id` values are required consistency values, not independent or fallback identity sources.
- Filename ID prefixes are consistency values only and never become fallback identity sources.
- When H1 contains one valid complete canonical ID, the source remains addressable under that ID even if metadata is missing, malformed, or mismatched.
- When H1 does not contain a valid complete canonical ID, metadata and filename text do not create one; the source remains validation-only.

Current identity comparison is exact. DRMCP does not repair case, surrounding whitespace, missing app namespace, missing domain namespace, or malformed sequence segments.

Authority: `spec:product.design_records.namespace_model.artifact_id_grammar` and the kind-specific PRODUCT authoring standards.

### Current spec path-derived canonical ref

Current spec identity is derived from the configured `app_namespace` and the file path relative to `<records_root>/spec/`.

Derivation steps:

1. Resolve the file path relative to the configured `records_root` and normalize path separators.
2. Remove the leading `spec/` portion.
3. Remove the `.md` extension.
4. Omit a final segment equal to `index`.
5. Convert remaining path separators to `.`.
6. Convert hyphens within each remaining segment to underscores.
7. Prefix the result with `spec:<app_namespace>`.

Result form:

```text
root index:
spec:<app_namespace>

non-root spec:
spec:<app_namespace>.<segment>(.<segment>)*
```

Each app and topic segment must match:

```text
[a-z0-9][a-z0-9_]*
```

DRMCP does not lowercase or otherwise repair invalid segments. An invalid identity segment leaves the candidate without a determinable canonical spec ref.

Examples, not exhaustive:

| spec-relative path | configured app namespace | canonical ref |
|---|---|---|
| `spec/index.md` | `exampleapp` | `spec:exampleapp` |
| `spec/search/index.md` | `exampleapp` | `spec:exampleapp.search` |
| `spec/search/query-contract.md` | `exampleapp` | `spec:exampleapp.search.query_contract` |

These examples illustrate derivation only and do not define repository-specific topics.

The path-derived canonical ref is authoritative. The H1-adjacent metadata `id` is a required visible consistency value and must match exactly.

A stale or mismatched metadata `id`:

- does not replace the path-derived ref;
- does not become an alias;
- does not become an alternate lookup key;
- is not rewritten automatically.

No legacy sequential spec ID, filename-derived alias, or implicit redirect participates in current spec identity unless a separate compatibility contract explicitly defines it.

Authority: `spec:product.design_records.spec_format.spec_id_as_ref`.

### Invalid and conflicting identity

- A source with one determinable canonical ID remains addressable even when other content is invalid.
- A source with no determinable canonical ID remains validation-only.
- When multiple sources produce the same canonical ID, the identity has no winner and all source paths remain validation inputs.

The shared behavior is defined by `spec:drmcp.design_records_mcp.schema.record_model` and `spec:drmcp.design_records_mcp.schema.discovery`.

### Legacy boundary

Issued legacy IDs are handled only by separately configured compatibility and legacy-index contracts. They do not define the grammar, parsing model, or identity source for current records.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Authority for current sequential artifact IDs. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Authority for current path-derived spec refs. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Shared addressable, invalid, and conflicting source behavior. |
| `spec:drmcp.design_records_mcp.schema.discovery` | Candidate discovery and invalid-source behavior. |
| `spec:drmcp.design_records_mcp.schema.record_source` | Source mapping by record kind. |

## Sources

- `DRMCP-TASK-MCP-003-03`: Current spec identity decisions.
- `DRMCP-TASK-MCP-003-04`: Shared current identity and legacy-boundary decisions.

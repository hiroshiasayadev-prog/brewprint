# Reference: Namespace scanning

- **id**: `spec:drmcp.design_records_mcp.namespace_scanning`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Defines DRMCP configuration and validation for current records roots, app association, active-index construction, and separation from optional legacy roots.

PRODUCT specifications own app namespace semantics, Design Records placement, and record path patterns.
This specification defines how DRMCP consumes those authorities.

## Configuration contract

DRMCP configuration declares one `repository_root` and a non-empty `current_roots` collection.

Each current-root entry contains:

| field | contract |
|---|---|
| `app_namespace` | Explicit app namespace associated with the root. It is not inferred from the filesystem. |
| `records_root` | Repository-relative path to the app's current Design Records root. |

Conceptual example:

```yaml
repository_root: C:/work/brewprint
current_roots:
  - app_namespace: product
    records_root: product/records
  - app_namespace: drmcp
    records_root: drmcp/records
```

The concrete configuration serialization is owned by the runtime configuration contract.
The semantic requirements in this specification apply regardless of serialization.

## Current-root rules

- `current_roots` must contain at least one entry.
- One `app_namespace` may appear in exactly one current-root entry.
- One canonical filesystem root may appear in exactly one current-root entry.
- `records_root` is resolved relative to the configured `repository_root`, not relative to the process working directory.
- The portable configuration form does not require a host-specific absolute `records_root`.
- After repository-relative normalization, each entry must satisfy `records_root == <app_namespace>/records`.
- The explicit `app_namespace` and `records_root` values are checked for consistency; the namespace is not derived from the path.
- DRMCP does not auto-discover `*/records/` directories.
- Folder-local ignore files do not add, remove, or redefine configured current roots.

The placement check consumes `spec:product.design_records.repository_layout`.
It does not redefine the generic placement rule.

## Root validation

Every configured current root is mandatory.

Active-index construction fails as a whole when any configured current root is:

- missing;
- not a directory;
- unreadable;
- outside the configured `repository_root` after canonical resolution;
- inconsistent with its explicit `app_namespace`;
- duplicated by another current-root entry;
- canonically identical to another configured root where uniqueness is required;
- overlapping a configured legacy root.

DRMCP does not omit an invalid configured current root and continue with a partial active index.
Exact diagnostic identifiers and response representation are owned by the validation and diagnostics contracts.

A valid current root may contain zero discoverable current records.
Such a root contributes an empty app-scoped portion to the active index and does not make configuration invalid.
Root validity and discovered record count are separate concerns.

## Active-index construction

DRMCP builds one active read index over discovered current records under all configured current roots.
Current parsing contracts determine whether a source becomes a unique addressable entry, a conflicted entry, or a validation-only source.

The active index:

- contains current records only;
- retains the explicit app association of each source root;
- uses current canonical identities defined by PRODUCT authorities and mapped by DRMCP parsing contracts;
- supports records from multiple configured app namespaces in one operational index;
- excludes all configured legacy roots and legacy archive records;
- marks a canonical identity as conflicted when two or more discovered current records produce that identity;
- does not create one addressable active-index entry for a conflicted identity;
- keeps unaffected current records available to normal read operations;
- retains all conflicting sources as validation inputs with source provenance;
- never selects one duplicate by filesystem traversal order.

Record-kind path patterns are consumed from `spec:product.design_records.repository_layout.record_discovery_paths`.
Current spec parsing and path-derived spec identity are defined by the current spec discovery contracts, not by this specification.

Duplicate canonical identity is an index-entry conflict, not a configured-root failure.
A duplicate does not invalidate the configured discovery scope or make unrelated records unavailable.
Exact diagnostics and response representation remain owned by the validation and read-operation contracts.

## Current and legacy root separation

Legacy compatibility is optional and configuration-gated.
DRMCP does not auto-discover `v01/` or any other legacy archive directory.

Current and legacy roots remain separate configuration and indexing scopes.
After path normalization and canonical filesystem resolution, a current root and legacy root must not:

- identify the same directory;
- place one root inside the other;
- resolve through links or aliases to overlapping filesystem trees.

Any equality, ancestor relationship, or descendant relationship between current and legacy roots is a configuration error.

Each `legacy_roots` entry contains one required `records_root` field.
The path is repository-relative and is resolved from `repository_root`.
A legacy-root entry does not declare an app namespace, namespace prefix, or archive identity.

Missing `legacy_roots` and an explicit empty list are equivalent valid current-only configurations.
Either form disables legacy fallback.

When one or more legacy roots are configured, every entry is mandatory.
Each configured legacy root must:

- exist;
- be a readable directory;
- remain inside `repository_root` after canonical resolution;
- be unique after canonicalization;
- not equal, contain, or be contained by another configured legacy root;
- not equal, contain, or be contained by a configured current root.

Link or alias resolution that creates equality or overlap is invalid.
Any invalid configured legacy root fails DRMCP startup.
DRMCP does not omit one invalid root or build a partial legacy lookup map from remaining roots.
A valid legacy root may contain zero accepted legacy sources.

Configured legacy roots build one separate read-only exact lookup map.
DRMCP scans eligible regular Markdown files recursively and derives an issued legacy ID from the filename mapping below.

### Legacy issued-ID lexical mapping

The PRODUCT compatibility authority owns which V01 families are accepted.
DRMCP owns the lexical parser mapping for those families so one exact issued ID can be recognized in a resolver or retrieval input and extracted from an archived source filename.
The mapping does not authorize another legacy family or change the PRODUCT retention policy.

| accepted family | exact issued-ID grammar |
|---|---|
| decision | `V01-ADR-<SEQUENCE>` |
| investigation | `V01-INV-<DOMAIN>-<SEQUENCE>` |
| requirement | `V01-REQ-<DOMAIN>-<SEQUENCE>` |
| work item | `V01-WORK-<DOMAIN>-<SEQUENCE>` |
| task | `V01-TASK-<DOMAIN>-<WORK_SEQUENCE>-<TASK_SEQUENCE>` |

Lexical tokens:

- `<DOMAIN>` matches `[A-Z][A-Z0-9]*`;
- `<SEQUENCE>` and `<WORK_SEQUENCE>` each match exactly three ASCII digits `[0-9]{3}`;
- `<TASK_SEQUENCE>` matches exactly two ASCII digits `[0-9]{2}`;
- all literal tokens and comparisons are case-sensitive.

After one complete issued ID, the filename must end in exactly one of these forms:

- `<issued-id>.md`;
- `<issued-id>-<slug>.md`, where `<slug>` is non-empty and matches `[a-z0-9][a-z0-9-]*`.

The parser matches the complete family-specific issued-ID grammar from the beginning of the filename stem.
The suffix boundary begins only after the complete sequence fields required by that family.
For example, `V01-TASK-MCP-003-01-description.md` yields `V01-TASK-MCP-003-01`.
A filename that does not satisfy exactly one mapping is not a legacy lookup candidate.
`V01-SPEC-*` has no mapping and is never adopted.
The mapping is not an alias, fuzzy-match, repair, path-inference, or prefix-completion registry.
Only one exact accepted issued-ID grammar may create a legacy lookup key.
Operation-specific rejection and warning-trigger behavior is defined by `spec:drmcp.design_records_mcp.tools.resolve_reference` and `spec:drmcp.design_records_mcp.tools.get_records`.
The configured root path is internal provenance only and does not qualify or rewrite the issued ID.

### Legacy enumeration boundary

Symlinked files, junction files, reparse-point files, and other filesystem-alias files are not legacy candidates.
Symlinked directories, junction directories, reparse-point directories, and other filesystem aliases are not traversed.
Every otherwise eligible candidate is canonicalized before adoption and must remain within both its configured legacy root and `repository_root` and outside every configured current root.
A candidate that violates this source boundary is excluded from the lookup map; it does not invalidate the configured legacy root or remove unrelated legacy entries.
Diagnostic category, severity, message, and source-location representation remain owned by `DRMCP-WORK-MCP-006`.

When two or more source files produce the same issued legacy ID, that ID has no selected winner.
Filesystem traversal order must not select one source.
Unrelated unique legacy IDs remain available.

The legacy lookup map never contributes records to the active index or `list_records`.
Legacy sources remain excluded from current repository-wide validation and all authoring targets.
The absence of legacy roots does not affect current-only operation.
Legacy-only startup is not supported because at least one current root is required.

## Explicit exclusions

This specification does not define:

- current spec metadata parsing;
- path-derived `spec:` identity details;
- sequential record field parsing;
- query filters, ordering, ranges, or pagination;
- exact retrieval behavior;
- resolver fallback order;
- validation diagnostic identifiers or response shapes;
- fixture design;
- runtime implementation mechanics;
- authoring target selection.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.namespace_model` | Authority for app namespace and current identity semantics. |
| `spec:product.design_records.repository_layout` | Authority for `<app>/records/` placement. |
| `spec:product.design_records.repository_layout.record_discovery_paths` | Authority for app-independent record path patterns. |
| `spec:product.design_records.spec_format` | Authority for current spec format and path-derived identity. |
| `spec:product.brewprint.compatibility.legacy_id_compatibility` | Authority for the accepted V01 sequential-family set. |
| `spec:drmcp.design_records_mcp.resolver` | Current-first orchestration and configured legacy-fallback outcomes. |
| `spec:drmcp.design_records_mcp.tools.resolve_reference` | Public resolver rejection and outcome behavior. |
| `spec:drmcp.design_records_mcp.tools.get_records` | Exact-retrieval classification and rejected string-item behavior. |
| `spec:drmcp.design_records_mcp.schema.discovery` | DRMCP record inclusion and current source-format discovery. |
| `spec:drmcp.design_records_mcp.schema.record_source` | DRMCP source representation for discovered records. |
| `spec:drmcp.design_records_mcp.schema.record_model` | DRMCP indexed record representation and canonical identity handling. |
| `spec:drmcp.design_records_mcp.schema.diagnostics` | Diagnostic identifiers and representations for configuration and indexing failures. |

## Sources

- `DRMCP-ADR-MCP-001`: Current-format-first operation, configuration-gated legacy fallback, and separate indexes.
- `DRMCP-REQ-MCP-001`: Configured current roots, active-index requirements, and root validation obligations.
- `DRMCP-TASK-MCP-003-02`: Accepted configured-root and index-separation decisions.
- `DRMCP-TASK-MCP-005-03`: Configured legacy-root and minimal lookup-map decisions.
- `DRMCP-TASK-MCP-005-04`: Rejected-input and operation-pointer synchronization.

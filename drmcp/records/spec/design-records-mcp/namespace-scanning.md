# Reference: Namespace scanning

- **id**: `spec:drmcp.design_records_mcp.namespace_scanning`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Defines the v1 `namespace_prefix` derivation formula, kind-level prefix application table, and multi-root scan behavior.

## Current contract

### namespace_prefix derivation

`records_root` takes the relative path form `<app-namespace>/records` from the repository root. The namespace prefix is derived mechanically from the app namespace directory name (the parent of `records/`).

Derivation formula: `namespace_prefix = strings.ToUpper(appNamespaceDir) + "-"`

| records_root | appNamespaceDir | namespace_prefix |
|---|---|---|
| `v01/records` | `v01` | `V01-` |
| `drmcp/records` | `drmcp` | `DRMCP-` |

### Kind-level prefix application

Artifacts within a records tree carry the namespace prefix at the following locations by record kind.

| kind | filename | H1 | metadata id |
|---|---|---|---|
| ADR | ✓ `V01-ADR-NNN-slug.md` | ✓ `# V01-ADR-NNN: title` | none |
| Spec | none (slug only) | none (arbitrary title) | ✓ front matter `design_record.id: V01-SPEC-slug` |
| Investigation | ✓ `V01-INV-DOMAIN-NNN-slug.md` | ✓ `# V01-INV-DOMAIN-NNN: title` | none |
| Workflow (REQ / WORK / TASK) | ✓ `V01-WORK-DOMAIN-NNN-slug.md` | ✓ `# V01-WORK-DOMAIN-NNN: title` | ✓ bullet `- **id**: V01-WORK-DOMAIN-NNN` |

The parser strips the namespace prefix from the kind-specific location to extract and validate the bare ID, then returns the prefixed full ID as `record.ID`. Bare ID grammar is defined in `spec:drmcp.design_records_mcp.schema.id_normalization`.

### Multi-root scan

**Default (multi-root) mode:** Auto-discovers all `*/records/` directories within the repository, derives the namespace prefix for each, and builds a unified index.

**Single-root mode:** When a single `--records-root <path>` is specified, only that records tree is targeted (backward-compatible behavior).

In multi-root mode, records from different app namespaces can be retrieved and ref-resolved in the same query. Public IDs across different namespace prefixes do not collide. Cross-namespace relations are held in the index and can be resolved.

> Source: V01-ADR-097, V01-ADR-099

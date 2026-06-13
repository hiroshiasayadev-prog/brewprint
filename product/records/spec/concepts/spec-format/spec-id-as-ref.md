# Concept: Spec ID-as-ref

- **id**: `spec:product.concepts.spec_format.spec_id_as_ref`
- **status**: accepted
- **date**: 2026-06-14
- **parent**: `spec:product.concepts.spec_format`

## What this is

This spec defines the local spec identity model, path-derived canonical spec ref derivation, and interim parent reference grammar for the spec-format area.

The visible H1-adjacent `id` is the canonical document-level spec ref. For new or migrated specs, that value must match the path-derived canonical ref. Compatibility behavior for aliases, redirects, split, merge, move, rename, and legacy metadata is exceptional follow-up owned by PRODUCT-WORK-SPEC-002.

## Concept model

| concept | rule |
|---|---|
| canonical spec ref | For new or migrated specs, the path-derived `spec:` ref is the canonical document-level spec ref. |
| H1-adjacent ID | The H1-adjacent `- **id**:` marker records the path-derived canonical ref visibly near the H1. |
| path/ref mapping | Path to ref and ref to path are one-to-one for new or migrated specs under this format. |
| rename / move behavior | Moving or renaming a spec changes the canonical spec ref. Old IDs are not preserved by default. |
| mismatch validation | Validators must report an error when a new or migrated spec's visible `id` does not match its path-derived canonical ref. Inventory, migration, or transient working states may emit warnings. |
| compatibility exceptions | Alias, redirect, stale-ref, split, merge, move, rename, and legacy metadata compatibility are explicit follow-up design, not the default identity model. |
| section refs | Section refs are not listed in front matter. If section-level refs are needed, a visible table must be introduced by a later compatibility design. |

## Rules

### Path-derived canonical spec ref derivation

Spec document IDs are path-derived canonical refs. They are written in the H1-adjacent `- **id**:` marker and use the `spec:` prefix.

| rule | contract |
|---|---|
| base prefix | `<app>/records/spec/` maps to `spec:<app>.` using lower-case app namespace. |
| path separator | Directory separators under `records/spec/` map to `.`. |
| extension | `.md` is removed. |
| `index.md` | `index` is omitted because `index.md` represents the entrypoint topic for its containing directory. |
| non-index file | File stem is kept as the final spec ID segment. |
| word separator | Hyphenated path/title segments are normalized to underscore in spec IDs. |
| write location | The derived ref is written to H1-adjacent `- **id**:`, not to front matter. |
| one-to-one mapping | For new or migrated specs, each canonical ref maps to exactly one spec path and each spec path maps to exactly one canonical ref. |
| rename / move | A path change changes the canonical spec ref. Old refs are not preserved unless a later compatibility design explicitly adds an alias or redirect. |
| mismatch handling | Validator must error when a new or migrated spec's visible `id` does not match the path-derived canonical ref. Inventory, migration, or transient working states may warn. |
| rewrite behavior | Validator must not rewrite IDs automatically; authoring or migration tooling must make ID-changing edits explicit. |

Examples:

| path | canonical spec ref |
|---|---|
| `product/records/spec/concepts/spec-format/index.md` | `spec:product.concepts.spec_format` |
| `product/records/spec/concepts/traceability/index.md` | `spec:product.concepts.traceability` |
| `product/records/spec/concepts/traceability/semantic-ref.md` | `spec:product.concepts.traceability.semantic_ref` |
| `drmcp/records/spec/design-records-mcp/tools.md` | `spec:drmcp.design_records_mcp.tools` |

This contract intentionally uses underscore as the word separator inside spec ID segments. Existing traceability docs that still require hyphen-only `spec:` segments must be updated or explicitly compatibility-mapped before this format is enforced.

### Parent reference grammar

H1-adjacent `parent` markers use this grammar:

```text
parent := "-" | "root" | spec_ref
spec_ref := "spec:" segment ("." segment)*
segment := [a-z0-9][a-z0-9_]*
```

| parent form | status | reason |
|---|---|---|
| `root` | allowed | Declares a top-level topic root. |
| `-` | allowed | Equivalent root marker for compact tables. |
| canonical `spec:` ref | allowed | Parent identity is the path-derived canonical spec ref. |
| physical path | prohibited | Stale on move; not a canonical reference. |
| file name | prohibited | Ambiguous and stale on rename. |
| H1 title string | prohibited | Stale on title change. |
| Markdown heading anchor | prohibited | Parser/render dependent and stale on heading change. |
| title-derived topic ref | prohibited in this contract | Deferred to PRODUCT-WORK-SPEC-002 if still needed. |

## Boundary

Compatibility exceptions were evaluated as part of PRODUCT-WORK-SPEC-002. None require a dedicated spec or mapping table.

| compatibility class | decision |
|---|---|
| Alias and redirect | Not needed now. No alias or redirect table introduced. If required during migration, a new work item must scope it explicitly. |
| Derived topic refs | Prohibited by the parent grammar above. No compatibility period. |
| Split, merge, move, and rename stale-ref behavior | Not needed now. No pending restructuring requires backward-compatible old refs. |
| Legacy `semantic_refs` / `sections` in front matter | Handled by `spec:product.concepts.spec_format.validation_policy` severity rules: warning for existing specs, error for new or migrated specs. No parsing shim introduced. |
| Hyphen-form IDs (e.g., `spec:trace.semantic-ref`) | Validation-policy warning severity applies until PRODUCT-WORK-SPEC-005 migration. No bidirectional compatibility table introduced. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.concepts.spec_format` | Parent Index for this concept. |
| `spec:product.concepts.spec_format.topics_table` | Uses canonical refs in authoritative topic rows. |
| `spec:product.concepts.spec_format.validation_policy` | Severity rules for existing unmigrated specs with hyphen-form IDs or legacy front matter. |

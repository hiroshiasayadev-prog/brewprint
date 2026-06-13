# Contract: Topics table

- **id**: `spec:product.concepts.spec_format.topics_table`
- **status**: accepted
- **date**: 2026-06-11
- **parent**: `spec:product.concepts.spec_format`
- **contract_class**: `format`

## What this is

This spec defines the `## Topics` table contract for `Index` specs and any `Overview` spec that acts as the authoritative topic parent for child specs.

It owns the table columns, parent declaration behavior, and duplicate-parent rule. It does not define aliases, redirects, or graph validation beyond the local parent declaration contract.

## Current contract

An `Index` spec, or an `Overview` spec that intentionally declares child topics, conforms to this contract by containing a navigation-first `## Topics` table with required columns `title`, `kind`, `ref`, and `summary`.

Each row targets the child spec by canonical `spec:` ref. Tooling resolves `ref` to a path through the path-derived canonical mapping defined by `spec:product.concepts.spec_format.spec_id_as_ref`; the table does not carry canonical physical file paths.

## Rules

| rule | contract |
|---|---|
| pure Index | `Index` specs are navigation-first and should contain `## What this is` and `## Topics` only, except for brief explanatory notes. |
| Overview+Topics | `Overview` specs may contain `## Topics` when splitting a substantive overview from its topic table would add unnecessary churn. |
| authoritative parent | A child spec must have exactly one authoritative parent declaration from an `Index` or `Overview` `## Topics` row plus a matching H1-adjacent `parent` marker in the child spec. |
| duplicate parent | Duplicate parent declarations for the same child `ref` are invalid. |
| filename implication | `index.md` does not imply `Index`; H1 kind controls the spec kind. |
| path resolution | Physical paths are derived from canonical refs; `file` is not a canonical `## Topics` column. |
| row-level parent | `parent` is not required per row. Parent identity is declared by the owning Index or Overview and the child spec's H1-adjacent `parent`. |

## Validation rules

Required `## Topics` table columns:

| column | required | meaning | validation |
|---|---:|---|---|
| `title` | yes | Human-readable child topic title. | Non-empty. Truthfulness is human-reviewed. |
| `kind` | yes | Child spec kind. | Must be one of accepted spec kinds. |
| `ref` | yes | Canonical child `spec:` ref. | Must match the path-derived canonical ref for exactly one spec file. |
| `summary` | yes | Short review/navigation summary. | Non-empty. Semantic accuracy is human-reviewed. |

The `file` column is not canonical and is not required by this contract. A tool may display derived paths, but it must not treat displayed paths as the source of truth for child identity.

For `## Topics` rows, the row parent is the declaring spec's H1-adjacent `id`. Cross-parent declarations and row-level parent exceptions are deferred to PRODUCT-WORK-SPEC-002.

## Errors

| condition | severity |
|---|---|
| Missing `## Topics` table in an `Index` spec | Error for new or migrated specs. |
| Missing required column | Error for migrated or new Index / Overview+Topics specs. |
| `file` is used instead of `ref` as the canonical child target column | Error for migrated or new Index / Overview+Topics specs. |
| Row-level `parent` is required or treated as canonical | Error for migrated or new Index / Overview+Topics specs. |
| Unresolved child `ref` | Error for migrated or new Index / Overview+Topics specs. |
| Duplicate parent declaration for the same child `ref` | Error. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.concepts.spec_format` | Parent Index for this contract. |
| `spec:product.concepts.spec_format.document_shape` | Defines `Index` and `Overview` kind expectations. |
| `spec:product.concepts.spec_format.spec_id_as_ref` | Defines parent reference grammar used by this table. |

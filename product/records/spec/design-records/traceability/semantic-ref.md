# Reference: Semantic ref

- **id**: `spec:product.design_records.traceability.semantic_ref`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.design_records.traceability`

## What this is

Defines the current `spec:` canonical ref class for Design Records traceability. It points to the authoritative path-derived spec identity contract and records dispositions for obsolete semantic-ref assumptions.

## Current spec ref model

| concept | current rule |
|---|---|
| Canonical spec ref | New and migrated specs use the path-derived `spec:` ref. |
| Ref location | The canonical ref is written in the visible H1-adjacent `id`. |
| Path/ref relation | Path and canonical ref are one-to-one for new and migrated specs. |
| Move or rename | Moving or renaming a spec changes the canonical ref. Old refs are not retained by default. |
| Derivation grammar | Owned by `spec:product.design_records.spec_format.spec_id_as_ref`; this file does not redefine it. |
| Section refs | Not active. A later visible-table contract is required before section refs become canonical. |

Physical paths are repository locations. They are not canonical relation values.

## Ref classes

| class | identity rule |
|---|---|
| Spec ref | Path-derived document-level `spec:` ref. |
| Record ID-as-ref | Complete public record ID for ADR, investigation, requirement, work item, or task. |
| Legacy issued ID | Complete legacy public ID preserved through Brewprint compatibility. |

Spec refs and record ID-as-refs are distinct. A record ID-as-ref is not a semantic prefix and not a bare kind grammar fragment.

## Obsolete assumptions

| removed assumption | final disposition |
|---|---|
| Append-only `spec:` stability independent of path | Superseded by `spec:product.design_records.spec_format.spec_id_as_ref`. New and migrated specs change canonical ref on move or rename. |
| Preserve the same canonical ref after file rename or move | Superseded by the path-derived identity model. |
| Preserve section refs through heading rename or section move | Delete. No current section-ref contract exists. |
| Front-matter `semantic_refs` declaration examples | Delete. H1-adjacent `id` is the current visible spec identity. |
| Front-matter `sections` heading-map examples | Delete. A visible-table contract is required before section refs become canonical. |
| Redirect and superseded placeholder schemas | Delete. Alias, redirect, stale-ref, split, merge, move, and rename compatibility are not current defaults. |
| Hyphen-only semantic-ref grammar | Superseded by the spec-format path-derived grammar with underscore ID segments. |
| `internal-design:` and `coverage:` as deferred semantic endpoints in this file | Existing evidence owners preserve the rationale; no current endpoint is defined here. |

## Non-goals

- Redefining the spec-format path derivation grammar.
- Defining aliases, redirects, stale-ref compatibility, split/merge compatibility, or superseded-ref tables.
- Defining current section-level canonical refs.
- Defining `yaml:`, `internal-design:`, `coverage:`, `fixture:`, or workflow semantic prefixes.
- Defining DRMCP request, response, parser, diagnostic, or writer behavior.

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.spec_format.spec_id_as_ref` | Authoritative path-derived spec identity contract. |
| `spec:product.design_records.traceability.artifact_refs` | Supported active reference classes. |
| `spec:product.design_records.traceability.metadata_schema` | Visible metadata and relation boundary. |

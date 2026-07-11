# USDM requirement: Spec identity

- **id**: `usdm:product.design_records.namespace_and_identity.spec_identity`
- **status**: draft
- **date**: 2026-07-09
- **kind**: requirement
- **parent**: `usdm:product.design_records.namespace_and_identity`

## What this is

Defines `spec:` ref identity requirements for Specifications.

This record includes path-derived spec refs, H1-adjacent IDs, and parent ref grammar.
This record does not include workflow artifact IDs, alias/redirect implementation, or section refs.

## Requirements: spec:product.design_records.spec_format.spec_id_as_ref

| id | requirement | notes |
|---|---|---|
| R001 | Implementations MUST derive canonical spec refs from paths for new or migrated specs. |  |
| R002 | Implementations MUST write the canonical spec ref in the H1-adjacent `id` marker. |  |
| R003 | Implementations MUST map `<app>/records/spec/` to `spec:<app>.` using lowercase app namespace. |  |
| R004 | Implementations MUST map directory separators under `records/spec/` to dots in spec refs. |  |
| R005 | Implementations MUST remove the `.md` extension from spec refs. |  |
| R006 | Implementations MUST omit `index` from the ref when `index.md` represents the containing directory entrypoint. |  |
| R007 | Implementations MUST keep a non-index file stem as the final spec ID segment. |  |
| R008 | Implementations MUST normalize hyphenated path segments to underscore spec ID segments. |  |
| R009 | Implementations MUST enforce one-to-one path/ref mapping for new or migrated specs. | This includes keeping a changed canonical spec ref consistent with the placement path. |
| R010 | Implementations MUST treat a spec move or rename as a canonical spec ref change unless explicit compatibility design adds an alias or redirect. |  |
| R011 | Validators MUST report an error when a new or migrated spec's visible `id` does not match its path-derived canonical ref. |  |
| R012 | Validators MUST NOT rewrite spec IDs automatically. |  |
| R013 | Implementations MUST accept only `root`, `-`, or a canonical `spec:` ref as parent marker values. |  |
| R014 | Implementations MUST reject any other parent marker value. |  |

## Requirements: spec:product.design_records.traceability.semantic_ref

| id | requirement | notes |
|---|---|---|
| R015 | Implementations MUST treat `spec:` refs as path-derived document-level canonical refs. |  |
| R016 | Implementations MUST treat physical paths as repository locations, not canonical relation values. |  |
| R017 | Implementations MUST NOT treat section refs as active canonical refs before a visible-table contract exists. |  |

## Requirements: spec:product.design_records.traceability.artifact_refs

| id | requirement | notes |
|---|---|---|
| R018 | Implementations MUST treat `spec:` refs as canonical document-level identities for new and migrated specs. |  |
| R019 | Implementations MUST use the path-derived H1-adjacent `id` as the canonical spec ref. |  |
| R020 | Implementations MUST NOT register spec refs through hidden front matter. |  |

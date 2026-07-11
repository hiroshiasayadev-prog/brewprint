# USDM requirement: Physical path boundary

- **id**: `usdm:product.design_records.repository_layout_and_discovery.physical_path_boundary`
- **status**: draft
- **date**: 2026-07-09
- **kind**: requirement
- **parent**: `usdm:product.design_records.repository_layout_and_discovery`

## What this is

Requirements for the boundary between physical paths and canonical artifact identity.

## Requirements: Repository layout model
> source: spec:product.design_records.repository_layout

| id | requirement | notes |
|---|---|---|
| R001 | The implementation must treat a physical path as a repository location, not as canonical artifact identity. | A path says where an artifact is located. Identity comes from a public record ID or a path-derived `spec:` ref. |

## Requirements: Spec ID-as-ref
> source: spec:product.design_records.spec_format.spec_id_as_ref

| id | requirement | notes |
|---|---|---|
| R002 | The implementation must treat the canonical identity of a Specification record as the path-derived `spec:` ref under `records/spec/`. | `records/spec/` under an app namespace directory maps to `spec:<app>.`. Directory separators map to `.`, `.md` is removed, `index.md` is omitted, non-index file stems become final segments, and hyphens are normalized to underscores. |
| R003 | The implementation must treat a new or migrated Specification record as invalid when its visible `id` does not match its path-derived canonical `spec:` ref. | This requirement does not allow automatic rewriting by the implementation. ID-changing edits must be explicit authoring or migration actions. |

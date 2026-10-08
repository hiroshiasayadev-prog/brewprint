# USDM requirement: Record discovery paths

- **id**: `usdm:product.design_records.repository_layout_and_discovery.record_discovery_paths`
- **status**: draft
- **date**: 2026-07-13
- **kind**: requirement
- **parent**: `usdm:product.design_records.repository_layout_and_discovery`

## What this is

Requirements for kind-specific discovery path patterns for Design Records.

## Requirements: Record discovery paths
> source: spec:product.design_records.repository_layout.record_discovery_paths

| id | requirement | notes |
|---|---|---|
| R001 | The implementation must treat `records/spec/**/*.md` as the discovery path pattern for Specification records. | Specification records use topic tree placement, so discovery is recursive. |
| R002 | The implementation must discover sequential Design Record sources by artifact-kind directory, one physical domain subdirectory, and Markdown extension without filtering by file-name public-ID text. | The scope includes `adr/*/*.md`, `investigations/*/*.md`, `requirements/*/*.md`, `work-items/*/*.md`, and `tasks/*/*.md`. Identity and file-name conformance are evaluated after physical discovery. |
| R004 | A sequential source whose file-name public-ID prefix differs from its H1 public ID must remain in the discovered current-source corpus. | The mismatch is a nonidentity repository-conformance violation, not a discovery exclusion. |

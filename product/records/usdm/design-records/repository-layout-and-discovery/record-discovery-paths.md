# USDM requirement: Record discovery paths

- **id**: `usdm:product.design_records.repository_layout_and_discovery.record_discovery_paths`
- **status**: draft
- **date**: 2026-07-09
- **kind**: requirement
- **parent**: `usdm:product.design_records.repository_layout_and_discovery`

## What this is

Requirements for kind-specific discovery path patterns for Design Records.

## Requirements: Record discovery paths
> source: spec:product.design_records.repository_layout.record_discovery_paths

| id | requirement | notes |
|---|---|---|
| R001 | The implementation must treat `records/spec/**/*.md` as the discovery path pattern for Specification records. | Specification records use topic tree placement, so discovery is recursive. |
| R002 | The implementation must discover sequential Design Records artifacts with discovery path patterns based on kind directories and domain subdirectories. | The scope includes `adr/*/<record_prefix>ADR-*-*.md`, `investigations/*/<record_prefix>INV-*-*.md`, `requirements/*/<record_prefix>REQ-*-*.md`, `work-items/*/<record_prefix>WORK-*-*.md`, and `tasks/*/<record_prefix>TASK-*-*.md`. `<record_prefix>` is a filename-pattern placeholder. This row does not define ID grammar or prefix derivation rules. |
| R003 | The implementation must be able to discover existing flat ADR records through the compatibility discovery path pattern. | The compatibility pattern is `adr/<record_prefix>ADR-*.md`. The standard placement for new ADRs is the domain subdirectory pattern. |

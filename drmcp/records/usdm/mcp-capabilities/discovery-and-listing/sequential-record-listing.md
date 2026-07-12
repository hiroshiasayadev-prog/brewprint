# USDM requirement: Sequential record listing

- **id**: `usdm:drmcp.mcp_capabilities.discovery_and_listing.sequential_record_listing`
- **status**: draft
- **date**: 2026-07-11
- **kind**: requirement
- **parent**: `usdm:drmcp.mcp_capabilities.discovery_and_listing`

## What this is

Product requirements for listing current sequential records within an explicitly selected scope.

## Requirements: Sequential record listing
> source: literal

| id | requirement | notes |
|---|---|---|
| R001 | DRMCP must list current sequential records within a scope selected by app namespace, artifact kind, and domain namespace. |  |
| R002 | DRMCP must require the app namespace, artifact kind, and domain namespace selectors for a sequential record listing. | An omitted or invalid selector must not broaden the listing scope. |
| R003 | DRMCP must return the canonical ref of each listed record. | Additional compact projection fields may be defined by downstream Specifications. |
| R004 | DRMCP must accept only artifact kinds that use the sequential record structure for sequential record listing. |  |
| R005 | When a requested artifact kind is not supported for sequential record listing, DRMCP must report the invalid selector and expose the supported sequential artifact kinds. | Exact diagnostic representation belongs to downstream Specifications. |
| R006 | When a requested domain namespace has no discovered record source within the selected app namespace and artifact kind, DRMCP must report the invalid selector and expose the discovered domain namespaces available within that scope. |  |
| R007 | When a selected domain contains discovered record sources but no record eligible for listing, DRMCP must return a successful empty result. | Eligibility and addressability are defined by downstream model and operation Specifications. |
| R008 | DRMCP must allow the maximum number of returned records to be limited. |  |
| R009 | The default listing limit must be 30 records. | Accepted minimum and maximum values are defined by downstream Specifications. |
| R010 | DRMCP must support ascending and descending ordering by the sequence values declared for the selected artifact kind. | This requirement does not prescribe an internal sequence representation. |
| R011 | The default sequence ordering must be descending. |  |
| R012 | DRMCP must indicate when additional matching records exist beyond the applied listing limit. | The concrete pagination or continuation representation is not fixed here. |
| R013 | DRMCP must not expose physical source paths in normal sequential record listing results. |  |

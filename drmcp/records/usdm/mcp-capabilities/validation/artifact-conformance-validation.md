# USDM requirement: Artifact conformance validation

- **id**: `usdm:drmcp.mcp_capabilities.validation.artifact_conformance_validation`
- **status**: draft
- **date**: 2026-07-12
- **kind**: requirement
- **parent**: `usdm:drmcp.mcp_capabilities.validation`

## What this is

Product requirements for validating uniquely selectable current records against the applicable base and artifact-specific format contracts.

## Requirements: Applicable artifact contract
> source: literal

| id | requirement | notes |
|---|---|---|
| R001 | DRMCP must validate every selected current record against the base and artifact-specific Specifications applicable to its artifact kind. |  |
| R002 | DRMCP must validate current records as current artifacts and must not treat a legacy format as an alternative valid current format. |  |

## Requirements: Local artifact conformance
> source: literal

| id | requirement | notes |
|---|---|---|
| R004 | DRMCP must validate each selected current record against the applicable source-document shape, identity, record-structure, placement, metadata, and H2-section rules. |  |
| R005 | DRMCP must validate metadata field presence, multiplicity, value form, value type, and value format as declared by the applicable artifact Specifications. |  |
| R006 | DRMCP must validate H2 heading presence, multiplicity, conditional applicability, unlisted-heading policy, and referenced section-body formats as declared by the applicable artifact Specifications. |  |
| R007 | DRMCP must not repair, complete, normalize, or infer nonconforming source content when determining artifact conformance. |  |
| R008 | DRMCP must not treat absence of a recommended artifact element as a violation of a required artifact rule. | Exact finding classification belongs to downstream Specifications. |

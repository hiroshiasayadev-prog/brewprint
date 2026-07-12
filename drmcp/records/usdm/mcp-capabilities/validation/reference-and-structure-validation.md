# USDM requirement: Reference and structure validation

- **id**: `usdm:drmcp.mcp_capabilities.validation.reference_and_structure_validation`
- **status**: draft
- **date**: 2026-07-12
- **kind**: requirement
- **parent**: `usdm:drmcp.mcp_capabilities.validation`

## What this is

Product requirements for validating declared artifact references and relationships defined by current record structures.

## Requirements: Declared reference validation
> source: literal

| id | requirement | notes |
|---|---|---|
| R003 | DRMCP must validate every declared artifact reference of a selected current record against the relation rules applicable to its artifact kind and reference-bearing surface. | Text that is not declared as a reference-bearing surface is not resolved solely because it resembles an artifact ref. |
| R004 | DRMCP must validate current reference targets against the complete current record state across configured app namespaces, regardless of the selected validation scope. | A target outside the selected validation scope remains lookup-only unless independently selected. |
| R005 | DRMCP must make it identifiable when a declared reference has no uniquely addressable current target, identifies a conflicting current identity, identifies a disallowed artifact kind, or violates another applicable relation constraint. | Exact finding representation belongs to downstream Specifications. |
| R006 | DRMCP must not repair, complete, normalize, or infer a declared reference when validating its target or relation constraints. |  |

## Requirements: Record-structure validation
> source: literal

| id | requirement | notes |
|---|---|---|
| R007 | DRMCP must validate relationships between current records that are declared or derived by the applicable artifact record structure. |  |
| R008 | DRMCP must validate tree-record parent, child, placement, identity, and child-declaration consistency as required by the applicable artifact Specifications. |  |
| R009 | DRMCP must make structural conflicts such as inconsistent parentage, duplicate child declarations, or prohibited cycles identifiable when the applicable artifact Specifications prohibit them. |  |

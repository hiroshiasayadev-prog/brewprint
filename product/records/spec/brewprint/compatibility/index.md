# Overview: Brewprint compatibility

- **id**: `spec:product.brewprint.compatibility`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.brewprint`

## What this is

Owns Brewprint V01 compatibility, historical attribution, issued-ID retention, and migration state.

## Current contract

| compatibility area | owner |
|---|---|
| Existing artifact ownership and effective attribution | `spec:product.brewprint.compatibility.existing_artifacts`. |
| Legacy ID family acceptance and retention | `spec:product.brewprint.compatibility.legacy_id_compatibility`. |
| Compatibility-only V01 spec identity | `spec:product.brewprint.compatibility.legacy_id_compatibility`. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Existing artifact compatibility | Reference | `spec:product.brewprint.compatibility.existing_artifacts` | Historical PRODUCT ownership, effective attribution, and migration state. |
| Legacy ID compatibility | Reference | `spec:product.brewprint.compatibility.legacy_id_compatibility` | Accepted V01 ID families, retention policy, and compatibility-only spec identity. |

## Rules

- Keep compatibility under `brewprint/compatibility/`.
- Do not create a top-level PRODUCT compatibility area.
- Do not represent effective attribution as a new ID or alias.
- Keep generic artifact grammar and spec identity rules as pointers.

## Boundary

| content | owner |
|---|---|
| Generic artifact ID grammar | `spec:product.design_records.namespace_model.artifact_id_grammar`. |
| Generic spec ID-as-ref behavior | `spec:product.design_records.spec_format.spec_id_as_ref`. |
| Current Brewprint namespace registry | `spec:product.brewprint.namespaces`. |
| DRMCP UI, MCP, projection, or tool behavior | DRMCP app-local specifications. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.brewprint` | Parent Brewprint overview. |
| `spec:product.brewprint.namespaces` | Current Brewprint namespace profile. |

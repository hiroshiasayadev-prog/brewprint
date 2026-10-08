# Overview: Design Records CLI

- **id**: `spec:drcli.design_records_cli`
- **status**: draft
- **date**: 2026-10-08
- **parent**: root

## What this is

Entry point for the complete externally observable DRCLI Specification.
DRCLI exposes Design Records through a portable command-line interface while PRODUCT remains authoritative for Design Records semantics.

## Current contract

DRCLI selects one repository per invocation and discovers current Design Records beneath that repository.
Current-source construction automatically discovers candidate `records` roots and resolves app namespaces from canonical record identity evidence.
DRCLI requires no repository manifest, persistent repository registration, per-repository installation, or preconfigured app-namespace and records-root association.

The artifact Specifications project PRODUCT-owned source, identity, metadata, and structure rules for DRCLI consumption.
The current-record model owns source discovery, admission, addressability, and logical scopes.
The operation Specifications own discovery, retrieval, search, reference resolution, validation, guidance, and authoring semantics.
Diagnostics own transport-independent semantic operation errors and warnings.
CLI and runtime Specifications map those semantic operations to portable commands, help, body input, output, exit status, repository traversal, and retained operational state.

Valid CLI usage is discoverable from root, group, and command help.
A caller does not need MCP request schemas or MCP transport knowledge to invoke DRCLI correctly.

## Non-goals

- Redefining PRODUCT-owned Design Records semantics.
- Replacing, modifying, deprecating, or emulating DRMCP transport.
- Selecting implementation language, parser framework, storage engine, filesystem-walking technology, or other implementation architecture.
- Requiring an MCP server, daemon, repository registration, or one DRCLI installation per repository.

## Topic map

- Artifact format projections belong to `spec:drcli.design_records_cli.artifacts`.
- Current-source and current-record state belong to `spec:drcli.design_records_cli.current_record_model`.
- Semantic operation diagnostics belong to `spec:drcli.design_records_cli.diagnostics`.
- Transport-independent operation contracts belong to `spec:drcli.design_records_cli.operations`.
- Command, help, traversal, output, distribution, and retained-state behavior belong to `spec:drcli.design_records_cli.cli`.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Design record artifacts | Overview | `spec:drcli.design_records_cli.artifacts` | Defines DRCLI projections of supported PRODUCT-governed Design Record artifact contracts. |
| Current record model | Overview | `spec:drcli.design_records_cli.current_record_model` | Defines selected-repository current-source discovery, admission, addressability, and scopes. |
| Operation diagnostics | Index | `spec:drcli.design_records_cli.diagnostics` | Defines semantic operation diagnostic envelopes and canonical codes. |
| Operations | Index | `spec:drcli.design_records_cli.operations` | Navigates discovery, retrieval, search, reference resolution, validation, guidance, and authoring operations. |
| CLI and runtime | Index | `spec:drcli.design_records_cli.cli` | Defines commands, help, repository traversal, body input, output, exit status, distribution, and retained state. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.repository_layout` | PRODUCT authority for Design Records repository placement semantics. |
| `spec:product.design_records.authoring_standards` | PRODUCT authority for Design Records authoring semantics. |
| `spec:product.design_records.traceability.artifact_refs` | PRODUCT authority for canonical record kinds and reference forms. |

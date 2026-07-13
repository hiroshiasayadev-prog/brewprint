# Overview: Current record model

- **id**: `spec:drmcp.design_records_mcp.current_record_model`
- **status**: draft
- **date**: 2026-07-13
- **parent**: `spec:drmcp.design_records_mcp`

## What this is

Defines the DRMCP model for discovering current sources, admitting current records, resolving identity conflicts, and constructing current-record scopes.

## Current contract

The current record model uses this progression:

1. Explicit configuration establishes configured records roots.
2. Artifact Specifications derive artifact source roots and discovery corpora.
3. Discovered sources become artifact candidates only when sequential H1 identity or tree path-derived identity determination succeeds.
4. Candidate claims are aggregated by canonical identity.
5. One claim produces one uniquely addressable current record.
6. Multiple claims produce an identity conflict without a winner.
7. Current-record scopes expose sources, records, conflicts, and tree nodes according to operation needs.

The model includes only current Design Records.
Legacy archives do not participate in current discovery, admission, addressability, or scopes.

Sequential records use the H1 public ID as the sole source-internal identity authority.
H1-adjacent metadata and file names remain nonidentity conformance surfaces.
Tree records retain path-derived identity.

## Non-goals

- Concrete configuration file or command-line syntax.
- MCP tool names, request schemas, response schemas, or diagnostic payloads.
- Artifact-specific source, metadata, identity, or section formats.
- Search matching, listing order, retrieval projection, validation finding, or output-limit rules.
- Legacy lookup or migration behavior.
- Internal parser objects, caches, indexes, or storage architecture.

## Topic map

- Source configuration and physical discovery boundaries belong to current source corpus.
- Source interpretation and identity determination belong to artifact candidate and admission.
- Claim aggregation and exact addressability belong to identity conflict and addressability.
- App, artifact, domain, tree, and validation boundaries belong to current record scopes.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Current source corpus | Concept | `spec:drmcp.design_records_mcp.current_record_model.current_source_corpus` | Defines configured sources, artifact source roots, discovery corpora, and source provenance. |
| Artifact candidate and admission | Concept | `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission` | Defines candidate formation, structure-specific identity determination, admission failures, and nonidentity violations. |
| Identity conflict and addressability | Concept | `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability` | Defines identity claims, conflicts, uniquely addressable records, and selectable records. |
| Current record scopes | Concept | `spec:drmcp.design_records_mcp.current_record_model.current_record_scopes` | Defines current-state boundaries for apps, artifact kinds, domains, tree nodes, and subtrees. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts` | Defines supported artifact contracts consumed by this model. |
| `spec:product.design_records.repository_layout` | Product authority for the Design Records repository layout. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for canonical artifact references. |

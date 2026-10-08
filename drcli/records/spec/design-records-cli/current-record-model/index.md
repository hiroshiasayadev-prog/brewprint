# Overview: Current record model

- **id**: `spec:drcli.design_records_cli.current_record_model`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli`

## What this is

Defines the DRCLI model for discovering current sources, admitting current records, resolving identity conflicts, and constructing current-record scopes.

## Current contract

The current record model uses this progression:

1. Per-invocation repository selection and ignore handling establish traversal eligibility.
2. DRCLI discovers candidate directories named exactly `records` beneath the selected repository.
3. Canonical record identities provide namespace evidence without using directory or file names as app identity.
4. Each candidate records root resolves only when the evidence yields exactly one app namespace.
5. Artifact Specifications derive artifact source roots and discovery corpora beneath each resolved records root.
6. Discovered sources become artifact candidates only when sequential H1 identity or tree path-derived identity determination succeeds.
7. Candidate claims are aggregated by canonical identity.
8. One claim produces one uniquely addressable current record.
9. Multiple claims produce an identity conflict without a winner.
10. Current-record scopes expose sources, records, conflicts, and tree nodes according to operation needs.

DRCLI does not require persistent repository registration, a repository manifest, or preconfigured app-namespace and records-root associations.
Physical directory and file names never canonicalize the app namespace.

The model includes only current Design Records.
Legacy archives do not participate in current discovery, admission, addressability, or scopes.

Sequential records use the H1 public ID as the sole source-internal identity authority.
H1-adjacent metadata and file names remain nonidentity conformance surfaces.
Tree records retain path-derived identity.

## Non-goals

- CLI repository-selection syntax or persistent repository configuration formats.
- Transport-specific tool names, request schemas, response schemas, or diagnostic payloads.
- Artifact-specific source, metadata, identity, or section formats.
- Search matching, listing order, retrieval projection, validation finding, or output-limit rules.
- Legacy lookup or migration behavior.
- Internal parser objects, caches, indexes, or storage architecture.

## Topic map

- Records-root discovery, namespace resolution, and physical discovery boundaries belong to current source corpus.
- Source interpretation and identity determination belong to artifact candidate and admission.
- Claim aggregation and exact addressability belong to identity conflict and addressability.
- App, artifact, domain, tree, and validation boundaries belong to current record scopes.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Current source corpus | Concept | `spec:drcli.design_records_cli.current_record_model.current_source_corpus` | Defines automatic records-root discovery, namespace resolution, artifact source roots, discovery corpora, and source provenance. |
| Artifact candidate and admission | Concept | `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission` | Defines candidate formation, structure-specific identity determination, admission failures, and nonidentity violations. |
| Identity conflict and addressability | Concept | `spec:drcli.design_records_cli.current_record_model.identity_conflict_and_addressability` | Defines identity claims, conflicts, uniquely addressable records, and selectable records. |
| Current record scopes | Concept | `spec:drcli.design_records_cli.current_record_model.current_record_scopes` | Defines current-state boundaries for apps, artifact kinds, domains, tree nodes, and subtrees. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts` | Defines supported artifact contracts consumed by this model. |
| `spec:product.design_records.repository_layout` | Product authority for the Design Records repository layout. |
| `spec:product.design_records.traceability.artifact_refs` | Product authority for canonical artifact references. |

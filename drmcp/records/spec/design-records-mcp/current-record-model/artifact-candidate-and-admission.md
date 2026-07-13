# Concept: Artifact candidate and admission

- **id**: `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission`
- **status**: draft
- **date**: 2026-07-13
- **parent**: `spec:drmcp.design_records_mcp.current_record_model`
- **usdm_covers**:
  - usdm:drmcp.mcp_capabilities.discovery_and_listing.current_source_configuration#R006-R007
  - usdm:drmcp.mcp_capabilities.validation.validation_scope_and_subjects#R012-R013,#R019-R022
  - usdm:drmcp.mcp_capabilities.validation.artifact_conformance_validation#R009-R010
  - usdm:product.design_records.namespace_and_identity.workflow_artifact_identity#R015-R017
  - usdm:product.design_records.traceability_and_relations.resolve_and_validation#R003,#R015

## What this is

Defines how a discovered current source becomes an artifact candidate, an unadmitted source, or an admitted current record.

## Concept model

| concept | definition |
|---|---|
| Expected artifact context | The configured app namespace, corpus-selected artifact kind, and structure-dependent placement context for one discovered source. |
| Canonical current identity | The one current artifact identity determined by the applicable structure and artifact identity rules. |
| Current artifact candidate | A discovered source interpretable as the expected artifact and associated with one agreeing canonical current identity. |
| Unadmitted current source | A discovered source that cannot become a current artifact candidate. |
| Admitted current record | A candidate whose canonical identity has one claim after repository-wide claim aggregation. |
| Nonconforming admitted record | An admitted current record with one or more nonidentity artifact-contract violations. |

A syntax-level parsed source representation may exist in an implementation.
The parsed representation is not a formal current-record-model concept.

## Rules

### Expected artifact context

The source corpus selects the expected artifact kind.
DRMCP does not route a source into another artifact corpus from its content, filename, or declared identity.

For a sequential artifact, the physical domain directory selects the expected domain namespace.
An H1 public ID with another domain does not create or populate that other domain scope.

### Candidate formation

A discovered current source becomes a current artifact candidate only when both conditions hold:

1. DRMCP can parse the source sufficiently to apply the identity rule selected by its record structure.
2. That identity rule determines one canonical identity that agrees with the expected artifact context.

Candidate formation does not require the source to satisfy nonidentity metadata, heading, title, or file-name conformance rules.
Candidate formation does not repair, complete, normalize, or infer source content or identity values.

### Structure-specific identity determination

| record structure | canonical identity rule | candidate-blocking identity condition |
|---|---|---|
| `sequential` | The complete public ID parsed from the H1 prefix is the sole source-internal canonical identity authority. | The H1 public ID is missing or malformed, or its app namespace, artifact kind, or domain namespace disagrees with the expected artifact context. |
| `tree` | The path relative to the artifact source root determines the canonical identity. | The source cannot be interpreted as the expected artifact. |

Sequential identity determination does not inspect H1-adjacent metadata or the file name for another identity claim.
Metadata and file-name text do not repair, complete, normalize, infer, or replace the H1 public ID.
A prohibited metadata `id` or a file-name public-ID mismatch is evaluated only after H1 identity determination succeeds.

For a tree artifact, a visible `id` or similar declaration is not the identity authority.
A declared value that differs from the path-derived identity is a conformance violation.
The mismatch does not prevent candidate formation or admission.

### Unadmitted sources

Examples of unadmitted-source conditions include:

- unreadable source content;
- source content that cannot be interpreted under the expected artifact contract;
- unavailable canonical identity;
- missing or malformed sequential H1 public ID;
- sequential H1 app-namespace disagreement;
- sequential H1 artifact-kind disagreement;
- sequential H1 domain-placement disagreement.

An unadmitted source remains identifiable by repository-root-relative provenance during broad validation.
An unadmitted source cannot be selected for detailed validation by source path.

### Admission and nonidentity violations

Each current artifact candidate contributes one claim for its canonical identity.
Final admission follows the claim-aggregation rules in `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability`.

A nonidentity violation does not remove candidate or admitted-record status.
Nonidentity violations include prohibited sequential metadata `id`, file-name public-ID mismatch, other metadata violations, H2 violations, title violations, and other artifact-format violations that do not prevent canonical identity determination.

A nonconforming admitted record remains eligible for:

- sequential listing;
- exact retrieval;
- lexical search;
- exact-ref validation.

An operation may report an operation-specific limitation when a requested projection needs unavailable parsed data.
The limitation does not remove the record from the repository-wide addressable set.

## Boundary

| concern | owner |
|---|---|
| Expected artifact context | This Specification. |
| Structure-specific candidate boundary | This Specification. |
| Unadmitted-source classification boundary | This Specification. |
| Nonidentity violations and continued admission eligibility | This Specification. |
| Candidate claim aggregation | `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability`. |
| Tree-node existence independent of record admission | `spec:drmcp.design_records_mcp.current_record_model.current_record_scopes`. |
| Sequential H1 identity and tree path-derived identity rules | Applicable artifact Specifications and `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration`. |
| Detailed conformance findings | Validation Specifications. |
| Parser representation and implementation architecture | Implementation. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.current_record_model.current_source_corpus` | Supplies discovered current sources and expected artifact corpora. |
| `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability` | Aggregates candidate identity claims. |
| `spec:drmcp.design_records_mcp.current_record_model.current_record_scopes` | Retains tree nodes and broad-scope admission failures. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration` | Defines shared sequential and tree identity forms. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` | Defines shared record structures and source placement. |

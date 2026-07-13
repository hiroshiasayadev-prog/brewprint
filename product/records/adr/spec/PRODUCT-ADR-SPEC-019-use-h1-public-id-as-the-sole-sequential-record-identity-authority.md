# PRODUCT-ADR-SPEC-019: Use H1 public ID as the sole sequential record identity authority

- **status**: accepted
- **date**: 2026-07-13
- **depends_on**:
- **supersedes**:
- **migrated_to_spec**: 2026-07-13

## Context

Sequential Design Record kinds use complete public IDs as canonical references.

The source surfaces carrying public-ID text were inconsistent across artifact kinds:

- ADR and Investigation used the H1 public ID and file-name prefix;
- Requirement, Work Item, and Task additionally persisted metadata `id`;
- repository discovery patterns could filter sources by file-name public-ID text.

Treating several surfaces as identity-bearing creates independent disagreement states for H1, metadata, and file names.
It also forces candidate formation to reject or repair a source before one canonical identity can be selected.

The identity contract needs one source-internal authority while retaining repository-placement and file-name conformance checks.

## Decision

For every sequential Design Record, the complete public ID in the H1 prefix is the sole source-internal canonical identity authority.

The H1 form is:

```text
# <PUBLIC-ID>: <Title>
```

A sequential H1 public ID must:

- conform to the grammar for the corpus-selected artifact kind;
- use the configured app namespace;
- use the artifact kind selected by the physical source corpus;
- use the domain namespace represented by the physical domain directory.

A missing or malformed H1 public ID, or disagreement with any required physical context, prevents canonical identity determination and leaves the source unadmitted.

### Metadata boundary

ADR, Investigation, Requirement, Work Item, and Task metadata do not persist an `id` field.

Metadata does not supply, repair, complete, normalize, infer, or replace sequential identity.
When a legacy metadata `id` remains in an otherwise identifiable sequential record, the H1 identity remains effective and the field is reported as a prohibited metadata-conformance violation.

### File-name boundary

The file-name public-ID prefix remains required as a repository-conformance projection of the H1 public ID.

The file name is not an identity authority and is not a candidate-blocking identity-bearing surface.
A file-name prefix mismatch does not change the H1-derived identity or by itself prevent current-record addressability.
Validation reports the mismatch as a nonidentity artifact-contract violation.

Sequential discovery uses the artifact-kind directory, one physical domain directory, and Markdown extension without filtering by file-name public-ID text.

### Physical-context boundary

Sequential candidate formation compares the H1 public ID with:

- the explicitly configured app namespace;
- the artifact kind selected by the artifact source root;
- the physical domain directory directly beneath that root.

A sequential source outside this domain-directory structure is outside the current sequential source corpus.
Flat ADR compatibility is not retained as a current-record exception because it provides no physical domain directory for the required context agreement.

### Tree-artifact boundary

This decision does not change tree-artifact identity.

Specification and other tree artifacts retain path-derived canonical identity.
A visible tree-artifact metadata `id` may remain as a projection checked against that path-derived identity when its artifact contract requires it.

## Rationale

One authority removes redundant identity state and eliminates disagreement precedence between equally persisted surfaces.

H1 is the appropriate sequential authority because it is visible in the record body, already contains the complete public ID for every sequential kind, and remains independent of relation metadata.
Using the configured app, corpus-selected kind, and physical domain as agreement context prevents H1 content from reclassifying a source into another scope.

Keeping the file-name prefix as conformance preserves repository readability and authoring discipline without making record addressability depend on a duplicated projection.
Ignoring prohibited metadata `id` for identity allows existing records to remain addressable while still exposing migration work through validation.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Require H1, metadata `id`, and file-name public-ID prefix to agree before candidate formation. | It retains three identity-bearing surfaces and multiplies admission-failure states. |
| Use metadata `id` as the canonical authority for all sequential records. | It adds redundant identity metadata to ADR and Investigation and hides the primary identity below H1. |
| Use the file-name public-ID prefix as the canonical authority. | Renaming or malformed slugs would affect addressability, and discovery would exclude sources before conformance could report them. |
| Select a precedence winner only when identity surfaces disagree. | Lower-precedence surfaces would still duplicate identity and require repair or conflict semantics. |
| Infer missing H1 segments from metadata, file name, or physical context. | Inference would conceal malformed source identity and make canonical identity depend on repair behavior. |
| Retain flat ADR current compatibility by deriving its domain from H1. | It would create an ADR-specific exception that bypasses required physical domain agreement. |

## Consequences

- Requirement, Work Item, and Task metadata schemas prohibit `id`.
- New sequential-record writers project one resolved public ID into H1 and the file-name prefix but do not generate metadata `id`.
- Sequential discovery does not filter by file-name public-ID text.
- Missing or malformed H1 identity and H1 physical-context disagreement remain admission failures.
- File-name mismatch and prohibited metadata `id` remain addressable-record conformance violations.
- Existing records are not bulk-migrated by this decision; validation identifies remaining prohibited metadata `id` fields and file-name mismatches.
- Flat ADR sources must move into a physical domain directory before participating in the current sequential corpus.
- Tree-artifact path-derived identity and visible identity projections remain unchanged.
- Retrieval operations consume the resulting H1-derived canonical identity but are specified separately.

## Evidence

- `spec:product.design_records.namespace_model.artifact_id_grammar` defines the resulting sequential identity source and context agreement.
- `spec:product.design_records.authoring_standards.adr_authoring` and the other sequential authoring Specifications define the unified source and metadata shape.
- `spec:product.design_records.repository_layout.record_discovery_paths` defines file-name-independent physical discovery.
- `spec:product.design_records.traceability.resolve_and_validation` defines lookup-source and invalid-state boundaries.
- `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission` defines candidate and nonidentity-conformance boundaries.

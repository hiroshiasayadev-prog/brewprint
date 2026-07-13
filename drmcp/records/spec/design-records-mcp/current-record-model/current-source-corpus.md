# Concept: Current source corpus

- **id**: `spec:drmcp.design_records_mcp.current_record_model.current_source_corpus`
- **status**: draft
- **date**: 2026-07-13
- **parent**: `spec:drmcp.design_records_mcp.current_record_model`
- **usdm_covers**:
  - usdm:drmcp.mcp_capabilities.discovery_and_listing.current_source_configuration#R001-R005,R008-R011
  - usdm:drmcp.mcp_capabilities.discovery_and_listing.record_scope_discovery#R010-R011
  - usdm:drmcp.mcp_capabilities.validation.validation_scope_and_subjects#R016
  - usdm:product.design_records.repository_layout_and_discovery.record_discovery_paths#R002,#R004

## What this is

Defines the configured physical sources and artifact-bounded corpus from which DRMCP discovers current Design Record sources.

## Concept model

| concept | definition |
|---|---|
| Configured current source | One explicit association between an app namespace and one physical records root. |
| Configured source set | The complete set of configured current sources used to construct current-record state. |
| Configured records root | The physical boundary corresponding to one app namespace's Design Records `records/` directory. |
| Artifact source root | The configured records root joined with the artifact directory declared by one artifact Specification. |
| Source-discovery corpus | The union of existing artifact source roots beneath all configured records roots. |
| Discovered current source | A source found inside one artifact source root before candidate formation or admission. |
| Source provenance | The repository-root-relative physical path retained for validation and repair context. |

## Rules

### Explicit source configuration

Each configured current source contains exactly these associations:

| value | rule |
|---|---|
| app namespace | Supplied explicitly by configuration. |
| records root | Supplied explicitly as the physical discovery boundary for that app namespace. |

DRMCP does not auto-discover configured current sources.
DRMCP does not infer an app namespace from directory names, source content, filenames, or declared record identities.
The records root and its parent directories do not contribute canonical identity segments.

### Configuration uniqueness

The configured source set satisfies these invariants:

- One app namespace has at most one configured records root.
- One physical records root has at most one configured app namespace.
- One physical source is processed through at most one configured current source.

Concrete path canonicalization and overlap diagnostics belong to configuration Specifications.

### Artifact source roots

For each supported artifact kind, DRMCP derives the artifact source root as:

```text
<CONFIGURED_RECORDS_ROOT>/<ARTIFACT_DIRECTORY>/
```

`<ARTIFACT_DIRECTORY>` comes from the artifact's identity-and-structure Specification.
An artifact kind is physically available for an app when its derived artifact source root exists.
An existing empty artifact source root remains available.

### Corpus boundary

DRMCP discovers a source for an artifact kind only within that artifact kind's source root.
For a sequential artifact, a discovered current source is a Markdown file placed directly inside one physical domain directory beneath the artifact source root.
A Markdown file placed directly in the sequential artifact source root or below an additional nested directory is outside the current sequential source corpus.
Sequential source inclusion does not depend on public-ID text in the file name.
A file-name public-ID mismatch remains inside the corpus and is evaluated later as nonidentity conformance.
Source content, filename text, or declared identity does not reclassify a source into another artifact corpus.
A source outside every declared artifact source root remains outside the current-record corpus.
A source outside every configured records root remains outside the current-record corpus.

The current corpus excludes legacy archive sources.
Legacy source availability does not affect current source discovery.

### Source provenance

DRMCP retains one repository-root-relative path for each discovered current source.
The path identifies unadmitted sources and identity-conflict members during broad validation.
Normal discovery, listing, retrieval, navigation, and search results do not expose physical source paths.

## Boundary

| concern | owner |
|---|---|
| Configured app namespace and records-root association | This Specification. |
| Configured-source uniqueness invariants | This Specification. |
| Artifact source-root derivation | This Specification and the applicable artifact identity-and-structure Specification. |
| Source-discovery corpus boundary | This Specification. |
| Candidate formation and identity agreement | `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission`. |
| Identity conflict and unique addressability | `spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability`. |
| Scope availability and inclusion | `spec:drmcp.design_records_mcp.current_record_model.current_record_scopes`. |
| Configuration representation and diagnostics | Other DRMCP Specifications. |
| Artifact-specific source and identity formats | `spec:drmcp.design_records_mcp.artifacts`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.record_structure` | Defines artifact-directory declarations and artifact source-root derivation. |
| `spec:drmcp.design_records_mcp.current_record_model.artifact_candidate_and_admission` | Consumes discovered current sources. |
| `spec:drmcp.design_records_mcp.current_record_model.current_record_scopes` | Uses configured sources and artifact source roots to establish scopes. |
| `spec:product.design_records.repository_layout.record_discovery_paths` | Product authority for standard artifact source placement. |

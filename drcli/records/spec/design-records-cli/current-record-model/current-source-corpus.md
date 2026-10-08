# Concept: Current source corpus

- **id**: `spec:drcli.design_records_cli.current_record_model.current_source_corpus`
- **status**: draft
- **date**: 2026-10-08
- **parent**: `spec:drcli.design_records_cli.current_record_model`
- **usdm_covers**:
  - usdm:drmcp.mcp_capabilities.discovery_and_listing.record_scope_discovery#R010-R011
  - usdm:drmcp.mcp_capabilities.validation.validation_scope_and_subjects#R016
  - usdm:product.design_records.repository_layout_and_discovery.record_discovery_paths#R002,#R004

## What this is

Defines automatic records-root discovery, app-namespace resolution, and the artifact-bounded corpus from which DRCLI discovers current Design Record sources in one selected repository.

## Concept model

| concept | definition |
|---|---|
| Selected repository root | The per-invocation repository boundary supplied by the CLI repository-selection contract. |
| Candidate records root | One traversal-eligible directory named exactly `records` at or beneath the selected repository. |
| Namespace evidence | One app namespace syntactically extracted from a canonical record identity in an in-corpus source. |
| Resolved current source | One resolved association between exactly one app namespace and one candidate records root. |
| Artifact source root | The resolved records root joined with the artifact directory declared by one artifact Specification. |
| Source-discovery corpus | The union of traversal-eligible current artifact corpora beneath all resolved records roots. |
| Discovered current source | A traversal-eligible source inside one artifact source root before candidate formation or admission. |
| Source provenance | The repository-root-relative physical path retained for validation and repair context. |

## Rules

### Candidate records-root discovery

DRCLI discovers candidate records roots only beneath the selected repository.
A traversal-eligible directory is a candidate when its directory name is exactly `records`.
The parent directory name never supplies app-namespace identity.
After DRCLI discovers a candidate records root, nested `records` directories inside that root are not additional candidate roots.

Git-compatible ignore pruning applies before candidate discovery.
A fully ignored `records` directory is not discovered and produces no discovery diagnostic.
`--no-ignore` removes ignore pruning only and then applies the same discovery and resolution rules.

A selected repository with no discovered candidate records root produces a valid empty current state.
DRCLI requires no repository manifest, persistent registration, per-repository installation, or app-directory naming convention.

### Namespace evidence

DRCLI collects namespace evidence only from traversal-eligible sources inside PRODUCT-defined current artifact corpora beneath one candidate records root.

| source structure | namespace evidence |
|---|---|
| Sequential source | The app segment of a syntactically valid H1 public ID under the canonical sequential ID grammar. |
| Specification source | The app segment of a syntactically valid H1-adjacent `spec:<app>...` ID, mapped to the corresponding app-namespace token. |

Directory names and file names never contribute namespace evidence.
The physical records-root path never contributes namespace evidence.
Evidence extraction tests only the identity syntax needed to extract the app namespace.
Candidate formation, admission, identity agreement, and nonidentity conformance remain later phases.

A malformed or unadmitted source may contribute namespace evidence when its canonical identity syntax is extractable.
An in-corpus source that contributes no namespace evidence remains eligible for the resolved root's source corpus.

### Root resolution and uniqueness

A candidate records root resolves only when its collected evidence contains exactly one distinct app namespace.

| root state | result |
|---|---|
| Exactly one distinct namespace | Create one resolved current source for that app namespace and candidate records root. |
| Zero namespaces | Request-wide `configuration_failure`; the root is unresolvable. |
| More than one namespace | Request-wide `configuration_failure`; the root is ambiguous and DRCLI selects no winner. |
| Empty discovered `records/` root | Zero namespaces; request-wide `configuration_failure`. |
| Malformed-only root | Resolve only when syntactically extractable evidence yields exactly one namespace. |
| Fully ignored root under default traversal | Omitted from discovery and produces no discovery diagnostic. |

Distinct resolved roots may coexist only when they resolve to distinct app namespaces.
Two resolved roots for one app namespace produce request-wide `configuration_failure`.
DRCLI does not merge duplicate app roots and does not choose one root.

The `configuration_failure` diagnostic uses the existing context-free operation code.
Its message identifies the cause and each affected repository-relative records root.
For a multi-namespace root, the message also identifies the conflicting namespaces.
No partial normal result is returned after any current-source resolution failure.

### Artifact source roots

For each supported artifact kind, DRCLI derives the artifact source root as:

```text
<RESOLVED_RECORDS_ROOT>/<ARTIFACT_DIRECTORY>/
```

`<ARTIFACT_DIRECTORY>` comes from the artifact's identity-and-structure Specification.
An artifact kind is physically available for an app when its derived artifact source root exists.
An existing empty artifact source root remains available.

### Corpus boundary

After a records root resolves, DRCLI associates every traversal-eligible source in that root's PRODUCT-defined current artifact corpora with the resolved app namespace.
This association includes malformed and unadmitted sources that contributed no namespace evidence.

DRCLI discovers a source for an artifact kind only within that artifact kind's source root.
For a sequential artifact, a discovered current source is a Markdown file placed directly inside one physical domain directory beneath the artifact source root.
A Markdown file placed directly in the sequential artifact source root or below an additional nested directory is outside the current sequential source corpus.
Sequential source inclusion does not depend on public-ID text in the file name.
A file-name public-ID mismatch remains inside the corpus and is evaluated later as nonidentity conformance.
Source content, filename text, or declared identity does not reclassify a source into another artifact corpus.
A source outside every declared artifact source root remains outside the current-record corpus.
A source outside every resolved records root remains outside the current-record corpus.

The current corpus excludes legacy archive sources.
Legacy source availability does not affect current source discovery.

### Source provenance

DRCLI retains one repository-root-relative path for each discovered current source.
The path identifies unadmitted sources and identity-conflict members during broad validation.
Normal discovery, listing, retrieval, navigation, and search results do not expose physical source paths.

## Boundary

| concern | owner |
|---|---|
| Candidate records-root discovery and namespace resolution | This Specification. |
| Resolved app-namespace and records-root associations | This Specification. |
| Root-resolution and duplicate-app-root failure triggers | This Specification. |
| Repository selection and traversal eligibility | `spec:drcli.design_records_cli.cli.repository_selection_and_traversal`. |
| Artifact source-root derivation | This Specification and the applicable artifact identity-and-structure Specification. |
| Source-discovery corpus boundary | This Specification. |
| Candidate formation and identity agreement | `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission`. |
| Identity conflict and unique addressability | `spec:drcli.design_records_cli.current_record_model.identity_conflict_and_addressability`. |
| Scope availability and inclusion | `spec:drcli.design_records_cli.current_record_model.current_record_scopes`. |
| Operation diagnostic code and envelope | `spec:drcli.design_records_cli.diagnostics`. |
| Artifact-specific source and identity formats | `spec:drcli.design_records_cli.artifacts`. |

## Related specs

| ref | relation |
|---|---|
| `spec:drcli.design_records_cli.artifacts.base.definitions.record_structure` | Defines artifact-directory declarations and artifact source-root derivation. |
| `spec:drcli.design_records_cli.current_record_model.artifact_candidate_and_admission` | Consumes discovered current sources. |
| `spec:drcli.design_records_cli.current_record_model.current_record_scopes` | Uses resolved current sources and artifact source roots to establish scopes. |
| `spec:drcli.design_records_cli.cli.repository_selection_and_traversal` | Defines the selected repository and traversal eligibility consumed by records-root discovery. |
| `spec:drcli.design_records_cli.diagnostics.operation_diagnostic_catalog` | Defines the existing context-free `configuration_failure` code. |
| `spec:product.design_records.repository_layout.record_discovery_paths` | Product authority for standard artifact source placement. |

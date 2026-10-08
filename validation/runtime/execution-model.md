# Runtime execution model

## Status

Working contract for source parsing, artifact interpretation, candidate formation, operation consumption, and conformance validation.

## Processing boundary

The runtime preserves this pipeline:

```text
source discovery
  -> shared source-structure parse
  -> artifact facts and identity determination
  -> candidate formation
  -> repository-wide identity aggregation and admission
  -> listing, retrieval, navigation, or search
  -> artifact conformance validation
```

These stages must not be collapsed into one artifact-specific validated parser.

## Shared source parsing

`artifact.base.source.parse` is the shared Design Record source-parser boundary.
It extracts recoverable source structure, including explicit source-location fields in declared occurrence records when reliable structure can still be retained.

Likely outputs include:

- H1 occurrences;
- H1 prefix and title candidates;
- H1-adjacent metadata occurrences;
- H2 sections;
- Markdown tables;
- source locations;
- explicit unavailable projections.

Shared parsing does not decide artifact conformance.
It does not own field requiredness, status vocabulary, required H2 sets, reference resolution, or finding classification.

## Artifact facts

An artifact module selects and interprets values relevant to its artifact contract from the parsed source representation.

Artifact facts preserve invalid actual values and any source-location or evidence fields explicitly declared by their fact types where structure is recoverable.
For example, a parsed status value outside the artifact vocabulary remains available as a fact so validation can report the actual value.

Artifact facts are validation-independent and may support listing, retrieval, search, identity determination, or validation.

## Identity determination and candidate gate

Only canonical identity availability and agreement with expected physical context gate candidate formation.

### Sequential artifacts

Sequential candidate formation requires an artifact-specific identity interpretation that obtains one complete H1 public ID and confirms agreement with:

- configured app namespace;
- corpus-selected artifact kind;
- physical domain namespace.

Missing, malformed, or context-inconsistent sequential identity prevents candidate formation.

### Tree artifacts

Tree-node identity is path-derived.
Tree-node construction does not require full source parsing or validation.
A directory node remains available without `index.md`.
An unreadable non-index source may still establish its path-derived leaf node.

## Nonidentity defects

Nonidentity defects do not remove candidate or admitted-record status.
Examples include:

- invalid status value;
- missing nonidentity metadata;
- prohibited sequential metadata `id`;
- file-name projection mismatch;
- required H2 absence;
- title violation;
- unresolved persisted reference.

A nonconforming admitted record remains eligible for listing, retrieval, search, and exact validation.

## Admission and addressability

Each candidate contributes one canonical identity claim.
Repository-wide aggregation determines whether the identity is:

```text
unresolved
uniquely addressable
conflicted
```

A conflict does not select a winner.
Admission and addressability remain independent of nonidentity conformance findings.

## Listing and retrieval

Listing and retrieval do not run complete validation as an eligibility gate.

A projection distinguishes:

```text
invalid but structurally available
  return the parsed actual value

structurally unavailable
  return the operation-defined unavailable or null projection
```

Examples:

- an unrecognized but parsed status string remains returnable;
- a metadata block that cannot be projected reliably may produce `metadata: null`;
- missing H2 contract requirements do not prevent returning the H2 headings that actually exist.

## Reference execution path

A raw reference follows this generic route:

```text
exact raw string
  -> artifact.base.reference.parse(raw)
     -> SharedRoutingOutcome
        -> UnrouteableReference
        | RoutedBaseReference
          -> BaseReference union
             tree -> TreeBaseReference
             sequential -> SequentialBaseReference
  -> ArtifactModuleCatalog lookup
  -> module-bound narrowing to the matching variant type
  -> module.reference.parse(concrete_base_ref)
     -> ArtifactReferenceParseOutcome
        -> InvalidArtifactReference
        | opaque ArtifactReference
  -> module.reference.canonical_form(reference)
     or artifact.base.reference.resolve_current(reference)
```

An `UnrouteableReference` remains an expected typed parse outcome and does not proceed to module lookup.
A `RoutedBaseReference` carries one `BaseReference` that contains only enough structure for routing; it does not establish complete canonical grammar conformance.
The active `BaseReference` variant type and structure-specific discriminator select a compatible artifact module.
The module-bound invocation verifies the selected module structure, narrows the `BaseReference` to the matching `TreeBaseReference` or `SequentialBaseReference` variant, and lets the selected artifact module confirm the complete canonical grammar.
A successful sequential parse retains the `SequentialBaseReference` and adds a native `map<string, string>` containing artifact-specific named segment values. A successful tree parse retains only the `TreeBaseReference`, because that base already represents the complete canonical tree identity.
Both native values are exposed to design-facing YAML as the single opaque `ArtifactReference` type.
`InvalidArtifactReference` remains an expected typed parse outcome and does not proceed to canonical formatting or current-state resolution. Only the `ArtifactReference` variant proceeds.

## Validation execution

Artifact validation consumes the data needed by its accepted rules, including:

- parsed source structure;
- artifact facts;
- identity result;
- physical source context;
- repository current-record state;
- reference resolution state;
- artifact tree state;
- normalized relation edges.

Artifact validation recipes own:

- which rule applies;
- expected values or conditions;
- finding classification;
- finding code;
- contract ref;
- affected subject mapping.

Generic checks and algorithms do not own artifact meaning.

## Failure boundary

Expected record or reference conditions remain normal typed outcomes.

Configuration or execution failures prevent reliable operation completion and follow the operation failure path.
They are not converted into artifact validation findings.

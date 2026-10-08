# Reference parsing boundary

## Status

Evidence-backed working boundary derived from current Product authority and DRMCP operation Specifications on 2026-07-16.
The shared-routing and artifact-specific parsing responsibilities are established here.
The shared-routing and artifact-specific parse outcome declarations are fixed.

## Authority reviewed

| authority | responsibility consumed here |
|---|---|
| `spec:product.design_records.traceability.artifact_refs` | Current record kinds and canonical reference families. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Sequential canonical public-ID forms and sequence widths. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Tree-form `spec:` reference derivation and segment syntax. |
| `spec:drmcp.design_records_mcp.artifacts.base.definitions.identity_declaration` | Shared tree and sequential identity envelopes. |
| Artifact-specific identity-and-structure Specifications | Artifact kind, structure, and artifact-specific segment contract. |
| Exact retrieval and tree navigation Specifications | Malformed-selector placement and no-repair boundary. |
| Current-record addressability Specification | Separation of malformed, unresolved, and conflicted outcomes. |

## Canonical reference families

| family | current canonical form | structure discriminator | artifact-specific grammar |
|---|---|---|---|
| Specification | `spec:<app>(.<path_segment>)*` | lowercase record kind before `:`; currently `spec` | Path segment syntax and path-derived identity rules. |
| ADR | `<APP>-ADR-<DOMAIN>-<SEQUENCE>` | uppercase artifact-kind segment `ADR` | One three-digit sequence. |
| Investigation | `<APP>-INV-<DOMAIN>-<SEQUENCE>` | uppercase artifact-kind segment `INV` | One three-digit sequence. |
| Requirement | `<APP>-REQ-<DOMAIN>-<SEQUENCE>` | uppercase artifact-kind segment `REQ` | One three-digit sequence. |
| Work Item | `<APP>-WORK-<DOMAIN>-<SEQUENCE>` | uppercase artifact-kind segment `WORK` | One three-digit sequence. |
| Task | `<APP>-TASK-<DOMAIN>-<WORK_SEQUENCE>-<TASK_SEQUENCE>` | uppercase artifact-kind segment `TASK` | Three-digit inherited Work sequence followed by two-digit Task sequence. |

Bare grammar fragments such as `REQ-*` or `TASK-*`, physical paths, file names, and section selectors are not canonical record refs.

## Required processing boundary

Canonical record-ref handling has two parsing stages rather than one all-knowing base parser:

```text
exact supplied string
  -> artifact.base.reference.parse(raw)
     -> SharedRoutingOutcome
        -> UnrouteableReference
        | RoutedBaseReference
          -> BaseReference union
  -> ArtifactModuleCatalog lookup
  -> module-bound narrowing
  -> artifact-specific canonical parse
     -> ArtifactReferenceParseOutcome
        -> InvalidArtifactReference
        | opaque ArtifactReference
  -> current identity resolution
     -> unresolved
     | conflicted
     | uniquely addressable record
```

The runtime must not trim, repair, complete, normalize, infer, or fuzzy-match the supplied string at either parsing stage.
The exact supplied string remains available to the consuming operation for result projection and diagnostics.

## Shared routing parse

The shared parser owns only enough structure to select a compatible artifact module.
`artifact.base.reference.parse(raw: string) -> SharedRoutingOutcome` returns one expected domain outcome without trimming or repairing `raw`.

The outcome types are:

```yaml
RoutedBaseReference:
  record:
    reference: BaseReference

UnrouteableReference:
  record:
    raw: string

SharedRoutingOutcome:
  union:
    - RoutedBaseReference
    - UnrouteableReference
```

`UnrouteableReference.raw` preserves the exact supplied string. The variant itself means that the string matches neither the sequential nor tree routing envelope; no separate reason field is required.
`RoutedBaseReference` deliberately wraps `BaseReference` so routing success versus failure is handled before any tree-versus-sequential branching.

A successful shared parse places the existing `BaseReference` union in `RoutedBaseReference.reference`:

```text
tree       -> TreeBaseReference
sequential -> SequentialBaseReference
```

The shared parser may establish:

- whether the input has a routeable tree or sequential envelope;
- the structure-specific discriminator used by the artifact module catalog;
- the app namespace represented by the exact input;
- the domain namespace represented by a sequential input;
- ordered artifact-specific sequential segments whose meaning remains unconfirmed;
- ordered tree canonical path segments after the app namespace;
- the exact raw input.

A `BaseReference` is therefore routing evidence, not proof that the supplied string is a valid canonical record ref.
Its detail and consumers must not describe it as a confirmed canonical reference.

The shared parser does not own:

- artifact-specific sequential segment count;
- artifact-specific names for sequential tail segments;
- sequence width or allocation semantics;
- Task parent-sequence semantics;
- whether a tree path segment corresponds physically to a directory node, `index.md`, or a non-index Markdown file;
- current identity existence;
- identity conflicts;
- operation diagnostic classification.

## Artifact-specific canonical parse

The selected module's `reference.parse` capability owns complete canonical-form confirmation for its artifact kind.

For a sequential module, this includes:

- the registered artifact-kind literal;
- exact segment count and order;
- artifact-specific segment formats;
- sequence widths;
- any artifact-specific canonical grammar rule.

For a tree module, this includes:

- the registered record-kind literal;
- valid app and path segment syntax;
- the module's accepted tree reference contract.

Artifact-specific parsing returns:

```yaml
ArtifactReference:
  opaque: true

InvalidArtifactReference:
  record:
    raw: string

ArtifactReferenceParseOutcome:
  union:
    - ArtifactReference
    - InvalidArtifactReference
```

`ArtifactReference` represents confirmed canonical grammar and may proceed to canonical formatting or current-state resolution.
Its design-facing declaration remains opaque, while the native representation is:

```text
ArtifactReference
  = TreeArtifactReference
  | SequentialArtifactReference

TreeArtifactReference
  base: TreeBaseReference

SequentialArtifactReference
  base: SequentialBaseReference
  fields: map<string, string>
```

For sequential artifacts, `fields` assigns artifact-specific semantic names to the canonical tail segments preserved by `base.artifact_segments`. It contains string keys and string values only and is not exposed as a design-facing `dict<string, string>` value. A successful parser maps every tail segment to exactly one module-declared field in source order, with no missing or additional field.
For tree artifacts, `TreeBaseReference` already contains the complete canonical identity, so the confirmed reference retains only that base value. No empty field map or additional tree identity structure is created.

`InvalidArtifactReference.raw` preserves the exact routed string that failed the selected artifact module's canonical reference contract. No finer failure reason is required at this boundary.

## Outcome ownership

| condition | stage | runtime meaning | DRMCP operation projection |
|---|---|---|---|
| Input cannot produce routing evidence | shared routing parse | `UnrouteableReference` | `malformed_selector` or operation-specific invalid-selector classification. |
| Input routes to a module but violates that artifact's canonical grammar | artifact-specific canonical parse | `InvalidArtifactReference` | `malformed_selector` or operation-specific invalid-selector classification. |
| Confirmed canonical ref has no current identity claim | current-state resolution | unresolved reference | `unresolved_record`, `node_not_found`, or consuming-operation equivalent. |
| Confirmed canonical ref has multiple current identity claims | current-state resolution | conflicted reference | `conflicted_record` or consuming-operation equivalent. |
| Confirmed canonical ref has one current identity claim | current-state resolution | uniquely addressable current record | normal operation processing. |

`unresolved` and `conflicted` must not appear in a parse-outcome type.
They require repository current state and belong after canonical parsing.

## Representative cases

| supplied input | shared route | artifact parse | later state |
|---|---|---|---|
| `DRMCP-TASK-MCP-016-01` | sequential, discriminator `TASK` | valid Task canonical form | resolve against current identity claims. |
| `DRMCP-TASK-MCP-16-1` | sequential, discriminator `TASK` | malformed Task sequence widths | no resolution. |
| `DRMCP-REQ-MCP-033` | sequential, discriminator `REQ` | valid Requirement canonical form | resolve against current identity claims. |
| `REQ-*` | no valid current canonical envelope | not invoked | malformed selector. |
| `spec:product.design_records.traceability` | tree, discriminator `spec` | valid Specification canonical form | resolve or navigate current tree state. |
| `spec:Product.design-records` | tree-like envelope | malformed Specification segment syntax | no resolution. |
| `spec:product.design_records.traceability#Sources` | tree-like envelope | not a canonical record ref | malformed selector for exact-record operations. |
| `product/records/spec/design-records/index.md` | no canonical record-ref envelope | not invoked | malformed selector. |

## Working sequential namespace-token contract

The working runtime contract adopts this lexical grammar for namespace tokens embedded in sequential canonical IDs:

```text
sequential_namespace_token := [A-Z][A-Z0-9_]*
```

This grammar applies to both `<APP_NAMESPACE>` and `<DOMAIN_NAMESPACE>` in sequential IDs.
A namespace token therefore cannot contain `-`; the hyphen character is reserved as the sequential ID segment separator.

This decision does not change tree-form Specification refs.
The `spec:` family continues to use the lowercase segment grammar already defined by Product authority:

```text
spec_segment := [a-z0-9][a-z0-9_]*
```

Under the sequential token contract, shared routing may deterministically split the exact input on `-` and interpret the source-ordered tokens as:

```text
0: app namespace
1: artifact-kind discriminator
2: domain namespace
3..: artifact-specific segments preserved without canonical validation
```

Shared routing validates the app and domain token grammar and obtains the artifact-kind discriminator from token `1`.
It does not validate the artifact-specific segment count, width, role, or meaning.
Those checks remain owned by the selected artifact module's canonical parser.

For example:

```text
DRMCP-TASK-MCP-016-01

app namespace: DRMCP
artifact-kind discriminator: TASK
domain namespace: MCP
artifact-specific segments: [016, 01]
```

The corresponding transparent routing-evidence shape is:

```yaml
SequentialBaseReference:
  raw: string
  app_namespace: string
  artifact_kind: SequentialReferenceDiscriminator
  domain_namespace: string
  artifact_segments: list<string>
```

`artifact_segments` preserves token order. The selected artifact module assigns names and validates meaning. For Task, the module may interpret the two values as `work_sequence` and `task_sequence`; the shared parser does not expose those names.

The generic Product namespace authority does not yet contain this lexical rule.
This working decision must later be routed into the applicable Product authority before the runtime design becomes canonical.
The current Brewprint registry remains profile data rather than the source of this grammar.

## Tree routing-evidence shape

A tree-form base reference preserves this transparent shape:

```yaml
TreeBaseReference:
  raw: string
  record_kind: TreeReferenceDiscriminator
  app_namespace: string
  path_segments: list<string>
```

For:

```text
spec:product.design_records.traceability.artifact_refs
```

shared routing produces the logical evidence:

```text
record kind: spec
app namespace: product
path segments: [design_records, traceability, artifact_refs]
```

`path_segments` contains canonical identity segments after the app namespace. It does not divide the final segment into a physical file name or distinguish a directory node from a leaf file. Both distinctions require current tree or source-path information and belong outside exact-reference parsing.

Source-derived identity construction separately owns the physical mapping from directory names, `index.md`, and non-index Markdown file stems to canonical tree segments.

## Consequence for the next runtime decision

The structure-specific `BaseReference` fields, `SharedRoutingOutcome`, `ArtifactReferenceParseOutcome`, and native `ArtifactReference` representation are fixed for the working runtime.
Sequential confirmed references retain their base plus artifact-specific `map<string, string>` fields. Tree confirmed references retain only their base because it already represents the complete canonical identity.
Exhaustive union switching removes the need for parse-success predicates, extraction functions, or failure-reason enums at either parsing stage.

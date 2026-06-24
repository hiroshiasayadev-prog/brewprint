# Reference: Semantic ref

- **id**: `spec:product.concepts.traceability.semantic_ref`
- **status**: draft
- **date**: 2026-06-23
- **parent**: `spec:product.concepts.traceability`

## What this is

Defines the `spec:` semantic ref grammar, stability rules, and document/section ref model for the traceability MVP.

## Semantic ref definition

A semantic ref is an identifier that stably references the concept represented by a brewprint docs artifact.

A semantic ref is not a physical path.
It is not a Markdown heading.
It is not a directory layout.

When file rename, document split, document merge, or section move occurs, a semantic ref should be maintained as long as it refers to the same concept.

MVP examples of active semantic refs:

```text
spec:trace
spec:trace.semantic-ref
spec:trace.resolve-and-validation
spec:project-artifact-model
```

Semantic refs and record ID-as-refs are distinct. A record ID-as-ref is a complete public record ID, not a semantic prefix and not a bare kind form. Existing issued records retain legacy public IDs; new sequential records use `spec:product.concepts.namespace_model.artifact_id_grammar`. New and migrated specs use path-derived `spec:` refs, while legacy `SPEC-*` public IDs are compatibility-only. Per V01-ADR-088, `internal-design:` and `coverage:` are outside the MVP active scope and will be reconsidered with future requirements.

Per V01-ADR-087, investigation `source_refs` and recorded `follow_up_results` use artifact ID-as-refs or semantic refs as canonical references depending on the target. Physical paths are not used as canonical references.

## Semantic ref grammar

The MVP prefix-ref grammar is:

```text
<prefix>:<domain>[.<concept>[.<subconcept>...]]
```

MVP examples of active semantic refs:

```text
spec:trace
spec:trace.semantic-ref
spec:trace.coverage-mapping
spec:trace.resolve-and-validation
spec:project-artifact-model
spec:project-artifact-model.responsibilities
```

### Character rules

| constraint | rule |
|---|---|
| allowed characters | `a-z 0-9 - . :` |
| prefix | lowercase ASCII |
| domain / concept / subconcept | lowercase ASCII |
| word separator | hyphen `-` |
| namespace separator | dot `.` |
| prefix separator | colon `:` |
| whitespace | not used |
| slash `/` | not used |
| file extension | not included |
| physical path | not included |

Non-ASCII semantic refs are not permitted in the MVP.
Human-readable titles go in Markdown headings or metadata fields.

## Document-level ref and section-level ref

The semantic ref schema has both document-level and section-level forms.

A document-level ref points to the concept represented by an entire artifact.
A section-level ref points to the concept represented by a specific section within an artifact.

A root document ref takes the form `<prefix>:<domain>` and points to a concept set or the whole root document.
A dot-notation ref points to a nested document or section identity beneath it.
Both are canonical as active `spec:` semantic refs.

The MVP only uses `spec:` as an active semantic ref, covering both document-level and section-level. `internal-design:` and `coverage:` have their endpoint identity deferred to a future decision and are not resolved in the MVP.

Example:

```yaml
semantic_refs:
  - spec:trace
  - spec:trace.semantic-ref
  - spec:project-artifact-model
sections:
  spec:trace.semantic-ref.definition: Semantic ref definition
  spec:trace.semantic-ref-grammar: Semantic ref grammar
  spec:project-artifact-model.responsibilities: Artifact responsibility matrix
```

`semantic_refs` declares document-level refs.
`sections` declares the mapping between section-level refs and Markdown heading text.

Do not write `{#anchor}` directly on Markdown headings as canonical identity.
Section anchors are managed via the `sections` mapping in front matter.

## Stability rules

Semantic refs are treated as append-only.

- A once-issued semantic ref must not be reused for a different concept
- When a heading is renamed, the semantic ref is preserved
- When a section is moved, the semantic ref is preserved
- When a file is renamed, the semantic ref is preserved
- When a document is split, the existing semantic ref is left with the closest successor document/section
- When a section is split, the existing semantic ref is left with the closest successor section; a new semantic ref is issued for the new concept
- When sections are merged, multiple semantic refs may point to the same section

These rules separate physical layout changes from trace stability.

## Redirect and superseded mapping

When a semantic ref needs to be deleted or its meaning changed, the existing ref must not be reused for a different concept.

The MVP does not define a complete schema for redirect/superseded mapping.
However, the following is reserved for handling in subsequent specs:

```yaml
redirects:
  spec:old.ref: spec:new.ref
superseded:
  spec:old.ref:
    by: spec:new.ref
    reason: optional
```

Until redirect/superseded mapping is introduced, existing semantic ref entries must not be deleted and reassigned to different concepts in the MVP.
When a reference target is renamed, moved, or split, the old semantic ref entry remains with the closest successor document/section; new concepts receive new semantic refs.

## Out of scope

This file does not define:

- Complete list of active/reserved prefixes
- Semantic realization mapping schema
- Resolver request/response
- MCP tool contract
- Brewprint DSL YAML entity-level refs

These are covered in sibling spec files.

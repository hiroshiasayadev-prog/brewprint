# Overview: PuppyDSL source model

- **id**: `spec:puppydsl.language.source_model`
- **status**: draft
- **date**: 2026-07-25
- **parent**: `spec:puppydsl.language`

## What this is

Defines how one PuppyDSL application root yields a declaration source set. The source model owns file identification, declaration roots, source headers, and the designated entry declaration.

## Current contract

Only files whose names end with the exact lowercase suffix `.puppy.yaml` are PuppyDSL declaration sources.

| file name | treatment |
|---|---|
| `types.puppy.yaml` | PuppyDSL source. |
| `main.puppy.yaml` | PuppyDSL source and designated entry declaration file. |
| `types.yaml` | Ordinary application data. |
| `types.yml` | Ordinary application data. |
| `types.puppy.yml` | Ordinary application data. |

Source discovery recursively scans one selected application root. Generated output and temporary generation directories are excluded by the application integration contract, not by declaration semantics.

Each declaration file contains exactly one YAML document. The document root contains exactly one supported declaration-kind mapping.

```text
types
functions
contexts
events
```

A file groups declarations of one kind only. The mapping key beneath the declaration root is the declaration identity.

## Source envelope

Every declaration file begins with this raw-source envelope:

```yaml
# responsibility: Defines one coherent declaration set.
# excludes: Identifies adjacent concerns owned elsewhere.

functions:
  example.run:
    detail: Runs one example operation.
    signature: >-
      () -> void
    implementation:
      stub: native
```

| rule | requirement |
|---|---|
| First non-empty line | One `# responsibility:` comment with non-empty single-line text. |
| Second non-empty line | One `# excludes:` comment with non-empty single-line text. |
| Header separation | Exactly one empty line between the required comments and declaration root. |
| Additional comments | Allowed only after the required header. |
| YAML documents | Multi-document source is prohibited. |

The source envelope is validated before ordinary YAML declaration decoding. The comments are source authority and are not decoded declaration fields.

## Declaration detail

Every type, function, context, and event declaration contains a non-empty `detail` field.

`detail` states the declaration's semantic responsibility. `detail` must not merely repeat fields, signatures, or implementation representation.

## Entry declaration

The selected application root contains one `main.puppy.yaml` file. That file declares the reserved unqualified function ID `main` beneath the ordinary `functions` root.

No other source file may declare `main`. `main.puppy.yaml` introduces no separate declaration grammar.

## Rules

- Declaration files use lower-snake directory segments and file stems.
- One source file contains one declaration root.
- Duplicate declaration identities are program errors, even when they occur in different files.
- Physical source ownership does not create function namespace segments.
- Ordinary YAML files beneath the same root are not interpreted as declarations.

## Boundary

| owned here | owned elsewhere |
|---|---|
| `.puppy.yaml` identification. | Application-specific data file schemas. |
| Declaration root set. | Generator-specific manifest discovery. |
| Required raw-source header. | Diagnostic output format. |
| `main.puppy.yaml` entry role. | Runtime host startup and shutdown. |
| Declaration source grouping. | Application-specific directory ownership conventions. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Identifiers | Reference | `spec:puppydsl.language.source_model.identifiers` | Identifier grammar for declarations, namespaces, fields, parameters, and locals. |

## Related specs

| ref | relation |
|---|---|
| `spec:puppydsl.language` | Parent language overview. |
| `spec:puppydsl.language.functions` | Defines the reserved `main` function contract. |
| `spec:puppydsl.language.program_validity` | Defines source-set registration and duplicate rejection. |

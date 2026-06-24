# Reference: Repository implementation-flow preservation

- **id**: `spec:product.bpdsl.repository_implementation_flow`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product.bpdsl`

## What this is

Preservation-only staging for `dsl/`, `src/`, generated-source, and handwritten bootstrap layout material removed from the generic repository-layout contract.
It preserves existing meaning without making the generic Design Records repository layout own BPDSL implementation flow.

## Preservation status

| rule | status |
|---|---|
| Ownership claim | No PRODUCT canonical ownership claim. |
| Expected final owner | `bpdsl/records/spec/**` for BPDSL-internal content after migration review. |
| Integration claim | Historical; not adopted by PRODUCT. Retained as migration-review evidence only. Historical model differs from current app-local BPDSL contracts. |
| Migration review obligation | Review all retained statements during BPDSL migration or when an explicit integration requirement is accepted. |

## Preserved repository structure

The previous repository-layout model described an app namespace as:

```text
<app>/
  records/
  dsl/
  src/
```

| directory | preserved statement | status |
|---|---|---|
| `records/` | Stores app-owned design records, workflow artifacts, specifications, and authoring guidance. | Generic Design Records pointer. |
| `dsl/` | Stores BPDSL definitions used as an implementation source. | Preservation-only BPDSL staging. |
| `src/` | Stores runtime or implementation source. | Preservation-only implementation-flow staging. |

## Preserved DSL-to-source relationship

For an implementation-bearing app with an operational DSL pipeline, the previous long-term target was:

```text
dsl/ -> generated src/
```

The previous text allowed handwritten `src/` derived from current specifications and internal design while BPDSL could not express and generate the required implementation.
The previous text stated that handwritten bootstrap source does not replace `records/spec/` as the canonical authority for current design contracts.

## Preserved rules

- `dsl/` and `src/` were described as concern-based directories.
- The previous contract did not require every app namespace to contain `dsl/` and `src/`.
- `dsl/` was optional unless the app owns DSL definitions.
- `src/` was optional unless the app owns implementation code.
- The target implementation flow for an operational DSL pipeline was `dsl/` to generated `src/`.
- Direct spec-to-handwritten-`src/` implementation was a permitted bootstrap state while DSL support was insufficient.
- App-specific structure below `dsl/` or `src/` was outside the previous generic repository-layout concept.

## Deferred integration disposition

| preserved statement | status |
|---|---|
| Direct spec-to-handwritten-`src/` implementation as an integration relationship. | Previous description; not adopted by PRODUCT. Not represented by current app-local BPDSL contracts; current app-local BPDSL scope neither adopts nor evaluates this relationship. Historical context; retained as migration-review evidence only. |
| Internal design as a source for handwritten route behavior. | Previous description; not adopted by PRODUCT. No current app-local BPDSL owner for an internal design layer. Historical context; retained as migration-review evidence only. |

## Source mapping

| source file | source sections | disposition |
|---|---|---|
| `product/records/spec/concepts/repository-layout/index.md` | `## What this is`, `## Current contract`, `### DSL-to-source relationship`, `## Rules`, and `## Boundary` content about `dsl/`, `src/`, generated source, and handwritten bootstrap source. | Preserved here. |
| `product/records/spec/concepts/project-artifact-model/design-flow.md` | Related DSL-to-source and handwritten bootstrap statements. | Preserved here and in `spec:product.bpdsl.design_flow`. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.bpdsl` | Temporary staging parent. |
| `spec:product.design_records.repository_layout` | Generic records placement contract after extraction. |

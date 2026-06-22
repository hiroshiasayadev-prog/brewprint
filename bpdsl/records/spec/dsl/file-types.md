# Reference: YAML file types

- **id**: `spec:bpdsl.dsl.file_types`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.overview`

## What this is

Classification rules for brewprint YAML files. Defines the FileKind enumeration, the classification algorithm (`as:` / `nodes:` / filename), the `as:` value catalog, unsupported-file handling, and file-path normalization.

> Source: V01-ADR-030, V01-ADR-051, V01-ADR-052

## File kinds

Files under `yaml/` and `render_index.yaml` at the project root are classified by the loader into one of the following kinds.

| FileKind | contents | top-level key |
|---|---|---|
| `node` | Node definition file | has `nodes:` |
| `view` | View definition file | has `as:` |
| `render_index` | Project-root `render_index.yaml` | identified by filename |
| `unsupported` | Anything else | — |

> Source: V01-ADR-030, V01-ADR-043

## Classification algorithm

`*.yaml` / `*.yml` files are classified in the following order.

1. **Filename**: if the filename is `render_index.yaml` → `render_index`
2. **`as:` key**: if the top-level mapping has `as:` → `view` (value stored in `viewAs`)
3. **`nodes:` key**: if the top-level mapping has `nodes:` → `node`
4. Otherwise: `unsupported`

Only files with `.yaml` or `.yml` extensions are classified; other extensions are ignored by the loader.

Filenames other than `render_index.yaml` do not affect classification (arbitrary filenames are permitted per V01-ADR-002 / V01-ADR-030).

> Source: V01-ADR-030, V01-ADR-043

## `as:` value catalog

| `as:` value | view | source |
|---|---|---|
| `api_table` | API Table view | V01-ADR-028 |
| `sequence_diagram` | Sequence Diagram scenario | V01-ADR-032 |
| `er_diagram` | ER Diagram view (cross-module) | V01-ADR-039 |

When adding a new view kind, append a row to this table.

Wireframe has no file-level view — it is embedded as DSL within `state` node definitions (V01-ADR-029, V01-ADR-042). `as: wireframe` does not exist.

> Source: V01-ADR-030, V01-ADR-039, V01-ADR-029, V01-ADR-042

## Unsupported file handling

Files classified as `unsupported` are not silently skipped. The loader emits an `unsupported_file` warning diagnostic. Validation still succeeds (warnings do not affect exit code).

This prevents a `view` YAML with a missing `as:` key from disappearing silently — if it were silently skipped, the missing render output would go unnoticed.

The file is excluded from further resolve / render processing.

> Source: V01-ADR-051

## File path normalization

`source.File.Path` is stored as a project-root-relative, slash-normalized path.

- Base: project root (parent of `yaml/`)
- Separator: `/` (no OS-native separators)
- Examples: `yaml/auth/dag.yaml`, `render_index.yaml`

Absolute paths and OS-native backslash separators are not retained. The loader normalizes via `filepath.ToSlash` or equivalent.

This guarantees deterministic golden tests, reproducible render output, and consistent log / diagnostic display.

> Source: V01-ADR-052

# PRODUCT-TASK-SPEC-009-02: Format migration of all 9 target spec files

- **id**: PRODUCT-TASK-SPEC-009-02
- **status**: done
- **date**: 2026-06-22
- **work_item**: PRODUCT-WORK-SPEC-009
- **estimate**: 2d
- **depends_on**:
  - PRODUCT-TASK-SPEC-009-01
- **outputs**:
  - 20 spec files under `product/records/spec/concepts/` (migrated and split from 9 source files)

## Goal

Migrate all 9 source files to accepted spec format. Two standalone files (`project-artifact-model/index.md`, `namespace-model/index.md`) are split into topic-level sibling files per PRODUCT-WORK-SPEC-009 updated boundary. The 7 traceability files are migrated 1:1. Total output: 20 files.

## Output file table

### project-artifact-model/ — 1 source → 5 output

| output file | kind | id | parent | source content |
|---|---|---|---|---|
| `project-artifact-model/index.md` | `Overview` | `spec:product.concepts.project_artifact_model` | `root` | Purpose, artifact classes summary, source-of-truth table, navigation table, MVP scope / future extensions |
| `project-artifact-model/artifact-responsibility-matrix.md` | `Reference` | `spec:product.concepts.project_artifact_model.artifact_responsibility_matrix` | `spec:product.concepts.project_artifact_model` | `## Artifact responsibility matrix` table |
| `project-artifact-model/design-flow.md` | `Reference` | `spec:product.concepts.project_artifact_model.design_flow` | `spec:product.concepts.project_artifact_model` | Design artifact flow diagram + explanation |
| `project-artifact-model/change-and-investigation-flow.md` | `Reference` | `spec:product.concepts.project_artifact_model.change_and_investigation_flow` | `spec:product.concepts.project_artifact_model` | Change and investigation flow diagram + explanation |
| `project-artifact-model/traceability-boundary.md` | `Reference` | `spec:product.concepts.project_artifact_model.traceability_boundary` | `spec:product.concepts.project_artifact_model` | Traceability and tool boundary diagram + MVP table |

### namespace-model/ — 1 source → 6 output

| output file | kind | id | parent | source content |
|---|---|---|---|---|
| `namespace-model/index.md` | `Overview` | `spec:product.concepts.namespace_model` | `root` | Purpose, app/domain namespace concept intro, current/future placement note, `## Topics` |
| `namespace-model/app-namespaces.md` | `Reference` | `spec:product.concepts.namespace_model.app_namespaces` | `spec:product.concepts.namespace_model` | 3 app namespace definitions, each with architecture diagram and domain namespace table |
| `namespace-model/domain-catalog.md` | `Reference` | `spec:product.concepts.namespace_model.domain_catalog` | `spec:product.concepts.namespace_model` | Canonical domain namespace catalog + existing prefixes outside catalog |
| `namespace-model/subdomain-model.md` | `Reference` | `spec:product.concepts.namespace_model.subdomain_model` | `spec:product.concepts.namespace_model` | Subdomain grouping model, representation, write-time advisory, DRMCP example |
| `namespace-model/v2-grammar.md` | `Reference` | `spec:product.concepts.namespace_model.v2_grammar` | `spec:product.concepts.namespace_model` | v2 artifact ID grammar, sequence format, mapping rule from existing IDs |
| `namespace-model/existing-artifacts.md` | `Reference` | `spec:product.concepts.namespace_model.existing_artifacts` | `spec:product.concepts.namespace_model` | Existing artifact ownership table |

### traceability/ — 7 source → 7 output (1:1)

| output file | kind | id | parent | notes |
|---|---|---|---|---|
| `traceability/index.md` | `Overview` | `spec:product.concepts.traceability` | `root` | Resolved in PRODUCT-TASK-SPEC-004-01. `## Topics` → 6 sibling files. |
| `traceability/semantic-ref.md` | `Reference` | `spec:product.concepts.traceability.semantic_ref` | `spec:product.concepts.traceability` | |
| `traceability/artifact-refs.md` | `Reference` | `spec:product.concepts.traceability.artifact_refs` | `spec:product.concepts.traceability` | |
| `traceability/metadata-schema.md` | `Reference` | `spec:product.concepts.traceability.metadata_schema` | `spec:product.concepts.traceability` | Add drift guards to 2 hybrid sections (see below). |
| `traceability/coverage-mapping.md` | `Reference` | `spec:product.concepts.traceability.coverage_mapping` | `spec:product.concepts.traceability` | |
| `traceability/resolve-and-validation.md` | `Reference` | `spec:product.concepts.traceability.resolve_and_validation` | `spec:product.concepts.traceability` | Resolved in PRODUCT-TASK-SPEC-004-01. Add drift guards to 2 hybrid sections (see below). Preserve resolver output and MCP writer placeholder sections as-is. |
| `traceability/out-of-scope.md` | `Reference` | `spec:product.concepts.traceability.out_of_scope` | `spec:product.concepts.traceability` | |

## Migration steps (each file)

1. Remove YAML front matter block entirely.
2. Add H1 in `# <Kind>: <Title>` format.
3. Add H1-adjacent metadata (`id`, `status`, `date`, `parent`) immediately after H1.
4. Translate all H1, H2 titles, and table headers to English. Body prose may remain Japanese.
5. Add `## What this is` as the first body section.
6. Add required sections for the spec kind per `product/records/guides/spec-authoring.md`.
7. For Overview files: add `## Topics` table pointing to child files with correct `ref` values.
8. Remove stale `docs/...` path references. Update to current paths where a live target exists; remove if no current target exists.

## Drift guards

**`traceability/metadata-schema.md`:**

| section | drift guard |
|---|---|
| Workflow and investigation reference metadata | `> **Drift guard:** PRODUCT owns concrete bidirectional integrity rule statements. DRMCP owns parser, response, and diagnostic details — do not add those here.` |
| Validation responsibility | `> **Drift guard:** PRODUCT owns invalid conditions and scope-boundary clauses. DRMCP owns diagnostic category names, JSON shape, and tool response vocabulary — do not add those here.` |

**`traceability/resolve-and-validation.md`:**

| section | action |
|---|---|
| Resolve / resolver input / lookup sources / section anchor lookup | Add: `> **Drift guard:** PRODUCT owns canonical input semantics and tool behavior descriptions expressed in PRODUCT metadata. DRMCP may expose these as tool API — do not let DRMCP API vocabulary accumulate here.` |
| Duplicate detection / unresolved refs / declared relation integrity | Add: `> **Drift guard:** PRODUCT owns invalid conditions and scope boundary. DRMCP owns diagnostic category names, JSON response shape, and tool response vocabulary — do not let those accumulate here.` |
| Resolver output | Preserve as-is — DRMCP-owned pointer. |
| MCP writer contract placeholder | Preserve as-is — do not expand into PRODUCT-owned API schema. |

## Done condition

| item | done when |
|---|---|
| 20 files exist | All output files listed above are present. |
| YAML front matter removed | No YAML front matter block in any output file. |
| English titles | All H1, H2 titles, and table headers are English across all 20 files. |
| Topics tables correct | `project-artifact-model/index.md`, `namespace-model/index.md`, and `traceability/index.md` each have a `## Topics` table covering all their declared children with correct `ref` values. |
| parent chain clean | Every child file's `parent:` matches the `id:` of its declared parent. |
| drift guards added | 2 guards in `metadata-schema.md`, 2 guards in `resolve-and-validation.md`; resolver output and MCP writer sections preserved. |
| stale paths removed | No `docs/...` references remain in any output file. |
| no ownership change | No content relocated across namespace or ownership boundary. |

## Verification

- Cross-check `traceability/index.md` `## Topics` rows against the 6 sibling filenames.
- Cross-check `project-artifact-model/index.md` `## Topics` rows against the 4 child filenames.
- Cross-check `namespace-model/index.md` `## Topics` rows against the 5 child filenames.
- Confirm staged files in `product/records/old/` were not modified.

## Evidence

### Migration (2026-06-22)

18 output files written under `product/records/spec/concepts/`:

- `traceability/` — 7 files migrated 1:1 (index.md rewritten as Overview; 6 siblings rewritten as Reference)
- `project-artifact-model/` — 1 source → 5 files (index.md rewritten as Overview; 4 new Reference files created)
- `namespace-model/` — 1 source → 6 files (index.md rewritten as Overview; 5 new Reference files created)

Note: T02 task spec said "20 files" but the breakdown tables produce 18. Discrepancy traced to session context; actual file count per the breakdown tables is 18.

All body prose subsequently retranslated to English (original T02 instruction "body prose may remain Japanese" was incorrect per user correction). M1 table-header findings from Opus review applied in the same pass.

All files:
- YAML front matter removed
- H1 in `# <Kind>: <Title>` format (English)
- H1-adjacent metadata (`id`, `status`, `date`, `parent`)
- `## What this is` added as first body section
- All H2/H3 titles and table headers translated to English
- `## Topics` tables added to 3 Overview files
- Drift guards added to `metadata-schema.md` (2 sections) and `resolve-and-validation.md` (4 sections)
- Stale `docs/...` path hyperlinks removed; table cell paths updated to `records/` or `<namespace>/records/` pattern
- `## Current contract` added to 3 Overview files (required section)
- REFERENCE_NO_TABLE: `design-flow.md` and `change-and-investigation-flow.md` received summary tables after initial validation failed

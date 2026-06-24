# PRODUCT-TASK-SPEC-012-04: Split artifact model and repository layout semantics

- **id**: PRODUCT-TASK-SPEC-012-04
- **status**: done
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 2d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-02
- **outputs**:
  - Cleaned Design Records artifact-model content
  - Cleaned Design Records repository-layout content
  - Preservation-only BPDSL staging content

## Goal

Separate Design Records responsibilities and record placement from BPDSL and implementation-flow material.

## Work

Process this semantic batch:

- `project-artifact-model/index.md`
- `project-artifact-model/artifact-responsibility-matrix.md`
- `project-artifact-model/change-and-investigation-flow.md`
- `project-artifact-model/design-flow.md`
- `project-artifact-model/traceability-boundary.md`
- `repository-layout/index.md`
- `repository-layout/record-discovery-paths.md`

Keep Design Records artifact roles, workflow roles, and `records/` placement semantics.
Remove concrete DRMCP behavior from PRODUCT normative text.
Preserve unresolved DSL, source, render, and implementation-flow content under temporary PRODUCT BPDSL staging.
Do not redesign or validate BPDSL semantics.
Do not perform relocation or broad ref synchronization.

## Done condition

- Design Records artifact-model content owns only app-independent responsibilities.
- Repository-layout content separates record placement from BPDSL implementation concerns.
- BPDSL-related content is preserved without canonical ownership claims.
- Deferred integration statements are identified for extraction.
- The seven-file semantic batch is independently reviewable.

## Verification

- Compare each section against `PRODUCT-INV-SPEC-005` classifications.
- Confirm BPDSL staging preserves meaning without adding new contracts.
- Confirm DRMCP implementation details are pointers or handoff candidates only.
- Confirm no file relocation is mixed into the semantic review.

## Evidence

### Execution summary

| item | result |
|---|---|
| Seven source files processed | Yes. See section disposition map. |
| Generic artifact-model result | Retained app-independent ADR, investigation, requirement, work item, task, spec, workflow, source-of-truth, and traceability-boundary semantics. |
| Generic repository-layout result | Retained `<app>/records/`, kind-first directories, domain-scoped sequential record placement, spec topic-tree placement, records-only namespace validity, and physical path versus canonical identity boundary. |
| BPDSL staging result | Added preservation-only staging under `product/records/spec/bpdsl/`. |
| DRMCP result | Replaced concrete operational behavior with T08 handoff notes. |
| Deferred integration result | Preserved unresolved material as `T05 pending`. |
| Generic relocation | Not performed. Cleaned generic files remain in `product/records/spec/concepts/**`. |
| App-local specs changed | No changes under `drmcp/records/spec/**` or `bpdsl/records/spec/**`. |
| `v01/**` changed | No. |
| Broad ref synchronization | Not performed. Only local pointer edits required by T04 were made. |
| Review correction status | Final minor correction applied. T04 set to `done` after strict validation. |
| Objective contradictions found | None. |

### Files changed

| file | action |
|---|---|
| `product/records/spec/concepts/project-artifact-model/index.md` | Rewritten in place. |
| `product/records/spec/concepts/project-artifact-model/artifact-responsibility-matrix.md` | Rewritten in place. |
| `product/records/spec/concepts/project-artifact-model/change-and-investigation-flow.md` | Rewritten in place. |
| `product/records/spec/concepts/project-artifact-model/design-flow.md` | Removed after complete preservation under BPDSL staging. |
| `product/records/spec/concepts/project-artifact-model/traceability-boundary.md` | Rewritten in place. |
| `product/records/spec/concepts/repository-layout/index.md` | Rewritten in place. |
| `product/records/spec/concepts/repository-layout/record-discovery-paths.md` | Rewritten in place. |
| `product/records/spec/bpdsl/index.md` | Updated `## Topics` for actual staging files. |
| `product/records/spec/bpdsl/design-flow.md` | Created. |
| `product/records/spec/bpdsl/artifact-responsibilities.md` | Created. |
| `product/records/spec/bpdsl/repository-implementation-flow.md` | Created. |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-04-split-artifact-model-and-repository-layout.md` | Status and Evidence updated. |

### Section disposition map

| source file | section | disposition | preserved or retained location |
|---|---|---|---|
| `project-artifact-model/index.md` | `## What this is` | retain generic | Rewritten in source file. |
| `project-artifact-model/index.md` | `## Current contract` | retain generic | Rewritten in source file. |
| `project-artifact-model/index.md` | `## Boundary` | retain generic / stage BPDSL / T08 handoff | Rewritten as non-goals and ownership pointers. |
| `project-artifact-model/index.md` | `## Artifact classes` | split | Generic rows retained; BPDSL and implementation rows staged in `spec:product.bpdsl.design_flow` and `spec:product.bpdsl.artifact_responsibilities`. |
| `project-artifact-model/index.md` | `## Source of truth and documentation ownership` | split | Generic owners retained; DRMCP and BPDSL converted to pointers; authoring, migration, and implementation-follow-up statements preserved in `## Disposition of previous ownership statements`. |
| `project-artifact-model/index.md` | `## Disposition of previous ownership statements` | retain generic / pointer / T05 pending where in scope | Preserves authoring guide IDs, missing internal-design guide status, legacy M-series migration ownership, `docs/impl/**` follow-up, and authoring-guidance canonicalization context. Legacy M-series and `docs/impl/**` are not T05 pending. |
| `project-artifact-model/index.md` | `## Navigation` | split / T05 pending | Generic references retained; authoring entrypoints preserved as pointers; internal-design and future integration entrypoints marked T05 pending. |
| `project-artifact-model/index.md` | `## MVP scope and future extensions` | T05 pending | Preserved in source file as `T05 pending`. |
| `project-artifact-model/index.md` | `## Topics` | split | Generic topics retained; design-flow topic removed after staging. |
| `artifact-responsibility-matrix.md` | `## What this is` | retain generic | Rewritten in source file. |
| `artifact-responsibility-matrix.md` | `## Artifact responsibility matrix` | split | Generic Design Records rows retained; implementation rows staged in `spec:product.bpdsl.artifact_responsibilities`. |
| `change-and-investigation-flow.md` | `## What this is` | retain generic | Rewritten in source file. |
| `change-and-investigation-flow.md` | `## Change and investigation flow` | retain generic / T05 pending | Generic diagram retained; internal-design and YAML edges preserved as `T05 pending`. |
| `change-and-investigation-flow.md` | `## Artifact ownership in the flow` | retain generic | Rewritten in source file. |
| `change-and-investigation-flow.md` | `## Rules` | retain generic | Rewritten in source file. |
| `change-and-investigation-flow.md` | `## Sources` | retain generic | Retained in source file. |
| `design-flow.md` | all sections | stage BPDSL / remove after transfer | Preserved in `spec:product.bpdsl.design_flow` and `spec:product.bpdsl.repository_implementation_flow`; source file removed. |
| `traceability-boundary.md` | `## What this is` | retain generic | Rewritten in source file. |
| `traceability-boundary.md` | `## Traceability and tool boundary` | retain generic / T08 handoff | Generic boundary retained; DRMCP implementation behavior handed to T08. |
| `traceability-boundary.md` | `## MVP scope` | retain generic / T05 pending / T08 handoff | Active identity semantics retained; deferred endpoints preserved as `T05 pending`; DRMCP operation handed to T08. |
| `traceability-boundary.md` | `## Sources` | retain generic | Retained in source file. |
| `repository-layout/index.md` | `## What this is` | retain generic / stage BPDSL | Records placement retained; `dsl/` and `src/` ownership extracted. |
| `repository-layout/index.md` | `## Current contract` | retain generic / stage BPDSL | Records tree retained; `dsl/`, `src/`, generation, and bootstrap text staged. |
| `repository-layout/index.md` | `## Rules` | retain generic / stage BPDSL | Record placement retained; implementation-flow rules staged. |
| `repository-layout/index.md` | `## Boundary` | retain generic / stage BPDSL | Rewritten as generic ownership table and staging note. |
| `repository-layout/index.md` | `## Topics` | retain generic | Record discovery topic retained. |
| `repository-layout/index.md` | `## Related specs` | retain generic / local pointer update | Design-flow pointer removed; BPDSL staging pointer added. |
| `record-discovery-paths.md` | `## What this is` | retain generic / T08 handoff | Path-pattern scope retained; DRMCP behavior excluded. |
| `record-discovery-paths.md` | `## Current contract` | retain generic / T08 handoff | Generic placeholder retained; namespace-prefix derivation handed to T08. |

### Retained generic roles and rules

| category | retained |
|---|---|
| Artifact roles | ADR, investigation, requirement, work item, task, and spec. |
| Workflow roles | Investigation to ADR/REQ/SPEC/WORK, ADR to SPEC, REQ to WORK, WORK to TASK, WORK to SPEC. |
| Traceability boundary | Traceability is not an independent design authority. PRODUCT owns generic identity/reference/validation semantics. Tools may implement those semantics. |
| Repository placement | `<app>/records/`, kind-first directories, domain subdirectories, `records/spec/` topic tree, `guides/`, records-only app namespaces. |
| Identity boundary | Physical paths are locations. Public IDs and path-derived `spec:` refs are canonical identities. |

### BPDSL staging file map

| staging file | source sections preserved |
|---|---|
| `product/records/spec/bpdsl/design-flow.md` | `project-artifact-model/design-flow.md` all sections, old related-spec meanings, old historical source references, and `project-artifact-model/index.md` design and implementation artifact rows. |
| `product/records/spec/bpdsl/artifact-responsibilities.md` | `artifact-responsibility-matrix.md` implementation-flow rows and `project-artifact-model/index.md` design and implementation artifact rows. External relation material is pointer-only to the primary T05 location. |
| `product/records/spec/bpdsl/repository-implementation-flow.md` | `repository-layout/index.md` `dsl/`, `src/`, generated-source, and handwritten bootstrap statements; related `design-flow.md` bootstrap statements. |

### T08 handoff items

| source | handoff item |
|---|---|
| `traceability-boundary.md` | Design Records MCP indexing behavior. |
| `traceability-boundary.md` | Design Records MCP resolving behavior. |
| `traceability-boundary.md` | Design Records MCP validation behavior. |
| `traceability-boundary.md` | Concrete DRMCP request, response, diagnostic, parser, persistence, UI, or tool behavior. |
| `record-discovery-paths.md` | `<namespace_prefix>` derivation from `records_root`. |
| `record-discovery-paths.md` | DRMCP-specific index inclusion conditions. |
| `record-discovery-paths.md` | DRMCP schema provenance for discovery path-pattern content. |

### T05 pending items

| item | preserved location |
|---|---|
| `internal-design:` endpoint. | Primary: `project-artifact-model/index.md`. Pointers: `artifact-responsibility-matrix.md`, `traceability-boundary.md`. |
| Spec/internal-design realization relation. | Primary: `project-artifact-model/index.md`. Context: `bpdsl/design-flow.md`. |
| YAML semantic ref and YAML relation ownership. | `project-artifact-model/index.md`, `change-and-investigation-flow.md`. |
| `maps_to` and `covers` relations. | `project-artifact-model/index.md`. |
| External relation or assurance artifacts. | Primary: `project-artifact-model/index.md`. Pointers: `artifact-responsibility-matrix.md`, `bpdsl/artifact-responsibilities.md`, `traceability-boundary.md`. |
| Mapping groups. | `project-artifact-model/index.md`. |
| Fixture or golden traceability. | `project-artifact-model/index.md`. |
| Orphan diagnostics. | `project-artifact-model/index.md`, `traceability-boundary.md`. |
| Task-status-derived progress projection. | `project-artifact-model/index.md`, `traceability-boundary.md`. |
| Workflow traversal queries. | `project-artifact-model/index.md`, `traceability-boundary.md`. |
| Task dependency cycle or execution order projection. | `project-artifact-model/index.md`. |
| MCP writer tools. | `project-artifact-model/index.md`. |
| Design Records-to-BPDSL constraint statements. | `bpdsl/design-flow.md`. |
| Direct spec-to-handwritten-`src/` integration relation. | `bpdsl/repository-implementation-flow.md`. |
| Internal-design authoring guide missing in phase 1. | `project-artifact-model/index.md`. |

### Pointer-retained and historical items

| item | disposition |
|---|---|
| ADR authoring guide ID `adr-authoring`. | Pointer retained in `project-artifact-model/index.md`; not T05 pending. |
| Artifact-boundary guide ID `artifact-boundary`. | Pointer retained in `project-artifact-model/index.md`; not T05 pending. |
| Investigation, requirement, work-item, and task authoring guide IDs. | Pointer retained in `project-artifact-model/index.md`; not T05 pending. |
| Legacy M-series migration ownership. | Historical ownership pointer retained in `project-artifact-model/index.md`; not T05 pending. |
| `docs/impl/**` implementation follow-up. | Non-normative implementation-follow-up pointer retained in `project-artifact-model/index.md`; not T05 pending. |
| Authoring-guidance canonicalization context. | Historical authoring-guidance pointer retained in `project-artifact-model/index.md`; not T05 pending. |

### Removed-source transfer proof

| removed source | proof |
|---|---|
| `product/records/spec/concepts/project-artifact-model/design-flow.md` | Original `## What this is`, `## Design artifact flow`, `## Source of truth roles`, and `## Rules` are preserved in `product/records/spec/bpdsl/design-flow.md`. The old `spec:product.concepts.repository_layout` relation is preserved as historical preservation context and through `product/records/spec/bpdsl/repository-implementation-flow.md`. The old `spec:product.brewprint.layout` relation is preserved as historical context in `product/records/spec/bpdsl/design-flow.md`. Original V01 ADR and investigation references are preserved in `product/records/spec/bpdsl/design-flow.md` `## Preservation context`. Generic current-spec authority is retained as a pointer in generic artifact-model text. Integration claims are labeled `T05 pending`. |

### Validation

Command:

```powershell
python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color
```

Exit code: `0`

Output:

```text
[strict]  All 49 file(s) OK.
```

### Scoped git status evidence

Command:

```powershell
git status --short -- product/records/spec
```

Exit code: `0`

Output:

```text
 M product/records/spec/brewprint/index.md
 M product/records/spec/concepts/namespace-model/app-namespaces.md
 D product/records/spec/concepts/namespace-model/domain-catalog.md
 M product/records/spec/concepts/namespace-model/existing-artifacts.md
 M product/records/spec/concepts/namespace-model/index.md
 D product/records/spec/concepts/namespace-model/legacy-id-compatibility.md
 M product/records/spec/concepts/namespace-model/subdomain-model.md
 M product/records/spec/concepts/project-artifact-model/artifact-responsibility-matrix.md
 M product/records/spec/concepts/project-artifact-model/change-and-investigation-flow.md
 D product/records/spec/concepts/project-artifact-model/design-flow.md
 M product/records/spec/concepts/project-artifact-model/index.md
 M product/records/spec/concepts/project-artifact-model/traceability-boundary.md
 M product/records/spec/concepts/repository-layout/index.md
 M product/records/spec/concepts/repository-layout/record-discovery-paths.md
?? product/records/spec/bpdsl/
?? product/records/spec/brewprint/compatibility/
?? product/records/spec/brewprint/namespaces/
?? product/records/spec/design-records/
?? product/records/spec/index.md
```

Command:

```powershell
git status --short -- drmcp/records/spec bpdsl/records/spec v01
```

Exit code: `0`

Output:

```text
```

### Diff and leakage checks

| command | exit code | result |
|---|---:|---|
| `git diff --name-status -- product/records/spec` | 0 | Shows unrelated earlier namespace/Brewprint changes plus T04 changes. T04 changes are the seven source files and `product/records/spec/bpdsl/`. |
| `git diff -- product/records/spec/concepts/project-artifact-model` | 0 | Shows generic artifact-model rewrite and deletion of `design-flow.md` after transfer. |
| `git diff -- product/records/spec/concepts/repository-layout` | 0 | Shows records-only generic layout rewrite and abstract record-prefix path patterns. |
| `git diff -- product/records/spec/bpdsl` | 0 | No output because the staging directory is untracked. |
| `git diff --no-index -- NUL product/records/spec/bpdsl/index.md` | 1 | New-file diff inspected. |
| `git diff --no-index -- NUL product/records/spec/bpdsl/design-flow.md` | 1 | New-file diff inspected. |
| `git diff --no-index -- NUL product/records/spec/bpdsl/artifact-responsibilities.md` | 1 | New-file diff inspected. |
| `git diff --no-index -- NUL product/records/spec/bpdsl/repository-implementation-flow.md` | 1 | New-file diff inspected. |

Leakage command:

```powershell
rg -n "DSL|dsl/|src/|render|generated source|target implementation|Design Records MCP|namespace_prefix|internal-design:|yaml:|maps_to|covers|coverage|MCP writer" `
  product/records/spec/concepts/project-artifact-model `
  product/records/spec/concepts/repository-layout `
  product/records/spec/bpdsl
```

Exit code: `0`

Interpretation:

| match class | result |
|---|---|
| Generic artifact-model BPDSL terms | Only non-goal, staging pointer, source citation, or T05 pending usage. |
| Generic repository-layout `dsl/` / `src/` terms | Only boundary exclusion or staging map usage. |
| Generic traceability DRMCP terms | Only T08 handoff usage. |
| `namespace_prefix` | Only T08 handoff usage in `record-discovery-paths.md`. |
| `internal-design:`, `yaml:`, `maps_to`, `covers`, `coverage`, `MCP writer` | Only T05 pending, source citation, or staging usage. |
| BPDSL staging terms | Expected preservation-only content. |

No broad ref synchronization command:

```powershell
git diff --word-diff=porcelain -- product/records/spec |
  Select-String "spec:product\.concepts"
```

Exit code: `0`

Interpretation:

The output includes unrelated pre-existing namespace-model changes from earlier restructuring work.
The T04-scoped variant:

```powershell
git diff --word-diff=porcelain -- product/records/spec/concepts/project-artifact-model product/records/spec/concepts/repository-layout product/records/spec/bpdsl |
  Select-String "spec:product\.concepts"
```

Exit code: `0`

The T04-scoped output shows only local pointer edits required by the semantic split.
It removes the obsolete `spec:product.concepts.project_artifact_model.design_flow` topic and adds local parent or related-spec pointers.

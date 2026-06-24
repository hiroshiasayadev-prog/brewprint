# PRODUCT-TASK-SPEC-012-09: Finalize temporary BPDSL staging

- **id**: PRODUCT-TASK-SPEC-012-09
- **status**: done
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-07
- **outputs**:
  - `product/records/spec/bpdsl/index.md`
  - Preservation-only BPDSL staging files defined by the T01 manifest

## Goal

Finalize temporary PRODUCT-side BPDSL staging without adopting or redesigning BPDSL semantics.

## Work

- Place preserved BPDSL-related material under the temporary staging area.
- Keep only existing DSL, source, render, implementation-flow, or BPDSL artifact descriptions required for preservation.
- Remove accidental Design Records ownership claims.
- Confirm the staging overview defines the accepted review boundary and exit obligation.
- Do not normalize `bpdsl/records/spec/**`.
- Do not define new schema, resolver, render, generation, runtime, MCP, or integration contracts.

## Done condition

- Every staged item is listed in the T01 manifest.
- Staged content preserves existing meaning without semantic expansion.
- The staging area makes no canonical ownership claim.
- The BPDSL migration trigger and four-way exit classification remain explicit.
- No unrelated BPDSL specification is added.

## Verification

- Compare staged content with the source sections recorded by T01 and `PRODUCT-INV-SPEC-005`.
- Confirm `bpdsl/records/spec/**` remains unchanged unless a separate accepted handoff requires a pointer correction.
- Confirm the staging area is not referenced as canonical BPDSL authority.
- Confirm no Design Records-to-BPDSL integration rule is introduced.

## Evidence

### 1. Execution summary

| item | result |
|---|---|
| Final status | `done` |
| Files changed | `product/records/spec/bpdsl/index.md`, `product/records/spec/bpdsl/design-flow.md`, `product/records/spec/bpdsl/artifact-responsibilities.md`, `product/records/spec/bpdsl/repository-implementation-flow.md`, `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-09-finalize-temporary-bpdsl-staging.md` |
| Files deleted | None |
| Files created | None |
| App-local BPDSL specs edited | No. `bpdsl/records/spec/**` is unchanged. |
| New integration contract created | No. |

### 2. Complete disposition map

#### T03 handoff concerns (from `app-namespaces.md ## BPDSL`, dispositioned by T03 pointer)

| concern | source evidence | current staging location | app-local overlap or owner | final T09 disposition | file update, if any | rationale |
|---|---|---|---|---|---|---|
| BPDSL type behavior | T03 handoff note in `spec:product.brewprint.namespaces.app_namespaces` Boundary | None (T03 replaced with pointer) | `spec:bpdsl.dsl.type_ref` — active app-local contract | `replace_with_pointer` | None; app-local owner identified in T09 Evidence. Existing PRODUCT text provides an owner-level pointer; exact canonical cross-owner ref synchronization remains T10 scope. | Active app-local owner exists; no staging needed |
| Identity / name resolution | T03 handoff note in `spec:product.brewprint.namespaces.app_namespaces` Boundary | None (T03 replaced with pointer) | `spec:bpdsl.dsl.naming` — active app-local contract | `replace_with_pointer` | None; app-local owner identified in T09 Evidence. Existing PRODUCT text provides an owner-level pointer; exact canonical cross-owner ref synchronization remains T10 scope. | Active app-local owner exists; no staging needed |
| Render behavior | T03 handoff note in `spec:product.brewprint.namespaces.app_namespaces` Boundary | None (T03 replaced with pointer) | `spec:bpdsl.views.overview` and children — active app-local contract | `replace_with_pointer` | None; app-local owner identified in T09 Evidence. Existing PRODUCT text provides an owner-level pointer; exact canonical cross-owner ref synchronization remains T10 scope. | Active app-local owner exists; no staging needed |
| YAML loading / file classification | T03 handoff note in `spec:product.brewprint.namespaces.app_namespaces` Boundary | None (T03 replaced with pointer) | `spec:bpdsl.dsl.file_types` — active app-local contract | `replace_with_pointer` | None; app-local owner identified in T09 Evidence. Existing PRODUCT text provides an owner-level pointer; exact canonical cross-owner ref synchronization remains T10 scope. | Active app-local owner exists; no staging needed |
| MCP behavior | T03 handoff note in `spec:product.brewprint.namespaces.app_namespaces` Boundary | None (T03 replaced with pointer) | `spec:bpdsl.mcp.overview` and children — active app-local contract | `replace_with_pointer` | None; app-local owner identified in T09 Evidence. Existing PRODUCT text provides an owner-level pointer; exact canonical cross-owner ref synchronization remains T10 scope. | Active app-local owner exists; no staging needed |
| Self-hosting behavior | T03 handoff note in `spec:product.brewprint.namespaces.app_namespaces` Boundary | None (T03 replaced with pointer) | No current app-local BPDSL owner identified. Current BPDSL DSL scope explicitly excludes code generation. | `delete_with_rationale` | None; no staged content to remove | No accepted owner; not a current app-local contract; term does not correspond to any existing BPDSL spec section |

#### T04 / T05 staging concerns (from the four staging files)

| concern | source evidence | current staging location | app-local overlap or owner | final T09 disposition | file update, if any | rationale |
|---|---|---|---|---|---|---|
| `<app>/dsl/` as target implementation source | T04 `design-flow.md` § Preserved source-of-truth roles and §Preserved rules | `spec:product.bpdsl.design_flow` | None. Current app-local BPDSL uses `yaml/` not `dsl/`. `spec:bpdsl.dsl.project_layout` defines `yaml/` as source area. | `retain_preservation_only` | Labeled as previous description differing from current app-local model | Historical target not reflected in current app-local contracts; needed for migration review |
| `<app>/src/` as generated or bootstrap source | T04 `design-flow.md` § Preserved source-of-truth roles and § Preserved rules | `spec:product.bpdsl.design_flow` | None. Current BPDSL scope: code generation out of scope; no `src/` in current BPDSL layout. | `retain_preservation_only` | Labeled as previous description differing from current app-local model | Historical target not reflected in current app-local contracts |
| `renders` as view derived from DSL | T04 `design-flow.md` § Preserved source-of-truth roles | `spec:product.bpdsl.design_flow` | `spec:bpdsl.views.overview` — active app-local contract. Current model: `yaml/ → renders/`. | `retain_preservation_only` (for historical staging context); render contract is `replace_with_pointer` | Already present; row labeled preservation-only | Historical `dsl/ → renders/` flow differs from current `yaml/ → renders/`; staging row correctly non-canonical |
| `internal design` supplements DSL mapping | T04 `design-flow.md` § Deferred integration disposition | `spec:product.bpdsl.design_flow` | No current app-local BPDSL owner. Current BPDSL DSL overview mentions `impl design` for exceptions/concurrency but not as a formal artifact class or layer. | `retain_preservation_only` | Changed from "T09 classification" to historical description label | No current owner; no accepted integration claim; preserve historical context for migration review |
| `internal design` supplements handwritten route | T04 `design-flow.md` § Deferred integration disposition | `spec:product.bpdsl.design_flow` | No current app-local BPDSL owner. | `retain_preservation_only` | Changed from "T09 classification" to historical description label | Same as above |
| `records/spec` constrains DSL semantics | T04 `design-flow.md` § Deferred integration disposition (integration claim) | `spec:product.bpdsl.design_flow` | None. Current app-local BPDSL does not reference Design Records in any constraint role; no accepted integration contract defines this relationship. | `retain_preservation_only` | Changed from "T09 classification as relation candidate" to neutral historical label noting no accepted integration contract currently defines this relationship | Historical claim preserved with explicit statement that it is not an active integration rule and is not represented by current app-local BPDSL contracts |
| `records/spec` constrains internal routes | T04 `design-flow.md` § Deferred integration disposition | `spec:product.bpdsl.design_flow` | Existing evidence owners: V01-INV-DOCS-003, V01-ADR-088 | `retain_preservation_only` | Already labeled; no T09 placeholder; no change | Pre-existing correct disposition |
| `internal design` as artifact class | T04 `artifact-responsibilities.md` § Preserved responsibility rows | `spec:product.bpdsl.artifact_responsibilities` | No current app-local BPDSL owner for internal design as an artifact class. | `retain_preservation_only` | Changed from "T09 classification" to historical description label | Historical artifact role without current owner; retained for migration review |
| "Internal design requires final disposition" meta-statement | T04 `artifact-responsibilities.md` § Deferred integration disposition | `spec:product.bpdsl.artifact_responsibilities` | Not applicable (task progress tracking, not preserved content) | `delete_with_rationale` | Row deleted | Task-progress tracking row, not preservation evidence; the `internal design` row above preserves the substantive content |
| `brewprint DSL definition` as target implementation source | T04 `artifact-responsibilities.md` § Preserved responsibility rows | `spec:product.bpdsl.artifact_responsibilities` | None. Current app-local BPDSL doesn't describe a "DSL definition" artifact in this sense. | `retain_preservation_only` | No T09 placeholder; no change | Already correctly labeled preservation-only |
| `source implementation` (generated / bootstrap) | T04 `artifact-responsibilities.md` § Preserved responsibility rows | `spec:product.bpdsl.artifact_responsibilities` | None. Code generation is out of scope in current BPDSL. | `retain_preservation_only` | No T09 placeholder; no change | Already correctly labeled |
| `render` responsibility row | T04 `artifact-responsibilities.md` § Preserved responsibility rows | `spec:product.bpdsl.artifact_responsibilities` | `spec:bpdsl.views.overview` — active app-local render contract | `retain_preservation_only` (staging context); render contract owned app-locally | No T09 placeholder; no change | Historical `dsl/ → renders/` staging row; current render contract is app-local |
| `target implementation` responsibility row | T04 `artifact-responsibilities.md` § Preserved responsibility rows | `spec:product.bpdsl.artifact_responsibilities` | None. No `src/` in current BPDSL scope. | `retain_preservation_only` | No T09 placeholder; no change | Already correctly labeled |
| `impl note` responsibility row | T04 `artifact-responsibilities.md` § Preserved responsibility rows | `spec:product.bpdsl.artifact_responsibilities` | No current owner identified. | `retain_preservation_only` | No T09 placeholder; no change | Already correctly labeled |
| `dsl/` directory statement | T04 `repository-implementation-flow.md` § Preserved repository structure | `spec:product.bpdsl.repository_implementation_flow` | None. Current app-local project layout uses `yaml/`, `renders/`, `render_index.yaml`. | `retain_preservation_only` | No T09 placeholder; no change | Already correctly labeled preservation-only; historical model differs from current |
| `src/` directory statement | T04 `repository-implementation-flow.md` § Preserved repository structure | `spec:product.bpdsl.repository_implementation_flow` | None. No `src/` in current BPDSL scope. | `retain_preservation_only` | No T09 placeholder; no change | Already correctly labeled |
| `dsl/ → generated src/` relationship | T04 `repository-implementation-flow.md` § Preserved DSL-to-source relationship | `spec:product.bpdsl.repository_implementation_flow` | None. Code generation out of scope in current BPDSL. | `retain_preservation_only` | No T09 placeholder; no change | Historical target model; correctly labeled |
| Handwritten `src/` bootstrap path | T04 `repository-implementation-flow.md` § Preserved DSL-to-source relationship | `spec:product.bpdsl.repository_implementation_flow` | None. | `retain_preservation_only` | No T09 placeholder; no change | Historical bootstrap state; correctly labeled |
| Direct spec-to-handwritten-`src/` integration relationship | T04 `repository-implementation-flow.md` § Deferred integration disposition | `spec:product.bpdsl.repository_implementation_flow` | None. Current app-local BPDSL scope does not adopt or evaluate this integration relationship. | `retain_preservation_only` | Changed from "T09 classification" to neutral historical label; retained as migration-review evidence only | Historical claim preserved as migration context only; not an active integration rule |
| Internal design as source for handwritten route behavior | T04 `repository-implementation-flow.md` § Deferred integration disposition | `spec:product.bpdsl.repository_implementation_flow` | No current app-local BPDSL owner. | `retain_preservation_only` | Changed from "T09 classification" to historical label | Historical claim; no current owner |
| Product-level policy for implementation-bearing apps | T04 `repository-implementation-flow.md` § Deferred integration disposition | `spec:product.bpdsl.repository_implementation_flow` | None. Not a specific preserved statement; no accepted owner; aspirational policy that was never adopted. | `delete_with_rationale` | Row deleted | No specific preservation content; the `dsl/`/`src/` layout evidence in §§ Preserved repository structure and Preserved DSL-to-source relationship already captures the historical model fully |
| `Preservation status` review-gate rows (all three staging files) | T04/T05 task-progress wording | All three staging files | Not applicable (task-progress tracking) | N/A — replaced with durable contract | Changed "T09 finalizes and reviews temporary staging" to "Migration review obligation: Review all retained statements during BPDSL migration or when an explicit integration requirement is accepted" | Task-progress wording replaced with durable exit-review obligation |
| `Integration claim` status rows (all three staging files) | T04 placeholder | All three staging files | Not applicable | N/A — updated to final disposition | Changed "preserved for T09 classification" to "Historical; not adopted by PRODUCT. Retained as migration-review evidence only. Historical model differs from current app-local BPDSL contracts." | Final disposition recorded |
| `spec:product.bpdsl` Boundary section | T02 task-progress wording | `spec:product.bpdsl` index.md | Not applicable | N/A — rewritten as durable rule | "Current restructuring review checks..." replaced with permanent content-rule statement | Task-progress wording replaced with durable contract stating what content must remain |

#### Disposition counts

| disposition | count |
|---:|---:|
| `retain_preservation_only` | 19 |
| `replace_with_pointer` | 5 |
| `delete_with_rationale` | 3 |

### 3. Staging file inventory

| staging file | source material preserved | retained purpose | app-local overlap | final migration-review obligation |
|---|---|---|---|---|
| `spec:product.bpdsl` (`index.md`) | T02-created overview defining temporary staging contract | Staging contract definition: purpose, non-goals, rules, exit conditions, four-way exit classification | None | Review exit conditions and remove or redefine after BPDSL migration or explicit integration requirement |
| `spec:product.bpdsl.design_flow` (`design-flow.md`) | `project-artifact-model/design-flow.md` all sections; design and implementation artifact rows from `project-artifact-model/index.md` | Preserves historical Design artifact flow (ADR → spec → internal design → DSL → src → render → impl), source-of-truth roles, DSL-to-source rules, and V01 historical source references | `spec:bpdsl.views.overview` covers current render contract; `spec:bpdsl.dsl.project_layout` covers current project layout | Migration review: determine which rows describe current app-local behavior (with `yaml/`), which are superseded, and which require a new accepted integration decision |
| `spec:product.bpdsl.artifact_responsibilities` (`artifact-responsibilities.md`) | BPDSL and implementation-flow responsibility rows from `artifact-responsibility-matrix.md`; design and implementation artifact rows from `project-artifact-model/index.md` | Preserves historical artifact classes: internal design, brewprint DSL definition, source implementation, render, target implementation, impl note | `spec:bpdsl.views.overview` covers render; no app-local owner for internal design artifact class | Migration review: determine final ownership for internal design artifact class; reconcile DSL-definition row with current `yaml/`-based model |
| `spec:product.bpdsl.repository_implementation_flow` (`repository-implementation-flow.md`) | `dsl/`, `src/`, generated-source, and handwritten bootstrap statements from `repository-layout/index.md` | Preserves historical repository model: `<app>/dsl/`, `<app>/src/`, `dsl/ → generated src/`, bootstrap path | `spec:bpdsl.dsl.project_layout` covers current project layout (`yaml/`, `renders/`, `render_index.yaml`) | Migration review: reconcile historical `dsl/`/`src/` layout against current `yaml/`/`renders/` layout; determine if historical bootstrap concepts remain relevant |

### 4. Placeholder resolution

Pre-edit counts under `product/records/spec/bpdsl/**`:

| placeholder text | pre-edit count | post-edit count |
|---|---:|---:|
| `T09 classification` | 12 | 0 |
| `T09 finalizes` | 3 | 0 |
| `T09 handoff` | 0 | 0 |

Post-edit search: `rg -n "T09 classification|T09 finalizes|T09 handoff" product/records/spec/bpdsl` — exit code 1 (no matches).

### 5. Current-contract mismatch

#### Historical model (preserved in staging)

The staging files preserve the following previous descriptions:

- `<app>/dsl/` stores BPDSL definitions used as an implementation source.
- `dsl/ → generated src/` was the long-term target for an implementation-bearing app with an operational DSL pipeline.
- Handwritten `<app>/src/` was a bootstrap path while DSL support was insufficient.
- `render` was described as views derived from DSL definitions.
- Design Records (`records/spec`) was described as constraining DSL semantics.
- Internal design was described as supplementing the route from spec to DSL or source.

#### Current app-local BPDSL model (from `bpdsl/records/spec/**`)

- Source area: `yaml/` (brewprint YAML files) is the single source of truth. No `dsl/` directory.
- Generated output: `renders/` — auto-generated by `brewprint render`. Not editable by humans.
- Project layout: `{project-root}/yaml/`, `renders/`, `render_index.yaml`. No `src/` directory.
- Code generation: Explicitly **out of scope** in `spec:bpdsl.dsl.overview`.
- Design Records: No reference to Design Records constraining DSL semantics in any current app-local spec.
- Internal design layer: Referenced only informally in `spec:bpdsl.dsl.overview` responsibility split table as "impl design" handling exceptions/concurrency/transactions. Not a formal artifact class.
- Render derives from `yaml/` YAML definitions via `spec:bpdsl.views.overview`, not from a `dsl/` intermediate.

#### T09 action

T09 does not resolve this mismatch. Every retained staging statement is labeled as a previous or historical description; wording is neutral about what current app-local BPDSL contracts address. Reconciliation is deferred to BPDSL migration review.

### 6. Compatibility and migration findings

The following findings are recorded as migration-review inputs. None are new normative staging rules.

| finding | detail |
|---|---|
| `dsl/` vs `yaml/` source model divergence | The staged historical model describes `<app>/dsl/` as the BPDSL source area. The current app-local BPDSL spec defines `yaml/` as the source. Whether these represent the same concept under a renamed directory, or a substantive design change, must be determined during migration review. |
| Code generation scope | The staged model assumed `dsl/ → generated src/` as the target flow. The current app-local BPDSL explicitly excludes code generation from scope. Migration review must confirm whether code generation is permanently out of scope or deferred. |
| Internal design artifact class | The staged content preserves "internal design" as a potential artifact class (wiring from spec to DSL/source). No current app-local BPDSL spec defines this class. Migration review must decide whether to adopt, reject, or redesign this concept. |
| `records/spec` constraining DSL semantics | The staged model described Design Records as constraining DSL semantics. No current app-local BPDSL spec references Design Records in any constraint role. Migration review must decide whether a Design Records-to-BPDSL integration relationship is required, and if so, define its exact form through an accepted requirement or ADR. |
| Self-hosting behavior | T03 handoff mentioned self-hosting behavior as BPDSL-owned. No current app-local BPDSL spec defines self-hosting behavior. If this concept was meaningful, migration review should determine whether it requires a new spec section or is obsolete. |

### 7. Scope evidence

Pre-edit commands:

```text
git status --short -- product/records/spec/bpdsl product/records/tasks/spec/PRODUCT-TASK-SPEC-012-09-finalize-temporary-bpdsl-staging.md bpdsl/records/spec drmcp/records/spec v01
```

Output:
```text
 M drmcp/records/spec/design-records-mcp/resolver.md
 M drmcp/records/spec/design-records-mcp/schema/discovery.md
?? product/records/spec/bpdsl/
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-09-finalize-temporary-bpdsl-staging.md
```

(drmcp modifications are pre-existing T08 changes, not from T09.)

```text
git diff --cached --name-status
```

Output: no output (nothing staged).

Post-edit commands:

```text
git status --short -- product/records/spec/bpdsl product/records/tasks/spec/PRODUCT-TASK-SPEC-012-09-finalize-temporary-bpdsl-staging.md
```

Output:
```text
?? product/records/spec/bpdsl/
?? product/records/tasks/spec/PRODUCT-TASK-SPEC-012-09-finalize-temporary-bpdsl-staging.md
```

```text
git diff --name-status -- product/records/spec/bpdsl product/records/tasks/spec/PRODUCT-TASK-SPEC-012-09-finalize-temporary-bpdsl-staging.md
```

Output: no output (untracked files do not appear in ordinary `git diff`).

```text
git status --short -- bpdsl/records/spec drmcp/records/spec v01
```

Output:
```text
 M drmcp/records/spec/design-records-mcp/resolver.md
 M drmcp/records/spec/design-records-mcp/schema/discovery.md
```

(These are pre-existing T08 changes. `bpdsl/records/spec` and `v01` show no output — unchanged.)

```text
git diff --cached --name-status
```

Output: no output (nothing staged).

### 8. Validation

Pre-edit:

```text
python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color
```

Exit code: 0. Output: `[strict]  All 47 file(s) OK.`

```text
python -X utf8 product/src/tools/validate_spec.py bpdsl/records/spec --strict --no-color
```

Exit code: 0. Output: `[strict]  All 37 file(s) OK.`

Post-edit:

```text
python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color
```

Exit code: 0. Output: `[strict]  All 47 file(s) OK.`

```text
python -X utf8 product/src/tools/validate_spec.py bpdsl/records/spec --strict --no-color
```

Exit code: 0. Output: `[strict]  All 37 file(s) OK.`

### 9. Leakage checks

| check | command | result |
|---|---|---|
| No T09 placeholders remain | `rg -n "T09 classification\|T09 finalizes\|T09 handoff" product/records/spec/bpdsl` | Exit 1 (no matches). |
| No canonical BPDSL ownership claim introduced | `rg -n "canonical BPDSL" product/records/spec/bpdsl` | Exit 0; all matches are negative statements ("not a canonical BPDSL contract", "no canonical BPDSL ownership claim", "make no canonical BPDSL ownership claim"). |
| No active Design Records-to-BPDSL integration rule | `rg -n "Design Records.*constrains\|constrains.*DSL\|integration rule\|integration contract" product/records/spec/bpdsl` | Exit 0; only matches are the "does not define a new accepted Design Records-to-BPDSL integration contract" prohibition and the "not an active Design Records-to-BPDSL integration claim" label on the historical row. No normative integration rule introduced. |
| App-local BPDSL files unchanged | `git status --short -- bpdsl/records/spec` | No output for `bpdsl/records/spec`. |
| No broad ref sync performed | Only local text changes in `product/records/spec/bpdsl/**` — no `spec:product.concepts.*` refs modified. | Confirmed by scope of edits. |

### 10. Review corrections

Two review corrections applied after initial T09 completion. T09 status remains `done`.

#### Correction 1: unsupported conflict claims removed

Locations corrected:

- `design-flow.md` `## Deferred integration disposition`, row for `records/spec constrains DSL semantics`: wording that asserted an unsupported conflict with current app-local model replaced with neutral text stating the relationship is not represented by current app-local BPDSL contracts and no accepted integration contract currently defines it.
- `repository-implementation-flow.md` `## Deferred integration disposition`, row for direct spec-to-handwritten-`src/` implementation: same class of conflict-claim replaced with neutral text stating the current app-local BPDSL scope neither adopts nor evaluates this relationship.
- T09 disposition map rows for both items above: app-local-overlap column and rationale column conflict-claim wording replaced with neutral equivalents.
- T09 §5 T09 action sentence: rewritten to state that retained staging statements use neutral wording about what current app-local BPDSL contracts address, rather than asserting conflict or difference.

Rationale: `yaml/` being the current BPDSL source does not logically prohibit an external Design Records contract from constraining DSL semantics. Current app-local BPDSL specs do not explicitly prohibit handwritten application source derived from specifications; they simply do not define this integration relationship. Asserting a conflict overstates what the current app-local specs say. The historical `dsl/ → generated src/` model may validly be described as differing from the current `yaml/`-based, code-generation-out-of-scope model; that description was not changed.

#### Correction 2: T03 pointer-evidence wording qualified

Locations corrected:

- T09 disposition map, all five T03 `replace_with_pointer` rows (`file update, if any` column): overstated wording replaced with: `None; app-local owner identified in T09 Evidence. Existing PRODUCT text provides an owner-level pointer; exact canonical cross-owner ref synchronization remains T10 scope.`

Rationale: The current PRODUCT files (`spec:product.design_records.namespace_model.app_namespaces` and `spec:product.brewprint.namespaces.app_namespaces`) contain only a general owner-level pointer to BPDSL app-local specifications or migration review, not exact canonical refs to each individual app-local contract. Exact ref synchronization is T10 scope per the manifest.

#### Post-correction verification

- Conflict-phrase and T03-pointer-phrase scoped search: exit 1 (no matches).
- T09 placeholder search: exit 1 (no matches).
- Strict validation: `product/records/spec` and `bpdsl/records/spec` both exit 0.
- Disposition counts unchanged: 19 `retain_preservation_only`, 5 `replace_with_pointer`, 3 `delete_with_rationale`.

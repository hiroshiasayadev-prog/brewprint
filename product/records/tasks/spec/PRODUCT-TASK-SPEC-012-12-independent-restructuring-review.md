# PRODUCT-TASK-SPEC-012-12: Independent restructuring review

- **id**: PRODUCT-TASK-SPEC-012-12
- **status**: done
- **date**: 2026-06-25
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-11
- **outputs**:
  - Independent review verdict and findings recorded in this task

## Goal

Obtain an independent review of the completed semantic-layer restructuring before final closure.

## Work

- Review the final tree against `PRODUCT-ADR-SPEC-001` and `PRODUCT-INV-SPEC-005`.
- Review semantic ownership, prohibited content, dependency direction, and compatibility placement.
- Review temporary BPDSL staging for non-canonical wording and exit obligations.
- Review semantic batches separately from relocation and mechanical ref changes.
- Confirm unadopted future material received an explicit disposition.
- Classify findings by severity and identify required corrections.
- Do not modify files during the review.

## Done condition

- The reviewer issues `PASS` or `NEEDS REVISION`.
- Every finding names an exact file or section.
- Blockers and optional refinements are distinguished.
- The review states whether the Work Item may close after corrections.

## Verification

- Confirm the reviewer used the accepted ADR and source Investigation.
- Confirm all Work Item completion conditions were checked.
- Confirm the review examined current files rather than only task summaries.
- Confirm no review edit was applied directly.

## Evidence

> This evidence block reflects a second-pass review-record correction. The
> first-pass review concluded `NEEDS REVISION` with 5 MUST-FIX findings; that
> verdict and the architectural conclusions are preserved. A second reviewer
> inspected the same files and observed that two earlier MUST-FIX bundles
> mixed current-ownership defects with migration-provenance wording, that one
> registry-row qualifier was over-classified, and that a malformed Markdown
> table defect was under-classified. The severity model below has been
> corrected; no spec, ADR, or other reviewed file was modified by either pass.

### 1. Review identity and independence

- **Reviewer model**: Claude Opus 4.7 (claude-opus-4-7).
- **Implementation role in T01–T11**: none. This review did not perform any restructuring work; it consumed the existing working tree only.
- **Source of truth**: current working-tree files (intentional unstaged migration changes were treated as authoritative, per task brief).
- **Reviewed files modified**: none. The only file modified by this review (across both passes) is this task file (T12). No spec, ADR, investigation, work item, DRMCP, or BPDSL file was edited. Nothing was staged or committed.

### 2. Files reviewed

PRODUCT root + areas inspected (current working-tree state):

- Root router: `product/records/spec/index.md`.
- Design Records area: `product/records/spec/design-records/` — every file.
  - `design-records/index.md`.
  - `design-records/authoring-standards/`: `index.md`, `adr-authoring.md`, `agent-authoring-policy.md`, `artifact-boundary.md`, `investigation-authoring.md`, `requirement-authoring.md`, `spec-authoring.md`, `task-authoring.md`, `work-item-authoring.md`, `writing-standard.md`.
  - `design-records/namespace-model/`: `index.md`, `app-namespaces.md`, `artifact-id-grammar.md`, `existing-artifacts.md`, `subdomain-model.md`.
  - `design-records/repository-layout/`: `index.md`, `record-discovery-paths.md`.
  - `design-records/artifact-model/`: `index.md`, `artifact-responsibility-matrix.md`, `change-and-investigation-flow.md`, `traceability-boundary.md`.
  - `design-records/spec-format/`: `index.md`, `overview.md`, `document-shape.md`, `topics-table.md`, `spec-id-as-ref.md`, `validation-policy.md`, `follow-up-boundary.md`.
  - `design-records/traceability/`: `index.md`, `semantic-ref.md`, `artifact-refs.md`, `metadata-schema.md`, `resolve-and-validation.md`.
- Brewprint area: `product/records/spec/brewprint/` — every file.
  - `brewprint/index.md`, `brewprint/layout/index.md`.
  - `brewprint/namespaces/`: `index.md`, `app-namespaces.md`, `domain-catalog.md`.
  - `brewprint/compatibility/`: `index.md`, `existing-artifacts.md`, `legacy-id-compatibility.md`.
- Temporary PRODUCT-side BPDSL staging: `product/records/spec/bpdsl/` — every file.
  - `bpdsl/index.md`, `bpdsl/design-flow.md`, `bpdsl/artifact-responsibilities.md`, `bpdsl/repository-implementation-flow.md`.
- Empty shell: `product/records/spec/concepts/` confirmed empty (no `.md` files).

DRMCP files inspected to verify cross-owner refs:

- Directory listing: `drmcp/records/spec/design-records-mcp/tools/` (15 files: `overview.md` + 14 per-tool specs).
- `drmcp/records/spec/design-records-mcp/tools/overview.md` H1-adjacent `id` confirmed: `spec:drmcp.design_records_mcp.tools.overview`. The area is an `Overview`, not an `Index`; the bare ref `spec:drmcp.design_records_mcp.tools` is not the canonical ID of any current DRMCP spec.
- `drmcp/records/spec/design-records-mcp/overview.md` and `resolver.md` confirmed they reference `spec:drmcp.design_records_mcp.tools.overview` (correct canonical form).

BPDSL files inspected: only enough to confirm staging non-canonical framing and absence of duplication. No correctness review of BPDSL canonical specs performed.

Task and decision sources reviewed:

- `prompt_chappy.md` (reviewer protocol startup).
- `product/records/spec/design-records/authoring-standards/task-authoring.md`.
- `PRODUCT-ADR-SPEC-001` (accepted ownership boundary).
- `PRODUCT-INV-SPEC-005` (source classification).
- `PRODUCT-WORK-SPEC-012` (work item).
- All thirteen tasks `PRODUCT-TASK-SPEC-012-01` … `-13` (metadata + Evidence sections, with deep read of T11 Evidence).

**All 47 current PRODUCT specs were inspected during the first pass.** The second-pass correction only re-inspected the specific lines named in the revised findings to verify ownership-vs-provenance classification and the malformed table separators.

### 3. Verdict

**NEEDS REVISION.**

The accepted ADR boundary is preserved at the architectural level: top-level areas, ownership claims, dependency direction, compatibility placement, BPDSL staging contract, and removal of `concepts/` all match `PRODUCT-ADR-SPEC-001`. T11 strict validation reports zero errors. No BLOCKER finding.

The verdict is `NEEDS REVISION` because:

- one cross-owner authority ref in `agent-authoring-policy.md` does not resolve;
- migration-task identifiers (`T08`, `T09`, `PRODUCT-TASK-SPEC-012-08`) appear as current owners, treatments, section headings, and Topics descriptions in current contract prose across several Design Records, Brewprint, and traceability specs;
- two `## ... handoff` section headings name a migration task as their identity;
- two Markdown tables in `namespace-model/app-namespaces.md` carry three-column separators under two-column headers — a structural defect the temporary validator did not detect.

Migration-provenance wording that does **not** assign current ownership (e.g., "Historical disposition evidence is recorded in T05", "preserved here so T04 does not silently drop them") is recorded as ADVISORY rather than MUST-FIX; it would improve readability if moved to `## Sources` / `## Provenance` blocks with full task IDs, but it does not block closure.

**T13 can correct all MUST-FIX findings locally without architectural rework.** No file move, no ADR change, no re-planning is required. The Work Item may close after T13 applies the corrections in §7.

### 4. Findings

#### BLOCKER

None.

#### MUST-FIX

**M1 — Unresolved active cross-owner authority claim.**

| field | value |
|---|---|
| ID | M1 |
| Severity | MUST-FIX |
| File | `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md` |
| Section | `## Non-goals` (line 15) and `## Related specs` (line 49) |
| Evidence | Line 15: "Tool API contracts — those belong to `spec:drmcp.design_records_mcp.tools`." Line 49: `\| spec:drmcp.design_records_mcp.tools \| Authoritative DRMCP tool API contract. \|`. |
| Violated source | `PRODUCT-ADR-SPEC-001` Design Records non-goals; `PRODUCT-WORK-SPEC-012` Completion Condition "Mechanical refs, parent markers, and topic tables match the accepted final paths." |
| Why it matters | Both occurrences are present-tense authority claims, not illustrative path-derivation examples. The ref `spec:drmcp.design_records_mcp.tools` does not resolve to any current spec ID. The actual DRMCP tools area carries the public ID `spec:drmcp.design_records_mcp.tools.overview`. Per review brief §8, this is the "stale reference that should point to an existing child" case. |
| Required correction | Replace both occurrences with `spec:drmcp.design_records_mcp.tools.overview`. Do not weaken the wording further — the Non-goals row legitimately points outward; only the ref needs to resolve. |
| T13 owner | Yes. Mechanical ref correction. No DRMCP-side change required. |
| Closure effect | Blocks closure: leaves an active unresolved cross-owner authority claim in a current Design Records spec. |

**M2 — Migration-task IDs used as current section headings and adjacent section-level owner claims.**

| field | value |
|---|---|
| ID | M2 |
| Severity | MUST-FIX |
| File | `product/records/spec/design-records/artifact-model/traceability-boundary.md`; `product/records/spec/design-records/repository-layout/record-discovery-paths.md` |
| Section | `## T08 handoff` (traceability-boundary.md line 52; record-discovery-paths.md line 39), and the in-section owner claim "Concrete DRMCP behavior belongs to T08." (traceability-boundary.md line 54). |
| Evidence | traceability-boundary.md line 52–62: section titled `## T08 handoff` with body "Concrete DRMCP behavior belongs to T08." and a 4-row table whose right column is uniformly "T08 app-local DRMCP review." record-discovery-paths.md line 39–47: section titled `## T08 handoff` with body "The following previous statements were removed from PRODUCT normative text and handed off to T08." and a 3-row table with "T08 DRMCP app-local review." / "T08 may retain app-local provenance if needed." |
| Violated source | `PRODUCT-ADR-SPEC-001` Decision §"Top-level PRODUCT spec areas" and §"Dependency direction" (cross-owner references remain pointers to canonical owners, not to migration tasks). |
| Why it matters | A current `Reference` spec must not contain a top-level section whose identity is a migration task. `T08` is `PRODUCT-TASK-SPEC-012-08` — a closed work vehicle. Once the Work Item closes, a future reader will see a section heading and an owner column called `T08` that mean nothing without the migration context. |
| Required correction | Rename each heading to an owner-oriented heading such as `## DRMCP boundary` (or `## Cross-owner boundary`). Replace the in-section owner claim "Concrete DRMCP behavior belongs to T08." with a canonical-owner form, e.g., "Concrete DRMCP behavior belongs to DRMCP app-local specifications (`spec:drmcp.design_records_mcp.tools.overview`, `spec:drmcp.design_records_mcp.schema.discovery`, etc.)." Replace every `T08 app-local DRMCP review.` / `T08 DRMCP app-local review.` / `T08 may retain app-local provenance if needed.` cell with the canonical owner pointer or, for the provenance row, evergreen wording without a migration-task identifier. |
| T13 owner | Yes. Heading rename + cell rewrites; mechanical. |
| Closure effect | Blocks closure: leaves task-vehicle headings inside the current Design Records contract. |

**M3 — Migration-task IDs used as current owners, treatments, or boundary destinations in Design Records specs.**

| field | value |
|---|---|
| ID | M3 |
| Severity | MUST-FIX |
| File | `product/records/spec/design-records/namespace-model/subdomain-model.md`; `product/records/spec/design-records/artifact-model/index.md`; `product/records/spec/design-records/artifact-model/change-and-investigation-flow.md`; `product/records/spec/design-records/artifact-model/traceability-boundary.md`; `product/records/spec/design-records/traceability/resolve-and-validation.md` |
| Section | `subdomain-model.md` `## Boundary` line 40: treatment cell "PRODUCT-TASK-SPEC-012-08 handoff to DRMCP app-local specifications." — `## Artifact model boundary` of `artifact-model/index.md` line 95 ("BPDSL-specific preserved material remains temporary for T09 classification.") and line 96 ("DRMCP operational behavior belongs to T08.") — `change-and-investigation-flow.md` `## Deferred implementation tracking disposition` line 54: final-disposition cell "T08 handoff." — `traceability-boundary.md` `## Traceability contract boundary` line 67 ("DRMCP operational behavior belongs to T08.") — `resolve-and-validation.md` `## Excluded implementation behavior` line 99: owner cell "DRMCP, T08 handoff." and `## Resolve and validation boundary` line 105 ("Workflow orphan diagnostics, progress projection, traversal, cycle detection, and execution-order checks belong to T08."). |
| Violated source | `PRODUCT-ADR-SPEC-001` §"Area ownership" (Design Records cross-owner references "remain pointers" — i.e., to canonical owners); task-authoring guide ("Tasks do not become canonical authority for accepted specs or decisions"). |
| Why it matters | These are present-tense ownership / treatment / boundary claims inside the current Design Records contract. They confuse the reader about who owns the referenced content and they will become opaque once `PRODUCT-WORK-SPEC-012` closes. |
| Required correction | Replace each task reference with the canonical owner: "DRMCP app-local specifications", "BPDSL app-local specifications", or specific resolvable canonical refs. For the `subdomain-model.md` row, the treatment column should read "DRMCP app-local specifications" rather than the full `PRODUCT-TASK-SPEC-012-08` task ID. |
| T13 owner | Yes. Each replacement is mechanical and localized. |
| Closure effect | Blocks closure: leaves task-vehicle ownership claims inside current Design Records specs across five files / seven lines. |

**M4 — Migration-task IDs used as current owners, treatments, headings, or Topics descriptions in Brewprint and traceability specs.**

| field | value |
|---|---|
| ID | M4 |
| Severity | MUST-FIX |
| File | `product/records/spec/brewprint/namespaces/app-namespaces.md`; `product/records/spec/brewprint/compatibility/existing-artifacts.md`; `product/records/spec/design-records/traceability/index.md`; `product/records/spec/design-records/traceability/semantic-ref.md` |
| Section | `brewprint/namespaces/app-namespaces.md` `## Boundary` line 34 (treatment cell "T08 handoff to DRMCP app-local specifications.") and line 35 (treatment cell "T09 handoff to BPDSL migration review or app-local specifications."); `brewprint/compatibility/existing-artifacts.md` `## Boundary` line 54 (treatment cell "T08 handoff to DRMCP app-local specifications."); `traceability/index.md` `## Topics` line 51 (summary cell "Current `spec:` ref class and obsolete semantic-ref assumptions removed by T05."); `traceability/semantic-ref.md` `## What this is` line 10 (same "removed by T05" phrasing) and `## Obsolete assumptions removed by T05` heading on line 35. |
| Violated source | `PRODUCT-ADR-SPEC-001` §"Area ownership" — Brewprint must point to canonical owners, not name a migration task as the current treatment; same boundary applies inside the traceability sub-area; spec-format Topics-table rules require the summary to describe the current topic rather than embed the migration task. |
| Why it matters | The Brewprint treatment columns are current registry-area boundary cells, not provenance notes. The `## Obsolete assumptions removed by T05` heading is a section heading whose identity is a migration task, and the parent index's Topics summary inherits the same wording. |
| Required correction | In the Brewprint Boundary rows, replace "T08 handoff" / "T09 handoff" with "DRMCP app-local specifications" and "BPDSL app-local specifications". In `semantic-ref.md`, rename the heading to `## Obsolete assumptions` (or `## Removed assumptions`) and rewrite the `## What this is` sentence to describe the current `spec:` ref class without naming the migration task. In `traceability/index.md`, update the Topics summary to match. |
| T13 owner | Yes. Heading rename + cell rewrites; mechanical. |
| Closure effect | Blocks closure for the same class of reason as M2 / M3. |

**M5 — Malformed Markdown table separators in current Design Records spec.**

| field | value |
|---|---|
| ID | M5 (promoted from former A3) |
| Severity | MUST-FIX |
| File | `product/records/spec/design-records/namespace-model/app-namespaces.md` |
| Section | `## Current contract` (header line 17 + separator line 18); `## Related specs` (header line 43 + separator line 44). |
| Evidence | Header rows declare two columns (`\| rule \| contract \|`, `\| ref \| relation \|`) but the separator rows declare three (`\|---\|---\|---\|`). Some Markdown parsers render a phantom third column; others fall back to plain prose. |
| Violated source | Spec-format `## Validation rules` / `## Errors` policy (document shape must be parser-friendly); `PRODUCT-WORK-SPEC-012` Completion Condition "Validation reports no unresolved references or structural diagnostics caused by this migration." The structural diagnostic exists; the temporary PRODUCT validator did not flag it. |
| Why it matters | A current Design Records `Reference` spec carries a Markdown structural defect. The spec is parsed differently by different tools and may break section-level navigation. The strict validator pass count of "47 file(s) OK" reflects a validator limitation, not the absence of a structural defect. |
| Required correction | Change both separator rows to `\|---\|---\|` (two cells) to match the two-column headers. |
| T13 owner | Yes. Two single-line edits. |
| Closure effect | Blocks closure: current Design Records spec has a structural Markdown defect. |

#### ADVISORY

**A1 — Illustrative path-derivation example uses a nonexistent file and the same ref as M1.**

| field | value |
|---|---|
| ID | A1 |
| Severity | ADVISORY |
| File | `product/records/spec/design-records/authoring-standards/spec-authoring.md` (line 43); `product/records/spec/design-records/spec-format/spec-id-as-ref.md` (line 53) |
| Section | "## Rules" / "### ID grammar" example tables. |
| Evidence | Both files include the example row `\| drmcp/records/spec/design-records-mcp/tools.md \| spec:drmcp.design_records_mcp.tools \|`. The path `tools.md` does not exist; the area uses `tools/overview.md` instead. T11 §10 classified these as `false_positive`. |
| Why it matters | Picking a different example removes the visual collision with the M1 stale ref. Does not block closure because the row is explicitly illustrative. |
| Required correction | Optional swap to a non-colliding example (e.g., a path / ref pair that actually resolves). |
| T13 owner | Only if the user accepts optional cleanup. |
| Closure effect | Optional. |

**A2 — Hyphen-form example ref in Boundary table.**

| field | value |
|---|---|
| ID | A2 |
| Severity | ADVISORY |
| File | `product/records/spec/design-records/spec-format/spec-id-as-ref.md` |
| Section | `## Boundary` line 88 (`spec:trace.semantic-ref`). |
| Evidence | The example string is wrapped in backticks and used to illustrate the legacy hyphen-form pattern. T11 §10 classified as `false_positive`. |
| Why it matters | Intentional illustration of an obsolete grammar form. |
| Required correction | None. |
| T13 owner | No action. |
| Closure effect | Optional / no action. |

**A3 — Migration-provenance wording in current specs.**

| field | value |
|---|---|
| ID | A3 |
| Severity | ADVISORY |
| File | `product/records/spec/design-records/artifact-model/index.md`; `product/records/spec/design-records/artifact-model/artifact-responsibility-matrix.md`; `product/records/spec/design-records/artifact-model/traceability-boundary.md`; `product/records/spec/design-records/traceability/artifact-refs.md`; `product/records/spec/design-records/traceability/metadata-schema.md`; `product/records/spec/design-records/traceability/resolve-and-validation.md`; `product/records/spec/brewprint/namespaces/domain-catalog.md` |
| Section | `artifact-model/index.md` `## Disposition of previous ownership statements` line 63 ("preserved here so T04 does not silently drop them"), line 72 (right-column "T05 Evidence."), line 73 ("T04 does not create a T05 action for this statement."), line 74 (same pattern), and `## Artifact model boundary` line 98 ("Historical per-item disposition evidence is recorded in T05.") — `artifact-responsibility-matrix.md` `## Extracted implementation responsibilities` line 32 (status cell "...preserved only in temporary BPDSL staging for T09 classification.") — `traceability-boundary.md` `## Traceability contract boundary` line 70 ("Historical disposition evidence is recorded in T05.") — `traceability/artifact-refs.md` `## Reference boundary` line 89 ("Historical disposition evidence is recorded in T05.") — `traceability/metadata-schema.md` `## Metadata boundary` line 86 ("Historical disposition evidence is recorded in T05.") — `traceability/resolve-and-validation.md` `## Resolve and validation boundary` line 106 ("Historical disposition evidence is recorded in T05.") — `brewprint/namespaces/domain-catalog.md` `## Future candidates` lines 35 and 36 (registry note "No scoped current app-aware record ID was found for this domain during T03 verification."). |
| Evidence | All listed lines name a migration task as the place where evidence lives, or describe a migration-time observation, without assigning current ownership or treatment. |
| Why it matters | Readability degrades after `PRODUCT-WORK-SPEC-012` closes. The lines do not violate semantic ownership, dependency direction, or spec-area boundaries. Migration provenance is legitimate evidence; it would be clearer in a `## Sources` / `## Provenance` block with full task IDs (e.g., `PRODUCT-TASK-SPEC-012-05`) rather than bare `T05` in body prose. |
| Required correction | Optional: relocate to a Sources / Provenance block and use full task IDs. The `domain-catalog.md` rows can either drop "during T03 verification" or adopt evergreen wording such as "in the current Brewprint inventory". |
| T13 owner | Optional. |
| Closure effect | Does not block closure. Was previously over-classified as part of M3 / M4 / M5; downgraded after the second-pass severity correction. |

**A4 — DRMCP-side `tools` area is an Overview, not an Index — out of scope.**

| field | value |
|---|---|
| ID | A4 |
| Severity | ADVISORY — out of scope for `PRODUCT-WORK-SPEC-012` |
| File | `drmcp/records/spec/design-records-mcp/tools/overview.md` (DRMCP-side, not edited by this review). |
| Section | Public ID line 3 (`spec:drmcp.design_records_mcp.tools.overview`). |
| Evidence | The DRMCP tools area is expressed as an `Overview`, not an `Index`. PRODUCT specs that want to reference "the DRMCP tools area" must use `spec:drmcp.design_records_mcp.tools.overview`. |
| Why it matters | A future DRMCP-side rename (e.g., to `tools/index.md`) would let downstream consumers use the bare `tools` ref. This is a DRMCP authoring choice, not a PRODUCT migration defect. |
| Required correction | **No T13 action. No closure effect.** Recorded only so future DRMCP-side work has a pointer. |
| T13 owner | No. DRMCP-side decision. |
| Closure effect | None. |

### 5. Area review results

**Root router** (`product/records/spec/index.md`): satisfies the ADR router contract. Routes by semantic ownership, names all three top-level areas, states ownership boundary, dependency direction, registry-fact prohibition, and authoring boundary. Does not restate child contracts. Does not invent a new top-level semantic owner.

**Design Records** (`design-records/**`): architecture correct. Generic authoring / format / namespace / repository-layout / traceability / artifact-model semantics are owned here and use pointers (not duplications) for DRMCP and BPDSL behavior. Brewprint registry tables are not duplicated. App-neutral examples are used. **Local defects requiring T13**: M1 (one stale cross-owner ref), M2 (two `T08 handoff` section headings + adjacent owner claims), M3 (task-vehicle ownership claims in current contract rows), M5 (malformed table separators in `namespace-model/app-namespaces.md`). **Advisory**: A1, A2, A3.

**Brewprint** (`brewprint/**`): architecture correct. Profile is observational; namespace registry present; compatibility lives under `brewprint/compatibility/` (not at top level). Generic Design Records rules not redefined. **Local defects requiring T13**: M4 (Boundary treatment rows in `namespaces/app-namespaces.md` and `compatibility/existing-artifacts.md`). **Advisory**: A3 (`domain-catalog.md` "during T03 verification").

**Temporary BPDSL staging** (`bpdsl/**`): architecture correct. `bpdsl/index.md` states purpose, ownership status, expected final owner, allowed content, context limit, migration trigger, non-goals, rules, boundary, and exit conditions. Children framed as preservation-only with "no PRODUCT canonical ownership claim" rows. Source-mapping tables retain old `concepts/` paths only as historical evidence. References to canonical BPDSL refs are pointers only. No canonical self-hosting owner invented.

**Compatibility placement**: V01 compatibility, historical attribution, issued-ID retention, and migration state are all under `brewprint/compatibility/`. Generic Design Records grammar is referenced rather than redefined. Correct.

**Future-material disposition**: T05 dispositions honored. Inactive `yaml:` / `internal-design:` / `coverage:` endpoints, fixture mechanisms, MCP writer contracts, and external-relation artifacts are removed from current normative specs; rationale preserved via V01-INV-DOCS-002/003 + V01-ADR-088 references. No catalog of unadopted mechanisms remains in current contract bodies.

**Dependency direction**: verified. Design Records → DRMCP has no normative dependency. The only Design Records → DRMCP refs are (a) `agent-authoring-policy.md`'s outward pointer (M1) and (b) `metadata_grammar` / `authoring_transaction_schema` references in authoring guides — these match the ADR's allowed pointer pattern. Brewprint → Design Records is consistent. Temporary BPDSL staging → canonical BPDSL is contextual-pointer only.

**Cross-owner refs**: the single unresolved active ref is M1. All other cross-owner pointers resolve.

### 6. Work Item completion-condition matrix

| Completion condition | Architectural placement | Local correction status |
|---|---|---|
| `product/records/spec/index.md` routes placement without duplicating child contracts. | satisfied | satisfied |
| `design-records/`, `brewprint/`, and temporary `bpdsl/` have explicit overview contracts. | satisfied | satisfied |
| No active PRODUCT spec remains under `product/records/spec/concepts/`. | satisfied | satisfied (`concepts/` empty shell) |
| App-independent Design Records semantics reside under `design-records/`. | satisfied | **requires T13 correction** — M1, M2, M3, M5 are local-prose / structural defects inside Design Records specs. |
| Brewprint current-state and registry material resides under `brewprint/`. | satisfied | **requires T13 correction** — M4 Brewprint Boundary rows. |
| V01 compatibility, historical attribution, issued-ID retention, and migration state reside under `brewprint/compatibility/`. | satisfied | satisfied |
| Temporary PRODUCT-side BPDSL staging states its non-canonical status, allowed content, prohibited work, review boundary, migration trigger, and exit condition. | satisfied | satisfied |
| PRODUCT Design Records specs do not restate DRMCP operational behavior or canonical BPDSL internals. | satisfied | **requires T13 correction** — cross-owner content is by pointer, but M1 still points to a non-resolving ref. |
| Unadopted future integration material is transferred to appropriate follow-up artifacts or removed after evidence transfer. | satisfied | satisfied |
| Semantic rewrites, relocation, app-local handoff, broad ref synchronization, and validation cleanup remain separately reviewable. | satisfied | satisfied — task split T02–T11 matches ADR change-class separation. |
| Mechanical refs, parent markers, and topic tables match the accepted final paths. | satisfied at the path level | **requires T13 correction** — M1 active ref correction. |
| Validation reports no unresolved references or structural diagnostics caused by this migration. | satisfied for unresolved-ref counts | **requires T13 correction** — strict validator passes, but the validator did not detect M5 (malformed table separators). Counting validator output as "fully structurally clean" without qualification overstates the result. |
| Pre-existing diagnostics are separated from migration-caused diagnostics. | satisfied | satisfied (T11 §10 classifies residuals). |
| `v01/` remains unchanged. | satisfied | satisfied (T11 §11; current `git status` shows no `v01/` modification). |
| Completion evidence lists exact files, review gates, and validation results for every task. | not independently verifiable | T11 evidence is detailed; T02–T10 evidence was sampled, not exhaustively re-audited. Treat as advisory only. |

Overall: architecture satisfied at the area level; M1–M5 are the closure-blocking local-prose / structural corrections. The Work Item may close after T13 applies them.

**Validator limitation noted explicitly**: the temporary PRODUCT spec validator (`product/src/tools/validate_spec.py`, strict mode) reports zero errors on 47 PRODUCT files but did not detect the M5 malformed table separators. T11's "47 file(s) OK" is accurate within validator scope; it is not a guarantee of structural correctness.

### 7. T13 correction boundary

**Required corrections (closure-blocking)**:

1. **M1** — In `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`, replace both occurrences of `spec:drmcp.design_records_mcp.tools` with `spec:drmcp.design_records_mcp.tools.overview` (lines 15 and 49).
2. **M2** — In `product/records/spec/design-records/artifact-model/traceability-boundary.md` (lines 52–67) and `product/records/spec/design-records/repository-layout/record-discovery-paths.md` (lines 39–47): rename the `## T08 handoff` sections to owner-oriented headings (e.g., `## DRMCP boundary`), rewrite "Concrete DRMCP behavior belongs to T08." to a canonical-owner form, and rewrite all `T08 ... review` / `T08 may retain ...` cells to canonical-owner pointers.
3. **M3** — Replace migration-task ownership / treatment / boundary wording with canonical-owner pointers in:
   - `design-records/namespace-model/subdomain-model.md` line 40 (treatment cell);
   - `design-records/artifact-model/index.md` lines 95–96 (`## Artifact model boundary`);
   - `design-records/artifact-model/change-and-investigation-flow.md` line 54 (final-disposition cell);
   - `design-records/artifact-model/traceability-boundary.md` line 67 (`## Traceability contract boundary`);
   - `design-records/traceability/resolve-and-validation.md` line 99 (owner cell) and line 105 (`## Resolve and validation boundary`).
4. **M4** — Replace Brewprint Boundary treatment cells and rename heading / Topics summary embedding `T05` in:
   - `brewprint/namespaces/app-namespaces.md` lines 34, 35 (`## Boundary`);
   - `brewprint/compatibility/existing-artifacts.md` line 54 (`## Boundary`);
   - `design-records/traceability/index.md` line 51 (Topics summary);
   - `design-records/traceability/semantic-ref.md` line 10 (`## What this is`) and line 35 (`## Obsolete assumptions removed by T05` heading → rename to e.g. `## Obsolete assumptions`).
5. **M5** — In `product/records/spec/design-records/namespace-model/app-namespaces.md`, change the separator rows on line 18 and line 44 from `\|---\|---\|---\|` to `\|---\|---\|` to match the two-column headers above them.

After applying M1–M5, re-run `validate_spec.py --strict` for all three roots and the combined graph check. Diagnostic counts should remain at zero PRODUCT / DRMCP / BPDSL errors.

**Optional advisories (do not block closure, do not apply unless the user accepts)**:

- **A1** — Swap the `tools.md` path-derivation example in `design-records/authoring-standards/spec-authoring.md` line 43 and `design-records/spec-format/spec-id-as-ref.md` line 53 to a non-colliding example.
- **A2** — Leave `spec:trace.semantic-ref` example as-is.
- **A3** — Relocate provenance-only wording (the lines and files listed in the A3 finding) into `## Sources` / `## Provenance` blocks with full task IDs (`PRODUCT-TASK-SPEC-012-04`, `-05`, etc.). The `domain-catalog.md` "during T03 verification" rows may be dropped or rephrased as "in the current Brewprint inventory".
- **A4** — Not a T13 action. Recorded only so future DRMCP-side work has a pointer.

T13 must not perform additional semantic rewrites, file moves, or DRMCP-side edits beyond M1–M5.

### 8. Scope evidence

- No reviewed file changed. Only this T12 task file changed (across both review passes).
- Nothing staged.
- `v01/` unchanged.
- No commit.
- No reviewed PRODUCT spec, DRMCP spec, BPDSL spec, ADR, investigation, requirement, or work item file modified by either review session.

# PRODUCT-WORK-SPEC-012: PRODUCT spec semantic-layer restructuring

- **id**: PRODUCT-WORK-SPEC-012
- **status**: done
- **date**: 2026-06-24
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **impact_refs**:
  - spec:product.design_records.authoring_standards
  - spec:product.design_records.namespace_model
  - spec:product.design_records.repository_layout
  - spec:product.design_records.spec_format
  - spec:product.design_records.traceability
  - spec:product.design_records.artifact_model
  - spec:product.brewprint
  - spec:product.bpdsl
- **tasks**:
  - PRODUCT-TASK-SPEC-012-01
  - PRODUCT-TASK-SPEC-012-02
  - PRODUCT-TASK-SPEC-012-03
  - PRODUCT-TASK-SPEC-012-04
  - PRODUCT-TASK-SPEC-012-05
  - PRODUCT-TASK-SPEC-012-06
  - PRODUCT-TASK-SPEC-012-07
  - PRODUCT-TASK-SPEC-012-08
  - PRODUCT-TASK-SPEC-012-09
  - PRODUCT-TASK-SPEC-012-10
  - PRODUCT-TASK-SPEC-012-11
  - PRODUCT-TASK-SPEC-012-12
  - PRODUCT-TASK-SPEC-012-13

## Goal

Implement the semantic ownership boundary accepted by `PRODUCT-ADR-SPEC-001`.

Replace the current `concepts/` ownership model with explicit PRODUCT spec areas:

- `design-records/`;
- `brewprint/`;
- temporary `bpdsl/` staging.

Complete the restructuring without combining semantic rewrites, relocation, app-local handoff, broad ref synchronization, and validation cleanup in one change.

## Boundary

In scope:

- Build an exact migration manifest from `PRODUCT-INV-SPEC-005`.
- Create `product/records/spec/index.md` as the placement router.
- Create or revise top-level area overviews.
- Rewrite mixed specifications in small semantic batches.
- Move single-owner specifications after semantic cleanup.
- Place Brewprint compatibility under `brewprint/compatibility/`.
- Isolate preservation-only BPDSL material under temporary PRODUCT staging.
- Hand off clearly app-local DRMCP or BPDSL content in owner-specific batches.
- Extract unadopted future integration material or remove it after evidence transfer.
- Synchronize refs only after target ownership and paths are accepted.
- Validate the resulting tree and remove migration-caused stale refs.

Out of scope:

- DRMCP redesign.
- Full BPDSL migration or BPDSL hierarchy normalization.
- Final BPDSL ownership decisions beyond the temporary staging contract.
- New Design Records-to-BPDSL integration semantics.
- New parser, resolver, render, generation, runtime, UI, or MCP contracts.
- Unrelated PRODUCT spec redesign.
- Modification of `v01/`.

## Impact Scope

| ref or area | impact |
|---|---|
| `product/records/spec/index.md` | Create the placement router. |
| `design-records/` | Become the owner of app-independent Design Records semantics. |
| `brewprint/` | Own current profile, namespaces, layout, and compatibility. |
| `bpdsl/` | Provide temporary non-canonical preservation staging. |
| `concepts/` | Remove as an active semantic area after migration. |
| App-local DRMCP and BPDSL specs | Receive only clearly owned handoff content or pointers. |
| Downstream refs | Update mechanically after target paths are accepted. |

## Task flow

1. `PRODUCT-TASK-SPEC-012-01`: Confirm the migration manifest and validation baseline.
2. `PRODUCT-TASK-SPEC-012-02`: Create the root placement router and top-level area overviews.
3. `PRODUCT-TASK-SPEC-012-03`: Split namespace, Brewprint profile, and compatibility semantics.
4. `PRODUCT-TASK-SPEC-012-04`: Split artifact-model and repository-layout semantics.
5. `PRODUCT-TASK-SPEC-012-05`: Reconcile traceability and extract future material.
6. `PRODUCT-TASK-SPEC-012-06`: Relocate authoring standards and spec format.
7. `PRODUCT-TASK-SPEC-012-07`: Relocate the cleaned semantic areas.
8. `PRODUCT-TASK-SPEC-012-08`: Apply the DRMCP app-local handoff.
9. `PRODUCT-TASK-SPEC-012-09`: Finalize temporary BPDSL staging.
10. `PRODUCT-TASK-SPEC-012-10`: Synchronize refs mechanically.
11. `PRODUCT-TASK-SPEC-012-11`: Validate and clean migration diagnostics.
12. `PRODUCT-TASK-SPEC-012-12`: Perform independent restructuring review.
13. `PRODUCT-TASK-SPEC-012-13`: Apply required corrections and close the Work Item.

Review gates:

- Accept the migration manifest before PRODUCT spec edits.
- Review each semantic batch before relocation.
- Review relocation separately from content changes.
- Review app-local handoff separately for DRMCP and BPDSL.
- Run broad ref synchronization only after paths are stable.

## Task Candidates

| task | scope |
|---|---|
| PRODUCT-TASK-SPEC-012-01 | Confirm the full migration manifest and pre-migration validation baseline. |
| PRODUCT-TASK-SPEC-012-02 | Create the root router and three top-level area overview contracts. |
| PRODUCT-TASK-SPEC-012-03 | Separate namespace semantics, Brewprint registry facts, and V01 compatibility. |
| PRODUCT-TASK-SPEC-012-04 | Separate Design Records artifact and layout semantics from BPDSL material. |
| PRODUCT-TASK-SPEC-012-05 | Reconcile active traceability and extract unadopted integration material. |
| PRODUCT-TASK-SPEC-012-06 | Relocate authoring standards and spec-format files without semantic rewrites. |
| PRODUCT-TASK-SPEC-012-07 | Relocate the remaining cleaned Design Records semantic areas. |
| PRODUCT-TASK-SPEC-012-08 | Preserve required DRMCP-owned operational content app-locally. |
| PRODUCT-TASK-SPEC-012-09 | Finalize temporary non-canonical BPDSL staging. |
| PRODUCT-TASK-SPEC-012-10 | Synchronize canonical refs, parent markers, topic tables, and links. |
| PRODUCT-TASK-SPEC-012-11 | Validate the final tree and correct migration-caused diagnostics. |
| PRODUCT-TASK-SPEC-012-12 | Review the completed restructuring independently. |
| PRODUCT-TASK-SPEC-012-13 | Apply required review corrections and close the Work Item. |

Semantic rewrite tasks should target three to seven files when practical.
Task creation must preserve the change-class separation defined by `PRODUCT-ADR-SPEC-001`.

## Completion Condition

- `product/records/spec/index.md` routes placement without duplicating child contracts.
- `design-records/`, `brewprint/`, and temporary `bpdsl/` have explicit overview contracts.
- No active PRODUCT spec remains under `product/records/spec/concepts/`.
- App-independent Design Records semantics reside under `design-records/`.
- Brewprint current-state and registry material resides under `brewprint/`.
- V01 compatibility, historical attribution, issued-ID retention, and migration state reside under `brewprint/compatibility/`.
- Temporary PRODUCT-side BPDSL staging states its non-canonical status, allowed content, prohibited work, review boundary, migration trigger, and exit condition.
- PRODUCT Design Records specs do not restate DRMCP operational behavior or canonical BPDSL internals.
- Unadopted future integration material is transferred to appropriate follow-up artifacts or removed after evidence transfer.
- Semantic rewrites, relocation, app-local handoff, broad ref synchronization, and validation cleanup remain separately reviewable.
- Mechanical refs, parent markers, and topic tables match the accepted final paths.
- Validation reports no unresolved references or structural diagnostics caused by this migration.
- Pre-existing diagnostics are separated from migration-caused diagnostics.
- `v01/` remains unchanged.
- Completion evidence lists exact files, review gates, and validation results for every task.

## Evidence

### 1. Closure verdict

| item | value |
|---|---|
| Status | `done` |
| Target ownership model implemented | `design-records/`, `brewprint/`, temporary `bpdsl/` per `PRODUCT-ADR-SPEC-001` |
| T12 verdict | `NEEDS REVISION` — 0 BLOCKER, 5 MUST-FIX, 4 ADVISORY |
| Required findings resolved by T13 | M1, M2, M3, M4, M5 — all five |
| BLOCKER remaining | none |
| Advisory dispositions | A1–A4 explicitly not applied; recorded in T13 Evidence §3 |

### 2. Final ownership tree

```text
product/records/spec/
  index.md                       spec:product — placement router; routes by semantic ownership
  design-records/                spec:product.design_records — app-independent Design Records semantics
    authoring-standards/         spec:product.design_records.authoring_standards
    namespace-model/             spec:product.design_records.namespace_model
    repository-layout/           spec:product.design_records.repository_layout
    spec-format/                 spec:product.design_records.spec_format
    artifact-model/              spec:product.design_records.artifact_model
    traceability/                spec:product.design_records.traceability
  brewprint/                     spec:product.brewprint — Brewprint profile, namespaces, layout, compatibility
    namespaces/                  spec:product.brewprint.namespaces
    layout/                      spec:product.brewprint.layout
    compatibility/               spec:product.brewprint.compatibility
  bpdsl/                         spec:product.bpdsl — temporary non-canonical staging only
  concepts/                      empty shell (no .md files — removed as active semantic area)
```

Cross-owner pointers only; no canonical BPDSL or DRMCP content owned here.

### 3. Child-task evidence matrix

| task | final status | primary output or gate | evidence result |
|---|---|---|---|
| PRODUCT-TASK-SPEC-012-01 | done | Migration manifest (40-file inventory) and pre-migration baseline (0 strict errors) | T01 Evidence: full file inventory, baseline PRODUCT/DRMCP/BPDSL strict validation 0 errors, baseline `spec:product.concepts.*` ref count |
| PRODUCT-TASK-SPEC-012-02 | done | `product/records/spec/index.md`; `design-records/index.md`; `brewprint/index.md`; `bpdsl/index.md` | T02 Evidence: four router/overview files created; each confirms owned content, prohibited content, and navigation links |
| PRODUCT-TASK-SPEC-012-03 | done | `brewprint/namespaces/`; `brewprint/compatibility/`; generic Design Records namespace-model semantics separated | T03 Evidence: six source files processed; generic semantics → `design-records/namespace-model/`; Brewprint profile facts → `brewprint/namespaces/`; compatibility → `brewprint/compatibility/` |
| PRODUCT-TASK-SPEC-012-04 | done | Cleaned `design-records/artifact-model/` and `design-records/repository-layout/`; BPDSL staging content | T04 Evidence: seven source files processed; artifact-model and repository-layout separated; BPDSL staging initiated |
| PRODUCT-TASK-SPEC-012-05 | done | Cleaned `design-records/traceability/`; unadopted material dispositioned via V01-ADR-088, V01-INV-DOCS-002, V01-INV-DOCS-003 | T05 Evidence: seven traceability source files processed; `coverage:`, `internal-design:`, and `yaml:` endpoints removed with evidence-transfer; `follow_up_candidates` boundary cleaned |
| PRODUCT-TASK-SPEC-012-06 | done | `design-records/authoring-standards/`; `design-records/spec-format/` | T06 Evidence: authoring-standards and spec-format directories moved from `concepts/` to `design-records/`; no semantic rewrites |
| PRODUCT-TASK-SPEC-012-07 | done | `design-records/namespace-model/`; `design-records/repository-layout/`; `design-records/traceability/`; `design-records/artifact-model/` | T07 Evidence: four remaining cleaned semantic areas relocated from `concepts/` to `design-records/`; `concepts/` emptied |
| PRODUCT-TASK-SPEC-012-08 | done | DRMCP app-local handoff applied; `drmcp/records/spec/design-records-mcp/` updated | T08 Evidence: DRMCP-owned content handed off; PRODUCT normative text confirmed not restating DRMCP operational behavior |
| PRODUCT-TASK-SPEC-012-09 | done | `product/records/spec/bpdsl/index.md` and preservation-only BPDSL staging files | T09 Evidence: BPDSL staging finalized with non-canonical status, migration trigger, exit conditions, and preservation-only framing |
| PRODUCT-TASK-SPEC-012-10 | done | Updated canonical refs, parent markers, topic tables, and downstream links | T10 Evidence: mechanical ref synchronization complete; `spec:product.concepts.*` refs replaced; parent/Topics markers updated; 0 active stale refs |
| PRODUCT-TASK-SPEC-012-11 | done | Validation results; two migration-caused stale refs corrected | T11 Evidence: 47+30+37=114 strict validator 0 errors; combined graph 0 duplicates/unresolved; 2 migration-caused corrections applied; 4 pre-existing/false-positive diagnostics classified; `concepts/` confirmed empty |
| PRODUCT-TASK-SPEC-012-12 | done | Independent review verdict and findings | T12 Evidence: NEEDS REVISION; 0 BLOCKER, 5 MUST-FIX, 4 ADVISORY; architectural placement confirmed satisfactory; M1–M5 required corrections identified; A1–A4 advisory dispositions stated; T13 correction boundary defined |
| PRODUCT-TASK-SPEC-012-13 | done | M1–M5 corrections; complete T13 and Work Item Evidence | T13 Evidence: all five MUST-FIX findings resolved; 14 files edited; strict validators 0 errors; combined graph PASS; exact finding checks pass; scope/Git checks pass |

### 4. Review-gate evidence

| gate | gate condition | result |
|---|---|---|
| Migration-manifest gate | T01 manifest accepted before PRODUCT spec edits | satisfied — T01 done before T02 started |
| Semantic-batch gate (T03) | T03 review accepted before relocation | satisfied — review gate passed; T07 relocation after T03 done |
| Semantic-batch gate (T04) | T04 review accepted before relocation | satisfied — review gate passed; T07 relocation after T04 done |
| Semantic-batch gate (T05) | T05 review accepted before relocation | satisfied — review gate passed; T07 relocation after T05 done |
| Semantic-batch gate (T06) | T06 review accepted before relocation | satisfied — authoring-standards and spec-format accepted before T07 |
| Relocation gate | T07 relocation reviewed separately from semantic rewrites | satisfied — T07 was a pure relocation task with no semantic rewrites |
| DRMCP handoff gate | T08 reviewed separately | satisfied — T08 confirmed; DRMCP-local content transferred; PRODUCT specs use pointers only |
| BPDSL staging gate | T09 reviewed separately | satisfied — T09 confirmed non-canonical framing, migration trigger, and exit obligation |
| Mechanical-ref gate | T10 reviewed separately from semantic and relocation changes | satisfied — T10 was mechanical ref-only synchronization |
| T11 validation gate | strict validators and combined graph pass | satisfied — T11 Evidence: all three roots 0 errors, combined graph clean |
| T12 independent-review gate | independent review of final tree | satisfied — T12 NEEDS REVISION; M1–M5 identified; Work Item clear to close after T13 |
| T13 correction gate | M1–M5 applied; validators re-run; graph re-verified | satisfied — T13 Evidence §2–§7 |

### 5. T12 finding disposition

**Required corrections applied by T13**:

| finding | severity | result |
|---|---|---|
| M1 — Unresolved active cross-owner ref in `agent-authoring-policy.md` | MUST-FIX | resolved — both occurrences of `spec:drmcp.design_records_mcp.tools` → `spec:drmcp.design_records_mcp.tools.overview` |
| M2 — T08 handoff section headings in `traceability-boundary.md` and `record-discovery-paths.md` | MUST-FIX | resolved — both `## T08 handoff` headings renamed `## DRMCP boundary`; tables rewritten with canonical-owner pointers |
| M3 — Task vehicles as current owners in five Design Records spec files | MUST-FIX | resolved — all seven occurrences replaced with canonical-owner text (`DRMCP app-local specifications`, `BPDSL app-local specifications`, or `spec:product.bpdsl`) |
| M4 — Task vehicles in Brewprint and traceability current wording (four files) | MUST-FIX | resolved — treatment cells, Topics summary, What this is text, and obsolete-heading renamed |
| M5 — Malformed three-column separators under two-column headers in `namespace-model/app-namespaces.md` | MUST-FIX | resolved — both `\|---\|---\|---\|` rows changed to `\|---\|---\|` |

**Optional advisories — not applied**:

| advisory | disposition |
|---|---|
| A1 — Illustrative tools.md example uses nonexistent path and same ref as M1 | Not applied. Non-blocking illustrative example. |
| A2 — Hyphen-form example ref `spec:trace.semantic-ref` in Boundary table | No action. Intentional obsolete illustration. |
| A3 — Migration-provenance wording in current specs (T04, T05 references) | Not applied. Non-blocking provenance wording. |
| A4 — DRMCP tools area is an Overview not an Index | No action. DRMCP-side authoring choice; outside PRODUCT scope. |

### 6. Final validation

| check | result |
|---|---|
| PRODUCT strict validation | `[strict]  All 47 file(s) OK.` — exit 0 |
| DRMCP strict validation | `[strict]  All 30 file(s) OK.` — exit 0 |
| BPDSL strict validation | `[strict]  All 37 file(s) OK.` — exit 0 |
| Combined graph — 114 indexed specs | pass |
| Combined graph — duplicate IDs | 0 — pass |
| Combined graph — unresolved parents | 0 — pass |
| Combined graph — unresolved active body refs | 0 active after false-positive classification — pass |
| Stale `spec:product.concepts.*` refs in active specs | 0 active — one historical source-map row in `bpdsl/design-flow.md:97` classified as T11-accepted historical evidence |
| Stale `product/records/spec/concepts/` physical paths in active specs | 0 active — all residuals in BPDSL source-map preservation rows (historical evidence) |
| `concepts/` audit | no `.md` files present — confirmed empty shell |
| Staged diff | nothing staged |
| `v01/` diff | unchanged |

### 7. Completion-condition matrix

| completion condition | status | evidence source | qualification |
|---|---|---|---|
| `product/records/spec/index.md` routes placement without duplicating child contracts. | satisfied | T02 Evidence + T11 tree audit; file confirmed present with three-area routing and no child-contract restatement | none |
| `design-records/`, `brewprint/`, and temporary `bpdsl/` have explicit overview contracts. | satisfied | T02 Evidence: all three area overviews created; `bpdsl/index.md` has non-canonical contract | none |
| No active PRODUCT spec remains under `product/records/spec/concepts/`. | satisfied | T11 §8 + T13 tree check: `concepts/` is an empty shell with zero `.md` files | none |
| App-independent Design Records semantics reside under `design-records/`. | satisfied | T03–T07 semantic batches; T11 area review; T12 area-review result "architecture correct" | M1–M3 local-prose defects resolved by T13 |
| Brewprint current-state and registry material resides under `brewprint/`. | satisfied | T03 Evidence + T12 Brewprint area result "architecture correct" | M4 Boundary row defects resolved by T13 |
| V01 compatibility, historical attribution, issued-ID retention, and migration state reside under `brewprint/compatibility/`. | satisfied | T03 Evidence: `brewprint/compatibility/` created with correct content; T11 tree confirms placement | none |
| Temporary PRODUCT-side BPDSL staging states its non-canonical status, allowed content, prohibited work, review boundary, migration trigger, and exit condition. | satisfied | T09 Evidence + T12 BPDSL area result "architecture correct"; `bpdsl/index.md` confirmed to contain all required contract rows | none |
| PRODUCT Design Records specs do not restate DRMCP operational behavior or canonical BPDSL internals. | satisfied | T08 handoff; T12 dependency-direction check; M1 resolved by T13 (active authority ref now resolves) | M1 was the only active unresolved cross-owner ref; corrected |
| Unadopted future integration material is transferred to appropriate follow-up artifacts or removed after evidence transfer. | satisfied | T05 Evidence: `yaml:`, `internal-design:`, `coverage:`, and fixture endpoints removed; rationale preserved via V01-ADR-088, V01-INV-DOCS-002, V01-INV-DOCS-003 | none |
| Semantic rewrites, relocation, app-local handoff, broad ref synchronization, and validation cleanup remain separately reviewable. | satisfied | T03–T05 semantic batches; T06–T07 relocation; T08 DRMCP handoff; T09 BPDSL staging; T10 ref sync; T11 validation — all separate tasks | none |
| Mechanical refs, parent markers, and topic tables match the accepted final paths. | satisfied | T10 Evidence: mechanical synchronization complete; T11 §5–§7 stale-ref audit; T13 combined graph 0 unresolved active | M1 active unresolved ref corrected by T13 |
| Validation reports no unresolved references or structural diagnostics caused by this migration. | satisfied | T11 and T13 strict validators: 0 errors all three roots; combined graph: 0 active unresolved refs; M5 Markdown structural defect corrected by T13 | Validator does not detect Markdown table-column mismatches (T12 finding); M5 corrected in T13 |
| Pre-existing diagnostics are separated from migration-caused diagnostics. | satisfied | T11 §10: four pre-existing/false-positive diagnostics classified; T13 graph check: five false positives classified (A1 ×2, A2 ×2, T11-accepted historical ×1) | none |
| `v01/` remains unchanged. | satisfied | T11 §11 Git evidence + T13 `git diff --name-status -- v01` empty output | none |
| Completion evidence lists exact files, review gates, and validation results for every task. | satisfied | This Work Item Evidence §3 (child-task matrix T01–T13), §4 (review-gate matrix), §6 (validation), §7 (this matrix) | T13 owns final closure verification |

### 8. Deferred and out-of-scope boundary

This closure does not perform:

- DRMCP redesign or DRMCP spec rewrite.
- Full BPDSL migration or BPDSL hierarchy normalization.
- BPDSL hierarchy normalization or final BPDSL ownership decisions.
- New Design Records-to-BPDSL integration design.
- New parser, resolver, render, generation, runtime, UI, or MCP contracts.
- A1–A4 optional cleanup.
- `v01/` modification.

### 9. Final scope evidence

| check | result |
|---|---|
| T13-edited files | 14 files (listed in T13 Evidence §1) |
| No file relocation | confirmed — all T13 changes are in-place content edits |
| Nothing staged | `git diff --cached --name-status` empty |
| No commit | confirmed |
| `v01/` unchanged | `git diff --name-status -- v01` empty |

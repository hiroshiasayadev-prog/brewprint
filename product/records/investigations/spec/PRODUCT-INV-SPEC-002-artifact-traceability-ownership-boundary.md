# PRODUCT-INV-SPEC-002: Artifact and traceability ownership boundary

- **id**: PRODUCT-INV-SPEC-002
- **status**: concluded
- **date**: 2026-06-10
- **requirement_refs**:
  - PRODUCT-REQ-SPEC-001
- **source_work_items**:
  - PRODUCT-WORK-SPEC-001

## Question

Which artifact and traceability spec sections are PRODUCT-owned cross-app semantics, DRMCP-owned tool contracts, or hybrid content that needs split or relocation before spec migration?

## Scope

| area | in scope |
|---|---|
| traceability specs | Inspect `product/records/spec/concepts/traceability/**` section by section. |
| project artifact model | Inspect `product/records/spec/concepts/project-artifact-model/index.md`. |
| ownership classification | Classify sections as PRODUCT-owned semantics, DRMCP-owned tool contract, or hybrid. |
| relocation candidates | Identify files or sections that may need relocation before migration. |
| migration prerequisites | Produce recommendations needed before PRODUCT-WORK-SPEC-004 and PRODUCT-WORK-SPEC-005. |

## Non-scope

| area | owner |
|---|---|
| actual relocation | PRODUCT-WORK-SPEC-004 / PRODUCT-WORK-SPEC-005 |
| format contract changes | PRODUCT-WORK-SPEC-001 or successor work |
| DRMCP implementation | DRMCP work items |
| bulk migration | PRODUCT-WORK-SPEC-005 |

## Expected evidence

| evidence | purpose |
|---|---|
| section-level classification table | Shows PRODUCT / DRMCP / hybrid ownership. |
| relocation candidate list | Feeds PRODUCT-WORK-SPEC-004. |
| ambiguity list | Captures unresolved ownership questions. |
| migration risk notes | Prevents migration before boundary decisions are made. |

## Done condition

| item | done when |
|---|---|
| classification complete | Target files/sections are classified. |
| relocation candidates identified | Candidate files/sections and reasons are listed. |
| decision handoff ready | PRODUCT-WORK-SPEC-004 can decide relocation plan from the investigation. |
| no implementation | No implementation or bulk migration is performed. |

## Source records

| ref | role |
|---|---|
| PRODUCT-REQ-SPEC-001 | Requirement for MCP-readable spec format and topic tree support. |
| PRODUCT-INV-SPEC-001 | Found ownership-sensitive traceability concepts as migration dependency. |
| PRODUCT-WORK-SPEC-001 | Defines this investigation as a migration gate. |

## Evidence

### Files reviewed

| file | reviewed scope |
|---|---|
| `product/records/spec/concepts/project-artifact-model/index.md` | Full document. |
| `product/records/spec/concepts/traceability/index.md` | Full document. |
| `product/records/spec/concepts/traceability/artifact-refs.md` | Full document. |
| `product/records/spec/concepts/traceability/coverage-mapping.md` | Full document. |
| `product/records/spec/concepts/traceability/metadata-schema.md` | Full document. |
| `product/records/spec/concepts/traceability/out-of-scope.md` | Full document. |
| `product/records/spec/concepts/traceability/resolve-and-validation.md` | Full document. |
| `product/records/spec/concepts/traceability/semantic-ref.md` | Full document. |
| `PRODUCT-REQ-SPEC-001` | Requirement and investigation expectations. |
| `PRODUCT-INV-SPEC-001` | Migration feasibility and ownership-sensitive finding. |
| `PRODUCT-WORK-SPEC-001` | Accepted format contract and ownership boundary matrix. |
| `PRODUCT-WORK-SPEC-004` | Downstream relocation decision work item. |
| `PRODUCT-WORK-SPEC-005` | Downstream migration / restructuring work item. |
| `V01-ADR-081` | Requirement / spec traceability basis. |
| `V01-ADR-083` | Project artifact boundary basis. |
| `V01-ADR-084` | Original semantic trace MVP boundary. |
| `V01-ADR-087` | DRMCP investigation / resolver responsibility. |
| `V01-ADR-088` | Reduction to canonical reference resolution foundation. |
| `V01-ADR-091` | Workflow artifact boundary. |
| `V01-ADR-092` | DRMCP workflow record / relation boundary. |

### Classification summary

| target | ownership | reason | relocation impact |
|---|---|---|---|
| `project-artifact-model/index.md` | PRODUCT | Defines cross-app artifact classes, source-of-truth boundaries, design/change/trace flows, and artifact responsibility rules. | Keep under PRODUCT. Do not relocate as DRMCP spec. |
| `traceability/index.md` | PRODUCT | Defines canonical reference resolution foundation scope and states that DRMCP is a tool boundary, not traceability owner. | Keep under PRODUCT. May migrate as `Overview` with `## Topics`. |
| `traceability/artifact-refs.md` | PRODUCT | Defines active / reserved / deferred reference kinds and ID-as-ref semantics. These are shared identity rules. | Keep under PRODUCT. DRMCP may implement resolver support separately. |
| `traceability/coverage-mapping.md` | PRODUCT | Defines deferred realization mapping and reintroduction triggers. This is cross-app trace scope governance. | Keep under PRODUCT. No DRMCP relocation. |
| `traceability/metadata-schema.md` | PRODUCT, with DRMCP-adjacent boundaries | Defines trace metadata semantics and allowed canonical references. It explicitly leaves complete parser / response schema and diagnostic category detail to DRMCP. | Keep under PRODUCT. Avoid adding concrete DRMCP response schema here. |
| `traceability/out-of-scope.md` | PRODUCT | Defines future extension boundaries for internal-design, coverage, YAML, fixture, workflow, and writer tools. | Keep under PRODUCT. No DRMCP relocation. |
| `traceability/resolve-and-validation.md` | hybrid, PRODUCT-led | Defines canonical resolver input, lookup sources, section lookup, duplicate / unresolved conditions, and validation scope. It delegates concrete request / response fields and status vocabulary to DRMCP tools spec. | Keep semantic rules under PRODUCT. Split only concrete tool API / diagnostic response details into DRMCP if expanded. |
| `traceability/semantic-ref.md` | PRODUCT | Defines semantic ref identity, grammar, document/section refs, stability, and redirect/superseded reservation. | Keep under PRODUCT. Needed before DRMCP resolver behavior. |

### Section-level classification

| file / section group | classification | owner | notes |
|---|---|---|---|
| `project-artifact-model/index.md` / purpose, artifact classes, responsibility matrix, design/change flows | PRODUCT-owned semantics | PRODUCT | Cross-app artifact model and source-of-truth boundary. |
| `project-artifact-model/index.md` / traceability and tool boundary | PRODUCT-owned architecture with DRMCP-adjacent references | PRODUCT | States where DRMCP sits as a tool boundary, not as owner of canonical reference semantics. No split is needed unless concrete DRMCP API schema is added. |
| `project-artifact-model/index.md` / detail specifications and guide links | PRODUCT-owned navigation | PRODUCT | Contains stale v01/docs paths; migration cleanup, not relocation. |
| `traceability/index.md` / purpose, MVP scope, exclusions, terms, source-of-truth boundary | PRODUCT-owned semantics | PRODUCT | Defines traceability meaning model and says DRMCP implements the tool boundary. |
| `traceability/index.md` / split spec table | PRODUCT-owned navigation | PRODUCT | Convert to `## Topics` during migration if needed. |
| `artifact-refs.md` / active, reserved, deferred prefixes and ID-as-ref | PRODUCT-owned identity semantics | PRODUCT | Shared reference vocabulary; DRMCP implements lookup. |
| `coverage-mapping.md` / all sections | PRODUCT-owned semantics | PRODUCT | Deferred realization mapping and reintroduction triggers. |
| `metadata-schema.md` / trace metadata, `semantic_refs`, `sections`, front matter | PRODUCT-owned metadata semantics | PRODUCT | Stable metadata semantics, not DRMCP response schema. |
| `metadata-schema.md` / workflow and investigation reference metadata | hybrid, PRODUCT-led | PRODUCT + DRMCP reference | PRODUCT owns the concrete bidirectional integrity rule statements; DRMCP owns parser / response / diagnostic details. |
| `metadata-schema.md` / validation responsibility | hybrid, PRODUCT-led | PRODUCT + DRMCP reference | PRODUCT owns invalid conditions; DRMCP owns diagnostic category and JSON shape. |
| `out-of-scope.md` / all sections | PRODUCT-owned scope boundary | PRODUCT | Future extension exclusions and triggers. |
| `resolve-and-validation.md` / resolve and resolver input | hybrid, PRODUCT-led | PRODUCT + DRMCP reference | PRODUCT owns canonical input semantics; DRMCP may expose them through tool API. |
| `resolve-and-validation.md` / lookup sources and section anchor lookup | hybrid, PRODUCT-led, stable today | PRODUCT + DRMCP reference | Tool behavior is described through PRODUCT-owned metadata; guard against accidental API-schema expansion. |
| `resolve-and-validation.md` / duplicate detection, unresolved refs, and declared relation integrity validation | hybrid, PRODUCT-led, drift-sensitive | PRODUCT + DRMCP reference | PRODUCT owns invalid conditions; DRMCP owns diagnostic category names, JSON shape, and tool response vocabulary. |
| `resolve-and-validation.md` / resolver output | DRMCP-owned tool contract pointer | DRMCP | Current section correctly delegates concrete fields to DRMCP tools spec. |
| `resolve-and-validation.md` / MCP writer contract placeholder | DRMCP-owned future contract pointer | DRMCP | Placeholder should not become PRODUCT-owned API schema. |
| `semantic-ref.md` / all sections | PRODUCT-owned semantics | PRODUCT | Shared semantic ref identity, grammar, stability, and redirect reservation. |

### Relocation candidates

| candidate | recommended action | reason | target owner |
|---|---|---|---|
| `project-artifact-model/index.md` whole file | Do not relocate. | Cross-app artifact model and source-of-truth boundaries are PRODUCT semantics. | PRODUCT |
| `traceability/**` whole spec set | Do not relocate wholesale. | Traceability meaning model is cross-app semantics. Moving it to DRMCP would make one tool own shared identity rules. | PRODUCT |
| `traceability/resolve-and-validation.md` concrete resolver API details | Split only if expanded beyond current pointer. | Current file defines semantic resolution/validation boundaries. Concrete request/response/status fields belong in DRMCP tools spec. | DRMCP for API shape; PRODUCT for semantic boundary |
| `traceability/metadata-schema.md` concrete parser / diagnostic response details | Split only if concrete schema is added. | Current file defines canonical metadata semantics. Complete parser/response schema and diagnostic categories are explicitly DRMCP-owned. | DRMCP for tool schema; PRODUCT for metadata semantics |
| `traceability/index.md` split spec table | Convert during format migration, not ownership relocation. | Topic navigation and scope entrypoint content. | PRODUCT |
| `project-artifact-model/index.md` stale `docs/...` paths and guide links | Cleanup during migration. | Stale v01/docs paths are not evidence of DRMCP ownership. | PRODUCT migration work |
| DRMCP pointer targets referenced from PRODUCT specs | Verify and update during PRODUCT-WORK-SPEC-004 or a child task. | PRODUCT specs point to legacy `docs/spec/design-records-mcp/...` targets; the relocation plan must confirm their current DRMCP-side destination. | PRODUCT-WORK-SPEC-004 planning with DRMCP follow-up if needed |

### Ambiguities and required decisions

| ambiguity | impact | owner to decide |
|---|---|---|
| PRODUCT semantic validation conditions vs DRMCP diagnostic categories | PRODUCT specs may accidentally define tool response schema. | PRODUCT-WORK-SPEC-004, with DRMCP follow-up. |
| Stable `spec:` grammar currently uses hyphen word separators, while PRODUCT-WORK-SPEC-001 path-derived spec IDs use underscore segments. | Migration validation may report false drift or break stable ref compatibility. | PRODUCT-WORK-SPEC-002. |
| Whether `resolve-and-validation.md` should migrate as PRODUCT `Contract` or as PRODUCT `Concept` / `Reference` split. | A `Contract` label may make future agents treat it as DRMCP API contract. | PRODUCT-WORK-SPEC-004 / PRODUCT-WORK-SPEC-005. |
| How much workflow relation metadata belongs in PRODUCT traceability vs workflow authoring guides. | Traceability specs could duplicate authoring guide contracts. | PRODUCT-WORK-SPEC-003 / 004. |
| Whether `Overview` with `## Topics` remains the migration target for `traceability/index.md`. | A pure `Index` migration would force unnecessary split. | PRODUCT-WORK-SPEC-005. |
| Stale front-matter `scope:` paths still point at legacy `docs/spec/...` locations. | Path-derived spec IDs and migration cleanup may drift if treated as ownership evidence. | PRODUCT-WORK-SPEC-002 / PRODUCT-WORK-SPEC-005. |
| Legacy DRMCP pointer destinations are not verified by this investigation. | PRODUCT-WORK-SPEC-004 may need a dependency-update step before final relocation planning. | PRODUCT-WORK-SPEC-004, with DRMCP follow-up if needed. |

### Migration risk notes

| risk | mitigation |
|---|---|
| Moving traceability specs under DRMCP would invert ownership. | Keep PRODUCT as owner of canonical reference semantics; let DRMCP own implementation and tool contracts. |
| Treating every validation rule as DRMCP-owned would make shared artifact semantics tool-specific. | Keep invalid conditions and scope boundaries in PRODUCT; define diagnostic JSON shape in DRMCP. |
| Treating `resolve-and-validation.md` as a DRMCP tool contract wholesale would blur semantic boundary and API boundary. | Split only concrete request/response/status wording if it grows. |
| Rewriting stable `spec:` refs during migration would violate append-only stability. | Preserve existing refs until PRODUCT-WORK-SPEC-002 defines alias / redirect / compatibility behavior. |
| Cleaning stale `docs/...` paths could be mistaken for ownership relocation. | Treat stale paths as migration cleanup, not ownership evidence. |
| Current MCP validation may not fully support PRODUCT namespace records. | Do not block PRODUCT ownership decisions on current DRMCP implementation gaps. |
| Drift-sensitive hybrid sections may accumulate DRMCP diagnostic vocabulary over time. | Add active drift guards for duplicate detection, unresolved refs, and declared relation integrity validation in PRODUCT-WORK-SPEC-004. |
| Stale `scope:` paths and legacy DRMCP pointers may be confused with ownership relocation. | Treat them as cleanup / dependency-update work, not evidence that semantic ownership moves to DRMCP. |

### Review advisories carried forward

| advisory | handling |
|---|---|
| `project-artifact-model/index.md` traceability/tool-boundary content is PRODUCT architecture, not a split candidate. | PRODUCT-WORK-SPEC-004 should record this as no-split-needed DRMCP-adjacent content. |
| `resolve-and-validation.md` hybrid sections have different drift risk profiles. | PRODUCT-WORK-SPEC-004 should guard duplicate detection, unresolved refs, and declared relation integrity validation from accumulating DRMCP API vocabulary. |
| Workflow artifact metadata integrity rules are PRODUCT semantics. | PRODUCT-WORK-SPEC-004 should preserve the rule statements in PRODUCT and delegate only diagnostic JSON shape to DRMCP. |
| Stale `scope:` paths are explicit migration cleanup. | PRODUCT-WORK-SPEC-002 / PRODUCT-WORK-SPEC-005 should handle them as path / ID cleanup, not ownership relocation. |
| DRMCP-side pointer targets need verification. | PRODUCT-WORK-SPEC-004 or a child task should verify the current DRMCP destination for legacy `docs/spec/design-records-mcp/...` links. |

### Recommendation

No whole file in the reviewed target set should be relocated to DRMCP before migration.

PRODUCT should continue to own:

- project artifact model semantics
- stable `spec:` identity and section-ref semantics
- active / reserved / deferred reference vocabulary
- canonical reference resolution scope
- semantic validation conditions and boundaries
- future extension triggers

DRMCP should own:

- concrete resolver / validation tool request and response schemas
- diagnostic category names and JSON response shape
- parser / index implementation contracts
- MCP writer tool contracts, if introduced

`PRODUCT-WORK-SPEC-004` should therefore produce a split decision, not a relocation-heavy plan. The expected relocation plan is mostly "keep PRODUCT-owned semantics in PRODUCT, and prevent future DRMCP API details from accumulating in PRODUCT traceability specs."

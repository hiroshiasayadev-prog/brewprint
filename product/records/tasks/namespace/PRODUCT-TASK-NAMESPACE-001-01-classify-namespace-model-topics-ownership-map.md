# PRODUCT-TASK-NAMESPACE-001-01: Classify current namespace-model topics and define the target ownership map

- **id**: PRODUCT-TASK-NAMESPACE-001-01
- **status**: completed
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-NAMESPACE-001
- **source_requirement**: V01-REQ-PRODUCT-001
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:

## Goal

Produce a written classification of all namespace-model files by section and ownership tier, and record the target topic file names, stable refs, and relocation plan that T02–T05 will execute.

## Work

| area | required work |
|---|---|
| file inventory | Read all files under `product/records/spec/concepts/namespace-model/` and the two DRMCP stubs: `drmcp/records/spec/design-records-mcp/namespace-scanning.md`, `drmcp/records/spec/design-records-mcp/schema/id-normalization.md`. |
| section classification | For each section in each file, assign an ownership tier. Sections that are uniform receive exactly one tier: `PRODUCT-semantic` (permanent app/domain namespace definitions), `PRODUCT-compatibility` (issued ID retention policy, attribution policy), or `DRMCP-implementation` (parser behavior, prefix derivation, scanning, normalization). Sections that mix responsibilities are assigned `mixed — split required`, with the destination of each clause or paragraph recorded. |
| target structure | Determine new permanent topic file names and stable refs for T02–T05. Record the final target Topics table for `index.md` so T02 can write it without re-deriving it. |
| consistency check | Verify the classification is consistent with V01-ADR-096 (PRODUCT ownership of existing artifacts), V01-ADR-097 (namespace-first repository layout and records-root context; the DRMCP prefix-derivation algorithm is not fully defined by this ADR), V01-ADR-099 (unified v2 grammar). |

## Done condition

| item | done when |
|---|---|
| classification complete | Every section of every namespace-model file is assigned an ownership tier or `mixed — split required` with destination of each part recorded. |
| target structure defined | New permanent file names and stable refs are recorded for T02–T05 topics. |
| target Topics table recorded | The full replacement Topics table for `index.md` is documented in Evidence. |
| ADR consistency verified | Classification does not contradict V01-ADR-096, V01-ADR-097, V01-ADR-099. |

## Verification

- Cross-check the ownership map against the work item's `## Evidence` section initial evidence bullets.
- Confirm no DRMCP-implementation content is classified as PRODUCT-owned.

## Evidence

### Section-by-section classification

#### `index.md` (`spec:product.concepts.namespace_model`)

All sections: `PRODUCT-semantic`. No content relocation needed.

Two kinds of update needed:
- **Prose**: The "What this is" and "Current contract" sections say "v2 artifact ID grammar" — terminology-only update to "artifact ID grammar" (T06 scope).
- **Topics table**: Entries for `v2_grammar`, `v1_namespace_algorithm`, and `v1_id_grammar` are stale temporary refs — replace with permanent refs (T02).

#### `app-namespaces.md` (`spec:product.concepts.namespace_model.app_namespaces`)

Entirely `PRODUCT-semantic`. No ownership or relocation changes needed. Terminology cleanup: the PRODUCT domain namespace table cell uses "Namespace model, catalog, v2 ID grammar" — update "v2 ID grammar" to "artifact ID grammar" (T06 scope).

#### `domain-catalog.md` (`spec:product.concepts.namespace_model.domain_catalog`)

Entirely `PRODUCT-semantic`. No ownership or relocation changes needed. Terminology cleanup: the domain namespace table cell uses "Namespace model, catalog, v2 ID grammar" — update "v2 ID grammar" to "artifact ID grammar" (T06 scope).

#### `subdomain-model.md` (`spec:product.concepts.namespace_model.subdomain_model`)

Entirely `PRODUCT-semantic`. The DRMCP example section (V01-REQ-MCP-001–032 attribution context) is illustrative text explaining the subdomain concept, not a normative ownership definition. No content relocation needed.

#### `v2-grammar.md` (`spec:product.concepts.namespace_model.v2_grammar`) — `mixed — split required`

| section | ownership tier | destination |
|---|---|---|
| Grammar: REQ/WORK/INV/ADR format, TASK format, subdomain note, ADR note | `PRODUCT-semantic` | `artifact-id-grammar.md` |
| Sequence format: table, allocation scope, TASK work-sequence inheritance | `PRODUCT-semantic` | `artifact-id-grammar.md` |
| Mapping rule from existing IDs: domain prefix → app namespace table (MCP→DRMCP, DATA→BPDSL, RESOLVE→BPDSL) | `PRODUCT-compatibility` (effective attribution) | `existing-artifacts.md` |
| PRODUCT namespace handling — "V01-REQ-PRODUCT-001 remains unchanged" clause | `PRODUCT-compatibility` | `legacy-id-compatibility.md` |
| PRODUCT namespace handling — "New PRODUCT artifacts use full app-aware form" clause and table | `PRODUCT-semantic` (new-artifact ownership) | `existing-artifacts.md`, citing `spec:product.concepts.namespace_model.artifact_id_grammar` |

File disposition: retain temporarily; deletion owned by T07.

#### `v1-id-grammar.md` (`spec:product.concepts.namespace_model.v1_id_grammar`) — entirely `DRMCP-implementation`

| section | ownership tier | destination |
|---|---|---|
| "What this is" preamble (relocated from DRMCP) | `DRMCP-implementation` | `id-normalization.md` (new preamble) |
| Public ID: formula `namespace_prefix + bare_id`, examples, tool-API statement | `DRMCP-implementation` | `id-normalization.md` |
| Bare ID grammar table (kind → bare ID form) | `DRMCP-implementation` | `id-normalization.md` |
| Bare ID disambiguation note | `DRMCP-implementation` | `id-normalization.md` |

File disposition: retain temporarily; deletion owned by T07.

#### `v1-namespace-algorithm.md` (`spec:product.concepts.namespace_model.v1_namespace_algorithm`) — entirely `DRMCP-implementation`

| section | ownership tier | destination |
|---|---|---|
| "What this is" preamble (relocated from DRMCP) | `DRMCP-implementation` | `namespace-scanning.md` (new preamble) |
| `namespace_prefix` derivation formula + table | `DRMCP-implementation` | `namespace-scanning.md` |
| Kind-level prefix application table + parser note | `DRMCP-implementation` | `namespace-scanning.md` |
| Multi-root scan: default mode and single-root mode | `DRMCP-implementation` | `namespace-scanning.md` |

File disposition: retain temporarily; deletion owned by T07.

#### `existing-artifacts.md` (`spec:product.concepts.namespace_model.existing_artifacts`) — `mixed — split required` (in-file separation; no relocation)

| content | ownership tier | action |
|---|---|---|
| "Per V01-ADR-096, all existing artifacts are treated as owned by the PRODUCT namespace." | `PRODUCT-semantic` (historical ownership) | Retain; separate as own section |
| Table: existing ID prefix → effective app namespace (rows: REQ-MCP-*, REQ-DATA-*, etc.) | `PRODUCT-compatibility` (effective attribution) | Retain; merge with incoming attribution content from `v2-grammar.md` |
| Incoming from `v2-grammar.md`: domain prefix → app namespace mapping table (MCP→DRMCP, DATA→BPDSL, RESOLVE→BPDSL) | `PRODUCT-compatibility` (effective attribution) | Merge with attribution table above; T03 integrates both |
| "For new artifacts, use `<APP_NAMESPACE>-...` or `PRODUCT`" | `PRODUCT-semantic` (new-artifact ownership) | Retain; separate as own section |
| Incoming from `v2-grammar.md`: "New PRODUCT artifacts use PRODUCT-REQ-SPEC-001, PRODUCT-WORK-NAMESPACE-001, etc." clause + table | `PRODUCT-semantic` (new-artifact ownership) | Merge with new-artifact ownership section; cite `spec:product.concepts.namespace_model.artifact_id_grammar` for grammar details |
| End reference to `spec:product.concepts.namespace_model.v2_grammar` | stale ref | Update to `spec:product.concepts.namespace_model.artifact_id_grammar` and `spec:product.concepts.namespace_model.legacy_id_compatibility` |

File disposition: update in-place during T03; no deletion.

#### `namespace-scanning.md` (DRMCP) — hollow `DRMCP-implementation` placeholder

Current state: "relocated to `spec:product.concepts.namespace_model.v1_namespace_algorithm`". All scanning content is temporarily absent.

Restore action (T05): write full contract from `v1-namespace-algorithm.md`.

#### `id-normalization.md` (DRMCP) — hollow `DRMCP-implementation` placeholder

Current state: "relocated to `spec:product.concepts.namespace_model.v1_id_grammar`". All ID normalization content is temporarily absent.

Restore action (T04): write full contract from `v1-id-grammar.md`.

---

### Target permanent PRODUCT topics (created by T02)

| file | stable ref | tier | source content |
|---|---|---|---|
| `artifact-id-grammar.md` | `spec:product.concepts.namespace_model.artifact_id_grammar` | `PRODUCT-semantic` | Grammar + sequence format sections of `v2-grammar.md` |
| `legacy-id-compatibility.md` | `spec:product.concepts.namespace_model.legacy_id_compatibility` | `PRODUCT-compatibility` | Issued-ID retention and legacy-form compatibility clauses from `v2-grammar.md`, including the V01-* unchanged rule. |

### Target DRMCP specs restored by T04 and T05

| file | stable ref | tier | source content |
|---|---|---|---|
| `drmcp/records/spec/design-records-mcp/schema/id-normalization.md` | `spec:drmcp.design_records_mcp.schema.id_normalization` | `DRMCP-implementation` | Full content of `v1-id-grammar.md` |
| `drmcp/records/spec/design-records-mcp/namespace-scanning.md` | `spec:drmcp.design_records_mcp.namespace_scanning` | `DRMCP-implementation` | Full content of `v1-namespace-algorithm.md` |

### Files with no ownership or relocation change

| file | stable ref |
|---|---|
| `app-namespaces.md` | `spec:product.concepts.namespace_model.app_namespaces` |
| `domain-catalog.md` | `spec:product.concepts.namespace_model.domain_catalog` |
| `subdomain-model.md` | `spec:product.concepts.namespace_model.subdomain_model` |

### Target Topics table for `index.md` (T02 writes this)

| title | kind | ref | summary |
|---|---|---|---|
| App namespaces | Reference | `spec:product.concepts.namespace_model.app_namespaces` | The three app namespaces (DRMCP, BPDSL, PRODUCT), their architecture overview, and domain namespace assignments. |
| Domain catalog | Reference | `spec:product.concepts.namespace_model.domain_catalog` | Canonical domain namespace catalog and existing artifact prefixes outside the catalog. |
| Subdomain model | Reference | `spec:product.concepts.namespace_model.subdomain_model` | Subdomain grouping model, representation, write-time advisory, and example. |
| Artifact ID grammar | Reference | `spec:product.concepts.namespace_model.artifact_id_grammar` | Canonical artifact ID grammar, sequence format, allocation scopes, and canonical-ref rule. |
| Legacy ID compatibility | Reference | `spec:product.concepts.namespace_model.legacy_id_compatibility` | Issued V01-* legacy ID families, retention policy, and non-canonical status for new records. |
| Existing artifacts | Reference | `spec:product.concepts.namespace_model.existing_artifacts` | Historical ownership under V01-ADR-096, effective app attribution, and new-artifact ownership policy. |

Entries removed from current table: `v2 grammar`, `v1 namespace resolution algorithm`, `v1 record ID grammar`.

---

### ADR consistency check

| ADR | key decision | consistent? |
|---|---|---|
| V01-ADR-096 | All existing artifacts → PRODUCT owned; per-app migration not performed; existing IDs unchanged | Yes. Historical ownership (`PRODUCT-semantic`) and effective attribution (`PRODUCT-compatibility`) are correctly separated. No ID changes. |
| V01-ADR-097 | App namespace-first repository layout; `<app-namespace>/records/`; actual migration deferred | Yes. The ADR establishes the directory naming convention. The `strings.ToUpper(appNamespaceDir) + "-"` prefix derivation formula is a DRMCP parser implementation detail, not a PRODUCT-semantic contract. Classifying `v1-namespace-algorithm.md` as `DRMCP-implementation` is consistent with the ADR's scope. |
| V01-ADR-099 | All existing IDs get `V01-` prefix; new artifacts use v2 grammar; ADR numbering unified | Yes. Grammar rules are `PRODUCT-semantic`. The mapping rule (logical attribution for grouping and display) is `PRODUCT-compatibility`. The V01-* epoch prefix convention for issued IDs is captured in the compatibility policy. No contradictions. |

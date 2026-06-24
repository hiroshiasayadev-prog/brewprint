# PRODUCT-TASK-SPEC-012-05: Reconcile traceability and extract future material

- **id**: PRODUCT-TASK-SPEC-012-05
- **status**: done
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 2d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-02
- **outputs**:
  - Cleaned Design Records traceability content
  - Follow-up artifacts or evidence disposition for unadopted mechanisms

## Goal

Keep active Design Records reference semantics while removing stale format assumptions and unadopted integration mechanisms.

## Work

Process this semantic batch:

- `traceability/index.md`
- `traceability/artifact-refs.md`
- `traceability/coverage-mapping.md`
- `traceability/metadata-schema.md`
- `traceability/out-of-scope.md`
- `traceability/resolve-and-validation.md`
- `traceability/semantic-ref.md`

Keep active `spec:` and design-record ID-as-ref semantics.
Reconcile metadata wording with the accepted visible spec-format contract.
Separate canonical invalid conditions from concrete DRMCP request and response behavior.
Transfer inactive `yaml:` refs, internal-design endpoints, coverage artifacts, and unrealized relations to suitable follow-up artifacts.
Remove obsolete or duplicated statements after evidence transfer.
Do not create a `deferred-integration/` spec area.

## Done condition

- Active traceability contracts contain only current semantics.
- Spec metadata assumptions match accepted spec-format policy.
- Concrete DRMCP resolver or writer behavior is removed from PRODUCT ownership.
- Every unadopted mechanism has a follow-up destination or deletion rationale.
- The seven-file semantic batch is independently reviewable.

## Verification

- Compare each section against `PRODUCT-INV-SPEC-005`.
- Confirm no inactive integration mechanism remains normative.
- Confirm no new Design Records-to-BPDSL integration rule is introduced.
- Confirm no relocation or broad ref synchronization occurs.

## Evidence

### Execution summary

| item | result |
|---|---|
| Seven traceability source files processed | Yes. |
| Additional T03/T04 preserved-marker locations processed | Yes. |
| New follow-up investigation created | No. Existing V01 evidence and T08/T09 handoffs were sufficient. |
| Permanent future-mechanism catalog created | No. |
| Generic traceability files relocated | No. |
| DRMCP, app-local BPDSL, and `v01/**` modified | No. |
| Broad ref synchronization performed | No. |
| Remaining objective contradictions | None found. |

### Seven source files processed

| file | final result |
|---|---|
| `product/records/spec/concepts/traceability/index.md` | Rewritten as active overview. |
| `product/records/spec/concepts/traceability/artifact-refs.md` | Rewritten as active reference-class boundary. |
| `product/records/spec/concepts/traceability/coverage-mapping.md` | Removed. |
| `product/records/spec/concepts/traceability/metadata-schema.md` | Rewritten around visible metadata and relation boundaries. |
| `product/records/spec/concepts/traceability/out-of-scope.md` | Removed. |
| `product/records/spec/concepts/traceability/resolve-and-validation.md` | Rewritten as semantic lookup and invalid-condition boundary. |
| `product/records/spec/concepts/traceability/semantic-ref.md` | Rewritten to align with path-derived spec refs. |

### Additional preserved-marker locations processed

| file | final result |
|---|---|
| `product/records/spec/concepts/namespace-model/index.md` | Registry-placement alternatives deleted from current contract after evidence transfer. |
| `product/records/spec/concepts/project-artifact-model/index.md` | Future mechanisms dispositioned as existing evidence owner, BPDSL staging, T08 handoff, or delete. |
| `product/records/spec/concepts/project-artifact-model/artifact-responsibility-matrix.md` | Matrix-local pending pointers replaced with final dispositions. |
| `product/records/spec/concepts/project-artifact-model/change-and-investigation-flow.md` | Implementation endpoint tracking replaced with final dispositions. |
| `product/records/spec/concepts/project-artifact-model/traceability-boundary.md` | Deferred traceability mechanisms replaced with final dispositions. |
| `product/records/spec/bpdsl/design-flow.md` | Integration claims marked not adopted by PRODUCT and preserved for T09 classification. |
| `product/records/spec/bpdsl/artifact-responsibilities.md` | Integration claims marked not adopted by PRODUCT and preserved for T09 classification. |
| `product/records/spec/bpdsl/repository-implementation-flow.md` | Integration claims marked not adopted by PRODUCT and preserved for T09 classification. |

### Section disposition map

| source file | H2 section | final disposition |
|---|---|---|
| `traceability/index.md` | `## What this is` | rewrite current |
| `traceability/index.md` | `## Current contract` | rewrite current |
| `traceability/index.md` | `## Purpose` | rewrite current |
| `traceability/index.md` | `## MVP scope` | rewrite current |
| `traceability/index.md` | `## Out of MVP scope` | delete |
| `traceability/index.md` | `## Terms` | rewrite current |
| `traceability/index.md` | `## Topics` | rewrite current |
| `traceability/index.md` | `## Source of truth boundary` | rewrite current |
| `traceability/index.md` | `## Sources` | rewrite current |
| `traceability/artifact-refs.md` | `## What this is` | rewrite current |
| `traceability/artifact-refs.md` | `## Purpose` | rewrite current |
| `traceability/artifact-refs.md` | `## Active prefixes` | rewrite current |
| `traceability/artifact-refs.md` | `## Reserved prefixes` | BPDSL staging / delete |
| `traceability/artifact-refs.md` | `## Deferred prefixes` | existing evidence owner |
| `traceability/artifact-refs.md` | `## ID-as-ref` | rewrite current |
| `traceability/artifact-refs.md` | `## Design record ID-as-ref boundary` | rewrite current |
| `traceability/artifact-refs.md` | `## Workflow ID-as-ref boundary` | rewrite current |
| `traceability/artifact-refs.md` | `## COV-*` | existing evidence owner |
| `traceability/artifact-refs.md` | `## Relation endpoint boundary` | existing evidence owner |
| `traceability/artifact-refs.md` | `## Scope-out prefixes` | delete / existing evidence owner |
| `traceability/artifact-refs.md` | `## Sources` | rewrite current |
| `traceability/coverage-mapping.md` | `## What this is` | remove file |
| `traceability/coverage-mapping.md` | `## MVP realization mapping boundary` | existing evidence owner |
| `traceability/coverage-mapping.md` | `## Internal design boundary` | existing evidence owner |
| `traceability/coverage-mapping.md` | `## Deferred mechanisms` | existing evidence owner |
| `traceability/coverage-mapping.md` | `## Reintroduction triggers` | existing evidence owner |
| `traceability/coverage-mapping.md` | `## Validation boundary` | rewrite current into `resolve-and-validation.md` |
| `traceability/coverage-mapping.md` | `## Future artifact placement` | delete |
| `traceability/coverage-mapping.md` | `## Sources` | existing evidence owner |
| `traceability/metadata-schema.md` | `## What this is` | rewrite current |
| `traceability/metadata-schema.md` | `## Trace metadata` | rewrite current |
| `traceability/metadata-schema.md` | `## Common fields` | rewrite current |
| `traceability/metadata-schema.md` | `## semantic_refs` | delete |
| `traceability/metadata-schema.md` | `## sections` | delete |
| `traceability/metadata-schema.md` | `## Internal design metadata boundary` | existing evidence owner |
| `traceability/metadata-schema.md` | `## Coverage metadata boundary` | existing evidence owner |
| `traceability/metadata-schema.md` | `## Workflow artifact metadata boundary` | rewrite current |
| `traceability/metadata-schema.md` | `## Investigation reference metadata` | rewrite current |
| `traceability/metadata-schema.md` | `## Validation responsibility` | rewrite current |
| `traceability/metadata-schema.md` | `## Out of scope` | delete / T08 handoff |
| `traceability/out-of-scope.md` | all H2 sections | remove file |
| `traceability/resolve-and-validation.md` | `## What this is` | rewrite current |
| `traceability/resolve-and-validation.md` | `## Resolve` | rewrite current |
| `traceability/resolve-and-validation.md` | `## Resolver input` | rewrite current |
| `traceability/resolve-and-validation.md` | `## Resolver output` | T08 handoff |
| `traceability/resolve-and-validation.md` | `## Lookup sources` | rewrite current |
| `traceability/resolve-and-validation.md` | `## Section anchor lookup` | delete |
| `traceability/resolve-and-validation.md` | `## Duplicate detection` | rewrite current |
| `traceability/resolve-and-validation.md` | `## Unresolved reference and declared relation integrity` | rewrite current |
| `traceability/resolve-and-validation.md` | `## Reserved and deferred ref handling` | existing evidence owner / BPDSL staging |
| `traceability/resolve-and-validation.md` | `## Validation` | rewrite current |
| `traceability/resolve-and-validation.md` | `## Validation boundary` | rewrite current |
| `traceability/resolve-and-validation.md` | `## MCP writer contract placeholder` | T08 handoff |
| `traceability/resolve-and-validation.md` | `## Out of scope` | delete / T08 handoff |
| `traceability/semantic-ref.md` | `## What this is` | rewrite current |
| `traceability/semantic-ref.md` | `## Semantic ref definition` | rewrite current |
| `traceability/semantic-ref.md` | `## Semantic ref grammar` | delete |
| `traceability/semantic-ref.md` | `## Document-level ref and section-level ref` | rewrite current / delete section-ref claim |
| `traceability/semantic-ref.md` | `## Stability rules` | delete |
| `traceability/semantic-ref.md` | `## Redirect and superseded mapping` | delete |
| `traceability/semantic-ref.md` | `## Out of scope` | rewrite current |

### Active traceability contract retained

| active contract | retained result |
|---|---|
| Path-derived canonical `spec:` refs | Retained and pointed to `spec:product.concepts.spec_format.spec_id_as_ref`. |
| Complete public record ID-as-refs | Retained for ADR, investigation, requirement, work item, and task records. |
| Investigation canonical-reference rules | Retained for `source_refs`, recorded `follow_up_results`, and `follow_up_candidates`. |
| Workflow declared-relation integrity | Retained for requirement/work-item/task relation fields. |
| Physical paths are noncanonical | Retained. |
| PRODUCT/DRMCP ownership split | Retained with concrete DRMCP behavior excluded. |

### Spec-format reconciliation

| stale assumption | result |
|---|---|
| Spec front-matter `semantic_refs` | Removed as active contract. |
| Spec front-matter `sections` | Removed as active contract. |
| Section-heading mappings | Removed as active contract. |
| Append-only semantic refs independent of path | Removed as active contract. |
| Move/rename stability | Removed as active contract. |
| Arbitrary aliases, redirects, and superseded mapping schemas | Removed as active contract. |
| Current section-level canonical refs | Removed pending a later visible-table contract. |

### Existing evidence owners used

| evidence owner | use |
|---|---|
| V01-INV-DOCS-002 | External coverage artifact, coverage mapping, assurance matrix, and reintroduction trigger evidence. |
| V01-INV-DOCS-003 | Internal-design endpoint and realization relation deferral evidence. |
| V01-ADR-083 | Historical BPDSL/internal-design/coverage artifact boundary context. |
| V01-ADR-084 | YAML reserve-only, fixture/golden, and relation-vocabulary historical context. |
| V01-ADR-087 | Investigation canonical-reference and resolver responsibility evidence. |
| V01-ADR-088 | Canonical reference foundation and endpoint/relation reduction evidence. |
| V01-ADR-092 | Workflow ID-as-ref and declared-relation integrity boundary. |
| PRODUCT-ADR-SPEC-001 | PRODUCT semantic ownership and BPDSL staging boundary. |
| PRODUCT-INV-SPEC-005 | Semantic-layer classification and extraction guidance. |

### Removed-file transfer proof

| removed file | proof |
|---|---|
| `traceability/coverage-mapping.md` | Current validation boundary moved into `resolve-and-validation.md`; coverage artifacts, `coverage:`, `COV-*`, mapping groups, assurance matrices, and relation lifecycle material are preserved by V01-INV-DOCS-002 and V01-ADR-088. |
| `traceability/out-of-scope.md` | Current non-goals are summarized in `traceability/index.md`; internal-design and coverage mechanisms are preserved by V01-INV-DOCS-002, V01-INV-DOCS-003, and V01-ADR-088; workflow future capabilities and writer behavior are handed to T08 where implementation behavior is relevant. |

### BPDSL staging transfer

| material | result |
|---|---|
| `yaml:` endpoint and YAML realization relations | Not adopted by PRODUCT; BPDSL-specific context remains under `spec:product.bpdsl` for T09 classification. |
| Design Records-to-BPDSL integration claims | Not adopted by PRODUCT; preserved in temporary BPDSL staging for T09 classification. |
| New `product/records/spec/bpdsl/traceability-integration.md` | Not created. Existing BPDSL staging preserved the needed meaning. |

### T08 handoff items

| item | handoff |
|---|---|
| DRMCP request, response, diagnostic, parser, persistence, indexing, UI, and tool behavior | T08. |
| MCP writer tools, dry-run diff, confirmation, conflict handling, format preservation, and permission boundaries | T08. |
| Workflow orphan diagnostics, progress projection, traversal, dependency-cycle checks, and execution-order projection | T08 handoff. |

### Deletion rationales

| removed or deleted material | rationale |
|---|---|
| Future prefix catalogs | Current specs must not preserve inactive semantic prefix catalogs. |
| External relation artifacts and coverage schemas | Existing evidence owners preserve rationale; no current contract remains. |
| Fixture/golden traceability | Verification assets may exist without project-level canonical trace endpoints. |
| Mapping groups and relation lifecycle proposals | No current requirement or accepted owner. |
| Namespace registry placement alternatives | No accepted requirement chooses root-level registry files or app-local declarations. |

### Validation

Command:

```powershell
python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color
```

Exit code: 0

Exact output:

```text
[strict]  All 47 file(s) OK.
```

### Search checks

| check | result |
|---|---|
| Literal preserved-marker search | Exit code 1; no matches. |
| Stale front-matter search | Remaining matches are only obsolete/deletion evidence, not active contracts. |
| Future-catalog search | Remaining matches are current non-goals, evidence-owner rows, BPDSL staging notes, or T08 handoff rows. |
| Ownership search | Remaining DRMCP terms are exclusion or ownership-boundary statements. |

### Scoped git status evidence

`git status --short -- product/records/spec product/records/investigations/spec` returned:

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
 M product/records/spec/concepts/traceability/artifact-refs.md
 D product/records/spec/concepts/traceability/coverage-mapping.md
 M product/records/spec/concepts/traceability/index.md
 M product/records/spec/concepts/traceability/metadata-schema.md
 D product/records/spec/concepts/traceability/out-of-scope.md
 M product/records/spec/concepts/traceability/resolve-and-validation.md
 M product/records/spec/concepts/traceability/semantic-ref.md
?? product/records/investigations/spec/PRODUCT-INV-SPEC-005-product-spec-semantic-layer-and-top-level-ownership-structure.md
?? product/records/spec/bpdsl/
?? product/records/spec/brewprint/compatibility/
?? product/records/spec/brewprint/namespaces/
?? product/records/spec/design-records/
?? product/records/spec/index.md
```

`git status --short -- drmcp/records/spec bpdsl/records/spec v01` returned no output.

### Scope confirmations

| scope item | confirmation |
|---|---|
| Generic traceability relocation | Not performed. |
| DRMCP app-local specs | Unchanged. |
| App-local BPDSL specs | Unchanged. |
| `v01/**` | Unchanged. |
| Broad ref synchronization | Not performed; ref changes were local semantic corrections and local topic removal. |

### Review correction evidence

#### Review correction summary

Two issues identified by independent review and resolved:

1. Detailed future-mechanism disposition catalogs removed from normative specs. Replaced with concise current boundary or non-goal sections.
2. Conditional final-disposition wording removed. Every mechanism now has one unconditional final disposition.

#### Normative future catalogs removed

| file | old section | replacement |
|---|---|---|
| `traceability/artifact-refs.md` | `## Removed future references` (6-row disposition table) | `## Reference boundary` (3-sentence concise statement) |
| `traceability/metadata-schema.md` | `## Excluded metadata and relation mechanisms` (6-row detail table) | `## Metadata boundary` (4-sentence concise statement) |
| `traceability/resolve-and-validation.md` | `## Removed future mechanisms` (6-row detail table with conditional wording) | `## Resolve and validation boundary` (4-sentence concise statement) |
| `project-artifact-model/index.md` | `## MVP scope and future extensions` (11-row detail table with conditional wording) | `## Artifact model boundary` (4-sentence concise statement + T05 pointer) |
| `project-artifact-model/traceability-boundary.md` | `## Deferred traceability mechanisms` (6-row detail table with conditional wording) | `## Traceability contract boundary` (4-sentence concise statement) |

#### Concise current boundaries retained

| file | retained |
|---|---|
| `traceability/index.md` | `## Non-goals` concise non-goal list retained. |
| `traceability/artifact-refs.md` | All active reference class content retained. |
| `traceability/metadata-schema.md` | Spec metadata boundary, removed front-matter schemas, investigation reference metadata, and workflow relation metadata retained. |
| `traceability/resolve-and-validation.md` | All active resolve and validation sections retained. |
| `project-artifact-model/index.md` | Current contract, artifact classes, source-of-truth ownership, disposition of previous statements, and navigation retained. |
| `project-artifact-model/traceability-boundary.md` | Traceability and tool boundary, MVP scope, and T08 handoff sections retained. |

#### Unconditional T08 handoff list

All DRMCP operational and projection capabilities now carry unconditional T08 handoff disposition:

| item | file corrected |
|---|---|
| Workflow orphan diagnostics | `project-artifact-model/index.md` (via section removal); `project-artifact-model/traceability-boundary.md` (via section removal); T05 Evidence |
| Task-status-derived progress projection | `project-artifact-model/index.md` (via section removal); `project-artifact-model/traceability-boundary.md` (via section removal) |
| Workflow traversal queries | `project-artifact-model/index.md` (via section removal); `project-artifact-model/traceability-boundary.md` (via section removal) |
| Task dependency cycle or execution order projection | `project-artifact-model/index.md` (via section removal) |
| MCP writer tools | `project-artifact-model/index.md` (via section removal) |
| Implementation-update projection from Design Records | `project-artifact-model/change-and-investigation-flow.md` (conditional → unconditional) |
| Workflow orphan diagnostics, progress projection, traversal, cycle, execution-order projection (T05 Evidence row) | T05 Evidence: "T08 if later scoped as DRMCP behavior." → "T08 handoff." |

#### BPDSL staging final disposition

In `product/records/spec/bpdsl/repository-implementation-flow.md`, the statement:

> Not adopted by PRODUCT; delete from current Design Records contract unless a later accepted owner scopes it.

Changed to:

> Not adopted by PRODUCT; preserved for T09 classification.

This resolves the conditional between preservation and deletion. T09 owns final classification.

#### Confirmation: detailed disposition history exists only in task Evidence

The per-item disposition history — evidence owners, deleted mechanisms, T08 handoffs, T09 staging, removed-file transfer proof — exists only in the sections above. No detailed per-item rows remain in normative current specs.

#### Validation

Command:

```powershell
python -X utf8 product/src/tools/validate_spec.py product/records/spec --strict --no-color
```

Exit code: 0

Exact output:

```text
[strict]  All 47 file(s) OK.
```

#### Conditional-wording search result

Command:

```powershell
rg -n "if DRMCP later scopes|if later scoped|otherwise delete|unless a later accepted owner"
  product/records/spec/concepts/traceability
  product/records/spec/concepts/project-artifact-model
  product/records/spec/bpdsl
  product/records/tasks/spec/PRODUCT-TASK-SPEC-012-05-...
```

Result: No matches. Exit code 1.

#### Future-catalog heading search result

Command:

```powershell
rg -n "Removed future references|Excluded metadata and relation mechanisms|Removed future mechanisms|MVP scope and future extensions|Deferred traceability mechanisms"
  product/records/spec/concepts/traceability
  product/records/spec/concepts/project-artifact-model
```

Result: No matches in normative spec files. The only prior match (stale internal pointer `## MVP scope and future extensions` in `project-artifact-model/index.md` line 72) was also corrected to `T05 Evidence.`

Interpretation: Detailed future-mechanism catalogs are gone. Concise current boundary or non-goal sections remain under clearer headings. Full disposition history exists only in T05 Evidence.

#### `T05 pending` search result

Command:

```powershell
rg -n "T05 pending"
  product/records/spec/concepts/traceability
  product/records/spec/concepts/project-artifact-model
  product/records/spec/bpdsl
  product/records/tasks/spec/PRODUCT-TASK-SPEC-012-05-...
```

Result: No matches.

#### Forbidden-path status result

Command:

```powershell
git status --short -- drmcp/records/spec bpdsl/records/spec v01
```

Result: No output. All forbidden paths unchanged.

#### Scoped git status after review corrections

`git status --short -- product/records/spec product/records/investigations/spec` returned:

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
 M product/records/spec/concepts/traceability/artifact-refs.md
 D product/records/spec/concepts/traceability/coverage-mapping.md
 M product/records/spec/concepts/traceability/index.md
 M product/records/spec/concepts/traceability/metadata-schema.md
 D product/records/spec/concepts/traceability/out-of-scope.md
 M product/records/spec/concepts/traceability/resolve-and-validation.md
 M product/records/spec/concepts/traceability/semantic-ref.md
?? product/records/investigations/spec/PRODUCT-INV-SPEC-005-product-spec-semantic-layer-and-top-level-ownership-structure.md
?? product/records/spec/bpdsl/
?? product/records/spec/brewprint/compatibility/
?? product/records/spec/brewprint/namespaces/
?? product/records/spec/design-records/
?? product/records/spec/index.md
```

`git status --short -- drmcp/records/spec bpdsl/records/spec v01` returned no output.

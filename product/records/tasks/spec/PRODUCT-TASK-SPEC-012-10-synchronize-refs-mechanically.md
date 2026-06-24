# PRODUCT-TASK-SPEC-012-10: Synchronize refs mechanically

- **id**: PRODUCT-TASK-SPEC-012-10
- **status**: done
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-SPEC-012
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 1.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-012-08
  - PRODUCT-TASK-SPEC-012-09
- **outputs**:
  - Updated canonical refs, parent markers, topic tables, and downstream links

## Goal

Synchronize references after semantic ownership and physical paths are stable.

## Work

- Update path-derived `spec:` refs for moved PRODUCT specifications.
- Update parent markers and top-level topic tables.
- Update cross-owner pointers to PRODUCT, DRMCP, and BPDSL specs.
- Update downstream design records and supporting docs that reference old canonical refs.
- Limit changes to mechanical reference synchronization.
- Record any statement requiring semantic judgment as a separate finding.
- Do not use this task to rewrite contracts.

## Done condition

- All accepted moved specs use their new canonical refs.
- Parent markers and topic tables match the final tree.
- No active reference points to removed `product.concepts` targets.
- Cross-owner pointers resolve to the correct semantic owner.
- Semantic questions are not silently resolved during mechanical edits.

## Verification

- Search scoped Markdown files for old `spec:product.concepts` refs.
- Compare every replacement with the T01 old-to-new mapping.
- Confirm no unrelated prose changed.
- Review broad changes as a mechanical diff.

## Evidence

### 1. Execution summary

| item | result |
|---|---|
| Status | `done`; mechanical synchronization, scoped residual audit, Git staging check, and strict validation are complete. |
| Files changed | 53 files. Exact paths are listed below. |
| Canonical refs replaced | 243 stale canonical-ref occurrences. |
| Supporting physical paths replaced | 2 active fallback-path occurrences in `prompt_chappy.md`. |
| Baseline | 428 matching lines containing 462 stale canonical-ref occurrences. |
| Residual | 186 matching lines containing 219 old canonical-ref occurrences. |
| Historical evidence preserved | 217 occurrences. |
| Removed-target history | 2 occurrences; both are historical output-map rows, not live pointers. |
| Active stale refs remaining | 0. |
| Parent or Topics corrections | None. Current parent markers and authoritative Topics rows already matched the accepted tree. |

Important files inspected but unchanged:

- `product/records/spec/index.md`
- `product/records/spec/design-records/index.md`
- `product/records/spec/brewprint/index.md`
- `product/records/spec/bpdsl/index.md`
- `drmcp/records/spec/design-records-mcp/resolver.md`
- `drmcp/records/spec/design-records-mcp/schema/discovery.md`
- all current files under `bpdsl/records/**`

### 2. Authoritative mapping used

One-to-one mappings encountered:

- authoring-standards root and retained child suffixes to `spec:product.design_records.authoring_standards`;
- spec-format root and retained child suffixes to `spec:product.design_records.spec_format`;
- artifact-ID grammar and subdomain-model children to `spec:product.design_records.namespace_model`;
- record-discovery-paths to `spec:product.design_records.repository_layout.record_discovery_paths`;
- traceability root and retained children to `spec:product.design_records.traceability`;
- project-artifact-model root and retained children to `spec:product.design_records.artifact_model`.

Context-sensitive mappings encountered:

- generic namespace semantics to `spec:product.design_records.namespace_model`;
- current Brewprint namespace assignments to `spec:product.brewprint.namespaces`;
- issued-ID and historical attribution policy to `spec:product.brewprint.compatibility`;
- generic new-artifact owner selection to `spec:product.design_records.namespace_model.existing_artifacts`;
- generic record placement to `spec:product.design_records.repository_layout`;
- Design Records artifact responsibilities to `spec:product.design_records.artifact_model.artifact_responsibility_matrix`;
- preserved DSL/source flow to the existing temporary `spec:product.bpdsl` staging records.

Removed targets encountered:

- the former `traceability.coverage_mapping` topic;
- the former `traceability.out_of_scope` topic.

Cross-owner mappings encountered:

- DRMCP resolver semantics to `spec:product.design_records.traceability.resolve_and_validation`;
- DRMCP record discovery to `spec:product.design_records.repository_layout.record_discovery_paths`;
- DRMCP authoring inputs to current Design Records authoring standards;
- DRMCP overview boundary to `spec:product.design_records.artifact_model`.

### 3. Occurrence disposition summary

| occurrence class | count | treatment |
|---|---:|---|
| `active_mechanical` | 243 | Updated to accepted current canonical refs. Split-owner cases use multiple accepted refs without reopening semantic decisions. |
| `historical_evidence` | 217 | Preserved as old values, source identities, migration maps, commands, baseline output, completed-task evidence, or an explicitly historical decision statement. |
| `removed_target` | 2 | Preserved as historical source identities. No replacement spec was invented. |
| `context_sensitive` | 0 remaining | Every active case was resolved from the accepted owner tree. |
| `active_stale` | 0 | No active canonical pointer remains on a removed `product.concepts` target. |

The counts are occurrence counts, not matching-line counts. Lines containing multiple refs were counted once per ref.

### 4. Files changed by area

#### PRODUCT specs

- `product/records/spec/bpdsl/artifact-responsibilities.md`
- `product/records/spec/bpdsl/repository-implementation-flow.md`
- `product/records/spec/brewprint/compatibility/existing-artifacts.md`
- `product/records/spec/brewprint/compatibility/index.md`
- `product/records/spec/brewprint/compatibility/legacy-id-compatibility.md`
- `product/records/spec/brewprint/layout/index.md`
- `product/records/spec/brewprint/namespaces/app-namespaces.md`
- `product/records/spec/brewprint/namespaces/domain-catalog.md`
- `product/records/spec/brewprint/namespaces/index.md`
- `product/records/spec/design-records/artifact-model/artifact-responsibility-matrix.md`
- `product/records/spec/design-records/artifact-model/change-and-investigation-flow.md`
- `product/records/spec/design-records/artifact-model/index.md`
- `product/records/spec/design-records/artifact-model/traceability-boundary.md`
- `product/records/spec/design-records/authoring-standards/adr-authoring.md`
- `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`
- `product/records/spec/design-records/authoring-standards/artifact-boundary.md`
- `product/records/spec/design-records/authoring-standards/investigation-authoring.md`
- `product/records/spec/design-records/authoring-standards/requirement-authoring.md`
- `product/records/spec/design-records/authoring-standards/spec-authoring.md`
- `product/records/spec/design-records/authoring-standards/task-authoring.md`
- `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
- `product/records/spec/design-records/authoring-standards/writing-standard.md`
- `product/records/spec/design-records/namespace-model/app-namespaces.md`
- `product/records/spec/design-records/namespace-model/artifact-id-grammar.md`
- `product/records/spec/design-records/namespace-model/existing-artifacts.md`
- `product/records/spec/design-records/namespace-model/index.md`
- `product/records/spec/design-records/namespace-model/subdomain-model.md`
- `product/records/spec/design-records/repository-layout/index.md`
- `product/records/spec/design-records/repository-layout/record-discovery-paths.md`
- `product/records/spec/design-records/spec-format/document-shape.md`
- `product/records/spec/design-records/spec-format/follow-up-boundary.md`
- `product/records/spec/design-records/spec-format/index.md`
- `product/records/spec/design-records/spec-format/overview.md`
- `product/records/spec/design-records/spec-format/spec-id-as-ref.md`
- `product/records/spec/design-records/spec-format/topics-table.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`
- `product/records/spec/design-records/traceability/artifact-refs.md`
- `product/records/spec/design-records/traceability/index.md`
- `product/records/spec/design-records/traceability/metadata-schema.md`
- `product/records/spec/design-records/traceability/resolve-and-validation.md`
- `product/records/spec/design-records/traceability/semantic-ref.md`

#### PRODUCT downstream records

- `product/records/investigations/spec/PRODUCT-INV-SPEC-005-product-spec-semantic-layer-and-top-level-ownership-structure.md`
- `product/records/requirements/PRODUCT-REQ-SPEC-002-migrate-artifact-authoring-guides-to-product-namespace.md`
- `product/records/work-items/namespace/PRODUCT-WORK-NAMESPACE-001-namespace-model-canonical-grammar-and-compatibility-boundary-cleanup.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-002-path-derived-canonical-spec-refs-and-ref-first-topic-index.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-010-cognitive-load-writing-standard-for-design-records.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-011-per-artifact-authoring-guide-migration.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-012-product-spec-semantic-layer-restructuring.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-10-synchronize-refs-mechanically.md`

No PRODUCT ADR required an active-pointer edit. `PRODUCT-INV-SPEC-005` required a split-owner `source_refs` update.

#### DRMCP records

- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-002-namespace-aware-authoring-transaction-conformance.md`
- `drmcp/records/spec/design-records-mcp/overview.md`

#### BPDSL records

None. Current BPDSL records contained no stale PRODUCT refs.

#### Supporting docs

- `prompt_chappy.md`
- `CLAUDE.md`

### 5. Preserved historical evidence

Important preserved groups:

| file or group | reason preserved |
|---|---|
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-06-relocate-authoring-standards-and-spec-format.md` | Old-to-new ID map and explicit handoff evidence. Rewriting old-value columns would falsify migration history. |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-07-relocate-cleaned-semantic-areas.md` | Old-to-new area map, old tree checks, and completed command output. |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-08-apply-drmcp-app-local-handoff.md` | Before-values for the already accepted resolver and discovery updates. |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-09-finalize-temporary-bpdsl-staging.md` | Evidence that broad ref synchronization was intentionally deferred to T10. |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-009-02-format-migration.md` | Historical output manifest, including identities later removed or split. |
| older namespace and spec-format tasks | Creation-time IDs, commands, acceptance evidence, and deleted-topic identities. |
| `product/records/spec/bpdsl/design-flow.md` | Explicit previous-context row for the mixed repository-layout relation. |
| `product/records/work-items/spec/PRODUCT-WORK-SPEC-011-per-artifact-authoring-guide-migration.md` | Historical decision and completion evidence describing the former mixed artifact matrix and former v2 grammar. Rewriting those statements as current ownership claims would falsify completed-work history. |

The residual physical-path search returns 572 matching lines. These are historical inventories, source maps, command output, and migration Evidence, plus the accepted T13 advisory below. They are not live navigation links.

### 6. Removed-target findings

| file | section | old topic | disposition | accepted replacement | follow-up |
|---|---|---|---|---|---|
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-009-02-format-migration.md` | `## Work`, traceability output map | `traceability.coverage_mapping` | Preserved historical output identity. | None. | No correction required unless completed-task history policy changes. |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-009-02-format-migration.md` | `## Work`, traceability output map | `traceability.out_of_scope` | Preserved historical output identity. | None. | No correction required unless completed-task history policy changes. |

No live pointer to either removed target remains.

### 7. Cross-owner pointer synchronization

PRODUCT to DRMCP and DRMCP to PRODUCT:

- Existing T08 resolver pointers remain on `spec:product.design_records.traceability.resolve_and_validation`.
- Existing T08 discovery pointers remain on `spec:product.design_records.repository_layout.record_discovery_paths`.
- `DRMCP-REQ-MCP-002` now cites current Design Records authoring, ID grammar, and repository-layout contracts.
- DRMCP overview now cites `spec:product.design_records.artifact_model`.
- No DRMCP format, authoring-transaction, namespace-prefix, or resolver redesign was performed.

PRODUCT to BPDSL:

- Temporary PRODUCT BPDSL staging now points to current Design Records owners where a generic relation was already accepted.
- Current `bpdsl/records/**` contained no stale PRODUCT refs.
- No self-hosting pointer was added.
- No app-local BPDSL contract or Design Records-to-BPDSL integration contract was introduced.

### 8. Parent and Topics audit

No parent or authoritative Topics mismatch required correction.

The audited tree is:

```text
spec:product
├─ spec:product.design_records
│  ├─ spec:product.design_records.authoring_standards
│  ├─ spec:product.design_records.namespace_model
│  ├─ spec:product.design_records.repository_layout
│  ├─ spec:product.design_records.spec_format
│  ├─ spec:product.design_records.traceability
│  └─ spec:product.design_records.artifact_model
├─ spec:product.brewprint
│  ├─ spec:product.brewprint.layout
│  ├─ spec:product.brewprint.namespaces
│  └─ spec:product.brewprint.compatibility
└─ spec:product.bpdsl
```

All 47 current PRODUCT spec files expose a parent matching this tree or the accepted child owner.

### 9. Residual stale-ref audit

Residual old refs: 219 occurrences across 186 lines.

| disposition | occurrences | active stale pointer remaining |
|---|---:|---|
| historical evidence | 217 | no |
| removed-target history | 2 | no |
| active stale refs | 0 | no |

Final context-sensitive dispositions:

| location | disposition |
|---|---|
| `product/records/spec/brewprint/layout/index.md` | Split the former mixed repository-layout pointer into the current Design Records placement owner and temporary BPDSL implementation-flow staging owner. |
| `product/records/spec/design-records/namespace-model/index.md` | Applied the same accepted split without changing the namespace contract. |
| `PRODUCT-WORK-NAMESPACE-001` metadata | Replaced the former mixed existing-artifacts ref with current generic ownership-selection and Brewprint compatibility refs. |
| `PRODUCT-WORK-SPEC-012` metadata | Replaced former mixed area refs with current Design Records roots and added the temporary BPDSL root. |
| `PRODUCT-INV-SPEC-005` metadata | Expanded moved mixed sources to their accepted Design Records, Brewprint, and temporary BPDSL owners. |
| `CLAUDE.md` and `prompt_chappy.md` | Updated active authoring-policy and filesystem-fallback pointers. |
| `PRODUCT-WORK-SPEC-011` historical decision | Preserved unchanged as completed-work history. The statement describes the former mixed matrix, so rewriting it as a current ownership claim would be false. |

The physical-path occurrence in `product/records/spec/design-records/artifact-model/traceability-boundary.md` remains unchanged as the explicit T07-to-T13 advisory.

### 10. Scope evidence

- No `v01/**` file was read for modification or changed.
- No file was relocated.
- No staging or commit command was issued.
- No substantive contract argument was rewritten.
- No T11 diagnostic cleanup was performed.
- No T12 review or T13 correction was performed.
- The accepted T08 resolver and discovery changes were inspected and preserved.
- The four T09 PRODUCT-side BPDSL staging files were not redesigned or normalized.
- Current working-tree content, not `HEAD`, was used as the semantic baseline.
- `git diff --cached --name-status` returned no staged changes.
- Strict validation passed for PRODUCT, DRMCP, and BPDSL spec roots.
- Final scoped grep found no active workflow metadata or current spec pointer using a removed `product.concepts` target.

### 11. Final verification

| check | result |
|---|---|
| `git status --short` | Migration worktree present and consistent with T01-T10 scope; no unexpected staging action performed. |
| `git diff --cached --name-status` | Empty. |
| PRODUCT strict validator | Exit 0: all 47 files OK. |
| DRMCP strict validator | Exit 0: all 30 files OK. |
| BPDSL strict validator | Exit 0: all 37 files OK. |
| Old canonical-ref grep | `rg` was unavailable in PowerShell. Scoped grep MCP replacement completed: 186 matching lines, all classified as historical evidence or removed-target history. |
| Active workflow metadata grep | 0 matches outside task history. |
| Old physical-path audit | 572 residual matching lines after updating the two active `prompt_chappy.md` fallback paths. Remaining matches are historical inventories, source maps, command output, migration evidence, or the accepted T13 advisory. |
| Active removed-target pointers | 0. |
| Parent and Topics audit | No correction required. |

Verdict: done. All T10 completion gates are satisfied. T11 validation-diagnostic cleanup, T12 independent review, and T13 review corrections remain separate downstream tasks.

# PRODUCT-TASK-NAMESPACE-001-07: Synchronize DRMCP indexes and dependent contract references

- **id**: PRODUCT-TASK-NAMESPACE-001-07
- **status**: completed
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-NAMESPACE-001
- **source_requirement**: V01-REQ-PRODUCT-001
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-NAMESPACE-001-02
  - PRODUCT-TASK-NAMESPACE-001-03
  - PRODUCT-TASK-NAMESPACE-001-04
  - PRODUCT-TASK-NAMESPACE-001-05
  - PRODUCT-TASK-NAMESPACE-001-06
- **outputs**:
  - `drmcp/records/spec/design-records-mcp/overview.md` (updated if stale refs found)
  - `drmcp/records/spec/design-records-mcp/schema/overview.md` (updated if stale refs found)
  - `drmcp/records/spec/design-records-mcp/tools/` (affected tool specs updated)
  - `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-*.md` (updated if stale refs found)
  - `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-002-*.md` (updated if stale refs found)
  - `product/records/spec/concepts/namespace-model/v2-grammar.md` (deleted)
  - `product/records/spec/concepts/namespace-model/v1-id-grammar.md` (deleted)
  - `product/records/spec/concepts/namespace-model/v1-namespace-algorithm.md` (deleted)

## Goal

Update DRMCP spec index files, tool specs, and requirement records that reference stale PRODUCT canonical refs or carry ownership descriptions that no longer reflect the restored DRMCP contracts from T04–T05. After DRMCP reference migration is complete and T06 has cleared PRODUCT references, delete the three obsolete temporary PRODUCT namespace-model files.

Note: DRMCP API redesign is out of scope. Only canonical ref updates and ownership source corrections are in scope for this task.

## Work

| area | required work |
|---|---|
| discover stale refs | Search for references to `v1_id_grammar`, `v1_namespace_algorithm`, `namespace_model.v2_grammar` across all files in `drmcp/records/`. Also search for stale Phase 2 relocation notes ("Relocated from … per PRODUCT-WORK-SPEC-004") in `id-normalization.md` and `namespace-scanning.md` — these were restored by T04/T05 and the relocation note is now incorrect. |
| update overview files | In `drmcp/records/spec/design-records-mcp/overview.md` and `schema/overview.md`: update any refs that pointed to PRODUCT namespace-model files for parser/scanning contracts to now point to the restored DRMCP specs. |
| update tool specs | Check `drmcp/records/spec/design-records-mcp/tools/` for any tool spec that references the old relocation stubs; update to the restored DRMCP spec refs. |
| update requirements | Check `DRMCP-REQ-MCP-001-*` and `DRMCP-REQ-MCP-002-*` for stale namespace-model or ownership source refs. Update only canonical ref fields; do not change requirement scope or acceptance criteria. |
| remove stale relocation notes | In the restored `id-normalization.md` and `namespace-scanning.md`, ensure no stale "relocated to PRODUCT" relocation history note remains. |
| delete obsolete PRODUCT files | After DRMCP reference migration is complete and T06 has confirmed all PRODUCT references are clean: delete `product/records/spec/concepts/namespace-model/v2-grammar.md`, `v1-id-grammar.md`, and `v1-namespace-algorithm.md`. Before deleting each file, confirm no active machine-readable ref still points to it. If T01 recorded a retention requirement for any file, skip its deletion and document in Evidence. |
| validate | Run `python product/src/tools/validate_spec.py --strict --no-color` on all changed DRMCP files. Confirm exit 0. |

## Done condition

| item | done when |
|---|---|
| stale DRMCP refs eliminated | No DRMCP spec references `v1_id_grammar`, `v1_namespace_algorithm`, or contains stale Phase 2 relocation notes. |
| DRMCP overview updated | Overview files reference the restored DRMCP namespace-scanning and id-normalization specs directly. |
| requirements updated | Requirement records reference the correct canonical ownership sources. |
| obsolete files deleted | `v2-grammar.md`, `v1-id-grammar.md`, and `v1-namespace-algorithm.md` are deleted from `product/records/spec/concepts/namespace-model/`, or T01 retention requirements are documented in Evidence for any retained file. |
| validator clean | `validate_spec.py --strict` exits 0 on all changed DRMCP files. |

## Verification

- Run stale-ref search after edits and confirm zero unexpected matches in DRMCP files.
- Confirm DRMCP API contract content (tool schemas, response formats) is unchanged.

## Evidence

**DRMCP stale-ref search:** Searched `drmcp/records/` for `v1_id_grammar`, `v1_namespace_algorithm`, `namespace_model.v2_grammar`, "relocated to PRODUCT", and "PRODUCT-WORK-SPEC-004".

Findings:
- `DRMCP-REQ-MCP-002` had `spec:product.concepts.namespace_model.v2_grammar` in `source_refs` (line 10) and body Evidence (line 30). Both updated to `spec:product.concepts.namespace_model.artifact_id_grammar`.
- `drmcp/records/spec/design-records-mcp/schema/discovery.md`: Phase 2 relocation note present, but it describes relocation of `record_discovery_paths` content (not the namespace-model files). Out of scope for this task; no change made.
- No references to `v1_id_grammar` or `v1_namespace_algorithm` found in DRMCP records.
- Restored `id-normalization.md` and `namespace-scanning.md` contain no stale relocation notes (written fresh by T04/T05).

**Active-ref pre-deletion check:** Confirmed no active machine-readable metadata ref pointed to any of the three files at time of deletion. Remaining occurrences were only in: the files' own `id` fields, internal cross-reference within `v1-namespace-algorithm.md`, and historical body prose in completed task records (exempt).

**T01 retention check:** T01 Evidence records no retention requirement for any of the three files beyond T07.

**Deleted:**
- `product/records/spec/concepts/namespace-model/v2-grammar.md`
- `product/records/spec/concepts/namespace-model/v1-id-grammar.md`
- `product/records/spec/concepts/namespace-model/v1-namespace-algorithm.md`

**`namespace-model/` after deletion:** 7 files — `index.md`, `app-namespaces.md`, `domain-catalog.md`, `subdomain-model.md`, `artifact-id-grammar.md`, `legacy-id-compatibility.md`, `existing-artifacts.md`. No v-named temporary topics remain.

**Validation:** `validate_spec.py --strict` on `drmcp/records/spec/` — exit 0, 30 files OK.

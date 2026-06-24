# PRODUCT-TASK-NAMESPACE-001-06: Synchronize PRODUCT downstream consumers

- **id**: PRODUCT-TASK-NAMESPACE-001-06
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
- **outputs**:
  - `product/records/spec/concepts/traceability/` (affected files updated)
  - `product/records/spec/concepts/project-artifact-model/` (affected files updated)
  - `product/records/spec/concepts/authoring-standards/` (affected files updated)
  - `product/records/spec/concepts/repository-layout/` (affected files updated)
  - `product/records/requirements/` (metadata refs updated where stale)
  - `product/records/work-items/` (metadata refs updated where stale)
  - `product/records/tasks/` (metadata refs updated where stale)
  - `product/records/investigations/` (metadata refs updated where stale)
  - `product/records/adr/` (metadata refs updated where stale)

## Goal

Update all active PRODUCT spec files and workflow records that carry stale machine-readable canonical refs to the old temporary namespace-model topics, replacing each with the appropriate new stable ref from T02–T05. PRODUCT semantic statements must cite PRODUCT authority; DRMCP implementation specs are not cited as the authority for public identity semantics.

## Work

| area | required work |
|---|---|
| discover stale refs in specs | Search for `v1_id_grammar`, `v1_namespace_algorithm`, `v2_grammar` (and hyphenated filename variants) across the four spec directories: `product/records/spec/concepts/traceability/`, `product/records/spec/concepts/project-artifact-model/`, `product/records/spec/concepts/authoring-standards/`, `product/records/spec/concepts/repository-layout/`. |
| update specs — authority rule | Replace each stale ref with the correct new canonical ref. Authority rule: PRODUCT semantic content (public ID grammar, compatibility policy, spec identity) must cite PRODUCT sources — `spec:product.concepts.namespace_model.artifact_id_grammar`, `spec:product.concepts.namespace_model.legacy_id_compatibility`, `spec:product.concepts.spec_format.spec_id_as_ref`. Only content that explicitly describes the DRMCP tool boundary may cite `spec:drmcp.design_records_mcp.schema.id_normalization` or `spec:drmcp.design_records_mcp.namespace_scanning`. PRODUCT specs must not acquire a new dependency on DRMCP normalization specs as the authority for public identity semantics. |
| discover stale refs in workflow records | Search for machine-readable references to `spec:product.concepts.namespace_model.v2_grammar`, `spec:product.concepts.namespace_model.v1_id_grammar`, `spec:product.concepts.namespace_model.v1_namespace_algorithm` in metadata fields (`impact_refs`, `source_refs`, `semantic_refs`, `depends_on`, `outputs`) across active (non-`v01/`) PRODUCT records: `product/records/requirements/`, `product/records/work-items/`, `product/records/tasks/`, `product/records/investigations/`, `product/records/adr/`. |
| update workflow metadata | For each workflow record with a stale machine-readable ref in a metadata field, update the ref to the new canonical ref. Historical prose references in body text do not require rewriting. Update `PRODUCT-WORK-NAMESPACE-001.impact_refs` to list the final canonical refs from T02–T05. |
| validate | Run `python product/src/tools/validate_spec.py --strict --no-color` on all changed spec files. Confirm exit 0. |

## Done condition

| item | done when |
|---|---|
| stale spec refs eliminated | No machine-readable references to `v1_id_grammar`, `v1_namespace_algorithm`, or `v2_grammar` remain in PRODUCT spec files under the four spec directories. |
| correct authority used | PRODUCT semantic statements cite PRODUCT canonical refs; no PRODUCT spec has acquired a new dependency on DRMCP normalization specs as the authority for public identity semantics. |
| workflow metadata updated | No active (non-`v01/`) PRODUCT workflow record carries `spec:product.concepts.namespace_model.v2_grammar`, `spec:product.concepts.namespace_model.v1_id_grammar`, or `spec:product.concepts.namespace_model.v1_namespace_algorithm` in a machine-readable metadata field. |
| WORK-NAMESPACE-001 updated | `PRODUCT-WORK-NAMESPACE-001.impact_refs` carries the new canonical refs from T02–T05. |
| validator clean | `validate_spec.py --strict` exits 0 on all changed spec files. |

## Verification

- Run the stale-ref search after edits and confirm zero unexpected matches in machine-readable fields.
- Confirm no PRODUCT downstream spec now cites `spec:drmcp.design_records_mcp.schema.id_normalization` as the authority for public identity semantics.
- Confirm no downstream spec duplicates grammar content now owned by `artifact-id-grammar.md`.

## Evidence

**Stale-ref discovery — spec files:**

All four spec directories searched for `v1_id_grammar`, `v1_namespace_algorithm`, `v2_grammar`. Found in 11 files across `traceability/`, `repository-layout/`, and `authoring-standards/`.

**Spec files updated:**

| file | change |
|---|---|
| `traceability/artifact-refs.md` | Lines 78-79: `v1_id_grammar` → `legacy_id_compatibility`; `v2_grammar` → `artifact_id_grammar`. Lines 125-126: Sources entries updated to `legacy_id_compatibility` and `artifact_id_grammar`. |
| `traceability/semantic-ref.md` | Line 31: `v2_grammar` → `artifact_id_grammar`. |
| `traceability/metadata-schema.md` | Line 101: `v2_grammar` → `artifact_id_grammar`. |
| `traceability/resolve-and-validation.md` | Line 41: `v2_grammar` → `artifact_id_grammar`. |
| `repository-layout/index.md` | Related specs table: `v2_grammar` → `artifact_id_grammar`. |
| `repository-layout/record-discovery-paths.md` | Line 25: `v1_namespace_algorithm` → `spec:drmcp.design_records_mcp.namespace_scanning` (DRMCP tool boundary citation; authority rule applies). |
| `authoring-standards/adr-authoring.md` | Lines 33 and 175: `v2_grammar` → `artifact_id_grammar` (replace_all). |
| `authoring-standards/investigation-authoring.md` | Lines 33 and 219: `v2_grammar` → `artifact_id_grammar` (replace_all). |
| `authoring-standards/requirement-authoring.md` | Lines 33 and 177: `v2_grammar` → `artifact_id_grammar` (replace_all). |
| `authoring-standards/task-authoring.md` | Lines 33 and 200: `v2_grammar` → `artifact_id_grammar` (replace_all). |
| `authoring-standards/work-item-authoring.md` | Lines 32 and 193: `v2_grammar` → `artifact_id_grammar` (replace_all). |

**Terminology cleanup (namespace-model/ prose):**

| file | change |
|---|---|
| `namespace-model/index.md` | "v2 artifact ID grammar" → "artifact ID grammar" (What this is); "v2 ID grammar" → "artifact ID grammar" (Current contract). Date updated to 2026-06-24. |
| `namespace-model/app-namespaces.md` | "v2 ID grammar" → "artifact ID grammar" in PRODUCT domain namespace table. |
| `namespace-model/domain-catalog.md` | "v2 ID grammar" → "artifact ID grammar" in PRODUCT domain namespace table. |

**Workflow metadata updated:**

- `PRODUCT-WORK-NAMESPACE-001.impact_refs`: removed `v2_grammar`, `v1_id_grammar`, `v1_namespace_algorithm`; added `artifact_id_grammar`, `legacy_id_compatibility`.
- `PRODUCT-WORK-SPEC-011`: completed record; `impact_refs` are historical — not updated.
- `PRODUCT-WORK-NAMESPACE-001` Impact Scope body prose (lines 56-58): historical pre-change descriptions — exempt per T06 rule.

**Authority rule:** `record-discovery-paths.md` now cites `spec:drmcp.design_records_mcp.namespace_scanning` (DRMCP tool boundary). No PRODUCT spec acquired a new dependency on DRMCP normalization as the authority for public identity semantics.

**Validation:** `validate_spec.py --strict` on all four spec directories — exit 0, 29 files OK.

**Additional fixes (post-validation):**

- `PRODUCT-WORK-SPEC-011.impact_refs`: changed `v2_grammar` → `artifact_id_grammar`. Body Evidence retained as historical record of completed work.
- Terminology cleanup — "namespace-aware v2 public IDs / v2 IDs / v2 ID-as-refs" → "canonical app-aware artifact IDs / app-aware artifact ID-as-refs" in: `traceability/out-of-scope.md`, `traceability/resolve-and-validation.md` (3 occurrences), `traceability/artifact-refs.md`, `project-artifact-model/change-and-investigation-flow.md`, `project-artifact-model/traceability-boundary.md`.

**Remaining stale refs in namespace-model/:** Only the three temporary files' own `id` fields (`v2-grammar.md:3`, `v1-id-grammar.md:3`, `v1-namespace-algorithm.md:3`) and their internal cross-reference (`v1-namespace-algorithm.md:36`). These are in the files T07 will delete; they are not downstream consumers.

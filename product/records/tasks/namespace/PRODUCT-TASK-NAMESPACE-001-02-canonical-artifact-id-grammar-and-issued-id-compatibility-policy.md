# PRODUCT-TASK-NAMESPACE-001-02: Establish the canonical artifact-ID grammar and issued-ID compatibility policy

- **id**: PRODUCT-TASK-NAMESPACE-001-02
- **status**: completed
- **date**: 2026-06-24
- **work_item**: PRODUCT-WORK-NAMESPACE-001
- **source_requirement**: V01-REQ-PRODUCT-001
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-NAMESPACE-001-01
- **outputs**:
  - `product/records/spec/concepts/namespace-model/artifact-id-grammar.md` (new)
  - `product/records/spec/concepts/namespace-model/legacy-id-compatibility.md` (new)
  - `product/records/spec/concepts/namespace-model/v2-grammar.md` (retained temporarily; deletion owned by T07)
  - `product/records/spec/concepts/namespace-model/index.md` (Topics table rebuilt)

## Goal

Create the two permanent replacement topics, rebuild the namespace-model topic index, and defer deletion of `v2-grammar.md` to T07.

## Work

| area | required work |
|---|---|
| create grammar topic | Create `artifact-id-grammar.md` with stable ref `spec:product.concepts.namespace_model.artifact_id_grammar`. Content: (1) ID format grammar for REQ/WORK/INV/ADR and TASK; (2) sequence format table; (3) sequence allocation scope — for REQ/WORK/INV/ADR, scope is app namespace + artifact kind + domain namespace; for TASK, the work-sequence segment is inherited from the parent Work Item's three-digit sequence, and the task-sequence is allocated within the parent Work Item; (4) subdomain exclusion — subdomain labels are grouping tools and are not segments in the public ID; (5) the complete public ID is the canonical record ID-as-ref; (6) bare forms (`REQ-*`, `WORK-*`, `TASK-*`, etc.) are internal grammar fragments and are not canonical external references. |
| create compatibility topic | Create `legacy-id-compatibility.md` with stable ref `spec:product.concepts.namespace_model.legacy_id_compatibility`. Content: (1) accepted legacy public-ID families: `V01-ADR-*`, `V01-SPEC-*`, `V01-INV-*`, `V01-REQ-*`, `V01-WORK-*`, `V01-TASK-*`; (2) issued IDs remain unchanged and resolvable; (3) legacy public IDs are not the canonical form for new records — new records must use the grammar defined in `spec:product.concepts.namespace_model.artifact_id_grammar`; (4) legacy indexed `V01-SPEC-*` IDs are compatibility-only and do not constitute canonical spec identity — for canonical spec identity, see `spec:product.concepts.spec_format.spec_id_as_ref`. Do not redefine spec identity in this topic; reference only. |
| retain v2-grammar temporarily | Do not delete `v2-grammar.md` in this task. Deletion is owned by T07 after T06 completes reference migration. If T01 records a concrete compatibility requirement for retention beyond T07, document it in Evidence. |
| rebuild Topics table | Write the final Topics table to `product/records/spec/concepts/namespace-model/index.md` using the target structure from T01 Evidence. This table covers all namespace-model topics including those affected by T03–T05. T04 and T05 must not modify `index.md`. |

## Done condition

| item | done when |
|---|---|
| grammar topic created | `artifact-id-grammar.md` exists with stable ref `spec:product.concepts.namespace_model.artifact_id_grammar`; contains ID grammar, sequence format, allocation scopes, task sequence inheritance, subdomain exclusion, canonical-ref rule, and bare-form exclusion. |
| compatibility topic created | `legacy-id-compatibility.md` exists with stable ref `spec:product.concepts.namespace_model.legacy_id_compatibility`; enumerates all six V01-* legacy families; states issued IDs remain unchanged; states legacy forms are not canonical for new records; references `spec:product.concepts.spec_format.spec_id_as_ref` for canonical spec identity without redefining it. |
| v2-grammar retained | `v2-grammar.md` remains in the repository; its replacement topics are created and correct. Deletion is deferred to T07. |
| Topics table rebuilt | `index.md` Topics table uses permanent responsibility-based names; no `v1` or `v2` as permanent topic labels. |

## Verification

- Confirm `artifact-id-grammar.md` contains no legacy mapping content — issued-ID families belong in the compatibility topic.
- Confirm `legacy-id-compatibility.md` contains no grammar definitions — format rules belong in the grammar topic.
- Confirm `legacy-id-compatibility.md` does not define spec canonical identity; it references `spec:product.concepts.spec_format.spec_id_as_ref` only.
- Confirm the Topics table in `index.md` matches the target structure from T01 Evidence exactly.
- Confirm `v2-grammar.md` still exists in the repository (deletion deferred to T07).

## Evidence

- Created `product/records/spec/concepts/namespace-model/artifact-id-grammar.md` with stable ref `spec:product.concepts.namespace_model.artifact_id_grammar`. Contains: ID grammar for REQ/WORK/INV/ADR and TASK, sequence format table, allocation scopes, TASK work-sequence inheritance, subdomain exclusion note, canonical-ref rule, bare-form exclusion, and Related specs.
- Created `product/records/spec/concepts/namespace-model/legacy-id-compatibility.md` with stable ref `spec:product.concepts.namespace_model.legacy_id_compatibility`. Contains: all six V01-* legacy families enumerated, retention policy (issued IDs unchanged and resolvable, legacy not canonical for new records), spec identity note (V01-SPEC-* compatibility-only; references `spec:product.concepts.spec_format.spec_id_as_ref`), and Related specs.
- Updated `product/records/spec/concepts/namespace-model/index.md` Topics table: removed `v2 grammar`, `v1 namespace resolution algorithm`, `v1 record ID grammar` entries; added `Artifact ID grammar` and `Legacy ID compatibility` entries; updated `Existing artifacts` summary to reflect three-section restructure. Prose terminology updates deferred to T06.
- `v2-grammar.md` retained in repository; deletion deferred to T07.
- Source split per T01 Evidence: Grammar + sequence format sections → `artifact-id-grammar.md`; "V01-* unchanged" compatibility clause from PRODUCT namespace handling → `legacy-id-compatibility.md`; mapping rule table and new-artifact ownership clause → `existing-artifacts.md` (T03).

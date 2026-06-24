# PRODUCT-WORK-NAMESPACE-001: Namespace model canonical grammar and compatibility boundary cleanup

- **id**: PRODUCT-WORK-NAMESPACE-001
- **status**: done
- **date**: 2026-06-23
- **source_requirement**: V01-REQ-PRODUCT-001
- **impact_refs**:
  - spec:product.concepts.namespace_model
  - spec:product.concepts.namespace_model.artifact_id_grammar
  - spec:product.concepts.namespace_model.legacy_id_compatibility
  - spec:product.concepts.namespace_model.existing_artifacts
  - spec:product.concepts.traceability.artifact_refs
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.schema.id_normalization
- **tasks**:
  - PRODUCT-TASK-NAMESPACE-001-01
  - PRODUCT-TASK-NAMESPACE-001-02
  - PRODUCT-TASK-NAMESPACE-001-03
  - PRODUCT-TASK-NAMESPACE-001-04
  - PRODUCT-TASK-NAMESPACE-001-05
  - PRODUCT-TASK-NAMESPACE-001-06
  - PRODUCT-TASK-NAMESPACE-001-07
  - PRODUCT-TASK-NAMESPACE-001-08
  - PRODUCT-TASK-NAMESPACE-001-09

## Goal

Establish a stable namespace and artifact-ID contract before DRMCP specification redesign.

Separate permanent PRODUCT-owned namespace semantics from legacy compatibility policy and DRMCP-owned parsing, normalization, and scanning behavior.

## Boundary

This work item owns:

- replacing temporary `v1` and `v2` naming with responsibility-based namespace-model topics;
- separating canonical artifact ID grammar from legacy ID compatibility;
- moving parser, prefix-stripping, records-root scanning, and multi-root lookup contracts back to DRMCP specs;
- correcting ownership versus effective attribution language for existing artifacts;
- updating namespace-model, traceability, project-artifact-model, authoring-standard, and DRMCP references;
- independent review and final validation of the resulting boundary.

This work item does not own:

- changing issued legacy artifact IDs;
- bulk migration of existing files;
- changing the accepted app/domain namespace or sequence-allocation decisions;
- implementing DRMCP behavior;
- redesigning DRMCP authoring or query APIs beyond restoring their contract ownership.

## Impact Scope

| ref | impact |
|---|---|
| `spec:product.concepts.namespace_model` | Rebuild the topic map around permanent semantics and compatibility boundaries. |
| `spec:product.concepts.namespace_model.v2_grammar` | Replace temporary version naming with the canonical artifact-ID grammar topic. |
| `spec:product.concepts.namespace_model.v1_id_grammar` | Split PRODUCT legacy compatibility from DRMCP parser and normalization behavior. |
| `spec:product.concepts.namespace_model.v1_namespace_algorithm` | Return scanning and prefix-resolution behavior to DRMCP ownership. |
| `spec:product.concepts.namespace_model.existing_artifacts` | Separate issued identity policy from effective app attribution. |
| `spec:product.concepts.traceability` | Point ID-as-ref rules to the resulting canonical grammar and compatibility topics. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Restore the concrete records-root and multi-root scanning contract. |
| `spec:drmcp.design_records_mcp.schema.id_normalization` | Restore parser-facing public-ID and normalization behavior. |

## Task flow

1. Classify current namespace-model topics and record the target ownership map.
2. Establish the canonical artifact-ID grammar topic and the issued-ID compatibility policy topic. Depends on 1. Parallel with 3–5.
3. Separate historical ownership, effective attribution, and new-artifact ownership in the existing-artifact topic, referencing the issued-ID compatibility policy where identity is relevant. Depends on 1. Parallel with 2, 4–5.
4. Restore DRMCP-owned normalization contract to DRMCP ownership. Depends on 1. Parallel with 2, 3, 5.
5. Restore DRMCP-owned namespace-scanning contract to DRMCP ownership. Depends on 1. Parallel with 2, 3, 4.
6. Synchronize PRODUCT downstream consumers after canonical refs and ownership placement are settled. Depends on 2–5.
7. Synchronize DRMCP indexes and dependent contract references, then delete the obsolete PRODUCT topic files. Depends on 2–6.
8. Run independent review focused on ownership leakage, compatibility completeness, and stale refs. Independent reviewer. Depends on 6 and 7.
9. Apply review findings, run validation, and close the Work Item. Depends on 8.

## Task Candidates

- Classify current namespace-model topics and define the target ownership map.
- Establish the canonical artifact-ID grammar and issued-ID compatibility policy.
- Separate historical ownership, effective attribution, and new-artifact ownership, referencing the issued-ID compatibility policy where identity is relevant.
- Restore DRMCP ID normalization ownership.
- Restore DRMCP namespace scanning ownership.
- Synchronize PRODUCT downstream consumers.
- Synchronize DRMCP indexes and dependent contract references.
- Perform independent namespace and ID-as-ref boundary review.
- Apply review findings, run validation, and close the Work Item.

## Completion Condition

- PRODUCT namespace-model files contain only PRODUCT-owned semantic and compatibility contracts.
- No canonical topic name uses `v1` or `v2` as a permanent responsibility label.
- DRMCP owns concrete prefix stripping, parser normalization, records-root derivation, multi-root scanning, and lookup behavior.
- The canonical artifact-ID grammar and issued-ID compatibility policy have distinct stable refs.
- Existing issued IDs remain unchanged and resolvable.
- Historical ownership decision, issued identity, effective attribution, and new-artifact ownership are stated without contradiction.
- Traceability and authoring standards reference the new canonical topics without duplicating grammar.
- Independent review has no unresolved must-fix findings.
- Scoped grep confirms no DRMCP implementation behavior remains within PRODUCT namespace-model files.
- Validation reports no new format, reference, or workflow-integrity errors.
- `v01/` is unmodified.
- The V01-REQ-PRODUCT-001.work_items reciprocal-link exception is recorded in the Evidence section.

## Evidence

Initial evidence:

- `v2-grammar.md` combines permanent grammar, migration mapping, legacy retention, and PRODUCT-specific transition text.
- `v1-id-grammar.md` combines PRODUCT compatibility policy with DRMCP parser and normalization behavior.
- `v1-namespace-algorithm.md` defines DRMCP-specific records-root, prefix derivation, parser, and multi-root scanning behavior under PRODUCT namespace-model.
- `existing-artifacts.md` mixes issued ownership policy with effective app attribution.

Compatibility note:

- The source requirement is the accepted read-only `V01-REQ-PRODUCT-001`.
- `v01/` remains immutable, so its reciprocal `work_items` field is not updated.
- This exception is recorded explicitly for final validation and review.

Final state (T09 close):

**New permanent PRODUCT topics:**
- `spec:product.concepts.namespace_model.artifact_id_grammar` — canonical sequential-record ID grammar
- `spec:product.concepts.namespace_model.legacy_id_compatibility` — V01-* issued-ID retention and compatibility policy

**Updated PRODUCT topics:**
- `spec:product.concepts.namespace_model` (index, date, topic table)
- `spec:product.concepts.namespace_model.existing_artifacts` (ownership decision, effective attribution table, new-artifact ownership)

**Restored DRMCP topics:**
- `spec:drmcp.design_records_mcp.schema.id_normalization`
- `spec:drmcp.design_records_mcp.namespace_scanning`

**Deleted obsolete PRODUCT topics:**
- `spec:product.concepts.namespace_model.v2_grammar`
- `spec:product.concepts.namespace_model.v1_id_grammar`
- `spec:product.concepts.namespace_model.v1_namespace_algorithm`

**Downstream consumers synchronized:** traceability/, project-artifact-model/, authoring-standards/ (all 5 guides), repository-layout/, DRMCP-REQ-MCP-002.

**Independent review (T08):** 2 must-fix findings (M1 grammar scope, M2 attribution table completeness) — both applied. 2 advisory findings — A1 deferred, A2 applied (status update).

**Final validation:** `validate_spec.py --strict` — PRODUCT 31 files OK, DRMCP 30 files OK. v01/ unmodified.


# PRODUCT-TASK-SPEC-028-01: Decide retrospective term-analysis scope

- **id**: PRODUCT-TASK-SPEC-028-01
- **status**: done
- **date**: 2026-07-08
- **work_item**: PRODUCT-WORK-SPEC-028
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - PRODUCT-WORK-SPEC-028

## Goal

Decide how PRODUCT-WORK-SPEC-028 will admit, capture, and route the semantic-analysis work already performed over PRODUCT-INV-SPEC-011.

## Work

Resolve these decision items:

| item | status | required judgment | decision summary |
|---|---|---|---|
| D-001 | decided | Whether existing `tools/term-inventory-analysis/` outputs may be used as retrospective evidence. | Existing `tools/term-inventory-analysis/` output may be used as retrospective source material for deriving product-side evidence. It is not durable repository evidence by itself. Only product-side captured summaries or derived artifacts count as repository-persistent evidence candidates. |
| D-002 | decided | Which analysis stages are accepted as already completed. | Completed stages: PRODUCT-INV-SPEC-011 raw inventory; leaf semantic analysis with 71 of 71 valid results; trigger-level reduction with 57 of 57 valid results; Tier A cross-trigger identity review with 54 of 54 valid jobs and spot audit PASS. Leaf and trigger-level reduction outputs still require product-side evidence capture before repository history can rely on them. Canonical vocabulary, term definition, deprecation, source rewrite, and projection remain outside the accepted completed stage set. |
| D-003 | decided | Which product-side evidence artifacts must be created or updated. | Required product-side evidence artifacts: create `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/leaf-analysis/`; create `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/trigger-reduction/`; update `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/cross-trigger-review/` only as needed to connect the existing Tier A review summary to the PRODUCT-WORK-SPEC-028 evidence policy. Each created directory should contain a human-readable `README.md` and a machine-readable summary JSON. Raw per-job tools output stays under ignored `tools/`. |
| D-004 | decided | Whether the Tier A cross-trigger review evidence needs additional independent review. | Additional independent review is not required for the existing Tier A cross-trigger review evidence. The existing review completed 54 of 54 jobs, passed validation with no errors or warnings, and received a PASS spot audit with no findings. D-003 may still require evidence synchronization to connect the existing summary to PRODUCT-WORK-SPEC-028. That synchronization is not an additional independent review. |
| D-005 | decided | Which next Work Items or Tasks are required before canonical vocabulary decisions. | Required next Tasks inside PRODUCT-WORK-SPEC-028: T02 investigation for product-side leaf-analysis evidence capture; T03 investigation for product-side trigger-reduction evidence capture; T04 synchronization for connecting existing cross-trigger-review evidence to the PRODUCT-WORK-SPEC-028 evidence policy; T05 coordination for follow-up routing before canonical vocabulary decisions; T06 integrated independent review over the captured evidence and routing; T07 synchronization after review PASS. Canonical vocabulary approval remains a separate downstream Work Item route, not a Task inside PRODUCT-WORK-SPEC-028. |
| D-006 | decided | Whether PRODUCT-REQ-SPEC-012 restart criteria can be partially or fully satisfied by this Work Item. | PRODUCT-WORK-SPEC-028 can partially satisfy PRODUCT-REQ-SPEC-012 restart criteria only after its evidence capture, follow-up routing, integrated review, and synchronization complete. Full restart remains blocked until downstream vocabulary, conflicting-meaning, qualified-term, and deprecation dispositions exist. |

Do not decide canonical vocabulary inside this Task.
Do not rewrite source records inside this Task.

## Done condition

- Every decision item is `decided`, `deferred`, or validly `blocked`.
- The accepted analysis scope is explicit.
- The retrospective evidence policy is explicit.
- The next Task graph or stop condition is explicit.

## Verification

PRODUCT-WORK-SPEC-028 still lists PRODUCT-TASK-SPEC-028-01 in its `tasks` field.
No canonical vocabulary, deprecation, source rewrite, or Specification projection was performed by this Task.
Every decision item is decided.
The accepted analysis scope, retrospective evidence policy, and next Task graph are explicit.

## Evidence

- PRODUCT-REQ-SPEC-014 requires product-owned semantic-analysis evidence.
- PRODUCT-INV-SPEC-011 is the raw evidence source.
- The existing cross-trigger review summary is currently stored as commit-safe evidence, but its parent Work Item boundary was missing before PRODUCT-WORK-SPEC-028.
- D-001 decision: Existing `tools/term-inventory-analysis/` output is admissible only as retrospective source material. Durable evidence requires product-side capture under repository-tracked records or data.
- D-001 basis: PRODUCT-REQ-SPEC-014 excludes treating ignored tools output as durable product evidence without product-side capture. PRODUCT-WORK-SPEC-028 treats `tools/term-inventory-analysis/` as ignored working output unless specific summaries are captured under product records. The cross-trigger review README records the same storage policy.
- D-002 decision: Accept the raw inventory, leaf semantic analysis, trigger-level reduction, and Tier A cross-trigger identity review as completed analysis stages.
- D-002 basis: PRODUCT-INV-SPEC-011 is concluded with 5,699 observations across 733 sources. Leaf analysis validation reports 71 selected jobs, 71 valid results, 0 missing results, and complete true. Trigger reduction validation reports 57 selected jobs, 57 valid results, 0 missing results, and complete true. The Tier A cross-trigger review summary reports 54 completed jobs, 0 missing jobs, valid true, complete true, and spot audit PASS.
- D-002 limit: Leaf and trigger-level reduction remain tools-side outputs until product-side summaries or derived artifacts capture them. Canonical vocabulary, term definition, deprecation, source rewrite, and projection remain excluded.
- D-003 decision: Product-side evidence capture requires a leaf-analysis evidence directory, a trigger-reduction evidence directory, and a possible cross-trigger-review evidence update.
- D-003 basis: PRODUCT-REQ-SPEC-014 requires product-owned evidence for trigger-level aggregation, reduction, cross-trigger candidate generation, routing, review completion, and spot audit. D-001 rejects raw ignored tools output as durable evidence. D-002 accepts leaf semantic analysis and trigger-level reduction as completed tools-side stages that still require product-side capture.
- D-003 output boundary: Evidence capture should use `README.md` plus machine-readable summary JSON. Raw per-job tools output remains ignored working output.
- D-004 decision: Do not require an additional independent review for the existing Tier A cross-trigger review evidence.
- D-004 basis: The product-side Tier A summary records 54 completed jobs, 0 missing jobs, 0 invalid jobs, 0 errors, 0 warnings, valid true, and complete true. The spot audit records overall verdict PASS, 7 audited jobs, 0 findings, and no follow-up correction required.
- D-004 limit: PRODUCT-WORK-SPEC-028 may still require synchronization or evidence-link updates for the existing cross-trigger review summary. That work is evidence capture or synchronization, not a new independent review gate.
- D-005 decision: Materialize follow-up Tasks for evidence capture, evidence-policy synchronization, follow-up route coordination, integrated independent review, and closure synchronization.
- D-005 task graph: T02 captures leaf-analysis evidence. T03 captures trigger-reduction evidence. T04 synchronizes existing cross-trigger-review evidence with this Work Item's evidence policy. T05 coordinates follow-up routes for canonical vocabulary, conflicting meanings, qualified terms, deprecation, and PRODUCT-REQ-SPEC-012 restart. T06 performs integrated independent review. T07 performs post-review synchronization.
- D-005 basis: The completed tools output is too large and too raw for normal human review. PRODUCT-WORK-SPEC-028 should make the work reviewable by creating compact product-side summaries and routing evidence instead of requiring reviewers to inspect ignored per-job output directly.
- D-005 limit: Canonical vocabulary approval, term definition, deprecation, source rewrite, and Specification projection remain outside PRODUCT-WORK-SPEC-028. T05 may create or route those downstream Work Items, but it must not decide their substantive outcomes.
- D-006 decision: PRODUCT-WORK-SPEC-028 can partially satisfy PRODUCT-REQ-SPEC-012 restart criteria only after the Work Item's evidence capture, follow-up routing, integrated review, and synchronization complete.
- D-006 basis: PRODUCT-WORK-SPEC-028 can produce reviewed evidence about foundational term boundaries and follow-up routes. It does not approve canonical vocabulary, define terms, split qualified terms, make retirement or deprecation decisions, rewrite source records, or project results into Specifications.
- D-006 limit: Full PRODUCT-REQ-SPEC-012 restart remains blocked until downstream vocabulary, conflicting-meaning, qualified-term, and deprecation dispositions exist.
- Loop state: decision_complete.

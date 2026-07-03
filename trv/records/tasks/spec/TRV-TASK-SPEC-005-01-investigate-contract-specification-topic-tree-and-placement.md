# TRV-TASK-SPEC-005-01: Investigate contract Specification topic tree and placement

- **id**: TRV-TASK-SPEC-005-01
- **status**: done
- **date**: 2026-07-03
- **work_item**: TRV-WORK-SPEC-005
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - TRV-TASK-SPEC-001-14
- **outputs**:
  - TRV-TASK-SPEC-005-01
  - TRV-INV-SPEC-003

## Goal

Produce one bounded Investigation of architecture-derived contract topics and candidate Specification Markdown placement.

## Work

- Read the complete reviewed W002 architecture Specification set.
- Inventory every component, application port, application model boundary, validation-flow handoff, application outcome, and architecture handoff that requires a W005 contract owner.
- Identify concerns already owned by PRODUCT, W002, W004, or current DRMCP authority.
- Compare coherent candidate `spec:trv` topic trees and physical Markdown placement.
- Evaluate topic cohesion, duplication risk, navigation, future W004 projection, and review boundaries.
- Record factual conflicts, uncertainties, and follow-up judgment candidates in TRV-INV-SPEC-003.

This Task must not decide the topic tree, author or amend ADRs or Specifications, change the Task graph, perform review, closure, or implementation.

## Done condition

- TRV-INV-SPEC-003 concludes one architecture-to-contract placement question.
- Every W002 contract-relevant boundary is inventoried or explicitly excluded.
- Candidate topic trees and Markdown placements are compared without adopting one.
- Open questions and decision candidates are exact enough for T02.

## Verification

- Confirm the Investigation follows active Investigation authoring rules.
- Confirm the scope is limited to contract Specification partition and placement.
- Confirm PRODUCT semantics, W002 architecture, and W004 detail remain unmodified.
- Confirm no decision or canonical authoring occurred.

## Evidence

- T14 created this Investigation owner as the first W005 responsibility.
- Created `TRV-INV-SPEC-003` under `trv/records/investigations/spec/` after correcting the duplicate Investigation number.
- Reviewed the complete W002 architecture Specification set, `spec:trv.model_runtime`, PRODUCT semantic authority, TRV-ADR-SPEC-001 through TRV-ADR-SPEC-005, retired W003, and W004.
- Inventoried every W002 application port, boundary model, validation-flow handoff, outcome, adapter boundary, and explicit exclusion.
- Compared four placement patterns without adopting one.
- Candidate A, the retired flat external-first tree, cannot cover the architecture-derived application and non-MCP adapter contracts.
- Candidate B, peer application and adapter areas, remains viable.
- Candidate C, one `spec:trv.contracts` root with application and adapter children, appears preferable but remains undecided.
- Candidate D, contracts nested under application architecture, conflicts with separate W002 and W005 completion boundaries.
- During T02, corrected the initial generic `external` grouping after it was found to omit parallel record/checklist and model-provider adapter areas.
- Recorded ten exact T02 judgment candidates, stale W003 references, ADR migration gaps, and shared-writer risks.
- Confirmed PRODUCT semantics, W002 architecture, W004 detailed design, current DRMCP, and production implementation remained unmodified.
- The standalone responsibility-boundary validator returned compliant results for all 23 common and investigation criteria after the adapter-boundary correction and final Evidence update.
- Result: `PASS`.

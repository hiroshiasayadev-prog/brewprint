# PRODUCT-TASK-SPEC-004-02: Review ownership boundary decision and relocation plan

- **id**: PRODUCT-TASK-SPEC-004-02
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-004
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-004-01
- **outputs**:
  - Review findings (inline in Evidence)

## Goal

Independent review of the draft ownership boundary decision and relocation plan from PRODUCT-TASK-SPEC-004-01. This decision has a wide blast radius — it determines target structure for PRODUCT-WORK-SPEC-005, PRODUCT-WORK-SPEC-009, and the already-drafted PRODUCT-TASK-SPEC-005-13..-16 — so it is a review gate before acceptance.

## Work

| area | what to check |
|---|---|
| completeness | Every relocation candidate from PRODUCT-INV-SPEC-002 and PRODUCT-INV-SPEC-004 is addressed in the decision. |
| consistency | The decision does not contradict either investigation's classification reasoning without stating why. |
| relocation plan ordering | Move order is logically sound (e.g. format-only migration before relocation, per the PRODUCT-WORK-SPEC-009 precedent). |
| downstream impact | Confirm what changes for PRODUCT-WORK-SPEC-005 (in particular PRODUCT-TASK-SPEC-005-13..-16's already-drafted output file list) and PRODUCT-WORK-SPEC-009 (whether `namespace-model/index.md` is added). |
| no premature execution | Confirm no actual file moves happened in PRODUCT-TASK-SPEC-004-01 — this task and -01 are decision-only. |

## Done condition

| item | done when |
|---|---|
| review complete | Review findings are attached in Evidence. |
| findings classified | Each finding is classified: must-fix before 004-03, or defer. |
| user sign-off | User approves proceeding to 004-03. |

## Verification

This task is itself a review gate. Proceed to PRODUCT-TASK-SPEC-004-03 only after user sign-off on findings.

## Evidence

Review basis: the draft decision in PRODUCT-TASK-SPEC-004-01 was reviewed against the user's own framing ("drmcp's spec is nothing more than about the actual implementation, and it should be the mirror of product spec's dependencies"). Decision logic is consistent with that intent. Source files are in pre-migration Japanese format and were not re-read for this review — the classification derives from prior reads in this session and PRODUCT-INV-SPEC-002/004's evidence.

No must-fix findings. Two notes recorded as defer:

| finding | classification | disposition |
|---|---|---|
| Decision doesn't specify what happens to `drmcp/overview.md §既存brewprint MCPとの責務境界` cross-ref cleanup | defer | This is an editorial note for PRODUCT-TASK-SPEC-005-14, not a relocation issue. Carry into -14's Work table as a note: cross-ref `project-artifact-model/index.md` boundary table instead of restating rule independently. |
| WORK-SPEC-002 ambiguity (hyphen vs. underscore in spec: refs) is carried forward unresolved | defer | Correctly out of WORK-SPEC-004's scope. No action needed in 004-03. |

User sign-off received 2026-06-17. Proceeding to PRODUCT-TASK-SPEC-004-03.

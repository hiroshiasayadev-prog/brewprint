# PRODUCT-TASK-SPEC-026-02: Amend term-inventory Requirement boundary

- **id**: PRODUCT-TASK-SPEC-026-02
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-026
- **task_type**: authoring
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-026-01
- **outputs**:
  - PRODUCT-REQ-SPEC-013

## Goal

Amend PRODUCT-REQ-SPEC-013 so the downstream inventory work owns direct JSONL observation gathering and coverage Evidence, not aggregation or helper tooling.

## Work

- Replace the Requirement boundary statement that assigned generic tool support to the downstream Work Item.
- State that direct JSONL observation authoring and coverage Evidence belong to the downstream inventory Work Item.
- State that aggregation, helper tooling, schema-validation tooling, classification, definition, preferred wording, retirement, and use-case extraction belong to later work only when the corpus justifies them.
- Clarify that later analytical work is conditional rather than mandatory.

## Done condition

PRODUCT-REQ-SPEC-013 accurately separates investigation evidence gathering from optional later aggregation, tooling, classification, and use-case work.

## Verification

- Confirm the Required Outcome still ends at corpus gathering and coverage.
- Confirm helper-tool implementation and aggregation are not downstream inventory completion conditions.
- Confirm JSONL observation shape remains unchanged.

## Evidence

- The user stated that separate sessions will author assigned JSONL directly.
- The user excluded validation tooling from the current investigation.
- The user classified aggregation and use-case extraction as separate responsibilities after the investigation result is known.
- PRODUCT-REQ-SPEC-013 was amended accordingly.
- DRMCP is non-operational. Filesystem authoring is the required fallback.

# PRODUCT-TASK-SPEC-010-01: Rule selection from INV-SPEC-003 candidates

- **id**: PRODUCT-TASK-SPEC-010-01
- **status**: done
- **date**: 2026-06-15
- **work_item**: PRODUCT-WORK-SPEC-010
- **source_requirement**: PRODUCT-REQ-SPEC-002
- **estimate**: 0.25d
- **depends_on**:
- **outputs**:

## Goal

Determine which PRODUCT-INV-SPEC-003 candidate rules to adopt, at what normative level, and resolve the three scoping decisions required before authoring.

## Work

- Review spec-side and AI output candidate tables from PRODUCT-INV-SPEC-003.
- Resolve scope decisions: (1) AI output rule placement; (2) vocabulary supplement; (3) BLUF handling.
- Assign normative levels (MUST / SHOULD) to each adopted rule.

## Done condition

| item | done when |
|---|---|
| Rule set confirmed | User approves adopted rules and deferred rules. |
| Normative levels assigned | Each adopted rule has MUST or SHOULD. |
| Scope decisions resolved | AI output placement, vocabulary supplement, and BLUF decisions recorded. |

## Verification

User confirmation in session on 2026-06-15.

## Evidence

- Rule set confirmed by user on 2026-06-15.
- Scope decisions: AI output rules go in the spec (CLAUDE.md / prompt_chappy.md reference it); vocabulary supplement removed from scope; BLUF dropped as a standalone rule (enacted structurally by `## What this is`).
- Normative levels: subordinating conjunction rule and 20-word cap as SHOULD; label-first bullets as SHOULD; all others as MUST.

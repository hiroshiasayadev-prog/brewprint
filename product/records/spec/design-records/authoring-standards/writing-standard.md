# Reference: Design record writing standard

- **id**: `spec:product.design_records.authoring_standards.writing_standard`
- **status**: accepted
- **date**: 2026-06-15
- **parent**: `spec:product.design_records.authoring_standards`

## What this is

Writing rules for brewprint design record prose. Governs sentence style, block type selection, and AI output behavior in design record authoring.

## Non-goals

- Vocabulary restriction lists.
- Full ASD-STE100, DITA, or Plain Language compliance.
- Bulk rewrite of existing records.
- Section title and heading format rules — those belong to `spec:product.design_records.spec_format.document_shape`.

## Rules

### Spec-side rules

| Rule | Level |
|---|---|
| Prefer structured blocks (table, bullet, labeled block); reserve prose for rationale and causality. | MUST |
| No throat-clearing. Remove `it is important to note that`, `this document describes`, and similar lead-in phrases. | MUST |
| Do not duplicate rationale across records. A one-line summary is allowed when linking to the source record. | MUST |
| Prefer verb forms over nominalizations (`use` not `utilize`; `define` not `provide a definition of`). | SHOULD |
| One sentence per claim. | MUST |
| Prefer active voice. Passive is allowed when the actor is irrelevant or unknown. | MUST |
| Target 20 words per sentence maximum. Longer sentences are allowed only when splitting reduces precision. | SHOULD |
| Do not use ambiguous pronouns (`this`, `that`, `it`); repeat the noun when the referent is unclear. | MUST |
| Use no more than one subordinating conjunction (`because`, `which`, `when`, `where`, `while`, `although`, `if`) per sentence. Split or use bullets instead. | SHOULD |
| Use a table for condition branching. Do not use if/then prose. | MUST |
| Use a table for matrices, state transitions, and input/output contracts. | MUST |
| Use bullets for constraints, non-goals, and acceptance criteria. | MUST |
| Use label-first bullet prefixes (`Rule:` / `Exception:` / `Reason:` / `Evidence:`) when classification aids readability. | SHOULD |
| Preserve domain terms (`invoke`, `boundary`, `artifact`, `diagnostic`, and equivalent domain vocabulary). Do not replace them with simpler synonyms. | MUST |

Note: BLUF (conclusion before rationale) is enacted structurally by the `## What this is` section. It is not a separate prose rule.

### AI output rules

These rules apply to AI-authored content that appears in design records: recommendations, rule proposals, and review verdicts.

| Rule | Level |
|---|---|
| Separate recommendations from normative text. | MUST |
| Mark non-exhaustive examples as `Examples, not exhaustive`. | MUST |
| Use `Candidate:` for undecided rules. Do not use `MUST` or `SHOULD` for unconfirmed rules. | MUST |
| For an embedded review verdict field, use `PASS` or `NEEDS REVISION` followed by one concise reason. | SHOULD |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.authoring_standards` | Parent Index for this spec. |
| `spec:product.design_records.spec_format.document_shape` | Section and heading format rules for spec records. |
| PRODUCT-INV-SPEC-003 | Source investigation. Candidate rules and writing standard survey. |
| PRODUCT-WORK-SPEC-010 | Source work item. |

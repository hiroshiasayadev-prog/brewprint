# Overview: Temporary BPDSL staging

- **id**: `spec:product.bpdsl`
- **status**: draft
- **date**: 2026-06-24
- **parent**: `spec:product`

## What this is

Temporary quarantine and staging for BPDSL-related material separated from mixed PRODUCT specifications.
It preserves context without claiming canonical PRODUCT ownership.

## Current contract

| topic | contract |
|---|---|
| Purpose | Isolate existing BPDSL-related material while Design Records specifications are cleaned. |
| Ownership status | This area is not the canonical BPDSL specification hierarchy. Placement creates no PRODUCT ownership claim. |
| Expected final owner | `bpdsl/records/spec/**` remains the expected owner for BPDSL-internal contracts, subject to migration review. |
| Allowed content | Existing DSL, source, render, implementation-flow, or BPDSL artifact descriptions requiring preservation. |
| Context limit | Include only context required to preserve existing meaning. |
| Migration trigger | Review every staged item during BPDSL migration or an explicit integration requirement. |

## Non-goals

- BPDSL redesign or hierarchy normalization.
- New BPDSL schema, type, resolver, render, generation, runtime, or MCP contracts.
- New Design Records-to-BPDSL integration design.
- Evaluation of BPDSL correctness or app-local specification quality.
- Unrelated new BPDSL specifications.

## Rules

- Preserve staged content without semantic redesign.
- Do not treat staged content as reviewed or accepted BPDSL truth.
- Do not add canonical BPDSL ownership claims.
- Do not normalize app-local BPDSL specifications during this restructuring.
- Use Design Records and app-local BPDSL contracts only as contextual pointers.

## Boundary

Content in this area must remain preservation-only, make no canonical BPDSL ownership claim, and introduce no new Design Records-to-BPDSL integration rules.

Review of this area checks separation, preservation-only framing, and absence of canonical ownership claims. It excludes BPDSL correctness, app-local duplication, final layout, final ownership, and integration design.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Design flow preservation | Reference | `spec:product.bpdsl.design_flow` | Preservation-only copy of Design Records-to-DSL-to-source flow material. |
| Artifact responsibility preservation | Reference | `spec:product.bpdsl.artifact_responsibilities` | Preservation-only copy of BPDSL and implementation-flow responsibility rows. |
| Repository implementation-flow preservation | Reference | `spec:product.bpdsl.repository_implementation_flow` | Preservation-only copy of `dsl/`, `src/`, generation, and bootstrap layout material. |

## Exit conditions

| disposition | result |
|---|---|
| Move to app-local BPDSL | Place BPDSL-internal content under `bpdsl/records/spec/**`. |
| Redefine as PRODUCT policy | Accept a genuine cross-cutting PRODUCT responsibility through an explicit decision. |
| Relocate elsewhere | Move evidence or policy to the appropriate design record or owner. |
| `delete` | Delete obsolete or duplicate content after evidence transfer. |

After exit review, remove this temporary area or explicitly redefine it through an accepted decision.

## Related specs

| ref | relation |
|---|---|
| `spec:product` | PRODUCT placement router and dependency direction. |
| `spec:product.design_records` | Generic semantics that must remain independent from BPDSL. |
| PRODUCT-ADR-SPEC-001 | Accepts the temporary staging and exit contract. |

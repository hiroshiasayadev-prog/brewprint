# PRODUCT-REQ-SPEC-014: Analyze Brewprint design-governance term inventory

- **id**: PRODUCT-REQ-SPEC-014
- **status**: accepted
- **date**: 2026-07-08
- **source_refs**:
  - PRODUCT-REQ-SPEC-012
  - PRODUCT-REQ-SPEC-013
  - PRODUCT-INV-SPEC-011

## Requirement

Analyze the PRODUCT-INV-SPEC-011 observation corpus to identify reusable governance-term meaning groups and unresolved semantic boundaries.

The analysis must convert the raw observation corpus into reviewable evidence for later vocabulary, definition, qualification, deprecation, and Task-type use-case work.

The analysis must preserve the raw 32 batch-owned observation files as source evidence.

The analysis must not treat automated similarity, clustering, or model output as canonical vocabulary approval.

## Evidence

- PRODUCT-INV-SPEC-011 concluded with 5,699 conforming observation records across 733 accounted sources.
- PRODUCT-WORK-SPEC-027 explicitly excluded aggregation, semantic clustering, term classification, definition, normalization, approval, replacement, and consolidation.
- PRODUCT-INV-SPEC-011 recommends a separate follow-up for machine-assisted aggregation and semantic analysis.
- PRODUCT-REQ-SPEC-012 remains deferred until foundational cross-cutting terms have accepted definitions, qualified-term splits, retirement decisions, or explicit unresolved dispositions.
- The repository already contains commit-safe evidence for Tier A cross-trigger semantic identity review under `product/records/investigations/spec/data/PRODUCT-INV-SPEC-011/cross-trigger-review/`.
- Raw per-job prompts and result files remain under ignored tools output and require product-side evidence capture before repository history can rely on them.

## Required Outcome

- The analysis scope and retrospective-evidence policy are explicitly decided.
- Trigger-level aggregation and reduction evidence from PRODUCT-INV-SPEC-011 observations is recorded in a product-owned artifact location.
- Cross-trigger identity candidate generation, routing, review completion, and spot-audit evidence are recorded in a product-owned artifact location.
- Identity groups, unresolved candidates, relationship hints, and confidence limits are separated from canonical vocabulary decisions.
- Follow-up routes are identified for at least these outcomes:
  - canonical vocabulary decisions;
  - conflicting-meaning investigation;
  - qualified-term split review;
  - deprecated or misleading wording review;
  - PRODUCT-REQ-SPEC-012 restart criteria.
- No Specification, skill, authoring guide, validator, or source artifact is rewritten by this requirement alone.

## Explicitly Excluded Scope

- Approving canonical vocabulary.
- Defining preferred wording.
- Retiring or deprecating terms.
- Rewriting existing source records.
- Projecting term decisions into Specifications, skills, authoring guides, or validators.
- Completing PRODUCT-REQ-SPEC-012 directly.
- Treating ignored tools output as durable product evidence without product-side capture.

## Boundary

This Requirement owns the need to turn PRODUCT-INV-SPEC-011 raw observations into reviewed semantic-analysis evidence.

This Requirement does not own canonical terminology adoption, current normative specification text, or source-record migration.

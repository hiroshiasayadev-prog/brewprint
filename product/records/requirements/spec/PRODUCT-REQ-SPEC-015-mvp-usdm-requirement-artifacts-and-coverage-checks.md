# PRODUCT-REQ-SPEC-015: MVP USDM requirement artifacts and coverage checks

- **id**: PRODUCT-REQ-SPEC-015
- **status**: accepted
- **date**: 2026-07-08
- **source_refs**:
  - DRMCP-WORK-MCP-022

## Requirement

Brewprint must introduce MVP USDM requirement artifacts for Design Records specifications.

The MVP must organize specification requirements by topic and make implementation coverage statically checkable.

The MVP must allow each implementation Specification to declare the USDM requirement IDs it covers.

The MVP must prevent requirement omissions during implementation-spec authoring and review.

The MVP must be usable before full artifact-model, repository-layout, migration, and DRMCP integrated-read support are complete.

## Evidence

- DRMCP-WORK-MCP-022 exposed a missing middle layer between Product specifications and implementation contracts.
- Domain component and input/output decisions are unsafe before Product concepts and constraints are inventoried.
- DRMCP-TASK-MCP-022-03 and DRMCP-TASK-MCP-022-04 worked better because validation requirements were inventoried before response and ownership mapping.
- Existing Specification metadata does not provide a static coverage check from implementation specs back to normalized implementation requirements.
- Early standalone tooling is acceptable because formal artifact-model and repository-layout migration would delay W022 recovery work.

## Required Outcome

- USDM overview, artifact-format, coverage-format, and standalone-tool contract specs exist under `product/records/spec/design-records/usdm/`.
- A USDM authoring standard exists under `product/records/spec/design-records/authoring-standards/usdm-authoring.md`.
- USDM records are placed under each app namespace at `<app>/records/usdm/` for the MVP.
- USDM artifact IDs use `usdm:<app_namespace>.<path.to.topic>`.
- USDM artifact kinds are limited to `index` and `requirement` for USDM records in the MVP.
- USDM requirement rows produce full requirement IDs as `usdm:<app_namespace>.<path.to.topic>#RNNN`.
- USDM `index` records require only `## What this is` in the MVP.
- USDM `requirement` records require `## What this is` and one or more `## Requirements: <spec ref>` tables.
- Requirement row IDs are unique and sequential within one USDM requirement record.
- Implementation Specifications may declare `usdm_covers` metadata.
- Standalone tools under `tools/usdm/` can validate USDM format and check coverage.
- The MVP specifies standalone tool behavior for at least `validate_usdm`, `check_usdm_coverage`, and `usdm_covered_by`.
- Full artifact-model integration remains a later Work Item.
- Full repository-layout integration remains a later Work Item.
- Existing design record migration remains a later Work Item.

## Explicitly Excluded Scope

- Migrating existing Product, DRMCP, or other app records to USDM.
- Making USDM a fully integrated canonical artifact kind in the initial MVP.
- Rewriting the artifact model or repository layout as part of the MVP.
- Requiring section-level coverage before file-level `usdm_covers` proves insufficient.
- Completing DRMCP integrated USDM read support before DRMCP read support is ready.
- Deciding DRMCP Domain component inputs and outputs inside this Product requirement.
- Replacing Product Specifications with USDM records.

## Boundary

This requirement owns the MVP need for USDM records, authoring rules, coverage metadata, and standalone coverage tooling.

This requirement does not own the later migration into the canonical artifact model, the repository layout, or DRMCP runtime integration.

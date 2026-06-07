# REQ-PRODUCT-001: App and domain namespace model for namespace-first design records

- **id**: REQ-PRODUCT-001
- **status**: accepted
- **date**: 2026-06-07
- **source_refs**:
  - REQ-MCP-013
- **work_items**:
  - WORK-PRODUCT-001

## Requirement

Design records need an explicit namespace model that separates app namespace from domain namespace before any namespace-first layout or major-version artifact migration is attempted.

The app namespace identifies the owning product area, application, subsystem, or cross-application product scope. Examples include `DRMCP` for Design Records MCP, `DRUI` for Design Records UI, `BPDSL` for Brewprint DSL, and `PRODUCT` for repository-wide or cross-application concerns.

The domain namespace identifies the concern area inside an app namespace. Examples include `MCP`, `DATA`, `RESOLVE`, `UI`, `LAYOUT`, `NAMESPACE`, `GOVERNANCE`, and `MIGRATION`.

The target v2 artifact ID grammar should preserve both axes, for example `<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<SEQUENCE>`, with task IDs extending the work sequence, for example `<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>`. Existing IDs such as `REQ-DATA-013` should be mappable to IDs such as `DRMCP-REQ-DATA-013` when the owning app namespace is Design Records MCP.

A product-level namespace must exist for requirements, decisions, policies, and migration records that affect multiple app namespaces.

## Evidence

- The current artifact layout and ID scheme primarily encode artifact kind and domain, for example `REQ-MCP-*`, `WORK-DATA-*`, and `TASK-DATA-*-*`.
- As Design Records MCP, Brewprint DSL, and the proposed Design Records UI become distinct app-level concerns, the current domain-only ID scheme cannot distinguish the owning app namespace from an internal domain namespace.
- UI-level namespace projection can hide this ambiguity temporarily, but it does not define ownership, cross-app scope, or a safe migration path.
- Cross-application requirements are expected, especially for repository layout, namespace governance, migration criteria, and major-version compatibility policy.

## Required Outcome

- A namespace model exists that explicitly distinguishes app namespace and domain namespace.
- A product-level namespace is defined for cross-application artifacts.
- A target v2 artifact ID grammar is specified that includes both app namespace and domain namespace.
- Existing domain-first IDs can be mapped to target namespace-aware IDs without relying on ad-hoc inference in UI code.
- The namespace model can serve as a prerequisite for future namespace catalog, namespace-aware MCP discovery, and namespace-first physical layout migration work.

## Explicitly Excluded Scope

- Immediate migration of existing artifacts to the v2 ID grammar.
- Immediate physical relocation of existing artifacts into namespace-first directories.
- Implementation of namespace-aware MCP APIs.
- Implementation of Design Records UI screens.
- Final compatibility, alias, or rollback policy for a major-version migration.

## Boundary

This requirement owns the need for a stable namespace model and target artifact ID grammar.

It does not own the final current-state specification, the migration implementation, the UI implementation, or the MCP API changes required to expose namespace-aware discovery. Those should be handled by follow-up spec, ADR, work item, and task records.

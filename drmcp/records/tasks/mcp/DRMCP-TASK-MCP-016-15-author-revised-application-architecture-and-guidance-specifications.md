# DRMCP-TASK-MCP-016-15: Author revised application architecture and Guidance Specifications

- **id**: DRMCP-TASK-MCP-016-15
- **status**: done
- **date**: 2026-07-04
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: authoring
- **estimate**: 2.0d
- **depends_on**:
  - DRMCP-TASK-MCP-016-14
- **outputs**:
  - spec:drmcp.application_architecture
  - spec:drmcp.application_architecture.application_boundary_and_components
  - spec:drmcp.application_architecture.dependency_and_responsibility
  - spec:drmcp.application_architecture.runtime_and_state
  - spec:drmcp.application_architecture.failure_and_evolution
  - spec:drmcp.design_records_mcp.namespace_scanning
  - spec:drmcp.design_records_mcp.schema.overview
  - spec:drmcp.design_records_mcp.schema.discovery
  - spec:drmcp.design_records_mcp.schema.id_normalization
  - spec:drmcp.design_records_mcp.schema.record_source
  - spec:drmcp.design_records_mcp.schema.authoring_guidance_source
  - spec:drmcp.design_records_mcp.tools.list_authoring_guides
  - spec:drmcp.design_records_mcp.tools.get_authoring_guidance
  - DRMCP-TASK-MCP-016-15

## Goal

Author the revised canonical Specification state from T12 decisions and ADR-010 through ADR-012.

## Work

- Replace the six-component architecture with the accepted five-component model.
- Remove Guidance Domain and Guidance Source ownership from the application architecture.
- Treat the portable standards package as a configured spec-only Current Records source under `design_records`.
- Generalize current source mapping so records roots and explicit spec-tree roots share normal current Spec semantics.
- Include package Specs in the normal request-scoped Current Records snapshot and active index.
- Define Guidance tools as fixed-scope aliases over shared record-query orchestration.
- Use canonical package Spec refs as guide IDs.
- Preserve first-H1 title, `## What this is` abstract, and verbatim body projection.
- Remove stale direct PRODUCT directory, filename-stem identity, and separate Guidance source contracts.
- Preserve Current Records versus Legacy Archive separation and deferred Authoring boundaries.

## Done condition

- The five architecture Specifications agree with ADR-010 through ADR-012.
- Current source contracts accept an explicit spec-tree root with an assigned app namespace.
- Normal current Spec discovery, identity, record model, and active index apply to the portable package.
- Guidance list and get use the `design_records` authoring-standards subtree.
- No Guidance-specific domain, source port, index, snapshot, parser, cache, or lifecycle remains.
- No stale `docs/guides`, direct PRODUCT source, filename-stem guide ID, or `## Abstract` projection remains in the assigned scope.
- T16 can review exact final artifacts without new authoring judgment.

## Verification

- Confirm all assigned Specification dates are `2026-07-04` after substantive change.
- Confirm the component count is five everywhere in the architecture topic.
- Confirm `Guidance Domain` and `Guidance Source` are absent from current normative text.
- Confirm the portable package uses `app_namespace: design_records` and canonical `spec:design_records.*` refs.
- Confirm DRMCP performs no runtime namespace rewrite.
- Confirm Guidance uses shared record-query orchestration without public use-case chaining.
- Confirm list scope excludes `spec:design_records.authoring_standards` and includes its child subtree.
- Confirm Current and Legacy state remain separate.
- Confirm exact interfaces, types, functions, concrete config serialization, implementation, tests, and fixtures remain outside scope.
- Inspect the complete scoped diff and whitespace result.

## Evidence

- The repository-root `prompt_chappy.md` was read first.
- `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational. Design Record authoring used filesystem fallback.
- T12 D-018 through D-024 supplied the revised decisions.
- T14 supplied DRMCP-REQ-MCP-003 and ADR-010 through ADR-012 as current authority.
- Updated all five application-architecture Specifications.
- Updated the schema Overview and current source configuration, discovery, identity, and source mapping Specifications.
- Updated the Guidance schema and both Guidance operation Specifications.
- Removed the separate Guidance Domain, source port, source directory, filename-stem identity, and standalone lifecycle model.
- Preserved the normal Current Records logical model and one active index.
- Preserved Legacy Archive as separate lookup state.
- Preserved deferred proposal and write-transaction architecture.
- Scoped stale-authority search returned zero matches for legacy guide directories, direct PRODUCT Guidance source, Guidance Domain, Guidance Source, filename-stem guide identity, six-component current authority, and `## Abstract` extraction.
- The complete combined scoped diff was inspected without truncation.
- Scoped whitespace verification passed. LF-to-CRLF notices were advisory only.
- No scoped file is staged.
- Repository-wide cleanliness was not inspected or claimed.
- The standalone semantic responsibility validator was not executed because no operational invocation tool is available. No validator PASS was synthesized.
- No independent review occurred.
- No stage or commit was performed.

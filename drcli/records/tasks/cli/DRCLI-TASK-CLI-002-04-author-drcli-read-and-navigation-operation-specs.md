# DRCLI-TASK-CLI-002-04: Author DRCLI read and navigation operation specs

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: authoring
- **estimate**: 1d-2d
- **depends_on**:
  - DRCLI-TASK-CLI-002-01
- **outputs**:
  - spec:drcli.design_records_cli.operations.discovery_and_listing
  - spec:drcli.design_records_cli.operations.discovery_and_listing.record_scope_discovery
  - spec:drcli.design_records_cli.operations.discovery_and_listing.sequential_record_listing
  - spec:drcli.design_records_cli.operations.discovery_and_listing.tree_child_listing
  - spec:drcli.design_records_cli.operations.discovery_and_listing.tree_overview
  - spec:drcli.design_records_cli.operations.retrieval
  - spec:drcli.design_records_cli.operations.retrieval.batch_retrieval_results
  - spec:drcli.design_records_cli.operations.retrieval.exact_record_retrieval
  - spec:drcli.design_records_cli.operations.retrieval.h2_section_retrieval
  - spec:drcli.design_records_cli.operations.retrieval.retrieval_output_limits
  - spec:drcli.design_records_cli.operations.search
  - spec:drcli.design_records_cli.operations.search.lexical_match_model
  - spec:drcli.design_records_cli.operations.search.lexical_record_search
  - spec:drcli.design_records_cli.operations.search.search_results_and_limits
  - spec:drcli.design_records_cli.operations.reference_resolution

## Goal

Create transport-neutral DRCLI read, navigation, search, and reference-resolution operation contracts.

## Work

Write only:

- `operations/discovery-and-listing/**`
- `operations/retrieval/**`
- `operations/search/**`
- `operations/reference-resolution/**`

Primary baselines:

- current DRMCP `operations/discovery-and-listing/**`
- current DRMCP `operations/retrieval/**`
- current DRMCP `operations/search/**`
- old DRMCP `resolve-reference.md` plus current PRODUCT reference authority for reference resolution.

Preserve bounded outputs, deterministic ordering, exact selectors, no fuzzy repair, logical-tree semantics, and source-content limits where applicable.

Do not define concrete CLI flags or command spelling.
Do not write `operations/index.md`; T09 owns it.

## Done condition

- Discovery/listing, retrieval, section retrieval, search, and reference resolution have complete DRCLI operation contracts.
- The contracts are independent of MCP request envelopes.
- CLI mapping can be added later without changing semantic operation behavior.
- No assigned path overlaps another parallel writer.

## Verification

- Trace each Required Outcome read capability to an operation contract.
- Confirm current DRMCP semantics were preserved unless DRCLI-REQ-CLI-001 explicitly requires a difference.
- Confirm no root or sibling writer files changed.

## Evidence

- Inputs: DRCLI-REQ-CLI-001, DRCLI-TASK-CLI-002-01, current DRMCP discovery/listing, retrieval, and search operation Specifications, old DRMCP `resolve_reference`, and current PRODUCT traceability and V01 compatibility Specifications.
- Authored 15 Specification records under the four Task-owned operation path families. No independent review was performed.
- Required Outcome trace: scope discovery -> `record_scope_discovery`; record listing -> `sequential_record_listing`; logical-tree navigation -> `tree_child_listing` and `tree_overview`; exact record retrieval -> `exact_record_retrieval`; exact section retrieval -> `h2_section_retrieval`; scoped literal/pattern search -> `lexical_record_search`; canonical reference resolution -> `reference_resolution`.
- Shared retrieval contracts preserve selector occurrence order, partial selector success, exact H2 matching, and all-or-nothing UTF-8 source-content budgets.
- Search contracts preserve literal and RE2 matching, deterministic match ordering, bounded snippets, fixed per-record limiting, and no fuzzy repair.
- Discovery contracts preserve bounded deterministic listing and logical-tree behavior while consuming selected-repository current state instead of DRMCP configured roots.
- Reference resolution preserves exact current-first behavior while delegating canonical-ref and accepted V01 issued-ID semantics to PRODUCT-owned authorities.
- Verification confirmed path-derived DRCLI IDs, H1-adjacent metadata, required `Request` / `Response` / `Errors` sections for interface Contracts, and `Topics` tables for authored Overviews.
- Verification found no MCP request/tool wording or stale DRMCP refs in the authored operation paths after transport-neutral adaptation.
- All writes were limited to the four declared operation path families plus this Task record. `operations/index.md`, sibling writer outputs, DRMCP, and DRCLI-WORK-CLI-002 were not modified. Git operations were not used.

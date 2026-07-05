# DRMCP-TASK-MCP-016-08: Correct authoring-guidance source Specifications

- **id**: DRMCP-TASK-MCP-016-08
- **status**: done
- **date**: 2026-07-04
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-016-05
- **outputs**:
  - spec:drmcp.design_records_mcp.schema.authoring_guidance_source
  - spec:drmcp.design_records_mcp.tools.list_authoring_guides
  - spec:drmcp.design_records_mcp.tools.get_authoring_guidance

## Goal

Correct the stale authoring-guidance source path in the current DRMCP Guidance Specifications.

## Work

- Replace `docs/guides/`, `docs/guides/*.md`, and `docs/guides/<id>.md` source claims with `product/records/spec/design-records/authoring-standards/` and its direct Markdown files.
- Discover only Markdown files directly under `product/records/spec/design-records/authoring-standards/`; do not recurse.
- Derive guide ID from the Markdown filename stem.
- Derive guide title from the first H1 text.
- Derive guide abstract from the body of the `## What this is` section.
- Project guide content as the full Markdown content verbatim.
- Preserve Guidance as a separate read-only retrieval surface.
- Preserve exact guide-ID lookup, ASCII ordering, and public-response physical-path hiding.
- Preserve the T03 responsibility split among operation-specific Application Use Cases, the Guidance Domain, the Guidance Source port, and Infrastructure I/O Adapters.
- Do not add preload, cache, catalog-store, watcher, background-refresh, retained Guidance state, or independent stateful-subsystem requirements.
- Do not decide deferred authoring transaction architecture or unrelated operation-contract gaps.

## Done condition

- All three scoped Specifications name `product/records/spec/design-records/authoring-standards/` as the authoritative source.
- Discovery is limited to direct `*.md` files in that directory.
- No scoped Specification retains `docs/guides` or `## Abstract` as the current projection contract.
- Guide ID is the filename stem and is not converted to a canonical Spec ref.
- Guide title is the first H1 text.
- Guide abstract is the `## What this is` section body.
- Guide content is the full Markdown content verbatim.
- Exact ID lookup, ASCII ordering, and public-response physical-path hiding remain intact.
- The correction adds no architecture ADR, preload, cache, catalog store, watcher, background refresh, retained Guidance state, or new state lifecycle.
- The corrected Specifications are ready for T09 integrated review.

## Verification

- Confirm this Task contract uses `## What this is` extraction instead of `## Abstract` extraction.
- Confirm all three scoped Specifications name `product/records/spec/design-records/authoring-standards/`.
- Confirm discovery means direct `*.md` files only and does not recurse.
- Search the three scoped Specifications for stale `docs/guides` and `## Abstract` references.
- Confirm guide ID remains the filename stem and is not converted to a canonical Spec ref.
- Confirm title uses first H1 text, abstract uses the `## What this is` section body, and content uses full Markdown verbatim.
- Confirm exact lookup semantics, ASCII ordering, and public-response physical-path hiding remain intact.
- Confirm no preload, cache, catalog store, watcher, background refresh, retained Guidance state, or independent stateful subsystem was introduced.
- Confirm status, ID, parent, and `contract_class` metadata remain intact except for the required date updates.
- Confirm no other DRMCP Specification, ADR, T07 output, PRODUCT file, Work Item graph, review verdict, closure state, implementation artifact, staged file, or commit changed.
- Inspect the complete scoped diff and run the scoped whitespace check.
- Record whether the standalone semantic responsibility validator was available and executed.

## Evidence

- The repository-root `prompt_chappy.md` was the first repository file read in this session.
- `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational. Design Record reads and writes used filesystem fallback.
- The previous execution was blocked because the authoritative PRODUCT authoring-standard files do not contain `## Abstract`, while the old contract required that section.
- The user explicitly decided `guide.abstract = ## What this is section body`.
- The T08 contract was amended before Specification authoring. Goal, Task ID, outputs, dependency, and task type were preserved.
- The authoritative directory contains ten direct Markdown files. Every direct file contains a `## What this is` section.
- Updated `spec:drmcp.design_records_mcp.schema.authoring_guidance_source`.
- Updated `spec:drmcp.design_records_mcp.tools.list_authoring_guides`.
- Updated `spec:drmcp.design_records_mcp.tools.get_authoring_guidance`.
- Replaced the old `docs/guides/` source with `product/records/spec/design-records/authoring-standards/`.
- Discovery is limited to direct `product/records/spec/design-records/authoring-standards/*.md` files and is not recursive.
- Guide identity remains the direct Markdown filename stem and is not converted to a canonical Spec ref.
- Guide title is the first H1 text.
- Guide abstract is the body of the `## What this is` section.
- Guide content is the full Markdown content verbatim.
- Public responses continue to hide physical source paths.
- Exact ID lookup, no whitespace trimming, no case normalization, no user-supplied path lookup, `guide_not_found`, and ASCII lexical ordering remain intact.
- No preload, cache, catalog store, watcher, background refresh, retained Guidance state, or independent stateful Guidance subsystem was added.
- The T03 responsibility boundary among operation-specific Application Use Cases, Guidance Domain, the Guidance Source port, and Infrastructure I/O Adapters was preserved.
- No architecture decision, ADR, T07 output, PRODUCT file, Work Item graph, implementation, test, fixture, review verdict, or closure state was changed by this Task.
- Scoped searches over the three Specifications returned zero matches for `docs/guides` and zero matches for `## Abstract`.
- The complete four-file scoped diff was inspected without truncation.
- The final scoped whitespace check passed. Line-ending conversion warnings were advisory only.
- The standalone semantic responsibility validator was not executed because no operational invocation tool is available in this session. The validator was not treated as PASS.
- No staged files exist. Stage and commit were not performed.

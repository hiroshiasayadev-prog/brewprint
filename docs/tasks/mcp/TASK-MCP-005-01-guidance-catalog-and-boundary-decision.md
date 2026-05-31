# TASK-MCP-005-01: Guidance catalog and boundary decision

- **id**: TASK-MCP-005-01
- **status**: done
- **date**: 2026-05-31
- **work_item**: WORK-MCP-005

## Goal

Define the initial authoring-guidance catalog model and confirm the phase-1 boundary for MCP-based guidance retrieval.

## Scope

### Included

- Define the initial `docs/guides/` directory layout.
- Define guide ID generation from filename stem.
- Define required guide structure:
  - H1 as title
  - `## Abstract` as catalog summary
- Define `list_authoring_guides` output fields without exposing physical path.
- Define `get_authoring_guidance` retrieval behavior as Markdown content retrieval by guide ID.
- Confirm coexistence strategy with existing authoring guides.

### Excluded

- MCP implementation.
- Tool schema implementation.
- Migration of existing authoring guides.
- Removal of legacy guide locations.

## Decisions

### Phase 1 source model

Authoring guidance is sourced from Markdown files under:

`docs/guides/`

Guide ID is the filename without extension.

### Metadata extraction

- Title: first H1
- Abstract: contents of `## Abstract`

`list_authoring_guides` returns catalog metadata only:

- `id`
- `title`
- `abstract`

It does not return source file path. Guide source path is an internal resolution detail derived from guide ID.

`get_authoring_guidance` resolves guide ID internally and returns the Markdown content for that guide.

### Coexistence strategy

Phase 1 adopts migration strategy A.

- New guide files are introduced under `docs/guides/`.
- Existing authoring guides remain in place.
- Existing guides are not removed or relocated during this phase.
- Cleanup and consolidation may be handled by future work after operational validation.

## Done Condition

- Source model is documented.
- Catalog metadata extraction rules are documented.
- `list_authoring_guides` response excludes physical path.
- `get_authoring_guidance` returns Markdown content by guide ID.
- Coexistence strategy is documented.
- WORK-MCP-005 reflects the agreed boundary.

## Evidence

- Created `docs/guides/` as the phase-1 guide source directory.
- Created initial guide files:
  - `adr-authoring.md`
  - `spec-authoring.md`
  - `requirement-authoring.md`
  - `work-item-authoring.md`
  - `task-authoring.md`
  - `investigation-authoring.md`
  - `artifact-boundary.md`
- Updated `docs/doc-policy.md` so artifact-specific authoring rules point to `docs/guides/`.
- Preserved existing authoring guide and README files in their original locations.

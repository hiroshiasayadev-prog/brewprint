# V01-TASK-MCP-005-05: doc-policy / README guide ID reference update

- **id**: V01-TASK-MCP-005-05
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-MCP-005
- **source_requirement**: V01-REQ-MCP-005
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-005-04
- **outputs**:
  - docs/doc-policy.md
  - docs/requirements/README.md
  - docs/work-items/README.md
  - docs/tasks/README.md
  - docs/investigations/README.md
  - docs/adr-authoring-guide.md
  - docs/spec-authoring-guide.md
  - docs/guides/artifact-boundary.md
  - docs/requirements/mcp/REQ-MCP-005-project-authoring-guidance-retrieval-support.md
  - docs/work-items/mcp/WORK-MCP-005-project-authoring-guidance-retrieval-support.md

## Goal

Update startup and README-style documentation so authoring guidance is referenced by MCP guide ID instead of treating `docs/guides/*.md` source paths as public contract.

## Work

- Update `docs/doc-policy.md` to point artifact-specific authoring rules to Design Records MCP authoring guidance tools.
- Replace authoring guide path references in the doc-policy role table with guide IDs where appropriate.
- Add guide ID entry notes to README-style artifact entry points.
- Preserve existing Markdown guide / README locations during phase 1.
- Remove stale requirement wording that suggested source path is returned as public metadata.
- Remove stale requirement wording that suggested section / scope retrieval.
- Keep guide source path as internal implementation detail.
- Thin `docs/doc-policy.md` so artifact boundary details and stale structure snapshots are not duplicated there.

## Done condition

- `docs/doc-policy.md` uses guide IDs for authoring guide references.
- `docs/doc-policy.md` stays thin and does not duplicate artifact boundary details or stale structure snapshots.
- README-style entry points mention their authoring guide ID and `artifact-boundary` where applicable.
- No phase-1 document says guide source path is part of public response contract.
- No phase-1 requirement text claims section / scope retrieval is supported.
- Existing Markdown guide and README files remain in place.
- V01-WORK-MCP-005 reflects this task and moves to verification-oriented status.

## Verification

- Checked `docs/doc-policy.md` authoring guide table now uses guide IDs.
- Checked README-style entry points for guide ID notes:
  - `docs/requirements/README.md`
  - `docs/work-items/README.md`
  - `docs/tasks/README.md`
  - `docs/investigations/README.md`
  - `docs/adr-authoring-guide.md`
  - `docs/spec-authoring-guide.md`
- Checked `docs/guides/artifact-boundary.md` startup boundary uses `get_authoring_guidance` and guide ID language.
- Checked `V01-REQ-MCP-005` no longer states source path as returned guide metadata.
- Checked `V01-REQ-MCP-005` no longer claims section / scope retrieval.
- Checked `docs/doc-policy.md` no longer carries the stale current spec structure snapshot.
- Checked `docs/doc-policy.md` delegates artifact boundary details to `artifact-boundary`.
- Checked `V01-WORK-MCP-005` task list includes `V01-TASK-MCP-005-05` and status is `verification_pending`.

## Evidence

- Updated `docs/doc-policy.md` to reference authoring guide IDs:
  - `spec-authoring`
  - `adr-authoring`
  - `requirement-authoring`
  - `work-item-authoring`
  - `task-authoring`
  - `investigation-authoring`
  - `artifact-boundary`
- Thinned `docs/doc-policy.md` by removing duplicated artifact boundary details and stale fixed spec structure listing.
- Added guide ID notes to artifact README / legacy guide entry points.
- Updated `docs/guides/artifact-boundary.md` so startup docs point to `get_authoring_guidance` by guide ID.
- Updated `docs/requirements/mcp/REQ-MCP-005-project-authoring-guidance-retrieval-support.md` to remove stale source path metadata wording and section / scope retrieval wording.
- Updated `docs/work-items/mcp/WORK-MCP-005-project-authoring-guidance-retrieval-support.md` metadata and impact state.

# V01-TASK-MCP-014-02: Update spec and guidance for fields plus body create contract

- **id**: V01-TASK-MCP-014-02
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-MCP-014
- **source_requirement**: V01-REQ-MCP-014
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-MCP-014-01
- **outputs**:
  - Updated SPEC-design-records-mcp-tools propose_record_create contract
  - Updated authoring guidance/examples for body-as-sections create

## Goal

Update the public `propose_record_create` spec and authoring guidance so callers can use structured `fields` together with section-only `body` without hard-coding generated H1 or metadata.

## Work

- Rewrote `docs/spec/design-records-mcp/tools.md` body source rules for `propose_record_create`.
- Updated the preferred create example to use `fields` plus section-only `body`.
- Clarified `fields.id` compatibility behavior and `new` placeholder rendering responsibility.
- Added MCP create notes to requirement, work item, task, and ADR authoring guides so full artifact format examples are not mistaken for `fields + body` request bodies.
- Left directory README files unchanged because they only point to canonical guide IDs and do not present MCP body examples.
- Left Go implementation unchanged; implementation remains owned by `V01-TASK-MCP-014-03`.

## Done condition

- `propose_record_create` spec no longer says `fields` forbids `body`.
- Spec documents valid create combinations:
  - `fields`
  - `fields` plus section-only `body`
  - legacy full-record `body` or `body_cache_id` without `fields`
- Spec keeps `body` plus `body_cache_id` invalid.
- Spec states MCP owns H1 and metadata rendering when `fields` is present.
- Spec states caller-supplied `body` in `fields + body` mode must omit H1, metadata block, metadata `id`, and guessed server-resolved ID.
- Spec states `new` placeholder creates render generated H1 and metadata with `target.resolved_id`.
- Spec keeps `fields.id` compatibility-tolerated only for exact top-level IDs that match after normalization; mismatches and `fields.id` with `new` placeholder creates remain invalid.
- Authoring guidance distinguishes full artifact file shape from MCP `fields + body` body payload shape.

## Verification

- `go test ./internal/designrecords ./internal/designrecordsmcp`
- `go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml --format json`
- Design Records MCP `validate_records(kind=task, id_range=V01-TASK-MCP-014-01..V01-TASK-MCP-014-04)`

## Evidence

- Updated `docs/spec/design-records-mcp/tools.md`:
  - common body source rules now allow `propose_record_create` `fields + body` when `body` is content sections only;
  - request contract table now lists structured metadata only, structured metadata plus content sections, and legacy full-record body create;
  - preferred example uses `fields + body` and the `body` begins at `## Goal`;
  - error handling no longer treats `fields + body` as invalid.
- Updated `docs/guides/requirement-authoring.md`, `docs/guides/work-item-authoring.md`, `docs/guides/task-authoring.md`, and `docs/guides/adr-authoring.md` with MCP create notes that prohibit H1, metadata block, metadata `id`, and guessed resolved ID in `fields + body` bodies.
- Kept `docs/requirements/README.md`, `docs/work-items/README.md`, and `docs/tasks/README.md` unchanged because they do not document MCP request bodies.
- `go test ./internal/designrecords ./internal/designrecordsmcp` passed on 2026-06-02.
- `go run ./cmd/brewprint validate --format json` failed because the CLI requires `--yaml-root`.
- Existing MCP/self-hosting task docs use `docs\uc\002-brewprint-self-hosting\yaml` as the relevant YAML root.
- `go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml --format json` passed on 2026-06-02 with `error_count: 0` and `warning_count: 0`.
- Design Records MCP `validate_records(kind=task, id_range=V01-TASK-MCP-014-01..V01-TASK-MCP-014-04)` passed with no diagnostics.

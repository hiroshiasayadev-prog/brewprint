# V01-TASK-MCP-008-04: MCP tools spec reflection

- **id**: V01-TASK-MCP-008-04
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-MCP-008
- **source_requirement**: V01-REQ-MCP-008
- **estimate**: 1d-2d
- **depends_on**:
  - V01-TASK-MCP-008-03
- **outputs**:
  - Updated `SPEC-design-records-mcp-tools` authoring transaction contract
  - Spec reflection notes from V01-ADR-093

## Goal

Accepted `V01-ADR-093` の authoring transaction model を `SPEC-design-records-mcp-tools` に反映し、Design Records MCP の public tool contract として proposal / accept / cache / validation semantics を定義する。

## Work

- `V01-ADR-093` を読み、tools spec に反映すべき contract と spec に委ねられた details を抽出する。
- `V01-REQ-MCP-008`, `V01-WORK-MCP-008`, and `V01-TASK-MCP-008-01` through `V01-TASK-MCP-008-03` を確認し、要求・レビュー結果・accept済み判断と矛盾しないようにする。
- `docs/spec/design-records-mcp/tools.md` の現行 read-only / write tool policy を更新する。
- 必要に応じて `docs/spec/design-records-mcp/schema.md` に proposal / cache / metadata block / section selector / diagnostics の schema-level contract を追加する。
- Spec には以下を明確に定義する。
  - authoring transaction tool surface;
  - propose / get / accept / discard flow;
  - proposal response fields including proposal ID, resolved target, expiry, diff, validation result, and note;
  - accept response semantics including written/not-written result;
  - base-state / staleness / collision guard;
  - `new` placeholder resolution and update-time rejection;
  - task create parent context and reciprocal metadata handling;
  - body source rule for `body` / `body_cache_id`, including neither-allowed cases;
  - proposal/body cache expiry and unknown/expired diagnostics;
  - kind-specific metadata block replacement;
  - named-section replacement with exactly-one-section safety condition;
  - MVP exclusions, especially general-purpose multi-record atomic transactions with rollback semantics.
- Do not implement the MCP tools in this task.

## Done condition

- Accepted V01-ADR-093 の decisions が tools spec に反映されている。
- Spec が V01-ADR-093 と矛盾していない。
- Public request / response / error behavior が implementation task に渡せる粒度で記述されている。
- Read-only policy が古いまま残って authoring write support と矛盾していない。
- V01-WORK-MCP-008 の後続 implementation planning に進める状態になっている。

## Verification

- V01-ADR-093 の decision list と spec reflection の対応を Evidence に記録する。
- 可能であれば `validate_records` で `V01-ADR-093`, `V01-WORK-MCP-008`, and this task の metadata を確認する。
- Repo-local command execution が必要な validation / rendering / tests は Codex / implementation agent に委譲するか、未実施として明示する。

## Evidence

2026-06-02: Accepted V01-ADR-093 was reflected into the public Design Records MCP spec.

Files modified:

- `docs/spec/design-records-mcp/tools.md`
- `docs/spec/design-records-mcp/schema.md`
- `docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md`

V01-ADR-093 decision mapping to spec sections:

- Propose -> accept flow, proposal creation no-write rule, accept as the write operation, get/discard proposal, and 3 day proposal expiry: `tools.md` `## Authoring transaction model`.
- MVP authoring tool surface: `tools.md` `## Tool set`, `## propose_record_create`, `## propose_record_update`, `## get_proposed_write`, `## accept_proposed_write`, and `## discard_proposed_write`.
- Proposal response fields: `tools.md` `### Common authoring response fields`.
- Accept response semantics, `written: true/false`, stale/changed target/ID collision rejection before writing, expired/discarded/unknown diagnostics, repair guidance, and no automatic rollback after accepted post-write validation failure: `tools.md` `## accept_proposed_write`.
- Artifact-oriented input and path-as-transparent-output boundary: `tools.md` `## Authoring transaction model`; schema concept in `schema.md` `### Authoring target identity`.
- `new` placeholder create support, update rejection, existing-index ID resolution, accept-time ID availability re-check, and task parent context: `tools.md` `## propose_record_create` and `## propose_record_update`.
- Workflow reciprocal metadata boundary: `tools.md` `## propose_record_create` and `## Authoring write boundary`.
- Body source and body cache behavior: `tools.md` `### Body source and body cache`; schema concept in `schema.md` `### Body cache model`.
- Metadata block replacement and named section replacement: `tools.md` `## propose_record_update`; schema concepts in `schema.md` `### Metadata block replacement target` and `### Section selector model`.
- MVP exclusions: `tools.md` `## Authoring write boundary`.

Summary of new / updated tool contracts:

- Added public contracts for `propose_record_create`, `propose_record_update`, `get_proposed_write`, `accept_proposed_write`, and `discard_proposed_write`.
- Defined proposal fields: proposal ID, state, operation, target kind, requested/resolved target identity, transparent path output, expiry, 3 day retention, unified diff, validation result, diagnostics, and no-write note.
- Defined accept fields: proposal ID, state, `written`, `files_written`, validation, repair guidance, and diagnostics.
- Preserved `suggest_next_record` as read-only suggestion and separated it from authoring create.

Summary of schema concepts:

- Added schema-level concepts for authoring target identity, proposal model, body cache model, metadata block replacement targets, section selector model, and authoring diagnostics.
- Left existing record model and read/navigation response schema unchanged except for adding V01-ADR-093 as a spec dependency.

Stale read-only wording:

- Updated the stale `tools.md` top-level read-only statement and old `Write tool policy`.
- `tools.md` now states existing read/navigation/guidance tools remain read-only, while authoring writes use proposal-first tools and only `accept_proposed_write` may write repository files.

Existing read/navigation tools:

- No semantic changes were made to `list_records`, `get_record`, `get_records`, `validate_records`, `resolve_reference`, `list_authoring_guides`, `get_authoring_guidance`, or `suggest_next_record`.
- Changes near those sections are limited to dependency metadata and documentation clarifying relationship with authoring tools.

Validation / command results:

- `validate_records(kind: decision, id_range: V01-ADR-093..V01-ADR-093)` returned `ok: true`.
- `validate_records(kind: work_item, id_range: V01-WORK-MCP-008..V01-WORK-MCP-008)` returned `ok: true`.
- `validate_records(kind: task, id_range: V01-TASK-MCP-008-04..V01-TASK-MCP-008-04)` returned `ok: true`.
- Targeted stale wording check against `tools.md` / `schema.md`: `rg -n "MVP tool は read-only|MVP では write 系 tool を提供しない|Write tool policy" docs/spec/design-records-mcp/tools.md docs/spec/design-records-mcp/schema.md` returned no matches.
- Targeted Markdown fence check: `tools.md` and `schema.md` have balanced code fence markers.
- `git diff --check -- docs/spec/design-records-mcp/tools.md docs/spec/design-records-mcp/schema.md docs/tasks/mcp/TASK-MCP-008-04-mcp-tools-spec-reflection.md` returned no whitespace errors. Git reported line-ending conversion warnings for the two spec files.
- No implementation files were modified by this task; no implementation tests were required.

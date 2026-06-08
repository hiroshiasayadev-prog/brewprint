# V01-TASK-MCP-008-02: ADR draft for authoring transaction model

- **id**: V01-TASK-MCP-008-02
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-MCP-008
- **source_requirement**: V01-REQ-MCP-008
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-MCP-008-01
- **outputs**:
  - Proposed ADR for Design Records MCP authoring transaction model
  - ADR drafting evidence based on V01-TASK-MCP-008-01

## Goal

`V01-TASK-MCP-008-01` の evidence を踏まえて、Design Records MCP authoring transaction model を固定する proposed ADR を起草する。

## Work

- `V01-TASK-MCP-008-01` の Evidence を読み、current write capability gap と ADR decision points を確認する。
- ADR authoring guide に従い、ADR が所有する判断と spec に委ねる schema details を分離する。
- `suggest_next_record` または既存 ADR 一覧で次の ADR 番号を確認する。
- 新規 ADR を proposed status で起草する。
- ADR には少なくとも以下の判断を含める。
  - Design Records MCP に authoring write capability を追加する。
  - Immediate write ではなく propose → accept transaction model を採用する。
  - Proposal creation は repository files を変更しない。
  - Write inputs は filesystem path ではなく artifact identity / kind / domain / section / structured fields を primary input とする。
  - Create operations は `new` placeholder を許可し、MCP 側で採番する。
  - Update operations では `new` placeholder を拒否する。
  - Body source は、必要時に `body` または `body_cache_id` のどちらか一方のみとする。
  - Proposal / body cache retention は 3 days とする。
  - MVP updates は add/remove convenience ではなく set-only とする。
  - Response semantics は write failure と validation failure を分離し、accept result は `written: true/false` 相当を明示する。
- ADR には以下を抱え込まない。
  - Exact tool names
  - Full JSON schemas
  - Complete error code list
  - Proposal ID format
  - Cache storage layout
  - Section matching algorithm details
  - Implementation package structure
  - Exhaustive test matrix

## Done condition

- 新規 ADR が proposed status で作成されている。
- ADR が `V01-REQ-MCP-008`, `V01-WORK-MCP-008`, `V01-TASK-MCP-008-01`, and this task を evidence / impact として適切に参照している。
- ADR が Design Records MCP authoring transaction model の主要判断を所有している。
- ADR が spec-owned schema details や implementation procedure を抱え込みすぎていない。
- 後続 `V01-TASK-MCP-008-03` で ADR review に進める状態になっている。

## Verification

- 作成した ADR を読み返し、ADR authoring guide の responsibility boundary に反していないことを確認する。
- 可能であれば `validate_records` で新規 ADR と this task の metadata を確認する。
- Repo-local command execution が必要な検証は Codex / implementation agent に委譲するか、未実施として明示する。

## Evidence

2026-06-01: ADR draft created.

Created:

- `V01-ADR-093`: Design Records MCP authoring transaction model

ADR path suggested by `suggest_next_record`:

- `docs/adr/093-design-records-mcp-authoring-transaction-model.md`

Draft status:

- `proposed`

The ADR captures the durable authoring transaction decisions:

- Design Records MCP will add authoring write capability.
- Writes use propose -> accept rather than immediate write.
- Proposal creation must not modify repository files.
- Write inputs are artifact-oriented rather than path-first.
- Create operations may use `new` placeholder IDs, resolved by MCP.
- Update operations reject `new` IDs.
- Task creation requires parent-aware context.
- Body source is exactly one of `body` or `body_cache_id` when required.
- Proposal/body cache retention is 3 days.
- MVP updates are set-only for front matter and named sections.
- Write failure and validation failure are distinct response states.

The ADR intentionally leaves exact tool names, JSON schemas, proposal ID format, cache storage layout, diagnostics, and implementation package structure to the Design Records MCP tools spec and implementation tasks.

Verification:

- Re-read `V01-TASK-MCP-008-01` Evidence before drafting.
- Read `adr-authoring` guidance before creating this task and ADR.
- Confirmed next decision ID with `suggest_next_record`, which returned `V01-ADR-093` and `existing_max_id: V01-ADR-092`.
- ADR responsibility boundary checked manually: it records design decisions and rationale, not full tool schemas or implementation procedure.
- Codex implementation / agent-usability review returned `Needs revision before Opus review`.
- Incorporated Codex major findings into V01-ADR-093: accept-time staleness / collision guard, workflow reciprocal metadata handling, unambiguous named-section replacement, kind-specific metadata block terminology, cache expiry diagnostics, and no automatic rollback after accepted post-write validation failure.


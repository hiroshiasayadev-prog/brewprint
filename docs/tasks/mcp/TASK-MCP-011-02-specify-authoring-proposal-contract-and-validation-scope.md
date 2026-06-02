# TASK-MCP-011-02: Specify authoring proposal contract and validation scope

- **id**: TASK-MCP-011-02
- **status**: done
- **date**: 2026-06-02
- **work_item**: WORK-MCP-011
- **source_requirement**: REQ-MCP-011
- **estimate**: 1d
- **depends_on**:
  - TASK-MCP-011-01
- **outputs**:
  - Updated public spec contract for authoring proposal create input normalization
  - Updated public spec contract for authoring proposal validation scope isolation
  - Spec evidence that separates REQ-MCP-011 and REQ-MCP-012 behavior from implementation fixes

## Goal

Update the Design Records MCP public spec contract for authoring proposal behavior using the evidence from `TASK-MCP-011-01`.

This task should make the expected public behavior explicit before implementation changes begin, so the later patch does not rely on undocumented dogfooding observations.

## Work

- Inspect the current public contract in `SPEC-design-records-mcp-tools` and `SPEC-design-records-mcp-schema` for authoring proposal create/update operations.
- Specify the accepted behavior for `propose_record_create` when `body`, `body_cache_id`, and structured `fields` are supplied.
- Specify whether top-level `id` is the canonical create target ID and whether `fields.id` is forbidden, optional, or required.
- Specify mismatch handling when both top-level `id` and `fields.id` are supplied.
- Specify domain normalization behavior between canonical ID domains such as `MCP` and repository path domains such as `mcp`.
- Specify proposal validation scope rules for create and update operations, including the affected record set and separation of proposal-local diagnostics from unrelated repository health diagnostics.
- Check whether `ADR-093` remains consistent with the specified authoring transaction contract. If the transaction model itself must change, record the ADR follow-up instead of silently changing the spec beyond the ADR boundary.
- Do not implement the Design Records MCP code changes in this task.
- Do not clean up unrelated existing repository validation errors in this task.

## Done condition

- `SPEC-design-records-mcp-tools` and/or `SPEC-design-records-mcp-schema` define the public behavior for all four failures reproduced in `TASK-MCP-011-01`:
  - `body` / `body_cache_id` / `fields` content-source selection;
  - top-level `id` versus `fields.id` responsibility;
  - domain case normalization or documented rejection;
  - proposal-local validation scope and unrelated diagnostic separation.
- The spec text distinguishes contract decisions from current implementation bugs.
- The spec text preserves the boundary between `REQ-MCP-011` input normalization and `REQ-MCP-012` validation scope isolation.
- Any ADR consistency concern is either ruled out with evidence or captured as a follow-up before implementation proceeds.
- The resulting spec contract is precise enough for `TASK-MCP-011-03` to implement without guessing.

## Verification

Expected checks include:

```powershell
validate_records(id_range={from: "TASK-MCP-011-01", to: "TASK-MCP-011-02"}, kind: task)
validate_records(id: "WORK-MCP-011", kind: work_item)
```

After the spec update, run targeted validation for the edited spec records as supported by Design Records MCP.

## Evidence

2026-06-02 Codex updated the public Design Records MCP authoring proposal contract without implementation or test changes.

Spec files updated:

- `SPEC-design-records-mcp-tools` (`docs/spec/design-records-mcp/tools.md`)
- `SPEC-design-records-mcp-schema` (`docs/spec/design-records-mcp/schema.md`)

Contract choices recorded:

- `body` / `body_cache_id` remain mutually exclusive body sources; supplying both keeps the existing `invalid_body_source` behavior.
- `fields` and full Markdown body input are mutually exclusive create content sources. `propose_record_create` must reject `fields` together with `body` or `body_cache_id` as `invalid_request`; the spec intentionally defines no precedence.
- Top-level create `id` is the canonical target ID input. `fields.id` is not required and must not be required by rendering. If supplied with an exact top-level ID it must match after canonical ID normalization; if supplied with a `new` placeholder request it is invalid because the final ID is server-resolved.
- Workflow create `domain` is compared with the ID domain case-insensitively. Canonical ID domains remain uppercase, while repository path domains use lowercase normalized directories.
- Proposal-local validation is scoped to the affected record set: the proposed target record plus any related records actually modified in the same proposal, such as required reciprocal metadata updates. Unrelated repository diagnostics must not be mixed into proposal-local `validation.diagnostics`; optional repository health reporting must be separated.
- Proposal-time and accept-time pre-write diagnostics must be reproducible by the same `validate_records` rules over the same affected record set in the same candidate state, or after that state is accepted/materialized.

Judgment notes:

- `REQ-MCP-011` owns input normalization: body/fields, top-level `id` versus `fields.id`, and domain normalization.
- `REQ-MCP-012` owns validation scope isolation and unrelated diagnostic separation.
- `ADR-093` remains consistent. It already defines artifact identity as the primary write input, body/body_cache mutual exclusion, proposal validation result semantics, and tools-spec ownership for concrete schemas. No ADR follow-up is needed for this contract update.
- The spec `design_record.depends_on` metadata was not expanded with `REQ-*` / `TASK-*` because the current schema only allows `ADR-*` / `SPEC-*` / `INV-*` dependencies for decision/spec records.

Validation observed after the spec update:

- `validate_records(kind=task, id_range={from:"TASK-MCP-011-01", to:"TASK-MCP-011-02"})`: `ok:true`, no diagnostics.
- `validate_records(kind=work_item, id_range={from:"WORK-MCP-011", to:"WORK-MCP-011"})`: `ok:true`, no diagnostics.
- `validate_records(kind=requirement, id_range={from:"REQ-MCP-011", to:"REQ-MCP-012"})`: `ok:true`, no diagnostics.
- `validate_records(kind=spec)`: `ok:true`, no diagnostics.

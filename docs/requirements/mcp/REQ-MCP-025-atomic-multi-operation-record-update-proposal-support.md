# REQ-MCP-025: atomic multi-operation record update proposal support

- **id**: REQ-MCP-025
- **status**: accepted
- **date**: 2026-06-05
- **source_refs**:
  - REQ-MCP-020
  - REQ-MCP-022
  - REQ-MCP-024
  - TASK-DATA-011-03
- **work_items**:
  - WORK-MCP-026

## Requirement

Design Records MCP authoring update operations MUST support an atomic multi-operation update proposal for a single existing record.

A caller MUST be able to combine multiple supported update operations into one retained proposal, especially a named section replacement and a metadata field-level update. The combined proposal MUST be validated against the final record state after all operations are applied, rather than validating each intermediate state independently.

This requirement assumes the field-level metadata update capability from `REQ-MCP-020` for metadata-side operations. The intended metadata operation is `metadata_field_replace` / `metadata_fields_replace`, not requiring callers to resubmit the complete metadata block.

## Evidence

Closing workflow artifacts commonly requires writing Evidence and then changing status to `done`.

For `TASK-DATA-011-03`, attempting to update `status` first hit metadata replacement friction and then the done-state Evidence gate:

```text
TASK-DATA-011-03 を done に更新します。まずメタデータ更新。
```

```text
id フィールドも必須でした。追加して再試行します。
```

```text
status done には Evidence が必須です。順序を逆にして、Evidence を先に書いてから status を更新します。
```

The Evidence gate itself is desirable, but the authoring API forces a high-frequency close operation into multiple proposal/accept cycles. That increases input size, retry surface, and ordering mistakes.

## Required Outcome

The update contract supports a single retained proposal containing multiple update operations for the same record.

Acceptance criteria:

- A caller can replace `## Evidence` and set `status: done` in the same proposal.
- Metadata-side status update uses the field-level metadata update semantics from `REQ-MCP-020`.
- Section-side update can use `named_section_replace`, including the heading-safe normalization behavior from `REQ-MCP-022` when applicable.
- Validation is performed against the final record content after all operations are applied.
- Done-state validation can pass when Evidence is added and status is changed to `done` within the same atomic proposal.
- The proposal diff includes all changes in one reviewable diff.
- Accept is atomic: either all operations are written or none are written.
- Operation ordering is deterministic and documented.
- Conflicting operations targeting the same metadata field or same section fail safely with diagnostics.

## Explicitly Excluded Scope

- Multi-record transactions.
- Automatically closing parent work items or source requirements.
- Bypassing done-state Evidence validation.
- Requiring full metadata block replacement for the status update path.
- Arbitrary text/string replacement outside the supported operation set.

## Boundary

This requirement belongs to the Design Records MCP authoring update contract. It addresses single-record atomicity and validation ordering for common lifecycle updates. It depends on `REQ-MCP-020` for field-level metadata update semantics and complements `REQ-MCP-022` for safe section replacement body handling.

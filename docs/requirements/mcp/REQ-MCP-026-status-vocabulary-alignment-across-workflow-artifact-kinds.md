# REQ-MCP-026: status vocabulary alignment across workflow artifact kinds

- **id**: REQ-MCP-026
- **status**: captured
- **date**: 2026-06-06
- **source_refs**:
  - REQ-MCP-023
- **work_items**:
  - WORK-MCP-022

## Requirement

Status vocabulary for workflow artifact kinds MUST be aligned so that valid values are predictable without trial-and-error probing.

Current vocabularies mix conventions: task uses short casual tokens (`todo`, `doing`, `blocked`, `done`), while work_item uses verbose compound tokens (`implementation_pending`, `in_progress`). This mismatch causes LLM callers to default to natural-English equivalents (`in_progress`, `not_started`) that are invalid for the target kind.

The alignment must satisfy:

- Status tokens across kinds follow a consistent naming convention where lifecycle phases are semantically equivalent.
- Tokens are short enough to recall without reference and unambiguous enough not to collide with common English phrases.
- Existing records are migrated to any updated canonical values.
- Parser, validation, authoring tooling, and authoring guides reflect the aligned vocabulary.

## Evidence

During TASK-DATA-011 authoring and the investigation that produced REQ-MCP-023, the following failures occurred in sequence:

```text
in_progress はタスクkindで無効でした。破棄して not_started で再起票します。
```

```text
not_started もタスクkindでは無効なようです。有効なステータスを確認します。
```

```text
有効なステータスは todo / doing / blocked / done でした。doing で再起票します。
```

The root cause is that `doing` and `todo` are not the first terms an LLM produces when describing an active or not-yet-started task. `in_progress` and `not_started` are higher-prior natural-language equivalents.

REQ-MCP-023 proposes synonym repair as a band-aid. This requirement addresses the underlying vocabulary design so that the band-aid is not the permanent solution.

## Required Outcome

Aligned status vocabulary is defined, documented, and implemented across all workflow artifact kinds.

Acceptance criteria:

- A vocabulary alignment review compares status tokens across `decision`, `requirement`, `work_item`, and `task` kinds and identifies gaps.
- For each gap, alignment options are evaluated on: semantic precision, recall likelihood for LLM callers, migration cost, and impact on existing tooling.
- Updated canonical values replace ambiguous or misaligned tokens where alignment is safe and semantically preserving.
- All existing records in the repository are migrated to the updated canonical values.
- Schema, parser, validator, and authoring tooling are updated consistently.
- Authoring guides for each kind surface the valid status vocabulary explicitly.
- After alignment, the synonym repair behavior introduced by REQ-MCP-023 may be narrowed or removed if the underlying confusion source is eliminated.

## Explicitly Excluded Scope

- Changing the proposal / accept flow or adding new authoring operations.
- Vocabulary changes that alter the semantic meaning of a lifecycle phase rather than its token.
- Removing validation guards for invalid status values.
- Automatically migrating caller-side code outside this repository.

## Boundary

This requirement belongs to the Design Records MCP schema and vocabulary contract. It depends on REQ-MCP-023 for the evidence of confusion and the immediate diagnostic improvement. It supersedes the need for a permanent synonym mapping table once vocabulary alignment is complete.

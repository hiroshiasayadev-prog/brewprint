# Task Authoring Guide

## Abstract

Task artifact を起票・更新するときの実践ルールを定める。Task は work item を前進させる短期 concrete work、完了条件、個別 status、verification evidence を所有する。

## Migration Note

This section is a maintainer note about phase-1 guide extraction history.
It is not an instruction to read the original file, and it is not part of the public authoring guidance retrieval contract.

Extracted from:

- `docs/tasks/README.md`

## Purpose

Task は、work item を前進させるための短期 concrete work、完了条件、個別 status、verification evidence を記録する artifact である。

Task は短期に閉じられる具体作業を所有する。

## Granularity

新規 task は、原則として着手後 `0.5d` から `3d` 程度で完了判定できる単位に分割する。

以下に該当する場合は split を検討する。

- 3日を明らかに超える見込みである。
- 独立した判断、仕様更新、実装、検証を複数含む。
- outputs または done condition を一つの完了判定として説明しにくい。
- 着手時に「どこまで終われば今日は閉じられるか」が明確でない。

## ID and Layout

Task ID:

```text
TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>
```

File path:

```text
docs/tasks/<domain>/TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>-<slug>.md
```

## Format

```markdown
# TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>: <title>

- **id**: TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>
- **status**: not_started / in_progress / blocked / done
- **date**: YYYY-MM-DD
- **work_item**: WORK-<DOMAIN>-NNN
- **source_requirement**: REQ-<DOMAIN>-NNN
- **estimate**: <rough estimate; usually within 0.5d-3d>
- **depends_on**:
  - TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>
- **outputs**:
  - <output>

## Goal

## Work

## Done condition

## Verification

## Evidence
```

`id` は H1 の task ID と一致させる。

MCP create note:
The Markdown above is the full artifact file shape.
When calling `propose_record_create` with `fields` plus `body`, put metadata in `fields` and make `body` content sections only, starting at `## Goal`.
Do not include the H1, metadata block, metadata `id`, or guessed server-resolved ID in that MCP `body`.
Pass `id: TASK-<DOMAIN>-<WORK-SEQUENCE>-new` (e.g. `TASK-MCP-014-new`); the MCP resolves the next number server-side. Never hardcode a guessed task number in `body` or `fields`.

Always include all sections in `body`, even those that will be filled in later.
Use `TBD` as a placeholder for sections that are not yet written (e.g. `## Evidence\n\nTBD`).
This ensures `propose_record_update` `named_section_replace` can target those sections later without falling back to filesystem edits.

Use exact `TASK-<DOMAIN>-<WORK-SEQUENCE>-NN` only when the specific task ID is intentional. If an exact ID would skip the next task sequence within the parent work item scope, `propose_record_create` may return a non-blocking `exact_id_sequence_gap` info diagnostic. Prefer `TASK-<DOMAIN>-<WORK-SEQUENCE>-new` when no reserved ID is required.

`work_item`、`source_requirement`、`depends_on` は `WORK-*` / `REQ-*` / `TASK-*` の ID-as-ref を用いる。

Required metadata は Design Records MCP validation の対象である。

- `id` / `status` / `date` / `work_item` / `source_requirement` / `estimate` は存在し、non-empty でなければならない。
- `date` は `YYYY-MM-DD` format とする。
- `depends_on` / `outputs` は list field として存在しなければならない。
- `depends_on` / `outputs` は empty list を許容する。
- `depends_on` / `outputs` の empty item は validation error とし、metadata diagnostic category は `empty_required_metadata` とする。

## Status

| status | meaning |
|---|---|
| `not_started` | 未着手 |
| `in_progress` | 作業中 |
| `blocked` | blocker により停止中 |
| `done` | Done condition を満たした |

## Required Sections

新規 task は原則として以下の section を持つ。

- `Goal`: この task 単独で達成すること。
- `Work`: 実施する具体作業。
- `Done condition`: status を `done` にできる判定条件。
- `Verification`: 完了判断の確認方法。
- `Evidence`: 実行結果、判断結果、検証結果。

MCP validation gate:
When a task has `status: done`, all sections listed above must exist and be non-empty.

Task が investigation、ADR、spec、implementation、fixture 等を作成・更新する場合、それらは outputs または evidence として記録する。

Task 本文がそれら artifact の source of truth になってはならない。

## Responsibility Boundary

Task が所有するもの:

- 短期 concrete work
- 完了条件
- 個別 status
- verification evidence
- outputs
- dependencies

Task が所有しないもの:

- requirement の正本
- 設計判断
- 現行仕様
- work item 全体の到達点
- task graph

## Status Ownership

- 個々の task の完了状態の正本は、当該 task artifact の `status` である。
- Parent work item は requirement 解消フロー全体の処理段階を status として保持する。
- Work item 本文に、配下 task の完了 checkbox または status copy を手動複製しない。

## Canonical Reference Boundary

Workflow artifact の identity と relation は ID-as-ref で扱う。

| relation | canonical reference |
|---|---|
| task -> source requirement | `REQ-*` |
| task -> parent work item | `WORK-*` |
| task -> dependency task | `TASK-*` |

Physical path は workflow artifact 間の canonical relation として support しない。

`req:` / `work:` / `task:` の semantic prefix は導入しない。

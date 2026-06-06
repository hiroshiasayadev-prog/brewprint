# Work Item Authoring Guide

## Abstract

Work item を起票・更新するときの実践ルールを定める。Work item は source requirement を解消する到達点、横断進捗、影響範囲、task graph を所有し、個別 task status の正本を所有しない。

## Migration Note

Maintainer note only. This is not an instruction to read another file, and it is not part of the public guide retrieval contract.

Extracted from:

- `docs/work-items/README.md`

## Purpose

Work item は、accepted または追跡対象となった requirement を完了させるための到達点、作業フロー全体、横断進捗、影響範囲、task graph を記録する artifact である。

すべての work item は source requirement を必ず持つ。

## ID and Layout

Work item ID:

```text
WORK-<DOMAIN>-NNN
```

File path:

```text
docs/work-items/<domain>/WORK-<DOMAIN>-NNN-<slug>.md
```

## Format

```markdown
# WORK-<DOMAIN>-NNN: <title>

- **id**: WORK-<DOMAIN>-NNN
- **status**: not_started / in_progress / blocked / done
- **date**: YYYY-MM-DD
- **source_requirement**: REQ-<DOMAIN>-NNN
- **impact_refs**:
  - <canonical reference>
- **tasks**:
  - TASK-<DOMAIN>-NNN-01

## Goal

## Boundary

## Impact Scope

## Task flow

## Task Candidates

## Completion Condition
```

`id` は H1 の work item ID と一致させる。

MCP create note:
The Markdown above is the full artifact file shape.
When calling `propose_record_create` with `fields` plus `body`, put metadata in `fields` and make `body` content sections only, starting at `## Goal`.
Do not include the H1, metadata block, metadata `id`, or guessed server-resolved ID in that MCP `body`.
Pass `id: WORK-<DOMAIN>-new` (e.g. `WORK-MCP-new`); the MCP resolves the next number server-side. Never hardcode a guessed work item number in `body` or `fields`.

Use exact `WORK-<DOMAIN>-NNN` only when that specific work item ID is intentional. If an exact ID would skip the next domain-scoped work item sequence, `propose_record_create` may return a non-blocking `exact_id_sequence_gap` info diagnostic. Prefer `WORK-<DOMAIN>-new` when no reserved ID is required.

Required metadata は Design Records MCP validation の対象である。

- `id` / `status` / `date` / `source_requirement` は存在し、non-empty でなければならない。
- `date` は `YYYY-MM-DD` format とする。
- `impact_refs` / `tasks` は list field として存在しなければならない。
- `impact_refs` / `tasks` は empty list を許容する。
- `impact_refs` / `tasks` の empty item は validation error とし、metadata diagnostic category は `empty_required_metadata` とする。

## Status

| status | meaning |
|---|---|
| `not_started` | 未着手 |
| `in_progress` | 作業中 |
| `blocked` | 依存判断・外部要因で停止中 |
| `done` | 必要な反映・検証が完了 |

## Responsibility Boundary

Work item が所有するもの:

- source requirement を解消する到達点
- 作業フロー全体
- 横断進捗
- 影響範囲
- task graph
- close 判定

Work item が所有しないもの:

- 現行仕様本文
- 設計判断の長い理由
- 具体的な作業手順
- 個別 task の status 正本

## Task Flow

Work item は Mermaid `flowchart` で task の順序・分岐・blocker・並列可能性を示せる。

Task flow は進行構造を示す view であり、個別 task の status の正本ではない。

## Completion Condition

Work item は requirement を解消するための全体的な完了条件を持つ。

配下 task がすべて `done` でも、必要な evidence、status synchronization、requirement への結果反映が未完了であれば、work item を `done` としない。

## Progress Boundary

- Work item の `status` は requirement 解消フロー全体の処理段階を表す。
- 各 task の完了状態の正本は task artifact の `status` とする。
- Work item 本文に、配下 task の完了 checkbox または status copy を手動複製しない。
- Status 集約 view が必要な場合は、将来の MCP support による derived projection として扱う。

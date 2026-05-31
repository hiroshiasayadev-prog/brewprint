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
- **status**: not_started / decision_pending / design_spec_pending / internal_design_pending / yaml_pending / implementation_pending / fixture_pending / verification_pending / done / blocked
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

`source_requirement` と `tasks` は `REQ-*` / `TASK-*` の ID-as-ref を用いる。

## Status

| status | meaning |
|---|---|
| `not_started` | 未着手 |
| `decision_pending` | ADR 等の設計判断待ち |
| `design_spec_pending` | design spec 反映待ち |
| `internal_design_pending` | internal design 反映待ち |
| `yaml_pending` | brewprint DSL YAML 更新待ち |
| `implementation_pending` | implementation / renderer / validator / MCP 更新待ち |
| `fixture_pending` | fixture / golden 更新待ち |
| `verification_pending` | 検証待ち |
| `done` | 必要な反映・検証が完了 |
| `blocked` | 依存判断・外部要因で停止中 |

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

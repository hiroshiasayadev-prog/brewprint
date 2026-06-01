# Requirement Authoring Guide

## Abstract

Requirement artifact を起票・更新するときの実践ルールを定める。Requirement は要求・不足・要望を stable ID で捕捉し、現行仕様や完了状態を所有しない。

## Migration Note

Maintainer note only. This is not an instruction to read another file, and it is not part of the public guide retrieval contract.

Extracted from:

- `docs/requirements/README.md`

## Purpose

Requirement は、brewprint project で見つかった要求・不足・要望・spec gap 候補を捕捉する artifact である。

Requirement は「何が必要か」を所有する。

## When to Create

Requirement は lazy に起票する。以下のいずれかが必要になった時点で作る。

- 追加要件・不足・要望・spec gap 候補を失わず残す必要がある。
- work item の起点が必要である。
- deferred / rejected の履歴を残す必要がある。
- bugfix / refactoring / internal design 整理について、あるべき状態との乖離を追跡する必要がある。

## ID and Layout

Requirement ID:

```text
REQ-<DOMAIN>-NNN
```

File path:

```text
docs/requirements/<domain>/REQ-<DOMAIN>-NNN-<slug>.md
```

## Format

```markdown
# REQ-<DOMAIN>-NNN: <title>

- **id**: REQ-<DOMAIN>-NNN
- **status**: captured / decision_needed / accepted / deferred / rejected
- **date**: YYYY-MM-DD
- **source_refs**:
  - <artifact ID or semantic ref>
- **work_items**:
  - WORK-<DOMAIN>-NNN

## Requirement

## Evidence

## Required Outcome

## Explicitly Excluded Scope

## Boundary
```

`id` は H1 の requirement ID と一致させる。

`source_refs` は physical path ではなく、artifact ID または semantic ref を用いる。

Required metadata は Design Records MCP validation の対象である。

- `id` / `status` / `date` は存在し、non-empty でなければならない。
- `date` は `YYYY-MM-DD` format とする。
- `source_refs` / `work_items` は list field として存在しなければならない。
- `source_refs` / `work_items` は empty list を許容する。
- `source_refs` / `work_items` の empty item は validation error とし、metadata diagnostic category は `empty_required_metadata` とする。

## Status

| status | meaning |
|---|---|
| `captured` | 要求・不足・要望として捕捉済み |
| `decision_needed` | 採用可否や設計判断が必要 |
| `accepted` | 要求として採用する |
| `deferred` | 要求としては認識するが後回し |
| `rejected` | 採用しない |

`accepted` は現行仕様への反映完了を意味しない。必要な spec / internal design / implementation 更新と検証は work item で追跡する。

## Responsibility Boundary

Requirement が所有するもの:

- 要求
- 不足
- 要望
- spec gap 候補
- stable requirement ID
- work item の起点

Requirement が所有しないもの:

- 現行仕様
- 設計判断
- 内部設計
- coverage
- 作業進捗
- 具体作業手順
- 完了状態

## Relation Boundary

- Requirement の `work_items` は `WORK-*` ID-as-ref で参照する。
- Work item は source requirement を `REQ-*` ID-as-ref で参照する。
- Workflow artifact 間の canonical relation として physical path は support しない。
- Requirement を coverage edge endpoint として扱うことは semantic trace MVP の scope 外である。

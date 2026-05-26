# Requirements

`docs/requirements/` は、brewprint project で見つかった要求・不足・要望・spec gap 候補を捕捉する artifact layer である。

Requirements は現行仕様、設計判断、内部設計、coverage、作業進捗、具体作業手順を所有しない。
現行仕様は `docs/spec/`、判断は `docs/adr/`、横断進捗は `docs/work-items/`、具体作業は `docs/tasks/` が所有する。

## When to create a requirement

requirement は lazy に起票する。以下のいずれかが必要になった時点で作る。

- 追加要件・不足・要望・spec gap 候補を失わず残す必要がある
- work item の起点が必要である
- deferred / rejected の履歴を残す必要がある
- bugfix / refactoring / internal design 整理について、あるべき状態との乖離を追跡する必要がある

## ID and layout

Requirement ID は domain-scoped sequence とし、ADR 番号と結合しない。

```text
REQ-<DOMAIN>-NNN
```

初期 layout は domain 単位とする。

```text
docs/requirements/<domain>/REQ-<DOMAIN>-NNN-<slug>.md
```

例:

```text
docs/requirements/mcp/REQ-MCP-001-design-records-semantic-trace-support.md
```

## Minimal metadata

初期運用では Markdown 冒頭の bullet metadata を使う。
Design Records MCP の parser / public record schema / validation diagnostic は `docs/spec/design-records-mcp/` が所有し、この README は authoring guidance を所有する。

```markdown
# REQ-<DOMAIN>-NNN: <title>

- **id**: REQ-<DOMAIN>-NNN
- **status**: captured / decision_needed / accepted / deferred / rejected
- **date**: YYYY-MM-DD
- **source_refs**:
  - <artifact ID or semantic ref>
- **work_items**:
  - WORK-<DOMAIN>-NNN
```

`id` は H1 の requirement ID と一致させる。
`source_refs` は physical path ではなく、artifact ID または semantic ref を用いる。

## Status

| status | meaning |
|---|---|
| `captured` | 要求・不足・要望として捕捉済み |
| `decision_needed` | 採用可否や設計判断が必要 |
| `accepted` | 要求として採用する |
| `deferred` | 要求としては認識するが後回し |
| `rejected` | 採用しない |

`accepted` は現行仕様への反映完了を意味しない。必要な spec / internal design / coverage / implementation 更新と検証は work item で追跡する。

## Workflow relation boundary

- Requirement の `work_items` は、対応する work item を `WORK-*` ID-as-ref で参照する。
- Work item は source requirement を `REQ-*` ID-as-ref で参照し、配下の短期 task を `TASK-*` ID-as-ref で参照する。
- Requirement / work item / task 間の canonical relation として physical path は support しない。
- Workflow artifact から他 artifact への参照規則は既存の canonical reference 方針に従い、本 document で追加拡張しない。

## Boundary

- requirement は「何が必要か」を所有する。
- requirement は source of truth ではない。
- work item は source requirement を必ず持つ。
- work item は requirement を解消する作業フロー全体と task graph を所有し、task は短期 concrete work と個別 status / evidence を所有する。
- requirement を coverage edge endpoint として扱うことは semantic trace MVP の scope 外である。

> 由来: ADR-081, ADR-083, ADR-084, ADR-091

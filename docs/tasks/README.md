# Tasks

> Authoring guide ID: `task-authoring`
> Boundary guide ID: `artifact-boundary`

`docs/tasks/` は、work item を前進させるための短期 concrete work、完了条件、個別 status、verification evidence を記録する artifact layer である。

Task は requirement の正本、設計判断、現行仕様、work item 全体の到達点または task graph を所有しない。
Requirement は `docs/requirements/`、到達点・横断フロー・task graph は `docs/work-items/`、設計判断は `docs/adr/`、現行仕様は `docs/spec/` が所有する。従来 `milestone` と呼んでいた実行計画の役割は、新形式では work item が担う。

> 由来: ADR-091

## Granularity

新規 task は、原則として着手後 `0.5d` から `3d` 程度で完了判定できる単位に分割する。

以下のいずれかに該当する場合、単一 task とせず split を検討する。

- 3日を明らかに超える見込みである
- 独立した判断、仕様更新、実装、検証を複数含む
- outputs または done condition を一つの完了判定として説明しにくい
- 着手時に「どこまで終われば今日は閉じられるか」が明確でない

この粒度基準は厳密な見積制度ではなく、過大な作業単位により着手可能性と達成感が損なわれることを防ぐ guard である。

## Initial ID and layout convention

ADR-091 が確定するのは workflow relation が `TASK-*` ID-as-ref を用いることまでである。以下の具体的な task ID / layout shape は、新形式を運用するための初期 authoring convention とし、後続 dogfooding または MCP contract refinement で必要に応じて見直せる。

Task ID は parent work item の domain / sequence を基礎に、work item 内 sequence を付与する。

```text
TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>
```

例:

```text
TASK-MCP-002-01
TASK-MCP-002-02
```

新規 task の初期 authoring layout は domain 単位とする。

```text
docs/tasks/<domain>/TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>-<slug>.md
```

既存の `docs/tasks/m*.md` は、work item 相当の計画と具体 task が混在していた legacy milestone-shaped work record である。ADR-091 の accepted 化のみを理由に即時移動または書換えせず、archive 移行または open legacy record の `WORK-*` / `TASK-*` 分解を行う migration work で扱う。

## Minimal metadata

```markdown
# TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>: <title>

- **id**: TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>
- **status**: todo / doing / blocked / done
- **date**: YYYY-MM-DD
- **work_item**: WORK-<DOMAIN>-NNN
- **source_requirement**: REQ-<DOMAIN>-NNN
- **estimate**: <rough estimate; usually within 0.5d-3d>
- **depends_on**:
  - TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>
- **outputs**:
  - <output>
```

`id` は H1 の task ID と一致させる。
`work_item`、`source_requirement`、`depends_on` は `WORK-*` / `REQ-*` / `TASK-*` の ID-as-ref を用いる。
Workflow artifact 間の canonical relation として physical path は support しない。
`outputs` など workflow 外の artifact への参照は、既存の canonical reference 方針に従い、この guide で新たな参照種別を追加定義しない。

## Required sections

新規 task は、原則として以下の section を持つ。

- `Goal`: この task 単独で達成すること
- `Work`: 実施する具体作業
- `Done condition`: status を `done` にできる判定条件
- `Verification`: 完了判断の確認方法
- `Evidence`: 実行結果、判断結果、検証結果

Task が investigation、ADR、spec、implementation、fixture 等を作成・更新する場合、それらは outputs または evidence として記録する。Task 本文がそれら artifact の source of truth になってはならない。

## Status ownership

- 個々の task の完了状態の正本は、当該 task artifact の `status` である。
- Parent work item は requirement 解消フロー全体の処理段階を status として保持するが、配下 task の完了 checkbox または status copy を手動複製しない。
- 将来、MCP が task record を取得・集約する場合、checkbox 相当の一覧は task status から導出する projection として提供する。

## Legacy milestone-shaped record boundary

新形式では `milestone` を独立した artifact layer または relation として導入しない。

- task は短期に閉じられる具体作業を所有する。
- work item は requirement 解消のための到達点、task graph、横断進捗を所有する。
- 旧 M-series の `milestone` 表記は、移行前の歴史的ラベルとしてのみ扱う。

既存 legacy record の archive 化、open legacy record の work item / task 分解、旧 index の再構成は、別 migration work で追跡する。

## Canonical reference boundary

Workflow artifact の identity と relation は ID-as-ref で扱う。

| relation | canonical reference |
|---|---|
| task -> source requirement | `REQ-*` |
| task -> parent work item | `WORK-*` |
| task -> dependency task | `TASK-*` |

Physical path は workflow artifact 間の canonical relation として support しない。
Workflow artifact から他 artifact への参照規則は既存の canonical reference 方針に従い、本 guide では追加拡張しない。
`req:` / `work:` / `task:` の semantic prefix は導入しない。

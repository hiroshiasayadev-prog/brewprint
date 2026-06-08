# Tasks

> Directory entrypoint.
> Canonical guide ID: `task-authoring`
> Boundary guide ID: `artifact-boundary`

この README は canonical authoring guidance ではない。

Task の format / metadata / status / granularity / dependency / evidence / boundary rule は Design Records MCP 経由で参照する。

- `list_authoring_guides` で利用可能な guide ID を確認する。
- `get_authoring_guidance` に guide ID `task-authoring` を指定する。
- `get_authoring_guidance` に guide ID `artifact-boundary` を指定する。

Guide source path は public contract ではない。Docs / task / prompt / review では guide ID を参照する。

## Directory Role

`docs/tasks/` は、parent work item を前進させる短期 concrete work artifact の置き場である。

Task は、具体作業、done condition、個別 status、verification evidence、outputs、dependencies を所有する。Requirement の正本、設計判断、現行仕様本文、parent work item の到達点や task graph は所有しない。

## Layout

Task file は domain ごとに配置する。

```text
docs/tasks/<domain>/TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>-<slug>.md
```

ID grammar、required metadata、required sections、status ownership、dependency rule は guide ID `task-authoring` を参照する。

## Legacy M-Series Boundary

新しい workflow model では `milestone` を artifact layer または relation として導入しない。

- task は短期 concrete work を所有する。
- work item は source requirement の到達点、task graph、横断進捗を所有する。
- 旧 M-series の `milestone` wording は historical label としてのみ扱う。

既存の `docs/tasks/m*.md` は legacy milestone-shaped work record である。新しい workflow artifact が存在することだけを理由に移動または書換えしない。Archive migration、または open legacy record の `WORK-*` / `TASK-*` 分解は別 migration work で扱う。

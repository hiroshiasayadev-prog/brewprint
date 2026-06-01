# Work Items

> Directory entrypoint.
> Canonical guide ID: `work-item-authoring`
> Boundary guide ID: `artifact-boundary`

この README は canonical authoring guidance ではない。

Work item の format / metadata / status / task flow / relation / boundary rule は Design Records MCP 経由で参照する。

- `list_authoring_guides` で利用可能な guide ID を確認する。
- `get_authoring_guidance` に guide ID `work-item-authoring` を指定する。
- `get_authoring_guidance` に guide ID `artifact-boundary` を指定する。

Guide source path は public contract ではない。Docs / task / prompt / review では guide ID を参照する。

## Directory Role

`docs/work-items/` は、source requirement を解消するための到達点、横断作業フロー、影響範囲、task graph を記録する work item artifact の置き場である。

すべての work item は source requirement を持つ。Requirement が「何が必要か」を所有し、work item はその requirement をどの artifact 更新と task flow で完了させるかを所有する。

Work item は、現行仕様本文、設計判断の長い理由、具体的な作業手順、個別 task の status 正本を所有しない。

## Layout

Work item file は domain ごとに配置する。

```text
docs/work-items/<domain>/WORK-<DOMAIN>-NNN-<slug>.md
```

ID grammar、required metadata、task-flow convention、completion rule は guide ID `work-item-authoring` を参照する。

## Legacy Milestone Boundary

`milestone` を新しい artifact layer、canonical identity、metadata field、または work item relation として導入しない。

旧 M-series record では historical label として milestone 語彙が残りうる。新規 work は requirement、work item、短期 task artifact で表現する。

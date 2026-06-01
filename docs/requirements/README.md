# Requirements

> Directory entrypoint.
> Canonical guide ID: `requirement-authoring`
> Boundary guide ID: `artifact-boundary`

この README は canonical authoring guidance ではない。

Requirement の format / metadata / status / relation / boundary rule は Design Records MCP 経由で参照する。

- `list_authoring_guides` で利用可能な guide ID を確認する。
- `get_authoring_guidance` に guide ID `requirement-authoring` を指定する。
- `get_authoring_guidance` に guide ID `artifact-boundary` を指定する。

Guide source path は public contract ではない。Docs / task / prompt / review では guide ID を参照する。

## Directory Role

`docs/requirements/` は、brewprint project で見つかった要求・不足・要望・spec gap 候補を捕捉する requirement artifact の置き場である。

Requirement は、現行仕様本文、設計判断、内部設計、coverage、作業進捗、具体作業手順を所有しない。現行仕様は spec、判断は ADR、横断進捗は work item、具体作業は task が所有する。

## Layout

Requirement file は domain ごとに配置する。

```text
docs/requirements/<domain>/REQ-<DOMAIN>-NNN-<slug>.md
```

例:

```text
docs/requirements/mcp/REQ-MCP-001-design-records-semantic-trace-support.md
```

ID grammar と required metadata は guide ID `requirement-authoring` を参照する。

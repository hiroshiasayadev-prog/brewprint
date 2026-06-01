# Investigations

> Directory entrypoint.
> Canonical guide ID: `investigation-authoring`
> Boundary guide ID: `artifact-boundary`

この README は canonical authoring guidance ではない。

Investigation の format / metadata / lifecycle / reference / recommendation / split / boundary rule は Design Records MCP 経由で参照する。

- `list_authoring_guides` で利用可能な guide ID を確認する。
- `get_authoring_guidance` に guide ID `investigation-authoring` を指定する。
- `get_authoring_guidance` に guide ID `artifact-boundary` を指定する。

Guide source path は public contract ではない。Docs / task / prompt / review では guide ID を参照する。

## Directory Role

`docs/investigations/` は、調査結果、根拠、影響範囲の仮説、未確定点、選択肢、後続 artifact 候補を保存する investigation artifact の置き場である。

Investigation は、決定、現行仕様、要求そのもの、横断進捗、完了状態、具体的な作業手順を所有しない。

## Layout

Investigation file は domain ごとに配置する。

```text
docs/investigations/<domain>/INV-<DOMAIN>-NNN-<slug>.md
```

ID grammar、required metadata、body structure、status、canonical reference、recommendation boundary は guide ID `investigation-authoring` を参照する。

## Records

Investigation 一覧の正本はこの README には持たせない。

Investigation record discovery は Design Records MCP の `list_records(kind=investigation)` を使う。

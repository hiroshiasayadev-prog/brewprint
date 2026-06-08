# Spec Authoring Guide

> Legacy compatibility entrypoint.
> Canonical guide ID: `spec-authoring`
> Boundary guide ID: `artifact-boundary`

このファイルは canonical authoring guidance ではない。

Spec の authoring rule は Design Records MCP 経由で参照する。

- `list_authoring_guides` で利用可能な guide ID を確認する。
- `get_authoring_guidance` に guide ID `spec-authoring` を指定して spec authoring rule を読む。
- `get_authoring_guidance` に guide ID `artifact-boundary` を指定して artifact responsibility boundary を読む。

Guide source path は public contract ではない。Docs / task / prompt / review では guide ID を参照する。

## Compatibility Note

このファイルは、古いリンクが `docs/spec-authoring-guide.md` を指している場合の互換入口としてだけ残す。

以前このファイルにあった長い本文は canonical guidance と重複していた。再利用すべき guidance は `spec-authoring` と `artifact-boundary` が扱う。詳細な front matter / origin-note example は、この entrypoint が二重正本として読まれないよう、ここには残さない。

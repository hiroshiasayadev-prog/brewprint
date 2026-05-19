# Investigations

`docs/investigations/` は、ADR-085 により導入された調査 artifact の置き場である。

investigation は、複雑な変更における調査結果、根拠、影響範囲、未確定点、選択肢、後続 artifact 候補を保存する。

investigation は、requirement / work item / task / ADR / spec / internal design / coverage / 別 investigation の起票・更新前に必ず必要な gate ではない。

investigation は、決定、現行仕様、要求そのもの、横断進捗、完了状態、具体的な作業手順を所有しない。

詳細 format / lifecycle は ADR-086 に従う。
investigation は `docs/investigations/<domain>/` 配下に置く。

## Current investigations

- `docs/investigations/docs/INV-DOCS-001-investigation-artifact-format-and-lifecycle.md` — investigation artifact の directory / ID / format / lifecycle / authoring boundary を調査した。

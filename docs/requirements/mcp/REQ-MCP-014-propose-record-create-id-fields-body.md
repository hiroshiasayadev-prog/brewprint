# REQ-MCP-014: propose_record_create の id/fields/body 呼び出し規約の整理

- **id**: REQ-MCP-014
- **status**: captured
- **date**: 2026-06-02
- **source_refs**:
- **work_items**:
  - WORK-MCP-014

## Requirement

`propose_record_create` の呼び出し規約に以下の問題がある。

1. `id` がトップレベルパラメータと `fields.id` の両方に必要で二重指定になっている
2. `fields` と `body` が排他であるため、「メタデータはMCPに生成させつつ本文も渡す」というユースケースが不可能
3. `body` で起票する場合、AIは resolved ID を事前に知れないため H1・メタデータブロックへのIDのハードコードが必要になり、番号ずれ時に `filename_id_mismatch` / `duplicate_id` が発生する

これらを解消し、以下の設計に統一する必要がある。

## Required Outcome

- `id` はトップレベルパラメータのみで指定する。`fields` への `id` 指定は不要とする
- `fields` と `body` を組み合わせて渡せる。`fields` はメタデータブロックを生成し、`body` はコンテンツ部分（`## Requirement` 以降）のみを担う
- H1 およびメタデータブロックは常に MCP が生成責任を持つ。`body` にこれらを含める必要はない
- `id: new` を渡せば server-side で解決されたIDが H1・メタデータブロックに反映される

## Evidence

- `fields` に `id` を含めないと `invalid_request` になることを確認
- `fields` + `body` を同時に渡すと `invalid_request` になることを確認
- `body` でのcreate時に番号がずれて `filename_id_mismatch` が発生することを確認

## Explicitly Excluded Scope

- `propose_record_update` の呼び出し規約変更はこの REQ のスコープ外

## Boundary

- `propose_record_create` のインターフェース変更のみを対象とする
- 既存の `fields`-only / `body`-only の動作を破壊しない後方互換性を考慮する

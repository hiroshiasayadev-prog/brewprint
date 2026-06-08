# V01-REQ-MCP-014: propose_record_create の id/fields/body 呼び出し規約の整理

- **id**: V01-REQ-MCP-014
- **status**: accepted
- **date**: 2026-06-02
- **source_refs**:
- **work_items**:
  - V01-WORK-MCP-014

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
Initial gap evidence before V01-WORK-MCP-014 showed the old contract was inconvenient for AI-assisted authoring:

- omitting `fields.id` caused `invalid_request` even though top-level `id` was present;
- supplying `fields + body` caused `invalid_request`, so callers could not ask MCP to generate metadata while supplying content sections;
- full-record `body` create required caller-generated H1 and metadata, which made server-side `new` ID resolution hard to use safely.

Close evidence on 2026-06-02:

- `V01-TASK-MCP-014-02` updated `SPEC-design-records-mcp-tools` and authoring guidance to document structured `fields`, section-only `body`, server-generated H1/metadata, and `new` placeholder rendering responsibility.
- `V01-TASK-MCP-014-03` implemented schema, validation, rendering, and regression tests in `internal/designrecords` and `internal/designrecordsmcp`.
- Runtime smoke in `V01-TASK-MCP-014-04` used the actual stdio JSON-RPC Design Records MCP path, `go run ./cmd/design-records-mcp --root .`, and passed the requested `propose_record_create` cases.
- Runtime smoke confirmed `propose_record_create` schema requires only `kind`, `id`, and `title`; `fields` is not schema-required; `body` and `body_cache_id` are exposed.
- Runtime smoke confirmed fields-only create succeeds without `fields.id` and generates H1/metadata from MCP-owned rendering.
- Runtime smoke confirmed `fields + body` create succeeds when `body` is section-only content, with MCP-generated H1 and metadata using the resolved ID.
- Runtime smoke confirmed `REQ-MCP-new` resolves server-side and renders generated H1/metadata with the resolved ID while caller body omits the resolved ID.
- Runtime smoke confirmed exact top-level ID plus matching `fields.id` remains compatibility-tolerated.
- Runtime smoke confirmed exact top-level ID plus mismatching `fields.id`, `new` placeholder plus `fields.id`, `body + body_cache_id`, `fields + body_cache_id`, and stale full-record body in `fields + body` mode are rejected with documented diagnostics.
- Runtime smoke confirmed body-only full-record create still works as a legacy compatibility mode, not the preferred create path.
- `go test ./internal/designrecords ./internal/designrecordsmcp` passed on 2026-06-02.
- `go test ./...` passed on 2026-06-02.
- Design Records MCP validation for `V01-REQ-MCP-014`, `V01-WORK-MCP-014`, and `V01-TASK-MCP-014-01..V01-TASK-MCP-014-04` passed with no diagnostics.

## Explicitly Excluded Scope

- `propose_record_update` の呼び出し規約変更はこの REQ のスコープ外
- V01-REQ-MCP-015 の cache / retry expansion は別 requirement として扱う

## Boundary

- `propose_record_create` のインターフェース変更のみを対象とする
- 既存の `fields`-only / `body`-only の動作を破壊しない後方互換性を考慮する

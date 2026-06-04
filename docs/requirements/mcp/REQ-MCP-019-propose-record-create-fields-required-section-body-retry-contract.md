# REQ-MCP-019: propose_record_create fields-required section body retry contract

- **id**: REQ-MCP-019
- **status**: accepted
- **date**: 2026-06-03
- **source_refs**:
  - REQ-MCP-014
  - REQ-MCP-015
  - WORK-MCP-012
  - WORK-MCP-014
- **work_items**:
  - WORK-MCP-018

## Requirement

`propose_record_create` の正規 create contract を、MCP-generated identity / metadata 前提へ整理する必要がある。

WORK-MCP-014 で `fields + body` が valid になり、`body` は H1 / metadata を含まない section-only content として扱う方針になった。一方で、legacy `body`-only / `body_cache_id`-only create は full-record body を caller が渡す互換 mode として残っており、MCP が H1 / metadata / resolved ID を生成責任として持つ設計と噛み合っていない。

REQ-MCP-015 / WORK-MCP-012 では、body を受け取った failed propose response が `body_cache` を返すことを close した。しかし、`fields + body` create の retry form である `fields + body_cache_id` を valid 化することは `propose_record_create` の input contract 変更であり、REQ-MCP-015 の範囲外として残した。

そのため、create contract の正規形として `fields` を必須化し、`body` / `body_cache_id` は section-only content source として扱う設計へ移行する必要がある。

## Evidence
- `WORK-MCP-018` completed the strict `propose_record_create` contract cleanup.
- Spec updates completed in `docs/spec/design-records-mcp/tools.md` and `docs/spec/design-records-mcp/schema.md`:
  - `fields` is required for create.
  - `fields`, `fields + body`, and `fields + body_cache_id` are valid create forms.
  - `body`-only / `body_cache_id`-only create are invalid.
  - `body + body_cache_id` and `fields + body + body_cache_id` remain invalid.
  - invalid request with submitted string `body` should preserve the body through returned `body_cache`.
- Implementation completed in `internal/designrecordsmcp/tools.go` and `internal/designrecords/authoring.go`.
- Regression tests updated in `internal/designrecordsmcp/tools_call_test.go` and `internal/designrecords/authoring_test.go`.
- `go test ./... -count=1` passed across all packages.
- Runtime smoke through actual Design Records MCP stdio JSON-RPC passed:
  - fields-only create: proposal created.
  - fields + body create: proposal created and returned body_cache.
  - fields + body_cache_id retry with returned cache id: proposal created.
  - body-only create: invalid_request and body_cache returned.
  - body_cache_id-only create: invalid_request and no body_cache.
  - fields + missing body_cache_id: body_cache_not_found.
  - body + body_cache_id: invalid_body_source and no body_cache.
  - fields + body + body_cache_id: invalid_body_source and no body_cache.

## Required Outcome

- `propose_record_create` の正規 create mode は `fields` 必須とする。
- `fields + body` は structured metadata + section-only content create として維持する。
- `fields + body_cache_id` を `fields + body` の retry form として valid にする。
- `body` / `body_cache_id` は H1 / metadata を含まない section-only content source として扱う。
- `body + body_cache_id` や `fields + body + body_cache_id` は proposal input として invalid のままにする。ただし submitted `body` が string として受け取れている場合、新しい `body_cache` を返して本文を保護してよい。
- legacy `body`-only / `body_cache_id`-only full-record create mode は deprecated とし、廃止または warning / compatibility boundary を明確化する。
- spec / guidance / tests / runtime smoke で、new code が `fields`, `fields + body`, `fields + body_cache_id` を使うことを明確にする。

## Explicitly Excluded Scope

- `propose_record_update` の body source contract 変更
- Design Records MCP authoring transaction model 全体の再設計
- generic filesystem write API
- workflow reciprocal relation model の変更

## Boundary

- REQ-MCP-014 と REQ-MCP-015 の後続 contract cleanup として扱う
- REQ-MCP-015 は failed propose body_cache return behavior の close 済み scope として再オープンしない
- WORK-MCP-014 の `fields + body` section-only create 方針を前提とする

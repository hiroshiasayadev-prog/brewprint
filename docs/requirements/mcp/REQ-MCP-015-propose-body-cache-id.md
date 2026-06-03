# REQ-MCP-015: propose 失敗時の body_cache_id 返却

- **id**: REQ-MCP-015
- **status**: accepted
- **date**: 2026-06-02
- **source_refs**:
- **work_items**:
  - WORK-MCP-012

## Requirement

`propose_record_create` / `propose_record_update` が `proposal_created: false` で失敗した場合、
`body_cache_id` が返却されない。

body_cache_id の設計意図は「失敗後の再proposeでbodyを再生成させず drift を防ぐ」ことにある（ADR-093）。
しかし実際には **失敗時にこそ** body_cache_id が必要であり、成功時のみ返却される現状は設計意図と逆になっている。

## Evidence
Initial gap evidence:

- Before REQ-MCP-015 close, body-cache behavior around failed propose responses was unclear and under-tested.
- The suspected update-side missing-cache gap did not reproduce; current implementation already returned `body_cache` for named-section no-match and ambiguous update failures.

Close evidence on 2026-06-03:

- `TASK-MCP-012-01` reproduced and classified failure behavior for `propose_record_create` / `propose_record_update`.
- `TASK-MCP-012-02` added regression coverage for request-level no-cache cases and post-body-receipt cache-return cases. `go test ./internal/designrecords ./internal/designrecordsmcp` passed, and `go test ./...` passed across all 17 packages.
- `TASK-MCP-012-03` completed runtime smoke through `go run ./cmd/design-records-mcp --root .` using JSON-RPC `tools/call`.
- Runtime smoke confirmed `propose_record_update` named-section no-match with submitted `body` returned `proposal_created:false`, `section_selector_no_match`, and `body_cache` with `body_cache_id`, `expires_at`, and `retention_days`.
- Runtime smoke confirmed `propose_record_create` `fields + body` with stale full-record body returned `proposal_created:false`, `invalid_request`, and `body_cache` with `body_cache_id`, `expires_at`, and `retention_days`.

Boundary retained for follow-up:

- `fields + body_cache_id` as a valid retry form for `fields + body` create is a `propose_record_create` input-contract change and is intentionally outside REQ-MCP-015.
- Legacy full-record create deprecation / fields-required create contract tightening is also outside REQ-MCP-015 and should be handled by a separate follow-up requirement.
## Required Outcome

- bodyを受け取った後にfailした場合（request-level invalidとは区別する）、`proposal_created: false` でも `body_cache_id` を返却する
- bodyを受け取る前に弾かれた場合（スキーマエラー等）は返却不要

## Explicitly Excluded Scope

- body を受け取る前のリクエストレベルエラーでの body_cache_id 返却

## Boundary

- ADR-093 の body_cache 設計の範囲内で対応する

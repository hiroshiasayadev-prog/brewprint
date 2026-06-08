# V01-TASK-MCP-018-05: Complete runtime smoke and close synchronization

- **id**: V01-TASK-MCP-018-05
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-018
- **source_requirement**: V01-REQ-MCP-019
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-018-04
- **outputs**:
  - Runtime smoke result for strict propose_record_create retry contract
  - Close synchronization updates for TASK / WORK / REQ

## Goal

V01-REQ-MCP-019 / V01-WORK-MCP-018 の実装後検証として、strict contract の runtime smoke と close synchronization を完了する。

## Work

- `go run ./cmd/design-records-mcp --root .` を使い、stdio JSON-RPC `tools/call` 経由で `propose_record_create` の runtime smoke を行う。
- fields-only、fields + body、fields + body_cache_id retry form を確認する。
- body-only、body_cache_id-only、body + body_cache_id、fields + body + body_cache_id、full-record body in section-only mode などの invalid path を確認する。
- invalid request with submitted body が body_cache を返すことを確認する。
- `go test ./...` を実行する。
- TASK / WORK / REQ の status、Evidence、work_items / tasks relation を同期する。

## Done condition

- Runtime smoke が actual Design Records MCP stdio path で完了している。
- `go test ./...` が通っている。
- `validate_records` が V01-REQ-MCP-019 / V01-WORK-MCP-018 / V01-TASK-MCP-018-01..05 に対して diagnostics なしで通っている。
- `V01-REQ-MCP-019` が accepted、`V01-WORK-MCP-018` が done、配下 task が done に同期されている。

## Verification

- Runtime smoke command / cases / result を Evidence に残す。
- Test command / result を Evidence に残す。
- Design Records MCP validation result を Evidence に残す。

## Evidence
- Runtime smoke completed through actual Design Records MCP stdio JSON-RPC path using Python subprocess driver.
- Command executed from repository root:
  - `python tmp.py`
- Runtime smoke covered:
  - fields-only create: proposal created.
  - fields + body create: proposal created and returned body_cache.
  - fields + body_cache_id retry using the returned cache id in the same MCP process: proposal created.
  - body-only create: invalid_request and returned body_cache for submitted body preservation.
  - body_cache_id-only create: invalid_request and no new body_cache.
  - fields + missing body_cache_id: body_cache_not_found.
  - body + body_cache_id: invalid_body_source and no body_cache.
  - fields + body + body_cache_id: invalid_body_source and no body_cache.
- Runtime smoke result: PASS.
- MCP process exit code: 0.
- Implementation verification from `V01-TASK-MCP-018-04`: `go test ./... -count=1` passed across all packages.
- Close synchronization updated `V01-TASK-MCP-018-01..05`, `V01-WORK-MCP-018`, and `V01-REQ-MCP-019`.

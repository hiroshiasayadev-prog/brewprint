# V01-TASK-MCP-018-04: Implement strict create contract and regression tests

- **id**: V01-TASK-MCP-018-04
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-018
- **source_requirement**: V01-REQ-MCP-019
- **estimate**: 1.5d
- **depends_on**:
  - V01-TASK-MCP-018-03
- **outputs**:
  - Updated MCP tool schema / validation / create rendering implementation
  - Regression tests for strict fields-required create contract

## Goal

`propose_record_create` の strict fields-required create contract と `fields + body_cache_id` retry form を実装し、regression tests で固定する。

## Work

- `fields` required化を MCP tool schema に反映する。
- `fields + body_cache_id` を `fields + body` の retry form として valid 化する。
- `body_cache_id` から復元した body を section-only content source として扱う。
- `body`-only / `body_cache_id`-only create を invalid にする。
- `body + body_cache_id` と `fields + body + body_cache_id` は invalid のまま維持する。
- submitted `body` が string として受け取れている invalid request では、新しい `body_cache` を返して本文を保護する。
- `V01-TASK-MCP-018-02` の影響調査に基づき、既存 tests / runtime smoke を strict contract に合わせて更新する。
- Regression tests を追加・更新する。

## Done condition

- fields-only、fields + body、fields + body_cache_id が valid path として通る。
- body-only、body_cache_id-only、body + body_cache_id、fields + body + body_cache_id、full-record body in section-only mode が expected diagnostic で落ちる。
- invalid request with submitted body が body_cache を返す behavior が tests で固定されている。
- `go test ./internal/designrecords ./internal/designrecordsmcp` が通る。

## Verification

- Unit / integration regression tests を実行する。
- Tool schema と implementation の allowed input combinations が spec と一致していることを確認する。

## Evidence
- Verdict: PASS.
- `go test ./... -count=1` passed across all packages.
- Files changed:
  - `internal/designrecordsmcp/tools.go`: added `fields` to `propose_record_create` schema required array.
  - `internal/designrecords/authoring.go`: removed `fields + body_cache_id` early rejection, added `fields == nil` invalid guard with body cache preservation, and removed the legacy body-only branch from `prepareCreate`.
  - `internal/designrecordsmcp/tools_call_test.go`: renamed `TestToolsProposeRecordCreateSchemaFieldsOptional` to `TestToolsProposeRecordCreateSchemaFieldsRequired` and inverted the assertion.
  - `internal/designrecords/authoring_test.go`: updated strict contract subcases and body cache classification tests.
- Behavior verified by tests:
  - `fields`, `fields + body`, and `fields + body_cache_id` valid paths.
  - missing cache for `fields + body_cache_id` returns `body_cache_not_found`.
  - `body`-only create returns `invalid_request` and preserves submitted body in `body_cache`.
  - `body_cache_id`-only create returns `invalid_request` without new body cache.
  - `body + body_cache_id` and `fields + body + body_cache_id` return `invalid_body_source` without body cache.
- New / updated coverage includes `fieldsAndCacheRetry`, `fieldsBodyCacheConflict`, body-only fields-missing cache preservation cases, and schema required assertion.
- Spec was not changed in this task; it used the completed `V01-TASK-MCP-018-03` contract.
- Runtime smoke / close synchronization remains for `V01-TASK-MCP-018-05`.

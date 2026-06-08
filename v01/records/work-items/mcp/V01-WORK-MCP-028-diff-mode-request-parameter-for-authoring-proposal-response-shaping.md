# V01-WORK-MCP-028: diff_mode request parameter for authoring proposal response shaping

- **id**: V01-WORK-MCP-028
- **status**: done
- **date**: 2026-06-07
- **source_requirement**: V01-REQ-MCP-031
- **impact_refs**:
  - docs/spec/design-records-mcp/tools.md
  - internal/designrecords/types.go
  - internal/designrecords/authoring.go
  - internal/designrecordsmcp/tools.go
  - internal/designrecords/authoring_test.go
- **tasks**:
  - V01-TASK-MCP-028-01
  - V01-TASK-MCP-028-02
  - V01-TASK-MCP-028-03

## Goal

`propose_record_create` と `propose_record_update` に `diff_mode` リクエストパラメータを追加する。
デフォルトを `summary`（unified diff text なし・ファイルメタデータのみ）とし、agent が不要な大容量 diff テキストを受け取らずに済むようにする。

`diff_mode` 値:

| 値 | 意味 |
|---|---|
| `summary` | デフォルト。`diff.text` を省略し、`diff.files` のみ返す |
| `patch` | 現行と同等。`diff.text` を含む完全な unified diff を返す |
| `none` | diff 詳細を省略。`diff: {"omitted": true}` のみ返す |

## Boundary

- Proposal の内部保持 diff・accept 動作・validation・diagnostics は `diff_mode` に依存しない。
- `get_proposed_write` は常に `patch` モード相当の完全 diff を返す（`diff_mode` 非対応）。
- 不正な `diff_mode` 値は `invalid_request` または同等のブロッキング diagnostic を返す。

## Impact Scope

- `docs/spec/design-records-mcp/tools.md` — `diff_mode` パラメータと応答形状を追記
- `internal/designrecords/types.go` — `DiffMode` 定数・`Diff.Omitted` フィールド・リクエスト構造体への `DiffMode` 追加
- `internal/designrecords/authoring.go` — バリデーション・応答シェーピング実装
- `internal/designrecordsmcp/tools.go` — 両ツール MCP スキーマへ `diff_mode` 追加
- `internal/designrecords/authoring_test.go` — 3 モードのテスト追加

## Task flow

```
V01-TASK-MCP-028-01 → [spec review gate] → V01-TASK-MCP-028-02 → V01-TASK-MCP-028-03
```

spec update 完了後、外部レビュー承認を得てから実装タスクに進む。

## Task Candidates

- V01-TASK-MCP-028-01: spec update — tools.md に `diff_mode` パラメータと各モード応答形状を追記
- V01-TASK-MCP-028-02: types + implementation — types.go / authoring.go / designrecordsmcp/tools.go
- V01-TASK-MCP-028-03: tests + verification — authoring_test.go + go test

## Completion Condition

- `propose_record_create` / `propose_record_update` が `diff_mode` を受け付け、`summary` (default) / `patch` / `none` で応答形状が変わる。
- 不正な `diff_mode` 値で `invalid_request` が返る。
- spec / tests が合格し、V01-REQ-MCP-031 の Required Outcome を全項目満たしている。

## Evidence

2026-06-07: 全タスク完了。

- V01-TASK-MCP-028-01: `docs/spec/design-records-mcp/tools.md` 更新。外部 LLM レビュー 3 回・承認済み。
- V01-TASK-MCP-028-02: `authoring.go` / `tools.go` 実装完了。`go test` 全件 PASS。
- V01-TASK-MCP-028-03: `authoring_test.go` に `TestDiffModeRequestParameter`（8 サブテスト）追加。`go test ./internal/designrecords/... ./internal/designrecordsmcp/...` 全件 PASS。

# V01-REQ-MCP-017: workflow artifact close validation for required narrative sections

- **id**: V01-REQ-MCP-017
- **status**: accepted
- **date**: 2026-06-03
- **source_refs**:
  - V01-WORK-MCP-014
- **work_items**:
  - V01-WORK-MCP-016

## Requirement

Workflow artifact が close 状態になっても、required narrative section の本文が空のまま validation を通過してしまう。

特に `WORK` が `status: done` の状態で `Goal` / `Boundary` / `Evidence` の本文が空でも検出されないと、close record としての説明責務を満たさない artifact が残る。

## Evidence
- `V01-WORK-MCP-014` は一度 `status: done` になったが、`Goal` と `Boundary` が空のまま Design Records MCP validation を通過していた。
- 現行 validation は metadata と relation の妥当性を主に確認しており、status に応じた required narrative section の non-empty 条件を検出していない。

Close evidence on 2026-06-03:

- `V01-WORK-MCP-016` completed the required narrative section validation work for this requirement.
- `V01-TASK-MCP-016-01` defined the status-gated policy matrix for `WORK done`, `TASK done`, and `REQ accepted`.
- `V01-TASK-MCP-016-02` updated `SPEC-design-records-mcp-tools` and `SPEC-design-records-mcp-schema` with diagnostic contract and policy details.
- `V01-TASK-MCP-016-03` implemented `missing_required_section` and `empty_required_section` diagnostics with regression tests.
- `V01-TASK-MCP-016-04` completed targeted tests, full tests, newline-delimited JSON-RPC runtime smoke, and close synchronization.
- Runtime smoke verified `validate_records` returned `ok:true` / `diagnostics:null` for `V01-TASK-MCP-016-01..V01-TASK-MCP-016-04`, `V01-WORK-MCP-016`, and `V01-REQ-MCP-017`.
- `go test ./internal/designrecords ./internal/designrecordsmcp` passed.
- `go test ./...` passed.

## Required Outcome

- Workflow artifact の close/status transition に必要な narrative section が空の場合、validation diagnostic を返す。
- 少なくとも `WORK status: done` では `Goal` / `Boundary` / `Evidence` の本文が non-empty であることを検証する。
- `TASK status: done` と `REQ status: accepted` に対する non-empty section policy も、既存 authoring guides と整合する形で定義する。
- Diagnostic は対象 record ID、section 名、status、severity を含み、修正しやすい形にする。

## Explicitly Excluded Scope

- narrative section の内容品質や十分性の自動判定
- Markdown 全体 formatter の導入
- status transition workflow の新規導入

## Boundary

- Design Records MCP validation / authoring guards における workflow artifact の required section non-empty 検証を対象とする。

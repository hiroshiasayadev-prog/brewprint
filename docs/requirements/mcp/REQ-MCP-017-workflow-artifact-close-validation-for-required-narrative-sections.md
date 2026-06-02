# REQ-MCP-017: workflow artifact close validation for required narrative sections

- **id**: REQ-MCP-017
- **status**: captured
- **date**: 2026-06-03
- **source_refs**:
  - WORK-MCP-014
- **work_items**:

## Requirement

Workflow artifact が close 状態になっても、required narrative section の本文が空のまま validation を通過してしまう。

特に `WORK` が `status: done` の状態で `Goal` / `Boundary` / `Evidence` の本文が空でも検出されないと、close record としての説明責務を満たさない artifact が残る。

## Evidence

- `WORK-MCP-014` は一度 `status: done` になったが、`Goal` と `Boundary` が空のまま Design Records MCP validation を通過していた。
- 現行 validation は metadata と relation の妥当性を主に確認しており、status に応じた required narrative section の non-empty 条件を検出していない。

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

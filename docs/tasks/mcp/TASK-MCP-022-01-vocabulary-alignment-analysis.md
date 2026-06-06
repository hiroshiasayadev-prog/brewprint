# TASK-MCP-022-01: vocabulary alignment analysis

- **id**: TASK-MCP-022-01
- **status**: done
- **date**: 2026-06-06
- **work_item**: WORK-MCP-022
- **source_requirement**: REQ-MCP-026
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - vocabulary comparison table
  - alignment options matrix with trade-off evaluation
  - recommendation with rationale

## Goal

現在の全 kind status vocabulary を棚卸しし、LLM recall likelihood との乖離を特定する。
alignment オプションを複数案評価し、推奨案と根拠を提示する。
これにより USER GATE ① で alignment direction を決定できる状態にする。

## Work

- 全 kind（`decision` / `spec` / `investigation` / `requirement` / `work_item` / `task`）の現在 status token を一覧化
- LLM が自然言語で最初に言うであろうトークンと現 canonical values を対比
- 主要な乖離箇所（特に `task` の `todo`/`doing` と `work_item` の compound tokens）を特定
- alignment オプション（A/B/C 以上）を semantic precision / LLM recall / migration cost / tooling impact の4軸で評価
- 推奨案と理由を提示

## Done condition

- vocabulary comparison table が完成している
- 最低3案の alignment option matrix が完成している
- 推奨案に対して根拠が明記されている
- ユーザーが GATE ① で選択できる状態になっている

## Verification

- 各 kind の現在 status values が `docs/spec/design-records-mcp/schema.md` § status テーブルと一致している
- options matrix の評価基準が REQ-MCP-026 § Required Outcome の acceptance criteria と対応している

## Evidence

現在の全 kind status token 棚卸しを完了した。

LLM recall 乖離の特定:

- `task.todo` → LLM first prior は `not_started`
- `task.doing` → LLM first prior は `in_progress`
- `work_item.*_pending` compound tokens → 読めるが `in_progress` に統合可能
- `decision`/`investigation` の `superseded`、`requirement` の `rejected`/`deferred` は意味的に必要なため変更対象外と確認

Migration cost 実測:

- 既存 task レコード: `todo` 使用 0 件、`doing` 使用 1 件（TASK-MCP-022-01 自身）
- 既存 work_item レコード: `not_started` 2件（変更なし）、`decision_pending` 3件（→ `blocked`）、`done` その他（変更なし）

ユーザー決定（USER GATE ①）:

- `task` と `work_item` 両方を `not_started` / `in_progress` / `blocked` / `done` の4トークンに統一
- スコープ: `task` と `work_item` のみ。`decision`/`investigation`/`requirement`/`spec` は現状維持
- Migration mapping: `task.todo` → `not_started`、`task.doing` → `in_progress`、`work_item.decision_pending` → `blocked`、`work_item.*_pending` → `in_progress`
- 次: ADR-094 起票（ADR-092 を refine する形）

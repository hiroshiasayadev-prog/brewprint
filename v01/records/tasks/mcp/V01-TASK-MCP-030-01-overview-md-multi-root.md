# V01-TASK-MCP-030-01: overview.md multi-root セクション改訂

- **id**: V01-TASK-MCP-030-01
- **status**: done
- **date**: 2026-06-09
- **work_item**: V01-WORK-MCP-030
- **source_requirement**: V01-REQ-MCP-033
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - drmcp/records/spec/design-records-mcp/overview.md 更新

## Goal

`drmcp/records/spec/design-records-mcp/overview.md` の multi-root スキャン関連記述を「MVP外」から実装仕様として改訂する。

## Work

- `### multi-root スキャン` セクションを書き換え：「MVP 外とする」→ multi-root 動作仕様として記述
- `## 対象 record` の単一 root 前提の注記（`records_root = v01/records` のみ）を更新
- `## Record scanning と namespace prefix` の記述を multi-root デフォルト対応に更新
- front matter `design_record.depends_on` に `V01-ADR-097` / `V01-ADR-099` を追加
- `last_updated` を更新

## Done condition

- `### multi-root スキャン` セクションが「MVP外」ではなく実装仕様として記述されている
- V01-REQ-MCP-033 の Required Outcome と矛盾しない
- レビュー承認済み

## Verification

overview.md の対象セクションを読み、V01-REQ-MCP-033 の Required Outcome と照合する。

## Evidence
overview.md の `## Record scanning と namespace prefix` セクションを multi-root 実装仕様に全面改訂。「MVP外」記述を削除し、`--root <repo>` auto-detect / `--records-root` backward compat / cross-namespace relation を記述。V01-ADR-097 / V01-ADR-099 を depends_on に追加。外部レビュー OK。

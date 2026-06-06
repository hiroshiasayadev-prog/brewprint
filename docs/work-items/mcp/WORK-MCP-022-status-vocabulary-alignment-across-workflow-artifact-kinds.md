# WORK-MCP-022: status vocabulary alignment across workflow artifact kinds

- **id**: WORK-MCP-022
- **status**: blocked
- **date**: 2026-06-06
- **source_requirement**: REQ-MCP-026
- **impact_refs**:
  - spec:design-records-mcp-schema
  - SPEC-design-records-mcp-schema
- **tasks**:
  - TASK-MCP-022-01
  - TASK-MCP-022-02
  - TASK-MCP-022-03
  - TASK-MCP-022-04
  - TASK-MCP-022-05
  - TASK-MCP-022-06

## Goal

`task` / `work_item` を中心に、workflow artifact kind 間の status vocabulary を整列する。
LLM callers が trial-and-error なしに valid status を使えるよう、recall likelihood が高いトークンを canonical values として確定し、spec・Go実装・authoring guides・既存レコードに反映する。

## Boundary

- 変更対象: `task` / `work_item` kind の status vocabulary（`requirement` は分析後に判断）
- 変更対象外: proposal/accept フロー、lifecycle phase の意味変更、バリデーションガードの除去
- REQ-MCP-023 (synonym repair) との interaction を確認し、026 完了後に 023 scope を再評価する

## Impact Scope

- `docs/spec/design-records-mcp/schema.md` — status テーブル更新
- `internal/designrecords/types.go` — RecordStatus 定数
- `internal/designrecords/validation.go` — isAllowedStatusForKind
- `docs/guides/task-authoring.md` — valid status tokens
- `docs/guides/work-item-authoring.md` — valid status tokens
- 既存 task / work_item records — migration

## Task flow
```
TASK-MCP-022-01 (analysis) ✓
  → USER GATE ①: alignment direction 決定 ✓
TASK-MCP-022-02 (ADR-094 drafting and acceptance) ✓
  → EXTERNAL REVIEW: Codex spec diff ✓
TASK-MCP-022-03 (spec schema.md status vocabulary update) ✓
  → EXTERNAL REVIEW: Codex code review ✓
TASK-MCP-022-04 (Go implementation: types.go / validation.go) ✓
TASK-MCP-022-05 (authoring guide update + tools.md examples) ✓
TASK-MCP-022-06 (existing record migration) ✓
  → USER GATE ②: REQ-MCP-023 scope reassessment ← pending
```

## Task Candidates
- TASK-MCP-022-01: vocabulary alignment analysis ✓
- TASK-MCP-022-02: ADR-094 drafting and acceptance ✓
- TASK-MCP-022-03: spec schema.md status vocabulary update ✓
- TASK-MCP-022-04: Go implementation (types.go / validation.go) ✓
- TASK-MCP-022-05: authoring guide update + tools.md examples ✓
- TASK-MCP-022-06: existing record migration ✓

## Completion Condition

- alignment 済み status vocabulary が spec・Go・guides に反映されている
- 既存レコードが新 canonical values にすべて migration 済みである
- `validate_records` で全レコード pass
- REQ-MCP-023 の残存 scope について判断が済んでいる

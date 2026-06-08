# V01-WORK-MCP-022: status vocabulary alignment across workflow artifact kinds

- **id**: V01-WORK-MCP-022
- **status**: done
- **date**: 2026-06-06
- **source_requirement**: V01-REQ-MCP-026
- **impact_refs**:
  - spec:design-records-mcp-schema
  - SPEC-design-records-mcp-schema
- **tasks**:
  - V01-TASK-MCP-022-01
  - V01-TASK-MCP-022-02
  - V01-TASK-MCP-022-03
  - V01-TASK-MCP-022-04
  - V01-TASK-MCP-022-05
  - V01-TASK-MCP-022-06

## Goal

`task` / `work_item` を中心に、workflow artifact kind 間の status vocabulary を整列する。
LLM callers が trial-and-error なしに valid status を使えるよう、recall likelihood が高いトークンを canonical values として確定し、spec・Go実装・authoring guides・既存レコードに反映する。

## Boundary

- 変更対象: `task` / `work_item` kind の status vocabulary（`requirement` は分析後に判断）
- 変更対象外: proposal/accept フロー、lifecycle phase の意味変更、バリデーションガードの除去
- V01-REQ-MCP-023 (synonym repair) との interaction を確認し、026 完了後に 023 scope を再評価する

## Impact Scope

- `docs/spec/design-records-mcp/schema.md` — status テーブル更新
- `internal/designrecords/types.go` — RecordStatus 定数
- `internal/designrecords/validation.go` — isAllowedStatusForKind
- `docs/guides/task-authoring.md` — valid status tokens
- `docs/guides/work-item-authoring.md` — valid status tokens
- 既存 task / work_item records — migration

## Task flow
```
V01-TASK-MCP-022-01 (analysis) ✓
  → USER GATE ①: alignment direction 決定 ✓
V01-TASK-MCP-022-02 (V01-ADR-094 drafting and acceptance) ✓
  → EXTERNAL REVIEW: Codex spec diff ✓
V01-TASK-MCP-022-03 (spec schema.md status vocabulary update) ✓
  → EXTERNAL REVIEW: Codex code review ✓
V01-TASK-MCP-022-04 (Go implementation: types.go / validation.go) ✓
V01-TASK-MCP-022-05 (authoring guide update + tools.md examples) ✓
V01-TASK-MCP-022-06 (existing record migration) ✓
  → USER GATE ②: V01-REQ-MCP-023 scope reassessment ← pending
```

## Task Candidates
- V01-TASK-MCP-022-01: vocabulary alignment analysis ✓
- V01-TASK-MCP-022-02: V01-ADR-094 drafting and acceptance ✓
- V01-TASK-MCP-022-03: spec schema.md status vocabulary update ✓
- V01-TASK-MCP-022-04: Go implementation (types.go / validation.go) ✓
- V01-TASK-MCP-022-05: authoring guide update + tools.md examples ✓
- V01-TASK-MCP-022-06: existing record migration ✓

## Completion Condition
- alignment 済み status vocabulary が spec・Go・guides に反映されている
- 既存レコードが新 canonical values にすべて migration 済みである
- `validate_records` で全レコード pass
- V01-REQ-MCP-023 の残存 scope について判断が済んでいる

## Evidence

- V01-TASK-MCP-022-01: vocabulary alignment analysis 完了（options matrix / migration cost 実測）
- V01-TASK-MCP-022-02: V01-ADR-094 drafting and acceptance 完了
- V01-TASK-MCP-022-03: spec schema.md status vocabulary 更新完了
- V01-TASK-MCP-022-04: Go implementation 完了（types.go / validation.go 更新、`go test ./...` pass、commit: 3181ba7）
- V01-TASK-MCP-022-05: authoring guides / tools.md examples 更新完了
- V01-TASK-MCP-022-06: 既存レコード migration 完了
- USER GATE ②: V01-REQ-MCP-023 を `accepted`（実質 superseded）でクローズ、criteria 1/5/6 を V01-REQ-MCP-028 へ移管、criterion 2 を V01-REQ-MCP-024 へ移管

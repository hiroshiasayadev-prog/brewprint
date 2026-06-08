# V01-TASK-MCP-027-01: add DiagnosticReciprocalUpdateIncluded and emit info diagnostic on include_required path

- **id**: V01-TASK-MCP-027-01
- **status**: done
- **date**: 2026-06-07
- **work_item**: V01-WORK-MCP-027
- **source_requirement**: V01-REQ-MCP-030
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - internal/designrecords/types.go
  - internal/designrecords/authoring.go

## Goal

`DiagnosticReciprocalUpdateIncluded` カテゴリを `types.go` に追加し、`authoring.go` の `requiredReciprocalUpdates` で `include_required` 時に info diagnostic を emit する。

## Work

- `types.go`: `DiagnosticReciprocalUpdateIncluded DiagnosticCategory = "reciprocal_update_included"` を追加
- `authoring.go`: `requiredReciprocalUpdates` の `include_required` パス（WorkItem と Task の両方）で `reciprocalUpdateIncludedDiagnostic` を返す
- ヘルパー関数 `reciprocalUpdateIncludedDiagnostic(recordID string, kind RecordKind, field, value string) Diagnostic` を追加

## Done condition

- `include_required` モード（デフォルト）で task/work_item create 時に `reciprocal_update_included` info diagnostic が返る
- `report_required_follow_up` モードの挙動は変わらない
- `required_follow_up_updates` は `include_required` 時に空のまま

## Verification

`go test ./internal/designrecords/... ./internal/designrecordsmcp/...`

## Evidence
2026-06-07: `DiagnosticReciprocalUpdateIncluded` を types.go に追加。`requiredReciprocalUpdates` の `include_required` パス（WorkItem・Task 両方）で `reciprocalUpdateIncludedDiagnostic` ヘルパーを呼び出す形に変更。コンパイル確認済み。

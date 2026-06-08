# V01-TASK-MCP-019-05: Run runtime smoke and close synchronization

- **id**: V01-TASK-MCP-019-05
- **status**: done
- **date**: 2026-06-05
- **work_item**: V01-WORK-MCP-019
- **source_requirement**: V01-REQ-MCP-021
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-MCP-019-04
- **outputs**:
  - Runtime smoke evidence for required heading case canonicalization
  - Close synchronization updates for the task, work item, and requirement
  - Final validation evidence before commit

## Goal

Complete runtime smoke and close synchronization for V01-REQ-MCP-021 / V01-WORK-MCP-019 after implementation.

## Work

- Run targeted runtime smoke for `propose_record_update` required-heading case-only fallback through the MCP boundary if feasible.
- Confirm `validate_records` public response exposes `section_heading_case_mismatch` with `actual_heading` when applicable.
- Confirm no broad fuzzy matching or optional/cross-kind heading canonicalization behavior was introduced.
- Run final validation for related workflow artifacts.
- Record smoke and validation evidence.
- Close this task, the parent work item, and the source requirement when evidence is complete.

## Done condition

- Runtime smoke demonstrates case-only required-heading repair through proposal flow or records why package-level tests are sufficient.
- Public validation diagnostic behavior is confirmed.
- `go test ./internal/designrecords ./internal/designrecordsmcp` is green.
- This task, the parent work item, and the source requirement are status-synchronized.
- Evidence is recorded before final commit.

## Verification

Recommended commands:

```powershell
cd C:\Users\imved\projects\brewprint
go test ./internal/designrecords -run TestRequiredSectionHeadingCaseMismatch -v
go test ./internal/designrecords -run TestProposeRecordUpdateRequiredHeadingCaseFallback -v
go test ./internal/designrecordsmcp -run TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields -v
go test ./internal/designrecords ./internal/designrecordsmcp
```

If MCP runtime smoke is executed manually, record the exact command and response summary in Evidence.

## Evidence
実施日: 2026-06-05

### 実行コマンドと結果

#### 1. `TestRequiredSectionHeadingCaseMismatch`

```
go test ./internal/designrecords -run TestRequiredSectionHeadingCaseMismatch -v
```

stdout:

```
=== RUN   TestRequiredSectionHeadingCaseMismatchDiagnostics
=== RUN   TestRequiredSectionHeadingCaseMismatchDiagnostics/gated_task_emits_strict_missing_error_plus_repair_info
=== RUN   TestRequiredSectionHeadingCaseMismatchDiagnostics/non-gated_task_does_not_emit_repair_info
--- PASS: TestRequiredSectionHeadingCaseMismatchDiagnostics (0.02s)
    --- PASS: TestRequiredSectionHeadingCaseMismatchDiagnostics/gated_task_emits_strict_missing_error_plus_repair_info (0.01s)
    --- PASS: TestRequiredSectionHeadingCaseMismatchDiagnostics/non-gated_task_does_not_emit_repair_info (0.01s)
PASS
ok  github.com/hiroshiasayadev-prog/brewprint/internal/designrecords
```

**結果: PASS**

検証済み内容:
- gated (done) タスクに `## Done Condition` がある場合、`missing_required_section` + `section_heading_case_mismatch` の両診断が発行される。
- `section_heading_case_mismatch` には `section: "Done condition"` (canonical)、`actual_heading: "Done Condition"` (actual) が含まれる。
- non-gated (todo) タスクでは `section_heading_case_mismatch` は発行されない。

stderr: なし

#### 2. `TestProposeRecordUpdateRequiredHeadingCaseFallback`

```
go test ./internal/designrecords -run TestProposeRecordUpdateRequiredHeadingCaseFallback -v
```

stdout:

```
=== RUN   TestProposeRecordUpdateRequiredHeadingCaseFallback
=== RUN   TestProposeRecordUpdateRequiredHeadingCaseFallback/case-only_task_required_heading_repair_creates_proposal_and_canonicalizes_diff
=== RUN   TestProposeRecordUpdateRequiredHeadingCaseFallback/fallback_applies_even_when_task_is_not_in_gated_status
=== RUN   TestProposeRecordUpdateRequiredHeadingCaseFallback/ambiguous_case-insensitive_required_headings_fail_closed
=== RUN   TestProposeRecordUpdateRequiredHeadingCaseFallback/non-case_mismatch_still_fails_with_no_fuzzy_matching
=== RUN   TestProposeRecordUpdateRequiredHeadingCaseFallback/optional_user-defined_headings_are_not_canonicalized
=== RUN   TestProposeRecordUpdateRequiredHeadingCaseFallback/cross-kind_required_headings_are_not_canonicalized
=== RUN   TestProposeRecordUpdateRequiredHeadingCaseFallback/exact_case-sensitive_match_remains_default
=== RUN   TestProposeRecordUpdateRequiredHeadingCaseFallback/fallback_honors_selector_level
--- PASS: TestProposeRecordUpdateRequiredHeadingCaseFallback (0.14s)
    --- PASS: TestProposeRecordUpdateRequiredHeadingCaseFallback/case-only_task_required_heading_repair_creates_proposal_and_canonicalizes_diff (0.02s)
    --- PASS: TestProposeRecordUpdateRequiredHeadingCaseFallback/fallback_applies_even_when_task_is_not_in_gated_status (0.02s)
    --- PASS: TestProposeRecordUpdateRequiredHeadingCaseFallback/ambiguous_case-insensitive_required_headings_fail_closed (0.02s)
    --- PASS: TestProposeRecordUpdateRequiredHeadingCaseFallback/non-case_mismatch_still_fails_with_no_fuzzy_matching (0.02s)
    --- PASS: TestProposeRecordUpdateRequiredHeadingCaseFallback/optional_user-defined_headings_are_not_canonicalized (0.02s)
    --- PASS: TestProposeRecordUpdateRequiredHeadingCaseFallback/cross-kind_required_headings_are_not_canonicalized (0.02s)
    --- PASS: TestProposeRecordUpdateRequiredHeadingCaseFallback/exact_case-sensitive_match_remains_default (0.02s)
    --- PASS: TestProposeRecordUpdateRequiredHeadingCaseFallback/fallback_honors_selector_level (0.02s)
PASS
ok  github.com/hiroshiasayadev-prog/brewprint/internal/designrecords
```

**結果: PASS (8 subtests)**

検証済み内容:
- selector `{heading: "Done condition", level: 2, match: "exact"}` でケース違いの `## Done Condition` にマッチし、proposal が作成される。
- diff に `+## Done condition` が含まれ、`+## Done Condition` は含まれない (canonical への書き換えが確認)。
- 曖昧な case-insensitive 一致 (同ファイルに `## Done Condition` と `## DONE CONDITION` が共存) は `section_selector_ambiguous` で fail closed。
- 非ケース違い (`## Done conditions`) は fuzzy matching なしで `section_selector_no_match`。
- optional / ユーザー定義 heading は canonicalization 対象外 (`section_selector_no_match`)。
- cross-kind required heading は canonicalization 対象外 (`section_selector_no_match`)。
- exact case-sensitive 一致が存在する場合は fallback より優先される。
- `level` 制約がある selector は level 違いでマッチしない (`section_selector_no_match`)。

stderr: なし

#### 3. `TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields`

```
go test ./internal/designrecordsmcp -run TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields -v
```

stdout:

```
=== RUN   TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields
--- PASS: TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields (0.01s)
PASS
ok  github.com/hiroshiasayadev-prog/brewprint/internal/designrecordsmcp
```

**結果: PASS**

検証済み内容 (MCP boundary — JSON-RPC `tools/call` レベル):
- raw JSON-RPC `{"jsonrpc":"2.0","id":921,"method":"tools/call","params":{"name":"validate_records","arguments":{"kind":"task"}}}` を直接送信し、JSON レスポンスをパースして全フィールドを検証。
- `diagnostics` 配列に以下フィールドを含む診断オブジェクトが存在する:
  - `category: "section_heading_case_mismatch"`
  - `severity: "info"`
  - `section: "Done condition"`
  - `actual_heading: "Done Condition"`
  - `status: "done"`
  - `candidate_headings` フィールドが存在する
- `ok: false` であることを確認 (canonical required section は依然として missing)。

stderr: なし

#### 4. フルスイート

```
go test ./internal/designrecords ./internal/designrecordsmcp
```

stdout:

```
ok  github.com/hiroshiasayadev-prog/brewprint/internal/designrecords      2.125s
ok  github.com/hiroshiasayadev-prog/brewprint/internal/designrecordsmcp   (cached)
```

**結果: PASS (両パッケージ)**

stderr: なし

---

### パッケージレベル MCP boundary テストで十分な理由

stdio プロセス経由のランタイムスモークを実施しなかった理由:

1. `TestToolsCallValidateRecordsExposesSectionHeadingCaseMismatchFields` は `NewServer(cfg)` + raw JSON-RPC を直接送信し、JSON レスポンスをパースして全フィールドを検証している。これは MCP boundary (tools_call 層) を完全にカバーする。
2. V01-REQ-MCP-021 の実装は designrecords (validation / authoring) 層と tools_call 層にある。stdio transport 層には変更なし。
3. stdio transport は thin wrapper であり、`handleLine` を通じて同じ `tools_call` 実装を呼び出す。パッケージレベルのテストが transport を変えずに同じコードパスを検証している。
4. 追加の stdio プロセス起動は redundant であり、テスト実行コストに見合う新たな観察値を提供しない。

**結論: パッケージレベル MCP boundary テストは Done condition の evidence として十分。**

# V01-TASK-MCP-015-01: Inventory named section replacement spacing behavior

- **id**: V01-TASK-MCP-015-01
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-015
- **source_requirement**: V01-REQ-MCP-016
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Current behavior inventory for propose_record_update named_section_replace spacing
  - Reproduction evidence for spacing loss and trailing-newline variants

## Goal

Identify the current `propose_record_update` `named_section_replace` formatting behavior and reproduce the spacing loss described by `V01-REQ-MCP-016`.

## Work

- Inspect the current named section replacement / splice implementation.
- Reproduce the case where a replaced section body is joined directly to the next same-level heading.
- Check at least these replacement body shapes:
  - no trailing newline
  - one trailing newline
  - already separated with a blank line
- Classify whether the issue is caller body normalization, section splice / join logic, or later rendering.

## Done condition

- The responsible implementation path is identified.
- The spacing behavior is described with concrete before / after examples or test observations.
- The required minimal fix point is clear enough for implementation work.

## Verification

- Use targeted unit tests, local reproduction, or runtime JSON-RPC smoke input to confirm the observed behavior.
- Record the commands or manual reproduction inputs used.

## Evidence
- **調査対象**: `internal/designrecords/authoring.go` `replaceNamedSection` 関数。
- **根本原因**:
  1. `strings.TrimSuffix(body, "\n")` は末尾 `\n` を 1 個しか除去しない。
  2. `strings.Join(out, "\n")` が各要素間に `\n` 1 個だけ挿入するため、body が `"text"` または `"text\n"` で終わる場合は blank line が生まれない。
  3. `match.EndLine` は次 heading 行番号そのものを指し、元の blank line は replacement 範囲内で捨てられる。
- **再現条件**: body が末尾 `\n` 1 個または末尾 `\n` なしで終わり、次の同レベル heading が存在する場合。
- **偶然通る条件**: body が末尾 `\n\n` で終わる場合。`TrimSuffix` が 1 個だけ除去するため、残りの `\n` が separator として機能する。
- **最小修正箇所**: `replaceNamedSection` の `strings.TrimSuffix` を trailing newline 全体の正規化に変え、次 heading が存在する場合だけ splice 時に blank line 要素を追加する。
- **失敗 regression test**: `TestReplaceNamedSectionSpacingPreservation` を `internal/designrecords/authoring_test.go` に追加済み。`no_trailing_newline` と `one_trailing_newline` は期待通り fail、`already_separated` は現状でも pass。
- **末尾 section regression test**: `TestReplaceNamedSectionSpacingLastSection` を追加済み。末尾 section 置換で余分な空行が生じないことを確認済み。
- **検証コマンド**: `go test ./internal/designrecords -run TestReplaceNamedSectionSpacing -v`。
- **変更ファイル**: `internal/designrecords/authoring_test.go` の failing regression test 追加のみ。実装ファイルへの変更は `V01-TASK-MCP-015-03` に委ねる。

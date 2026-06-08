# V01-TASK-MCP-015-03: Implement named section replacement spacing normalization and tests

- **id**: V01-TASK-MCP-015-03
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-015
- **source_requirement**: V01-REQ-MCP-016
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-MCP-015-02
- **outputs**:
  - Updated named_section_replace formatting implementation
  - Regression tests for replacement body trailing newline variants

## Goal

Implement the minimal `named_section_replace` spacing normalization and add regression tests that protect the behavior required by `V01-REQ-MCP-016`.

## Work

- Patch the responsible section replacement / splice logic so replaced section content and the next heading are separated by the canonical blank line.
- Keep the change local to `propose_record_update` `named_section_replace` formatting.
- Add regression coverage for replacement bodies with no trailing newline, one trailing newline, and an already-present blank-line separator.
- Ensure metadata block replacement and selector matching behavior remain unchanged.

## Done condition

- The implementation produces stable output independent of replacement body trailing newline shape.
- Regression tests fail before the fix and pass after the fix, or otherwise explicitly prove the corrected behavior.
- No unrelated authoring contract or Markdown whole-document formatting behavior is introduced.

## Verification

- Run targeted tests for `internal/designrecords` and/or `internal/designrecordsmcp` covering the changed behavior.
- Run broader relevant package tests if the touched code is shared by multiple authoring paths.

## Evidence

### 実装修正

**ファイル**: `internal/designrecords/authoring.go`、`replaceNamedSection` 関数。

**変更内容** (diff 概要):
```
-	replacementBody := strings.TrimSuffix(body, "\n")
+	replacementBody := strings.TrimRight(body, "\n")
 	replacement := []string{strings.Repeat("#", match.Heading.Level) + " " + match.Heading.Text}
 	if replacementBody != "" {
 		replacement = append(replacement, replacementBody)
 	}
 	out := append([]string{}, lines[:match.StartLine]...)
 	out = append(out, replacement...)
+	if match.EndLine < len(lines) {
+		out = append(out, "")
+	}
 	out = append(out, lines[match.EndLine:]...)
```

**変更理由**:
1. `strings.TrimSuffix(body, "\n")` は末尾 `\n` を 1 個しか除去しないため、`body = "text\n\n"` の場合は `"text\n"` が残り、`strings.Join` 時に偶然 blank line が生じていた。`strings.TrimRight(body, "\n")` に変えることで全末尾 `\n` を除去し、どの trailing newline 形状でも同じ `replacementBody` になる。
2. `match.EndLine` は次 heading 行番号そのものを指す。元の blank line は `lines[StartLine:EndLine]` の範囲内にあり、splice で捨てられる。次 heading が存在する場合（`match.EndLine < len(lines)`）だけ空要素 `""` を追加することで、`strings.Join("\n")` が canonical blank line (`\n\n`) を生成する。
3. 末尾 section（`match.EndLine == len(lines)`）では空要素を追加しないため、余分な trailing blank line は生じない。

### テスト結果

```
go test ./internal/designrecords -run TestReplaceNamedSectionSpacing -v

=== RUN   TestReplaceNamedSectionSpacingPreservation/no_trailing_newline  PASS
=== RUN   TestReplaceNamedSectionSpacingPreservation/one_trailing_newline PASS
=== RUN   TestReplaceNamedSectionSpacingPreservation/already_separated    PASS
=== RUN   TestReplaceNamedSectionSpacingLastSection/no_trailing_newline   PASS
=== RUN   TestReplaceNamedSectionSpacingLastSection/one_trailing_newline  PASS
ok  github.com/hiroshiasayadev-prog/brewprint/internal/designrecords
```

```
go test ./internal/designrecords ./internal/designrecordsmcp

ok  github.com/hiroshiasayadev-prog/brewprint/internal/designrecords
ok  github.com/hiroshiasayadev-prog/brewprint/internal/designrecordsmcp
```

### 変更ファイル一覧

- `internal/designrecords/authoring.go` — `replaceNamedSection` の TrimSuffix→TrimRight 変更 + blank line 挿入条件追加
- `docs/tasks/mcp/TASK-MCP-015-02-define-minimal-spacing-contract-for-named-section-replacement.md` — Evidence 追記・status → done
- `docs/tasks/mcp/TASK-MCP-015-03-implement-named-section-replacement-spacing-normalization-and-tests.md` — Evidence 追記・status → done

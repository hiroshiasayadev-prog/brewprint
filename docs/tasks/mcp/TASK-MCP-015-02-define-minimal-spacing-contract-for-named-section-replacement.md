# TASK-MCP-015-02: Define minimal spacing contract for named section replacement

- **id**: TASK-MCP-015-02
- **status**: done
- **date**: 2026-06-03
- **work_item**: WORK-MCP-015
- **source_requirement**: REQ-MCP-016
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-015-01
- **outputs**:
  - Minimal canonical separator rule for named_section_replace
  - Spec or guidance update decision

## Goal

Define the minimal formatting contract for `propose_record_update` `named_section_replace` so implementation does not broaden into a Markdown formatter.

## Work

- Decide the canonical separator rule between replaced section content and the next heading.
- Confirm the rule is limited to section replacement output formatting.
- Decide whether `SPEC-design-records-mcp-tools` or authoring guidance needs an explicit note.
- If documentation is needed, identify the exact target section and patch scope.

## Done condition

- The intended separator rule is stated clearly enough for tests to encode.
- The decision preserves existing metadata block replacement and section selector semantics.
- Any required spec / guidance update target is identified, or explicitly classified as unnecessary.

## Verification

- Compare the decision against `REQ-MCP-016` Required Outcome and Explicitly Excluded Scope.
- Confirm that no whole-document formatting behavior is introduced by the proposed rule.

## Evidence

### Spacing contract

**Rule**: `named_section_replace` は置換後の section body と次の同レベル見出しの間に、canonical blank line（空行 1 個）を保証する。

**適用条件**:
- 次の同レベル heading が存在する場合のみ blank line を挿入する。
- 末尾 section（次 heading なし）には余分な blank line を追加しない。
- caller が供給する replacement body の末尾 `\n` 数に依存しない（0 個・1 個・複数個すべて同一出力）。

**スコープ境界**:
- `named_section_replace` の section replacement formatting に限定する。
- `metadata_block_replace` の意味論は変えない。
- section selector matching の意味論は変えない。
- Markdown whole-document formatter には広げない。

### SPEC / guidance 更新要否の判断

**SPEC 更新不要**。理由:
- `SPEC-design-records-mcp-tools` の public contract は `propose_record_update` の入出力型・エラーコード・フィールド定義を所有する。
- section 置換後の内部 formatting（blank line 保証）は caller-facing の動作変更ではなく、authoring output の品質修正である。
- REQ-MCP-016 Required Outcome は「blank line を維持する」であり、spec に公式の separator rule section を追加しなくても、実装とテストで契約を表現できる。
- Markdown formatting の「常識的な読みやすさ」は spec の記述対象よりも実装 invariant に属する。
- 将来の SPEC 更新が必要になる条件: caller が separator の有無をコントロールするフラグを公開する場合、または複数の separator mode をサポートする場合。今回はいずれも対象外。

**authoring guidance 更新不要**。理由:
- `named_section_replace` の body 供給方法（末尾 `\n` の有無）について、caller への注意点を guidance に追記する必要はない。修正後は body の末尾 `\n` 数によらず安定した出力を保証するため、caller 側に制約を課さない。

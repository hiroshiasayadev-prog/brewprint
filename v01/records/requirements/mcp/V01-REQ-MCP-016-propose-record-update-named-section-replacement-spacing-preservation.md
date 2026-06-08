# V01-REQ-MCP-016: propose_record_update named section replacement spacing preservation

- **id**: V01-REQ-MCP-016
- **status**: accepted
- **date**: 2026-06-03
- **source_refs**:
  - V01-WORK-MCP-014
- **work_items**:
  - V01-WORK-MCP-015

## Requirement

`propose_record_update` の `named_section_replace` が section 本文を置換するとき、次の同レベル見出しとの間の空行を詰めてしまう。

Markdown としては解析可能でも、authoring output としては読みづらく、手動整形や後続差分のノイズを誘発する。

## Evidence

- `V01-WORK-MCP-014` の `Goal` / `Boundary` を MCP の `named_section_replace` で更新したところ、section 本文末尾と次見出しの間の空行が消えた。
- 末尾改行を含めて再 proposal しても、diff は同じ形になったため、呼び出し側の body ではなく replacement/join logic 側の整形仕様または実装に起因する可能性が高い。
- Close evidence from `V01-WORK-MCP-015`:
  - `replaceNamedSection` now trims all trailing replacement-body newlines and inserts the canonical blank line only when a next heading exists.
  - Regression tests cover no trailing newline, one trailing newline, already-separated replacement body, and last-section behavior.
  - `go test ./internal/designrecords -run TestReplaceNamedSectionSpacing -v`: pass.
  - `go test ./internal/designrecords ./internal/designrecordsmcp`: pass.
  - `go test ./...`: pass.
  - MCP runtime smoke through `go run ./cmd/design-records-mcp --root .` confirmed `propose_record_update` `named_section_replace` proposal diffs retain a blank line between replacement body and the next heading for all required trailing-newline variants.
  - Runtime smoke proposals were discarded; no smoke proposal was accepted or written to repository files.

## Required Outcome

- `named_section_replace` は置換後 section と次見出しの間に Markdown として読みやすい空行を維持する。
- 呼び出し body の末尾改行有無に依存せず、安定した canonical formatting を生成する。
- 既存の metadata block replacement や section selector の意味論は変えない。

## Explicitly Excluded Scope

- `propose_record_create` の呼び出し規約変更
- Markdown 全体 formatter の導入
- unrelated design record content rewrite

## Boundary

- `propose_record_update` の `named_section_replace` における section replacement formatting のみを対象とする。

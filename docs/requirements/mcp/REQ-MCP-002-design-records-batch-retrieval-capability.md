# REQ-MCP-002: Design records の複数取得 capability が必要

- **id**: REQ-MCP-002
- **status**: captured
- **date**: 2026-05-25
- **source_refs**:
  - ADR-077
  - ADR-087
- **work_items**:

## 要求

LLM が関連する複数の design record を確認するとき、record ごとの反復取得だけに依存せず、必要な metadata / headings / 必要に応じた本文を効率よく取得できる query capability が必要である。

## 発見根拠

M19 の contract refinement と review では、関連する ADR / spec / investigation record の確認のために `get_record` 相当の単一取得を繰り返す利用パターンが観測された。

これは現行 `get_record` の責務違反ではないが、LLM-first な調査・レビュー経路では read round-trip と文脈回収の負荷を増やす。

## Candidate solution

候補として、既存の単一取得 contract を壊さない追加 query tool `get_records` を検討する。

想定する検討事項:

- ID 配列による複数取得
- `include_body` の扱いと本文を含む場合の response size / 件数上限
- 一部 ID が存在しない場合の partial result / diagnostic contract
- record ordering の安定性

## Boundary

- 本 requirement は複数取得 capability の必要性を捕捉するものであり、tool 名や request / response schema を確定しない。
- M19 の必須完了条件には含めない。
- requirement / work item / task artifact の MCP integration は別 requirement で扱う。

## Next decision

M19 implementation / review の運用結果を追加 evidence とし、batch retrieval tool を public contract として採用するか判断する。

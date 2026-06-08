# V01-REQ-MCP-002: Design records の複数取得 capability が必要

- **id**: V01-REQ-MCP-002
- **status**: accepted
- **date**: 2026-05-25
- **source_refs**:
  - V01-ADR-077
  - V01-ADR-087
  - V01-ADR-090
- **work_items**:
  - V01-WORK-MCP-002

## 要求

LLM が関連する複数の design record を確認するとき、record ごとの反復取得だけに依存せず、必要な metadata / headings / 必要に応じた本文を効率よく取得できる query capability が必要である。

## 発見根拠

M19 の contract refinement と review では、関連する ADR / spec / investigation record の確認のために `get_record` 相当の単一取得を繰り返す利用パターンが観測された。

これは現行 `get_record` の責務違反ではないが、LLM-first な調査・レビュー経路では read round-trip と文脈回収の負荷を増やす。

## Accepted solution

`V01-ADR-090` により、既存の単一取得 contract を壊さない追加 public read-only tool `get_records` を採用する。

判断済みの境界:

- 明示された ID 配列による複数 detail 取得を担い、filter / range query は `list_records` に維持する。
- `include_body` は request 全体に適用し、raw body は truncate しない。
- 一部 ID が存在しない場合は item-level partial result を返す。
- duplicate requested ID は first occurrence のみ返し、`info` diagnostic で可視化する。
- response size の数値上限は public contract として定義しない。

## Boundary

- 本 requirement は複数取得 capability の必要性を捕捉するものであり、tool 名や request / response schema を確定しない。
- M19 の必須完了条件には含めない。
- requirement / work item / task artifact の MCP integration は別 requirement で扱う。

## Next step

`V01-WORK-MCP-002` において、`V01-ADR-090` に基づく public contract / spec 更新、implementation、tests、runtime verification を追跡する。

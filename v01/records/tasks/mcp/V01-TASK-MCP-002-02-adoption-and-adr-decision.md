# V01-TASK-MCP-002-02: Batch retrieval の採用判断と ADR 要否を確定する

- **id**: V01-TASK-MCP-002-02
- **status**: done
- **date**: 2026-05-26
- **work_item**: V01-WORK-MCP-002
- **source_requirement**: V01-REQ-MCP-002
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-MCP-002-01
- **outputs**:
  - capability の採用・延期・不採用判断
  - ADR 起票要否の判断、または必要な ADR

## Goal

V01-TASK-MCP-002-01 の evidence を基に、batch retrieval capability を public contract として進めるかを判断し、設計判断を ADR に残す必要があるかを確定する。

## Work

- capability の利用価値、追加 contract の責務境界、response size / partial result / ordering の論点を確認する。
- 既存 V01-ADR-077 / V01-ADR-087 の範囲内の補足で済むか、新たな accepted decision が必要か判断する。
- 採用する場合は V01-TASK-MCP-002-03 が更新すべき spec 範囲を指定する。
- 延期または不採用の場合は、V01-REQ-MCP-002 と V01-WORK-MCP-002 に反映すべき理由を整理する。

## Done condition

- 採用・延期・不採用のいずれかが根拠付きで確定している。
- ADR が必要なら起票・判断の成果物が存在し、不要なら不要理由が記録されている。
- 採用時は V01-TASK-MCP-002-03 が開始可能である。

## Verification

- 未判断の tool schema を暗黙に確定していないことを確認する。
- 採用判断が V01-REQ-MCP-002 の Boundary と矛盾しないことを確認する。

## Evidence

### 採用判断

- Batch retrieval capability は採用する。
- 新しい public read-only tool 名は `get_records` とし、既存 `get_record` の batch counterpart として追加する判断とする。
- 利用価値は V01-TASK-MCP-002-01 で確認済みであり、複数の design record の headings / body を確認する場面で単一 `get_record` の反復取得が発生することを根拠とする。

### Contract の責務境界

- `get_records` は、呼び出し側が明示した複数 record ID の detail retrieval のみを担う。
- candidate discovery、filter、`kind` / `status` / `id_range` / `limit` による列挙は既存 `list_records` の責務に維持し、`get_records` には持たせない。
- canonical reference の解決は `resolve_reference`、integrity validation は `validate_records` の責務に維持する。
- 対象 record kind は現行 Design Records MCP が index / get 対象として公開している `decision` / `spec` / `investigation` とする。`requirement` / `work item` / `task` の MCP support は V01-REQ-MCP-003 の対象であり、本 tool 判断に含めない。

### Request / response behavior の判断

- Request は `ids` と request 全体に一律適用する `include_body` を持つ。`include_body` の default は既存 `get_record` と揃えて `false` とする。
- Response は dedupe 後の first occurrence order を維持する。
- 一部 ID が存在しない場合、batch 全体を失敗させず item-level の `retrieval_status: "not_found"` と `record_not_found` diagnostic を返す partial result contract とする。
- 同一 ID が複数回指定された場合、最初の一件だけを取得結果として返し、重複分は top-level の `duplicate_requested_id_ignored` / `info` diagnostic として可視化する。duplicate は request error としない。
- Item-level の取得結果 field は `retrieval_status` とし、取得した record 自体の lifecycle `status` と区別する。
- `include_body: true` の body は既存 `get_record` と同様に raw 完全本文を返す。本文の truncate は行わない。
- Response total length / body size の数値上限は public contract として定義しない。サイズ都合を Design Records MCP の意味論に持ち込まず、呼び出し側は必要に応じて取得 ID 群を分割できるためである。

### Representative example boundary

- Contract example には `V01-ADR-077`（decision）、`SPEC-design-records-mcp-tools`（spec）、`V01-INV-DOCS-001`（investigation）を混在させ、現行 indexed record kind を横断取得できることを示す。
- Example には duplicate `V01-ADR-077` と未存在の `V01-INV-DOCS-999` を含め、duplicate info diagnostic と partial result behavior を示す。
- `REQ-*` / `WORK-*` / `TASK-*` は現行 MCP 取得対象外であるため example に含めない。

### ADR 要否

- 新 ADR が必要である。
- 理由は、新 public tool の追加、`get_record` の単一取得 error behavior と異なる partial result、duplicate input の info diagnostic 化、`list_records` と `get_records` の責務境界を新たな design decision として確定するためである。
- `V01-ADR-090: Design Records MCP batch retrieval tool boundary` を `accepted` として起票し、Design Records MCP から record として取得できることを確認した。
- `V01-ADR-090` により、本 task の ADR 成果物条件を満たした。`V01-TASK-MCP-002-03` は `get_records` の public contract / spec 更新へ着手可能である。

# V01-TASK-MCP-002-01: Batch retrieval の利用 evidence と調査要否を整理する

- **id**: V01-TASK-MCP-002-01
- **status**: done
- **date**: 2026-05-26
- **work_item**: V01-WORK-MCP-002
- **source_requirement**: V01-REQ-MCP-002
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - batch retrieval capability の必要性を判断するための evidence summary
  - investigation artifact 起票要否の判断

## Goal

Design Records MCP の利用で、複数 record の取得負荷がどの場面で発生しているかを確認し、V01-REQ-MCP-002 を設計判断へ進めるための根拠を短く整理する。

## Work

- M19 review / close で観測された単一 `get_record` の反復利用パターンを確認する。
- 現行 `list_records` / `get_record` contract で不足する読み取り経路を整理する。
- 独立した investigation artifact に保存すべき複雑な比較・不確実性があるか判断する。
- 次 task で判断すべき最小 capability 候補と論点を列挙する。

## Done condition

- 採用判断に必要な evidence と未確定点が整理されている。
- investigation artifact を起票するか、不要とする理由が記録されている。
- V01-TASK-MCP-002-02 が開始可能になっている。

## Verification

- V01-REQ-MCP-002 の発見根拠と矛盾しないことを確認する。
- 既存 public tool contract を未確認のまま前提にしていないことを確認する。

## Evidence

### 観測した根拠

- `docs/tasks/m19-design-records-semantic-trace-support.md` の Close note は、M19 の follow-up requirement として `V01-REQ-MCP-002`（batch retrieval capability）を captured 済みと記録している。
- `SPEC-design-records-mcp-tools` は、`list_records` を本文読取り前の候補絞り込み、`get_record` を単一の `id` から metadata / headings / 必要に応じた body を取得する tool として定義している。現行 tool set に複数 ID をまとめて詳細取得する contract はない。
- 本 dogfooding（2026-05-26）では、関連する設計根拠と現行 contract の確認にあたり、`list_records` で accepted ADR を把握した後、`get_record` を `V01-ADR-077`、`V01-ADR-087`、`SPEC-design-records-mcp-tools`、`V01-ADR-085` に対して個別に実行した。関連 record の headings / body を複数確認する経路では、単一取得の反復が実際に発生した。
- `V01-ADR-077` は、`get_record` を候補絞り込み後の raw body / path 取得のための P0 tool として採用している。`V01-ADR-087` は、ADR / spec / investigation を横断する record-oriented query を維持するため、確認対象 record の種類が増える方向を決定している。

### 現行 contract で不足する読み取り経路

- metadata による候補絞り込みは `list_records` で行えるため、全 record 本文の一括返却は必要性が確認されていない。
- 一方、絞り込み後に複数 record の headings または raw body を比較・確認する場合、現行 `get_record` は `id` を一件だけ受け付けるため、record 件数分の反復 query が必要になる。
- この不足は `get_record` の既存責務違反ではなく、複数根拠を同時に確認する LLM-first な review / investigation 経路での read round-trip 負荷として扱う。

### Investigation artifact 起票要否

- 独立した investigation artifact は起票しない。
- 理由は、必要性を裏付ける観測が `V01-REQ-MCP-002`、M19 close note、現行 tool spec、および関連 accepted ADR に限定され、複雑な比較調査や追加の影響範囲仮説を保存する必要がないためである。
- `V01-ADR-085` と `docs/investigations/README.md` は investigation を必須 gate とせず、既存 requirement / work item / task だけでは根拠・影響範囲・判断候補を安全に保存できない場合に限って使用する境界を定めている。今回の evidence と未確定点は本 task 内で保持できる。

### V01-TASK-MCP-002-02 に渡す未確定論点

- batch retrieval capability を public contract に追加するか、現行の単一取得を維持するか。
- 採用する場合に、ID 配列、`include_body` の適用単位、本文を含む際の response size / 件数上限、一部 ID 不在時の partial result / diagnostic、返却順序をどう定義するか。
- 上記は本 task では確定せず、採用判断と ADR 要否を扱う `V01-TASK-MCP-002-02` の判断対象とする。

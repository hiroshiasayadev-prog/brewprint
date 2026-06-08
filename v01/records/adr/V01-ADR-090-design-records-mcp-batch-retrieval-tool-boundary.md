# V01-ADR-090: Design Records MCP batch retrieval tool boundary

- **status**: accepted
- **date**: 2026-05-26
- **depends_on**: V01-ADR-077, V01-ADR-087
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

V01-ADR-077 は Design Records MCP の read-only query 境界として `list_records` / `get_record` / `validate_records` を定め、`get_record` を record ID から metadata、headings、必要に応じて raw body を取得する単一取得 tool として採用した。

V01-ADR-087 は Design Records MCP の record index を `decision` / `spec` / `investigation` の横断取得・検証へ拡張し、設計対話で複数種別の record を参照する導線を確立した。

M19 close 後の dogfooding では、関連する ADR、spec、investigation record の判断根拠を確認する際、候補を絞った後に複数 record の headings または body を確認するため、単一 `get_record` を繰り返す経路が実際に生じた。この観測は `V01-REQ-MCP-002` および `V01-TASK-MCP-002-01` に記録されている。

この負荷は既存 `get_record` の責務違反ではない。一方で、LLM-first な調査・レビューで選択済みの複数根拠をまとめて確認する read path がなく、read round-trip と文脈回収の負荷が残る。

## 決定

Design Records MCP の public read-only tool として、複数 record の detail retrieval を行う `get_records` を追加する。

本ADRは V01-ADR-077 / V01-ADR-087 の既存 tool 境界を破棄せず、選択済み record の batch retrieval 経路を追加する形で refine する。

### 1. `get_records` の責務

`get_records` は、呼び出し側が明示した複数 record ID について、既存 `get_record` と同じ品質の detail representation をまとめて取得する tool とする。

`get_records` が扱うもの:

- 明示された複数 ID の metadata / path / headings の取得
- requested scope に応じた raw body の取得
- missing requested ID の item-level 可視化
- duplicate requested ID の informational 可視化

`get_records` が扱わないもの:

- candidate discovery
- `kind` / `status` / `id_range` / `limit` による filter / listing
- canonical reference resolution
- metadata integrity validation
- body の要約・整形・正規化
- record 間比較や自動選別

candidate discovery と filter / listing は `list_records`、canonical reference resolution は `resolve_reference`、integrity validation は `validate_records` の責務に維持する。

### 2. 対象 record kind

`get_records` は、起票時点で Design Records MCP が index / get 対象として公開している以下の record kind を扱う。

- `decision`
- `spec`
- `investigation`

`requirement` / `work item` / `task` は現行 Design Records MCP の record kind ではなく、本ADRの対象に含めない。これらの MCP support を追加する場合は、別 requirement / decision において `get_records` への影響を判断する。

### 3. Request の基本形

`get_records` は、取得対象を表す `ids` と、request 全体に一律適用する `include_body` を受け取る。

- `ids` は明示された record ID の配列とする。
- `include_body` の default は既存 `get_record` と揃えて `false` とする。
- record ごとに異なる body 取得条件を指定する query plan 形式は採用しない。
- `kind` / `status` / `id_range` / `limit` は request に持たせない。

`get_records.ids[]` は、Design Records MCP の index に対する exact record ID lookup key としてのみ評価する。`get_records` は canonical reference resolution または input kind classification を行わず、`resolve_reference` の `unsupported` contract を導入しない。文字列として受理された値が indexed record ID に一致しない場合は、semantic ref、未対応 artifact ID、physical path、case や whitespace の異なる値であっても、item-level の `not_found` として扱う。

具体的な request schema と invalid request behavior は spec で定義する。

### 4. Partial result

一部の requested ID が存在しない場合、`get_records` は batch 全体を失敗させず、取得できた record を返す。

存在しない requested ID は item-level result として返し、取得状態を `retrieval_status: "not_found"`、診断を `record_not_found` として表現する。

取得に成功した item は `retrieval_status: "found"` とし、取得した record representation を返す。

`retrieval_status` は record 自体の lifecycle を表す `status` と異なる field とする。例えば、取得した ADR の lifecycle `status: "accepted"` と、取得処理結果 `retrieval_status: "found"` を混同しない。

既存 `get_record` の単一 ID 不在時の error behavior は変更しない。単一取得と複数取得では、有効な取得結果を保持する必要性が異なるためである。

### 5. Ordering と duplicate requested ID

Response item の順序は、request の `ids` における first occurrence order を維持する。

同一 ID が複数回指定された場合、`get_records` は最初の出現に対する item 一件だけを返す。重複指定は request error とせず、top-level の `duplicate_requested_id_ignored` / `info` diagnostic により可視化する。

これにより、無駄な response 重複を避けつつ、agent / client が冗長な request を送っていることを dogfooding や改善時に観測できる。

### 6. Raw body と response size

`include_body: true` の場合、body は既存 `get_record` と同じく raw Markdown 本文として扱い、要約・整形・正規化・truncate を行わない。

Response total length または body size の数値上限は public contract として定義しない。取得件数の分割で代替可能な transport / context budget を、Design Records MCP の意味論上の制限として固定しないためである。

実行環境上、完全な raw body を返せない場合に、不完全な本文を完全な body として返してはならない。具体的な execution error behavior は spec / implementation で定義する。

### 7. Contract example が示すべき境界

Spec の representative example は、少なくとも以下を示すものとする。

- `V01-ADR-077` のような `decision` record
- `SPEC-design-records-mcp-tools` のような `spec` record
- `V01-INV-DOCS-001` のような `investigation` record
- duplicate requested ID に対する informational diagnostic
- missing requested ID に対する item-level partial result

`REQ-*` / `WORK-*` / `TASK-*` は現行取得対象外であるため、現行 contract example の成功対象として含めない。

## 理由

### なぜ `get_records` を追加するか

`list_records` は本文を読む前に候補を絞ることには有効だが、選択済みの複数 record の headings / body を取得する経路は提供しない。

設計判断の確認では、単一の record だけでなく、根拠 ADR、現行 spec、関連 investigation を並べて読むことがある。これを一件ずつ取得させることは、Design Records MCP が既に持つ record-oriented read boundary の自然な不足である。

### なぜ filter / range query を追加しないか

`kind` / `status` / `id_range` / `limit` による候補列挙は既に `list_records` の責務である。`get_records` に同じ query 能力を追加すると、候補探索と選択済み詳細取得の境界が重複する。

`get_records` を明示 ID の取得に限定することで、利用者は `list_records` で探索し、`get_records` で選択済み根拠をまとめて読むという単純な導線を利用できる。

### なぜ partial result とするか

複数取得では、一件の stale ID または typo のために、他の正常な requested record の取得結果を捨てる合理性がない。

Missing item を item-level diagnostic として残すことで、取得成功結果と参照不備を同一 response 上で確認できる。

### なぜ duplicate を error にしないか

Duplicate requested ID は、record が取得不能であることを意味しない。重複分を返すと token と response 解釈の負荷が増えるが、request 全体を失敗させるほどの契約違反でもない。

First occurrence の結果だけを返し、info diagnostic として可視化することで、取得の継続性と冗長 request の観測可能性を両立する。

### なぜ public response size limit を定めないか

Public limit を定めても、利用者は request を分割して取得できる。数値上限を contract に持ち込むと、transport や client context budget に依存する制約が record retrieval の意味論に混入する。

一方、raw body が完全本文であるという性質は既存 `get_record` と整合させる必要があるため、truncate は許容しない。

## 却下した代替案

### 代替案A: `list_records` に body 取得を追加する

却下する。候補探索と選択済み record の詳細取得が混在し、`list_records` の絞り込み責務が膨らむ。

### 代替案B: `get_records` に filter / range query を持たせる

却下する。`list_records` と query 境界が重複し、同じ候補探索を二つの tool で定義することになる。

### 代替案C: missing requested ID が一件でもあれば batch 全体を error とする

却下する。取得可能な record の結果まで失い、batch retrieval の利用価値を下げる。

### 代替案D: duplicate requested ID を silently deduplicate する

却下する。取得結果は簡潔になるが、agent / client の冗長 request が観測不能になる。

### 代替案E: body を record 単位で truncate する

却下する。`get_record` が raw body を完全な元本文として返す既存の価値と整合しない。

### 代替案F: response total length の public numeric limit を定義する

却下する。取得単位の分割で代替できる環境依存制約を public semantic contract として固定する必要がない。

## 影響

### Design Records MCP spec への影響

`docs/spec/design-records-mcp/tools.md` は、`get_records` の public tool contract、request / response schema、item-level partial result、duplicate diagnostic、raw body behavior、representative example を定義する必要がある。

必要に応じて `docs/spec/design-records-mcp/overview.md` / `schema.md` も tool set と diagnostic vocabulary の追加に追従する。

### Implementation / tests への影響

Design Records MCP implementation は、既存 `get_record` の record representation を再利用しつつ、複数 ID 取得、first occurrence ordering、missing item partial result、duplicate informational diagnostic を実装する必要がある。

具体的な作業項目と完了条件は `V01-WORK-MCP-002` および対応する task file で追跡する。

### Requirement / work item への影響

`V01-REQ-MCP-002` は capability 採用結果を反映し、`V01-WORK-MCP-002` は spec 更新、implementation、tests、runtime verification へ進む。

## Evidence

- commit: tbd
- impl commit: tbd
- 参考: V01-ADR-077, V01-ADR-087, V01-REQ-MCP-002, V01-WORK-MCP-002, V01-TASK-MCP-002-01, V01-TASK-MCP-002-02, M19 close 後の Design Records MCP dogfooding

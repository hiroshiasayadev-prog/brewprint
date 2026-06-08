# V01-REQ-DATA-001: M15 v1.1.0-spec を minimum-expressiveness release として完了する

- **id**: V01-REQ-DATA-001
- **status**: accepted
- **date**: 2026-05-29
- **source_refs**:
  - V01-INV-DATA-001
  - V01-INV-DATA-002
  - V01-ADR-060
  - V01-ADR-061
  - V01-ADR-062
  - V01-ADR-063
  - V01-ADR-064
  - V01-ADR-067
  - V01-ADR-069
- **work_items**:
  - V01-WORK-DATA-001

## 要求

legacy M15 (`data layer expressiveness (v1.1)`) は、単なる flow-safety release として縮小せず、UC-002 self-hosting で観測された代表的な closed vocabulary の `str + note` 退避を machine-readable な型制約へ回収する **minimum-expressiveness release** として `v1.1.0-spec` を成立させる必要がある。

M15 の close boundary は、Phase A〜B4 の完遂に加えて、M15 Phase C から派生したうち release boundary として閉じた最小範囲である V01-ADR-069 minimum および V01-ADR-067 enum minimum を含むものとする。

## 発見根拠

- M15 は当初から、TypeRef / flow wiring / return source / DAG render だけでなく、UC-002 self-hosting に必要となる enum / discriminated object / inline struct / container complexity を Phase C として保持していた。
- `V01-INV-DATA-001` は、M15 が複数 ADR 系列へ派生したまま legacy record と workflow artifact へ回収されていないことを整理した。
- `V01-INV-DATA-002` は、現行 UC-002 YAML に schema / contract debt が 54 件あり、そのうち `enum_like_closed_vocabulary` が 19 件存在することを確認した。
- V01-ADR-067 は、TypeRef variant を増やさず named enum model として閉じた語彙を表現し、初期移行対象を 3 enum model に限定している。
- V01-ADR-069 は、M15 v1.1.0-spec で確定する範囲として parser safety limit、anonymous inline struct 非導入、`opaque_type_ref` warning 等を明示している。
- 後続 boundary review を踏まえ、ユーザーは `F1: B4 + enum minimum` を M15 / v1.1.0-spec の release boundary として採用した。

## 必要な結果

### M15 baseline の完遂

- V01-ADR-060〜V01-ADR-063 に由来する Phase A〜B3 の spec / implementation / verification evidence が、legacy checkbox ではなく後続 work item 上の evidence として確認できる。
- V01-ADR-064 に由来する Phase B4 の DAG renderer 実装、必要な current fixture / golden 更新、および regression verification が完了する。
- `v1.0.0-spec` の過去 snapshot は git tag に保持されるものとし、現行 UC-001 fixture は v1.1 現行仕様の regression fixture として更新可能であることを前提とする。

### V01-ADR-069 minimum の反映

- V01-ADR-069 §10 が M15 v1.1.0-spec に含めると定めた minimum scope を、spec / implementation / tests に反映する。
- minimum scope は、nested `list<T>` / `dict<T>` の維持、parser safety limit、anonymous inline struct 非導入、および `opaque_type_ref` warning を含む。
- `opaque_type_ref` は container TypeRef 内の `any` を対象とする warning であり、UC-002 に残る bare `any + note` debt 全体を可視化・解消するものとは扱わない。

### V01-ADR-067 enum minimum の反映

- V01-ADR-067 の acceptance 判断を行い、採用する場合は同 ADR §7 の初期移行対象に限定して enum model を導入する。
- 初期導入対象は以下の 3 enum model とする。
  - `mcp_object_type`
  - `mcp_diagnostic_severity`
  - `reference_tree_direction`
- enum model 定義追加と、以下の UC-002 field の `str + note` から named enum model TypeRef への切替は、unresolved model または移行漏れを生じさせない同一実行単位として扱う。
  - `object_selector.object`
  - `object_ref.object`
  - `diagnostic.severity`
  - `get_reference_tree_request.direction`
  - `get_reference_tree_response.direction`
- `get_references.direction`、`reference.direction`、object-dependent `kind`、`impact_severity`、`impact_fixability` は、初期移行対象に自動追加しない。

### Release close

- `v1.1.0-spec` は、flow wiring / return source / DAG render の整合に加え、UC-002 の共通 closed vocabulary の一部を machine-readable enum model として表現できる release として記録する。
- legacy M15 record は historical record として維持し、close 時に本 requirement が採用した boundary と、後続へ送った範囲を記録する。

## 明示的に除外する範囲

以下は重要な follow-up 候補であるが、本 requirement の `v1.1.0-spec` close blocker には含めない。

- V01-ADR-070 / V01-ADR-071 / V01-ADR-072 / V01-ADR-075 に由来する file-private helper model、private model render、model catalog、model file render
- V01-ADR-073 に由来する tagged union / discriminator payload の型表現
- V01-ADR-074 に由来する DAG asset node の TypeRef hint 表示
  - v1.1 では enum を machine-readable な named model として導入するが、DAG Mermaid 本体で `asset_name: type_hint` として可視化することまでは要求しない。
  - V01-ADR-074 を後続で採用する場合、enum model は既存の named model hint 規則により表示可能であり、enum 専用 render rule を別途導入する必要はない。
- V01-ADR-078 / V01-ADR-079 / V01-ADR-080 に由来する MCP semantic identity / state machine identity
- recursive `ObjectRef`、untagged union list、selector combination matrix、numeric range / default / cross-field behavior、usage-site-dependent vocabulary の機械可読化
- notes retreat の完全解消

## 後続 work item への制約

- 本 requirement は F1 boundary の採用を所有するが、具体 task graph、commit 順序、変更ファイル一覧、verification evidence は後続 work item / task が所有する。
- V01-ADR-069 minimum と V01-ADR-067 enum minimum は `docs/spec/type-ref.md` / `docs/spec/diagnostics.md` 等の更新 surface が重なり得るため、後続 work item は spec reflection と implementation の整列順を明示する。
- V01-ADR-070 / V01-ADR-073 / V01-ADR-074 / V01-ADR-078〜080 を、verification の都合や readability 改善を理由に暗黙に本 requirement の blocker へ逆流させない。

## Boundary

- 本 requirement は M15 / `v1.1.0-spec` close に必要な採用済み capability boundary を保持するものであり、現行仕様本文、設計判断本文、実装手順、進捗、fixture evidence を所有しない。
- V01-ADR-067 は `V01-WORK-DATA-001` により accepted として確定し、enum minimum の spec / implementation / UC-002 initial migration まで完了した。
- V01-ADR-069 minimum は UC-002 notes debt 全体の救済策ではなく、M15 が採用済み判断として確定すべき TypeRef safety / debt visibility の minimum baseline として扱う。
- F2 相当の helper-shape capability は v1.1 後の優先 follow-up 候補であるが、本 requirement の完了条件には含めない。

## Outcome

`V01-WORK-DATA-001` により、本 requirement が要求した F1 boundary は解消された。

- Phase A〜B3: V01-ADR-060〜V01-ADR-063 由来の TypeRef / flow wiring / foreach collected asset / task return source / initialized source baseline は実装・tests evidence として確認済み。
- Phase B4: V01-ADR-064 の `returns.source` / initialized source DAG render、UC-001 current golden regeneration、renderer regression は完了済み。
- V01-ADR-069 minimum: parser safety limit、anonymous inline struct 非導入、`opaque_type_ref` warning は spec / implementation / tests に反映済み。
- V01-ADR-067 enum minimum: `mcp_object_type` / `mcp_diagnostic_severity` / `reference_tree_direction` の 3 enum model と、初期 5 field migration は atomic に反映済み。
- Full verification: `go test ./...` passed after UC-001 current render regeneration.
- `v1.1.0-spec` snapshot is ready to tag after commit. Tag issuance itself remains a separate git operation.

Remaining follow-ups are explicitly outside this requirement: helper model / model render series, tagged union, DAG TypeRef hint, MCP / state identity, UC-002 duplicate task QID issue, and remaining notes retreat debt.

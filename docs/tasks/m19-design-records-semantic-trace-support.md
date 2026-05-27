# M19: Design Records MCP semantic trace support

- **status**: closed
- **closed_at**: 2026-05-26

## Goal

ADR-087 / ADR-088 と semantic traceability spec に従い、Design Records MCP に investigation record integration と canonical semantic/artifact ref resolve / validation を実装する。

M19 は `REQ-MCP-001` を実現する concrete implementation milestone であり、横断進捗は `WORK-MCP-001` が所有する。

## Source artifacts

- requirement: `REQ-MCP-001`
- work item: `WORK-MCP-001`
- ADR-087: Design Records MCP investigation support and semantic ref resolve
- ADR-088: Reduce semantic trace MVP to a canonical reference resolution foundation
- `docs/spec/design-records-mcp/`
- `docs/spec/concepts/traceability/`
- `docs/internal-design/resolver/semantic-ref-index.md`（implementation-facing design。MVP endpoint / acceptance target ではない）
- `docs/investigations/docs/INV-DOCS-001-investigation-artifact-format-and-lifecycle.md`
- `docs/investigations/docs/INV-DOCS-002-external-coverage-artifact-necessity.md`
- `docs/investigations/docs/INV-DOCS-003-internal-design-semantic-trace-mvp-necessity.md`

## Scope

M19 が扱うもの:

- Design Records MCP spec の concrete resolver / validation tool contract refinement
- `kind: investigation` の parser / index / query / validation 実装
- record response の common fields + kind-specific detail object への切替
- active `spec:` semantic ref の resolve 実装
- Design Records MCP record ID-as-ref (`ADR-*` / `SPEC-*` / `INV-*`) の resolve 実装
- investigation の `source_refs` / 記載済み `follow_up_results` / `follow_up_candidates` validation 実装
- `follow_up_candidates` の canonical form 検査と、canonical だが unresolved な candidate を `info` として可視化する実装
- noncanonical physical path reference diagnostic の contract と実装。`source_refs` / `follow_up_results` は error、`follow_up_candidates` は `info` とする
- contract / implementation / tests の同一切替単位での追従

## Non-goals

- `internal-design:` / `coverage:` / `COV-*` の resolve
- coverage mapping index / validation / query
- `maps_to` / `covers` / `validates` relation の導入または検証
- `yaml:` prefix の active 化
- fixture / golden traceability
- requirement / work item を Design Records MCP の record kind または public resolver contract に含める判断
- MCP writer tools
- UC-002 self-hosting 再構築

## Phase A: contract refinement

- [x] `docs/spec/design-records-mcp/{overview,schema,tools}.md` の ADR-087 / ADR-088 反映状況を確認する
- [x] semantic/artifact ref resolve tool の名称、request、response schema を確定する
- [x] active `spec:` ref と record ID-as-ref の resolver input / output contract を確定する
- [x] investigation validation diagnostic category / severity を確定する
- [x] ADR / spec の `depends_on` が `INV-*` を参照できる contract と、investigation record integration 前の既知 unresolved の扱いを確定する
- [x] noncanonical physical path reference diagnostic の category / severity を確定する
- [x] record response の新 contract への切替条件を確定する
- [x] `docs/spec/concepts/traceability/resolve-and-validation.md` と矛盾がないことを確認する

### Phase A contract result

- Public resolver tool 名は `resolve_reference` とする。Request は `{ "ref": string }`、response は `ref` / `ref_kind` / `status` / `target` / `diagnostics` を必須とする。
- Active resolve input は active `spec:` semantic ref と record ID-as-ref (`ADR-*` / `SPEC-*` / `INV-*`) のみとする。`internal-design:` / `coverage:` / `COV-*` / `REQ-*` / `WORK-*` / physical path / unsupported ID form は direct query では `status: unsupported` を返す。Reserved prefix `yaml:` の public resolver input / direct query response behavior、および investigation metadata validation behavior は MVP で定義しない。
- `spec:` document-level target は document path を、section-level target は document path と heading text を返す。入力 canonical ref は top-level `ref` に保持し、target に重複して返さない。MVP は section-level ref と document-level ref の親子 relation を public response として定義せず、section ref の文字列 prefix から親 document ref は推定しない。Record ID target は indexed record の path / kind / title / status を返す。
- Supported form が解決不能なら `unresolved_reference`、同一 ref が複数 target に解決されるなら `ambiguous_reference` とし、任意の一件を選択しない。Validation 側では duplicate を error として返す。
- Semantic ref validation category は `invalid_semantic_ref_declaration` / `missing_section_target` / `ambiguous_section_target` / `duplicate_semantic_ref` とする。Investigation validation category は `unresolved_source_ref` / `unresolved_follow_up_result` / `unresolved_follow_up_candidate`、`noncanonical_source_ref` / `noncanonical_follow_up_result` / `noncanonical_follow_up_candidate`、および metadata field 由来の `unsupported_reference` とする。
- `source_refs` / `follow_up_results` の unresolved・noncanonical・unsupported は error とし、`follow_up_candidates` の対応する状態は info とする。Investigation metadata が duplicate semantic ref または duplicate record ID を指す場合は field-specific diagnostic を追加せず、`duplicate_semantic_ref` / `duplicate_id` のみで報告する。`validate_records.ok` は error diagnostic の有無だけで決まる。
- Public record response は common fields + kind-specific detail object のみを返し、旧 flat metadata field との compatibility 併存を設けない。Parser / index / list / get / validate と tests を同一切替単位で更新する。
- ADR / spec の `depends_on` は `INV-*` を valid canonical record ID-as-ref として参照できる。開始時 baseline では `ADR-086 depends_on references missing record INV-DOCS-001` が既知 error として確認されており、Phase B/C/D で investigation integration 後の解消を検証する。

## Phase B0: internal design refinement

- [x] `docs/internal-design/resolver/semantic-ref-index.md` を finalized Phase A contract に追従させる
- [x] investigation metadata parser / record discovery / index integration route を具体化する
- [x] `resolve_reference` の lookup source と `validate_records` の shared resolution route を具体化する
- [x] common fields + kind-specific detail response への切替箇所と、旧 flat response を残さない移行 route を具体化する
- [x] `yaml:`、REQ / WORK / task、coverage / internal-design relation を implementation scope に持ち込まないことを確認する

## Phase B: investigation record integration

- [x] investigation metadata parser を実装する
- [x] `kind: investigation` を discovery / index / list / get / validate に追加する
- [x] common fields + investigation detail response を実装する
- [x] existing decision / spec response を新 contract に切り替える
- [x] flat response と detail object の不整合な併存を作らない

## Phase C: resolver / validation implementation

- [x] active `spec:` semantic ref の lookup を実装する
- [x] record ID-as-ref (`ADR-*` / `SPEC-*` / `INV-*`) の lookup を実装する
- [x] active `spec:` semantic ref / record ID の duplicate detection を実装する
- [x] investigation `source_refs` unresolved を error とする
- [x] investigation の記載済み `follow_up_results` unresolved を error とする
- [x] investigation `follow_up_candidates` の artifact reference が canonical form であることを検査する
- [x] canonical form の unresolved `follow_up_candidates` を orphan error にせず、予定された後続 artifact が未作成であることを示す `info` diagnostic として返す
- [x] `source_refs` / `follow_up_results` の physical path input は noncanonical `error`、artifact reference として記載された `follow_up_candidates` の physical path input は noncanonical candidate を示す `info` diagnostic として扱う
- [x] `internal-design:` / `coverage:` / `COV-*` を MVP required resolver input に含めない

## Phase D: verification

- [x] `INV-DOCS-001` を investigation record として取得し、record ID-as-ref `INV-DOCS-001` を解決できる test を追加する
- [x] `ADR-086` の `depends_on: INV-DOCS-001` が investigation record integration 後に resolve でき、現行 `missing_depends_on_target` error が解消される test を追加する
- [x] `INV-DOCS-001` の legacy path-based `follow_up_candidates` が、canonical resolution 成功の前提ではなく noncanonical candidate の `info` diagnostic 期待入力として扱われる test を追加する
- [x] active `spec:` semantic ref の document / section resolve test を追加する。Section response は path と heading text を返し、section ref の文字列 prefix から親 document ref を推定しないことを検証する
- [x] `spec:` declaration grammar error、section heading missing、同一 document 内の section heading ambiguity の validation test を追加する
- [x] supported canonical ref の unresolved response と、duplicate target に対する ambiguous response が任意 target を返さない contract test を追加する
- [x] investigation metadata が duplicate semantic ref または duplicate record ID を指す場合、field-specific diagnostic を追加せず `duplicate_semantic_ref` / `duplicate_id` のみを返す validation test を追加する
- [x] `internal-design:` / `coverage:` / `COV-*` / unsupported ID form の direct query が `status: unsupported` を返す contract test を追加する
- [x] `ADR-088` / `INV-DOCS-002` / `INV-DOCS-003` の ID-as-ref resolve test を追加する
- [x] unresolved `source_refs` / `follow_up_results` の error test を追加する
- [x] canonical form の unresolved `follow_up_candidates` が orphan error ではなく `info` diagnostic になる test を追加する
- [x] `source_refs` / `follow_up_results` の path-based input が noncanonical `error`、`follow_up_candidates` の path-based input が noncanonical `info` になる test を追加する
- [x] `internal-design:` / `coverage:` / `COV-*` を required MVP acceptance target としない contract test を追加する
- [x] old flat response を返さない contract test を追加する

## Done criteria

- [x] Design Records MCP spec が ADR-087 / ADR-088 に沿う concrete contract を定義している
- [x] Design Records MCP implementation が `investigation` kind を扱う
- [x] record response が ADR-087 の kind-specific detail contract に追従する
- [x] active `spec:` semantic ref と required record ID-as-ref の resolve が動作する
- [x] investigation canonical reference validation と noncanonical path diagnostic が動作する
- [x] coverage / internal-design realization relation を前提にしない acceptance tests が通る
- [x] `WORK-MCP-001` の status を反映できる状態になっている

## Close note

- `kind: investigation`、kind-specific detail response、`resolve_reference`、semantic / investigation reference validation を実装し、MCP runtime で確認した。
- Runtime `get_record("INV-DOCS-001")` は `kind: "investigation"` と investigation detail を返すことを確認した。
- Runtime `validate_records` は `ok: true` を返し、開始時 baseline の `ADR-086 depends_on references missing record INV-DOCS-001` error が解消されたことを確認した。
- Root document semantic ref (`spec:trace` / `spec:project-artifact-model`) は既存 docs intent に基づく canonical ref として grammar / implementation gap を修正し、runtime validation error が解消されたことを確認した。
- `validate_records` には `follow_up_candidates` 由来の info diagnostic が30件残る（`noncanonical_follow_up_candidate`: 15、`unsupported_reference`: 13、`unresolved_follow_up_candidate`: 2）。これは M19 contract 上の informational debt であり、M19 close blocker ではない。
- Follow-up requirement として `REQ-MCP-002`（batch retrieval capability）および `REQ-MCP-003`（requirement / work item / task MCP support）を captured 済みとする。

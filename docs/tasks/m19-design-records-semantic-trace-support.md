# M19: Design Records MCP semantic trace support

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

- [ ] `docs/spec/design-records-mcp/{overview,schema,tools}.md` の ADR-087 / ADR-088 反映状況を確認する
- [ ] semantic/artifact ref resolve tool の名称、request、response schema を確定する
- [ ] active `spec:` ref と record ID-as-ref の resolver input / output contract を確定する
- [ ] investigation validation diagnostic category / severity を確定する
- [ ] ADR / spec の `depends_on` が `INV-*` を参照できる contract と、investigation record integration 前の既知 unresolved の扱いを確定する
- [ ] noncanonical physical path reference diagnostic の category / severity を確定する
- [ ] record response の新 contract への切替条件を確定する
- [ ] `docs/spec/concepts/traceability/resolve-and-validation.md` と矛盾がないことを確認する

## Phase B: investigation record integration

- [ ] investigation metadata parser を実装する
- [ ] `kind: investigation` を discovery / index / list / get / validate に追加する
- [ ] common fields + investigation detail response を実装する
- [ ] existing decision / spec response を新 contract に切り替える
- [ ] flat response と detail object の不整合な併存を作らない

## Phase C: resolver / validation implementation

- [ ] active `spec:` semantic ref の lookup を実装する
- [ ] record ID-as-ref (`ADR-*` / `SPEC-*` / `INV-*`) の lookup を実装する
- [ ] active `spec:` semantic ref / record ID の duplicate detection を実装する
- [ ] investigation `source_refs` unresolved を error とする
- [ ] investigation の記載済み `follow_up_results` unresolved を error とする
- [ ] investigation `follow_up_candidates` の artifact reference が canonical form であることを検査する
- [ ] canonical form の unresolved `follow_up_candidates` を orphan error にせず、予定された後続 artifact が未作成であることを示す `info` diagnostic として返す
- [ ] `source_refs` / `follow_up_results` の physical path input は noncanonical `error`、artifact reference として記載された `follow_up_candidates` の physical path input は noncanonical candidate を示す `info` diagnostic として扱う
- [ ] `internal-design:` / `coverage:` / `COV-*` を MVP required resolver input に含めない

## Phase D: verification

- [ ] `INV-DOCS-001` を investigation record として取得し、record ID-as-ref `INV-DOCS-001` を解決できる test を追加する
- [ ] `ADR-086` の `depends_on: INV-DOCS-001` が investigation record integration 後に resolve でき、現行 `missing_depends_on_target` error が解消される test を追加する
- [ ] `INV-DOCS-001` の legacy path-based `follow_up_candidates` が、canonical resolution 成功の前提ではなく noncanonical candidate の `info` diagnostic 期待入力として扱われる test を追加する
- [ ] active `spec:` semantic ref の document / section resolve test を追加する
- [ ] `ADR-088` / `INV-DOCS-002` / `INV-DOCS-003` の ID-as-ref resolve test を追加する
- [ ] unresolved `source_refs` / `follow_up_results` の error test を追加する
- [ ] canonical form の unresolved `follow_up_candidates` が orphan error ではなく `info` diagnostic になる test を追加する
- [ ] `source_refs` / `follow_up_results` の path-based input が noncanonical `error`、`follow_up_candidates` の path-based input が noncanonical `info` になる test を追加する
- [ ] `internal-design:` / `coverage:` / `COV-*` を required MVP acceptance target としない contract test を追加する
- [ ] old flat response を返さない contract test を追加する

## Done criteria

- [ ] Design Records MCP spec が ADR-087 / ADR-088 に沿う concrete contract を定義している
- [ ] Design Records MCP implementation が `investigation` kind を扱う
- [ ] record response が ADR-087 の kind-specific detail contract に追従する
- [ ] active `spec:` semantic ref と required record ID-as-ref の resolve が動作する
- [ ] investigation canonical reference validation と noncanonical path diagnostic が動作する
- [ ] coverage / internal-design realization relation を前提にしない acceptance tests が通る
- [ ] `WORK-MCP-001` の status を反映できる状態になっている

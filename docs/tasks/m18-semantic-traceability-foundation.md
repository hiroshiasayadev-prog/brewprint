# M18: semantic traceability foundation

## Goal

semantic traceability MVP を、docs 上で運用できる最小状態にする。

M18 では、ADR-081〜084 で定義された requirements / work-items / internal-design / coverage / semantic ref の責務境界を、spec / policy / 最小 artifact layout に反映する。

## Scope

M18 が扱うもの:

- `docs/spec/concepts/traceability/` の draft review / refinement
- `docs/doc-policy.md` の ADR-081〜084 追従
- `docs/adr-authoring-guide.md` の責務表更新
- `docs/requirements/` / `docs/work-items/` / `docs/internal-design/` / `docs/coverage/` の最小配置方針
- semantic ref / coverage mapping の最小 example
- Design Records MCP に trace support を入れるかどうかの後続判断

## Non-goals

M18 では以下を扱わない。

- `yaml:` active 化
- brewprint DSL YAML entity-level semantic ref
- fixture / golden traceability
- `validates` relation
- MCP writer tools の具体 request / response schema
- UC-002 self-hosting 再構築
- spec directory 全体 taxonomy migration
- fixture `render_expected/` comparison semantics

## References

- ADR-081: requirements layer と semantic traceability
- ADR-082: golden fixture と self-hosting requirement の責務境界
- ADR-083: project artifact boundary と YAML as primary implementation source
- ADR-084: semantic trace MVP scope と artifact boundary
- `docs/spec/concepts/traceability/`

## Phase A: traceability spec review

- [ ] `docs/spec/concepts/traceability/index.md` をレビューする
- [ ] `docs/spec/concepts/traceability/semantic-ref.md` をレビューする
- [ ] `docs/spec/concepts/traceability/artifact-refs.md` をレビューする
- [ ] `docs/spec/concepts/traceability/metadata-schema.md` をレビューする
- [ ] `docs/spec/concepts/traceability/coverage-mapping.md` をレビューする
- [ ] `docs/spec/concepts/traceability/resolve-and-validation.md` をレビューする
- [ ] `docs/spec/concepts/traceability/out-of-scope.md` をレビューする
- [ ] file 間の用語揺れを確認する
- [ ] semantic ref grammar / prefix / ID-as-ref の矛盾を確認する
- [ ] coverage mapping schema が ADR-084 の MVP scope と一致しているか確認する
- [ ] 必要な修正を反映する
- [ ] status を `draft` から `confirmed` にする条件を判断する

## Phase B: docs policy alignment

- [ ] `docs/doc-policy.md` の docs 構成に `requirements/` を追加する
- [ ] `docs/doc-policy.md` の docs 構成に `work-items/` を追加する
- [ ] `docs/doc-policy.md` の docs 構成に `internal-design/` を追加する
- [ ] `docs/doc-policy.md` の docs 構成に `coverage/` を追加する
- [ ] `docs/doc-policy.md` の YAML 記述を ADR-083 に合わせて更新する
- [ ] `brewprint DSL YAML` と `trace metadata YAML` の用語を反映する
- [ ] `docs/uc/` を golden fixture corpus として ADR-082 に合わせて更新する
- [ ] semantic ref を physical path ではなく primary key とする方針を追加する

## Phase C: ADR authoring guide alignment

- [ ] `docs/adr-authoring-guide.md` の責務表に requirements を追加する
- [ ] `docs/adr-authoring-guide.md` の責務表に work-items を追加する
- [ ] `docs/adr-authoring-guide.md` の責務表に internal-design を追加する
- [ ] `docs/adr-authoring-guide.md` の責務表に coverage を追加する
- [ ] UC docs / fixture が gap discovery log や migration state を所有するように読める記述を修正する
- [ ] ADR に semantic ref / trace detail を抱え込ませすぎない方針を追加する
- [ ] ADR は現行仕様本文ではなく、spec 由来注記と semantic ref から辿る方針を確認する

## Phase D: minimal artifact layout

- [ ] `docs/requirements/` の最小配置方針を決める
- [ ] `docs/work-items/` の最小配置方針を決める
- [ ] `docs/internal-design/` の最小配置方針を決める
- [ ] `docs/coverage/` の最小配置方針を決める
- [ ] 各 directory に README / index / placeholder が必要か判断する
- [ ] requirement ID domain の初期 vocabulary を決めるか判断する
- [ ] work item ID domain の初期 vocabulary を決めるか判断する
- [ ] coverage mapping ID domain の初期 vocabulary を決めるか判断する

## Phase E: first trace example

- [ ] `spec:trace.*` の section refs を確認する
- [ ] traceability spec 自体を使った最小 example を決める
- [ ] 対応する `internal-design:` ref の要否を判断する
- [ ] 最小 coverage mapping example を作る
- [ ] resolver / validation の手動チェック観点を整理する
- [ ] example が fixture / golden / `validates` relation に流れていないか確認する

## Phase F: Design Records MCP follow-up decision

- [ ] Design Records MCP に semantic ref resolve を入れるか判断する
- [ ] Design Records MCP に trace metadata validation を入れるか判断する
- [ ] Design Records MCP に coverage mapping query を入れるか判断する
- [ ] MCP writer tools は M18 scope 外であることを確認する
- [ ] 必要なら後続 ADR / requirement / milestone を起票する

## Done criteria

- [ ] `docs/spec/concepts/traceability/` の MVP spec が review 済みである
- [ ] traceability spec の status を confirmed にするか、remaining draft reason が明示されている
- [ ] `docs/doc-policy.md` が ADR-081〜084 と矛盾しない
- [ ] `docs/adr-authoring-guide.md` が ADR-081〜084 と矛盾しない
- [ ] requirements / work-items / internal-design / coverage の最小配置方針が決まっている
- [ ] 最小 semantic trace example が存在する、または作らない理由が明示されている
- [ ] Design Records MCP trace support の扱いが後続作業として整理されている

# M18: semantic traceability foundation

## Goal

Semantic traceability MVP を、docs 上で運用できる最小状態にする。

ADR-088 により、MVP は semantic realization mapping ではなく **canonical reference resolution foundation** に縮小された。M18 は、`spec:` semantic ref、record ID-as-ref、および investigation canonical reference validation の責務境界を spec / policy / authoring guidance / M19 handoff に同期する。

## Scope

M18 が扱うもの:

- `docs/spec/concepts/project-artifact-model/` による artifact system 全体像の current spec 化
- `docs/spec/concepts/traceability/` の ADR-088 追従
- `docs/doc-policy.md` と authoring README の scope 同期
- requirements / work-items / internal-design / investigations の配置方針
- external relation / assurance artifact は MVP で配置・authoring entrance を持たず、必要時に新設判断する境界の整理
- active `spec:` semantic ref と investigation canonical reference の docs-side boundary
- ADR-087 / ADR-088 に基づく Design Records MCP implementation handoff の M19 への同期

## Non-goals

M18 では以下を扱わない。

- `internal-design:` / `coverage:` / `COV-*` の MVP active 化
- semantic realization relation (`maps_to` / `covers` / `validates`) の operational 導入
- external coverage mapping artifact の MVP 導入
- `yaml:` active 化、brewprint DSL YAML entity-level semantic ref
- fixture / golden traceability
- MCP writer tools の具体 request / response schema
- Design Records MCP の concrete resolve / validation contract と実装（M19 で追跡する）
- UC-002 self-hosting 再構築
- spec directory 全体 taxonomy migration

## References

- ADR-081: requirements layer と semantic traceability
- ADR-082: golden fixture と self-hosting requirement の責務境界
- ADR-083: project artifact boundary と YAML as primary implementation source
- ADR-084: semantic trace MVP scope と artifact boundary
- ADR-085: investigation artifact boundary
- ADR-086: investigation artifact format and lifecycle
- ADR-087: Design Records MCP investigation support and semantic ref resolve
- ADR-088: Reduce semantic trace MVP to a canonical reference resolution foundation
- `docs/investigations/docs/INV-DOCS-002-external-coverage-artifact-necessity.md`
- `docs/investigations/docs/INV-DOCS-003-internal-design-semantic-trace-mvp-necessity.md`
- `docs/spec/concepts/project-artifact-model/`
- `docs/spec/concepts/traceability/`

## Phase A: original traceability draft review

- [x] `docs/spec/concepts/traceability/` 一式をレビューする
- [x] semantic ref grammar / prefix / ID-as-ref の矛盾を確認する
- [x] ADR-087 の investigation canonical reference boundary を反映する
- [x] initial mapping draft と `COV-TRACE-001` example を作成し、relation / endpoint 問題を発見する
- [x] independent review findings を ADR 判断へ切り出す

## Phase B: investigation and decision refinement

- [x] `INV-DOCS-002` により external coverage artifact の MVP 必要性を評価する
- [x] `INV-DOCS-003` により `internal-design:` endpoint / semantic realization relation 自体の MVP 必要性を評価する
- [x] ADR-088 を corrected decision として accepted にし、MVP を canonical reference resolution foundation に縮小する
- [x] `internal-design:` / `coverage:` / `COV-*` / realization relation を MVP scope 外へ送る判断を記録する

## Phase C: traceability spec synchronization

- [x] `index.md` を `spec:` + record / investigation canonical references 中心へ更新する
- [x] `semantic-ref.md` の MVP active example を `spec:` のみに整理する
- [x] `artifact-refs.md` で `spec:` を active、`yaml:` を reserved、`internal-design:` / `coverage:` を deferred と整理する
- [x] `metadata-schema.md` から internal-design / coverage mapping metadata の MVP contract を外す
- [x] `coverage-mapping.md` を realization mapping defer boundary / reintroduction trigger の spec として更新する
- [x] `resolve-and-validation.md` を `spec:` / record ID / investigation canonical ref validation 中心へ更新する
- [x] `out-of-scope.md` に internal-design endpoint / external coverage / relation 再導入 trigger を記録する
- [x] `docs/coverage/traceability/semantic-ref.yaml` を MVP example artifact から除去する

## Phase D: policy and authoring guidance synchronization

- [x] `docs/spec/concepts/project-artifact-model/index.md` から internal-design relation owner の旧中間案を除去する
- [x] `docs/spec/concepts/project-artifact-model/index.md` を canonical reference foundation boundary に同期する
- [x] `docs/doc-policy.md` の MVP active prefix / coverage 説明を ADR-088 に同期する
- [x] `docs/internal-design/README.md` に layer 存続と MVP endpoint 非採用の境界を反映する
- [x] MVP では external relation / assurance artifact 用 directory や authoring entrance を設けず、必要時に配置を含めて新設判断する境界を反映する
- [x] `docs/spec-authoring-guide.md` に `docs/spec/concepts/` の入口と project artifact model 導線を反映する

## Phase E: implementation handoff synchronization

- [x] `docs/internal-design/resolver/semantic-ref-index.md` を implementation-facing design として残し、MVP mapping example / endpoint の役割を外す
- [x] M19 から `COV-TRACE-001`、coverage mapping validation、internal-design relation acceptance を外す
- [x] M19 を active `spec:` / record ID / investigation canonical reference validation の implementation milestone に同期する
- [x] `REQ-MCP-001` / `WORK-MCP-001` の ADR-088 および acceptance boundary 追従を反映する
- [x] `docs/spec/design-records-mcp/` の concrete contract draft に old mapping / endpoint 前提が残っていないか確認し、必要な boundary 同期を反映する

## Phase F: final review and close gate

- [x] ADR-088 に基づく spec / policy / README / task / example cleanup 一式を独立レビューする
- [x] review findings のうち M18 同期差分へ反映すべきものを反映し、M19 / future convention work へ送る事項を整理する
- [x] independent review により M18 close は妥当と判断され、milestone status を `closed` に反映する

## Review outcome

- 初期 M18 draft は `spec:` → `internal-design:` の `COV-TRACE-001` mapping を MVP example として作成したが、これは resolver 自身を対象とする bootstrap example であり、realization relation の independent requirement ではなかった。
- `INV-DOCS-002` は external coverage artifact を MVP に残す実需が確認できないことを示した。
- `INV-DOCS-003` は `internal-design:` endpoint と realization relation 自体についても MVP 必須性が確認できないことを示した。
- Corrected ADR-088 は、MVP を `spec:` semantic ref、record ID-as-ref、investigation canonical reference validation からなる foundation に縮小した。
- `docs/internal-design/` layer は implementation-facing documentation boundary として残るが、MVP semantic endpoint ではない。
- External relation / assurance artifact は MVP required authoring entrance / example / acceptance target ではなく、必要性が成立した時点で配置と責務を含めて新設判断する。
- Final independent review は M18 close を妥当と確認した。Traceability spec は M19 で concrete resolver / validation contract と実装検証が追従するまで `draft` を維持する。

## Deferred follow-ups from final review

- M19 で、`ADR-086` の `depends_on: INV-DOCS-001` が investigation record integration 後に resolve でき、現行の `missing_depends_on_target` error が解消されることを test する。
- M19 contract refinement で、ADR / spec の `depends_on` が `INV-*` を参照できる boundary と implementation 移行前の既知 unresolved の扱いを確定する。
- `migrated_to_spec` の付与基準は convention が未確定であるため、ADR-088 では未指定へ戻し、必要なら別の小さな doc-policy / authoring-guide 作業で定める。
- `coverage-mapping.md` の filename / semantic ref は historical identity を保ったまま boundary spec として読めるため、本 milestone では rename しない。

## Done criteria

- [x] ADR-088 の corrected MVP boundary が traceability leaf spec に反映されている
- [x] project artifact model / doc-policy / internal-design README が ADR-088 と整合し、external artifact 用 directory を MVP layout として予約していない
- [x] M19 が coverage / internal-design realization relation を acceptance target としない形に同期されている
- [x] 旧 coverage mapping example が current MVP artifact として残っていない
- [x] REQ-MCP-001 / WORK-MCP-001 と Design Records MCP spec が ADR-088 の M19 boundary と整合する
- [x] project artifact model への導線が spec authoring entrance に反映されている
- [x] 更新後 artifact 一式について independent review が完了し、M18 close が妥当と確認されている

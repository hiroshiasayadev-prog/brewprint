# TASK-DATA-001-02: Phase A〜B4 の evidence と ADR-064 renderer を完了する

- **id**: TASK-DATA-001-02
- **status**: done
- **date**: 2026-05-29
- **work_item**: WORK-DATA-001
- **source_requirement**: REQ-DATA-001
- **estimate**: 2d
- **depends_on**:
  - TASK-DATA-001-01
- **outputs**:
  - ADR-060〜ADR-063 の implementation / test evidence 確認結果
  - ADR-064 DAG renderer 実装と current fixture / golden 更新
  - Phase A〜B4 regression verification evidence

## Goal

M15 baseline である Phase A〜B4 を閉じるため、既に完了扱いとなっている Phase A〜B3 の証跡を独立に確認し、未完の ADR-064 DAG render 実装・fixture・regression verification を完了する。

## Work

- ADR-060〜ADR-063、関連 spec、implementation、tests を照合し、legacy M15 checkbox に依存せず Phase A〜B3 の完了 evidence を整理する。
- ADR-064 の accepted decision と `docs/spec/views/dag.md` の反映内容を照合し、renderer implementation の不足を確定する。
- `returns.source` および initialized source の data line render を implementation と tests / golden に反映する。
- 必要な UC-001 current fixture / golden を v1.1 現行仕様に対する regression evidence として更新する。
- `v1.0.0-spec` tag に保持された過去 snapshot と、更新する current fixture を混同しない。
- ADR-074 の DAG asset TypeRef hint は本 task に混入させない。

## Done condition

- ADR-060〜ADR-063 の implementation / tests / evidence の存在確認結果が記録されている。
- ADR-064 の accepted render boundary に必要な実装と golden / fixture 更新が完了している。
- Phase A〜B4 に関する対象 tests / regression verification が passing であり、結果が evidence として記録されている。
- ADR-074 の TypeRef hint 表示を追加していない。

## Verification

- ADR-060〜ADR-064 と対応 spec / implementation / tests / golden の差分を照合する。
- renderer 関連 test と、影響範囲に応じた regression test を実行する。
- 更新した current fixture / golden が v1.1 向け変更であり、過去 snapshot を破壊する扱いになっていないことを確認する。

## Evidence

- Codex implementation gap review returned `Blocking gaps` for TASK-DATA-001-02.
- ADR-060〜ADR-063 resolver baseline was reported as implemented and tested by existing tests (`flow_wiring_type_test.go`, `foreach_returns_test.go`, `return_source_test.go`).
- Blocking gaps were concentrated in ADR-064 DAG renderer / UC-001 current golden:
  - `subgraph returns` still emitted even though ADR-064 abolished it.
  - `_end` return data line from `returns.source` was missing.
  - initialized store boundary / `initStoreNode` rendering was missing.
  - DAG task detail Returns table omitted `source` column.
  - UC-001 `validate_cart` fixture lacked `returns.source: validated_items`.
- Implemented draft ADR-064 renderer changes:
  - Removed `subgraph returns` emission from flow and single-task DAG renderers.
  - Added `subgraph initializes` / `initStoreNode` support for flow renderer and initial single-task path.
  - Added return data line from `returns.source` to `_end` with label `returns as <returns.name>`.
  - Ensured task / foreach / join produced return values are rendered as asset nodes before return data line use.
  - Added `source` column to Returns detail table via shared `writeReturnsSection`.
- Updated UC-001 `cart/task/validate_cart.yaml` with `returns.source: validated_items` to exercise foreach collected asset return.
- Local verification after `gofmt` confirmed implementation compiles, but DAG golden tests fail because UC-001 current golden is still pre-ADR-064 shape.
- Passing checks: `go test ./internal/render/project`, `go test ./cmd/brewprint`.
- Expected failing checks: `go test ./internal/render/dag` and `go test ./...`, both failing only on `TestRenderDAGGolden` mismatches.
- Observed mismatches align with ADR-064 renderer changes: old goldens still expect `subgraph returns`; actual output emits no returns boundary, emits `subgraph initializes` where applicable, and adds Returns table `source` column.
- Regenerated UC-001 current renders with `go run ./cmd/brewprint render --yaml-root docs\uc\001-ec-checkout-flow\yaml --out docs\uc\001-ec-checkout-flow\renders --clean`; output: `rendered 23 file(s)`.
- `go test ./internal/render/dag`: passed.
- `go test ./internal/render/project`: passed.
- `go test ./cmd/brewprint`: passed.
- `go test ./...`: passed.
- UC-001 current DAG goldens now represent v1.1 ADR-064 behavior; `v1.0.0-spec` remains preserved by tag and was not retroactively modified.
- ADR-074 TypeRef hint rendering was not introduced in this task.

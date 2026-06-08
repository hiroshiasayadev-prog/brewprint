# V01-TASK-RESOLVE-001-03: regression tests と UC-002 validate / render verification を行う

- **id**: V01-TASK-RESOLVE-001-03
- **status**: done
- **date**: 2026-05-31
- **work_item**: V01-WORK-RESOLVE-001
- **source_requirement**: V01-REQ-RESOLVE-001
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-RESOLVE-001-02
- **outputs**:
  - resolver regression tests
  - UC-002 validate / render verification evidence
  - V01-WORK-RESOLVE-001 close readiness evidence

## Goal

`V01-TASK-RESOLVE-001-02` の resolver 修正が V01-ADR-058 の file-private sub node contract を満たし、UC-002 duplicate task QID / unresolved flow task issue を解消したことを検証する。

## Work

- resolver regression tests を追加または更新する。
- 少なくとも以下を確認する。
  - 別 file 間の同名 private sub task local ID は許容される。
  - 同一 file 内の同名 private sub task local ID は diagnostic になる。
  - 同一 module 内の同名 main task は `duplicate_node` になる。
  - local flow step は同一 file の private sub task を優先して解決する。
- `go test ./...` を実行する。
- `go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml` を実行し、duplicate task QID / unresolved flow task が解消したことを確認する。
- 必要に応じて UC-002 render を実行し、render pipeline が同 issue で停止しないことを確認する。
- 検証結果を task / work item / requirement の close evidence として反映できる形に整理する。

## Done condition

- resolver regression tests が追加または更新されている。
- `go test ./...` が pass している。
- UC-002 validate が duplicate task QID / unresolved flow task issue で fail しない。
- UC-002 render が同 issue で停止しない。
- 残る failure がある場合、それが本 requirement の範囲外かどうか分類されている。

## Verification

- `go test ./...`
- `go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml`
- UC-002 render command。具体 command は現行 CLI / render workflow に従って実行する。

## Evidence

- Regression tests were added / updated as part of the `V01-TASK-RESOLVE-001-02` implementation.
- Regression coverage includes:
  - cross-file same private sub task local ID is allowed;
  - same-file duplicate private sub task local ID emits `duplicate_sub_node`;
  - same-module duplicate main task emits `duplicate_node`;
  - local flow step resolves same-file private sub task first;
  - duplicate private sub task handling does not cascade into `unresolved_flow_task`;
  - cross-file same private sub task local ID with same `returns.name` does not collide in asset ID / object key;
  - public-shaped private alias is not registered in public lookup indexes;
  - full QID transition action does not validate through a private alias;
  - asset query paths do not first-hit the wrong file-private asset.
- Verification reported by implementation handoff:
  - `go test ./...` -> pass
  - `go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml` -> ok
  - `go run ./cmd/brewprint render --yaml-root docs\uc\002-brewprint-self-hosting\yaml --out $env:TEMP\brewprint-uc002-render-review2 --clean` -> rendered 11 file(s)
- UC-002 duplicate task QID / unresolved flow task issue is resolved.
- UC-002 has no remaining diagnostics in the reported verification.
- No remaining failure was reported for this requirement scope.
- Close readiness: `V01-WORK-RESOLVE-001` can proceed to close review after this task update, because V01-TASK-RESOLVE-001-01 / 02 / 03 are now done.

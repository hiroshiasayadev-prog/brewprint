# V01-TASK-DATA-001-05: M15 close と v1.1.0-spec snapshot evidence を確定する

- **id**: V01-TASK-DATA-001-05
- **status**: done
- **date**: 2026-05-29
- **work_item**: V01-WORK-DATA-001
- **source_requirement**: V01-REQ-DATA-001
- **estimate**: 1d
- **depends_on**:
  - V01-TASK-DATA-001-02
  - V01-TASK-DATA-001-03
  - V01-TASK-DATA-001-04
- **outputs**:
  - Full verification と v1.1.0-spec snapshot evidence
  - legacy M15 historical close boundary 記録
  - V01-REQ-DATA-001 / V01-WORK-DATA-001 の close outcome 更新

## Goal

F1 boundary の全反映・検証結果を照合し、legacy M15 を minimum-expressiveness release として閉じ、`v1.1.0-spec` snapshot を確定できる evidence と workflow artifact の close outcome を記録する。

## Work

- `V01-TASK-DATA-001-01`〜`04` の decision / spec / implementation / YAML / fixture / test evidence を照合する。
- `V01-REQ-DATA-001` と `V01-WORK-DATA-001` の completion condition が満たされたかを判定する。
- 影響範囲に応じた full verification を実施し、release snapshot を妨げる blocker の有無を確認する。
- Legacy `docs/tasks/m15-data-layer-expressiveness.md` に historical close note を反映する。
  - 実 close boundary が F1 であること
  - `minimum-expressiveness release` として閉じること
  - V01-ADR-069 minimum / V01-ADR-067 enum minimum を含めたこと
  - V01-ADR-070 / V01-ADR-073 / V01-ADR-074 / V01-ADR-078〜080 および remaining notes debt を follow-up へ送ったこと
- `V01-REQ-DATA-001` / `V01-WORK-DATA-001` に解消結果、verification evidence、残る follow-up boundary を反映する。
- `v1.1.0-spec` tag 発行に必要な snapshot evidence と、tag 発行実施の扱いを記録する。

## Done condition

- F1 boundary に必要な decision / spec / implementation / YAML / fixture / verification evidence が揃っている。
- `V01-WORK-DATA-001` の completion condition の充足または未充足理由が記録されている。
- Legacy M15 が進捗正本ではなく historical record として、close boundary と follow-up 分離を保持している。
- `V01-REQ-DATA-001` / `V01-WORK-DATA-001` の status / outcome / evidence が実結果と一致している。
- `v1.1.0-spec` snapshot を確定できる状態、または発行を妨げる blocker が明示されている。
- Follow-up scope が M15 blocker として再混入していない。

## Verification

- 配下 task evidence と work item completion condition を照合する。
- 必要な automated tests / fixture regression / repository-wide verification を実行する。
- Workflow artifact relation が ID-as-ref を維持し、task status の手動 checkbox 複製を導入していないことを確認する。
- `v1.0.0-spec` snapshot と v1.1 current fixture / snapshot evidence の扱いを混同していないことを確認する。

## Evidence

- Verified upstream task completion:
  - `V01-TASK-DATA-001-01`: V01-ADR-067 enum minimum acceptance gate completed.
  - `V01-TASK-DATA-001-02`: V01-ADR-060〜V01-ADR-064 baseline / renderer / UC-001 current golden regeneration completed.
  - `V01-TASK-DATA-001-03`: V01-ADR-069 parser safety / `opaque_type_ref` warning implementation and tests completed.
  - `V01-TASK-DATA-001-04`: V01-ADR-067 enum model implementation and UC-002 initial 3 enum model / 5 field atomic migration completed.
- Release boundary is F1: Phase A〜B4 + V01-ADR-069 minimum + V01-ADR-067 enum minimum.
- Full verification evidence for this close:
  - `go run ./cmd/brewprint render --yaml-root docs\uc\001-ec-checkout-flow\yaml --out docs\uc\001-ec-checkout-flow\renders --clean`: `rendered 23 file(s)`.
  - `go test ./internal/render/dag`: passed.
  - `go test ./internal/render/project`: passed.
  - `go test ./cmd/brewprint`: passed.
  - `go test ./internal/resolve`: passed as part of previous `V01-TASK-DATA-001-03/04` verification and full regression.
  - `go test ./...`: passed after UC-001 current render regeneration.
- UC-001 current renders now act as v1.1 current golden evidence for V01-ADR-064. `v1.0.0-spec` remains preserved by git tag and was not retroactively modified.
- UC-002 enum migration was completed for exactly the initial 3 enum model / 5 field scope. UC-002 full validate / render remains blocked by pre-existing duplicate task QID / unresolved flow task diagnostics and is not treated as an enum migration or M15 close blocker.
- Follow-up scope explicitly deferred outside M15 / `v1.1.0-spec`: V01-ADR-070 / V01-ADR-071 / V01-ADR-072 / V01-ADR-075 helper/model render series, V01-ADR-073 tagged union, V01-ADR-074 DAG asset TypeRef hint, V01-ADR-078〜080 MCP / state identity, and remaining UC-002 notes retreat debt.
- Legacy `docs/tasks/m15-data-layer-expressiveness.md`, `V01-REQ-DATA-001`, `V01-WORK-DATA-001`, and `docs/TASKS.md` were synchronized with this close outcome.
- `v1.1.0-spec` snapshot is ready to tag after commit. Tag issuance itself is intentionally recorded as pending commit/tag operation, not performed in this document edit.

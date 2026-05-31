# TASK-DATA-001-04: ADR-067 enum minimum を atomic に反映する

- **id**: TASK-DATA-001-04
- **status**: done
- **date**: 2026-05-29
- **work_item**: WORK-DATA-001
- **source_requirement**: REQ-DATA-001
- **estimate**: 2.5d
- **depends_on**:
  - TASK-DATA-001-01
  - TASK-DATA-001-03
- **outputs**:
  - ADR-067 enum minimum に対応する spec / implementation / tests 更新
  - UC-002 の初期 3 enum model 定義と 5 field migration
  - enum migration scope / deferred debt の evidence

## Goal

ADR-067 の accepted enum minimum を spec・implementation・UC-002 YAML・tests に一貫して反映し、UC-002 の中央 closed vocabulary の一部を `str + note` 退避から machine-readable な named enum model へ移行する。

## Work

- ADR-067 の accepted boundary と `TASK-DATA-001-03` 後の共有 spec surface を前提に、enum model syntax / validation / nominal compatibility の spec 差分を反映する。
- enum model 表現、validation diagnostics、TypeRef compatibility に必要な implementation / tests を追加または更新する。
- 以下の 3 enum model definition を UC-002 YAML に追加する。
  - `mcp_object_type`
  - `mcp_diagnostic_severity`
  - `reference_tree_direction`
- 以下の初期 field を `str + note` から named enum model TypeRef へ切り替える。
  - `object_selector.object`
  - `object_ref.object`
  - `diagnostic.severity`
  - `get_reference_tree_request.direction`
  - `get_reference_tree_response.direction`
- Enum definitions と初期 field migration は同一実行単位として反映し、unresolved model または definition-only の中間状態を成果物に残さない。
- 以下は類似 debt として観測されても本 task の migration に追加しない。
  - `get_references.direction`
  - `reference.direction`
  - object-dependent `kind`
  - `impact_severity`
  - `impact_fixability`
- Helper model、tagged union、DAG TypeRef hint、MCP / state identity を本 task に混入させない。

## Done condition

- ADR-067 enum minimum の spec / implementation / diagnostics / tests が整合している。
- 初期 3 enum model definition と 5 field migration が atomic に反映され、YAML validation が成立している。
- Nominal enum compatibility と invalid enum model / duplicate enum value 等の採用済み diagnostic boundary が local tests で確認されている。
- Follow-up へ送る enum-like debt と、除外 scope が evidence に記録されている。
- ADR-070 / ADR-073 / ADR-074 / ADR-078〜080 が本 task の完了条件へ逆流していない。

## Verification

- ADR-067 acceptance 内容と spec / YAML / implementation diff を照合する。
- Enum parser / semantic validation / TypeRef compatibility / diagnostic tests を local environment で実行する。
- UC-002 YAML の model definition と利用箇所を読み戻し、初期5 field のみが意図した scope で移行されていることを確認する。
- 影響範囲に応じた repository test を local environment で実行し、結果を evidence に記録する。

## Evidence

- Spec-only draft reflection performed for Codex review.
- Updated `docs/spec/nodes.md` with `model.kind: enum`, required `values`, enum field constraints, and absence of `fields` / `element` / `value` on enum models.
- Updated `docs/spec/type-ref.md` with enum named model TypeRef usage and nominal compatibility boundary, including no implicit compatibility between enum and `str`.
- Updated `docs/spec/diagnostics.md` with `invalid_enum_model` and `duplicate_enum_value`, while keeping `invalid_enum_value` out of v1.1 minimum.
- Codex review returned OK with minor fixes; fixed `invalid_enum_model` diagnostic condition and moved enum diagnostic detail under model validation responsibility.
- Implemented enum model support in raw YAML / semantic model / builder.
- Added `enum` to valid model kinds and implemented `invalid_enum_model` / `duplicate_enum_value` validation.
- Added enum validation and nominal compatibility tests.
- Added UC-002 enum definitions for `mcp_object_type`, `mcp_diagnostic_severity`, and `reference_tree_direction` in `common.yaml`.
- Migrated exactly the initial 5 fields to named enum TypeRef: `object_selector.object`, `object_ref.object`, `diagnostic.severity`, `get_reference_tree_request.direction`, `get_reference_tree_response.direction`.
- Did not migrate `get_references.direction`, `reference.direction`, object-dependent `kind`, `impact_severity`, or `impact_fixability`.
- Local verification completed after `TASK-DATA-001-03` verification.
- `go test ./internal/resolve`: passed.
- `go test ./...`: passed.
- UC-002 full `validate` / `render` still fail with pre-existing duplicate task QID / unresolved flow task diagnostics for repeated `build_response`, `query_service`, and `validate_request` task IDs across MCP task files.
- No enum-related validation error was observed in the UC-002 validation output: no `invalid_enum_model`, `duplicate_enum_value`, `invalid_type_ref`, or enum model unresolved diagnostic appeared.
- Therefore this task treats the enum implementation and atomic 3 enum model / 5 field migration as complete, while leaving the UC-002 duplicate task issue outside this task boundary.

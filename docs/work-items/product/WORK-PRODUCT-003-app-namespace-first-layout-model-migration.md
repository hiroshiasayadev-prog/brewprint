# WORK-PRODUCT-003: App namespace-first layout model 仕様化と migration 方針決定

- **id**: WORK-PRODUCT-003
- **status**: not_started
- **date**: 2026-06-08
- **source_requirement**: REQ-PRODUCT-003
- **impact_refs**:
- **tasks**:

## Goal

REQ-PRODUCT-003 が定義した app namespace-first repository layout model を spec として固定し、migration の実行可否と互換方針を決定する。実ファイル移動は本 WORK のスコープ外とする。

## Boundary

このWORKが所有するもの:

- namespace-model spec への app namespace-first layout model セクション追加
- 現状 `docs/` と `internal/` の migration impact inventory
- compatibility / legacy path support 方針の決定
- migration 実行 or explicit defer の判断と記録

このWORKが所有しないもの:

- 実際のファイル・ディレクトリの移動
- BPDSL による DSL 生成の実装
- MCP API のパス変更対応
- UI MVP の情報設計

## Task flow

```mermaid
flowchart TD
  T01[TASK-01: namespace-model spec\nlayout model セクション追加] -->|spec review ゲート| T02[TASK-02: migration impact inventory\ndocs/ と internal/ の棚卸し]
  T02 --> T03[TASK-03: compatibility / legacy path\n方針決定]
  T03 --> T04[TASK-04: migration plan or explicit defer]
  T04 --> DONE[完了]
```

## Task Candidates

- TASK-01: `docs/spec/concepts/namespace-model/index.md` に app namespace-first layout model を追記（`records/` / `dsl/` / `src/` の意味・用途・任意性、`dsl/ → src/` 生成パターンを全 app namespace の意図する方向として明示） → spec review ゲート
- TASK-02: 現状 `docs/` と `internal/` の移動対象ファイル・ディレクトリを棚卸しし、app namespace 別の migration 対象パス一覧を作成
- TASK-03: 旧パス（`docs/requirements/mcp/` 等）の扱い、MCP のパス変更対応方針、migration timing を決定し記録
- TASK-04: migration を v0.2.0 で実行するか explicit defer にするかを判断し、実行する場合は別 WORK を切り出す

## Completion Condition

- namespace-model spec に app namespace-first layout model が記述されている
- migration 対象の棚卸しが完了し、パス一覧が存在する
- compatibility / legacy path 方針が決定されている
- migration 実行または explicit defer が記録されている
- REQ-PRODUCT-003 が accepted に更新されている

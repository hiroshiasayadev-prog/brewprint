# TASK-PRODUCT-001-03: v2 artifact ID grammar と mapping rule の spec 化

- **id**: TASK-PRODUCT-001-03
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-PRODUCT-001
- **source_requirement**: REQ-PRODUCT-001
- **estimate**: 0.5d
- **depends_on**:
  - TASK-PRODUCT-001-02
- **outputs**:
  - v2 artifact ID grammar 仕様（spec section）
  - existing ID → namespace-aware ID の mapping rule（spec section）

## Goal

REQ-PRODUCT-001 が要求する v2 artifact ID grammar を spec として正式に記述し、既存 domain-first ID から namespace-aware ID への mapping rule を定義する。

## Work

- REQ-PRODUCT-001 の grammar 骨格（`<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<SEQUENCE>`）を spec section として詳細化する
- TASK ID の format（`<APP_NAMESPACE>-TASK-<DOMAIN_NAMESPACE>-<WORK_SEQUENCE>-<TASK_SEQUENCE>`）も仕様化する
- 既存 domain-first ID（例: `REQ-MCP-013`）から namespace-aware ID（例: `DRMCP-REQ-MCP-013`）への mapping rule を記述する
- ユーザーレビューに提出する

## Done condition

- v2 artifact ID grammar が spec section として存在する
- mapping rule が spec section として存在する
- ユーザーのレビュー承認を得ている

## Verification

- ユーザーの明示的なレビュー承認がある
- TASK-PRODUCT-001-04 の namespace catalog 作成に着手できる状態である

## Evidence
- v2 artifact ID grammar セクションを docs/spec/concepts/namespace-model/index.md に追加
- REQ/WORK/INV grammar、TASK grammar、ADR 例外、sequence format、mapping rule を定義
- PRODUCT prefix 特殊ケースと移行方针を明記
- domain namespace トークンは例示であり詳細は REQ-PRODUCT-002 で決定と追記
- ユーザーレビュー承認: 2026-06-07

# TASK-PRODUCT-001-02: 各 app の architecture sketch と domain namespace 候補リストアップ

- **id**: TASK-PRODUCT-001-02
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-PRODUCT-001
- **source_requirement**: REQ-PRODUCT-001
- **estimate**: 1d
- **depends_on**:
  - TASK-PRODUCT-001-01
- **outputs**:
  - 各 app の architecture sketch（mermaid）
  - domain namespace 候補一覧（spec draft）

## Goal

TASK-PRODUCT-001-01 で確定した app namespace 一覧をもとに、各 app のアーキテクチャを sketch し、domain namespace の候補を導出する。

## Work

- 確定した各 app namespace（DRMCP / BPDSL / PRODUCT 他）について、主要コンポーネントと責務を mermaid で図示する
- 各 app の figure から domain namespace 候補を導出し、説明とともにリストアップする
- spec draft としてまとめ、ユーザーレビューに提出する

## Done condition

- 全確定 app namespace に対して architecture sketch が存在する
- 各 app の domain namespace 候補が記述されている
- ユーザーのレビュー承認を得ている

## Verification

- ユーザーの明示的なレビュー承認がある
- TASK-PRODUCT-001-03 の v2 grammar spec 化に着手できる状態である

## Evidence
- spec draft 作成: docs/spec/concepts/namespace-model/index.md
- DRMCP / BPDSL / PRODUCT の architecture sketch (mermaid) を記載
- domain namespace catalog (canonical / non-canonical) を記載
- SELFHOST を canonical domain から外し non-canonical existing prefix テーブルに分離
- 現在の概念 spec と将来の機械可読 namespace registry の配置境界を明記
- ユーザーレビュー承認: 2026-06-07

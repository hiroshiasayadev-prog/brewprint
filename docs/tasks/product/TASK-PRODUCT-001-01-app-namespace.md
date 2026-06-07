# TASK-PRODUCT-001-01: app namespace 候補一覧の整理と確定

- **id**: TASK-PRODUCT-001-01
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-PRODUCT-001
- **source_requirement**: REQ-PRODUCT-001
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - 確定 app namespace 一覧（ユーザー承認済み）

## Goal

WORK-PRODUCT-001 が前提とする app namespace 一覧をユーザー判断のもとで確定する。

## Work

- REQ-PRODUCT-001 / ADR-095 / ADR-096 で言及されている app namespace 候補（DRMCP / BPDSL / PRODUCT）を整理する
- 曖昧な帰属が残る領域（self-hosting / DRUI / governance 等）の現状を調査し、独立 app namespace とするか DRMCP / PRODUCT の domain namespace とするかの候補を提示する
- ユーザーに候補一覧を提示し、確定を得る

## Done condition

- app namespace 一覧がユーザーによって確定されている
- 各 app namespace の簡潔な説明が記録されている

## Verification

- ユーザーの明示的な確認がある
- TASK-PRODUCT-001-02 の architecture sketch に着手できる状態である

## Evidence
- app namespace 一覧をユーザーが確定: DRMCP / BPDSL / PRODUCT
- DRUI は BPDSL 運用可能後まで保留
- SELFHOST は特定 app の domain ではなく cross-app 検証活動として整理
- 確定日: 2026-06-07（会話内ユーザー判断）

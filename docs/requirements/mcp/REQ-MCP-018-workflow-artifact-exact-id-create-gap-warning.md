# REQ-MCP-018: workflow artifact exact ID create gap warning

- **id**: REQ-MCP-018
- **status**: accepted
- **date**: 2026-06-03
- **source_refs**:
  - WORK-MCP-014
- **work_items**:
  - WORK-MCP-017

## Requirement

Workflow artifact create で exact ID を指定した場合、同一 kind/domain の連番に欠番を作る可能性があっても warning が出ない。

WORK / TASK は domain 内または parent scope 内で連番として運用するため、通常は `WORK-<DOMAIN>-new` や `TASK-<DOMAIN>-<WORK-SEQUENCE>-new` を使って server-side 採番に任せるのが安全である。Exact ID create は必要な場合に残すとしても、欠番を作る可能性は caller に明示されるべきである。

## Evidence

- `WORK-MCP-014` を exact ID で起票した結果、既存の `WORK-MCP-011` / `WORK-MCP-012` の次に `WORK-MCP-014` が作成され、`WORK-MCP-013` が欠番になった。
- MCP は exact ID を指定通り作成したが、`WORK-MCP-new` を使うべき状況であることや、前番号欠番の可能性を warning しなかった。

## Required Outcome

- Workflow artifact の exact ID create が同一 kind/domain/scope の連番欠番を作る可能性がある場合、proposal diagnostics または note で warning を返す。
- Warning は create を禁止しない。既存履歴や予約番号を尊重するため、exact ID create は引き続き可能とする。
- Diagnostic / note は、必要に応じて `*-new` placeholder の利用を促す。
- TASK の場合は parent work item scope の連番として判定する。

## Explicitly Excluded Scope

- 既存欠番の自動修復
- exact ID create の禁止
- REQ 番号と WORK 番号の一致強制
- 過去に作成済みの workflow artifact ID の rename / migration

## Boundary

- `propose_record_create` における workflow artifact exact ID create の diagnostics / authoring guidance を対象とする。

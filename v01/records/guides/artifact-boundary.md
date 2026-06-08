# Artifact Boundary Guide

## Abstract

brewprint の主要 artifact 間の責務境界をまとめる。起票・更新時は、対象 artifact の authoring guide とあわせてこの boundary guide を確認する。

## Migration Note

Maintainer note only. This is not an instruction to read another file, and it is not part of the public guide retrieval contract.

Extracted from:

- `docs/doc-policy.md`
- `docs/adr-authoring-guide.md`
- `docs/requirements/README.md`
- `docs/work-items/README.md`
- `docs/tasks/README.md`
- `docs/investigations/README.md`

## Core Rule

各 artifact は、自分が所有する一次情報だけを保持する。

他 artifact が所有する情報は、必要に応じて参照する。二重管理しない。

## Artifact Responsibility Matrix

| artifact | owns | does not own |
|---|---|---|
| spec | 現行仕様、schema、validation rule、diagnostic、MCP schema、renderer rule | ADR の長い判断背景、実装手順、task checklist |
| ADR | 設計判断、背景、理由、却下案、影響範囲、evidence | 現行仕様本文、作業 checklist、fixture migration 状態、実装手順 |
| requirement | 要求、不足、要望、spec gap 候補、stable requirement ID | 現行仕様、設計判断、内部設計、作業進捗、具体作業手順、完了状態 |
| work item | requirement 解消の到達点、作業フロー全体、横断進捗、影響範囲、task graph | 現行仕様本文、設計判断の長い理由、具体的な作業手順、個別 task status 正本 |
| task | 短期 concrete work、完了条件、個別 status、verification evidence | requirement 正本、設計判断、現行仕様、work item 全体の到達点、task graph |
| investigation | 調査結果、根拠、影響範囲の仮説、未確定点、選択肢、後続 artifact 候補 | 決定、現行仕様、要求そのもの、横断進捗、完了状態、具体的な作業手順 |
| internal design | spec を実装へ落とす internal wiring / route | 設計判断、要求、横断進捗、MVP semantic endpoint contract |
| UC / fixture | 実例、固定入力、期待出力、fixture-local 検証補助 | 汎用仕様、project-level trace relation、横断 gap / migration 進捗 |

## Spec vs ADR

Spec は現在の正しい仕様を所有する。

ADR は、なぜその判断をしたかを所有する。

ADR 本文の仕様記述は起票時点の snapshot であり、後続 ADR や spec 更新で覆されうる。

Spec と ADR が矛盾する場合、現行仕様としては spec を優先し、矛盾は docs stale / ADR conflict として報告する。

## Requirement / Work Item / Task

Requirement は「何が必要か」を所有する。

Work item は requirement を解消するための到達点、作業フロー、影響範囲、task graph を所有する。

Task は work item を前進させる短期 concrete work、完了条件、個別 status、verification evidence を所有する。

Work item 本文に、配下 task の完了 checkbox または status copy を手動複製しない。Task status の正本は各 task artifact である。

## Investigation Boundary

Investigation は、複雑な変更の調査材料を保存する。

Investigation は gate ではない。単純な変更や判断済みの内容では、investigation を必ず起票する必要はない。

Investigation は決定を所有しない。推奨案は書いてよいが、採用判断は後続 artifact に委ねる。

## Canonical Reference Boundary

Workflow artifact 間の relation は ID-as-ref を用いる。

- `REQ-*`
- `WORK-*`
- `TASK-*`

Design Records MCP が扱う record ID-as-ref は以下を含む。

- `ADR-*`
- `SPEC-*`
- `INV-*`
- `REQ-*`
- `WORK-*`
- `TASK-*`

Physical path は supported canonical relation として扱わない。

`internal-design:` / `coverage:` / `COV-*`、および `maps_to` / `covers` を用いた semantic realization relation は MVP に含めない。

## Startup Boundary

Startup docs は入口方針と安全ルールを持つ。

Artifact-specific authoring rules は、Design Records MCP の `get_authoring_guidance` で対象 artifact の guide ID を指定して読む。

Boundary 判断が必要な場合は、guide ID `artifact-boundary` を読む。

Startup docs に全 authoring rule を再掲しない。

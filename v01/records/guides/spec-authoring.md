# Spec Authoring Guide

## Abstract

Spec を作成・更新するときの実践ルールを定める。Spec は現行仕様の唯一の正を所有し、ADR の判断背景や具体実装手順を所有しない。

## Migration Note

This section is a maintainer note about phase-1 guide extraction history.
It is not an instruction to read the original files, and it is not part of the public authoring guidance retrieval contract.

Extracted from:

- `docs/spec-authoring-guide.md`
- `docs/doc-policy.md`

## Purpose

Spec は brewprint の現在の正しい仕様を記録する artifact である。

現行仕様を知りたいときは spec を参照する。なぜそう決まったかを知りたいときは ADR を参照する。

## Spec Owns

Spec が所有するもの:

- 現在の schema
- field definition
- validation rule
- diagnostic code
- semantic reference declaration
- MCP request / response schema
- renderer rule
- current behavior contract

## Spec Does Not Own

Spec が所有しないもの:

- 設計判断の長い背景説明
- 却下した代替案の詳細
- 実装手順
- task checklist
- investigation の探索ログ
- fixture migration の作業状態

## ADR Relationship

ADR は判断の履歴であり、spec は現在の仕様である。

ADR 本文の仕様記述は起票時点の snapshot である。後続 ADR や spec 更新で覆されうる。

Spec と ADR が矛盾する場合、現行仕様としては spec を優先し、矛盾は docs stale / ADR conflict として報告する。

## Update Rules

- Spec 更新時は、関連する現行仕様本文を直接更新する。
- 由来が重要な場合は ADR / requirement / work item への短い由来注記を置く。
- Spec に作業状態やチェックリストを置かない。
- Spec 内の semantic ref declaration は、resolver / validation contract と矛盾しないようにする。
- Front matter や metadata がある spec では、本文更新と合わせて必要な metadata も更新する。

## Reference Style

Spec から他 artifact を参照する場合は、可能な限り canonical ID-as-ref または semantic ref を使う。

Physical path は所在説明には使えるが、canonical relation として扱わない。

## Review Points

Spec 更新では以下を見る。

- 現行仕様として自己完結しているか
- ADR の snapshot 記述を二重管理していないか
- validation / renderer / MCP schema と矛盾していないか
- field / status / diagnostic の語彙が曖昧でないか
- future work を現行仕様として書いていないか

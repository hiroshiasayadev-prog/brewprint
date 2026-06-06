# 094: workflow artifact status vocabulary alignment

- **status**: accepted
- **date**: 2026-06-06
- **depends_on**: ADR-092, ADR-089
- **supersedes**: 
- **migrated_to_spec**: 

## 背景

Design Records MCP の workflow artifact kind 間で status vocabulary が統一されていない。

`task` は `todo` / `doing` / `blocked` / `done` というカジュアルな短形トークンを使い、`work_item` は `not_started` / `decision_pending` / `design_spec_pending` / `internal_design_pending` / `yaml_pending` / `implementation_pending` / `fixture_pending` / `verification_pending` / `done` / `blocked` という brewprint 固有のステージ特化トークン10種を使う。

REQ-MCP-026 の evidence として記録されている通り、LLM callers は `task` を操作する際に `in_progress` や `not_started` という自然言語の first prior を使い、`invalid_status_for_kind` で繰り返し失敗した。`todo` や `doing` は LLM が workflow 文脈で最初に想起するトークンではない。

また `work_item` のステージ特化トークンは brewprint の内部作業フェーズをステータスに押し込んだ設計であり、Linear / GitHub Projects などの標準的な workflow ツールが採用するライフサイクルフェーズ vocabulary からかけ離れている。これらのステージ情報は work item 本文の Task flow セクションや配下 task の状態から読み取れるため、status として保持する必要がない。

## 決定

`task` と `work_item` の両 kind の status vocabulary を以下の4トークンに統一する。

```
not_started / in_progress / blocked / done
```

| token | meaning |
|---|---|
| `not_started` | 未着手 |
| `in_progress` | 作業中 |
| `blocked` | 外部要因・依存判断待ちで停止中 |
| `done` | 完了 |

`decision` / `investigation` / `requirement` / `spec` の status vocabulary はこの ADR では変更しない。

既存レコードの migration mapping:

| 変更前 | 変更後 | 対象 kind |
|---|---|---|
| `todo` | `not_started` | `task` |
| `doing` | `in_progress` | `task` |
| `decision_pending` | `blocked` | `work_item` |
| `design_spec_pending` | `in_progress` | `work_item` |
| `internal_design_pending` | `in_progress` | `work_item` |
| `yaml_pending` | `in_progress` | `work_item` |
| `implementation_pending` | `in_progress` | `work_item` |
| `fixture_pending` | `in_progress` | `work_item` |
| `verification_pending` | `in_progress` | `work_item` |
| `not_started` | `not_started` | `work_item`（変更なし） |
| `blocked` | `blocked` | 両 kind（変更なし） |
| `done` | `done` | 両 kind（変更なし） |

## 理由

Linear / GitHub Projects / Jira など標準的な workflow ツールはライフサイクルフェーズ（未着手・作業中・完了）をステータスとして持ち、サブステージはラベル・本文・タスク構造で表現する。この設計により、ステータスは recall しやすい少数の語彙に収まる。

`not_started` と `in_progress` は LLM が workflow 文脈で最も高確率に想起するトークンであり、trial-and-error なしに正しいステータスを使える。

`decision_pending` は「外部の設計判断が出るまで進めない」状態であり、これは `blocked` の典型的な意味と一致する。他の `*_pending` ステージトークンは「現在そのフェーズの作業をしている」状態であり、`in_progress` に統合できる。ステージの詳細は work item 本文の Task flow セクションと配下 task が担う。

`task` の `todo` / `doing` は、現行の `work_item.not_started` / `in_progress` と意味的に等価だが表記が異なるため、同一ライフサイクルフェーズに対して異なるトークンを LLM が学習しなければならない。統一することでこの認知負荷を除去する。

## 却下した代替案

**全 kind を4トークンに統一する案**: `decision`（ADR）の `superseded` と `requirement` の `rejected` / `deferred` は `done` では表現できない意味を持つ。`superseded` は「かつて accepted だったが別 ADR に置き換えられた」、`rejected` は「実施しないと決めた」を意味し、`done`（完了）と混同すると記録の信頼性を損なう。スコープを `task` / `work_item` に限定する。

**`work_item` ステージトークンを維持する案**: ステージ情報のステータス保持は standard workflow tool の設計に反し、LLM の recall likelihood を下げる。ステージ情報は本文セクションと task 構造で代替できる。

**`task` のみ変更する案（Option A）**: `work_item` も同様の recall 問題を持ち、standard vocabulary との乖離は同程度である。両 kind を同時に統一する方がスコープ・migration コストの比率として優れる。

## 影響

- `docs/spec/design-records-mcp/schema.md` § status テーブル: `task` / `work_item` の行を更新する
- `internal/designrecords/types.go`: `RecordStatusTodo` / `RecordStatusDoing` を削除し、`RecordStatusInProgress` を追加する。ステージ特化定数を削除する
- `internal/designrecords/validation.go`: `isAllowedStatusForKind` の `task` / `work_item` 分岐を更新する
- `docs/guides/task-authoring.md` / `docs/guides/work-item-authoring.md`: status テーブルを更新する
- 既存 `task` / `work_item` レコード: migration mapping に従い status 値を更新する
- REQ-MCP-023 (synonym repair): `in_progress` → `doing` の変換は不要になる。026 完了後に 023 scope を再評価する

## Evidence

- REQ-MCP-026: evidence セクション（`in_progress` / `not_started` での失敗記録）
- TASK-MCP-022-01: vocabulary alignment analysis（migration cost 実測、options matrix）
- commit: tbd
- impl commit: tbd

# V01-WORK-PRODUCT-002: Subdomain grouping model の仕様化と propose ツール write-time advisory 実装

- **id**: V01-WORK-PRODUCT-002
- **status**: done
- **date**: 2026-06-07
- **source_requirement**: V01-REQ-PRODUCT-002
- **impact_refs**:
- **tasks**:

## Goal

V01-REQ-PRODUCT-002 が定義した subdomain grouping model を spec に記述し、propose 系ツールの write-time advisory として実装する。

## Boundary

このWORKが所有するもの:

- namespace-model spec への subdomain model 記述追加
- design-records-mcp schema spec への `subdomain:` フィールド定義追加
- propose 系ツール（propose_record_create / propose_record_update）への write-time advisory 実装

このWORKが所有しないもの:

- 既存 records への `subdomain:` フィールドの一括付与
- subdomain を軸にした検索・フィルタリング API の追加
- UI での subdomain 表示

## Task flow

```mermaid
flowchart TD
  T01[TASK-01: namespace-model spec\nsubdomain 記述追加] -->|spec review ゲート| T02[TASK-02: design-records-mcp schema spec\nsubdomain フィールド追加]
  T02 -->|spec review ゲート| T03[TASK-03: propose ツール\nwrite-time advisory 実装]
  T03 --> DONE[完了]
```

## Task Candidates

- TASK-01: `docs/spec/concepts/namespace-model/index.md` に subdomain model を追記（subdomain の定義・label 形式・動的導出・write-time advisory の概要）→ spec review ゲート
- TASK-02: `SPEC-design-records-mcp-schema` に `subdomain:` フィールドを追加（任意フィールド、文字列型）→ spec review ゲート
- TASK-03: propose_record_create / propose_record_update が `subdomain` 新規値を検出した際に同 domain 内の既存値を列挙して返す（block なし）

## Completion Condition

- namespace-model spec に subdomain model が記述されている
- design-records-mcp schema spec に `subdomain:` フィールドが定義されている
- propose 系ツールが新規 subdomain 値を検出した場合に既存値を列挙する advisory を返す
- V01-REQ-PRODUCT-002 に本 WORK が反映されている

## Evidence

TASK-01（namespace-model spec 更新）、TASK-02（schema spec 更新）、TASK-03（Go 実装）をすべて完了し、Completion Condition をすべて満たすことを確認した。

| タスク | 成果物 | commit |
|---|---|---|
| TASK-01: namespace-model spec に `## Subdomain model` セクション追加 | `docs/spec/concepts/namespace-model/index.md` | c26de9a |
| TASK-02: schema spec に `subdomain` フィールド定義追加（record source table / workflow artifact bullet metadata / field definitions） | `docs/spec/design-records-mcp/schema.md` | 54ac722 |
| TASK-03: Go 実装（`types.go` Subdomain フィールド・`DiagnosticNewSubdomainValue` 追加、`parser.go` subdomain 解析追加、`authoring.go` write-time advisory 実装） | `internal/designrecords/` 配下 3 ファイル | 35888cb |

V01-REQ-PRODUCT-002 は `accepted` に更新済み。spec review ゲートを 2 回経由（TASK-01・TASK-02 後）し、各レビュー指摘（Record model top-level 行削除・sections frontmatter 修正・Record source table subdomain 追加・bullet metadata 形式修正）を反映してから次のタスクに進んだ。

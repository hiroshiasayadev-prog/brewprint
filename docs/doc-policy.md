# brewprint ドキュメント運用方針

> このdocは、AIアシスタントとの協働におけるドキュメント管理の入口方針を定める。
> 別会話のAIアシスタントは、このdocを最初に読むこと。

---

## 1. 最重要ルール

### ファイル操作

AIアシスタントのサンドボックス環境（bash_tool / create_file 等）は揮発性であり、セッション終了と同時に消える。
**ファイル操作はすべてMCPツール経由でローカルリポジトリに対して行うこと。例外なし。**

- 書き込み操作は、ユーザーが明示的に指示した場合のみ実行する。
- 書き込み指示がない場合は、変更案または dry-run diff を提示して許可を得る。
- 指示されていないファイルを勝手に変更しない。

### 読み込み範囲

- docs に存在する可能性がある情報は、推測せず先に読む。
- **全 doc を最初から読まない。**
- 現行仕様を把握したいときは spec を読む。
- 判断理由を辿りたいときだけ ADR を読む。
- 作業に関連する requirement / work item / task / spec / ADR / investigation / internal-design / UC / YAML だけ読む。
- closed な旧 M-series 記録の詳細は原則読まない。
- 長い Markdown は headings / section read / head / tail を使い、必要箇所だけ読む。
- 根拠が足りない場合のみ全文を読む。

---

## 2. プロジェクト概要

brewprint は **人間と LLM の共通設計言語**。

- 人間向け: Mermaid 図などの readable view。
- LLM 向け: signature / dep tree / inspect / design record retrieval などの MCP interface。
- brewprint DSL YAML: 対象 system / design model を表す primary implementation source。
- trace metadata YAML: semantic trace 運用のための metadata。brewprint DSL YAML とは別責務。

実装は Go で行う。brewprint DSL YAML の AST を Go で保持し、MCP tools として LLM に公開する。

---

## 3. ドキュメント運用方針

ドキュメント運用は **spec-first** で行う（ADR-050）。

`docs/doc-policy.md` は入口方針と安全ルールだけを所有する。
Artifact-specific な書式・起票手順・責務境界はここに再掲しない。

### Authoring guidance

起票・更新時は、Design Records MCP の authoring guidance tools を使う。

- `list_authoring_guides`: 利用可能な guide ID / title / abstract を確認する。
- `get_authoring_guidance`: 必要な guide ID を指定して Markdown 全文を読む。

Guide source path は public contract ではない。Guide ID を public reference として扱う。

| 対象 | guide ID |
|---|---|
| spec | `spec-authoring` |
| ADR | `adr-authoring` |
| requirement | `requirement-authoring` |
| work item | `work-item-authoring` |
| task | `task-authoring` |
| investigation | `investigation-authoring` |
| artifact responsibility / boundary 判断 | `artifact-boundary` |

`internal-design` と UC / fixture の authoring guide は phase 1 では未整備である。必要な場合は関連 README / task / spec を読む。

### Responsibility boundary

Artifact 間の関係、所有範囲、非所有範囲、workflow artifact の境界、spec / ADR / investigation の境界判断は `artifact-boundary` guide が所有する。

`docs/doc-policy.md` に boundary detail を再掲しない。

---

## 4. Design Records MCP

Design Records MCP は、ADR / spec / investigation / requirement / work item / task の record 検索・検証・取得、および authoring guide の discovery / retrieval に利用する。

主な入口:

- `list_records`: record 候補を探す。
- `get_record`: 単一 record を取得する。
- `get_records`: 複数 record をまとめて取得する。
- `resolve_reference`: canonical reference を解決する。
- `validate_records`: record metadata / relation を検証する。
- `list_authoring_guides`: authoring guide catalog を取得する。
- `get_authoring_guidance`: guide ID で authoring guide Markdown 全文を取得する。

Authoring guidance は Design Records record kind ではない。Guide ID は resolver target / record ID ではない。
Guide source path は public response contract に含めない。

Tool の request / response 仕様は `docs/spec/design-records-mcp/tools.md` を参照する。

---

## 5. セッション開始時の手順

1. `docs/doc-policy.md` を読む。
2. 作業対象が不明な場合は、Design Records MCP で record 一覧や必要な guide 一覧を確認する。
3. 作業対象が決まったら、関連する requirement / work item / task / spec / ADR / investigation だけを読む。
4. 起票・更新を行う場合は、対象 artifact の guide ID を `get_authoring_guidance` で取得して読む。
5. 責務境界に迷う場合は `artifact-boundary` を読む。
6. 旧 M-series 記録の確認が必要な場合のみ、`docs/TASKS.md` または `docs/tasks/mXX-*.md` を legacy record として読む。

---

## 6. Legacy / archive notes

- `docs/TASKS.md` は旧 M-series index として migration まで維持する。
- `docs/tasks/mXX-*.md` は legacy milestone-shaped work record として扱う。
- 新規 work は requirement から work item を起票し、具体作業を `docs/tasks/<domain>/TASK-*.md` の短期 task に分解する。
- 新形式に別途 milestone artifact は作らない。

---

## 7. v1で確定した運用方針

ADR-057 で brewprint v1.0.0-spec を凍結した際に、以下の運用が確定した。

### Release snapshots 運用

- 仕様 + 実装スナップショットの git tag は `v{MAJOR}.{MINOR}.{PATCH}-spec` 形式。
- 凍結対象: `docs/adr/*` / `docs/spec/**` / `docs/uc/**` / Go 実装ツリー全体。
- 運用頻度: メジャーな仕様マイルストーンに合わせて切る。毎 milestone では切らない。
- 公開 contract の version は別軸で管理する。

詳細は ADR-057 を参照する。

### DISCLAIMER.md

- プロジェクト root に `DISCLAIMER.md` を新設する方針が確定（ADR-057）。
- 文面はユーザーが起草する。AI アシスタントは起案しない。
- 法務的主張（業務時間外開発、会社リソース不使用、公知技術の組合せ）を記載する。

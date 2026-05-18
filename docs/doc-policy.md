# brewprint ドキュメント運用方針

> このdocは、AIアシスタントとの協働におけるドキュメント管理の入口方針を定める。
> 別会話のAIアシスタントは、このdocを最初に読むこと。

---

## 1. 最重要ルール

### ファイル操作

AIアシスタントのサンドボックス環境（bash_tool / create_file 等）は揮発性であり、セッション終了と同時に消える。
**ファイル操作はすべてMCPツール経由でローカルリポジトリに対して行うこと。例外なし。**

- 書き込み操作は、ユーザーが明示的に指示した場合のみ実行する。
- 書き込み指示がない場合は、変更案またはdry-run diffを提示して許可を得る。
- 指示されていないファイルを勝手に変更しない。

### 読み込み範囲

- docsに存在する可能性がある情報は、推測せず先に読む。
- **全docを最初から読まなくていい。** ADRタイトルと `docs/TASKS.md` で文脈を把握し、必要なものだけ読む。
- **現行仕様を把握したいときはspecを読む。** ADRは根拠を辿りたいときに参照する。
- 作業に関連するspec / ADR / UC / YAMLだけ読む。
- closed milestone の詳細taskは原則読まない。
- 長いMarkdownは head / headings / section read を使い、必要箇所だけ読む。
- 根拠が足りない場合のみ全文を読む。

---

## 2. プロジェクト概要

brewprintは**人間とLLMの共通設計言語**。

- 人間向け → Mermaid図（md形式でrender）
- LLM向け → signature / dep tree / inspect（MCP経由）
- YAMLはその裏側にある中間表現

実装はGoで行う。YAMLのASTをGoで保持し、MCPツールとしてLLMに公開する。

---

## 3. ドキュメントの基本方針

ドキュメント運用は **spec-first** で行う（ADR-050）。

| 文書 | 役割 | 作成・更新ルール |
|---|---|---|
| `docs/spec/` | 現行仕様の唯一の正 | `docs/spec-authoring-guide.md` |
| `docs/adr/` | 設計判断の履歴 | `docs/adr-authoring-guide.md` |
| `docs/uc/` | golden fixture corpus（入力YAMLと期待出力） | 各fixtureのREADME / task file |
| `docs/tasks/` | 作業項目、順序、完了条件 | milestone別task file |
| `docs/impl/` | 実装作業の引継ぎ・レビューメモ | 必要時のみ参照 |

### spec と ADR の関係

- 現行仕様を知りたいときは **spec** を参照する。
- なぜそう決まったかを知りたいときは **ADR** を参照する。
- ADR本文の仕様記述は起票時点のスナップショットであり、後続ADRやspec更新で覆されうる。
- specとADRが矛盾する場合、現行仕様としてはspecを優先し、矛盾は docs stale / ADR conflict として報告する。

---

## 4. ドキュメント構成

```
docs/
  doc-policy.md              ← このファイル。セッション開始時に読む入口方針
  adr-authoring-guide.md     ← ADRの起票・更新・フォーマット
  spec-authoring-guide.md    ← specの作成・更新・フォーマット
  spec/                      ← 現行仕様の唯一の正
  adr/                       ← Architecture Decision Records
  uc/                        ← golden fixture corpus
  impl/                      ← 実装作業の引継ぎ・レビューメモ
  TASKS.md                   ← milestone index
  tasks/                     ← milestone別task file
    mXX-*.md
```

`docs/TASKS.md` は milestone index として運用する。
詳細taskは `docs/tasks/mXX-*.md` に置き、セッション開始時は `docs/TASKS.md` のみ読む。
作業対象 milestone が決まってから該当 task file を読む。

### 現状のspec構成

```
docs/spec/
  overview.md
  project-layout.md
  file-types.md
  naming.md
  nodes.md
  edges.md
  diagnostics.md
  mcp.md
  views/
    dag.md
    er.md
    state-diagram.md
    sequence-diagram.md
    api-table.md
    wireframe.md
```

---

## 5. セッション開始時の手順

1. `docs/doc-policy.md`（このdoc）を読む。
2. `docs/TASKS.md` を milestone index として読む。詳細task fileはまだ読まない。
3. `docs/adr/` の一覧を取得し、acceptedなADRのタイトルを確認する。
4. 作業に関連するspec / ADR / UC / YAMLだけ読む。
5. 作業対象 milestone が決まってから、該当する `docs/tasks/mXX-*.md` を読む。

---

## 6. 既存ADRの扱い

ADRは、設計判断の履歴として扱う。
現行仕様を確認する目的で、過去ADRを網羅的に読む必要はない。

ADRを読むのは、以下の場合に限る。

- specの由来や判断理由を確認したいとき
- 既存の設計判断を変更・上書きする可能性があるとき
- specとADRの矛盾が疑われるとき
- ADR自体を更新・移行するとき

ADR本文の仕様記述がspecと矛盾する場合、現行仕様としてはspecを優先し、矛盾は docs stale / ADR conflict として報告する。

---

## 7. Design Records MCP

ADR / spec の record 検索・検証・取得には Design Records MCP を利用できる。
tool の request / response 仕様は `docs/spec/design-records-mcp/tools.md` を参照する。

---

## 8. v1で確定した運用方針

ADR-057でbrewprint v1.0.0-spec を凍結した際に、以下の運用が確定した。

### Release snapshots運用

- 仕様+実装スナップショットのgitタグは `v{MAJOR}.{MINOR}.{PATCH}-spec` 形式。
- 凍結対象: `docs/adr/*` / `docs/spec/**` / `docs/uc/**` / Go実装ツリー全体。
- 運用頻度: メジャーな仕様マイルストーンに合わせて切る。毎milestoneでは切らない。
- 公開contractのバージョン（MCP vNなど）は別軸で管理する。

詳細は [ADR-057 §4](adr/057-brewprint-v1-snapshot.md) を参照する。

### DISCLAIMER.md

- プロジェクトルートに `DISCLAIMER.md` を新設する方針が確定（ADR-057 §5）。
- 文面はユーザーが起草する。AIアシスタントは起案しない。
- 法務的主張（業務時間外開発、会社リソース不使用、公知技術の組合せ）を記載する。



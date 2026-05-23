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
- brewprint DSL YAML → 対象 system / design model を表す primary implementation source
- trace metadata YAML → semantic trace 運用のための metadata。brewprint DSL YAML とは別責務

実装はGoで行う。brewprint DSL YAML のASTをGoで保持し、MCPツールとしてLLMに公開する。
semantic trace MVP では、physical path ではなく canonical reference を用い、active semantic ref は `spec:` のみに限定する。
MVP は record ID-as-ref と investigation の canonical reference resolve / validation foundation を提供し、`internal-design:` / `coverage:` / `COV-*` / semantic realization relation は concrete requirement が生じるまで扱わない。

---

## 3. ドキュメントの基本方針

ドキュメント運用は **spec-first** で行う（ADR-050）。

| 文書 | 役割 | 作成・更新ルール |
|---|---|---|
| `docs/spec/` | 現行仕様の唯一の正 | `docs/spec-authoring-guide.md` |
| `docs/adr/` | 設計判断の履歴 | `docs/adr-authoring-guide.md` |
| `docs/requirements/` | 要求・不足・要望と stable requirement ID | `docs/requirements/README.md` |
| `docs/work-items/` | requirement から派生する横断作業・進捗・影響範囲 | `docs/work-items/README.md` |
| `docs/internal-design/` | 現行仕様を実装へ落とす internal wiring / route | `docs/internal-design/README.md` |
| `docs/investigations/` | 調査結果・根拠・影響範囲・選択肢の保存 | `docs/investigations/README.md` |
| `docs/uc/` | golden fixture corpus（入力brewprint DSL YAMLと期待出力） | 各fixtureのREADME / task file |
| `docs/tasks/` | 作業項目、順序、完了条件 | milestone別task file |
| `docs/impl/` | 実装作業の引継ぎ・レビューメモ | 必要時のみ参照 |

### spec と ADR の関係

- 現行仕様を知りたいときは **spec** を参照する。
- なぜそう決まったかを知りたいときは **ADR** を参照する。
- ADR本文の仕様記述は起票時点のスナップショットであり、後続ADRやspec更新で覆されうる。
- specとADRが矛盾する場合、現行仕様としてはspecを優先し、矛盾は docs stale / ADR conflict として報告する。

### semantic trace artifact の扱い

`docs/requirements/` は、要求・不足・要望を stable requirement ID で保持する。
`docs/work-items/` は、source requirement を持つ横断作業・進捗・影響範囲を保持する。
`docs/internal-design/` は、spec を実装へ落とす internal wiring / route を保持する。ただし MVP semantic trace endpoint ではない。
External relation / assurance artifact は、gap / evidence / sign-off / lifecycle 等を中央管理する concrete requirement が生じた場合に、配置と責務を含めて新設判断する。MVP では directory や authoring entrance を設けない。

semantic trace MVP の active semantic prefix は `spec:` のみとし、Design Records MCP が扱う `ADR-*` / `SPEC-*` / `INV-*` と investigation の canonical references を resolve / validate 対象とする。
`internal-design:` / `coverage:` / `COV-*`、および `maps_to` / `covers` を用いた semantic realization relation は MVP に含めない。
`yaml:` は brewprint DSL YAML 用の reserved prefix であり、MVP では active trace 対象にしない。
fixture / golden は project-level semantic trace foundation の対象外とする。

### investigations の扱い

`docs/investigations/` は、複雑な変更における調査結果、根拠、影響範囲、未確定点、選択肢、後続 artifact 候補を保存する。

investigation は、requirement / work item / task / ADR / spec / internal design / 将来の external relation artifact / 別 investigation の起票・更新前に必ず必要な gate ではない。

investigation は、決定、現行仕様、要求そのもの、横断進捗、完了状態、具体的な作業手順を所有しない。

investigation の directory / ID / metadata / status / lifecycle / authoring format は `docs/investigations/README.md` が所有する。

---

## 4. ドキュメント構成

```
docs/
  doc-policy.md              ← このファイル。セッション開始時に読む入口方針
  adr-authoring-guide.md     ← ADRの起票・更新・フォーマット
  spec-authoring-guide.md    ← specの作成・更新・フォーマット
  spec/                      ← 現行仕様の唯一の正
  adr/                       ← Architecture Decision Records
  requirements/              ← 要求・不足・要望
  work-items/                ← requirement由来の横断作業・進捗
  internal-design/           ← internal wiring / route
  investigations/            ← 調査artifact
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
  design-records-mcp/
    overview.md
    schema.md
    tools.md
  concepts/
    project-artifact-model/
    traceability/
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
4. 作業に関連するspec / ADR / investigation / UC / YAMLだけ読む。
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

現行実装では、ADR / spec の record 検索・検証・取得に Design Records MCP を利用できる。
ADR-087 により、Design Records MCP は investigation record integration、docs artifact 間の semantic/artifact ref resolve、investigation の `source_refs` / 記載済み `follow_up_results` の参照切れ検査も担う方針が確定している。
ADR-088 により、MVP resolve / validation は `spec:` semantic ref、record ID-as-ref、investigation canonical references に限定され、internal-design / coverage relation resolution は必須範囲から外れた。
これら ADR-087 / ADR-088 追従の実装と contract 切替は M19 で追跡する。M19 完了までは、investigation / semantic ref resolve が現行 tool で利用可能であると仮定しない。
resolver が読む artifact と record kind として公開する artifact は同一集合である必要はない。
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



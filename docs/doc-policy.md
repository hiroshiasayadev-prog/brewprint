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
- 作業に関連する requirement / work item / task / spec / ADR / investigation / internal-design / UC / YAML だけ読む。
- closed な旧 M-series 記録の詳細は原則読まない。
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
| `docs/work-items/` | requirement を解消する到達点・作業フロー全体・横断進捗・影響範囲・task graph | `docs/work-items/README.md` |
| `docs/internal-design/` | 現行仕様を実装へ落とす internal wiring / route | `docs/internal-design/README.md` |
| `docs/investigations/` | 調査結果・根拠・影響範囲・選択肢の保存 | `docs/investigations/README.md` |
| `docs/uc/` | golden fixture corpus（入力brewprint DSL YAMLと期待出力） | 各fixtureのREADME / task file |
| `docs/tasks/` | 短期 concrete work・完了条件・個別 status・verification evidence | `docs/tasks/README.md` |
| `docs/TASKS.md` / `docs/tasks/m*.md` | 旧 M-series の index / legacy milestone-shaped work record | migration までの参照用。新しい artifact layer ではない |
| `docs/impl/` | 実装作業の引継ぎ・レビューメモ | 必要時のみ参照 |

### spec と ADR の関係

- 現行仕様を知りたいときは **spec** を参照する。
- なぜそう決まったかを知りたいときは **ADR** を参照する。
- ADR本文の仕様記述は起票時点のスナップショットであり、後続ADRやspec更新で覆されうる。
- specとADRが矛盾する場合、現行仕様としてはspecを優先し、矛盾は docs stale / ADR conflict として報告する。

### semantic trace artifact の扱い

`docs/requirements/` は、要求・不足・要望を stable requirement ID で保持する。
`docs/work-items/` は、source requirement を持つ requirement 解消の到達点・作業フロー全体・横断進捗・影響範囲・task graph を保持する。従来 `milestone` と呼んでいた実行計画の役割は、新形式では work item が担う。
`docs/tasks/` は、work item 配下の短期 concrete work と個別 status / evidence を保持する。Task status の正本は各 task artifact とし、work item に checkbox として手動複製しない。
`milestone` を新しい artifact layer、canonical identity、metadata field、または work item 間の relation として導入しない。既存 `docs/tasks/m*.md` は移行まで legacy milestone-shaped work record として扱う。
Workflow artifact 間の canonical relation は `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref を用い、physical path は supported canonical relation に含めない。
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
  work-items/                ← requirement解消の到達点・作業フロー・横断進捗・task graph
  internal-design/           ← internal wiring / route
  investigations/            ← 調査artifact
  uc/                        ← golden fixture corpus
  impl/                      ← 実装作業の引継ぎ・レビューメモ
  TASKS.md                   ← 旧 M-series index（legacy migration まで掲載）
  tasks/                     ← 短期 concrete task / legacy milestone-shaped work record
    README.md                ← 新規 task のauthoring guidance
    <domain>/TASK-*.md       ← 新形式の短期 task
    mXX-*.md                 ← 既存 legacy milestone-shaped work record（移行対象）
```

`docs/TASKS.md` は旧 M-series の index として migration まで維持する。既存の `docs/tasks/mXX-*.md` は、work item 相当の計画を保持した legacy milestone-shaped work record として参照する。
新規 work は requirement から work item を起票し、具体作業を `docs/tasks/<domain>/TASK-*.md` の短期 task に分解する。新形式に別途 milestone artifact は作らない。
セッション開始時は `docs/TASKS.md` のみ読み、作業対象が決まってから関連する work item / task / legacy milestone record を読む。

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
2. `docs/TASKS.md` を旧 M-series index として読む。legacy milestone-shaped work record や個別 task file はまだ読まない。
3. `docs/adr/` の一覧を取得し、acceptedなADRのタイトルを確認する。
4. 作業に関連する requirement / work item / task / spec / ADR / investigation / internal-design / UC / YAML だけ読む。
5. 新形式の作業対象が決まったら関連する work item と `docs/tasks/<domain>/TASK-*.md` を読む。旧 M-series 記録の確認が必要な場合のみ `docs/tasks/mXX-*.md` を legacy record として読む。

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

現行実装では、ADR / spec / investigation の record 検索・検証・取得に Design Records MCP を利用できる。
ADR-087 / ADR-088 に基づく investigation integration と canonical reference resolve / validation は M19 で実装追従済みである。
ADR-090 に基づき、明示 ID 配列による複数 record detail retrieval の `get_records` も利用できる。
現行 resolve / validation は `spec:` semantic ref、`ADR-*` / `SPEC-*` / `INV-*` ID-as-ref、investigation canonical references を対象とし、internal-design / coverage relation resolution は必須範囲外である。
ADR-091 が定めた workflow artifact 間の `REQ-*` / `WORK-*` / `TASK-*` ID-as-ref の MCP support は REQ-MCP-003 の後続判断対象であり、現行 tool で利用可能であると仮定しない。Physical path は supported canonical relation に含めない。
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



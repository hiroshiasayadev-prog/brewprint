# 077: Design Records MCP MVP boundary and tool prioritization

- **status**: accepted
- **date**: 2026-05-11
- **depends_on**: ADR-076
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

ADR-076では、ADR/spec運用を machine-readable metadata と MCP query / validation で支援するため、Design Records MCP MVPを導入する方針を採用した。

ADR-076は、Design Records MCPの完全なschemaやtool仕様を決めず、MVPを最小範囲に留める境界を定めた。
また、Design Records MCPは既存brewprint MCPの `QueryService` へそのまま混ぜず、既存brewprint MCPとは独立して起動・検証できる構成を第一候補とした。

一方で、Design Records MCP MVPでどのtoolを優先実装するか、どこまでをMVPのread-only機能に含めるかは、ADR-076では詳細に決めていない。

過去のgeneric support調査では、関連ADR候補を絞るために複数ADR本文を読む必要があった。
その大半は本文理解ではなく、`status` / `supersedes` / `migrated_to_spec` などの確認に使われた。
この観測から、Design Records MCP MVPでは、本文全文取得そのものよりも、本文を読む前にrecordを絞り込むためのindex queryとvalidationを優先する必要がある。

また、ADR番号から本文全文やpathを取得できない場合、結局filesystem toolへ戻る必要がある。
そのため、Design Records MCPは汎用filesystemの代替にはしない一方で、design recordとして意味を持つ読み取り、path解決、本文取得、次ADR番号の提案までは扱う必要がある。

## 決定

Design Records MCP MVPのtool優先度と責務境界を以下のように定める。

本ADRは ADR-076 のMVP tool境界を refine する。
ADR-076ではMVP toolを `list_records` / `get_record` / `validate_records` に限定したが、本ADRではこれらをP0必須toolと位置づけ直し、`suggest_next_record` をP1の任意補助toolとして許容する。

### P0: MVP必須tool

MVPでは以下のread-only toolを必須とする。

- `list_records`
- `get_record`
- `validate_records`

P0 toolの目的は、ADR/spec本文を読む前に、record indexと基本metadataによって読むべきrecordを絞り込み、metadataの不整合を検出することである。

### P1: MVPに含めてもよい補助tool

以下はMVPに含めてもよいが、P0完了後に扱う。

- `suggest_next_record`

`suggest_next_record` は、既存record indexから次のADR番号と推奨pathを提案するread-only toolである。
実際のファイル作成は行わない。

### MVP外

以下はMVPには含めない。

- `trace_record`
- `list_gaps`
- `create_record`
- `update_record`
- `set_evidence`
- その他のwrite系tool

`trace_record` と `list_gaps` は、MVPのvalidatorを実データに当てて必要なgraph traversalやgap分類の形が見えてから、後続ADRまたはspecで追加を判断する。

write系toolは、ユーザー確認、dry-run diff、衝突処理、template責務、commit運用などを伴うため、read-only MVPからは外す。

### `list_records` の責務

`list_records` は、ADR/spec record indexを構造化して返すquery toolである。
単なる全件一覧ではなく、本文を読む前の候補絞り込みを主目的とする。

MVPでは、少なくとも以下の情報を返す。

- `id`
- `kind`
- `title`
- `status`
- `path`
- `depends_on`
- `supersedes`
- `migrated_to_spec`

`title` は Markdown本文のH1から抽出する。
ADRでは `# NNN: タイトル` 形式を前提とし、MVPでは `design_record` front matter に `title` フィールドを持たせない。
filenameからのtitle推定はMVPでは行わない。

MVPでは、少なくとも以下の絞り込みを扱う。

- `kind`
- `status`
- `id`
- `id_range`

新規ADR起票後の確認や直近ADRの把握をfilesystem listingへ戻さず行えるように、`list_records` はID順の絞り込み取得をサポートする。

MVPでは、並び順の基準はrecord IDに限定し、mtimeやgit履歴による並び替えは扱わない。

MVPでは、以下の形を採用する。

- `order_by: id`
- `order: asc | desc`
- `limit: N`

`head` / `tail` はMVPでは採用しない。
これは、filesystemの行単位head/tailとの混同を避け、将来 `order_by` の値を拡張できる形にするためである。

### `get_record` の責務

`get_record` は、record IDからmetadata、path、見出し、必要に応じてMarkdown本文を取得するread toolである。

ADR番号から本文全文やpathを取得できないと、format確認や編集前確認のたびにfilesystem toolへ戻る必要がある。
そのため、`get_record` はMVP必須toolとする。

MVPでは、`include_body` を持つ。

- `include_body: false` の場合、metadata、path、title、headingsなどを返す
- `include_body: true` の場合、上記に加えてMarkdown本文を返す

`include_body` のdefaultは `false` とする。

`get_record` が本文を返す場合、Markdown本文は編集・整形・要約・正規化せず、元ファイル内容をそのまま返す。
構造化metadataやheadingsは、本文とは別フィールドとして返す。

これは、format確認や後続編集時に、本文が実ファイルと一致していることを保証し、filesystemで再読込する必要を避けるためである。

### `validate_records` の責務

`validate_records` は、Design Records MCPのmetadata indexが信頼できる状態かを検証するtoolである。

MVPでは、少なくとも以下を検査する。

- `id` 重複
- `id` とファイル名番号の不一致
- `kind` 別 `status` 値域違反
- `depends_on` 参照切れ
- `supersedes` 参照切れ
- `migrated_to_spec` 形式違反
- record pathの実在確認

MVPでは、`depends_on` / `supersedes` の参照検査はID存在確認に限定する。
参照元・参照先の `status` 組み合わせ、たとえば `proposed` が `proposed` に `depends_on` できるか、`accepted` が `superseded` に `depends_on` してよいか、`superseded` recordの `depends_on` を検査対象にするかは、後続ADRで判断する。

MVPの `validate_records` は、record metadataの基本整合性検査に限定する。
MVP diagnostic category は以下とする。

- `duplicate_id`
- `filename_id_mismatch`
- `invalid_h1_title`
- `invalid_status_for_kind`
- `missing_depends_on_target`
- `missing_supersedes_target`
- `invalid_migrated_to_spec`
- `missing_record_path`

`accepted_but_not_migrated` や `missing_design_record` のような運用gap診断はMVP外とする。
これらを `validate_records` の追加diagnostic categoryにするか、独立した `list_gaps` toolへ分離するかは、後続ADRで判断する。

specは、ADRで決まったdiagnostic categoryのinput/output schemaと表現形式を定義する。

### `suggest_next_record` の責務

`suggest_next_record` は、新規ADR起票を補助するread-only toolである。

MVPでは、`suggest_next_record` の対象は `kind: decision` に限定する。
spec新規作成のpath提案は、番号制ではなく内容ベース命名になるためMVPでは扱わない。

MVPに含める場合、既存record indexから次のADR IDと推奨pathを返す。

例:

```json
{
  "kind": "decision",
  "title": "Design Records MCP MVP boundary and tool prioritization",
  "next_id": "ADR-077",
  "next_number": 77,
  "suggested_path": "docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md",
  "existing_max_id": "ADR-076"
}
```

`suggest_next_record` はファイル作成を行わない。
ファイル作成、本文生成、既存ファイル更新、Evidence更新はwrite系責務であり、MVP外とする。

### filesystemとの責務境界

Design Records MCPは、汎用filesystem toolの代替ではない。

Design Records MCPが扱うのは、design recordとして意味を持つ読み取り・index・validation・起票補助である。

Design Records MCPが扱う:

- record IDからmetadata / path / headings / raw bodyを取得する
- record一覧を構造化して返す
- record metadataの基本不整合を検証する
- 次のADR番号と推奨pathを提案する

Design Records MCPが扱わない:

- 任意ファイルの読み書き
- Markdown一般編集
- ADR本文の自動生成・自動更新
- commit hashの自動書き換え
- git操作

## 理由

### 1. 本文を読む前の絞り込みが最大の効果を持つ

過去の調査では、関連ADR候補を絞るために複数ADR本文を読む必要があった。
その多くは本文理解ではなく、`status`、`supersedes`、`migrated_to_spec` の確認に使われた。

そのため、MVPでは本文取得よりも先に、`list_records` によるmetadata index queryと、`validate_records` によるmetadata信頼性の確保を優先する。

### 2. `title` はH1から抽出する

titleをfront matterへ追加すると、既存ADRへmetadataを付与する際の移行負荷が増える。
一方、ADRのH1は既存doc-policy上の標準形式であり、人間が読む正式タイトルでもある。

そのため、MVPではtitleをfront matterへ複製せず、H1から抽出する。
H1が存在しない、または `# NNN: タイトル` 形式に合わない場合は、`invalid_h1_title` として検出する。

### 3. `get_record` はfilesystem戻りを避けるために必要である

ADR番号から本文やpathを取得できない場合、Design Records MCPで候補を絞った後も、最終的にはfilesystem toolでファイルを探して読む必要がある。

特にformat確認や編集前確認では、元ファイルと一致した本文が必要になる。
そのため、`get_record` は本文を改変せず、そのまま返す必要がある。

### 4. write系toolはread-only MVPから外すべきである

起票、更新、Evidence commit hash更新などのwrite系toolは便利だが、ユーザー確認、dry-run diff、衝突処理、template責務、git運用との境界を伴う。

これらをMVPに含めると、Design Records MCPの初期目的であるindex / read / validationよりも実装範囲が大きくなる。

そのため、MVPではread-only toolに限定し、write系toolは後続ADRで扱う。

### 5. `suggest_next_record` はwriteではなく運用ルールのqueryである

次ADR番号と推奨pathの算出は、単なるfilesystem操作ではなく、ADR運用ルールに基づくqueryである。

ただし、実際のファイル作成まで含めるとwrite toolになる。
そのため、MVPに含める場合も `suggest_next_record` は提案に留め、作成は行わない。

## 却下した代替案

### 1. `get_record` をMVPから外す

却下する。

`get_record` がない場合、record IDから本文やpathを取得するためにfilesystem toolへ戻る必要がある。
これではDesign Records MCPがrecord indexを持つ利点が弱くなる。

### 2. `get_record` が本文を整形・要約して返す

却下する。

format確認や後続編集では、実ファイルと一致したMarkdown本文が必要である。
本文が整形・要約・正規化されると、編集前にfilesystemで再読込する必要が生じる。

### 3. `trace_record` / `list_gaps` をMVPに含める

却下する。

これらは有用だが、MVPのvalidatorを実データに当ててから、必要なgraph traversalやgap分類の形を判断する方がよい。
初期MVPでは、record metadataの基本整合性検査に限定し、運用gap診断は後続ADRで扱う。

### 4. 起票やEvidence更新などのwrite系toolをMVPに含める

却下する。

write系toolは、dry-run diff、ユーザー確認、衝突処理、template責務、git運用との境界が必要である。
これらはread-only MVPの範囲を超える。

## 影響

### docs/spec/design-records-mcp/**

Design Records MCP specでは、本ADRに基づき、MVP toolの詳細schemaを定義する。

少なくとも以下を定義する。

- `list_records`
- `get_record`
- `validate_records`
- optional: `suggest_next_record`
- MVP diagnostic category

また、`list_records` の `title` はH1由来であり、MVPでは `design_record.title` を持たせないことを明記する。
`get_record` の本文返却はraw bodyとして扱い、構造化metadataやheadingsとは別フィールドで返すことを明記する。

### 実装

MVP実装では、まずread-only index / validation / record取得を実装する。

実装順序の推奨は以下とする。

1. front matter parser + record index
2. `list_records`
3. `validate_records`
4. `get_record`
5. optional: `suggest_next_record`（`kind: decision` のみ）

write系toolは実装しない。

### 既存ADR/spec

ADR-076のbootstrap方針に従い、初期対象recordへ `design_record` metadataを付与した上で、`list_records` と `validate_records` の実データ検証に使う。

## Evidence

- commit: beb3f51
- impl commit: tbd
- 参考: generic support調査時のADR読み込み計測、ADR-076のDesign Records MCP導入方針

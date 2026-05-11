# 076: Design Records MCP

- **status**: proposed
- **date**: 2026-05-11
- **depends_on**: ADR-050, ADR-068
- **supersedes**:
- **migrated_to_spec**:

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

brewprint は ADR-050 により spec-first のドキュメント運用へ移行している。

この運用では、現行仕様は `docs/spec/**` に置き、ADR は設計判断の背景・却下案・トレードオフを記録する。
また、既存ADRの仕様記述は一括移行せず、触れたタイミングで漸進的にspecへ移す方針を採っている。

ADR-068 では、ADR / spec / task / UC docs の責務境界を「所有」「参照」「影響範囲」という観点で整理した。
同ADRは、DecisionRecord / Spec / TaskTracker などの抽象artifact型名、他プロジェクト適用、MCP schema化は別ADRで検討するとしていた。
本ADRは、そのうち ADR/spec 運用を機械可読metadataとMCP queryで支援する方針を扱う。

一方で、ADR数が増えたことで、以下の情報を Markdown本文・filename・自然言語リンク・人間の記憶だけで維持することが難しくなっている。

- ADRが現行specへ反映済みか
- ADRがどのspecに影響するか
- spec sectionがどのADR由来か
- accepted ADRのうちmigration未完了のものはどれか
- superseded ADRが現行specから参照され続けていないか
- ADR同士の depends_on / supersedes に参照切れがないか
- 別セッションのLLMが最初に読むべきADR/specはどれか

これらは、人間が読む説明というより、検索・検証・影響分析の対象である。
そのため、自然言語本文だけではなく、機械可読なmetadataとして扱う必要がある。

## 決定

brewprint 内に Design Records MCP MVP を導入する方針を採用する。

Design Records MCP は、`docs/adr/**` と `docs/spec/**` の Markdown front matter を読み取り、ADR/specのID・種別・状態・基本的な関係・migration状態をindex化し、MCP toolから参照・検証できるようにする。

ただし、本ADRでは Design Records MCP の完全なschemaやtool仕様は決めない。
本ADRで決めるのは、ADR/spec運用を machine-readable metadata と MCP query / validation で支援する方針、およびMVPを最小範囲に留める境界である。

初期実装は brewprint 内部で使うことを前提にする。
最初から独立OSS、汎用CLI、Web UI、複数プロジェクト横断基盤としては設計しない。

一方で、front matter schema は brewprint 固有概念に閉じすぎない形にし、将来的な外部切り出しを妨げないようにする。

MVPの実装範囲は、brewprint本体の進行をブロックしない程度の、2〜3日程度の短期で完結する規模を維持する。
schemaやtoolの拡張は、MVP稼働後の必要性確認を経た後続ADRで扱う。

### MVP対象

MVPのindex対象は ADR と spec に限定する。

task file / UC docs / impl notes は、ADR-068で責務境界上のartifactとして扱われる。
しかし、これらは作業状態・実例・実装引継ぎの性質が強く、ADR/specとは更新頻度や正とする情報が異なる。
そのため、MVPではDesign Records MCPのindex対象に含めない。

task file / UC docs / impl notes を扱うかどうかは、MVP運用後に必要性を確認してから後続ADRで検討する。

### MVPスコープ

MVPでは以下を扱う。

- Markdown front matter の読み取り
- design record index の作成
- `id` / `kind` / `status` の管理
- `depends_on` / `supersedes` の基本的な参照確認
- `migrated_to_spec` の一覧化
- ADR/spec間の基本的な参照切れ検出
- MCP tool による list / get / validate

MVP tool は以下に限定する。

- `list_records`
- `get_record`
- `validate_records`

`trace_record` や `list_gaps` はMVPには含めない。
これらは、MVPのvalidatorを実データに当てて必要なgapやtraceの形が見えてから、後続ADRまたはspecで追加を判断する。

### MVPスコープ外

MVPでは以下を扱わない。

- 自然言語本文から依存関係を推定すること
- spec本文との厳密な意味照合
- git履歴解析
- コード静的解析
- Web UI
- 複数プロジェクト横断管理
- OSSとしての汎用CLI
- brewprint DSLとの深い統合
- spec section単位の完全な正規化
- `trace_record`
- `list_gaps`
- `topics` metadata
- `affects` metadata
- `refines` / `conflicts_with` metadata
- `migration.state` の正規化
- task file / UC docs / impl notes のindex化
- 既存brewprint MCPとの統合方式の決定

### front matter 方針

ADR/specのfront matterには、段階的に `design_record` metadataを追加する。

ただし、MVPではschemaを最小に留める。
最初から `topics`、`affects`、`refines`、`conflicts_with`、`owns`、section単位の由来情報などは正規化しない。

`status` の値域は `kind` ごとに既存 `doc-policy.md` の定義に従う。
`kind: decision` では `proposed` / `accepted` / `superseded` を使う。
`kind: spec` では `confirmed` / `draft` / `wip` を使う。
Design Records MCP の validator は、`kind` 別の `status` 値域を検査する。

ADR側の最小例:

```yaml
design_record:
  id: ADR-076
  kind: decision
  status: proposed
  depends_on:
    - ADR-050
    - ADR-068
  supersedes: []
  migrated_to_spec: null
```

spec側の最小例:

```yaml
design_record:
  id: SPEC-mcp-design-records
  kind: spec
  status: draft
  depends_on:
    - ADR-076
```

MVPでは、既存ADRフォーマットの `migrated_to_spec: YYYY-MM-DD` と互換のあるmetadataを優先する。
`migration.state` のような正規化された状態語彙は、実データで必要性を確認してから後続ADRで決める。

既存ADR/specへのfront matter追加は一括では行わない。
新規ADRから優先し、既存文書は触れたタイミングで漸進的に追加する。

### bootstrap方針

MVPの検証対象が薄くなりすぎないように、最初に少数の代表的なADR/specへ `design_record` metadataを付与する。

初期bootstrap対象は以下を候補とする。

- ADR-067〜ADR-076
- ADR-050: spec-first documentation policy
- `docs/spec/mcp/tools/design-records.md`（本ADRと同時または直後に新設）

それ以前のADR（ADR-001〜ADR-049、ADR-051〜ADR-066）は一括移行しない。
整合性レビュー、新ADR起票、関連spec更新などで触れたタイミングで、漸進的にmetadataを追加する。

### 既存brewprint MCPとの関係

Design Records MCP は、既存brewprint MCPとは対象データソースが異なる。

既存brewprint MCPは、brewprint YAMLから構築された `ResolvedProject` 上のsemantic objectをqueryするためのlayerである。
一方、Design Records MCPは、`docs/adr/**` と `docs/spec/**` のMarkdown front matterをquery / validation対象とする。

既存brewprint MCPへ統合するか、別MCPサーバーとして提供するかは、本ADRでは決めない。
ただし、brewprint YAML semantic object向けの既存 `QueryService` の責務を、docs管理へそのまま拡張しない。

統合方式、tool namespace、実装package構成は、MVP実装前の後続ADRまたはspecで決める。

## 理由

### 1. spec-first方針と整合する

Design Records MCP は、ADR-050のspec-first方針を置き換えるものではない。

現行仕様は引き続きspecに置く。
ADRは判断履歴として残す。
Design Records MCP は、その関係を機械的に辿り、検証しやすくする補助である。

### 2. ADR/spec関係はquery対象である

`status`、`depends_on`、`supersedes`、`migrated_to_spec` などは、自然言語説明ではなく、検索・検証・影響分析の対象である。

そのため、Markdown本文やfilenameだけではなく、front matter上の構造化metadataとして扱う方が適している。

### 3. LLMとの設計対話を支援する

brewprint MCP は、ADR-054で設計対話を支援するquery layerとして位置づけられている。

現在のMCPは主にbrewprint YAMLから構築されたsemantic objectを対象にしているが、brewprint自体の設計運用もまた、LLMとの対話対象である。

Design Records MCP により、別セッションのLLMは、関連するADR/spec、未移行の判断、影響範囲、参照切れを機械的に確認できる。

### 4. 最初から汎用化しないことでスコープ肥大化を防ぐ

Design Records MCP は他プロジェクトでも使える可能性がある。
しかし、最初からOSS化や汎用CLIを目指すと、schema migration、multi-project対応、GitHub連携、UI、権限管理などにスコープが膨らむ。

まずは brewprint 内部の実データに当て、front matter index + validator としてMVPを作る。
外部切り出しは、brewprint内で実用して粗が見えてから判断する。

### 5. MVPを小さく保つ

Design Records MCPは、schemaを大きく設計しようとすれば、`topics`、`affects`、`refines`、`conflicts_with`、section単位のtraceability、task/UC連携などを含みうる。

しかし、既存ADR/specへmetadataを付ける前にこれらを決め切ると、実データに合わないschemaを固定する危険がある。

MVPでは、既存doc-policyと互換しやすい `id` / `kind` / `status` / `depends_on` / `supersedes` / `migrated_to_spec` に絞る。
追加metadataは、MVPのvalidator結果と実運用上の不足を見てから後続ADRで判断する。

## 却下した代替案

### 1. filenameとMarkdown本文だけで管理し続ける

却下する。

ADR/spec数が少ない間は成立するが、参照関係・migration状態・影響範囲が増えるほど、人間の記憶と自然言語リンクに依存する。
参照切れやstale docsを機械的に検出できない。

### 2. 最初から汎用Design Records基盤として作る

却下する。

brewprint外でも使える可能性はあるが、初期段階で汎用化を目的にすると、現在必要なADR/spec運用改善よりも周辺設計が大きくなる。
MVPではbrewprint内部に閉じる。

### 3. 自然言語本文から依存関係を推定する

MVPでは却下する。

LLMや自然言語解析で依存関係を推定することは将来的に補助としてあり得る。
しかし、Design Records MCPの初期目的は、明示された機械可読metadataを検証することである。
推定結果を正とすると、検証基盤としての信頼性が落ちる。

### 4. spec section単位まで最初から正規化する

MVPでは却下する。

section単位のtraceabilityは有用だが、既存docsへ一括適用すると負荷が高い。
初期はfile単位のADR/spec関係と基本的な参照確認に留め、section IDは後続ADRで検討する。

### 5. Design Records MCPを既存QueryServiceへそのまま混ぜる

却下する。

既存QueryServiceは、brewprint YAMLから構築された `ResolvedProject` 上のsemantic objectをqueryする責務を持つ。
ADR/spec front matterを対象にしたdocs管理queryを同じ責務へ混ぜると、既存brewprint MCPの境界が曖昧になる。

統合提供する可能性は残すが、少なくとも既存QueryServiceの責務をdocs管理へそのまま拡張しない。

## 影響

### docs/doc-policy.md

将来的に以下の追記が必要になる可能性がある。

- ADR/spec front matterに `design_record` metadataを持たせる
- ADR番号・filenameではなく `design_record.id` を正とする
- `migrated_to_spec` をDesign Records MCPでも参照可能なmetadataとして扱う
- セッション開始時にDesign Records MCPで関連recordを引く

### docs/adr-authoring-guide.md

ADR起票時のfront matter推奨項目として、以下を追記する可能性がある。

- `design_record.id`
- `design_record.kind`
- `design_record.status`
- `design_record.depends_on`
- `design_record.supersedes`
- `design_record.migrated_to_spec`

### docs/spec/mcp/**

Design Records MCPのtool specを追加する可能性がある。

MVP段階では、以下を単一specにまとめる。

- `docs/spec/mcp/tools/design-records.md`

このspecでは、少なくとも以下を定義する。

- `list_records`
- `get_record`
- `validate_records`
- 最小 `design_record` front matter schema
- validation diagnostic の基本形

`trace_record`、`list_gaps`、詳細なgraph traversal、gap分類は後続ADRまたは後続specで扱う。

### 実装

実装位置はこのADRでは固定しない。

候補としては、以下が考えられる。

- 既存brewprint MCPとは別のMCPサーバー
- 既存MCPサーバー内の別namespace
- `internal/docgov`
- `cmd/brewprint` 配下のdocs validation command

ただし、このADRではGo package構成、tool namespace、別サーバー化の有無までは決めない。

### 既存ADR/spec

既存ADR/specへの `design_record` metadata追加は一括移行しない。

初期bootstrap対象を除き、既存文書は触れたタイミングで漸進的に追加する。

## Evidence

- commit: f92c228
- impl commit: tbd
- 参考: ADR-050のspec-first運用、ADR-068のADR authoring guide、ADR-054のMCP設計対話coverage

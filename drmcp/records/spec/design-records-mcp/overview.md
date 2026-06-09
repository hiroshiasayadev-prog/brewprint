---
scope: docs/spec/design-records-mcp/overview.md
status: draft
last_updated: 2026-05-27
summary: >
  Design Records MCP の目的、対象範囲、既存 brewprint MCP との責務境界、
  MVP の非目標を定義する。
depends_on:
  - docs/adr/076-design-records-mcp.md
  - docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
  - docs/adr/090-design-records-mcp-batch-retrieval-tool-boundary.md
  - docs/adr/092-design-records-mcp-workflow-artifact-record-and-relation-boundary.md
design_record:
  id: V01-SPEC-design-records-mcp-overview
  kind: spec
  status: draft
  depends_on:
    - V01-ADR-076
    - V01-ADR-077
    - V01-ADR-087
    - V01-ADR-088
    - V01-ADR-090
    - V01-ADR-092
---

# Design Records MCP overview

## 目的

Design Records MCP は、brewprint の design record / workflow artifact 運用を machine-readable metadata と MCP query / validation で支援する補助 MCP である。ADR / spec / investigation に加え、requirement / work item / task も対象とする。

主目的は以下である。

- ADR / spec / investigation / requirement / work item / task の record index を構築する
- record の ID / kind / status / path と kind 固有 metadata を構造化して取得できるようにする
- 選択済みの複数 record ID について detail representation をまとめて取得できるようにする
- record metadata の基本不整合を検出する
- docs artifact 間の semantic/artifact ref を解決し、参照切れ検査に利用できるようにする
- 別セッションの LLM が、読むべき design record を本文読解前に絞り込めるようにする

Design Records MCP は、spec-first ドキュメント運用を置き換えない。
現行仕様の唯一の正は引き続き `drmcp/records/spec/**` であり、ADR は設計判断の根拠記録である。
Design Records MCP は、その関係を機械的に辿りやすくする query / validation layer として扱う。

> 由来: V01-ADR-076 §決定, V01-ADR-076 §理由

## 対象 record

index / query / validation 対象は以下とする。パスは `<records_root>` からの相対。MVP では `records_root = v01/records`、`namespace_prefix = V01-`。

| kind | discovery パス | public ID 例 |
|---|---|---|
| `decision` | `<records_root>/adr/*.md` | `V01-ADR-076` |
| `spec` | `<records_root>/spec/**/*.md`（`design_record.id` + `design_record.kind` を持つ file のみ） | `V01-SPEC-design-records-mcp-overview` |
| `investigation` | `<records_root>/investigations/<domain>/<namespace_prefix>INV-*-*.md` | `V01-INV-MCP-001` |
| `requirement` | `<records_root>/requirements/<domain>/<namespace_prefix>REQ-*-*.md` | `V01-REQ-MCP-001` |
| `work_item` | `<records_root>/work-items/<domain>/<namespace_prefix>WORK-*-*.md` | `V01-WORK-DRMCP-001` |
| `task` | `<records_root>/tasks/<domain>/<namespace_prefix>TASK-*-*.md` | `V01-TASK-MCP-001-01` |

`design_record` を持たない既存 spec は index 対象外とし、`missing_design_record` diagnostic も出さない。
legacy M-series は`task` record discovery の対象に含めない。

本specは record kind 全体を閉じた列挙として確定しない。後続判断により他の artifact kind を追加しうる。
MVP では UC docs / impl notes は record kind として index 対象に含めない。

> 由来: V01-ADR-076 §MVP対象, V01-ADR-087 §1, V01-ADR-091 §1〜§6, V01-ADR-092 §1

## 既存 brewprint MCP との責務境界

Design Records MCP は、既存 brewprint MCP とは別の対象データソースを扱う。

| MCP | 対象 | 主な責務 |
|---|---|---|
| brewprint MCP | brewprint YAML から構築された `ResolvedProject` | semantic object の query / inspect / impact analysis |
| Design Records MCP | `<records_root>/adr/**` / `<records_root>/investigations/**` / `<records_root>/requirements/**` / `<records_root>/work-items/**` / `<records_root>/tasks/**` の箇条書きmetadata と、`design_record` を持つ `<records_root>/spec/**` の YAML front matter | design record / workflow artifact の index / read / validation、および traceability spec に従う semantic/artifact ref resolve |

Design Records MCP は、既存 brewprint MCP とは独立して起動・検証できる構成を第一候補とする。
既存 `QueryService` の責務を docs 管理へそのまま拡張しない。

この spec は既存 brewprint MCP の `docs/spec/mcp/tools/**` には混ぜず、`drmcp/records/spec/design-records-mcp/**` に置く。

> 由来: V01-ADR-076 §既存brewprint MCPとの関係

## Resolver responsibility

Design Records MCP は、traceability spec が定める canonical reference model に従い、ref の resolve とその結果を用いた validation を担う。

V01-ADR-088 / V01-ADR-092 により、MVP で必須とする resolver input は active `spec:` semantic ref、Design Records MCP が扱う record の public ID-as-ref（namespace_prefix 付き完全形。MVP では `V01-ADR-*` / `V01-SPEC-*` / `V01-INV-*` / `V01-REQ-*` / `V01-WORK-*` / `V01-TASK-*` 等）、investigation canonical reference validation、および workflow relation integrity validation とする。

Investigation metadata が canonical reference として追加利用できる workflow の public ID-as-ref は requirement / work item 系（MVP では `V01-REQ-*` / `V01-WORK-*`）に限定し、task 系（`V01-TASK-*`）は direct resolver input および workflow artifact 間 relation のみで support する。

`internal-design:` / `coverage:` / `COV-*`、coverage mapping、semantic realization relation は MVP required resolver scope に含めない。Orphan workflow artifact diagnostics、task status 由来 progress projection、workflow 専用 traversal tool も MVP に含めない。

resolver が lookup source として読む artifact と、`list_records` / `get_record` が record kind として公開する artifact は同一集合である必要はない。

Resolver の public tool 名は `resolve_reference` とし、active `spec:` semantic ref と namespace_prefix 付き public record ID-as-ref（`tools.md` 参照）を active lookup input として扱う。`internal-design:` / `coverage:` / `COV-*`、physical path、および未採用 ID form は `unsupported` response とする。Reserved prefix である `yaml:` の public resolver input / direct query response behavior は MVP で定義しない。

> 由来: V01-ADR-087 §4, V01-ADR-088, V01-ADR-092 §3〜§7

## Record scanning と namespace prefix

Design Records MCP は、リポジトリ内の1つの app namespace の records ツリーをスキャンして index を構築する。スキャン対象は `--records-root` フラグで指定する。

**MVP のスキャン対象は `v01/records` のみ（デフォルト）とする。** namespace prefix は `V01-` となる。

### records_root と namespace prefix 導出

`records_root` は、リポジトリルートからの相対パスで `<app-namespace>/records` の形をとる。app namespace ディレクトリ名（`records/` の親）から namespace prefix を機械的に導出する。

導出式: `namespace_prefix = strings.ToUpper(appNamespaceDir) + "-"`

MVP での適用:

| records_root | appNamespaceDir | namespace_prefix |
|---|---|---|
| `v01/records` | `v01` | `V01-` |

将来の app namespace 対応（MVP 外）での導出式適用例（参考）:

| records_root | appNamespaceDir | namespace_prefix |
|---|---|---|
| `drmcp/records` | `drmcp` | `DRMCP-` |

注: 現時点の `drmcp/records/` 配下のファイルは `V01-SPEC-*` ID を持つ移管前コンテンツであり、`drmcp/records` を records_root として使用するには別途ファイル名・ID の re-prefixing が必要となる。この対応は MVP 外とする。

### kind 別 prefix 適用箇所

records ツリー内の artifact は record kind ごとに以下の箇所に namespace prefix を持つ。

| kind | ファイル名 | H1 | metadata id |
|---|---|---|---|
| ADR | ✓ `V01-ADR-NNN-slug.md` | ✓ `# V01-ADR-NNN: title` | なし（箇条書き metadata に `id:` フィールドなし） |
| Spec | なし（slug のみ） | なし（任意 title） | ✓ front matter `design_record.id: V01-SPEC-slug` |
| Investigation | ✓ `V01-INV-DOMAIN-NNN-slug.md` | ✓ `# V01-INV-DOMAIN-NNN: title` | なし |
| Workflow (REQ / WORK / TASK) | ✓ `V01-WORK-DOMAIN-NNN-slug.md` | ✓ `# V01-WORK-DOMAIN-NNN: title` | ✓ 箇条書き `- **id**: V01-WORK-DOMAIN-NNN` |

parser は kind 別の適用箇所から namespace prefix をストリップして bare ID を抽出・検証し、prefix を付与した完全 ID を `record.ID` として返す。bare ID 文法の検証（`ADR-NNN` 形式・`WORK-DOMAIN-NNN` 形式等）は prefix ストリップ後の文字列に対して行う。完全 ID が tools で公開される public record ID となる。

### multi-root スキャン

複数の app namespace records ツリーを同時にスキャンする機能は MVP 外とする。MVP では1プロセスにつき1つの records_root だけを index し、別 root の record は未スキャンとして扱う。root 間の merge / duplicate detection / cross-root relation validation は行わない。

> 由来: V01-ADR-097, V01-ADR-099

## MVP tool set

MVP の P0 tool は以下である。

- `list_records`
- `get_record`
- `get_records`
- `validate_records`
- `resolve_reference`

P1 の任意補助 tool として、以下を許容する。

- `suggest_next_record`

P0 tool は read-only であり、record 本文を読む前の候補絞り込みと metadata 整合性検証を目的とする。

> 由来: V01-ADR-077 §P0: MVP必須tool, V01-ADR-077 §P1: MVPに含めてもよい補助tool, V01-ADR-090 §決定

## MVP 外

MVP では以下を扱わない。

- `trace_record`
- `list_gaps`
- `create_record`
- `update_record`
- `set_evidence`
- その他の write 系 tool
- 自然言語本文から依存関係を推定すること
- spec 本文との厳密な意味照合
- git 履歴解析
- code static analysis
- Web UI
- 複数プロジェクト横断管理
- 汎用 OSS CLI としての公開 contract
- section 単位の完全な traceability
- `topics` / `affects` / `refines` / `conflicts_with` metadata
- legacy M-series task record / UC docs / impl notes の record kind としての index 化
- `internal-design:` / `coverage:` / `COV-*` の resolve と semantic realization relation validation
- coverage mapping query
- orphan requirement / orphan work item / orphan task diagnostics
- task status から work item progress を導出する projection
- workflow 専用 traversal / tree / graph query tool
- task dependency cycle detection / execution order projection
- investigation metadata から `TASK-*` を canonical reference として辿ること

MVP は ADR / investigation / requirement / work item / task の箇条書きmetadata、spec の YAML front matter、H1、path から得られる明示情報だけを扱う。
自然言語本文の推定や運用 gap 診断は、MVP の validator を実データへ当てた後に追加可否を判断する。

> 由来: V01-ADR-076 §MVPスコープ外, V01-ADR-077 §MVP外, V01-ADR-092 §7

## filesystem tool との責務境界

Design Records MCP は汎用 filesystem tool の代替ではない。

Design Records MCP が扱うもの:

- 単一または明示された複数の record ID から metadata / path / headings / raw body を取得する
- record 一覧を構造化して返す
- record metadata の基本不整合を検証する
- 次の ADR 番号と推奨 path を提案する

Design Records MCP が扱わないもの:

- 任意ファイルの読み書き
- Markdown 一般編集
- ADR 本文の自動生成・自動更新
- commit hash の自動書き換え
- git 操作

> 由来: V01-ADR-077 §filesystemとの責務境界, V01-ADR-090 §決定

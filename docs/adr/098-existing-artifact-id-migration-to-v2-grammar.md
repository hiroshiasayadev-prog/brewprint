# 098: 既存 artifact ID の v2 grammar への一括移行

- **status**: accepted
- **date**: 2026-06-08
- **depends_on**: ADR-096, ADR-097
- **supersedes**: ADR-096（「既存 artifact ID 変更なし」の決定のみ）
- **migrated_to_spec**: 

## 背景

ADR-096 は「既存 artifact ID は変更しない」と決定した。しかしその後、ADR-097 により app namespace-first 物理レイアウトへの移行が確定した。

物理レイアウト移行と ID 移行が不一致のまま進むと以下の問題が生じる。

- **UI（REQ-MCP-013）**: artifact を app namespace 単位で表示する際、`REQ-MCP-001`（旧形式）と `DRMCP-REQ-MCP-033`（新形式）が混在し、表示ロジックに「旧形式か否か」の分岐が必要になる。
- **MCP 挙動**: namespace-aware な browse / discovery API を実装する際にも、同様の形式分岐が必要になる。

物理レイアウト移行（WORK-PRODUCT-003）のタイミングで ID も揃えることで、分岐コストを永続させずに済む。現時点では artifact 数が有限であり、script による一括変換が現実的である。

## 決定

既存の全 artifact ID を v2 grammar（`<APP>-<KIND>-<DOMAIN>-<INDEX>`）へ一括移行する。

変換はスクリプトで自動化し、以下をアトミックに更新する:

- ファイル名（ID prefix を含む部分）
- record 本文の `id:` フィールド
- 全ファイル横断の cross-reference（`source_refs` / `work_items` / `depends_on` / `impact_refs` 等に含まれる ID 文字列）

## ID 変換マッピング

### prefix 追加のみ（単純変換）

| 旧形式 | 新形式 | app namespace |
|---|---|---|
| `REQ-MCP-NNN` | `DRMCP-REQ-MCP-NNN` | DRMCP |
| `WORK-MCP-NNN` | `DRMCP-WORK-MCP-NNN` | DRMCP |
| `TASK-MCP-NNN-NN` | `DRMCP-TASK-MCP-NNN-NN` | DRMCP |
| `REQ-DATA-NNN` | `BPDSL-REQ-DATA-NNN` | BPDSL |
| `WORK-DATA-NNN` | `BPDSL-WORK-DATA-NNN` | BPDSL |
| `TASK-DATA-NNN-NN` | `BPDSL-TASK-DATA-NNN-NN` | BPDSL |
| `REQ-RESOLVE-NNN` | `BPDSL-REQ-RESOLVE-NNN` | BPDSL |
| `WORK-RESOLVE-NNN` | `BPDSL-WORK-RESOLVE-NNN` | BPDSL |
| `REQ-SELFHOST-NNN` | `BPDSL-REQ-SELFHOST-NNN` | BPDSL |
| `WORK-SELFHOST-NNN` | `BPDSL-WORK-SELFHOST-NNN` | BPDSL |
| `TASK-SELFHOST-NNN-NN` | `BPDSL-TASK-SELFHOST-NNN-NN` | BPDSL |
| `INV-DATA-NNN` | `BPDSL-INV-DATA-NNN` | BPDSL |
| `INV-DOCS-NNN` | `PRODUCT-INV-DOCS-NNN` | PRODUCT（DOCS domain は暫定） |

### prefix 追加 + domain 変換（PRODUCT 系）

現行の `PRODUCT` は domain identifier として使われてきた移行期の形式である。v2 では domain を明示する。現時点の全 PRODUCT artifacts は `NAMESPACE` domain に属する（namespace model / governance / migration の関心領域）。

| 旧形式 | 新形式 |
|---|---|
| `REQ-PRODUCT-NNN` | `PRODUCT-REQ-NAMESPACE-NNN` |
| `WORK-PRODUCT-NNN` | `PRODUCT-WORK-NAMESPACE-NNN` |
| `TASK-PRODUCT-NNN-NN` | `PRODUCT-TASK-NAMESPACE-NNN-NN` |
| `INV-PRODUCT-NNN` | `PRODUCT-INV-NAMESPACE-NNN` |

## 理由

**UI / MCP 実装の単純化**: ID が統一形式になれば、`<APP>-<KIND>-<DOMAIN>-<INDEX>` の grammar だけで帰属 app namespace を parse できる。旧形式との分岐ロジックが不要になる。

**移行コストの前倒し**: 物理レイアウト移行（WORK-PRODUCT-003）と同一作業ウィンドウで実施することで、二度手間を避ける。artifact 数は現時点で有限（docs 634 ファイル）であり、script による自動変換が現実的。

**PRODUCT domain の明示化**: `REQ-PRODUCT-003` の `PRODUCT` が app namespace なのか domain なのか曖昧だった問題が解消される。

## 却下した代替案

**旧形式を据え置き、UI / MCP でマッピングテーブルを持つ案**: namespace-model spec の mapping rule（論理マッピング規則）で対応できるが、コードベースに形式判定ロジックが永続する。UI・MCP・将来のツールそれぞれが同じマッピングを実装する必要があり、メンテナンスコストが高い。

## 影響

- `docs/` 全 Markdown ファイルの ID 文字列が一括更新される
- Go コードに artifact ID 形式の validator（正規表現等）が存在する場合、新形式に対応した更新が必要（WORK-PRODUCT-003 内で確認・対応）
- ADR-096 の「既存 artifact ID 変更なし」の決定を本 ADR が supersede する。ADR-096 の他の決定（所有権・物理 migration タイミング）は本 ADR のスコープ外
- namespace-model spec の mapping rule セクション（論理マッピング規則の参照用途として記載）は migration 完了後に削除または「移行済み」として更新する

## Evidence

- ADR-096: 既存 artifact の PRODUCT namespace 所有と per-app migration 非実施（本 ADR が ID 変更なし決定を supersede）
- ADR-097: app namespace-first repository directory layout の採用
- REQ-MCP-013: Design Records MCP record browser UI（UI 実装時の形式分岐問題の根拠）
- REQ-PRODUCT-003: App namespace-first repository directory layout model
- WORK-PRODUCT-003: App namespace-first layout model 仕様化と migration 方針決定
- INV-PRODUCT-001: migration impact inventory（変換対象スケールの根拠）

# V01-INV-PRODUCT-001: app namespace-first repository layout への migration impact inventory

- **status**: concluded
- **date**: 2026-06-08
- **trigger**: V01-WORK-PRODUCT-003 TASK-02。V01-ADR-097 で確定した app namespace-first layout model への migration を計画するにあたり、`docs/` と `internal/` の移動対象・規模・エッジケースを事前に把握する必要があった
- **scope**: `docs/` 配下の全ファイルの app namespace 別分類、`internal/` Go パッケージの移動対象特定、帰属が曖昧なエッジケースの列挙と判断
- **non_scope**: 実際のファイル移動、互換方針の決定、MCP API のパス変更対応
- **source_refs**:
  - V01-ADR-097
  - V01-REQ-PRODUCT-003
  - V01-WORK-PRODUCT-003
- **follow_up_candidates**:
  - V01-WORK-PRODUCT-003 TASK-03: compatibility / legacy path 方針決定
  - V01-WORK-PRODUCT-003 TASK-04: migration plan or explicit defer

## 調査スコープ

- `docs/` 配下の全ファイルを列挙し、app namespace（drmcp / bpdsl / product）別に分類する
- `internal/` 配下の Go パッケージを列挙し、移動先 app namespace を特定する
- 帰属が複数 namespace にまたがる・不明なものをエッジケースとして記録し、判断を取得する

## 調査結果

### 全体スケール

| 対象 | ファイル数 |
|---|---|
| `docs/` 配下 総計 | 634 |
| `internal/` Go ファイル 総計 | ~110 |

`docs/` 内訳:

| 種別 | 件数 |
|---|---|
| ADR | 97 |
| Requirements | 46 |
| Work-items | 48 |
| Tasks | 202 |
| Spec | 40 |
| Investigations | 12 |
| Guides | 7 |
| Impl | 30 |
| UC (Use Cases) | 143 |
| Internal-design | 1 |
| Top-level docs | 8 |

### docs/ 分類

#### drmcp/records/ への移動対象（概算 ~185 files）

**Workflow artifacts**:
- `docs/requirements/mcp/` 全体（32 files）
- `docs/work-items/mcp/` 全体（27 files）
- `docs/tasks/mcp/` 全体（111 files）

**Spec**:
- `docs/spec/design-records-mcp/`（3 files: overview.md, schema.md, tools.md）
- `docs/spec/mcp/`（4 files: overview.md, schema.md, versioning.md, errors.md + tools/）

**ADR**:
- V01-ADR-076〜090（Design Records MCP MVP から semantic trace, workflow artifacts, authoring transaction まで）
- V01-ADR-093（authoring transaction model）
- V01-ADR-096（既存 artifact の PRODUCT namespace 所有と per-app migration 非実施）

#### bpdsl/records/ への移動対象（概算 ~200+ files）

**Workflow artifacts**:
- `docs/requirements/data/`（8 files）
- `docs/requirements/resolve/`（1 file）
- `docs/requirements/self-hosting/`（1 file）← エッジケース判断結果参照
- `docs/work-items/data/`（15 files）
- `docs/work-items/resolve/`（1 file）
- `docs/work-items/self-hosting/`（1 file）← エッジケース判断結果参照
- `docs/tasks/data/`（59 files）
- `docs/tasks/resolve/`（3 files）
- `docs/tasks/self-hosting/`（3 files）← エッジケース判断結果参照

**Spec**:
- `docs/spec/file-types.md`, `type-ref.md`, `naming.md`, `overview.md`, `nodes.md`, `edges.md`, `diagnostics.md`
- `docs/spec/views/` 全体（8 files）
- `docs/spec/project-layout.md`

**ADR**:
- V01-ADR-001〜075（node type 定義・name resolution・DAG/ER/sequence/state/wireframe・task IO・type system 等、BPDSL コア設計判断群）
- V01-ADR-083（project artifact boundary — YAML as implementation source）

**Impl / UC**:
- `docs/impl/go-m*.md`（BPDSL 実装マイルストーン記録）
- `docs/uc/001-ec-checkout-flow/`（BPDSL example）
- `docs/uc/003-task-file-helper-model/`（BPDSL feature demo）

#### product/records/ への移動対象（概算 ~54 files）

**Workflow artifacts**:
- `docs/requirements/product/` 全体（3 files）
- `docs/work-items/product/` 全体（3 files）
- `docs/tasks/product/` 全体（4 files）

**Spec**:
- `docs/spec/concepts/` 全体（namespace-model, repository-layout, project-artifact-model, traceability）

**Guides**:
- `docs/guides/artifact-boundary.md`, `adr-authoring.md`, `investigation-authoring.md`

**ADR**:
- V01-ADR-050（spec-first documentation policy）
- V01-ADR-068（ADR authoring guide）
- V01-ADR-081, 082, 084, 085, 086, 087, 088, 089（workflow artifact / investigation / semantic trace 設計判断）
- V01-ADR-091, 092, 094（workflow artifact status vocabulary）
- V01-ADR-095（YAML DSL と DRMCP の結合境界）← エッジケース判断結果参照
- V01-ADR-097（app namespace-first repository layout）

**Investigations**:
- `docs/investigations/docs/` 全体（7 files）

**UC**:
- `docs/uc/002-brewprint-self-hosting/`（product-wide self-hosting practice）

**Top-level docs**:
- `docs/doc-policy.md`, `docs/third-party-notices.md`, `docs/adr-authoring-guide.md`, `docs/spec-authoring-guide.md` 等

### internal/ 分類

#### drmcp/src/ への移動対象（29 Go files）

- `internal/designrecords/`（22 files）— Design Records store の読み書き・索引・解析・validation・reference resolution
- `internal/designrecordsmcp/`（7 files）— MCP JSON-RPC server, tools ディスパッチ

内部サブ構造案: `drmcp/src/records/`（designrecords 相当）+ `drmcp/src/mcp/`（designrecordsmcp 相当）

#### bpdsl/src/ への移動対象（81 Go files）

- `internal/rawyaml/`（13 files）— YAML AST parsing, raw model representation
- `internal/resolve/`（20 files）— identity resolution, type resolution, validation
- `internal/semantic/`（18 files）— semantic object model（resolved, type-checked）
- `internal/source/`（3 files）— YAML source file loading & classification
- `internal/render/`（26 files）— render engine（dag, er, model, sequence, state, wireframe 等）
- `internal/query/`（21 files）— reference resolution, impact analysis, diagram element query

#### CLI エントリポイント

- `internal/mcp/`（5 files）— `cmd/brewprint/main.go` からのみ import される汎用 MCP server base。DRMCP / BPDSL のどちらからも直接 import されていない ← エッジケース判断結果参照
- `cmd/brewprint/main.go` — CLI エントリポイント。`internal/mcp/` を使ってサーバーを起動する

## エッジケース判断記録

調査で浮上した3点について、2026-06-08 に判断を取得した。

### 1. `internal/mcp/`（汎用 MCP base library）

**調査所見**: `cmd/brewprint/main.go` からのみ import されており、`internal/designrecordsmcp/` を含む他のパッケージからは import されていない。DRMCP と BPDSL は事実上完全に分離済み。

**判断**: 共有ライブラリではない。`internal/mcp/` は `cmd/brewprint/main.go`（BPDSL CLI オーケストレーション層）に付随するため、`bpdsl/src/` または `cmd/` 相当の場所に含める。

### 2. SELFHOST artifacts（REQ-SELFHOST-\* / WORK-SELFHOST-\* / TASK-SELFHOST-\*）

**調査所見**: 「任意 app に適用できる cross-app 検証活動」として namespace-model spec に記載されているが、実態は BPDSL の dogfooding（BPDSL が自身の仕様をユースケとして検証する活動）。DRMCP はすでに自身のワークフローで DRMCP を運用しているため、SELFHOST として別扱いする必要がない。

**判断**: `bpdsl/records/` へ。SELFHOST は BPDSL 固有の検証活動であり、cross-app namespace を設ける必要はない。

### 3. V01-ADR-095（YAML DSL と Design Records MCP の結合境界）

**調査所見**: DRMCP と BPDSL 両方の結合境界を定める ADR。帰属が二重。

**判断**: `product/records/` へ。cross-app 間の境界を定める判断は PRODUCT namespace の関心。

## 結論

- `docs/` + `internal/` の migration 対象規模は合計 ~740 ファイル（docs 634 + Go ~110）
- app namespace 別の分類方針は上記の通り確定した
- TASK-03（compatibility / legacy path 方針）・TASK-04（migration plan or explicit defer）に引き渡す

# REQ-PRODUCT-003: App namespace-first repository directory layout model

- **id**: REQ-PRODUCT-003
- **status**: accepted
- **date**: 2026-06-08
- **source_refs**:
  - REQ-PRODUCT-001
  - REQ-PRODUCT-002
- **work_items**:
  - WORK-PRODUCT-003
  - WORK-PRODUCT-004

## Requirement

現在の `docs/`（kind-first）と `internal/`（app namespace 不明）という物理レイアウトを、app namespace を軸とする構造に再定義する。

各 app namespace はトップレベルディレクトリを持ち、その中に以下の3つの sub-directory を必要に応じて持つ。

- `records/` — 判断履歴（ADR・spec）とワークフロー（requirement・work item・task・investigation）を格納する。現在の `docs/` 配下の design records に相当する。
- `dsl/` — BPDSL YAML 定義を格納する。domain model の source of truth。
- `src/` — 実装コードを格納する。原則として `dsl/` から生成されることを意図するが、BPDSL が対応するまでの間は手書きも許容する。

3つの sub-directory はいずれも必須ではない。対象 app namespace に当該 concern が存在する場合のみ置く。

予想されるトップレベル app namespace ディレクトリ:

| directory | app namespace | 現状の対応 |
|---|---|---|
| `drmcp/` | Design Records MCP | `docs/` の drmcp 関連 + `internal/designrecords/` |
| `bpdsl/` | Brewprint DSL | `docs/` の bpdsl 関連 |
| `product/` | cross-app / repository-wide | `docs/requirements/product/` 等の cross-app records |

`dsl/ → src/` 生成パターンは全 app namespace が長期的に目指す方向である。現状の `internal/designrecords/` は BPDSL が追い付いていない状態の `drmcp/src/` に相当する。

## Evidence
- REQ-PRODUCT-001 が定義した app namespace / domain namespace モデルの物理表現がまだ存在しない。
- 現状の `docs/` は kind-first（`requirements/mcp/`、`requirements/data/`）であり、どの record が何の app namespace に属するかがディレクトリ構造から読み取れない。
- 現状の `internal/` は Go 実装の慣習的な置き場であり、app namespace との対応が不明である。
- REQ-PRODUCT-001 は「namespace-first physical layout migration work」を将来の prerequisite 対象として明示していた。
- `internal/designrecords/` の実装は BPDSL YAML から生成されるべきものが手書きされている状態であり、`dsl/ → src/` パターンへの移行が必要である。

**外部先行事例との照合:**

- **app namespace-first ディレクトリ**: Nx は `apps/<app-name>` 構造を標準とし、プロジェクトをスコープ（application）単位でグルーピングすることを推奨する（[Nx Folder Structure](https://nx.dev/docs/concepts/decisions/folder-structure)）。Go monorepo においても `internal/` を bounded context の境界として使うパターンが確立されている（[Bounded Contexts in a Go Monorepo](https://dev.to/gabrielanhaia/bounded-contexts-in-a-go-monorepo-how-internal-becomes-the-boundary-1dod)）。
- **`dsl/ → src/` 生成パターン**: Protobuf/gRPC エコシステムでは `proto/<domain>/v1/` にスキーマ定義を置き、`gen/go/` に生成コードを出力するパターンが標準化されている。Buf monorepo ではドメイン別モジュール（`proto/user/`、`proto/billing/` 等）と一元的なコード生成設定を組み合わせる（[Buf Monorepo for Go gRPC](https://medium.com/@cassius.paim/hands-on-buf-monorepo-for-go-grpc-a-multi-module-protobuf-architecture-2fd47d16b6a2)、[Go: Structuring with Protocol Buffers](https://blog.dsb.dev/posts/structuring-repositories-with-protocol-buffers/)）。今回の `dsl/` → `src/` は構造的にこのパターンと同一である。
- **`records/` per app namespace**: Martin Fowler は ADR をコードと同じリポジトリに置くことを推奨する（[Architecture Decision Record](https://martinfowler.com/bliki/ArchitectureDecisionRecord.html)）。app namespace 別に `records/` を分割するのはこの per-service docs アプローチの自然な延長であり、反トレンドではない。

## Required Outcome

- app namespace を軸とする repository top-level directory 構造が定義されている。
- `records/` / `dsl/` / `src/` の意味・用途・任意性が仕様化されている。
- `dsl/ → src/` 生成パターンが全 app namespace の意図する方向として明示されている。
- 現状の `docs/` と `internal/` の migration target が宣言されている。
- migration 実行のための work item が立てられている。

## Explicitly Excluded Scope

- 実際のファイル・ディレクトリの移動（migration 実行は work item に委ねる）
- BPDSL による DSL 生成の実装
- MCP API のパス変更対応
- `records/` 内部のサブディレクトリ構造の変更（現行の `adr/`、`spec/`、`requirements/`、`work-items/`、`tasks/`、`investigations/` 構造は維持）

## Boundary

このREQは repository layout model の定義を所有する。migration の実行・DSL 生成実装・MCP API 変更は所有しない。

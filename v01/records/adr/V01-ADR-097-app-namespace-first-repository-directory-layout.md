# V01-ADR-097: app namespace-first repository directory layout の採用

- **status**: accepted
- **date**: 2026-06-08
- **depends_on**: V01-ADR-095, V01-ADR-096
- **supersedes**: 
- **migrated_to_spec**: docs/spec/concepts/repository-layout/index.md

## 背景

V01-REQ-PRODUCT-001 / V01-REQ-PRODUCT-002 により app namespace と domain namespace のモデルが定義された。しかし物理ディレクトリ構造はこれに追随しておらず、現状以下の問題がある。

- `docs/` は kind-first（`requirements/mcp/`、`requirements/data/` など）であり、どの record が何の app namespace に属するかがディレクトリ構造から読み取れない。
- `internal/` は Go 実装の慣習的な置き場であり、app namespace との対応が不明である。
- `internal/designrecords/` の実装は BPDSL YAML から生成されるべきものが手書きされている状態であり、「DSL 定義 → 実装」という意図がディレクトリ構造に反映されていない。

V01-REQ-PRODUCT-001 は「namespace-first physical layout migration work」を将来の prerequisite 対象として予告していた。V01-REQ-PRODUCT-003 はこの物理レイアウトモデルの仕様化を要求する。

## 決定

repository のトップレベルを app namespace 基準で切る。各 app namespace ディレクトリは以下の3つの sub-directory を必要に応じて持つ。

```
<app-namespace>/
  records/   # 判断履歴（ADR・spec）とワークフロー（requirement・work item・task・investigation）
  dsl/       # BPDSL YAML 定義（domain model の source of truth）
  src/       # 実装コード（原則として dsl/ から生成）
```

3つの sub-directory はいずれも必須ではない。対象 app namespace に当該 concern が存在する場合のみ置く。

予想されるトップレベル app namespace ディレクトリ:

| directory | app namespace | 現状の対応 |
|---|---|---|
| `drmcp/` | Design Records MCP | `docs/` の drmcp 関連 + `internal/designrecords/` |
| `bpdsl/` | Brewprint DSL | `docs/` の bpdsl 関連 |
| `product/` | cross-app / repository-wide | `docs/requirements/product/` 等の cross-app records |

`dsl/ → src/` 生成パターンは全 app namespace が長期的に目指す方向とする。現状の `internal/designrecords/` は BPDSL が追い付いていない状態の `drmcp/src/` に相当する。

実際のファイル・ディレクトリ移動は本 ADR のスコープ外とし、V01-WORK-PRODUCT-003 に委ねる。

## 理由

**namespace model との一致**: V01-ADR-095 / V01-REQ-PRODUCT-001 が定義した app namespace モデルを物理ディレクトリに反映することで、ディレクトリ構造を見るだけで「どの app の concern か」が分かるようになる。

**外部先行事例との整合**:

- app namespace-first directory は Nx の `apps/<app-name>` 構造や DDD の bounded context per directory と一致しており、業界標準のモノレポ設計に沿っている。
- `dsl/ → src/` 生成パターンは Protobuf/gRPC エコシステムの `proto/<domain>/` → `gen/go/` と構造的に同一であり、Buf monorepo で広く採用されている実績がある。
- `records/` per app namespace は Martin Fowler が推奨する「ADR をコードと同じリポジトリに置く」アプローチの per-service 拡張であり、反トレンドではない。

**DSL 駆動開発の意図の明示**: `drmcp/src/` が `drmcp/dsl/` から生成されるべき存在であることをディレクトリ構造で示せる。現状の `internal/` という置き場ではこの意図が読み取れない。

## 却下した代替案

**kind-first 維持（`docs/requirements/mcp/` 等の現状維持）**: 既存の kind-first 構造はツールの走査には扱いやすいが、app namespace の境界がディレクトリ構造に現れない。UI や namespace-aware MCP discovery を実装する際に「どの app の concern か」を推論で補う必要が生じる。

**`src/` のみ app namespace で切り、`docs/` は kind-first のまま**: 実装と設計記録でディレクトリ軸が異なるため、「DRMCP に関わる全ての artifact」を辿るのにディレクトリが分散する。namespace モデルを一貫させる意義が薄れる。

## 影響

- `docs/` 配下の design records は各 app namespace の `records/` を migration target とする
- `internal/designrecords/` は `drmcp/src/` を migration target とする
- 実際の migration タイミング・手順・互換方針は V01-WORK-PRODUCT-003 で決定する
- `records/` 内部のサブディレクトリ構造（`adr/`、`spec/`、`requirements/`、`work-items/`、`tasks/`、`investigations/`）は変更しない

## Evidence

- V01-REQ-PRODUCT-001: App and domain namespace model
- V01-REQ-PRODUCT-002: Domain namespace internal subdomain grouping model
- V01-REQ-PRODUCT-003: App namespace-first repository directory layout model
- V01-ADR-095: YAML DSL と Design Records MCP の結合境界
- V01-ADR-096: 既存 artifact の PRODUCT namespace 所有と per-app migration 非実施
- [Nx Folder Structure](https://nx.dev/docs/concepts/decisions/folder-structure)
- [Buf Monorepo for Go gRPC](https://medium.com/@cassius.paim/hands-on-buf-monorepo-for-go-grpc-a-multi-module-protobuf-architecture-2fd47d16b6a2)
- [Go: Structuring with Protocol Buffers](https://blog.dsb.dev/posts/structuring-repositories-with-protocol-buffers/)
- [Bounded Contexts in a Go Monorepo](https://dev.to/gabrielanhaia/bounded-contexts-in-a-go-monorepo-how-internal-becomes-the-boundary-1dod)
- [Architecture Decision Record (Martin Fowler)](https://martinfowler.com/bliki/ArchitectureDecisionRecord.html)

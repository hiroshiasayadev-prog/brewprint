---
scope: docs/spec/concepts/repository-layout/index.md
status: draft
last_updated: 2026-06-08
summary: >
  brewprint repository 自体のディレクトリレイアウトモデルを定義する。
  top-level を app namespace 基準で切り、各 app namespace が records/ / dsl/ / src/ を持つ構造を定める。
  dsl/ → src/ 生成パターンを全 app namespace の意図する方向として明示する。
depends_on:
  - docs/adr/097-app-namespace-first-repository-directory-layout.md
  - docs/adr/095-yaml-dsl-design-records-mcp.md
  - docs/adr/096-artifact-product-namespace-per-app-migration.md
semantic_refs:
  - spec:repository-layout
sections:
  spec:repository-layout.model: Repository directory model
  spec:repository-layout.sub-directories: Sub-directory definitions
  spec:repository-layout.generation-pattern: dsl → src generation pattern
  spec:repository-layout.records-internal: records/ internal structure
  spec:repository-layout.current-state: Current state mapping
---

# Repository layout model

## 目的

brewprint repository 自体のディレクトリ構造を、app namespace を軸として定義する。本 spec は BPDSL を使うプロジェクトのレイアウト（`yaml/` / `renders/` 等、`docs/spec/project-layout.md` で定義）とは独立した、repository 管理構造の仕様である。

本 spec は `PRODUCT` namespace の関心（cross-app governance）として管理される。

## この spec が所有するもの / 所有しないもの

所有するもの:

- app namespace-first top-level directory 構造の定義
- `records/` / `dsl/` / `src/` の意味・用途・任意性
- `dsl/ → src/` 生成パターンの intent 定義
- `records/` 内部のサブディレクトリ構造の方針
- 現状の `docs/` と `internal/` の migration target 宣言

所有しないもの:

- 実際のファイル・ディレクトリの移動（V01-WORK-PRODUCT-003 に委ねる）
- BPDSL による DSL 生成の実装
- MCP API のパス変更対応
- namespace catalog の formal schema・機械可読 registry 形式
- BPDSL プロジェクトのレイアウト（`docs/spec/project-layout.md` が所有）

## Repository directory model

top-level を app namespace 基準で切る。

```
<app-namespace>/
  records/   # 判断履歴（ADR・spec）とワークフロー（REQ・WORK・TASK・INV）
  dsl/       # BPDSL YAML 定義（domain model の source of truth）
  src/       # 実装コード（原則として dsl/ から生成）
```

現在定義されている app namespace ディレクトリ:

| directory | app namespace | 性質 |
|---|---|---|
| `drmcp/` | Design Records MCP | 設計記録管理 MCP サーバー |
| `bpdsl/` | Brewprint DSL | brewprint 設計記述 DSL |
| `product/` | cross-app / repository-wide | cross-app ポリシー・ガバナンス |

`product/` は実行時コンポーネントを持たないため `records/` のみを持つ。`dsl/` と `src/` は持たない。

## Sub-directory definitions

### records/

判断履歴とワークフローを格納する。

- ADR（Architecture Decision Records）
- spec（仕様）
- requirement / work item / task / investigation（ワークフロー artifact）

現状の `docs/` 配下の design records がここに対応する。

### dsl/

BPDSL YAML 定義を格納する。app namespace が扱う domain model の source of truth。

この dir が存在する場合、`src/` の実装コードは原則として `dsl/` から生成される。

### src/

実装コードを格納する。原則として `dsl/` から生成されることを意図するが、BPDSL が対応するまでの間は手書きも許容する。

現状の `internal/` 配下の Go 実装コードがここに対応する。

### 任意性

3つの sub-directory はいずれも必須ではない。対象 app namespace に当該 concern が存在する場合のみ置く。

## dsl → src generation pattern

`dsl/ → src/` 生成パターンは全 app namespace が長期的に目指す方向とする。

現状:

- `bpdsl/`: `dsl/` に BPDSL YAML 定義を持ち、`src/` が生成対象
- `drmcp/`: `src/` のみ存在（現 `internal/designrecords/`）。BPDSL が追い付いた段階で `dsl/` を追加し生成パターンに移行する

この intent はディレクトリ構造に以下の意味を付与する: `drmcp/src/` の手書き実装は BPDSL 未対応の暫定状態であることを明示する。

Protobuf/gRPC エコシステムにおける `proto/<domain>/` → `gen/go/` パターンが外部先行事例として対応する（V01-ADR-097 参照）。

## records/ internal structure

`records/` 内部のサブディレクトリ構造は現行の kind-first 構造を維持する。

```
records/
  adr/
  spec/
  requirements/
  work-items/
  tasks/
  investigations/
```

`records/` 内部の kind-first → namespace-first への再編は本 spec のスコープ外とする。

## Current state mapping

現状の `docs/` と `internal/` から target layout への migration 対象:

| 現状 | target | 対象 app |
|---|---|---|
| `docs/` の drmcp 関連 records | `drmcp/records/` | DRMCP |
| `docs/` の bpdsl 関連 records | `bpdsl/records/` | BPDSL |
| `docs/requirements/product/` 等 cross-app records | `product/records/` | PRODUCT |
| `internal/designrecords/` | `drmcp/src/` | DRMCP |

実際の migration 実行・互換方針・タイミングは V01-WORK-PRODUCT-003 で決定する。

## 由来

- V01-ADR-097: app namespace-first repository directory layout の採用
- V01-ADR-095: YAML DSL と Design Records MCP の結合境界
- V01-ADR-096: 既存 artifact の PRODUCT namespace 所有と per-app migration 非実施
- V01-REQ-PRODUCT-003: App namespace-first repository directory layout model
- V01-WORK-PRODUCT-003: App namespace-first layout model 仕様化と migration 方針決定

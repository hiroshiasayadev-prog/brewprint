---
scope: docs/spec/naming.md
status: wip
last_updated: 2026-04-29
summary: >
  brewprint の名前空間とID解決ルールを定義する。
  モジュール=フォルダ階層、QualifiedID形式、同モジュール内ID直書き、sentinel方式パースを含む。
  resolve層レビューで肉付け予定（actor global / FK解決等は未収載）。
depends_on:
  - docs/adr/002-folder-as-namespace.md
  - docs/adr/003-name-resolution-rules.md
  - docs/adr/027-module-nesting-and-name-resolution.md
open_issues:
  - actor の global 定義 / 名前解決ルール（ADR-031 / ADR-025 由来。resolve 層レビューで追記予定）
  - FK 解決ルール（ADR-033 由来。resolve 層レビューで追記予定）
  - cross-edge 管理（ADR-020 由来）
---

# 名前解決仕様

> このspecは最小骨格。resolve層レビューで未収載項目（actor global / FK解決 / cross-edge等）を追記する予定。

## 1. モジュールとフォルダ階層

`yaml/` 配下のディレクトリ階層がそのままモジュール階層となる。

```
yaml/
  auth/
    dag.yaml        ← モジュール: auth
  commerce/
    cart/
      dag.yaml      ← モジュール: commerce.cart
    order/
      dag.yaml      ← モジュール: commerce.order
```

ディレクトリの深さに制限はない（任意の段数のネストを許容）。モジュール名は `.` 区切りで表現する。

> 由来: ADR-002 §決定, ADR-027 §モジュールネストを許容する

## 2. QualifiedID

ノードのフルパス参照は以下の形式を取る。

```
<モジュールパス>.<ノード種別>.<ID>
```

- `<モジュールパス>`: 1段以上の任意の深さのドット区切りパス（例: `auth`, `commerce.cart`）
- `<ノード種別>`: 後述の予約語（sentinel）
- `<ID>`: モジュール内のノードID

例:

```yaml
auth.task.login
commerce.cart.store.cart_db
analysis.state.session
```

> 由来: ADR-027 §命名規則を拡張する

## 3. ノード種別予約語（sentinel）

QualifiedID のパースは、ノード種別予約語の出現位置でモジュールパスと ID の境界を判定する。

| 予約語 | ノード種別 |
|--------|-----------|
| `task` | タスク |
| `model` | モデル |
| `store` | ストア |
| `actor` | アクター |
| `event` | イベント |
| `state` | ステート |
| `branch` | ブランチ |
| `fork` | フォーク |
| `join` | ジョイン |

パース例:

```
auth.oauth.task.login
           ↑
           sentinel を発見 → 左側がモジュールパス、右側が ID
→ module: auth.oauth
→ type:   task
→ id:     login
```

> 由来: ADR-027 §ノード種別予約語をsentinelとして正式化

## 4. 同モジュール内ID直書き

同モジュール内のノード参照では QualifiedID のフルパスは要求しない。ID 直書きで解決される。
モジュールを跨ぐ参照は QualifiedID（フルパス）必須。

```yaml
# 同モジュール内（auth/dag.yaml内での参照）
flow:
  - from: login           # auth.task.login に解決される
    to: token

# モジュール跨ぎ（フルパス必須）
edges:
  - from: auth.task.login
    to: analysis.state.session
```

> 由来: ADR-003 §決定, ADR-027 §フルパス必須（同モジュール内ID直書きは継承）

## 5. 未収載項目

以下は resolve 層レビューで追記される予定（現時点では関連 ADR を直接参照）:

- **actor の global 定義**: actor は global namespace を持つ（ADR-031）。プロジェクト全体で一意な actor ID を共有する仕組み
- **actor の YAML 配置**: `actors.yaml` のようなファイルで定義する（ADR-025）
- **FK 解決ルール**: ER における FK 参照の名前解決（ADR-033）
- **cross-edge 管理**: モジュール跨ぎエッジの管理ルール（ADR-020）

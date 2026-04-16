# 007: assetとstoreの境界定義

- **status**: accepted
- **date**: 2026-04-17

## 背景

DAGのノード種別として `asset`（旧: artifact）と `store`（旧: state）を持つが、
「データを扱う」という点で両者の境界が曖昧になりうる。
特に「複数のassetを内包するstruct」をどちらに分類するかが論点になった。

## 決定

> **asset = 型定義（データの形・schema）**
> **store = 実行時にデータを保持する実体（インスタンス）**

| 種別 | 本質 | 問い |
|------|------|------|
| `asset` | 型定義 | "どんな形をしているか？" |
| `store` | 実行時インスタンス | "どこに・何が住んでいるか？" |

### assetのkind

| kind | 例 |
|------|---|
| `scalar` | int, float, str, bytes |
| `struct` | `User { id, name, email }`。他のassetをフィールドに持つネストも可 |
| `list` | `List<User>`（型としての ordered sequence） |
| `dict` | `Dict<str, User>` |

### storeのkind

| kind | 例 |
|------|---|
| `db` | PostgreSQLのテーブル・コレクション |
| `session` | Streamlit session_state・HTTPセッション |
| `collection` | 特定のassetを保持しクエリを持つリポジトリ（`of: <asset_id>`） |
| `context` | React Context・DIコンテナ |

### collection の制約

`store.kind: collection` の `of:` に指定できるのは **assetのIDのみ**。
storeのネストは禁止。storeのネストが必要に見えるケースは `task` でつなぐ設計に切り出す。

## 理由

OOPにおける **class（型定義）vs instance（実行時実体）** の区別をそのまま適用した。

- 「複数のassetを内包するstruct」は `asset.kind: struct` で表現できる（型のネストであり新種別不要）
- 「実際にデータを保持しクエリを受け付けるもの」は `store` に限定することで、型と実体の混在を防ぐ

参照した公知概念：
- OOP の class vs instance（型理論の基本区別）
- DDD の Repository パターン（`store.kind: collection` の設計根拠）

## 影響

- assetのフィールドに別のasset IDを型として参照することは正規の使い方
- storeがstoreを参照する構造は仕様上禁止
- `spec/nodes.md` にてフィールド定義を詳細化する

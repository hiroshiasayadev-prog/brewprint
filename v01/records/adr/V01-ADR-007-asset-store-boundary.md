# V01-ADR-007: assetとstoreの境界定義

- **status**: superseded
- **date**: 2026-04-17
- **superseded by**: V01-ADR-010（asset = 型定義の定義をmodelに移管）

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

### collectionの制約

`store.kind: collection` の `of:` に指定できるのは **assetのIDのみ**。
storeのネストは禁止。storeのネストが必要に見えるケースは `task` でつなぐ設計に切り出す。

### collectionのnoteクエリ設計

`collection` には `note` フィールドで自然言語のクエリ仕様を記述できる。

```yaml
- id: user_store
  type: store
  kind: collection
  of: user
  note: |
    - active_users: is_active = true のもの
    - find_by_email: email が一致するもの
```

**noteに書くべきもの（collection内で完結するもの）**
- 単一collectionのフィールドに対するフィルタ・検索
- 「言葉で書ける程度」の条件（等値・範囲・真偽フラグ等）

**taskに切り出すべきもの（noteに書かないもの）**
- 複数storeにまたがる結合
- 集計・変換・ソートを伴う複雑なクエリ
- 結果を別のassetに変換する処理

この境界は `spec/overview.md` の「structのmethodsはnoteで書ける程度のもの、それを超えたらtaskに切り出す」という原則と同じ思想を `collection` に適用したもの。

**noteの二重の役割**（ADR 008と同じ位置づけ）
- 人間にとって: クエリ仕様のdocstring
- LLMにとって: `inspect` ツールがsemantic validationの根拠として使用するcontract

noteに書かれた内容は機械的validationの対象外だが、brewprintのMCPツールがLLMに渡すことで意味的整合性のチェックに使われる。

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

## Evidence
- commit: f911107
- impl commit: n/a（V01-ADR-010によりsuperseded）
- 参考: OOPのclass vs instance区別、DDDのRepositoryパターン参考
- 注: 本ADR内の "asset"（型定義）はV01-ADR-010以降 "model" に改称済み。本ADRの記述は当時の用語のまま保存

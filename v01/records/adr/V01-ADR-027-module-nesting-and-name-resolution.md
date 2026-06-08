# V01-ADR-027: モジュールネスト許容と名前解決ルール拡張

- **status**: accepted
- **date**: 2026-04-21
- **supersedes**: V01-ADR-003

## 背景

V01-ADR-003では名前解決ルールを `モジュール名.ノード種別.ID`（3階層固定）と定めた。
これはモジュールが1階層であることを暗黙的に前提にしていた。

しかし複数サービスをモノレポで管理するユースケース（例: `auth/oauth/`、`payment/stripe/`）では、
フォルダのネストが現実に発生する。3階層固定ルールではこれに対応できない。

また、brewprintは長期的にschema-aware editorを必要とすることが確定しており、
ユーザーがフルパスを手書きする前提が崩れた（エディタが補完・生成する）。
これにより、フルパス必須でも学習コストの問題は生じないと判断した。

## 決定

### モジュールネストを許容する

フォルダ階層は何段でもよい。

```
modules/
  auth/
    task.yaml
    oauth/           ← ネストOK
      task.yaml
  payment/
    stripe/
      task.yaml
```

### 命名規則を拡張する

`モジュール名.ノード種別.ID`（3階層固定）を廃止し、以下に拡張する：

```
<モジュールパス>.<ノード種別>.<ID>
```

`<モジュールパス>` は1段以上の任意の深さのドット区切りパス。

```yaml
# 1段モジュール（従来通り）
from: auth.task.login

# 2段ネスト
from: auth.oauth.task.callback

# 3段ネスト
from: payment.stripe.webhook.task.receive
```

### ノード種別予約語をsentinelとして正式化

パーサーはノード種別予約語の出現位置でモジュールパスとIDの境界を判定する。

予約語（sentinel）一覧：

| 予約語 | ノード種別 |
|--------|-----------|
| `task` | タスク |
| `asset` | アセット |
| `store` | ストア |
| `event` | イベント |
| `state` | ステート |
| `branch` | ブランチ |
| `fork` | フォーク |
| `join` | ジョイン |
| `model` | モデル |

パース例：

```
auth.oauth.task.login
           ↑
           sentinelを発見 → 左側がモジュールパス、右側がID
→ module: auth.oauth
→ type:   task
→ id:     login
```

### フルパス必須（同モジュール内ID直書きは継承）

クロスモジュール参照はフルパス必須。V01-ADR-003の同モジュール内ID直書きルールは継続有効。

```yaml
# 同モジュール内（auth/oauth/task.yaml内での参照）
edges:
  - from: callback   # auth.oauth.task.callbackに解決される

# モジュール跨ぎ（フルパス必須）
edges:
  - from: auth.oauth.task.callback
    to: payment.stripe.task.charge
```

## 理由

### ネスト許容

大規模モノレポでは `auth/oauth/`、`payment/stripe/` などのネストが現実に発生する。
フラット固定では `auth_oauth` のような歪なフォルダ名を強制することになり、
実装コードベースとの対応が取りにくくなる。

### フルパス必須（leaf名参照を禁止）

leaf名参照（例: `oauth.task.login`）を許すと、`auth/oauth/` と `payment/oauth/` が
共存した際に曖昧性が生じる。エディタ補完が前提のため手書きコストは問題にならず、
明示性を優先してフルパス必須とする。

### sentinel方式によるパース

ノード種別はbrewprint仕様内の閉じた予約語セットであり、将来的に大幅に増えることはない。
sentinelとして使うことでモジュールパスの深さに関係なく一意にパースできる。

## 影響

- V01-ADR-003の `モジュール名.ノード種別.ID`（3階層固定）は本ADRに supersede される
- Goパーサーの名前解決ロジックを拡張する（sentinel検索方式）
- `spec/overview.md` の名前解決セクションを本ADRに基づき更新する
- クロスエッジを持つ既存の spec / uc の参照形式に変更はない（1段モジュールは従来通り動作する）

## Evidence
- commit: 75cb64b
- impl commit: tbd
- 参考: Goのimport path方式参考

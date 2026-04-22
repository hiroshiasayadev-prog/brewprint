# 033: `fk:` 値の名前解決ルール

- **status**: accepted
- **date**: 2026-04-23

## 背景

ADR-021 で `fk: <model-id>.<field-name>` という記法を定めたが、例が同モジュール内（`fk: role.id`）のみで、
クロスモジュール参照の場合の記法が未規定だった。

ADR-027 はクロスモジュール参照全般にフルパス必須を定めているが、
`fk:` の値がこのルールに従うかどうかが明示されていなかった。

UC-001 の設計中に `cart.user_id → auth.credential.username` のようなクロスモジュール FK が必要になり、
記法の確定が必要となった。

## 決定

`fk:` の値の名前解決は ADR-027 のルールに完全に準拠する。

### 同モジュール内参照

モジュールパスを省略し、`<model-id>.<field-name>` のみ書く。

```yaml
# auth/model/token.yaml 内での参照
- name: user_id
  type: str
  fk: credential.username   # auth.model.credential.username に解決される
```

### クロスモジュール参照

ADR-027 の sentinel 方式に従いフルパス必須。`model` sentinel を含む完全なパスを書く。

```yaml
# cart/model/cart.yaml 内での参照
- name: user_id
  type: str
  fk: auth.model.credential.username   # クロスモジュール → フルパス必須
```

### グローバルスコープは actor のみ

ADR-031 の通り、グローバルスコープを持つノードは actor のみ。
model を含む他の全ノード種別はモジュールスコープであるため、クロスモジュール参照には必ずフルパスが必要になる。
これは `fk:` に限らず brewprint 全体の一貫したルール。

## 理由

- ADR-021 の `fk: role.id` という例は同モジュール内前提の記述であり、クロスモジュールを排除する意図ではなかった
- ADR-027 の同モジュール内省略・クロスモジュールフルパス必須というルールを `fk:` にも適用するのが最も一貫性がある
- actor 以外にグローバルスコープのノードが存在しない以上、クロスモジュール FK にフルパスを要求するのは必然

## 影響

- `spec/nodes.md` の `fk:` フィールド説明にクロスモジュール記法の例を追記する
- ADR-021 の `fk:` 記述は「同モジュール内の例のみ」であることを注記する（supersede は不要）
- UC-001 の cart/order モデルのクロスモジュール FK は `auth.model.credential.username` 形式で記述する

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: ADR-027（sentinel方式・同モジュール省略ルール）、ADR-031（actor global定義）

# V01-ADR-039: ER図の横断view定義

- **status**: accepted
- **date**: 2026-04-24
- **depends on**: V01-ADR-017, V01-ADR-026, V01-ADR-028

## 背景

`spec/views/er.md` は「ER図はモジュール単位で描画する」と定義していた。
しかし UC-001 のように複数モジュールにまたがる DB スキーマを1枚で俯瞰したいユースケースが存在し、
「全体」の粒度をどう定義するかが未決定だった。

また V01-ADR-028 で API Table は明示的な view YAML を持つと定めており、
ER図だけが「暗黙のモジュール単位」にとどまると、view 定義の一貫性が崩れる。

## 決定

### 1. ER図にも独立した view YAML を導入する

view YAML を置くことで、モジュールを横断した ER 図を生成できる。

```yaml
as: er_diagram
id: ec_er
note: ECサイト全体のER図
modules:
  - module: auth
  - module: catalog
  - module: cart
  - module: order
  - module: payment
```

- `id`: ER図の識別子
- `note`: ER図の説明（任意）
- `modules`: 集計対象モジュールの一覧
- `modules[].module`: モジュールパス（V01-ADR-027 の名前解決ルールに従う）

### 2. view YAML なし = モジュール単位（現行動作を維持）

`render_er` を view YAML なしで呼び出した場合は、従来通りモジュール単位で描画する。
er.md の「モジュール単位」はデフォルト動作の定義として読む。

### 3. modules[] の各エントリはそのモジュール直下の store.kind: db のみを対象とする

`include_submodules` は持たない。
サブモジュールを含めたい場合は `modules[]` に明示的に列挙する。

理由: ER図はデータ層の静的スキーマ図であり、
API Table（エンドポイント集約）と異なりサブモジュールの自動巻き込みが必要なユースケースが薄い。
明示列挙のほうが「何が図に出るか」が YAMLから自明になる。

### 4. クロスモジュール FK はリレーション線として描画する

view YAML に複数モジュールが含まれる場合、
モジュールをまたぐ `fk:` もリレーション線として描画する。
（例: `cart.cart_item.item_id FK → catalog.item.id`）

view YAML に含まれないモジュールへの FK は `json` 型カラムとして表示し、リレーション線は引かない。

## 理由

- API Table（V01-ADR-028）と同じ「明示的 view YAML」パターンを採用することで、view 定義の設計を一貫させる
- 「全体」「認証系のみ」など複数粒度の ER 図を view YAML の組み合わせで表現できる
- view YAML なし動作を残すことで既存仕様（er.md）との後退互換を維持する
- `include_submodules` を持たない設計により、図の構成要素が YAML から自明になる

却下した代替案:
- `index_er.yaml` をディレクトリルートに置いて配下を自動集約 → 集約範囲がファイル配置に依存し、部分図が作りにくい
- ER図はモジュール単位のみ維持し UC-001 README は複数図を並べる → クロスモジュール FK が1枚で見えない。UC の検証意図を満たせない

## 影響

- `spec/views/er.md` の「renderスコープ」セクションを更新し、view YAML による横断描画を追記する
- UC-001 の `views/er.yaml` を追加し、`### ER Diagram: ec全体` の render 例として使用する
- `render_er` MCP ツールは view YAML の有無を見て動作を切り替える（実装 tbd）

## Evidence
- commit: 42472a7
- impl commit: tbd
- 参考: V01-ADR-028（API Table の view YAML パターン）

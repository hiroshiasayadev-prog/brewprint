# 045: render_index.yaml スキーマ

- **status**: accepted
- **date**: 2026-04-26
- **supersedes**: なし

## 背景

ADR-043でbrewprintプロジェクトのルートに `render_index.yaml` を置き、`renders/` のgroup構成を制御することを決定した。
本ADRでは `render_index.yaml` のスキーマ詳細を規定する。

規定が必要な論点は以下の通り。

- group idの命名規則
- 1 moduleが複数groupに属することを許可するか
- どのgroupにも属さないuncovered moduleの扱い
- `groups` 配列の順序とrender出力の表示順の関係
- ADR-027のmodule nestingとの整合（ネストしたmoduleを指定したときの子moduleの扱い）

## 決定

### 1. トップレベル構造

```yaml
# render_index.yaml
groups:
  - id: auth
    label: 認証
    modules: [auth]
  - id: commerce
    label: 商取引
    modules: [cart, order, payment]
  - id: catalog
    label: カタログ
    modules: [catalog, inventory]
```

| フィールド | 型 | 必須 | 説明 |
|---|---|---|---|
| `groups` | array | 必須 | groupの定義リスト。空配列はvalidation error |

### 2. group オブジェクト

| フィールド | 型 | 必須 | 説明 |
|---|---|---|---|
| `id` | string | 必須 | group識別子。`renders/{id}/` ディレクトリ名になる |
| `label` | string | 任意 | 人間向け表示名。省略時は `id` をそのまま使う |
| `modules` | array[string] | 必須 | このgroupに属するmodule名のリスト。空配列はvalidation error |

### 3. group id の命名規則

- 使用可能文字: `[a-z0-9_]`（小文字英数字とアンダースコア）
- アンダースコア始まり（`_` prefix）は **validation error**。`_cross/` との混同を防ぐため予約する
- 空文字はvalidation error
- ※ module名も同じく `_` 始まり禁止とする。暗黙group idがmodule名をそのまま使うため、module名がアンダースコア始まりだとgroup id命名規則と衝突する。詳細は理由セクション参照。

```yaml
# OK
- id: commerce
- id: payment_webhook

# NG（validation error）
- id: _internal
- id: Commerce
- id: ""
```

### 4. module 重複の禁止

1つのmoduleは最大1つのgroupにしか属せない。複数のgroupに同じmodule名が現れた場合はvalidation error。

```yaml
# NG（orderが2つのgroupに属している）
groups:
  - id: commerce
    modules: [cart, order]
  - id: fulfillment
    modules: [order, inventory]
```

### 5. uncovered module の扱い

`yaml/` 配下に存在するmoduleが `render_index.yaml` のどのgroupにも属さない場合、そのmoduleは **group id = module名の暗黙groupとして扱う**。

```
# yaml/配下にmoduleが [auth, cart, order, payment, catalog, inventory] あるとする
# render_index.yaml で以下を定義した場合:
groups:
  - id: commerce
    modules: [cart, order, payment]

# 結果:
# renders/commerce/   ← 明示group
# renders/auth/       ← 暗黙group（auth module単独）
# renders/catalog/    ← 暗黙group（catalog module単独）
# renders/inventory/  ← 暗黙group（inventory module単独）
```

warningは出力するが、エラーにはしない。意図的に一部だけgroupingすることを許容する。

### 6. groups 配列の順序

`groups` 配列の定義順が `renders/index.md` のテーブル行順および各group `index.md` の表示順と一致する。
暗黙groupは明示groupの後にアルファベット順で追加される。

### 7. ネストしたmoduleの扱い（ADR-027との整合）

ADR-027のmodule nestingに基づき、`modules` フィールドに指定するのは **最上位module名のみ**とする。
`payment/webhooks` のようなスラッシュ区切りパスはvalidation errorとする。

ネストしたmodule（例: `payment/webhooks`）を親module（`payment`）でgroupに含めた場合、**子moduleも同groupに含まれる**。

```yaml
# payment を指定した場合、payment/webhooks も commerce groupに含まれる
groups:
  - id: commerce
    modules: [cart, order, payment]
```

子moduleだけを別groupに分離することはできない。親moduleを分離したうえで子moduleを別途指定することも禁止する（module重複禁止ルールに抵触するため）。

### 8. ネストしたmodule内taskのrender出力ファイル名

ネストしたmodule内のtask（例: `payment/webhooks` 内のtask）のrender出力ファイル命名ルールは、Go実装ADRで規定する。

## 理由

### 暗黙group idとアンダースコア禁止ルールの整合

ADR-043で「group idにアンダースコア始まりはvalidation error」と定めている。
ADR-027はmodule名の命名規則を規定していないが、アンダースコア始まりのmodule名を使うと暗黙group idが同じくアンダースコア始まりになり、validation errorと衝突する。
そのため **module名もアンダースコア始まりを禁止する**。アンダースコア始まりのmodule名が `yaml/` 配下に存在した場合、そのmoduleのrenderはスキップしwarningを出力する。

### module重複を禁止する理由

同じmoduleが複数groupに属すると、renderの配置が一意に定まらず、`renders/` の構造がambiguousになる。
1 module = 1 group の制約によりrender出力の決定性を保証する。

### uncovered moduleをwarningにとどめる理由

大規模プロジェクトでは `render_index.yaml` の更新が実態に遅れることが多い。
新moduleを追加するたびに `render_index.yaml` の更新を強制するとDXが悪化する。
暗黙groupとして動作させることで、定義漏れでもrenderが壊れない安全性を確保する。

### 順序を配列順と一致させる理由

`renders/index.md` の表示順が毎回変わると、diffが読みにくくなる。
YAMLの配列順という明示的な定義に従うことで決定性を持たせる。

### ネストしたmoduleを親で一括指定とする理由

ADR-027のmodule nestingは名前解決のためのルールであり、groupingはそれに依存する形で動く。
子moduleだけを別groupに分離できると、module nestingの構造と `renders/` の構造が乖離し、追跡が困難になる。親module単位でgroupingを制御することで、YAMLのディレクトリ構造と `renders/` 構造の対応を保つ。

## 影響

- Go rendererはYAML parseフェーズで `render_index.yaml` をvalidationし、上記エラー条件を検出したらbuildを中止する。
- uncovered moduleおよびアンダースコア始まりmodule名はwarningとしてstderrに出力し、そのmoduleのrenderをスキップする。
- ネストしたmodule内のtask render出力ファイル命名はGo実装ADRで規定する。
- `docs/spec/` への追記は不要（render_index.yamlはbrewprint YAML仕様の対象外のツール設定ファイル）。

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: 特になし

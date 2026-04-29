---
scope: docs/spec/project-layout.md
status: draft
last_updated: 2026-04-30
summary: >
  brewprint プロジェクトのディレクトリ構造を定義する。
  yaml/ / renders/ / render_index.yaml の配置、render_index.yaml のスキーマ、
  renders/ の出力構造、render output filenameと衝突時の扱いを含む。
depends_on:
  - docs/adr/043-project-root-layout-and-render-output.md
  - docs/adr/045-render-index-schema.md
  - docs/adr/046-render-output-placement-for-state-sequence-wireframe-preview.md
  - docs/adr/053-render-output-filename-policy-for-nested-modules.md
---

# プロジェクトレイアウト仕様

## 1. プロジェクトルート構造

brewprint プロジェクトは以下のディレクトリ構造を持つ。

```
{project-root}/
  yaml/               ← brewprint YAML群（single source of truth）
  renders/            ← Go rendererによる自動生成物
  render_index.yaml   ← render出力のグルーピング設定（人間が書く、任意）
  README.md
```

- `yaml/` がプロジェクトの YAML ソースを保持する単一ディレクトリ
- `renders/` は Go renderer が `brewprint render` で生成する成果物。人間は直接編集しない（golden test 目的での gitcommit は許容するが、編集権限は renderer のみ）
- `render_index.yaml` は省略可能。省略時は「1 module = 1 group」のデフォルト動作

`master.yaml` のようなプロジェクトルート台帳ファイルは存在しない。プロジェクトのモジュール構成は `yaml/` 配下のディレクトリ構造から導出される。

> 由来: ADR-043 §1。`master.yaml` 不在は ADR-002 §決定 で当初定めた `master.yaml` を ADR-043 が覆した結果。

## 2. yaml/ の構造

`yaml/` 配下のディレクトリがそのままモジュール階層となる。詳細は [naming.md](./naming.md) §モジュールとフォルダ階層 を参照。

`yaml/` の親ディレクトリがプロジェクトルートとなる。loader は `--yaml-root` で指定されたディレクトリの basename が `yaml` の場合、その親に `render_index.yaml` を探す。

> 由来: ADR-002 §決定, ADR-043 §1

## 3. renders/ の出力構造

```
renders/
  index.md                  ← masterインデックス（全groupへのリンクテーブル）
  {group-id}/
    index.md                ← groupインデックス（group内の全render一覧テーブル）
    dag-{task-id}.md
    state-{fsm-id}.md
    seq-{scenario-id}.md
    wireframe-{fsm-id}-{state-id}.html
  _cross/
    er.md
    api.md
  _preview/
    wireframe.html
```

### group ディレクトリ

通常 group は `{group-id}/` ディレクトリに配置される。`group-id` の命名規則は本spec §4 を参照。

各 group ディレクトリ直下に以下の render が配置される（存在するもののみ）:

- `dag-{task-id}.md` — task の DAG render
- `state-{fsm-id}.md` — FSM の state diagram render
- `seq-{scenario-id}.md` — sequence diagram シナリオ render
- `wireframe-{fsm-id}-{state-id}.html` — wireframe render

v1では、nested module の module path は出力ファイル名に含めない。render 出力ファイル名は semantic object の local ID から決定する。

例:

```text
yaml/payment/webhooks/task/process_payment.yaml
→ renders/commerce/dag-process_payment.md
```

同一 group 内で複数の render output が同一 relative path に解決される場合、renderer / placement validation は error として停止する。silent overwrite は禁止する。

> 由来: ADR-053 §決定

### `_cross/` — module横断 view

ER / API Table のような複数モジュールにまたがる view は `_cross/` に配置する。アンダースコアプレフィックスは通常の group id と区別するための予約。

将来の cross-cutting view（global event flow 等）も同ディレクトリ配下に追加する。

### `_preview/` — preview harness

wireframe の preview ハーネスは `_preview/wireframe.html` に配置する。

> 由来: ADR-043 §2, ADR-046

## 4. render_index.yaml スキーマ

### 4.1 トップレベル構造

```yaml
groups:
  - id: auth
    label: 認証
    modules: [auth]
  - id: commerce
    label: 商取引
    modules: [cart, order, payment]
```

| フィールド | 型 | 必須 | 説明 |
|---|---|---|---|
| `groups` | array | 必須 | group の定義リスト。空配列は validation error |

### 4.2 group オブジェクト

| フィールド | 型 | 必須 | 説明 |
|---|---|---|---|
| `id` | string | 必須 | group 識別子。`renders/{id}/` ディレクトリ名になる |
| `label` | string | 任意 | 人間向け表示名。省略時は `id` をそのまま使う |
| `modules` | array[string] | 必須 | この group に属する module 名のリスト。空配列は validation error |

### 4.3 group id / module 名の命名規則

- 使用可能文字: `[a-z0-9_]`（小文字英数字とアンダースコア）
- アンダースコア始まり（`_` prefix）は **validation error**。`_cross/` / `_preview/` との混同を防ぐため予約
- 空文字は validation error
- module 名も同じく `_` 始まり禁止（暗黙 group id が module 名をそのまま使うため）

```yaml
# OK
- id: commerce
- id: payment_webhook

# NG（validation error）
- id: _internal
- id: Commerce
- id: ""
```

### 4.4 module 重複の禁止

1つの module は最大1つの group にしか属せない。複数 group に同じ module 名が現れた場合は validation error。

### 4.5 uncovered module の扱い

`yaml/` 配下に存在する module が `render_index.yaml` のどの group にも属さない場合、そのmodule は **group id = module 名の暗黙 group として扱う**。

明示 group の後にアルファベット順で追加される。warning は出力するが、エラーにはしない。

### 4.6 groups 配列の順序

`groups` 配列の定義順が `renders/index.md` のテーブル行順および各 group `index.md` の表示順と一致する。
暗黙 group は明示 group の後にアルファベット順で追加される。

### 4.7 ネスト module の扱い

`modules` フィールドに指定するのは **最上位 module 名のみ**。スラッシュ区切りパス（`payment/webhooks` 等）は validation error。

ネストした module（例: `payment/webhooks`）を親 module（`payment`）で group に含めた場合、**子 module も同 group に含まれる**。子 module だけを別 group に分離することはできない。

> 由来: ADR-045 §1〜§7

## 5. master index.md フォーマット

```markdown
# {project-name} render index

| group | DAG | State | Sequence | Wireframe | ER | API |
|---|---|---|---|---|---|---|
| [認証](auth/index.md) | 1 | 1 | - | 2 | - | - |
| [商取引](commerce/index.md) | 3 | 1 | 2 | 2 | - | - |
| *(cross)* | - | - | - | - | [er](_cross/er.md) | [api](_cross/api.md) |
| *(preview)* | - | - | - | [wireframe preview](_preview/wireframe.html) | - | - |
```

group 列の表示ルール:
- 明示 group: `label`（省略時は `id`）
- 暗黙 group: `id`（= module 名）
- cross 行: 固定で `*(cross)*`
- preview 行: 固定で `*(preview)*`

master `index.md` は通常 group へのリンクのみ。個別 render への直リンクは group `index.md` が提供する。`_cross/` と `_preview/` は特殊 render ディレクトリのため、master `index.md` から `*(cross)*` / `*(preview)*` 行として直接リンクしてよい。

> 由来: ADR-043 §4

## 6. CLI

`brewprint render --yaml-root <path> --out <path> [--clean]` で renders/ を生成する。
詳細は [mcp.md](./mcp.md) および実装側の go-mN-summary.md を参照。

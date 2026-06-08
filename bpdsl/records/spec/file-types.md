---
scope: docs/spec/file-types.md
status: draft
last_updated: 2026-04-29
summary: >
  brewprint YAMLファイルの種別分類ルールを定義する。
  `as:` / `nodes:` / ファイル名による判定アルゴリズム、`as:`値一覧、
  Unsupportedファイルの扱い、ファイルパスの正規化規約を含む。
depends_on:
  - docs/adr/028-api-table-route-composition.md
  - docs/adr/030-yaml-file-type-declaration.md
  - docs/adr/032-sequence-diagram-scenario-schema.md
  - docs/adr/039-er-diagram-composed-view.md
  - docs/adr/043-project-root-layout-and-render-output.md
  - docs/adr/051-unsupported-yaml-file-warning.md
  - docs/adr/052-source-file-path-normalization.md
---

# ファイル種別仕様

## 1. ファイル種別

brewprint プロジェクトの `yaml/` 配下と直下の `render_index.yaml` は、loader によって以下のいずれかに分類される。

| FileKind | 内容 | トップレベルキーの特徴 |
|---|---|---|
| `node` | ノード定義ファイル | `nodes:` を持つ |
| `view` | View 定義ファイル | `as:` を持つ |
| `render_index` | プロジェクトルート直下の `render_index.yaml` | ファイル名で識別 |
| `unsupported` | 上記いずれにも該当しないファイル | — |

> 由来: V01-ADR-030 §決定, V01-ADR-043 §1

## 2. 分類アルゴリズム

`*.yaml` / `*.yml` ファイルは以下の順で判定する。

1. **ファイル名**: ファイル名が `render_index.yaml` であれば `render_index`
2. **`as:` キー**: トップレベルマッピングに `as:` が存在すれば `view`（`viewAs` フィールドに値を保持）
3. **`nodes:` キー**: トップレベルマッピングに `nodes:` が存在すれば `node`
4. それ以外は `unsupported`

判定対象は `*.yaml` および `*.yml` の拡張子を持つファイルのみ。それ以外の拡張子は loader が無視する。

`render_index.yaml` 以外のファイル名は分類判定に影響しない（V01-ADR-002 / V01-ADR-030 の方針: 任意のファイル名を許容する）。

> 由来: V01-ADR-030 §決定, V01-ADR-043 §1

## 3. `as:` 値一覧

| `as:` 値 | 対応する view | 由来 |
|---|---|---|
| `api_table` | API Table view | V01-ADR-028 |
| `sequence_diagram` | Sequence Diagram シナリオ | V01-ADR-032 |
| `er_diagram` | ER Diagram view（モジュール横断） | V01-ADR-039 |

新たな view 種別を追加する際は本表に追記する。

### Wireframe について

wireframe はファイルレベルの view を持たない。state ノード内に DSL として埋め込まれる形でのみ存在する（V01-ADR-029, V01-ADR-042）。
そのため `as: wireframe` という値は存在せず、本表に含めない。

> 由来: V01-ADR-030 §決定, V01-ADR-039 §1, V01-ADR-029, V01-ADR-042

## 4. Unsupported ファイルの扱い

`unsupported` に分類されたファイルは silent に通さず、`unsupported_file` という warning diagnostic を出力する。validation 自体は成功扱い（warning は exit code に影響しない）。

これは「`as:` を書き忘れた view YAML が静かに消える」事故を防ぐための方針。silent skip にすると、view ファイルとして書いたつもりが種別不明として無視され、render 出力から欠落しても気づけない。

警告を出した上でファイル自体は処理対象外とする。後段の resolve / render は当該ファイルを参照しない。

> 由来: V01-ADR-051

## 5. ファイルパス正規化

`source.File.Path` フィールドはプロジェクトルート相対のスラッシュ正規化形式で保持する。

- 基点: プロジェクトルート（`yaml/` の親ディレクトリ）
- 区切り: `/`（OS ネイティブセパレータは使わない）
- 例: `yaml/auth/dag.yaml`, `render_index.yaml`

絶対パスや OS 依存セパレータ（`\`）を含む形式は保持しない。loader が `filepath.ToSlash` 等で正規化する。

この規約により、golden test の決定性、render 出力の再現性、ログ・diagnostic 表示の一貫性を担保する。

> 由来: V01-ADR-052

## 6. 実装参照

ファイル分類は `internal/source/classify.go` の `classifyFile` 関数で行われる。
loader は `internal/source/loader.go` の `Loader.Load` で `yaml/` 配下を walk し、追加で `yaml/` の親ディレクトリ直下の `render_index.yaml` を探索する。

# V01-ADR-051: Unsupported YAMLファイルへのwarning出力

- **status**: accepted
- **date**: 2026-04-29
- **depends on**: V01-ADR-030, V01-ADR-050

## 背景

source層のloader（`internal/source/loader.go` の `Loader.Load`）は `yaml/` 配下を walk し、各 `*.yaml` ファイルを `classifyFile`（`internal/source/classify.go`）で4種別（`node` / `view` / `render_index` / `unsupported`）に分類する。

現状、`unsupported` に分類されたファイル（`as:` も `nodes:` も持たず、`render_index.yaml` でもないファイル）は silent にスキップされる。後段の resolve / render は当該ファイルを参照しないだけで、警告も診断も出ない。

これには以下の問題がある:

- view YAML を新規作成した際に `as:` を書き忘れると、ファイル全体が無言で無視される
- typo（`as:` を `is:` 等と書いてしまう）も同様に検出されない
- 大規模プロジェクトでは「render に出ない」事象から原因ファイルを逆引きするのが困難

`as:` を必須としたV01-ADR-030の方針（view定義は明示的にファイル種別を宣言する）を実効性あるものにするには、宣言が欠落している事象を検出して通知する仕組みが必要。

## 決定

`unsupported` に分類されたファイルに対して `unsupported_file` という diagnostic を warning severity で発行する。

- severity: `warning`
- code: `unsupported_file`
- 発行タイミング: source層の loader フェーズ（classify 直後）
- validation 全体は成功扱い（warning は exit code に影響しない）

警告を出した上でファイル自体は処理対象外とする。後段の resolve / render は当該ファイルを参照しない（現状動作を維持）。

## 理由

### error ではなく warning にする理由

- 既存プロジェクトに新規ファイル形式が混入している過渡的状態でも build を止めない
- 「単に分類できないファイル」と「明確な不正」は区別すべき。`unsupported` は前者
- 後段で実質害がない（resolve / render は当該ファイルを無視するため整合性は保たれる）

### error にしない代替案を却下した理由

`as:` 書き忘れは error にすべき、という判断もありうるが、

- loader 段階では「as: を書き忘れた view YAML」と「単なるメモ用 YAML」を区別できない
- 区別できない以上、より厳しい severity（error）を一律適用すると過剰検知になる

将来的に「`yaml/` 配下にある `*.yaml` のうち分類不能なものは error」という別ルールを導入する余地は残すが、本ADRでは warning に留める。

### silent skip を却下した理由

- view YAML の `as:` 書き忘れに気づけない（背景セクション参照）
- spec-first 方針（V01-ADR-050）として、loader の振る舞いは外部仕様に明示されるべき。silent skip は仕様化しても利用者にとって価値がない

## 影響

- `docs/spec/file-types.md` §4 にUnsupported ファイルの扱いを記載
- `docs/spec/diagnostics.md` の diagnostic code 一覧に `unsupported_file` (warning) を追加
- 実装は `internal/source` 層に warning diagnostic を出す経路を追加する必要がある（M1 時点では diagnostic 配信パスが未整備のため、impl側で別途扱う）

## Evidence
- commit: 261ed60
- impl commit: tbd
- 参考: V01-ADR-030 §決定, V01-ADR-050 §決定

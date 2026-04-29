# 052: source.File.Path のプロジェクトルート相対 + slash正規化

- **status**: accepted
- **date**: 2026-04-29
- **depends on**: ADR-043, ADR-050

## 背景

source層の `rawyaml.File` 構造体には `ID` と `Path` の2フィールドがある。

- `ID`: プロジェクトルート相対のスラッシュ正規化文字列。`normalizeFileID`（`internal/source/file.go`）で生成
- `Path`: ファイル読み込みに使用した絶対パス（`filepath.WalkDir` および `filepath.Join` の生の出力）

現状の `Path` は OS ネイティブセパレータを含む絶対パス（Windows なら `C:\Users\imved\projects\brewprint\yaml\auth\dag.yaml` 形式、Unix なら `/path/to/yaml/auth/dag.yaml` 形式）になっている。

これには以下の問題がある:

- **golden test の決定性が崩れる**: `Path` が出力に含まれる test（goldenとして固定したいdiagnostic等）でOS依存の差分が出る
- **render 出力の再現性に影響**: 将来 diagnostic に file path を出す際にOS依存のpathが見える
- **ログ表示の一貫性が崩れる**: Windows と Unix で diagnostic 表示が変わる

`ID` は既にスラッシュ正規化されているため、`Path` も同様の規約に揃えるのが自然。

## 決定

`rawyaml.File.Path` をプロジェクトルート相対のスラッシュ正規化形式で保持する。

- 基点: プロジェクトルート（`yaml/` の親ディレクトリ）
- 区切り: `/`（OS ネイティブセパレータは使わない）
- 例: `yaml/auth/dag.yaml`, `render_index.yaml`

絶対パスや OS 依存セパレータ（`\`）を含む形式は保持しない。loader が `filepath.Rel` + `filepath.ToSlash` で正規化する。

`ID` と `Path` の関係:

- `yaml/` 配下のファイル: `Path` = `yaml/` を含む相対パス、`ID` = `yaml/` を含まない相対パス
- `render_index.yaml`: `Path` = `render_index.yaml`、`ID` = `render_index.yaml`（同じ）

両者の差分は「`yaml/` プレフィックスを含むかどうか」。`ID` はモジュール解決のキーとして使われるため `yaml/` を剥がした形になっており、`Path` はファイルシステム上の位置を表すため `yaml/` を含む。

## 理由

### プロジェクトルート相対にする理由

- ADR-043 でプロジェクトのルートが定義されており、それを基点にすることで「ファイル位置」を一意に特定できる
- 絶対パスは OS / 環境依存であり、再現性のあるテストやログに不適

### slash 正規化する理由

- 既存の `normalizeFileID`（`ID` の正規化）と同じ思想で一貫
- Windows / Unix 双方で同一の文字列表現になるため golden test が安定する
- YAML / Markdown のリンクや diagnostic 表示でそのまま使える

### `ID` と `Path` を分けたまま残す理由

- `ID` は名前解決キー（モジュール = `yaml/` 配下のディレクトリ階層）として使われるため、`yaml/` プレフィックスを剥がす必要がある
- `Path` はファイルシステム上の位置として使われるため、プロジェクトルートからの完全な相対パスを保つべき
- 両者を1つのフィールドに統合すると、用途の混同が起きる

却下した代替案:

- **絶対パスを保持**: OS依存・環境依存になり再現性が損なわれる
- **`Path` を廃止して `ID` だけにする**: `render_index.yaml` のように `yaml/` 外のファイルがあるため、ID とパスを統合できない

## 影響

- `docs/spec/file-types.md` §5 に File.Path 正規化規約を記載
- 実装は `internal/source/loader.go` の `loadFile` 呼び出し前後で `Path` を正規化する経路を追加する必要がある
- `internal/source/file.go` に `normalizeFilePath`（仮称）相当のヘルパーが必要になる可能性がある（実装側で命名・関数分離は判断）
- 既存の golden test に影響がある場合は更新する

## Evidence
- commit: tbd
- impl commit: tbd
- 参考: ADR-043 §1, ADR-050 §決定

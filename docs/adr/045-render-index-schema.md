# 045: render_index.yaml スキーマ

- **status**: accepted
- **date**: 2026-04-26
- **migrated_to_spec**: 2026-04-29
- **supersedes**:

> このADRの現行仕様詳細は [docs/spec/project-layout.md](../spec/project-layout.md) §4 を参照。

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

### 決定の概要

1. **トップレベル構造**: `groups` 配列（仕様詳細: [project-layout.md](../spec/project-layout.md) §4.1）
2. **group オブジェクト**: `id` / `label` / `modules`（仕様詳細: [project-layout.md](../spec/project-layout.md) §4.2）
3. **id / module 名命名規則**: `[a-z0-9_]`、`_` 始まり禁止（仕様詳細: [project-layout.md](../spec/project-layout.md) §4.3）
4. **module 重複の禁止**: 1 module = 1 group（仕様詳細: [project-layout.md](../spec/project-layout.md) §4.4）
5. **uncovered module の扱い**: 暗黙 group として扱う、warning 出力（仕様詳細: [project-layout.md](../spec/project-layout.md) §4.5）
6. **groups 配列の順序**: 定義順 = 表示順（仕様詳細: [project-layout.md](../spec/project-layout.md) §4.6）
7. **ネスト module の扱い**: 最上位 module 名のみ指定、子 module は親に従属（仕様詳細: [project-layout.md](../spec/project-layout.md) §4.7）
8. **ネスト module 内 task の render 出力ファイル名**: Go 実装 ADR で規定（未確定）

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
- commit: 3200fa2
- impl commit: tbd
- 参考: 特になし

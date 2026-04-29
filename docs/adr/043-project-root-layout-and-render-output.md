# 043: プロジェクトルートレイアウトとrender出力構造

- **status**: accepted
- **date**: 2026-04-26
- **migrated_to_spec**: 2026-04-29
- **supersedes**: なし

> このADRの現行仕様詳細は [docs/spec/project-layout.md](../spec/project-layout.md) §1, §3, §5 を参照。

## 背景

UC-001（EC Checkout Flow）の作業を通じて、brewprintの「プロジェクト」は1ディレクトリで完結する単位として自然に機能することが確認された。
現状では `docs/uc/001-ec-checkout-flow/` 配下に `yaml/` と `docs/` を持つ構造になっているが、以下の問題がある。

1. **render-*.md が手書き**
   - `docs/render-dag.md` 等はGo rendererが自動生成すべき成果物であるにもかかわらず、現状は人間が手書きしたgolden fixtureである。
   - 「手書きのドキュメント」と「Goの生成物」が同じ `docs/` ディレクトリに混在することで、どちらが正となるか曖昧になる。

2. **render出力の粒度が固定されている**
   - `render-dag.md` にプロジェクト内の全DAG renderが1ファイルに集約されているため、システムが大規模になると追跡不能になる。
   - 一方で、1 task = 1 ファイルに固定すると、小規模プロジェクトでは逆に分散しすぎる。

3. **グルーピングをユーザーが制御できない**
   - module単位に固定するとdomain概念を表現できない。任意のグルーピング設定が必要。

4. **render_index.yaml の配置が未定**
   - グルーピング設定ファイルをどこに置くか、プロジェクト構造の中での位置づけが規定されていない。

## 決定

### スコープの明示

本ADRは **brewprintプロジェクト本体の構造**（`yaml/` / `renders/` / `render_index.yaml` / `README.md`）のみを規定する。

UC運用上付加されるファイル（`HANDOFF.md` / `TASKS-*.md` / `docs/coverage.md` 等）はbrewprintプロジェクト仕様の対象外であり、`doc-policy.md` 側の責務とする。

### 決定の概要

1. **プロジェクトルート構造**: `yaml/` / `renders/` / `render_index.yaml` / `README.md`（仕様詳細: [project-layout.md](../spec/project-layout.md) §1）
2. **renders/ の出力構造**: `index.md` / `{group-id}/` / `_cross/` / `_preview/`（仕様詳細: [project-layout.md](../spec/project-layout.md) §3）
3. **render_index.yaml によるグルーピング制御**: スキーマ詳細はADR-045（仕様詳細: [project-layout.md](../spec/project-layout.md) §4）
4. **master index.md フォーマット**: group一覧テーブル形式（仕様詳細: [project-layout.md](../spec/project-layout.md) §5）
5. **CLIインターフェース**: `brewprint render ./yaml/ --out ./renders/`（実装ADR側で確定）

`renders/` はGoが生成する成果物ディレクトリであり、人間が直接編集しない。
goldenテスト目的でgitにcommitすることは許容するが、その場合も編集権限はGo rendererのみとする。

State / Sequence / Wireframe / Preview の詳細な配置規則はADR-046で規定する。
`render_index.yaml` のスキーマ詳細はADR-045で規定する。

## 理由

### renders/ を docs/ と分ける理由

`docs/` という名称は「人間が書いたドキュメント」の慣習的な配置先であり、生成物を置くと混乱を招く。
`renders/` はYAMLをrenderした成果物であることを名称から明示でき、brewprintの「YAMLがsingle source of truth」という方針と一貫している。

また、生成物ディレクトリを分離することで `.gitignore` への追加・除外判断もしやすくなる。

### グルーピングをconfigで制御する理由

module固定では小規模で過不足なく、大規模では粗すぎるという問題が起きる。
domain / bounded context / チーム構成など、グルーピングの基準はプロジェクトによって異なるため、`render_index.yaml` でユーザーが制御できるようにする。

デフォルト「1 module = 1 group」は設定なしでも動作する最小構成を保証する。

### _cross/ を分ける理由

ER / API Tableはプロジェクト横断的なviewであり、特定のgroupに属させると配置が恣意的になる。
アンダースコアプレフィックスで通常のgroup IDと区別する。将来のcross-cutting view（global event flow等）も同ディレクトリに収める。

### group index.md を設ける理由

masterの `renders/index.md` が全renderへの直リンクを持つと、大規模プロジェクトでは1ファイルが肥大化して追跡不能になる。
master → group index → 個別renderの2段構造にすることで、スケーラビリティを確保する。

### modules フィールドとADR-010 / ADR-027との整合

`groups[].modules` はADR-010が定義する「モジュール = ディレクトリ」と同義のmodule名を指定する。
ADR-027のmodule nestingとの整合（ネストしたmoduleを指定したとき子moduleを含むかどうか等）はADR-045で規定する。

## 影響

- `docs/uc/001-ec-checkout-flow/` の既存 `docs/render-*.md` を、本ADRで決定した `renders/` 構造に合わせてgolden fixtureとして再配置する（`docs/TASKS.md` に追記）。
- `docs/doc-policy.md` のuc運用セクション（section 5）を本構造に合わせて更新する（完了済み）。
- `render_index.yaml` のスキーマ詳細はADR-045で規定する。
- Go renderer実装時は本ADRのレイアウト方針に従う。
- CLIインターフェースはGo実装ADRで規定する。

## Evidence
- commit: 3200fa2
- impl commit: tbd
- 参考: 特になし

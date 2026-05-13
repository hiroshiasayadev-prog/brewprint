# brewprint ドキュメント運用方針

> このdocはClaude（AIアシスタント）との協働における、ドキュメント管理の方針を定める。
> Claudeへの指示も兼ねているため、別会話のClaudeはこのdocを最初に読むこと。

---

### ⛔ ファイル操作の禁止事項（最重要）

Claudeのサンドボックス環境（bash_tool / create_file 等）は揮発性であり、
セッション終了と同時に消える。ここへの書き込みはユーザーのリポジトリに届かない。
**ファイル操作はすべてMCPツール経由でローカルリポジトリに対して行うこと。例外なし。**

---

## 1. プロジェクト概要

brewprintは**人間とLLMの共通設計言語**。

- 人間向け → Mermaid図（md形式でrender）
- LLM向け → signature / dep tree / inspect（MCP経由）
- YAMLはその裏側にある中間表現にすぎない

実装はGoで行う。YAMLのASTをGoで保持し、MCPツールとしてLLMに公開する。

ドキュメント運用は **spec-first** で行う（ADR-050）。
- **spec** が「現行仕様の唯一の正」。実装・LLM・読者はspecだけを読めば現状を把握できる
- **ADR** は「設計判断の根拠記録」。なぜその仕様になったか、何を却下したか、トレードオフは何だったかを記録する
- ADRは起票時点での仕様案・規定を含んでよいが、それは**起票時点のスナップショット**であり、後続ADRやspec更新で覆されうる前提で書かれる

---

## 2. ドキュメント構成

```
docs/
  doc-policy.md       ← このファイル。Claudeはセッション開始時に必ず読む
  spec/               ← 現行仕様の唯一の正（言語仕様・スキーマ定義・MCPインターフェース）
  adr/                ← Architecture Decision Records（設計判断の根拠記録）
  uc/                 ← ユースケース集（実例YAML + 期待するrender結果）
  impl/               ← 実装作業の引継ぎ・レビューメモ
  TASKS.md            ← milestone index。詳細taskは tasks/ を参照
  tasks/              ← milestone別task file
    mXX-*.md
```

`docs/TASKS.md` は milestone index として運用する。
詳細taskは `docs/tasks/mXX-*.md` に置き、セッション開始時は `docs/TASKS.md` のみ読む。
作業対象 milestone が決まってから該当 task file を読み、closed milestone は原則読まない。

### 現状のspec構成

```
docs/spec/
  overview.md            ← brewprintとは何か、全体像
  project-layout.md      ← プロジェクトディレクトリ構造、yaml/、renders/、render_index.yaml
  file-types.md          ← ファイル分類、as:、nodes:、Unsupported扱い
  naming.md              ← QualifiedID、名前解決、actor global、module nesting、FK解決
  nodes.md               ← ノード種別ごとのスキーマ
  edges.md               ← flow:、transitions:、reads/writes、$シジル
  diagnostics.md         ← validation diagnostic
  mcp.md                 ← MCP query layer
  views/
    dag.md
    er.md
    state-diagram.md
    sequence-diagram.md
    api-table.md
    wireframe.md
```

`project-layout.md` / `file-types.md` / `naming.md` は ADR-050で新設が決まった spec。
漸進移行ルール（§7）に従い、関連ADRに触れたタイミングで作成する。

### Design Records MCP

ADR / spec の record 検索・検証・取得には Design Records MCP を利用できる。
tool の request / response 仕様は `docs/spec/design-records-mcp/tools.md` を参照する。

---

## 3. ADR運用

### ADRが残る理由

spec-firstに転換した今でも、ADRは残り続ける。
ADRは「**なぜそう決まったか**」を残すドキュメントであり、specには書かれない設計判断のトレードオフ・却下案・歴史を記録する。

- 別会話のClaudeが「これ変えていい？」を判断する際、根拠ベースで議論できる
- 過去の決定を覆す場合、何が動機だったかを辿れる
- 設計判断の積み重ねがbrewprintの一貫性を支える

### ADRとspecの責務境界

ADRとspecは**時間軸の違い**で役割を分担する。

| ドキュメント | 何を書くか | 役割 |
|---|---|---|
| spec | **現在の仕様**。スキーマ、ルール、フィールド定義、構造、振る舞い | 現行の唯一の正 |
| ADR | **起票時点の決定**: 背景、選択肢、却下理由、トレードオフ、起票時点での仕様案・規定 | 過去の判断の凍結された記録 |

現状を知りたいときは**specを参照**する。なぜそうなったかを辿りたいときは**ADRを参照**する。

ADRは起票時点での仕様案・規定を持ってよい。ただし**ADR本文の仕様記述は起票時点のスナップショット**であり、後続ADRやspec更新で覆されうる前提で書かれる。
ADRの記述は**遡及修正しない**。後続ADR・spec更新で覆された場合も、旧ADRには起票時点の記述を残す。supersededラベルや`migrated_to_spec`メタ情報で関係性を示す。

ADRに書いてよい仕様記述の例（詳細はADR-050 §決定 §2を参照）:

- 起票時点での決定の核（例: 「`as:` キーでファイルを分類する」）
- 決定の意味を伝えるための具体例
- 却下案を説明するための仕様候補比較
- ADRのスコープ内で完結する詳細仕様

ADR起票・レビュー時の実践ルール、実例由来ADRの書き方、ADR / spec / task / UC docs の責務境界は `docs/adr-authoring-guide.md` を参照する。

### ファイル名規則

```
docs/adr/
  NNN-タイトル.md   （例: 001-node-type-splitting.md）
```

NNNは3桁ゼロ埋めの連番。

### ADRのフォーマット

```markdown
# NNN: タイトル

- **status**: proposed / accepted / superseded
- **date**: YYYY-MM-DD
- **supersedes**: （該当する場合、旧ADR番号）
- **migrated_to_spec**: YYYY-MM-DD（spec移管済みの場合のみ）

> このADRは起票時点での決定を記録したスナップショットである。
> 現在の仕様は spec を参照すること。

## 背景

なぜこの決定が必要だったか。

## 決定

何を決めたか。起票時点での仕様案・規定を含めてよい。
仕様詳細が別specに移管済み・新設予定の場合はリンクで参照させる。

> 仕様詳細: [docs/spec/xxx.md](../spec/xxx.md) §N （該当する場合）

## 理由

なぜそう決めたか。却下した代替案も書く。

## 影響

この決定が他の仕様・実装に与える影響。

## Evidence
- commit: <ADR起票時のcommit hash>
- impl commit: <実装反映時のcommit hash。未着手なら "tbd"。doc運用ADR等で該当しない場合は "該当なし">
- 参考: <"dagsterのassets参考" / "Goのinterface慣習" 程度の軽い記述。なければ省略>
```

冒頭の「起票時点のスナップショット」注記は新規ADRに入れることを推奨する。doc-policy.md §3に既述の方針なので必須ではないが、ADR単体で読まれた場合の見落としを防ぐ。

### Evidenceの書き方

- **commit / impl commit**: git logから拾う。ADR起票と実装反映が同コミットなら1行でOK
- **参考**: OSS名・言語慣習名のレベル。URL/取得日は書かない
  - ✅ `dagsterのsoftware-defined assets参考` / `Goのinterface慣習` / `特になし`
  - ❌ `https://docs.dagster.io/... (retrieved 2026-04-19)` ← 不要

### statusの基準

- `proposed` — 議論中・まだ覆りうる
- `accepted` — 確定。変更する場合は新しいADRでsupersedesする
- `superseded` — 旧ADRの場合。新ADR番号をsupersedesに記載

---

## 4. spec運用

### 対象

言語仕様・YAMLスキーマ・MCPインターフェース定義・プロジェクトレイアウト・名前解決ルール・diagnostic定義など。
**現行仕様の唯一の正**として扱う。

### Front Matter

各specファイルの先頭に以下を置く：

```markdown
---
scope: docs/spec/ファイル名.md
status: confirmed / draft / wip
last_updated: YYYY-MM-DD
summary: >
  このdocが何を定義するかを3行以内で。
depends_on:
  - docs/adr/NNN-xxx.md   # 関連する設計判断（複数可）
---
```

### セクション末尾の由来注記

特定の決定に基づくセクションには、本文中末尾に由来注記を入れる。

```markdown
## ファイル分類ルール

brewprint YAMLは以下の3種別に分類される...

- ノード定義ファイル: `nodes:` キーをトップレベルに持つ
- View定義ファイル: `as:` キーをトップレベルに持つ
- render_index.yaml: ファイル名で識別する

> 由来: ADR-030 §決定, ADR-043 §2
```

Front Matterはspec全体の出自、セクション注記は局所的な決定の出自を示す。

### specの責務

- 現状の仕様を、可能なら**コード/ツールから直接参照できるレベルの精度**で書く
- specは「現在のスナップショット」。歴史記述や却下案はADR側に書く
- specが破綻していれば即修正する（ADRと違い遡及修正可。spec修正の経緯はcommit履歴で辿れる）

---

## 5. uc運用

### 対象

実際のユースケースをYAMLで書いたもの。スキーマ設計のinputとして使う。

### 命名・構造

UCはディレクトリ単位で管理する。各UCディレクトリがblueprintの「プロジェクト」単位に対応する。

```
docs/uc/
  NNN-タイトル/
    README.md             ← 概要・ファイル構成・ドキュメント目次・TODO/spec gap
    HANDOFF.md            ← 作業引継ぎメモ（一時的。完了後削除可）
    TASKS-UC-NNN.md       ← spec gap・追跡タスク（一時的。完了後削除可）
    render_index.yaml     ← render出力のグルーピング設定（人間が書く。省略時は1 module = 1 group）
    yaml/                 ← 実例YAML群（single source of truth）
    renders/              ← Go rendererによる自動生成物（人間が直接編集しない）
      index.md            ← masterインデックス（全groupへのリンクテーブル）
      {group-id}/
        index.md          ← groupインデックス
        dag-{task}.md
        state-{fsm}.md
        seq-{scenario}.md
        wireframe-{fsm}-{state}.html
      _cross/             ← module横断view（ER / API Table）
        er.md
        api.md
      _preview/           ← preview harness
        wireframe.html
    docs/
      coverage.md         ← カバレッジ表（人間が書く）
```

`renders/` はGo rendererが `brewprint render ./yaml/ --out ./renders/` で生成する。
goldenテスト目的でgitにcommitすることは許容するが、編集権限はGo rendererのみ。

`render_index.yaml` でgroupを定義することで、renders/の粒度をプロジェクト規模に合わせて制御できる。

```yaml
# render_index.yaml（例）
groups:
  - id: auth
    label: 認証
    modules: [auth]
  - id: commerce
    label: 商取引
    modules: [cart, order, payment]
```

詳細はADR-043（および移行後は spec/project-layout.md）を参照。

### フォーマット（README.md）

```markdown
# UC-NNN: タイトル

## 概要

何を表現したいユースケースか。

## ファイル構成

（ディレクトリツリー）

## ドキュメント

（docs/ 以下の各ファイルへのリンク表）

## TODO / spec gap

未解決事項は TASKS-UC-NNN.md を参照。
```

---

## 6. セッション開始時のClaude向け手順

1. `docs/doc-policy.md`（このdoc）を読む
2. `docs/TASKS.md` を milestone index として読む。詳細task fileはまだ読まない
3. `docs/adr/`の一覧を取得し、`accepted`なADRのタイトルを把握する
4. 作業に関連するspec / ucを必要に応じて読む
5. 作業対象 milestone が決まってから、該当する `docs/tasks/mXX-*.md` を読む

**全docを最初から読まなくていい。** ADRタイトルと `docs/TASKS.md` で文脈を把握し、必要なものだけ読む。
**closed milestone の詳細taskは原則読まなくていい。**
**現行仕様を把握したいときはspecを読む。** ADRは根拠を辿りたいときに参照する。

---

## 7. 既存ADRの漸進移行ルール

ADR-050でspec-firstに転換したため、既存49 ADR（ADR-001〜049）には現行仕様の記述が残っている。
これらは一括書き換えせず、**触れたタイミングで漸進的に移行**する。

### 移行作業の手順

1. 該当ADRの「決定」セクションから、現行仕様にあたる記述（表・スキーマ・ルール本体）を抽出
2. 対応するspecに反映（または新規specを作成）
3. ADRの「決定」セクションを「何を決めたか」の概要 + spec参照リンクに書き換える
4. ADR冒頭に `> 仕様詳細: [docs/spec/xxx.md](../spec/xxx.md)` を追加
5. spec側に `> 由来: ADR-NNN §M` を該当箇所に追記
6. ADRに `migrated_to_spec: YYYY-MM-DD` というメタ行を追加

### 移行のトリガー

- 整合性レビューでそのADRに触れたとき
- そのADRに関連する新ADRを起票するとき
- そのADRが指す仕様にバグや疑問が見つかったとき

未移行のADRは「決定」セクションが現行仕様を兼ねた状態で残る。これは過渡期の措置として許容する。

---

## 8. Claudeが自動で行うこと

- 設計決定が確定したらADRを書く（proposedで起票 → 議論完了でacceptedに更新）
- specを新規作成・更新する場合はFront Matterも更新する
- 既存ADRを覆す決定をした場合は旧ADRをsuperseededに更新し新ADRを起票する
- ADRに触れて仕様詳細を更新するときは、漸進移行ルール（§7）に従ってspecへの移行を実施する

---

## 9. ファイル操作方針

- **新規ファイル作成** → `filesystem:write_file`
- **既存ファイルの部分更新** → `str-replace:str_replace_in_file` を優先する（全文書き直しはtokenの無駄）
- **複数箇所を一度に更新** → `filesystem:edit_file`（複数editをまとめて渡せる）

str-replaceはold_strがファイル内に1箇所だけ存在する必要がある。ユニークでない文字列を指定するとエラーになるため、見出し行など十分にユニークな部分を含めて指定すること。

---

## 10. md-sectionを使ったdoc読み込みパターン

`md-section` MCPツールが使える場合、全文読みよりtoken効率がよい。

```
# 見出し一覧取得
md-section:list_headings
  path: C:\Users\imved\projects\brewprint\docs\spec\xxx.md

# 必要なセクションだけ取得
md-section:read_section
  path: C:\Users\imved\projects\brewprint\docs\spec\xxx.md
  heading: セクション名
  include_subheadings: true
```

推奨手順：
1. `filesystem:read_text_file` + `head: 30` でFront Matterだけ読む
2. `md-section:list_headings` で構造把握
3. `md-section:read_section` で必要なセクションだけ取得
4. それでも足りない場合のみ全文取得

---

## 11. 未解決事項

以下は方針未定。議論が進み次第このdocを更新すること。

- **ADR-010の複数論点混在**: 「CA強制 / ディレクトリ構造 / model-asset分離」が1 ADRに混在。ADR-011で部分的に上書きされている。分割するか追記のみにするかは **v1後に検討する**（ADR-057 §6）。漸進移行ルール（§7）に従い、ADR-010に触れたタイミングでspec移管を進める

## 12. v1で確定した運用方針

ADR-057でbrewprint v1.0.0-spec を凍結した際に、以下の運用が確定した。

### Release snapshots運用

- 仕様+実装スナップショットのgitタグは `v{MAJOR}.{MINOR}.{PATCH}-spec` 形式
  - 例: `v1.0.0-spec`
- 凍結対象: `docs/adr/*` / `docs/spec/**` / `docs/uc/**` / Go実装ツリー全体
- 運用頻度: メジャーな仕様マイルストーンに合わせて切る。毎milestoneでは切らない
- 公開contractのバージョン（MCP vNなど）は別軸で管理する

詳細は [ADR-057 §4](adr/057-brewprint-v1-snapshot.md) を参照。

### DISCLAIMER.md

- プロジェクトルートに `DISCLAIMER.md` を新設する方針が確定（ADR-057 §5）
- 文面はユーザーが起草する。Claudeは起案しない
- 法務的主張（業務時間外開発、会社リソース不使用、公知技術の組合せ）を記載する

### v1の参照UC

- UC-001（EC Checkout Flow）を v1.0.0-spec のcanonical fixtureとして固定（ADR-057 §3）
- v1範囲のspec仕様はUC-001で表現可能であることが暗黙の検証基準
- 新規UCを追加する場合もUC-001を破壊しない方向で進める

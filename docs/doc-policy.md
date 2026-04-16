# brewprint ドキュメント運用方針

> このdocはClaude（AIアシスタント）との協働における、ドキュメント管理の方針を定める。
> Claudeへの指示も兼ねているため、別会話のClaudeはこのdocを最初に読むこと。

---

## 1. プロジェクト概要

brewprintは**人間とLLMの共通設計言語**。

- 人間向け → Mermaid図（md形式でrender）
- LLM向け → signature / dep tree / inspect（MCP経由）
- YAMLはその裏側にある中間表現にすぎない

実装はGoで行う。YAMLのASTをGoで保持し、MCPツールとしてLLMに公開する。

---

## 2. ドキュメント構成

```
docs/
  doc-policy.md       ← このファイル。Claudeはセッション開始時に必ず読む
  spec/               ← 言語仕様・スキーマ定義
  adr/                ← Architecture Decision Records（設計判断の記録）
  uc/                 ← ユースケース集（実例YAML + 期待するrender結果）
```

---

## 3. ADR運用

### ADRが中心である理由

brewprintは仕様そのものより「なぜこう決めたか」の積み重ねが重要。
YAMLスキーマ・名前解決ルール・ノード型設計など、判断の根拠を残すことで：

- 別会話のClaudeが「これ変えていい？」を自分で判断できる
- 後から覆す場合も根拠ベースで議論できる

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

## 背景

なぜこの決定が必要だったか。

## 決定

何を決めたか。

## 理由

なぜそう決めたか。却下した代替案も書く。

## 影響

この決定が他の仕様・実装に与える影響。
```

### statusの基準

- `proposed` — 議論中・まだ覆りうる
- `accepted` — 確定。変更する場合は新しいADRでsupersedesする
- `superseded` — 旧ADRの場合。新ADR番号をsupersedesに記載

---

## 4. spec運用

### 対象

言語仕様・YAMLスキーマ・MCPインターフェース定義など。

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
  - docs/adr/NNN-xxx.md   # 関連する設計判断
---
```

---

## 5. uc運用

### 対象

実際のユースケースをYAMLで書いたもの。スキーマ設計のinputとして使う。

### 命名

```
docs/uc/
  NNN-タイトル.md   （例: 001-simple-dag.md）
```

### フォーマット

```markdown
# NNN: タイトル

## 概要

何を表現したいユースケースか。

## YAML

​```yaml
# 実例
​```

## 期待するrender

- DAG図: （Mermaid）
- signature: （MCPレスポンスイメージ）
```

---

## 6. セッション開始時のClaude向け手順

1. `docs/doc-policy.md`（このdoc）を読む
2. `docs/adr/`の一覧を取得し、`accepted`なADRのタイトルを把握する
3. 作業に関連するspec / ucを必要に応じて読む

**全docを最初から読まなくていい。** ADRタイトルで文脈を把握し、必要なものだけ読む。

---

## 7. Claudeが自動で行うこと

- 設計決定が確定したらADRを書く（proposedで起票 → 議論完了でacceptedに更新）
- specを新規作成・更新する場合はFront Matterも更新する
- 既存ADRを覆す決定をした場合は旧ADRをsuperseededに更新し新ADRを起票する

---

## 8. ファイル操作方針

- **新規ファイル作成** → `filesystem:write_file`
- **既存ファイルの部分更新** → `str-replace:str_replace_in_file` を優先する（全文書き直しはtokenの無駄）
- **複数箇所を一度に更新** → `filesystem:edit_file`（複数editをまとめて渡せる）

str-replaceはold_strがファイル内に1箇所だけ存在する必要がある。ユニークでない文字列を指定するとエラーになるため、見出し行など十分にユニークな部分を含めて指定すること。

---

## 9. md-sectionを使ったdoc読み込みパターン

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

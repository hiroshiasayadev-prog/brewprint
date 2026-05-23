# Investigations

`docs/investigations/` は、ADR-085 により導入された調査 artifact の置き場である。

investigation は、複雑な変更における調査結果、根拠、影響範囲、未確定点、選択肢、後続 artifact 候補を保存する。

investigation は、requirement / work item / task / ADR / spec / internal design / coverage / 別 investigation の起票・更新前に必ず必要な gate ではない。

investigation は、決定、現行仕様、要求そのもの、横断進捗、完了状態、具体的な作業手順を所有しない。

詳細 format / lifecycle は ADR-086 に従う。
MCP index / query / validation と canonical reference 方針は ADR-087 に従う。
この README は、investigation を書くときの入口と初期実務ガイドを持つ。

## Directory layout

investigation は `docs/investigations/<domain>/` 配下に置く。

例:

```text
docs/investigations/docs/
docs/investigations/adr/
docs/investigations/spec/
```

`docs/investigations/README.md` は root に置き、この directory の入口説明を所有する。

## ID and file name

investigation ID は domain ごとの連番で採番する。

```text
INV-<DOMAIN>-NNN
```

- `DOMAIN` は uppercase の短い domain label とする。
- `NNN` は investigation の `DOMAIN` ごとに 001 から始まる3桁ゼロ埋め連番とする。
- investigation ID は、他 artifact の ID / 番号体系から独立して採番する。
- ADR number、requirement ID、work item ID、task milestone、coverage mapping ID などとは結合しない。

file name は ID を大文字で残す。

```text
INV-<DOMAIN>-NNN-<slug>.md
```

例:

```text
INV-DOCS-001-investigation-artifact-format-and-lifecycle.md
```

## Metadata

investigation は Markdown 冒頭に bullet metadata を置く。
初期運用では YAML front matter を必須化しない。

### Required metadata

以下は required metadata とする。

- `status`
- `date`
- `trigger`
- `scope`
- `non_scope`
- `source_refs`
- `follow_up_candidates`

### Optional metadata

以下は optional metadata とする。
該当する情報がある場合にのみ書く。
空 field を義務付けない。

- `supersedes`
- `related_requirements`
- `related_work_items`
- `related_adrs`
- `related_specs`
- `related_internal_design`
- `related_coverage`
- `follow_up_results`

optional の `related_*` は、investigation document 内の補助参照である。
関連 artifact 側の primary trace edge を所有するものではない。
関連 artifact の責務は requirement / work item / spec / internal design 等の primary owner が引き続き所有する。External coverage artifact は MVP operational scope 外であり、導入する場合は責務を改めて判断する。

## Status

investigation の status は以下とする。

| status | 意味 |
|---|---|
| `investigating` | 調査中 |
| `concluded` | 調査結果がまとまり、後続判断に渡せる状態 |
| `superseded` | 後続 investigation または別 artifact により置き換えられた状態 |

`concluded` は、後続 artifact の採用判断や実装完了を意味しない。
調査 artifact として判断材料がまとまったことだけを表す。

`proposed` は ADR status と混同しやすいため、investigation status には使わない。

## Scope and non-scope

metadata には短い `scope` / `non_scope` を置く。
本文には `## 調査スコープ` / `## 非スコープ` を置き、詳細を書く。

scope が広がりすぎる場合は、元 investigation を無制限に拡張せず、別 investigation として切り出してよい。

## Trigger, sources, and follow-ups

investigation は、起点、調査根拠、後続候補、実際に生まれた artifact を区別する。

- `trigger`: この investigation が起票された理由または起点 artifact
- `source_refs`: 調査根拠として参照する artifact
- `follow_up_candidates`: 調査結果から起票・更新されうる artifact
- `follow_up_results`: 実際に作成・更新された artifact

`trigger` / `source_refs` / `follow_up_candidates` は required metadata とする。
`follow_up_results` は optional metadata とする。

`follow_up_candidates` が結果として空になることは許容する。
後続 artifact を生まない結論も、investigation の正当な帰結である。

`follow_up_results` は進捗管理 field ではない。
この investigation を根拠に実際に作成・更新された artifact の記録に限る。
作業状態や完了状態の管理は work item / task が所有する。

`source_refs` は調査根拠であり、MVP の canonical reference として Design Records MCP が扱う record ID-as-ref (`ADR-*` / `SPEC-*` / `INV-*`) または active `spec:` semantic ref を書く。記載値が解決できない場合は validation error とする。

`follow_up_results` は実際に作成・更新された artifact の参照である。記載する場合は、MVP の canonical reference として record ID-as-ref または active `spec:` semantic ref を書き、解決できない場合は validation error とする。

`follow_up_candidates` に artifact reference を記載する場合は、canonical reference として record ID-as-ref または active `spec:` semantic ref を書く。未作成 artifact の候補を含みうるため、canonical form で記載された参照先がまだ存在しないこと自体は validation error としない。未作成であることは、予定された後続 artifact が未解決であることを示す `info` diagnostic として可視化する。

ADR-088 により、`internal-design:` / `coverage:` / `COV-*` は MVP canonical reference / validation target として要求しない。

physical path は `source_refs` / `follow_up_results` / artifact reference として記載された `follow_up_candidates` の canonical reference として使わない。既存 document の path-based value を compatibility input として読む必要がある場合、`source_refs` / `follow_up_results` にある path は noncanonical `error`、`follow_up_candidates` にある path は noncanonical candidate を示す `info` diagnostic として扱う。

`trigger` / optional の `related_*` の resolve / validation rule は後続 contract で定義する。

## Body structure

標準的な investigation は、以下の構成を正式テンプレートとする。
調査内容に応じて、不要な section は省略してよい。

```markdown
# INV-<DOMAIN>-NNN: <title>

- **status**: investigating
- **date**: YYYY-MM-DD
- **trigger**: <起点 artifact または起票理由>
- **scope**: <短い調査スコープ>
- **non_scope**: <短い非スコープ>
- **source_refs**:
  - <artifact ID-as-ref or semantic ref>
- **follow_up_candidates**:
  - <candidate artifact ID-as-ref / semantic ref / human-readable candidate, or なし>
- **follow_up_results**:  # 実際に作成・更新された artifact がある場合のみ
  - <artifact ID-as-ref or semantic ref>

## 調査スコープ

## 非スコープ

## 背景

## 調査したもの

## 調査項目ごとの確認結果

### Q1: <調査項目>

#### 確認対象

#### 観測事実

#### 候補

#### 判断に必要な観点

#### 後続判断先

## 横断的な観測事実

## 後続判断に渡す候補

## 推奨案

## 後続 artifact 候補

## 未確定点
```

investigation は決定を所有しない。
ただし、調査者の見立てとして `推奨案` を書いてよい。
その場合は「〜と考えられる」「〜が妥当と見られる」のように、判断ではなく調査結果に基づく候補であることが分かる表現にする。

`推奨案` は、後続 ADR / README / doc-policy / task file 等の判断を代替しない。
判断が必要な内容は、`後続判断に渡す候補` または `推奨案` として整理し、採用判断は後続 artifact に委ねる。

## Starting another investigation

investigation の調査中に別領域の調査が必要になった場合、別 investigation を起票してよい。

ただし、別 investigation の起票は必須ではない。
軽微な追加確認は元 investigation 内に留めてよい。

別 investigation を起票する場合、元 investigation の `follow_up_candidates` または `follow_up_results` に記録する。

## Current investigations

- `docs/investigations/docs/INV-DOCS-001-investigation-artifact-format-and-lifecycle.md` — investigation artifact の directory / ID / format / lifecycle / authoring boundary を調査した。
- `docs/investigations/docs/INV-DOCS-002-external-coverage-artifact-necessity.md` — semantic trace MVP で external coverage artifact が必要かを調査した。
- `docs/investigations/docs/INV-DOCS-003-internal-design-semantic-trace-mvp-necessity.md` — `internal-design:` endpoint / realization relation の MVP 必要性を調査した。

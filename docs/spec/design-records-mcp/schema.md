---
scope: docs/spec/design-records-mcp/schema.md
status: draft
last_updated: 2026-05-12
summary: >
  Design Records MCP MVP が読む design_record metadata schema、record model、
  H1 title 抽出、diagnostic category を定義する。
depends_on:
  - docs/adr/076-design-records-mcp.md
  - docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md
design_record:
  id: SPEC-design-records-mcp-schema
  kind: spec
  status: draft
  depends_on:
    - ADR-076
    - ADR-077
---

# Design Records MCP schema

## Record source

Design Records MCP MVP は Markdown file から design record を構築する。

MVP で読む source は以下である。

| source | 用途 |
|---|---|
| ADR 箇条書きmetadata | decision record の `status` / `date` / `depends_on` / `supersedes` / `migrated_to_spec` |
| spec YAML front matter | spec record の `scope` / `status` / `design_record` metadata。top-level `depends_on` は doc-policy 用出自 path list として読むが、record dependency には使わない |
| Markdown H1 | `title` 抽出 |
| file path | record path / filename ID validation |
| Markdown headings | `get_record` response の headings |
| Markdown body | `get_record(include_body=true)` の raw body |

MVP では Markdown 本文の自然言語から依存関係や migration 状態を推定しない。

> 由来: ADR-076 §front matter 方針, ADR-077 §get_record の責務

## Metadata source

MVP では、ADR と spec で metadata source が異なる。

ADR は既存フォーマットを維持し、H1直下の箇条書きmetadataを読む。

```markdown
# 076: Design Records MCP

- **status**: accepted
- **date**: 2026-05-11
- **depends_on**: ADR-050, ADR-068
- **supersedes**:
- **migrated_to_spec**:
```

spec は既存の YAML front matter を読む。
spec 側では `design_record` metadata を追加できる。

MVP では、ADR-076 の ADR 向け `design_record` front matter 例は採用しない。
ADR record は既存 ADR フォーマットの H1 直下箇条書きmetadataから構築し、ADR に YAML front matter / `design_record` block を追加しない。

spec 側の例:

```yaml
design_record:
  id: SPEC-design-records-mcp-schema
  kind: spec
  status: draft
  depends_on:
    - ADR-076
    - ADR-077
```

### ADR bullet metadata 文法

ADR の箇条書きmetadataは、H1直後から始まる metadata block から読む。
MVP では、metadata block は H1 の次行から、最初の H2 行または blockquote 行の直前までとする。
H2 行は `## ` で始まる行、blockquote 行は `>` で始まる行を指す。
metadata block 内の空行は許容する。

認識する metadata 行は以下の形式に限定する。

```text
- **<key>**: <value>
```

制約:

- bold marker `**` は必須とする
- key は case-sensitive とする
- 認識する key は `status` / `date` / `depends_on` / `supersedes` / `migrated_to_spec` のみとする
- 認識しない key は MVP では無視する
- `value` は前後 whitespace を trim する
- `value` が空または whitespace-only の場合、その metadata は未指定として扱う

list value は comma 区切りとする。
`depends_on` / `supersedes` は、non-empty の場合 comma で分割し、各値を trim する。
空 value の `depends_on` / `supersedes` は empty list に正規化する。
空 value の `migrated_to_spec` は `null` に正規化する。

`date` は ADR metadata として parse してよいが、MVP の record field には含めない。

> 由来: ADR-076 §front matter 方針

## Field definitions

Internal record model では、ADR / spec の metadata source 差分を吸収し、以下の field に正規化する。

| field | required | type | meaning |
|---|---:|---|---|
| `id` | yes | string | record ID。ADR は H1 から導出し、spec は `design_record.id` から読む |
| `kind` | yes | string | record kind。ADR は `decision` として導出し、spec は `design_record.kind` から読む |
| `status` | yes | string | kind 別 status。ADR は箇条書きmetadata、spec は top-level front matter を canonical とする |
| `depends_on` | no | list<string> | 依存する record ID。ADR は箇条書きmetadata、spec は `design_record.depends_on` から読む |
| `supersedes` | no | list<string> | この record が置き換える record ID。ADR metadata では空欄可 |
| `migrated_to_spec` | no | string or null | ADR の仕様記述が spec へ移管済みの場合の日付。ADR metadata では空欄可 |

MVP では `title` は metadata に持たせない。
`title` は Markdown H1 から抽出する。

ADR 箇条書きmetadataの `date` は parse してよいが、MVP では record field として持たせない。

MVP では以下の metadata は持たせない。

- `topics`
- `affects`
- `refines`
- `conflicts_with`
- `owns`
- `migration.state`
- section 単位の由来情報

> 由来: ADR-076 §front matter 方針, ADR-077 §list_records の責務

## Discovery and index inclusion

MVP の discovery は以下の規則に従う。

| kind | discovery rule |
|---|---|
| `decision` | `docs/adr/*.md` の Markdown file を ADR 候補として読む |
| `spec` | `docs/spec/**/*.md` のうち YAML front matter に `design_record.id` と `design_record.kind` を持つ file のみを spec record として読む |

`design_record` を持たない spec は MVP index 対象外とする。
その場合も `missing_design_record` diagnostic は出さない。

`design_record.kind` が `decision` / `spec` 以外の場合、MVP では index 対象外とし、diagnostic は出さない。

> 由来: ADR-076 §bootstrap方針, ADR-077 §validate_records の責務

## `id`

`id` は record の安定識別子である。

MVP で扱う ID 形式は以下とする。

| kind | ID 例 | 備考 |
|---|---|---|
| `decision` | `ADR-076` | ADR 番号を3桁ゼロ埋めで持つ |
| `spec` | `SPEC-design-records-mcp-schema` | spec 用の stable ID |

`decision` record の canonical ID は、H1 の番号から `ADR-NNN` として導出する。
filename 先頭の番号は canonical ID との一致検査にのみ使う。
H1 が不正な場合は `invalid_h1_title` を出し、filename 由来の ID を canonical ID として採用しない。

canonical ID の番号と filename 先頭の番号が一致していなければならない。

例:

```text
H1: # 076: Design Records MCP
id: ADR-076
path: docs/adr/076-design-records-mcp.md
```

一致しない場合、`filename_id_mismatch` とする。

> 由来: ADR-077 §validate_records の責務

## `kind`

MVP で許可する `kind` は以下である。

| kind | meaning |
|---|---|
| `decision` | ADR |
| `spec` | spec |

MVP では task / UC / impl note を `kind` として扱わない。

> 由来: ADR-076 §MVP対象

## `status`

`status` の値域は `kind` ごとに異なる。

| kind | allowed status |
|---|---|
| `decision` | `proposed` / `accepted` / `superseded` |
| `spec` | `confirmed` / `draft` / `wip` |

`decision` record の `status` は ADR 箇条書きmetadataから読む。

`spec` record の canonical `status` は top-level YAML front matter の `status` とする。
`design_record.status` が存在する場合、top-level `status` と同値でなければならない。
不一致の場合は `spec_status_mismatch` とする。

`kind` に対して許可されない `status` は `invalid_status_for_kind` とする。

> 由来: ADR-076 §front matter 方針, ADR-077 §validate_records の責務

## `depends_on`

`depends_on` は、この record が判断・仕様上依存する record ID の list である。

`decision` record の `depends_on` は ADR 箇条書きmetadataから読む。

`spec` record の record dependency は `design_record.depends_on` から読む。
spec top-level front matter の `depends_on` は doc-policy 用の出自 path list であり、record dependency としては扱わない。

MVP では参照先 ID の存在確認のみを行う。
参照元・参照先の status 組み合わせは検査しない。

存在しない ID を参照している場合、`missing_depends_on_target` とする。

> 由来: ADR-077 §validate_records の責務

## `supersedes`

`supersedes` は、この record が置き換える record ID の list である。

`decision` record の `supersedes` は ADR 箇条書きmetadataから読む。
`spec` record の `supersedes` は MVP では常に empty list に正規化する。
spec の `design_record.supersedes` に値が存在しても MVP では無視し、diagnostic は出さない。

MVP では参照先 ID の存在確認のみを行う。
置き換え対象の status や逆参照の整合性は検査しない。

存在しない ID を参照している場合、`missing_supersedes_target` とする。

> 由来: ADR-077 §validate_records の責務

## `migrated_to_spec`

`migrated_to_spec` は、ADR の仕様記述が spec へ移管済みであることを表す metadata である。

`decision` record の `migrated_to_spec` は ADR 箇条書きmetadataから読む。
`spec` record の `migrated_to_spec` は MVP では常に `null` に正規化する。
spec の `design_record.migrated_to_spec` に値が存在しても MVP では無視し、diagnostic は出さない。

ADR 箇条書きmetadataの `migrated_to_spec` が空欄または whitespace-only の場合、`null` に正規化する。

non-empty の場合、MVP では `YYYY-MM-DD` 形式のみ有効とする。
path、record ID、任意文字列による移管先表現は MVP 外とする。

MVP では `migration.state` のような正規化状態語彙は導入しない。
`migrated_to_spec` の形式が不正な場合、`invalid_migrated_to_spec` とする。

> 由来: ADR-076 §front matter 方針, ADR-077 §validate_records の責務

## Title extraction

MVP では `title` を Markdown H1 から抽出する。

ADR の H1 は以下の形式のみ valid とする。

```text
^#\s+(?P<num>\d{3}):\s+(?P<title>.+?)\s*$
```

制約:

- `num` は3桁ゼロ埋め必須とする
- `ADR-` prefix は許可しない
- 区切りは ASCII colon `:` とする
- colon 直後には1文字以上の whitespace が必要である
- `title` は trim 後 non-empty でなければならない
- canonical ID は `ADR-<num>` として導出する
- filename 先頭番号との比較は3桁ゼロ埋め文字列一致とする

例:

```markdown
# 076: Design Records MCP
```

H1 が存在しない、または期待形式に合わない場合、`invalid_h1_title` とする。

spec record の H1 は `# <title>` 形式とし、先頭に番号を要求しない。
spec record の title は、H1 行から leading `#` とその直後の whitespace を除き、前後 whitespace を trim した残りとする。

filename からの title 推定は MVP では行わない。

> 由来: ADR-077 §list_records の責務, ADR-077 §理由

## Record model

Design Records MCP の internal record model は、少なくとも以下の情報を持つ。

| field | source | meaning |
|---|---|---|
| `id` | ADR: H1 / spec: `design_record.id` | record ID |
| `kind` | ADR: fixed `decision` / spec: `design_record.kind` | record kind |
| `title` | H1 | human-readable title |
| `status` | ADR: bullet metadata / spec: top-level front matter | record status |
| `path` | filesystem | Markdown file path |
| `depends_on` | ADR: bullet metadata / spec: `design_record.depends_on` | dependency IDs |
| `supersedes` | ADR: bullet metadata / spec: empty list in MVP | superseded IDs |
| `migrated_to_spec` | ADR: bullet metadata / spec: null in MVP | ADR migration marker |
| `headings` | Markdown parse | heading list |
| `body` | Markdown file | raw body, requested only when needed |

`headings` は ATX heading のみを対象とする。
YAML front matter 内、および fenced code block 内の `#` で始まる行は heading として扱わない。
setext heading は MVP では扱わない。

`body` は `get_record(include_body=true)` の場合だけ response に含める。
本文は整形・要約・正規化せず、元ファイル内容をそのまま返す。

> 由来: ADR-077 §list_records の責務, ADR-077 §get_record の責務

## Diagnostic category

MVP の `validate_records` は以下の diagnostic category を返す。

Diagnostic は検査軸ごとに独立して発火し、1 record に複数 diagnostic が付いてよい。

| category | severity | meaning |
|---|---|---|
| `duplicate_id` | error | 複数 record が同じ正規化後 record ID を持つ |
| `filename_id_mismatch` | error | `decision` record の ID 番号と filename 番号が一致しない |
| `invalid_h1_title` | error | H1 が存在しない、または期待形式に合わない |
| `invalid_status_for_kind` | error | `kind` に対して許可されない `status` を持つ |
| `spec_status_mismatch` | error | spec top-level `status` と `design_record.status` が一致しない |
| `missing_depends_on_target` | error | `depends_on` の参照先 ID が存在しない |
| `missing_supersedes_target` | error | `supersedes` の参照先 ID が存在しない |
| `invalid_migrated_to_spec` | error | `migrated_to_spec` の値が不正 |
| `missing_record_path` | error | discovery で候補 path を検出したが、read/stat に失敗した |

MVP では以下を diagnostic category に含めない。

- `accepted_but_not_migrated`
- `missing_design_record`
- status 組み合わせの妥当性
- spec section 単位の由来不足
- 自然言語本文と metadata の意味的不一致

`missing_record_path` は、filesystem scan または path normalization により record 候補 path を検出したが、実際の read/stat に失敗した場合に出す。
例として、scan 後に file が削除された場合、permission denied、symlink target missing、path normalization 後の path が存在しない場合を含む。

> 由来: ADR-077 §validate_records の責務

## Bootstrap metadata

MVP 検証用に、少数の代表 record を bootstrap 対象にする。

ADR では既存の箇条書きmetadataを利用し、追加の YAML front matter は導入しない。
spec では YAML front matter 内の `design_record` metadata を利用する。

初期候補は以下である。

- ADR-050
- ADR-067〜ADR-077
- `docs/spec/design-records-mcp/**`

これ以外の既存 ADR/spec には一括付与しない。
整合性レビュー、新規 ADR 起票、関連 spec 更新で触れたタイミングで漸進的に追加する。

> 由来: ADR-076 §bootstrap方針

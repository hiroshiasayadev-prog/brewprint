---
scope: docs/spec/design-records-mcp/schema.md
status: draft
last_updated: 2026-05-26
summary: >
  Design Records MCP MVP が読む design_record metadata schema、record model、
  H1 title 抽出、diagnostic category を定義する。
depends_on:
  - docs/adr/076-design-records-mcp.md
  - docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md
  - docs/adr/086-investigation-artifact-format-and-lifecycle.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
  - docs/adr/090-design-records-mcp-batch-retrieval-tool-boundary.md
design_record:
  id: SPEC-design-records-mcp-schema
  kind: spec
  status: draft
  depends_on:
    - ADR-076
    - ADR-077
    - ADR-086
    - ADR-087
    - ADR-088
    - ADR-090
---

# Design Records MCP schema

## Record source

Design Records MCP MVP は Markdown file から design record を構築する。

MVP で読む source は以下である。

| source | 用途 |
|---|---|
| ADR 箇条書きmetadata | decision record の `status` / `date` / `depends_on` / `supersedes` / `migrated_to_spec` |
| investigation 箇条書きmetadata | investigation record の `status` / `date` / `trigger` / `scope` / `non_scope` / `source_refs` / `follow_up_candidates` / optional related metadata / `follow_up_results` |
| spec YAML front matter | spec record の `scope` / `status` / `design_record` metadata。top-level `depends_on` は doc-policy 用出自 path list として読むが、record dependency には使わない |
| Markdown H1 | `title` 抽出 |
| file path | record path / filename ID validation |
| Markdown headings | `get_record` / `get_records` の found record response の headings |
| Markdown body | `get_record(include_body=true)` / `get_records(include_body=true)` の found record raw body |

MVP では Markdown 本文の自然言語から依存関係や migration 状態を推定しない。

> 由来: ADR-076 §front matter 方針, ADR-077 §get_record の責務

## Metadata source

Design Records MCP では、decision / spec / investigation で metadata source が異なる。

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
    - ADR-087
```

investigation は ADR-086 の形式に従い、H1直下の箇条書きmetadataを読む。

```markdown
# INV-MCP-001: Design Records MCP investigation support

- **status**: concluded
- **date**: 2026-05-23
- **trigger**: ADR-087
- **scope**: investigation MCP integration
- **non_scope**: writer tools
- **source_refs**:
  - ADR-087
- **follow_up_candidates**:
  - なし
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

### Investigation bullet metadata 文法

Investigation の metadata block は ADR と同様に H1 直後から最初の H2 行または blockquote 行の直前までとする。
`status` / `date` / `trigger` / `scope` / `non_scope` / `source_refs` / `follow_up_candidates` は required metadata とする。
`supersedes` / `related_requirements` / `related_work_items` / `related_adrs` / `related_specs` / `related_internal_design` / `related_coverage` / `follow_up_results` は、記載がある場合のみ読む optional metadata とする。

`source_refs` と `follow_up_results` の各値は canonical reference として Design Records MCP が扱う record ID-as-ref (`ADR-*` / `SPEC-*` / `INV-*`) または active `spec:` semantic ref を用いる。記載値は解決可能でなければならず、unresolved は error とする。physical path が入力に現れた場合、compatibility input として読み取ってよいが canonical form ではなく、error diagnostic の対象とする。
ADR-088 により、`internal-design:` / `coverage:` / `COV-*` は MVP canonical reference / resolver input として要求しない。
`follow_up_candidates` に artifact reference が記載される場合も、record ID-as-ref または active `spec:` semantic ref の canonical form を用いる。候補は未作成 artifact を指しうるため、canonical form の unresolved は error にせず、予定された後続 artifact が未作成であることを示す `info` diagnostic として返す。Physical path による candidate は canonical form ではなく、noncanonical candidate を示す `info` diagnostic として返す。
`trigger` / `related_*` の resolve / validation rule はこの版では確定しない。

> 由来: ADR-086 §4〜§7, ADR-087 §5〜§8

## Field definitions

Internal record model は、metadata source の差分を吸収し、共通 field と kind 固有 detail object に正規化する。Public response もこの shape のみを返し、既存 flat metadata field と kind 固有 detail object を同時に返す compatibility contract は設けない。Spec 更新後、parser / index / list / get / validate と対応 tests は同一切替単位でこの response shape に移行する。

Common fields:

| field | required | type | meaning |
|---|---:|---|---|
| `id` | yes | string | record ID。decision / investigation は H1 から導出し、spec は `design_record.id` から読む |
| `kind` | yes | string | record kind |
| `title` | yes | string | Markdown H1 由来の human-readable title |
| `status` | yes | string | kind 別 status |
| `path` | yes | string | repository root からの Markdown file path |

Kind-specific details:

| detail object | fields |
|---|---|
| `decision` | `depends_on`, `supersedes`, `migrated_to_spec` |
| `spec` | `depends_on` |
| `investigation` | `trigger`, `scope`, `non_scope`, `source_refs`, `follow_up_candidates`, optional `supersedes`, `related_*`, `follow_up_results` |

`headings` と requested raw `body` は `get_record` の取得内容として common response に追加でき、`get_records` の `retrieval_status: "found"` item でも同一 record representation を再利用する。

ADR 箇条書きmetadataの `date` は parse してよいが、record response field としては持たせない。
investigation の `date` も同様に metadata として parse するが、common response field には含めない。

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
| `investigation` | `docs/investigations/*/INV-*-*.md` の Markdown file を investigation 候補として読む |

`design_record` を持たない spec は index 対象外とする。
その場合も `missing_design_record` diagnostic は出さない。

spec の `design_record.kind` が `spec` 以外の場合、この版では index 対象外とし、diagnostic は出さない。

本specは後続の `requirement` その他の record kind 追加を制限しない。

> 由来: ADR-076 §bootstrap方針, ADR-077 §validate_records の責務

## `id`

`id` は record の安定識別子である。

MVP で扱う ID 形式は以下とする。

| kind | ID 例 | 備考 |
|---|---|---|
| `decision` | `ADR-076` | ADR 番号を3桁ゼロ埋めで持つ |
| `spec` | `SPEC-design-records-mcp-schema` | spec 用の stable ID |
| `investigation` | `INV-MCP-001` | ADR-086 に従う domain-scoped ID |

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

現在 index / query / validation 対象とする `kind` は以下である。

| kind | meaning |
|---|---|
| `decision` | ADR |
| `spec` | spec |
| `investigation` | investigation artifact |

この表は record kind の閉じた列挙ではない。ADR-081 により予約された `requirement` その他の kind は、後続判断で追加しうる。
MVP では task / UC / impl note を record kind として扱わない。

> 由来: ADR-076 §MVP対象, ADR-087 §1

## `status`

`status` の値域は `kind` ごとに異なる。

| kind | allowed status |
|---|---|
| `decision` | `proposed` / `accepted` / `superseded` |
| `spec` | `confirmed` / `draft` / `wip` |
| `investigation` | `investigating` / `concluded` / `superseded` |

`decision` record の `status` は ADR 箇条書きmetadataから読む。

`spec` record の canonical `status` は top-level YAML front matter の `status` とする。
`design_record.status` が存在する場合、top-level `status` と同値でなければならない。
不一致の場合は `spec_status_mismatch` とする。

`investigation` record の `status` は investigation 箇条書きmetadataから読む。

`kind` に対して許可されない `status` は `invalid_status_for_kind` とする。

> 由来: ADR-076 §front matter 方針, ADR-077 §validate_records の責務, ADR-086 §5

## `depends_on`

`depends_on` は、この record が判断・仕様上依存する record ID の list である。

`decision` record の `depends_on` は ADR 箇条書きmetadataから読む。

`spec` record の record dependency は `design_record.depends_on` から読む。
spec top-level front matter の `depends_on` は doc-policy 用の出自 path list であり、record dependency としては扱わない。

MVP では参照先 ID の存在確認のみを行う。
参照元・参照先の status 組み合わせは検査しない。

`depends_on` が参照できる canonical record ID-as-ref は `ADR-*` / `SPEC-*` / `INV-*` とする。したがって ADR / spec が investigation record (`INV-*`) に依存することは valid である。

存在しない ID を参照している場合、`missing_depends_on_target` とする。現行 implementation は investigation record integration 前であるため、`ADR-086` が参照する `INV-DOCS-001` に対して同 diagnostic を返すが、これは invalid dependency ではなく M19 で解消すべき implementation gap である。

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

investigation record の H1 は以下の形式のみ valid とする。

```text
^#\s+(?P<id>INV-[A-Z0-9-]+-\d{3}):\s+(?P<title>.+?)\s*$
```

investigation の canonical ID は H1 の `id` から取得する。filename は `INV-<DOMAIN>-NNN-<slug>.md` 形式であり、H1 の canonical ID と prefix が一致しなければならない。不一致は filename / ID mismatch として診断対象にする。

filename からの title 推定は MVP では行わない。

> 由来: ADR-077 §list_records の責務, ADR-077 §理由

## Record model

Design Records MCP の internal record model は、少なくとも以下の情報を持つ。

| field | source | meaning |
|---|---|---|
| `id` | decision/investigation: H1 / spec: `design_record.id` | record ID |
| `kind` | source kind / spec: `design_record.kind` | record kind |
| `title` | H1 | human-readable title |
| `status` | decision/investigation: bullet metadata / spec: top-level front matter | record status |
| `path` | filesystem | Markdown file path |
| `decision` | ADR bullet metadata | decision-specific detail object |
| `spec` | spec `design_record` metadata | spec-specific detail object |
| `investigation` | investigation bullet metadata | investigation-specific detail object |
| `headings` | Markdown parse | heading list for `get_record` |
| `body` | Markdown file | raw body, requested only when needed by `get_record` または `get_records` |

`get_records` の response item wrapper は record model 自体ではない。`retrieval_status: "found"` item の `record` は上記 record model を返し、`retrieval_status: "not_found"` item の `record` は `null` とする。取得状態を表す `retrieval_status` は record lifecycle の `status` に混在させない。

`headings` は ATX heading のみを対象とする。
YAML front matter 内、および fenced code block 内の `#` で始まる行は heading として扱わない。
setext heading は MVP では扱わない。

`body` は `get_record(include_body=true)` または `get_records(include_body=true)` の found record response にだけ含める。
本文は整形・要約・正規化・truncate を行わず、元ファイル内容をそのまま返す。

> 由来: ADR-077 §list_records の責務, ADR-077 §get_record の責務

## Diagnostic category

`resolve_reference` は direct query の結果として、supported form だが target が存在しない場合に `unresolved_reference`、同一 ref が複数 target に一致して単一解決できない場合に `ambiguous_reference`、MVP resolver 対象外として behavior を定義する入力には `unsupported_reference` を返す。Reserved prefix `yaml:` の public resolver input / direct query response behavior、および investigation metadata validation behavior は MVP で定義しない。これら resolution response の diagnostic と、`validate_records` が参照元 field や index defect に対して返す下記 validation diagnostic は区別する。

`get_records` は retrieval / request-level diagnostic として以下を返す。

| category | severity | placement | required additional fields | meaning |
|---|---|---|---|---|
| `record_not_found` | error | item-level | `requested_id` | requested exact record ID lookup key が index に存在しない |
| `duplicate_requested_id_ignored` | info | top-level | `requested_id`, `first_index`, `duplicate_indexes` | 同一 requested ID の2回目以降を無視し、first occurrence の item のみ返した |

`first_index` と `duplicate_indexes` は request `ids` array の zero-based index とする。`duplicate_requested_id_ignored` は重複した requested ID ごとに一件返す。これらは request / retrieval に対する diagnostic であり、record metadata defect を示す `record_id` field は用いない。

MVP の `validate_records` は以下の diagnostic category を返す。

Diagnostic は検査軸ごとに独立して発火し、1 record に複数 diagnostic が付いてよい。

| category | severity | meaning |
|---|---|---|
| `duplicate_id` | error | 複数 record が同じ正規化後 record ID を持つ |
| `filename_id_mismatch` | error | `decision` または `investigation` record の canonical ID と filename ID 部分が一致しない |
| `invalid_h1_title` | error | H1 が存在しない、または期待形式に合わない |
| `invalid_status_for_kind` | error | `kind` に対して許可されない `status` を持つ |
| `spec_status_mismatch` | error | spec top-level `status` と `design_record.status` が一致しない |
| `missing_depends_on_target` | error | `depends_on` の参照先 ID が存在しない |
| `missing_supersedes_target` | error | `supersedes` の参照先 ID が存在しない |
| `invalid_migrated_to_spec` | error | `migrated_to_spec` の値が不正 |
| `missing_record_path` | error | discovery で候補 path を検出したが、read/stat に失敗した |
| `invalid_semantic_ref_declaration` | error | spec front matter の `semantic_refs` entry または `sections` key が active `spec:` grammar に従わない |
| `missing_section_target` | error | spec front matter の `sections` value と一致する Markdown heading が存在しない |
| `ambiguous_section_target` | error | spec front matter の `sections` value が同一 document 内の複数 heading に一致し、section target を単一解決できない |
| `duplicate_semantic_ref` | error | active `spec:` document-level または section-level ref が複数 target へ宣言され、単一解決できない |
| `unresolved_source_ref` | error | investigation `source_refs` に記載された supported canonical ref が解決不能 |
| `unresolved_follow_up_result` | error | investigation `follow_up_results` に記載された supported canonical ref が解決不能 |
| `unresolved_follow_up_candidate` | info | investigation `follow_up_candidates` に記載された supported canonical ref が未解決 |
| `noncanonical_source_ref` | error | investigation `source_refs` に physical path が記載された |
| `noncanonical_follow_up_result` | error | investigation `follow_up_results` に physical path が記載された |
| `noncanonical_follow_up_candidate` | info | investigation `follow_up_candidates` に physical path が記載された |
| `unsupported_reference` | error / info | MVP が unsupported と定義する metadata reference。`source_refs` / `follow_up_results` では error、`follow_up_candidates` では info。Reserved `yaml:` はこの category の対象に含めず、MVP では behavior を定義しない |

Investigation reference diagnostic (`unresolved_*` / `noncanonical_*` / metadata field 由来の `unsupported_reference`) は、通常の `category` / `severity` / `record_id` / `path` / `message` に加えて、`field`、`value`、`ref_status` を必須で返す。`field` は `source_refs` / `follow_up_results` / `follow_up_candidates` のいずれか、`value` は入力 reference 文字列、`ref_status` は `unresolved` / `unsupported` / `noncanonical` のいずれかとする。Investigation metadata が duplicate semantic ref または duplicate record ID を指して単一解決できない場合は field-specific diagnostic を追加せず、index defect を示す `duplicate_semantic_ref` または `duplicate_id` のみを返す。これら duplicate diagnostic および spec declaration / section lookup diagnostic は investigation metadata field 由来の追加 field を要求しない。

`follow_up_candidates` の参照先が未作成であること自体は error としない。`validate_records.ok` は error diagnostic の有無だけで決まり、info diagnostic が存在しても failure にはならない。Coverage mapping、semantic realization relation、`internal-design:` / `coverage:` / `COV-*` の解決・診断は MVP diagnostic scope に含めない。

MVP では以下を diagnostic category に含めない。

- `accepted_but_not_migrated`
- `missing_design_record`
- status 組み合わせの妥当性
- spec section 単位の由来不足
- 自然言語本文と metadata の意味的不一致

`missing_record_path` は、filesystem scan または path normalization により record 候補 path を検出したが、実際の read/stat に失敗した場合に出す。
例として、scan 後に file が削除された場合、permission denied、symlink target missing、path normalization 後の path が存在しない場合を含む。

> 由来: ADR-077 §validate_records の責務, ADR-090 §Partial result / Ordering と duplicate requested ID

## Bootstrap metadata

MVP 検証用に、少数の代表 record を bootstrap 対象にする。

ADR では既存の箇条書きmetadataを利用し、追加の YAML front matter は導入しない。
spec では YAML front matter 内の `design_record` metadata を利用する。

初期候補は以下である。

- ADR-050
- ADR-067〜ADR-077
- ADR-086〜ADR-088
- `docs/investigations/docs/INV-DOCS-001-investigation-artifact-format-and-lifecycle.md`
- `docs/spec/design-records-mcp/**`

これ以外の既存 ADR/spec には一括付与しない。
整合性レビュー、新規 ADR 起票、関連 spec 更新で触れたタイミングで漸進的に追加する。

> 由来: ADR-076 §bootstrap方針

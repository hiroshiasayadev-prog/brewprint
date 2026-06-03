---
scope: docs/spec/design-records-mcp/schema.md
status: draft
last_updated: 2026-06-03
summary: >
  Design Records MCP MVP が読む design record / workflow artifact metadata schema、
  record model、authoring guidance source model、authoring transaction schema concept、
  H1 title 抽出、diagnostic category を定義する。
depends_on:
  - docs/adr/076-design-records-mcp.md
  - docs/adr/077-design-records-mcp-mvp-boundary-and-tool-prioritization.md
  - docs/adr/086-investigation-artifact-format-and-lifecycle.md
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
  - docs/adr/090-design-records-mcp-batch-retrieval-tool-boundary.md
  - docs/adr/092-design-records-mcp-workflow-artifact-record-and-relation-boundary.md
  - docs/adr/093-design-records-mcp-authoring-transaction-model.md
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
    - ADR-092
    - ADR-093
---

# Design Records MCP schema

## Record source

Design Records MCP MVP は Markdown file から design record を構築する。

MVP で読む source は以下である。

| source | 用途 |
|---|---|
| ADR 箇条書きmetadata | decision record の `status` / `date` / `depends_on` / `supersedes` / `migrated_to_spec` |
| investigation 箇条書きmetadata | investigation record の `status` / `date` / `trigger` / `scope` / `non_scope` / `source_refs` / `follow_up_candidates` / optional related metadata / `follow_up_results` |
| requirement 箇条書きmetadata | requirement record の `id` / `status` / `date` / `source_refs` / `work_items` |
| work item 箇条書きmetadata | work item record の `id` / `status` / `date` / `source_requirement` / `impact_refs` / `tasks` |
| task 箇条書きmetadata | task record の `id` / `status` / `date` / `work_item` / `source_requirement` / `estimate` / `depends_on` / `outputs` |
| spec YAML front matter | spec record の `scope` / `status` / `design_record` metadata。top-level `depends_on` は doc-policy 用出自 path list として読むが、record dependency には使わない |
| Markdown H1 | `title` 抽出 |
| file path | record path / filename ID validation |
| Markdown headings | `get_record` / `get_records` の found record response の headings |
| Markdown body | `get_record(include_body=true)` / `get_records(include_body=true)` の found record raw body |
| authoring guide Markdown | `docs/guides/*.md` から `list_authoring_guides` / `get_authoring_guidance` 用の guide ID / title / abstract / content を取得する |

MVP では Markdown 本文の自然言語から依存関係や migration 状態を推定しない。

> 由来: ADR-076 §front matter 方針, ADR-077 §get_record の責務

## Metadata source

Design Records MCP では、decision / spec / investigation / requirement / work item / task で metadata source が異なる。

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

ADR-092 により、investigation の上記 validated reference field に追加で使用できる workflow artifact ID-as-ref は `REQ-<DOMAIN>-NNN` / `WORK-<DOMAIN>-NNN` とする。`TASK-*` は workflow artifact 間 relation と direct resolver input では supported だが、investigation metadata field の canonical reference form には含めない。`TASK-*` が investigation の `source_refs` / `follow_up_results` に現れた場合は `unsupported_reference` error、`follow_up_candidates` に現れた場合は `unsupported_reference` info とする。

### Workflow artifact bullet metadata 文法

Requirement / work item / task の metadata block は ADR / investigation と同様に H1 直後から最初の H2 行または blockquote 行の直前までとする。認識する metadata 行は `- **<key>**: <value>` と、その直下のインデント付き list item とする。

Workflow artifact ID grammar は以下とする。

```text
REQ-<DOMAIN>-NNN
WORK-<DOMAIN>-NNN
TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>
```

- `<DOMAIN>` は uppercase ASCII letter / digit / hyphen で構成し、先頭と末尾を hyphen にしない。
- Requirement / work item の `NNN` と task の `<WORK-SEQUENCE>` は3桁ゼロ埋め decimal sequence とする。
- Task の `<TASK-SEQUENCE>` は2桁ゼロ埋め decimal sequence とする。
- ID は metadata `id`、H1 ID、filename prefix で一致しなければならない。
- Workflow relation は metadata field に記載された ID-as-ref だけから読み、task ID 文字列または path から親 relation を推定しない。

Requirement の認識 field:

- required: `id`, `status`, `date`, `source_refs`, `work_items`

Work item の認識 field:

- required: `id`, `status`, `date`, `source_requirement`, `impact_refs`, `tasks`

Task の認識 field:

- required: `id`, `status`, `date`, `work_item`, `source_requirement`, `estimate`, `depends_on`, `outputs`

Workflow artifact の required metadata は presence validation の対象とする。

- Required scalar field は存在し、かつ non-empty でなければならない。
- Required list field は存在しなければならない。
- Required list field の empty list は、artifact-specific rule が non-empty を要求しない限り valid とする。
- Required list field 内の empty item は validation error とし、metadata diagnostic category は `empty_required_metadata` とする。
- `date` は required scalar metadata として扱い、strict `YYYY-MM-DD` format でなければならない。

`task.depends_on` は field が存在し、値と直下の list item がともに空の場合、empty list `[]` として正規化する。この場合、workflow relation diagnostic は生成しない。

`source_refs` / `impact_refs` / `outputs` の workflow 外 reference rule は既存 canonical reference 方針に従い、本 schema は ADR-092 により新しい reference resolution rule を追加しない。ただし、これらの field 自体は required list metadata として presence validation の対象とする。`work_items` / `source_requirement` / `tasks` / `work_item` / `depends_on` は workflow relation field として下記 relation integrity validation の対象とする。

> 由来: ADR-086 §4〜§7, ADR-087 §5〜§8, ADR-091 §6, ADR-092 §3〜§6

## Authoring guidance source model

Authoring guidance は Design Records record model とは別の read-only guidance source として扱う。

Guide source directory:

```text
docs/guides/
```

Guide file discovery rule:

```text
docs/guides/*.md
```

Guide ID は filename stem とする。

例:

```text
docs/guides/adr-authoring.md -> adr-authoring
```

Guide title は first H1 から抽出する。Guide abstract は `## Abstract` section の本文から抽出する。

`list_authoring_guides` は `id` / `title` / `abstract` のみを返す。`get_authoring_guidance` は `id` / `title` / `content` を返す。

Guide source file path は public response contract に含めない。Path は guide ID から内部解決する implementation detail である。

Guide は record kind ではない。Design Records record ID、record status、record path、record headings、record diagnostics、canonical reference resolver target として扱わない。

## Authoring transaction schema concepts

Authoring transaction は Design Records record model そのものではない。
Proposal / body cache は retained operational object であり、`list_records` / `get_record` / `get_records` / `resolve_reference` の record target には含めない。

### Authoring target identity

Authoring target identity は artifact identity を primary key とする。
Public request は physical path を primary input として受け取らない。

Authoring target identity fields:

| field | meaning |
|---|---|
| `kind` | target record kind |
| `requested_id` | caller input ID。Create では `new` placeholder を含んでよい |
| `resolved_id` | MCP が record index から解決した final ID |
| `domain` | workflow artifact domain。Requirement / work item / task create で使用する |
| `parent_id` | parent-aware ID resolution に使った parent record ID。Task create では required |
| `path` | resolved repository-relative path。Transparency output only |

`path` は relocation や slug generation の結果を説明するための output であり、authoring request の canonical target identity ではない。

Create operation の canonical target ID input は top-level request `id` である。
Structured `fields.id` is not part of the primary authoring target identity and is not required.
If a create request supplies `fields.id`, it is a duplicate consistency input only: it must match the exact top-level ID after canonical ID normalization, and it must be omitted when the top-level ID uses a `new` placeholder.
Mismatch or placeholder-time `fields.id` is an invalid request, not a record validation diagnostic.

For domain-scoped workflow artifact creates, request `domain` comparison with the ID domain is case-insensitive.
Canonical record IDs keep their uppercase domain segment.
Repository-relative paths use the lowercase normalized domain directory.
For example, `domain: "mcp"` with `id: "REQ-MCP-011"` resolves to canonical domain `MCP` and repository path domain `mcp`.

> 由来: REQ-MCP-011, TASK-MCP-011-01

### Proposal model

Proposal は write candidate の retained representation である。
Proposal creation does not write repository files.

Proposal fields:

| field | meaning |
|---|---|
| `proposal_id` | opaque lookup key |
| `state` | `proposed` / `accepted` / `discarded` |
| `operation` | `create` / `update` |
| `target_kind` | resolved target kind |
| `target` | authoring target identity |
| `base_state` | accept-time staleness detection に必要な target file / index state |
| `diff` | previewable diff |
| `validation` | proposal-time validation result |
| `required_follow_up_updates` | acceptance before satisfaction が forbidden な required follow-up list |
| `expires_at` | proposal expiry timestamp |
| `retention_days` | `3` |

`base_state` の concrete hash / timestamp / index snapshot shape は implementation detail である。
Public contract は、accept が stale target / changed target / ID collision を write 前に検出し、`written: false` と diagnostics を返すことである。

Proposal retention is 3 days.
Expired proposals are not valid authoring targets.

Proposal `validation` is scoped to the proposal-local affected record set in the candidate repository state.
The affected record set is the proposed target record plus any related records whose files are actually modified by the same proposal, such as required reciprocal workflow metadata updates.
Unrelated repository diagnostics are repository health information, not proposal-local diagnostics.
If exposed, repository health must be represented separately from proposal `validation` and must not change proposal-local `validation.ok`.

Proposal-local diagnostics must use the same validation categories and field contracts as `validate_records`.
They must be reproducible by the same `validate_records` rules against the same affected record set in the same candidate state, or after the candidate state has been accepted/materialized.

> 由来: REQ-MCP-012, TASK-MCP-011-01

### Body cache model

Body cache は、large Markdown body を proposal retry のために一時保持する operational object である。
It is not a design record and is not addressable by resolver.

Body cache fields:

| field | meaning |
|---|---|
| `body_cache_id` | opaque lookup key |
| `expires_at` | cache expiry timestamp |
| `retention_days` | `3` |

Operations requiring large Markdown body input accept exactly one of `body` or `body_cache_id`.
Supplying both is invalid and must not create a proposal or body cache.
Operations that do not require body input may omit both.
Unknown or expired body cache IDs must produce diagnostics and must not create proposals.
Body cache entries remain reusable within the 3 day retention period, including after they have been used to create a proposal.

For create operations, structured `fields` and a full Markdown body source (`body` or `body_cache_id`) are mutually exclusive content sources.
Supplying both is an invalid request; the schema does not define precedence between them.

### Metadata block replacement target

Metadata block replacement targets the kind-specific metadata block.

| kind | metadata block |
|---|---|
| `spec` | recognized spec metadata fields inside YAML front matter |
| `decision` | H1-following ADR bullet metadata block |
| `requirement` | H1-following requirement bullet metadata block |
| `work_item` | H1-following work item bullet metadata block |
| `task` | H1-following task bullet metadata block |

For `spec`, metadata replacement is scoped to recognized fields only. Unknown or auxiliary YAML front matter fields must be preserved. The recognized spec metadata fields are `scope`, top-level `status`, and `design_record.id` / `design_record.kind` / `design_record.status` / `design_record.depends_on`.

Required recognized fields are validated by the same field vocabulary used for record parsing and validation.
Missing required recognized fields produce `missing_required_metadata`.
Empty required scalar fields or empty list items produce `empty_required_metadata`.
Invalid recognized values produce `invalid_metadata_value` or an existing kind-specific diagnostic such as `invalid_status_for_kind`, `spec_status_mismatch`, or `invalid_migrated_to_spec`.

### Section selector model

Named section replacement uses an ATX heading selector.
The selector resolves within one Markdown record body.

MVP selector fields:

| field | meaning |
|---|---|
| `heading` | exact heading text after ATX marker removal and whitespace trim |
| `match` | `exact` only |
| `level` | optional ATX heading level constraint |

Matching is case-sensitive.
No Unicode normalization, punctuation folding, prefix matching, or contains matching is part of MVP.
If the selector resolves to zero sections, authoring returns `section_selector_no_match`.
If it resolves to multiple sections, authoring returns `section_selector_ambiguous`.
Neither case may create a proposal or write files.
Diagnostics should include candidate headings when possible.

Section selector resolution uses the same Markdown heading source rules as the `headings` field in the record model.
YAML front matter content and fenced code block content are not heading sources for section selectors.
Setext headings are not section sources in the MVP.

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
| `requirement` | `source_refs`, `work_items` |
| `work_item` | `source_requirement`, `impact_refs`, `tasks` |
| `task` | `work_item`, `source_requirement`, `estimate`, `depends_on`, `outputs` |

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
| `requirement` | `docs/requirements/*/REQ-*-*.md` の Markdown file を requirement 候補として読む |
| `work_item` | `docs/work-items/*/WORK-*-*.md` の Markdown file を work item 候補として読む |
| `task` | `docs/tasks/*/TASK-*-*.md` の Markdown file を新形式 task 候補として読む。`docs/tasks/m*.md` は含めない |

`design_record` を持たない spec は index 対象外とする。
その場合も `missing_design_record` diagnostic は出さない。

spec の `design_record.kind` が `spec` 以外の場合、この版では index 対象外とし、diagnostic は出さない。

本specは後続の他の record kind 追加を制限しない。

> 由来: ADR-076 §bootstrap方針, ADR-077 §validate_records の責務, ADR-092 §1

## `id`

`id` は record の安定識別子である。

MVP で扱う ID 形式は以下とする。

| kind | ID 例 | 備考 |
|---|---|---|
| `decision` | `ADR-076` | ADR 番号を3桁ゼロ埋めで持つ |
| `spec` | `SPEC-design-records-mcp-schema` | spec 用の stable ID |
| `investigation` | `INV-MCP-001` | ADR-086 に従う domain-scoped ID |
| `requirement` | `REQ-MCP-003` | domain-scoped ID。`REQ-<DOMAIN>-NNN` |
| `work_item` | `WORK-MCP-003` | domain-scoped ID。`WORK-<DOMAIN>-NNN` |
| `task` | `TASK-MCP-003-01` | parent work item sequence と task sequence を含む ID。`TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>` |

`decision` record の canonical ID は、H1 の番号から `ADR-NNN` として導出する。
filename 先頭の番号は canonical ID との一致検査にのみ使う。
H1 が不正な場合は `invalid_h1_title` を出し、filename 由来の ID を canonical ID として採用しない。

canonical ID の番号と filename 先頭の番号が一致していなければならない。

Requirement / work item / task の canonical ID は metadata `id` と H1 の ID から取得し、両者および filename の ID prefix が一致しなければならない。Workflow ID の syntax が上記 grammar に従わない場合、または metadata / H1 / filename prefix が一致しない場合は `invalid_workflow_id` または `filename_id_mismatch` とする。

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
| `requirement` | requirement artifact |
| `work_item` | work item artifact |
| `task` | 新形式の短期 task artifact |

この表は record kind の閉じた列挙ではない。後続判断により他の kind を追加しうる。
MVP では legacy M-series task record / UC / impl note を record kind として扱わない。

> 由来: ADR-076 §MVP対象, ADR-087 §1, ADR-092 §1

## `status`

`status` の値域は `kind` ごとに異なる。

| kind | allowed status |
|---|---|
| `decision` | `proposed` / `accepted` / `superseded` |
| `spec` | `confirmed` / `draft` / `wip` |
| `investigation` | `investigating` / `concluded` / `superseded` |
| `requirement` | `captured` / `decision_needed` / `accepted` / `deferred` / `rejected` |
| `work_item` | `not_started` / `decision_pending` / `design_spec_pending` / `internal_design_pending` / `yaml_pending` / `implementation_pending` / `fixture_pending` / `verification_pending` / `done` / `blocked` |
| `task` | `todo` / `doing` / `blocked` / `done` |

`decision` record の `status` は ADR 箇条書きmetadataから読む。

`spec` record の canonical `status` は top-level YAML front matter の `status` とする。
`design_record.status` が存在する場合、top-level `status` と同値でなければならない。
不一致の場合は `spec_status_mismatch` とする。

`investigation` record の `status` は investigation 箇条書きmetadataから読む。
`requirement` / `work_item` / `task` record の `status` はそれぞれの箇条書きmetadataから読む。

`kind` に対して許可されない `status` は `invalid_status_for_kind` とする。

> 由来: ADR-076 §front matter 方針, ADR-077 §validate_records の責務, ADR-086 §5, ADR-091, ADR-092

## `depends_on`

`depends_on` という field 名は、design record dependency と task graph dependency で異なる意味を持つため、kind ごとに解釈する。

### Decision / spec dependency

`decision` record の `depends_on` は、その判断が依存する record ID の list として ADR 箇条書きmetadataから読む。

`spec` record の `depends_on` は、その仕様が依存する record ID の list として `design_record.depends_on` から読む。Spec top-level front matter の `depends_on` は doc-policy 用の出自 path listであり、record dependency としては扱わない。

Decision / spec dependency が参照できる canonical record ID-as-ref は `ADR-*` / `SPEC-*` / `INV-*` とする。したがって ADR / spec が investigation record (`INV-*`) に依存することは valid である。

存在しない ID を参照している場合、`missing_depends_on_target` とする。MVP では参照元・参照先の status 組み合わせは検査しない。

### Task dependency relation

`task` record の `depends_on` は、同 task の実行が依存する task artifact を指す workflow relation field として task 箇条書きmetadataから読む。Canonical target form は `TASK-*` とする。

`task.depends_on` は decision / spec dependency ではなく、`unresolved_workflow_relation` / `invalid_workflow_relation_target` の対象とする。MVP は参照先の存在確認を行うが、same-work-item 制約、cycle detection、execution order projection は扱わない。

> 由来: ADR-077 §validate_records の責務, ADR-091 §3・§6, ADR-092 §4・§7

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

Workflow artifact record の H1 は以下の形式のみ valid とする。

```text
^#\s+(?P<id>REQ-[A-Z0-9]+(?:-[A-Z0-9]+)*-\d{3}):\s+(?P<title>.+?)\s*$
^#\s+(?P<id>WORK-[A-Z0-9]+(?:-[A-Z0-9]+)*-\d{3}):\s+(?P<title>.+?)\s*$
^#\s+(?P<id>TASK-[A-Z0-9]+(?:-[A-Z0-9]+)*-\d{3}-\d{2}):\s+(?P<title>.+?)\s*$
```

Requirement / work item / task の canonical ID は metadata `id` と H1 の `id` が一致したとき、その ID とする。H1 または metadata `id` が grammar に従わない場合は `invalid_workflow_id`、両者または filename の ID prefix が一致しない場合は `filename_id_mismatch` とする。

filename からの title 推定は MVP では行わない。

> 由来: ADR-077 §list_records の責務, ADR-077 §理由, ADR-092 §3

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
| `requirement` | requirement bullet metadata | requirement-specific detail object |
| `work_item` | work item bullet metadata | work-item-specific detail object |
| `task` | task bullet metadata | task-specific detail object |
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

Authoring transaction tools additionally use the following diagnostic categories.
These diagnostics may appear in proposal, accept, discard, get-proposal, or body-cache retry responses.

| category | severity | meaning |
|---|---|---|
| `proposal_not_found` | error | requested proposal ID が存在しない |
| `proposal_expired` | error | requested proposal ID は expiry を過ぎている |
| `proposal_discarded` | error | proposal は discard 済みであり accept できない |
| `proposal_already_accepted` | error | proposal は accepted 済みであり再適用できない |
| `proposal_stale` | error | proposal base state と current target state が一致しない |
| `target_changed` | error | target record kind / path / identity が proposal 作成時と異なる |
| `id_collision` | error | create proposal の resolved ID が accept 前に使用済みになった |
| `required_follow_up_not_satisfied` | error | required reciprocal metadata update などの follow-up が満たされていない |
| `invalid_body_source` | error | body source rule 違反。`body` と `body_cache_id` の両方指定、または required body source の欠落。Create operation の `fields` と full body source の同時指定は request shape violation として `invalid_request` を用いる |
| `body_cache_not_found` | error | requested body cache ID が存在しない |
| `body_cache_expired` | error | requested body cache ID は expiry を過ぎている |
| `proposal_preparation_failed` | error | proposal preparation failed before proposal persistence |
| `section_selector_no_match` | error | named section selector が target record 内の section に一致しない |
| `section_selector_ambiguous` | error | named section selector が複数 section に一致し、単一 target に解決できない |

`section_selector_no_match` / `section_selector_ambiguous` diagnostics should include `candidate_headings` when possible.
Candidate heading entries contain at least `heading`, `level`, and `ordinal`.

MVP の `validate_records` は以下の diagnostic category を返す。

Diagnostic は検査軸ごとに独立して発火し、1 record に複数 diagnostic が付いてよい。

| category | severity | meaning |
|---|---|---|
| `duplicate_id` | error | 複数 record が同じ正規化後 record ID を持つ |
| `filename_id_mismatch` | error | `decision` / `investigation` / workflow artifact record の canonical ID または metadata ID と filename ID 部分が一致しない |
| `invalid_h1_title` | error | H1 が存在しない、または期待形式に合わない |
| `invalid_workflow_id` | error | requirement / work item / task の metadata ID または H1 ID が workflow ID grammar に従わない |
| `missing_required_metadata` | error | requirement / work item / task の required metadata field が存在しない |
| `empty_required_metadata` | error | required scalar metadata field が empty、または required list metadata field に empty item が含まれる |
| `missing_required_section` | error | workflow artifact が gated status にあるとき、required narrative section heading が存在しない |
| `empty_required_section` | error | workflow artifact が gated status にあるとき、required narrative section heading は存在するが section body が empty または whitespace-only である |
| `invalid_metadata_value` | error | required metadata field が non-empty だが value contract を満たさない。例: workflow artifact `date` が strict `YYYY-MM-DD` format ではない |
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
| `unsupported_reference` | error / info | MVP が unsupported と定義する investigation metadata reference。`source_refs` / `follow_up_results` では error、`follow_up_candidates` では info。Investigation field に現れる `TASK-*` を含む。Reserved `yaml:` はこの category の対象に含めず、MVP では behavior を定義しない |
| `unresolved_workflow_relation` | error | workflow relation field に記載された supported `REQ-*` / `WORK-*` / `TASK-*` が解決不能 |
| `invalid_workflow_relation_target` | error | workflow relation field に、field が要求する kind / ID form ではない target が記載された |
| `workflow_relation_mismatch` | error | `REQ.work_items` と `WORK.source_requirement`、または `WORK.tasks` と `TASK.work_item` の宣言済み双方向 relation が一致しない |
| `workflow_source_requirement_mismatch` | error | task の `source_requirement` が parent work item の `source_requirement` と一致しない |

Workflow metadata diagnostic (`missing_required_metadata` / `empty_required_metadata` / `invalid_metadata_value`) は、通常の `category` / `severity` / `record_id` / `path` / `message` に加えて、`field` を必須で返す。入力 value が存在する場合は `value` も返す。

Workflow required section diagnostic (`missing_required_section` / `empty_required_section`) は、通常の `category` / `severity` / `record_id` / `path` / `message` に加えて、`section` と `status` を必須で返す。
`section` は required narrative section heading text とする。
`status` は required section rule を発火させた workflow artifact status とする。

Workflow relation diagnostic は、通常の `category` / `severity` / `record_id` / `path` / `message` に加えて、`field`、`value`、`ref_status` を必須で返す。`field` は `work_items` / `source_requirement` / `tasks` / `work_item` / `depends_on` のいずれか、`value` は入力 ID-as-ref、`ref_status` は `unresolved` / `invalid_target` / `mismatch` のいずれかとする。対象 ID が特定できる場合は `target_id` も返す。

Investigation reference diagnostic (`unresolved_*` / `noncanonical_*` / metadata field 由来の `unsupported_reference`) は、通常の `category` / `severity` / `record_id` / `path` / `message` に加えて、`field`、`value`、`ref_status` を必須で返す。`field` は `source_refs` / `follow_up_results` / `follow_up_candidates` のいずれか、`value` は入力 reference 文字列、`ref_status` は `unresolved` / `unsupported` / `noncanonical` のいずれかとする。Investigation metadata が duplicate semantic ref または duplicate record ID を指して単一解決できない場合は field-specific diagnostic を追加せず、index defect を示す `duplicate_semantic_ref` または `duplicate_id` のみを返す。これら duplicate diagnostic および spec declaration / section lookup diagnostic は investigation metadata field 由来の追加 field を要求しない。

Required narrative section policy:

| artifact kind | gated status | required non-empty narrative sections |
|---|---|---|
| `work_item` | `done` | `Goal`, `Boundary`, `Evidence` |
| `task` | `done` | `Goal`, `Work`, `Done condition`, `Verification`, `Evidence` |
| `requirement` | `accepted` | `Requirement`, `Required Outcome` |

`requirement` の `accepted` は close/completion state ではなく adoption-readiness gate として扱う。
したがって `Evidence` / `Boundary` / `Explicitly Excluded Scope` は `REQ accepted` の required non-empty section には含めない。

Required narrative section body は、heading 行を除いた section body の前後 whitespace を trim した結果、少なくとも 1 つの non-whitespace character を含む場合に non-empty とする。
Whitespace-only body は empty とする。
本文の品質・十分性・意味内容は判定しないため、`Pending` や `None` のような placeholder text も non-empty として扱う。

`follow_up_candidates` の参照先が未作成であること自体は error としない。`validate_records.ok` は error diagnostic の有無だけで決まり、info diagnostic が存在しても failure にはならない。Coverage mapping、semantic realization relation、`internal-design:` / `coverage:` / `COV-*` の解決・診断は MVP diagnostic scope に含めない。Workflow relation validation は宣言済み relation の存在と整合性に限り、未接続 artifact の orphan diagnostics、task dependency cycle detection、execution order projection、task status 由来 progress projection は含めない。

MVP では以下を diagnostic category に含めない。

- `accepted_but_not_migrated`
- `missing_design_record`
- status 組み合わせの妥当性
- spec section 単位の由来不足
- 自然言語本文と metadata の意味的不一致
- orphan requirement / orphan work item / orphan task
- task dependency cycle / execution order projection
- task status 由来 progress projection

`missing_record_path` は、filesystem scan または path normalization により record 候補 path を検出したが、実際の read/stat に失敗した場合に出す。
例として、scan 後に file が削除された場合、permission denied、symlink target missing、path normalization 後の path が存在しない場合を含む。

> 由来: ADR-077 §validate_records の責務, ADR-090 §Partial result / Ordering と duplicate requested ID, ADR-092 §4〜§7, REQ-MCP-017 / TASK-MCP-016-01 required narrative section policy

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

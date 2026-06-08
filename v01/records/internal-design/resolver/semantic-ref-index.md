---
scope: docs/internal-design/resolver/semantic-ref-index.md
status: draft
last_updated: 2026-05-25
summary: >
  Design Records MCP が canonical semantic/artifact ref を解決するための
  index、record integration、validation route の implementation-facing design を記録する。
depends_on:
  - docs/adr/087-design-records-mcp-investigation-support-and-semantic-ref-resolve.md
  - docs/adr/088-reduce-semantic-trace-mvp-to-canonical-reference-resolution-foundation.md
  - docs/spec/design-records-mcp/schema.md
  - docs/spec/design-records-mcp/tools.md
---

# Semantic ref index resolver

## Purpose

この document は、M19 Phase A で確定した Design Records MCP public contract を、既存 Go implementation へ追従させるための internal design である。

Public tool 名、request / response schema、diagnostic category と severity の正は `docs/spec/design-records-mcp/{schema,tools}.md` が所有する。本 document はそれらを変更せず、parser / discovery / index / resolver / validator / MCP transport の実装 route を具体化する。

V01-ADR-088 により、本 document 自体は semantic trace MVP の active endpoint または required acceptance target ではない。

## Finalized implementation scope

M19 で実装する public behavior は以下に限定する。

| area | required implementation |
|---|---|
| record integration | `decision` / `spec` / `investigation` の index / list / get / validate 対応 |
| response shape | common fields + kind-specific detail object。旧 flat metadata response は返さない |
| public resolver | `resolve_reference` を追加し、active `spec:` と `ADR-*` / `SPEC-*` / `INV-*` を解決する |
| semantic ref validation | active `spec:` declaration grammar / duplicate / section lookup を検査する |
| investigation validation | `source_refs` / `follow_up_results` / `follow_up_candidates` の M19 対象 reference を診断する |
| known gap closure | `V01-ADR-086` の `depends_on: V01-INV-DOCS-001` が investigation integration 後に解決する |

M19 で実装 behavior を定義しないもの:

- reserved `yaml:` の public resolver input / direct query response / investigation metadata validation
- `internal-design:` / `coverage:` / `COV-*` の resolve または validation
- `REQ-*` / `WORK-*` / task identity の record / resolver support
- coverage mapping / semantic realization relation
- batch retrieval (`get_records` 等)
- MCP writer tools

## Current implementation map and gaps

現行実装は二層に分かれている。

| package / file | current responsibility | M19 gap |
|---|---|---|
| `internal/designrecords/types.go` | record / diagnostic / tool request-response type | `investigation` kind、detail object、`info` severity、resolver type、reference diagnostic fields がない |
| `internal/designrecords/parser.go` | ADR bullet metadata と spec front matter の parse | investigation bullet metadata と active `spec:` declaration parse がない |
| `internal/designrecords/index.go` | `docs/adr/*.md` と `docs/spec/**/*.md` の discovery | `docs/investigations/*/INV-*-*.md` discovery と reference index material がない |
| `internal/designrecords/tools.go` | list / get / validate / suggest の domain handler | response が flat、mixed-kind ordering が旧仕様、`resolve_reference` がない |
| `internal/designrecords/validation.go` | duplicate ID、status、depends_on / supersedes 等の検査 | investigation / semantic ref / noncanonical reference diagnostics がない |
| `internal/designrecordsmcp/tools.go` | MCP tool schema 公開 | `investigation` enum と `resolve_reference` schema がない |
| `internal/designrecordsmcp/tools_call.go` | strict JSON argument decode と domain handler dispatch | `resolve_reference` dispatch がない |

実装順は domain package (`internal/designrecords`) を先に閉じ、transport package (`internal/designrecordsmcp`) を最後に public tool surface へ追従させる。Transport に parsing / resolution rule を重複実装しない。

## Target record model

### Common and detail representation

`Record` と public response DTO は、common fields と kind-specific detail object に正規化する。

Common fields:

- `id`
- `kind`
- `title`
- `status`
- `path`

Detail object:

| kind | detail fields |
|---|---|
| `decision` | `depends_on`, `supersedes`, `migrated_to_spec` |
| `spec` | `depends_on` |
| `investigation` | `trigger`, `scope`, `non_scope`, `source_refs`, `follow_up_candidates`, optional `supersedes`, `related_*`, `follow_up_results` |

`headings` と requested raw `body` は `get_record` のみに追加する common retrieval content とする。

### Migration boundary

現行 `Record`, `ListedRecord`, `GetRecordRecord` は `DependsOn` / `Supersedes` / `MigratedToSpec` を flat field として公開している。M19 では以下を同一変更単位で切り替える。

1. type 定義を detail object shape に更新する。
2. ADR / spec parser が kind detail を生成するように更新する。
3. investigation parser が investigation detail を生成する。
4. `listedRecord` / `getRecordResponseRecord` の mapper を新 shape 専用に更新する。
5. existing list / get tests を新 shape へ更新し、flat field が返らない test を追加する。

旧 flat public response と新 detail response の互換併存期間は設けない。内部で parse 補助構造を用いる場合も、public DTO に flat metadata を残さない。

## Phase B: investigation record integration route

### Record kind and status

`RecordKindInvestigation = "investigation"` を追加する。

Investigation status は以下のみを許容する。

- `investigating`
- `concluded`
- `superseded`

`newListRecordsScope` / `newValidationScope` / MCP tool input schema の kind enum は `investigation` を受け付ける。`id_range` は引き続き `decision` 専用とし、investigation range filter は追加しない。

### Discovery

`BuildIndex` に `discoverInvestigationRecords` を追加する。Discovery route は以下とする。

1. repository root 配下の `docs/investigations/` を起点に再帰 walk する。
2. Markdown file のうち filename が `INV-*-*.md` 形式のものだけを investigation candidate として読む。
3. response path は既存と同じ repository-relative slash path に正規化する。
4. read/stat failure は既存 `PathIssue` route を用いて `missing_record_path` に接続する。

M19 は requirement / work item / task directory を discovery 対象に追加しない。

### Investigation parsing

Investigation source は V01-ADR-086 / schema に従う H1 直下の bullet metadata である。`parseInvestigationRecord` は ADR parser と同じ metadata block boundary（H1 後から最初の H2 または blockquote 直前まで）を使い、以下を読む。

Required scalar metadata:

- `status`
- `date`（parse してよいが public response へ出さない）
- `trigger`
- `scope`
- `non_scope`

Required list metadata:

- `source_refs`
- `follow_up_candidates`

Optional list metadata:

- `supersedes`
- `related_requirements`
- `related_work_items`
- `related_adrs`
- `related_specs`
- `related_internal_design`
- `related_coverage`
- `follow_up_results`

ADR の comma-separated list と異なり、investigation list field は metadata key の次に続く indented Markdown list item (`  - value`) を収集する。入力値は trim するが、reference validation 前に path や prefix を正規化・変換しない。

H1 は `# INV-<DOMAIN>-NNN: <title>` 形式で読み、H1 の ID を canonical ID とする。Filename の `INV-<DOMAIN>-NNN` prefix と一致しない場合は既存 `filename_id_mismatch` category を再利用する。

### Index output

investigation record が index に入ると、既存 `recordsByID` に `INV-*` が現れる。これにより、既存 dependency validation route は `V01-ADR-086 -> V01-INV-DOCS-001` を record ID-as-ref として解決でき、現行 `missing_depends_on_target` baseline error を自然に解消する。

## Phase C: shared reference resolution route

### Index materials

`Index` は record list に加え、resolver / validator が共有する lookup material を保持する。

| lookup material | key | target |
|---|---|---|
| record ID index | normalized `ADR-*` / `SPEC-*` / `INV-*` | `[]Record` |
| semantic ref index | exact active `spec:` ref string | `[]SemanticTarget` |

`[]` で保持するのは duplicate を fail-closed に扱うためである。Resolver は duplicate の際に任意 target を選ばず、validator は index defect diagnostic を返す。

### Active `spec:` extraction

Spec parser は既存 `design_record` metadata に加え、front matter の active semantic reference declaration を読む。

- document-level: `semantic_refs`
- section-level: `sections` mapping (`spec:` key -> heading text)

Section target validation は `extractHeadings` が返す Markdown heading を利用し、front matter / fenced code block 内の heading-like text を target にしない。

### Shared resolver service

Domain package に単一の resolver route を置き、`resolve_reference` と `validate_records` が同じ lookup function を使う。

Resolver classification order:

1. active `spec:` grammar に一致する入力を semantic ref lookup へ送る。
2. `ADR-*` / `SPEC-*` / `INV-*` grammar に一致する入力を record ID lookup へ送る。
3. M19 が explicit unsupported と定義する prefix / ID / physical path は `unsupported_reference` として返す。
4. reserved `yaml:` は M19 で behavior を定義しないため、resolver / metadata validator の実装対象にしない。

Resolved target shape:

| input | target output |
|---|---|
| document-level `spec:` | `target_type: document`, `path` |
| section-level `spec:` | `target_type: section`, `path`, `section` |
| record ID-as-ref | `target_type: record`, `path`, `record_id`, `record_kind`, `title`, `status` |

Supported form だが target がない場合は `unresolved_reference`。同一 input に複数 target がある場合は `ambiguous_reference` を返し、target は返さない。

### Duplicate and ambiguity handling

Validation は duplicate の原因を index defect として報告する。

- duplicate record ID: `duplicate_id`
- duplicate semantic ref declaration: `duplicate_semantic_ref`
- `sections` value が同一 document 内の複数 heading と一致する: `ambiguous_section_target`

Investigation metadata が duplicate target を指しても、`ambiguous_source_ref` 等の field-specific category は追加しない。Duplicate index defect の diagnostic のみを返す。

## Phase C: validation route

### Existing validation retained

以下の現行 route は detail object へ field access を移した上で維持する。

- `duplicate_id`
- `filename_id_mismatch`
- `invalid_h1_title`
- `invalid_status_for_kind`
- `spec_status_mismatch`
- `missing_depends_on_target`
- `missing_supersedes_target`
- `invalid_migrated_to_spec`
- `missing_record_path`

### New semantic ref diagnostics

Spec declaration / lookup validation は以下を返す。

| diagnostic | trigger |
|---|---|
| `invalid_semantic_ref_declaration` | `semantic_refs` entry または `sections` key が active `spec:` grammar に合わない |
| `missing_section_target` | `sections` value に一致する Markdown heading が存在しない |
| `ambiguous_section_target` | 同一 spec document 内で複数 heading に一致する |
| `duplicate_semantic_ref` | 同一 active `spec:` ref が複数 target に宣言された |

### Investigation reference diagnostics

Investigation validation は `investigation` detail の以下 field のみに適用する。

| field | unresolved supported canonical ref | physical path | explicit unsupported M19 ref form |
|---|---|---|---|
| `source_refs` | `unresolved_source_ref` error | `noncanonical_source_ref` error | `unsupported_reference` error |
| `follow_up_results` | `unresolved_follow_up_result` error | `noncanonical_follow_up_result` error | `unsupported_reference` error |
| `follow_up_candidates` | `unresolved_follow_up_candidate` info | `noncanonical_follow_up_candidate` info | `unsupported_reference` info |

これらの investigation field diagnostic は `field`、`value`、`ref_status` を必須で返す。`ref_status` は `unresolved` / `unsupported` / `noncanonical` のみとする。

`trigger` / `related_*`、および reserved `yaml:` の metadata validation behavior は M19 では実装しない。

### Diagnostic type migration

現行 `DiagnosticSeverity` は `error` のみであり、`Diagnostic` は reference field location を保持しない。M19 では以下を追加する。

- `DiagnosticSeverityInfo = "info"`
- optional `Field` JSON field
- optional `Value` JSON field
- optional `RefStatus` JSON field

`ValidateRecordsResponse.OK` の判定は現行の error scan を維持し、`info` のみでは `false` にしない。

## Tool and transport integration

### Domain handlers

`internal/designrecords` に以下を実装する。

| handler | implementation route |
|---|---|
| `ListRecords` | investigation filter と新 detail response を返す。mixed-kind `order_by: id` は canonical ID ASCII lexical、同一 ID は path lexical で安定化する |
| `GetRecord` | new detail response と optional raw body を返す。duplicate ID を暗黙に first-match 解決する既存 behavior は resolver の fail-closed rule と混同せず、contract test に従い必要なら見直す |
| `ResolveReference` | shared resolver service を呼び出し public resolution response を返す |
| `ValidateRecords` | shared lookup material を用いて既存 + semantic / investigation diagnostics を返す |

### MCP transport

`internal/designrecordsmcp` は business rule を持たず、次のみを行う。

- `tools.go`: `list_records` / `validate_records` の kind enum に `investigation` を追加し、`resolve_reference` input schema を公開する。
- `tools_call.go`: `resolve_reference` request を strict decode し、domain handler に dispatch する。
- transport tests: tool listing / schema / dispatch / unknown field rejection を更新する。

Resolver grammar、target selection、diagnostic severity は domain package にのみ実装する。

## Implementation sequence

Phase B/C の code change は以下の順で行う。

1. `types.go`: kind / status / detail DTO / resolver DTO / diagnostic fields と severity を追加し、public response model を切り替える。
2. `parser.go`（必要なら補助 file 分割）: investigation metadata parser と active `spec:` declaration parse を追加する。
3. `index.go`: investigation discovery と resolver lookup material 構築を追加する。
4. `tools.go`: list / get mapper、ordering、`ResolveReference` handler を実装する。
5. `validation.go`: shared resolution を用いる semantic / investigation validation と info severity を実装する。
6. `internal/designrecordsmcp/{tools.go,tools_call.go}`: public tool schema / dispatch を更新する。
7. tests: parser/index、list/get response migration、resolver、validation、transport、known baseline error 解消を追加・更新する。

この順は public DTO の flat/detail 併存を避けるため、実装 commit 内では package tests と同時に切り替える。途中状態を public MCP として提供しない。

## Verification targets

Concrete checklist は M19 Phase D が所有する。B0 から implementation へ渡す必須確認点は以下である。

- `V01-INV-DOCS-001` を investigation record として取得できる。
- `V01-ADR-086 -> V01-INV-DOCS-001` の既知 `missing_depends_on_target` error が消える。
- `spec:` document / section resolve が contract shape で返る。
- `ADR-*` / `SPEC-*` / `INV-*` resolve が contract shape で返る。
- invalid / missing / ambiguous / duplicate semantic ref declaration が定義済み category で診断される。
- investigation unresolved / noncanonical / explicit unsupported input の severity boundary が field ごとに守られる。
- duplicate target を指す investigation metadata に field-specific ambiguous diagnostic を追加しない。
- `yaml:`、REQ / WORK / task、coverage / internal-design relation を acceptance 対象に持ち込まない。
- list / get が old flat metadata field を返さない。

## Non-goals

- Public MCP contract の追加変更
- `internal-design:` / `coverage:` / `yaml:` endpoint の active 化
- `yaml:` metadata diagnostic の定義
- requirement / work item / task MCP support
- multiple-record batch read tool
- semantic realization relation validation
- coverage mapping query tool
- MCP writer tools

> 由来: V01-ADR-083, V01-ADR-087, V01-ADR-088; V01-INV-DOCS-002; V01-INV-DOCS-003; M19 Phase A contract result

# Milestone 16: Design Records MCP MVP

- **status**: closed
- **scope**: docs/spec/design-records-mcp / docs/adr / docs/spec / cmd/design-records-mcp / internal/designrecords / tests
- **source**: ADR-076 (Design Records MCP) / ADR-077 (MVP boundary and tool prioritization) / docs/spec/design-records-mcp/**
- **last_updated**: 2026-05-12

---

## Context

ADR-050 により、brewprint のドキュメント運用は spec-first に転換した。
現行仕様は `docs/spec/**` に置き、ADR は設計判断の根拠記録として扱う。

ADR-076 では、ADR/spec 運用を machine-readable metadata と MCP query / validation で支援する Design Records MCP の導入方針を決めた。
ADR-077 では、MVP の tool 優先度と read-only 境界を決めた。

本 milestone は、`docs/spec/design-records-mcp/**` に定義した MVP 仕様を実装に落とし込む。
対象は ADR / spec record の index / read / validation であり、既存 brewprint MCP の `ResolvedProject` / `QueryService` とは別責務として扱う。

---

## Implementation layout

MVP implementation uses the following layout:

```text
cmd/
  brewprint/
  design-records-mcp/

internal/
  designrecords/
```

`cmd/design-records-mcp/` は Design Records MCP の独立起動 binary とする。

`internal/designrecords/` が所有するもの:

- ADR/spec Markdown metadata parsing
- record index construction
- validation
- Design Records MCP tool handlers

Non-goals:

- 既存 `internal/mcp` に Design Records MCP を混ぜない
- 既存 `internal/query` / `ResolvedProject` を docs 管理へ拡張しない
- MVP では `cmd/brewprint` の subcommand として統合しない

---

## Non-goals

M16 MVP では以下を行わない。

- `trace_record`
- `list_gaps`
- `create_record`
- `update_record`
- `set_evidence`
- その他 write 系 tool
- task file / UC docs / impl notes の index 化
- 自然言語本文から依存関係を推定すること
- spec 本文との厳密な意味照合
- git 履歴解析
- code static analysis
- Web UI
- 複数プロジェクト横断管理
- 汎用 OSS CLI としての公開 contract 整備
- section 単位の完全な traceability
- `topics` / `affects` / `refines` / `conflicts_with` metadata

---

## Phase 0: implementation skeleton

Design Records MCP を既存 brewprint MCP から独立して起動できる構成で立ち上げる。

- [x] `cmd/design-records-mcp/` を作成する
- [x] `internal/designrecords/` を作成する
- [x] repository root を cwd または起動引数から決める
- [x] `docs/adr/*.md` と `docs/spec/**/*.md` を読むための filesystem boundary を実装する
- [x] 既存 `internal/mcp` / `internal/query` に依存しない構成にする

Done criteria:

- [x] `go run ./cmd/design-records-mcp` 相当で独立起動できる
- [x] root path を指定または cwd から解決できる
- [x] 既存 brewprint YAML semantic build に依存していない

---

## Phase 1: record parser and index

ADR/spec Markdown から MVP record index を構築する。

### ADR parsing

- [x] `docs/adr/*.md` を `decision` record 候補として scan する
- [x] ADR H1 を `^#\s+(?P<num>\d{3}):\s+(?P<title>.+?)\s*$` で parse する
- [x] canonical ID を `ADR-<num>` として導出する
- [x] ADR files with missing or invalid H1 remain validation candidates so `invalid_h1_title` can be emitted
- [x] When ADR H1 is invalid, do not derive the canonical record ID from the filename
- [x] filename 先頭番号と H1 番号を3桁ゼロ埋め文字列一致で検査する
- [x] H1 直後の bullet metadata block を parse する
- [x] ADR metadata block starts after H1 and ends before the first H2 line or blockquote line; empty lines are allowed inside the block
- [x] ADR metadata lines are recognized only in the form `- **<key>**: <value>`; bold marker is required
- [x] ADR metadata keys are case-sensitive; values are trimmed; empty or whitespace-only values are treated as unspecified
- [x] `status` / `date` / `depends_on` / `supersedes` / `migrated_to_spec` key を認識する
- [x] 未認識 key は MVP では無視する
- [x] `depends_on` / `supersedes` は comma 区切り list として正規化する
- [x] 空 `depends_on` / `supersedes` は empty list に正規化する
- [x] 空 `migrated_to_spec` は null に正規化する
- [x] non-empty `migrated_to_spec` は `YYYY-MM-DD` のみ有効とする
- [x] `date` は parse してよいが record field には含めない

### Spec parsing

- [x] `docs/spec/**/*.md` を scan する
- [x] YAML front matter を parse する
- [x] `design_record.id` と `design_record.kind` を持つ file のみ spec record として index する
- [x] `design_record` を持たない spec は silent skip し、`missing_design_record` diagnostic は出さない
- [x] `design_record.kind` が `decision` / `spec` 以外の場合も MVP では silent skip する
- [x] spec record の canonical status は top-level front matter `status` とする
- [x] `design_record.status` が存在する場合、top-level `status` と一致するか検査する
- [x] top-level `depends_on` は doc-policy 用 path list として扱い、record dependency には使わない
- [x] record dependency は `design_record.depends_on` から読む
- [x] spec record の `supersedes` は empty list に正規化する
- [x] spec record の `migrated_to_spec` は null に正規化する
- [x] spec の `design_record.supersedes` / `design_record.migrated_to_spec` に値があっても MVP では無視し、diagnostic は出さない

### Common parsing

- [x] `title` は H1 から抽出する
- [x] spec record title は H1 行から leading `#` と whitespace を除き、前後 whitespace を trim する
- [x] `headings` は ATX heading のみ抽出する
- [x] YAML front matter 内 / fenced code block 内の `#` は heading として扱わない
- [x] setext heading は MVP では扱わない
- [x] Spec records with missing or invalid ATX H1 emit `invalid_h1_title`; title is not inferred from filename or front matter
- [x] `body` は raw Markdown file content として保持または必要時取得できるようにする

Done criteria:

- [x] ADR-050 / ADR-067〜ADR-077 が decision record として index できる
- [x] `docs/spec/design-records-mcp/**` が spec record として index できる
- [x] `design_record` を持たない既存 spec は silent skip される
- [x] raw body は整形・要約・正規化されない

---

## Phase 2: list_records

`list_records` tool を実装する。

- [x] `kind` filter を実装する
- [x] `status` filter を実装する
- [x] `id` exact filter を実装する
- [x] `id_range` filter を実装する
- [x] `id_range` endpoints are inclusive
- [x] `id_range.from` and `id_range.to` are independently optional; one-sided ranges are supported
- [x] `id_range` は `ADR-NNN` の decision record 専用とする
- [x] `kind` 省略 + `id_range` 指定時は `kind: decision` と同等に扱う
- [x] `kind: spec` と `id_range` の併用は request error とする
- [x] `SPEC-*` range 指定は request error とする
- [x] `order_by: id` を実装する
- [x] `order: asc | desc` を実装する
- [x] `limit` を実装する
- [x] `head` / `tail` は実装しない

Response fields:

- [x] `id`
- [x] `kind`
- [x] `title`
- [x] `status`
- [x] `path`
- [x] `depends_on`
- [x] `supersedes`
- [x] `migrated_to_spec`

Done criteria:

- [x] 最新 ADR を `order_by:id`, `order:desc`, `limit` で取得できる
- [x] ADR-067〜ADR-077 を `id_range` で取得できる
- [x] `kind: spec` で Design Records MCP spec record を取得できる
- [x] `kind: spec` + `id_range` は request error になる

---

## Phase 3: validate_records

`validate_records` tool を実装する。

Diagnostic category:

- [x] `duplicate_id`
- [x] `filename_id_mismatch`
- [x] `invalid_h1_title`
- [x] `invalid_status_for_kind`
- [x] `spec_status_mismatch`
- [x] `missing_depends_on_target`
- [x] `missing_supersedes_target`
- [x] `invalid_migrated_to_spec`
- [x] `missing_record_path`

Validation rules:

- [x] duplicate normalized record ID を検出する
- [x] ADR H1 番号と filename 番号の不一致を検出する
- [x] ADR H1 が期待形式に合わない場合 `invalid_h1_title` を出す
- [x] kind 別 status 値域違反を検出する
- [x] spec top-level `status` と `design_record.status` の不一致を `spec_status_mismatch` とする
- [x] `depends_on` 参照先 ID の存在確認を行う
- [x] `supersedes` 参照先 ID の存在確認を行う
- [x] ADR `migrated_to_spec` の non-empty 値が `YYYY-MM-DD` でない場合 `invalid_migrated_to_spec` を出す
- [x] scan/path normalization 後の candidate path に対する read/stat 失敗を `missing_record_path` とする
- [x] Diagnostic は検査軸ごとに独立して発火し、1 record に複数 diagnostic が付いてよい

MVP out:

- [x] `accepted_but_not_migrated` は実装しない
- [x] `missing_design_record` は実装しない
- [x] status combination validation は実装しない
- [x] semantic mismatch between body and metadata は実装しない
- [x] spec section origin completeness は実装しない

Done criteria:

- [x] `validate_records` が全 record を検証できる
- [x] `kind` / `id_range` filter に対応する
- [x] `id_range` rule は `list_records` と同じ
- [x] validation response returns top-level `ok` and `diagnostics`; `ok` is true when there are no error diagnostics
- [x] diagnostic response は `category` / `severity` / `record_id` / `path` / `message` / `target_id` を返せる

---

## Phase 4: get_record

`get_record` tool を実装する。

- [x] `id` で record を取得する
- [x] `include_body` を受け取る
- [x] `include_body` default は false とする
- [x] `include_body=false` でも metadata / path / title / headings を返す
- [x] `include_body=true` の場合、raw Markdown body を追加する
- [x] body は整形・要約・正規化しない
- [x] metadata / headings は body とは別 field として返す
- [x] 存在しない ID は `record_not_found` tool error とする

Done criteria:

- [x] ADR-076 を ID から取得できる
- [x] `include_body=true` で元 Markdown 本文をそのまま返せる
- [x] `include_body=false` では body を返さない
- [x] headings に H1 / H2 などの ATX heading が含まれる

---

## Phase 5: tool error handling

MVP tool error code を実装する。

- [x] `record_not_found`
- [x] `invalid_request`
- [x] `unsupported_kind`
- [x] `id_range_requires_decision_kind`

Rules:

- [x] `list_records` に `kind: task` など不正な kind を指定した場合は `invalid_request` とする
- [x] If `suggest_next_record` is implemented, `kind: spec` など対象外 kind を指定した場合は `unsupported_kind` とする
- [x] `kind: spec` と `id_range` の併用は `id_range_requires_decision_kind` とする
- [x] `SPEC-*` range 指定は `id_range_requires_decision_kind` とする

Done criteria:

- [x] request validation error が machine-readable に返る
- [x] tool execution error と validation diagnostic を混同しない

---

## Phase 6: optional suggest_next_record

P0 完了後、余力があれば `suggest_next_record` を実装する。

- [x] 対象は `kind: decision` のみとする
- [x] existing max decision record number を取得する
- [x] `next_number = max + 1` とする
- [x] 欠番は埋めない
- [x] `next_id = ADR-NNN` を返す
- [x] `suggested_path` を返す
- [x] filename slug を title から生成する
- [x] slug は ASCII 英数字 lowercase、非英数字を `-`、連続 `-` を1つ、前後 `-` 除去、非 ASCII は `-` とする
- [x] slug が空の場合、`docs/adr/{NNN}.md` を suggested path としてよい
- [x] ファイル作成は行わない

Done criteria:

- [x] 既存最大 ADR が 077 の場合、078 を提案できる
- [x] suggested path を返せる
- [x] ファイル副作用がない

---

## Phase 7: tests and fixtures

MVP parser / tools の regression test を追加する。

- [x] ADR H1 parse tests
- [x] ADR bullet metadata parse tests
- [x] spec YAML front matter / design_record parse tests
- [x] record index tests
- [x] list_records filter / sort / limit tests
- [x] validate_records diagnostic tests
- [x] get_record include_body tests
- [x] tool error handling tests
- [x] optional: suggest_next_record tests

Test fixtures should cover:

- [x] valid ADR H1
- [x] invalid ADR H1
- [x] filename ID mismatch
- [x] empty depends_on / supersedes / migrated_to_spec
- [x] invalid migrated_to_spec
- [x] spec status mismatch
- [x] design_record missing spec silent skip
- [x] `kind: spec` + `id_range` request error
- [x] fenced code block headings exclusion
- [x] If `suggest_next_record` is implemented, slug generation covers ASCII and non-ASCII titles

---

## Done criteria for M16

M16 is done when:

- [x] `cmd/design-records-mcp/` can run independently
- [x] `internal/designrecords/` owns parser / index / validation / tool handlers
- [x] ADR/spec record index can be built from repo docs
- [x] `list_records` is implemented
- [x] `validate_records` is implemented
- [x] `get_record` is implemented
- [x] P1 `suggest_next_record` is either implemented or explicitly deferred in this task file
- [x] MVP diagnostic categories match `docs/spec/design-records-mcp/schema.md`
- [x] tool behavior matches `docs/spec/design-records-mcp/tools.md`
- [x] tests cover parser, index, validation, and P0 tools
- [x] implementation does not depend on existing brewprint YAML semantic build / `ResolvedProject`

---

## Follow-up candidates

These are explicitly outside M16 MVP and should be handled by later ADR/spec/task if needed.

- `trace_record`
- `list_gaps`
- write tools for ADR/spec updates
- metadata gap diagnostics such as `accepted_but_not_migrated`
- section-level traceability
- task / UC / impl note indexing
- existing brewprint MCP integration or shared launcher
- expose Design Records MCP handlers through MCP transport and add integration/smoke tests
- external generic Design Records CLI / OSS packaging

M16 close note: `cmd/design-records-mcp` currently runs independently as an index summary binary. MCP server registration / transport exposure is not part of this M16 close.

# TASK-MCP-003-03: Workflow artifact MCP support の public contract を spec に反映する

- **id**: TASK-MCP-003-03
- **status**: done
- **date**: 2026-05-26
- **work_item**: WORK-MCP-003
- **source_requirement**: REQ-MCP-003
- **estimate**: 1d
- **depends_on**:
  - TASK-MCP-003-02
- **outputs**:
  - 採用された workflow artifact MCP support の spec 更新
  - implementation が追従すべき request / response / validation boundary

## Goal

TASK-MCP-003-02 で採用された最小 public contract を、Design Records MCP の現行 spec に整合する形で明文化し、implementation が未確定事項へ依存せず着手できる状態にする。

## Work

- 採用判断に応じて `SPEC-design-records-mcp-overview` / `schema` / `tools` の必要箇所を更新する。
- Record/query surface、resolver input、relation validation、diagnostic category の採用範囲を明文化する。
- MVP 外と判断された orphan diagnostics / progress projection 等を非対象として明記する。
- ID-as-ref と physical path 非対応の境界を contract に維持する。

## Done condition

- 採用された public surface が spec 上で矛盾なく定義されている。
- Implementation / tests が必要とする schema と diagnostic behavior が未確定のまま残っていない。
- TASK-MCP-003-04 が開始可能である。

## Verification

- TASK-MCP-003-02 の判断記録と spec 差分を照合する。
- 既存 `get_record(s)` / `resolve_reference` / `validate_records` の責務分離が崩れていないことを確認する。

## Evidence

- 2026-05-27: `ADR-092` accepted および `TASK-MCP-003-02` done を受け、spec 反映に着手した。
- 更新対象は `docs/spec/design-records-mcp/{overview,schema,tools}.md` と、採用済み canonical reference / tool boundary を保持する `docs/spec/concepts/{project-artifact-model,traceability}/**` の関連記述とする。
- Implementation は未追従であるため、runtime availability の確認と implementation evidence は後続 `TASK-MCP-003-04` / `05` で扱う。
- 2026-05-27: `SPEC-design-records-mcp-overview` / `schema` / `tools` に workflow record kind、resolver input、declared relation integrity、investigation metadata の `REQ-*` / `WORK-*` 限定 boundary、MVP 外 scope を反映した。
- 2026-05-27: `spec:project-artifact-model` および `spec:trace` 配下の artifact refs / metadata schema / resolve-and-validation / out-of-scope に、ADR-092 accepted 後の canonical reference / validation boundary を同期した。
- 2026-05-27: 読み戻しで、reserved `yaml:` に限定すべき未定義 boundary の曖昧表現と、design record の `depends_on` と `task.depends_on` の field-name collision を検出し、tools / schema で責務分離を明記して解消した。
- 2026-05-27: Spec 更新により、`TASK-MCP-003-04` は workflow document discovery、record serialization、resolver、declared relation validation、対応 tests の implementation に着手可能となった。
- 2026-05-27: 最終整合確認で、overview の metadata source 説明、schema summary、tools の `invalid_workflow_id` diagnostic 一覧を workflow artifact support と一致させた。
- 2026-05-27: Spec reflection review は `OK with non-blocking follow-ups`。Implementation 着手前に、空の `task.depends_on:` を `[]` に正規化して diagnostic を出さない parser contract を `schema.md` に追記し、resolver response の concrete field / status vocabulary は `tools.md` が所有するよう `spec:trace.resolve-and-validation` の重複定義を解消した。Entry docs の stale wording cleanup は後続 follow-up として保持する。

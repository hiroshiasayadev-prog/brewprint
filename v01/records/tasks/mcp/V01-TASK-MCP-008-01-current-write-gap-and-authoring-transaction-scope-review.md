# V01-TASK-MCP-008-01: Current write gap and authoring transaction scope review

- **id**: V01-TASK-MCP-008-01
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-MCP-008
- **source_requirement**: V01-REQ-MCP-008
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Current Design Records MCP write capability gap summary
  - Authoring transaction MVP scope confirmation
  - ADR drafting inputs for V01-TASK-MCP-008-02

## Goal

現行 Design Records MCP の read / navigation / validation / guidance capabilities と、V01-REQ-MCP-008 が求める authoring transaction capability の差分を確認し、ADR 起草前に MVP scope と設計判断点を明確にする。

## Work

- `V01-REQ-MCP-008` と `V01-WORK-MCP-008` を確認し、要求された authoring transaction scope を整理する。
- `SPEC-design-records-mcp-tools` の現行 tools contract を確認し、write / proposal / cache / accept に相当する public surface が存在しないことを確認する。
- 現行 Design Records MCP の tool set が read / navigation / validation / guidance に偏っているかを確認する。
- Existing related requirements / work items such as `V01-REQ-MCP-005`, `V01-REQ-MCP-007`, `V01-WORK-MCP-005`, and `V01-WORK-MCP-007` を必要範囲で確認し、artifact-oriented access の既存方針と矛盾しないか整理する。
- ADR 化すべき判断点を分類する。
  - propose → accept transaction model
  - pathless artifact-oriented write input
  - `new` placeholder ID resolution
  - `body` / `body_cache_id` exclusive source contract
  - 3 days proposal / body cache retention
  - set-only MVP for front matter / section updates
  - validation failure vs write failure response semantics

## Done condition

- Current Design Records MCP write capability gap が、spec / current tools / related requirements の evidence として整理されている。
- V01-REQ-MCP-008 の MVP scope が、ADR / spec / implementation に渡せる粒度で確認されている。
- ADR draft に含めるべき decision points と、ADR ではなく spec に委ねるべき schema details が分離されている。
- 後続 `V01-TASK-MCP-008-02` が ADR draft に進める状態になっている。

## Verification

- Relevant requirement / work item / spec / current MCP tool contract の確認結果を Evidence に記録する。
- 可能であれば `validate_records` で `V01-REQ-MCP-008`, `V01-WORK-MCP-008`, and this task の metadata relation を確認する。
- Repo-local command execution が必要な検証は Codex / implementation agent に委譲するか、未実施として明示する。

## Evidence

### Files inspected

- `AGENTS.md`
- `docs/prompt_chappy.md`
- `docs/doc-policy.md`
- `docs/requirements/mcp/REQ-MCP-008-design-records-authoring-transaction-support.md`
- `docs/work-items/mcp/WORK-MCP-008-design-records-authoring-transaction-support.md`
- `docs/tasks/mcp/TASK-MCP-008-01-current-write-gap-and-authoring-transaction-scope-review.md`
- `docs/spec/design-records-mcp/tools.md`
- `docs/spec/design-records-mcp/schema.md`
- `docs/requirements/mcp/REQ-MCP-005-project-authoring-guidance-retrieval-support.md` via Design Records MCP
- `docs/requirements/mcp/REQ-MCP-007-list-records-workflow-artifact-range-filter-support.md` via Design Records MCP
- `docs/work-items/mcp/WORK-MCP-005-project-authoring-guidance-retrieval-support.md` via Design Records MCP
- `docs/work-items/mcp/WORK-MCP-007-list-records-workflow-artifact-range-filter-support.md` via Design Records MCP
- Design Records MCP authoring guides: `requirement-authoring`, `work-item-authoring`, `task-authoring`, `artifact-boundary`
- Accepted ADR title list under `docs/adr/`, especially the current Design Records MCP lineage referenced by the specs: V01-ADR-076, V01-ADR-077, V01-ADR-087, V01-ADR-088, V01-ADR-090, V01-ADR-092
- `internal/designrecords`
- `internal/designrecordsmcp`
- Targeted tests under `internal/designrecords/*test.go` and `internal/designrecordsmcp/*test.go`

### Current public tool surface

Current public Design Records MCP tools are:

- `list_records`
- `get_record`
- `get_records`
- `validate_records`
- `resolve_reference`
- `list_authoring_guides`
- `get_authoring_guidance`
- `suggest_next_record`

Evidence:

- `docs/spec/design-records-mcp/tools.md` lists seven P0 tools and one P1 helper. The P0 tools are `list_records`, `get_record`, `get_records`, `validate_records`, `resolve_reference`, `list_authoring_guides`, and `get_authoring_guidance`. The only P1 helper is `suggest_next_record`.
- `internal/designrecordsmcp/tools.go` registers the same eight tool names in `Tools()`.
- `internal/designrecordsmcp/tools_call.go` dispatches only those eight tool names and returns an unknown-tool error for anything else.
- `internal/designrecordsmcp/jsonrpc_test.go` asserts those same eight names appear in `tools/list`.

### Current spec contract summary

The current contract is intentionally read-only.

- `docs/spec/design-records-mcp/tools.md` states that the MVP tool interface is read-only and that file creation, file update, Evidence rewrite, and commit operations are not performed.
- The same spec has a dedicated `Write tool policy` section saying that MVP does not provide write tools.
- Explicitly out-of-scope examples include `create_record`, `update_record`, `set_evidence`, `add_record_metadata`, and `migrate_record_to_spec`.
- The spec says any future write tools need separate ADR / spec definition for dry-run diff, user confirmation, conflict handling, template responsibility, and git-operation boundaries.
- `suggest_next_record` is only a read-only ADR ID/path suggestion helper. The spec and implementation both state that it does not create a file.
- `docs/spec/design-records-mcp/schema.md` defines sources read by the MCP: ADR/investigation/workflow bullet metadata, spec front matter, Markdown H1/headings/body, file path for discovery/validation, and authoring guide Markdown. It does not define an authoring proposal, write transaction, mutable front matter update, mutable section update, or body cache schema.

### Current implementation / tests summary

Implementation matches the read-only contract.

- `internal/designrecordsmcp/tools.go` exposes input schemas for read/navigation/validation/guidance tools only. There are no registered input schemas for creation, update, proposal, accept, discard, body cache, front matter replacement, or section replacement.
- `internal/designrecordsmcp/tools_call.go` dispatches only the current eight read/navigation/guidance/helper tools.
- `internal/designrecords/tools.go` implements `ListRecords`, `GetRecord`, `GetRecords`, `ValidateRecords`, and `SuggestNextRecord`. `SuggestNextRecord` computes the next decision ID/path from the index and does not write files.
- `internal/designrecords/authoring_guidance.go` reads guide catalog/content from `docs/guides/*.md`; it does not author records.
- Existing tests cover list/get/get_records/validate/resolve/guidance/suggest behavior, including workflow artifact read/range navigation and a `TestSuggestNextRecordNoSideEffects` test that asserts no suggested file is created.
- The test inventory does not include proposal, accept, discard, body cache, write recovery, front matter update, named section update, or create/update transaction tests.

### Confirmed write capability gap

No current public tool supports any of the following V01-REQ-MCP-008 authoring-write capabilities:

- record creation for requirement, work item, task, decision, or spec skeleton;
- section update;
- front matter update;
- proposal creation;
- accept/write;
- discard proposal;
- body cache / `body_cache_id` retry;
- write recovery reporting;
- explicit `written: true/false` write outcome semantics;
- validation/repair hints tied to a write or proposed write.

This is not a bug in the current Design Records MCP. It is a public contract / authoring capability gap: the existing accepted/spec-backed MVP deliberately excluded write tools.

### Confirmed artifact-oriented read/navigation vs path-oriented authoring gap

Current read/navigation/guidance is artifact-oriented:

- `list_records` filters by kind/status/id/id_range and supports workflow artifact ranges for `REQ-*`, `WORK-*`, and `TASK-*`.
- `get_record` and `get_records` use exact record IDs and can return metadata/headings/body plus transparent relative `path`.
- `resolve_reference` resolves canonical `spec:` refs and record ID-as-ref targets to document/section/record targets.
- `validate_records` validates indexed records and workflow artifact relations by record kind/range.
- `list_authoring_guides` and `get_authoring_guidance` use guide IDs rather than guide file paths.

Current authoring writes are still path-oriented because no Design Records MCP write tool exists. Creating or updating `REQ-*`, `WORK-*`, `TASK-*`, ADR, or spec skeleton files still requires direct filesystem editing of physical Markdown files, even when the user intent is expressed in artifact IDs or workflow terms.

Related evidence:

- V01-REQ-MCP-005 / V01-WORK-MCP-005 established artifact-independent authoring guidance retrieval by guide ID and explicitly kept guide source paths out of the public contract.
- V01-REQ-MCP-007 / V01-WORK-MCP-007 extended artifact-oriented range navigation for workflow artifacts.
- Those completed scopes support read/navigation/guidance ergonomics, but they did not introduce record authoring writes.

### V01-REQ-MCP-008 MVP scope confirmation

The V01-REQ-MCP-008 MVP scope is still valid and not already implemented.

Already supported by current MCP:

- artifact-oriented record discovery/retrieval/validation/reference resolution for decision/spec/investigation/requirement/work_item/task;
- artifact-oriented authoring guidance discovery/retrieval by guide ID;
- workflow artifact range navigation for requirement/work item/task;
- P1 read-only ADR next-ID/path suggestion through `suggest_next_record`.

Needs new public tool contract, spec updates, implementation, and tests:

- propose record create;
- propose record update;
- get proposed write;
- accept proposed write;
- discard proposed write;
- create requirement/work item/task/decision records and spec skeletons;
- `new` placeholder ID resolution for create operations only;
- front matter whole replacement;
- named Markdown section whole replacement;
- previewable diff responses before writing;
- `body` / `body_cache_id` exclusive body source validation;
- 3 day proposal/body cache retention;
- body cache retry guidance when large-body proposal/write setup fails before persistence;
- validation failure vs write failure distinction;
- `written: true/false` response semantics;
- actionable validation diagnostics and repair hints for failed writes/proposals.

### ADR decision points for `V01-TASK-MCP-008-02`

The next ADR should capture durable design decisions, not tool-schema minutiae:

- whether Design Records MCP should add authoring write capability at all, despite the current read-only MVP boundary;
- propose -> accept as the transaction model, including the rule that proposal creation must not write repository files;
- pathless / artifact-oriented write inputs as the primary public contract, with physical paths returned only for transparency;
- supported MVP create/update surface: requirement, work item, task, decision, spec skeleton create; front matter whole replacement; named Markdown section whole replacement;
- `new` placeholder ID resolution on the MCP side for create operations, and rejection for update operations;
- single body source rule: exactly one of `body` or `body_cache_id` when a body source is required;
- 3 day retention for proposal and body caches;
- body cache recovery as the retry mechanism for large submitted bodies that fail before proposal persistence;
- response distinction between validation failure and write failure, including explicit `written` semantics;
- exclusion of generic filesystem write, path-first authoring APIs, partial Markdown AST editing, add/remove relation convenience operations, multi-record transactions, automatic close cascades, formatter integration, indefinite retention, and force-accepting invalid proposals from the MVP.

### Spec-owned details not to over-specify in the ADR

The ADR should leave these to the tools/schema spec and implementation tasks:

- final tool names;
- exact request/response JSON field names;
- proposal ID format;
- proposal/body cache storage layout;
- diff format and normalization details;
- full JSON schemas for create/update payloads;
- exact template text for requirement/work item/task/decision/spec skeleton creation;
- exact front matter serialization rules;
- exact named-section matching algorithm and ambiguity diagnostics;
- exact validation diagnostic categories and repair-hint wording;
- cache cleanup timing mechanics beyond the 3 day retention decision;
- implementation package structure and test fixture layout.

### Test / command results

Targeted checks performed:

- `rg --files internal/designrecords internal/designrecordsmcp`
  - Confirmed the relevant implementation and test files are under the two requested packages.
- Targeted `rg` for current tool names and write-related terms (`write`, `create`, `update`, `proposal`, `accept`, `cache`, `body_cache`, `front_matter`, `section`) across `docs/spec/design-records-mcp/tools.md`, `docs/spec/design-records-mcp/schema.md`, `internal/designrecords`, and `internal/designrecordsmcp`.
  - Confirmed write-related terms only appear as read-only policy, out-of-scope examples, test fixture helpers, raw body retrieval, semantic section resolution/validation, or suggested path support. No public write/proposal/cache transaction implementation was found.
- Design Records MCP validation:
  - `validate_records(kind: requirement, id_range: V01-REQ-MCP-008..V01-REQ-MCP-008)` returned `ok: true` with no diagnostics.
  - `validate_records(kind: work_item, id_range: V01-WORK-MCP-008..V01-WORK-MCP-008)` returned `ok: true` with no diagnostics.
  - `validate_records(kind: task, id_range: V01-TASK-MCP-008-01..V01-TASK-MCP-008-01)` returned `ok: true` with no diagnostics.
- `go test ./internal/designrecords ./internal/designrecordsmcp`
  - Passed.
  - Output summary:
    - `ok   github.com/hiroshiasayadev-prog/brewprint/internal/designrecords 1.450s`
    - `ok   github.com/hiroshiasayadev-prog/brewprint/internal/designrecordsmcp (cached)`

### Conclusion

`V01-TASK-MCP-008-01` satisfies its done condition.

The current Design Records MCP is internally consistent with its accepted/spec-backed read-only MVP contract. The gap confirmed here is not an implementation bug; it is the absence of a public artifact-oriented authoring transaction contract. `V01-WORK-MCP-008` should proceed to `V01-TASK-MCP-008-02` to draft an ADR for the transaction model before changing the tools spec or implementation.

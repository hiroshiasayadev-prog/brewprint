# V01-TASK-MCP-008-05: MCP tools spec reflection review

- **id**: V01-TASK-MCP-008-05
- **status**: done
- **date**: 2026-06-02
- **work_item**: V01-WORK-MCP-008
- **source_requirement**: V01-REQ-MCP-008
- **estimate**: 0.5d-1d
- **depends_on**:
  - V01-TASK-MCP-008-04
- **outputs**:
  - Spec reflection review result
  - Spec reflection fixes before implementation
  - Follow-up requirement for spec domain tree / placement discovery

## Goal

`V01-TASK-MCP-008-04` で反映された Design Records MCP authoring transaction spec を independent reviewer / Opus にレビューさせ、implementation に進む前の spec ambiguity を解消する。

## Work

- `V01-ADR-093`, `SPEC-design-records-mcp-tools`, `SPEC-design-records-mcp-schema`, `V01-REQ-MCP-008`, `V01-WORK-MCP-008`, and `V01-TASK-MCP-008-04` を照合する。
- Review findings を分類し、implementation 前に必要な spec / ADR / requirement 修正を反映する。
- Spec skeleton creation / `SPEC-new` の扱いを見直し、MVP から除外する。
- Spec placement discovery / domain tree support を follow-up requirement として切り出す。
- Blocking / major findings が残っていない状態にする。

## Done condition

- Spec reflection review result が Evidence に記録されている。
- Implementation 前に必要な findings が反映済みである。
- `SPEC-new` / spec skeleton create の MVP boundary が明確になっている。
- Spec placement discovery / domain tree support が follow-up requirement として捕捉されている。
- `V01-TASK-MCP-008-06` の implementation and tests に進める状態になっている。

## Verification

- `V01-ADR-093`, `V01-REQ-MCP-008`, `V01-WORK-MCP-008`, `V01-TASK-MCP-008-05`, and related specs の validation を確認する。
- Repo-local command execution が必要な validation / rendering / tests は Codex / implementation agent に委譲するか、未実施として明示する。

## Evidence

2026-06-02: Spec reflection review completed.

Review result:

- Verdict: Needs revision before implementation
- Structural conclusion: The main authoring transaction tool surface and proposal lifecycle were aligned with V01-ADR-093, but several implementation-facing ambiguities had to be resolved before implementation.

Findings addressed:

- Spec metadata block replacement now uses scoped recognized-field replacement for `spec` records. Unknown / auxiliary YAML front matter fields must be preserved.
- Spec skeleton creation and `SPEC-new` placeholder create are excluded from the MVP because safe spec placement cannot be derived from ID alone.
- `V01-ADR-093` `migrated_to_spec` was set to `2026-06-02`.
- Section selector resolution now explicitly uses the same heading source rules as record headings: YAML front matter and fenced code block content are excluded, and setext headings are outside MVP.
- `new` placeholder grammar is now explicit: `new` is a literal token in the sequence position, not a substring match.
- Body cache entries remain reusable within the 3 day retention period, including after successful proposal creation.
- `propose_record_update` now includes concrete response shape guidance.
- `get_records` spec example `depends_on` was updated for `SPEC-design-records-mcp-tools`.

Follow-up requirement created:

- `V01-REQ-MCP-010`: Spec domain tree and placement discovery support

Boundary decision:

- Existing spec records remain update targets for metadata block replacement and named section replacement.
- New spec skeleton creation is excluded from V01-REQ-MCP-008 MVP and deferred to a later scope after spec placement discovery is available.

Implementation handoff:

- No implementation files were edited as part of this task.
- `V01-TASK-MCP-008-06` may proceed after validation confirms the updated records and specs remain indexed.

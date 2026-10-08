# DRCLI-TASK-CLI-002-06: Author DRCLI guidance operation specs

- **status**: done
- **date**: 2026-10-07
- **work_item**: DRCLI-WORK-CLI-002
- **task_type**: authoring
- **estimate**: 0.5d-1d
- **depends_on**:
  - DRCLI-TASK-CLI-002-01
- **outputs**:
  - spec:drcli.design_records_cli.operations.guidance
  - spec:drcli.design_records_cli.operations.guidance.list
  - spec:drcli.design_records_cli.operations.guidance.exact_read

## Goal

Create DRCLI authoring-guidance list and exact-read contracts.

## Work

Write only:

- `drcli/records/spec/design-records-cli/operations/guidance/**`

Use as baseline:

- old DRMCP `list-authoring-guides.md`;
- old DRMCP `get-authoring-guidance.md`;
- accepted DRMCP-REQ-MCP-003 package/guidance semantics;
- current PRODUCT portable standards authority.

Preserve the fixed `design_records` package scope, exact canonical refs, list projection, complete verbatim guide content, no fuzzy lookup, and no separate legacy guide model.

Do not define concrete command spelling; T08 owns CLI mapping.

## Done condition

- Guidance list and detail contracts are complete.
- Guide content comes from the portable standards package treated as normal current Specs.
- Full Markdown detail is verbatim and untruncated.
- No legacy guide source becomes canonical.

## Verification

- Trace DRCLI-REQ-CLI-001 guidance outcomes to the new contracts.
- Confirm package refs and fixed scope are current.
- Confirm only the assigned path family changed.

## Evidence

- Authored only `drcli/records/spec/design-records-cli/operations/guidance/**` plus this Task's lifecycle, outputs, and Evidence.
- Created `spec:drcli.design_records_cli.operations.guidance` as the fixed portable-guidance boundary.
- Created `spec:drcli.design_records_cli.operations.guidance.list` with fixed `design_records` / `spec` / `spec:design_records.authoring_standards.*` scope, root exclusion, canonical-ID/title/abstract projection, ASCII lexical ordering, and no partial catalog fallback.
- Created `spec:drcli.design_records_cli.operations.guidance.exact_read` with exact canonical-ref lookup, no normalization or fuzzy/alias/path fallback, and complete verbatim untruncated Markdown content.
- Preserved PRODUCT authority: guidance reads the portable package as normal current Specs and does not create a legacy guide source or independent semantic model.
- Preserved DRCLI-REQ-CLI-001 guidance outcomes without defining concrete command or option spelling; CLI mapping remains outside this Task.
- Used filesystem authoring because the current agent-authoring policy marks DRMCP authoring transactions non-operational.
- Verification was scoped to the three guidance Specs and this Task. No DRMCP file, sibling lane output, root DRCLI Spec index, or parent Work Item lifecycle was modified.
- The current PRODUCT source and repository-local derived package both confirm the fixed `spec:design_records.authoring_standards` package ref shape. The derived `bin/design-records/` snapshot currently lacks source `usdm-authoring.md`; PRODUCT remains semantic authority, so this Task did not regenerate or treat that snapshot as independent authority.
- No independent review was performed by this authoring Task.

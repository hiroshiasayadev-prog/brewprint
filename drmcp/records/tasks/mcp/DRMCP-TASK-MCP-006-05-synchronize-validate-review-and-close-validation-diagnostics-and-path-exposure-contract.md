# DRMCP-TASK-MCP-006-05: Synchronize, validate, review, and close validation, diagnostics, and path-exposure contracts

- **id**: DRMCP-TASK-MCP-006-05
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-006
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1.5d
- **depends_on**:
  - DRMCP-TASK-MCP-006-04
- **outputs**:
  - spec:drmcp.design_records_mcp.overview
  - spec:drmcp.design_records_mcp.tools.overview
  - spec:drmcp.design_records_mcp.mvp_scope
  - spec:drmcp.design_records_mcp.responsibility_boundary
  - DRMCP-WORK-MCP-006
  - DRMCP-TASK-MCP-001-06

## Goal

Synchronize the final cross-spec pointers for the accepted W006 contract.

Validate the complete normative change set, complete independent review, correct findings, and close W006 without reopening T01 through T04 decisions.

## Work

- Confirm the exact `DRMCP-TASK-MCP-006-*` inventory before creating this Task.
- Consume the accepted T01 authority baseline and T02 through T04 contracts.
- Confirm that T04 is `done` and its final review has no remaining finding.
- Aggregate the complete T01 through T04 changed-file and recheck-only manifests.
- Recheck W003 through W005 non-regression boundaries.
- Confirm that W007 disposition and PRODUCT validation-policy owner-pointer work remain outside W006.
- Inspect `overview.md`, `tools/overview.md`, `mvp-scope.md`, and `responsibility-boundary.md` one file at a time.
- Record `change` or `recheck-only / no change` for every synchronization candidate.
- Modify only files with a direct contradiction or missing authoritative pointer.
- Keep overview contracts navigation-first and avoid copying detailed operation or schema rules.
- Confirm the final normative changed-file manifest.
- Run scoped strict validation against every changed normative specification.
- Run tracked and untracked whitespace checks with separate commands.
- Record an independent final review prompt before review begins.
- Correct every accepted finding and rerun affected verification.
- After final review returns `PASS`, synchronize T05, W006, and hub Task closure.
- Treat post-closure whitespace output as external evidence and do not write it back into checked files.

This Task does not define new validation semantics, diagnostic categories, resolver statuses, read behavior, authoring transaction behavior, location shapes, or path grammar.

## Done condition

- Every synchronization candidate has an explicit disposition.
- The final normative changed-file manifest is complete, and the T05 direct-change manifest matches Task metadata outputs.
- Overview and responsibility contracts contain no direct contradiction with T02 through T04.
- W003 through W005 accepted boundaries remain unchanged.
- W007 remains the owner of DRMCP validation Work Item disposition.
- PRODUCT validation-policy owner-pointer synchronization remains outside W006.
- Fixtures and implementation remain outside W006.
- Every changed normative specification passes scoped strict validation.
- Tracked and untracked changed files pass applicable whitespace checks.
- Independent final review returns `PASS` with no remaining blocking, major, or minor finding.
- T04 remains `done`.
- This Task, W006, and `DRMCP-TASK-MCP-001-06` are `done`.
- Repository-wide clean status is not inferred.

## Verification

- Compare final summaries with `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`.
- Compare validation execution with T02, diagnostic representation with T03, and location and path exposure with T04.
- Compare normal read behavior with W004 and resolver behavior with W005.
- Confirm that W007 and PRODUCT owner-pointer records were not absorbed into W006.
- Run the strict spec validator against only the final changed normative specification set.
- Run `git diff --check` against tracked T05 files.
- Run `git diff --no-index --check -- NUL` against this Task while it remains untracked.
- Run independent final review before closure synchronization.
- Run one final external whitespace check after closure edits and do not write that result back into checked files.

## Evidence

### Exact Task inventory

The exact directory `drmcp/records/tasks/mcp/` was listed on 2026-06-28.

Existing W006 child Tasks were T01 through T04.
No `DRMCP-TASK-MCP-006-05` record existed.
The next Task is therefore `DRMCP-TASK-MCP-006-05`.

The canonical path is:

`drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-05-synchronize-validate-review-and-close-validation-diagnostics-and-path-exposure-contract.md`.

### Accepted baseline

| phase | accepted input | T05 treatment |
|---|---|---|
| T01 | Authority matrix, W003-W005 inputs, contradiction inventory, manifests, W007 boundary, and Task split. | Consume without reopening. |
| T02 | Request scope, subject selection, relation lookup, failure boundaries, wrapper, and `ok` semantics. | Synchronize pointers only. |
| T03 | Envelope, categories, severities, associations, conflicts, ordering, suppression, authoring compatibility, and PRODUCT pointers. | Synchronize pointers only. |
| T04 | Location object, portable path grammar, containment, identity, sorting, exposure, fail-closed behavior, authoring paths, and privileged-operation boundary. | Synchronize pointers only. |

T04 is `done`.
Its final limited re-review returned `PASS`.
`F-MIN-01`, `F-MIN-02`, and `F-MIN-03` are closed.
No blocking, major, minor, or advisory finding remains.

### Authority and exclusion boundary

- PRODUCT owns semantic validity.
- DRMCP owns validation execution, diagnostic representation, operation contracts, and MCP path exposure.
- W003 owns current discovery, retained sources, record state, conflicts, provenance, and active-index construction.
- W004 owns list and exact-retrieval requests, ordering, partial success, warnings, wrappers, successful projection, and normal path hiding.
- W005 owns resolver grammar order, fallback eligibility, legacy lookup, statuses, and successful target projection.
- W007 owns DRMCP validation Work Item disposition and rebaseline.
- `PRODUCT-WORK-SPEC-015` owns PRODUCT validation-policy owner-pointer synchronization.
- W006 does not own fixtures or runtime implementation.

### Aggregated T01 through T04 normative manifest

T01 changed no normative specification.
The accepted T02 through T04 union is:

- `drmcp/records/spec/design-records-mcp/schema/diagnostics.md`;
- `drmcp/records/spec/design-records-mcp/tools/validate-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/list-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/get-records.md`;
- `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`;
- `drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md`;
- `drmcp/records/spec/design-records-mcp/tools/authoring-transaction-model.md`;
- `drmcp/records/spec/design-records-mcp/tools/propose-record-create.md`;
- `drmcp/records/spec/design-records-mcp/tools/propose-record-update.md`;
- `drmcp/records/spec/design-records-mcp/tools/accept-proposed-write.md`.

This is the pre-T05 normative manifest.
T05 adds a candidate only when a direct contradiction or missing pointer requires an edit.

### Final synchronization candidate disposition

| file | disposition | reason and bounded change |
|---|---|---|
| `overview.md` | `change` | Replaced the ambiguous physical-path ownership summary with pointers to W004, W005, the diagnostic schema, the no-current-absolute-path boundary, and the separate future privileged-operation requirement. |
| `tools/overview.md` | `change` | Replaced a Work Item-level diagnostic authority pointer with the canonical diagnostic schema and preserved operation ownership of triggers and response placement. |
| `mvp-scope.md` | `change` | Separated W004 warning triggers, shared diagnostic representation, validation execution, and PRODUCT semantic invalidity. Added resolver to the successful path-free projection summary and preserved the absolute-path prohibition. |
| `responsibility-boundary.md` | `change` | Replaced W006 ownership of `validation semantics` with validation execution and diagnostic mapping while retaining PRODUCT semantic authority. Replaced the broad physical-path statement with portable repository-relative exception rules and the future privileged absolute-path boundary. |

No detailed operation, diagnostic, resolver, authoring, location, or path grammar was copied into the overview contracts.

### Recheck-only disposition

| file | disposition | reason |
|---|---|---|
| `schema/discovery.md` | `recheck-only / no change` | Current candidate, invalid-source retention, PRODUCT path authority, and W004/W006 handoffs remain consistent. |
| `schema/record-model.md` | `recheck-only / no change` | Repository-relative operational path retention and public read versus repair-diagnostic ownership remain consistent. |
| `schema/record-source.md` | `recheck-only / no change` | Current metadata and path-derived identity sources remain unchanged; no T05 path-exposure contract is defined here. |
| `namespace-scanning.md` | `recheck-only / no change` | Configured-root, containment, current/legacy separation, and diagnostic handoffs remain consistent with W003, W005, and W006. |
| `resolver.md` | `recheck-only / no change` | Current-first orchestration and successful non-path target projection remain W005-owned; diagnostic and validation concerns remain delegated. |
| `tools/propose-record-update.md` | `recheck-only / no change` | Existing authoring target, diff, patch, and diagnostic location fields use the accepted repository-relative T04 representation. |
| `tools/get-proposed-write.md` | `recheck-only / no change` | No direct contradiction or missing T05 authority pointer was found. |

### Final normative changed-file manifest

The complete W006 normative contract manifest is the accepted T02 through T04 ten-file union plus these four T05 synchronization changes:

- `drmcp/records/spec/design-records-mcp/overview.md`;
- `drmcp/records/spec/design-records-mcp/tools/overview.md`;
- `drmcp/records/spec/design-records-mcp/mvp-scope.md`;
- `drmcp/records/spec/design-records-mcp/responsibility-boundary.md`.

The complete final normative set therefore contains 14 files.
The T05 direct normative change set contains exactly the four files above and matches the Task metadata outputs.

### Repository-local verification commands

Strict validation covers only the four normative specifications changed directly by T05:

```powershell
python -X utf8 product/src/tools/validate_spec.py `
  drmcp/records/spec/design-records-mcp/overview.md `
  drmcp/records/spec/design-records-mcp/tools/overview.md `
  drmcp/records/spec/design-records-mcp/mvp-scope.md `
  drmcp/records/spec/design-records-mcp/responsibility-boundary.md `
  --strict --no-color
```

Tracked whitespace covers every tracked file changed directly by T05:

```powershell
git diff --check -- `
  drmcp/records/spec/design-records-mcp/overview.md `
  drmcp/records/spec/design-records-mcp/tools/overview.md `
  drmcp/records/spec/design-records-mcp/mvp-scope.md `
  drmcp/records/spec/design-records-mcp/responsibility-boundary.md `
  drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md

$tracked_exit = $LASTEXITCODE
```

The new Task is checked separately while it remains untracked:

```powershell
git diff --no-index --check -- NUL `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-05-synchronize-validate-review-and-close-validation-diagnostics-and-path-exposure-contract.md

$untracked_exit = $LASTEXITCODE

"tracked_exit=$tracked_exit"
"untracked_exit=$untracked_exit"
```

For the untracked Task, exit code `1` is expected when no whitespace error is printed. Exit code `2` or higher is a failure.

### Current verification state

- Required instruction and authoring standards: read.
- Exact W006 Task inventory: complete.
- T05 Task creation: complete.
- T05 linkage to W006: complete.
- Synchronization candidate inspection: complete.
- Recheck-only inspection: complete.
- Final changed-file manifest: complete; 14 normative files total, four changed directly by T05.
- Filesystem-visible structure check: the four changed specs retain one H1, required H1-adjacent metadata, and canonical top-level sections; the T05 Task retains canonical Task metadata and H2 sections.
- Repository-local strict validator: `[strict]  All 4 file(s) OK.`
- Tracked whitespace: `tracked_exit=0`; no whitespace error was reported.
- Initial untracked Task whitespace: `untracked_exit=3` because line 354 had a new blank line at EOF.
- The extra EOF blank line was removed after that check.
- Corrected untracked Task whitespace recheck: `untracked_exit=1`; no whitespace error was reported.
- That result was then written into this Task and W006, so the checked bytes changed.
- LF-to-CRLF working-copy conversion warnings are non-blocking.
- Final pre-review tracked and untracked whitespace recheck after evidence synchronization: `tracked_exit=0`, `untracked_exit=1`; no whitespace error was reported.
- The final pre-review recheck result remained external and was not written back before review.
- Independent final review: `PASS`.
- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- Advisories: none.
- Repository-wide clean status: not inferred.

### Independent final review result

- Verdict: `PASS`.
- T04 `F-MIN-01`, `F-MIN-02`, and `F-MIN-03` remain closed with no regression.
- All four synchronization candidate dispositions were accepted.
- T01 through T04 and W003 through W005 non-regression assessments passed.
- W007 disposition and PRODUCT validation-policy owner-pointer ownership remain outside W006.
- The complete normative manifest is 14 files and the T05 direct normative change set is four files.
- Recorded strict validation and tracked/untracked whitespace evidence were accepted.
- Independent repository-local reruns were `NOT RUN` because the reviewer had no command-execution boundary.
- W006 completion conditions were assessed as satisfied in substance.
- T05, W006, and hub T06 closure readiness was `READY`.
- No normative contract correction was required.
- T05, W006, and `DRMCP-TASK-MCP-001-06` were synchronized to `done` on 2026-06-28.
- One post-closure tracked and untracked whitespace check remains external and must not be written back into checked files.

### Independent final review prompt

```text
C:\Users\imved\projects\brewprint

`DRMCP-WORK-MCP-006`と`DRMCP-TASK-MCP-006-05`のindependent final reviewを行う。

ファイルは変更しないこと。
DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
sandboxへrepositoryを複製しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。
repository-wide clean statusを推測しないこと。
無制限なrepository traversalや広範なsearchを行わず、以下のexact pathだけを必要な範囲で読むこと。

## 最初に読む

- `prompt_chappy.md`
- `product/records/spec/design-records/authoring-standards/task-authoring.md`
- `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
- `product/records/spec/design-records/authoring-standards/spec-authoring.md`
- `product/records/spec/design-records/authoring-standards/writing-standard.md`
- `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## Authority and workflow records

- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-design-records-mcp-contract-baseline-and-realignment.md`
- `drmcp/records/investigations/mcp/DRMCP-INV-MCP-002-design-records-mcp-contract-consistency-and-realignment-audit.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-06-track-validation-diagnostics-and-path-exposure-contract.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-01-establish-validation-diagnostics-and-path-exposure-correction-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-02-define-current-repository-and-relation-validation-execution-contract.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-03-define-machine-readable-diagnostic-representation-and-semantic-invalidity-mapping.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-04-define-source-location-and-exceptional-path-exposure-contract.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-006-05-synchronize-validate-review-and-close-validation-diagnostics-and-path-exposure-contract.md`

## Upstream non-regression owners

- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-003-current-discovery-and-active-index-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-005-resolver-and-configured-legacy-fallback-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md`

## Complete normative manifest

Read all 14 files:

- `drmcp/records/spec/design-records-mcp/schema/diagnostics.md`
- `drmcp/records/spec/design-records-mcp/tools/validate-records.md`
- `drmcp/records/spec/design-records-mcp/tools/list-records.md`
- `drmcp/records/spec/design-records-mcp/tools/get-records.md`
- `drmcp/records/spec/design-records-mcp/tools/resolve-reference.md`
- `drmcp/records/spec/design-records-mcp/schema/authoring-transaction-schema.md`
- `drmcp/records/spec/design-records-mcp/tools/authoring-transaction-model.md`
- `drmcp/records/spec/design-records-mcp/tools/propose-record-create.md`
- `drmcp/records/spec/design-records-mcp/tools/propose-record-update.md`
- `drmcp/records/spec/design-records-mcp/tools/accept-proposed-write.md`
- `drmcp/records/spec/design-records-mcp/overview.md`
- `drmcp/records/spec/design-records-mcp/tools/overview.md`
- `drmcp/records/spec/design-records-mcp/mvp-scope.md`
- `drmcp/records/spec/design-records-mcp/responsibility-boundary.md`

T05 direct synchronization changes are the final four files.

Recheck-only dispositions may be verified at these exact paths when needed:

- `drmcp/records/spec/design-records-mcp/schema/discovery.md`
- `drmcp/records/spec/design-records-mcp/schema/record-model.md`
- `drmcp/records/spec/design-records-mcp/schema/record-source.md`
- `drmcp/records/spec/design-records-mcp/namespace-scanning.md`
- `drmcp/records/spec/design-records-mcp/resolver.md`
- `drmcp/records/spec/design-records-mcp/tools/get-proposed-write.md`

`tools/propose-record-update.md` is already in the complete normative manifest and also has a recheck-only T05 disposition.

## Accepted baseline — reopen禁止

直接矛盾がない限り、T01からT04のaccepted decisionsを再設計しないこと。
新しいvalidation semantics、diagnostic category、resolver status、read behavior、authoring behavior、location shape、path grammarを導入しないこと。
W003 discovery/index、W004 list/exact retrieval、W005 resolver/fallback、W007 disposition、PRODUCT validation-policy owner-pointer scopeを再オープンしないこと。

T04 accepted final state:

- Verdict: `PASS`
- `F-MIN-01`, `F-MIN-02`, `F-MIN-03`: CLOSED
- blocking, major, minor, advisory finding: none
- status: `done`

## Review focus

以下を確認すること。

- 4つのsynchronization candidateすべてに根拠付きdispositionがある。
- 変更はpointerまたは正確な境界要約に限定され、詳細契約をoverviewへ重複していない。
- PRODUCT semantic validity、DRMCP validation execution、diagnostic representationの所有が混同されていない。
- normal successful list、exact retrieval、resolver projectionはpath-freeである。
- current operationまたはshared diagnostic fieldがabsolute physical pathを公開していない。
- portable diagnostic `location`とauthoring target/diff/patch/`files_written` pathが混同されていない。
- future absolute `physical_path`はseparate host-enabled privileged operationだけに限定される。
- T02 request scope、subject selection、relation lookup、request/execution failure、wrapper、`ok` semanticsが維持されている。
- T03 envelope、category、severity、association、current/legacy conflict、ordering、duplicate suppression、authoring compatibility、PRODUCT pointerが維持されている。
- T04 location shape、portable path、containment、identity、sort、exposure、fail-closed、authoring path、absolute-path boundaryが維持されている。
- W003からW005のnon-regression boundaryが維持されている。
- W007 dispositionとPRODUCT owner-pointer synchronizationがW006へ混入していない。
- complete normative manifestは14 filesで、T05 direct-change manifestは4 filesである。
- T05 metadata outputsが4 direct normative changesとW006/hub synchronizationを正確に表す。
- Taskへ記録されたstrict validationとwhitespace evidenceが、実際のcommand outputおよび対象file setと一致する。
- repository-wide clean statusを推測していない。

repository-local commandを実行できる場合のみ、記録済み結果を独立再確認してよい。
実行できない場合は`NOT RUN`と明記し、記録済みevidenceと独立実行を分けること。

## Output format

1. Verdict: `PASS` / `NEEDS REVISION`
2. Previous-finding disposition, if applicable
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. Synchronization-candidate disposition assessment
8. T01-T04 non-regression assessment
9. W003-W005 non-regression assessment
10. W007 and PRODUCT owner-boundary assessment
11. Final changed-file manifest assessment
12. Validation and whitespace evidence assessment
13. T04 closure-state assessment
14. W006 completion-condition assessment
15. T05 and W006 closure readiness
```

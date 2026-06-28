# DRMCP-TASK-MCP-007-04: Record PRODUCT handoff and close validation disposition

- **id**: DRMCP-TASK-MCP-007-04
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-007
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-007-03
- **outputs**:
  - DRMCP-WORK-MCP-007

## Goal

Record the accepted DRMCP validation-owner handoff for `PRODUCT-WORK-SPEC-015`.

Run final scoped verification and independent review before closing `DRMCP-WORK-MCP-007`.

## Work

- Record the exact pre-creation `DRMCP-TASK-MCP-007-*` inventory.
- Confirm T01, T02, and T03 are complete and accepted.
- Record the exact retained implementation-owner target set.
- Compare the authoritative PRODUCT pointer specs with the retained owner boundary.
- Record every known pointer mismatch for `PRODUCT-WORK-SPEC-015` without editing PRODUCT records.
- Confirm W003 through W006 remain accepted fixed boundaries.
- Confirm W008 and W009 ownership remains unchanged.
- Confirm W-SPEC-001 and W-SPEC-002 remain `not_started` implementation Work Items.
- Record the T04 changed-file and recheck-only manifests.
- Record targeted status and whitespace commands.
- Record a ready-to-run independent T04 review prompt.
- Keep T04 and W007 `in_progress` until external verification, independent review, and finding correction complete.
- Close T04 and W007 only after the review reports no unresolved blocking, major, or minor finding.

This Task does not:

- edit `PRODUCT-WORK-SPEC-015`;
- edit PRODUCT validation-policy or follow-up-boundary specs;
- start hub `DRMCP-TASK-MCP-001-07`;
- start W-SPEC-001 or W-SPEC-002 implementation;
- create fixtures or tests;
- edit implementation source;
- reopen T01 through T03;
- reopen W003 through W006.

## Done condition

- The exact T04 pre-creation inventory is recorded.
- T01, T02, and T03 remain `done`.
- The accepted per-file owner is exactly `DRMCP-WORK-SPEC-001`.
- The accepted Topics and graph owner is exactly `DRMCP-WORK-SPEC-002`.
- The retained Work Items remain separate.
- W-SPEC-002 retains its dependency on the accepted W-SPEC-001 detector boundary.
- PRODUCT semantic rules and severity policy remain PRODUCT-owned.
- The authoritative PRODUCT pointer specs are assessed against the accepted target set.
- Remaining PRODUCT-side pointer inventory and synchronization are delegated to `PRODUCT-WORK-SPEC-015`.
- W008 remains the shared fixture owner.
- W009 remains the general current-read implementation and test owner.
- No PRODUCT record, hub Task, implementation source, fixture, or test file changes.
- The changed-file and recheck-only manifests match the actual writes.
- Targeted whitespace evidence is available.
- Independent review reports no unresolved blocking, major, or minor finding.
- T04 changes to `done` only after verification and review.
- W007 changes to `done` only after T04 closure evidence is recorded.

## Verification

- Compare the T04 handoff package with the accepted T02 decisions and T03 rebaseline.
- Confirm `spec:product.design_records.spec_format.validation_policy` assigns local per-file detectors to W-SPEC-001 and cross-file graph checks to W-SPEC-002.
- Confirm the known local Topics column-shape row still points to W-SPEC-002 and must be synchronized by `PRODUCT-WORK-SPEC-015`.
- Confirm `spec:product.design_records.spec_format.follow_up_boundary` uses the same two durable implementation owners.
- Confirm no PRODUCT semantic or severity text is copied into DRMCP authority.
- Confirm `PRODUCT-WORK-SPEC-015` remains the PRODUCT-side synchronization owner.
- Confirm W-SPEC-001 and W-SPEC-002 remain `not_started`.
- Confirm W007 remains `in_progress` through independent review.
- Confirm PRODUCT-WORK-SPEC-015 and hub T07 remain `not_started` during T04.
- Run the targeted Git status and whitespace commands recorded in Evidence.
- Run the independent T04 review prompt recorded in Evidence.

Task and Work Item records are outside the strict spec-format validator scope.
No repository-local command result may be recorded unless externally executed.
Repository-wide clean status must not be inferred.

## Evidence

### Pre-creation inventory

The exact directory `drmcp/records/tasks/mcp/` was listed before T04 creation.

Existing `DRMCP-TASK-MCP-007-*` records were:

- `DRMCP-TASK-MCP-007-01-establish-validation-work-item-audit-baseline.md`;
- `DRMCP-TASK-MCP-007-02-decide-validation-work-item-dispositions.md`;
- `DRMCP-TASK-MCP-007-03-rebaseline-retained-validation-work-items.md`.

No `DRMCP-TASK-MCP-007-04*` record existed.
The canonical new ID is `DRMCP-TASK-MCP-007-04`.
No repository-wide traversal was used.

### Accepted upstream state

- T01 audit: `done`; independent review `PASS`.
- T02 disposition decision: `done`; independent review `PASS`.
- T03 rebaseline: `done`; independent review `PASS`.
- T03 post-closure whitespace result: `tracked_exit=0`, `untracked_exit=1`.
- No whitespace error or exit code `2` or greater was reported.
- W007 remains `in_progress` until this Task completes.

### Accepted PRODUCT handoff package

| PRODUCT validation area | accepted DRMCP implementation owner |
|---|---|
| Per-file H1, metadata, spec-kind, `contract_class`, required-section, local Topics shape, path-derived ID, and front-matter checks | `DRMCP-WORK-SPEC-001` |
| Topics edge extraction, exact child lookup, parent consistency, duplicate authoritative-parent detection, and cycle detection | `DRMCP-WORK-SPEC-002` |

The target identities are retained rather than replaced.
No superseding or absorbed Work Item ID exists.

The retained owner boundary includes these fixed relations:

- W-SPEC-001 consumes W003 parsed state and W006 validation and diagnostic contracts.
- W-SPEC-002 consumes the accepted W-SPEC-001 detector result boundary.
- W-SPEC-002 also consumes W003 active-index state and W006 validation and diagnostic contracts.
- W008 owns shared fixtures, manifests, and fixture-local checks.
- W009 owns general current-read implementation and tests outside the retained validators.
- PRODUCT owns rule text and severity semantics.

### PRODUCT pointer assessment

`spec:product.design_records.spec_format.follow_up_boundary` already names the accepted durable owners:

- `DRMCP-WORK-SPEC-001` for parser-aware per-file validation;
- `DRMCP-WORK-SPEC-002` for cross-file Topics graph validation.

`spec:product.design_records.spec_format.validation_policy` is not fully aligned at rule granularity.

The known mismatch is:

| validation rule | current PRODUCT owner | accepted owner |
|---|---|---|
| Invalid local `## Topics` table columns, including canonical `file` instead of `ref` | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-001` |

The row is a local table-shape detector.
T03 assigns local Topics presence, header, and column-shape detection to W-SPEC-001.
W-SPEC-002 consumes the accepted W-SPEC-001 detector result and does not own local table-shape detection.

T04 does not edit the PRODUCT record because PRODUCT-side pointer synchronization belongs to `PRODUCT-WORK-SPEC-015`.

`PRODUCT-WORK-SPEC-015` must:

- inventory every repeated validation-owner pointer;
- change the known local Topics column-shape owner from W-SPEC-002 to W-SPEC-001;
- preserve the PRODUCT-owned rule text and severity;
- synchronize any additional stale pointer or wording found by its scoped inventory;
- record explicit no-change evidence for pointers that already match;
- run scoped PRODUCT validation and cross-owner review;
- record PRODUCT-side closure evidence.

T04 does not assume that every PRODUCT pointer-only record has already been inventoried.
That remaining inventory belongs to PRODUCT-WORK-SPEC-015 T01.

### Lifecycle boundary

After T04 closure synchronization:

- `DRMCP-TASK-MCP-007-04`: `done`;
- `DRMCP-WORK-MCP-007`: `done`;
- `DRMCP-WORK-SPEC-001`: `not_started`;
- `DRMCP-WORK-SPEC-002`: `not_started`;
- `PRODUCT-WORK-SPEC-015`: `not_started`;
- `DRMCP-TASK-MCP-001-07`: `not_started`.

PRODUCT-WORK-SPEC-015 may now begin.
Hub T07 remains a later cross-owner lifecycle tracker.

### Changed-file manifest

| path | change |
|---|---|
| `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-04-record-product-handoff-and-close-validation-disposition.md` | New T04 Task; closed as `done` after limited re-review PASS. |
| `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md` | Record final handoff evidence and close W007 as `done`. |

### Recheck-only manifest

| record or authority | result |
|---|---|
| `DRMCP-TASK-MCP-007-01` | Accepted upstream Task; no edit. |
| `DRMCP-TASK-MCP-007-02` | Accepted upstream Task; no edit. |
| `DRMCP-TASK-MCP-007-03` | Accepted upstream Task; no edit. |
| `DRMCP-WORK-SPEC-001` | Accepted retained owner; remains `not_started`; no edit. |
| `DRMCP-WORK-SPEC-002` | Accepted retained owner; remains `not_started`; no edit. |
| `DRMCP-WORK-MCP-003` through `DRMCP-WORK-MCP-006` | Fixed accepted boundaries; no edit. |
| `DRMCP-WORK-MCP-008` | Shared fixture owner; no edit. |
| `DRMCP-WORK-MCP-009` | General current-read implementation owner; no edit. |
| `DRMCP-REQ-MCP-001` | Reciprocal links already present; no edit. |
| PRODUCT validation-policy | Read only; known local Topics column-shape owner mismatch handed to `PRODUCT-WORK-SPEC-015`; no T04 edit. |
| PRODUCT follow-up-boundary | Ownership authority read only; no edit. |
| `PRODUCT-WORK-SPEC-015` | Downstream owner; remains `not_started`; no edit. |
| `DRMCP-TASK-MCP-001-07` | Later hub tracker; remains `not_started`; no edit. |

### Initial independent T04 review result

Initial verdict: `NEEDS REVISION`.

- Blocking findings: none.
- Major finding F-MAJ-01: the local Topics table column-shape row in PRODUCT validation-policy points to W-SPEC-002, but the accepted T03 boundary assigns local table-shape detection to W-SPEC-001.
- Minor findings: none.
- Advisory A-01: repository-local Git commands were not independently rerun.
- Advisory A-02: the Task inventory ran before `prompt_chappy.md` was read, but no write occurred before instructions and authoring standards were read.

F-MAJ-01 correction:

- Removed the claim that PRODUCT validation-policy is fully aligned.
- Recorded the exact mismatched rule and current owner.
- Recorded `DRMCP-WORK-SPEC-001` as the accepted owner for the local Topics column-shape detector.
- Preserved T04's no-PRODUCT-edit boundary.
- Handed the known owner-pointer synchronization to `PRODUCT-WORK-SPEC-015`.
- Required PRODUCT rule text and severity to remain unchanged.
- Updated the recheck-only disposition and limited re-review prompt.

F-MAJ-01 was closed by targeted whitespace verification and independent limited re-review PASS.
T04 and W007 changed to `done` on 2026-06-28.

### Independent limited re-review result

Verdict: `PASS`.

- Previous finding F-MAJ-01: `CLOSED`.
- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- Advisory A-01: repository-local Git commands were not independently rerun.
- Advisory A-02: the initial inventory ordering deviation remained non-blocking.
- PRODUCT pointer correction assessment: `PASS`.
- PRODUCT-WORK-SPEC-015 handoff correction assessment: `PASS`.
- Regression assessment: `PASS`.
- Changed-file manifest assessment: `PASS`.
- Validation evidence assessment: `PASS`.
- Lifecycle assessment: `PASS`.
- T04 closure readiness: `READY`.
- W007 closure readiness: `READY AFTER T04 CLOSURE SYNCHRONIZATION`.
- PRODUCT-WORK-SPEC-015 start readiness: `READY AFTER W007 IS done`.

The accepted downstream input is:

- W-SPEC-001 and W-SPEC-002 remain the durable owner set;
- the known local Topics column-shape pointer must move from W-SPEC-002 to W-SPEC-001;
- PRODUCT rule text and severity remain unchanged;
- PRODUCT-WORK-SPEC-015 must complete repeated-pointer inventory, scoped validation, and cross-owner review.

### Validation boundary

Repository-local commands are not available through the filesystem MCP and were not executed by this assistant.

The exact Task directory inventory ran before `prompt_chappy.md` was read in this session.
No repository write occurred before `prompt_chappy.md` and the applicable authoring standards were read.
The ordering deviation does not change the inventory result, but independent review should treat it as an explicit advisory.

The user executed the targeted two-path status and whitespace commands from the repository root.

Targeted status matched the T04 changed-file manifest:

- W007 was tracked and modified;
- the T04 Task was untracked;
- no additional path was included in the targeted status result.

External whitespace results were:

- `tracked_exit=0` for W007;
- `untracked_exit=1` for the new T04 Task;
- no whitespace error reported;
- no exit code `2` or greater reported;
- LF-to-CRLF working-copy warnings for both files were non-blocking.

The untracked exit code `1` is the expected difference result against `NUL`.

Task and Work Item records are outside the strict spec-format validator scope.
No repository validator or automated test applies to the T04 documentation-only change.
Repository-wide clean status is not inferred.

The F-MAJ-01 correction changed the checked T04 Task and W007 bytes.
The targeted whitespace result supplied to the limited reviewer was `tracked_exit=0` and `untracked_exit=1`, with no whitespace error or exit code `2` or greater.
The result was not written back before review.

This closure synchronization changes the checked T04 Task and W007 bytes again.
One post-closure targeted whitespace check remains external and must not be written back into either checked file.

### External commands required

Run from:

`C:\Users\imved\projects\brewprint`

```powershell
git status --short -- `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-04-record-product-handoff-and-close-validation-disposition.md `
  drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md
```

For an untracked T04 Task and tracked W007:

```powershell
$trackedPath = "drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md"
$untrackedPath = "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-04-record-product-handoff-and-close-validation-disposition.md"

git diff --check -- $trackedPath
$tracked_exit = $LASTEXITCODE

git diff --no-index --check -- NUL $untrackedPath
$untracked_exit = $LASTEXITCODE

"tracked_exit=$tracked_exit"
"untracked_exit=$untracked_exit"
```

Expected result:

- `tracked_exit=0`;
- `untracked_exit=1`;
- no whitespace error;
- no exit code `2` or greater.

If T04 is staged, use `git diff --cached --check -- <T04 path>`.
LF-to-CRLF warnings are non-blocking when no whitespace error exists.

### Independent T04 review prompt

```text
C:\Users\imved\projects\brewprint

`DRMCP-TASK-MCP-007-04`の独立handoff and closure limited re-reviewを行う。

ファイルは変更しないこと。
DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
sandboxへrepositoryを複製しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。
repository-wide clean statusを推測しないこと。
無制限なrepository traversalを行わないこと。

## 最初に読む

- `prompt_chappy.md`
- `product/records/spec/design-records/authoring-standards/task-authoring.md`
- `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
- `product/records/spec/design-records/authoring-standards/writing-standard.md`
- `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## T04 and accepted DRMCP state

- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-01-establish-validation-work-item-audit-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-02-decide-validation-work-item-dispositions.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-03-rebaseline-retained-validation-work-items.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-04-record-product-handoff-and-close-validation-disposition.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md`
- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md`

## PRODUCT handoff boundary

- `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`
- `product/records/spec/design-records/spec-format/follow-up-boundary.md`
- `product/records/requirements/PRODUCT-REQ-SPEC-001-mcp-readable-spec-format-and-topic-tree.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md`

## Fixed non-overlap records

- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-003-current-discovery-and-active-index-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-005-resolver-and-configured-legacy-fallback-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-009-current-format-read-implementation.md`

## Review scope

F-MAJ-01 correctionの限定re-reviewを行うこと。
T04のhandoff package全体を再設計しないこと。
PRODUCT record、hub T07、W-SPEC-001/002、implementation、fixtureを編集または開始しないこと。

Previous finding:

- F-MAJ-01: local Topics table column-shape owner mismatch was omitted from the handoff.

F-MAJ-01が正確に修正され、regressionがないかを確認すること。

以下を確認すること。

- T01、T02、T03が`done`でaccepted reviewを持つ。
- T04のpre-creation inventoryがT01-T03のみを記録している。
- accepted owner targetがW-SPEC-001とW-SPEC-002の2件に一意である。
- W-SPEC-001がper-file validation ownerである。
- W-SPEC-002がTopics and graph validation ownerである。
- W-SPEC-002からW-SPEC-001へのdependencyが維持されている。
- PRODUCT semantic rulesとseverity policyをDRMCP authorityへ複製していない。
- validation-policyのlocal Topics column-shape rowがW-SPEC-002を指す既知の不一致として記録されている。
- accepted boundaryではそのrowのownerがW-SPEC-001である。
- follow-up-boundaryが同じ2 durable ownerを使用する。
- T04でPRODUCT editを行わず、PRODUCT-WORK-SPEC-015へ既知の同期対象を引き渡す判断が妥当である。
- PRODUCT-WORK-SPEC-015がlocal Topics rowの同期、repeated pointer inventory、追加同期、no-change evidence、PRODUCT validation、cross-owner reviewを所有する。
- PRODUCT rule textとseverityを変更しないhandoffになっている。
- W008 fixture ownershipとW009 general current-read ownershipが維持されている。
- W-SPEC-001/002、PRODUCT-WORK-SPEC-015、hub T07が開始されていない。
- changed-file manifestがT04とW007の2件だけである。
- recheck-only manifestが妥当である。
- T04とW007がreview前に`done`へ変更されていない。
- validation evidenceを捏造していない。
- TaskとWork Item authoring shapeが妥当である。

repository-local commandを実行できる場合、T04に記録された2-path statusとwhitespace commandを実行すること。

期待値:

- `tracked_exit=0`
- `untracked_exit=1`
- whitespace errorなし
- exit code 2以上なし

LF-to-CRLF warningはwhitespace errorがなければnon-blocking。

出力形式:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. PRODUCT pointer correction assessment
8. PRODUCT-WORK-SPEC-015 handoff correction assessment
9. Regression assessment
10. Changed-file manifest assessment
11. Validation evidence assessment
12. Lifecycle assessment
13. T04 closure readiness
14. W007 closure readiness
15. PRODUCT-WORK-SPEC-015 start readiness
```

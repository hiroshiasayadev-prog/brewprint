# DRMCP-TASK-MCP-007-03: Rebaseline retained validation Work Items

- **id**: DRMCP-TASK-MCP-007-03
- **status**: done
- **date**: 2026-06-28
- **work_item**: DRMCP-WORK-MCP-007
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-007-02
- **outputs**:
  - DRMCP-WORK-SPEC-001
  - DRMCP-WORK-SPEC-002
  - DRMCP-REQ-MCP-001
  - DRMCP-WORK-MCP-007

## Goal

Rebaseline the two retained DRMCP validation implementation Work Items and synchronize their DRMCP workflow links.

Preserve the accepted T02 disposition, ownership split, and PRODUCT handoff target set.

## Work

### Accepted T02 decisions

- Retain `DRMCP-WORK-SPEC-001` as the parser-aware per-file validation implementation owner.
- Retain `DRMCP-WORK-SPEC-002` as the cross-file Topics graph validation implementation owner.
- Keep both Work Items separate.
- Make W-SPEC-002 depend on the accepted W-SPEC-001 detector boundary and callable result contract.
- Use `DRMCP-REQ-MCP-001` as `source_requirement` for both retained Work Items.
- Keep PRODUCT semantic rules and severity policy in PRODUCT authority records.
- Keep shared fixtures and fixture-local checks in `DRMCP-WORK-MCP-008`.
- Keep per-file detector tests in W-SPEC-001.
- Keep Topics graph tests in W-SPEC-002.
- Keep general current-read tests outside retained validators in `DRMCP-WORK-MCP-009`.
- Preserve `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002` as the exact PRODUCT pointer targets.

### Rebaseline work

- Record the exact pre-creation `DRMCP-TASK-MCP-007-*` inventory.
- Rebaseline W-SPEC-001 to the current Work Item metadata and section shape.
- Rebaseline W-SPEC-002 to the current Work Item metadata and section shape.
- Remove obsolete migration-era dependencies and assumptions.
- Preserve W003 parsed-state and canonical-identity authority.
- Preserve W006 validation-execution and diagnostic-representation authority.
- Express the W-SPEC-002 dependency through Work Item prose and Task flow.
- Add both retained Work Items to `DRMCP-REQ-MCP-001.work_items` without reordering existing entries.
- Add this Task to W007 and record the rebaseline result.
- Recheck W008 and W009 before deciding whether either requires synchronization.
- Record changed-file and recheck-only manifests.
- Record applicable validation boundaries and external whitespace commands.
- Record a ready-to-run independent T03 review prompt.

### Explicit non-goals

This Task does not:

- start T04;
- start or edit `PRODUCT-WORK-SPEC-015`;
- start or edit `DRMCP-TASK-MCP-001-07`;
- implement parser-aware detectors;
- implement Topics graph algorithms;
- create or edit fixture files;
- edit PRODUCT authority records or owner pointers;
- edit W003 through W006;
- close W007;
- change this Task to `done` before scoped verification, independent review, and finding correction.

## Done condition

- The exact pre-creation inventory is recorded.
- W-SPEC-001 keeps its ID and has current canonical metadata and sections.
- W-SPEC-002 keeps its ID and has current canonical metadata and sections.
- Both Work Items use `DRMCP-REQ-MCP-001` as `source_requirement`.
- W-SPEC-001 owns the complete retained per-file detector and test scope.
- W-SPEC-002 owns the complete retained Topics graph detector and test scope.
- W-SPEC-002 depends on the accepted W-SPEC-001 detector boundary.
- W-SPEC-002 does not duplicate local detector ownership.
- W003 and W006 authorities are consumed without duplication.
- PRODUCT semantic authorities are referenced without copied ownership.
- W008 and W009 have explicit change or no-change dispositions.
- `DRMCP-REQ-MCP-001` contains reciprocal entries for both retained Work Items.
- W007 lists this Task and records the T03 synchronization result.
- The exact PRODUCT pointer target set remains unchanged.
- No PRODUCT record, W003-W006 record, hub T07 record, implementation source, or fixture changes.
- The changed-file and recheck-only manifests match the actual filesystem writes.
- Scoped validation and whitespace evidence are available before closure.
- Independent review reports no unresolved blocking, major, or minor finding.
- T03 remains `in_progress` until review and correction complete.
- W007 remains `in_progress` until T04 closure work completes.

## Verification

- Compare both rebaselined Work Items with the current Work Item authoring standard.
- Compare W-SPEC-001 scope with PRODUCT document-shape, ID-as-ref, and validation-policy authorities.
- Compare W-SPEC-002 scope with PRODUCT Topics-table, ID-as-ref, and validation-policy authorities.
- Confirm W003 parsing, identity, source-state, and active-index ownership is not reopened.
- Confirm W006 validation execution, diagnostic category, severity, ordering, location, and representation ownership is not reopened.
- Confirm W008 fixture ownership and W009 retained-validator exclusion remain sufficient.
- Confirm reciprocal Requirement linkage and absence of duplicate entries.
- Confirm no PRODUCT record changed.
- Confirm T03 remains `in_progress` through independent review and W007 remains `in_progress` after T03 closure.
- Run targeted Git status and whitespace checks from the repository root.
- Run the independent review prompt recorded in Evidence.

Task and Work Item records are outside the strict spec-format validator scope.
The changed Requirement is not assumed to have an applicable repository validator without an identified exact command.
No repository-local command result may be recorded unless externally executed.

## Evidence

### Pre-creation inventory

The exact directory `drmcp/records/tasks/mcp/` was listed once before Task creation.

Existing `DRMCP-TASK-MCP-007-*` records were:

- `DRMCP-TASK-MCP-007-01-establish-validation-work-item-audit-baseline.md`;
- `DRMCP-TASK-MCP-007-02-decide-validation-work-item-dispositions.md`.

No `DRMCP-TASK-MCP-007-03*` record existed.
The canonical new ID was therefore `DRMCP-TASK-MCP-007-03`.
No repository-wide search or traversal was used.

### Rebaseline results

`DRMCP-WORK-SPEC-001` keeps its existing identity and remains `not_started`.
Its source Requirement is now `DRMCP-REQ-MCP-001`.
The record now owns per-file detector implementation, detector tests, and runtime integration tests.
It consumes W003 parsed state, W006 validation and diagnostic contracts, PRODUCT semantic authorities, and W008 fixtures.

`DRMCP-WORK-SPEC-002` keeps its existing identity and remains `not_started`.
Its source Requirement is now `DRMCP-REQ-MCP-001`.
The record now owns accepted Topics-edge extraction, exact canonical child lookup, parent consistency, duplicate-parent detection, cycle detection, graph tests, and runtime integration tests.
It consumes the accepted W-SPEC-001 detector result boundary.

Obsolete migration-era dependencies and assumptions were removed from both records.
No implementation source or test file changed.

### Dependency assessment

W-SPEC-001 and W-SPEC-002 remain separate Work Items.

W-SPEC-002 depends on W-SPEC-001 through canonical prose and Task flow.
Work Item metadata does not add a `depends_on` field.

W-SPEC-002 does not own:

- H1 count or shape detection;
- metadata detection;
- spec-kind detection;
- `contract_class` detection;
- required-section detection;
- local Topics table presence, header, or row-shape detection.

### Fixture and test ownership

| artifact or verification | owner |
|---|---|
| Shared fixture files | `DRMCP-WORK-MCP-008` |
| Fixture manifests | `DRMCP-WORK-MCP-008` |
| Fixture-local structural checks | `DRMCP-WORK-MCP-008` |
| Per-file detector tests | `DRMCP-WORK-SPEC-001` |
| Per-file runtime integration tests | `DRMCP-WORK-SPEC-001` |
| Topics graph algorithm tests | `DRMCP-WORK-SPEC-002` |
| Topics graph runtime integration tests | `DRMCP-WORK-SPEC-002` |
| General current-read tests outside retained validators | `DRMCP-WORK-MCP-009` |

### Requirement synchronization

`DRMCP-REQ-MCP-001.work_items` receives these reciprocal entries:

- `DRMCP-WORK-SPEC-001`;
- `DRMCP-WORK-SPEC-002`.

Existing entries remain in their original order.
No duplicate entry is added.
No Requirement body change is required.

### W008 and W009 disposition

`DRMCP-WORK-MCP-008` requires no T03 change.
Its Boundary already assigns shared fixtures, manifests, and fixture-local checks to W008.
It already excludes production implementation and runtime behavior assertions.

`DRMCP-WORK-MCP-009` requires no T03 change.
Its Boundary already excludes retained spec-format validator implementation.
It already assigns only general current-read implementation and tests to W009.

Both records remain recheck-only.

### PRODUCT handoff preservation

The exact downstream pointer target set remains:

| PRODUCT pointer class | target |
|---|---|
| Per-file validation rules | `DRMCP-WORK-SPEC-001` |
| Topics and graph validation rules | `DRMCP-WORK-SPEC-002` |

No PRODUCT record changed.
`PRODUCT-WORK-SPEC-015` remains `not_started`.

### Changed-file manifest

| path | change |
|---|---|
| `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-03-rebaseline-retained-validation-work-items.md` | New T03 Task; closed as `done` after independent review. |
| `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md` | Full canonical rebaseline for retained per-file validator implementation. |
| `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md` | Full canonical rebaseline for retained Topics graph implementation. |
| `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md` | Add two reciprocal Work Item entries. |
| `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md` | Add T03 and synchronize rebaseline and handoff evidence. |

No conditional W008 or W009 change was required.

### Recheck-only manifest

| record or authority | result |
|---|---|
| `DRMCP-WORK-MCP-003` | Rechecked; no edit. |
| `DRMCP-WORK-MCP-004` | Rechecked; no edit. |
| `DRMCP-WORK-MCP-005` | Rechecked; no edit. |
| `DRMCP-WORK-MCP-006` | Rechecked; no edit. |
| `DRMCP-WORK-MCP-008` | Rechecked; no edit. |
| `DRMCP-WORK-MCP-009` | Rechecked; no edit. |
| PRODUCT semantic authorities | Authority read only; no edit. |
| `PRODUCT-WORK-SPEC-015` | Rechecked; remains `not_started`. |
| `DRMCP-TASK-MCP-001-07` | Rechecked; remains `not_started`. |

### Validation boundary

Repository-local commands were not available through the filesystem MCP and were not executed by this assistant.
Repository-wide clean status is not inferred.

The user executed the targeted status and whitespace commands from the repository root.

Targeted status matched the T03 changed-file manifest:

- four existing files were tracked and modified;
- the T03 Task was untracked;
- no additional path was included in the targeted status result.

External whitespace results were:

- `tracked_exit=0` for W-SPEC-001, W-SPEC-002, `DRMCP-REQ-MCP-001`, and W007;
- `untracked_exit=1` for the new T03 Task;
- no whitespace error reported;
- no exit code `2` or greater reported;
- LF-to-CRLF working-copy warnings for all five files were non-blocking.

The untracked exit code `1` is the expected difference result against `NUL`.

Task and Work Item records are outside the strict spec-format validator scope.
No exact applicable repository validator command was identified for the changed Requirement.
No repository validator or automated test result is claimed.

This Evidence synchronization changed the checked T03 Task and W007 bytes.
The final pre-review targeted whitespace check repeated `tracked_exit=0` and `untracked_exit=1` with no whitespace error or exit code `2` or greater.
That result was supplied directly to the reviewer and was not written back before review.

### Lifecycle state

- `DRMCP-TASK-MCP-007-03`: `done`.
- `DRMCP-WORK-MCP-007`: `in_progress`.
- `PRODUCT-WORK-SPEC-015`: `not_started`.
- `DRMCP-TASK-MCP-001-07`: `not_started`.
- T04: not created and not started.

T03 closure prerequisites are complete.
T04 handoff is ready after the required post-closure whitespace check; T04 remains uncreated and not started.

### Residual limitations

- Repository-wide Git state remains unknown; only the five T03 paths were inspected.
- Repository-local validators and tests were not run.
- Runtime implementation and automated tests remain future child-Task work under W-SPEC-001 and W-SPEC-002.
- PRODUCT pointer synchronization remains deferred to `PRODUCT-WORK-SPEC-015`.
- One post-closure targeted whitespace check remains external and must not be written back into the checked files.

### External commands required

Run from:

`C:\Users\imved\projects\brewprint`

First inspect only the T03 changed paths:

```powershell
git status --short -- `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-03-rebaseline-retained-validation-work-items.md `
  drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md `
  drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md `
  drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md `
  drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md
```

For the four existing tracked files:

```powershell
$trackedPaths = @(
  "drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md",
  "drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md",
  "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md",
  "drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md"
)

git diff --check -- $trackedPaths
$tracked_exit = $LASTEXITCODE
"tracked_exit=$tracked_exit"
```

For the new untracked T03 Task:

```powershell
$taskPath = "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-03-rebaseline-retained-validation-work-items.md"

git diff --no-index --check -- NUL $taskPath
$untracked_exit = $LASTEXITCODE
"untracked_exit=$untracked_exit"
```

Expected no-whitespace-error result:

- `tracked_exit=0`;
- `untracked_exit=1`;
- no whitespace error;
- no exit code `2` or greater.

If the Task is staged, use:

```powershell
git diff --cached --check -- `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-03-rebaseline-retained-validation-work-items.md
```

If any existing changed file is staged, include it in a targeted `git diff --cached --check` call.
LF-to-CRLF warnings are non-blocking when no whitespace error exists.

### Independent T03 review result

Independent review verdict: `PASS`.

- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- T02 decisions, Work Item separation, and the W-SPEC-002 dependency were accepted.
- Both retained Work Item rebaselines and their test ownership were accepted.
- W003 through W006 non-duplication boundaries were accepted.
- W008 fixture ownership and W009 general current-read ownership were accepted.
- Reciprocal Requirement synchronization was accepted.
- PRODUCT semantic authority and the two-target handoff set were preserved.
- Changed-file and recheck-only manifests were accepted.
- External whitespace evidence was accepted.
- Advisory A-01 records that Git commands were not independently rerun by the reviewer.
- Advisory A-02 records that T04 absence was confirmed from lifecycle records without another directory inventory.
- Neither advisory changes the verdict or closure readiness.
- T03 closure readiness: `READY`.
- T04 handoff readiness: `READY AFTER T03 CLOSURE`.
- T03 changed to `done` on 2026-06-28.
- W007 remains `in_progress`.
- This closure synchronization changes the checked T03 Task and W007 bytes.
- One post-closure targeted whitespace check must run after these edits.
- The post-closure result must remain external and must not be written back into either checked file.

### Independent T03 review prompt

```text
C:\Users\imved\projects\brewprint

`DRMCP-TASK-MCP-007-03`の独立rebaseline reviewを行う。

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
- `product/records/spec/design-records/authoring-standards/requirement-authoring.md`
- `product/records/spec/design-records/authoring-standards/writing-standard.md`
- `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## T03 and accepted baseline

- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-01-establish-validation-work-item-audit-baseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-02-decide-validation-work-item-dispositions.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-03-rebaseline-retained-validation-work-items.md`
- `drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md`

## Rebaselined Work Items

- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md`

## Fixed non-overlap and downstream boundaries

- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-003-current-discovery-and-active-index-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-005-resolver-and-configured-legacy-fallback-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-006-validation-diagnostics-and-path-exposure-contract-realignment.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-008-current-and-legacy-read-fixture-baseline.md`
- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-009-current-format-read-implementation.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md`

## PRODUCT authority and handoff boundary

- `product/records/requirements/PRODUCT-REQ-SPEC-001-mcp-readable-spec-format-and-topic-tree.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`
- `product/records/spec/design-records/spec-format/document-shape.md`
- `product/records/spec/design-records/spec-format/topics-table.md`
- `product/records/spec/design-records/spec-format/spec-id-as-ref.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`
- `product/records/spec/design-records/spec-format/follow-up-boundary.md`

PRODUCT recordsはauthority inputとして読むだけとし、編集しないこと。
W003-W006はaccepted fixed inputであり、明確なregressionがない限り再設計または再審査しないこと。

## Review scope

T03のrebaselineとDRMCP-side synchronizationだけをreviewすること。
T04、PRODUCT-WORK-SPEC-015、hub T07、実装、fixture authoringを開始しないこと。

以下を確認すること。

- T02のaccepted decisionが変更されていない。
- W-SPEC-001とW-SPEC-002の既存IDと`retain` identityが維持されている。
- 両Work Itemがcurrent canonical metadataとsection shapeを持つ。
- 両Work Itemの`source_requirement`が`DRMCP-REQ-MCP-001`である。
- W-SPEC-001にreal H1 count/shape、metadata、spec kind、`contract_class`、required section、local Topics shape、path-derived ID mismatch、front-matter policyのdetector scopeがある。
- W-SPEC-001にdetector unit testとruntime integration test ownershipがある。
- W-SPEC-001がW003 parsed stateとcanonical identityを再定義していない。
- W-SPEC-001がW006 validation execution、diagnostic category、severity、ordering、location、representationを再定義していない。
- W-SPEC-002に`title/kind/ref/summary` row、canonical child lookup、declaring Index/Overview parent、child parent-marker consistency、duplicate parent、cycle detectionがある。
- W-SPEC-002が`file`、row-level `parent`、alias、redirect、temporary tooling dependencyを復活させていない。
- W-SPEC-002からW-SPEC-001へのdependencyがcanonical proseとTask flowで明確である。
- Work Item metadataに`depends_on`を追加していない。
- W-SPEC-002がH1、metadata、spec kind、contract class、required section、local Topics shape detectorを重複実装しない。
- PRODUCT semantic rule textとseverity policyをDRMCP authorityとして複製していない。
- W008がshared fixtures、manifests、fixture-local checksを一意に所有する。
- W-SPEC-001がper-file detector testsを所有する。
- W-SPEC-002がgraph algorithm testsを所有する。
- W009がgeneral current-read testsだけを所有し、retained validator implementationを吸収しない。
- `DRMCP-REQ-MCP-001.work_items`にW-SPEC-001とW-SPEC-002が重複なしで追加されている。
- 既存Requirement entryが削除または並べ替えられていない。
- PRODUCT pointer target setがW-SPEC-001とW-SPEC-002のまま維持されている。
- W008とW009の非変更判断が既存Boundaryに基づいている。
- changed-file manifestが実際の5 recordsと一致する。
- recheck-only manifestが妥当である。
- T03が`in_progress`である。
- W007が`in_progress`である。
- PRODUCT-WORK-SPEC-015が`not_started`である。
- hub `DRMCP-TASK-MCP-001-07`が`not_started`である。
- T04が未作成・未開始である。
- validation evidenceを捏造していない。
- TaskとWork Itemのauthoring shapeが妥当である。

Repository-local commandを実行できる場合、T03 Taskに記録された対象path限定のstatusとwhitespace commandを実行すること。

期待値:

- existing tracked files: `tracked_exit=0`
- untracked T03 Task: `untracked_exit=1`
- whitespace errorなし
- exit code 2以上なし

Taskがstaged済みの場合はTaskへ`git diff --cached --check`を使うこと。
LF-to-CRLF warningはwhitespace errorがなければnon-blocking。
TaskとWork Itemはstrict spec-format validator対象外。
Requirement validatorの適用可否を推測しないこと。

出力形式:

1. Verdict: PASS / NEEDS REVISION
2. Blocking findings
3. Major findings
4. Minor findings
5. Advisories
6. T02 decision preservation assessment
7. W-SPEC-001 rebaseline assessment
8. W-SPEC-002 rebaseline assessment
9. Separation and dependency assessment
10. W003-W006 non-duplication assessment
11. Fixture and test ownership assessment
12. Source-Requirement and reciprocal-link assessment
13. PRODUCT handoff preservation assessment
14. W008/W009 disposition assessment
15. Changed-file manifest assessment
16. Recheck-only manifest assessment
17. Validation evidence assessment
18. Lifecycle assessment
19. T03 closure readiness
20. T04 handoff readiness
```

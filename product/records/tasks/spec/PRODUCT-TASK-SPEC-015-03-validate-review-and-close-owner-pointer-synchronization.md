# PRODUCT-TASK-SPEC-015-03: Validate, review, and close owner-pointer synchronization

- **id**: PRODUCT-TASK-SPEC-015-03
- **status**: done
- **date**: 2026-06-28
- **work_item**: PRODUCT-WORK-SPEC-015
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-015-02
- **outputs**:
  - PRODUCT-WORK-SPEC-015
  - DRMCP-TASK-MCP-001-07

## Goal

Validate the accepted validation-owner synchronization and obtain an independent cross-owner review.

Close W015 only after all findings are resolved, then record the required handoff state in hub T07 without closing the hub from the PRODUCT side.

## Work

- Confirm the accepted T01 inventory and T02 synchronization results.
- Recheck the two synchronized owner rows against the retained W-SPEC-001 and W-SPEC-002 boundaries.
- Confirm rule text, severity, row order, and adjacent graph-owner rows remain unchanged.
- Recheck `follow-up-boundary.md`, `spec-id-as-ref.md`, and `topics-table.md` without changing them unless a concrete contradiction exists.
- Assess every W015 Completion Condition separately.
- Confirm `PRODUCT-REQ-SPEC-001` lists W015 in `work_items`.
- Record the exact T03 changed-file manifest and cumulative review target set.
- Record scoped external validation, targeted status, and whitespace commands.
- Obtain an independent final review using the prompt in Evidence.
- Apply only scoped corrections required by review findings.
- Keep this Task and W015 open until validation and review complete.
- After review acceptance, close W015 and this Task.
- After W015 closure, record the owner-side handoff state in hub T07 without marking T07 done.

This Task does not:

- redesign PRODUCT validation rules or severity;
- change the accepted W-SPEC-001 and W-SPEC-002 implementation boundary;
- start W-SPEC-001 or W-SPEC-002;
- close hub T07 from the PRODUCT side;
- modify unrelated working-tree changes;
- infer repository-wide clean status.

## Done condition

- T01 and T02 accepted results remain valid.
- The local Topics column-shape row points to W-SPEC-001.
- The parent grammar violation row points to W-SPEC-001.
- Unresolved child refs, duplicate authoritative parents, parent consistency, and topic cycles remain W-SPEC-002 concerns.
- Rule text, severity, row order, and table formatting remain unchanged.
- `follow-up-boundary.md` remains unchanged unless review identifies a concrete contradiction.
- W-SPEC-001 remains the per-file implementation owner.
- W-SPEC-002 remains the cross-file Topics graph implementation owner.
- No duplicated implementation authority remains between W-SPEC-001 and W-SPEC-002.
- `PRODUCT-REQ-SPEC-001` lists W015 in `work_items`.
- Scoped validation and targeted whitespace checks pass.
- Independent review reports no unresolved blocking, major, or minor finding.
- Final evidence records the target set, changed files, validation results, review verdict, and residual limitations.
- W015 is marked `done` only after all Completion Conditions are satisfied.
- Hub T07 records the W015 completion handoff and remains open for separate cross-owner closure.
- This Task is marked `done` only after closure synchronization is complete.

## Verification

- Compare the current validation-policy rows with the accepted T02 synchronization set.
- Confirm W-SPEC-001 owns local per-file structure and metadata checks.
- Confirm W-SPEC-002 consumes accepted W-SPEC-001 detector results and owns only cross-file graph checks.
- Confirm `follow-up-boundary.md` still names both retained canonical Work Items.
- Confirm `spec-id-as-ref.md` owns parent grammar semantics.
- Confirm `topics-table.md` owns local Topics columns and authoritative-parent semantics.
- Confirm W-SPEC-001 and W-SPEC-002 remain `not_started`.
- Confirm W015 and hub T07 remain `in_progress` before independent review.
- Run the strict validator, targeted status, and whitespace commands recorded in Evidence.
- Run the independent final review prompt recorded in Evidence.
- Re-run affected validation and whitespace commands after any correction.

Task and Work Item records are outside the strict spec-format validator scope.
No repository-local command result may be recorded unless externally executed.
Repository-wide clean status must not be inferred.

## Evidence

### Pre-creation inventory

The exact directory `product/records/tasks/spec/` was listed before creation.

- `PRODUCT-TASK-SPEC-015-01` existed.
- `PRODUCT-TASK-SPEC-015-02` existed.
- No `PRODUCT-TASK-SPEC-015-03` or equivalent validate-review-close Task existed.
- No repository-wide traversal was used.

### Accepted T01 and T02 result

- T01 is complete and its independent review returned `PASS`.
- T02 is complete and its final independent review returned `PASS`.
- T02 closed F-MIN-01, F-MIN-02, and F-MIN-03.
- No blocking, major, or minor T02 finding remains.
- The strict validator returned `[strict]  All 1 file(s) OK.` and `validator_exit=0`.
- The final post-closure whitespace result was `tracked_exit=0` and `untracked_exit=1`.
- No whitespace error or exit code `2` or greater was reported.
- W015's stale T02 post-closure pending note was synchronized to this accepted result.

Accepted owner synchronization:

| validation rule | previous owner | accepted owner |
|---|---|---|
| Invalid local `## Topics` table columns, including canonical `file` instead of `ref` | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-001` |
| Parent grammar violation | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-001` |

Retained W-SPEC-002 concerns:

- unresolved child `ref`;
- duplicate authoritative parent;
- declaring-parent and child-marker consistency;
- topic cycle.

### Final W015 Completion Condition assessment

| Completion Condition | final assessment | closure evidence |
|---|---|---|
| Affected PRODUCT owner pointers match the accepted DRMCP disposition. | Satisfied. | Final review `PASS`. |
| Obsolete Work Item IDs do not remain as current implementation owners. | Satisfied within the reviewed scope. | Final review `PASS`. |
| PRODUCT rule text and severity remain unchanged. | Satisfied. | Final review `PASS`. |
| PRODUCT records reference canonical DRMCP Work Items rather than lifecycle Tasks. | Satisfied. | Final review `PASS`. |
| No duplicated implementation authority remains. | Satisfied. | Final review `PASS`. |
| All changed PRODUCT records pass scoped validation. | Satisfied under the recorded validation boundary. | Final pre-review whitespace passed; final review `PASS`. |
| Cross-owner independent review has no blocking or major finding. | Satisfied. | No blocking, major, or minor finding. |
| `PRODUCT-REQ-SPEC-001` lists W015. | Satisfied. | No correction required. |
| Final evidence records target set, changed files, validation results, verdict, and limitations. | Satisfied. | Closure synchronization recorded. |

### PRODUCT-REQ-SPEC-001 relation assessment

`PRODUCT-REQ-SPEC-001` already lists `PRODUCT-WORK-SPEC-015` in `work_items`.

The reciprocal source relation is correct.
No Requirement edit is required.

### Recheck-only assessment

| record | result |
|---|---|
| `spec:product.design_records.spec_format.follow_up_boundary` | Both retained implementation Work Items are named correctly. No edit required. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Parent grammar remains a local per-file metadata rule. No edit required. |
| `spec:product.design_records.spec_format.topics_table` | Local Topics columns remain PRODUCT semantics. Cross-file graph checks remain separate. No edit required. |
| `DRMCP-WORK-SPEC-001` | Remains `not_started` and owns per-file parser-aware validation. No edit required. |
| `DRMCP-WORK-SPEC-002` | Remains `not_started` and owns cross-file Topics graph validation. No edit required. |
| `DRMCP-TASK-MCP-001-07` | Remains `in_progress`. No pre-review edit required. |
| `PRODUCT-REQ-SPEC-001` | Reciprocal W015 relation is present. No edit required. |

### Target and changed-file manifests

T03 initial changed-file manifest:

| path | change |
|---|---|
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-03-validate-review-and-close-owner-pointer-synchronization.md` | New T03 Task with validation scope, Completion Condition assessment, external commands, and independent review prompt. |
| `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md` | Add the T03 relation and initial progress Evidence; synchronize the accepted T02 final whitespace result. |

Current cumulative review target set:

- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-01-inventory-validation-owner-pointers.md`;
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-02-synchronize-validation-owner-pointers.md`;
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-03-validate-review-and-close-owner-pointer-synchronization.md`;
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`;
- `product/records/spec/design-records/spec-format/validation-policy.md`;
- the explicit recheck-only records listed above.

The closure-time manifest may add only:

- scoped finding corrections, if required;
- W015 lifecycle and final Evidence;
- T03 lifecycle and final Evidence;
- hub T07 owner-side handoff Evidence after W015 closure.

Any expansion must be recorded exactly before closure.

### Validation boundary

Repository-local commands are not available through the filesystem MCP.
This assistant did not execute Git, Python, PowerShell, validator, formatter, or test commands.

Static filesystem inspection confirmed:

- the two synchronized validation-policy owner values;
- the unchanged adjacent W-SPEC-002 rows;
- the retained W-SPEC-001 and W-SPEC-002 boundary;
- the follow-up-boundary no-change result;
- the Requirement relation;
- the pre-review lifecycle state.

Static inspection is not repository-local validation.
Repository-wide clean status is unknown and must not be inferred.

### External verification result

The user executed the targeted T03 commands from the repository root.

Targeted status matched the initial T03 changed-file manifest:

- W015 was tracked and modified;
- the new T03 Task was untracked;
- no additional path appeared in the targeted result.

Strict validation result:

- `[strict]  All 1 file(s) OK.`;
- `validator_exit=0`.

Targeted whitespace result:

- `tracked_exit=0`;
- `untracked_exit=1`;
- no whitespace error;
- no exit code `2` or greater.

LF-to-CRLF working-copy warnings for W015 and T03 are non-blocking because no whitespace error was reported.
The untracked exit code `1` is the expected difference result against `NUL`.

These results were externally executed by the user and were not executed independently by this assistant.

This Evidence synchronization changed the T03 Task and W015 bytes after the initial whitespace check.
The strict validator result remained applicable because validation-policy did not change.
The user-executed final pre-review whitespace check returned `tracked_exit=0` and `untracked_exit=1`, with no whitespace error or exit code `2` or greater.

### External commands required

Run from:

`C:\Users\imved\projects\brewprint`

Strict validation of the synchronized PRODUCT spec:

```powershell
$specPath = "product/records/spec/design-records/spec-format/validation-policy.md"
python -X utf8 product/src/tools/validate_spec.py --strict --no-color $specPath
$validator_exit = $LASTEXITCODE
"validator_exit=$validator_exit"
```

Expected result:

- `[strict]  All 1 file(s) OK.`;
- `validator_exit=0`.

Targeted T03 status:

```powershell
git status --short -- `
  product/records/tasks/spec/PRODUCT-TASK-SPEC-015-03-validate-review-and-close-owner-pointer-synchronization.md `
  product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md
```

Expected paths:

- the new T03 Task;
- W015;
- no additional path in this targeted result.

Targeted whitespace checks for the current final-review bytes:

```powershell
$trackedPaths = @(
  "product/records/spec/design-records/spec-format/validation-policy.md",
  "product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md"
)
$untrackedPath = "product/records/tasks/spec/PRODUCT-TASK-SPEC-015-03-validate-review-and-close-owner-pointer-synchronization.md"

git diff --check -- $trackedPaths
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

If T03 is staged, use:

```powershell
git diff --cached --check -- `
  product/records/tasks/spec/PRODUCT-TASK-SPEC-015-03-validate-review-and-close-owner-pointer-synchronization.md
```

LF-to-CRLF warnings are non-blocking when no whitespace error exists.

### Independent final review prompt

```text
C:\Users\imved\projects\brewprint

`PRODUCT-TASK-SPEC-015-03`の独立cross-owner final reviewを行う。

ファイルは変更しないこと。
DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
sandboxへrepositoryを複製しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。
repository-wide clean statusを推測しないこと。
無制限なrepository traversalを行わないこと。
既存の無関係なworking-tree変更には触れないこと。

## 最初に読む

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/task-authoring.md`
3. `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
4. `product/records/spec/design-records/authoring-standards/writing-standard.md`
5. `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## T03 current changed files

- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-03-validate-review-and-close-owner-pointer-synchronization.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`

## Accepted T01 and T02 records

- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-01-inventory-validation-owner-pointers.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-02-synchronize-validation-owner-pointers.md`

Accepted result:

- T01 and T02 are complete.
- T02 final review returned `PASS`.
- F-MIN-01, F-MIN-02, and F-MIN-03 are closed.
- No blocking, major, or minor T02 finding remains.
- Strict validation returned `[strict]  All 1 file(s) OK.` and `validator_exit=0`.
- Final post-closure whitespace returned `tracked_exit=0` and `untracked_exit=1`.
- T03 targeted status contained only the modified W015 and untracked T03 Task.
- T03 strict validation returned `[strict]  All 1 file(s) OK.` and `validator_exit=0`.
- Initial T03 whitespace returned `tracked_exit=0` and `untracked_exit=1`, with no whitespace error or exit code `2` or greater.
- Final pre-review whitespace after Evidence synchronization returned `tracked_exit=0` and `untracked_exit=1`, with no whitespace error or exit code `2` or greater.

## Parent and synchronized spec

- `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`

Accepted owner changes:

- Invalid local `## Topics` columns: W-SPEC-002 to W-SPEC-001.
- Parent grammar violation: W-SPEC-002 to W-SPEC-001.

Required preservation:

- rule text;
- severity;
- row order;
- table formatting;
- unresolved child `ref` ownership by W-SPEC-002;
- duplicate authoritative-parent ownership by W-SPEC-002;
- declaring-parent and child-marker consistency ownership by W-SPEC-002;
- topic-cycle ownership by W-SPEC-002.

## Recheck-only records

- `product/records/spec/design-records/spec-format/follow-up-boundary.md`
- `product/records/spec/design-records/spec-format/spec-id-as-ref.md`
- `product/records/spec/design-records/spec-format/topics-table.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md`
- `product/records/requirements/PRODUCT-REQ-SPEC-001-mcp-readable-spec-format-and-topic-tree.md`

## Review scope

Confirm all of the following without reopening accepted W007, T01, or T02 decisions:

- T01 and T02 accepted results are represented accurately.
- Exactly two validation-policy owner pointers changed.
- Rule text, severity, row order, and table formatting remain unchanged.
- W-SPEC-001 owns per-file parser-aware validation.
- W-SPEC-002 owns cross-file Topics graph validation.
- W-SPEC-002 consumes accepted W-SPEC-001 detector results.
- No duplicated implementation authority remains.
- `follow-up-boundary.md` requires no change.
- `PRODUCT-REQ-SPEC-001` lists W015 in `work_items`.
- Every W015 Completion Condition is assessed correctly.
- The T03 changed-file manifest matches the actual T03 writes.
- The cumulative review target set is complete.
- Validation evidence is accurate and not fabricated.
- Repository-wide clean status is not inferred.
- T03 and W015 remain open before review acceptance.
- Hub T07 remains open and has not been closed by PRODUCT.
- The planned closure sequence is valid.
- Final evidence will contain the target set, changed files, validation results, review verdict, and residual limitations.
- No blocking, major, or minor finding remains.
- W015 closure readiness.
- Downstream hub T07 handoff readiness.

If repository-local commands are available, run the exact strict-validator, targeted status, and whitespace commands recorded in T03 Evidence.

Expected results:

- `validator_exit=0`;
- `[strict]  All 1 file(s) OK.`;
- targeted status contains only T03 and W015;
- `tracked_exit=0`;
- `untracked_exit=1`;
- whitespace errorなし;
- exit code 2以上なし.

LF-to-CRLF warnings are non-blocking when no whitespace error exists.

Output format:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. T01 and T02 accepted-result assessment
8. Exact owner-pointer assessment
9. Rule-text, severity, and row-order preservation assessment
10. W-SPEC-001 / W-SPEC-002 boundary assessment
11. Follow-up-boundary no-change assessment
12. PRODUCT-REQ-SPEC-001 relation assessment
13. W015 Completion Condition assessment
14. Changed-file manifest assessment
15. Validation evidence assessment
16. Lifecycle assessment
17. W015 closure readiness
18. Downstream hub T07 handoff readiness
```

### Independent final review result

- Verdict: `PASS`.
- F-MIN-01, F-MIN-02, and F-MIN-03 remain closed.
- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- W015 closure readiness: `READY`.
- Hub T07 handoff readiness: `READY`.
- No scoped correction was required.

### Closure synchronization

Final closure-time changed files:

- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-03-validate-review-and-close-owner-pointer-synchronization.md`;
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`;
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md`.

Final lifecycle state:

- W015 is `done`.
- W-SPEC-001 remains `not_started`.
- W-SPEC-002 remains `not_started`.
- Hub T07 remains `in_progress` and is not closed by PRODUCT.
- The PRODUCT owner-side handoff is recorded in hub T07.
- This Task is `done` after the handoff synchronization.

### Residual limitations

- Repository-local commands were not independently rerun by the reviewer.
- Byte-level Git diff comparison against pre-T02 bytes was not independently performed.
- Repository-wide working-tree state is unknown.
- This closure synchronization changes T03, W015, and hub T07 bytes, so final post-closure whitespace verification remains external.

# PRODUCT-TASK-SPEC-015-02: Synchronize validation owner pointers

- **id**: PRODUCT-TASK-SPEC-015-02
- **status**: done
- **date**: 2026-06-28
- **work_item**: PRODUCT-WORK-SPEC-015
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-015-01
- **outputs**:
  - spec:product.design_records.spec_format.validation_policy
  - PRODUCT-WORK-SPEC-015

## Goal

Synchronize the two confirmed PRODUCT validation owner pointers with the retained DRMCP implementation owners.

Prepare scoped validation evidence and an independent review prompt without changing PRODUCT validation semantics.

## Work

- Change exactly two owner cells in `spec:product.design_records.spec_format.validation_policy`.
- Move the local Topics column-shape row from `DRMCP-WORK-SPEC-002` to `DRMCP-WORK-SPEC-001`.
- Move the parent grammar violation row from `DRMCP-WORK-SPEC-002` to `DRMCP-WORK-SPEC-001`.
- Preserve validation rule text, severity, row order, and table formatting.
- Preserve W-SPEC-002 ownership for unresolved child refs, duplicate authoritative parents, parent consistency, and cycles.
- Normalize the W015 Impact Scope wording from pointer candidates to retained implementation owners.
- Add this Task to W015 and record the exact synchronization and no-change sets.
- Recheck follow-up-boundary and retained lifecycle states without editing them.
- Record the changed-file manifest, scoped validation boundary, external PowerShell, and independent review prompt.

This Task does not:

- redesign validation rules or severity;
- change `spec:product.design_records.spec_format.follow_up_boundary` without a concrete contradiction;
- start `DRMCP-WORK-SPEC-001` or `DRMCP-WORK-SPEC-002`;
- close `PRODUCT-WORK-SPEC-015` or `DRMCP-TASK-MCP-001-07`;
- edit unrelated working-tree changes;
- infer repository-wide clean status.

## Done condition

- `validation-policy.md` has exactly two owner-cell changes.
- Both changed owner cells point to `DRMCP-WORK-SPEC-001`.
- Rule text, severity, row order, and unrelated owner cells remain unchanged.
- W015 lists this Task and uses retained-owner wording for W-SPEC-001 and W-SPEC-002.
- W015 changes remain limited to the T02 relation, Evidence, and wording-only normalization.
- `follow-up-boundary.md` remains unchanged.
- W-SPEC-001 and W-SPEC-002 remain `not_started`.
- W015 and hub T07 remain `in_progress` and not closed.
- The changed-file manifest contains only the expected three files.
- Strict spec-format validation evidence exists for `validation-policy.md`.
- Targeted whitespace evidence exists for all three changed files.
- Independent review reports no unresolved blocking, major, or minor finding.
- This Task changes to `done` only after validation, review, and finding correction complete.

## Verification

- Compare both changed validation-policy rows with the T01 synchronization set.
- Confirm rule text and severity are byte-for-byte unchanged in both rows.
- Confirm unresolved child `ref`, duplicate parent declaration, and topic cycle remain owned by W-SPEC-002.
- Confirm cross-file parent consistency remains owned by W-SPEC-002 through its retained Work Item boundary.
- Confirm W015 changes are limited to the T02 relation, Evidence, and two Impact Scope descriptions.
- Confirm `follow-up-boundary.md` has no change.
- Confirm W-SPEC-001 and W-SPEC-002 remain `not_started`.
- Confirm W015 and hub T07 remain `in_progress`.
- Run the strict spec-format validator only against the changed spec file.
- Run the targeted status and whitespace commands recorded in Evidence.
- Run the independent T02 review prompt recorded in Evidence.

Task and Work Item records are outside the strict spec-format validator scope.
No repository-local command result may be recorded unless externally executed.
Repository-wide clean status must not be inferred.

## Evidence

### Accepted input

- `PRODUCT-TASK-SPEC-015-01`: `done`; independent review `PASS`.
- The accepted per-file implementation owner is `DRMCP-WORK-SPEC-001`.
- The accepted cross-file Topics graph implementation owner is `DRMCP-WORK-SPEC-002`.
- W-SPEC-002 consumes the accepted W-SPEC-001 detector result.
- W-SPEC-002 does not own local Topics table shape or H1-adjacent metadata grammar.

### Exact synchronization set

| validation rule | previous owner | synchronized owner | preserved fields |
|---|---|---|---|
| Invalid local `## Topics` table columns, including canonical `file` instead of `ref` | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-001` | Rule text, severity, row order, and table shape. |
| Parent grammar violation | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-001` | Rule text, severity, row order, and table shape. |

No other validation-policy owner changes are in scope.

### W015 wording-only normalization

- `DRMCP-WORK-SPEC-001` is described as the retained per-file parser-aware validation implementation owner.
- `DRMCP-WORK-SPEC-002` is described as the retained cross-file Topics graph-validation implementation owner.
- Both Work Item IDs remain unchanged.
- W015 receives only the T02 relation, Evidence, and these two wording changes.

### No-change set

- `spec:product.design_records.spec_format.follow_up_boundary`: no contradiction found; no edit.
- Unresolved child `ref`: remains owned by `DRMCP-WORK-SPEC-002`.
- Duplicate parent declaration: remains owned by `DRMCP-WORK-SPEC-002`.
- Declaring-parent and child-marker consistency: remains owned by `DRMCP-WORK-SPEC-002`.
- Topic cycle: remains owned by `DRMCP-WORK-SPEC-002`.
- `DRMCP-WORK-SPEC-001`: remains `not_started`.
- `DRMCP-WORK-SPEC-002`: remains `not_started`.
- `DRMCP-TASK-MCP-001-07`: remains `in_progress`.
- `PRODUCT-WORK-SPEC-015`: remains `in_progress`.

### Changed-file manifest

| path | change |
|---|---|
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-02-synchronize-validation-owner-pointers.md` | New T02 Task; closes after validation and independent review PASS. |
| `product/records/spec/design-records/spec-format/validation-policy.md` | Change exactly two implementation-owner cells from W-SPEC-002 to W-SPEC-001. |
| `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md` | Add T02 relation and Evidence; normalize two Impact Scope descriptions. |

Recheck-only records remain unchanged:

- `product/records/spec/design-records/spec-format/follow-up-boundary.md`;
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md`;
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md`;
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md`.

### Validation boundary

Repository-local commands are not available through the filesystem MCP and were not executed by this assistant.

Post-write filesystem reread confirmed the two owner-cell values, the unchanged adjacent graph-owner rows, the W015 relation and wording, and the retained lifecycle states. This is static inspection, not repository-local validation.

The strict spec-format validator applies only to:

- `product/records/spec/design-records/spec-format/validation-policy.md`.

The new Task and updated Work Item are outside the strict spec-format validator scope.
The user executed the initial strict validator, targeted status, and whitespace checks.
The user also executed the post-F-MIN-01-correction and post-F-MIN-02-synchronization whitespace checks.
The final independent review accepted the last pre-closure result as external review evidence.
This closure update changes the T02 Task and W015 bytes, so one final post-closure whitespace check remains external.
Repository-wide clean status is not inferred.

### External verification result

The user executed the strict validator from the repository root.

Result:

- `[strict]  All 1 file(s) OK.`;
- `validator_exit=0`.

The user executed targeted status for the three expected files.

Result:

- `validation-policy.md`: tracked and modified;
- W015: tracked and modified;
- the new T02 Task: untracked;
- no additional path appeared in the targeted result.

The user executed the targeted whitespace checks.

Result:

- `tracked_exit=0`;
- `untracked_exit=1`;
- no whitespace error;
- no exit code `2` or greater.

LF-to-CRLF warnings for all three files were non-blocking.
The untracked exit code `1` is the expected difference result against `NUL`.

### Initial independent review result and correction

Initial verdict: `NEEDS REVISION`.

- Blocking findings: none.
- Major findings: none.
- Minor finding F-MIN-01: W015 copied the current T02 status into Work Item Evidence.
- Advisories: repository-local commands and byte-level Git diff confirmation were not independently run by the reviewer.

F-MIN-01 correction:

- Removed the sentence that copied T02's current status into W015 Evidence.
- Replaced it with pending-evidence wording that does not state a Task lifecycle value.
- Recorded the user-executed validator, targeted status, and whitespace results in T02 and W015 Evidence.
- Preserved the owner-pointer changes, retained-owner wording, changed-file manifest, and lifecycle boundaries.

The correction changed the checked T02 Task and W015 bytes.

The user executed the required post-correction targeted whitespace check.

Result:

- `tracked_exit=0`;
- `untracked_exit=1`;
- no whitespace error;
- no exit code `2` or greater.

The first limited re-review returned `NEEDS REVISION`.

- F-MIN-01: `CLOSED`.
- Blocking findings: none.
- Major findings: none.
- Minor finding F-MIN-02: the post-correction whitespace result had not yet been recorded in T02 and W015 Evidence.
- Regression assessment: `PASS` within the reviewer's static scope.
- Lifecycle assessment: `PASS`.

F-MIN-02 correction:

- Recorded the post-F-MIN-01-correction whitespace result in T02 and W015 Evidence.
- Preserved the initial strict validator and targeted status evidence.
- Preserved the two owner-only changes, retained-owner wording, no-change set, and lifecycle boundary.

This F-MIN-02 evidence synchronization changed the checked T02 Task and W015 bytes.

The user then executed the final pre-review targeted whitespace check.

Result:

- `tracked_exit=0`;
- `untracked_exit=1`;
- no whitespace error;
- no exit code `2` or greater.

The final limited re-review returned `PASS`.

- F-MIN-01: `CLOSED`.
- F-MIN-02: `CLOSED`.
- F-MIN-03: `CLOSED`.
- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- Validation evidence assessment: `PASS`.
- T02 closure readiness: `READY`.
- T03 start readiness: `READY` after T02 closure.

The reviewer accepted the final whitespace result as external review evidence without writing it back into the checked files.
This Task closes on that accepted review result.
One final post-closure targeted whitespace check remains required because this lifecycle and Evidence update changed the checked bytes.

### External commands required

Run from:

`C:\Users\imved\projects\brewprint`

Strict spec-format validation:

```powershell
$specPath = "product/records/spec/design-records/spec-format/validation-policy.md"
python -X utf8 product/src/tools/validate_spec.py --strict --no-color $specPath
$validator_exit = $LASTEXITCODE
"validator_exit=$validator_exit"
```

Expected result:

- `validator_exit=0`;
- `[strict]  All 1 file(s) OK.`

Targeted changed-file status:

```powershell
git status --short -- `
  product/records/tasks/spec/PRODUCT-TASK-SPEC-015-02-synchronize-validation-owner-pointers.md `
  product/records/spec/design-records/spec-format/validation-policy.md `
  product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md
```

Targeted whitespace checks:

```powershell
$trackedPaths = @(
  "product/records/spec/design-records/spec-format/validation-policy.md",
  "product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md"
)
$untrackedPath = "product/records/tasks/spec/PRODUCT-TASK-SPEC-015-02-synchronize-validation-owner-pointers.md"

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

If the new Task is staged, use `git diff --cached --check -- <T02 path>`.
LF-to-CRLF warnings are non-blocking when no whitespace error exists.

### Independent T02 review prompt

```text
C:\Users\imved\projects\brewprint

`PRODUCT-TASK-SPEC-015-02`の独立validation-owner pointer synchronization reviewを行う。

ファイルは変更しないこと。
DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
sandboxへrepositoryを複製しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。
repository-wide clean statusを推測しないこと。
無制限なrepository traversalを行わないこと。

## 最初に読む

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/task-authoring.md`
3. `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
4. `product/records/spec/design-records/authoring-standards/writing-standard.md`
5. `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## T02 changed files

- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-02-synchronize-validation-owner-pointers.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`

## Accepted baseline and recheck-only records

- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-01-inventory-validation-owner-pointers.md`
- `product/records/spec/design-records/spec-format/follow-up-boundary.md`
- `product/records/spec/design-records/spec-format/spec-id-as-ref.md`
- `product/records/spec/design-records/spec-format/topics-table.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md`

## Accepted owner boundary

- W-SPEC-001 owns per-file H1, metadata, spec-kind, `contract_class`, required-section, local Topics shape, path-derived ID, and front-matter checks.
- W-SPEC-002 owns Topics edge extraction, exact child lookup, declaring-parent and child-marker consistency, duplicate authoritative-parent detection, and cycle detection.
- W-SPEC-002 consumes the accepted W-SPEC-001 detector result and does not own local Topics shape or H1-adjacent metadata grammar.

## Review scope

Confirm all of the following without redesigning validation rules or owner boundaries:

- validation-policy has exactly two owner-only changes;
- the local Topics columns row changes from W-SPEC-002 to W-SPEC-001;
- the parent grammar violation row changes from W-SPEC-002 to W-SPEC-001;
- rule text, severity, row order, and table formatting are unchanged;
- unresolved child `ref`, duplicate authoritative parent, parent consistency, and topic cycle remain W-SPEC-002 concerns;
- W015 wording normalization matches the accepted retain disposition;
- both retained Work Item IDs remain unchanged;
- W015 changes are limited to the T02 relation, Evidence, and two Impact Scope descriptions;
- follow-up-boundary remains unchanged;
- W-SPEC-001 and W-SPEC-002 remain `not_started`;
- W015 and hub T07 remain `in_progress` and are not closed;
- changed-file manifest contains only the expected three files;
- strict validator scope includes only validation-policy among the changed files;
- validation and whitespace evidence are accurate and not fabricated;
- Task and Work Item authoring shape is valid;
- no blocking, major, or minor finding remains;
- T02 closure readiness;
- T03 start readiness.

If repository-local commands are available, run the exact strict-validator, targeted status, and whitespace commands recorded in T02 Evidence.

Expected results:

- `validator_exit=0`;
- `[strict]  All 1 file(s) OK.`;
- `tracked_exit=0`;
- `untracked_exit=1`;
- whitespace errorなし;
- exit code 2以上なし.

LF-to-CRLF warning is non-blocking when no whitespace error exists.

Output format:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. Exact owner-pointer change assessment
8. Rule-text and severity preservation assessment
9. W015 wording-normalization assessment
10. Follow-up-boundary no-change assessment
11. Lifecycle assessment
12. Changed-file manifest assessment
13. Validation evidence assessment
14. T02 closure readiness
15. T03 start readiness
```

### Independent limited re-review prompt

```text
C:\Users\imved\projects\brewprint

`PRODUCT-TASK-SPEC-015-02`のF-MIN-01 correction limited re-reviewを行う。

ファイルは変更しないこと。
DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
sandboxへrepositoryを複製しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。
repository-wide clean statusを推測しないこと。
無制限なrepository traversalを行わないこと。

## 最初に読む

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/task-authoring.md`
3. `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
4. `product/records/spec/design-records/authoring-standards/writing-standard.md`
5. `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## Changed files

- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-02-synchronize-validation-owner-pointers.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`

## Previous review

Initial verdict: `NEEDS REVISION`.

- Blocking findings: none.
- Major findings: none.
- F-MIN-01: W015 copied T02's canonical current status into Work Item Evidence.
- Initial user-executed verification returned `validator_exit=0`, `[strict]  All 1 file(s) OK.`, `tracked_exit=0`, and `untracked_exit=1` with no whitespace error or exit code `2` or greater.

## Correction scope

Confirm only the F-MIN-01 correction and resulting Evidence synchronization.
Do not redesign the accepted owner boundary or reopen T01.

Check all of the following:

- W015 no longer states T02's current lifecycle value.
- W015 uses pending-evidence wording instead of copying Task status.
- No other current child-Task status was introduced into W015 by the correction.
- T02 records the initial validator, targeted status, and whitespace results accurately.
- T02 records the initial review verdict and F-MIN-01 accurately.
- The two owner-only changes remain unchanged and correct.
- Rule text, severity, row order, and adjacent W-SPEC-002 owner rows remain unchanged.
- W015 retained-owner wording and both Work Item IDs remain unchanged.
- follow-up-boundary remains unchanged.
- W-SPEC-001 and W-SPEC-002 have not started.
- W015 and hub T07 have not been closed.
- The overall T02 changed-file manifest remains the same three files.
- Post-correction whitespace evidence is present and accurate.
- No blocking, major, or minor finding remains.
- T02 closure readiness.
- T03 start readiness.

If repository-local commands are available, run the targeted whitespace commands recorded in T02 Evidence after the correction.

Expected result:

- `tracked_exit=0`;
- `untracked_exit=1`;
- whitespace errorなし;
- exit code 2以上なし.

LF-to-CRLF warning is non-blocking when no whitespace error exists.

Output format:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. F-MIN-01 correction assessment
8. Evidence synchronization assessment
9. Regression assessment
10. Changed-file manifest assessment
11. Validation evidence assessment
12. Lifecycle assessment
13. T02 closure readiness
14. T03 start readiness
```

### Independent F-MIN-02 limited re-review prompt

```text
C:\Users\imved\projects\brewprint

`PRODUCT-TASK-SPEC-015-02`のF-MIN-02 evidence correction limited re-reviewを行う。

ファイルは変更しないこと。
DRMCPは現在利用できない。
repositoryの読み込みにはfilesystem MCPを使用すること。
sandboxへrepositoryを複製しないこと。
repository-local commandを実行できない場合、実行したと捏造しないこと。
repository-wide clean statusを推測しないこと。
無制限なrepository traversalを行わないこと。

## 最初に読む

1. `prompt_chappy.md`
2. `product/records/spec/design-records/authoring-standards/task-authoring.md`
3. `product/records/spec/design-records/authoring-standards/work-item-authoring.md`
4. `product/records/spec/design-records/authoring-standards/writing-standard.md`
5. `product/records/spec/design-records/authoring-standards/agent-authoring-policy.md`

## Changed files

- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-02-synchronize-validation-owner-pointers.md`
- `product/records/spec/design-records/spec-format/validation-policy.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`

## Previous findings

- F-MIN-01: `CLOSED`. W015 no longer copies T02's current status.
- F-MIN-02: post-F-MIN-01-correction whitespace evidence was not recorded.

User-executed post-F-MIN-01-correction result:

- `tracked_exit=0`;
- `untracked_exit=1`;
- whitespace errorなし;
- exit code 2以上なし.

## Correction scope

Confirm only the F-MIN-02 evidence correction and regression safety.
Do not redesign the accepted owner boundary or reopen T01 or F-MIN-01.

Check all of the following:

- T02 records the post-F-MIN-01-correction whitespace result accurately.
- W015 records the same result accurately without copying a Task lifecycle value.
- F-MIN-02 is closed.
- F-MIN-01 remains closed.
- The initial strict validator result remains valid because validation-policy did not change during finding corrections.
- The two owner-only changes remain unchanged and correct.
- Rule text, severity, row order, and adjacent W-SPEC-002 owner rows remain unchanged.
- W015 retained-owner wording and both Work Item IDs remain unchanged.
- follow-up-boundary remains unchanged.
- W-SPEC-001 and W-SPEC-002 have not started.
- W015 and hub T07 have not been closed.
- The overall T02 changed-file manifest remains the same three files.
- Final pre-review whitespace evidence after the F-MIN-02 synchronization is present and accurate.
- No blocking, major, or minor finding remains.
- T02 closure readiness.
- T03 start readiness.

If repository-local commands are available, run the targeted whitespace commands recorded in T02 Evidence after the F-MIN-02 synchronization.

Expected result:

- `tracked_exit=0`;
- `untracked_exit=1`;
- whitespace errorなし;
- exit code 2以上なし.

LF-to-CRLF warning is non-blocking when no whitespace error exists.

Output format:

1. Verdict: PASS / NEEDS REVISION
2. Previous-finding disposition
3. Blocking findings
4. Major findings
5. Minor findings
6. Advisories
7. F-MIN-02 correction assessment
8. F-MIN-01 regression assessment
9. Evidence synchronization assessment
10. Owner-pointer regression assessment
11. Changed-file manifest assessment
12. Validation evidence assessment
13. Lifecycle assessment
14. T02 closure readiness
15. T03 start readiness
```

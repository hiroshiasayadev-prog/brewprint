# PRODUCT-TASK-SPEC-015-01: Inventory validation owner pointers

- **id**: PRODUCT-TASK-SPEC-015-01
- **status**: done
- **date**: 2026-06-28
- **work_item**: PRODUCT-WORK-SPEC-015
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-007-04
- **outputs**:
  - PRODUCT-WORK-SPEC-015
  - DRMCP-TASK-MCP-001-07

## Goal

Inventory PRODUCT references to the retained DRMCP validation Work Items.

Identify exact owner-pointer changes for T02 without changing PRODUCT rules or severity.

## Work

- Record the exact pre-creation `PRODUCT-TASK-SPEC-015-*` inventory.
- Confirm the accepted W007 handoff and retained owner set.
- Search only `product/records/` for `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002`.
- Read only files containing scoped matches and required recheck-only authorities.
- Classify each pointer as correct per-file ownership, correct graph ownership, known mismatch, additional stale pointer, wording-only ambiguity, or no-change.
- Record exact paths, context, current owner, accepted owner, and T02 disposition.
- Preserve PRODUCT rule text and severity.
- Record changed-file and recheck-only manifests.
- Record targeted external status and whitespace commands.
- Record a ready-to-run independent T01 review prompt.
- Synchronize the hub T07 lifecycle only when its recorded gate is satisfied.

This Task does not:

- edit PRODUCT validation-policy or follow-up-boundary specs;
- change validation rule text or severity;
- start W-SPEC-001 or W-SPEC-002;
- implement parser, validator, graph, fixture, or test behavior;
- reopen W007, T04, or W003 through W006;
- close `PRODUCT-WORK-SPEC-015` or hub `DRMCP-TASK-MCP-001-07`.

## Done condition

- The exact pre-creation Task inventory is recorded.
- The accepted durable owner set is recorded without re-deciding W007.
- All scoped PRODUCT matches are classified.
- Every actionable mismatch has an exact T02 disposition.
- PRODUCT rule text and severity remain unchanged.
- The changed-file and recheck-only manifests match the actual T01 writes.
- Targeted external status and whitespace evidence is available.
- Independent review reports no unresolved blocking, major, or minor finding.
- T01 changes to `done` only after verification, review, finding correction, and closure synchronization.

## Verification

- Confirm the accepted per-file owner is `DRMCP-WORK-SPEC-001`.
- Confirm the accepted cross-file Topics graph owner is `DRMCP-WORK-SPEC-002`.
- Confirm W-SPEC-002 consumes the accepted W-SPEC-001 detector result.
- Confirm W-SPEC-002 does not own local Topics presence, header, column, or row-shape detection.
- Confirm every scoped PRODUCT match is represented in the inventory.
- Confirm T02 changes only owner pointers identified as actionable.
- Confirm `PRODUCT-REQ-SPEC-001` already lists `PRODUCT-WORK-SPEC-015`.
- Confirm W007 is `done`, W015 is `in_progress`, and hub T07 is `in_progress` but not `done`.
- Run the targeted status and whitespace commands recorded in Evidence.
- Run the independent T01 review prompt recorded in Evidence.

Task and Work Item records are outside the strict spec-format validator scope.
No repository-local command result may be recorded unless externally executed.
Repository-wide clean status must not be inferred.

## Evidence

### Pre-creation inventory

The exact directory `product/records/tasks/spec/` was listed before T01 creation.

No `PRODUCT-TASK-SPEC-015-*` record existed.
The canonical new ID is `PRODUCT-TASK-SPEC-015-01`.
No repository-wide traversal was used.

### Accepted upstream state

- `DRMCP-WORK-MCP-007`: `done`.
- `DRMCP-TASK-MCP-007-04`: `done`.
- F-MAJ-01: `CLOSED`.
- `DRMCP-WORK-SPEC-001`: retained and `not_started`.
- `DRMCP-WORK-SPEC-002`: retained and `not_started`.
- `PRODUCT-WORK-SPEC-015` was `not_started` before this Task opened.
- `DRMCP-TASK-MCP-001-07` was `not_started` before this Task opened.

Accepted durable owner set:

| validation area | accepted owner |
|---|---|
| Per-file H1, metadata, spec-kind, `contract_class`, required-section, local Topics shape, path-derived ID, and front-matter checks | `DRMCP-WORK-SPEC-001` |
| Topics edge extraction, exact child lookup, parent consistency, duplicate authoritative-parent detection, and cycle detection | `DRMCP-WORK-SPEC-002` |

W-SPEC-002 consumes the accepted W-SPEC-001 detector result.
W-SPEC-002 does not own local Topics presence, header, column, or row-shape detection.

### Scoped search evidence

Two fixed-string grep calls were limited to `product/records/` and Markdown files.

| pattern | files inspected | files with matches | matches | truncated |
|---|---:|---:|---:|---|
| `DRMCP-WORK-SPEC-001` | 157 | 16 | 49 | no |
| `DRMCP-WORK-SPEC-002` | 157 | 13 | 24 | no |

Only matching files and explicit recheck-only authorities were read.

### Pointer inventory

| exact path | rule or context | current owner or reference | accepted owner | classification | T02 disposition |
|---|---|---|---|---|---|
| `product/records/spec/design-records/spec-format/validation-policy.md` | Missing or multiple real H1, H1 shape, H1-adjacent metadata, `contract_class`, path-derived ID, front matter, `## What this is`, spec kind, and required-section rows | `DRMCP-WORK-SPEC-001` | `DRMCP-WORK-SPEC-001` | Correct per-file detector | No change. |
| `product/records/spec/design-records/spec-format/validation-policy.md` | Invalid local `## Topics` columns, including canonical `file` instead of `ref` | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-001` | Known mismatch | Change owner only to W-SPEC-001. Preserve rule text and severity. |
| `product/records/spec/design-records/spec-format/validation-policy.md` | Unresolved child `ref` | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-002` | Correct cross-file graph validation | No change. |
| `product/records/spec/design-records/spec-format/validation-policy.md` | Duplicate parent declaration | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-002` | Correct cross-file graph validation | No change. |
| `product/records/spec/design-records/spec-format/validation-policy.md` | Parent grammar violation | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-001` | Additional stale pointer | Change owner only to W-SPEC-001. Preserve rule text and severity. |
| `product/records/spec/design-records/spec-format/validation-policy.md` | Topic cycle | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-002` | Correct cross-file graph validation | No change. |
| `product/records/spec/design-records/spec-format/follow-up-boundary.md` | Parser-aware per-file implementation follow-up | `DRMCP-WORK-SPEC-001` | `DRMCP-WORK-SPEC-001` | Correct per-file owner | No change. |
| `product/records/spec/design-records/spec-format/follow-up-boundary.md` | Index Topics graph implementation follow-up | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-002` | Correct cross-file graph owner | No change. |
| `product/records/spec/design-records/spec-format/follow-up-boundary.md` | Current DRMCP implementation boundary names both durable follow-ups | W-SPEC-001 / W-SPEC-002 | W-SPEC-001 / W-SPEC-002 | No-change | No change. The wording remains compatible with both retained Work Items being future implementation owners. |
| `product/records/spec/brewprint/namespaces/domain-catalog.md` | Current DRMCP `SPEC` domain ID examples | W-SPEC-001 and W-SPEC-002 | Not an owner assignment | No-change | No change. This is namespace inventory evidence. |
| `product/records/investigations/spec/PRODUCT-INV-SPEC-001-spec-format-topic-tree-design-and-migration-feasibility.md` | Historical follow-up candidates and combined implementation recommendation | W-SPEC-001 / W-SPEC-002 | Current granular split is W-SPEC-001 per-file and W-SPEC-002 graph | Wording-only ambiguity | No change. The concluded investigation records historical recommendation context, not current normative ownership. |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-001-02-review-spec-format-contract.md` | Historical follow-up split evidence | W-SPEC-001 / W-SPEC-002 | Same retained identities | No-change | No change. |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-001-03-create-follow-up-split.md` | W-SPEC-001 parser-aware follow-up and W-SPEC-002 graph follow-up | W-SPEC-001 / W-SPEC-002 | Same retained identities and split | No-change | No change. |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-001-04-close-work-spec-001.md` | Historical closure evidence for implementation-phase isolation | W-SPEC-001 / W-SPEC-002 | Same retained identities | No-change | No change. |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-03-split-namespace-profile-and-compatibility.md` | Active `DRMCP` / `SPEC` domain ID examples | W-SPEC-001 and W-SPEC-002 | Not an owner assignment | No-change | No change. |
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-012-08-apply-drmcp-app-local-handoff.md` | Parser and YAML-front-matter compatibility gaps | `DRMCP-WORK-SPEC-001` | `DRMCP-WORK-SPEC-001` | Correct per-file detector | No change. |
| `product/records/work-items/spec/PRODUCT-WORK-SPEC-001-spec-format-contract-and-follow-up-split.md` | Historical implementation-phase follow-up and combined W-SPEC-001/002 wording | W-SPEC-001 / W-SPEC-002 | Current granular split is W-SPEC-001 per-file and W-SPEC-002 graph | Wording-only ambiguity | No change. The `done` Work Item records historical planning and does not act as current rule-level authority. |
| `product/records/work-items/spec/PRODUCT-WORK-SPEC-002-path-derived-canonical-spec-refs-and-ref-first-topic-index.md` | DRMCP validation implementation non-scope | W-SPEC-001 / W-SPEC-002 | Same durable owner set | No-change | No change. |
| `product/records/work-items/spec/PRODUCT-WORK-SPEC-003-spec-authoring-guide-format-update.md` | DRMCP validation implementation non-scope | W-SPEC-001 / W-SPEC-002 | Same durable owner set | No-change | No change. |
| `product/records/work-items/spec/PRODUCT-WORK-SPEC-004-ownership-boundary-decision-and-relocation-plan.md` | DRMCP validation implementation non-scope | W-SPEC-001 / W-SPEC-002 | Same durable owner set | No-change | No change. |
| `product/records/work-items/spec/PRODUCT-WORK-SPEC-005-existing-spec-format-migration-and-restructuring.md` | DRMCP implementation non-scope | W-SPEC-001 / W-SPEC-002 | Same durable owner set | No-change | No change. |
| `product/records/work-items/spec/PRODUCT-WORK-SPEC-006-temporary-spec-format-validator-tooling.md` | DRMCP production implementation non-scope | W-SPEC-001 / W-SPEC-002 | Same durable owner set | No-change | No change. |
| `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md` | Impact Scope still calls W-SPEC-001 and W-SPEC-002 pointer candidates after the retain disposition was accepted | W-SPEC-001 / W-SPEC-002 | Same retained durable owner set | Wording-only ambiguity | In T02, replace candidate wording with retained-owner wording. Do not change either ID. |

### T02 synchronization set

T02 has exactly two owner-only changes in `validation-policy.md`:

| validation rule | current owner | accepted owner |
|---|---|---|
| Invalid local `## Topics` table columns, including canonical `file` instead of `ref` | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-001` |
| Parent grammar violation | `DRMCP-WORK-SPEC-002` | `DRMCP-WORK-SPEC-001` |

Reason for the additional stale pointer:

- `spec:product.design_records.spec_format.spec_id_as_ref` defines the H1-adjacent `parent` grammar.
- W-SPEC-001 owns H1-adjacent metadata shape detection.
- W-SPEC-002 explicitly excludes H1-adjacent metadata detection.
- W-SPEC-002 still owns cross-file parent consistency and duplicate authoritative-parent detection.

T02 also has one wording-only synchronization in W015 Impact Scope:

- replace `Existing parser-aware validation pointer candidate` with retained per-file implementation-owner wording;
- replace `Existing Topics graph-validation pointer candidate` with retained graph implementation-owner wording;
- keep both Work Item IDs unchanged.

No additional current PRODUCT owner pointer requires synchronization.
No follow-up-boundary edit is required by this inventory.

### Hub T07 lifecycle assessment

The existing hub gate says T07 remains `not_started` until both child Work Items begin.

Current child state after this T01 opening:

- `DRMCP-WORK-MCP-007`: `done`; it has begun and completed.
- `PRODUCT-WORK-SPEC-015`: `in_progress`; it has begun.

The gate is satisfied.
Hub `DRMCP-TASK-MCP-001-07` changes to `in_progress` and remains not done.

### Changed-file manifest

| path | change |
|---|---|
| `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-01-inventory-validation-owner-pointers.md` | New T01 Task; records inventory, manifests, external commands, independent review, and final `done` closure. |
| `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md` | Change status to `in_progress`, add T01 to `tasks`, and record T01 opening evidence. |
| `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md` | Change hub status to `in_progress` and record the satisfied lifecycle gate. |

### Recheck-only manifest

| record or authority | result |
|---|---|
| `product/records/spec/design-records/spec-format/validation-policy.md` | Read only. Two owner-only T02 changes identified. No T01 edit. |
| `product/records/spec/design-records/spec-format/follow-up-boundary.md` | Durable owners and wording accepted. No edit. |
| `product/records/requirements/PRODUCT-REQ-SPEC-001-mcp-readable-spec-format-and-topic-tree.md` | Already lists W015 in `work_items`. No edit. |
| `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md` | Accepted upstream Work Item is `done`. No edit. |
| `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-04-record-product-handoff-and-close-validation-disposition.md` | Accepted upstream Task is `done`; F-MAJ-01 is closed. No edit. |
| `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md` | Retained per-file owner remains `not_started`. No edit. |
| `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md` | Retained graph owner remains `not_started`. No edit. |

### Validation boundary

Repository-local commands are not available through the filesystem MCP and were not executed by this assistant.

The user executed the targeted three-path status command from the repository root.
The result matched the changed-file manifest:

- `DRMCP-TASK-MCP-001-07` was tracked and modified;
- `PRODUCT-WORK-SPEC-015` was tracked and modified;
- the new T01 Task was untracked;
- no additional path appeared in the targeted result.

The user executed the targeted whitespace commands.
External results were:

- `tracked_exit=0`;
- `untracked_exit=1`;
- no whitespace error;
- no exit code `2` or greater.

The untracked exit code `1` is the expected difference result against `NUL`.
LF-to-CRLF working-copy warnings for all three files are non-blocking because no whitespace error was reported.

Task and Work Item records are outside the strict spec-format validator scope.
No strict spec-format validation applies to the T01 changed-file manifest.
Repository-wide clean status is not inferred.

The final pre-review targeted whitespace check also returned:

- `tracked_exit=0`;
- `untracked_exit=1`;
- no whitespace error;
- no exit code `2` or greater.

The independent review returned `PASS` with no blocking, major, or minor finding.

Review advisories were non-blocking:

- repository-local Git commands were not independently rerun because the reviewer had no command-execution tool;
- the pre-creation inventory could not be independently reconstructed from the current filesystem state, but no contradiction was found.

The review confirmed inventory completeness, both owner-only T02 changes, the W015 wording-only normalization, hub lifecycle state, changed-file manifest, and authoring shape.
T01 is ready and is closed as `done`.

This closure update changes the checked T01 and W015 bytes.
One final post-closure targeted whitespace check is required.

### External commands required

Run from:

`C:\Users\imved\projects\brewprint`

```powershell
git status --short -- `
  product/records/tasks/spec/PRODUCT-TASK-SPEC-015-01-inventory-validation-owner-pointers.md `
  product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md
```

For the new untracked T01 Task and the two tracked records:

```powershell
$trackedPaths = @(
  "product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md",
  "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md"
)
$untrackedPath = "product/records/tasks/spec/PRODUCT-TASK-SPEC-015-01-inventory-validation-owner-pointers.md"

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

If T01 is staged, use `git diff --cached --check -- <T01 path>`.
LF-to-CRLF warnings are non-blocking when no whitespace error exists.

### Independent T01 review prompt

```text
C:\Users\imved\projects\brewprint

`PRODUCT-TASK-SPEC-015-01`の独立validation-owner pointer inventory reviewを行う。

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

## T01 changed files

- `product/records/tasks/spec/PRODUCT-TASK-SPEC-015-01-inventory-validation-owner-pointers.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-015-synchronize-validation-owner-pointers.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md`

## Accepted upstream state

- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-007-validation-work-item-disposition-and-rebaseline.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-007-04-record-product-handoff-and-close-validation-disposition.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-001-parser-aware-spec-format-validation.md`
- `drmcp/records/work-items/spec/DRMCP-WORK-SPEC-002-index-topics-graph-validation.md`

Accepted owner set:

- W-SPEC-001 owns per-file H1, metadata, spec-kind, contract-class, required-section, local Topics shape, path-derived ID, and front-matter checks.
- W-SPEC-002 owns Topics edge extraction, exact child lookup, parent consistency, duplicate authoritative-parent detection, and cycle detection.
- W-SPEC-002 consumes accepted W-SPEC-001 detector results.
- W-SPEC-002 does not own local Topics presence, header, column, or row-shape detection.

## PRODUCT authorities and recheck-only records

- `product/records/spec/design-records/spec-format/validation-policy.md`
- `product/records/spec/design-records/spec-format/follow-up-boundary.md`
- `product/records/spec/design-records/spec-format/spec-id-as-ref.md`
- `product/records/spec/design-records/spec-format/topics-table.md`
- `product/records/spec/design-records/spec-format/document-shape.md`
- `product/records/requirements/PRODUCT-REQ-SPEC-001-mcp-readable-spec-format-and-topic-tree.md`

## Review scope

Confirm the T01 inventory is complete and correctly classified without redesigning W007 or W-SPEC-001/002.

Check all of the following:

- The exact Task directory inventory found no prior `PRODUCT-TASK-SPEC-015-*` file.
- T01 is `in_progress`, not `done`.
- W015 is `in_progress` and lists T01.
- W007 and T04 remain accepted and `done`.
- W-SPEC-001 and W-SPEC-002 remain `not_started`.
- The scoped PRODUCT grep method is limited to `product/records/` and the two exact Work Item IDs.
- Every matching PRODUCT file is represented in the inventory or explicitly grouped without loss.
- The known local Topics column-shape row is classified as W-SPEC-001 ownership.
- The `parent grammar violation` row is correctly classified as local H1-adjacent metadata grammar owned by W-SPEC-001.
- Unresolved child ref, duplicate authoritative-parent, parent consistency, and topic cycle remain W-SPEC-002 concerns.
- PRODUCT rule text and severity are unchanged.
- Follow-up-boundary requires no T01 or T02 edit unless the review identifies a concrete stale owner statement.
- Historical Task, Work Item, Investigation, and namespace-ID references are not incorrectly treated as current rule-level owner pointers.
- T02 synchronization set contains exactly two owner-only changes and one W015 wording-only normalization.
- PRODUCT-REQ-SPEC-001 already lists W015 and remains unchanged.
- Hub T07 gate is applied correctly: W007 is done, W015 is in progress, T07 is in progress and not done.
- Changed-file and recheck-only manifests match the actual writes.
- Repository-local results are not fabricated.
- Task and Work Item authoring shape is valid.

If repository-local commands are available, run the exact three-path status and whitespace commands recorded in T01.

Expected whitespace result:

- `tracked_exit=0`
- `untracked_exit=1`
- whitespace errorなし
- exit code 2以上なし

LF-to-CRLF warning is non-blocking when no whitespace error exists.

Output format:

1. Verdict: PASS / NEEDS REVISION
2. Blocking findings
3. Major findings
4. Minor findings
5. Advisories
6. Scoped inventory completeness assessment
7. Known Topics-column mismatch assessment
8. Parent-grammar pointer assessment
9. Historical-reference disposition assessment
10. T02 synchronization-set assessment
11. Hub T07 lifecycle assessment
12. Changed-file manifest assessment
13. Validation evidence assessment
14. T01 closure readiness
15. T02 start readiness
```

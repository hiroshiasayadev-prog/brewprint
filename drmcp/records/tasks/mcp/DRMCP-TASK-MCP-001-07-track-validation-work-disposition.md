# DRMCP-TASK-MCP-001-07: Track validation-work disposition

- **id**: DRMCP-TASK-MCP-001-07
- **status**: done
- **date**: 2026-06-26
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d coordination
- **depends_on**:
  - DRMCP-TASK-MCP-001-01
- **outputs**:
  - DRMCP-WORK-MCP-007
  - PRODUCT-WORK-SPEC-015

## Goal

Accept the coordinated disposition of existing DRMCP validation Work Items and PRODUCT owner pointers.

## Work

- Track `DRMCP-WORK-MCP-007` and `PRODUCT-WORK-SPEC-015` as the exact owner-local Work Items selected by T01.
- Delegate the disposition of `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002` to the selected owner.
- Delegate matching PRODUCT validation-policy owner-pointer changes to the correctly owned PRODUCT Work Item.
- Require one explicit disposition for each existing Work Item: retain, supersede, absorb, or close.
- Track all required child Work Items through review and `done`.
- Record exact Work Item IDs and accepted evidence here.

This Task does not decide validation semantics or edit owner pointers directly.
All detailed decisions and updates belong to the selected Work Items.

## Done condition

- Every affected existing Work Item has an explicit accepted disposition.
- PRODUCT validation-policy owner pointers match the accepted disposition.
- No replacement Work Item creates duplicated authority.
- Every selected child Work Item is `done`.
- Cross-owner review has no blocking or major findings.
- Exact Work Item IDs and evidence pointers are recorded here.

## Verification

- Compare the final dispositions with `spec:product.design_records.spec_format.validation_policy`.
- Confirm reciprocal references and source Requirements for all created Work Items.
- Confirm that this Task contains no direct validation-policy implementation evidence.

## Evidence

### Accepted disposition and closure pointers

| record | accepted state | evidence pointer |
|---|---|---|
| `DRMCP-WORK-SPEC-001` | `retain`; `not_started` | `DRMCP-TASK-MCP-007-02` accepted the disposition. `DRMCP-TASK-MCP-007-03` rebaselined the retained owner. |
| `DRMCP-WORK-SPEC-002` | `retain`; `not_started` | `DRMCP-TASK-MCP-007-02` accepted the disposition. `DRMCP-TASK-MCP-007-03` rebaselined the retained owner. |
| `DRMCP-WORK-MCP-007` | `done` | `DRMCP-TASK-MCP-007-04` recorded the PRODUCT handoff and closed the DRMCP owner-side Work Item. |
| `PRODUCT-WORK-SPEC-015` | `done` | `PRODUCT-TASK-SPEC-015-03` validated the synchronized pointers, recorded cross-owner review `PASS`, and closed the PRODUCT owner-side Work Item. |

No replacement Work Item was created for either retained validation owner.
The durable implementation-owner set remains exactly `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002`.

### Final Done condition assessment

| Done condition | assessment | evidence |
|---|---|---|
| Every affected existing Work Item has an explicit accepted disposition. | Satisfied. | W-SPEC-001 and W-SPEC-002 are both accepted as `retain`. |
| PRODUCT validation-policy owner pointers match the accepted disposition. | Satisfied. | Local Topics column-shape and parent grammar rows point to W-SPEC-001. Cross-file Topics graph rows remain with W-SPEC-002. |
| No replacement Work Item creates duplicated authority. | Satisfied. | No replacement identity exists. The retained boundaries remain separate and non-overlapping. |
| Every selected child Work Item is `done`. | Satisfied. | W007 and W015 are both `done`. |
| Cross-owner review has no blocking or major findings. | Satisfied. | Final verdict `PASS`; blocking, major, and minor findings are none. |
| Exact Work Item IDs and evidence pointers are recorded here. | Satisfied. | The accepted disposition and closure table records both retained owners, both child Work Items, and their closure Tasks. |

### Independent hub closure review

- Verdict: `PASS`.
- Blocking findings: none.
- Major findings: none.
- Minor findings: none.
- Child Work Item states: `DRMCP-WORK-MCP-007` is `done`; `PRODUCT-WORK-SPEC-015` is `done`.
- Owner-pointer consistency: `PASS`.
- Duplicated-authority assessment: `PASS`.
- W-SPEC-001 lifecycle: retained and `not_started`.
- W-SPEC-002 lifecycle: retained and `not_started`.
- PRODUCT rule text, severity, row order, and table shape remain owned by PRODUCT and were not moved into this Task.
- This Task records lifecycle and evidence pointers only. It does not own validation-policy implementation evidence.

### Parent Work Item assessment

`DRMCP-WORK-MCP-001` lists this Task and both owner-side Work Items.
No stale relation, source-Requirement mismatch, or lifecycle contradiction was found.
The parent remains unchanged because Task status is canonical here and must not be copied into the Work Item.

### Validation boundary

Review used bounded static inspection through the filesystem MCP.
No repository-local Git, Python, PowerShell, validator, formatter, or test command was executed by this assistant.
Repository-wide clean status remains unknown.
Byte-level Git diff comparison was not performed.

The strict spec-format validator does not apply to this Task record.
The following target-scoped structural check remains external.

Run from `C:\Users\imved\projects\brewprint`:

```powershell
@'
from pathlib import Path

path = Path(r"drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md")
text = path.read_text(encoding="utf-8")

outside_fences = []
in_fence = False
for line in text.splitlines():
    if line.startswith("```"):
        in_fence = not in_fence
        continue
    if not in_fence:
        outside_fences.append(line)

outside_text = "\n".join(outside_fences)

assert sum(line == "# DRMCP-TASK-MCP-001-07: Track validation-work disposition" for line in outside_fences) == 1
assert "- **status**: done" in outside_fences

for heading in (
    "## Goal",
    "## Work",
    "## Done condition",
    "## Verification",
    "## Evidence",
):
    assert sum(line == heading for line in outside_fences) == 1, heading

for required_ref in (
    "DRMCP-WORK-MCP-007",
    "PRODUCT-WORK-SPEC-015",
    "DRMCP-WORK-SPEC-001",
    "DRMCP-WORK-SPEC-002",
    "DRMCP-TASK-MCP-007-04",
    "PRODUCT-TASK-SPEC-015-03",
):
    assert required_ref in outside_text, required_ref

assert all(line.strip() != "TBD" for line in outside_fences)
print("task_shape=OK")
'@ | python -X utf8 -

$validator_exit = $LASTEXITCODE
"validator_exit=$validator_exit"
```

Expected result:

- `task_shape=OK`;
- `validator_exit=0`.

Targeted status command:

```powershell
git status --short -- `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md
```

Expected result:

- only the target path appears in this targeted result when it is modified;
- unrelated working-tree state is outside this check.

Targeted whitespace command:

```powershell
git diff --check -- `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md

"tracked_exit=$LASTEXITCODE"
```

Expected result:

- `tracked_exit=0`;
- no whitespace error.

This closure update changes the target file bytes.
A final post-closure targeted whitespace check remains external and must not be written back into this Task.

### Final changed-file manifest

| path | change |
|---|---|
| `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-07-track-validation-work-disposition.md` | Record independent hub closure review, external verification commands, residual limitations, and final lifecycle state. |

T07 changes to `done` on 2026-06-28 after every Done condition is satisfied.

# DRMCP-TASK-MCP-011-06: Independently re-review rebuild read-runtime architecture findings

- **id**: DRMCP-TASK-MCP-011-06
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-011
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-011-05
- **outputs**:
  - DRMCP-TASK-MCP-011-06

## Goal

Independently determine whether F-MAJ-01 and F-MAJ-02 are correctly resolved in canonical artifacts.

Check direct cross-artifact regressions caused by the correction.

Do not repeat the full W011 design review.

Do not perform correction or lifecycle synchronization.

## Work

### Responsibility and execution boundary

This Task is a read-only independent finding-closure re-review.

The reviewer must:

- judge F-MAJ-01 and F-MAJ-02 separately;
- check only direct regressions caused by the T05 correction;
- read the named authority and evidence directly;
- preserve accepted decisions outside the correction boundary;
- stop before correction, lifecycle synchronization, design closure, implementation planning, or production implementation.

The reviewer must not:

- modify files;
- fix findings;
- close findings based only on T05 claims;
- repeat the full W011 design review;
- review W003 through W010 as a new design scope;
- change Task, Work Item, ADR, Requirement, or Specification lifecycle state.

### Review authority

Read:

```text
drmcp/records/tasks/mcp/
DRMCP-TASK-MCP-011-04-independently-review-rebuild-read-runtime-architecture.md
DRMCP-TASK-MCP-011-05-correct-rebuild-read-runtime-architecture-review-findings.md

drmcp/records/work-items/mcp/
DRMCP-WORK-MCP-011-rebuild-read-runtime-application-architecture.md
```

Treat T04 as the authority for the original findings.

Treat T05 as correction-authoring evidence, not finding-closure evidence.

Carry these finding scopes forward exactly:

| finding | affected decisions | new user decision required |
|---|---|---|
| F-MAJ-01 | D-001, D-002, D-007, D-009 | no |
| F-MAJ-02 | D-006 | no |

### Governing Requirements

Read:

```text
drmcp/records/requirements/mcp/
DRMCP-REQ-MCP-001-multi-root-multi-namespace-mcp-tool-contract.md
DRMCP-REQ-MCP-002-namespace-aware-authoring-transaction-conformance.md
```

Read `DRMCP-REQ-MCP-002` only to confirm deferred ownership of authoring integration.

Do not review the complete `DRMCP-REQ-MCP-002` design.

### Corrected decision ledger

Read:

```text
drmcp/records/tasks/mcp/
DRMCP-TASK-MCP-011-01-run-rebuild-read-runtime-architecture-decision-loop.md
```

Primary decision review set:

```text
D-001
D-002
D-006
D-007
D-009
```

Review D-003, D-004, D-005, and D-008 only for direct regression caused by T05.

### Corrected ADRs

Read:

```text
drmcp/records/adr/mcp/
DRMCP-ADR-MCP-002-request-scoped-read-runtime-snapshot-and-lifecycle.md
DRMCP-ADR-MCP-003-layered-read-runtime-application-and-adapter-boundaries.md
DRMCP-ADR-MCP-005-validation-orchestration-over-fresh-snapshots.md
DRMCP-ADR-MCP-006-read-runtime-go-package-boundaries.md
```

### Corrected Specifications

Read:

```text
drmcp/records/spec/implementation/index.md

drmcp/records/spec/design-records-mcp/
overview.md
namespace-scanning.md
responsibility-boundary.md

drmcp/records/spec/design-records-mcp/tools/
overview.md
authoring-transaction-model.md
accept-proposed-write.md
```

### Unchanged authoring contracts for consistency checking

Read:

```text
drmcp/records/spec/design-records-mcp/schema/
authoring-transaction-schema.md
record-source.md
metadata-grammar.md

drmcp/records/spec/design-records-mcp/tools/
propose-record-update.md
```

Use these four files only for direct F-MAJ-02 consistency checking.

Do not review the full existing authoring contract.

Do not perform current-format correction.

### Review boundary

Judge only:

```text
1. F-MAJ-01 closure
2. F-MAJ-02 closure
3. Direct regression caused by the correction
```

Do not repeat the full design review.

Do not reopen D-003, D-004, D-005, or D-008 unless the correction directly changed their accepted meaning.

Do not review W003 through W010 as a new design scope.

Report a new finding only when T05 caused the defect or directly exposed a material defect.

### F-MAJ-01 closure criteria

Mark F-MAJ-01 `CLOSED` only when every criterion below passes.

#### Exact operation set

The W011 architecture covers only:

```text
list_records
get_records
resolve_reference
validate_records
```

The W011 architecture does not cover:

```text
list_authoring_guides
get_authoring_guidance
propose_record_create
propose_record_update
get_proposed_write
accept_proposed_write
discard_proposed_write
```

The public catalog must still contain the authoring-guidance and authoring-transaction tools.

#### Cross-artifact consistency

The following artifacts must express the same four-operation boundary:

```text
W011 Boundary and Completion Condition
T01 D-001
T01 D-002
T01 D-007
T01 D-009
ADR-002
ADR-003
ADR-006
spec:drmcp.implementation
Design Records MCP overview
namespace scanning
responsibility boundary
tools overview
```

#### No remaining all-tool ambiguity

The following forms are defects only when their context applies the W011 architecture to all tools:

```text
each MCP invocation
each MCP tool invocation
each public MCP tool
each public MCP tool dispatch
each public operation
```

Do not report a finding merely because a phrase exists in an unrelated or general context.

#### Package inventory

Confirm all conditions:

- The package tree is complete for the four W011 read-runtime operations.
- The package tree is not the complete Design Records MCP server inventory.
- Authoring-guidance and authoring-transaction package placement is outside W011.
- An implementation Task author does not need to decide whether authoring packages belong in the W011 package tree.

### F-MAJ-02 closure criteria

Mark F-MAJ-02 `CLOSED` only when every criterion below passes.

#### Current W011 validation authority

The current normative W011 runtime is standalone `validate_records`.

Confirm:

- each request uses a fresh request-scoped snapshot;
- validation passes begin after source loading completes;
- per-source validation precedes relation and Topics graph validation;
- validators perform no filesystem I/O;
- validators do not call public MCP tools;
- findings aggregation is outside validators;
- findings ordering is outside validators;
- semantic duplicate suppression is outside validators;
- MCP projection is outside validators.

#### Persisted-write constraint

D-006 must retain both constraints:

- a future persisted-write validation caller does not reuse candidate or pre-write state;
- the caller rebuilds validation input from persisted filesystem state.

Do not accept a correction that deletes or weakens this constraint.

Do not accept replacement with proposal candidate validation.

#### Deferred authoring integration

Confirm:

- current authoring-transaction integration is deferred to `DRMCP-REQ-MCP-002`;
- W011 does not own current authoring-transaction use cases;
- W011 does not own authoring-transaction snapshot lifecycle;
- W011 does not own authoring-transaction source loading;
- W011 does not own authoring-transaction package architecture;
- YAML and V01-SPEC authoring semantics are not treated as integrated with current-format validation semantics.

#### Authoring contract correction

Confirm these Specifications do not require current normative integration with the W011 fresh-snapshot architecture:

```text
tools/authoring-transaction-model.md
tools/accept-proposed-write.md
```

Confirm the correction did not remove or change these existing authoring semantics:

- proposal lifecycle;
- accept-time eligibility;
- pre-write validation;
- `written`;
- `files_written`;
- actual-write reporting;
- repair guidance;
- no automatic rollback;
- candidate and affected-record-set validation boundary.

Deferral to `DRMCP-REQ-MCP-002` must not delete existing authoring interface behavior.

### Direct regression criteria

Check at minimum:

- D-003, D-004, D-005, and D-008 retain their accepted meaning.
- ADR `status`, `date`, `depends_on`, `supersedes`, and `migrated_to_spec` metadata did not change unintentionally.
- Public read request behavior did not change.
- Public read response behavior did not change.
- Public status and warning behavior did not change.
- Public diagnostic and error behavior did not change.
- Current and legacy index separation remains intact.
- `validate_records` subject remains intact.
- `validate_records` ordering remains intact.
- `validate_records` diagnostic boundary remains intact.
- `validate_records` execution-error boundary remains intact.
- Authoring-guidance tools remain in the public catalog.
- Authoring-transaction tools remain in the public catalog.
- The correction does not decide `DRMCP-REQ-MCP-002` current-format design.
- T05 does not claim independent finding closure.
- W011 does not claim design closure.
- W011 does not claim implementation readiness.

### Git review boundary

Inspect only these 14 T05-changed files:

```text
drmcp/records/tasks/mcp/
DRMCP-TASK-MCP-011-01-run-rebuild-read-runtime-architecture-decision-loop.md
DRMCP-TASK-MCP-011-05-correct-rebuild-read-runtime-architecture-review-findings.md

drmcp/records/work-items/mcp/
DRMCP-WORK-MCP-011-rebuild-read-runtime-application-architecture.md

drmcp/records/adr/mcp/
DRMCP-ADR-MCP-002-request-scoped-read-runtime-snapshot-and-lifecycle.md
DRMCP-ADR-MCP-003-layered-read-runtime-application-and-adapter-boundaries.md
DRMCP-ADR-MCP-005-validation-orchestration-over-fresh-snapshots.md
DRMCP-ADR-MCP-006-read-runtime-go-package-boundaries.md

drmcp/records/spec/implementation/index.md

drmcp/records/spec/design-records-mcp/
overview.md
namespace-scanning.md
responsibility-boundary.md

drmcp/records/spec/design-records-mcp/tools/
overview.md
authoring-transaction-model.md
accept-proposed-write.md
```

Use these tools when available:

```text
git.inspect_worktree
git.inspect_diff
```

Use:

```text
cwd: C:\Users\imved\projects\brewprint
include_untracked: true
```

Split inspection when one request exceeds the tool limit:

```text
1. T01, T05, W011, ADR-002, ADR-003, ADR-005, ADR-006
2. implementation Specification and the six Design Records MCP Specifications
```

Inspect the complete textual patch.

Do not interpret `result: pass` as a clean working tree.

Do not check or infer repository-wide cleanliness.

Treat LF-to-CRLF warnings separately from whitespace findings.

### Verdict and disposition

For each original finding, use one disposition:

```text
CLOSED
OPEN
```

Use one overall verdict:

| verdict | condition |
|---|---|
| `PASS` | Both findings are `CLOSED`, with no blocking, major, or required minor direct regression. |
| `NEEDS REVISION` | Either finding is `OPEN`, or the correction directly introduced a material regression. |
| `NOT READY` | T05 is not `done`, correction artifacts are incomplete, required diff is incomplete, or another review prerequisite is missing. |
| `BLOCKED` | Named authority or required evidence is unavailable, or scoped authority precedence cannot be resolved. |

### Finding rules

When F-MAJ-01 or F-MAJ-02 remains `OPEN`, record:

- Finding ID.
- Closure status.
- Remaining contradiction or omission.
- Exact affected artifacts and sections.
- Required correction outcome.
- `New user decision required: yes` or `no`.

Use these IDs for correction-caused or correction-exposed regressions:

```text
F-REG-BLOCK-01
F-REG-MAJ-01
F-REG-MIN-01
```

Do not report editorial preference as a finding.

### Output contract

Return exactly these sections:

```text
1. Verdict
2. Reviewed files
3. Review prerequisites
4. F-MAJ-01 disposition
5. F-MAJ-02 disposition
6. Direct regression findings
7. Git verification
8. Implementation-planning readiness
9. Exact next gate
```

Implementation-planning readiness mapping:

| verdict | output |
|---|---|
| `PASS` | `READY FOR DESIGN CLOSURE` |
| `NEEDS REVISION` | `NOT READY — FINDING CORRECTION REQUIRED` |
| `NOT READY` | `NOT READY — REVIEW PREREQUISITE INCOMPLETE` |
| `BLOCKED` | `BLOCKED — AUTHORITY OR EVIDENCE MISSING` |

Exact next gate after `PASS`:

```text
Synchronize T06 PASS and perform dedicated W011 design closure.
```

Exact next gate after `NEEDS REVISION`:

```text
Correct only the remaining open finding or direct regression,
then perform another independent closure re-review.
```

## Done condition

- The review remains read-only.
- F-MAJ-01 and F-MAJ-02 are judged separately.
- Only direct regression caused by T05 is checked.
- The full W011 design review is not repeated.
- No file is changed.
- No finding is corrected by the reviewer.
- No lifecycle state is synchronized by the reviewer.
- Design closure does not begin.
- Implementation planning does not begin.
- The exact output format is used.
- A `PASS` routes to design closure synchronization.
- `NEEDS REVISION` routes to correction of only the remaining finding or direct regression.

## Verification

- Confirm T05 is `done` before beginning review judgment.
- Confirm the review reads all named authority and evidence within the exact boundary.
- Confirm F-MAJ-01 and F-MAJ-02 use the original affected decision sets.
- Confirm both findings retain `New user decision required: no` unless correction evidence directly proves otherwise.
- Confirm direct regression review does not expand into full W011 review.
- Confirm the reviewer modifies no file.
- Confirm the reviewer performs no correction or lifecycle synchronization.
- Confirm the reviewer does not start design closure or implementation planning.
- Confirm Git inspection is restricted to the 14 named T05-changed files.
- Confirm the complete scoped diff is reviewed.
- Confirm repository-wide cleanliness is not checked or inferred.

## Evidence

- Contract authored for an independent read-only finding-closure re-review.
- The accepted independent execution result is persisted below without rerunning the review.
- No correction, design closure, implementation planning, or production implementation was performed by the reviewer.

### Verdict

`PASS`

### Finding dispositions

| finding | disposition | new user decision required |
|---|---|---|
| F-MAJ-01 | `CLOSED` | no |
| F-MAJ-02 | `CLOSED` | no |

### Direct regression

None.

The re-review confirmed that the correction preserved:

- D-003, D-004, D-005, and D-008;
- public read request and response behavior;
- status, warning, diagnostic, and error behavior;
- current and legacy index separation;
- `validate_records` subject, ordering, diagnostic, and execution-error boundaries;
- authoring-guidance and authoring-transaction public tool catalog entries;
- deferred ownership of current authoring integration by `DRMCP-REQ-MCP-002`.

### Git evidence

- Exact 14-file correction boundary reviewed.
- Group 1 patch: 69,483 of 69,483 bytes.
- Group 2 patch: 29,229 of 29,229 bytes.
- Complete patch confirmed with no truncation.
- Whitespace pass with no findings.
- No staged files.
- LF-to-CRLF messages were advisory only.
- Repository-wide cleanliness was not checked or inferred.

### Readiness and next gate

Implementation-planning readiness:

```text
READY FOR DESIGN CLOSURE
```

Exact next gate:

```text
Synchronize T06 PASS and perform dedicated W011 design closure.
```

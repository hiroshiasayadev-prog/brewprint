# PRODUCT-TASK-SPEC-020-05: Independently review Task responsibility checklists

- **id**: PRODUCT-TASK-SPEC-020-05
- **status**: done
- **date**: 2026-07-02
- **work_item**: PRODUCT-WORK-SPEC-020
- **task_type**: review
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-020-04
  - PRODUCT-TASK-SPEC-019-18
- **outputs**:
  - PRODUCT-TASK-SPEC-020-05

## Goal

Independently review the complete checklist artifact set and issue one integrated verdict.

## Work

- Establish independence from T04 checklist authoring.
- Start semantic review only after W019 records an accepted integrated-review route and closure state.
- Review the decided checklist format, placement, partitioning, and criterion schema.
- Review every common criterion.
- Review every canonical Task-type criterion set.
- Trace checklist content to accepted Task-authoring authority and the accepted validator contract.
- Verify direct machine consumption without content inference.
- Record one `PASS` or `NEEDS REVISION` verdict and a complete named finding set.

This Task must not:

- edit checklist artifacts;
- modify ADRs or Specifications;
- repair findings;
- implement the validator;
- synchronize lifecycle;
- stage or commit changes.

## Done condition

- One independent verdict covers the full checklist artifact set.
- Every canonical Task type is reviewed.
- Every criterion is checked for one binary claim and Task-local judgeability.
- Every material finding identifies the artifact, criterion, required outcome, and owner type.
- The exact next gate is recorded.

## Verification

- Confirm W019 has an accepted integrated-review route and closure state.
- Confirm reviewer independence from T04.
- Confirm all checklist artifacts are included.
- Confirm the review does not infer missing checklist content.
- Confirm no artifact modification occurred.

## Evidence

### Verdict

`PASS`.

The checklist artifact semantics, type coverage, Task-local evaluation boundary, decomposition-versus-execution split, and corrected compactness Evidence pass review.
No blocking, major, or required minor finding remains.

### Reviewer independence

- This session did not author PRODUCT-TASK-SPEC-020-04 or any reviewed prompt asset.
- This session did not modify any checklist artifact, authority, W019 record, W022 record, ADR, or Specification.
- Author reports, self-verification, historical W022 T04 Evidence, and prior-session summaries were not used as proof.
- Current full text and scoped Git Evidence were inspected directly.
- This session changed only PRODUCT-TASK-SPEC-020-05.

### Review precondition

- PRODUCT-TASK-SPEC-019-18 retains its historical `NEEDS REVISION` verdict and F-MAJ-01 through F-MAJ-03.
- PRODUCT-TASK-SPEC-019-19 records explicit user acceptance of all three findings as non-blocking workflow exceptions.
- The findings are not represented as `CLOSED` or as a T18 `PASS`.
- PRODUCT-WORK-SPEC-019 is `done`.
- The accepted W019 route releases this integrated W020 checklist review.
- Review precondition result: `READY`.

### Reviewed artifacts

Workflow and authoring authority:

- `prompt_chappy.md`.
- `skills/design-convergence-workflow/SKILL.md`.
- `skills/design-convergence-workflow/design-review-gate.md`.
- `skills/design-convergence-workflow/work-item-decomposition.md`.
- `skills/design-convergence-workflow/work-item-execution.md`.
- `spec:product.design_records.authoring_standards.agent_authoring_policy`.
- `spec:product.design_records.authoring_standards.writing_standard`.
- `spec:product.design_records.authoring_standards.task_authoring`.
- `spec:product.design_records.authoring_standards.work_item_authoring`.
- `spec:product.responsibility_boundary_validator`.

W019 release authority:

- PRODUCT-WORK-SPEC-019.
- PRODUCT-TASK-SPEC-019-18.
- PRODUCT-TASK-SPEC-019-19.

W020 authoring route:

- PRODUCT-WORK-SPEC-020.
- PRODUCT-TASK-SPEC-020-02.
- PRODUCT-TASK-SPEC-020-03.
- PRODUCT-TASK-SPEC-020-04.
- PRODUCT-TASK-SPEC-020-05.

Additional decomposition and execution authority:

- PRODUCT-ADR-SPEC-004.
- PRODUCT-ADR-SPEC-005.
- PRODUCT-ADR-SPEC-010.
- PRODUCT-WORK-SPEC-022.
- PRODUCT-TASK-SPEC-022-01 through PRODUCT-TASK-SPEC-022-05.

Checklist artifact set:

- `skills/task-responsibility-boundary-validator/SKILL.md`.
- `skills/task-responsibility-boundary-validator/prompts/evaluator-instructions.md`.
- `skills/task-responsibility-boundary-validator/prompts/common.md`.
- All eleven files under `skills/task-responsibility-boundary-validator/prompts/task-types/`:
  `investigation`, `decision`, `authoring`, `implementation`, `review`, `correction`, `verification`, `coordination`, `work_item_decomposition`, `work_item_execution`, and `synchronization`.

### Complete checklist-set result

- Evaluator instructions exist.
- The common checklist exists.
- The type directory contains exactly one file for each of the eleven canonical `task_type` values and no additional type file.
- `SKILL.md` deterministically composes evaluator instructions, common criteria, and exactly the declared canonical type file.
- Invocation-time criterion addition or removal is prohibited.
- `task_authoring` is explicitly authoritative when a prompt fragment conflicts with canonical authority.
- Criterion IDs are unique within each file.
- Evaluator instructions require exact ID, boolean result, concise reason, and relevant Task section.
- Missing Task-local information produces `false`; external inference is prohibited.
- No criterion requires inspection of a parent Work Item, child Work Item, repository state, or execution result.
- No validator implementation or external response API contract is defined by the checklist set.

### Criterion and type coverage result

| checklist | result |
|---|---|
| common | PASS |
| investigation | PASS |
| decision | PASS |
| authoring | PASS |
| implementation | PASS |
| review | PASS |
| correction | PASS |
| verification | PASS |
| coordination | PASS |
| work_item_decomposition | PASS |
| work_item_execution | PASS |
| synchronization | PASS |

Each criterion expresses one binary responsibility-boundary claim.
The common checklist covers one primary outcome, one completion judgment, section alignment, supporting-action limits, acceptance or verification boundary, owner, release decision, required independence, and unresolved-decision separation.
Type files project the canonical primary outcome, completion judgment, and prohibited overlaps without unnecessary common-criterion repetition.

### Work Item decomposition result

`PASS`.

- The checklist owns child Work Item creation or split, coarse parent-level routing, and decomposition completion only.
- It excludes child completion, child-owned Tasks, deliverables, procedures, decisions, and reviews.
- It excludes dependency, blocker, owner, writer-order, review-order, and release-order changes.
- It excludes separate implementation, review, correction, and synchronization outcomes.
- Every criterion is judgeable from the Task record without inspecting parent or child Work Item content.

### Work Item execution result

`PASS`.

- The checklist requires exactly one child reference through `work_item_ref`.
- It limits Goal and Work to the child Work Item's coarse purpose and parent-level execution role.
- It uses referenced child completion as the sole completion boundary.
- It excludes separate child deliverable, judgment, and verification-result completion requirements.
- It excludes child-owned Tasks, deliverables, procedures, decisions, and reviews.
- It excludes child creation, split, redefinition, and internal coordination.
- It excludes separate implementation, review, correction, and synchronization outcomes.
- It judges the written Task contract only and does not inspect the referenced Work Item's actual status or content.

### Decomposition versus execution result

`PASS`.

`work_item_decomposition` ends when child boundaries and coarse routing exist.
`work_item_execution` represents one already-created child and waits on that child's completion boundary.
The two checklists do not share a primary outcome or completion judgment.

### Compactness verification

Character counts exclude the terminal newline from each file.
Token values are character-based estimates.

Verified required values:

| item | criteria | lines | words | characters | approximate tokens |
|---|---:|---:|---:|---:|---:|
| work_item_decomposition | 7 | 7 | 104 | 702 | 176 |
| work_item_execution | 8 | 8 | 111 | 774 | 194 |
| evaluator instructions + common + work_item_execution | n/a | 27 | 342 | 2222 | 556 |

All line, word, and character counts in the T04 compactness table now match current files under the terminal-newline-excluding convention.
The directly derived coordination token estimate was updated from 109 to 108.
The maximum runtime composition and the required decomposition and execution counts remain correct.

### Scope-boundary result

`PASS`.

The artifact set contains no validator implementation, DRMCP integration, provider or model selection, retry, timeout, decode policy, automatic Task correction, graph validation, complete executor-readiness judgment, or external response API contract.

### Blocking findings

None.

### Major findings

None.

### Minor findings

None.

### Corrected factual discrepancy

- The six T04 character counts and one directly derived token estimate were corrected before final acceptance.
- The discrepancy affected only compactness Evidence and did not alter checklist semantics.
- The user explicitly authorized direct correction without a separate correction and finding-closure review route.

### Advisories

- LF-to-CRLF conversion warnings are advisory only.

### Scoped Git Evidence

- Scope: `skills/task-responsibility-boundary-validator/`, PRODUCT-TASK-SPEC-020-04, and PRODUCT-TASK-SPEC-020-05.
- Scoped worktree inspection result: `pass`.
- Scoped state contains sixteen untracked, unstaged files.
- Scoped staged changes: none.
- Scoped whitespace result: `pass`; no whitespace findings.
- Scoped textual patch result: `pass`.
- Returned patch: 33,689 of 33,689 bytes; not truncated.
- Untracked `git diff --no-index` exit code `1` was treated as the expected file-difference result.
- Repository-wide cleanliness was not inspected or claimed.

### Exact next gate

```text
PRODUCT-TASK-SPEC-020-06
closure synchronization
```

PRODUCT-TASK-SPEC-020-06 is released by this `PASS` verdict.

### Explicitly not performed

- No checklist artifact was edited.
- PRODUCT-TASK-SPEC-020-04 compactness Evidence was corrected before final acceptance under explicit user authorization.
- No W019 or W022 record was edited.
- No ADR or Specification was edited.
- No Task graph or lifecycle synchronization was performed.
- No validator implementation, DRMCP integration, stage, or commit was performed.

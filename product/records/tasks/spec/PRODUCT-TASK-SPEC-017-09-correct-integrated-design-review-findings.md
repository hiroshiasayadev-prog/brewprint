# PRODUCT-TASK-SPEC-017-09: Correct integrated design-review findings

- **id**: PRODUCT-TASK-SPEC-017-09
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-017
- **source_requirement**: PRODUCT-REQ-SPEC-006
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-017-08
- **outputs**:
  - PRODUCT-WORK-SPEC-016
  - PRODUCT-WORK-SPEC-017
  - PRODUCT-TASK-SPEC-016-08
  - PRODUCT-TASK-SPEC-016-09
  - PRODUCT-TASK-SPEC-016-10
  - PRODUCT-TASK-SPEC-016-11
  - PRODUCT-TASK-SPEC-017-09
  - PRODUCT-TASK-SPEC-017-11

## Goal

Correct only `F-MAJ-01` and `F-MAJ-02` from `PRODUCT-TASK-SPEC-017-08`.

Preserve the accepted Requirement, ADR, and Specification semantics.

Leave finding-closure judgment to `PRODUCT-TASK-SPEC-017-10`.

## Work

### Correction authority

- Use `PRODUCT-TASK-SPEC-017-08` as the only finding authority.
- Do not add a user decision.
- Do not reopen an accepted decision.
- Do not change Requirement, ADR, or canonical Specification content.
- Do not judge either finding closed.

### Exact correction target set

| record | allowed correction |
|---|---|
| `PRODUCT-WORK-SPEC-016` | Align `## Task flow` and `## Task Candidates` with the single integrated-review route. |
| `PRODUCT-WORK-SPEC-017` | Correct `F-MAJ-01` in `## Impact Scope` and `## Completion Condition`. Align `## Task flow` and `## Task Candidates` with the executable two-branch review route. |
| `PRODUCT-TASK-SPEC-016-08` | Mark the separate W016 review route as not required and superseded by W017 T08. Keep lifecycle status unchanged. |
| `PRODUCT-TASK-SPEC-016-09` | Mark the separate W016 correction route as not required and superseded by W017 T09. Keep lifecycle status unchanged. |
| `PRODUCT-TASK-SPEC-016-10` | Mark the separate W016 re-review route as not required and superseded by W017 T10. Keep lifecycle status unchanged. |
| `PRODUCT-TASK-SPEC-016-11` | Base metadata dependency on W017 T08. Add the conditional acceptance gate for direct PASS or corrected NEEDS REVISION. |
| `PRODUCT-TASK-SPEC-017-09` | Record execution Evidence, exact changed records, and finding-to-change mappings. |
| `PRODUCT-TASK-SPEC-017-11` | Base metadata dependency on W017 T08. Add the conditional acceptance gate for direct PASS or corrected NEEDS REVISION. |

Do not add another correction target.

Stop with `BLOCKED` without changing any record when an additional target is unavoidable.
Record the exact reason in this Task Evidence.

### F-MAJ-01 correction contract

Target: `PRODUCT-WORK-SPEC-017`.

Allowed sections:

- `## Impact Scope`
- `## Completion Condition`

Required result:

- Only Work Items persist required non-empty `source_refs`.
- Tasks persist no source field.
- Tasks do not persist `source_refs`.
- Task provenance is reached through `work_item`.
- Legacy Task `source_requirement` migration is removal-only.
- Bootstrap workflow Tasks retain their old metadata shape only as a workflow-local exclusion.
- The bootstrap exclusion is not a canonical metadata exception.

Do not change the accepted Requirement, ADR, or Specification semantics.

### F-MAJ-02 correction contract

#### Integrated review ownership

- `PRODUCT-TASK-SPEC-017-08` is the only cross-requirement integrated review owner for REQ-005 and REQ-006.
- W016 does not require a separate independent review.
- W016 closure consumes W017 integrated-review evidence.
- W016 T08, T09, and T10 do not execute as a competing review route.
- Keep W016 T08, T09, and T10 lifecycle status unchanged.
- Do not add `canceled` or another noncanonical lifecycle status.
- State their `not_required` and superseded-route disposition in each Task body or Evidence.

#### Review branch contract

| T08 verdict | required route |
|---|---|
| `PASS` | T09 `not_required`; T10 `not_required`; W016 T11 and W017 T11 may execute from accepted T08 Evidence. |
| `NEEDS REVISION` | W017 T09 correction, then W017 T10 independent finding-closure re-review, then W016 T11 and W017 T11. |

Do not execute T09 or T10 as synthetic no-op Tasks to satisfy metadata dependencies.

The current verdict is `NEEDS REVISION`.

#### Closure dependency and conditional gate

For both W016 T11 and W017 T11:

- Set metadata `depends_on` to `PRODUCT-TASK-SPEC-017-08`.
- Keep the two closure Tasks separate because each Task synchronizes one Work Item lifecycle.
- Make both closure Tasks consume the same authoritative integrated-review Evidence.
- Add a body-level conditional acceptance gate.

The conditional acceptance gate must enforce:

| T08 verdict | closure evidence |
|---|---|
| `PASS` | Accepted T08 Evidence is sufficient. T09 and T10 remain not required. |
| `NEEDS REVISION` | T09 must be complete. T10 must independently close every required finding. |

Closure must not execute while T08 remains `NEEDS REVISION` without accepted T10 finding dispositions.

#### Correction ownership

W017 T09 owns correction of `F-MAJ-01` and `F-MAJ-02`.

W017 T09 may change only:

- the named Work Item sections;
- the named Task graph records;
- closure dependencies and conditional acceptance gates;
- obsolete duplicate-route dispositions;
- this Task's outputs and Evidence.

W017 T09 must not change:

- Requirements;
- ADRs;
- canonical Specifications;
- W017 T08 review result;
- W017 T10 independent re-review contract;
- accepted decision-ledger entries;
- production source, tests, or fixtures;
- existing-record migration targets.

### Finding-to-target mapping

| finding | exact targets | required result |
|---|---|---|
| `F-MAJ-01` | `PRODUCT-WORK-SPEC-017` | Remove stale Task `source_refs` wording from `## Impact Scope` and `## Completion Condition`. State Work Item-only `source_refs`, Task provenance through `work_item`, and removal-only Task source migration. |
| `F-MAJ-02` | `PRODUCT-WORK-SPEC-016`, `PRODUCT-WORK-SPEC-017`, `PRODUCT-TASK-SPEC-016-08`, `PRODUCT-TASK-SPEC-016-09`, `PRODUCT-TASK-SPEC-016-10`, `PRODUCT-TASK-SPEC-016-11`, `PRODUCT-TASK-SPEC-017-09`, `PRODUCT-TASK-SPEC-017-11` | Establish W017 T08 as the only integrated review owner. Make PASS and NEEDS REVISION routes executable. Preserve separate W016 and W017 closure ownership. |

### T10 handoff

`PRODUCT-TASK-SPEC-017-10` remains unchanged and read-only during correction.

T10 independently re-reviews:

- `F-MAJ-01` against `PRODUCT-WORK-SPEC-017`;
- `F-MAJ-02` against every exact correction target;
- direct consistency effects only;
- this Task's finding-to-change Evidence.

T10 must not rely on the correction author's closure claim.

## Done condition

- The stale W017 Work Item wording named by `F-MAJ-01` is removed.
- Only Work Items own persisted `source_refs`.
- Tasks persist no source field and use `work_item` for provenance.
- The `F-MAJ-02` review, correction, re-review, and closure graph is reachable for both verdict branches.
- W017 T08 is the only integrated review owner.
- Correction and independent finding-closure re-review remain separate.
- W016 and W017 retain separate closure ownership.
- Both closure Tasks consume the same authoritative integrated-review Evidence.
- Accepted Requirement, ADR, and Specification semantics remain unchanged.
- T10 receives the exact bounded re-review set and finding-to-change mapping.
- This Task does not classify either finding as closed.

## Verification

- Map every changed line to `F-MAJ-01` or `F-MAJ-02`.
- Confirm only the eight records in `outputs` changed.
- Confirm Requirements, ADRs, and canonical Specifications are unchanged.
- Confirm W017 T08 and its verdict are unchanged.
- Confirm W017 T10 is unchanged and absent from the writable set.
- Confirm the PASS branch does not require T09 or T10 no-op execution.
- Confirm the NEEDS REVISION branch is T09, then T10, then both closure Tasks.
- Confirm W016 and W017 closure Tasks reference the same accepted review Evidence.
- Confirm the correction author did not perform finding closure.
- Confirm no migration, production implementation, lifecycle closure, stage, or commit occurred.
- Run scoped whitespace inspection for the exact changed records.
- Do not inspect or claim repository-wide clean status.

## Evidence

### Execution state

- Task status: `done`.
- Execution date: `2026-07-01`.
- Correction authority: `PRODUCT-TASK-SPEC-017-08`, `F-MAJ-01` and `F-MAJ-02`.
- New user decision required: no.
- Accepted Requirement, ADR, and Specification semantics: unchanged.
- DRMCP availability: non-operational under the current agent authoring policy.
- Filesystem fallback reason: DRMCP retrieval and authoring transactions are not operational.
- Additional correction target required: no.

### Changed records

- `product/records/work-items/spec/PRODUCT-WORK-SPEC-016-define-typed-single-responsibility-task-contract.md`
- `product/records/work-items/spec/PRODUCT-WORK-SPEC-017-define-generic-workflow-source-relations.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-016-08-review-final-req-005-design.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-016-09-correct-review-findings.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-016-10-rereview-finding-closure.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-016-11-synchronize-lifecycle-and-evidence-closure.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-017-09-correct-integrated-design-review-findings.md`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-017-11-synchronize-source-relation-design-closure.md`

### F-MAJ-01 mapping

- Changed artifact: `PRODUCT-WORK-SPEC-017`.
- Changed sections: `## Impact Scope`, `## Completion Condition`, and direct workflow consistency sections.
- Removed stale statement: Task `source_requirement` migrates to Task `source_refs`.
- Removed stale statement: Work Items and Tasks share one generic persisted `source_refs` contract.
- Resulting contract: only Work Items persist required non-empty `source_refs`.
- Resulting Task contract: Tasks persist no source field.
- Resulting provenance path: Task provenance is reached through `work_item`.
- Resulting migration contract: legacy Task `source_requirement` removal is replacement-free.
- Bootstrap disposition: old Task metadata remains a workflow-local exclusion only.
- Canonical metadata exception created: no.

### F-MAJ-02 mapping

- `PRODUCT-WORK-SPEC-016`: removed the independent duplicate review route and routed acceptance through W017 T08.
- `PRODUCT-WORK-SPEC-017`: established W017 T08 as the only integrated review owner and added both verdict branches.
- `PRODUCT-TASK-SPEC-016-08`: retained as a superseded, not-required, non-executable review route owned by W017 T08.
- `PRODUCT-TASK-SPEC-016-09`: retained as a superseded, not-required, non-executable correction route owned by W017 T09.
- `PRODUCT-TASK-SPEC-016-10`: retained as a superseded, not-required, non-executable re-review route owned by W017 T10.
- `PRODUCT-TASK-SPEC-016-11`: metadata predecessor changed to W017 T08. Conditional closure acceptance gate added.
- `PRODUCT-TASK-SPEC-017-11`: metadata predecessor changed to W017 T08. Conditional closure acceptance gate added.
- `PRODUCT-TASK-SPEC-017-09`: correction ownership, exact targets, execution result, and re-review handoff recorded.
- W016 and W017 closure ownership remains separate.
- Synthetic no-op Task execution is prohibited.

### Integrated review ownership

- W017 T08 owns the authoritative cross-requirement verdict.
- W017 T09 owns correction of `F-MAJ-01` and `F-MAJ-02`.
- W017 T10 owns independent finding-closure re-review.
- W016 T11 owns only W016 lifecycle and Evidence closure.
- W017 T11 owns only W017 lifecycle and Evidence closure.

### Branch result

| T08 verdict | executable route |
|---|---|
| `PASS` | T08 -> W016 T11 and W017 T11. T09 and T10 are not required. |
| `NEEDS REVISION` | T08 -> T09 -> T10 -> W016 T11 and W017 T11. |

- Current active branch: `NEEDS REVISION`.
- Current next gate: `PRODUCT-TASK-SPEC-017-10`.
- Neither closure Task is executable before accepted T10 finding dispositions.

### T10 handoff

T10 remains read-only and unchanged.

T10 re-review boundary:

- `F-MAJ-01` against `PRODUCT-WORK-SPEC-017`;
- `F-MAJ-02` against all eight correction targets;
- direct consistency effects caused by this correction;
- this Task's finding-to-change Evidence.

T09 does not classify either finding as closed.

### Verification result

- Exact writable set contained eight records.
- No additional target was required.
- W017 T08 was unchanged.
- W017 T10 was unchanged.
- Requirements were unchanged.
- ADRs were unchanged.
- Specifications were unchanged.
- W016 T08, T09, and T10 remain `not_started`.
- W016 and W017 remain `in_progress`.
- W016 T11 and W017 T11 remain `not_started`.
- Both T11 metadata dependencies are W017 T08.
- The direct `PASS` branch does not require T09 or T10 execution.
- The `NEEDS REVISION` branch requires T09, then accepted T10 dispositions, then both closure Tasks.
- Scoped full-text and whitespace verification completed for the eight records.
- Repository-wide clean status was not inspected or claimed.

### Prohibited-work confirmation

- Requirement change: none.
- ADR change: none.
- Specification change: none.
- W017 T08 change: none.
- W017 T10 change: none.
- Decision-ledger change: none.
- Migration execution: none.
- Production implementation: none.
- Lifecycle closure: none.
- Finding-closure judgment: none.
- Stage or commit: none.

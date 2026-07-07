# PRODUCT-TASK-SPEC-025-07: Extract boundary-vocabulary entries from W009–W012

- **id**: PRODUCT-TASK-SPEC-025-07
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-025
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-025-02
- **outputs**:
  - PRODUCT-TASK-SPEC-025-07

## Goal

Extract boundary-vocabulary log entries from every `status: done` Task under `product/records/tasks/spec/` whose parent Work Item is in the `PRODUCT-WORK-SPEC-009` through `PRODUCT-WORK-SPEC-012` range (~30 Tasks), and record them directly in this Task's own `## Evidence`, under the Investigation-Task lightweight Evidence exception (PRODUCT-TASK-SPEC-025-02).

## Work

- Fix this Task's corpus at start: the exact set of `status: done` Task IDs under `product/records/tasks/spec/` whose `work_item` falls in `PRODUCT-WORK-SPEC-009`–`PRODUCT-WORK-SPEC-012`.
- Read each in-scope Task's Work and Done Condition sections. Some older Tasks have no `task_type` field; infer the responsibility cluster from title, Work, and Done Condition instead of skipping them.
- Find phrases where a responsibility is expressed without its owning verb (e.g. a synonym standing in for `author`, `review`, `coordinate`), or where a verb carries an unusual meaning in context (e.g. `Fix`).
- Record each finding as a 4-column entry: phrase (exact), source Task ID, vocabulary target (what responsibility or term it stands in for), effective meaning in that context.
- Record a scanned entry for every Task read, including Tasks with zero findings. Do not force findings.
- Do not rank by frequency.
- Do not decide canonical terms.
- Record the accepted user judgment and reason for using the lightweight exception (PRODUCT-TASK-SPEC-025-01 D-006) directly in `## Evidence`.
- Write only to this Task's own Evidence. Do not write to any sibling Task's file.

Responsibility clusters (classification target for findings, not a further split axis):

- decision / authoring / coordination (decide, define, fix, determine, materialize, project, encode, update, align, amend)
- review / verification / correction (review, inspect, verify, validate, check, re-check, confirm, assess, correct, repair, close finding)
- investigation / decision / design (investigate, research, analyze, compare, identify, conclude, recommend, select)
- graph / lifecycle / execution (coordinate, materialize, create, route, release, synchronize, close, execute, decompose, propagate)

This Task must not:

- author or edit `skills/task-boundary-vocabulary/` files directly;
- decide canonical term definitions;
- create a separate formal Investigation record (covered by the accepted exception);
- rewrite or reinterpret source Task content;
- write to any Task file other than this one.

## Done condition

- Every `status: done` Task in the `PRODUCT-WORK-SPEC-009`–`PRODUCT-WORK-SPEC-012` range at this Task's fixed cutoff has a scanned entry, including zero-finding Tasks.
- Every logged entry cites an exact, existing source Task ID.
- The accepted user judgment and reason for using the lightweight exception is recorded in this Task's own Evidence.

## Verification

- Confirm every entry cites a real, existing Task ID within the assigned range.
- Confirm no canonical term decision was made.
- Confirm the exception-use judgment and reason are recorded in Evidence.
- Confirm scanned-Task count matches the fixed corpus for this range.

## Evidence

### Execution boundary and lightweight exception

- Fixed cutoff: 2026-07-03, before this Evidence was authored.
- Fixed corpus: 30 `status: done` Tasks whose `work_item` is `PRODUCT-WORK-SPEC-009`, `PRODUCT-WORK-SPEC-010`, `PRODUCT-WORK-SPEC-011`, or `PRODUCT-WORK-SPEC-012`.
- Accepted user judgment: PRODUCT-TASK-SPEC-025-01 D-006 permits this bounded research result to remain in this Task's Evidence.
- Exception reason: a formal Investigation record is disproportionate for simple corpus-log extraction and would recreate the heavyweight route rejected by the user. Relabeling this work as `authoring` was also rejected because the extraction does not produce a canonical artifact outside this Task.
- Downstream conclusion route: PRODUCT-TASK-SPEC-025-04 will reconcile and conclude on this log, including a possible conclusion that no entry requires canonical projection.
- DRMCP is non-operational. Filesystem authoring was the required fallback.

### Fixed corpus and scan ledger

| source Task | inferred responsibility cluster | scan result |
|---|---|---|
| PRODUCT-TASK-SPEC-009-01 | investigation / execution | 2 findings |
| PRODUCT-TASK-SPEC-009-02 | authoring / correction | zero findings; the Work uses explicit migration actions |
| PRODUCT-TASK-SPEC-009-03 | verification / correction | 1 finding |
| PRODUCT-TASK-SPEC-009-04 | review | 3 findings |
| PRODUCT-TASK-SPEC-009-05 | correction / lifecycle | 2 findings |
| PRODUCT-TASK-SPEC-010-01 | decision | 2 findings |
| PRODUCT-TASK-SPEC-010-02 | authoring | zero findings; the Work uses explicit authoring actions |
| PRODUCT-TASK-SPEC-010-03 | review / correction | 2 findings |
| PRODUCT-TASK-SPEC-010-04 | synchronization / lifecycle | 2 findings |
| PRODUCT-TASK-SPEC-011-01 | correction | zero findings; the Work uses the direct verb `Correct` |
| PRODUCT-TASK-SPEC-011-02 | authoring | 1 finding |
| PRODUCT-TASK-SPEC-011-03 | authoring / correction | zero findings; the Work uses explicit revision actions |
| PRODUCT-TASK-SPEC-011-04 | authoring | 1 finding |
| PRODUCT-TASK-SPEC-011-05 | authoring | 1 finding |
| PRODUCT-TASK-SPEC-011-06 | authoring | 1 finding |
| PRODUCT-TASK-SPEC-011-07 | authoring | 1 finding |
| PRODUCT-TASK-SPEC-011-08 | verification / review / correction | 3 findings |
| PRODUCT-TASK-SPEC-012-01 | investigation / verification / coordination | 2 findings |
| PRODUCT-TASK-SPEC-012-02 | authoring | zero findings; the Work uses explicit authoring actions |
| PRODUCT-TASK-SPEC-012-03 | authoring / correction | zero findings; the Work uses explicit keep, move, place, and replace actions |
| PRODUCT-TASK-SPEC-012-04 | authoring / correction | 1 finding |
| PRODUCT-TASK-SPEC-012-05 | investigation / authoring / coordination | 1 finding |
| PRODUCT-TASK-SPEC-012-06 | execution | zero findings; the Work uses the direct verb `Move` |
| PRODUCT-TASK-SPEC-012-07 | execution | zero findings; the Work uses the direct verb `Move` |
| PRODUCT-TASK-SPEC-012-08 | coordination / verification / authoring | 2 findings |
| PRODUCT-TASK-SPEC-012-09 | coordination / lifecycle | 1 finding |
| PRODUCT-TASK-SPEC-012-10 | synchronization | 1 finding |
| PRODUCT-TASK-SPEC-012-11 | verification / correction | 1 finding |
| PRODUCT-TASK-SPEC-012-12 | review | 2 findings |
| PRODUCT-TASK-SPEC-012-13 | correction / lifecycle | 2 findings |

### Boundary-vocabulary log

The vocabulary target is a provisional responsibility classification. It does not decide canonical terminology.

| phrase (exact) | source Task ID | vocabulary target | effective meaning in that context |
|---|---|---|---|
| `inventory run` | PRODUCT-TASK-SPEC-009-01 | investigation / verification | Execute the baseline inventory command and record its result. |
| `staging complete` | PRODUCT-TASK-SPEC-009-01 | graph / lifecycle / execution | Prepare the complete preservation-copy set before migration. |
| `fix errors` | PRODUCT-TASK-SPEC-009-03 | review / verification / correction | Repair every strict-validator diagnostic until validation passes. |
| `review complete` | PRODUCT-TASK-SPEC-009-04 | review / verification / correction | Produce the independent review report for the full target set. |
| `findings classified` | PRODUCT-TASK-SPEC-009-04 | review / verification / correction | Assign each review finding a required-fix or defer disposition. |
| `user sign-off` | PRODUCT-TASK-SPEC-009-04 | decision / authoring / coordination | Obtain explicit human approval to pass the review gate. |
| `corrections applied` | PRODUCT-TASK-SPEC-009-05 | review / verification / correction | Repair all must-fix findings from the preceding review. |
| `WORK-009 updated` | PRODUCT-TASK-SPEC-009-05 | graph / lifecycle / execution | Synchronize parent Work Item Evidence and lifecycle state. |
| `Rule set confirmed` | PRODUCT-TASK-SPEC-010-01 | investigation / decision / design | Approve the selected and deferred rule set. |
| `Scope decisions resolved` | PRODUCT-TASK-SPEC-010-01 | investigation / decision / design | Settle every named scope choice before authoring. |
| `Apply required findings` | PRODUCT-TASK-SPEC-010-03 | review / verification / correction | Implement the corrections demanded by review findings. |
| `Spec passes review` | PRODUCT-TASK-SPEC-010-03 | review / verification / correction | Accept that no required spec change remains. |
| `Work item closed` | PRODUCT-TASK-SPEC-010-04 | graph / lifecycle / execution | Move the parent Work Item to its terminal lifecycle state. |
| `Traceability complete` | PRODUCT-TASK-SPEC-010-04 | graph / lifecycle / execution | Persist the complete Task relation list on the parent Work Item. |
| `Boundary published` | PRODUCT-TASK-SPEC-011-02 | decision / authoring / coordination | Create and register the canonical artifact-boundary specification. |
| `Guide published` | PRODUCT-TASK-SPEC-011-04 | decision / authoring / coordination | Create and register the canonical requirement authoring guide. |
| `Guide published` | PRODUCT-TASK-SPEC-011-05 | decision / authoring / coordination | Create and register the canonical Work Item authoring guide. |
| `Guide published` | PRODUCT-TASK-SPEC-011-06 | decision / authoring / coordination | Create and register the canonical Task authoring guide. |
| `Guide published` | PRODUCT-TASK-SPEC-011-07 | decision / authoring / coordination | Create and register the canonical Investigation authoring guide. |
| `Spec coverage reviewed` | PRODUCT-TASK-SPEC-011-08 | review / verification / correction | Evaluate spec-authoring coverage and record the verdict or gaps. |
| `Drift resolved` | PRODUCT-TASK-SPEC-011-08 | review / verification / correction | Repair cross-guide ownership, terminology, metadata, and lifecycle inconsistencies. |
| `Work item updated` | PRODUCT-TASK-SPEC-011-08 | graph / lifecycle / execution | Propagate the final review result into the parent Work Item. |
| `one explicit disposition` | PRODUCT-TASK-SPEC-012-01 | investigation / decision / design | Route each source file or mixed section to one action, owner, or decision gate. |
| `Correct task scope or dependency errors` | PRODUCT-TASK-SPEC-012-01 | graph / lifecycle / execution | Repair Task boundaries or dependency edges before implementation starts. |
| `preserved without canonical ownership claims` | PRODUCT-TASK-SPEC-012-04 | decision / authoring / coordination | Retain unresolved material as non-canonical staging evidence. |
| `follow-up destination or deletion rationale` | PRODUCT-TASK-SPEC-012-05 | investigation / decision / design | Assign every unadopted mechanism a terminal follow-up or deletion disposition. |
| `app-local preservation` | PRODUCT-TASK-SPEC-012-08 | decision / authoring / coordination | Retain removed operational contract content under its app-local owner. |
| `retained app-local owner or deletion rationale` | PRODUCT-TASK-SPEC-012-08 | investigation / decision / design | Prove each removed statement has an owner or a justified deletion. |
| `four-way exit classification` | PRODUCT-TASK-SPEC-012-09 | investigation / decision / design | Assign every staged item to one accepted exit route. |
| `mechanical reference synchronization` | PRODUCT-TASK-SPEC-012-10 | graph / lifecycle / execution | Update refs, parents, Topics rows, and links without semantic redesign. |
| `Pre-existing diagnostics remain separately attributed` | PRODUCT-TASK-SPEC-012-11 | investigation / decision / design | Classify non-migration diagnostics without absorbing them into correction scope. |
| `` `PASS` or `NEEDS REVISION` `` | PRODUCT-TASK-SPEC-012-12 | review / verification / correction | Issue the bounded independent-review verdict. |
| `required corrections` | PRODUCT-TASK-SPEC-012-12 | review / verification / correction | Identify the exact correction set required before closure. |
| `Apply all blocker and must-fix findings` | PRODUCT-TASK-SPEC-012-13 | review / verification / correction | Repair every closure-blocking finding from the independent review. |
| `disposition of optional findings` | PRODUCT-TASK-SPEC-012-13 | investigation / decision / design | Record whether each advisory is applied, deferred, or left unchanged. |

### Verification result

| check | result |
|---|---|
| Fixed corpus count | 30 Tasks: W009 5, W010 4, W011 8, W012 13. |
| Scanned entries | 30 of 30 Tasks recorded. |
| Finding entries | 35 entries from 22 Tasks. |
| Zero-finding Tasks | 8 Tasks, each recorded in the scan ledger. |
| Phrase scope | Every phrase is copied from the source Task's `## Work` or `## Done condition`. |
| Source validity | Every finding cites an existing Task in PRODUCT-WORK-SPEC-009 through PRODUCT-WORK-SPEC-012. |
| Canonical-term decisions | None. Vocabulary targets are classification aids for downstream reconciliation only. |
| Exception judgment | PRODUCT-TASK-SPEC-025-01 D-006 and its reason are recorded above. |
| Write boundary | Only PRODUCT-TASK-SPEC-025-07 was modified. No sibling Task or `skills/task-boundary-vocabulary/` file was changed. |

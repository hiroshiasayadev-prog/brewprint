# PRODUCT-TASK-SPEC-025-09: Extract boundary-vocabulary entries from W018–W019

- **id**: PRODUCT-TASK-SPEC-025-09
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-025
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-025-02
- **outputs**:
  - PRODUCT-TASK-SPEC-025-09

## Goal

Extract boundary-vocabulary log entries from every `status: done` Task under `product/records/tasks/spec/` whose parent Work Item is in the `PRODUCT-WORK-SPEC-018` through `PRODUCT-WORK-SPEC-019` range (~42 Tasks), and record them directly in this Task's own `## Evidence`, under the Investigation-Task lightweight Evidence exception (PRODUCT-TASK-SPEC-025-02).

## Work

- Fix this Task's corpus at start: the exact set of `status: done` Task IDs under `product/records/tasks/spec/` whose `work_item` falls in `PRODUCT-WORK-SPEC-018`–`PRODUCT-WORK-SPEC-019`.
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

- Every `status: done` Task in the `PRODUCT-WORK-SPEC-018`–`PRODUCT-WORK-SPEC-019` range at this Task's fixed cutoff has a scanned entry, including zero-finding Tasks.
- Every logged entry cites an exact, existing source Task ID.
- The accepted user judgment and reason for using the lightweight exception is recorded in this Task's own Evidence.

## Verification

- Confirm every entry cites a real, existing Task ID within the assigned range.
- Confirm no canonical term decision was made.
- Confirm the exception-use judgment and reason are recorded in Evidence.
- Confirm scanned-Task count matches the fixed corpus for this range.

## Evidence

### Lightweight-exception judgment

- Accepted judgment: PRODUCT-TASK-SPEC-025-01 D-006 permits this bounded Investigation to record results directly in Task Evidence.
- Reason: a separate formal Investigation record is disproportionate for corpus-log extraction and would recreate the rejected heavyweight route.
- Downstream conclusion owner: PRODUCT-TASK-SPEC-025-04 will verify, reconcile, cluster, and either project or reject these candidates.
- Boundary: this Task records raw candidates only. It makes no canonical-term decision and changes no vocabulary skill file.
- Access mode: DRMCP is non-operational. Filesystem read and authoring were the required fallback.

### Fixed corpus

- Cutoff: 2026-07-03 19:20 JST.
- Selection rule: `status: done` Tasks whose `work_item` is PRODUCT-WORK-SPEC-018 or PRODUCT-WORK-SPEC-019.
- Fixed count: 42 Tasks.
- Range result: PRODUCT-TASK-SPEC-018-01 through PRODUCT-TASK-SPEC-018-20, and PRODUCT-TASK-SPEC-019-01 through PRODUCT-TASK-SPEC-019-22.
- Every selected Task had an explicit `task_type` field.

### Boundary-vocabulary log

| phrase (exact) | source Task ID | vocabulary target | effective meaning in context |
|---|---|---|---|
| F-001 — “Persist every accepted item as D-001 through D-023.” | PRODUCT-TASK-SPEC-018-01 | decision-ledger recording | Write each accepted judgment into the Task-local decision ledger. |
| F-002 — “Derive the exact next authoring Task boundaries.” | PRODUCT-TASK-SPEC-018-02 | downstream-boundary determination | Determine the next authoring Task scopes from the completed ADR route. |
| F-003 — “Assign downstream references and Evidence to downstream Tasks.” | PRODUCT-TASK-SPEC-018-03 | canonical ownership projection | Amend ADR text so downstream records own their references and Evidence. |
| F-004 — “Point `prompt_chappy.md` to the successor and its phase companions.” | PRODUCT-TASK-SPEC-018-05 | instruction authoring | Edit the active instruction pointer to reference the successor workflow files. |
| F-005 — “Project the accepted Requirement identity boundary into Requirement authoring rules.” | PRODUCT-TASK-SPEC-018-06 | Specification authoring | Encode an accepted decision as a normative Requirement-authoring rule. |
| F-006 — “Record T19 as the accepted finding-closure review.” | PRODUCT-TASK-SPEC-018-08 | closure synchronization | Propagate the accepted review authority into closure Evidence. |
| F-007 — “Treat T07 prerequisite `P-BLK-01` as a mechanically necessary missing-owner repair.” | PRODUCT-TASK-SPEC-018-09 | graph-repair classification | Classify the prerequisite gap as a workflow-owner repair rather than a review finding. |
| F-008 — “Fix canonical target changes and every required graph amendment.” | PRODUCT-TASK-SPEC-018-11 | decision | Decide and state target and graph changes without applying them. |
| F-009 — “Change T07 from `blocked` to `not_started`.” | PRODUCT-TASK-SPEC-018-13 | lifecycle graph update | Update the Task lifecycle state after materializing the missing authoring route. |
| F-010 — “Assign child Work Item creation and split to the new type in `work_item_authoring`.” | PRODUCT-TASK-SPEC-018-14 | Specification authoring | Encode the new Task-type ownership rule in canonical authoring guidance. |
| F-011 — “Keep T08 reserved until T19 closes every required finding.” | PRODUCT-TASK-SPEC-018-15 | release-route coordination | Defer closure-Task materialization until the independent finding gate completes. |
| F-012 — “Fix the durable amendment boundary and exact authoring targets.” | PRODUCT-TASK-SPEC-018-16 | decision | Decide the durable boundary and target set; do not author them. |
| F-013 — “State that `coordination` owns Task graph change.” | PRODUCT-TASK-SPEC-018-18 | finding correction | Add the missing ownership wording to the active instruction file. |
| F-014 — “Decide each finding as `CLOSED` or `OPEN`.” | PRODUCT-TASK-SPEC-018-19 | independent finding-closure review | Issue an independent review disposition, not a design decision. |
| F-015 — “Route the accepted T19 result to T08.” | PRODUCT-TASK-SPEC-018-20 | graph coordination | Connect the accepted review result to the closure-synchronization owner. |
| F-016 — “Persist every explicit answer before advancing the cursor.” | PRODUCT-TASK-SPEC-019-01 | decision-ledger recording | Write each accepted answer into the Task-local ledger before continuing. |
| F-017 — “Amend the incomplete W019 boundary and downstream route.” | PRODUCT-TASK-SPEC-019-02 | graph coordination | Update the incomplete Work Item boundary and successor route. |
| F-018 — “Fix the corrected canonical target and downstream Investigation scope.” | PRODUCT-TASK-SPEC-019-03 | decision | Decide the corrected target and research scope without authoring them. |
| F-019 — “Persist T05 responsibility, dependency, outputs, excluded scope, and release condition.” | PRODUCT-TASK-SPEC-019-04 | Task materialization | Write the complete bounded contract for the new Investigation Task. |
| F-020 — “Reserve `PRODUCT-INV-SPEC-007` as the bounded Investigation output.” | PRODUCT-TASK-SPEC-019-04 | output allocation | Allocate the Investigation ID and output boundary without authoring the record. |
| F-021 — “Synchronize the W019 Task flow and Task Candidates with the materialized route.” | PRODUCT-TASK-SPEC-019-04 | graph coordination | Align duplicate Work Item graph projections; this is not lifecycle synchronization. |
| F-022 — “Replace the abstract immediate post-T05 reconciliation step with exact T06 and T07 routing.” | PRODUCT-TASK-SPEC-019-06 | graph coordination | Replace an abstract route with concrete Task owners and dependencies. |
| F-023 — “Treat W018 T11 J-001 as the accepted authority for MC-002.” | PRODUCT-TASK-SPEC-019-06 | fixed-input binding | Bind an earlier accepted judgment as authoritative input to the new route. |
| F-024 — “Consume W018 T11 J-001 without reopening the Task-type split.” | PRODUCT-TASK-SPEC-019-07 | decision input reuse | Use an accepted prior judgment as a fixed premise. |
| F-025 — “Fix Requirement disposition.” | PRODUCT-TASK-SPEC-019-07 | decision | Decide how the Requirement continues. |
| F-026 — “Fix W019 continuation or split disposition.” | PRODUCT-TASK-SPEC-019-07 | decision | Decide whether the Work Item continues or splits. |
| F-027 — “Fix canonical target decisions.” | PRODUCT-TASK-SPEC-019-07 | decision | Decide the exact canonical targets without writing them. |
| F-028 — “Fix shared-writer policy.” | PRODUCT-TASK-SPEC-019-07 | decision | Decide the writer-order policy without changing the graph. |
| F-029 — “State every required graph change without applying it.” | PRODUCT-TASK-SPEC-019-07 | decision | Record the required graph outcome while leaving execution to coordination. |
| F-030 — “Fix the checklist-authoring Work Item boundary.” | PRODUCT-TASK-SPEC-019-09 | decision | Decide the child Work Item responsibility boundary. |
| F-031 — “Fix the implementation Work Item boundary.” | PRODUCT-TASK-SPEC-019-09 | decision | Decide the separate implementation Work Item boundary. |
| F-032 — “Fix coarse release order between W019 and the two downstream Work Items.” | PRODUCT-TASK-SPEC-019-09 | decision | Decide the high-level release sequence without editing Work Items. |
| F-033 — “Record only coarse downstream routing in W019.” | PRODUCT-TASK-SPEC-019-10 | Work Item decomposition projection | Update the parent Work Item with child identities and high-level routing only. |
| F-034 — “Record the external W018 finding-closure gate for the later `task_authoring` writer and W019 integrated review.” | PRODUCT-TASK-SPEC-019-11 | release-route coordination | Persist the cross-Work-Item dependency gate for later writers and review. |
| F-035 — “Hand the completed route to T13 without authoring ADR content.” | PRODUCT-TASK-SPEC-019-12 | routing handoff | Make the completed routing result the fixed input to downstream coordination. |
| F-036 — “Consume the completed T12 ADR-routing ledger without reopening its decisions.” | PRODUCT-TASK-SPEC-019-13 | coordination input reuse | Use the accepted routing ledger as fixed graph-materialization input. |
| F-037 — “Serialize T14 through T17 in canonical authoring order.” | PRODUCT-TASK-SPEC-019-13 | writer-order coordination | Set dependencies so canonical writers run in a deterministic sequence. |
| F-038 — “Set ADR dependencies consistently:” | PRODUCT-TASK-SPEC-019-14 | ADR authoring | Write dependency metadata that matches the routed ADR boundaries. |
| F-039 — “Preserve PRODUCT-REQ-SPEC-007 unchanged when its motivating problem and Required Outcome remain accurate.” | PRODUCT-TASK-SPEC-019-15 | no-amendment authoring disposition | Conclude that the Requirement needs no edit and record that disposition. |
| F-040 — “Project PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017 into current normative rules.” | PRODUCT-TASK-SPEC-019-16 | Specification authoring | Encode accepted ADR decisions as current normative Specification text. |
| F-041 — “Register the new topic directly under `spec:product`.” | PRODUCT-TASK-SPEC-019-16 | Specification index authoring | Add the new Specification area to the PRODUCT parent topic map. |
| F-042 — “State that the dedicated validator Specification owns semantic validator behavior.” | PRODUCT-TASK-SPEC-019-17 | Specification authoring | Add an ownership rule and semantic reference to `task_authoring`. |
| F-043 — “Establish reviewer independence from T14 through T17 authoring.” | PRODUCT-TASK-SPEC-019-18 | review prerequisite verification | Confirm the reviewer did not author the reviewed outputs. |
| F-044 — “Trace T01, T03, T07, and T09 decisions through PRODUCT-INV-SPEC-007, T12 routing, and final canonical projection.” | PRODUCT-TASK-SPEC-019-18 | integrated review | Evaluate end-to-end decision traceability across the reviewed design. |
| F-045 — “Accept T18 directly when its verdict is `PASS`.” | PRODUCT-TASK-SPEC-019-19 | closure synchronization | Use the review verdict as closure authority; do not issue a new review verdict. |
| F-046 — “Accept an explicit user disposition only when the user classifies every named finding as a non-blocking workflow exception and accepts the reviewed semantic design.” | PRODUCT-TASK-SPEC-019-19 | closure synchronization exception | Propagate a user-owned exception into closure without independently closing findings. |
| F-047 — “Evaluate every W019 Completion Condition.” | PRODUCT-TASK-SPEC-019-19 | closure verification | Mechanically check every closure prerequisite before lifecycle propagation. |
| F-048 — “Record the current release order without modifying the graph.” | PRODUCT-TASK-SPEC-019-21 | decision-ledger recording | Persist the selected release order as a decision while leaving graph edits downstream. |
| F-049 — “Set W020 to `in_progress`.” | PRODUCT-TASK-SPEC-019-22 | lifecycle graph update | Apply the accepted early-release decision to the child Work Item lifecycle. |

### Scanned-Task ledger

| source Task ID | declared task type | scan result |
|---|---|---|
| PRODUCT-TASK-SPEC-018-01 | decision | F-001 |
| PRODUCT-TASK-SPEC-018-02 | decision | F-002 |
| PRODUCT-TASK-SPEC-018-03 | authoring | F-003 |
| PRODUCT-TASK-SPEC-018-04 | authoring | zero findings |
| PRODUCT-TASK-SPEC-018-05 | authoring | F-004 |
| PRODUCT-TASK-SPEC-018-06 | authoring | F-005 |
| PRODUCT-TASK-SPEC-018-07 | review | zero findings |
| PRODUCT-TASK-SPEC-018-08 | synchronization | F-006 |
| PRODUCT-TASK-SPEC-018-09 | coordination | F-007 |
| PRODUCT-TASK-SPEC-018-10 | investigation | zero findings |
| PRODUCT-TASK-SPEC-018-11 | decision | F-008 |
| PRODUCT-TASK-SPEC-018-12 | decision | zero findings |
| PRODUCT-TASK-SPEC-018-13 | coordination | F-009 |
| PRODUCT-TASK-SPEC-018-14 | authoring | F-010 |
| PRODUCT-TASK-SPEC-018-15 | coordination | F-011 |
| PRODUCT-TASK-SPEC-018-16 | decision | F-012 |
| PRODUCT-TASK-SPEC-018-17 | authoring | zero findings |
| PRODUCT-TASK-SPEC-018-18 | correction | F-013 |
| PRODUCT-TASK-SPEC-018-19 | review | F-014 |
| PRODUCT-TASK-SPEC-018-20 | coordination | F-015 |
| PRODUCT-TASK-SPEC-019-01 | decision | F-016 |
| PRODUCT-TASK-SPEC-019-02 | coordination | F-017 |
| PRODUCT-TASK-SPEC-019-03 | decision | F-018 |
| PRODUCT-TASK-SPEC-019-04 | coordination | F-019, F-020, F-021 |
| PRODUCT-TASK-SPEC-019-05 | investigation | zero findings |
| PRODUCT-TASK-SPEC-019-06 | coordination | F-022, F-023 |
| PRODUCT-TASK-SPEC-019-07 | decision | F-024 through F-029 |
| PRODUCT-TASK-SPEC-019-08 | coordination | zero findings |
| PRODUCT-TASK-SPEC-019-09 | decision | F-030 through F-032 |
| PRODUCT-TASK-SPEC-019-10 | work_item_decomposition | F-033 |
| PRODUCT-TASK-SPEC-019-11 | coordination | F-034 |
| PRODUCT-TASK-SPEC-019-12 | decision | F-035 |
| PRODUCT-TASK-SPEC-019-13 | coordination | F-036, F-037 |
| PRODUCT-TASK-SPEC-019-14 | authoring | F-038 |
| PRODUCT-TASK-SPEC-019-15 | authoring | F-039 |
| PRODUCT-TASK-SPEC-019-16 | authoring | F-040, F-041 |
| PRODUCT-TASK-SPEC-019-17 | authoring | F-042 |
| PRODUCT-TASK-SPEC-019-18 | review | F-043, F-044 |
| PRODUCT-TASK-SPEC-019-19 | synchronization | F-045 through F-047 |
| PRODUCT-TASK-SPEC-019-20 | coordination | zero findings |
| PRODUCT-TASK-SPEC-019-21 | decision | F-048 |
| PRODUCT-TASK-SPEC-019-22 | coordination | F-049 |

### Verification result

- Corpus coverage: 42 of 42 fixed Tasks scanned.
- Finding coverage: 49 candidates across 34 Tasks; 8 Tasks recorded with zero findings.
- Source validity: every finding cites one existing Task in PRODUCT-WORK-SPEC-018 or PRODUCT-WORK-SPEC-019.
- Section boundary: every exact phrase comes from the source Task's `## Work` section.
- Canonical boundary: no candidate was promoted to a canonical term or written into `skills/task-boundary-vocabulary/`.
- Write boundary: only this Task's status and Evidence were updated.

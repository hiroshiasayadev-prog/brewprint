# PRODUCT-TASK-SPEC-025-08: Extract boundary-vocabulary entries from W013–W017

- **id**: PRODUCT-TASK-SPEC-025-08
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-025
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-025-02
- **outputs**:
  - PRODUCT-TASK-SPEC-025-08

## Goal

Extract boundary-vocabulary log entries from every `status: done` Task under `product/records/tasks/spec/` whose parent Work Item is in the `PRODUCT-WORK-SPEC-013` through `PRODUCT-WORK-SPEC-017` range (~32 Tasks), and record them directly in this Task's own `## Evidence`, under the Investigation-Task lightweight Evidence exception (PRODUCT-TASK-SPEC-025-02).

## Work

- Fix this Task's corpus at start: the exact set of `status: done` Task IDs under `product/records/tasks/spec/` whose `work_item` falls in `PRODUCT-WORK-SPEC-013`–`PRODUCT-WORK-SPEC-017`.
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

- Every `status: done` Task in the `PRODUCT-WORK-SPEC-013`–`PRODUCT-WORK-SPEC-017` range at this Task's fixed cutoff has a scanned entry, including zero-finding Tasks.
- Every logged entry cites an exact, existing source Task ID.
- The accepted user judgment and reason for using the lightweight exception is recorded in this Task's own Evidence.

## Verification

- Confirm every entry cites a real, existing Task ID within the assigned range.
- Confirm no canonical term decision was made.
- Confirm the exception-use judgment and reason are recorded in Evidence.
- Confirm scanned-Task count matches the fixed corpus for this range.

## Evidence

### Execution boundary and lightweight-exception judgment

- Fixed cutoff: 2026-07-03, using the current contents of `product/records/tasks/spec/`.
- Fixed corpus: 32 Tasks with `status: done` whose `work_item` is `PRODUCT-WORK-SPEC-013` through `PRODUCT-WORK-SPEC-017`.
- Per-Work-Item count: W013 = 8, W014 = 2, W015 = 3, W016 = 8, W017 = 11.
- Excluded: `PRODUCT-TASK-SPEC-016-08`, `PRODUCT-TASK-SPEC-016-09`, and `PRODUCT-TASK-SPEC-016-10` had `status: not_started` at cutoff.
- Inference boundary: W013 through W015 Tasks have no `task_type` field. Their clusters were inferred from title, Work, and Done condition.
- Accepted judgment: PRODUCT-TASK-SPEC-025-01 D-006 permits this Task to use the lightweight Evidence exception.
- Reason: a formal Investigation record is disproportionate for bounded log extraction. It would reproduce the heavyweight route already rejected for this work.
- Rejected alternative: relabeling the Task as `authoring` would require an artifact outside this Task's scope.
- Downstream conclusion: PRODUCT-TASK-SPEC-025-04 verifies, reconciles, and clusters this log. T04 then authors confirmed entries or concludes that no further action is required.
- Access mode: DRMCP is non-operational. Filesystem authoring was the required fallback.

### Boundary-vocabulary findings

These are raw corpus findings. `vocabulary target` classifies the responsibility represented by the phrase. It does not select a canonical term.

| phrase (exact) | source Task ID | vocabulary target | effective meaning in context |
|---|---|---|---|
| ``Treat `product/records/spec/design-records/` as the authoritative source tree for the package.`` | PRODUCT-TASK-SPEC-013-01 | decision / define boundary | Declare the fixed source-of-authority boundary consumed by later package work. |
| `Distinguish package boundary from source-authoring findings.` | PRODUCT-TASK-SPEC-013-01 | investigation / design classification | Separate responsibility domains so audit findings are not generation instructions. |
| `Define the producer/consumer responsibility split.` | PRODUCT-TASK-SPEC-013-02 | decision / design | Fix normative ownership between package production and consumption. |
| `Record that the producer does not own persistent warning evidence.` | PRODUCT-TASK-SPEC-013-03 | define / author contract | Establish a normative exclusion. `record` does more than log an execution fact. |
| `Select the public generator entrypoint and its PRODUCT-owned placement.` | PRODUCT-TASK-SPEC-013-04 | decision | Choose the accepted command surface and implementation owner. |
| `Fix the authoritative source and repository-local destination.` | PRODUCT-TASK-SPEC-013-04 | decision / freeze contract | Make source and destination fixed downstream inputs. |
| `Define the minimum T06 verification cases.` | PRODUCT-TASK-SPEC-013-05 | design / verification contract | Specify the bounded acceptance cases for the implementation Task. |
| `Synchronize the parent Work Item relations and Evidence.` | PRODUCT-TASK-SPEC-013-05 | synchronization | Propagate the accepted contract into parent relations and evidence. |
| ``Add a dedicated release fixture module under `product/tests/tools/`.`` | PRODUCT-TASK-SPEC-013-07 | implementation / create | Implement a new fixture module. `add` stands in for creating code and tests. |
| `Consolidate accepted T06 and T07 generation, test, release, ignore, and independent-review evidence.` | PRODUCT-TASK-SPEC-013-08 | synchronization | Aggregate accepted evidence into one release-ready producer record. |
| ``Prepare producer handoff evidence for `DRMCP-REQ-MCP-003` through `DRMCP-WORK-MCP-002`.`` | PRODUCT-TASK-SPEC-013-08 | coordination / synchronization | Project the accepted producer boundary to the downstream owner. |
| `Route stale DRMCP contract claims to their owning Work Items.` | PRODUCT-TASK-SPEC-014-01 | coordination | Assign follow-up responsibility without correcting downstream artifacts. |
| `Apply required corrections from blocking or major findings.` | PRODUCT-TASK-SPEC-014-02 | correction | Repair the exact changes required by independent review findings. |
| ``Mark `PRODUCT-WORK-SPEC-014` done after its completion conditions are satisfied.`` | PRODUCT-TASK-SPEC-014-02 | lifecycle synchronization | Advance the parent lifecycle after its closure gate passes. |
| ``Update `DRMCP-TASK-MCP-001-02` with accepted child completion evidence and close the coordination Task.`` | PRODUCT-TASK-SPEC-014-02 | synchronization / close | Propagate child completion and close the upstream coordination Task. |
| `Classify each pointer as correct per-file ownership, correct graph ownership, known mismatch, additional stale pointer, wording-only ambiguity, or no-change.` | PRODUCT-TASK-SPEC-015-01 | investigation / classify | Categorize each pointer without changing its owner. |
| `Synchronize the hub T07 lifecycle only when its recorded gate is satisfied.` | PRODUCT-TASK-SPEC-015-01 | lifecycle synchronization | Propagate lifecycle state only after the predefined gate passes. |
| ``Move the local Topics column-shape row from `DRMCP-WORK-SPEC-002` to `DRMCP-WORK-SPEC-001`.`` | PRODUCT-TASK-SPEC-015-02 | authoring / synchronization | Update a canonical ownership pointer to the accepted owner. |
| `Normalize the W015 Impact Scope wording from pointer candidates to retained implementation owners.` | PRODUCT-TASK-SPEC-015-02 | correction / authoring | Align stale wording with the accepted owner disposition. |
| `Assess every W015 Completion Condition separately.` | PRODUCT-TASK-SPEC-015-03 | review / verification | Evaluate closure readiness condition by condition. |
| `Apply only scoped corrections required by review findings.` | PRODUCT-TASK-SPEC-015-03 | correction | Repair only the named review scope. |
| `After review acceptance, close W015 and this Task.` | PRODUCT-TASK-SPEC-015-03 | lifecycle synchronization | Advance both lifecycle records after independent acceptance. |
| `Resolve facts already fixed by accepted authority without asking the user.` | PRODUCT-TASK-SPEC-016-01 | investigation / decision preparation | Derive accepted facts rather than adopt a new decision. |
| `Select the first unblocked decision for T02.` | PRODUCT-TASK-SPEC-016-01 | coordination / decision sequencing | Choose the next decision-loop cursor, not the substantive outcome. |
| `Persist each explicit answer before advancing.` | PRODUCT-TASK-SPEC-016-02 | decision recording | Store each user judgment before moving the decision cursor. |
| `Route W017-owned questions to W017 without deciding them.` | PRODUCT-TASK-SPEC-016-04 | coordination | Transfer decision ownership without adopting the decision. |
| `Classify each item as no ADR, new ADR, amendment, or supersession.` | PRODUCT-TASK-SPEC-016-05 | ADR routing / coordination | Determine the durable-record route without authoring an ADR. |
| `Group only decisions that form one coherent durable choice.` | PRODUCT-TASK-SPEC-016-05 | ADR boundary partitioning | Partition decisions into bounded ADR units. |
| `Create, amend, or supersede only ADRs identified by T05.` | PRODUCT-TASK-SPEC-016-06 | authoring | Produce exactly the routed ADR set. |
| `Apply decided T02 and T04 semantics to the exact T03 target sections.` | PRODUCT-TASK-SPEC-016-07 | authoring / project decision | Write accepted decisions into canonical Specification sections. |
| `Reflect required ADR choices as current normative rules.` | PRODUCT-TASK-SPEC-016-07 | authoring / project decision | Translate ADR choices into current Specification text. |
| `Remove stale Task-contract wording within W016 scope.` | PRODUCT-TASK-SPEC-016-07 | correction | Delete normative text that contradicts the accepted contract. |
| ``During closure execution, synchronize only this Task and `PRODUCT-WORK-SPEC-016`.`` | PRODUCT-TASK-SPEC-016-11 | lifecycle synchronization | Update only the exact closure-owned lifecycle records. |
| `Resolve facts already fixed by accepted authority without asking the user.` | PRODUCT-TASK-SPEC-017-01 | investigation / decision preparation | Derive accepted provenance facts rather than adopt a new choice. |
| `Select the first unblocked decision for T02.` | PRODUCT-TASK-SPEC-017-01 | coordination / decision sequencing | Choose the next cursor in the decision loop. |
| `Persist each explicit answer before advancing.` | PRODUCT-TASK-SPEC-017-02 | decision recording | Store each source-relation judgment before moving the cursor. |
| `Route W016-owned Task-type questions to W016 without deciding them.` | PRODUCT-TASK-SPEC-017-04 | coordination | Transfer decision ownership and preserve the Work Item boundary. |
| `Classify each item as no ADR, new ADR, amendment, or supersession.` | PRODUCT-TASK-SPEC-017-05 | ADR routing / coordination | Determine the ADR disposition without authoring an ADR. |
| `Group only decisions that form one coherent durable choice.` | PRODUCT-TASK-SPEC-017-05 | ADR boundary partitioning | Partition provenance decisions into coherent ADR units. |
| `Create, amend, or supersede only ADRs identified by T05.` | PRODUCT-TASK-SPEC-017-06 | authoring | Produce only the routed ADR set. |
| `Apply decided T02 and T04 semantics to the exact T03 target sections.` | PRODUCT-TASK-SPEC-017-07 | authoring / project decision | Write accepted source-relation decisions into canonical sections. |
| `Reflect required ADR choices as current normative rules.` | PRODUCT-TASK-SPEC-017-07 | authoring / project decision | Translate accepted ADR choices into normative text. |
| `Replace Requirement-only provenance and persisted Requirement reverse membership in canonical Specifications.` | PRODUCT-TASK-SPEC-017-07 | correction / authoring | Remove the stale provenance model and author its replacement. |
| ``Record `PASS` or `NEEDS REVISION` with classified findings.`` | PRODUCT-TASK-SPEC-017-08 | review verdict | Issue the independent verdict and owned finding set. |
| ``Align `## Task flow` and `## Task Candidates` with the single integrated-review route.`` | PRODUCT-TASK-SPEC-017-09 | correction | Repair workflow projections to match the accepted review route. |
| ``Mark the separate W016 review route as not required and superseded by W017 T08. Keep lifecycle status unchanged.`` | PRODUCT-TASK-SPEC-017-09 | correction / graph projection | Correct the route description without a lifecycle transition. |
| `Classify each finding as closed or still open with exact evidence.` | PRODUCT-TASK-SPEC-017-10 | review / finding closure | Independently decide each named finding disposition. |
| `Route any newly exposed design decision back to the decision workflow.` | PRODUCT-TASK-SPEC-017-10 | coordination | Stop re-review and return unresolved judgment to its owner. |
| ``During closure execution, synchronize only this Task and `PRODUCT-WORK-SPEC-017`.`` | PRODUCT-TASK-SPEC-017-11 | lifecycle synchronization | Update only the exact source-relation closure records. |

### Scanned-Task ledger

| source Task ID | scan result |
|---|---|
| PRODUCT-TASK-SPEC-013-01 | 2 findings logged. |
| PRODUCT-TASK-SPEC-013-02 | 1 finding logged. |
| PRODUCT-TASK-SPEC-013-03 | 1 finding logged. |
| PRODUCT-TASK-SPEC-013-04 | 2 findings logged. |
| PRODUCT-TASK-SPEC-013-05 | 2 findings logged. |
| PRODUCT-TASK-SPEC-013-06 | 0 findings. Implementation verbs and completion language were explicit. |
| PRODUCT-TASK-SPEC-013-07 | 1 finding logged. |
| PRODUCT-TASK-SPEC-013-08 | 2 findings logged. |
| PRODUCT-TASK-SPEC-014-01 | 1 finding logged. |
| PRODUCT-TASK-SPEC-014-02 | 3 findings logged. |
| PRODUCT-TASK-SPEC-015-01 | 2 findings logged. |
| PRODUCT-TASK-SPEC-015-02 | 2 findings logged. |
| PRODUCT-TASK-SPEC-015-03 | 3 findings logged. |
| PRODUCT-TASK-SPEC-016-01 | 2 findings logged. |
| PRODUCT-TASK-SPEC-016-02 | 1 finding logged. |
| PRODUCT-TASK-SPEC-016-03 | 0 findings. The investigation responsibility used direct verbs. |
| PRODUCT-TASK-SPEC-016-04 | 1 finding logged. |
| PRODUCT-TASK-SPEC-016-05 | 2 findings logged. |
| PRODUCT-TASK-SPEC-016-06 | 1 finding logged. |
| PRODUCT-TASK-SPEC-016-07 | 3 findings logged. |
| PRODUCT-TASK-SPEC-016-11 | 1 finding logged. |
| PRODUCT-TASK-SPEC-017-01 | 2 findings logged. |
| PRODUCT-TASK-SPEC-017-02 | 1 finding logged. |
| PRODUCT-TASK-SPEC-017-03 | 0 findings. The investigation responsibility was explicit. |
| PRODUCT-TASK-SPEC-017-04 | 1 finding logged. |
| PRODUCT-TASK-SPEC-017-05 | 2 findings logged. |
| PRODUCT-TASK-SPEC-017-06 | 1 finding logged. |
| PRODUCT-TASK-SPEC-017-07 | 3 findings logged. |
| PRODUCT-TASK-SPEC-017-08 | 1 finding logged. |
| PRODUCT-TASK-SPEC-017-09 | 2 findings logged. |
| PRODUCT-TASK-SPEC-017-10 | 2 findings logged. |
| PRODUCT-TASK-SPEC-017-11 | 1 finding logged. |

### Verification result

- Scanned Tasks: 32 of 32 fixed-corpus Tasks.
- Logged findings: 49.
- Zero-finding scanned Tasks: 3.
- Every finding cites an existing Task ID inside W013 through W017.
- Every in-scope `status: done` Task has one scanned-ledger row.
- No frequency ranking was performed.
- No canonical term or definition was selected.
- No sibling Task, formal Investigation record, or `skills/task-boundary-vocabulary/` file was created or modified.

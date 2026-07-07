# PRODUCT-TASK-SPEC-025-10: Extract boundary-vocabulary entries from W020–W025

- **id**: PRODUCT-TASK-SPEC-025-10
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-025
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-025-02
- **outputs**:
  - PRODUCT-TASK-SPEC-025-10

## Goal

Extract boundary-vocabulary log entries from every `status: done` Task under `product/records/tasks/spec/` whose parent Work Item is in the `PRODUCT-WORK-SPEC-020` through `PRODUCT-WORK-SPEC-025` range (~35 Tasks), and record them directly in this Task's own `## Evidence`, under the Investigation-Task lightweight Evidence exception (PRODUCT-TASK-SPEC-025-02).

## Work

- Fix this Task's corpus at start: the exact set of `status: done` Task IDs under `product/records/tasks/spec/` whose `work_item` falls in `PRODUCT-WORK-SPEC-020`–`PRODUCT-WORK-SPEC-025`. This range includes `PRODUCT-WORK-SPEC-025` itself (the current Work Item); exclude this Work Item's own Tasks from the scanned corpus, since they are not yet a stable finished source at the time of scanning.
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

- Every `status: done` Task in the `PRODUCT-WORK-SPEC-020`–`PRODUCT-WORK-SPEC-025` range (excluding `PRODUCT-WORK-SPEC-025`'s own Tasks) at this Task's fixed cutoff has a scanned entry, including zero-finding Tasks.
- Every logged entry cites an exact, existing source Task ID.
- The accepted user judgment and reason for using the lightweight exception is recorded in this Task's own Evidence.

## Verification

- Confirm every entry cites a real, existing Task ID within the assigned range.
- Confirm no canonical term decision was made.
- Confirm the exception-use judgment and reason are recorded in Evidence.
- Confirm scanned-Task count matches the fixed corpus for this range.

## Evidence

### Lightweight Evidence exception

- Accepted user judgment: PRODUCT-TASK-SPEC-025-01 D-006 permits this `investigation` Task to record the bounded corpus result directly in its own Evidence.
- Reason: a separate formal Investigation record is disproportionate for phrase-log extraction and would recreate the heavyweight route rejected by the user.
- Downstream conclusion: PRODUCT-TASK-SPEC-025-04 depends on this Task and will verify, reconcile, cluster, and either project confirmed entries into the canonical vocabulary files or determine that no further action is required.

### Fixed corpus

- Cutoff: repository state observed on 2026-07-03 when this Task started.
- Included: every `status: done` Task owned by PRODUCT-WORK-SPEC-020 through PRODUCT-WORK-SPEC-024.
- Excluded: every PRODUCT-WORK-SPEC-025 Task, as required by this Task contract.
- Excluded: PRODUCT-TASK-SPEC-021-04 because its cutoff status is `blocked`.
- Fixed corpus size: 38 Tasks.

### Scanned-task ledger

| source Task ID | findings |
|---|---:|
| PRODUCT-TASK-SPEC-020-01 | 1 |
| PRODUCT-TASK-SPEC-020-02 | 0 |
| PRODUCT-TASK-SPEC-020-03 | 1 |
| PRODUCT-TASK-SPEC-020-04 | 0 |
| PRODUCT-TASK-SPEC-020-05 | 0 |
| PRODUCT-TASK-SPEC-020-06 | 0 |
| PRODUCT-TASK-SPEC-021-01 | 1 |
| PRODUCT-TASK-SPEC-021-02 | 1 |
| PRODUCT-TASK-SPEC-021-03 | 1 |
| PRODUCT-TASK-SPEC-021-05 | 1 |
| PRODUCT-TASK-SPEC-021-06 | 1 |
| PRODUCT-TASK-SPEC-021-07 | 1 |
| PRODUCT-TASK-SPEC-021-08 | 1 |
| PRODUCT-TASK-SPEC-021-09 | 1 |
| PRODUCT-TASK-SPEC-021-10 | 1 |
| PRODUCT-TASK-SPEC-021-11 | 1 |
| PRODUCT-TASK-SPEC-021-12 | 0 |
| PRODUCT-TASK-SPEC-021-13 | 1 |
| PRODUCT-TASK-SPEC-021-14 | 1 |
| PRODUCT-TASK-SPEC-021-16 | 1 |
| PRODUCT-TASK-SPEC-022-01 | 1 |
| PRODUCT-TASK-SPEC-022-02 | 1 |
| PRODUCT-TASK-SPEC-022-03 | 1 |
| PRODUCT-TASK-SPEC-022-04 | 0 |
| PRODUCT-TASK-SPEC-022-05 | 0 |
| PRODUCT-TASK-SPEC-023-01 | 1 |
| PRODUCT-TASK-SPEC-023-02 | 1 |
| PRODUCT-TASK-SPEC-023-03 | 1 |
| PRODUCT-TASK-SPEC-023-04 | 1 |
| PRODUCT-TASK-SPEC-023-05 | 1 |
| PRODUCT-TASK-SPEC-023-06 | 1 |
| PRODUCT-TASK-SPEC-023-07 | 1 |
| PRODUCT-TASK-SPEC-023-08 | 1 |
| PRODUCT-TASK-SPEC-023-09 | 1 |
| PRODUCT-TASK-SPEC-023-10 | 0 |
| PRODUCT-TASK-SPEC-023-11 | 0 |
| PRODUCT-TASK-SPEC-024-01 | 1 |
| PRODUCT-TASK-SPEC-024-02 | 2 |

### Boundary-vocabulary findings

The vocabulary target column classifies the responsibility expressed by the source phrase. It does not select a canonical term.

| phrase (exact) | source Task ID | vocabulary target | effective meaning in that context |
|---|---|---|---|
| `Materialize the W020 checklist contract, Investigation, authoring, review, and closure graph.` | PRODUCT-TASK-SPEC-020-01 | coordination | Create the required Tasks, dependencies, gates, and release order for W020. |
| `Produce one Investigation that verifies checklist coverage, authority alignment, placement effects, and authoring risks.` | PRODUCT-TASK-SPEC-020-03 | investigation | Investigate one bounded checklist-coverage question and record the result in an Investigation artifact. |
| `Materialize one initial W021 decision, Investigation, and post-Investigation coordination route.` | PRODUCT-TASK-SPEC-021-01 | coordination | Create the initial W021 Task graph and its dependency chain. |
| `Produce one implementation-boundary decision ledger for the temporary standalone validator.` | PRODUCT-TASK-SPEC-021-02 | decision | Decide the bounded implementation questions and persist their terminal outcomes in a ledger. |
| `Produce one bounded implementation impact Investigation for an executor-ready standalone validator graph.` | PRODUCT-TASK-SPEC-021-03 | investigation | Investigate repository seams, constraints, writers, and verification candidates without adopting implementation choices. |
| `Repair the W021 Task graph so PRODUCT work stops before app-local design and implementation.` | PRODUCT-TASK-SPEC-021-05 | coordination | Replace an invalid release route by blocking obsolete work and creating successor decision and coordination owners. |
| `Produce one bounded decision ledger for correcting W021 ownership and bootstrapping the standalone validator app-local design route.` | PRODUCT-TASK-SPEC-021-06 | decision | Decide the corrected PRODUCT/TRV ownership boundary and downstream route. |
| `Materialize the accepted PRODUCT conceptual-design, TRV namespace bootstrap, independent successor handoff, review, and closure graph.` | PRODUCT-TASK-SPEC-021-07 | coordination | Create the complete post-decision Task graph, dependencies, writers, review gate, handoff, and closure route. |
| ``Classify every T06 decision as `required`, `covered`, `not_required`, or `blocked`.`` | PRODUCT-TASK-SPEC-021-08 | decision / ADR routing | Decide the ADR disposition for each prior decision before canonical authoring. |
| `Materialize the exact ADR-authoring route from T08 and release the serialized canonical authoring chain.` | PRODUCT-TASK-SPEC-021-09 | coordination | Create required ADR writer Tasks and establish deterministic writer and release order. |
| `Activate the TRV app namespace and SPEC domain through one bounded Brewprint profile and namespace-overview authoring set.` | PRODUCT-TASK-SPEC-021-10 | authoring | Create and update canonical namespace, domain, layout, and overview Specifications so TRV becomes active. |
| `Project the T06 PRODUCT-owned semantic contract and corrected TRV ownership boundary into the canonical validator Specification.` | PRODUCT-TASK-SPEC-021-11 | authoring | Update the canonical Specification from already decided semantic and ownership inputs. |
| `Create the independent TRV app-local design Work Item selected by T06.` | PRODUCT-TASK-SPEC-021-13 | work_item_decomposition | Materialize one decided child Work Item boundary without creating its internal Task graph. |
| `Propagate the accepted W021 review and successor-handoff result into lifecycle, Evidence, relation, and closure state.` | PRODUCT-TASK-SPEC-021-14 | synchronization | Mechanically update Work Item lifecycle, Evidence, and relations from accepted prior results. |
| `Amend PRODUCT-ADR-SPEC-016 with the non-material TRV ownership clarification selected by corrected T08 B-001.` | PRODUCT-TASK-SPEC-021-16 | authoring | Modify one existing ADR from a fixed routed authoring input without reopening the decision. |
| `Fix one bounded contract for representing one already-created child Work Item as one parent-graph execution unit.` | PRODUCT-TASK-SPEC-022-01 | decision | Decide and freeze the `work_item_execution` Task contract. |
| ``Project the decided `work_item_execution` contract into one bounded canonical authority set.`` | PRODUCT-TASK-SPEC-022-02 | authoring | Write the decided contract into ADRs, Specifications, workflow support, and activation guidance. |
| ``Release the existing W020 checklist authoring owner after the `work_item_execution` authority becomes canonical.`` | PRODUCT-TASK-SPEC-022-03 | coordination | Remove a resolved blocker and add the dependency that permits the existing writer to proceed. |
| `Fix one bounded semantic contract for terminal Work Item and Task cancellation.` | PRODUCT-TASK-SPEC-023-01 | decision | Decide and freeze the cancellation lifecycle semantics. |
| ``Create one bounded Investigation of repository impact and conflicts for the decided `cancelled` lifecycle contract.`` | PRODUCT-TASK-SPEC-023-02 | investigation | Investigate direct lifecycle consumers, mismatches, owners, and writer conflicts in one formal Investigation. |
| ``Materialize one exact post-Investigation Task route for the decided `cancelled` lifecycle contract.`` | PRODUCT-TASK-SPEC-023-03 | coordination | Create only the required reconciliation decision and later graph owner after Investigation. |
| `Resolve the exact cancellation contract gaps identified by PRODUCT-INV-SPEC-010.` | PRODUCT-TASK-SPEC-023-04 | decision | Decide the unresolved body-readiness, propagation-owner, and validator-invocation questions. |
| `Materialize the exact ADR-routing, authoring, review, and closure route after T04 resolves the Investigation findings.` | PRODUCT-TASK-SPEC-023-05 | coordination | Create the next routing owners and defer exact writers until routing completes. |
| `Classify every terminal cancellation decision into one complete ADR route and coherent ADR boundary set.` | PRODUCT-TASK-SPEC-023-06 | decision / ADR routing | Decide ADR coverage, disposition, boundary partitioning, and downstream authoring boundaries. |
| `Materialize the exact ADR, canonical authoring, workflow-support, integrated review, and accepted-route closure graph from T06.` | PRODUCT-TASK-SPEC-023-07 | coordination | Create and order the exact writers, integrated review, and direct-PASS closure route. |
| `Create and amend the exact ADR set routed by T06 so every durable cancellation decision has historically honest authority.` | PRODUCT-TASK-SPEC-023-08 | authoring | Author one new ADR and amend the routed existing ADRs from fixed decisions. |
| `Project the accepted cancellation decisions and ADR authority into one coherent canonical and workflow-support contract.` | PRODUCT-TASK-SPEC-023-09 | authoring | Update current normative Specifications and workflow guidance from accepted decisions and ADR authority. |
| `Fix one bounded MVP contract for Work Item framing and its exact downstream authoring route.` | PRODUCT-TASK-SPEC-024-01 | decision | Decide and freeze the framing workflow contract and uniquely determined downstream route. |
| ``Activate framing in `prompt_chappy.md` before repository-persistent Work Item planning.`` | PRODUCT-TASK-SPEC-024-02 | authoring | Update the instruction source so the new framing workflow becomes mandatory at the selected entry point. |
| `Record verification Evidence and directly close PRODUCT-WORK-SPEC-024 when every Done condition is satisfied.` | PRODUCT-TASK-SPEC-024-02 | synchronization / lifecycle closure | Mechanically update the parent Work Item to `done` inside the authoring Task under the accepted simple-closure exception. |

### Verification result

- Scanned Tasks: 38 of 38 fixed corpus entries.
- Findings: 30.
- Zero-finding Tasks: 9.
- Every source Task ID exists in PRODUCT-WORK-SPEC-020 through PRODUCT-WORK-SPEC-024.
- PRODUCT-TASK-SPEC-021-04 was excluded because it was `blocked` at cutoff.
- PRODUCT-WORK-SPEC-025 Tasks were excluded from the source corpus.
- No frequency ranking was added.
- No canonical vocabulary term or definition was selected.
- No formal Investigation record was created.
- No sibling Task or `skills/task-boundary-vocabulary/` file was modified.
- DRMCP is non-operational. Filesystem authoring was the required fallback.

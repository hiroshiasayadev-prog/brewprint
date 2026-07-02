# PRODUCT-WORK-SPEC-022: Add work_item_execution Task type

- **id**: PRODUCT-WORK-SPEC-022
- **status**: done
- **date**: 2026-07-02
- **source_refs**:
  - PRODUCT-REQ-SPEC-005
  - PRODUCT-WORK-SPEC-018
  - PRODUCT-TASK-SPEC-018-11
  - PRODUCT-TASK-SPEC-018-14
  - PRODUCT-WORK-SPEC-020
  - PRODUCT-TASK-SPEC-020-04
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-010
- **impact_refs**:
  - PRODUCT-ADR-SPEC-004
  - PRODUCT-ADR-SPEC-005
  - PRODUCT-ADR-SPEC-010
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.authoring_standards.work_item_authoring
  - PRODUCT-WORK-SPEC-020
  - PRODUCT-TASK-SPEC-020-04
- **tasks**:
  - PRODUCT-TASK-SPEC-022-01
  - PRODUCT-TASK-SPEC-022-02
  - PRODUCT-TASK-SPEC-022-03
  - PRODUCT-TASK-SPEC-022-04
  - PRODUCT-TASK-SPEC-022-05

## Goal

Add `work_item_execution` as the Task type for representing one already-created child Work Item in a parent Work Item Task graph.

Keep `work_item_decomposition` as the owner of child Work Item creation and split.

Restore W020 checklist authoring with the expanded Task taxonomy.

## Boundary

This Work Item owns:

- the `work_item_execution` Task contract;
- one Task-to-child-Work-Item relation;
- child-status and Task-completion semantics;
- boundaries with `coordination`, `work_item_decomposition`, and `synchronization`;
- canonical ADR, Specification, and workflow-support projection;
- W020 checklist projection;
- one final integrated independent review;
- final lifecycle and Evidence synchronization.

This Work Item does not own:

- removal or replacement of `work_item_decomposition`;
- child Work Item internal Tasks, deliverables, decisions, procedures, or review Evidence;
- a direct Work Item-to-Work Item ownership relation;
- existing-record migration;
- validator implementation or DRMCP integration;
- production implementation;
- stage or commit work.

## Impact Scope

| target | impact |
|---|---|
| `PRODUCT-ADR-SPEC-004` | Add the eleventh closed Task type. |
| `PRODUCT-ADR-SPEC-005` | Define the execution, decomposition, coordination, and synchronization boundaries. |
| `PRODUCT-ADR-SPEC-010` | Add conditional child Work Item execution to design-convergence routing. |
| `spec:product.design_records.authoring_standards.task_authoring` | Add `work_item_ref`, the type contract, completion semantics, and adjacent boundaries. |
| `spec:product.design_records.authoring_standards.work_item_authoring` | Define parent hub and child ownership rules. |
| Design-convergence workflow skill | Add the execution handoff and companion authority. |
| `prompt_chappy.md` | Activate the new workflow companion. |
| W020 checklist artifact set | Add the new type and complete the remaining type projections. |

## Task flow

```text
PRODUCT-TASK-SPEC-022-01 contract decision
  -> PRODUCT-TASK-SPEC-022-02 canonical authority authoring
  -> PRODUCT-TASK-SPEC-022-03 W020 checklist release coordination
  -> PRODUCT-TASK-SPEC-020-04 checklist artifact authoring
  -> PRODUCT-TASK-SPEC-022-04 integrated independent review
     -> PASS: PRODUCT-TASK-SPEC-022-05 closure synchronization
     -> NEEDS REVISION: finding-specific coordination, correction, and independent closure review
```

T02 owns canonical authority authoring.
T03 releases the existing W020 authoring owner and performs no checklist writing.
W020 T04 remains the sole checklist writer.
W022 T04 follows T02 and W020 T04.
Correction and finding-closure review Tasks are created only after named findings exist.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `PRODUCT-TASK-SPEC-022-01` | `decision` | Fix the type, relation, completion, status, boundary, ADR, migration, and target contract. | none |
| `PRODUCT-TASK-SPEC-022-02` | `authoring` | Project the decided contract into canonical ADR, Specification, and workflow-support authority. | T01 |
| `PRODUCT-TASK-SPEC-022-03` | `coordination` | Release W020 T04 after canonical authority and remove duplicate checklist ownership from W022. | T02 |
| `PRODUCT-TASK-SPEC-022-04` | `review` | Independently review the final combined W022 design and W020 checklist projection. | T02 and W020 T04 |
| `PRODUCT-TASK-SPEC-022-05` | `synchronization` | Synchronize accepted lifecycle, Evidence, relations, and Work Item closure. | accepted T04 route |

## Completion Condition

- `work_item_execution` has one primary outcome and one completion judgment.
- `work_item_ref` represents exactly one already-created child Work Item.
- The referenced child Work Item must be `done` before the execution Task can be `done`.
- `work_item_decomposition` remains valid and non-overlapping.
- Parent Task content does not duplicate child-owned execution detail.
- ADR, Specification, workflow-support, and checklist projections agree.
- W020 checklist authoring is complete for all eleven canonical Task types.
- One final integrated independent review returns `PASS`, or every required finding is independently `CLOSED`.
- Lifecycle, Evidence, and relations express the accepted result.
- No implementation, stage, or commit work is performed.

## Evidence

- The user fixed the `work_item_execution` name and one-Task-to-one-child-Work-Item intent on 2026-07-02.
- The user selected a dedicated scalar `work_item_ref` relation and child completion as the Task completion boundary.
- Current canonical lifecycle has no `canceled` status. W022 does not add one.
- Existing `PRODUCT-INV-SPEC-006`, `PRODUCT-INV-SPEC-008`, W020 T04 blocker Evidence, and accepted Task authorities provide the bounded impact evidence used by T01.
- The initial eight-Task draft was untracked and unexecuted. It was replaced by this five-Task route before authoring.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
- T01 fixed D-001 through D-017 and completed with all decisions terminal.
- T02 authored the accepted ADR, Specification, workflow-support, and activation projections.
- T03 preserved W020 T04 as the sole checklist writer and released its authoring route.
- W020 T04 completed the eleven-type checklist artifact set.
- T04 returned `PASS` with no findings or direct regressions.
- T05 synchronized the direct PASS route and completed W022 closure.

### Completion Condition results

| condition | result |
|---|---|
| `work_item_execution` has one primary outcome and completion judgment. | PASS |
| `work_item_ref` represents exactly one already-created child Work Item. | PASS |
| The child must be `done` before the execution Task can be `done`. | PASS |
| `work_item_decomposition` remains valid and non-overlapping. | PASS |
| Parent Task content does not duplicate child-owned execution detail. | PASS |
| ADR, Specification, workflow-support, and checklist projections agree. | PASS |
| W020 authoring covers all eleven canonical Task types. | PASS |
| Final integrated independent review is accepted. | PASS |
| Lifecycle, Evidence, and relations express the accepted result. | PASS |
| No implementation, stage, or commit work occurred. | PASS |

### Closure state

- PRODUCT-TASK-SPEC-022-01 through PRODUCT-TASK-SPEC-022-05 are `done`.
- PRODUCT-TASK-SPEC-022-04 is the accepted review Evidence.
- PRODUCT-TASK-SPEC-022-05 is the closure synchronization Evidence.
- Work Item `tasks` and every Task `work_item` relation remain coherent.
- W020 full checklist review remains a separate responsibility of PRODUCT-TASK-SPEC-020-05.
- No canonical content, review verdict, finding set, dependency, or Task graph was changed during closure.
- No implementation, stage, or commit work was performed.

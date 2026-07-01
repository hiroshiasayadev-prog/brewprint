# PRODUCT-TASK-SPEC-016-02: Run REQ-005 interactive decision loop

- **id**: PRODUCT-TASK-SPEC-016-02
- **status**: done
- **date**: 2026-06-30
- **work_item**: PRODUCT-WORK-SPEC-016
- **source_requirement**: PRODUCT-REQ-SPEC-005
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-016-01
- **outputs**:
  - PRODUCT-TASK-SPEC-016-02

## Goal

Persist every unresolved REQ-005 decision one at a time.

Finish with each required decision decided, deferred, or blocked for a named reason.

## Work

- Resume from the current decision cursor.
- Ask exactly one decision per user turn.
- Check each answer against accepted Requirements and prior decisions.
- Persist each explicit answer before advancing.
- Update status, summary, dependencies, ADR route, canonical target, and cursor.
- Keep W017-owned source-relation decisions outside this loop.

## Done condition

- Every required decision has status `decided`, `deferred`, or validly `blocked`.
- Every explicit user answer is durably recorded.
- No more than one decision is `in_discussion`.
- No accepted constraint is reopened without an explicit contradiction.
- ADR routing and Specification targets are ready for downstream Tasks.

## Verification

- Inspect the scoped T02 diff after every persisted answer.
- Confirm the current decision ID and cursor match.
- Confirm no unrelated decision row changed.
- Confirm no decision is marked `recorded` before canonical synchronization.
- Confirm no production implementation, ADR authoring, Specification authoring, or independent review begins in T02.

## Evidence

### Bootstrap disposition

`BOOTSTRAP-001` was decided by the user on 2026-06-30.

- Use the current canonical Task metadata contract.
- Use `source_requirement` and `work_item`.
- Do not add an unaccepted primary Task type field.
- Do not require migration of this workflow Task.
- Treat the bootstrap choice as workflow authoring policy only.

### Loop state

Loop status: done

Current decision: none

Inventory source: PRODUCT-TASK-SPEC-016-01

### Decision confirmation loop

| ID | Topic | Status | Depends on | Decision summary | ADR route | Canonical target |
|---|---|---|---|---|---|---|
| D-001 | Primary Task type field name | decided | — | Use required persisted field `task_type`. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |
| D-002 | Allowed primary Task type closed set | decided | D-001 | Use the closed set `investigation`, `decision`, `authoring`, `implementation`, `review`, `correction`, `verification`, `coordination`, and `synchronization`. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-003 | Investigation type contract | decided | D-002 | Own creation or update of one Investigation record for one bounded research question. Completion follows the Investigation record's own completion requirements. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-004 | Decision type contract | decided | D-002 | Own one bounded decision ledger. Completion requires every owned decision to be `decided`, `deferred`, or validly `blocked` and the Task Done condition to be satisfied. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-005 | Authoring type contract | decided | D-002 | Own creation or content update of one bounded artifact set from already decided inputs. Completion follows the authored artifacts' own authoring and completion requirements and the Task Done condition. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-006 | Implementation type contract | decided | D-002 | Own one bounded implementation outcome. `## Implementation contract` is mandatory and maps each target to required change, acceptance criterion, and verification. Completion requires the Task Done condition and declared verification to pass. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-007 | Review type contract | decided | D-002 | Own one bounded independent verdict and finding set. Completion requires `PASS` or `NEEDS REVISION` plus complete finding evidence and the Task Done condition. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-008 | Correction type contract | decided | D-002 | Own correction of one bounded named finding set and its direct consistency effects. Completion requires the specified corrections, the Task Done condition, and declared verification; finding closure remains separate review work. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-009 | Verification type contract | decided | D-002 | Own one bounded acceptance gate with predefined objective checks. Completion requires all checks to be executed, expected and actual results recorded, and the overall result set to `PASS`, `FAIL`, or validly `BLOCKED`. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-010 | Coordination type contract | decided | D-002 | Own the parent Work Item overview of child Work Items and their responsibility boundaries. Child Work Items may be created before their detailed Task composition is known. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-011 | Synchronization type contract | decided | D-002 | Own propagation of an already accepted result into bounded lifecycle, Evidence, and relation state. Completion requires all specified records to express the same accepted result without stale state. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-012 | Decision and authoring boundary | decided | D-004, D-005 | `decision` ends when the choice and canonical target are fixed in the decision ledger; `authoring` begins when those fixed inputs are written into canonical artifacts. | candidate | `spec:product.design_records.authoring_standards.artifact_boundary` |
| D-013 | Review and verification boundary | decided | D-007, D-009 | Classify by primary outcome: independent verdict and semantic findings are `review`; predefined objective check results are `verification`. Review may use commands as evidence without requiring a separate verification Task. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-014 | Correction and finding-closure review boundary | decided | D-007, D-008 | `correction` owns the named fix and direct verification; a separate independent `review` Task decides each finding as `CLOSED` or `OPEN`. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-015 | Coordination and synchronization boundary | decided | D-010, D-011 | `coordination` defines or changes child Work Item overview and responsibility boundaries; `synchronization` only propagates already accepted results into lifecycle, Evidence, and relation state. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-016 | Judgment allowed inside implementation Tasks | decided | D-006 | Permit only locally observable-equivalent implementation choices that do not change accepted contract boundaries or acceptance criteria; contract-affecting choices return to `decision`. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-017 | Type alignment for Goal, Work, Done condition, and Verification | decided | D-002 through D-011 | Goal, Work, Done condition, and Verification must all serve one primary outcome matching `task_type`; supporting actions may not own a second outcome or completion judgment. | candidate | `spec:product.design_records.authoring_standards.task_authoring` |
| D-018 | Multi-file single-responsibility test | decided | D-002 | Multiple files remain one Task only when all changes serve one primary outcome, one completion judgment, and one acceptance or verification boundary without separate ownership or release decisions. | not_required | `spec:product.design_records.authoring_standards.task_authoring` |

### Recorded decisions

- D-001: The required persisted primary Task type field is `task_type`.
- D-002: The closed `task_type` set is `investigation`, `decision`, `authoring`, `implementation`, `review`, `correction`, `verification`, `coordination`, and `synchronization`.
- D-003: An `investigation` Task owns creation or update of one Investigation record for one bounded research question. Its completion judgment follows the Investigation record's own completion requirements. It does not adopt a design decision, synchronize ADR or Specification state, implement production changes, perform independent review, correct findings, or synchronize workflow lifecycle.
- D-004: A `decision` Task owns one bounded decision ledger. It may manage the decision inventory, ask one decision at a time, and record answers, dependencies, cursor, ADR route, and canonical targets. Completion requires every owned decision to be `decided`, `deferred`, or validly `blocked` and the Task's own Done condition to be satisfied. It does not create Investigation records, author ADRs or Specifications, implement production changes, perform independent review, correct or close findings, or perform final lifecycle synchronization.
- D-005: An `authoring` Task owns creation or content update of one bounded artifact set from already decided inputs. Completion follows the authored artifacts' own authoring and completion requirements and the Task's own Done condition. It may author ADRs, Specifications, Requirements, Work Items, Tasks, or equivalent design records. It does not make unresolved decisions, create Investigation records, implement production changes, perform independent review, correct named findings, close findings, or perform workflow lifecycle synchronization. Investigation-record authoring remains `investigation`; finding-driven artifact changes remain `correction`; accepted-state lifecycle or relation propagation remains subject to the `synchronization` contract.
- D-006: An `implementation` Task owns one bounded implementation outcome that realizes accepted design or Specification content. The Task may change production code and directly required focused tests or fixtures when they share one acceptance boundary. `## Implementation contract` is mandatory for this Task type. The section maps each target to required change, acceptance criterion, and verification. Completion requires the Task's Done condition and declared verification to pass. The Task does not contain unresolved design decisions, create Investigation, ADR, or Specification records, perform independent review, correct named findings, close findings, or coordinate or synchronize workflow state. Verification whose primary outcome is independent confirmation belongs to `verification`.
- D-007: A `review` Task owns one bounded independent verdict and finding set. Completion requires `PASS` or `NEEDS REVISION`, complete finding severity, evidence, target, and correction boundary, and the Task's own Done condition. The Task may compare accepted authority with reviewed artifacts and perform read-only inspection. It does not change reviewed artifacts, correct or close findings, adopt unresolved design decisions, author ADRs or Specifications, implement production changes, or synchronize workflow lifecycle. The reviewer must be independent of the authoring, implementation, or correction work under review. Multiple artifacts or files may share one review Task only when one common acceptance boundary is judged.
- D-008: A `correction` Task owns correction of one bounded named finding set and its direct consistency effects. Completion requires every specified correction to be implemented, the Task's own Done condition to be satisfied, and declared verification to pass. It does not decide finding closure, perform independent review, adopt new unresolved design decisions, make unrelated improvements, or synchronize workflow lifecycle. Correction remains with the owner of the affected artifact or implementation; responsibility must not be shifted to a generated or downstream artifact. Independent finding closure belongs to a subsequent `review` Task.
- D-009: A `verification` Task owns one bounded acceptance gate composed of predefined objective checks. It may execute commands, tests, schema checks, exact comparisons, or read-only inspections and record their evidence. Completion requires every declared check to be executed, expected and actual results to be recorded, and the overall result to be `PASS`, `FAIL`, or validly `BLOCKED`. It does not change the verified artifacts, make undefined design judgments, correct findings, issue an independent review verdict, or synchronize workflow lifecycle. Failure detection is recorded and routed to the responsible owner without repair. Multiple checks may share one Task only when they form one common acceptance gate.
- D-010: A `coordination` Task owns the parent Work Item overview of child Work Items and their responsibility boundaries. It creates or updates the minimal child Work Item inventory needed to show what each child owns. It does not duplicate each child Work Item's procedures, Task list, dependencies, release conditions, or next execution step. Child Work Items may be created from their overview and responsibility boundary before their detailed Task composition is known. Each child Work Item may then begin with its own investigation Task and derive only the Task types and sequence appropriate to that Work Item. Completion requires the necessary child Work Items to exist and their responsibilities to be distinguishable without material overlap or omission. The Task does not execute child-owned investigation, decision, authoring, implementation, review, correction, or verification work.
- D-011: A `synchronization` Task owns propagation of an already accepted result into a bounded set of Task or Work Item lifecycle, Evidence, completion-result, and relation state. Completion requires every specified record to express the same accepted result without stale state. The Task may only apply changes mechanically derivable from upstream accepted results. It does not make new design decisions, create substantive deliverables, decompose new child Work Items or Tasks, implement production changes, perform review, or correct findings. Any change requiring judgment belongs to a separate `decision` or `coordination` Task.
- D-012: A `decision` Task ends when the selected outcome, rationale, and canonical target are fixed in its decision ledger. An `authoring` Task begins when those already fixed inputs are written into canonical artifacts. A `decision` Task does not create or update canonical ADR or Specification body content. An `authoring` Task does not make new choices or resolve multiple valid interpretations. If authoring exposes an unresolved choice, authoring stops and a `decision` Task owns that choice before authoring resumes.
- D-013: `review` and `verification` are classified by primary outcome rather than method. A `review` Task owns an independent semantic verdict and may derive findings from accepted authority, inspected artifacts, and command or test evidence. A `verification` Task owns predefined objective check results and does not derive new semantic findings. A review may execute commands as evidence without requiring a separate verification Task. Verification is separated when the acceptance gate itself must be independently owned, reused, or aggregated.
- D-014: A `correction` Task owns implementation of the named finding set and direct verification of those changes. It does not mark findings `CLOSED`. A subsequent independent `review` Task evaluates each original finding as `CLOSED` or `OPEN`. That closure review may be limited to the named findings and their direct effects rather than repeating the full original review. The closure reviewer must be independent of the correction executor. Newly discovered issues are recorded as new findings. Correction and finding-closure review must not share one Task.
- D-015: `coordination` defines or changes the parent Work Item overview of child Work Items and their responsibility boundaries. `synchronization` only propagates already accepted results into bounded lifecycle, Evidence, completion-result, and relation state. Creating a child Work Item or changing its responsibility boundary is coordination. Updating status or Evidence from an accepted review or verification result is synchronization. If synchronization reveals that new decomposition or responsibility judgment is required, it stops and separate coordination work owns that decision. One Task must not combine both responsibilities.
- D-016: An `implementation` Task may make only local implementation choices whose externally observable result remains equivalent and whose alternatives all satisfy the accepted contract and `## Implementation contract` acceptance criteria. Examples include private names, helper decomposition, internal control flow, equivalent standard-library APIs, test setup, and local data structures within accepted architecture. Implementation stops and returns to `decision` when a choice changes public API or schema, responsibility or dependency boundaries, error categories or lifecycle, persistence or compatibility, security or performance guarantees, external dependencies, externally observable behavior, or the declared acceptance criteria.
- D-017: `## Goal`, `## Work`, `## Done condition`, and `## Verification` must all serve one primary outcome matching `task_type`. Goal declares that outcome, Work contains only actions needed to produce it, Done condition expresses its completion judgment, and Verification confirms that judgment without adding new acceptance requirements. Supporting actions associated with another type are allowed only when they do not own a second deliverable or completion judgment. `implementation` Tasks must additionally align these sections with `## Implementation contract`.
- D-018: Multiple files may remain in one Task only when every change is necessary for one `task_type`-aligned primary outcome, all changes share one completion judgment and one acceptance or verification boundary, no changed part is independently complete as another deliverable or verdict, and no file group requires a separate owner, release decision, or unresolved design decision. Production code with directly required focused tests or fixtures, multiple Specifications reflecting one accepted decision, and one named finding set spanning several locations may remain together. Implementation plus independent review, correction plus finding closure, independently releasable features, or changes with different primary Task types must be split.

### Loop closure

- All 18 decisions are `decided`.
- No decision remains `in_discussion`, `open`, or `blocked`.
- ADR routing and canonical targets are ready for downstream conflict investigation and routing Tasks.
- T03 and later Tasks have not started.

### Authority-resolved scope

- Exactly one primary type, one primary responsibility, one primary outcome, and one completion judgment are mandatory.
- Authoring plus independent review is prohibited.
- Correction plus independent finding closure is prohibited.
- Unresolved design decisions are prohibited inside implementation Tasks.
- Existing Task migration is outside W016.
- REQ-006 source-relation semantics remain owned by W017.

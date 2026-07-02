# PRODUCT-TASK-SPEC-019-07: Decide post-Investigation reconciliation

- **id**: PRODUCT-TASK-SPEC-019-07
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-06
- **outputs**:
  - PRODUCT-TASK-SPEC-019-07

## Goal

Produce one bounded post-Investigation reconciliation decision ledger for the remaining W019 semantic, identity, target, workflow-use, and writer-order judgments.

## Work

- Classify every investigated mismatch relevant to W019.
- Fix Requirement disposition.
- Fix W019 continuation or split disposition.
- Fix canonical target decisions.
- Fix shared-writer policy.
- State every required graph change without applying it.
- State the exact input required by ADR routing.
- Preserve T01 and T03 as completed historical checkpoints.
- Consume W018 T11 J-001 without reopening the Task-type split.
- Use an interactive one-question-at-a-time decision loop.
- Persist each accepted answer before advancing the cursor.

This Task does not change the Task graph.
This Task does not select an ADR disposition.
This Task does not author canonical artifacts or implementation work.

### Judgment inventory

| item | topic | lifecycle state | fixed input or decision boundary |
|---|---|---|---|
| R-001 | Validator ownership boundary | `decided` | Distinguish structural validation, standalone semantic evaluation, and out-of-scope DRMCP integration. |
| R-002 | Canonical validator contract target | `decided` | Select the exact canonical semantic ref and physical target candidate without authoring it. |
| R-003 | Historical decision disposition | `decided` | Reconcile T01 D-006 and T03 D-001 while preserving completed Tasks. |
| R-004 | MC-002 authority consumption and writer serialization | `decided` | Consume W018 T11 J-001 as fixed input and decide only dependency, serialization, and review-blocking effects. |
| R-005 | Authoring and pre-release usage boundary | `decided` | Decide caller ownership, result consumption, enforcement ownership, and required graph changes. |
| R-006 | Requirement, Work Item, and implementation-handoff disposition | `decided` | Decide Requirement and Work Item identity plus the separate implementation handoff boundary. |

### R-001: Validator ownership boundary

- Lifecycle state: `decided`.
- Decision: PRODUCT owns the semantic Task responsibility-boundary contract and the temporary standalone validator product boundary.
- Decision: The temporary standalone validator owns one-Task reading, checklist selection and injection, local LLM orchestration, criterion-result receipt, logical-AND aggregation, and semantic-result versus execution-failure separation.
- Decision: Current DRMCP remains responsible for structural parsing, structural validation, diagnostics, indexing, and current tool projections for accepted Design Record contracts.
- Decision: Future DRMCP integration remains outside W019 and requires a separate Requirement or Work Item.
- Reason: This split preserves the accepted standalone-tool boundary while keeping current DRMCP ownership limited to its existing structural contract responsibilities.
- Final mismatch classification: MC-001 is a `semantic_conflict` resolved by separating structural DRMCP validation from standalone semantic Task evaluation.
- Canonical target: The standalone semantic contract requires a PRODUCT canonical target selected by R-002. Current DRMCP artifacts are not targets.
- Required graph change: None selected by R-001. Downstream authoring and graph coordination must preserve this ownership split.
- Shared-writer or blocker effect: Any later `task_authoring` wording change must not assign standalone semantic execution to current DRMCP.
- ADR-routing handoff input: Route the ownership split without treating future DRMCP integration as part of W019.

### R-002: Canonical validator contract target

- Lifecycle state: `decided`.
- Decision: Create one dedicated PRODUCT Specification at `spec:product.responsibility_boundary_validator`.
- Physical target: `product/records/spec/responsibility-boundary-validator/index.md`.
- Parent target: `spec:product`.
- Decision: The dedicated Specification owns validator criteria composition, criterion-result semantics, rationale requirements, logical-AND aggregation, structural-precondition boundary, execution-failure separation, and the temporary standalone validator boundary.
- Decision: `spec:product.design_records.authoring_standards.task_authoring` continues to own the evaluated Task contract.
- Decision: `task_authoring` receives only a narrow workflow rule that runs this validator after Task authoring. It does not own validator behavior.
- Reason: The validator is a standalone PRODUCT concern but is not part of current DRMCP or the Design Records Task document contract itself.
- Final mismatch classification: The missing canonical validator target is `workflow_graph_drift`, not a reason to extend current DRMCP ownership.
- Required graph change: Later coordination must materialize authoring for the dedicated Specification, its parent topic registration, and the narrow `task_authoring` relation.
- Shared-writer or blocker effect: Any `task_authoring` update must follow the W018 canonical repair and preserve the dedicated validator target.
- ADR-routing handoff input: Treat the dedicated validator Specification and the narrow `task_authoring` usage relation as separate authoring boundaries.

### R-003: Historical decision disposition

- Lifecycle state: `decided`.
- Decision: Preserve T01 D-006 unchanged as historical decision Evidence.
- Decision: Classify the T01 D-006 “MCP” wording as `stale_representation`.
- Decision: Use T03 D-001 as the terminal corrected product-boundary decision for the temporary standalone validator.
- Decision: Downstream canonical authoring consumes T03 D-001 together with R-001 and R-002.
- Reason: T03 corrects the product boundary without invalidating the remaining T01 semantic decisions.
- Canonical target: `spec:product.responsibility_boundary_validator` and the narrow `task_authoring` usage relation selected by R-002.
- Required graph change: None from the historical Task itself. Later authoring must use the corrected boundary.
- Shared-writer or blocker effect: T01 is not a writable target and does not block downstream work once the corrected input is explicit.
- ADR-routing handoff input: Preserve T01 as historical Evidence and route current authority from T03 plus R-001 and R-002.

### R-004: MC-002 authority consumption and writer serialization

- Lifecycle state: `decided`.
- Fixed authority: W018 T11 J-001 assigns parent-to-child Work Item decomposition to `work_item_decomposition`.
- Fixed authority: W018 T11 J-001 assigns Task-graph changes to `coordination`.
- Decision: W019 consumes the W018 Task-type split as fixed authority and does not reopen it.
- Decision: W019 may continue dedicated validator Specification design, checklist design, validator contract design, downstream implementation planning, and standalone implementation before W018 completes.
- Decision: Only the W019 write to `spec:product.design_records.authoring_standards.task_authoring` waits for accepted W018 canonical repair and review.
- Decision: W019 integrated review remains blocked until the W019 `task_authoring` usage rule is authored after W018 acceptance.
- Decision: W019 graph coordination may be materialized earlier when it records this shared-writer dependency and review blocker explicitly.
- Reason: This isolates one shared-writer conflict without blocking independently completable validator design and implementation work.
- Final mismatch classification: MC-002 is `consistent_refinement` for W019 ownership, with one bounded shared-writer serialization constraint inherited from W018.
- Canonical target: Dedicated validator authoring targets are independent. The shared target is `spec:product.design_records.authoring_standards.task_authoring`.
- Required graph change: Later W019 coordination must serialize the `task_authoring` writer after accepted W018 repair and keep integrated review blocked until that writer completes.
- Shared-writer or blocker effect: W018 accepted canonical repair and review is the release gate only for the W019 `task_authoring` writer and final integrated review, not for the rest of W019.
- ADR-routing handoff input: Preserve the fixed Task-type split and route only the local shared-writer order and review blocker.

### R-005: Authoring and pre-release usage boundary

- Lifecycle state: `decided`.
- Decision: Run the validator once immediately after Task authoring and once after Task execution when final Evidence has been written.
- Decision: Both validation points consume the same validator and the same criterion-level result contract.
- Decision: The validator reports criterion verdicts, rationale, overall logical-AND result, and execution failure separately. It does not own workflow enforcement.
- Decision: A fully compliant result permits the caller to continue its normal route.
- Decision: Any semantic violation routes to explicit human judgment asking whether the violation is accepted for that Task.
- Decision: A human-accepted violation must preserve the violated criteria, rationale, acceptance decision, and acceptance reason as Evidence. A rejected violation routes back to Task correction or Task-boundary reconsideration.
- Caller ownership: The authoring workflow owns the post-authoring validation call. The Task completion or release workflow owns the post-Evidence validation call.
- Reason: The first check catches an invalid responsibility boundary before execution, while the second checks the completed Task including its actual Evidence and outcome.
- Final mismatch classification: The missing two-point usage and exception route is `workflow_graph_drift`.
- Required graph change: Later coordination must add one post-authoring validation gate, one post-Evidence validation gate, and a human exception-decision branch for semantic violations.
- Shared-writer or blocker effect: The post-authoring gate depends on the delayed `task_authoring` usage-rule writer selected by R-004. Validator contract and implementation may proceed independently.
- ADR-routing handoff input: Route one shared validator contract, two caller-owned invocation points, and one human-owned violation-acceptance branch.

### R-006: Requirement, Work Item, and implementation-handoff disposition

- Lifecycle state: `decided`.
- Decision: Amend `PRODUCT-REQ-SPEC-007` to require validation immediately after Task authoring and after final Evidence is written.
- Decision: Amend `PRODUCT-REQ-SPEC-007` to require explicit human acceptance or rejection of any reported semantic violation.
- Decision: Preserve `PRODUCT-REQ-SPEC-005` unchanged because the typed single-responsibility Task contract itself is not altered.
- Decision: Continue W019 as one design Work Item through Requirement reconciliation, validator-contract authoring, canonical Specification projection, integrated review, and closure.
- Decision: Route temporary standalone validator implementation to one separate downstream Work Item.
- Decision: Defer the implementation app namespace, source ownership, code placement, concrete model, provider, runtime, retry policy, and implementation error taxonomy to that downstream Work Item.
- Reason: W019 owns the reusable product contract and workflow-use semantics. Concrete delivery has a distinct implementation outcome and completion judgment.
- Final mismatch classification: Requirement and Work Item identity are a `consistent_refinement` of the accepted semantic-validation need.
- Canonical targets: Amend `PRODUCT-REQ-SPEC-007`; preserve `PRODUCT-REQ-SPEC-005`; author `spec:product.responsibility_boundary_validator`; add only the narrow usage relation to `spec:product.design_records.authoring_standards.task_authoring`.
- Required graph change: Later coordination must materialize bounded Requirement and Specification authoring, shared-writer serialization, integrated review, and the separate downstream implementation Work Item route. T07 does not perform those changes.
- Shared-writer or blocker effect: Only the `task_authoring` writer and final W019 integrated review wait for the accepted W018 repair identified by R-004.
- ADR-routing handoff input: Preserve REQ-005, amend REQ-007, continue W019, and keep implementation plus its exact application boundary in a separate downstream Work Item.

### Final reconciliation summary

| subject | final classification | disposition |
|---|---|---|
| MC-001 validator ownership overlap | `semantic_conflict` | Split PRODUCT semantic ownership, standalone execution, current DRMCP structural ownership, and future integration. |
| Missing canonical validator target | `workflow_graph_drift` | Create `spec:product.responsibility_boundary_validator`. |
| T01 D-006 historical “MCP” wording | `stale_representation` | Preserve T01 and use T03 plus R-001 and R-002 as current authority. |
| MC-002 W018 authority consumption | `consistent_refinement` | Consume the Task-type split and serialize only the shared `task_authoring` writer and final review. |
| Missing two-point invocation and exception route | `workflow_graph_drift` | Validate post-authoring and post-Evidence; route violations to explicit human judgment. |
| Requirement and Work Item identity | `consistent_refinement` | Amend REQ-007, preserve REQ-005, continue W019, and split implementation into a downstream Work Item. |

Required graph changes remain unperformed.
ADR routing remains a downstream responsibility.
No ADR disposition is selected by T07.

### ADR-routing handoff

The terminal ledger must provide these inputs without selecting an ADR disposition:

- final mismatch classifications;
- terminal R-001 through R-006 states and decisions;
- exact Requirement and Work Item dispositions;
- exact canonical semantic refs and physical authoring targets;
- required graph changes and release blockers;
- fixed shared-writer order and W018 dependency boundary;
- preserved historical decision treatment;
- validator ownership and workflow-use boundaries.

## Done condition

- R-001 through R-006 are each `decided`, `deferred`, or validly `blocked`.
- Every relevant mismatch has one final classification.
- Requirement and W019 identity dispositions are explicit.
- Canonical targets are exact or have named blockers.
- Shared-writer policy is fixed.
- Every required graph change is stated but not performed.
- MC-002 is not re-decided.
- ADR-routing inputs are complete.
- No ADR disposition is selected.
- No graph or canonical artifact is changed.

## Verification

- Confirm exactly six reconciliation items exist.
- Confirm each item has a stable ID and lifecycle state.
- Confirm MC-002 uses W018 T11 J-001 as fixed input.
- Confirm no item asks whether `coordination` owns Task-graph changes.
- Confirm Requirement, Work Item, target, graph, and shared-writer dispositions are covered.
- Confirm ADR routing remains a later responsibility.
- Confirm no ADR, Specification, Work Item, graph, review, closure, or implementation work occurred.

## Evidence

### Loop state

- Status: `decision_complete`.
- Current item: none.
- Terminal items: R-001 through R-006.
- Open items: none.
- Blocked items: none.

### Direct inputs

- `PRODUCT-TASK-SPEC-019-01`
- `PRODUCT-TASK-SPEC-019-03`
- `PRODUCT-TASK-SPEC-019-05`
- `PRODUCT-INV-SPEC-007`
- `PRODUCT-TASK-SPEC-018-11`, `### J-001 decision`
- `PRODUCT-REQ-SPEC-007`
- `PRODUCT-REQ-SPEC-005`
- `spec:product.design_records.authoring_standards.task_authoring`

- T07 was materialized by `PRODUCT-TASK-SPEC-019-06`.
- R-001 accepted the four-part ownership split.
- R-002 selected `spec:product.responsibility_boundary_validator` as the dedicated PRODUCT target.
- R-003 preserved T01 D-006 as stale historical wording and selected T03 D-001 as current authority.
- Current W018 state: T14 is done; T15 through T20 are not materialized.
- R-004 limited the W018 wait to the shared `task_authoring` writer and W019 integrated review.
- R-005 selected post-authoring and post-Evidence validation with human violation acceptance.
- R-006 amended the REQ-007 contract boundary, preserved REQ-005, continued W019, and separated implementation into a downstream Work Item.
- R-001 through R-006 are terminal.
- The decision loop is complete.
- Required graph changes are recorded but unperformed.
- No ADR routing, graph change, canonical authoring, review, closure, implementation, stage, or commit was performed.

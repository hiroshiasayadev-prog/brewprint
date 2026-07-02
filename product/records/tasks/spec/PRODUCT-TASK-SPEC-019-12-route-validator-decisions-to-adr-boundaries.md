# PRODUCT-TASK-SPEC-019-12: Route validator decisions to ADR boundaries

- **id**: PRODUCT-TASK-SPEC-019-12
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-019-11
- **outputs**:
  - PRODUCT-TASK-SPEC-019-12

## Goal

Produce one complete ADR-routing ledger for the terminal W019 validator-design and downstream-boundary decisions.

## Work

- Read T01, T03, T07, T09, and the concluded W019 Investigation.
- Inventory every terminal decision that requires an ADR route.
- Classify each decision as `required`, `covered`, `not_required`, or `blocked`.
- Check accepted non-superseded ADR coverage.
- Partition required durable decisions into coherent ADR boundaries.
- Select `create`, `amend`, `reuse`, or `supersede` for every required or covered boundary.
- Name exact affected Requirement, Specification, Work Item, and workflow targets.
- Record exact blockers when a route cannot complete.
- Hand the completed route to T13 without authoring ADR content.

This Task must not:

- author or amend ADR body content;
- amend PRODUCT-REQ-SPEC-007;
- create the validator Specification;
- modify `task_authoring`;
- create downstream authoring, review, or closure Tasks;
- perform checklist authoring, implementation, review, correction, synchronization, stage, or commit work.

### Routing result

- Status: `routing_complete`.
- Required new ADRs: three.
- Existing ADR amendments: none.
- Existing ADR supersessions: none.
- Blocked decisions: none.
- Next owner: PRODUCT-TASK-SPEC-019-13.

### Decision routing ledger

| source decision | routing outcome | ADR disposition | ADR boundary or existing ADR | reason | canonical targets |
|---|---|---|---|---|---|
| T01 D-001 | `required` | `create` | B-001 / PRODUCT-ADR-SPEC-015 | Overall and criterion-level binary result semantics constrain every implementation and caller. | `spec:product.responsibility_boundary_validator` |
| T01 D-002 | `required` | `create` | B-001 / PRODUCT-ADR-SPEC-015 | Criterion-level rationale is part of the observable semantic judgment contract. | `spec:product.responsibility_boundary_validator` |
| T01 D-003 | `required` | `create` | B-001 / PRODUCT-ADR-SPEC-015 | The Task-only evidence boundary determines what the validator may claim. | `spec:product.responsibility_boundary_validator` |
| T01 D-004 | `required` | `create` | B-001 / PRODUCT-ADR-SPEC-015 | Automatic checklist selection from declared `task_type` is a durable validator behavior. | `spec:product.responsibility_boundary_validator` |
| T01 D-005 | `required` | `create` | B-001 / PRODUCT-ADR-SPEC-015 | Common plus type-specific criterion composition constrains checklist authoring and implementation. | `spec:product.responsibility_boundary_validator`; PRODUCT-WORK-SPEC-020 |
| T01 D-006 | `not_required` | none | none | Its “MCP” wording is preserved historical Evidence and is replaced as current authority by T03 D-001, T07 R-001, and T07 R-002. | No write to T01; current state reaches PRODUCT-ADR-SPEC-016 and the validator Specification. |
| T01 D-007 | `required` | `create` | B-001 / PRODUCT-ADR-SPEC-015 | Logical-AND aggregation is a deterministic product-level result rule. | `spec:product.responsibility_boundary_validator` |
| T01 D-008 | `not_required` | none | none | Section references, optional excerpts, and exclusion of line numbers are bounded result-shape details without an independently durable trade-off. | `spec:product.responsibility_boundary_validator` |
| T01 D-009 | `required` | `create` | B-001 / PRODUCT-ADR-SPEC-015 | Structural precondition failure versus semantic non-compliance is a durable claim boundary. | `spec:product.responsibility_boundary_validator` |
| T01 D-010 | `required` | `create` | B-001 / PRODUCT-ADR-SPEC-015 | Execution failure must remain distinct from semantic judgment across runtimes and implementations. | `spec:product.responsibility_boundary_validator`; PRODUCT-WORK-SPEC-021 |
| T01 D-011 | `not_required` | none | none | Omission of checklist revision and stable criterion IDs from the external result is a bounded contract-detail choice. | `spec:product.responsibility_boundary_validator` |
| T01 D-012 | `required` | `create` | B-003 / PRODUCT-ADR-SPEC-017 | Invocation context and enforcement ownership constrain authoring and completion workflows. | PRODUCT-REQ-SPEC-007; `spec:product.responsibility_boundary_validator`; `spec:product.design_records.authoring_standards.task_authoring` |
| T03 D-001 | `required` | `create` | B-002 / PRODUCT-ADR-SPEC-016 | The temporary standalone shape and exclusion of current DRMCP integration establish a durable ownership boundary. | `spec:product.responsibility_boundary_validator`; PRODUCT-WORK-SPEC-021 |
| T07 R-001 | `required` | `create` | B-002 / PRODUCT-ADR-SPEC-016 | PRODUCT, standalone validator, current DRMCP, and future integration require explicit non-overlapping ownership. | `spec:product.responsibility_boundary_validator`; narrow ownership clarification in `spec:product.design_records.authoring_standards.task_authoring` |
| T07 R-002 | `required` | `create` | B-002 / PRODUCT-ADR-SPEC-016 | A dedicated PRODUCT Specification is the durable canonical owner of validator semantics. | `spec:product.responsibility_boundary_validator`; parent topic registration under `spec:product` |
| T07 R-003 | `covered` | `reuse` | PRODUCT-ADR-SPEC-006 and PRODUCT-ADR-SPEC-014 | Existing authority already preserves decision checkpoints and requires append-only successor work without rewriting completed Tasks. | No write to T01 or T03; downstream records cite them as historical inputs. |
| T07 R-004 | `covered` | `reuse` | PRODUCT-ADR-SPEC-012 | Existing authority already requires deterministic shared-writer order and one final integrated review after the final writer. | W019 Task graph; `spec:product.design_records.authoring_standards.task_authoring` writer dependency |
| T07 R-005 | `required` | `create` | B-003 / PRODUCT-ADR-SPEC-017 | Two mandatory invocation points and human-owned violation acceptance are durable workflow and exception-policy choices. | PRODUCT-REQ-SPEC-007; `spec:product.responsibility_boundary_validator`; `spec:product.design_records.authoring_standards.task_authoring` |
| T07 R-006 | `required` | `create` plus existing-policy reuse | B-003 / PRODUCT-ADR-SPEC-017; PRODUCT-ADR-SPEC-009 and PRODUCT-ADR-SPEC-011 reused | Mandatory two-point validation and human acceptance require B-003. Requirement continuity, W019 continuation, and separate implementation completion are already governed by ADR-009 and ADR-011. | PRODUCT-REQ-SPEC-007; PRODUCT-WORK-SPEC-019; PRODUCT-WORK-SPEC-021 |
| T09 D-001 | `covered` | `reuse` | PRODUCT-ADR-SPEC-005, PRODUCT-ADR-SPEC-009, PRODUCT-ADR-SPEC-011, and PRODUCT-ADR-SPEC-012 | Independent outcomes, completion judgments, release timing, and final reviews already require separate Work Items for checklist authoring and implementation. | PRODUCT-WORK-SPEC-020; PRODUCT-WORK-SPEC-021 |

### B-001: Semantic Task responsibility-validation contract

- ADR ID: PRODUCT-ADR-SPEC-015.
- Disposition: `create`.
- Proposed title: Define Task-local semantic responsibility validation semantics.
- Included decisions: T01 D-001 through D-005, D-007, D-009, and D-010.
- Bounded question: What evidence may one Task responsibility validator evaluate, what criterion results must it produce, and how are compliance, precondition failure, and execution failure distinguished?
- Cohesion reason: These decisions jointly define one observable semantic evaluation contract. Changing one independently can invalidate the meaning of the others.
- Excluded detail: Exact checklist wording, criterion identifiers, external response field names, physical storage format, model, provider, runtime, timeout, and retry policy.
- Required alternatives to preserve:
  - Task-local evidence versus external-context inference;
  - deterministic criterion aggregation versus free-form overall model judgment;
  - semantic false versus structural or execution failure;
  - automatic checklist selection versus caller-selected criteria.
- Consequence targets:
  - `spec:product.responsibility_boundary_validator`;
  - PRODUCT-WORK-SPEC-020 checklist contract;
  - PRODUCT-WORK-SPEC-021 implementation acceptance boundary.
- Dependency candidates: PRODUCT-ADR-SPEC-004 and PRODUCT-ADR-SPEC-005.
- Authoring handoff: T13 must create one bounded ADR-authoring owner for PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017. The writer authors this ADR before B-002 and B-003 dependent text.

### B-002: Temporary standalone ownership and integration boundary

- ADR ID: PRODUCT-ADR-SPEC-016.
- Disposition: `create`.
- Proposed title: Separate temporary semantic Task validation from current DRMCP.
- Included decisions: T03 D-001; T07 R-001 and R-002. T01 D-006 remains historical input only.
- Bounded question: Which owner defines and executes temporary semantic Task validation, what remains owned by current DRMCP, and where is future integration decided?
- Cohesion reason: Product shape, semantic ownership, orchestration ownership, structural DRMCP ownership, and dedicated canonical placement form one ownership architecture.
- Required ownership split:
  - PRODUCT owns the semantic contract and temporary product boundary;
  - the standalone validator owns Task reading, checklist selection and injection, local-model orchestration, criterion-result receipt, and deterministic aggregation;
  - current DRMCP retains parsing, structural validation, diagnostics, indexing, and current tool projections;
  - future DRMCP integration requires a separate Requirement or Work Item.
- Excluded detail: Exact implementation app namespace, source path, executable interface, model, runtime, and future integration design.
- Consequence targets:
  - `spec:product.responsibility_boundary_validator`;
  - parent topic registration under `spec:product`;
  - narrow ownership wording in `spec:product.design_records.authoring_standards.task_authoring`;
  - PRODUCT-WORK-SPEC-021.
- Dependency candidates: PRODUCT-ADR-SPEC-001 and PRODUCT-ADR-SPEC-015.
- Authoring handoff: T13 must preserve PRODUCT-ADR-SPEC-015 semantics and author PRODUCT-ADR-SPEC-016 as a separate independently supersedable ownership boundary.

### B-003: Invocation timing and human-owned violation exceptions

- ADR ID: PRODUCT-ADR-SPEC-017.
- Disposition: `create`.
- Proposed title: Validate Task responsibility after authoring and final Evidence with human-owned exceptions.
- Included decisions: T01 D-012; T07 R-005; the validation and human-acceptance portions of T07 R-006.
- Bounded question: When must the validator run, which workflow consumes its result, and who decides whether a semantic violation is accepted?
- Cohesion reason: Invocation timing, caller enforcement, violation handling, and required exception Evidence together define one workflow policy. Separating them could create a gate without an owner or an exception without an audit trail.
- Required workflow boundary:
  - run once immediately after Task authoring;
  - run once after final Evidence is written;
  - use the same validator and result contract at both points;
  - keep workflow enforcement outside the validator;
  - route semantic violations to explicit human acceptance or rejection;
  - preserve violated criteria, rationale, acceptance decision, and acceptance reason when accepted;
  - route rejected violations to Task correction or responsibility-boundary reconsideration.
- Excluded detail: Concrete UI, CLI, transport, identity mechanism, notification channel, or implementation retry behavior.
- Consequence targets:
  - PRODUCT-REQ-SPEC-007 amendment;
  - `spec:product.responsibility_boundary_validator`;
  - narrow usage rule in `spec:product.design_records.authoring_standards.task_authoring`;
  - downstream authoring and completion workflow integration.
- Dependency candidates: PRODUCT-ADR-SPEC-009, PRODUCT-ADR-SPEC-015, and PRODUCT-ADR-SPEC-016.
- Authoring handoff: T13 must place the `task_authoring` projection after independent W018 finding closure and place W019 integrated review after that final writer.

### Existing ADR reuse summary

| existing ADR | reused authority | affected decisions |
|---|---|---|
| PRODUCT-ADR-SPEC-005 | Split independent outcomes, owners, completion judgments, and release decisions. | T09 D-001 |
| PRODUCT-ADR-SPEC-006 | Preserve decision Tasks as workflow checkpoints rather than canonical design state. | T07 R-003 |
| PRODUCT-ADR-SPEC-009 | Keep design convergence and production implementation in separate completion boundaries. | T07 R-006; T09 D-001 |
| PRODUCT-ADR-SPEC-011 | Preserve Requirement and Work Item identity when the same acceptance and completion meaning remains; split independently completable scope. | T07 R-006; T09 D-001 |
| PRODUCT-ADR-SPEC-012 | Serialize shared writers and review the final combined state once per Work Item. | T07 R-004; T09 D-001 |
| PRODUCT-ADR-SPEC-014 | Preserve completed Tasks through append-only successor work. | T07 R-003 |

All reused ADRs are accepted and non-superseded.
No existing ADR fully owns B-001, B-002, or B-003.
No in-place ADR amendment would honestly represent those new durable choices.
No existing selected alternative is reversed, so no supersession is required.

### Canonical authoring handoff

After the three ADRs are authored in dependency order, bounded authoring must project:

1. PRODUCT-REQ-SPEC-007 amendment for mandatory post-authoring and post-Evidence validation plus explicit human violation judgment.
2. `spec:product.responsibility_boundary_validator` creation and parent-topic registration.
3. A narrow `task_authoring` ownership and usage relation only after W018 findings F-BLK-01 and F-MAJ-01 are independently closed.
4. W019 integrated independent review after the final shared writer.
5. W019 closure synchronization only after the accepted review route.

Exact checklist content remains PRODUCT-WORK-SPEC-020.
Concrete standalone implementation remains PRODUCT-WORK-SPEC-021.
Current DRMCP artifacts remain unchanged.

## Done condition

- Every terminal W019 decision has one resolved ADR-routing outcome.
- Every `required` decision belongs to one coherent ADR boundary.
- Every `covered` decision names an accepted non-superseded ADR.
- Every `not_required` decision records one reason and canonical target.
- Every ADR boundary has an exact disposition and authoring handoff.
- No unresolved route remains except a validly named blocker.
- No ADR content is authored.

## Verification

- Routing covers T01 D-001 through D-012, T03 D-001, T07 R-001 through R-006, and T09 D-001.
- Every terminal decision has one resolved primary routing outcome.
- PRODUCT-ADR-SPEC-015 through PRODUCT-ADR-SPEC-017 are unique unused IDs after PRODUCT-ADR-SPEC-014.
- B-001, B-002, and B-003 can be changed or superseded independently without creating an invalid combined boundary.
- B-001 does not absorb exact checklist wording or runtime details.
- B-002 does not absorb future DRMCP integration design.
- B-003 does not absorb concrete interaction implementation.
- Existing ADR reuse is limited to accepted non-superseded authority.
- No amendment or supersession is selected.
- Every canonical target is exact.
- No ADR, Requirement, Specification, Work Item, graph, review, correction, synchronization, stage, or commit change occurred.

## Evidence

- T07 completed the post-Investigation reconciliation ledger.
- T09 added the independent checklist-authoring and implementation Work Item boundary.
- T11 materialized this routing owner.
- PRODUCT-INV-SPEC-007 found that ADR-004 and ADR-005 define the evaluated Task semantics but do not own validator-specific behavior or the temporary standalone product shape.
- The user accepted the three-boundary split: semantic evaluation contract, standalone ownership boundary, and workflow invocation plus human exception policy.
- PRODUCT-ADR-SPEC-015, PRODUCT-ADR-SPEC-016, and PRODUCT-ADR-SPEC-017 are selected as new ADRs.
- PRODUCT-ADR-SPEC-005, PRODUCT-ADR-SPEC-006, PRODUCT-ADR-SPEC-009, PRODUCT-ADR-SPEC-011, PRODUCT-ADR-SPEC-012, and PRODUCT-ADR-SPEC-014 are selected for reuse.
- T13 is released to materialize exact authoring, shared-writer, integrated-review, and closure Tasks.
- No ADR body, Requirement, Specification, Work Item graph, review, correction, synchronization, implementation, stage, or commit work was performed.

# PRODUCT-INV-SPEC-008: Task responsibility checklist coverage and placement

- **status**: concluded
- **date**: 2026-07-01
- **trigger**: PRODUCT-WORK-SPEC-020 mandatory checklist coverage and placement Investigation
- **scope**: Canonical Task responsibility-boundary coverage required by the common and task-type-specific checklists, plus authority, placement, partitioning, workflow, and writer effects.
- **non_scope**: Final criterion IDs or wording, checklist artifact authoring, evaluator implementation, external response schema design, review, correction, synchronization, stage, and commit.
- **source_refs**:
  - PRODUCT-WORK-SPEC-020
  - PRODUCT-ADR-SPEC-015
  - spec:product.design_records.authoring_standards.task_authoring
  - spec:product.design_records.repository_layout
  - spec:product.brewprint.layout
- **follow_up_candidates**:
  - checklist artifact authoring under the accepted W020 authoring route
  - integrated independent checklist review

## Investigation scope

This Investigation answers one question:

> Which existing canonical Task-authoring rules must be covered by the common and task-type-specific responsibility-boundary checklists, and do the accepted checklist placement and partitioning introduce authority, coverage, writer, or workflow conflicts?

The Investigation covers:

- universal Task responsibility rules;
- all ten canonical `task_type` values;
- Task-local semantic evidence limits;
- responsibility rules excluded from checklist projection;
- the accepted `skills/task-responsibility-boundary-validator/` placement;
- common, type-specific, evaluator-instruction, and skill writer boundaries;
- graph-change and shared-writer candidates.

## Out of scope

- Final `Cxx` or `Txx` criterion assignment.
- Final affirmative checklist sentences.
- Checklist, `SKILL.md`, or evaluator-instruction authoring.
- JSON field names for an external validator response.
- Model, provider, runtime, retry, timeout, or decode policy.
- Validator implementation or DRMCP integration.
- New Task-authoring rules or changed canonical semantics.
- Review, correction, lifecycle synchronization, stage, or commit.

## Background

PRODUCT-WORK-SPEC-020 owns one checklist artifact set for semantic Task responsibility-boundary validation.

The accepted artifact contract fixes Markdown prompt assets under `skills/task-responsibility-boundary-validator/`.
The contract separates common criteria, one file per canonical `task_type`, evaluator instructions, and a minimal `SKILL.md`.
The checklist may only condense existing canonical authoring rules.

PRODUCT-ADR-SPEC-015 limits each invocation to one Task record.
The ADR prohibits external inference and composes the applied checklist from common and declared-type criteria.

`spec:product.design_records.authoring_standards.task_authoring` owns the canonical Task responsibility model.
The checklist must project that model without becoming another authoring authority.

## What was investigated

### Canonical authorities read

| authority | sections used | investigation use |
|---|---|---|
| `spec:product.design_records.authoring_standards.task_authoring` | `### File shape`; `### Metadata schema`; `### Status lifecycle`; `#### Task type contract`; `#### Single responsibility`; `#### Common section alignment`; `#### Implementation contract`; `#### Adjacent responsibility boundaries`; `#### Decision workflow Evidence`; `#### Task continuation and reconvergence routes`; `#### General Task rules`; `### Canonical reference policy`; `## Authoring interface requirements` | Canonical responsibility rules, type semantics, adjacent boundaries, and exclusion classification. |
| PRODUCT-ADR-SPEC-015 | `## Decision`; `## Consequences` | Task-local evidence, checklist composition, and semantic versus structural failure boundaries. |
| PRODUCT-WORK-SPEC-020 | `## Boundary`; `## Task flow`; `## Completion Condition` | W020 ownership, downstream route, and completion boundary. |
| PRODUCT-TASK-SPEC-020-02 | `### Decision inventory`, rows D-001 through D-003; `### Fixed inputs` | Accepted format, placement, partitioning, authority, and cognitive-load constraints. |
| `spec:product.design_records.repository_layout` | `## Current contract`; `## Rules`; `## Boundary` | Normative Design Record placement boundary. |
| `spec:product.brewprint.layout` | `## What this is`; `## Other current repository areas`; `## Maintenance rule` | Non-normative repository inventory and placement interpretation. |
| `spec:product.design_records.authoring_standards.investigation_authoring` | `### File shape`; `### Metadata schema`; `### Status lifecycle`; `### Kind-specific authoring rules`; `### Canonical reference policy` | Investigation shape, metadata, lifecycle, and decision boundary. |
| `spec:product.design_records.authoring_standards.writing_standard` | `### Spec-side rules`; `### AI output rules` | Investigation prose and candidate wording discipline. |
| `spec:product.design_records.authoring_standards.agent_authoring_policy` | `### DRMCP retrieval`; `### Authoring transaction preference` | Filesystem authoring mode. |
| `skills/design-convergence-workflow/SKILL.md`; `skills/design-convergence-workflow/impact-investigation.md` | `SKILL.md`, `## Canonical ownership` and `## Shared writers`; `impact-investigation.md`, `## Graph-change candidates` and `## Shared-writer candidates` | Graph-change, shared-writer, and follow-up ownership checks. |

### Repository placement inspected

The repository currently contains `skills/` with two existing skill directories.
The accepted checklist target directory does not yet exist.

No repository-wide traversal was used.

## Findings

### Evaluation boundary

The checklist evaluates whether one Task record defines one coherent responsibility.
The checklist does not prove that external artifacts, commands, graph changes, or reviews actually satisfy their contracts.

This distinction follows PRODUCT-ADR-SPEC-015, `## Decision`:

- one Task record is the only semantic Evidence source;
- missing Task-local content is non-compliance;
- external records and repository state are not inferred.

A checklist item may inspect Task-local Goal, Work, Done condition, Verification, Evidence, the declared `task_type`, and embedded ledgers.
A checklist item must not require repository lookup to establish semantic compliance.

In the tables below, `task_authoring` means `spec:product.design_records.authoring_standards.task_authoring`.
`writing_standard` and `agent_authoring_policy` mean their exact semantic refs listed above.

### Common coverage inventory

| coverage topic | exact authority | observed canonical rule | classification | checklist coverage | projection boundary |
|---|---|---|---|---|---|
| Primary outcome unity | `task_authoring`, `#### Single responsibility`, first paragraph | Every Task owns one primary outcome matching `task_type`. | common | required | Judge the Task-local declared outcome, not external delivery. |
| Completion judgment unity | `task_authoring`, `#### Single responsibility`, first paragraph | Every Task owns one completion judgment matching the outcome and `task_type`. | common | required | Judge one completion boundary across the Task text. |
| Declared-type alignment | `task_authoring`, `#### Task type contract`, all rows; `#### Common section alignment` | Goal, Work, Done condition, and Verification must express the declared type responsibility. | common plus selected type | required | Use the declared valid `task_type`; invalid or missing type is a structural precondition failure under PRODUCT-ADR-SPEC-015. |
| Goal alignment | `task_authoring`, `#### Common section alignment`, `Goal`; `#### General Task rules`, first bullet | Goal declares the Task-owned outcome and does not restate the Work Item goal. | common | required | Judge semantic ownership, not the exact heading format. |
| Work alignment | `task_authoring`, `#### Common section alignment`, `Work`; `#### General Task rules`, second bullet | Work contains only actions needed for the owned outcome and is not an execution log. | common | required | Supporting steps are allowed only within the same outcome. |
| Done-condition alignment | `task_authoring`, `#### Common section alignment`, `Done condition`; `#### General Task rules`, third bullet | Done condition defines one observable completion judgment for the outcome. | common | required | Status values alone do not define completion. |
| Verification alignment | `task_authoring`, `#### Common section alignment`, `Verification`; `#### Implementation contract`, final rule | Verification confirms the Done condition and adds no new acceptance requirement. | common | required | Judge consistency with the existing acceptance boundary. |
| Supporting-action boundary | `task_authoring`, `#### Common section alignment`, final paragraphs | Another-type action may remain only without a separate deliverable or completion judgment. | common | required | A support action becomes a split candidate when it owns either boundary. |
| Outcome, judgment, or type split | `task_authoring`, `#### Single responsibility`, split list | Different outcome, completion judgment, or Task type requires a separate Task. | common | required | Judge differences visible in the Task record. |
| Acceptance or verification split | `task_authoring`, `#### Single responsibility`, split list and multi-file rule | A separate acceptance or verification boundary requires a separate Task. | common | required | Multiple checks may remain only under one acceptance gate. |
| Owner, release, or independence split | `task_authoring`, `#### Single responsibility`, split list; `#### Task continuation and reconvergence routes`, second paragraph | Different owner, release decision, or required independence requires a separate Task. | common | required | Actual person identity is external; judge the ownership and independence contract declared by the Task. |
| Unresolved-decision split | `task_authoring`, `#### Single responsibility`, multi-file rule; `#### Adjacent responsibility boundaries`, authoring and implementation rows | A file group cannot contain an unresolved design decision. Authoring and implementation stop when contract judgment is required. | common | required | Judge whether the Task claims to select unresolved contract meaning. |
| Multi-file cohesion | `task_authoring`, `#### Single responsibility`, multi-file paragraphs | Multiple files may share one Task only under one outcome, judgment, acceptance boundary, owner, and release decision. | common | required | File count is not a responsibility test. |
| Continuation versus new Task | `task_authoring`, `#### Task continuation and reconvergence routes`, first two paragraphs | An incomplete Task may expand only when type, outcome, and judgment remain unchanged. | common | required | Judge whether the current Task contract combines a newly independent responsibility. |

These topics are coverage inventory labels.
They are not final checklist statements or criterion identifiers.

### Canonical Task-type outcome and overlap inventory

Every row below requires one type-specific checklist projection.

| task type | canonical primary outcome | canonical completion judgment | prohibited overlaps and adjacent boundary | exact authority |
|---|---|---|---|---|
| `investigation` | One Investigation for one bounded research question. | The Investigation satisfies its completion requirements and the Task Done condition. | No decision adoption, canonical authoring, implementation, independent review, correction, or synchronization. Investigation owns evidence, uncertainty, and options. | `task_authoring`, `#### Task type contract`, row `investigation`. |
| `decision` | One bounded decision ledger. | Every owned item is terminal and the Task Done condition is satisfied. | No Investigation authoring, canonical ADR or Specification authoring, implementation, independent review, correction, or final synchronization. Decision fixes selected outcomes; authoring writes fixed inputs. | `task_authoring`, `#### Task type contract`, row `decision`; `#### Adjacent responsibility boundaries`, row `decision versus authoring`; `#### Decision workflow Evidence`. |
| `authoring` | One bounded artifact set from decided inputs. | The artifacts satisfy their authoring requirements and the Task Done condition. | No unresolved decision work, implementation, independent review, finding correction, or lifecycle synchronization. Investigation creation and finding repair use other types. | `task_authoring`, `#### Task type contract`, row `authoring`; `#### Adjacent responsibility boundaries`, authoring stop rule. |
| `implementation` | One bounded implementation outcome. | The implementation contract, Task Done condition, and declared verification pass. | No unresolved design decisions, canonical design authoring, independent review, finding correction, coordination, or synchronization. Local choices remain observable-equivalent. | `task_authoring`, `#### Task type contract`, row `implementation`; `#### Implementation contract`; `#### Adjacent responsibility boundaries`, row `implementation detail versus contract decision`. |
| `review` | One bounded independent verdict and finding set. | Result is `PASS` or `NEEDS REVISION`, findings are complete, and the Task Done condition is satisfied. | No authoring, implementation, correction, self-closure of repaired findings, or lifecycle synchronization. Review owns semantic soundness; verification owns predefined objective checks. | `task_authoring`, `#### Task type contract`, row `review`; `#### Adjacent responsibility boundaries`, rows `review versus verification` and `correction versus finding closure`. |
| `correction` | One bounded named finding set repaired. | Named repairs, Task Done condition, and direct verification pass. | No independent finding closure, unrelated improvement, new decision adoption, or lifecycle synchronization. A later independent review closes findings. | `task_authoring`, `#### Task type contract`, row `correction`; `#### Adjacent responsibility boundaries`, row `correction versus finding closure`; `#### Task continuation and reconvergence routes`. |
| `verification` | One bounded objective acceptance gate. | Every predefined check has expected and actual results, with `PASS`, `FAIL`, or valid `BLOCKED`. | No artifact modification, undefined semantic judgment, repair, independent review verdict, or lifecycle synchronization. Several checks require one shared gate. | `task_authoring`, `#### Task type contract`, row `verification`; `#### Adjacent responsibility boundaries`, review boundary and separate-verification rule. |
| `coordination` | One bounded workflow-graph change. | Required Task, dependency, blocker, owner, writer-order, review-order, and release-route changes are persisted. | No Work Item decomposition, child-owned deliverables, implementation, review, correction, or synchronization. Coordination changes graph and route structure. | `task_authoring`, `#### Task type contract`, row `coordination`; `#### Adjacent responsibility boundaries`, coordination rows. |
| `work_item_decomposition` | One bounded parent-to-child Work Item decomposition. | Required child Work Items have distinct responsibilities, parent routing, and the Task Done condition. | No Task-graph coordination, child-owned deliverables, implementation, review, correction, or synchronization. Work Item identity is decided before decomposition. | `task_authoring`, `#### Task type contract`, row `work_item_decomposition`; `#### Adjacent responsibility boundaries`, row `coordination versus work_item_decomposition`. |
| `synchronization` | One bounded propagation of an accepted result. | Lifecycle, Evidence, completion-result, and relation state express the same accepted result. | No new judgment, decomposition, substantive deliverables, implementation, review, or correction. Graph or choice changes return to coordination or decision. | `task_authoring`, `#### Task type contract`, row `synchronization`; `#### Adjacent responsibility boundaries`, row `coordination versus synchronization`; synchronization stop rules. |

### Task-local type-specific coverage and semantic limits

| task type | checklist-visible responsibility evidence | rule that must not be strengthened or externally inferred | format, metadata, or lifecycle-only rule excluded from semantic coverage |
|---|---|---|---|
| `investigation` | Goal identifies one Investigation and one bounded question. Work remains research. Done uses Investigation completion as the sole judgment. | Do not claim the referenced Investigation actually satisfies its authoring contract without reading it. Do not require a checklist-specific research method. | Investigation ID, path, metadata, and ten-section shape belong to Investigation authoring validation. |
| `decision` | Goal and ledger own one decision boundary. Every item has a terminal route. Work does not author canonical ADR or Specification content. | Do not require one ledger table shape. Do not treat downstream ADR, Specification, review, or closure progress as decision completion. | Task status and metadata formatting remain structural. Exact ADR artifact shape belongs to ADR authoring. |
| `authoring` | Goal names one decided artifact set. Work projects fixed inputs. Done and Verification use one artifact acceptance boundary. | Do not infer that source decisions are accepted or that external artifacts satisfy their guides. Do not choose among materially different interpretations. | Artifact file paths, headings, metadata, and prose-style compliance are separate authoring checks. |
| `implementation` | Goal owns one implementation outcome. Acceptance remains within one implementation contract. Done and Verification add no new acceptance. | Do not infer code state or command success beyond Task-local Evidence. Do not convert local implementation choices into new contract requirements. | Exact `Implementation contract` heading order and table shape are structural. Semantic acceptance alignment remains in coverage. |
| `review` | Goal owns one independent verdict and finding set. Work evaluates rather than repairs. Done returns the canonical verdict boundary. | Do not infer actual reviewer identity or independence from repository history. The Task must declare the independence boundary without self-closing repaired findings. | Exact verdict field presentation and Task status are format or lifecycle concerns. |
| `correction` | Goal names one existing finding set. Work repairs only those findings. Verification is direct. Done does not claim independent closure. | Do not infer that findings exist or are closed outside the Task. Do not expand a repair into unrelated improvement or new choice. | Finding record formatting and lifecycle state belong to their owning artifacts. |
| `verification` | Goal owns one predefined objective gate. Work executes checks without modifying artifacts. Evidence records expected and actual results. | Do not reproduce commands or establish real-world results through semantic evaluation. Do not turn undefined semantic judgment into objective verification. | Command syntax, environment, and result serialization are execution-contract concerns. |
| `coordination` | Goal names one graph or route change. Work changes Tasks, dependencies, blockers, owners, order, or release route only. | Do not infer that the graph was persisted. Do not include child Work Item creation or child-owned deliverables. | Relation syntax, exact file edits, and Task lifecycle updates are structural authoring concerns. |
| `work_item_decomposition` | Goal owns one parent-to-child split. Work creates or splits child Work Items and parent routing, not their deliverables. | Do not infer child existence or non-overlap from external records. Do not decide Work Item identity during decomposition. | Work Item IDs, paths, metadata, and relation formatting belong to Work Item authoring validation. |
| `synchronization` | Goal names one accepted result to propagate. Work changes only named mechanically derivable state. Done uses one propagation judgment. | Do not infer accepted source state. Do not choose among propagation options or repair missing work. | Exact status values, metadata syntax, and relation serialization remain lifecycle or structural checks. |

### Rules excluded from checklist projection

| excluded rule category | exact authority | reason for exclusion |
|---|---|---|
| Task ID grammar and file path | `task_authoring`, `### ID grammar`; `### File path layout` | Structural identity and placement, not semantic responsibility. |
| Exact H1, metadata position, and heading presence | `task_authoring`, `### File shape` | Document-shape validation. Semantic section roles remain covered separately. |
| Metadata field grammar and relation normalization | `task_authoring`, `### Metadata schema`; `### Canonical reference policy` | Structural validation. Valid declared `task_type` is a semantic evaluator precondition. |
| Status lifecycle and `TBD` gating | `task_authoring`, `### Status lifecycle` | Lifecycle and completeness validation, not responsibility partitioning. |
| Exact implementation heading order and table shape | `task_authoring`, `#### Implementation contract` | Structural presentation. The single acceptance-boundary meaning remains covered. |
| Create and update interface mechanics | `task_authoring`, `## Authoring interface requirements` | Authoring-tool contract, not Task-local responsibility semantics. |
| Prose style | `writing_standard`, `### Spec-side rules`; `### AI output rules` | Writing quality and candidate labeling, not responsibility ownership. |
| Agent filesystem or DRMCP operating mode | `agent_authoring_policy`, `### DRMCP retrieval`; `### Authoring transaction preference` | Agent execution policy, not Task content semantics. |
| External artifact truth | PRODUCT-ADR-SPEC-015, `## Decision` | The evaluator may use only one Task record. It cannot inspect actual artifact or graph state. |
| Validator runtime and external response schema | PRODUCT-WORK-SPEC-020, `## Boundary`; PRODUCT-TASK-SPEC-020-02, `### Fixed inputs` | W020 checklist authoring does not own implementation or external API schema. |

Exclusion does not weaken the canonical rules.
Other validators and authoring workflows remain responsible for those rules.

### Authority alignment result

The accepted checklist contract is a derived projection of `task_authoring`.
No checklist topic requires a new Task-authoring rule.

The authority order is coherent:

1. `task_authoring` remains the canonical authoring authority.
2. PRODUCT-ADR-SPEC-015 defines Task-local semantic evaluation behavior.
3. The checklist assets condense only responsibility-boundary rules.
4. The implementation consumes the assets without redefining their meaning.

A checklist conflict with `task_authoring` must resolve in favor of `task_authoring`.
This rule is already fixed by the checklist artifact contract.

### Placement and partitioning effects

| concern | observed effect | conflict result | required authoring boundary |
|---|---|---|---|
| Canonical authority separation | `spec:product.design_records.repository_layout` governs Design Records under `<app>/records/`. The accepted `skills/` target is outside that tree. | no conflict | Checklist assets must identify themselves as derived evaluator assets and point to canonical authority. |
| Standards discovery | The target is outside `product/records/spec/design-records/authoring-standards/` and `product/records/guides/`. | no conflict | Normal standards discovery must not load the checklist set. |
| Independent-authority appearance | Markdown under `skills/` can look instructional without an authority marker. | manageable risk | The minimal skill and evaluator instructions must preserve the canonical-authority pointer and conflict precedence. Final wording belongs to authoring. |
| Temporary implementation separation | The target stores evaluator prompt assets. W020 excludes validator source implementation. | no conflict | No runtime source, provider configuration, or executable implementation belongs in this directory under W020. |
| Common versus type files | Repeating universal rules in every type file creates drift and prompt bloat. | manageable risk | Common files own universal topics. Type files own only type outcome, completion, prohibited overlap, and adjacent-boundary deltas. |
| `SKILL.md` boundary | T02 requires a minimal skill wrapper. | no conflict | The skill owns loading, selection, and canonical-authority routing. It must not duplicate criteria. |
| Evaluator-instruction boundary | T02 separates shared evaluation and structured model-output instructions from checklist files. | no conflict with qualification | Evaluator instructions own shared procedure and the model-facing structured output contract. They do not define the external validator API. |
| Checklist-file boundary | T02 separates one common file and one file per type. | no conflict | Checklist files own only responsibility coverage topics rendered as final criteria by the authoring Task. |
| T04 ownership | All prompt assets serve one artifact-set outcome, one completion judgment, one review boundary, and one release route. | no split required | One authoring Task can own the whole set. |
| Writer order | No other current Task is authorized to write these new assets before integrated review. | no shared-writer conflict | T04 is the sole initial writer. T05 remains an independent reviewer. Finding correction is materialized only after named findings. |

The phrase `response schema` in the T04 exclusion could be read broadly.
The T04 Done condition narrows the exclusion to an external response contract.
T02 already fixes a separate model-facing structured JSON instruction.
T04 can follow both rules without selecting external field names.

### Semantic conflict and stale-representation result

| candidate | classification result | evidence | route effect |
|---|---|---|---|
| Checklist content versus canonical Task rules | no semantic conflict | Every required topic maps to `task_authoring`. Excluded topics remain canonical elsewhere. | T04 may proceed. |
| `skills/` placement versus Design Record layout | no semantic conflict | `repository_layout` owns only `<app>/records/` placement. | No layout amendment required. |
| Root `skills/` omission from current inventory tables | no material stale representation | `brewprint.layout` is non-normative and permits other root development areas. The existing `skills/` root predates this target. | No T04 blocker or required layout update. |
| Evaluator JSON instruction versus external response schema exclusion | no material conflict under accepted distinction | T02 fixes model-facing output instructions. T04 excludes external response contracts. | T04 must preserve the distinction. |
| Common and type-specific duplication | risk, not current conflict | No checklist files exist yet. | T04 verification must reject avoidable duplication. |

### Graph-change candidates

No mandatory graph change was found before T04.

The existing route provides:

- one completed artifact-contract decision;
- one concluded coverage Investigation;
- one bounded authoring Task;
- one integrated independent review Task;
- one closure synchronization Task.

A new decision or coordination Task becomes necessary only when T04 discovers a materially different interpretation, external response schema choice, missing owner, or changed writer order.
That condition is not present in the current authority.

### Shared-writer candidates

No additional shared writer is required for initial authoring.

| target | initial writer | later writer conditions | required order |
|---|---|---|---|
| Minimal `SKILL.md` | W020 checklist authoring owner | Finding correction owner, only after named findings | authoring, review, conditional correction, closure review |
| Evaluator instructions | W020 checklist authoring owner | Finding correction owner, only after named findings | authoring, review, conditional correction, closure review |
| Common checklist | W020 checklist authoring owner | Finding correction owner, only after named findings | authoring, review, conditional correction, closure review |
| Ten type-specific checklists | W020 checklist authoring owner | Finding correction owner, only after named findings | authoring, review, conditional correction, closure review |

The integrated reviewer must not author or repair the reviewed assets.

### Uncertainty and missing evidence

| ID | uncertainty or evidence limit | next owner | effect |
|---|---|---|---|
| U-001 | Exact checklist filenames are not fixed beyond the accepted partition. | W020 checklist authoring owner | Authoring detail. No design decision is required when names preserve deterministic selection. |
| U-002 | Task-local evaluation cannot verify actual external artifact, graph, command, or reviewer state. | W020 checklist authoring owner | Final criteria must judge Task-local responsibility claims only. |
| U-003 | Criterion granularity may vary while preserving the same canonical rule. | W020 checklist authoring owner, then integrated reviewer | Authoring must minimize duplication and must not strengthen meaning. |
| U-004 | Model-facing structured JSON instructions must remain distinct from an external validator response schema. | W020 checklist authoring owner | T04 must stop if external field design becomes necessary. |

No uncertainty currently blocks T04.

## Cross-cutting observations

- Responsibility coverage is semantic, not a replacement for structural Task validation.
- The common checklist should contain only universal responsibility rules.
- Type files should contain only declared-type deltas and adjacent boundaries.
- Every type has external completion facts that a Task-local evaluator cannot independently verify.
- Final criteria must assess the Task's ownership contract and recorded Task-local Evidence.
- Exact headings, metadata, paths, lifecycle, and prose style remain outside semantic responsibility coverage.
- The accepted `skills/` placement separates evaluator assets from canonical authoring standards.
- One authoring Task can own all files because the artifact set has one outcome and one acceptance boundary.

## Follow-up judgment candidates

No mandatory follow-up design judgment was found.

| candidate condition | required owner | route |
|---|---|---|
| T04 cannot express a coverage topic without strengthening or weakening canonical meaning. | new decision owner created through graph coordination | Stop authoring and return to decision. |
| T04 requires external validator response field names. | validator implementation or contract decision owner | Keep the choice outside W020 checklist authoring. |
| Another Task becomes a concurrent writer for the prompt assets. | graph coordination owner | Persist writer order before authoring continues. |
| Integrated review records named findings. | finding-specific coordination owner | Create correction and independent closure-review Tasks. |

These are conditional stop routes.
They are not current graph changes.

## Recommendation

T04 appears ready to start.

T04 should use this Investigation as the complete coverage inventory.
The authoring Task should:

- project all common topics listed above;
- project all ten type outcomes, completion judgments, and prohibited overlaps;
- judge only Task-local responsibility evidence;
- keep structural, metadata, lifecycle, prose, runtime, and external response rules outside the checklist;
- keep universal topics out of type-specific files;
- preserve the canonical-authority pointer and precedence rule;
- author evaluator-internal structured output instructions without defining an external validator API.

No additional decision, graph amendment, Work Item split, or shared-writer order is required before T04.

## Follow-up artifact candidates

- The accepted checklist artifact set under `skills/task-responsibility-boundary-validator/`.
- One minimal skill wrapper.
- One shared evaluator-instruction asset.
- One common responsibility checklist.
- One checklist for each of the ten canonical Task types.
- One integrated independent review result after authoring.
- Conditional finding-specific correction artifacts only when review records named findings.

## Open questions

No blocking open question remains.

The following authoring checks remain open for T04 and T05:

1. Does every final criterion preserve the exact canonical strength?
2. Can every final criterion be judged from one Task record?
3. Are common topics absent from type files unless a type-specific relation requires them?
4. Does the evaluator instruction avoid external validator response design?
5. Does the final asset set remain directly consumable without implementation inference?

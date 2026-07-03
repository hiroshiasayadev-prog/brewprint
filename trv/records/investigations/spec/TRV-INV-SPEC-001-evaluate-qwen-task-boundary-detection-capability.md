# TRV-INV-SPEC-001: Evaluate Qwen Task boundary detection capability

- **status**: concluded
- **date**: 2026-07-03
- **trigger**: Direct Ollama delegation reproduced TRV-style Task responsibility checks and raised whether a separate TRV application still adds material value.
- **scope**: Evaluate whether `qwen3.6-27b:q4_k_m-tools` can detect deliberate Task responsibility-boundary violations and distinguish one compliant control Task.
- **non_scope**: Repository-wide design review, production benchmarking, model selection, final TRV scope decisions, and implementation.
- **source_refs**:
  - TRV-REQ-SPEC-001
  - spec:product.responsibility_boundary_validator
  - spec:product.design_records.authoring_standards.task_authoring
- **follow_up_candidates**:
  - Reassess TRV application scope against direct Ollama delegation
  - Create a labeled Task responsibility regression corpus
  - Define the minimum TRV-specific value beyond raw model invocation
- **related_work_items**:
  - TRV-WORK-SPEC-001
- **related_specs**:
  - spec:trv
  - spec:trv.model_runtime

## Investigation scope

This Investigation asks one bounded question:

> Can the current Qwen 27B model act as a useful first-pass semantic lint engine for Task responsibility-boundary validation?

The Investigation compares deliberate Task violations with one compliant control.
The Investigation also records response-shape and rationale-quality weaknesses that affect TRV runtime value.

## Out of scope

- Statistical accuracy claims from a representative Task corpus.
- Comparison with other local or hosted models.
- Repository-wide design or Work Item graph review.
- Validation of large real-world decision ledgers.
- Final adoption, cancellation, or redesign of TRV.
- Production runtime, retry, persistence, transport, or integration implementation.

## Background

Direct Ollama delegation successfully evaluated completed PRODUCT Task records against common and Task-type checklists.
Those evaluations returned compliant results but did not prove violation-detection capability.

A synthetic test was therefore run without writing the test Tasks to the repository.
The test used the same PRODUCT responsibility criteria that TRV is intended to apply.

## What was investigated

### Model and invocation conditions

| item | value |
|---|---|
| Model | `qwen3.6-27b:q4_k_m-tools` |
| Temperature | `0` |
| Reasoning mode | `off` |
| Evidence boundary | Synthetic Task text only |
| Checklist | Common criteria plus the declared Task-type criteria |
| Output mode | Structured JSON requested |
| Trial count | One pass per synthetic Task |
| Persistence | Synthetic Tasks were not written to the repository |

The host tool blocked one batch request.
The synthetic Tasks were therefore evaluated individually.

### Synthetic Task set

| Task | declared type | planted condition |
|---|---|---|
| A | `decision` | Canonical Specification authoring and implementation were combined with decision work. |
| B | `review` | The reviewer corrected findings and closed the parent Work Item. |
| C | `authoring` | The Task made an unresolved protocol decision before authoring. |
| D | `coordination` | ADR authoring was combined with Task graph materialization. |
| E | `decision` | Verification added a new latency acceptance condition absent from the Done condition. |
| F | `authoring` | Compliant control that projected accepted decisions into one artifact set. |

## Findings

### Task-level detection result

| result | count |
|---|---|
| Violating Tasks with at least one relevant non-compliant result | 5 of 5 |
| Compliant control Tasks with no false non-compliant result | 1 of 1 |

Every deliberately violating Task produced at least one relevant failure.
The compliant authoring control returned all 23 common and type-specific criteria as compliant.

### Planted violation classes

| violation class | detection result | evidence from model output |
|---|---|---|
| Decision Task owns canonical Specification authoring | detected | Decision criterion prohibiting canonical Specification authoring returned false. |
| Decision Task owns implementation | detected | Decision criterion prohibiting implementation returned false. |
| Review Task owns finding correction | detected | Review criterion prohibiting finding correction returned false. |
| Review Task owns parent lifecycle closure | missed by the type-specific criterion | The model identified multiple completion and acceptance boundaries, but the lifecycle-synchronization criterion returned true. |
| Authoring Task owns unresolved design judgment | detected | Common unresolved-decision criteria and authoring projection criteria returned false. |
| Coordination Task owns ADR authoring | detected | Common separate-deliverable criteria and the coordination-only-change criterion returned false. |
| Verification adds a new acceptance condition | detected | The Verification criterion returned false for the added p95 latency threshold. |

The model detected six of seven deliberately planted violation classes.
The missed class depended on repository-specific lifecycle ownership rather than surface wording alone.

### Strong detections

The model correctly identified these responsibility patterns:

- one Task owning multiple primary outcomes;
- multiple independent completion judgments;
- a supporting action owning a separate deliverable;
- authoring from undecided rather than decided inputs;
- Task-type-specific prohibited overlaps;
- Verification introducing a requirement absent from the Done condition;
- ADR authoring escaping a coordination-only boundary.

### Rationale and locator weaknesses

Several criterion results were useful even when the reason was imprecise.
Observed weaknesses included:

- finding correction described as implementation;
- a design decision described as Investigation work;
- a benchmark written in Verification described as Work content;
- one-owner feasibility justified from weak evidence;
- lifecycle synchronization missed despite generic boundary failures;
- section attribution that did not always match the cited Task section.

Binary criterion results appeared more reliable than the detailed rationale taxonomy.

### Structured-output weakness

One response contained the expected criterion items but wrapped them in an unexpected object shape.
The runtime rejected the response against the requested JSON Schema.

This failure did not establish semantic failure.
It demonstrated that raw model invocation needs completeness checks, schema validation, and a retry or execution-failure route.

### Limits of the result

The experiment used six small handcrafted Tasks and one trial per Task.
The result does not establish production recall or precision.

The experiment did not test:

- long real-world Task records;
- large decision ledgers;
- subtle contradictions across many decisions;
- repeated-run stability;
- evidence-locator accuracy at scale;
- ambiguous Task prose.

## Cross-cutting observations

- Explicit checklist criteria substantially constrained model behavior.
- Task-local evaluation is a tractable use case for a 27B model.
- Repository-specific ownership terms remain harder than visible responsibility mixing.
- Larger context permits larger inputs but does not establish stronger semantic judgment.
- A compliant result from the model remains weaker evidence than an independent design review.
- Direct Ollama delegation already supplies model invocation, source loading, output limits, and schema enforcement.
- TRV value therefore depends on repeatability and workflow guarantees rather than unique access to model reasoning.

## Follow-up judgment candidates

- Decide whether TRV remains a standalone application or becomes a thin validator adapter.
- Decide whether the current multi-stage TRV design remains proportionate to its likely runtime responsibility.
- Define the minimum model-quality evidence required before production implementation.
- Decide which criterion failures require human confirmation because repository-specific semantics are weakly detected.
- Decide whether rationale text is advisory while binary criterion results remain the primary machine output.

## Recommendation

Qwen 27B appears suitable as a first-pass semantic lint engine.
Qwen 27B does not appear suitable as the final authority for Task or repository-wide design correctness.

A labeled regression corpus appears preferable before continuing full TRV implementation planning.
The corpus should contain compliant Tasks, single-fault Tasks, multi-fault Tasks, and repository-specific lifecycle cases.

TRV is likely to add the most value through:

- automatic checklist selection and composition;
- complete-criterion enforcement;
- deterministic schema and result validation;
- malformed-response retry and execution-failure classification;
- checklist and prompt version control;
- workflow invocation after authoring and final Evidence;
- human exception recording;
- regression testing across model and prompt changes.

A large standalone application is not yet justified by this experiment alone.
A thinner application boundary may be sufficient, but that direction requires a separate decision.

## Follow-up artifact candidates

- A new TRV Investigation with a labeled regression corpus and repeated trials.
- A reconciliation decision for `TRV-REQ-SPEC-001` and `TRV-WORK-SPEC-001`.
- An amendment or replacement Work Item if the accepted TRV scope is reduced.
- A focused Specification for evaluator completeness, schema validation, retry, and regression behavior.

## Open questions

- What recall and false-positive rate does Qwen achieve on representative real Task records?
- How stable are criterion results across repeated runs and prompt variations?
- Can repository-specific lifecycle and independence semantics be improved through examples or checklist wording?
- Does a smaller model preserve the useful detection rate?
- Does a stronger model materially improve rationale and evidence-locator accuracy?
- Which existing TRV architecture and contract work remains necessary when direct delegation is already available?

## Appendix: Synthetic inputs and returned results

### Common invocation envelope

Each successful evaluation used these model settings:

```text
model: qwen3.6-27b:q4_k_m-tools
reasoning_effort: off
temperature: 0
max_output_tokens: 3072
semantic evidence: inline synthetic Task text only
criteria: common checklist plus the declared Task-type checklist
```

The evaluator instruction was:

```text
Evaluate this synthetic Task against every listed criterion.
Use only the Task text.
If evidence is insufficient, return false.
Preserve criterion order.
Return JSON only.
```

The requested output shape was:

```json
{
  "results": [
    {
      "id": "C01",
      "result": true,
      "reason": "Task-local reason"
    }
  ]
}
```

The exact checklist wording came from:

- `skills/task-responsibility-boundary-validator/prompts/common.md`;
- the matching file under `skills/task-responsibility-boundary-validator/prompts/task-types/`.

The appendix records every non-compliant result returned for each violating Task.
Compliant results are summarized by count to avoid duplicating the complete checklist text.

### Task A input: decision mixed with Specification authoring and implementation

```markdown
# Decide cache backend and deliver it
- task_type: decision

## Goal
Choose the cache backend and make the selected backend available to the application.

## Work
- Compare Redis and an in-process cache.
- Record the selected option and rationale.
- Update the canonical caching Specification with the selected contract.
- Build the selected cache adapter.

## Done condition
- Every decision item is decided.
- The Specification is updated.
- The selected adapter passes its tests.

## Verification
- Check the decision ledger, Specification diff, and adapter tests.
```

#### Task A returned result

The model produced the expected criterion items but used this unexpected outer shape:

```json
{
  "items": [
    {"id": "C01", "result": false, "reason": "..."}
  ]
}
```

The requested schema required a top-level array for this first trial.
The runtime returned `OUTPUT_VALIDATION_FAILED` even though the criterion content was available.

Returned non-compliant criteria:

| criterion | returned reason |
|---|---|
| `C01` | The Task combines a decision outcome with an implementation outcome. |
| `C02` | The Done condition contains decision, Specification, and adapter-test judgments. |
| `C03` | The Goal states both selection and delivery outcomes. |
| `C04` | Building the adapter is not needed to complete the decision. |
| `C06` | The Done condition defines three observable judgments. |
| `C08` | The adapter is a separate deliverable. |
| `C09` | Passing adapter tests is a separate completion judgment. |
| `C10` | The decision record and adapter code have separate acceptance boundaries. |
| `C12` | The decision and implementation have different release decisions. |
| `C13` | The decision and implementation do not share one independence boundary. |
| `C14` | The Task combines backend selection with adapter delivery. |
| `T04` | The Work owns canonical Specification authoring. |
| `T05` | The Work owns implementation. |
| `T09` | Task completion depends on downstream adapter tests. |

### Task B input: review mixed with correction and parent closure

```markdown
# Independently review revised API contract
- task_type: review

## Goal
Issue an independent verdict on the revised API contract and leave the reviewed files ready for closure.

## Work
- Review the revised API Specification and its decision trace.
- Record findings.
- Directly fix minor wording and broken references found during review.
- Mark the parent Work Item done when no blocking finding remains.

## Done condition
- Result is PASS or NEEDS REVISION with finding evidence.
- All minor findings discovered by this Task are corrected.
- The parent Work Item is done when the result is PASS.

## Verification
- Re-read the corrected files and confirm the parent status.

## Independence
The reviewer did not author the original revision.
```

#### Task B returned result

Returned non-compliant criteria:

| criterion | returned reason |
|---|---|
| `C04` | Directly fixing wording and references is outside the bounded review verdict. |
| `C06` | The Done condition combines verdict, finding correction, and parent closure. |
| `C08` | Corrected files are a separate deliverable from the review verdict. |
| `C09` | Finding correction has a separate completion judgment. |
| `C10` | Review, correction, and parent status do not share one acceptance boundary. |
| `C12` | The review verdict and corrected files have separate release decisions. |
| `C13` | Review and correction do not share the same independence boundary. |
| `T07` | The model classified direct correction as implementation work. |
| `T08` | The Task owns finding correction. |

Important returned compliant results:

| criterion | returned result | observation |
|---|---|---|
| `T04` | `true` | The explicit independence declaration was recognized. |
| `T06` | `true` | The model did not classify correction as authoring. This appears too permissive. |
| `T09` | `true` | Parent Work Item closure was not recognized as lifecycle synchronization. |

Task B therefore detected correction overlap but missed the repository-specific closure classification.

### Task C input: authoring mixed with an unresolved design decision

```text
Title: Author transport contract
Type: authoring
Goal: Author the transport contract from accepted decisions.
Work:
- Decide whether the interface uses transport option A or option B.
- Write the selected option and message contract into the canonical Specification.
Done:
- One option is selected.
- The Specification contains the selected option and complete contract.
Verification:
- Confirm the option choice is reasonable and the Specification is internally consistent.
```

The option names were neutralized after an earlier tool request was blocked.
The semantic violation remained unchanged.

#### Task C returned result

Returned non-compliant criteria:

| criterion | returned reason |
|---|---|
| `C04` | Selecting an option is a design decision, not projection from accepted decisions. |
| `C07` | Checking whether the option is reasonable adds an acceptance requirement absent from Done. |
| `C14` | The Task combines unresolved design judgment with authoring. |
| `T01` | The artifact set is not derived only from already decided inputs. |
| `T03` | The Work does not merely project decided inputs. |
| `T04` | The Task owns an unresolved design decision. |
| `T05` | The model classified the decision activity as Investigation work. This classification is imprecise. |

The core authoring-boundary failures were detected despite the inaccurate `T05` rationale.

### Task D input: coordination mixed with ADR authoring

```text
Title: Coordinate post-decision graph
Type: coordination
Goal: Materialize the accepted downstream workflow graph.
Work:
- Create one authoring Task and one review Task.
- Set their dependencies and release order.
- Author the ADR rationale explaining why review follows authoring.
Done:
- Both Tasks and dependencies are persisted.
- The ADR is accepted.
Verification:
- Inspect the Task graph and ADR text.
```

#### Task D returned result

Returned non-compliant criteria:

| criterion | returned reason |
|---|---|
| `C04` | ADR authoring is not needed to materialize the workflow graph. |
| `C06` | The Done condition combines graph persistence with ADR acceptance. |
| `C08` | The ADR is a separate deliverable. |
| `C09` | ADR acceptance is a separate completion judgment. |
| `C10` | Graph persistence and ADR acceptance have separate acceptance boundaries. |
| `C12` | Graph and ADR work have different release decisions. |
| `C13` | The ADR is outside the graph-change independence boundary. |
| `C14` | The Task combines coordination and separate design-document work. |
| `T03` | ADR authoring is outside the allowed coordination change set. |

The model correctly kept creation of a review Task distinct from performing review.
The review prohibition criterion therefore remained compliant.

### Task E input: Verification adds a hidden acceptance requirement

```text
Title: Decide pagination policy
Type: decision
Goal: Choose one pagination policy for list endpoints.
Work:
- Compare cursor and offset pagination.
- Record one selected option and rationale.
Done:
- D-001 is decided.
Verification:
- Confirm D-001 is terminal.
- Benchmark list latency and require p95 below 50 ms before marking the Task complete.
```

#### Task E returned result

Returned non-compliant criteria:

| criterion | returned reason |
|---|---|
| `C04` | The model treated the benchmark as work outside the decision outcome. The benchmark was actually written in Verification. |
| `C07` | The p95 threshold adds an acceptance requirement absent from the Done condition. |
| `T05` | The model classified the benchmark as implementation or testing work. |

The important detection was `C07`.
The reason correctly identified the hidden acceptance gate.
The section attribution for `C04` and `T05` was imprecise.

### Task F input: compliant authoring control

```text
Title: Author accepted logging contract
Type: authoring
Goal: Project accepted logging decisions into one bounded canonical artifact set.
Work:
- Update the logging Specification from decided items D-001 through D-004.
- Update the adjacent example file so it matches the same accepted contract.
Done:
- Both files express D-001 through D-004 consistently and satisfy their authoring format requirements.
Verification:
- Compare both files with D-001 through D-004.
- Run the declared format checks.
```

The control originally used a retry-policy label.
A later tool request used an equivalent logging-policy label after a host safety block.
The responsibility structure remained identical.

#### Task F returned result

The model returned all common and authoring criteria as compliant:

```text
C01-C14: true
T01-T09: true
Overall criterion count: 23 of 23 true
```

Representative reasons included:

| criterion | returned reason summary |
|---|---|
| `C01` | The Specification and example form one bounded artifact-set outcome. |
| `C06` | Both files share one consistency and format completion judgment. |
| `C07` | Comparison and format checks only verify the Done condition. |
| `C14` | Inputs are explicitly identified as decided items. |
| `T01` | The artifact set is bounded and derived from D-001 through D-004. |
| `T03` | The Work projects accepted decisions into the artifacts. |
| `T04` | No unresolved design decision is present. |

### Tool-level attempts and failures

| attempt | input | returned outcome |
|---|---|---|
| Combined six-Task batch | All synthetic Tasks, all matching criteria, one structured result array | Host safety check blocked the request before model execution. |
| Task A first individual run | Decision Task A and common plus decision criteria | Model evaluation completed, but returned an unexpected wrapper. Runtime reported `OUTPUT_VALIDATION_FAILED`. |
| Task C first wording | Authoring Task with named protocol alternatives | Host safety check blocked the request. Neutral option names succeeded. |
| Task F first wording | Compliant retry-policy control | Host safety check blocked the request. Equivalent logging-policy wording succeeded. |

Host safety blocks are not model-quality evidence.
They explain why the final experiment used one Task per invocation and minor neutral wording changes.

### Appendix interpretation rule

The returned reasons above are concise faithful summaries of the model output.
They preserve the criterion result and substantive rationale.
They are not presented as exact byte-for-byte model transcripts.

The actual structured outputs remain session-local tool results.
No synthetic Task or raw model response was persisted outside this Investigation.

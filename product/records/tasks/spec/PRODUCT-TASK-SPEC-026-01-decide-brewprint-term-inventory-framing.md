# PRODUCT-TASK-SPEC-026-01: Decide Brewprint term-inventory framing

- **id**: PRODUCT-TASK-SPEC-026-01
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-026
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - PRODUCT-WORK-SPEC-026
  - PRODUCT-TASK-SPEC-026-01
  - PRODUCT-TASK-SPEC-026-02
  - PRODUCT-TASK-SPEC-026-03

## Goal

Fix one source disposition and, when proceeding, one actionable downstream Work Item contract for PRODUCT-REQ-SPEC-013.

## Work

- Align the user's Desired Outcome with the Requirement's Required Outcome.
- Decide the source disposition.
- Decide whether helper-tool work belongs inside the same downstream Work Item boundary.
- Fix the downstream Goal, Boundary, Completion Condition, direct source, unknown handling, and initial route.
- Materialize only Tasks uniquely required by accepted framing decisions.
- Route the Requirement clarification to authoring before downstream Work Item creation.

This Task does not scan the corpus, author JSONL observations, implement helper tools, classify terms, or define vocabulary.

### Decision ledger

| ID | topic | status | decision summary | reason | dependency | downstream owner or target |
|---|---|---|---|---|---|---|
| D-001 | Desired Outcome | `decided` | Build a machine-readable observed-use corpus of Brewprint design-governance terms and phrases. | The user accepted the Purpose, Gathering criterion, and JSONL schema. | none | PRODUCT-REQ-SPEC-013 |
| D-002 | Outcome alignment | `decided` | Desired Outcome matches the Required Outcome. | Both stop at candidate discovery, observed-use evidence, schema conformance, and corpus coverage. | D-001 | PRODUCT-REQ-SPEC-013 |
| D-003 | Source disposition | `decided` | `proceed`. | The user directed creating the Requirement and then starting Work Item planning. | D-002 | downstream inventory Work Item |
| D-004 | Unknown handling | `decided` | Carry corpus scope, batch partitioning, extraction route, and output placement into the downstream decision loop. No pre-framing Investigation is required. | These choices do not change the Requirement's outcome or completion meaning. | D-003 | downstream decision Task |
| D-005 | Helper-tool ownership | `decided` | Exclude helper-tool implementation, schema validation tooling, aggregation, and use-case extraction from the downstream inventory Work Item. Separate sessions may author assigned JSONL directly. Later work may decide whether aggregation or tooling is useful after observing corpus size and shape. | The inventory Work Item owns investigation evidence only. Aggregation and tool creation have separate completion judgments and may prove unnecessary. | D-004 | downstream Work Item Boundary |
| D-006 | Downstream Goal | `decided` | Produce one coverage-accountable JSONL corpus of observed Brewprint design-governance term usage across the corpus scope selected by its decision loop. | This is the one investigation outcome required by PRODUCT-REQ-SPEC-013. | D-005 | downstream Work Item |
| D-007 | Downstream Boundary | `decided` | Own corpus-scope selection, batch partitioning, parallel source scanning, direct JSONL observation authoring, and coverage evidence. Exclude aggregation, clustering, use-case extraction, definition, normalization, validation-tool implementation, and vocabulary projection. | Keeps the Work Item limited to investigation and preserves later choices until corpus evidence exists. | D-005 | downstream Work Item |
| D-008 | Downstream Completion Condition | `decided` | Every source file in the selected corpus scope is accounted for by observations or an explicit no-candidate result; each batch records required coverage metadata; every observation uses `bp-wide-term-observation-v1`. | Matches the Requirement while avoiding a required aggregation or tooling outcome. | D-006, D-007 | downstream Work Item |
| D-009 | Direct source | `decided` | PRODUCT-REQ-SPEC-013 is the direct material source. | It independently and materially motivates the downstream inventory work. | D-003 | downstream Work Item `source_refs` |
| D-010 | Initial downstream route | `decided` | Start the downstream Work Item with one `decision` Task that fixes exact corpus scope, output placement, batch partitioning, parallel extraction ownership, and coverage accounting before extraction Tasks are materialized. | These graph and scope choices remain unresolved but belong inside the accepted investigation boundary. | D-005 through D-008 | downstream decision Task |

### Current cursor

- Decision: none
- Loop state: `decision_complete`
- Every owned decision is terminal.

## Done condition

- Every decision item is `decided`, `deferred`, or validly `blocked`.
- One source disposition and reason are explicit.
- The full proceed contract exists when the disposition is `proceed`.
- Every Task uniquely required by the accepted framing route is materialized.
- PRODUCT-TASK-SPEC-026-02 is materialized to amend the Requirement boundary.
- PRODUCT-TASK-SPEC-026-03 is materialized to create the fixed downstream inventory Work Item after T02.
- No downstream corpus or tool work is performed by this Task.

## Verification

- Confirm each accepted user answer appears once in the ledger.
- Confirm no proceed-contract field remains open before completion.
- Confirm no speculative downstream Task is materialized.
- Confirm Task and Work Item responsibility boundaries remain aligned with the framing workflow.

## Evidence

- The user approved the Requirement text covering Purpose, Gathering criterion, JSONL schema, and Required Outcome.
- The user directed proceeding to Work Item planning.
- The user assigned exact corpus scope, Task partitioning, and detailed investigation planning to the Work Item decision loop.
- The user decided that JSONL observations will be authored directly by separate sessions.
- The user excluded validation tooling, aggregation, use-case extraction, and helper-tool implementation from the inventory Work Item.
- The user directed evaluating aggregation or tool needs only after observing the investigation result.
- D-001 through D-010 are decided.
- T02 and T03 are directly materialized from the fixed Requirement clarification and downstream Work Item contract.
- DRMCP is non-operational. Filesystem authoring is the required fallback.

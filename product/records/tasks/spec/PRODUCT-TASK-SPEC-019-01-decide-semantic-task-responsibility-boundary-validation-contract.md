# PRODUCT-TASK-SPEC-019-01: Decide semantic Task responsibility-boundary validation contract

- **id**: PRODUCT-TASK-SPEC-019-01
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-019
- **task_type**: decision
- **estimate**: 1d
- **depends_on**: []
- **outputs**:
  - PRODUCT-TASK-SPEC-019-01

## Goal

Decide the complete semantic Task responsibility-boundary validation contract required before impact Investigation and canonical authoring.

## Work

- Preserve accepted Requirement constraints as terminal decision inputs.
- Preserve prior explicit user decisions without asking them again.
- Decide each unresolved criterion, aggregation, rationale, and failure-boundary item one at a time.
- Persist every explicit answer before advancing the cursor.
- Identify likely canonical targets and provisional ADR routes without authoring them.

This Task does not perform Investigation, graph coordination, ADR authoring, Specification authoring, implementation, review, correction, or synchronization.

## Done condition

- Every owned decision is `decided`, `deferred`, or validly `blocked`.
- At most one decision is `in_discussion`.
- No required decision exists only in chat.
- The complete decision set is ready for mandatory impact and conflict Investigation.

## Verification

- Confirm every decision has a stable local ID and status.
- Confirm every decided item records an outcome, reason, canonical target, and ADR route.
- Confirm the cursor names the first unblocked non-terminal item.
- Confirm accepted Requirement constraints were not reopened.
- Confirm no downstream artifact was authored by this Task.

## Evidence

### Loop state

- Status: `decision_complete`
- Current decision: `none`
- Terminal decisions: `D-001` through `D-012`
- Open decisions: none
- Blocked decisions: none

### Direct authority

- `PRODUCT-REQ-SPEC-007`
- `PRODUCT-REQ-SPEC-005`
- `spec:product.design_records.authoring_standards.task_authoring`
- `skills/design-convergence-workflow/SKILL.md`
- `skills/design-convergence-workflow/interactive-decision-loop.md`

### Decision status definitions

| status | meaning |
|---|---|
| `open` | The decision is known and unblocked but has not been asked. |
| `in_discussion` | The decision is the current cursor. |
| `decided` | One explicit outcome is fixed. |
| `blocked` | Named missing authority or evidence prevents a decision. |
| `deferred` | The user explicitly moved the item outside this Work Item. |
| `superseded` | A later decision entry replaced this item before Task completion. |

### Decision inventory

| ID | Topic | Status | Depends on | Decision summary | Reason | Canonical target | ADR route |
|---|---|---|---|---|---|---|---|
| D-001 | Result granularity | `decided` | none | Return one machine-readable overall binary result and one binary result for each evaluated criterion. | `PRODUCT-REQ-SPEC-007` requires both overall and criterion-level judgments. | PRODUCT validator contract Specification | `candidate` |
| D-002 | Criterion-level rationale | `decided` | D-001 | Return Task-local supporting evidence or a concise reason for every criterion judgment, including both compliant and non-compliant results. | The user requires item-level judgment grounds, not only a final result. | PRODUCT validator contract Specification | `candidate` |
| D-003 | Evidence source boundary | `decided` | D-002 | Judge only from the Task record. Treat missing Task-local evidence as non-compliance and do not infer external context. | `PRODUCT-REQ-SPEC-007` fixes the Task-local evidence boundary. | PRODUCT validator contract Specification | `candidate` |
| D-004 | Checklist selection | `decided` | D-003 | Select the validation checklist from the declared `task_type` and inject that checklist automatically. | The validator must evaluate responsibility against the declared Task type without caller-selected criteria. | PRODUCT validator contract Specification; downstream tool contract | `candidate` |
| D-005 | Checklist composition | `decided` | D-004 | Compose the applied checklist from one common responsibility-boundary criterion set and one declared-`task_type` criterion set. Evaluate every combined criterion independently. | Common criteria avoid duplicated cross-type rules. Type-specific criteria preserve each Task type's prohibited-overlap and completion semantics. | PRODUCT validator contract Specification | `candidate` |
| D-006 | Validator product shape | `decided` | D-001 through D-005 | Provide a read-only semantic Task linter. Predefine checklist content by `task_type`. The MCP selects and attaches the applicable common and type-specific criteria, sends only the Task record and injected checklist to the LLM, and returns criterion-level judgments, reasons, and an overall result. The validator does not generate corrections or Task-splitting proposals. | This keeps checklist ownership deterministic, MCP orchestration explicit, and LLM responsibility limited to semantic evaluation of the supplied Task-local evidence. | PRODUCT validator contract Specification; downstream tool boundary | `candidate` |
| D-007 | Overall-result aggregation | `decided` | D-006 | MCP derives the overall compliance flag by applying logical AND to all completed criterion judgments. The LLM does not produce a separate free-form overall judgment. The flag is `true` only when every evaluated criterion is `true`; any `false` criterion makes it `false`. | Every criterion represents a required responsibility-boundary condition. MCP-side aggregation keeps the final result deterministic and avoids asking the local LLM for a redundant summary verdict. | PRODUCT validator contract Specification; downstream tool contract | `candidate` |
| D-008 | Rationale evidence shape | `decided` | D-006 | Require a concise reason and one or more Task section references for every criterion judgment. Permit exact excerpts as optional supporting evidence. When evidence is missing, identify the expected section and the missing content. Do not require line numbers. | Section references make judgments reviewable without forcing large excerpts. Optional excerpts preserve flexibility, and line numbers would become stale after edits. | PRODUCT validator contract Specification | `candidate` |
| D-009 | Structural-precondition boundary | `decided` | D-006 | Do not start semantic validation when the Task cannot be read or parsed, or when `task_type` is missing or invalid. Report a precondition failure. When the Task and `task_type` are valid but required sections or content are missing, execute the checklist and mark the affected criteria `false`. | Checklist selection requires a valid `task_type`. Missing Task-local content remains part of the semantic compliance claim and must not be hidden as an execution failure. | PRODUCT validator contract Specification; downstream validation boundary | `candidate` |
| D-010 | Execution-failure boundary | `decided` | D-007 | Keep model, runtime, timeout, response-decode, and incomplete-response failures separate from semantic compliance. Do not emit overall or criterion judgments when execution fails. Do not synthesize missing LLM judgments. | A `false` compliance result means the Task was evaluated and found non-compliant. Execution failure does not establish Task quality and must not be converted into a semantic verdict. | PRODUCT validator contract Specification; downstream tool contract | `candidate` |
| D-011 | Checklist identity trace | `decided` | D-006 | Do not require checklist identity, checklist revision, stable criterion IDs, or equivalent trace metadata in the validation result. Require only the criterion judgments and reasons needed by the active MCP response contract. The MCP owns checklist selection and result association. | The validator targets a local LLM and should avoid trace requirements that add output burden without improving the core responsibility-boundary judgment. | PRODUCT validator contract Specification; downstream tool contract | `not_required` |
| D-012 | Authoring and release usage | `decided` | D-007, D-010 | Use one validator and one result contract during Task authoring and before Task release. The validator reports semantic compliance only. Calling workflows decide whether to use the result as authoring feedback or as a release gate. | Separating semantic judgment from workflow enforcement keeps the validator read-only and avoids duplicating result contracts by invocation context. | PRODUCT validator contract Specification; downstream workflow contract | `candidate` |

### Decision closure

- D-001 through D-012 are `decided`.
- No decision remains `open`, `in_discussion`, or `blocked`.
- Mandatory impact and conflict Investigation is the next workflow responsibility.
- No ADR, Specification, implementation, review, correction, or synchronization work was performed by this Task.

### Stop conditions

Stop the loop when:

- a user answer permits materially different interpretations;
- accepted Requirement or Specification authority conflicts with the proposed result;
- a decision requires Investigation evidence not yet available;
- a proposed outcome expands into complete executor-readiness or graph validation;
- a proposed outcome requires concrete model, provider, runtime, MCP schema, checklist wording, or storage decisions excluded by the Work Item.

### Expected downstream route

```text
terminal T01 decision ledger
  -> mandatory impact and conflict Investigation
  -> conditional reconciliation and graph coordination
  -> ADR routing and boundary partitioning
  -> conditional ADR authoring
  -> canonical PRODUCT Specification authoring
  -> integrated independent review
  -> closure synchronization or finding-driven repair route
```

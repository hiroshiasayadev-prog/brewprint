# PRODUCT-TASK-SPEC-020-02: Decide checklist artifact contract

- **id**: PRODUCT-TASK-SPEC-020-02
- **status**: done
- **date**: 2026-07-01
- **work_item**: PRODUCT-WORK-SPEC-020
- **task_type**: decision
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-020-01
- **outputs**:
  - PRODUCT-TASK-SPEC-020-02

## Goal

Decide the machine-consumable checklist artifact contract required before checklist authoring.

## Work

- Decide the checklist source format.
- Decide repository placement and file partitioning.
- Decide the criterion entry schema and internal identity policy.
- Preserve common plus declared-`task_type` composition.
- Preserve one binary judgment and one concise reason per criterion.
- Keep exact criterion wording in the later authoring Task.
- Persist one user judgment at a time.

This Task must not:

- author checklist artifacts;
- perform the mandatory Investigation;
- choose a model, runtime, provider, retry, or timeout policy;
- define external validator response field names;
- modify current DRMCP artifacts;
- implement the validator;
- perform review, correction, synchronization, stage, or commit work.

### Decision inventory

| ID | Topic | Status | Depends on | Decision summary | Reason | Canonical target | ADR route |
|---|---|---|---|---|---|---|---|
| D-001 | Checklist source format | `decided` | none | Use Markdown as the canonical checklist source format and compose the selected Markdown fragments directly into the evaluator prompt. | Markdown keeps the canonical artifact identical to the instruction text seen by the weaker local model, avoids YAML syntax noise and a separate rendering contract, and still permits deterministic common-plus-`task_type` selection through file partitioning. | Checklist artifact set | `not_required` |
| D-002 | Repository placement and file partitioning | `decided` | D-001 | Place the derived evaluator prompt assets under `skills/task-responsibility-boundary-validator/`, with a minimal `SKILL.md`, separate evaluator instructions and common checklist Markdown, and one Markdown checklist per canonical `task_type`. | The checklist set is an execution-oriented condensation of existing authoring standards, not an additional authoring authority. Keeping it under `skills/` avoids exposing it through standards discovery, avoids inflating normal authoring context, and keeps the content independent from the temporary validator implementation. | `skills/task-responsibility-boundary-validator/` | `not_required` |
| D-003 | Criterion entry schema and internal identity | `decided` | D-001, D-002 | Represent each criterion as one short affirmative sentence prefixed by a compact internal key: `Cxx` for common criteria and `Txx` for declared-`task_type` criteria. Keep shared evaluation instructions and the structured JSON response contract outside checklist files and state them once in evaluator instructions. Treat criterion keys as invocation-local mapping aids rather than stable external identifiers. | This minimizes prompt length and cognitive load, avoids repeated per-criterion scaffolding, prevents Markdown checklist structure from competing with JSON output, and still gives the implementation an unambiguous key for each binary judgment and reason. | Checklist artifact set | `not_required` |

### Current cursor

- Current item: none.
- Loop state: `complete`.

### Fixed inputs

- One invocation evaluates one Task record.
- Applied criteria equal the common set plus the declared-`task_type` set.
- Checklist and evaluator-prompt design must minimize cognitive load for a relatively weak local model.
- Repeated instruction text and nonessential context must be eliminated wherever one shared instruction can govern all criteria.
- Checklist representation must not compete with or obscure the model's required structured JSON response.
- Every applied criterion receives one binary judgment and one concise Task-local reason.
- Missing required Task-local content produces criterion non-compliance.
- External validation results do not require checklist revision or stable criterion IDs.
- Checklist content may only condense requirements already established by canonical authoring standards.
- Checklist content must not introduce, strengthen, weaken, or reinterpret authoring requirements beyond those standards.
- Canonical authoring standards remain authoritative whenever checklist wording is incomplete, stale, or conflicting.
- The checklist set is not part of normal authoring-standard discovery and is loaded only for validator execution or checklist maintenance.
- Exact checklist wording belongs to T04, not this decision Task.

## Done condition

- D-001 through D-003 are `decided`, `deferred`, or validly `blocked`.
- The selected contract is sufficient for T03 Investigation and T04 authoring.
- No checklist wording or artifact is authored.
- No implementation choice is introduced.

## Verification

- Confirm at most one item is `in_discussion`.
- Confirm every decided item has a concise reason and target.
- Confirm the contract preserves common plus type-specific composition.
- Confirm the contract remains implementation-independent.
- Confirm no artifact authoring occurred.

## Evidence

- PRODUCT-ADR-SPEC-015 fixes checklist composition and semantic judgment behavior.
- PRODUCT-WORK-SPEC-020 owns exact format, placement, schema, wording, and artifact authoring.
- PRODUCT-TASK-SPEC-019-21 releases this decision before W019 closure.
- User decisions on 2026-07-01 selected Markdown prompt assets under `skills/`, restricted checklist meaning to canonical authoring standards, and selected compact affirmative criteria with separate shared JSON-output instructions.

# Task vocabulary reference: review / verification / correction

General reference of real phrasing used by `review`, `verification`, and `correction` Tasks in this repository's corpus. Not a boundary-violation list — see `skills/task-boundary-vocabulary/` for that.

Source: `PRODUCT-TASK-SPEC-025-06` through `-10` Finding logs, corpus-range extraction, 2026-07-03.

## review

- `Confirm` — judge scope, safety, and reasonableness at an independent review gate (`PRODUCT-TASK-SPEC-001-02`, `-004-02`)
- `is accounted for` — establish source-to-output completeness without requiring byte-identical text (`PRODUCT-TASK-SPEC-005-03/07/11`)
- `what to check` — express review responsibility as checklist criteria rather than an owning verb (`PRODUCT-TASK-SPEC-005-03`)
- `Compare` — review split artifacts against a preserved source for structure, coverage, and boundaries (`PRODUCT-TASK-SPEC-007-04`)
- `review complete` / `Spec passes review` — produce the independent review verdict for a target set (`PRODUCT-TASK-SPEC-009-04`, `-010-03`)
- `Apply required findings` before the compound correction+synchronization split emerged as its own pattern — treated as review-gate language in this corpus's earlier Work Items (`PRODUCT-TASK-SPEC-010-03`)
- `Spec coverage reviewed` — evaluate spec-authoring coverage and record the verdict (`PRODUCT-TASK-SPEC-011-08`)
- `` `PASS` or `NEEDS REVISION` `` / `Record `PASS` or `NEEDS REVISION` with classified findings` — issue the bounded independent-review verdict (`PRODUCT-TASK-SPEC-012-12`, `-017-08`)
- `required corrections` — identify the exact correction set required before closure, as a review output (`PRODUCT-TASK-SPEC-012-12`)
- `Decide each finding as CLOSED or OPEN` / `Classify each finding as closed or still open with exact evidence` — independent finding-closure review (`PRODUCT-TASK-SPEC-018-19`, `-017-10`)
- `Establish reviewer independence from [prior authoring]` — confirm the reviewer did not author the reviewed outputs (`PRODUCT-TASK-SPEC-019-18`)
- `Trace [decisions] through [the design chain]` — evaluate end-to-end decision traceability in an integrated review (`PRODUCT-TASK-SPEC-019-18`)
- `Produce one Investigation that verifies [coverage/alignment/risk]` — a review-flavored Investigation framing, still owned by `investigation` (`PRODUCT-TASK-SPEC-020-03`)

## verification

- `Confirm bpdsl/... is byte-identical` — objectively verify a staging copy preserves the source corpus (`PRODUCT-TASK-SPEC-005-05/09`)
- `Check required/prohibited sections` / `report mismatch` — implemented validator behavior, not a manual verification act (`PRODUCT-TASK-SPEC-006-01`)
- `Pre-existing diagnostics remain separately attributed` — classify non-migration diagnostics without absorbing them into correction scope (`PRODUCT-TASK-SPEC-012-11`)
- `retained app-local owner or deletion rationale` — prove each removed statement has an owner or justified deletion (`PRODUCT-TASK-SPEC-012-08`)
- `Assess every [Work Item] Completion Condition separately` — evaluate closure readiness condition by condition (`PRODUCT-TASK-SPEC-015-03`)
- `Evaluate every [Work Item] Completion Condition` — mechanically check every closure prerequisite before lifecycle propagation (`PRODUCT-TASK-SPEC-019-19`)

## correction

- `Apply all must-fix findings` — repair every blocking finding from a prior review (repeated identical phrase across 6 independent Tasks: `PRODUCT-TASK-SPEC-004-03`, `-005-04/08/12/16/21`; see also `skills/task-boundary-vocabulary/correction.md`)
- `fix errors` — repair every strict-validator diagnostic until validation passes (`PRODUCT-TASK-SPEC-009-03`)
- `corrections applied` — repair all must-fix findings from a preceding review (`PRODUCT-TASK-SPEC-009-05`)
- `Apply required corrections from blocking or major findings` — repair the exact changes required by review findings (`PRODUCT-TASK-SPEC-014-02`)
- `Normalize the [Impact Scope] wording from pointer candidates to retained implementation owners` — align stale wording with an accepted disposition (`PRODUCT-TASK-SPEC-015-02`)
- `Apply only scoped corrections required by review findings` — repair only the named review scope (`PRODUCT-TASK-SPEC-015-03`)
- `Remove stale Task-contract wording within [a Work Item's] scope` — delete normative text contradicting an accepted contract (`PRODUCT-TASK-SPEC-016-07`)
- `Replace [a stale model] and [X] in canonical Specifications` — remove a stale model and author its replacement (`PRODUCT-TASK-SPEC-017-07`)
- `Align [Task flow / Task Candidates] with the single integrated-review route` — repair workflow projections to match an accepted route (`PRODUCT-TASK-SPEC-017-09`)
- `Mark the separate [X] route as not required and superseded by [Y]` — correct a route description without a lifecycle transition (`PRODUCT-TASK-SPEC-017-09`)
- `State that coordination owns Task graph change` — add missing ownership wording to an active instruction file (`PRODUCT-TASK-SPEC-018-18`)

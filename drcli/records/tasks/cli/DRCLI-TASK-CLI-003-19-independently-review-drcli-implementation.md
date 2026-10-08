# DRCLI-TASK-CLI-003-19: Independently review integrated DRCLI implementation

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: review
- **estimate**: 2d
- **depends_on**:
  - DRCLI-TASK-CLI-003-18
- **outputs**:
  - DRCLI-TASK-CLI-003-19

## Goal

Produce one independent integrated code verdict and named findings.

## Work

### Read

- `DRCLI-TASK-CLI-003-19`.
- `drcli/records/work-items/cli/DRCLI-WORK-CLI-003-implement-portable-design-records-cli.md`
- `drcli/records/tasks/cli/DRCLI-TASK-CLI-003-18-verify-complete-portable-cli-integration.md`

### Execute

- An independent reviewer who did not write or correct the code examines exact source/test files for T04-T17 against all 82 spec refs and PRODUCT authority. Check API boundaries, Git ignore, source scopes, error placement/finding identity, write-staleness and coupled atomicity, retained-state portability, CLI help/UTF-8, and Linux/Windows runtime evidence. Return PASS or NEEDS REVISION; every finding has ID, severity, affected target and nominated correction owner. A required finding routes to new coordination, later correction and distinct independent closure-review Tasks only after findings exist. Never change implementation.
- Success boundary: Independent evidence-backed verdict states all required findings and whether T20 closure is eligible.

### Stop

- Record BLOCKED on absent exact predecessor, missing OS verification, missing authority or unapproved changes. Do not guess or rewrite another owner's work.

### Prohibited operations

- No source Specification, PRODUCT, DRMCP, historical completed Task, unrelated code or other Task ownership changes; no stage/commit.

### Output

- Task-local actual results, evidence, exact changed outputs when applicable, and downstream handoff.

## Done condition

- Independent evidence-backed verdict states all required findings and whether T20 closure is eligible.

## Verification

- Compare code, test evidence and 82-row parent spec trace, without trusting self-attestation.

## Evidence

TBD

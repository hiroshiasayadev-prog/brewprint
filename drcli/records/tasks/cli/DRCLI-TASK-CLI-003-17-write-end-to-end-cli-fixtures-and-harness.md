# DRCLI-TASK-CLI-003-17: Implement end-to-end CLI acceptance harness

- **status**: not_started
- **date**: 2026-10-08
- **work_item**: DRCLI-WORK-CLI-003
- **task_type**: implementation
- **estimate**: 3d
- **depends_on**:
  - DRCLI-TASK-CLI-003-16
- **outputs**:
  - drcli/tests/integration/integration_test.go
  - drcli/tests/integration/harness_test.go
  - drcli/tests/fixtures/current/
  - drcli/tests/fixtures/invalid/
  - drcli/tests/fixtures/ignore/

## Goal

Deliver one reproducible subprocess end-to-end contract test harness.

## Work

### Read

- `DRCLI-TASK-CLI-003-17`.
- `drcli/records/spec/design-records-cli/index.md`
- `drcli/records/spec/design-records-cli/operations/index.md`
- `drcli/records/spec/design-records-cli/cli/command-surface.md`

### Execute

- Implement TestEndToEnd* Go subprocess harness invoking compiled DRCLI with isolated temp repositories and retained state. Create six artifact kind fixture sets for valid, malformed, missing-index, conflicting identity, ignored roots, and distinct app roots. Assert all 19 command paths, help, cwd/--repo, --no-ignore, JSON and exit statuses, retrieval/search limits, source finding codes, direct create/update/no-op and process-independent cache/proposal scenarios. Never change production source.
- Success boundary: All 19 command paths and negative cases use exact output assertions, source fixtures are immutable, cache/proposal tests invoke separate processes and the command exits 0.
- Focused command: `go -C drcli test ./tests/integration -run '^TestEndToEnd' -count=1`; require exit 0.

### Stop

- Record BLOCKED on absent exact predecessor, missing OS verification, missing authority or unapproved changes. Do not guess or rewrite another owner's work.

### Prohibited operations

- No source Specification, PRODUCT, DRMCP, historical completed Task, unrelated code or other Task ownership changes; no stage/commit.

### Output

- Task-local actual results, evidence, exact changed outputs when applicable, and downstream handoff.

## Implementation contract

| target | required change | acceptance criterion | verification |
|---|---|---|---|
| `drcli/tests/integration/integration_test.go` | Create the fixed end-to-end tests or fixture inputs. | All 19 command paths and negative cases use exact output assertions, source fixtures are immutable, cache/proposal tests invoke separate processes and the command exits 0. | `go -C drcli test ./tests/integration -run '^TestEndToEnd' -count=1` |
| `drcli/tests/integration/harness_test.go` | Create the fixed end-to-end tests or fixture inputs. | All 19 command paths and negative cases use exact output assertions, source fixtures are immutable, cache/proposal tests invoke separate processes and the command exits 0. | `go -C drcli test ./tests/integration -run '^TestEndToEnd' -count=1` |
| `drcli/tests/fixtures/current/` | Create the fixed end-to-end tests or fixture inputs. | All 19 command paths and negative cases use exact output assertions, source fixtures are immutable, cache/proposal tests invoke separate processes and the command exits 0. | `go -C drcli test ./tests/integration -run '^TestEndToEnd' -count=1` |
| `drcli/tests/fixtures/invalid/` | Create the fixed end-to-end tests or fixture inputs. | All 19 command paths and negative cases use exact output assertions, source fixtures are immutable, cache/proposal tests invoke separate processes and the command exits 0. | `go -C drcli test ./tests/integration -run '^TestEndToEnd' -count=1` |
| `drcli/tests/fixtures/ignore/` | Create the fixed end-to-end tests or fixture inputs. | All 19 command paths and negative cases use exact output assertions, source fixtures are immutable, cache/proposal tests invoke separate processes and the command exits 0. | `go -C drcli test ./tests/integration -run '^TestEndToEnd' -count=1` |

## Done condition

- All 19 command paths and negative cases use exact output assertions, source fixtures are immutable, cache/proposal tests invoke separate processes and the command exits 0.

## Verification

- Run `go -C drcli test ./tests/integration -run '^TestEndToEnd' -count=1`; inspect exact fixture changes and subprocess output.

## Evidence

TBD

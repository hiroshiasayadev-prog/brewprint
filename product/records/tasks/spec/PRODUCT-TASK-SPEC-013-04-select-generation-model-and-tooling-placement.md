# PRODUCT-TASK-SPEC-013-04: Select generation model and tooling placement

- **id**: PRODUCT-TASK-SPEC-013-04
- **status**: done
- **date**: 2026-06-26
- **work_item**: PRODUCT-WORK-SPEC-013
- **source_requirement**: PRODUCT-REQ-SPEC-003
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-013-02
  - PRODUCT-TASK-SPEC-013-03
- **outputs**:
  - PRODUCT-ADR-SPEC-002
  - PRODUCT-WORK-SPEC-013 generation-model evidence update

## Goal

Select the deterministic package-generation model and implementation placement.

Preserve PRODUCT semantic authority while keeping the public generation contract minimal.

## Work

- Compare explicit rebuild generation with watcher or incremental synchronization.
- Select the public generator entrypoint and its PRODUCT-owned placement.
- Fix the authoritative source and repository-local destination.
- Define temporary generation, destination replacement, and rollback boundaries.
- Define distribution-build integration.
- Define the automated-test path seam without widening the public CLI.
- Define the authority boundary for manual edits to the generated package.
- Record the accepted decision and rationale in `PRODUCT-ADR-SPEC-002`.
- Defer concrete checks, diagnostics, exit behavior, cleanup details, and concurrency handling to T05.
- Defer implementation and `scripts/verify.bat` modification to T06.

## Done condition

- An accepted ADR records the generation model and tooling placement.
- The authoritative source remains `product/records/spec/design-records/`.
- The public generator is a single Python command under `product/src/tools/`.
- The repository-local destination is fixed as `bin/design-records/`.
- Generation uses full rebuild, temporary output, replacement, and rollback boundaries.
- The generated package remains a non-authoritative derived artifact.
- Distribution builds must invoke the generator.
- Public path overrides are excluded while internal test seams remain allowed.
- T05 and T06 retain their existing check and implementation ownership.
- No unrelated distribution-layout migration is introduced.

## Verification

- `PRODUCT-ADR-SPEC-002` has status `accepted`.
- The ADR depends on `PRODUCT-ADR-SPEC-001` and supersedes no decision.
- The ADR does not move `bin/design-records-mcp.exe`.
- The ADR fixes `bin/design-records/` as the repository-local package destination.
- The ADR keeps source and destination overrides out of the public CLI.
- The ADR states that manual generated-tree edits are non-authoritative and non-persistent.
- The ADR leaves concrete operational checks and implementation to T05 and T06.
- No generator implementation or `scripts/verify.bat` change is performed by this Task.

## Evidence

### Accepted generation model

| concern | accepted decision |
|---|---|
| generation trigger | Explicit public generator command. |
| public entrypoint | `product/src/tools/generate_design_records_package.py`. |
| authoritative source | `product/records/spec/design-records/`. |
| repository-local destination | `bin/design-records/`. |
| executable placement | Keep `bin/design-records-mcp.exe`. |
| update model | Full rebuild on every invocation. |
| publication model | Generate a temporary tree, then replace the destination. |
| failure recovery | Preserve or restore the previous destination when replacement fails. |
| repository treatment | Do not commit the generated package tree. |
| build integration | `scripts/verify.bat` must invoke the generator during T06. |
| automated tests | Use temporary roots or destinations through internal functions. |
| public CLI | No source or destination override. |
| manual edits | Non-authoritative, discarded by later generation, and never reverse-synchronized. |

### ADR boundary

`PRODUCT-ADR-SPEC-002` is the authoritative decision record for this Task.

The ADR implements the ownership direction established by `PRODUCT-ADR-SPEC-001` and `DRMCP-ADR-MCP-001`.
It does not supersede either decision.

### Deferred scope

T05 owns concrete generation checks, warning emission, operational failures, exit behavior, cleanup, and concurrent-run handling.

T06 owns generator implementation, automated tests, real package generation, and `scripts/verify.bat` integration.

### Closure

The generation model and tooling placement are decided and recorded.
The Done condition is satisfied.
T05 may consume `PRODUCT-ADR-SPEC-002` as its accepted input.

# PRODUCT-TASK-SPEC-013-06: Implement portable package generator and produce initial package

- **id**: PRODUCT-TASK-SPEC-013-06
- **status**: done
- **date**: 2026-06-26
- **work_item**: PRODUCT-WORK-SPEC-013
- **source_requirement**: PRODUCT-REQ-SPEC-003
- **estimate**: 2d
- **depends_on**:
  - PRODUCT-TASK-SPEC-013-03
  - PRODUCT-TASK-SPEC-013-04
  - PRODUCT-TASK-SPEC-013-05
- **outputs**:
  - `product/src/tools/generate_design_records_package.py`
  - generator automated tests
  - `scripts/verify.bat` integration
  - initial generated `bin/design-records/` package
  - operational generation evidence

## Goal

Implement the accepted portable Design Records package generator.

Produce and operationally verify the initial repository-local package without creating a second semantic authority.

## Work

- Implement the fixed public generator at `product/src/tools/generate_design_records_package.py`.
- Resolve the repository root from the generator file location.
- Copy the complete `product/records/spec/design-records/` file tree into a unique temporary tree.
- Apply only the accepted canonical prefix rewrite from `spec:product.design_records` to `spec:design_records`.
- Verify source and generated relative file-path equality.
- Verify that content changes are limited to the accepted canonical rewrite.
- Emit non-blocking warnings for duplicate, unresolved, external, unrewritten, and source-authoring boundary findings.
- Replace `bin/design-records/` only after pre-publication checks pass.
- Preserve or restore the previous destination when replacement or publication confirmation fails.
- Preserve recovery-capable artifacts when rollback fails.
- Keep cleanup-only failures non-blocking after valid publication.
- Add standard-library automated tests for the accepted operational contract.
- Invoke the tests and public generator from `scripts/verify.bat` without path overrides.
- Generate the initial `bin/design-records/` package and record the operational result.
- Confirm the generated tree remains ignored and uncommitted.

## Done condition

- The fixed public generator is implemented without public source or destination overrides.
- Whole-tree copy, canonical rewrite, warnings, replacement, cleanup, and rollback behavior are automated-test covered.
- Operational failures return exit code `1`.
- Semantic warnings retain exit code `0`.
- `scripts/verify.bat` runs the automated tests and generator without changing executable placement.
- The real `bin/design-records/` package is generated successfully.
- The generated package file set matches the source file set.
- The generated package is ignored and not committed.
- Execution commands and results are recorded in Evidence.
- Independent review reports no blocking or major findings before closure.

## Verification

- Run the generator automated tests through the Python standard-library test runner.
- Verify nested Markdown and non-Markdown file copying.
- Verify missing and unexpected file detection.
- Verify stale destination removal through full replacement.
- Verify exact root and suffix canonical ref rewriting.
- Verify ordinary prose, public record IDs, physical paths, and external canonical refs remain unchanged.
- Verify accepted-rewrite-only content integrity.
- Verify all required warning classes and multiple warnings per file.
- Verify warnings do not change the successful exit code.
- Verify source, temporary-tree, copy, rewrite, check, replacement, rollback, and publication-confirmation failures.
- Verify success cleanup and pre-replacement destination preservation.
- Verify publication-confirmation failure restores the previous destination when available.
- Verify rollback failure preserves recovery artifacts and returns exit code `1`.
- Verify cleanup-only failure after valid publication remains successful and emits a warning.
- Inspect imports and root resolution for DRMCP, app-registry, process-state, and current-working-directory dependencies.
- Run the real public generator and compare source and destination relative file-path sets.
- Confirm `bin/design-records/` is ignored or untracked and is not staged.

## Evidence

- Authoring fallback: Design Records MCP retrieved the parent Work Item and Tasks but returned `record_not_found` for `PRODUCT-ADR-SPEC-002` and `PRODUCT-ADR-SPEC-003`. The accepted ADR files were read directly from their known PRODUCT paths, so Task authoring uses filesystem fallback.
- Existing Python convention: `product/src/tools/validate_spec.py` is a standalone standard-library command. No existing Python test directory or third-party Python test framework exists.
- New test directory required: `product/tests/tools/`.
- Publication-confirmation recovery decision: if the newly published destination cannot be enumerated, restore the previous destination from the invocation-specific backup when available. A failed restoration is an operational failure and retains recovery-capable artifacts.
- Automated tests: `python -X utf8 -m unittest discover -s product\tests\tools -p "test_*.py" -v` exited `0`; all 27 tests passed after review corrections.
- Initial generation: `python -X utf8 product\src\tools\generate_design_records_package.py` exited `0` and generated 34 files.
- Warning result: generation emitted 79 non-blocking warnings. Every warning used either `external canonical ref` or `source-authoring boundary finding`; no operational failure occurred.
- Repository-local distribution verification: `scripts\verify.bat` exited `0`. Go tests, Python tests, both executable builds, package generation, and the final `OK` stage completed.
- Generated-artifact boundary: `git check-ignore -v bin/design-records/index.md` exited `0` and matched `.gitignore` rule `bin/`.
- Git status confirmation: `bin/design-records/` did not appear in `git status --short`; the generated package was not staged or committed.
- File-set result: the generator reported 34 destination files, matching the 34-file source tree inspected before implementation.
- Independent review verdict: PASS. No blocking or major findings were reported. One minor coverage gap requested focused tests for destination-backup failure and rollback failure after publication confirmation.
- Review correction: added `test_backup_move_failure_preserves_existing_destination` and `test_publication_confirmation_rollback_failure_retains_recovery_artifacts`.
- Post-correction verification: the 27-test suite exited `0`; `scripts\verify.bat` also exited `0`, including Go tests, Python tests, both executable builds, package generation, 79 non-blocking warnings, and final `OK`.
- Advisory disposition: exact warning detail-set assertions are deferred to T07 portability and release fixture work because the advisory did not affect T06 closure readiness.
- Closure: all T06 Done conditions are satisfied and the independent reviewer marked T06 closure readiness as ready.

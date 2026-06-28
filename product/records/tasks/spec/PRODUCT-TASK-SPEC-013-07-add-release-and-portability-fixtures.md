# PRODUCT-TASK-SPEC-013-07: Add release and portability fixtures

- **id**: PRODUCT-TASK-SPEC-013-07
- **status**: done
- **date**: 2026-06-28
- **work_item**: PRODUCT-WORK-SPEC-013
- **source_requirement**: PRODUCT-REQ-SPEC-003
- **estimate**: 1.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-013-05
  - PRODUCT-TASK-SPEC-013-06
- **outputs**:
  - `product/tests/tools/test_generate_design_records_package_release.py`
  - `.gitignore` test-source tracking exception
  - release and portability fixture evidence

## Goal

Add focused release and portability fixtures for the accepted portable package generator contract.

Verify reproducibility, public-command independence, exact warning reporting, localized semantic defects, and release-script integration without redefining T02-T06 decisions.

## Work

- Add a dedicated release fixture module under `product/tests/tools/`.
- Verify repeated generation produces the same relative path set, file bytes, and ordered warnings.
- Verify full replacement removes stale destination content and leaves no invocation-specific artifact in the package.
- Verify the public command runs from a repository-external current working directory against an isolated repository-shaped fixture.
- Run the subprocess in Python isolated mode and supply irrelevant host-state environment values.
- Verify the public command uses its script location, fixed source, and fixed destination.
- Verify prefix rewriting across metadata, Topics, relation tables, and canonical body refs.
- Verify ordinary prose, public artifact IDs, physical paths, external refs, malformed lookalikes, non-UTF-8 bytes, and unrelated line endings remain unchanged.
- Verify the complete warning fixture shape for duplicate, unresolved, external, unrewritten, and source-authoring boundary findings.
- Verify warning ordering is deterministic and warning scanning does not change content.
- Verify duplicate and unresolved refs remain localized, do not block generation, and do not remove unrelated files.
- Verify public-command operational failure returns exit code `1`.
- Verify `scripts/verify.bat` invokes the discovered test suite and propagates generator failure.
- Reuse temporary fixture trees instead of committing a generated package or large golden tree.
- Keep the nested test source visible to Git without changing the `bin/` ignore boundary.
- Do not change the generator unless a fixture exposes an accepted-contract violation.

## Done condition

- A dedicated release fixture module covers reproducibility, byte-level output stability, stale-file replacement, and artifact-name isolation.
- A process-level fixture verifies repository-external current working directory execution and fixed path resolution.
- Host independence is checked through isolated subprocess execution and standard-library-only import inspection.
- Prefix rewrite preservation rules are verified with explicit expected bytes.
- Warning fixtures assert required classes, relative paths, important details, deterministic order, and unchanged content.
- Duplicate and unresolved defects remain present in generated output while unrelated files are generated.
- Warning-only generation exits `0`; operational public-command failure exits `1`.
- Existing T06 failure-path tests are not duplicated without a release-level reason.
- `scripts/verify.bat` needs no additional invocation when unittest discovery includes the new module.
- The new fixture module is visible to normal Git status and add operations.
- `bin/design-records/` remains ignored, derived, and uncommitted.
- Verification evidence is recorded before independent review.
- Task status remains `in_progress` until independent review passes and findings are closed.

## Verification

- Run `python -X utf8 -m unittest discover -s product\tests\tools -p "test_*.py" -v`.
- Run `python -X utf8 product\src\tools\generate_design_records_package.py`.
- Run `scripts\verify.bat`.
- Inspect scoped Git changes and whitespace.
- Confirm `bin/design-records/index.md` remains ignored.
- Confirm generated package files do not appear in the scoped status.
- Obtain independent review against the accepted T02-T06 contract.

## Evidence

- Authoring fallback: Design Records MCP retrieved the parent Work Item and prior Tasks. The create proposal was not accepted because current indexing reported `PRODUCT-REQ-SPEC-003` as an invalid workflow target. The Task and reciprocal Work Item update use known-path filesystem fallback.
- Investigation result: the existing 27 tests cover rewrite primitives, copy integrity, warning class presence, operational failures, rollback, cleanup, public argument rejection, unit-level root resolution, and a limited dependency string check.
- Gap: the existing suite does not compare repeated package bytes and ordered warnings.
- Gap: the existing suite does not execute the public command from an external current working directory.
- Gap: warning tests do not fix relative paths and important detail shape for all required classes.
- Gap: duplicate and unresolved defects are not verified together through successful package generation with unrelated output retained.
- Fixture placement: use a separate release fixture module to keep T06 unit and failure-injection tests distinct.
- Golden decision: no persistent generated-package golden is required. Explicit expected bytes and repeated output snapshots provide deterministic verification with less maintenance coupling.
- Persistent fixture decision: use temporary focused fixture trees. No committed package tree or `bin/` fixture is required.
- `scripts/verify.bat` decision: no new test command is required because existing unittest discovery includes every `test_*.py` module under `product/tests/tools/`.
- Implementation result: added eight focused release tests in `product/tests/tools/test_generate_design_records_package_release.py`.
- Reproducibility fixture: repeated replacement generation compares the complete relative path-to-bytes snapshot and ordered warning tuple. The second run also removes injected stale and tampered destination content.
- Prefix fixture: explicit expected bytes cover metadata, Topics, relation rows, body refs, prose, public IDs, physical paths, external refs, malformed lookalikes, CRLF, LF-only text, and non-UTF-8 bytes.
- Warning fixture: one file emits duplicate, unresolved, external, unrewritten, and source-authoring boundary classes. The test fixes class, relative path, deterministic order, semantic detail elements, and unchanged file bytes without fixing private wrapper wording.
- Localized-defect fixture: duplicate and unresolved refs remain in generated files, unrelated output remains present, and generation exits `0`.
- Process fixture: an isolated fake repository runs the public command from an external current working directory with irrelevant host-state environment values. A separate missing-source run expects public exit `1`.
- Host-independence fixture: AST import inspection requires standard-library-only imports and rejects environment, current-working-directory, registry, DRMCP, and network-module inputs.
- Release-script fixture: the existing unittest discovery and generator invocations must each be followed by `if errorlevel 1 exit /b 1`.
- Git tracking correction: `.gitignore` previously ignored every nested `tools/` directory. Added directory traversal, re-ignore, and `*.py` re-inclusion rules so only Python files directly under `product/tests/tools/` appear in normal Git status. The `bin/` rule is unchanged.
- T07 did not change the generator or `scripts/verify.bat`.
- Scoped Git inspection before repository-local verification: whitespace result `pass`; T07 Task, parent Work Item, and both Python test modules appear as untracked worktree files. `.gitignore` and the pre-existing T06 `scripts/verify.bat` change appear as modified.
- Automated test result: `python -X utf8 -m unittest discover -s product\\tests\\tools -p "test_*.py" -v` exited `0`; all 35 tests passed in 0.776 seconds.
- Public generator result: `python -X utf8 product\\src\\tools\\generate_design_records_package.py` exited `0`; generation produced 34 files and 79 non-blocking warnings.
- Repository-local verification result: `scripts\\verify.bat` exited `0` and printed final `OK`.
- Generated-artifact ignore result: `git check-ignore -v bin/design-records/index.md` exited `0`.
- Verification correction: the first test run exposed one over-coupled assertion for the internal operational-error label. The fixture now asserts public stderr error presence and the stable `path does not exist` detail instead of the private wrapper operation name.
- Final verification summary: `tests_exit=0`, `generator_exit=0`, `verify_exit=0`, and `ignore_exit=0`.
- Independent review result: `NEEDS REVISION` with no blocking or major findings and two minor findings.
- F-MIN-01 correction: re-ignore all direct contents under `product/tests/tools/` before re-including only `*.py`; non-Python files remain ignored.
- F-MIN-02 correction: retain deterministic warning order and exact class/path checks, but project each detail into canonical refs, declaring path sets, and boundary group sets instead of matching private wrapper sentences.
- Post-correction verification passed: `tests_exit=0`, `generator_exit=0`, `verify_exit=0`, `bin_ignore_exit=0`, `nonpython_ignore_exit=0`, and `python_ignore_exit=1`.
- F-MIN-01 closure evidence: `bin/design-records/index.md` remains ignored; a non-Python probe under `product/tests/tools/` remains ignored; and the Python release test remains visible to Git.
- F-MIN-02 closure evidence: all 35 tests, the public generator, and `scripts\\verify.bat` passed after replacing private detail-sentence equality with semantic detail projection.
- Independent re-review verdict: `PASS`.
- Previous-finding disposition: F-MIN-01 `CLOSED`; F-MIN-02 `CLOSED`.
- Final finding summary: no blocking, major, minor, or advisory findings.
- Closure readiness: `READY FOR CLOSURE`.
- T07 is closed as `done`. T08 remains responsible for final producer/consumer review and handoff.

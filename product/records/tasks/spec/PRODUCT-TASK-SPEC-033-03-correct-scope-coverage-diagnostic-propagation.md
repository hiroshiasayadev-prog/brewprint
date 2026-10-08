# PRODUCT-TASK-SPEC-033-03: Correct scope coverage diagnostic propagation

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-033
- **task_type**: implementation
- **estimate**: 0.25d
- **depends_on**:
  - PRODUCT-TASK-SPEC-033-02
- **outputs**:
  - tools/usdm/usdm_tools.py
  - tools/usdm/tests

## Goal

Close REV-033-001 by propagating scope-relevant USDM and coverage scan errors through `check_usdm_scope_coverage`.

## Work

- Preserve all accepted hierarchical and effective-coverage semantics.
- Propagate relevant `scan_usdm` and `scan_coverage` error diagnostics into the scoped response.
- Ensure any propagated error diagnostic makes `ok: false`.
- Do not broaden or redefine scope matching.
- Add regression tests for malformed coverage and orphan hierarchy.
- Re-run scoped tests and current-corpus validation.

## Done condition

REV-033-001 is corrected with passing regression tests and no unrelated behavior change.

## Verification

At minimum:
- USDM tool tests;
- similarity tests as regression safety;
- current repository `validate_usdm`;
- scoped Git/whitespace inspection.

## Evidence

Files changed:
- `tools/usdm/usdm_tools.py`
- `tools/usdm/tests/test_hierarchical_coverage.py`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-033-03-correct-scope-coverage-diagnostic-propagation.md`

Verification:
- `python -m unittest discover -s tools/usdm/tests -p "test_*.py"` -> PASS, 36 tests, exit 0.
- `python -m unittest discover -s tools/usdm/similarity/tests -p "test_*.py"` -> PASS, 23 tests, exit 0.
- `python tools/usdm/usdm_tools.py validate_usdm --repo-root .` -> PASS, `ok: true`, 46 USDM records, 353 requirements, diagnostics `[]`, exit 0.
- Explicit reviewer reproductions via `python -m unittest -v tools.usdm.tests.test_hierarchical_coverage.ScopeCoverageTests.test_scope_propagates_malformed_coverage_diagnostic tools.usdm.tests.test_hierarchical_coverage.ScopeCoverageTests.test_scope_propagates_orphan_hierarchy_diagnostic` -> PASS, 2 tests, exit 0.
- Scoped `git.inspect_diff` over the three owned paths -> pass; changes limited to the existing hierarchical implementation plus this correction/test/task evidence.
- Scoped `git.inspect_worktree` with whitespace checking over the three owned paths -> pass, no whitespace findings. LF/CRLF conversion warnings are advisory only.

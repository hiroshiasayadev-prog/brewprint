# PRODUCT-TASK-SPEC-033-01: Implement hierarchical USDM tooling

- **status**: done
- **date**: 2026-09-30
- **work_item**: PRODUCT-WORK-SPEC-033
- **task_type**: implementation
- **estimate**: 1d
- **depends_on**:
- **outputs**:
  - tools/usdm/usdm_tools.py
  - tools/usdm/tests
  - tools/usdm/similarity/usdm_loader.py
  - tools/usdm/similarity/tests
  - tools/usdm-mcp/server.py

## Goal

Implement the accepted hierarchical USDM row identity and effective coverage behavior in repository-local tooling.

## Work

- Update row/full-ID parsing for `RNNN(-NN)*`.
- Validate immediate-parent existence inside one requirement record.
- Keep `RNNN-RNNN` ranges top-level only and support hierarchical comma tokens.
- Compute direct/derived/warning/blocking coverage states recursively.
- Project the exact additive response contract.
- Preserve cross-app full-ID direct coverage.
- Update similarity scope/load paths for hierarchical requirement IDs.
- Update the MCP wrapper for `include_warnings`.
- Add focused regression tests and run scoped verification.

## Implementation contract

| target | required change | acceptance criterion | verification |
|---|---|---|---|
| `tools/usdm/usdm_tools.py` | Hierarchical ID grammar, orphan validation, effective coverage evaluator, updated public responses. | Behavior matches artifact-format, coverage-format, and coverage-tools Specifications without changing their semantics. | USDM tool tests plus current-corpus `validate_usdm`. |
| `tools/usdm/tests/` | Hierarchy and coverage regression tests. | Tests cover child/grandchild IDs, orphan rejection, direct/derived/warning/blocking states, top-level ranges, hierarchical comma lists, and cross-app direct coverage. | Test suite passes. |
| `tools/usdm/similarity/usdm_loader.py` and tests | Accept hierarchical requirement IDs for row loading and exact requirement scopes. | Child rows are loadable/search-scope eligible without changing record/topic/app scope meaning. | Similarity loader tests pass without requiring live Qdrant/Ollama. |
| `tools/usdm-mcp/server.py` | Add `include_warnings` to scope coverage and align descriptions/examples. | MCP forwards the new parameter and remains backward compatible for omitted optional arguments. | Import/smoke or focused server test. |

## Done condition

All specified implementation surfaces conform to the reviewed design, focused tests pass, the current corpus validates, and this Task records exact test evidence.

## Verification

Run at minimum:

- existing and new tests under `tools/usdm/tests/`;
- applicable local tests under `tools/usdm/similarity/tests/` that do not require external live services;
- `validate_usdm` against the Brewprint repository;
- scoped Git diff/whitespace inspection for owned files.

## Evidence

Changed files:

- `tools/usdm/usdm_tools.py`
- `tools/usdm/tests/test_scope_coverage.py`
- `tools/usdm/tests/test_hierarchical_coverage.py`
- `tools/usdm/tests/test_mcp_wrapper_hierarchy.py`
- `tools/usdm/similarity/usdm_loader.py`
- `tools/usdm/similarity/tests/test_hierarchical_loader.py`
- `tools/usdm-mcp/server.py`
- `product/records/tasks/spec/PRODUCT-TASK-SPEC-033-01-implement-hierarchical-usdm-tooling.md`

Verification used a read-only temporary mirror of the exact owned source/test files because this session has repository filesystem/Git access but no arbitrary Windows command runner.

- `uv run --with pytest python -m unittest discover -s tools/usdm/tests -p 'test_*.py'`: exit 0; 34 tests run; OK.
- `uv run --with pytest python -m unittest discover -s tools/usdm/similarity/tests -p 'test_*.py'`: exit 0; 23 tests run; OK.
- `python3 -m py_compile tools/usdm/usdm_tools.py tools/usdm/similarity/usdm_loader.py tools/usdm-mcp/server.py`: exit 0.
- Current-corpus validation copied the 46 current `*/records/usdm/**/*.md` files into the same mirror and ran `python3 tools/usdm/usdm_tools.py validate_usdm --repo-root .`: exit 0; `ok: true`; `usdm_records: 46`; `requirements: 353`; `diagnostics: []`.
- The configured USDM MCP independently validated the live repository root with the same result: `ok: true`, 46 records, 353 requirements, no diagnostics.
- A preliminary plain-system-Python discovery run could not import the pre-existing migration test because that temporary environment lacked `pytest`; rerunning the complete suite through `uv run --with pytest` produced the passing 34-test result above.

Scoped Git/whitespace inspection covered only the owned paths. It reported exactly the eight files listed above as modified/untracked, no staged files, and no whitespace findings. Git emitted LF-to-CRLF conversion advisories only.

The post-Evidence semantic responsibility validator was not invoked: accepted `TRV-ADR-SPEC-006` suspends that validator and prohibits the deprecated prompt skill from acting as a completion gate.

No staging, commit, checkout, reset, restore, clean, or unrelated-file modification was performed. `PRODUCT-TASK-SPEC-033-02` was not executed.

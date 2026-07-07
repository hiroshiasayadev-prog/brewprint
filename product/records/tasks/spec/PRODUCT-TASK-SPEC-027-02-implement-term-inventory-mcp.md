# PRODUCT-TASK-SPEC-027-02: Implement term-inventory MCP

- **id**: PRODUCT-TASK-SPEC-027-02
- **status**: done
- **date**: 2026-07-04
- **work_item**: PRODUCT-WORK-SPEC-027
- **task_type**: implementation
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-027-01
- **outputs**:
  - tools/term-inventory-mcp/

## Goal

Implement one lightweight MCP that safely supports the term-inventory investigation without performing semantic extraction, aggregation, clustering, or use-case analysis.

## Work

- Follow the repository-local Python MCP pattern used by `tools/grep-mcp/`, `tools/ollama-mcp/`, and `tools/repo-ops-mcp/`.
- Expose bounded tools for manifest construction, deterministic batch partitioning, batch retrieval, constrained per-file result recording, and structural validation.
- Exclude this workflow's generated T03–T38 records and PRODUCT-INV-SPEC-011 from manifest enumeration so the inventory cannot observe its own execution artifacts.
- Restrict all source and output paths to configured allowed roots.
- Treat file size, modification time, and optional hashes as best-effort snapshot metadata.
- Report source drift without assuming that the original source remains immutable during parallel work.
- Generate observation IDs and schema-owned fields mechanically where possible.
- Reject malformed enums, missing required fields, duplicate file dispositions, duplicate observation IDs, and output-path escapes.
- Keep each batch's writer output separate from every other batch.
- Add focused unit tests and Windows start scripts.

The implementation must not call an LLM, identify candidate terms, infer meanings, merge batches semantically, or classify observations.

## Implementation contract

| target | required change | acceptance criterion | verification |
|---|---|---|---|
| `tools/term-inventory-mcp/server.py` | Implement `build_inventory_manifest`, `partition_inventory_manifest`, `get_inventory_batch`, `record_inventory_file_result`, and `validate_inventory_batch` as bounded MCP tools, including the accepted T03–T38 and PRODUCT-INV-SPEC-011 self-exclusion. | The server can enumerate only configured scope roots minus the explicit current-workflow exclusions, produce deterministic byte-balanced batches, return exact batch files, write only structurally conforming per-file results, and validate batch completeness and JSONL shape. | Unit tests call implementation functions directly and cover valid and invalid requests, including self-exclusion while retaining T02. |
| `tools/term-inventory-mcp/server.py` | Record best-effort snapshot metadata and source-drift diagnostics without requiring source immutability. | A changed source is reported as changed but does not by itself invalidate an otherwise conforming observation result; missing or unreadable sources remain explicit coverage outcomes. | Tests modify a fixture after manifest creation and verify the drift result. |
| `tools/term-inventory-mcp/pyproject.toml` and `uv.lock` | Define a self-contained Python MCP project using the same stable dependency line as existing repository MCP tools. | `uv sync` and the test command resolve from the project directory without global package installation. | Run project dependency resolution and unit tests. |
| `tools/term-inventory-mcp/start.ps1` and `start.cmd` | Start the stdio MCP behind the repository-local `mcp-proxy`, with configurable allowed root, output root, and port. | Windows launch scripts expose `/mcp`, `/sse`, and `/status` consistently with existing tools. | Start in check mode or run a smoke request through the proxy. |
| `tools/term-inventory-mcp/README.md` | Document tool schemas, safety boundary, snapshot semantics, output ownership, and example calls. | Another session can start the server and use every tool without reading implementation source. | Review examples against implemented function signatures. |
| `tools/term-inventory-mcp/tests/` | Cover scope authorization, deterministic ordering, byte-balanced partitioning, schema enums, observation-ID generation, no-candidate and failed results, duplicate dispositions, source drift, and output isolation. | All declared behaviors have at least one direct test and the suite passes. | `uv run python -m unittest discover -s tests -v` |

## Done condition

- All five MCP tools are implemented and documented.
- Structural validation enforces `bp-wide-term-observation-v1` and `bp-wide-term-batch-v1` requirements relevant to collection.
- Manifest snapshots tolerate concurrent source changes through explicit drift reporting.
- Parallel batches cannot write the same output file through the accepted interface.
- The focused unit-test suite and one proxy smoke check pass.
- No semantic extraction, aggregation, clustering, or use-case analysis is implemented.

## Verification

- Run `uv run python -m unittest discover -s tests -v` from `tools/term-inventory-mcp/`.
- Start the MCP through `start.ps1` in a bounded test configuration.
- Build a manifest over fixture roots, partition it, retrieve one batch, record one observation and one no-candidate result, then validate the batch.
- Change one source after manifest creation and confirm non-blocking drift diagnostics.
- Confirm an output escape and a duplicate source disposition are rejected.

## Evidence

- The user accepted a lightweight MCP and structural validator because unconstrained direct JSONL authoring could produce unusable investigation evidence.
- The user explicitly accepted best-effort source snapshots because other sessions may modify source artifacts during the investigation.
- `tools/term-inventory-mcp/plan_batches.py` is an early read-only planning aid and may be reused or replaced by the implementation.
- Aggregation and semantic tooling remain outside this Task and Work Item.
- DRMCP is non-operational. Filesystem authoring is the required fallback.
- Filesystem fallback created `server.py`, `pyproject.toml`, `uv.lock`, `README.md`, `start.ps1`, `start.cmd`, and `tests/test_server.py` under `tools/term-inventory-mcp/`.
- The implementation exposes all five required MCP tools and keeps semantic extraction, aggregation, clustering, classification, and use-case analysis outside the server.
- The focused unit-test suite covers authorization, traversal and symlink rejection, deterministic manifest creation, coherent-32 assignment, result recording, drift, output isolation, and batch validation.
- Post-T04 correction added explicit self-exclusion handling to `server.py` and `plan_batches.py`, plus a regression test that retains T02 while excluding T03, T38, and PRODUCT-INV-SPEC-011.
- The pre-correction unit-test verification passed on 2026-07-04. `uv run python -m unittest discover -s tests -v` ran 16 tests in 3.928 seconds with `OK (skipped=1)`; the single skip was the symlink fixture because the Windows account lacks symlink-creation privilege. The amended regression suite was not re-executed in this coordination session because the connected MCP process remained loaded with the prior server module.
- Dependency resolution passed on 2026-07-04. `uv lock` resolved 34 packages, including `mcp 1.28.1` and `mcp-proxy 0.12.0`, and regenerated `uv.lock` from `pyproject.toml`.
- Proxy smoke passed on 2026-07-04 through `http://127.0.0.1:8934/mcp`. The client completed MCP initialization and `list_tools()` returned exactly `build_inventory_manifest`, `partition_inventory_manifest`, `get_inventory_batch`, `record_inventory_file_result`, and `validate_inventory_batch`.
- Scoped Git whitespace inspection passed for the visible Task change. `tools/term-inventory-mcp/` remains hidden by the repository `.gitignore` rule.
- Stage and commit were not performed.

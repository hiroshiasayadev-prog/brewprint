# PRODUCT-TASK-SPEC-029-05: Implement USDM MCP wrapper

- **id**: PRODUCT-TASK-SPEC-029-05
- **status**: done
- **date**: 2026-07-09
- **work_item**: PRODUCT-WORK-SPEC-029
- **task_type**: implementation
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-029-03
  - PRODUCT-TASK-SPEC-029-04
- **outputs**:
  - `tools/usdm-mcp/`
  - `.gitignore`
  - PRODUCT-TASK-SPEC-029-05
  - PRODUCT-WORK-SPEC-029

## Goal

Expose the standalone USDM tools through a repository-local MCP server.

## Boundary

This Task owns only the MCP wrapper around the existing standalone USDM tool implementation.

This Task does not own changes to USDM validation semantics, coverage semantics, real USDM records, smoke fixtures, DRMCP integration, existing-record migration, or PRODUCT-WORK-SPEC-029 closure.

## Work

- Added `tools/usdm-mcp/server.py` as a FastMCP stdio server.
- Added `tools/usdm-mcp/pyproject.toml` with MCP runtime dependencies.
- Added `tools/usdm-mcp/start.ps1` for local `mcp-proxy` startup.
- Added `tools/usdm-mcp/README.md` with the startup command shape.
- Updated `.gitignore` so `tools/usdm-mcp/` is trackable while other ignored tool output remains ignored.
- Inserted this Task before the Domain handoff task candidate.

## Implementation contract

| target | required change | acceptance criterion | verification |
|---|---|---|---|
| `tools/usdm-mcp/server.py` | Provide a FastMCP server exposing `validate_usdm`, `check_usdm_coverage`, and `usdm_covered_by`. | Each MCP tool delegates to the existing `tools/usdm/usdm_tools.py` implementation and returns the same JSON-compatible response shape. | Static inspection and MCP runtime calls through `@usdm-mcp`. |
| `tools/usdm-mcp/start.ps1` | Provide a local `mcp-proxy` startup wrapper. | The script accepts root, host, port, and stateless options and passes `USDM_MCP_ROOT` to the server. | Static comparison against `tools/grep-mcp/start.ps1`. |
| `.gitignore` | Allow normal tracking of `tools/usdm-mcp/`. | `tools/usdm-mcp/` is unignored without unignoring all `tools/` output. | Scoped git inspection. |
| PRODUCT records | Register this Task and preserve the later handoff as a separate candidate. | PRODUCT-WORK-SPEC-029 lists PRODUCT-TASK-SPEC-029-05 and renames the handoff candidate to PRODUCT-TASK-SPEC-029-06. | Scoped git inspection. |

## Done condition

- `tools/usdm-mcp/` contains an MCP server, project file, startup script, and README.
- The MCP server exposes the three USDM operations.
- The MCP server restricts `repo_root` to paths inside `USDM_MCP_ROOT`.
- `.gitignore` permits `tools/usdm-mcp/` tracking.
- PRODUCT-WORK-SPEC-029 registers this Task and preserves the handoff route as a later candidate.

## Verification

Static verification performed:

| check | result |
|---|---|
| `tools/usdm-mcp/server.py` imports `FastMCP` and defines three `@mcp.tool()` functions. | PASS |
| `server.py` resolves `USDM_MCP_ROOT` and rejects `repo_root` values escaping that root. | PASS |
| `server.py` delegates to `tools/usdm/usdm_tools.py` functions. | PASS |
| `start.ps1` follows the existing `tools/grep-mcp/start.ps1` `mcp-proxy` shape with `USDM_MCP_ROOT`. | PASS |
| `.gitignore` unignores `tools/usdm-mcp/` and its contents. | PASS |

Runtime verification through `@usdm-mcp` executed after local MCP startup.

| check | result |
|---|---|
| `@usdm-mcp.validate_usdm` with default arguments | PASS; returned `ok: true`, `usdm_records: 0`, `requirements: 0`, and no diagnostics. |
| `@usdm-mcp.check_usdm_coverage` with default arguments | PASS; returned `ok: true`, `requirements: 0`, `covered: 0`, empty `uncovered`, empty `dangling`, and no diagnostics. |

## Evidence

- The MCP startup command is documented in `tools/usdm-mcp/README.md`.
- The direct command shape is:

```bat
start "usdm-mcp" cmd /k "mcp-proxy --host=127.0.0.1 --port=8184 --cwd=C:\Users\imved\projects\brewprint --env USDM_MCP_ROOT C:\Users\imved\projects\brewprint -- uv run --project C:\Users\imved\projects\brewprint\tools\usdm-mcp python C:\Users\imved\projects\brewprint\tools\usdm-mcp\server.py"
```

- MCP runtime calls through `@usdm-mcp` confirmed the wrapper exposes `validate_usdm` and `check_usdm_coverage`.
- The PowerShell wrapper shape is:

```powershell
.\tools\usdm-mcp\start.ps1 -Root C:\Users\imved\projects\brewprint -Port 8184
```

- No files were staged, committed, or pushed.

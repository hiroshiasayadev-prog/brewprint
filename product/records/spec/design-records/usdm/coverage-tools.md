# Contract: USDM coverage tools

- **id**: `spec:product.design_records.usdm.coverage_tools`
- **status**: draft
- **date**: 2026-07-08
- **parent**: `spec:product.design_records.usdm`
- **contract_class**: `interface`

## What this is

This contract defines standalone MVP tools for validating USDM records and checking USDM coverage.

The tools live under `tools/usdm/` during the MVP. They may later move behind DRMCP or another MCP surface.

## Request

### Shared request fields

| field | required | meaning |
|---|---:|---|
| `repo_root` | yes | Repository root to scan. |
| `app_namespace` | no | Optional app namespace filter. |

When `app_namespace` is omitted, the tool may scan every discovered app namespace that has `records/usdm/` or `records/spec/`.

### `validate_usdm`

| field | required | meaning |
|---|---:|---|
| `repo_root` | yes | Repository root. |
| `app_namespace` | no | Optional app namespace filter. |

### `check_usdm_coverage`

| field | required | meaning |
|---|---:|---|
| `repo_root` | yes | Repository root. |
| `app_namespace` | no | Optional app namespace filter. |
| `include_dangling` | no | When true, include dangling `usdm_covers` entries. Default is true. |

### `usdm_covered_by`

| field | required | meaning |
|---|---:|---|
| `repo_root` | yes | Repository root. |
| `requirement_id` | yes | Full USDM requirement ID. |

## Response

### Operation set

| operation | purpose |
|---|---|
| `validate_usdm` | Validate USDM record format, USDM IDs, row IDs, duplicate full requirement IDs, and malformed coverage entries. |
| `check_usdm_coverage` | List USDM requirement IDs not covered by any discovered implementation Specification. |
| `usdm_covered_by` | List implementation Specification refs that cover one full USDM requirement ID. |

The tools are repository-static. They do not execute implementation code.

### Shared diagnostic fields

| field | meaning |
|---|---|
| `category` | Machine-readable issue category. |
| `severity` | `error`, `warning`, or `info`. |
| `path` | Repository-relative path when available. |
| `message` | Human-readable summary. |
| `value` | Offending value when useful. |

### `validate_usdm` response

| field | meaning |
|---|---|
| `ok` | False when any diagnostic has severity `error`. |
| `usdm_records` | Number of discovered USDM records. |
| `requirements` | Number of discovered full USDM requirement IDs. |
| `diagnostics` | Format, ID, duplicate, and malformed coverage diagnostics. |

### `check_usdm_coverage` response

| field | meaning |
|---|---|
| `ok` | False when uncovered requirements or dangling coverage entries exist. |
| `requirements` | Number of discovered full USDM requirement IDs. |
| `covered` | Number of USDM requirement IDs covered by at least one Specification. |
| `uncovered` | Full USDM requirement IDs not covered by any discovered Specification. |
| `dangling` | Coverage entries that reference missing full USDM requirement IDs. |
| `diagnostics` | Tool diagnostics for malformed inputs or scan failures. |

### `usdm_covered_by` response

| field | meaning |
|---|---|
| `ok` | False when the requested requirement ID is malformed or missing. |
| `requirement_id` | Requested full USDM requirement ID. |
| `exists` | Whether the requirement exists in discovered USDM records. |
| `covered_by` | Specification refs that declare the requirement ID in `usdm_covers`. |
| `diagnostics` | Tool diagnostics for malformed, missing, or scan-failure states. |

## Errors

| condition | handling |
|---|---|
| `repo_root` is missing or unreadable | Return `ok: false` and an error diagnostic. |
| `requirement_id` is malformed | Return `ok: false` and an error diagnostic. |
| USDM record cannot be parsed enough to determine metadata | Return an error diagnostic from `validate_usdm`. |
| Specification metadata cannot be parsed enough to inspect `usdm_covers` | Return an error diagnostic for coverage operations. |
| No USDM records are discovered | Return `ok: true` for `validate_usdm` with zero counts; coverage tools return zero counts unless a specific missing requirement was requested. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.usdm` | Parent overview. |
| `spec:product.design_records.usdm.artifact_format` | Defines USDM record and requirement ID format. |
| `spec:product.design_records.usdm.coverage_format` | Defines `usdm_covers` metadata and coverage semantics. |
| PRODUCT-REQ-SPEC-015 | Source requirement. |

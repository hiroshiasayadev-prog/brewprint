# Contract: USDM coverage tools

- **id**: `spec:product.design_records.usdm.coverage_tools`
- **status**: draft
- **date**: 2026-09-30
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

### `check_usdm_scope_coverage`

| field | required | meaning |
|---|---:|---|
| `repo_root` | yes | Repository root. |
| `scope_ids` | yes | USDM app, topic, record, or full requirement IDs to report. |
| `include_covered` | no | When true, include direct and derived covered row IDs. Directly covered rows include covering Specification refs. Default is true. |
| `include_not_covered` | no | When true, include blocking uncovered row IDs. Default is true. |
| `include_warnings` | no | When true, include non-blocking refinement-warning row IDs. Default is true. |
| `include_empty_records` | no | When true, include record items with no visible coverage fields. Default is false. |

## Response

### Operation set

| operation | purpose |
|---|---|
| `validate_usdm` | Validate USDM record format, hierarchy, row IDs, duplicate full requirement IDs, and malformed coverage entries. |
| `check_usdm_coverage` | Evaluate direct and derived coverage and report blocking uncovered requirements plus refinement warnings. |
| `usdm_covered_by` | Report one requirement's effective coverage state and its direct covering Specification refs. |
| `check_usdm_scope_coverage` | Return compact record-grouped direct, derived, blocking, and warning coverage for selected scopes. |

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
| `ok` | False when blocking uncovered requirements, dangling coverage entries, or error diagnostics exist. Warning-only refinement gaps do not make `ok` false. |
| `requirements` | Number of discovered full USDM requirement IDs. |
| `covered` | Number of effectively covered requirement IDs, including direct and derived coverage. |
| `direct_covered` | Number of requirement IDs directly listed by at least one Specification. |
| `derived_covered` | Number of non-leaf requirement IDs covered only because every direct child is effectively covered. |
| `uncovered` | Full requirement IDs that are effectively uncovered and have no directly covered ancestor. |
| `refinement_warnings` | Full requirement IDs that are effectively uncovered but have a directly covered ancestor. |
| `dangling` | Coverage entries that reference missing full USDM requirement IDs. |
| `diagnostics` | Tool diagnostics for malformed inputs or scan failures. |

### `usdm_covered_by` response

| field | meaning |
|---|---|
| `ok` | False when the requested requirement ID is malformed or missing, or coverage evaluation has an error. |
| `requirement_id` | Requested full USDM requirement ID. |
| `exists` | Whether the requirement exists in discovered USDM records. |
| `coverage_state` | `direct`, `derived`, `warning`, `uncovered`, or null when the row does not exist or coverage evaluation is unavailable. Direct coverage takes precedence over derived coverage. |
| `covered_by` | Specification refs that directly declare or compactly expand to the requirement ID in `usdm_covers`. Empty for `derived`, `warning`, `uncovered`, missing, or unavailable states. |
| `direct_covered_ancestor` | The nearest directly covered ancestor full requirement ID when `coverage_state` is `warning`; otherwise null. |
| `diagnostics` | Tool diagnostics for malformed, missing, or scan-failure states. |

All fields above are present in the response.
Malformed or missing IDs use `exists: false`, `coverage_state: null`, an empty `covered_by` list, and `direct_covered_ancestor: null` together with diagnostics.

### `check_usdm_scope_coverage` response

| field | meaning |
|---|---|
| `ok` | False when scope resolution, blocking uncovered requirements, or error diagnostics exist. Warning-only refinement gaps do not make `ok` false. |
| `scope_ids` | Requested USDM scope IDs. |
| `records` | Number of expanded USDM requirement records. |
| `requirements` | Number of expanded full USDM requirement IDs. |
| `covered_requirements` | Number of effectively covered requirement IDs. |
| `direct_covered_requirements` | Number of directly covered requirement IDs. |
| `derived_covered_requirements` | Number of derived-covered requirement IDs. |
| `not_covered_requirements` | Number of blocking uncovered requirement IDs. |
| `refinement_warning_requirements` | Number of non-blocking refinement-warning requirement IDs. |
| `items` | Record-grouped compact coverage report. |
| `diagnostics` | Tool diagnostics for malformed inputs or scan failures. |

Each item uses `record_id` as the grouping key.

Item fields:

| field | type | presence | meaning |
|---|---|---|---|
| `record_id` | string | always | Full USDM requirement-record ID. |
| `covered` | object from compact row ID to list of Specification refs | when `include_covered` is true and at least one direct-covered row is visible, or when `include_empty_records` requires an otherwise empty item | Existing direct-coverage map; retained for compatibility. |
| `derived_covered` | list of compact row IDs | when `include_covered` is true and at least one derived-covered row is visible, or when `include_empty_records` requires an otherwise empty item | Rows effectively covered only through complete direct-child coverage. |
| `not_covered` | list of compact row IDs | when `include_not_covered` is true and at least one blocking uncovered row is visible, or when `include_empty_records` requires an otherwise empty item | Blocking uncovered rows. |
| `refinement_warnings` | list of compact row IDs | when `include_warnings` is true and at least one warning row is visible, or when `include_empty_records` requires an otherwise empty item | Effectively uncovered descendants beneath a directly covered ancestor. |

When `include_empty_records` forces an otherwise empty item to appear, every enabled category field is present with an empty object or list as appropriate.
Disabled categories are omitted even on empty items.

A row appears in exactly one effective coverage category.
Direct coverage takes precedence over derived coverage.
The `covered` map always represents direct coverage only; derived coverage never fabricates Specification refs.

## Errors

| condition | handling |
|---|---|
| `repo_root` is missing or unreadable | Return `ok: false` and an error diagnostic. |
| `requirement_id` is malformed | Return `ok: false` and an error diagnostic. |
| `scope_ids` is empty or contains an invalid scope ID | Return `ok: false` and an error diagnostic. |
| USDM record cannot be parsed enough to determine metadata | Return an error diagnostic from `validate_usdm`. |
| Specification metadata cannot be parsed enough to inspect `usdm_covers` | Return an error diagnostic for coverage operations. |
| No USDM records are discovered | Return `ok: true` for `validate_usdm` with zero counts; coverage tools return zero counts unless a specific missing requirement was requested. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.usdm` | Parent overview. |
| `spec:product.design_records.usdm.artifact_format` | Defines USDM record and requirement ID format. |
| `spec:product.design_records.usdm.coverage_format` | Defines `usdm_covers` metadata and coverage semantics. |
| PRODUCT-REQ-SPEC-015 | Original MVP source requirement. |
| PRODUCT-REQ-SPEC-016 | Hierarchical decomposition and derived-coverage source requirement. |
| PRODUCT-ADR-SPEC-020 | Hierarchical coverage reporting decision. |

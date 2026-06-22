# Contract: `validate_records`

- **id**: `spec:drmcp.design_records_mcp.tools.validate_records`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`validate_records` verifies that the Design Records MCP metadata index is in a trustworthy state. In addition to basic record metadata integrity checks, it validates active `spec:` semantic refs, record public ID-as-refs (with namespace_prefix; MVP: `V01-ADR-*` / `V01-SPEC-*` / `V01-INV-*` / `V01-REQ-*` / `V01-WORK-*` / `V01-TASK-*`), resolvability of investigation `source_refs` / `follow_up_results` as canonical references, and declared workflow relation integrity. When `follow_up_candidates` contain artifact references, their canonical form is validated.

For workflow artifacts, status-gated required narrative section validation is also performed: only section heading existence and non-empty body content are checked; body quality, sufficiency, and semantic validity are not evaluated.

MVP excluded scope per V01-ADR-088 / V01-ADR-092: `internal-design:` / `coverage:` / `COV-*`, semantic realization relations, coverage mapping queries, orphan workflow artifact diagnostics, progress projection, and workflow traversal queries. Investigation metadata canonical workflow reference extension is limited to `REQ-*` / `WORK-*`; `TASK-*` in investigation metadata fields is an unsupported reference.

> Source: V01-ADR-077 §validate_records の責務

## Request

```json
{
  "kind": "decision",
  "id_range": {
    "from": "V01-ADR-067",
    "to": "V01-ADR-077"
  }
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `kind` | no | string | Restrict validation to this record kind |
| `id_range` | no | object | Restrict validation to this ID range |

When request is empty, all MVP-indexed records are validated.

`id_range` endpoint family, effective `kind`, one-sided range, and unsupported range rules follow `list_records`. Workflow artifact ID ranges in public ID form (`V01-REQ-<DOMAIN>-NNN` / `V01-WORK-<DOMAIN>-NNN` / `V01-TASK-<DOMAIN>-NNN-MM`) are supported. `SPEC-*` / `INV-*` range, mixed family, mixed domain, mixed task work sequence, malformed endpoints, and `kind`/endpoint family mismatch are request errors.

## Response

```json
{
  "ok": false,
  "diagnostics": [
    {
      "category": "missing_required_section",
      "severity": "error",
      "record_id": "V01-WORK-MCP-014",
      "path": "v01/records/work-items/mcp/V01-WORK-MCP-014-normalize-propose-record-create-id-fields-body-contract.md",
      "message": "required section \"Boundary\" must be non-empty when work_item status is \"done\"",
      "section": "Boundary",
      "status": "done"
    }
  ]
}
```

| field | meaning |
|---|---|
| `ok` | `true` when no error-severity diagnostics are present |
| `diagnostics` | Diagnostic list |

Standard diagnostic object fields:

| field | required | meaning |
|---|---:|---|
| `category` | yes | Diagnostic category |
| `severity` | yes | `error`, `warning`, or `info` |
| `record_id` | no | ID of the record containing the issue |
| `path` | no | Path of the record containing the issue |
| `message` | yes | Human-readable message |
| `target_id` | no | Target ID for broken reference diagnostics |

Severity effects:
- `error`: sets `validation.ok` to `false`; blocks retained proposal creation.
- `warning`: returned in retained proposals but does not set `validation.ok` to `false` and does not block proposal creation.
- `info`: does not set `validation.ok` to `false`.

### Diagnostic categories

MVP diagnostic categories and their required additional fields are defined in `spec:drmcp.design_records_mcp.schema.diagnostics`. Key categories include:

- Metadata integrity: `duplicate_id`, `filename_id_mismatch`, `invalid_h1_title`, `invalid_workflow_id`, `missing_required_metadata`, `empty_required_metadata`, `invalid_metadata_value`, `invalid_status_for_kind`, `spec_status_mismatch`, `missing_depends_on_target`, `missing_supersedes_target`, `invalid_migrated_to_spec`, `missing_record_path`
- Required narrative sections: `missing_required_section`, `empty_required_section`, `section_heading_case_mismatch`
- Spec semantic ref: `invalid_semantic_ref_declaration`, `missing_section_target`, `ambiguous_section_target`, `duplicate_semantic_ref`
- Investigation refs: `unresolved_source_ref`, `unresolved_follow_up_result`, `unresolved_follow_up_candidate`, `noncanonical_source_ref`, `noncanonical_follow_up_result`, `noncanonical_follow_up_candidate`, `unsupported_reference`
- Workflow relations: `unresolved_workflow_relation`, `invalid_workflow_relation_target`, `workflow_relation_mismatch`, `workflow_source_requirement_mismatch`

Valid status values by record kind:

| record kind | valid status values |
|---|---|
| `decision` | `proposed`, `accepted`, `superseded` |
| `spec` | `confirmed`, `draft`, `wip` |
| `investigation` | `investigating`, `concluded`, `superseded` |
| `requirement` | `captured`, `decision_needed`, `accepted`, `deferred`, `rejected` |
| `work_item` | `not_started`, `in_progress`, `blocked`, `done` |
| `task` | `not_started`, `in_progress`, `blocked`, `done` |

`depends_on` can reference `ADR-*` / `SPEC-*` / `INV-*` canonical record ID-as-refs. Therefore `V01-ADR-086` with `depends_on: V01-INV-DOCS-001` is contract-valid. A `missing_depends_on_target` for that reference before investigation index integration is implemented is a known implementation gap that must resolve after M19 implementation.

`ok` is `true` when there are no `error` diagnostics. `warning` and `info` diagnostics do not set `ok` to `false`.

### Authoring-specific diagnostics

When `validate_records` is used in the context of authoring proposals (proposal-local validation), the following additional diagnostic categories from the authoring transaction model also appear. Full definitions are in `spec:drmcp.design_records_mcp.tools.propose_record_update` and `spec:drmcp.design_records_mcp.tools.propose_record_create`.

- `no_op_update`: `propose_record_update` result produces no content change.
- `reciprocal_follow_up_mode_required`: Required reciprocal update exists but was not included.
- `conflicting_operations`: Two operations in `operations` array target the same metadata field.
- `multiple_section_replace_not_supported`: Two or more `named_section_replace` operations in one `operations` array.
- `missing_required_metadata_batch`: Multiple required fields absent on create.

Authoring diagnostic additional fields:

| field | applicable categories | meaning |
|---|---|---|
| `allowed_values` | `invalid_metadata_value` (status) | Permitted status values for the target kind |
| `required_fields` | `missing_required_metadata_batch` | List of absent required field names |
| `target_kind` | `missing_required_metadata_batch` | Record kind defining the field requirements |
| `repair_suggestion` | `invalid_metadata_value`, `reciprocal_follow_up_mode_required` | Advisory minimal fix; caller must verify before applying; not automatically applied |

## Errors

| code | condition |
|---|---|
| `invalid_request` | Invalid `kind` or other malformed field |
| `invalid_id_range` | Malformed or unsupported `id_range` (same rules as `list_records`) |

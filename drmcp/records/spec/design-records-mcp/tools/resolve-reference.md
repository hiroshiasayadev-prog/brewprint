# Contract: `resolve_reference`

- **id**: `spec:drmcp.design_records_mcp.tools.resolve_reference`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`resolve_reference` resolves an MVP canonical reference to a single document / section / record target. Validation uses the same resolver lookup rules and must not have independent resolution logic.

## Request

```json
{
  "ref": "spec:trace.semantic-ref.definition"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `ref` | yes | string | Canonical reference candidate to resolve. Leading and trailing whitespace are not permitted; the input string is evaluated as-is. |

Supported input forms. Record ID-as-ref uses the full ID as returned by the index (with namespace_prefix). MVP: `namespace_prefix = V01-`.

| input form | ref kind | lookup source |
|---|---|---|
| Active `spec:` document-level ref | `semantic_ref` | Spec front matter `semantic_refs` |
| Active `spec:` section-level ref | `semantic_ref` | Spec front matter `sections` |
| `<namespace_prefix>ADR-NNN` (e.g. `V01-ADR-097`) | `record_id` | `decision` record index |
| `<namespace_prefix>SPEC-<slug>` (e.g. `V01-SPEC-design-records-mcp-overview`) | `record_id` | `spec` record index |
| `<namespace_prefix>INV-<DOMAIN>-NNN` (e.g. `V01-INV-MCP-001`) | `record_id` | `investigation` record index |
| `<namespace_prefix>REQ-<DOMAIN>-NNN` (e.g. `V01-REQ-MCP-001`) | `record_id` | `requirement` record index |
| `<namespace_prefix>WORK-<DOMAIN>-NNN` (e.g. `V01-WORK-DRMCP-001`) | `record_id` | `work_item` record index |
| `<namespace_prefix>TASK-<DOMAIN>-<WRK>-<TSK>` (e.g. `V01-TASK-MCP-001-01`) | `record_id` | `task` record index |

Workflow ID grammar: requirement / work item sequence and task work sequence are 3-digit zero-padded; task sequence is 2-digit zero-padded.

`internal-design:` / `coverage:`, `COV-*`, physical paths, and grammar-invalid ID forms are not supported inputs. Direct query returns `status: "unsupported"` (not a tool execution error). `yaml:` is a reserved prefix; MVP does not define public resolver input or direct query response behavior for it.

## Response

All MVP direct query responses include these top-level fields:

| field | required | meaning |
|---|---:|---|
| `ref` | yes | Input string from the request |
| `ref_kind` | yes | `semantic_ref` / `record_id` / `unsupported` |
| `status` | yes | `resolved` / `unresolved` / `unsupported` |
| `target` | yes | Target object when `resolved`; `null` otherwise |
| `diagnostics` | yes | Resolution diagnostic list; empty on successful resolution |

Resolved section-level `spec:` example:

```json
{
  "ref": "spec:trace.semantic-ref.definition",
  "ref_kind": "semantic_ref",
  "status": "resolved",
  "target": {
    "target_type": "section",
    "path": "v01/records/spec/concepts/traceability/semantic-ref.md",
    "section": "Semantic ref definition"
  },
  "diagnostics": []
}
```

- Resolved document-level `spec:` target: `target_type: "document"`, `path`. No `section` field.
- Resolved section-level `spec:` target: `target_type: "section"`, `path`, `section`. Input canonical ref is held in top-level `ref` and is not repeated in `target`.
- MVP does not define a public parent-child relationship between section-level and document-level refs, and does not infer a parent document ref from a section-level ref's string prefix.
- Resolved record ID-as-ref target: `target_type: "record"`, `path`, `record_id`, `record_kind`, `title`, `status`.

When a supported form target does not exist: `status: "unresolved"`, `target: null`, `diagnostics` includes `unresolved_reference`.

When the same `spec:` ref or record ID resolves to multiple targets, one must not be arbitrarily chosen. Returns `status: "unresolved"`, `target: null`, `diagnostics` includes `ambiguous_reference`. Validation reports the same cause as `duplicate_semantic_ref` or `duplicate_id` error.

Unsupported example:

```json
{
  "ref": "internal-design:resolver.semantic-ref-index",
  "ref_kind": "unsupported",
  "status": "unsupported",
  "target": null,
  "diagnostics": [
    {
      "category": "unsupported_reference",
      "severity": "info",
      "message": "reference form is outside the MVP resolver contract"
    }
  ]
}
```

`unsupported_reference` in a direct query is `info` severity — it indicates an input boundary, not a resolver failure. When an unsupported input appears in an investigation metadata validation target field, the severity follows the `validate_records` contract (see `spec:drmcp.design_records_mcp.schema.diagnostics`).

## Errors

No tool execution errors are defined for this tool beyond `invalid_request`. Unsupported and unresolved inputs produce diagnostic responses, not tool errors.

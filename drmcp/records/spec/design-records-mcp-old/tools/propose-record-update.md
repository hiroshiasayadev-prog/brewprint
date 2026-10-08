# Contract: `propose_record_update`

- **id**: `spec:drmcp.design_records_mcp.tools.propose_record_update`
- **status**: draft
- **date**: 2026-06-28
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`propose_record_update` creates a retained proposal for a single-operation update (whole metadata block replacement, field-level metadata patch, or named Markdown section replacement) or an atomic multi-operation update combining multiple supported operations into one retained proposal for the same record. It does not write repository files.

A no-op update (where the proposed persisted content would be byte-equivalent to the current file content) MUST NOT create a retained proposal.

MVP update support: `decision`, `spec`, `requirement`, `work_item`, `task`. Update operations reject any ID containing the literal `new` token in the sequence position.

## Request

Metadata block replacement:

```json
{
  "kind": "task",
  "id": "V01-TASK-MCP-008-04",
  "update": {
    "type": "metadata_block_replace",
    "metadata": {
      "id": "V01-TASK-MCP-008-04",
      "status": "done",
      "date": "2026-06-01",
      "work_item": "V01-WORK-MCP-008",
      "source_requirement": "V01-REQ-MCP-008",
      "estimate": "1d-2d",
      "depends_on": ["V01-TASK-MCP-008-03"],
      "outputs": ["Updated V01-SPEC-design-records-mcp-tools"]
    }
  }
}
```

Named section replacement:

```json
{
  "kind": "task",
  "id": "V01-TASK-MCP-008-04",
  "update": {
    "type": "named_section_replace",
    "section_selector": {
      "heading": "Evidence",
      "match": "exact"
    }
  },
  "body": "2026-06-02: Spec reflection completed.\n"
}
```

Metadata field replacement:

```json
{
  "kind": "task",
  "id": "V01-TASK-MCP-020-02",
  "update": {
    "type": "metadata_fields_replace",
    "metadata": {
      "status": "done"
    }
  }
}
```

Operations array:

```json
{
  "kind": "task",
  "id": "V01-TASK-MCP-008-04",
  "operations": [
    {
      "type": "metadata_fields_replace",
      "metadata": { "status": "done" }
    },
    {
      "type": "named_section_replace",
      "section_selector": { "heading": "Evidence", "match": "exact" },
      "body": "2026-06-07: Implementation verified.\n"
    }
  ]
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `kind` | yes | string | Update target kind |
| `id` | yes | string | Exact existing record public ID. `new` placeholder is invalid. |
| `update` | conditional | object | Single update operation object. Mutually exclusive with `operations`. |
| `operations` | conditional | array | Atomic multi-operation list. Mutually exclusive with `update`. |
| `body` | conditional | string | Replacement Markdown body for `named_section_replace` when using `update` |
| `body_cache_id` | conditional | string | Cached replacement body for `named_section_replace` when using `update` |
| `diff_mode` | no | string | Diff return format: `summary` (default) / `patch` / `none` |

Exactly one of `update` or `operations` must be present. Supplying both returns `invalid_request` and creates no proposal. Empty `operations` array returns `invalid_request` and creates no proposal.

### update.type values

| value | meaning |
|---|---|
| `metadata_block_replace` | Replace the kind-specific metadata block as a whole |
| `metadata_fields_replace` | Patch one or more metadata fields; preserve unspecified existing fields |
| `named_section_replace` | Replace exactly one Markdown section as a whole |

### Metadata block replacement

Targets the kind-specific metadata block:

| kind | replacement target |
|---|---|
| `spec` | Recognized spec metadata fields inside YAML front matter |
| `decision` | H1-following ADR bullet metadata block |
| `requirement` | H1-following requirement bullet metadata block |
| `work_item` | H1-following work item bullet metadata block |
| `task` | H1-following task bullet metadata block |

For `spec`, replacement is scoped to recognized fields only. Unknown or auxiliary YAML front matter fields must be preserved. Recognized spec metadata fields: `scope`, top-level `status`, and `design_record.id` / `design_record.kind` / `design_record.status` / `design_record.depends_on`.

Decision metadata replacement validates recognized ADR fields: `status`, `date`, `depends_on`, `supersedes`, `migrated_to_spec`.

Workflow artifact metadata replacement validates required fields per `spec:drmcp.design_records_mcp.schema.metadata_grammar`.

Proposal-local validation uses the shared repository taxonomy from `spec:drmcp.design_records_mcp.schema.diagnostics`. Missing required fields produce `missing_required_field`. A present empty value produces `empty_required_field` only when the applicable PRODUCT authority requires non-empty content. Invalid present values produce `invalid_field_value`. The resulting entries use structured `subject`, `field`, `value`, and applicable `target` or `location` associations.

Spec metadata block replacement example:

```json
{
  "kind": "spec",
  "id": "V01-SPEC-design-records-mcp-tools",
  "update": {
    "type": "metadata_block_replace",
    "metadata": {
      "scope": "drmcp/records/spec/design-records-mcp/tools.md",
      "status": "draft",
      "design_record": {
        "id": "V01-SPEC-design-records-mcp-tools",
        "kind": "spec",
        "status": "draft",
        "depends_on": ["V01-ADR-076", "V01-ADR-077", "V01-ADR-087", "V01-ADR-088", "V01-ADR-090", "V01-ADR-092", "V01-ADR-093"]
      }
    }
  }
}
```

### Metadata field replacement

Reads the current record metadata, applies only the caller-supplied field changes, and preserves all unspecified existing fields.

`update.metadata` must contain only the fields to be patched. After applying the patch, the complete result is validated using the same rules as `metadata_block_replace`.

`body` and `body_cache_id` must be omitted for `metadata_fields_replace`. Supplying either returns `invalid_body_source` and creates no proposal.

### Named section replacement

Valid only when `section_selector` resolves to exactly one Markdown ATX section in the target record. Section matching uses the same ATX heading source rules as the `headings` field in the record model; YAML front matter and fenced code block content are not section sources; setext headings are not section sources in MVP.

`section_selector` fields:

| field | required | meaning |
|---|---:|---|
| `heading` | yes | Heading text to match |
| `match` | no | MVP: `exact` only |
| `level` | no | Optional ATX heading level constraint |

Exact matching compares heading text after removing ATX marker syntax and trimming whitespace. Case-sensitive. No Unicode normalization, punctuation folding, or prefix/contains matching in MVP. When `level` is supplied, both text and level must match.

**Narrow case-only fallback** (workflow artifacts only):

- Applies to `requirement`, `work_item`, `task` only.
- Applies only when `section_selector.heading` is a canonical heading whose required presence or gated substantive-content requirement is declared for the target kind by the applicable PRODUCT authority: `spec:product.design_records.authoring_standards.requirement_authoring`, `spec:product.design_records.authoring_standards.work_item_authoring`, or `spec:product.design_records.authoring_standards.task_authoring`.
- The target record does not need to currently be in the gated status; record kind and requested heading determine eligibility.
- Attempted only after exact matching finds zero matches.
- Case-insensitive comparison; no Unicode normalization, punctuation folding, typo correction, prefix, or contains matching.
- When `level` is supplied, the same level constraint applies before determining zero, one, or ambiguous case-insensitive matches.
- If exactly one eligible case-insensitive match: selector resolves through fallback; proposal creation proceeds unless an independent error applies.
- When proceeding through fallback, the proposal diff MUST rewrite the matched heading line to the canonical `section_selector.heading` text.
- If multiple case-insensitive matches: `section_selector_ambiguous`; no proposal created.
- Non-case differences are governed by exact selector rules and must not use this fallback.

The replacement range includes the matched heading line and all following lines until the next heading whose level is ≤ the matched heading level. Nested headings within the matched section are replaced.

Zero matches: `section_selector_no_match`; no proposal created.
Multiple matches: `section_selector_ambiguous`; no proposal created.
Diagnostics should include `candidate_headings` (heading text, level, ordinal) when possible.

**Heading-safe replacement body normalization:**

Evaluated after `section_selector` resolves to exactly one section. The comparison uses the resolved section heading text and level (not the raw `section_selector` input values), so it applies regardless of whether `section_selector.level` was supplied.

When the first non-empty line of the replacement body is an ATX heading whose text equals the resolved selected section heading text and whose ATX level equals the resolved selected section heading level, that heading line is stripped before retained proposal creation.

- Applies to both direct `body` and `body_cache_id` replacement content.
- Only the first matching heading line is stripped.
- Body-internal headings after the first content line are preserved as section content.
- When the first non-empty line heading does not match the resolved section text or level, no stripping occurs.
- When stripping occurs, a `section_replacement_body_heading_stripped` warning diagnostic is returned with `stripped_heading` (stripped heading text) and `stripped_level` (ATX level as integer). This warning does not block proposal creation.
- Error-severity diagnostics continue to block proposal creation regardless of this normalization.

### Operations array

`operations` applies multiple operations atomically to the same record in one retained proposal. `update` and `operations` are mutually exclusive (supplying both: `invalid_request`, no proposal). Empty `operations` array: `invalid_request`. A single-element `operations` array has the same semantics as the corresponding single `update`.

Each element fields:

| field | required | type | meaning |
|---|---:|---|---|
| `type` | yes | string | Same values as `update.type` |
| `metadata` | conditional | object | Required for `metadata_block_replace` and `metadata_fields_replace` |
| `section_selector` | conditional | object | Required for `named_section_replace` |
| `body` | conditional | string | Replacement body for `named_section_replace`; mutually exclusive with `body_cache_id` |
| `body_cache_id` | conditional | string | Cached replacement body; mutually exclusive with `body` |

Per-operation body source rules match the single `update` operation rules. `body`/`body_cache_id` must be omitted for `metadata_block_replace`/`metadata_fields_replace`. Exactly one of `body`/`body_cache_id` must be present for `named_section_replace`. Supplying both in one element: `invalid_body_source`, no proposal. Top-level `body` and `body_cache_id` are only for single `update` requests; supplying either alongside `operations` returns `invalid_body_source`.

For each `named_section_replace` element that supplies inline `body`, if preparation fails after the body is received, the response SHOULD include a retryable `body_cache` entry for that body.

**Operation ordering** (deterministic, regardless of array position):

1. `metadata_block_replace` and `metadata_fields_replace` operations, in array order.
2. `named_section_replace` operations, in array order.

**Conflict detection** (runs before any operation is applied; conflicting array: no proposal):

- Two or more operations targeting the same metadata field: `conflicting_operations`. For `metadata_fields_replace`, conflict is determined by comparing `metadata` keys. A `metadata_block_replace` conflicts with any other metadata operation in the same array.
- Two or more `named_section_replace` operations in one `operations` array: `multiple_section_replace_not_supported`. MVP constraint regardless of whether they target the same or different sections.

**Validation** is performed against the final record state after all operations are applied, not an intermediate state.

**No-op detection:** when the combined result is byte-equivalent to the current file, no proposal is created. Response has `proposal_created: false` and a `no_op_update` info diagnostic, identical to single-operation no-op.

**Response shape:** `operations` proposals use the same retained proposal response as single-operation proposals. When `diff_mode` is `patch`, `diff.text` covers all changes from all operations in a single unified diff for the target file. When `diff_mode` is `summary` or `patch`, `diff.files` has one entry for the target record with `change: modify`.

## Response

When an update request produces changed content:

```json
{
  "proposal_id": "pw_opaque",
  "state": "proposed",
  "operation": "update",
  "target_kind": "task",
  "target": {
    "requested_id": "V01-TASK-MCP-008-04",
    "resolved_id": "V01-TASK-MCP-008-04",
    "kind": "task",
    "domain": "MCP",
    "path": "v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md"
  },
  "expires_at": "2026-06-05T00:00:00Z",
  "retention_days": 3,
  "diff": {
    "format": "unified",
    "files": [
      {
        "path": "v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md",
        "change": "modify",
        "record_id": "V01-TASK-MCP-008-04",
        "record_kind": "task"
      }
    ],
    "text": "diff --git a/v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md b/v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md\nindex oldhash..newhash 100644\n--- a/v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md\n+++ b/v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md\n@@ -1,7 +1,7 @@\n # V01-TASK-MCP-008-04: MCP tools spec reflection\n \n - **id**: V01-TASK-MCP-008-04\n-- **status**: todo\n+- **status**: done\n - **date**: 2026-06-01\n"
  },
  "validation": {
    "ok": true,
    "diagnostics": []
  },
  "diagnostics": [],
  "note": "No repository files have been written. Call accept_proposed_write with this proposal_id to apply the diff."
}
```

`target.resolved_id` is the existing record ID for update proposals. `diff.files[].change` is `modify`.

When `diff_mode` is `patch`, `diff.text` MUST compare the current persisted file content with the proposed persisted content after operation semantics. It MUST NOT render the entire target record as newly added content unless the file is actually new. For modify proposals, `diff.text` MUST use git-style unified diff format including `diff --git a/<path> b/<path>`, hash line, `--- a/<path>`, `+++ b/<path>`, `@@` hunk headers, and changed lines plus bounded context.

No-op update response:

```json
{
  "proposal_created": false,
  "operation": "update",
  "target_kind": "task",
  "target": {
    "requested_id": "V01-TASK-MCP-008-04",
    "resolved_id": "V01-TASK-MCP-008-04",
    "kind": "task",
    "domain": "MCP",
    "path": "v01/records/tasks/mcp/V01-TASK-MCP-008-04-mcp-tools-spec-reflection.md"
  },
  "validation": {
    "ok": true,
    "diagnostics": []
  },
  "diagnostics": [
    {
      "category": "no_op_update",
      "severity": "info",
      "message": "update produced no persisted content changes",
      "subject": {
        "type": "record",
        "ref": "V01-TASK-MCP-008-04",
        "record_kind": "task"
      }
    }
  ]
}
```

No-op response fields:

| field | required | meaning |
|---|---:|---|
| `proposal_created` | yes | `false` |
| `operation` | yes | `update` |
| `target_kind` | yes | Resolved target record kind |
| `target` | yes | Requested / resolved target identity object |
| `validation` | yes | Proposal-local validation result; `ok: true` when no error diagnostics |
| `diagnostics` | yes | Includes `no_op_update` info diagnostic |
| `diff` | no | Omitted or `null` (no retained proposal) |
| `proposal_id` | no | Omitted (no retained proposal) |

`no_op_update` MUST have severity `info`. It MUST include `message` and a structured `subject` identifying the target record. It MUST NOT use scalar diagnostic `record_id` or standalone diagnostic `path`. The enclosing response `target` continues to expose the resolved authoring target under the authoring transaction contract. A T04-defined `location` may be included when that path-exposure policy allows it. `no_op_update` MUST NOT be returned as a tool execution error and MUST NOT set `validation.ok` to `false`.

No-op detection is evaluated after all operation semantics and normalization (including metadata field preservation, metadata validation, section selector resolution, and heading-safe replacement body normalization).

## Errors

| code | condition |
|---|---|
| `invalid_request` | Both `update` and `operations` supplied, empty `operations` array, `new` placeholder in `id`, or other invalid request |
| `invalid_body_source` | Both `body` and `body_cache_id` supplied (top-level or per-operation), or rule violation |
| `body_cache_not_found` | `body_cache_id` does not exist |
| `body_cache_expired` | `body_cache_id` is past expiry |
| `section_selector_no_match` | Named section selector matched no sections |
| `section_selector_ambiguous` | Named section selector matched multiple sections |
| `proposal_preparation_failed` | Proposal preparation failed before persistence |

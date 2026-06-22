# Reference: Fields

- **id**: `spec:drmcp.design_records_mcp.schema.fields`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines the internal record model field set (common and kind-specific), title extraction rules, and field validation contracts for all record kinds.

## Current contract

The internal record model absorbs differences between metadata sources and normalizes them into common fields plus kind-specific detail objects. Public responses return only this shape; no compatibility contract serving both the old flat metadata fields and the new kind-specific detail object simultaneously is provided. After the spec update, the parser / index / list / get / validate and their tests migrate to this response shape as a single cut.

### Common fields

| field | required | type | meaning |
|---|---:|---|---|
| `id` | yes | string | Record ID. For decision and investigation: derived from H1. For spec: read from `design_record.id`. |
| `kind` | yes | string | Record kind. |
| `title` | yes | string | Human-readable title from Markdown H1. |
| `status` | yes | string | Kind-specific status value. |
| `path` | yes | string | Markdown file path from repository root. |

### Kind-specific detail objects

| detail object | fields |
|---|---|
| `decision` | `depends_on`, `supersedes`, `migrated_to_spec` |
| `spec` | `depends_on` |
| `investigation` | `trigger`, `scope`, `non_scope`, `source_refs`, `follow_up_candidates`, optional `supersedes`, `related_*`, `follow_up_results` |
| `requirement` | `source_refs`, `work_items`, optional `subdomain` |
| `work_item` | `source_requirement`, `impact_refs`, `tasks`, optional `subdomain` |
| `task` | `work_item`, `source_requirement`, `estimate`, `depends_on`, `outputs`, optional `subdomain` |

`headings` and requested raw `body` can be added to the common response as `get_record` retrieval content; `get_records` items with `retrieval_status: "found"` reuse the same record representation.

ADR bullet metadata `date` may be parsed but is not included as a record response field. Investigation `date` is likewise parsed but not included in the common response field.

MVP does not include the following metadata fields: `topics`, `affects`, `refines`, `conflicts_with`, `owns`, `migration.state`, or per-section origin information.

> Source: V01-ADR-076 §front matter 方針, V01-ADR-077 §list_records の責務

### `id`

`id` is the stable record identifier. Tools expose the public ID (namespace_prefix-prefixed form).

| kind | public ID example | notes |
|---|---|---|
| `decision` | `V01-ADR-076` | ADR number, 3-digit zero-padded |
| `spec` | `V01-SPEC-design-records-mcp-schema` | stable spec ID |
| `investigation` | `V01-INV-MCP-001` | domain-scoped ID per V01-ADR-086 |
| `requirement` | `V01-REQ-MCP-003` | domain-scoped ID |
| `work_item` | `V01-WORK-MCP-003` | domain-scoped ID |
| `task` | `V01-TASK-MCP-003-01` | includes parent work item sequence and task sequence |

For `decision` records: the canonical ID is derived from H1 by stripping namespace_prefix, building bare `ADR-NNN`, and re-attaching namespace_prefix. The filename ID prefix is used only for consistency verification. When H1 is invalid, `invalid_h1_title` is issued and the filename-derived ID is not adopted as the canonical ID.

For requirement / work item / task: the canonical ID is taken from both the metadata `id` field and the H1 ID; all three (metadata `id`, H1 ID, and filename prefix) must agree. When workflow ID syntax does not match the bare ID grammar, or metadata / H1 / filename prefix disagree, `invalid_workflow_id` or `filename_id_mismatch` is issued.

Example (MVP `v01/records`, ADR):

```text
H1: # V01-ADR-076: Design Records MCP
id: V01-ADR-076
path: v01/records/adr/V01-ADR-076-design-records-mcp.md
```

Mismatch produces `filename_id_mismatch`.

> Source: V01-ADR-077 §validate_records の責務, V01-ADR-097, V01-ADR-099

### `kind`

Record kinds currently indexed, queried, and validated:

| kind | meaning |
|---|---|
| `decision` | ADR |
| `spec` | Spec |
| `investigation` | Investigation artifact |
| `requirement` | Requirement artifact |
| `work_item` | Work item artifact |
| `task` | New-format short-term task artifact |

This is not a closed enumeration. Other kinds may be added by subsequent decisions. MVP excludes legacy M-series task records, UC docs, and impl notes from record kind indexing.

> Source: V01-ADR-076 §MVP対象, V01-ADR-087 §1, V01-ADR-092 §1

### `status`

Allowed status values by kind:

| kind | allowed status |
|---|---|
| `decision` | `proposed` / `accepted` / `superseded` |
| `spec` | `confirmed` / `draft` / `wip` |
| `investigation` | `investigating` / `concluded` / `superseded` |
| `requirement` | `captured` / `decision_needed` / `accepted` / `deferred` / `rejected` |
| `work_item` | `not_started` / `in_progress` / `blocked` / `done` |
| `task` | `not_started` / `in_progress` / `blocked` / `done` |

Status sources: `decision` — ADR bullet metadata; `spec` — top-level YAML front matter `status` (canonical); `investigation` / `requirement` / `work_item` / `task` — their respective bullet metadata.

For `spec`: if `design_record.status` exists it must match the top-level `status`; mismatch produces `spec_status_mismatch`.

Status values not allowed for the record kind produce `invalid_status_for_kind`.

> Source: V01-ADR-076 §front matter 方針, V01-ADR-077 §validate_records の責務, V01-ADR-086 §5, V01-ADR-091, V01-ADR-092, V01-ADR-094

### `depends_on`

The field name `depends_on` has different semantics by kind.

**Decision / spec dependency:**

`decision` record `depends_on` is read from ADR bullet metadata as a list of record IDs this decision depends on.

`spec` record `depends_on` is read from `design_record.depends_on` as a list of record IDs this spec depends on. The spec top-level front matter `depends_on` is a doc-policy origin path list and is not treated as record dependency.

Canonical public ID-as-ref target forms for decision / spec dependency: namespace_prefix-prefixed `ADR-*` / `SPEC-*` / `INV-*` kinds (MVP: `V01-ADR-*` / `V01-SPEC-*` / `V01-INV-*`). ADR and spec records may validly depend on investigation records.

Non-existent ID reference: `missing_depends_on_target`. MVP does not check source/target status combinations.

**Task dependency relation:**

`task` record `depends_on` is a workflow relation field read from task bullet metadata, pointing to task artifacts this task's execution depends on. Canonical target form: public ID namespace_prefix-prefixed `TASK-*` kind (MVP: `V01-TASK-*`).

`task.depends_on` is subject to `unresolved_workflow_relation` / `invalid_workflow_relation_target` rather than design dependency rules. MVP verifies reference target existence but does not enforce same-work-item constraints, cycle detection, or execution order projection.

> Source: V01-ADR-077 §validate_records の責務, V01-ADR-091 §3・§6, V01-ADR-092 §4・§7

### `supersedes`

`supersedes` is the list of record IDs this record replaces.

`decision` record `supersedes` is read from ADR bullet metadata. `spec` record `supersedes` normalizes to empty list in MVP; `design_record.supersedes` values are ignored in MVP without a diagnostic.

MVP verifies only that referenced IDs exist. Target status and reverse-reference consistency are not checked.

Non-existent ID: `missing_supersedes_target`.

> Source: V01-ADR-077 §validate_records の責務

### `migrated_to_spec`

`migrated_to_spec` is ADR metadata indicating that the ADR's spec content has been migrated to a spec record.

`decision` record `migrated_to_spec` is read from ADR bullet metadata. `spec` record `migrated_to_spec` normalizes to `null` in MVP; `design_record.migrated_to_spec` values are ignored without a diagnostic.

Empty or whitespace-only values normalize to `null`. When non-empty, only `YYYY-MM-DD` format is valid in MVP. Path, record ID, and free-text migration-target representations are outside MVP scope.

MVP does not introduce a normalized migration state vocabulary (e.g. `migration.state`). Invalid format produces `invalid_migrated_to_spec`.

> Source: V01-ADR-076 §front matter 方針, V01-ADR-077 §validate_records の責務

### Title extraction

MVP extracts `title` from the Markdown H1. H1 IDs use the public ID form (namespace_prefix-prefixed). The parser strips namespace_prefix and validates against the bare ID grammar, then returns the prefixed canonical ID.

**Decision (ADR) H1 processing:**

1. Strip namespace_prefix from the raw H1 text.
2. Match against bare form: `^#\s+ADR-(?P<num>\d{3}):\s+(?P<title>.+?)\s*$`
3. `num` must be 3-digit zero-padded.
4. Separator is ASCII colon `:`.
5. At least one whitespace character required after the colon.
6. `title` must be non-empty after trim.
7. Canonical ID: `<namespace_prefix>ADR-<num>` (MVP example: `V01-ADR-076`).
8. Filename public ID prefix comparison is string equality.

Example (MVP, `namespace_prefix = V01-`):

```markdown
# V01-ADR-076: Design Records MCP
```

When H1 is absent or does not match the expected format: `invalid_h1_title`.

**Spec H1:** format `# <title>`. No leading number or ID required. Title is extracted by removing the leading `#` and adjacent whitespace, then trimming.

**Investigation H1:** strip namespace_prefix, then match bare form:

```
^#\s+(?P<id>INV-[A-Z0-9-]+-\d{3}):\s+(?P<title>.+?)\s*$
```

Canonical ID: `<namespace_prefix>` + H1 bare `id` (MVP example: `V01-INV-MCP-001`). Filename format: `<namespace_prefix>INV-<DOMAIN>-NNN-<slug>.md`. Filename and H1 canonical ID must agree; mismatch is a diagnostic target.

**Workflow artifact H1:** strip namespace_prefix, then match one of these bare forms:

```
^#\s+(?P<id>REQ-[A-Z0-9]+(?:-[A-Z0-9]+)*-\d{3}):\s+(?P<title>.+?)\s*$
^#\s+(?P<id>WORK-[A-Z0-9]+(?:-[A-Z0-9]+)*-\d{3}):\s+(?P<title>.+?)\s*$
^#\s+(?P<id>TASK-[A-Z0-9]+(?:-[A-Z0-9]+)*-\d{3}-\d{2}):\s+(?P<title>.+?)\s*$
```

Canonical ID: `<namespace_prefix>` + H1 bare `id`. Metadata `id` field and H1-derived canonical ID must agree. When H1 or metadata `id` does not follow the grammar: `invalid_workflow_id`. When these or the filename ID prefix disagree: `filename_id_mismatch`.

Title is not inferred from filename in MVP.

> Source: V01-ADR-077 §list_records の責務, V01-ADR-077 §理由, V01-ADR-092 §3, V01-ADR-099

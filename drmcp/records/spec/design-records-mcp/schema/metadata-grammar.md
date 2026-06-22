# Reference: Metadata grammar

- **id**: `spec:drmcp.design_records_mcp.schema.metadata_grammar`
- **status**: draft
- **date**: 2026-06-17
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines the bullet metadata grammar for ADR, investigation, and workflow artifact record kinds: how metadata blocks are delimited, which fields are recognized, and how values are parsed and normalized.

## Current contract

### ADR bullet metadata grammar

ADR bullet metadata is read from the metadata block that starts immediately after H1. In MVP, the metadata block spans from the line after H1 up to (but not including) the first H2 line or blockquote line. H2 lines begin with `## ` and blockquote lines begin with `>`. Blank lines within the metadata block are permitted.

ADR recognized metadata keys:

| key | list value | normalization |
|---|---|---|
| `status` | no | trimmed string |
| `date` | no | trimmed; parsed but not included in record response field |
| `depends_on` | yes (comma-separated) | split and trimmed; empty → empty list |
| `supersedes` | yes (comma-separated) | split and trimmed; empty → empty list |
| `migrated_to_spec` | no | trimmed; empty/whitespace → `null` |

Example:

```markdown
# V01-ADR-076: Design Records MCP

- **status**: accepted
- **date**: 2026-05-11
- **depends_on**: V01-ADR-050, V01-ADR-068
- **supersedes**:
- **migrated_to_spec**:
```

Recognized metadata line format:

```text
- **<key>**: <value>
```

Constraints:

- Bold marker `**` is required.
- Keys are case-sensitive.
- Recognized keys: `status` / `date` / `depends_on` / `supersedes` / `migrated_to_spec` only. Unrecognized keys are ignored in MVP.
- `value` is whitespace-trimmed (leading and trailing).
- Empty or whitespace-only `value` is treated as unspecified for that metadata field.

List values use comma separation. `depends_on` / `supersedes` are split on commas when non-empty, with each value trimmed. Empty `depends_on` / `supersedes` normalize to an empty list. Empty `migrated_to_spec` normalizes to `null`.

`date` may be parsed as ADR metadata but is not included in the MVP record field response.

> Source: V01-ADR-086 §4–§5

### Investigation bullet metadata grammar

The investigation metadata block follows the same delimitation rule as ADR: H1-adjacent, ending before the first H2 or blockquote line.

Required fields: `status` / `date` / `trigger` / `scope` / `non_scope` / `source_refs` / `follow_up_candidates`.

Optional fields (read only when present): `supersedes` / `related_requirements` / `related_work_items` / `related_adrs` / `related_specs` / `related_internal_design` / `related_coverage` / `follow_up_results`.

Example:

```markdown
# V01-INV-MCP-001: Design Records MCP investigation support

- **status**: concluded
- **date**: 2026-05-23
- **trigger**: V01-ADR-087
- **scope**: investigation MCP integration
- **non_scope**: writer tools
- **source_refs**:
  - V01-ADR-087
- **follow_up_candidates**:
  - なし
```

Reference field rules:

- `source_refs` and `follow_up_results` values must be canonical references: public ID-as-ref (`V01-ADR-*` / `V01-SPEC-*` / `V01-INV-*` etc.) or active `spec:` semantic ref. Unresolved supported refs are an error. Physical paths are read as compatibility input but are not canonical form and produce an error diagnostic.
- `internal-design:` / `coverage:` / `COV-*` are not required as MVP canonical references / resolver inputs per V01-ADR-088.
- `follow_up_candidates`: canonical form is the same as above. Unresolved refs indicate a planned artifact not yet created, returned as an `info` diagnostic (not an error). Physical path candidates produce a `noncanonical_follow_up_candidate` info diagnostic.
- `trigger` / `related_*` resolve and validation rules are not finalized in this version.
- Workflow artifact public IDs usable as canonical references in validated investigation reference fields: `V01-REQ-<DOMAIN>-NNN` / `V01-WORK-<DOMAIN>-NNN` etc. (full form with namespace_prefix). `TASK-*` types are supported for workflow artifact relations and direct resolver input only — they are not supported as canonical reference form in investigation metadata fields. `TASK-*` in `source_refs` / `follow_up_results` produces an `unsupported_reference` error; in `follow_up_candidates` produces `unsupported_reference` info.

> Source: V01-ADR-086 §4–§7, V01-ADR-087 §5–§8, V01-ADR-091 §6, V01-ADR-092 §3–§6

### Workflow artifact bullet metadata grammar

Requirement / work item / task metadata blocks follow the same delimitation rule: H1-adjacent, ending before the first H2 or blockquote line. Recognized metadata lines are `- **<key>**: <value>` and indented list items directly below.

Bare ID grammar (internal validation form after namespace_prefix stripping). Public IDs carry the namespace_prefix (e.g. `V01-WORK-DRMCP-001`).

```text
REQ-<DOMAIN>-NNN
WORK-<DOMAIN>-NNN
TASK-<DOMAIN>-<WORK-SEQUENCE>-<TASK-SEQUENCE>
```

- `<DOMAIN>`: uppercase ASCII letters / digits / hyphens; must not begin or end with a hyphen.
- `NNN` and `<WORK-SEQUENCE>`: 3-digit zero-padded decimal sequence.
- `<TASK-SEQUENCE>`: 2-digit zero-padded decimal sequence.
- Public ID (`metadata id`, H1 ID, and filename prefix) must all agree.
- Workflow relations are read only from ID-as-refs declared in metadata fields — not inferred from task ID strings or paths.

**Requirement** recognized fields:

- required: `id`, `status`, `date`, `source_refs`, `work_items`
- optional: `subdomain`

**Work item** recognized fields:

- required: `id`, `status`, `date`, `source_requirement`, `impact_refs`, `tasks`
- optional: `subdomain`

**Task** recognized fields:

- required: `id`, `status`, `date`, `work_item`, `source_requirement`, `estimate`, `depends_on`, `outputs`
- optional: `subdomain`

`subdomain` is an optional string field representing a concept-area grouping within the domain namespace. It is not part of the artifact ID. Valid values are derived dynamically from existing records in the same domain (no pre-defined catalog). For the subdomain model, see the namespace-model spec.

**Presence validation rules:**

- Required scalar fields must exist and be non-empty.
- Required list fields must exist.
- Empty list for a required list field is valid unless the artifact-specific rule requires non-empty.
- Empty items within a required list field are a validation error (`empty_required_metadata`).
- `date` is a required scalar validated against strict `YYYY-MM-DD` format.

`task.depends_on`: when the field exists and both the value and direct child list items are empty, it normalizes to empty list `[]`. No workflow relation diagnostic is generated in this case.

`source_refs` / `impact_refs` / `outputs`: workflow-external reference rules follow the existing canonical reference policy. These fields are required list fields subject to presence validation.

`work_items` / `source_requirement` / `tasks` / `work_item` / `depends_on`: workflow relation fields subject to relation integrity validation.

> Source: V01-ADR-086 §4–§7, V01-ADR-091 §6, V01-ADR-092 §3–§6

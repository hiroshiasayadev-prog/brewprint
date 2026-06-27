# Reference: Authoring transaction schema

- **id**: `spec:drmcp.design_records_mcp.schema.authoring_transaction_schema`
- **status**: draft
- **date**: 2026-06-28
- **parent**: `spec:drmcp.design_records_mcp.schema.overview`

## What this is

Defines the authoring transaction schema concepts: authoring target identity, proposal model, body cache model, metadata block replacement target, and section selector model.

## Current contract

Authoring transactions are not part of the Design Records record model. Proposals and body caches are retained operational objects; they are not included as targets of `list_records` / `get_record` / `get_records` / `resolve_reference`.

### Authoring target identity

Authoring target identity uses artifact identity as its primary key. Public requests do not accept a physical path as the primary input.

| field | meaning |
|---|---|
| `kind` | Target record kind |
| `requested_id` | Caller-supplied input ID. For create, may include a `new` placeholder. |
| `resolved_id` | Final ID resolved by MCP from the record index |
| `domain` | Workflow artifact domain. Used for requirement / work item / task create. |
| `parent_id` | Parent record ID used for parent-aware ID resolution. Required for task create. |
| `path` | Resolved repository-relative path. Transparency output only. |

`path` is an output that explains relocation or slug generation results; it is not the canonical authoring target identity for requests.

For create operations, the canonical target ID input is the top-level request `id`. `fields.id` is not a primary authoring target identity field and is not required. When a create request supplies `fields.id`, it is a duplicate consistency input only — it must exactly match the top-level ID after canonical ID normalization, and must be omitted when the top-level ID uses a `new` placeholder. Mismatch or placeholder-time `fields.id` is an invalid request, not a record validation diagnostic.

For domain-scoped workflow artifact creates, `domain` comparison against the ID domain segment is case-insensitive. Canonical record IDs keep the uppercase domain segment. Repository-relative paths use the lowercase normalized domain directory.

Example: `domain: "mcp"` + `id: "V01-REQ-MCP-011"` → canonical domain `MCP`, path domain `mcp`.

> Source: V01-REQ-MCP-011, V01-TASK-MCP-011-01

### Proposal model

A proposal is the retained representation of a write candidate. Proposal creation does not write repository files.

| field | meaning |
|---|---|
| `proposal_id` | Opaque lookup key |
| `state` | `proposed` / `accepted` / `discarded` |
| `operation` | `create` / `update` |
| `target_kind` | Resolved target kind |
| `target` | Authoring target identity |
| `base_state` | Target file / index state needed for accept-time staleness detection |
| `diff` | Previewable diff |
| `validation` | Proposal-time validation result |
| `required_follow_up_updates` | Required follow-up list that must be satisfied before acceptance |
| `expires_at` | Proposal expiry timestamp |
| `retention_days` | `3` |

`base_state` concrete shape (hash / timestamp / index snapshot) is an implementation detail. The public contract is that accept detects stale target / changed target / ID collision before writing, and returns `written: false` with diagnostics when detected.

Proposal retention is 3 days. Expired proposals are not valid authoring targets.

Proposal `validation` is scoped to the proposal-local affected record set in the candidate repository state. The affected record set is the proposed target record plus any related records whose files are actually modified by the same proposal (e.g. required reciprocal workflow metadata updates). Unrelated repository diagnostics are repository health information and must not be included in or change proposal-local `validation.ok`.

Proposal-local diagnostics must use the shared repository-validation categories and structured associations from `spec:drmcp.design_records_mcp.schema.diagnostics`. They must be reproducible by the same `validate_records` rules against the same affected record set in the same candidate or accepted/materialized state.

> Source: V01-REQ-MCP-012, V01-TASK-MCP-011-01

### Body cache model

Body cache is an operational object that temporarily holds a large Markdown body for proposal retries. It is not a design record and is not addressable by the resolver.

| field | meaning |
|---|---|
| `body_cache_id` | Opaque lookup key |
| `expires_at` | Cache expiry timestamp |
| `retention_days` | `3` |

Operations requiring large Markdown body input accept exactly one of `body` or `body_cache_id`, unless the tool-specific contract explicitly permits omission. Supplying both `body` and `body_cache_id` is invalid and must not create a proposal or body cache. Operations not requiring body input may omit both.

Unknown or expired body cache IDs must produce diagnostics and must not create proposals. Body cache entries remain reusable within the 3-day retention period, including after they have been used to create a proposal.

For `propose_record_create`:

- Structured `fields` are required.
- `body` and `body_cache_id` are section-only content sources and may be combined with `fields`, but not with each other.
- A create request supplying `body` without `fields` is invalid; when the submitted body is a string, the failed response should include a new body cache entry so the caller can retry with `fields + body_cache_id`.
- A create request supplying only `body_cache_id` without `fields` is invalid and cannot create a new body cache because no submitted body is present.

### Metadata block replacement target

Metadata block replacement targets the kind-specific metadata block.

| kind | metadata block |
|---|---|
| `spec` | Recognized spec metadata fields inside YAML front matter |
| `decision` | H1-following ADR bullet metadata block |
| `requirement` | H1-following requirement bullet metadata block |
| `work_item` | H1-following work item bullet metadata block |
| `task` | H1-following task bullet metadata block |

For `spec`, metadata replacement is scoped to recognized fields only. Unknown or auxiliary YAML front matter fields must be preserved. Recognized spec metadata fields: `scope`, top-level `status`, and `design_record.id` / `design_record.kind` / `design_record.status` / `design_record.depends_on`.

Required recognized fields are validated by the same field vocabulary and PRODUCT authority used for repository validation. Missing required fields produce `missing_required_field`. A present empty value produces `empty_required_field` only when the applicable authority requires non-empty content. Invalid present values produce `invalid_field_value`. The diagnostic uses structured `subject`, `field`, `value`, and applicable `target` or `location` associations from `spec:drmcp.design_records_mcp.schema.diagnostics`; authoring does not introduce separate kind-specific validation categories.

### Section selector model

Named section replacement uses an ATX heading selector that resolves within one Markdown record body.

MVP selector fields:

| field | meaning |
|---|---|
| `heading` | Exact heading text after ATX marker removal and whitespace trim |
| `match` | `exact` only |
| `level` | Optional ATX heading level constraint |

Matching is case-sensitive. No Unicode normalization, punctuation folding, prefix matching, or contains matching in MVP.

- Zero sections matched: `section_selector_no_match`.
- Multiple sections matched: `section_selector_ambiguous`.

Neither case may create a proposal or write files. Diagnostics should include candidate headings when possible.

Section selector resolution uses the same Markdown heading source rules as the `headings` field in the record model. YAML front matter content and fenced code block content are not heading sources. Setext headings are not section sources in MVP.

# Reference: Agent authoring policy

- **id**: `spec:product.concepts.authoring_standards.agent_authoring_policy`
- **status**: draft
- **date**: 2026-06-15
- **parent**: `spec:product.concepts.authoring_standards`

## What this is

Behavioral rules for AI agents working with brewprint design records. Governs record retrieval, authoring transaction preference, prose compliance, and namespace usage. DRMCP-dependent rules are deferred pending DRMCP operationalization.

## Non-goals

- Per-agent configuration (e.g., CLAUDE.md, AGENTS.md). Those files reference this spec; they do not replicate it.
- Tool API contracts — those belong to `spec:drmcp.design_records_mcp.tools`.
- Prose style rules — those belong to `spec:product.concepts.authoring_standards.writing_standard`.

## Rules

### Prose compliance

| Rule | Level |
|---|---|
| Follow `spec:product.concepts.authoring_standards.writing_standard` for all design record prose. | MUST |

### Namespace usage

| Rule | Level |
|---|---|
| When creating records, include the namespace prefix in the ID (e.g., `PRODUCT-REQ-new`, `DRMCP-WORK-new`). Do not use prefix-less IDs. | MUST |
| Create new REQ / WORK / TASK / ADR / SPEC under an active namespace (`product`, `drmcp`, `bpdsl`). `v01/records/` is a read-only snapshot. | MUST |

### DRMCP retrieval (TBD)

Rule set deferred. To cover: retrieval-first rule, filesystem fallback conditions, and tool entry points (`list_records`, `get_record`, `get_records`, `resolve_reference`, `validate_records`).

**Current operating mode**: DRMCP is non-operational. Use filesystem operations for all record retrieval and authoring.

### Authoring transaction preference (TBD)

Rule set deferred. To cover: propose → review diff/diagnostics → accept flow, body vs `body_cache_id` contract, fallback conditions, and `*-new` placeholder usage.

## Related specs

| ref | relation |
|---|---|
| `spec:product.concepts.authoring_standards` | Parent Index. |
| `spec:product.concepts.authoring_standards.writing_standard` | Prose style rules for design record authoring. |
| `spec:drmcp.design_records_mcp.tools` | Authoritative DRMCP tool API contract. |

# Docs Review Instructions

## Scope

This directory contains project documentation, ADRs, specs, task files, UC docs, and implementation notes.

Do not treat all docs as equally authoritative.

## Authority model

- specs describe the current intended behavior
- ADRs record decisions and rationale at the time they were written
- task files track work status and may become stale
- handoff notes are temporary and may be incomplete
- implementation notes describe what happened, not necessarily what should remain true

## ADR rules

ADR files use H1 plus bullet metadata, not YAML front matter.

Example:

- **status**: accepted
- **date**: YYYY-MM-DD
- **depends_on**:
- **supersedes**:
- **migrated_to_spec**:

Do not add YAML front matter to ADRs unless explicitly instructed.

## Spec rules

Spec files use YAML front matter.

When editing specs, keep:

- scope
- status
- last_updated
- summary
- depends_on

If a spec has `design_record`, keep it consistent with the Design Records MCP spec.

## Review focus

When reviewing docs, check:

- whether ADR and spec responsibilities are mixed
- whether current spec contradicts accepted ADRs
- whether ADR wording is historical but being treated as current spec
- whether MVP-out items accidentally become required
- whether metadata source is unambiguous
- whether implementation can follow the spec without guessing

Do not approve text merely because it sounds plausible.
# V01-REQ-MCP-011: Authoring proposal input contract normalization

- **id**: V01-REQ-MCP-011
- **status**: accepted
- **date**: 2026-06-02
- **source_refs**:
  - V01-REQ-MCP-008
  - V01-ADR-093
  - SPEC-design-records-mcp-tools
  - SPEC-design-records-mcp-schema
- **work_items**:
  - V01-WORK-MCP-011

## Requirement

Design Records MCP authoring proposal inputs need a normalized and explicit contract for structured fields, full Markdown bodies, IDs, and domains.

The current authoring transaction surface exposes multiple input routes for creating records, but observed behavior is ambiguous or inconvenient in several places:

- when `body` and `fields` are both supplied, `body` appears to take precedence and structured `fields` are ignored, without an explicit contract-level rule;
- create operations require `fields.id` even when the top-level `id` parameter already supplies the record ID;
- `domain` is compared case-sensitively against the uppercase domain segment in IDs such as `V01-REQ-MCP-011`, even though repository paths use lowercase domain directories such as `docs/requirements/mcp/`.

These behaviors make authoring transaction tools difficult to use safely from AI assistants, because the caller must know undocumented precedence, duplicate ID placement, and case behavior.

## Evidence

During dogfooding of `propose_record_create` for a new MCP requirement, the following behavior was observed:

- omitting `fields.id` caused an `invalid_request` even though top-level `id` was present;
- supplying both `body` and `fields` caused the generated proposal to follow `body` rather than structured `fields`;
- using `domain: "mcp"` with `id: "V01-REQ-MCP-011"` failed with `domain "mcp" does not match ID domain "MCP"`.

The domain behavior is especially misleading because existing MCP requirement files live under the lowercase `mcp` directory, while their canonical IDs use uppercase `MCP`.

Close evidence on 2026-06-02: `V01-WORK-MCP-011` specified, implemented, regression-tested, and runtime-smoked the input normalization contract. `V01-TASK-MCP-011-04` confirmed fields-only create success, top-level ID use without `fields.id`, case-insensitive domain handling with uppercase canonical domains and lowercase path domains, and the expected `invalid_request` / `invalid_body_source` rejection behavior.

## Required Outcome

A follow-up work item should define and implement a clear input normalization contract for authoring proposal create operations.

The expected properties are:

- exactly one primary content source is selected for record creation;
- if both `body` and `fields` are supplied, the behavior is either rejected as `invalid_request` or explicitly specified in the public tool contract;
- top-level `id` is the canonical create target ID and does not need to be duplicated in `fields.id`, unless the public schema intentionally documents the duplication;
- if both top-level `id` and `fields.id` are supplied, they must match after the same normalization used by the tool contract;
- domain input is normalized consistently between canonical IDs and repository path domains;
- diagnostics explain which input field is invalid and how to repair it.

## Explicitly Excluded Scope

This requirement does not require:

- changing the canonical workflow artifact ID format;
- changing repository directory names;
- adding generic filesystem-path authoring inputs;
- changing record validation rules unrelated to authoring request normalization;
- implementing spec skeleton creation or `SPEC-new` support.

## Boundary

This requirement owns the authoring proposal input contract for create operations.

It does not own proposal validation scope, full repository validation behavior, existing repository metadata cleanup, or authoring transaction persistence semantics. Proposal validation scope isolation is captured separately by `V01-REQ-MCP-012`.

# REQ-MCP-010: Spec domain tree and placement discovery support

- **id**: REQ-MCP-010
- **status**: captured
- **date**: 2026-06-02
- **source_refs**:
  - REQ-MCP-008
  - WORK-MCP-008
  - ADR-093
  - SPEC-design-records-mcp-tools
- **work_items**:

## Requirement

Design Records MCP needs spec domain tree and placement discovery support before artifact-oriented spec creation can be safely added.

Spec records do not have a simple one-to-one ID-to-path derivation rule. For example, `SPEC-design-records-mcp-tools` is located at `docs/spec/design-records-mcp/tools.md`, and that placement cannot be derived safely from `SPEC-new` or `SPEC-<slug>` alone.

The MCP should provide a way to inspect the current spec domain tree and discover existing spec document placement before any future spec skeleton creation flow is introduced.

## Evidence

During WORK-MCP-008 spec reflection review, spec skeleton creation was found to require placement information that the current MCP public surface does not provide.

The authoring transaction MVP therefore excludes `SPEC-new` and spec skeleton creation. Existing spec records may still be updated by metadata block or named section replacement when the target record already exists and can be resolved by record ID.

Without a spec domain tree / placement discovery tool, any implementation of `SPEC-new` would either:

- invent a path derivation rule not represented in the public contract;
- require filesystem path input, which conflicts with pathless authoring; or
- risk placing new spec documents in an unintended domain.

## Required Outcome

A follow-up work item should define a Design Records MCP read/navigation capability for spec placement discovery.

Candidate capability:

- list spec domains;
- list spec documents under a domain;
- optionally include document IDs and paths;
- optionally filter by an exact spec domain;
- return repository-relative paths as transparent output, not as primary authoring input.

A possible shape is a `list_spec_tree` or broader `list_record_tree` tool that can return:

- domain name;
- domain path;
- child domains, if supported;
- existing spec document IDs;
- existing spec document titles;
- existing spec document paths.

The final tool name and schema should be decided by the follow-up work item / spec update.

## Explicitly Excluded Scope

This requirement does not require:

- spec skeleton creation;
- `SPEC-new` support;
- generic filesystem tree browsing;
- path-first authoring APIs;
- moving existing spec documents;
- refactoring the current spec directory layout;
- implementation of authoring writes.

## Boundary

This requirement captures the missing spec placement discovery capability that was intentionally excluded from REQ-MCP-008 MVP authoring transaction support.

It does not own the final tree tool schema, spec document placement policy, or future spec skeleton creation contract. Those should be handled by a follow-up work item and, if needed, a later ADR / spec update.

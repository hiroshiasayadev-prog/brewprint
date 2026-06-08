# V01-REQ-MCP-005: Project authoring guidance retrieval support

- **id**: V01-REQ-MCP-005
- **status**: accepted
- **date**: 2026-05-31
- **source_refs**:
  - docs/doc-policy.md
  - docs/adr-authoring-guide.md
  - docs/spec-authoring-guide.md
  - docs/requirements/README.md
  - docs/work-items/README.md
  - docs/tasks/README.md
  - docs/investigations/README.md
- **work_items**:
  - V01-WORK-MCP-005

## Requirement

Design Records MCP needs a project authoring guidance retrieval surface so AI assistants can discover and read the minimum necessary authoring guidance for the current task without treating every project-wide guide as mandatory startup context.

The project needs a paired tool flow:

- `list_authoring_guides` to discover available authoring guides and their intended use.
- `get_authoring_guidance` to retrieve the relevant Markdown guidance by guide ID.

## Motivation

Current authoring guides such as ADR authoring, spec authoring, requirements, work items, tasks, investigations, and doc policy are project-wide operational knowledge. They are useful across projects that adopt brewprint, but reading them all during startup creates excessive context cost and makes `docs/doc-policy.md` grow into a large operational manual.

The desired direction is to keep README / doc-policy as a thin entry point and move task-specific guidance loading behind Design Records MCP discovery and retrieval.

## Required Outcome

- Design Records MCP exposes a discovery tool for available authoring guides.
- Design Records MCP exposes a retrieval tool for a selected authoring guide.
- The returned guide metadata distinguishes guide identity, title, and abstract without exposing source path as public response contract.
- `docs/doc-policy.md` and top-level README-style entry points can reference the MCP guidance retrieval flow instead of duplicating detailed authoring rules.
- The design remains usable by projects that adopt brewprint, not only by the brewprint repository itself.

## Explicitly Excluded Scope

- Rewriting every authoring guide in this requirement.
- Removing existing Markdown authoring guides before MCP support exists.
- Making all project docs mandatory startup reads.
- Replacing Design Records MCP record retrieval tools such as `get_record` / `get_records`.
- Defining a generic non-authoring knowledge-base search tool.

## Boundary

This requirement captures the need for authoring guidance discovery and retrieval. It does not by itself decide the final response schema, tool names beyond the requested working names, guide ID vocabulary, or doc-policy migration details. Those belong to the linked work item and follow-up tasks / ADR / spec updates.

# DRMCP-ADR-MCP-011: Inward ownership and Guidance query aliases

- **status**: accepted
- **date**: 2026-07-04
- **depends_on**:
  - DRMCP-ADR-MCP-010
  - DRMCP-ADR-MCP-012
- **supersedes**:
  - DRMCP-ADR-MCP-008
- **migrated_to_spec**: null

## Context

ADR-008 assigned Guidance identity and projection behavior to a separate Guidance Domain and Guidance Source port.
The revised architecture treats portable standards as ordinary Current Records.

Guidance still needs distinct public operations and response contracts.
It does not need distinct storage, parsing, identity, logical-tree, or lifecycle ownership.

The architecture must preserve inward dependencies and public use-case isolation while reusing normal record-query capabilities.

## Decision

Dependencies point inward.

The MCP Inbound Adapter depends on Application Use Cases.
Application Use Cases depend on Record Domain / Logical Tree and application-owned source contracts.
Infrastructure I/O Adapters implement those source contracts.
Record Domain does not depend on Application, MCP, filesystem, or runtime configuration.
Composition / Lifecycle alone wires concrete components.

Application Use Cases contain:

| group | responsibility |
|---|---|
| Operation-specific use cases | Public request policy, sequencing, fixed scope, response projection, and result semantics. |
| Shared application orchestration | Reusable request-scoped Current Records loading, logical-state assembly, exact lookup, and query primitives. |

Public use cases do not call one another.
`list_authoring_guides` does not call `list_records`.
`get_authoring_guidance` does not call `get_records`.
They reuse shared internal orchestration and Record Domain behavior.

Guidance list fixes this scope:

- app namespace `design_records`;
- record kind `spec`;
- canonical child subtree `spec:design_records.authoring_standards.*`;
- exclusion of the root `spec:design_records.authoring_standards`.

Guidance detail accepts one exact canonical ref inside that child subtree.
It does not accept filename stems, basenames, physical paths, titles, fuzzy inputs, aliases, or inferred candidates.

Guidance projection ownership is split as follows:

| concern | owner |
|---|---|
| Canonical Spec identity, parsed H1, headings, sections, and complete body | Record Domain / Logical Tree. |
| Fixed app, kind, subtree, root exclusion, ordering, response shape, and operation errors | Guidance Application Use Case. |
| Concrete Current Records source access | Infrastructure I/O Adapters behind normal Current Records source contracts. |
| Construction and source wiring | Composition / Lifecycle. |

The public projection is:

| field | rule |
|---|---|
| `id` | Canonical package Spec ref. |
| `title` | First H1 text. |
| `abstract` | Body of the `## What this is` section. |
| `content` | Complete Markdown source verbatim. |

List ordering uses canonical-ref ASCII lexical order.
Normal Guidance responses hide physical paths.

PRODUCT remains the external semantic authority.
The portable package is a distribution of PRODUCT semantics.
Record Domain applies normal current Spec and validation rules without interpreting the package as a separate standard.

Failures retain three ownership classes:

| class | owner and handling |
|---|---|
| Request rejection | MCP owns protocol decoding. Application owns operation-specific invalid scope or input. |
| Expected semantic outcome | Record Domain produces modeled record states. Application projects zero match, unresolved, invalid, ambiguous, or guide-not-found behavior. |
| Execution failure | Infrastructure reports access failure. Application aborts when the required Current Records result is not complete and trustworthy. |

Downstream design may refine packages, interfaces, data types, query helpers, section extraction, adapter APIs, filesystem libraries, and error-code representation inside this ownership model.

Work returns to architecture when a change would add a top-level component, bypass inward dependencies, add a separate package state model, move Guidance policy into Domain or Infrastructure, or change trustworthy-result semantics.

## Rationale

Guidance tools expose a specialized public view, not a specialized record source.
Application policy owns the fixed query and response projection.

Record Domain already owns the parsed information required by Guidance.
Reusing shared orchestration avoids public use-case chaining and duplicate index behavior.

Canonical package refs preserve portable identity across hosts.
`## What this is` is the current Spec-owned summary section and therefore supplies the Guidance abstract.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Keep a Guidance Domain | No independent domain semantics remain after package Specs enter Current Records. |
| Keep a Guidance Source port | Normal Current Records source contracts already provide required source access. |
| Let Guidance call public list or get tools | Public use cases would become transitively coupled. |
| Use filename stems as guide IDs | Filename identity discards canonical package namespace and conflicts with ADR-001. |
| Use `## Abstract` | Current authoring-standard Specs use `## What this is`. |
| Put fixed Guidance filtering in Record Domain | App, subtree, ordering, and response policy belong to the use case. |

## Consequences

- ADR-008 becomes historical and superseded.
- Guidance Domain and Guidance Source disappear from current architecture.
- Guidance operations share normal Current Records query and retrieval primitives.
- Generic public record operations do not need Guidance-specific request options.
- Canonical package refs become public guide IDs.
- Exact error identifiers remain operation-contract details.
- New Guidance storage or lifecycle requirements must return to architecture.

Affected design areas:

- `spec:drmcp.application_architecture.dependency_and_responsibility`;
- `spec:drmcp.application_architecture.failure_and_evolution`;
- `spec:drmcp.design_records_mcp.schema.authoring_guidance_source`;
- `spec:drmcp.design_records_mcp.tools.list_authoring_guides`;
- `spec:drmcp.design_records_mcp.tools.get_authoring_guidance`.

## Evidence

- Source Requirements: `DRMCP-REQ-MCP-003` and `DRMCP-REQ-MCP-005`.
- Source Work Item: `DRMCP-WORK-MCP-016`.
- Review findings: F-BLK-01 and F-MAJ-01 in `DRMCP-TASK-MCP-016-09`.
- Revised decisions: D-021 and D-022 in `DRMCP-TASK-MCP-016-12`.
- ADR routing authority: B-05 in `DRMCP-TASK-MCP-016-13`.
- Component dependency: `DRMCP-ADR-MCP-010`.
- State dependency: `DRMCP-ADR-MCP-012`.

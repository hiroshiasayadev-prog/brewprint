# DRMCP-ADR-MCP-004: Internal state and operation contract separation

- **status**: accepted
- **date**: 2026-06-30
- **depends_on**:
  - DRMCP-ADR-MCP-003
- **supersedes**: []
- **migrated_to_spec**: 2026-06-30

## Context

The accepted read contracts distinguish valid records, invalid but addressable sources, duplicate conflicts, diagnostics, and operation-specific public responses.

One nullable `Record` type shared across parsing, indexing, validation, application, and MCP would make absence and invalidity ambiguous.
It would also let public tool response changes reshape core parsing and index state.

The rebuild needs a stable boundary between internal semantic state and each operation's public meaning.

## Decision

The runtime uses separate internal types for:

- raw and parsed source state;
- canonical index entries;
- duplicate identity conflict state;
- invalid but addressable index state;
- validation findings and operation diagnostics;
- operation-specific public projections.

No single nullable `Record` type spans every layer.
Missing parsed values remain missing.
Invalid parsed values remain available as invalid data.
Conflict state does not select a winner.

Each application use case owns one typed input and one typed output contract.
Application contracts do not import MCP SDK or protocol wrapper types.

Core services return domain values, lookup states, conflict state, and validation findings.
Core does not define complete tool responses.

The application use case projects core results into the accepted operation output.
Expected not-found, conflict, disabled-fallback, unresolved, and semantic-invalidity states remain typed result data when the public tool contract requires a normal response.

An application execution error is reserved for a failure that prevents construction of a trustworthy operation result.
Examples include an unbuildable mandatory snapshot or a required diagnostic location that cannot be represented.

The MCP adapter owns:

- protocol and SDK wrappers;
- request decoding;
- structural JSON validation;
- application invocation;
- response encoding.

The MCP adapter does not redesign operation fields, statuses, diagnostics, or error semantics.

## Rationale

Separate types preserve the meaning and lifecycle of each state.
An invalid but addressable source is not a partially successful valid record.
A duplicate conflict is not a record with an optional conflict flag.

Operation-owned contracts keep public response meaning close to sequencing and projection.
Core remains reusable and does not become coupled to MCP envelopes.

Distinguishing expected result states from execution errors preserves partial-success and diagnostic contracts without returning normal responses from untrustworthy state.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Share one nullable `Record` across all layers | Missing, invalid, conflicted, and absent states would be ambiguous. |
| Let core define complete public tool responses | Core would depend on operation and transport concerns. |
| Let the MCP adapter assemble semantic responses directly | Transport would become the owner of fields, statuses, and diagnostic meaning. |
| Treat every not-found or conflict as an execution error | Accepted normal result and partial-success behavior would be lost. |
| Treat every failure as normal result data | Untrustworthy snapshots or incomplete required diagnostics could be misreported as valid responses. |

## Consequences

- Parsers, indexes, validators, and use cases require explicit conversion boundaries.
- Public response changes do not automatically change raw or indexed source types.
- Application tests can verify operation projection without MCP protocol setup.
- MCP adapter tests focus on decode, structural validation, invocation, and encode.
- Result-data versus execution-error classification must remain explicit in each operation contract.
- Shared diagnostic representation remains a separate schema contract consumed by operation outputs.

Affected Specification targets:

- `spec:drmcp.design_records_mcp.schema.overview`;
- `spec:drmcp.design_records_mcp.schema.record_model`;
- `spec:drmcp.design_records_mcp.schema.diagnostics`;
- `spec:drmcp.design_records_mcp.tools.overview`;
- operation-specific tool contracts.

## Evidence

- Source Requirement: `DRMCP-REQ-MCP-001`.
- Source Work Item: `DRMCP-WORK-MCP-011`.
- Accepted decisions: D-003 and D-008 in `DRMCP-TASK-MCP-011-01`.
- Logical layer authority: `DRMCP-ADR-MCP-003`.

# Overview: Tools

- **id**: `spec:drmcp.design_records_mcp.tools.overview`
- **status**: draft
- **date**: 2026-06-30
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Provides the navigation-first catalog for Design Records MCP tools and points each operation to its owning contract.

## Current contract

### Contract hierarchy

| concern | authority |
|---|---|
| Public tool catalog and navigation | This overview. |
| Request, response, and operation error behavior | The individual tool contract. |
| Shared parsed fields and internal record state | The corresponding `schema.*` contract. |
| Warning and diagnostic entry representation | `spec:drmcp.design_records_mcp.schema.diagnostics`. Individual operation contracts own trigger conditions and response placement. |

This overview does not define a shared record response shape. Individual operations select the projection they return.

### Application and transport boundary

`list_records`, `get_records`, `resolve_reference`, and `validate_records` each invoke one dedicated application use case. No generic App Router is part of the accepted architecture for those four operations.

Each of those application use cases owns a typed input, typed output, operation sequencing, and projection into the accepted public meaning.

Core services return domain values, lookup states, conflicts, and validation findings. Core services do not construct complete tool responses.

The MCP adapter owns protocol schemas, decoding, structural JSON validation, invocation, and encoding. MCP SDK types do not enter application or core contracts.

Expected semantic states remain normal result data when an operation contract defines them. A failure that prevents a trustworthy result is an execution error.

`spec:drmcp.implementation` owns the layer, port, adapter, and package architecture for those four operations.

`list_authoring_guides` and `get_authoring_guidance` remain in the public catalog. Their application architecture remains with the authoring-guidance contracts.

Authoring-transaction tools remain in the public catalog. Their application, snapshot, and package architecture is deferred to `DRMCP-REQ-MCP-002`.

This scope clarification does not change any public request, response, status, diagnostic, warning, or error behavior.

### Public read and navigation tools

| tool | purpose | contract owner |
|---|---|---|
| `list_records` | Return a compact current-record listing for one active app, sequential kind, and domain. | `DRMCP-WORK-MCP-004` |
| `get_records` | Perform sole exact retrieval for one through 20 current or configured legacy refs. A single retrieval uses one `refs` element. | `DRMCP-WORK-MCP-004` |
| `resolve_reference` | Resolve references through the current-first and configured legacy-fallback contract. | `DRMCP-WORK-MCP-005` |
| `validate_records` | Validate current records and return machine-readable diagnostics. | `DRMCP-WORK-MCP-006` |
| `list_authoring_guides` | Return the authoring-guidance catalog defined by its operation contract. | Authoring-guidance contracts. |
| `get_authoring_guidance` | Retrieve authoring guidance defined by its operation contract. | Authoring-guidance contracts. |

`get_record` is not a public tool. The server tool catalog omits it, and no compatibility alias routes it to `get_records`.

`suggest_next_record` is removed. Sequential allocation belongs to namespace-aware authoring placeholders under the authoring contracts.

### Query and exact-retrieval summary

| operation | summary |
|---|---|
| `list_records` | Queries the active index only. Returns compact `ref`, `title`, `status`, and `date` entries plus `has_more` and operation warnings. Specs and legacy archive records are excluded. |
| `get_records` | Accepts exact `refs` and request-wide `include_body`. Returns successful records only, preserves first-occurrence order, and does not invoke the resolver. |

Normal list and exact-retrieval responses do not expose physical paths, source provenance, resolver traces, or internal index state.

The warning taxonomy, severity, shared fields, source-location representation, and exceptional path exposure are not defined here.

### Authoring transaction tools

Proposal creation does not modify repository files. Only `accept_proposed_write` may write repository files.

| tool | purpose |
|---|---|
| `propose_record_create` | Create an artifact-oriented write proposal. |
| `propose_record_update` | Create a metadata-block or named-section replacement proposal. |
| `get_proposed_write` | Retrieve a retained proposal by proposal ID. |
| `accept_proposed_write` | Accept a proposal and attempt repository writes. |
| `discard_proposed_write` | Discard a proposal and prevent later acceptance. |

Authoring request, response, diagnostic, and write-boundary behavior remains defined by the authoring operation contracts.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| `list_records` | Contract | `spec:drmcp.design_records_mcp.tools.list_records` | Compact active-index listing for one app, sequential kind, and domain. |
| `get_records` | Contract | `spec:drmcp.design_records_mcp.tools.get_records` | Sole exact retrieval for one through 20 current or configured legacy refs. |
| `list_authoring_guides` | Contract | `spec:drmcp.design_records_mcp.tools.list_authoring_guides` | Return the authoring-guidance catalog defined by the operation contract. |
| `get_authoring_guidance` | Contract | `spec:drmcp.design_records_mcp.tools.get_authoring_guidance` | Retrieve authoring guidance defined by the operation contract. |
| `resolve_reference` | Contract | `spec:drmcp.design_records_mcp.tools.resolve_reference` | Resolve a reference under the current-first and configured fallback contract. |
| `validate_records` | Contract | `spec:drmcp.design_records_mcp.tools.validate_records` | Validate current records and return diagnostic results. |
| Authoring transaction model | Reference | `spec:drmcp.design_records_mcp.tools.authoring_transaction_model` | Shared concepts for the 5 write tools: proposal lifecycle, body cache, diff_mode, affected record set. |
| `propose_record_create` | Contract | `spec:drmcp.design_records_mcp.tools.propose_record_create` | Create a retained proposal for a new record. |
| `propose_record_update` | Contract | `spec:drmcp.design_records_mcp.tools.propose_record_update` | Create a retained proposal for a metadata or section update. |
| `get_proposed_write` | Contract | `spec:drmcp.design_records_mcp.tools.get_proposed_write` | Retrieve a retained proposal by ID. |
| `accept_proposed_write` | Contract | `spec:drmcp.design_records_mcp.tools.accept_proposed_write` | Accept a proposal and write to repository files. |
| `discard_proposed_write` | Contract | `spec:drmcp.design_records_mcp.tools.discard_proposed_write` | Discard a proposal; prevents future acceptance. |

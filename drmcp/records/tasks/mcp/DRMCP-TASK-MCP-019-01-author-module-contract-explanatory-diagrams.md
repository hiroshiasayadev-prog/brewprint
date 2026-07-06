# DRMCP-TASK-MCP-019-01: Author module-contract explanatory diagrams

- **id**: DRMCP-TASK-MCP-019-01
- **status**: done
- **date**: 2026-07-06
- **work_item**: DRMCP-WORK-MCP-019
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**: []
- **outputs**:
  - DRMCP-TASK-MCP-019-01
  - `spec:drmcp.implementation.contracts`
  - `spec:drmcp.implementation.contracts.composition_lifecycle.contract_boundary`
  - `spec:drmcp.implementation.contracts.mcp_inbound_adapter.contract_boundary`
  - `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary`
  - `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary`
  - `spec:drmcp.implementation.contracts.infrastructure_io_adapters.contract_boundary`

## Goal

Author a small explanatory diagram set for the W018 module-contract Specification baseline.

The diagrams clarify ownership, collaboration, state assembly, request-state visibility, and source-boundary rules without changing the accepted contract.

## Work

- Read the active ChatGPT repository instruction source and design-record authoring standards.
- Read the W018 target Specification files and DRMCP-ADR-MCP-013.
- Select the minimum diagram set needed to clarify the accepted W018 baseline.
- Add Mermaid diagrams to the six target Specification files.
- Keep diagrams responsibility-oriented and implementation-neutral.
- Avoid changes to ADR-013 and W018 records.
- Run scoped Git and whitespace inspection.

## Done condition

- The selected diagram set is authored in the target Specification files.
- Each diagram is explanatory and aligned with DRMCP-ADR-MCP-013.
- No diagram adds package layout, Go signatures, structs, interfaces, functions, concrete algorithms, fixtures, tests, or implementation planning.
- No W018 decision is reopened.
- Scoped Git and whitespace inspection reports no whitespace failure for the changed files.

## Verification

- Confirmed the diagram set covers module topology, request path, Application orchestration, request-state visibility, Domain responsibility split, and Infrastructure source-family split.
- Confirmed no target Specification metadata date was updated because the diagrams are explanatory and do not change normative contract content.
- Confirmed DRMCP-ADR-MCP-013 was not changed.
- Confirmed W018 records were not changed.
- Scoped Git and whitespace inspection was run for the changed files.

## Evidence

Proposed diagram inventory and placement:

| diagram | file | purpose |
|---|---|---|
| Topology view | `spec:drmcp.implementation.contracts` | Navigation view of the five W018 module-contract domains and inward-owned source contract edge. |
| Lifecycle and request-state boundary | `spec:drmcp.implementation.contracts.composition_lifecycle.contract_boundary` | Separates server-lifetime dependency construction from request-state assembly. |
| Request path view | `spec:drmcp.implementation.contracts.mcp_inbound_adapter.contract_boundary` | Shows MCP mapping to one matching Application Use Case and no semantic reclassification. |
| Application-owned orchestration view | `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary` | Shows public use cases, shared orchestration, Domain collaborators, and source contracts. |
| Request-state visibility view | `spec:drmcp.implementation.contracts.application_use_cases.contract_boundary` | Shows allowed state readers, excluded observers, and request-end discard. |
| Domain responsibility split | `spec:drmcp.implementation.contracts.record_domain_logical_tree.contract_boundary` | Shows parser, logical tree, relation graph, resolution, local validation, and relation validation responsibilities. |
| Source-family split view | `spec:drmcp.implementation.contracts.infrastructure_io_adapters.contract_boundary` | Shows Current and Legacy source access separation and inward-owned contract implementation. |

No user judgment was needed before writing because the diagram set is an explanatory projection of accepted W018 authority.

Changed Specification files:

- `drmcp/records/spec/implementation/contracts/index.md`
- `drmcp/records/spec/implementation/contracts/composition-lifecycle/contract-boundary.md`
- `drmcp/records/spec/implementation/contracts/mcp-inbound-adapter/contract-boundary.md`
- `drmcp/records/spec/implementation/contracts/application-use-cases/contract-boundary.md`
- `drmcp/records/spec/implementation/contracts/record-domain-logical-tree/contract-boundary.md`
- `drmcp/records/spec/implementation/contracts/infrastructure-io-adapters/contract-boundary.md`

Workflow records created:

- `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-019-add-module-contract-explanatory-diagrams.md`
- `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-019-01-author-module-contract-explanatory-diagrams.md`

DRMCP is non-operational.
Filesystem authoring was the required fallback.

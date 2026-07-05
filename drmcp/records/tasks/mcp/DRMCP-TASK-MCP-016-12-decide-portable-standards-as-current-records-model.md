# DRMCP-TASK-MCP-016-12: Decide portable standards as Current Records model

- **id**: DRMCP-TASK-MCP-016-12
- **status**: done
- **date**: 2026-07-04
- **work_item**: DRMCP-WORK-MCP-016
- **task_type**: decision
- **estimate**: 1.0d
- **depends_on**:
  - DRMCP-TASK-MCP-016-11
- **outputs**:
  - DRMCP-TASK-MCP-016-12

## Goal

Decide the revised portable standards and Guidance architecture required to resolve T09 findings.

## Work

- Reconcile DRMCP-REQ-MCP-003 and DRMCP-ADR-MCP-001 with the W016 application architecture.
- Decide how the portable standards package enters Current Records.
- Decide the resulting top-level component model.
- Decide Guidance query, identity, and projection ownership.
- Decide request-state and lifecycle treatment.
- Decide the physical spec-tree mapping boundary without selecting concrete configuration serialization.
- Decide the required originating-artifact disposition.

### Decision inventory

| ID | topic | status | depends on | decision summary | reason / evidence | canonical target | ADR route |
|---|---|---|---|---|---|---|---|
| D-018 | Portable package runtime model | `decided` | none | Treat the portable Design Records standards package as ordinary Current Records under the reserved `design_records` app namespace. Use the normal current parsing, logical-tree, active-index, exact retrieval, resolution, and validation semantics. DRMCP performs no runtime ref rewrite. | PRODUCT-REQ-SPEC-003 already produces a ref-rewritten `spec:design_records.*` spec tree. Reusing Current Records avoids a second record model and index. | Requirement, component, dependency, runtime, discovery, identity, and Guidance contracts | required |
| D-019 | Physical package source mapping | `decided` | D-018 | Configuration associates the selected portable spec-tree root with `app_namespace: design_records` as a spec-only Current Records source. A records-root source and a spec-tree source use the same current Spec semantics after root mapping. Exact configuration serialization remains downstream. | The package root is `<exe-dir>/design-records/`, whose root `index.md` is the `spec:design_records` record. It is not physically `<app>/records/spec/`. | Requirement, namespace scanning, discovery, source mapping, identity mapping | required |
| D-020 | Whole-application component graph | `decided` | D-018 | Remove Guidance Domain as a top-level component. Use five components: Composition / Lifecycle, MCP Inbound Adapter, Application Use Cases, Record Domain / Logical Tree, and Infrastructure I/O Adapters. | Guidance uses ordinary Spec records and adds no independent semantic model, source lifecycle, or substitution boundary. | Component ADR and architecture boundary view | required |
| D-021 | Guidance operation ownership | `decided` | D-018, D-020 | Keep `list_authoring_guides` and `get_authoring_guidance` as operation-specific Application Use Cases over shared record-query orchestration. Fix scope to `design_records`, kind `spec`, and the `spec:design_records.authoring_standards.*` subtree. Public use cases do not call `list_records` or `get_records`. | The tools are aliases at the application-policy level, not separate storage or parsing capabilities. | Dependency ADR, Guidance schema and operation contracts | required |
| D-022 | Guidance identity and projection | `decided` | D-021 | Use the canonical package Spec ref as guide ID. Use first H1 text as title, `## What this is` body as abstract, and complete Markdown as verbatim content. Exclude the authoring-standards root from normal list results. | Canonical refs preserve package identity. The user explicitly selected `## What this is` as the abstract source. | Guidance schema and operation contracts | required |
| D-023 | Request state and lifecycle | `decided` | D-018, D-019 | Include the configured `design_records` spec-tree source in the same fresh immutable Current Records snapshot used by Read and Validation operations. Guidance consumes that snapshot. Do not create a package-specific index, snapshot, source port, cache, or lifecycle. Legacy Archive remains separate. | Current package Specs have the same record semantics. Only their configured physical root and app association differ. | Lifecycle ADR and runtime view | required |
| D-024 | Requirement and prior authority disposition | `decided` | D-018 through D-023 | Amend DRMCP-REQ-MCP-003 without changing its requirement identity. Retain DRMCP-ADR-MCP-001. Supersede ADR-007, ADR-008, and ADR-009 through replacement architecture ADRs. Preserve T03, T05, T06, T07, T08, and T09 as historical Evidence. | The portability need remains the same. The W016 component, ownership, and lifecycle choices change materially. | DRMCP-REQ-MCP-003, ADR routing, replacement ADRs | not_required |

### Status definitions

- `open`: A required choice is known but not selected.
- `in_discussion`: The one active user judgment.
- `decided`: The selected outcome is explicit.
- `blocked`: A named dependency prevents selection.
- `deferred`: The choice is intentionally routed outside this Work Item.
- `superseded`: A later decision item replaced the entry.

### Current cursor

- Decision: none
- Loop state: `decision_complete`
- Next action: Route D-018 through D-024 to ADR and canonical authoring.

## Done condition

- D-018 through D-024 are terminal.
- The package uses ordinary Current Records semantics without a second logical index.
- The physical spec-tree mapping is explicit without fixing concrete configuration serialization.
- Guidance is a thin Application alias over shared record-query orchestration.
- The five-component graph and request-state treatment are coherent.
- Requirement and ADR dispositions are explicit.
- No ADR or Specification body is authored.

## Verification

- Confirm no item remains `open`, `in_discussion`, or `blocked`.
- Confirm the selected model satisfies DRMCP-REQ-MCP-003 portability intent.
- Confirm package refs remain `spec:design_records.*` and are not rewritten by DRMCP.
- Confirm Current Records and Legacy Archive remain separate.
- Confirm public use cases do not call one another.
- Confirm exact package config fields, interfaces, types, and implementation details remain downstream.

## Evidence

- The repository-root `prompt_chappy.md` was read first.
- `CLAUDE.md` and `AGENTS.md` were not read.
- DRMCP is non-operational. Design Record access used filesystem fallback.
- T09 F-BLK-01 exposed the conflict between the portable-package authority and the direct PRODUCT directory Guidance model.
- T09 F-MAJ-01 exposed stale `## Abstract` projection text.
- PRODUCT-REQ-SPEC-003 defines a whole-tree copied package at `<exe-dir>/design-records/` with refs rewritten to `spec:design_records.*`.
- DRMCP-REQ-MCP-003 and DRMCP-ADR-MCP-001 require indexed package Specs as Guidance authority.
- The user decided that package Specs should receive ordinary Current Records treatment.
- The user decided that Guidance tools should be fixed-scope aliases over the normal record model.
- The user confirmed `guide.abstract = ## What this is section body`.
- The selected model removes Guidance-specific domain, source, index, snapshot, and lifecycle responsibilities.
- No canonical artifact was authored by this Task.
- No stage or commit was performed.

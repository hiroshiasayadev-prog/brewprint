# TRV-TASK-SPEC-005-02: Decide contract Specification topic tree and Markdown placement

- **id**: TRV-TASK-SPEC-005-02
- **status**: done
- **date**: 2026-07-03
- **work_item**: TRV-WORK-SPEC-005
- **task_type**: decision
- **estimate**: 1d
- **depends_on**:
  - TRV-TASK-SPEC-005-01
- **outputs**:
  - TRV-TASK-SPEC-005-02

## Goal

Decide one coherent `spec:trv` topic tree and Markdown placement for architecture-derived contract Specifications.

## Work

- Consume the completed TRV-INV-SPEC-003 findings without treating its recommendation as a decision.
- Decide the parent and child Specification refs for the W005 contract boundary.
- Decide the physical Markdown directory and file placement corresponding to those semantic refs.
- Assign one normative ownership boundary to every selected Specification file.
- Decide where application-port, application-model, validation-flow, outcome, MCP-adapter, record/checklist-adapter, model-provider-adapter, Task-input, caller, and compatibility contracts belong.
- Preserve PRODUCT semantic ownership and W002 architecture ownership.
- Keep exact implementation-ready representation and mechanics in W004.
- Record ADR-routing candidates without authoring ADRs.

This Task must not author or amend Specifications or ADRs, change the Task graph, perform review, closure, or implementation.

## Done condition

- Every placement item is `decided`, `deferred`, or validly `blocked`.
- The selected topic tree and physical Markdown placement are explicit.
- Every selected file has one coherent normative responsibility.
- Cross-file ownership and exclusions prevent duplicate contract authority.
- T03 can materialize exact writers without new placement judgment.

## Verification

- Confirm the decision consumes the Investigation evidence.
- Confirm semantic refs and physical paths are distinct but aligned.
- Confirm no canonical content was authored.
- Confirm PRODUCT, W002, and W004 ownership remains intact.

## Evidence

- T01 completed `TRV-INV-SPEC-003` and supplied the placement inventory.
- During D-004 discussion, the user identified that the first Investigation draft used one generic `external` area and omitted parallel record/checklist and model-provider adapter areas.
- Corrected `TRV-INV-SPEC-003` to align the candidate tree with all three W002 adapter components.
- Decision scope is limited to contract Specification topic tree, physical placement, normative ownership, and required downstream routing candidates.
- Loop state: `completed`.
- Current cursor: none.
- The standalone responsibility-boundary validator returned compliant results for all 23 common and decision criteria after all D-001 through D-010 decisions, the final tree, and `spec:trv.model_runtime` retirement were recorded.
- Result: `PASS`.

| ID | Topic | Status | Depends on | Decision summary | Reason | Canonical target | ADR route |
|---|---|---|---|---|---|---|---|
| D-001 | Contract-area root and top-level placement | `decided` | TRV-INV-SPEC-003 | Create one unified `spec:trv.contracts` root under `trv/records/spec/contracts/`. Decide all child classifications and files separately; D-001 does not authorize an `external` child or any other specific subtree. | One contract Overview preserves W005 as one independently reviewed contract area. Keeping contracts outside `application-architecture/` preserves the separate W002 and W005 completion boundaries. | `spec:trv.contracts` | `candidate` |
| D-002 | Application-contract file partition | `decided` | D-001 | Create `spec:trv.contracts.application` with one Overview and four port-aligned child Specifications: `validation_use_case`, `task_record_source`, `checklist_catalog`, and `model_evaluation`. Co-locate each boundary model with its owning port contract. Do not create a generic application-model catalog. | W002 defines distinct application-owned ports and failure boundaries. One shared adapter may implement the Task and checklist ports, but adapter grouping does not merge their contracts. Co-location avoids pushing exact representation into W005. | `spec:trv.contracts.application` and four child refs | `candidate` |
| D-003 | Application outcome ownership | `decided` | D-002 | Create `spec:trv.contracts.application.outcomes` as a dedicated application outcome contract. It owns the three mutually distinct application outcome classes, their allowed content boundary, validation-use-case construction ownership, and meaning-preserving MCP projection boundary. | Outcomes are shared by the validation use case and MCP projection. A dedicated contract prevents duplicate outcome definitions across application and MCP-adapter Specifications. PRODUCT remains the semantic authority; exact tagged unions remain W004-owned. | `spec:trv.contracts.application.outcomes` | `candidate` |
| D-004 | Adapter-area child partition | `decided` | D-001 | Create `spec:trv.contracts.adapters` with one Overview and three parallel child areas: `mcp`, `record_checklist`, and `model_provider`. Each child area uses an Overview or Index at `index.md`; detailed obligations belong to non-index child Specifications decided separately. | W002 defines the MCP, record/checklist, and model-provider adapters as sibling top-level components. Mirroring those boundaries avoids the incomplete generic `external` classification and keeps application-port contracts separate from adapter contracts. | `spec:trv.contracts.adapters`, `.mcp`, `.record_checklist`, and `.model_provider` | `candidate` |
| D-005 | MCP-adapter file partition | `decided` | D-004 | Create three non-index child Specifications under `spec:trv.contracts.adapters.mcp`: `interface`, `task_input`, and `caller_workflow`. Keep `index.md` as an Overview only. | The MCP interface owns tool, request/response, transport validation, outcome projection, and transport failure boundaries. Task identity and caller workflow are durable concerns with distinct ADR and PRODUCT authority, so keeping them separate avoids burying them in transport schema. Exact normalization, field syntax, and implementation mechanics remain W004-owned. | `spec:trv.contracts.adapters.mcp.interface`, `.task_input`, and `.caller_workflow` | `candidate` |
| D-006 | Record/checklist-adapter file partition | `decided` | D-004 | Create two non-index child Specifications under `spec:trv.contracts.adapters.record_checklist`: `task_record_access` and `checklist_access`. Keep `index.md` as an Overview only. | The adapter may be one implementation, but Task source access and checklist source access have different inputs, returned application models, and failure boundaries. Shared filesystem, parser, cache, and package mechanics remain W004-owned and do not require one merged contract. | `spec:trv.contracts.adapters.record_checklist.task_record_access` and `.checklist_access` | `candidate` |
| D-007 | Model-provider partition and legacy `model_runtime` retirement | `decided` | D-002, D-004 | Create two non-index child Specifications under `spec:trv.contracts.adapters.model_provider`: `ollama_adapter` and `ollama_runtime`. Retire and delete `spec:trv.model_runtime` and `trv/records/spec/model-runtime/index.md`; route the provider-neutral application port through `spec:trv.contracts.application.model_evaluation` and provider/runtime obligations through the two model-provider contracts. | The standalone model-runtime topic existed because W002 had no complete contract tree. The accepted W005 partition now gives the application port, TRV adapter, and external runtime distinct normative owners, so retaining an extra sibling Overview would add duplicate navigation without an independent responsibility. Exact endpoint, schema, retry, configuration, and implementation mechanics remain W004-owned. | `spec:trv.contracts.application.model_evaluation`, `spec:trv.contracts.adapters.model_provider.ollama_adapter`, and `.ollama_runtime`; retire `spec:trv.model_runtime` | `candidate` |
| D-008 | Compatibility contract placement | `decided` | D-001, D-004 | Create `spec:trv.contracts.compatibility` as one non-index cross-cutting Specification at `trv/records/spec/contracts/compatibility.md`. It defines which semantic behavior must survive transport, source, provider, adapter, and future DRMCP replacement, and which mechanics may change. | Compatibility spans application outcomes, PRODUCT semantics, caller judgment, invocation meaning, and replaceable adapter mechanics. Placing it under one adapter would misstate its scope, while a separate directory is unnecessary without child topics. | `spec:trv.contracts.compatibility` | `candidate` |
| D-009 | Stale W003 and retired model-runtime correction scope | `decided` | D-001, D-002, D-003, D-004, D-005, D-006, D-007, D-008 | Correct six retained current TRV Specifications: `spec:trv`, `spec:trv.application_architecture`, and its `component_model`, `dependency_model`, `validation_flow`, and `boundary` children. Register the new contract tree, replace active W003 handoffs, and replace references to `spec:trv.model_runtime` with the accepted application and model-provider contract refs. Delete `trv/records/spec/model-runtime/index.md`. Do not modify historical W002/W003 Work Items or Tasks as part of Specification correction. | The six retained Specifications contain active W003 handoffs or references to the now-redundant model-runtime topic. Retaining or merely shrinking that topic would duplicate the accepted contract tree. The correction changes current normative projection and navigation without changing the W002 component or dependency architecture. | Six retained TRV Specifications plus deletion of `spec:trv.model_runtime` | `candidate` |
| D-010 | Accepted ADR target alignment | `decided` | D-001, D-004, D-005, D-006, D-007, D-008, D-009 | Route non-material in-place amendments for TRV-ADR-SPEC-002 through TRV-ADR-SPEC-005. Preserve every selected alternative, core architecture, rationale, status, date, and historical Evidence. Update stale consequence and Specification-target statements: ADR-002 projects to `application.model_evaluation`, `model_provider.ollama_adapter`, and `model_provider.ollama_runtime`; ADR-003 projects to the MCP interface, Task-input, caller-workflow, and application-outcomes contracts; ADR-004 projects to MCP Task input and Task-record access; ADR-005 projects to the root compatibility contract. Do not create or supersede an ADR. Update `migrated_to_spec` only after canonical Specification authoring is complete. | ADR-002 through ADR-005 retain their durable decisions. Only their W003-era handoffs, obsolete `spec:trv.model_runtime` target, and pending-T05 consequence wording are stale. The routing rules permit non-material amendment when the selected alternative and core rationale remain valid and only references, consequences, or extracted responsibilities change. | TRV-ADR-SPEC-002 through TRV-ADR-SPEC-005 and the selected contract Specifications | `amend` |

### Selected Specification tree

```text
trv/records/spec/contracts/
  index.md
  compatibility.md
  application/
    index.md
    validation-use-case.md
    task-record-source.md
    checklist-catalog.md
    model-evaluation.md
    outcomes.md
  adapters/
    index.md
    mcp/
      index.md
      interface.md
      task-input.md
      caller-workflow.md
    record-checklist/
      index.md
      task-record-access.md
      checklist-access.md
    model-provider/
      index.md
      ollama-adapter.md
      ollama-runtime.md
```

Every `index.md` is an Overview or Index only.
Detailed normative obligations belong to the named non-index child Specifications.

Retained current Specifications to correct:

- `spec:trv`;
- `spec:trv.application_architecture`;
- `spec:trv.application_architecture.component_model`;
- `spec:trv.application_architecture.dependency_model`;
- `spec:trv.application_architecture.validation_flow`;
- `spec:trv.application_architecture.boundary`.

Retired current Specification:

- delete `spec:trv.model_runtime` at `trv/records/spec/model-runtime/index.md` after its obligations are projected into the selected contract tree.

Status definitions:

- `open`: Unblocked judgment not yet discussed.
- `in_discussion`: The single current user judgment.
- `decided`: Explicit outcome persisted.
- `blocked`: Named authority or dependency prevents progress.
- `deferred`: Explicitly excluded from W005.

Stop conditions:

- A selected placement conflicts with path-derived Specification identity.
- A choice changes W002 architecture or PRODUCT semantics.
- A contract boundary requires W004 implementation-ready detail.
- An accepted ADR requires material reversal without a supersession route.

Expected downstream route:

- T03 graph coordination after every item is terminal.
- Conditional ADR routing or amendment.
- Serialized Specification authoring and stale-reference correction.
- One integrated independent review.
- Verdict-gated closure synchronization.

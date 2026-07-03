# TRV-INV-SPEC-003: Architecture-derived contract Specification topic tree and placement

- **status**: concluded
- **date**: 2026-07-03
- **trigger**: TRV-WORK-SPEC-005 requires placement-first contract Specification design after the W003 route was retired.
- **scope**: Inventory W002 contract boundaries and compare coherent `spec:trv` topic trees and Markdown placement.
- **non_scope**: Do not select a topic tree, author normative contracts, change architecture, or define implementation-ready representations.
- **source_refs**:
  - TRV-WORK-SPEC-002
  - TRV-WORK-SPEC-003
  - TRV-WORK-SPEC-004
  - TRV-WORK-SPEC-005
  - spec:trv
  - spec:trv.application_architecture
  - spec:trv.model_runtime
  - spec:product.responsibility_boundary_validator
- **follow_up_candidates**:
  - TRV contract Specification topic-tree and placement decision
  - TRV contract ADR routing and canonical authoring route
- **related_work_items**:
  - TRV-WORK-SPEC-002
  - TRV-WORK-SPEC-003
  - TRV-WORK-SPEC-004
  - TRV-WORK-SPEC-005
- **related_adrs**:
  - TRV-ADR-SPEC-001
  - TRV-ADR-SPEC-002
  - TRV-ADR-SPEC-003
  - TRV-ADR-SPEC-004
  - TRV-ADR-SPEC-005
- **related_specs**:
  - spec:trv
  - spec:trv.application_architecture.component_model
  - spec:trv.application_architecture.dependency_model
  - spec:trv.application_architecture.validation_flow
  - spec:trv.application_architecture.boundary
  - spec:trv.model_runtime
  - spec:product.responsibility_boundary_validator

## Investigation scope

This Investigation answers one question:

> Which `spec:trv` topic tree and Markdown placement can own every architecture-derived contract without duplicating PRODUCT semantics, W002 architecture, or W004 detail?

The investigation covers application ports, boundary models, validation-flow handoffs, application outcomes, the three W002 adapter boundaries, caller contracts, and compatibility.

## Out of scope

- Selecting the final topic tree or file set.
- Authoring or amending ADRs or Specifications.
- Changing W002 architecture decisions or ownership.
- Defining exact Go declarations, schemas, algorithms, configuration names, commands, or tests.
- Designing current DRMCP integration.
- Changing the retired W003 Task graph.

## Background

W002 closed the reviewed application architecture.
W002 defines five top-level components and four application ports.
W002 also defines application models, validation-flow stages, and three outcome classes.

W003 routed and authored durable external-interface ADRs.
W003 then preselected four flat Specification targets before deciding the architecture-derived contract tree.
The user retired that route because the four targets did not cover the application contracts derived from W002.

W005 replaces W003 contract Specification authoring.
W005 requires topic-tree and Markdown-placement judgment before normative authoring.

## What was investigated

The investigation reviewed:

- `spec:trv` and the complete W002 architecture Specification set;
- `spec:trv.model_runtime`;
- the MCP, record/checklist, and model-provider adapter boundaries defined by W002;
- TRV-ADR-SPEC-001 through TRV-ADR-SPEC-005;
- the PRODUCT semantic validator contract;
- the retired W003 Work Item and its four-file writer;
- the W004 detailed-design boundary;
- active Specification path and authoring rules.

The investigation compared four placement patterns:

1. the retired flat external-first pattern;
2. peer application-contract and adapter-contract areas;
3. one unified `contracts/` area with application and adapter children;
4. contract files nested under the architecture tree.

## Findings

### Correction discovered during T02

The first Investigation draft grouped MCP, Task input, and caller behavior under one `external` area.
That grouping was incomplete because it did not give the record/checklist adapter and model-provider adapter parallel contract areas.

W002 defines the MCP adapter, record/checklist adapter, and model-provider adapter as sibling top-level components.
The corrected candidate trees therefore use one adapter-aligned partition with a child area for each adapter.
This correction does not change the recommendation for one unified `spec:trv.contracts` root, but it changes the required child tree and follow-up judgments.

### Architecture-derived contract inventory

| architecture item | observed W002 state | required W005 owner or exclusion | duplication boundary |
|---|---|---|---|
| Inbound validation-use-case port | Application-owned entry boundary used by the MCP adapter. | Application validation-use-case contract. | W002 keeps port ownership and dependency direction. W004 owns exact interface declarations. |
| Application invocation input | MCP data becomes application-owned input before use-case invocation. | Application validation-use-case contract. | External request fields belong to the MCP contract. Exact structs belong to W004. |
| Validation use case | Orchestrates one complete Task responsibility-boundary evaluation. | Application validation-use-case contract. | PRODUCT owns semantic criteria and aggregation. W002 owns stage ownership. |
| Task-record source port | Requests Task content and declared `task_type`. | Application Task-record source contract. | The record/checklist adapter contract owns source access. W004 owns filesystem and parsing mechanics. Current DRMCP is excluded. |
| Task boundary model | Carries persisted Task content and declared `task_type`. | Co-locate with the Task-record source contract. | PRODUCT owns Task semantics. W004 owns exact data types. |
| Checklist-catalog port | Requests common and declared-type criteria. | Application checklist-catalog contract. | The record/checklist adapter contract owns source access. PRODUCT owns checklist composition. W004 owns storage and discovery mechanics. |
| Checklist boundary model | Carries applied criteria into the application core. | Co-locate with the checklist-catalog contract. | Exact criterion representation belongs to W004. |
| Model-evaluation port | Accepts one complete prompt and returns candidates or execution failure. | Application model-evaluation contract. | The model-provider adapter contract owns Ollama-specific translation and execution. W002 owns provider-neutral direction. W004 owns exact schemas and retries. |
| Evaluation candidate model | Carries decoded provider output without Ollama types. | Co-locate with the model-evaluation contract. | PRODUCT owns completeness and correspondence semantics. W004 owns exact representation. |
| Complete prompt handoff | The application core builds the prompt before provider invocation. | Validation-use-case and model-evaluation contracts must share one non-duplicated boundary statement. | Exact prompt text and correlation encoding belong to W004. |
| Application precondition handoff | The use case validates preconditions before checklist retrieval. | Application validation-use-case contract. | Structural parsing mechanics belong to the source adapter and W004. |
| Provider-result validation handoff | The use case validates completeness and criterion correspondence after decoding. | Application validation-use-case contract referencing PRODUCT authority. | Do not restate PRODUCT semantic rules. |
| Structural precondition failure | Application-owned outcome with PRODUCT-defined meaning. | Application outcome contract or a dedicated outcome section in the validation-use-case contract. | PRODUCT owns meaning. MCP owns external projection only. |
| Completed semantic evaluation | Application-owned criterion results and mechanically derived overall result. | Application outcome contract or validation-use-case outcome section. | PRODUCT owns criterion and aggregation semantics. |
| Execution failure | Application-owned outcome when complete semantic evaluation cannot finish. | Application outcome contract or validation-use-case outcome section. | Provider mechanics belong to the model-provider adapter contract and W004; `spec:trv.model_runtime` only routes the boundary. |
| MCP input validation | MCP adapter rejects invalid MCP or JSON shape before application invocation. | MCP-adapter interface contract. | Exact schema syntax and library types belong to W004. |
| MCP outcome projection | MCP adapter maps application outcomes without changing meaning. | MCP-adapter interface contract. | Application outcome construction remains inside the application contract. |
| Public Task identity | Accepted ADR uses one repository-root-relative `task_path`. | MCP-adapter Task-input contract. | Normalization, symlink, separator, and filesystem mechanics belong to W004. |
| Caller human judgment | Violations require caller-owned acceptance or rejection. | MCP-adapter caller-workflow contract referencing PRODUCT-ADR-SPEC-017. | PRODUCT owns the human exception semantics. |
| Future DRMCP compatibility | Semantic behavior survives replaceable transport and source mechanics. | Cross-cutting compatibility contract. | Current DRMCP contract remains unchanged. |
| Record and checklist source access | One adapter implements Task and checklist source mechanics behind two application ports. | Record/checklist adapter area with distinct Task-access and checklist-access contracts, unless T02 accepts a bounded combined contract. | Adapter grouping does not merge the two application-port responsibilities. Exact parser and storage mechanics remain W004-owned. |
| Model-provider adapter and Ollama runtime boundary | The adapter owns provider translation and execution; external deployment owns Ollama and model files. | Model-provider adapter area with an Ollama-adapter contract and a bounded external-runtime contract or section. | W005 owns the adapter-facing obligation, not installation or process supervision. Exact protocol and configuration remain W004-owned. |
| Startup and composition | Startup loads configuration and wires top-level components. | Explicit W005 exclusion. | W002 owns responsibility. W004 owns exact configuration, constructors, wiring, and lifecycle mechanics. |
| Static component and dependency graph | W002 fixes component ownership and inward dependencies. | W005 exclusion with architecture references. | Contract authoring must not restate architecture diagrams or ownership tables. |

Every W002 application port, model boundary, validation handoff, application outcome, and adapter boundary has a candidate normative owner or explicit exclusion.

### Current authority and repository gaps

| artifact | observed state | proposed classification | next owner |
|---|---|---|---|
| `spec:trv` | Topic map still routes remaining contract design to W003. | `stale_representation` candidate. | Later W005 Specification authoring. |
| `spec:trv.application_architecture.boundary` | W003 is named as external MCP and application-contract owner. | `stale_representation` candidate. | Later W005 Specification authoring or bounded architecture-reference correction. |
| `spec:trv.application_architecture.validation_flow` | External MCP contract routes to W003. Application port details route broadly to W004. | Mixed stale reference and ownership gap candidate. | T02 placement decision, then authoring coordination. |
| `spec:trv.model_runtime` | Its `index.md` currently defines detailed model-port, adapter, and runtime obligations. | Overview-detail mismatch and overlap risk, not an architecture contradiction. | T02 must route detail into application and model-provider contracts, then reduce this file to Overview-level ownership and navigation. |
| TRV-ADR-SPEC-003 through TRV-ADR-SPEC-005 | Accepted durable choices exist, but `migrated_to_spec` is null. | Consistent pending normative projection. | Later ADR routing and Specification authoring. |
| W003 T05 | Preselects four flat external-oriented Specification targets. | Historical workflow route, not accepted placement authority. | No writer. W005 replaces it. |
| W004 | Waits for reviewed W005 contract closure and owns exact representations and mechanics. | Consistent. | No W005 change required. |
| PRODUCT semantic validator | Owns criterion selection, results, aggregation, outcome meanings, invocation, and human exception semantics. | Controlling exclusion authority. | W005 references without restatement. |

No accepted architecture rule conflicts with W005.
The main risk is duplicate authority across architecture, application-port contracts, adapter contracts, and detailed design.

### Candidate A: flat external-first tree

Candidate paths:

```text
trv/records/spec/
  mcp-interface/index.md
  task-input/index.md
  caller-integration/index.md
  compatibility/index.md
```

Candidate refs:

- `spec:trv.mcp_interface`
- `spec:trv.task_input`
- `spec:trv.caller_integration`
- `spec:trv.compatibility`

| criterion | result |
|---|---|
| External contract cohesion | Strong. The tree matches TRV-ADR-SPEC-003 through TRV-ADR-SPEC-005. |
| W002 application-port coverage | Insufficient. No owner exists for inbound, Task, checklist, or model application contracts. |
| Application-model coverage | Insufficient. Boundary models would remain in architecture prose or drift into W004. |
| Duplication risk | High. External files would absorb application meaning to become complete. |
| Navigation | Simple but misleading. The tree presents external delivery as the whole contract surface. |
| W004 projection | Weak. W004 would need to infer several missing application contracts. |
| Review boundary | Too narrow for W005 Completion Conditions. |

Candidate A cannot satisfy W005 without adding application-contract and non-MCP adapter-contract owners.

### Candidate B: peer application and adapter areas

Candidate paths:

```text
trv/records/spec/
  application-contracts/
    index.md
    validation-use-case.md
    task-record-source.md
    checklist-catalog.md
    model-evaluation.md
    outcomes.md
  adapter-contracts/
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
  compatibility.md
```

Candidate refs follow the path-derived hierarchy.

| criterion | result |
|---|---|
| Application contract cohesion | Strong. Ports and application outcomes receive direct owners. |
| Adapter contract cohesion | Strong. Each W002 adapter receives a parallel child area. |
| PRODUCT separation | Strong when application files reference PRODUCT semantics. |
| Architecture separation | Strong. Contract topics remain siblings of architecture topics. |
| Navigation | Moderate. `spec:trv` must route multiple peer contract areas. |
| Cross-cutting compatibility | Clear as a peer topic. |
| Duplication risk | Moderate. Each application port and implementing adapter need an explicit non-overlap rule. |
| W004 projection | Strong. Exact declarations and mechanics can map to stable contract topics. |
| Review boundary | Coherent, but no single contract-area Overview exists. |

Candidate B is viable.
The main trade-off is fragmented top-level navigation.

### Candidate C: unified contracts root

Candidate paths:

```text
trv/records/spec/contracts/
  index.md
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
  compatibility.md
```

Candidate refs follow the path-derived hierarchy under `spec:trv.contracts`.
Every `index.md` is an Overview or Index only; detailed obligations live in non-index child Specifications.

| criterion | result |
|---|---|
| Contract-area cohesion | Strong. One Overview owns navigation and cross-boundary rules. |
| Application and adapter separation | Strong. Application-owned ports remain separate from adapter-owned mechanics. |
| Adapter alignment | Strong. MCP, record/checklist, and model-provider areas mirror the three W002 adapter components. |
| PRODUCT separation | Strong when the application area references PRODUCT authority. |
| Architecture separation | Strong. Contracts remain downstream siblings of the architecture area. |
| Navigation | Strong. `spec:trv` registers one contract area. |
| Cross-cutting compatibility | Strong. Compatibility can sit at the contract root. |
| Duplication risk | Lowest among compared patterns when the Overview defines port-versus-adapter ownership. |
| W004 projection | Strong. W004 can map exact types and mechanics to stable child contracts. |
| Review boundary | Strong. One integrated contract tree matches W005. |
| Cost | Highest path depth and largest initial file set. |

Candidate C is viable.
File granularity inside each adapter area still requires T02 judgment.

### Candidate D: contracts nested under application architecture

Candidate paths:

```text
trv/records/spec/application-architecture/contracts/
  ...
```

| criterion | result |
|---|---|
| Traceability to W002 | Strong. The physical tree directly follows architecture. |
| Adapter and caller placement | Weak. Adapter-facing and caller-facing contracts are not architecture views. |
| Independent change boundary | Weak. Contract changes would appear subordinate to architecture changes. |
| `spec:trv.model_runtime` alignment | Weak. The sibling runtime topic would sit outside its related contract tree. |
| Duplication risk | High. Architecture and contract prose would share one topic area. |
| Review boundary | Misaligned with separate W002 and W005 completion judgments. |

Candidate D conflicts with the accepted separation between architecture and contract closure.

### File-granularity findings

| granularity question | evidence | judgment needed |
|---|---|---|
| One file per application port | Task and checklist ports have distinct responsibilities and failure boundaries. | Decide separate files or one record-access contract with clearly separated sections. |
| Dedicated outcomes file | Three outcomes cross application and MCP projection boundaries. | Decide dedicated application outcome contract or validation-use-case ownership. |
| Dedicated application-models file | Models are boundary-specific and may become a miscellaneous type catalog. | Prefer co-location unless T02 finds a coherent independent model contract. |
| Adapter-area partition | W002 defines MCP, record/checklist, and model-provider as sibling top-level adapters. | Prefer one `adapters/` Overview with one child area per adapter; reject one generic `external/` bucket. |
| MCP-adapter file partition | MCP owns transport validation, Task input, outcome projection, and caller-facing behavior. | Decide interface, Task-input, and optional caller-workflow child files. |
| Record/checklist-adapter file partition | One adapter implements two application ports with different source responsibilities. | Decide separate Task-access and checklist-access files or one bounded combined source-access contract. |
| Model-provider file partition | Ollama translation/execution and the external runtime boundary are distinct from the application port. | Decide separate Ollama-adapter and runtime-boundary files or one bounded provider contract. |
| `model_runtime` index boundary | The current `index.md` contains detailed port and provider rules despite index/Overview placement. | Reduce it to an Overview that routes to application and model-provider child contracts. |
| Compatibility placement | Compatibility spans application semantics, replaceable adapters, and future DRMCP. | Decide contract-root child or another explicit cross-cutting location. |
| Caller workflow file | PRODUCT owns acceptance semantics, while TRV owns concrete interaction. | Decide whether caller workflow remains separate from MCP response behavior. |

### Shared-writer candidates

| target | candidate writers | required order candidate | risk |
|---|---|---|---|
| `spec:trv` | Contract-tree registration writer and later W004 registration writer. | W005 registration before W004. | Concurrent edits can lose Topics rows or stale state. |
| `spec:trv.application_architecture.boundary` | W005 stale-reference correction and any later architecture reconvergence writer. | W005 bounded correction after T02 placement. | Contract ownership can continue pointing to retired W003. |
| `spec:trv.application_architecture.validation_flow` | W005 handoff-reference correction and future architecture writer. | W005 bounded correction after contract refs exist. | Flow can route contracts to wrong owners. |
| `spec:trv.model_runtime` | W005 Overview correction writer and future W004 detailed writer. | W005 reduces it to navigation after application and model-provider refs are fixed. W004 follows reviewed W005. | Detailed port and provider rules can remain duplicated in an Overview. |
| Application outcome contract and MCP-adapter contract | Separate W005 writers when authoring is split. | Application outcome writer before MCP projection writer. | MCP projection may redefine or synthesize application meaning. |

T03 must serialize any accepted shared writers.

## Cross-cutting observations

- The old four-file tree is an MCP-facing interface tree, not a complete architecture-derived contract tree.
- Architecture-derived contracts should follow declared application ports and the three top-level adapter components, not an ad hoc `external` grouping or shared implementation convenience.
- Boundary models are best owned beside their producing or consuming contract.
- A generic application-model catalog risks becoming W004 type design under a W005 name.
- PRODUCT semantics must remain references, not copied rules.
- W002 may state coarse port payload classes without owning full contract obligations.
- W005 needs one explicit rule for architecture-versus-contract duplication.
- W004 already has the correct downstream dependency on reviewed W005 closure.
- Current DRMCP remains an exclusion, except as a future compatibility subject.

## Follow-up judgment candidates

| candidate | question | options evidenced by this Investigation |
|---|---|---|
| J-001 | What is the contract-area root? | Peer top-level areas, or one `spec:trv.contracts` root. |
| J-002 | How are application contracts partitioned? | One file per port, grouped record-access contracts, or another bounded partition. |
| J-003 | Who owns application outcomes? | Dedicated outcomes Specification, or validation-use-case contract sections. |
| J-004 | What are the adapter-area children? | One `adapters/` area with MCP, record/checklist, and model-provider children, or another architecture-aligned partition. |
| J-005 | How is the MCP-adapter area partitioned? | Separate interface, Task-input, and caller-workflow contracts, or bounded combinations that preserve ownership. |
| J-006 | How is the record/checklist-adapter area partitioned? | Separate Task-access and checklist-access contracts, or one bounded combined source-access contract. |
| J-007 | How are model-provider contracts related to `spec:trv.model_runtime`? | Keep `model_runtime` as an Overview and use separate Ollama-adapter and runtime-boundary contracts, or another non-duplicating partition. |
| J-008 | Where does compatibility live? | Contract-root child, top-level peer, or an explicitly justified adapter child. |
| J-009 | Which stale W003 references are corrected by W005? | `spec:trv`, architecture boundary, validation flow, model runtime, or a narrower subset. |
| J-010 | Do accepted ADRs require amendment after nested refs are chosen? | Reuse unchanged, bounded target-reference amendment, or broader ADR routing. |

T02 owns these judgments.

## Recommendation

Candidate C appears preferable because it matches W005 as one independently reviewed contract area.

The application and adapter subtrees should remain separate.
The adapter subtree should mirror the MCP, record/checklist, and model-provider components.
Boundary models should usually stay with their owning port or use-case contract.
A generic application-model file appears unnecessary unless T02 identifies an independent normative model boundary.

The recommendation is not a decision.
Candidate B remains viable when shallower top-level paths are preferred over one contract registry.

## Follow-up artifact candidates

- T02 decision ledger for J-001 through J-010.
- T03 graph coordination after the final paths and owners are fixed.
- One contract-area Overview when T02 selects a unified root.
- Bounded application-contract Specifications for the selected port and outcome partition.
- Adapter Overviews and bounded MCP, record/checklist, model-provider, caller, and compatibility Specifications.
- ADR routing for TRV-ADR-SPEC-003 through TRV-ADR-SPEC-005 target alignment.
- Bounded corrections for stale W003 references in current Specifications.
- One integrated independent review after the final writer.

## Open questions

- Whether each of the three adapter areas needs all proposed non-index child files.
- Whether the record/checklist adapter uses separate Task-access and checklist-access contracts or one bounded combined contract.
- Whether the model-provider runtime boundary needs a separate file from the Ollama-adapter contract.
- Whether `spec:trv.model_runtime` retains any normative detail beyond Overview-level routing.
- Whether caller interaction has enough concrete TRV behavior for a separate Specification.
- Whether nested contract refs require ADR amendment or only new Specification migration records.

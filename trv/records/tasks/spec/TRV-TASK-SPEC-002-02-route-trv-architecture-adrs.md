# TRV-TASK-SPEC-002-02: Route TRV architecture ADRs

- **id**: TRV-TASK-SPEC-002-02
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-002
- **task_type**: decision
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-002-01
- **outputs**:
  - TRV-TASK-SPEC-002-02

## Goal

Produce one complete ADR-routing ledger for the accepted TRV application-architecture decisions.

## Work

- Treat TRV-TASK-SPEC-001-02 D-017, D-005, and D-011 as fixed inputs.
- Inspect accepted PRODUCT ADRs only for actual reusable coverage.
- Classify each decision as `required`, `covered`, `not_required`, or validly `blocked`.
- Select `create`, `amend`, `reuse`, or `supersede` for every ADR-owned choice.
- Partition coherent ADR boundaries for:
  - ports-and-adapters architecture and inward dependency direction;
  - model-provider port and externally managed remote Ollama ownership;
  - logical component ownership and startup/dependency-wiring boundary.
- Combine decisions only when their alternatives, rationale, consequences, and supersession lifecycle are inseparable.
- Record exact canonical architecture Specification targets.
- Preserve exact package, file, symbol, interface, and implementation choices for W004.

This Task must not:

- reopen the selected architecture;
- author or amend an ADR or Specification body;
- make contract or detailed-design decisions;
- change the Task graph;
- perform review, correction, synchronization, implementation, stage, or commit work.

## Done condition

- D-017, D-005, and D-011 each have one terminal ADR route.
- Every `required` decision belongs to one coherent ADR boundary.
- Every `covered` decision names one accepted non-superseded ADR.
- Every `not_required` decision records a reason and canonical target.
- Every create, amend, reuse, or supersede disposition is explicit.
- T03 can materialize exact authoring and review owners without new routing judgment.

## Verification

- Confirm no architecture decision was changed.
- Confirm existing ADR coverage is exact rather than approximate.
- Confirm ADR boundaries can be superseded independently when appropriate.
- Confirm no contract, detailed-design, or implementation content entered the ledger.
- Confirm no ADR or Specification body changed.

## Evidence

- TRV-TASK-SPEC-001-02 contains the terminal architecture decisions.
- TRV-TASK-SPEC-002-01 created this routing owner.
- PRODUCT-ADR-SPEC-016 establishes TRV app-local ownership but does not select TRV internal architecture, dependency direction, provider seam, or runtime topology.

| decision | route | disposition | ADR boundary | canonical target | reason |
|---|---|---|---|---|---|
| D-017 | `required` | `create` | TRV-ADR-SPEC-001 | `spec:trv.application_architecture` | Ports and adapters, inward dependencies, and the application-core boundary are durable architecture choices with viable alternatives. |
| D-011 | `required` | `create` | TRV-ADR-SPEC-001 | `spec:trv.application_architecture` | Logical component ownership is inseparable from the selected dependency architecture and should share one supersession boundary. |
| D-005 | `required` | `create` | TRV-ADR-SPEC-002 | `spec:trv.model_runtime` | Provider-neutral core ownership and externally managed remote Ollama topology can evolve independently from the component architecture. |

Architecture ADR boundaries:

- `TRV-ADR-SPEC-001`: Use ports and adapters with explicit TRV component ownership and inward dependency direction.
- `TRV-ADR-SPEC-002`: Use a provider-neutral model port with an externally managed remote Ollama HTTP runtime.
- Both ADRs depend on PRODUCT-ADR-SPEC-016 for the accepted PRODUCT-versus-TRV ownership split.
- No existing accepted ADR fully covers either TRV-local choice.
- No ADR amendment or supersession is required.
- Exact packages, source files, symbols, and concrete interfaces remain W004-owned.
- Result: `PASS`.

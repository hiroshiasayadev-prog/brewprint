# TRV-WORK-SPEC-004: Define TRV implementation-ready detailed Specifications

- **id**: TRV-WORK-SPEC-004
- **status**: blocked
- **date**: 2026-07-03
- **source_refs**:
  - TRV-TASK-SPEC-001-06
  - TRV-TASK-SPEC-001-02
  - TRV-REQ-SPEC-001
  - TRV-WORK-SPEC-005
- **impact_refs**:
  - spec:trv
- **tasks**:
  - TRV-TASK-SPEC-004-01

## Goal

Define and independently review implementation-ready TRV detailed Specifications after architecture and contract closure.

## Boundary

This Work Item owns:

- the remaining detailed-design decision inventory and interactive decision loop;
- exact package, directory, file, symbol, interface, type, constructor, and dependency-wiring contracts;
- exact Task parsing, path safety, checklist discovery, prompt composition, and criterion-identity behavior;
- exact Ollama HTTP request, response, schema, retry, timeout, and failure handling;
- exact MCP transport models and result projection;
- exact environment configuration, startup validation, launcher, and external `mcp-proxy` contract;
- exact test, fixture, fake HTTP server, build, stdio smoke, proxy smoke, and optional live-smoke procedures;
- detailed-design ADR routing and required ADR authoring;
- detailed Specifications and namespace projection;
- one integrated independent detailed-design review;
- detailed-design lifecycle and closure synchronization.

This Work Item does not own:

- changes to reviewed architecture or application contract without an explicit return-to-decision route;
- production implementation, implementation Task authoring, or implementation Work Item decomposition;
- current DRMCP integration.

## Impact Scope

| target | impact |
|---|---|
| `spec:trv` | Receives reviewed detailed-Specification registration and implementation-readiness state. |
| Future TRV detailed-design ADRs | Record durable detailed choices when routing requires them. |
| Future TRV detailed Specifications | Define exact source, interface, schema, failure, configuration, launcher, test, and command contracts. |

## Task flow

```text
parent T15 W005 contract-Specification execution complete
  -> TRV-TASK-SPEC-004-01 child graph coordination
     -> detailed-design decision loop
     -> post-decision ADR routing, authoring, integrated review, and closure coordination
```

T01 materializes only the decision and post-decision owners. Later Tasks remain child-owned.

## Task Candidates

| task | task type | responsibility | dependency |
|---|---|---|---|
| `TRV-TASK-SPEC-004-01` | `coordination` | Materialize the detailed decision and post-decision graph owners. | Parent T15 |

## Completion Condition

- Every remaining implementation-facing design judgment is terminal.
- Reviewed architecture and contract remain unchanged or use an explicit reconvergence route.
- Every detailed decision has an ADR routing outcome.
- Required detailed-design ADRs are accepted.
- Current TRV detailed Specifications define exact implementation targets and verification commands.
- The final writer map is sufficient for later implementation decomposition.
- One integrated independent review returns `PASS`, or every required finding is independently closed.
- Lifecycle, Evidence, relations, and Work Item closure are synchronized.
- No production implementation or implementation Work Item is created.

## Evidence

- TRV-TASK-SPEC-001-02 fixed the detailed-design topics but did not claim implementation readiness.
- TRV-TASK-SPEC-001-02 D-015 requires detailed-Specification closure before implementation decomposition.
- TRV-TASK-SPEC-001-06 created this independent detailed-design boundary.
- TRV-TASK-SPEC-004-01 was created as the child-local detailed-design graph owner.
- W003 and parent T09 were retired.
- Parent T15 gates the child route until reviewed W005 contract-Specification closure.
- Detailed decision and authoring Tasks remain unmaterialized until T01 executes.
- TRV-ADR-SPEC-006 blocks detailed design while semantic-validator delivery is suspended.
- Resume requires reviewed controlled vocabulary, renewed semantic-feasibility evidence, and an explicit restoration decision.

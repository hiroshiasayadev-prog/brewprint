# Concept: Failure and evolution

- **id**: `spec:drmcp.implementation.application_architecture.failure_and_evolution`
- **status**: draft
- **date**: 2026-07-04
- **parent**: `spec:drmcp.implementation.application_architecture`

## What this is

Failure ownership and architecture-change boundary for DRMCP. This view owns failure classes, trustworthy-result semantics, downstream-local refinement, and architecture-return triggers.

## Concept model

### Failure classes

| failure class | source and owner | required handling |
|---|---|---|
| Request rejection | MCP Inbound Adapter owns protocol and schema decoding failures. Application Use Cases own operation-specific request invalidity. | Reject the request. Do not return a normal operation response. |
| Expected semantic outcome | Record Domain / Logical Tree produces modeled record states or findings. Application Use Cases select the operation-specific projection, including Guidance-specific outcomes. | Return the declared normal result, diagnostic, partial success, or operation error defined by the use case. |
| Execution failure | Infrastructure I/O Adapters report concrete access failures. Application Use Cases judge whether the promised result remains complete and trustworthy. | Fail the operation when mandatory state or required source-backed evidence cannot be constructed. |

Expected semantic outcomes include unresolved identity, unsupported input, duplicate conflict, invalid records or relations, unreadable indexed targets, zero matches, and guide-not-found meaning.

The MCP Inbound Adapter encodes the result or error selected by Application. The adapter does not reclassify semantic outcomes or execution failures.

### Trustworthy-result rule

An operation may continue only when it can produce the complete trustworthy result promised by its contract.

An operation fails when mandatory configuration, source state, record state, or required diagnostic location cannot support that result.
The operation does not return partial normal data from incomplete or untrustworthy state.

## Rules

### Downstream-local refinement

Downstream module contracts and detailed specifications may refine these concerns inside the accepted architecture:

- package and file layout;
- concrete interfaces and data types;
- parser, resolver, and validator algorithms;
- adapter APIs and filesystem libraries;
- rule encoding;
- request-local snapshot helper structure;
- error-code representation;
- performance refinements that preserve accepted ownership, freshness, and lifetime semantics.

### Architecture-return triggers

Work must return to application-architecture decision when a change would:

- add or remove a top-level component;
- reverse or bypass an accepted dependency;
- move responsibility across component boundaries;
- introduce request-spanning mutable state;
- introduce a new resource lifecycle;
- change source authority or introduce a package-specific record model;
- change Current Records and Legacy Archive separation;
- require cross-use-case coordination outside shared application orchestration;
- change trustworthy-result or failure semantics;
- make an external provider or process part of runtime behavior.

The deferred proposal/body-cache boundary in `spec:drmcp.implementation.application_architecture.runtime_and_state` is an architecture-return trigger.
The deferred write-transaction and post-write validation boundary in that view is also an architecture-return trigger.

Deferred authoring internals are not active failure contracts. This view does not define proposal errors, transaction failures, rollback behavior, or retained-state recovery.

## Boundary

This view defines who classifies and selects failure behavior at application level. Operation specifications own exact result schemas, diagnostic forms, and error codes.

A local refinement remains local only when it preserves the five-component graph, inward dependency direction, responsibility ownership, state lifetime, Current-versus-Legacy separation, and trustworthy-result rule.

## Non-goals

- Exact error codes or payloads.
- Operation-specific diagnostic schemas.
- Retry policy.
- Proposal-state failure behavior.
- Write atomicity, rollback, or recovery design.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.implementation.application_architecture` | Parent Overview and authoritative four-view map. |
| `spec:drmcp.implementation.application_architecture.application_boundary_and_components` | Owns the five-component graph and application boundary. |
| `spec:drmcp.implementation.application_architecture.dependency_and_responsibility` | Owns dependency direction and policy placement. |
| `spec:drmcp.implementation.application_architecture.runtime_and_state` | Owns request-state lifetime and deferred authoring boundaries. |

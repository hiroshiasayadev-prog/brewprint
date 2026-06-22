# Overview: Edge definitions

- **id**: `spec:bpdsl.dsl.edges.overview`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.overview`

## What this is

Entry point for BPDSL edge specs. Covers the four edge kinds that connect nodes: data flow wiring, state transitions, cross-layer access, and cross-file references.

## Current contract

Each node definition in `nodes:` holds signature only — types and names. All inter-node wiring is consolidated in separate sections, not embedded in node definitions (V01-ADR-015).

### Edge kind summary

| kind | YAML section | layer | description |
|---|---|---|---|
| Data flow | `flow:` | Processing | Wiring between nodes in the DAG. Declares which output feeds which input. |
| State transitions | `transitions:` | Application | FSM state machine transitions triggered by events. |
| Cross-edges | `reads:` / `writes:` fields on task | Cross-layer | Task read/write access to stores. Declared on the task node, not in flow. |
| Cross-file references | QualifiedID | Any | Module-crossing node references using full-path IDs. |

`flow:` is established as the DAG wiring keyword for the Processing layer; the Application layer uses `transitions:` to keep the two orthogonal (V01-ADR-019).

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Data flow edges | Reference | `spec:bpdsl.dsl.edges.data_flow` | `flow:` entry types (step, fork, branch, foreach), sigil system, type compatibility, and task return wiring. |
| State transition edges | Reference | `spec:bpdsl.dsl.edges.state_transitions` | `transitions:` section syntax, transition fields, action placement, and guard semantics. |
| Cross-edges | Reference | `spec:bpdsl.dsl.edges.cross_edges` | `reads:` / `writes:` cross-layer task fields and deprecated edge kinds. |
| Cross-file references | Reference | `spec:bpdsl.dsl.edges.cross_file_refs` | QualifiedID syntax for module-crossing node references. |

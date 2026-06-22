# Overview: Node definitions

- **id**: `spec:bpdsl.dsl.nodes.overview`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.overview`

## What this is

Entry point for BPDSL node kind specs. Covers the shared file structure, common fields, and the diagram-presence boundary matrix for all ten node kinds.

## Current contract

### File structure

One file = one main node (`main: true`) + zero or more sub-nodes (V01-ADR-011).

```yaml
nodes:
  - id: <string>
    type: <node-type>
    main: true          # main node only; one per file

  - id: <sub-node>      # sub-node (file-private)
    type: <node-type>

flow:                   # Processing layer wiring
  - step: ...

transitions:            # Application layer state transitions
  - from: ...
```

Sub-nodes are file-private and cannot be referenced from outside the file. Only the main node's ID forms a public QualifiedID subject to project-wide uniqueness.

Bare IDs in `flow:`, `reads:`, `writes:` are resolved first against sub-nodes in the same file, then against main nodes in the same module.

### Common fields

Fields shared across all node kinds:

| field | required | type | description |
|---|---|---|---|
| `id` | ✓ | string | Public ID for main nodes; file-local ID for sub-nodes. |
| `type` | ✓ | enum | Node kind (see table below). |
| `note` | optional | string | Human docstring and LLM semantic contract (V01-ADR-008). |

### Node kind boundary matrix

| kind | layer | file placement | Seq | State | DAG | ER | API |
|---|---|---|:---:|:---:|:---:|:---:|:---:|
| `task` | Processing | `task/*.yaml` | △※1 | ❌ | ✅ | ❌ | △※1 |
| `model` | Data | `model/*.yaml` | ❌ | ❌ | ❌ | ✅ | ❌ |
| `asset` | Processing | none (derived from task.returns) | ❌ | ❌ | ✅ | ❌ | ❌ |
| `store` | Processing / Data | `store/*.yaml` | ❌ | ❌ | ✅ | ✅ | ❌ |
| `actor` | Application | project-global; any filename | ✅ | ❌ | ❌ | ❌ | ❌ |
| `event` | Application | co-located in state.yaml etc. | ✅ | ✅ | ❌ | ❌ | ❌ |
| `state` | Application | `state.yaml` | ❌ | ✅ | ❌ | ❌ | ❌ |
| `branch` | Processing | sub-node in task/*.yaml | ❌ | ❌ | ✅ | ❌ | ❌ |
| `fork` | Processing | sub-node in task/*.yaml | ❌ | ❌ | ✅ | ❌ | ❌ |
| `join` | Processing | sub-node in task/*.yaml | ❌ | ❌ | ✅ | ❌ | ❌ |

Legend: Seq = Sequence Diagram / State = State Diagram / ER = ER Diagram / API = API Table.
※1 `endpoint: true` only. In Sequence Diagram, task appears as an arrow label, not a lifeline (V01-ADR-017).

`foreach` was removed as a node kind (V01-ADR-016). It is a control flow construct written in the `flow:` section.

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Processing nodes | Reference | `spec:bpdsl.dsl.nodes.processing` | Field definitions for task, asset, branch, fork, and join. |
| Data nodes | Reference | `spec:bpdsl.dsl.nodes.data` | Field definitions for model and store. |
| Application nodes | Reference | `spec:bpdsl.dsl.nodes.application` | Field definitions for actor, event, and state. |

# Reference: Application nodes

- **id**: `spec:bpdsl.dsl.nodes.application`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.nodes.overview`

## What this is

Field definitions for the three Application-layer node kinds: `actor`, `event`, and `state`. These nodes appear in the Sequence Diagram and State Diagram views. They do not appear in the DAG or ER Diagram.

## actor

> Source: V01-ADR-004, V01-ADR-006, V01-ADR-031

Application layer. Human or external system. Appears as a participant column in the Sequence Diagram. Corresponds to UML Actor.

`actor` is project-global — it does not belong to any module. It can be defined in any file with any filename. The parser enforces project-wide ID uniqueness (V01-ADR-031).

```yaml
# actors.yaml (filename is arbitrary)
nodes:
  - id: stripe
    type: actor
    note: "External payment service"

  - id: scheduler
    type: actor
    note: "Cron scheduler"

  - id: end_user
    type: actor
    note: "End user of the service"
```

### actor fields

| field | required | type | description | source |
|---|---|---|---|---|
| (no kind-specific fields) | — | — | Expressed with `id` / `type` / `note` only. No module path needed. References always use direct ID. | V01-ADR-004, V01-ADR-021, V01-ADR-031 |

---

## event

> Source: V01-ADR-017, V01-ADR-018

Application layer. Control flow trigger. Does **not** appear in the DAG `flow:`. Used in Sequence Diagram and State Diagram.

```yaml
- id: login_submitted
  type: event
  source: ui
  payload:
    model: login_form
  note: "Login form submit"

- id: payment_webhook_received
  type: event
  source: external
  actor: stripe
  payload:
    model: payment_event
  note: "Payment completion notification from Stripe"

- id: connection_status_changed
  type: event
  source: er
  watches: db_connection_store
  payload:
    model: connection_status
  note: "Fires when db_connection_store status changes"
```

### event fields

| field | required | type | description | source |
|---|---|---|---|---|
| `source` | ✓ | enum | `ui` / `external` / `er` / `internal` | V01-ADR-018, V01-ADR-034 |
| `actor` | required for `external` | actor-id | Originating actor ID. Must be declared as `type: actor` somewhere in the project. | V01-ADR-018 |
| `payload` | optional | payload | Model reference for the data this event carries. | V01-ADR-018 |
| `watches` | required for `er` | store-id | Store ID to watch for changes. | V01-ADR-018 |

### payload object

| field | required | type | description |
|---|---|---|---|
| `model` | ✓ | model-id | Payload type reference. Treated as model-id (not TypeRef) in v1.1. |

### source semantics

| source | meaning | `actor` | `payload` |
|---|---|---|---|
| `ui` | User interaction (click, form submit, etc.). Implicitly generates a UI participant column. | not needed | optional (form data etc.) |
| `external` | Input from an external system or scheduler. `actor:` names the origin. | **required** | optional (received data) |
| `er` | Fires when a store value changes. `watches:` required. | not needed | optional |
| `internal` | Fires internally within the app (task completion, FSM runtime, internal timer, etc.). What is watched is described in `note:` (V01-ADR-034). | not needed | optional |

---

## state

> Source: V01-ADR-017, V01-ADR-019

Application layer. FSM state node. Used in State Diagram. Distinct from `store` (data holder) — see distinction table below (V01-ADR-019).

```yaml
- id: idle
  type: state
  initial: true
  note: "User is not interacting"

- id: loading
  type: state
  note: "Authentication request in progress"

- id: authenticated
  type: state
  final: true

- id: error
  type: state
  final: true
  note: "Authentication error state"
```

### state fields

| field | required | type | description | source |
|---|---|---|---|---|
| `initial` | optional | bool | `true` = FSM initial state. One per file. | V01-ADR-019 |
| `final` | optional | bool | `true` = terminal state. Multiple allowed. | V01-ADR-019 |
| `note` | optional | string | State meaning and the condition for being in this state. | V01-ADR-019 |
| `wireframe` | optional | object | UI skeleton for this state. Single root container node. Details defined in the wireframe view spec. | V01-ADR-029, V01-ADR-042 |

`wireframe` field details: see [`spec:bpdsl.views.wireframe`](../../views/wireframe.md). Design decisions for `main` container and `layout` object: V01-ADR-042.

### state vs store distinction

| node | layer | concept |
|---|---|---|
| `store` | Processing / Data | Runtime data holder. Target of DAG reads / writes. |
| `state` | Application | FSM state definition. Used as source / target of transitions. |

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.dsl.nodes.overview` | Parent overview; node kind boundary matrix. |
| `spec:bpdsl.dsl.edges.state_transitions` | State transition syntax (`transitions:` section). |
| `spec:bpdsl.views.state_diagram` | State Diagram render contract using `state` and `event` nodes. |
| `spec:bpdsl.views.sequence_diagram` | Sequence Diagram render contract using `actor` and `event` nodes. |

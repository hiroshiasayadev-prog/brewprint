# Reference: State transition edges

- **id**: `spec:bpdsl.dsl.edges.state_transitions`
- **status**: draft
- **date**: 2026-06-16
- **parent**: `spec:bpdsl.dsl.edges.overview`

## What this is

Syntax and field definitions for the `transitions:` section, which declares Application-layer FSM state machine transitions. `flow:` is reserved for Processing-layer DAG wiring; Application-layer transitions use `transitions:` to keep the two orthogonal (V01-ADR-019).

## Transition entries

```yaml
nodes:
  - id: idle
    type: state
    initial: true

  - id: loading
    type: state

  - id: login_submitted
    type: event
    source: ui
    payload:
      model: login_form

transitions:
  - from: idle
    on: login_submitted
    to: loading
    action: auth.task.login
    note: "Call the login API"

  - from: loading
    on: login_succeeded
    to: authenticated

  - from: loading
    on: login_failed
    to: error
    guard: "retryCount < 3"
    note: "Only if retry limit is not reached"
```

### Transition entry fields

| field | required | type | description | source |
|---|---|---|---|---|
| `from` | ✓ | state-id | Source state ID. | V01-ADR-019 |
| `on` | ✓ | event-id | Triggering event ID. | V01-ADR-019 |
| `to` | ✓ | state-id | Target state ID. | V01-ADR-019 |
| `action` | optional | task-id | Task to execute during this transition (full path or local ID). | V01-ADR-019 |
| `guard` | optional | string | Transition condition text. Evaluation is outside brewprint scope. | V01-ADR-019 |
| `note` | optional | string | Supplemental description. | V01-ADR-019 |

## Action placement rule

`action` belongs on the `transition`, not on the `event`. Rationale: an action is determined by the `(state, event)` pair. The same event from different source states may trigger different actions (V01-ADR-019, Mealy machine pattern).

```yaml
transitions:
  - from: idle
    on: login_submitted
    to: loading
    action: auth.task.login      # standard login

  - from: session_expired
    on: login_submitted
    to: loading
    action: auth.task.reauth     # re-authentication (different task)
```

## Guard semantics

`guard` evaluation is outside brewprint scope (implementation-language dependent). Complex conditions should be expressed as an extract task or as task logic. In YAML, only a natural language text description is written (V01-ADR-019).

## Event co-location

Events used exclusively in a State Diagram are typically co-located with `state` nodes in the same file's `nodes:` section. Events shared with other files (e.g., a Sequence Diagram) are referenced via cross-file reference (V01-ADR-003, V01-ADR-019).

## Related specs

| ref | relation |
|---|---|
| `spec:bpdsl.dsl.edges.overview` | Parent overview; edge kind summary. |
| `spec:bpdsl.dsl.nodes.application` | `state` and `event` node definitions. |
| `spec:bpdsl.dsl.edges.cross_file_refs` | Cross-file reference syntax for shared event IDs. |
| `spec:bpdsl.views.state_diagram` | State Diagram render contract. |

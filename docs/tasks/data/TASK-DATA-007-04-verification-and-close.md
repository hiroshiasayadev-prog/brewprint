# TASK-DATA-007-04: Verification and close

- **id**: TASK-DATA-007-04
- **status**: done
- **date**: 2026-06-02
- **work_item**: WORK-DATA-007
- **source_requirement**: REQ-DATA-005
- **estimate**: 0.5d
- **depends_on**:
  - TASK-DATA-007-03
- **outputs**:
  - WORK-DATA-007 close synchronization
  - REQ-DATA-005 acceptance evidence
  - Final validation evidence

## Goal

Close WORK-DATA-007 after ADR-074 acceptance, DAG view spec alignment, renderer implementation, fixture / golden updates, and verification evidence are complete.

## Work

- Confirm TASK-DATA-007-01 through TASK-DATA-007-03 are done.
- Confirm ADR-074 is accepted and aligned to REQ-DATA-005 / WORK-DATA-007.
- Confirm DAG view spec owns the accepted asset TypeRef hint render rule.
- Confirm renderer implementation and UC-001 golden evidence were completed by TASK-DATA-007-03.
- Confirm tests and UC validation passed.
- Mark WORK-DATA-007 as done and record close evidence.
- Mark REQ-DATA-005 as accepted and record close evidence.
- Validate task / work item / requirement / ADR metadata.

## Done Condition

- WORK-DATA-007 status is `done`.
- REQ-DATA-005 status is `accepted`.
- Close evidence records the implementation result and verification commands.
- Design Records MCP validation passes for TASK-DATA-007-04, WORK-DATA-007, REQ-DATA-005, and ADR-074.
- No implementation, fixture, golden, UC YAML, TypeRef compatibility, or diagnostic changes are performed by this close task.

## Verification

- Design Records MCP validation for TASK-DATA-007-03 passed before close.
- Design Records MCP validation for WORK-DATA-007 passed before close.
- TASK-DATA-007-03 implementation evidence reports `go test ./...` PASS.
- TASK-DATA-007-03 implementation evidence reports UC-001 / UC-002 validate cleanly with 0 errors and 0 warnings.
- Final Design Records MCP validation is recorded below.

## Evidence

Closed on 2026-06-02.

### Inputs confirmed

- TASK-DATA-007-01: done; ADR-074 acceptance / split review completed with `revise-before-acceptance` input.
- TASK-DATA-007-02: done; ADR-074 revised to accepted, collision inventory completed, and DAG view spec aligned.
- TASK-DATA-007-03: done; renderer and fixture update completed.
- ADR-074: accepted; active boundary is REQ-DATA-005 / WORK-DATA-007.

### Implementation result confirmed from TASK-DATA-007-03

- `internal/render/dag/type_hint.go` added TypeRef hint calculation logic.
- `internal/render/dag/flow_renderer.go` updated for `ambiguousHints`, `calcHint`, params boundary, and asset references with TypeRef hints.
- `internal/render/dag/renderer.go` updated for boundary node hint support and simple renderer hints.
- `internal/render/dag/type_hint_test.go` added tests for asset hints, named model local IDs, and collision detection.
- `internal/render/dag/private_models_test.go` updated for hint-aware assertions.
- UC-001 DAG golden files were updated for TypeRef hint labels.

### Behavior confirmed

- Params boundary assets render hints, e.g. `form([form: login_form])`, `cart_id([cart_id: str])`, `cart_items([cart_items: cart_item_list])`.
- Task returns render hints, e.g. `auth_token([auth_token: token])`, `draft_order([draft_order: order])`.
- Join returns render hints, e.g. `pending_order([pending_order: order])`.
- Foreach collected asset renders top-level inline list hint, e.g. `validated_items([validated_items: list])`.
- Named list model remains named, e.g. `cart_item_list`, not `list`.
- Full QID TypeRef renders as local id, e.g. `payment.model.payment_event` -> `payment_event`.
- Ambiguous local ID hints are omitted.
- Shortened QID fallback was not implemented.
- No new diagnostics were added.
- TypeRef compatibility was not changed.
- UC YAML was not changed.

### Test / validation evidence confirmed

- `go test ./...` passed.
- UC-001 validate passed with 0 errors / 0 warnings.
- UC-002 validate passed with 0 errors / 0 warnings.
- Design Records MCP validation for TASK-DATA-007-03 passed.
- Design Records MCP validation for WORK-DATA-007 passed.

### Final close status

WORK-DATA-007 is closed as done.
REQ-DATA-005 is accepted.

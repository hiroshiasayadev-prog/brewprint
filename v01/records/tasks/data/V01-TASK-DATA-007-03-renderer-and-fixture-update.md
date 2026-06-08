# V01-TASK-DATA-007-03: Renderer and fixture update

- **id**: V01-TASK-DATA-007-03
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-DATA-007
- **source_requirement**: V01-REQ-DATA-005
- **estimate**: 1d-2d
- **depends_on**:
  - V01-TASK-DATA-007-02
- **outputs**:
  - DAG renderer TypeRef hint implementation
  - Updated DAG fixture / golden evidence for asset TypeRef hints
  - Tests covering accepted V01-ADR-074 / dag.md minimum behavior

## Goal

Implement the accepted V01-ADR-074 / dag.md minimum for DAG asset TypeRef hint labels and update the relevant fixture / golden evidence.

This task follows V01-TASK-DATA-007-02, where V01-ADR-074 was accepted and the active DAG view spec was aligned. The implementation must follow that accepted minimum without reintroducing shortened QID fallback into V01-WORK-DATA-007.

## Work

- Review V01-ADR-074 and `docs/spec/views/dag.md` for the accepted asset TypeRef hint rule.
- Inspect the current DAG renderer implementation and tests.
- Implement asset label TypeRef hint rendering for:
  - main task params boundary assets,
  - task returns assets,
  - join returns assets,
  - foreach collected assets.
- Implement top-level `type_hint` calculation:
  - primitive TypeRef shows primitive name,
  - named model TypeRef shows model local id,
  - inline `list<T>` shows `list`,
  - inline `dict<T>` shows `dict`,
  - named list / dict model remains named model hint.
- Ensure full container TypeRefs are not expanded in Mermaid labels.
- Ensure invalid / unresolved TypeRef cases do not add a DAG-render diagnostic surface; when a hint cannot be safely calculated, omit the hint and keep the asset name.
- Ensure ambiguous named model local-id collision behavior follows the accepted minimum: omit the ambiguous TypeRef hint rather than rendering local id, full QID, shortened QID, or suffix-qualified id.
- Update representative fixtures / golden outputs for UC-001 / UC-002 or existing DAG renderer fixtures so the new asset labels are covered.
- Add or update tests for params boundary asset hints, task returns hints, inline list / dict hints, named list / dict model hints, and hint omission where feasible.
- Keep TypeRef syntax, TypeRef compatibility, diagnostics, ADR text, spec text, and UC YAML changes outside this task unless a narrow test fixture requires generated output refresh.

## Included Scope

- DAG renderer behavior for accepted TypeRef hint labels.
- Fixture / golden updates directly caused by the accepted label change.
- Unit / integration tests covering renderer output and representative UC evidence.
- Validation that shortened QID fallback is not implemented as part of the V01-WORK-DATA-007 minimum.

## Excluded Scope

- Revising V01-ADR-074 or DAG spec wording.
- Implementing shortened QID fallback / suffix-qualified id disambiguation.
- Changing TypeRef syntax, TypeRef compatibility, or resolver semantics.
- Adding new diagnostics for DAG render hints.
- UC YAML migration unrelated to fixture / golden refresh.
- V01-ADR-073 tagged union / discriminator payload support.
- V01-ADR-078 / V01-ADR-079 / V01-ADR-080 MCP semantic identity / state machine identity.
- UC-002 duplicate task QID / unresolved flow task repair.
- Remaining UC-002 notes retreat cleanup.
- Reopening M15, V01-WORK-DATA-001, V01-WORK-DATA-002, V01-WORK-DATA-003, or V01-WORK-DATA-004.

## Done Condition

- DAG renderer emits accepted top-level TypeRef hints on supported asset labels.
- Params boundary assets and ordinary returned assets are covered.
- Full TypeRef container details remain out of Mermaid labels.
- Ambiguous / invalid / unresolved hint cases omit the hint rather than rendering misleading or expanded identity.
- Fixture / golden evidence is updated for representative cases.
- Relevant tests pass.
- No excluded scope is changed.

## Verification

- Run the smallest relevant DAG renderer tests.
- Run any affected resolve / query tests if renderer implementation depends on resolved TypeRef metadata.
- Run UC-001 / UC-002 render or validation commands needed to confirm fixture / golden evidence.
- Validate `V01-TASK-DATA-007-03` metadata with Design Records MCP.
- Validate `V01-WORK-DATA-007` metadata / relations with Design Records MCP.
- Record exact commands and results in Evidence before marking this task done.

## Evidence
Completed on 2026-06-03.

### Commands run

```
git status --short
go test ./internal/resolve ./internal/render/dag ./internal/query
go run ./cmd/brewprint validate --yaml-root docs/uc/001-ec-checkout-flow/yaml --format json
go run ./cmd/brewprint validate --yaml-root docs/uc/002-brewprint-self-hosting/yaml --format json
go test ./internal/render/dag/... -v
go test ./...
```

### Test results

- `go test ./internal/render/dag/...` — PASS (all subtests)
- `go test ./...` — PASS (all packages, no FAIL)
- UC-001 validate: `errors=0 warnings=0`
- UC-002 validate: `errors=0 warnings=0`

### Files changed

Implementation:

- `internal/render/dag/type_hint.go` (new): `assetTypeHint`, `namedModelLocalID`, `calcAssetHint`, `collectFlowAssetTypeRefs`, `computeAmbiguousHints`
- `internal/render/dag/flow_renderer.go`: added `ambiguousHints` field, `calcHint` method, `writeParamsBoundary` method; updated `assetRef` to accept TypeRef; updated `writeTaskReturn`, `renderForeach`, `renderFork` callers; removed pre-registration of `main.Returns.Name` so join/foreach returns are rendered with correct stadium shape
- `internal/render/dag/renderer.go`: added `boundaryNode` struct; updated `writeBoundary` to support per-node hints; updated simple renderer to compute `ambiguous` hints and pass them to boundary and returns rendering; added `paramBoundaryNodes` helper

Tests:

- `internal/render/dag/type_hint_test.go` (new): unit tests for `assetTypeHint`, `namedModelLocalID`, `calcAssetHint` ambiguity, `computeAmbiguousHints`
- `internal/render/dag/private_models_test.go`: updated assertions — private model local IDs in type hint labels are correct V01-ADR-074 behavior; `credentials` (nested helper, not a direct TypeRef) still asserted absent from Mermaid body

Golden files updated (all 7 UC-001 DAG goldens):

- `docs/uc/001-ec-checkout-flow/renders/auth/dag-login.md` — `form([form: login_form])`, `auth_token([auth_token: token])`
- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-add_to_cart.md` — `cart_id([cart_id: str])`, `item_id([item_id: str])`, `qty([qty: int])`, `updated_cart([updated_cart: cart])`
- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-validate_cart.md` — `cart_items([cart_items: cart_item_list])`, `validated_items([validated_items: list])` (foreach collected, inline list → "list")
- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-checkout.md` — `cart_id([cart_id: str])`, `shipping_address([shipping_address: address])`, `draft_order([draft_order: order])`, `reserved([reserved: order])`, `notified([notified: order])`, `pending_order([pending_order: order])` (join returns now rendered with stadium shape)
- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_order.md` — `order_id([order_id: str])`, `order_asset([order: order])`
- `docs/uc/001-ec-checkout-flow/renders/commerce/dag-process_payment.md` — `event([event: payment_event])` (full QID `payment.model.payment_event` → local ID `payment_event`)
- `docs/uc/001-ec-checkout-flow/renders/catalog/dag-get_items.md` — `items([items: item_list])`

### Coverage of accepted V01-ADR-074 minimum

| requirement | covered | evidence |
|---|---|---|
| params boundary asset hint | ✓ | `form([form: login_form])`, `cart_id([cart_id: str])`, `cart_items([cart_items: cart_item_list])` |
| task returns asset hint | ✓ | `auth_token([auth_token: token])`, `draft_order([draft_order: order])` |
| join returns asset hint | ✓ | `pending_order([pending_order: order])` |
| foreach collected asset hint | ✓ | `validated_items([validated_items: list])` |
| primitive hint (str / int / any) | ✓ | `cart_id([cart_id: str])`, `qty([qty: int])` |
| inline list hint | ✓ | `validated_items([validated_items: list])` |
| named list model stays local id | ✓ | `cart_items([cart_items: cart_item_list])` (not "list") |
| full container TypeRef not in label | ✓ | `list<cart_item>` never appears in Mermaid labels |
| ambiguous hint omission | ✓ | unit tests `TestCalcAssetHint_Ambiguity`, `TestComputeAmbiguousHints` |
| invalid / unresolved hint omission | ✓ | `calcAssetHint(nil, ...)` returns "" |
| shortened QID fallback not implemented | ✓ | no shortened QID logic added |

### Scope checks

- shortened QID fallback implemented: no
- new diagnostics added: no
- TypeRef compatibility changed: no
- UC YAML changed: no
- unrelated work touched: no
- V01-ADR-074 revised: no
- spec revised: no

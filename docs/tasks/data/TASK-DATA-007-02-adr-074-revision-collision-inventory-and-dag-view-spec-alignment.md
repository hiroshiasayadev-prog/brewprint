# TASK-DATA-007-02: ADR-074 revision, collision inventory, and DAG view spec alignment

- **id**: TASK-DATA-007-02
- **status**: done
- **date**: 2026-06-02
- **work_item**: WORK-DATA-007
- **source_requirement**: REQ-DATA-005
- **estimate**: 1d-2d
- **depends_on**:
  - TASK-DATA-007-01
- **outputs**:
  - UC-001 / UC-002 named model local-id collision inventory
  - ADR-074 revision / split decision for collision disambiguation
  - Accepted DAG TypeRef hint minimum contract
  - DAG view spec alignment for asset TypeRef hint labels

## Goal

Revise ADR-074 toward the accepted WORK-DATA-007 minimum, decide whether shortened QID fallback belongs in the minimum or should be split / deferred, and align the DAG view spec with the accepted TypeRef hint render contract.

This task exists because TASK-DATA-007-01 concluded `revise-before-acceptance`, and the Opus review confirmed that the next task must close the shortened-QID fallback ambiguity before implementation or fixture work starts.

## Work

- Review TASK-DATA-007-01 evidence and the Opus review findings as input.
- Inventory UC-001 / UC-002 current YAML for named model local-id collisions within the same DAG render scope.
- Use the collision inventory to decide whether ADR-074 should be:
  - partially revised with collision behavior kept inside ADR-074,
  - revised with shortened QID fallback deferred as non-minimum behavior,
  - or split so collision disambiguation becomes a separate proposed ADR / follow-up decision.
- If shortened QID fallback is deferred from the minimum, explicitly define the minimum collision-time behavior in ADR-074 and the DAG spec. Candidate minimum behavior is to omit the ambiguous type hint rather than render misleading local-id hints or full QIDs in Mermaid labels.
- Revise ADR-074 before acceptance so that:
  - the active boundary is REQ-DATA-005 / WORK-DATA-007,
  - M15 / WORK-DATA-001 remain closed and are not reopened,
  - top-level TypeRef hint rules remain the core accepted behavior,
  - params boundary assets remain included,
  - `subgraph returns` remains excluded,
  - full TypeRef remains in Markdown detail sections,
  - Mermaid labels do not display full container TypeRefs,
  - unresolved / invalid TypeRef handling stays with existing diagnostics and hint omission,
  - shortened QID fallback is either accepted, explicitly deferred, or split.
- Update ADR-074 Evidence / close-boundary wording so it no longer reads as only an M15 deferred item.
- Update `docs/spec/views/dag.md` to define the accepted minimum label format and TypeRef hint calculation rule.
- Update stale DAG spec boundary wording that still points to `WORK-DATA-003` as the future scope for ADR-074, replacing it with REQ-DATA-005 / WORK-DATA-007.
- Keep TypeRef syntax, TypeRef compatibility, diagnostics, renderer implementation, fixtures, golden output, and UC YAML changes outside this task unless a narrow wording-only spec reference is required.

## Decision Questions

1. Do UC-001 / UC-002 currently contain same-DAG render scope named model local-id collisions?
2. If no collision is observed, is it sufficient to defer shortened QID fallback from the WORK-DATA-007 minimum?
3. If collision is observed, should the minimum omit ambiguous hints, accept shortened QID fallback, or split collision disambiguation into a separate ADR?
4. Should ADR-074 be partially revised, or should collision disambiguation be extracted into a separate proposed ADR?
5. What exact collision-time behavior should the accepted DAG TypeRef hint minimum specify?
6. Which stale ADR-074 / dag.md boundary lines must be updated so future readers do not treat ADR-074 as a still-deferred M15 item?

## Included Scope

- Collision inventory for current UC-001 / UC-002 YAML and DAG render scopes.
- ADR-074 revision or split decision before acceptance.
- ADR-074 wording update for active REQ-DATA-005 / WORK-DATA-007 boundary.
- DAG view spec alignment for accepted asset TypeRef hint label behavior.
- Stale boundary reference cleanup directly tied to ADR-074 / WORK-DATA-007.

## Excluded Scope

- Renderer implementation changes.
- Fixture / golden output regeneration.
- UC-001 / UC-002 YAML migration or behavior changes.
- TypeRef syntax or compatibility changes.
- New diagnostics for DAG render hints.
- ADR-073 tagged union / discriminator payload support.
- ADR-078 / ADR-079 / ADR-080 MCP semantic identity / state machine identity.
- UC-002 duplicate task QID / unresolved flow task repair.
- Remaining UC-002 notes retreat cleanup.
- Reopening M15, WORK-DATA-001, WORK-DATA-002, WORK-DATA-003, or WORK-DATA-004.

## Done Condition

- UC-001 / UC-002 collision inventory result is recorded.
- ADR-074 is either revised toward acceptance or collision disambiguation is explicitly split / deferred with a documented rationale.
- If shortened QID fallback is not in the minimum, the accepted minimum still defines deterministic collision-time behavior.
- `docs/spec/views/dag.md` defines the accepted DAG asset TypeRef hint label rule.
- Stale `WORK-DATA-003` / M15-only boundary wording directly related to ADR-074 is updated or explicitly left with rationale.
- No renderer, fixture, golden, UC YAML, TypeRef compatibility, or diagnostic implementation changes are performed by this task.

## Verification

- Validate `TASK-DATA-007-02` metadata with Design Records MCP.
- Validate `WORK-DATA-007` metadata / relations with Design Records MCP.
- Confirm ADR-074 status and wording reflect the chosen revise / split path.
- Confirm DAG view spec includes the accepted minimum label rule and no unrelated render behavior.
- Confirm excluded scopes remain untouched.
- If command execution is available, run the smallest relevant grep / validation commands needed to support the collision inventory; otherwise record that collision inventory requires Codex execution.

## Evidence

Partial progress on 2026-06-02.

### Collision inventory input

Inspected current UC-001 / UC-002 YAML using filesystem reads only. No renderer, resolver, validation, fixture, golden, or UC YAML changes were performed.

UC-001 task DAGs inspected:

| DAG file | asset TypeRef sources inspected | named model local-id collision result |
|---|---|---|
| `auth/task/login.yaml` | `login_form`, `token` | none observed |
| `cart/task/add_to_cart.yaml` | `cart` plus primitive params | none observed |
| `cart/task/validate_cart.yaml` | `cart_item_list`, `cart_item` | none observed |
| `catalog/task/get_items.yaml` | `item_list` | none observed |
| `order/task/checkout.yaml` | `address`, `order` | none observed |
| `order/task/process_order.yaml` | `order` plus primitive params | none observed |
| `payment/webhooks/task/process_payment.yaml` | `payment.model.payment_event` | none observed |

UC-002 task DAGs inspected:

| DAG file | asset TypeRef sources inspected | named model local-id collision result |
|---|---|---|
| `mcp/task/analyze_impact.yaml` | `analyze_impact_request`, `analyze_impact_response`, `any` | none observed |
| `mcp/task/get_references.yaml` | `get_references_request`, `get_references_response`, `any` | none observed |
| `mcp/task/get_reference_tree.yaml` | `get_reference_tree_request`, `get_reference_tree_response`, `any` | none observed |
| `mcp/task/get_signature.yaml` | `get_signature_request`, `get_signature_response`, `any` | none observed |
| `mcp/task/get_source.yaml` | `get_source_request`, `get_source_response`, `any` | none observed |
| `mcp/task/inspect.yaml` | `inspect_request`, `inspect_response`, `any` | none observed |
| `mcp/task/list_endpoints.yaml` | `list_endpoints_request`, `list_endpoints_response`, `any` | none observed |
| `mcp/task/list_objects.yaml` | `list_objects_request`, `list_objects_response`, `any` | none observed |

Model definition placement checked at inventory level:

- UC-001 model local IDs are distributed across `auth`, `cart`, `catalog`, `order`, and `payment` modules. The task DAGs inspected above do not reference two distinct named models with the same local id in a single DAG render scope.
- UC-002 model files are under the single `mcp` module. The request / response model local IDs referenced by the task DAGs are distinct per DAG. `mcp/model/common.yaml` defines shared helper models such as `mcp_object_type`, `mcp_diagnostic_severity`, `reference_tree_direction`, and `object_selector`; none create a same-DAG duplicate local-id asset hint in the inspected task params / returns.

### Collision inventory result

Result: `no current UC-001 / UC-002 same-DAG named model local-id collision observed`.

This supports deferring `shortened QID fallback` from the WORK-DATA-007 minimum, as long as the accepted minimum still defines deterministic collision-time behavior. The recommended minimum behavior remains: if a named model local id is ambiguous in a DAG render scope and shortened-QID fallback is not part of the minimum, omit the type hint for the ambiguous asset label rather than rendering a misleading local id or a full QID in Mermaid.

### Codex verification

Codex performed repo-local verification for the same collision inventory and reported `PASS_NO_COLLISION`.

Commands run:

- `git status --short`
- `go test ./internal/resolve ./internal/render/dag ./internal/query`
- `go run ./cmd/brewprint validate --yaml-root docs/uc/001-ec-checkout-flow/yaml --format json`
- `go run ./cmd/brewprint validate --yaml-root docs/uc/002-brewprint-self-hosting/yaml --format json`
- `go test -overlay <temp overlay outside repo> -run TestCollisionInventory00702 -count=1 -v ./internal/query`
  - First overlay attempt failed due relative path resolution from package cwd.
  - Rerun with `BREWPRINT_REPO_ROOT` passed.
  - Temporary files were outside the repo and removed.

Codex result summary:

- Current UC-001 / UC-002 do not require shortened QID fallback in the WORK-DATA-007 minimum.
- Both UC roots validate cleanly with 0 errors and 0 warnings.
- Existing CLI / query / MCP does not expose a direct collision inventory command, so Codex used a temporary Go overlay test against `source.Loader` and `resolve.Build`.
- Current renderer does not implement ADR-074 hints yet; the inventory checked candidate hint surfaces from resolved semantic structures.
- Scope was limited to the listed UC-001 / UC-002 task DAG files, with an intentionally over-inclusive params / returns / join / foreach asset scope.

### Collision inventory decision input

The filesystem inspection and Codex repo-local verification agree: no same-DAG named model local-id collision is currently observed in UC-001 / UC-002.

This is sufficient evidence to proceed with ADR-074 revision using the following WORK-DATA-007 minimum:

- `shortened QID fallback` is deferred from the minimum.
- The minimum still defines deterministic collision-time behavior.
- If a named model local id is ambiguous in a DAG render scope, omit the type hint for the ambiguous asset label rather than rendering a misleading local id or a full QID in Mermaid.

### Inventory limitation

The inventory does not prove future projects or future fixtures cannot create same-DAG named model local-id collisions. It only proves that current UC-001 / UC-002 do not require shortened QID fallback to validate the WORK-DATA-007 minimum. Collision disambiguation can be split into follow-up behavior if a future fixture introduces a real same-DAG collision.

### ADR-074 revision result

ADR-074 was revised from `proposed` to `accepted` under the active REQ-DATA-005 / WORK-DATA-007 boundary.

Revision summary:

- The top note now states that ADR-074 is accepted for REQ-DATA-005 / WORK-DATA-007.
- M15 / WORK-DATA-001 are explicitly not reopened.
- The core top-level hint rules remain accepted:
  - primitive hints show the primitive name,
  - named model hints show the model local id,
  - inline `list<T>` and `dict<T>` collapse to `list` / `dict`,
  - named list / dict models remain named model hints,
  - params boundary assets are included,
  - `subgraph returns` remains excluded,
  - full TypeRef remains in Markdown detail sections.
- `shortened QID fallback` / suffix-qualified id calculation was removed from the accepted minimum.
- The deterministic collision-time behavior is now: if another named model identity in the same DAG render scope would produce the same local id hint, omit that ambiguous TypeRef hint and render the asset name only.
- ADR-074 Evidence now records the TASK-DATA-007-02 / Codex `PASS_NO_COLLISION` inventory and the WORK-DATA-007 minimum collision behavior.

Decision: partial revise, not split.

Rationale: current UC-001 / UC-002 have no same-DAG collision, so a separate collision-disambiguation ADR is not required for the WORK-DATA-007 minimum. The accepted ADR keeps the readability hint contract complete by defining hint omission for ambiguous local ids. A future fixture with real collision pressure can create a follow-up ADR / task for shortened QID disambiguation.

### DAG view spec alignment result

`docs/spec/views/dag.md` was updated to align with the accepted ADR-074 minimum.

Spec alignment summary:

- The `asset` node render rule now uses `asset_name([asset_name: type_hint])`.
- The spec defines top-level hint calculation for primitive, named model, inline `list<T>`, and inline `dict<T>` TypeRefs.
- Named list / dict models are specified as named model hints, not collapsed to `list` / `dict`.
- Ambiguous named model local-id hints are omitted instead of using shortened QID fallback.
- Invalid / unresolved TypeRef hints are omitted, with diagnostics remaining the owner of invalid / unresolved explanation.
- Full TypeRef / full identity are explicitly kept out of Mermaid labels and left to detail sections / MCP inspect / model render / catalog render.
- `subgraph params` example and bullets now show that params boundary assets are TypeRef hint targets.
- The stale `WORK-DATA-003` future-scope wording was replaced with ADR-074 / REQ-DATA-005 / WORK-DATA-007 ownership.
- The basic DAG Mermaid example was updated for `config: app_config`, `raw: raw_data`, and `result: result_data`.

### Decision question answers

1. Do UC-001 / UC-002 currently contain same-DAG render scope named model local-id collisions?
   - No. Filesystem inspection and Codex repo-local verification both reported no observed same-DAG named model local-id collision.

2. If no collision is observed, is it sufficient to defer shortened QID fallback from the WORK-DATA-007 minimum?
   - Yes, as long as deterministic collision-time behavior is still specified. ADR-074 and dag.md now specify hint omission for ambiguous named model local ids.

3. If collision is observed, should the minimum omit ambiguous hints, accept shortened QID fallback, or split collision disambiguation into a separate ADR?
   - No current collision is observed. The accepted minimum uses hint omission. If future fixtures introduce real collision pressure, shortened QID fallback should be handled as follow-up behavior rather than silently expanding the minimum.

4. Should ADR-074 be partially revised, or should collision disambiguation be extracted into a separate proposed ADR?
   - Partially revise ADR-074. Splitting is not needed for the current minimum because collision disambiguation is not accepted; the minimum only needs hint omission.

5. What exact collision-time behavior should the accepted DAG TypeRef hint minimum specify?
   - If a named model local id is ambiguous in the same DAG render scope, omit the TypeRef hint for the ambiguous asset label and render only the asset name.

6. Which stale ADR-074 / dag.md boundary lines must be updated so future readers do not treat ADR-074 as a still-deferred M15 item?
   - ADR-074 top note and Evidence close boundary were updated to name REQ-DATA-005 / WORK-DATA-007 and preserve M15 / WORK-DATA-001 closure. dag.md stale `WORK-DATA-003` wording was replaced with ADR-074 / REQ-DATA-005 / WORK-DATA-007 ownership.

### Verification result

- `TASK-DATA-007-02` performed collision inventory, ADR revision, and DAG view spec alignment only.
- ADR-074 is now `accepted` under REQ-DATA-005 / WORK-DATA-007.
- `shortened QID fallback` is not part of the WORK-DATA-007 minimum.
- Deterministic collision-time behavior is specified as hint omission for ambiguous named model local ids.
- `docs/spec/views/dag.md` contains the accepted asset TypeRef hint label rule.
- No renderer implementation, fixture / golden regeneration, UC YAML migration, TypeRef compatibility change, or new diagnostic was performed by this task.
- Excluded scopes remain excluded: ADR-073, ADR-078 / ADR-079 / ADR-080, UC-002 duplicate task QID repair, remaining notes retreat cleanup, and reopening M15 / WORK-DATA-001〜004.

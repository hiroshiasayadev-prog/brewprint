# Overview: PRODUCT specifications

- **id**: `spec:product`
- **status**: draft
- **date**: 2026-07-01
- **parent**: `root`

## What this is

Placement router for PRODUCT specifications.
It routes content by semantic ownership and does not restate child contracts.

## Current contract

| area | owns | must route elsewhere |
|---|---|---|
| `design-records/` | App-independent Design Records identity, authoring, format, record placement, traceability, and artifact responsibility semantics. | Brewprint registry facts, Brewprint compatibility state, DRMCP behavior, and canonical BPDSL internals. |
| `brewprint/` | Brewprint profile, current repository state, current namespace assignments, and Brewprint compatibility history. | Generic Design Records rules, DRMCP operational contracts, and canonical BPDSL internals. |
| `bpdsl/` | Temporary preservation of BPDSL-related material removed from mixed PRODUCT specifications. | Canonical BPDSL ownership, BPDSL redesign, new integration design, and unrelated new BPDSL specifications. |
| `responsibility-boundary-validator/` | Standalone semantic Task responsibility-boundary validation behavior, result semantics, outcome separation, and workflow-use boundary. | Generic Design Records rules, exact checklist artifacts, executable implementation, current DRMCP behavior, and future DRMCP integration. |

## Topics

| title | kind | ref | summary |
|---|---|---|---|
| Design Records | Overview | `spec:product.design_records` | App-independent Design Records semantics and ownership boundary. |
| Brewprint | Overview | `spec:product.brewprint` | Brewprint-specific profile, current state, namespaces, and compatibility. |
| Temporary BPDSL staging | Overview | `spec:product.bpdsl` | Non-canonical preservation area pending BPDSL migration review. |
| Responsibility-boundary validator | Overview | `spec:product.responsibility_boundary_validator` | Standalone semantic contract for Task responsibility-boundary validation. |

## Placement rules

| question | route |
|---|---|
| Does the content define app-independent Design Records semantics and remain valid when Brewprint assignments change? | Place it under `design-records/`. |
| Does the content record current Brewprint layout, assignments, or V01 compatibility? | Place it under `brewprint/`. |
| Is existing BPDSL-related content preserved only because final ownership is unresolved? | Stage it temporarily under `bpdsl/`. |
| Does the content define standalone semantic Task responsibility-boundary validation behavior or result semantics? | Place it under `responsibility-boundary-validator/`. |
| Does the content define DRMCP parser, storage, UI, tool, or authoring behavior? | Place it in DRMCP app-local specifications. Keep only a pointer in PRODUCT when needed. |
| Does the content define canonical BPDSL language, schema, resolver, render, generation, runtime, or MCP behavior? | Place it in BPDSL app-local specifications. Keep only a pointer in PRODUCT when needed. |
| Does the content describe an unadopted future integration mechanism? | Create an appropriate follow-up record or remove it after evidence transfer. Do not create a spec area. |

Physical location is not ownership evidence.
Classify content by the contract it defines before choosing a path.

## Dependency direction

| dependency | rule |
|---|---|
| Design Records to DRMCP | No normative dependency. |
| DRMCP to Design Records | Allowed. DRMCP may implement and expose Design Records contracts. |
| Design Records to BPDSL | No normative dependency. |
| BPDSL to Design Records | No normative integration dependency until an explicit requirement accepts one. |
| Brewprint to Design Records | Allowed. Brewprint may instantiate generic Design Records contracts. |
| Temporary PRODUCT BPDSL staging to canonical BPDSL | Contextual reference only. Staging creates no ownership claim. |
| Responsibility-boundary validator to Design Records | Allowed. The validator may consume PRODUCT-owned Task semantics. |
| Responsibility-boundary validator to DRMCP | No current normative dependency. Future integration requires separate authority. |

## Authoring boundary

- Generic contracts use app-neutral, non-normative examples where practical.
- Generic contracts must not copy current app or domain registry tables.
- Cross-owner contracts remain pointers instead of copied contract text.
- The validator area references Design Records Task semantics without moving validator ownership into `design-records/`.
- Child overviews own detailed rules for their areas.

## Related specs

| ref | relation |
|---|---|
| PRODUCT-ADR-SPEC-001 | Accepts the semantic ownership boundary and dependency direction. |

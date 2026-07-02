# PRODUCT-ADR-SPEC-001: PRODUCT spec semantic ownership boundary

- **status**: accepted
- **date**: 2026-07-01
- **depends_on**: []
- **supersedes**: []
- **migrated_to_spec**: 2026-07-01

## Context

The current `product/records/spec/` tree contains `concepts/` and `brewprint/`.

`concepts/` is not a semantic owner.
It mixes app-independent Design Records semantics with project profile, compatibility, BPDSL material, and app-local behavior.

`PRODUCT-INV-SPEC-005` classified all current PRODUCT spec files.
The investigation recommends explicit semantic areas and a placement router.

The earlier `PRODUCT-INV-SPEC-004` used a binary PRODUCT-versus-DRMCP model.
That model cannot classify the current mixed content safely.

This decision defines ownership boundaries only.
It does not move files, rewrite PRODUCT specs, redesign DRMCP, or perform BPDSL migration.

## Decision

### Top-level PRODUCT spec areas

The target top-level structure is:

```text
product/records/spec/
  index.md
  design-records/
  brewprint/
  bpdsl/
  responsibility-boundary-validator/
```

`concepts/` is not retained as a target semantic area.

No top-level `compatibility/` area is created.
No `deferred-integration/` spec area is created.

Exact child paths remain implementation-planning decisions, except `brewprint/compatibility/`, which this ADR fixes.

### Root placement router

`product/records/spec/index.md` is the placement router for PRODUCT specifications.
It does not restate child contracts in detail.

The root index defines:

| responsibility | required treatment |
|---|---|
| Top-level areas | Route content to `design-records/`, `brewprint/`, temporary `bpdsl/`, or the standalone `responsibility-boundary-validator/` area. |
| Ownership boundary | State each area's owned and prohibited content. |
| Dependency direction | Keep Design Records semantics independent from DRMCP and BPDSL. |
| Placement tests | Route by semantic ownership, not current physical location. |
| Examples | Prefer app-neutral examples in generic contracts. |
| Registry facts | Prohibit duplicate current app and domain registry tables in generic contracts. |
| Navigation | Link to each top-level area overview using canonical refs. |

### Area ownership

| area | owns | prohibited content |
|---|---|---|
| `design-records/` | App-independent Design Records identity, responsibilities, authoring, format, record placement, and traceability semantics. | DRMCP parser, storage, UI, tool, or authoring behavior; canonical BPDSL internals; Brewprint registry facts; Brewprint compatibility state. |
| `brewprint/` | Brewprint-specific profile, current repository state, current namespace assignments, and Brewprint compatibility history. | Generic Design Records rules; DRMCP operational contracts; canonical BPDSL language, render, resolver, generation, runtime, or MCP contracts. |
| `bpdsl/` | Temporary preservation of BPDSL-related material removed from mixed PRODUCT specifications. | Canonical BPDSL ownership claims; BPDSL redesign; new integration design; unrelated new BPDSL specifications. |
| `responsibility-boundary-validator/` | Standalone semantic Task responsibility-boundary validation behavior, result semantics, outcome separation, and workflow-use boundary. | Generic Design Records authoring rules; exact checklist artifacts; executable implementation; current DRMCP behavior; future DRMCP integration. |

Cross-owner references remain pointers.
A PRODUCT spec uses a pointer to an app-local contract instead of restating its content.

The standalone validator area may consume the PRODUCT Task responsibility contract.
Placement under PRODUCT does not assign validator implementation to Design Records or DRMCP.

### Brewprint compatibility placement

Brewprint V01 compatibility belongs under `product/records/spec/brewprint/compatibility/`.

This area owns:

- V01 compatibility policy;
- historical attribution;
- issued-ID retention;
- migration state.

Compatibility is project-specific history.
It is not an app-independent Design Records semantic area.

### Temporary BPDSL quarantine contract

`product/records/spec/bpdsl/` is a temporary quarantine and staging area.
It is not the canonical BPDSL specification hierarchy.

A future `product/records/spec/bpdsl/index.md` defines:

| topic | contract |
|---|---|
| Purpose | Preserve BPDSL-related content removed from mixed Design Records specifications. |
| Ownership status | Placement creates no PRODUCT ownership claim. `bpdsl/records/spec/**` remains the expected owner for BPDSL-internal contracts, subject to migration review. |
| Allowed content | Existing DSL, source, render, implementation-flow, or BPDSL artifact descriptions. Include only context required to preserve meaning. |
| Prohibited work | No BPDSL redesign, normalization, schema definition, resolver contract, render contract, generation contract, runtime contract, or MCP contract. |
| Review boundary | Review separation, preservation, and absence of canonical claims. Exclude BPDSL correctness, duplication, final layout, final ownership, and integration design. |
| Migration trigger | Review every staged item during BPDSL migration or an explicit integration requirement. |
| Exit condition | Move to app-local BPDSL specs, redefine as PRODUCT policy, relocate elsewhere, or delete. Then remove or explicitly redefine this area through an accepted decision. |

Temporary placement does not create a normative dependency from Design Records to BPDSL.

### Unadopted future integration material

Current specifications contain current contracts.
They do not preserve catalogs of unadopted mechanisms.

Inactive `yaml:` references, internal-design endpoints, coverage artifacts, and unrealized relations are extracted or removed after evidence transfer.

Possible follow-up artifacts include:

- Investigation;
- Requirement;
- ADR candidate;
- Work Item.

This ADR does not choose the destination for every statement.

### Dependency direction

| dependency | decision |
|---|---|
| Design Records → DRMCP | No normative dependency. |
| DRMCP → Design Records | Allowed. DRMCP may implement and expose Design Records contracts. |
| Design Records ↔ BPDSL integration | No normative dependency in either direction. Integration is deferred until a concrete requirement defines endpoints, direction, validation, and owner. |
| Brewprint profile → Design Records | Allowed. Brewprint may instantiate generic Design Records contracts. |
| Temporary PRODUCT BPDSL staging → canonical BPDSL | Contextual reference only. Staging creates no ownership claim. |
| Responsibility-boundary validator → Design Records | Allowed. The validator may evaluate Tasks against PRODUCT-owned Task semantics. |
| Responsibility-boundary validator → DRMCP | No current normative dependency. Future integration requires separate authority. |

### Implementation sequencing constraints

The restructuring separates these change classes:

| change class | constraint |
|---|---|
| Ownership decision | Accept the semantic boundary before PRODUCT spec edits. |
| Semantic rewrite batches | Keep each batch small enough for reviewable diffs. Target three to seven files. |
| File relocation | Move cleaned files without semantic rewrites. |
| App-local handoff | Transfer DRMCP or BPDSL-owned material in owner-specific batches. |
| Mechanical ref synchronization | Update refs only after target ownership and paths are accepted. |
| Validation and stale-ref cleanup | Attribute diagnostics to the accepted migration scope. |

Semantic rewrites, relocation, and broad ref synchronization are not combined in one change.

The full migration requires a separate Work Item after this ADR is accepted.
PRODUCT specifications remain unchanged while this ADR is proposed.

## Rationale

Explicit semantic areas make placement reviewable.
They also prevent physical location from becoming an accidental ownership claim.

The Design Records boundary remains useful when DRMCP or BPDSL internals change.
Brewprint profile and compatibility content can change without rewriting generic contracts.

Temporary BPDSL staging bounds the current work.
Temporary BPDSL staging avoids both premature migration and indefinite PRODUCT ownership.

A direct validator area keeps standalone semantic behavior outside generic Design Records ownership.
The direct area also avoids implying that current DRMCP must implement the validator.

Separating semantic rewrites from moves and ref synchronization prevents another broad mixed diff.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Retain `concepts/` as the generic owner | The name does not define a stable semantic boundary. |
| Use only PRODUCT-versus-DRMCP ownership | The model cannot classify Brewprint profile, compatibility, BPDSL, or deferred integration material. |
| Create top-level `compatibility/` | Current compatibility material is Brewprint-specific. |
| Create `deferred-integration/` specs | Unadopted mechanisms are not current contracts. |
| Place the standalone responsibility validator under `design-records/`. | The validator evaluates Design Records Tasks but is not a generic Design Records authoring or artifact-semantics contract. |
| Move all BPDSL material directly now | Final BPDSL ownership requires a separate migration review. |
| Combine rewrites, moves, and ref synchronization | The combined diff would be difficult to review and validate. |

## Consequences

- `PRODUCT-INV-SPEC-005` becomes the source classification for later planning.
- A separate Work Item will plan and execute the staged restructuring.
- PRODUCT spec files change only through bounded authoring after acceptance.
- The root router and area overviews are authored only after acceptance.
- Mechanical reference changes occur after semantic ownership and paths are accepted.
- The temporary BPDSL area requires an explicit migration trigger and exit review.
- `responsibility-boundary-validator/` is a direct PRODUCT child area.
- `spec:product` registers `spec:product.responsibility_boundary_validator` directly.
- `spec:product.design_records` does not own or register the validator.

## Evidence

- PRODUCT-INV-SPEC-005: Full file-level classification and required mixed-file section classification.
- PRODUCT-WORK-NAMESPACE-001: Historical namespace restructuring context only.
- PRODUCT-INV-SPEC-004: Evidence that the earlier binary ownership model was insufficient.
- PRODUCT-TASK-SPEC-019-07 R-002: Selects the direct PRODUCT validator target.
- PRODUCT-ADR-SPEC-016: Defines the standalone validator and current DRMCP separation.
- PRODUCT-TASK-SPEC-019-16: Projects the direct PRODUCT-root ownership amendment and canonical Specification.

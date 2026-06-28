# DRMCP-REQ-MCP-004: End-to-end Design Records MCP realignment milestone

- **id**: DRMCP-REQ-MCP-004
- **status**: accepted
- **date**: 2026-06-25
- **source_refs**:
  - DRMCP-ADR-MCP-001
  - DRMCP-INV-MCP-002
  - DRMCP-REQ-MCP-001
  - DRMCP-REQ-MCP-002
  - DRMCP-REQ-MCP-003
  - PRODUCT-REQ-SPEC-003
  - PRODUCT-ADR-SPEC-001
- **work_items**:
  - DRMCP-WORK-MCP-002

## Requirement

The Design Records MCP contract realignment must be delivered as one coordinated milestone across the current read baseline, portable standards-package production, package consumption, and package-backed authoring runtime.

Each implementation area must remain owned by its source Requirement and dedicated Work Item.
The milestone must expose the cross-Work-Item dependency order, hard gates, accepted parallelism, and integrated completion evidence without duplicating child implementation plans.

## Evidence

- `DRMCP-ADR-MCP-001` defines a replacement baseline and phased delivery order across read operations, portable authoring standards, and artifact authoring.
- `DRMCP-REQ-MCP-001`, `DRMCP-REQ-MCP-002`, and `DRMCP-REQ-MCP-003` intentionally separate read, authoring, and portable-package consumer responsibilities.
- `PRODUCT-REQ-SPEC-003` separately owns portable-package production, drift detection, and release validation.
- `DRMCP-WORK-MCP-001` and `PRODUCT-WORK-SPEC-013` already provide detailed plans for two independent workstreams.
- Authoring runtime implementation is unsafe until a released PRODUCT package can be loaded and validated by DRMCP.
- Read-baseline correction and PRODUCT package production can proceed in parallel, but their downstream integration requires explicit gates.

## Required Outcome

### Milestone workstreams

The milestone must coordinate, without absorbing, these workstreams:

1. Current-format read baseline under `DRMCP-REQ-MCP-001`.
2. Portable standards-package production under `PRODUCT-REQ-SPEC-003`.
3. Portable package consumer P0 under `DRMCP-REQ-MCP-003`.
4. Package guidance and authoring-integration follow-up under `DRMCP-REQ-MCP-003` where required beyond P0.
5. Package-backed authoring transaction delivery under `DRMCP-REQ-MCP-002`.
6. Integrated contract, portability, and implementation review.

Each workstream must use a dedicated Work Item with its own source Requirement.
Child Work Items own their internal tasks, detailed dependencies, implementation evidence, and closure decision.

### Hard gates

The milestone must preserve these hard gates:

- Brewprint compatibility must remove `V01-SPEC-*` before legacy fallback implementation is accepted.
- Corrected read contracts and fixtures must precede read-baseline implementation for the affected area.
- `DRMCP-WORK-SPEC-001` and `DRMCP-WORK-SPEC-002` disposition must be coordinated with PRODUCT validation-policy owner pointers before either Work Item is superseded or absorbed.
- PRODUCT package production must establish a reviewable producer contract before DRMCP package-consumer implementation depends on it.
- A released or reviewable package fixture must exist before DRMCP package loader and load-time validation are accepted.
- DRMCP package-consumer P0 must be complete before package-backed authoring runtime implementation is accepted.
- Authoring runtime must not hard-code PRODUCT semantics or depend directly on the Brewprint repository layout.
- Integrated milestone closure requires all required child Work Items to be done and independently reviewable.

### Parallel work

The milestone must allow these activities to proceed in parallel where their local Work Items permit:

- current read-baseline contract correction and PRODUCT portable-package design;
- PRODUCT package generation work and DRMCP package-consumer contract design after the producer interface is reviewable;
- REQ-002 authoring contract design and earlier read/package work;
- child Work Item planning before all upstream implementation is complete.

Parallel planning does not waive runtime or acceptance gates.

### Milestone-level tracking

Milestone tasks must track child Work Item lifecycle at a coarse level:

- confirm or create the child Work Item;
- confirm its source Requirement and boundary;
- allow detailed execution to remain inside that Work Item;
- verify its completion evidence and review gate;
- record the milestone dependency it satisfies;
- close the milestone task only after the child Work Item reaches `done` or receives an explicit accepted disposition.

Milestone tasks must not copy child task status or implementation logs into the milestone record.

### Integrated completion

The milestone must finish with an integrated review proving that:

- current read behavior follows the corrected contract;
- legacy compatibility remains explicit, exact, read-only, and configuration-gated;
- the portable package remains a PRODUCT-owned distribution artifact;
- DRMCP loads and validates the package without host `product` dependency;
- authoring runtime consumes package semantics rather than duplicating them;
- workflow, spec, and investigation authoring phases respect their accepted sequence;
- normal API surfaces preserve canonical identity and path-hiding boundaries;
- no child Work Item remains open for a capability claimed complete by the milestone.

## Explicitly Excluded Scope

- Replacing the source Requirements of child Work Items.
- Defining child implementation details or copying child task plans.
- Treating the milestone Work Item as the semantic owner of read, package, or authoring contracts.
- Closing a child Requirement solely because a milestone task is complete.
- BPDSL design or migration.
- UI work.
- General-purpose package registries or network distribution.
- Reopening accepted ownership decisions in `PRODUCT-ADR-SPEC-001` or `DRMCP-ADR-MCP-001`.

## Boundary

This Requirement owns only end-to-end delivery coordination, dependency visibility, milestone gates, and integrated closure evidence.

The child Requirements retain their existing ownership:

- `DRMCP-REQ-MCP-001`: current read baseline and legacy fallback;
- `PRODUCT-REQ-SPEC-003`: portable package production and release validation;
- `DRMCP-REQ-MCP-003`: package loading, load-time validation, guidance projection, and runtime integration;
- `DRMCP-REQ-MCP-002`: package-backed authoring transaction behavior.

The milestone Work Item must reference child Work Items as delivery dependencies.
It must not list their implementation tasks as its own tasks or alter their source-requirement relations.

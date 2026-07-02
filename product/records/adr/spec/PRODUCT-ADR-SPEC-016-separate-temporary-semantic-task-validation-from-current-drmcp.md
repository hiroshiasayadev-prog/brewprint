# PRODUCT-ADR-SPEC-016: Separate temporary semantic Task validation from current DRMCP

- **status**: accepted
- **date**: 2026-07-02
- **depends_on**:
  - PRODUCT-ADR-SPEC-001
  - PRODUCT-ADR-SPEC-015
- **supersedes**: []
- **migrated_to_spec**: null

## Context

The immediate need is a lightweight semantic Task responsibility validator.
The current DRMCP owns structural Design Record operations and current tool projections.

Combining the temporary validator with current DRMCP would expand the current DRMCP contract.
That expansion would introduce separate tool-surface, diagnostic, architecture, and lifecycle decisions.

The semantic contract also needs one canonical PRODUCT owner.
The standalone application now has the app namespace `TRV`.
Future DRMCP integration must remain separately decidable.

## Decision

PRODUCT owns:

- the semantic Task responsibility-validation contract;
- the temporary standalone validator product boundary;
- the dedicated canonical Specification at `spec:product.responsibility_boundary_validator`.

`TRV` owns app-local design and later implementation for the standalone validator.

For each invocation, `TRV` owns:

- reading one Task record;
- selecting the checklist from `task_type`;
- injecting common and type-specific criteria;
- orchestrating the local semantic model;
- receiving criterion-level results and reasons;
- applying the deterministic aggregation defined by PRODUCT-ADR-SPEC-015;
- separating semantic results from standalone execution failures.

Current DRMCP retains its current ownership of:

- structural parsing;
- structural validation;
- diagnostics;
- indexing;
- current tool projections for accepted Design Record contracts.

W019 does not integrate the temporary validator into DRMCP.
W019 does not modify current DRMCP Specifications, tools, diagnostics, source, or tests.

Any future DRMCP integration requires a separate Requirement or Work Item.
That future authority decides the integration contract and ownership effects.

Treat the older `MCP` wording in PRODUCT-TASK-SPEC-019-01 D-006 as historical Evidence.
Do not rewrite the completed Task.
Use PRODUCT-TASK-SPEC-019-03 D-001 and the W019 reconciliation decisions as current authority.

## Rationale

The standalone boundary satisfies the immediate validation need without expanding current DRMCP.
The split keeps semantic product ownership separate from current structural tool ownership.

A dedicated PRODUCT Specification provides one stable owner for reusable semantic behavior.
The `TRV` namespace makes the existing standalone ownership boundary explicit.
Future integration can consume the PRODUCT contract without predetermining the TRV interface or architecture.

Preserving the completed T01 record keeps the workflow history honest.
The later decision corrects the current product boundary without rewriting earlier Evidence.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Add semantic Task evaluation to current DRMCP in W019. | The change requires separate integration, diagnostics, architecture, and lifecycle decisions. |
| Let current DRMCP own the semantic contract. | The semantic contract is a standalone PRODUCT concern, not current DRMCP behavior. |
| Place validator behavior in `task_authoring`. | That Specification owns the evaluated Task contract, not validator execution semantics. |
| Leave the canonical owner implicit. | Requirement, checklist, workflow, and implementation work could diverge. |
| Rewrite T01 D-006 to remove the historical `MCP` wording. | The completed Task must remain historical Evidence. |

## Consequences

- `spec:product.responsibility_boundary_validator` becomes the canonical semantic owner.
- The parent `spec:product` topic must register that Specification.
- `spec:product.design_records.authoring_standards.task_authoring` may contain only a narrow ownership and usage relation.
- `TRV` owns app-local Requirement, ADR, Specification, and later implementation work.
- PRODUCT-WORK-SPEC-021 owns reviewed PRODUCT conceptual design and the namespace bootstrap needed for handoff.
- TRV source path, executable interface, model, runtime, packaging, and deployment remain app-local decisions.
- Future DRMCP integration remains outside W019 and requires separate design authority.
- Current DRMCP artifacts remain unchanged.

## Evidence

- `PRODUCT-TASK-SPEC-019-03`: D-001 temporary standalone boundary and future integration deferral.
- `PRODUCT-TASK-SPEC-019-07`: R-001 ownership split, R-002 canonical target, and R-003 historical treatment.
- `PRODUCT-TASK-SPEC-019-12`: B-002 routing boundary and `create` disposition.
- `PRODUCT-ADR-SPEC-001`: PRODUCT root semantic-placement authority.
- `PRODUCT-TASK-SPEC-021-06`: D-002 selected `TRV`; D-003 fixed PRODUCT and TRV ownership boundaries.
- `PRODUCT-TASK-SPEC-021-08`: B-001 routed this non-material ownership clarification to ADR-016 amendment.

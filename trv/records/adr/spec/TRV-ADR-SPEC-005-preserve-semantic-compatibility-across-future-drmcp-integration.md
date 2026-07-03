# TRV-ADR-SPEC-005: Preserve semantic compatibility across future DRMCP integration

- **status**: accepted
- **date**: 2026-07-02
- **depends_on**:
  - PRODUCT-ADR-SPEC-015
  - PRODUCT-ADR-SPEC-016
  - PRODUCT-ADR-SPEC-017
  - TRV-ADR-SPEC-003
  - TRV-ADR-SPEC-004
- **supersedes**: []
- **migrated_to_spec**: null

## Context

The standalone TRV application is a temporary delivery boundary.
A future Work Item may integrate semantic Task validation into DRMCP.

Future integration can replace transport, record access, configuration, and adapter mechanics.
The integration must not silently change the PRODUCT-owned semantic validator contract.

Compatibility therefore needs a clear preservation boundary.
The boundary must reference PRODUCT authority rather than duplicate semantic rules under TRV.

## Decision

Preserve semantic compatibility across any future DRMCP integration.

Preserve these behaviors:

- criterion-level boolean results and concise reasons;
- PRODUCT-defined criterion-set validation;
- PRODUCT-defined overall result;
- distinct semantic, structural, and execution outcomes;
- caller-owned human judgment;
- post-authoring and post-Evidence invocation semantics.

A future DRMCP integration may replace these mechanics:

- MCP transport;
- server and tool identity;
- Task-path input;
- configuration shape;
- adapter implementations;
- standalone source implementation.

Do not perform current DRMCP integration through this ADR.
Do not change the current DRMCP tool contract.

Semantic compatibility means preserving the behavior owned by PRODUCT authority.
TRV must not restate or take ownership of that semantic contract.

## Rationale

The semantic validator behavior is useful beyond the first standalone delivery.
The standalone transport and source mechanics exist only to satisfy the immediate application boundary.

Preserving semantics allows a later integration to replace temporary mechanics without changing validation meaning.
Allowing mechanical replacement avoids freezing the standalone interface as a permanent DRMCP contract.

Referencing PRODUCT authority maintains one semantic owner.
The reference also prevents TRV and future DRMCP Specifications from developing competing definitions.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Preserve the entire standalone interface unchanged. | Future DRMCP integration may require different transport, identity, input, and configuration mechanics. |
| Preserve only the final overall result. | Criterion judgments, failure separation, and human workflow semantics would be lost. |
| Copy PRODUCT semantic rules into this ADR. | Duplicate authority would create drift and conflicting ownership. |
| Integrate TRV into current DRMCP now. | Current integration is explicitly outside W003 and requires separate design authority. |
| Let future integration redefine failure classes. | Structural, execution, and semantic outcomes would no longer remain compatible. |

## Consequences

- T05 must define the current compatibility boundary without duplicating PRODUCT semantics.
- A future DRMCP Work Item may replace the standalone interface and implementation mechanics.
- Future integration must preserve the PRODUCT-owned semantic behaviors listed by this ADR.
- Future integration must preserve caller-owned human judgment from PRODUCT-ADR-SPEC-017.
- Current DRMCP Specifications, tools, diagnostics, source, and tests remain unchanged.
- TRV-ADR-SPEC-003 and TRV-ADR-SPEC-004 remain current standalone choices, not permanent DRMCP requirements.
- Any future semantic change must return to PRODUCT authority rather than being introduced as integration detail.

## Evidence

- TRV-TASK-SPEC-001-02 D-014 selected semantic compatibility with replaceable standalone mechanics.
- TRV-TASK-SPEC-003-02 routed D-014 into this ADR boundary.
- PRODUCT-ADR-SPEC-015 defines Task-local semantic evaluation and distinct failure outcomes.
- PRODUCT-ADR-SPEC-016 separates the standalone TRV application from current DRMCP.
- PRODUCT-ADR-SPEC-017 defines the two invocation points and caller-owned human judgment.
- TRV-ADR-SPEC-003 and TRV-ADR-SPEC-004 define the replaceable standalone interface and input choices.

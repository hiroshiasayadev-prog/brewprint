# TRV-TASK-SPEC-002-12: Amend TRV architecture ADRs

- **id**: TRV-TASK-SPEC-002-12
- **status**: done
- **date**: 2026-07-02
- **work_item**: TRV-WORK-SPEC-002
- **task_type**: authoring
- **estimate**: 0.5d
- **depends_on**:
  - TRV-TASK-SPEC-002-11
- **outputs**:
  - TRV-ADR-SPEC-001
  - TRV-ADR-SPEC-002
  - TRV-TASK-SPEC-002-12

## Goal

Amend the two accepted TRV architecture ADRs with the non-material responsibility clarifications routed by T10.

## Work

- Amend TRV-ADR-SPEC-001 using T10 Boundary A and T09 D-003, D-004, D-007, D-008, D-009, D-011, and D-012.
- Preserve the selected ports-and-adapters alternative, inward-dependency rationale, five top-level responsibility areas, and no-domain-layer choice.
- Clarify the inbound application port, separate Task-record and checklist-catalog ports, core-owned prompt construction, MCP-versus-use-case validation boundary, application outcome ownership, and startup/composition responsibility.
- Amend TRV-ADR-SPEC-002 using T10 Boundary B and T09 D-010.
- Preserve the provider-neutral application port, externally managed remote Ollama runtime, and adapter-owned provider mechanics.
- Clarify that the model port receives the completed prompt and returns decoded criterion-result candidates or execution failure, while the core validates completeness and derives overall compliance.
- Keep exact prompt wording, result fields, Go declarations, HTTP schema, retry, timeout, and configuration names outside the ADRs.
- Preserve `status: accepted`; do not create or supersede an ADR.
- Set `migrated_to_spec` to `null` on both amended ADRs. T15 restores the migration date only after the revised Specification projection passes review.

This Task must not:

- change a T09 decision or T10 routing outcome;
- author or reorganize Specifications;
- choose exact implementation structure;
- perform independent review, correction, synchronization, implementation, stage, or commit work.

## Done condition

- TRV-ADR-SPEC-001 expresses the complete Boundary A durable ownership choices without reversing its accepted alternative.
- TRV-ADR-SPEC-002 expresses the complete Boundary B model-port responsibility without reversing its accepted alternative.
- Both ADRs preserve honest history and remain independently supersedable.
- No detailed-design or external MCP contract content is introduced.

## Verification

- Compare both amended ADRs with the T10 routing ledger.
- Confirm every included decision is represented once in the correct ADR boundary.
- Confirm no selected alternative, core rationale, status, or supersession relation changed.
- Confirm exact W003 and W004 details remain excluded.

## Evidence

- T10 routed D-003, D-004, D-007, D-008, D-009, D-011, and D-012 to a non-material amendment of TRV-ADR-SPEC-001.
- T10 routed D-010 to a non-material amendment of TRV-ADR-SPEC-002.
- Amended TRV-ADR-SPEC-001 `## Decision`, `## Rationale`, `## Rejected alternatives`, `## Consequences`, and `## Evidence`.
- TRV-ADR-SPEC-001 now records the five component areas, inbound port, two outbound input ports, core-owned prompt construction, MCP-versus-use-case validation boundary, application outcomes, and startup responsibility.
- Amended TRV-ADR-SPEC-002 `## Decision`, `## Rationale`, `## Rejected alternatives`, `## Consequences`, and `## Evidence`.
- TRV-ADR-SPEC-002 now records the completed-prompt input, decoded criterion-result or execution-failure output, Ollama execution boundary, and core-owned completeness and aggregation checks.
- Both ADRs retain `status: accepted`, the original date, dependencies, and empty supersession lists.
- Both ADRs set `migrated_to_spec: null` pending T13 authoring and T14 review.
- No ADR was created or superseded.
- No Specification, W003 contract, W004 detailed design, implementation, review, stage, or commit work occurred.
- Scoped diff inspection confirmed the declared three-file write boundary.
- Result: `PASS`.

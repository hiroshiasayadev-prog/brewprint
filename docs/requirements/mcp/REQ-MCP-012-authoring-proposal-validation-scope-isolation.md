# REQ-MCP-012: Authoring proposal validation scope isolation

- **id**: REQ-MCP-012
- **status**: accepted
- **date**: 2026-06-02
- **source_refs**:
  - REQ-MCP-008
  - ADR-093
  - SPEC-design-records-mcp-tools
  - SPEC-design-records-mcp-schema
- **work_items**:

## Requirement

Design Records MCP authoring proposal validation needs scope isolation so proposal-local diagnostics are not polluted by unrelated pre-existing repository diagnostics.

Authoring transactions should make it clear whether a proposal itself is valid. If `propose_record_create` or `propose_record_update` performs a full repository scan and emits unrelated diagnostics, the caller cannot reliably tell whether the proposed change is invalid or whether the repository already contained independent validation errors.

Proposal diagnostics must be actionable and reproducible from the same affected record set.

## Evidence

During dogfooding of authoring proposal creation, unrelated validation errors appeared during proposal handling, including investigation reference diagnostics that were not reproduced by a direct `validate_records(kind=investigation)` check.

Separately, direct validation showed known repository issues in other record kinds, such as:

- `WORK-DATA-010` has an invalid work item status and references absent `TASK-DATA-010-03`;
- `TASK-MCP-005-01` through `TASK-MCP-005-03` are missing metadata required by the current task validation contract.

Those existing repository issues may be real cleanup work, but they should not make an unrelated authoring proposal appear proposal-invalid unless they are part of the proposal impact set.

Close evidence on 2026-06-02: `WORK-MCP-011` included this coupled validation-scope requirement in its spec, implementation, regression tests, and runtime smoke. `TASK-MCP-011-04` confirmed proposal-local `validation.diagnostics` stayed empty for successful create/update proposals despite unrelated existing repository diagnostics, and `accept_proposed_write` pre-write validation completed without being blocked by unrelated diagnostics.

## Required Outcome

A follow-up work item should define and implement validation scope rules for authoring proposals.

The expected properties are:

- propose-time validation validates the proposed record and directly affected reciprocal records only;
- unrelated pre-existing repository diagnostics are not returned as proposal-local blocking diagnostics;
- if the tool chooses to report pre-existing repository health issues, they are clearly separated from proposal-local validation under a distinct category or response field;
- diagnostics emitted by propose-time validation can be reproduced by an explicit `validate_records` call over the same affected record set;
- create and update operations use the same affected-record-set model unless the public contract documents a difference;
- proposal responses remain actionable for AI assistants without requiring a full repository cleanup first.

## Explicitly Excluded Scope

This requirement does not require:

- fixing all existing validation errors in the repository;
- weakening workflow artifact validation strictness;
- changing investigation, work item, or task metadata rules;
- suppressing diagnostics for records directly affected by a proposal;
- adding multi-record authoring transactions;
- changing canonical reference grammar.

## Boundary

This requirement owns the validation scope and diagnostic isolation behavior for Design Records MCP authoring proposal operations.

It does not own authoring input normalization issues such as `body` versus `fields`, duplicate `id` placement, or domain case normalization. Those are captured separately by `REQ-MCP-011`.

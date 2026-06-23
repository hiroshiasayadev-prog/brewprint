# PRODUCT-TASK-SPEC-011-01: Correct v2 ADR grammar

- **id**: PRODUCT-TASK-SPEC-011-01
- **status**: done
- **date**: 2026-06-23
- **work_item**: PRODUCT-WORK-SPEC-011
- **source_requirement**: PRODUCT-REQ-SPEC-002
- **estimate**: 0.2d
- **depends_on**:
- **outputs**:
  - `product/records/spec/concepts/namespace-model/v2-grammar.md`

## Goal

Correct the v2 artifact ID grammar so new ADR IDs use the same app, kind, domain, and sequence structure as other sequential artifacts.

## Work

- Add ADR to the shared REQ / WORK / INV grammar family.
- Define the new ADR form as `<APP_NAMESPACE>-ADR-<DOMAIN_NAMESPACE>-<SEQUENCE>`.
- Define ADR sequences as three-digit, zero-padded values.
- Add a namespace-aware ADR example.
- Preserve existing V01 ADR IDs as compatibility records.

## Done condition

| item | done when |
|---|---|
| ADR grammar corrected | The v2 grammar includes ADR in the app + kind + domain + sequence form. |
| Sequence format defined | ADR uses a three-digit, zero-padded sequence. |
| Example present | The grammar includes a namespace-aware ADR example. |
| Compatibility preserved | Existing V01 ADR IDs remain unchanged. |

## Verification

- Confirm the grammar heading groups REQ / WORK / INV / ADR.
- Confirm the sequence table includes ADR with `001`-style formatting.
- Confirm `DRMCP-ADR-MCP-001` appears as an example.
- Confirm the compatibility note retains existing `V01-ADR-NNN` IDs.

## Evidence

- `product/records/spec/concepts/namespace-model/v2-grammar.md` defines REQ / WORK / INV / ADR as `<APP_NAMESPACE>-<ARTIFACT_KIND>-<DOMAIN_NAMESPACE>-<SEQUENCE>`.
- The sequence table defines REQ / WORK / INV / ADR as three-digit, zero-padded sequences.
- The examples include `DRMCP-ADR-MCP-001`.
- The ADR compatibility note states that existing ADRs retain `V01-ADR-NNN`.

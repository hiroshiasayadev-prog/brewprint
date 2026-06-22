# PRODUCT-TASK-SPEC-005-15: DRMCP — Opus review of design-records-mcp/ spec files

- **id**: PRODUCT-TASK-SPEC-005-15
- **status**: done
- **date**: 2026-06-17
- **work_item**: PRODUCT-WORK-SPEC-005
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-005-14
- **outputs**:
  - Review findings (inline in Evidence)

## Goal

Independent Opus review of the `drmcp/records/spec/design-records-mcp/` migration output against the content in `drmcp/old/`. Confirm completeness, format compliance, and translation accuracy before removing the staging files. Given the larger scope (30 output files vs. BPDSL's 16), pay particular attention to the `schema.md` → `tools.md` split boundary and to content that referenced the authoring-transaction model from both files.

## Work

| area | what to check |
|---|---|
| completeness | All content from `drmcp/old/overview.md` / `schema.md` / `tools.md` is accounted for in the new tree. No section silently dropped. |
| format compliance | H1 format, H1-adjacent metadata, required sections by kind, `## Topics` columns. |
| parent chain | Every child spec's `parent:` matches an existing `## Topics` row in the declared parent, tracing clean to `overview.md` (`parent: -`). |
| English translation | No Japanese remains in H1 titles, H2 section titles, or table headers. Body prose may retain Japanese until a separate pass if needed. |
| tool contract shape | All 13 tool files have `contract_class: interface` and complete `## Request` / `## Response` / `## Errors` sections matching the source `tools.md` tool definition. |
| schema/tools split boundary | `schema/authoring-transaction-schema.md` (data shapes) and `tools/authoring-transaction-model.md` (shared tool-response concepts) do not duplicate each other's content — confirm the division is real, not arbitrary. |
| diagnostics cross-ref | `tools/validate-records.md` correctly defers to `schema/diagnostics.md` instead of duplicating the diagnostic category table. |
| internal refs | Cross-spec `[]()` links point to files that exist under `design-records-mcp/` (not to `drmcp/old/`). |
| V01 ID references | Confirm disposition of embedded `V01-ADR-*` / `V01-INV-*` example blocks and ADR `depends_on` lists — consistent with how the BPDSL batch (PRODUCT-TASK-SPEC-005-03) handled ADR-attribution refs. |
| deferred relocation candidates | Per PRODUCT-INV-SPEC-004: flag `namespace-scanning.md`, `schema/id-normalization.md`, `schema/discovery.md`, and `resolver.md` as PRODUCT-owned-semantics content currently migrated in place under DRMCP. Do not relocate them in this task — record them as deferred relocation candidates in Evidence, pending PRODUCT-WORK-SPEC-004's accepted plan. |
| validator | `validate_spec.py drmcp/records/spec/design-records-mcp/ --strict` exits 0. |

## Done condition

| item | done when |
|---|---|
| review complete | Opus review report is attached in Evidence. |
| findings classified | Each finding is classified: must-fix before 005-16, or defer. |
| deferred relocation candidates recorded | Evidence lists the PRODUCT-INV-SPEC-004 candidates found in the migrated output, explicitly marked "format-migrated, not relocated." |
| user sign-off | User approves proceeding to 005-16. |

## Verification

This task is itself a review gate. Proceed to PRODUCT-TASK-SPEC-005-16 only after user sign-off on findings.

## Evidence

### Verdict: PASS — no must-fix issues

Review run: 2026-06-17. Reviewer: Opus 4.7 (external). Cross-check by Claude Code: confirmed all factual claims against actual files.

### Per-area results

| area | result | notes |
|---|---|---|
| completeness | ok | All sections from overview.md / schema.md / tools.md accounted for across 30 files. |
| format compliance | ok | H1 prefix, H1-adjacent metadata, required sections by kind, contract_class: interface — all correct on all 30 files. |
| parent chain | ok | Every parent: value resolves to an existing id; Topics rows in all 3 overview files match their children; root traces to parent: -. |
| English translation | ok | No Japanese in any H1/H2/H3/H4 or table header. Japanese remains only in: > Source: attribution line section names, example JSON strings (acceptable). |
| tool contract shape | ok | All 13 tool Contract files have contract_class: interface and complete Request/Response/Errors sections matching source tools.md. |
| schema/tools split boundary | ok | schema/authoring-transaction-schema.md (data shapes) and tools/authoring-transaction-model.md (operational concepts) are non-redundant; tools file cross-refs schema file. |
| diagnostics cross-ref | ok | tools/validate-records.md lists authoring categories in short bullets and defers to spec:drmcp.design_records_mcp.schema.diagnostics — does not duplicate full table. |
| internal refs | ok | All 88 spec:drmcp.* cross-references resolve to existing file IDs in the tree. |
| V01 ID references | ok | Front-matter depends_on lists removed; ADR attributions preserved as > Source: lines; example ID blocks preserved verbatim. Consistent with BPDSL precedent. |
| deferred relocation candidates | ok | All 4 files have deferred relocation note; notes are accurate (see section below). |
| validator | ok | validate_spec.py drmcp/records/spec/design-records-mcp/ --strict exits 0. (Confirmed in PRODUCT-TASK-SPEC-005-14.) |

### Deferred relocation candidates

| file | note present | classification |
|---|---|---|
| `resolver.md` | yes | format-migrated, not relocated |
| `namespace-scanning.md` | yes | format-migrated, not relocated |
| `schema/id-normalization.md` | yes | format-migrated, not relocated |
| `schema/discovery.md` | yes | format-migrated, not relocated |

Pending Phase 2 relocation per PRODUCT-WORK-SPEC-004 accepted plan.

### Deferred findings

1. **Authoring-only category cross-ref wording** (`tools/validate-records.md:107`): Cross-ref updated to point to `spec:drmcp.design_records_mcp.tools.propose_record_update` and `spec:drmcp.design_records_mcp.tools.propose_record_create`. Source schema.md also did not define these categories there — not a migration oversight. **Resolved 2026-06-22.**

### User sign-off

Obtained 2026-06-17.

# DRMCP-TASK-MCP-001-03: Track discovery and active-index contract correction

- **id**: DRMCP-TASK-MCP-001-03
- **status**: done
- **date**: 2026-06-26
- **work_item**: DRMCP-WORK-MCP-001
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d coordination
- **depends_on**:
  - DRMCP-TASK-MCP-001-01
  - DRMCP-TASK-MCP-001-02
- **outputs**: []

## Goal

Accept the corrected current discovery, spec parsing, and active-index contract gate.

## Work

- Track `DRMCP-WORK-MCP-003` as the selected child Work Item.
- Delegate active-root discovery, current spec parsing, duplicate identity, and index-separation contract work to `DRMCP-WORK-MCP-003`.
- Confirm that `DRMCP-WORK-MCP-003` remains sourced from `DRMCP-REQ-MCP-001`.
- Track the child Work Item through contract review and `done`.
- Record the child Work Item ID and accepted evidence here.

This Task does not modify DRMCP discovery, schema, or index specs.
All detailed contract work belongs to the selected child Work Item.

## Done condition

- `DRMCP-WORK-MCP-003` is `done`.
- Current roots and optional legacy roots have separate discovery and index contracts.
- Current specs use H1-adjacent metadata and path-derived canonical `spec:` refs.
- YAML front matter is rejected as an active spec metadata source.
- Duplicate current identity behavior is deterministic and diagnostic.
- The child review has no blocking or major findings.
- The `DRMCP-WORK-MCP-003` evidence pointer is recorded here.

## Verification

- Review the corrected discovery and shared-schema contracts.
- Confirm that legacy archive roots are excluded from active discovery.
- Confirm that this Task contains no direct contract implementation evidence.

## Evidence

- Selected child Work Item: `DRMCP-WORK-MCP-003`.
- Child source Requirement: `DRMCP-REQ-MCP-001`.
- Ownership: DRMCP owns current-root loading, discovery, parsing behavior, and active-index construction while PRODUCT supplies canonical semantics.
- Child execution state: `DRMCP-WORK-MCP-003` completed T01 through T05 and is `done`.
- Completion evidence accepted from `DRMCP-WORK-MCP-003`:
  - final changed contract set passed strict validation: `[strict] All 30 file(s) OK.`;
  - final independent re-review verdict: `PASS`;
  - previous major ownership-boundary finding: `CLOSED`;
  - blocking findings: none;
  - major findings: none;
  - current and legacy discovery/index contracts remain separate;
  - W004, W005, W006, fixture, implementation, and test scope remain delegated.
- Hub acceptance completed on 2026-06-27.

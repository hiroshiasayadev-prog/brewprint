# TASK-MCP-021-02: update named_section_replace normalization contract

- **id**: TASK-MCP-021-02
- **status**: done
- **date**: 2026-06-05
- **work_item**: WORK-MCP-021
- **source_requirement**: REQ-MCP-022
- **estimate**: 0.5d
- **depends_on**:
  - TASK-MCP-021-01
- **outputs**:
  - Updated Design Records MCP tool/spec contract for heading-safe named_section_replace body normalization
  - Diagnostic contract for section_replacement_body_heading_stripped warning

## Goal

Update the Design Records MCP authoring update contract so `named_section_replace` explicitly documents heading-safe replacement body normalization and its warning diagnostic.

## Work

- Update the relevant MCP tool/spec documentation for `propose_record_update`.
- Document that only the first non-empty matching Markdown heading line is stripped.
- Document that direct `body` and `body_cache_id` sources follow the same rule.
- Document warning diagnostic shape/category, including stripped heading text and level.
- Keep non-matching heading behavior and selector validation behavior unchanged.

## Done condition

- The public contract describes when normalization happens.
- The diagnostic category and warning semantics are documented.
- The contract does not claim multi-section or arbitrary replacement support.

## Verification

- Review the updated documentation against REQ-MCP-022 acceptance criteria.
- Confirm excluded scope remains excluded.

## Evidence
2026-06-05: Contract documentation updated per REQ-MCP-022.

Changes made:

- `docs/spec/design-records-mcp/tools.md` — `#### Named section replacement`: Added and revised "Heading-safe replacement body normalization" block documenting:
  - Normalization happens after `section_selector` resolves exactly one target section.
  - When the first non-empty line of the replacement body is a Markdown ATX heading matching the resolved selected section heading text and resolved selected section level, that one heading line is stripped before retained proposal creation.
  - The rule applies even when `section_selector.level` was omitted, because matching uses the resolved section level, not the raw selector input.
  - Normalization applies to both direct `body` and `body_cache_id` sources.
  - Only the first matching heading line is stripped; body-internal headings are preserved.
  - Non-matching heading text or level: no stripping, existing selector validation applies.
  - `section_replacement_body_heading_stripped` warning is returned with `stripped_heading` and `stripped_level` fields.
  - Warning diagnostics do not block retained proposal creation; error-severity diagnostics still block.
  - Multi-section replacement, arbitrary string replacement, and canonical workflow section name changes are explicitly excluded.

- `docs/spec/design-records-mcp/schema.md` — `## Diagnostic category`: Added `section_replacement_body_heading_stripped` (warning) to authoring transaction diagnostic table, with prose specifying `stripped_heading` and `stripped_level` as required additional fields.

REQ-MCP-022 acceptance criteria verified against updated spec: all criteria covered.

2026-06-05: Contract documentation revised per Codex review (NEEDS_FIX verdict).

Blocking findings fixed:

- Level matching was underspecified: normalization block now states that comparison uses the resolved selected section heading and level, not the raw `section_selector` input values, and applies regardless of whether `section_selector.level` was supplied.
- Diagnostic severity conflict resolved: tools.md shared diagnostic object now lists `warning` as a valid severity, with an explicit note that `warning` does not cause `validation.ok: false` and does not block proposal creation.

Non-blocking findings fixed:

- `stripped_heading` and `stripped_level` field names now explicitly named in the tools.md normalization block.
- Stale Evidence wording that referred to matching `section_selector.heading` and optionally `section_selector.level` was replaced with resolved-section wording, so future implementation agents are not pointed at raw selector matching.

Implementation note for TASK-MCP-021-03:

The implementation should resolve the named section first via `section_selector`, obtaining the resolved heading text and heading level from the matched section. It should then compare and strip the first non-empty replacement-body heading against those resolved values before constructing replacement lines and the retained proposal diff.

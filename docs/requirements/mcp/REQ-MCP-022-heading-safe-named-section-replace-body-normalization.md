# REQ-MCP-022: heading-safe named_section_replace body normalization

- **id**: REQ-MCP-022
- **status**: captured
- **date**: 2026-06-05
- **source_refs**:
  - TASK-DATA-011-01
  - REQ-MCP-021
- **work_items**:

## Requirement

`propose_record_update` with `update.type = named_section_replace` MUST safely normalize replacement bodies that accidentally include the selected section heading as the first content line.

When the caller has already selected a target section through `section_selector`, and the replacement body starts with the same Markdown heading line, the authoring layer SHOULD strip that redundant heading line before creating the retained proposal and return a warning diagnostic explaining the normalization.

This prevents accidental duplicated headings such as `## Evidence` inside the already-selected `Evidence` section.

## Evidence

During task close synchronization for `TASK-DATA-011-01`, an Evidence section update was proposed with a replacement body that included the `## Evidence` heading. The resulting diff duplicated the `## Evidence` heading, so the proposal had to be discarded and recreated with the heading removed.

Observed workflow note:

```text
Diff を確認すると ## Evidence が二重になっています。body に見出しを含めてしまったのが原因です。破棄して修正します。
```

Because `named_section_replace` already identifies the target section via `section_selector`, a leading duplicate of that same heading is redundant syntax noise, not intended section content.

## Required Outcome

For `named_section_replace`, the authoring update flow supports heading-safe replacement body normalization.

Acceptance criteria:

- When `body` first non-empty line is a Markdown heading matching the selected `section_selector.heading` and `section_selector.level`, that heading line is stripped before proposal creation.
- The same normalization applies when replacement content is supplied through `body_cache_id`.
- The retained proposal diff reflects the normalized replacement body.
- A warning diagnostic is returned, with a category such as `section_replacement_body_heading_stripped`, identifying the stripped heading and level.
- The normalization is limited to the first matching heading line only.
- Body-internal headings after the first content line are preserved as content.
- If the first heading does not match the selected section heading and level, existing validation/selector behavior applies.

## Explicitly Excluded Scope

- Multi-section replacement.
- Arbitrary text/string replacement.
- Interpreting multiple H2 sections inside one replacement body.
- Silently mutating repository files without retained proposal review and accept.
- Changing canonical workflow artifact section names.

## Boundary

This requirement belongs to the Design Records MCP authoring update contract. It is specifically about safe normalization for `named_section_replace` replacement bodies and body cache content. It does not change task semantics or DATA-domain behavior.

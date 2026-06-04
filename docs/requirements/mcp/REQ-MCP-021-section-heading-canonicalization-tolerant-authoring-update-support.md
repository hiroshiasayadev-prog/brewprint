# REQ-MCP-021: section heading canonicalization tolerant authoring update support

- **id**: REQ-MCP-021
- **status**: captured
- **date**: 2026-06-05
- **source_refs**:
  - TASK-DATA-011-01
  - REQ-MCP-020
- **work_items**:

## Requirement

Design Records MCP authoring updates SHOULD tolerate section heading case differences when the intended required section is otherwise unambiguous.

For workflow artifact sections with canonical headings, the authoring layer SHOULD be able to detect headings that differ only by letter case, propose the canonical heading correction, and continue with the intended update when it is safe to do so.

For example, `## Done Condition` SHOULD be recognized as the canonical required section `## Done condition` when there is no ambiguity.

## Evidence

During task close synchronization for `TASK-DATA-011-01`, the existing section heading used `## Done Condition` with uppercase `C`, while the validator required `## Done condition` with lowercase `c`.

The workflow discarded the proposed write, directly edited the file to correct the heading, and then retried the status update. That direct-edit detour indicates a missing MCP authoring capability: a trivial heading-case mismatch should not force manual filesystem edits.

Observed operator note:

```text
セクション名が Done Condition（大文字 C）になっていますが、validator は Done condition（小文字 c）を要求しています。まず破棄して、ファイルを直接修正します。
```

## Required Outcome

The authoring and validation flow supports safe canonicalization of section headings that differ only by case from required canonical headings.

Acceptance criteria:

- A workflow artifact with `## Done Condition` can be repaired to `## Done condition` through an MCP proposal instead of direct filesystem editing.
- Case-only heading mismatches are reported as warnings or info diagnostics, not opaque hard failures, when an unambiguous canonical heading exists.
- The proposed diff shows the canonical heading correction.
- Ambiguous heading matches still fail safely and report candidate headings.
- Non-case heading differences remain governed by existing selector and validation rules unless explicitly supported by a later requirement.

## Explicitly Excluded Scope

- Broad fuzzy matching for arbitrary misspelled section names.
- Changing canonical section names in authoring guides.
- Silently mutating files without a retained proposal and accept step.
- Relaxing required-section validation after canonicalization is applied.

## Boundary

This requirement belongs to the Design Records MCP authoring and workflow artifact validation boundary. It is about safe proposal-based repair of canonical section headings, not DATA-domain task semantics.

# PRODUCT-TASK-SPEC-001-02: Review spec format contract

- **id**: PRODUCT-TASK-SPEC-001-02
- **status**: done
- **date**: 2026-06-10
- **work_item**: PRODUCT-WORK-SPEC-001
- **source_requirement**: PRODUCT-REQ-SPEC-001
- **source_investigation**: PRODUCT-INV-SPEC-001
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-001-01
- **outputs**:
  - Review result for `product/records/spec/concepts/spec-format/index.md`

## Goal

Review the drafted spec format contract before any follow-up WORK / INV / DRMCP implementation work is created from it.

## Work

| review area | required check |
|---|---|
| PRODUCT scope | Confirm the contract is cross-app PRODUCT governance and not DRMCP implementation. |
| kind set | Confirm accepted/deferred spec kinds and revisit conditions are reasonable. |
| section matrix | Confirm required / recommended / prohibited sections are validation-friendly. |
| Overview+Topics | Confirm the rule avoids stale metadata and avoids unnecessary split work. |
| spec ID-as-ref derivation | Confirm `index.md` omission, underscore segment policy, no separate `ref`, and mismatch warning rules are safe. |
| parent grammar | Confirm interim grammar is stable and does not use path / filename / H1 title / derived topic ref. |
| stable spec ID-as-ref boundary | Confirm stable `spec:` IDs remain canonical and derived topic refs are deferred. |
| migration safety | Confirm existing specs are not migrated in this task. |

## Done condition

| item | done when |
|---|---|
| review completed | Codex / Opus / user review result exists. |
| blocking issues | Any blocking issue is reflected back into TASK-SPEC-001-01 or the spec draft. |
| user approval | User approves the format contract or explicitly sends it back for revision. |

## Verification

- Review output includes PASS / NEEDS_REVISION / FAIL.
- Review output references the exact spec draft path.

## Evidence

- Review result: OK with minor fixes.
- Reviewer: Codex.
- Files reviewed:
  - `v01/records/prompt_chappy.md`
  - `AGENTS.md`
  - `product/records/spec/concepts/spec-format/index.md`
  - `product/records/work-items/spec/PRODUCT-WORK-SPEC-001-spec-format-contract-and-follow-up-split.md`
  - `product/records/tasks/spec/PRODUCT-TASK-SPEC-001-01-draft-spec-format-contract.md`
  - `product/records/tasks/spec/PRODUCT-TASK-SPEC-001-02-review-spec-format-contract.md`
  - `product/records/tasks/spec/PRODUCT-TASK-SPEC-001-03-create-follow-up-split.md`
  - `product/records/tasks/spec/PRODUCT-TASK-SPEC-001-04-close-work-spec-001.md`
- Minor fix 1 applied: TASK-001-02 review checklist now explicitly checks `index.md` omission, underscore segment policy, no separate `ref`, and mismatch warning rules.
- Minor fix 2 applied: `spec-format/index.md` mismatch wording now requires warning when visible `id` does not match the path-derived default suggestion.
- Non-issues confirmed:
  - `spec-format/index.md` has no YAML front matter.
  - H1-adjacent metadata is `id/status/date/parent`.
  - Separate `- **ref**:` marker is not required or present.
  - Spec ID examples use underscore segments, including related specs.
  - PRODUCT-WORK-SPEC-001 Required Section Matrix includes H1-adjacent `id/status/date/parent`.
  - Parent grammar is `segment := [a-z0-9][a-z0-9_]*`.
  - Follow-up split includes PRODUCT-WORK-SPEC-002/003/004/005, PRODUCT-INV-SPEC-002, DRMCP-WORK-SPEC-001/002.
  - TASK-001-01/03/04 are aligned with current scope.
- Machine check reported 44 pass / 0 fail.
- `git diff -- <targets>` produced no output because the target files are currently untracked.
- `git status --short -- <targets>` reported the target files as `??`.
- Out-of-scope confirmation:
  - `v01/records/**` was not changed.
  - DRMCP implementation code was not changed.
  - Existing spec bulk migration was not performed.

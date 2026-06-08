# V01-TASK-MCP-016-01: Define required narrative section close policy

- **id**: V01-TASK-MCP-016-01
- **status**: done
- **date**: 2026-06-03
- **work_item**: V01-WORK-MCP-016
- **source_requirement**: V01-REQ-MCP-017
- **estimate**: 0.5d
- **depends_on**:
- **outputs**:
  - Close-state required narrative section policy matrix for WORK/TASK/REQ

## Goal

Define the required narrative section non-empty policy for workflow artifacts when they enter close states.

## Work

- Inspect existing authoring guides for requirement, work item, and task required sections.
- Decide the minimum close-state matrix for:
  - `WORK status: done`
  - `TASK status: done`
  - `REQ status: accepted`
- Classify any ambiguous section requirements as either enforced, deferred, or intentionally informational.
- Record the resulting matrix in this task evidence so implementation tasks do not guess policy.

## Done condition

This task is done when the close-state required narrative section matrix is explicit and ready to drive spec and implementation updates.

## Verification

- Compare the matrix against current authoring guidance for requirement, work item, and task artifacts.
- Confirm the matrix does not introduce content quality judgment beyond non-empty section body validation.

## Evidence
Policy review completed on 2026-06-03 using Opus/reviewer output and current authoring guidance.

Review verdict: OK with minor fixes.

Policy matrix:

| Artifact | Status gate | Required non-empty sections | Not enforced | Basis |
|---|---|---|---|---|
| WORK | `done` | `Goal`, `Boundary`, `Evidence` | `Impact Scope`, `Task flow`, `Task Candidates`, `Completion Condition` | Newly enforced by `V01-REQ-MCP-017`; the accepted requirement explicitly requires at least this guard for `WORK status: done`. |
| TASK | `done` | `Goal`, `Work`, `Done condition`, `Verification`, `Evidence` | none | Already grounded in `task-authoring` required section policy. |
| REQ | `accepted` | `Requirement`, `Required Outcome` | `Evidence`, `Boundary`, `Explicitly Excluded Scope` | Adoption-readiness gate, not close/completion gate. `Requirement` is the requirement statement; `Required Outcome` is the satisfaction contract consumed by work items. |

REQ accepted policy decision:

- `REQ status: accepted` is not a close/completion state. It means the requirement is adopted; implementation/spec completion continues through work items.
- Therefore the validation gate is adoption-readiness, not completion evidence.
- `Requirement` and `Required Outcome` are required non-empty.
- `Evidence`, `Boundary`, and `Explicitly Excluded Scope` are intentionally not enforced for `REQ accepted`.
- `Explicitly Excluded Scope` must not be required non-empty because a requirement can legitimately have no excluded scope; forcing filler text would exceed non-empty validation and drift toward content-quality judgment.

Diagnostic policy for follow-up spec/implementation tasks:

- Use separate categories:
  - `missing_required_section`: required heading is absent for the gated status.
  - `empty_required_section`: required heading exists but its body is empty or whitespace-only.
- Both categories should be `error` severity because warnings would not close the validation gap identified by `V01-REQ-MCP-017`.
- Diagnostic payload must include at least standard fields plus `record_id`, `path`, `section`, `status`, `severity`, and actionable `message`.
- Missing-section and empty-section should remain separate, following the existing metadata precedent of distinguishing missing vs empty required metadata.

Non-empty definition:

- Exclude the heading line.
- Trim surrounding whitespace from the section body.
- The section body is non-empty when it contains at least one non-whitespace character.
- Whitespace-only body is empty.
- Literal placeholder text such as `Pending` passes this non-empty guard. This is an accepted limitation because `V01-REQ-MCP-017` explicitly excludes narrative quality or sufficiency judgment.

Catalog target correction for `V01-TASK-MCP-016-02`:

- The design-record validation diagnostic catalog target is `docs/spec/design-records-mcp/schema.md`, not `docs/spec/diagnostics.md`.
- `docs/spec/diagnostics.md` belongs to the brewprint DSL semantic diagnostics catalog and should not receive this Design Records MCP diagnostic unless a later spec review finds a separate reason.

Verification result:

- Compared the matrix against `requirement-authoring`, `work-item-authoring`, and `task-authoring`.
- Confirmed that TASK required sections are already explicit in `task-authoring`.
- Confirmed that WORK/REQ narrative non-empty policy is newly enforced by `V01-REQ-MCP-017` and should be cited as such, not misrepresented as pre-existing guide policy.
- Confirmed the policy stays within non-empty section body validation and does not introduce narrative quality assessment.

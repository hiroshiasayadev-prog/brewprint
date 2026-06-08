# V01-TASK-DATA-006-04: Close and follow-up split

- **id**: V01-TASK-DATA-006-04
- **status**: done
- **date**: 2026-06-01
- **work_item**: V01-WORK-DATA-006
- **source_requirement**: V01-REQ-DATA-002
- **estimate**: 0.5d
- **depends_on**:
  - V01-TASK-DATA-006-03
- **outputs**:
  - V01-WORK-DATA-006 close evidence
  - Remaining deferred helper-shape debt classification
  - V01-WORK-DATA-006 status synchronized to done

## Goal

Close V01-WORK-DATA-006 after the selected UC-002 helper-shape migration and verification review are complete, while keeping remaining deferred debt outside this work item.

This task does not perform additional YAML migration, fixture regeneration, render update, implementation change, or successor work item creation.

## Work

- Review V01-WORK-DATA-006 and V01-TASK-DATA-006-01 / V01-TASK-DATA-006-02 / V01-TASK-DATA-006-03.
- Record close evidence for the selected migration set.
- Classify remaining deferred helper-shape and notes-retreat debt as outside V01-WORK-DATA-006.
- Synchronize V01-WORK-DATA-006 metadata status and task relation.
- Preserve unrelated dirty and untracked files.

## Done condition

- V01-TASK-DATA-006-04 records the close evidence from the completed migration and verification tasks.
- V01-WORK-DATA-006 includes V01-TASK-DATA-006-04 in task metadata.
- V01-WORK-DATA-006 status is synchronized to `done`.
- Remaining deferred debt is not treated as part of the selected 7-candidate migration.
- No unrelated dirty or untracked files are edited, staged, reverted, or committed.

## Verification

- Run `git status --short` before editing to identify unrelated dirty / untracked files.
- Run `git log --oneline -1` to record current HEAD context.
- Validate task records through Design Records MCP.
- Validate work item records through Design Records MCP.

## Evidence

Completed on 2026-06-01.

Initial repository check:

- `git status --short` showed pre-existing unrelated modified and untracked files, including design-records implementation/spec files, UC-002 render/YAML outputs from the completed migration, and other workflow artifacts. These files were preserved.
- `git log --oneline -1` returned `09ba38b docs(data): triage M15 deferred follow-ups`.

Close evidence:

- V01-TASK-DATA-006-01 selected 7 UC-002 model-file response helper-shape candidates for migration.
- V01-TASK-DATA-006-02 migrated the selected candidates as same-file private helper models.
- The owning response models remain public.
- The selected helpers were not cut out to standalone public model files.
- UC-002 task-file `query_result:any` patterns remain unchanged.
- Tagged union / discriminator payloads, identity semantics, optional semantics, literal constraints, enum / vocabulary constraints, and other notes-retreat debt remain out of scope for V01-WORK-DATA-006.
- V01-TASK-DATA-006-03 review verdict was OK to proceed.
- Current renderer evidence for UC-002 is `rendered 40 file(s)`.
- Temp render output and canonical render output file lists matched.
- Temp render output and canonical render output SHA-256 hashes matched.
- `go run ./cmd/brewprint validate --yaml-root docs\uc\002-brewprint-self-hosting\yaml` passed with `ok`.
- `go test ./internal/resolve ./internal/render/model ./internal/render/project` passed.
- Design Records MCP validation passed for task records.
- Design Records MCP validation passed for work item records.

Remaining deferred helper-shape classification:

- The 8 UC-002 MCP task-file `query_result:any` patterns remain deferred because task `params[].model` must not reference same-file task-file private helper models under the accepted DATA-004 / V01-REQ-DATA-003 policy.
- No public model files were created for those task-file result shapes in V01-WORK-DATA-006.
- The selected model-file response-local migration set is complete; no selected candidate remains open inside V01-WORK-DATA-006.
- Broader tagged union / discriminator / identity / optional / literal / enum debt remains outside V01-WORK-DATA-006 and must not be folded into this close task.


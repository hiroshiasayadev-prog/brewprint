# WORK-DATA-006: Migrate UC-002 helper-shape candidates

- **id**: WORK-DATA-006
- **status**: done
- **date**: 2026-06-01
- **source_requirement**: REQ-DATA-002
- **impact_refs**:
  - REQ-DATA-002
  - REQ-DATA-003
  - WORK-DATA-002
  - WORK-DATA-003
  - WORK-DATA-004
  - TASK-DATA-003-04
  - ADR-070
  - ADR-071
  - ADR-075
- **tasks**:
  - TASK-DATA-006-01
  - TASK-DATA-006-02
  - TASK-DATA-006-03
  - TASK-DATA-006-04

## Goal

Migrate the UC-002 helper-shape candidates that were classified by WORK-DATA-003 and left out of WORK-DATA-004 after the required render and signature-policy blockers were resolved.

WORK-DATA-003 made model-file render available and classified the UC-002 model response helper candidates. WORK-DATA-004 resolved task-file private helper params / returns exposure policy. This work item owns the next migration step without reopening either work item.

## Boundary

### Included

- Review TASK-DATA-003-04 candidate classification as the migration input.
- Migrate selected UC-002 model response helper-shape candidates where the public response model can remain public and nested response-local shapes can become same-file private helper models.
- Re-evaluate the WORK-DATA-004 wait candidates after the params / returns policy is now enforced.
- For task-file candidates blocked by `params[].model` policy, decide whether to:
  - keep `any` / note-based shape,
  - use public model files,
  - migrate only returns-side helper shapes,
  - or defer to a later design.
- Update UC-002 YAML, fixtures, golden render outputs, and verification evidence only for the selected migration set.
- Preserve model-file render and task-file signature policy as already implemented capabilities.

### Candidate input from TASK-DATA-003-04

Model-file helper migration candidates:

- N-005: `analyze_impact_response.impacts`
- N-006: `analyze_impact_response.coverage`
- N-014: `get_reference_tree_response.nodes`
- N-015: `get_reference_tree_response.edges`
- N-023: `get_source_response.snippet`
- N-029: `list_endpoints_response.tables`
- N-033: `list_objects_response.objects`

Task-file policy-blocked candidates to re-evaluate:

- UC-002 MCP task files where `query_service` returns `model: any` and `build_response.params[].model` consumes `query_result` as `any`.
- These candidates must not be migrated into task-file private helper params, because WORK-DATA-004 made `params[].model` references to task-file private helper models invalid.

### Excluded

- Reopening WORK-DATA-002 task-file helper minimum.
- Reopening WORK-DATA-003 model-file render implementation or boundary.
- Reopening WORK-DATA-004 signature policy implementation.
- Tagged union / ADR-073 discriminator and variants rendering.
- ADR-074 DAG asset TypeRef hint.
- ADR-078 / ADR-079 / ADR-080 MCP semantic identity or helper exposure schema.
- Broad UC-002 notes-retreat cleanup outside the classified helper-shape candidates.
- Enum / range / default / support-matrix constraints not required for the selected helper-shape migration.
- Request-side or generic container debt that is not a response-local helper-shape migration.
- M15 / v1.1.0-spec reopening.

## Impact Scope

| layer | current state | handling in this work item |
|---|---|---|
| source requirement | REQ-DATA-002 accepted helper/model-render follow-up | Provides the migration lineage |
| model-file render | WORK-DATA-003 done | Use as available capability; do not reopen |
| task-file signature policy | WORK-DATA-004 done | Use as constraint; do not reopen |
| UC-002 candidates | TASK-DATA-003-04 classified | Use as input inventory |
| YAML / fixtures / golden outputs | UC-002 still has helper-shape note debt | Update only selected migration targets |

## Task Flow

```mermaid
flowchart TD
  T1["TASK-DATA-006-01: Migration set selection"]
  T2["TASK-DATA-006-02: UC-002 YAML migration"]
  T3["TASK-DATA-006-03: Fixture / golden verification"]
  T4["TASK-DATA-006-04: Close and follow-up split"]
  T1 --> T2 --> T3 --> T4
```

## Task Candidates

- `TASK-DATA-006-01`: Select the exact UC-002 helper-shape migration set from TASK-DATA-003-04 candidates.
- `TASK-DATA-006-02`: Migrate selected UC-002 YAML to model-file or allowed task-file helper shapes.
- `TASK-DATA-006-03`: Update fixtures / golden outputs and run verification.
- `TASK-DATA-006-04`: Close the migration work item and split any remaining notes-retreat follow-ups.

TASK-DATA-006-01 has been created to record the selected migration set.
TASK-DATA-006-02 has been created and completed to migrate the selected model-file response helper shapes and regenerate UC-002 render fixtures.
TASK-DATA-006-03 has been created and completed as a separate verification review of the selected migration boundary, temp/canonical render equality, and validation evidence.
TASK-DATA-006-04 has been created and completed to synchronize close evidence, deferred debt classification, and work item status.

## Completion Condition

This work item can be marked `done` when the selected UC-002 helper-shape migration set is implemented, verified with fixtures / golden outputs, and remaining candidates are explicitly classified as deferred, obsolete, or successor work without reopening WORK-DATA-002, WORK-DATA-003, WORK-DATA-004, tagged union, MCP identity, or M15 scope.

## Close Evidence

WORK-DATA-006 is closed as `done` for the selected UC-002 helper-shape migration.

Closed scope:

- TASK-DATA-006-01 selected 7 model-file response-local helper-shape candidates.
- TASK-DATA-006-02 migrated the selected candidates as same-file private helper models.
- TASK-DATA-006-03 verified the migration boundary and render evidence with an OK to proceed verdict.
- TASK-DATA-006-04 synchronized the close evidence and work item status.

Migration result:

- The owning response models remain public.
- The selected helpers were not cut out to standalone public model files.
- UC-002 task-file `query_result:any` patterns remain unchanged.
- Current UC-002 render evidence is `rendered 40 file(s)`.
- Temp and canonical render file lists matched.
- Temp and canonical render SHA-256 hashes matched.
- YAML validation, focused Go tests, and Design Records MCP task / work item validation passed.

Remaining deferred debt:

- The 8 UC-002 MCP task-file `query_result:any` patterns remain deferred under the accepted DATA-004 / REQ-DATA-003 task `params[].model` private-helper policy.
- Tagged union / discriminator payloads, identity semantics, optional semantics, literal constraints, enum / vocabulary constraints, and broader notes-retreat debt remain outside WORK-DATA-006.
- WORK-DATA-002, WORK-DATA-003, and WORK-DATA-004 were not reopened.
- No WORK-DATA-005, REQ-MCP-006, REQ-MCP-007, WORK-MCP-006, tagged union, ADR-074 DAG asset TypeRef hint, or MCP semantic identity work was mixed into this close.

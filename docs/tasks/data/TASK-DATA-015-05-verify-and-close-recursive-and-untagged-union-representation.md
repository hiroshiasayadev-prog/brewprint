# TASK-DATA-015-05: Verify and close recursive and untagged-union representation

- **id**: TASK-DATA-015-05
- **status**: done
- **date**: 2026-06-07
- **work_item**: WORK-DATA-015
- **source_requirement**: REQ-DATA-008
- **estimate**: 0.5d
- **depends_on**:
  - TASK-DATA-015-03
  - TASK-DATA-015-04
- **outputs**:
  - Verification summary for WORK-DATA-015
  - WORK-DATA-015 close evidence
  - REQ-DATA-008 close / accepted update input if appropriate

## Goal

Verify the selected recursive named reference / unsupported untagged union boundary and close WORK-DATA-015 when all selected follow-up work is complete.

## Work

- Review TASK-DATA-015-01 through TASK-DATA-015-04 results.
- Confirm recursive named model reference support is specified and either implemented or explicitly deferred with rationale.
- Confirm untagged union / general `oneOf` remains unsupported and no hidden broadening occurred.
- Confirm any UC-002 cleanup, fixture, or render follow-up is complete or explicitly deferred.
- Update WORK-DATA-015 close evidence when ready.
- Identify whether REQ-DATA-008 can move from `captured` to `accepted` or should remain open.

## Included Scope

- Verification and close synchronization.
- WORK-DATA-015 evidence update.
- Requirement status recommendation.

## Excluded Scope

- New implementation work.
- New UC-002 YAML migration.
- Golden regeneration.
- Untagged union / general `oneOf` support.

## Done condition

- WORK-DATA-015 is closed or explicitly left open with remaining blocker rationale.
- Verification evidence is recorded.
- Requirement follow-up status is identified.

## Verification

- Validate affected task / work item / requirement records after close updates.
- Record any repo-local test command outputs if implementation or YAML changed in prior tasks.

## Evidence

Completed on 2026-06-08.

### Scope reviewed

- `TASK-DATA-015-01`: contract boundary decision.
- `TASK-DATA-015-02`: TypeRef spec update.
- `TASK-DATA-015-03`: recursive named model reference runtime investigation.
- `TASK-DATA-015-04`: UC-002 recursive / untagged-like YAML cleanup.
- `TASK-DATA-015-06`: independent review of spec and task split.

### Verification summary

Recursive named model reference support is now specified and verified:

- `docs/spec/type-ref.md` documents recursive named model references as supported.
- Inline recursive shapes remain unsupported.
- `TASK-DATA-015-03` classified current runtime behavior as already-supported.
- Focused validation and render checks accepted `object_ref.parent: object_ref`.
- Renderer displayed the recursive field as a named reference and did not infinitely expand it.

Untagged union / general oneOf remains unsupported:

- `docs/spec/type-ref.md` explicitly rejects inline `union<...>`, `oneOf<...>`, `anyOf<...>`, and scalar union syntax.
- ADR-073 remains limited to tagged / discriminated union.
- No untagged union, general oneOf, anyOf, scalar union, or ADR-073 broadening was introduced.

UC-002 cleanup is complete for the selected surfaces:

- N-044: `object_ref.parent` migrated from `any` to recursive named model reference `object_ref`.
- N-009: `diagnostic.related` migrated from `any` to `list<diagnostic_related>`.
- `diagnostic_related` was added as a tagged union envelope with `kind` discriminator and `source_location` / `object_ref` variants.

### Commands recorded from completed tasks

- `go test ./internal/resolve ./internal/render/model ./cmd/brewprint`
  - Result: PASS.
- `go run ./cmd/brewprint validate --yaml-root docs/uc/002-brewprint-self-hosting/yaml --format json`
  - Result: PASS, `error_count: 0`, `warning_count: 0`.
- `go run ./cmd/brewprint render --yaml-root docs/uc/002-brewprint-self-hosting/yaml --out docs/uc/002-brewprint-self-hosting/render --clean`
  - Result: PASS, `rendered 47 files`.
- MCP `validate_records` for `TASK-DATA-015-04`: PASS.
- MCP `validate_records` for `WORK-DATA-015`: PASS.

### Close decision

`WORK-DATA-015` can be closed as done.

`REQ-DATA-008` can move from `captured` to `accepted` because the representation requirement was decided, specified, verified, and applied to the selected UC-002 surfaces without introducing untagged union support.

### Workspace note

The only dirty file reported before this close update was `tmp.py`, which is unrelated to WORK-DATA-015 and should not be included in this close commit.

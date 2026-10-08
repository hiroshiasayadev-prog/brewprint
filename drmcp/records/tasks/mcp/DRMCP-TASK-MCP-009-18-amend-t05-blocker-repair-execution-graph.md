# DRMCP-TASK-MCP-009-18: Amend T05 blocker repair execution graph

- **id**: DRMCP-TASK-MCP-009-18
- **status**: done
- **date**: 2026-06-30
- **work_item**: DRMCP-WORK-MCP-009
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 0.5d
- **depends_on**:
  - DRMCP-TASK-MCP-009-12
- **outputs**:
  - DRMCP-WORK-MCP-009
  - DRMCP-TASK-MCP-009-05
  - DRMCP-TASK-MCP-009-13
  - DRMCP-TASK-MCP-009-14
  - DRMCP-TASK-MCP-009-15
  - DRMCP-TASK-MCP-009-16
  - DRMCP-TASK-MCP-009-17
  - DRMCP-TASK-MCP-009-18
  - DRMCP-TASK-MCP-009-19
  - DRMCP-TASK-MCP-009-20
  - DRMCP-TASK-MCP-009-21
  - DRMCP-TASK-MCP-009-22
  - DRMCP-TASK-MCP-009-23

## Goal

Persist the post-release T13 and T14 blockers as a reviewed repair amendment.

Preserve accepted T10 through T12 history and the existing T15 through T17 verification and closure roles.

## Work

- Record that T13 completed its sole source edit but its mandatory adapter test exposed a stale test configuration outside its writer boundary.
- Record that T14 completed its three-file cleanup but package compilation exposed stale list and get calls in the protected T06 test file.
- Keep both executor changes intact and unstaged.
- Keep T13 and T14 blocked until the repair wave and downstream acceptance complete.
- Add T19 independent amendment review and T20 release synchronization.
- Add T21 as the sole writer for `drmcp/src/internal/designrecordsmcp/tools_call_test.go`.
- Add T22 as the sole temporary repair writer for `drmcp/src/internal/designrecords/resolve_reference_test.go`.
- Preserve T06 production ownership and resolver semantics.
- Limit the T22 ownership transfer to removal of retired list and get assertions from one named resolver test.
- Release T21 and T22 only after T19 returns `PASS` and T20 records release.
- Allow T21 and T22 to execute in parallel because their writable paths are disjoint.
- Add T23 as the sole pre-T15 repair Evidence and lifecycle synchronization owner.
- Change T15 to depend on T23 and verify the complete six-file correction and repair boundary.
- Keep T16 as the independent finding and mapping acceptance owner.
- Keep T17 as the sole lifecycle and Evidence synchronization owner.
- Do not modify production source, tests, fixtures, schemas, manifests, ADRs, Requirements, or Specifications.

### Amended execution graph

```text
accepted T10 graph authoring
  -> accepted T11 independent review
  -> accepted T12 release
  -> T13 source edit result BLOCKED by external stale test fixture
  -> T14 package cleanup result BLOCKED by protected stale resolver test
  -> T18 blocker repair graph amendment
  -> T19 independent amendment review
  -> T20 repair release synchronization
  -> T21 adapter current-root test fixture repair ┐
  -> T22 resolver stale API test repair          ├─ parallel
                                                 ┘
  -> T23 repair Evidence synchronization
  -> existing T15 aggregate verification
  -> existing T16 independent finding closure
  -> existing T17 correction closure synchronization
```

### Model routing

| Task | model route | reason |
|---|---|---|
| T18 | ChatGPT coordinator | Persistent graph and cross-lane ownership amendment. |
| T19 | Claude Code Opus | Independent ownership, dependency, and contract review. |
| T20 | Claude Code Haiku | Mechanical review acceptance and repair release. |
| T21 | Claude Code Haiku | One exact test fixture replacement with one focused command. |
| T22 | Claude Code Haiku | One exact stale assertion removal with frozen resolver behavior. |
| T23 | Claude Code Haiku | Mechanical persistence of two accepted executor results. |
| T15 | Claude Code Haiku | Existing deterministic aggregate verification. |
| T16 | Claude Code Opus | Existing independent finding and mapping acceptance. |
| T17 | Claude Code Haiku | Existing mechanical closure synchronization. |

### Writer ownership

| path or artifact | sole writer |
|---|---|
| W009, T05, T13 through T23 graph amendment | T18 |
| Amendment verdict | T19, read-only except its own Evidence during later synchronization |
| Repair release state | T20 |
| `drmcp/src/internal/designrecordsmcp/tools_call_test.go` | T21 |
| `drmcp/src/internal/designrecords/resolve_reference_test.go` | T22 for this repair wave only |
| T21 and T22 focused Evidence and lifecycle | T23 |
| Aggregate verification | T15, read-only |
| Finding and mapping acceptance | T16, read-only |
| T05 correction lifecycle and Evidence | T17 |

T06 retains ownership of `resolver.go`, resolver behavior, and the surviving semantic-reference assertions in `resolve_reference_test.go`.
T22 does not reopen T06 implementation or review.

## Done condition

- T18 through T23 exist under W009 with exact IDs and filenames.
- T10 through T12 accepted history remains unchanged.
- T13 and T14 blocker results are recorded without attributing executor fault.
- T21 and T22 have disjoint exact writer boundaries.
- T22 has a bounded temporary ownership transfer that preserves T06 resolver semantics.
- T23 persists both repair results before aggregate verification.
- T15 depends on T23 and owns the six-file aggregate gate.
- T16 and T17 retain their existing primary responsibilities.
- Model routing, writer ownership, dependencies, focused tests, aggregate verification, review, release, and closure ownership are explicit.
- No production source or Go test is changed during graph authoring.
- The amendment is ready for independent review.

## Verification

- Confirm W009 lists T01 through T23 exactly once.
- Confirm T19 depends on T18 and T20 depends on T19.
- Confirm T21 and T22 depend on T20.
- Confirm T23 depends on T21 and T22.
- Confirm T15 depends on T23.
- Confirm T16 still depends on T15 and T17 still depends on T16.
- Confirm T21 and T22 writer paths do not overlap.
- Confirm no new Task owns `resolver.go`.
- Confirm T06 remains unchanged.
- Inspect only the amended Design Record paths with scoped Git and whitespace checks.

## Evidence

- T13 source result: `idRangeSchema` removal complete; formatter, symbol absence, and scoped whitespace checks passed.
- T13 blocker: `TestToolsCallToolErrors/get_authoring_guidance_unknown_id` reached `normalizeConfig` through the authoring-guidance path while using `designrecords.Config{Root: "."}` with no current roots.
- T13 blocker path: `drmcp/src/internal/designrecordsmcp/tools_call_test.go`, outside the T13 writer boundary.
- T14 source and owned-test result: obsolete list and get symbols, legacy tests, and compact `kind` assignment were removed; formatter and scoped presence or absence checks passed.
- T14 blocker: package compilation found retired `ListRecords` and `GetRecord` calls in `TestResolveReferenceUsesSemanticRefsFromNonRecordSpec`.
- T14 blocker path: `drmcp/src/internal/designrecords/resolve_reference_test.go`, outside the T14 writer boundary and inside the accepted T06 lane.
- Both blockers result from incomplete released executor contracts, not executor implementation mistakes.
- The minimum repair uses two test-only writers and preserves production source and accepted resolver behavior.
- T10 through T12 remain accepted historical gates and are not reopened.
- DRMCP authoring transaction tools were unavailable. Filesystem authoring was used under `drmcp/records`.
- No production source, Go test, fixture, schema, manifest, ADR, Requirement, or Specification changed during T18.
- T19 independently reviewed this amendment and returned `PASS` with no blocking, major, or minor finding.
- T20 synchronized that acceptance and released only T21 and T22 for parallel execution.
- T23, T15, T16, and T17 remain dependency-gated.

# PRODUCT-TASK-SPEC-013-03: Define source-authoring warning handling contract

- **id**: PRODUCT-TASK-SPEC-013-03
- **status**: done
- **date**: 2026-06-26
- **work_item**: PRODUCT-WORK-SPEC-013
- **source_requirement**: PRODUCT-REQ-SPEC-003
- **estimate**: 0.5d
- **depends_on**:
  - PRODUCT-TASK-SPEC-013-01
  - PRODUCT-TASK-SPEC-013-02
- **outputs**:
  - Source-authoring warning handling contract recorded in this Task Evidence
  - T01 and DRMCP consumer requirement synchronized to the non-persistent warning-output boundary

## Goal

Define the minimum handling boundary for producer-reported semantic warnings.

Separate warning output from persistent evidence management.

Reserve manually created source-authoring Requirements for cases where a human reviewer decides that tracked correction is required.

Keep warning closure outside the dependency chain for package generation and subsequent tasks.

## Work

- Record that the producer emits semantic warnings during generation or check execution.
- Record that the producer does not own persistent warning evidence.
- Record that the producer does not own warning disposition, lifecycle management, owner assignment, or a warning routing registry.
- Record that there is no obligation to save warning output to T05, T06, or T08 Evidence.
- Record that a human reviewer decides whether any warning requires tracked source correction.
- Record that a source-authoring Requirement is manually created only when tracked correction is required.
- Record that the relevant warning output may be copied into the Requirement Evidence.
- A new record is not required when the human reviewer determines that tracked correction is unnecessary, or when the same finding is already covered by an existing Requirement. Repeated reports of the same finding do not require duplicate records.
- Record that source authoring owns semantic correction.
- Record that warnings do not block generation, distribution, or T04–T08.
- Confirm the boundary between semantic warnings and operational generation failures as established by T02.
- Do not define diagnostic codes, JSON schemas, CLI formatting, exit codes, scanner algorithms, or consumer diagnostics.

## Done condition

The following are clearly recorded in Task Evidence:

- Producer responsibility is bounded by warning emission.
- Persistent warning evidence is not required.
- Human review and optional manual Requirement creation are defined.
- Warning closure is not a generation dependency.
- No contradictions with T01, T02, PRODUCT-REQ-SPEC-003, PRODUCT-WORK-SPEC-013, or DRMCP-REQ-MCP-003.
- Independent review accepts the warning-handling contract without blocking or major findings.

## Verification

- No warning registry introduced.
- No warning lifecycle state introduced.
- No warning evidence field schema introduced.
- No obligation to save warning output introduced.
- No per-warning Task, Work Item, or Requirement creation required.
- Producer is not assigned semantic correction ownership.
- Ordinary semantic warnings are not generation blockers.
- Operational failure boundary from T02 is unchanged.
- No new Task for T04 or later created.
- No unrelated files changed.

## Evidence

### Current normative summary

| concern | contract |
| --- | --- |
| producer | emits warnings during generation or check execution |
| persistent producer-owned warning evidence | not required |
| human review | decides whether tracked source correction is required |
| follow-up | manually create a source-authoring Requirement only when required |
| warning output | may be copied into the Requirement Evidence |
| follow-up threshold | create a new source-authoring Requirement only when tracked correction is required and no existing Requirement already covers the finding |
| generation | continues while semantic warnings or source-authoring follow-ups remain open |

### Independent review

- **verdict**: PASS
- **reviewed contract**:
  - producer responsibility is limited to warning emission during generation or check execution;
  - the producer does not own persistent warning evidence, warning disposition, lifecycle management, owner assignment, or automatic routing;
  - warning output does not have to be retained in T05, T06, or T08 Evidence;
  - a human reviewer decides whether tracked source correction is required;
  - a source-authoring Requirement is manually created only when tracked correction is required and no existing Requirement already covers the finding;
  - repeated reports of the same finding do not require duplicate records;
  - source authoring owns semantic correction;
  - semantic warnings and open source-authoring follow-ups do not block generation, distribution, or T04–T08;
  - operational generation failures remain the producer errors defined by T02.
- **previous findings**:
  - ambiguous `duplicate` wording: closed;
  - stale `warning-evidence tooling` wording: closed;
  - stale no-Task statement: closed;
  - stale T01 recording wording: closed;
  - Done condition lifecycle contradiction: closed;
  - Work Item follow-up threshold mismatch: closed.
- **finding result**: No blocking, major, or minor findings.
- **closure**: The Done condition is satisfied. This Task is `done`, and T04 may consume the accepted warning-handling contract.

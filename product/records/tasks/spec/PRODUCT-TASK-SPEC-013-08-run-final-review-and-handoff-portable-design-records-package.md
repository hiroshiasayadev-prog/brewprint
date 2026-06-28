# PRODUCT-TASK-SPEC-013-08: Run final review and hand off portable Design Records package

- **id**: PRODUCT-TASK-SPEC-013-08
- **status**: done
- **date**: 2026-06-28
- **work_item**: PRODUCT-WORK-SPEC-013
- **source_requirement**: PRODUCT-REQ-SPEC-003
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-013-06
  - PRODUCT-TASK-SPEC-013-07
- **outputs**:
  - PRODUCT-WORK-SPEC-013 final release evidence and closure preparation
  - DRMCP-WORK-MCP-002 producer handoff evidence
  - independent final review prompt

## Goal

Complete the producer-side final contract review and release-evidence consolidation.

Prepare the accepted package boundary for DRMCP consumer work without implementing or defining consumer runtime behavior.

## Work

- Review the released producer contract against `PRODUCT-REQ-SPEC-003` and accepted PRODUCT ADRs.
- Confirm the authoritative source, generated destination, logical roots, fixed namespace, and canonical prefix rewrite.
- Confirm whole-tree rebuild, warning-only success, operational-failure behavior, and generated-artifact authority boundaries.
- Confirm the generator does not depend on process working directory, host app registry, DRMCP runtime, network access, or third-party packages.
- Review duplicate, unresolved, external, unrewritten, and source-authoring boundary warning behavior.
- Confirm warnings do not authorize repair, filtering, deletion, generalization, or reinterpretation.
- Consolidate accepted T06 and T07 generation, test, release, ignore, and independent-review evidence.
- Confirm the reciprocal relation from `PRODUCT-REQ-SPEC-003` to `PRODUCT-WORK-SPEC-013`.
- Prepare producer handoff evidence for `DRMCP-REQ-MCP-003` through `DRMCP-WORK-MCP-002`.
- Record the consumer-owned remaining scope and the absence of a consumer P0 Work Item.
- Record residual limitations without treating them as producer release blockers.
- Obtain an independent final review and close every blocking, major, and minor finding before closure.

## Done condition

- The final producer contract review finds no unresolved conflict with accepted PRODUCT authority.
- Exact release evidence is recorded without copying all 79 warning messages.
- `PRODUCT-REQ-SPEC-003` lists `PRODUCT-WORK-SPEC-013` in `work_items`.
- The DRMCP handoff identifies accepted producer inputs, outputs, boundaries, evidence, and residual limitations.
- The next consumer-side action is identified without creating or implementing a consumer Work Item.
- An independent final review returns `PASS`.
- Every blocking, major, and minor final-review finding is closed.
- This Task and `PRODUCT-WORK-SPEC-013` remain `in_progress` until the independent review passes.
- After review closure, this Task and `PRODUCT-WORK-SPEC-013` are changed to `done` together with final closure evidence.

## Verification

- Compare the producer contract with `PRODUCT-REQ-SPEC-003`, `PRODUCT-ADR-SPEC-001`, `PRODUCT-ADR-SPEC-002`, and `PRODUCT-ADR-SPEC-003`.
- Inspect the generator, release fixtures, `scripts/verify.bat`, and `.gitignore` against the accepted boundary.
- Confirm the five required semantic warning classes remain non-blocking.
- Confirm generated duplicate and unresolved defects remain present and localized.
- Confirm the package is derived and ignored under `bin/design-records/`.
- Confirm `PRODUCT-REQ-SPEC-003` has the required reciprocal Work Item relation.
- Confirm no Work Item sourced from `DRMCP-REQ-MCP-003` currently exists before recording the next consumer action.
- Use accepted T07 command evidence when repository-local command execution is unavailable.
- Inspect only the defined T08 Git scope and run scoped whitespace checks through `git.inspect_worktree`.
- Submit the prepared independent-review prompt to a separate reviewer.

## Evidence

### Authoring and collision checks

- Design Records MCP authoring tools were unavailable in the current tool surface. Known-path filesystem authoring was used under the current agent-authoring fallback rule.
- The expected file name did not exist in `product/records/tasks/spec/` before creation.
- A scoped search under `product/records/` found no existing `PRODUCT-TASK-SPEC-013-08` ID before creation.

### Producer contract review

The current producer implementation matches the accepted contract:

| concern | reviewed result |
|---|---|
| authoritative source | `product/records/spec/design-records/` |
| generated destination | `bin/design-records/`, equivalent to `<exe-dir>/design-records/` for the repository-local executable |
| logical package root | `design-records/` |
| package spec root | `spec:design_records` |
| guidance root | `spec:design_records.authoring_standards` |
| semantic transformation | Exact canonical prefix rewrite from `spec:product.design_records` to `spec:design_records` only. |
| preserved content | Ordinary prose, public IDs, physical paths, external refs, noncanonical lookalikes, unrelated line endings, and non-UTF-8 bytes remain unchanged. |
| generation model | Deterministic whole-tree rebuild through a temporary tree and destination replacement. |
| warning-only result | Successful generation returns exit `0`. |
| operational failure | Failed operational generation returns exit `1`. |
| semantic authority | The generated tree is derived and has no independent semantic authority. |
| Git boundary | `bin/design-records/` remains covered by the existing `bin/` ignore rule. |
| independence | The public generator resolves from its script location and has no host registry, DRMCP runtime, network, environment-root, current-working-directory, or third-party dependency. |

No new producer-contract conflict was found during this review.
The independent final re-review later confirmed this assessment and closed the only T08 minor finding.

### Warning and defect boundary

The generator reports these semantic warning classes:

- `duplicate canonical ref`;
- `unresolved internal ref`;
- `external canonical ref`;
- `unrewritten source ref`;
- `source-authoring boundary finding`.

All five classes are non-blocking.
The producer does not repair, filter, delete, generalize, or reinterpret warning-producing content.
Duplicate and unresolved defects remain in the generated package.
Source-authoring warning closure is not a package-release prerequisite.
Warning output is not a producer-owned persistent artifact.
A human may copy relevant warning output into a separately owned Requirement only when tracked source correction is required.

### Accepted release evidence

The following results are accepted from T06 and the final T07 re-review:

| evidence | result |
|---|---|
| generation command | `python -X utf8 product\src\tools\generate_design_records_package.py` |
| source path | `product/records/spec/design-records/` |
| destination path | `bin/design-records/` |
| generated files | 34 |
| semantic warnings | 79, non-blocking |
| generator exit | `0` |
| test command | `python -X utf8 -m unittest discover -s product\tests\tools -p "test_*.py" -v` |
| tests | 35 passed |
| test exit | `0` |
| repository verification | `scripts\verify.bat` exited `0` and printed final `OK` |
| generated artifact ignore | `git check-ignore -v bin/design-records/index.md` exited `0` |
| scoped whitespace | pass |
| T07 independent re-review | `PASS`; F-MIN-01 and F-MIN-02 closed; no blocking, major, minor, or advisory findings |

Repository-local commands were not rerun during T08 because the current tool surface has no repository command-execution capability.
The table records accepted T07 evidence and does not claim a new execution.

### Requirement relation

`PRODUCT-REQ-SPEC-003` already contains:

```text
work_items:
  - PRODUCT-WORK-SPEC-013
```

The reciprocal Requirement relation is correct.
No Requirement rewrite is required.

### DRMCP handoff

Producer authorities:

- Requirement: `PRODUCT-REQ-SPEC-003`.
- Work Item: `PRODUCT-WORK-SPEC-013`.

Accepted producer interface:

- source root: `product/records/spec/design-records/`;
- generated package location: `bin/design-records/`;
- logical package root: `design-records/`;
- fixed package namespace root: `spec:design_records`;
- guidance root: `spec:design_records.authoring_standards`;
- transformation: canonical prefix rewrite only;
- warnings: five semantic classes, all non-blocking;
- operational failure: exit `1` and no accepted partial publication;
- generation command: `python -X utf8 product\src\tools\generate_design_records_package.py`;
- accepted portability evidence: repeated byte-stable generation, repository-external current-working-directory execution, irrelevant host-state environment values, standard-library-only imports, and fixed script-location root resolution.

Consumer-owned remaining scope under `DRMCP-REQ-MCP-003`:

- package-root configuration;
- loader behavior;
- package availability handling;
- recursive Markdown discovery;
- localized indexing and diagnostics;
- duplicate, unresolved, unreadable, unparseable, and unrewritten-ref consumer behavior;
- guidance list and exact-get projection;
- authoring and validation integration;
- proposal reproducibility and staleness guards;
- package-dependent capability exposure.

The producer does not own or define those consumer behaviors.
A scoped search found no existing Work Item sourced from `DRMCP-REQ-MCP-003`.
The next consumer-side action is P0 Work Item creation through the `DRMCP-WORK-MCP-002` T03 milestone sequence after producer closure readiness is accepted.
No consumer implementation Work Item was created by T08.

### Residual limitations

- 79 semantic warnings remain and are not release blockers.
- No package manifest exists.
- No package version negotiation exists.
- No remote registry exists.
- No incremental generation exists.
- Concurrent generation against one destination is unsupported.
- Consumer loading, indexing, guidance projection, and authoring integration remain unimplemented.
- The generated tree is not manual semantic authority.
- Package source correction may require a different owner and separate work.

### Independent final review

- Initial independent review verdict: `NEEDS REVISION`.
- Initial finding: T08 F-MIN-01 identified stale parent Work Item wording that treated the not-yet-created DRMCP P0 Work Item as the direct handoff target and omitted minor-finding closure from the Work Item review gate.
- Correction: Phase H, Task Candidate T08, and the handoff Completion Condition now route producer evidence through `DRMCP-WORK-MCP-002` for T03 P0 Work Item creation.
- Correction: the parent Work Item review gate now requires no blocking, major, or minor findings.
- Independent re-review verdict: `PASS`.
- T08 F-MIN-01: `CLOSED`.
- T07 F-MIN-01 and F-MIN-02 remain `CLOSED`.
- Final finding summary: no blocking, major, or minor findings.
- Scoped whitespace result: `pass`; LF-to-CRLF conversion warnings were advisory only.
- Repository-local tests, generator, and `scripts\\verify.bat` were not rerun for the wording-only correction. Accepted T07 evidence remains the execution evidence.

### Final closure state

- `PRODUCT-TASK-SPEC-013-08`: `done`.
- `PRODUCT-WORK-SPEC-013`: `done`.
- Producer closure is complete.
- The next consumer-side action is P0 Work Item creation through `DRMCP-WORK-MCP-002` T03.
- Consumer implementation remains open and is not implied complete by producer closure.

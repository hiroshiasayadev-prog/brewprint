# PRODUCT-TASK-SPEC-025-06: Extract boundary-vocabulary entries from W001–W007

- **id**: PRODUCT-TASK-SPEC-025-06
- **status**: done
- **date**: 2026-07-03
- **work_item**: PRODUCT-WORK-SPEC-025
- **task_type**: investigation
- **estimate**: 1d
- **depends_on**:
  - PRODUCT-TASK-SPEC-025-02
- **outputs**:
  - PRODUCT-TASK-SPEC-025-06

## Goal

Extract boundary-vocabulary log entries from every `status: done` Task under `product/records/tasks/spec/` whose parent Work Item is in the `PRODUCT-WORK-SPEC-001` through `PRODUCT-WORK-SPEC-007` range (~35 Tasks), and record them directly in this Task's own `## Evidence`, under the Investigation-Task lightweight Evidence exception (PRODUCT-TASK-SPEC-025-02).

## Work

- Fix this Task's corpus at start: the exact set of `status: done` Task IDs under `product/records/tasks/spec/` whose `work_item` falls in `PRODUCT-WORK-SPEC-001`–`PRODUCT-WORK-SPEC-007`.
- Read each in-scope Task's Work and Done Condition sections. Some older Tasks have no `task_type` field; infer the responsibility cluster from title, Work, and Done Condition instead of skipping them.
- Find phrases where a responsibility is expressed without its owning verb (e.g. a synonym standing in for `author`, `review`, `coordinate`), or where a verb carries an unusual meaning in context (e.g. `Fix`).
- Record each finding as a 4-column entry: phrase (exact), source Task ID, vocabulary target (what responsibility or term it stands in for), effective meaning in that context.
- Record a scanned entry for every Task read, including Tasks with zero findings. Do not force findings.
- Do not rank by frequency.
- Do not decide canonical terms.
- Record the accepted user judgment and reason for using the lightweight exception (PRODUCT-TASK-SPEC-025-01 D-006) directly in `## Evidence`.
- Write only to this Task's own Evidence. Do not write to any sibling Task's file.

Responsibility clusters (classification target for findings, not a further split axis):

- decision / authoring / coordination (decide, define, fix, determine, materialize, project, encode, update, align, amend)
- review / verification / correction (review, inspect, verify, validate, check, re-check, confirm, assess, correct, repair, close finding)
- investigation / decision / design (investigate, research, analyze, compare, identify, conclude, recommend, select)
- graph / lifecycle / execution (coordinate, materialize, create, route, release, synchronize, close, execute, decompose, propagate)

This Task must not:

- author or edit `skills/task-boundary-vocabulary/` files directly;
- decide canonical term definitions;
- create a separate formal Investigation record (covered by the accepted exception);
- rewrite or reinterpret source Task content;
- write to any Task file other than this one.

## Done condition

- Every `status: done` Task in the `PRODUCT-WORK-SPEC-001`–`PRODUCT-WORK-SPEC-007` range at this Task's fixed cutoff has a scanned entry, including zero-finding Tasks.
- Every logged entry cites an exact, existing source Task ID.
- The accepted user judgment and reason for using the lightweight exception is recorded in this Task's own Evidence.

## Verification

- Confirm every entry cites a real, existing Task ID within the assigned range.
- Confirm no canonical term decision was made.
- Confirm the exception-use judgment and reason are recorded in Evidence.
- Confirm scanned-Task count matches the fixed corpus for this range.

## Evidence

### Execution basis

- Cutoff: the corpus was fixed on 2026-07-03 before source reading.
- Corpus query: `product/records/tasks/spec/PRODUCT-TASK-SPEC-001-*` through `PRODUCT-TASK-SPEC-007-*`, filtered to `status: done`.
- Fixed corpus: 35 Tasks. W002 and W003 had no matching Task files. Every inspected file in the range was `done` at cutoff.
- Source boundary: only each Task's `## Work` and `## Done condition` sections were used for phrase extraction.
- Exception judgment: PRODUCT-TASK-SPEC-025-01 D-006 records the user's accepted judgment that a formal Investigation record is disproportionate for this bounded log extraction.
- Exception reason: a formal Investigation record would reproduce the heavyweight route already rejected for this work. PRODUCT-TASK-SPEC-025-04 will conclude the log through downstream review and authoring.
- Canonicalization boundary: vocabulary targets below are classification candidates only. This Task did not select canonical terms or definitions.
- Repeated phrases are retained per source Task. No frequency ranking was performed.

### Finding log

| phrase (exact) | source Task ID | vocabulary target | effective meaning in that context |
|---|---|---|---|
| `Create or confirm` | PRODUCT-TASK-SPEC-001-01 | authoring / verification | Produce the target spec file when absent, or verify an existing file as the deliverable. |
| `Define` | PRODUCT-TASK-SPEC-001-01 | decision / authoring | Encode accepted normative clauses; the verb does not distinguish selecting a rule from writing it. |
| `Confirm` | PRODUCT-TASK-SPEC-001-02 | review | Judge semantic scope, safety, and reasonableness during independent review. |
| `reflected back into` | PRODUCT-TASK-SPEC-001-02 | correction routing | Propagate a blocking issue into the draft Task or artifact for repair. |
| `created or explicitly queued` | PRODUCT-TASK-SPEC-001-03 | coordination / materialization | Either persist a follow-up artifact now or record deferred downstream work with rationale. |
| `Mark completed tasks done with evidence` | PRODUCT-TASK-SPEC-001-04 | synchronization | Propagate accepted completion into Task lifecycle state and Evidence. |
| `Add evidence to PRODUCT-WORK-SPEC-001` | PRODUCT-TASK-SPEC-001-04 | synchronization | Update the parent Work Item with mechanically derived completion Evidence. |
| `Combine` | PRODUCT-TASK-SPEC-004-01 | decision synthesis | Merge two Investigation classifications into one proposed ownership decision. |
| `Recommend whether` | PRODUCT-TASK-SPEC-004-01 | recommendation / decision preparation | Propose a downstream scope choice without applying the update in the same Task. |
| `Confirm` | PRODUCT-TASK-SPEC-004-02 | review | Assess completeness, consistency, ordering, and downstream impact at the review gate. |
| `must-fix before 004-03, or defer` | PRODUCT-TASK-SPEC-004-02 | review finding disposition | Classify each finding by whether it blocks the next Task or may remain deferred. |
| `Apply all must-fix findings` | PRODUCT-TASK-SPEC-004-03 | correction | Repair every blocking finding recorded by the prior review. |
| `Record the accepted relocation plan as a future follow-up` | PRODUCT-TASK-SPEC-004-03 | coordination / synchronization | Persist an accepted downstream route so the decision does not dead-end. |
| `Mark PRODUCT-WORK-SPEC-004 status: done` | PRODUCT-TASK-SPEC-004-03 | synchronization | Close the parent Work Item after accepted artifacts and handoff state are present. |
| `Move 8 flat root-level files` | PRODUCT-TASK-SPEC-005-01 | authoring / migration preparation | Relocate source files into a read-only staging area before canonical authoring. |
| `Reduce top-level bpdsl/records/spec/overview.md Topics table from 16 rows to 3` | PRODUCT-TASK-SPEC-005-01 | authoring | Rewrite the canonical navigation projection to the accepted three-domain structure. |
| `is accounted for` | PRODUCT-TASK-SPEC-005-03 | review | Establish source-to-output completeness without requiring byte-identical text. |
| `what to check` | PRODUCT-TASK-SPEC-005-03 | review | Express review responsibility as checklist criteria rather than an owning verb. |
| `Apply all must-fix findings` | PRODUCT-TASK-SPEC-005-04 | correction | Repair findings produced by the preceding independent review. |
| `Add BPDSL DSL migration evidence entry to PRODUCT-WORK-SPEC-005` | PRODUCT-TASK-SPEC-005-04 | synchronization | Propagate the completed migration batch into parent Work Item Evidence. |
| `Copy all 12 files` | PRODUCT-TASK-SPEC-005-05 | materialization / authoring preparation | Create a frozen staging snapshot used by later authoring and review. |
| `Confirm bpdsl/old_mcp/ is byte-identical` | PRODUCT-TASK-SPEC-005-05 | verification | Objectively verify that the staging copy preserves the source corpus. |
| `Restructure numbered sections` | PRODUCT-TASK-SPEC-005-06 | authoring | Rewrite existing source content into the accepted Contract section shape. |
| `is accounted for` | PRODUCT-TASK-SPEC-005-07 | review | Judge migration completeness against the staged source corpus. |
| `Apply all must-fix findings` | PRODUCT-TASK-SPEC-005-08 | correction | Repair blocking review findings before cleanup. |
| `Add mcp/ migration evidence entry to PRODUCT-WORK-SPEC-005` | PRODUCT-TASK-SPEC-005-08 | synchronization | Propagate the completed mcp migration batch into parent Work Item Evidence. |
| `Copy the 7 existing .md files` | PRODUCT-TASK-SPEC-005-09 | materialization / authoring preparation | Create the read-only views staging corpus for later migration and review. |
| `Confirm bpdsl/old_views/ is byte-identical` | PRODUCT-TASK-SPEC-005-09 | verification | Objectively verify the staging snapshot against the source files. |
| `Restructure sections` | PRODUCT-TASK-SPEC-005-10 | authoring | Rewrite each view spec into the accepted format-contract section shape. |
| `identify which current sections map to` | PRODUCT-TASK-SPEC-005-10 | authoring analysis | Analyze source structure only to support the owned canonical rewrite. |
| `is accounted for` | PRODUCT-TASK-SPEC-005-11 | review | Judge source-to-output completeness for the views migration. |
| `Apply all must-fix findings` | PRODUCT-TASK-SPEC-005-12 | correction | Repair blocking review findings before removing staging data. |
| `Add views/ migration evidence entry to PRODUCT-WORK-SPEC-005` | PRODUCT-TASK-SPEC-005-12 | synchronization | Propagate batch and namespace-level completion into parent Work Item Evidence. |
| `Move all 3 existing files` | PRODUCT-TASK-SPEC-005-13 | authoring / migration preparation | Relocate the pre-migration DRMCP corpus into a read-only staging area. |
| `Create empty drmcp/records/spec/design-records-mcp/schema/ and drmcp/records/spec/design-records-mcp/tools/ directories` | PRODUCT-TASK-SPEC-005-13 | materialization / authoring preparation | Materialize the target directory structure without authoring target files yet. |
| `moves out` | PRODUCT-TASK-SPEC-005-14 | authoring / decomposition | Extract detailed content from the root overview into narrower canonical child specs. |
| `record them as deferred relocation candidates` | PRODUCT-TASK-SPEC-005-15 | review / routing | Preserve ownership findings for later execution without relocating content during review. |
| `Each finding is classified: must-fix before 005-16, or defer` | PRODUCT-TASK-SPEC-005-15 | review finding disposition | Assign each review finding a blocking or deferred route. |
| `Apply all must-fix findings` | PRODUCT-TASK-SPEC-005-16 | correction | Repair every blocking review finding before finalizing migration. |
| `Carry forward the PRODUCT-TASK-SPEC-005-15 deferred relocation candidates list` | PRODUCT-TASK-SPEC-005-16 | synchronization | Propagate unresolved but accepted follow-up state into parent Work Item Evidence. |
| `Relocate` | PRODUCT-TASK-SPEC-005-17 | authoring / migration | Move normative namespace semantics to the PRODUCT-owned canonical specification. |
| `Replace DRMCP file body with pointer` | PRODUCT-TASK-SPEC-005-17 | authoring | Remove duplicated normative content and leave a canonical cross-reference. |
| `Relocate` | PRODUCT-TASK-SPEC-005-18 | authoring / migration | Move repository discovery conventions into the PRODUCT-owned layout specification. |
| `DRMCP file trimmed` | PRODUCT-TASK-SPEC-005-18 | correction / authoring | Remove relocated clauses while retaining the DRMCP-owned filter behavior and pointer. |
| `Trim` | PRODUCT-TASK-SPEC-005-19 | correction / authoring | Remove duplicated resolver semantics from the DRMCP spec. |
| `Add drift guards` | PRODUCT-TASK-SPEC-005-19 | authoring | Insert ownership-boundary notes into the PRODUCT specification. |
| `Apply all must-fix findings` | PRODUCT-TASK-SPEC-005-21 | correction | Repair blocking findings from the Phase 2 review. |
| `Add Phase 2 relocation evidence entry` | PRODUCT-TASK-SPEC-005-21 | synchronization | Propagate the accepted relocation result into Work Item Evidence. |
| `evaluate closing WORK-005` | PRODUCT-TASK-SPEC-005-21 | lifecycle judgment / synchronization | Determine closure readiness from dependency state before updating Work Item lifecycle. |
| `Check required/prohibited sections` | PRODUCT-TASK-SPEC-006-01 | implementation | Implement validator behavior that checks section presence; the verb does not mean execute a verification Task. |
| `report mismatch` | PRODUCT-TASK-SPEC-006-01 | implementation | Implement diagnostic emission for declared-ID and path-derived-ID divergence. |
| `Rename` | PRODUCT-TASK-SPEC-007-01 | authoring / migration preparation | Move the current canonical file into a temporary review-source identity. |
| `Compare` | PRODUCT-TASK-SPEC-007-04 | review | Review split artifacts against the preserved source for structure, coverage, and boundaries. |
| `Update` | PRODUCT-TASK-SPEC-007-05 | correction | Apply review-decided changes to canonical specs and directly related Work Item wording. |
| `Remove the temporary all-in-one review source` | PRODUCT-TASK-SPEC-007-06 | correction / cleanup | Delete the temporary comparison artifact after corrections and review complete. |
| `update PRODUCT-WORK-SPEC-007 for closure` | PRODUCT-TASK-SPEC-007-06 | synchronization | Propagate final cleanup and accepted completion into the parent Work Item. |

### Scanned-task ledger

| Task ID | findings logged | scan result |
|---|---:|---|
| PRODUCT-TASK-SPEC-001-01 | 2 | `Create or confirm`; `Define` |
| PRODUCT-TASK-SPEC-001-02 | 2 | `Confirm`; `reflected back into` |
| PRODUCT-TASK-SPEC-001-03 | 1 | `created or explicitly queued` |
| PRODUCT-TASK-SPEC-001-04 | 2 | lifecycle and Evidence synchronization phrases |
| PRODUCT-TASK-SPEC-004-01 | 2 | `Combine`; `Recommend whether` |
| PRODUCT-TASK-SPEC-004-02 | 2 | review judgment and finding-disposition phrases |
| PRODUCT-TASK-SPEC-004-03 | 3 | correction, follow-up routing, and Work Item closure phrases |
| PRODUCT-TASK-SPEC-005-01 | 2 | staging relocation and navigation rewrite phrases |
| PRODUCT-TASK-SPEC-005-02 | 0 | No finding. Work uses explicit authoring language with no unusual contextual meaning. |
| PRODUCT-TASK-SPEC-005-03 | 2 | completeness and checklist-based review phrases |
| PRODUCT-TASK-SPEC-005-04 | 2 | correction and parent-Evidence synchronization phrases |
| PRODUCT-TASK-SPEC-005-05 | 2 | staging materialization and byte-identity verification phrases |
| PRODUCT-TASK-SPEC-005-06 | 1 | `Restructure numbered sections` |
| PRODUCT-TASK-SPEC-005-07 | 1 | `is accounted for` |
| PRODUCT-TASK-SPEC-005-08 | 2 | correction and parent-Evidence synchronization phrases |
| PRODUCT-TASK-SPEC-005-09 | 2 | staging materialization and byte-identity verification phrases |
| PRODUCT-TASK-SPEC-005-10 | 2 | canonical restructuring and supporting source analysis phrases |
| PRODUCT-TASK-SPEC-005-11 | 1 | `is accounted for` |
| PRODUCT-TASK-SPEC-005-12 | 2 | correction and parent-Evidence synchronization phrases |
| PRODUCT-TASK-SPEC-005-13 | 2 | staging relocation and target-structure materialization phrases |
| PRODUCT-TASK-SPEC-005-14 | 1 | `moves out` |
| PRODUCT-TASK-SPEC-005-15 | 2 | deferred routing and finding-disposition phrases |
| PRODUCT-TASK-SPEC-005-16 | 2 | correction and deferred-state synchronization phrases |
| PRODUCT-TASK-SPEC-005-17 | 2 | canonical relocation and pointer-replacement phrases |
| PRODUCT-TASK-SPEC-005-18 | 2 | canonical relocation and source-spec trimming phrases |
| PRODUCT-TASK-SPEC-005-19 | 2 | `Trim`; `Add drift guards` |
| PRODUCT-TASK-SPEC-005-20 | 0 | No finding. Work uses an explicit review checklist and does not substitute an unusual owner verb. |
| PRODUCT-TASK-SPEC-005-21 | 3 | correction, Evidence propagation, and closure-readiness phrases |
| PRODUCT-TASK-SPEC-006-01 | 2 | `Check` and `report` used as implemented validator behavior |
| PRODUCT-TASK-SPEC-007-01 | 1 | `Rename` |
| PRODUCT-TASK-SPEC-007-02 | 0 | No finding. Work uses explicit creation language for canonical child specs. |
| PRODUCT-TASK-SPEC-007-03 | 0 | No finding. Work uses explicit creation language for the navigation Index. |
| PRODUCT-TASK-SPEC-007-04 | 1 | `Compare` |
| PRODUCT-TASK-SPEC-007-05 | 1 | `Update` used for review-driven correction |
| PRODUCT-TASK-SPEC-007-06 | 2 | cleanup and Work Item closure synchronization phrases |

### Verification result

- Scanned Tasks: 35 of 35 fixed-corpus Tasks.
- Finding entries: 56.
- Every finding cites an existing Task ID in PRODUCT-WORK-SPEC-001 through PRODUCT-WORK-SPEC-007.
- Zero-finding Tasks are recorded explicitly.
- No frequency ranking or canonical-term decision was made.
- No formal Investigation record was created under the accepted lightweight Evidence exception.
- The post-Evidence automated semantic validator was not invoked. TRV-ADR-SPEC-006 suspends validator delivery and prohibits using the deprecated prompt skill as a completion gate.
- Manual contract inspection against `spec:product.design_records.authoring_standards.task_authoring` found one bounded Evidence-only investigation outcome and no owned decision adoption, canonical authoring, implementation, independent review, correction, or synchronization.
- Only PRODUCT-TASK-SPEC-025-06 was modified.

# Design review gate

## Purpose

Independently review the final combined design state and route named findings without allowing the author to approve or repair their own work.

This gate covers:

- integrated independent review;
- finding classification;
- conditional correction Task creation;
- independent finding-closure review;
- return to new decision work when judgment is required.

Closure synchronization remains a separate phase.

## Review preconditions

Return `NOT READY` unless:

- the decision boundary is terminal;
- mandatory impact Investigation is complete;
- reconciliation and graph changes are complete;
- every decision has an ADR route;
- required ADR authoring is complete;
- originating-artifact and Specification authoring is complete;
- shared writers completed in deterministic order;
- exact reviewed artifacts are named;
- the authoring session has stopped changing the design;
- the reviewer is independent of authoring judgment.

Do not review an incomplete design as though it were final.

## One integrated review per Work Item

Use one final integrated review Task after all required writers complete.

Review the final combined state, not isolated intermediate diffs.

When a design boundary has an independent completion judgment, split it into another Work Item and give each Work Item its own integrated review.
Do not create several competing integrated verdicts inside one Work Item.

## Review boundary

The review Task names exact:

- parent Work Item;
- decision Task and decision IDs;
- impact Investigation;
- reconciliation and coordination Evidence;
- ADR-routing Task;
- new, amended, reused, or superseded ADRs;
- changed Requirements and Work Items;
- changed Specification files or sections;
- changed skills or workflow support files;
- deferred and blocked scope;
- relevant authoring and lifecycle standards.

Do not perform repository-wide traversal to compensate for a missing review contract.
Read adjacent artifacts only when a scoped contradiction requires them.

## Reviewer independence

The reviewer is read-only.

The reviewer must not:

- change decision status;
- edit reviewed artifacts;
- repair findings;
- amend the graph;
- synchronize lifecycle or Evidence;
- start implementation;
- treat the author's summary as proof.

The review may expose a new design gap directly visible in the scoped trace.
That gap returns to decision rather than being solved by the reviewer.

## Review trace

For every reviewed decision, verify:

```text
decision Task
  -> Investigation impact
  -> reconciliation and graph route
  -> ADR required / covered / not_required / blocked
  -> exact canonical artifact projection
  -> final combined Work Item state
```

For `deferred` or validly `blocked` items, verify:

```text
status
  -> exact reason
  -> dependency or destination
  -> proof that current completion remains coherent
```

## Review criteria

### Work Item integrity

Verify:

- one coherent resolution and completion boundary;
- direct material sources remain correct;
- Task responsibilities and dependencies match the accepted route;
- conditional branches are represented without speculative no-op Tasks;
- shared-writer and review order are explicit;
- production implementation is not hidden inside design closure.

### Decision integrity

Verify:

- material judgments are represented;
- explicit user choices are accurate;
- accepted authority answered repository-resolvable facts;
- dependencies and blockers are coherent;
- completed decisions were not rewritten for downstream progress;
- later revisions use new Tasks;
- no required decision exists only in chat or narrative Evidence.

### Investigation and routing integrity

Verify:

- every material impact was investigated;
- mismatch classes match evidence;
- Requirement and Work Item identity dispositions are coherent;
- graph changes have explicit owners;
- ADR routing is complete;
- ADR boundaries avoid both omission and fragmentation.

### ADR integrity

Verify:

- durable decisions have required ADRs;
- `covered` and `not_required` results are justified;
- reused ADRs remain accepted and non-superseded;
- amendments do not conceal reversals;
- supersession is explicit and historically honest;
- rationale and consequences are sufficient without duplicating Specifications.

### Canonical artifact integrity

Verify:

- every normative decision appears in the correct Specification or originating artifact;
- current-state text is normative rather than historical discussion;
- stale contradictions are removed or dispositioned;
- Requirements remain satisfied;
- exclusions are explicit;
- refs and relations follow repository policy;
- no hidden implementation-time design judgment remains.

### Workflow support integrity

When skills or instruction sources changed, verify:

- activation points to the intended successor;
- companions have coherent, non-overlapping authority;
- replaced authority is removed when required;
- canonical standards are referenced rather than contradicted;
- the workflow can return from authoring, review, and closure to new decision work safely.

## Verdicts

Use:

| verdict | meaning |
|---|---|
| `PASS` | No blocking, major, or required minor finding prevents closure. |
| `NEEDS REVISION` | One or more material findings require repair or renewed decision work. |
| `NOT READY` | Review prerequisites are incomplete. |
| `BLOCKED` | Required authority or Evidence is unavailable within the scoped boundary. |

Do not use `PASS WITH CHANGES` when changes were not independently reviewed.

## Finding severity

### Blocking

Use when:

- accepted Requirement or controlling authority is contradicted;
- required judgment remains unresolved while closure is claimed;
- canonical ownership cannot be determined;
- materially incompatible implementations remain possible;
- review independence or Evidence is invalid;
- required ADR supersession is absent.

### Major

Use when:

- a durable decision lacks a required ADR;
- ADR and Specification materially disagree;
- normative decisions are missing from canonical artifacts;
- workflow state or traceability is materially false;
- stale text would direct downstream work incorrectly;
- completion conditions cannot be satisfied as written;
- replaced workflow authority remains active contrary to the accepted activation route.

### Minor

Use when:

- bounded references, consequences, Evidence, or wording are incomplete;
- clarity defects do not change the selected design;
- a non-critical status or relation needs correction.

### Advisory

Use for optional improvements not required for closure.
Do not inflate stylistic preferences into findings.

## Finding format

Each material finding includes:

- finding ID;
- severity;
- affected decision IDs;
- affected artifact and section;
- observed contradiction or omission;
- required outcome;
- whether new user judgment is required;
- required owner type: correction, decision, or coordination.

## Finding routing

### Projection defect with no new judgment

Use `correction` when:

- the accepted decision is clear;
- the defect is incomplete or incorrect projection;
- stale text, missing refs, or contradictory wording can be repaired without choosing among alternatives.

Use coordination to create exact correction and independent finding-closure review Tasks derived from named findings.

### New or unresolved judgment

Return to `decision` when:

- a new choice is required;
- accepted decisions conflict;
- it is unclear whether Requirement or decision should change;
- correction would have to invent a boundary or rationale.

When the earlier decision Task is complete, coordination creates a new decision Task.
Affected downstream work remains blocked.

### Missing route or owner

Use coordination first when:

- no Task owns the required repair or decision;
- dependencies or writer order must change;
- downstream work must be blocked or released;
- another Work Item is required.

## Conditional Task materialization

Do not create correction or finding-closure review Tasks before named findings exist.

After `NEEDS REVISION`:

- group findings only when one repair owner and completion judgment apply;
- create exact correction Tasks with exact writable targets;
- create later independent review Tasks for closure;
- derive dependencies and Done conditions from actual findings;
- do not create placeholders or synthetic no-op Tasks.

## Correction boundary

A correction Task:

- names exact finding IDs;
- writes only exact affected artifacts and direct consistency effects;
- preserves accepted decisions not implicated by findings;
- stops when judgment is required;
- does not close findings;
- does not perform lifecycle closure;
- avoids unrelated cleanup.

## Finding-closure review

A later independent review verifies:

- each named finding is `CLOSED` or remains `OPEN`;
- direct cross-artifact consistency effects are correct;
- accepted decisions were not unintentionally changed;
- no direct regression was introduced;
- required new decision Evidence exists;
- revised ADR and Specification state agree.

Do not reopen the full Work Item unless the correction directly exposes a blocking contradiction.
A new finding must be caused by or directly exposed by the correction under review.

## New judgment after completed review

When review or later work exposes new design judgment after original Tasks are complete:

- retain completed decision, authoring, and review Tasks unchanged;
- create a new decision Task;
- create new authoring Tasks for revised canonical targets;
- create a new integrated review Task for the revised combined state;
- use prior completed records as historical Evidence and inputs.

A non-judgment finding stays on the correction and closure-review route.

## Review output

Default output:

1. Verdict
2. Reviewer independence
3. Reviewed artifacts
4. Decision-to-final-state trace result
5. Blocking findings
6. Major findings
7. Minor findings
8. Advisories
9. Implementation-planning readiness
10. Exact next gate

## Closure handoff

Proceed to `closure-synchronization.md` only after:

- initial integrated review returns `PASS`; or
- every required finding is independently `CLOSED`.

The review Task records the verdict and findings.
It does not update Work Item lifecycle or closure Evidence.

## Stop conditions

Stop and report the exact blocker when:

- a named review artifact is missing;
- the intended decision is unclear;
- accepted authority conflicts without precedence;
- the review boundary omits required Evidence;
- correction would require new judgment;
- closure would claim Evidence that does not exist.

# Boundary vocabulary: `decision` task_type

A `decision` Task owns judgment and a decision ledger. It does not own `authoring`, `review`, `implementation`, `correction`, `synchronization`, `coordination`, `investigation`, or Work Item decomposition/execution. This file lists confirmed real and realistic phrasings that describe one of those other responsibilities while avoiding its canonical verb.

## canonical: authoring (owned by `authoring`, not `decision`)

Pattern: writing or editing canonical content — an ADR, a Specification, a skill file — directly, under any verb other than "author".

Known paraphrases that ARE authoring, confirmed missed by every tested model (qwen3.6-27b, gpt-oss:20b, Claude Haiku 4.5; `reasoning_effort` off and confirmed genuine medium; with and without an accompanying self-declared denial statement):

- "encode the accepted decision text directly into the affected skill files, replacing the corresponding section content"
- "adjust that file's wording so it matches the ledger's accepted phrasing exactly"

(Source: TRV-INV-SPEC-004 sample_02/sample_03, synthetic mutations of `PRODUCT-TASK-SPEC-024-01`.)

Known SAFE pattern — NOT authoring, do not flag:

- "Quote the accepted decision wording verbatim in the ledger's decision summary column for downstream reference." — copying text into the Task's own ledger/summary is not writing external canonical content. (Source: TRV-INV-SPEC-004 sample_04.)

Ambiguous verb warning — "Fix":

In this repo's real corpus, `decision`-type Tasks routinely use "Fix" to mean "decide and settle", not "author" or "implement". Example: "Fix canonical target changes and every required graph amendment" (`PRODUCT-TASK-SPEC-018-11`) and "Fix the durable amendment boundary and exact authoring targets" (`PRODUCT-TASK-SPEC-018-16`) both resolve, in their own Done conditions, to "stated but not performed" — the Task decides *what* the target is, without writing it. Do not flag "Fix ... target(s)" as authoring by itself; check the Done condition for whether the target is described as performed/written or merely identified/decided.

Corroborating instances (PRODUCT-TASK-SPEC-025-04 reconciliation, 2026-07-03): the same "Fix X" = decide-not-perform shape recurs at least 9 more times in real `decision`-type Tasks — `PRODUCT-TASK-SPEC-019-03` ("Fix the corrected canonical target and downstream Investigation scope"), `PRODUCT-TASK-SPEC-019-07` ("Fix Requirement disposition"; "Fix W019 continuation or split disposition"; "Fix canonical target decisions"; "Fix shared-writer policy"; "State every required graph change without applying it"), and `PRODUCT-TASK-SPEC-019-09` ("Fix the checklist-authoring Work Item boundary"; "Fix the implementation Work Item boundary"; "Fix coarse release order between W019 and the two downstream Work Items"). This raises the pattern from 2 to 12 confirmed real-corpus instances.

Known SAFE pattern — NOT authoring, do not flag:

- "Define the accepted app-local Requirement, decision, ADR, Specification, integrated review, and closure boundary." — this reads like a list of authoring outputs, but in context (a `work_item_decomposition` Task scoping a child Work Item) it defines the *boundary* the child will later own, not the content itself. (Source: `PRODUCT-TASK-SPEC-021-13`.) Applies the same caution to `decision`-type Tasks that scope a downstream contract without writing it.

## canonical: independent review (owned by `review`, not `decision`)

Pattern: evaluating or confirming work from a perspective other than the original author or decision-maker, including a self-directed re-check framed as internal to the Task.

Known paraphrase that IS independent review, confirmed missed by every tested model under every tested condition:

- "Before closing, re-check each ledger item against the Specification from a separate perspective and note any mismatches for correction."

(Source: TRV-INV-SPEC-004 sample_05, synthetic mutation of `PRODUCT-TASK-SPEC-024-01`.)

Every tested model independently rationalized this as *not* independent review, drawing an unprompted distinction between a task's internal self-recheck and "true" independent review — treating the former as categorically exempt. Per this checklist's intent, that distinction is wrong: "from a separate perspective" is the review signal regardless of who performs it or whether it's framed as internal. Flag it.

## Open questions / not yet confirmed

- Whether "Fix" carries the same decide-vs-perform ambiguity in other `task_type`s (`coordination`, `work_item_decomposition`) has not been checked.
- No confirmed paraphrase examples exist yet for `implementation`, `synchronization`, `investigation`, or decomposition/execution actions disguised inside a `decision` Task. Add them here when found — do not guess.
- Single-occurrence candidate, not yet promoted (PRODUCT-TASK-SPEC-025-04 reconciliation, 2026-07-03): `PRODUCT-TASK-SPEC-019-12` (declared `decision`) contains "Hand the completed route to T13 without authoring ADR content" — a coordination/handoff-flavored action. Only one instance found; needs more before promotion to a confirmed `coordination`-inside-`decision` entry.
- Terminology-drift note, not a boundary-vocabulary finding: the "classify each item as no ADR / new ADR / amendment / supersession" ADR-routing task shape is declared `coordination` in `PRODUCT-TASK-SPEC-016-05` and `PRODUCT-TASK-SPEC-017-05`, but the same shape (`PRODUCT-TASK-SPEC-019-12`, "route validator decisions to ADR boundaries") is declared `decision`. This may reflect a real convention shift across the corpus rather than a paraphrase to catch — flagged for whoever next touches ADR-routing Task authoring guidance, out of scope for this vocabulary skill.
- Single-occurrence candidate, not yet promoted: `PRODUCT-TASK-SPEC-009-04` (declared `review`) contains "user sign-off" as a required Done-condition item — decision/coordination-flavored (obtaining explicit human approval) embedded in a review Task. Only one instance found.
- Single-occurrence candidate, not yet promoted: `PRODUCT-TASK-SPEC-011-08` (declared `review`/verification) contains "Work item updated" — lifecycle/synchronization-flavored, embedded in a review-cluster Task's own closure. Only one instance found; may simply be the corpus's normal small-scope review-closure convention rather than a boundary crossing.

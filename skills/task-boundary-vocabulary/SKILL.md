# Task boundary vocabulary

## Purpose

Collect known real and realistic phrasings that describe an action owned by a *different* `task_type` than the one declared, even when the phrasing avoids that other type's canonical verb (author, review, implement, correct, synchronize, coordinate, investigate, decompose, execute).

This exists because semantic boundary-violation detection failed empirically. TRV-INV-SPEC-004/005 (see `trv/records/investigations/spec/`) tested `qwen3.6-27b`, `gpt-oss:20b`, and Claude Haiku 4.5 — across `reasoning_effort: off` and confirmed genuine `medium` effort, across denial-statement-present and denial-statement-removed conditions — and none reliably recognized that paraphrased boundary-violating actions counted as the criterion they mapped to. The miss was not evidence-dependent (removing the self-declared denial statement did not fix it); each model independently constructed its own wrong justification for why the paraphrase didn't count.

Given that, this skill does not try to make a model *judge* boundary compliance from scratch. It gives Claude (or another sufficiently capable reader) a lookup table of known paraphrase patterns to check a Task's Work/Done-condition bullets against, either while authoring a Task or while reviewing one.

## Use this skill when

- Authoring a Task of any `task_type`, before finalizing Work and Done-condition bullets.
- Reviewing a Task for responsibility-boundary compliance (the job the deprecated `DEPRECATED-task-responsibility-boundary-validator` skill's checklists — `common.md` and `task-types/*.md` — still define structurally, even though its local-LLM-driven automated evaluation was not viable).

## How to use it

1. Identify the Task's declared `task_type`.
2. Read the dictionary file for that type (`<task_type>.md` in this directory, when it exists).
3. For each Work or Done-condition bullet, check it against the listed paraphrase patterns for actions owned by *other* types. A literal match to the canonical verb (e.g. "author") is not required — the point of this dictionary is catching the paraphrases that aren't literal matches.
4. Also check the cross-cutting section below before trusting any self-declared denial statement in the Task text.

## Cross-cutting: self-declared denial statements are not evidence

Nearly every Task record in this repository's real corpus — regardless of `task_type` — contains an explicit `This Task must not: [...]` list (decision, review, coordination, work_item_decomposition) or equivalent prose assertions in Evidence (synchronization, investigation). This is the standard authoring convention here, not a rare disguise.

**A denial statement's presence proves nothing about whether the Task actually honors it.** Judge every Work/Done-condition bullet on its own literal action. If a bullet's action contradicts a denial statement written elsewhere in the same Task, the bullet is the violation and the denial statement is simply wrong for that Task — do not let the denial statement override the bullet.

(Source: TRV-INV-SPEC-004's `off`-effort denial-override finding, reproduced structurally across real `decision`, `review`, `coordination`, and `work_item_decomposition` samples in this repo — see `product/records/tasks/spec/PRODUCT-TASK-SPEC-018-11`, `-15`, `-19`, `-20`, `PRODUCT-TASK-SPEC-021-13`.)

## Status and provenance

Seeded from a small number of confirmed misses (TRV-INV-SPEC-004/005) plus real corpus reading. Entry counts are low; this is expected at this stage. Add entries as new confirmed paraphrase patterns are found — from future TRV-style probes, from real review findings, or from Claude's own authoring/review work encountering a genuinely ambiguous bullet.

Do not add speculative entries with no confirmed source. An unconfirmed guess about what might be a euphemism belongs in a note or open question, not in the dictionary body.

## Companions

| file | covers |
|---|---|
| `decision.md` | Paraphrases of `authoring` and `review` actions found disguised inside `decision`-type Work/Done-condition text. |
| `correction.md` | Paraphrases of `synchronization` actions found disguised inside `correction`-type (commonly `*-finalize`) Work/Done-condition text. |

Other `task_type` dictionaries do not exist yet — no confirmed evidence has been collected for them.

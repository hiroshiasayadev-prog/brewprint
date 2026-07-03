# TRV-INV-SPEC-004: Decision Task batching and lexical priming effects

- **status**: concluded
- **date**: 2026-07-03
- **trigger**: Ad hoc re-evaluation of TRV-INV-SPEC-002's real-Task mutation method with neutral case labels exposed a different failure pattern than the labeled run: the target violation was missed while unrelated compliant criteria failed instead.
- **scope**: Evaluate, for one `decision`-type Task, whether case-label neutrality changes Qwen's failure pattern, whether the resulting cross-criterion noise is caused by inserted-content semantics or by criterion-batch size, and whether single-criterion isolation or grouped batching improves detection or reduces noise.
- **non_scope**: Statistical recall or precision claims, other Task types, other models (35b-a3b excluded per request), rule-pack redesign, adoption of any batching strategy, companion-decision-record context supply (explicitly rejected as infeasible for this model size).
- **source_refs**:
  - TRV-INV-SPEC-002
  - TRV-REQ-SPEC-001
  - spec:product.responsibility_boundary_validator
  - spec:product.design_records.authoring_standards.task_authoring
- **follow_up_candidates**:
  - Decide whether C08, C09, C10, C12, C13 should be consolidated (recurring co-failure cluster across four separate probes now)
  - Decide whether C04's literal wording ("only actions needed") should be revised; it is the single least stable criterion under both isolation and the denial-override rule
  - Evaluate whether the vocabulary-overlap priming effect (T04/T05 definition words appearing in Work) generalizes to other criterion/task-type pairs
  - Decide whether the denial-override rule concept should be abandoned entirely in favor of a deterministic post-hoc consistency check outside the model (comparing Work/Done-condition text against denial statements programmatically before or after the semantic evaluation)
  - Decide whether grouped batching should be used at all for `decision`-type Tasks, given it deleted a real detection (sample_05) in addition to reducing noise elsewhere; no evaluation-strategy adoption is currently recommended
  - Decide whether Tasks carrying a self-declared denial statement should be routed to mandatory human review instead of relying on any tested model configuration
- **related_work_items**:
  - TRV-WORK-SPEC-001
- **related_specs**:
  - spec:trv
  - spec:trv.model_runtime

## Investigation scope

This Investigation asks three bounded questions, all on one real `decision`-type Task (`PRODUCT-TASK-SPEC-024-01`):

> 1. With case labels stripped of content hints, does Qwen's failure pattern on the same mutation set from TRV-INV-SPEC-002 change?
> 2. Is the resulting cross-criterion false-negative/false-positive noise driven by inserted-content semantics, or by mere structural change (Work list length, insertion position)?
> 3. Does single-criterion isolation (1 criterion per call) or grouped batching (criteria split into ~7-9-item groups) improve detection of the real violation, reduce spurious noise, or both?

All evaluations used `qwen3.6-27b:q4_k_m-tools`, `temperature=0`, `reasoning_effort=off`. No repository files were modified as part of running this Investigation; only the evaluator prompts sent to the local model varied.

## Out of scope

- Any change to `common.md`, `decision.md`, or `evaluator-instructions.md`.
- Adoption of grouped batching as the production evaluation strategy.
- Testing `qwen3.6:35b-a3b` (explicitly excluded by the requester).
- Supplying companion decision-record context to the evaluator (explicitly rejected by the requester as impractical for this model size).
- Testing the R1/R2 interpretation rules proposed in prior conversation against the newly found noise pattern.

## Background

TRV-INV-SPEC-002 established that Qwen can miss a real responsibility violation entirely under a full 23-criterion batch prompt, and that a supplemental rule pack fixes detection at the cost of new false positives.

A follow-up session ran the same mutation method against a `decision`-type Task (`PRODUCT-TASK-SPEC-024-01`, which owns D-001 through D-021 including D-014's same-Work-Item direct materialization exception) and found the same two weaknesses: over-reliance on a self-declared denial statement, and cascade failures when a real violation is detected. A methodological flaw was also found and corrected: the case labels used in that session (e.g. `"CASE: M4, disguised independent review inserted"`) leaked the intended answer into the evaluator prompt, invalidating the results as a blind test.

This Investigation reruns the same five-variant mutation set with fully neutral labels (no case identifiers, no content hints, identical instruction text across all variants) to obtain an unbiased baseline, then investigates the resulting noise pattern through targeted structural probes and two remediation strategies: full isolation and grouped batching.

## What was investigated

### Base Task and variant set

| item | value |
|---|---|
| Base Task | `PRODUCT-TASK-SPEC-024-01` (task_type: `decision`) |
| Checklist | `common.md` (C01-C14) + `decision.md` (T01-T09), unmodified |
| Model | `qwen3.6-27b:q4_k_m-tools`, temperature 0, reasoning off |
| Labeling | Fully neutral; identical instruction text for every call; no case identifiers |

| variant | content | expected result |
|---|---|---|
| sample_01 | Unmodified real Task | 23/23 compliant |
| sample_02 | Explicit skill-file authoring inserted into Work/Done condition | T04 false |
| sample_03 | Same violation, reworded to avoid authoring-adjacent verbs | T04 false |
| sample_04 | Legitimate ledger-only quoting inserted (no external file touched) | 23/23 compliant |
| sample_05 | Disguised independent-review action inserted into Work/Done condition | T06 false |

### Structural probes (neutral-label baseline diagnosis)

Four additional single-bullet insertions were tested against the unmodified Control text to separate "any insertion" from "specific vocabulary" as the noise trigger:

| probe | inserted content | theme |
|---|---|---|
| Determinism check | none (Control re-run) | baseline stability |
| Trivial insertion | "Timestamp the ledger with the decision date for record-keeping." | unrelated to authoring/review |
| Read-only "skill file" mention | "Note in the ledger which skill file each decision eventually affects." | mentions "skill file", no write action |
| Read-only "authoring" mention | "Confirm the authoring Task ID before recording it in the ledger." | mentions "authoring", no write action |

### Isolation and grouping probes

| strategy | calls per Task | grouping |
|---|---|---|
| Full batch (baseline) | 1 | all 23 criteria in one call |
| Full isolation | 23 (partially run) | 1 criterion per call |
| Grouped batching | 3 | C01-C07, C08-C14, T01-T09 |

Isolation was run for the 6 criteria that misfired on sample_04 in full batch (C04, C07, C08, C09, C10, T09), but budget constraints limited the completed probe to C04, C08, and T09 against the unmodified Control text. Grouped batching was run in full (3 calls each) against sample_02, sample_04, and sample_05.

### Continuation probes: lexical desensitization rule and functional regrouping

A follow-up session tested two additional remediation strategies raised as open questions above, using `qwen3.6-27b:q4_k_m-tools`, `temperature=0`, `reasoning_effort=off`.

| probe | mechanism | variants tested |
|---|---|---|
| Lexical desensitization rule | One added evaluator-instruction sentence: word overlap between a Work bullet and a criterion's own definition text is not evidence for or against that criterion by itself; judge only the concrete action described. Full 23-criteria batch. | sample_02, sample_04, sample_05 |
| Functional regrouping | Criteria split into two thematic groups instead of the type-based groups used in the grouped-batching probe above: a boundary-violation group (C04, C08, C09, C10, C12, C13, T03, T04, T05, T06, T07, T08) and a structural group (C01, C02, C03, C05, C06, C07, C11, C14, T01, T02, T09). No rule text added. | sample_05, both groups |

The desensitization rule was tested alone, without the denial-override rule tested above. The functional regrouping probe targets the open question of whether a boundary-violation-themed group, rather than the criterion-type-themed T-group used above, would preserve the T06 detection that type-based grouping lost.

### Methodological correction: `reasoning_effort` was silently forced to `off` for every probe above

Every probe in this Investigation, including the continuation probes above, requested `reasoning_effort=off` intentionally per the original test plan. However, a second follow-up session discovered that the delegation tool's `output_schema` parameter (used throughout this Investigation to get structured JSON back) **silently overrides any requested `reasoning_effort` to `off`**, regardless of what is passed — the tool returns `override_reason: "structured_output_requires_final_content"` when this happens. This was not noticed until the requester questioned the results, since `off` was also the intentionally-requested value for every probe above, so the override was invisible in this Investigation's own data. A re-test using `reasoning_effort=medium` with the `output_schema` parameter removed (plain-text JSON instead) confirmed `effective_effort=medium` actually took effect (generation time increased roughly 3.5x, from ~48s to ~168s per call for sample_02).

Two variants were re-run against `qwen3.6-27b:q4_k_m-tools` with genuine `reasoning_effort=medium`, no rule, full 23-criterion batch, to check whether real reasoning changes the picture:

| variant | target criterion | target detected? | false criteria (noise) |
|---|---|---|---|
| sample_02 | T04 | **No — stayed true** (unchanged from the off-effort baseline) | none — 23/23 true |
| sample_05 | T06 | **No — stayed true** (regression — the off-effort, no-rule baseline correctly caught this) | none — 23/23 true |

Genuine medium-effort reasoning did not fix the T04 miss, and it lost the T06 detection that the (unintentionally) thinking-disabled baseline had achieved unaided. This is a fifth independent condition — alongside type-based grouping, thematic grouping, the denial-override rule, and the desensitization rule — under which sample_05's T06 detection does not survive. Unlike the other four, this one was not an intervention designed to reduce noise; it was simply turning reasoning on. The noise reduction it produced (0 false criteria on both variants) looks identical in shape to what the desensitization rule and grouped batching achieved, but arrived at with no evaluator-instruction or batching changes at all — only more compute per call. This raises a serious question about how much of this Investigation's earlier noise analysis (the lexical-priming findings, the C04/C08/C09/C10 cluster, all of it) is actually a property of `qwen3.6-27b:q4_k_m-tools`'s semantic evaluation versus an artifact of running that evaluation with reasoning disabled. **This Investigation's findings above should now be read as characterizing `qwen3.6-27b:q4_k_m-tools` with `reasoning_effort=off` specifically, not the model in general.**

### Follow-up probe: does removing the self-declared denial statement change the medium-effort miss?

The requester raised a specific hypothesis: the model might be treating the Task's own self-declared denial statements (e.g. "This Task does not author... canonical Specification content", "Confirm no formal Investigation, independent review... Task was created") as sufficient evidence on their own, regardless of what the actual Work bullets say — i.e. self-report being mistaken for a resolved criterion. To test this in isolation from the `reasoning_effort` question, sample_02 and sample_05 were each re-run under confirmed `reasoning_effort=medium` with the relevant self-declaration line(s) removed from the Task text (the Work-section and Done-condition denial sentences for the sample_02/T04 test; the Verification-section denial clause for the sample_05/T06 test), leaving the injected violation bullet as the only remaining signal.

| variant (denial text removed) | target criterion | target detected? | model's stated reasoning |
|---|---|---|---|
| sample_02, no T04 denial | T04 | **No — still missed** | "The task records decisions and encodes text into skill files but does not author a canonical ADR or Specification itself." |
| sample_05, no T06 denial | T06 | **No — still missed** | "Task performs a self-recheck for mismatches but does not own independent review." |

The hypothesis does not hold for these two cases under `medium` effort. With the denial statement removed, the model did not fall back to "insufficient evidence, therefore false" (which the evaluator instructions require) — it constructed its own justification for why the injected bullet does not count as the violation, without any self-declared text to lean on. For T04, it failed to connect "encode the accepted decision text directly into the affected skill files, replacing the corresponding section content" to "canonical ADR or Specification authoring". For T06, it drew an unprompted (and incorrect, per the checklist's intent) distinction between a task's internal self-recheck and "independent review", treating the former as categorically exempt. Both are genuine semantic misreadings, not evidence-dependent shortcuts. This suggests the `off`-condition denial-override finding from earlier in this Investigation and the `medium`-condition miss documented here may be two different failure modes that happen to produce the same missed result — one driven by over-trusting self-declared text (found under `off`), the other by the model's own (mistaken) semantic judgment persisting independent of that text (found under `medium`). This was tested on two samples only and should not be generalized without further probes.

## Findings

### Neutral-label baseline result (compare to TRV-INV-SPEC-002's labeled run)

| variant | target criterion | target detected? | additional false criteria (noise) |
|---|---|---|---|
| sample_01 | none | n/a | none — 23/23 true |
| sample_02 | T04 | **No** — stayed true | C04, C08, C09, C10, T09 (5) |
| sample_03 | T04 | **No** — stayed true | C04, C08, C09, C10 (4) |
| sample_04 | none (legitimate) | n/a | C04, C07, C08, C09, C10, T09 (6) |
| sample_05 | T06 | **Yes** | C04, C07, C08, C09, C10, C12, C13, T07, T08, T09 (10) |

Stripping the content-revealing case label did not fix the core detection miss on T04; it changed the failure's shape. Under the labeled run (prior session), the identical T04 mutation produced a clean 23/23 false-negative with no noise. Under the neutral-label run, the same mutation still misses T04 but now also misjudges unrelated, pre-existing, legitimate Work bullets (`Directly materialize the uniquely required T02 authoring Task`, `Update the parent Work Item Task flow and Task list`) as boundary violations. The label itself was influencing which failure mode appeared, not just whether one appeared.

### Structural probe results: lexical priming, not structural change

| probe | C04/C08/C09/C10/T09 pattern |
|---|---|
| Determinism check (Control re-run) | Identical to first Control run, including a duplicate-key artifact in the JSON output. Fully deterministic at this temperature/reasoning setting. |
| Trivial insertion ("timestamp... record-keeping") | 23/23 true. No noise. |
| Read-only "skill file" mention | C04, C08, C09, C10, T09 false — same pattern as sample_02/03, despite the inserted action being read-only and non-violating. |
| Read-only "authoring" mention | C04, C08, C09, C10, T09 false — same pattern again. |

The noise is not caused by Work-list length or insertion position. It is triggered by specific vocabulary — words that overlap with T04's and T05's own criterion definitions ("skill file", "authoring") — appearing anywhere in Work, regardless of whether the action described is a real violation. A semantically inert insertion does not trigger it; a semantically inert insertion that happens to use authoring-adjacent vocabulary does.

### Isolation probe results: criterion-specific, not uniform

| criterion (on unmodified Control) | full-batch result | isolated result |
|---|---|---|
| C04 | true | **false** — flipped to a new false positive under isolation |
| C08 | true | true — stable |
| T09 | true | true — stable |

Isolating C04 did not recover the correct judgment; it produced a false positive that does not exist in the full-batch evaluation of the same unmodified text. C04's literal wording ("The Work contains only actions needed to produce the primary outcome") appears to depend on having the fuller Done-condition/Verification context available in the same call to correctly recognize `T02` materialization as in-scope. Without that context, the model defaults to a stricter and incorrect reading. C08 and T09 did not show this fragility. Isolation is not a uniform fix; it is criterion-dependent, and for at least one criterion it makes things worse than full-batch.

### Grouped batching probe results: a precision/recall trade-off, not a fix

| Task | grouping | result |
|---|---|---|
| sample_02 (real T04 violation) | 3 groups (C01-07 / C08-14 / T01-09) | 23/23 true — noise eliminated, but T04 detection still missed |
| sample_04 (legitimate) | 3 groups (C01-07 / C08-14 / T01-09) | 23/23 true — matches expected result exactly |
| sample_05 (real T06 violation, detected under full batch) | 3 groups (C01-07 / C08-14 / T01-09) | 23/23 true — noise eliminated, **and the previously-correct T06 detection was also lost** |

Splitting the 23 criteria into three ~7-9-item groups fully eliminated the spurious cascade noise on all three variants tested. sample_04 became fully correct, which looked promising in isolation. But sample_05 shows the same intervention destroying a detection that full-batch got right without any rule pack: T06 was correctly false in sample_05's full-batch baseline, and reverted to true under grouping. The T06 reason text in the grouped run fell back on the same self-declared-denial citation pattern documented elsewhere in this Investigation ("Verification confirms 'no formal... independent review... Task was created'"), ignoring the actual injected Work bullet that grouping had, in the full-batch run, apparently helped the model notice.

Grouped batching is therefore not a one-directional improvement. It reduces false positives caused by lexical priming on legitimate content, but it can just as easily suppress a true positive that full-batch evaluation produced on its own. Removing criteria from the same call removes both the contaminating context and, evidently, some of the context the model was using — inconsistently — to catch real violations. There is no evidence in this Investigation that grouping is net-positive across cases; it moved the error in different directions on different variants.

### Denial-override rule isolated from the vocabulary rule

A follow-up probe isolated the denial-override interpretation rule ("if a general denial statement and a contradicting literal Work/Done-condition bullet both exist, treat the literal action as authoritative") from the vocabulary-mapping rule it was previously combined with in prior conversation. The rule was tested alone, in both full-batch and grouped configurations, against sample_02 (real T04 violation) and sample_04 (legitimate).

| condition | sample | result |
|---|---|---|
| Full batch + rule | sample_02 | T04 still **true** (missed). Noise increased to 7 false criteria (C04, C08, C09, C10, C12, C13, T09) — worse than the no-rule baseline's 5. |
| Full batch + rule | sample_04 | 8 false criteria (C04, C07, C08, C09, C10, C12, C13, T09) — worse than the no-rule baseline's 6. |
| Grouped, T-group only + rule | sample_02 | All 9 T-criteria true. T04 still missed. T05's reason text explicitly rationalized the injected violation as "documentation... not implementation", actively explaining it away despite the rule being present in the same prompt. |
| Grouped, C01-C07 group + rule | sample_04 | C04 and C07 newly false. Neither was false in the same group without the rule. C04's reason explicitly invoked the rule, citing the denial statement as conflicting with `Directly materialize the uniquely required T02 authoring Task` — a legitimate, D-014-sanctioned bullet that has nothing to do with the denial's actual subject (authoring canonical Specification content). |

The rule never triggered correctly on the actual target (the skill-file-encoding bullet in sample_02), in either full-batch or grouped configuration. It did trigger, incorrectly, on an unrelated legitimate bullet in sample_04's C01-C07 group, via surface keyword overlap between "authoring Task" and "does not author... content" rather than genuine semantic equivalence. In every tested configuration the rule was neutral-to-harmful: it added false positives without ever adding true positives.

### Lexical desensitization rule results: noise eliminated, but the one unaided detection is destroyed

| variant | target criterion | target detected? | false criteria (noise) |
|---|---|---|---|
| sample_02 | T04 | No — stayed true | none — 23/23 true |
| sample_04 | none (legitimate) | n/a | none — 23/23 true |
| sample_05 | T06 | **No — stayed true** | C04, C08, C09, C10 (4) |

The rule performed exactly as hoped on the two variants it targets: sample_02's C04/C08/C09/C10/T09 noise dropped to zero, and sample_04 went from 6 false criteria (the no-rule baseline) to a clean 23/23. But sample_05 — the one variant where full-batch evaluation caught a real violation unaided — lost that detection under the rule. Noise dropped from 10 false criteria to 4, but T06 flipped back to true (missed), with the same self-declared-denial citation pattern seen elsewhere in this Investigation ("Verification confirms 'no formal... independent review... Task was created'").

This is the same precision/recall trade-off documented for grouped batching above, reproduced through a completely different mechanism. Grouped batching removes noise by removing criteria from the same call; the desensitization rule removes noise by telling the model to discount the very vocabulary overlap that, in the no-rule baseline, apparently helped it notice the disguised-review bullet in sample_05. Two independent interventions — one structural, one instructional — trade the same detection for the same noise reduction.

### Functional regrouping results: thematic grouping does not recover T06 either

| group | criteria | T06 detected? | false criteria |
|---|---|---|---|
| Boundary-violation group (12) | C04, C08, C09, C10, C12, C13, T03, T04, T05, T06, T07, T08 | **No — stayed true** | C04 (1) |
| Structural group (11) | C01, C02, C03, C05, C06, C07, C11, C14, T01, T02, T09 | n/a (T06 not in this group) | none — 11/11 true |

Grouping T06 together with the other boundary-violation criteria — rather than with the rest of the `decision`-type checklist as in the T-group probe above — did not recover the detection. The boundary-violation group's own T06 reason text fell back on the same self-declared-denial citation pattern used elsewhere. This rules out "T06 was diluted by unrelated structural criteria sharing the same call" as the mechanism behind grouping's earlier loss of this detection: a group built entirely from thematically related boundary criteria still lost it. Noise did drop sharply in this group (1 false criterion vs. the original full-batch baseline's 10), so the noise-reduction effect of grouping is confirmed again; only the detection-preservation question remains open.

## Cross-cutting observations

- Case-label neutrality is necessary for a valid blind test, and it materially changes which failure mode is observed. Prior single-label results (this session's and TRV-INV-SPEC-002's) should not be read as fully representative of unlabeled production behavior.
- The C04/C08/C09/C10 co-failure cluster has now appeared in three separate probes (sample_02, sample_03, sample_04) under full-batch evaluation, always together, always around the same two pre-existing legitimate Work bullets. This is consistent with the semantic-overlap concern raised in prior conversation about consolidating C08/C09/C10/C12/C13.
- The noise trigger is lexical, not semantic: read-only, non-violating mentions of "skill file" or "authoring" are sufficient. A rule instructing the evaluator to "ignore vocabulary overlap and judge only the described action" was not tested in this Investigation and remains a candidate.
- Grouped batching is not a reliable precision-only intervention. It removed noise on sample_04, but on sample_05 it removed a correct detection that full-batch achieved unaided. Any recommendation for grouped batching must account for this regression, not just the noise-reduction cases.
- Full isolation was expected to be the strictest, most conservative evaluation strategy. It was not uniformly conservative: it introduced a false positive on C04 that neither full-batch nor grouped batching produced for the same unmodified text.
- The core T04 miss (self-declared denial statement overriding a contradicting Work bullet) persisted across every batching strategy tested — full batch, grouped batch, and by inference isolation (not directly tested against a real violation). This confirms the miss is orthogonal to batch size; it requires a semantic/interpretation-level fix, not a prompt-structure fix.
- The denial-override rule, tested in isolation from the vocabulary-mapping rule it was previously paired with, never fixed detection in any configuration and actively worsened precision in three of four tested configurations. This indicates the earlier apparent success of the combined rule pack (in prior conversation, before this Investigation) was carried entirely by the vocabulary-mapping component, not by the denial-override principle. The denial-override rule's own mechanism — surface text matching between a denial statement and a candidate bullet — is exactly as vulnerable to keyword collision as the lexical priming effect documented above; it is not a fix for that effect, it is another instance of it.
- Two further, mechanistically unrelated interventions converge on the same trade-off found for grouped batching: the lexical desensitization rule (instructional, applied to full batch) and functional/thematic regrouping (structural, grouping by theme rather than criterion type) both eliminate sample_05's noise while also eliminating its one unaided true detection (T06). Four independent intervention mechanisms now — type-based grouping, thematic grouping, the denial-override rule, and the desensitization rule — have each been shown to destroy this specific detection. This convergence suggests the T06 detection in full-batch evaluation depends on some cross-criterion signal (plausibly the same vocabulary-overlap channel identified as the noise source) rather than on stable semantic recognition of the disguised action. The model may not be reliably right for the reason a validator would need it to be right.
- **All findings above were produced with `reasoning_effort` silently forced to `off`** (see the methodological correction note under "What was investigated"), not by deliberate choice validated against a reasoning-enabled baseline. A minimal genuine-`medium`-effort re-test (sample_02, sample_05) reproduced the T04 miss and *also* lost the T06 detection — a fifth mechanism producing the same trade-off, this time from more compute rather than a structural or instructional intervention. Every noise-reduction result in this Investigation should be treated as provisional until re-run with confirmed non-`off` reasoning; it is not yet known how much of the C04/C08/C09/C10 lexical-priming pattern survives when the model is actually allowed to think.

## Follow-up judgment candidates

- Decide whether grouped batching should be used at all for `decision`-type Tasks, given that it deleted a real detection (sample_05/T06) as well as reducing noise; a per-Task or per-criterion decision may be required rather than a blanket policy.
- Decide whether C08, C09, C10, C12, C13 warrant consolidation given the repeated co-failure pattern, and whether C04's wording should be revised given its unique fragility under both isolation and the denial-override rule.
- [tested] The lexical-desensitization rule was tested in a follow-up session (see Findings above): it eliminates noise on sample_02/sample_04 but also eliminates sample_05's unaided T06 detection, the same trade-off found for grouped batching. Decide whether this rule is worth adopting anyway for Tasks/checklists where the noise cost outweighs the risk of losing this specific detection class, or whether it should be abandoned alongside grouped batching for the same reason.
- **[critical, supersedes prior priority]** Re-run this Investigation's full probe set (neutral-label baseline, structural probes, isolation, grouped batching, denial-override rule, desensitization rule, functional regrouping) with confirmed genuine `reasoning_effort=medium` or higher — not `output_schema`-forced `off`. The `output_schema` parameter must be omitted or the tool's structured-output override must be otherwise bypassed; verify `effective_effort` in the tool response matches the request before trusting any result. Every finding in this record was produced under an unintended `off` condition and needs re-validation before any adoption decision is made.
- Investigate the two T04/T06 misses as distinct semantic-classification gaps now that the self-declaration hypothesis has been tested and did not explain them under `medium` effort: does the model need explicit examples mapping "encode text into a skill file" to "Specification authoring", and "self-recheck against the Specification" to "independent review", or is this a deeper checklist-wording ambiguity (e.g. T06's "independent" could plausibly be read as requiring an external reviewer, which the checklist does not define)?
- Decide whether interpretation-rule-based detection fixes for this class of miss should be abandoned in favor of a deterministic, non-model consistency check (e.g. a script that flags any Task where a denial phrase and a semantically related action verb co-occur, for human or separate-pass review) given that two independent rule attempts (the earlier combined pack, and this Investigation's isolated denial-override rule) both failed to achieve clean detection.
- Decide whether `decision`-type Tasks with a self-declared denial statement should simply be routed to mandatory human review, given that no tested model configuration reliably both avoids noise and preserves detection.

## Recommendation

No batching or isolation strategy tested in this Investigation should be adopted as-is. Full batch (23 criteria/call) is the only strategy that caught a real violation unaided (T06 on sample_05), but it does so with heavy, inconsistent noise (4-10 spurious false criteria depending on the variant) and it never caught the T04-class violation regardless. Grouped batching (3-way split) cleaned up that noise on the legitimate-content and T04-miss variants, but on sample_05 it deleted the one unaided detection full-batch had. Full single-criterion isolation is the most expensive option (up to 23x call cost) and was shown to introduce a new false positive (C04) that neither of the other two strategies produced. None of the three is dominant; each wins on some variants and loses on others, and the losses are not confined to noise — grouping's sample_05 result shows it can cost a real detection outright.

The denial-override interpretation rule, tested here in isolation and in combination with grouped batching, does not appear preferable to leaving the checklist unmodified. It never achieved detection of the target violation in any tested configuration, and it introduced new false positives via keyword collision that were not present without it. Prompt-level interpretation rules built on general principles ("if a denial and a contradiction exist, prefer the contradiction") appear to depend on the model already recognizing semantic equivalence between the described action and the denied action — which is the actual hard part this class of rule was meant to solve, not a precondition it can assume.

Given that rule-based, isolation-based, and grouping-based interventions have each been tested without producing a configuration that reliably improves recall without also risking precision or an existing detection, a non-model deterministic pre-check (flagging denial/action co-occurrence for separate handling before or after semantic evaluation) appears more promising than further prompt-structure iteration for the denial-override failure mode specifically. For the general reliability question, this Investigation does not support recommending any single evaluation strategy over full-batch as a default; the honest conclusion is that this model, on this checklist, does not yet have a configuration that is both low-noise and non-regressive across the variant types tested.

A follow-up session tested the lexical desensitization rule and a thematic (boundary-violation-based) regrouping of sample_05, specifically to answer two of this Investigation's open questions. Neither closes the gap identified above. The desensitization rule is a fourth intervention — alongside full isolation, grouped batching, and the denial-override rule — that trades sample_05's one unaided detection for cleaner output on the other variants; it does not warrant a different recommendation than grouped batching received. Thematic regrouping rules out one candidate explanation (dilution by unrelated structural criteria) for why grouping loses T06, but does not identify the actual mechanism. The core recommendation is unchanged and now more strongly supported: no tested intervention, across five independent mechanisms (isolation, type-based grouping, thematic grouping, the denial-override rule, and the desensitization rule), reliably preserves both low noise and the sample_05 detection at once.

**Supersedes the above**: a methodological flaw was found after this Investigation's continuation probes were written up — every result in this record, including the desensitization rule and functional regrouping probes, ran with `reasoning_effort` silently forced to `off` by the `output_schema` parameter. A minimal genuine-reasoning re-test (medium effort, sample_02 and sample_05, no rule) reproduced the T04 miss and additionally lost the T06 detection, reproducing the same noise-vs-detection trade-off through a sixth mechanism (more compute, no structural or instructional change). This does not overturn the recommendation above — if anything it strengthens the case that no tested condition preserves both low noise and the sample_05 detection — but it means the specific numbers in this record (which criteria fail, by how much, under which named intervention) are not yet validated against a reasoning-enabled evaluator and should not be used to make adoption decisions until re-run properly. The immediate next step is re-running the full probe matrix with confirmed `effective_effort` matching the request.

## Follow-up artifact candidates

- A decision record on whether to consolidate C08/C09/C10/C12/C13 in `common.md`.
- A decision record on C04's wording.
- Regression probes reusing the four structural-priming insertions (trivial / skill-file-mention / authoring-mention / determinism-check) as a lightweight lexical-sensitivity check for any future checklist or evaluator-instruction change.
- A design spike for a deterministic denial/action co-occurrence pre-check, run outside the model, as an alternative to prompt-level interpretation rules for this failure class.

## Open questions

- Does the C04 isolation fragility reproduce on other Tasks, or is it specific to `PRODUCT-TASK-SPEC-024-01`'s particular Done-condition/Verification phrasing?
- Would a 2-way split (C01-C14 / T01-T09) or a different grouping boundary preserve the precision benefit of the 3-way split while also improving recall on T04-class violations?
- Does the lexical priming effect reproduce on `authoring`-type or other task-type checklists, where the criterion-defining vocabulary differs?
- Would explicitly listing "words that must not be treated as evidence on their own" (e.g. skill file, authoring, canonical) as part of the evaluator instructions reduce the priming effect without reintroducing the denial-override miss?
- Was the isolation probe's incomplete coverage (3 of 6 misfired criteria tested) sufficient to characterize the fragility pattern, or would testing C07, C09, C10 in isolation change the picture?
- Would a Task-specific action-to-denial mapping rule (rather than the general denial-override principle tested here) achieve detection without the keyword-collision false positives, and is authoring such a rule per-Task practical at scale?
- Does the denial-override rule's keyword-collision failure mode (firing on "authoring Task" instead of "author... content") generalize to other denial statements in other Tasks, or is it specific to this Task's coincidental reuse of the word "authoring" in both a legitimate bullet and the denial sentence?
- Why did grouping the T06 criterion away from the other 22 remove the model's ability to notice the disguised-review bullet, when full-batch caught it unaided? Is the model using some other criterion's context as an indirect cue for T06 in full-batch, and if so, which one? [Partially answered: it is not dilution by unrelated structural criteria — a thematic boundary-violation-only group loses the detection too. The cue, whatever it is, appears to require the full 23-criterion context specifically, not just topical proximity to other boundary criteria.]
- Does T06 detection in full-batch evaluation depend on the same vocabulary-overlap channel that causes the C04/C08/C09/C10 noise cluster? The desensitization rule suppresses both the noise and the T06 detection simultaneously, which is consistent with a shared mechanism, but this Investigation did not isolate the two effects.

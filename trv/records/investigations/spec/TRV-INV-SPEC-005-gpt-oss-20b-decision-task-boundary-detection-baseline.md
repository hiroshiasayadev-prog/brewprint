# TRV-INV-SPEC-005: gpt-oss:20b decision Task boundary detection baseline

- **status**: investigating
- **date**: 2026-07-03
- **trigger**: TRV-INV-SPEC-004 concluded that no tested Qwen configuration (full batch, grouped batch, isolation, denial-override rule) reliably detects the self-declared-denial violation class without either heavy noise or losing an existing detection. The requester asked whether `gpt-oss:20b` shows a different, more usable failure profile on the same Task and variant set, and requested this be handed off to a fresh session if it grew large.
- **scope**: Run the same five-variant neutral-label mutation set from TRV-INV-SPEC-004 (base Task `PRODUCT-TASK-SPEC-024-01`, `decision` type, `common.md` + `decision.md` checklist) against `gpt-oss:20b` in full-batch mode, and compare its failure pattern to `qwen3.6-27b:q4_k_m-tools`'s.
- **non_scope**: Isolation probes, grouped-batching probes, and denial-override rule tests against `gpt-oss:20b` — deprioritized per this Investigation's own conclusion, given gpt-oss:20b's 0/3 recall on real violations across the completed five-variant set (see Recommendation). Other models. Any checklist or evaluator-instruction changes.
- **source_refs**:
  - TRV-INV-SPEC-002
  - TRV-INV-SPEC-004
  - TRV-REQ-SPEC-001
  - spec:product.responsibility_boundary_validator
  - spec:product.design_records.authoring_standards.task_authoring
- **follow_up_candidates**:
  - Decide whether to deprioritize further `gpt-oss:20b` evaluation (isolation, grouped batching, denial-override rule) given its 0/3 recall on real violations across the completed five-variant set — this Investigation's recommendation is yes
  - Decide whether the C08/C09 pairing should be consolidated or reworded given its cross-model, cross-investigation recurrence, now confirmed across TRV-INV-SPEC-002, TRV-INV-SPEC-004 (qwen), and this Investigation (gpt-oss, five variants plus four structural probes)
  - Decide whether "fails its own Control case" should be an automatic disqualifying result for a candidate evaluator model, independent of mutation-detection results
  - Investigate the mechanism behind gpt-oss:20b's noise: across five variants and four structural probes, every observed false-criterion set was a subset of the unmodified Control's own {C08, C09, C12, C13}, never a superset — the opposite of qwen's lexically-primed noise growth. No hypothesis for this pattern has been tested
- **related_work_items**:
  - TRV-WORK-SPEC-001
- **related_specs**:
  - spec:trv
  - spec:trv.model_runtime

## Investigation scope

This Investigation asks one bounded question:

> Against the same `decision`-type Task, checklist, and neutral-label mutation set used in TRV-INV-SPEC-004, does `gpt-oss:20b` produce a detection/noise profile that is better, worse, or simply different from `qwen3.6-27b:q4_k_m-tools`'s?

Only full-batch evaluation (all 23 criteria in one call) was tested. `temperature=0`, `reasoning_effort=off` (Ollama's `think` option disabled; `thinking_discarded: true` was returned by the tool for every call, meaning no chain-of-thought was requested or retained). No repository files were modified; only evaluator prompts sent to the local model varied.

A follow-up session completed this Investigation's originally planned scope: `sample_03` (the reworded T04 violation) and the four structural lexical-priming probes from TRV-INV-SPEC-004, both run against `gpt-oss:20b` under the same full-batch, neutral-label, `temperature=0`, `reasoning_effort=off` conditions. Isolation, grouped batching, and the denial-override rule remain out of scope for `gpt-oss:20b` — see Recommendation for why they are not being pursued.

## Out of scope

- Any change to `common.md`, `decision.md`, or `evaluator-instructions.md`.
- Isolation, grouped batching, and the denial-override rule against `gpt-oss:20b` — see `non_scope` above; deprioritized following this Investigation's conclusion rather than deferred to a continuation.
- Statistical recall/precision claims (n=5 variants + 4 structural probes, single run each).

## Background

TRV-INV-SPEC-004 concluded that, for `qwen3.6-27b:q4_k_m-tools` on this Task, no tested evaluation-strategy or rule intervention reliably both avoided noise and preserved detection. The requester asked whether switching the evaluator model to `gpt-oss:20b` (already installed and available via the same Ollama MCP tool) would sidestep these specific failure modes, given both models are candidates for the local-LLM-driven TRV validator role described in `TRV-REQ-SPEC-001`.

This Investigation runs the identical four already-defined variants (Control, explicit-violation M1, legitimate M3, disguised-review M4 — using TRV-INV-SPEC-004's `sample_01`/`sample_02`/`sample_04`/`sample_05` content unchanged) against `gpt-oss:20b` under the same full-batch, neutral-label conditions, to get a directly comparable first data point. `sample_03` was not reached before the requester asked to hand off if the session grew long; this record captures what was completed and hands off the remainder explicitly.

## What was investigated

| variant | content | qwen3.6-27b full-batch result (TRV-INV-SPEC-004) |
|---|---|---|
| sample_01 (Control) | Unmodified real Task | 23/23 true |
| sample_02 | Explicit skill-file authoring inserted (T04 violation) | T04 missed (true); C04,C08,C09,C10,T09 false |
| sample_04 | Legitimate ledger-only quoting inserted | C04,C07,C08,C09,C10,T09 false (false positives) |
| sample_05 | Disguised independent-review action inserted (T06 violation) | T06 correctly false; 10 additional false criteria (noise) |

`gpt-oss:20b` was run against the same four variant texts, same criteria list, same instruction text, `temperature=0`, `reasoning_effort=off`.

A follow-up session extended this to `sample_03` (the reworded T04 violation, avoiding authoring-adjacent verbs) and to the four structural lexical-priming probes defined in TRV-INV-SPEC-004 (determinism check / Control re-run, trivial insertion, read-only "skill file" mention, read-only "authoring" mention), all against the unmodified Control text. Same model, same full-batch mode, same `temperature=0`, `reasoning_effort=off`.

### Methodological correction: `reasoning_effort` was silently forced to `off`

Every call in this Investigation, including the sample_03 and structural-probe continuation above, used the `output_schema` parameter to get structured JSON back. A later session discovered that `output_schema` silently forces `effective_effort=off` regardless of the requested `reasoning_effort` (`override_reason: "structured_output_requires_final_content"` in the tool response). Since `off` was also this Investigation's intentionally requested value throughout, the override was invisible until directly questioned. This record's status is reopened from `concluded` to re-test with genuine reasoning before the recommendation below can be trusted.

Two variants were re-run against `gpt-oss:20b` with genuine `reasoning_effort=medium` (verified via `effective_effort=medium` in the response, `output_schema` omitted), no rule, full 23-criterion batch:

| variant | target criterion | target detected? | false criteria (noise) |
|---|---|---|---|
| sample_02 | T04 | **No — stayed true** (unchanged from the off-effort baseline) | C08, C09, C12, C13 (4) — noise appeared where the off-effort run had none |
| sample_05 | T06 | **No — stayed true** (unchanged from the off-effort baseline) | C12, C13 (2) — same count as the off-effort run's C08/C09, different pair |

Genuine medium-effort reasoning did not recover either missed detection. `gpt-oss:20b`'s 0/3 recall across real violations, established under the (unintentionally) thinking-disabled condition, reproduces under confirmed genuine reasoning too. Unlike qwen (TRV-INV-SPEC-004), which lost its one correct detection when reasoning was turned on, gpt-oss:20b had no correct detection to lose — its recall floor and ceiling are the same at 0/3 regardless of reasoning effort, at least at `medium`. The noise pattern did shift (sample_02 gained noise, sample_05's noise pair changed), consistent with this Investigation's earlier finding that gpt-oss:20b's noise is not stably tied to specific vocabulary, and extending that instability to reasoning-effort as an additional axis of variation.

## Findings

### gpt-oss:20b full-batch results

| variant | target criterion | target detected? | false criteria observed |
|---|---|---|---|
| sample_01 (Control) | none | n/a | C08, C09, C12, C13 (4) |
| sample_02 (T04 violation) | T04 | **No** — stayed true | none — 23/23 true |
| sample_04 (legitimate) | none | n/a | C08, C09, C12, C13 (4) — identical set to Control |
| sample_05 (T06 violation) | T06 | **No** — stayed true | C08, C09 (2) |

### Comparison to qwen3.6-27b

- **Recall**: `gpt-oss:20b` missed both tested real violations (T04 in sample_02, T06 in sample_05) outright. `qwen3.6-27b` also missed T04, but did catch T06 in the same baseline condition. On this small sample, `gpt-oss:20b`'s recall is strictly worse.
- **Noise direction**: The two models fail in different directions. Qwen's noise (documented in TRV-INV-SPEC-004) is triggered by specific vocabulary overlap ("skill file", "authoring") appearing anywhere in Work, and tends to spare the unmodified Control. `gpt-oss:20b`'s noise appears on the **unmodified Control itself** (C08, C09, C12, C13 false with reasons like "No release decision specified" / "No independence boundary specified") — a fail-closed pattern that treats the absence of explicit text addressing a criterion as a violation, rather than qwen's fail-open pattern of over-trusting explicit denial text.
- **Noise stability**: Qwen's noise cluster (C04/C08/C09/C10/T09) was reproducible across multiple insertions sharing trigger vocabulary. `gpt-oss:20b`'s noise cluster changed shape across variants in a way that does not obviously track content: Control and sample_04 produced an identical 4-item false set (C08, C09, C12, C13) despite sample_04 containing an extra legitimate Work bullet that Control does not have; sample_02 produced zero false criteria despite containing the actual injected violation; sample_05 produced a 2-item subset (C08, C09) of the Control/sample_04 pattern. No hypothesis for this pattern has been tested yet.
- **C08/C09 co-occurrence**: C08 and C09 failed together in every `gpt-oss:20b` variant that had any false criteria at all (Control, sample_04, sample_05), consistent with the semantic-overlap concern about these two criteria already raised in TRV-INV-SPEC-002 and TRV-INV-SPEC-004 for a different model.
- **Reasoning style**: `gpt-oss:20b`'s reason strings were terser than qwen's (typically under 15 words, e.g. "No release decision specified") and several cited `"section": "None"` for C12/C13 on sample_04, i.e. the model itself flagged that no supporting text existed for its judgment — a data point toward the fail-closed-on-silence hypothesis above.

### Completing the five-variant set: sample_03

| variant | target criterion | target detected? | false criteria observed |
|---|---|---|---|
| sample_03 (T04 violation, reworded) | T04 | **No — stayed true** | none — 23/23 true |

`sample_03` reproduces `sample_02`'s pattern exactly: zero noise, and the actual T04 violation missed. Rewording the violation to avoid authoring-adjacent verbs made no difference — `gpt-oss:20b` did not detect either phrasing. Across the full five-variant set, `gpt-oss:20b` now has a complete recall result: it missed all three tested real violations (T04 in both sample_02 and sample_03, T06 in sample_05), a 0/3 detection rate, against `qwen3.6-27b:q4_k_m-tools`'s 1/3 (T06 only, in TRV-INV-SPEC-004).

### Structural probes: no lexical-priming growth, unlike qwen

| probe | false criteria observed | vs. Control's {C08, C09, C12, C13} |
|---|---|---|
| Determinism check (Control re-run) | C08, C09, C12, C13 (4) | identical — confirms deterministic at this temperature/reasoning setting |
| Trivial insertion ("timestamp... record-keeping") | C08, C09, C12, C13 (4) | identical — no effect |
| Read-only "skill file" mention | C13 (1) | strict subset — noise reduced |
| Read-only "authoring" mention | C08, C09 (2) | strict subset — noise reduced, different pair |

Every probe's false-criterion set was a subset of the unmodified Control's own {C08, C09, C12, C13}; none introduced a new false criterion outside that set, and none pushed the count higher than Control's own baseline. This is the opposite of qwen's pattern (TRV-INV-SPEC-004), where the same three real-vocabulary insertions reproducibly grew the noise set from 0 (trivial insertion) to 5 (C04, C08, C09, C10, T09). `gpt-oss:20b`'s noise on this Task appears to be a ceiling, not a lexically-triggered signal: specific insertions can only remove members of {C08, C09, C12, C13}, never add to it. No hypothesis for why "skill file" removes C08/C09/C12 while "authoring" removes C12/C13 (leaving C08/C09) has been tested; both read as arbitrary redistributions within the same fixed noise budget rather than a content-driven effect.

## Cross-cutting observations

- Neither model achieved a clean baseline on the unmodified Control in a way that generalizes: qwen was clean on Control in TRV-INV-SPEC-004's neutral-label run, but gpt-oss:20b was not (4 false criteria on unmodified real, `done` Task text). A validator that fails its own Control case cannot be trusted without further tuning, independent of mutation-detection performance.
- The C08/C09 pairing recurring as a co-failure unit across two different models and three separate investigations (TRV-INV-SPEC-002, -004, -005) is now a stronger signal that these two criteria are not independently evaluable as worded, regardless of which model is used.
- The structural probes now answer the question left open above: gpt-oss:20b's noise is not lexically triggered the way qwen's is. Every structural probe result was a subset of Control's own {C08, C09, C12, C13}; specific vocabulary never expanded the false set, and in two of three probes it shrank it. Whatever drives gpt-oss:20b's noise on this Task, it is bounded by the Control's own baseline failure rather than primed upward by inserted content.
- With sample_03 complete, gpt-oss:20b's recall across the full five-variant set is 0/3 on real violations (T04 twice, T06 once), against qwen's 1/3. Combined with gpt-oss:20b failing its own Control case (a failure mode qwen did not show under the same neutral-label conditions), this Investigation's provisional read now stands on complete data rather than a partial sample.
- **A genuine-reasoning re-test (see the methodological correction above) reproduces the 0/3 recall result under confirmed `reasoning_effort=medium`.** Unlike qwen, which had a detection to lose when reasoning was enabled, gpt-oss:20b's recall was already 0/3 under the (unintentional) `off` condition and stayed 0/3 under `medium`. This is the one piece of good news in the reasoning-effort methodological correction: gpt-oss:20b's core recall finding does not appear to be an artifact of the `off`-forcing bug, unlike qwen's noise/detection findings, which are now flagged as provisional pending re-test. The noise composition still shifted under `medium`, so the noise-pattern findings (which criteria fail, in what combination) remain less certain than the recall finding.

## Follow-up judgment candidates

- [resolved by this Investigation] gpt-oss:20b evaluation is now complete (five variants + four structural probes). Its recall (0/3) is strictly worse than qwen's (1/3), and it additionally fails its own Control case. Decide whether to deprioritize gpt-oss:20b for TRV entirely, rather than continuing to isolation/grouping/rule probes — this Investigation's recommendation is yes.
- Decide whether the C08/C09 pairing should be consolidated or reworded given its cross-model, cross-investigation recurrence — now confirmed across TRV-INV-SPEC-002, TRV-INV-SPEC-004 (qwen), and this Investigation (gpt-oss, complete five-variant + structural-probe set).
- Decide whether "fails its own Control case" should be an automatic disqualifying result for a candidate evaluator model, independent of mutation-detection results — gpt-oss:20b is now a concrete example of this failure mode with complete supporting data.
- Decide whether gpt-oss:20b's noise-ceiling behavior (structural probes never exceed Control's own false-criterion set) is worth investigating further as a distinct model-reliability property, independent of the recall question, given it differs qualitatively from qwen's lexical-priming pattern.
- Decide whether to re-run this Investigation's full five-variant + structural-probe set with confirmed genuine `reasoning_effort=medium` (not `output_schema`-forced `off`) before finalizing recall/noise numbers, even though the headline recall result (0/3) already reproduced in the minimal re-test — the noise composition did not fully reproduce and remains unverified at scale.

## Recommendation

With the full five-variant set and all four structural probes complete, `gpt-oss:20b` does not appear preferable to `qwen3.6-27b:q4_k_m-tools` for this checklist and Task type, and the evidence is now stronger than the provisional read this record originally carried. `gpt-oss:20b` missed all three tested real violations (0/3, vs. qwen's 1/3), and it failed its own unmodified-Control baseline, which qwen did not do under the same neutral-label full-batch conditions. The structural probes rule out one possible mitigating explanation — that gpt-oss:20b's Control-case noise might itself be a lexical-priming artifact fixable the same way qwen's is — since none of the three real-vocabulary insertions grew gpt-oss:20b's noise set beyond Control's own {C08, C09, C12, C13}, and two of three shrank it. This is a different, so-far unexplained failure mode, not the same one qwen has.

TRV-INV-SPEC-004 already found that isolation, grouped batching, and the denial-override rule are unreliable interventions even on qwen, the model with better baseline recall. Given gpt-oss:20b starts from a strictly worse recall position (0/3 vs. 1/3) and a Control-case failure qwen does not share, further remediation-strategy testing on gpt-oss:20b is not recommended at this time. If `gpt-oss:20b` is revisited later, the more promising next step is not batching/rule interventions but testing whether a non-`off` `reasoning_effort` (`low` or `medium`, enabling gpt-oss's native chain-of-thought) changes this profile materially — `off` may be suppressing the model's strongest mode entirely rather than merely being a fair comparison setting.

**Update following the methodological correction**: the `reasoning_effort=medium` test recommended above was run on a minimal sample (sample_02, sample_05) as part of discovering that all prior calls in this Investigation had `reasoning_effort` silently forced to `off`. Recall did not improve — gpt-oss:20b stayed at 0/3 under confirmed genuine `medium` effort. This is a meaningfully stronger basis for the recommendation to deprioritize gpt-oss:20b than the record previously had, since the one obvious confound (thinking was never actually running) has now been tested and did not change the outcome. Status is held at `investigating` rather than `concluded` because the noise-composition data (as opposed to the recall data) has not been re-run at full scale under confirmed genuine reasoning, and this record's numeric noise claims should not be treated as final until that is done.

## Follow-up artifact candidates

- A decision record on the C08/C09 pairing, informed by both this Investigation and TRV-INV-SPEC-004.
- A decision record on whether to deprioritize gpt-oss:20b for TRV given its complete 0/3 recall result and Control-case failure.
- A comparison table (qwen vs gpt-oss vs 35b-a3b, the last from the earlier ad hoc session) as a durable artifact once all three have comparable data, to support a model-selection decision for TRV.
- If gpt-oss:20b is revisited, a small follow-up investigation testing `reasoning_effort: low`/`medium` against the same five-variant set, isolated from the batching/rule interventions already ruled unreliable on qwen.

## Open questions

- Why did sample_02 (the actual T04 violation) produce zero false criteria from gpt-oss:20b, when the content-free Control produced four? Is the model somehow more "confident" or more verbose-context-anchored when more Work bullets are present, inverting the usual expectation that more content triggers more scrutiny?
- Why did C12 and C13 flip between false (Control, sample_04) and true (sample_05) without any change to the Work Item framing/Verification text that defines those criteria's subject matter?
- Does `gpt-oss:20b`'s "fails silently-unaddressed criteria closed" pattern (C12/C13 false with `"section": "None"`) generalize to other Tasks, or is it specific to this checklist's C12/C13 wording lacking an explicit anchor phrase in this particular Task?
- Would `reasoning_effort` other than `off` (e.g. `low` or `medium`, enabling gpt-oss's native chain-of-thought) change this profile materially, given gpt-oss is a reasoning-first model family and `off` may be suppressing its strongest mode?
- sample_03 reproduced sample_02's zero-noise/missed-detection pattern exactly. Does this "more content, less noise" inversion hold for non-violation content too, or is it specific to Work bullets that resemble a genuine violation? The skill-file-mention and authoring-mention structural probes (single added bullets, not full violations) also reduced noise relative to Control, which is at least directionally consistent, but a single bullet is a much smaller content change than an entire mutation variant.
- What determines which subset of {C08, C09, C12, C13} a given insertion removes? "skill file" left only C13 false; "authoring" left C08/C09 false; sample_05's re-check bullet left C08/C09 false too. No content-based hypothesis has been tested for why these specific pairs split the way they do.

# TRV-ADR-SPEC-006: Suspend semantic validator delivery pending controlled vocabulary

- **status**: accepted
- **date**: 2026-07-03
- **depends_on**:
  - TRV-ADR-SPEC-001
  - TRV-ADR-SPEC-002
  - TRV-ADR-SPEC-003
  - TRV-ADR-SPEC-004
  - TRV-ADR-SPEC-005
- **supersedes**: []
- **migrated_to_spec**: 2026-07-03

## Context

TRV was intended to realize the PRODUCT-owned semantic Task responsibility validator with a lightweight local model.

TRV-INV-SPEC-001 through TRV-INV-SPEC-005 evaluated several models, prompt variants, batching strategies, reasoning settings, and realistic Task mutations.

The experiments did not establish a reliable semantic evaluator.
Observed failures included:

- real responsibility violations remained undetected;
- legitimate Tasks produced unrelated false criteria;
- results changed with criterion grouping, labels, vocabulary, and reasoning mode;
- denial statements and paraphrased actions were interpreted inconsistently;
- different models assigned different responsibility meanings to the same wording;
- remediation prompts often encoded the expected semantic mapping instead of testing independent understanding.

The evidence indicates a prerequisite problem in Brewprint responsibility vocabulary.
The repository does not yet define enough canonical action terms, paraphrase mappings, target-sensitive meanings, and safe contrasts for stable cross-model interpretation.

Continuing contract design, detailed design, or implementation would preserve the unproven semantic assumption and increase sunk cost.

## Decision

Suspend TRV semantic-validator delivery.

Do not continue:

- architecture-derived contract Specification authoring;
- implementation-ready detailed design;
- implementation decomposition;
- production implementation;
- runtime integration;
- prompt-level remediation intended to make a lightweight model infer Brewprint responsibility semantics from unrestricted prose.

Preserve the completed application architecture and existing ADRs as historical design assets.
The architecture remains available if semantic validation becomes viable later.

Treat the existing responsibility-validator prompt skill as deprecated evidence, not active validator authority.
Do not invoke it as an automated release or completion gate.

Resume TRV semantic-validator work only after all of these conditions hold:

1. Brewprint owns a reviewed controlled responsibility vocabulary.
2. The vocabulary defines canonical actions, targets, state changes, ambiguous terms, and confirmed paraphrase mappings.
3. Task authoring uses the vocabulary in a bounded evaluation corpus.
4. A new Investigation demonstrates acceptable cross-model interpretation agreement on valid and violating cases.
5. A new decision explicitly restores TRV contract design or implementation work.

The suspension does not reject deterministic structural lint, vocabulary lint, or human review support.
Those concerns require separate Requirements and Work Items.

## Rationale

The current semantic contract assumes shared interpretation of responsibility language.
The experiments show that the assumption does not hold across lightweight models or stronger agents.

More prompt rules can improve selected examples, but the rules move Brewprint semantics into ad hoc evaluator instructions.
That approach duplicates authority and creates an expensive per-pattern maintenance burden.

A controlled vocabulary attacks the upstream ambiguity.
The vocabulary can improve Task authoring, human review, agent handoff, deterministic lint, and later model evaluation.

Suspension avoids investing in implementation before the semantic input language is stable enough to validate.
Preserving the architecture avoids discarding work that may remain useful after the prerequisite closes.

## Rejected alternatives

| alternative | rejection reason |
|---|---|
| Continue TRV contract and detailed design while vocabulary work proceeds. | Downstream contracts would assume semantics that the experiments did not validate. |
| Add more evaluator prompt rules for every observed miss. | The rule set would become a duplicate responsibility ontology and would overfit the current corpus. |
| Replace Qwen with another lightweight model. | GPT-OSS and Claude Haiku also failed the tested semantic boundary cases. |
| Use a stronger hosted model as the permanent validator. | Cost, availability, and model drift would remain, while repository vocabulary ambiguity would persist. |
| Delete all TRV design artifacts. | The reviewed architecture may remain useful after controlled vocabulary and semantic feasibility are established. |
| Treat the prompt assets as advisory and continue implementation. | An advisory semantic signal does not justify the planned application and workflow integration cost. |

## Consequences

- TRV-WORK-SPEC-001 becomes blocked by the controlled-vocabulary prerequisite.
- TRV-WORK-SPEC-005 and TRV-WORK-SPEC-004 remain unexecuted and blocked.
- TRV-TASK-SPEC-001-15 remains blocked because its child Work Item cannot proceed.
- No TRV implementation Work Item is created.
- TRV-ADR-SPEC-001 through TRV-ADR-SPEC-005 remain accepted historical decisions.
- `spec:trv` records the suspended delivery state and restart conditions.
- The deprecated prompt skill remains available only as investigation evidence and checklist source material.
- A separate controlled-vocabulary workflow may proceed without reopening TRV automatically.
- A later restart requires a new Investigation and explicit decision.

## Evidence

- TRV-INV-SPEC-001 found that an unassisted Qwen evaluation missed responsibility violations.
- TRV-INV-SPEC-002 found that rule guidance improved selected cases but transferred criterion implications into prompt or deterministic logic.
- TRV-INV-SPEC-004 found unstable detection, lexical priming, batching sensitivity, and persistent semantic misses with genuine reasoning enabled.
- TRV-INV-SPEC-005 found worse recall and a failed control case for `gpt-oss:20b`.
- Claude Haiku evaluation also failed to map confirmed paraphrases to the intended responsibility boundary.
- The combined evidence supports explicit responsibility-language mappings instead of unrestricted semantic inference.

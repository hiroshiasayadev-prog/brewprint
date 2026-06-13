# PRODUCT-INV-SPEC-003: Cognitive Load Writing Standard Investigation

- **status**: concluded
- **date**: 2026-06-14
- **trigger**: brewprint spec English-only policy introduced; AI output verbosity and spec readability identified as pain points
- **scope**: Survey of structured writing standards applicable to brewprint design records; identification of adoptable rules for spec prose style and AI output control
- **non_scope**: Bulk rewrite of existing records; full DITA or ASD-STE100 compliance; vocabulary restriction management
- **source_refs**:
  - ASD-STE100 (Simplified Technical English): https://www.asd-ste100.org
  - Plain Language guidelines: https://www.plainlanguage.gov
  - Information Mapping: https://www.informationmapping.com
  - DITA specification: https://docs.oasis-open.org/dita/dita/v1.3/dita-v1.3-part0-overview.html
  - Pyramid Principle (Barbara Minto, 1987)
- **follow_up_candidates**:
  - Cognitive Load Writing Standard as REQ or SPEC (Examples, not exhaustive)
  - Bomb-check rules for AI output in prompt_chappy.md / CLAUDE.md (Examples, not exhaustive)

## 調査スコープ

Survey structured writing standards used in technical documentation.
Identify rules applicable to brewprint design records without full adoption of any single standard.
Separate rules targeting spec prose from rules targeting AI output.

## 非スコープ

- Full DITA adoption
- Full ASD-STE100 compliance
- Vocabulary restriction list management
- Caveman-style ungrammatical English
- Bulk rewrite of existing records

## 背景

brewprint spec was moved to English-only authoring.
AI output (Claude, ChatGPT) tends toward verbosity; long output raises cognitive load and increases risk of unreviewed normative text being accepted.
Spec readability is also impacted by unconstrained prose style.

## 調査したもの

- Information Mapping (Robert Horn)
- DITA (Darwin Information Typing Architecture, IBM)
- ASD-STE100 (Simplified Technical English, aerospace/defense standard)
- Plain Language (plainlanguage.gov, US government)
- Pyramid Principle / BLUF (Barbara Minto / military origin)
- Caveman style (reference only)

## 調査項目ごとの確認結果

### Information Mapping / DITA

- Classifies information into typed blocks: Procedure, Concept, Reference, Fact, Principle.
- brewprint kind / spec / task separation aligns with this philosophy.
- Full adoption not necessary; block type concept is directly usable.

### ASD-STE100

| Area | Finding |
|---|---|
| Vocabulary restriction | Requires managing an approved word list. Not practical for a single-person project. Not adopted. |
| Grammar restriction | 1 sentence = 1 claim, active voice only, max 20 words/sentence, numbered steps for procedures. Adoptable as-is. |
| Technical names | Domain terms are protected from restriction. Examples, not exhaustive: `invoke`, `boundary`, `artifact`, `diagnostic`. |
| Condition branching | STE recommends if/then in prose. Table is lower cognitive load in markdown context. Deviate from STE here. |

### Plain Language

- Core principle: use words your audience already knows, not the simplest possible words.
- For a technical audience, `invoke`, `boundary`, `diagnostic` are known words. Do not replace them.
- Remove bureaucratic language only: `utilize` → `use`, `in order to` → `to`, `it is important to note that` → remove.

### Pyramid Principle / BLUF

- Conclusion before rationale.
- MECE as a criterion for section boundaries.
- BLUF adoptable as a per-section rule.

### Caveman Style (Reference)

| Technique | Adopt | brewprint interpretation |
|---|---|---|
| label-first | Yes | `Rule:` / `Exception:` / `Reason:` / `Evidence:` as bullet prefixes |
| key-value | Yes | metadata / contract / status / scope blocks |
| table-first | Yes | matrix, comparison, state transitions, input/output contracts |
| short bullets | Yes | constraints, non-goals, acceptance criteria |
| broken English | No | grammatical structure preserved |
| extreme omission | No | subject / verb / condition not omitted |

## 横断的な観測事実

- Information Mapping, DITA, Plain Language, ASD-STE100 share the same core principle: chunk by type, label explicitly, keep each block relevant and consistent.
- Grammar restriction is lower-cost than vocabulary restriction for adoption.
- `prefer structured blocks; reserve prose for rationale and causality` is the single highest-leverage rule.
- Condition branching belongs in tables, not if/then prose, in a markdown document context.
- AI output should be treated as a writing standard subject, not only spec prose.

## 後続判断に渡す候補

### Spec-side rules (candidate)

| Rule | Source |
|---|---|
| prefer structured blocks; reserve prose for rationale and causality | Information Mapping / ASD-STE100 |
| BLUF for prose sections only | Pyramid Principle |
| no throat-clearing | Plain Language |
| no duplicate rationale; one-line summary allowed when linking to source record | brewprint traceability model |
| prefer verb over noun | Plain Language |
| 1 sentence = 1 claim | ASD-STE100 |
| prefer active voice; passive allowed when actor is irrelevant or unknown | ASD-STE100 |
| target max 20 words/sentence; longer allowed only when splitting reduces precision | ASD-STE100 / English clarity |
| avoid ambiguous pronouns (this / that / it); repeat the noun when referent is unclear | English clarity |
| do not use more than one of `because`, `which`, `when`, `where`, `while`, `although`, `if` in one sentence; split or use bullets instead | English clarity |
| table for condition branching | deviation from ASD-STE100 |
| table for matrices, state transitions, input/output | Information Mapping |
| bullets for constraints, non-goals, acceptance criteria | Information Mapping |
| label-first bullets (Rule: / Exception: / Reason: / Evidence:) | Caveman style |
| preserve domain terms | Plain Language |

### AI output rules (candidate)

| Rule | Purpose |
|---|---|
| Separate recommendation from normative text | Prevent accidental rule creation |
| Mark examples as `Examples, not exhaustive` | Prevent example-to-rule escalation |
| Use `Candidate:` for undecided rules, not MUST/SHOULD | Prevent premature normative language |
| `verdict` output format for review (PASS / NEEDS REVISION + reason only) | Reduce review output verbosity |

## 推奨案

Introduce writing standard in two layers:

1. **Spec-side**: Start with block type rules and sentence style. Fix section title constraints incrementally via ADR as friction is observed.
2. **AI output**: Add bomb-check rules to prompt_chappy.md and CLAUDE.md separately from spec writing standard.

Do not attempt full adoption of any single standard. Extract only the rules that reduce cognitive load without increasing authoring cost.

## 後続 artifact 候補

- Cognitive Load Writing Standard — REQ or SPEC under product domain (Examples, not exhaustive)
- AI output control rules — prompt_chappy.md / CLAUDE.md update (Examples, not exhaustive)

## 未確定点

- Whether INV section headings should be migrated to English under the spec English-only policy
- Target record kind for the writing standard (REQ vs SPEC)
- Candidate rules may overlap with or conflict with existing authoring guides and specs; cross-check required before adoption

# PRODUCT-REQ-SPEC-012: Define Task type use cases and controlled responsibility vocabulary

- **id**: PRODUCT-REQ-SPEC-012
- **status**: deferred
- **date**: 2026-07-03
- **source_refs**:
  - PRODUCT-REQ-SPEC-011
  - PRODUCT-TASK-SPEC-025-11
  - spec:product.design_records.authoring_standards.task_authoring

## Requirement

Define corpus-grounded use cases and controlled responsibility vocabulary for every accepted `task_type`.

The result must make each Task's intended action, owned object, completion judgment, and adjacent responsibility boundary understandable to humans and lower-capability models without relying on unconstrained synonyms.

## Evidence

- `spec:product.design_records.authoring_standards.task_authoring` defines 11 accepted Task types, their primary outcomes, completion judgments, prohibited overlaps, and selected adjacent boundaries. It does not provide a use-case inventory or intent-specific preferred phrasing.
- PRODUCT-TASK-SPEC-025-11 reorganized 219 corpus findings into `skills/task-vocabulary-reference/`. That artifact records observed phrasing but is not normative authoring authority.
- PRODUCT-TASK-SPEC-025-11 grouped part of the corpus by vocabulary target because some source logs did not retain each source Task's declared type. The resulting reference cannot by itself establish which phrasing is valid for a Task type.
- TRV-INV-SPEC-004 and TRV-INV-SPEC-005 found that tested models missed paraphrased responsibility violations and produced unstable judgments from vocabulary overlap. General synonym breadth is therefore not a reliable responsibility contract.
- Frequently used terms such as `canonical`, `accepted`, `normative`, `owner`, `boundary`, `projection`, `materialize`, `route`, `propagate`, and `synchronize` affect Task-type descriptions but do not yet have one Brewprint-wide controlled definition set.
- The user directed deferring this Requirement until the Brewprint-wide concept vocabulary has been investigated.

Current disposition: `deferred`.

Deferral reason: Task-type use-case definitions would depend on unresolved Brewprint-wide terms and could require substantial rework after those terms are classified or defined.

Restart condition: Resume this Requirement after each foundational cross-cutting term used by the Task-type contract has an accepted definition, a qualified-term split, a retirement decision, or an explicit unresolved disposition.

## Required Outcome

- Each accepted `task_type` has a corpus-grounded inventory of owned use cases.
- Each use case identifies its owned object, expected result, completion judgment, and adjacent responsibilities that must be routed elsewhere.
- Each supported intent maps to preferred wording and at least one preferred sentence shape.
- Ambiguous verbs and nouns map to intent-specific replacements instead of receiving one global meaning.
- Cross-type stop and route signals are explicit.
- Corpus examples cite real Task IDs and remain evidence rather than automatic normative authority.
- The result is consistent with `spec:product.design_records.authoring_standards.task_authoring` and with accepted Brewprint-wide concept definitions.

## Explicitly Excluded Scope

- Defining Brewprint-wide cross-cutting concepts; that work is a prerequisite owned by a separate Requirement.
- Treating every phrase in `skills/task-vocabulary-reference/` as approved vocabulary.
- Bulk rewriting existing Task records.
- Reimplementing TRV or another automated semantic validator.
- Restricting ordinary prose that does not affect responsibility, ownership, completion, routing, or artifact authority.

## Boundary

This Requirement owns Task-type use-case classification and intent-specific responsibility phrasing.

This Requirement does not own the foundational meaning of Brewprint-wide terms, general writing-style rules, Task lifecycle semantics, or validator implementation.

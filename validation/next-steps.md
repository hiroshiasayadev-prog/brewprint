# Validation design next steps

## Status

This document is a working sequence for continuing the design without reopening already settled gaps.

## Immediate next sequence

### 0. Confirm the runtime execution contract slice

The reference and artifact-module design now constrains every later capability input and output, so confirm this bounded runtime slice before resuming the remaining inventory:

1. runtime value and type model;
2. function declaration contract;
3. artifact module manifest and capability slots;
4. module bootstrap, catalog, and routing;
5. design-facing invocation syntax;
6. minimum `if` and bounded `for` semantics.

This is not a general-purpose language design pass.
It only establishes the typed execution concepts already required by concrete validation and reference-routing consumers.

### 1. Finalize the atomic component inventory

Complete one bounded list of capabilities that are genuinely reusable and independently meaningful.

Recommended groups:

- Markdown/source selectors;
- text and scalar transforms;
- collection analysis;
- artifact/base reference facts;
- tree facts;
- graph algorithms;
- generic validation checks.

For each candidate, record only:

| field | purpose |
|---|---|
| capability ID | Stable proposed namespace and name. |
| input | Minimal typed inputs. |
| output | Minimal typed output. |
| owner | `common`, `artifact/base`, `artifact/<kind>`, or `operation/validation`. |
| first consumers | Artifact rules that require it. |
| custom needed | Whether normal composition is insufficient. |

Do not define the full recipe DSL during this step.

### 2. Define artifact/base reusable recipes

After the atomic inventory, identify constraints shared by artifact kinds:

- current-source YAML front matter prohibition;
- H1 count, placement, marker, and title;
- metadata field occurrence and unlisted-field validation;
- sequential H1 identity projection checks;
- sequential filename projection;
- standard H2 presence conditions;
- recommended-heading advisory behavior;
- substantive-content check;
- tree directory `index.md` presence;
- direct self-reference and duplicate checks where parameterization is sufficient.

The base recipe must be parameterized by artifact declarations rather than contain a central artifact-kind switch.

### 3. Map each artifact constraint to capabilities

Create one table per artifact:

| artifact rule | facts/transforms | generic check | classification | custom component |
|---|---|---|---|---|

The mapping should cover every constraint in `artifact-constraint-inventory.md`.

This step is where the remaining uncertainty about duplicate prohibition across all ref collections should be resolved exhaustively.

### 4. Identify the truly custom components

Expected candidates include:

- Work Item provenance-edge normalization;
- Specification Topics-to-physical-child projection;
- Task inherited Work Item sequence comparison;
- field-specific canonical-looking `ref_or_literal` interpretation when base reference facts are insufficient.

A custom component should emit facts or edges where possible. Generic collection or graph checks should still determine equality, duplicates, or cycles.

### 5. Complete the remaining minimal framework contract

After the confirmed runtime execution slice and steps 1 through 4:

- component declaration;
- recipe declaration;
- typed values;
- context dependency declaration;
- output references;
- condition and iteration behavior;
- evaluation result shape;
- static validation;
- namespace dependency rules.

### 6. Route canonical authoring

After the component and mapping design is stable:

- decide ADR boundaries for the declarative validation architecture;
- update artifact Specifications with closed validation mappings where they own semantics;
- add dedicated Validation Specifications for finding codes, subjects, deduplication, and canonical order;
- update operation Specifications only where external projection or operation ownership changes;
- run integrated independent review before implementation planning.

## Known canonical mismatches to resolve

### Requirement metadata

Current file:

```text
drmcp/records/spec/design-records-mcp/artifacts/requirement/h1-adjacent-metadata.md
```

Currently declares `status`, `date`, and `source_refs` as optional with weak formats.

Latest Product authority requires:

- all three persisted fields;
- lifecycle vocabulary for `status`;
- strict date;
- required `source_refs` with empty list allowed.

This is a stale DRMCP artifact Specification, not an unresolved user decision.

### Tree directory `index.md`

Current file:

```text
drmcp/records/spec/design-records-mcp/current-record-model/current-record-scopes.md
```

Currently says broad validation reports a missing directory `index.md` as an advisory.

Latest accepted direction says:

- every tree directory node requires `index.md`;
- absence is a contract violation;
- node and subtree availability remain unchanged;
- the rule is tree-artifact-wide, not Specification-only.

Canonical authoring must update the classification without accidentally changing the current tree-node and navigation model.

### Topics completeness

The Product Topics table contract defines required columns, resolvability, and parent behavior.
The latest validation direction additionally requires exact set equality between declared refs and physical direct children.

Canonical ownership of this stronger completeness rule should be made explicit between:

- Specification artifact validation;
- Product Topics table contract;
- tree/source structure contracts.

## Decisions not to reopen during the next inventory pass

- YAML is a wiring DSL, not a general programming language.
- Built-in components, custom components, and recipes share one invocation model.
- Common capabilities do not own artifact validity or finding classification.
- Canonical artifact ref parsing is not a `common` capability.
- Artifact facts remain validation-independent.
- Artifact validation owns rule meaning and `contract_violation` versus `advisory`.
- Generic validation checks do not know artifact meaning.
- Operation diagnostics and validation findings remain separate.
- Validation findings use `classification`, not `severity`.
- Tree directory `index.md` presence is an artifact/base validation rule.
- Different relation types use separate graphs.
- Generic cycle detection consumes artifact-normalized edges.
- Missing or invalid Requirement status suppresses the conditional Required Outcome finding.

## Details intentionally deferred

Do not spend the next pass deciding:

- the final evaluation JSON or YAML shape;
- dynamic subject-building syntax;
- dedup-key syntax;
- exact `when` or `requires` grammar;
- registry implementation;
- plugin version negotiation;
- retries or parallel execution;
- final finding message templates;
- line and column locator precision.

These details should be derived from concrete component and artifact mapping needs rather than designed speculatively.

## Completion condition for the next design slice

The next slice is complete when:

- every required atomic component has one owner and at least one concrete consumer;
- no component is merely a trivial helper exposed without independent reuse value;
- every artifact constraint maps to reusable facts/checks or a named custom component;
- artifact/base contains only genuinely cross-artifact meaning;
- unresolved framework syntax questions are listed but not prematurely decided;
- known current-spec mismatches remain visible for later canonical routing.

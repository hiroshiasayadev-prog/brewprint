# PRODUCT-TASK-SPEC-013-02: Define portable package interface contract

- **id**: PRODUCT-TASK-SPEC-013-02
- **status**: done
- **date**: 2026-06-26
- **work_item**: PRODUCT-WORK-SPEC-013
- **source_requirement**: PRODUCT-REQ-SPEC-003
- **estimate**: 2d
- **depends_on**:
  - PRODUCT-TASK-SPEC-013-01
- **outputs**:
  - Portable package interface contract recorded in this Task Evidence
  - Fixed package namespace, bundled root, spec roots, ref-prefix rewrite, warning boundary, localized indexing behavior, and producer/consumer split decisions

## Goal

Define one reviewable producer/consumer interface for the portable PRODUCT-owned Design Records standards package.

The interface gives producers and consumers stable roots, namespace, ref-rewrite rules, warning boundaries, and localized indexing behavior for the first package.

It does not define package-management machinery, consumer implementation internals, diagnostic schemas, startup policy, loader algorithms, cache behavior, or app-local runtime behavior.

## Work

- Use the corrected whole-tree source boundary from `PRODUCT-TASK-SPEC-013-01` as the package-content input.
- Read the producer expectations in `PRODUCT-REQ-SPEC-003` and the consumer expectations in `DRMCP-REQ-MCP-003`.
- Define the fixed package namespace `design_records`.
- Define the fixed bundled physical root `<exe-dir>/design-records/`.
- Define the whole-tree copy boundary from `product/records/spec/design-records/`.
- Define package spec tree root `spec:design_records`.
- Define authoring guidance root `spec:design_records.authoring_standards`.
- Define the canonical spec-ref prefix mapping from `spec:product.design_records` to `spec:design_records`.
- Define semantic source defects as producer warnings rather than package-generation blockers.
- Define operational generation failures as the producer's only package-generation errors.
- Define consumer root selection, recursive Markdown discovery, partial-file failure behavior, and package-root unavailability.
- Define localized duplicate, unresolved, and unrewritten-ref behavior without package-wide semantic rejection.
- Define guidance listing from indexable authoring-standards children and exact canonical-ref guidance get.
- Define the producer/consumer responsibility split.
- Record unresolved producer/consumer contradictions as blockers rather than selecting behavior by implementation inference.
- Keep package versioning, identity metadata, manifests, capability declarations, source-state metadata, relative refs, namespace remapping, aliases, redirects, network distribution, and multi-package composition outside the first-version contract.
- Do not choose the copy implementation, tooling placement, loader implementation, parser implementation, cache implementation, diagnostic schema, startup termination behavior, or process-exit behavior in this Task.
- Do not generate the package or modify PRODUCT semantic source specs.

## Done condition

Task Evidence defines one internally consistent portable-package interface that covers namespace, bundled root, whole-tree copy boundary, spec roots, canonical ref-prefix mapping, producer warning versus operational-error boundary, consumer partial loading, localized duplicate/unresolved/unrewritten-ref behavior, guidance list/get behavior, and producer/consumer responsibility split.

Independent review accepted the contract, so this Task's Done condition is satisfied.

## Verification

- Confirm every package-interface expectation in `PRODUCT-REQ-SPEC-003` and `DRMCP-REQ-MCP-003` is either defined here or explicitly assigned to a later producer or consumer Task.
- Confirm the selected namespace and bundled root are independent of Brewprint app registries, repository roots, and DRMCP runtime configuration.
- Confirm package spec tree root and guidance root are deterministic from the package root.
- Confirm the prefix mapping covers canonical spec refs without requiring full-text string replacement.
- Confirm semantic defects are producer warnings, not generation blockers.
- Confirm operational generation failures remain producer errors.
- Confirm the consumer does not package-wide validate semantic closure or auto-repair refs.
- Confirm individual file failure, duplicate canonical IDs, unresolved refs, and unrewritten source-prefix refs have localized effects.
- Confirm `<package-root>/index.md` has no special parse-success gate.
- Confirm guidance list and guidance get behavior are exact canonical package-ref behavior, not basename, filename, path, title, fuzzy, alias, or inferred lookup.
- Confirm producer and consumer responsibilities are separated.
- Confirm prohibited first-version mechanisms are not introduced by implication.
- Confirm no PRODUCT source spec, generated package, implementation, fixture, or unrelated record changed.

## Evidence

### 1. Fixed package roots and namespace

The authoritative source tree is:

```text
product/records/spec/design-records/
```

The package producer copies that whole tree to the bundled installation location:

```text
<exe-dir>/design-records/
```

The logical package root is:

```text
design-records/
```

The fixed package namespace is:

```text
design_records
```

The package spec root is:

```text
spec:design_records
```

The authoring-guidance root is:

```text
spec:design_records.authoring_standards
```

Root resolution must not depend on the process working directory.

### 2. Whole-tree copy and canonical ref rewrite

Package production performs a whole-tree copy. The only semantic transformation is exact canonical spec-ref prefix rewriting:

```text
spec:product.design_records -> spec:design_records
spec:product.design_records.<suffix> -> spec:design_records.<suffix>
```

The rewrite applies only to canonical `spec:` refs, including visible spec `id` values, Topics refs, Related specs refs, body tokens recognized as canonical spec refs, and other canonical spec-ref fields.

The rewrite does not apply to ordinary prose containing `product`, Design Record public IDs such as `PRODUCT-REQ-*`, physical paths, non-canonical examples, arbitrary strings, or refs outside the exact `spec:product.design_records` prefix. Full-text replacement is prohibited.

### 3. Source-authoring responsibility

Source spec authoring owns semantic correctness of `product/records/spec/design-records/`.

Source authoring owns:

- keeping the source tree app-independent;
- moving app wiring and app-local facts to an appropriate separately owned domain;
- correcting external canonical refs;
- correcting duplicate IDs;
- correcting unresolved refs;
- correcting source-prefix and namespace mistakes.

These are not package-producer or package-consumer authoring responsibilities.

### 4. Producer warnings and operational failures

The producer copies the whole source tree, applies the accepted canonical prefix rewrite, may inspect the resulting package, emits warnings for semantic defects during generation or check execution, and does not repair, remove, generalize, filter, or reinterpret source semantics. Warning output is not a persistent producer-owned evidence artifact.

Semantic findings do not block package generation or distribution.

Producer warning conditions include:

- canonical `spec:` refs outside `spec:design_records`;
- unresolved package refs;
- duplicate canonical spec refs;
- unrewritten `spec:product.design_records` refs;
- apparent app-local, wiring, migration, or project-tracking material.

Producer errors are limited to operational generation failures, including:

- source directory cannot be read;
- destination cannot be created or written;
- copy execution fails;
- prefix-rewrite execution fails;
- generation cannot produce the package tree.

T02 does not define exact warning codes, payload schemas, or implementation algorithms.

### 5. Consumer root selection and partial loading

When an explicit configured root is supplied, the consumer uses only that configured root and does not silently fall back to the bundled package.

When no configured root is supplied, the bundled default is:

```text
<exe-dir>/design-records/
```

If the selected package root is absent, unreadable as a directory, or cannot be enumerated, the package is unavailable. T02 does not choose startup exit behavior, process failure behavior, or a capability-degradation protocol.

The consumer discovers Markdown files recursively:

```text
<package-root>/**/*.md
```

Discovery does not require Topics-graph reachability. For an individual Markdown file that cannot be read or parsed, the consumer emits a warning identifying that file, excludes only that file from the index, and continues loading and using other readable files.

`<package-root>/index.md` has no special parse-success gate. It is an ordinary root spec when readable and indexable.

### 6. Package index behavior

A readable spec with a valid canonical ID under `spec:design_records` or `spec:design_records.*` is eligible for the canonical package index.

When multiple documents declare the same canonical package ref, the consumer does not use first-wins or last-wins. It marks only that ref ambiguous. Exact get or resolve for that ref cannot select a document, unrelated unique refs remain usable, and the whole package is not rejected.

When a readable document contains an unresolved canonical ref, the document remains readable. List and get remain available where they do not require that resolution. Only resolution of that ref returns unresolved. The document and package are not rejected.

The consumer does not repair unrewritten source-prefix refs. An unrewritten source-prefix body ref remains unresolved. A document whose visible ID remains under `spec:product.design_records` is not entered into the canonical `spec:design_records` package index. Other documents remain usable and the package remains loaded.

The consumer does not resolve package refs through a host app registry.

### 7. Guidance list and exact get behavior

The package-wide ref index covers:

```text
<package-root>/**/*.md
```

Authoring-guidance list scope covers readable, indexable child specs under:

```text
<package-root>/authoring-standards/
```

equivalently:

```text
spec:design_records.authoring_standards.*
```

The guidance root itself, `spec:design_records.authoring_standards`, is excluded from normal guidance listing and remains available through explicit exact get.

Guidance get accepts only an exact canonical ref, for example:

```text
spec:design_records.authoring_standards.task_authoring
```

The first contract does not support basename lookup, filename lookup, physical-path lookup, title lookup, fuzzy lookup, aliases, or inferred candidates.

### 8. Producer / consumer responsibility matrix

| concern | source authoring | package producer | package consumer |
| --- | --- | --- | --- |
| App-independent semantics | Owns correctness | Warns | Does not repair |
| External canonical refs | Corrects | Warns | Resolve remains unresolved |
| Duplicate canonical IDs | Corrects | Warns | Affected ref is ambiguous |
| Unresolved refs | Corrects | Warns | Affected resolve fails |
| Prefix rewriting | Supplies canonical source refs | Performs exact rewrite | Does not repair misses |
| Whole-tree copy | N/A | Owns | Consumes |
| Individual unreadable file | N/A | May record during generation | Warns, skips file, continues |
| Root selection | N/A | Supplies bundled default | Owns configured override |
| Host registry use | N/A | Package does not require it | Must not use it for package refs |
| Guidance listing | Defines source standards | Copies them | Projects authoring-standards children |
| Semantic defect severity | Corrective authoring concern | Warning | Local runtime effect only |
| Operational generation failure | N/A | Error | N/A |

### 9. Deferred concerns

T02 explicitly defers:

- package versioning;
- package identity metadata;
- manifest design;
- capability declaration;
- source-state metadata;
- diagnostic codes and payloads;
- parser implementation;
- generator implementation;
- loader implementation;
- cache and invalidation implementation;
- startup and process-exit behavior;
- remote distribution;
- aliases and namespace remapping;
- relative refs;
- multi-package composition.

### 10. Decision closure

The accepted interface is minimal and portable:

- source authoring owns semantic correction;
- the producer performs whole-tree copy plus exact canonical prefix rewriting;
- semantic defects become producer warnings and review evidence, not generation blockers;
- operational generation failures remain producer errors;
- the consumer performs operational loading and local indexing only;
- the consumer does not package-wide validate semantic closure;
- the consumer does not auto-repair refs;
- individual file failures are localized;
- duplicate, unresolved, and unrewritten-ref behavior is localized;
- guidance listing is limited to indexable authoring-standards children;
- guidance get is exact canonical package-ref lookup.

Independent review is complete. This interface is accepted input for generation-model and warning-handling follow-up tasks.

### Independent review

- verdict: PASS
- finding result: no blocking, major, or minor findings
- reviewed contract:
  - the package is a whole-tree copy of `product/records/spec/design-records/`
  - the bundled location is `<exe-dir>/design-records/`
  - the fixed package namespace is `design_records`
  - package refs use `spec:design_records`
  - authoring guidance uses `spec:design_records.authoring_standards`
  - package production performs only exact canonical `spec:product.design_records` to `spec:design_records` prefix rewriting
  - source authoring owns semantic correction
  - semantic defects are producer warnings rather than generation blockers
  - operational generation failures remain producer errors
  - consumers perform operational loading and localized indexing without package-wide semantic closure validation or automatic ref repair
  - individual file failures, duplicate refs, unresolved refs, and unrewritten refs have localized effects
  - guidance listing covers indexable authoring-standards children
  - guidance retrieval accepts exact canonical package refs only
  - package self-containment means operational independence from Brewprint, host `product` namespace, host app registry, and process working directory
  - T03 defines human review and optional manual source-authoring Requirement creation; T05 defines warning-emission checks, and T06 warning output need not be retained in Task Evidence
- closure: the T02 Done condition is satisfied and the generation-model and warning-handling follow-up tasks may consume this interface as accepted input

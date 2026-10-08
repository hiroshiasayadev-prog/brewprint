# Validation capability architecture

## Status

This document preserves the current design direction.

It distinguishes:

- confirmed architectural decisions;
- corrections made during discussion;
- provisional runtime details that were explored but intentionally deferred.

It is not yet a canonical framework or validation Specification.

## Design objective

Validation should be assembled from small reusable capabilities rather than implemented as one validator per artifact or embedded directly inside the external validation operations.

The target qualities are:

- artifact rules remain readable and reviewable;
- common algorithms are implemented once;
- adding an artifact kind does not require modifying a central switch statement;
- unusual artifact semantics can use custom components without turning the YAML format into a programming language;
- `validate_scope` and `validate_records` remain projection operations rather than owners of artifact meaning.

## Declarative recipe direction

Artifact validation composition should be expressed in YAML recipes.

YAML is a wiring DSL only. It may:

- select values or facts;
- pass typed inputs to components;
- transform values through named components;
- compose steps;
- apply limited control such as a condition or iteration;
- call built-in components, artifact-specific custom components, or another recipe through one common `use` interface.

YAML should not become a general-purpose expression language.

Complex semantic normalization belongs in a custom component. A custom component should preferably emit reusable facts, values, or graph edges that generic algorithms can consume.

## Capability categories

| category | role |
|---|---|
| built-in atomic component | Small implementation provided by the framework or common capability libraries. |
| custom atomic component | Artifact-specific implementation used when normal selectors and transforms cannot express the required fact. |
| declarative recipe | YAML composition of facts, transforms, checks, conditions, and nested recipes. |

All three are invoked through the same conceptual form:

```yaml
use: <fully-qualified-capability-id>
with:
  <input-name>: <value-or-output-reference>
```

## Confirmed namespace responsibilities

### `framework`

Owns only the minimum mechanism needed to define and run components and recipes:

- common callable-unit interface;
- typed inputs and outputs;
- component and recipe registry;
- step-output references;
- declared execution-context dependencies;
- execution success or failure;
- static recipe validation;
- minimal control constructs.

It does not know artifact kinds, finding codes, status vocabularies, or validation classifications.

### `common`

Owns artifact-independent facts, transforms, collection algorithms, tree traversal, and graph algorithms.

A common capability may:

- extract generic Markdown structures;
- trim or compare text;
- count or compare collections;
- discover duplicates;
- enumerate tree nodes;
- derive direct children or descendants from a supplied tree;
- detect graph cycles from supplied edges.

A common capability must not decide:

- whether a field is required;
- whether an H2 is prohibited;
- whether a value is a valid Requirement status;
- whether a missing item is a contract violation or advisory;
- which relation graph an artifact uses;
- what a canonical artifact reference means.

### `artifact/base`

Owns artifact semantics shared by artifact kinds:

- shared H1 facts and validation recipes;
- H1-adjacent metadata notation and field occurrence facts;
- sequential and tree structure semantics;
- artifact reference grammar and canonical reference interpretation;
- source-placement facts;
- sequential file-name projection validation;
- generic tree-artifact `index.md` requirement;
- shared substantive-content behavior;
- base validation recipes parameterized by artifact declarations.

Reference parsing does not belong in `common` because canonical refs encode artifact kinds, app namespaces, and structure-dependent identity rules.

### `artifact/<kind>`

Owns the meaning of one artifact kind.

Each artifact namespace may contain:

- validation-independent facts;
- custom components for special semantic normalization;
- artifact-specific validation recipes;
- field inventories and permitted target kinds;
- status and task-type vocabularies;
- conditional H2 applicability;
- relation edge recipes;
- finding classification, code, contract ref, and subject mapping.

### `operation/validation`

Owns generic validation operation machinery:

- generic checks over supplied actual and expected values;
- execution of applicable artifact validation contracts;
- conversion of failed evaluations into the external finding envelope;
- finding aggregation, ordering preservation, and operation-limit integration;
- broad and exact validation operation projection.

It does not know Requirement, Task, Topics, `accepted`, or any other artifact-specific meaning.

## Proposed directory structure

```text
capabilities/
├─ framework/
├─ common/
│  ├─ collection/
│  ├─ graph/
│  ├─ markdown/
│  ├─ text/
│  └─ tree/
├─ operation/
│  └─ validation/
│     ├─ check/
│     ├─ execution/
│     ├─ finding/
│     └─ registry/
└─ artifact/
   ├─ base/
   │  ├─ facts/
   │  ├─ reference/
   │  └─ validation/
   ├─ spec/
   │  ├─ facts/
   │  ├─ components/
   │  └─ validation/
   ├─ requirement/
   │  ├─ facts/
   │  ├─ components/
   │  └─ validation/
   ├─ decision/
   │  ├─ facts/
   │  ├─ components/
   │  └─ validation/
   ├─ investigation/
   │  ├─ facts/
   │  ├─ components/
   │  └─ validation/
   ├─ work-item/
   │  ├─ facts/
   │  ├─ components/
   │  └─ validation/
   └─ task/
      ├─ facts/
      ├─ components/
      └─ validation/
```

Empty `components/` directories need not be created until an artifact actually requires custom implementation.

## Dependency direction

The intended conceptual dependency direction is:

```text
framework
  <- common
  <- artifact facts and components
  <- artifact validation recipes
  <- operation validation execution and projection
```

More precisely:

- `common` depends only on `framework`;
- artifact facts may use `framework`, `common`, and artifact/base facts;
- artifact validation recipes may use artifact facts, common facts, and generic operation validation checks;
- operation execution discovers and runs artifact validation recipes but does not define their semantics.

Artifact facts remain validation-independent even when artifact validation recipes call generic validation checks.

## Facts versus validation meaning

The same low-level fact may support several operations.

Example:

```text
metadata occurrences of `status`
```

This fact does not say whether `status` is mandatory or which values are allowed.

The Requirement validation recipe adds that meaning:

```text
Requirement status
  -> exactly one occurrence
  -> scalar string
  -> one of captured / decision_needed / accepted / deferred / rejected
  -> failure classification: contract_violation
```

The separation is:

| layer | owns |
|---|---|
| common or artifact fact | what values and locations exist |
| generic validation check | whether supplied actual values satisfy supplied expectations |
| artifact validation recipe | which expectation applies and what the finding means |
| operation executor | how findings are collected and projected externally |

## Corrected common component boundary

### Candidates that belong in `common`

#### Markdown and source structure

- select H1 data from a parsed source;
- select metadata field occurrences by supplied field name;
- select H2 occurrences by supplied heading;
- enumerate H2 sections;
- enumerate Markdown tables within a supplied body;
- select a table column by supplied column name;
- obtain a section body.

#### Text and values

- trim text;
- exact equality;
- supplied-pattern matching;
- membership in a supplied value set;
- generic value-form classification where the form vocabulary itself is artifact-independent.

#### Collections

- count;
- find duplicates and occurrence locations;
- ordered equality;
- set equality;
- set difference;
- subset;
- contains.

#### Trees

- direct children;
- descendants;
- directory nodes;
- corresponding supplied-record lookup when the tree model is already constructed.

#### Graphs

- directed cycle detection from supplied edges.

Reachability or connected-component algorithms should be added only when a real accepted validation rule requires them.

### Capabilities that do not belong in `common`

- canonical artifact ref parsing;
- artifact kind and app namespace extraction from canonical refs;
- target-kind permission policy;
- path-to-Spec-ref derivation;
- Task identity segment meaning;
- sequential filename projection meaning;
- missing tree `index.md` classification;
- Requirement status vocabulary;
- finding classification and code.

These belong to artifact/base or the applicable artifact namespace.

## Reference capability correction

The earlier candidate `common.reference.parse` was rejected.

Canonical reference grammar is not artifact-independent. It includes:

- record kind;
- app namespace;
- sequential or tree identity structure;
- domain or path segments;
- artifact-specific identity segments.

The current placement direction is:

```text
artifact/base/reference/
├─ parse
├─ canonical-form
├─ segments
├─ resolve
└─ resolution-state
```

Field-specific permitted target kinds and unresolved-reference policy remain under `artifact/<kind>/validation`.

## File-name helper correction

The earlier `common.path.filename-stem` candidate was removed from the capability inventory.

A file stem is simply the file name without its final extension, but exposing every trivial path helper as a capability would create unnecessary inventory noise.

Sequential file validation should operate at artifact/base recipe level using the accepted projection rule:

```text
<PUBLIC-ID>.md
or
<PUBLIC-ID>-<slug>.md
```

## Tree `index.md` ownership correction

The `index.md` requirement is not Specification-specific.

It belongs to tree artifact base validation:

```text
artifact/base/facts/tree-directory-nodes
artifact/base/validation/tree-index-presence
```

The rule is intended to apply to Specification trees and future tree artifacts such as USDM.

Generic tree enumeration remains in `common.tree`.
The meaning that every artifact tree directory requires a corresponding `index.md` remains in artifact/base validation.

## Graph boundary

Generic graph algorithms receive already normalized edges.

| concern | owner |
|---|---|
| cycle detection | `common.graph` |
| Task `depends_on` edge construction | `artifact/task/facts` |
| Decision dependency or supersession edges | `artifact/decision/facts` |
| Work Item provenance edge normalization | `artifact/work-item/components` when normal recipe composition is insufficient |
| whether a detected cycle is invalid | applicable artifact validation recipe |

Different relations remain separate graphs. A dependency edge and a supersession edge are not combined merely because both are directed relations.

## External finding and diagnostic boundary

The current external validation finding envelope is:

```text
classification
code
message
contract_ref
subject
```

`classification` is exactly:

- `contract_violation`;
- `advisory`.

There is no validation `severity` field.

Operation diagnostics are separate from validation findings:

- request-wide or selector-level failures use `error` placement;
- non-failing operation conditions use `warnings` placement;
- diagnostic objects do not contain `severity`;
- diagnostic class is determined by registered code and placement.

## Confirmed checker boundary

A generic check should not directly know a finding code, artifact kind, contract ref, or classification.

The accepted conceptual result states are:

- `passed`;
- `failed`;
- `indeterminate`.

A component or context execution failure is not `indeterminate`; it follows the operation failure path and may become `execution_failure` or `configuration_failure` according to the operation contract.

Artifact validation recipes supply the finding meaning when a check fails.

## Provisional engine details intentionally not frozen

The following ideas were explored and accepted as plausible, but discussion was intentionally pulled back before treating their exact syntax or shape as final:

- exact evaluation object fields such as `actual`, `expected`, `evidence`, and `reason`;
- one evaluation containing multiple failure items;
- static versus dynamically mapped finding subjects;
- exact semantic deduplication keys;
- recipe-order-driven canonical finding order;
- separate `when` applicability and `requires` evaluability constructs;
- exact YAML syntax for failure iteration and subject construction;
- finding message templates;
- registry implementation and plugin/versioning behavior.

These ideas should be reconsidered only after the atomic component inventory and artifact-to-component mapping are complete.

## Minimal framework concepts retained for later design

The framework will probably need these concepts, but exact schemas remain pending:

- typed scalar, list, record, ref, resolution, graph-edge, and truth values;
- explicit source-location or evidence fields in declared record types when a capability requires them;
- declared execution-context requirements such as current records, resolver, or artifact tree;
- recipe step IDs and output references;
- static detection of missing components, bad input names, incompatible types, invalid output references, duplicate step IDs, and recipe cycles;
- minimal condition and iteration support.

Arbitrary expressions, retries, parallel execution, advanced version negotiation, and unrestricted plugins are outside the current design target.

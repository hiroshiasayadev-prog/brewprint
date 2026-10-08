# Executable design open questions

## Status

These questions must be resolved through one complete pilot operation rather than speculative language design.

## Canonical placement

Pending:

- whether executable declarations remain under `validation/` or move to a repository-wide runtime area after a non-validation operation proves the same model;
- whether generated implementation-design Markdown belongs under a canonical Design Record tree or remains build output;
- whether the failed `drmcp/records/spec/implementation` area is retired, archived, or superseded after this model is proven.

Current direction:

- preserve this proposal under `validation/executable-design/` as working notes;
- do not extend the stale implementation architecture as authority;
- do not create handwritten per-operation implementation Specifications before the YAML-first pilot.

## First pilot operation

Pending:

- whether the first complete slice is `discover_record_scopes` or another bounded discovery-and-listing operation.

Selection criteria:

- exercises runtime- and operation-owned contexts;
- requires repository configuration, artifact module catalog use, and filesystem discovery;
- includes request validation and result variants;
- maps operation outcomes to diagnostics;
- reaches at least one narrow filesystem native primitive;
- is bounded enough to complete through native leaves.

## Traceability target granularity

Pending:

- whether `implements` targets whole Specification refs, USDM requirement IDs, requirement records, or new stable responsibility IDs inside operation Specifications;
- how coarse document-level coverage is reported without falsely claiming full implementation;
- how one authority responsibility implemented jointly by several declarations is represented.

Current constraint:

- free-form heading text must not become the only durable identity for required coverage.

## Which declarations require `implements`

Pending:

- every declaration;
- only operation-, domain-, and artifact-facing declarations;
- or every declaration with either `implements` or an explicit infrastructure classification.

The decision must avoid both fake authority refs and untraceable semantic declarations.

## Coverage status

Pending:

- whether missing required coverage is a bootstrap error, design-lint error, or closure-only error;
- how optional, deferred, superseded, and intentionally unsupported authorities are represented;
- whether coverage is checked repository-wide or per executable package/module set.

## Generated document identity

Pending:

- whether generated pages receive canonical Design Record refs;
- whether one Markdown file represents one function/type or groups related declarations;
- how generated filenames remain stable across declaration moves;
- whether generated pages are committed or produced only in CI and local tooling.

Current direction:

- generated identity should derive from stable declaration IDs rather than YAML file paths;
- grouping is a presentation choice and must not alter declaration identity.

## DAG projection

Pending:

- exact graph syntax and whether Mermaid is sufficient for nested `if`, `switch`, and `for`;
- whether the canonical generated representation is a graph, structured pseudocode, or both;
- how to show stub calls and incomplete implementation status without producing misleading complete-looking diagrams.

## Type projection

Pending:

- which local types qualify for display on one function page;
- whether top-level field expansion includes generic instantiations;
- how opaque type native conformance is reported without exposing hidden fields;
- whether transport-only request and response records require direct `implements` refs.

Current direction:

- function pages expand transparent records one level only;
- canonical type pages own the complete top-level field or union-variant list;
- recursive expansion is prohibited.

## Acceptance evidence

Pending:

- declaration format for executable operation scenarios;
- whether scenarios live in YAML alongside declarations or in a separate test declaration kind;
- how one scenario links to authority responsibilities;
- distinction among bootstrap validation, deterministic examples, contract tests, and native unit tests.

Current constraint:

- `implements` coverage alone must not be treated as proof of correct external behavior.

## Native implementation boundary

Pending:

- objective criteria for deciding whether a function remains declarative or becomes native;
- maximum acceptable semantic breadth of one native binding;
- how native Go unit-test coverage links back to the YAML function declaration;
- whether generated output reports native source symbols or only stable binding IDs.

Current direction:

- the YAML declaration owns semantic identity and signature;
- native bindings should be narrow leaves;
- operation policy and artifact-semantic branching stay declarative whenever expressible.

## Incremental and partial generation

Confirmed direction:

- every called function and referenced type must be declared in every validation profile;
- undefined declarations are validation errors and no `allow undefined` mode exists;
- `implementation.stub` with intent `native`, `steps`, or `undecided` is the only supported incomplete function state;
- `stub: native` and `stub: steps` identify a decided final implementation form, while `stub: undecided` preserves an unresolved ownership decision and emits an authoring warning;
- stubs participate in signature, context, call-graph, traceability, and generated-document validation;
- partial generation is required for design review;
- reachable stubs block design closure and executable activation;
- generated output must visibly identify partial status.

Pending:

- final names and command surface for authoring, design-closure, and activation profiles;
- whether design closure permits declared but unavailable native bindings;
- how stale generated output is detected;
- whether the generator emits a single repository-wide gap report or one report per entry function.

The current model is detailed in `partial-design-validation.md`.

## Relationship to USDM tooling

Pending:

- whether existing USDM coverage tools are generalized to declaration coverage;
- whether implementation-design coverage uses the same `covers` direction, the inverse `implements` direction, or a normalized internal relation graph;
- whether USDM requirements become the preferred stable targets for operation implementation coverage.

Current opportunity:

- reuse the established coverage concepts while keeping requirement authority and executable declaration authority distinct.

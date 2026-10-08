# Declaration traceability and coverage format

## Status

This document defines a working proposal for declaration-level traceability.
The exact reference grammar and mandatory-declaration set remain subject to pilot implementation evidence.

## Purpose

Make the relationship between external authority and executable implementation design explicit, queryable, and mechanically checkable.

The primary relation is:

```text
YAML declaration
  -> implements
  -> accepted external responsibility
```

The reverse relation is generated:

```text
external responsibility
  -> implemented by
  -> YAML declarations
```

This permits USDM-like coverage analysis without maintaining a separate handwritten operation-to-YAML mapping table.

## Proposed declaration field

Design-facing declarations may contain an `implements` field.

```yaml
functions:
  operations.get-records:
    detail: >-
      Processes exact current-record selectors and projects one independent
      result for every effective selector.
    implements:
      - spec:drmcp.design_records_mcp.operations.retrieval.exact_record_retrieval
    signature: >-
      (request: GetRecordsRequest) -> GetRecordsResponse
```

The same relation may appear on types, contexts, events, and modules when the declaration itself realizes an externally meaningful responsibility.

```yaml
types:
  ArtifactReferenceParseOutcome:
    detail: >-
      Represents complete artifact-reference parse success or accepted malformed input.
    implements:
      - spec:drmcp.design_records_mcp.current_record_model.identity_conflict_and_addressability
    union:
      - ArtifactReference
      - InvalidArtifactReference
```

## Semantic meaning

`implements` means:

> This declaration exists as part of the executable realization of the referenced responsibility and its declared semantics must remain compatible with that authority.

It does not mean:

- the declaration alone completely satisfies the entire referenced document;
- the referenced external behavior is proven correct;
- a generic helper becomes operation-specific merely because one operation calls it;
- generated Markdown becomes a second source of authority.

## Reference granularity

Document-level refs may be too broad when one Specification contains many independent responsibilities.

The coverage model ultimately needs stable responsibility-level targets, such as:

- USDM requirement IDs;
- canonical requirement records;
- accepted Specification section or rule IDs;
- another stable responsibility identifier explicitly owned by the external authority.

Free-form Markdown headings are not sufficient as durable coverage identities unless the owning format makes them stable identifiers.

Until a responsibility-level vocabulary is accepted, document-level refs may be used provisionally and reported as coarse coverage.

## Direct and transitive coverage

Coverage reports distinguish:

| relation | meaning |
|---|---|
| Direct coverage | A declaration explicitly lists the authority ref in `implements`. |
| Transitive implementation reachability | A directly covering function reaches another declaration through its typed call graph or type graph. |
| Native realization | A declaration is implemented by a resolved native binding. |
| Scenario evidence | An executable scenario exercises the authority through an operation entry function. |

A helper called by a covering operation does not automatically directly implement every authority ref of its caller.

Transitive reachability is generated evidence, not an implicit rewrite of the helper's `implements` field.

## Declaration classification

Forcing fake external refs onto infrastructure declarations would corrupt coverage.

The pilot must determine one of these models:

### Model A: every declaration has authority

Every declaration has a nonempty `implements` list, including runtime and common primitives, because runtime design Specifications provide their authority.

### Model B: externally meaningful declarations only

Operation-, domain-, and artifact-facing declarations require `implements`.
Runtime infrastructure and low-level common declarations may omit it but must be reachable from at least one traced declaration or designated runtime entry.

### Model C: explicit traceability class

Every declaration explicitly chooses one classification:

```yaml
implements:
  - <authority-ref>
```

or:

```yaml
infrastructure: true
```

Model C is the clearest mechanically, but it adds declaration noise. No model is accepted yet.

## Coverage outputs

The validator or generator should produce at least:

### Authority coverage

```text
authority ref
  -> direct declarations
  -> operation entry functions reaching those declarations
  -> executable scenarios providing evidence
```

### Declaration traceability

```text
declaration ID
  -> direct implements refs
  -> callers
  -> callees
  -> implementation kind
  -> native binding when applicable
```

### Gaps

- required authority with no direct declaration;
- declaration with an invalid or unavailable authority ref;
- declaration requiring traceability but lacking it;
- authority covered only by unreachable declarations;
- native binding not reachable from any operation or designated runtime entry;
- generated projection whose source declaration no longer exists.

## Coverage rules

The initial validator should enforce:

- `implements`, when present, is a nonempty duplicate-free ordered list of stable refs;
- every referenced authority target resolves under the configured authority catalog;
- one declaration cannot list the same ref more than once;
- generated coverage is deterministic;
- generated coverage does not infer direct implementation from call reachability;
- coverage validation is separate from signature, context, and execution validation;
- lack of coverage cannot be silently converted into a runtime warning during operation execution.

Whether missing required coverage is a bootstrap error or a design-lint error remains open until declaration classification is decided.

## Type-level coverage

Type coverage is useful when a type itself encodes an accepted semantic distinction, for example:

- malformed versus canonical reference;
- unresolved versus conflicted versus uniquely addressable current identity;
- tree node with or without a corresponding record;
- contract violation versus advisory finding.

A transparent transport DTO that merely mirrors one operation response may instead inherit its rationale from the operation function and not need an independent external requirement ref, depending on the accepted declaration-classification model.

The pilot should avoid assigning refs mechanically to every struct. Type-level `implements` should explain why that type boundary exists.

## Function-level coverage

Function coverage is strongest at semantic boundaries:

- operation entry functions;
- request validation policy;
- current scope selection;
- exact current-record selection;
- tree-node selection;
- diagnostic projection;
- artifact-specific identity or validation behavior.

Low-level primitives should reference their own runtime or artifact authority rather than every operation that consumes them.

## Native binding traceability

Native Go symbols do not independently declare public authority.

The binding chain is:

```text
external authority
  -> YAML function declaration
  -> native binding ID
  -> registered Go implementation
```

The YAML function owns the design-facing ID, detail, signature, context contract, and authority refs.
The Go implementation must conform to that declaration and must not introduce undeclared fields, result states, or operation policy.

# YAML-first executable design workflow

## Status

This document records the current proposed workflow. It does not yet fix the final declaration schema or generated Markdown format.

## Goal

Turn one accepted external operation contract into an executable, typed, traceable implementation design without separately hand-authoring the same design as prose.

The workflow must expose implementation gaps during design rather than after design approval.

## Entry condition

A design slice starts from an accepted external behavior contract, normally one DRMCP operation Specification and the related diagnostic, current-record, and artifact Specifications.

The slice must identify:

- the public operation being realized;
- the authoritative external Specifications;
- request-wide and item-level continuation boundaries;
- expected semantic outcomes;
- trustworthy-result and execution-failure boundaries.

The workflow does not rewrite the external contract.

## Refinement loop

### 1. Declare the operation entry function

Create one typed entry function representing the application operation.

The initial function contains:

- a stable function ID;
- semantic `detail`;
- external responsibilities referenced through `implements`;
- typed input and output;
- required runtime contexts;
- a coarse, source-ordered processing DAG.

Every call target in the first DAG must already have a declaration. Functions whose bodies are not yet designed use a complete typed stub declaration, making unresolved implementation work explicit without losing signature, context, or traceability validation.

### 2. Keep the first DAG semantic

The entry function describes operation responsibilities rather than filesystem mechanics.

Preferred level:

```text
validate request
  -> select current scope
  -> process effective selectors
  -> project operation results
  -> assemble response
```

Premature level:

```text
stat directory
  -> walk files
  -> split path
  -> compare extension
```

Filesystem and parsing mechanics appear only after semantic functions are refined far enough to require native leaves.

### 3. Validate the partial design in authoring mode

The runtime or design validator reports at least:

- unknown functions or types as validation errors;
- explicit stubs as authoring gaps;
- unbound locals;
- signature mismatches;
- unavailable contexts;
- invalid `opens` declarations or active-context routes;
- direct or indirect provider dependency cycles;
- call-graph cycles among declared functions;
- non-exhaustive union switches;
- invalid native bindings;
- missing or invalid traceability references.

Authoring-mode validation permits declared stubs without pretending they are implemented. It preserves full type, context, DAG, and traceability validation, allows partial documentation generation, and prohibits execution or design closure while a reachable stub remains.

Exact stub syntax and validation profiles are defined in `partial-design-validation.md`.

### 4. Refine one stub function at a time

For each stub declaration, choose one of two executable realizations:

```text
composite semantic function
  -> implemented by additional YAML steps

native primitive function
  -> implemented by one narrow native binding
```

A composite function should preserve meaningful operation or artifact language.
A native function should have a small responsibility that can be implemented and tested independently.

Examples of likely native leaves include:

- enumerate one configured source boundary;
- read one source file;
- parse one Markdown structure;
- compare or normalize one primitive value under an accepted contract;
- apply one deterministic collection operation;
- compile and execute one regular expression.

Operation policy, artifact meaning, expected outcome branching, and cross-function sequencing should remain visible in YAML whenever the runtime language can express them directly.

### 5. Introduce types from actual handoffs

Types are declared when a function boundary requires them.

The workflow does not invent a complete domain model before operation flow exists.

For each new type, record:

- semantic `detail`;
- external responsibilities implemented by the type, when applicable;
- top-level fields or union variants;
- whether the type is transparent or opaque;
- which functions construct or consume it.

A generated function page shows only top-level fields for directly used transparent records. Recursive field expansion belongs on the canonical generated type page.

### 6. Preserve expected outcomes as values

Expected semantic states remain typed result variants and are handled explicitly in the operation DAG.

Execution failures remain outside ordinary YAML-visible result branching and propagate to the operation failure boundary.

Operation-specific diagnostic codes are projected from semantic outcomes at the operation or adapter boundary. Shared domain types do not embed public diagnostic codes merely because one operation maps them to those codes.

### 7. Generate review artifacts continuously

At every refinement step, generation should produce:

- operation and function DAGs;
- called-function and `detail` tables;
- involved type summaries;
- native leaf inventory;
- reachable stub inventory;
- external responsibility coverage;
- orphan declaration inventory;
- bootstrap validation results.

This makes design review and implementation-feasibility review the same iterative activity.

### 8. Close the design slice

A slice is closed only when all of the following hold:

- every call target and referenced type is declared;
- every type and module reference exists;
- no reachable stub remains;
- every function has exactly one declarative or native implementation;
- all signatures and field projections type-check;
- all required contexts are statically available;
- all expected unions are handled exhaustively where consumed;
- all native bindings resolve and match their declarations;
- all required external responsibilities are covered by at least one applicable declaration;
- every generated document is reproducible from the declaration set;
- required executable scenarios or tests pass.

Coverage alone does not close a slice when behavior has no executable evidence.

## Suggested first pilot

Use one operation with meaningful selector and continuation behavior rather than a trivial helper.

A retrieval or exact-validation operation is a strong pilot because it forces the design to exercise:

- request validation;
- bounded per-selector iteration;
- ArtifactReference parsing;
- current-record resolution;
- expected outcome switching;
- selector-level continuation;
- source or validation processing;
- operation diagnostic projection;
- response assembly.

The pilot should be completed through native leaves before generalizing declaration-format rules from it.

## Review questions

A reviewer should be able to answer from generated output:

- Which external responsibility caused this function or type to exist?
- Which function invokes it?
- Is its implementation declarative or native?
- What values cross the boundary?
- Which expected outcomes are possible and where are they handled?
- Which runtime contexts are required?
- Which design gaps remain unresolved?
- Which native Go units remain to be implemented?

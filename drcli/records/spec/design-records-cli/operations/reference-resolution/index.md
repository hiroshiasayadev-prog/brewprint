# Contract: Reference resolution

- **id**: `spec:drcli.design_records_cli.operations.reference_resolution`
- **status**: draft
- **date**: 2026-10-07
- **parent**: `spec:drcli.design_records_cli.operations`
- **contract_class**: `interface`

## What this is

Defines the transport-neutral `resolve_reference` operation for exact canonical and PRODUCT-approved compatibility reference resolution.
The operation returns at most one target and remains separate from retrieval and validation.

## Non-goals

- Defining CLI command names, flags, option spelling, or process transport.
- Redefining PRODUCT canonical-reference grammar or compatibility families.
- Fuzzy matching, completion, normalization, alias repair, or path-to-ref conversion.
- Resolving H2 headings or section selectors as canonical references.
- Returning source content, validation findings, or physical source paths.

## Request

The logical request contains one field:

| field | requirement | value |
|---|---|---|
| `ref` | required | Reference candidate evaluated exactly as supplied. |

No additional logical request field is accepted.
A missing or non-string `ref` is an invalid request.
An unsupported additional field is also an invalid request.

A string is not trimmed, repaired, normalized, completed, or case-adjusted.
An empty string or surrounding whitespace therefore remains part of the supplied candidate.

### Reference authority

DRCLI does not define reference grammar independently.

| input family | authority | resolution source |
|---|---|---|
| Current spec ref | `spec:product.design_records.traceability.artifact_refs` and `spec:product.design_records.traceability.semantic_ref` | Selected-repository current state. |
| Current sequential public ID | `spec:product.design_records.traceability.artifact_refs` | Selected-repository current state. |
| Accepted V01 issued ID | `spec:product.brewprint.compatibility.legacy_id_compatibility` | PRODUCT-approved compatibility data available to the selected repository. |

Physical paths are not reference candidates.
Current section-level canonical refs are not active.
No additional semantic prefix is introduced by this operation.

### Resolution order
The operation applies this sequence:

1. Validate the logical request shape.
2. Classify `ref` under current PRODUCT canonical-reference rules.
3. For a current reference, perform one exact lookup in selected-repository current state.
4. Return a resolved current target immediately when exactly one addressable target exists.
5. When the current stage has no resolved target, test exact eligibility under PRODUCT compatibility policy.
6. Return `unsupported` when neither current nor compatibility authority accepts the exact candidate.
7. For an accepted compatibility input, perform exact compatibility lookup.
8. Return the resulting normal outcome without a fuzzy, repaired, or inferred second pass.

A resolved current target stops compatibility evaluation.
A current candidate with no addressable current target may continue to compatibility eligibility only when PRODUCT compatibility also accepts the exact string.

The operation does not scan arbitrary directories to repair an unresolved reference.
The operation does not derive reference targets from natural-language body text, file names, or physical paths.

## Response

A normal response contains:

```text
ref
status
target
```

`ref` is the exact supplied candidate.
`status` uses exactly these values:
| status | target | meaning |
|---|---|---|
| `resolved` | one target | Exactly one selectable target was resolved. |
| `unresolved` | `null` | An accepted input has no selectable target. |
| `unsupported` | `null` | No PRODUCT-owned accepted input family accepts the exact candidate. |

`unresolved` and `unsupported` are normal operation outcomes.
They are not request-shape or execution failures.

### Current target

A resolved current target contains:

| field | presence | value |
|---|---|---|
| `target_type` | always | `current_spec` or `current_sequential_record`. |
| `ref` | always | Resolved canonical current reference. |
| `kind` | always | PRODUCT-owned record kind. |
| `title` | always | Parsed title or `null` when unavailable. |
| `status` | always | Parsed lifecycle status or `null` when unavailable. |

The operation does not repair or default an unavailable title or lifecycle status.
PRODUCT owns the record-kind vocabulary and canonical identity semantics.

### Compatibility target
An accepted V01 issued ID that resolves successfully contains:

| field | presence | value |
|---|---|---|
| `target_type` | always | `legacy_sequential_record`. |
| `ref` | always | Exact issued legacy ID. |
| `kind` | always | PRODUCT-owned sequential record kind represented by the accepted family. |

The compatibility target does not add title, lifecycle status, metadata, headings, or source content.
Compatibility-family acceptance and issued-ID retention remain PRODUCT-owned.

### Outcome distinctions

| condition | normal outcome |
|---|---|
| Exact current ref has one addressable current target | `resolved`. |
| Exact current ref has no current identity claim | `unresolved`. |
| Exact current ref has conflicting current identity claims | `unresolved`. |
| Accepted V01 issued ID has one selectable compatibility target | `resolved`. |
| Accepted V01 issued ID has no selectable compatibility target | `unresolved`. |
| Candidate matches neither current nor accepted compatibility reference rules | `unsupported`. |

The following inputs are not repaired:

- app-prefixless or partial current IDs;
- physical paths;
- fuzzy or incomplete references;
- metadata-only aliases;
- section-like values that are not current canonical refs;
- `yaml:`, `fixture:`, `internal-design:`, or `coverage:` inputs;
- values that require whitespace, case, prefix, domain, or sequence repair.

Such a string returns `unsupported` unless PRODUCT authority independently accepts that exact string.

Normal targets never expose physical paths, source locations, resolver traces, or section anchors.

### Retrieval boundary

Reference resolution does not perform exact record retrieval.
Exact retrieval does not inherit reference-resolution fallback behavior.

Callers use retrieval operations after resolution when record or section content is required.

## Errors

| condition | operation behavior |
|---|---|
| Logical request shape is invalid | Return the DRCLI invalid-request diagnostic; no normal resolver outcome is produced. |
| Selected-repository current state cannot be constructed reliably | Return a DRCLI request-wide failure; no normal resolver outcome is produced. |
| Resolution cannot complete reliably | Return a DRCLI execution failure; no normal resolver outcome is produced. |

Shared diagnostic representation, codes, messages, and machine-readable shape belong to `spec:drcli.design_records_cli.diagnostics`.
This contract owns the normal resolution outcome semantics, not diagnostic serialization.
## Boundary

| concern | owner |
|---|---|
| Canonical record kinds and reference forms | `spec:product.design_records.traceability.artifact_refs`. |
| Current spec-ref model | `spec:product.design_records.traceability.semantic_ref`. |
| PRODUCT lookup sources and invalid reference conditions | `spec:product.design_records.traceability.resolve_and_validation`. |
| Accepted V01 issued-ID families and retention | `spec:product.brewprint.compatibility.legacy_id_compatibility`. |
| Selected-repository addressability | `spec:drcli.design_records_cli.current_record_model`. |
| Current-first orchestration and normal resolver outcomes | This Specification. |
| Record and section content retrieval | `spec:drcli.design_records_cli.operations.retrieval`. |
| Shared failure and advisory representation | `spec:drcli.design_records_cli.diagnostics`. |

## Related specs

| ref | relation |
|---|---|
| `spec:product.design_records.traceability.artifact_refs` | Defines canonical record kinds and current reference forms. |
| `spec:product.design_records.traceability.semantic_ref` | Defines the current path-derived spec-ref boundary. |
| `spec:product.design_records.traceability.resolve_and_validation` | Defines PRODUCT lookup sources and invalid conditions. |
| `spec:product.brewprint.compatibility.legacy_id_compatibility` | Defines accepted V01 issued-ID reference families. |
| `spec:drcli.design_records_cli.current_record_model` | Supplies selected-repository current addressability. |
| `spec:drcli.design_records_cli.operations.retrieval` | Owns exact record and H2 section content retrieval. |
| `spec:drcli.design_records_cli.diagnostics` | Owns shared DRCLI diagnostic representation. |
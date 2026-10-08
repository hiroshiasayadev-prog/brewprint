# Reference: Resolver responsibility

- **id**: `spec:drmcp.design_records_mcp.resolver`
- **status**: draft
- **date**: 2026-06-28
- **parent**: `spec:drmcp.design_records_mcp.overview`

## What this is

Defines DRMCP resolver orchestration for current canonical references and the gate into configured legacy fallback.

PRODUCT specifications own canonical reference grammar and identity semantics.
DRMCP owns grammar evaluation order, index selection, lookup outcomes, and successful target projection.

## Current authority

| concern | authority |
|---|---|
| Current sequential canonical IDs | `spec:product.design_records.namespace_model.artifact_id_grammar` |
| Current spec canonical refs | `spec:product.design_records.spec_format.spec_id_as_ref` |
| Canonical traceability inputs and lookup sources | `spec:product.design_records.traceability.resolve_and_validation` |
| Accepted Brewprint legacy families | `spec:product.brewprint.compatibility.legacy_id_compatibility` |
| Legacy issued-ID lexical parser mapping | `spec:drmcp.design_records_mcp.namespace_scanning` |
| Current identity mapping and addressability | `spec:drmcp.design_records_mcp.schema.id_normalization` and `spec:drmcp.design_records_mcp.schema.record_model` |
| Public resolver operation | `spec:drmcp.design_records_mcp.tools.resolve_reference` |

## Current-first orchestration

`resolve_reference` evaluates one input exactly as supplied.
It does not trim, repair, complete, rewrite, or infer identity segments.

The resolver uses this sequence:

1. Evaluate the input against current canonical grammar.
2. When current grammar accepts the input, query the active index by the exact canonical identity.
3. When the active index returns one addressable current target, return that target and stop.
4. When the current stage produces no resolved target, evaluate accepted legacy-family grammar eligibility.
5. Return `unsupported` only when neither current grammar nor accepted legacy-family grammar accepts the input.
6. When accepted legacy-family grammar matches, continue under the configured legacy fallback contract in this specification.

The current stage produces no resolved target when:

- the input does not match a current canonical grammar; or
- the input matches current grammar but the active index has no addressable target for that exact identity.

A resolved current target prevents legacy grammar evaluation and legacy-index lookup.
The resolver never rewrites a current input into a legacy input.

When one string can satisfy both current and accepted legacy grammar, the resolver still queries the active index first.
Legacy eligibility begins only after the current stage remains unresolved.

Accepted legacy-family recognition combines the PRODUCT-owned accepted family set with the exact issued-ID lexical mapping in `spec:drmcp.design_records_mcp.namespace_scanning`.
Recognition is independent of `legacy_roots` availability.
When accepted legacy grammar matches after the current stage remains unresolved, the resolver uses the configured legacy lookup contract below.

## Configured legacy fallback

Legacy fallback uses the separate read-only exact lookup map defined by `spec:drmcp.design_records_mcp.namespace_scanning`.
It does not query the active index, scan unconfigured directories, normalize archived metadata into the current record model, or invoke `get_records`.

The resolver applies these outcomes after one exact accepted legacy-family match:

| legacy condition | resolver outcome |
|---|---|
| `legacy_roots` is missing or empty. | Return `unresolved` with no target. |
| The lookup map has no source for the issued legacy ID. | Return `unresolved` with no target. |
| Duplicate sources prevent selection of one source. | Return `unresolved` with no target. |
| One indexed source exists but cannot be read. | Return `unresolved` with no target. |
| One readable indexed source exists. | Return `resolved` with one minimal legacy target. |

Operational distinctions between disabled fallback, missing target, duplicate conflict, and unreadable source remain W006-owned diagnostic concerns.
They do not add public resolver statuses.
`unsupported` remains limited to inputs accepted by neither current nor accepted legacy grammar.

A successful legacy target contains exactly:

| field | value |
|---|---|
| `target_type` | `legacy_sequential_record` |
| `ref` | Exact issued legacy ID. |
| `kind` | Record kind derived from the accepted legacy ID family. |

The legacy target does not include title, lifecycle status, metadata, headings, body, physical path, source location, provenance, index state, or resolver trace.
Callers use `get_records` with the returned `ref` when archived source content is required.

## Current input families

| input family | current contract | lookup scope |
|---|---|---|
| Current sequential record ref | Complete app-aware decision, investigation, requirement, work-item, or task ID. | Active index only. |
| Current spec ref | Path-derived document-level `spec:` ref. | Active index only. |

Current spec input classification is lexical.
Any string that matches the current path-derived document-level `spec:` grammar is a current resolver input and queries only the active index.
When that exact identity has no addressable target, the current stage remains `unresolved`.

The resolver does not consult front-matter `semantic_refs`, front-matter `sections`, heading data, or an alias registry to determine where a supplied string originated.
A string historically used as a section alias but lexically accepted by current spec grammar remains a current spec input.
Section targets, headings, and anchors are not separate lookup surfaces.
Physical paths are not canonical resolver inputs.

## Rejected-input boundary

A string that matches neither current canonical grammar nor one accepted legacy issued-ID grammar is `unsupported`.
The resolver evaluates the supplied value exactly and does not transform a rejected value into another candidate.

Rejected input classes include:

- `V01-SPEC-*`;
- app-prefixless sequential IDs;
- physical paths;
- fuzzy or partial references;
- legacy YAML-only alias spellings that fail current canonical grammar;
- direct `yaml:` inputs;
- `fixture:` inputs;
- `internal-design:` inputs;
- `coverage:` inputs;
- `COV-*` inputs;
- section-like `spec:` values that do not match current spec grammar;
- metadata-only alias spellings that fail current canonical grammar;
- values requiring whitespace, case, prefix, domain, or sequence repair.

These values are not repaired or redirected to the active index, legacy lookup map, filesystem scanning, fixture lookup, or validation execution.
A valid request containing one of these strings produces the operation-specific `unsupported` outcome defined by `spec:drmcp.design_records_mcp.tools.resolve_reference`.

Classification does not use semantic origin.
A supplied `spec:` string that matches current spec grammar queries only the active index even when an earlier document used the same string as a section alias.
When that lookup finds no addressable target, the result remains `unresolved`.
The resolver does not perform section-target lookup, heading lookup, or front-matter alias lookup.

An accepted current or accepted legacy input with no selectable readable target remains `unresolved`; it is not reclassified as rejected.

`get_records` applies its own exact classification and warning-trigger contract.
It does not inherit resolver status vocabulary or current-first fallback orchestration.

## Operation separation

`resolve_reference` is a resolution operation.
`get_records` is the sole exact-retrieval operation.

The operations remain separate:

- `resolve_reference` does not call `get_records`;
- `get_records` does not call `resolve_reference`;
- `get_records` does not inherit current-first fallback orchestration;
- resolver target projection does not redefine the `get_records` response.

Lookup-source kinds and public listing kinds may differ.
The resolver operation contract defines its supported target surface without using retired `get_record` or broad `list_records` behavior as authority.

## Current successful target boundary

A successful current target uses canonical identity and parsed record fields.
It does not expose a physical path, source location, index path, or resolver trace.

The public field contract is defined by `spec:drmcp.design_records_mcp.tools.resolve_reference`.
Current spec and current sequential targets remain distinguishable through the target discriminator.

## Delegated boundaries

The following concerns remain outside this specification:

- concrete runtime configuration serialization;
- exact legacy retrieval response projection;
- diagnostic object shape, category, severity, message, and source location;
- warning representation;
- validation execution;
- exceptional physical-path representation.

Legacy-root validation and exact lookup-map construction are defined by `spec:drmcp.design_records_mcp.namespace_scanning`.
Exact archived-source retrieval is defined by `spec:drmcp.design_records_mcp.tools.get_records`.
Diagnostic, validation, source-location, and exceptional path representation belong to `DRMCP-WORK-MCP-006`.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.tools.resolve_reference` | Public request, current lookup, outcomes, and successful current target projection. |
| `spec:drmcp.design_records_mcp.tools.get_records` | Sole exact-retrieval operation. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Configured current roots, legacy-root validation, and separate legacy lookup-map construction. |
| `spec:drmcp.design_records_mcp.schema.id_normalization` | Current canonical identity mapping. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Current addressability and conflict behavior. |
| `DRMCP-WORK-MCP-006` | Diagnostic, validation, source-location, and exceptional path representation owner. |

## Sources

- `DRMCP-ADR-MCP-001`: Current-first resolution and configured legacy fallback.
- `DRMCP-REQ-MCP-001`: Resolver and archive-isolation requirements.
- `DRMCP-TASK-MCP-005-01`: Accepted W005 authority and ownership baseline.
- `DRMCP-TASK-MCP-005-02`: Current-first normative reflection.
- `DRMCP-TASK-MCP-005-03`: Configured legacy lookup and final legacy outcome reflection.
- `DRMCP-TASK-MCP-005-04`: Rejected-input and cross-operation pointer synchronization.

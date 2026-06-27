# Contract: `resolve_reference`

- **id**: `spec:drmcp.design_records_mcp.tools.resolve_reference`
- **status**: draft
- **date**: 2026-06-27
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`resolve_reference` evaluates one reference through current canonical grammar and the active index before any configured legacy fallback.

The operation returns one resolved target at most.
It is separate from exact retrieval and validation execution.

## Request

```json
{
  "ref": "spec:drmcp.design_records_mcp.resolver"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `ref` | yes | string | Reference candidate evaluated exactly as supplied. |

Only `ref` is accepted as a top-level request field.
A missing or non-string `ref`, or an unsupported top-level field, is an `invalid_request` condition.

A string value is not trimmed or repaired.
An empty string or a value with leading or trailing whitespace remains a string-shaped resolver input and does not become a request-shape error.

## Current accepted input families

| input family | authority | lookup scope |
|---|---|---|
| Current sequential canonical ref | `spec:product.design_records.namespace_model.artifact_id_grammar` | Active index only. |
| Current path-derived document-level `spec:` ref | `spec:product.design_records.spec_format.spec_id_as_ref` | Active index only. |

Current sequential refs are complete app-aware canonical IDs for:

- decision;
- investigation;
- requirement;
- work item;
- task.

Bare grammar fragments are not current canonical refs.
The resolver does not infer an app namespace, domain namespace, or sequence segment.

Current spec resolution uses lexical classification against the path-derived document-level `spec:` grammar.
Any supplied string that matches that grammar is a current spec input and queries only the active index.
When the exact identity is absent, the current stage remains `unresolved`.

The resolver does not inspect front-matter `semantic_refs`, front-matter `sections`, heading data, or alias registries to determine the semantic origin of a string.
A string historically used as a section alias but lexically accepted by current spec grammar remains a current spec input.
Section targets, headings, and anchors are not separate lookup surfaces.

The following values are not current lookup keys:

- values from front-matter `semantic_refs` or `sections` that do not independently match current canonical grammar;
- section-like `spec:` values that do not match current spec grammar;
- legacy sequential spec IDs;
- metadata aliases that do not independently match current canonical grammar;
- source paths.

## Current-first resolution sequence

| step | condition | behavior |
|---|---|---|
| 1 | Request shape is valid. | Evaluate `ref` against current canonical grammar exactly as supplied. |
| 2 | Current grammar accepts `ref`. | Query the active index by that exact canonical identity. |
| 3 | One addressable current target exists. | Return `resolved` with the current target and stop. |
| 4 | The current stage has no resolved target. | Evaluate accepted legacy-family grammar eligibility. |
| 5 | Neither current grammar nor accepted legacy-family grammar accepts `ref`. | Return `unsupported`. |
| 6 | Accepted legacy-family grammar accepts `ref`. | Continue under the configured legacy fallback contract below. |

The current stage has no resolved target when current grammar does not accept the input or the active-index lookup has no addressable target.

A resolved current target prevents:

- legacy grammar evaluation;
- configured legacy-index lookup;
- current-to-legacy rewriting;
- a second target selection pass.

When one string can match both current and accepted legacy grammar, active-index lookup runs first.
Legacy eligibility begins only after the current stage remains unresolved.

When accepted legacy-family grammar matches, the resolver queries only the separate configured legacy lookup map defined by `spec:drmcp.design_records_mcp.namespace_scanning`.
It does not scan unconfigured directories, query the active index again, invoke `get_records`, or normalize archived records into the current record model.

## Response

The public outcome vocabulary is fixed to three statuses:

| status | target | condition |
|---|---|---|
| `resolved` | Current or legacy target object. | One addressable current target is found, or one readable legacy source is found after accepted legacy fallback. |
| `unresolved` | `null` | An accepted current or legacy input has no selectable readable target. |
| `unsupported` | `null` | Neither current grammar nor accepted legacy-family grammar accepts `ref`. |

The accepted family set is defined by `spec:product.brewprint.compatibility.legacy_id_compatibility`.
Exact issued-ID lexical recognition uses the parser mapping in `spec:drmcp.design_records_mcp.namespace_scanning`.
Grammar acceptance does not depend on whether `legacy_roots` is configured or usable.

After accepted legacy grammar matches, each of the following returns `unresolved` with `target: null`:

- `legacy_roots` is missing or empty;
- the legacy lookup map contains no source for the exact issued ID;
- duplicate sources prevent one source from being selected;
- one indexed source exists but cannot be read.

One readable source for the exact issued ID returns `resolved` with the legacy target defined below.
The operation does not add statuses such as `disabled`, `unavailable`, or `conflicted`.

Diagnostic and warning fields that distinguish disabled fallback, missing target, duplicate conflict, and unreadable source are owned by `DRMCP-WORK-MCP-006`.
This contract does not define diagnostic object shape, category names, severity, messages, source locations, or exceptional path representation.

## Current successful target projection

A resolved current target uses this non-path projection:

| field | required | type | meaning |
|---|---:|---|---|
| `target_type` | yes | string enum | `current_spec` or `current_sequential_record`. |
| `ref` | yes | string | Resolved canonical current identity. |
| `kind` | yes | string enum | `decision`, `spec`, `investigation`, `requirement`, `work_item`, or `task`. |
| `title` | yes | string or null | Parsed title, or `null` when unavailable on an invalid but addressable source. |
| `status` | yes | string or null | Parsed lifecycle status, or `null` when unavailable on an invalid but addressable source. |

Parsed invalid `title` or `status` values remain unchanged when available.
The operation does not repair, default, or infer those values.
`null` represents unavailable parsed data and does not become normalized metadata.

Current spec example:

```json
{
  "ref": "spec:drmcp.design_records_mcp.resolver",
  "status": "resolved",
  "target": {
    "target_type": "current_spec",
    "ref": "spec:drmcp.design_records_mcp.resolver",
    "kind": "spec",
    "title": "Resolver responsibility",
    "status": "draft"
  }
}
```

Current sequential record example:

```json
{
  "ref": "DRMCP-WORK-MCP-005",
  "status": "resolved",
  "target": {
    "target_type": "current_sequential_record",
    "ref": "DRMCP-WORK-MCP-005",
    "kind": "work_item",
    "title": "Resolver and configured legacy-fallback contract realignment",
    "status": "in_progress"
  }
}
```

## Legacy successful target projection

A resolved legacy target contains exactly:

| field | required | type | meaning |
|---|---:|---|---|
| `target_type` | yes | string enum | Fixed value `legacy_sequential_record`. |
| `ref` | yes | string | Exact issued legacy ID. |
| `kind` | yes | string enum | `decision`, `investigation`, `requirement`, `work_item`, or `task`, derived from the accepted legacy family. |

```json
{
  "ref": "V01-REQ-MCP-001",
  "status": "resolved",
  "target": {
    "target_type": "legacy_sequential_record",
    "ref": "V01-REQ-MCP-001",
    "kind": "requirement"
  }
}
```

The legacy target does not contain `title`, lifecycle `status`, metadata, headings, or body.
Callers use `get_records` with the returned `ref` when archived source content is required.

A normal successful target must not include:

- physical path;
- source location;
- active-index path;
- source provenance;
- section heading or section anchor;
- resolver trace;
- duplicate canonical identity fields such as `record_id` alongside `ref`.

Exceptional path representation remains owned by `DRMCP-WORK-MCP-006`.

## Exact-retrieval boundary

`get_records` is the sole exact-retrieval operation.

- `resolve_reference` does not call `get_records`.
- `get_records` does not call `resolve_reference`.
- `resolve_reference` does not redefine `get_records` request, partial-success, warning, heading, body, or successful-record projection.
- `get_records` does not inherit current-first fallback orchestration.

## Rejected input behavior

The following table applies when the request shape is valid and `ref` is a string.
Each value is evaluated exactly as supplied.

| input class | resolver outcome | operation behavior |
|---|---|---|
| `V01-SPEC-*` | `unsupported` | No current or accepted legacy lookup. |
| App-prefixless sequential ID | `unsupported` | No app, domain, or prefix inference. |
| Physical path | `unsupported` | No path-to-ref conversion or filesystem lookup. |
| Fuzzy or partial reference | `unsupported` | No matching, completion, or second candidate pass. |
| Legacy YAML-only alias spelling that fails current canonical grammar | `unsupported` | No front-matter alias lookup or repair. |
| Direct `yaml:` input | `unsupported` | No YAML lookup surface is active. |
| `fixture:` input | `unsupported` | No fixture lookup surface is active. |
| `internal-design:` input | `unsupported` | No internal-design lookup surface is active. |
| `coverage:` input | `unsupported` | No coverage lookup surface is active. |
| `COV-*` input | `unsupported` | No coverage-ID lookup surface is active. |
| Section-like `spec:` value that does not match current spec grammar | `unsupported` | No repair, section-target lookup, heading lookup, or alias lookup. |
| Metadata-only alias spelling that fails current canonical grammar | `unsupported` | Referring metadata does not register a target alias. |
| Value requiring whitespace, case, prefix, domain, or sequence repair | `unsupported` | No trimming, case repair, completion, or sequence repair. |
| Empty string | `unsupported` | The string remains an item input and is not promoted to a request-shape error. |

These rejected strings are not repaired or redirected to the active index, legacy lookup map, filesystem scanning, fixture lookup, or validation execution.
They produce `unsupported` with `target: null`.

Classification does not use semantic origin.
A supplied `spec:` string that matches current spec grammar queries only the active index, even when an earlier document used the same string as a section alias.
When that exact lookup finds no addressable target, the operation returns `unresolved` with `target: null`.
The operation does not consult front-matter alias registries or perform section-target or heading lookup.

An exact current or accepted legacy input that has no selectable readable target produces `unresolved`, not `unsupported`.
The accepted V01 decision, investigation, requirement, work-item, and task families therefore remain distinct from rejected `V01-SPEC-*` inputs.

Diagnostic category, severity, message, source-location, and exceptional path representation remain owned by `DRMCP-WORK-MCP-006`.

## Errors

| code | condition |
|---|---|
| `invalid_request` | The top-level request violates the request-shape rules. |

Unsupported and unresolved string inputs are normal resolver outcomes, not tool execution errors.
Their diagnostic representation remains delegated to `DRMCP-WORK-MCP-006`.

## Boundary

| concern | owner |
|---|---|
| Current-root discovery, current parsing, canonical identity, active-index construction, and addressability | `DRMCP-WORK-MCP-003` |
| Exact retrieval and `get_records` response | `DRMCP-WORK-MCP-004` |
| Current-first resolver orchestration, configured legacy lookup, final status mapping, and successful non-path target projection | `DRMCP-WORK-MCP-005` |
| Diagnostic and warning representation, validation execution, source location, and exceptional path exposure | `DRMCP-WORK-MCP-006` |

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.resolver` | Resolver responsibility and orchestration summary. |
| `spec:drmcp.design_records_mcp.tools.get_records` | Sole exact-retrieval operation. |
| `spec:drmcp.design_records_mcp.schema.id_normalization` | Current canonical identity mapping. |
| `spec:drmcp.design_records_mcp.schema.record_model` | Current addressability and conflict behavior. |
| `spec:drmcp.design_records_mcp.schema.fields` | Parsed current field vocabulary. |
| `spec:drmcp.design_records_mcp.namespace_scanning` | Configured legacy-root validation and separate exact lookup map. |
| `DRMCP-TASK-MCP-005-02` | Current-first normative reflection owner. |
| `DRMCP-TASK-MCP-005-03` | Configured legacy fallback and final outcome reflection owner. |
| `DRMCP-TASK-MCP-005-04` | Rejected-input and cross-spec pointer synchronization owner. |

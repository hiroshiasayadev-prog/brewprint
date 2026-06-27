# Contract: `validate_records`

- **id**: `spec:drmcp.design_records_mcp.tools.validate_records`
- **status**: draft
- **date**: 2026-06-28
- **parent**: `spec:drmcp.design_records_mcp.tools.overview`
- **contract_class**: `interface`

## What this is

`validate_records` validates selected current Design Records sources, records, duplicate-conflict groups, and declared relations.

This contract owns:

- request selectors and empty-request meaning;
- selection of W003-retained validation inputs;
- startup, configuration, request, execution, and validation-result boundaries;
- current and accepted legacy relation-target lookup behavior;
- the normal response wrapper and execution meaning of `ok`.

PRODUCT specifications own semantic invalidity for metadata, visible document shape, canonical IDs, canonical refs, and declared relation integrity.
This contract cites those authorities rather than copying obsolete V01, hidden YAML front-matter, section-ref, spec-status, or artifact-ID rules.

The diagnostic envelope, category vocabulary, severity vocabulary, deterministic ordering, and deduplication are defined by the diagnostics contract completed under DRMCP-WORK-MCP-006.
Source-location fields and exceptional physical-path exposure are defined separately by the path-exposure contract.

## Request

Repository-wide validation uses an empty object:

```json
{}
```

App-scoped validation uses one configured app namespace:

```json
{
  "app_namespace": "drmcp"
}
```

Exact-record validation uses one current canonical ref:

```json
{
  "ref": "DRMCP-WORK-MCP-006"
}
```

| field | required | type | meaning |
|---|---:|---|---|
| `app_namespace` | no | string | Validate all retained current inputs for exactly one configured current app. |
| `ref` | no | string | Validate one exact current canonical record or spec ref. |

Request rules:

- Empty request validates all configured current roots.
- `app_namespace` and `ref` are mutually exclusive.
- `app_namespace` must identify one configured current app exactly.
- `ref` accepts current canonical refs only.
- Accepted legacy issued IDs are not validation-scope selectors.
- Physical paths are not validation-scope selectors.
- `domain`, `kind`, `status`, `id_range`, one-sided ranges, and list-query filters are not supported.
- Unknown request fields are invalid.
- Selection does not call the public `resolve_reference` operation.

## Validation subjects

| request scope | validation subjects |
|---|---|
| Empty request | Configuration and active-index state for all configured current roots; every discovered current source; parse-failed source; identity-less validation-only source; addressable current record; and duplicate-conflict group. |
| `app_namespace` | The same retained input classes, limited to the selected configured current app. |
| `ref` | The exact current source or addressable record identified by the ref. A conflicted ref selects the complete duplicate-conflict group for that identity. |

Validation subjects are determined directly by request scope.
They are not discovered by traversing relations.

Repository-wide validation checks every retained current subject across all configured current roots.
App-scoped validation checks every retained current subject in the selected app.
Exact-ref validation remains focused on the selected source, record, or duplicate-conflict group.

Identity-less validation-only sources are selectable through repository-wide or app-scoped validation.
They cannot be selected by exact ref because they have no addressable current identity.

## Execution prerequisites

Every configured current root is mandatory.
Every configured legacy root is mandatory when legacy fallback is configured.

`validate_records` cannot produce a normal response when the complete trustworthy input state required by the request cannot be built.
This includes an invalid mandatory current root, an invalid configured legacy root, unavailable required active-index state, or unavailable required configured legacy lookup state.

Execution rules:

- DRMCP does not omit an invalid configured root and continue with a partial active index or partial legacy lookup map.
- App-scoped validation does not bypass a failure in another mandatory configured current root.
- Configuration or index construction failure is not represented as ordinary source, record, or relation invalidity.
- A valid current root with zero discovered records is valid and contributes an empty subject set.
- Missing `legacy_roots` and an explicit empty list are valid and mean legacy fallback is disabled.

## Current relation validation

Every selected validation subject checks its declared current canonical relation targets against the complete active index across all configured current roots.

Current relation rules:

- Lookup uses exact current canonical identity.
- Same-app restriction is prohibited; cross-app current relations are supported.
- One uniquely addressable active-index target satisfies target existence.
- No matching target is unresolved.
- A duplicate-conflict identity has no selected target and does not satisfy target existence.
- Unsupported values are not repaired, prefix-completed, path-interpreted, or fuzzily normalized.
- The public `resolve_reference` operation is not called.
- PRODUCT-owned reciprocity and relation-integrity checks use the same resolved target state when those checks apply.
- A relation target is not recursively added as a validation subject.

A target inside the selected repository-wide or app scope is validated independently because that target is already a selected subject.
A target outside an app-scoped or exact-ref-scoped subject set is lookup-only for the selected subject's relation checks.

## Legacy relation validation

A selected current subject may declare a relation target in an accepted legacy sequential-ID family.
Legacy archive records are lookup targets only and never repository-validation subjects.

| legacy lookup state | relation result |
|---|---|
| Fallback configured; one readable source matches the exact issued ID. | Relation target exists. |
| Fallback disabled because `legacy_roots` is missing or empty. | Relation is not resolved; preserve a distinct disabled-fallback outcome. |
| Fallback configured; no source matches the exact issued ID. | Relation is not resolved; preserve an unresolved-target outcome. |
| Two or more sources produce the same issued ID. | Relation is not resolved; preserve a duplicate-conflict outcome. |
| One indexed source exists but cannot be read. | Relation is not resolved; preserve an unreadable-source outcome. |
| Value does not match accepted legacy-family grammar. | Unsupported relation value; perform no legacy lookup. |
| Configured legacy roots or required lookup state cannot be built. | Startup or repository-validation execution failure; produce no normal response. |

Legacy lookup uses the exact issued-ID mapping and separate configured lookup state defined by the namespace-scanning and resolver contracts.
Validation does not call the public `resolve_reference` operation.
Disabled fallback, unresolved target, duplicate conflict, and unreadable source remain distinguishable inputs to the diagnostics contract.

## Semantic authority

`validate_records` applies semantic invalid conditions from these PRODUCT authorities:

| authority | owned semantics |
|---|---|
| `spec:product.design_records.traceability.resolve_and_validation` | Canonical lookup sources, duplicate identity, and declared relation invalidity. |
| `spec:product.design_records.traceability.metadata_schema` | Current investigation and workflow relation fields. |
| `spec:product.design_records.spec_format.document_shape` | Visible spec document shape. |
| `spec:product.design_records.spec_format.spec_id_as_ref` | Path-derived current spec identity and parent-ref grammar. |
| `spec:product.design_records.spec_format.validation_policy` | Spec-format validation policy and migration-sensitive severity inputs. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Current sequential artifact ID and canonical record-ref grammar. |

DRMCP owns parsing, indexing, request handling, concrete diagnostics, and response representation.
PRODUCT owns whether parsed content or declared relations are semantically valid.

## Response

A successfully started validation execution returns one normal response wrapper:

```json
{
  "ok": true,
  "scope": {
    "app_namespace": "drmcp"
  },
  "summary": {
    "sources": 42,
    "addressable_records": 39,
    "validation_only_sources": 1,
    "conflict_groups": 1
  },
  "diagnostics": []
}
```

| field | type | meaning |
|---|---|---|
| `ok` | boolean | `false` when at least one diagnostic has the diagnostics contract's blocking severity; `true` otherwise. |
| `scope` | object | Effective repository-wide, app, or exact-ref selector. |
| `summary` | object | Counts of selected current validation subjects. |
| `diagnostics` | array | Unified source, record, duplicate, current-relation, and legacy-relation findings. |

`scope` forms:

| request | response `scope` |
|---|---|
| `{}` | `{}` |
| `{ "app_namespace": "drmcp" }` | `{ "app_namespace": "drmcp" }` |
| `{ "ref": "DRMCP-WORK-MCP-006" }` | `{ "ref": "DRMCP-WORK-MCP-006" }` |

`summary` fields:

| field | meaning |
|---|---|
| `sources` | Number of selected discovered current sources. |
| `addressable_records` | Number of selected uniquely addressable current records. |
| `validation_only_sources` | Number of selected sources retained for validation without an addressable current identity. |
| `conflict_groups` | Number of selected duplicate-identity conflict groups. |

Relation-target lookup counts are not included in `summary`.
Validation does not add a separate top-level `warnings` collection.
Non-blocking findings remain diagnostics under the T03 severity vocabulary.

`ok` is `true` when no blocking diagnostic exists, including when `diagnostics` is empty.
`ok` is `false` when at least one blocking diagnostic exists.
Request errors and startup or execution failures do not return this normal wrapper.

## Errors

| condition | handling |
|---|---|
| Malformed request object or field type. | Request error; no normal validation response. |
| `app_namespace` and `ref` supplied together. | Request error; no normal validation response. |
| Unsupported or unknown selector supplied. | Request error; no normal validation response. |
| Selected `app_namespace` is not configured. | Request error; no normal validation response. |
| Selected `ref` does not match current canonical grammar. | Request error; no normal validation response. |
| Selected current canonical `ref` has no addressable source, record, or duplicate-conflict group. | Request error; no normal validation response. |
| Mandatory configuration or required index state cannot be built. | Startup or execution failure; no normal validation response. |
| Valid configured app contains zero retained validation subjects. | Successful validation with an empty selected subject set. |
| Selected source, record, duplicate group, current relation, or legacy relation is invalid. | Validation diagnostic inside the normal response. |

Exact request-error codes, diagnostic categories, diagnostic severity, ordering, and deduplication remain owned by the diagnostics contract.
Unsupported relation values inside a selected record are validation diagnostics, not request errors.
An unresolved exact `ref` selector is not converted into `ok: true` with an empty diagnostic set.

## Related specs

| ref | relation |
|---|---|
| `spec:drmcp.design_records_mcp.namespace_scanning` | Configured roots, retained current inputs, active-index construction, and separate legacy lookup state. |
| `spec:drmcp.design_records_mcp.resolver` | Current-first and accepted legacy lookup orchestration boundary. |
| `spec:drmcp.design_records_mcp.tools.resolve_reference` | Public resolver contract that repository validation does not invoke. |
| `spec:drmcp.design_records_mcp.schema.diagnostics` | Diagnostic envelope and vocabulary; realigned by the next W006 task. |
| `spec:product.design_records.traceability.resolve_and_validation` | PRODUCT-owned lookup and invalidity semantics. |
| `spec:product.design_records.namespace_model.artifact_id_grammar` | Current sequential artifact ID grammar. |

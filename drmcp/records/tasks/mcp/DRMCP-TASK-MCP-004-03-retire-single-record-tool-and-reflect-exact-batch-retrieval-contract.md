# DRMCP-TASK-MCP-004-03: Retire single-record tool and reflect exact batch-retrieval contract

- **id**: DRMCP-TASK-MCP-004-03
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-004
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-004-02
- **outputs**:
  - spec:drmcp.design_records_mcp.tools.get_record
  - spec:drmcp.design_records_mcp.tools.get_records

## Goal

Retire `get_record` from the public tool surface.

Establish `get_records` as the sole exact-retrieval operation with ordered partial success, request-wide body inclusion, and no normal physical-path exposure.

## Work

- Rewrite `spec:drmcp.design_records_mcp.tools.get_record` as an explicit retirement marker.
- Remove every current claim that `get_record` is invokable or retained as a compatibility alias.
- Rewrite `spec:drmcp.design_records_mcp.tools.get_records` around required `refs` and optional request-wide `include_body`.
- Define request-shape rejection for missing, non-array, empty, over-limit, non-string, unknown-field, and invalid `include_body` inputs.
- Define accepted exact current sequential refs, active `spec:` refs, and configured legacy sequential IDs through their owning authorities.
- Prohibit prefix inference, case repair, whitespace repair, path interpretation, fuzzy matching, partial completion, and family guessing.
- Define exact-string ordered deduplication, first-occurrence record order, partial success, and top-level warning placement.
- Return successful records only with canonical `ref`, normalized parsed metadata, headings, and conditional verbatim body.
- Omit missing parsed metadata fields instead of fabricating values.
- Exclude physical paths, source provenance, resolver traces, and internal index state from normal retrieval records.
- Recheck `schema.fields`, `schema.record_model`, and `schema.record_source` for required operation pointers.
- Synchronize `DRMCP-WORK-MCP-004` with the Task link and current execution evidence.

Changed-file scope:

- create this Task;
- update `DRMCP-WORK-MCP-004` task linkage and Evidence;
- rewrite `spec:drmcp.design_records_mcp.tools.get_record`;
- rewrite `spec:drmcp.design_records_mcp.tools.get_records`;
- change shared schema files only when an existing pointer is insufficient;
- change no tools overview, MVP scope, responsibility boundary, resolver, diagnostics, validation, fixture, implementation, or automated-test file.

Ownership exclusions:

- W003 retains current-root discovery, source parsing, canonical identity, active-index construction, normalized record fields, duplicate-conflict construction, and invalid-source retention.
- W005 retains resolver invocation, current-first resolution, configured legacy fallback, fallback ordering, and current/legacy lookup orchestration.
- W006 retains warning and diagnostic category names, severity, shared warning fields, source-location representation, validation behavior, and exceptional path exposure.
- T04 retains tool-catalog, MVP-scope, responsibility-boundary, and broader shared-pointer synchronization.

## Done condition

- `get_record` cannot be read as an available public tool or compatibility alias.
- `get_record` identifies `get_records` with one `refs` element as the replacement.
- `get_records` accepts only `refs` and `include_body` as top-level request fields.
- `refs` is required, is an array, contains 1 through 20 strings, and is not repaired.
- `include_body` is boolean and defaults to `false`.
- Accepted input families point to current sequential, active spec, and PRODUCT-owned configured legacy authorities.
- Exact retrieval does not invoke the resolver or infer another reference family.
- Duplicate input uses exact-string first-occurrence deduplication and triggers a top-level warning.
- Malformed, unsupported, unresolved, unavailable legacy, and duplicate inputs do not discard retrievable records.
- The response contains successful `records` only and top-level `warnings`.
- Successful record order follows first occurrence in the request.
- Every successful record contains canonical `ref`, normalized parsed metadata, and headings.
- `body` is omitted when `include_body` is `false`.
- `body` contains complete source Markdown verbatim when `include_body` is `true`.
- Missing parsed metadata fields remain omitted; parsed invalid values remain unchanged.
- Normal retrieval records contain no physical path, source provenance, internal index state, or resolver trace.
- W003, W005, W006, and T04 ownership boundaries remain intact.
- Scoped strict validation passes.
- Independent review reports no blocking or major findings.
- Required review corrections are applied and recorded before Task closure.

## Verification

- Compare the rewritten operations against `DRMCP-TASK-MCP-004-01` and the accepted T03 baseline.
- Compare record field availability against `spec:drmcp.design_records_mcp.schema.fields` and `spec:drmcp.design_records_mcp.schema.record_model`.
- Compare heading and body source availability against `spec:drmcp.design_records_mcp.schema.record_source`.
- Compare accepted legacy-family authority against `spec:product.brewprint.compatibility.legacy_id_compatibility` without restating its grammar.
- Confirm no `get_record` availability, `ids`, failure-item wrapper, `retrieval_status`, `record: null`, physical-path response, V01 spec input, resolver invocation, or fuzzy repair claim remains in the changed operation specs.
- Confirm warning placement does not define W006 category names, severity, shared fields, or source-location schema.
- Run the scoped strict spec validator and `git diff --check`.
- Obtain independent review before changing the Task status to `done`.

## Evidence

### Accepted contract input

| concern | accepted input |
|---|---|
| Public surface | `get_record` is retired. `get_records` is the only exact-retrieval operation. |
| Request | Required `refs`; optional request-wide `include_body`; 1 through 20 string elements. |
| Exact current input | Current sequential canonical refs and active path-derived `spec:` refs query the active index only. |
| Exact legacy input | Legacy sequential IDs accepted by PRODUCT compatibility authority query the configured legacy index only. |
| No repair | No prefix inference, case or whitespace repair, path interpretation, fuzzy matching, partial completion, or family guessing. |
| Ordering | Successful records follow first-occurrence request order. |
| Duplicates | Exact-string first occurrence wins; later occurrences trigger warnings and return no duplicate record. |
| Partial success | Malformed, unsupported, unresolved, unavailable legacy, and duplicate inputs do not discard retrievable records. |
| Response | Successful `records` only plus top-level `warnings`. |
| Record projection | Canonical `ref`, parsed normalized metadata, headings, and conditional verbatim body. |
| Path boundary | No physical path, source provenance, internal index state, or resolver trace in normal record responses. |

T01 fixed these inputs. T03 does not reopen them.

### Changed files

- Created `drmcp/records/tasks/mcp/DRMCP-TASK-MCP-004-03-retire-single-record-tool-and-reflect-exact-batch-retrieval-contract.md`.
- Rewritten `drmcp/records/spec/design-records-mcp/tools/get-record.md` as a retirement marker.
- Rewritten `drmcp/records/spec/design-records-mcp/tools/get-records.md` as the sole exact batch-retrieval contract.
- Updated `drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md` with the T03 link and execution evidence.

### Shared schema recheck

- `schema.fields` already delegates public retrieval representation to W004.
- `schema.record_model` already delegates omission, null, warning, heading, body, and path representation to W004.
- `schema.record_source` already defines headings and body as source material and delegates public inclusion to W004.
- The rewritten operation contracts revealed no missing pointer. No shared schema file changed.

### Ownership exclusions

- W003 remains authoritative for discovery, parsing, identity, normalized fields, addressability, conflict state, and invalid-source retention.
- W005 remains authoritative for resolver and legacy-fallback orchestration.
- W006 remains authoritative for warning entry schema, category names, severity, source locations, and exceptional path exposure.
- T04 remains authoritative for catalog and broader cross-spec synchronization.

### Removed stale claims

- Public single-record `get_record` availability and compatibility alias behavior.
- `ids` as the batch request field.
- Unbounded batch count.
- `V01-SPEC-*` as an exact input example.
- `items` as a mixed success and failure collection.
- `retrieval_status`, `record: null`, and per-item diagnostic wrappers.
- Physical paths in successful record responses.
- `get_record` as the shared response authority.
- Exact retrieval through resolver normalization or family guessing.

### Scoped stale-claim confirmation

A scoped search over `tools/get-record.md` and `tools/get-records.md` found no positive stale claim for `ids`, `V01-SPEC-*`, path-bearing examples, `record_not_found`, legacy item diagnostics, or `get_record` retrieval behavior.

The only match for a removed response token is the normative sentence that prohibits `items`, `retrieval_status`, `record: null`, and per-item wrappers.

### Unresolved decisions

None.

### Validation and review state

Scoped strict validator command:

```powershell
python -X utf8 product/src/tools/validate_spec.py `
  drmcp/records/spec/design-records-mcp/tools/get-record.md `
  drmcp/records/spec/design-records-mcp/tools/get-records.md `
  --strict `
  --no-color
```

Whitespace and patch-integrity command:

```powershell
git diff --check -- `
  drmcp/records/tasks/mcp/DRMCP-TASK-MCP-004-03-retire-single-record-tool-and-reflect-exact-batch-retrieval-contract.md `
  drmcp/records/work-items/mcp/DRMCP-WORK-MCP-004-query-and-exact-retrieval-contract-realignment.md `
  drmcp/records/spec/design-records-mcp/tools/get-record.md `
  drmcp/records/spec/design-records-mcp/tools/get-records.md
```

- Scoped strict validator result: `[strict]  All 2 file(s) OK.`
- `git diff --check` result: PASS; no whitespace errors reported.
- Git emitted LF-to-CRLF working-copy warnings for `tools/get-record.md` and `tools/get-records.md`.
- The line-ending warnings are non-blocking because `git diff --check` reported no error.
- Independent review verdict: `PASS`.
- Blocking, major, minor findings, and advisories: none.
- Review confirmed the scoped validator and `git diff --check` evidence as valid.
- Review classified the LF-to-CRLF working-copy warnings as non-blocking.
- T03 closure readiness: `ready`.
- Task closed as `done` on 2026-06-27.

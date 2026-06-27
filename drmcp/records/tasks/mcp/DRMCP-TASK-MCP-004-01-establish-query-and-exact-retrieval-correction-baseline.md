# DRMCP-TASK-MCP-004-01: Establish query and exact-retrieval correction baseline

- **id**: DRMCP-TASK-MCP-004-01
- **status**: done
- **date**: 2026-06-27
- **work_item**: DRMCP-WORK-MCP-004
- **source_requirement**: DRMCP-REQ-MCP-001
- **estimate**: 1d
- **depends_on**:
  - DRMCP-TASK-MCP-003-05
- **outputs**: []

## Goal

Establish the accepted query and exact-retrieval correction baseline for `DRMCP-WORK-MCP-004`.

Confirm the affected-file manifest, W003 inputs, and W005/W006 exclusions before normative spec edits begin.

## Work

- Read the accepted ADR, source Requirement, contract audit, hub Work Item, W003 output, W005 boundary, and W006 boundary.
- Inspect the current `list_records`, `get_record`, and `get_records` operation contracts.
- Inspect the tool catalog, MVP scope, responsibility summary, and shared record-model pointers.
- Classify each candidate baseline claim as accepted, clarified, stale, delegated, deferred, or decision-required.
- Record exact files that later Tasks must rewrite, synchronize, recheck, or leave to another Work Item.
- Separate operation warning triggers from diagnostic category, severity, and source-location representation.
- Separate exact retrieval from current-first reference resolution and legacy fallback orchestration.
- Record unresolved response-representation decisions one at a time for user judgment.

This Task does not edit normative DRMCP specs.
It establishes the manifest and decision boundary used by T02 through T05.

## Done condition

- Every candidate baseline item has a recorded classification.
- Every affected or delegated spec has a recorded disposition.
- `get_record` retirement and `get_records` ownership are explicit.
- Normal listing kinds and spec/legacy exclusions are explicit.
- Exact retrieval and resolver ownership are separated.
- Warning triggers and diagnostic taxonomy ownership are separated.
- Normal response path hiding and W006 exceptional path ownership are explicit.
- T02 through T05 can proceed without reopening W003 decisions.
- No unresolved T01-level decision remains.

## Verification

- Compare the baseline against `DRMCP-ADR-MCP-001` and `DRMCP-REQ-MCP-001`.
- Compare active-index and source-material assumptions against final W003 contracts.
- Compare exact retrieval ownership against W005 resolver scope.
- Compare warning and path-exposure ownership against W006 scope.
- Search the scoped DRMCP spec tree for stale `get_record`, `id_range`, response-path, and batch-item claims.
- Confirm that no normative DRMCP spec changed during this Task.

## Evidence

### Authority and upstream baseline

| concern | authority or input | T01 use |
|---|---|---|
| Current-format-first read baseline | `DRMCP-ADR-MCP-001` | Accept active-only normal listing, exact accepted legacy retrieval, no fuzzy repair, and normal path hiding. |
| Required query and retrieval outcomes | `DRMCP-REQ-MCP-001` | Require explicit scope, ordering, duplicate handling, partial success, and exact lookup behavior. |
| Audit evidence | `DRMCP-INV-MCP-002` | Use F06-F08 and F19 plus the file-disposition inventory as historical gap evidence. |
| Active-index and current identity contract | `DRMCP-WORK-MCP-003` | Consume configured current roots, separate indexes, complete current IDs, path-derived spec refs, and invalid-source retention. |
| Resolver and legacy fallback | `DRMCP-WORK-MCP-005` | Delegate current-first grammar evaluation, fallback order, legacy-root orchestration, and issued-ID preservation. |
| Diagnostics and path exceptions | `DRMCP-WORK-MCP-006` | Delegate category names, severity, shared diagnostic fields, source locations, and exceptional path exposure. |
| Semantic fields | `spec:drmcp.design_records_mcp.schema.fields` | Consume parsed common and kind-specific fields without redefining PRODUCT semantics. |
| Addressability and source provenance | `spec:drmcp.design_records_mcp.schema.record_model` | Consume valid, invalid-but-addressable, identity-less, and duplicate-conflict states. |
| Heading and body source availability | `spec:drmcp.design_records_mcp.schema.record_source` | Treat headings and body as source material only; W004 decides public inclusion. |

W003 decisions are fixed inputs.
T01 does not reopen discovery, parsing, identity, index construction, or source-retention behavior.

### Candidate baseline classification: `list_records`

| candidate claim | classification | authority and clarification |
|---|---|---|
| Query the active index only. | accepted and authoritative | Required by the ADR, Requirement, and W003 current/legacy separation. |
| Require exactly one active app scope. | accepted and authoritative | Replaces omitted namespace scope and cross-app broad listing. T02 defines the request field name. |
| Require one supported sequential kind. | accepted and authoritative | Normal listing is kind-specific. Mixed-kind listing is removed. |
| Require one domain scope. | accepted and authoritative | Prevents implicit cross-domain listing and range-derived scope. T02 defines the request field name. |
| Support decision, investigation, requirement, work item, and task. | accepted and authoritative | These are the supported sequential normal-listing kinds. |
| Exclude spec from normal listing. | accepted and authoritative | Specs remain exact-retrieval targets through active path-derived refs. |
| Exclude legacy archive records from normal listing. | accepted and authoritative | Legacy records remain confined to the separate legacy index and accepted exact operations. |
| Allow an optional status filter using the selected kind vocabulary. | retained with clarification | Status is meaningful only after app, kind, and domain scope are fixed. Request validation belongs to T02; PRODUCT owns vocabularies. |
| Order by canonical ID. | accepted and authoritative | Existing mixed-kind and path tie-break rules are removed. |
| Default to descending order; allow ascending only. | accepted and authoritative | `order_by` becomes unnecessary because canonical ID is the sole ordering key. |
| Default limit to 20; allow 1 through 100. | accepted and authoritative | Unbounded listing and omitted hard caps are rejected. |
| Return compact `ref`, `title`, `status`, and `date`. | accepted and authoritative | `ref` carries the canonical current identity. All four fields remain present. Missing `title`, `status`, or `date` values are `null` for invalid-but-addressable records. Kind and physical path are not returned. |
| Return `has_more`. | accepted and authoritative | Replaces range-based narrowing and signals truncation without exposing total repository state. |
| Return an empty result for a valid zero-match query. | accepted and authoritative | Zero matches do not produce a warning or error. |
| Exclude physical path from the normal response. | accepted and authoritative | W006 owns narrow diagnostic, patch, and debug exceptions. |
| Support exact `id`, `id_range`, one-sided range, or broad omitted filters. | stale and must be replaced | Current `list-records.md` behavior conflicts with the accepted scoped compact query. |
| Return full common and kind-specific metadata objects. | stale and must be replaced | Normal listing returns only the compact projection. |
| Resolve duplicate IDs by path ordering. | stale and must be replaced | W003 defines no winner for a conflicted canonical identity. |

### Candidate baseline classification: exact retrieval

| candidate claim | classification | authority and clarification |
|---|---|---|
| Remove public `get_record`. | accepted and authoritative | T03 retires the operation and removes it from the public catalog and MVP surface. |
| Use `get_records` as the sole exact retrieval operation. | accepted and authoritative | The operation owns current, spec, and accepted configured-legacy exact retrieval. |
| Accept 1 through 20 refs per request. | accepted and authoritative | Missing, empty, non-array, non-string, and over-limit input are invalid requests. Exact diagnostic representation remains outside T01. |
| Accept current sequential canonical refs. | accepted and authoritative | Lookup uses the active index only. |
| Accept active path-derived `spec:` refs. | accepted and authoritative | Lookup uses the active index only. |
| Accept approved legacy sequential IDs only when legacy roots are configured. | accepted and authoritative | Lookup uses the legacy index only. W005 owns root configuration and fallback orchestration, not exact retrieval response behavior. |
| Do not infer an app prefix. | accepted and authoritative | App-prefixless bare IDs are not accepted exact refs. |
| Do not repair case or whitespace. | accepted and authoritative | Exact inputs remain byte-for-byte lookup candidates after request-shape validation. |
| Do not interpret filesystem paths. | accepted and authoritative | Paths are neither canonical refs nor exact retrieval inputs. |
| Do not perform fuzzy normalization or invoke the resolver. | accepted and authoritative | Exact retrieval remains distinct from `resolve_reference`. |
| Preserve first-occurrence request order. | accepted and authoritative | Successful records follow ordered first-occurrence deduplication. |
| Deduplicate repeated refs and issue a warning trigger. | accepted and authoritative | W004 owns the trigger and placement. W006 owns category name, severity representation, and shared fields. |
| Continue after unresolved, malformed, unsupported, or unavailable legacy refs. | accepted and authoritative | Each failed ref produces a warning trigger; retrievable records remain successful. |
| Return successful records only. | accepted and authoritative | Existing `items` entries with `retrieval_status: not_found` and `record: null` are removed. |
| Apply `include_body` to the entire request. | accepted and authoritative | Per-item body selection is unsupported. Default is `false`. |
| Include normalized metadata and headings for each successful record. | retained with clarification | W003 supplies parsed values and addressability. T03 must define missing-field representation for invalid-but-addressable records. |
| Include body only when requested. | accepted and authoritative | Body is returned verbatim and is not summarized, reformatted, or truncated by the operation. |
| Exclude physical path from normal retrieval responses. | accepted and authoritative | W006 owns source-location diagnostics and explicit exceptional surfaces. |
| Return one item for every requested ref, including failures. | stale and must be replaced | Existing mixed success/failure item wrappers conflict with successful-record-only output. |
| Use `get_record` representation as the shared response contract. | stale and must be replaced | `get_records` becomes self-contained after `get_record` retirement. |
| Leave request count and total response size unbounded. | stale and must be replaced | Request count is capped at 20. Fixture and implementation limits beyond the public contract are deferred. |

### Affected-file manifest

| file | disposition | reason and later owner |
|---|---|---|
| `tools/list-records.md` | rewrite in T02 | Replace broad optional filters, exact ID, ranges, mixed-kind ordering, path tie-breaks, and full metadata output. |
| `tools/get-record.md` | rewrite as retired public operation in T03 | Preserve an explicit retirement marker while removing the operation from the public tool surface. |
| `tools/get-records.md` | rewrite in T03 | Define sole exact retrieval, 1-20 refs, ordered deduplication, successful-record-only partial success, body behavior, and path hiding. |
| `tools/overview.md` | rewrite affected sections in T04 | Remove `get_record`, stale ranges, path-bearing common response examples, and duplicated retrieval diagnostics. Keep the overview navigation-first. |
| `mvp-scope.md` | rewrite affected sections in T04 | Remove `get_record`, stale P0 language, YAML source claims, and obsolete `suggest_next_record` references without redesigning authoring phases. |
| `responsibility-boundary.md` | narrow in T04 | Remove the stale next-record suggestion and synchronize the query/retrieval and filesystem boundaries. |
| `schema/fields.md` | consume and recheck in T04 | Keep semantic field ownership unchanged. Add only operation pointers if T02/T03 require them. |
| `schema/record-model.md` | consume and recheck in T04 | Keep internal provenance and index-state semantics unchanged. Synchronize public representation pointers only if required. |
| `overview.md` | recheck in T04; no change currently required | The current overview already delegates normal list and exact-retrieval representation to W004. |
| `schema/record-source.md` | input only; no W004 edit planned | W003 already limits headings and body to source availability and delegates public inclusion to W004. |
| `schema/diagnostics.md` | stale retrieval section delegated to W006 | Current `record_not_found` error and duplicate-info taxonomy conflicts with the accepted warning boundary. W004 must not define the replacement taxonomy. |
| `tools/validate-records.md` | stale range coupling delegated to W006 | Validation currently imports `list_records` range behavior. W006 owns validation request semantics and diagnostic response changes. |
| `resolver.md` | stale `get_record` reference delegated to W005 | Resolver lookup-source wording must follow the retired tool surface without changing W004 exact retrieval. |
| `schema/authoring-transaction-schema.md` | pointer-only follow-up outside T01 | The exclusion list names `get_record`. T04 may perform a mechanical public-tool pointer update only if it does not alter authoring semantics. |

### Retained claims

- Exact retrieval remains distinct from search, filtering, validation, and reference resolution.
- `get_records` keeps request-wide body inclusion.
- Returned body content remains verbatim.
- Batch retrieval keeps ordered first-occurrence deduplication and partial success.
- Current and legacy lookups remain confined to separate indexes.
- Normal list and retrieval responses use canonical identity instead of physical path.

### Rewritten or removed claims

- Replace omitted or optional app, kind, and domain scope with required explicit scope.
- Remove spec and legacy records from normal listing.
- Remove exact ID and range filters from `list_records`.
- Remove mixed-kind and path-based tie-breaking.
- Remove full metadata objects from listing results.
- Remove public `get_record`.
- Remove `get_records.items[]` failure placeholders and `record: null` entries.
- Remove V01 spec examples and physical paths from normal retrieval examples.
- Remove diagnostic category and severity authority from operation overview text.

### Delegated scope

| concern | owner |
|---|---|
| Current-root discovery, parsing, identity, active-index state, and source retention | Completed `DRMCP-WORK-MCP-003` contract. |
| Current-first resolution order and configured legacy fallback orchestration | `DRMCP-WORK-MCP-005`. |
| Diagnostic category names, severity, shared fields, and source locations | `DRMCP-WORK-MCP-006`. |
| Validation request and execution semantics | `DRMCP-WORK-MCP-006`. |
| Current and legacy fixtures | `DRMCP-WORK-MCP-008`. |
| Parser, index, and MCP implementation | `DRMCP-WORK-MCP-009` and `DRMCP-WORK-MCP-010`. |
| Automated implementation tests | Implementation Work Items. |
| Authoring transactions and authoring response paths | Requirements and Work Items outside W004. |

### Deferred implementation detail

- Exact Go types and internal pagination strategy.
- Body-size and transport-size safeguards that do not change the public contract.
- Fixture file names and concrete test cases.
- Warning category strings, severity values, and shared diagnostic field names.
- Internal active-index and legacy-index lookup APIs.

### Accepted decision: invalid-but-addressable records in `list_records`

Decision: include every addressable record in normal listing.

| source state | compact listing behavior |
|---|---|
| All compact fields are available | Return `ref`, `title`, `status`, and `date` with parsed values. |
| `title`, `status`, or `date` is missing | Keep every compact field present. Return the missing value as `null`. |
| A compact field contains a parsed but invalid value | Return the parsed value unchanged. Do not replace it with `null` or a fabricated value. |

A missing compact field produces an operation warning trigger.
W006 owns the warning category, severity, shared fields, and source-location representation.

Reason: excluding an invalid-but-addressable record would hide the defect from normal discovery. Field omission would also make the compact result shape unstable. A fixed nullable projection preserves discoverability without fabricating values.

### Current verification result

- Required authority and boundary records read: complete.
- Minimum affected specs inspected: complete.
- Scoped stale-contract search: complete.
- Normative DRMCP spec changes: none.
- Workflow changes: hub Task and W004 moved to `in_progress`; T01 created and linked.
- Unresolved T01-level decisions: none.
- Scoped spec validation was not required because T01 changed no normative spec.
- Independent baseline review verdict: `PASS`.
- Independent review findings: none.
- Independent review advisories: none.
- T01 closure readiness: `ready`.
- T01 closed as `done` on 2026-06-27.
